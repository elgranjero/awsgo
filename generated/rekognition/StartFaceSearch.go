package rekognition

// StartFaceSearch is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// Starts the asynchronous search for faces in a collection that match the faces
// of persons detected in a stored video.
//
// The video must be stored in an Amazon S3 bucket. Use Video to specify the bucket
// name and the filename of the video. StartFaceSearch returns a job identifier (
// JobId ) which you use to get the search results once the search has completed.
// When searching is finished, Amazon Rekognition Video publishes a completion
// status to the Amazon Simple Notification Service topic that you specify in
// NotificationChannel . To get the search results, first check that the status
// value published to the Amazon SNS topic is SUCCEEDED . If so, call GetFaceSearch and pass
// the job identifier ( JobId ) from the initial call to StartFaceSearch . For more
// information, see [Searching stored videos for faces].
//
// [Searching stored videos for faces]: https://docs.aws.amazon.com/rekognition/latest/dg/procedure-person-search-videos.html
