package wisdom

// CreateKnowledgeBase is generated as a reference stub.
// Executable command wiring lives under cmd/wisdom.go.
//
// Creates a knowledge base.
//
// When using this API, you cannot reuse [Amazon AppIntegrations] DataIntegrations with external knowledge
// bases such as Salesforce and ServiceNow. If you do, you'll get an
// InvalidRequestException error.
//
// For example, you're programmatically managing your external knowledge base, and
// you want to add or remove one of the fields that is being ingested from
// Salesforce. Do the following:
//
// - Call [DeleteKnowledgeBase].
//
// - Call [DeleteDataIntegration].
//
// - Call [CreateDataIntegration]to recreate the DataIntegration or a create different one.
//
// - Call CreateKnowledgeBase.
//
// [Amazon AppIntegrations]: https://docs.aws.amazon.com/appintegrations/latest/APIReference/Welcome.html
// [DeleteKnowledgeBase]: https://docs.aws.amazon.com/wisdom/latest/APIReference/API_DeleteKnowledgeBase.html
// [DeleteDataIntegration]: https://docs.aws.amazon.com/appintegrations/latest/APIReference/API_DeleteDataIntegration.html
// [CreateDataIntegration]: https://docs.aws.amazon.com/appintegrations/latest/APIReference/API_CreateDataIntegration.html
