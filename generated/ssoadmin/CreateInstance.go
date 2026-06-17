package ssoadmin

// CreateInstance is generated as a reference stub.
// Executable command wiring lives under cmd/ssoadmin.go.
//
// Creates an instance of IAM Identity Center for a standalone Amazon Web Services
// account that is not managed by Organizations or a member Amazon Web Services
// account in an organization. You can create only one instance per account and
// across all Amazon Web Services Regions.
//
// The CreateInstance request is rejected if the following apply:
//
// - The instance is created within the organization management account.
//
// - An instance already exists in the same account.
