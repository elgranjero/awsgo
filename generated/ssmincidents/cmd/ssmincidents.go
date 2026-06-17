package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssmincidents"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ssmincidentsCmd represents the ssmincidents command
var _ssmincidentsCmd = &cobra.Command{
	Use:   "ssmincidents",
	Short: "AWS ssmincidents CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := ssmincidents.NewFromConfig(cfg)
		if _ssmincidentsBatchGetIncidentFindings {
			ssmincidents_BatchGetIncidentFindings(cfg, client)
			return
		}
		if _ssmincidentsCreateReplicationSet {
			ssmincidents_CreateReplicationSet(cfg, client)
			return
		}
		if _ssmincidentsCreateResponsePlan {
			ssmincidents_CreateResponsePlan(cfg, client)
			return
		}
		if _ssmincidentsCreateTimelineEvent {
			ssmincidents_CreateTimelineEvent(cfg, client)
			return
		}
		if _ssmincidentsDeleteIncidentRecord {
			ssmincidents_DeleteIncidentRecord(cfg, client)
			return
		}
		if _ssmincidentsDeleteReplicationSet {
			ssmincidents_DeleteReplicationSet(cfg, client)
			return
		}
		if _ssmincidentsDeleteResourcePolicy {
			ssmincidents_DeleteResourcePolicy(cfg, client)
			return
		}
		if _ssmincidentsDeleteResponsePlan {
			ssmincidents_DeleteResponsePlan(cfg, client)
			return
		}
		if _ssmincidentsDeleteTimelineEvent {
			ssmincidents_DeleteTimelineEvent(cfg, client)
			return
		}
		if _ssmincidentsGetIncidentRecord {
			ssmincidents_GetIncidentRecord(cfg, client)
			return
		}
		if _ssmincidentsGetReplicationSet {
			ssmincidents_GetReplicationSet(cfg, client)
			return
		}
		if _ssmincidentsGetResourcePolicies {
			ssmincidents_GetResourcePolicies(cfg, client)
			return
		}
		if _ssmincidentsGetResponsePlan {
			ssmincidents_GetResponsePlan(cfg, client)
			return
		}
		if _ssmincidentsGetTimelineEvent {
			ssmincidents_GetTimelineEvent(cfg, client)
			return
		}
		if _ssmincidentsListIncidentFindings {
			ssmincidents_ListIncidentFindings(cfg, client)
			return
		}
		if _ssmincidentsListIncidentRecords {
			ssmincidents_ListIncidentRecords(cfg, client)
			return
		}
		if _ssmincidentsListRelatedItems {
			ssmincidents_ListRelatedItems(cfg, client)
			return
		}
		if _ssmincidentsListReplicationSets {
			ssmincidents_ListReplicationSets(cfg, client)
			return
		}
		if _ssmincidentsListResponsePlans {
			ssmincidents_ListResponsePlans(cfg, client)
			return
		}
		if _ssmincidentsListTagsForResource {
			ssmincidents_ListTagsForResource(cfg, client)
			return
		}
		if _ssmincidentsListTimelineEvents {
			ssmincidents_ListTimelineEvents(cfg, client)
			return
		}
		if _ssmincidentsPutResourcePolicy {
			ssmincidents_PutResourcePolicy(cfg, client)
			return
		}
		if _ssmincidentsStartIncident {
			ssmincidents_StartIncident(cfg, client)
			return
		}
		if _ssmincidentsTagResource {
			ssmincidents_TagResource(cfg, client)
			return
		}
		if _ssmincidentsUntagResource {
			ssmincidents_UntagResource(cfg, client)
			return
		}
		if _ssmincidentsUpdateDeletionProtection {
			ssmincidents_UpdateDeletionProtection(cfg, client)
			return
		}
		if _ssmincidentsUpdateIncidentRecord {
			ssmincidents_UpdateIncidentRecord(cfg, client)
			return
		}
		if _ssmincidentsUpdateRelatedItems {
			ssmincidents_UpdateRelatedItems(cfg, client)
			return
		}
		if _ssmincidentsUpdateReplicationSet {
			ssmincidents_UpdateReplicationSet(cfg, client)
			return
		}
		if _ssmincidentsUpdateResponsePlan {
			ssmincidents_UpdateResponsePlan(cfg, client)
			return
		}
		if _ssmincidentsUpdateTimelineEvent {
			ssmincidents_UpdateTimelineEvent(cfg, client)
			return
		}

	},
}

