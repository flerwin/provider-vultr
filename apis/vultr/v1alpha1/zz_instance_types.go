/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BackupsScheduleObservation struct {
}

type BackupsScheduleParameters struct {

	// Day of month to run. Use values between 1 and 28.
	// +kubebuilder:validation:Optional
	Dom *float64 `json:"dom,omitempty" tf:"dom,omitempty"`

	// Day of week to run. 1 = Sunday, 2 = Monday, 3 = Tuesday, 4 = Wednesday, 5 = Thursday, 6 = Friday, 7 = Saturday
	// +kubebuilder:validation:Optional
	Dow *float64 `json:"dow,omitempty" tf:"dow,omitempty"`

	// Hour of day to run in UTC.
	// +kubebuilder:validation:Optional
	Hour *float64 `json:"hour,omitempty" tf:"hour,omitempty"`

	// Type of backup schedule Possible values are daily, weekly, monthly, daily_alt_even, or daily_alt_odd.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type InstanceObservation struct {

	// The server's allowed bandwidth usage in GB.
	AllowedBandwidth *float64 `json:"allowedBandwidth,omitempty" tf:"allowed_bandwidth,omitempty"`

	// The date the server was added to your Vultr account.
	DateCreated *string `json:"dateCreated,omitempty" tf:"date_created,omitempty"`

	// The description of the disk(s) on the server.
	Disk *float64 `json:"disk,omitempty" tf:"disk,omitempty"`

	// Array of which features are enabled.
	Features []*string `json:"features,omitempty" tf:"features,omitempty"`

	// The server's IPv4 gateway.
	GatewayV4 *string `json:"gatewayV4,omitempty" tf:"gateway_v4,omitempty"`

	// ID of the server.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The server's internal IP address.
	InternalIP *string `json:"internalIp,omitempty" tf:"internal_ip,omitempty"`

	// The server's current KVM URL. This URL will change periodically. It is not advised to cache this value.
	Kvm *string `json:"kvm,omitempty" tf:"kvm,omitempty"`

	// The server's main IP address.
	MainIP *string `json:"mainIp,omitempty" tf:"main_ip,omitempty"`

	// The server's IPv4 netmask.
	NetmaskV4 *string `json:"netmaskV4,omitempty" tf:"netmask_v4,omitempty"`

	// The string description of the operating system installed on the server.
	Os *string `json:"os,omitempty" tf:"os,omitempty"`

	// Whether the server is powered on or not.
	PowerStatus *string `json:"powerStatus,omitempty" tf:"power_status,omitempty"`

	// The amount of memory available on the server in MB.
	RAM *float64 `json:"ram,omitempty" tf:"ram,omitempty"`

	// A more detailed server status (none, locked, installingbooting, isomounting, ok).
	ServerStatus *string `json:"serverStatus,omitempty" tf:"server_status,omitempty"`

	// The status of the server's subscription.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The main IPv6 network address.
	V6MainIP *string `json:"v6MainIp,omitempty" tf:"v6_main_ip,omitempty"`

	// The IPv6 subnet.
	V6Network *string `json:"v6Network,omitempty" tf:"v6_network,omitempty"`

	// The IPv6 network size in bits.
	V6NetworkSize *float64 `json:"v6NetworkSize,omitempty" tf:"v6_network_size,omitempty"`

	// The number of virtual CPUs available on the server.
	VcpuCount *float64 `json:"vcpuCount,omitempty" tf:"vcpu_count,omitempty"`
}

