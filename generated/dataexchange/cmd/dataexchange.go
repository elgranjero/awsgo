package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dataexchange"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// dataexchangeCmd represents the dataexchange command
var _dataexchangeCmd = &cobra.Command{
	Use:   "dataexchange",
	Short: "AWS dataexchange CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := dataexchange.NewFromConfig(cfg)
		if _dataexchangeAcceptDataGrant {
			dataexchange_AcceptDataGrant(cfg, client)
			return
		}
		if _dataexchangeCancelJob {
			dataexchange_CancelJob(cfg, client)
			return
		}
		if _dataexchangeCreateDataGrant {
			dataexchange_CreateDataGrant(cfg, client)
			return
		}
		if _dataexchangeCreateDataSet {
			dataexchange_CreateDataSet(cfg, client)
			return
		}
		if _dataexchangeCreateEventAction {
			dataexchange_CreateEventAction(cfg, client)
			return
		}
		if _dataexchangeCreateJob {
			dataexchange_CreateJob(cfg, client)
			return
		}
		if _dataexchangeCreateRevision {
			dataexchange_CreateRevision(cfg, client)
			return
		}
		if _dataexchangeDeleteAsset {
			dataexchange_DeleteAsset(cfg, client)
			return
		}
		if _dataexchangeDeleteDataGrant {
			dataexchange_DeleteDataGrant(cfg, client)
			return
		}
		if _dataexchangeDeleteDataSet {
			dataexchange_DeleteDataSet(cfg, client)
			return
		}
		if _dataexchangeDeleteEventAction {
			dataexchange_DeleteEventAction(cfg, client)
			return
		}
		if _dataexchangeDeleteRevision {
			dataexchange_DeleteRevision(cfg, client)
			return
		}
		if _dataexchangeGetAsset {
			dataexchange_GetAsset(cfg, client)
			return
		}
		if _dataexchangeGetDataGrant {
			dataexchange_GetDataGrant(cfg, client)
			return
		}
		if _dataexchangeGetDataSet {
			dataexchange_GetDataSet(cfg, client)
			return
		}
		if _dataexchangeGetEventAction {
			dataexchange_GetEventAction(cfg, client)
			return
		}
		if _dataexchangeGetJob {
			dataexchange_GetJob(cfg, client)
			return
		}
		if _dataexchangeGetReceivedDataGrant {
			dataexchange_GetReceivedDataGrant(cfg, client)
			return
		}
		if _dataexchangeGetRevision {
			dataexchange_GetRevision(cfg, client)
			return
		}
		if _dataexchangeListDataGrants {
			dataexchange_ListDataGrants(cfg, client)
			return
		}
		if _dataexchangeListDataSetRevisions {
			dataexchange_ListDataSetRevisions(cfg, client)
			return
		}
		if _dataexchangeListDataSets {
			dataexchange_ListDataSets(cfg, client)
			return
		}
		if _dataexchangeListEventActions {
			dataexchange_ListEventActions(cfg, client)
			return
		}
		if _dataexchangeListJobs {
			dataexchange_ListJobs(cfg, client)
			return
		}
		if _dataexchangeListReceivedDataGrants {
			dataexchange_ListReceivedDataGrants(cfg, client)
			return
		}
		if _dataexchangeListRevisionAssets {
			dataexchange_ListRevisionAssets(cfg, client)
			return
		}
		if _dataexchangeListTagsForResource {
			dataexchange_ListTagsForResource(cfg, client)
			return
		}
		if _dataexchangeRevokeRevision {
			dataexchange_RevokeRevision(cfg, client)
			return
		}
		if _dataexchangeSendApiAsset {
			dataexchange_SendApiAsset(cfg, client)
			return
		}
		if _dataexchangeSendDataSetNotification {
			dataexchange_SendDataSetNotification(cfg, client)
			return
		}
		if _dataexchangeStartJob {
			dataexchange_StartJob(cfg, client)
			return
		}
		if _dataexchangeTagResource {
			dataexchange_TagResource(cfg, client)
			return
		}
		if _dataexchangeUntagResource {
			dataexchange_UntagResource(cfg, client)
			return
		}
		if _dataexchangeUpdateAsset {
			dataexchange_UpdateAsset(cfg, client)
			return
		}
		if _dataexchangeUpdateDataSet {
			dataexchange_UpdateDataSet(cfg, client)
			return
		}
		if _dataexchangeUpdateEventAction {
			dataexchange_UpdateEventAction(cfg, client)
			return
		}
		if _dataexchangeUpdateRevision {
			dataexchange_UpdateRevision(cfg, client)
			return
		}

	},
}

