package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// eventbridgeCmd represents the eventbridge command
var _eventbridgeCmd = &cobra.Command{
	Use:   "eventbridge",
	Short: "AWS eventbridge CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := eventbridge.NewFromConfig(cfg)
		if _eventbridgeActivateEventSource {
			eventbridge_ActivateEventSource(cfg, client)
			return
		}
		if _eventbridgeCancelReplay {
			eventbridge_CancelReplay(cfg, client)
			return
		}
		if _eventbridgeCreateApiDestination {
			eventbridge_CreateApiDestination(cfg, client)
			return
		}
		if _eventbridgeCreateArchive {
			eventbridge_CreateArchive(cfg, client)
			return
		}
		if _eventbridgeCreateConnection {
			eventbridge_CreateConnection(cfg, client)
			return
		}
		if _eventbridgeCreateEndpoint {
			eventbridge_CreateEndpoint(cfg, client)
			return
		}
		if _eventbridgeCreateEventBus {
			eventbridge_CreateEventBus(cfg, client)
			return
		}
		if _eventbridgeCreatePartnerEventSource {
			eventbridge_CreatePartnerEventSource(cfg, client)
			return
		}
		if _eventbridgeDeactivateEventSource {
			eventbridge_DeactivateEventSource(cfg, client)
			return
		}
		if _eventbridgeDeauthorizeConnection {
			eventbridge_DeauthorizeConnection(cfg, client)
			return
		}
		if _eventbridgeDeleteApiDestination {
			eventbridge_DeleteApiDestination(cfg, client)
			return
		}
		if _eventbridgeDeleteArchive {
			eventbridge_DeleteArchive(cfg, client)
			return
		}
		if _eventbridgeDeleteConnection {
			eventbridge_DeleteConnection(cfg, client)
			return
		}
		if _eventbridgeDeleteEndpoint {
			eventbridge_DeleteEndpoint(cfg, client)
			return
		}
		if _eventbridgeDeleteEventBus {
			eventbridge_DeleteEventBus(cfg, client)
			return
		}
		if _eventbridgeDeletePartnerEventSource {
			eventbridge_DeletePartnerEventSource(cfg, client)
			return
		}
		if _eventbridgeDeleteRule {
			eventbridge_DeleteRule(cfg, client)
			return
		}
		if _eventbridgeDescribeApiDestination {
			eventbridge_DescribeApiDestination(cfg, client)
			return
		}
		if _eventbridgeDescribeArchive {
			eventbridge_DescribeArchive(cfg, client)
			return
		}
		if _eventbridgeDescribeConnection {
			eventbridge_DescribeConnection(cfg, client)
			return
		}
		if _eventbridgeDescribeEndpoint {
			eventbridge_DescribeEndpoint(cfg, client)
			return
		}
		if _eventbridgeDescribeEventBus {
			eventbridge_DescribeEventBus(cfg, client)
			return
		}
		if _eventbridgeDescribeEventSource {
			eventbridge_DescribeEventSource(cfg, client)
			return
		}
		if _eventbridgeDescribePartnerEventSource {
			eventbridge_DescribePartnerEventSource(cfg, client)
			return
		}
		if _eventbridgeDescribeReplay {
			eventbridge_DescribeReplay(cfg, client)
			return
		}
		if _eventbridgeDescribeRule {
			eventbridge_DescribeRule(cfg, client)
			return
		}
		if _eventbridgeDisableRule {
			eventbridge_DisableRule(cfg, client)
			return
		}
		if _eventbridgeEnableRule {
			eventbridge_EnableRule(cfg, client)
			return
		}
		if _eventbridgeListApiDestinations {
			eventbridge_ListApiDestinations(cfg, client)
			return
		}
		if _eventbridgeListArchives {
			eventbridge_ListArchives(cfg, client)
			return
		}
		if _eventbridgeListConnections {
			eventbridge_ListConnections(cfg, client)
			return
		}
		if _eventbridgeListEndpoints {
			eventbridge_ListEndpoints(cfg, client)
			return
		}
		if _eventbridgeListEventBuses {
			eventbridge_ListEventBuses(cfg, client)
			return
		}
		if _eventbridgeListEventSources {
			eventbridge_ListEventSources(cfg, client)
			return
		}
		if _eventbridgeListPartnerEventSourceAccounts {
			eventbridge_ListPartnerEventSourceAccounts(cfg, client)
			return
		}
		if _eventbridgeListPartnerEventSources {
			eventbridge_ListPartnerEventSources(cfg, client)
			return
		}
		if _eventbridgeListReplays {
			eventbridge_ListReplays(cfg, client)
			return
		}
		if _eventbridgeListRuleNamesByTarget {
			eventbridge_ListRuleNamesByTarget(cfg, client)
			return
		}
		if _eventbridgeListRules {
			eventbridge_ListRules(cfg, client)
			return
		}
		if _eventbridgeListTagsForResource {
			eventbridge_ListTagsForResource(cfg, client)
			return
		}
		if _eventbridgeListTargetsByRule {
			eventbridge_ListTargetsByRule(cfg, client)
			return
		}
		if _eventbridgePutEvents {
			eventbridge_PutEvents(cfg, client)
			return
		}
		if _eventbridgePutPartnerEvents {
			eventbridge_PutPartnerEvents(cfg, client)
			return
		}
		if _eventbridgePutPermission {
			eventbridge_PutPermission(cfg, client)
			return
		}
		if _eventbridgePutRule {
			eventbridge_PutRule(cfg, client)
			return
		}
		if _eventbridgePutTargets {
			eventbridge_PutTargets(cfg, client)
			return
		}
		if _eventbridgeRemovePermission {
			eventbridge_RemovePermission(cfg, client)
			return
		}
		if _eventbridgeRemoveTargets {
			eventbridge_RemoveTargets(cfg, client)
			return
		}
		if _eventbridgeStartReplay {
			eventbridge_StartReplay(cfg, client)
			return
		}
		if _eventbridgeTagResource {
			eventbridge_TagResource(cfg, client)
			return
		}
		if _eventbridgeTestEventPattern {
			eventbridge_TestEventPattern(cfg, client)
			return
		}
		if _eventbridgeUntagResource {
			eventbridge_UntagResource(cfg, client)
			return
		}
		if _eventbridgeUpdateApiDestination {
			eventbridge_UpdateApiDestination(cfg, client)
			return
		}
		if _eventbridgeUpdateArchive {
			eventbridge_UpdateArchive(cfg, client)
			return
		}
		if _eventbridgeUpdateConnection {
			eventbridge_UpdateConnection(cfg, client)
			return
		}
		if _eventbridgeUpdateEndpoint {
			eventbridge_UpdateEndpoint(cfg, client)
			return
		}
		if _eventbridgeUpdateEventBus {
			eventbridge_UpdateEventBus(cfg, client)
			return
		}

	},
}

