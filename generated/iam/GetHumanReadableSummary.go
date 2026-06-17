package iam

// GetHumanReadableSummary is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Retrieves a human readable summary for a given entity. At this time, the only
// supported entity type is delegation-request
//
// This method uses a Large Language Model (LLM) to generate the summary.
//
// If a delegation request has no owner or owner account, GetHumanReadableSummary
// for that delegation request can be called by any account. If the owner account
// is assigned but there is no owner id, only identities within that owner account
// can call GetHumanReadableSummary for the delegation request to retrieve a
// summary of that request. Once the delegation request is fully owned, the owner
// of the request gets a default permission to get that delegation request. For
// more details, read default permissions granted to delegation requests. These rules are identical to [GetDelegationRequest] API behavior, such that a
// party who has permissions to call [GetDelegationRequest]for a given delegation request will always be
// able to retrieve the human readable summary for that request.
//
// [GetDelegationRequest]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetDelegationRequest.html
