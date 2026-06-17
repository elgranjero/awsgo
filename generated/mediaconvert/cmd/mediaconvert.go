package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// mediaconvertCmd represents the mediaconvert command
var _mediaconvertCmd = &cobra.Command{
	Use:   "mediaconvert",
	Short: "AWS mediaconvert CLI",
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
		client := mediaconvert.NewFromConfig(cfg)
		if _mediaconvertAssociateCertificate {
			mediaconvert_AssociateCertificate(cfg, client)
			return
		}
		if _mediaconvertCancelJob {
			mediaconvert_CancelJob(cfg, client)
			return
		}
		if _mediaconvertCreateJob {
			mediaconvert_CreateJob(cfg, client)
			return
		}
		if _mediaconvertCreateJobTemplate {
			mediaconvert_CreateJobTemplate(cfg, client)
			return
		}
		if _mediaconvertCreatePreset {
			mediaconvert_CreatePreset(cfg, client)
			return
		}
		if _mediaconvertCreateQueue {
			mediaconvert_CreateQueue(cfg, client)
			return
		}
		if _mediaconvertCreateResourceShare {
			mediaconvert_CreateResourceShare(cfg, client)
			return
		}
		if _mediaconvertDeleteJobTemplate {
			mediaconvert_DeleteJobTemplate(cfg, client)
			return
		}
		if _mediaconvertDeletePolicy {
			mediaconvert_DeletePolicy(cfg, client)
			return
		}
		if _mediaconvertDeletePreset {
			mediaconvert_DeletePreset(cfg, client)
			return
		}
		if _mediaconvertDeleteQueue {
			mediaconvert_DeleteQueue(cfg, client)
			return
		}
		if _mediaconvertDescribeEndpoints {
			mediaconvert_DescribeEndpoints(cfg, client)
			return
		}
		if _mediaconvertDisassociateCertificate {
			mediaconvert_DisassociateCertificate(cfg, client)
			return
		}
		if _mediaconvertGetJob {
			mediaconvert_GetJob(cfg, client)
			return
		}
		if _mediaconvertGetJobTemplate {
			mediaconvert_GetJobTemplate(cfg, client)
			return
		}
		if _mediaconvertGetJobsQueryResults {
			mediaconvert_GetJobsQueryResults(cfg, client)
			return
		}
		if _mediaconvertGetPolicy {
			mediaconvert_GetPolicy(cfg, client)
			return
		}
		if _mediaconvertGetPreset {
			mediaconvert_GetPreset(cfg, client)
			return
		}
		if _mediaconvertGetQueue {
			mediaconvert_GetQueue(cfg, client)
			return
		}
		if _mediaconvertListJobTemplates {
			mediaconvert_ListJobTemplates(cfg, client)
			return
		}
		if _mediaconvertListJobs {
			mediaconvert_ListJobs(cfg, client)
			return
		}
		if _mediaconvertListPresets {
			mediaconvert_ListPresets(cfg, client)
			return
		}
		if _mediaconvertListQueues {
			mediaconvert_ListQueues(cfg, client)
			return
		}
		if _mediaconvertListTagsForResource {
			mediaconvert_ListTagsForResource(cfg, client)
			return
		}
		if _mediaconvertListVersions {
			mediaconvert_ListVersions(cfg, client)
			return
		}
		if _mediaconvertProbe {
			mediaconvert_Probe(cfg, client)
			return
		}
		if _mediaconvertPutPolicy {
			mediaconvert_PutPolicy(cfg, client)
			return
		}
		if _mediaconvertSearchJobs {
			mediaconvert_SearchJobs(cfg, client)
			return
		}
		if _mediaconvertStartJobsQuery {
			mediaconvert_StartJobsQuery(cfg, client)
			return
		}
		if _mediaconvertTagResource {
			mediaconvert_TagResource(cfg, client)
			return
		}
		if _mediaconvertUntagResource {
			mediaconvert_UntagResource(cfg, client)
			return
		}
		if _mediaconvertUpdateJobTemplate {
			mediaconvert_UpdateJobTemplate(cfg, client)
			return
		}
		if _mediaconvertUpdatePreset {
			mediaconvert_UpdatePreset(cfg, client)
			return
		}
		if _mediaconvertUpdateQueue {
			mediaconvert_UpdateQueue(cfg, client)
			return
		}

	},
}

