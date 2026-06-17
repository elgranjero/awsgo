package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/notifications"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// notificationsCmd represents the notifications command
var _notificationsCmd = &cobra.Command{
	Use:   "notifications",
	Short: "AWS notifications CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := notifications.NewFromConfig(cfg)
		if _notificationsAssociateChannel {
			notifications_AssociateChannel(cfg, client)
			return
		}
		if _notificationsAssociateManagedNotificationAccountContact {
			notifications_AssociateManagedNotificationAccountContact(cfg, client)
			return
		}
		if _notificationsAssociateManagedNotificationAdditionalChannel {
			notifications_AssociateManagedNotificationAdditionalChannel(cfg, client)
			return
		}
		if _notificationsAssociateOrganizationalUnit {
			notifications_AssociateOrganizationalUnit(cfg, client)
			return
		}
		if _notificationsCreateEventRule {
			notifications_CreateEventRule(cfg, client)
			return
		}
		if _notificationsCreateNotificationConfiguration {
			notifications_CreateNotificationConfiguration(cfg, client)
			return
		}
		if _notificationsDeleteEventRule {
			notifications_DeleteEventRule(cfg, client)
			return
		}
		if _notificationsDeleteNotificationConfiguration {
			notifications_DeleteNotificationConfiguration(cfg, client)
			return
		}
		if _notificationsDeregisterNotificationHub {
			notifications_DeregisterNotificationHub(cfg, client)
			return
		}
		if _notificationsDisableNotificationsAccessForOrganization {
			notifications_DisableNotificationsAccessForOrganization(cfg, client)
			return
		}
		if _notificationsDisassociateChannel {
			notifications_DisassociateChannel(cfg, client)
			return
		}
		if _notificationsDisassociateManagedNotificationAccountContact {
			notifications_DisassociateManagedNotificationAccountContact(cfg, client)
			return
		}
		if _notificationsDisassociateManagedNotificationAdditionalChannel {
			notifications_DisassociateManagedNotificationAdditionalChannel(cfg, client)
			return
		}
		if _notificationsDisassociateOrganizationalUnit {
			notifications_DisassociateOrganizationalUnit(cfg, client)
			return
		}
		if _notificationsEnableNotificationsAccessForOrganization {
			notifications_EnableNotificationsAccessForOrganization(cfg, client)
			return
		}
		if _notificationsGetEventRule {
			notifications_GetEventRule(cfg, client)
			return
		}
		if _notificationsGetManagedNotificationChildEvent {
			notifications_GetManagedNotificationChildEvent(cfg, client)
			return
		}
		if _notificationsGetManagedNotificationConfiguration {
			notifications_GetManagedNotificationConfiguration(cfg, client)
			return
		}
		if _notificationsGetManagedNotificationEvent {
			notifications_GetManagedNotificationEvent(cfg, client)
			return
		}
		if _notificationsGetNotificationConfiguration {
			notifications_GetNotificationConfiguration(cfg, client)
			return
		}
		if _notificationsGetNotificationEvent {
			notifications_GetNotificationEvent(cfg, client)
			return
		}
		if _notificationsGetNotificationsAccessForOrganization {
			notifications_GetNotificationsAccessForOrganization(cfg, client)
			return
		}
		if _notificationsListChannels {
			notifications_ListChannels(cfg, client)
			return
		}
		if _notificationsListEventRules {
			notifications_ListEventRules(cfg, client)
			return
		}
		if _notificationsListManagedNotificationChannelAssociations {
			notifications_ListManagedNotificationChannelAssociations(cfg, client)
			return
		}
		if _notificationsListManagedNotificationChildEvents {
			notifications_ListManagedNotificationChildEvents(cfg, client)
			return
		}
		if _notificationsListManagedNotificationConfigurations {
			notifications_ListManagedNotificationConfigurations(cfg, client)
			return
		}
		if _notificationsListManagedNotificationEvents {
			notifications_ListManagedNotificationEvents(cfg, client)
			return
		}
		if _notificationsListMemberAccounts {
			notifications_ListMemberAccounts(cfg, client)
			return
		}
		if _notificationsListNotificationConfigurations {
			notifications_ListNotificationConfigurations(cfg, client)
			return
		}
		if _notificationsListNotificationEvents {
			notifications_ListNotificationEvents(cfg, client)
			return
		}
		if _notificationsListNotificationHubs {
			notifications_ListNotificationHubs(cfg, client)
			return
		}
		if _notificationsListOrganizationalUnits {
			notifications_ListOrganizationalUnits(cfg, client)
			return
		}
		if _notificationsListTagsForResource {
			notifications_ListTagsForResource(cfg, client)
			return
		}
		if _notificationsRegisterNotificationHub {
			notifications_RegisterNotificationHub(cfg, client)
			return
		}
		if _notificationsTagResource {
			notifications_TagResource(cfg, client)
			return
		}
		if _notificationsUntagResource {
			notifications_UntagResource(cfg, client)
			return
		}
		if _notificationsUpdateEventRule {
			notifications_UpdateEventRule(cfg, client)
			return
		}
		if _notificationsUpdateNotificationConfiguration {
			notifications_UpdateNotificationConfiguration(cfg, client)
			return
		}

	},
}

