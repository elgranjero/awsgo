package cloudtrail

// StartImport is generated as a reference stub.
// Executable command wiring lives under cmd/cloudtrail.go.
//
// Starts an import of logged trail events from a source S3 bucket to a
//
// destination event data store. By default, CloudTrail only imports events
// contained in the S3 bucket's CloudTrail prefix and the prefixes inside the
// CloudTrail prefix, and does not check prefixes for other Amazon Web Services
// services. If you want to import CloudTrail events contained in another prefix,
// you must include the prefix in the S3LocationUri . For more considerations about
// importing trail events, see [Considerations for copying trail events]in the CloudTrail User Guide.
//
// When you start a new import, the Destinations and ImportSource parameters are
// required. Before starting a new import, disable any access control lists (ACLs)
// attached to the source S3 bucket. For more information about disabling ACLs, see
// [Controlling ownership of objects and disabling ACLs for your bucket].
//
// When you retry an import, the ImportID parameter is required.
//
// If the destination event data store is for an organization, you must use the
// management account to import trail events. You cannot use the delegated
// administrator account for the organization.
//
// [Considerations for copying trail events]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-copy-trail-to-lake.html#cloudtrail-trail-copy-considerations
// [Controlling ownership of objects and disabling ACLs for your bucket]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html
