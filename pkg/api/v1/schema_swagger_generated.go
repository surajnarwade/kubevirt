// Automatically generated by swagger-doc. DO NOT EDIT!

package v1

func (CloudInitNoCloudSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"":               "Represents a cloud-init nocloud user data source\nMore info: http://cloudinit.readthedocs.io/en/latest/topics/datasources/nocloud.html",
		"secretRef":      "UserDataSecretRef references a k8s secret that contains NoCloud userdata\n+ optional",
		"userDataBase64": "UserDataBase64 contains NoCloud cloud-init userdata as a base64 encoded string\n+ optional",
	}
}

func (DomainSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"resources": "Resources describes the Compute Resources required by this vm.",
		"cpu":       "CPU allow specified the detailed CPU topology inside the vm.\n+optional",
		"firmware":  "Firmware\n+optional",
		"clock":     "Clock sets the clock and timers of the vm.\n+optional",
		"features":  "Features like acpi, apic, hyperv\n+optional",
		"devices":   "Devices allows adding disks, network interfaces, ...",
	}
}

func (ResourceRequirements) SwaggerDoc() map[string]string {
	return map[string]string{
		"requests": "Requests is a description of the initial vm resources.\nValid resource keys are \"memory\" and \"cpu\".\n+optional",
	}
}

func (CPU) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "CPU allow specifying the CPU topology",
		"cores": "Cores specifies the number of cores inside the vm.\nMust be a value greater or equal 1.",
	}
}

func (Firmware) SwaggerDoc() map[string]string {
	return map[string]string{
		"uid": "UID reported by the vm bios\nDefaults to a random generated uid",
	}
}

func (Devices) SwaggerDoc() map[string]string {
	return map[string]string{
		"disks":    "Disks describes disks, cdroms, floppy and luns which are connected to the vm",
		"watchdog": "Watchdog describes a watchdog device which can be added to the vm",
	}
}

func (Disk) SwaggerDoc() map[string]string {
	return map[string]string{
		"name":       "Name is the device name",
		"volumeName": "Name of the volume which is referenced\nMust match the Name of a Volume.",
	}
}

func (DiskDevice) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "Represents the target of a volume to mount.\nOnly one of its members may be specified.",
		"disk":   "Attach a volume as a disk to the vm",
		"lun":    "Attach a volume as a LUN to the vm",
		"floppy": "Attach a volume as a floppy to the vm",
		"cdrom":  "Attach a volume as a cdrom to the vm",
	}
}

func (DiskTarget) SwaggerDoc() map[string]string {
	return map[string]string{
		"dev":      "Device indicates the \"logical\" device name. The actual device name\nspecified is not guaranteed to map to the device name in the guest OS. Treat\nit as a device ordering hint.",
		"readonly": "ReadOnly\nDefaults to false",
	}
}

func (LunTarget) SwaggerDoc() map[string]string {
	return map[string]string{
		"dev":      "Device indicates the \"logical\" device name. The actual device name\nspecified is not guaranteed to map to the device name in the guest OS. Treat\nit as a device ordering hint.",
		"readonly": "ReadOnly\nDefaults to false",
	}
}

func (FloppyTarget) SwaggerDoc() map[string]string {
	return map[string]string{
		"dev":      "Device indicates the \"logical\" device name. The actual device name\nspecified is not guaranteed to map to the device name in the guest OS. Treat\nit as a device ordering hint.",
		"readonly": "ReadOnly\nDefaults to false",
		"tray":     "Tray indicates if the tray of the device is open or closed.\nAllowed values are \"open\" and \"closed\"\nDefaults to closed\n+optional",
	}
}

func (CDRomTarget) SwaggerDoc() map[string]string {
	return map[string]string{
		"dev":      "Device indicates the \"logical\" device name. The actual device name\nspecified is not guaranteed to map to the device name in the guest OS. Treat\nit as a device ordering hint.",
		"readonly": "ReadOnly\nDefaults to true",
		"tray":     "Tray indicates if the tray of the device is open or closed.\nAllowed values are \"open\" and \"closed\"\nDefaults to closed\n+optional",
	}
}

