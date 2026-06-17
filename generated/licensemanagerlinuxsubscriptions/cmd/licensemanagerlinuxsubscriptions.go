package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/licensemanagerlinuxsubscriptions"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// licensemanagerlinuxsubscriptionsCmd represents the licensemanagerlinuxsubscriptions command
var _licensemanagerlinuxsubscriptionsCmd = &cobra.Command{
	Use:   "licensemanagerlinuxsubscriptions",
	Short: "AWS licensemanagerlinuxsubscriptions CLI",
	Run: func(cmd *cobra.Command, args []string) {
		_awsOutput = resolveAWSOutput(_awsProfile, cmd.Flags().Changed("output"))
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := licensemanagerlinuxsubscriptions.NewFromConfig(cfg)
		if _licensemanagerlinuxsubscriptionsDeregisterSubscriptionProvider {
			licensemanagerlinuxsubscriptions_DeregisterSubscriptionProvider(cfg, client)
			return
		}
		if _licensemanagerlinuxsubscriptionsGetRegisteredSubscriptionProvider {
			licensemanagerlinuxsubscriptions_GetRegisteredSubscriptionProvider(cfg, client)
			return
		}
		if _licensemanagerlinuxsubscriptionsGetServiceSettings {
			licensemanagerlinuxsubscriptions_GetServiceSettings(cfg, client)
			return
		}
		if _licensemanagerlinuxsubscriptionsListLinuxSubscriptionInstances {
			licensemanagerlinuxsubscriptions_ListLinuxSubscriptionInstances(cfg, client)
			return
		}
		if _licensemanagerlinuxsubscriptionsListLinuxSubscriptions {
			licensemanagerlinuxsubscriptions_ListLinuxSubscriptions(cfg, client)
			return
		}
		if _licensemanagerlinuxsubscriptionsListRegisteredSubscriptionProviders {
			licensemanagerlinuxsubscriptions_ListRegisteredSubscriptionProviders(cfg, client)
			return
		}
		if _licensemanagerlinuxsubscriptionsListTagsForResource {
			licensemanagerlinuxsubscriptions_ListTagsForResource(cfg, client)
			return
		}
		if _licensemanagerlinuxsubscriptionsRegisterSubscriptionProvider {
			licensemanagerlinuxsubscriptions_RegisterSubscriptionProvider(cfg, client)
			return
		}
		if _licensemanagerlinuxsubscriptionsTagResource {
			licensemanagerlinuxsubscriptions_TagResource(cfg, client)
			return
		}
		if _licensemanagerlinuxsubscriptionsUntagResource {
			licensemanagerlinuxsubscriptions_UntagResource(cfg, client)
			return
		}
		if _licensemanagerlinuxsubscriptionsUpdateServiceSettings {
			licensemanagerlinuxsubscriptions_UpdateServiceSettings(cfg, client)
			return
		}

	},
}

var (
	_licensemanagerlinuxsubscriptionsDeregisterSubscriptionProvider      bool
	_licensemanagerlinuxsubscriptionsGetRegisteredSubscriptionProvider   bool
	_licensemanagerlinuxsubscriptionsGetServiceSettings                  bool
	_licensemanagerlinuxsubscriptionsListLinuxSubscriptionInstances      bool
	_licensemanagerlinuxsubscriptionsListLinuxSubscriptions              bool
	_licensemanagerlinuxsubscriptionsListRegisteredSubscriptionProviders bool
	_licensemanagerlinuxsubscriptionsListTagsForResource                 bool
	_licensemanagerlinuxsubscriptionsRegisterSubscriptionProvider        bool
	_licensemanagerlinuxsubscriptionsTagResource                         bool
	_licensemanagerlinuxsubscriptionsUntagResource                       bool
	_licensemanagerlinuxsubscriptionsUpdateServiceSettings               bool

	_licensemanagerlinuxsubscriptionsAllowUpdate                         string
	_licensemanagerlinuxsubscriptionsFilters                             string
	_licensemanagerlinuxsubscriptionsLinuxSubscriptionsDiscovery         string
	_licensemanagerlinuxsubscriptionsLinuxSubscriptionsDiscoverySettings string
	_licensemanagerlinuxsubscriptionsMaxResults                          string
	_licensemanagerlinuxsubscriptionsNextToken                           string
	_licensemanagerlinuxsubscriptionsResourceArn                         string
	_licensemanagerlinuxsubscriptionsSecretArn                           string
	_licensemanagerlinuxsubscriptionsSubscriptionProviderArn             string
	_licensemanagerlinuxsubscriptionsSubscriptionProviderSource          string
	_licensemanagerlinuxsubscriptionsSubscriptionProviderSources         string
	_licensemanagerlinuxsubscriptionsTagKeys                             []string
	_licensemanagerlinuxsubscriptionsTags                                string
)

