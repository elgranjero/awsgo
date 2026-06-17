package rekognition

// StartTextDetection is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// Starts asynchronous detection of text in a stored video.
//
// Amazon Rekognition Video can detect text in a video stored in an Amazon S3
// bucket. Use Videoto specify the bucket name and the filename of the video.
// StartTextDetection returns a job identifier ( JobId ) which you use to get the
// results of the operation. When text detection is finished, Amazon Rekognition
// Video publishes a completion status to the Amazon Simple Notification Service
// topic that you specify in NotificationChannel .
//
// To get the results of the text detection operation, first check that the status
// value published to the Amazon SNS topic is SUCCEEDED . if so, call GetTextDetection and pass
// the job identifier ( JobId ) from the initial call to StartTextDetection .
