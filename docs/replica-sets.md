# VirtualMachineReplicaSet

## Overview

In order to allow scaling up or down similar VMs, add a way to specify
`VirtualMachine` templates and the amount of replicas required, to let the
runtime create these `VirtualMachine`s.

## Use-cases

### Ephemeral VMs

Scaling ephemeral VMs which only have read-only mounts, or work with a backing
store, to keep temporary data, which can be deleted after the VM gets
undefined.

## Implementation

A new object `VirtualMachineReplicaSet`, backed with a controller will be
implemented:

```yaml
apiVersion: kubevirt.io/v1alpha1
kind: VirtualMachineReplicaSet
metadata:
  name: myreplicaset
spec:
  replicas: 3
  selector:
    matchLabels:
      mylabel: mylabel
  template:
    metadata:
      name: test
      labels:
        mylabel: mylabel
    spec:
      domain:
        devices:
      [...]
```

`spec.template` is equal to a `VirtualMachineSpec`. `spec.replicas` specifies
how many instances should be created out of `spec.template`. `spec.selector`
contains selectors, which need to match `spec.template.metadata.labels`.

The status looks like this:

```yaml
status:
  conditions: null
  replicas: 3
  readyReplicas : 2
```
In case of a scaling error, a `ReplicaFailure` condition is added to the
`status.conditions`. Further it shows the number of `VirtualMachine`s which
are in a non-final state and which match `spec.selector` in the
`status.replicas` field.  `status.readyReplicas` indicates how many of these
replicas meet the ready condition.

*Note* that at the moment when writing this proposal, there exist no
readiness checks for VirtualMachines in Kubevirt. Therefore a `VirtualMachine` is
considered to be ready, when reported by virt-handler as running or migrating.

In case of a delete failure:

```yaml
status:
  conditions:
  - type: "ReplicaFailure"
    status: True
    reason: "FailureDelete"
    message: "no permission to delete VMs"
    lastTransmissionTime: "..."
  replicas: 4
  readyReplicas: 3
```

In case of a create failure:

```yaml
status:
  conditions:
  - type: "ReplicaFailure"
    status: True
    reason: "FailureCreate"
    message: "no permission to create VMs"
    lastTransmissionTime: "..."
  replicas: 2
  readyReplicas: 3
```

### Guarantees

The VirtualMachineReplicaSet  does **not** guarantee that there will never be
more than the wanted replicas active in the cluster. Based on readiness checks,
unknown VirtualMachine states and graceful deletes, it might decide to already
create new replicas in advance, to make sure that the amount of ready replicas
stays close to the expected replica count.

### Chosen Implementation

The implementation of the VirtualMachineReplicaSet is a reimplementation of the
ReplicaSet for Pods in Kubernetes. It does not wrap around a Kubernetes
ReplicaSet.


There two *hard* reasons why one was chose over the other:

 1. Allow possible VirtualMachine related optimizations by implementing a
    custom deletion order. When scaling down, it can make sense to take
    migrating VirtualMachines into account. A possible deletion order would be:
    not ready VMs, migrating VMs, ready VMs. This would be impossible to
    represent by wrapping around the ReplicaSet.
 2. Make Live Migrations an orthogonal feature for every VirtualMachine
    controller. If a controller does not have to care about the
    VirtualMachine/Pod relationship, it does not have to care in any special
    way about migrations. By wrapping around existing controllers, we will
    always have to hide migration target Pods from the controller, until the
    VirtualMachine was migrated, and then have to make it somehow visible
    again. This solution leads to tricky synchronization problems which may not
    even be completely solvable.

There are several *soft* reasons why one was chosen over the other:

 1. Provide a clear VirtualMachine abstraction via our VirtualMachine
    object. This allows domain specific modelling of controllers.
 2. Hide the whole relationship between VirtualMachines and Pods behind on
    controller, so that others don't have to care about it.
 3. The actual business-logic of scaling the VirtualMachines is very easy to
    understand and simple to implement. We get almost all the necessary
    infrastructure code from k8s libraries. The infrastructure code has a
    similar complexity (entity listeners, creating/deleting resources,
    ... ) for both implementations.  For different types of controllers this
    argument may not apply (e.g.  DaemonSet equivalent for VirtualMachines).
 4. Because of the simplicity of the business-logic, the equal complexity
    regarding to the controller infrastructure and the fact that there exists a
    ReplicaSet reference implementation from which we can learn/copy, it seems
    beneficial to really express our domain and avoid design errors, by
    implementing another complex flow around the existing ReplicaSet.

### Milestones

 * Basic functionality
 * Support label changes
 * Define a well known scale-down order
 * Support graceful delete [1]
 * Support controller references [2]
 * Support adopting orphaned Pods [2]

The basic functionality includes scaling up, down and reporting errors if
scaling does not work. In this stage it is the full responsibility of the user
to attach labels to the VMs in a way, so that selectors of multiple
VirtualMachineReplicaSets don't overlap.

### References

1. https://kubernetes.io/docs/tasks/run-application/force-delete-stateful-set-pod/#delete-pods
2. https://github.com/kubernetes/community/blob/58b1c30d95719749068497ba35dfe4c64b21aa72/contributors/design-proposals/api-machinery/controller-ref.md