// Remove a third-party subscription provider from the Bring Your Own License
// (BYOL) subscriptions registered to your account.
func licensemanagerlinuxsubscriptions_DeregisterSubscriptionProvider(cfg aws.Config, client *licensemanagerlinuxsubscriptions.Client) {
	input := &licensemanagerlinuxsubscriptions.DeregisterSubscriptionProviderInput{
		// SubscriptionProviderArn: *string, // Required
	}

	if len(_licensemanagerlinuxsubscriptionsSubscriptionProviderArn) > 0 {
		input.SubscriptionProviderArn = aws.String(_licensemanagerlinuxsubscriptionsSubscriptionProviderArn)
	}

	if resp, err := client.DeregisterSubscriptionProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get details for a Bring Your Own License (BYOL) subscription that's registered
// to your account.
func licensemanagerlinuxsubscriptions_GetRegisteredSubscriptionProvider(cfg aws.Config, client *licensemanagerlinuxsubscriptions.Client) {
	input := &licensemanagerlinuxsubscriptions.GetRegisteredSubscriptionProviderInput{
		// SubscriptionProviderArn: *string, // Required
	}

	if len(_licensemanagerlinuxsubscriptionsSubscriptionProviderArn) > 0 {
		input.SubscriptionProviderArn = aws.String(_licensemanagerlinuxsubscriptionsSubscriptionProviderArn)
	}

	if resp, err := client.GetRegisteredSubscriptionProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the Linux subscriptions service settings for your account.
func licensemanagerlinuxsubscriptions_GetServiceSettings(cfg aws.Config, client *licensemanagerlinuxsubscriptions.Client) {
	input := &licensemanagerlinuxsubscriptions.GetServiceSettingsInput{}

	if resp, err := client.GetServiceSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the running Amazon EC2 instances that were discovered with commercial
// Linux subscriptions.
func licensemanagerlinuxsubscriptions_ListLinuxSubscriptionInstances(cfg aws.Config, client *licensemanagerlinuxsubscriptions.Client) {
	input := &licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesInput{}

	if len(_licensemanagerlinuxsubscriptionsFilters) > 0 {
		if err := assignInputField(input, "Filters", _licensemanagerlinuxsubscriptionsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerlinuxsubscriptionsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _licensemanagerlinuxsubscriptionsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerlinuxsubscriptionsNextToken) > 0 {
		input.NextToken = aws.String(_licensemanagerlinuxsubscriptionsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLinuxSubscriptionInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesOutput
	p := licensemanagerlinuxsubscriptions.NewListLinuxSubscriptionInstancesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists the Linux subscriptions that have been discovered. If you have linked
// your organization, the returned results will include data aggregated across your
// accounts in Organizations.
func licensemanagerlinuxsubscriptions_ListLinuxSubscriptions(cfg aws.Config, client *licensemanagerlinuxsubscriptions.Client) {
	input := &licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsInput{}

	if len(_licensemanagerlinuxsubscriptionsFilters) > 0 {
		if err := assignInputField(input, "Filters", _licensemanagerlinuxsubscriptionsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerlinuxsubscriptionsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _licensemanagerlinuxsubscriptionsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerlinuxsubscriptionsNextToken) > 0 {
		input.NextToken = aws.String(_licensemanagerlinuxsubscriptionsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLinuxSubscriptions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsOutput
	p := licensemanagerlinuxsubscriptions.NewListLinuxSubscriptionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// List Bring Your Own License (BYOL) subscription registration resources for your
// account.
func licensemanagerlinuxsubscriptions_ListRegisteredSubscriptionProviders(cfg aws.Config, client *licensemanagerlinuxsubscriptions.Client) {
	input := &licensemanagerlinuxsubscriptions.ListRegisteredSubscriptionProvidersInput{}

	if len(_licensemanagerlinuxsubscriptionsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _licensemanagerlinuxsubscriptionsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerlinuxsubscriptionsNextToken) > 0 {
		input.NextToken = aws.String(_licensemanagerlinuxsubscriptionsNextToken)
	}
	if len(_licensemanagerlinuxsubscriptionsSubscriptionProviderSources) > 0 {
		if err := assignInputField(input, "SubscriptionProviderSources", _licensemanagerlinuxsubscriptionsSubscriptionProviderSources); err != nil {
			log.Errorf("invalid --subscription-provider-sources: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListRegisteredSubscriptionProviders(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*licensemanagerlinuxsubscriptions.ListRegisteredSubscriptionProvidersOutput
	p := licensemanagerlinuxsubscriptions.NewListRegisteredSubscriptionProvidersPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// List the metadata tags that are assigned to the specified Amazon Web Services
// resource.
func licensemanagerlinuxsubscriptions_ListTagsForResource(cfg aws.Config, client *licensemanagerlinuxsubscriptions.Client) {
	input := &licensemanagerlinuxsubscriptions.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_licensemanagerlinuxsubscriptionsResourceArn) > 0 {
		input.ResourceArn = aws.String(_licensemanagerlinuxsubscriptionsResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Register the supported third-party subscription provider for your Bring Your
// Own License (BYOL) subscription.
func licensemanagerlinuxsubscriptions_RegisterSubscriptionProvider(cfg aws.Config, client *licensemanagerlinuxsubscriptions.Client) {
	input := &licensemanagerlinuxsubscriptions.RegisterSubscriptionProviderInput{
		// SecretArn: *string, // Required
		// SubscriptionProviderSource: types.SubscriptionProviderSource, // Required
	}

	if len(_licensemanagerlinuxsubscriptionsSecretArn) > 0 {
		input.SecretArn = aws.String(_licensemanagerlinuxsubscriptionsSecretArn)
	}
	if len(_licensemanagerlinuxsubscriptionsSubscriptionProviderSource) > 0 {
		if err := assignInputField(input, "SubscriptionProviderSource", _licensemanagerlinuxsubscriptionsSubscriptionProviderSource); err != nil {
			log.Errorf("invalid --subscription-provider-source: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerlinuxsubscriptionsTags) > 0 {
		if err := assignInputField(input, "Tags", _licensemanagerlinuxsubscriptionsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.RegisterSubscriptionProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Add metadata tags to the specified Amazon Web Services resource.
func licensemanagerlinuxsubscriptions_TagResource(cfg aws.Config, client *licensemanagerlinuxsubscriptions.Client) {
	input := &licensemanagerlinuxsubscriptions.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_licensemanagerlinuxsubscriptionsResourceArn) > 0 {
		input.ResourceArn = aws.String(_licensemanagerlinuxsubscriptionsResourceArn)
	}
	if len(_licensemanagerlinuxsubscriptionsTags) > 0 {
		if err := assignInputField(input, "Tags", _licensemanagerlinuxsubscriptionsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Remove one or more metadata tag from the specified Amazon Web Services resource.
func licensemanagerlinuxsubscriptions_UntagResource(cfg aws.Config, client *licensemanagerlinuxsubscriptions.Client) {
	input := &licensemanagerlinuxsubscriptions.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_licensemanagerlinuxsubscriptionsResourceArn) > 0 {
		input.ResourceArn = aws.String(_licensemanagerlinuxsubscriptionsResourceArn)
	}
	if len(_licensemanagerlinuxsubscriptionsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _licensemanagerlinuxsubscriptionsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the service settings for Linux subscriptions.
func licensemanagerlinuxsubscriptions_UpdateServiceSettings(cfg aws.Config, client *licensemanagerlinuxsubscriptions.Client) {
	input := &licensemanagerlinuxsubscriptions.UpdateServiceSettingsInput{
		// LinuxSubscriptionsDiscovery: types.LinuxSubscriptionsDiscovery, // Required
		// LinuxSubscriptionsDiscoverySettings: *types.LinuxSubscriptionsDiscoverySettings, // Required
	}

	if len(_licensemanagerlinuxsubscriptionsLinuxSubscriptionsDiscovery) > 0 {
		if err := assignInputField(input, "LinuxSubscriptionsDiscovery", _licensemanagerlinuxsubscriptionsLinuxSubscriptionsDiscovery); err != nil {
			log.Errorf("invalid --linux-subscriptions-discovery: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerlinuxsubscriptionsLinuxSubscriptionsDiscoverySettings) > 0 {
		if err := assignInputField(input, "LinuxSubscriptionsDiscoverySettings", _licensemanagerlinuxsubscriptionsLinuxSubscriptionsDiscoverySettings); err != nil {
			log.Errorf("invalid --linux-subscriptions-discovery-settings: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerlinuxsubscriptionsAllowUpdate) > 0 {
		if err := assignInputField(input, "AllowUpdate", _licensemanagerlinuxsubscriptionsAllowUpdate); err != nil {
			log.Errorf("invalid --allow-update: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateServiceSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_licensemanagerlinuxsubscriptionsCmd)
	_licensemanagerlinuxsubscriptionsCmd.Flags().SortFlags = false

	_licensemanagerlinuxsubscriptionsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_licensemanagerlinuxsubscriptionsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_licensemanagerlinuxsubscriptionsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_licensemanagerlinuxsubscriptionsCmd.Flags().StringVarP(&_licensemanagerlinuxsubscriptionsAllowUpdate, "allow-update", "", "", "Allow Update")
	_licensemanagerlinuxsubscriptionsCmd.Flags().StringVarP(&_licensemanagerlinuxsubscriptionsFilters, "filters", "", "", "Filters")
	_licensemanagerlinuxsubscriptionsCmd.Flags().StringVarP(&_licensemanagerlinuxsubscriptionsLinuxSubscriptionsDiscovery, "linux-subscriptions-discovery", "", "", "Linux Subscriptions Discovery")
	_licensemanagerlinuxsubscriptionsCmd.Flags().StringVarP(&_licensemanagerlinuxsubscriptionsLinuxSubscriptionsDiscoverySettings, "linux-subscriptions-discovery-settings", "", "", "Linux Subscriptions Discovery Settings")
	_licensemanagerlinuxsubscriptionsCmd.Flags().StringVarP(&_licensemanagerlinuxsubscriptionsMaxResults, "max-results", "", "", "Max Results")
	_licensemanagerlinuxsubscriptionsCmd.Flags().StringVarP(&_licensemanagerlinuxsubscriptionsNextToken, "next-token", "", "", "Next Token")
	_licensemanagerlinuxsubscriptionsCmd.Flags().StringVarP(&_licensemanagerlinuxsubscriptionsResourceArn, "resource-arn", "", "", "Resource ARN")
	_licensemanagerlinuxsubscriptionsCmd.Flags().StringVarP(&_licensemanagerlinuxsubscriptionsSecretArn, "secret-arn", "", "", "Secret ARN")
	_licensemanagerlinuxsubscriptionsCmd.Flags().StringVarP(&_licensemanagerlinuxsubscriptionsSubscriptionProviderArn, "subscription-provider-arn", "", "", "Subscription Provider ARN")
	_licensemanagerlinuxsubscriptionsCmd.Flags().StringVarP(&_licensemanagerlinuxsubscriptionsSubscriptionProviderSource, "subscription-provider-source", "", "", "Subscription Provider Source")
	_licensemanagerlinuxsubscriptionsCmd.Flags().StringVarP(&_licensemanagerlinuxsubscriptionsSubscriptionProviderSources, "subscription-provider-sources", "", "", "Subscription Provider Sources")
	_licensemanagerlinuxsubscriptionsCmd.Flags().StringSliceVarP(&_licensemanagerlinuxsubscriptionsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_licensemanagerlinuxsubscriptionsCmd.Flags().StringVarP(&_licensemanagerlinuxsubscriptionsTags, "tags", "", "", "Tags")

	_licensemanagerlinuxsubscriptionsCmd.Flags().BoolVarP(&_licensemanagerlinuxsubscriptionsDeregisterSubscriptionProvider, "deregister-subscription-provider", "", false, "Deregister Subscription Provider")
	_licensemanagerlinuxsubscriptionsCmd.Flags().BoolVarP(&_licensemanagerlinuxsubscriptionsGetRegisteredSubscriptionProvider, "get-registered-subscription-provider", "", false, "Get Registered Subscription Provider")
	_licensemanagerlinuxsubscriptionsCmd.Flags().BoolVarP(&_licensemanagerlinuxsubscriptionsGetServiceSettings, "get-service-settings", "", false, "Get Service Settings")
	_licensemanagerlinuxsubscriptionsCmd.Flags().BoolVarP(&_licensemanagerlinuxsubscriptionsListLinuxSubscriptionInstances, "list-linux-subscription-instances", "", false, "List Linux Subscription Instances")
	_licensemanagerlinuxsubscriptionsCmd.Flags().BoolVarP(&_licensemanagerlinuxsubscriptionsListLinuxSubscriptions, "list-linux-subscriptions", "", false, "List Linux Subscriptions")
	_licensemanagerlinuxsubscriptionsCmd.Flags().BoolVarP(&_licensemanagerlinuxsubscriptionsListRegisteredSubscriptionProviders, "list-registered-subscription-providers", "", false, "List Registered Subscription Providers")
	_licensemanagerlinuxsubscriptionsCmd.Flags().BoolVarP(&_licensemanagerlinuxsubscriptionsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_licensemanagerlinuxsubscriptionsCmd.Flags().BoolVarP(&_licensemanagerlinuxsubscriptionsRegisterSubscriptionProvider, "register-subscription-provider", "", false, "Register Subscription Provider")
	_licensemanagerlinuxsubscriptionsCmd.Flags().BoolVarP(&_licensemanagerlinuxsubscriptionsTagResource, "tag-resource", "", false, "Tag Resource")
	_licensemanagerlinuxsubscriptionsCmd.Flags().BoolVarP(&_licensemanagerlinuxsubscriptionsUntagResource, "untag-resource", "", false, "Untag Resource")
	_licensemanagerlinuxsubscriptionsCmd.Flags().BoolVarP(&_licensemanagerlinuxsubscriptionsUpdateServiceSettings, "update-service-settings", "", false, "Update Service Settings")

}
