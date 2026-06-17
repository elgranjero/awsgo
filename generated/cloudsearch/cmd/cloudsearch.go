package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cloudsearchCmd represents the cloudsearch command
var _cloudsearchCmd = &cobra.Command{
	Use:   "cloudsearch",
	Short: "AWS cloudsearch CLI",
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
		client := cloudsearch.NewFromConfig(cfg)
		if _cloudsearchBuildSuggesters {
			cloudsearch_BuildSuggesters(cfg, client)
			return
		}
		if _cloudsearchCreateDomain {
			cloudsearch_CreateDomain(cfg, client)
			return
		}
		if _cloudsearchDefineAnalysisScheme {
			cloudsearch_DefineAnalysisScheme(cfg, client)
			return
		}
		if _cloudsearchDefineExpression {
			cloudsearch_DefineExpression(cfg, client)
			return
		}
		if _cloudsearchDefineIndexField {
			cloudsearch_DefineIndexField(cfg, client)
			return
		}
		if _cloudsearchDefineSuggester {
			cloudsearch_DefineSuggester(cfg, client)
			return
		}
		if _cloudsearchDeleteAnalysisScheme {
			cloudsearch_DeleteAnalysisScheme(cfg, client)
			return
		}
		if _cloudsearchDeleteDomain {
			cloudsearch_DeleteDomain(cfg, client)
			return
		}
		if _cloudsearchDeleteExpression {
			cloudsearch_DeleteExpression(cfg, client)
			return
		}
		if _cloudsearchDeleteIndexField {
			cloudsearch_DeleteIndexField(cfg, client)
			return
		}
		if _cloudsearchDeleteSuggester {
			cloudsearch_DeleteSuggester(cfg, client)
			return
		}
		if _cloudsearchDescribeAnalysisSchemes {
			cloudsearch_DescribeAnalysisSchemes(cfg, client)
			return
		}
		if _cloudsearchDescribeAvailabilityOptions {
			cloudsearch_DescribeAvailabilityOptions(cfg, client)
			return
		}
		if _cloudsearchDescribeDomainEndpointOptions {
			cloudsearch_DescribeDomainEndpointOptions(cfg, client)
			return
		}
		if _cloudsearchDescribeDomains {
			cloudsearch_DescribeDomains(cfg, client)
			return
		}
		if _cloudsearchDescribeExpressions {
			cloudsearch_DescribeExpressions(cfg, client)
			return
		}
		if _cloudsearchDescribeIndexFields {
			cloudsearch_DescribeIndexFields(cfg, client)
			return
		}
		if _cloudsearchDescribeScalingParameters {
			cloudsearch_DescribeScalingParameters(cfg, client)
			return
		}
		if _cloudsearchDescribeServiceAccessPolicies {
			cloudsearch_DescribeServiceAccessPolicies(cfg, client)
			return
		}
		if _cloudsearchDescribeSuggesters {
			cloudsearch_DescribeSuggesters(cfg, client)
			return
		}
		if _cloudsearchIndexDocuments {
			cloudsearch_IndexDocuments(cfg, client)
			return
		}
		if _cloudsearchListDomainNames {
			cloudsearch_ListDomainNames(cfg, client)
			return
		}
		if _cloudsearchUpdateAvailabilityOptions {
			cloudsearch_UpdateAvailabilityOptions(cfg, client)
			return
		}
		if _cloudsearchUpdateDomainEndpointOptions {
			cloudsearch_UpdateDomainEndpointOptions(cfg, client)
			return
		}
		if _cloudsearchUpdateScalingParameters {
			cloudsearch_UpdateScalingParameters(cfg, client)
			return
		}
		if _cloudsearchUpdateServiceAccessPolicies {
			cloudsearch_UpdateServiceAccessPolicies(cfg, client)
			return
		}

	},
}

