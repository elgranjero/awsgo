package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/health"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// healthCmd represents the health command
var _healthCmd = &cobra.Command{
	Use:   "health",
	Short: "AWS health CLI",
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
		client := health.NewFromConfig(cfg)
		if _healthDescribeAffectedAccountsForOrganization {
			health_DescribeAffectedAccountsForOrganization(cfg, client)
			return
		}
		if _healthDescribeAffectedEntities {
			health_DescribeAffectedEntities(cfg, client)
			return
		}
		if _healthDescribeAffectedEntitiesForOrganization {
			health_DescribeAffectedEntitiesForOrganization(cfg, client)
			return
		}
		if _healthDescribeEntityAggregates {
			health_DescribeEntityAggregates(cfg, client)
			return
		}
		if _healthDescribeEntityAggregatesForOrganization {
			health_DescribeEntityAggregatesForOrganization(cfg, client)
			return
		}
		if _healthDescribeEventAggregates {
			health_DescribeEventAggregates(cfg, client)
			return
		}
		if _healthDescribeEventDetails {
			health_DescribeEventDetails(cfg, client)
			return
		}
		if _healthDescribeEventDetailsForOrganization {
			health_DescribeEventDetailsForOrganization(cfg, client)
			return
		}
		if _healthDescribeEventTypes {
			health_DescribeEventTypes(cfg, client)
			return
		}
		if _healthDescribeEvents {
			health_DescribeEvents(cfg, client)
			return
		}
		if _healthDescribeEventsForOrganization {
			health_DescribeEventsForOrganization(cfg, client)
			return
		}
		if _healthDescribeHealthServiceStatusForOrganization {
			health_DescribeHealthServiceStatusForOrganization(cfg, client)
			return
		}
		if _healthDisableHealthServiceAccessForOrganization {
			health_DisableHealthServiceAccessForOrganization(cfg, client)
			return
		}
		if _healthEnableHealthServiceAccessForOrganization {
			health_EnableHealthServiceAccessForOrganization(cfg, client)
			return
		}

	},
}

var (
	_healthDescribeAffectedAccountsForOrganization    bool
	_healthDescribeAffectedEntities                   bool
	_healthDescribeAffectedEntitiesForOrganization    bool
	_healthDescribeEntityAggregates                   bool
	_healthDescribeEntityAggregatesForOrganization    bool
	_healthDescribeEventAggregates                    bool
	_healthDescribeEventDetails                       bool
	_healthDescribeEventDetailsForOrganization        bool
	_healthDescribeEventTypes                         bool
	_healthDescribeEvents                             bool
	_healthDescribeEventsForOrganization              bool
	_healthDescribeHealthServiceStatusForOrganization bool
	_healthDisableHealthServiceAccessForOrganization  bool
	_healthEnableHealthServiceAccessForOrganization   bool

	_healthAggregateField                   string
	_healthAwsAccountIds                    []string
	_healthEventArn                         string
	_healthEventArns                        []string
	_healthFilter                           string
	_healthLocale                           string
	_healthMaxResults                       string
	_healthNextToken                        string
	_healthOrganizationEntityAccountFilters string
	_healthOrganizationEntityFilters        string
	_healthOrganizationEventDetailFilters   string
)

