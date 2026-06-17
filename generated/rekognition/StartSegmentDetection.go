package rekognition

// StartSegmentDetection is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// Starts asynchronous detection of segment detection in a stored video.
//
// Amazon Rekognition Video can detect segments in a video stored in an Amazon S3
// bucket. Use Videoto specify the bucket name and the filename of the video.
// StartSegmentDetection returns a job identifier ( JobId ) which you use to get
// the results of the operation. When segment detection is finished, Amazon
// Rekognition Video publishes a completion status to the Amazon Simple
// Notification Service topic that you specify in NotificationChannel .
//
// You can use the Filters (StartSegmentDetectionFilters ) input parameter to specify the minimum detection
// confidence returned in the response. Within Filters , use ShotFilter (StartShotDetectionFilter ) to
// filter detected shots. Use TechnicalCueFilter (StartTechnicalCueDetectionFilter ) to filter technical cues.
//
// To get the results of the segment detection operation, first check that the
// status value published to the Amazon SNS topic is SUCCEEDED . if so, call GetSegmentDetection and
// pass the job identifier ( JobId ) from the initial call to StartSegmentDetection
// .
//
// For more information, see Detecting video segments in stored video in the
// Amazon Rekognition Developer Guide.
