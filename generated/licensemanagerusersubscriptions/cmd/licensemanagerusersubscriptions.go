package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/licensemanagerusersubscriptions"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// licensemanagerusersubscriptionsCmd represents the licensemanagerusersubscriptions command
var _licensemanagerusersubscriptionsCmd = &cobra.Command{
	Use:   "licensemanagerusersubscriptions",
	Short: "AWS licensemanagerusersubscriptions CLI",
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
		client := licensemanagerusersubscriptions.NewFromConfig(cfg)
		if _licensemanagerusersubscriptionsAssociateUser {
			licensemanagerusersubscriptions_AssociateUser(cfg, client)
			return
		}
		if _licensemanagerusersubscriptionsCreateLicenseServerEndpoint {
			licensemanagerusersubscriptions_CreateLicenseServerEndpoint(cfg, client)
			return
		}
		if _licensemanagerusersubscriptionsDeleteLicenseServerEndpoint {
			licensemanagerusersubscriptions_DeleteLicenseServerEndpoint(cfg, client)
			return
		}
		if _licensemanagerusersubscriptionsDeregisterIdentityProvider {
			licensemanagerusersubscriptions_DeregisterIdentityProvider(cfg, client)
			return
		}
		if _licensemanagerusersubscriptionsDisassociateUser {
			licensemanagerusersubscriptions_DisassociateUser(cfg, client)
			return
		}
		if _licensemanagerusersubscriptionsListIdentityProviders {
			licensemanagerusersubscriptions_ListIdentityProviders(cfg, client)
			return
		}
		if _licensemanagerusersubscriptionsListInstances {
			licensemanagerusersubscriptions_ListInstances(cfg, client)
			return
		}
		if _licensemanagerusersubscriptionsListLicenseServerEndpoints {
			licensemanagerusersubscriptions_ListLicenseServerEndpoints(cfg, client)
			return
		}
		if _licensemanagerusersubscriptionsListProductSubscriptions {
			licensemanagerusersubscriptions_ListProductSubscriptions(cfg, client)
			return
		}
		if _licensemanagerusersubscriptionsListTagsForResource {
			licensemanagerusersubscriptions_ListTagsForResource(cfg, client)
			return
		}
		if _licensemanagerusersubscriptionsListUserAssociations {
			licensemanagerusersubscriptions_ListUserAssociations(cfg, client)
			return
		}
		if _licensemanagerusersubscriptionsRegisterIdentityProvider {
			licensemanagerusersubscriptions_RegisterIdentityProvider(cfg, client)
			return
		}
		if _licensemanagerusersubscriptionsStartProductSubscription {
			licensemanagerusersubscriptions_StartProductSubscription(cfg, client)
			return
		}
		if _licensemanagerusersubscriptionsStopProductSubscription {
			licensemanagerusersubscriptions_StopProductSubscription(cfg, client)
			return
		}
		if _licensemanagerusersubscriptionsTagResource {
			licensemanagerusersubscriptions_TagResource(cfg, client)
			return
		}
		if _licensemanagerusersubscriptionsUntagResource {
			licensemanagerusersubscriptions_UntagResource(cfg, client)
			return
		}
		if _licensemanagerusersubscriptionsUpdateIdentityProviderSettings {
			licensemanagerusersubscriptions_UpdateIdentityProviderSettings(cfg, client)
			return
		}

	},
}

