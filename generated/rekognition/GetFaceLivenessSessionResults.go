package rekognition

// GetFaceLivenessSessionResults is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// Retrieves the results of a specific Face Liveness session. It requires the
// sessionId as input, which was created using CreateFaceLivenessSession . Returns
// the corresponding Face Liveness confidence score, a reference image that
// includes a face bounding box, and audit images that also contain face bounding
// boxes. The Face Liveness confidence score ranges from 0 to 100.
//
// The number of audit images returned by GetFaceLivenessSessionResults is defined
// by the AuditImagesLimit paramater when calling CreateFaceLivenessSession .
// Reference images are always returned when possible.
