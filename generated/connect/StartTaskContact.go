package connect

// StartTaskContact is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Initiates a flow to start a new task contact. For more information about task
// contacts, see [Concepts: Tasks in Amazon Connect]in the Amazon Connect Administrator Guide.
//
// When using PreviousContactId and RelatedContactId input parameters, note the
// following:
//
// - PreviousContactId
//
// - Any updates to user-defined task contact attributes on any contact linked
// through the same PreviousContactId will affect every contact in the chain.
//
// - There can be a maximum of 12 linked task contacts in a chain. That is, 12
// task contacts can be created that share the same PreviousContactId .
//
// - RelatedContactId
//
// - Copies contact attributes from the related task contact to the new contact.
//
// - Any update on attributes in a new task contact does not update attributes
// on previous contact.
//
// - There’s no limit on the number of task contacts that can be created that
// use the same RelatedContactId .
//
// In addition, when calling StartTaskContact include only one of these
// parameters: ContactFlowID , QuickConnectID , or TaskTemplateID . Only one
// parameter is required as long as the task template has a flow configured to run
// it. If more than one parameter is specified, or only the TaskTemplateID is
// specified but it does not have a flow configured, the request returns an error
// because Amazon Connect cannot identify the unique flow to run when the task is
// created.
//
// A ServiceQuotaExceededException occurs when the number of open tasks exceeds
// the active tasks quota or there are already 12 tasks referencing the same
// PreviousContactId . For more information about service quotas for task contacts,
// see [Amazon Connect service quotas]in the Amazon Connect Administrator Guide.
//
// [Amazon Connect service quotas]: https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html
// [Concepts: Tasks in Amazon Connect]: https://docs.aws.amazon.com/connect/latest/adminguide/tasks.html
