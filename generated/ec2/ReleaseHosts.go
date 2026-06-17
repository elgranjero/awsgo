package ec2

// ReleaseHosts is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// When you no longer want to use an On-Demand Dedicated Host it can be released.
// On-Demand billing is stopped and the host goes into released state. The host ID
// of Dedicated Hosts that have been released can no longer be specified in another
// request, for example, to modify the host. You must stop or terminate all
// instances on a host before it can be released.
//
// When Dedicated Hosts are released, it may take some time for them to stop
// counting toward your limit and you may receive capacity errors when trying to
// allocate new Dedicated Hosts. Wait a few minutes and then try again.
//
// Released hosts still appear in a DescribeHosts response.
