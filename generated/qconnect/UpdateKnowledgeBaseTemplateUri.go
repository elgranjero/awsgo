package qconnect

// UpdateKnowledgeBaseTemplateUri is generated as a reference stub.
// Executable command wiring lives under cmd/qconnect.go.
//
// Updates the template URI of a knowledge base. This is only supported for
// knowledge bases of type EXTERNAL. Include a single variable in ${variable}
// format; this interpolated by Amazon Q in Connect using ingested content. For
// example, if you ingest a Salesforce article, it has an Id value, and you can
// set the template URI to
// https://myInstanceName.lightning.force.com/lightning/r/Knowledge__kav/*${Id}*/view
// .
