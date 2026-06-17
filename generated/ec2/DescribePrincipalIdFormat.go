package ec2

// DescribePrincipalIdFormat is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Describes the ID format settings for the root user and all IAM roles and IAM
// users that have explicitly specified a longer ID (17-character ID) preference.
//
// By default, all IAM roles and IAM users default to the same ID settings as the
// root user, unless they explicitly override the settings. This request is useful
// for identifying those IAM users and IAM roles that have overridden the default
// ID settings.
//
// The following resource types support longer IDs: bundle | conversion-task |
// customer-gateway | dhcp-options | elastic-ip-allocation | elastic-ip-association
// | export-task | flow-log | image | import-task | instance | internet-gateway |
// network-acl | network-acl-association | network-interface |
// network-interface-attachment | prefix-list | reservation | route-table |
// route-table-association | security-group | snapshot | subnet |
// subnet-cidr-block-association | volume | vpc | vpc-cidr-block-association |
// vpc-endpoint | vpc-peering-connection | vpn-connection | vpn-gateway .
