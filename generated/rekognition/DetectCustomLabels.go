package rekognition

// DetectCustomLabels is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// This operation applies only to Amazon Rekognition Custom Labels.
//
// Detects custom labels in a supplied image by using an Amazon Rekognition Custom
// Labels model.
//
// You specify which version of a model version to use by using the
// ProjectVersionArn input parameter.
//
// You pass the input image as base64-encoded image bytes or as a reference to an
// image in an Amazon S3 bucket. If you use the AWS CLI to call Amazon Rekognition
// operations, passing image bytes is not supported. The image must be either a PNG
// or JPEG formatted file.
//
// For each object that the model version detects on an image, the API returns a (
// CustomLabel ) object in an array ( CustomLabels ). Each CustomLabel object
// provides the label name ( Name ), the level of confidence that the image
// contains the object ( Confidence ), and object location information, if it
// exists, for the label on the image ( Geometry ).
//
// To filter labels that are returned, specify a value for MinConfidence .
// DetectCustomLabelsLabels only returns labels with a confidence that's higher
// than the specified value.
//
// The value of MinConfidence maps to the assumed threshold values created during
// training. For more information, see Assumed threshold in the Amazon Rekognition
// Custom Labels Developer Guide. Amazon Rekognition Custom Labels metrics
// expresses an assumed threshold as a floating point value between 0-1. The range
// of MinConfidence normalizes the threshold value to a percentage value (0-100).
// Confidence responses from DetectCustomLabels are also returned as a percentage.
// You can use MinConfidence to change the precision and recall or your model. For
// more information, see Analyzing an image in the Amazon Rekognition Custom Labels
// Developer Guide.
//
// If you don't specify a value for MinConfidence , DetectCustomLabels returns
// labels based on the assumed threshold of each label.
//
// This is a stateless API operation. That is, the operation does not persist any
// data.
//
// This operation requires permissions to perform the
// rekognition:DetectCustomLabels action.
//
// For more information, see Analyzing an image in the Amazon Rekognition Custom
// Labels Developer Guide.
