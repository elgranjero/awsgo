package evs

// CreateEnvironment is generated as a reference stub.
// Executable command wiring lives under cmd/evs.go.
//
// Creates an Amazon EVS environment that runs VCF software, such as SDDC Manager,
// NSX Manager, and vCenter Server.
//
// During environment creation, Amazon EVS performs validations on DNS settings,
// provisions VLAN subnets and hosts, and deploys the supplied version of VCF.
//
// It can take several hours to create an environment. After the deployment
// completes, you can configure VCF in the vSphere user interface according to your
// needs.
//
// When creating a new environment, the default ESX version for the selected VCF
// version will be used, you cannot choose a specific ESX version in
// CreateEnvironment action. When a host has been added with a specific ESX
// version, it can only be upgraded using vCenter Lifecycle Manager.
//
// You cannot use the dedicatedHostId and placementGroupId parameters together in
// the same CreateEnvironment action. This results in a ValidationException
// response.
