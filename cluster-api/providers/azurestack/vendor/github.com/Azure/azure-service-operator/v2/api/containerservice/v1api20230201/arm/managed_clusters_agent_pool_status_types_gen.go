// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type ManagedClustersAgentPool_STATUS struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of an agent pool.
	Properties *ManagedClusterAgentPoolProfileProperties_STATUS `json:"properties,omitempty"`

	// Type: Resource type
	Type *string `json:"type,omitempty"`
}

// Properties for the container service agent pool profile.
type ManagedClusterAgentPoolProfileProperties_STATUS struct {
	// AvailabilityZones: The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType
	// property is 'VirtualMachineScaleSets'.
	AvailabilityZones []string `json:"availabilityZones"`

	// Count: Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive)
	// for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.
	Count *int `json:"count,omitempty"`

	// CreationData: CreationData to be used to specify the source Snapshot ID if the node pool will be created/upgraded using
	// a snapshot.
	CreationData *CreationData_STATUS `json:"creationData,omitempty"`

	// CurrentOrchestratorVersion: If orchestratorVersion is a fully specified version <major.minor.patch>, this field will be
	// exactly equal to it. If orchestratorVersion is <major.minor>, this field will contain the full <major.minor.patch>
	// version being used.
	CurrentOrchestratorVersion *string `json:"currentOrchestratorVersion,omitempty"`

	// EnableAutoScaling: Whether to enable auto-scaler
	EnableAutoScaling *bool `json:"enableAutoScaling,omitempty"`

	// EnableEncryptionAtHost: This is only supported on certain VM sizes and in certain Azure regions. For more information,
	// see: https://docs.microsoft.com/azure/aks/enable-host-encryption
	EnableEncryptionAtHost *bool `json:"enableEncryptionAtHost,omitempty"`

	// EnableFIPS: See [Add a FIPS-enabled node
	// pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview) for more
	// details.
	EnableFIPS *bool `json:"enableFIPS,omitempty"`

	// EnableNodePublicIP: Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses.
	// A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine
	// to minimize hops. For more information see [assigning a public IP per
	// node](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools). The
	// default is false.
	EnableNodePublicIP *bool `json:"enableNodePublicIP,omitempty"`

	// EnableUltraSSD: Whether to enable UltraSSD
	EnableUltraSSD *bool `json:"enableUltraSSD,omitempty"`

	// GpuInstanceProfile: GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.
	GpuInstanceProfile *GPUInstanceProfile_STATUS `json:"gpuInstanceProfile,omitempty"`

	// HostGroupID: This is of the form:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/hostGroups/{hostGroupName}.
	// For more information see [Azure dedicated hosts](https://docs.microsoft.com/azure/virtual-machines/dedicated-hosts).
	HostGroupID *string `json:"hostGroupID,omitempty"`

	// KubeletConfig: The Kubelet configuration on the agent pool nodes.
	KubeletConfig *KubeletConfig_STATUS `json:"kubeletConfig,omitempty"`

	// KubeletDiskType: Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral
	// storage.
	KubeletDiskType *KubeletDiskType_STATUS `json:"kubeletDiskType,omitempty"`

	// LinuxOSConfig: The OS configuration of Linux agent nodes.
	LinuxOSConfig *LinuxOSConfig_STATUS `json:"linuxOSConfig,omitempty"`

	// MaxCount: The maximum number of nodes for auto-scaling
	MaxCount *int `json:"maxCount,omitempty"`

	// MaxPods: The maximum number of pods that can run on a node.
	MaxPods *int `json:"maxPods,omitempty"`

	// MinCount: The minimum number of nodes for auto-scaling
	MinCount *int `json:"minCount,omitempty"`

	// Mode: A cluster must have at least one 'System' Agent Pool at all times. For additional information on agent pool
	// restrictions  and best practices, see: https://docs.microsoft.com/azure/aks/use-system-pools
	Mode *AgentPoolMode_STATUS `json:"mode,omitempty"`

	// NodeImageVersion: The version of node image
	NodeImageVersion *string `json:"nodeImageVersion,omitempty"`

	// NodeLabels: The node labels to be persisted across all nodes in agent pool.
	NodeLabels map[string]string `json:"nodeLabels"`

	// NodePublicIPPrefixID: This is of the form:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}
	NodePublicIPPrefixID *string `json:"nodePublicIPPrefixID,omitempty"`

	// NodeTaints: The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
	NodeTaints []string `json:"nodeTaints"`

	// OrchestratorVersion: Both patch version <major.minor.patch> (e.g. 1.20.13) and <major.minor> (e.g. 1.20) are supported.
	// When <major.minor> is specified, the latest supported GA patch version is chosen automatically. Updating the cluster
	// with the same <major.minor> once it has been created (e.g. 1.14.x -> 1.14) will not trigger an upgrade, even if a newer
	// patch version is available. As a best practice, you should upgrade all node pools in an AKS cluster to the same
	// Kubernetes version. The node pool version must have the same major version as the control plane. The node pool minor
	// version must be within two minor versions of the control plane version. The node pool version cannot be greater than the
	// control plane version. For more information see [upgrading a node
	// pool](https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool).
	OrchestratorVersion *string `json:"orchestratorVersion,omitempty"`
	OsDiskSizeGB        *int    `json:"osDiskSizeGB,omitempty"`

	// OsDiskType: The default is 'Ephemeral' if the VM supports it and has a cache disk larger than the requested
	// OSDiskSizeGB. Otherwise,  defaults to 'Managed'. May not be changed after creation. For more information see [Ephemeral
	// OS](https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os).
	OsDiskType *OSDiskType_STATUS `json:"osDiskType,omitempty"`

	// OsSKU: Specifies the OS SKU used by the agent pool. The default is Ubuntu if OSType is Linux. The default is Windows2019
	// when  Kubernetes <= 1.24 or Windows2022 when Kubernetes >= 1.25 if OSType is Windows.
	OsSKU *OSSKU_STATUS `json:"osSKU,omitempty"`

	// OsType: The operating system type. The default is Linux.
	OsType *OSType_STATUS `json:"osType,omitempty"`

	// PodSubnetID: If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is
	// of the form:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}
	PodSubnetID *string `json:"podSubnetID,omitempty"`

	// PowerState: When an Agent Pool is first created it is initially Running. The Agent Pool can be stopped by setting this
	// field to Stopped. A stopped Agent Pool stops all of its VMs and does not accrue billing charges. An Agent Pool can only
	// be stopped if it is Running and provisioning state is Succeeded
	PowerState *PowerState_STATUS `json:"powerState,omitempty"`

	// ProvisioningState: The current deployment or provisioning state.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// ProximityPlacementGroupID: The ID for Proximity Placement Group.
	ProximityPlacementGroupID *string `json:"proximityPlacementGroupID,omitempty"`

	// ScaleDownMode: This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.
	ScaleDownMode *ScaleDownMode_STATUS `json:"scaleDownMode,omitempty"`

	// ScaleSetEvictionPolicy: This cannot be specified unless the scaleSetPriority is 'Spot'. If not specified, the default is
	// 'Delete'.
	ScaleSetEvictionPolicy *ScaleSetEvictionPolicy_STATUS `json:"scaleSetEvictionPolicy,omitempty"`

	// ScaleSetPriority: The Virtual Machine Scale Set priority. If not specified, the default is 'Regular'.
	ScaleSetPriority *ScaleSetPriority_STATUS `json:"scaleSetPriority,omitempty"`

	// SpotMaxPrice: Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any
	// on-demand price. For more details on spot pricing, see [spot VMs
	// pricing](https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing)
	SpotMaxPrice *float64 `json:"spotMaxPrice,omitempty"`

	// Tags: The tags to be persisted on the agent pool virtual machine scale set.
	Tags map[string]string `json:"tags"`

	// Type: The type of Agent Pool.
	Type *AgentPoolType_STATUS `json:"type,omitempty"`

	// UpgradeSettings: Settings for upgrading the agentpool
	UpgradeSettings *AgentPoolUpgradeSettings_STATUS `json:"upgradeSettings,omitempty"`

	// VmSize: VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods
	// might fail to run correctly. For more details on restricted VM sizes, see:
	// https://docs.microsoft.com/azure/aks/quotas-skus-regions
	VmSize *string `json:"vmSize,omitempty"`

	// VnetSubnetID: If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified,
	// this applies to nodes and pods, otherwise it applies to just nodes. This is of the form:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}
	VnetSubnetID *string `json:"vnetSubnetID,omitempty"`

	// WorkloadRuntime: Determines the type of workload a node can run.
	WorkloadRuntime *WorkloadRuntime_STATUS `json:"workloadRuntime,omitempty"`
}

