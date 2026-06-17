package rekognition

// StartCelebrityRecognition is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// Starts asynchronous recognition of celebrities in a stored video.
//
// Amazon Rekognition Video can detect celebrities in a video must be stored in an
// Amazon S3 bucket. Use Videoto specify the bucket name and the filename of the video.
// StartCelebrityRecognition returns a job identifier ( JobId ) which you use to
// get the results of the analysis. When celebrity recognition analysis is
// finished, Amazon Rekognition Video publishes a completion status to the Amazon
// Simple Notification Service topic that you specify in NotificationChannel . To
// get the results of the celebrity recognition analysis, first check that the
// status value published to the Amazon SNS topic is SUCCEEDED . If so, call GetCelebrityRecognition and
// pass the job identifier ( JobId ) from the initial call to
// StartCelebrityRecognition .
//
// For more information, see Recognizing celebrities in the Amazon Rekognition
// Developer Guide.
