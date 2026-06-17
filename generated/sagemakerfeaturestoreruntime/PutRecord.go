package sagemakerfeaturestoreruntime

// PutRecord is generated as a reference stub.
// Executable command wiring lives under cmd/sagemakerfeaturestoreruntime.go.
//
// The PutRecord API is used to ingest a list of Records into your feature group.
//
// If a new record’s EventTime is greater, the new record is written to both the
// OnlineStore and OfflineStore . Otherwise, the record is a historic record and it
// is written only to the OfflineStore .
//
// You can specify the ingestion to be applied to the OnlineStore , OfflineStore ,
// or both by using the TargetStores request parameter.
//
// You can set the ingested record to expire at a given time to live (TTL)
// duration after the record’s event time, ExpiresAt = EventTime + TtlDuration , by
// specifying the TtlDuration parameter. A record level TtlDuration is set when
// specifying the TtlDuration parameter using the PutRecord API call. If the input
// TtlDuration is null or unspecified, TtlDuration is set to the default feature
// group level TtlDuration . A record level TtlDuration supersedes the group level
// TtlDuration .
