// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type VirtualNetworksVirtualNetworkPeering_STATUS struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the virtual network peering.
	Properties *VirtualNetworkPeeringPropertiesFormat_STATUS `json:"properties,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// Properties of the virtual network peering.
type VirtualNetworkPeeringPropertiesFormat_STATUS struct {
	// AllowForwardedTraffic: Whether the forwarded traffic from the VMs in the local virtual network will be
	// allowed/disallowed in remote virtual network.
	AllowForwardedTraffic *bool `json:"allowForwardedTraffic,omitempty"`

	// AllowGatewayTransit: If gateway links can be used in remote virtual networking to link to this virtual network.
	AllowGatewayTransit *bool `json:"allowGatewayTransit,omitempty"`

	// AllowVirtualNetworkAccess: Whether the VMs in the local virtual network space would be able to access the VMs in remote
	// virtual network space.
	AllowVirtualNetworkAccess *bool `json:"allowVirtualNetworkAccess,omitempty"`

	// DoNotVerifyRemoteGateways: If we need to verify the provisioning state of the remote gateway.
	DoNotVerifyRemoteGateways *bool `json:"doNotVerifyRemoteGateways,omitempty"`

	// PeeringState: The status of the virtual network peering.
	PeeringState *VirtualNetworkPeeringPropertiesFormat_PeeringState_STATUS `json:"peeringState,omitempty"`

	// ProvisioningState: The provisioning state of the virtual network peering resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// RemoteAddressSpace: The reference to the remote virtual network address space.
	RemoteAddressSpace *AddressSpace_STATUS `json:"remoteAddressSpace,omitempty"`

	// RemoteBgpCommunities: The reference to the remote virtual network's Bgp Communities.
	RemoteBgpCommunities *VirtualNetworkBgpCommunities_STATUS `json:"remoteBgpCommunities,omitempty"`

	// RemoteVirtualNetwork: The reference to the remote virtual network. The remote virtual network can be in the same or
	// different region (preview). See here to register for the preview and learn more
	// (https://docs.microsoft.com/en-us/azure/virtual-network/virtual-network-create-peering).
	RemoteVirtualNetwork *SubResource_STATUS `json:"remoteVirtualNetwork,omitempty"`

	// ResourceGuid: The resourceGuid property of the Virtual Network peering resource.
	ResourceGuid *string `json:"resourceGuid,omitempty"`

	// UseRemoteGateways: If remote gateways can be used on this virtual network. If the flag is set to true, and
	// allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for
	// transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a
	// gateway.
	UseRemoteGateways *bool `json:"useRemoteGateways,omitempty"`
}

type VirtualNetworkPeeringPropertiesFormat_PeeringState_STATUS string

const (
	VirtualNetworkPeeringPropertiesFormat_PeeringState_STATUS_Connected    = VirtualNetworkPeeringPropertiesFormat_PeeringState_STATUS("Connected")
	VirtualNetworkPeeringPropertiesFormat_PeeringState_STATUS_Disconnected = VirtualNetworkPeeringPropertiesFormat_PeeringState_STATUS("Disconnected")
	VirtualNetworkPeeringPropertiesFormat_PeeringState_STATUS_Initiated    = VirtualNetworkPeeringPropertiesFormat_PeeringState_STATUS("Initiated")
)

// Mapping from string to VirtualNetworkPeeringPropertiesFormat_PeeringState_STATUS
var virtualNetworkPeeringPropertiesFormat_PeeringState_STATUS_Values = map[string]VirtualNetworkPeeringPropertiesFormat_PeeringState_STATUS{
	"connected":    VirtualNetworkPeeringPropertiesFormat_PeeringState_STATUS_Connected,
	"disconnected": VirtualNetworkPeeringPropertiesFormat_PeeringState_STATUS_Disconnected,
	"initiated":    VirtualNetworkPeeringPropertiesFormat_PeeringState_STATUS_Initiated,
}
