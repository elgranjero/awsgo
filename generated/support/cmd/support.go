package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/support"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// supportCmd represents the support command
var _supportCmd = &cobra.Command{
	Use:   "support",
	Short: "AWS support CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := support.NewFromConfig(cfg)
		if _supportAddAttachmentsToSet {
			support_AddAttachmentsToSet(cfg, client)
			return
		}
		if _supportAddCommunicationToCase {
			support_AddCommunicationToCase(cfg, client)
			return
		}
		if _supportCreateCase {
			support_CreateCase(cfg, client)
			return
		}
		if _supportDescribeAttachment {
			support_DescribeAttachment(cfg, client)
			return
		}
		if _supportDescribeCases {
			support_DescribeCases(cfg, client)
			return
		}
		if _supportDescribeCommunications {
			support_DescribeCommunications(cfg, client)
			return
		}
		if _supportDescribeCreateCaseOptions {
			support_DescribeCreateCaseOptions(cfg, client)
			return
		}
		if _supportDescribeServices {
			support_DescribeServices(cfg, client)
			return
		}
		if _supportDescribeSeverityLevels {
			support_DescribeSeverityLevels(cfg, client)
			return
		}
		if _supportDescribeSupportedLanguages {
			support_DescribeSupportedLanguages(cfg, client)
			return
		}
		if _supportDescribeTrustedAdvisorCheckRefreshStatuses {
			support_DescribeTrustedAdvisorCheckRefreshStatuses(cfg, client)
			return
		}
		if _supportDescribeTrustedAdvisorCheckResult {
			support_DescribeTrustedAdvisorCheckResult(cfg, client)
			return
		}
		if _supportDescribeTrustedAdvisorCheckSummaries {
			support_DescribeTrustedAdvisorCheckSummaries(cfg, client)
			return
		}
		if _supportDescribeTrustedAdvisorChecks {
			support_DescribeTrustedAdvisorChecks(cfg, client)
			return
		}
		if _supportRefreshTrustedAdvisorCheck {
			support_RefreshTrustedAdvisorCheck(cfg, client)
			return
		}
		if _supportResolveCase {
			support_ResolveCase(cfg, client)
			return
		}

	},
}

var (
	_supportAddAttachmentsToSet                        bool
	_supportAddCommunicationToCase                     bool
	_supportCreateCase                                 bool
	_supportDescribeAttachment                         bool
	_supportDescribeCases                              bool
	_supportDescribeCommunications                     bool
	_supportDescribeCreateCaseOptions                  bool
	_supportDescribeServices                           bool
	_supportDescribeSeverityLevels                     bool
	_supportDescribeSupportedLanguages                 bool
	_supportDescribeTrustedAdvisorCheckRefreshStatuses bool
	_supportDescribeTrustedAdvisorCheckResult          bool
	_supportDescribeTrustedAdvisorCheckSummaries       bool
	_supportDescribeTrustedAdvisorChecks               bool
	_supportRefreshTrustedAdvisorCheck                 bool
	_supportResolveCase                                bool

	_supportAfterTime             string
	_supportAttachmentId          string
	_supportAttachmentSetId       string
	_supportAttachments           string
	_supportBeforeTime            string
	_supportCaseId                string
	_supportCaseIdList            []string
	_supportCategoryCode          string
	_supportCcEmailAddresses      []string
	_supportCheckId               string
	_supportCheckIds              string
	_supportCommunicationBody     string
	_supportDisplayId             string
	_supportIncludeCommunications string
	_supportIncludeResolvedCases  string
	_supportIssueType             string
	_supportLanguage              string
	_supportMaxResults            string
	_supportNextToken             string
	_supportServiceCode           string
	_supportServiceCodeList       []string
	_supportSeverityCode          string
	_supportSubject               string
)

