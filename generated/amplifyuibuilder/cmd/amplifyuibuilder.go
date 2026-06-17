package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/amplifyuibuilder"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// amplifyuibuilderCmd represents the amplifyuibuilder command
var _amplifyuibuilderCmd = &cobra.Command{
	Use:   "amplifyuibuilder",
	Short: "AWS amplifyuibuilder CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := amplifyuibuilder.NewFromConfig(cfg)
		if _amplifyuibuilderCreateComponent {
			amplifyuibuilder_CreateComponent(cfg, client)
			return
		}
		if _amplifyuibuilderCreateForm {
			amplifyuibuilder_CreateForm(cfg, client)
			return
		}
		if _amplifyuibuilderCreateTheme {
			amplifyuibuilder_CreateTheme(cfg, client)
			return
		}
		if _amplifyuibuilderDeleteComponent {
			amplifyuibuilder_DeleteComponent(cfg, client)
			return
		}
		if _amplifyuibuilderDeleteForm {
			amplifyuibuilder_DeleteForm(cfg, client)
			return
		}
		if _amplifyuibuilderDeleteTheme {
			amplifyuibuilder_DeleteTheme(cfg, client)
			return
		}
		if _amplifyuibuilderExchangeCodeForToken {
			amplifyuibuilder_ExchangeCodeForToken(cfg, client)
			return
		}
		if _amplifyuibuilderExportComponents {
			amplifyuibuilder_ExportComponents(cfg, client)
			return
		}
		if _amplifyuibuilderExportForms {
			amplifyuibuilder_ExportForms(cfg, client)
			return
		}
		if _amplifyuibuilderExportThemes {
			amplifyuibuilder_ExportThemes(cfg, client)
			return
		}
		if _amplifyuibuilderGetCodegenJob {
			amplifyuibuilder_GetCodegenJob(cfg, client)
			return
		}
		if _amplifyuibuilderGetComponent {
			amplifyuibuilder_GetComponent(cfg, client)
			return
		}
		if _amplifyuibuilderGetForm {
			amplifyuibuilder_GetForm(cfg, client)
			return
		}
		if _amplifyuibuilderGetMetadata {
			amplifyuibuilder_GetMetadata(cfg, client)
			return
		}
		if _amplifyuibuilderGetTheme {
			amplifyuibuilder_GetTheme(cfg, client)
			return
		}
		if _amplifyuibuilderListCodegenJobs {
			amplifyuibuilder_ListCodegenJobs(cfg, client)
			return
		}
		if _amplifyuibuilderListComponents {
			amplifyuibuilder_ListComponents(cfg, client)
			return
		}
		if _amplifyuibuilderListForms {
			amplifyuibuilder_ListForms(cfg, client)
			return
		}
		if _amplifyuibuilderListTagsForResource {
			amplifyuibuilder_ListTagsForResource(cfg, client)
			return
		}
		if _amplifyuibuilderListThemes {
			amplifyuibuilder_ListThemes(cfg, client)
			return
		}
		if _amplifyuibuilderPutMetadataFlag {
			amplifyuibuilder_PutMetadataFlag(cfg, client)
			return
		}
		if _amplifyuibuilderRefreshToken {
			amplifyuibuilder_RefreshToken(cfg, client)
			return
		}
		if _amplifyuibuilderStartCodegenJob {
			amplifyuibuilder_StartCodegenJob(cfg, client)
			return
		}
		if _amplifyuibuilderTagResource {
			amplifyuibuilder_TagResource(cfg, client)
			return
		}
		if _amplifyuibuilderUntagResource {
			amplifyuibuilder_UntagResource(cfg, client)
			return
		}
		if _amplifyuibuilderUpdateComponent {
			amplifyuibuilder_UpdateComponent(cfg, client)
			return
		}
		if _amplifyuibuilderUpdateForm {
			amplifyuibuilder_UpdateForm(cfg, client)
			return
		}
		if _amplifyuibuilderUpdateTheme {
			amplifyuibuilder_UpdateTheme(cfg, client)
			return
		}

	},
}