func (Volume) SwaggerDoc() map[string]string {
	return map[string]string{
		"":     "Volume represents a named volume in a vm.",
		"name": "Volume's name.\nMust be a DNS_LABEL and unique within the vm.\nMore info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
	}
}

func (VolumeSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                      "Represents the source of a volume to mount.\nOnly one of its members may be specified.",
		"iscsi":                 "ISCSI represents an ISCSI Disk resource which is directly attached to the vm via qemu.\n+optional",
		"persistentVolumeClaim": "PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace.\nDirectly attached to the vm via qemu.\nMore info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims\n+optional",
		"cloudInitNoCloud":      "CloudInitNoCloud represents a cloud-init NoCloud user-data source.\nThe NoCloud data will be added as a disk to the vm. A proper cloud-init installation is required inside the guest.\nMore info: http://cloudinit.readthedocs.io/en/latest/topics/datasources/nocloud.html\n+optional",
		"registryDisk":          "RegistryDisk references a docker image, embedding a qcow or raw disk\nMore info: https://kubevirt.gitbooks.io/user-guide/registry-disk.html\n+optional",
	}
}

func (RegistryDiskSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "Represents a docker image with an embedded disk",
		"image": "Image is the name of the image with the embedded disk",
	}
}

func (ClockOffset) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "Exactly one of its members must be set.",
		"utc":      "UTC sets the guest clock to UTC on each boot. If an offset is specified,\nguest changes to the clock will be kept during reboots and are not reset.",
		"timezone": "Timezone sets the guest clock to the specified timezone.\nZone name follows the TZ environment variable format (e.g. 'America/New_York')",
	}
}

func (ClockOffsetUTC) SwaggerDoc() map[string]string {
	return map[string]string{
		"":              "UTC sets the guest clock to UTC on each boot.",
		"offsetSeconds": "OffsetSeconds specifies an offset in seconds, relative to UTC. If set,\nguest changes to the clock will be kept during reboots and not reset.",
	}
}

func (Clock) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "Represents the clock and timers of a vm",
		"timer": "Timer specifies whih timers are attached to the vm",
	}
}

func (Timer) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "Represents all available timers in a vm",
		"hpet":   "HPET (High Precision Event Timer) - multiple timers with periodic interrupts.",
		"kvm":    "KVM \t(KVM clock) - lets guests read the host’s wall clock time (paravirtualized). For linux guests.",
		"pit":    "PIT (Programmable Interval Timer) - a timer with periodic interrupts.",
		"rtc":    "RTC (Real Time Clock) - a continuously running timer with periodic interrupts.",
		"hyperv": "Hyperv (Hypervclock) - lets guests read the host’s wall clock time (paravirtualized). For windows guests.",
	}
}

func (RTCTimer) SwaggerDoc() map[string]string {
	return map[string]string{
		"tickPolicy": "TickPolicy determines what happens when QEMU misses a deadline for injecting a tick to the guest\nOne of \"delay\", \"catchup\"",
		"present":    "Enabled set to false makes sure that the machine type or a preset can't add the timer.\nDefaults to true\n+optional",
		"track":      "Track the guest or the wall clock",
	}
}

func (HPETTimer) SwaggerDoc() map[string]string {
	return map[string]string{
		"tickPolicy": "TickPolicy determines what happens when QEMU misses a deadline for injecting a tick to the guest\nOne of \"delay\", \"catchup\", \"merge\", \"discard\"",
		"present":    "Enabled set to false makes sure that the machine type or a preset can't add the timer.\nDefaults to true\n+optional",
	}
}

func (PITTimer) SwaggerDoc() map[string]string {
	return map[string]string{
		"tickPolicy": "TickPolicy determines what happens when QEMU misses a deadline for injecting a tick to the guest\nOne of \"delay\", \"catchup\", \"discard\"",
		"present":    "Enabled set to false makes sure that the machine type or a preset can't add the timer.\nDefaults to true\n+optional",
	}
}

