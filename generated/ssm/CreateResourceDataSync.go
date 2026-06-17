package ssm

// CreateResourceDataSync is generated as a reference stub.
// Executable command wiring lives under cmd/ssm.go.
//
// A resource data sync helps you view data from multiple sources in a single
// location. Amazon Web Services Systems Manager offers two types of resource data
// sync: SyncToDestination and SyncFromSource .
//
// You can configure Systems Manager Inventory to use the SyncToDestination type
// to synchronize Inventory data from multiple Amazon Web Services Regions to a
// single Amazon Simple Storage Service (Amazon S3) bucket. For more information,
// see [Creating a resource data sync for Inventory]in the Amazon Web Services Systems Manager User Guide.
//
// You can configure Systems Manager Explorer to use the SyncFromSource type to
// synchronize operational work items (OpsItems) and operational data (OpsData)
// from multiple Amazon Web Services Regions to a single Amazon S3 bucket. This
// type can synchronize OpsItems and OpsData from multiple Amazon Web Services
// accounts and Amazon Web Services Regions or EntireOrganization by using
// Organizations. For more information, see [Setting up Systems Manager Explorer to display data from multiple accounts and Regions]in the Amazon Web Services Systems
// Manager User Guide.
//
// A resource data sync is an asynchronous operation that returns immediately.
// After a successful initial sync is completed, the system continuously syncs
// data. To check the status of a sync, use the ListResourceDataSync.
//
// By default, data isn't encrypted in Amazon S3. We strongly recommend that you
// enable encryption in Amazon S3 to ensure secure data storage. We also recommend
// that you secure access to the Amazon S3 bucket by creating a restrictive bucket
// policy.
//
// [Setting up Systems Manager Explorer to display data from multiple accounts and Regions]: https://docs.aws.amazon.com/systems-manager/latest/userguide/Explorer-resource-data-sync.html
// [Creating a resource data sync for Inventory]: https://docs.aws.amazon.com/systems-manager/latest/userguide/inventory-create-resource-data-sync.html