// Adds one or more attachments to an attachment set.
// An attachment set is a temporary container for attachments that you add to a
// case or case communication. The set is available for 1 hour after it's created.
// The expiryTime returned in the response is when the set expires.
//
// - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan to
// use the Amazon Web Services Support API.
//
// - If you call the Amazon Web Services Support API from an account that
// doesn't have a Business, Enterprise On-Ramp, or Enterprise Support plan, the
// SubscriptionRequiredException error message appears. For information about
// changing your support plan, see [Amazon Web Services Support].
//
// [Amazon Web Services Support]: http://aws.amazon.com/premiumsupport/
func support_AddAttachmentsToSet(cfg aws.Config, client *support.Client) {
	input := &support.AddAttachmentsToSetInput{
		// Attachments: []types.Attachment, // Required
	}

	if len(_supportAttachments) > 0 {
		if err := assignInputField(input, "Attachments", _supportAttachments); err != nil {
			log.Errorf("invalid --attachments: %s", err.Error())
			return
		}
	}
	if len(_supportAttachmentSetId) > 0 {
		input.AttachmentSetId = aws.String(_supportAttachmentSetId)
	}

	if resp, err := client.AddAttachmentsToSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds additional customer communication to an Amazon Web Services Support case.
// Use the caseId parameter to identify the case to which to add communication.
// You can list a set of email addresses to copy on the communication by using the
// ccEmailAddresses parameter. The communicationBody value contains the text of
// the communication.
//
// - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan to
// use the Amazon Web Services Support API.
//
// - If you call the Amazon Web Services Support API from an account that
// doesn't have a Business, Enterprise On-Ramp, or Enterprise Support plan, the
// SubscriptionRequiredException error message appears. For information about
// changing your support plan, see [Amazon Web Services Support].
//
// [Amazon Web Services Support]: http://aws.amazon.com/premiumsupport/
func support_AddCommunicationToCase(cfg aws.Config, client *support.Client) {
	input := &support.AddCommunicationToCaseInput{
		// CommunicationBody: *string, // Required
	}

	if len(_supportCommunicationBody) > 0 {
		input.CommunicationBody = aws.String(_supportCommunicationBody)
	}
	if len(_supportAttachmentSetId) > 0 {
		input.AttachmentSetId = aws.String(_supportAttachmentSetId)
	}
	if len(_supportCaseId) > 0 {
		input.CaseId = aws.String(_supportCaseId)
	}
	if len(_supportCcEmailAddresses) > 0 {
		input.CcEmailAddresses = append([]string(nil), _supportCcEmailAddresses...)
	}

	if resp, err := client.AddCommunicationToCase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a case in the Amazon Web Services Support Center. This operation is
// similar to how you create a case in the Amazon Web Services Support Center [Create Case]page.
//
// The Amazon Web Services Support API doesn't support requesting service limit
// increases. You can submit a service limit increase in the following ways:
//
// - Submit a request from the Amazon Web Services Support Center [Create Case]page.
//
// - Use the Service Quotas [RequestServiceQuotaIncrease]operation.
//
// A successful CreateCase request returns an Amazon Web Services Support case
// number. You can use the DescribeCasesoperation and specify the case number to get existing
// Amazon Web Services Support cases. After you create a case, use the AddCommunicationToCaseoperation
// to add additional communication or attachments to an existing case.
//
// The caseId is separate from the displayId that appears in the [Amazon Web Services Support Center]. Use the DescribeCases
// operation to get the displayId .
//
// - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan to
// use the Amazon Web Services Support API.
//
// - If you call the Amazon Web Services Support API from an account that
// doesn't have a Business, Enterprise On-Ramp, or Enterprise Support plan, the
// SubscriptionRequiredException error message appears. For information about
// changing your support plan, see [Amazon Web Services Support].
//
// [Amazon Web Services Support]: http://aws.amazon.com/premiumsupport/
// [Create Case]: https://console.aws.amazon.com/support/home#/case/create
// [RequestServiceQuotaIncrease]: https://docs.aws.amazon.com/servicequotas/2019-06-24/apireference/API_RequestServiceQuotaIncrease.html
// [Amazon Web Services Support Center]: https://console.aws.amazon.com/support
func support_CreateCase(cfg aws.Config, client *support.Client) {
	input := &support.CreateCaseInput{
		// CommunicationBody: *string, // Required
		// Subject: *string, // Required
	}

	if len(_supportCommunicationBody) > 0 {
		input.CommunicationBody = aws.String(_supportCommunicationBody)
	}
	if len(_supportSubject) > 0 {
		input.Subject = aws.String(_supportSubject)
	}
	if len(_supportAttachmentSetId) > 0 {
		input.AttachmentSetId = aws.String(_supportAttachmentSetId)
	}
	if len(_supportCategoryCode) > 0 {
		input.CategoryCode = aws.String(_supportCategoryCode)
	}
	if len(_supportCcEmailAddresses) > 0 {
		input.CcEmailAddresses = append([]string(nil), _supportCcEmailAddresses...)
	}
	if len(_supportIssueType) > 0 {
		input.IssueType = aws.String(_supportIssueType)
	}
	if len(_supportLanguage) > 0 {
		input.Language = aws.String(_supportLanguage)
	}
	if len(_supportServiceCode) > 0 {
		input.ServiceCode = aws.String(_supportServiceCode)
	}
	if len(_supportSeverityCode) > 0 {
		input.SeverityCode = aws.String(_supportSeverityCode)
	}

	if resp, err := client.CreateCase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the attachment that has the specified ID. Attachments can include
// screenshots, error logs, or other files that describe your issue. Attachment IDs
// are generated by the case management system when you add an attachment to a case
// or case communication. Attachment IDs are returned in the AttachmentDetailsobjects that are
// returned by the DescribeCommunicationsoperation.
//
// - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan to
// use the Amazon Web Services Support API.
//
// - If you call the Amazon Web Services Support API from an account that
// doesn't have a Business, Enterprise On-Ramp, or Enterprise Support plan, the
// SubscriptionRequiredException error message appears. For information about
// changing your support plan, see [Amazon Web Services Support].
//
// [Amazon Web Services Support]: http://aws.amazon.com/premiumsupport/
func support_DescribeAttachment(cfg aws.Config, client *support.Client) {
	input := &support.DescribeAttachmentInput{
		// AttachmentId: *string, // Required
	}

	if len(_supportAttachmentId) > 0 {
		input.AttachmentId = aws.String(_supportAttachmentId)
	}

	if resp, err := client.DescribeAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of cases that you specify by passing one or more case IDs. You
// can use the afterTime and beforeTime parameters to filter the cases by date.
// You can set values for the includeResolvedCases and includeCommunications
// parameters to specify how much information to return.
//
// The response returns the following in JSON format:
//
// - One or more [CaseDetails]data types.
//
// - One or more nextToken values, which specify where to paginate the returned
// records represented by the CaseDetails objects.
//
// Case data is available for 12 months after creation. If a case was created more
// than 12 months ago, a request might return an error.
//
// - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan to
// use the Amazon Web Services Support API.
//
// - If you call the Amazon Web Services Support API from an account that
// doesn't have a Business, Enterprise On-Ramp, or Enterprise Support plan, the
// SubscriptionRequiredException error message appears. For information about
// changing your support plan, see [Amazon Web Services Support].
//
// [Amazon Web Services Support]: http://aws.amazon.com/premiumsupport/
// [CaseDetails]: https://docs.aws.amazon.com/awssupport/latest/APIReference/API_CaseDetails.html
func support_DescribeCases(cfg aws.Config, client *support.Client) {
	input := &support.DescribeCasesInput{}

	if len(_supportAfterTime) > 0 {
		input.AfterTime = aws.String(_supportAfterTime)
	}
	if len(_supportBeforeTime) > 0 {
		input.BeforeTime = aws.String(_supportBeforeTime)
	}
	if len(_supportCaseIdList) > 0 {
		input.CaseIdList = append([]string(nil), _supportCaseIdList...)
	}
	if len(_supportDisplayId) > 0 {
		input.DisplayId = aws.String(_supportDisplayId)
	}
	if len(_supportIncludeCommunications) > 0 {
		if err := assignInputField(input, "IncludeCommunications", _supportIncludeCommunications); err != nil {
			log.Errorf("invalid --include-communications: %s", err.Error())
			return
		}
	}
	if len(_supportIncludeResolvedCases) > 0 {
		if err := assignInputField(input, "IncludeResolvedCases", _supportIncludeResolvedCases); err != nil {
			log.Errorf("invalid --include-resolved-cases: %s", err.Error())
			return
		}
	}
	if len(_supportLanguage) > 0 {
		input.Language = aws.String(_supportLanguage)
	}
	if len(_supportMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _supportMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_supportNextToken) > 0 {
		input.NextToken = aws.String(_supportNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeCases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*support.DescribeCasesOutput
	p := support.NewDescribeCasesPaginator(client, input)
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

// Returns communications and attachments for one or more support cases. Use the
// afterTime and beforeTime parameters to filter by date. You can use the caseId
// parameter to restrict the results to a specific case.
//
// Case data is available for 12 months after creation. If a case was created more
// than 12 months ago, a request for data might cause an error.
//
// You can use the maxResults and nextToken parameters to control the pagination
// of the results. Set maxResults to the number of cases that you want to display
// on each page, and use nextToken to specify the resumption of pagination.
//
// - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan to
// use the Amazon Web Services Support API.
//
// - If you call the Amazon Web Services Support API from an account that
// doesn't have a Business, Enterprise On-Ramp, or Enterprise Support plan, the
// SubscriptionRequiredException error message appears. For information about
// changing your support plan, see [Amazon Web Services Support].
//
// [Amazon Web Services Support]: http://aws.amazon.com/premiumsupport/
func support_DescribeCommunications(cfg aws.Config, client *support.Client) {
	input := &support.DescribeCommunicationsInput{
		// CaseId: *string, // Required
	}

	if len(_supportCaseId) > 0 {
		input.CaseId = aws.String(_supportCaseId)
	}
	if len(_supportAfterTime) > 0 {
		input.AfterTime = aws.String(_supportAfterTime)
	}
	if len(_supportBeforeTime) > 0 {
		input.BeforeTime = aws.String(_supportBeforeTime)
	}
	if len(_supportMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _supportMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_supportNextToken) > 0 {
		input.NextToken = aws.String(_supportNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeCommunications(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*support.DescribeCommunicationsOutput
	p := support.NewDescribeCommunicationsPaginator(client, input)
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

// Returns a list of CreateCaseOption types along with the corresponding supported
// hours and language availability. You can specify the language categoryCode ,
// issueType and serviceCode used to retrieve the CreateCaseOptions.
//
// - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan to
// use the Amazon Web Services Support API.
//
// - If you call the Amazon Web Services Support API from an account that
// doesn't have a Business, Enterprise On-Ramp, or Enterprise Support plan, the
// SubscriptionRequiredException error message appears. For information about
// changing your support plan, see [Amazon Web Services Support].
//
// [Amazon Web Services Support]: http://aws.amazon.com/premiumsupport/
func support_DescribeCreateCaseOptions(cfg aws.Config, client *support.Client) {
	input := &support.DescribeCreateCaseOptionsInput{
		// CategoryCode: *string, // Required
		// IssueType: *string, // Required
		// Language: *string, // Required
		// ServiceCode: *string, // Required
	}

	if len(_supportCategoryCode) > 0 {
		input.CategoryCode = aws.String(_supportCategoryCode)
	}
	if len(_supportIssueType) > 0 {
		input.IssueType = aws.String(_supportIssueType)
	}
	if len(_supportLanguage) > 0 {
		input.Language = aws.String(_supportLanguage)
	}
	if len(_supportServiceCode) > 0 {
		input.ServiceCode = aws.String(_supportServiceCode)
	}

	if resp, err := client.DescribeCreateCaseOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the current list of Amazon Web Services services and a list of service
// categories for each service. You then use service names and categories in your CreateCase
// requests. Each Amazon Web Services service has its own set of categories.
//
// The service codes and category codes correspond to the values that appear in
// the Service and Category lists on the Amazon Web Services Support Center [Create Case]page.
// The values in those fields don't necessarily match the service codes and
// categories returned by the DescribeServices operation. Always use the service
// codes and categories that the DescribeServices operation returns, so that you
// have the most recent set of service and category codes.
//
// - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan to
// use the Amazon Web Services Support API.
//
// - If you call the Amazon Web Services Support API from an account that
// doesn't have a Business, Enterprise On-Ramp, or Enterprise Support plan, the
// SubscriptionRequiredException error message appears. For information about
// changing your support plan, see [Amazon Web Services Support].
//
// [Amazon Web Services Support]: http://aws.amazon.com/premiumsupport/
// [Create Case]: https://console.aws.amazon.com/support/home#/case/create
func support_DescribeServices(cfg aws.Config, client *support.Client) {
	input := &support.DescribeServicesInput{}

	if len(_supportLanguage) > 0 {
		input.Language = aws.String(_supportLanguage)
	}
	if len(_supportServiceCodeList) > 0 {
		input.ServiceCodeList = append([]string(nil), _supportServiceCodeList...)
	}

	if resp, err := client.DescribeServices(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the list of severity levels that you can assign to a support case. The
// severity level for a case is also a field in the CaseDetailsdata type that you include for
// a CreateCaserequest.
//
// - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan to
// use the Amazon Web Services Support API.
//
// - If you call the Amazon Web Services Support API from an account that
// doesn't have a Business, Enterprise On-Ramp, or Enterprise Support plan, the
// SubscriptionRequiredException error message appears. For information about
// changing your support plan, see [Amazon Web Services Support].
//
// [Amazon Web Services Support]: http://aws.amazon.com/premiumsupport/
func support_DescribeSeverityLevels(cfg aws.Config, client *support.Client) {
	input := &support.DescribeSeverityLevelsInput{}

	if len(_supportLanguage) > 0 {
		input.Language = aws.String(_supportLanguage)
	}

	if resp, err := client.DescribeSeverityLevels(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of supported languages for a specified categoryCode , issueType
// and serviceCode . The returned supported languages will include a ISO 639-1 code
// for the language , and the language display name.
//
// - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan to
// use the Amazon Web Services Support API.
//
// - If you call the Amazon Web Services Support API from an account that
// doesn't have a Business, Enterprise On-Ramp, or Enterprise Support plan, the
// SubscriptionRequiredException error message appears. For information about
// changing your support plan, see [Amazon Web Services Support].
//
// [Amazon Web Services Support]: http://aws.amazon.com/premiumsupport/
func support_DescribeSupportedLanguages(cfg aws.Config, client *support.Client) {
	input := &support.DescribeSupportedLanguagesInput{
		// CategoryCode: *string, // Required
		// IssueType: *string, // Required
		// ServiceCode: *string, // Required
	}

	if len(_supportCategoryCode) > 0 {
		input.CategoryCode = aws.String(_supportCategoryCode)
	}
	if len(_supportIssueType) > 0 {
		input.IssueType = aws.String(_supportIssueType)
	}
	if len(_supportServiceCode) > 0 {
		input.ServiceCode = aws.String(_supportServiceCode)
	}

	if resp, err := client.DescribeSupportedLanguages(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the refresh status of the Trusted Advisor checks that have the
// specified check IDs. You can get the check IDs by calling the DescribeTrustedAdvisorChecksoperation.
//
// Some checks are refreshed automatically, and you can't return their refresh
// statuses by using the DescribeTrustedAdvisorCheckRefreshStatuses operation. If
// you call this operation for these checks, you might see an InvalidParameterValue
// error.
//
// - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan to
// use the Amazon Web Services Support API.
//
// - If you call the Amazon Web Services Support API from an account that
// doesn't have a Business, Enterprise On-Ramp, or Enterprise Support plan, the
// SubscriptionRequiredException error message appears. For information about
// changing your support plan, see [Amazon Web Services Support].
//
// To call the Trusted Advisor operations in the Amazon Web Services Support API,
// you must use the US East (N. Virginia) endpoint. Currently, the US West (Oregon)
// and Europe (Ireland) endpoints don't support the Trusted Advisor operations. For
// more information, see [About the Amazon Web Services Support API]in the Amazon Web Services Support User Guide.
//
// [Amazon Web Services Support]: http://aws.amazon.com/premiumsupport/
// [About the Amazon Web Services Support API]: https://docs.aws.amazon.com/awssupport/latest/user/about-support-api.html#endpoint
func support_DescribeTrustedAdvisorCheckRefreshStatuses(cfg aws.Config, client *support.Client) {
	input := &support.DescribeTrustedAdvisorCheckRefreshStatusesInput{
		// CheckIds: []*string, // Required
	}

	if len(_supportCheckIds) > 0 {
		if err := assignInputField(input, "CheckIds", _supportCheckIds); err != nil {
			log.Errorf("invalid --check-ids: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeTrustedAdvisorCheckRefreshStatuses(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the results of the Trusted Advisor check that has the specified check
// ID. You can get the check IDs by calling the DescribeTrustedAdvisorChecksoperation.
//
// The response contains a TrustedAdvisorCheckResult object, which contains these three objects:
//
// # TrustedAdvisorCategorySpecificSummary
//
// # TrustedAdvisorResourceDetail
//
// # TrustedAdvisorResourcesSummary
//
// In addition, the response contains these fields:
//
// - status - The alert status of the check can be ok (green), warning (yellow),
// error (red), or not_available .
//
// - timestamp - The time of the last refresh of the check.
//
// - checkId - The unique identifier for the check.
//
// - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan to
// use the Amazon Web Services Support API.
//
// - If you call the Amazon Web Services Support API from an account that
// doesn't have a Business, Enterprise On-Ramp, or Enterprise Support plan, the
// SubscriptionRequiredException error message appears. For information about
// changing your support plan, see [Amazon Web Services Support].
//
// To call the Trusted Advisor operations in the Amazon Web Services Support API,
// you must use the US East (N. Virginia) endpoint. Currently, the US West (Oregon)
// and Europe (Ireland) endpoints don't support the Trusted Advisor operations. For
// more information, see [About the Amazon Web Services Support API]in the Amazon Web Services Support User Guide.
//
// [Amazon Web Services Support]: http://aws.amazon.com/premiumsupport/
// [About the Amazon Web Services Support API]: https://docs.aws.amazon.com/awssupport/latest/user/about-support-api.html#endpoint
func support_DescribeTrustedAdvisorCheckResult(cfg aws.Config, client *support.Client) {
	input := &support.DescribeTrustedAdvisorCheckResultInput{
		// CheckId: *string, // Required
	}

	if len(_supportCheckId) > 0 {
		input.CheckId = aws.String(_supportCheckId)
	}
	if len(_supportLanguage) > 0 {
		input.Language = aws.String(_supportLanguage)
	}

	if resp, err := client.DescribeTrustedAdvisorCheckResult(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the results for the Trusted Advisor check summaries for the check IDs
// that you specified. You can get the check IDs by calling the DescribeTrustedAdvisorChecksoperation.
//
// The response contains an array of TrustedAdvisorCheckSummary objects.
//
// - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan to
// use the Amazon Web Services Support API.
//
// - If you call the Amazon Web Services Support API from an account that
// doesn't have a Business, Enterprise On-Ramp, or Enterprise Support plan, the
// SubscriptionRequiredException error message appears. For information about
// changing your support plan, see [Amazon Web Services Support].
//
// To call the Trusted Advisor operations in the Amazon Web Services Support API,
// you must use the US East (N. Virginia) endpoint. Currently, the US West (Oregon)
// and Europe (Ireland) endpoints don't support the Trusted Advisor operations. For
// more information, see [About the Amazon Web Services Support API]in the Amazon Web Services Support User Guide.
//
// [Amazon Web Services Support]: http://aws.amazon.com/premiumsupport/
// [About the Amazon Web Services Support API]: https://docs.aws.amazon.com/awssupport/latest/user/about-support-api.html#endpoint
func support_DescribeTrustedAdvisorCheckSummaries(cfg aws.Config, client *support.Client) {
	input := &support.DescribeTrustedAdvisorCheckSummariesInput{
		// CheckIds: []*string, // Required
	}

	if len(_supportCheckIds) > 0 {
		if err := assignInputField(input, "CheckIds", _supportCheckIds); err != nil {
			log.Errorf("invalid --check-ids: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeTrustedAdvisorCheckSummaries(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about all available Trusted Advisor checks, including the
// name, ID, category, description, and metadata. You must specify a language code.
//
// The response contains a TrustedAdvisorCheckDescription object for each check. You must set the Amazon Web
// Services Region to us-east-1.
//
// - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan to
// use the Amazon Web Services Support API.
//
// - If you call the Amazon Web Services Support API from an account that
// doesn't have a Business, Enterprise On-Ramp, or Enterprise Support plan, the
// SubscriptionRequiredException error message appears. For information about
// changing your support plan, see [Amazon Web Services Support].
//
// - The names and descriptions for Trusted Advisor checks are subject to
// change. We recommend that you specify the check ID in your code to uniquely
// identify a check.
//
// To call the Trusted Advisor operations in the Amazon Web Services Support API,
// you must use the US East (N. Virginia) endpoint. Currently, the US West (Oregon)
// and Europe (Ireland) endpoints don't support the Trusted Advisor operations. For
// more information, see [About the Amazon Web Services Support API]in the Amazon Web Services Support User Guide.
//
// [Amazon Web Services Support]: http://aws.amazon.com/premiumsupport/
// [About the Amazon Web Services Support API]: https://docs.aws.amazon.com/awssupport/latest/user/about-support-api.html#endpoint
func support_DescribeTrustedAdvisorChecks(cfg aws.Config, client *support.Client) {
	input := &support.DescribeTrustedAdvisorChecksInput{
		// Language: *string, // Required
	}

	if len(_supportLanguage) > 0 {
		input.Language = aws.String(_supportLanguage)
	}

	if resp, err := client.DescribeTrustedAdvisorChecks(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Refreshes the Trusted Advisor check that you specify using the check ID. You
// can get the check IDs by calling the DescribeTrustedAdvisorChecksoperation.
//
// Some checks are refreshed automatically. If you call the
// RefreshTrustedAdvisorCheck operation to refresh them, you might see the
// InvalidParameterValue error.
//
// The response contains a TrustedAdvisorCheckRefreshStatus object.
//
// - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan to
// use the Amazon Web Services Support API.
//
// - If you call the Amazon Web Services Support API from an account that
// doesn't have a Business, Enterprise On-Ramp, or Enterprise Support plan, the
// SubscriptionRequiredException error message appears. For information about
// changing your support plan, see [Amazon Web Services Support].
//
// To call the Trusted Advisor operations in the Amazon Web Services Support API,
// you must use the US East (N. Virginia) endpoint. Currently, the US West (Oregon)
// and Europe (Ireland) endpoints don't support the Trusted Advisor operations. For
// more information, see [About the Amazon Web Services Support API]in the Amazon Web Services Support User Guide.
//
// [Amazon Web Services Support]: http://aws.amazon.com/premiumsupport/
// [About the Amazon Web Services Support API]: https://docs.aws.amazon.com/awssupport/latest/user/about-support-api.html#endpoint
func support_RefreshTrustedAdvisorCheck(cfg aws.Config, client *support.Client) {
	input := &support.RefreshTrustedAdvisorCheckInput{
		// CheckId: *string, // Required
	}

	if len(_supportCheckId) > 0 {
		input.CheckId = aws.String(_supportCheckId)
	}

	if resp, err := client.RefreshTrustedAdvisorCheck(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Resolves a support case. This operation takes a caseId and returns the initial
// and final state of the case.
//
// - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan to
// use the Amazon Web Services Support API.
//
// - If you call the Amazon Web Services Support API from an account that
// doesn't have a Business, Enterprise On-Ramp, or Enterprise Support plan, the
// SubscriptionRequiredException error message appears. For information about
// changing your support plan, see [Amazon Web Services Support].
//
// [Amazon Web Services Support]: http://aws.amazon.com/premiumsupport/
func support_ResolveCase(cfg aws.Config, client *support.Client) {
	input := &support.ResolveCaseInput{}

	if len(_supportCaseId) > 0 {
		input.CaseId = aws.String(_supportCaseId)
	}

	if resp, err := client.ResolveCase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_supportCmd)
	_supportCmd.Flags().SortFlags = false

	_supportCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_supportCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_supportCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_supportCmd.Flags().StringVarP(&_supportAfterTime, "after-time", "", "", "After Time")
	_supportCmd.Flags().StringVarP(&_supportAttachmentId, "attachment-id", "", "", "Attachment ID")
	_supportCmd.Flags().StringVarP(&_supportAttachmentSetId, "attachment-set-id", "", "", "Attachment Set ID")
	_supportCmd.Flags().StringVarP(&_supportAttachments, "attachments", "", "", "Attachments")
	_supportCmd.Flags().StringVarP(&_supportBeforeTime, "before-time", "", "", "Before Time")
	_supportCmd.Flags().StringVarP(&_supportCaseId, "case-id", "", "", "Case ID")
	_supportCmd.Flags().StringSliceVarP(&_supportCaseIdList, "case-id-list", "", nil, "Case ID List")
	_supportCmd.Flags().StringVarP(&_supportCategoryCode, "category-code", "", "", "Category Code")
	_supportCmd.Flags().StringSliceVarP(&_supportCcEmailAddresses, "cc-email-addresses", "", nil, "Cc Email Addresses")
	_supportCmd.Flags().StringVarP(&_supportCheckId, "check-id", "", "", "Check ID")
	_supportCmd.Flags().StringVarP(&_supportCheckIds, "check-ids", "", "", "Check Ids")
	_supportCmd.Flags().StringVarP(&_supportCommunicationBody, "communication-body", "", "", "Communication Body")
	_supportCmd.Flags().StringVarP(&_supportDisplayId, "display-id", "", "", "Display ID")
	_supportCmd.Flags().StringVarP(&_supportIncludeCommunications, "include-communications", "", "", "Include Communications")
	_supportCmd.Flags().StringVarP(&_supportIncludeResolvedCases, "include-resolved-cases", "", "", "Include Resolved Cases")
	_supportCmd.Flags().StringVarP(&_supportIssueType, "issue-type", "", "", "Issue Type")
	_supportCmd.Flags().StringVarP(&_supportLanguage, "language", "", "", "Language")
	_supportCmd.Flags().StringVarP(&_supportMaxResults, "max-results", "", "", "Max Results")
	_supportCmd.Flags().StringVarP(&_supportNextToken, "next-token", "", "", "Next Token")
	_supportCmd.Flags().StringVarP(&_supportServiceCode, "service-code", "", "", "Service Code")
	_supportCmd.Flags().StringSliceVarP(&_supportServiceCodeList, "service-code-list", "", nil, "Service Code List")
	_supportCmd.Flags().StringVarP(&_supportSeverityCode, "severity-code", "", "", "Severity Code")
	_supportCmd.Flags().StringVarP(&_supportSubject, "subject", "", "", "Subject")

	_supportCmd.Flags().BoolVarP(&_supportAddAttachmentsToSet, "add-attachments-to-set", "", false, "Add Attachments To Set")
	_supportCmd.Flags().BoolVarP(&_supportAddCommunicationToCase, "add-communication-to-case", "", false, "Add Communication To Case")
	_supportCmd.Flags().BoolVarP(&_supportCreateCase, "create-case", "", false, "Create Case")
	_supportCmd.Flags().BoolVarP(&_supportDescribeAttachment, "describe-attachment", "", false, "Describe Attachment")
	_supportCmd.Flags().BoolVarP(&_supportDescribeCases, "describe-cases", "", false, "Describe Cases")
	_supportCmd.Flags().BoolVarP(&_supportDescribeCommunications, "describe-communications", "", false, "Describe Communications")
	_supportCmd.Flags().BoolVarP(&_supportDescribeCreateCaseOptions, "describe-create-case-options", "", false, "Describe Create Case Options")
	_supportCmd.Flags().BoolVarP(&_supportDescribeServices, "describe-services", "", false, "Describe Services")
	_supportCmd.Flags().BoolVarP(&_supportDescribeSeverityLevels, "describe-severity-levels", "", false, "Describe Severity Levels")
	_supportCmd.Flags().BoolVarP(&_supportDescribeSupportedLanguages, "describe-supported-languages", "", false, "Describe Supported Languages")
	_supportCmd.Flags().BoolVarP(&_supportDescribeTrustedAdvisorCheckRefreshStatuses, "describe-trusted-advisor-check-refresh-statuses", "", false, "Describe Trusted Advisor Check Refresh Statuses")
	_supportCmd.Flags().BoolVarP(&_supportDescribeTrustedAdvisorCheckResult, "describe-trusted-advisor-check-result", "", false, "Describe Trusted Advisor Check Result")
	_supportCmd.Flags().BoolVarP(&_supportDescribeTrustedAdvisorCheckSummaries, "describe-trusted-advisor-check-summaries", "", false, "Describe Trusted Advisor Check Summaries")
	_supportCmd.Flags().BoolVarP(&_supportDescribeTrustedAdvisorChecks, "describe-trusted-advisor-checks", "", false, "Describe Trusted Advisor Checks")
	_supportCmd.Flags().BoolVarP(&_supportRefreshTrustedAdvisorCheck, "refresh-trusted-advisor-check", "", false, "Refresh Trusted Advisor Check")
	_supportCmd.Flags().BoolVarP(&_supportResolveCase, "resolve-case", "", false, "Resolve Case")

}
