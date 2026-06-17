package iotsitewise

// CreateBulkImportJob is generated as a reference stub.
// Executable command wiring lives under cmd/iotsitewise.go.
//
// Defines a job to ingest data to IoT SiteWise from Amazon S3. For more
// information, see [Create a bulk import job (CLI)]in the Amazon Simple Storage Service User Guide.
//
// Before you create a bulk import job, you must enable IoT SiteWise warm tier or
// IoT SiteWise cold tier. For more information about how to configure storage
// settings, see [PutStorageConfiguration].
//
// Bulk import is designed to store historical data to IoT SiteWise.
//
// - Newly ingested data in the hot tier triggers notifications and computations.
//
// - After data moves from the hot tier to the warm or cold tier based on
// retention settings, it does not trigger computations or notifications.
//
// - Data older than 7 days does not trigger computations or notifications.
//
// [Create a bulk import job (CLI)]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/CreateBulkImportJob.html
// [PutStorageConfiguration]: https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_PutStorageConfiguration.html
