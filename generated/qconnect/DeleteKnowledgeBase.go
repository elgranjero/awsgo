package qconnect

// DeleteKnowledgeBase is generated as a reference stub.
// Executable command wiring lives under cmd/qconnect.go.
//
// Deletes the knowledge base.
//
// When you use this API to delete an external knowledge base such as Salesforce
// or ServiceNow, you must also delete the [Amazon AppIntegrations]DataIntegration. This is because you
// can't reuse the DataIntegration after it's been associated with an external
// knowledge base. However, you can delete and recreate it. See [DeleteDataIntegration]and [CreateDataIntegration] in the Amazon
// AppIntegrations API Reference.
//
// [Amazon AppIntegrations]: https://docs.aws.amazon.com/appintegrations/latest/APIReference/Welcome.html
// [DeleteDataIntegration]: https://docs.aws.amazon.com/appintegrations/latest/APIReference/API_DeleteDataIntegration.html
// [CreateDataIntegration]: https://docs.aws.amazon.com/appintegrations/latest/APIReference/API_CreateDataIntegration.html
