package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cloudwatcheventsCmd represents the cloudwatchevents command
var _cloudwatcheventsCmd = &cobra.Command{
	Use:   "cloudwatchevents",
	Short: "AWS cloudwatchevents CLI",
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
		client := cloudwatchevents.NewFromConfig(cfg)
		if _cloudwatcheventsActivateEventSource {
			cloudwatchevents_ActivateEventSource(cfg, client)
			return
		}
		if _cloudwatcheventsCancelReplay {
			cloudwatchevents_CancelReplay(cfg, client)
			return
		}
		if _cloudwatcheventsCreateApiDestination {
			cloudwatchevents_CreateApiDestination(cfg, client)
			return
		}
		if _cloudwatcheventsCreateArchive {
			cloudwatchevents_CreateArchive(cfg, client)
			return
		}
		if _cloudwatcheventsCreateConnection {
			cloudwatchevents_CreateConnection(cfg, client)
			return
		}
		if _cloudwatcheventsCreateEventBus {
			cloudwatchevents_CreateEventBus(cfg, client)
			return
		}
		if _cloudwatcheventsCreatePartnerEventSource {
			cloudwatchevents_CreatePartnerEventSource(cfg, client)
			return
		}
		if _cloudwatcheventsDeactivateEventSource {
			cloudwatchevents_DeactivateEventSource(cfg, client)
			return
		}
		if _cloudwatcheventsDeauthorizeConnection {
			cloudwatchevents_DeauthorizeConnection(cfg, client)
			return
		}
		if _cloudwatcheventsDeleteApiDestination {
			cloudwatchevents_DeleteApiDestination(cfg, client)
			return
		}
		if _cloudwatcheventsDeleteArchive {
			cloudwatchevents_DeleteArchive(cfg, client)
			return
		}
		if _cloudwatcheventsDeleteConnection {
			cloudwatchevents_DeleteConnection(cfg, client)
			return
		}
		if _cloudwatcheventsDeleteEventBus {
			cloudwatchevents_DeleteEventBus(cfg, client)
			return
		}
		if _cloudwatcheventsDeletePartnerEventSource {
			cloudwatchevents_DeletePartnerEventSource(cfg, client)
			return
		}
		if _cloudwatcheventsDeleteRule {
			cloudwatchevents_DeleteRule(cfg, client)
			return
		}
		if _cloudwatcheventsDescribeApiDestination {
			cloudwatchevents_DescribeApiDestination(cfg, client)
			return
		}
		if _cloudwatcheventsDescribeArchive {
			cloudwatchevents_DescribeArchive(cfg, client)
			return
		}
		if _cloudwatcheventsDescribeConnection {
			cloudwatchevents_DescribeConnection(cfg, client)
			return
		}
		if _cloudwatcheventsDescribeEventBus {
			cloudwatchevents_DescribeEventBus(cfg, client)
			return
		}
		if _cloudwatcheventsDescribeEventSource {
			cloudwatchevents_DescribeEventSource(cfg, client)
			return
		}
		if _cloudwatcheventsDescribePartnerEventSource {
			cloudwatchevents_DescribePartnerEventSource(cfg, client)
			return
		}
		if _cloudwatcheventsDescribeReplay {
			cloudwatchevents_DescribeReplay(cfg, client)
			return
		}
		if _cloudwatcheventsDescribeRule {
			cloudwatchevents_DescribeRule(cfg, client)
			return
		}
		if _cloudwatcheventsDisableRule {
			cloudwatchevents_DisableRule(cfg, client)
			return
		}
		if _cloudwatcheventsEnableRule {
			cloudwatchevents_EnableRule(cfg, client)
			return
		}
		if _cloudwatcheventsListApiDestinations {
			cloudwatchevents_ListApiDestinations(cfg, client)
			return
		}
		if _cloudwatcheventsListArchives {
			cloudwatchevents_ListArchives(cfg, client)
			return
		}
		if _cloudwatcheventsListConnections {
			cloudwatchevents_ListConnections(cfg, client)
			return
		}
		if _cloudwatcheventsListEventBuses {
			cloudwatchevents_ListEventBuses(cfg, client)
			return
		}
		if _cloudwatcheventsListEventSources {
			cloudwatchevents_ListEventSources(cfg, client)
			return
		}
		if _cloudwatcheventsListPartnerEventSourceAccounts {
			cloudwatchevents_ListPartnerEventSourceAccounts(cfg, client)
			return
		}
		if _cloudwatcheventsListPartnerEventSources {
			cloudwatchevents_ListPartnerEventSources(cfg, client)
			return
		}
		if _cloudwatcheventsListReplays {
			cloudwatchevents_ListReplays(cfg, client)
			return
		}
		if _cloudwatcheventsListRuleNamesByTarget {
			cloudwatchevents_ListRuleNamesByTarget(cfg, client)
			return
		}
		if _cloudwatcheventsListRules {
			cloudwatchevents_ListRules(cfg, client)
			return
		}
		if _cloudwatcheventsListTagsForResource {
			cloudwatchevents_ListTagsForResource(cfg, client)
			return
		}
		if _cloudwatcheventsListTargetsByRule {
			cloudwatchevents_ListTargetsByRule(cfg, client)
			return
		}
		if _cloudwatcheventsPutEvents {
			cloudwatchevents_PutEvents(cfg, client)
			return
		}
		if _cloudwatcheventsPutPartnerEvents {
			cloudwatchevents_PutPartnerEvents(cfg, client)
			return
		}
		if _cloudwatcheventsPutPermission {
			cloudwatchevents_PutPermission(cfg, client)
			return
		}
		if _cloudwatcheventsPutRule {
			cloudwatchevents_PutRule(cfg, client)
			return
		}
		if _cloudwatcheventsPutTargets {
			cloudwatchevents_PutTargets(cfg, client)
			return
		}
		if _cloudwatcheventsRemovePermission {
			cloudwatchevents_RemovePermission(cfg, client)
			return
		}
		if _cloudwatcheventsRemoveTargets {
			cloudwatchevents_RemoveTargets(cfg, client)
			return
		}
		if _cloudwatcheventsStartReplay {
			cloudwatchevents_StartReplay(cfg, client)
			return
		}
		if _cloudwatcheventsTagResource {
			cloudwatchevents_TagResource(cfg, client)
			return
		}
		if _cloudwatcheventsTestEventPattern {
			cloudwatchevents_TestEventPattern(cfg, client)
			return
		}
		if _cloudwatcheventsUntagResource {
			cloudwatchevents_UntagResource(cfg, client)
			return
		}
		if _cloudwatcheventsUpdateApiDestination {
			cloudwatchevents_UpdateApiDestination(cfg, client)
			return
		}
		if _cloudwatcheventsUpdateArchive {
			cloudwatchevents_UpdateArchive(cfg, client)
			return
		}
		if _cloudwatcheventsUpdateConnection {
			cloudwatchevents_UpdateConnection(cfg, client)
			return
		}

	},
}

