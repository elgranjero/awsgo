package rekognition

// ListDatasetEntries is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// This operation applies only to Amazon Rekognition Custom Labels.
//
// Lists the entries (images) within a dataset. An entry is a JSON Line that
// contains the information for a single image, including the image location,
// assigned labels, and object location bounding boxes. For more information, see [Creating a manifest file].
//
// JSON Lines in the response include information about non-terminal errors found
// in the dataset. Non terminal errors are reported in errors lists within each
// JSON Line. The same information is reported in the training and testing
// validation result manifests that Amazon Rekognition Custom Labels creates during
// model training.
//
// You can filter the response in variety of ways, such as choosing which labels
// to return and returning JSON Lines created after a specific date.
//
// This operation requires permissions to perform the
// rekognition:ListDatasetEntries action.
//
// [Creating a manifest file]: https://docs.aws.amazon.com/rekognition/latest/customlabels-dg/md-manifest-files.html