type AgentPoolMode_STATUS string

const (
	AgentPoolMode_STATUS_System = AgentPoolMode_STATUS("System")
	AgentPoolMode_STATUS_User   = AgentPoolMode_STATUS("User")
)

// Mapping from string to AgentPoolMode_STATUS
var agentPoolMode_STATUS_Values = map[string]AgentPoolMode_STATUS{
	"system": AgentPoolMode_STATUS_System,
	"user":   AgentPoolMode_STATUS_User,
}

type AgentPoolType_STATUS string

const (
	AgentPoolType_STATUS_AvailabilitySet         = AgentPoolType_STATUS("AvailabilitySet")
	AgentPoolType_STATUS_VirtualMachineScaleSets = AgentPoolType_STATUS("VirtualMachineScaleSets")
)

// Mapping from string to AgentPoolType_STATUS
var agentPoolType_STATUS_Values = map[string]AgentPoolType_STATUS{
	"availabilityset":         AgentPoolType_STATUS_AvailabilitySet,
	"virtualmachinescalesets": AgentPoolType_STATUS_VirtualMachineScaleSets,
}

type AgentPoolUpgradeSettings_STATUS struct {
	MaxSurge *string `json:"maxSurge,omitempty"`
}

type CreationData_STATUS struct {
	SourceResourceId *string `json:"sourceResourceId,omitempty"`
}