var (
	_dataexchangeAcceptDataGrant         bool
	_dataexchangeCancelJob               bool
	_dataexchangeCreateDataGrant         bool
	_dataexchangeCreateDataSet           bool
	_dataexchangeCreateEventAction       bool
	_dataexchangeCreateJob               bool
	_dataexchangeCreateRevision          bool
	_dataexchangeDeleteAsset             bool
	_dataexchangeDeleteDataGrant         bool
	_dataexchangeDeleteDataSet           bool
	_dataexchangeDeleteEventAction       bool
	_dataexchangeDeleteRevision          bool
	_dataexchangeGetAsset                bool
	_dataexchangeGetDataGrant            bool
	_dataexchangeGetDataSet              bool
	_dataexchangeGetEventAction          bool
	_dataexchangeGetJob                  bool
	_dataexchangeGetReceivedDataGrant    bool
	_dataexchangeGetRevision             bool
	_dataexchangeListDataGrants          bool
	_dataexchangeListDataSetRevisions    bool
	_dataexchangeListDataSets            bool
	_dataexchangeListEventActions        bool
	_dataexchangeListJobs                bool
	_dataexchangeListReceivedDataGrants  bool
	_dataexchangeListRevisionAssets      bool
	_dataexchangeListTagsForResource     bool
	_dataexchangeRevokeRevision          bool
	_dataexchangeSendApiAsset            bool
	_dataexchangeSendDataSetNotification bool
	_dataexchangeStartJob                bool
	_dataexchangeTagResource             bool
	_dataexchangeUntagResource           bool
	_dataexchangeUpdateAsset             bool
	_dataexchangeUpdateDataSet           bool
	_dataexchangeUpdateEventAction       bool
	_dataexchangeUpdateRevision          bool

	_dataexchangeAcceptanceState        string
	_dataexchangeAction                 string
	_dataexchangeAssetId                string
	_dataexchangeAssetType              string
	_dataexchangeBody                   string
	_dataexchangeClientToken            string
	_dataexchangeComment                string
	_dataexchangeDataGrantArn           string
	_dataexchangeDataGrantId            string
	_dataexchangeDataSetId              string
	_dataexchangeDescription            string
	_dataexchangeDetails                string
	_dataexchangeEndsAt                 string
	_dataexchangeEvent                  string
	_dataexchangeEventActionId          string
	_dataexchangeEventSourceId          string
	_dataexchangeFinalized              string
	_dataexchangeGrantDistributionScope string
	_dataexchangeJobId                  string
	_dataexchangeMaxResults             string
	_dataexchangeMethod                 string
	_dataexchangeName                   string
	_dataexchangeNextToken              string
	_dataexchangeOrigin                 string
	_dataexchangePath                   string
	_dataexchangeQueryStringParameters  string
	_dataexchangeReceiverPrincipal      string
	_dataexchangeRequestHeaders         string
	_dataexchangeResourceArn            string
	_dataexchangeRevisionId             string
	_dataexchangeRevocationComment      string
	_dataexchangeScope                  string
	_dataexchangeSourceDataSetId        string
	_dataexchangeTagKeys                []string
	_dataexchangeTags                   string
	_dataexchangeType                   string
)

