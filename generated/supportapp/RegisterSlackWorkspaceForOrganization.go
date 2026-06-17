package supportapp

// RegisterSlackWorkspaceForOrganization is generated as a reference stub.
// Executable command wiring lives under cmd/supportapp.go.
//
// Registers a Slack workspace for your Amazon Web Services account. To call this
// API, your account must be part of an organization in Organizations.
//
// If you're the management account and you want to register Slack workspaces for
// your organization, you must complete the following tasks:
//
// - Sign in to the [Amazon Web Services Support Center]and authorize the Slack workspaces where you want your
// organization to have access to. See [Authorize a Slack workspace]in the Amazon Web Services Support User
// Guide.
//
// - Call the RegisterSlackWorkspaceForOrganization API to authorize each Slack
// workspace for the organization.
//
// After the management account authorizes the Slack workspace, member accounts
// can call this API to authorize the same Slack workspace for their individual
// accounts. Member accounts don't need to authorize the Slack workspace manually
// through the [Amazon Web Services Support Center].
//
// To use the Amazon Web Services Support App, each account must then complete the
// following tasks:
//
// - Create an Identity and Access Management (IAM) role with the required
// permission. For more information, see [Managing access to the Amazon Web Services Support App].
//
// - Configure a Slack channel to use the Amazon Web Services Support App for
// support cases for that account. For more information, see [Configuring a Slack channel].
//
// [Authorize a Slack workspace]: https://docs.aws.amazon.com/awssupport/latest/user/authorize-slack-workspace.html
// [Managing access to the Amazon Web Services Support App]: https://docs.aws.amazon.com/awssupport/latest/user/support-app-permissions.html
// [Amazon Web Services Support Center]: https://console.aws.amazon.com/support/app
// [Configuring a Slack channel]: https://docs.aws.amazon.com/awssupport/latest/user/add-your-slack-channel.html
