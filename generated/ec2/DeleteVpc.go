package ec2

// DeleteVpc is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Deletes the specified VPC. You must detach or delete all gateways and resources
// that are associated with the VPC before you can delete it. For example, you must
// terminate all instances running in the VPC, delete all security groups
// associated with the VPC (except the default one), delete all route tables
// associated with the VPC (except the default one), and so on. When you delete the
// VPC, it deletes the default security group, network ACL, and route table for the
// VPC.
//
// If you created a flow log for the VPC that you are deleting, note that flow
// logs for deleted VPCs are eventually automatically removed.