var (
	_cloudwatcheventsActivateEventSource            bool
	_cloudwatcheventsCancelReplay                   bool
	_cloudwatcheventsCreateApiDestination           bool
	_cloudwatcheventsCreateArchive                  bool
	_cloudwatcheventsCreateConnection               bool
	_cloudwatcheventsCreateEventBus                 bool
	_cloudwatcheventsCreatePartnerEventSource       bool
	_cloudwatcheventsDeactivateEventSource          bool
	_cloudwatcheventsDeauthorizeConnection          bool
	_cloudwatcheventsDeleteApiDestination           bool
	_cloudwatcheventsDeleteArchive                  bool
	_cloudwatcheventsDeleteConnection               bool
	_cloudwatcheventsDeleteEventBus                 bool
	_cloudwatcheventsDeletePartnerEventSource       bool
	_cloudwatcheventsDeleteRule                     bool
	_cloudwatcheventsDescribeApiDestination         bool
	_cloudwatcheventsDescribeArchive                bool
	_cloudwatcheventsDescribeConnection             bool
	_cloudwatcheventsDescribeEventBus               bool
	_cloudwatcheventsDescribeEventSource            bool
	_cloudwatcheventsDescribePartnerEventSource     bool
	_cloudwatcheventsDescribeReplay                 bool
	_cloudwatcheventsDescribeRule                   bool
	_cloudwatcheventsDisableRule                    bool
	_cloudwatcheventsEnableRule                     bool
	_cloudwatcheventsListApiDestinations            bool
	_cloudwatcheventsListArchives                   bool
	_cloudwatcheventsListConnections                bool
	_cloudwatcheventsListEventBuses                 bool
	_cloudwatcheventsListEventSources               bool
	_cloudwatcheventsListPartnerEventSourceAccounts bool
	_cloudwatcheventsListPartnerEventSources        bool
	_cloudwatcheventsListReplays                    bool
	_cloudwatcheventsListRuleNamesByTarget          bool
	_cloudwatcheventsListRules                      bool
	_cloudwatcheventsListTagsForResource            bool
	_cloudwatcheventsListTargetsByRule              bool
	_cloudwatcheventsPutEvents                      bool
	_cloudwatcheventsPutPartnerEvents               bool
	_cloudwatcheventsPutPermission                  bool
	_cloudwatcheventsPutRule                        bool
	_cloudwatcheventsPutTargets                     bool
	_cloudwatcheventsRemovePermission               bool
	_cloudwatcheventsRemoveTargets                  bool
	_cloudwatcheventsStartReplay                    bool
	_cloudwatcheventsTagResource                    bool
	_cloudwatcheventsTestEventPattern               bool
	_cloudwatcheventsUntagResource                  bool
	_cloudwatcheventsUpdateApiDestination           bool
	_cloudwatcheventsUpdateArchive                  bool
	_cloudwatcheventsUpdateConnection               bool

	_cloudwatcheventsAccount                      string
	_cloudwatcheventsAction                       string
	_cloudwatcheventsArchiveName                  string
	_cloudwatcheventsAuthParameters               string
	_cloudwatcheventsAuthorizationType            string
	_cloudwatcheventsCondition                    string
	_cloudwatcheventsConnectionArn                string
	_cloudwatcheventsConnectionState              string
	_cloudwatcheventsDescription                  string
	_cloudwatcheventsDestination                  string
	_cloudwatcheventsEntries                      string
	_cloudwatcheventsEvent                        string
	_cloudwatcheventsEventBusName                 string
	_cloudwatcheventsEventEndTime                 string
	_cloudwatcheventsEventPattern                 string
	_cloudwatcheventsEventSourceArn               string
	_cloudwatcheventsEventSourceName              string
	_cloudwatcheventsEventStartTime               string
	_cloudwatcheventsForce                        string
	_cloudwatcheventsHttpMethod                   string
	_cloudwatcheventsIds                          []string
	_cloudwatcheventsInvocationEndpoint           string
	_cloudwatcheventsInvocationRateLimitPerSecond string
	_cloudwatcheventsLimit                        string
	_cloudwatcheventsName                         string
	_cloudwatcheventsNamePrefix                   string
	_cloudwatcheventsNextToken                    string
	_cloudwatcheventsPolicy                       string
	_cloudwatcheventsPrincipal                    string
	_cloudwatcheventsRemoveAllPermissions         string
	_cloudwatcheventsReplayName                   string
	_cloudwatcheventsResourceARN                  string
	_cloudwatcheventsRetentionDays                string
	_cloudwatcheventsRoleArn                      string
	_cloudwatcheventsRule                         string
	_cloudwatcheventsScheduleExpression           string
	_cloudwatcheventsState                        string
	_cloudwatcheventsStatementId                  string
	_cloudwatcheventsTagKeys                      []string
	_cloudwatcheventsTags                         string
	_cloudwatcheventsTargetArn                    string
	_cloudwatcheventsTargets                      string
)

