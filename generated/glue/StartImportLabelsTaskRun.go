package glue

// StartImportLabelsTaskRun is generated as a reference stub.
// Executable command wiring lives under cmd/glue.go.
//
// Enables you to provide additional labels (examples of truth) to be used to
// teach the machine learning transform and improve its quality. This API operation
// is generally used as part of the active learning workflow that starts with the
// StartMLLabelingSetGenerationTaskRun call and that ultimately results in
// improving the quality of your machine learning transform.
//
// After the StartMLLabelingSetGenerationTaskRun finishes, Glue machine learning
// will have generated a series of questions for humans to answer. (Answering these
// questions is often called 'labeling' in the machine learning workflows). In the
// case of the FindMatches transform, these questions are of the form, “What is
// the correct way to group these rows together into groups composed entirely of
// matching records?” After the labeling process is finished, users upload their
// answers/labels with a call to StartImportLabelsTaskRun . After
// StartImportLabelsTaskRun finishes, all future runs of the machine learning
// transform use the new and improved labels and perform a higher-quality
// transformation.
//
// By default, StartMLLabelingSetGenerationTaskRun continually learns from and
// combines all labels that you upload unless you set Replace to true. If you set
// Replace to true, StartImportLabelsTaskRun deletes and forgets all previously
// uploaded labels and learns only from the exact set that you upload. Replacing
// labels can be helpful if you realize that you previously uploaded incorrect
// labels, and you believe that they are having a negative effect on your transform
// quality.
//
// You can check on the status of your task run by calling the GetMLTaskRun
// operation.
