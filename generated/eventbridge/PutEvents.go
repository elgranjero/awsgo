package eventbridge

// PutEvents is generated as a reference stub.
// Executable command wiring lives under cmd/eventbridge.go.
//
// Sends custom events to Amazon EventBridge so that they can be matched to rules.
//
// You can batch multiple event entries into one request for efficiency. However,
// the total entry size must be less than 256KB. You can calculate the entry size
// before you send the events. For more information, see [Calculating PutEvents event entry size]in the Amazon EventBridge
// User Guide .
//
// PutEvents accepts the data in JSON format. For the JSON number (integer) data
// type, the constraints are: a minimum value of -9,223,372,036,854,775,808 and a
// maximum value of 9,223,372,036,854,775,807.
//
// PutEvents will only process nested JSON up to 1000 levels deep.
//
// [Calculating PutEvents event entry size]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-putevents.html#eb-putevent-size
