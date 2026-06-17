package rekognition

// StartFaceDetection is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// Starts asynchronous detection of faces in a stored video.
//
// Amazon Rekognition Video can detect faces in a video stored in an Amazon S3
// bucket. Use Videoto specify the bucket name and the filename of the video.
// StartFaceDetection returns a job identifier ( JobId ) that you use to get the
// results of the operation. When face detection is finished, Amazon Rekognition
// Video publishes a completion status to the Amazon Simple Notification Service
// topic that you specify in NotificationChannel . To get the results of the face
// detection operation, first check that the status value published to the Amazon
// SNS topic is SUCCEEDED . If so, call GetFaceDetection and pass the job identifier ( JobId ) from
// the initial call to StartFaceDetection .
//
// For more information, see Detecting faces in a stored video in the Amazon
// Rekognition Developer Guide.