var (
	_mediaconvertAssociateCertificate    bool
	_mediaconvertCancelJob               bool
	_mediaconvertCreateJob               bool
	_mediaconvertCreateJobTemplate       bool
	_mediaconvertCreatePreset            bool
	_mediaconvertCreateQueue             bool
	_mediaconvertCreateResourceShare     bool
	_mediaconvertDeleteJobTemplate       bool
	_mediaconvertDeletePolicy            bool
	_mediaconvertDeletePreset            bool
	_mediaconvertDeleteQueue             bool
	_mediaconvertDescribeEndpoints       bool
	_mediaconvertDisassociateCertificate bool
	_mediaconvertGetJob                  bool
	_mediaconvertGetJobTemplate          bool
	_mediaconvertGetJobsQueryResults     bool
	_mediaconvertGetPolicy               bool
	_mediaconvertGetPreset               bool
	_mediaconvertGetQueue                bool
	_mediaconvertListJobTemplates        bool
	_mediaconvertListJobs                bool
	_mediaconvertListPresets             bool
	_mediaconvertListQueues              bool
	_mediaconvertListTagsForResource     bool
	_mediaconvertListVersions            bool
	_mediaconvertProbe                   bool
	_mediaconvertPutPolicy               bool
	_mediaconvertSearchJobs              bool
	_mediaconvertStartJobsQuery          bool
	_mediaconvertTagResource             bool
	_mediaconvertUntagResource           bool
	_mediaconvertUpdateJobTemplate       bool
	_mediaconvertUpdatePreset            bool
	_mediaconvertUpdateQueue             bool

	_mediaconvertAccelerationSettings    string
	_mediaconvertArn                     string
	_mediaconvertBillingTagsSource       string
	_mediaconvertCategory                string
	_mediaconvertClientRequestToken      string
	_mediaconvertConcurrentJobs          string
	_mediaconvertDescription             string
	_mediaconvertFilterList              string
	_mediaconvertHopDestinations         string
	_mediaconvertId                      string
	_mediaconvertInputFile               string
	_mediaconvertInputFiles              string
	_mediaconvertJobEngineVersion        string
	_mediaconvertJobId                   string
	_mediaconvertJobTemplate             string
	_mediaconvertListBy                  string
	_mediaconvertMaxResults              string
	_mediaconvertMode                    string
	_mediaconvertName                    string
	_mediaconvertNextToken               string
	_mediaconvertOrder                   string
	_mediaconvertPolicy                  string
	_mediaconvertPricingPlan             string
	_mediaconvertPriority                string
	_mediaconvertQueue                   string
	_mediaconvertReservationPlanSettings string
	_mediaconvertRole                    string
	_mediaconvertSettings                string
	_mediaconvertSimulateReservedQueue   string
	_mediaconvertStatus                  string
	_mediaconvertStatusUpdateInterval    string
	_mediaconvertSupportCaseId           string
	_mediaconvertTagKeys                 []string
	_mediaconvertTags                    string
	_mediaconvertUserMetadata            string
)