type GPUInstanceProfile_STATUS string

const (
	GPUInstanceProfile_STATUS_MIG1G = GPUInstanceProfile_STATUS("MIG1g")
	GPUInstanceProfile_STATUS_MIG2G = GPUInstanceProfile_STATUS("MIG2g")
	GPUInstanceProfile_STATUS_MIG3G = GPUInstanceProfile_STATUS("MIG3g")
	GPUInstanceProfile_STATUS_MIG4G = GPUInstanceProfile_STATUS("MIG4g")
	GPUInstanceProfile_STATUS_MIG7G = GPUInstanceProfile_STATUS("MIG7g")
)

// Mapping from string to GPUInstanceProfile_STATUS
var gPUInstanceProfile_STATUS_Values = map[string]GPUInstanceProfile_STATUS{
	"mig1g": GPUInstanceProfile_STATUS_MIG1G,
	"mig2g": GPUInstanceProfile_STATUS_MIG2G,
	"mig3g": GPUInstanceProfile_STATUS_MIG3G,
	"mig4g": GPUInstanceProfile_STATUS_MIG4G,
	"mig7g": GPUInstanceProfile_STATUS_MIG7G,
}

type KubeletConfig_STATUS struct {
	AllowedUnsafeSysctls  []string `json:"allowedUnsafeSysctls"`
	ContainerLogMaxFiles  *int     `json:"containerLogMaxFiles,omitempty"`
	ContainerLogMaxSizeMB *int     `json:"containerLogMaxSizeMB,omitempty"`
	CpuCfsQuota           *bool    `json:"cpuCfsQuota,omitempty"`
	CpuCfsQuotaPeriod     *string  `json:"cpuCfsQuotaPeriod,omitempty"`
	CpuManagerPolicy      *string  `json:"cpuManagerPolicy,omitempty"`
	FailSwapOn            *bool    `json:"failSwapOn,omitempty"`
	ImageGcHighThreshold  *int     `json:"imageGcHighThreshold,omitempty"`
	ImageGcLowThreshold   *int     `json:"imageGcLowThreshold,omitempty"`
	PodMaxPids            *int     `json:"podMaxPids,omitempty"`
	TopologyManagerPolicy *string  `json:"topologyManagerPolicy,omitempty"`
}