var (
	_cloudsearchBuildSuggesters               bool
	_cloudsearchCreateDomain                  bool
	_cloudsearchDefineAnalysisScheme          bool
	_cloudsearchDefineExpression              bool
	_cloudsearchDefineIndexField              bool
	_cloudsearchDefineSuggester               bool
	_cloudsearchDeleteAnalysisScheme          bool
	_cloudsearchDeleteDomain                  bool
	_cloudsearchDeleteExpression              bool
	_cloudsearchDeleteIndexField              bool
	_cloudsearchDeleteSuggester               bool
	_cloudsearchDescribeAnalysisSchemes       bool
	_cloudsearchDescribeAvailabilityOptions   bool
	_cloudsearchDescribeDomainEndpointOptions bool
	_cloudsearchDescribeDomains               bool
	_cloudsearchDescribeExpressions           bool
	_cloudsearchDescribeIndexFields           bool
	_cloudsearchDescribeScalingParameters     bool
	_cloudsearchDescribeServiceAccessPolicies bool
	_cloudsearchDescribeSuggesters            bool
	_cloudsearchIndexDocuments                bool
	_cloudsearchListDomainNames               bool
	_cloudsearchUpdateAvailabilityOptions     bool
	_cloudsearchUpdateDomainEndpointOptions   bool
	_cloudsearchUpdateScalingParameters       bool
	_cloudsearchUpdateServiceAccessPolicies   bool

	_cloudsearchAccessPolicies        string
	_cloudsearchAnalysisScheme        string
	_cloudsearchAnalysisSchemeName    string
	_cloudsearchAnalysisSchemeNames   []string
	_cloudsearchDeployed              string
	_cloudsearchDomainEndpointOptions string
	_cloudsearchDomainName            string
	_cloudsearchDomainNames           []string
	_cloudsearchExpression            string
	_cloudsearchExpressionName        string
	_cloudsearchExpressionNames       []string
	_cloudsearchFieldNames            []string
	_cloudsearchIndexField            string
	_cloudsearchIndexFieldName        string
	_cloudsearchMultiAZ               string
	_cloudsearchScalingParameters     string
	_cloudsearchSuggester             string
	_cloudsearchSuggesterName         string
	_cloudsearchSuggesterNames        []string
)

