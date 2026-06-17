package rekognition

// UpdateDatasetEntries is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// This operation applies only to Amazon Rekognition Custom Labels.
//
// Adds or updates one or more entries (images) in a dataset. An entry is a JSON
// Line which contains the information for a single image, including the image
// location, assigned labels, and object location bounding boxes. For more
// information, see Image-Level labels in manifest files and Object localization in
// manifest files in the Amazon Rekognition Custom Labels Developer Guide.
//
// If the source-ref field in the JSON line references an existing image, the
// existing image in the dataset is updated. If source-ref field doesn't reference
// an existing image, the image is added as a new image to the dataset.
//
// You specify the changes that you want to make in the Changes input parameter.
// There isn't a limit to the number JSON Lines that you can change, but the size
// of Changes must be less than 5MB.
//
// UpdateDatasetEntries returns immediatly, but the dataset update might take a
// while to complete. Use DescribeDatasetto check the current status. The dataset updated
// successfully if the value of Status is UPDATE_COMPLETE .
//
// To check if any non-terminal errors occured, call ListDatasetEntries and check for the presence
// of errors lists in the JSON Lines.
//
// Dataset update fails if a terminal error occurs ( Status = UPDATE_FAILED ).
// Currently, you can't access the terminal error information from the Amazon
// Rekognition Custom Labels SDK.
//
// This operation requires permissions to perform the
// rekognition:UpdateDatasetEntries action.