type KubeletDiskType_STATUS string

const (
	KubeletDiskType_STATUS_OS        = KubeletDiskType_STATUS("OS")
	KubeletDiskType_STATUS_Temporary = KubeletDiskType_STATUS("Temporary")
)

// Mapping from string to KubeletDiskType_STATUS
var kubeletDiskType_STATUS_Values = map[string]KubeletDiskType_STATUS{
	"os":        KubeletDiskType_STATUS_OS,
	"temporary": KubeletDiskType_STATUS_Temporary,
}

type LinuxOSConfig_STATUS struct {
	SwapFileSizeMB             *int                 `json:"swapFileSizeMB,omitempty"`
	Sysctls                    *SysctlConfig_STATUS `json:"sysctls,omitempty"`
	TransparentHugePageDefrag  *string              `json:"transparentHugePageDefrag,omitempty"`
	TransparentHugePageEnabled *string              `json:"transparentHugePageEnabled,omitempty"`
}

type OSDiskType_STATUS string

const (
	OSDiskType_STATUS_Ephemeral = OSDiskType_STATUS("Ephemeral")
	OSDiskType_STATUS_Managed   = OSDiskType_STATUS("Managed")
)

// Mapping from string to OSDiskType_STATUS
var oSDiskType_STATUS_Values = map[string]OSDiskType_STATUS{
	"ephemeral": OSDiskType_STATUS_Ephemeral,
	"managed":   OSDiskType_STATUS_Managed,
}

type OSSKU_STATUS string

const (
	OSSKU_STATUS_CBLMariner  = OSSKU_STATUS("CBLMariner")
	OSSKU_STATUS_Ubuntu      = OSSKU_STATUS("Ubuntu")
	OSSKU_STATUS_Windows2019 = OSSKU_STATUS("Windows2019")
	OSSKU_STATUS_Windows2022 = OSSKU_STATUS("Windows2022")
)

// Mapping from string to OSSKU_STATUS
var oSSKU_STATUS_Values = map[string]OSSKU_STATUS{
	"cblmariner":  OSSKU_STATUS_CBLMariner,
	"ubuntu":      OSSKU_STATUS_Ubuntu,
	"windows2019": OSSKU_STATUS_Windows2019,
	"windows2022": OSSKU_STATUS_Windows2022,
}

type OSType_STATUS string

const (
	OSType_STATUS_Linux   = OSType_STATUS("Linux")
	OSType_STATUS_Windows = OSType_STATUS("Windows")
)

// Mapping from string to OSType_STATUS
var oSType_STATUS_Values = map[string]OSType_STATUS{
	"linux":   OSType_STATUS_Linux,
	"windows": OSType_STATUS_Windows,
}

type ScaleDownMode_STATUS string

const (
	ScaleDownMode_STATUS_Deallocate = ScaleDownMode_STATUS("Deallocate")
	ScaleDownMode_STATUS_Delete     = ScaleDownMode_STATUS("Delete")
)

// Mapping from string to ScaleDownMode_STATUS
var scaleDownMode_STATUS_Values = map[string]ScaleDownMode_STATUS{
	"deallocate": ScaleDownMode_STATUS_Deallocate,
	"delete":     ScaleDownMode_STATUS_Delete,
}

type ScaleSetEvictionPolicy_STATUS string

const (
	ScaleSetEvictionPolicy_STATUS_Deallocate = ScaleSetEvictionPolicy_STATUS("Deallocate")
	ScaleSetEvictionPolicy_STATUS_Delete     = ScaleSetEvictionPolicy_STATUS("Delete")
)

// Mapping from string to ScaleSetEvictionPolicy_STATUS
var scaleSetEvictionPolicy_STATUS_Values = map[string]ScaleSetEvictionPolicy_STATUS{
	"deallocate": ScaleSetEvictionPolicy_STATUS_Deallocate,
	"delete":     ScaleSetEvictionPolicy_STATUS_Delete,
}

type ScaleSetPriority_STATUS string