// Activates a partner event source that has been deactivated. Once activated,
// your matching event bus will start receiving events from the event source.
func cloudwatchevents_ActivateEventSource(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.ActivateEventSourceInput{
		// Name: *string, // Required
	}

	if len(_cloudwatcheventsName) > 0 {
		input.Name = aws.String(_cloudwatcheventsName)
	}

	if resp, err := client.ActivateEventSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels the specified replay.
func cloudwatchevents_CancelReplay(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.CancelReplayInput{
		// ReplayName: *string, // Required
	}

	if len(_cloudwatcheventsReplayName) > 0 {
		input.ReplayName = aws.String(_cloudwatcheventsReplayName)
	}

	if resp, err := client.CancelReplay(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an API destination, which is an HTTP invocation endpoint configured as
// a target for events.
func cloudwatchevents_CreateApiDestination(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.CreateApiDestinationInput{
		// ConnectionArn: *string, // Required
		// HttpMethod: types.ApiDestinationHttpMethod, // Required
		// InvocationEndpoint: *string, // Required
		// Name: *string, // Required
	}

	if len(_cloudwatcheventsConnectionArn) > 0 {
		input.ConnectionArn = aws.String(_cloudwatcheventsConnectionArn)
	}
	if len(_cloudwatcheventsHttpMethod) > 0 {
		if err := assignInputField(input, "HttpMethod", _cloudwatcheventsHttpMethod); err != nil {
			log.Errorf("invalid --http-method: %s", err.Error())
			return
		}
	}
	if len(_cloudwatcheventsInvocationEndpoint) > 0 {
		input.InvocationEndpoint = aws.String(_cloudwatcheventsInvocationEndpoint)
	}
	if len(_cloudwatcheventsName) > 0 {
		input.Name = aws.String(_cloudwatcheventsName)
	}
	if len(_cloudwatcheventsDescription) > 0 {
		input.Description = aws.String(_cloudwatcheventsDescription)
	}
	if len(_cloudwatcheventsInvocationRateLimitPerSecond) > 0 {
		if err := assignInputField(input, "InvocationRateLimitPerSecond", _cloudwatcheventsInvocationRateLimitPerSecond); err != nil {
			log.Errorf("invalid --invocation-rate-limit-per-second: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateApiDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an archive of events with the specified settings. When you create an
// archive, incoming events might not immediately start being sent to the archive.
// Allow a short period of time for changes to take effect. If you do not specify a
// pattern to filter events sent to the archive, all events are sent to the archive
// except replayed events. Replayed events are not sent to an archive.
func cloudwatchevents_CreateArchive(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.CreateArchiveInput{
		// ArchiveName: *string, // Required
		// EventSourceArn: *string, // Required
	}

	if len(_cloudwatcheventsArchiveName) > 0 {
		input.ArchiveName = aws.String(_cloudwatcheventsArchiveName)
	}
	if len(_cloudwatcheventsEventSourceArn) > 0 {
		input.EventSourceArn = aws.String(_cloudwatcheventsEventSourceArn)
	}
	if len(_cloudwatcheventsDescription) > 0 {
		input.Description = aws.String(_cloudwatcheventsDescription)
	}
	if len(_cloudwatcheventsEventPattern) > 0 {
		input.EventPattern = aws.String(_cloudwatcheventsEventPattern)
	}
	if len(_cloudwatcheventsRetentionDays) > 0 {
		if err := assignInputField(input, "RetentionDays", _cloudwatcheventsRetentionDays); err != nil {
			log.Errorf("invalid --retention-days: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateArchive(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a connection. A connection defines the authorization type and
// credentials to use for authorization with an API destination HTTP endpoint.
func cloudwatchevents_CreateConnection(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.CreateConnectionInput{
		// AuthParameters: *types.CreateConnectionAuthRequestParameters, // Required
		// AuthorizationType: types.ConnectionAuthorizationType, // Required
		// Name: *string, // Required
	}

	if len(_cloudwatcheventsAuthParameters) > 0 {
		if err := assignInputField(input, "AuthParameters", _cloudwatcheventsAuthParameters); err != nil {
			log.Errorf("invalid --auth-parameters: %s", err.Error())
			return
		}
	}
	if len(_cloudwatcheventsAuthorizationType) > 0 {
		if err := assignInputField(input, "AuthorizationType", _cloudwatcheventsAuthorizationType); err != nil {
			log.Errorf("invalid --authorization-type: %s", err.Error())
			return
		}
	}
	if len(_cloudwatcheventsName) > 0 {
		input.Name = aws.String(_cloudwatcheventsName)
	}
	if len(_cloudwatcheventsDescription) > 0 {
		input.Description = aws.String(_cloudwatcheventsDescription)
	}

	if resp, err := client.CreateConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new event bus within your account. This can be a custom event bus
// which you can use to receive events from your custom applications and services,
// or it can be a partner event bus which can be matched to a partner event source.
func cloudwatchevents_CreateEventBus(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.CreateEventBusInput{
		// Name: *string, // Required
	}

	if len(_cloudwatcheventsName) > 0 {
		input.Name = aws.String(_cloudwatcheventsName)
	}
	if len(_cloudwatcheventsEventSourceName) > 0 {
		input.EventSourceName = aws.String(_cloudwatcheventsEventSourceName)
	}
	if len(_cloudwatcheventsTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudwatcheventsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEventBus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Called by an SaaS partner to create a partner event source. This operation is
// not used by Amazon Web Services customers.
//
// Each partner event source can be used by one Amazon Web Services account to
// create a matching partner event bus in that Amazon Web Services account. A SaaS
// partner must create one partner event source for each Amazon Web Services
// account that wants to receive those event types.
//
// A partner event source creates events based on resources within the SaaS
// partner's service or application.
//
// An Amazon Web Services account that creates a partner event bus that matches
// the partner event source can use that event bus to receive events from the
// partner, and then process them using Amazon Web Services Events rules and
// targets.
//
// Partner event source names follow this format:
//
// partner_name/event_namespace/event_name
//
// partner_name is determined during partner registration and identifies the
// partner to Amazon Web Services customers. event_namespace is determined by the
// partner and is a way for the partner to categorize their events. event_name is
// determined by the partner, and should uniquely identify an event-generating
// resource within the partner system. The combination of event_namespace and
// event_name should help Amazon Web Services customers decide whether to create an
// event bus to receive these events.
func cloudwatchevents_CreatePartnerEventSource(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.CreatePartnerEventSourceInput{
		// Account: *string, // Required
		// Name: *string, // Required
	}

	if len(_cloudwatcheventsAccount) > 0 {
		input.Account = aws.String(_cloudwatcheventsAccount)
	}
	if len(_cloudwatcheventsName) > 0 {
		input.Name = aws.String(_cloudwatcheventsName)
	}

	if resp, err := client.CreatePartnerEventSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// You can use this operation to temporarily stop receiving events from the
// specified partner event source. The matching event bus is not deleted.
//
// When you deactivate a partner event source, the source goes into PENDING state.
// If it remains in PENDING state for more than two weeks, it is deleted.
//
// To activate a deactivated partner event source, use [ActivateEventSource].
//
// [ActivateEventSource]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_ActivateEventSource.html
func cloudwatchevents_DeactivateEventSource(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.DeactivateEventSourceInput{
		// Name: *string, // Required
	}

	if len(_cloudwatcheventsName) > 0 {
		input.Name = aws.String(_cloudwatcheventsName)
	}

	if resp, err := client.DeactivateEventSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes all authorization parameters from the connection. This lets you remove
// the secret from the connection so you can reuse it without having to create a
// new connection.
func cloudwatchevents_DeauthorizeConnection(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.DeauthorizeConnectionInput{
		// Name: *string, // Required
	}

	if len(_cloudwatcheventsName) > 0 {
		input.Name = aws.String(_cloudwatcheventsName)
	}

	if resp, err := client.DeauthorizeConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified API destination.
func cloudwatchevents_DeleteApiDestination(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.DeleteApiDestinationInput{
		// Name: *string, // Required
	}

	if len(_cloudwatcheventsName) > 0 {
		input.Name = aws.String(_cloudwatcheventsName)
	}

	if resp, err := client.DeleteApiDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified archive.
func cloudwatchevents_DeleteArchive(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.DeleteArchiveInput{
		// ArchiveName: *string, // Required
	}

	if len(_cloudwatcheventsArchiveName) > 0 {
		input.ArchiveName = aws.String(_cloudwatcheventsArchiveName)
	}

	if resp, err := client.DeleteArchive(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a connection.
func cloudwatchevents_DeleteConnection(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.DeleteConnectionInput{
		// Name: *string, // Required
	}

	if len(_cloudwatcheventsName) > 0 {
		input.Name = aws.String(_cloudwatcheventsName)
	}

	if resp, err := client.DeleteConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified custom event bus or partner event bus. All rules
// associated with this event bus need to be deleted. You can't delete your
// account's default event bus.
func cloudwatchevents_DeleteEventBus(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.DeleteEventBusInput{
		// Name: *string, // Required
	}

	if len(_cloudwatcheventsName) > 0 {
		input.Name = aws.String(_cloudwatcheventsName)
	}

	if resp, err := client.DeleteEventBus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is used by SaaS partners to delete a partner event source. This
// operation is not used by Amazon Web Services customers.
//
// When you delete an event source, the status of the corresponding partner event
// bus in the Amazon Web Services customer account becomes DELETED.
func cloudwatchevents_DeletePartnerEventSource(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.DeletePartnerEventSourceInput{
		// Account: *string, // Required
		// Name: *string, // Required
	}

	if len(_cloudwatcheventsAccount) > 0 {
		input.Account = aws.String(_cloudwatcheventsAccount)
	}
	if len(_cloudwatcheventsName) > 0 {
		input.Name = aws.String(_cloudwatcheventsName)
	}

	if resp, err := client.DeletePartnerEventSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified rule.
// Before you can delete the rule, you must remove all targets, using [RemoveTargets].
//
// When you delete a rule, incoming events might continue to match to the deleted
// rule. Allow a short period of time for changes to take effect.
//
// If you call delete rule multiple times for the same rule, all calls will
// succeed. When you call delete rule for a non-existent custom eventbus,
// ResourceNotFoundException is returned.
//
// Managed rules are rules created and managed by another Amazon Web Services
// service on your behalf. These rules are created by those other Amazon Web
// Services services to support functionality in those services. You can delete
// these rules using the Force option, but you should do so only if you are sure
// the other service is not still using that rule.
//
// [RemoveTargets]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_RemoveTargets.html
func cloudwatchevents_DeleteRule(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.DeleteRuleInput{
		// Name: *string, // Required
	}

	if len(_cloudwatcheventsName) > 0 {
		input.Name = aws.String(_cloudwatcheventsName)
	}
	if len(_cloudwatcheventsEventBusName) > 0 {
		input.EventBusName = aws.String(_cloudwatcheventsEventBusName)
	}
	if len(_cloudwatcheventsForce) > 0 {
		if err := assignInputField(input, "Force", _cloudwatcheventsForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about an API destination.
func cloudwatchevents_DescribeApiDestination(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.DescribeApiDestinationInput{
		// Name: *string, // Required
	}

	if len(_cloudwatcheventsName) > 0 {
		input.Name = aws.String(_cloudwatcheventsName)
	}

	if resp, err := client.DescribeApiDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about an archive.
func cloudwatchevents_DescribeArchive(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.DescribeArchiveInput{
		// ArchiveName: *string, // Required
	}

	if len(_cloudwatcheventsArchiveName) > 0 {
		input.ArchiveName = aws.String(_cloudwatcheventsArchiveName)
	}

	if resp, err := client.DescribeArchive(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about a connection.
func cloudwatchevents_DescribeConnection(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.DescribeConnectionInput{
		// Name: *string, // Required
	}

	if len(_cloudwatcheventsName) > 0 {
		input.Name = aws.String(_cloudwatcheventsName)
	}

	if resp, err := client.DescribeConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays details about an event bus in your account. This can include the
// external Amazon Web Services accounts that are permitted to write events to your
// default event bus, and the associated policy. For custom event buses and partner
// event buses, it displays the name, ARN, policy, state, and creation time.
//
// To enable your account to receive events from other accounts on its default
// event bus, use [PutPermission].
//
// For more information about partner event buses, see [CreateEventBus].
//
// [PutPermission]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_PutPermission.html
// [CreateEventBus]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_CreateEventBus.html
func cloudwatchevents_DescribeEventBus(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.DescribeEventBusInput{}

	if len(_cloudwatcheventsName) > 0 {
		input.Name = aws.String(_cloudwatcheventsName)
	}

	if resp, err := client.DescribeEventBus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation lists details about a partner event source that is shared with
// your account.
func cloudwatchevents_DescribeEventSource(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.DescribeEventSourceInput{
		// Name: *string, // Required
	}

	if len(_cloudwatcheventsName) > 0 {
		input.Name = aws.String(_cloudwatcheventsName)
	}

	if resp, err := client.DescribeEventSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// An SaaS partner can use this operation to list details about a partner event
// source that they have created. Amazon Web Services customers do not use this
// operation. Instead, Amazon Web Services customers can use [DescribeEventSource]to see details about
// a partner event source that is shared with them.
//
// [DescribeEventSource]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_DescribeEventSource.html
func cloudwatchevents_DescribePartnerEventSource(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.DescribePartnerEventSourceInput{
		// Name: *string, // Required
	}

	if len(_cloudwatcheventsName) > 0 {
		input.Name = aws.String(_cloudwatcheventsName)
	}

	if resp, err := client.DescribePartnerEventSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about a replay. Use DescribeReplay to determine the progress
// of a running replay. A replay processes events to replay based on the time in
// the event, and replays them using 1 minute intervals. If you use StartReplay
// and specify an EventStartTime and an EventEndTime that covers a 20 minute time
// range, the events are replayed from the first minute of that 20 minute range
// first. Then the events from the second minute are replayed. You can use
// DescribeReplay to determine the progress of a replay. The value returned for
// EventLastReplayedTime indicates the time within the specified time range
// associated with the last event replayed.
func cloudwatchevents_DescribeReplay(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.DescribeReplayInput{
		// ReplayName: *string, // Required
	}

	if len(_cloudwatcheventsReplayName) > 0 {
		input.ReplayName = aws.String(_cloudwatcheventsReplayName)
	}

	if resp, err := client.DescribeReplay(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified rule.
// DescribeRule does not list the targets of a rule. To see the targets associated
// with a rule, use [ListTargetsByRule].
//
// [ListTargetsByRule]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_ListTargetsByRule.html
func cloudwatchevents_DescribeRule(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.DescribeRuleInput{
		// Name: *string, // Required
	}

	if len(_cloudwatcheventsName) > 0 {
		input.Name = aws.String(_cloudwatcheventsName)
	}
	if len(_cloudwatcheventsEventBusName) > 0 {
		input.EventBusName = aws.String(_cloudwatcheventsEventBusName)
	}

	if resp, err := client.DescribeRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables the specified rule. A disabled rule won't match any events, and won't
// self-trigger if it has a schedule expression.
//
// When you disable a rule, incoming events might continue to match to the
// disabled rule. Allow a short period of time for changes to take effect.
func cloudwatchevents_DisableRule(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.DisableRuleInput{
		// Name: *string, // Required
	}

	if len(_cloudwatcheventsName) > 0 {
		input.Name = aws.String(_cloudwatcheventsName)
	}
	if len(_cloudwatcheventsEventBusName) > 0 {
		input.EventBusName = aws.String(_cloudwatcheventsEventBusName)
	}

	if resp, err := client.DisableRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the specified rule. If the rule does not exist, the operation fails.
// When you enable a rule, incoming events might not immediately start matching to
// a newly enabled rule. Allow a short period of time for changes to take effect.
func cloudwatchevents_EnableRule(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.EnableRuleInput{
		// Name: *string, // Required
	}

	if len(_cloudwatcheventsName) > 0 {
		input.Name = aws.String(_cloudwatcheventsName)
	}
	if len(_cloudwatcheventsEventBusName) > 0 {
		input.EventBusName = aws.String(_cloudwatcheventsEventBusName)
	}

	if resp, err := client.EnableRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of API destination in the account in the current Region.
func cloudwatchevents_ListApiDestinations(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.ListApiDestinationsInput{}

	if len(_cloudwatcheventsConnectionArn) > 0 {
		input.ConnectionArn = aws.String(_cloudwatcheventsConnectionArn)
	}
	if len(_cloudwatcheventsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatcheventsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatcheventsNamePrefix) > 0 {
		input.NamePrefix = aws.String(_cloudwatcheventsNamePrefix)
	}
	if len(_cloudwatcheventsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatcheventsNextToken)
	}

	if resp, err := client.ListApiDestinations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists your archives. You can either list all the archives or you can provide a
// prefix to match to the archive names. Filter parameters are exclusive.
func cloudwatchevents_ListArchives(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.ListArchivesInput{}

	if len(_cloudwatcheventsEventSourceArn) > 0 {
		input.EventSourceArn = aws.String(_cloudwatcheventsEventSourceArn)
	}
	if len(_cloudwatcheventsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatcheventsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatcheventsNamePrefix) > 0 {
		input.NamePrefix = aws.String(_cloudwatcheventsNamePrefix)
	}
	if len(_cloudwatcheventsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatcheventsNextToken)
	}
	if len(_cloudwatcheventsState) > 0 {
		if err := assignInputField(input, "State", _cloudwatcheventsState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListArchives(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of connections from the account.
func cloudwatchevents_ListConnections(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.ListConnectionsInput{}

	if len(_cloudwatcheventsConnectionState) > 0 {
		if err := assignInputField(input, "ConnectionState", _cloudwatcheventsConnectionState); err != nil {
			log.Errorf("invalid --connection-state: %s", err.Error())
			return
		}
	}
	if len(_cloudwatcheventsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatcheventsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatcheventsNamePrefix) > 0 {
		input.NamePrefix = aws.String(_cloudwatcheventsNamePrefix)
	}
	if len(_cloudwatcheventsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatcheventsNextToken)
	}

	if resp, err := client.ListConnections(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the event buses in your account, including the default event bus,
// custom event buses, and partner event buses.
func cloudwatchevents_ListEventBuses(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.ListEventBusesInput{}

	if len(_cloudwatcheventsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatcheventsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatcheventsNamePrefix) > 0 {
		input.NamePrefix = aws.String(_cloudwatcheventsNamePrefix)
	}
	if len(_cloudwatcheventsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatcheventsNextToken)
	}

	if resp, err := client.ListEventBuses(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// You can use this to see all the partner event sources that have been shared
// with your Amazon Web Services account. For more information about partner event
// sources, see [CreateEventBus].
//
// [CreateEventBus]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_CreateEventBus.html
func cloudwatchevents_ListEventSources(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.ListEventSourcesInput{}

	if len(_cloudwatcheventsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatcheventsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatcheventsNamePrefix) > 0 {
		input.NamePrefix = aws.String(_cloudwatcheventsNamePrefix)
	}
	if len(_cloudwatcheventsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatcheventsNextToken)
	}

	if resp, err := client.ListEventSources(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// An SaaS partner can use this operation to display the Amazon Web Services
// account ID that a particular partner event source name is associated with. This
// operation is not used by Amazon Web Services customers.
func cloudwatchevents_ListPartnerEventSourceAccounts(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.ListPartnerEventSourceAccountsInput{
		// EventSourceName: *string, // Required
	}

	if len(_cloudwatcheventsEventSourceName) > 0 {
		input.EventSourceName = aws.String(_cloudwatcheventsEventSourceName)
	}
	if len(_cloudwatcheventsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatcheventsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatcheventsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatcheventsNextToken)
	}

	if resp, err := client.ListPartnerEventSourceAccounts(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// An SaaS partner can use this operation to list all the partner event source
// names that they have created. This operation is not used by Amazon Web Services
// customers.
func cloudwatchevents_ListPartnerEventSources(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.ListPartnerEventSourcesInput{
		// NamePrefix: *string, // Required
	}

	if len(_cloudwatcheventsNamePrefix) > 0 {
		input.NamePrefix = aws.String(_cloudwatcheventsNamePrefix)
	}
	if len(_cloudwatcheventsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatcheventsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatcheventsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatcheventsNextToken)
	}

	if resp, err := client.ListPartnerEventSources(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists your replays. You can either list all the replays or you can provide a
// prefix to match to the replay names. Filter parameters are exclusive.
func cloudwatchevents_ListReplays(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.ListReplaysInput{}

	if len(_cloudwatcheventsEventSourceArn) > 0 {
		input.EventSourceArn = aws.String(_cloudwatcheventsEventSourceArn)
	}
	if len(_cloudwatcheventsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatcheventsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatcheventsNamePrefix) > 0 {
		input.NamePrefix = aws.String(_cloudwatcheventsNamePrefix)
	}
	if len(_cloudwatcheventsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatcheventsNextToken)
	}
	if len(_cloudwatcheventsState) > 0 {
		if err := assignInputField(input, "State", _cloudwatcheventsState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListReplays(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the rules for the specified target. You can see which of the rules in
// Amazon EventBridge can invoke a specific target in your account.
func cloudwatchevents_ListRuleNamesByTarget(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.ListRuleNamesByTargetInput{
		// TargetArn: *string, // Required
	}

	if len(_cloudwatcheventsTargetArn) > 0 {
		input.TargetArn = aws.String(_cloudwatcheventsTargetArn)
	}
	if len(_cloudwatcheventsEventBusName) > 0 {
		input.EventBusName = aws.String(_cloudwatcheventsEventBusName)
	}
	if len(_cloudwatcheventsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatcheventsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatcheventsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatcheventsNextToken)
	}

	if resp, err := client.ListRuleNamesByTarget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists your Amazon EventBridge rules. You can either list all the rules or you
// can provide a prefix to match to the rule names.
//
// ListRules does not list the targets of a rule. To see the targets associated
// with a rule, use [ListTargetsByRule].
//
// [ListTargetsByRule]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_ListTargetsByRule.html
func cloudwatchevents_ListRules(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.ListRulesInput{}

	if len(_cloudwatcheventsEventBusName) > 0 {
		input.EventBusName = aws.String(_cloudwatcheventsEventBusName)
	}
	if len(_cloudwatcheventsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatcheventsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatcheventsNamePrefix) > 0 {
		input.NamePrefix = aws.String(_cloudwatcheventsNamePrefix)
	}
	if len(_cloudwatcheventsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatcheventsNextToken)
	}

	if resp, err := client.ListRules(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays the tags associated with an EventBridge resource. In EventBridge,
// rules and event buses can be tagged.
func cloudwatchevents_ListTagsForResource(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_cloudwatcheventsResourceARN) > 0 {
		input.ResourceARN = aws.String(_cloudwatcheventsResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the targets assigned to the specified rule.
func cloudwatchevents_ListTargetsByRule(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.ListTargetsByRuleInput{
		// Rule: *string, // Required
	}

	if len(_cloudwatcheventsRule) > 0 {
		input.Rule = aws.String(_cloudwatcheventsRule)
	}
	if len(_cloudwatcheventsEventBusName) > 0 {
		input.EventBusName = aws.String(_cloudwatcheventsEventBusName)
	}
	if len(_cloudwatcheventsLimit) > 0 {
		if err := assignInputField(input, "Limit", _cloudwatcheventsLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_cloudwatcheventsNextToken) > 0 {
		input.NextToken = aws.String(_cloudwatcheventsNextToken)
	}

	if resp, err := client.ListTargetsByRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends custom events to Amazon EventBridge so that they can be matched to rules.
func cloudwatchevents_PutEvents(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.PutEventsInput{
		// Entries: []types.PutEventsRequestEntry, // Required
	}

	if len(_cloudwatcheventsEntries) > 0 {
		if err := assignInputField(input, "Entries", _cloudwatcheventsEntries); err != nil {
			log.Errorf("invalid --entries: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutEvents(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is used by SaaS partners to write events to a customer's partner event
// bus. Amazon Web Services customers do not use this operation.
func cloudwatchevents_PutPartnerEvents(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.PutPartnerEventsInput{
		// Entries: []types.PutPartnerEventsRequestEntry, // Required
	}

	if len(_cloudwatcheventsEntries) > 0 {
		if err := assignInputField(input, "Entries", _cloudwatcheventsEntries); err != nil {
			log.Errorf("invalid --entries: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutPartnerEvents(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Running PutPermission permits the specified Amazon Web Services account or
// Amazon Web Services organization to put events to the specified event bus.
// Amazon EventBridge (CloudWatch Events) rules in your account are triggered by
// these events arriving to an event bus in your account.
//
// For another account to send events to your account, that external account must
// have an EventBridge rule with your account's event bus as a target.
//
// To enable multiple Amazon Web Services accounts to put events to your event
// bus, run PutPermission once for each of these accounts. Or, if all the accounts
// are members of the same Amazon Web Services organization, you can run
// PutPermission once specifying Principal as "*" and specifying the Amazon Web
// Services organization ID in Condition , to grant permissions to all accounts in
// that organization.
//
// If you grant permissions using an organization, then accounts in that
// organization must specify a RoleArn with proper permissions when they use
// PutTarget to add your account's event bus as a target. For more information, see [Sending and Receiving Events Between Amazon Web Services Accounts]
// in the Amazon EventBridge User Guide.
//
// The permission policy on the event bus cannot exceed 10 KB in size.
//
// [Sending and Receiving Events Between Amazon Web Services Accounts]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-cross-account-event-delivery.html
func cloudwatchevents_PutPermission(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.PutPermissionInput{}

	if len(_cloudwatcheventsAction) > 0 {
		input.Action = aws.String(_cloudwatcheventsAction)
	}
	if len(_cloudwatcheventsCondition) > 0 {
		if err := assignInputField(input, "Condition", _cloudwatcheventsCondition); err != nil {
			log.Errorf("invalid --condition: %s", err.Error())
			return
		}
	}
	if len(_cloudwatcheventsEventBusName) > 0 {
		input.EventBusName = aws.String(_cloudwatcheventsEventBusName)
	}
	if len(_cloudwatcheventsPolicy) > 0 {
		input.Policy = aws.String(_cloudwatcheventsPolicy)
	}
	if len(_cloudwatcheventsPrincipal) > 0 {
		input.Principal = aws.String(_cloudwatcheventsPrincipal)
	}
	if len(_cloudwatcheventsStatementId) > 0 {
		input.StatementId = aws.String(_cloudwatcheventsStatementId)
	}

	if resp, err := client.PutPermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates the specified rule. Rules are enabled by default, or based
// on value of the state. You can disable a rule using [DisableRule].
//
// A single rule watches for events from a single event bus. Events generated by
// Amazon Web Services services go to your account's default event bus. Events
// generated by SaaS partner services or applications go to the matching partner
// event bus. If you have custom applications or services, you can specify whether
// their events go to your default event bus or a custom event bus that you have
// created. For more information, see [CreateEventBus].
//
// If you are updating an existing rule, the rule is replaced with what you
// specify in this PutRule command. If you omit arguments in PutRule , the old
// values for those arguments are not kept. Instead, they are replaced with null
// values.
//
// When you create or update a rule, incoming events might not immediately start
// matching to new or updated rules. Allow a short period of time for changes to
// take effect.
//
// A rule must contain at least an EventPattern or ScheduleExpression. Rules with
// EventPatterns are triggered when a matching event is observed. Rules with
// ScheduleExpressions self-trigger based on the given schedule. A rule can have
// both an EventPattern and a ScheduleExpression, in which case the rule triggers
// on matching events as well as on a schedule.
//
// When you initially create a rule, you can optionally assign one or more tags to
// the rule. Tags can help you organize and categorize your resources. You can also
// use them to scope user permissions, by granting a user permission to access or
// change only rules with certain tag values. To use the PutRule operation and
// assign tags, you must have both the events:PutRule and events:TagResource
// permissions.
//
// If you are updating an existing rule, any tags you specify in the PutRule
// operation are ignored. To update the tags of an existing rule, use [TagResource]and [UntagResource].
//
// Most services in Amazon Web Services treat : or / as the same character in
// Amazon Resource Names (ARNs). However, EventBridge uses an exact match in event
// patterns and rules. Be sure to use the correct ARN characters when creating
// event patterns so that they match the ARN syntax in the event you want to match.
//
// In EventBridge, it is possible to create rules that lead to infinite loops,
// where a rule is fired repeatedly. For example, a rule might detect that ACLs
// have changed on an S3 bucket, and trigger software to change them to the desired
// state. If the rule is not written carefully, the subsequent change to the ACLs
// fires the rule again, creating an infinite loop.
//
// To prevent this, write the rules so that the triggered actions do not re-fire
// the same rule. For example, your rule could fire only if ACLs are found to be in
// a bad state, instead of after any change.
//
// An infinite loop can quickly cause higher than expected charges. We recommend
// that you use budgeting, which alerts you when charges exceed your specified
// limit. For more information, see [Managing Your Costs with Budgets].
//
// [CreateEventBus]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_CreateEventBus.html
// [TagResource]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_TagResource.html
// [UntagResource]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_UntagResource.html
// [DisableRule]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_DisableRule.html
// [Managing Your Costs with Budgets]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/budgets-managing-costs.html
func cloudwatchevents_PutRule(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.PutRuleInput{
		// Name: *string, // Required
	}

	if len(_cloudwatcheventsName) > 0 {
		input.Name = aws.String(_cloudwatcheventsName)
	}
	if len(_cloudwatcheventsDescription) > 0 {
		input.Description = aws.String(_cloudwatcheventsDescription)
	}
	if len(_cloudwatcheventsEventBusName) > 0 {
		input.EventBusName = aws.String(_cloudwatcheventsEventBusName)
	}
	if len(_cloudwatcheventsEventPattern) > 0 {
		input.EventPattern = aws.String(_cloudwatcheventsEventPattern)
	}
	if len(_cloudwatcheventsRoleArn) > 0 {
		input.RoleArn = aws.String(_cloudwatcheventsRoleArn)
	}
	if len(_cloudwatcheventsScheduleExpression) > 0 {
		input.ScheduleExpression = aws.String(_cloudwatcheventsScheduleExpression)
	}
	if len(_cloudwatcheventsState) > 0 {
		if err := assignInputField(input, "State", _cloudwatcheventsState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}
	if len(_cloudwatcheventsTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudwatcheventsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds the specified targets to the specified rule, or updates the targets if
// they are already associated with the rule.
//
// Targets are the resources that are invoked when a rule is triggered.
//
// You can configure the following as targets for Events:
//
// [API destination]
//
// - Amazon API Gateway REST API endpoints
//
// - API Gateway
//
// - Batch job queue
//
// - CloudWatch Logs group
//
// - CodeBuild project
//
// - CodePipeline
//
// - Amazon EC2 CreateSnapshot API call
//
// - Amazon EC2 RebootInstances API call
//
// - Amazon EC2 StopInstances API call
//
// - Amazon EC2 TerminateInstances API call
//
// - Amazon ECS tasks
//
// - Event bus in a different Amazon Web Services account or Region.
//
// You can use an event bus in the US East (N. Virginia) us-east-1, US West
//
// (Oregon) us-west-2, or Europe (Ireland) eu-west-1 Regions as a target for a
// rule.
//
// - Firehose delivery stream (Firehose)
//
// - Inspector assessment template (Amazon Inspector)
//
// - Kinesis stream (Kinesis Data Stream)
//
// - Lambda function
//
// - Redshift clusters (Data API statement execution)
//
// - Amazon SNS topic
//
// - Amazon SQS queues (includes FIFO queues
//
// - SSM Automation
//
// - SSM OpsItem
//
// - SSM Run Command
//
// - Step Functions state machines
//
// Creating rules with built-in targets is supported only in the Amazon Web
// Services Management Console. The built-in targets are EC2 CreateSnapshot API
// call , EC2 RebootInstances API call , EC2 StopInstances API call , and EC2
// TerminateInstances API call .
//
// For some target types, PutTargets provides target-specific parameters. If the
// target is a Kinesis data stream, you can optionally specify which shard the
// event goes to by using the KinesisParameters argument. To invoke a command on
// multiple EC2 instances with one rule, you can use the RunCommandParameters
// field.
//
// To be able to make API calls against the resources that you own, Amazon
// EventBridge needs the appropriate permissions. For Lambda and Amazon SNS
// resources, EventBridge relies on resource-based policies. For EC2 instances,
// Kinesis Data Streams, Step Functions state machines and API Gateway REST APIs,
// EventBridge relies on IAM roles that you specify in the RoleARN argument in
// PutTargets . For more information, see [Authentication and Access Control] in the Amazon EventBridge User Guide.
//
// If another Amazon Web Services account is in the same region and has granted
// you permission (using PutPermission ), you can send events to that account. Set
// that account's event bus as a target of the rules in your account. To send the
// matched events to the other account, specify that account's event bus as the Arn
// value when you run PutTargets . If your account sends events to another account,
// your account is charged for each sent event. Each event sent to another account
// is charged as a custom event. The account receiving the event is not charged.
// For more information, see [Amazon EventBridge Pricing].
//
// Input , InputPath , and InputTransformer are not available with PutTarget if
// the target is an event bus of a different Amazon Web Services account.
//
// If you are setting the event bus of another account as the target, and that
// account granted permission to your account through an organization instead of
// directly by the account ID, then you must specify a RoleArn with proper
// permissions in the Target structure. For more information, see [Sending and Receiving Events Between Amazon Web Services Accounts] in the Amazon
// EventBridge User Guide.
//
// For more information about enabling cross-account events, see [PutPermission].
//
// Input, InputPath, and InputTransformer are mutually exclusive and optional
// parameters of a target. When a rule is triggered due to a matched event:
//
// - If none of the following arguments are specified for a target, then the
// entire event is passed to the target in JSON format (unless the target is Amazon
// EC2 Run Command or Amazon ECS task, in which case nothing from the event is
// passed to the target).
//
// - If Input is specified in the form of valid JSON, then the matched event is
// overridden with this constant.
//
// - If InputPath is specified in the form of JSONPath (for example, $.detail ),
// then only the part of the event specified in the path is passed to the target
// (for example, only the detail part of the event is passed).
//
// - If InputTransformer is specified, then one or more specified JSONPaths are
// extracted from the event and used as values in a template that you specify as
// the input to the target.
//
// When you specify InputPath or InputTransformer , you must use JSON dot notation,
// not bracket notation.
//
// When you add targets to a rule and the associated rule triggers soon after, new
// or updated targets might not be immediately invoked. Allow a short period of
// time for changes to take effect.
//
// This action can partially fail if too many requests are made at the same time.
// If that happens, FailedEntryCount is non-zero in the response and each entry in
// FailedEntries provides the ID of the failed target and the error code.
//
// [Authentication and Access Control]: https://docs.aws.amazon.com/eventbridge/latest/userguide/auth-and-access-control-eventbridge.html
// [Amazon EventBridge Pricing]: http://aws.amazon.com/eventbridge/pricing/
// [API destination]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-api-destinations.html
// [Sending and Receiving Events Between Amazon Web Services Accounts]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-cross-account-event-delivery.html
// [PutPermission]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_PutPermission.html
func cloudwatchevents_PutTargets(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.PutTargetsInput{
		// Rule: *string, // Required
		// Targets: []types.Target, // Required
	}

	if len(_cloudwatcheventsRule) > 0 {
		input.Rule = aws.String(_cloudwatcheventsRule)
	}
	if len(_cloudwatcheventsTargets) > 0 {
		if err := assignInputField(input, "Targets", _cloudwatcheventsTargets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}
	if len(_cloudwatcheventsEventBusName) > 0 {
		input.EventBusName = aws.String(_cloudwatcheventsEventBusName)
	}

	if resp, err := client.PutTargets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Revokes the permission of another Amazon Web Services account to be able to put
// events to the specified event bus. Specify the account to revoke by the
// StatementId value that you associated with the account when you granted it
// permission with PutPermission . You can find the StatementId by using [DescribeEventBus].
//
// [DescribeEventBus]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_DescribeEventBus.html
func cloudwatchevents_RemovePermission(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.RemovePermissionInput{}

	if len(_cloudwatcheventsEventBusName) > 0 {
		input.EventBusName = aws.String(_cloudwatcheventsEventBusName)
	}
	if len(_cloudwatcheventsRemoveAllPermissions) > 0 {
		if err := assignInputField(input, "RemoveAllPermissions", _cloudwatcheventsRemoveAllPermissions); err != nil {
			log.Errorf("invalid --remove-all-permissions: %s", err.Error())
			return
		}
	}
	if len(_cloudwatcheventsStatementId) > 0 {
		input.StatementId = aws.String(_cloudwatcheventsStatementId)
	}

	if resp, err := client.RemovePermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified targets from the specified rule. When the rule is
// triggered, those targets are no longer be invoked.
//
// When you remove a target, when the associated rule triggers, removed targets
// might continue to be invoked. Allow a short period of time for changes to take
// effect.
//
// This action can partially fail if too many requests are made at the same time.
// If that happens, FailedEntryCount is non-zero in the response and each entry in
// FailedEntries provides the ID of the failed target and the error code.
func cloudwatchevents_RemoveTargets(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.RemoveTargetsInput{
		// Ids: []string, // Required
		// Rule: *string, // Required
	}

	if len(_cloudwatcheventsIds) > 0 {
		input.Ids = append([]string(nil), _cloudwatcheventsIds...)
	}
	if len(_cloudwatcheventsRule) > 0 {
		input.Rule = aws.String(_cloudwatcheventsRule)
	}
	if len(_cloudwatcheventsEventBusName) > 0 {
		input.EventBusName = aws.String(_cloudwatcheventsEventBusName)
	}
	if len(_cloudwatcheventsForce) > 0 {
		if err := assignInputField(input, "Force", _cloudwatcheventsForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}

	if resp, err := client.RemoveTargets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the specified replay. Events are not necessarily replayed in the exact
// same order that they were added to the archive. A replay processes events to
// replay based on the time in the event, and replays them using 1 minute
// intervals. If you specify an EventStartTime and an EventEndTime that covers a
// 20 minute time range, the events are replayed from the first minute of that 20
// minute range first. Then the events from the second minute are replayed. You can
// use DescribeReplay to determine the progress of a replay. The value returned
// for EventLastReplayedTime indicates the time within the specified time range
// associated with the last event replayed.
func cloudwatchevents_StartReplay(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.StartReplayInput{
		// Destination: *types.ReplayDestination, // Required
		// EventEndTime: *time.Time, // Required
		// EventSourceArn: *string, // Required
		// EventStartTime: *time.Time, // Required
		// ReplayName: *string, // Required
	}

	if len(_cloudwatcheventsDestination) > 0 {
		if err := assignInputField(input, "Destination", _cloudwatcheventsDestination); err != nil {
			log.Errorf("invalid --destination: %s", err.Error())
			return
		}
	}
	if len(_cloudwatcheventsEventEndTime) > 0 {
		if err := assignInputField(input, "EventEndTime", _cloudwatcheventsEventEndTime); err != nil {
			log.Errorf("invalid --event-end-time: %s", err.Error())
			return
		}
	}
	if len(_cloudwatcheventsEventSourceArn) > 0 {
		input.EventSourceArn = aws.String(_cloudwatcheventsEventSourceArn)
	}
	if len(_cloudwatcheventsEventStartTime) > 0 {
		if err := assignInputField(input, "EventStartTime", _cloudwatcheventsEventStartTime); err != nil {
			log.Errorf("invalid --event-start-time: %s", err.Error())
			return
		}
	}
	if len(_cloudwatcheventsReplayName) > 0 {
		input.ReplayName = aws.String(_cloudwatcheventsReplayName)
	}
	if len(_cloudwatcheventsDescription) > 0 {
		input.Description = aws.String(_cloudwatcheventsDescription)
	}

	if resp, err := client.StartReplay(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns one or more tags (key-value pairs) to the specified EventBridge
// resource. Tags can help you organize and categorize your resources. You can also
// use them to scope user permissions by granting a user permission to access or
// change only resources with certain tag values. In EventBridge, rules and event
// buses can be tagged.
//
// Tags don't have any semantic meaning to Amazon Web Services and are interpreted
// strictly as strings of characters.
//
// You can use the TagResource action with a resource that already has tags. If
// you specify a new tag key, this tag is appended to the list of tags associated
// with the resource. If you specify a tag key that is already associated with the
// resource, the new tag value that you specify replaces the previous value for
// that tag.
//
// You can associate as many as 50 tags with a resource.
func cloudwatchevents_TagResource(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_cloudwatcheventsResourceARN) > 0 {
		input.ResourceARN = aws.String(_cloudwatcheventsResourceARN)
	}
	if len(_cloudwatcheventsTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudwatcheventsTags); err != nil {
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

// Tests whether the specified event pattern matches the provided event.
// Most services in Amazon Web Services treat : or / as the same character in
// Amazon Resource Names (ARNs). However, EventBridge uses an exact match in event
// patterns and rules. Be sure to use the correct ARN characters when creating
// event patterns so that they match the ARN syntax in the event you want to match.
func cloudwatchevents_TestEventPattern(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.TestEventPatternInput{
		// Event: *string, // Required
		// EventPattern: *string, // Required
	}

	if len(_cloudwatcheventsEvent) > 0 {
		input.Event = aws.String(_cloudwatcheventsEvent)
	}
	if len(_cloudwatcheventsEventPattern) > 0 {
		input.EventPattern = aws.String(_cloudwatcheventsEventPattern)
	}

	if resp, err := client.TestEventPattern(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes one or more tags from the specified EventBridge resource. In Amazon
// EventBridge (CloudWatch Events), rules and event buses can be tagged.
func cloudwatchevents_UntagResource(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_cloudwatcheventsResourceARN) > 0 {
		input.ResourceARN = aws.String(_cloudwatcheventsResourceARN)
	}
	if len(_cloudwatcheventsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _cloudwatcheventsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an API destination.
func cloudwatchevents_UpdateApiDestination(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.UpdateApiDestinationInput{
		// Name: *string, // Required
	}

	if len(_cloudwatcheventsName) > 0 {
		input.Name = aws.String(_cloudwatcheventsName)
	}
	if len(_cloudwatcheventsConnectionArn) > 0 {
		input.ConnectionArn = aws.String(_cloudwatcheventsConnectionArn)
	}
	if len(_cloudwatcheventsDescription) > 0 {
		input.Description = aws.String(_cloudwatcheventsDescription)
	}
	if len(_cloudwatcheventsHttpMethod) > 0 {
		if err := assignInputField(input, "HttpMethod", _cloudwatcheventsHttpMethod); err != nil {
			log.Errorf("invalid --http-method: %s", err.Error())
			return
		}
	}
	if len(_cloudwatcheventsInvocationEndpoint) > 0 {
		input.InvocationEndpoint = aws.String(_cloudwatcheventsInvocationEndpoint)
	}
	if len(_cloudwatcheventsInvocationRateLimitPerSecond) > 0 {
		if err := assignInputField(input, "InvocationRateLimitPerSecond", _cloudwatcheventsInvocationRateLimitPerSecond); err != nil {
			log.Errorf("invalid --invocation-rate-limit-per-second: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateApiDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified archive.
func cloudwatchevents_UpdateArchive(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.UpdateArchiveInput{
		// ArchiveName: *string, // Required
	}

	if len(_cloudwatcheventsArchiveName) > 0 {
		input.ArchiveName = aws.String(_cloudwatcheventsArchiveName)
	}
	if len(_cloudwatcheventsDescription) > 0 {
		input.Description = aws.String(_cloudwatcheventsDescription)
	}
	if len(_cloudwatcheventsEventPattern) > 0 {
		input.EventPattern = aws.String(_cloudwatcheventsEventPattern)
	}
	if len(_cloudwatcheventsRetentionDays) > 0 {
		if err := assignInputField(input, "RetentionDays", _cloudwatcheventsRetentionDays); err != nil {
			log.Errorf("invalid --retention-days: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateArchive(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates settings for a connection.
func cloudwatchevents_UpdateConnection(cfg aws.Config, client *cloudwatchevents.Client) {
	input := &cloudwatchevents.UpdateConnectionInput{
		// Name: *string, // Required
	}

	if len(_cloudwatcheventsName) > 0 {
		input.Name = aws.String(_cloudwatcheventsName)
	}
	if len(_cloudwatcheventsAuthParameters) > 0 {
		if err := assignInputField(input, "AuthParameters", _cloudwatcheventsAuthParameters); err != nil {
			log.Errorf("invalid --auth-parameters: %s", err.Error())
			return
		}
	}
	if len(_cloudwatcheventsAuthorizationType) > 0 {
		if err := assignInputField(input, "AuthorizationType", _cloudwatcheventsAuthorizationType); err != nil {
			log.Errorf("invalid --authorization-type: %s", err.Error())
			return
		}
	}
	if len(_cloudwatcheventsDescription) > 0 {
		input.Description = aws.String(_cloudwatcheventsDescription)
	}

	if resp, err := client.UpdateConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_cloudwatcheventsCmd)
	_cloudwatcheventsCmd.Flags().SortFlags = false

	_cloudwatcheventsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_cloudwatcheventsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_cloudwatcheventsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsAccount, "account", "", "", "Account")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsAction, "action", "", "", "Action")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsArchiveName, "archive-name", "", "", "Archive Name")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsAuthParameters, "auth-parameters", "", "", "Auth Parameters")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsAuthorizationType, "authorization-type", "", "", "Authorization Type")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsCondition, "condition", "", "", "Condition")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsConnectionArn, "connection-arn", "", "", "Connection ARN")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsConnectionState, "connection-state", "", "", "Connection State")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsDescription, "description", "", "", "Description")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsDestination, "destination", "", "", "Destination")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsEntries, "entries", "", "", "Entries")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsEvent, "event", "", "", "Event")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsEventBusName, "event-bus-name", "", "", "Event Bus Name")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsEventEndTime, "event-end-time", "", "", "Event End Time")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsEventPattern, "event-pattern", "", "", "Event Pattern")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsEventSourceArn, "event-source-arn", "", "", "Event Source ARN")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsEventSourceName, "event-source-name", "", "", "Event Source Name")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsEventStartTime, "event-start-time", "", "", "Event Start Time")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsForce, "force", "", "", "Force")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsHttpMethod, "http-method", "", "", "HTTP Method")
	_cloudwatcheventsCmd.Flags().StringSliceVarP(&_cloudwatcheventsIds, "ids", "", nil, "Ids")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsInvocationEndpoint, "invocation-endpoint", "", "", "Invocation Endpoint")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsInvocationRateLimitPerSecond, "invocation-rate-limit-per-second", "", "", "Invocation Rate Limit Per Second")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsLimit, "limit", "", "", "Limit")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsName, "name", "", "", "Name")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsNamePrefix, "name-prefix", "", "", "Name Prefix")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsNextToken, "next-token", "", "", "Next Token")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsPolicy, "policy", "", "", "Policy")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsPrincipal, "principal", "", "", "Principal")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsRemoveAllPermissions, "remove-all-permissions", "", "", "Remove All Permissions")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsReplayName, "replay-name", "", "", "Replay Name")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsResourceARN, "resource-arn", "", "", "Resource ARN")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsRetentionDays, "retention-days", "", "", "Retention Days")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsRoleArn, "role-arn", "", "", "Role ARN")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsRule, "rule", "", "", "Rule")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsScheduleExpression, "schedule-expression", "", "", "Schedule Expression")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsState, "state", "", "", "State")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsStatementId, "statement-id", "", "", "Statement ID")
	_cloudwatcheventsCmd.Flags().StringSliceVarP(&_cloudwatcheventsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsTags, "tags", "", "", "Tags")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsTargetArn, "target-arn", "", "", "Target ARN")
	_cloudwatcheventsCmd.Flags().StringVarP(&_cloudwatcheventsTargets, "targets", "", "", "Targets")

	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsActivateEventSource, "activate-event-source", "", false, "Activate Event Source")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsCancelReplay, "cancel-replay", "", false, "Cancel Replay")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsCreateApiDestination, "create-api-destination", "", false, "Create API Destination")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsCreateArchive, "create-archive", "", false, "Create Archive")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsCreateConnection, "create-connection", "", false, "Create Connection")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsCreateEventBus, "create-event-bus", "", false, "Create Event Bus")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsCreatePartnerEventSource, "create-partner-event-source", "", false, "Create Partner Event Source")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsDeactivateEventSource, "deactivate-event-source", "", false, "Deactivate Event Source")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsDeauthorizeConnection, "deauthorize-connection", "", false, "Deauthorize Connection")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsDeleteApiDestination, "delete-api-destination", "", false, "Delete API Destination")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsDeleteArchive, "delete-archive", "", false, "Delete Archive")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsDeleteConnection, "delete-connection", "", false, "Delete Connection")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsDeleteEventBus, "delete-event-bus", "", false, "Delete Event Bus")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsDeletePartnerEventSource, "delete-partner-event-source", "", false, "Delete Partner Event Source")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsDeleteRule, "delete-rule", "", false, "Delete Rule")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsDescribeApiDestination, "describe-api-destination", "", false, "Describe API Destination")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsDescribeArchive, "describe-archive", "", false, "Describe Archive")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsDescribeConnection, "describe-connection", "", false, "Describe Connection")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsDescribeEventBus, "describe-event-bus", "", false, "Describe Event Bus")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsDescribeEventSource, "describe-event-source", "", false, "Describe Event Source")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsDescribePartnerEventSource, "describe-partner-event-source", "", false, "Describe Partner Event Source")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsDescribeReplay, "describe-replay", "", false, "Describe Replay")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsDescribeRule, "describe-rule", "", false, "Describe Rule")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsDisableRule, "disable-rule", "", false, "Disable Rule")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsEnableRule, "enable-rule", "", false, "Enable Rule")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsListApiDestinations, "list-api-destinations", "", false, "List API Destinations")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsListArchives, "list-archives", "", false, "List Archives")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsListConnections, "list-connections", "", false, "List Connections")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsListEventBuses, "list-event-buses", "", false, "List Event Buses")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsListEventSources, "list-event-sources", "", false, "List Event Sources")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsListPartnerEventSourceAccounts, "list-partner-event-source-accounts", "", false, "List Partner Event Source Accounts")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsListPartnerEventSources, "list-partner-event-sources", "", false, "List Partner Event Sources")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsListReplays, "list-replays", "", false, "List Replays")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsListRuleNamesByTarget, "list-rule-names-by-target", "", false, "List Rule Names By Target")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsListRules, "list-rules", "", false, "List Rules")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsListTargetsByRule, "list-targets-by-rule", "", false, "List Targets By Rule")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsPutEvents, "put-events", "", false, "Put Events")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsPutPartnerEvents, "put-partner-events", "", false, "Put Partner Events")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsPutPermission, "put-permission", "", false, "Put Permission")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsPutRule, "put-rule", "", false, "Put Rule")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsPutTargets, "put-targets", "", false, "Put Targets")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsRemovePermission, "remove-permission", "", false, "Remove Permission")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsRemoveTargets, "remove-targets", "", false, "Remove Targets")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsStartReplay, "start-replay", "", false, "Start Replay")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsTagResource, "tag-resource", "", false, "Tag Resource")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsTestEventPattern, "test-event-pattern", "", false, "Test Event Pattern")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsUntagResource, "untag-resource", "", false, "Untag Resource")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsUpdateApiDestination, "update-api-destination", "", false, "Update API Destination")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsUpdateArchive, "update-archive", "", false, "Update Archive")
	_cloudwatcheventsCmd.Flags().BoolVarP(&_cloudwatcheventsUpdateConnection, "update-connection", "", false, "Update Connection")

}
