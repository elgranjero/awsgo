package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// servicequotasCmd represents the servicequotas command
var _servicequotasCmd = &cobra.Command{
	Use:   "servicequotas",
	Short: "AWS servicequotas CLI",
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
		client := servicequotas.NewFromConfig(cfg)
		if _servicequotasAssociateServiceQuotaTemplate {
			servicequotas_AssociateServiceQuotaTemplate(cfg, client)
			return
		}
		if _servicequotasCreateSupportCase {
			servicequotas_CreateSupportCase(cfg, client)
			return
		}
		if _servicequotasDeleteServiceQuotaIncreaseRequestFromTemplate {
			servicequotas_DeleteServiceQuotaIncreaseRequestFromTemplate(cfg, client)
			return
		}
		if _servicequotasDisassociateServiceQuotaTemplate {
			servicequotas_DisassociateServiceQuotaTemplate(cfg, client)
			return
		}
		if _servicequotasGetAssociationForServiceQuotaTemplate {
			servicequotas_GetAssociationForServiceQuotaTemplate(cfg, client)
			return
		}
		if _servicequotasGetAutoManagementConfiguration {
			servicequotas_GetAutoManagementConfiguration(cfg, client)
			return
		}
		if _servicequotasGetAWSDefaultServiceQuota {
			servicequotas_GetAWSDefaultServiceQuota(cfg, client)
			return
		}
		if _servicequotasGetQuotaUtilizationReport {
			servicequotas_GetQuotaUtilizationReport(cfg, client)
			return
		}
		if _servicequotasGetRequestedServiceQuotaChange {
			servicequotas_GetRequestedServiceQuotaChange(cfg, client)
			return
		}
		if _servicequotasGetServiceQuota {
			servicequotas_GetServiceQuota(cfg, client)
			return
		}
		if _servicequotasGetServiceQuotaIncreaseRequestFromTemplate {
			servicequotas_GetServiceQuotaIncreaseRequestFromTemplate(cfg, client)
			return
		}
		if _servicequotasListAWSDefaultServiceQuotas {
			servicequotas_ListAWSDefaultServiceQuotas(cfg, client)
			return
		}
		if _servicequotasListRequestedServiceQuotaChangeHistory {
			servicequotas_ListRequestedServiceQuotaChangeHistory(cfg, client)
			return
		}
		if _servicequotasListRequestedServiceQuotaChangeHistoryByQuota {
			servicequotas_ListRequestedServiceQuotaChangeHistoryByQuota(cfg, client)
			return
		}
		if _servicequotasListServiceQuotaIncreaseRequestsInTemplate {
			servicequotas_ListServiceQuotaIncreaseRequestsInTemplate(cfg, client)
			return
		}
		if _servicequotasListServiceQuotas {
			servicequotas_ListServiceQuotas(cfg, client)
			return
		}
		if _servicequotasListServices {
			servicequotas_ListServices(cfg, client)
			return
		}
		if _servicequotasListTagsForResource {
			servicequotas_ListTagsForResource(cfg, client)
			return
		}
		if _servicequotasPutServiceQuotaIncreaseRequestIntoTemplate {
			servicequotas_PutServiceQuotaIncreaseRequestIntoTemplate(cfg, client)
			return
		}
		if _servicequotasRequestServiceQuotaIncrease {
			servicequotas_RequestServiceQuotaIncrease(cfg, client)
			return
		}
		if _servicequotasStartAutoManagement {
			servicequotas_StartAutoManagement(cfg, client)
			return
		}
		if _servicequotasStartQuotaUtilizationReport {
			servicequotas_StartQuotaUtilizationReport(cfg, client)
			return
		}
		if _servicequotasStopAutoManagement {
			servicequotas_StopAutoManagement(cfg, client)
			return
		}
		if _servicequotasTagResource {
			servicequotas_TagResource(cfg, client)
			return
		}
		if _servicequotasUntagResource {
			servicequotas_UntagResource(cfg, client)
			return
		}
		if _servicequotasUpdateAutoManagement {
			servicequotas_UpdateAutoManagement(cfg, client)
			return
		}

	},
}

