package rekognition

// CreateStreamProcessor is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// Creates an Amazon Rekognition stream processor that you can use to detect and
// recognize faces or to detect labels in a streaming video.
//
// Amazon Rekognition Video is a consumer of live video from Amazon Kinesis Video
// Streams. There are two different settings for stream processors in Amazon
// Rekognition: detecting faces and detecting labels.
//
// - If you are creating a stream processor for detecting faces, you provide as
// input a Kinesis video stream ( Input ) and a Kinesis data stream ( Output )
// stream for receiving the output. You must use the FaceSearch option in
// Settings , specifying the collection that contains the faces you want to
// recognize. After you have finished analyzing a streaming video, use StopStreamProcessorto stop
// processing.
//
// - If you are creating a stream processor to detect labels, you provide as
// input a Kinesis video stream ( Input ), Amazon S3 bucket information ( Output
// ), and an Amazon SNS topic ARN ( NotificationChannel ). You can also provide a
// KMS key ID to encrypt the data sent to your Amazon S3 bucket. You specify what
// you want to detect by using the ConnectedHome option in settings, and
// selecting one of the following: PERSON , PET , PACKAGE , ALL You can also
// specify where in the frame you want Amazon Rekognition to monitor with
// RegionsOfInterest . When you run the StartStreamProcessoroperation on a label detection stream
// processor, you input start and stop information to determine the length of the
// processing time.
//
// Use Name to assign an identifier for the stream processor. You use Name to
// manage the stream processor. For example, you can start processing the source
// video by calling StartStreamProcessorwith the Name field.
//
// This operation requires permissions to perform the
// rekognition:CreateStreamProcessor action. If you want to tag your stream
// processor, you also require permission to perform the rekognition:TagResource
// operation.
