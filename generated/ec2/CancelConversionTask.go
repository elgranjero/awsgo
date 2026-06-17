package ec2

// CancelConversionTask is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Cancels an active conversion task. The task can be the import of an instance or
// volume. The action removes all artifacts of the conversion, including a
// partially uploaded volume or instance. If the conversion is complete or is in
// the process of transferring the final disk image, the command fails and returns
// an exception.