func (KVMTimer) SwaggerDoc() map[string]string {
	return map[string]string{
		"present": "Enabled set to false makes sure that the machine type or a preset can't add the timer.\nDefaults to true\n+optional",
	}
}

func (HypervTimer) SwaggerDoc() map[string]string {
	return map[string]string{
		"present": "Enabled set to false makes sure that the machine type or a preset can't add the timer.\nDefaults to true\n+optional",
	}
}

func (Features) SwaggerDoc() map[string]string {
	return map[string]string{
		"acpi":   "ACPI enables/disables ACPI insidejsondata guest\nDefaults to enabled\n+optional",
		"apic":   "Defaults to the machine type setting\n+optional",
		"hyperv": "Defaults to the machine type setting\n+optional",
	}
}

func (FeatureState) SwaggerDoc() map[string]string {
	return map[string]string{
		"":        "Represents if a feature is enabled or disabled",
		"enabled": "Enabled determines if the feature should be enabled or disabled on the guest\nDefaults to true\n+optional",
	}
}

func (FeatureAPIC) SwaggerDoc() map[string]string {
	return map[string]string{
		"enabled":        "Enabled determines if the feature should be enabled or disabled on the guest\nDefaults to true\n+optional",
		"endOfInterrupt": "EndOfInterrupt enables the end of interrupt notification in the guest\nDefaults to false\n+optional",
	}
}

func (FeatureSpinlocks) SwaggerDoc() map[string]string {
	return map[string]string{
		"enabled":   "Enabled determines if the feature should be enabled or disabled on the guest\nDefaults to true\n+optional",
		"spinlocks": "Retries indicates the number of retries\nMust be a value greater or equal 4096\nDefaults to 4096\n+optional",
	}
}

func (FeatureVendorID) SwaggerDoc() map[string]string {
	return map[string]string{
		"enabled":  "Enabled determines if the feature should be enabled or disabled on the guest\nDefaults to true\n+optional",
		"vendorid": "VendorID sets the hypervisor vendor id, visible to the vm\nString up to twelve characters",
	}
}

func (FeatureHyperv) SwaggerDoc() map[string]string {
	return map[string]string{
		"":           "Hyperv specific features",
		"relaxed":    "Relaxed relaxes constraints on timer\nDefaults to the machine type setting\n+optional",
		"vapic":      "VAPIC indicates weather virtual APIC is enabled\nDefaults to the machine type setting\n+optional",
		"spinlocks":  "Spinlocks indicates if spinlocks should be made available to the guest\n+optional",
		"vpindex":    "VPIndex enables the Virtual Processor Index to help windows identifying virtual processors\nDefaults to the machine type setting\n+optional",
		"runtime":    "Runtime\nDefaults to the machine type setting\n+optional",
		"synic":      "SyNIC enable Synthetic Interrupt Controller\nDefaults to the machine type setting\n+optional",
		"synictimer": "SyNICTimer enable Synthetic Interrupt Controller timer\nDefaults to the machine type setting\n+optional",
		"reset":      "Reset enables Hyperv reboot/reset for the vm\nDefaults to the machine type setting\n+optional",
		"vendorid":   "VendorID allows setting the hypervisor vendor id\nDefaults to the machine type setting\n+optional",
	}
}

func (Watchdog) SwaggerDoc() map[string]string {
	return map[string]string{
		"":     "Named watchdog device",
		"name": "Name of the watchdog",
	}
}

func (WatchdogDevice) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "Hardware watchdog device\nExactly one of its members must be set.",
		"i6300esb": "i6300esb watchdog device\n+optional",
	}
}

func (I6300ESBWatchdog) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "i6300esb watchdog device",
		"action": "The action to take. Valid values are poweroff, reset, shutdown.\nDefaults to reset",
	}
}
