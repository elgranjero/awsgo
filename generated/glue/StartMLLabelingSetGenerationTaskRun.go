package glue

// StartMLLabelingSetGenerationTaskRun is generated as a reference stub.
// Executable command wiring lives under cmd/glue.go.
//
// Starts the active learning workflow for your machine learning transform to
// improve the transform's quality by generating label sets and adding labels.
//
// When the StartMLLabelingSetGenerationTaskRun finishes, Glue will have generated
// a "labeling set" or a set of questions for humans to answer.
//
// In the case of the FindMatches transform, these questions are of the form,
// “What is the correct way to group these rows together into groups composed
// entirely of matching records?”
//
// After the labeling process is finished, you can upload your labels with a call
// to StartImportLabelsTaskRun . After StartImportLabelsTaskRun finishes, all
// future runs of the machine learning transform will use the new and improved
// labels and perform a higher-quality transformation.
//
// Note: The role used to write the generated labeling set to the OutputS3Path is
// the role associated with the Machine Learning Transform, specified in the
// CreateMLTransform API.