var (
	_servicequotasAssociateServiceQuotaTemplate                 bool
	_servicequotasCreateSupportCase                             bool
	_servicequotasDeleteServiceQuotaIncreaseRequestFromTemplate bool
	_servicequotasDisassociateServiceQuotaTemplate              bool
	_servicequotasGetAssociationForServiceQuotaTemplate         bool
	_servicequotasGetAutoManagementConfiguration                bool
	_servicequotasGetAWSDefaultServiceQuota                     bool
	_servicequotasGetQuotaUtilizationReport                     bool
	_servicequotasGetRequestedServiceQuotaChange                bool
	_servicequotasGetServiceQuota                               bool
	_servicequotasGetServiceQuotaIncreaseRequestFromTemplate    bool
	_servicequotasListAWSDefaultServiceQuotas                   bool
	_servicequotasListRequestedServiceQuotaChangeHistory        bool
	_servicequotasListRequestedServiceQuotaChangeHistoryByQuota bool
	_servicequotasListServiceQuotaIncreaseRequestsInTemplate    bool
	_servicequotasListServiceQuotas                             bool
	_servicequotasListServices                                  bool
	_servicequotasListTagsForResource                           bool
	_servicequotasPutServiceQuotaIncreaseRequestIntoTemplate    bool
	_servicequotasRequestServiceQuotaIncrease                   bool
	_servicequotasStartAutoManagement                           bool
	_servicequotasStartQuotaUtilizationReport                   bool
	_servicequotasStopAutoManagement                            bool
	_servicequotasTagResource                                   bool
	_servicequotasUntagResource                                 bool
	_servicequotasUpdateAutoManagement                          bool

	_servicequotasAwsRegion             string
	_servicequotasContextId             string
	_servicequotasDesiredValue          string
	_servicequotasExclusionList         string
	_servicequotasMaxResults            string
	_servicequotasNextToken             string
	_servicequotasNotificationArn       string
	_servicequotasOptInLevel            string
	_servicequotasOptInType             string
	_servicequotasQuotaAppliedAtLevel   string
	_servicequotasQuotaCode             string
	_servicequotasQuotaRequestedAtLevel string
	_servicequotasReportId              string
	_servicequotasRequestId             string
	_servicequotasResourceARN           string
	_servicequotasServiceCode           string
	_servicequotasStatus                string
	_servicequotasSupportCaseAllowed    string
	_servicequotasTagKeys               []string
	_servicequotasTags                  string
)

