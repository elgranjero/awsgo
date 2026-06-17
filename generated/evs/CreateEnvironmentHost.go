package evs

// CreateEnvironmentHost is generated as a reference stub.
// Executable command wiring lives under cmd/evs.go.
//
// Creates an ESX host and adds it to an Amazon EVS environment. Amazon EVS
// supports 4-16 hosts per environment.
//
// This action can only be used after the Amazon EVS environment is deployed.
//
// You can use the dedicatedHostId parameter to specify an Amazon EC2 Dedicated
// Host for ESX host creation.
//
// You can use the placementGroupId parameter to specify a cluster or partition
// placement group to launch EC2 instances into.
//
// If you don't specify an ESX version when adding hosts using
// CreateEnvironmentHost action, Amazon EVS automatically uses the default ESX
// version associated with your environment's VCF version. To find the default ESX
// version for a particular VCF version, use the GetVersions action.
//
// You cannot use the dedicatedHostId and placementGroupId parameters together in
// the same CreateEnvironmentHost action. This results in a ValidationException
// response.
