package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appfabric"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// appfabricCmd represents the appfabric command
var _appfabricCmd = &cobra.Command{
	Use:   "appfabric",
	Short: "AWS appfabric CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := appfabric.NewFromConfig(cfg)
		if _appfabricBatchGetUserAccessTasks {
			appfabric_BatchGetUserAccessTasks(cfg, client)
			return
		}
		if _appfabricConnectAppAuthorization {
			appfabric_ConnectAppAuthorization(cfg, client)
			return
		}
		if _appfabricCreateAppAuthorization {
			appfabric_CreateAppAuthorization(cfg, client)
			return
		}
		if _appfabricCreateAppBundle {
			appfabric_CreateAppBundle(cfg, client)
			return
		}
		if _appfabricCreateIngestion {
			appfabric_CreateIngestion(cfg, client)
			return
		}
		if _appfabricCreateIngestionDestination {
			appfabric_CreateIngestionDestination(cfg, client)
			return
		}
		if _appfabricDeleteAppAuthorization {
			appfabric_DeleteAppAuthorization(cfg, client)
			return
		}
		if _appfabricDeleteAppBundle {
			appfabric_DeleteAppBundle(cfg, client)
			return
		}
		if _appfabricDeleteIngestion {
			appfabric_DeleteIngestion(cfg, client)
			return
		}
		if _appfabricDeleteIngestionDestination {
			appfabric_DeleteIngestionDestination(cfg, client)
			return
		}
		if _appfabricGetAppAuthorization {
			appfabric_GetAppAuthorization(cfg, client)
			return
		}
		if _appfabricGetAppBundle {
			appfabric_GetAppBundle(cfg, client)
			return
		}
		if _appfabricGetIngestion {
			appfabric_GetIngestion(cfg, client)
			return
		}
		if _appfabricGetIngestionDestination {
			appfabric_GetIngestionDestination(cfg, client)
			return
		}
		if _appfabricListAppAuthorizations {
			appfabric_ListAppAuthorizations(cfg, client)
			return
		}
		if _appfabricListAppBundles {
			appfabric_ListAppBundles(cfg, client)
			return
		}
		if _appfabricListIngestionDestinations {
			appfabric_ListIngestionDestinations(cfg, client)
			return
		}
		if _appfabricListIngestions {
			appfabric_ListIngestions(cfg, client)
			return
		}
		if _appfabricListTagsForResource {
			appfabric_ListTagsForResource(cfg, client)
			return
		}
		if _appfabricStartIngestion {
			appfabric_StartIngestion(cfg, client)
			return
		}
		if _appfabricStartUserAccessTasks {
			appfabric_StartUserAccessTasks(cfg, client)
			return
		}
		if _appfabricStopIngestion {
			appfabric_StopIngestion(cfg, client)
			return
		}
		if _appfabricTagResource {
			appfabric_TagResource(cfg, client)
			return
		}
		if _appfabricUntagResource {
			appfabric_UntagResource(cfg, client)
			return
		}
		if _appfabricUpdateAppAuthorization {
			appfabric_UpdateAppAuthorization(cfg, client)
			return
		}
		if _appfabricUpdateIngestionDestination {
			appfabric_UpdateIngestionDestination(cfg, client)
			return
		}

	},
}

