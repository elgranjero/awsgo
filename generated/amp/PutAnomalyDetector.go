package amp

// PutAnomalyDetector is generated as a reference stub.
// Executable command wiring lives under cmd/amp.go.
//
// When you call PutAnomalyDetector , the operation creates a new anomaly detector
// if one doesn't exist, or updates an existing one. Each call to this operation
// triggers a complete retraining of the detector, which includes querying the
// minimum required samples and backfilling the detector with historical data. This
// process occurs regardless of whether you're making a minor change like updating
// the evaluation interval or making more substantial modifications. The operation
// serves as the single method for creating, updating, and retraining anomaly
// detectors.
