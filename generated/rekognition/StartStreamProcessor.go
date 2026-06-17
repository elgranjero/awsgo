package rekognition

// StartStreamProcessor is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// Starts processing a stream processor. You create a stream processor by calling CreateStreamProcessor
// . To tell StartStreamProcessor which stream processor to start, use the value
// of the Name field specified in the call to CreateStreamProcessor .
//
// If you are using a label detection stream processor to detect labels, you need
// to provide a Start selector and a Stop selector to determine the length of the
// stream processing time.