var (
	_appfabricBatchGetUserAccessTasks    bool
	_appfabricConnectAppAuthorization    bool
	_appfabricCreateAppAuthorization     bool
	_appfabricCreateAppBundle            bool
	_appfabricCreateIngestion            bool
	_appfabricCreateIngestionDestination bool
	_appfabricDeleteAppAuthorization     bool
	_appfabricDeleteAppBundle            bool
	_appfabricDeleteIngestion            bool
	_appfabricDeleteIngestionDestination bool
	_appfabricGetAppAuthorization        bool
	_appfabricGetAppBundle               bool
	_appfabricGetIngestion               bool
	_appfabricGetIngestionDestination    bool
	_appfabricListAppAuthorizations      bool
	_appfabricListAppBundles             bool
	_appfabricListIngestionDestinations  bool
	_appfabricListIngestions             bool
	_appfabricListTagsForResource        bool
	_appfabricStartIngestion             bool
	_appfabricStartUserAccessTasks       bool
	_appfabricStopIngestion              bool
	_appfabricTagResource                bool
	_appfabricUntagResource              bool
	_appfabricUpdateAppAuthorization     bool
	_appfabricUpdateIngestionDestination bool

	_appfabricApp                            string
	_appfabricAppAuthorizationIdentifier     string
	_appfabricAppBundleIdentifier            string
	_appfabricAuthRequest                    string
	_appfabricAuthType                       string
	_appfabricClientToken                    string
	_appfabricCredential                     string
	_appfabricCustomerManagedKeyIdentifier   string
	_appfabricDestinationConfiguration       string
	_appfabricEmail                          string
	_appfabricIngestionDestinationIdentifier string
	_appfabricIngestionIdentifier            string
	_appfabricIngestionType                  string
	_appfabricMaxResults                     string
	_appfabricNextToken                      string
	_appfabricProcessingConfiguration        string
	_appfabricResourceArn                    string
	_appfabricTagKeys                        []string
	_appfabricTags                           string
	_appfabricTaskIdList                     []string
	_appfabricTenant                         string
	_appfabricTenantId                       string
)