// Indexes the search suggestions. For more information, see [Configuring Suggesters] in the Amazon
// CloudSearch Developer Guide.
//
// [Configuring Suggesters]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/getting-suggestions.html#configuring-suggesters
func cloudsearch_BuildSuggesters(cfg aws.Config, client *cloudsearch.Client) {
	input := &cloudsearch.BuildSuggestersInput{
		// DomainName: *string, // Required
	}

	if len(_cloudsearchDomainName) > 0 {
		input.DomainName = aws.String(_cloudsearchDomainName)
	}

	if resp, err := client.BuildSuggesters(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new search domain. For more information, see [Creating a Search Domain] in the Amazon
// CloudSearch Developer Guide.
//
// [Creating a Search Domain]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/creating-domains.html
func cloudsearch_CreateDomain(cfg aws.Config, client *cloudsearch.Client) {
	input := &cloudsearch.CreateDomainInput{
		// DomainName: *string, // Required
	}

	if len(_cloudsearchDomainName) > 0 {
		input.DomainName = aws.String(_cloudsearchDomainName)
	}

	if resp, err := client.CreateDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Configures an analysis scheme that can be applied to a text or text-array field
// to define language-specific text processing options. For more information, see [Configuring Analysis Schemes]
// in the Amazon CloudSearch Developer Guide.
//
// [Configuring Analysis Schemes]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-analysis-schemes.html
func cloudsearch_DefineAnalysisScheme(cfg aws.Config, client *cloudsearch.Client) {
	input := &cloudsearch.DefineAnalysisSchemeInput{
		// AnalysisScheme: *types.AnalysisScheme, // Required
		// DomainName: *string, // Required
	}

	if len(_cloudsearchAnalysisScheme) > 0 {
		if err := assignInputField(input, "AnalysisScheme", _cloudsearchAnalysisScheme); err != nil {
			log.Errorf("invalid --analysis-scheme: %s", err.Error())
			return
		}
	}
	if len(_cloudsearchDomainName) > 0 {
		input.DomainName = aws.String(_cloudsearchDomainName)
	}

	if resp, err := client.DefineAnalysisScheme(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Configures an Expression for the search domain. Used to create new expressions and modify
// existing ones. If the expression exists, the new configuration replaces the old
// one. For more information, see [Configuring Expressions]in the Amazon CloudSearch Developer Guide.
//
// [Configuring Expressions]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-expressions.html
func cloudsearch_DefineExpression(cfg aws.Config, client *cloudsearch.Client) {
	input := &cloudsearch.DefineExpressionInput{
		// DomainName: *string, // Required
		// Expression: *types.Expression, // Required
	}

	if len(_cloudsearchDomainName) > 0 {
		input.DomainName = aws.String(_cloudsearchDomainName)
	}
	if len(_cloudsearchExpression) > 0 {
		if err := assignInputField(input, "Expression", _cloudsearchExpression); err != nil {
			log.Errorf("invalid --expression: %s", err.Error())
			return
		}
	}

	if resp, err := client.DefineExpression(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Configures an IndexField for the search domain. Used to create new fields and modify
// existing ones. You must specify the name of the domain you are configuring and
// an index field configuration. The index field configuration specifies a unique
// name, the index field type, and the options you want to configure for the field.
// The options you can specify depend on the IndexFieldType. If the field exists, the new
// configuration replaces the old one. For more information, see [Configuring Index Fields]in the Amazon
// CloudSearch Developer Guide.
//
// [Configuring Index Fields]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-index-fields.html
func cloudsearch_DefineIndexField(cfg aws.Config, client *cloudsearch.Client) {
	input := &cloudsearch.DefineIndexFieldInput{
		// DomainName: *string, // Required
		// IndexField: *types.IndexField, // Required
	}

	if len(_cloudsearchDomainName) > 0 {
		input.DomainName = aws.String(_cloudsearchDomainName)
	}
	if len(_cloudsearchIndexField) > 0 {
		if err := assignInputField(input, "IndexField", _cloudsearchIndexField); err != nil {
			log.Errorf("invalid --index-field: %s", err.Error())
			return
		}
	}

	if resp, err := client.DefineIndexField(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Configures a suggester for a domain. A suggester enables you to display
// possible matches before users finish typing their queries. When you configure a
// suggester, you must specify the name of the text field you want to search for
// possible matches and a unique name for the suggester. For more information, see [Getting Search Suggestions]
// in the Amazon CloudSearch Developer Guide.
//
// [Getting Search Suggestions]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/getting-suggestions.html
func cloudsearch_DefineSuggester(cfg aws.Config, client *cloudsearch.Client) {
	input := &cloudsearch.DefineSuggesterInput{
		// DomainName: *string, // Required
		// Suggester: *types.Suggester, // Required
	}

	if len(_cloudsearchDomainName) > 0 {
		input.DomainName = aws.String(_cloudsearchDomainName)
	}
	if len(_cloudsearchSuggester) > 0 {
		if err := assignInputField(input, "Suggester", _cloudsearchSuggester); err != nil {
			log.Errorf("invalid --suggester: %s", err.Error())
			return
		}
	}

	if resp, err := client.DefineSuggester(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an analysis scheme. For more information, see [Configuring Analysis Schemes] in the Amazon
// CloudSearch Developer Guide.
//
// [Configuring Analysis Schemes]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-analysis-schemes.html
func cloudsearch_DeleteAnalysisScheme(cfg aws.Config, client *cloudsearch.Client) {
	input := &cloudsearch.DeleteAnalysisSchemeInput{
		// AnalysisSchemeName: *string, // Required
		// DomainName: *string, // Required
	}

	if len(_cloudsearchAnalysisSchemeName) > 0 {
		input.AnalysisSchemeName = aws.String(_cloudsearchAnalysisSchemeName)
	}
	if len(_cloudsearchDomainName) > 0 {
		input.DomainName = aws.String(_cloudsearchDomainName)
	}

	if resp, err := client.DeleteAnalysisScheme(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Permanently deletes a search domain and all of its data. Once a domain has been
// deleted, it cannot be recovered. For more information, see [Deleting a Search Domain]in the Amazon
// CloudSearch Developer Guide.
//
// [Deleting a Search Domain]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/deleting-domains.html
func cloudsearch_DeleteDomain(cfg aws.Config, client *cloudsearch.Client) {
	input := &cloudsearch.DeleteDomainInput{
		// DomainName: *string, // Required
	}

	if len(_cloudsearchDomainName) > 0 {
		input.DomainName = aws.String(_cloudsearchDomainName)
	}

	if resp, err := client.DeleteDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes an Expression from the search domain. For more information, see [Configuring Expressions] in the Amazon
// CloudSearch Developer Guide.
//
// [Configuring Expressions]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-expressions.html
func cloudsearch_DeleteExpression(cfg aws.Config, client *cloudsearch.Client) {
	input := &cloudsearch.DeleteExpressionInput{
		// DomainName: *string, // Required
		// ExpressionName: *string, // Required
	}

	if len(_cloudsearchDomainName) > 0 {
		input.DomainName = aws.String(_cloudsearchDomainName)
	}
	if len(_cloudsearchExpressionName) > 0 {
		input.ExpressionName = aws.String(_cloudsearchExpressionName)
	}

	if resp, err := client.DeleteExpression(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes an IndexField from the search domain. For more information, see [Configuring Index Fields] in the Amazon
// CloudSearch Developer Guide.
//
// [Configuring Index Fields]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-index-fields.html
func cloudsearch_DeleteIndexField(cfg aws.Config, client *cloudsearch.Client) {
	input := &cloudsearch.DeleteIndexFieldInput{
		// DomainName: *string, // Required
		// IndexFieldName: *string, // Required
	}

	if len(_cloudsearchDomainName) > 0 {
		input.DomainName = aws.String(_cloudsearchDomainName)
	}
	if len(_cloudsearchIndexFieldName) > 0 {
		input.IndexFieldName = aws.String(_cloudsearchIndexFieldName)
	}

	if resp, err := client.DeleteIndexField(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a suggester. For more information, see [Getting Search Suggestions] in the Amazon CloudSearch
// Developer Guide.
//
// [Getting Search Suggestions]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/getting-suggestions.html
func cloudsearch_DeleteSuggester(cfg aws.Config, client *cloudsearch.Client) {
	input := &cloudsearch.DeleteSuggesterInput{
		// DomainName: *string, // Required
		// SuggesterName: *string, // Required
	}

	if len(_cloudsearchDomainName) > 0 {
		input.DomainName = aws.String(_cloudsearchDomainName)
	}
	if len(_cloudsearchSuggesterName) > 0 {
		input.SuggesterName = aws.String(_cloudsearchSuggesterName)
	}

	if resp, err := client.DeleteSuggester(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the analysis schemes configured for a domain. An analysis scheme defines
// language-specific text processing options for a text field. Can be limited to
// specific analysis schemes by name. By default, shows all analysis schemes and
// includes any pending changes to the configuration. Set the Deployed option to
// true to show the active configuration and exclude pending changes. For more
// information, see [Configuring Analysis Schemes]in the Amazon CloudSearch Developer Guide.
//
// [Configuring Analysis Schemes]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-analysis-schemes.html
func cloudsearch_DescribeAnalysisSchemes(cfg aws.Config, client *cloudsearch.Client) {
	input := &cloudsearch.DescribeAnalysisSchemesInput{
		// DomainName: *string, // Required
	}

	if len(_cloudsearchDomainName) > 0 {
		input.DomainName = aws.String(_cloudsearchDomainName)
	}
	if len(_cloudsearchAnalysisSchemeNames) > 0 {
		input.AnalysisSchemeNames = append([]string(nil), _cloudsearchAnalysisSchemeNames...)
	}
	if len(_cloudsearchDeployed) > 0 {
		if err := assignInputField(input, "Deployed", _cloudsearchDeployed); err != nil {
			log.Errorf("invalid --deployed: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeAnalysisSchemes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the availability options configured for a domain. By default, shows the
// configuration with any pending changes. Set the Deployed option to true to show
// the active configuration and exclude pending changes. For more information, see [Configuring Availability Options]
// in the Amazon CloudSearch Developer Guide.
//
// [Configuring Availability Options]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-availability-options.html
func cloudsearch_DescribeAvailabilityOptions(cfg aws.Config, client *cloudsearch.Client) {
	input := &cloudsearch.DescribeAvailabilityOptionsInput{
		// DomainName: *string, // Required
	}

	if len(_cloudsearchDomainName) > 0 {
		input.DomainName = aws.String(_cloudsearchDomainName)
	}
	if len(_cloudsearchDeployed) > 0 {
		if err := assignInputField(input, "Deployed", _cloudsearchDeployed); err != nil {
			log.Errorf("invalid --deployed: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeAvailabilityOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the domain's endpoint options, specifically whether all requests to the
// domain must arrive over HTTPS. For more information, see [Configuring Domain Endpoint Options]in the Amazon
// CloudSearch Developer Guide.
//
// [Configuring Domain Endpoint Options]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-domain-endpoint-options.html
func cloudsearch_DescribeDomainEndpointOptions(cfg aws.Config, client *cloudsearch.Client) {
	input := &cloudsearch.DescribeDomainEndpointOptionsInput{
		// DomainName: *string, // Required
	}

	if len(_cloudsearchDomainName) > 0 {
		input.DomainName = aws.String(_cloudsearchDomainName)
	}
	if len(_cloudsearchDeployed) > 0 {
		if err := assignInputField(input, "Deployed", _cloudsearchDeployed); err != nil {
			log.Errorf("invalid --deployed: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeDomainEndpointOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the search domains owned by this account. Can be limited
// to specific domains. Shows all domains by default. To get the number of
// searchable documents in a domain, use the console or submit a matchall request
// to your domain's search endpoint: q=matchall&q.parser=structured&size=0 . For
// more information, see [Getting Information about a Search Domain]in the Amazon CloudSearch Developer Guide.
//
// [Getting Information about a Search Domain]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/getting-domain-info.html
func cloudsearch_DescribeDomains(cfg aws.Config, client *cloudsearch.Client) {
	input := &cloudsearch.DescribeDomainsInput{}

	if len(_cloudsearchDomainNames) > 0 {
		input.DomainNames = append([]string(nil), _cloudsearchDomainNames...)
	}

	if resp, err := client.DescribeDomains(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the expressions configured for the search domain. Can be limited to
// specific expressions by name. By default, shows all expressions and includes any
// pending changes to the configuration. Set the Deployed option to true to show
// the active configuration and exclude pending changes. For more information, see [Configuring Expressions]
// in the Amazon CloudSearch Developer Guide.
//
// [Configuring Expressions]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-expressions.html
func cloudsearch_DescribeExpressions(cfg aws.Config, client *cloudsearch.Client) {
	input := &cloudsearch.DescribeExpressionsInput{
		// DomainName: *string, // Required
	}

	if len(_cloudsearchDomainName) > 0 {
		input.DomainName = aws.String(_cloudsearchDomainName)
	}
	if len(_cloudsearchDeployed) > 0 {
		if err := assignInputField(input, "Deployed", _cloudsearchDeployed); err != nil {
			log.Errorf("invalid --deployed: %s", err.Error())
			return
		}
	}
	if len(_cloudsearchExpressionNames) > 0 {
		input.ExpressionNames = append([]string(nil), _cloudsearchExpressionNames...)
	}

	if resp, err := client.DescribeExpressions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the index fields configured for the search domain. Can
// be limited to specific fields by name. By default, shows all fields and includes
// any pending changes to the configuration. Set the Deployed option to true to
// show the active configuration and exclude pending changes. For more information,
// see [Getting Domain Information]in the Amazon CloudSearch Developer Guide.
//
// [Getting Domain Information]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/getting-domain-info.html
func cloudsearch_DescribeIndexFields(cfg aws.Config, client *cloudsearch.Client) {
	input := &cloudsearch.DescribeIndexFieldsInput{
		// DomainName: *string, // Required
	}

	if len(_cloudsearchDomainName) > 0 {
		input.DomainName = aws.String(_cloudsearchDomainName)
	}
	if len(_cloudsearchDeployed) > 0 {
		if err := assignInputField(input, "Deployed", _cloudsearchDeployed); err != nil {
			log.Errorf("invalid --deployed: %s", err.Error())
			return
		}
	}
	if len(_cloudsearchFieldNames) > 0 {
		input.FieldNames = append([]string(nil), _cloudsearchFieldNames...)
	}

	if resp, err := client.DescribeIndexFields(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the scaling parameters configured for a domain. A domain's scaling
// parameters specify the desired search instance type and replication count. For
// more information, see [Configuring Scaling Options]in the Amazon CloudSearch Developer Guide.
//
// [Configuring Scaling Options]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-scaling-options.html
func cloudsearch_DescribeScalingParameters(cfg aws.Config, client *cloudsearch.Client) {
	input := &cloudsearch.DescribeScalingParametersInput{
		// DomainName: *string, // Required
	}

	if len(_cloudsearchDomainName) > 0 {
		input.DomainName = aws.String(_cloudsearchDomainName)
	}

	if resp, err := client.DescribeScalingParameters(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the access policies that control access to the domain's
// document and search endpoints. By default, shows the configuration with any
// pending changes. Set the Deployed option to true to show the active
// configuration and exclude pending changes. For more information, see [Configuring Access for a Search Domain]in the
// Amazon CloudSearch Developer Guide.
//
// [Configuring Access for a Search Domain]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-access.html
func cloudsearch_DescribeServiceAccessPolicies(cfg aws.Config, client *cloudsearch.Client) {
	input := &cloudsearch.DescribeServiceAccessPoliciesInput{
		// DomainName: *string, // Required
	}

	if len(_cloudsearchDomainName) > 0 {
		input.DomainName = aws.String(_cloudsearchDomainName)
	}
	if len(_cloudsearchDeployed) > 0 {
		if err := assignInputField(input, "Deployed", _cloudsearchDeployed); err != nil {
			log.Errorf("invalid --deployed: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeServiceAccessPolicies(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the suggesters configured for a domain. A suggester enables you to display
// possible matches before users finish typing their queries. Can be limited to
// specific suggesters by name. By default, shows all suggesters and includes any
// pending changes to the configuration. Set the Deployed option to true to show
// the active configuration and exclude pending changes. For more information, see [Getting Search Suggestions]
// in the Amazon CloudSearch Developer Guide.
//
// [Getting Search Suggestions]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/getting-suggestions.html
func cloudsearch_DescribeSuggesters(cfg aws.Config, client *cloudsearch.Client) {
	input := &cloudsearch.DescribeSuggestersInput{
		// DomainName: *string, // Required
	}

	if len(_cloudsearchDomainName) > 0 {
		input.DomainName = aws.String(_cloudsearchDomainName)
	}
	if len(_cloudsearchDeployed) > 0 {
		if err := assignInputField(input, "Deployed", _cloudsearchDeployed); err != nil {
			log.Errorf("invalid --deployed: %s", err.Error())
			return
		}
	}
	if len(_cloudsearchSuggesterNames) > 0 {
		input.SuggesterNames = append([]string(nil), _cloudsearchSuggesterNames...)
	}

	if resp, err := client.DescribeSuggesters(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tells the search domain to start indexing its documents using the latest
// indexing options. This operation must be invoked to activate options whose OptionStatusis
// RequiresIndexDocuments .
func cloudsearch_IndexDocuments(cfg aws.Config, client *cloudsearch.Client) {
	input := &cloudsearch.IndexDocumentsInput{
		// DomainName: *string, // Required
	}

	if len(_cloudsearchDomainName) > 0 {
		input.DomainName = aws.String(_cloudsearchDomainName)
	}

	if resp, err := client.IndexDocuments(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all search domains owned by an account.
func cloudsearch_ListDomainNames(cfg aws.Config, client *cloudsearch.Client) {
	input := &cloudsearch.ListDomainNamesInput{}

	if resp, err := client.ListDomainNames(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Configures the availability options for a domain. Enabling the Multi-AZ option
// expands an Amazon CloudSearch domain to an additional Availability Zone in the
// same Region to increase fault tolerance in the event of a service disruption.
// Changes to the Multi-AZ option can take about half an hour to become active. For
// more information, see [Configuring Availability Options]in the Amazon CloudSearch Developer Guide.
//
// [Configuring Availability Options]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-availability-options.html
func cloudsearch_UpdateAvailabilityOptions(cfg aws.Config, client *cloudsearch.Client) {
	input := &cloudsearch.UpdateAvailabilityOptionsInput{
		// DomainName: *string, // Required
		// MultiAZ: *bool, // Required
	}

	if len(_cloudsearchDomainName) > 0 {
		input.DomainName = aws.String(_cloudsearchDomainName)
	}
	if len(_cloudsearchMultiAZ) > 0 {
		if err := assignInputField(input, "MultiAZ", _cloudsearchMultiAZ); err != nil {
			log.Errorf("invalid --multi-az: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAvailabilityOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the domain's endpoint options, specifically whether all requests to the
// domain must arrive over HTTPS. For more information, see [Configuring Domain Endpoint Options]in the Amazon
// CloudSearch Developer Guide.
//
// [Configuring Domain Endpoint Options]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-domain-endpoint-options.html
func cloudsearch_UpdateDomainEndpointOptions(cfg aws.Config, client *cloudsearch.Client) {
	input := &cloudsearch.UpdateDomainEndpointOptionsInput{
		// DomainEndpointOptions: *types.DomainEndpointOptions, // Required
		// DomainName: *string, // Required
	}

	if len(_cloudsearchDomainEndpointOptions) > 0 {
		if err := assignInputField(input, "DomainEndpointOptions", _cloudsearchDomainEndpointOptions); err != nil {
			log.Errorf("invalid --domain-endpoint-options: %s", err.Error())
			return
		}
	}
	if len(_cloudsearchDomainName) > 0 {
		input.DomainName = aws.String(_cloudsearchDomainName)
	}

	if resp, err := client.UpdateDomainEndpointOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Configures scaling parameters for a domain. A domain's scaling parameters
// specify the desired search instance type and replication count. Amazon
// CloudSearch will still automatically scale your domain based on the volume of
// data and traffic, but not below the desired instance type and replication count.
// If the Multi-AZ option is enabled, these values control the resources used per
// Availability Zone. For more information, see [Configuring Scaling Options]in the Amazon CloudSearch
// Developer Guide.
//
// [Configuring Scaling Options]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-scaling-options.html
func cloudsearch_UpdateScalingParameters(cfg aws.Config, client *cloudsearch.Client) {
	input := &cloudsearch.UpdateScalingParametersInput{
		// DomainName: *string, // Required
		// ScalingParameters: *types.ScalingParameters, // Required
	}

	if len(_cloudsearchDomainName) > 0 {
		input.DomainName = aws.String(_cloudsearchDomainName)
	}
	if len(_cloudsearchScalingParameters) > 0 {
		if err := assignInputField(input, "ScalingParameters", _cloudsearchScalingParameters); err != nil {
			log.Errorf("invalid --scaling-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateScalingParameters(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Configures the access rules that control access to the domain's document and
// search endpoints. For more information, see [Configuring Access for an Amazon CloudSearch Domain].
//
// [Configuring Access for an Amazon CloudSearch Domain]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-access.html
func cloudsearch_UpdateServiceAccessPolicies(cfg aws.Config, client *cloudsearch.Client) {
	input := &cloudsearch.UpdateServiceAccessPoliciesInput{
		// AccessPolicies: *string, // Required
		// DomainName: *string, // Required
	}

	if len(_cloudsearchAccessPolicies) > 0 {
		input.AccessPolicies = aws.String(_cloudsearchAccessPolicies)
	}
	if len(_cloudsearchDomainName) > 0 {
		input.DomainName = aws.String(_cloudsearchDomainName)
	}

	if resp, err := client.UpdateServiceAccessPolicies(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_cloudsearchCmd)
	_cloudsearchCmd.Flags().SortFlags = false

	_cloudsearchCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_cloudsearchCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_cloudsearchCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_cloudsearchCmd.Flags().StringVarP(&_cloudsearchAccessPolicies, "access-policies", "", "", "Access Policies")
	_cloudsearchCmd.Flags().StringVarP(&_cloudsearchAnalysisScheme, "analysis-scheme", "", "", "Analysis Scheme")
	_cloudsearchCmd.Flags().StringVarP(&_cloudsearchAnalysisSchemeName, "analysis-scheme-name", "", "", "Analysis Scheme Name")
	_cloudsearchCmd.Flags().StringSliceVarP(&_cloudsearchAnalysisSchemeNames, "analysis-scheme-names", "", nil, "Analysis Scheme Names")
	_cloudsearchCmd.Flags().StringVarP(&_cloudsearchDeployed, "deployed", "", "", "Deployed")
	_cloudsearchCmd.Flags().StringVarP(&_cloudsearchDomainEndpointOptions, "domain-endpoint-options", "", "", "Domain Endpoint Options")
	_cloudsearchCmd.Flags().StringVarP(&_cloudsearchDomainName, "domain-name", "", "", "Domain Name")
	_cloudsearchCmd.Flags().StringSliceVarP(&_cloudsearchDomainNames, "domain-names", "", nil, "Domain Names")
	_cloudsearchCmd.Flags().StringVarP(&_cloudsearchExpression, "expression", "", "", "Expression")
	_cloudsearchCmd.Flags().StringVarP(&_cloudsearchExpressionName, "expression-name", "", "", "Expression Name")
	_cloudsearchCmd.Flags().StringSliceVarP(&_cloudsearchExpressionNames, "expression-names", "", nil, "Expression Names")
	_cloudsearchCmd.Flags().StringSliceVarP(&_cloudsearchFieldNames, "field-names", "", nil, "Field Names")
	_cloudsearchCmd.Flags().StringVarP(&_cloudsearchIndexField, "index-field", "", "", "Index Field")
	_cloudsearchCmd.Flags().StringVarP(&_cloudsearchIndexFieldName, "index-field-name", "", "", "Index Field Name")
	_cloudsearchCmd.Flags().StringVarP(&_cloudsearchMultiAZ, "multi-az", "", "", "Multi AZ")
	_cloudsearchCmd.Flags().StringVarP(&_cloudsearchScalingParameters, "scaling-parameters", "", "", "Scaling Parameters")
	_cloudsearchCmd.Flags().StringVarP(&_cloudsearchSuggester, "suggester", "", "", "Suggester")
	_cloudsearchCmd.Flags().StringVarP(&_cloudsearchSuggesterName, "suggester-name", "", "", "Suggester Name")
	_cloudsearchCmd.Flags().StringSliceVarP(&_cloudsearchSuggesterNames, "suggester-names", "", nil, "Suggester Names")

	_cloudsearchCmd.Flags().BoolVarP(&_cloudsearchBuildSuggesters, "build-suggesters", "", false, "Build Suggesters")
	_cloudsearchCmd.Flags().BoolVarP(&_cloudsearchCreateDomain, "create-domain", "", false, "Create Domain")
	_cloudsearchCmd.Flags().BoolVarP(&_cloudsearchDefineAnalysisScheme, "define-analysis-scheme", "", false, "Define Analysis Scheme")
	_cloudsearchCmd.Flags().BoolVarP(&_cloudsearchDefineExpression, "define-expression", "", false, "Define Expression")
	_cloudsearchCmd.Flags().BoolVarP(&_cloudsearchDefineIndexField, "define-index-field", "", false, "Define Index Field")
	_cloudsearchCmd.Flags().BoolVarP(&_cloudsearchDefineSuggester, "define-suggester", "", false, "Define Suggester")
	_cloudsearchCmd.Flags().BoolVarP(&_cloudsearchDeleteAnalysisScheme, "delete-analysis-scheme", "", false, "Delete Analysis Scheme")
	_cloudsearchCmd.Flags().BoolVarP(&_cloudsearchDeleteDomain, "delete-domain", "", false, "Delete Domain")
	_cloudsearchCmd.Flags().BoolVarP(&_cloudsearchDeleteExpression, "delete-expression", "", false, "Delete Expression")
	_cloudsearchCmd.Flags().BoolVarP(&_cloudsearchDeleteIndexField, "delete-index-field", "", false, "Delete Index Field")
	_cloudsearchCmd.Flags().BoolVarP(&_cloudsearchDeleteSuggester, "delete-suggester", "", false, "Delete Suggester")
	_cloudsearchCmd.Flags().BoolVarP(&_cloudsearchDescribeAnalysisSchemes, "describe-analysis-schemes", "", false, "Describe Analysis Schemes")
	_cloudsearchCmd.Flags().BoolVarP(&_cloudsearchDescribeAvailabilityOptions, "describe-availability-options", "", false, "Describe Availability Options")
	_cloudsearchCmd.Flags().BoolVarP(&_cloudsearchDescribeDomainEndpointOptions, "describe-domain-endpoint-options", "", false, "Describe Domain Endpoint Options")
	_cloudsearchCmd.Flags().BoolVarP(&_cloudsearchDescribeDomains, "describe-domains", "", false, "Describe Domains")
	_cloudsearchCmd.Flags().BoolVarP(&_cloudsearchDescribeExpressions, "describe-expressions", "", false, "Describe Expressions")
	_cloudsearchCmd.Flags().BoolVarP(&_cloudsearchDescribeIndexFields, "describe-index-fields", "", false, "Describe Index Fields")
	_cloudsearchCmd.Flags().BoolVarP(&_cloudsearchDescribeScalingParameters, "describe-scaling-parameters", "", false, "Describe Scaling Parameters")
	_cloudsearchCmd.Flags().BoolVarP(&_cloudsearchDescribeServiceAccessPolicies, "describe-service-access-policies", "", false, "Describe Service Access Policies")
	_cloudsearchCmd.Flags().BoolVarP(&_cloudsearchDescribeSuggesters, "describe-suggesters", "", false, "Describe Suggesters")
	_cloudsearchCmd.Flags().BoolVarP(&_cloudsearchIndexDocuments, "index-documents", "", false, "Index Documents")
	_cloudsearchCmd.Flags().BoolVarP(&_cloudsearchListDomainNames, "list-domain-names", "", false, "List Domain Names")
	_cloudsearchCmd.Flags().BoolVarP(&_cloudsearchUpdateAvailabilityOptions, "update-availability-options", "", false, "Update Availability Options")
	_cloudsearchCmd.Flags().BoolVarP(&_cloudsearchUpdateDomainEndpointOptions, "update-domain-endpoint-options", "", false, "Update Domain Endpoint Options")
	_cloudsearchCmd.Flags().BoolVarP(&_cloudsearchUpdateScalingParameters, "update-scaling-parameters", "", false, "Update Scaling Parameters")
	_cloudsearchCmd.Flags().BoolVarP(&_cloudsearchUpdateServiceAccessPolicies, "update-service-access-policies", "", false, "Update Service Access Policies")

}
