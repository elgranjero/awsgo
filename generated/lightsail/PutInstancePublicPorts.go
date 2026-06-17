package lightsail

// PutInstancePublicPorts is generated as a reference stub.
// Executable command wiring lives under cmd/lightsail.go.
//
// Opens ports for a specific Amazon Lightsail instance, and specifies the IP
// addresses allowed to connect to the instance through the ports, and the
// protocol. This action also closes all currently open ports that are not included
// in the request. Include all of the ports and the protocols you want to open in
// your PutInstancePublicPorts request. Or use the OpenInstancePublicPorts action
// to open ports without closing currently open ports.
//
// The PutInstancePublicPorts action supports tag-based access control via
// resource tags applied to the resource identified by instanceName . For more
// information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