var (
	_ssmincidentsBatchGetIncidentFindings bool
	_ssmincidentsCreateReplicationSet     bool
	_ssmincidentsCreateResponsePlan       bool
	_ssmincidentsCreateTimelineEvent      bool
	_ssmincidentsDeleteIncidentRecord     bool
	_ssmincidentsDeleteReplicationSet     bool
	_ssmincidentsDeleteResourcePolicy     bool
	_ssmincidentsDeleteResponsePlan       bool
	_ssmincidentsDeleteTimelineEvent      bool
	_ssmincidentsGetIncidentRecord        bool
	_ssmincidentsGetReplicationSet        bool
	_ssmincidentsGetResourcePolicies      bool
	_ssmincidentsGetResponsePlan          bool
	_ssmincidentsGetTimelineEvent         bool
	_ssmincidentsListIncidentFindings     bool
	_ssmincidentsListIncidentRecords      bool
	_ssmincidentsListRelatedItems         bool
	_ssmincidentsListReplicationSets      bool
	_ssmincidentsListResponsePlans        bool
	_ssmincidentsListTagsForResource      bool
	_ssmincidentsListTimelineEvents       bool
	_ssmincidentsPutResourcePolicy        bool
	_ssmincidentsStartIncident            bool
	_ssmincidentsTagResource              bool
	_ssmincidentsUntagResource            bool
	_ssmincidentsUpdateDeletionProtection bool
	_ssmincidentsUpdateIncidentRecord     bool
	_ssmincidentsUpdateRelatedItems       bool
	_ssmincidentsUpdateReplicationSet     bool
	_ssmincidentsUpdateResponsePlan       bool
	_ssmincidentsUpdateTimelineEvent      bool

	_ssmincidentsActions                             string
	_ssmincidentsArn                                 string
	_ssmincidentsChatChannel                         string
	_ssmincidentsClientToken                         string
	_ssmincidentsDeletionProtected                   string
	_ssmincidentsDisplayName                         string
	_ssmincidentsEngagements                         []string
	_ssmincidentsEventData                           string
	_ssmincidentsEventId                             string
	_ssmincidentsEventReferences                     string
	_ssmincidentsEventTime                           string
	_ssmincidentsEventType                           string
	_ssmincidentsFilters                             string
	_ssmincidentsFindingIds                          []string
	_ssmincidentsImpact                              string
	_ssmincidentsIncidentRecordArn                   string
	_ssmincidentsIncidentTemplate                    string
	_ssmincidentsIncidentTemplateDedupeString        string
	_ssmincidentsIncidentTemplateImpact              string
	_ssmincidentsIncidentTemplateNotificationTargets string
	_ssmincidentsIncidentTemplateSummary             string
	_ssmincidentsIncidentTemplateTags                string
	_ssmincidentsIncidentTemplateTitle               string
	_ssmincidentsIntegrations                        string
	_ssmincidentsMaxResults                          string
	_ssmincidentsName                                string
	_ssmincidentsNextToken                           string
	_ssmincidentsNotificationTargets                 string
	_ssmincidentsPolicy                              string
	_ssmincidentsPolicyId                            string
	_ssmincidentsRegions                             string
	_ssmincidentsRelatedItems                        string
	_ssmincidentsRelatedItemsUpdate                  string
	_ssmincidentsResourceArn                         string
	_ssmincidentsResponsePlanArn                     string
	_ssmincidentsSortBy                              string
	_ssmincidentsSortOrder                           string
	_ssmincidentsStatus                              string
	_ssmincidentsSummary                             string
	_ssmincidentsTagKeys                             []string
	_ssmincidentsTags                                string
	_ssmincidentsTitle                               string
	_ssmincidentsTriggerDetails                      string
)

