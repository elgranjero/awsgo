package polly

// StartSpeechSynthesisTask is generated as a reference stub.
// Executable command wiring lives under cmd/polly.go.
//
// Allows the creation of an asynchronous synthesis task, by starting a new
// SpeechSynthesisTask . This operation requires all the standard information
// needed for speech synthesis, plus the name of an Amazon S3 bucket for the
// service to store the output of the synthesis task and two optional parameters (
// OutputS3KeyPrefix and SnsTopicArn ). Once the synthesis task is created, this
// operation will return a SpeechSynthesisTask object, which will include an
// identifier of this task as well as the current status. The SpeechSynthesisTask
// object is available for 72 hours after starting the asynchronous synthesis task.