var (
	_notificationsAssociateChannel                                 bool
	_notificationsAssociateManagedNotificationAccountContact       bool
	_notificationsAssociateManagedNotificationAdditionalChannel    bool
	_notificationsAssociateOrganizationalUnit                      bool
	_notificationsCreateEventRule                                  bool
	_notificationsCreateNotificationConfiguration                  bool
	_notificationsDeleteEventRule                                  bool
	_notificationsDeleteNotificationConfiguration                  bool
	_notificationsDeregisterNotificationHub                        bool
	_notificationsDisableNotificationsAccessForOrganization        bool
	_notificationsDisassociateChannel                              bool
	_notificationsDisassociateManagedNotificationAccountContact    bool
	_notificationsDisassociateManagedNotificationAdditionalChannel bool
	_notificationsDisassociateOrganizationalUnit                   bool
	_notificationsEnableNotificationsAccessForOrganization         bool
	_notificationsGetEventRule                                     bool
	_notificationsGetManagedNotificationChildEvent                 bool
	_notificationsGetManagedNotificationConfiguration              bool
	_notificationsGetManagedNotificationEvent                      bool
	_notificationsGetNotificationConfiguration                     bool
	_notificationsGetNotificationEvent                             bool
	_notificationsGetNotificationsAccessForOrganization            bool
	_notificationsListChannels                                     bool
	_notificationsListEventRules                                   bool
	_notificationsListManagedNotificationChannelAssociations       bool
	_notificationsListManagedNotificationChildEvents               bool
	_notificationsListManagedNotificationConfigurations            bool
	_notificationsListManagedNotificationEvents                    bool
	_notificationsListMemberAccounts                               bool
	_notificationsListNotificationConfigurations                   bool
	_notificationsListNotificationEvents                           bool
	_notificationsListNotificationHubs                             bool
	_notificationsListOrganizationalUnits                          bool
	_notificationsListTagsForResource                              bool
	_notificationsRegisterNotificationHub                          bool
	_notificationsTagResource                                      bool
	_notificationsUntagResource                                    bool
	_notificationsUpdateEventRule                                  bool
	_notificationsUpdateNotificationConfiguration                  bool

	_notificationsAggregateManagedNotificationEventArn string
	_notificationsAggregateNotificationEventArn        string
	_notificationsAggregationDuration                  string
	_notificationsArn                                  string
	_notificationsChannelArn                           string
	_notificationsChannelIdentifier                    string
	_notificationsContactIdentifier                    string
	_notificationsDescription                          string
	_notificationsEndTime                              string
	_notificationsEventPattern                         string
	_notificationsEventRuleSource                      string
	_notificationsEventType                            string
	_notificationsIncludeChildEvents                   string
	_notificationsLocale                               string
	_notificationsManagedNotificationConfigurationArn  string
	_notificationsMaxResults                           string
	_notificationsMemberAccount                        string
	_notificationsName                                 string
	_notificationsNextToken                            string
	_notificationsNotificationConfigurationArn         string
	_notificationsNotificationHubRegion                string
	_notificationsOrganizationalUnitId                 string
	_notificationsRegions                              []string
	_notificationsRelatedAccount                       string
	_notificationsSource                               string
	_notificationsStartTime                            string
	_notificationsStatus                               string
	_notificationsSubtype                              string
	_notificationsTagKeys                              []string
	_notificationsTags                                 string
)