// Retrieves details about all specified findings for an incident, including
// descriptive details about each finding. A finding represents a recent
// application environment change made by an CodeDeploy deployment or an
// CloudFormation stack creation or update that can be investigated as a potential
// cause of the incident.
func ssmincidents_BatchGetIncidentFindings(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.BatchGetIncidentFindingsInput{
		// FindingIds: []string, // Required
		// IncidentRecordArn: *string, // Required
	}

	if len(_ssmincidentsFindingIds) > 0 {
		input.FindingIds = append([]string(nil), _ssmincidentsFindingIds...)
	}
	if len(_ssmincidentsIncidentRecordArn) > 0 {
		input.IncidentRecordArn = aws.String(_ssmincidentsIncidentRecordArn)
	}

	if resp, err := client.BatchGetIncidentFindings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A replication set replicates and encrypts your data to the provided Regions
// with the provided KMS key.
func ssmincidents_CreateReplicationSet(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.CreateReplicationSetInput{
		// Regions: map[string]types.RegionMapInputValue, // Required
	}

	if len(_ssmincidentsRegions) > 0 {
		if err := assignInputField(input, "Regions", _ssmincidentsRegions); err != nil {
			log.Errorf("invalid --regions: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsClientToken) > 0 {
		input.ClientToken = aws.String(_ssmincidentsClientToken)
	}
	if len(_ssmincidentsTags) > 0 {
		if err := assignInputField(input, "Tags", _ssmincidentsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateReplicationSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a response plan that automates the initial response to incidents. A
// response plan engages contacts, starts chat channel collaboration, and initiates
// runbooks at the beginning of an incident.
func ssmincidents_CreateResponsePlan(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.CreateResponsePlanInput{
		// IncidentTemplate: *types.IncidentTemplate, // Required
		// Name: *string, // Required
	}

	if len(_ssmincidentsIncidentTemplate) > 0 {
		if err := assignInputField(input, "IncidentTemplate", _ssmincidentsIncidentTemplate); err != nil {
			log.Errorf("invalid --incident-template: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsName) > 0 {
		input.Name = aws.String(_ssmincidentsName)
	}
	if len(_ssmincidentsActions) > 0 {
		if err := assignInputField(input, "Actions", _ssmincidentsActions); err != nil {
			log.Errorf("invalid --actions: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsChatChannel) > 0 {
		if err := assignInputField(input, "ChatChannel", _ssmincidentsChatChannel); err != nil {
			log.Errorf("invalid --chat-channel: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsClientToken) > 0 {
		input.ClientToken = aws.String(_ssmincidentsClientToken)
	}
	if len(_ssmincidentsDisplayName) > 0 {
		input.DisplayName = aws.String(_ssmincidentsDisplayName)
	}
	if len(_ssmincidentsEngagements) > 0 {
		input.Engagements = append([]string(nil), _ssmincidentsEngagements...)
	}
	if len(_ssmincidentsIntegrations) > 0 {
		if err := assignInputField(input, "Integrations", _ssmincidentsIntegrations); err != nil {
			log.Errorf("invalid --integrations: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsTags) > 0 {
		if err := assignInputField(input, "Tags", _ssmincidentsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateResponsePlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a custom timeline event on the incident details page of an incident
// record. Incident Manager automatically creates timeline events that mark key
// moments during an incident. You can create custom timeline events to mark
// important events that Incident Manager can detect automatically.
func ssmincidents_CreateTimelineEvent(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.CreateTimelineEventInput{
		// EventData: *string, // Required
		// EventTime: *time.Time, // Required
		// EventType: *string, // Required
		// IncidentRecordArn: *string, // Required
	}

	if len(_ssmincidentsEventData) > 0 {
		input.EventData = aws.String(_ssmincidentsEventData)
	}
	if len(_ssmincidentsEventTime) > 0 {
		if err := assignInputField(input, "EventTime", _ssmincidentsEventTime); err != nil {
			log.Errorf("invalid --event-time: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsEventType) > 0 {
		input.EventType = aws.String(_ssmincidentsEventType)
	}
	if len(_ssmincidentsIncidentRecordArn) > 0 {
		input.IncidentRecordArn = aws.String(_ssmincidentsIncidentRecordArn)
	}
	if len(_ssmincidentsClientToken) > 0 {
		input.ClientToken = aws.String(_ssmincidentsClientToken)
	}
	if len(_ssmincidentsEventReferences) > 0 {
		if err := assignInputField(input, "EventReferences", _ssmincidentsEventReferences); err != nil {
			log.Errorf("invalid --event-references: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTimelineEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an incident record from Incident Manager.
func ssmincidents_DeleteIncidentRecord(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.DeleteIncidentRecordInput{
		// Arn: *string, // Required
	}

	if len(_ssmincidentsArn) > 0 {
		input.Arn = aws.String(_ssmincidentsArn)
	}

	if resp, err := client.DeleteIncidentRecord(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes all Regions in your replication set. Deleting the replication set
// deletes all Incident Manager data.
func ssmincidents_DeleteReplicationSet(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.DeleteReplicationSetInput{
		// Arn: *string, // Required
	}

	if len(_ssmincidentsArn) > 0 {
		input.Arn = aws.String(_ssmincidentsArn)
	}

	if resp, err := client.DeleteReplicationSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the resource policy that Resource Access Manager uses to share your
// Incident Manager resource.
func ssmincidents_DeleteResourcePolicy(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.DeleteResourcePolicyInput{
		// PolicyId: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_ssmincidentsPolicyId) > 0 {
		input.PolicyId = aws.String(_ssmincidentsPolicyId)
	}
	if len(_ssmincidentsResourceArn) > 0 {
		input.ResourceArn = aws.String(_ssmincidentsResourceArn)
	}

	if resp, err := client.DeleteResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified response plan. Deleting a response plan stops all linked
// CloudWatch alarms and EventBridge events from creating an incident with this
// response plan.
func ssmincidents_DeleteResponsePlan(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.DeleteResponsePlanInput{
		// Arn: *string, // Required
	}

	if len(_ssmincidentsArn) > 0 {
		input.Arn = aws.String(_ssmincidentsArn)
	}

	if resp, err := client.DeleteResponsePlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a timeline event from an incident.
func ssmincidents_DeleteTimelineEvent(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.DeleteTimelineEventInput{
		// EventId: *string, // Required
		// IncidentRecordArn: *string, // Required
	}

	if len(_ssmincidentsEventId) > 0 {
		input.EventId = aws.String(_ssmincidentsEventId)
	}
	if len(_ssmincidentsIncidentRecordArn) > 0 {
		input.IncidentRecordArn = aws.String(_ssmincidentsIncidentRecordArn)
	}

	if resp, err := client.DeleteTimelineEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the details for the specified incident record.
func ssmincidents_GetIncidentRecord(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.GetIncidentRecordInput{
		// Arn: *string, // Required
	}

	if len(_ssmincidentsArn) > 0 {
		input.Arn = aws.String(_ssmincidentsArn)
	}

	if resp, err := client.GetIncidentRecord(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve your Incident Manager replication set.
func ssmincidents_GetReplicationSet(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.GetReplicationSetInput{
		// Arn: *string, // Required
	}

	if len(_ssmincidentsArn) > 0 {
		input.Arn = aws.String(_ssmincidentsArn)
	}

	if resp, err := client.GetReplicationSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the resource policies attached to the specified response plan.
func ssmincidents_GetResourcePolicies(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.GetResourcePoliciesInput{
		// ResourceArn: *string, // Required
	}

	if len(_ssmincidentsResourceArn) > 0 {
		input.ResourceArn = aws.String(_ssmincidentsResourceArn)
	}
	if len(_ssmincidentsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmincidentsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsNextToken) > 0 {
		input.NextToken = aws.String(_ssmincidentsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetResourcePolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssmincidents.GetResourcePoliciesOutput
	p := ssmincidents.NewGetResourcePoliciesPaginator(client, input)
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

// Retrieves the details of the specified response plan.
func ssmincidents_GetResponsePlan(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.GetResponsePlanInput{
		// Arn: *string, // Required
	}

	if len(_ssmincidentsArn) > 0 {
		input.Arn = aws.String(_ssmincidentsArn)
	}

	if resp, err := client.GetResponsePlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a timeline event based on its ID and incident record.
func ssmincidents_GetTimelineEvent(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.GetTimelineEventInput{
		// EventId: *string, // Required
		// IncidentRecordArn: *string, // Required
	}

	if len(_ssmincidentsEventId) > 0 {
		input.EventId = aws.String(_ssmincidentsEventId)
	}
	if len(_ssmincidentsIncidentRecordArn) > 0 {
		input.IncidentRecordArn = aws.String(_ssmincidentsIncidentRecordArn)
	}

	if resp, err := client.GetTimelineEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of the IDs of findings, plus their last modified times, that
// have been identified for a specified incident. A finding represents a recent
// application environment change made by an CloudFormation stack creation or
// update or an CodeDeploy deployment that can be investigated as a potential cause
// of the incident.
func ssmincidents_ListIncidentFindings(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.ListIncidentFindingsInput{
		// IncidentRecordArn: *string, // Required
	}

	if len(_ssmincidentsIncidentRecordArn) > 0 {
		input.IncidentRecordArn = aws.String(_ssmincidentsIncidentRecordArn)
	}
	if len(_ssmincidentsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmincidentsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsNextToken) > 0 {
		input.NextToken = aws.String(_ssmincidentsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIncidentFindings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssmincidents.ListIncidentFindingsOutput
	p := ssmincidents.NewListIncidentFindingsPaginator(client, input)
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

// Lists all incident records in your account. Use this command to retrieve the
// Amazon Resource Name (ARN) of the incident record you want to update.
func ssmincidents_ListIncidentRecords(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.ListIncidentRecordsInput{}

	if len(_ssmincidentsFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmincidentsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmincidentsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsNextToken) > 0 {
		input.NextToken = aws.String(_ssmincidentsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIncidentRecords(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssmincidents.ListIncidentRecordsOutput
	p := ssmincidents.NewListIncidentRecordsPaginator(client, input)
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

// List all related items for an incident record.
func ssmincidents_ListRelatedItems(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.ListRelatedItemsInput{
		// IncidentRecordArn: *string, // Required
	}

	if len(_ssmincidentsIncidentRecordArn) > 0 {
		input.IncidentRecordArn = aws.String(_ssmincidentsIncidentRecordArn)
	}
	if len(_ssmincidentsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmincidentsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsNextToken) > 0 {
		input.NextToken = aws.String(_ssmincidentsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRelatedItems(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssmincidents.ListRelatedItemsOutput
	p := ssmincidents.NewListRelatedItemsPaginator(client, input)
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

// Lists details about the replication set configured in your account.
func ssmincidents_ListReplicationSets(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.ListReplicationSetsInput{}

	if len(_ssmincidentsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmincidentsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsNextToken) > 0 {
		input.NextToken = aws.String(_ssmincidentsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListReplicationSets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssmincidents.ListReplicationSetsOutput
	p := ssmincidents.NewListReplicationSetsPaginator(client, input)
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

// Lists all response plans in your account.
func ssmincidents_ListResponsePlans(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.ListResponsePlansInput{}

	if len(_ssmincidentsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmincidentsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsNextToken) > 0 {
		input.NextToken = aws.String(_ssmincidentsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListResponsePlans(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssmincidents.ListResponsePlansOutput
	p := ssmincidents.NewListResponsePlansPaginator(client, input)
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

// Lists the tags that are attached to the specified response plan or incident.
func ssmincidents_ListTagsForResource(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_ssmincidentsResourceArn) > 0 {
		input.ResourceArn = aws.String(_ssmincidentsResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists timeline events for the specified incident record.
func ssmincidents_ListTimelineEvents(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.ListTimelineEventsInput{
		// IncidentRecordArn: *string, // Required
	}

	if len(_ssmincidentsIncidentRecordArn) > 0 {
		input.IncidentRecordArn = aws.String(_ssmincidentsIncidentRecordArn)
	}
	if len(_ssmincidentsFilters) > 0 {
		if err := assignInputField(input, "Filters", _ssmincidentsFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmincidentsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsNextToken) > 0 {
		input.NextToken = aws.String(_ssmincidentsNextToken)
	}
	if len(_ssmincidentsSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _ssmincidentsSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _ssmincidentsSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTimelineEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssmincidents.ListTimelineEventsOutput
	p := ssmincidents.NewListTimelineEventsPaginator(client, input)
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

// Adds a resource policy to the specified response plan. The resource policy is
// used to share the response plan using Resource Access Manager (RAM). For more
// information about cross-account sharing, see [Cross-Region and cross-account incident management].
//
// [Cross-Region and cross-account incident management]: https://docs.aws.amazon.com/incident-manager/latest/userguide/incident-manager-cross-account-cross-region.html
func ssmincidents_PutResourcePolicy(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.PutResourcePolicyInput{
		// Policy: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_ssmincidentsPolicy) > 0 {
		input.Policy = aws.String(_ssmincidentsPolicy)
	}
	if len(_ssmincidentsResourceArn) > 0 {
		input.ResourceArn = aws.String(_ssmincidentsResourceArn)
	}

	if resp, err := client.PutResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Used to start an incident from CloudWatch alarms, EventBridge events, or
// manually.
func ssmincidents_StartIncident(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.StartIncidentInput{
		// ResponsePlanArn: *string, // Required
	}

	if len(_ssmincidentsResponsePlanArn) > 0 {
		input.ResponsePlanArn = aws.String(_ssmincidentsResponsePlanArn)
	}
	if len(_ssmincidentsClientToken) > 0 {
		input.ClientToken = aws.String(_ssmincidentsClientToken)
	}
	if len(_ssmincidentsImpact) > 0 {
		if err := assignInputField(input, "Impact", _ssmincidentsImpact); err != nil {
			log.Errorf("invalid --impact: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsRelatedItems) > 0 {
		if err := assignInputField(input, "RelatedItems", _ssmincidentsRelatedItems); err != nil {
			log.Errorf("invalid --related-items: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsTitle) > 0 {
		input.Title = aws.String(_ssmincidentsTitle)
	}
	if len(_ssmincidentsTriggerDetails) > 0 {
		if err := assignInputField(input, "TriggerDetails", _ssmincidentsTriggerDetails); err != nil {
			log.Errorf("invalid --trigger-details: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartIncident(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a tag to a response plan.
func ssmincidents_TagResource(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_ssmincidentsResourceArn) > 0 {
		input.ResourceArn = aws.String(_ssmincidentsResourceArn)
	}
	if len(_ssmincidentsTags) > 0 {
		if err := assignInputField(input, "Tags", _ssmincidentsTags); err != nil {
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

// Removes a tag from a resource.
func ssmincidents_UntagResource(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_ssmincidentsResourceArn) > 0 {
		input.ResourceArn = aws.String(_ssmincidentsResourceArn)
	}
	if len(_ssmincidentsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _ssmincidentsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update deletion protection to either allow or deny deletion of the final Region
// in a replication set.
func ssmincidents_UpdateDeletionProtection(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.UpdateDeletionProtectionInput{
		// Arn: *string, // Required
		// DeletionProtected: *bool, // Required
	}

	if len(_ssmincidentsArn) > 0 {
		input.Arn = aws.String(_ssmincidentsArn)
	}
	if len(_ssmincidentsDeletionProtected) > 0 {
		if err := assignInputField(input, "DeletionProtected", _ssmincidentsDeletionProtected); err != nil {
			log.Errorf("invalid --deletion-protected: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsClientToken) > 0 {
		input.ClientToken = aws.String(_ssmincidentsClientToken)
	}

	if resp, err := client.UpdateDeletionProtection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the details of an incident record. You can use this operation to update
// an incident record from the defined chat channel. For more information about
// using actions in chat channels, see [Interacting through chat].
//
// [Interacting through chat]: https://docs.aws.amazon.com/incident-manager/latest/userguide/chat.html#chat-interact
func ssmincidents_UpdateIncidentRecord(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.UpdateIncidentRecordInput{
		// Arn: *string, // Required
	}

	if len(_ssmincidentsArn) > 0 {
		input.Arn = aws.String(_ssmincidentsArn)
	}
	if len(_ssmincidentsChatChannel) > 0 {
		if err := assignInputField(input, "ChatChannel", _ssmincidentsChatChannel); err != nil {
			log.Errorf("invalid --chat-channel: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsClientToken) > 0 {
		input.ClientToken = aws.String(_ssmincidentsClientToken)
	}
	if len(_ssmincidentsImpact) > 0 {
		if err := assignInputField(input, "Impact", _ssmincidentsImpact); err != nil {
			log.Errorf("invalid --impact: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsNotificationTargets) > 0 {
		if err := assignInputField(input, "NotificationTargets", _ssmincidentsNotificationTargets); err != nil {
			log.Errorf("invalid --notification-targets: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsStatus) > 0 {
		if err := assignInputField(input, "Status", _ssmincidentsStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsSummary) > 0 {
		input.Summary = aws.String(_ssmincidentsSummary)
	}
	if len(_ssmincidentsTitle) > 0 {
		input.Title = aws.String(_ssmincidentsTitle)
	}

	if resp, err := client.UpdateIncidentRecord(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Add or remove related items from the related items tab of an incident record.
func ssmincidents_UpdateRelatedItems(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.UpdateRelatedItemsInput{
		// IncidentRecordArn: *string, // Required
		// RelatedItemsUpdate: types.RelatedItemsUpdate, // Required
	}

	if len(_ssmincidentsIncidentRecordArn) > 0 {
		input.IncidentRecordArn = aws.String(_ssmincidentsIncidentRecordArn)
	}
	if len(_ssmincidentsRelatedItemsUpdate) > 0 {
		if err := assignInputField(input, "RelatedItemsUpdate", _ssmincidentsRelatedItemsUpdate); err != nil {
			log.Errorf("invalid --related-items-update: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsClientToken) > 0 {
		input.ClientToken = aws.String(_ssmincidentsClientToken)
	}

	if resp, err := client.UpdateRelatedItems(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Add or delete Regions from your replication set.
func ssmincidents_UpdateReplicationSet(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.UpdateReplicationSetInput{
		// Actions: []types.UpdateReplicationSetAction, // Required
		// Arn: *string, // Required
	}

	if len(_ssmincidentsActions) > 0 {
		if err := assignInputField(input, "Actions", _ssmincidentsActions); err != nil {
			log.Errorf("invalid --actions: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsArn) > 0 {
		input.Arn = aws.String(_ssmincidentsArn)
	}
	if len(_ssmincidentsClientToken) > 0 {
		input.ClientToken = aws.String(_ssmincidentsClientToken)
	}

	if resp, err := client.UpdateReplicationSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified response plan.
func ssmincidents_UpdateResponsePlan(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.UpdateResponsePlanInput{
		// Arn: *string, // Required
	}

	if len(_ssmincidentsArn) > 0 {
		input.Arn = aws.String(_ssmincidentsArn)
	}
	if len(_ssmincidentsActions) > 0 {
		if err := assignInputField(input, "Actions", _ssmincidentsActions); err != nil {
			log.Errorf("invalid --actions: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsChatChannel) > 0 {
		if err := assignInputField(input, "ChatChannel", _ssmincidentsChatChannel); err != nil {
			log.Errorf("invalid --chat-channel: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsClientToken) > 0 {
		input.ClientToken = aws.String(_ssmincidentsClientToken)
	}
	if len(_ssmincidentsDisplayName) > 0 {
		input.DisplayName = aws.String(_ssmincidentsDisplayName)
	}
	if len(_ssmincidentsEngagements) > 0 {
		input.Engagements = append([]string(nil), _ssmincidentsEngagements...)
	}
	if len(_ssmincidentsIncidentTemplateDedupeString) > 0 {
		input.IncidentTemplateDedupeString = aws.String(_ssmincidentsIncidentTemplateDedupeString)
	}
	if len(_ssmincidentsIncidentTemplateImpact) > 0 {
		if err := assignInputField(input, "IncidentTemplateImpact", _ssmincidentsIncidentTemplateImpact); err != nil {
			log.Errorf("invalid --incident-template-impact: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsIncidentTemplateNotificationTargets) > 0 {
		if err := assignInputField(input, "IncidentTemplateNotificationTargets", _ssmincidentsIncidentTemplateNotificationTargets); err != nil {
			log.Errorf("invalid --incident-template-notification-targets: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsIncidentTemplateSummary) > 0 {
		input.IncidentTemplateSummary = aws.String(_ssmincidentsIncidentTemplateSummary)
	}
	if len(_ssmincidentsIncidentTemplateTags) > 0 {
		if err := assignInputField(input, "IncidentTemplateTags", _ssmincidentsIncidentTemplateTags); err != nil {
			log.Errorf("invalid --incident-template-tags: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsIncidentTemplateTitle) > 0 {
		input.IncidentTemplateTitle = aws.String(_ssmincidentsIncidentTemplateTitle)
	}
	if len(_ssmincidentsIntegrations) > 0 {
		if err := assignInputField(input, "Integrations", _ssmincidentsIntegrations); err != nil {
			log.Errorf("invalid --integrations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateResponsePlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a timeline event. You can update events of type Custom Event .
func ssmincidents_UpdateTimelineEvent(cfg aws.Config, client *ssmincidents.Client) {
	input := &ssmincidents.UpdateTimelineEventInput{
		// EventId: *string, // Required
		// IncidentRecordArn: *string, // Required
	}

	if len(_ssmincidentsEventId) > 0 {
		input.EventId = aws.String(_ssmincidentsEventId)
	}
	if len(_ssmincidentsIncidentRecordArn) > 0 {
		input.IncidentRecordArn = aws.String(_ssmincidentsIncidentRecordArn)
	}
	if len(_ssmincidentsClientToken) > 0 {
		input.ClientToken = aws.String(_ssmincidentsClientToken)
	}
	if len(_ssmincidentsEventData) > 0 {
		input.EventData = aws.String(_ssmincidentsEventData)
	}
	if len(_ssmincidentsEventReferences) > 0 {
		if err := assignInputField(input, "EventReferences", _ssmincidentsEventReferences); err != nil {
			log.Errorf("invalid --event-references: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsEventTime) > 0 {
		if err := assignInputField(input, "EventTime", _ssmincidentsEventTime); err != nil {
			log.Errorf("invalid --event-time: %s", err.Error())
			return
		}
	}
	if len(_ssmincidentsEventType) > 0 {
		input.EventType = aws.String(_ssmincidentsEventType)
	}

	if resp, err := client.UpdateTimelineEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_ssmincidentsCmd)
	_ssmincidentsCmd.Flags().SortFlags = false

	_ssmincidentsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_ssmincidentsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_ssmincidentsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsActions, "actions", "", "", "Actions")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsArn, "arn", "", "", "ARN")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsChatChannel, "chat-channel", "", "", "Chat Channel")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsClientToken, "client-token", "", "", "Client Token")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsDeletionProtected, "deletion-protected", "", "", "Deletion Protected")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsDisplayName, "display-name", "", "", "Display Name")
	_ssmincidentsCmd.Flags().StringSliceVarP(&_ssmincidentsEngagements, "engagements", "", nil, "Engagements")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsEventData, "event-data", "", "", "Event Data")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsEventId, "event-id", "", "", "Event ID")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsEventReferences, "event-references", "", "", "Event References")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsEventTime, "event-time", "", "", "Event Time")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsEventType, "event-type", "", "", "Event Type")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsFilters, "filters", "", "", "Filters")
	_ssmincidentsCmd.Flags().StringSliceVarP(&_ssmincidentsFindingIds, "finding-ids", "", nil, "Finding Ids")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsImpact, "impact", "", "", "Impact")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsIncidentRecordArn, "incident-record-arn", "", "", "Incident Record ARN")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsIncidentTemplate, "incident-template", "", "", "Incident Template")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsIncidentTemplateDedupeString, "incident-template-dedupe-string", "", "", "Incident Template Dedupe String")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsIncidentTemplateImpact, "incident-template-impact", "", "", "Incident Template Impact")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsIncidentTemplateNotificationTargets, "incident-template-notification-targets", "", "", "Incident Template Notification Targets")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsIncidentTemplateSummary, "incident-template-summary", "", "", "Incident Template Summary")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsIncidentTemplateTags, "incident-template-tags", "", "", "Incident Template Tags")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsIncidentTemplateTitle, "incident-template-title", "", "", "Incident Template Title")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsIntegrations, "integrations", "", "", "Integrations")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsMaxResults, "max-results", "", "", "Max Results")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsName, "name", "", "", "Name")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsNextToken, "next-token", "", "", "Next Token")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsNotificationTargets, "notification-targets", "", "", "Notification Targets")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsPolicy, "policy", "", "", "Policy")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsPolicyId, "policy-id", "", "", "Policy ID")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsRegions, "regions", "", "", "Regions")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsRelatedItems, "related-items", "", "", "Related Items")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsRelatedItemsUpdate, "related-items-update", "", "", "Related Items Update")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsResourceArn, "resource-arn", "", "", "Resource ARN")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsResponsePlanArn, "response-plan-arn", "", "", "Response Plan ARN")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsSortBy, "sort-by", "", "", "Sort By")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsSortOrder, "sort-order", "", "", "Sort Order")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsStatus, "status", "", "", "Status")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsSummary, "summary", "", "", "Summary")
	_ssmincidentsCmd.Flags().StringSliceVarP(&_ssmincidentsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsTags, "tags", "", "", "Tags")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsTitle, "title", "", "", "Title")
	_ssmincidentsCmd.Flags().StringVarP(&_ssmincidentsTriggerDetails, "trigger-details", "", "", "Trigger Details")

	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsBatchGetIncidentFindings, "batch-get-incident-findings", "", false, "Batch Get Incident Findings")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsCreateReplicationSet, "create-replication-set", "", false, "Create Replication Set")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsCreateResponsePlan, "create-response-plan", "", false, "Create Response Plan")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsCreateTimelineEvent, "create-timeline-event", "", false, "Create Timeline Event")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsDeleteIncidentRecord, "delete-incident-record", "", false, "Delete Incident Record")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsDeleteReplicationSet, "delete-replication-set", "", false, "Delete Replication Set")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsDeleteResourcePolicy, "delete-resource-policy", "", false, "Delete Resource Policy")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsDeleteResponsePlan, "delete-response-plan", "", false, "Delete Response Plan")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsDeleteTimelineEvent, "delete-timeline-event", "", false, "Delete Timeline Event")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsGetIncidentRecord, "get-incident-record", "", false, "Get Incident Record")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsGetReplicationSet, "get-replication-set", "", false, "Get Replication Set")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsGetResourcePolicies, "get-resource-policies", "", false, "Get Resource Policies")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsGetResponsePlan, "get-response-plan", "", false, "Get Response Plan")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsGetTimelineEvent, "get-timeline-event", "", false, "Get Timeline Event")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsListIncidentFindings, "list-incident-findings", "", false, "List Incident Findings")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsListIncidentRecords, "list-incident-records", "", false, "List Incident Records")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsListRelatedItems, "list-related-items", "", false, "List Related Items")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsListReplicationSets, "list-replication-sets", "", false, "List Replication Sets")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsListResponsePlans, "list-response-plans", "", false, "List Response Plans")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsListTimelineEvents, "list-timeline-events", "", false, "List Timeline Events")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsPutResourcePolicy, "put-resource-policy", "", false, "Put Resource Policy")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsStartIncident, "start-incident", "", false, "Start Incident")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsTagResource, "tag-resource", "", false, "Tag Resource")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsUntagResource, "untag-resource", "", false, "Untag Resource")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsUpdateDeletionProtection, "update-deletion-protection", "", false, "Update Deletion Protection")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsUpdateIncidentRecord, "update-incident-record", "", false, "Update Incident Record")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsUpdateRelatedItems, "update-related-items", "", false, "Update Related Items")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsUpdateReplicationSet, "update-replication-set", "", false, "Update Replication Set")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsUpdateResponsePlan, "update-response-plan", "", false, "Update Response Plan")
	_ssmincidentsCmd.Flags().BoolVarP(&_ssmincidentsUpdateTimelineEvent, "update-timeline-event", "", false, "Update Timeline Event")

}