var (
	_amplifyuibuilderCreateComponent      bool
	_amplifyuibuilderCreateForm           bool
	_amplifyuibuilderCreateTheme          bool
	_amplifyuibuilderDeleteComponent      bool
	_amplifyuibuilderDeleteForm           bool
	_amplifyuibuilderDeleteTheme          bool
	_amplifyuibuilderExchangeCodeForToken bool
	_amplifyuibuilderExportComponents     bool
	_amplifyuibuilderExportForms          bool
	_amplifyuibuilderExportThemes         bool
	_amplifyuibuilderGetCodegenJob        bool
	_amplifyuibuilderGetComponent         bool
	_amplifyuibuilderGetForm              bool
	_amplifyuibuilderGetMetadata          bool
	_amplifyuibuilderGetTheme             bool
	_amplifyuibuilderListCodegenJobs      bool
	_amplifyuibuilderListComponents       bool
	_amplifyuibuilderListForms            bool
	_amplifyuibuilderListTagsForResource  bool
	_amplifyuibuilderListThemes           bool
	_amplifyuibuilderPutMetadataFlag      bool
	_amplifyuibuilderRefreshToken         bool
	_amplifyuibuilderStartCodegenJob      bool
	_amplifyuibuilderTagResource          bool
	_amplifyuibuilderUntagResource        bool
	_amplifyuibuilderUpdateComponent      bool
	_amplifyuibuilderUpdateForm           bool
	_amplifyuibuilderUpdateTheme          bool

	_amplifyuibuilderAppId              string
	_amplifyuibuilderBody               string
	_amplifyuibuilderClientToken        string
	_amplifyuibuilderCodegenJobToCreate string
	_amplifyuibuilderComponentToCreate  string
	_amplifyuibuilderEnvironmentName    string
	_amplifyuibuilderFeatureName        string
	_amplifyuibuilderFormToCreate       string
	_amplifyuibuilderId                 string
	_amplifyuibuilderMaxResults         string
	_amplifyuibuilderNextToken          string
	_amplifyuibuilderProvider           string
	_amplifyuibuilderRefreshTokenBody   string
	_amplifyuibuilderRequest            string
	_amplifyuibuilderResourceArn        string
	_amplifyuibuilderTagKeys            []string
	_amplifyuibuilderTags               string
	_amplifyuibuilderThemeToCreate      string
	_amplifyuibuilderUpdatedComponent   string
	_amplifyuibuilderUpdatedForm        string
	_amplifyuibuilderUpdatedTheme       string
)

// Creates a new component for an Amplify app.
func amplifyuibuilder_CreateComponent(cfg aws.Config, client *amplifyuibuilder.Client) {
	input := &amplifyuibuilder.CreateComponentInput{
		// AppId: *string, // Required
		// ComponentToCreate: *types.CreateComponentData, // Required
		// EnvironmentName: *string, // Required
	}

	if len(_amplifyuibuilderAppId) > 0 {
		input.AppId = aws.String(_amplifyuibuilderAppId)
	}
	if len(_amplifyuibuilderComponentToCreate) > 0 {
		if err := assignInputField(input, "ComponentToCreate", _amplifyuibuilderComponentToCreate); err != nil {
			log.Errorf("invalid --component-to-create: %s", err.Error())
			return
		}
	}
	if len(_amplifyuibuilderEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_amplifyuibuilderEnvironmentName)
	}
	if len(_amplifyuibuilderClientToken) > 0 {
		input.ClientToken = aws.String(_amplifyuibuilderClientToken)
	}

	if resp, err := client.CreateComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new form for an Amplify app.
func amplifyuibuilder_CreateForm(cfg aws.Config, client *amplifyuibuilder.Client) {
	input := &amplifyuibuilder.CreateFormInput{
		// AppId: *string, // Required
		// EnvironmentName: *string, // Required
		// FormToCreate: *types.CreateFormData, // Required
	}

	if len(_amplifyuibuilderAppId) > 0 {
		input.AppId = aws.String(_amplifyuibuilderAppId)
	}
	if len(_amplifyuibuilderEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_amplifyuibuilderEnvironmentName)
	}
	if len(_amplifyuibuilderFormToCreate) > 0 {
		if err := assignInputField(input, "FormToCreate", _amplifyuibuilderFormToCreate); err != nil {
			log.Errorf("invalid --form-to-create: %s", err.Error())
			return
		}
	}
	if len(_amplifyuibuilderClientToken) > 0 {
		input.ClientToken = aws.String(_amplifyuibuilderClientToken)
	}

	if resp, err := client.CreateForm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a theme to apply to the components in an Amplify app.