// Associates a delivery [Channel] with a particular NotificationConfiguration . Supported
// Channels include Amazon Q Developer in chat applications, the Console Mobile
// Application, and emails (notifications-contacts).
//
// [Channel]: https://docs.aws.amazon.com/notifications/latest/userguide/managing-delivery-channels.html
func notifications_AssociateChannel(cfg aws.Config, client *notifications.Client) {
	input := &notifications.AssociateChannelInput{
		// Arn: *string, // Required
		// NotificationConfigurationArn: *string, // Required
	}

	if len(_notificationsArn) > 0 {
		input.Arn = aws.String(_notificationsArn)
	}
	if len(_notificationsNotificationConfigurationArn) > 0 {
		input.NotificationConfigurationArn = aws.String(_notificationsNotificationConfigurationArn)
	}

	if resp, err := client.AssociateChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates an Account Contact with a particular ManagedNotificationConfiguration
// .
func notifications_AssociateManagedNotificationAccountContact(cfg aws.Config, client *notifications.Client) {
	input := &notifications.AssociateManagedNotificationAccountContactInput{
		// ContactIdentifier: types.AccountContactType, // Required
		// ManagedNotificationConfigurationArn: *string, // Required
	}

	if len(_notificationsContactIdentifier) > 0 {
		if err := assignInputField(input, "ContactIdentifier", _notificationsContactIdentifier); err != nil {
			log.Errorf("invalid --contact-identifier: %s", err.Error())
			return
		}
	}
	if len(_notificationsManagedNotificationConfigurationArn) > 0 {
		input.ManagedNotificationConfigurationArn = aws.String(_notificationsManagedNotificationConfigurationArn)
	}

	if resp, err := client.AssociateManagedNotificationAccountContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates an additional Channel with a particular
// ManagedNotificationConfiguration .
//
// Supported Channels include Amazon Q Developer in chat applications, the Console
// Mobile Application, and emails (notifications-contacts).
func notifications_AssociateManagedNotificationAdditionalChannel(cfg aws.Config, client *notifications.Client) {
	input := &notifications.AssociateManagedNotificationAdditionalChannelInput{
		// ChannelArn: *string, // Required
		// ManagedNotificationConfigurationArn: *string, // Required
	}

	if len(_notificationsChannelArn) > 0 {
		input.ChannelArn = aws.String(_notificationsChannelArn)
	}
	if len(_notificationsManagedNotificationConfigurationArn) > 0 {
		input.ManagedNotificationConfigurationArn = aws.String(_notificationsManagedNotificationConfigurationArn)
	}

	if resp, err := client.AssociateManagedNotificationAdditionalChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates an organizational unit with a notification configuration.
func notifications_AssociateOrganizationalUnit(cfg aws.Config, client *notifications.Client) {
	input := &notifications.AssociateOrganizationalUnitInput{
		// NotificationConfigurationArn: *string, // Required
		// OrganizationalUnitId: *string, // Required
	}

	if len(_notificationsNotificationConfigurationArn) > 0 {
		input.NotificationConfigurationArn = aws.String(_notificationsNotificationConfigurationArn)
	}
	if len(_notificationsOrganizationalUnitId) > 0 {
		input.OrganizationalUnitId = aws.String(_notificationsOrganizationalUnitId)
	}

	if resp, err := client.AssociateOrganizationalUnit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an [EventRule]EventRule that is associated with a specified
// NotificationConfiguration .
//
// [EventRule]: https://docs.aws.amazon.com/notifications/latest/userguide/glossary.html
func notifications_CreateEventRule(cfg aws.Config, client *notifications.Client) {
	input := &notifications.CreateEventRuleInput{
		// EventType: *string, // Required
		// NotificationConfigurationArn: *string, // Required
		// Regions: []string, // Required
		// Source: *string, // Required
	}

	if len(_notificationsEventType) > 0 {
		input.EventType = aws.String(_notificationsEventType)
	}
	if len(_notificationsNotificationConfigurationArn) > 0 {
		input.NotificationConfigurationArn = aws.String(_notificationsNotificationConfigurationArn)
	}
	if len(_notificationsRegions) > 0 {
		input.Regions = append([]string(nil), _notificationsRegions...)
	}
	if len(_notificationsSource) > 0 {
		input.Source = aws.String(_notificationsSource)
	}
	if len(_notificationsEventPattern) > 0 {
		input.EventPattern = aws.String(_notificationsEventPattern)
	}

	if resp, err := client.CreateEventRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new NotificationConfiguration .
func notifications_CreateNotificationConfiguration(cfg aws.Config, client *notifications.Client) {
	input := &notifications.CreateNotificationConfigurationInput{
		// Description: *string, // Required
		// Name: *string, // Required
	}

	if len(_notificationsDescription) > 0 {
		input.Description = aws.String(_notificationsDescription)
	}
	if len(_notificationsName) > 0 {
		input.Name = aws.String(_notificationsName)
	}
	if len(_notificationsAggregationDuration) > 0 {
		if err := assignInputField(input, "AggregationDuration", _notificationsAggregationDuration); err != nil {
			log.Errorf("invalid --aggregation-duration: %s", err.Error())
			return
		}
	}
	if len(_notificationsTags) > 0 {
		if err := assignInputField(input, "Tags", _notificationsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateNotificationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an EventRule .
func notifications_DeleteEventRule(cfg aws.Config, client *notifications.Client) {
	input := &notifications.DeleteEventRuleInput{
		// Arn: *string, // Required
	}

	if len(_notificationsArn) > 0 {
		input.Arn = aws.String(_notificationsArn)
	}

	if resp, err := client.DeleteEventRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a NotificationConfiguration .
func notifications_DeleteNotificationConfiguration(cfg aws.Config, client *notifications.Client) {
	input := &notifications.DeleteNotificationConfigurationInput{
		// Arn: *string, // Required
	}

	if len(_notificationsArn) > 0 {
		input.Arn = aws.String(_notificationsArn)
	}

	if resp, err := client.DeleteNotificationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deregisters a NotificationConfiguration in the specified Region.
// You can't deregister the last NotificationHub in the account. NotificationEvents
// stored in the deregistered NotificationConfiguration are no longer be visible.
// Recreating a new NotificationConfiguration in the same Region restores access
// to those NotificationEvents .
func notifications_DeregisterNotificationHub(cfg aws.Config, client *notifications.Client) {
	input := &notifications.DeregisterNotificationHubInput{
		// NotificationHubRegion: *string, // Required
	}

	if len(_notificationsNotificationHubRegion) > 0 {
		input.NotificationHubRegion = aws.String(_notificationsNotificationHubRegion)
	}

	if resp, err := client.DeregisterNotificationHub(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables service trust between User Notifications and Amazon Web Services
// Organizations.
func notifications_DisableNotificationsAccessForOrganization(cfg aws.Config, client *notifications.Client) {
	input := &notifications.DisableNotificationsAccessForOrganizationInput{}

	if resp, err := client.DisableNotificationsAccessForOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a Channel from a specified NotificationConfiguration . Supported
// Channels include Amazon Q Developer in chat applications, the Console Mobile
// Application, and emails (notifications-contacts).
func notifications_DisassociateChannel(cfg aws.Config, client *notifications.Client) {
	input := &notifications.DisassociateChannelInput{
		// Arn: *string, // Required
		// NotificationConfigurationArn: *string, // Required
	}

	if len(_notificationsArn) > 0 {
		input.Arn = aws.String(_notificationsArn)
	}
	if len(_notificationsNotificationConfigurationArn) > 0 {
		input.NotificationConfigurationArn = aws.String(_notificationsNotificationConfigurationArn)
	}

	if resp, err := client.DisassociateChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates an Account Contact with a particular
// ManagedNotificationConfiguration .
func notifications_DisassociateManagedNotificationAccountContact(cfg aws.Config, client *notifications.Client) {
	input := &notifications.DisassociateManagedNotificationAccountContactInput{
		// ContactIdentifier: types.AccountContactType, // Required
		// ManagedNotificationConfigurationArn: *string, // Required
	}

	if len(_notificationsContactIdentifier) > 0 {
		if err := assignInputField(input, "ContactIdentifier", _notificationsContactIdentifier); err != nil {
			log.Errorf("invalid --contact-identifier: %s", err.Error())
			return
		}
	}
	if len(_notificationsManagedNotificationConfigurationArn) > 0 {
		input.ManagedNotificationConfigurationArn = aws.String(_notificationsManagedNotificationConfigurationArn)
	}

	if resp, err := client.DisassociateManagedNotificationAccountContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates an additional Channel from a particular
// ManagedNotificationConfiguration .
//
// Supported Channels include Amazon Q Developer in chat applications, the Console
// Mobile Application, and emails (notifications-contacts).
func notifications_DisassociateManagedNotificationAdditionalChannel(cfg aws.Config, client *notifications.Client) {
	input := &notifications.DisassociateManagedNotificationAdditionalChannelInput{
		// ChannelArn: *string, // Required
		// ManagedNotificationConfigurationArn: *string, // Required
	}

	if len(_notificationsChannelArn) > 0 {
		input.ChannelArn = aws.String(_notificationsChannelArn)
	}
	if len(_notificationsManagedNotificationConfigurationArn) > 0 {
		input.ManagedNotificationConfigurationArn = aws.String(_notificationsManagedNotificationConfigurationArn)
	}

	if resp, err := client.DisassociateManagedNotificationAdditionalChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the association between an organizational unit and a notification
// configuration.
func notifications_DisassociateOrganizationalUnit(cfg aws.Config, client *notifications.Client) {
	input := &notifications.DisassociateOrganizationalUnitInput{
		// NotificationConfigurationArn: *string, // Required
		// OrganizationalUnitId: *string, // Required
	}

	if len(_notificationsNotificationConfigurationArn) > 0 {
		input.NotificationConfigurationArn = aws.String(_notificationsNotificationConfigurationArn)
	}
	if len(_notificationsOrganizationalUnitId) > 0 {
		input.OrganizationalUnitId = aws.String(_notificationsOrganizationalUnitId)
	}

	if resp, err := client.DisassociateOrganizationalUnit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables service trust between User Notifications and Amazon Web Services
// Organizations.
func notifications_EnableNotificationsAccessForOrganization(cfg aws.Config, client *notifications.Client) {
	input := &notifications.EnableNotificationsAccessForOrganizationInput{}

	if resp, err := client.EnableNotificationsAccessForOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a specified EventRule .
func notifications_GetEventRule(cfg aws.Config, client *notifications.Client) {
	input := &notifications.GetEventRuleInput{
		// Arn: *string, // Required
	}

	if len(_notificationsArn) > 0 {
		input.Arn = aws.String(_notificationsArn)
	}

	if resp, err := client.GetEventRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the child event of a specific given ManagedNotificationEvent .
func notifications_GetManagedNotificationChildEvent(cfg aws.Config, client *notifications.Client) {
	input := &notifications.GetManagedNotificationChildEventInput{
		// Arn: *string, // Required
	}

	if len(_notificationsArn) > 0 {
		input.Arn = aws.String(_notificationsArn)
	}
	if len(_notificationsLocale) > 0 {
		if err := assignInputField(input, "Locale", _notificationsLocale); err != nil {
			log.Errorf("invalid --locale: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetManagedNotificationChildEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a specified ManagedNotificationConfiguration .
func notifications_GetManagedNotificationConfiguration(cfg aws.Config, client *notifications.Client) {
	input := &notifications.GetManagedNotificationConfigurationInput{
		// Arn: *string, // Required
	}

	if len(_notificationsArn) > 0 {
		input.Arn = aws.String(_notificationsArn)
	}

	if resp, err := client.GetManagedNotificationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a specified ManagedNotificationEvent .
func notifications_GetManagedNotificationEvent(cfg aws.Config, client *notifications.Client) {
	input := &notifications.GetManagedNotificationEventInput{
		// Arn: *string, // Required
	}

	if len(_notificationsArn) > 0 {
		input.Arn = aws.String(_notificationsArn)
	}
	if len(_notificationsLocale) > 0 {
		if err := assignInputField(input, "Locale", _notificationsLocale); err != nil {
			log.Errorf("invalid --locale: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetManagedNotificationEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a specified NotificationConfiguration .
func notifications_GetNotificationConfiguration(cfg aws.Config, client *notifications.Client) {
	input := &notifications.GetNotificationConfigurationInput{
		// Arn: *string, // Required
	}

	if len(_notificationsArn) > 0 {
		input.Arn = aws.String(_notificationsArn)
	}

	if resp, err := client.GetNotificationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a specified NotificationEvent .
// User Notifications stores notifications in the individual Regions you register
// as notification hubs and the Region of the source event rule.
// GetNotificationEvent only returns notifications stored in the same Region in
// which the action is called. User Notifications doesn't backfill notifications to
// new Regions selected as notification hubs. For this reason, we recommend that
// you make calls in your oldest registered notification hub. For more information,
// see [Notification hubs]in the Amazon Web Services User Notifications User Guide.
//
// [Notification hubs]: https://docs.aws.amazon.com/notifications/latest/userguide/notification-hubs.html
func notifications_GetNotificationEvent(cfg aws.Config, client *notifications.Client) {
	input := &notifications.GetNotificationEventInput{
		// Arn: *string, // Required
	}

	if len(_notificationsArn) > 0 {
		input.Arn = aws.String(_notificationsArn)
	}
	if len(_notificationsLocale) > 0 {
		if err := assignInputField(input, "Locale", _notificationsLocale); err != nil {
			log.Errorf("invalid --locale: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetNotificationEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the AccessStatus of Service Trust Enablement for User Notifications and
// Amazon Web Services Organizations.
func notifications_GetNotificationsAccessForOrganization(cfg aws.Config, client *notifications.Client) {
	input := &notifications.GetNotificationsAccessForOrganizationInput{}

	if resp, err := client.GetNotificationsAccessForOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of Channels for a NotificationConfiguration .
func notifications_ListChannels(cfg aws.Config, client *notifications.Client) {
	input := &notifications.ListChannelsInput{
		// NotificationConfigurationArn: *string, // Required
	}

	if len(_notificationsNotificationConfigurationArn) > 0 {
		input.NotificationConfigurationArn = aws.String(_notificationsNotificationConfigurationArn)
	}
	if len(_notificationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _notificationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_notificationsNextToken) > 0 {
		input.NextToken = aws.String(_notificationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListChannels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*notifications.ListChannelsOutput
	p := notifications.NewListChannelsPaginator(client, input)
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

// Returns a list of EventRules according to specified filters, in reverse
// chronological order (newest first).
func notifications_ListEventRules(cfg aws.Config, client *notifications.Client) {
	input := &notifications.ListEventRulesInput{
		// NotificationConfigurationArn: *string, // Required
	}

	if len(_notificationsNotificationConfigurationArn) > 0 {
		input.NotificationConfigurationArn = aws.String(_notificationsNotificationConfigurationArn)
	}
	if len(_notificationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _notificationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_notificationsNextToken) > 0 {
		input.NextToken = aws.String(_notificationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEventRules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*notifications.ListEventRulesOutput
	p := notifications.NewListEventRulesPaginator(client, input)
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

// Returns a list of Account contacts and Channels associated with a
// ManagedNotificationConfiguration , in paginated format.
func notifications_ListManagedNotificationChannelAssociations(cfg aws.Config, client *notifications.Client) {
	input := &notifications.ListManagedNotificationChannelAssociationsInput{
		// ManagedNotificationConfigurationArn: *string, // Required
	}

	if len(_notificationsManagedNotificationConfigurationArn) > 0 {
		input.ManagedNotificationConfigurationArn = aws.String(_notificationsManagedNotificationConfigurationArn)
	}
	if len(_notificationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _notificationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_notificationsNextToken) > 0 {
		input.NextToken = aws.String(_notificationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListManagedNotificationChannelAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*notifications.ListManagedNotificationChannelAssociationsOutput
	p := notifications.NewListManagedNotificationChannelAssociationsPaginator(client, input)
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

// Returns a list of ManagedNotificationChildEvents for a specified aggregate
// ManagedNotificationEvent , ordered by creation time in reverse chronological
// order (newest first).
func notifications_ListManagedNotificationChildEvents(cfg aws.Config, client *notifications.Client) {
	input := &notifications.ListManagedNotificationChildEventsInput{
		// AggregateManagedNotificationEventArn: *string, // Required
	}

	if len(_notificationsAggregateManagedNotificationEventArn) > 0 {
		input.AggregateManagedNotificationEventArn = aws.String(_notificationsAggregateManagedNotificationEventArn)
	}
	if len(_notificationsEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _notificationsEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_notificationsLocale) > 0 {
		if err := assignInputField(input, "Locale", _notificationsLocale); err != nil {
			log.Errorf("invalid --locale: %s", err.Error())
			return
		}
	}
	if len(_notificationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _notificationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_notificationsNextToken) > 0 {
		input.NextToken = aws.String(_notificationsNextToken)
	}
	if len(_notificationsOrganizationalUnitId) > 0 {
		input.OrganizationalUnitId = aws.String(_notificationsOrganizationalUnitId)
	}
	if len(_notificationsRelatedAccount) > 0 {
		input.RelatedAccount = aws.String(_notificationsRelatedAccount)
	}
	if len(_notificationsStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _notificationsStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListManagedNotificationChildEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*notifications.ListManagedNotificationChildEventsOutput
	p := notifications.NewListManagedNotificationChildEventsPaginator(client, input)
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

// Returns a list of Managed Notification Configurations according to specified
// filters, ordered by creation time in reverse chronological order (newest first).
func notifications_ListManagedNotificationConfigurations(cfg aws.Config, client *notifications.Client) {
	input := &notifications.ListManagedNotificationConfigurationsInput{}

	if len(_notificationsChannelIdentifier) > 0 {
		input.ChannelIdentifier = aws.String(_notificationsChannelIdentifier)
	}
	if len(_notificationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _notificationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_notificationsNextToken) > 0 {
		input.NextToken = aws.String(_notificationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListManagedNotificationConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*notifications.ListManagedNotificationConfigurationsOutput
	p := notifications.NewListManagedNotificationConfigurationsPaginator(client, input)
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

// Returns a list of Managed Notification Events according to specified filters,
// ordered by creation time in reverse chronological order (newest first).
func notifications_ListManagedNotificationEvents(cfg aws.Config, client *notifications.Client) {
	input := &notifications.ListManagedNotificationEventsInput{}

	if len(_notificationsEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _notificationsEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_notificationsLocale) > 0 {
		if err := assignInputField(input, "Locale", _notificationsLocale); err != nil {
			log.Errorf("invalid --locale: %s", err.Error())
			return
		}
	}
	if len(_notificationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _notificationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_notificationsNextToken) > 0 {
		input.NextToken = aws.String(_notificationsNextToken)
	}
	if len(_notificationsOrganizationalUnitId) > 0 {
		input.OrganizationalUnitId = aws.String(_notificationsOrganizationalUnitId)
	}
	if len(_notificationsRelatedAccount) > 0 {
		input.RelatedAccount = aws.String(_notificationsRelatedAccount)
	}
	if len(_notificationsSource) > 0 {
		input.Source = aws.String(_notificationsSource)
	}
	if len(_notificationsStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _notificationsStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListManagedNotificationEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*notifications.ListManagedNotificationEventsOutput
	p := notifications.NewListManagedNotificationEventsPaginator(client, input)
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

// Returns a list of member accounts associated with a notification configuration.
func notifications_ListMemberAccounts(cfg aws.Config, client *notifications.Client) {
	input := &notifications.ListMemberAccountsInput{
		// NotificationConfigurationArn: *string, // Required
	}

	if len(_notificationsNotificationConfigurationArn) > 0 {
		input.NotificationConfigurationArn = aws.String(_notificationsNotificationConfigurationArn)
	}
	if len(_notificationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _notificationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_notificationsMemberAccount) > 0 {
		input.MemberAccount = aws.String(_notificationsMemberAccount)
	}
	if len(_notificationsNextToken) > 0 {
		input.NextToken = aws.String(_notificationsNextToken)
	}
	if len(_notificationsOrganizationalUnitId) > 0 {
		input.OrganizationalUnitId = aws.String(_notificationsOrganizationalUnitId)
	}
	if len(_notificationsStatus) > 0 {
		if err := assignInputField(input, "Status", _notificationsStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListMemberAccounts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*notifications.ListMemberAccountsOutput
	p := notifications.NewListMemberAccountsPaginator(client, input)
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

// Returns a list of abbreviated NotificationConfigurations according to specified
// filters, in reverse chronological order (newest first).
func notifications_ListNotificationConfigurations(cfg aws.Config, client *notifications.Client) {
	input := &notifications.ListNotificationConfigurationsInput{}

	if len(_notificationsChannelArn) > 0 {
		input.ChannelArn = aws.String(_notificationsChannelArn)
	}
	if len(_notificationsEventRuleSource) > 0 {
		input.EventRuleSource = aws.String(_notificationsEventRuleSource)
	}
	if len(_notificationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _notificationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_notificationsNextToken) > 0 {
		input.NextToken = aws.String(_notificationsNextToken)
	}
	if len(_notificationsStatus) > 0 {
		if err := assignInputField(input, "Status", _notificationsStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_notificationsSubtype) > 0 {
		if err := assignInputField(input, "Subtype", _notificationsSubtype); err != nil {
			log.Errorf("invalid --subtype: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListNotificationConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*notifications.ListNotificationConfigurationsOutput
	p := notifications.NewListNotificationConfigurationsPaginator(client, input)
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

// Returns a list of NotificationEvents according to specified filters, in reverse
// chronological order (newest first).
//
// User Notifications stores notifications in the individual Regions you register
// as notification hubs and the Region of the source event rule.
// ListNotificationEvents only returns notifications stored in the same Region in
// which the action is called. User Notifications doesn't backfill notifications to
// new Regions selected as notification hubs. For this reason, we recommend that
// you make calls in your oldest registered notification hub. For more information,
// see [Notification hubs]in the Amazon Web Services User Notifications User Guide.
//
// [Notification hubs]: https://docs.aws.amazon.com/notifications/latest/userguide/notification-hubs.html
func notifications_ListNotificationEvents(cfg aws.Config, client *notifications.Client) {
	input := &notifications.ListNotificationEventsInput{}

	if len(_notificationsAggregateNotificationEventArn) > 0 {
		input.AggregateNotificationEventArn = aws.String(_notificationsAggregateNotificationEventArn)
	}
	if len(_notificationsEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _notificationsEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_notificationsIncludeChildEvents) > 0 {
		if err := assignInputField(input, "IncludeChildEvents", _notificationsIncludeChildEvents); err != nil {
			log.Errorf("invalid --include-child-events: %s", err.Error())
			return
		}
	}
	if len(_notificationsLocale) > 0 {
		if err := assignInputField(input, "Locale", _notificationsLocale); err != nil {
			log.Errorf("invalid --locale: %s", err.Error())
			return
		}
	}
	if len(_notificationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _notificationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_notificationsNextToken) > 0 {
		input.NextToken = aws.String(_notificationsNextToken)
	}
	if len(_notificationsOrganizationalUnitId) > 0 {
		input.OrganizationalUnitId = aws.String(_notificationsOrganizationalUnitId)
	}
	if len(_notificationsSource) > 0 {
		input.Source = aws.String(_notificationsSource)
	}
	if len(_notificationsStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _notificationsStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListNotificationEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*notifications.ListNotificationEventsOutput
	p := notifications.NewListNotificationEventsPaginator(client, input)
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

// Returns a list of NotificationHubs .
func notifications_ListNotificationHubs(cfg aws.Config, client *notifications.Client) {
	input := &notifications.ListNotificationHubsInput{}

	if len(_notificationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _notificationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_notificationsNextToken) > 0 {
		input.NextToken = aws.String(_notificationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListNotificationHubs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*notifications.ListNotificationHubsOutput
	p := notifications.NewListNotificationHubsPaginator(client, input)
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

// Returns a list of organizational units associated with a notification
// configuration.
func notifications_ListOrganizationalUnits(cfg aws.Config, client *notifications.Client) {
	input := &notifications.ListOrganizationalUnitsInput{
		// NotificationConfigurationArn: *string, // Required
	}

	if len(_notificationsNotificationConfigurationArn) > 0 {
		input.NotificationConfigurationArn = aws.String(_notificationsNotificationConfigurationArn)
	}
	if len(_notificationsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _notificationsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_notificationsNextToken) > 0 {
		input.NextToken = aws.String(_notificationsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListOrganizationalUnits(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*notifications.ListOrganizationalUnitsOutput
	p := notifications.NewListOrganizationalUnitsPaginator(client, input)
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

// Returns a list of tags for a specified Amazon Resource Name (ARN).
// For more information, see [Tagging your Amazon Web Services resources] in the Tagging Amazon Web Services Resources User
// Guide.
//
// This is only supported for NotificationConfigurations .
//
// [Tagging your Amazon Web Services resources]: https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html
func notifications_ListTagsForResource(cfg aws.Config, client *notifications.Client) {
	input := &notifications.ListTagsForResourceInput{
		// Arn: *string, // Required
	}

	if len(_notificationsArn) > 0 {
		input.Arn = aws.String(_notificationsArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers a NotificationConfiguration in the specified Region.
// There is a maximum of one NotificationConfiguration per Region. You can have a
// maximum of 3 NotificationHub resources at a time.
func notifications_RegisterNotificationHub(cfg aws.Config, client *notifications.Client) {
	input := &notifications.RegisterNotificationHubInput{
		// NotificationHubRegion: *string, // Required
	}

	if len(_notificationsNotificationHubRegion) > 0 {
		input.NotificationHubRegion = aws.String(_notificationsNotificationHubRegion)
	}

	if resp, err := client.RegisterNotificationHub(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tags the resource with a tag key and value.
// For more information, see [Tagging your Amazon Web Services resources] in the Tagging Amazon Web Services Resources User
// Guide.
//
// This is only supported for NotificationConfigurations .
//
// [Tagging your Amazon Web Services resources]: https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html
func notifications_TagResource(cfg aws.Config, client *notifications.Client) {
	input := &notifications.TagResourceInput{
		// Arn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_notificationsArn) > 0 {
		input.Arn = aws.String(_notificationsArn)
	}
	if len(_notificationsTags) > 0 {
		if err := assignInputField(input, "Tags", _notificationsTags); err != nil {
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

// Untags a resource with a specified Amazon Resource Name (ARN).
// For more information, see [Tagging your Amazon Web Services resources] in the Tagging Amazon Web Services Resources User
// Guide.
//
// [Tagging your Amazon Web Services resources]: https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html
func notifications_UntagResource(cfg aws.Config, client *notifications.Client) {
	input := &notifications.UntagResourceInput{
		// Arn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_notificationsArn) > 0 {
		input.Arn = aws.String(_notificationsArn)
	}
	if len(_notificationsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _notificationsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing EventRule .
func notifications_UpdateEventRule(cfg aws.Config, client *notifications.Client) {
	input := &notifications.UpdateEventRuleInput{
		// Arn: *string, // Required
	}

	if len(_notificationsArn) > 0 {
		input.Arn = aws.String(_notificationsArn)
	}
	if len(_notificationsEventPattern) > 0 {
		input.EventPattern = aws.String(_notificationsEventPattern)
	}
	if len(_notificationsRegions) > 0 {
		input.Regions = append([]string(nil), _notificationsRegions...)
	}

	if resp, err := client.UpdateEventRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a NotificationConfiguration .
func notifications_UpdateNotificationConfiguration(cfg aws.Config, client *notifications.Client) {
	input := &notifications.UpdateNotificationConfigurationInput{
		// Arn: *string, // Required
	}

	if len(_notificationsArn) > 0 {
		input.Arn = aws.String(_notificationsArn)
	}
	if len(_notificationsAggregationDuration) > 0 {
		if err := assignInputField(input, "AggregationDuration", _notificationsAggregationDuration); err != nil {
			log.Errorf("invalid --aggregation-duration: %s", err.Error())
			return
		}
	}
	if len(_notificationsDescription) > 0 {
		input.Description = aws.String(_notificationsDescription)
	}
	if len(_notificationsName) > 0 {
		input.Name = aws.String(_notificationsName)
	}

	if resp, err := client.UpdateNotificationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_notificationsCmd)
	_notificationsCmd.Flags().SortFlags = false

	_notificationsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_notificationsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_notificationsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_notificationsCmd.Flags().StringVarP(&_notificationsAggregateManagedNotificationEventArn, "aggregate-managed-notification-event-arn", "", "", "Aggregate Managed Notification Event ARN")
	_notificationsCmd.Flags().StringVarP(&_notificationsAggregateNotificationEventArn, "aggregate-notification-event-arn", "", "", "Aggregate Notification Event ARN")
	_notificationsCmd.Flags().StringVarP(&_notificationsAggregationDuration, "aggregation-duration", "", "", "Aggregation Duration")
	_notificationsCmd.Flags().StringVarP(&_notificationsArn, "arn", "", "", "ARN")
	_notificationsCmd.Flags().StringVarP(&_notificationsChannelArn, "channel-arn", "", "", "Channel ARN")
	_notificationsCmd.Flags().StringVarP(&_notificationsChannelIdentifier, "channel-identifier", "", "", "Channel Identifier")
	_notificationsCmd.Flags().StringVarP(&_notificationsContactIdentifier, "contact-identifier", "", "", "Contact Identifier")
	_notificationsCmd.Flags().StringVarP(&_notificationsDescription, "description", "", "", "Description")
	_notificationsCmd.Flags().StringVarP(&_notificationsEndTime, "end-time", "", "", "End Time")
	_notificationsCmd.Flags().StringVarP(&_notificationsEventPattern, "event-pattern", "", "", "Event Pattern")
	_notificationsCmd.Flags().StringVarP(&_notificationsEventRuleSource, "event-rule-source", "", "", "Event Rule Source")
	_notificationsCmd.Flags().StringVarP(&_notificationsEventType, "event-type", "", "", "Event Type")
	_notificationsCmd.Flags().StringVarP(&_notificationsIncludeChildEvents, "include-child-events", "", "", "Include Child Events")
	_notificationsCmd.Flags().StringVarP(&_notificationsLocale, "locale", "", "", "Locale")
	_notificationsCmd.Flags().StringVarP(&_notificationsManagedNotificationConfigurationArn, "managed-notification-configuration-arn", "", "", "Managed Notification Configuration ARN")
	_notificationsCmd.Flags().StringVarP(&_notificationsMaxResults, "max-results", "", "", "Max Results")
	_notificationsCmd.Flags().StringVarP(&_notificationsMemberAccount, "member-account", "", "", "Member Account")
	_notificationsCmd.Flags().StringVarP(&_notificationsName, "name", "", "", "Name")
	_notificationsCmd.Flags().StringVarP(&_notificationsNextToken, "next-token", "", "", "Next Token")
	_notificationsCmd.Flags().StringVarP(&_notificationsNotificationConfigurationArn, "notification-configuration-arn", "", "", "Notification Configuration ARN")
	_notificationsCmd.Flags().StringVarP(&_notificationsNotificationHubRegion, "notification-hub-region", "", "", "Notification Hub Region")
	_notificationsCmd.Flags().StringVarP(&_notificationsOrganizationalUnitId, "organizational-unit-id", "", "", "Organizational Unit ID")
	_notificationsCmd.Flags().StringSliceVarP(&_notificationsRegions, "regions", "", nil, "Regions")
	_notificationsCmd.Flags().StringVarP(&_notificationsRelatedAccount, "related-account", "", "", "Related Account")
	_notificationsCmd.Flags().StringVarP(&_notificationsSource, "source", "", "", "Source")
	_notificationsCmd.Flags().StringVarP(&_notificationsStartTime, "start-time", "", "", "Start Time")
	_notificationsCmd.Flags().StringVarP(&_notificationsStatus, "status", "", "", "Status")
	_notificationsCmd.Flags().StringVarP(&_notificationsSubtype, "subtype", "", "", "Subtype")
	_notificationsCmd.Flags().StringSliceVarP(&_notificationsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_notificationsCmd.Flags().StringVarP(&_notificationsTags, "tags", "", "", "Tags")

	_notificationsCmd.Flags().BoolVarP(&_notificationsAssociateChannel, "associate-channel", "", false, "Associate Channel")
	_notificationsCmd.Flags().BoolVarP(&_notificationsAssociateManagedNotificationAccountContact, "associate-managed-notification-account-contact", "", false, "Associate Managed Notification Account Contact")
	_notificationsCmd.Flags().BoolVarP(&_notificationsAssociateManagedNotificationAdditionalChannel, "associate-managed-notification-additional-channel", "", false, "Associate Managed Notification Additional Channel")
	_notificationsCmd.Flags().BoolVarP(&_notificationsAssociateOrganizationalUnit, "associate-organizational-unit", "", false, "Associate Organizational Unit")
	_notificationsCmd.Flags().BoolVarP(&_notificationsCreateEventRule, "create-event-rule", "", false, "Create Event Rule")
	_notificationsCmd.Flags().BoolVarP(&_notificationsCreateNotificationConfiguration, "create-notification-configuration", "", false, "Create Notification Configuration")
	_notificationsCmd.Flags().BoolVarP(&_notificationsDeleteEventRule, "delete-event-rule", "", false, "Delete Event Rule")
	_notificationsCmd.Flags().BoolVarP(&_notificationsDeleteNotificationConfiguration, "delete-notification-configuration", "", false, "Delete Notification Configuration")
	_notificationsCmd.Flags().BoolVarP(&_notificationsDeregisterNotificationHub, "deregister-notification-hub", "", false, "Deregister Notification Hub")
	_notificationsCmd.Flags().BoolVarP(&_notificationsDisableNotificationsAccessForOrganization, "disable-notifications-access-for-organization", "", false, "Disable Notifications Access For Organization")
	_notificationsCmd.Flags().BoolVarP(&_notificationsDisassociateChannel, "disassociate-channel", "", false, "Disassociate Channel")
	_notificationsCmd.Flags().BoolVarP(&_notificationsDisassociateManagedNotificationAccountContact, "disassociate-managed-notification-account-contact", "", false, "Disassociate Managed Notification Account Contact")
	_notificationsCmd.Flags().BoolVarP(&_notificationsDisassociateManagedNotificationAdditionalChannel, "disassociate-managed-notification-additional-channel", "", false, "Disassociate Managed Notification Additional Channel")
	_notificationsCmd.Flags().BoolVarP(&_notificationsDisassociateOrganizationalUnit, "disassociate-organizational-unit", "", false, "Disassociate Organizational Unit")
	_notificationsCmd.Flags().BoolVarP(&_notificationsEnableNotificationsAccessForOrganization, "enable-notifications-access-for-organization", "", false, "Enable Notifications Access For Organization")
	_notificationsCmd.Flags().BoolVarP(&_notificationsGetEventRule, "get-event-rule", "", false, "Get Event Rule")
	_notificationsCmd.Flags().BoolVarP(&_notificationsGetManagedNotificationChildEvent, "get-managed-notification-child-event", "", false, "Get Managed Notification Child Event")
	_notificationsCmd.Flags().BoolVarP(&_notificationsGetManagedNotificationConfiguration, "get-managed-notification-configuration", "", false, "Get Managed Notification Configuration")
	_notificationsCmd.Flags().BoolVarP(&_notificationsGetManagedNotificationEvent, "get-managed-notification-event", "", false, "Get Managed Notification Event")
	_notificationsCmd.Flags().BoolVarP(&_notificationsGetNotificationConfiguration, "get-notification-configuration", "", false, "Get Notification Configuration")
	_notificationsCmd.Flags().BoolVarP(&_notificationsGetNotificationEvent, "get-notification-event", "", false, "Get Notification Event")
	_notificationsCmd.Flags().BoolVarP(&_notificationsGetNotificationsAccessForOrganization, "get-notifications-access-for-organization", "", false, "Get Notifications Access For Organization")
	_notificationsCmd.Flags().BoolVarP(&_notificationsListChannels, "list-channels", "", false, "List Channels")
	_notificationsCmd.Flags().BoolVarP(&_notificationsListEventRules, "list-event-rules", "", false, "List Event Rules")
	_notificationsCmd.Flags().BoolVarP(&_notificationsListManagedNotificationChannelAssociations, "list-managed-notification-channel-associations", "", false, "List Managed Notification Channel Associations")
	_notificationsCmd.Flags().BoolVarP(&_notificationsListManagedNotificationChildEvents, "list-managed-notification-child-events", "", false, "List Managed Notification Child Events")
	_notificationsCmd.Flags().BoolVarP(&_notificationsListManagedNotificationConfigurations, "list-managed-notification-configurations", "", false, "List Managed Notification Configurations")
	_notificationsCmd.Flags().BoolVarP(&_notificationsListManagedNotificationEvents, "list-managed-notification-events", "", false, "List Managed Notification Events")
	_notificationsCmd.Flags().BoolVarP(&_notificationsListMemberAccounts, "list-member-accounts", "", false, "List Member Accounts")
	_notificationsCmd.Flags().BoolVarP(&_notificationsListNotificationConfigurations, "list-notification-configurations", "", false, "List Notification Configurations")
	_notificationsCmd.Flags().BoolVarP(&_notificationsListNotificationEvents, "list-notification-events", "", false, "List Notification Events")
	_notificationsCmd.Flags().BoolVarP(&_notificationsListNotificationHubs, "list-notification-hubs", "", false, "List Notification Hubs")
	_notificationsCmd.Flags().BoolVarP(&_notificationsListOrganizationalUnits, "list-organizational-units", "", false, "List Organizational Units")
	_notificationsCmd.Flags().BoolVarP(&_notificationsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_notificationsCmd.Flags().BoolVarP(&_notificationsRegisterNotificationHub, "register-notification-hub", "", false, "Register Notification Hub")
	_notificationsCmd.Flags().BoolVarP(&_notificationsTagResource, "tag-resource", "", false, "Tag Resource")
	_notificationsCmd.Flags().BoolVarP(&_notificationsUntagResource, "untag-resource", "", false, "Untag Resource")
	_notificationsCmd.Flags().BoolVarP(&_notificationsUpdateEventRule, "update-event-rule", "", false, "Update Event Rule")
	_notificationsCmd.Flags().BoolVarP(&_notificationsUpdateNotificationConfiguration, "update-notification-configuration", "", false, "Update Notification Configuration")

}