var (
	_eventbridgeActivateEventSource            bool
	_eventbridgeCancelReplay                   bool
	_eventbridgeCreateApiDestination           bool
	_eventbridgeCreateArchive                  bool
	_eventbridgeCreateConnection               bool
	_eventbridgeCreateEndpoint                 bool
	_eventbridgeCreateEventBus                 bool
	_eventbridgeCreatePartnerEventSource       bool
	_eventbridgeDeactivateEventSource          bool
	_eventbridgeDeauthorizeConnection          bool
	_eventbridgeDeleteApiDestination           bool
	_eventbridgeDeleteArchive                  bool
	_eventbridgeDeleteConnection               bool
	_eventbridgeDeleteEndpoint                 bool
	_eventbridgeDeleteEventBus                 bool
	_eventbridgeDeletePartnerEventSource       bool
	_eventbridgeDeleteRule                     bool
	_eventbridgeDescribeApiDestination         bool
	_eventbridgeDescribeArchive                bool
	_eventbridgeDescribeConnection             bool
	_eventbridgeDescribeEndpoint               bool
	_eventbridgeDescribeEventBus               bool
	_eventbridgeDescribeEventSource            bool
	_eventbridgeDescribePartnerEventSource     bool
	_eventbridgeDescribeReplay                 bool
	_eventbridgeDescribeRule                   bool
	_eventbridgeDisableRule                    bool
	_eventbridgeEnableRule                     bool
	_eventbridgeListApiDestinations            bool
	_eventbridgeListArchives                   bool
	_eventbridgeListConnections                bool
	_eventbridgeListEndpoints                  bool
	_eventbridgeListEventBuses                 bool
	_eventbridgeListEventSources               bool
	_eventbridgeListPartnerEventSourceAccounts bool
	_eventbridgeListPartnerEventSources        bool
	_eventbridgeListReplays                    bool
	_eventbridgeListRuleNamesByTarget          bool
	_eventbridgeListRules                      bool
	_eventbridgeListTagsForResource            bool
	_eventbridgeListTargetsByRule              bool
	_eventbridgePutEvents                      bool
	_eventbridgePutPartnerEvents               bool
	_eventbridgePutPermission                  bool
	_eventbridgePutRule                        bool
	_eventbridgePutTargets                     bool
	_eventbridgeRemovePermission               bool
	_eventbridgeRemoveTargets                  bool
	_eventbridgeStartReplay                    bool
	_eventbridgeTagResource                    bool
	_eventbridgeTestEventPattern               bool
	_eventbridgeUntagResource                  bool
	_eventbridgeUpdateApiDestination           bool
	_eventbridgeUpdateArchive                  bool
	_eventbridgeUpdateConnection               bool
	_eventbridgeUpdateEndpoint                 bool
	_eventbridgeUpdateEventBus                 bool

	_eventbridgeAccount                          string
	_eventbridgeAction                           string
	_eventbridgeArchiveName                      string
	_eventbridgeAuthParameters                   string
	_eventbridgeAuthorizationType                string
	_eventbridgeCondition                        string
	_eventbridgeConnectionArn                    string
	_eventbridgeConnectionState                  string
	_eventbridgeDeadLetterConfig                 string
	_eventbridgeDescription                      string
	_eventbridgeDestination                      string
	_eventbridgeEndpointId                       string
	_eventbridgeEntries                          string
	_eventbridgeEvent                            string
	_eventbridgeEventBusName                     string
	_eventbridgeEventBuses                       string
	_eventbridgeEventEndTime                     string
	_eventbridgeEventPattern                     string
	_eventbridgeEventSourceArn                   string
	_eventbridgeEventSourceName                  string
	_eventbridgeEventStartTime                   string
	_eventbridgeForce                            string
	_eventbridgeHomeRegion                       string
	_eventbridgeHttpMethod                       string
	_eventbridgeIds                              []string
	_eventbridgeInvocationConnectivityParameters string
	_eventbridgeInvocationEndpoint               string
	_eventbridgeInvocationRateLimitPerSecond     string
	_eventbridgeKmsKeyIdentifier                 string
	_eventbridgeLimit                            string
	_eventbridgeLogConfig                        string
	_eventbridgeMaxResults                       string
	_eventbridgeName                             string
	_eventbridgeNamePrefix                       string
	_eventbridgeNextToken                        string
	_eventbridgePolicy                           string
	_eventbridgePrincipal                        string
	_eventbridgeRemoveAllPermissions             string
	_eventbridgeReplayName                       string
	_eventbridgeReplicationConfig                string
	_eventbridgeResourceARN                      string
	_eventbridgeRetentionDays                    string
	_eventbridgeRoleArn                          string
	_eventbridgeRoutingConfig                    string
	_eventbridgeRule                             string
	_eventbridgeScheduleExpression               string
	_eventbridgeState                            string
	_eventbridgeStatementId                      string
	_eventbridgeTagKeys                          []string
	_eventbridgeTags                             string
	_eventbridgeTargetArn                        string
	_eventbridgeTargets                          string
)