func amplifyuibuilder_CreateTheme(cfg aws.Config, client *amplifyuibuilder.Client) {
	input := &amplifyuibuilder.CreateThemeInput{
		// AppId: *string, // Required
		// EnvironmentName: *string, // Required
		// ThemeToCreate: *types.CreateThemeData, // Required
	}

	if len(_amplifyuibuilderAppId) > 0 {
		input.AppId = aws.String(_amplifyuibuilderAppId)
	}
	if len(_amplifyuibuilderEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_amplifyuibuilderEnvironmentName)
	}
	if len(_amplifyuibuilderThemeToCreate) > 0 {
		if err := assignInputField(input, "ThemeToCreate", _amplifyuibuilderThemeToCreate); err != nil {
			log.Errorf("invalid --theme-to-create: %s", err.Error())
			return
		}
	}
	if len(_amplifyuibuilderClientToken) > 0 {
		input.ClientToken = aws.String(_amplifyuibuilderClientToken)
	}

	if resp, err := client.CreateTheme(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a component from an Amplify app.
func amplifyuibuilder_DeleteComponent(cfg aws.Config, client *amplifyuibuilder.Client) {
	input := &amplifyuibuilder.DeleteComponentInput{
		// AppId: *string, // Required
		// EnvironmentName: *string, // Required
		// Id: *string, // Required
	}

	if len(_amplifyuibuilderAppId) > 0 {
		input.AppId = aws.String(_amplifyuibuilderAppId)
	}
	if len(_amplifyuibuilderEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_amplifyuibuilderEnvironmentName)
	}
	if len(_amplifyuibuilderId) > 0 {
		input.Id = aws.String(_amplifyuibuilderId)
	}

	if resp, err := client.DeleteComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a form from an Amplify app.
func amplifyuibuilder_DeleteForm(cfg aws.Config, client *amplifyuibuilder.Client) {
	input := &amplifyuibuilder.DeleteFormInput{
		// AppId: *string, // Required
		// EnvironmentName: *string, // Required
		// Id: *string, // Required
	}

	if len(_amplifyuibuilderAppId) > 0 {
		input.AppId = aws.String(_amplifyuibuilderAppId)
	}
	if len(_amplifyuibuilderEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_amplifyuibuilderEnvironmentName)
	}
	if len(_amplifyuibuilderId) > 0 {
		input.Id = aws.String(_amplifyuibuilderId)
	}

	if resp, err := client.DeleteForm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a theme from an Amplify app.
func amplifyuibuilder_DeleteTheme(cfg aws.Config, client *amplifyuibuilder.Client) {
	input := &amplifyuibuilder.DeleteThemeInput{
		// AppId: *string, // Required
		// EnvironmentName: *string, // Required
		// Id: *string, // Required
	}

	if len(_amplifyuibuilderAppId) > 0 {
		input.AppId = aws.String(_amplifyuibuilderAppId)
	}
	if len(_amplifyuibuilderEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_amplifyuibuilderEnvironmentName)
	}
	if len(_amplifyuibuilderId) > 0 {
		input.Id = aws.String(_amplifyuibuilderId)
	}

	if resp, err := client.DeleteTheme(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is for internal use.
// Amplify uses this action to exchange an access code for a token.
func amplifyuibuilder_ExchangeCodeForToken(cfg aws.Config, client *amplifyuibuilder.Client) {
	input := &amplifyuibuilder.ExchangeCodeForTokenInput{
		// Provider: types.TokenProviders, // Required
		// Request: *types.ExchangeCodeForTokenRequestBody, // Required
	}

	if len(_amplifyuibuilderProvider) > 0 {
		if err := assignInputField(input, "Provider", _amplifyuibuilderProvider); err != nil {
			log.Errorf("invalid --provider: %s", err.Error())
			return
		}
	}
	if len(_amplifyuibuilderRequest) > 0 {
		if err := assignInputField(input, "Request", _amplifyuibuilderRequest); err != nil {
			log.Errorf("invalid --request: %s", err.Error())
			return
		}
	}

	if resp, err := client.ExchangeCodeForToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Exports component configurations to code that is ready to integrate into an
// Amplify app.
func amplifyuibuilder_ExportComponents(cfg aws.Config, client *amplifyuibuilder.Client) {
	input := &amplifyuibuilder.ExportComponentsInput{
		// AppId: *string, // Required
		// EnvironmentName: *string, // Required
	}

	if len(_amplifyuibuilderAppId) > 0 {
		input.AppId = aws.String(_amplifyuibuilderAppId)
	}
	if len(_amplifyuibuilderEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_amplifyuibuilderEnvironmentName)
	}
	if len(_amplifyuibuilderNextToken) > 0 {
		input.NextToken = aws.String(_amplifyuibuilderNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ExportComponents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*amplifyuibuilder.ExportComponentsOutput
	p := amplifyuibuilder.NewExportComponentsPaginator(client, input)
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

// Exports form configurations to code that is ready to integrate into an Amplify
// app.
func amplifyuibuilder_ExportForms(cfg aws.Config, client *amplifyuibuilder.Client) {
	input := &amplifyuibuilder.ExportFormsInput{
		// AppId: *string, // Required
		// EnvironmentName: *string, // Required
	}

	if len(_amplifyuibuilderAppId) > 0 {
		input.AppId = aws.String(_amplifyuibuilderAppId)
	}
	if len(_amplifyuibuilderEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_amplifyuibuilderEnvironmentName)
	}
	if len(_amplifyuibuilderNextToken) > 0 {
		input.NextToken = aws.String(_amplifyuibuilderNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ExportForms(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*amplifyuibuilder.ExportFormsOutput
	p := amplifyuibuilder.NewExportFormsPaginator(client, input)
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

// Exports theme configurations to code that is ready to integrate into an Amplify
// app.
func amplifyuibuilder_ExportThemes(cfg aws.Config, client *amplifyuibuilder.Client) {
	input := &amplifyuibuilder.ExportThemesInput{
		// AppId: *string, // Required
		// EnvironmentName: *string, // Required
	}

	if len(_amplifyuibuilderAppId) > 0 {
		input.AppId = aws.String(_amplifyuibuilderAppId)
	}
	if len(_amplifyuibuilderEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_amplifyuibuilderEnvironmentName)
	}
	if len(_amplifyuibuilderNextToken) > 0 {
		input.NextToken = aws.String(_amplifyuibuilderNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ExportThemes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*amplifyuibuilder.ExportThemesOutput
	p := amplifyuibuilder.NewExportThemesPaginator(client, input)
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

// Returns an existing code generation job.
func amplifyuibuilder_GetCodegenJob(cfg aws.Config, client *amplifyuibuilder.Client) {
	input := &amplifyuibuilder.GetCodegenJobInput{
		// AppId: *string, // Required
		// EnvironmentName: *string, // Required
		// Id: *string, // Required
	}

	if len(_amplifyuibuilderAppId) > 0 {
		input.AppId = aws.String(_amplifyuibuilderAppId)
	}
	if len(_amplifyuibuilderEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_amplifyuibuilderEnvironmentName)
	}
	if len(_amplifyuibuilderId) > 0 {
		input.Id = aws.String(_amplifyuibuilderId)
	}

	if resp, err := client.GetCodegenJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an existing component for an Amplify app.
func amplifyuibuilder_GetComponent(cfg aws.Config, client *amplifyuibuilder.Client) {
	input := &amplifyuibuilder.GetComponentInput{
		// AppId: *string, // Required
		// EnvironmentName: *string, // Required
		// Id: *string, // Required
	}

	if len(_amplifyuibuilderAppId) > 0 {
		input.AppId = aws.String(_amplifyuibuilderAppId)
	}
	if len(_amplifyuibuilderEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_amplifyuibuilderEnvironmentName)
	}
	if len(_amplifyuibuilderId) > 0 {
		input.Id = aws.String(_amplifyuibuilderId)
	}

	if resp, err := client.GetComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an existing form for an Amplify app.
func amplifyuibuilder_GetForm(cfg aws.Config, client *amplifyuibuilder.Client) {
	input := &amplifyuibuilder.GetFormInput{
		// AppId: *string, // Required
		// EnvironmentName: *string, // Required
		// Id: *string, // Required
	}

	if len(_amplifyuibuilderAppId) > 0 {
		input.AppId = aws.String(_amplifyuibuilderAppId)
	}
	if len(_amplifyuibuilderEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_amplifyuibuilderEnvironmentName)
	}
	if len(_amplifyuibuilderId) > 0 {
		input.Id = aws.String(_amplifyuibuilderId)
	}

	if resp, err := client.GetForm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns existing metadata for an Amplify app.
func amplifyuibuilder_GetMetadata(cfg aws.Config, client *amplifyuibuilder.Client) {
	input := &amplifyuibuilder.GetMetadataInput{
		// AppId: *string, // Required
		// EnvironmentName: *string, // Required
	}

	if len(_amplifyuibuilderAppId) > 0 {
		input.AppId = aws.String(_amplifyuibuilderAppId)
	}
	if len(_amplifyuibuilderEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_amplifyuibuilderEnvironmentName)
	}

	if resp, err := client.GetMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an existing theme for an Amplify app.
func amplifyuibuilder_GetTheme(cfg aws.Config, client *amplifyuibuilder.Client) {
	input := &amplifyuibuilder.GetThemeInput{
		// AppId: *string, // Required
		// EnvironmentName: *string, // Required
		// Id: *string, // Required
	}

	if len(_amplifyuibuilderAppId) > 0 {
		input.AppId = aws.String(_amplifyuibuilderAppId)
	}
	if len(_amplifyuibuilderEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_amplifyuibuilderEnvironmentName)
	}
	if len(_amplifyuibuilderId) > 0 {
		input.Id = aws.String(_amplifyuibuilderId)
	}

	if resp, err := client.GetTheme(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of code generation jobs for a specified Amplify app and
// backend environment.
func amplifyuibuilder_ListCodegenJobs(cfg aws.Config, client *amplifyuibuilder.Client) {
	input := &amplifyuibuilder.ListCodegenJobsInput{
		// AppId: *string, // Required
		// EnvironmentName: *string, // Required
	}

	if len(_amplifyuibuilderAppId) > 0 {
		input.AppId = aws.String(_amplifyuibuilderAppId)
	}
	if len(_amplifyuibuilderEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_amplifyuibuilderEnvironmentName)
	}
	if len(_amplifyuibuilderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _amplifyuibuilderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_amplifyuibuilderNextToken) > 0 {
		input.NextToken = aws.String(_amplifyuibuilderNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCodegenJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*amplifyuibuilder.ListCodegenJobsOutput
	p := amplifyuibuilder.NewListCodegenJobsPaginator(client, input)
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

// Retrieves a list of components for a specified Amplify app and backend
// environment.
func amplifyuibuilder_ListComponents(cfg aws.Config, client *amplifyuibuilder.Client) {
	input := &amplifyuibuilder.ListComponentsInput{
		// AppId: *string, // Required
		// EnvironmentName: *string, // Required
	}

	if len(_amplifyuibuilderAppId) > 0 {
		input.AppId = aws.String(_amplifyuibuilderAppId)
	}
	if len(_amplifyuibuilderEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_amplifyuibuilderEnvironmentName)
	}
	if len(_amplifyuibuilderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _amplifyuibuilderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_amplifyuibuilderNextToken) > 0 {
		input.NextToken = aws.String(_amplifyuibuilderNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListComponents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*amplifyuibuilder.ListComponentsOutput
	p := amplifyuibuilder.NewListComponentsPaginator(client, input)
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

// Retrieves a list of forms for a specified Amplify app and backend environment.
func amplifyuibuilder_ListForms(cfg aws.Config, client *amplifyuibuilder.Client) {
	input := &amplifyuibuilder.ListFormsInput{
		// AppId: *string, // Required
		// EnvironmentName: *string, // Required
	}

	if len(_amplifyuibuilderAppId) > 0 {
		input.AppId = aws.String(_amplifyuibuilderAppId)
	}
	if len(_amplifyuibuilderEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_amplifyuibuilderEnvironmentName)
	}
	if len(_amplifyuibuilderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _amplifyuibuilderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_amplifyuibuilderNextToken) > 0 {
		input.NextToken = aws.String(_amplifyuibuilderNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListForms(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*amplifyuibuilder.ListFormsOutput
	p := amplifyuibuilder.NewListFormsPaginator(client, input)
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

// Returns a list of tags for a specified Amazon Resource Name (ARN).
func amplifyuibuilder_ListTagsForResource(cfg aws.Config, client *amplifyuibuilder.Client) {
	input := &amplifyuibuilder.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_amplifyuibuilderResourceArn) > 0 {
		input.ResourceArn = aws.String(_amplifyuibuilderResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of themes for a specified Amplify app and backend environment.
func amplifyuibuilder_ListThemes(cfg aws.Config, client *amplifyuibuilder.Client) {
	input := &amplifyuibuilder.ListThemesInput{
		// AppId: *string, // Required
		// EnvironmentName: *string, // Required
	}

	if len(_amplifyuibuilderAppId) > 0 {
		input.AppId = aws.String(_amplifyuibuilderAppId)
	}
	if len(_amplifyuibuilderEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_amplifyuibuilderEnvironmentName)
	}
	if len(_amplifyuibuilderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _amplifyuibuilderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_amplifyuibuilderNextToken) > 0 {
		input.NextToken = aws.String(_amplifyuibuilderNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListThemes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*amplifyuibuilder.ListThemesOutput
	p := amplifyuibuilder.NewListThemesPaginator(client, input)
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

// Stores the metadata information about a feature on a form.
func amplifyuibuilder_PutMetadataFlag(cfg aws.Config, client *amplifyuibuilder.Client) {
	input := &amplifyuibuilder.PutMetadataFlagInput{
		// AppId: *string, // Required
		// Body: *types.PutMetadataFlagBody, // Required
		// EnvironmentName: *string, // Required
		// FeatureName: *string, // Required
	}

	if len(_amplifyuibuilderAppId) > 0 {
		input.AppId = aws.String(_amplifyuibuilderAppId)
	}
	if len(_amplifyuibuilderBody) > 0 {
		if err := assignInputField(input, "Body", _amplifyuibuilderBody); err != nil {
			log.Errorf("invalid --body: %s", err.Error())
			return
		}
	}
	if len(_amplifyuibuilderEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_amplifyuibuilderEnvironmentName)
	}
	if len(_amplifyuibuilderFeatureName) > 0 {
		input.FeatureName = aws.String(_amplifyuibuilderFeatureName)
	}

	if resp, err := client.PutMetadataFlag(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is for internal use.
// Amplify uses this action to refresh a previously issued access token that might
// have expired.
func amplifyuibuilder_RefreshToken(cfg aws.Config, client *amplifyuibuilder.Client) {
	input := &amplifyuibuilder.RefreshTokenInput{
		// Provider: types.TokenProviders, // Required
		// RefreshTokenBody: *types.RefreshTokenRequestBody, // Required
	}

	if len(_amplifyuibuilderProvider) > 0 {
		if err := assignInputField(input, "Provider", _amplifyuibuilderProvider); err != nil {
			log.Errorf("invalid --provider: %s", err.Error())
			return
		}
	}
	if len(_amplifyuibuilderRefreshTokenBody) > 0 {
		if err := assignInputField(input, "RefreshTokenBody", _amplifyuibuilderRefreshTokenBody); err != nil {
			log.Errorf("invalid --refresh-token-body: %s", err.Error())
			return
		}
	}

	if resp, err := client.RefreshToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a code generation job for a specified Amplify app and backend
// environment.
func amplifyuibuilder_StartCodegenJob(cfg aws.Config, client *amplifyuibuilder.Client) {
	input := &amplifyuibuilder.StartCodegenJobInput{
		// AppId: *string, // Required
		// CodegenJobToCreate: *types.StartCodegenJobData, // Required
		// EnvironmentName: *string, // Required
	}

	if len(_amplifyuibuilderAppId) > 0 {
		input.AppId = aws.String(_amplifyuibuilderAppId)
	}
	if len(_amplifyuibuilderCodegenJobToCreate) > 0 {
		if err := assignInputField(input, "CodegenJobToCreate", _amplifyuibuilderCodegenJobToCreate); err != nil {
			log.Errorf("invalid --codegen-job-to-create: %s", err.Error())
			return
		}
	}
	if len(_amplifyuibuilderEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_amplifyuibuilderEnvironmentName)
	}
	if len(_amplifyuibuilderClientToken) > 0 {
		input.ClientToken = aws.String(_amplifyuibuilderClientToken)
	}

	if resp, err := client.StartCodegenJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tags the resource with a tag key and value.
func amplifyuibuilder_TagResource(cfg aws.Config, client *amplifyuibuilder.Client) {
	input := &amplifyuibuilder.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_amplifyuibuilderResourceArn) > 0 {
		input.ResourceArn = aws.String(_amplifyuibuilderResourceArn)
	}
	if len(_amplifyuibuilderTags) > 0 {
		if err := assignInputField(input, "Tags", _amplifyuibuilderTags); err != nil {
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

// Untags a resource with a specified Amazon Resource Name (ARN).
func amplifyuibuilder_UntagResource(cfg aws.Config, client *amplifyuibuilder.Client) {
	input := &amplifyuibuilder.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_amplifyuibuilderResourceArn) > 0 {
		input.ResourceArn = aws.String(_amplifyuibuilderResourceArn)
	}
	if len(_amplifyuibuilderTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _amplifyuibuilderTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing component.
func amplifyuibuilder_UpdateComponent(cfg aws.Config, client *amplifyuibuilder.Client) {
	input := &amplifyuibuilder.UpdateComponentInput{
		// AppId: *string, // Required
		// EnvironmentName: *string, // Required
		// Id: *string, // Required
		// UpdatedComponent: *types.UpdateComponentData, // Required
	}

	if len(_amplifyuibuilderAppId) > 0 {
		input.AppId = aws.String(_amplifyuibuilderAppId)
	}
	if len(_amplifyuibuilderEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_amplifyuibuilderEnvironmentName)
	}
	if len(_amplifyuibuilderId) > 0 {
		input.Id = aws.String(_amplifyuibuilderId)
	}
	if len(_amplifyuibuilderUpdatedComponent) > 0 {
		if err := assignInputField(input, "UpdatedComponent", _amplifyuibuilderUpdatedComponent); err != nil {
			log.Errorf("invalid --updated-component: %s", err.Error())
			return
		}
	}
	if len(_amplifyuibuilderClientToken) > 0 {
		input.ClientToken = aws.String(_amplifyuibuilderClientToken)
	}

	if resp, err := client.UpdateComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing form.
func amplifyuibuilder_UpdateForm(cfg aws.Config, client *amplifyuibuilder.Client) {
	input := &amplifyuibuilder.UpdateFormInput{
		// AppId: *string, // Required
		// EnvironmentName: *string, // Required
		// Id: *string, // Required
		// UpdatedForm: *types.UpdateFormData, // Required
	}

	if len(_amplifyuibuilderAppId) > 0 {
		input.AppId = aws.String(_amplifyuibuilderAppId)
	}
	if len(_amplifyuibuilderEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_amplifyuibuilderEnvironmentName)
	}
	if len(_amplifyuibuilderId) > 0 {
		input.Id = aws.String(_amplifyuibuilderId)
	}
	if len(_amplifyuibuilderUpdatedForm) > 0 {
		if err := assignInputField(input, "UpdatedForm", _amplifyuibuilderUpdatedForm); err != nil {
			log.Errorf("invalid --updated-form: %s", err.Error())
			return
		}
	}
	if len(_amplifyuibuilderClientToken) > 0 {
		input.ClientToken = aws.String(_amplifyuibuilderClientToken)
	}

	if resp, err := client.UpdateForm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing theme.
func amplifyuibuilder_UpdateTheme(cfg aws.Config, client *amplifyuibuilder.Client) {
	input := &amplifyuibuilder.UpdateThemeInput{
		// AppId: *string, // Required
		// EnvironmentName: *string, // Required
		// Id: *string, // Required
		// UpdatedTheme: *types.UpdateThemeData, // Required
	}

	if len(_amplifyuibuilderAppId) > 0 {
		input.AppId = aws.String(_amplifyuibuilderAppId)
	}
	if len(_amplifyuibuilderEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_amplifyuibuilderEnvironmentName)
	}
	if len(_amplifyuibuilderId) > 0 {
		input.Id = aws.String(_amplifyuibuilderId)
	}
	if len(_amplifyuibuilderUpdatedTheme) > 0 {
		if err := assignInputField(input, "UpdatedTheme", _amplifyuibuilderUpdatedTheme); err != nil {
			log.Errorf("invalid --updated-theme: %s", err.Error())
			return
		}
	}
	if len(_amplifyuibuilderClientToken) > 0 {
		input.ClientToken = aws.String(_amplifyuibuilderClientToken)
	}

	if resp, err := client.UpdateTheme(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_amplifyuibuilderCmd)
	_amplifyuibuilderCmd.Flags().SortFlags = false

	_amplifyuibuilderCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_amplifyuibuilderCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_amplifyuibuilderCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_amplifyuibuilderCmd.Flags().StringVarP(&_amplifyuibuilderAppId, "app-id", "", "", "App ID")
	_amplifyuibuilderCmd.Flags().StringVarP(&_amplifyuibuilderBody, "body", "", "", "Body")
	_amplifyuibuilderCmd.Flags().StringVarP(&_amplifyuibuilderClientToken, "client-token", "", "", "Client Token")
	_amplifyuibuilderCmd.Flags().StringVarP(&_amplifyuibuilderCodegenJobToCreate, "codegen-job-to-create", "", "", "Codegen Job To Create")
	_amplifyuibuilderCmd.Flags().StringVarP(&_amplifyuibuilderComponentToCreate, "component-to-create", "", "", "Component To Create")
	_amplifyuibuilderCmd.Flags().StringVarP(&_amplifyuibuilderEnvironmentName, "environment-name", "", "", "Environment Name")
	_amplifyuibuilderCmd.Flags().StringVarP(&_amplifyuibuilderFeatureName, "feature-name", "", "", "Feature Name")
	_amplifyuibuilderCmd.Flags().StringVarP(&_amplifyuibuilderFormToCreate, "form-to-create", "", "", "Form To Create")
	_amplifyuibuilderCmd.Flags().StringVarP(&_amplifyuibuilderId, "id", "", "", "ID")
	_amplifyuibuilderCmd.Flags().StringVarP(&_amplifyuibuilderMaxResults, "max-results", "", "", "Max Results")
	_amplifyuibuilderCmd.Flags().StringVarP(&_amplifyuibuilderNextToken, "next-token", "", "", "Next Token")
	_amplifyuibuilderCmd.Flags().StringVarP(&_amplifyuibuilderProvider, "provider", "", "", "Provider")
	_amplifyuibuilderCmd.Flags().StringVarP(&_amplifyuibuilderRefreshTokenBody, "refresh-token-body", "", "", "Refresh Token Body")
	_amplifyuibuilderCmd.Flags().StringVarP(&_amplifyuibuilderRequest, "request", "", "", "Request")
	_amplifyuibuilderCmd.Flags().StringVarP(&_amplifyuibuilderResourceArn, "resource-arn", "", "", "Resource ARN")
	_amplifyuibuilderCmd.Flags().StringSliceVarP(&_amplifyuibuilderTagKeys, "tag-keys", "", nil, "Tag Keys")
	_amplifyuibuilderCmd.Flags().StringVarP(&_amplifyuibuilderTags, "tags", "", "", "Tags")
	_amplifyuibuilderCmd.Flags().StringVarP(&_amplifyuibuilderThemeToCreate, "theme-to-create", "", "", "Theme To Create")
	_amplifyuibuilderCmd.Flags().StringVarP(&_amplifyuibuilderUpdatedComponent, "updated-component", "", "", "Updated Component")
	_amplifyuibuilderCmd.Flags().StringVarP(&_amplifyuibuilderUpdatedForm, "updated-form", "", "", "Updated Form")
	_amplifyuibuilderCmd.Flags().StringVarP(&_amplifyuibuilderUpdatedTheme, "updated-theme", "", "", "Updated Theme")

	_amplifyuibuilderCmd.Flags().BoolVarP(&_amplifyuibuilderCreateComponent, "create-component", "", false, "Create Component")
	_amplifyuibuilderCmd.Flags().BoolVarP(&_amplifyuibuilderCreateForm, "create-form", "", false, "Create Form")
	_amplifyuibuilderCmd.Flags().BoolVarP(&_amplifyuibuilderCreateTheme, "create-theme", "", false, "Create Theme")
	_amplifyuibuilderCmd.Flags().BoolVarP(&_amplifyuibuilderDeleteComponent, "delete-component", "", false, "Delete Component")
	_amplifyuibuilderCmd.Flags().BoolVarP(&_amplifyuibuilderDeleteForm, "delete-form", "", false, "Delete Form")
	_amplifyuibuilderCmd.Flags().BoolVarP(&_amplifyuibuilderDeleteTheme, "delete-theme", "", false, "Delete Theme")
	_amplifyuibuilderCmd.Flags().BoolVarP(&_amplifyuibuilderExchangeCodeForToken, "exchange-code-for-token", "", false, "Exchange Code For Token")
	_amplifyuibuilderCmd.Flags().BoolVarP(&_amplifyuibuilderExportComponents, "export-components", "", false, "Export Components")
	_amplifyuibuilderCmd.Flags().BoolVarP(&_amplifyuibuilderExportForms, "export-forms", "", false, "Export Forms")
	_amplifyuibuilderCmd.Flags().BoolVarP(&_amplifyuibuilderExportThemes, "export-themes", "", false, "Export Themes")
	_amplifyuibuilderCmd.Flags().BoolVarP(&_amplifyuibuilderGetCodegenJob, "get-codegen-job", "", false, "Get Codegen Job")
	_amplifyuibuilderCmd.Flags().BoolVarP(&_amplifyuibuilderGetComponent, "get-component", "", false, "Get Component")
	_amplifyuibuilderCmd.Flags().BoolVarP(&_amplifyuibuilderGetForm, "get-form", "", false, "Get Form")
	_amplifyuibuilderCmd.Flags().BoolVarP(&_amplifyuibuilderGetMetadata, "get-metadata", "", false, "Get Metadata")
	_amplifyuibuilderCmd.Flags().BoolVarP(&_amplifyuibuilderGetTheme, "get-theme", "", false, "Get Theme")
	_amplifyuibuilderCmd.Flags().BoolVarP(&_amplifyuibuilderListCodegenJobs, "list-codegen-jobs", "", false, "List Codegen Jobs")
	_amplifyuibuilderCmd.Flags().BoolVarP(&_amplifyuibuilderListComponents, "list-components", "", false, "List Components")
	_amplifyuibuilderCmd.Flags().BoolVarP(&_amplifyuibuilderListForms, "list-forms", "", false, "List Forms")
	_amplifyuibuilderCmd.Flags().BoolVarP(&_amplifyuibuilderListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_amplifyuibuilderCmd.Flags().BoolVarP(&_amplifyuibuilderListThemes, "list-themes", "", false, "List Themes")
	_amplifyuibuilderCmd.Flags().BoolVarP(&_amplifyuibuilderPutMetadataFlag, "put-metadata-flag", "", false, "Put Metadata Flag")
	_amplifyuibuilderCmd.Flags().BoolVarP(&_amplifyuibuilderRefreshToken, "refresh-token", "", false, "Refresh Token")
	_amplifyuibuilderCmd.Flags().BoolVarP(&_amplifyuibuilderStartCodegenJob, "start-codegen-job", "", false, "Start Codegen Job")
	_amplifyuibuilderCmd.Flags().BoolVarP(&_amplifyuibuilderTagResource, "tag-resource", "", false, "Tag Resource")
	_amplifyuibuilderCmd.Flags().BoolVarP(&_amplifyuibuilderUntagResource, "untag-resource", "", false, "Untag Resource")
	_amplifyuibuilderCmd.Flags().BoolVarP(&_amplifyuibuilderUpdateComponent, "update-component", "", false, "Update Component")
	_amplifyuibuilderCmd.Flags().BoolVarP(&_amplifyuibuilderUpdateForm, "update-form", "", false, "Update Form")
	_amplifyuibuilderCmd.Flags().BoolVarP(&_amplifyuibuilderUpdateTheme, "update-theme", "", false, "Update Theme")

}
