package appflow

// ResetConnectorMetadataCache is generated as a reference stub.
// Executable command wiring lives under cmd/appflow.go.
//
// Resets metadata about your connector entities that Amazon AppFlow stored in its
// cache. Use this action when you want Amazon AppFlow to return the latest
// information about the data that you have in a source application.
//
// Amazon AppFlow returns metadata about your entities when you use the
// ListConnectorEntities or DescribeConnectorEntities actions. Following these
// actions, Amazon AppFlow caches the metadata to reduce the number of API requests
// that it must send to the source application. Amazon AppFlow automatically resets
// the cache once every hour, but you can use this action when you want to get the
// latest metadata right away.