const (
	ScaleSetPriority_STATUS_Regular = ScaleSetPriority_STATUS("Regular")
	ScaleSetPriority_STATUS_Spot    = ScaleSetPriority_STATUS("Spot")
)

// Mapping from string to ScaleSetPriority_STATUS
var scaleSetPriority_STATUS_Values = map[string]ScaleSetPriority_STATUS{
	"regular": ScaleSetPriority_STATUS_Regular,
	"spot":    ScaleSetPriority_STATUS_Spot,
}

type WorkloadRuntime_STATUS string

const (
	WorkloadRuntime_STATUS_OCIContainer = WorkloadRuntime_STATUS("OCIContainer")
	WorkloadRuntime_STATUS_WasmWasi     = WorkloadRuntime_STATUS("WasmWasi")
)

// Mapping from string to WorkloadRuntime_STATUS
var workloadRuntime_STATUS_Values = map[string]WorkloadRuntime_STATUS{
	"ocicontainer": WorkloadRuntime_STATUS_OCIContainer,
	"wasmwasi":     WorkloadRuntime_STATUS_WasmWasi,
}

type SysctlConfig_STATUS struct {
	FsAioMaxNr                     *int    `json:"fsAioMaxNr,omitempty"`
	FsFileMax                      *int    `json:"fsFileMax,omitempty"`
	FsInotifyMaxUserWatches        *int    `json:"fsInotifyMaxUserWatches,omitempty"`
	FsNrOpen                       *int    `json:"fsNrOpen,omitempty"`
	KernelThreadsMax               *int    `json:"kernelThreadsMax,omitempty"`
	NetCoreNetdevMaxBacklog        *int    `json:"netCoreNetdevMaxBacklog,omitempty"`
	NetCoreOptmemMax               *int    `json:"netCoreOptmemMax,omitempty"`
	NetCoreRmemDefault             *int    `json:"netCoreRmemDefault,omitempty"`
	NetCoreRmemMax                 *int    `json:"netCoreRmemMax,omitempty"`
	NetCoreSomaxconn               *int    `json:"netCoreSomaxconn,omitempty"`
	NetCoreWmemDefault             *int    `json:"netCoreWmemDefault,omitempty"`
	NetCoreWmemMax                 *int    `json:"netCoreWmemMax,omitempty"`
	NetIpv4IpLocalPortRange        *string `json:"netIpv4IpLocalPortRange,omitempty"`
	NetIpv4NeighDefaultGcThresh1   *int    `json:"netIpv4NeighDefaultGcThresh1,omitempty"`
	NetIpv4NeighDefaultGcThresh2   *int    `json:"netIpv4NeighDefaultGcThresh2,omitempty"`
	NetIpv4NeighDefaultGcThresh3   *int    `json:"netIpv4NeighDefaultGcThresh3,omitempty"`
	NetIpv4TcpFinTimeout           *int    `json:"netIpv4TcpFinTimeout,omitempty"`
	NetIpv4TcpKeepaliveProbes      *int    `json:"netIpv4TcpKeepaliveProbes,omitempty"`
	NetIpv4TcpKeepaliveTime        *int    `json:"netIpv4TcpKeepaliveTime,omitempty"`
	NetIpv4TcpMaxSynBacklog        *int    `json:"netIpv4TcpMaxSynBacklog,omitempty"`
	NetIpv4TcpMaxTwBuckets         *int    `json:"netIpv4TcpMaxTwBuckets,omitempty"`
	NetIpv4TcpTwReuse              *bool   `json:"netIpv4TcpTwReuse,omitempty"`
	NetIpv4TcpkeepaliveIntvl       *int    `json:"netIpv4TcpkeepaliveIntvl,omitempty"`
	NetNetfilterNfConntrackBuckets *int    `json:"netNetfilterNfConntrackBuckets,omitempty"`
	NetNetfilterNfConntrackMax     *int    `json:"netNetfilterNfConntrackMax,omitempty"`
	VmMaxMapCount                  *int    `json:"vmMaxMapCount,omitempty"`
	VmSwappiness                   *int    `json:"vmSwappiness,omitempty"`
	VmVfsCachePressure             *int    `json:"vmVfsCachePressure,omitempty"`
}