// Returns a list of accounts in the organization from Organizations that are
// affected by the provided event. For more information about the different types
// of Health events, see [Event].
//
// Before you can call this operation, you must first enable Health to work with
// Organizations. To do this, call the [EnableHealthServiceAccessForOrganization]operation from your organization's
// management account.
//
// This API operation uses pagination. Specify the nextToken parameter in the next
// request to return more results.
//
// [Event]: https://docs.aws.amazon.com/health/latest/APIReference/API_Event.html
// [EnableHealthServiceAccessForOrganization]: https://docs.aws.amazon.com/health/latest/APIReference/API_EnableHealthServiceAccessForOrganization.html
func health_DescribeAffectedAccountsForOrganization(cfg aws.Config, client *health.Client) {
	input := &health.DescribeAffectedAccountsForOrganizationInput{
		// EventArn: *string, // Required
	}

	if len(_healthEventArn) > 0 {
		input.EventArn = aws.String(_healthEventArn)
	}
	if len(_healthMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _healthMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_healthNextToken) > 0 {
		input.NextToken = aws.String(_healthNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeAffectedAccountsForOrganization(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*health.DescribeAffectedAccountsForOrganizationOutput
	p := health.NewDescribeAffectedAccountsForOrganizationPaginator(client, input)
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

// Returns a list of entities that have been affected by the specified events,
// based on the specified filter criteria. Entities can refer to individual
// customer resources, groups of customer resources, or any other construct,
// depending on the Amazon Web Services service. Events that have impact beyond
// that of the affected entities, or where the extent of impact is unknown, include
// at least one entity indicating this.
//
// At least one event ARN is required.
//
// - This API operation uses pagination. Specify the nextToken parameter in the
// next request to return more results.
//
// - This operation supports resource-level permissions. You can use this
// operation to allow or deny access to specific Health events. For more
// information, see [Resource- and action-based conditions]in the Health User Guide.
//
// [Resource- and action-based conditions]: https://docs.aws.amazon.com/health/latest/ug/security_iam_id-based-policy-examples.html#resource-action-based-conditions
func health_DescribeAffectedEntities(cfg aws.Config, client *health.Client) {
	input := &health.DescribeAffectedEntitiesInput{
		// Filter: *types.EntityFilter, // Required
	}

	if len(_healthFilter) > 0 {
		if err := assignInputField(input, "Filter", _healthFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_healthLocale) > 0 {
		input.Locale = aws.String(_healthLocale)
	}
	if len(_healthMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _healthMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_healthNextToken) > 0 {
		input.NextToken = aws.String(_healthNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeAffectedEntities(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*health.DescribeAffectedEntitiesOutput
	p := health.NewDescribeAffectedEntitiesPaginator(client, input)
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

// Returns a list of entities that have been affected by one or more events for
// one or more accounts in your organization in Organizations, based on the filter
// criteria. Entities can refer to individual customer resources, groups of
// customer resources, or any other construct, depending on the Amazon Web Services
// service.
//
// At least one event Amazon Resource Name (ARN) and account ID are required.
//
// Before you can call this operation, you must first enable Health to work with
// Organizations. To do this, call the [EnableHealthServiceAccessForOrganization]operation from your organization's
// management account.
//
// - This API operation uses pagination. Specify the nextToken parameter in the
// next request to return more results.
//
// - This operation doesn't support resource-level permissions. You can't use
// this operation to allow or deny access to specific Health events. For more
// information, see [Resource- and action-based conditions]in the Health User Guide.
//
// [Resource- and action-based conditions]: https://docs.aws.amazon.com/health/latest/ug/security_iam_id-based-policy-examples.html#resource-action-based-conditions
// [EnableHealthServiceAccessForOrganization]: https://docs.aws.amazon.com/health/latest/APIReference/API_EnableHealthServiceAccessForOrganization.html
func health_DescribeAffectedEntitiesForOrganization(cfg aws.Config, client *health.Client) {
	input := &health.DescribeAffectedEntitiesForOrganizationInput{}

	if len(_healthLocale) > 0 {
		input.Locale = aws.String(_healthLocale)
	}
	if len(_healthMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _healthMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_healthNextToken) > 0 {
		input.NextToken = aws.String(_healthNextToken)
	}
	if len(_healthOrganizationEntityAccountFilters) > 0 {
		if err := assignInputField(input, "OrganizationEntityAccountFilters", _healthOrganizationEntityAccountFilters); err != nil {
			log.Errorf("invalid --organization-entity-account-filters: %s", err.Error())
			return
		}
	}
	if len(_healthOrganizationEntityFilters) > 0 {
		if err := assignInputField(input, "OrganizationEntityFilters", _healthOrganizationEntityFilters); err != nil {
			log.Errorf("invalid --organization-entity-filters: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeAffectedEntitiesForOrganization(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*health.DescribeAffectedEntitiesForOrganizationOutput
	p := health.NewDescribeAffectedEntitiesForOrganizationPaginator(client, input)
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

// Returns the number of entities that are affected by each of the specified
// events.
func health_DescribeEntityAggregates(cfg aws.Config, client *health.Client) {
	input := &health.DescribeEntityAggregatesInput{}

	if len(_healthEventArns) > 0 {
		input.EventArns = append([]string(nil), _healthEventArns...)
	}

	if resp, err := client.DescribeEntityAggregates(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of entity aggregates for your Organizations that are affected by
// each of the specified events.
func health_DescribeEntityAggregatesForOrganization(cfg aws.Config, client *health.Client) {
	input := &health.DescribeEntityAggregatesForOrganizationInput{
		// EventArns: []string, // Required
	}

	if len(_healthEventArns) > 0 {
		input.EventArns = append([]string(nil), _healthEventArns...)
	}
	if len(_healthAwsAccountIds) > 0 {
		input.AwsAccountIds = append([]string(nil), _healthAwsAccountIds...)
	}

	if resp, err := client.DescribeEntityAggregatesForOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the number of events of each event type (issue, scheduled change, and
// account notification). If no filter is specified, the counts of all events in
// each category are returned.
//
// This API operation uses pagination. Specify the nextToken parameter in the next
// request to return more results.
func health_DescribeEventAggregates(cfg aws.Config, client *health.Client) {
	input := &health.DescribeEventAggregatesInput{
		// AggregateField: types.EventAggregateField, // Required
	}

	if len(_healthAggregateField) > 0 {
		if err := assignInputField(input, "AggregateField", _healthAggregateField); err != nil {
			log.Errorf("invalid --aggregate-field: %s", err.Error())
			return
		}
	}
	if len(_healthFilter) > 0 {
		if err := assignInputField(input, "Filter", _healthFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_healthMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _healthMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_healthNextToken) > 0 {
		input.NextToken = aws.String(_healthNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeEventAggregates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*health.DescribeEventAggregatesOutput
	p := health.NewDescribeEventAggregatesPaginator(client, input)
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

// Returns detailed information about one or more specified events. Information
// includes standard event data (Amazon Web Services Region, service, and so on, as
// returned by [DescribeEvents]), a detailed event description, and possible additional metadata
// that depends upon the nature of the event. Affected entities are not included.
// To retrieve the entities, use the [DescribeAffectedEntities]operation.
//
// If a specified event can't be retrieved, an error message is returned for that
// event.
//
// This operation supports resource-level permissions. You can use this operation
// to allow or deny access to specific Health events. For more information, see [Resource- and action-based conditions]in
// the Health User Guide.
//
// [DescribeEvents]: https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEvents.html
// [DescribeAffectedEntities]: https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeAffectedEntities.html
// [Resource- and action-based conditions]: https://docs.aws.amazon.com/health/latest/ug/security_iam_id-based-policy-examples.html#resource-action-based-conditions
func health_DescribeEventDetails(cfg aws.Config, client *health.Client) {
	input := &health.DescribeEventDetailsInput{
		// EventArns: []string, // Required
	}

	if len(_healthEventArns) > 0 {
		input.EventArns = append([]string(nil), _healthEventArns...)
	}
	if len(_healthLocale) > 0 {
		input.Locale = aws.String(_healthLocale)
	}

	if resp, err := client.DescribeEventDetails(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns detailed information about one or more specified events for one or more
// Amazon Web Services accounts in your organization. This information includes
// standard event data (such as the Amazon Web Services Region and service), an
// event description, and (depending on the event) possible metadata. This
// operation doesn't return affected entities, such as the resources related to the
// event. To return affected entities, use the [DescribeAffectedEntitiesForOrganization]operation.
//
// Before you can call this operation, you must first enable Health to work with
// Organizations. To do this, call the [EnableHealthServiceAccessForOrganization]operation from your organization's
// management account.
//
// When you call the DescribeEventDetailsForOrganization operation, specify the
// organizationEventDetailFilters object in the request. Depending on the Health
// event type, note the following differences:
//
// - To return event details for a public event, you must specify a null value
// for the awsAccountId parameter. If you specify an account ID for a public
// event, Health returns an error message because public events aren't specific to
// an account.
//
// - To return event details for an event that is specific to an account in your
// organization, you must specify the awsAccountId parameter in the request. If
// you don't specify an account ID, Health returns an error message because the
// event is specific to an account in your organization.
//
// For more information, see [Event].
//
// This operation doesn't support resource-level permissions. You can't use this
// operation to allow or deny access to specific Health events. For more
// information, see [Resource- and action-based conditions]in the Health User Guide.
//
// [Resource- and action-based conditions]: https://docs.aws.amazon.com/health/latest/ug/security_iam_id-based-policy-examples.html#resource-action-based-conditions
// [DescribeAffectedEntitiesForOrganization]: https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeAffectedEntitiesForOrganization.html
// [Event]: https://docs.aws.amazon.com/health/latest/APIReference/API_Event.html
// [EnableHealthServiceAccessForOrganization]: https://docs.aws.amazon.com/health/latest/APIReference/API_EnableHealthServiceAccessForOrganization.html
func health_DescribeEventDetailsForOrganization(cfg aws.Config, client *health.Client) {
	input := &health.DescribeEventDetailsForOrganizationInput{
		// OrganizationEventDetailFilters: []types.EventAccountFilter, // Required
	}

	if len(_healthOrganizationEventDetailFilters) > 0 {
		if err := assignInputField(input, "OrganizationEventDetailFilters", _healthOrganizationEventDetailFilters); err != nil {
			log.Errorf("invalid --organization-event-detail-filters: %s", err.Error())
			return
		}
	}
	if len(_healthLocale) > 0 {
		input.Locale = aws.String(_healthLocale)
	}

	if resp, err := client.DescribeEventDetailsForOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the event types that meet the specified filter criteria. You can use
// this API operation to find information about the Health event, such as the
// category, Amazon Web Services service, and event code. The metadata for each
// event appears in the [EventType]object.
//
// If you don't specify a filter criteria, the API operation returns all event
// types, in no particular order.
//
// This API operation uses pagination. Specify the nextToken parameter in the next
// request to return more results.
//
// [EventType]: https://docs.aws.amazon.com/health/latest/APIReference/API_EventType.html
func health_DescribeEventTypes(cfg aws.Config, client *health.Client) {
	input := &health.DescribeEventTypesInput{}

	if len(_healthFilter) > 0 {
		if err := assignInputField(input, "Filter", _healthFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_healthLocale) > 0 {
		input.Locale = aws.String(_healthLocale)
	}
	if len(_healthMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _healthMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_healthNextToken) > 0 {
		input.NextToken = aws.String(_healthNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeEventTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*health.DescribeEventTypesOutput
	p := health.NewDescribeEventTypesPaginator(client, input)
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

// Returns information about events that meet the specified filter criteria.
// Events are returned in a summary form and do not include the detailed
// description, any additional metadata that depends on the event type, or any
// affected resources. To retrieve that information, use the [DescribeEventDetails]and [DescribeAffectedEntities] operations.
//
// If no filter criteria are specified, all events are returned. Results are
// sorted by lastModifiedTime , starting with the most recent event.
//
// - When you call the DescribeEvents operation and specify an entity for the
// entityValues parameter, Health might return public events that aren't specific
// to that resource. For example, if you call DescribeEvents and specify an ID
// for an Amazon Elastic Compute Cloud (Amazon EC2) instance, Health might return
// events that aren't specific to that resource or service. To get events that are
// specific to a service, use the services parameter in the filter object. For
// more information, see [Event].
//
// - This API operation uses pagination. Specify the nextToken parameter in the
// next request to return more results.
//
// [DescribeEventDetails]: https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventDetails.html
// [DescribeAffectedEntities]: https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeAffectedEntities.html
// [Event]: https://docs.aws.amazon.com/health/latest/APIReference/API_Event.html
func health_DescribeEvents(cfg aws.Config, client *health.Client) {
	input := &health.DescribeEventsInput{}

	if len(_healthFilter) > 0 {
		if err := assignInputField(input, "Filter", _healthFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_healthLocale) > 0 {
		input.Locale = aws.String(_healthLocale)
	}
	if len(_healthMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _healthMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_healthNextToken) > 0 {
		input.NextToken = aws.String(_healthNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*health.DescribeEventsOutput
	p := health.NewDescribeEventsPaginator(client, input)
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

// Returns information about events across your organization in Organizations. You
// can use the filters parameter to specify the events that you want to return.
// Events are returned in a summary form and don't include the affected accounts,
// detailed description, any additional metadata that depends on the event type, or
// any affected resources. To retrieve that information, use the following
// operations:
//
// [DescribeAffectedAccountsForOrganization]
//
// [DescribeEventDetailsForOrganization]
//
// [DescribeAffectedEntitiesForOrganization]
//
// If you don't specify a filter , the DescribeEventsForOrganizations returns all
// events across your organization. Results are sorted by lastModifiedTime ,
// starting with the most recent event.
//
// For more information about the different types of Health events, see [Event].
//
// Before you can call this operation, you must first enable Health to work with
// Organizations. To do this, call the [EnableHealthServiceAccessForOrganization]operation from your organization's
// management account.
//
// This API operation uses pagination. Specify the nextToken parameter in the next
// request to return more results.
//
// [DescribeEventDetailsForOrganization]: https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventDetailsForOrganization.html
// [DescribeAffectedEntitiesForOrganization]: https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeAffectedEntitiesForOrganization.html
// [DescribeAffectedAccountsForOrganization]: https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeAffectedAccountsForOrganization.html
// [Event]: https://docs.aws.amazon.com/health/latest/APIReference/API_Event.html
// [EnableHealthServiceAccessForOrganization]: https://docs.aws.amazon.com/health/latest/APIReference/API_EnableHealthServiceAccessForOrganization.html
func health_DescribeEventsForOrganization(cfg aws.Config, client *health.Client) {
	input := &health.DescribeEventsForOrganizationInput{}

	if len(_healthFilter) > 0 {
		if err := assignInputField(input, "Filter", _healthFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_healthLocale) > 0 {
		input.Locale = aws.String(_healthLocale)
	}
	if len(_healthMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _healthMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_healthNextToken) > 0 {
		input.NextToken = aws.String(_healthNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeEventsForOrganization(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*health.DescribeEventsForOrganizationOutput
	p := health.NewDescribeEventsForOrganizationPaginator(client, input)
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

// This operation provides status information on enabling or disabling Health to
// work with your organization. To call this operation, you must use the
// organization's management account.
func health_DescribeHealthServiceStatusForOrganization(cfg aws.Config, client *health.Client) {
	input := &health.DescribeHealthServiceStatusForOrganizationInput{}

	if resp, err := client.DescribeHealthServiceStatusForOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables Health from working with Organizations. To call this operation, you
// must sign in to the organization's management account. For more information, see
// [Aggregating Health events]in the Health User Guide.
//
// This operation doesn't remove the service-linked role from the management
// account in your organization. You must use the IAM console, API, or Command Line
// Interface (CLI) to remove the service-linked role. For more information, see [Deleting a Service-Linked Role]in
// the IAM User Guide.
//
// You can also disable the organizational feature by using the Organizations [DisableAWSServiceAccess] API
// operation. After you call this operation, Health stops aggregating events for
// all other Amazon Web Services accounts in your organization. If you call the
// Health API operations for organizational view, Health returns an error. Health
// continues to aggregate health events for your Amazon Web Services account.
//
// [DisableAWSServiceAccess]: https://docs.aws.amazon.com/organizations/latest/APIReference/API_DisableAWSServiceAccess.html
// [Aggregating Health events]: https://docs.aws.amazon.com/health/latest/ug/aggregate-events.html
// [Deleting a Service-Linked Role]: https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html#delete-service-linked-role
func health_DisableHealthServiceAccessForOrganization(cfg aws.Config, client *health.Client) {
	input := &health.DisableHealthServiceAccessForOrganizationInput{}

	if resp, err := client.DisableHealthServiceAccessForOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables Health to work with Organizations. You can use the organizational view
// feature to aggregate events from all Amazon Web Services accounts in your
// organization in a centralized location.
//
// This operation also creates a service-linked role for the management account in
// the organization.
//
// To call this operation, you must meet the following requirements:
//
// - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan
// from [Amazon Web Services Support]to use the Health API. If you call the Health API from an Amazon Web
// Services account that doesn't have a Business, Enterprise On-Ramp, or Enterprise
// Support plan, you receive a SubscriptionRequiredException error.
//
// - You must have permission to call this operation from the organization's
// management account. For example IAM policies, see [Health identity-based policy examples].
//
// If you don't have the required support plan, you can instead use the Health
// console to enable the organizational view feature. For more information, see [Aggregating Health events]in
// the Health User Guide.
//
// [Amazon Web Services Support]: http://aws.amazon.com/premiumsupport/
// [Aggregating Health events]: https://docs.aws.amazon.com/health/latest/ug/aggregate-events.html
// [Health identity-based policy examples]: https://docs.aws.amazon.com/health/latest/ug/security_iam_id-based-policy-examples.html
func health_EnableHealthServiceAccessForOrganization(cfg aws.Config, client *health.Client) {
	input := &health.EnableHealthServiceAccessForOrganizationInput{}

	if resp, err := client.EnableHealthServiceAccessForOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_healthCmd)
	_healthCmd.Flags().SortFlags = false

	_healthCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_healthCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_healthCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_healthCmd.Flags().StringVarP(&_healthAggregateField, "aggregate-field", "", "", "Aggregate Field")
	_healthCmd.Flags().StringSliceVarP(&_healthAwsAccountIds, "aws-account-ids", "", nil, "AWS Account Ids")
	_healthCmd.Flags().StringVarP(&_healthEventArn, "event-arn", "", "", "Event ARN")
	_healthCmd.Flags().StringSliceVarP(&_healthEventArns, "event-arns", "", nil, "Event Arns")
	_healthCmd.Flags().StringVarP(&_healthFilter, "filter", "", "", "Filter")
	_healthCmd.Flags().StringVarP(&_healthLocale, "locale", "", "", "Locale")
	_healthCmd.Flags().StringVarP(&_healthMaxResults, "max-results", "", "", "Max Results")
	_healthCmd.Flags().StringVarP(&_healthNextToken, "next-token", "", "", "Next Token")
	_healthCmd.Flags().StringVarP(&_healthOrganizationEntityAccountFilters, "organization-entity-account-filters", "", "", "Organization Entity Account Filters")
	_healthCmd.Flags().StringVarP(&_healthOrganizationEntityFilters, "organization-entity-filters", "", "", "Organization Entity Filters")
	_healthCmd.Flags().StringVarP(&_healthOrganizationEventDetailFilters, "organization-event-detail-filters", "", "", "Organization Event Detail Filters")

	_healthCmd.Flags().BoolVarP(&_healthDescribeAffectedAccountsForOrganization, "describe-affected-accounts-for-organization", "", false, "Describe Affected Accounts For Organization")
	_healthCmd.Flags().BoolVarP(&_healthDescribeAffectedEntities, "describe-affected-entities", "", false, "Describe Affected Entities")
	_healthCmd.Flags().BoolVarP(&_healthDescribeAffectedEntitiesForOrganization, "describe-affected-entities-for-organization", "", false, "Describe Affected Entities For Organization")
	_healthCmd.Flags().BoolVarP(&_healthDescribeEntityAggregates, "describe-entity-aggregates", "", false, "Describe Entity Aggregates")
	_healthCmd.Flags().BoolVarP(&_healthDescribeEntityAggregatesForOrganization, "describe-entity-aggregates-for-organization", "", false, "Describe Entity Aggregates For Organization")
	_healthCmd.Flags().BoolVarP(&_healthDescribeEventAggregates, "describe-event-aggregates", "", false, "Describe Event Aggregates")
	_healthCmd.Flags().BoolVarP(&_healthDescribeEventDetails, "describe-event-details", "", false, "Describe Event Details")
	_healthCmd.Flags().BoolVarP(&_healthDescribeEventDetailsForOrganization, "describe-event-details-for-organization", "", false, "Describe Event Details For Organization")
	_healthCmd.Flags().BoolVarP(&_healthDescribeEventTypes, "describe-event-types", "", false, "Describe Event Types")
	_healthCmd.Flags().BoolVarP(&_healthDescribeEvents, "describe-events", "", false, "Describe Events")
	_healthCmd.Flags().BoolVarP(&_healthDescribeEventsForOrganization, "describe-events-for-organization", "", false, "Describe Events For Organization")
	_healthCmd.Flags().BoolVarP(&_healthDescribeHealthServiceStatusForOrganization, "describe-health-service-status-for-organization", "", false, "Describe Health Service Status For Organization")
	_healthCmd.Flags().BoolVarP(&_healthDisableHealthServiceAccessForOrganization, "disable-health-service-access-for-organization", "", false, "Disable Health Service Access For Organization")
	_healthCmd.Flags().BoolVarP(&_healthEnableHealthServiceAccessForOrganization, "enable-health-service-access-for-organization", "", false, "Enable Health Service Access For Organization")

}
