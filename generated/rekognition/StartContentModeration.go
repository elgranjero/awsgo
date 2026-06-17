package rekognition

// StartContentModeration is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// Starts asynchronous detection of inappropriate, unwanted, or offensive content
//
// in a stored video. For a list of moderation labels in Amazon Rekognition, see [Using the image and video moderation APIs].
//
// Amazon Rekognition Video can moderate content in a video stored in an Amazon S3
// bucket. Use Videoto specify the bucket name and the filename of the video.
// StartContentModeration returns a job identifier ( JobId ) which you use to get
// the results of the analysis. When content analysis is finished, Amazon
// Rekognition Video publishes a completion status to the Amazon Simple
// Notification Service topic that you specify in NotificationChannel .
//
// To get the results of the content analysis, first check that the status value
// published to the Amazon SNS topic is SUCCEEDED . If so, call GetContentModeration and pass the job
// identifier ( JobId ) from the initial call to StartContentModeration .
//
// For more information, see Moderating content in the Amazon Rekognition
// Developer Guide.
//
// [Using the image and video moderation APIs]: https://docs.aws.amazon.com/rekognition/latest/dg/moderation.html#moderation-api