// Activates a partner event source that has been deactivated. Once activated,
// your matching event bus will start receiving events from the event source.
func eventbridge_ActivateEventSource(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.ActivateEventSourceInput{
		// Name: *string, // Required
	}

	if len(_eventbridgeName) > 0 {
		input.Name = aws.String(_eventbridgeName)
	}

	if resp, err := client.ActivateEventSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels the specified replay.
func eventbridge_CancelReplay(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.CancelReplayInput{
		// ReplayName: *string, // Required
	}

	if len(_eventbridgeReplayName) > 0 {
		input.ReplayName = aws.String(_eventbridgeReplayName)
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
//
// API destinations do not support private destinations, such as interface VPC
// endpoints.
//
// For more information, see [API destinations] in the EventBridge User Guide.
//
// [API destinations]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-api-destinations.html
func eventbridge_CreateApiDestination(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.CreateApiDestinationInput{
		// ConnectionArn: *string, // Required
		// HttpMethod: types.ApiDestinationHttpMethod, // Required
		// InvocationEndpoint: *string, // Required
		// Name: *string, // Required
	}

	if len(_eventbridgeConnectionArn) > 0 {
		input.ConnectionArn = aws.String(_eventbridgeConnectionArn)
	}
	if len(_eventbridgeHttpMethod) > 0 {
		if err := assignInputField(input, "HttpMethod", _eventbridgeHttpMethod); err != nil {
			log.Errorf("invalid --http-method: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeInvocationEndpoint) > 0 {
		input.InvocationEndpoint = aws.String(_eventbridgeInvocationEndpoint)
	}
	if len(_eventbridgeName) > 0 {
		input.Name = aws.String(_eventbridgeName)
	}
	if len(_eventbridgeDescription) > 0 {
		input.Description = aws.String(_eventbridgeDescription)
	}
	if len(_eventbridgeInvocationRateLimitPerSecond) > 0 {
		if err := assignInputField(input, "InvocationRateLimitPerSecond", _eventbridgeInvocationRateLimitPerSecond); err != nil {
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
//
// If you have specified that EventBridge use a customer managed key for
// encrypting the source event bus, we strongly recommend you also specify a
// customer managed key for any archives for the event bus as well.
//
// For more information, see [Encrypting archives] in the Amazon EventBridge User Guide.
//
// [Encrypting archives]: https://docs.aws.amazon.com/eventbridge/latest/userguide/encryption-archives.html
func eventbridge_CreateArchive(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.CreateArchiveInput{
		// ArchiveName: *string, // Required
		// EventSourceArn: *string, // Required
	}

	if len(_eventbridgeArchiveName) > 0 {
		input.ArchiveName = aws.String(_eventbridgeArchiveName)
	}
	if len(_eventbridgeEventSourceArn) > 0 {
		input.EventSourceArn = aws.String(_eventbridgeEventSourceArn)
	}
	if len(_eventbridgeDescription) > 0 {
		input.Description = aws.String(_eventbridgeDescription)
	}
	if len(_eventbridgeEventPattern) > 0 {
		input.EventPattern = aws.String(_eventbridgeEventPattern)
	}
	if len(_eventbridgeKmsKeyIdentifier) > 0 {
		input.KmsKeyIdentifier = aws.String(_eventbridgeKmsKeyIdentifier)
	}
	if len(_eventbridgeRetentionDays) > 0 {
		if err := assignInputField(input, "RetentionDays", _eventbridgeRetentionDays); err != nil {
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
//
// For more information, see [Connections for endpoint targets] in the Amazon EventBridge User Guide.
//
// [Connections for endpoint targets]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-target-connection.html
func eventbridge_CreateConnection(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.CreateConnectionInput{
		// AuthParameters: *types.CreateConnectionAuthRequestParameters, // Required
		// AuthorizationType: types.ConnectionAuthorizationType, // Required
		// Name: *string, // Required
	}

	if len(_eventbridgeAuthParameters) > 0 {
		if err := assignInputField(input, "AuthParameters", _eventbridgeAuthParameters); err != nil {
			log.Errorf("invalid --auth-parameters: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeAuthorizationType) > 0 {
		if err := assignInputField(input, "AuthorizationType", _eventbridgeAuthorizationType); err != nil {
			log.Errorf("invalid --authorization-type: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeName) > 0 {
		input.Name = aws.String(_eventbridgeName)
	}
	if len(_eventbridgeDescription) > 0 {
		input.Description = aws.String(_eventbridgeDescription)
	}
	if len(_eventbridgeInvocationConnectivityParameters) > 0 {
		if err := assignInputField(input, "InvocationConnectivityParameters", _eventbridgeInvocationConnectivityParameters); err != nil {
			log.Errorf("invalid --invocation-connectivity-parameters: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeKmsKeyIdentifier) > 0 {
		input.KmsKeyIdentifier = aws.String(_eventbridgeKmsKeyIdentifier)
	}

	if resp, err := client.CreateConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a global endpoint. Global endpoints improve your application's
// availability by making it regional-fault tolerant. To do this, you define a
// primary and secondary Region with event buses in each Region. You also create a
// Amazon Route 53 health check that will tell EventBridge to route events to the
// secondary Region when an "unhealthy" state is encountered and events will be
// routed back to the primary Region when the health check reports a "healthy"
// state.
func eventbridge_CreateEndpoint(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.CreateEndpointInput{
		// EventBuses: []types.EndpointEventBus, // Required
		// Name: *string, // Required
		// RoutingConfig: *types.RoutingConfig, // Required
	}

	if len(_eventbridgeEventBuses) > 0 {
		if err := assignInputField(input, "EventBuses", _eventbridgeEventBuses); err != nil {
			log.Errorf("invalid --event-buses: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeName) > 0 {
		input.Name = aws.String(_eventbridgeName)
	}
	if len(_eventbridgeRoutingConfig) > 0 {
		if err := assignInputField(input, "RoutingConfig", _eventbridgeRoutingConfig); err != nil {
			log.Errorf("invalid --routing-config: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeDescription) > 0 {
		input.Description = aws.String(_eventbridgeDescription)
	}
	if len(_eventbridgeReplicationConfig) > 0 {
		if err := assignInputField(input, "ReplicationConfig", _eventbridgeReplicationConfig); err != nil {
			log.Errorf("invalid --replication-config: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeRoleArn) > 0 {
		input.RoleArn = aws.String(_eventbridgeRoleArn)
	}

	if resp, err := client.CreateEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new event bus within your account. This can be a custom event bus
// which you can use to receive events from your custom applications and services,
// or it can be a partner event bus which can be matched to a partner event source.
func eventbridge_CreateEventBus(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.CreateEventBusInput{
		// Name: *string, // Required
	}

	if len(_eventbridgeName) > 0 {
		input.Name = aws.String(_eventbridgeName)
	}
	if len(_eventbridgeDeadLetterConfig) > 0 {
		if err := assignInputField(input, "DeadLetterConfig", _eventbridgeDeadLetterConfig); err != nil {
			log.Errorf("invalid --dead-letter-config: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeDescription) > 0 {
		input.Description = aws.String(_eventbridgeDescription)
	}
	if len(_eventbridgeEventSourceName) > 0 {
		input.EventSourceName = aws.String(_eventbridgeEventSourceName)
	}
	if len(_eventbridgeKmsKeyIdentifier) > 0 {
		input.KmsKeyIdentifier = aws.String(_eventbridgeKmsKeyIdentifier)
	}
	if len(_eventbridgeLogConfig) > 0 {
		if err := assignInputField(input, "LogConfig", _eventbridgeLogConfig); err != nil {
			log.Errorf("invalid --log-config: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeTags) > 0 {
		if err := assignInputField(input, "Tags", _eventbridgeTags); err != nil {
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
// - partner_name is determined during partner registration, and identifies the
// partner to Amazon Web Services customers.
//
// - event_namespace is determined by the partner, and is a way for the partner
// to categorize their events.
//
// - event_name is determined by the partner, and should uniquely identify an
// event-generating resource within the partner system.
//
// The event_name must be unique across all Amazon Web Services customers. This is
//
// because the event source is a shared resource between the partner and customer
// accounts, and each partner event source unique in the partner account.
//
// The combination of event_namespace and event_name should help Amazon Web
// Services customers decide whether to create an event bus to receive these
// events.
func eventbridge_CreatePartnerEventSource(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.CreatePartnerEventSourceInput{
		// Account: *string, // Required
		// Name: *string, // Required
	}

	if len(_eventbridgeAccount) > 0 {
		input.Account = aws.String(_eventbridgeAccount)
	}
	if len(_eventbridgeName) > 0 {
		input.Name = aws.String(_eventbridgeName)
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
func eventbridge_DeactivateEventSource(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.DeactivateEventSourceInput{
		// Name: *string, // Required
	}

	if len(_eventbridgeName) > 0 {
		input.Name = aws.String(_eventbridgeName)
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
func eventbridge_DeauthorizeConnection(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.DeauthorizeConnectionInput{
		// Name: *string, // Required
	}

	if len(_eventbridgeName) > 0 {
		input.Name = aws.String(_eventbridgeName)
	}

	if resp, err := client.DeauthorizeConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified API destination.
func eventbridge_DeleteApiDestination(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.DeleteApiDestinationInput{
		// Name: *string, // Required
	}

	if len(_eventbridgeName) > 0 {
		input.Name = aws.String(_eventbridgeName)
	}

	if resp, err := client.DeleteApiDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified archive.
func eventbridge_DeleteArchive(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.DeleteArchiveInput{
		// ArchiveName: *string, // Required
	}

	if len(_eventbridgeArchiveName) > 0 {
		input.ArchiveName = aws.String(_eventbridgeArchiveName)
	}

	if resp, err := client.DeleteArchive(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a connection.
func eventbridge_DeleteConnection(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.DeleteConnectionInput{
		// Name: *string, // Required
	}

	if len(_eventbridgeName) > 0 {
		input.Name = aws.String(_eventbridgeName)
	}

	if resp, err := client.DeleteConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an existing global endpoint. For more information about global
// endpoints, see [Making applications Regional-fault tolerant with global endpoints and event replication]in the Amazon EventBridge User Guide .
//
// [Making applications Regional-fault tolerant with global endpoints and event replication]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-global-endpoints.html
func eventbridge_DeleteEndpoint(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.DeleteEndpointInput{
		// Name: *string, // Required
	}

	if len(_eventbridgeName) > 0 {
		input.Name = aws.String(_eventbridgeName)
	}

	if resp, err := client.DeleteEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified custom event bus or partner event bus. All rules
// associated with this event bus need to be deleted. You can't delete your
// account's default event bus.
func eventbridge_DeleteEventBus(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.DeleteEventBusInput{
		// Name: *string, // Required
	}

	if len(_eventbridgeName) > 0 {
		input.Name = aws.String(_eventbridgeName)
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
func eventbridge_DeletePartnerEventSource(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.DeletePartnerEventSourceInput{
		// Account: *string, // Required
		// Name: *string, // Required
	}

	if len(_eventbridgeAccount) > 0 {
		input.Account = aws.String(_eventbridgeAccount)
	}
	if len(_eventbridgeName) > 0 {
		input.Name = aws.String(_eventbridgeName)
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
func eventbridge_DeleteRule(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.DeleteRuleInput{
		// Name: *string, // Required
	}

	if len(_eventbridgeName) > 0 {
		input.Name = aws.String(_eventbridgeName)
	}
	if len(_eventbridgeEventBusName) > 0 {
		input.EventBusName = aws.String(_eventbridgeEventBusName)
	}
	if len(_eventbridgeForce) > 0 {
		if err := assignInputField(input, "Force", _eventbridgeForce); err != nil {
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
func eventbridge_DescribeApiDestination(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.DescribeApiDestinationInput{
		// Name: *string, // Required
	}

	if len(_eventbridgeName) > 0 {
		input.Name = aws.String(_eventbridgeName)
	}

	if resp, err := client.DescribeApiDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about an archive.
func eventbridge_DescribeArchive(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.DescribeArchiveInput{
		// ArchiveName: *string, // Required
	}

	if len(_eventbridgeArchiveName) > 0 {
		input.ArchiveName = aws.String(_eventbridgeArchiveName)
	}

	if resp, err := client.DescribeArchive(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about a connection.
func eventbridge_DescribeConnection(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.DescribeConnectionInput{
		// Name: *string, // Required
	}

	if len(_eventbridgeName) > 0 {
		input.Name = aws.String(_eventbridgeName)
	}

	if resp, err := client.DescribeConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the information about an existing global endpoint. For more information
// about global endpoints, see [Making applications Regional-fault tolerant with global endpoints and event replication]in the Amazon EventBridge User Guide .
//
// [Making applications Regional-fault tolerant with global endpoints and event replication]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-global-endpoints.html
func eventbridge_DescribeEndpoint(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.DescribeEndpointInput{
		// Name: *string, // Required
	}

	if len(_eventbridgeName) > 0 {
		input.Name = aws.String(_eventbridgeName)
	}
	if len(_eventbridgeHomeRegion) > 0 {
		input.HomeRegion = aws.String(_eventbridgeHomeRegion)
	}

	if resp, err := client.DescribeEndpoint(context.TODO(), input); err != nil {
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
func eventbridge_DescribeEventBus(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.DescribeEventBusInput{}

	if len(_eventbridgeName) > 0 {
		input.Name = aws.String(_eventbridgeName)
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
func eventbridge_DescribeEventSource(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.DescribeEventSourceInput{
		// Name: *string, // Required
	}

	if len(_eventbridgeName) > 0 {
		input.Name = aws.String(_eventbridgeName)
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
func eventbridge_DescribePartnerEventSource(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.DescribePartnerEventSourceInput{
		// Name: *string, // Required
	}

	if len(_eventbridgeName) > 0 {
		input.Name = aws.String(_eventbridgeName)
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
func eventbridge_DescribeReplay(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.DescribeReplayInput{
		// ReplayName: *string, // Required
	}

	if len(_eventbridgeReplayName) > 0 {
		input.ReplayName = aws.String(_eventbridgeReplayName)
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
func eventbridge_DescribeRule(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.DescribeRuleInput{
		// Name: *string, // Required
	}

	if len(_eventbridgeName) > 0 {
		input.Name = aws.String(_eventbridgeName)
	}
	if len(_eventbridgeEventBusName) > 0 {
		input.EventBusName = aws.String(_eventbridgeEventBusName)
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
func eventbridge_DisableRule(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.DisableRuleInput{
		// Name: *string, // Required
	}

	if len(_eventbridgeName) > 0 {
		input.Name = aws.String(_eventbridgeName)
	}
	if len(_eventbridgeEventBusName) > 0 {
		input.EventBusName = aws.String(_eventbridgeEventBusName)
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
func eventbridge_EnableRule(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.EnableRuleInput{
		// Name: *string, // Required
	}

	if len(_eventbridgeName) > 0 {
		input.Name = aws.String(_eventbridgeName)
	}
	if len(_eventbridgeEventBusName) > 0 {
		input.EventBusName = aws.String(_eventbridgeEventBusName)
	}

	if resp, err := client.EnableRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of API destination in the account in the current Region.
func eventbridge_ListApiDestinations(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.ListApiDestinationsInput{}

	if len(_eventbridgeConnectionArn) > 0 {
		input.ConnectionArn = aws.String(_eventbridgeConnectionArn)
	}
	if len(_eventbridgeLimit) > 0 {
		if err := assignInputField(input, "Limit", _eventbridgeLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeNamePrefix) > 0 {
		input.NamePrefix = aws.String(_eventbridgeNamePrefix)
	}
	if len(_eventbridgeNextToken) > 0 {
		input.NextToken = aws.String(_eventbridgeNextToken)
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
func eventbridge_ListArchives(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.ListArchivesInput{}

	if len(_eventbridgeEventSourceArn) > 0 {
		input.EventSourceArn = aws.String(_eventbridgeEventSourceArn)
	}
	if len(_eventbridgeLimit) > 0 {
		if err := assignInputField(input, "Limit", _eventbridgeLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeNamePrefix) > 0 {
		input.NamePrefix = aws.String(_eventbridgeNamePrefix)
	}
	if len(_eventbridgeNextToken) > 0 {
		input.NextToken = aws.String(_eventbridgeNextToken)
	}
	if len(_eventbridgeState) > 0 {
		if err := assignInputField(input, "State", _eventbridgeState); err != nil {
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
func eventbridge_ListConnections(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.ListConnectionsInput{}

	if len(_eventbridgeConnectionState) > 0 {
		if err := assignInputField(input, "ConnectionState", _eventbridgeConnectionState); err != nil {
			log.Errorf("invalid --connection-state: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeLimit) > 0 {
		if err := assignInputField(input, "Limit", _eventbridgeLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeNamePrefix) > 0 {
		input.NamePrefix = aws.String(_eventbridgeNamePrefix)
	}
	if len(_eventbridgeNextToken) > 0 {
		input.NextToken = aws.String(_eventbridgeNextToken)
	}

	if resp, err := client.ListConnections(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List the global endpoints associated with this account. For more information
// about global endpoints, see [Making applications Regional-fault tolerant with global endpoints and event replication]in the Amazon EventBridge User Guide .
//
// [Making applications Regional-fault tolerant with global endpoints and event replication]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-global-endpoints.html
func eventbridge_ListEndpoints(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.ListEndpointsInput{}

	if len(_eventbridgeHomeRegion) > 0 {
		input.HomeRegion = aws.String(_eventbridgeHomeRegion)
	}
	if len(_eventbridgeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _eventbridgeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeNamePrefix) > 0 {
		input.NamePrefix = aws.String(_eventbridgeNamePrefix)
	}
	if len(_eventbridgeNextToken) > 0 {
		input.NextToken = aws.String(_eventbridgeNextToken)
	}

	if resp, err := client.ListEndpoints(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the event buses in your account, including the default event bus,
// custom event buses, and partner event buses.
func eventbridge_ListEventBuses(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.ListEventBusesInput{}

	if len(_eventbridgeLimit) > 0 {
		if err := assignInputField(input, "Limit", _eventbridgeLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeNamePrefix) > 0 {
		input.NamePrefix = aws.String(_eventbridgeNamePrefix)
	}
	if len(_eventbridgeNextToken) > 0 {
		input.NextToken = aws.String(_eventbridgeNextToken)
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
func eventbridge_ListEventSources(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.ListEventSourcesInput{}

	if len(_eventbridgeLimit) > 0 {
		if err := assignInputField(input, "Limit", _eventbridgeLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeNamePrefix) > 0 {
		input.NamePrefix = aws.String(_eventbridgeNamePrefix)
	}
	if len(_eventbridgeNextToken) > 0 {
		input.NextToken = aws.String(_eventbridgeNextToken)
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
func eventbridge_ListPartnerEventSourceAccounts(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.ListPartnerEventSourceAccountsInput{
		// EventSourceName: *string, // Required
	}

	if len(_eventbridgeEventSourceName) > 0 {
		input.EventSourceName = aws.String(_eventbridgeEventSourceName)
	}
	if len(_eventbridgeLimit) > 0 {
		if err := assignInputField(input, "Limit", _eventbridgeLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeNextToken) > 0 {
		input.NextToken = aws.String(_eventbridgeNextToken)
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
func eventbridge_ListPartnerEventSources(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.ListPartnerEventSourcesInput{
		// NamePrefix: *string, // Required
	}

	if len(_eventbridgeNamePrefix) > 0 {
		input.NamePrefix = aws.String(_eventbridgeNamePrefix)
	}
	if len(_eventbridgeLimit) > 0 {
		if err := assignInputField(input, "Limit", _eventbridgeLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeNextToken) > 0 {
		input.NextToken = aws.String(_eventbridgeNextToken)
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
func eventbridge_ListReplays(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.ListReplaysInput{}

	if len(_eventbridgeEventSourceArn) > 0 {
		input.EventSourceArn = aws.String(_eventbridgeEventSourceArn)
	}
	if len(_eventbridgeLimit) > 0 {
		if err := assignInputField(input, "Limit", _eventbridgeLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeNamePrefix) > 0 {
		input.NamePrefix = aws.String(_eventbridgeNamePrefix)
	}
	if len(_eventbridgeNextToken) > 0 {
		input.NextToken = aws.String(_eventbridgeNextToken)
	}
	if len(_eventbridgeState) > 0 {
		if err := assignInputField(input, "State", _eventbridgeState); err != nil {
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
//
// The maximum number of results per page for requests is 100.
func eventbridge_ListRuleNamesByTarget(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.ListRuleNamesByTargetInput{
		// TargetArn: *string, // Required
	}

	if len(_eventbridgeTargetArn) > 0 {
		input.TargetArn = aws.String(_eventbridgeTargetArn)
	}
	if len(_eventbridgeEventBusName) > 0 {
		input.EventBusName = aws.String(_eventbridgeEventBusName)
	}
	if len(_eventbridgeLimit) > 0 {
		if err := assignInputField(input, "Limit", _eventbridgeLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeNextToken) > 0 {
		input.NextToken = aws.String(_eventbridgeNextToken)
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
// The maximum number of results per page for requests is 100.
//
// ListRules does not list the targets of a rule. To see the targets associated
// with a rule, use [ListTargetsByRule].
//
// [ListTargetsByRule]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_ListTargetsByRule.html
func eventbridge_ListRules(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.ListRulesInput{}

	if len(_eventbridgeEventBusName) > 0 {
		input.EventBusName = aws.String(_eventbridgeEventBusName)
	}
	if len(_eventbridgeLimit) > 0 {
		if err := assignInputField(input, "Limit", _eventbridgeLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeNamePrefix) > 0 {
		input.NamePrefix = aws.String(_eventbridgeNamePrefix)
	}
	if len(_eventbridgeNextToken) > 0 {
		input.NextToken = aws.String(_eventbridgeNextToken)
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
func eventbridge_ListTagsForResource(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_eventbridgeResourceARN) > 0 {
		input.ResourceARN = aws.String(_eventbridgeResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the targets assigned to the specified rule.
// The maximum number of results per page for requests is 100.
func eventbridge_ListTargetsByRule(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.ListTargetsByRuleInput{
		// Rule: *string, // Required
	}

	if len(_eventbridgeRule) > 0 {
		input.Rule = aws.String(_eventbridgeRule)
	}
	if len(_eventbridgeEventBusName) > 0 {
		input.EventBusName = aws.String(_eventbridgeEventBusName)
	}
	if len(_eventbridgeLimit) > 0 {
		if err := assignInputField(input, "Limit", _eventbridgeLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeNextToken) > 0 {
		input.NextToken = aws.String(_eventbridgeNextToken)
	}

	if resp, err := client.ListTargetsByRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends custom events to Amazon EventBridge so that they can be matched to rules.
// You can batch multiple event entries into one request for efficiency. However,
// the total entry size must be less than 256KB. You can calculate the entry size
// before you send the events. For more information, see [Calculating PutEvents event entry size]in the Amazon EventBridge
// User Guide .
//
// PutEvents accepts the data in JSON format. For the JSON number (integer) data
// type, the constraints are: a minimum value of -9,223,372,036,854,775,808 and a
// maximum value of 9,223,372,036,854,775,807.
//
// PutEvents will only process nested JSON up to 1000 levels deep.
//
// [Calculating PutEvents event entry size]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-putevents.html#eb-putevent-size
func eventbridge_PutEvents(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.PutEventsInput{
		// Entries: []types.PutEventsRequestEntry, // Required
	}

	if len(_eventbridgeEntries) > 0 {
		if err := assignInputField(input, "Entries", _eventbridgeEntries); err != nil {
			log.Errorf("invalid --entries: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeEndpointId) > 0 {
		input.EndpointId = aws.String(_eventbridgeEndpointId)
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
//
// For information on calculating event batch size, see [Calculating EventBridge PutEvents event entry size] in the EventBridge User
// Guide.
//
// [Calculating EventBridge PutEvents event entry size]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-putevent-size.html
func eventbridge_PutPartnerEvents(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.PutPartnerEventsInput{
		// Entries: []types.PutPartnerEventsRequestEntry, // Required
	}

	if len(_eventbridgeEntries) > 0 {
		if err := assignInputField(input, "Entries", _eventbridgeEntries); err != nil {
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
// Amazon EventBridge rules in your account are triggered by these events arriving
// to an event bus in your account.
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
func eventbridge_PutPermission(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.PutPermissionInput{}

	if len(_eventbridgeAction) > 0 {
		input.Action = aws.String(_eventbridgeAction)
	}
	if len(_eventbridgeCondition) > 0 {
		if err := assignInputField(input, "Condition", _eventbridgeCondition); err != nil {
			log.Errorf("invalid --condition: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeEventBusName) > 0 {
		input.EventBusName = aws.String(_eventbridgeEventBusName)
	}
	if len(_eventbridgePolicy) > 0 {
		input.Policy = aws.String(_eventbridgePolicy)
	}
	if len(_eventbridgePrincipal) > 0 {
		input.Principal = aws.String(_eventbridgePrincipal)
	}
	if len(_eventbridgeStatementId) > 0 {
		input.StatementId = aws.String(_eventbridgeStatementId)
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
// To create a rule that filters for management events from Amazon Web Services
// services, see [Receiving read-only management events from Amazon Web Services services]in the EventBridge User Guide.
//
// [CreateEventBus]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_CreateEventBus.html
// [TagResource]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_TagResource.html
// [UntagResource]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_UntagResource.html
// [DisableRule]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_DisableRule.html
// [Managing Your Costs with Budgets]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/budgets-managing-costs.html
// [Receiving read-only management events from Amazon Web Services services]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-service-event-cloudtrail.html#eb-service-event-cloudtrail-management
func eventbridge_PutRule(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.PutRuleInput{
		// Name: *string, // Required
	}

	if len(_eventbridgeName) > 0 {
		input.Name = aws.String(_eventbridgeName)
	}
	if len(_eventbridgeDescription) > 0 {
		input.Description = aws.String(_eventbridgeDescription)
	}
	if len(_eventbridgeEventBusName) > 0 {
		input.EventBusName = aws.String(_eventbridgeEventBusName)
	}
	if len(_eventbridgeEventPattern) > 0 {
		input.EventPattern = aws.String(_eventbridgeEventPattern)
	}
	if len(_eventbridgeRoleArn) > 0 {
		input.RoleArn = aws.String(_eventbridgeRoleArn)
	}
	if len(_eventbridgeScheduleExpression) > 0 {
		input.ScheduleExpression = aws.String(_eventbridgeScheduleExpression)
	}
	if len(_eventbridgeState) > 0 {
		if err := assignInputField(input, "State", _eventbridgeState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeTags) > 0 {
		if err := assignInputField(input, "Tags", _eventbridgeTags); err != nil {
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
// The maximum number of entries per request is 10.
//
// Each rule can have up to five (5) targets associated with it at one time.
//
// For a list of services you can configure as targets for events, see [EventBridge targets] in the
// Amazon EventBridge User Guide .
//
// Creating rules with built-in targets is supported only in the Amazon Web
// Services Management Console. The built-in targets are:
//
// - Amazon EBS CreateSnapshot API call
//
// - Amazon EC2 RebootInstances API call
//
// - Amazon EC2 StopInstances API call
//
// - Amazon EC2 TerminateInstances API call
//
// For some target types, PutTargets provides target-specific parameters. If the
// target is a Kinesis data stream, you can optionally specify which shard the
// event goes to by using the KinesisParameters argument. To invoke a command on
// multiple EC2 instances with one rule, you can use the RunCommandParameters
// field.
//
// To be able to make API calls against the resources that you own, Amazon
// EventBridge needs the appropriate permissions:
//
// - For Lambda and Amazon SNS resources, EventBridge relies on resource-based
// policies.
//
// - For EC2 instances, Kinesis Data Streams, Step Functions state machines and
// API Gateway APIs, EventBridge relies on IAM roles that you specify in the
// RoleARN argument in PutTargets .
//
// For more information, see [Authentication and Access Control] in the Amazon EventBridge User Guide .
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
// If you have an IAM role on a cross-account event bus target, a PutTargets call
// without a role on the same target (same Id and Arn ) will not remove the role.
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
// [Sending and Receiving Events Between Amazon Web Services Accounts]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-cross-account-event-delivery.html
// [PutPermission]: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_PutPermission.html
// [EventBridge targets]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-targets.html
func eventbridge_PutTargets(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.PutTargetsInput{
		// Rule: *string, // Required
		// Targets: []types.Target, // Required
	}

	if len(_eventbridgeRule) > 0 {
		input.Rule = aws.String(_eventbridgeRule)
	}
	if len(_eventbridgeTargets) > 0 {
		if err := assignInputField(input, "Targets", _eventbridgeTargets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeEventBusName) > 0 {
		input.EventBusName = aws.String(_eventbridgeEventBusName)
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
func eventbridge_RemovePermission(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.RemovePermissionInput{}

	if len(_eventbridgeEventBusName) > 0 {
		input.EventBusName = aws.String(_eventbridgeEventBusName)
	}
	if len(_eventbridgeRemoveAllPermissions) > 0 {
		if err := assignInputField(input, "RemoveAllPermissions", _eventbridgeRemoveAllPermissions); err != nil {
			log.Errorf("invalid --remove-all-permissions: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeStatementId) > 0 {
		input.StatementId = aws.String(_eventbridgeStatementId)
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
// A successful execution of RemoveTargets doesn't guarantee all targets are
// removed from the rule, it means that the target(s) listed in the request are
// removed.
//
// When you remove a target, when the associated rule triggers, removed targets
// might continue to be invoked. Allow a short period of time for changes to take
// effect.
//
// This action can partially fail if too many requests are made at the same time.
// If that happens, FailedEntryCount is non-zero in the response and each entry in
// FailedEntries provides the ID of the failed target and the error code.
//
// The maximum number of entries per request is 10.
func eventbridge_RemoveTargets(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.RemoveTargetsInput{
		// Ids: []string, // Required
		// Rule: *string, // Required
	}

	if len(_eventbridgeIds) > 0 {
		input.Ids = append([]string(nil), _eventbridgeIds...)
	}
	if len(_eventbridgeRule) > 0 {
		input.Rule = aws.String(_eventbridgeRule)
	}
	if len(_eventbridgeEventBusName) > 0 {
		input.EventBusName = aws.String(_eventbridgeEventBusName)
	}
	if len(_eventbridgeForce) > 0 {
		if err := assignInputField(input, "Force", _eventbridgeForce); err != nil {
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
func eventbridge_StartReplay(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.StartReplayInput{
		// Destination: *types.ReplayDestination, // Required
		// EventEndTime: *time.Time, // Required
		// EventSourceArn: *string, // Required
		// EventStartTime: *time.Time, // Required
		// ReplayName: *string, // Required
	}

	if len(_eventbridgeDestination) > 0 {
		if err := assignInputField(input, "Destination", _eventbridgeDestination); err != nil {
			log.Errorf("invalid --destination: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeEventEndTime) > 0 {
		if err := assignInputField(input, "EventEndTime", _eventbridgeEventEndTime); err != nil {
			log.Errorf("invalid --event-end-time: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeEventSourceArn) > 0 {
		input.EventSourceArn = aws.String(_eventbridgeEventSourceArn)
	}
	if len(_eventbridgeEventStartTime) > 0 {
		if err := assignInputField(input, "EventStartTime", _eventbridgeEventStartTime); err != nil {
			log.Errorf("invalid --event-start-time: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeReplayName) > 0 {
		input.ReplayName = aws.String(_eventbridgeReplayName)
	}
	if len(_eventbridgeDescription) > 0 {
		input.Description = aws.String(_eventbridgeDescription)
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
func eventbridge_TagResource(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_eventbridgeResourceARN) > 0 {
		input.ResourceARN = aws.String(_eventbridgeResourceARN)
	}
	if len(_eventbridgeTags) > 0 {
		if err := assignInputField(input, "Tags", _eventbridgeTags); err != nil {
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
func eventbridge_TestEventPattern(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.TestEventPatternInput{
		// Event: *string, // Required
		// EventPattern: *string, // Required
	}

	if len(_eventbridgeEvent) > 0 {
		input.Event = aws.String(_eventbridgeEvent)
	}
	if len(_eventbridgeEventPattern) > 0 {
		input.EventPattern = aws.String(_eventbridgeEventPattern)
	}

	if resp, err := client.TestEventPattern(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes one or more tags from the specified EventBridge resource. In Amazon
// EventBridge, rules and event buses can be tagged.
func eventbridge_UntagResource(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_eventbridgeResourceARN) > 0 {
		input.ResourceARN = aws.String(_eventbridgeResourceARN)
	}
	if len(_eventbridgeTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _eventbridgeTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an API destination.
func eventbridge_UpdateApiDestination(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.UpdateApiDestinationInput{
		// Name: *string, // Required
	}

	if len(_eventbridgeName) > 0 {
		input.Name = aws.String(_eventbridgeName)
	}
	if len(_eventbridgeConnectionArn) > 0 {
		input.ConnectionArn = aws.String(_eventbridgeConnectionArn)
	}
	if len(_eventbridgeDescription) > 0 {
		input.Description = aws.String(_eventbridgeDescription)
	}
	if len(_eventbridgeHttpMethod) > 0 {
		if err := assignInputField(input, "HttpMethod", _eventbridgeHttpMethod); err != nil {
			log.Errorf("invalid --http-method: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeInvocationEndpoint) > 0 {
		input.InvocationEndpoint = aws.String(_eventbridgeInvocationEndpoint)
	}
	if len(_eventbridgeInvocationRateLimitPerSecond) > 0 {
		if err := assignInputField(input, "InvocationRateLimitPerSecond", _eventbridgeInvocationRateLimitPerSecond); err != nil {
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
func eventbridge_UpdateArchive(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.UpdateArchiveInput{
		// ArchiveName: *string, // Required
	}

	if len(_eventbridgeArchiveName) > 0 {
		input.ArchiveName = aws.String(_eventbridgeArchiveName)
	}
	if len(_eventbridgeDescription) > 0 {
		input.Description = aws.String(_eventbridgeDescription)
	}
	if len(_eventbridgeEventPattern) > 0 {
		input.EventPattern = aws.String(_eventbridgeEventPattern)
	}
	if len(_eventbridgeKmsKeyIdentifier) > 0 {
		input.KmsKeyIdentifier = aws.String(_eventbridgeKmsKeyIdentifier)
	}
	if len(_eventbridgeRetentionDays) > 0 {
		if err := assignInputField(input, "RetentionDays", _eventbridgeRetentionDays); err != nil {
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
func eventbridge_UpdateConnection(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.UpdateConnectionInput{
		// Name: *string, // Required
	}

	if len(_eventbridgeName) > 0 {
		input.Name = aws.String(_eventbridgeName)
	}
	if len(_eventbridgeAuthParameters) > 0 {
		if err := assignInputField(input, "AuthParameters", _eventbridgeAuthParameters); err != nil {
			log.Errorf("invalid --auth-parameters: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeAuthorizationType) > 0 {
		if err := assignInputField(input, "AuthorizationType", _eventbridgeAuthorizationType); err != nil {
			log.Errorf("invalid --authorization-type: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeDescription) > 0 {
		input.Description = aws.String(_eventbridgeDescription)
	}
	if len(_eventbridgeInvocationConnectivityParameters) > 0 {
		if err := assignInputField(input, "InvocationConnectivityParameters", _eventbridgeInvocationConnectivityParameters); err != nil {
			log.Errorf("invalid --invocation-connectivity-parameters: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeKmsKeyIdentifier) > 0 {
		input.KmsKeyIdentifier = aws.String(_eventbridgeKmsKeyIdentifier)
	}

	if resp, err := client.UpdateConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update an existing endpoint. For more information about global endpoints, see [Making applications Regional-fault tolerant with global endpoints and event replication]
// in the Amazon EventBridge User Guide .
//
// [Making applications Regional-fault tolerant with global endpoints and event replication]: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-global-endpoints.html
func eventbridge_UpdateEndpoint(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.UpdateEndpointInput{
		// Name: *string, // Required
	}

	if len(_eventbridgeName) > 0 {
		input.Name = aws.String(_eventbridgeName)
	}
	if len(_eventbridgeDescription) > 0 {
		input.Description = aws.String(_eventbridgeDescription)
	}
	if len(_eventbridgeEventBuses) > 0 {
		if err := assignInputField(input, "EventBuses", _eventbridgeEventBuses); err != nil {
			log.Errorf("invalid --event-buses: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeReplicationConfig) > 0 {
		if err := assignInputField(input, "ReplicationConfig", _eventbridgeReplicationConfig); err != nil {
			log.Errorf("invalid --replication-config: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeRoleArn) > 0 {
		input.RoleArn = aws.String(_eventbridgeRoleArn)
	}
	if len(_eventbridgeRoutingConfig) > 0 {
		if err := assignInputField(input, "RoutingConfig", _eventbridgeRoutingConfig); err != nil {
			log.Errorf("invalid --routing-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified event bus.
func eventbridge_UpdateEventBus(cfg aws.Config, client *eventbridge.Client) {
	input := &eventbridge.UpdateEventBusInput{}

	if len(_eventbridgeDeadLetterConfig) > 0 {
		if err := assignInputField(input, "DeadLetterConfig", _eventbridgeDeadLetterConfig); err != nil {
			log.Errorf("invalid --dead-letter-config: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeDescription) > 0 {
		input.Description = aws.String(_eventbridgeDescription)
	}
	if len(_eventbridgeKmsKeyIdentifier) > 0 {
		input.KmsKeyIdentifier = aws.String(_eventbridgeKmsKeyIdentifier)
	}
	if len(_eventbridgeLogConfig) > 0 {
		if err := assignInputField(input, "LogConfig", _eventbridgeLogConfig); err != nil {
			log.Errorf("invalid --log-config: %s", err.Error())
			return
		}
	}
	if len(_eventbridgeName) > 0 {
		input.Name = aws.String(_eventbridgeName)
	}

	if resp, err := client.UpdateEventBus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_eventbridgeCmd)
	_eventbridgeCmd.Flags().SortFlags = false

	_eventbridgeCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_eventbridgeCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_eventbridgeCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeAccount, "account", "", "", "Account")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeAction, "action", "", "", "Action")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeArchiveName, "archive-name", "", "", "Archive Name")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeAuthParameters, "auth-parameters", "", "", "Auth Parameters")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeAuthorizationType, "authorization-type", "", "", "Authorization Type")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeCondition, "condition", "", "", "Condition")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeConnectionArn, "connection-arn", "", "", "Connection ARN")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeConnectionState, "connection-state", "", "", "Connection State")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeDeadLetterConfig, "dead-letter-config", "", "", "Dead Letter Config")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeDescription, "description", "", "", "Description")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeDestination, "destination", "", "", "Destination")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeEndpointId, "endpoint-id", "", "", "Endpoint ID")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeEntries, "entries", "", "", "Entries")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeEvent, "event", "", "", "Event")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeEventBusName, "event-bus-name", "", "", "Event Bus Name")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeEventBuses, "event-buses", "", "", "Event Buses")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeEventEndTime, "event-end-time", "", "", "Event End Time")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeEventPattern, "event-pattern", "", "", "Event Pattern")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeEventSourceArn, "event-source-arn", "", "", "Event Source ARN")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeEventSourceName, "event-source-name", "", "", "Event Source Name")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeEventStartTime, "event-start-time", "", "", "Event Start Time")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeForce, "force", "", "", "Force")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeHomeRegion, "home-region", "", "", "Home Region")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeHttpMethod, "http-method", "", "", "HTTP Method")
	_eventbridgeCmd.Flags().StringSliceVarP(&_eventbridgeIds, "ids", "", nil, "Ids")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeInvocationConnectivityParameters, "invocation-connectivity-parameters", "", "", "Invocation Connectivity Parameters")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeInvocationEndpoint, "invocation-endpoint", "", "", "Invocation Endpoint")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeInvocationRateLimitPerSecond, "invocation-rate-limit-per-second", "", "", "Invocation Rate Limit Per Second")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeKmsKeyIdentifier, "kms-key-identifier", "", "", "KMS Key Identifier")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeLimit, "limit", "", "", "Limit")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeLogConfig, "log-config", "", "", "Log Config")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeMaxResults, "max-results", "", "", "Max Results")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeName, "name", "", "", "Name")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeNamePrefix, "name-prefix", "", "", "Name Prefix")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeNextToken, "next-token", "", "", "Next Token")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgePolicy, "policy", "", "", "Policy")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgePrincipal, "principal", "", "", "Principal")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeRemoveAllPermissions, "remove-all-permissions", "", "", "Remove All Permissions")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeReplayName, "replay-name", "", "", "Replay Name")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeReplicationConfig, "replication-config", "", "", "Replication Config")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeResourceARN, "resource-arn", "", "", "Resource ARN")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeRetentionDays, "retention-days", "", "", "Retention Days")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeRoleArn, "role-arn", "", "", "Role ARN")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeRoutingConfig, "routing-config", "", "", "Routing Config")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeRule, "rule", "", "", "Rule")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeScheduleExpression, "schedule-expression", "", "", "Schedule Expression")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeState, "state", "", "", "State")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeStatementId, "statement-id", "", "", "Statement ID")
	_eventbridgeCmd.Flags().StringSliceVarP(&_eventbridgeTagKeys, "tag-keys", "", nil, "Tag Keys")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeTags, "tags", "", "", "Tags")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeTargetArn, "target-arn", "", "", "Target ARN")
	_eventbridgeCmd.Flags().StringVarP(&_eventbridgeTargets, "targets", "", "", "Targets")

	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeActivateEventSource, "activate-event-source", "", false, "Activate Event Source")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeCancelReplay, "cancel-replay", "", false, "Cancel Replay")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeCreateApiDestination, "create-api-destination", "", false, "Create API Destination")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeCreateArchive, "create-archive", "", false, "Create Archive")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeCreateConnection, "create-connection", "", false, "Create Connection")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeCreateEndpoint, "create-endpoint", "", false, "Create Endpoint")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeCreateEventBus, "create-event-bus", "", false, "Create Event Bus")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeCreatePartnerEventSource, "create-partner-event-source", "", false, "Create Partner Event Source")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeDeactivateEventSource, "deactivate-event-source", "", false, "Deactivate Event Source")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeDeauthorizeConnection, "deauthorize-connection", "", false, "Deauthorize Connection")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeDeleteApiDestination, "delete-api-destination", "", false, "Delete API Destination")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeDeleteArchive, "delete-archive", "", false, "Delete Archive")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeDeleteConnection, "delete-connection", "", false, "Delete Connection")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeDeleteEndpoint, "delete-endpoint", "", false, "Delete Endpoint")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeDeleteEventBus, "delete-event-bus", "", false, "Delete Event Bus")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeDeletePartnerEventSource, "delete-partner-event-source", "", false, "Delete Partner Event Source")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeDeleteRule, "delete-rule", "", false, "Delete Rule")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeDescribeApiDestination, "describe-api-destination", "", false, "Describe API Destination")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeDescribeArchive, "describe-archive", "", false, "Describe Archive")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeDescribeConnection, "describe-connection", "", false, "Describe Connection")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeDescribeEndpoint, "describe-endpoint", "", false, "Describe Endpoint")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeDescribeEventBus, "describe-event-bus", "", false, "Describe Event Bus")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeDescribeEventSource, "describe-event-source", "", false, "Describe Event Source")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeDescribePartnerEventSource, "describe-partner-event-source", "", false, "Describe Partner Event Source")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeDescribeReplay, "describe-replay", "", false, "Describe Replay")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeDescribeRule, "describe-rule", "", false, "Describe Rule")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeDisableRule, "disable-rule", "", false, "Disable Rule")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeEnableRule, "enable-rule", "", false, "Enable Rule")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeListApiDestinations, "list-api-destinations", "", false, "List API Destinations")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeListArchives, "list-archives", "", false, "List Archives")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeListConnections, "list-connections", "", false, "List Connections")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeListEndpoints, "list-endpoints", "", false, "List Endpoints")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeListEventBuses, "list-event-buses", "", false, "List Event Buses")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeListEventSources, "list-event-sources", "", false, "List Event Sources")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeListPartnerEventSourceAccounts, "list-partner-event-source-accounts", "", false, "List Partner Event Source Accounts")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeListPartnerEventSources, "list-partner-event-sources", "", false, "List Partner Event Sources")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeListReplays, "list-replays", "", false, "List Replays")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeListRuleNamesByTarget, "list-rule-names-by-target", "", false, "List Rule Names By Target")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeListRules, "list-rules", "", false, "List Rules")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeListTargetsByRule, "list-targets-by-rule", "", false, "List Targets By Rule")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgePutEvents, "put-events", "", false, "Put Events")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgePutPartnerEvents, "put-partner-events", "", false, "Put Partner Events")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgePutPermission, "put-permission", "", false, "Put Permission")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgePutRule, "put-rule", "", false, "Put Rule")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgePutTargets, "put-targets", "", false, "Put Targets")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeRemovePermission, "remove-permission", "", false, "Remove Permission")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeRemoveTargets, "remove-targets", "", false, "Remove Targets")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeStartReplay, "start-replay", "", false, "Start Replay")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeTagResource, "tag-resource", "", false, "Tag Resource")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeTestEventPattern, "test-event-pattern", "", false, "Test Event Pattern")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeUntagResource, "untag-resource", "", false, "Untag Resource")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeUpdateApiDestination, "update-api-destination", "", false, "Update API Destination")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeUpdateArchive, "update-archive", "", false, "Update Archive")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeUpdateConnection, "update-connection", "", false, "Update Connection")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeUpdateEndpoint, "update-endpoint", "", false, "Update Endpoint")
	_eventbridgeCmd.Flags().BoolVarP(&_eventbridgeUpdateEventBus, "update-event-bus", "", false, "Update Event Bus")

}