var (
	_licensemanagerusersubscriptionsAssociateUser                  bool
	_licensemanagerusersubscriptionsCreateLicenseServerEndpoint    bool
	_licensemanagerusersubscriptionsDeleteLicenseServerEndpoint    bool
	_licensemanagerusersubscriptionsDeregisterIdentityProvider     bool
	_licensemanagerusersubscriptionsDisassociateUser               bool
	_licensemanagerusersubscriptionsListIdentityProviders          bool
	_licensemanagerusersubscriptionsListInstances                  bool
	_licensemanagerusersubscriptionsListLicenseServerEndpoints     bool
	_licensemanagerusersubscriptionsListProductSubscriptions       bool
	_licensemanagerusersubscriptionsListTagsForResource            bool
	_licensemanagerusersubscriptionsListUserAssociations           bool
	_licensemanagerusersubscriptionsRegisterIdentityProvider       bool
	_licensemanagerusersubscriptionsStartProductSubscription       bool
	_licensemanagerusersubscriptionsStopProductSubscription        bool
	_licensemanagerusersubscriptionsTagResource                    bool
	_licensemanagerusersubscriptionsUntagResource                  bool
	_licensemanagerusersubscriptionsUpdateIdentityProviderSettings bool

	_licensemanagerusersubscriptionsDomain                   string
	_licensemanagerusersubscriptionsFilters                  string
	_licensemanagerusersubscriptionsIdentityProvider         string
	_licensemanagerusersubscriptionsIdentityProviderArn      string
	_licensemanagerusersubscriptionsInstanceId               string
	_licensemanagerusersubscriptionsInstanceUserArn          string
	_licensemanagerusersubscriptionsLicenseServerEndpointArn string
	_licensemanagerusersubscriptionsLicenseServerSettings    string
	_licensemanagerusersubscriptionsMaxResults               string
	_licensemanagerusersubscriptionsNextToken                string
	_licensemanagerusersubscriptionsProduct                  string
	_licensemanagerusersubscriptionsProductUserArn           string
	_licensemanagerusersubscriptionsResourceArn              string
	_licensemanagerusersubscriptionsServerType               string
	_licensemanagerusersubscriptionsSettings                 string
	_licensemanagerusersubscriptionsTagKeys                  []string
	_licensemanagerusersubscriptionsTags                     string
	_licensemanagerusersubscriptionsUpdateSettings           string
	_licensemanagerusersubscriptionsUsername                 string
)