// Associates your quota request template with your organization. When a new
// Amazon Web Services account is created in your organization, the quota increase
// requests in the template are automatically applied to the account. You can add a
// quota increase request for any adjustable quota to your template.
func servicequotas_AssociateServiceQuotaTemplate(cfg aws.Config, client *servicequotas.Client) {
	input := &servicequotas.AssociateServiceQuotaTemplateInput{}

	if resp, err := client.AssociateServiceQuotaTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Support case for an existing quota increase request. This call only
// creates a Support case if the request has a Pending status.
func servicequotas_CreateSupportCase(cfg aws.Config, client *servicequotas.Client) {
	input := &servicequotas.CreateSupportCaseInput{
		// RequestId: *string, // Required
	}

	if len(_servicequotasRequestId) > 0 {
		input.RequestId = aws.String(_servicequotasRequestId)
	}

	if resp, err := client.CreateSupportCase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the quota increase request for the specified quota from your quota
// request template.
func servicequotas_DeleteServiceQuotaIncreaseRequestFromTemplate(cfg aws.Config, client *servicequotas.Client) {
	input := &servicequotas.DeleteServiceQuotaIncreaseRequestFromTemplateInput{
		// AwsRegion: *string, // Required
		// QuotaCode: *string, // Required
		// ServiceCode: *string, // Required
	}

	if len(_servicequotasAwsRegion) > 0 {
		input.AwsRegion = aws.String(_servicequotasAwsRegion)
	}
	if len(_servicequotasQuotaCode) > 0 {
		input.QuotaCode = aws.String(_servicequotasQuotaCode)
	}
	if len(_servicequotasServiceCode) > 0 {
		input.ServiceCode = aws.String(_servicequotasServiceCode)
	}

	if resp, err := client.DeleteServiceQuotaIncreaseRequestFromTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables your quota request template. After a template is disabled, the quota
// increase requests in the template are not applied to new Amazon Web Services
// accounts in your organization. Disabling a quota request template does not apply
// its quota increase requests.
func servicequotas_DisassociateServiceQuotaTemplate(cfg aws.Config, client *servicequotas.Client) {
	input := &servicequotas.DisassociateServiceQuotaTemplateInput{}

	if resp, err := client.DisassociateServiceQuotaTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the status of the association for the quota request template.
func servicequotas_GetAssociationForServiceQuotaTemplate(cfg aws.Config, client *servicequotas.Client) {
	input := &servicequotas.GetAssociationForServiceQuotaTemplateInput{}

	if resp, err := client.GetAssociationForServiceQuotaTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about your [Service Quotas Automatic Management] configuration. Automatic Management monitors
// your Service Quotas utilization and notifies you before you run out of your
// allocated quotas.
//
// [Service Quotas Automatic Management]: https://docs.aws.amazon.com/servicequotas/latest/userguide/automatic-management.html
func servicequotas_GetAutoManagementConfiguration(cfg aws.Config, client *servicequotas.Client) {
	input := &servicequotas.GetAutoManagementConfigurationInput{}

	if resp, err := client.GetAutoManagementConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the default value for the specified quota. The default value does not
// reflect any quota increases.
func servicequotas_GetAWSDefaultServiceQuota(cfg aws.Config, client *servicequotas.Client) {
	input := &servicequotas.GetAWSDefaultServiceQuotaInput{
		// QuotaCode: *string, // Required
		// ServiceCode: *string, // Required
	}

	if len(_servicequotasQuotaCode) > 0 {
		input.QuotaCode = aws.String(_servicequotasQuotaCode)
	}
	if len(_servicequotasServiceCode) > 0 {
		input.ServiceCode = aws.String(_servicequotasServiceCode)
	}

	if resp, err := client.GetAWSDefaultServiceQuota(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the quota utilization report for your Amazon Web Services account.
// This operation returns paginated results showing your quota usage across all
// Amazon Web Services services, sorted by utilization percentage in descending
// order (highest utilization first).
//
// You must first initiate a report using the StartQuotaUtilizationReport
// operation. The report generation process is asynchronous and may take several
// seconds to complete. Poll this operation periodically to check the status and
// retrieve results when the report is ready.
//
// Each report contains up to 1,000 quota records per page. Use the NextToken
// parameter to retrieve additional pages of results. Reports are automatically
// deleted after 15 minutes.
func servicequotas_GetQuotaUtilizationReport(cfg aws.Config, client *servicequotas.Client) {
	input := &servicequotas.GetQuotaUtilizationReportInput{
		// ReportId: *string, // Required
	}

	if len(_servicequotasReportId) > 0 {
		input.ReportId = aws.String(_servicequotasReportId)
	}
	if len(_servicequotasMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _servicequotasMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_servicequotasNextToken) > 0 {
		input.NextToken = aws.String(_servicequotasNextToken)
	}

	if resp, err := client.GetQuotaUtilizationReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the specified quota increase request.
func servicequotas_GetRequestedServiceQuotaChange(cfg aws.Config, client *servicequotas.Client) {
	input := &servicequotas.GetRequestedServiceQuotaChangeInput{
		// RequestId: *string, // Required
	}

	if len(_servicequotasRequestId) > 0 {
		input.RequestId = aws.String(_servicequotasRequestId)
	}

	if resp, err := client.GetRequestedServiceQuotaChange(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the applied quota value for the specified account-level or
// resource-level quota. For some quotas, only the default values are available. If
// the applied quota value is not available for a quota, the quota is not
// retrieved.
func servicequotas_GetServiceQuota(cfg aws.Config, client *servicequotas.Client) {
	input := &servicequotas.GetServiceQuotaInput{
		// QuotaCode: *string, // Required
		// ServiceCode: *string, // Required
	}

	if len(_servicequotasQuotaCode) > 0 {
		input.QuotaCode = aws.String(_servicequotasQuotaCode)
	}
	if len(_servicequotasServiceCode) > 0 {
		input.ServiceCode = aws.String(_servicequotasServiceCode)
	}
	if len(_servicequotasContextId) > 0 {
		input.ContextId = aws.String(_servicequotasContextId)
	}

	if resp, err := client.GetServiceQuota(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the specified quota increase request in your quota
// request template.
func servicequotas_GetServiceQuotaIncreaseRequestFromTemplate(cfg aws.Config, client *servicequotas.Client) {
	input := &servicequotas.GetServiceQuotaIncreaseRequestFromTemplateInput{
		// AwsRegion: *string, // Required
		// QuotaCode: *string, // Required
		// ServiceCode: *string, // Required
	}

	if len(_servicequotasAwsRegion) > 0 {
		input.AwsRegion = aws.String(_servicequotasAwsRegion)
	}
	if len(_servicequotasQuotaCode) > 0 {
		input.QuotaCode = aws.String(_servicequotasQuotaCode)
	}
	if len(_servicequotasServiceCode) > 0 {
		input.ServiceCode = aws.String(_servicequotasServiceCode)
	}

	if resp, err := client.GetServiceQuotaIncreaseRequestFromTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the default values for the quotas for the specified Amazon Web Services
// service. A default value does not reflect any quota increases.
func servicequotas_ListAWSDefaultServiceQuotas(cfg aws.Config, client *servicequotas.Client) {
	input := &servicequotas.ListAWSDefaultServiceQuotasInput{
		// ServiceCode: *string, // Required
	}

	if len(_servicequotasServiceCode) > 0 {
		input.ServiceCode = aws.String(_servicequotasServiceCode)
	}
	if len(_servicequotasMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _servicequotasMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_servicequotasNextToken) > 0 {
		input.NextToken = aws.String(_servicequotasNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAWSDefaultServiceQuotas(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicequotas.ListAWSDefaultServiceQuotasOutput
	p := servicequotas.NewListAWSDefaultServiceQuotasPaginator(client, input)
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

// Retrieves the quota increase requests for the specified Amazon Web Services
// service. Filter responses to return quota requests at either the account level,
// resource level, or all levels. Responses include any open or closed requests
// within 90 days.
func servicequotas_ListRequestedServiceQuotaChangeHistory(cfg aws.Config, client *servicequotas.Client) {
	input := &servicequotas.ListRequestedServiceQuotaChangeHistoryInput{}

	if len(_servicequotasMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _servicequotasMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_servicequotasNextToken) > 0 {
		input.NextToken = aws.String(_servicequotasNextToken)
	}
	if len(_servicequotasQuotaRequestedAtLevel) > 0 {
		if err := assignInputField(input, "QuotaRequestedAtLevel", _servicequotasQuotaRequestedAtLevel); err != nil {
			log.Errorf("invalid --quota-requested-at-level: %s", err.Error())
			return
		}
	}
	if len(_servicequotasServiceCode) > 0 {
		input.ServiceCode = aws.String(_servicequotasServiceCode)
	}
	if len(_servicequotasStatus) > 0 {
		if err := assignInputField(input, "Status", _servicequotasStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListRequestedServiceQuotaChangeHistory(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicequotas.ListRequestedServiceQuotaChangeHistoryOutput
	p := servicequotas.NewListRequestedServiceQuotaChangeHistoryPaginator(client, input)
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

// Retrieves the quota increase requests for the specified quota. Filter responses
// to return quota requests at either the account level, resource level, or all
// levels.
func servicequotas_ListRequestedServiceQuotaChangeHistoryByQuota(cfg aws.Config, client *servicequotas.Client) {
	input := &servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaInput{
		// QuotaCode: *string, // Required
		// ServiceCode: *string, // Required
	}

	if len(_servicequotasQuotaCode) > 0 {
		input.QuotaCode = aws.String(_servicequotasQuotaCode)
	}
	if len(_servicequotasServiceCode) > 0 {
		input.ServiceCode = aws.String(_servicequotasServiceCode)
	}
	if len(_servicequotasMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _servicequotasMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_servicequotasNextToken) > 0 {
		input.NextToken = aws.String(_servicequotasNextToken)
	}
	if len(_servicequotasQuotaRequestedAtLevel) > 0 {
		if err := assignInputField(input, "QuotaRequestedAtLevel", _servicequotasQuotaRequestedAtLevel); err != nil {
			log.Errorf("invalid --quota-requested-at-level: %s", err.Error())
			return
		}
	}
	if len(_servicequotasStatus) > 0 {
		if err := assignInputField(input, "Status", _servicequotasStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListRequestedServiceQuotaChangeHistoryByQuota(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput
	p := servicequotas.NewListRequestedServiceQuotaChangeHistoryByQuotaPaginator(client, input)
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

// Lists the quota increase requests in the specified quota request template.
func servicequotas_ListServiceQuotaIncreaseRequestsInTemplate(cfg aws.Config, client *servicequotas.Client) {
	input := &servicequotas.ListServiceQuotaIncreaseRequestsInTemplateInput{}

	if len(_servicequotasAwsRegion) > 0 {
		input.AwsRegion = aws.String(_servicequotasAwsRegion)
	}
	if len(_servicequotasMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _servicequotasMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_servicequotasNextToken) > 0 {
		input.NextToken = aws.String(_servicequotasNextToken)
	}
	if len(_servicequotasServiceCode) > 0 {
		input.ServiceCode = aws.String(_servicequotasServiceCode)
	}

	if disablePaginator() {
		if resp, err := client.ListServiceQuotaIncreaseRequestsInTemplate(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicequotas.ListServiceQuotaIncreaseRequestsInTemplateOutput
	p := servicequotas.NewListServiceQuotaIncreaseRequestsInTemplatePaginator(client, input)
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

// Lists the applied quota values for the specified Amazon Web Services service.
// For some quotas, only the default values are available. If the applied quota
// value is not available for a quota, the quota is not retrieved. Filter responses
// to return applied quota values at either the account level, resource level, or
// all levels.
func servicequotas_ListServiceQuotas(cfg aws.Config, client *servicequotas.Client) {
	input := &servicequotas.ListServiceQuotasInput{
		// ServiceCode: *string, // Required
	}

	if len(_servicequotasServiceCode) > 0 {
		input.ServiceCode = aws.String(_servicequotasServiceCode)
	}
	if len(_servicequotasMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _servicequotasMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_servicequotasNextToken) > 0 {
		input.NextToken = aws.String(_servicequotasNextToken)
	}
	if len(_servicequotasQuotaAppliedAtLevel) > 0 {
		if err := assignInputField(input, "QuotaAppliedAtLevel", _servicequotasQuotaAppliedAtLevel); err != nil {
			log.Errorf("invalid --quota-applied-at-level: %s", err.Error())
			return
		}
	}
	if len(_servicequotasQuotaCode) > 0 {
		input.QuotaCode = aws.String(_servicequotasQuotaCode)
	}

	if disablePaginator() {
		if resp, err := client.ListServiceQuotas(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicequotas.ListServiceQuotasOutput
	p := servicequotas.NewListServiceQuotasPaginator(client, input)
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

// Lists the names and codes for the Amazon Web Services services integrated with
// Service Quotas.
func servicequotas_ListServices(cfg aws.Config, client *servicequotas.Client) {
	input := &servicequotas.ListServicesInput{}

	if len(_servicequotasMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _servicequotasMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_servicequotasNextToken) > 0 {
		input.NextToken = aws.String(_servicequotasNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListServices(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicequotas.ListServicesOutput
	p := servicequotas.NewListServicesPaginator(client, input)
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

// Returns a list of the tags assigned to the specified applied quota.
func servicequotas_ListTagsForResource(cfg aws.Config, client *servicequotas.Client) {
	input := &servicequotas.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_servicequotasResourceARN) > 0 {
		input.ResourceARN = aws.String(_servicequotasResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a quota increase request to your quota request template.
func servicequotas_PutServiceQuotaIncreaseRequestIntoTemplate(cfg aws.Config, client *servicequotas.Client) {
	input := &servicequotas.PutServiceQuotaIncreaseRequestIntoTemplateInput{
		// AwsRegion: *string, // Required
		// DesiredValue: *float64, // Required
		// QuotaCode: *string, // Required
		// ServiceCode: *string, // Required
	}

	if len(_servicequotasAwsRegion) > 0 {
		input.AwsRegion = aws.String(_servicequotasAwsRegion)
	}
	if len(_servicequotasDesiredValue) > 0 {
		if err := assignInputField(input, "DesiredValue", _servicequotasDesiredValue); err != nil {
			log.Errorf("invalid --desired-value: %s", err.Error())
			return
		}
	}
	if len(_servicequotasQuotaCode) > 0 {
		input.QuotaCode = aws.String(_servicequotasQuotaCode)
	}
	if len(_servicequotasServiceCode) > 0 {
		input.ServiceCode = aws.String(_servicequotasServiceCode)
	}

	if resp, err := client.PutServiceQuotaIncreaseRequestIntoTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Submits a quota increase request for the specified quota at the account or
// resource level.
func servicequotas_RequestServiceQuotaIncrease(cfg aws.Config, client *servicequotas.Client) {
	input := &servicequotas.RequestServiceQuotaIncreaseInput{
		// DesiredValue: *float64, // Required
		// QuotaCode: *string, // Required
		// ServiceCode: *string, // Required
	}

	if len(_servicequotasDesiredValue) > 0 {
		if err := assignInputField(input, "DesiredValue", _servicequotasDesiredValue); err != nil {
			log.Errorf("invalid --desired-value: %s", err.Error())
			return
		}
	}
	if len(_servicequotasQuotaCode) > 0 {
		input.QuotaCode = aws.String(_servicequotasQuotaCode)
	}
	if len(_servicequotasServiceCode) > 0 {
		input.ServiceCode = aws.String(_servicequotasServiceCode)
	}
	if len(_servicequotasContextId) > 0 {
		input.ContextId = aws.String(_servicequotasContextId)
	}
	if len(_servicequotasSupportCaseAllowed) > 0 {
		if err := assignInputField(input, "SupportCaseAllowed", _servicequotasSupportCaseAllowed); err != nil {
			log.Errorf("invalid --support-case-allowed: %s", err.Error())
			return
		}
	}

	if resp, err := client.RequestServiceQuotaIncrease(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts [Service Quotas Automatic Management] for an Amazon Web Services account, including notification preferences
// and excluded quotas configurations. Automatic Management monitors your Service
// Quotas utilization and notifies you before you run out of your allocated quotas.
//
// [Service Quotas Automatic Management]: https://docs.aws.amazon.com/servicequotas/latest/userguide/automatic-management.html
func servicequotas_StartAutoManagement(cfg aws.Config, client *servicequotas.Client) {
	input := &servicequotas.StartAutoManagementInput{
		// OptInLevel: types.OptInLevel, // Required
		// OptInType: types.OptInType, // Required
	}

	if len(_servicequotasOptInLevel) > 0 {
		if err := assignInputField(input, "OptInLevel", _servicequotasOptInLevel); err != nil {
			log.Errorf("invalid --opt-in-level: %s", err.Error())
			return
		}
	}
	if len(_servicequotasOptInType) > 0 {
		if err := assignInputField(input, "OptInType", _servicequotasOptInType); err != nil {
			log.Errorf("invalid --opt-in-type: %s", err.Error())
			return
		}
	}
	if len(_servicequotasExclusionList) > 0 {
		if err := assignInputField(input, "ExclusionList", _servicequotasExclusionList); err != nil {
			log.Errorf("invalid --exclusion-list: %s", err.Error())
			return
		}
	}
	if len(_servicequotasNotificationArn) > 0 {
		input.NotificationArn = aws.String(_servicequotasNotificationArn)
	}

	if resp, err := client.StartAutoManagement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates the generation of a quota utilization report for your Amazon Web
// Services account. This asynchronous operation analyzes your quota usage across
// all Amazon Web Services services and returns a unique report identifier that you
// can use to retrieve the results.
//
// The report generation process may take several seconds to complete, depending
// on the number of quotas in your account. Use the GetQuotaUtilizationReport
// operation to check the status and retrieve the results when the report is ready.
func servicequotas_StartQuotaUtilizationReport(cfg aws.Config, client *servicequotas.Client) {
	input := &servicequotas.StartQuotaUtilizationReportInput{}

	if resp, err := client.StartQuotaUtilizationReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops [Service Quotas Automatic Management] for an Amazon Web Services account and removes all associated
// configurations. Automatic Management monitors your Service Quotas utilization
// and notifies you before you run out of your allocated quotas.
//
// [Service Quotas Automatic Management]: https://docs.aws.amazon.com/servicequotas/latest/userguide/automatic-management.html
func servicequotas_StopAutoManagement(cfg aws.Config, client *servicequotas.Client) {
	input := &servicequotas.StopAutoManagementInput{}

	if resp, err := client.StopAutoManagement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds tags to the specified applied quota. You can include one or more tags to
// add to the quota.
func servicequotas_TagResource(cfg aws.Config, client *servicequotas.Client) {
	input := &servicequotas.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_servicequotasResourceARN) > 0 {
		input.ResourceARN = aws.String(_servicequotasResourceARN)
	}
	if len(_servicequotasTags) > 0 {
		if err := assignInputField(input, "Tags", _servicequotasTags); err != nil {
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

// Removes tags from the specified applied quota. You can specify one or more tags
// to remove.
func servicequotas_UntagResource(cfg aws.Config, client *servicequotas.Client) {
	input := &servicequotas.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_servicequotasResourceARN) > 0 {
		input.ResourceARN = aws.String(_servicequotasResourceARN)
	}
	if len(_servicequotasTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _servicequotasTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates your [Service Quotas Automatic Management] configuration, including notification preferences and excluded
// quotas. Automatic Management monitors your Service Quotas utilization and
// notifies you before you run out of your allocated quotas.
//
// [Service Quotas Automatic Management]: https://docs.aws.amazon.com/servicequotas/latest/userguide/automatic-management.html
func servicequotas_UpdateAutoManagement(cfg aws.Config, client *servicequotas.Client) {
	input := &servicequotas.UpdateAutoManagementInput{}

	if len(_servicequotasExclusionList) > 0 {
		if err := assignInputField(input, "ExclusionList", _servicequotasExclusionList); err != nil {
			log.Errorf("invalid --exclusion-list: %s", err.Error())
			return
		}
	}
	if len(_servicequotasNotificationArn) > 0 {
		input.NotificationArn = aws.String(_servicequotasNotificationArn)
	}
	if len(_servicequotasOptInType) > 0 {
		if err := assignInputField(input, "OptInType", _servicequotasOptInType); err != nil {
			log.Errorf("invalid --opt-in-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAutoManagement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_servicequotasCmd)
	_servicequotasCmd.Flags().SortFlags = false

	_servicequotasCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_servicequotasCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_servicequotasCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_servicequotasCmd.Flags().StringVarP(&_servicequotasAwsRegion, "aws-region", "", "", "AWS Region")
	_servicequotasCmd.Flags().StringVarP(&_servicequotasContextId, "context-id", "", "", "Context ID")
	_servicequotasCmd.Flags().StringVarP(&_servicequotasDesiredValue, "desired-value", "", "", "Desired Value")
	_servicequotasCmd.Flags().StringVarP(&_servicequotasExclusionList, "exclusion-list", "", "", "Exclusion List")
	_servicequotasCmd.Flags().StringVarP(&_servicequotasMaxResults, "max-results", "", "", "Max Results")
	_servicequotasCmd.Flags().StringVarP(&_servicequotasNextToken, "next-token", "", "", "Next Token")
	_servicequotasCmd.Flags().StringVarP(&_servicequotasNotificationArn, "notification-arn", "", "", "Notification ARN")
	_servicequotasCmd.Flags().StringVarP(&_servicequotasOptInLevel, "opt-in-level", "", "", "Opt In Level")
	_servicequotasCmd.Flags().StringVarP(&_servicequotasOptInType, "opt-in-type", "", "", "Opt In Type")
	_servicequotasCmd.Flags().StringVarP(&_servicequotasQuotaAppliedAtLevel, "quota-applied-at-level", "", "", "Quota Applied At Level")
	_servicequotasCmd.Flags().StringVarP(&_servicequotasQuotaCode, "quota-code", "", "", "Quota Code")
	_servicequotasCmd.Flags().StringVarP(&_servicequotasQuotaRequestedAtLevel, "quota-requested-at-level", "", "", "Quota Requested At Level")
	_servicequotasCmd.Flags().StringVarP(&_servicequotasReportId, "report-id", "", "", "Report ID")
	_servicequotasCmd.Flags().StringVarP(&_servicequotasRequestId, "request-id", "", "", "Request ID")
	_servicequotasCmd.Flags().StringVarP(&_servicequotasResourceARN, "resource-arn", "", "", "Resource ARN")
	_servicequotasCmd.Flags().StringVarP(&_servicequotasServiceCode, "service-code", "", "", "Service Code")
	_servicequotasCmd.Flags().StringVarP(&_servicequotasStatus, "status", "", "", "Status")
	_servicequotasCmd.Flags().StringVarP(&_servicequotasSupportCaseAllowed, "support-case-allowed", "", "", "Support Case Allowed")
	_servicequotasCmd.Flags().StringSliceVarP(&_servicequotasTagKeys, "tag-keys", "", nil, "Tag Keys")
	_servicequotasCmd.Flags().StringVarP(&_servicequotasTags, "tags", "", "", "Tags")

	_servicequotasCmd.Flags().BoolVarP(&_servicequotasAssociateServiceQuotaTemplate, "associate-service-quota-template", "", false, "Associate Service Quota Template")
	_servicequotasCmd.Flags().BoolVarP(&_servicequotasCreateSupportCase, "create-support-case", "", false, "Create Support Case")
	_servicequotasCmd.Flags().BoolVarP(&_servicequotasDeleteServiceQuotaIncreaseRequestFromTemplate, "delete-service-quota-increase-request-from-template", "", false, "Delete Service Quota Increase Request From Template")
	_servicequotasCmd.Flags().BoolVarP(&_servicequotasDisassociateServiceQuotaTemplate, "disassociate-service-quota-template", "", false, "Disassociate Service Quota Template")
	_servicequotasCmd.Flags().BoolVarP(&_servicequotasGetAssociationForServiceQuotaTemplate, "get-association-for-service-quota-template", "", false, "Get Association For Service Quota Template")
	_servicequotasCmd.Flags().BoolVarP(&_servicequotasGetAutoManagementConfiguration, "get-auto-management-configuration", "", false, "Get Auto Management Configuration")
	_servicequotasCmd.Flags().BoolVarP(&_servicequotasGetAWSDefaultServiceQuota, "get-aws-default-service-quota", "", false, "Get AWS Default Service Quota")
	_servicequotasCmd.Flags().BoolVarP(&_servicequotasGetQuotaUtilizationReport, "get-quota-utilization-report", "", false, "Get Quota Utilization Report")
	_servicequotasCmd.Flags().BoolVarP(&_servicequotasGetRequestedServiceQuotaChange, "get-requested-service-quota-change", "", false, "Get Requested Service Quota Change")
	_servicequotasCmd.Flags().BoolVarP(&_servicequotasGetServiceQuota, "get-service-quota", "", false, "Get Service Quota")
	_servicequotasCmd.Flags().BoolVarP(&_servicequotasGetServiceQuotaIncreaseRequestFromTemplate, "get-service-quota-increase-request-from-template", "", false, "Get Service Quota Increase Request From Template")
	_servicequotasCmd.Flags().BoolVarP(&_servicequotasListAWSDefaultServiceQuotas, "list-aws-default-service-quotas", "", false, "List AWS Default Service Quotas")
	_servicequotasCmd.Flags().BoolVarP(&_servicequotasListRequestedServiceQuotaChangeHistory, "list-requested-service-quota-change-history", "", false, "List Requested Service Quota Change History")
	_servicequotasCmd.Flags().BoolVarP(&_servicequotasListRequestedServiceQuotaChangeHistoryByQuota, "list-requested-service-quota-change-history-by-quota", "", false, "List Requested Service Quota Change History By Quota")
	_servicequotasCmd.Flags().BoolVarP(&_servicequotasListServiceQuotaIncreaseRequestsInTemplate, "list-service-quota-increase-requests-in-template", "", false, "List Service Quota Increase Requests In Template")
	_servicequotasCmd.Flags().BoolVarP(&_servicequotasListServiceQuotas, "list-service-quotas", "", false, "List Service Quotas")
	_servicequotasCmd.Flags().BoolVarP(&_servicequotasListServices, "list-services", "", false, "List Services")
	_servicequotasCmd.Flags().BoolVarP(&_servicequotasListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_servicequotasCmd.Flags().BoolVarP(&_servicequotasPutServiceQuotaIncreaseRequestIntoTemplate, "put-service-quota-increase-request-into-template", "", false, "Put Service Quota Increase Request Into Template")
	_servicequotasCmd.Flags().BoolVarP(&_servicequotasRequestServiceQuotaIncrease, "request-service-quota-increase", "", false, "Request Service Quota Increase")
	_servicequotasCmd.Flags().BoolVarP(&_servicequotasStartAutoManagement, "start-auto-management", "", false, "Start Auto Management")
	_servicequotasCmd.Flags().BoolVarP(&_servicequotasStartQuotaUtilizationReport, "start-quota-utilization-report", "", false, "Start Quota Utilization Report")
	_servicequotasCmd.Flags().BoolVarP(&_servicequotasStopAutoManagement, "stop-auto-management", "", false, "Stop Auto Management")
	_servicequotasCmd.Flags().BoolVarP(&_servicequotasTagResource, "tag-resource", "", false, "Tag Resource")
	_servicequotasCmd.Flags().BoolVarP(&_servicequotasUntagResource, "untag-resource", "", false, "Untag Resource")
	_servicequotasCmd.Flags().BoolVarP(&_servicequotasUpdateAutoManagement, "update-auto-management", "", false, "Update Auto Management")

}
