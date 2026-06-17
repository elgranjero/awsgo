package configservice

// PutConfigurationAggregator is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Creates and updates the configuration aggregator with the selected source
// accounts and regions. The source account can be individual account(s) or an
// organization.
//
// accountIds that are passed will be replaced with existing accounts. If you want
// to add additional accounts into the aggregator, call
// DescribeConfigurationAggregators to get the previous accounts and then append
// new ones.
//
// Config should be enabled in source accounts and regions you want to aggregate.
//
// If your source type is an organization, you must be signed in to the management
// account or a registered delegated administrator and all the features must be
// enabled in your organization. If the caller is a management account, Config
// calls EnableAwsServiceAccess API to enable integration between Config and
// Organizations. If the caller is a registered delegated administrator, Config
// calls ListDelegatedAdministrators API to verify whether the caller is a valid
// delegated administrator.
//
// To register a delegated administrator, see [Register a Delegated Administrator] in the Config developer guide.
//
// # Tags are added at creation and cannot be updated with this operation
//
// PutConfigurationAggregator is an idempotent API. Subsequent requests won’t
// create a duplicate resource if one was already created. If a following request
// has different tags values, Config will ignore these differences and treat it as
// an idempotent request of the previous. In this case, tags will not be updated,
// even if they are different.
//
// Use [TagResource] and [UntagResource] to update tags after creation.
//
// [Register a Delegated Administrator]: https://docs.aws.amazon.com/config/latest/developerguide/set-up-aggregator-cli.html#register-a-delegated-administrator-cli
// [TagResource]: https://docs.aws.amazon.com/config/latest/APIReference/API_TagResource.html
// [UntagResource]: https://docs.aws.amazon.com/config/latest/APIReference/API_UntagResource.html
