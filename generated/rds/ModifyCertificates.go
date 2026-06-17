package rds

// ModifyCertificates is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Override the system-default Secure Sockets Layer/Transport Layer Security
// (SSL/TLS) certificate for Amazon RDS for new DB instances, or remove the
// override.
//
// By using this operation, you can specify an RDS-approved SSL/TLS certificate
// for new DB instances that is different from the default certificate provided by
// RDS. You can also use this operation to remove the override, so that new DB
// instances use the default certificate provided by RDS.
//
// You might need to override the default certificate in the following situations:
//
// - You already migrated your applications to support the latest certificate
// authority (CA) certificate, but the new CA certificate is not yet the RDS
// default CA certificate for the specified Amazon Web Services Region.
//
// - RDS has already moved to a new default CA certificate for the specified
// Amazon Web Services Region, but you are still in the process of supporting the
// new CA certificate. In this case, you temporarily need additional time to finish
// your application changes.
//
// For more information about rotating your SSL/TLS certificate for RDS DB
// engines, see [Rotating Your SSL/TLS Certificate]in the Amazon RDS User Guide.
//
// For more information about rotating your SSL/TLS certificate for Aurora DB
// engines, see [Rotating Your SSL/TLS Certificate]in the Amazon Aurora User Guide.
//
// [Rotating Your SSL/TLS Certificate]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL-certificate-rotation.html
