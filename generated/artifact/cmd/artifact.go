package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/artifact"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// artifactCmd represents the artifact command
var _artifactCmd = &cobra.Command{
	Use:   "artifact",
	Short: "AWS artifact CLI",
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
		client := artifact.NewFromConfig(cfg)
		if _artifactGetAccountSettings {
			artifact_GetAccountSettings(cfg, client)
			return
		}
		if _artifactGetReport {
			artifact_GetReport(cfg, client)
			return
		}
		if _artifactGetReportMetadata {
			artifact_GetReportMetadata(cfg, client)
			return
		}
		if _artifactGetTermForReport {
			artifact_GetTermForReport(cfg, client)
			return
		}
		if _artifactListCustomerAgreements {
			artifact_ListCustomerAgreements(cfg, client)
			return
		}
		if _artifactListReportVersions {
			artifact_ListReportVersions(cfg, client)
			return
		}
		if _artifactListReports {
			artifact_ListReports(cfg, client)
			return
		}
		if _artifactPutAccountSettings {
			artifact_PutAccountSettings(cfg, client)
			return
		}

	},
}

var (
	_artifactGetAccountSettings     bool
	_artifactGetReport              bool
	_artifactGetReportMetadata      bool
	_artifactGetTermForReport       bool
	_artifactListCustomerAgreements bool
	_artifactListReportVersions     bool
	_artifactListReports            bool
	_artifactPutAccountSettings     bool

	_artifactMaxResults                     string
	_artifactNextToken                      string
	_artifactNotificationSubscriptionStatus string
	_artifactReportId                       string
	_artifactReportVersion                  string
	_artifactTermToken                      string
)

// Get the account settings for Artifact.
func artifact_GetAccountSettings(cfg aws.Config, client *artifact.Client) {
	input := &artifact.GetAccountSettingsInput{}

	if resp, err := client.GetAccountSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the content for a single report.
func artifact_GetReport(cfg aws.Config, client *artifact.Client) {
	input := &artifact.GetReportInput{
		// ReportId: *string, // Required
		// TermToken: *string, // Required
	}

	if len(_artifactReportId) > 0 {
		input.ReportId = aws.String(_artifactReportId)
	}
	if len(_artifactTermToken) > 0 {
		input.TermToken = aws.String(_artifactTermToken)
	}
	if len(_artifactReportVersion) > 0 {
		if err := assignInputField(input, "ReportVersion", _artifactReportVersion); err != nil {
			log.Errorf("invalid --report-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the metadata for a single report.
func artifact_GetReportMetadata(cfg aws.Config, client *artifact.Client) {
	input := &artifact.GetReportMetadataInput{
		// ReportId: *string, // Required
	}

	if len(_artifactReportId) > 0 {
		input.ReportId = aws.String(_artifactReportId)
	}
	if len(_artifactReportVersion) > 0 {
		if err := assignInputField(input, "ReportVersion", _artifactReportVersion); err != nil {
			log.Errorf("invalid --report-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetReportMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the Term content associated with a single report.
func artifact_GetTermForReport(cfg aws.Config, client *artifact.Client) {
	input := &artifact.GetTermForReportInput{
		// ReportId: *string, // Required
	}

	if len(_artifactReportId) > 0 {
		input.ReportId = aws.String(_artifactReportId)
	}
	if len(_artifactReportVersion) > 0 {
		if err := assignInputField(input, "ReportVersion", _artifactReportVersion); err != nil {
			log.Errorf("invalid --report-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetTermForReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List active customer-agreements applicable to calling identity.
func artifact_ListCustomerAgreements(cfg aws.Config, client *artifact.Client) {
	input := &artifact.ListCustomerAgreementsInput{}

	if len(_artifactMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _artifactMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_artifactNextToken) > 0 {
		input.NextToken = aws.String(_artifactNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCustomerAgreements(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*artifact.ListCustomerAgreementsOutput
	p := artifact.NewListCustomerAgreementsPaginator(client, input)
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

// List available report versions for a given report.
func artifact_ListReportVersions(cfg aws.Config, client *artifact.Client) {
	input := &artifact.ListReportVersionsInput{
		// ReportId: *string, // Required
	}

	if len(_artifactReportId) > 0 {
		input.ReportId = aws.String(_artifactReportId)
	}
	if len(_artifactMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _artifactMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_artifactNextToken) > 0 {
		input.NextToken = aws.String(_artifactNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListReportVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*artifact.ListReportVersionsOutput
	p := artifact.NewListReportVersionsPaginator(client, input)
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

// List available reports.
func artifact_ListReports(cfg aws.Config, client *artifact.Client) {
	input := &artifact.ListReportsInput{}

	if len(_artifactMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _artifactMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_artifactNextToken) > 0 {
		input.NextToken = aws.String(_artifactNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListReports(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*artifact.ListReportsOutput
	p := artifact.NewListReportsPaginator(client, input)
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

// Put the account settings for Artifact.
func artifact_PutAccountSettings(cfg aws.Config, client *artifact.Client) {
	input := &artifact.PutAccountSettingsInput{}

	if len(_artifactNotificationSubscriptionStatus) > 0 {
		if err := assignInputField(input, "NotificationSubscriptionStatus", _artifactNotificationSubscriptionStatus); err != nil {
			log.Errorf("invalid --notification-subscription-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutAccountSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_artifactCmd)
	_artifactCmd.Flags().SortFlags = false

	_artifactCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_artifactCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_artifactCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_artifactCmd.Flags().StringVarP(&_artifactMaxResults, "max-results", "", "", "Max Results")
	_artifactCmd.Flags().StringVarP(&_artifactNextToken, "next-token", "", "", "Next Token")
	_artifactCmd.Flags().StringVarP(&_artifactNotificationSubscriptionStatus, "notification-subscription-status", "", "", "Notification Subscription Status")
	_artifactCmd.Flags().StringVarP(&_artifactReportId, "report-id", "", "", "Report ID")
	_artifactCmd.Flags().StringVarP(&_artifactReportVersion, "report-version", "", "", "Report Version")
	_artifactCmd.Flags().StringVarP(&_artifactTermToken, "term-token", "", "", "Term Token")

	_artifactCmd.Flags().BoolVarP(&_artifactGetAccountSettings, "get-account-settings", "", false, "Get Account Settings")
	_artifactCmd.Flags().BoolVarP(&_artifactGetReport, "get-report", "", false, "Get Report")
	_artifactCmd.Flags().BoolVarP(&_artifactGetReportMetadata, "get-report-metadata", "", false, "Get Report Metadata")
	_artifactCmd.Flags().BoolVarP(&_artifactGetTermForReport, "get-term-for-report", "", false, "Get Term For Report")
	_artifactCmd.Flags().BoolVarP(&_artifactListCustomerAgreements, "list-customer-agreements", "", false, "List Customer Agreements")
	_artifactCmd.Flags().BoolVarP(&_artifactListReportVersions, "list-report-versions", "", false, "List Report Versions")
	_artifactCmd.Flags().BoolVarP(&_artifactListReports, "list-reports", "", false, "List Reports")
	_artifactCmd.Flags().BoolVarP(&_artifactPutAccountSettings, "put-account-settings", "", false, "Put Account Settings")

}
