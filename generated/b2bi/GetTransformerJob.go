package b2bi

// GetTransformerJob is generated as a reference stub.
// Executable command wiring lives under cmd/b2bi.go.
//
// Returns the details of the transformer run, based on the Transformer job ID.
//
// If 30 days have elapsed since your transformer job was started, the system
// deletes it. So, if you run GetTransformerJob and supply a transformerId and
// transformerJobId for a job that was started more than 30 days previously, you
// receive a 404 response.
