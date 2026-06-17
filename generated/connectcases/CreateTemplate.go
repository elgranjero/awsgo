package connectcases

// CreateTemplate is generated as a reference stub.
// Executable command wiring lives under cmd/connectcases.go.
//
// Creates a template in the Cases domain. This template is used to define the
// case object model (that is, to define what data can be captured on cases) in a
// Cases domain. A template must have a unique name within a domain, and it must
// reference existing field IDs and layout IDs. Additionally, multiple fields with
// same IDs are not allowed within the same Template. A template can be either
// Active or Inactive, as indicated by its status. Inactive templates cannot be
// used to create cases.
//
// Other template APIs are:
//
// [DeleteTemplate]
//
// [GetTemplate]
//
// [ListTemplates]
//
// [UpdateTemplate]
//
// [DeleteTemplate]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_DeleteTemplate.html
// [ListTemplates]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_ListTemplates.html
// [UpdateTemplate]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_UpdateTemplate.html
// [GetTemplate]: https://docs.aws.amazon.com/connect/latest/APIReference/API_connect-cases_GetTemplate.html
