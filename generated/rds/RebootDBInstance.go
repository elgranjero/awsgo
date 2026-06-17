package rds

// RebootDBInstance is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// You might need to reboot your DB instance, usually for maintenance reasons. For
// example, if you make certain modifications, or if you change the DB parameter
// group associated with the DB instance, you must reboot the instance for the
// changes to take effect.
//
// Rebooting a DB instance restarts the database engine service. Rebooting a DB
// instance results in a momentary outage, during which the DB instance status is
// set to rebooting.
//
// For more information about rebooting, see [Rebooting a DB Instance] in the Amazon RDS User Guide.
//
// This command doesn't apply to RDS Custom.
//
// If your DB instance is part of a Multi-AZ DB cluster, you can reboot the DB
// cluster with the RebootDBCluster operation.
//
// [Rebooting a DB Instance]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_RebootInstance.html
