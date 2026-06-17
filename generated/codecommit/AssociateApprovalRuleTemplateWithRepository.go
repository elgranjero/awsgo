package codecommit

// AssociateApprovalRuleTemplateWithRepository is generated as a reference stub.
// Executable command wiring lives under cmd/codecommit.go.
//
// Creates an association between an approval rule template and a specified
// repository. Then, the next time a pull request is created in the repository
// where the destination reference (if specified) matches the destination reference
// (branch) for the pull request, an approval rule that matches the template
// conditions is automatically created for that pull request. If no destination
// references are specified in the template, an approval rule that matches the
// template contents is created for all pull requests in that repository.
