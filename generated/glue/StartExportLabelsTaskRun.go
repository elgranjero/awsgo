package glue

// StartExportLabelsTaskRun is generated as a reference stub.
// Executable command wiring lives under cmd/glue.go.
//
// Begins an asynchronous task to export all labeled data for a particular
// transform. This task is the only label-related API call that is not part of the
// typical active learning workflow. You typically use StartExportLabelsTaskRun
// when you want to work with all of your existing labels at the same time, such as
// when you want to remove or change labels that were previously submitted as
// truth. This API operation accepts the TransformId whose labels you want to
// export and an Amazon Simple Storage Service (Amazon S3) path to export the
// labels to. The operation returns a TaskRunId . You can check on the status of
// your task run by calling the GetMLTaskRun API.
