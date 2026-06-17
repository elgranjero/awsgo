package lightsail

// UpdateRelationalDatabaseParameters is generated as a reference stub.
// Executable command wiring lives under cmd/lightsail.go.
//
// Allows the update of one or more parameters of a database in Amazon Lightsail.
//
// Parameter updates don't cause outages; therefore, their application is not
// subject to the preferred maintenance window. However, there are two ways in
// which parameter updates are applied: dynamic or pending-reboot . Parameters
// marked with a dynamic apply type are applied immediately. Parameters marked
// with a pending-reboot apply type are applied only after the database is
// rebooted using the reboot relational database operation.
//
// The update relational database parameters operation supports tag-based access
// control via resource tags applied to the resource identified by
// relationalDatabaseName. For more information, see the [Amazon Lightsail Developer Guide].
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