// Associates an AWS Certificate Manager (ACM) Amazon Resource Name (ARN) with AWS
// Elemental MediaConvert.
func mediaconvert_AssociateCertificate(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.AssociateCertificateInput{
		// Arn: *string, // Required
	}

	if len(_mediaconvertArn) > 0 {
		input.Arn = aws.String(_mediaconvertArn)
	}

	if resp, err := client.AssociateCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Permanently cancel a job. Once you have canceled a job, you can't start it
// again.
func mediaconvert_CancelJob(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.CancelJobInput{
		// Id: *string, // Required
	}

	if len(_mediaconvertId) > 0 {
		input.Id = aws.String(_mediaconvertId)
	}

	if resp, err := client.CancelJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a new transcoding job. For information about jobs and job settings, see
// the User Guide at http://docs.aws.amazon.com/mediaconvert/latest/ug/what-is.html
func mediaconvert_CreateJob(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.CreateJobInput{
		// Role: *string, // Required
		// Settings: *types.JobSettings, // Required
	}

	if len(_mediaconvertRole) > 0 {
		input.Role = aws.String(_mediaconvertRole)
	}
	if len(_mediaconvertSettings) > 0 {
		if err := assignInputField(input, "Settings", _mediaconvertSettings); err != nil {
			log.Errorf("invalid --settings: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertAccelerationSettings) > 0 {
		if err := assignInputField(input, "AccelerationSettings", _mediaconvertAccelerationSettings); err != nil {
			log.Errorf("invalid --acceleration-settings: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertBillingTagsSource) > 0 {
		if err := assignInputField(input, "BillingTagsSource", _mediaconvertBillingTagsSource); err != nil {
			log.Errorf("invalid --billing-tags-source: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_mediaconvertClientRequestToken)
	}
	if len(_mediaconvertHopDestinations) > 0 {
		if err := assignInputField(input, "HopDestinations", _mediaconvertHopDestinations); err != nil {
			log.Errorf("invalid --hop-destinations: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertJobEngineVersion) > 0 {
		input.JobEngineVersion = aws.String(_mediaconvertJobEngineVersion)
	}
	if len(_mediaconvertJobTemplate) > 0 {
		input.JobTemplate = aws.String(_mediaconvertJobTemplate)
	}
	if len(_mediaconvertPriority) > 0 {
		if err := assignInputField(input, "Priority", _mediaconvertPriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertQueue) > 0 {
		input.Queue = aws.String(_mediaconvertQueue)
	}
	if len(_mediaconvertSimulateReservedQueue) > 0 {
		if err := assignInputField(input, "SimulateReservedQueue", _mediaconvertSimulateReservedQueue); err != nil {
			log.Errorf("invalid --simulate-reserved-queue: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertStatusUpdateInterval) > 0 {
		if err := assignInputField(input, "StatusUpdateInterval", _mediaconvertStatusUpdateInterval); err != nil {
			log.Errorf("invalid --status-update-interval: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertTags) > 0 {
		if err := assignInputField(input, "Tags", _mediaconvertTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertUserMetadata) > 0 {
		if err := assignInputField(input, "UserMetadata", _mediaconvertUserMetadata); err != nil {
			log.Errorf("invalid --user-metadata: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a new job template. For information about job templates see the User
// Guide at http://docs.aws.amazon.com/mediaconvert/latest/ug/what-is.html
func mediaconvert_CreateJobTemplate(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.CreateJobTemplateInput{
		// Name: *string, // Required
		// Settings: *types.JobTemplateSettings, // Required
	}

	if len(_mediaconvertName) > 0 {
		input.Name = aws.String(_mediaconvertName)
	}
	if len(_mediaconvertSettings) > 0 {
		if err := assignInputField(input, "Settings", _mediaconvertSettings); err != nil {
			log.Errorf("invalid --settings: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertAccelerationSettings) > 0 {
		if err := assignInputField(input, "AccelerationSettings", _mediaconvertAccelerationSettings); err != nil {
			log.Errorf("invalid --acceleration-settings: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertCategory) > 0 {
		input.Category = aws.String(_mediaconvertCategory)
	}
	if len(_mediaconvertDescription) > 0 {
		input.Description = aws.String(_mediaconvertDescription)
	}
	if len(_mediaconvertHopDestinations) > 0 {
		if err := assignInputField(input, "HopDestinations", _mediaconvertHopDestinations); err != nil {
			log.Errorf("invalid --hop-destinations: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertPriority) > 0 {
		if err := assignInputField(input, "Priority", _mediaconvertPriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertQueue) > 0 {
		input.Queue = aws.String(_mediaconvertQueue)
	}
	if len(_mediaconvertStatusUpdateInterval) > 0 {
		if err := assignInputField(input, "StatusUpdateInterval", _mediaconvertStatusUpdateInterval); err != nil {
			log.Errorf("invalid --status-update-interval: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertTags) > 0 {
		if err := assignInputField(input, "Tags", _mediaconvertTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateJobTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a new preset. For information about job templates see the User Guide at
// http://docs.aws.amazon.com/mediaconvert/latest/ug/what-is.html
func mediaconvert_CreatePreset(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.CreatePresetInput{
		// Name: *string, // Required
		// Settings: *types.PresetSettings, // Required
	}

	if len(_mediaconvertName) > 0 {
		input.Name = aws.String(_mediaconvertName)
	}
	if len(_mediaconvertSettings) > 0 {
		if err := assignInputField(input, "Settings", _mediaconvertSettings); err != nil {
			log.Errorf("invalid --settings: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertCategory) > 0 {
		input.Category = aws.String(_mediaconvertCategory)
	}
	if len(_mediaconvertDescription) > 0 {
		input.Description = aws.String(_mediaconvertDescription)
	}
	if len(_mediaconvertTags) > 0 {
		if err := assignInputField(input, "Tags", _mediaconvertTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePreset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a new transcoding queue. For information about queues, see Working With
// Queues in the User Guide at
// https://docs.aws.amazon.com/mediaconvert/latest/ug/working-with-queues.html
func mediaconvert_CreateQueue(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.CreateQueueInput{
		// Name: *string, // Required
	}

	if len(_mediaconvertName) > 0 {
		input.Name = aws.String(_mediaconvertName)
	}
	if len(_mediaconvertConcurrentJobs) > 0 {
		if err := assignInputField(input, "ConcurrentJobs", _mediaconvertConcurrentJobs); err != nil {
			log.Errorf("invalid --concurrent-jobs: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertDescription) > 0 {
		input.Description = aws.String(_mediaconvertDescription)
	}
	if len(_mediaconvertPricingPlan) > 0 {
		if err := assignInputField(input, "PricingPlan", _mediaconvertPricingPlan); err != nil {
			log.Errorf("invalid --pricing-plan: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertReservationPlanSettings) > 0 {
		if err := assignInputField(input, "ReservationPlanSettings", _mediaconvertReservationPlanSettings); err != nil {
			log.Errorf("invalid --reservation-plan-settings: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertStatus) > 0 {
		if err := assignInputField(input, "Status", _mediaconvertStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertTags) > 0 {
		if err := assignInputField(input, "Tags", _mediaconvertTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateQueue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a new resource share request for MediaConvert resources with AWS Support.
func mediaconvert_CreateResourceShare(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.CreateResourceShareInput{
		// JobId: *string, // Required
		// SupportCaseId: *string, // Required
	}

	if len(_mediaconvertJobId) > 0 {
		input.JobId = aws.String(_mediaconvertJobId)
	}
	if len(_mediaconvertSupportCaseId) > 0 {
		input.SupportCaseId = aws.String(_mediaconvertSupportCaseId)
	}

	if resp, err := client.CreateResourceShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Permanently delete a job template you have created.
func mediaconvert_DeleteJobTemplate(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.DeleteJobTemplateInput{
		// Name: *string, // Required
	}

	if len(_mediaconvertName) > 0 {
		input.Name = aws.String(_mediaconvertName)
	}

	if resp, err := client.DeleteJobTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Permanently delete a policy that you created.
func mediaconvert_DeletePolicy(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.DeletePolicyInput{}

	if resp, err := client.DeletePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Permanently delete a preset you have created.
func mediaconvert_DeletePreset(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.DeletePresetInput{
		// Name: *string, // Required
	}

	if len(_mediaconvertName) > 0 {
		input.Name = aws.String(_mediaconvertName)
	}

	if resp, err := client.DeletePreset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Permanently delete a queue you have created.
func mediaconvert_DeleteQueue(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.DeleteQueueInput{
		// Name: *string, // Required
	}

	if len(_mediaconvertName) > 0 {
		input.Name = aws.String(_mediaconvertName)
	}

	if resp, err := client.DeleteQueue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Send a request with an empty body to the regional API endpoint to get your
// account API endpoint. Note that DescribeEndpoints is no longer required. We
// recommend that you send your requests directly to the regional endpoint instead.
//
// Deprecated: DescribeEndpoints and account specific endpoints are no longer
// required. We recommend that you send your requests directly to the regional
// endpoint instead.
func mediaconvert_DescribeEndpoints(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.DescribeEndpointsInput{}

	if len(_mediaconvertMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediaconvertMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertMode) > 0 {
		if err := assignInputField(input, "Mode", _mediaconvertMode); err != nil {
			log.Errorf("invalid --mode: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertNextToken) > 0 {
		input.NextToken = aws.String(_mediaconvertNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeEndpoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediaconvert.DescribeEndpointsOutput
	p := mediaconvert.NewDescribeEndpointsPaginator(client, input)
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

// Removes an association between the Amazon Resource Name (ARN) of an AWS
// Certificate Manager (ACM) certificate and an AWS Elemental MediaConvert
// resource.
func mediaconvert_DisassociateCertificate(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.DisassociateCertificateInput{
		// Arn: *string, // Required
	}

	if len(_mediaconvertArn) > 0 {
		input.Arn = aws.String(_mediaconvertArn)
	}

	if resp, err := client.DisassociateCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve the JSON for a specific transcoding job.
func mediaconvert_GetJob(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.GetJobInput{
		// Id: *string, // Required
	}

	if len(_mediaconvertId) > 0 {
		input.Id = aws.String(_mediaconvertId)
	}

	if resp, err := client.GetJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve the JSON for a specific job template.
func mediaconvert_GetJobTemplate(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.GetJobTemplateInput{
		// Name: *string, // Required
	}

	if len(_mediaconvertName) > 0 {
		input.Name = aws.String(_mediaconvertName)
	}

	if resp, err := client.GetJobTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve a JSON array of up to twenty of your most recent jobs matched by a
// jobs query.
func mediaconvert_GetJobsQueryResults(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.GetJobsQueryResultsInput{
		// Id: *string, // Required
	}

	if len(_mediaconvertId) > 0 {
		input.Id = aws.String(_mediaconvertId)
	}

	if resp, err := client.GetJobsQueryResults(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve the JSON for your policy.
func mediaconvert_GetPolicy(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.GetPolicyInput{}

	if resp, err := client.GetPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve the JSON for a specific preset.
func mediaconvert_GetPreset(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.GetPresetInput{
		// Name: *string, // Required
	}

	if len(_mediaconvertName) > 0 {
		input.Name = aws.String(_mediaconvertName)
	}

	if resp, err := client.GetPreset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve the JSON for a specific queue.
func mediaconvert_GetQueue(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.GetQueueInput{
		// Name: *string, // Required
	}

	if len(_mediaconvertName) > 0 {
		input.Name = aws.String(_mediaconvertName)
	}

	if resp, err := client.GetQueue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve a JSON array of up to twenty of your job templates. This will return
// the templates themselves, not just a list of them. To retrieve the next twenty
// templates, use the nextToken string returned with the array
func mediaconvert_ListJobTemplates(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.ListJobTemplatesInput{}

	if len(_mediaconvertCategory) > 0 {
		input.Category = aws.String(_mediaconvertCategory)
	}
	if len(_mediaconvertListBy) > 0 {
		if err := assignInputField(input, "ListBy", _mediaconvertListBy); err != nil {
			log.Errorf("invalid --list-by: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediaconvertMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertNextToken) > 0 {
		input.NextToken = aws.String(_mediaconvertNextToken)
	}
	if len(_mediaconvertOrder) > 0 {
		if err := assignInputField(input, "Order", _mediaconvertOrder); err != nil {
			log.Errorf("invalid --order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListJobTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediaconvert.ListJobTemplatesOutput
	p := mediaconvert.NewListJobTemplatesPaginator(client, input)
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

// Retrieve a JSON array of up to twenty of your most recently created jobs. This
// array includes in-process, completed, and errored jobs. This will return the
// jobs themselves, not just a list of the jobs. To retrieve the twenty next most
// recent jobs, use the nextToken string returned with the array.
func mediaconvert_ListJobs(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.ListJobsInput{}

	if len(_mediaconvertMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediaconvertMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertNextToken) > 0 {
		input.NextToken = aws.String(_mediaconvertNextToken)
	}
	if len(_mediaconvertOrder) > 0 {
		if err := assignInputField(input, "Order", _mediaconvertOrder); err != nil {
			log.Errorf("invalid --order: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertQueue) > 0 {
		input.Queue = aws.String(_mediaconvertQueue)
	}
	if len(_mediaconvertStatus) > 0 {
		if err := assignInputField(input, "Status", _mediaconvertStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediaconvert.ListJobsOutput
	p := mediaconvert.NewListJobsPaginator(client, input)
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

// Retrieve a JSON array of up to twenty of your presets. This will return the
// presets themselves, not just a list of them. To retrieve the next twenty
// presets, use the nextToken string returned with the array.
func mediaconvert_ListPresets(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.ListPresetsInput{}

	if len(_mediaconvertCategory) > 0 {
		input.Category = aws.String(_mediaconvertCategory)
	}
	if len(_mediaconvertListBy) > 0 {
		if err := assignInputField(input, "ListBy", _mediaconvertListBy); err != nil {
			log.Errorf("invalid --list-by: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediaconvertMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertNextToken) > 0 {
		input.NextToken = aws.String(_mediaconvertNextToken)
	}
	if len(_mediaconvertOrder) > 0 {
		if err := assignInputField(input, "Order", _mediaconvertOrder); err != nil {
			log.Errorf("invalid --order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListPresets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediaconvert.ListPresetsOutput
	p := mediaconvert.NewListPresetsPaginator(client, input)
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

// Retrieve a JSON array of up to twenty of your queues. This will return the
// queues themselves, not just a list of them. To retrieve the next twenty queues,
// use the nextToken string returned with the array.
func mediaconvert_ListQueues(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.ListQueuesInput{}

	if len(_mediaconvertListBy) > 0 {
		if err := assignInputField(input, "ListBy", _mediaconvertListBy); err != nil {
			log.Errorf("invalid --list-by: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediaconvertMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertNextToken) > 0 {
		input.NextToken = aws.String(_mediaconvertNextToken)
	}
	if len(_mediaconvertOrder) > 0 {
		if err := assignInputField(input, "Order", _mediaconvertOrder); err != nil {
			log.Errorf("invalid --order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListQueues(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediaconvert.ListQueuesOutput
	p := mediaconvert.NewListQueuesPaginator(client, input)
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

// Retrieve the tags for a MediaConvert resource.
func mediaconvert_ListTagsForResource(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.ListTagsForResourceInput{
		// Arn: *string, // Required
	}

	if len(_mediaconvertArn) > 0 {
		input.Arn = aws.String(_mediaconvertArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve a JSON array of all available Job engine versions and the date they
// expire.
func mediaconvert_ListVersions(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.ListVersionsInput{}

	if len(_mediaconvertMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediaconvertMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertNextToken) > 0 {
		input.NextToken = aws.String(_mediaconvertNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediaconvert.ListVersionsOutput
	p := mediaconvert.NewListVersionsPaginator(client, input)
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

// Use Probe to obtain detailed information about your input media files. Probe
// returns a JSON that includes container, codec, frame rate, resolution, track
// count, audio layout, captions, and more. You can use this information to learn
// more about your media files, or to help make decisions while automating your
// transcoding workflow.
func mediaconvert_Probe(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.ProbeInput{}

	if len(_mediaconvertInputFiles) > 0 {
		if err := assignInputField(input, "InputFiles", _mediaconvertInputFiles); err != nil {
			log.Errorf("invalid --input-files: %s", err.Error())
			return
		}
	}

	if resp, err := client.Probe(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create or change your policy. For more information about policies, see the user
// guide at http://docs.aws.amazon.com/mediaconvert/latest/ug/what-is.html
func mediaconvert_PutPolicy(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.PutPolicyInput{
		// Policy: *types.Policy, // Required
	}

	if len(_mediaconvertPolicy) > 0 {
		if err := assignInputField(input, "Policy", _mediaconvertPolicy); err != nil {
			log.Errorf("invalid --policy: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve a JSON array that includes job details for up to twenty of your most
// recent jobs. Optionally filter results further according to input file, queue,
// or status. To retrieve the twenty next most recent jobs, use the nextToken
// string returned with the array.
func mediaconvert_SearchJobs(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.SearchJobsInput{}

	if len(_mediaconvertInputFile) > 0 {
		input.InputFile = aws.String(_mediaconvertInputFile)
	}
	if len(_mediaconvertMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediaconvertMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertNextToken) > 0 {
		input.NextToken = aws.String(_mediaconvertNextToken)
	}
	if len(_mediaconvertOrder) > 0 {
		if err := assignInputField(input, "Order", _mediaconvertOrder); err != nil {
			log.Errorf("invalid --order: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertQueue) > 0 {
		input.Queue = aws.String(_mediaconvertQueue)
	}
	if len(_mediaconvertStatus) > 0 {
		if err := assignInputField(input, "Status", _mediaconvertStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediaconvert.SearchJobsOutput
	p := mediaconvert.NewSearchJobsPaginator(client, input)
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

// Start an asynchronous jobs query using the provided filters. To receive the
// list of jobs that match your query, call the GetJobsQueryResults API using the
// query ID returned by this API.
func mediaconvert_StartJobsQuery(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.StartJobsQueryInput{}

	if len(_mediaconvertFilterList) > 0 {
		if err := assignInputField(input, "FilterList", _mediaconvertFilterList); err != nil {
			log.Errorf("invalid --filter-list: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediaconvertMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertNextToken) > 0 {
		input.NextToken = aws.String(_mediaconvertNextToken)
	}
	if len(_mediaconvertOrder) > 0 {
		if err := assignInputField(input, "Order", _mediaconvertOrder); err != nil {
			log.Errorf("invalid --order: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartJobsQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Add tags to a MediaConvert queue, preset, or job template. For information
// about tagging, see the User Guide at
// https://docs.aws.amazon.com/mediaconvert/latest/ug/tagging-resources.html
func mediaconvert_TagResource(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.TagResourceInput{
		// Arn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_mediaconvertArn) > 0 {
		input.Arn = aws.String(_mediaconvertArn)
	}
	if len(_mediaconvertTags) > 0 {
		if err := assignInputField(input, "Tags", _mediaconvertTags); err != nil {
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

// Remove tags from a MediaConvert queue, preset, or job template. For information
// about tagging, see the User Guide at
// https://docs.aws.amazon.com/mediaconvert/latest/ug/tagging-resources.html
func mediaconvert_UntagResource(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.UntagResourceInput{
		// Arn: *string, // Required
	}

	if len(_mediaconvertArn) > 0 {
		input.Arn = aws.String(_mediaconvertArn)
	}
	if len(_mediaconvertTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _mediaconvertTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modify one of your existing job templates.
func mediaconvert_UpdateJobTemplate(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.UpdateJobTemplateInput{
		// Name: *string, // Required
	}

	if len(_mediaconvertName) > 0 {
		input.Name = aws.String(_mediaconvertName)
	}
	if len(_mediaconvertAccelerationSettings) > 0 {
		if err := assignInputField(input, "AccelerationSettings", _mediaconvertAccelerationSettings); err != nil {
			log.Errorf("invalid --acceleration-settings: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertCategory) > 0 {
		input.Category = aws.String(_mediaconvertCategory)
	}
	if len(_mediaconvertDescription) > 0 {
		input.Description = aws.String(_mediaconvertDescription)
	}
	if len(_mediaconvertHopDestinations) > 0 {
		if err := assignInputField(input, "HopDestinations", _mediaconvertHopDestinations); err != nil {
			log.Errorf("invalid --hop-destinations: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertPriority) > 0 {
		if err := assignInputField(input, "Priority", _mediaconvertPriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertQueue) > 0 {
		input.Queue = aws.String(_mediaconvertQueue)
	}
	if len(_mediaconvertSettings) > 0 {
		if err := assignInputField(input, "Settings", _mediaconvertSettings); err != nil {
			log.Errorf("invalid --settings: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertStatusUpdateInterval) > 0 {
		if err := assignInputField(input, "StatusUpdateInterval", _mediaconvertStatusUpdateInterval); err != nil {
			log.Errorf("invalid --status-update-interval: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateJobTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modify one of your existing presets.
func mediaconvert_UpdatePreset(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.UpdatePresetInput{
		// Name: *string, // Required
	}

	if len(_mediaconvertName) > 0 {
		input.Name = aws.String(_mediaconvertName)
	}
	if len(_mediaconvertCategory) > 0 {
		input.Category = aws.String(_mediaconvertCategory)
	}
	if len(_mediaconvertDescription) > 0 {
		input.Description = aws.String(_mediaconvertDescription)
	}
	if len(_mediaconvertSettings) > 0 {
		if err := assignInputField(input, "Settings", _mediaconvertSettings); err != nil {
			log.Errorf("invalid --settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePreset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modify one of your existing queues.
func mediaconvert_UpdateQueue(cfg aws.Config, client *mediaconvert.Client) {
	input := &mediaconvert.UpdateQueueInput{
		// Name: *string, // Required
	}

	if len(_mediaconvertName) > 0 {
		input.Name = aws.String(_mediaconvertName)
	}
	if len(_mediaconvertConcurrentJobs) > 0 {
		if err := assignInputField(input, "ConcurrentJobs", _mediaconvertConcurrentJobs); err != nil {
			log.Errorf("invalid --concurrent-jobs: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertDescription) > 0 {
		input.Description = aws.String(_mediaconvertDescription)
	}
	if len(_mediaconvertReservationPlanSettings) > 0 {
		if err := assignInputField(input, "ReservationPlanSettings", _mediaconvertReservationPlanSettings); err != nil {
			log.Errorf("invalid --reservation-plan-settings: %s", err.Error())
			return
		}
	}
	if len(_mediaconvertStatus) > 0 {
		if err := assignInputField(input, "Status", _mediaconvertStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateQueue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_mediaconvertCmd)
	_mediaconvertCmd.Flags().SortFlags = false

	_mediaconvertCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_mediaconvertCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_mediaconvertCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertAccelerationSettings, "acceleration-settings", "", "", "Acceleration Settings")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertArn, "arn", "", "", "ARN")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertBillingTagsSource, "billing-tags-source", "", "", "Billing Tags Source")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertCategory, "category", "", "", "Category")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertConcurrentJobs, "concurrent-jobs", "", "", "Concurrent Jobs")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertDescription, "description", "", "", "Description")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertFilterList, "filter-list", "", "", "Filter List")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertHopDestinations, "hop-destinations", "", "", "Hop Destinations")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertId, "id", "", "", "ID")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertInputFile, "input-file", "", "", "Input File")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertInputFiles, "input-files", "", "", "Input Files")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertJobEngineVersion, "job-engine-version", "", "", "Job Engine Version")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertJobId, "job-id", "", "", "Job ID")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertJobTemplate, "job-template", "", "", "Job Template")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertListBy, "list-by", "", "", "List By")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertMaxResults, "max-results", "", "", "Max Results")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertMode, "mode", "", "", "Mode")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertName, "name", "", "", "Name")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertNextToken, "next-token", "", "", "Next Token")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertOrder, "order", "", "", "Order")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertPolicy, "policy", "", "", "Policy")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertPricingPlan, "pricing-plan", "", "", "Pricing Plan")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertPriority, "priority", "", "", "Priority")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertQueue, "queue", "", "", "Queue")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertReservationPlanSettings, "reservation-plan-settings", "", "", "Reservation Plan Settings")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertRole, "role", "", "", "Role")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertSettings, "settings", "", "", "Settings")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertSimulateReservedQueue, "simulate-reserved-queue", "", "", "Simulate Reserved Queue")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertStatus, "status", "", "", "Status")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertStatusUpdateInterval, "status-update-interval", "", "", "Status Update Interval")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertSupportCaseId, "support-case-id", "", "", "Support Case ID")
	_mediaconvertCmd.Flags().StringSliceVarP(&_mediaconvertTagKeys, "tag-keys", "", nil, "Tag Keys")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertTags, "tags", "", "", "Tags")
	_mediaconvertCmd.Flags().StringVarP(&_mediaconvertUserMetadata, "user-metadata", "", "", "User Metadata")

	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertAssociateCertificate, "associate-certificate", "", false, "Associate Certificate")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertCancelJob, "cancel-job", "", false, "Cancel Job")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertCreateJob, "create-job", "", false, "Create Job")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertCreateJobTemplate, "create-job-template", "", false, "Create Job Template")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertCreatePreset, "create-preset", "", false, "Create Preset")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertCreateQueue, "create-queue", "", false, "Create Queue")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertCreateResourceShare, "create-resource-share", "", false, "Create Resource Share")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertDeleteJobTemplate, "delete-job-template", "", false, "Delete Job Template")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertDeletePolicy, "delete-policy", "", false, "Delete Policy")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertDeletePreset, "delete-preset", "", false, "Delete Preset")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertDeleteQueue, "delete-queue", "", false, "Delete Queue")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertDescribeEndpoints, "describe-endpoints", "", false, "Describe Endpoints")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertDisassociateCertificate, "disassociate-certificate", "", false, "Disassociate Certificate")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertGetJob, "get-job", "", false, "Get Job")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertGetJobTemplate, "get-job-template", "", false, "Get Job Template")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertGetJobsQueryResults, "get-jobs-query-results", "", false, "Get Jobs Query Results")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertGetPolicy, "get-policy", "", false, "Get Policy")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertGetPreset, "get-preset", "", false, "Get Preset")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertGetQueue, "get-queue", "", false, "Get Queue")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertListJobTemplates, "list-job-templates", "", false, "List Job Templates")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertListJobs, "list-jobs", "", false, "List Jobs")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertListPresets, "list-presets", "", false, "List Presets")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertListQueues, "list-queues", "", false, "List Queues")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertListVersions, "list-versions", "", false, "List Versions")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertProbe, "probe", "", false, "Probe")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertPutPolicy, "put-policy", "", false, "Put Policy")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertSearchJobs, "search-jobs", "", false, "Search Jobs")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertStartJobsQuery, "start-jobs-query", "", false, "Start Jobs Query")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertTagResource, "tag-resource", "", false, "Tag Resource")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertUntagResource, "untag-resource", "", false, "Untag Resource")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertUpdateJobTemplate, "update-job-template", "", false, "Update Job Template")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertUpdatePreset, "update-preset", "", false, "Update Preset")
	_mediaconvertCmd.Flags().BoolVarP(&_mediaconvertUpdateQueue, "update-queue", "", false, "Update Queue")

}
