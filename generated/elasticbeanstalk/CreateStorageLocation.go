package elasticbeanstalk

// CreateStorageLocation is generated as a reference stub.
// Executable command wiring lives under cmd/elasticbeanstalk.go.
//
// Creates a bucket in Amazon S3 to store application versions, logs, and other
// files used by Elastic Beanstalk environments. The Elastic Beanstalk console and
// EB CLI call this API the first time you create an environment in a region. If
// the storage location already exists, CreateStorageLocation still returns the
// bucket name but does not create a new bucket.
