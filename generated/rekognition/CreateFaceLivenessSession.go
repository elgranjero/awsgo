package rekognition

// CreateFaceLivenessSession is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// This API operation initiates a Face Liveness session. It returns a SessionId ,
// which you can use to start streaming Face Liveness video and get the results for
// a Face Liveness session.
//
// You can use the OutputConfig option in the Settings parameter to provide an
// Amazon S3 bucket location. The Amazon S3 bucket stores reference images and
// audit images. If no Amazon S3 bucket is defined, raw bytes are sent instead.
//
// You can use AuditImagesLimit to limit the number of audit images returned when
// GetFaceLivenessSessionResults is called. This number is between 0 and 4. By
// default, it is set to 0. The limit is best effort and based on the duration of
// the selfie-video.
