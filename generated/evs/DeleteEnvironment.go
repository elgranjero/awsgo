package evs

// DeleteEnvironment is generated as a reference stub.
// Executable command wiring lives under cmd/evs.go.
//
// Deletes an Amazon EVS environment.
//
// Amazon EVS environments will only be enabled for deletion once the hosts are
// deleted. You can delete hosts using the DeleteEnvironmentHost action.
//
// Environment deletion also deletes the associated Amazon EVS VLAN subnets and
// Amazon Web Services Secrets Manager secrets that Amazon EVS created. Amazon Web
// Services resources that you create are not deleted. These resources may continue
// to incur costs.
