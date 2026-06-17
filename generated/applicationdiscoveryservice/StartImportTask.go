package applicationdiscoveryservice

// StartImportTask is generated as a reference stub.
// Executable command wiring lives under cmd/applicationdiscoveryservice.go.
//
// Starts an import task, which allows you to import details of your on-premises
// environment directly into Amazon Web Services Migration Hub without having to
// use the Amazon Web Services Application Discovery Service (Application Discovery
// Service) tools such as the Amazon Web Services Application Discovery Service
// Agentless Collector or Application Discovery Agent. This gives you the option to
// perform migration assessment and planning directly from your imported data,
// including the ability to group your devices as applications and track their
// migration status.
//
// To start an import request, do this:
//
// - Download the specially formatted comma separated value (CSV) import
// template, which you can find here: [https://s3.us-west-2.amazonaws.com/templates-7cffcf56-bd96-4b1c-b45b-a5b42f282e46/import_template.csv].
//
// - Fill out the template with your server and application data.
//
// - Upload your import file to an Amazon S3 bucket, and make a note of it's
// Object URL. Your import file must be in the CSV format.
//
// - Use the console or the StartImportTask command with the Amazon Web Services
// CLI or one of the Amazon Web Services SDKs to import the records from your file.
//
// For more information, including step-by-step procedures, see [Migration Hub Import] in the Amazon Web
// Services Application Discovery Service User Guide.
//
// There are limits to the number of import tasks you can create (and delete) in
// an Amazon Web Services account. For more information, see [Amazon Web Services Application Discovery Service Limits]in the Amazon Web
// Services Application Discovery Service User Guide.
//
// [Amazon Web Services Application Discovery Service Limits]: https://docs.aws.amazon.com/application-discovery/latest/userguide/ads_service_limits.html
// [https://s3.us-west-2.amazonaws.com/templates-7cffcf56-bd96-4b1c-b45b-a5b42f282e46/import_template.csv]: https://s3.us-west-2.amazonaws.com/templates-7cffcf56-bd96-4b1c-b45b-a5b42f282e46/import_template.csv
// [Migration Hub Import]: https://docs.aws.amazon.com/application-discovery/latest/userguide/discovery-import.html