type InstanceParameters struct {

	// Whether an activation email will be sent when the server is ready.
	// +kubebuilder:validation:Optional
	ActivationEmail *bool `json:"activationEmail,omitempty" tf:"activation_email,omitempty"`

	// The ID of the Vultr application to be installed on the server. See List Applications
	// +kubebuilder:validation:Optional
	AppID *float64 `json:"appId,omitempty" tf:"app_id,omitempty"`

	// Whether automatic backups will be enabled for this server (these have an extra charge associated with them). Values can be enabled or disabled.
	// +kubebuilder:validation:Optional
	Backups *string `json:"backups,omitempty" tf:"backups,omitempty"`

	// A block that defines the way backups should be scheduled. While this is an optional field if backups are enabled this field is mandatory. The configuration of a backups_schedule is listed below.
	// +kubebuilder:validation:Optional
	BackupsSchedule []BackupsScheduleParameters `json:"backupsSchedule,omitempty" tf:"backups_schedule,omitempty"`

	// Whether DDOS protection will be enabled on the server (there is an additional charge for this).
	// +kubebuilder:validation:Optional
	DdosProtection *bool `json:"ddosProtection,omitempty" tf:"ddos_protection,omitempty"`

	// Whether the server has IPv6 networking activated.
	// +kubebuilder:validation:Optional
	EnableIPv6 *bool `json:"enableIpv6,omitempty" tf:"enable_ipv6,omitempty"`

	// The ID of the firewall group to assign to the server.
	// +kubebuilder:validation:Optional
	FirewallGroupID *string `json:"firewallGroupId,omitempty" tf:"firewall_group_id,omitempty"`

	// The hostname to assign to the server.
	// The hostname of the instance. Updating the hostname will cause a force new. This behavior is in place to prevent accidental reinstalls. Issuing an update to the hostname on UI or API issues a reinstall of the OS.
	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// The ID of the Vultr marketplace application to be installed on the server. See List Applications Note marketplace applications are denoted by type: marketplace and you must use the image_id not the id.
	// +kubebuilder:validation:Optional
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// The ID of the ISO file to be installed on the server. See List ISO
	// +kubebuilder:validation:Optional
	IsoID *string `json:"isoId,omitempty" tf:"iso_id,omitempty"`

	// A label for the server.
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// The ID of the operating system to be installed on the server. See List OS
	// +kubebuilder:validation:Optional
	OsID *float64 `json:"osId,omitempty" tf:"os_id,omitempty"`

	// The ID of the plan that you want the instance to subscribe to. See List Plans
	// +kubebuilder:validation:Required
	Plan *string `json:"plan" tf:"plan,omitempty"`

	// (Deprecated: use vpc_ids instead) A list of private network IDs to be attached to the server.
	// +kubebuilder:validation:Optional
	PrivateNetworkIds []*string `json:"privateNetworkIds,omitempty" tf:"private_network_ids,omitempty"`

	// The ID of the region that the instance is to be created in. See List Regions
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// ID of the floating IP to use as the main IP of this server.
	// +kubebuilder:validation:Optional
	ReservedIPID *string `json:"reservedIpId,omitempty" tf:"reserved_ip_id,omitempty"`

	// A list of SSH key IDs to apply to the server on install (only valid for Linux/FreeBSD).
	// +kubebuilder:validation:Optional
	SSHKeyIds []*string `json:"sshKeyIds,omitempty" tf:"ssh_key_ids,omitempty"`

	// The ID of the startup script you want added to the server.
	// +kubebuilder:validation:Optional
	ScriptID *string `json:"scriptId,omitempty" tf:"script_id,omitempty"`

	// The ID of the Vultr snapshot that the server will restore for the initial installation. See List Snapshots
	// +kubebuilder:validation:Optional
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// (Deprecated: use tags instead)  The tag to assign to the server.
	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// A list of tags to apply to the instance.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Generic data store, which some provisioning tools and cloud operating systems use as a configuration file. It is generally consumed only once after an instance has been launched, but individual needs may vary.
	// +kubebuilder:validation:Optional
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`

	// A list of VPC IDs to be attached to the server.
	// +kubebuilder:validation:Optional
	VPCIds []*string `json:"vpcIds,omitempty" tf:"vpc_ids,omitempty"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API. Provides a Vultr instance resource. This can be used to create, read, modify, and delete instances on your Vultr account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vultr}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec"`
	Status            InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
