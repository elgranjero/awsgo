package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/ec2"
)

var fields_accept_address_transfer = []leanruntime.Field{
	{Name: "Address", Flag: "address", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_accept_capacity_reservation_billing_ownership = []leanruntime.Field{
	{Name: "CapacityReservationId", Flag: "capacity-reservation-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_accept_reserved_instances_exchange_quote = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ReservedInstanceIds", Flag: "reserved-instance-ids", Type: "[]string", Required: true},
	{Name: "TargetConfigurations", Flag: "target-configurations", Type: "[]types.TargetConfigurationRequest", Required: false},
}

var fields_accept_transit_gateway_multicast_domain_associations = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: false},
	{Name: "TransitGatewayAttachmentId", Flag: "transit-gateway-attachment-id", Type: "*string", Required: false},
	{Name: "TransitGatewayMulticastDomainId", Flag: "transit-gateway-multicast-domain-id", Type: "*string", Required: false},
}

var fields_accept_transit_gateway_peering_attachment = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TransitGatewayAttachmentId", Flag: "transit-gateway-attachment-id", Type: "*string", Required: true},
}

var fields_accept_transit_gateway_vpc_attachment = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TransitGatewayAttachmentId", Flag: "transit-gateway-attachment-id", Type: "*string", Required: true},
}

var fields_accept_vpc_endpoint_connections = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ServiceId", Flag: "service-id", Type: "*string", Required: true},
	{Name: "VpcEndpointIds", Flag: "vpc-endpoint-ids", Type: "[]string", Required: true},
}

var fields_accept_vpc_peering_connection = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VpcPeeringConnectionId", Flag: "vpc-peering-connection-id", Type: "*string", Required: true},
}

var fields_advertise_byoip_cidr = []leanruntime.Field{
	{Name: "Asn", Flag: "asn", Type: "*string", Required: false},
	{Name: "Cidr", Flag: "cidr", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "NetworkBorderGroup", Flag: "network-border-group", Type: "*string", Required: false},
}

var fields_allocate_address = []leanruntime.Field{
	{Name: "Address", Flag: "address", Type: "*string", Required: false},
	{Name: "CustomerOwnedIpv4Pool", Flag: "customer-owned-ipv4-pool", Type: "*string", Required: false},
	{Name: "Domain", Flag: "domain", Type: "types.DomainType", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamPoolId", Flag: "ipam-pool-id", Type: "*string", Required: false},
	{Name: "NetworkBorderGroup", Flag: "network-border-group", Type: "*string", Required: false},
	{Name: "PublicIpv4Pool", Flag: "public-ipv4-pool", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_allocate_hosts = []leanruntime.Field{
	{Name: "AssetIds", Flag: "asset-ids", Type: "[]string", Required: false},
	{Name: "AutoPlacement", Flag: "auto-placement", Type: "types.AutoPlacement", Required: false},
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "AvailabilityZoneId", Flag: "availability-zone-id", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "HostMaintenance", Flag: "host-maintenance", Type: "types.HostMaintenance", Required: false},
	{Name: "HostRecovery", Flag: "host-recovery", Type: "types.HostRecovery", Required: false},
	{Name: "InstanceFamily", Flag: "instance-family", Type: "*string", Required: false},
	{Name: "InstanceType", Flag: "instance-type", Type: "*string", Required: false},
	{Name: "OutpostArn", Flag: "outpost-arn", Type: "*string", Required: false},
	{Name: "Quantity", Flag: "quantity", Type: "*int32", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_allocate_ipam_pool_cidr = []leanruntime.Field{
	{Name: "AllowedCidrs", Flag: "allowed-cidrs", Type: "[]string", Required: false},
	{Name: "Cidr", Flag: "cidr", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisallowedCidrs", Flag: "disallowed-cidrs", Type: "[]string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamPoolId", Flag: "ipam-pool-id", Type: "*string", Required: true},
	{Name: "NetmaskLength", Flag: "netmask-length", Type: "*int32", Required: false},
	{Name: "PreviewNextCidr", Flag: "preview-next-cidr", Type: "*bool", Required: false},
}

var fields_apply_security_groups_to_client_vpn_target_network = []leanruntime.Field{
	{Name: "ClientVpnEndpointId", Flag: "client-vpn-endpoint-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: true},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_assign_ipv6_addresses = []leanruntime.Field{
	{Name: "Ipv6AddressCount", Flag: "ipv6-address-count", Type: "*int32", Required: false},
	{Name: "Ipv6Addresses", Flag: "ipv6-addresses", Type: "[]string", Required: false},
	{Name: "Ipv6PrefixCount", Flag: "ipv6-prefix-count", Type: "*int32", Required: false},
	{Name: "Ipv6Prefixes", Flag: "ipv6-prefixes", Type: "[]string", Required: false},
	{Name: "NetworkInterfaceId", Flag: "network-interface-id", Type: "*string", Required: true},
}

var fields_assign_private_ip_addresses = []leanruntime.Field{
	{Name: "AllowReassignment", Flag: "allow-reassignment", Type: "*bool", Required: false},
	{Name: "Ipv4PrefixCount", Flag: "ipv4-prefix-count", Type: "*int32", Required: false},
	{Name: "Ipv4Prefixes", Flag: "ipv4-prefixes", Type: "[]string", Required: false},
	{Name: "NetworkInterfaceId", Flag: "network-interface-id", Type: "*string", Required: true},
	{Name: "PrivateIpAddresses", Flag: "private-ip-addresses", Type: "[]string", Required: false},
	{Name: "SecondaryPrivateIpAddressCount", Flag: "secondary-private-ip-address-count", Type: "*int32", Required: false},
}

var fields_assign_private_nat_gateway_address = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "NatGatewayId", Flag: "nat-gateway-id", Type: "*string", Required: true},
	{Name: "PrivateIpAddressCount", Flag: "private-ip-address-count", Type: "*int32", Required: false},
	{Name: "PrivateIpAddresses", Flag: "private-ip-addresses", Type: "[]string", Required: false},
}

var fields_associate_address = []leanruntime.Field{
	{Name: "AllocationId", Flag: "allocation-id", Type: "*string", Required: false},
	{Name: "AllowReassociation", Flag: "allow-reassociation", Type: "*bool", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: false},
	{Name: "NetworkInterfaceId", Flag: "network-interface-id", Type: "*string", Required: false},
	{Name: "PrivateIpAddress", Flag: "private-ip-address", Type: "*string", Required: false},
	{Name: "PublicIp", Flag: "public-ip", Type: "*string", Required: false},
}

var fields_associate_capacity_reservation_billing_owner = []leanruntime.Field{
	{Name: "CapacityReservationId", Flag: "capacity-reservation-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "UnusedReservationBillingOwnerId", Flag: "unused-reservation-billing-owner-id", Type: "*string", Required: true},
}

var fields_associate_client_vpn_target_network = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ClientVpnEndpointId", Flag: "client-vpn-endpoint-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "SubnetId", Flag: "subnet-id", Type: "*string", Required: true},
}

var fields_associate_dhcp_options = []leanruntime.Field{
	{Name: "DhcpOptionsId", Flag: "dhcp-options-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_associate_enclave_certificate_iam_role = []leanruntime.Field{
	{Name: "CertificateArn", Flag: "certificate-arn", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
}

var fields_associate_iam_instance_profile = []leanruntime.Field{
	{Name: "IamInstanceProfile", Flag: "iam-instance-profile", Type: "*types.IamInstanceProfileSpecification", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_associate_instance_event_window = []leanruntime.Field{
	{Name: "AssociationTarget", Flag: "association-target", Type: "*types.InstanceEventWindowAssociationRequest", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceEventWindowId", Flag: "instance-event-window-id", Type: "*string", Required: true},
}

var fields_associate_ipam_byoasn = []leanruntime.Field{
	{Name: "Asn", Flag: "asn", Type: "*string", Required: true},
	{Name: "Cidr", Flag: "cidr", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_associate_ipam_resource_discovery = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamId", Flag: "ipam-id", Type: "*string", Required: true},
	{Name: "IpamResourceDiscoveryId", Flag: "ipam-resource-discovery-id", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_associate_nat_gateway_address = []leanruntime.Field{
	{Name: "AllocationIds", Flag: "allocation-ids", Type: "[]string", Required: true},
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "AvailabilityZoneId", Flag: "availability-zone-id", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "NatGatewayId", Flag: "nat-gateway-id", Type: "*string", Required: true},
	{Name: "PrivateIpAddresses", Flag: "private-ip-addresses", Type: "[]string", Required: false},
}

var fields_associate_route_server = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "RouteServerId", Flag: "route-server-id", Type: "*string", Required: true},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_associate_route_table = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: false},
	{Name: "PublicIpv4Pool", Flag: "public-ipv4-pool", Type: "*string", Required: false},
	{Name: "RouteTableId", Flag: "route-table-id", Type: "*string", Required: true},
	{Name: "SubnetId", Flag: "subnet-id", Type: "*string", Required: false},
}

var fields_associate_security_group_vpc = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_associate_subnet_cidr_block = []leanruntime.Field{
	{Name: "Ipv6CidrBlock", Flag: "ipv6-cidr-block", Type: "*string", Required: false},
	{Name: "Ipv6IpamPoolId", Flag: "ipv6-ipam-pool-id", Type: "*string", Required: false},
	{Name: "Ipv6NetmaskLength", Flag: "ipv6-netmask-length", Type: "*int32", Required: false},
	{Name: "SubnetId", Flag: "subnet-id", Type: "*string", Required: true},
}

var fields_associate_transit_gateway_multicast_domain = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: true},
	{Name: "TransitGatewayAttachmentId", Flag: "transit-gateway-attachment-id", Type: "*string", Required: true},
	{Name: "TransitGatewayMulticastDomainId", Flag: "transit-gateway-multicast-domain-id", Type: "*string", Required: true},
}

var fields_associate_transit_gateway_policy_table = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TransitGatewayAttachmentId", Flag: "transit-gateway-attachment-id", Type: "*string", Required: true},
	{Name: "TransitGatewayPolicyTableId", Flag: "transit-gateway-policy-table-id", Type: "*string", Required: true},
}

var fields_associate_transit_gateway_route_table = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TransitGatewayAttachmentId", Flag: "transit-gateway-attachment-id", Type: "*string", Required: true},
	{Name: "TransitGatewayRouteTableId", Flag: "transit-gateway-route-table-id", Type: "*string", Required: true},
}

var fields_associate_trunk_interface = []leanruntime.Field{
	{Name: "BranchInterfaceId", Flag: "branch-interface-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GreKey", Flag: "gre-key", Type: "*int32", Required: false},
	{Name: "TrunkInterfaceId", Flag: "trunk-interface-id", Type: "*string", Required: true},
	{Name: "VlanId", Flag: "vlan-id", Type: "*int32", Required: false},
}

var fields_associate_vpc_cidr_block = []leanruntime.Field{
	{Name: "AmazonProvidedIpv6CidrBlock", Flag: "amazon-provided-ipv6-cidr-block", Type: "*bool", Required: false},
	{Name: "CidrBlock", Flag: "cidr-block", Type: "*string", Required: false},
	{Name: "Ipv4IpamPoolId", Flag: "ipv4-ipam-pool-id", Type: "*string", Required: false},
	{Name: "Ipv4NetmaskLength", Flag: "ipv4-netmask-length", Type: "*int32", Required: false},
	{Name: "Ipv6CidrBlock", Flag: "ipv6-cidr-block", Type: "*string", Required: false},
	{Name: "Ipv6CidrBlockNetworkBorderGroup", Flag: "ipv6-cidr-block-network-border-group", Type: "*string", Required: false},
	{Name: "Ipv6IpamPoolId", Flag: "ipv6-ipam-pool-id", Type: "*string", Required: false},
	{Name: "Ipv6NetmaskLength", Flag: "ipv6-netmask-length", Type: "*int32", Required: false},
	{Name: "Ipv6Pool", Flag: "ipv6-pool", Type: "*string", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_attach_classic_link_vpc = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Groups", Flag: "groups", Type: "[]string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_attach_internet_gateway = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InternetGatewayId", Flag: "internet-gateway-id", Type: "*string", Required: true},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_attach_network_interface = []leanruntime.Field{
	{Name: "DeviceIndex", Flag: "device-index", Type: "*int32", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EnaQueueCount", Flag: "ena-queue-count", Type: "*int32", Required: false},
	{Name: "EnaSrdSpecification", Flag: "ena-srd-specification", Type: "*types.EnaSrdSpecification", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "NetworkCardIndex", Flag: "network-card-index", Type: "*int32", Required: false},
	{Name: "NetworkInterfaceId", Flag: "network-interface-id", Type: "*string", Required: true},
}

var fields_attach_verified_access_trust_provider = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VerifiedAccessInstanceId", Flag: "verified-access-instance-id", Type: "*string", Required: true},
	{Name: "VerifiedAccessTrustProviderId", Flag: "verified-access-trust-provider-id", Type: "*string", Required: true},
}

var fields_attach_volume = []leanruntime.Field{
	{Name: "Device", Flag: "device", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EbsCardIndex", Flag: "ebs-card-index", Type: "*int32", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "VolumeId", Flag: "volume-id", Type: "*string", Required: true},
}

var fields_attach_vpn_gateway = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
	{Name: "VpnGatewayId", Flag: "vpn-gateway-id", Type: "*string", Required: true},
}

var fields_authorize_client_vpn_ingress = []leanruntime.Field{
	{Name: "AccessGroupId", Flag: "access-group-id", Type: "*string", Required: false},
	{Name: "AuthorizeAllGroups", Flag: "authorize-all-groups", Type: "*bool", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ClientVpnEndpointId", Flag: "client-vpn-endpoint-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TargetNetworkCidr", Flag: "target-network-cidr", Type: "*string", Required: true},
}

var fields_authorize_security_group_egress = []leanruntime.Field{
	{Name: "CidrIp", Flag: "cidr-ip", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "FromPort", Flag: "from-port", Type: "*int32", Required: false},
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "IpPermissions", Flag: "ip-permissions", Type: "[]types.IpPermission", Required: false},
	{Name: "IpProtocol", Flag: "ip-protocol", Type: "*string", Required: false},
	{Name: "SourceSecurityGroupName", Flag: "source-security-group-name", Type: "*string", Required: false},
	{Name: "SourceSecurityGroupOwnerId", Flag: "source-security-group-owner-id", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "ToPort", Flag: "to-port", Type: "*int32", Required: false},
}

var fields_authorize_security_group_ingress = []leanruntime.Field{
	{Name: "CidrIp", Flag: "cidr-ip", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "FromPort", Flag: "from-port", Type: "*int32", Required: false},
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: false},
	{Name: "IpPermissions", Flag: "ip-permissions", Type: "[]types.IpPermission", Required: false},
	{Name: "IpProtocol", Flag: "ip-protocol", Type: "*string", Required: false},
	{Name: "SourceSecurityGroupName", Flag: "source-security-group-name", Type: "*string", Required: false},
	{Name: "SourceSecurityGroupOwnerId", Flag: "source-security-group-owner-id", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "ToPort", Flag: "to-port", Type: "*int32", Required: false},
}

var fields_bundle_instance = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Storage", Flag: "storage", Type: "*types.Storage", Required: true},
}

var fields_cancel_bundle_task = []leanruntime.Field{
	{Name: "BundleId", Flag: "bundle-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_cancel_capacity_reservation = []leanruntime.Field{
	{Name: "CapacityReservationId", Flag: "capacity-reservation-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_cancel_capacity_reservation_fleets = []leanruntime.Field{
	{Name: "CapacityReservationFleetIds", Flag: "capacity-reservation-fleet-ids", Type: "[]string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_cancel_conversion_task = []leanruntime.Field{
	{Name: "ConversionTaskId", Flag: "conversion-task-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ReasonMessage", Flag: "reason-message", Type: "*string", Required: false},
}

var fields_cancel_declarative_policies_report = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ReportId", Flag: "report-id", Type: "*string", Required: true},
}

var fields_cancel_export_task = []leanruntime.Field{
	{Name: "ExportTaskId", Flag: "export-task-id", Type: "*string", Required: true},
}

var fields_cancel_image_launch_permission = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: true},
}

var fields_cancel_import_task = []leanruntime.Field{
	{Name: "CancelReason", Flag: "cancel-reason", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ImportTaskId", Flag: "import-task-id", Type: "*string", Required: false},
}

var fields_cancel_reserved_instances_listing = []leanruntime.Field{
	{Name: "ReservedInstancesListingId", Flag: "reserved-instances-listing-id", Type: "*string", Required: true},
}

var fields_cancel_spot_fleet_requests = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "SpotFleetRequestIds", Flag: "spot-fleet-request-ids", Type: "[]string", Required: true},
	{Name: "TerminateInstances", Flag: "terminate-instances", Type: "*bool", Required: true},
}

var fields_cancel_spot_instance_requests = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "SpotInstanceRequestIds", Flag: "spot-instance-request-ids", Type: "[]string", Required: true},
}

var fields_confirm_product_instance = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "ProductCode", Flag: "product-code", Type: "*string", Required: true},
}

var fields_copy_fpga_image = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "SourceFpgaImageId", Flag: "source-fpga-image-id", Type: "*string", Required: true},
	{Name: "SourceRegion", Flag: "source-region", Type: "*string", Required: true},
}

var fields_copy_image = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CopyImageTags", Flag: "copy-image-tags", Type: "*bool", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DestinationAvailabilityZone", Flag: "destination-availability-zone", Type: "*string", Required: false},
	{Name: "DestinationAvailabilityZoneId", Flag: "destination-availability-zone-id", Type: "*string", Required: false},
	{Name: "DestinationOutpostArn", Flag: "destination-outpost-arn", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Encrypted", Flag: "encrypted", Type: "*bool", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SnapshotCopyCompletionDurationMinutes", Flag: "snapshot-copy-completion-duration-minutes", Type: "*int64", Required: false},
	{Name: "SourceImageId", Flag: "source-image-id", Type: "*string", Required: true},
	{Name: "SourceRegion", Flag: "source-region", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_copy_snapshot = []leanruntime.Field{
	{Name: "CompletionDurationMinutes", Flag: "completion-duration-minutes", Type: "*int32", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DestinationAvailabilityZone", Flag: "destination-availability-zone", Type: "*string", Required: false},
	{Name: "DestinationOutpostArn", Flag: "destination-outpost-arn", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Encrypted", Flag: "encrypted", Type: "*bool", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "PresignedUrl", Flag: "presigned-url", Type: "*string", Required: false},
	{Name: "SourceRegion", Flag: "source-region", Type: "*string", Required: true},
	{Name: "SourceSnapshotId", Flag: "source-snapshot-id", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "destinationRegion", Flag: "destination-region", Type: "*string", Required: false},
}

var fields_copy_volumes = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Iops", Flag: "iops", Type: "*int32", Required: false},
	{Name: "MultiAttachEnabled", Flag: "multi-attach-enabled", Type: "*bool", Required: false},
	{Name: "Size", Flag: "size", Type: "*int32", Required: false},
	{Name: "SourceVolumeId", Flag: "source-volume-id", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "Throughput", Flag: "throughput", Type: "*int32", Required: false},
	{Name: "VolumeType", Flag: "volume-type", Type: "types.VolumeType", Required: false},
}

var fields_create_capacity_manager_data_export = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "OutputFormat", Flag: "output-format", Type: "types.OutputFormat", Required: true},
	{Name: "S3BucketName", Flag: "s3-bucket-name", Type: "*string", Required: true},
	{Name: "S3BucketPrefix", Flag: "s3-bucket-prefix", Type: "*string", Required: false},
	{Name: "Schedule", Flag: "schedule", Type: "types.Schedule", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_capacity_reservation = []leanruntime.Field{
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "AvailabilityZoneId", Flag: "availability-zone-id", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CommitmentDuration", Flag: "commitment-duration", Type: "*int64", Required: false},
	{Name: "DeliveryPreference", Flag: "delivery-preference", Type: "types.CapacityReservationDeliveryPreference", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EbsOptimized", Flag: "ebs-optimized", Type: "*bool", Required: false},
	{Name: "EndDate", Flag: "end-date", Type: "*time.Time", Required: false},
	{Name: "EndDateType", Flag: "end-date-type", Type: "types.EndDateType", Required: false},
	{Name: "EphemeralStorage", Flag: "ephemeral-storage", Type: "*bool", Required: false},
	{Name: "InstanceCount", Flag: "instance-count", Type: "*int32", Required: true},
	{Name: "InstanceMatchCriteria", Flag: "instance-match-criteria", Type: "types.InstanceMatchCriteria", Required: false},
	{Name: "InstancePlatform", Flag: "instance-platform", Type: "types.CapacityReservationInstancePlatform", Required: true},
	{Name: "InstanceType", Flag: "instance-type", Type: "*string", Required: true},
	{Name: "OutpostArn", Flag: "outpost-arn", Type: "*string", Required: false},
	{Name: "PlacementGroupArn", Flag: "placement-group-arn", Type: "*string", Required: false},
	{Name: "StartDate", Flag: "start-date", Type: "*time.Time", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "Tenancy", Flag: "tenancy", Type: "types.CapacityReservationTenancy", Required: false},
}

var fields_create_capacity_reservation_by_splitting = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceCount", Flag: "instance-count", Type: "*int32", Required: true},
	{Name: "SourceCapacityReservationId", Flag: "source-capacity-reservation-id", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_capacity_reservation_fleet = []leanruntime.Field{
	{Name: "AllocationStrategy", Flag: "allocation-strategy", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EndDate", Flag: "end-date", Type: "*time.Time", Required: false},
	{Name: "InstanceMatchCriteria", Flag: "instance-match-criteria", Type: "types.FleetInstanceMatchCriteria", Required: false},
	{Name: "InstanceTypeSpecifications", Flag: "instance-type-specifications", Type: "[]types.ReservationFleetInstanceSpecification", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "Tenancy", Flag: "tenancy", Type: "types.FleetCapacityReservationTenancy", Required: false},
	{Name: "TotalTargetCapacity", Flag: "total-target-capacity", Type: "*int32", Required: true},
}

var fields_create_carrier_gateway = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_create_client_vpn_endpoint = []leanruntime.Field{
	{Name: "AuthenticationOptions", Flag: "authentication-options", Type: "[]types.ClientVpnAuthenticationRequest", Required: true},
	{Name: "ClientCidrBlock", Flag: "client-cidr-block", Type: "*string", Required: false},
	{Name: "ClientConnectOptions", Flag: "client-connect-options", Type: "*types.ClientConnectOptions", Required: false},
	{Name: "ClientLoginBannerOptions", Flag: "client-login-banner-options", Type: "*types.ClientLoginBannerOptions", Required: false},
	{Name: "ClientRouteEnforcementOptions", Flag: "client-route-enforcement-options", Type: "*types.ClientRouteEnforcementOptions", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ConnectionLogOptions", Flag: "connection-log-options", Type: "*types.ConnectionLogOptions", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisconnectOnSessionTimeout", Flag: "disconnect-on-session-timeout", Type: "*bool", Required: false},
	{Name: "DnsServers", Flag: "dns-servers", Type: "[]string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EndpointIpAddressType", Flag: "endpoint-ip-address-type", Type: "types.EndpointIpAddressType", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "SelfServicePortal", Flag: "self-service-portal", Type: "types.SelfServicePortal", Required: false},
	{Name: "ServerCertificateArn", Flag: "server-certificate-arn", Type: "*string", Required: true},
	{Name: "SessionTimeoutHours", Flag: "session-timeout-hours", Type: "*int32", Required: false},
	{Name: "SplitTunnel", Flag: "split-tunnel", Type: "*bool", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "TrafficIpAddressType", Flag: "traffic-ip-address-type", Type: "types.TrafficIpAddressType", Required: false},
	{Name: "TransportProtocol", Flag: "transport-protocol", Type: "types.TransportProtocol", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: false},
	{Name: "VpnPort", Flag: "vpn-port", Type: "*int32", Required: false},
}

var fields_create_client_vpn_route = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ClientVpnEndpointId", Flag: "client-vpn-endpoint-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DestinationCidrBlock", Flag: "destination-cidr-block", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TargetVpcSubnetId", Flag: "target-vpc-subnet-id", Type: "*string", Required: true},
}

var fields_create_coip_cidr = []leanruntime.Field{
	{Name: "Cidr", Flag: "cidr", Type: "*string", Required: true},
	{Name: "CoipPoolId", Flag: "coip-pool-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_create_coip_pool = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "LocalGatewayRouteTableId", Flag: "local-gateway-route-table-id", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_customer_gateway = []leanruntime.Field{
	{Name: "BgpAsn", Flag: "bgp-asn", Type: "*int32", Required: false},
	{Name: "BgpAsnExtended", Flag: "bgp-asn-extended", Type: "*int64", Required: false},
	{Name: "CertificateArn", Flag: "certificate-arn", Type: "*string", Required: false},
	{Name: "DeviceName", Flag: "device-name", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpAddress", Flag: "ip-address", Type: "*string", Required: false},
	{Name: "PublicIp", Flag: "public-ip", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "Type", Flag: "type", Type: "types.GatewayType", Required: true},
}

var fields_create_default_subnet = []leanruntime.Field{
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "AvailabilityZoneId", Flag: "availability-zone-id", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Ipv6Native", Flag: "ipv6-native", Type: "*bool", Required: false},
}

var fields_create_default_vpc = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_create_delegate_mac_volume_ownership_task = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MacCredentials", Flag: "mac-credentials", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_dhcp_options = []leanruntime.Field{
	{Name: "DhcpConfigurations", Flag: "dhcp-configurations", Type: "[]types.NewDhcpConfiguration", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_egress_only_internet_gateway = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_create_fleet = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Context", Flag: "context", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ExcessCapacityTerminationPolicy", Flag: "excess-capacity-termination-policy", Type: "types.FleetExcessCapacityTerminationPolicy", Required: false},
	{Name: "LaunchTemplateConfigs", Flag: "launch-template-configs", Type: "[]types.FleetLaunchTemplateConfigRequest", Required: true},
	{Name: "OnDemandOptions", Flag: "on-demand-options", Type: "*types.OnDemandOptionsRequest", Required: false},
	{Name: "ReplaceUnhealthyInstances", Flag: "replace-unhealthy-instances", Type: "*bool", Required: false},
	{Name: "SpotOptions", Flag: "spot-options", Type: "*types.SpotOptionsRequest", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "TargetCapacitySpecification", Flag: "target-capacity-specification", Type: "*types.TargetCapacitySpecificationRequest", Required: true},
	{Name: "TerminateInstancesWithExpiration", Flag: "terminate-instances-with-expiration", Type: "*bool", Required: false},
	{Name: "Type", Flag: "type", Type: "types.FleetType", Required: false},
	{Name: "ValidFrom", Flag: "valid-from", Type: "*time.Time", Required: false},
	{Name: "ValidUntil", Flag: "valid-until", Type: "*time.Time", Required: false},
}

var fields_create_flow_logs = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DeliverCrossAccountRole", Flag: "deliver-cross-account-role", Type: "*string", Required: false},
	{Name: "DeliverLogsPermissionArn", Flag: "deliver-logs-permission-arn", Type: "*string", Required: false},
	{Name: "DestinationOptions", Flag: "destination-options", Type: "*types.DestinationOptionsRequest", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "LogDestination", Flag: "log-destination", Type: "*string", Required: false},
	{Name: "LogDestinationType", Flag: "log-destination-type", Type: "types.LogDestinationType", Required: false},
	{Name: "LogFormat", Flag: "log-format", Type: "*string", Required: false},
	{Name: "LogGroupName", Flag: "log-group-name", Type: "*string", Required: false},
	{Name: "MaxAggregationInterval", Flag: "max-aggregation-interval", Type: "*int32", Required: false},
	{Name: "ResourceIds", Flag: "resource-ids", Type: "[]string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.FlowLogsResourceType", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "TrafficType", Flag: "traffic-type", Type: "types.TrafficType", Required: false},
}

var fields_create_fpga_image = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InputStorageLocation", Flag: "input-storage-location", Type: "*types.StorageLocation", Required: true},
	{Name: "LogsStorageLocation", Flag: "logs-storage-location", Type: "*types.StorageLocation", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_image = []leanruntime.Field{
	{Name: "BlockDeviceMappings", Flag: "block-device-mappings", Type: "[]types.BlockDeviceMapping", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NoReboot", Flag: "no-reboot", Type: "*bool", Required: false},
	{Name: "SnapshotLocation", Flag: "snapshot-location", Type: "types.SnapshotLocationEnum", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_image_usage_report = []leanruntime.Field{
	{Name: "AccountIds", Flag: "account-ids", Type: "[]string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: true},
	{Name: "ResourceTypes", Flag: "resource-types", Type: "[]types.ImageUsageResourceTypeRequest", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_instance_connect_endpoint = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.IpAddressType", Required: false},
	{Name: "PreserveClientIp", Flag: "preserve-client-ip", Type: "*bool", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "SubnetId", Flag: "subnet-id", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_instance_event_window = []leanruntime.Field{
	{Name: "CronExpression", Flag: "cron-expression", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "TimeRanges", Flag: "time-ranges", Type: "[]types.InstanceEventWindowTimeRangeRequest", Required: false},
}

var fields_create_instance_export_task = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExportToS3Task", Flag: "export-to-s3-task", Type: "*types.ExportToS3TaskSpecification", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "TargetEnvironment", Flag: "target-environment", Type: "types.ExportEnvironment", Required: true},
}

var fields_create_internet_gateway = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_interruptible_capacity_reservation_allocation = []leanruntime.Field{
	{Name: "CapacityReservationId", Flag: "capacity-reservation-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceCount", Flag: "instance-count", Type: "*int32", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_ipam = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EnablePrivateGua", Flag: "enable-private-gua", Type: "*bool", Required: false},
	{Name: "MeteredAccount", Flag: "metered-account", Type: "types.IpamMeteredAccount", Required: false},
	{Name: "OperatingRegions", Flag: "operating-regions", Type: "[]types.AddIpamOperatingRegion", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "Tier", Flag: "tier", Type: "types.IpamTier", Required: false},
}

var fields_create_ipam_external_resource_verification_token = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamId", Flag: "ipam-id", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_ipam_policy = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamId", Flag: "ipam-id", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_ipam_pool = []leanruntime.Field{
	{Name: "AddressFamily", Flag: "address-family", Type: "types.AddressFamily", Required: true},
	{Name: "AllocationDefaultNetmaskLength", Flag: "allocation-default-netmask-length", Type: "*int32", Required: false},
	{Name: "AllocationMaxNetmaskLength", Flag: "allocation-max-netmask-length", Type: "*int32", Required: false},
	{Name: "AllocationMinNetmaskLength", Flag: "allocation-min-netmask-length", Type: "*int32", Required: false},
	{Name: "AllocationResourceTags", Flag: "allocation-resource-tags", Type: "[]types.RequestIpamResourceTag", Required: false},
	{Name: "AutoImport", Flag: "auto-import", Type: "*bool", Required: false},
	{Name: "AwsService", Flag: "aws-service", Type: "types.IpamPoolAwsService", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamScopeId", Flag: "ipam-scope-id", Type: "*string", Required: true},
	{Name: "Locale", Flag: "locale", Type: "*string", Required: false},
	{Name: "PublicIpSource", Flag: "public-ip-source", Type: "types.IpamPoolPublicIpSource", Required: false},
	{Name: "PubliclyAdvertisable", Flag: "publicly-advertisable", Type: "*bool", Required: false},
	{Name: "SourceIpamPoolId", Flag: "source-ipam-pool-id", Type: "*string", Required: false},
	{Name: "SourceResource", Flag: "source-resource", Type: "*types.IpamPoolSourceResourceRequest", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_ipam_prefix_list_resolver = []leanruntime.Field{
	{Name: "AddressFamily", Flag: "address-family", Type: "types.AddressFamily", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamId", Flag: "ipam-id", Type: "*string", Required: true},
	{Name: "Rules", Flag: "rules", Type: "[]types.IpamPrefixListResolverRuleRequest", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_ipam_prefix_list_resolver_target = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DesiredVersion", Flag: "desired-version", Type: "*int64", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamPrefixListResolverId", Flag: "ipam-prefix-list-resolver-id", Type: "*string", Required: true},
	{Name: "PrefixListId", Flag: "prefix-list-id", Type: "*string", Required: true},
	{Name: "PrefixListRegion", Flag: "prefix-list-region", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "TrackLatestVersion", Flag: "track-latest-version", Type: "*bool", Required: true},
}

var fields_create_ipam_resource_discovery = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "OperatingRegions", Flag: "operating-regions", Type: "[]types.AddIpamOperatingRegion", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_ipam_scope = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ExternalAuthorityConfiguration", Flag: "external-authority-configuration", Type: "*types.ExternalAuthorityConfiguration", Required: false},
	{Name: "IpamId", Flag: "ipam-id", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_key_pair = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "KeyFormat", Flag: "key-format", Type: "types.KeyFormat", Required: false},
	{Name: "KeyName", Flag: "key-name", Type: "*string", Required: true},
	{Name: "KeyType", Flag: "key-type", Type: "types.KeyType", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_launch_template = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "LaunchTemplateData", Flag: "launch-template-data", Type: "*types.RequestLaunchTemplateData", Required: true},
	{Name: "LaunchTemplateName", Flag: "launch-template-name", Type: "*string", Required: true},
	{Name: "Operator", Flag: "operator", Type: "*types.OperatorRequest", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "VersionDescription", Flag: "version-description", Type: "*string", Required: false},
}

var fields_create_launch_template_version = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "LaunchTemplateData", Flag: "launch-template-data", Type: "*types.RequestLaunchTemplateData", Required: true},
	{Name: "LaunchTemplateId", Flag: "launch-template-id", Type: "*string", Required: false},
	{Name: "LaunchTemplateName", Flag: "launch-template-name", Type: "*string", Required: false},
	{Name: "ResolveAlias", Flag: "resolve-alias", Type: "*bool", Required: false},
	{Name: "SourceVersion", Flag: "source-version", Type: "*string", Required: false},
	{Name: "VersionDescription", Flag: "version-description", Type: "*string", Required: false},
}

var fields_create_local_gateway_route = []leanruntime.Field{
	{Name: "DestinationCidrBlock", Flag: "destination-cidr-block", Type: "*string", Required: false},
	{Name: "DestinationPrefixListId", Flag: "destination-prefix-list-id", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "LocalGatewayRouteTableId", Flag: "local-gateway-route-table-id", Type: "*string", Required: true},
	{Name: "LocalGatewayVirtualInterfaceGroupId", Flag: "local-gateway-virtual-interface-group-id", Type: "*string", Required: false},
	{Name: "NetworkInterfaceId", Flag: "network-interface-id", Type: "*string", Required: false},
}

var fields_create_local_gateway_route_table = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "LocalGatewayId", Flag: "local-gateway-id", Type: "*string", Required: true},
	{Name: "Mode", Flag: "mode", Type: "types.LocalGatewayRouteTableMode", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_local_gateway_route_table_virtual_interface_group_association = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "LocalGatewayRouteTableId", Flag: "local-gateway-route-table-id", Type: "*string", Required: true},
	{Name: "LocalGatewayVirtualInterfaceGroupId", Flag: "local-gateway-virtual-interface-group-id", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_local_gateway_route_table_vpc_association = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "LocalGatewayRouteTableId", Flag: "local-gateway-route-table-id", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_create_local_gateway_virtual_interface = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "LocalAddress", Flag: "local-address", Type: "*string", Required: true},
	{Name: "LocalGatewayVirtualInterfaceGroupId", Flag: "local-gateway-virtual-interface-group-id", Type: "*string", Required: true},
	{Name: "OutpostLagId", Flag: "outpost-lag-id", Type: "*string", Required: true},
	{Name: "PeerAddress", Flag: "peer-address", Type: "*string", Required: true},
	{Name: "PeerBgpAsn", Flag: "peer-bgp-asn", Type: "*int32", Required: false},
	{Name: "PeerBgpAsnExtended", Flag: "peer-bgp-asn-extended", Type: "*int64", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "Vlan", Flag: "vlan", Type: "*int32", Required: true},
}

var fields_create_local_gateway_virtual_interface_group = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "LocalBgpAsn", Flag: "local-bgp-asn", Type: "*int32", Required: false},
	{Name: "LocalBgpAsnExtended", Flag: "local-bgp-asn-extended", Type: "*int64", Required: false},
	{Name: "LocalGatewayId", Flag: "local-gateway-id", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_mac_system_integrity_protection_modification_task = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "MacCredentials", Flag: "mac-credentials", Type: "*string", Required: false},
	{Name: "MacSystemIntegrityProtectionConfiguration", Flag: "mac-system-integrity-protection-configuration", Type: "*types.MacSystemIntegrityProtectionConfigurationRequest", Required: false},
	{Name: "MacSystemIntegrityProtectionStatus", Flag: "mac-system-integrity-protection-status", Type: "types.MacSystemIntegrityProtectionSettingStatus", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_managed_prefix_list = []leanruntime.Field{
	{Name: "AddressFamily", Flag: "address-family", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Entries", Flag: "entries", Type: "[]types.AddPrefixListEntry", Required: false},
	{Name: "MaxEntries", Flag: "max-entries", Type: "*int32", Required: true},
	{Name: "PrefixListName", Flag: "prefix-list-name", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_nat_gateway = []leanruntime.Field{
	{Name: "AllocationId", Flag: "allocation-id", Type: "*string", Required: false},
	{Name: "AvailabilityMode", Flag: "availability-mode", Type: "types.AvailabilityMode", Required: false},
	{Name: "AvailabilityZoneAddresses", Flag: "availability-zone-addresses", Type: "[]types.AvailabilityZoneAddress", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ConnectivityType", Flag: "connectivity-type", Type: "types.ConnectivityType", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "PrivateIpAddress", Flag: "private-ip-address", Type: "*string", Required: false},
	{Name: "SecondaryAllocationIds", Flag: "secondary-allocation-ids", Type: "[]string", Required: false},
	{Name: "SecondaryPrivateIpAddressCount", Flag: "secondary-private-ip-address-count", Type: "*int32", Required: false},
	{Name: "SecondaryPrivateIpAddresses", Flag: "secondary-private-ip-addresses", Type: "[]string", Required: false},
	{Name: "SubnetId", Flag: "subnet-id", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: false},
}

var fields_create_network_acl = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_create_network_acl_entry = []leanruntime.Field{
	{Name: "CidrBlock", Flag: "cidr-block", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Egress", Flag: "egress", Type: "*bool", Required: true},
	{Name: "IcmpTypeCode", Flag: "icmp-type-code", Type: "*types.IcmpTypeCode", Required: false},
	{Name: "Ipv6CidrBlock", Flag: "ipv6-cidr-block", Type: "*string", Required: false},
	{Name: "NetworkAclId", Flag: "network-acl-id", Type: "*string", Required: true},
	{Name: "PortRange", Flag: "port-range", Type: "*types.PortRange", Required: false},
	{Name: "Protocol", Flag: "protocol", Type: "*string", Required: true},
	{Name: "RuleAction", Flag: "rule-action", Type: "types.RuleAction", Required: true},
	{Name: "RuleNumber", Flag: "rule-number", Type: "*int32", Required: true},
}

var fields_create_network_insights_access_scope = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ExcludePaths", Flag: "exclude-paths", Type: "[]types.AccessScopePathRequest", Required: false},
	{Name: "MatchPaths", Flag: "match-paths", Type: "[]types.AccessScopePathRequest", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_network_insights_path = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Destination", Flag: "destination", Type: "*string", Required: false},
	{Name: "DestinationIp", Flag: "destination-ip", Type: "*string", Required: false},
	{Name: "DestinationPort", Flag: "destination-port", Type: "*int32", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "FilterAtDestination", Flag: "filter-at-destination", Type: "*types.PathRequestFilter", Required: false},
	{Name: "FilterAtSource", Flag: "filter-at-source", Type: "*types.PathRequestFilter", Required: false},
	{Name: "Protocol", Flag: "protocol", Type: "types.Protocol", Required: true},
	{Name: "Source", Flag: "source", Type: "*string", Required: true},
	{Name: "SourceIp", Flag: "source-ip", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_network_interface = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ConnectionTrackingSpecification", Flag: "connection-tracking-specification", Type: "*types.ConnectionTrackingSpecificationRequest", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EnablePrimaryIpv6", Flag: "enable-primary-ipv6", Type: "*bool", Required: false},
	{Name: "Groups", Flag: "groups", Type: "[]string", Required: false},
	{Name: "InterfaceType", Flag: "interface-type", Type: "types.NetworkInterfaceCreationType", Required: false},
	{Name: "Ipv4PrefixCount", Flag: "ipv4-prefix-count", Type: "*int32", Required: false},
	{Name: "Ipv4Prefixes", Flag: "ipv4-prefixes", Type: "[]types.Ipv4PrefixSpecificationRequest", Required: false},
	{Name: "Ipv6AddressCount", Flag: "ipv6-address-count", Type: "*int32", Required: false},
	{Name: "Ipv6Addresses", Flag: "ipv6-addresses", Type: "[]types.InstanceIpv6Address", Required: false},
	{Name: "Ipv6PrefixCount", Flag: "ipv6-prefix-count", Type: "*int32", Required: false},
	{Name: "Ipv6Prefixes", Flag: "ipv6-prefixes", Type: "[]types.Ipv6PrefixSpecificationRequest", Required: false},
	{Name: "Operator", Flag: "operator", Type: "*types.OperatorRequest", Required: false},
	{Name: "PrivateIpAddress", Flag: "private-ip-address", Type: "*string", Required: false},
	{Name: "PrivateIpAddresses", Flag: "private-ip-addresses", Type: "[]types.PrivateIpAddressSpecification", Required: false},
	{Name: "SecondaryPrivateIpAddressCount", Flag: "secondary-private-ip-address-count", Type: "*int32", Required: false},
	{Name: "SubnetId", Flag: "subnet-id", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_network_interface_permission = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: false},
	{Name: "AwsService", Flag: "aws-service", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "NetworkInterfaceId", Flag: "network-interface-id", Type: "*string", Required: true},
	{Name: "Permission", Flag: "permission", Type: "types.InterfacePermissionType", Required: true},
}

var fields_create_placement_group = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: false},
	{Name: "LinkedGroupId", Flag: "linked-group-id", Type: "*string", Required: false},
	{Name: "Operator", Flag: "operator", Type: "*types.OperatorRequest", Required: false},
	{Name: "PartitionCount", Flag: "partition-count", Type: "*int32", Required: false},
	{Name: "SpreadLevel", Flag: "spread-level", Type: "types.SpreadLevel", Required: false},
	{Name: "Strategy", Flag: "strategy", Type: "types.PlacementStrategy", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_public_ipv4_pool = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "NetworkBorderGroup", Flag: "network-border-group", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_replace_root_volume_task = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DeleteReplacedRootVolume", Flag: "delete-replaced-root-volume", Type: "*bool", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "SnapshotId", Flag: "snapshot-id", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "VolumeInitializationRate", Flag: "volume-initialization-rate", Type: "*int64", Required: false},
}

var fields_create_reserved_instances_listing = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "InstanceCount", Flag: "instance-count", Type: "*int32", Required: true},
	{Name: "PriceSchedules", Flag: "price-schedules", Type: "[]types.PriceScheduleSpecification", Required: true},
	{Name: "ReservedInstancesId", Flag: "reserved-instances-id", Type: "*string", Required: true},
}

var fields_create_restore_image_task = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ObjectKey", Flag: "object-key", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_route = []leanruntime.Field{
	{Name: "CarrierGatewayId", Flag: "carrier-gateway-id", Type: "*string", Required: false},
	{Name: "CoreNetworkArn", Flag: "core-network-arn", Type: "*string", Required: false},
	{Name: "DestinationCidrBlock", Flag: "destination-cidr-block", Type: "*string", Required: false},
	{Name: "DestinationIpv6CidrBlock", Flag: "destination-ipv6-cidr-block", Type: "*string", Required: false},
	{Name: "DestinationPrefixListId", Flag: "destination-prefix-list-id", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EgressOnlyInternetGatewayId", Flag: "egress-only-internet-gateway-id", Type: "*string", Required: false},
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: false},
	{Name: "LocalGatewayId", Flag: "local-gateway-id", Type: "*string", Required: false},
	{Name: "NatGatewayId", Flag: "nat-gateway-id", Type: "*string", Required: false},
	{Name: "NetworkInterfaceId", Flag: "network-interface-id", Type: "*string", Required: false},
	{Name: "OdbNetworkArn", Flag: "odb-network-arn", Type: "*string", Required: false},
	{Name: "RouteTableId", Flag: "route-table-id", Type: "*string", Required: true},
	{Name: "TransitGatewayId", Flag: "transit-gateway-id", Type: "*string", Required: false},
	{Name: "VpcEndpointId", Flag: "vpc-endpoint-id", Type: "*string", Required: false},
	{Name: "VpcPeeringConnectionId", Flag: "vpc-peering-connection-id", Type: "*string", Required: false},
}

var fields_create_route_server = []leanruntime.Field{
	{Name: "AmazonSideAsn", Flag: "amazon-side-asn", Type: "*int64", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "PersistRoutes", Flag: "persist-routes", Type: "types.RouteServerPersistRoutesAction", Required: false},
	{Name: "PersistRoutesDuration", Flag: "persist-routes-duration", Type: "*int64", Required: false},
	{Name: "SnsNotificationsEnabled", Flag: "sns-notifications-enabled", Type: "*bool", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_route_server_endpoint = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "RouteServerId", Flag: "route-server-id", Type: "*string", Required: true},
	{Name: "SubnetId", Flag: "subnet-id", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_route_server_peer = []leanruntime.Field{
	{Name: "BgpOptions", Flag: "bgp-options", Type: "*types.RouteServerBgpOptionsRequest", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "PeerAddress", Flag: "peer-address", Type: "*string", Required: true},
	{Name: "RouteServerEndpointId", Flag: "route-server-endpoint-id", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_route_table = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_create_secondary_network = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Ipv4CidrBlock", Flag: "ipv4-cidr-block", Type: "*string", Required: true},
	{Name: "NetworkType", Flag: "network-type", Type: "types.SecondaryNetworkType", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_secondary_subnet = []leanruntime.Field{
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "AvailabilityZoneId", Flag: "availability-zone-id", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Ipv4CidrBlock", Flag: "ipv4-cidr-block", Type: "*string", Required: true},
	{Name: "SecondaryNetworkId", Flag: "secondary-network-id", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_security_group = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: false},
}

var fields_create_snapshot = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Location", Flag: "location", Type: "types.SnapshotLocationEnum", Required: false},
	{Name: "OutpostArn", Flag: "outpost-arn", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "VolumeId", Flag: "volume-id", Type: "*string", Required: true},
}

var fields_create_snapshots = []leanruntime.Field{
	{Name: "CopyTagsFromSource", Flag: "copy-tags-from-source", Type: "types.CopyTagsFromSource", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceSpecification", Flag: "instance-specification", Type: "*types.InstanceSpecification", Required: true},
	{Name: "Location", Flag: "location", Type: "types.SnapshotLocationEnum", Required: false},
	{Name: "OutpostArn", Flag: "outpost-arn", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_spot_datafeed_subscription = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Prefix", Flag: "prefix", Type: "*string", Required: false},
}

var fields_create_store_image_task = []leanruntime.Field{
	{Name: "Bucket", Flag: "bucket", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: true},
	{Name: "S3ObjectTags", Flag: "s3-object-tags", Type: "[]types.S3ObjectTag", Required: false},
}

var fields_create_subnet = []leanruntime.Field{
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "AvailabilityZoneId", Flag: "availability-zone-id", Type: "*string", Required: false},
	{Name: "CidrBlock", Flag: "cidr-block", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Ipv4IpamPoolId", Flag: "ipv4-ipam-pool-id", Type: "*string", Required: false},
	{Name: "Ipv4NetmaskLength", Flag: "ipv4-netmask-length", Type: "*int32", Required: false},
	{Name: "Ipv6CidrBlock", Flag: "ipv6-cidr-block", Type: "*string", Required: false},
	{Name: "Ipv6IpamPoolId", Flag: "ipv6-ipam-pool-id", Type: "*string", Required: false},
	{Name: "Ipv6Native", Flag: "ipv6-native", Type: "*bool", Required: false},
	{Name: "Ipv6NetmaskLength", Flag: "ipv6-netmask-length", Type: "*int32", Required: false},
	{Name: "OutpostArn", Flag: "outpost-arn", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_create_subnet_cidr_reservation = []leanruntime.Field{
	{Name: "Cidr", Flag: "cidr", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ReservationType", Flag: "reservation-type", Type: "types.SubnetCidrReservationType", Required: true},
	{Name: "SubnetId", Flag: "subnet-id", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_tags = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Resources", Flag: "resources", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_create_traffic_mirror_filter = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_traffic_mirror_filter_rule = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DestinationCidrBlock", Flag: "destination-cidr-block", Type: "*string", Required: true},
	{Name: "DestinationPortRange", Flag: "destination-port-range", Type: "*types.TrafficMirrorPortRangeRequest", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Protocol", Flag: "protocol", Type: "*int32", Required: false},
	{Name: "RuleAction", Flag: "rule-action", Type: "types.TrafficMirrorRuleAction", Required: true},
	{Name: "RuleNumber", Flag: "rule-number", Type: "*int32", Required: true},
	{Name: "SourceCidrBlock", Flag: "source-cidr-block", Type: "*string", Required: true},
	{Name: "SourcePortRange", Flag: "source-port-range", Type: "*types.TrafficMirrorPortRangeRequest", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "TrafficDirection", Flag: "traffic-direction", Type: "types.TrafficDirection", Required: true},
	{Name: "TrafficMirrorFilterId", Flag: "traffic-mirror-filter-id", Type: "*string", Required: true},
}

var fields_create_traffic_mirror_session = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "NetworkInterfaceId", Flag: "network-interface-id", Type: "*string", Required: true},
	{Name: "PacketLength", Flag: "packet-length", Type: "*int32", Required: false},
	{Name: "SessionNumber", Flag: "session-number", Type: "*int32", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "TrafficMirrorFilterId", Flag: "traffic-mirror-filter-id", Type: "*string", Required: true},
	{Name: "TrafficMirrorTargetId", Flag: "traffic-mirror-target-id", Type: "*string", Required: true},
	{Name: "VirtualNetworkId", Flag: "virtual-network-id", Type: "*int32", Required: false},
}

var fields_create_traffic_mirror_target = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GatewayLoadBalancerEndpointId", Flag: "gateway-load-balancer-endpoint-id", Type: "*string", Required: false},
	{Name: "NetworkInterfaceId", Flag: "network-interface-id", Type: "*string", Required: false},
	{Name: "NetworkLoadBalancerArn", Flag: "network-load-balancer-arn", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_transit_gateway = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Options", Flag: "options", Type: "*types.TransitGatewayRequestOptions", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_transit_gateway_connect = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Options", Flag: "options", Type: "*types.CreateTransitGatewayConnectRequestOptions", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "TransportTransitGatewayAttachmentId", Flag: "transport-transit-gateway-attachment-id", Type: "*string", Required: true},
}

var fields_create_transit_gateway_connect_peer = []leanruntime.Field{
	{Name: "BgpOptions", Flag: "bgp-options", Type: "*types.TransitGatewayConnectRequestBgpOptions", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InsideCidrBlocks", Flag: "inside-cidr-blocks", Type: "[]string", Required: true},
	{Name: "PeerAddress", Flag: "peer-address", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "TransitGatewayAddress", Flag: "transit-gateway-address", Type: "*string", Required: false},
	{Name: "TransitGatewayAttachmentId", Flag: "transit-gateway-attachment-id", Type: "*string", Required: true},
}

var fields_create_transit_gateway_metering_policy = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "MiddleboxAttachmentIds", Flag: "middlebox-attachment-ids", Type: "[]string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "TransitGatewayId", Flag: "transit-gateway-id", Type: "*string", Required: true},
}

var fields_create_transit_gateway_metering_policy_entry = []leanruntime.Field{
	{Name: "DestinationCidrBlock", Flag: "destination-cidr-block", Type: "*string", Required: false},
	{Name: "DestinationPortRange", Flag: "destination-port-range", Type: "*string", Required: false},
	{Name: "DestinationTransitGatewayAttachmentId", Flag: "destination-transit-gateway-attachment-id", Type: "*string", Required: false},
	{Name: "DestinationTransitGatewayAttachmentType", Flag: "destination-transit-gateway-attachment-type", Type: "types.TransitGatewayAttachmentResourceType", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "MeteredAccount", Flag: "metered-account", Type: "types.TransitGatewayMeteringPayerType", Required: true},
	{Name: "PolicyRuleNumber", Flag: "policy-rule-number", Type: "*int32", Required: true},
	{Name: "Protocol", Flag: "protocol", Type: "*string", Required: false},
	{Name: "SourceCidrBlock", Flag: "source-cidr-block", Type: "*string", Required: false},
	{Name: "SourcePortRange", Flag: "source-port-range", Type: "*string", Required: false},
	{Name: "SourceTransitGatewayAttachmentId", Flag: "source-transit-gateway-attachment-id", Type: "*string", Required: false},
	{Name: "SourceTransitGatewayAttachmentType", Flag: "source-transit-gateway-attachment-type", Type: "types.TransitGatewayAttachmentResourceType", Required: false},
	{Name: "TransitGatewayMeteringPolicyId", Flag: "transit-gateway-metering-policy-id", Type: "*string", Required: true},
}

var fields_create_transit_gateway_multicast_domain = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Options", Flag: "options", Type: "*types.CreateTransitGatewayMulticastDomainRequestOptions", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "TransitGatewayId", Flag: "transit-gateway-id", Type: "*string", Required: true},
}

var fields_create_transit_gateway_peering_attachment = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Options", Flag: "options", Type: "*types.CreateTransitGatewayPeeringAttachmentRequestOptions", Required: false},
	{Name: "PeerAccountId", Flag: "peer-account-id", Type: "*string", Required: true},
	{Name: "PeerRegion", Flag: "peer-region", Type: "*string", Required: true},
	{Name: "PeerTransitGatewayId", Flag: "peer-transit-gateway-id", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "TransitGatewayId", Flag: "transit-gateway-id", Type: "*string", Required: true},
}

var fields_create_transit_gateway_policy_table = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "TransitGatewayId", Flag: "transit-gateway-id", Type: "*string", Required: true},
}

var fields_create_transit_gateway_prefix_list_reference = []leanruntime.Field{
	{Name: "Blackhole", Flag: "blackhole", Type: "*bool", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "PrefixListId", Flag: "prefix-list-id", Type: "*string", Required: true},
	{Name: "TransitGatewayAttachmentId", Flag: "transit-gateway-attachment-id", Type: "*string", Required: false},
	{Name: "TransitGatewayRouteTableId", Flag: "transit-gateway-route-table-id", Type: "*string", Required: true},
}

var fields_create_transit_gateway_route = []leanruntime.Field{
	{Name: "Blackhole", Flag: "blackhole", Type: "*bool", Required: false},
	{Name: "DestinationCidrBlock", Flag: "destination-cidr-block", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TransitGatewayAttachmentId", Flag: "transit-gateway-attachment-id", Type: "*string", Required: false},
	{Name: "TransitGatewayRouteTableId", Flag: "transit-gateway-route-table-id", Type: "*string", Required: true},
}

var fields_create_transit_gateway_route_table = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "TransitGatewayId", Flag: "transit-gateway-id", Type: "*string", Required: true},
}

var fields_create_transit_gateway_route_table_announcement = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "PeeringAttachmentId", Flag: "peering-attachment-id", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "TransitGatewayRouteTableId", Flag: "transit-gateway-route-table-id", Type: "*string", Required: true},
}

var fields_create_transit_gateway_vpc_attachment = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Options", Flag: "options", Type: "*types.CreateTransitGatewayVpcAttachmentRequestOptions", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "TransitGatewayId", Flag: "transit-gateway-id", Type: "*string", Required: true},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_create_verified_access_endpoint = []leanruntime.Field{
	{Name: "ApplicationDomain", Flag: "application-domain", Type: "*string", Required: false},
	{Name: "AttachmentType", Flag: "attachment-type", Type: "types.VerifiedAccessEndpointAttachmentType", Required: true},
	{Name: "CidrOptions", Flag: "cidr-options", Type: "*types.CreateVerifiedAccessEndpointCidrOptions", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainCertificateArn", Flag: "domain-certificate-arn", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EndpointDomainPrefix", Flag: "endpoint-domain-prefix", Type: "*string", Required: false},
	{Name: "EndpointType", Flag: "endpoint-type", Type: "types.VerifiedAccessEndpointType", Required: true},
	{Name: "LoadBalancerOptions", Flag: "load-balancer-options", Type: "*types.CreateVerifiedAccessEndpointLoadBalancerOptions", Required: false},
	{Name: "NetworkInterfaceOptions", Flag: "network-interface-options", Type: "*types.CreateVerifiedAccessEndpointEniOptions", Required: false},
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: false},
	{Name: "RdsOptions", Flag: "rds-options", Type: "*types.CreateVerifiedAccessEndpointRdsOptions", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "SseSpecification", Flag: "sse-specification", Type: "*types.VerifiedAccessSseSpecificationRequest", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "VerifiedAccessGroupId", Flag: "verified-access-group-id", Type: "*string", Required: true},
}

var fields_create_verified_access_group = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: false},
	{Name: "SseSpecification", Flag: "sse-specification", Type: "*types.VerifiedAccessSseSpecificationRequest", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "VerifiedAccessInstanceId", Flag: "verified-access-instance-id", Type: "*string", Required: true},
}

var fields_create_verified_access_instance = []leanruntime.Field{
	{Name: "CidrEndpointsCustomSubDomain", Flag: "cidr-endpoints-custom-sub-domain", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "FIPSEnabled", Flag: "fips-enabled", Type: "*bool", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_verified_access_trust_provider = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DeviceOptions", Flag: "device-options", Type: "*types.CreateVerifiedAccessTrustProviderDeviceOptions", Required: false},
	{Name: "DeviceTrustProviderType", Flag: "device-trust-provider-type", Type: "types.DeviceTrustProviderType", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "NativeApplicationOidcOptions", Flag: "native-application-oidc-options", Type: "*types.CreateVerifiedAccessNativeApplicationOidcOptions", Required: false},
	{Name: "OidcOptions", Flag: "oidc-options", Type: "*types.CreateVerifiedAccessTrustProviderOidcOptions", Required: false},
	{Name: "PolicyReferenceName", Flag: "policy-reference-name", Type: "*string", Required: true},
	{Name: "SseSpecification", Flag: "sse-specification", Type: "*types.VerifiedAccessSseSpecificationRequest", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "TrustProviderType", Flag: "trust-provider-type", Type: "types.TrustProviderType", Required: true},
	{Name: "UserTrustProviderType", Flag: "user-trust-provider-type", Type: "types.UserTrustProviderType", Required: false},
}

var fields_create_volume = []leanruntime.Field{
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "AvailabilityZoneId", Flag: "availability-zone-id", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Encrypted", Flag: "encrypted", Type: "*bool", Required: false},
	{Name: "Iops", Flag: "iops", Type: "*int32", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "MultiAttachEnabled", Flag: "multi-attach-enabled", Type: "*bool", Required: false},
	{Name: "Operator", Flag: "operator", Type: "*types.OperatorRequest", Required: false},
	{Name: "OutpostArn", Flag: "outpost-arn", Type: "*string", Required: false},
	{Name: "Size", Flag: "size", Type: "*int32", Required: false},
	{Name: "SnapshotId", Flag: "snapshot-id", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "Throughput", Flag: "throughput", Type: "*int32", Required: false},
	{Name: "VolumeInitializationRate", Flag: "volume-initialization-rate", Type: "*int32", Required: false},
	{Name: "VolumeType", Flag: "volume-type", Type: "types.VolumeType", Required: false},
}

var fields_create_vpc = []leanruntime.Field{
	{Name: "AmazonProvidedIpv6CidrBlock", Flag: "amazon-provided-ipv6-cidr-block", Type: "*bool", Required: false},
	{Name: "CidrBlock", Flag: "cidr-block", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceTenancy", Flag: "instance-tenancy", Type: "types.Tenancy", Required: false},
	{Name: "Ipv4IpamPoolId", Flag: "ipv4-ipam-pool-id", Type: "*string", Required: false},
	{Name: "Ipv4NetmaskLength", Flag: "ipv4-netmask-length", Type: "*int32", Required: false},
	{Name: "Ipv6CidrBlock", Flag: "ipv6-cidr-block", Type: "*string", Required: false},
	{Name: "Ipv6CidrBlockNetworkBorderGroup", Flag: "ipv6-cidr-block-network-border-group", Type: "*string", Required: false},
	{Name: "Ipv6IpamPoolId", Flag: "ipv6-ipam-pool-id", Type: "*string", Required: false},
	{Name: "Ipv6NetmaskLength", Flag: "ipv6-netmask-length", Type: "*int32", Required: false},
	{Name: "Ipv6Pool", Flag: "ipv6-pool", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "VpcEncryptionControl", Flag: "vpc-encryption-control", Type: "*types.VpcEncryptionControlConfiguration", Required: false},
}

var fields_create_vpc_block_public_access_exclusion = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InternetGatewayExclusionMode", Flag: "internet-gateway-exclusion-mode", Type: "types.InternetGatewayExclusionMode", Required: true},
	{Name: "SubnetId", Flag: "subnet-id", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: false},
}

var fields_create_vpc_encryption_control = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_create_vpc_endpoint = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DnsOptions", Flag: "dns-options", Type: "*types.DnsOptionsSpecification", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.IpAddressType", Required: false},
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: false},
	{Name: "PrivateDnsEnabled", Flag: "private-dns-enabled", Type: "*bool", Required: false},
	{Name: "ResourceConfigurationArn", Flag: "resource-configuration-arn", Type: "*string", Required: false},
	{Name: "RouteTableIds", Flag: "route-table-ids", Type: "[]string", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: false},
	{Name: "ServiceNetworkArn", Flag: "service-network-arn", Type: "*string", Required: false},
	{Name: "ServiceRegion", Flag: "service-region", Type: "*string", Required: false},
	{Name: "SubnetConfigurations", Flag: "subnet-configurations", Type: "[]types.SubnetConfiguration", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "VpcEndpointType", Flag: "vpc-endpoint-type", Type: "types.VpcEndpointType", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_create_vpc_endpoint_connection_notification = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ConnectionEvents", Flag: "connection-events", Type: "[]string", Required: true},
	{Name: "ConnectionNotificationArn", Flag: "connection-notification-arn", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ServiceId", Flag: "service-id", Type: "*string", Required: false},
	{Name: "VpcEndpointId", Flag: "vpc-endpoint-id", Type: "*string", Required: false},
}

var fields_create_vpc_endpoint_service_configuration = []leanruntime.Field{
	{Name: "AcceptanceRequired", Flag: "acceptance-required", Type: "*bool", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GatewayLoadBalancerArns", Flag: "gateway-load-balancer-arns", Type: "[]string", Required: false},
	{Name: "NetworkLoadBalancerArns", Flag: "network-load-balancer-arns", Type: "[]string", Required: false},
	{Name: "PrivateDnsName", Flag: "private-dns-name", Type: "*string", Required: false},
	{Name: "SupportedIpAddressTypes", Flag: "supported-ip-address-types", Type: "[]string", Required: false},
	{Name: "SupportedRegions", Flag: "supported-regions", Type: "[]string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_create_vpc_peering_connection = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "PeerOwnerId", Flag: "peer-owner-id", Type: "*string", Required: false},
	{Name: "PeerRegion", Flag: "peer-region", Type: "*string", Required: false},
	{Name: "PeerVpcId", Flag: "peer-vpc-id", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_create_vpn_concentrator = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "TransitGatewayId", Flag: "transit-gateway-id", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.VpnConcentratorType", Required: true},
}

var fields_create_vpn_connection = []leanruntime.Field{
	{Name: "CustomerGatewayId", Flag: "customer-gateway-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Options", Flag: "options", Type: "*types.VpnConnectionOptionsSpecification", Required: false},
	{Name: "PreSharedKeyStorage", Flag: "pre-shared-key-storage", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "TransitGatewayId", Flag: "transit-gateway-id", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "*string", Required: true},
	{Name: "VpnConcentratorId", Flag: "vpn-concentrator-id", Type: "*string", Required: false},
	{Name: "VpnGatewayId", Flag: "vpn-gateway-id", Type: "*string", Required: false},
}

var fields_create_vpn_connection_route = []leanruntime.Field{
	{Name: "DestinationCidrBlock", Flag: "destination-cidr-block", Type: "*string", Required: true},
	{Name: "VpnConnectionId", Flag: "vpn-connection-id", Type: "*string", Required: true},
}

var fields_create_vpn_gateway = []leanruntime.Field{
	{Name: "AmazonSideAsn", Flag: "amazon-side-asn", Type: "*int64", Required: false},
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "Type", Flag: "type", Type: "types.GatewayType", Required: true},
}

var fields_delete_capacity_manager_data_export = []leanruntime.Field{
	{Name: "CapacityManagerDataExportId", Flag: "capacity-manager-data-export-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_delete_carrier_gateway = []leanruntime.Field{
	{Name: "CarrierGatewayId", Flag: "carrier-gateway-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_delete_client_vpn_endpoint = []leanruntime.Field{
	{Name: "ClientVpnEndpointId", Flag: "client-vpn-endpoint-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_delete_client_vpn_route = []leanruntime.Field{
	{Name: "ClientVpnEndpointId", Flag: "client-vpn-endpoint-id", Type: "*string", Required: true},
	{Name: "DestinationCidrBlock", Flag: "destination-cidr-block", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TargetVpcSubnetId", Flag: "target-vpc-subnet-id", Type: "*string", Required: false},
}

var fields_delete_coip_cidr = []leanruntime.Field{
	{Name: "Cidr", Flag: "cidr", Type: "*string", Required: true},
	{Name: "CoipPoolId", Flag: "coip-pool-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_delete_coip_pool = []leanruntime.Field{
	{Name: "CoipPoolId", Flag: "coip-pool-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_delete_customer_gateway = []leanruntime.Field{
	{Name: "CustomerGatewayId", Flag: "customer-gateway-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_delete_dhcp_options = []leanruntime.Field{
	{Name: "DhcpOptionsId", Flag: "dhcp-options-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_delete_egress_only_internet_gateway = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EgressOnlyInternetGatewayId", Flag: "egress-only-internet-gateway-id", Type: "*string", Required: true},
}

var fields_delete_fleets = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "FleetIds", Flag: "fleet-ids", Type: "[]string", Required: true},
	{Name: "TerminateInstances", Flag: "terminate-instances", Type: "*bool", Required: true},
}

var fields_delete_flow_logs = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "FlowLogIds", Flag: "flow-log-ids", Type: "[]string", Required: true},
}

var fields_delete_fpga_image = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "FpgaImageId", Flag: "fpga-image-id", Type: "*string", Required: true},
}

var fields_delete_image_usage_report = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ReportId", Flag: "report-id", Type: "*string", Required: true},
}

var fields_delete_instance_connect_endpoint = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceConnectEndpointId", Flag: "instance-connect-endpoint-id", Type: "*string", Required: true},
}

var fields_delete_instance_event_window = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ForceDelete", Flag: "force-delete", Type: "*bool", Required: false},
	{Name: "InstanceEventWindowId", Flag: "instance-event-window-id", Type: "*string", Required: true},
}

var fields_delete_internet_gateway = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InternetGatewayId", Flag: "internet-gateway-id", Type: "*string", Required: true},
}

var fields_delete_ipam = []leanruntime.Field{
	{Name: "Cascade", Flag: "cascade", Type: "*bool", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamId", Flag: "ipam-id", Type: "*string", Required: true},
}

var fields_delete_ipam_external_resource_verification_token = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamExternalResourceVerificationTokenId", Flag: "ipam-external-resource-verification-token-id", Type: "*string", Required: true},
}

var fields_delete_ipam_policy = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamPolicyId", Flag: "ipam-policy-id", Type: "*string", Required: true},
}

var fields_delete_ipam_pool = []leanruntime.Field{
	{Name: "Cascade", Flag: "cascade", Type: "*bool", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamPoolId", Flag: "ipam-pool-id", Type: "*string", Required: true},
}

var fields_delete_ipam_prefix_list_resolver = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamPrefixListResolverId", Flag: "ipam-prefix-list-resolver-id", Type: "*string", Required: true},
}

var fields_delete_ipam_prefix_list_resolver_target = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamPrefixListResolverTargetId", Flag: "ipam-prefix-list-resolver-target-id", Type: "*string", Required: true},
}

var fields_delete_ipam_resource_discovery = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamResourceDiscoveryId", Flag: "ipam-resource-discovery-id", Type: "*string", Required: true},
}

var fields_delete_ipam_scope = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamScopeId", Flag: "ipam-scope-id", Type: "*string", Required: true},
}

var fields_delete_key_pair = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "KeyName", Flag: "key-name", Type: "*string", Required: false},
	{Name: "KeyPairId", Flag: "key-pair-id", Type: "*string", Required: false},
}

var fields_delete_launch_template = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "LaunchTemplateId", Flag: "launch-template-id", Type: "*string", Required: false},
	{Name: "LaunchTemplateName", Flag: "launch-template-name", Type: "*string", Required: false},
}

var fields_delete_launch_template_versions = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "LaunchTemplateId", Flag: "launch-template-id", Type: "*string", Required: false},
	{Name: "LaunchTemplateName", Flag: "launch-template-name", Type: "*string", Required: false},
	{Name: "Versions", Flag: "versions", Type: "[]string", Required: true},
}

var fields_delete_local_gateway_route = []leanruntime.Field{
	{Name: "DestinationCidrBlock", Flag: "destination-cidr-block", Type: "*string", Required: false},
	{Name: "DestinationPrefixListId", Flag: "destination-prefix-list-id", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "LocalGatewayRouteTableId", Flag: "local-gateway-route-table-id", Type: "*string", Required: true},
}

var fields_delete_local_gateway_route_table = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "LocalGatewayRouteTableId", Flag: "local-gateway-route-table-id", Type: "*string", Required: true},
}

var fields_delete_local_gateway_route_table_virtual_interface_group_association = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "LocalGatewayRouteTableVirtualInterfaceGroupAssociationId", Flag: "local-gateway-route-table-virtual-interface-group-association-id", Type: "*string", Required: true},
}

var fields_delete_local_gateway_route_table_vpc_association = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "LocalGatewayRouteTableVpcAssociationId", Flag: "local-gateway-route-table-vpc-association-id", Type: "*string", Required: true},
}

var fields_delete_local_gateway_virtual_interface = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "LocalGatewayVirtualInterfaceId", Flag: "local-gateway-virtual-interface-id", Type: "*string", Required: true},
}

var fields_delete_local_gateway_virtual_interface_group = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "LocalGatewayVirtualInterfaceGroupId", Flag: "local-gateway-virtual-interface-group-id", Type: "*string", Required: true},
}

var fields_delete_managed_prefix_list = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "PrefixListId", Flag: "prefix-list-id", Type: "*string", Required: true},
}

var fields_delete_nat_gateway = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "NatGatewayId", Flag: "nat-gateway-id", Type: "*string", Required: true},
}

var fields_delete_network_acl = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "NetworkAclId", Flag: "network-acl-id", Type: "*string", Required: true},
}

var fields_delete_network_acl_entry = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Egress", Flag: "egress", Type: "*bool", Required: true},
	{Name: "NetworkAclId", Flag: "network-acl-id", Type: "*string", Required: true},
	{Name: "RuleNumber", Flag: "rule-number", Type: "*int32", Required: true},
}

var fields_delete_network_insights_access_scope = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "NetworkInsightsAccessScopeId", Flag: "network-insights-access-scope-id", Type: "*string", Required: true},
}

var fields_delete_network_insights_access_scope_analysis = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "NetworkInsightsAccessScopeAnalysisId", Flag: "network-insights-access-scope-analysis-id", Type: "*string", Required: true},
}

var fields_delete_network_insights_analysis = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "NetworkInsightsAnalysisId", Flag: "network-insights-analysis-id", Type: "*string", Required: true},
}

var fields_delete_network_insights_path = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "NetworkInsightsPathId", Flag: "network-insights-path-id", Type: "*string", Required: true},
}

var fields_delete_network_interface = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "NetworkInterfaceId", Flag: "network-interface-id", Type: "*string", Required: true},
}

var fields_delete_network_interface_permission = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Force", Flag: "force", Type: "*bool", Required: false},
	{Name: "NetworkInterfacePermissionId", Flag: "network-interface-permission-id", Type: "*string", Required: true},
}

var fields_delete_placement_group = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
}

var fields_delete_public_ipv4_pool = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "NetworkBorderGroup", Flag: "network-border-group", Type: "*string", Required: false},
	{Name: "PoolId", Flag: "pool-id", Type: "*string", Required: true},
}

var fields_delete_queued_reserved_instances = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ReservedInstancesIds", Flag: "reserved-instances-ids", Type: "[]string", Required: true},
}

var fields_delete_route = []leanruntime.Field{
	{Name: "DestinationCidrBlock", Flag: "destination-cidr-block", Type: "*string", Required: false},
	{Name: "DestinationIpv6CidrBlock", Flag: "destination-ipv6-cidr-block", Type: "*string", Required: false},
	{Name: "DestinationPrefixListId", Flag: "destination-prefix-list-id", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "RouteTableId", Flag: "route-table-id", Type: "*string", Required: true},
}

var fields_delete_route_server = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "RouteServerId", Flag: "route-server-id", Type: "*string", Required: true},
}

var fields_delete_route_server_endpoint = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "RouteServerEndpointId", Flag: "route-server-endpoint-id", Type: "*string", Required: true},
}

var fields_delete_route_server_peer = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "RouteServerPeerId", Flag: "route-server-peer-id", Type: "*string", Required: true},
}

var fields_delete_route_table = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "RouteTableId", Flag: "route-table-id", Type: "*string", Required: true},
}

var fields_delete_secondary_network = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "SecondaryNetworkId", Flag: "secondary-network-id", Type: "*string", Required: true},
}

var fields_delete_secondary_subnet = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "SecondarySubnetId", Flag: "secondary-subnet-id", Type: "*string", Required: true},
}

var fields_delete_security_group = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: false},
}

var fields_delete_snapshot = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "SnapshotId", Flag: "snapshot-id", Type: "*string", Required: true},
}

var fields_delete_spot_datafeed_subscription = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_delete_subnet = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "SubnetId", Flag: "subnet-id", Type: "*string", Required: true},
}

var fields_delete_subnet_cidr_reservation = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "SubnetCidrReservationId", Flag: "subnet-cidr-reservation-id", Type: "*string", Required: true},
}

var fields_delete_tags = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Resources", Flag: "resources", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_traffic_mirror_filter = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TrafficMirrorFilterId", Flag: "traffic-mirror-filter-id", Type: "*string", Required: true},
}

var fields_delete_traffic_mirror_filter_rule = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TrafficMirrorFilterRuleId", Flag: "traffic-mirror-filter-rule-id", Type: "*string", Required: true},
}

var fields_delete_traffic_mirror_session = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TrafficMirrorSessionId", Flag: "traffic-mirror-session-id", Type: "*string", Required: true},
}

var fields_delete_traffic_mirror_target = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TrafficMirrorTargetId", Flag: "traffic-mirror-target-id", Type: "*string", Required: true},
}

var fields_delete_transit_gateway = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TransitGatewayId", Flag: "transit-gateway-id", Type: "*string", Required: true},
}

var fields_delete_transit_gateway_connect = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TransitGatewayAttachmentId", Flag: "transit-gateway-attachment-id", Type: "*string", Required: true},
}

var fields_delete_transit_gateway_connect_peer = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TransitGatewayConnectPeerId", Flag: "transit-gateway-connect-peer-id", Type: "*string", Required: true},
}

var fields_delete_transit_gateway_metering_policy = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TransitGatewayMeteringPolicyId", Flag: "transit-gateway-metering-policy-id", Type: "*string", Required: true},
}

var fields_delete_transit_gateway_metering_policy_entry = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "PolicyRuleNumber", Flag: "policy-rule-number", Type: "*int32", Required: true},
	{Name: "TransitGatewayMeteringPolicyId", Flag: "transit-gateway-metering-policy-id", Type: "*string", Required: true},
}

var fields_delete_transit_gateway_multicast_domain = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TransitGatewayMulticastDomainId", Flag: "transit-gateway-multicast-domain-id", Type: "*string", Required: true},
}

var fields_delete_transit_gateway_peering_attachment = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TransitGatewayAttachmentId", Flag: "transit-gateway-attachment-id", Type: "*string", Required: true},
}

var fields_delete_transit_gateway_policy_table = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TransitGatewayPolicyTableId", Flag: "transit-gateway-policy-table-id", Type: "*string", Required: true},
}

var fields_delete_transit_gateway_prefix_list_reference = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "PrefixListId", Flag: "prefix-list-id", Type: "*string", Required: true},
	{Name: "TransitGatewayRouteTableId", Flag: "transit-gateway-route-table-id", Type: "*string", Required: true},
}

var fields_delete_transit_gateway_route = []leanruntime.Field{
	{Name: "DestinationCidrBlock", Flag: "destination-cidr-block", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TransitGatewayRouteTableId", Flag: "transit-gateway-route-table-id", Type: "*string", Required: true},
}

var fields_delete_transit_gateway_route_table = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TransitGatewayRouteTableId", Flag: "transit-gateway-route-table-id", Type: "*string", Required: true},
}

var fields_delete_transit_gateway_route_table_announcement = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TransitGatewayRouteTableAnnouncementId", Flag: "transit-gateway-route-table-announcement-id", Type: "*string", Required: true},
}

var fields_delete_transit_gateway_vpc_attachment = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TransitGatewayAttachmentId", Flag: "transit-gateway-attachment-id", Type: "*string", Required: true},
}

var fields_delete_verified_access_endpoint = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VerifiedAccessEndpointId", Flag: "verified-access-endpoint-id", Type: "*string", Required: true},
}

var fields_delete_verified_access_group = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VerifiedAccessGroupId", Flag: "verified-access-group-id", Type: "*string", Required: true},
}

var fields_delete_verified_access_instance = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VerifiedAccessInstanceId", Flag: "verified-access-instance-id", Type: "*string", Required: true},
}

var fields_delete_verified_access_trust_provider = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VerifiedAccessTrustProviderId", Flag: "verified-access-trust-provider-id", Type: "*string", Required: true},
}

var fields_delete_volume = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VolumeId", Flag: "volume-id", Type: "*string", Required: true},
}

var fields_delete_vpc = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_delete_vpc_block_public_access_exclusion = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ExclusionId", Flag: "exclusion-id", Type: "*string", Required: true},
}

var fields_delete_vpc_encryption_control = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VpcEncryptionControlId", Flag: "vpc-encryption-control-id", Type: "*string", Required: true},
}

var fields_delete_vpc_endpoint_connection_notifications = []leanruntime.Field{
	{Name: "ConnectionNotificationIds", Flag: "connection-notification-ids", Type: "[]string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_delete_vpc_endpoint_service_configurations = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ServiceIds", Flag: "service-ids", Type: "[]string", Required: true},
}

var fields_delete_vpc_endpoints = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VpcEndpointIds", Flag: "vpc-endpoint-ids", Type: "[]string", Required: true},
}

var fields_delete_vpc_peering_connection = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VpcPeeringConnectionId", Flag: "vpc-peering-connection-id", Type: "*string", Required: true},
}

var fields_delete_vpn_concentrator = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VpnConcentratorId", Flag: "vpn-concentrator-id", Type: "*string", Required: true},
}

var fields_delete_vpn_connection = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VpnConnectionId", Flag: "vpn-connection-id", Type: "*string", Required: true},
}

var fields_delete_vpn_connection_route = []leanruntime.Field{
	{Name: "DestinationCidrBlock", Flag: "destination-cidr-block", Type: "*string", Required: true},
	{Name: "VpnConnectionId", Flag: "vpn-connection-id", Type: "*string", Required: true},
}

var fields_delete_vpn_gateway = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VpnGatewayId", Flag: "vpn-gateway-id", Type: "*string", Required: true},
}

var fields_deprovision_byoip_cidr = []leanruntime.Field{
	{Name: "Cidr", Flag: "cidr", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_deprovision_ipam_byoasn = []leanruntime.Field{
	{Name: "Asn", Flag: "asn", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamId", Flag: "ipam-id", Type: "*string", Required: true},
}

var fields_deprovision_ipam_pool_cidr = []leanruntime.Field{
	{Name: "Cidr", Flag: "cidr", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamPoolId", Flag: "ipam-pool-id", Type: "*string", Required: true},
}

var fields_deprovision_public_ipv4_pool_cidr = []leanruntime.Field{
	{Name: "Cidr", Flag: "cidr", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "PoolId", Flag: "pool-id", Type: "*string", Required: true},
}

var fields_deregister_image = []leanruntime.Field{
	{Name: "DeleteAssociatedSnapshots", Flag: "delete-associated-snapshots", Type: "*bool", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: true},
}

var fields_deregister_instance_event_notification_attributes = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceTagAttribute", Flag: "instance-tag-attribute", Type: "*types.DeregisterInstanceTagAttributeRequest", Required: true},
}

var fields_deregister_transit_gateway_multicast_group_members = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GroupIpAddress", Flag: "group-ip-address", Type: "*string", Required: false},
	{Name: "NetworkInterfaceIds", Flag: "network-interface-ids", Type: "[]string", Required: false},
	{Name: "TransitGatewayMulticastDomainId", Flag: "transit-gateway-multicast-domain-id", Type: "*string", Required: false},
}

var fields_deregister_transit_gateway_multicast_group_sources = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GroupIpAddress", Flag: "group-ip-address", Type: "*string", Required: false},
	{Name: "NetworkInterfaceIds", Flag: "network-interface-ids", Type: "[]string", Required: false},
	{Name: "TransitGatewayMulticastDomainId", Flag: "transit-gateway-multicast-domain-id", Type: "*string", Required: false},
}

var fields_describe_account_attributes = []leanruntime.Field{
	{Name: "AttributeNames", Flag: "attribute-names", Type: "[]types.AccountAttributeName", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_describe_address_transfers = []leanruntime.Field{
	{Name: "AllocationIds", Flag: "allocation-ids", Type: "[]string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_addresses = []leanruntime.Field{
	{Name: "AllocationIds", Flag: "allocation-ids", Type: "[]string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "PublicIps", Flag: "public-ips", Type: "[]string", Required: false},
}

var fields_describe_addresses_attribute = []leanruntime.Field{
	{Name: "AllocationIds", Flag: "allocation-ids", Type: "[]string", Required: false},
	{Name: "Attribute", Flag: "attribute", Type: "types.AddressAttributeName", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_aggregate_id_format = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_describe_availability_zones = []leanruntime.Field{
	{Name: "AllAvailabilityZones", Flag: "all-availability-zones", Type: "*bool", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "ZoneIds", Flag: "zone-ids", Type: "[]string", Required: false},
	{Name: "ZoneNames", Flag: "zone-names", Type: "[]string", Required: false},
}

var fields_describe_aws_network_performance_metric_subscriptions = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_bundle_tasks = []leanruntime.Field{
	{Name: "BundleIds", Flag: "bundle-ids", Type: "[]string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
}

var fields_describe_byoip_cidrs = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_capacity_block_extension_history = []leanruntime.Field{
	{Name: "CapacityReservationIds", Flag: "capacity-reservation-ids", Type: "[]string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_capacity_block_extension_offerings = []leanruntime.Field{
	{Name: "CapacityBlockExtensionDurationHours", Flag: "capacity-block-extension-duration-hours", Type: "*int32", Required: true},
	{Name: "CapacityReservationId", Flag: "capacity-reservation-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_capacity_block_offerings = []leanruntime.Field{
	{Name: "AllAvailabilityZones", Flag: "all-availability-zones", Type: "*bool", Required: false},
	{Name: "CapacityDurationHours", Flag: "capacity-duration-hours", Type: "*int32", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EndDateRange", Flag: "end-date-range", Type: "*time.Time", Required: false},
	{Name: "InstanceCount", Flag: "instance-count", Type: "*int32", Required: false},
	{Name: "InstanceType", Flag: "instance-type", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartDateRange", Flag: "start-date-range", Type: "*time.Time", Required: false},
	{Name: "UltraserverCount", Flag: "ultraserver-count", Type: "*int32", Required: false},
	{Name: "UltraserverType", Flag: "ultraserver-type", Type: "*string", Required: false},
}

var fields_describe_capacity_block_status = []leanruntime.Field{
	{Name: "CapacityBlockIds", Flag: "capacity-block-ids", Type: "[]string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_capacity_blocks = []leanruntime.Field{
	{Name: "CapacityBlockIds", Flag: "capacity-block-ids", Type: "[]string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_capacity_manager_data_exports = []leanruntime.Field{
	{Name: "CapacityManagerDataExportIds", Flag: "capacity-manager-data-export-ids", Type: "[]string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_capacity_reservation_billing_requests = []leanruntime.Field{
	{Name: "CapacityReservationIds", Flag: "capacity-reservation-ids", Type: "[]string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Role", Flag: "role", Type: "types.CallerRole", Required: true},
}

var fields_describe_capacity_reservation_fleets = []leanruntime.Field{
	{Name: "CapacityReservationFleetIds", Flag: "capacity-reservation-fleet-ids", Type: "[]string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_capacity_reservation_topology = []leanruntime.Field{
	{Name: "CapacityReservationIds", Flag: "capacity-reservation-ids", Type: "[]string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_capacity_reservations = []leanruntime.Field{
	{Name: "CapacityReservationIds", Flag: "capacity-reservation-ids", Type: "[]string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_carrier_gateways = []leanruntime.Field{
	{Name: "CarrierGatewayIds", Flag: "carrier-gateway-ids", Type: "[]string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_classic_link_instances = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_client_vpn_authorization_rules = []leanruntime.Field{
	{Name: "ClientVpnEndpointId", Flag: "client-vpn-endpoint-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_client_vpn_connections = []leanruntime.Field{
	{Name: "ClientVpnEndpointId", Flag: "client-vpn-endpoint-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_client_vpn_endpoints = []leanruntime.Field{
	{Name: "ClientVpnEndpointIds", Flag: "client-vpn-endpoint-ids", Type: "[]string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_client_vpn_routes = []leanruntime.Field{
	{Name: "ClientVpnEndpointId", Flag: "client-vpn-endpoint-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_client_vpn_target_networks = []leanruntime.Field{
	{Name: "AssociationIds", Flag: "association-ids", Type: "[]string", Required: false},
	{Name: "ClientVpnEndpointId", Flag: "client-vpn-endpoint-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_coip_pools = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PoolIds", Flag: "pool-ids", Type: "[]string", Required: false},
}

var fields_describe_conversion_tasks = []leanruntime.Field{
	{Name: "ConversionTaskIds", Flag: "conversion-task-ids", Type: "[]string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_describe_customer_gateways = []leanruntime.Field{
	{Name: "CustomerGatewayIds", Flag: "customer-gateway-ids", Type: "[]string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
}

var fields_describe_declarative_policies_reports = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReportIds", Flag: "report-ids", Type: "[]string", Required: false},
}

var fields_describe_dhcp_options = []leanruntime.Field{
	{Name: "DhcpOptionsIds", Flag: "dhcp-options-ids", Type: "[]string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_egress_only_internet_gateways = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EgressOnlyInternetGatewayIds", Flag: "egress-only-internet-gateway-ids", Type: "[]string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_elastic_gpus = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ElasticGpuIds", Flag: "elastic-gpu-ids", Type: "[]string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_export_image_tasks = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ExportImageTaskIds", Flag: "export-image-task-ids", Type: "[]string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_export_tasks = []leanruntime.Field{
	{Name: "ExportTaskIds", Flag: "export-task-ids", Type: "[]string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
}

var fields_describe_fast_launch_images = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "ImageIds", Flag: "image-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_fast_snapshot_restores = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_fleet_history = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EventType", Flag: "event-type", Type: "types.FleetEventType", Required: false},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_describe_fleet_instances = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_fleets = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "FleetIds", Flag: "fleet-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_flow_logs = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filter", Flag: "filter", Type: "[]types.Filter", Required: false},
	{Name: "FlowLogIds", Flag: "flow-log-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_fpga_image_attribute = []leanruntime.Field{
	{Name: "Attribute", Flag: "attribute", Type: "types.FpgaImageAttributeName", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "FpgaImageId", Flag: "fpga-image-id", Type: "*string", Required: true},
}

var fields_describe_fpga_images = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "FpgaImageIds", Flag: "fpga-image-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Owners", Flag: "owners", Type: "[]string", Required: false},
}

var fields_describe_host_reservation_offerings = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "[]types.Filter", Required: false},
	{Name: "MaxDuration", Flag: "max-duration", Type: "*int32", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MinDuration", Flag: "min-duration", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OfferingId", Flag: "offering-id", Type: "*string", Required: false},
}

var fields_describe_host_reservations = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "[]types.Filter", Required: false},
	{Name: "HostReservationIdSet", Flag: "host-reservation-id-set", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_hosts = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "[]types.Filter", Required: false},
	{Name: "HostIds", Flag: "host-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_iam_instance_profile_associations = []leanruntime.Field{
	{Name: "AssociationIds", Flag: "association-ids", Type: "[]string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_id_format = []leanruntime.Field{
	{Name: "Resource", Flag: "resource", Type: "*string", Required: false},
}

var fields_describe_identity_id_format = []leanruntime.Field{
	{Name: "PrincipalArn", Flag: "principal-arn", Type: "*string", Required: true},
	{Name: "Resource", Flag: "resource", Type: "*string", Required: false},
}

var fields_describe_image_attribute = []leanruntime.Field{
	{Name: "Attribute", Flag: "attribute", Type: "types.ImageAttributeName", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: true},
}

var fields_describe_image_references = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ImageIds", Flag: "image-ids", Type: "[]string", Required: true},
	{Name: "IncludeAllResourceTypes", Flag: "include-all-resource-types", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceTypes", Flag: "resource-types", Type: "[]types.ResourceTypeRequest", Required: false},
}

var fields_describe_image_usage_report_entries = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "ImageIds", Flag: "image-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReportIds", Flag: "report-ids", Type: "[]string", Required: false},
}

var fields_describe_image_usage_reports = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "ImageIds", Flag: "image-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReportIds", Flag: "report-ids", Type: "[]string", Required: false},
}

var fields_describe_images = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ExecutableUsers", Flag: "executable-users", Type: "[]string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "ImageIds", Flag: "image-ids", Type: "[]string", Required: false},
	{Name: "IncludeDeprecated", Flag: "include-deprecated", Type: "*bool", Required: false},
	{Name: "IncludeDisabled", Flag: "include-disabled", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Owners", Flag: "owners", Type: "[]string", Required: false},
}

var fields_describe_import_image_tasks = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "ImportTaskIds", Flag: "import-task-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_import_snapshot_tasks = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "ImportTaskIds", Flag: "import-task-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_instance_attribute = []leanruntime.Field{
	{Name: "Attribute", Flag: "attribute", Type: "types.InstanceAttributeName", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_describe_instance_connect_endpoints = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "InstanceConnectEndpointIds", Flag: "instance-connect-endpoint-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_instance_credit_specifications = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_instance_event_notification_attributes = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_describe_instance_event_windows = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "InstanceEventWindowIds", Flag: "instance-event-window-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_instance_image_metadata = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_instance_sql_ha_history_states = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_describe_instance_sql_ha_states = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_instance_status = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IncludeAllInstances", Flag: "include-all-instances", Type: "*bool", Required: false},
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_instance_topology = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "GroupNames", Flag: "group-names", Type: "[]string", Required: false},
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_instance_type_offerings = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "LocationType", Flag: "location-type", Type: "types.LocationType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_instance_types = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "InstanceTypes", Flag: "instance-types", Type: "[]types.InstanceType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_instances = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_internet_gateways = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "InternetGatewayIds", Flag: "internet-gateway-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_ipam_byoasn = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_ipam_external_resource_verification_tokens = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IpamExternalResourceVerificationTokenIds", Flag: "ipam-external-resource-verification-token-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_ipam_policies = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IpamPolicyIds", Flag: "ipam-policy-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_ipam_pools = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IpamPoolIds", Flag: "ipam-pool-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_ipam_prefix_list_resolver_targets = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IpamPrefixListResolverId", Flag: "ipam-prefix-list-resolver-id", Type: "*string", Required: false},
	{Name: "IpamPrefixListResolverTargetIds", Flag: "ipam-prefix-list-resolver-target-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_ipam_prefix_list_resolvers = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IpamPrefixListResolverIds", Flag: "ipam-prefix-list-resolver-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_ipam_resource_discoveries = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IpamResourceDiscoveryIds", Flag: "ipam-resource-discovery-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_ipam_resource_discovery_associations = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IpamResourceDiscoveryAssociationIds", Flag: "ipam-resource-discovery-association-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_ipam_scopes = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IpamScopeIds", Flag: "ipam-scope-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_ipams = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IpamIds", Flag: "ipam-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_ipv6_pools = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PoolIds", Flag: "pool-ids", Type: "[]string", Required: false},
}

var fields_describe_key_pairs = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IncludePublicKey", Flag: "include-public-key", Type: "*bool", Required: false},
	{Name: "KeyNames", Flag: "key-names", Type: "[]string", Required: false},
	{Name: "KeyPairIds", Flag: "key-pair-ids", Type: "[]string", Required: false},
}

var fields_describe_launch_template_versions = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "LaunchTemplateId", Flag: "launch-template-id", Type: "*string", Required: false},
	{Name: "LaunchTemplateName", Flag: "launch-template-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MaxVersion", Flag: "max-version", Type: "*string", Required: false},
	{Name: "MinVersion", Flag: "min-version", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResolveAlias", Flag: "resolve-alias", Type: "*bool", Required: false},
	{Name: "Versions", Flag: "versions", Type: "[]string", Required: false},
}

var fields_describe_launch_templates = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "LaunchTemplateIds", Flag: "launch-template-ids", Type: "[]string", Required: false},
	{Name: "LaunchTemplateNames", Flag: "launch-template-names", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_local_gateway_route_table_virtual_interface_group_associations = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "LocalGatewayRouteTableVirtualInterfaceGroupAssociationIds", Flag: "local-gateway-route-table-virtual-interface-group-association-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_local_gateway_route_table_vpc_associations = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "LocalGatewayRouteTableVpcAssociationIds", Flag: "local-gateway-route-table-vpc-association-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_local_gateway_route_tables = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "LocalGatewayRouteTableIds", Flag: "local-gateway-route-table-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_local_gateway_virtual_interface_groups = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "LocalGatewayVirtualInterfaceGroupIds", Flag: "local-gateway-virtual-interface-group-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_local_gateway_virtual_interfaces = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "LocalGatewayVirtualInterfaceIds", Flag: "local-gateway-virtual-interface-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_local_gateways = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "LocalGatewayIds", Flag: "local-gateway-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_locked_snapshots = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SnapshotIds", Flag: "snapshot-ids", Type: "[]string", Required: false},
}

var fields_describe_mac_hosts = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "HostIds", Flag: "host-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_mac_modification_tasks = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MacModificationTaskIds", Flag: "mac-modification-task-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_managed_prefix_lists = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PrefixListIds", Flag: "prefix-list-ids", Type: "[]string", Required: false},
}

var fields_describe_moving_addresses = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PublicIps", Flag: "public-ips", Type: "[]string", Required: false},
}

var fields_describe_nat_gateways = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filter", Flag: "filter", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NatGatewayIds", Flag: "nat-gateway-ids", Type: "[]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_network_acls = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NetworkAclIds", Flag: "network-acl-ids", Type: "[]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_network_insights_access_scope_analyses = []leanruntime.Field{
	{Name: "AnalysisStartTimeBegin", Flag: "analysis-start-time-begin", Type: "*time.Time", Required: false},
	{Name: "AnalysisStartTimeEnd", Flag: "analysis-start-time-end", Type: "*time.Time", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NetworkInsightsAccessScopeAnalysisIds", Flag: "network-insights-access-scope-analysis-ids", Type: "[]string", Required: false},
	{Name: "NetworkInsightsAccessScopeId", Flag: "network-insights-access-scope-id", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_network_insights_access_scopes = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NetworkInsightsAccessScopeIds", Flag: "network-insights-access-scope-ids", Type: "[]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_network_insights_analyses = []leanruntime.Field{
	{Name: "AnalysisEndTime", Flag: "analysis-end-time", Type: "*time.Time", Required: false},
	{Name: "AnalysisStartTime", Flag: "analysis-start-time", Type: "*time.Time", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NetworkInsightsAnalysisIds", Flag: "network-insights-analysis-ids", Type: "[]string", Required: false},
	{Name: "NetworkInsightsPathId", Flag: "network-insights-path-id", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_network_insights_paths = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NetworkInsightsPathIds", Flag: "network-insights-path-ids", Type: "[]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_network_interface_attribute = []leanruntime.Field{
	{Name: "Attribute", Flag: "attribute", Type: "types.NetworkInterfaceAttribute", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "NetworkInterfaceId", Flag: "network-interface-id", Type: "*string", Required: true},
}

var fields_describe_network_interface_permissions = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NetworkInterfacePermissionIds", Flag: "network-interface-permission-ids", Type: "[]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_network_interfaces = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NetworkInterfaceIds", Flag: "network-interface-ids", Type: "[]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_outpost_lags = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OutpostLagIds", Flag: "outpost-lag-ids", Type: "[]string", Required: false},
}

var fields_describe_placement_groups = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "GroupIds", Flag: "group-ids", Type: "[]string", Required: false},
	{Name: "GroupNames", Flag: "group-names", Type: "[]string", Required: false},
}

var fields_describe_prefix_lists = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PrefixListIds", Flag: "prefix-list-ids", Type: "[]string", Required: false},
}

var fields_describe_principal_id_format = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Resources", Flag: "resources", Type: "[]string", Required: false},
}

var fields_describe_public_ipv4_pools = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PoolIds", Flag: "pool-ids", Type: "[]string", Required: false},
}

var fields_describe_regions = []leanruntime.Field{
	{Name: "AllRegions", Flag: "all-regions", Type: "*bool", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "RegionNames", Flag: "region-names", Type: "[]string", Required: false},
}

var fields_describe_replace_root_volume_tasks = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReplaceRootVolumeTaskIds", Flag: "replace-root-volume-task-ids", Type: "[]string", Required: false},
}

var fields_describe_reserved_instances = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "OfferingClass", Flag: "offering-class", Type: "types.OfferingClassType", Required: false},
	{Name: "OfferingType", Flag: "offering-type", Type: "types.OfferingTypeValues", Required: false},
	{Name: "ReservedInstancesIds", Flag: "reserved-instances-ids", Type: "[]string", Required: false},
}

var fields_describe_reserved_instances_listings = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "ReservedInstancesId", Flag: "reserved-instances-id", Type: "*string", Required: false},
	{Name: "ReservedInstancesListingId", Flag: "reserved-instances-listing-id", Type: "*string", Required: false},
}

var fields_describe_reserved_instances_modifications = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReservedInstancesModificationIds", Flag: "reserved-instances-modification-ids", Type: "[]string", Required: false},
}

var fields_describe_reserved_instances_offerings = []leanruntime.Field{
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "AvailabilityZoneId", Flag: "availability-zone-id", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IncludeMarketplace", Flag: "include-marketplace", Type: "*bool", Required: false},
	{Name: "InstanceTenancy", Flag: "instance-tenancy", Type: "types.Tenancy", Required: false},
	{Name: "InstanceType", Flag: "instance-type", Type: "types.InstanceType", Required: false},
	{Name: "MaxDuration", Flag: "max-duration", Type: "*int64", Required: false},
	{Name: "MaxInstanceCount", Flag: "max-instance-count", Type: "*int32", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MinDuration", Flag: "min-duration", Type: "*int64", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OfferingClass", Flag: "offering-class", Type: "types.OfferingClassType", Required: false},
	{Name: "OfferingType", Flag: "offering-type", Type: "types.OfferingTypeValues", Required: false},
	{Name: "ProductDescription", Flag: "product-description", Type: "types.RIProductDescription", Required: false},
	{Name: "ReservedInstancesOfferingIds", Flag: "reserved-instances-offering-ids", Type: "[]string", Required: false},
}

var fields_describe_route_server_endpoints = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RouteServerEndpointIds", Flag: "route-server-endpoint-ids", Type: "[]string", Required: false},
}

var fields_describe_route_server_peers = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RouteServerPeerIds", Flag: "route-server-peer-ids", Type: "[]string", Required: false},
}

var fields_describe_route_servers = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RouteServerIds", Flag: "route-server-ids", Type: "[]string", Required: false},
}

var fields_describe_route_tables = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RouteTableIds", Flag: "route-table-ids", Type: "[]string", Required: false},
}

var fields_describe_scheduled_instance_availability = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "FirstSlotStartTimeRange", Flag: "first-slot-start-time-range", Type: "*types.SlotDateTimeRangeRequest", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MaxSlotDurationInHours", Flag: "max-slot-duration-in-hours", Type: "*int32", Required: false},
	{Name: "MinSlotDurationInHours", Flag: "min-slot-duration-in-hours", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Recurrence", Flag: "recurrence", Type: "*types.ScheduledInstanceRecurrenceRequest", Required: true},
}

var fields_describe_scheduled_instances = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ScheduledInstanceIds", Flag: "scheduled-instance-ids", Type: "[]string", Required: false},
	{Name: "SlotStartTimeRange", Flag: "slot-start-time-range", Type: "*types.SlotStartTimeRangeRequest", Required: false},
}

var fields_describe_secondary_interfaces = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SecondaryInterfaceIds", Flag: "secondary-interface-ids", Type: "[]string", Required: false},
}

var fields_describe_secondary_networks = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SecondaryNetworkIds", Flag: "secondary-network-ids", Type: "[]string", Required: false},
}

var fields_describe_secondary_subnets = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SecondarySubnetIds", Flag: "secondary-subnet-ids", Type: "[]string", Required: false},
}

var fields_describe_security_group_references = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GroupId", Flag: "group-id", Type: "[]string", Required: true},
}

var fields_describe_security_group_rules = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SecurityGroupRuleIds", Flag: "security-group-rule-ids", Type: "[]string", Required: false},
}

var fields_describe_security_group_vpc_associations = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_security_groups = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "GroupIds", Flag: "group-ids", Type: "[]string", Required: false},
	{Name: "GroupNames", Flag: "group-names", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_service_link_virtual_interfaces = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceLinkVirtualInterfaceIds", Flag: "service-link-virtual-interface-ids", Type: "[]string", Required: false},
}

var fields_describe_snapshot_attribute = []leanruntime.Field{
	{Name: "Attribute", Flag: "attribute", Type: "types.SnapshotAttributeName", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "SnapshotId", Flag: "snapshot-id", Type: "*string", Required: true},
}

var fields_describe_snapshot_tier_status = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_snapshots = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OwnerIds", Flag: "owner-ids", Type: "[]string", Required: false},
	{Name: "RestorableByUserIds", Flag: "restorable-by-user-ids", Type: "[]string", Required: false},
	{Name: "SnapshotIds", Flag: "snapshot-ids", Type: "[]string", Required: false},
}

var fields_describe_spot_datafeed_subscription = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_describe_spot_fleet_instances = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SpotFleetRequestId", Flag: "spot-fleet-request-id", Type: "*string", Required: true},
}

var fields_describe_spot_fleet_request_history = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EventType", Flag: "event-type", Type: "types.EventType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SpotFleetRequestId", Flag: "spot-fleet-request-id", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_describe_spot_fleet_requests = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SpotFleetRequestIds", Flag: "spot-fleet-request-ids", Type: "[]string", Required: false},
}

var fields_describe_spot_instance_requests = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SpotInstanceRequestIds", Flag: "spot-instance-request-ids", Type: "[]string", Required: false},
}

var fields_describe_spot_price_history = []leanruntime.Field{
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "AvailabilityZoneId", Flag: "availability-zone-id", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "InstanceTypes", Flag: "instance-types", Type: "[]types.InstanceType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProductDescriptions", Flag: "product-descriptions", Type: "[]string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_describe_stale_security_groups = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_describe_store_image_tasks = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "ImageIds", Flag: "image-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_subnets = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: false},
}

var fields_describe_tags = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_traffic_mirror_filter_rules = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TrafficMirrorFilterId", Flag: "traffic-mirror-filter-id", Type: "*string", Required: false},
	{Name: "TrafficMirrorFilterRuleIds", Flag: "traffic-mirror-filter-rule-ids", Type: "[]string", Required: false},
}

var fields_describe_traffic_mirror_filters = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TrafficMirrorFilterIds", Flag: "traffic-mirror-filter-ids", Type: "[]string", Required: false},
}

var fields_describe_traffic_mirror_sessions = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TrafficMirrorSessionIds", Flag: "traffic-mirror-session-ids", Type: "[]string", Required: false},
}

var fields_describe_traffic_mirror_targets = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TrafficMirrorTargetIds", Flag: "traffic-mirror-target-ids", Type: "[]string", Required: false},
}

var fields_describe_transit_gateway_attachments = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TransitGatewayAttachmentIds", Flag: "transit-gateway-attachment-ids", Type: "[]string", Required: false},
}

var fields_describe_transit_gateway_connect_peers = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TransitGatewayConnectPeerIds", Flag: "transit-gateway-connect-peer-ids", Type: "[]string", Required: false},
}

var fields_describe_transit_gateway_connects = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TransitGatewayAttachmentIds", Flag: "transit-gateway-attachment-ids", Type: "[]string", Required: false},
}

var fields_describe_transit_gateway_metering_policies = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TransitGatewayMeteringPolicyIds", Flag: "transit-gateway-metering-policy-ids", Type: "[]string", Required: false},
}

var fields_describe_transit_gateway_multicast_domains = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TransitGatewayMulticastDomainIds", Flag: "transit-gateway-multicast-domain-ids", Type: "[]string", Required: false},
}

var fields_describe_transit_gateway_peering_attachments = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TransitGatewayAttachmentIds", Flag: "transit-gateway-attachment-ids", Type: "[]string", Required: false},
}

var fields_describe_transit_gateway_policy_tables = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TransitGatewayPolicyTableIds", Flag: "transit-gateway-policy-table-ids", Type: "[]string", Required: false},
}

var fields_describe_transit_gateway_route_table_announcements = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TransitGatewayRouteTableAnnouncementIds", Flag: "transit-gateway-route-table-announcement-ids", Type: "[]string", Required: false},
}

var fields_describe_transit_gateway_route_tables = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TransitGatewayRouteTableIds", Flag: "transit-gateway-route-table-ids", Type: "[]string", Required: false},
}

var fields_describe_transit_gateway_vpc_attachments = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TransitGatewayAttachmentIds", Flag: "transit-gateway-attachment-ids", Type: "[]string", Required: false},
}

var fields_describe_transit_gateways = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TransitGatewayIds", Flag: "transit-gateway-ids", Type: "[]string", Required: false},
}

var fields_describe_trunk_interface_associations = []leanruntime.Field{
	{Name: "AssociationIds", Flag: "association-ids", Type: "[]string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_verified_access_endpoints = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VerifiedAccessEndpointIds", Flag: "verified-access-endpoint-ids", Type: "[]string", Required: false},
	{Name: "VerifiedAccessGroupId", Flag: "verified-access-group-id", Type: "*string", Required: false},
	{Name: "VerifiedAccessInstanceId", Flag: "verified-access-instance-id", Type: "*string", Required: false},
}

var fields_describe_verified_access_groups = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VerifiedAccessGroupIds", Flag: "verified-access-group-ids", Type: "[]string", Required: false},
	{Name: "VerifiedAccessInstanceId", Flag: "verified-access-instance-id", Type: "*string", Required: false},
}

var fields_describe_verified_access_instance_logging_configurations = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VerifiedAccessInstanceIds", Flag: "verified-access-instance-ids", Type: "[]string", Required: false},
}

var fields_describe_verified_access_instances = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VerifiedAccessInstanceIds", Flag: "verified-access-instance-ids", Type: "[]string", Required: false},
}

var fields_describe_verified_access_trust_providers = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VerifiedAccessTrustProviderIds", Flag: "verified-access-trust-provider-ids", Type: "[]string", Required: false},
}

var fields_describe_volume_attribute = []leanruntime.Field{
	{Name: "Attribute", Flag: "attribute", Type: "types.VolumeAttributeName", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VolumeId", Flag: "volume-id", Type: "*string", Required: true},
}

var fields_describe_volume_status = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VolumeIds", Flag: "volume-ids", Type: "[]string", Required: false},
}

var fields_describe_volumes = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VolumeIds", Flag: "volume-ids", Type: "[]string", Required: false},
}

var fields_describe_volumes_modifications = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VolumeIds", Flag: "volume-ids", Type: "[]string", Required: false},
}

var fields_describe_vpc_attribute = []leanruntime.Field{
	{Name: "Attribute", Flag: "attribute", Type: "types.VpcAttributeName", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_describe_vpc_block_public_access_exclusions = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ExclusionIds", Flag: "exclusion-ids", Type: "[]string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_vpc_block_public_access_options = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_describe_vpc_classic_link = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "VpcIds", Flag: "vpc-ids", Type: "[]string", Required: false},
}

var fields_describe_vpc_classic_link_dns_support = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VpcIds", Flag: "vpc-ids", Type: "[]string", Required: false},
}

var fields_describe_vpc_encryption_controls = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VpcEncryptionControlIds", Flag: "vpc-encryption-control-ids", Type: "[]string", Required: false},
	{Name: "VpcIds", Flag: "vpc-ids", Type: "[]string", Required: false},
}

var fields_describe_vpc_endpoint_associations = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VpcEndpointIds", Flag: "vpc-endpoint-ids", Type: "[]string", Required: false},
}

var fields_describe_vpc_endpoint_connection_notifications = []leanruntime.Field{
	{Name: "ConnectionNotificationId", Flag: "connection-notification-id", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_vpc_endpoint_connections = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_vpc_endpoint_service_configurations = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceIds", Flag: "service-ids", Type: "[]string", Required: false},
}

var fields_describe_vpc_endpoint_service_permissions = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceId", Flag: "service-id", Type: "*string", Required: true},
}

var fields_describe_vpc_endpoint_services = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceNames", Flag: "service-names", Type: "[]string", Required: false},
	{Name: "ServiceRegions", Flag: "service-regions", Type: "[]string", Required: false},
}

var fields_describe_vpc_endpoints = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VpcEndpointIds", Flag: "vpc-endpoint-ids", Type: "[]string", Required: false},
}

var fields_describe_vpc_peering_connections = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VpcPeeringConnectionIds", Flag: "vpc-peering-connection-ids", Type: "[]string", Required: false},
}

var fields_describe_vpcs = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VpcIds", Flag: "vpc-ids", Type: "[]string", Required: false},
}

var fields_describe_vpn_concentrators = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VpnConcentratorIds", Flag: "vpn-concentrator-ids", Type: "[]string", Required: false},
}

var fields_describe_vpn_connections = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "VpnConnectionIds", Flag: "vpn-connection-ids", Type: "[]string", Required: false},
}

var fields_describe_vpn_gateways = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "VpnGatewayIds", Flag: "vpn-gateway-ids", Type: "[]string", Required: false},
}

var fields_detach_classic_link_vpc = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_detach_internet_gateway = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InternetGatewayId", Flag: "internet-gateway-id", Type: "*string", Required: true},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_detach_network_interface = []leanruntime.Field{
	{Name: "AttachmentId", Flag: "attachment-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Force", Flag: "force", Type: "*bool", Required: false},
}

var fields_detach_verified_access_trust_provider = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VerifiedAccessInstanceId", Flag: "verified-access-instance-id", Type: "*string", Required: true},
	{Name: "VerifiedAccessTrustProviderId", Flag: "verified-access-trust-provider-id", Type: "*string", Required: true},
}

var fields_detach_volume = []leanruntime.Field{
	{Name: "Device", Flag: "device", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Force", Flag: "force", Type: "*bool", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: false},
	{Name: "VolumeId", Flag: "volume-id", Type: "*string", Required: true},
}

var fields_detach_vpn_gateway = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
	{Name: "VpnGatewayId", Flag: "vpn-gateway-id", Type: "*string", Required: true},
}

var fields_disable_address_transfer = []leanruntime.Field{
	{Name: "AllocationId", Flag: "allocation-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_disable_allowed_images_settings = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_disable_aws_network_performance_metric_subscription = []leanruntime.Field{
	{Name: "Destination", Flag: "destination", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Metric", Flag: "metric", Type: "types.MetricType", Required: false},
	{Name: "Source", Flag: "source", Type: "*string", Required: false},
	{Name: "Statistic", Flag: "statistic", Type: "types.StatisticType", Required: false},
}

var fields_disable_capacity_manager = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_disable_ebs_encryption_by_default = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_disable_fast_launch = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Force", Flag: "force", Type: "*bool", Required: false},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: true},
}

var fields_disable_fast_snapshot_restores = []leanruntime.Field{
	{Name: "AvailabilityZoneIds", Flag: "availability-zone-ids", Type: "[]string", Required: false},
	{Name: "AvailabilityZones", Flag: "availability-zones", Type: "[]string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "SourceSnapshotIds", Flag: "source-snapshot-ids", Type: "[]string", Required: true},
}

var fields_disable_image = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: true},
}

var fields_disable_image_block_public_access = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_disable_image_deprecation = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: true},
}

var fields_disable_image_deregistration_protection = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: true},
}

var fields_disable_instance_sql_ha_standby_detections = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: true},
}

var fields_disable_ipam_organization_admin_account = []leanruntime.Field{
	{Name: "DelegatedAdminAccountId", Flag: "delegated-admin-account-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_disable_ipam_policy = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamPolicyId", Flag: "ipam-policy-id", Type: "*string", Required: true},
	{Name: "OrganizationTargetId", Flag: "organization-target-id", Type: "*string", Required: false},
}

var fields_disable_route_server_propagation = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "RouteServerId", Flag: "route-server-id", Type: "*string", Required: true},
	{Name: "RouteTableId", Flag: "route-table-id", Type: "*string", Required: true},
}

var fields_disable_serial_console_access = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_disable_snapshot_block_public_access = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_disable_transit_gateway_route_table_propagation = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TransitGatewayAttachmentId", Flag: "transit-gateway-attachment-id", Type: "*string", Required: false},
	{Name: "TransitGatewayRouteTableAnnouncementId", Flag: "transit-gateway-route-table-announcement-id", Type: "*string", Required: false},
	{Name: "TransitGatewayRouteTableId", Flag: "transit-gateway-route-table-id", Type: "*string", Required: true},
}

var fields_disable_vgw_route_propagation = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: true},
	{Name: "RouteTableId", Flag: "route-table-id", Type: "*string", Required: true},
}

var fields_disable_vpc_classic_link = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_disable_vpc_classic_link_dns_support = []leanruntime.Field{
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: false},
}

var fields_disassociate_address = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "PublicIp", Flag: "public-ip", Type: "*string", Required: false},
}

var fields_disassociate_capacity_reservation_billing_owner = []leanruntime.Field{
	{Name: "CapacityReservationId", Flag: "capacity-reservation-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "UnusedReservationBillingOwnerId", Flag: "unused-reservation-billing-owner-id", Type: "*string", Required: true},
}

var fields_disassociate_client_vpn_target_network = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
	{Name: "ClientVpnEndpointId", Flag: "client-vpn-endpoint-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_disassociate_enclave_certificate_iam_role = []leanruntime.Field{
	{Name: "CertificateArn", Flag: "certificate-arn", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
}

var fields_disassociate_iam_instance_profile = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
}

var fields_disassociate_instance_event_window = []leanruntime.Field{
	{Name: "AssociationTarget", Flag: "association-target", Type: "*types.InstanceEventWindowDisassociationRequest", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceEventWindowId", Flag: "instance-event-window-id", Type: "*string", Required: true},
}

var fields_disassociate_ipam_byoasn = []leanruntime.Field{
	{Name: "Asn", Flag: "asn", Type: "*string", Required: true},
	{Name: "Cidr", Flag: "cidr", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_disassociate_ipam_resource_discovery = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamResourceDiscoveryAssociationId", Flag: "ipam-resource-discovery-association-id", Type: "*string", Required: true},
}

var fields_disassociate_nat_gateway_address = []leanruntime.Field{
	{Name: "AssociationIds", Flag: "association-ids", Type: "[]string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "MaxDrainDurationSeconds", Flag: "max-drain-duration-seconds", Type: "*int32", Required: false},
	{Name: "NatGatewayId", Flag: "nat-gateway-id", Type: "*string", Required: true},
}

var fields_disassociate_route_server = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "RouteServerId", Flag: "route-server-id", Type: "*string", Required: true},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_disassociate_route_table = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_disassociate_security_group_vpc = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_disassociate_subnet_cidr_block = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
}

var fields_disassociate_transit_gateway_multicast_domain = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: true},
	{Name: "TransitGatewayAttachmentId", Flag: "transit-gateway-attachment-id", Type: "*string", Required: true},
	{Name: "TransitGatewayMulticastDomainId", Flag: "transit-gateway-multicast-domain-id", Type: "*string", Required: true},
}

var fields_disassociate_transit_gateway_policy_table = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TransitGatewayAttachmentId", Flag: "transit-gateway-attachment-id", Type: "*string", Required: true},
	{Name: "TransitGatewayPolicyTableId", Flag: "transit-gateway-policy-table-id", Type: "*string", Required: true},
}

var fields_disassociate_transit_gateway_route_table = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TransitGatewayAttachmentId", Flag: "transit-gateway-attachment-id", Type: "*string", Required: true},
	{Name: "TransitGatewayRouteTableId", Flag: "transit-gateway-route-table-id", Type: "*string", Required: true},
}

var fields_disassociate_trunk_interface = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_disassociate_vpc_cidr_block = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
}

var fields_enable_address_transfer = []leanruntime.Field{
	{Name: "AllocationId", Flag: "allocation-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TransferAccountId", Flag: "transfer-account-id", Type: "*string", Required: true},
}

var fields_enable_allowed_images_settings = []leanruntime.Field{
	{Name: "AllowedImagesSettingsState", Flag: "allowed-images-settings-state", Type: "types.AllowedImagesSettingsEnabledState", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_enable_aws_network_performance_metric_subscription = []leanruntime.Field{
	{Name: "Destination", Flag: "destination", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Metric", Flag: "metric", Type: "types.MetricType", Required: false},
	{Name: "Source", Flag: "source", Type: "*string", Required: false},
	{Name: "Statistic", Flag: "statistic", Type: "types.StatisticType", Required: false},
}

var fields_enable_capacity_manager = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "OrganizationsAccess", Flag: "organizations-access", Type: "*bool", Required: false},
}

var fields_enable_ebs_encryption_by_default = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_enable_fast_launch = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: true},
	{Name: "LaunchTemplate", Flag: "launch-template", Type: "*types.FastLaunchLaunchTemplateSpecificationRequest", Required: false},
	{Name: "MaxParallelLaunches", Flag: "max-parallel-launches", Type: "*int32", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: false},
	{Name: "SnapshotConfiguration", Flag: "snapshot-configuration", Type: "*types.FastLaunchSnapshotConfigurationRequest", Required: false},
}

var fields_enable_fast_snapshot_restores = []leanruntime.Field{
	{Name: "AvailabilityZoneIds", Flag: "availability-zone-ids", Type: "[]string", Required: false},
	{Name: "AvailabilityZones", Flag: "availability-zones", Type: "[]string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "SourceSnapshotIds", Flag: "source-snapshot-ids", Type: "[]string", Required: true},
}

var fields_enable_image = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: true},
}

var fields_enable_image_block_public_access = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ImageBlockPublicAccessState", Flag: "image-block-public-access-state", Type: "types.ImageBlockPublicAccessEnabledState", Required: true},
}

var fields_enable_image_deprecation = []leanruntime.Field{
	{Name: "DeprecateAt", Flag: "deprecate-at", Type: "*time.Time", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: true},
}

var fields_enable_image_deregistration_protection = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: true},
	{Name: "WithCooldown", Flag: "with-cooldown", Type: "*bool", Required: false},
}

var fields_enable_instance_sql_ha_standby_detections = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: true},
	{Name: "SqlServerCredentials", Flag: "sql-server-credentials", Type: "*string", Required: false},
}

var fields_enable_ipam_organization_admin_account = []leanruntime.Field{
	{Name: "DelegatedAdminAccountId", Flag: "delegated-admin-account-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_enable_ipam_policy = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamPolicyId", Flag: "ipam-policy-id", Type: "*string", Required: true},
	{Name: "OrganizationTargetId", Flag: "organization-target-id", Type: "*string", Required: false},
}

var fields_enable_reachability_analyzer_organization_sharing = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_enable_route_server_propagation = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "RouteServerId", Flag: "route-server-id", Type: "*string", Required: true},
	{Name: "RouteTableId", Flag: "route-table-id", Type: "*string", Required: true},
}

var fields_enable_serial_console_access = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_enable_snapshot_block_public_access = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "State", Flag: "state", Type: "types.SnapshotBlockPublicAccessState", Required: true},
}

var fields_enable_transit_gateway_route_table_propagation = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TransitGatewayAttachmentId", Flag: "transit-gateway-attachment-id", Type: "*string", Required: false},
	{Name: "TransitGatewayRouteTableAnnouncementId", Flag: "transit-gateway-route-table-announcement-id", Type: "*string", Required: false},
	{Name: "TransitGatewayRouteTableId", Flag: "transit-gateway-route-table-id", Type: "*string", Required: true},
}

var fields_enable_vgw_route_propagation = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: true},
	{Name: "RouteTableId", Flag: "route-table-id", Type: "*string", Required: true},
}

var fields_enable_volume_io = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VolumeId", Flag: "volume-id", Type: "*string", Required: true},
}

var fields_enable_vpc_classic_link = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_enable_vpc_classic_link_dns_support = []leanruntime.Field{
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: false},
}

var fields_export_client_vpn_client_certificate_revocation_list = []leanruntime.Field{
	{Name: "ClientVpnEndpointId", Flag: "client-vpn-endpoint-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_export_client_vpn_client_configuration = []leanruntime.Field{
	{Name: "ClientVpnEndpointId", Flag: "client-vpn-endpoint-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_export_image = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DiskImageFormat", Flag: "disk-image-format", Type: "types.DiskImageFormat", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: true},
	{Name: "RoleName", Flag: "role-name", Type: "*string", Required: false},
	{Name: "S3ExportLocation", Flag: "s3-export-location", Type: "*types.ExportTaskS3LocationRequest", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_export_transit_gateway_routes = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "S3Bucket", Flag: "s3-bucket", Type: "*string", Required: true},
	{Name: "TransitGatewayRouteTableId", Flag: "transit-gateway-route-table-id", Type: "*string", Required: true},
}

var fields_export_verified_access_instance_client_configuration = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VerifiedAccessInstanceId", Flag: "verified-access-instance-id", Type: "*string", Required: true},
}

var fields_get_active_vpn_tunnel_status = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VpnConnectionId", Flag: "vpn-connection-id", Type: "*string", Required: true},
	{Name: "VpnTunnelOutsideIpAddress", Flag: "vpn-tunnel-outside-ip-address", Type: "*string", Required: true},
}

var fields_get_allowed_images_settings = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_get_associated_enclave_certificate_iam_roles = []leanruntime.Field{
	{Name: "CertificateArn", Flag: "certificate-arn", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_get_associated_ipv6_pool_cidrs = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PoolId", Flag: "pool-id", Type: "*string", Required: true},
}

var fields_get_aws_network_performance_data = []leanruntime.Field{
	{Name: "DataQueries", Flag: "data-queries", Type: "[]types.DataQuery", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_get_capacity_manager_attributes = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_get_capacity_manager_metric_data = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "FilterBy", Flag: "filter-by", Type: "[]types.CapacityManagerCondition", Required: false},
	{Name: "GroupBy", Flag: "group-by", Type: "[]types.GroupBy", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MetricNames", Flag: "metric-names", Type: "[]types.Metric", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Period", Flag: "period", Type: "*int32", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_get_capacity_manager_metric_dimensions = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "FilterBy", Flag: "filter-by", Type: "[]types.CapacityManagerCondition", Required: false},
	{Name: "GroupBy", Flag: "group-by", Type: "[]types.GroupBy", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MetricNames", Flag: "metric-names", Type: "[]types.Metric", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_get_capacity_reservation_usage = []leanruntime.Field{
	{Name: "CapacityReservationId", Flag: "capacity-reservation-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_coip_pool_usage = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PoolId", Flag: "pool-id", Type: "*string", Required: true},
}

var fields_get_console_output = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "Latest", Flag: "latest", Type: "*bool", Required: false},
}

var fields_get_console_screenshot = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "WakeUp", Flag: "wake-up", Type: "*bool", Required: false},
}

var fields_get_declarative_policies_report_summary = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ReportId", Flag: "report-id", Type: "*string", Required: true},
}

var fields_get_default_credit_specification = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceFamily", Flag: "instance-family", Type: "types.UnlimitedSupportedInstanceFamily", Required: true},
}

var fields_get_ebs_default_kms_key_id = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_get_ebs_encryption_by_default = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_get_enabled_ipam_policy = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_get_flow_logs_integration_template = []leanruntime.Field{
	{Name: "ConfigDeliveryS3DestinationArn", Flag: "config-delivery-s3-destination-arn", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "FlowLogId", Flag: "flow-log-id", Type: "*string", Required: true},
	{Name: "IntegrateServices", Flag: "integrate-services", Type: "*types.IntegrateServices", Required: true},
}

var fields_get_groups_for_capacity_reservation = []leanruntime.Field{
	{Name: "CapacityReservationId", Flag: "capacity-reservation-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_host_reservation_purchase_preview = []leanruntime.Field{
	{Name: "HostIdSet", Flag: "host-id-set", Type: "[]string", Required: true},
	{Name: "OfferingId", Flag: "offering-id", Type: "*string", Required: true},
}

var fields_get_image_ancestry = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: true},
}

var fields_get_image_block_public_access_state = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_get_instance_metadata_defaults = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_get_instance_tpm_ek_pub = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "KeyFormat", Flag: "key-format", Type: "types.EkPubKeyFormat", Required: true},
	{Name: "KeyType", Flag: "key-type", Type: "types.EkPubKeyType", Required: true},
}

var fields_get_instance_types_from_instance_requirements = []leanruntime.Field{
	{Name: "ArchitectureTypes", Flag: "architecture-types", Type: "[]types.ArchitectureType", Required: true},
	{Name: "Context", Flag: "context", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceRequirements", Flag: "instance-requirements", Type: "*types.InstanceRequirementsRequest", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VirtualizationTypes", Flag: "virtualization-types", Type: "[]types.VirtualizationType", Required: true},
}

var fields_get_instance_uefi_data = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_get_ipam_address_history = []leanruntime.Field{
	{Name: "Cidr", Flag: "cidr", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "IpamScopeId", Flag: "ipam-scope-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: false},
}

var fields_get_ipam_discovered_accounts = []leanruntime.Field{
	{Name: "DiscoveryRegion", Flag: "discovery-region", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IpamResourceDiscoveryId", Flag: "ipam-resource-discovery-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_ipam_discovered_public_addresses = []leanruntime.Field{
	{Name: "AddressRegion", Flag: "address-region", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IpamResourceDiscoveryId", Flag: "ipam-resource-discovery-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_ipam_discovered_resource_cidrs = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IpamResourceDiscoveryId", Flag: "ipam-resource-discovery-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceRegion", Flag: "resource-region", Type: "*string", Required: true},
}

var fields_get_ipam_policy_allocation_rules = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IpamPolicyId", Flag: "ipam-policy-id", Type: "*string", Required: true},
	{Name: "Locale", Flag: "locale", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.IpamPolicyResourceType", Required: false},
}

var fields_get_ipam_policy_organization_targets = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IpamPolicyId", Flag: "ipam-policy-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_ipam_pool_allocations = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IpamPoolAllocationId", Flag: "ipam-pool-allocation-id", Type: "*string", Required: false},
	{Name: "IpamPoolId", Flag: "ipam-pool-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_ipam_pool_cidrs = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IpamPoolId", Flag: "ipam-pool-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_ipam_prefix_list_resolver_rules = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IpamPrefixListResolverId", Flag: "ipam-prefix-list-resolver-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_ipam_prefix_list_resolver_version_entries = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamPrefixListResolverId", Flag: "ipam-prefix-list-resolver-id", Type: "*string", Required: true},
	{Name: "IpamPrefixListResolverVersion", Flag: "ipam-prefix-list-resolver-version", Type: "*int64", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_ipam_prefix_list_resolver_versions = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IpamPrefixListResolverId", Flag: "ipam-prefix-list-resolver-id", Type: "*string", Required: true},
	{Name: "IpamPrefixListResolverVersions", Flag: "ipam-prefix-list-resolver-versions", Type: "[]int64", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_ipam_resource_cidrs = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "IpamPoolId", Flag: "ipam-pool-id", Type: "*string", Required: false},
	{Name: "IpamScopeId", Flag: "ipam-scope-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: false},
	{Name: "ResourceOwner", Flag: "resource-owner", Type: "*string", Required: false},
	{Name: "ResourceTag", Flag: "resource-tag", Type: "*types.RequestIpamResourceTag", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.IpamResourceType", Required: false},
}

var fields_get_launch_template_data = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_get_managed_prefix_list_associations = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PrefixListId", Flag: "prefix-list-id", Type: "*string", Required: true},
}

var fields_get_managed_prefix_list_entries = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PrefixListId", Flag: "prefix-list-id", Type: "*string", Required: true},
	{Name: "TargetVersion", Flag: "target-version", Type: "*int64", Required: false},
}

var fields_get_network_insights_access_scope_analysis_findings = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NetworkInsightsAccessScopeAnalysisId", Flag: "network-insights-access-scope-analysis-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_network_insights_access_scope_content = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "NetworkInsightsAccessScopeId", Flag: "network-insights-access-scope-id", Type: "*string", Required: true},
}

var fields_get_password_data = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_get_reserved_instances_exchange_quote = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ReservedInstanceIds", Flag: "reserved-instance-ids", Type: "[]string", Required: true},
	{Name: "TargetConfigurations", Flag: "target-configurations", Type: "[]types.TargetConfigurationRequest", Required: false},
}

var fields_get_route_server_associations = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "RouteServerId", Flag: "route-server-id", Type: "*string", Required: true},
}

var fields_get_route_server_propagations = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "RouteServerId", Flag: "route-server-id", Type: "*string", Required: true},
	{Name: "RouteTableId", Flag: "route-table-id", Type: "*string", Required: false},
}

var fields_get_route_server_routing_database = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RouteServerId", Flag: "route-server-id", Type: "*string", Required: true},
}

var fields_get_security_groups_for_vpc = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_get_serial_console_access_status = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_get_snapshot_block_public_access_state = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_get_spot_placement_scores = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceRequirementsWithMetadata", Flag: "instance-requirements-with-metadata", Type: "*types.InstanceRequirementsWithMetadataRequest", Required: false},
	{Name: "InstanceTypes", Flag: "instance-types", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegionNames", Flag: "region-names", Type: "[]string", Required: false},
	{Name: "SingleAvailabilityZone", Flag: "single-availability-zone", Type: "*bool", Required: false},
	{Name: "TargetCapacity", Flag: "target-capacity", Type: "*int32", Required: true},
	{Name: "TargetCapacityUnitType", Flag: "target-capacity-unit-type", Type: "types.TargetCapacityUnitType", Required: false},
}

var fields_get_subnet_cidr_reservations = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SubnetId", Flag: "subnet-id", Type: "*string", Required: true},
}

var fields_get_transit_gateway_attachment_propagations = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TransitGatewayAttachmentId", Flag: "transit-gateway-attachment-id", Type: "*string", Required: true},
}

var fields_get_transit_gateway_metering_policy_entries = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TransitGatewayMeteringPolicyId", Flag: "transit-gateway-metering-policy-id", Type: "*string", Required: true},
}

var fields_get_transit_gateway_multicast_domain_associations = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TransitGatewayMulticastDomainId", Flag: "transit-gateway-multicast-domain-id", Type: "*string", Required: true},
}

var fields_get_transit_gateway_policy_table_associations = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TransitGatewayPolicyTableId", Flag: "transit-gateway-policy-table-id", Type: "*string", Required: true},
}

var fields_get_transit_gateway_policy_table_entries = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TransitGatewayPolicyTableId", Flag: "transit-gateway-policy-table-id", Type: "*string", Required: true},
}

var fields_get_transit_gateway_prefix_list_references = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TransitGatewayRouteTableId", Flag: "transit-gateway-route-table-id", Type: "*string", Required: true},
}

var fields_get_transit_gateway_route_table_associations = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TransitGatewayRouteTableId", Flag: "transit-gateway-route-table-id", Type: "*string", Required: true},
}

var fields_get_transit_gateway_route_table_propagations = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TransitGatewayRouteTableId", Flag: "transit-gateway-route-table-id", Type: "*string", Required: true},
}

var fields_get_verified_access_endpoint_policy = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VerifiedAccessEndpointId", Flag: "verified-access-endpoint-id", Type: "*string", Required: true},
}

var fields_get_verified_access_endpoint_targets = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VerifiedAccessEndpointId", Flag: "verified-access-endpoint-id", Type: "*string", Required: true},
}

var fields_get_verified_access_group_policy = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VerifiedAccessGroupId", Flag: "verified-access-group-id", Type: "*string", Required: true},
}

var fields_get_vpc_resources_blocking_encryption_enforcement = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_get_vpn_connection_device_sample_configuration = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InternetKeyExchangeVersion", Flag: "internet-key-exchange-version", Type: "*string", Required: false},
	{Name: "SampleType", Flag: "sample-type", Type: "*string", Required: false},
	{Name: "VpnConnectionDeviceTypeId", Flag: "vpn-connection-device-type-id", Type: "*string", Required: true},
	{Name: "VpnConnectionId", Flag: "vpn-connection-id", Type: "*string", Required: true},
}

var fields_get_vpn_connection_device_types = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_vpn_tunnel_replacement_status = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VpnConnectionId", Flag: "vpn-connection-id", Type: "*string", Required: true},
	{Name: "VpnTunnelOutsideIpAddress", Flag: "vpn-tunnel-outside-ip-address", Type: "*string", Required: true},
}

var fields_import_client_vpn_client_certificate_revocation_list = []leanruntime.Field{
	{Name: "CertificateRevocationList", Flag: "certificate-revocation-list", Type: "*string", Required: true},
	{Name: "ClientVpnEndpointId", Flag: "client-vpn-endpoint-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_import_image = []leanruntime.Field{
	{Name: "Architecture", Flag: "architecture", Type: "*string", Required: false},
	{Name: "BootMode", Flag: "boot-mode", Type: "types.BootModeValues", Required: false},
	{Name: "ClientData", Flag: "client-data", Type: "*types.ClientData", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DiskContainers", Flag: "disk-containers", Type: "[]types.ImageDiskContainer", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Encrypted", Flag: "encrypted", Type: "*bool", Required: false},
	{Name: "Hypervisor", Flag: "hypervisor", Type: "*string", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "LicenseSpecifications", Flag: "license-specifications", Type: "[]types.ImportImageLicenseConfigurationRequest", Required: false},
	{Name: "LicenseType", Flag: "license-type", Type: "*string", Required: false},
	{Name: "Platform", Flag: "platform", Type: "*string", Required: false},
	{Name: "RoleName", Flag: "role-name", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "UsageOperation", Flag: "usage-operation", Type: "*string", Required: false},
}

var fields_import_instance = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DiskImages", Flag: "disk-images", Type: "[]types.DiskImage", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "LaunchSpecification", Flag: "launch-specification", Type: "*types.ImportInstanceLaunchSpecification", Required: false},
	{Name: "Platform", Flag: "platform", Type: "types.PlatformValues", Required: true},
}

var fields_import_key_pair = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "KeyName", Flag: "key-name", Type: "*string", Required: true},
	{Name: "PublicKeyMaterial", Flag: "public-key-material", Type: "[]byte", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_import_snapshot = []leanruntime.Field{
	{Name: "ClientData", Flag: "client-data", Type: "*types.ClientData", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DiskContainer", Flag: "disk-container", Type: "*types.SnapshotDiskContainer", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Encrypted", Flag: "encrypted", Type: "*bool", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "RoleName", Flag: "role-name", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_import_volume = []leanruntime.Field{
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "AvailabilityZoneId", Flag: "availability-zone-id", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Image", Flag: "image", Type: "*types.DiskImageDetail", Required: true},
	{Name: "Volume", Flag: "volume", Type: "*types.VolumeDetail", Required: true},
}

var fields_list_images_in_recycle_bin = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ImageIds", Flag: "image-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_snapshots_in_recycle_bin = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SnapshotIds", Flag: "snapshot-ids", Type: "[]string", Required: false},
}

var fields_list_volumes_in_recycle_bin = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VolumeIds", Flag: "volume-ids", Type: "[]string", Required: false},
}

var fields_lock_snapshot = []leanruntime.Field{
	{Name: "CoolOffPeriod", Flag: "cool-off-period", Type: "*int32", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ExpirationDate", Flag: "expiration-date", Type: "*time.Time", Required: false},
	{Name: "LockDuration", Flag: "lock-duration", Type: "*int32", Required: false},
	{Name: "LockMode", Flag: "lock-mode", Type: "types.LockMode", Required: true},
	{Name: "SnapshotId", Flag: "snapshot-id", Type: "*string", Required: true},
}

var fields_modify_address_attribute = []leanruntime.Field{
	{Name: "AllocationId", Flag: "allocation-id", Type: "*string", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_modify_availability_zone_group = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: true},
	{Name: "OptInStatus", Flag: "opt-in-status", Type: "types.ModifyAvailabilityZoneOptInStatus", Required: true},
}

var fields_modify_capacity_reservation = []leanruntime.Field{
	{Name: "Accept", Flag: "accept", Type: "*bool", Required: false},
	{Name: "AdditionalInfo", Flag: "additional-info", Type: "*string", Required: false},
	{Name: "CapacityReservationId", Flag: "capacity-reservation-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EndDate", Flag: "end-date", Type: "*time.Time", Required: false},
	{Name: "EndDateType", Flag: "end-date-type", Type: "types.EndDateType", Required: false},
	{Name: "InstanceCount", Flag: "instance-count", Type: "*int32", Required: false},
	{Name: "InstanceMatchCriteria", Flag: "instance-match-criteria", Type: "types.InstanceMatchCriteria", Required: false},
}

var fields_modify_capacity_reservation_fleet = []leanruntime.Field{
	{Name: "CapacityReservationFleetId", Flag: "capacity-reservation-fleet-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EndDate", Flag: "end-date", Type: "*time.Time", Required: false},
	{Name: "RemoveEndDate", Flag: "remove-end-date", Type: "*bool", Required: false},
	{Name: "TotalTargetCapacity", Flag: "total-target-capacity", Type: "*int32", Required: false},
}

var fields_modify_client_vpn_endpoint = []leanruntime.Field{
	{Name: "ClientConnectOptions", Flag: "client-connect-options", Type: "*types.ClientConnectOptions", Required: false},
	{Name: "ClientLoginBannerOptions", Flag: "client-login-banner-options", Type: "*types.ClientLoginBannerOptions", Required: false},
	{Name: "ClientRouteEnforcementOptions", Flag: "client-route-enforcement-options", Type: "*types.ClientRouteEnforcementOptions", Required: false},
	{Name: "ClientVpnEndpointId", Flag: "client-vpn-endpoint-id", Type: "*string", Required: true},
	{Name: "ConnectionLogOptions", Flag: "connection-log-options", Type: "*types.ConnectionLogOptions", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisconnectOnSessionTimeout", Flag: "disconnect-on-session-timeout", Type: "*bool", Required: false},
	{Name: "DnsServers", Flag: "dns-servers", Type: "*types.DnsServersOptionsModifyStructure", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "SelfServicePortal", Flag: "self-service-portal", Type: "types.SelfServicePortal", Required: false},
	{Name: "ServerCertificateArn", Flag: "server-certificate-arn", Type: "*string", Required: false},
	{Name: "SessionTimeoutHours", Flag: "session-timeout-hours", Type: "*int32", Required: false},
	{Name: "SplitTunnel", Flag: "split-tunnel", Type: "*bool", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: false},
	{Name: "VpnPort", Flag: "vpn-port", Type: "*int32", Required: false},
}

var fields_modify_default_credit_specification = []leanruntime.Field{
	{Name: "CpuCredits", Flag: "cpu-credits", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceFamily", Flag: "instance-family", Type: "types.UnlimitedSupportedInstanceFamily", Required: true},
}

var fields_modify_ebs_default_kms_key_id = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: true},
}

var fields_modify_fleet = []leanruntime.Field{
	{Name: "Context", Flag: "context", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ExcessCapacityTerminationPolicy", Flag: "excess-capacity-termination-policy", Type: "types.FleetExcessCapacityTerminationPolicy", Required: false},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "LaunchTemplateConfigs", Flag: "launch-template-configs", Type: "[]types.FleetLaunchTemplateConfigRequest", Required: false},
	{Name: "TargetCapacitySpecification", Flag: "target-capacity-specification", Type: "*types.TargetCapacitySpecificationRequest", Required: false},
}

var fields_modify_fpga_image_attribute = []leanruntime.Field{
	{Name: "Attribute", Flag: "attribute", Type: "types.FpgaImageAttributeName", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "FpgaImageId", Flag: "fpga-image-id", Type: "*string", Required: true},
	{Name: "LoadPermission", Flag: "load-permission", Type: "*types.LoadPermissionModifications", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "OperationType", Flag: "operation-type", Type: "types.OperationType", Required: false},
	{Name: "ProductCodes", Flag: "product-codes", Type: "[]string", Required: false},
	{Name: "UserGroups", Flag: "user-groups", Type: "[]string", Required: false},
	{Name: "UserIds", Flag: "user-ids", Type: "[]string", Required: false},
}

var fields_modify_hosts = []leanruntime.Field{
	{Name: "AutoPlacement", Flag: "auto-placement", Type: "types.AutoPlacement", Required: false},
	{Name: "HostIds", Flag: "host-ids", Type: "[]string", Required: true},
	{Name: "HostMaintenance", Flag: "host-maintenance", Type: "types.HostMaintenance", Required: false},
	{Name: "HostRecovery", Flag: "host-recovery", Type: "types.HostRecovery", Required: false},
	{Name: "InstanceFamily", Flag: "instance-family", Type: "*string", Required: false},
	{Name: "InstanceType", Flag: "instance-type", Type: "*string", Required: false},
}

var fields_modify_id_format = []leanruntime.Field{
	{Name: "Resource", Flag: "resource", Type: "*string", Required: true},
	{Name: "UseLongIds", Flag: "use-long-ids", Type: "*bool", Required: true},
}

var fields_modify_identity_id_format = []leanruntime.Field{
	{Name: "PrincipalArn", Flag: "principal-arn", Type: "*string", Required: true},
	{Name: "Resource", Flag: "resource", Type: "*string", Required: true},
	{Name: "UseLongIds", Flag: "use-long-ids", Type: "*bool", Required: true},
}

var fields_modify_image_attribute = []leanruntime.Field{
	{Name: "Attribute", Flag: "attribute", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*types.AttributeValue", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: true},
	{Name: "ImdsSupport", Flag: "imds-support", Type: "*types.AttributeValue", Required: false},
	{Name: "LaunchPermission", Flag: "launch-permission", Type: "*types.LaunchPermissionModifications", Required: false},
	{Name: "OperationType", Flag: "operation-type", Type: "types.OperationType", Required: false},
	{Name: "OrganizationArns", Flag: "organization-arns", Type: "[]string", Required: false},
	{Name: "OrganizationalUnitArns", Flag: "organizational-unit-arns", Type: "[]string", Required: false},
	{Name: "ProductCodes", Flag: "product-codes", Type: "[]string", Required: false},
	{Name: "UserGroups", Flag: "user-groups", Type: "[]string", Required: false},
	{Name: "UserIds", Flag: "user-ids", Type: "[]string", Required: false},
	{Name: "Value", Flag: "value", Type: "*string", Required: false},
}

var fields_modify_instance_attribute = []leanruntime.Field{
	{Name: "Attribute", Flag: "attribute", Type: "types.InstanceAttributeName", Required: false},
	{Name: "BlockDeviceMappings", Flag: "block-device-mappings", Type: "[]types.InstanceBlockDeviceMappingSpecification", Required: false},
	{Name: "DisableApiStop", Flag: "disable-api-stop", Type: "*types.AttributeBooleanValue", Required: false},
	{Name: "DisableApiTermination", Flag: "disable-api-termination", Type: "*types.AttributeBooleanValue", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EbsOptimized", Flag: "ebs-optimized", Type: "*types.AttributeBooleanValue", Required: false},
	{Name: "EnaSupport", Flag: "ena-support", Type: "*types.AttributeBooleanValue", Required: false},
	{Name: "Groups", Flag: "groups", Type: "[]string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "InstanceInitiatedShutdownBehavior", Flag: "instance-initiated-shutdown-behavior", Type: "*types.AttributeValue", Required: false},
	{Name: "InstanceType", Flag: "instance-type", Type: "*types.AttributeValue", Required: false},
	{Name: "Kernel", Flag: "kernel", Type: "*types.AttributeValue", Required: false},
	{Name: "Ramdisk", Flag: "ramdisk", Type: "*types.AttributeValue", Required: false},
	{Name: "SourceDestCheck", Flag: "source-dest-check", Type: "*types.AttributeBooleanValue", Required: false},
	{Name: "SriovNetSupport", Flag: "sriov-net-support", Type: "*types.AttributeValue", Required: false},
	{Name: "UserData", Flag: "user-data", Type: "*types.BlobAttributeValue", Required: false},
	{Name: "Value", Flag: "value", Type: "*string", Required: false},
}

var fields_modify_instance_capacity_reservation_attributes = []leanruntime.Field{
	{Name: "CapacityReservationSpecification", Flag: "capacity-reservation-specification", Type: "*types.CapacityReservationSpecification", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_modify_instance_connect_endpoint = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceConnectEndpointId", Flag: "instance-connect-endpoint-id", Type: "*string", Required: true},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.IpAddressType", Required: false},
	{Name: "PreserveClientIp", Flag: "preserve-client-ip", Type: "*bool", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
}

var fields_modify_instance_cpu_options = []leanruntime.Field{
	{Name: "CoreCount", Flag: "core-count", Type: "*int32", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "NestedVirtualization", Flag: "nested-virtualization", Type: "types.NestedVirtualizationSpecification", Required: false},
	{Name: "ThreadsPerCore", Flag: "threads-per-core", Type: "*int32", Required: false},
}

var fields_modify_instance_credit_specification = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceCreditSpecifications", Flag: "instance-credit-specifications", Type: "[]types.InstanceCreditSpecificationRequest", Required: true},
}

var fields_modify_instance_event_start_time = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceEventId", Flag: "instance-event-id", Type: "*string", Required: true},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "NotBefore", Flag: "not-before", Type: "*time.Time", Required: true},
}

var fields_modify_instance_event_window = []leanruntime.Field{
	{Name: "CronExpression", Flag: "cron-expression", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceEventWindowId", Flag: "instance-event-window-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "TimeRanges", Flag: "time-ranges", Type: "[]types.InstanceEventWindowTimeRangeRequest", Required: false},
}

var fields_modify_instance_maintenance_options = []leanruntime.Field{
	{Name: "AutoRecovery", Flag: "auto-recovery", Type: "types.InstanceAutoRecoveryState", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "RebootMigration", Flag: "reboot-migration", Type: "types.InstanceRebootMigrationState", Required: false},
}

var fields_modify_instance_metadata_defaults = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "HttpEndpoint", Flag: "http-endpoint", Type: "types.DefaultInstanceMetadataEndpointState", Required: false},
	{Name: "HttpPutResponseHopLimit", Flag: "http-put-response-hop-limit", Type: "*int32", Required: false},
	{Name: "HttpTokens", Flag: "http-tokens", Type: "types.MetadataDefaultHttpTokensState", Required: false},
	{Name: "HttpTokensEnforced", Flag: "http-tokens-enforced", Type: "types.DefaultHttpTokensEnforcedState", Required: false},
	{Name: "InstanceMetadataTags", Flag: "instance-metadata-tags", Type: "types.DefaultInstanceMetadataTagsState", Required: false},
}

var fields_modify_instance_metadata_options = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "HttpEndpoint", Flag: "http-endpoint", Type: "types.InstanceMetadataEndpointState", Required: false},
	{Name: "HttpProtocolIpv6", Flag: "http-protocol-ipv6", Type: "types.InstanceMetadataProtocolState", Required: false},
	{Name: "HttpPutResponseHopLimit", Flag: "http-put-response-hop-limit", Type: "*int32", Required: false},
	{Name: "HttpTokens", Flag: "http-tokens", Type: "types.HttpTokensState", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "InstanceMetadataTags", Flag: "instance-metadata-tags", Type: "types.InstanceMetadataTagsState", Required: false},
}

var fields_modify_instance_network_performance_options = []leanruntime.Field{
	{Name: "BandwidthWeighting", Flag: "bandwidth-weighting", Type: "types.InstanceBandwidthWeighting", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_modify_instance_placement = []leanruntime.Field{
	{Name: "Affinity", Flag: "affinity", Type: "types.Affinity", Required: false},
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: false},
	{Name: "HostId", Flag: "host-id", Type: "*string", Required: false},
	{Name: "HostResourceGroupArn", Flag: "host-resource-group-arn", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "PartitionNumber", Flag: "partition-number", Type: "*int32", Required: false},
	{Name: "Tenancy", Flag: "tenancy", Type: "types.HostTenancy", Required: false},
}

var fields_modify_ipam = []leanruntime.Field{
	{Name: "AddOperatingRegions", Flag: "add-operating-regions", Type: "[]types.AddIpamOperatingRegion", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EnablePrivateGua", Flag: "enable-private-gua", Type: "*bool", Required: false},
	{Name: "IpamId", Flag: "ipam-id", Type: "*string", Required: true},
	{Name: "MeteredAccount", Flag: "metered-account", Type: "types.IpamMeteredAccount", Required: false},
	{Name: "RemoveOperatingRegions", Flag: "remove-operating-regions", Type: "[]types.RemoveIpamOperatingRegion", Required: false},
	{Name: "Tier", Flag: "tier", Type: "types.IpamTier", Required: false},
}

var fields_modify_ipam_policy_allocation_rules = []leanruntime.Field{
	{Name: "AllocationRules", Flag: "allocation-rules", Type: "[]types.IpamPolicyAllocationRuleRequest", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamPolicyId", Flag: "ipam-policy-id", Type: "*string", Required: true},
	{Name: "Locale", Flag: "locale", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.IpamPolicyResourceType", Required: true},
}

var fields_modify_ipam_pool = []leanruntime.Field{
	{Name: "AddAllocationResourceTags", Flag: "add-allocation-resource-tags", Type: "[]types.RequestIpamResourceTag", Required: false},
	{Name: "AllocationDefaultNetmaskLength", Flag: "allocation-default-netmask-length", Type: "*int32", Required: false},
	{Name: "AllocationMaxNetmaskLength", Flag: "allocation-max-netmask-length", Type: "*int32", Required: false},
	{Name: "AllocationMinNetmaskLength", Flag: "allocation-min-netmask-length", Type: "*int32", Required: false},
	{Name: "AutoImport", Flag: "auto-import", Type: "*bool", Required: false},
	{Name: "ClearAllocationDefaultNetmaskLength", Flag: "clear-allocation-default-netmask-length", Type: "*bool", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamPoolId", Flag: "ipam-pool-id", Type: "*string", Required: true},
	{Name: "RemoveAllocationResourceTags", Flag: "remove-allocation-resource-tags", Type: "[]types.RequestIpamResourceTag", Required: false},
}

var fields_modify_ipam_prefix_list_resolver = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamPrefixListResolverId", Flag: "ipam-prefix-list-resolver-id", Type: "*string", Required: true},
	{Name: "Rules", Flag: "rules", Type: "[]types.IpamPrefixListResolverRuleRequest", Required: false},
}

var fields_modify_ipam_prefix_list_resolver_target = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DesiredVersion", Flag: "desired-version", Type: "*int64", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamPrefixListResolverTargetId", Flag: "ipam-prefix-list-resolver-target-id", Type: "*string", Required: true},
	{Name: "TrackLatestVersion", Flag: "track-latest-version", Type: "*bool", Required: false},
}

var fields_modify_ipam_resource_cidr = []leanruntime.Field{
	{Name: "CurrentIpamScopeId", Flag: "current-ipam-scope-id", Type: "*string", Required: true},
	{Name: "DestinationIpamScopeId", Flag: "destination-ipam-scope-id", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Monitored", Flag: "monitored", Type: "*bool", Required: true},
	{Name: "ResourceCidr", Flag: "resource-cidr", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "ResourceRegion", Flag: "resource-region", Type: "*string", Required: true},
}

var fields_modify_ipam_resource_discovery = []leanruntime.Field{
	{Name: "AddOperatingRegions", Flag: "add-operating-regions", Type: "[]types.AddIpamOperatingRegion", Required: false},
	{Name: "AddOrganizationalUnitExclusions", Flag: "add-organizational-unit-exclusions", Type: "[]types.AddIpamOrganizationalUnitExclusion", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamResourceDiscoveryId", Flag: "ipam-resource-discovery-id", Type: "*string", Required: true},
	{Name: "RemoveOperatingRegions", Flag: "remove-operating-regions", Type: "[]types.RemoveIpamOperatingRegion", Required: false},
	{Name: "RemoveOrganizationalUnitExclusions", Flag: "remove-organizational-unit-exclusions", Type: "[]types.RemoveIpamOrganizationalUnitExclusion", Required: false},
}

var fields_modify_ipam_scope = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ExternalAuthorityConfiguration", Flag: "external-authority-configuration", Type: "*types.ExternalAuthorityConfiguration", Required: false},
	{Name: "IpamScopeId", Flag: "ipam-scope-id", Type: "*string", Required: true},
	{Name: "RemoveExternalAuthorityConfiguration", Flag: "remove-external-authority-configuration", Type: "*bool", Required: false},
}

var fields_modify_launch_template = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DefaultVersion", Flag: "default-version", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "LaunchTemplateId", Flag: "launch-template-id", Type: "*string", Required: false},
	{Name: "LaunchTemplateName", Flag: "launch-template-name", Type: "*string", Required: false},
}

var fields_modify_local_gateway_route = []leanruntime.Field{
	{Name: "DestinationCidrBlock", Flag: "destination-cidr-block", Type: "*string", Required: false},
	{Name: "DestinationPrefixListId", Flag: "destination-prefix-list-id", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "LocalGatewayRouteTableId", Flag: "local-gateway-route-table-id", Type: "*string", Required: true},
	{Name: "LocalGatewayVirtualInterfaceGroupId", Flag: "local-gateway-virtual-interface-group-id", Type: "*string", Required: false},
	{Name: "NetworkInterfaceId", Flag: "network-interface-id", Type: "*string", Required: false},
}

var fields_modify_managed_prefix_list = []leanruntime.Field{
	{Name: "AddEntries", Flag: "add-entries", Type: "[]types.AddPrefixListEntry", Required: false},
	{Name: "CurrentVersion", Flag: "current-version", Type: "*int64", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamPrefixListResolverSyncEnabled", Flag: "ipam-prefix-list-resolver-sync-enabled", Type: "*bool", Required: false},
	{Name: "MaxEntries", Flag: "max-entries", Type: "*int32", Required: false},
	{Name: "PrefixListId", Flag: "prefix-list-id", Type: "*string", Required: true},
	{Name: "PrefixListName", Flag: "prefix-list-name", Type: "*string", Required: false},
	{Name: "RemoveEntries", Flag: "remove-entries", Type: "[]types.RemovePrefixListEntry", Required: false},
}

var fields_modify_network_interface_attribute = []leanruntime.Field{
	{Name: "AssociatePublicIpAddress", Flag: "associate-public-ip-address", Type: "*bool", Required: false},
	{Name: "AssociatedSubnetIds", Flag: "associated-subnet-ids", Type: "[]string", Required: false},
	{Name: "Attachment", Flag: "attachment", Type: "*types.NetworkInterfaceAttachmentChanges", Required: false},
	{Name: "ConnectionTrackingSpecification", Flag: "connection-tracking-specification", Type: "*types.ConnectionTrackingSpecificationRequest", Required: false},
	{Name: "Description", Flag: "description", Type: "*types.AttributeValue", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EnaSrdSpecification", Flag: "ena-srd-specification", Type: "*types.EnaSrdSpecification", Required: false},
	{Name: "EnablePrimaryIpv6", Flag: "enable-primary-ipv6", Type: "*bool", Required: false},
	{Name: "Groups", Flag: "groups", Type: "[]string", Required: false},
	{Name: "NetworkInterfaceId", Flag: "network-interface-id", Type: "*string", Required: true},
	{Name: "SourceDestCheck", Flag: "source-dest-check", Type: "*types.AttributeBooleanValue", Required: false},
}

var fields_modify_private_dns_name_options = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EnableResourceNameDnsAAAARecord", Flag: "enable-resource-name-dns-aaaa-record", Type: "*bool", Required: false},
	{Name: "EnableResourceNameDnsARecord", Flag: "enable-resource-name-dns-arecord", Type: "*bool", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
	{Name: "PrivateDnsHostnameType", Flag: "private-dns-hostname-type", Type: "types.HostnameType", Required: false},
}

var fields_modify_public_ip_dns_name_options = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "HostnameType", Flag: "hostname-type", Type: "types.PublicIpDnsOption", Required: true},
	{Name: "NetworkInterfaceId", Flag: "network-interface-id", Type: "*string", Required: true},
}

var fields_modify_reserved_instances = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ReservedInstancesIds", Flag: "reserved-instances-ids", Type: "[]string", Required: true},
	{Name: "TargetConfigurations", Flag: "target-configurations", Type: "[]types.ReservedInstancesConfiguration", Required: true},
}

var fields_modify_route_server = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "PersistRoutes", Flag: "persist-routes", Type: "types.RouteServerPersistRoutesAction", Required: false},
	{Name: "PersistRoutesDuration", Flag: "persist-routes-duration", Type: "*int64", Required: false},
	{Name: "RouteServerId", Flag: "route-server-id", Type: "*string", Required: true},
	{Name: "SnsNotificationsEnabled", Flag: "sns-notifications-enabled", Type: "*bool", Required: false},
}

var fields_modify_security_group_rules = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "SecurityGroupRules", Flag: "security-group-rules", Type: "[]types.SecurityGroupRuleUpdate", Required: true},
}

var fields_modify_snapshot_attribute = []leanruntime.Field{
	{Name: "Attribute", Flag: "attribute", Type: "types.SnapshotAttributeName", Required: false},
	{Name: "CreateVolumePermission", Flag: "create-volume-permission", Type: "*types.CreateVolumePermissionModifications", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GroupNames", Flag: "group-names", Type: "[]string", Required: false},
	{Name: "OperationType", Flag: "operation-type", Type: "types.OperationType", Required: false},
	{Name: "SnapshotId", Flag: "snapshot-id", Type: "*string", Required: true},
	{Name: "UserIds", Flag: "user-ids", Type: "[]string", Required: false},
}

var fields_modify_snapshot_tier = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "SnapshotId", Flag: "snapshot-id", Type: "*string", Required: true},
	{Name: "StorageTier", Flag: "storage-tier", Type: "types.TargetStorageTier", Required: false},
}

var fields_modify_spot_fleet_request = []leanruntime.Field{
	{Name: "Context", Flag: "context", Type: "*string", Required: false},
	{Name: "ExcessCapacityTerminationPolicy", Flag: "excess-capacity-termination-policy", Type: "types.ExcessCapacityTerminationPolicy", Required: false},
	{Name: "LaunchTemplateConfigs", Flag: "launch-template-configs", Type: "[]types.LaunchTemplateConfig", Required: false},
	{Name: "OnDemandTargetCapacity", Flag: "on-demand-target-capacity", Type: "*int32", Required: false},
	{Name: "SpotFleetRequestId", Flag: "spot-fleet-request-id", Type: "*string", Required: true},
	{Name: "TargetCapacity", Flag: "target-capacity", Type: "*int32", Required: false},
}

var fields_modify_subnet_attribute = []leanruntime.Field{
	{Name: "AssignIpv6AddressOnCreation", Flag: "assign-ipv6-address-on-creation", Type: "*types.AttributeBooleanValue", Required: false},
	{Name: "CustomerOwnedIpv4Pool", Flag: "customer-owned-ipv4-pool", Type: "*string", Required: false},
	{Name: "DisableLniAtDeviceIndex", Flag: "disable-lni-at-device-index", Type: "*types.AttributeBooleanValue", Required: false},
	{Name: "EnableDns64", Flag: "enable-dns64", Type: "*types.AttributeBooleanValue", Required: false},
	{Name: "EnableLniAtDeviceIndex", Flag: "enable-lni-at-device-index", Type: "*int32", Required: false},
	{Name: "EnableResourceNameDnsAAAARecordOnLaunch", Flag: "enable-resource-name-dns-aaaa-record-on-launch", Type: "*types.AttributeBooleanValue", Required: false},
	{Name: "EnableResourceNameDnsARecordOnLaunch", Flag: "enable-resource-name-dns-arecord-on-launch", Type: "*types.AttributeBooleanValue", Required: false},
	{Name: "MapCustomerOwnedIpOnLaunch", Flag: "map-customer-owned-ip-on-launch", Type: "*types.AttributeBooleanValue", Required: false},
	{Name: "MapPublicIpOnLaunch", Flag: "map-public-ip-on-launch", Type: "*types.AttributeBooleanValue", Required: false},
	{Name: "PrivateDnsHostnameTypeOnLaunch", Flag: "private-dns-hostname-type-on-launch", Type: "types.HostnameType", Required: false},
	{Name: "SubnetId", Flag: "subnet-id", Type: "*string", Required: true},
}

var fields_modify_traffic_mirror_filter_network_services = []leanruntime.Field{
	{Name: "AddNetworkServices", Flag: "add-network-services", Type: "[]types.TrafficMirrorNetworkService", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "RemoveNetworkServices", Flag: "remove-network-services", Type: "[]types.TrafficMirrorNetworkService", Required: false},
	{Name: "TrafficMirrorFilterId", Flag: "traffic-mirror-filter-id", Type: "*string", Required: true},
}

var fields_modify_traffic_mirror_filter_rule = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DestinationCidrBlock", Flag: "destination-cidr-block", Type: "*string", Required: false},
	{Name: "DestinationPortRange", Flag: "destination-port-range", Type: "*types.TrafficMirrorPortRangeRequest", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Protocol", Flag: "protocol", Type: "*int32", Required: false},
	{Name: "RemoveFields", Flag: "remove-fields", Type: "[]types.TrafficMirrorFilterRuleField", Required: false},
	{Name: "RuleAction", Flag: "rule-action", Type: "types.TrafficMirrorRuleAction", Required: false},
	{Name: "RuleNumber", Flag: "rule-number", Type: "*int32", Required: false},
	{Name: "SourceCidrBlock", Flag: "source-cidr-block", Type: "*string", Required: false},
	{Name: "SourcePortRange", Flag: "source-port-range", Type: "*types.TrafficMirrorPortRangeRequest", Required: false},
	{Name: "TrafficDirection", Flag: "traffic-direction", Type: "types.TrafficDirection", Required: false},
	{Name: "TrafficMirrorFilterRuleId", Flag: "traffic-mirror-filter-rule-id", Type: "*string", Required: true},
}

var fields_modify_traffic_mirror_session = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "PacketLength", Flag: "packet-length", Type: "*int32", Required: false},
	{Name: "RemoveFields", Flag: "remove-fields", Type: "[]types.TrafficMirrorSessionField", Required: false},
	{Name: "SessionNumber", Flag: "session-number", Type: "*int32", Required: false},
	{Name: "TrafficMirrorFilterId", Flag: "traffic-mirror-filter-id", Type: "*string", Required: false},
	{Name: "TrafficMirrorSessionId", Flag: "traffic-mirror-session-id", Type: "*string", Required: true},
	{Name: "TrafficMirrorTargetId", Flag: "traffic-mirror-target-id", Type: "*string", Required: false},
	{Name: "VirtualNetworkId", Flag: "virtual-network-id", Type: "*int32", Required: false},
}

var fields_modify_transit_gateway = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Options", Flag: "options", Type: "*types.ModifyTransitGatewayOptions", Required: false},
	{Name: "TransitGatewayId", Flag: "transit-gateway-id", Type: "*string", Required: true},
}

var fields_modify_transit_gateway_metering_policy = []leanruntime.Field{
	{Name: "AddMiddleboxAttachmentIds", Flag: "add-middlebox-attachment-ids", Type: "[]string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "RemoveMiddleboxAttachmentIds", Flag: "remove-middlebox-attachment-ids", Type: "[]string", Required: false},
	{Name: "TransitGatewayMeteringPolicyId", Flag: "transit-gateway-metering-policy-id", Type: "*string", Required: true},
}

var fields_modify_transit_gateway_prefix_list_reference = []leanruntime.Field{
	{Name: "Blackhole", Flag: "blackhole", Type: "*bool", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "PrefixListId", Flag: "prefix-list-id", Type: "*string", Required: true},
	{Name: "TransitGatewayAttachmentId", Flag: "transit-gateway-attachment-id", Type: "*string", Required: false},
	{Name: "TransitGatewayRouteTableId", Flag: "transit-gateway-route-table-id", Type: "*string", Required: true},
}

var fields_modify_transit_gateway_vpc_attachment = []leanruntime.Field{
	{Name: "AddSubnetIds", Flag: "add-subnet-ids", Type: "[]string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Options", Flag: "options", Type: "*types.ModifyTransitGatewayVpcAttachmentRequestOptions", Required: false},
	{Name: "RemoveSubnetIds", Flag: "remove-subnet-ids", Type: "[]string", Required: false},
	{Name: "TransitGatewayAttachmentId", Flag: "transit-gateway-attachment-id", Type: "*string", Required: true},
}

var fields_modify_verified_access_endpoint = []leanruntime.Field{
	{Name: "CidrOptions", Flag: "cidr-options", Type: "*types.ModifyVerifiedAccessEndpointCidrOptions", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "LoadBalancerOptions", Flag: "load-balancer-options", Type: "*types.ModifyVerifiedAccessEndpointLoadBalancerOptions", Required: false},
	{Name: "NetworkInterfaceOptions", Flag: "network-interface-options", Type: "*types.ModifyVerifiedAccessEndpointEniOptions", Required: false},
	{Name: "RdsOptions", Flag: "rds-options", Type: "*types.ModifyVerifiedAccessEndpointRdsOptions", Required: false},
	{Name: "VerifiedAccessEndpointId", Flag: "verified-access-endpoint-id", Type: "*string", Required: true},
	{Name: "VerifiedAccessGroupId", Flag: "verified-access-group-id", Type: "*string", Required: false},
}

var fields_modify_verified_access_endpoint_policy = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: false},
	{Name: "PolicyEnabled", Flag: "policy-enabled", Type: "*bool", Required: false},
	{Name: "SseSpecification", Flag: "sse-specification", Type: "*types.VerifiedAccessSseSpecificationRequest", Required: false},
	{Name: "VerifiedAccessEndpointId", Flag: "verified-access-endpoint-id", Type: "*string", Required: true},
}

var fields_modify_verified_access_group = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VerifiedAccessGroupId", Flag: "verified-access-group-id", Type: "*string", Required: true},
	{Name: "VerifiedAccessInstanceId", Flag: "verified-access-instance-id", Type: "*string", Required: false},
}

var fields_modify_verified_access_group_policy = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: false},
	{Name: "PolicyEnabled", Flag: "policy-enabled", Type: "*bool", Required: false},
	{Name: "SseSpecification", Flag: "sse-specification", Type: "*types.VerifiedAccessSseSpecificationRequest", Required: false},
	{Name: "VerifiedAccessGroupId", Flag: "verified-access-group-id", Type: "*string", Required: true},
}

var fields_modify_verified_access_instance = []leanruntime.Field{
	{Name: "CidrEndpointsCustomSubDomain", Flag: "cidr-endpoints-custom-sub-domain", Type: "*string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VerifiedAccessInstanceId", Flag: "verified-access-instance-id", Type: "*string", Required: true},
}

var fields_modify_verified_access_instance_logging_configuration = []leanruntime.Field{
	{Name: "AccessLogs", Flag: "access-logs", Type: "*types.VerifiedAccessLogOptions", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VerifiedAccessInstanceId", Flag: "verified-access-instance-id", Type: "*string", Required: true},
}

var fields_modify_verified_access_trust_provider = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DeviceOptions", Flag: "device-options", Type: "*types.ModifyVerifiedAccessTrustProviderDeviceOptions", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "NativeApplicationOidcOptions", Flag: "native-application-oidc-options", Type: "*types.ModifyVerifiedAccessNativeApplicationOidcOptions", Required: false},
	{Name: "OidcOptions", Flag: "oidc-options", Type: "*types.ModifyVerifiedAccessTrustProviderOidcOptions", Required: false},
	{Name: "SseSpecification", Flag: "sse-specification", Type: "*types.VerifiedAccessSseSpecificationRequest", Required: false},
	{Name: "VerifiedAccessTrustProviderId", Flag: "verified-access-trust-provider-id", Type: "*string", Required: true},
}

var fields_modify_volume = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Iops", Flag: "iops", Type: "*int32", Required: false},
	{Name: "MultiAttachEnabled", Flag: "multi-attach-enabled", Type: "*bool", Required: false},
	{Name: "Size", Flag: "size", Type: "*int32", Required: false},
	{Name: "Throughput", Flag: "throughput", Type: "*int32", Required: false},
	{Name: "VolumeId", Flag: "volume-id", Type: "*string", Required: true},
	{Name: "VolumeType", Flag: "volume-type", Type: "types.VolumeType", Required: false},
}

var fields_modify_volume_attribute = []leanruntime.Field{
	{Name: "AutoEnableIO", Flag: "auto-enable-io", Type: "*types.AttributeBooleanValue", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VolumeId", Flag: "volume-id", Type: "*string", Required: true},
}

var fields_modify_vpc_attribute = []leanruntime.Field{
	{Name: "EnableDnsHostnames", Flag: "enable-dns-hostnames", Type: "*types.AttributeBooleanValue", Required: false},
	{Name: "EnableDnsSupport", Flag: "enable-dns-support", Type: "*types.AttributeBooleanValue", Required: false},
	{Name: "EnableNetworkAddressUsageMetrics", Flag: "enable-network-address-usage-metrics", Type: "*types.AttributeBooleanValue", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_modify_vpc_block_public_access_exclusion = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ExclusionId", Flag: "exclusion-id", Type: "*string", Required: true},
	{Name: "InternetGatewayExclusionMode", Flag: "internet-gateway-exclusion-mode", Type: "types.InternetGatewayExclusionMode", Required: true},
}

var fields_modify_vpc_block_public_access_options = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InternetGatewayBlockMode", Flag: "internet-gateway-block-mode", Type: "types.InternetGatewayBlockMode", Required: true},
}

var fields_modify_vpc_encryption_control = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EgressOnlyInternetGatewayExclusion", Flag: "egress-only-internet-gateway-exclusion", Type: "types.VpcEncryptionControlExclusionStateInput", Required: false},
	{Name: "ElasticFileSystemExclusion", Flag: "elastic-file-system-exclusion", Type: "types.VpcEncryptionControlExclusionStateInput", Required: false},
	{Name: "InternetGatewayExclusion", Flag: "internet-gateway-exclusion", Type: "types.VpcEncryptionControlExclusionStateInput", Required: false},
	{Name: "LambdaExclusion", Flag: "lambda-exclusion", Type: "types.VpcEncryptionControlExclusionStateInput", Required: false},
	{Name: "Mode", Flag: "mode", Type: "types.VpcEncryptionControlMode", Required: false},
	{Name: "NatGatewayExclusion", Flag: "nat-gateway-exclusion", Type: "types.VpcEncryptionControlExclusionStateInput", Required: false},
	{Name: "VirtualPrivateGatewayExclusion", Flag: "virtual-private-gateway-exclusion", Type: "types.VpcEncryptionControlExclusionStateInput", Required: false},
	{Name: "VpcEncryptionControlId", Flag: "vpc-encryption-control-id", Type: "*string", Required: true},
	{Name: "VpcLatticeExclusion", Flag: "vpc-lattice-exclusion", Type: "types.VpcEncryptionControlExclusionStateInput", Required: false},
	{Name: "VpcPeeringExclusion", Flag: "vpc-peering-exclusion", Type: "types.VpcEncryptionControlExclusionStateInput", Required: false},
}

var fields_modify_vpc_endpoint = []leanruntime.Field{
	{Name: "AddRouteTableIds", Flag: "add-route-table-ids", Type: "[]string", Required: false},
	{Name: "AddSecurityGroupIds", Flag: "add-security-group-ids", Type: "[]string", Required: false},
	{Name: "AddSubnetIds", Flag: "add-subnet-ids", Type: "[]string", Required: false},
	{Name: "DnsOptions", Flag: "dns-options", Type: "*types.DnsOptionsSpecification", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.IpAddressType", Required: false},
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: false},
	{Name: "PrivateDnsEnabled", Flag: "private-dns-enabled", Type: "*bool", Required: false},
	{Name: "RemoveRouteTableIds", Flag: "remove-route-table-ids", Type: "[]string", Required: false},
	{Name: "RemoveSecurityGroupIds", Flag: "remove-security-group-ids", Type: "[]string", Required: false},
	{Name: "RemoveSubnetIds", Flag: "remove-subnet-ids", Type: "[]string", Required: false},
	{Name: "ResetPolicy", Flag: "reset-policy", Type: "*bool", Required: false},
	{Name: "SubnetConfigurations", Flag: "subnet-configurations", Type: "[]types.SubnetConfiguration", Required: false},
	{Name: "VpcEndpointId", Flag: "vpc-endpoint-id", Type: "*string", Required: true},
}

var fields_modify_vpc_endpoint_connection_notification = []leanruntime.Field{
	{Name: "ConnectionEvents", Flag: "connection-events", Type: "[]string", Required: false},
	{Name: "ConnectionNotificationArn", Flag: "connection-notification-arn", Type: "*string", Required: false},
	{Name: "ConnectionNotificationId", Flag: "connection-notification-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_modify_vpc_endpoint_service_configuration = []leanruntime.Field{
	{Name: "AcceptanceRequired", Flag: "acceptance-required", Type: "*bool", Required: false},
	{Name: "AddGatewayLoadBalancerArns", Flag: "add-gateway-load-balancer-arns", Type: "[]string", Required: false},
	{Name: "AddNetworkLoadBalancerArns", Flag: "add-network-load-balancer-arns", Type: "[]string", Required: false},
	{Name: "AddSupportedIpAddressTypes", Flag: "add-supported-ip-address-types", Type: "[]string", Required: false},
	{Name: "AddSupportedRegions", Flag: "add-supported-regions", Type: "[]string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "PrivateDnsName", Flag: "private-dns-name", Type: "*string", Required: false},
	{Name: "RemoveGatewayLoadBalancerArns", Flag: "remove-gateway-load-balancer-arns", Type: "[]string", Required: false},
	{Name: "RemoveNetworkLoadBalancerArns", Flag: "remove-network-load-balancer-arns", Type: "[]string", Required: false},
	{Name: "RemovePrivateDnsName", Flag: "remove-private-dns-name", Type: "*bool", Required: false},
	{Name: "RemoveSupportedIpAddressTypes", Flag: "remove-supported-ip-address-types", Type: "[]string", Required: false},
	{Name: "RemoveSupportedRegions", Flag: "remove-supported-regions", Type: "[]string", Required: false},
	{Name: "ServiceId", Flag: "service-id", Type: "*string", Required: true},
}

var fields_modify_vpc_endpoint_service_payer_responsibility = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "PayerResponsibility", Flag: "payer-responsibility", Type: "types.PayerResponsibility", Required: true},
	{Name: "ServiceId", Flag: "service-id", Type: "*string", Required: true},
}

var fields_modify_vpc_endpoint_service_permissions = []leanruntime.Field{
	{Name: "AddAllowedPrincipals", Flag: "add-allowed-principals", Type: "[]string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "RemoveAllowedPrincipals", Flag: "remove-allowed-principals", Type: "[]string", Required: false},
	{Name: "ServiceId", Flag: "service-id", Type: "*string", Required: true},
}

var fields_modify_vpc_peering_connection_options = []leanruntime.Field{
	{Name: "AccepterPeeringConnectionOptions", Flag: "accepter-peering-connection-options", Type: "*types.PeeringConnectionOptionsRequest", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "RequesterPeeringConnectionOptions", Flag: "requester-peering-connection-options", Type: "*types.PeeringConnectionOptionsRequest", Required: false},
	{Name: "VpcPeeringConnectionId", Flag: "vpc-peering-connection-id", Type: "*string", Required: true},
}

var fields_modify_vpc_tenancy = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceTenancy", Flag: "instance-tenancy", Type: "types.VpcTenancy", Required: true},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_modify_vpn_connection = []leanruntime.Field{
	{Name: "CustomerGatewayId", Flag: "customer-gateway-id", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TransitGatewayId", Flag: "transit-gateway-id", Type: "*string", Required: false},
	{Name: "VpnConnectionId", Flag: "vpn-connection-id", Type: "*string", Required: true},
	{Name: "VpnGatewayId", Flag: "vpn-gateway-id", Type: "*string", Required: false},
}

var fields_modify_vpn_connection_options = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "LocalIpv4NetworkCidr", Flag: "local-ipv4-network-cidr", Type: "*string", Required: false},
	{Name: "LocalIpv6NetworkCidr", Flag: "local-ipv6-network-cidr", Type: "*string", Required: false},
	{Name: "RemoteIpv4NetworkCidr", Flag: "remote-ipv4-network-cidr", Type: "*string", Required: false},
	{Name: "RemoteIpv6NetworkCidr", Flag: "remote-ipv6-network-cidr", Type: "*string", Required: false},
	{Name: "VpnConnectionId", Flag: "vpn-connection-id", Type: "*string", Required: true},
}

var fields_modify_vpn_tunnel_certificate = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VpnConnectionId", Flag: "vpn-connection-id", Type: "*string", Required: true},
	{Name: "VpnTunnelOutsideIpAddress", Flag: "vpn-tunnel-outside-ip-address", Type: "*string", Required: true},
}

var fields_modify_vpn_tunnel_options = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "PreSharedKeyStorage", Flag: "pre-shared-key-storage", Type: "*string", Required: false},
	{Name: "SkipTunnelReplacement", Flag: "skip-tunnel-replacement", Type: "*bool", Required: false},
	{Name: "TunnelOptions", Flag: "tunnel-options", Type: "*types.ModifyVpnTunnelOptionsSpecification", Required: true},
	{Name: "VpnConnectionId", Flag: "vpn-connection-id", Type: "*string", Required: true},
	{Name: "VpnTunnelOutsideIpAddress", Flag: "vpn-tunnel-outside-ip-address", Type: "*string", Required: true},
}

var fields_monitor_instances = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: true},
}

var fields_move_address_to_vpc = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "PublicIp", Flag: "public-ip", Type: "*string", Required: true},
}

var fields_move_byoip_cidr_to_ipam = []leanruntime.Field{
	{Name: "Cidr", Flag: "cidr", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamPoolId", Flag: "ipam-pool-id", Type: "*string", Required: true},
	{Name: "IpamPoolOwner", Flag: "ipam-pool-owner", Type: "*string", Required: true},
}

var fields_move_capacity_reservation_instances = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DestinationCapacityReservationId", Flag: "destination-capacity-reservation-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceCount", Flag: "instance-count", Type: "*int32", Required: true},
	{Name: "SourceCapacityReservationId", Flag: "source-capacity-reservation-id", Type: "*string", Required: true},
}

var fields_provision_byoip_cidr = []leanruntime.Field{
	{Name: "Cidr", Flag: "cidr", Type: "*string", Required: true},
	{Name: "CidrAuthorizationContext", Flag: "cidr-authorization-context", Type: "*types.CidrAuthorizationContext", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "MultiRegion", Flag: "multi-region", Type: "*bool", Required: false},
	{Name: "NetworkBorderGroup", Flag: "network-border-group", Type: "*string", Required: false},
	{Name: "PoolTagSpecifications", Flag: "pool-tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "PubliclyAdvertisable", Flag: "publicly-advertisable", Type: "*bool", Required: false},
}

var fields_provision_ipam_byoasn = []leanruntime.Field{
	{Name: "Asn", Flag: "asn", Type: "*string", Required: true},
	{Name: "AsnAuthorizationContext", Flag: "asn-authorization-context", Type: "*types.AsnAuthorizationContext", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamId", Flag: "ipam-id", Type: "*string", Required: true},
}

var fields_provision_ipam_pool_cidr = []leanruntime.Field{
	{Name: "Cidr", Flag: "cidr", Type: "*string", Required: false},
	{Name: "CidrAuthorizationContext", Flag: "cidr-authorization-context", Type: "*types.IpamCidrAuthorizationContext", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamExternalResourceVerificationTokenId", Flag: "ipam-external-resource-verification-token-id", Type: "*string", Required: false},
	{Name: "IpamPoolId", Flag: "ipam-pool-id", Type: "*string", Required: true},
	{Name: "NetmaskLength", Flag: "netmask-length", Type: "*int32", Required: false},
	{Name: "VerificationMethod", Flag: "verification-method", Type: "types.VerificationMethod", Required: false},
}

var fields_provision_public_ipv4_pool_cidr = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamPoolId", Flag: "ipam-pool-id", Type: "*string", Required: true},
	{Name: "NetmaskLength", Flag: "netmask-length", Type: "*int32", Required: true},
	{Name: "NetworkBorderGroup", Flag: "network-border-group", Type: "*string", Required: false},
	{Name: "PoolId", Flag: "pool-id", Type: "*string", Required: true},
}

var fields_purchase_capacity_block = []leanruntime.Field{
	{Name: "CapacityBlockOfferingId", Flag: "capacity-block-offering-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstancePlatform", Flag: "instance-platform", Type: "types.CapacityReservationInstancePlatform", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_purchase_capacity_block_extension = []leanruntime.Field{
	{Name: "CapacityBlockExtensionOfferingId", Flag: "capacity-block-extension-offering-id", Type: "*string", Required: true},
	{Name: "CapacityReservationId", Flag: "capacity-reservation-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_purchase_host_reservation = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CurrencyCode", Flag: "currency-code", Type: "types.CurrencyCodeValues", Required: false},
	{Name: "HostIdSet", Flag: "host-id-set", Type: "[]string", Required: true},
	{Name: "LimitPrice", Flag: "limit-price", Type: "*string", Required: false},
	{Name: "OfferingId", Flag: "offering-id", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_purchase_reserved_instances_offering = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceCount", Flag: "instance-count", Type: "*int32", Required: true},
	{Name: "LimitPrice", Flag: "limit-price", Type: "*types.ReservedInstanceLimitPrice", Required: false},
	{Name: "PurchaseTime", Flag: "purchase-time", Type: "*time.Time", Required: false},
	{Name: "ReservedInstancesOfferingId", Flag: "reserved-instances-offering-id", Type: "*string", Required: true},
}

var fields_purchase_scheduled_instances = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "PurchaseRequests", Flag: "purchase-requests", Type: "[]types.PurchaseRequest", Required: true},
}

var fields_reboot_instances = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: true},
}

var fields_register_image = []leanruntime.Field{
	{Name: "Architecture", Flag: "architecture", Type: "types.ArchitectureValues", Required: false},
	{Name: "BillingProducts", Flag: "billing-products", Type: "[]string", Required: false},
	{Name: "BlockDeviceMappings", Flag: "block-device-mappings", Type: "[]types.BlockDeviceMapping", Required: false},
	{Name: "BootMode", Flag: "boot-mode", Type: "types.BootModeValues", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EnaSupport", Flag: "ena-support", Type: "*bool", Required: false},
	{Name: "ImageLocation", Flag: "image-location", Type: "*string", Required: false},
	{Name: "ImdsSupport", Flag: "imds-support", Type: "types.ImdsSupportValues", Required: false},
	{Name: "KernelId", Flag: "kernel-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RamdiskId", Flag: "ramdisk-id", Type: "*string", Required: false},
	{Name: "RootDeviceName", Flag: "root-device-name", Type: "*string", Required: false},
	{Name: "SriovNetSupport", Flag: "sriov-net-support", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "TpmSupport", Flag: "tpm-support", Type: "types.TpmSupportValues", Required: false},
	{Name: "UefiData", Flag: "uefi-data", Type: "*string", Required: false},
	{Name: "VirtualizationType", Flag: "virtualization-type", Type: "*string", Required: false},
}

var fields_register_instance_event_notification_attributes = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceTagAttribute", Flag: "instance-tag-attribute", Type: "*types.RegisterInstanceTagAttributeRequest", Required: true},
}

var fields_register_transit_gateway_multicast_group_members = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GroupIpAddress", Flag: "group-ip-address", Type: "*string", Required: false},
	{Name: "NetworkInterfaceIds", Flag: "network-interface-ids", Type: "[]string", Required: true},
	{Name: "TransitGatewayMulticastDomainId", Flag: "transit-gateway-multicast-domain-id", Type: "*string", Required: true},
}

var fields_register_transit_gateway_multicast_group_sources = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GroupIpAddress", Flag: "group-ip-address", Type: "*string", Required: false},
	{Name: "NetworkInterfaceIds", Flag: "network-interface-ids", Type: "[]string", Required: true},
	{Name: "TransitGatewayMulticastDomainId", Flag: "transit-gateway-multicast-domain-id", Type: "*string", Required: true},
}

var fields_reject_capacity_reservation_billing_ownership = []leanruntime.Field{
	{Name: "CapacityReservationId", Flag: "capacity-reservation-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_reject_transit_gateway_multicast_domain_associations = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: false},
	{Name: "TransitGatewayAttachmentId", Flag: "transit-gateway-attachment-id", Type: "*string", Required: false},
	{Name: "TransitGatewayMulticastDomainId", Flag: "transit-gateway-multicast-domain-id", Type: "*string", Required: false},
}

var fields_reject_transit_gateway_peering_attachment = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TransitGatewayAttachmentId", Flag: "transit-gateway-attachment-id", Type: "*string", Required: true},
}

var fields_reject_transit_gateway_vpc_attachment = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TransitGatewayAttachmentId", Flag: "transit-gateway-attachment-id", Type: "*string", Required: true},
}

var fields_reject_vpc_endpoint_connections = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ServiceId", Flag: "service-id", Type: "*string", Required: true},
	{Name: "VpcEndpointIds", Flag: "vpc-endpoint-ids", Type: "[]string", Required: true},
}

var fields_reject_vpc_peering_connection = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VpcPeeringConnectionId", Flag: "vpc-peering-connection-id", Type: "*string", Required: true},
}

var fields_release_address = []leanruntime.Field{
	{Name: "AllocationId", Flag: "allocation-id", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "NetworkBorderGroup", Flag: "network-border-group", Type: "*string", Required: false},
	{Name: "PublicIp", Flag: "public-ip", Type: "*string", Required: false},
}

var fields_release_hosts = []leanruntime.Field{
	{Name: "HostIds", Flag: "host-ids", Type: "[]string", Required: true},
}

var fields_release_ipam_pool_allocation = []leanruntime.Field{
	{Name: "Cidr", Flag: "cidr", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "IpamPoolAllocationId", Flag: "ipam-pool-allocation-id", Type: "*string", Required: true},
	{Name: "IpamPoolId", Flag: "ipam-pool-id", Type: "*string", Required: true},
}

var fields_replace_iam_instance_profile_association = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
	{Name: "IamInstanceProfile", Flag: "iam-instance-profile", Type: "*types.IamInstanceProfileSpecification", Required: true},
}

var fields_replace_image_criteria_in_allowed_images_settings = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ImageCriteria", Flag: "image-criteria", Type: "[]types.ImageCriterionRequest", Required: false},
}

var fields_replace_network_acl_association = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "NetworkAclId", Flag: "network-acl-id", Type: "*string", Required: true},
}

var fields_replace_network_acl_entry = []leanruntime.Field{
	{Name: "CidrBlock", Flag: "cidr-block", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Egress", Flag: "egress", Type: "*bool", Required: true},
	{Name: "IcmpTypeCode", Flag: "icmp-type-code", Type: "*types.IcmpTypeCode", Required: false},
	{Name: "Ipv6CidrBlock", Flag: "ipv6-cidr-block", Type: "*string", Required: false},
	{Name: "NetworkAclId", Flag: "network-acl-id", Type: "*string", Required: true},
	{Name: "PortRange", Flag: "port-range", Type: "*types.PortRange", Required: false},
	{Name: "Protocol", Flag: "protocol", Type: "*string", Required: true},
	{Name: "RuleAction", Flag: "rule-action", Type: "types.RuleAction", Required: true},
	{Name: "RuleNumber", Flag: "rule-number", Type: "*int32", Required: true},
}

var fields_replace_route = []leanruntime.Field{
	{Name: "CarrierGatewayId", Flag: "carrier-gateway-id", Type: "*string", Required: false},
	{Name: "CoreNetworkArn", Flag: "core-network-arn", Type: "*string", Required: false},
	{Name: "DestinationCidrBlock", Flag: "destination-cidr-block", Type: "*string", Required: false},
	{Name: "DestinationIpv6CidrBlock", Flag: "destination-ipv6-cidr-block", Type: "*string", Required: false},
	{Name: "DestinationPrefixListId", Flag: "destination-prefix-list-id", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EgressOnlyInternetGatewayId", Flag: "egress-only-internet-gateway-id", Type: "*string", Required: false},
	{Name: "GatewayId", Flag: "gateway-id", Type: "*string", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: false},
	{Name: "LocalGatewayId", Flag: "local-gateway-id", Type: "*string", Required: false},
	{Name: "LocalTarget", Flag: "local-target", Type: "*bool", Required: false},
	{Name: "NatGatewayId", Flag: "nat-gateway-id", Type: "*string", Required: false},
	{Name: "NetworkInterfaceId", Flag: "network-interface-id", Type: "*string", Required: false},
	{Name: "OdbNetworkArn", Flag: "odb-network-arn", Type: "*string", Required: false},
	{Name: "RouteTableId", Flag: "route-table-id", Type: "*string", Required: true},
	{Name: "TransitGatewayId", Flag: "transit-gateway-id", Type: "*string", Required: false},
	{Name: "VpcEndpointId", Flag: "vpc-endpoint-id", Type: "*string", Required: false},
	{Name: "VpcPeeringConnectionId", Flag: "vpc-peering-connection-id", Type: "*string", Required: false},
}

var fields_replace_route_table_association = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "RouteTableId", Flag: "route-table-id", Type: "*string", Required: true},
}

var fields_replace_transit_gateway_route = []leanruntime.Field{
	{Name: "Blackhole", Flag: "blackhole", Type: "*bool", Required: false},
	{Name: "DestinationCidrBlock", Flag: "destination-cidr-block", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TransitGatewayAttachmentId", Flag: "transit-gateway-attachment-id", Type: "*string", Required: false},
	{Name: "TransitGatewayRouteTableId", Flag: "transit-gateway-route-table-id", Type: "*string", Required: true},
}

var fields_replace_vpn_tunnel = []leanruntime.Field{
	{Name: "ApplyPendingMaintenance", Flag: "apply-pending-maintenance", Type: "*bool", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VpnConnectionId", Flag: "vpn-connection-id", Type: "*string", Required: true},
	{Name: "VpnTunnelOutsideIpAddress", Flag: "vpn-tunnel-outside-ip-address", Type: "*string", Required: true},
}

var fields_report_instance_status = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "Instances", Flag: "instances", Type: "[]string", Required: true},
	{Name: "ReasonCodes", Flag: "reason-codes", Type: "[]types.ReportInstanceReasonCodes", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
	{Name: "Status", Flag: "status", Type: "types.ReportStatusType", Required: true},
}

var fields_request_spot_fleet = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "SpotFleetRequestConfig", Flag: "spot-fleet-request-config", Type: "*types.SpotFleetRequestConfigData", Required: true},
}

var fields_request_spot_instances = []leanruntime.Field{
	{Name: "AvailabilityZoneGroup", Flag: "availability-zone-group", Type: "*string", Required: false},
	{Name: "BlockDurationMinutes", Flag: "block-duration-minutes", Type: "*int32", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceCount", Flag: "instance-count", Type: "*int32", Required: false},
	{Name: "InstanceInterruptionBehavior", Flag: "instance-interruption-behavior", Type: "types.InstanceInterruptionBehavior", Required: false},
	{Name: "LaunchGroup", Flag: "launch-group", Type: "*string", Required: false},
	{Name: "LaunchSpecification", Flag: "launch-specification", Type: "*types.RequestSpotLaunchSpecification", Required: false},
	{Name: "SpotPrice", Flag: "spot-price", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "Type", Flag: "type", Type: "types.SpotInstanceType", Required: false},
	{Name: "ValidFrom", Flag: "valid-from", Type: "*time.Time", Required: false},
	{Name: "ValidUntil", Flag: "valid-until", Type: "*time.Time", Required: false},
}

var fields_reset_address_attribute = []leanruntime.Field{
	{Name: "AllocationId", Flag: "allocation-id", Type: "*string", Required: true},
	{Name: "Attribute", Flag: "attribute", Type: "types.AddressAttributeName", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_reset_ebs_default_kms_key_id = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_reset_fpga_image_attribute = []leanruntime.Field{
	{Name: "Attribute", Flag: "attribute", Type: "types.ResetFpgaImageAttributeName", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "FpgaImageId", Flag: "fpga-image-id", Type: "*string", Required: true},
}

var fields_reset_image_attribute = []leanruntime.Field{
	{Name: "Attribute", Flag: "attribute", Type: "types.ResetImageAttributeName", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: true},
}

var fields_reset_instance_attribute = []leanruntime.Field{
	{Name: "Attribute", Flag: "attribute", Type: "types.InstanceAttributeName", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_reset_network_interface_attribute = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "NetworkInterfaceId", Flag: "network-interface-id", Type: "*string", Required: true},
	{Name: "SourceDestCheck", Flag: "source-dest-check", Type: "*string", Required: false},
}

var fields_reset_snapshot_attribute = []leanruntime.Field{
	{Name: "Attribute", Flag: "attribute", Type: "types.SnapshotAttributeName", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "SnapshotId", Flag: "snapshot-id", Type: "*string", Required: true},
}

var fields_restore_address_to_classic = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "PublicIp", Flag: "public-ip", Type: "*string", Required: true},
}

var fields_restore_image_from_recycle_bin = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: true},
}

var fields_restore_managed_prefix_list_version = []leanruntime.Field{
	{Name: "CurrentVersion", Flag: "current-version", Type: "*int64", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "PrefixListId", Flag: "prefix-list-id", Type: "*string", Required: true},
	{Name: "PreviousVersion", Flag: "previous-version", Type: "*int64", Required: true},
}

var fields_restore_snapshot_from_recycle_bin = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "SnapshotId", Flag: "snapshot-id", Type: "*string", Required: true},
}

var fields_restore_snapshot_tier = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "PermanentRestore", Flag: "permanent-restore", Type: "*bool", Required: false},
	{Name: "SnapshotId", Flag: "snapshot-id", Type: "*string", Required: true},
	{Name: "TemporaryRestoreDays", Flag: "temporary-restore-days", Type: "*int32", Required: false},
}

var fields_restore_volume_from_recycle_bin = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "VolumeId", Flag: "volume-id", Type: "*string", Required: true},
}

var fields_revoke_client_vpn_ingress = []leanruntime.Field{
	{Name: "AccessGroupId", Flag: "access-group-id", Type: "*string", Required: false},
	{Name: "ClientVpnEndpointId", Flag: "client-vpn-endpoint-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "RevokeAllGroups", Flag: "revoke-all-groups", Type: "*bool", Required: false},
	{Name: "TargetNetworkCidr", Flag: "target-network-cidr", Type: "*string", Required: true},
}

var fields_revoke_security_group_egress = []leanruntime.Field{
	{Name: "CidrIp", Flag: "cidr-ip", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "FromPort", Flag: "from-port", Type: "*int32", Required: false},
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: true},
	{Name: "IpPermissions", Flag: "ip-permissions", Type: "[]types.IpPermission", Required: false},
	{Name: "IpProtocol", Flag: "ip-protocol", Type: "*string", Required: false},
	{Name: "SecurityGroupRuleIds", Flag: "security-group-rule-ids", Type: "[]string", Required: false},
	{Name: "SourceSecurityGroupName", Flag: "source-security-group-name", Type: "*string", Required: false},
	{Name: "SourceSecurityGroupOwnerId", Flag: "source-security-group-owner-id", Type: "*string", Required: false},
	{Name: "ToPort", Flag: "to-port", Type: "*int32", Required: false},
}

var fields_revoke_security_group_ingress = []leanruntime.Field{
	{Name: "CidrIp", Flag: "cidr-ip", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "FromPort", Flag: "from-port", Type: "*int32", Required: false},
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: false},
	{Name: "IpPermissions", Flag: "ip-permissions", Type: "[]types.IpPermission", Required: false},
	{Name: "IpProtocol", Flag: "ip-protocol", Type: "*string", Required: false},
	{Name: "SecurityGroupRuleIds", Flag: "security-group-rule-ids", Type: "[]string", Required: false},
	{Name: "SourceSecurityGroupName", Flag: "source-security-group-name", Type: "*string", Required: false},
	{Name: "SourceSecurityGroupOwnerId", Flag: "source-security-group-owner-id", Type: "*string", Required: false},
	{Name: "ToPort", Flag: "to-port", Type: "*int32", Required: false},
}

var fields_run_instances = []leanruntime.Field{
	{Name: "AdditionalInfo", Flag: "additional-info", Type: "*string", Required: false},
	{Name: "BlockDeviceMappings", Flag: "block-device-mappings", Type: "[]types.BlockDeviceMapping", Required: false},
	{Name: "CapacityReservationSpecification", Flag: "capacity-reservation-specification", Type: "*types.CapacityReservationSpecification", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CpuOptions", Flag: "cpu-options", Type: "*types.CpuOptionsRequest", Required: false},
	{Name: "CreditSpecification", Flag: "credit-specification", Type: "*types.CreditSpecificationRequest", Required: false},
	{Name: "DisableApiStop", Flag: "disable-api-stop", Type: "*bool", Required: false},
	{Name: "DisableApiTermination", Flag: "disable-api-termination", Type: "*bool", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EbsOptimized", Flag: "ebs-optimized", Type: "*bool", Required: false},
	{Name: "ElasticGpuSpecification", Flag: "elastic-gpu-specification", Type: "[]types.ElasticGpuSpecification", Required: false},
	{Name: "ElasticInferenceAccelerators", Flag: "elastic-inference-accelerators", Type: "[]types.ElasticInferenceAccelerator", Required: false},
	{Name: "EnablePrimaryIpv6", Flag: "enable-primary-ipv6", Type: "*bool", Required: false},
	{Name: "EnclaveOptions", Flag: "enclave-options", Type: "*types.EnclaveOptionsRequest", Required: false},
	{Name: "HibernationOptions", Flag: "hibernation-options", Type: "*types.HibernationOptionsRequest", Required: false},
	{Name: "IamInstanceProfile", Flag: "iam-instance-profile", Type: "*types.IamInstanceProfileSpecification", Required: false},
	{Name: "ImageId", Flag: "image-id", Type: "*string", Required: false},
	{Name: "InstanceInitiatedShutdownBehavior", Flag: "instance-initiated-shutdown-behavior", Type: "types.ShutdownBehavior", Required: false},
	{Name: "InstanceMarketOptions", Flag: "instance-market-options", Type: "*types.InstanceMarketOptionsRequest", Required: false},
	{Name: "InstanceType", Flag: "instance-type", Type: "types.InstanceType", Required: false},
	{Name: "Ipv6AddressCount", Flag: "ipv6-address-count", Type: "*int32", Required: false},
	{Name: "Ipv6Addresses", Flag: "ipv6-addresses", Type: "[]types.InstanceIpv6Address", Required: false},
	{Name: "KernelId", Flag: "kernel-id", Type: "*string", Required: false},
	{Name: "KeyName", Flag: "key-name", Type: "*string", Required: false},
	{Name: "LaunchTemplate", Flag: "launch-template", Type: "*types.LaunchTemplateSpecification", Required: false},
	{Name: "LicenseSpecifications", Flag: "license-specifications", Type: "[]types.LicenseConfigurationRequest", Required: false},
	{Name: "MaintenanceOptions", Flag: "maintenance-options", Type: "*types.InstanceMaintenanceOptionsRequest", Required: false},
	{Name: "MaxCount", Flag: "max-count", Type: "*int32", Required: true},
	{Name: "MetadataOptions", Flag: "metadata-options", Type: "*types.InstanceMetadataOptionsRequest", Required: false},
	{Name: "MinCount", Flag: "min-count", Type: "*int32", Required: true},
	{Name: "Monitoring", Flag: "monitoring", Type: "*types.RunInstancesMonitoringEnabled", Required: false},
	{Name: "NetworkInterfaces", Flag: "network-interfaces", Type: "[]types.InstanceNetworkInterfaceSpecification", Required: false},
	{Name: "NetworkPerformanceOptions", Flag: "network-performance-options", Type: "*types.InstanceNetworkPerformanceOptionsRequest", Required: false},
	{Name: "Operator", Flag: "operator", Type: "*types.OperatorRequest", Required: false},
	{Name: "Placement", Flag: "placement", Type: "*types.Placement", Required: false},
	{Name: "PrivateDnsNameOptions", Flag: "private-dns-name-options", Type: "*types.PrivateDnsNameOptionsRequest", Required: false},
	{Name: "PrivateIpAddress", Flag: "private-ip-address", Type: "*string", Required: false},
	{Name: "RamdiskId", Flag: "ramdisk-id", Type: "*string", Required: false},
	{Name: "SecondaryInterfaces", Flag: "secondary-interfaces", Type: "[]types.InstanceSecondaryInterfaceSpecificationRequest", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "SecurityGroups", Flag: "security-groups", Type: "[]string", Required: false},
	{Name: "SubnetId", Flag: "subnet-id", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "UserData", Flag: "user-data", Type: "*string", Required: false},
}

var fields_run_scheduled_instances = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceCount", Flag: "instance-count", Type: "*int32", Required: false},
	{Name: "LaunchSpecification", Flag: "launch-specification", Type: "*types.ScheduledInstancesLaunchSpecification", Required: true},
	{Name: "ScheduledInstanceId", Flag: "scheduled-instance-id", Type: "*string", Required: true},
}

var fields_search_local_gateway_routes = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "LocalGatewayRouteTableId", Flag: "local-gateway-route-table-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_search_transit_gateway_multicast_groups = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TransitGatewayMulticastDomainId", Flag: "transit-gateway-multicast-domain-id", Type: "*string", Required: true},
}

var fields_search_transit_gateway_routes = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TransitGatewayRouteTableId", Flag: "transit-gateway-route-table-id", Type: "*string", Required: true},
}

var fields_send_diagnostic_interrupt = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceId", Flag: "instance-id", Type: "*string", Required: true},
}

var fields_start_declarative_policies_report = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "S3Bucket", Flag: "s3-bucket", Type: "*string", Required: true},
	{Name: "S3Prefix", Flag: "s3-prefix", Type: "*string", Required: false},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
	{Name: "TargetId", Flag: "target-id", Type: "*string", Required: true},
}

var fields_start_instances = []leanruntime.Field{
	{Name: "AdditionalInfo", Flag: "additional-info", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: true},
}

var fields_start_network_insights_access_scope_analysis = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "NetworkInsightsAccessScopeId", Flag: "network-insights-access-scope-id", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_start_network_insights_analysis = []leanruntime.Field{
	{Name: "AdditionalAccounts", Flag: "additional-accounts", Type: "[]string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "FilterInArns", Flag: "filter-in-arns", Type: "[]string", Required: false},
	{Name: "FilterOutArns", Flag: "filter-out-arns", Type: "[]string", Required: false},
	{Name: "NetworkInsightsPathId", Flag: "network-insights-path-id", Type: "*string", Required: true},
	{Name: "TagSpecifications", Flag: "tag-specifications", Type: "[]types.TagSpecification", Required: false},
}

var fields_start_vpc_endpoint_service_private_dns_verification = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "ServiceId", Flag: "service-id", Type: "*string", Required: true},
}

var fields_stop_instances = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Force", Flag: "force", Type: "*bool", Required: false},
	{Name: "Hibernate", Flag: "hibernate", Type: "*bool", Required: false},
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: true},
	{Name: "SkipOsShutdown", Flag: "skip-os-shutdown", Type: "*bool", Required: false},
}

var fields_terminate_client_vpn_connections = []leanruntime.Field{
	{Name: "ClientVpnEndpointId", Flag: "client-vpn-endpoint-id", Type: "*string", Required: true},
	{Name: "ConnectionId", Flag: "connection-id", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: false},
}

var fields_terminate_instances = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "Force", Flag: "force", Type: "*bool", Required: false},
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: true},
	{Name: "SkipOsShutdown", Flag: "skip-os-shutdown", Type: "*bool", Required: false},
}

var fields_unassign_ipv6_addresses = []leanruntime.Field{
	{Name: "Ipv6Addresses", Flag: "ipv6-addresses", Type: "[]string", Required: false},
	{Name: "Ipv6Prefixes", Flag: "ipv6-prefixes", Type: "[]string", Required: false},
	{Name: "NetworkInterfaceId", Flag: "network-interface-id", Type: "*string", Required: true},
}

var fields_unassign_private_ip_addresses = []leanruntime.Field{
	{Name: "Ipv4Prefixes", Flag: "ipv4-prefixes", Type: "[]string", Required: false},
	{Name: "NetworkInterfaceId", Flag: "network-interface-id", Type: "*string", Required: true},
	{Name: "PrivateIpAddresses", Flag: "private-ip-addresses", Type: "[]string", Required: false},
}

var fields_unassign_private_nat_gateway_address = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "MaxDrainDurationSeconds", Flag: "max-drain-duration-seconds", Type: "*int32", Required: false},
	{Name: "NatGatewayId", Flag: "nat-gateway-id", Type: "*string", Required: true},
	{Name: "PrivateIpAddresses", Flag: "private-ip-addresses", Type: "[]string", Required: true},
}

var fields_unlock_snapshot = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "SnapshotId", Flag: "snapshot-id", Type: "*string", Required: true},
}

var fields_unmonitor_instances = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "InstanceIds", Flag: "instance-ids", Type: "[]string", Required: true},
}

var fields_update_capacity_manager_organizations_access = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "OrganizationsAccess", Flag: "organizations-access", Type: "*bool", Required: true},
}

var fields_update_interruptible_capacity_reservation_allocation = []leanruntime.Field{
	{Name: "CapacityReservationId", Flag: "capacity-reservation-id", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "TargetInstanceCount", Flag: "target-instance-count", Type: "*int32", Required: true},
}

var fields_update_security_group_rule_descriptions_egress = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: false},
	{Name: "IpPermissions", Flag: "ip-permissions", Type: "[]types.IpPermission", Required: false},
	{Name: "SecurityGroupRuleDescriptions", Flag: "security-group-rule-descriptions", Type: "[]types.SecurityGroupRuleDescription", Required: false},
}

var fields_update_security_group_rule_descriptions_ingress = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "GroupId", Flag: "group-id", Type: "*string", Required: false},
	{Name: "GroupName", Flag: "group-name", Type: "*string", Required: false},
	{Name: "IpPermissions", Flag: "ip-permissions", Type: "[]types.IpPermission", Required: false},
	{Name: "SecurityGroupRuleDescriptions", Flag: "security-group-rule-descriptions", Type: "[]types.SecurityGroupRuleDescription", Required: false},
}

var fields_withdraw_byoip_cidr = []leanruntime.Field{
	{Name: "Cidr", Flag: "cidr", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-address-transfer": {
			Name:   "accept-address-transfer",
			Fields: fields_accept_address_transfer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptAddressTransferInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_address_transfer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptAddressTransfer(ctx, input)
			},
		},
		"accept-capacity-reservation-billing-ownership": {
			Name:   "accept-capacity-reservation-billing-ownership",
			Fields: fields_accept_capacity_reservation_billing_ownership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptCapacityReservationBillingOwnershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_capacity_reservation_billing_ownership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptCapacityReservationBillingOwnership(ctx, input)
			},
		},
		"accept-reserved-instances-exchange-quote": {
			Name:   "accept-reserved-instances-exchange-quote",
			Fields: fields_accept_reserved_instances_exchange_quote,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptReservedInstancesExchangeQuoteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_reserved_instances_exchange_quote, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptReservedInstancesExchangeQuote(ctx, input)
			},
		},
		"accept-transit-gateway-multicast-domain-associations": {
			Name:   "accept-transit-gateway-multicast-domain-associations",
			Fields: fields_accept_transit_gateway_multicast_domain_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptTransitGatewayMulticastDomainAssociationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_transit_gateway_multicast_domain_associations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptTransitGatewayMulticastDomainAssociations(ctx, input)
			},
		},
		"accept-transit-gateway-peering-attachment": {
			Name:   "accept-transit-gateway-peering-attachment",
			Fields: fields_accept_transit_gateway_peering_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptTransitGatewayPeeringAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_transit_gateway_peering_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptTransitGatewayPeeringAttachment(ctx, input)
			},
		},
		"accept-transit-gateway-vpc-attachment": {
			Name:   "accept-transit-gateway-vpc-attachment",
			Fields: fields_accept_transit_gateway_vpc_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptTransitGatewayVpcAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_transit_gateway_vpc_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptTransitGatewayVpcAttachment(ctx, input)
			},
		},
		"accept-vpc-endpoint-connections": {
			Name:   "accept-vpc-endpoint-connections",
			Fields: fields_accept_vpc_endpoint_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptVpcEndpointConnectionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_vpc_endpoint_connections, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptVpcEndpointConnections(ctx, input)
			},
		},
		"accept-vpc-peering-connection": {
			Name:   "accept-vpc-peering-connection",
			Fields: fields_accept_vpc_peering_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptVpcPeeringConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_vpc_peering_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptVpcPeeringConnection(ctx, input)
			},
		},
		"advertise-byoip-cidr": {
			Name:   "advertise-byoip-cidr",
			Fields: fields_advertise_byoip_cidr,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AdvertiseByoipCidrInput{}
				if _, err := leanruntime.ApplyInput(input, fields_advertise_byoip_cidr, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AdvertiseByoipCidr(ctx, input)
			},
		},
		"allocate-address": {
			Name:   "allocate-address",
			Fields: fields_allocate_address,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AllocateAddressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_allocate_address, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AllocateAddress(ctx, input)
			},
		},
		"allocate-hosts": {
			Name:   "allocate-hosts",
			Fields: fields_allocate_hosts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AllocateHostsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_allocate_hosts, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AllocateHosts(ctx, input)
			},
		},
		"allocate-ipam-pool-cidr": {
			Name:   "allocate-ipam-pool-cidr",
			Fields: fields_allocate_ipam_pool_cidr,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AllocateIpamPoolCidrInput{}
				if _, err := leanruntime.ApplyInput(input, fields_allocate_ipam_pool_cidr, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AllocateIpamPoolCidr(ctx, input)
			},
		},
		"apply-security-groups-to-client-vpn-target-network": {
			Name:   "apply-security-groups-to-client-vpn-target-network",
			Fields: fields_apply_security_groups_to_client_vpn_target_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ApplySecurityGroupsToClientVpnTargetNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_apply_security_groups_to_client_vpn_target_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ApplySecurityGroupsToClientVpnTargetNetwork(ctx, input)
			},
		},
		"assign-ipv6-addresses": {
			Name:   "assign-ipv6-addresses",
			Fields: fields_assign_ipv6_addresses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssignIpv6AddressesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_assign_ipv6_addresses, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssignIpv6Addresses(ctx, input)
			},
		},
		"assign-private-ip-addresses": {
			Name:   "assign-private-ip-addresses",
			Fields: fields_assign_private_ip_addresses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssignPrivateIpAddressesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_assign_private_ip_addresses, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssignPrivateIpAddresses(ctx, input)
			},
		},
		"assign-private-nat-gateway-address": {
			Name:   "assign-private-nat-gateway-address",
			Fields: fields_assign_private_nat_gateway_address,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssignPrivateNatGatewayAddressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_assign_private_nat_gateway_address, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssignPrivateNatGatewayAddress(ctx, input)
			},
		},
		"associate-address": {
			Name:   "associate-address",
			Fields: fields_associate_address,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateAddressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_address, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateAddress(ctx, input)
			},
		},
		"associate-capacity-reservation-billing-owner": {
			Name:   "associate-capacity-reservation-billing-owner",
			Fields: fields_associate_capacity_reservation_billing_owner,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateCapacityReservationBillingOwnerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_capacity_reservation_billing_owner, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateCapacityReservationBillingOwner(ctx, input)
			},
		},
		"associate-client-vpn-target-network": {
			Name:   "associate-client-vpn-target-network",
			Fields: fields_associate_client_vpn_target_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateClientVpnTargetNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_client_vpn_target_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateClientVpnTargetNetwork(ctx, input)
			},
		},
		"associate-dhcp-options": {
			Name:   "associate-dhcp-options",
			Fields: fields_associate_dhcp_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateDhcpOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_dhcp_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateDhcpOptions(ctx, input)
			},
		},
		"associate-enclave-certificate-iam-role": {
			Name:   "associate-enclave-certificate-iam-role",
			Fields: fields_associate_enclave_certificate_iam_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateEnclaveCertificateIamRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_enclave_certificate_iam_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateEnclaveCertificateIamRole(ctx, input)
			},
		},
		"associate-iam-instance-profile": {
			Name:   "associate-iam-instance-profile",
			Fields: fields_associate_iam_instance_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateIamInstanceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_iam_instance_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateIamInstanceProfile(ctx, input)
			},
		},
		"associate-instance-event-window": {
			Name:   "associate-instance-event-window",
			Fields: fields_associate_instance_event_window,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateInstanceEventWindowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_instance_event_window, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateInstanceEventWindow(ctx, input)
			},
		},
		"associate-ipam-byoasn": {
			Name:   "associate-ipam-byoasn",
			Fields: fields_associate_ipam_byoasn,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateIpamByoasnInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_ipam_byoasn, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateIpamByoasn(ctx, input)
			},
		},
		"associate-ipam-resource-discovery": {
			Name:   "associate-ipam-resource-discovery",
			Fields: fields_associate_ipam_resource_discovery,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateIpamResourceDiscoveryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_ipam_resource_discovery, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateIpamResourceDiscovery(ctx, input)
			},
		},
		"associate-nat-gateway-address": {
			Name:   "associate-nat-gateway-address",
			Fields: fields_associate_nat_gateway_address,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateNatGatewayAddressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_nat_gateway_address, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateNatGatewayAddress(ctx, input)
			},
		},
		"associate-route-server": {
			Name:   "associate-route-server",
			Fields: fields_associate_route_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateRouteServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_route_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateRouteServer(ctx, input)
			},
		},
		"associate-route-table": {
			Name:   "associate-route-table",
			Fields: fields_associate_route_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateRouteTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_route_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateRouteTable(ctx, input)
			},
		},
		"associate-security-group-vpc": {
			Name:   "associate-security-group-vpc",
			Fields: fields_associate_security_group_vpc,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateSecurityGroupVpcInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_security_group_vpc, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateSecurityGroupVpc(ctx, input)
			},
		},
		"associate-subnet-cidr-block": {
			Name:   "associate-subnet-cidr-block",
			Fields: fields_associate_subnet_cidr_block,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateSubnetCidrBlockInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_subnet_cidr_block, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateSubnetCidrBlock(ctx, input)
			},
		},
		"associate-transit-gateway-multicast-domain": {
			Name:   "associate-transit-gateway-multicast-domain",
			Fields: fields_associate_transit_gateway_multicast_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateTransitGatewayMulticastDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_transit_gateway_multicast_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateTransitGatewayMulticastDomain(ctx, input)
			},
		},
		"associate-transit-gateway-policy-table": {
			Name:   "associate-transit-gateway-policy-table",
			Fields: fields_associate_transit_gateway_policy_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateTransitGatewayPolicyTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_transit_gateway_policy_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateTransitGatewayPolicyTable(ctx, input)
			},
		},
		"associate-transit-gateway-route-table": {
			Name:   "associate-transit-gateway-route-table",
			Fields: fields_associate_transit_gateway_route_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateTransitGatewayRouteTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_transit_gateway_route_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateTransitGatewayRouteTable(ctx, input)
			},
		},
		"associate-trunk-interface": {
			Name:   "associate-trunk-interface",
			Fields: fields_associate_trunk_interface,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateTrunkInterfaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_trunk_interface, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateTrunkInterface(ctx, input)
			},
		},
		"associate-vpc-cidr-block": {
			Name:   "associate-vpc-cidr-block",
			Fields: fields_associate_vpc_cidr_block,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateVpcCidrBlockInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_vpc_cidr_block, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateVpcCidrBlock(ctx, input)
			},
		},
		"attach-classic-link-vpc": {
			Name:   "attach-classic-link-vpc",
			Fields: fields_attach_classic_link_vpc,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachClassicLinkVpcInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_classic_link_vpc, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachClassicLinkVpc(ctx, input)
			},
		},
		"attach-internet-gateway": {
			Name:   "attach-internet-gateway",
			Fields: fields_attach_internet_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachInternetGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_internet_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachInternetGateway(ctx, input)
			},
		},
		"attach-network-interface": {
			Name:   "attach-network-interface",
			Fields: fields_attach_network_interface,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachNetworkInterfaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_network_interface, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachNetworkInterface(ctx, input)
			},
		},
		"attach-verified-access-trust-provider": {
			Name:   "attach-verified-access-trust-provider",
			Fields: fields_attach_verified_access_trust_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachVerifiedAccessTrustProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_verified_access_trust_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachVerifiedAccessTrustProvider(ctx, input)
			},
		},
		"attach-volume": {
			Name:   "attach-volume",
			Fields: fields_attach_volume,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachVolumeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_volume, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachVolume(ctx, input)
			},
		},
		"attach-vpn-gateway": {
			Name:   "attach-vpn-gateway",
			Fields: fields_attach_vpn_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachVpnGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_vpn_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachVpnGateway(ctx, input)
			},
		},
		"authorize-client-vpn-ingress": {
			Name:   "authorize-client-vpn-ingress",
			Fields: fields_authorize_client_vpn_ingress,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AuthorizeClientVpnIngressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_authorize_client_vpn_ingress, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AuthorizeClientVpnIngress(ctx, input)
			},
		},
		"authorize-security-group-egress": {
			Name:   "authorize-security-group-egress",
			Fields: fields_authorize_security_group_egress,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AuthorizeSecurityGroupEgressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_authorize_security_group_egress, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AuthorizeSecurityGroupEgress(ctx, input)
			},
		},
		"authorize-security-group-ingress": {
			Name:   "authorize-security-group-ingress",
			Fields: fields_authorize_security_group_ingress,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AuthorizeSecurityGroupIngressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_authorize_security_group_ingress, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AuthorizeSecurityGroupIngress(ctx, input)
			},
		},
		"bundle-instance": {
			Name:   "bundle-instance",
			Fields: fields_bundle_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BundleInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_bundle_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BundleInstance(ctx, input)
			},
		},
		"cancel-bundle-task": {
			Name:   "cancel-bundle-task",
			Fields: fields_cancel_bundle_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelBundleTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_bundle_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelBundleTask(ctx, input)
			},
		},
		"cancel-capacity-reservation": {
			Name:   "cancel-capacity-reservation",
			Fields: fields_cancel_capacity_reservation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelCapacityReservationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_capacity_reservation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelCapacityReservation(ctx, input)
			},
		},
		"cancel-capacity-reservation-fleets": {
			Name:   "cancel-capacity-reservation-fleets",
			Fields: fields_cancel_capacity_reservation_fleets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelCapacityReservationFleetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_capacity_reservation_fleets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelCapacityReservationFleets(ctx, input)
			},
		},
		"cancel-conversion-task": {
			Name:   "cancel-conversion-task",
			Fields: fields_cancel_conversion_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelConversionTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_conversion_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelConversionTask(ctx, input)
			},
		},
		"cancel-declarative-policies-report": {
			Name:   "cancel-declarative-policies-report",
			Fields: fields_cancel_declarative_policies_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelDeclarativePoliciesReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_declarative_policies_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelDeclarativePoliciesReport(ctx, input)
			},
		},
		"cancel-export-task": {
			Name:   "cancel-export-task",
			Fields: fields_cancel_export_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelExportTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_export_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelExportTask(ctx, input)
			},
		},
		"cancel-image-launch-permission": {
			Name:   "cancel-image-launch-permission",
			Fields: fields_cancel_image_launch_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelImageLaunchPermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_image_launch_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelImageLaunchPermission(ctx, input)
			},
		},
		"cancel-import-task": {
			Name:   "cancel-import-task",
			Fields: fields_cancel_import_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelImportTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_import_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelImportTask(ctx, input)
			},
		},
		"cancel-reserved-instances-listing": {
			Name:   "cancel-reserved-instances-listing",
			Fields: fields_cancel_reserved_instances_listing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelReservedInstancesListingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_reserved_instances_listing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelReservedInstancesListing(ctx, input)
			},
		},
		"cancel-spot-fleet-requests": {
			Name:   "cancel-spot-fleet-requests",
			Fields: fields_cancel_spot_fleet_requests,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelSpotFleetRequestsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_spot_fleet_requests, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelSpotFleetRequests(ctx, input)
			},
		},
		"cancel-spot-instance-requests": {
			Name:   "cancel-spot-instance-requests",
			Fields: fields_cancel_spot_instance_requests,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelSpotInstanceRequestsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_spot_instance_requests, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelSpotInstanceRequests(ctx, input)
			},
		},
		"confirm-product-instance": {
			Name:   "confirm-product-instance",
			Fields: fields_confirm_product_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ConfirmProductInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_confirm_product_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ConfirmProductInstance(ctx, input)
			},
		},
		"copy-fpga-image": {
			Name:   "copy-fpga-image",
			Fields: fields_copy_fpga_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CopyFpgaImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_copy_fpga_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CopyFpgaImage(ctx, input)
			},
		},
		"copy-image": {
			Name:   "copy-image",
			Fields: fields_copy_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CopyImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_copy_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CopyImage(ctx, input)
			},
		},
		"copy-snapshot": {
			Name:   "copy-snapshot",
			Fields: fields_copy_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CopySnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_copy_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CopySnapshot(ctx, input)
			},
		},
		"copy-volumes": {
			Name:   "copy-volumes",
			Fields: fields_copy_volumes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CopyVolumesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_copy_volumes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CopyVolumes(ctx, input)
			},
		},
		"create-capacity-manager-data-export": {
			Name:   "create-capacity-manager-data-export",
			Fields: fields_create_capacity_manager_data_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCapacityManagerDataExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_capacity_manager_data_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCapacityManagerDataExport(ctx, input)
			},
		},
		"create-capacity-reservation": {
			Name:   "create-capacity-reservation",
			Fields: fields_create_capacity_reservation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCapacityReservationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_capacity_reservation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCapacityReservation(ctx, input)
			},
		},
		"create-capacity-reservation-by-splitting": {
			Name:   "create-capacity-reservation-by-splitting",
			Fields: fields_create_capacity_reservation_by_splitting,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCapacityReservationBySplittingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_capacity_reservation_by_splitting, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCapacityReservationBySplitting(ctx, input)
			},
		},
		"create-capacity-reservation-fleet": {
			Name:   "create-capacity-reservation-fleet",
			Fields: fields_create_capacity_reservation_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCapacityReservationFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_capacity_reservation_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCapacityReservationFleet(ctx, input)
			},
		},
		"create-carrier-gateway": {
			Name:   "create-carrier-gateway",
			Fields: fields_create_carrier_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCarrierGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_carrier_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCarrierGateway(ctx, input)
			},
		},
		"create-client-vpn-endpoint": {
			Name:   "create-client-vpn-endpoint",
			Fields: fields_create_client_vpn_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateClientVpnEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_client_vpn_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateClientVpnEndpoint(ctx, input)
			},
		},
		"create-client-vpn-route": {
			Name:   "create-client-vpn-route",
			Fields: fields_create_client_vpn_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateClientVpnRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_client_vpn_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateClientVpnRoute(ctx, input)
			},
		},
		"create-coip-cidr": {
			Name:   "create-coip-cidr",
			Fields: fields_create_coip_cidr,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCoipCidrInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_coip_cidr, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCoipCidr(ctx, input)
			},
		},
		"create-coip-pool": {
			Name:   "create-coip-pool",
			Fields: fields_create_coip_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCoipPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_coip_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCoipPool(ctx, input)
			},
		},
		"create-customer-gateway": {
			Name:   "create-customer-gateway",
			Fields: fields_create_customer_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCustomerGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_customer_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCustomerGateway(ctx, input)
			},
		},
		"create-default-subnet": {
			Name:   "create-default-subnet",
			Fields: fields_create_default_subnet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDefaultSubnetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_default_subnet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDefaultSubnet(ctx, input)
			},
		},
		"create-default-vpc": {
			Name:   "create-default-vpc",
			Fields: fields_create_default_vpc,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDefaultVpcInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_default_vpc, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDefaultVpc(ctx, input)
			},
		},
		"create-delegate-mac-volume-ownership-task": {
			Name:   "create-delegate-mac-volume-ownership-task",
			Fields: fields_create_delegate_mac_volume_ownership_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDelegateMacVolumeOwnershipTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_delegate_mac_volume_ownership_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDelegateMacVolumeOwnershipTask(ctx, input)
			},
		},
		"create-dhcp-options": {
			Name:   "create-dhcp-options",
			Fields: fields_create_dhcp_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDhcpOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_dhcp_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDhcpOptions(ctx, input)
			},
		},
		"create-egress-only-internet-gateway": {
			Name:   "create-egress-only-internet-gateway",
			Fields: fields_create_egress_only_internet_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEgressOnlyInternetGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_egress_only_internet_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEgressOnlyInternetGateway(ctx, input)
			},
		},
		"create-fleet": {
			Name:   "create-fleet",
			Fields: fields_create_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFleet(ctx, input)
			},
		},
		"create-flow-logs": {
			Name:   "create-flow-logs",
			Fields: fields_create_flow_logs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFlowLogsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_flow_logs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFlowLogs(ctx, input)
			},
		},
		"create-fpga-image": {
			Name:   "create-fpga-image",
			Fields: fields_create_fpga_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFpgaImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_fpga_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFpgaImage(ctx, input)
			},
		},
		"create-image": {
			Name:   "create-image",
			Fields: fields_create_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateImage(ctx, input)
			},
		},
		"create-image-usage-report": {
			Name:   "create-image-usage-report",
			Fields: fields_create_image_usage_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateImageUsageReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_image_usage_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateImageUsageReport(ctx, input)
			},
		},
		"create-instance-connect-endpoint": {
			Name:   "create-instance-connect-endpoint",
			Fields: fields_create_instance_connect_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInstanceConnectEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_instance_connect_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInstanceConnectEndpoint(ctx, input)
			},
		},
		"create-instance-event-window": {
			Name:   "create-instance-event-window",
			Fields: fields_create_instance_event_window,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInstanceEventWindowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_instance_event_window, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInstanceEventWindow(ctx, input)
			},
		},
		"create-instance-export-task": {
			Name:   "create-instance-export-task",
			Fields: fields_create_instance_export_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInstanceExportTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_instance_export_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInstanceExportTask(ctx, input)
			},
		},
		"create-internet-gateway": {
			Name:   "create-internet-gateway",
			Fields: fields_create_internet_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInternetGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_internet_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInternetGateway(ctx, input)
			},
		},
		"create-interruptible-capacity-reservation-allocation": {
			Name:   "create-interruptible-capacity-reservation-allocation",
			Fields: fields_create_interruptible_capacity_reservation_allocation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInterruptibleCapacityReservationAllocationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_interruptible_capacity_reservation_allocation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInterruptibleCapacityReservationAllocation(ctx, input)
			},
		},
		"create-ipam": {
			Name:   "create-ipam",
			Fields: fields_create_ipam,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIpamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ipam, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIpam(ctx, input)
			},
		},
		"create-ipam-external-resource-verification-token": {
			Name:   "create-ipam-external-resource-verification-token",
			Fields: fields_create_ipam_external_resource_verification_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIpamExternalResourceVerificationTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ipam_external_resource_verification_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIpamExternalResourceVerificationToken(ctx, input)
			},
		},
		"create-ipam-policy": {
			Name:   "create-ipam-policy",
			Fields: fields_create_ipam_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIpamPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ipam_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIpamPolicy(ctx, input)
			},
		},
		"create-ipam-pool": {
			Name:   "create-ipam-pool",
			Fields: fields_create_ipam_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIpamPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ipam_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIpamPool(ctx, input)
			},
		},
		"create-ipam-prefix-list-resolver": {
			Name:   "create-ipam-prefix-list-resolver",
			Fields: fields_create_ipam_prefix_list_resolver,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIpamPrefixListResolverInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ipam_prefix_list_resolver, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIpamPrefixListResolver(ctx, input)
			},
		},
		"create-ipam-prefix-list-resolver-target": {
			Name:   "create-ipam-prefix-list-resolver-target",
			Fields: fields_create_ipam_prefix_list_resolver_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIpamPrefixListResolverTargetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ipam_prefix_list_resolver_target, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIpamPrefixListResolverTarget(ctx, input)
			},
		},
		"create-ipam-resource-discovery": {
			Name:   "create-ipam-resource-discovery",
			Fields: fields_create_ipam_resource_discovery,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIpamResourceDiscoveryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ipam_resource_discovery, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIpamResourceDiscovery(ctx, input)
			},
		},
		"create-ipam-scope": {
			Name:   "create-ipam-scope",
			Fields: fields_create_ipam_scope,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIpamScopeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ipam_scope, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIpamScope(ctx, input)
			},
		},
		"create-key-pair": {
			Name:   "create-key-pair",
			Fields: fields_create_key_pair,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateKeyPairInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_key_pair, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateKeyPair(ctx, input)
			},
		},
		"create-launch-template": {
			Name:   "create-launch-template",
			Fields: fields_create_launch_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLaunchTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_launch_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLaunchTemplate(ctx, input)
			},
		},
		"create-launch-template-version": {
			Name:   "create-launch-template-version",
			Fields: fields_create_launch_template_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLaunchTemplateVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_launch_template_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLaunchTemplateVersion(ctx, input)
			},
		},
		"create-local-gateway-route": {
			Name:   "create-local-gateway-route",
			Fields: fields_create_local_gateway_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLocalGatewayRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_local_gateway_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLocalGatewayRoute(ctx, input)
			},
		},
		"create-local-gateway-route-table": {
			Name:   "create-local-gateway-route-table",
			Fields: fields_create_local_gateway_route_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLocalGatewayRouteTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_local_gateway_route_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLocalGatewayRouteTable(ctx, input)
			},
		},
		"create-local-gateway-route-table-virtual-interface-group-association": {
			Name:   "create-local-gateway-route-table-virtual-interface-group-association",
			Fields: fields_create_local_gateway_route_table_virtual_interface_group_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_local_gateway_route_table_virtual_interface_group_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociation(ctx, input)
			},
		},
		"create-local-gateway-route-table-vpc-association": {
			Name:   "create-local-gateway-route-table-vpc-association",
			Fields: fields_create_local_gateway_route_table_vpc_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLocalGatewayRouteTableVpcAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_local_gateway_route_table_vpc_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLocalGatewayRouteTableVpcAssociation(ctx, input)
			},
		},
		"create-local-gateway-virtual-interface": {
			Name:   "create-local-gateway-virtual-interface",
			Fields: fields_create_local_gateway_virtual_interface,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLocalGatewayVirtualInterfaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_local_gateway_virtual_interface, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLocalGatewayVirtualInterface(ctx, input)
			},
		},
		"create-local-gateway-virtual-interface-group": {
			Name:   "create-local-gateway-virtual-interface-group",
			Fields: fields_create_local_gateway_virtual_interface_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLocalGatewayVirtualInterfaceGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_local_gateway_virtual_interface_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLocalGatewayVirtualInterfaceGroup(ctx, input)
			},
		},
		"create-mac-system-integrity-protection-modification-task": {
			Name:   "create-mac-system-integrity-protection-modification-task",
			Fields: fields_create_mac_system_integrity_protection_modification_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMacSystemIntegrityProtectionModificationTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_mac_system_integrity_protection_modification_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMacSystemIntegrityProtectionModificationTask(ctx, input)
			},
		},
		"create-managed-prefix-list": {
			Name:   "create-managed-prefix-list",
			Fields: fields_create_managed_prefix_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateManagedPrefixListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_managed_prefix_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateManagedPrefixList(ctx, input)
			},
		},
		"create-nat-gateway": {
			Name:   "create-nat-gateway",
			Fields: fields_create_nat_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNatGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_nat_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNatGateway(ctx, input)
			},
		},
		"create-network-acl": {
			Name:   "create-network-acl",
			Fields: fields_create_network_acl,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNetworkAclInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_network_acl, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNetworkAcl(ctx, input)
			},
		},
		"create-network-acl-entry": {
			Name:   "create-network-acl-entry",
			Fields: fields_create_network_acl_entry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNetworkAclEntryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_network_acl_entry, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNetworkAclEntry(ctx, input)
			},
		},
		"create-network-insights-access-scope": {
			Name:   "create-network-insights-access-scope",
			Fields: fields_create_network_insights_access_scope,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNetworkInsightsAccessScopeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_network_insights_access_scope, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNetworkInsightsAccessScope(ctx, input)
			},
		},
		"create-network-insights-path": {
			Name:   "create-network-insights-path",
			Fields: fields_create_network_insights_path,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNetworkInsightsPathInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_network_insights_path, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNetworkInsightsPath(ctx, input)
			},
		},
		"create-network-interface": {
			Name:   "create-network-interface",
			Fields: fields_create_network_interface,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNetworkInterfaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_network_interface, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNetworkInterface(ctx, input)
			},
		},
		"create-network-interface-permission": {
			Name:   "create-network-interface-permission",
			Fields: fields_create_network_interface_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNetworkInterfacePermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_network_interface_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNetworkInterfacePermission(ctx, input)
			},
		},
		"create-placement-group": {
			Name:   "create-placement-group",
			Fields: fields_create_placement_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePlacementGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_placement_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePlacementGroup(ctx, input)
			},
		},
		"create-public-ipv4-pool": {
			Name:   "create-public-ipv4-pool",
			Fields: fields_create_public_ipv4_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePublicIpv4PoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_public_ipv4_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePublicIpv4Pool(ctx, input)
			},
		},
		"create-replace-root-volume-task": {
			Name:   "create-replace-root-volume-task",
			Fields: fields_create_replace_root_volume_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateReplaceRootVolumeTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_replace_root_volume_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateReplaceRootVolumeTask(ctx, input)
			},
		},
		"create-reserved-instances-listing": {
			Name:   "create-reserved-instances-listing",
			Fields: fields_create_reserved_instances_listing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateReservedInstancesListingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_reserved_instances_listing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateReservedInstancesListing(ctx, input)
			},
		},
		"create-restore-image-task": {
			Name:   "create-restore-image-task",
			Fields: fields_create_restore_image_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRestoreImageTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_restore_image_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRestoreImageTask(ctx, input)
			},
		},
		"create-route": {
			Name:   "create-route",
			Fields: fields_create_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRoute(ctx, input)
			},
		},
		"create-route-server": {
			Name:   "create-route-server",
			Fields: fields_create_route_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRouteServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_route_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRouteServer(ctx, input)
			},
		},
		"create-route-server-endpoint": {
			Name:   "create-route-server-endpoint",
			Fields: fields_create_route_server_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRouteServerEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_route_server_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRouteServerEndpoint(ctx, input)
			},
		},
		"create-route-server-peer": {
			Name:   "create-route-server-peer",
			Fields: fields_create_route_server_peer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRouteServerPeerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_route_server_peer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRouteServerPeer(ctx, input)
			},
		},
		"create-route-table": {
			Name:   "create-route-table",
			Fields: fields_create_route_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRouteTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_route_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRouteTable(ctx, input)
			},
		},
		"create-secondary-network": {
			Name:   "create-secondary-network",
			Fields: fields_create_secondary_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSecondaryNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_secondary_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSecondaryNetwork(ctx, input)
			},
		},
		"create-secondary-subnet": {
			Name:   "create-secondary-subnet",
			Fields: fields_create_secondary_subnet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSecondarySubnetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_secondary_subnet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSecondarySubnet(ctx, input)
			},
		},
		"create-security-group": {
			Name:   "create-security-group",
			Fields: fields_create_security_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSecurityGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_security_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSecurityGroup(ctx, input)
			},
		},
		"create-snapshot": {
			Name:   "create-snapshot",
			Fields: fields_create_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSnapshot(ctx, input)
			},
		},
		"create-snapshots": {
			Name:   "create-snapshots",
			Fields: fields_create_snapshots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSnapshotsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_snapshots, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSnapshots(ctx, input)
			},
		},
		"create-spot-datafeed-subscription": {
			Name:   "create-spot-datafeed-subscription",
			Fields: fields_create_spot_datafeed_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSpotDatafeedSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_spot_datafeed_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSpotDatafeedSubscription(ctx, input)
			},
		},
		"create-store-image-task": {
			Name:   "create-store-image-task",
			Fields: fields_create_store_image_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStoreImageTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_store_image_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStoreImageTask(ctx, input)
			},
		},
		"create-subnet": {
			Name:   "create-subnet",
			Fields: fields_create_subnet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSubnetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_subnet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSubnet(ctx, input)
			},
		},
		"create-subnet-cidr-reservation": {
			Name:   "create-subnet-cidr-reservation",
			Fields: fields_create_subnet_cidr_reservation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSubnetCidrReservationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_subnet_cidr_reservation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSubnetCidrReservation(ctx, input)
			},
		},
		"create-tags": {
			Name:   "create-tags",
			Fields: fields_create_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTags(ctx, input)
			},
		},
		"create-traffic-mirror-filter": {
			Name:   "create-traffic-mirror-filter",
			Fields: fields_create_traffic_mirror_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTrafficMirrorFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_traffic_mirror_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTrafficMirrorFilter(ctx, input)
			},
		},
		"create-traffic-mirror-filter-rule": {
			Name:   "create-traffic-mirror-filter-rule",
			Fields: fields_create_traffic_mirror_filter_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTrafficMirrorFilterRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_traffic_mirror_filter_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTrafficMirrorFilterRule(ctx, input)
			},
		},
		"create-traffic-mirror-session": {
			Name:   "create-traffic-mirror-session",
			Fields: fields_create_traffic_mirror_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTrafficMirrorSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_traffic_mirror_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTrafficMirrorSession(ctx, input)
			},
		},
		"create-traffic-mirror-target": {
			Name:   "create-traffic-mirror-target",
			Fields: fields_create_traffic_mirror_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTrafficMirrorTargetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_traffic_mirror_target, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTrafficMirrorTarget(ctx, input)
			},
		},
		"create-transit-gateway": {
			Name:   "create-transit-gateway",
			Fields: fields_create_transit_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTransitGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_transit_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTransitGateway(ctx, input)
			},
		},
		"create-transit-gateway-connect": {
			Name:   "create-transit-gateway-connect",
			Fields: fields_create_transit_gateway_connect,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTransitGatewayConnectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_transit_gateway_connect, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTransitGatewayConnect(ctx, input)
			},
		},
		"create-transit-gateway-connect-peer": {
			Name:   "create-transit-gateway-connect-peer",
			Fields: fields_create_transit_gateway_connect_peer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTransitGatewayConnectPeerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_transit_gateway_connect_peer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTransitGatewayConnectPeer(ctx, input)
			},
		},
		"create-transit-gateway-metering-policy": {
			Name:   "create-transit-gateway-metering-policy",
			Fields: fields_create_transit_gateway_metering_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTransitGatewayMeteringPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_transit_gateway_metering_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTransitGatewayMeteringPolicy(ctx, input)
			},
		},
		"create-transit-gateway-metering-policy-entry": {
			Name:   "create-transit-gateway-metering-policy-entry",
			Fields: fields_create_transit_gateway_metering_policy_entry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTransitGatewayMeteringPolicyEntryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_transit_gateway_metering_policy_entry, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTransitGatewayMeteringPolicyEntry(ctx, input)
			},
		},
		"create-transit-gateway-multicast-domain": {
			Name:   "create-transit-gateway-multicast-domain",
			Fields: fields_create_transit_gateway_multicast_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTransitGatewayMulticastDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_transit_gateway_multicast_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTransitGatewayMulticastDomain(ctx, input)
			},
		},
		"create-transit-gateway-peering-attachment": {
			Name:   "create-transit-gateway-peering-attachment",
			Fields: fields_create_transit_gateway_peering_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTransitGatewayPeeringAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_transit_gateway_peering_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTransitGatewayPeeringAttachment(ctx, input)
			},
		},
		"create-transit-gateway-policy-table": {
			Name:   "create-transit-gateway-policy-table",
			Fields: fields_create_transit_gateway_policy_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTransitGatewayPolicyTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_transit_gateway_policy_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTransitGatewayPolicyTable(ctx, input)
			},
		},
		"create-transit-gateway-prefix-list-reference": {
			Name:   "create-transit-gateway-prefix-list-reference",
			Fields: fields_create_transit_gateway_prefix_list_reference,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTransitGatewayPrefixListReferenceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_transit_gateway_prefix_list_reference, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTransitGatewayPrefixListReference(ctx, input)
			},
		},
		"create-transit-gateway-route": {
			Name:   "create-transit-gateway-route",
			Fields: fields_create_transit_gateway_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTransitGatewayRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_transit_gateway_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTransitGatewayRoute(ctx, input)
			},
		},
		"create-transit-gateway-route-table": {
			Name:   "create-transit-gateway-route-table",
			Fields: fields_create_transit_gateway_route_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTransitGatewayRouteTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_transit_gateway_route_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTransitGatewayRouteTable(ctx, input)
			},
		},
		"create-transit-gateway-route-table-announcement": {
			Name:   "create-transit-gateway-route-table-announcement",
			Fields: fields_create_transit_gateway_route_table_announcement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTransitGatewayRouteTableAnnouncementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_transit_gateway_route_table_announcement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTransitGatewayRouteTableAnnouncement(ctx, input)
			},
		},
		"create-transit-gateway-vpc-attachment": {
			Name:   "create-transit-gateway-vpc-attachment",
			Fields: fields_create_transit_gateway_vpc_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTransitGatewayVpcAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_transit_gateway_vpc_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTransitGatewayVpcAttachment(ctx, input)
			},
		},
		"create-verified-access-endpoint": {
			Name:   "create-verified-access-endpoint",
			Fields: fields_create_verified_access_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVerifiedAccessEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_verified_access_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVerifiedAccessEndpoint(ctx, input)
			},
		},
		"create-verified-access-group": {
			Name:   "create-verified-access-group",
			Fields: fields_create_verified_access_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVerifiedAccessGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_verified_access_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVerifiedAccessGroup(ctx, input)
			},
		},
		"create-verified-access-instance": {
			Name:   "create-verified-access-instance",
			Fields: fields_create_verified_access_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVerifiedAccessInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_verified_access_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVerifiedAccessInstance(ctx, input)
			},
		},
		"create-verified-access-trust-provider": {
			Name:   "create-verified-access-trust-provider",
			Fields: fields_create_verified_access_trust_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVerifiedAccessTrustProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_verified_access_trust_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVerifiedAccessTrustProvider(ctx, input)
			},
		},
		"create-volume": {
			Name:   "create-volume",
			Fields: fields_create_volume,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVolumeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_volume, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVolume(ctx, input)
			},
		},
		"create-vpc": {
			Name:   "create-vpc",
			Fields: fields_create_vpc,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVpcInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vpc, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVpc(ctx, input)
			},
		},
		"create-vpc-block-public-access-exclusion": {
			Name:   "create-vpc-block-public-access-exclusion",
			Fields: fields_create_vpc_block_public_access_exclusion,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVpcBlockPublicAccessExclusionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vpc_block_public_access_exclusion, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVpcBlockPublicAccessExclusion(ctx, input)
			},
		},
		"create-vpc-encryption-control": {
			Name:   "create-vpc-encryption-control",
			Fields: fields_create_vpc_encryption_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVpcEncryptionControlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vpc_encryption_control, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVpcEncryptionControl(ctx, input)
			},
		},
		"create-vpc-endpoint": {
			Name:   "create-vpc-endpoint",
			Fields: fields_create_vpc_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVpcEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vpc_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVpcEndpoint(ctx, input)
			},
		},
		"create-vpc-endpoint-connection-notification": {
			Name:   "create-vpc-endpoint-connection-notification",
			Fields: fields_create_vpc_endpoint_connection_notification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVpcEndpointConnectionNotificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vpc_endpoint_connection_notification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVpcEndpointConnectionNotification(ctx, input)
			},
		},
		"create-vpc-endpoint-service-configuration": {
			Name:   "create-vpc-endpoint-service-configuration",
			Fields: fields_create_vpc_endpoint_service_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVpcEndpointServiceConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vpc_endpoint_service_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVpcEndpointServiceConfiguration(ctx, input)
			},
		},
		"create-vpc-peering-connection": {
			Name:   "create-vpc-peering-connection",
			Fields: fields_create_vpc_peering_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVpcPeeringConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vpc_peering_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVpcPeeringConnection(ctx, input)
			},
		},
		"create-vpn-concentrator": {
			Name:   "create-vpn-concentrator",
			Fields: fields_create_vpn_concentrator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVpnConcentratorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vpn_concentrator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVpnConcentrator(ctx, input)
			},
		},
		"create-vpn-connection": {
			Name:   "create-vpn-connection",
			Fields: fields_create_vpn_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVpnConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vpn_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVpnConnection(ctx, input)
			},
		},
		"create-vpn-connection-route": {
			Name:   "create-vpn-connection-route",
			Fields: fields_create_vpn_connection_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVpnConnectionRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vpn_connection_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVpnConnectionRoute(ctx, input)
			},
		},
		"create-vpn-gateway": {
			Name:   "create-vpn-gateway",
			Fields: fields_create_vpn_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVpnGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vpn_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVpnGateway(ctx, input)
			},
		},
		"delete-capacity-manager-data-export": {
			Name:   "delete-capacity-manager-data-export",
			Fields: fields_delete_capacity_manager_data_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCapacityManagerDataExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_capacity_manager_data_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCapacityManagerDataExport(ctx, input)
			},
		},
		"delete-carrier-gateway": {
			Name:   "delete-carrier-gateway",
			Fields: fields_delete_carrier_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCarrierGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_carrier_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCarrierGateway(ctx, input)
			},
		},
		"delete-client-vpn-endpoint": {
			Name:   "delete-client-vpn-endpoint",
			Fields: fields_delete_client_vpn_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteClientVpnEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_client_vpn_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteClientVpnEndpoint(ctx, input)
			},
		},
		"delete-client-vpn-route": {
			Name:   "delete-client-vpn-route",
			Fields: fields_delete_client_vpn_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteClientVpnRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_client_vpn_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteClientVpnRoute(ctx, input)
			},
		},
		"delete-coip-cidr": {
			Name:   "delete-coip-cidr",
			Fields: fields_delete_coip_cidr,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCoipCidrInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_coip_cidr, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCoipCidr(ctx, input)
			},
		},
		"delete-coip-pool": {
			Name:   "delete-coip-pool",
			Fields: fields_delete_coip_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCoipPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_coip_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCoipPool(ctx, input)
			},
		},
		"delete-customer-gateway": {
			Name:   "delete-customer-gateway",
			Fields: fields_delete_customer_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCustomerGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_customer_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCustomerGateway(ctx, input)
			},
		},
		"delete-dhcp-options": {
			Name:   "delete-dhcp-options",
			Fields: fields_delete_dhcp_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDhcpOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_dhcp_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDhcpOptions(ctx, input)
			},
		},
		"delete-egress-only-internet-gateway": {
			Name:   "delete-egress-only-internet-gateway",
			Fields: fields_delete_egress_only_internet_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEgressOnlyInternetGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_egress_only_internet_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEgressOnlyInternetGateway(ctx, input)
			},
		},
		"delete-fleets": {
			Name:   "delete-fleets",
			Fields: fields_delete_fleets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFleetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_fleets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFleets(ctx, input)
			},
		},
		"delete-flow-logs": {
			Name:   "delete-flow-logs",
			Fields: fields_delete_flow_logs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFlowLogsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_flow_logs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFlowLogs(ctx, input)
			},
		},
		"delete-fpga-image": {
			Name:   "delete-fpga-image",
			Fields: fields_delete_fpga_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFpgaImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_fpga_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFpgaImage(ctx, input)
			},
		},
		"delete-image-usage-report": {
			Name:   "delete-image-usage-report",
			Fields: fields_delete_image_usage_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteImageUsageReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_image_usage_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteImageUsageReport(ctx, input)
			},
		},
		"delete-instance-connect-endpoint": {
			Name:   "delete-instance-connect-endpoint",
			Fields: fields_delete_instance_connect_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInstanceConnectEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_instance_connect_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInstanceConnectEndpoint(ctx, input)
			},
		},
		"delete-instance-event-window": {
			Name:   "delete-instance-event-window",
			Fields: fields_delete_instance_event_window,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInstanceEventWindowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_instance_event_window, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInstanceEventWindow(ctx, input)
			},
		},
		"delete-internet-gateway": {
			Name:   "delete-internet-gateway",
			Fields: fields_delete_internet_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInternetGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_internet_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInternetGateway(ctx, input)
			},
		},
		"delete-ipam": {
			Name:   "delete-ipam",
			Fields: fields_delete_ipam,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIpamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ipam, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIpam(ctx, input)
			},
		},
		"delete-ipam-external-resource-verification-token": {
			Name:   "delete-ipam-external-resource-verification-token",
			Fields: fields_delete_ipam_external_resource_verification_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIpamExternalResourceVerificationTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ipam_external_resource_verification_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIpamExternalResourceVerificationToken(ctx, input)
			},
		},
		"delete-ipam-policy": {
			Name:   "delete-ipam-policy",
			Fields: fields_delete_ipam_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIpamPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ipam_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIpamPolicy(ctx, input)
			},
		},
		"delete-ipam-pool": {
			Name:   "delete-ipam-pool",
			Fields: fields_delete_ipam_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIpamPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ipam_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIpamPool(ctx, input)
			},
		},
		"delete-ipam-prefix-list-resolver": {
			Name:   "delete-ipam-prefix-list-resolver",
			Fields: fields_delete_ipam_prefix_list_resolver,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIpamPrefixListResolverInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ipam_prefix_list_resolver, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIpamPrefixListResolver(ctx, input)
			},
		},
		"delete-ipam-prefix-list-resolver-target": {
			Name:   "delete-ipam-prefix-list-resolver-target",
			Fields: fields_delete_ipam_prefix_list_resolver_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIpamPrefixListResolverTargetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ipam_prefix_list_resolver_target, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIpamPrefixListResolverTarget(ctx, input)
			},
		},
		"delete-ipam-resource-discovery": {
			Name:   "delete-ipam-resource-discovery",
			Fields: fields_delete_ipam_resource_discovery,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIpamResourceDiscoveryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ipam_resource_discovery, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIpamResourceDiscovery(ctx, input)
			},
		},
		"delete-ipam-scope": {
			Name:   "delete-ipam-scope",
			Fields: fields_delete_ipam_scope,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIpamScopeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ipam_scope, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIpamScope(ctx, input)
			},
		},
		"delete-key-pair": {
			Name:   "delete-key-pair",
			Fields: fields_delete_key_pair,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteKeyPairInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_key_pair, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteKeyPair(ctx, input)
			},
		},
		"delete-launch-template": {
			Name:   "delete-launch-template",
			Fields: fields_delete_launch_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLaunchTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_launch_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLaunchTemplate(ctx, input)
			},
		},
		"delete-launch-template-versions": {
			Name:   "delete-launch-template-versions",
			Fields: fields_delete_launch_template_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLaunchTemplateVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_launch_template_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLaunchTemplateVersions(ctx, input)
			},
		},
		"delete-local-gateway-route": {
			Name:   "delete-local-gateway-route",
			Fields: fields_delete_local_gateway_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLocalGatewayRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_local_gateway_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLocalGatewayRoute(ctx, input)
			},
		},
		"delete-local-gateway-route-table": {
			Name:   "delete-local-gateway-route-table",
			Fields: fields_delete_local_gateway_route_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLocalGatewayRouteTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_local_gateway_route_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLocalGatewayRouteTable(ctx, input)
			},
		},
		"delete-local-gateway-route-table-virtual-interface-group-association": {
			Name:   "delete-local-gateway-route-table-virtual-interface-group-association",
			Fields: fields_delete_local_gateway_route_table_virtual_interface_group_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_local_gateway_route_table_virtual_interface_group_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociation(ctx, input)
			},
		},
		"delete-local-gateway-route-table-vpc-association": {
			Name:   "delete-local-gateway-route-table-vpc-association",
			Fields: fields_delete_local_gateway_route_table_vpc_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLocalGatewayRouteTableVpcAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_local_gateway_route_table_vpc_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLocalGatewayRouteTableVpcAssociation(ctx, input)
			},
		},
		"delete-local-gateway-virtual-interface": {
			Name:   "delete-local-gateway-virtual-interface",
			Fields: fields_delete_local_gateway_virtual_interface,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLocalGatewayVirtualInterfaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_local_gateway_virtual_interface, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLocalGatewayVirtualInterface(ctx, input)
			},
		},
		"delete-local-gateway-virtual-interface-group": {
			Name:   "delete-local-gateway-virtual-interface-group",
			Fields: fields_delete_local_gateway_virtual_interface_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLocalGatewayVirtualInterfaceGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_local_gateway_virtual_interface_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLocalGatewayVirtualInterfaceGroup(ctx, input)
			},
		},
		"delete-managed-prefix-list": {
			Name:   "delete-managed-prefix-list",
			Fields: fields_delete_managed_prefix_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteManagedPrefixListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_managed_prefix_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteManagedPrefixList(ctx, input)
			},
		},
		"delete-nat-gateway": {
			Name:   "delete-nat-gateway",
			Fields: fields_delete_nat_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNatGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_nat_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNatGateway(ctx, input)
			},
		},
		"delete-network-acl": {
			Name:   "delete-network-acl",
			Fields: fields_delete_network_acl,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNetworkAclInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_network_acl, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNetworkAcl(ctx, input)
			},
		},
		"delete-network-acl-entry": {
			Name:   "delete-network-acl-entry",
			Fields: fields_delete_network_acl_entry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNetworkAclEntryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_network_acl_entry, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNetworkAclEntry(ctx, input)
			},
		},
		"delete-network-insights-access-scope": {
			Name:   "delete-network-insights-access-scope",
			Fields: fields_delete_network_insights_access_scope,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNetworkInsightsAccessScopeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_network_insights_access_scope, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNetworkInsightsAccessScope(ctx, input)
			},
		},
		"delete-network-insights-access-scope-analysis": {
			Name:   "delete-network-insights-access-scope-analysis",
			Fields: fields_delete_network_insights_access_scope_analysis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNetworkInsightsAccessScopeAnalysisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_network_insights_access_scope_analysis, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNetworkInsightsAccessScopeAnalysis(ctx, input)
			},
		},
		"delete-network-insights-analysis": {
			Name:   "delete-network-insights-analysis",
			Fields: fields_delete_network_insights_analysis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNetworkInsightsAnalysisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_network_insights_analysis, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNetworkInsightsAnalysis(ctx, input)
			},
		},
		"delete-network-insights-path": {
			Name:   "delete-network-insights-path",
			Fields: fields_delete_network_insights_path,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNetworkInsightsPathInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_network_insights_path, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNetworkInsightsPath(ctx, input)
			},
		},
		"delete-network-interface": {
			Name:   "delete-network-interface",
			Fields: fields_delete_network_interface,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNetworkInterfaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_network_interface, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNetworkInterface(ctx, input)
			},
		},
		"delete-network-interface-permission": {
			Name:   "delete-network-interface-permission",
			Fields: fields_delete_network_interface_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNetworkInterfacePermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_network_interface_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNetworkInterfacePermission(ctx, input)
			},
		},
		"delete-placement-group": {
			Name:   "delete-placement-group",
			Fields: fields_delete_placement_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePlacementGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_placement_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePlacementGroup(ctx, input)
			},
		},
		"delete-public-ipv4-pool": {
			Name:   "delete-public-ipv4-pool",
			Fields: fields_delete_public_ipv4_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePublicIpv4PoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_public_ipv4_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePublicIpv4Pool(ctx, input)
			},
		},
		"delete-queued-reserved-instances": {
			Name:   "delete-queued-reserved-instances",
			Fields: fields_delete_queued_reserved_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteQueuedReservedInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_queued_reserved_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteQueuedReservedInstances(ctx, input)
			},
		},
		"delete-route": {
			Name:   "delete-route",
			Fields: fields_delete_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRoute(ctx, input)
			},
		},
		"delete-route-server": {
			Name:   "delete-route-server",
			Fields: fields_delete_route_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRouteServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_route_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRouteServer(ctx, input)
			},
		},
		"delete-route-server-endpoint": {
			Name:   "delete-route-server-endpoint",
			Fields: fields_delete_route_server_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRouteServerEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_route_server_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRouteServerEndpoint(ctx, input)
			},
		},
		"delete-route-server-peer": {
			Name:   "delete-route-server-peer",
			Fields: fields_delete_route_server_peer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRouteServerPeerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_route_server_peer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRouteServerPeer(ctx, input)
			},
		},
		"delete-route-table": {
			Name:   "delete-route-table",
			Fields: fields_delete_route_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRouteTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_route_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRouteTable(ctx, input)
			},
		},
		"delete-secondary-network": {
			Name:   "delete-secondary-network",
			Fields: fields_delete_secondary_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSecondaryNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_secondary_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSecondaryNetwork(ctx, input)
			},
		},
		"delete-secondary-subnet": {
			Name:   "delete-secondary-subnet",
			Fields: fields_delete_secondary_subnet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSecondarySubnetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_secondary_subnet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSecondarySubnet(ctx, input)
			},
		},
		"delete-security-group": {
			Name:   "delete-security-group",
			Fields: fields_delete_security_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSecurityGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_security_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSecurityGroup(ctx, input)
			},
		},
		"delete-snapshot": {
			Name:   "delete-snapshot",
			Fields: fields_delete_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSnapshot(ctx, input)
			},
		},
		"delete-spot-datafeed-subscription": {
			Name:   "delete-spot-datafeed-subscription",
			Fields: fields_delete_spot_datafeed_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSpotDatafeedSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_spot_datafeed_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSpotDatafeedSubscription(ctx, input)
			},
		},
		"delete-subnet": {
			Name:   "delete-subnet",
			Fields: fields_delete_subnet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSubnetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_subnet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSubnet(ctx, input)
			},
		},
		"delete-subnet-cidr-reservation": {
			Name:   "delete-subnet-cidr-reservation",
			Fields: fields_delete_subnet_cidr_reservation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSubnetCidrReservationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_subnet_cidr_reservation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSubnetCidrReservation(ctx, input)
			},
		},
		"delete-tags": {
			Name:   "delete-tags",
			Fields: fields_delete_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTags(ctx, input)
			},
		},
		"delete-traffic-mirror-filter": {
			Name:   "delete-traffic-mirror-filter",
			Fields: fields_delete_traffic_mirror_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTrafficMirrorFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_traffic_mirror_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTrafficMirrorFilter(ctx, input)
			},
		},
		"delete-traffic-mirror-filter-rule": {
			Name:   "delete-traffic-mirror-filter-rule",
			Fields: fields_delete_traffic_mirror_filter_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTrafficMirrorFilterRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_traffic_mirror_filter_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTrafficMirrorFilterRule(ctx, input)
			},
		},
		"delete-traffic-mirror-session": {
			Name:   "delete-traffic-mirror-session",
			Fields: fields_delete_traffic_mirror_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTrafficMirrorSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_traffic_mirror_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTrafficMirrorSession(ctx, input)
			},
		},
		"delete-traffic-mirror-target": {
			Name:   "delete-traffic-mirror-target",
			Fields: fields_delete_traffic_mirror_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTrafficMirrorTargetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_traffic_mirror_target, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTrafficMirrorTarget(ctx, input)
			},
		},
		"delete-transit-gateway": {
			Name:   "delete-transit-gateway",
			Fields: fields_delete_transit_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTransitGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_transit_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTransitGateway(ctx, input)
			},
		},
		"delete-transit-gateway-connect": {
			Name:   "delete-transit-gateway-connect",
			Fields: fields_delete_transit_gateway_connect,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTransitGatewayConnectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_transit_gateway_connect, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTransitGatewayConnect(ctx, input)
			},
		},
		"delete-transit-gateway-connect-peer": {
			Name:   "delete-transit-gateway-connect-peer",
			Fields: fields_delete_transit_gateway_connect_peer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTransitGatewayConnectPeerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_transit_gateway_connect_peer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTransitGatewayConnectPeer(ctx, input)
			},
		},
		"delete-transit-gateway-metering-policy": {
			Name:   "delete-transit-gateway-metering-policy",
			Fields: fields_delete_transit_gateway_metering_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTransitGatewayMeteringPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_transit_gateway_metering_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTransitGatewayMeteringPolicy(ctx, input)
			},
		},
		"delete-transit-gateway-metering-policy-entry": {
			Name:   "delete-transit-gateway-metering-policy-entry",
			Fields: fields_delete_transit_gateway_metering_policy_entry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTransitGatewayMeteringPolicyEntryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_transit_gateway_metering_policy_entry, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTransitGatewayMeteringPolicyEntry(ctx, input)
			},
		},
		"delete-transit-gateway-multicast-domain": {
			Name:   "delete-transit-gateway-multicast-domain",
			Fields: fields_delete_transit_gateway_multicast_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTransitGatewayMulticastDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_transit_gateway_multicast_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTransitGatewayMulticastDomain(ctx, input)
			},
		},
		"delete-transit-gateway-peering-attachment": {
			Name:   "delete-transit-gateway-peering-attachment",
			Fields: fields_delete_transit_gateway_peering_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTransitGatewayPeeringAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_transit_gateway_peering_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTransitGatewayPeeringAttachment(ctx, input)
			},
		},
		"delete-transit-gateway-policy-table": {
			Name:   "delete-transit-gateway-policy-table",
			Fields: fields_delete_transit_gateway_policy_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTransitGatewayPolicyTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_transit_gateway_policy_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTransitGatewayPolicyTable(ctx, input)
			},
		},
		"delete-transit-gateway-prefix-list-reference": {
			Name:   "delete-transit-gateway-prefix-list-reference",
			Fields: fields_delete_transit_gateway_prefix_list_reference,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTransitGatewayPrefixListReferenceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_transit_gateway_prefix_list_reference, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTransitGatewayPrefixListReference(ctx, input)
			},
		},
		"delete-transit-gateway-route": {
			Name:   "delete-transit-gateway-route",
			Fields: fields_delete_transit_gateway_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTransitGatewayRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_transit_gateway_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTransitGatewayRoute(ctx, input)
			},
		},
		"delete-transit-gateway-route-table": {
			Name:   "delete-transit-gateway-route-table",
			Fields: fields_delete_transit_gateway_route_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTransitGatewayRouteTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_transit_gateway_route_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTransitGatewayRouteTable(ctx, input)
			},
		},
		"delete-transit-gateway-route-table-announcement": {
			Name:   "delete-transit-gateway-route-table-announcement",
			Fields: fields_delete_transit_gateway_route_table_announcement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTransitGatewayRouteTableAnnouncementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_transit_gateway_route_table_announcement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTransitGatewayRouteTableAnnouncement(ctx, input)
			},
		},
		"delete-transit-gateway-vpc-attachment": {
			Name:   "delete-transit-gateway-vpc-attachment",
			Fields: fields_delete_transit_gateway_vpc_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTransitGatewayVpcAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_transit_gateway_vpc_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTransitGatewayVpcAttachment(ctx, input)
			},
		},
		"delete-verified-access-endpoint": {
			Name:   "delete-verified-access-endpoint",
			Fields: fields_delete_verified_access_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVerifiedAccessEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_verified_access_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVerifiedAccessEndpoint(ctx, input)
			},
		},
		"delete-verified-access-group": {
			Name:   "delete-verified-access-group",
			Fields: fields_delete_verified_access_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVerifiedAccessGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_verified_access_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVerifiedAccessGroup(ctx, input)
			},
		},
		"delete-verified-access-instance": {
			Name:   "delete-verified-access-instance",
			Fields: fields_delete_verified_access_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVerifiedAccessInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_verified_access_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVerifiedAccessInstance(ctx, input)
			},
		},
		"delete-verified-access-trust-provider": {
			Name:   "delete-verified-access-trust-provider",
			Fields: fields_delete_verified_access_trust_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVerifiedAccessTrustProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_verified_access_trust_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVerifiedAccessTrustProvider(ctx, input)
			},
		},
		"delete-volume": {
			Name:   "delete-volume",
			Fields: fields_delete_volume,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVolumeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_volume, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVolume(ctx, input)
			},
		},
		"delete-vpc": {
			Name:   "delete-vpc",
			Fields: fields_delete_vpc,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVpcInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vpc, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVpc(ctx, input)
			},
		},
		"delete-vpc-block-public-access-exclusion": {
			Name:   "delete-vpc-block-public-access-exclusion",
			Fields: fields_delete_vpc_block_public_access_exclusion,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVpcBlockPublicAccessExclusionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vpc_block_public_access_exclusion, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVpcBlockPublicAccessExclusion(ctx, input)
			},
		},
		"delete-vpc-encryption-control": {
			Name:   "delete-vpc-encryption-control",
			Fields: fields_delete_vpc_encryption_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVpcEncryptionControlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vpc_encryption_control, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVpcEncryptionControl(ctx, input)
			},
		},
		"delete-vpc-endpoint-connection-notifications": {
			Name:   "delete-vpc-endpoint-connection-notifications",
			Fields: fields_delete_vpc_endpoint_connection_notifications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVpcEndpointConnectionNotificationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vpc_endpoint_connection_notifications, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVpcEndpointConnectionNotifications(ctx, input)
			},
		},
		"delete-vpc-endpoint-service-configurations": {
			Name:   "delete-vpc-endpoint-service-configurations",
			Fields: fields_delete_vpc_endpoint_service_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVpcEndpointServiceConfigurationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vpc_endpoint_service_configurations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVpcEndpointServiceConfigurations(ctx, input)
			},
		},
		"delete-vpc-endpoints": {
			Name:   "delete-vpc-endpoints",
			Fields: fields_delete_vpc_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVpcEndpointsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vpc_endpoints, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVpcEndpoints(ctx, input)
			},
		},
		"delete-vpc-peering-connection": {
			Name:   "delete-vpc-peering-connection",
			Fields: fields_delete_vpc_peering_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVpcPeeringConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vpc_peering_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVpcPeeringConnection(ctx, input)
			},
		},
		"delete-vpn-concentrator": {
			Name:   "delete-vpn-concentrator",
			Fields: fields_delete_vpn_concentrator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVpnConcentratorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vpn_concentrator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVpnConcentrator(ctx, input)
			},
		},
		"delete-vpn-connection": {
			Name:   "delete-vpn-connection",
			Fields: fields_delete_vpn_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVpnConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vpn_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVpnConnection(ctx, input)
			},
		},
		"delete-vpn-connection-route": {
			Name:   "delete-vpn-connection-route",
			Fields: fields_delete_vpn_connection_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVpnConnectionRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vpn_connection_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVpnConnectionRoute(ctx, input)
			},
		},
		"delete-vpn-gateway": {
			Name:   "delete-vpn-gateway",
			Fields: fields_delete_vpn_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVpnGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vpn_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVpnGateway(ctx, input)
			},
		},
		"deprovision-byoip-cidr": {
			Name:   "deprovision-byoip-cidr",
			Fields: fields_deprovision_byoip_cidr,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeprovisionByoipCidrInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deprovision_byoip_cidr, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeprovisionByoipCidr(ctx, input)
			},
		},
		"deprovision-ipam-byoasn": {
			Name:   "deprovision-ipam-byoasn",
			Fields: fields_deprovision_ipam_byoasn,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeprovisionIpamByoasnInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deprovision_ipam_byoasn, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeprovisionIpamByoasn(ctx, input)
			},
		},
		"deprovision-ipam-pool-cidr": {
			Name:   "deprovision-ipam-pool-cidr",
			Fields: fields_deprovision_ipam_pool_cidr,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeprovisionIpamPoolCidrInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deprovision_ipam_pool_cidr, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeprovisionIpamPoolCidr(ctx, input)
			},
		},
		"deprovision-public-ipv4-pool-cidr": {
			Name:   "deprovision-public-ipv4-pool-cidr",
			Fields: fields_deprovision_public_ipv4_pool_cidr,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeprovisionPublicIpv4PoolCidrInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deprovision_public_ipv4_pool_cidr, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeprovisionPublicIpv4PoolCidr(ctx, input)
			},
		},
		"deregister-image": {
			Name:   "deregister-image",
			Fields: fields_deregister_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterImage(ctx, input)
			},
		},
		"deregister-instance-event-notification-attributes": {
			Name:   "deregister-instance-event-notification-attributes",
			Fields: fields_deregister_instance_event_notification_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterInstanceEventNotificationAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_instance_event_notification_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterInstanceEventNotificationAttributes(ctx, input)
			},
		},
		"deregister-transit-gateway-multicast-group-members": {
			Name:   "deregister-transit-gateway-multicast-group-members",
			Fields: fields_deregister_transit_gateway_multicast_group_members,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterTransitGatewayMulticastGroupMembersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_transit_gateway_multicast_group_members, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterTransitGatewayMulticastGroupMembers(ctx, input)
			},
		},
		"deregister-transit-gateway-multicast-group-sources": {
			Name:   "deregister-transit-gateway-multicast-group-sources",
			Fields: fields_deregister_transit_gateway_multicast_group_sources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterTransitGatewayMulticastGroupSourcesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_transit_gateway_multicast_group_sources, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterTransitGatewayMulticastGroupSources(ctx, input)
			},
		},
		"describe-account-attributes": {
			Name:   "describe-account-attributes",
			Fields: fields_describe_account_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccountAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_account_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAccountAttributes(ctx, input)
			},
		},
		"describe-address-transfers": {
			Name:   "describe-address-transfers",
			Fields: fields_describe_address_transfers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAddressTransfersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_address_transfers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAddressTransfers(ctx, input)
				}
				var results []*svc.DescribeAddressTransfersOutput
				p := svc.NewDescribeAddressTransfersPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-addresses": {
			Name:   "describe-addresses",
			Fields: fields_describe_addresses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAddressesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_addresses, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAddresses(ctx, input)
			},
		},
		"describe-addresses-attribute": {
			Name:   "describe-addresses-attribute",
			Fields: fields_describe_addresses_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAddressesAttributeInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_addresses_attribute, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAddressesAttribute(ctx, input)
				}
				var results []*svc.DescribeAddressesAttributeOutput
				p := svc.NewDescribeAddressesAttributePaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-aggregate-id-format": {
			Name:   "describe-aggregate-id-format",
			Fields: fields_describe_aggregate_id_format,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAggregateIdFormatInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_aggregate_id_format, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAggregateIdFormat(ctx, input)
			},
		},
		"describe-availability-zones": {
			Name:   "describe-availability-zones",
			Fields: fields_describe_availability_zones,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAvailabilityZonesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_availability_zones, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAvailabilityZones(ctx, input)
			},
		},
		"describe-aws-network-performance-metric-subscriptions": {
			Name:   "describe-aws-network-performance-metric-subscriptions",
			Fields: fields_describe_aws_network_performance_metric_subscriptions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAwsNetworkPerformanceMetricSubscriptionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_aws_network_performance_metric_subscriptions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAwsNetworkPerformanceMetricSubscriptions(ctx, input)
				}
				var results []*svc.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput
				p := svc.NewDescribeAwsNetworkPerformanceMetricSubscriptionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-bundle-tasks": {
			Name:   "describe-bundle-tasks",
			Fields: fields_describe_bundle_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBundleTasksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_bundle_tasks, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBundleTasks(ctx, input)
			},
		},
		"describe-byoip-cidrs": {
			Name:   "describe-byoip-cidrs",
			Fields: fields_describe_byoip_cidrs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeByoipCidrsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_byoip_cidrs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeByoipCidrs(ctx, input)
				}
				var results []*svc.DescribeByoipCidrsOutput
				p := svc.NewDescribeByoipCidrsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-capacity-block-extension-history": {
			Name:   "describe-capacity-block-extension-history",
			Fields: fields_describe_capacity_block_extension_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCapacityBlockExtensionHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_capacity_block_extension_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeCapacityBlockExtensionHistory(ctx, input)
				}
				var results []*svc.DescribeCapacityBlockExtensionHistoryOutput
				p := svc.NewDescribeCapacityBlockExtensionHistoryPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-capacity-block-extension-offerings": {
			Name:   "describe-capacity-block-extension-offerings",
			Fields: fields_describe_capacity_block_extension_offerings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCapacityBlockExtensionOfferingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_capacity_block_extension_offerings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeCapacityBlockExtensionOfferings(ctx, input)
				}
				var results []*svc.DescribeCapacityBlockExtensionOfferingsOutput
				p := svc.NewDescribeCapacityBlockExtensionOfferingsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-capacity-block-offerings": {
			Name:   "describe-capacity-block-offerings",
			Fields: fields_describe_capacity_block_offerings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCapacityBlockOfferingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_capacity_block_offerings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeCapacityBlockOfferings(ctx, input)
				}
				var results []*svc.DescribeCapacityBlockOfferingsOutput
				p := svc.NewDescribeCapacityBlockOfferingsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-capacity-block-status": {
			Name:   "describe-capacity-block-status",
			Fields: fields_describe_capacity_block_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCapacityBlockStatusInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_capacity_block_status, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeCapacityBlockStatus(ctx, input)
				}
				var results []*svc.DescribeCapacityBlockStatusOutput
				p := svc.NewDescribeCapacityBlockStatusPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-capacity-blocks": {
			Name:   "describe-capacity-blocks",
			Fields: fields_describe_capacity_blocks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCapacityBlocksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_capacity_blocks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeCapacityBlocks(ctx, input)
				}
				var results []*svc.DescribeCapacityBlocksOutput
				p := svc.NewDescribeCapacityBlocksPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-capacity-manager-data-exports": {
			Name:   "describe-capacity-manager-data-exports",
			Fields: fields_describe_capacity_manager_data_exports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCapacityManagerDataExportsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_capacity_manager_data_exports, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeCapacityManagerDataExports(ctx, input)
				}
				var results []*svc.DescribeCapacityManagerDataExportsOutput
				p := svc.NewDescribeCapacityManagerDataExportsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-capacity-reservation-billing-requests": {
			Name:   "describe-capacity-reservation-billing-requests",
			Fields: fields_describe_capacity_reservation_billing_requests,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCapacityReservationBillingRequestsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_capacity_reservation_billing_requests, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeCapacityReservationBillingRequests(ctx, input)
				}
				var results []*svc.DescribeCapacityReservationBillingRequestsOutput
				p := svc.NewDescribeCapacityReservationBillingRequestsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-capacity-reservation-fleets": {
			Name:   "describe-capacity-reservation-fleets",
			Fields: fields_describe_capacity_reservation_fleets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCapacityReservationFleetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_capacity_reservation_fleets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeCapacityReservationFleets(ctx, input)
				}
				var results []*svc.DescribeCapacityReservationFleetsOutput
				p := svc.NewDescribeCapacityReservationFleetsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-capacity-reservation-topology": {
			Name:   "describe-capacity-reservation-topology",
			Fields: fields_describe_capacity_reservation_topology,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCapacityReservationTopologyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_capacity_reservation_topology, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCapacityReservationTopology(ctx, input)
			},
		},
		"describe-capacity-reservations": {
			Name:   "describe-capacity-reservations",
			Fields: fields_describe_capacity_reservations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCapacityReservationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_capacity_reservations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeCapacityReservations(ctx, input)
				}
				var results []*svc.DescribeCapacityReservationsOutput
				p := svc.NewDescribeCapacityReservationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-carrier-gateways": {
			Name:   "describe-carrier-gateways",
			Fields: fields_describe_carrier_gateways,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCarrierGatewaysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_carrier_gateways, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeCarrierGateways(ctx, input)
				}
				var results []*svc.DescribeCarrierGatewaysOutput
				p := svc.NewDescribeCarrierGatewaysPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-classic-link-instances": {
			Name:   "describe-classic-link-instances",
			Fields: fields_describe_classic_link_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClassicLinkInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_classic_link_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeClassicLinkInstances(ctx, input)
				}
				var results []*svc.DescribeClassicLinkInstancesOutput
				p := svc.NewDescribeClassicLinkInstancesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-client-vpn-authorization-rules": {
			Name:   "describe-client-vpn-authorization-rules",
			Fields: fields_describe_client_vpn_authorization_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClientVpnAuthorizationRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_client_vpn_authorization_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeClientVpnAuthorizationRules(ctx, input)
				}
				var results []*svc.DescribeClientVpnAuthorizationRulesOutput
				p := svc.NewDescribeClientVpnAuthorizationRulesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-client-vpn-connections": {
			Name:   "describe-client-vpn-connections",
			Fields: fields_describe_client_vpn_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClientVpnConnectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_client_vpn_connections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeClientVpnConnections(ctx, input)
				}
				var results []*svc.DescribeClientVpnConnectionsOutput
				p := svc.NewDescribeClientVpnConnectionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-client-vpn-endpoints": {
			Name:   "describe-client-vpn-endpoints",
			Fields: fields_describe_client_vpn_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClientVpnEndpointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_client_vpn_endpoints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeClientVpnEndpoints(ctx, input)
				}
				var results []*svc.DescribeClientVpnEndpointsOutput
				p := svc.NewDescribeClientVpnEndpointsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-client-vpn-routes": {
			Name:   "describe-client-vpn-routes",
			Fields: fields_describe_client_vpn_routes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClientVpnRoutesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_client_vpn_routes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeClientVpnRoutes(ctx, input)
				}
				var results []*svc.DescribeClientVpnRoutesOutput
				p := svc.NewDescribeClientVpnRoutesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-client-vpn-target-networks": {
			Name:   "describe-client-vpn-target-networks",
			Fields: fields_describe_client_vpn_target_networks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClientVpnTargetNetworksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_client_vpn_target_networks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeClientVpnTargetNetworks(ctx, input)
				}
				var results []*svc.DescribeClientVpnTargetNetworksOutput
				p := svc.NewDescribeClientVpnTargetNetworksPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-coip-pools": {
			Name:   "describe-coip-pools",
			Fields: fields_describe_coip_pools,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCoipPoolsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_coip_pools, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeCoipPools(ctx, input)
				}
				var results []*svc.DescribeCoipPoolsOutput
				p := svc.NewDescribeCoipPoolsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-conversion-tasks": {
			Name:   "describe-conversion-tasks",
			Fields: fields_describe_conversion_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConversionTasksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_conversion_tasks, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeConversionTasks(ctx, input)
			},
		},
		"describe-customer-gateways": {
			Name:   "describe-customer-gateways",
			Fields: fields_describe_customer_gateways,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCustomerGatewaysInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_customer_gateways, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCustomerGateways(ctx, input)
			},
		},
		"describe-declarative-policies-reports": {
			Name:   "describe-declarative-policies-reports",
			Fields: fields_describe_declarative_policies_reports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDeclarativePoliciesReportsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_declarative_policies_reports, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDeclarativePoliciesReports(ctx, input)
			},
		},
		"describe-dhcp-options": {
			Name:   "describe-dhcp-options",
			Fields: fields_describe_dhcp_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDhcpOptionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_dhcp_options, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDhcpOptions(ctx, input)
				}
				var results []*svc.DescribeDhcpOptionsOutput
				p := svc.NewDescribeDhcpOptionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-egress-only-internet-gateways": {
			Name:   "describe-egress-only-internet-gateways",
			Fields: fields_describe_egress_only_internet_gateways,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEgressOnlyInternetGatewaysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_egress_only_internet_gateways, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeEgressOnlyInternetGateways(ctx, input)
				}
				var results []*svc.DescribeEgressOnlyInternetGatewaysOutput
				p := svc.NewDescribeEgressOnlyInternetGatewaysPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-elastic-gpus": {
			Name:   "describe-elastic-gpus",
			Fields: fields_describe_elastic_gpus,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeElasticGpusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_elastic_gpus, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeElasticGpus(ctx, input)
			},
		},
		"describe-export-image-tasks": {
			Name:   "describe-export-image-tasks",
			Fields: fields_describe_export_image_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeExportImageTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_export_image_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeExportImageTasks(ctx, input)
				}
				var results []*svc.DescribeExportImageTasksOutput
				p := svc.NewDescribeExportImageTasksPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-export-tasks": {
			Name:   "describe-export-tasks",
			Fields: fields_describe_export_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeExportTasksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_export_tasks, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeExportTasks(ctx, input)
			},
		},
		"describe-fast-launch-images": {
			Name:   "describe-fast-launch-images",
			Fields: fields_describe_fast_launch_images,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFastLaunchImagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_fast_launch_images, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeFastLaunchImages(ctx, input)
				}
				var results []*svc.DescribeFastLaunchImagesOutput
				p := svc.NewDescribeFastLaunchImagesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-fast-snapshot-restores": {
			Name:   "describe-fast-snapshot-restores",
			Fields: fields_describe_fast_snapshot_restores,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFastSnapshotRestoresInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_fast_snapshot_restores, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeFastSnapshotRestores(ctx, input)
				}
				var results []*svc.DescribeFastSnapshotRestoresOutput
				p := svc.NewDescribeFastSnapshotRestoresPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-fleet-history": {
			Name:   "describe-fleet-history",
			Fields: fields_describe_fleet_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFleetHistoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_fleet_history, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFleetHistory(ctx, input)
			},
		},
		"describe-fleet-instances": {
			Name:   "describe-fleet-instances",
			Fields: fields_describe_fleet_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFleetInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_fleet_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFleetInstances(ctx, input)
			},
		},
		"describe-fleets": {
			Name:   "describe-fleets",
			Fields: fields_describe_fleets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFleetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_fleets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeFleets(ctx, input)
				}
				var results []*svc.DescribeFleetsOutput
				p := svc.NewDescribeFleetsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-flow-logs": {
			Name:   "describe-flow-logs",
			Fields: fields_describe_flow_logs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFlowLogsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_flow_logs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeFlowLogs(ctx, input)
				}
				var results []*svc.DescribeFlowLogsOutput
				p := svc.NewDescribeFlowLogsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-fpga-image-attribute": {
			Name:   "describe-fpga-image-attribute",
			Fields: fields_describe_fpga_image_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFpgaImageAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_fpga_image_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFpgaImageAttribute(ctx, input)
			},
		},
		"describe-fpga-images": {
			Name:   "describe-fpga-images",
			Fields: fields_describe_fpga_images,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFpgaImagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_fpga_images, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeFpgaImages(ctx, input)
				}
				var results []*svc.DescribeFpgaImagesOutput
				p := svc.NewDescribeFpgaImagesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-host-reservation-offerings": {
			Name:   "describe-host-reservation-offerings",
			Fields: fields_describe_host_reservation_offerings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeHostReservationOfferingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_host_reservation_offerings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeHostReservationOfferings(ctx, input)
				}
				var results []*svc.DescribeHostReservationOfferingsOutput
				p := svc.NewDescribeHostReservationOfferingsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-host-reservations": {
			Name:   "describe-host-reservations",
			Fields: fields_describe_host_reservations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeHostReservationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_host_reservations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeHostReservations(ctx, input)
				}
				var results []*svc.DescribeHostReservationsOutput
				p := svc.NewDescribeHostReservationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-hosts": {
			Name:   "describe-hosts",
			Fields: fields_describe_hosts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeHostsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_hosts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeHosts(ctx, input)
				}
				var results []*svc.DescribeHostsOutput
				p := svc.NewDescribeHostsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-iam-instance-profile-associations": {
			Name:   "describe-iam-instance-profile-associations",
			Fields: fields_describe_iam_instance_profile_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIamInstanceProfileAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_iam_instance_profile_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeIamInstanceProfileAssociations(ctx, input)
				}
				var results []*svc.DescribeIamInstanceProfileAssociationsOutput
				p := svc.NewDescribeIamInstanceProfileAssociationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-id-format": {
			Name:   "describe-id-format",
			Fields: fields_describe_id_format,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIdFormatInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_id_format, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeIdFormat(ctx, input)
			},
		},
		"describe-identity-id-format": {
			Name:   "describe-identity-id-format",
			Fields: fields_describe_identity_id_format,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIdentityIdFormatInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_identity_id_format, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeIdentityIdFormat(ctx, input)
			},
		},
		"describe-image-attribute": {
			Name:   "describe-image-attribute",
			Fields: fields_describe_image_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeImageAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_image_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeImageAttribute(ctx, input)
			},
		},
		"describe-image-references": {
			Name:   "describe-image-references",
			Fields: fields_describe_image_references,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeImageReferencesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_image_references, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeImageReferences(ctx, input)
				}
				var results []*svc.DescribeImageReferencesOutput
				p := svc.NewDescribeImageReferencesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-image-usage-report-entries": {
			Name:   "describe-image-usage-report-entries",
			Fields: fields_describe_image_usage_report_entries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeImageUsageReportEntriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_image_usage_report_entries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeImageUsageReportEntries(ctx, input)
				}
				var results []*svc.DescribeImageUsageReportEntriesOutput
				p := svc.NewDescribeImageUsageReportEntriesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-image-usage-reports": {
			Name:   "describe-image-usage-reports",
			Fields: fields_describe_image_usage_reports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeImageUsageReportsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_image_usage_reports, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeImageUsageReports(ctx, input)
				}
				var results []*svc.DescribeImageUsageReportsOutput
				p := svc.NewDescribeImageUsageReportsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-images": {
			Name:   "describe-images",
			Fields: fields_describe_images,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeImagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_images, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeImages(ctx, input)
				}
				var results []*svc.DescribeImagesOutput
				p := svc.NewDescribeImagesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-import-image-tasks": {
			Name:   "describe-import-image-tasks",
			Fields: fields_describe_import_image_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeImportImageTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_import_image_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeImportImageTasks(ctx, input)
				}
				var results []*svc.DescribeImportImageTasksOutput
				p := svc.NewDescribeImportImageTasksPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-import-snapshot-tasks": {
			Name:   "describe-import-snapshot-tasks",
			Fields: fields_describe_import_snapshot_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeImportSnapshotTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_import_snapshot_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeImportSnapshotTasks(ctx, input)
				}
				var results []*svc.DescribeImportSnapshotTasksOutput
				p := svc.NewDescribeImportSnapshotTasksPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-instance-attribute": {
			Name:   "describe-instance-attribute",
			Fields: fields_describe_instance_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstanceAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_instance_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeInstanceAttribute(ctx, input)
			},
		},
		"describe-instance-connect-endpoints": {
			Name:   "describe-instance-connect-endpoints",
			Fields: fields_describe_instance_connect_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstanceConnectEndpointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_instance_connect_endpoints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeInstanceConnectEndpoints(ctx, input)
				}
				var results []*svc.DescribeInstanceConnectEndpointsOutput
				p := svc.NewDescribeInstanceConnectEndpointsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-instance-credit-specifications": {
			Name:   "describe-instance-credit-specifications",
			Fields: fields_describe_instance_credit_specifications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstanceCreditSpecificationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_instance_credit_specifications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeInstanceCreditSpecifications(ctx, input)
				}
				var results []*svc.DescribeInstanceCreditSpecificationsOutput
				p := svc.NewDescribeInstanceCreditSpecificationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-instance-event-notification-attributes": {
			Name:   "describe-instance-event-notification-attributes",
			Fields: fields_describe_instance_event_notification_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstanceEventNotificationAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_instance_event_notification_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeInstanceEventNotificationAttributes(ctx, input)
			},
		},
		"describe-instance-event-windows": {
			Name:   "describe-instance-event-windows",
			Fields: fields_describe_instance_event_windows,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstanceEventWindowsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_instance_event_windows, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeInstanceEventWindows(ctx, input)
				}
				var results []*svc.DescribeInstanceEventWindowsOutput
				p := svc.NewDescribeInstanceEventWindowsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-instance-image-metadata": {
			Name:   "describe-instance-image-metadata",
			Fields: fields_describe_instance_image_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstanceImageMetadataInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_instance_image_metadata, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeInstanceImageMetadata(ctx, input)
				}
				var results []*svc.DescribeInstanceImageMetadataOutput
				p := svc.NewDescribeInstanceImageMetadataPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-instance-sql-ha-history-states": {
			Name:   "describe-instance-sql-ha-history-states",
			Fields: fields_describe_instance_sql_ha_history_states,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstanceSqlHaHistoryStatesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_instance_sql_ha_history_states, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeInstanceSqlHaHistoryStates(ctx, input)
			},
		},
		"describe-instance-sql-ha-states": {
			Name:   "describe-instance-sql-ha-states",
			Fields: fields_describe_instance_sql_ha_states,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstanceSqlHaStatesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_instance_sql_ha_states, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeInstanceSqlHaStates(ctx, input)
			},
		},
		"describe-instance-status": {
			Name:   "describe-instance-status",
			Fields: fields_describe_instance_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstanceStatusInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_instance_status, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeInstanceStatus(ctx, input)
				}
				var results []*svc.DescribeInstanceStatusOutput
				p := svc.NewDescribeInstanceStatusPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-instance-topology": {
			Name:   "describe-instance-topology",
			Fields: fields_describe_instance_topology,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstanceTopologyInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_instance_topology, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeInstanceTopology(ctx, input)
				}
				var results []*svc.DescribeInstanceTopologyOutput
				p := svc.NewDescribeInstanceTopologyPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-instance-type-offerings": {
			Name:   "describe-instance-type-offerings",
			Fields: fields_describe_instance_type_offerings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstanceTypeOfferingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_instance_type_offerings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeInstanceTypeOfferings(ctx, input)
				}
				var results []*svc.DescribeInstanceTypeOfferingsOutput
				p := svc.NewDescribeInstanceTypeOfferingsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-instance-types": {
			Name:   "describe-instance-types",
			Fields: fields_describe_instance_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstanceTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_instance_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeInstanceTypes(ctx, input)
				}
				var results []*svc.DescribeInstanceTypesOutput
				p := svc.NewDescribeInstanceTypesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-instances": {
			Name:   "describe-instances",
			Fields: fields_describe_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeInstances(ctx, input)
				}
				var results []*svc.DescribeInstancesOutput
				p := svc.NewDescribeInstancesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-internet-gateways": {
			Name:   "describe-internet-gateways",
			Fields: fields_describe_internet_gateways,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInternetGatewaysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_internet_gateways, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeInternetGateways(ctx, input)
				}
				var results []*svc.DescribeInternetGatewaysOutput
				p := svc.NewDescribeInternetGatewaysPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-ipam-byoasn": {
			Name:   "describe-ipam-byoasn",
			Fields: fields_describe_ipam_byoasn,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIpamByoasnInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_ipam_byoasn, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeIpamByoasn(ctx, input)
			},
		},
		"describe-ipam-external-resource-verification-tokens": {
			Name:   "describe-ipam-external-resource-verification-tokens",
			Fields: fields_describe_ipam_external_resource_verification_tokens,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIpamExternalResourceVerificationTokensInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_ipam_external_resource_verification_tokens, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeIpamExternalResourceVerificationTokens(ctx, input)
			},
		},
		"describe-ipam-policies": {
			Name:   "describe-ipam-policies",
			Fields: fields_describe_ipam_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIpamPoliciesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_ipam_policies, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeIpamPolicies(ctx, input)
			},
		},
		"describe-ipam-pools": {
			Name:   "describe-ipam-pools",
			Fields: fields_describe_ipam_pools,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIpamPoolsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_ipam_pools, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeIpamPools(ctx, input)
				}
				var results []*svc.DescribeIpamPoolsOutput
				p := svc.NewDescribeIpamPoolsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-ipam-prefix-list-resolver-targets": {
			Name:   "describe-ipam-prefix-list-resolver-targets",
			Fields: fields_describe_ipam_prefix_list_resolver_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIpamPrefixListResolverTargetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_ipam_prefix_list_resolver_targets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeIpamPrefixListResolverTargets(ctx, input)
				}
				var results []*svc.DescribeIpamPrefixListResolverTargetsOutput
				p := svc.NewDescribeIpamPrefixListResolverTargetsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-ipam-prefix-list-resolvers": {
			Name:   "describe-ipam-prefix-list-resolvers",
			Fields: fields_describe_ipam_prefix_list_resolvers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIpamPrefixListResolversInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_ipam_prefix_list_resolvers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeIpamPrefixListResolvers(ctx, input)
				}
				var results []*svc.DescribeIpamPrefixListResolversOutput
				p := svc.NewDescribeIpamPrefixListResolversPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-ipam-resource-discoveries": {
			Name:   "describe-ipam-resource-discoveries",
			Fields: fields_describe_ipam_resource_discoveries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIpamResourceDiscoveriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_ipam_resource_discoveries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeIpamResourceDiscoveries(ctx, input)
				}
				var results []*svc.DescribeIpamResourceDiscoveriesOutput
				p := svc.NewDescribeIpamResourceDiscoveriesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-ipam-resource-discovery-associations": {
			Name:   "describe-ipam-resource-discovery-associations",
			Fields: fields_describe_ipam_resource_discovery_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIpamResourceDiscoveryAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_ipam_resource_discovery_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeIpamResourceDiscoveryAssociations(ctx, input)
				}
				var results []*svc.DescribeIpamResourceDiscoveryAssociationsOutput
				p := svc.NewDescribeIpamResourceDiscoveryAssociationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-ipam-scopes": {
			Name:   "describe-ipam-scopes",
			Fields: fields_describe_ipam_scopes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIpamScopesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_ipam_scopes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeIpamScopes(ctx, input)
				}
				var results []*svc.DescribeIpamScopesOutput
				p := svc.NewDescribeIpamScopesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-ipams": {
			Name:   "describe-ipams",
			Fields: fields_describe_ipams,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIpamsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_ipams, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeIpams(ctx, input)
				}
				var results []*svc.DescribeIpamsOutput
				p := svc.NewDescribeIpamsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-ipv6-pools": {
			Name:   "describe-ipv6-pools",
			Fields: fields_describe_ipv6_pools,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIpv6PoolsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_ipv6_pools, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeIpv6Pools(ctx, input)
				}
				var results []*svc.DescribeIpv6PoolsOutput
				p := svc.NewDescribeIpv6PoolsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-key-pairs": {
			Name:   "describe-key-pairs",
			Fields: fields_describe_key_pairs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeKeyPairsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_key_pairs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeKeyPairs(ctx, input)
			},
		},
		"describe-launch-template-versions": {
			Name:   "describe-launch-template-versions",
			Fields: fields_describe_launch_template_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLaunchTemplateVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_launch_template_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeLaunchTemplateVersions(ctx, input)
				}
				var results []*svc.DescribeLaunchTemplateVersionsOutput
				p := svc.NewDescribeLaunchTemplateVersionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-launch-templates": {
			Name:   "describe-launch-templates",
			Fields: fields_describe_launch_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLaunchTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_launch_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeLaunchTemplates(ctx, input)
				}
				var results []*svc.DescribeLaunchTemplatesOutput
				p := svc.NewDescribeLaunchTemplatesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-local-gateway-route-table-virtual-interface-group-associations": {
			Name:   "describe-local-gateway-route-table-virtual-interface-group-associations",
			Fields: fields_describe_local_gateway_route_table_virtual_interface_group_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_local_gateway_route_table_virtual_interface_group_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(ctx, input)
				}
				var results []*svc.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput
				p := svc.NewDescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-local-gateway-route-table-vpc-associations": {
			Name:   "describe-local-gateway-route-table-vpc-associations",
			Fields: fields_describe_local_gateway_route_table_vpc_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLocalGatewayRouteTableVpcAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_local_gateway_route_table_vpc_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeLocalGatewayRouteTableVpcAssociations(ctx, input)
				}
				var results []*svc.DescribeLocalGatewayRouteTableVpcAssociationsOutput
				p := svc.NewDescribeLocalGatewayRouteTableVpcAssociationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-local-gateway-route-tables": {
			Name:   "describe-local-gateway-route-tables",
			Fields: fields_describe_local_gateway_route_tables,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLocalGatewayRouteTablesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_local_gateway_route_tables, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeLocalGatewayRouteTables(ctx, input)
				}
				var results []*svc.DescribeLocalGatewayRouteTablesOutput
				p := svc.NewDescribeLocalGatewayRouteTablesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-local-gateway-virtual-interface-groups": {
			Name:   "describe-local-gateway-virtual-interface-groups",
			Fields: fields_describe_local_gateway_virtual_interface_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLocalGatewayVirtualInterfaceGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_local_gateway_virtual_interface_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeLocalGatewayVirtualInterfaceGroups(ctx, input)
				}
				var results []*svc.DescribeLocalGatewayVirtualInterfaceGroupsOutput
				p := svc.NewDescribeLocalGatewayVirtualInterfaceGroupsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-local-gateway-virtual-interfaces": {
			Name:   "describe-local-gateway-virtual-interfaces",
			Fields: fields_describe_local_gateway_virtual_interfaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLocalGatewayVirtualInterfacesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_local_gateway_virtual_interfaces, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeLocalGatewayVirtualInterfaces(ctx, input)
				}
				var results []*svc.DescribeLocalGatewayVirtualInterfacesOutput
				p := svc.NewDescribeLocalGatewayVirtualInterfacesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-local-gateways": {
			Name:   "describe-local-gateways",
			Fields: fields_describe_local_gateways,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLocalGatewaysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_local_gateways, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeLocalGateways(ctx, input)
				}
				var results []*svc.DescribeLocalGatewaysOutput
				p := svc.NewDescribeLocalGatewaysPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-locked-snapshots": {
			Name:   "describe-locked-snapshots",
			Fields: fields_describe_locked_snapshots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLockedSnapshotsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_locked_snapshots, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeLockedSnapshots(ctx, input)
			},
		},
		"describe-mac-hosts": {
			Name:   "describe-mac-hosts",
			Fields: fields_describe_mac_hosts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMacHostsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_mac_hosts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeMacHosts(ctx, input)
				}
				var results []*svc.DescribeMacHostsOutput
				p := svc.NewDescribeMacHostsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-mac-modification-tasks": {
			Name:   "describe-mac-modification-tasks",
			Fields: fields_describe_mac_modification_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMacModificationTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_mac_modification_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeMacModificationTasks(ctx, input)
				}
				var results []*svc.DescribeMacModificationTasksOutput
				p := svc.NewDescribeMacModificationTasksPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-managed-prefix-lists": {
			Name:   "describe-managed-prefix-lists",
			Fields: fields_describe_managed_prefix_lists,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeManagedPrefixListsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_managed_prefix_lists, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeManagedPrefixLists(ctx, input)
				}
				var results []*svc.DescribeManagedPrefixListsOutput
				p := svc.NewDescribeManagedPrefixListsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-moving-addresses": {
			Name:   "describe-moving-addresses",
			Fields: fields_describe_moving_addresses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMovingAddressesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_moving_addresses, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeMovingAddresses(ctx, input)
				}
				var results []*svc.DescribeMovingAddressesOutput
				p := svc.NewDescribeMovingAddressesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-nat-gateways": {
			Name:   "describe-nat-gateways",
			Fields: fields_describe_nat_gateways,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeNatGatewaysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_nat_gateways, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeNatGateways(ctx, input)
				}
				var results []*svc.DescribeNatGatewaysOutput
				p := svc.NewDescribeNatGatewaysPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-network-acls": {
			Name:   "describe-network-acls",
			Fields: fields_describe_network_acls,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeNetworkAclsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_network_acls, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeNetworkAcls(ctx, input)
				}
				var results []*svc.DescribeNetworkAclsOutput
				p := svc.NewDescribeNetworkAclsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-network-insights-access-scope-analyses": {
			Name:   "describe-network-insights-access-scope-analyses",
			Fields: fields_describe_network_insights_access_scope_analyses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeNetworkInsightsAccessScopeAnalysesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_network_insights_access_scope_analyses, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeNetworkInsightsAccessScopeAnalyses(ctx, input)
				}
				var results []*svc.DescribeNetworkInsightsAccessScopeAnalysesOutput
				p := svc.NewDescribeNetworkInsightsAccessScopeAnalysesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-network-insights-access-scopes": {
			Name:   "describe-network-insights-access-scopes",
			Fields: fields_describe_network_insights_access_scopes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeNetworkInsightsAccessScopesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_network_insights_access_scopes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeNetworkInsightsAccessScopes(ctx, input)
				}
				var results []*svc.DescribeNetworkInsightsAccessScopesOutput
				p := svc.NewDescribeNetworkInsightsAccessScopesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-network-insights-analyses": {
			Name:   "describe-network-insights-analyses",
			Fields: fields_describe_network_insights_analyses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeNetworkInsightsAnalysesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_network_insights_analyses, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeNetworkInsightsAnalyses(ctx, input)
				}
				var results []*svc.DescribeNetworkInsightsAnalysesOutput
				p := svc.NewDescribeNetworkInsightsAnalysesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-network-insights-paths": {
			Name:   "describe-network-insights-paths",
			Fields: fields_describe_network_insights_paths,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeNetworkInsightsPathsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_network_insights_paths, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeNetworkInsightsPaths(ctx, input)
				}
				var results []*svc.DescribeNetworkInsightsPathsOutput
				p := svc.NewDescribeNetworkInsightsPathsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-network-interface-attribute": {
			Name:   "describe-network-interface-attribute",
			Fields: fields_describe_network_interface_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeNetworkInterfaceAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_network_interface_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeNetworkInterfaceAttribute(ctx, input)
			},
		},
		"describe-network-interface-permissions": {
			Name:   "describe-network-interface-permissions",
			Fields: fields_describe_network_interface_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeNetworkInterfacePermissionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_network_interface_permissions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeNetworkInterfacePermissions(ctx, input)
				}
				var results []*svc.DescribeNetworkInterfacePermissionsOutput
				p := svc.NewDescribeNetworkInterfacePermissionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-network-interfaces": {
			Name:   "describe-network-interfaces",
			Fields: fields_describe_network_interfaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeNetworkInterfacesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_network_interfaces, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeNetworkInterfaces(ctx, input)
				}
				var results []*svc.DescribeNetworkInterfacesOutput
				p := svc.NewDescribeNetworkInterfacesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-outpost-lags": {
			Name:   "describe-outpost-lags",
			Fields: fields_describe_outpost_lags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeOutpostLagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_outpost_lags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeOutpostLags(ctx, input)
			},
		},
		"describe-placement-groups": {
			Name:   "describe-placement-groups",
			Fields: fields_describe_placement_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePlacementGroupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_placement_groups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePlacementGroups(ctx, input)
			},
		},
		"describe-prefix-lists": {
			Name:   "describe-prefix-lists",
			Fields: fields_describe_prefix_lists,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePrefixListsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_prefix_lists, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribePrefixLists(ctx, input)
				}
				var results []*svc.DescribePrefixListsOutput
				p := svc.NewDescribePrefixListsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-principal-id-format": {
			Name:   "describe-principal-id-format",
			Fields: fields_describe_principal_id_format,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePrincipalIdFormatInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_principal_id_format, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribePrincipalIdFormat(ctx, input)
				}
				var results []*svc.DescribePrincipalIdFormatOutput
				p := svc.NewDescribePrincipalIdFormatPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-public-ipv4-pools": {
			Name:   "describe-public-ipv4-pools",
			Fields: fields_describe_public_ipv4_pools,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePublicIpv4PoolsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_public_ipv4_pools, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribePublicIpv4Pools(ctx, input)
				}
				var results []*svc.DescribePublicIpv4PoolsOutput
				p := svc.NewDescribePublicIpv4PoolsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-regions": {
			Name:   "describe-regions",
			Fields: fields_describe_regions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRegionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_regions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRegions(ctx, input)
			},
		},
		"describe-replace-root-volume-tasks": {
			Name:   "describe-replace-root-volume-tasks",
			Fields: fields_describe_replace_root_volume_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReplaceRootVolumeTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_replace_root_volume_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReplaceRootVolumeTasks(ctx, input)
				}
				var results []*svc.DescribeReplaceRootVolumeTasksOutput
				p := svc.NewDescribeReplaceRootVolumeTasksPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-reserved-instances": {
			Name:   "describe-reserved-instances",
			Fields: fields_describe_reserved_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReservedInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_reserved_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeReservedInstances(ctx, input)
			},
		},
		"describe-reserved-instances-listings": {
			Name:   "describe-reserved-instances-listings",
			Fields: fields_describe_reserved_instances_listings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReservedInstancesListingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_reserved_instances_listings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeReservedInstancesListings(ctx, input)
			},
		},
		"describe-reserved-instances-modifications": {
			Name:   "describe-reserved-instances-modifications",
			Fields: fields_describe_reserved_instances_modifications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReservedInstancesModificationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_reserved_instances_modifications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReservedInstancesModifications(ctx, input)
				}
				var results []*svc.DescribeReservedInstancesModificationsOutput
				p := svc.NewDescribeReservedInstancesModificationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-reserved-instances-offerings": {
			Name:   "describe-reserved-instances-offerings",
			Fields: fields_describe_reserved_instances_offerings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReservedInstancesOfferingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_reserved_instances_offerings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReservedInstancesOfferings(ctx, input)
				}
				var results []*svc.DescribeReservedInstancesOfferingsOutput
				p := svc.NewDescribeReservedInstancesOfferingsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-route-server-endpoints": {
			Name:   "describe-route-server-endpoints",
			Fields: fields_describe_route_server_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRouteServerEndpointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_route_server_endpoints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeRouteServerEndpoints(ctx, input)
				}
				var results []*svc.DescribeRouteServerEndpointsOutput
				p := svc.NewDescribeRouteServerEndpointsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-route-server-peers": {
			Name:   "describe-route-server-peers",
			Fields: fields_describe_route_server_peers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRouteServerPeersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_route_server_peers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeRouteServerPeers(ctx, input)
				}
				var results []*svc.DescribeRouteServerPeersOutput
				p := svc.NewDescribeRouteServerPeersPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-route-servers": {
			Name:   "describe-route-servers",
			Fields: fields_describe_route_servers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRouteServersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_route_servers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeRouteServers(ctx, input)
				}
				var results []*svc.DescribeRouteServersOutput
				p := svc.NewDescribeRouteServersPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-route-tables": {
			Name:   "describe-route-tables",
			Fields: fields_describe_route_tables,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRouteTablesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_route_tables, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeRouteTables(ctx, input)
				}
				var results []*svc.DescribeRouteTablesOutput
				p := svc.NewDescribeRouteTablesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-scheduled-instance-availability": {
			Name:   "describe-scheduled-instance-availability",
			Fields: fields_describe_scheduled_instance_availability,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeScheduledInstanceAvailabilityInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_scheduled_instance_availability, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeScheduledInstanceAvailability(ctx, input)
				}
				var results []*svc.DescribeScheduledInstanceAvailabilityOutput
				p := svc.NewDescribeScheduledInstanceAvailabilityPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-scheduled-instances": {
			Name:   "describe-scheduled-instances",
			Fields: fields_describe_scheduled_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeScheduledInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_scheduled_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeScheduledInstances(ctx, input)
				}
				var results []*svc.DescribeScheduledInstancesOutput
				p := svc.NewDescribeScheduledInstancesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-secondary-interfaces": {
			Name:   "describe-secondary-interfaces",
			Fields: fields_describe_secondary_interfaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSecondaryInterfacesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_secondary_interfaces, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSecondaryInterfaces(ctx, input)
				}
				var results []*svc.DescribeSecondaryInterfacesOutput
				p := svc.NewDescribeSecondaryInterfacesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-secondary-networks": {
			Name:   "describe-secondary-networks",
			Fields: fields_describe_secondary_networks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSecondaryNetworksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_secondary_networks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSecondaryNetworks(ctx, input)
				}
				var results []*svc.DescribeSecondaryNetworksOutput
				p := svc.NewDescribeSecondaryNetworksPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-secondary-subnets": {
			Name:   "describe-secondary-subnets",
			Fields: fields_describe_secondary_subnets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSecondarySubnetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_secondary_subnets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSecondarySubnets(ctx, input)
				}
				var results []*svc.DescribeSecondarySubnetsOutput
				p := svc.NewDescribeSecondarySubnetsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-security-group-references": {
			Name:   "describe-security-group-references",
			Fields: fields_describe_security_group_references,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSecurityGroupReferencesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_security_group_references, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSecurityGroupReferences(ctx, input)
			},
		},
		"describe-security-group-rules": {
			Name:   "describe-security-group-rules",
			Fields: fields_describe_security_group_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSecurityGroupRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_security_group_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSecurityGroupRules(ctx, input)
				}
				var results []*svc.DescribeSecurityGroupRulesOutput
				p := svc.NewDescribeSecurityGroupRulesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-security-group-vpc-associations": {
			Name:   "describe-security-group-vpc-associations",
			Fields: fields_describe_security_group_vpc_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSecurityGroupVpcAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_security_group_vpc_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSecurityGroupVpcAssociations(ctx, input)
				}
				var results []*svc.DescribeSecurityGroupVpcAssociationsOutput
				p := svc.NewDescribeSecurityGroupVpcAssociationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-security-groups": {
			Name:   "describe-security-groups",
			Fields: fields_describe_security_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSecurityGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_security_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSecurityGroups(ctx, input)
				}
				var results []*svc.DescribeSecurityGroupsOutput
				p := svc.NewDescribeSecurityGroupsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-service-link-virtual-interfaces": {
			Name:   "describe-service-link-virtual-interfaces",
			Fields: fields_describe_service_link_virtual_interfaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeServiceLinkVirtualInterfacesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_service_link_virtual_interfaces, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeServiceLinkVirtualInterfaces(ctx, input)
			},
		},
		"describe-snapshot-attribute": {
			Name:   "describe-snapshot-attribute",
			Fields: fields_describe_snapshot_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSnapshotAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_snapshot_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSnapshotAttribute(ctx, input)
			},
		},
		"describe-snapshot-tier-status": {
			Name:   "describe-snapshot-tier-status",
			Fields: fields_describe_snapshot_tier_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSnapshotTierStatusInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_snapshot_tier_status, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSnapshotTierStatus(ctx, input)
				}
				var results []*svc.DescribeSnapshotTierStatusOutput
				p := svc.NewDescribeSnapshotTierStatusPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-snapshots": {
			Name:   "describe-snapshots",
			Fields: fields_describe_snapshots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSnapshotsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_snapshots, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSnapshots(ctx, input)
				}
				var results []*svc.DescribeSnapshotsOutput
				p := svc.NewDescribeSnapshotsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-spot-datafeed-subscription": {
			Name:   "describe-spot-datafeed-subscription",
			Fields: fields_describe_spot_datafeed_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSpotDatafeedSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_spot_datafeed_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSpotDatafeedSubscription(ctx, input)
			},
		},
		"describe-spot-fleet-instances": {
			Name:   "describe-spot-fleet-instances",
			Fields: fields_describe_spot_fleet_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSpotFleetInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_spot_fleet_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSpotFleetInstances(ctx, input)
			},
		},
		"describe-spot-fleet-request-history": {
			Name:   "describe-spot-fleet-request-history",
			Fields: fields_describe_spot_fleet_request_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSpotFleetRequestHistoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_spot_fleet_request_history, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSpotFleetRequestHistory(ctx, input)
			},
		},
		"describe-spot-fleet-requests": {
			Name:   "describe-spot-fleet-requests",
			Fields: fields_describe_spot_fleet_requests,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSpotFleetRequestsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_spot_fleet_requests, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSpotFleetRequests(ctx, input)
				}
				var results []*svc.DescribeSpotFleetRequestsOutput
				p := svc.NewDescribeSpotFleetRequestsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-spot-instance-requests": {
			Name:   "describe-spot-instance-requests",
			Fields: fields_describe_spot_instance_requests,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSpotInstanceRequestsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_spot_instance_requests, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSpotInstanceRequests(ctx, input)
				}
				var results []*svc.DescribeSpotInstanceRequestsOutput
				p := svc.NewDescribeSpotInstanceRequestsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-spot-price-history": {
			Name:   "describe-spot-price-history",
			Fields: fields_describe_spot_price_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSpotPriceHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_spot_price_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSpotPriceHistory(ctx, input)
				}
				var results []*svc.DescribeSpotPriceHistoryOutput
				p := svc.NewDescribeSpotPriceHistoryPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-stale-security-groups": {
			Name:   "describe-stale-security-groups",
			Fields: fields_describe_stale_security_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStaleSecurityGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_stale_security_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeStaleSecurityGroups(ctx, input)
				}
				var results []*svc.DescribeStaleSecurityGroupsOutput
				p := svc.NewDescribeStaleSecurityGroupsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-store-image-tasks": {
			Name:   "describe-store-image-tasks",
			Fields: fields_describe_store_image_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStoreImageTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_store_image_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeStoreImageTasks(ctx, input)
				}
				var results []*svc.DescribeStoreImageTasksOutput
				p := svc.NewDescribeStoreImageTasksPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-subnets": {
			Name:   "describe-subnets",
			Fields: fields_describe_subnets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSubnetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_subnets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSubnets(ctx, input)
				}
				var results []*svc.DescribeSubnetsOutput
				p := svc.NewDescribeSubnetsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-tags": {
			Name:   "describe-tags",
			Fields: fields_describe_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTagsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_tags, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTags(ctx, input)
				}
				var results []*svc.DescribeTagsOutput
				p := svc.NewDescribeTagsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-traffic-mirror-filter-rules": {
			Name:   "describe-traffic-mirror-filter-rules",
			Fields: fields_describe_traffic_mirror_filter_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTrafficMirrorFilterRulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_traffic_mirror_filter_rules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTrafficMirrorFilterRules(ctx, input)
			},
		},
		"describe-traffic-mirror-filters": {
			Name:   "describe-traffic-mirror-filters",
			Fields: fields_describe_traffic_mirror_filters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTrafficMirrorFiltersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_traffic_mirror_filters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTrafficMirrorFilters(ctx, input)
				}
				var results []*svc.DescribeTrafficMirrorFiltersOutput
				p := svc.NewDescribeTrafficMirrorFiltersPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-traffic-mirror-sessions": {
			Name:   "describe-traffic-mirror-sessions",
			Fields: fields_describe_traffic_mirror_sessions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTrafficMirrorSessionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_traffic_mirror_sessions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTrafficMirrorSessions(ctx, input)
				}
				var results []*svc.DescribeTrafficMirrorSessionsOutput
				p := svc.NewDescribeTrafficMirrorSessionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-traffic-mirror-targets": {
			Name:   "describe-traffic-mirror-targets",
			Fields: fields_describe_traffic_mirror_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTrafficMirrorTargetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_traffic_mirror_targets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTrafficMirrorTargets(ctx, input)
				}
				var results []*svc.DescribeTrafficMirrorTargetsOutput
				p := svc.NewDescribeTrafficMirrorTargetsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-transit-gateway-attachments": {
			Name:   "describe-transit-gateway-attachments",
			Fields: fields_describe_transit_gateway_attachments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTransitGatewayAttachmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_transit_gateway_attachments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTransitGatewayAttachments(ctx, input)
				}
				var results []*svc.DescribeTransitGatewayAttachmentsOutput
				p := svc.NewDescribeTransitGatewayAttachmentsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-transit-gateway-connect-peers": {
			Name:   "describe-transit-gateway-connect-peers",
			Fields: fields_describe_transit_gateway_connect_peers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTransitGatewayConnectPeersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_transit_gateway_connect_peers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTransitGatewayConnectPeers(ctx, input)
				}
				var results []*svc.DescribeTransitGatewayConnectPeersOutput
				p := svc.NewDescribeTransitGatewayConnectPeersPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-transit-gateway-connects": {
			Name:   "describe-transit-gateway-connects",
			Fields: fields_describe_transit_gateway_connects,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTransitGatewayConnectsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_transit_gateway_connects, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTransitGatewayConnects(ctx, input)
				}
				var results []*svc.DescribeTransitGatewayConnectsOutput
				p := svc.NewDescribeTransitGatewayConnectsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-transit-gateway-metering-policies": {
			Name:   "describe-transit-gateway-metering-policies",
			Fields: fields_describe_transit_gateway_metering_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTransitGatewayMeteringPoliciesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_transit_gateway_metering_policies, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTransitGatewayMeteringPolicies(ctx, input)
			},
		},
		"describe-transit-gateway-multicast-domains": {
			Name:   "describe-transit-gateway-multicast-domains",
			Fields: fields_describe_transit_gateway_multicast_domains,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTransitGatewayMulticastDomainsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_transit_gateway_multicast_domains, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTransitGatewayMulticastDomains(ctx, input)
				}
				var results []*svc.DescribeTransitGatewayMulticastDomainsOutput
				p := svc.NewDescribeTransitGatewayMulticastDomainsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-transit-gateway-peering-attachments": {
			Name:   "describe-transit-gateway-peering-attachments",
			Fields: fields_describe_transit_gateway_peering_attachments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTransitGatewayPeeringAttachmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_transit_gateway_peering_attachments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTransitGatewayPeeringAttachments(ctx, input)
				}
				var results []*svc.DescribeTransitGatewayPeeringAttachmentsOutput
				p := svc.NewDescribeTransitGatewayPeeringAttachmentsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-transit-gateway-policy-tables": {
			Name:   "describe-transit-gateway-policy-tables",
			Fields: fields_describe_transit_gateway_policy_tables,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTransitGatewayPolicyTablesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_transit_gateway_policy_tables, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTransitGatewayPolicyTables(ctx, input)
				}
				var results []*svc.DescribeTransitGatewayPolicyTablesOutput
				p := svc.NewDescribeTransitGatewayPolicyTablesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-transit-gateway-route-table-announcements": {
			Name:   "describe-transit-gateway-route-table-announcements",
			Fields: fields_describe_transit_gateway_route_table_announcements,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTransitGatewayRouteTableAnnouncementsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_transit_gateway_route_table_announcements, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTransitGatewayRouteTableAnnouncements(ctx, input)
				}
				var results []*svc.DescribeTransitGatewayRouteTableAnnouncementsOutput
				p := svc.NewDescribeTransitGatewayRouteTableAnnouncementsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-transit-gateway-route-tables": {
			Name:   "describe-transit-gateway-route-tables",
			Fields: fields_describe_transit_gateway_route_tables,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTransitGatewayRouteTablesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_transit_gateway_route_tables, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTransitGatewayRouteTables(ctx, input)
				}
				var results []*svc.DescribeTransitGatewayRouteTablesOutput
				p := svc.NewDescribeTransitGatewayRouteTablesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-transit-gateway-vpc-attachments": {
			Name:   "describe-transit-gateway-vpc-attachments",
			Fields: fields_describe_transit_gateway_vpc_attachments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTransitGatewayVpcAttachmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_transit_gateway_vpc_attachments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTransitGatewayVpcAttachments(ctx, input)
				}
				var results []*svc.DescribeTransitGatewayVpcAttachmentsOutput
				p := svc.NewDescribeTransitGatewayVpcAttachmentsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-transit-gateways": {
			Name:   "describe-transit-gateways",
			Fields: fields_describe_transit_gateways,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTransitGatewaysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_transit_gateways, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTransitGateways(ctx, input)
				}
				var results []*svc.DescribeTransitGatewaysOutput
				p := svc.NewDescribeTransitGatewaysPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-trunk-interface-associations": {
			Name:   "describe-trunk-interface-associations",
			Fields: fields_describe_trunk_interface_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTrunkInterfaceAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_trunk_interface_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTrunkInterfaceAssociations(ctx, input)
				}
				var results []*svc.DescribeTrunkInterfaceAssociationsOutput
				p := svc.NewDescribeTrunkInterfaceAssociationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-verified-access-endpoints": {
			Name:   "describe-verified-access-endpoints",
			Fields: fields_describe_verified_access_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVerifiedAccessEndpointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_verified_access_endpoints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeVerifiedAccessEndpoints(ctx, input)
				}
				var results []*svc.DescribeVerifiedAccessEndpointsOutput
				p := svc.NewDescribeVerifiedAccessEndpointsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-verified-access-groups": {
			Name:   "describe-verified-access-groups",
			Fields: fields_describe_verified_access_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVerifiedAccessGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_verified_access_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeVerifiedAccessGroups(ctx, input)
				}
				var results []*svc.DescribeVerifiedAccessGroupsOutput
				p := svc.NewDescribeVerifiedAccessGroupsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-verified-access-instance-logging-configurations": {
			Name:   "describe-verified-access-instance-logging-configurations",
			Fields: fields_describe_verified_access_instance_logging_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVerifiedAccessInstanceLoggingConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_verified_access_instance_logging_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeVerifiedAccessInstanceLoggingConfigurations(ctx, input)
				}
				var results []*svc.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput
				p := svc.NewDescribeVerifiedAccessInstanceLoggingConfigurationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-verified-access-instances": {
			Name:   "describe-verified-access-instances",
			Fields: fields_describe_verified_access_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVerifiedAccessInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_verified_access_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeVerifiedAccessInstances(ctx, input)
				}
				var results []*svc.DescribeVerifiedAccessInstancesOutput
				p := svc.NewDescribeVerifiedAccessInstancesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-verified-access-trust-providers": {
			Name:   "describe-verified-access-trust-providers",
			Fields: fields_describe_verified_access_trust_providers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVerifiedAccessTrustProvidersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_verified_access_trust_providers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeVerifiedAccessTrustProviders(ctx, input)
				}
				var results []*svc.DescribeVerifiedAccessTrustProvidersOutput
				p := svc.NewDescribeVerifiedAccessTrustProvidersPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-volume-attribute": {
			Name:   "describe-volume-attribute",
			Fields: fields_describe_volume_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVolumeAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_volume_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVolumeAttribute(ctx, input)
			},
		},
		"describe-volume-status": {
			Name:   "describe-volume-status",
			Fields: fields_describe_volume_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVolumeStatusInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_volume_status, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeVolumeStatus(ctx, input)
				}
				var results []*svc.DescribeVolumeStatusOutput
				p := svc.NewDescribeVolumeStatusPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-volumes": {
			Name:   "describe-volumes",
			Fields: fields_describe_volumes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVolumesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_volumes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeVolumes(ctx, input)
				}
				var results []*svc.DescribeVolumesOutput
				p := svc.NewDescribeVolumesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-volumes-modifications": {
			Name:   "describe-volumes-modifications",
			Fields: fields_describe_volumes_modifications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVolumesModificationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_volumes_modifications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeVolumesModifications(ctx, input)
				}
				var results []*svc.DescribeVolumesModificationsOutput
				p := svc.NewDescribeVolumesModificationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-vpc-attribute": {
			Name:   "describe-vpc-attribute",
			Fields: fields_describe_vpc_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVpcAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_vpc_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVpcAttribute(ctx, input)
			},
		},
		"describe-vpc-block-public-access-exclusions": {
			Name:   "describe-vpc-block-public-access-exclusions",
			Fields: fields_describe_vpc_block_public_access_exclusions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVpcBlockPublicAccessExclusionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_vpc_block_public_access_exclusions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVpcBlockPublicAccessExclusions(ctx, input)
			},
		},
		"describe-vpc-block-public-access-options": {
			Name:   "describe-vpc-block-public-access-options",
			Fields: fields_describe_vpc_block_public_access_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVpcBlockPublicAccessOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_vpc_block_public_access_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVpcBlockPublicAccessOptions(ctx, input)
			},
		},
		"describe-vpc-classic-link": {
			Name:   "describe-vpc-classic-link",
			Fields: fields_describe_vpc_classic_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVpcClassicLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_vpc_classic_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVpcClassicLink(ctx, input)
			},
		},
		"describe-vpc-classic-link-dns-support": {
			Name:   "describe-vpc-classic-link-dns-support",
			Fields: fields_describe_vpc_classic_link_dns_support,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVpcClassicLinkDnsSupportInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_vpc_classic_link_dns_support, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeVpcClassicLinkDnsSupport(ctx, input)
				}
				var results []*svc.DescribeVpcClassicLinkDnsSupportOutput
				p := svc.NewDescribeVpcClassicLinkDnsSupportPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-vpc-encryption-controls": {
			Name:   "describe-vpc-encryption-controls",
			Fields: fields_describe_vpc_encryption_controls,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVpcEncryptionControlsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_vpc_encryption_controls, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVpcEncryptionControls(ctx, input)
			},
		},
		"describe-vpc-endpoint-associations": {
			Name:   "describe-vpc-endpoint-associations",
			Fields: fields_describe_vpc_endpoint_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVpcEndpointAssociationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_vpc_endpoint_associations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVpcEndpointAssociations(ctx, input)
			},
		},
		"describe-vpc-endpoint-connection-notifications": {
			Name:   "describe-vpc-endpoint-connection-notifications",
			Fields: fields_describe_vpc_endpoint_connection_notifications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVpcEndpointConnectionNotificationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_vpc_endpoint_connection_notifications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeVpcEndpointConnectionNotifications(ctx, input)
				}
				var results []*svc.DescribeVpcEndpointConnectionNotificationsOutput
				p := svc.NewDescribeVpcEndpointConnectionNotificationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-vpc-endpoint-connections": {
			Name:   "describe-vpc-endpoint-connections",
			Fields: fields_describe_vpc_endpoint_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVpcEndpointConnectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_vpc_endpoint_connections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeVpcEndpointConnections(ctx, input)
				}
				var results []*svc.DescribeVpcEndpointConnectionsOutput
				p := svc.NewDescribeVpcEndpointConnectionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-vpc-endpoint-service-configurations": {
			Name:   "describe-vpc-endpoint-service-configurations",
			Fields: fields_describe_vpc_endpoint_service_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVpcEndpointServiceConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_vpc_endpoint_service_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeVpcEndpointServiceConfigurations(ctx, input)
				}
				var results []*svc.DescribeVpcEndpointServiceConfigurationsOutput
				p := svc.NewDescribeVpcEndpointServiceConfigurationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-vpc-endpoint-service-permissions": {
			Name:   "describe-vpc-endpoint-service-permissions",
			Fields: fields_describe_vpc_endpoint_service_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVpcEndpointServicePermissionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_vpc_endpoint_service_permissions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeVpcEndpointServicePermissions(ctx, input)
				}
				var results []*svc.DescribeVpcEndpointServicePermissionsOutput
				p := svc.NewDescribeVpcEndpointServicePermissionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-vpc-endpoint-services": {
			Name:   "describe-vpc-endpoint-services",
			Fields: fields_describe_vpc_endpoint_services,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVpcEndpointServicesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_vpc_endpoint_services, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVpcEndpointServices(ctx, input)
			},
		},
		"describe-vpc-endpoints": {
			Name:   "describe-vpc-endpoints",
			Fields: fields_describe_vpc_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVpcEndpointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_vpc_endpoints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeVpcEndpoints(ctx, input)
				}
				var results []*svc.DescribeVpcEndpointsOutput
				p := svc.NewDescribeVpcEndpointsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-vpc-peering-connections": {
			Name:   "describe-vpc-peering-connections",
			Fields: fields_describe_vpc_peering_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVpcPeeringConnectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_vpc_peering_connections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeVpcPeeringConnections(ctx, input)
				}
				var results []*svc.DescribeVpcPeeringConnectionsOutput
				p := svc.NewDescribeVpcPeeringConnectionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-vpcs": {
			Name:   "describe-vpcs",
			Fields: fields_describe_vpcs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVpcsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_vpcs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeVpcs(ctx, input)
				}
				var results []*svc.DescribeVpcsOutput
				p := svc.NewDescribeVpcsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-vpn-concentrators": {
			Name:   "describe-vpn-concentrators",
			Fields: fields_describe_vpn_concentrators,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVpnConcentratorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_vpn_concentrators, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeVpnConcentrators(ctx, input)
				}
				var results []*svc.DescribeVpnConcentratorsOutput
				p := svc.NewDescribeVpnConcentratorsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-vpn-connections": {
			Name:   "describe-vpn-connections",
			Fields: fields_describe_vpn_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVpnConnectionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_vpn_connections, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVpnConnections(ctx, input)
			},
		},
		"describe-vpn-gateways": {
			Name:   "describe-vpn-gateways",
			Fields: fields_describe_vpn_gateways,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVpnGatewaysInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_vpn_gateways, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVpnGateways(ctx, input)
			},
		},
		"detach-classic-link-vpc": {
			Name:   "detach-classic-link-vpc",
			Fields: fields_detach_classic_link_vpc,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachClassicLinkVpcInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_classic_link_vpc, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachClassicLinkVpc(ctx, input)
			},
		},
		"detach-internet-gateway": {
			Name:   "detach-internet-gateway",
			Fields: fields_detach_internet_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachInternetGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_internet_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachInternetGateway(ctx, input)
			},
		},
		"detach-network-interface": {
			Name:   "detach-network-interface",
			Fields: fields_detach_network_interface,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachNetworkInterfaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_network_interface, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachNetworkInterface(ctx, input)
			},
		},
		"detach-verified-access-trust-provider": {
			Name:   "detach-verified-access-trust-provider",
			Fields: fields_detach_verified_access_trust_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachVerifiedAccessTrustProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_verified_access_trust_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachVerifiedAccessTrustProvider(ctx, input)
			},
		},
		"detach-volume": {
			Name:   "detach-volume",
			Fields: fields_detach_volume,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachVolumeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_volume, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachVolume(ctx, input)
			},
		},
		"detach-vpn-gateway": {
			Name:   "detach-vpn-gateway",
			Fields: fields_detach_vpn_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachVpnGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_vpn_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachVpnGateway(ctx, input)
			},
		},
		"disable-address-transfer": {
			Name:   "disable-address-transfer",
			Fields: fields_disable_address_transfer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableAddressTransferInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_address_transfer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableAddressTransfer(ctx, input)
			},
		},
		"disable-allowed-images-settings": {
			Name:   "disable-allowed-images-settings",
			Fields: fields_disable_allowed_images_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableAllowedImagesSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_allowed_images_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableAllowedImagesSettings(ctx, input)
			},
		},
		"disable-aws-network-performance-metric-subscription": {
			Name:   "disable-aws-network-performance-metric-subscription",
			Fields: fields_disable_aws_network_performance_metric_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableAwsNetworkPerformanceMetricSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_aws_network_performance_metric_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableAwsNetworkPerformanceMetricSubscription(ctx, input)
			},
		},
		"disable-capacity-manager": {
			Name:   "disable-capacity-manager",
			Fields: fields_disable_capacity_manager,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableCapacityManagerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_capacity_manager, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableCapacityManager(ctx, input)
			},
		},
		"disable-ebs-encryption-by-default": {
			Name:   "disable-ebs-encryption-by-default",
			Fields: fields_disable_ebs_encryption_by_default,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableEbsEncryptionByDefaultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_ebs_encryption_by_default, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableEbsEncryptionByDefault(ctx, input)
			},
		},
		"disable-fast-launch": {
			Name:   "disable-fast-launch",
			Fields: fields_disable_fast_launch,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableFastLaunchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_fast_launch, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableFastLaunch(ctx, input)
			},
		},
		"disable-fast-snapshot-restores": {
			Name:   "disable-fast-snapshot-restores",
			Fields: fields_disable_fast_snapshot_restores,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableFastSnapshotRestoresInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_fast_snapshot_restores, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableFastSnapshotRestores(ctx, input)
			},
		},
		"disable-image": {
			Name:   "disable-image",
			Fields: fields_disable_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableImage(ctx, input)
			},
		},
		"disable-image-block-public-access": {
			Name:   "disable-image-block-public-access",
			Fields: fields_disable_image_block_public_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableImageBlockPublicAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_image_block_public_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableImageBlockPublicAccess(ctx, input)
			},
		},
		"disable-image-deprecation": {
			Name:   "disable-image-deprecation",
			Fields: fields_disable_image_deprecation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableImageDeprecationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_image_deprecation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableImageDeprecation(ctx, input)
			},
		},
		"disable-image-deregistration-protection": {
			Name:   "disable-image-deregistration-protection",
			Fields: fields_disable_image_deregistration_protection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableImageDeregistrationProtectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_image_deregistration_protection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableImageDeregistrationProtection(ctx, input)
			},
		},
		"disable-instance-sql-ha-standby-detections": {
			Name:   "disable-instance-sql-ha-standby-detections",
			Fields: fields_disable_instance_sql_ha_standby_detections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableInstanceSqlHaStandbyDetectionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_instance_sql_ha_standby_detections, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableInstanceSqlHaStandbyDetections(ctx, input)
			},
		},
		"disable-ipam-organization-admin-account": {
			Name:   "disable-ipam-organization-admin-account",
			Fields: fields_disable_ipam_organization_admin_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableIpamOrganizationAdminAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_ipam_organization_admin_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableIpamOrganizationAdminAccount(ctx, input)
			},
		},
		"disable-ipam-policy": {
			Name:   "disable-ipam-policy",
			Fields: fields_disable_ipam_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableIpamPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_ipam_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableIpamPolicy(ctx, input)
			},
		},
		"disable-route-server-propagation": {
			Name:   "disable-route-server-propagation",
			Fields: fields_disable_route_server_propagation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableRouteServerPropagationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_route_server_propagation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableRouteServerPropagation(ctx, input)
			},
		},
		"disable-serial-console-access": {
			Name:   "disable-serial-console-access",
			Fields: fields_disable_serial_console_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableSerialConsoleAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_serial_console_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableSerialConsoleAccess(ctx, input)
			},
		},
		"disable-snapshot-block-public-access": {
			Name:   "disable-snapshot-block-public-access",
			Fields: fields_disable_snapshot_block_public_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableSnapshotBlockPublicAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_snapshot_block_public_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableSnapshotBlockPublicAccess(ctx, input)
			},
		},
		"disable-transit-gateway-route-table-propagation": {
			Name:   "disable-transit-gateway-route-table-propagation",
			Fields: fields_disable_transit_gateway_route_table_propagation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableTransitGatewayRouteTablePropagationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_transit_gateway_route_table_propagation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableTransitGatewayRouteTablePropagation(ctx, input)
			},
		},
		"disable-vgw-route-propagation": {
			Name:   "disable-vgw-route-propagation",
			Fields: fields_disable_vgw_route_propagation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableVgwRoutePropagationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_vgw_route_propagation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableVgwRoutePropagation(ctx, input)
			},
		},
		"disable-vpc-classic-link": {
			Name:   "disable-vpc-classic-link",
			Fields: fields_disable_vpc_classic_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableVpcClassicLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_vpc_classic_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableVpcClassicLink(ctx, input)
			},
		},
		"disable-vpc-classic-link-dns-support": {
			Name:   "disable-vpc-classic-link-dns-support",
			Fields: fields_disable_vpc_classic_link_dns_support,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableVpcClassicLinkDnsSupportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_vpc_classic_link_dns_support, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableVpcClassicLinkDnsSupport(ctx, input)
			},
		},
		"disassociate-address": {
			Name:   "disassociate-address",
			Fields: fields_disassociate_address,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateAddressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_address, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateAddress(ctx, input)
			},
		},
		"disassociate-capacity-reservation-billing-owner": {
			Name:   "disassociate-capacity-reservation-billing-owner",
			Fields: fields_disassociate_capacity_reservation_billing_owner,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateCapacityReservationBillingOwnerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_capacity_reservation_billing_owner, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateCapacityReservationBillingOwner(ctx, input)
			},
		},
		"disassociate-client-vpn-target-network": {
			Name:   "disassociate-client-vpn-target-network",
			Fields: fields_disassociate_client_vpn_target_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateClientVpnTargetNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_client_vpn_target_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateClientVpnTargetNetwork(ctx, input)
			},
		},
		"disassociate-enclave-certificate-iam-role": {
			Name:   "disassociate-enclave-certificate-iam-role",
			Fields: fields_disassociate_enclave_certificate_iam_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateEnclaveCertificateIamRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_enclave_certificate_iam_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateEnclaveCertificateIamRole(ctx, input)
			},
		},
		"disassociate-iam-instance-profile": {
			Name:   "disassociate-iam-instance-profile",
			Fields: fields_disassociate_iam_instance_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateIamInstanceProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_iam_instance_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateIamInstanceProfile(ctx, input)
			},
		},
		"disassociate-instance-event-window": {
			Name:   "disassociate-instance-event-window",
			Fields: fields_disassociate_instance_event_window,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateInstanceEventWindowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_instance_event_window, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateInstanceEventWindow(ctx, input)
			},
		},
		"disassociate-ipam-byoasn": {
			Name:   "disassociate-ipam-byoasn",
			Fields: fields_disassociate_ipam_byoasn,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateIpamByoasnInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_ipam_byoasn, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateIpamByoasn(ctx, input)
			},
		},
		"disassociate-ipam-resource-discovery": {
			Name:   "disassociate-ipam-resource-discovery",
			Fields: fields_disassociate_ipam_resource_discovery,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateIpamResourceDiscoveryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_ipam_resource_discovery, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateIpamResourceDiscovery(ctx, input)
			},
		},
		"disassociate-nat-gateway-address": {
			Name:   "disassociate-nat-gateway-address",
			Fields: fields_disassociate_nat_gateway_address,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateNatGatewayAddressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_nat_gateway_address, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateNatGatewayAddress(ctx, input)
			},
		},
		"disassociate-route-server": {
			Name:   "disassociate-route-server",
			Fields: fields_disassociate_route_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateRouteServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_route_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateRouteServer(ctx, input)
			},
		},
		"disassociate-route-table": {
			Name:   "disassociate-route-table",
			Fields: fields_disassociate_route_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateRouteTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_route_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateRouteTable(ctx, input)
			},
		},
		"disassociate-security-group-vpc": {
			Name:   "disassociate-security-group-vpc",
			Fields: fields_disassociate_security_group_vpc,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateSecurityGroupVpcInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_security_group_vpc, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateSecurityGroupVpc(ctx, input)
			},
		},
		"disassociate-subnet-cidr-block": {
			Name:   "disassociate-subnet-cidr-block",
			Fields: fields_disassociate_subnet_cidr_block,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateSubnetCidrBlockInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_subnet_cidr_block, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateSubnetCidrBlock(ctx, input)
			},
		},
		"disassociate-transit-gateway-multicast-domain": {
			Name:   "disassociate-transit-gateway-multicast-domain",
			Fields: fields_disassociate_transit_gateway_multicast_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateTransitGatewayMulticastDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_transit_gateway_multicast_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateTransitGatewayMulticastDomain(ctx, input)
			},
		},
		"disassociate-transit-gateway-policy-table": {
			Name:   "disassociate-transit-gateway-policy-table",
			Fields: fields_disassociate_transit_gateway_policy_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateTransitGatewayPolicyTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_transit_gateway_policy_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateTransitGatewayPolicyTable(ctx, input)
			},
		},
		"disassociate-transit-gateway-route-table": {
			Name:   "disassociate-transit-gateway-route-table",
			Fields: fields_disassociate_transit_gateway_route_table,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateTransitGatewayRouteTableInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_transit_gateway_route_table, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateTransitGatewayRouteTable(ctx, input)
			},
		},
		"disassociate-trunk-interface": {
			Name:   "disassociate-trunk-interface",
			Fields: fields_disassociate_trunk_interface,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateTrunkInterfaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_trunk_interface, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateTrunkInterface(ctx, input)
			},
		},
		"disassociate-vpc-cidr-block": {
			Name:   "disassociate-vpc-cidr-block",
			Fields: fields_disassociate_vpc_cidr_block,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateVpcCidrBlockInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_vpc_cidr_block, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateVpcCidrBlock(ctx, input)
			},
		},
		"enable-address-transfer": {
			Name:   "enable-address-transfer",
			Fields: fields_enable_address_transfer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableAddressTransferInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_address_transfer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableAddressTransfer(ctx, input)
			},
		},
		"enable-allowed-images-settings": {
			Name:   "enable-allowed-images-settings",
			Fields: fields_enable_allowed_images_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableAllowedImagesSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_allowed_images_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableAllowedImagesSettings(ctx, input)
			},
		},
		"enable-aws-network-performance-metric-subscription": {
			Name:   "enable-aws-network-performance-metric-subscription",
			Fields: fields_enable_aws_network_performance_metric_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableAwsNetworkPerformanceMetricSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_aws_network_performance_metric_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableAwsNetworkPerformanceMetricSubscription(ctx, input)
			},
		},
		"enable-capacity-manager": {
			Name:   "enable-capacity-manager",
			Fields: fields_enable_capacity_manager,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableCapacityManagerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_capacity_manager, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableCapacityManager(ctx, input)
			},
		},
		"enable-ebs-encryption-by-default": {
			Name:   "enable-ebs-encryption-by-default",
			Fields: fields_enable_ebs_encryption_by_default,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableEbsEncryptionByDefaultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_ebs_encryption_by_default, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableEbsEncryptionByDefault(ctx, input)
			},
		},
		"enable-fast-launch": {
			Name:   "enable-fast-launch",
			Fields: fields_enable_fast_launch,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableFastLaunchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_fast_launch, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableFastLaunch(ctx, input)
			},
		},
		"enable-fast-snapshot-restores": {
			Name:   "enable-fast-snapshot-restores",
			Fields: fields_enable_fast_snapshot_restores,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableFastSnapshotRestoresInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_fast_snapshot_restores, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableFastSnapshotRestores(ctx, input)
			},
		},
		"enable-image": {
			Name:   "enable-image",
			Fields: fields_enable_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableImage(ctx, input)
			},
		},
		"enable-image-block-public-access": {
			Name:   "enable-image-block-public-access",
			Fields: fields_enable_image_block_public_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableImageBlockPublicAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_image_block_public_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableImageBlockPublicAccess(ctx, input)
			},
		},
		"enable-image-deprecation": {
			Name:   "enable-image-deprecation",
			Fields: fields_enable_image_deprecation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableImageDeprecationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_image_deprecation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableImageDeprecation(ctx, input)
			},
		},
		"enable-image-deregistration-protection": {
			Name:   "enable-image-deregistration-protection",
			Fields: fields_enable_image_deregistration_protection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableImageDeregistrationProtectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_image_deregistration_protection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableImageDeregistrationProtection(ctx, input)
			},
		},
		"enable-instance-sql-ha-standby-detections": {
			Name:   "enable-instance-sql-ha-standby-detections",
			Fields: fields_enable_instance_sql_ha_standby_detections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableInstanceSqlHaStandbyDetectionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_instance_sql_ha_standby_detections, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableInstanceSqlHaStandbyDetections(ctx, input)
			},
		},
		"enable-ipam-organization-admin-account": {
			Name:   "enable-ipam-organization-admin-account",
			Fields: fields_enable_ipam_organization_admin_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableIpamOrganizationAdminAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_ipam_organization_admin_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableIpamOrganizationAdminAccount(ctx, input)
			},
		},
		"enable-ipam-policy": {
			Name:   "enable-ipam-policy",
			Fields: fields_enable_ipam_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableIpamPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_ipam_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableIpamPolicy(ctx, input)
			},
		},
		"enable-reachability-analyzer-organization-sharing": {
			Name:   "enable-reachability-analyzer-organization-sharing",
			Fields: fields_enable_reachability_analyzer_organization_sharing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableReachabilityAnalyzerOrganizationSharingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_reachability_analyzer_organization_sharing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableReachabilityAnalyzerOrganizationSharing(ctx, input)
			},
		},
		"enable-route-server-propagation": {
			Name:   "enable-route-server-propagation",
			Fields: fields_enable_route_server_propagation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableRouteServerPropagationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_route_server_propagation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableRouteServerPropagation(ctx, input)
			},
		},
		"enable-serial-console-access": {
			Name:   "enable-serial-console-access",
			Fields: fields_enable_serial_console_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableSerialConsoleAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_serial_console_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableSerialConsoleAccess(ctx, input)
			},
		},
		"enable-snapshot-block-public-access": {
			Name:   "enable-snapshot-block-public-access",
			Fields: fields_enable_snapshot_block_public_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableSnapshotBlockPublicAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_snapshot_block_public_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableSnapshotBlockPublicAccess(ctx, input)
			},
		},
		"enable-transit-gateway-route-table-propagation": {
			Name:   "enable-transit-gateway-route-table-propagation",
			Fields: fields_enable_transit_gateway_route_table_propagation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableTransitGatewayRouteTablePropagationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_transit_gateway_route_table_propagation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableTransitGatewayRouteTablePropagation(ctx, input)
			},
		},
		"enable-vgw-route-propagation": {
			Name:   "enable-vgw-route-propagation",
			Fields: fields_enable_vgw_route_propagation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableVgwRoutePropagationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_vgw_route_propagation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableVgwRoutePropagation(ctx, input)
			},
		},
		"enable-volume-io": {
			Name:   "enable-volume-io",
			Fields: fields_enable_volume_io,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableVolumeIOInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_volume_io, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableVolumeIO(ctx, input)
			},
		},
		"enable-vpc-classic-link": {
			Name:   "enable-vpc-classic-link",
			Fields: fields_enable_vpc_classic_link,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableVpcClassicLinkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_vpc_classic_link, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableVpcClassicLink(ctx, input)
			},
		},
		"enable-vpc-classic-link-dns-support": {
			Name:   "enable-vpc-classic-link-dns-support",
			Fields: fields_enable_vpc_classic_link_dns_support,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableVpcClassicLinkDnsSupportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_vpc_classic_link_dns_support, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableVpcClassicLinkDnsSupport(ctx, input)
			},
		},
		"export-client-vpn-client-certificate-revocation-list": {
			Name:   "export-client-vpn-client-certificate-revocation-list",
			Fields: fields_export_client_vpn_client_certificate_revocation_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportClientVpnClientCertificateRevocationListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_client_vpn_client_certificate_revocation_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportClientVpnClientCertificateRevocationList(ctx, input)
			},
		},
		"export-client-vpn-client-configuration": {
			Name:   "export-client-vpn-client-configuration",
			Fields: fields_export_client_vpn_client_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportClientVpnClientConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_client_vpn_client_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportClientVpnClientConfiguration(ctx, input)
			},
		},
		"export-image": {
			Name:   "export-image",
			Fields: fields_export_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportImage(ctx, input)
			},
		},
		"export-transit-gateway-routes": {
			Name:   "export-transit-gateway-routes",
			Fields: fields_export_transit_gateway_routes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportTransitGatewayRoutesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_transit_gateway_routes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportTransitGatewayRoutes(ctx, input)
			},
		},
		"export-verified-access-instance-client-configuration": {
			Name:   "export-verified-access-instance-client-configuration",
			Fields: fields_export_verified_access_instance_client_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportVerifiedAccessInstanceClientConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_verified_access_instance_client_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportVerifiedAccessInstanceClientConfiguration(ctx, input)
			},
		},
		"get-active-vpn-tunnel-status": {
			Name:   "get-active-vpn-tunnel-status",
			Fields: fields_get_active_vpn_tunnel_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetActiveVpnTunnelStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_active_vpn_tunnel_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetActiveVpnTunnelStatus(ctx, input)
			},
		},
		"get-allowed-images-settings": {
			Name:   "get-allowed-images-settings",
			Fields: fields_get_allowed_images_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAllowedImagesSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_allowed_images_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAllowedImagesSettings(ctx, input)
			},
		},
		"get-associated-enclave-certificate-iam-roles": {
			Name:   "get-associated-enclave-certificate-iam-roles",
			Fields: fields_get_associated_enclave_certificate_iam_roles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAssociatedEnclaveCertificateIamRolesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_associated_enclave_certificate_iam_roles, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAssociatedEnclaveCertificateIamRoles(ctx, input)
			},
		},
		"get-associated-ipv6-pool-cidrs": {
			Name:   "get-associated-ipv6-pool-cidrs",
			Fields: fields_get_associated_ipv6_pool_cidrs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAssociatedIpv6PoolCidrsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_associated_ipv6_pool_cidrs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetAssociatedIpv6PoolCidrs(ctx, input)
				}
				var results []*svc.GetAssociatedIpv6PoolCidrsOutput
				p := svc.NewGetAssociatedIpv6PoolCidrsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-aws-network-performance-data": {
			Name:   "get-aws-network-performance-data",
			Fields: fields_get_aws_network_performance_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAwsNetworkPerformanceDataInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_aws_network_performance_data, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetAwsNetworkPerformanceData(ctx, input)
				}
				var results []*svc.GetAwsNetworkPerformanceDataOutput
				p := svc.NewGetAwsNetworkPerformanceDataPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-capacity-manager-attributes": {
			Name:   "get-capacity-manager-attributes",
			Fields: fields_get_capacity_manager_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCapacityManagerAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_capacity_manager_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCapacityManagerAttributes(ctx, input)
			},
		},
		"get-capacity-manager-metric-data": {
			Name:   "get-capacity-manager-metric-data",
			Fields: fields_get_capacity_manager_metric_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCapacityManagerMetricDataInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_capacity_manager_metric_data, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetCapacityManagerMetricData(ctx, input)
				}
				var results []*svc.GetCapacityManagerMetricDataOutput
				p := svc.NewGetCapacityManagerMetricDataPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-capacity-manager-metric-dimensions": {
			Name:   "get-capacity-manager-metric-dimensions",
			Fields: fields_get_capacity_manager_metric_dimensions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCapacityManagerMetricDimensionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_capacity_manager_metric_dimensions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetCapacityManagerMetricDimensions(ctx, input)
				}
				var results []*svc.GetCapacityManagerMetricDimensionsOutput
				p := svc.NewGetCapacityManagerMetricDimensionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-capacity-reservation-usage": {
			Name:   "get-capacity-reservation-usage",
			Fields: fields_get_capacity_reservation_usage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCapacityReservationUsageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_capacity_reservation_usage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCapacityReservationUsage(ctx, input)
			},
		},
		"get-coip-pool-usage": {
			Name:   "get-coip-pool-usage",
			Fields: fields_get_coip_pool_usage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCoipPoolUsageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_coip_pool_usage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCoipPoolUsage(ctx, input)
			},
		},
		"get-console-output": {
			Name:   "get-console-output",
			Fields: fields_get_console_output,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConsoleOutputInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_console_output, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConsoleOutput(ctx, input)
			},
		},
		"get-console-screenshot": {
			Name:   "get-console-screenshot",
			Fields: fields_get_console_screenshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConsoleScreenshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_console_screenshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConsoleScreenshot(ctx, input)
			},
		},
		"get-declarative-policies-report-summary": {
			Name:   "get-declarative-policies-report-summary",
			Fields: fields_get_declarative_policies_report_summary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDeclarativePoliciesReportSummaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_declarative_policies_report_summary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDeclarativePoliciesReportSummary(ctx, input)
			},
		},
		"get-default-credit-specification": {
			Name:   "get-default-credit-specification",
			Fields: fields_get_default_credit_specification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDefaultCreditSpecificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_default_credit_specification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDefaultCreditSpecification(ctx, input)
			},
		},
		"get-ebs-default-kms-key-id": {
			Name:   "get-ebs-default-kms-key-id",
			Fields: fields_get_ebs_default_kms_key_id,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEbsDefaultKmsKeyIdInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ebs_default_kms_key_id, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEbsDefaultKmsKeyId(ctx, input)
			},
		},
		"get-ebs-encryption-by-default": {
			Name:   "get-ebs-encryption-by-default",
			Fields: fields_get_ebs_encryption_by_default,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEbsEncryptionByDefaultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ebs_encryption_by_default, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEbsEncryptionByDefault(ctx, input)
			},
		},
		"get-enabled-ipam-policy": {
			Name:   "get-enabled-ipam-policy",
			Fields: fields_get_enabled_ipam_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEnabledIpamPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_enabled_ipam_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEnabledIpamPolicy(ctx, input)
			},
		},
		"get-flow-logs-integration-template": {
			Name:   "get-flow-logs-integration-template",
			Fields: fields_get_flow_logs_integration_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFlowLogsIntegrationTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_flow_logs_integration_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFlowLogsIntegrationTemplate(ctx, input)
			},
		},
		"get-groups-for-capacity-reservation": {
			Name:   "get-groups-for-capacity-reservation",
			Fields: fields_get_groups_for_capacity_reservation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGroupsForCapacityReservationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_groups_for_capacity_reservation, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetGroupsForCapacityReservation(ctx, input)
				}
				var results []*svc.GetGroupsForCapacityReservationOutput
				p := svc.NewGetGroupsForCapacityReservationPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-host-reservation-purchase-preview": {
			Name:   "get-host-reservation-purchase-preview",
			Fields: fields_get_host_reservation_purchase_preview,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetHostReservationPurchasePreviewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_host_reservation_purchase_preview, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetHostReservationPurchasePreview(ctx, input)
			},
		},
		"get-image-ancestry": {
			Name:   "get-image-ancestry",
			Fields: fields_get_image_ancestry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetImageAncestryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_image_ancestry, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetImageAncestry(ctx, input)
			},
		},
		"get-image-block-public-access-state": {
			Name:   "get-image-block-public-access-state",
			Fields: fields_get_image_block_public_access_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetImageBlockPublicAccessStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_image_block_public_access_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetImageBlockPublicAccessState(ctx, input)
			},
		},
		"get-instance-metadata-defaults": {
			Name:   "get-instance-metadata-defaults",
			Fields: fields_get_instance_metadata_defaults,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInstanceMetadataDefaultsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_instance_metadata_defaults, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInstanceMetadataDefaults(ctx, input)
			},
		},
		"get-instance-tpm-ek-pub": {
			Name:   "get-instance-tpm-ek-pub",
			Fields: fields_get_instance_tpm_ek_pub,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInstanceTpmEkPubInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_instance_tpm_ek_pub, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInstanceTpmEkPub(ctx, input)
			},
		},
		"get-instance-types-from-instance-requirements": {
			Name:   "get-instance-types-from-instance-requirements",
			Fields: fields_get_instance_types_from_instance_requirements,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInstanceTypesFromInstanceRequirementsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_instance_types_from_instance_requirements, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetInstanceTypesFromInstanceRequirements(ctx, input)
				}
				var results []*svc.GetInstanceTypesFromInstanceRequirementsOutput
				p := svc.NewGetInstanceTypesFromInstanceRequirementsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-instance-uefi-data": {
			Name:   "get-instance-uefi-data",
			Fields: fields_get_instance_uefi_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInstanceUefiDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_instance_uefi_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInstanceUefiData(ctx, input)
			},
		},
		"get-ipam-address-history": {
			Name:   "get-ipam-address-history",
			Fields: fields_get_ipam_address_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIpamAddressHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_ipam_address_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetIpamAddressHistory(ctx, input)
				}
				var results []*svc.GetIpamAddressHistoryOutput
				p := svc.NewGetIpamAddressHistoryPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-ipam-discovered-accounts": {
			Name:   "get-ipam-discovered-accounts",
			Fields: fields_get_ipam_discovered_accounts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIpamDiscoveredAccountsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_ipam_discovered_accounts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetIpamDiscoveredAccounts(ctx, input)
				}
				var results []*svc.GetIpamDiscoveredAccountsOutput
				p := svc.NewGetIpamDiscoveredAccountsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-ipam-discovered-public-addresses": {
			Name:   "get-ipam-discovered-public-addresses",
			Fields: fields_get_ipam_discovered_public_addresses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIpamDiscoveredPublicAddressesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ipam_discovered_public_addresses, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIpamDiscoveredPublicAddresses(ctx, input)
			},
		},
		"get-ipam-discovered-resource-cidrs": {
			Name:   "get-ipam-discovered-resource-cidrs",
			Fields: fields_get_ipam_discovered_resource_cidrs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIpamDiscoveredResourceCidrsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_ipam_discovered_resource_cidrs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetIpamDiscoveredResourceCidrs(ctx, input)
				}
				var results []*svc.GetIpamDiscoveredResourceCidrsOutput
				p := svc.NewGetIpamDiscoveredResourceCidrsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-ipam-policy-allocation-rules": {
			Name:   "get-ipam-policy-allocation-rules",
			Fields: fields_get_ipam_policy_allocation_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIpamPolicyAllocationRulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ipam_policy_allocation_rules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIpamPolicyAllocationRules(ctx, input)
			},
		},
		"get-ipam-policy-organization-targets": {
			Name:   "get-ipam-policy-organization-targets",
			Fields: fields_get_ipam_policy_organization_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIpamPolicyOrganizationTargetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ipam_policy_organization_targets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIpamPolicyOrganizationTargets(ctx, input)
			},
		},
		"get-ipam-pool-allocations": {
			Name:   "get-ipam-pool-allocations",
			Fields: fields_get_ipam_pool_allocations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIpamPoolAllocationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_ipam_pool_allocations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetIpamPoolAllocations(ctx, input)
				}
				var results []*svc.GetIpamPoolAllocationsOutput
				p := svc.NewGetIpamPoolAllocationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-ipam-pool-cidrs": {
			Name:   "get-ipam-pool-cidrs",
			Fields: fields_get_ipam_pool_cidrs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIpamPoolCidrsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_ipam_pool_cidrs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetIpamPoolCidrs(ctx, input)
				}
				var results []*svc.GetIpamPoolCidrsOutput
				p := svc.NewGetIpamPoolCidrsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-ipam-prefix-list-resolver-rules": {
			Name:   "get-ipam-prefix-list-resolver-rules",
			Fields: fields_get_ipam_prefix_list_resolver_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIpamPrefixListResolverRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_ipam_prefix_list_resolver_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetIpamPrefixListResolverRules(ctx, input)
				}
				var results []*svc.GetIpamPrefixListResolverRulesOutput
				p := svc.NewGetIpamPrefixListResolverRulesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-ipam-prefix-list-resolver-version-entries": {
			Name:   "get-ipam-prefix-list-resolver-version-entries",
			Fields: fields_get_ipam_prefix_list_resolver_version_entries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIpamPrefixListResolverVersionEntriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_ipam_prefix_list_resolver_version_entries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetIpamPrefixListResolverVersionEntries(ctx, input)
				}
				var results []*svc.GetIpamPrefixListResolverVersionEntriesOutput
				p := svc.NewGetIpamPrefixListResolverVersionEntriesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-ipam-prefix-list-resolver-versions": {
			Name:   "get-ipam-prefix-list-resolver-versions",
			Fields: fields_get_ipam_prefix_list_resolver_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIpamPrefixListResolverVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_ipam_prefix_list_resolver_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetIpamPrefixListResolverVersions(ctx, input)
				}
				var results []*svc.GetIpamPrefixListResolverVersionsOutput
				p := svc.NewGetIpamPrefixListResolverVersionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-ipam-resource-cidrs": {
			Name:   "get-ipam-resource-cidrs",
			Fields: fields_get_ipam_resource_cidrs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIpamResourceCidrsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_ipam_resource_cidrs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetIpamResourceCidrs(ctx, input)
				}
				var results []*svc.GetIpamResourceCidrsOutput
				p := svc.NewGetIpamResourceCidrsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-launch-template-data": {
			Name:   "get-launch-template-data",
			Fields: fields_get_launch_template_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLaunchTemplateDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_launch_template_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLaunchTemplateData(ctx, input)
			},
		},
		"get-managed-prefix-list-associations": {
			Name:   "get-managed-prefix-list-associations",
			Fields: fields_get_managed_prefix_list_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetManagedPrefixListAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_managed_prefix_list_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetManagedPrefixListAssociations(ctx, input)
				}
				var results []*svc.GetManagedPrefixListAssociationsOutput
				p := svc.NewGetManagedPrefixListAssociationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-managed-prefix-list-entries": {
			Name:   "get-managed-prefix-list-entries",
			Fields: fields_get_managed_prefix_list_entries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetManagedPrefixListEntriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_managed_prefix_list_entries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetManagedPrefixListEntries(ctx, input)
				}
				var results []*svc.GetManagedPrefixListEntriesOutput
				p := svc.NewGetManagedPrefixListEntriesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-network-insights-access-scope-analysis-findings": {
			Name:   "get-network-insights-access-scope-analysis-findings",
			Fields: fields_get_network_insights_access_scope_analysis_findings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetNetworkInsightsAccessScopeAnalysisFindingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_network_insights_access_scope_analysis_findings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetNetworkInsightsAccessScopeAnalysisFindings(ctx, input)
				}
				var results []*svc.GetNetworkInsightsAccessScopeAnalysisFindingsOutput
				p := svc.NewGetNetworkInsightsAccessScopeAnalysisFindingsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-network-insights-access-scope-content": {
			Name:   "get-network-insights-access-scope-content",
			Fields: fields_get_network_insights_access_scope_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetNetworkInsightsAccessScopeContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_network_insights_access_scope_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetNetworkInsightsAccessScopeContent(ctx, input)
			},
		},
		"get-password-data": {
			Name:   "get-password-data",
			Fields: fields_get_password_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPasswordDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_password_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPasswordData(ctx, input)
			},
		},
		"get-reserved-instances-exchange-quote": {
			Name:   "get-reserved-instances-exchange-quote",
			Fields: fields_get_reserved_instances_exchange_quote,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReservedInstancesExchangeQuoteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_reserved_instances_exchange_quote, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetReservedInstancesExchangeQuote(ctx, input)
			},
		},
		"get-route-server-associations": {
			Name:   "get-route-server-associations",
			Fields: fields_get_route_server_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRouteServerAssociationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_route_server_associations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRouteServerAssociations(ctx, input)
			},
		},
		"get-route-server-propagations": {
			Name:   "get-route-server-propagations",
			Fields: fields_get_route_server_propagations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRouteServerPropagationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_route_server_propagations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRouteServerPropagations(ctx, input)
			},
		},
		"get-route-server-routing-database": {
			Name:   "get-route-server-routing-database",
			Fields: fields_get_route_server_routing_database,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRouteServerRoutingDatabaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_route_server_routing_database, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRouteServerRoutingDatabase(ctx, input)
			},
		},
		"get-security-groups-for-vpc": {
			Name:   "get-security-groups-for-vpc",
			Fields: fields_get_security_groups_for_vpc,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSecurityGroupsForVpcInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_security_groups_for_vpc, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetSecurityGroupsForVpc(ctx, input)
				}
				var results []*svc.GetSecurityGroupsForVpcOutput
				p := svc.NewGetSecurityGroupsForVpcPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-serial-console-access-status": {
			Name:   "get-serial-console-access-status",
			Fields: fields_get_serial_console_access_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSerialConsoleAccessStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_serial_console_access_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSerialConsoleAccessStatus(ctx, input)
			},
		},
		"get-snapshot-block-public-access-state": {
			Name:   "get-snapshot-block-public-access-state",
			Fields: fields_get_snapshot_block_public_access_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSnapshotBlockPublicAccessStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_snapshot_block_public_access_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSnapshotBlockPublicAccessState(ctx, input)
			},
		},
		"get-spot-placement-scores": {
			Name:   "get-spot-placement-scores",
			Fields: fields_get_spot_placement_scores,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSpotPlacementScoresInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_spot_placement_scores, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetSpotPlacementScores(ctx, input)
				}
				var results []*svc.GetSpotPlacementScoresOutput
				p := svc.NewGetSpotPlacementScoresPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-subnet-cidr-reservations": {
			Name:   "get-subnet-cidr-reservations",
			Fields: fields_get_subnet_cidr_reservations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSubnetCidrReservationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_subnet_cidr_reservations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSubnetCidrReservations(ctx, input)
			},
		},
		"get-transit-gateway-attachment-propagations": {
			Name:   "get-transit-gateway-attachment-propagations",
			Fields: fields_get_transit_gateway_attachment_propagations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTransitGatewayAttachmentPropagationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_transit_gateway_attachment_propagations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetTransitGatewayAttachmentPropagations(ctx, input)
				}
				var results []*svc.GetTransitGatewayAttachmentPropagationsOutput
				p := svc.NewGetTransitGatewayAttachmentPropagationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-transit-gateway-metering-policy-entries": {
			Name:   "get-transit-gateway-metering-policy-entries",
			Fields: fields_get_transit_gateway_metering_policy_entries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTransitGatewayMeteringPolicyEntriesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_transit_gateway_metering_policy_entries, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTransitGatewayMeteringPolicyEntries(ctx, input)
			},
		},
		"get-transit-gateway-multicast-domain-associations": {
			Name:   "get-transit-gateway-multicast-domain-associations",
			Fields: fields_get_transit_gateway_multicast_domain_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTransitGatewayMulticastDomainAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_transit_gateway_multicast_domain_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetTransitGatewayMulticastDomainAssociations(ctx, input)
				}
				var results []*svc.GetTransitGatewayMulticastDomainAssociationsOutput
				p := svc.NewGetTransitGatewayMulticastDomainAssociationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-transit-gateway-policy-table-associations": {
			Name:   "get-transit-gateway-policy-table-associations",
			Fields: fields_get_transit_gateway_policy_table_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTransitGatewayPolicyTableAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_transit_gateway_policy_table_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetTransitGatewayPolicyTableAssociations(ctx, input)
				}
				var results []*svc.GetTransitGatewayPolicyTableAssociationsOutput
				p := svc.NewGetTransitGatewayPolicyTableAssociationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-transit-gateway-policy-table-entries": {
			Name:   "get-transit-gateway-policy-table-entries",
			Fields: fields_get_transit_gateway_policy_table_entries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTransitGatewayPolicyTableEntriesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_transit_gateway_policy_table_entries, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTransitGatewayPolicyTableEntries(ctx, input)
			},
		},
		"get-transit-gateway-prefix-list-references": {
			Name:   "get-transit-gateway-prefix-list-references",
			Fields: fields_get_transit_gateway_prefix_list_references,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTransitGatewayPrefixListReferencesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_transit_gateway_prefix_list_references, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetTransitGatewayPrefixListReferences(ctx, input)
				}
				var results []*svc.GetTransitGatewayPrefixListReferencesOutput
				p := svc.NewGetTransitGatewayPrefixListReferencesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-transit-gateway-route-table-associations": {
			Name:   "get-transit-gateway-route-table-associations",
			Fields: fields_get_transit_gateway_route_table_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTransitGatewayRouteTableAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_transit_gateway_route_table_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetTransitGatewayRouteTableAssociations(ctx, input)
				}
				var results []*svc.GetTransitGatewayRouteTableAssociationsOutput
				p := svc.NewGetTransitGatewayRouteTableAssociationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-transit-gateway-route-table-propagations": {
			Name:   "get-transit-gateway-route-table-propagations",
			Fields: fields_get_transit_gateway_route_table_propagations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTransitGatewayRouteTablePropagationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_transit_gateway_route_table_propagations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetTransitGatewayRouteTablePropagations(ctx, input)
				}
				var results []*svc.GetTransitGatewayRouteTablePropagationsOutput
				p := svc.NewGetTransitGatewayRouteTablePropagationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-verified-access-endpoint-policy": {
			Name:   "get-verified-access-endpoint-policy",
			Fields: fields_get_verified_access_endpoint_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVerifiedAccessEndpointPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_verified_access_endpoint_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVerifiedAccessEndpointPolicy(ctx, input)
			},
		},
		"get-verified-access-endpoint-targets": {
			Name:   "get-verified-access-endpoint-targets",
			Fields: fields_get_verified_access_endpoint_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVerifiedAccessEndpointTargetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_verified_access_endpoint_targets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVerifiedAccessEndpointTargets(ctx, input)
			},
		},
		"get-verified-access-group-policy": {
			Name:   "get-verified-access-group-policy",
			Fields: fields_get_verified_access_group_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVerifiedAccessGroupPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_verified_access_group_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVerifiedAccessGroupPolicy(ctx, input)
			},
		},
		"get-vpc-resources-blocking-encryption-enforcement": {
			Name:   "get-vpc-resources-blocking-encryption-enforcement",
			Fields: fields_get_vpc_resources_blocking_encryption_enforcement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVpcResourcesBlockingEncryptionEnforcementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_vpc_resources_blocking_encryption_enforcement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVpcResourcesBlockingEncryptionEnforcement(ctx, input)
			},
		},
		"get-vpn-connection-device-sample-configuration": {
			Name:   "get-vpn-connection-device-sample-configuration",
			Fields: fields_get_vpn_connection_device_sample_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVpnConnectionDeviceSampleConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_vpn_connection_device_sample_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVpnConnectionDeviceSampleConfiguration(ctx, input)
			},
		},
		"get-vpn-connection-device-types": {
			Name:   "get-vpn-connection-device-types",
			Fields: fields_get_vpn_connection_device_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVpnConnectionDeviceTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_vpn_connection_device_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetVpnConnectionDeviceTypes(ctx, input)
				}
				var results []*svc.GetVpnConnectionDeviceTypesOutput
				p := svc.NewGetVpnConnectionDeviceTypesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-vpn-tunnel-replacement-status": {
			Name:   "get-vpn-tunnel-replacement-status",
			Fields: fields_get_vpn_tunnel_replacement_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVpnTunnelReplacementStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_vpn_tunnel_replacement_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVpnTunnelReplacementStatus(ctx, input)
			},
		},
		"import-client-vpn-client-certificate-revocation-list": {
			Name:   "import-client-vpn-client-certificate-revocation-list",
			Fields: fields_import_client_vpn_client_certificate_revocation_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportClientVpnClientCertificateRevocationListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_client_vpn_client_certificate_revocation_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportClientVpnClientCertificateRevocationList(ctx, input)
			},
		},
		"import-image": {
			Name:   "import-image",
			Fields: fields_import_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportImage(ctx, input)
			},
		},
		"import-instance": {
			Name:   "import-instance",
			Fields: fields_import_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportInstance(ctx, input)
			},
		},
		"import-key-pair": {
			Name:   "import-key-pair",
			Fields: fields_import_key_pair,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportKeyPairInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_key_pair, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportKeyPair(ctx, input)
			},
		},
		"import-snapshot": {
			Name:   "import-snapshot",
			Fields: fields_import_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportSnapshot(ctx, input)
			},
		},
		"import-volume": {
			Name:   "import-volume",
			Fields: fields_import_volume,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportVolumeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_volume, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportVolume(ctx, input)
			},
		},
		"list-images-in-recycle-bin": {
			Name:   "list-images-in-recycle-bin",
			Fields: fields_list_images_in_recycle_bin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListImagesInRecycleBinInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_images_in_recycle_bin, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListImagesInRecycleBin(ctx, input)
				}
				var results []*svc.ListImagesInRecycleBinOutput
				p := svc.NewListImagesInRecycleBinPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-snapshots-in-recycle-bin": {
			Name:   "list-snapshots-in-recycle-bin",
			Fields: fields_list_snapshots_in_recycle_bin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSnapshotsInRecycleBinInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_snapshots_in_recycle_bin, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSnapshotsInRecycleBin(ctx, input)
				}
				var results []*svc.ListSnapshotsInRecycleBinOutput
				p := svc.NewListSnapshotsInRecycleBinPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-volumes-in-recycle-bin": {
			Name:   "list-volumes-in-recycle-bin",
			Fields: fields_list_volumes_in_recycle_bin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVolumesInRecycleBinInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_volumes_in_recycle_bin, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListVolumesInRecycleBin(ctx, input)
			},
		},
		"lock-snapshot": {
			Name:   "lock-snapshot",
			Fields: fields_lock_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.LockSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_lock_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.LockSnapshot(ctx, input)
			},
		},
		"modify-address-attribute": {
			Name:   "modify-address-attribute",
			Fields: fields_modify_address_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyAddressAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_address_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyAddressAttribute(ctx, input)
			},
		},
		"modify-availability-zone-group": {
			Name:   "modify-availability-zone-group",
			Fields: fields_modify_availability_zone_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyAvailabilityZoneGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_availability_zone_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyAvailabilityZoneGroup(ctx, input)
			},
		},
		"modify-capacity-reservation": {
			Name:   "modify-capacity-reservation",
			Fields: fields_modify_capacity_reservation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyCapacityReservationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_capacity_reservation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyCapacityReservation(ctx, input)
			},
		},
		"modify-capacity-reservation-fleet": {
			Name:   "modify-capacity-reservation-fleet",
			Fields: fields_modify_capacity_reservation_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyCapacityReservationFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_capacity_reservation_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyCapacityReservationFleet(ctx, input)
			},
		},
		"modify-client-vpn-endpoint": {
			Name:   "modify-client-vpn-endpoint",
			Fields: fields_modify_client_vpn_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyClientVpnEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_client_vpn_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyClientVpnEndpoint(ctx, input)
			},
		},
		"modify-default-credit-specification": {
			Name:   "modify-default-credit-specification",
			Fields: fields_modify_default_credit_specification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyDefaultCreditSpecificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_default_credit_specification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyDefaultCreditSpecification(ctx, input)
			},
		},
		"modify-ebs-default-kms-key-id": {
			Name:   "modify-ebs-default-kms-key-id",
			Fields: fields_modify_ebs_default_kms_key_id,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyEbsDefaultKmsKeyIdInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_ebs_default_kms_key_id, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyEbsDefaultKmsKeyId(ctx, input)
			},
		},
		"modify-fleet": {
			Name:   "modify-fleet",
			Fields: fields_modify_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyFleet(ctx, input)
			},
		},
		"modify-fpga-image-attribute": {
			Name:   "modify-fpga-image-attribute",
			Fields: fields_modify_fpga_image_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyFpgaImageAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_fpga_image_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyFpgaImageAttribute(ctx, input)
			},
		},
		"modify-hosts": {
			Name:   "modify-hosts",
			Fields: fields_modify_hosts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyHostsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_hosts, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyHosts(ctx, input)
			},
		},
		"modify-id-format": {
			Name:   "modify-id-format",
			Fields: fields_modify_id_format,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyIdFormatInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_id_format, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyIdFormat(ctx, input)
			},
		},
		"modify-identity-id-format": {
			Name:   "modify-identity-id-format",
			Fields: fields_modify_identity_id_format,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyIdentityIdFormatInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_identity_id_format, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyIdentityIdFormat(ctx, input)
			},
		},
		"modify-image-attribute": {
			Name:   "modify-image-attribute",
			Fields: fields_modify_image_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyImageAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_image_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyImageAttribute(ctx, input)
			},
		},
		"modify-instance-attribute": {
			Name:   "modify-instance-attribute",
			Fields: fields_modify_instance_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyInstanceAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_instance_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyInstanceAttribute(ctx, input)
			},
		},
		"modify-instance-capacity-reservation-attributes": {
			Name:   "modify-instance-capacity-reservation-attributes",
			Fields: fields_modify_instance_capacity_reservation_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyInstanceCapacityReservationAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_instance_capacity_reservation_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyInstanceCapacityReservationAttributes(ctx, input)
			},
		},
		"modify-instance-connect-endpoint": {
			Name:   "modify-instance-connect-endpoint",
			Fields: fields_modify_instance_connect_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyInstanceConnectEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_instance_connect_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyInstanceConnectEndpoint(ctx, input)
			},
		},
		"modify-instance-cpu-options": {
			Name:   "modify-instance-cpu-options",
			Fields: fields_modify_instance_cpu_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyInstanceCpuOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_instance_cpu_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyInstanceCpuOptions(ctx, input)
			},
		},
		"modify-instance-credit-specification": {
			Name:   "modify-instance-credit-specification",
			Fields: fields_modify_instance_credit_specification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyInstanceCreditSpecificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_instance_credit_specification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyInstanceCreditSpecification(ctx, input)
			},
		},
		"modify-instance-event-start-time": {
			Name:   "modify-instance-event-start-time",
			Fields: fields_modify_instance_event_start_time,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyInstanceEventStartTimeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_instance_event_start_time, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyInstanceEventStartTime(ctx, input)
			},
		},
		"modify-instance-event-window": {
			Name:   "modify-instance-event-window",
			Fields: fields_modify_instance_event_window,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyInstanceEventWindowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_instance_event_window, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyInstanceEventWindow(ctx, input)
			},
		},
		"modify-instance-maintenance-options": {
			Name:   "modify-instance-maintenance-options",
			Fields: fields_modify_instance_maintenance_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyInstanceMaintenanceOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_instance_maintenance_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyInstanceMaintenanceOptions(ctx, input)
			},
		},
		"modify-instance-metadata-defaults": {
			Name:   "modify-instance-metadata-defaults",
			Fields: fields_modify_instance_metadata_defaults,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyInstanceMetadataDefaultsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_instance_metadata_defaults, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyInstanceMetadataDefaults(ctx, input)
			},
		},
		"modify-instance-metadata-options": {
			Name:   "modify-instance-metadata-options",
			Fields: fields_modify_instance_metadata_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyInstanceMetadataOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_instance_metadata_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyInstanceMetadataOptions(ctx, input)
			},
		},
		"modify-instance-network-performance-options": {
			Name:   "modify-instance-network-performance-options",
			Fields: fields_modify_instance_network_performance_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyInstanceNetworkPerformanceOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_instance_network_performance_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyInstanceNetworkPerformanceOptions(ctx, input)
			},
		},
		"modify-instance-placement": {
			Name:   "modify-instance-placement",
			Fields: fields_modify_instance_placement,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyInstancePlacementInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_instance_placement, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyInstancePlacement(ctx, input)
			},
		},
		"modify-ipam": {
			Name:   "modify-ipam",
			Fields: fields_modify_ipam,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyIpamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_ipam, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyIpam(ctx, input)
			},
		},
		"modify-ipam-policy-allocation-rules": {
			Name:   "modify-ipam-policy-allocation-rules",
			Fields: fields_modify_ipam_policy_allocation_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyIpamPolicyAllocationRulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_ipam_policy_allocation_rules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyIpamPolicyAllocationRules(ctx, input)
			},
		},
		"modify-ipam-pool": {
			Name:   "modify-ipam-pool",
			Fields: fields_modify_ipam_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyIpamPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_ipam_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyIpamPool(ctx, input)
			},
		},
		"modify-ipam-prefix-list-resolver": {
			Name:   "modify-ipam-prefix-list-resolver",
			Fields: fields_modify_ipam_prefix_list_resolver,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyIpamPrefixListResolverInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_ipam_prefix_list_resolver, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyIpamPrefixListResolver(ctx, input)
			},
		},
		"modify-ipam-prefix-list-resolver-target": {
			Name:   "modify-ipam-prefix-list-resolver-target",
			Fields: fields_modify_ipam_prefix_list_resolver_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyIpamPrefixListResolverTargetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_ipam_prefix_list_resolver_target, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyIpamPrefixListResolverTarget(ctx, input)
			},
		},
		"modify-ipam-resource-cidr": {
			Name:   "modify-ipam-resource-cidr",
			Fields: fields_modify_ipam_resource_cidr,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyIpamResourceCidrInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_ipam_resource_cidr, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyIpamResourceCidr(ctx, input)
			},
		},
		"modify-ipam-resource-discovery": {
			Name:   "modify-ipam-resource-discovery",
			Fields: fields_modify_ipam_resource_discovery,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyIpamResourceDiscoveryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_ipam_resource_discovery, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyIpamResourceDiscovery(ctx, input)
			},
		},
		"modify-ipam-scope": {
			Name:   "modify-ipam-scope",
			Fields: fields_modify_ipam_scope,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyIpamScopeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_ipam_scope, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyIpamScope(ctx, input)
			},
		},
		"modify-launch-template": {
			Name:   "modify-launch-template",
			Fields: fields_modify_launch_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyLaunchTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_launch_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyLaunchTemplate(ctx, input)
			},
		},
		"modify-local-gateway-route": {
			Name:   "modify-local-gateway-route",
			Fields: fields_modify_local_gateway_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyLocalGatewayRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_local_gateway_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyLocalGatewayRoute(ctx, input)
			},
		},
		"modify-managed-prefix-list": {
			Name:   "modify-managed-prefix-list",
			Fields: fields_modify_managed_prefix_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyManagedPrefixListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_managed_prefix_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyManagedPrefixList(ctx, input)
			},
		},
		"modify-network-interface-attribute": {
			Name:   "modify-network-interface-attribute",
			Fields: fields_modify_network_interface_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyNetworkInterfaceAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_network_interface_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyNetworkInterfaceAttribute(ctx, input)
			},
		},
		"modify-private-dns-name-options": {
			Name:   "modify-private-dns-name-options",
			Fields: fields_modify_private_dns_name_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyPrivateDnsNameOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_private_dns_name_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyPrivateDnsNameOptions(ctx, input)
			},
		},
		"modify-public-ip-dns-name-options": {
			Name:   "modify-public-ip-dns-name-options",
			Fields: fields_modify_public_ip_dns_name_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyPublicIpDnsNameOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_public_ip_dns_name_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyPublicIpDnsNameOptions(ctx, input)
			},
		},
		"modify-reserved-instances": {
			Name:   "modify-reserved-instances",
			Fields: fields_modify_reserved_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyReservedInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_reserved_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyReservedInstances(ctx, input)
			},
		},
		"modify-route-server": {
			Name:   "modify-route-server",
			Fields: fields_modify_route_server,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyRouteServerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_route_server, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyRouteServer(ctx, input)
			},
		},
		"modify-security-group-rules": {
			Name:   "modify-security-group-rules",
			Fields: fields_modify_security_group_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifySecurityGroupRulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_security_group_rules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifySecurityGroupRules(ctx, input)
			},
		},
		"modify-snapshot-attribute": {
			Name:   "modify-snapshot-attribute",
			Fields: fields_modify_snapshot_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifySnapshotAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_snapshot_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifySnapshotAttribute(ctx, input)
			},
		},
		"modify-snapshot-tier": {
			Name:   "modify-snapshot-tier",
			Fields: fields_modify_snapshot_tier,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifySnapshotTierInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_snapshot_tier, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifySnapshotTier(ctx, input)
			},
		},
		"modify-spot-fleet-request": {
			Name:   "modify-spot-fleet-request",
			Fields: fields_modify_spot_fleet_request,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifySpotFleetRequestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_spot_fleet_request, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifySpotFleetRequest(ctx, input)
			},
		},
		"modify-subnet-attribute": {
			Name:   "modify-subnet-attribute",
			Fields: fields_modify_subnet_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifySubnetAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_subnet_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifySubnetAttribute(ctx, input)
			},
		},
		"modify-traffic-mirror-filter-network-services": {
			Name:   "modify-traffic-mirror-filter-network-services",
			Fields: fields_modify_traffic_mirror_filter_network_services,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyTrafficMirrorFilterNetworkServicesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_traffic_mirror_filter_network_services, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyTrafficMirrorFilterNetworkServices(ctx, input)
			},
		},
		"modify-traffic-mirror-filter-rule": {
			Name:   "modify-traffic-mirror-filter-rule",
			Fields: fields_modify_traffic_mirror_filter_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyTrafficMirrorFilterRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_traffic_mirror_filter_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyTrafficMirrorFilterRule(ctx, input)
			},
		},
		"modify-traffic-mirror-session": {
			Name:   "modify-traffic-mirror-session",
			Fields: fields_modify_traffic_mirror_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyTrafficMirrorSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_traffic_mirror_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyTrafficMirrorSession(ctx, input)
			},
		},
		"modify-transit-gateway": {
			Name:   "modify-transit-gateway",
			Fields: fields_modify_transit_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyTransitGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_transit_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyTransitGateway(ctx, input)
			},
		},
		"modify-transit-gateway-metering-policy": {
			Name:   "modify-transit-gateway-metering-policy",
			Fields: fields_modify_transit_gateway_metering_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyTransitGatewayMeteringPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_transit_gateway_metering_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyTransitGatewayMeteringPolicy(ctx, input)
			},
		},
		"modify-transit-gateway-prefix-list-reference": {
			Name:   "modify-transit-gateway-prefix-list-reference",
			Fields: fields_modify_transit_gateway_prefix_list_reference,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyTransitGatewayPrefixListReferenceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_transit_gateway_prefix_list_reference, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyTransitGatewayPrefixListReference(ctx, input)
			},
		},
		"modify-transit-gateway-vpc-attachment": {
			Name:   "modify-transit-gateway-vpc-attachment",
			Fields: fields_modify_transit_gateway_vpc_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyTransitGatewayVpcAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_transit_gateway_vpc_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyTransitGatewayVpcAttachment(ctx, input)
			},
		},
		"modify-verified-access-endpoint": {
			Name:   "modify-verified-access-endpoint",
			Fields: fields_modify_verified_access_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyVerifiedAccessEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_verified_access_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyVerifiedAccessEndpoint(ctx, input)
			},
		},
		"modify-verified-access-endpoint-policy": {
			Name:   "modify-verified-access-endpoint-policy",
			Fields: fields_modify_verified_access_endpoint_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyVerifiedAccessEndpointPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_verified_access_endpoint_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyVerifiedAccessEndpointPolicy(ctx, input)
			},
		},
		"modify-verified-access-group": {
			Name:   "modify-verified-access-group",
			Fields: fields_modify_verified_access_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyVerifiedAccessGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_verified_access_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyVerifiedAccessGroup(ctx, input)
			},
		},
		"modify-verified-access-group-policy": {
			Name:   "modify-verified-access-group-policy",
			Fields: fields_modify_verified_access_group_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyVerifiedAccessGroupPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_verified_access_group_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyVerifiedAccessGroupPolicy(ctx, input)
			},
		},
		"modify-verified-access-instance": {
			Name:   "modify-verified-access-instance",
			Fields: fields_modify_verified_access_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyVerifiedAccessInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_verified_access_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyVerifiedAccessInstance(ctx, input)
			},
		},
		"modify-verified-access-instance-logging-configuration": {
			Name:   "modify-verified-access-instance-logging-configuration",
			Fields: fields_modify_verified_access_instance_logging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyVerifiedAccessInstanceLoggingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_verified_access_instance_logging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyVerifiedAccessInstanceLoggingConfiguration(ctx, input)
			},
		},
		"modify-verified-access-trust-provider": {
			Name:   "modify-verified-access-trust-provider",
			Fields: fields_modify_verified_access_trust_provider,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyVerifiedAccessTrustProviderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_verified_access_trust_provider, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyVerifiedAccessTrustProvider(ctx, input)
			},
		},
		"modify-volume": {
			Name:   "modify-volume",
			Fields: fields_modify_volume,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyVolumeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_volume, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyVolume(ctx, input)
			},
		},
		"modify-volume-attribute": {
			Name:   "modify-volume-attribute",
			Fields: fields_modify_volume_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyVolumeAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_volume_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyVolumeAttribute(ctx, input)
			},
		},
		"modify-vpc-attribute": {
			Name:   "modify-vpc-attribute",
			Fields: fields_modify_vpc_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyVpcAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_vpc_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyVpcAttribute(ctx, input)
			},
		},
		"modify-vpc-block-public-access-exclusion": {
			Name:   "modify-vpc-block-public-access-exclusion",
			Fields: fields_modify_vpc_block_public_access_exclusion,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyVpcBlockPublicAccessExclusionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_vpc_block_public_access_exclusion, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyVpcBlockPublicAccessExclusion(ctx, input)
			},
		},
		"modify-vpc-block-public-access-options": {
			Name:   "modify-vpc-block-public-access-options",
			Fields: fields_modify_vpc_block_public_access_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyVpcBlockPublicAccessOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_vpc_block_public_access_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyVpcBlockPublicAccessOptions(ctx, input)
			},
		},
		"modify-vpc-encryption-control": {
			Name:   "modify-vpc-encryption-control",
			Fields: fields_modify_vpc_encryption_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyVpcEncryptionControlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_vpc_encryption_control, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyVpcEncryptionControl(ctx, input)
			},
		},
		"modify-vpc-endpoint": {
			Name:   "modify-vpc-endpoint",
			Fields: fields_modify_vpc_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyVpcEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_vpc_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyVpcEndpoint(ctx, input)
			},
		},
		"modify-vpc-endpoint-connection-notification": {
			Name:   "modify-vpc-endpoint-connection-notification",
			Fields: fields_modify_vpc_endpoint_connection_notification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyVpcEndpointConnectionNotificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_vpc_endpoint_connection_notification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyVpcEndpointConnectionNotification(ctx, input)
			},
		},
		"modify-vpc-endpoint-service-configuration": {
			Name:   "modify-vpc-endpoint-service-configuration",
			Fields: fields_modify_vpc_endpoint_service_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyVpcEndpointServiceConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_vpc_endpoint_service_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyVpcEndpointServiceConfiguration(ctx, input)
			},
		},
		"modify-vpc-endpoint-service-payer-responsibility": {
			Name:   "modify-vpc-endpoint-service-payer-responsibility",
			Fields: fields_modify_vpc_endpoint_service_payer_responsibility,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyVpcEndpointServicePayerResponsibilityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_vpc_endpoint_service_payer_responsibility, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyVpcEndpointServicePayerResponsibility(ctx, input)
			},
		},
		"modify-vpc-endpoint-service-permissions": {
			Name:   "modify-vpc-endpoint-service-permissions",
			Fields: fields_modify_vpc_endpoint_service_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyVpcEndpointServicePermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_vpc_endpoint_service_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyVpcEndpointServicePermissions(ctx, input)
			},
		},
		"modify-vpc-peering-connection-options": {
			Name:   "modify-vpc-peering-connection-options",
			Fields: fields_modify_vpc_peering_connection_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyVpcPeeringConnectionOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_vpc_peering_connection_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyVpcPeeringConnectionOptions(ctx, input)
			},
		},
		"modify-vpc-tenancy": {
			Name:   "modify-vpc-tenancy",
			Fields: fields_modify_vpc_tenancy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyVpcTenancyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_vpc_tenancy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyVpcTenancy(ctx, input)
			},
		},
		"modify-vpn-connection": {
			Name:   "modify-vpn-connection",
			Fields: fields_modify_vpn_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyVpnConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_vpn_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyVpnConnection(ctx, input)
			},
		},
		"modify-vpn-connection-options": {
			Name:   "modify-vpn-connection-options",
			Fields: fields_modify_vpn_connection_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyVpnConnectionOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_vpn_connection_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyVpnConnectionOptions(ctx, input)
			},
		},
		"modify-vpn-tunnel-certificate": {
			Name:   "modify-vpn-tunnel-certificate",
			Fields: fields_modify_vpn_tunnel_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyVpnTunnelCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_vpn_tunnel_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyVpnTunnelCertificate(ctx, input)
			},
		},
		"modify-vpn-tunnel-options": {
			Name:   "modify-vpn-tunnel-options",
			Fields: fields_modify_vpn_tunnel_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ModifyVpnTunnelOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_modify_vpn_tunnel_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ModifyVpnTunnelOptions(ctx, input)
			},
		},
		"monitor-instances": {
			Name:   "monitor-instances",
			Fields: fields_monitor_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.MonitorInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_monitor_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.MonitorInstances(ctx, input)
			},
		},
		"move-address-to-vpc": {
			Name:   "move-address-to-vpc",
			Fields: fields_move_address_to_vpc,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.MoveAddressToVpcInput{}
				if _, err := leanruntime.ApplyInput(input, fields_move_address_to_vpc, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.MoveAddressToVpc(ctx, input)
			},
		},
		"move-byoip-cidr-to-ipam": {
			Name:   "move-byoip-cidr-to-ipam",
			Fields: fields_move_byoip_cidr_to_ipam,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.MoveByoipCidrToIpamInput{}
				if _, err := leanruntime.ApplyInput(input, fields_move_byoip_cidr_to_ipam, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.MoveByoipCidrToIpam(ctx, input)
			},
		},
		"move-capacity-reservation-instances": {
			Name:   "move-capacity-reservation-instances",
			Fields: fields_move_capacity_reservation_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.MoveCapacityReservationInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_move_capacity_reservation_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.MoveCapacityReservationInstances(ctx, input)
			},
		},
		"provision-byoip-cidr": {
			Name:   "provision-byoip-cidr",
			Fields: fields_provision_byoip_cidr,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ProvisionByoipCidrInput{}
				if _, err := leanruntime.ApplyInput(input, fields_provision_byoip_cidr, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ProvisionByoipCidr(ctx, input)
			},
		},
		"provision-ipam-byoasn": {
			Name:   "provision-ipam-byoasn",
			Fields: fields_provision_ipam_byoasn,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ProvisionIpamByoasnInput{}
				if _, err := leanruntime.ApplyInput(input, fields_provision_ipam_byoasn, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ProvisionIpamByoasn(ctx, input)
			},
		},
		"provision-ipam-pool-cidr": {
			Name:   "provision-ipam-pool-cidr",
			Fields: fields_provision_ipam_pool_cidr,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ProvisionIpamPoolCidrInput{}
				if _, err := leanruntime.ApplyInput(input, fields_provision_ipam_pool_cidr, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ProvisionIpamPoolCidr(ctx, input)
			},
		},
		"provision-public-ipv4-pool-cidr": {
			Name:   "provision-public-ipv4-pool-cidr",
			Fields: fields_provision_public_ipv4_pool_cidr,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ProvisionPublicIpv4PoolCidrInput{}
				if _, err := leanruntime.ApplyInput(input, fields_provision_public_ipv4_pool_cidr, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ProvisionPublicIpv4PoolCidr(ctx, input)
			},
		},
		"purchase-capacity-block": {
			Name:   "purchase-capacity-block",
			Fields: fields_purchase_capacity_block,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PurchaseCapacityBlockInput{}
				if _, err := leanruntime.ApplyInput(input, fields_purchase_capacity_block, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PurchaseCapacityBlock(ctx, input)
			},
		},
		"purchase-capacity-block-extension": {
			Name:   "purchase-capacity-block-extension",
			Fields: fields_purchase_capacity_block_extension,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PurchaseCapacityBlockExtensionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_purchase_capacity_block_extension, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PurchaseCapacityBlockExtension(ctx, input)
			},
		},
		"purchase-host-reservation": {
			Name:   "purchase-host-reservation",
			Fields: fields_purchase_host_reservation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PurchaseHostReservationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_purchase_host_reservation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PurchaseHostReservation(ctx, input)
			},
		},
		"purchase-reserved-instances-offering": {
			Name:   "purchase-reserved-instances-offering",
			Fields: fields_purchase_reserved_instances_offering,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PurchaseReservedInstancesOfferingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_purchase_reserved_instances_offering, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PurchaseReservedInstancesOffering(ctx, input)
			},
		},
		"purchase-scheduled-instances": {
			Name:   "purchase-scheduled-instances",
			Fields: fields_purchase_scheduled_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PurchaseScheduledInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_purchase_scheduled_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PurchaseScheduledInstances(ctx, input)
			},
		},
		"reboot-instances": {
			Name:   "reboot-instances",
			Fields: fields_reboot_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RebootInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reboot_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RebootInstances(ctx, input)
			},
		},
		"register-image": {
			Name:   "register-image",
			Fields: fields_register_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterImage(ctx, input)
			},
		},
		"register-instance-event-notification-attributes": {
			Name:   "register-instance-event-notification-attributes",
			Fields: fields_register_instance_event_notification_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterInstanceEventNotificationAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_instance_event_notification_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterInstanceEventNotificationAttributes(ctx, input)
			},
		},
		"register-transit-gateway-multicast-group-members": {
			Name:   "register-transit-gateway-multicast-group-members",
			Fields: fields_register_transit_gateway_multicast_group_members,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterTransitGatewayMulticastGroupMembersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_transit_gateway_multicast_group_members, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterTransitGatewayMulticastGroupMembers(ctx, input)
			},
		},
		"register-transit-gateway-multicast-group-sources": {
			Name:   "register-transit-gateway-multicast-group-sources",
			Fields: fields_register_transit_gateway_multicast_group_sources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterTransitGatewayMulticastGroupSourcesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_transit_gateway_multicast_group_sources, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterTransitGatewayMulticastGroupSources(ctx, input)
			},
		},
		"reject-capacity-reservation-billing-ownership": {
			Name:   "reject-capacity-reservation-billing-ownership",
			Fields: fields_reject_capacity_reservation_billing_ownership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectCapacityReservationBillingOwnershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_capacity_reservation_billing_ownership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectCapacityReservationBillingOwnership(ctx, input)
			},
		},
		"reject-transit-gateway-multicast-domain-associations": {
			Name:   "reject-transit-gateway-multicast-domain-associations",
			Fields: fields_reject_transit_gateway_multicast_domain_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectTransitGatewayMulticastDomainAssociationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_transit_gateway_multicast_domain_associations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectTransitGatewayMulticastDomainAssociations(ctx, input)
			},
		},
		"reject-transit-gateway-peering-attachment": {
			Name:   "reject-transit-gateway-peering-attachment",
			Fields: fields_reject_transit_gateway_peering_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectTransitGatewayPeeringAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_transit_gateway_peering_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectTransitGatewayPeeringAttachment(ctx, input)
			},
		},
		"reject-transit-gateway-vpc-attachment": {
			Name:   "reject-transit-gateway-vpc-attachment",
			Fields: fields_reject_transit_gateway_vpc_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectTransitGatewayVpcAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_transit_gateway_vpc_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectTransitGatewayVpcAttachment(ctx, input)
			},
		},
		"reject-vpc-endpoint-connections": {
			Name:   "reject-vpc-endpoint-connections",
			Fields: fields_reject_vpc_endpoint_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectVpcEndpointConnectionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_vpc_endpoint_connections, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectVpcEndpointConnections(ctx, input)
			},
		},
		"reject-vpc-peering-connection": {
			Name:   "reject-vpc-peering-connection",
			Fields: fields_reject_vpc_peering_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectVpcPeeringConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_vpc_peering_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectVpcPeeringConnection(ctx, input)
			},
		},
		"release-address": {
			Name:   "release-address",
			Fields: fields_release_address,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReleaseAddressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_release_address, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReleaseAddress(ctx, input)
			},
		},
		"release-hosts": {
			Name:   "release-hosts",
			Fields: fields_release_hosts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReleaseHostsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_release_hosts, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReleaseHosts(ctx, input)
			},
		},
		"release-ipam-pool-allocation": {
			Name:   "release-ipam-pool-allocation",
			Fields: fields_release_ipam_pool_allocation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReleaseIpamPoolAllocationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_release_ipam_pool_allocation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReleaseIpamPoolAllocation(ctx, input)
			},
		},
		"replace-iam-instance-profile-association": {
			Name:   "replace-iam-instance-profile-association",
			Fields: fields_replace_iam_instance_profile_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReplaceIamInstanceProfileAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_replace_iam_instance_profile_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReplaceIamInstanceProfileAssociation(ctx, input)
			},
		},
		"replace-image-criteria-in-allowed-images-settings": {
			Name:   "replace-image-criteria-in-allowed-images-settings",
			Fields: fields_replace_image_criteria_in_allowed_images_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReplaceImageCriteriaInAllowedImagesSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_replace_image_criteria_in_allowed_images_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReplaceImageCriteriaInAllowedImagesSettings(ctx, input)
			},
		},
		"replace-network-acl-association": {
			Name:   "replace-network-acl-association",
			Fields: fields_replace_network_acl_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReplaceNetworkAclAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_replace_network_acl_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReplaceNetworkAclAssociation(ctx, input)
			},
		},
		"replace-network-acl-entry": {
			Name:   "replace-network-acl-entry",
			Fields: fields_replace_network_acl_entry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReplaceNetworkAclEntryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_replace_network_acl_entry, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReplaceNetworkAclEntry(ctx, input)
			},
		},
		"replace-route": {
			Name:   "replace-route",
			Fields: fields_replace_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReplaceRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_replace_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReplaceRoute(ctx, input)
			},
		},
		"replace-route-table-association": {
			Name:   "replace-route-table-association",
			Fields: fields_replace_route_table_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReplaceRouteTableAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_replace_route_table_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReplaceRouteTableAssociation(ctx, input)
			},
		},
		"replace-transit-gateway-route": {
			Name:   "replace-transit-gateway-route",
			Fields: fields_replace_transit_gateway_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReplaceTransitGatewayRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_replace_transit_gateway_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReplaceTransitGatewayRoute(ctx, input)
			},
		},
		"replace-vpn-tunnel": {
			Name:   "replace-vpn-tunnel",
			Fields: fields_replace_vpn_tunnel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReplaceVpnTunnelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_replace_vpn_tunnel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReplaceVpnTunnel(ctx, input)
			},
		},
		"report-instance-status": {
			Name:   "report-instance-status",
			Fields: fields_report_instance_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReportInstanceStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_report_instance_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReportInstanceStatus(ctx, input)
			},
		},
		"request-spot-fleet": {
			Name:   "request-spot-fleet",
			Fields: fields_request_spot_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RequestSpotFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_request_spot_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RequestSpotFleet(ctx, input)
			},
		},
		"request-spot-instances": {
			Name:   "request-spot-instances",
			Fields: fields_request_spot_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RequestSpotInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_request_spot_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RequestSpotInstances(ctx, input)
			},
		},
		"reset-address-attribute": {
			Name:   "reset-address-attribute",
			Fields: fields_reset_address_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetAddressAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_address_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetAddressAttribute(ctx, input)
			},
		},
		"reset-ebs-default-kms-key-id": {
			Name:   "reset-ebs-default-kms-key-id",
			Fields: fields_reset_ebs_default_kms_key_id,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetEbsDefaultKmsKeyIdInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_ebs_default_kms_key_id, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetEbsDefaultKmsKeyId(ctx, input)
			},
		},
		"reset-fpga-image-attribute": {
			Name:   "reset-fpga-image-attribute",
			Fields: fields_reset_fpga_image_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetFpgaImageAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_fpga_image_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetFpgaImageAttribute(ctx, input)
			},
		},
		"reset-image-attribute": {
			Name:   "reset-image-attribute",
			Fields: fields_reset_image_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetImageAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_image_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetImageAttribute(ctx, input)
			},
		},
		"reset-instance-attribute": {
			Name:   "reset-instance-attribute",
			Fields: fields_reset_instance_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetInstanceAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_instance_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetInstanceAttribute(ctx, input)
			},
		},
		"reset-network-interface-attribute": {
			Name:   "reset-network-interface-attribute",
			Fields: fields_reset_network_interface_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetNetworkInterfaceAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_network_interface_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetNetworkInterfaceAttribute(ctx, input)
			},
		},
		"reset-snapshot-attribute": {
			Name:   "reset-snapshot-attribute",
			Fields: fields_reset_snapshot_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetSnapshotAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_snapshot_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetSnapshotAttribute(ctx, input)
			},
		},
		"restore-address-to-classic": {
			Name:   "restore-address-to-classic",
			Fields: fields_restore_address_to_classic,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreAddressToClassicInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_address_to_classic, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreAddressToClassic(ctx, input)
			},
		},
		"restore-image-from-recycle-bin": {
			Name:   "restore-image-from-recycle-bin",
			Fields: fields_restore_image_from_recycle_bin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreImageFromRecycleBinInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_image_from_recycle_bin, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreImageFromRecycleBin(ctx, input)
			},
		},
		"restore-managed-prefix-list-version": {
			Name:   "restore-managed-prefix-list-version",
			Fields: fields_restore_managed_prefix_list_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreManagedPrefixListVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_managed_prefix_list_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreManagedPrefixListVersion(ctx, input)
			},
		},
		"restore-snapshot-from-recycle-bin": {
			Name:   "restore-snapshot-from-recycle-bin",
			Fields: fields_restore_snapshot_from_recycle_bin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreSnapshotFromRecycleBinInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_snapshot_from_recycle_bin, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreSnapshotFromRecycleBin(ctx, input)
			},
		},
		"restore-snapshot-tier": {
			Name:   "restore-snapshot-tier",
			Fields: fields_restore_snapshot_tier,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreSnapshotTierInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_snapshot_tier, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreSnapshotTier(ctx, input)
			},
		},
		"restore-volume-from-recycle-bin": {
			Name:   "restore-volume-from-recycle-bin",
			Fields: fields_restore_volume_from_recycle_bin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreVolumeFromRecycleBinInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_volume_from_recycle_bin, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreVolumeFromRecycleBin(ctx, input)
			},
		},
		"revoke-client-vpn-ingress": {
			Name:   "revoke-client-vpn-ingress",
			Fields: fields_revoke_client_vpn_ingress,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RevokeClientVpnIngressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_revoke_client_vpn_ingress, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RevokeClientVpnIngress(ctx, input)
			},
		},
		"revoke-security-group-egress": {
			Name:   "revoke-security-group-egress",
			Fields: fields_revoke_security_group_egress,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RevokeSecurityGroupEgressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_revoke_security_group_egress, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RevokeSecurityGroupEgress(ctx, input)
			},
		},
		"revoke-security-group-ingress": {
			Name:   "revoke-security-group-ingress",
			Fields: fields_revoke_security_group_ingress,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RevokeSecurityGroupIngressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_revoke_security_group_ingress, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RevokeSecurityGroupIngress(ctx, input)
			},
		},
		"run-instances": {
			Name:   "run-instances",
			Fields: fields_run_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RunInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_run_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RunInstances(ctx, input)
			},
		},
		"run-scheduled-instances": {
			Name:   "run-scheduled-instances",
			Fields: fields_run_scheduled_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RunScheduledInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_run_scheduled_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RunScheduledInstances(ctx, input)
			},
		},
		"search-local-gateway-routes": {
			Name:   "search-local-gateway-routes",
			Fields: fields_search_local_gateway_routes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchLocalGatewayRoutesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_local_gateway_routes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchLocalGatewayRoutes(ctx, input)
				}
				var results []*svc.SearchLocalGatewayRoutesOutput
				p := svc.NewSearchLocalGatewayRoutesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"search-transit-gateway-multicast-groups": {
			Name:   "search-transit-gateway-multicast-groups",
			Fields: fields_search_transit_gateway_multicast_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchTransitGatewayMulticastGroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_transit_gateway_multicast_groups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchTransitGatewayMulticastGroups(ctx, input)
				}
				var results []*svc.SearchTransitGatewayMulticastGroupsOutput
				p := svc.NewSearchTransitGatewayMulticastGroupsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"search-transit-gateway-routes": {
			Name:   "search-transit-gateway-routes",
			Fields: fields_search_transit_gateway_routes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchTransitGatewayRoutesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_transit_gateway_routes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchTransitGatewayRoutes(ctx, input)
				}
				var results []*svc.SearchTransitGatewayRoutesOutput
				p := svc.NewSearchTransitGatewayRoutesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"send-diagnostic-interrupt": {
			Name:   "send-diagnostic-interrupt",
			Fields: fields_send_diagnostic_interrupt,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendDiagnosticInterruptInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_diagnostic_interrupt, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendDiagnosticInterrupt(ctx, input)
			},
		},
		"start-declarative-policies-report": {
			Name:   "start-declarative-policies-report",
			Fields: fields_start_declarative_policies_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDeclarativePoliciesReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_declarative_policies_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDeclarativePoliciesReport(ctx, input)
			},
		},
		"start-instances": {
			Name:   "start-instances",
			Fields: fields_start_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartInstances(ctx, input)
			},
		},
		"start-network-insights-access-scope-analysis": {
			Name:   "start-network-insights-access-scope-analysis",
			Fields: fields_start_network_insights_access_scope_analysis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartNetworkInsightsAccessScopeAnalysisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_network_insights_access_scope_analysis, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartNetworkInsightsAccessScopeAnalysis(ctx, input)
			},
		},
		"start-network-insights-analysis": {
			Name:   "start-network-insights-analysis",
			Fields: fields_start_network_insights_analysis,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartNetworkInsightsAnalysisInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_network_insights_analysis, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartNetworkInsightsAnalysis(ctx, input)
			},
		},
		"start-vpc-endpoint-service-private-dns-verification": {
			Name:   "start-vpc-endpoint-service-private-dns-verification",
			Fields: fields_start_vpc_endpoint_service_private_dns_verification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartVpcEndpointServicePrivateDnsVerificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_vpc_endpoint_service_private_dns_verification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartVpcEndpointServicePrivateDnsVerification(ctx, input)
			},
		},
		"stop-instances": {
			Name:   "stop-instances",
			Fields: fields_stop_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopInstances(ctx, input)
			},
		},
		"terminate-client-vpn-connections": {
			Name:   "terminate-client-vpn-connections",
			Fields: fields_terminate_client_vpn_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TerminateClientVpnConnectionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_terminate_client_vpn_connections, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TerminateClientVpnConnections(ctx, input)
			},
		},
		"terminate-instances": {
			Name:   "terminate-instances",
			Fields: fields_terminate_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TerminateInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_terminate_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TerminateInstances(ctx, input)
			},
		},
		"unassign-ipv6-addresses": {
			Name:   "unassign-ipv6-addresses",
			Fields: fields_unassign_ipv6_addresses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UnassignIpv6AddressesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_unassign_ipv6_addresses, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UnassignIpv6Addresses(ctx, input)
			},
		},
		"unassign-private-ip-addresses": {
			Name:   "unassign-private-ip-addresses",
			Fields: fields_unassign_private_ip_addresses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UnassignPrivateIpAddressesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_unassign_private_ip_addresses, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UnassignPrivateIpAddresses(ctx, input)
			},
		},
		"unassign-private-nat-gateway-address": {
			Name:   "unassign-private-nat-gateway-address",
			Fields: fields_unassign_private_nat_gateway_address,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UnassignPrivateNatGatewayAddressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_unassign_private_nat_gateway_address, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UnassignPrivateNatGatewayAddress(ctx, input)
			},
		},
		"unlock-snapshot": {
			Name:   "unlock-snapshot",
			Fields: fields_unlock_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UnlockSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_unlock_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UnlockSnapshot(ctx, input)
			},
		},
		"unmonitor-instances": {
			Name:   "unmonitor-instances",
			Fields: fields_unmonitor_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UnmonitorInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_unmonitor_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UnmonitorInstances(ctx, input)
			},
		},
		"update-capacity-manager-organizations-access": {
			Name:   "update-capacity-manager-organizations-access",
			Fields: fields_update_capacity_manager_organizations_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCapacityManagerOrganizationsAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_capacity_manager_organizations_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCapacityManagerOrganizationsAccess(ctx, input)
			},
		},
		"update-interruptible-capacity-reservation-allocation": {
			Name:   "update-interruptible-capacity-reservation-allocation",
			Fields: fields_update_interruptible_capacity_reservation_allocation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateInterruptibleCapacityReservationAllocationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_interruptible_capacity_reservation_allocation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateInterruptibleCapacityReservationAllocation(ctx, input)
			},
		},
		"update-security-group-rule-descriptions-egress": {
			Name:   "update-security-group-rule-descriptions-egress",
			Fields: fields_update_security_group_rule_descriptions_egress,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSecurityGroupRuleDescriptionsEgressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_security_group_rule_descriptions_egress, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSecurityGroupRuleDescriptionsEgress(ctx, input)
			},
		},
		"update-security-group-rule-descriptions-ingress": {
			Name:   "update-security-group-rule-descriptions-ingress",
			Fields: fields_update_security_group_rule_descriptions_ingress,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSecurityGroupRuleDescriptionsIngressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_security_group_rule_descriptions_ingress, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSecurityGroupRuleDescriptionsIngress(ctx, input)
			},
		},
		"withdraw-byoip-cidr": {
			Name:   "withdraw-byoip-cidr",
			Fields: fields_withdraw_byoip_cidr,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.WithdrawByoipCidrInput{}
				if _, err := leanruntime.ApplyInput(input, fields_withdraw_byoip_cidr, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.WithdrawByoipCidr(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("ec2", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
