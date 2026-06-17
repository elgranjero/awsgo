package lookoutequipment

// CreateInferenceScheduler is generated as a reference stub.
// Executable command wiring lives under cmd/lookoutequipment.go.
//
// Creates a scheduled inference. Scheduling an inference is setting up a
//
// continuous real-time inference plan to analyze new measurement data. When
// setting up the schedule, you provide an S3 bucket location for the input data,
// assign it a delimiter between separate entries in the data, set an offset delay
// if desired, and set the frequency of inferencing. You must also provide an S3
// bucket location for the output data.