// Associates the user to an EC2 instance to utilize user-based subscriptions.
// Your estimated bill for charges on the number of users and related costs will
// take 48 hours to appear for billing periods that haven't closed (marked as
// Pending billing status) in Amazon Web Services Billing. For more information,
// see [Viewing your monthly charges]in the Amazon Web Services Billing User Guide.
//
// [Viewing your monthly charges]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/invoice.html
func licensemanagerusersubscriptions_AssociateUser(cfg aws.Config, client *licensemanagerusersubscriptions.Client) {
	input := &licensemanagerusersubscriptions.AssociateUserInput{
		// IdentityProvider: types.IdentityProvider, // Required
		// InstanceId: *string, // Required
		// Username: *string, // Required
	}

	if len(_licensemanagerusersubscriptionsIdentityProvider) > 0 {
		if err := assignInputField(input, "IdentityProvider", _licensemanagerusersubscriptionsIdentityProvider); err != nil {
			log.Errorf("invalid --identity-provider: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerusersubscriptionsInstanceId) > 0 {
		input.InstanceId = aws.String(_licensemanagerusersubscriptionsInstanceId)
	}
	if len(_licensemanagerusersubscriptionsUsername) > 0 {
		input.Username = aws.String(_licensemanagerusersubscriptionsUsername)
	}
	if len(_licensemanagerusersubscriptionsDomain) > 0 {
		input.Domain = aws.String(_licensemanagerusersubscriptionsDomain)
	}
	if len(_licensemanagerusersubscriptionsTags) > 0 {
		if err := assignInputField(input, "Tags", _licensemanagerusersubscriptionsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a network endpoint for the Remote Desktop Services (RDS) license server.
func licensemanagerusersubscriptions_CreateLicenseServerEndpoint(cfg aws.Config, client *licensemanagerusersubscriptions.Client) {
	input := &licensemanagerusersubscriptions.CreateLicenseServerEndpointInput{
		// IdentityProviderArn: *string, // Required
		// LicenseServerSettings: *types.LicenseServerSettings, // Required
	}

	if len(_licensemanagerusersubscriptionsIdentityProviderArn) > 0 {
		input.IdentityProviderArn = aws.String(_licensemanagerusersubscriptionsIdentityProviderArn)
	}
	if len(_licensemanagerusersubscriptionsLicenseServerSettings) > 0 {
		if err := assignInputField(input, "LicenseServerSettings", _licensemanagerusersubscriptionsLicenseServerSettings); err != nil {
			log.Errorf("invalid --license-server-settings: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerusersubscriptionsTags) > 0 {
		if err := assignInputField(input, "Tags", _licensemanagerusersubscriptionsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLicenseServerEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a LicenseServerEndpoint resource.
func licensemanagerusersubscriptions_DeleteLicenseServerEndpoint(cfg aws.Config, client *licensemanagerusersubscriptions.Client) {
	input := &licensemanagerusersubscriptions.DeleteLicenseServerEndpointInput{
		// LicenseServerEndpointArn: *string, // Required
		// ServerType: types.ServerType, // Required
	}

	if len(_licensemanagerusersubscriptionsLicenseServerEndpointArn) > 0 {
		input.LicenseServerEndpointArn = aws.String(_licensemanagerusersubscriptionsLicenseServerEndpointArn)
	}
	if len(_licensemanagerusersubscriptionsServerType) > 0 {
		if err := assignInputField(input, "ServerType", _licensemanagerusersubscriptionsServerType); err != nil {
			log.Errorf("invalid --server-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteLicenseServerEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deregisters the Active Directory identity provider from License Manager
// user-based subscriptions.
func licensemanagerusersubscriptions_DeregisterIdentityProvider(cfg aws.Config, client *licensemanagerusersubscriptions.Client) {
	input := &licensemanagerusersubscriptions.DeregisterIdentityProviderInput{}

	if len(_licensemanagerusersubscriptionsIdentityProvider) > 0 {
		if err := assignInputField(input, "IdentityProvider", _licensemanagerusersubscriptionsIdentityProvider); err != nil {
			log.Errorf("invalid --identity-provider: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerusersubscriptionsIdentityProviderArn) > 0 {
		input.IdentityProviderArn = aws.String(_licensemanagerusersubscriptionsIdentityProviderArn)
	}
	if len(_licensemanagerusersubscriptionsProduct) > 0 {
		input.Product = aws.String(_licensemanagerusersubscriptionsProduct)
	}

	if resp, err := client.DeregisterIdentityProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates the user from an EC2 instance providing user-based subscriptions.
func licensemanagerusersubscriptions_DisassociateUser(cfg aws.Config, client *licensemanagerusersubscriptions.Client) {
	input := &licensemanagerusersubscriptions.DisassociateUserInput{}

	if len(_licensemanagerusersubscriptionsDomain) > 0 {
		input.Domain = aws.String(_licensemanagerusersubscriptionsDomain)
	}
	if len(_licensemanagerusersubscriptionsIdentityProvider) > 0 {
		if err := assignInputField(input, "IdentityProvider", _licensemanagerusersubscriptionsIdentityProvider); err != nil {
			log.Errorf("invalid --identity-provider: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerusersubscriptionsInstanceId) > 0 {
		input.InstanceId = aws.String(_licensemanagerusersubscriptionsInstanceId)
	}
	if len(_licensemanagerusersubscriptionsInstanceUserArn) > 0 {
		input.InstanceUserArn = aws.String(_licensemanagerusersubscriptionsInstanceUserArn)
	}
	if len(_licensemanagerusersubscriptionsUsername) > 0 {
		input.Username = aws.String(_licensemanagerusersubscriptionsUsername)
	}

	if resp, err := client.DisassociateUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the Active Directory identity providers for user-based subscriptions.
func licensemanagerusersubscriptions_ListIdentityProviders(cfg aws.Config, client *licensemanagerusersubscriptions.Client) {
	input := &licensemanagerusersubscriptions.ListIdentityProvidersInput{}

	if len(_licensemanagerusersubscriptionsFilters) > 0 {
		if err := assignInputField(input, "Filters", _licensemanagerusersubscriptionsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerusersubscriptionsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _licensemanagerusersubscriptionsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerusersubscriptionsNextToken) > 0 {
		input.NextToken = aws.String(_licensemanagerusersubscriptionsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIdentityProviders(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*licensemanagerusersubscriptions.ListIdentityProvidersOutput
	p := licensemanagerusersubscriptions.NewListIdentityProvidersPaginator(client, input)
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

// Lists the EC2 instances providing user-based subscriptions.
func licensemanagerusersubscriptions_ListInstances(cfg aws.Config, client *licensemanagerusersubscriptions.Client) {
	input := &licensemanagerusersubscriptions.ListInstancesInput{}

	if len(_licensemanagerusersubscriptionsFilters) > 0 {
		if err := assignInputField(input, "Filters", _licensemanagerusersubscriptionsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerusersubscriptionsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _licensemanagerusersubscriptionsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerusersubscriptionsNextToken) > 0 {
		input.NextToken = aws.String(_licensemanagerusersubscriptionsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*licensemanagerusersubscriptions.ListInstancesOutput
	p := licensemanagerusersubscriptions.NewListInstancesPaginator(client, input)
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

// List the Remote Desktop Services (RDS) License Server endpoints
func licensemanagerusersubscriptions_ListLicenseServerEndpoints(cfg aws.Config, client *licensemanagerusersubscriptions.Client) {
	input := &licensemanagerusersubscriptions.ListLicenseServerEndpointsInput{}

	if len(_licensemanagerusersubscriptionsFilters) > 0 {
		if err := assignInputField(input, "Filters", _licensemanagerusersubscriptionsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerusersubscriptionsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _licensemanagerusersubscriptionsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerusersubscriptionsNextToken) > 0 {
		input.NextToken = aws.String(_licensemanagerusersubscriptionsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLicenseServerEndpoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*licensemanagerusersubscriptions.ListLicenseServerEndpointsOutput
	p := licensemanagerusersubscriptions.NewListLicenseServerEndpointsPaginator(client, input)
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

// Lists the user-based subscription products available from an identity provider.
func licensemanagerusersubscriptions_ListProductSubscriptions(cfg aws.Config, client *licensemanagerusersubscriptions.Client) {
	input := &licensemanagerusersubscriptions.ListProductSubscriptionsInput{
		// IdentityProvider: types.IdentityProvider, // Required
	}

	if len(_licensemanagerusersubscriptionsIdentityProvider) > 0 {
		if err := assignInputField(input, "IdentityProvider", _licensemanagerusersubscriptionsIdentityProvider); err != nil {
			log.Errorf("invalid --identity-provider: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerusersubscriptionsFilters) > 0 {
		if err := assignInputField(input, "Filters", _licensemanagerusersubscriptionsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerusersubscriptionsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _licensemanagerusersubscriptionsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerusersubscriptionsNextToken) > 0 {
		input.NextToken = aws.String(_licensemanagerusersubscriptionsNextToken)
	}
	if len(_licensemanagerusersubscriptionsProduct) > 0 {
		input.Product = aws.String(_licensemanagerusersubscriptionsProduct)
	}

	if disablePaginator() {
		if resp, err := client.ListProductSubscriptions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*licensemanagerusersubscriptions.ListProductSubscriptionsOutput
	p := licensemanagerusersubscriptions.NewListProductSubscriptionsPaginator(client, input)
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

// Returns the list of tags for the specified resource.
func licensemanagerusersubscriptions_ListTagsForResource(cfg aws.Config, client *licensemanagerusersubscriptions.Client) {
	input := &licensemanagerusersubscriptions.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_licensemanagerusersubscriptionsResourceArn) > 0 {
		input.ResourceArn = aws.String(_licensemanagerusersubscriptionsResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists user associations for an identity provider.
func licensemanagerusersubscriptions_ListUserAssociations(cfg aws.Config, client *licensemanagerusersubscriptions.Client) {
	input := &licensemanagerusersubscriptions.ListUserAssociationsInput{
		// IdentityProvider: types.IdentityProvider, // Required
		// InstanceId: *string, // Required
	}

	if len(_licensemanagerusersubscriptionsIdentityProvider) > 0 {
		if err := assignInputField(input, "IdentityProvider", _licensemanagerusersubscriptionsIdentityProvider); err != nil {
			log.Errorf("invalid --identity-provider: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerusersubscriptionsInstanceId) > 0 {
		input.InstanceId = aws.String(_licensemanagerusersubscriptionsInstanceId)
	}
	if len(_licensemanagerusersubscriptionsFilters) > 0 {
		if err := assignInputField(input, "Filters", _licensemanagerusersubscriptionsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerusersubscriptionsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _licensemanagerusersubscriptionsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerusersubscriptionsNextToken) > 0 {
		input.NextToken = aws.String(_licensemanagerusersubscriptionsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListUserAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*licensemanagerusersubscriptions.ListUserAssociationsOutput
	p := licensemanagerusersubscriptions.NewListUserAssociationsPaginator(client, input)
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

// Registers an identity provider for user-based subscriptions.
func licensemanagerusersubscriptions_RegisterIdentityProvider(cfg aws.Config, client *licensemanagerusersubscriptions.Client) {
	input := &licensemanagerusersubscriptions.RegisterIdentityProviderInput{
		// IdentityProvider: types.IdentityProvider, // Required
		// Product: *string, // Required
	}

	if len(_licensemanagerusersubscriptionsIdentityProvider) > 0 {
		if err := assignInputField(input, "IdentityProvider", _licensemanagerusersubscriptionsIdentityProvider); err != nil {
			log.Errorf("invalid --identity-provider: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerusersubscriptionsProduct) > 0 {
		input.Product = aws.String(_licensemanagerusersubscriptionsProduct)
	}
	if len(_licensemanagerusersubscriptionsSettings) > 0 {
		if err := assignInputField(input, "Settings", _licensemanagerusersubscriptionsSettings); err != nil {
			log.Errorf("invalid --settings: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerusersubscriptionsTags) > 0 {
		if err := assignInputField(input, "Tags", _licensemanagerusersubscriptionsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.RegisterIdentityProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a product subscription for a user with the specified identity provider.
// Your estimated bill for charges on the number of users and related costs will
// take 48 hours to appear for billing periods that haven't closed (marked as
// Pending billing status) in Amazon Web Services Billing. For more information,
// see [Viewing your monthly charges]in the Amazon Web Services Billing User Guide.
//
// [Viewing your monthly charges]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/invoice.html
func licensemanagerusersubscriptions_StartProductSubscription(cfg aws.Config, client *licensemanagerusersubscriptions.Client) {
	input := &licensemanagerusersubscriptions.StartProductSubscriptionInput{
		// IdentityProvider: types.IdentityProvider, // Required
		// Product: *string, // Required
		// Username: *string, // Required
	}

	if len(_licensemanagerusersubscriptionsIdentityProvider) > 0 {
		if err := assignInputField(input, "IdentityProvider", _licensemanagerusersubscriptionsIdentityProvider); err != nil {
			log.Errorf("invalid --identity-provider: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerusersubscriptionsProduct) > 0 {
		input.Product = aws.String(_licensemanagerusersubscriptionsProduct)
	}
	if len(_licensemanagerusersubscriptionsUsername) > 0 {
		input.Username = aws.String(_licensemanagerusersubscriptionsUsername)
	}
	if len(_licensemanagerusersubscriptionsDomain) > 0 {
		input.Domain = aws.String(_licensemanagerusersubscriptionsDomain)
	}
	if len(_licensemanagerusersubscriptionsTags) > 0 {
		if err := assignInputField(input, "Tags", _licensemanagerusersubscriptionsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartProductSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a product subscription for a user with the specified identity provider.
func licensemanagerusersubscriptions_StopProductSubscription(cfg aws.Config, client *licensemanagerusersubscriptions.Client) {
	input := &licensemanagerusersubscriptions.StopProductSubscriptionInput{}

	if len(_licensemanagerusersubscriptionsDomain) > 0 {
		input.Domain = aws.String(_licensemanagerusersubscriptionsDomain)
	}
	if len(_licensemanagerusersubscriptionsIdentityProvider) > 0 {
		if err := assignInputField(input, "IdentityProvider", _licensemanagerusersubscriptionsIdentityProvider); err != nil {
			log.Errorf("invalid --identity-provider: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerusersubscriptionsProduct) > 0 {
		input.Product = aws.String(_licensemanagerusersubscriptionsProduct)
	}
	if len(_licensemanagerusersubscriptionsProductUserArn) > 0 {
		input.ProductUserArn = aws.String(_licensemanagerusersubscriptionsProductUserArn)
	}
	if len(_licensemanagerusersubscriptionsUsername) > 0 {
		input.Username = aws.String(_licensemanagerusersubscriptionsUsername)
	}

	if resp, err := client.StopProductSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds tags to a resource.
func licensemanagerusersubscriptions_TagResource(cfg aws.Config, client *licensemanagerusersubscriptions.Client) {
	input := &licensemanagerusersubscriptions.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_licensemanagerusersubscriptionsResourceArn) > 0 {
		input.ResourceArn = aws.String(_licensemanagerusersubscriptionsResourceArn)
	}
	if len(_licensemanagerusersubscriptionsTags) > 0 {
		if err := assignInputField(input, "Tags", _licensemanagerusersubscriptionsTags); err != nil {
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

// Removes tags from a resource.
func licensemanagerusersubscriptions_UntagResource(cfg aws.Config, client *licensemanagerusersubscriptions.Client) {
	input := &licensemanagerusersubscriptions.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_licensemanagerusersubscriptionsResourceArn) > 0 {
		input.ResourceArn = aws.String(_licensemanagerusersubscriptionsResourceArn)
	}
	if len(_licensemanagerusersubscriptionsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _licensemanagerusersubscriptionsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates additional product configuration settings for the registered identity
// provider.
func licensemanagerusersubscriptions_UpdateIdentityProviderSettings(cfg aws.Config, client *licensemanagerusersubscriptions.Client) {
	input := &licensemanagerusersubscriptions.UpdateIdentityProviderSettingsInput{
		// UpdateSettings: *types.UpdateSettings, // Required
	}

	if len(_licensemanagerusersubscriptionsUpdateSettings) > 0 {
		if err := assignInputField(input, "UpdateSettings", _licensemanagerusersubscriptionsUpdateSettings); err != nil {
			log.Errorf("invalid --update-settings: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerusersubscriptionsIdentityProvider) > 0 {
		if err := assignInputField(input, "IdentityProvider", _licensemanagerusersubscriptionsIdentityProvider); err != nil {
			log.Errorf("invalid --identity-provider: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerusersubscriptionsIdentityProviderArn) > 0 {
		input.IdentityProviderArn = aws.String(_licensemanagerusersubscriptionsIdentityProviderArn)
	}
	if len(_licensemanagerusersubscriptionsProduct) > 0 {
		input.Product = aws.String(_licensemanagerusersubscriptionsProduct)
	}

	if resp, err := client.UpdateIdentityProviderSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_licensemanagerusersubscriptionsCmd)
	_licensemanagerusersubscriptionsCmd.Flags().SortFlags = false

	_licensemanagerusersubscriptionsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_licensemanagerusersubscriptionsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_licensemanagerusersubscriptionsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_licensemanagerusersubscriptionsCmd.Flags().StringVarP(&_licensemanagerusersubscriptionsDomain, "domain", "", "", "Domain")
	_licensemanagerusersubscriptionsCmd.Flags().StringVarP(&_licensemanagerusersubscriptionsFilters, "filters", "", "", "Filters")
	_licensemanagerusersubscriptionsCmd.Flags().StringVarP(&_licensemanagerusersubscriptionsIdentityProvider, "identity-provider", "", "", "Identity Provider")
	_licensemanagerusersubscriptionsCmd.Flags().StringVarP(&_licensemanagerusersubscriptionsIdentityProviderArn, "identity-provider-arn", "", "", "Identity Provider ARN")
	_licensemanagerusersubscriptionsCmd.Flags().StringVarP(&_licensemanagerusersubscriptionsInstanceId, "instance-id", "", "", "Instance ID")
	_licensemanagerusersubscriptionsCmd.Flags().StringVarP(&_licensemanagerusersubscriptionsInstanceUserArn, "instance-user-arn", "", "", "Instance User ARN")
	_licensemanagerusersubscriptionsCmd.Flags().StringVarP(&_licensemanagerusersubscriptionsLicenseServerEndpointArn, "license-server-endpoint-arn", "", "", "License Server Endpoint ARN")
	_licensemanagerusersubscriptionsCmd.Flags().StringVarP(&_licensemanagerusersubscriptionsLicenseServerSettings, "license-server-settings", "", "", "License Server Settings")
	_licensemanagerusersubscriptionsCmd.Flags().StringVarP(&_licensemanagerusersubscriptionsMaxResults, "max-results", "", "", "Max Results")
	_licensemanagerusersubscriptionsCmd.Flags().StringVarP(&_licensemanagerusersubscriptionsNextToken, "next-token", "", "", "Next Token")
	_licensemanagerusersubscriptionsCmd.Flags().StringVarP(&_licensemanagerusersubscriptionsProduct, "product", "", "", "Product")
	_licensemanagerusersubscriptionsCmd.Flags().StringVarP(&_licensemanagerusersubscriptionsProductUserArn, "product-user-arn", "", "", "Product User ARN")
	_licensemanagerusersubscriptionsCmd.Flags().StringVarP(&_licensemanagerusersubscriptionsResourceArn, "resource-arn", "", "", "Resource ARN")
	_licensemanagerusersubscriptionsCmd.Flags().StringVarP(&_licensemanagerusersubscriptionsServerType, "server-type", "", "", "Server Type")
	_licensemanagerusersubscriptionsCmd.Flags().StringVarP(&_licensemanagerusersubscriptionsSettings, "settings", "", "", "Settings")
	_licensemanagerusersubscriptionsCmd.Flags().StringSliceVarP(&_licensemanagerusersubscriptionsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_licensemanagerusersubscriptionsCmd.Flags().StringVarP(&_licensemanagerusersubscriptionsTags, "tags", "", "", "Tags")
	_licensemanagerusersubscriptionsCmd.Flags().StringVarP(&_licensemanagerusersubscriptionsUpdateSettings, "update-settings", "", "", "Update Settings")
	_licensemanagerusersubscriptionsCmd.Flags().StringVarP(&_licensemanagerusersubscriptionsUsername, "username", "", "", "Username")

	_licensemanagerusersubscriptionsCmd.Flags().BoolVarP(&_licensemanagerusersubscriptionsAssociateUser, "associate-user", "", false, "Associate User")
	_licensemanagerusersubscriptionsCmd.Flags().BoolVarP(&_licensemanagerusersubscriptionsCreateLicenseServerEndpoint, "create-license-server-endpoint", "", false, "Create License Server Endpoint")
	_licensemanagerusersubscriptionsCmd.Flags().BoolVarP(&_licensemanagerusersubscriptionsDeleteLicenseServerEndpoint, "delete-license-server-endpoint", "", false, "Delete License Server Endpoint")
	_licensemanagerusersubscriptionsCmd.Flags().BoolVarP(&_licensemanagerusersubscriptionsDeregisterIdentityProvider, "deregister-identity-provider", "", false, "Deregister Identity Provider")
	_licensemanagerusersubscriptionsCmd.Flags().BoolVarP(&_licensemanagerusersubscriptionsDisassociateUser, "disassociate-user", "", false, "Disassociate User")
	_licensemanagerusersubscriptionsCmd.Flags().BoolVarP(&_licensemanagerusersubscriptionsListIdentityProviders, "list-identity-providers", "", false, "List Identity Providers")
	_licensemanagerusersubscriptionsCmd.Flags().BoolVarP(&_licensemanagerusersubscriptionsListInstances, "list-instances", "", false, "List Instances")
	_licensemanagerusersubscriptionsCmd.Flags().BoolVarP(&_licensemanagerusersubscriptionsListLicenseServerEndpoints, "list-license-server-endpoints", "", false, "List License Server Endpoints")
	_licensemanagerusersubscriptionsCmd.Flags().BoolVarP(&_licensemanagerusersubscriptionsListProductSubscriptions, "list-product-subscriptions", "", false, "List Product Subscriptions")
	_licensemanagerusersubscriptionsCmd.Flags().BoolVarP(&_licensemanagerusersubscriptionsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_licensemanagerusersubscriptionsCmd.Flags().BoolVarP(&_licensemanagerusersubscriptionsListUserAssociations, "list-user-associations", "", false, "List User Associations")
	_licensemanagerusersubscriptionsCmd.Flags().BoolVarP(&_licensemanagerusersubscriptionsRegisterIdentityProvider, "register-identity-provider", "", false, "Register Identity Provider")
	_licensemanagerusersubscriptionsCmd.Flags().BoolVarP(&_licensemanagerusersubscriptionsStartProductSubscription, "start-product-subscription", "", false, "Start Product Subscription")
	_licensemanagerusersubscriptionsCmd.Flags().BoolVarP(&_licensemanagerusersubscriptionsStopProductSubscription, "stop-product-subscription", "", false, "Stop Product Subscription")
	_licensemanagerusersubscriptionsCmd.Flags().BoolVarP(&_licensemanagerusersubscriptionsTagResource, "tag-resource", "", false, "Tag Resource")
	_licensemanagerusersubscriptionsCmd.Flags().BoolVarP(&_licensemanagerusersubscriptionsUntagResource, "untag-resource", "", false, "Untag Resource")
	_licensemanagerusersubscriptionsCmd.Flags().BoolVarP(&_licensemanagerusersubscriptionsUpdateIdentityProviderSettings, "update-identity-provider-settings", "", false, "Update Identity Provider Settings")

}