// Gets user access details in a batch request.
// This action polls data from the tasks that are kicked off by the
// StartUserAccessTasks action.
func appfabric_BatchGetUserAccessTasks(cfg aws.Config, client *appfabric.Client) {
	input := &appfabric.BatchGetUserAccessTasksInput{
		// AppBundleIdentifier: *string, // Required
		// TaskIdList: []string, // Required
	}

	if len(_appfabricAppBundleIdentifier) > 0 {
		input.AppBundleIdentifier = aws.String(_appfabricAppBundleIdentifier)
	}
	if len(_appfabricTaskIdList) > 0 {
		input.TaskIdList = append([]string(nil), _appfabricTaskIdList...)
	}

	if resp, err := client.BatchGetUserAccessTasks(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Establishes a connection between Amazon Web Services AppFabric and an
// application, which allows AppFabric to call the APIs of the application.
func appfabric_ConnectAppAuthorization(cfg aws.Config, client *appfabric.Client) {
	input := &appfabric.ConnectAppAuthorizationInput{
		// AppAuthorizationIdentifier: *string, // Required
		// AppBundleIdentifier: *string, // Required
	}

	if len(_appfabricAppAuthorizationIdentifier) > 0 {
		input.AppAuthorizationIdentifier = aws.String(_appfabricAppAuthorizationIdentifier)
	}
	if len(_appfabricAppBundleIdentifier) > 0 {
		input.AppBundleIdentifier = aws.String(_appfabricAppBundleIdentifier)
	}
	if len(_appfabricAuthRequest) > 0 {
		if err := assignInputField(input, "AuthRequest", _appfabricAuthRequest); err != nil {
			log.Errorf("invalid --auth-request: %s", err.Error())
			return
		}
	}

	if resp, err := client.ConnectAppAuthorization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an app authorization within an app bundle, which allows AppFabric to
// connect to an application.
func appfabric_CreateAppAuthorization(cfg aws.Config, client *appfabric.Client) {
	input := &appfabric.CreateAppAuthorizationInput{
		// App: *string, // Required
		// AppBundleIdentifier: *string, // Required
		// AuthType: types.AuthType, // Required
		// Credential: types.Credential, // Required
		// Tenant: *types.Tenant, // Required
	}

	if len(_appfabricApp) > 0 {
		input.App = aws.String(_appfabricApp)
	}
	if len(_appfabricAppBundleIdentifier) > 0 {
		input.AppBundleIdentifier = aws.String(_appfabricAppBundleIdentifier)
	}
	if len(_appfabricAuthType) > 0 {
		if err := assignInputField(input, "AuthType", _appfabricAuthType); err != nil {
			log.Errorf("invalid --auth-type: %s", err.Error())
			return
		}
	}
	if len(_appfabricCredential) > 0 {
		if err := assignInputField(input, "Credential", _appfabricCredential); err != nil {
			log.Errorf("invalid --credential: %s", err.Error())
			return
		}
	}
	if len(_appfabricTenant) > 0 {
		if err := assignInputField(input, "Tenant", _appfabricTenant); err != nil {
			log.Errorf("invalid --tenant: %s", err.Error())
			return
		}
	}
	if len(_appfabricClientToken) > 0 {
		input.ClientToken = aws.String(_appfabricClientToken)
	}
	if len(_appfabricTags) > 0 {
		if err := assignInputField(input, "Tags", _appfabricTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAppAuthorization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an app bundle to collect data from an application using AppFabric.
func appfabric_CreateAppBundle(cfg aws.Config, client *appfabric.Client) {
	input := &appfabric.CreateAppBundleInput{}

	if len(_appfabricClientToken) > 0 {
		input.ClientToken = aws.String(_appfabricClientToken)
	}
	if len(_appfabricCustomerManagedKeyIdentifier) > 0 {
		input.CustomerManagedKeyIdentifier = aws.String(_appfabricCustomerManagedKeyIdentifier)
	}
	if len(_appfabricTags) > 0 {
		if err := assignInputField(input, "Tags", _appfabricTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAppBundle(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a data ingestion for an application.
func appfabric_CreateIngestion(cfg aws.Config, client *appfabric.Client) {
	input := &appfabric.CreateIngestionInput{
		// App: *string, // Required
		// AppBundleIdentifier: *string, // Required
		// IngestionType: types.IngestionType, // Required
		// TenantId: *string, // Required
	}

	if len(_appfabricApp) > 0 {
		input.App = aws.String(_appfabricApp)
	}
	if len(_appfabricAppBundleIdentifier) > 0 {
		input.AppBundleIdentifier = aws.String(_appfabricAppBundleIdentifier)
	}
	if len(_appfabricIngestionType) > 0 {
		if err := assignInputField(input, "IngestionType", _appfabricIngestionType); err != nil {
			log.Errorf("invalid --ingestion-type: %s", err.Error())
			return
		}
	}
	if len(_appfabricTenantId) > 0 {
		input.TenantId = aws.String(_appfabricTenantId)
	}
	if len(_appfabricClientToken) > 0 {
		input.ClientToken = aws.String(_appfabricClientToken)
	}
	if len(_appfabricTags) > 0 {
		if err := assignInputField(input, "Tags", _appfabricTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateIngestion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an ingestion destination, which specifies how an application's ingested
// data is processed by Amazon Web Services AppFabric and where it's delivered.
func appfabric_CreateIngestionDestination(cfg aws.Config, client *appfabric.Client) {
	input := &appfabric.CreateIngestionDestinationInput{
		// AppBundleIdentifier: *string, // Required
		// DestinationConfiguration: types.DestinationConfiguration, // Required
		// IngestionIdentifier: *string, // Required
		// ProcessingConfiguration: types.ProcessingConfiguration, // Required
	}

	if len(_appfabricAppBundleIdentifier) > 0 {
		input.AppBundleIdentifier = aws.String(_appfabricAppBundleIdentifier)
	}
	if len(_appfabricDestinationConfiguration) > 0 {
		if err := assignInputField(input, "DestinationConfiguration", _appfabricDestinationConfiguration); err != nil {
			log.Errorf("invalid --destination-configuration: %s", err.Error())
			return
		}
	}
	if len(_appfabricIngestionIdentifier) > 0 {
		input.IngestionIdentifier = aws.String(_appfabricIngestionIdentifier)
	}
	if len(_appfabricProcessingConfiguration) > 0 {
		if err := assignInputField(input, "ProcessingConfiguration", _appfabricProcessingConfiguration); err != nil {
			log.Errorf("invalid --processing-configuration: %s", err.Error())
			return
		}
	}
	if len(_appfabricClientToken) > 0 {
		input.ClientToken = aws.String(_appfabricClientToken)
	}
	if len(_appfabricTags) > 0 {
		if err := assignInputField(input, "Tags", _appfabricTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateIngestionDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an app authorization. You must delete the associated ingestion before
// you can delete an app authorization.
func appfabric_DeleteAppAuthorization(cfg aws.Config, client *appfabric.Client) {
	input := &appfabric.DeleteAppAuthorizationInput{
		// AppAuthorizationIdentifier: *string, // Required
		// AppBundleIdentifier: *string, // Required
	}

	if len(_appfabricAppAuthorizationIdentifier) > 0 {
		input.AppAuthorizationIdentifier = aws.String(_appfabricAppAuthorizationIdentifier)
	}
	if len(_appfabricAppBundleIdentifier) > 0 {
		input.AppBundleIdentifier = aws.String(_appfabricAppBundleIdentifier)
	}

	if resp, err := client.DeleteAppAuthorization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an app bundle. You must delete all associated app authorizations before
// you can delete an app bundle.
func appfabric_DeleteAppBundle(cfg aws.Config, client *appfabric.Client) {
	input := &appfabric.DeleteAppBundleInput{
		// AppBundleIdentifier: *string, // Required
	}

	if len(_appfabricAppBundleIdentifier) > 0 {
		input.AppBundleIdentifier = aws.String(_appfabricAppBundleIdentifier)
	}

	if resp, err := client.DeleteAppBundle(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an ingestion. You must stop (disable) the ingestion and you must delete
// all associated ingestion destinations before you can delete an app ingestion.
func appfabric_DeleteIngestion(cfg aws.Config, client *appfabric.Client) {
	input := &appfabric.DeleteIngestionInput{
		// AppBundleIdentifier: *string, // Required
		// IngestionIdentifier: *string, // Required
	}

	if len(_appfabricAppBundleIdentifier) > 0 {
		input.AppBundleIdentifier = aws.String(_appfabricAppBundleIdentifier)
	}
	if len(_appfabricIngestionIdentifier) > 0 {
		input.IngestionIdentifier = aws.String(_appfabricIngestionIdentifier)
	}

	if resp, err := client.DeleteIngestion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an ingestion destination.
// This deletes the association between an ingestion and it's destination. It
// doesn't delete previously ingested data or the storage destination, such as the
// Amazon S3 bucket where the data is delivered. If the ingestion destination is
// deleted while the associated ingestion is enabled, the ingestion will fail and
// is eventually disabled.
func appfabric_DeleteIngestionDestination(cfg aws.Config, client *appfabric.Client) {
	input := &appfabric.DeleteIngestionDestinationInput{
		// AppBundleIdentifier: *string, // Required
		// IngestionDestinationIdentifier: *string, // Required
		// IngestionIdentifier: *string, // Required
	}

	if len(_appfabricAppBundleIdentifier) > 0 {
		input.AppBundleIdentifier = aws.String(_appfabricAppBundleIdentifier)
	}
	if len(_appfabricIngestionDestinationIdentifier) > 0 {
		input.IngestionDestinationIdentifier = aws.String(_appfabricIngestionDestinationIdentifier)
	}
	if len(_appfabricIngestionIdentifier) > 0 {
		input.IngestionIdentifier = aws.String(_appfabricIngestionIdentifier)
	}

	if resp, err := client.DeleteIngestionDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an app authorization.
func appfabric_GetAppAuthorization(cfg aws.Config, client *appfabric.Client) {
	input := &appfabric.GetAppAuthorizationInput{
		// AppAuthorizationIdentifier: *string, // Required
		// AppBundleIdentifier: *string, // Required
	}

	if len(_appfabricAppAuthorizationIdentifier) > 0 {
		input.AppAuthorizationIdentifier = aws.String(_appfabricAppAuthorizationIdentifier)
	}
	if len(_appfabricAppBundleIdentifier) > 0 {
		input.AppBundleIdentifier = aws.String(_appfabricAppBundleIdentifier)
	}

	if resp, err := client.GetAppAuthorization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an app bundle.
func appfabric_GetAppBundle(cfg aws.Config, client *appfabric.Client) {
	input := &appfabric.GetAppBundleInput{
		// AppBundleIdentifier: *string, // Required
	}

	if len(_appfabricAppBundleIdentifier) > 0 {
		input.AppBundleIdentifier = aws.String(_appfabricAppBundleIdentifier)
	}

	if resp, err := client.GetAppBundle(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an ingestion.
func appfabric_GetIngestion(cfg aws.Config, client *appfabric.Client) {
	input := &appfabric.GetIngestionInput{
		// AppBundleIdentifier: *string, // Required
		// IngestionIdentifier: *string, // Required
	}

	if len(_appfabricAppBundleIdentifier) > 0 {
		input.AppBundleIdentifier = aws.String(_appfabricAppBundleIdentifier)
	}
	if len(_appfabricIngestionIdentifier) > 0 {
		input.IngestionIdentifier = aws.String(_appfabricIngestionIdentifier)
	}

	if resp, err := client.GetIngestion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an ingestion destination.
func appfabric_GetIngestionDestination(cfg aws.Config, client *appfabric.Client) {
	input := &appfabric.GetIngestionDestinationInput{
		// AppBundleIdentifier: *string, // Required
		// IngestionDestinationIdentifier: *string, // Required
		// IngestionIdentifier: *string, // Required
	}

	if len(_appfabricAppBundleIdentifier) > 0 {
		input.AppBundleIdentifier = aws.String(_appfabricAppBundleIdentifier)
	}
	if len(_appfabricIngestionDestinationIdentifier) > 0 {
		input.IngestionDestinationIdentifier = aws.String(_appfabricIngestionDestinationIdentifier)
	}
	if len(_appfabricIngestionIdentifier) > 0 {
		input.IngestionIdentifier = aws.String(_appfabricIngestionIdentifier)
	}

	if resp, err := client.GetIngestionDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of all app authorizations configured for an app bundle.
func appfabric_ListAppAuthorizations(cfg aws.Config, client *appfabric.Client) {
	input := &appfabric.ListAppAuthorizationsInput{
		// AppBundleIdentifier: *string, // Required
	}

	if len(_appfabricAppBundleIdentifier) > 0 {
		input.AppBundleIdentifier = aws.String(_appfabricAppBundleIdentifier)
	}
	if len(_appfabricMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appfabricMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appfabricNextToken) > 0 {
		input.NextToken = aws.String(_appfabricNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAppAuthorizations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appfabric.ListAppAuthorizationsOutput
	p := appfabric.NewListAppAuthorizationsPaginator(client, input)
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

// Returns a list of app bundles.
func appfabric_ListAppBundles(cfg aws.Config, client *appfabric.Client) {
	input := &appfabric.ListAppBundlesInput{}

	if len(_appfabricMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appfabricMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appfabricNextToken) > 0 {
		input.NextToken = aws.String(_appfabricNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAppBundles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appfabric.ListAppBundlesOutput
	p := appfabric.NewListAppBundlesPaginator(client, input)
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

// Returns a list of all ingestion destinations configured for an ingestion.
func appfabric_ListIngestionDestinations(cfg aws.Config, client *appfabric.Client) {
	input := &appfabric.ListIngestionDestinationsInput{
		// AppBundleIdentifier: *string, // Required
		// IngestionIdentifier: *string, // Required
	}

	if len(_appfabricAppBundleIdentifier) > 0 {
		input.AppBundleIdentifier = aws.String(_appfabricAppBundleIdentifier)
	}
	if len(_appfabricIngestionIdentifier) > 0 {
		input.IngestionIdentifier = aws.String(_appfabricIngestionIdentifier)
	}
	if len(_appfabricMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appfabricMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appfabricNextToken) > 0 {
		input.NextToken = aws.String(_appfabricNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIngestionDestinations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appfabric.ListIngestionDestinationsOutput
	p := appfabric.NewListIngestionDestinationsPaginator(client, input)
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

// Returns a list of all ingestions configured for an app bundle.
func appfabric_ListIngestions(cfg aws.Config, client *appfabric.Client) {
	input := &appfabric.ListIngestionsInput{
		// AppBundleIdentifier: *string, // Required
	}

	if len(_appfabricAppBundleIdentifier) > 0 {
		input.AppBundleIdentifier = aws.String(_appfabricAppBundleIdentifier)
	}
	if len(_appfabricMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appfabricMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appfabricNextToken) > 0 {
		input.NextToken = aws.String(_appfabricNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIngestions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appfabric.ListIngestionsOutput
	p := appfabric.NewListIngestionsPaginator(client, input)
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

// Returns a list of tags for a resource.
func appfabric_ListTagsForResource(cfg aws.Config, client *appfabric.Client) {
	input := &appfabric.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_appfabricResourceArn) > 0 {
		input.ResourceArn = aws.String(_appfabricResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts (enables) an ingestion, which collects data from an application.
func appfabric_StartIngestion(cfg aws.Config, client *appfabric.Client) {
	input := &appfabric.StartIngestionInput{
		// AppBundleIdentifier: *string, // Required
		// IngestionIdentifier: *string, // Required
	}

	if len(_appfabricAppBundleIdentifier) > 0 {
		input.AppBundleIdentifier = aws.String(_appfabricAppBundleIdentifier)
	}
	if len(_appfabricIngestionIdentifier) > 0 {
		input.IngestionIdentifier = aws.String(_appfabricIngestionIdentifier)
	}

	if resp, err := client.StartIngestion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the tasks to search user access status for a specific email address.
// The tasks are stopped when the user access status data is found. The tasks are
// terminated when the API calls to the application time out.
func appfabric_StartUserAccessTasks(cfg aws.Config, client *appfabric.Client) {
	input := &appfabric.StartUserAccessTasksInput{
		// AppBundleIdentifier: *string, // Required
		// Email: *string, // Required
	}

	if len(_appfabricAppBundleIdentifier) > 0 {
		input.AppBundleIdentifier = aws.String(_appfabricAppBundleIdentifier)
	}
	if len(_appfabricEmail) > 0 {
		input.Email = aws.String(_appfabricEmail)
	}

	if resp, err := client.StartUserAccessTasks(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops (disables) an ingestion.
func appfabric_StopIngestion(cfg aws.Config, client *appfabric.Client) {
	input := &appfabric.StopIngestionInput{
		// AppBundleIdentifier: *string, // Required
		// IngestionIdentifier: *string, // Required
	}

	if len(_appfabricAppBundleIdentifier) > 0 {
		input.AppBundleIdentifier = aws.String(_appfabricAppBundleIdentifier)
	}
	if len(_appfabricIngestionIdentifier) > 0 {
		input.IngestionIdentifier = aws.String(_appfabricIngestionIdentifier)
	}

	if resp, err := client.StopIngestion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns one or more tags (key-value pairs) to the specified resource.
func appfabric_TagResource(cfg aws.Config, client *appfabric.Client) {
	input := &appfabric.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_appfabricResourceArn) > 0 {
		input.ResourceArn = aws.String(_appfabricResourceArn)
	}
	if len(_appfabricTags) > 0 {
		if err := assignInputField(input, "Tags", _appfabricTags); err != nil {
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

// Removes a tag or tags from a resource.
func appfabric_UntagResource(cfg aws.Config, client *appfabric.Client) {
	input := &appfabric.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_appfabricResourceArn) > 0 {
		input.ResourceArn = aws.String(_appfabricResourceArn)
	}
	if len(_appfabricTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _appfabricTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an app authorization within an app bundle, which allows AppFabric to
// connect to an application.
//
// If the app authorization was in a connected state, updating the app
// authorization will set it back to a PendingConnect state.
func appfabric_UpdateAppAuthorization(cfg aws.Config, client *appfabric.Client) {
	input := &appfabric.UpdateAppAuthorizationInput{
		// AppAuthorizationIdentifier: *string, // Required
		// AppBundleIdentifier: *string, // Required
	}

	if len(_appfabricAppAuthorizationIdentifier) > 0 {
		input.AppAuthorizationIdentifier = aws.String(_appfabricAppAuthorizationIdentifier)
	}
	if len(_appfabricAppBundleIdentifier) > 0 {
		input.AppBundleIdentifier = aws.String(_appfabricAppBundleIdentifier)
	}
	if len(_appfabricCredential) > 0 {
		if err := assignInputField(input, "Credential", _appfabricCredential); err != nil {
			log.Errorf("invalid --credential: %s", err.Error())
			return
		}
	}
	if len(_appfabricTenant) > 0 {
		if err := assignInputField(input, "Tenant", _appfabricTenant); err != nil {
			log.Errorf("invalid --tenant: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAppAuthorization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an ingestion destination, which specifies how an application's ingested
// data is processed by Amazon Web Services AppFabric and where it's delivered.
func appfabric_UpdateIngestionDestination(cfg aws.Config, client *appfabric.Client) {
	input := &appfabric.UpdateIngestionDestinationInput{
		// AppBundleIdentifier: *string, // Required
		// DestinationConfiguration: types.DestinationConfiguration, // Required
		// IngestionDestinationIdentifier: *string, // Required
		// IngestionIdentifier: *string, // Required
	}

	if len(_appfabricAppBundleIdentifier) > 0 {
		input.AppBundleIdentifier = aws.String(_appfabricAppBundleIdentifier)
	}
	if len(_appfabricDestinationConfiguration) > 0 {
		if err := assignInputField(input, "DestinationConfiguration", _appfabricDestinationConfiguration); err != nil {
			log.Errorf("invalid --destination-configuration: %s", err.Error())
			return
		}
	}
	if len(_appfabricIngestionDestinationIdentifier) > 0 {
		input.IngestionDestinationIdentifier = aws.String(_appfabricIngestionDestinationIdentifier)
	}
	if len(_appfabricIngestionIdentifier) > 0 {
		input.IngestionIdentifier = aws.String(_appfabricIngestionIdentifier)
	}

	if resp, err := client.UpdateIngestionDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_appfabricCmd)
	_appfabricCmd.Flags().SortFlags = false

	_appfabricCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_appfabricCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_appfabricCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_appfabricCmd.Flags().StringVarP(&_appfabricApp, "app", "", "", "App")
	_appfabricCmd.Flags().StringVarP(&_appfabricAppAuthorizationIdentifier, "app-authorization-identifier", "", "", "App Authorization Identifier")
	_appfabricCmd.Flags().StringVarP(&_appfabricAppBundleIdentifier, "app-bundle-identifier", "", "", "App Bundle Identifier")
	_appfabricCmd.Flags().StringVarP(&_appfabricAuthRequest, "auth-request", "", "", "Auth Request")
	_appfabricCmd.Flags().StringVarP(&_appfabricAuthType, "auth-type", "", "", "Auth Type")
	_appfabricCmd.Flags().StringVarP(&_appfabricClientToken, "client-token", "", "", "Client Token")
	_appfabricCmd.Flags().StringVarP(&_appfabricCredential, "credential", "", "", "Credential")
	_appfabricCmd.Flags().StringVarP(&_appfabricCustomerManagedKeyIdentifier, "customer-managed-key-identifier", "", "", "Customer Managed Key Identifier")
	_appfabricCmd.Flags().StringVarP(&_appfabricDestinationConfiguration, "destination-configuration", "", "", "Destination Configuration")
	_appfabricCmd.Flags().StringVarP(&_appfabricEmail, "email", "", "", "Email")
	_appfabricCmd.Flags().StringVarP(&_appfabricIngestionDestinationIdentifier, "ingestion-destination-identifier", "", "", "Ingestion Destination Identifier")
	_appfabricCmd.Flags().StringVarP(&_appfabricIngestionIdentifier, "ingestion-identifier", "", "", "Ingestion Identifier")
	_appfabricCmd.Flags().StringVarP(&_appfabricIngestionType, "ingestion-type", "", "", "Ingestion Type")
	_appfabricCmd.Flags().StringVarP(&_appfabricMaxResults, "max-results", "", "", "Max Results")
	_appfabricCmd.Flags().StringVarP(&_appfabricNextToken, "next-token", "", "", "Next Token")
	_appfabricCmd.Flags().StringVarP(&_appfabricProcessingConfiguration, "processing-configuration", "", "", "Processing Configuration")
	_appfabricCmd.Flags().StringVarP(&_appfabricResourceArn, "resource-arn", "", "", "Resource ARN")
	_appfabricCmd.Flags().StringSliceVarP(&_appfabricTagKeys, "tag-keys", "", nil, "Tag Keys")
	_appfabricCmd.Flags().StringVarP(&_appfabricTags, "tags", "", "", "Tags")
	_appfabricCmd.Flags().StringSliceVarP(&_appfabricTaskIdList, "task-id-list", "", nil, "Task ID List")
	_appfabricCmd.Flags().StringVarP(&_appfabricTenant, "tenant", "", "", "Tenant")
	_appfabricCmd.Flags().StringVarP(&_appfabricTenantId, "tenant-id", "", "", "Tenant ID")

	_appfabricCmd.Flags().BoolVarP(&_appfabricBatchGetUserAccessTasks, "batch-get-user-access-tasks", "", false, "Batch Get User Access Tasks")
	_appfabricCmd.Flags().BoolVarP(&_appfabricConnectAppAuthorization, "connect-app-authorization", "", false, "Connect App Authorization")
	_appfabricCmd.Flags().BoolVarP(&_appfabricCreateAppAuthorization, "create-app-authorization", "", false, "Create App Authorization")
	_appfabricCmd.Flags().BoolVarP(&_appfabricCreateAppBundle, "create-app-bundle", "", false, "Create App Bundle")
	_appfabricCmd.Flags().BoolVarP(&_appfabricCreateIngestion, "create-ingestion", "", false, "Create Ingestion")
	_appfabricCmd.Flags().BoolVarP(&_appfabricCreateIngestionDestination, "create-ingestion-destination", "", false, "Create Ingestion Destination")
	_appfabricCmd.Flags().BoolVarP(&_appfabricDeleteAppAuthorization, "delete-app-authorization", "", false, "Delete App Authorization")
	_appfabricCmd.Flags().BoolVarP(&_appfabricDeleteAppBundle, "delete-app-bundle", "", false, "Delete App Bundle")
	_appfabricCmd.Flags().BoolVarP(&_appfabricDeleteIngestion, "delete-ingestion", "", false, "Delete Ingestion")
	_appfabricCmd.Flags().BoolVarP(&_appfabricDeleteIngestionDestination, "delete-ingestion-destination", "", false, "Delete Ingestion Destination")
	_appfabricCmd.Flags().BoolVarP(&_appfabricGetAppAuthorization, "get-app-authorization", "", false, "Get App Authorization")
	_appfabricCmd.Flags().BoolVarP(&_appfabricGetAppBundle, "get-app-bundle", "", false, "Get App Bundle")
	_appfabricCmd.Flags().BoolVarP(&_appfabricGetIngestion, "get-ingestion", "", false, "Get Ingestion")
	_appfabricCmd.Flags().BoolVarP(&_appfabricGetIngestionDestination, "get-ingestion-destination", "", false, "Get Ingestion Destination")
	_appfabricCmd.Flags().BoolVarP(&_appfabricListAppAuthorizations, "list-app-authorizations", "", false, "List App Authorizations")
	_appfabricCmd.Flags().BoolVarP(&_appfabricListAppBundles, "list-app-bundles", "", false, "List App Bundles")
	_appfabricCmd.Flags().BoolVarP(&_appfabricListIngestionDestinations, "list-ingestion-destinations", "", false, "List Ingestion Destinations")
	_appfabricCmd.Flags().BoolVarP(&_appfabricListIngestions, "list-ingestions", "", false, "List Ingestions")
	_appfabricCmd.Flags().BoolVarP(&_appfabricListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_appfabricCmd.Flags().BoolVarP(&_appfabricStartIngestion, "start-ingestion", "", false, "Start Ingestion")
	_appfabricCmd.Flags().BoolVarP(&_appfabricStartUserAccessTasks, "start-user-access-tasks", "", false, "Start User Access Tasks")
	_appfabricCmd.Flags().BoolVarP(&_appfabricStopIngestion, "stop-ingestion", "", false, "Stop Ingestion")
	_appfabricCmd.Flags().BoolVarP(&_appfabricTagResource, "tag-resource", "", false, "Tag Resource")
	_appfabricCmd.Flags().BoolVarP(&_appfabricUntagResource, "untag-resource", "", false, "Untag Resource")
	_appfabricCmd.Flags().BoolVarP(&_appfabricUpdateAppAuthorization, "update-app-authorization", "", false, "Update App Authorization")
	_appfabricCmd.Flags().BoolVarP(&_appfabricUpdateIngestionDestination, "update-ingestion-destination", "", false, "Update Ingestion Destination")

}
