package medicalimaging

// GetDICOMImportJob is generated as a reference stub.
// Executable command wiring lives under cmd/medicalimaging.go.
//
// Get the import job properties to learn more about the job or job progress.
//
// The jobStatus refers to the execution of the import job. Therefore, an import
// job can return a jobStatus as COMPLETED even if validation issues are
// discovered during the import process. If a jobStatus returns as COMPLETED , we
// still recommend you review the output manifests written to S3, as they provide
// details on the success or failure of individual P10 object imports.
