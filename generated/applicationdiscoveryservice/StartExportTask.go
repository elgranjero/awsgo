package applicationdiscoveryservice

// StartExportTask is generated as a reference stub.
// Executable command wiring lives under cmd/applicationdiscoveryservice.go.
//
// Begins the export of a discovered data report to an Amazon S3 bucket managed by
// Amazon Web Services.
//
// Exports might provide an estimate of fees and savings based on certain
// information that you provide. Fee estimates do not include any taxes that might
// apply. Your actual fees and savings depend on a variety of factors, including
// your actual usage of Amazon Web Services services, which might vary from the
// estimates provided in this report.
//
// If you do not specify preferences or agentIds in the filter, a summary of all
// servers, applications, tags, and performance is generated. This data is an
// aggregation of all server data collected through on-premises tooling, file
// import, application grouping and applying tags.
//
// If you specify agentIds in a filter, the task exports up to 72 hours of
// detailed data collected by the identified Application Discovery Agent, including
// network, process, and performance details. A time range for exported agent data
// may be set by using startTime and endTime . Export of detailed agent data is
// limited to five concurrently running exports. Export of detailed agent data is
// limited to two exports per day.
//
// If you enable ec2RecommendationsPreferences in preferences , an Amazon EC2
// instance matching the characteristics of each server in Application Discovery
// Service is generated. Changing the attributes of the
// ec2RecommendationsPreferences changes the criteria of the recommendation.