// This operation accepts a data grant.
func dataexchange_AcceptDataGrant(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.AcceptDataGrantInput{
		// DataGrantArn: *string, // Required
	}

	if len(_dataexchangeDataGrantArn) > 0 {
		input.DataGrantArn = aws.String(_dataexchangeDataGrantArn)
	}

	if resp, err := client.AcceptDataGrant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation cancels a job. Jobs can be cancelled only when they are in the
// WAITING state.
func dataexchange_CancelJob(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.CancelJobInput{
		// JobId: *string, // Required
	}

	if len(_dataexchangeJobId) > 0 {
		input.JobId = aws.String(_dataexchangeJobId)
	}

	if resp, err := client.CancelJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation creates a data grant.
func dataexchange_CreateDataGrant(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.CreateDataGrantInput{
		// GrantDistributionScope: types.GrantDistributionScope, // Required
		// Name: *string, // Required
		// ReceiverPrincipal: *string, // Required
		// SourceDataSetId: *string, // Required
	}

	if len(_dataexchangeGrantDistributionScope) > 0 {
		if err := assignInputField(input, "GrantDistributionScope", _dataexchangeGrantDistributionScope); err != nil {
			log.Errorf("invalid --grant-distribution-scope: %s", err.Error())
			return
		}
	}
	if len(_dataexchangeName) > 0 {
		input.Name = aws.String(_dataexchangeName)
	}
	if len(_dataexchangeReceiverPrincipal) > 0 {
		input.ReceiverPrincipal = aws.String(_dataexchangeReceiverPrincipal)
	}
	if len(_dataexchangeSourceDataSetId) > 0 {
		input.SourceDataSetId = aws.String(_dataexchangeSourceDataSetId)
	}
	if len(_dataexchangeDescription) > 0 {
		input.Description = aws.String(_dataexchangeDescription)
	}
	if len(_dataexchangeEndsAt) > 0 {
		if err := assignInputField(input, "EndsAt", _dataexchangeEndsAt); err != nil {
			log.Errorf("invalid --ends-at: %s", err.Error())
			return
		}
	}
	if len(_dataexchangeTags) > 0 {
		if err := assignInputField(input, "Tags", _dataexchangeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataGrant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation creates a data set.
func dataexchange_CreateDataSet(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.CreateDataSetInput{
		// AssetType: types.AssetType, // Required
		// Description: *string, // Required
		// Name: *string, // Required
	}

	if len(_dataexchangeAssetType) > 0 {
		if err := assignInputField(input, "AssetType", _dataexchangeAssetType); err != nil {
			log.Errorf("invalid --asset-type: %s", err.Error())
			return
		}
	}
	if len(_dataexchangeDescription) > 0 {
		input.Description = aws.String(_dataexchangeDescription)
	}
	if len(_dataexchangeName) > 0 {
		input.Name = aws.String(_dataexchangeName)
	}
	if len(_dataexchangeTags) > 0 {
		if err := assignInputField(input, "Tags", _dataexchangeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation creates an event action.
func dataexchange_CreateEventAction(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.CreateEventActionInput{
		// Action: *types.Action, // Required
		// Event: *types.Event, // Required
	}

	if len(_dataexchangeAction) > 0 {
		if err := assignInputField(input, "Action", _dataexchangeAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_dataexchangeEvent) > 0 {
		if err := assignInputField(input, "Event", _dataexchangeEvent); err != nil {
			log.Errorf("invalid --event: %s", err.Error())
			return
		}
	}
	if len(_dataexchangeTags) > 0 {
		if err := assignInputField(input, "Tags", _dataexchangeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEventAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation creates a job.
func dataexchange_CreateJob(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.CreateJobInput{
		// Details: *types.RequestDetails, // Required
		// Type: types.Type, // Required
	}

	if len(_dataexchangeDetails) > 0 {
		if err := assignInputField(input, "Details", _dataexchangeDetails); err != nil {
			log.Errorf("invalid --details: %s", err.Error())
			return
		}
	}
	if len(_dataexchangeType) > 0 {
		if err := assignInputField(input, "Type", _dataexchangeType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
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

// This operation creates a revision for a data set.
func dataexchange_CreateRevision(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.CreateRevisionInput{
		// DataSetId: *string, // Required
	}

	if len(_dataexchangeDataSetId) > 0 {
		input.DataSetId = aws.String(_dataexchangeDataSetId)
	}
	if len(_dataexchangeComment) > 0 {
		input.Comment = aws.String(_dataexchangeComment)
	}
	if len(_dataexchangeTags) > 0 {
		if err := assignInputField(input, "Tags", _dataexchangeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRevision(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation deletes an asset.
func dataexchange_DeleteAsset(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.DeleteAssetInput{
		// AssetId: *string, // Required
		// DataSetId: *string, // Required
		// RevisionId: *string, // Required
	}

	if len(_dataexchangeAssetId) > 0 {
		input.AssetId = aws.String(_dataexchangeAssetId)
	}
	if len(_dataexchangeDataSetId) > 0 {
		input.DataSetId = aws.String(_dataexchangeDataSetId)
	}
	if len(_dataexchangeRevisionId) > 0 {
		input.RevisionId = aws.String(_dataexchangeRevisionId)
	}

	if resp, err := client.DeleteAsset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation deletes a data grant.
func dataexchange_DeleteDataGrant(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.DeleteDataGrantInput{
		// DataGrantId: *string, // Required
	}

	if len(_dataexchangeDataGrantId) > 0 {
		input.DataGrantId = aws.String(_dataexchangeDataGrantId)
	}

	if resp, err := client.DeleteDataGrant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation deletes a data set.
func dataexchange_DeleteDataSet(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.DeleteDataSetInput{
		// DataSetId: *string, // Required
	}

	if len(_dataexchangeDataSetId) > 0 {
		input.DataSetId = aws.String(_dataexchangeDataSetId)
	}

	if resp, err := client.DeleteDataSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation deletes the event action.
func dataexchange_DeleteEventAction(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.DeleteEventActionInput{
		// EventActionId: *string, // Required
	}

	if len(_dataexchangeEventActionId) > 0 {
		input.EventActionId = aws.String(_dataexchangeEventActionId)
	}

	if resp, err := client.DeleteEventAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation deletes a revision.
func dataexchange_DeleteRevision(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.DeleteRevisionInput{
		// DataSetId: *string, // Required
		// RevisionId: *string, // Required
	}

	if len(_dataexchangeDataSetId) > 0 {
		input.DataSetId = aws.String(_dataexchangeDataSetId)
	}
	if len(_dataexchangeRevisionId) > 0 {
		input.RevisionId = aws.String(_dataexchangeRevisionId)
	}

	if resp, err := client.DeleteRevision(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation returns information about an asset.
func dataexchange_GetAsset(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.GetAssetInput{
		// AssetId: *string, // Required
		// DataSetId: *string, // Required
		// RevisionId: *string, // Required
	}

	if len(_dataexchangeAssetId) > 0 {
		input.AssetId = aws.String(_dataexchangeAssetId)
	}
	if len(_dataexchangeDataSetId) > 0 {
		input.DataSetId = aws.String(_dataexchangeDataSetId)
	}
	if len(_dataexchangeRevisionId) > 0 {
		input.RevisionId = aws.String(_dataexchangeRevisionId)
	}

	if resp, err := client.GetAsset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation returns information about a data grant.
func dataexchange_GetDataGrant(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.GetDataGrantInput{
		// DataGrantId: *string, // Required
	}

	if len(_dataexchangeDataGrantId) > 0 {
		input.DataGrantId = aws.String(_dataexchangeDataGrantId)
	}

	if resp, err := client.GetDataGrant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation returns information about a data set.
func dataexchange_GetDataSet(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.GetDataSetInput{
		// DataSetId: *string, // Required
	}

	if len(_dataexchangeDataSetId) > 0 {
		input.DataSetId = aws.String(_dataexchangeDataSetId)
	}

	if resp, err := client.GetDataSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation retrieves information about an event action.
func dataexchange_GetEventAction(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.GetEventActionInput{
		// EventActionId: *string, // Required
	}

	if len(_dataexchangeEventActionId) > 0 {
		input.EventActionId = aws.String(_dataexchangeEventActionId)
	}

	if resp, err := client.GetEventAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation returns information about a job.
func dataexchange_GetJob(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.GetJobInput{
		// JobId: *string, // Required
	}

	if len(_dataexchangeJobId) > 0 {
		input.JobId = aws.String(_dataexchangeJobId)
	}

	if resp, err := client.GetJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation returns information about a received data grant.
func dataexchange_GetReceivedDataGrant(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.GetReceivedDataGrantInput{
		// DataGrantArn: *string, // Required
	}

	if len(_dataexchangeDataGrantArn) > 0 {
		input.DataGrantArn = aws.String(_dataexchangeDataGrantArn)
	}

	if resp, err := client.GetReceivedDataGrant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation returns information about a revision.
func dataexchange_GetRevision(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.GetRevisionInput{
		// DataSetId: *string, // Required
		// RevisionId: *string, // Required
	}

	if len(_dataexchangeDataSetId) > 0 {
		input.DataSetId = aws.String(_dataexchangeDataSetId)
	}
	if len(_dataexchangeRevisionId) > 0 {
		input.RevisionId = aws.String(_dataexchangeRevisionId)
	}

	if resp, err := client.GetRevision(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation returns information about all data grants.
func dataexchange_ListDataGrants(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.ListDataGrantsInput{}

	if len(_dataexchangeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _dataexchangeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_dataexchangeNextToken) > 0 {
		input.NextToken = aws.String(_dataexchangeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDataGrants(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*dataexchange.ListDataGrantsOutput
	p := dataexchange.NewListDataGrantsPaginator(client, input)
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

// This operation lists a data set's revisions sorted by CreatedAt in descending
// order.
func dataexchange_ListDataSetRevisions(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.ListDataSetRevisionsInput{
		// DataSetId: *string, // Required
	}

	if len(_dataexchangeDataSetId) > 0 {
		input.DataSetId = aws.String(_dataexchangeDataSetId)
	}
	if len(_dataexchangeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _dataexchangeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_dataexchangeNextToken) > 0 {
		input.NextToken = aws.String(_dataexchangeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDataSetRevisions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*dataexchange.ListDataSetRevisionsOutput
	p := dataexchange.NewListDataSetRevisionsPaginator(client, input)
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

// This operation lists your data sets. When listing by origin OWNED, results are
// sorted by CreatedAt in descending order. When listing by origin ENTITLED, there
// is no order.
func dataexchange_ListDataSets(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.ListDataSetsInput{}

	if len(_dataexchangeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _dataexchangeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_dataexchangeNextToken) > 0 {
		input.NextToken = aws.String(_dataexchangeNextToken)
	}
	if len(_dataexchangeOrigin) > 0 {
		input.Origin = aws.String(_dataexchangeOrigin)
	}

	if disablePaginator() {
		if resp, err := client.ListDataSets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*dataexchange.ListDataSetsOutput
	p := dataexchange.NewListDataSetsPaginator(client, input)
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

// This operation lists your event actions.
func dataexchange_ListEventActions(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.ListEventActionsInput{}

	if len(_dataexchangeEventSourceId) > 0 {
		input.EventSourceId = aws.String(_dataexchangeEventSourceId)
	}
	if len(_dataexchangeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _dataexchangeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_dataexchangeNextToken) > 0 {
		input.NextToken = aws.String(_dataexchangeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEventActions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*dataexchange.ListEventActionsOutput
	p := dataexchange.NewListEventActionsPaginator(client, input)
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

// This operation lists your jobs sorted by CreatedAt in descending order.
func dataexchange_ListJobs(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.ListJobsInput{}

	if len(_dataexchangeDataSetId) > 0 {
		input.DataSetId = aws.String(_dataexchangeDataSetId)
	}
	if len(_dataexchangeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _dataexchangeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_dataexchangeNextToken) > 0 {
		input.NextToken = aws.String(_dataexchangeNextToken)
	}
	if len(_dataexchangeRevisionId) > 0 {
		input.RevisionId = aws.String(_dataexchangeRevisionId)
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

	var results []*dataexchange.ListJobsOutput
	p := dataexchange.NewListJobsPaginator(client, input)
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

// This operation returns information about all received data grants.
func dataexchange_ListReceivedDataGrants(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.ListReceivedDataGrantsInput{}

	if len(_dataexchangeAcceptanceState) > 0 {
		if err := assignInputField(input, "AcceptanceState", _dataexchangeAcceptanceState); err != nil {
			log.Errorf("invalid --acceptance-state: %s", err.Error())
			return
		}
	}
	if len(_dataexchangeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _dataexchangeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_dataexchangeNextToken) > 0 {
		input.NextToken = aws.String(_dataexchangeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListReceivedDataGrants(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*dataexchange.ListReceivedDataGrantsOutput
	p := dataexchange.NewListReceivedDataGrantsPaginator(client, input)
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

// This operation lists a revision's assets sorted alphabetically in descending
// order.
func dataexchange_ListRevisionAssets(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.ListRevisionAssetsInput{
		// DataSetId: *string, // Required
		// RevisionId: *string, // Required
	}

	if len(_dataexchangeDataSetId) > 0 {
		input.DataSetId = aws.String(_dataexchangeDataSetId)
	}
	if len(_dataexchangeRevisionId) > 0 {
		input.RevisionId = aws.String(_dataexchangeRevisionId)
	}
	if len(_dataexchangeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _dataexchangeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_dataexchangeNextToken) > 0 {
		input.NextToken = aws.String(_dataexchangeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRevisionAssets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*dataexchange.ListRevisionAssetsOutput
	p := dataexchange.NewListRevisionAssetsPaginator(client, input)
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

// This operation lists the tags on the resource.
func dataexchange_ListTagsForResource(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_dataexchangeResourceArn) > 0 {
		input.ResourceArn = aws.String(_dataexchangeResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation revokes subscribers' access to a revision.
func dataexchange_RevokeRevision(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.RevokeRevisionInput{
		// DataSetId: *string, // Required
		// RevisionId: *string, // Required
		// RevocationComment: *string, // Required
	}

	if len(_dataexchangeDataSetId) > 0 {
		input.DataSetId = aws.String(_dataexchangeDataSetId)
	}
	if len(_dataexchangeRevisionId) > 0 {
		input.RevisionId = aws.String(_dataexchangeRevisionId)
	}
	if len(_dataexchangeRevocationComment) > 0 {
		input.RevocationComment = aws.String(_dataexchangeRevocationComment)
	}

	if resp, err := client.RevokeRevision(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation invokes an API Gateway API asset. The request is proxied to the
// provider’s API Gateway API.
func dataexchange_SendApiAsset(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.SendApiAssetInput{
		// AssetId: *string, // Required
		// DataSetId: *string, // Required
		// RevisionId: *string, // Required
	}

	if len(_dataexchangeAssetId) > 0 {
		input.AssetId = aws.String(_dataexchangeAssetId)
	}
	if len(_dataexchangeDataSetId) > 0 {
		input.DataSetId = aws.String(_dataexchangeDataSetId)
	}
	if len(_dataexchangeRevisionId) > 0 {
		input.RevisionId = aws.String(_dataexchangeRevisionId)
	}
	if len(_dataexchangeBody) > 0 {
		input.Body = aws.String(_dataexchangeBody)
	}
	if len(_dataexchangeMethod) > 0 {
		input.Method = aws.String(_dataexchangeMethod)
	}
	if len(_dataexchangePath) > 0 {
		input.Path = aws.String(_dataexchangePath)
	}
	if len(_dataexchangeQueryStringParameters) > 0 {
		if err := assignInputField(input, "QueryStringParameters", _dataexchangeQueryStringParameters); err != nil {
			log.Errorf("invalid --query-string-parameters: %s", err.Error())
			return
		}
	}
	if len(_dataexchangeRequestHeaders) > 0 {
		if err := assignInputField(input, "RequestHeaders", _dataexchangeRequestHeaders); err != nil {
			log.Errorf("invalid --request-headers: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendApiAsset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The type of event associated with the data set.
func dataexchange_SendDataSetNotification(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.SendDataSetNotificationInput{
		// DataSetId: *string, // Required
		// Type: types.NotificationType, // Required
	}

	if len(_dataexchangeDataSetId) > 0 {
		input.DataSetId = aws.String(_dataexchangeDataSetId)
	}
	if len(_dataexchangeType) > 0 {
		if err := assignInputField(input, "Type", _dataexchangeType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_dataexchangeClientToken) > 0 {
		input.ClientToken = aws.String(_dataexchangeClientToken)
	}
	if len(_dataexchangeComment) > 0 {
		input.Comment = aws.String(_dataexchangeComment)
	}
	if len(_dataexchangeDetails) > 0 {
		if err := assignInputField(input, "Details", _dataexchangeDetails); err != nil {
			log.Errorf("invalid --details: %s", err.Error())
			return
		}
	}
	if len(_dataexchangeScope) > 0 {
		if err := assignInputField(input, "Scope", _dataexchangeScope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendDataSetNotification(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation starts a job.
func dataexchange_StartJob(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.StartJobInput{
		// JobId: *string, // Required
	}

	if len(_dataexchangeJobId) > 0 {
		input.JobId = aws.String(_dataexchangeJobId)
	}

	if resp, err := client.StartJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation tags a resource.
func dataexchange_TagResource(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_dataexchangeResourceArn) > 0 {
		input.ResourceArn = aws.String(_dataexchangeResourceArn)
	}
	if len(_dataexchangeTags) > 0 {
		if err := assignInputField(input, "Tags", _dataexchangeTags); err != nil {
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

// This operation removes one or more tags from a resource.
func dataexchange_UntagResource(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_dataexchangeResourceArn) > 0 {
		input.ResourceArn = aws.String(_dataexchangeResourceArn)
	}
	if len(_dataexchangeTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _dataexchangeTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation updates an asset.
func dataexchange_UpdateAsset(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.UpdateAssetInput{
		// AssetId: *string, // Required
		// DataSetId: *string, // Required
		// Name: *string, // Required
		// RevisionId: *string, // Required
	}

	if len(_dataexchangeAssetId) > 0 {
		input.AssetId = aws.String(_dataexchangeAssetId)
	}
	if len(_dataexchangeDataSetId) > 0 {
		input.DataSetId = aws.String(_dataexchangeDataSetId)
	}
	if len(_dataexchangeName) > 0 {
		input.Name = aws.String(_dataexchangeName)
	}
	if len(_dataexchangeRevisionId) > 0 {
		input.RevisionId = aws.String(_dataexchangeRevisionId)
	}

	if resp, err := client.UpdateAsset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation updates a data set.
func dataexchange_UpdateDataSet(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.UpdateDataSetInput{
		// DataSetId: *string, // Required
	}

	if len(_dataexchangeDataSetId) > 0 {
		input.DataSetId = aws.String(_dataexchangeDataSetId)
	}
	if len(_dataexchangeDescription) > 0 {
		input.Description = aws.String(_dataexchangeDescription)
	}
	if len(_dataexchangeName) > 0 {
		input.Name = aws.String(_dataexchangeName)
	}

	if resp, err := client.UpdateDataSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation updates the event action.
func dataexchange_UpdateEventAction(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.UpdateEventActionInput{
		// EventActionId: *string, // Required
	}

	if len(_dataexchangeEventActionId) > 0 {
		input.EventActionId = aws.String(_dataexchangeEventActionId)
	}
	if len(_dataexchangeAction) > 0 {
		if err := assignInputField(input, "Action", _dataexchangeAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEventAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation updates a revision.
func dataexchange_UpdateRevision(cfg aws.Config, client *dataexchange.Client) {
	input := &dataexchange.UpdateRevisionInput{
		// DataSetId: *string, // Required
		// RevisionId: *string, // Required
	}

	if len(_dataexchangeDataSetId) > 0 {
		input.DataSetId = aws.String(_dataexchangeDataSetId)
	}
	if len(_dataexchangeRevisionId) > 0 {
		input.RevisionId = aws.String(_dataexchangeRevisionId)
	}
	if len(_dataexchangeComment) > 0 {
		input.Comment = aws.String(_dataexchangeComment)
	}
	if len(_dataexchangeFinalized) > 0 {
		if err := assignInputField(input, "Finalized", _dataexchangeFinalized); err != nil {
			log.Errorf("invalid --finalized: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRevision(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_dataexchangeCmd)
	_dataexchangeCmd.Flags().SortFlags = false

	_dataexchangeCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_dataexchangeCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_dataexchangeCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeAcceptanceState, "acceptance-state", "", "", "Acceptance State")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeAction, "action", "", "", "Action")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeAssetId, "asset-id", "", "", "Asset ID")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeAssetType, "asset-type", "", "", "Asset Type")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeBody, "body", "", "", "Body")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeClientToken, "client-token", "", "", "Client Token")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeComment, "comment", "", "", "Comment")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeDataGrantArn, "data-grant-arn", "", "", "Data Grant ARN")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeDataGrantId, "data-grant-id", "", "", "Data Grant ID")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeDataSetId, "data-set-id", "", "", "Data Set ID")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeDescription, "description", "", "", "Description")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeDetails, "details", "", "", "Details")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeEndsAt, "ends-at", "", "", "Ends At")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeEvent, "event", "", "", "Event")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeEventActionId, "event-action-id", "", "", "Event Action ID")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeEventSourceId, "event-source-id", "", "", "Event Source ID")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeFinalized, "finalized", "", "", "Finalized")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeGrantDistributionScope, "grant-distribution-scope", "", "", "Grant Distribution Scope")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeJobId, "job-id", "", "", "Job ID")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeMaxResults, "max-results", "", "", "Max Results")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeMethod, "method", "", "", "Method")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeName, "name", "", "", "Name")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeNextToken, "next-token", "", "", "Next Token")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeOrigin, "origin", "", "", "Origin")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangePath, "path", "", "", "Path")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeQueryStringParameters, "query-string-parameters", "", "", "Query String Parameters")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeReceiverPrincipal, "receiver-principal", "", "", "Receiver Principal")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeRequestHeaders, "request-headers", "", "", "Request Headers")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeResourceArn, "resource-arn", "", "", "Resource ARN")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeRevisionId, "revision-id", "", "", "Revision ID")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeRevocationComment, "revocation-comment", "", "", "Revocation Comment")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeScope, "scope", "", "", "Scope")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeSourceDataSetId, "source-data-set-id", "", "", "Source Data Set ID")
	_dataexchangeCmd.Flags().StringSliceVarP(&_dataexchangeTagKeys, "tag-keys", "", nil, "Tag Keys")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeTags, "tags", "", "", "Tags")
	_dataexchangeCmd.Flags().StringVarP(&_dataexchangeType, "type", "", "", "Type")

	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeAcceptDataGrant, "accept-data-grant", "", false, "Accept Data Grant")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeCancelJob, "cancel-job", "", false, "Cancel Job")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeCreateDataGrant, "create-data-grant", "", false, "Create Data Grant")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeCreateDataSet, "create-data-set", "", false, "Create Data Set")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeCreateEventAction, "create-event-action", "", false, "Create Event Action")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeCreateJob, "create-job", "", false, "Create Job")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeCreateRevision, "create-revision", "", false, "Create Revision")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeDeleteAsset, "delete-asset", "", false, "Delete Asset")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeDeleteDataGrant, "delete-data-grant", "", false, "Delete Data Grant")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeDeleteDataSet, "delete-data-set", "", false, "Delete Data Set")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeDeleteEventAction, "delete-event-action", "", false, "Delete Event Action")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeDeleteRevision, "delete-revision", "", false, "Delete Revision")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeGetAsset, "get-asset", "", false, "Get Asset")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeGetDataGrant, "get-data-grant", "", false, "Get Data Grant")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeGetDataSet, "get-data-set", "", false, "Get Data Set")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeGetEventAction, "get-event-action", "", false, "Get Event Action")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeGetJob, "get-job", "", false, "Get Job")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeGetReceivedDataGrant, "get-received-data-grant", "", false, "Get Received Data Grant")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeGetRevision, "get-revision", "", false, "Get Revision")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeListDataGrants, "list-data-grants", "", false, "List Data Grants")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeListDataSetRevisions, "list-data-set-revisions", "", false, "List Data Set Revisions")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeListDataSets, "list-data-sets", "", false, "List Data Sets")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeListEventActions, "list-event-actions", "", false, "List Event Actions")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeListJobs, "list-jobs", "", false, "List Jobs")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeListReceivedDataGrants, "list-received-data-grants", "", false, "List Received Data Grants")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeListRevisionAssets, "list-revision-assets", "", false, "List Revision Assets")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeRevokeRevision, "revoke-revision", "", false, "Revoke Revision")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeSendApiAsset, "send-api-asset", "", false, "Send API Asset")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeSendDataSetNotification, "send-data-set-notification", "", false, "Send Data Set Notification")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeStartJob, "start-job", "", false, "Start Job")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeTagResource, "tag-resource", "", false, "Tag Resource")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeUntagResource, "untag-resource", "", false, "Untag Resource")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeUpdateAsset, "update-asset", "", false, "Update Asset")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeUpdateDataSet, "update-data-set", "", false, "Update Data Set")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeUpdateEventAction, "update-event-action", "", false, "Update Event Action")
	_dataexchangeCmd.Flags().BoolVarP(&_dataexchangeUpdateRevision, "update-revision", "", false, "Update Revision")

}
