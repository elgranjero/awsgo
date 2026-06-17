package rekognition

// StartPersonTracking is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// End of support notice: On October 31, 2025, AWS will discontinue support for
//
// Amazon Rekognition People Pathing. After October 31, 2025, you will no longer be
// able to use the Rekognition People Pathing capability. For more information,
// visit this [blog post].
//
// Starts the asynchronous tracking of a person's path in a stored video.
//
// Amazon Rekognition Video can track the path of people in a video stored in an
// Amazon S3 bucket. Use Videoto specify the bucket name and the filename of the video.
// StartPersonTracking returns a job identifier ( JobId ) which you use to get the
// results of the operation. When label detection is finished, Amazon Rekognition
// publishes a completion status to the Amazon Simple Notification Service topic
// that you specify in NotificationChannel .
//
// To get the results of the person detection operation, first check that the
// status value published to the Amazon SNS topic is SUCCEEDED . If so, call GetPersonTracking and
// pass the job identifier ( JobId ) from the initial call to StartPersonTracking .
//
// [blog post]: https://aws.amazon.com/blogs/machine-learning/transitioning-from-amazon-rekognition-people-pathing-exploring-other-alternatives/
