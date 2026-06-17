package rekognition

// StartLabelDetection is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// Starts asynchronous detection of labels in a stored video.
//
// Amazon Rekognition Video can detect labels in a video. Labels are instances of
// real-world entities. This includes objects like flower, tree, and table; events
// like wedding, graduation, and birthday party; concepts like landscape, evening,
// and nature; and activities like a person getting out of a car or a person
// skiing.
//
// The video must be stored in an Amazon S3 bucket. Use Video to specify the bucket
// name and the filename of the video. StartLabelDetection returns a job
// identifier ( JobId ) which you use to get the results of the operation. When
// label detection is finished, Amazon Rekognition Video publishes a completion
// status to the Amazon Simple Notification Service topic that you specify in
// NotificationChannel .
//
// To get the results of the label detection operation, first check that the
// status value published to the Amazon SNS topic is SUCCEEDED . If so, call GetLabelDetection and
// pass the job identifier ( JobId ) from the initial call to StartLabelDetection .
//
// # Optional Parameters
//
// StartLabelDetection has the GENERAL_LABELS Feature applied by default. This
// feature allows you to provide filtering criteria to the Settings parameter. You
// can filter with sets of individual labels or with label categories. You can
// specify inclusive filters, exclusive filters, or a combination of inclusive and
// exclusive filters. For more information on filtering, see [Detecting labels in a video].
//
// You can specify MinConfidence to control the confidence threshold for the
// labels returned. The default is 50.
//
// [Detecting labels in a video]: https://docs.aws.amazon.com/rekognition/latest/dg/labels-detecting-labels-video.html
