package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssmcontacts"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ssmcontactsCmd represents the ssmcontacts command
var _ssmcontactsCmd = &cobra.Command{
	Use:   "ssmcontacts",
	Short: "AWS ssmcontacts CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := ssmcontacts.NewFromConfig(cfg)
		if _ssmcontactsAcceptPage {
			ssmcontacts_AcceptPage(cfg, client)
			return
		}
		if _ssmcontactsActivateContactChannel {
			ssmcontacts_ActivateContactChannel(cfg, client)
			return
		}
		if _ssmcontactsCreateContact {
			ssmcontacts_CreateContact(cfg, client)
			return
		}
		if _ssmcontactsCreateContactChannel {
			ssmcontacts_CreateContactChannel(cfg, client)
			return
		}
		if _ssmcontactsCreateRotation {
			ssmcontacts_CreateRotation(cfg, client)
			return
		}
		if _ssmcontactsCreateRotationOverride {
			ssmcontacts_CreateRotationOverride(cfg, client)
			return
		}
		if _ssmcontactsDeactivateContactChannel {
			ssmcontacts_DeactivateContactChannel(cfg, client)
			return
		}
		if _ssmcontactsDeleteContact {
			ssmcontacts_DeleteContact(cfg, client)
			return
		}
		if _ssmcontactsDeleteContactChannel {
			ssmcontacts_DeleteContactChannel(cfg, client)
			return
		}
		if _ssmcontactsDeleteRotation {
			ssmcontacts_DeleteRotation(cfg, client)
			return
		}
		if _ssmcontactsDeleteRotationOverride {
			ssmcontacts_DeleteRotationOverride(cfg, client)
			return
		}
		if _ssmcontactsDescribeEngagement {
			ssmcontacts_DescribeEngagement(cfg, client)
			return
		}
		if _ssmcontactsDescribePage {
			ssmcontacts_DescribePage(cfg, client)
			return
		}
		if _ssmcontactsGetContact {
			ssmcontacts_GetContact(cfg, client)
			return
		}
		if _ssmcontactsGetContactChannel {
			ssmcontacts_GetContactChannel(cfg, client)
			return
		}
		if _ssmcontactsGetContactPolicy {
			ssmcontacts_GetContactPolicy(cfg, client)
			return
		}
		if _ssmcontactsGetRotation {
			ssmcontacts_GetRotation(cfg, client)
			return
		}
		if _ssmcontactsGetRotationOverride {
			ssmcontacts_GetRotationOverride(cfg, client)
			return
		}
		if _ssmcontactsListContactChannels {
			ssmcontacts_ListContactChannels(cfg, client)
			return
		}
		if _ssmcontactsListContacts {
			ssmcontacts_ListContacts(cfg, client)
			return
		}
		if _ssmcontactsListEngagements {
			ssmcontacts_ListEngagements(cfg, client)
			return
		}
		if _ssmcontactsListPageReceipts {
			ssmcontacts_ListPageReceipts(cfg, client)
			return
		}
		if _ssmcontactsListPageResolutions {
			ssmcontacts_ListPageResolutions(cfg, client)
			return
		}
		if _ssmcontactsListPagesByContact {
			ssmcontacts_ListPagesByContact(cfg, client)
			return
		}
		if _ssmcontactsListPagesByEngagement {
			ssmcontacts_ListPagesByEngagement(cfg, client)
			return
		}
		if _ssmcontactsListPreviewRotationShifts {
			ssmcontacts_ListPreviewRotationShifts(cfg, client)
			return
		}
		if _ssmcontactsListRotationOverrides {
			ssmcontacts_ListRotationOverrides(cfg, client)
			return
		}
		if _ssmcontactsListRotationShifts {
			ssmcontacts_ListRotationShifts(cfg, client)
			return
		}
		if _ssmcontactsListRotations {
			ssmcontacts_ListRotations(cfg, client)
			return
		}
		if _ssmcontactsListTagsForResource {
			ssmcontacts_ListTagsForResource(cfg, client)
			return
		}
		if _ssmcontactsPutContactPolicy {
			ssmcontacts_PutContactPolicy(cfg, client)
			return
		}
		if _ssmcontactsSendActivationCode {
			ssmcontacts_SendActivationCode(cfg, client)
			return
		}
		if _ssmcontactsStartEngagement {
			ssmcontacts_StartEngagement(cfg, client)
			return
		}
		if _ssmcontactsStopEngagement {
			ssmcontacts_StopEngagement(cfg, client)
			return
		}
		if _ssmcontactsTagResource {
			ssmcontacts_TagResource(cfg, client)
			return
		}
		if _ssmcontactsUntagResource {
			ssmcontacts_UntagResource(cfg, client)
			return
		}
		if _ssmcontactsUpdateContact {
			ssmcontacts_UpdateContact(cfg, client)
			return
		}
		if _ssmcontactsUpdateContactChannel {
			ssmcontacts_UpdateContactChannel(cfg, client)
			return
		}
		if _ssmcontactsUpdateRotation {
			ssmcontacts_UpdateRotation(cfg, client)
			return
		}

	},
}

var (
	_ssmcontactsAcceptPage                bool
	_ssmcontactsActivateContactChannel    bool
	_ssmcontactsCreateContact             bool
	_ssmcontactsCreateContactChannel      bool
	_ssmcontactsCreateRotation            bool
	_ssmcontactsCreateRotationOverride    bool
	_ssmcontactsDeactivateContactChannel  bool
	_ssmcontactsDeleteContact             bool
	_ssmcontactsDeleteContactChannel      bool
	_ssmcontactsDeleteRotation            bool
	_ssmcontactsDeleteRotationOverride    bool
	_ssmcontactsDescribeEngagement        bool
	_ssmcontactsDescribePage              bool
	_ssmcontactsGetContact                bool
	_ssmcontactsGetContactChannel         bool
	_ssmcontactsGetContactPolicy          bool
	_ssmcontactsGetRotation               bool
	_ssmcontactsGetRotationOverride       bool
	_ssmcontactsListContactChannels       bool
	_ssmcontactsListContacts              bool
	_ssmcontactsListEngagements           bool
	_ssmcontactsListPageReceipts          bool
	_ssmcontactsListPageResolutions       bool
	_ssmcontactsListPagesByContact        bool
	_ssmcontactsListPagesByEngagement     bool
	_ssmcontactsListPreviewRotationShifts bool
	_ssmcontactsListRotationOverrides     bool
	_ssmcontactsListRotationShifts        bool
	_ssmcontactsListRotations             bool
	_ssmcontactsListTagsForResource       bool
	_ssmcontactsPutContactPolicy          bool
	_ssmcontactsSendActivationCode        bool
	_ssmcontactsStartEngagement           bool
	_ssmcontactsStopEngagement            bool
	_ssmcontactsTagResource               bool
	_ssmcontactsUntagResource             bool
	_ssmcontactsUpdateContact             bool
	_ssmcontactsUpdateContactChannel      bool
	_ssmcontactsUpdateRotation            bool

	_ssmcontactsAcceptCode           string
	_ssmcontactsAcceptCodeValidation string
	_ssmcontactsAcceptType           string
	_ssmcontactsActivationCode       string
	_ssmcontactsAlias                string
	_ssmcontactsAliasPrefix          string
	_ssmcontactsContactArn           string
	_ssmcontactsContactChannelId     string
	_ssmcontactsContactId            string
	_ssmcontactsContactIds           []string
	_ssmcontactsContent              string
	_ssmcontactsDeferActivation      string
	_ssmcontactsDeliveryAddress      string
	_ssmcontactsDisplayName          string
	_ssmcontactsEndTime              string
	_ssmcontactsEngagementId         string
	_ssmcontactsIdempotencyToken     string
	_ssmcontactsIncidentId           string
	_ssmcontactsMaxResults           string
	_ssmcontactsMembers              []string
	_ssmcontactsName                 string
	_ssmcontactsNewContactIds        []string
	_ssmcontactsNextToken            string
	_ssmcontactsNote                 string
	_ssmcontactsOverrides            string
	_ssmcontactsPageId               string
	_ssmcontactsPlan                 string
	_ssmcontactsPolicy               string
	_ssmcontactsPublicContent        string
	_ssmcontactsPublicSubject        string
	_ssmcontactsReason               string
	_ssmcontactsRecurrence           string
	_ssmcontactsResourceARN          string
	_ssmcontactsRotationId           string
	_ssmcontactsRotationNamePrefix   string
	_ssmcontactsRotationOverrideId   string
	_ssmcontactsRotationStartTime    string
	_ssmcontactsSender               string
	_ssmcontactsStartTime            string
	_ssmcontactsSubject              string
	_ssmcontactsTagKeys              []string
	_ssmcontactsTags                 string
	_ssmcontactsTimeRangeValue       string
	_ssmcontactsTimeZoneId           string
	_ssmcontactsType                 string
)

// Used to acknowledge an engagement to a contact channel during an incident.
func ssmcontacts_AcceptPage(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.AcceptPageInput{
		// AcceptCode: *string, // Required
		// AcceptType: types.AcceptType, // Required
		// PageId: *string, // Required
	}

	if len(_ssmcontactsAcceptCode) > 0 {
		input.AcceptCode = aws.String(_ssmcontactsAcceptCode)
	}
	if len(_ssmcontactsAcceptType) > 0 {
		if err := assignInputField(input, "AcceptType", _ssmcontactsAcceptType); err != nil {
			log.Errorf("invalid --accept-type: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsPageId) > 0 {
		input.PageId = aws.String(_ssmcontactsPageId)
	}
	if len(_ssmcontactsAcceptCodeValidation) > 0 {
		if err := assignInputField(input, "AcceptCodeValidation", _ssmcontactsAcceptCodeValidation); err != nil {
			log.Errorf("invalid --accept-code-validation: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsContactChannelId) > 0 {
		input.ContactChannelId = aws.String(_ssmcontactsContactChannelId)
	}
	if len(_ssmcontactsNote) > 0 {
		input.Note = aws.String(_ssmcontactsNote)
	}

	if resp, err := client.AcceptPage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Activates a contact's contact channel. Incident Manager can't engage a contact
// until the contact channel has been activated.
func ssmcontacts_ActivateContactChannel(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.ActivateContactChannelInput{
		// ActivationCode: *string, // Required
		// ContactChannelId: *string, // Required
	}

	if len(_ssmcontactsActivationCode) > 0 {
		input.ActivationCode = aws.String(_ssmcontactsActivationCode)
	}
	if len(_ssmcontactsContactChannelId) > 0 {
		input.ContactChannelId = aws.String(_ssmcontactsContactChannelId)
	}

	if resp, err := client.ActivateContactChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Contacts are either the contacts that Incident Manager engages during an
// incident or the escalation plans that Incident Manager uses to engage contacts
// in phases during an incident.
func ssmcontacts_CreateContact(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.CreateContactInput{
		// Alias: *string, // Required
		// Plan: *types.Plan, // Required
		// Type: types.ContactType, // Required
	}

	if len(_ssmcontactsAlias) > 0 {
		input.Alias = aws.String(_ssmcontactsAlias)
	}
	if len(_ssmcontactsPlan) > 0 {
		if err := assignInputField(input, "Plan", _ssmcontactsPlan); err != nil {
			log.Errorf("invalid --plan: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsType) > 0 {
		if err := assignInputField(input, "Type", _ssmcontactsType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsDisplayName) > 0 {
		input.DisplayName = aws.String(_ssmcontactsDisplayName)
	}
	if len(_ssmcontactsIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_ssmcontactsIdempotencyToken)
	}
	if len(_ssmcontactsTags) > 0 {
		if err := assignInputField(input, "Tags", _ssmcontactsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A contact channel is the method that Incident Manager uses to engage your
// contact.
func ssmcontacts_CreateContactChannel(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.CreateContactChannelInput{
		// ContactId: *string, // Required
		// DeliveryAddress: *types.ContactChannelAddress, // Required
		// Name: *string, // Required
		// Type: types.ChannelType, // Required
	}

	if len(_ssmcontactsContactId) > 0 {
		input.ContactId = aws.String(_ssmcontactsContactId)
	}
	if len(_ssmcontactsDeliveryAddress) > 0 {
		if err := assignInputField(input, "DeliveryAddress", _ssmcontactsDeliveryAddress); err != nil {
			log.Errorf("invalid --delivery-address: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsName) > 0 {
		input.Name = aws.String(_ssmcontactsName)
	}
	if len(_ssmcontactsType) > 0 {
		if err := assignInputField(input, "Type", _ssmcontactsType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsDeferActivation) > 0 {
		if err := assignInputField(input, "DeferActivation", _ssmcontactsDeferActivation); err != nil {
			log.Errorf("invalid --defer-activation: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_ssmcontactsIdempotencyToken)
	}

	if resp, err := client.CreateContactChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a rotation in an on-call schedule.
func ssmcontacts_CreateRotation(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.CreateRotationInput{
		// ContactIds: []string, // Required
		// Name: *string, // Required
		// Recurrence: *types.RecurrenceSettings, // Required
		// TimeZoneId: *string, // Required
	}

	if len(_ssmcontactsContactIds) > 0 {
		input.ContactIds = append([]string(nil), _ssmcontactsContactIds...)
	}
	if len(_ssmcontactsName) > 0 {
		input.Name = aws.String(_ssmcontactsName)
	}
	if len(_ssmcontactsRecurrence) > 0 {
		if err := assignInputField(input, "Recurrence", _ssmcontactsRecurrence); err != nil {
			log.Errorf("invalid --recurrence: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsTimeZoneId) > 0 {
		input.TimeZoneId = aws.String(_ssmcontactsTimeZoneId)
	}
	if len(_ssmcontactsIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_ssmcontactsIdempotencyToken)
	}
	if len(_ssmcontactsStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _ssmcontactsStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsTags) > 0 {
		if err := assignInputField(input, "Tags", _ssmcontactsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRotation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an override for a rotation in an on-call schedule.
func ssmcontacts_CreateRotationOverride(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.CreateRotationOverrideInput{
		// EndTime: *time.Time, // Required
		// NewContactIds: []string, // Required
		// RotationId: *string, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_ssmcontactsEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _ssmcontactsEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsNewContactIds) > 0 {
		input.NewContactIds = append([]string(nil), _ssmcontactsNewContactIds...)
	}
	if len(_ssmcontactsRotationId) > 0 {
		input.RotationId = aws.String(_ssmcontactsRotationId)
	}
	if len(_ssmcontactsStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _ssmcontactsStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_ssmcontactsIdempotencyToken)
	}

	if resp, err := client.CreateRotationOverride(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// To no longer receive Incident Manager engagements to a contact channel, you can
// deactivate the channel.
func ssmcontacts_DeactivateContactChannel(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.DeactivateContactChannelInput{
		// ContactChannelId: *string, // Required
	}

	if len(_ssmcontactsContactChannelId) > 0 {
		input.ContactChannelId = aws.String(_ssmcontactsContactChannelId)
	}

	if resp, err := client.DeactivateContactChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// To remove a contact from Incident Manager, you can delete the contact. However,
// deleting a contact does not remove it from escalation plans and related response
// plans. Deleting an escalation plan also does not remove it from all related
// response plans. To modify an escalation plan, we recommend using the UpdateContactaction to
// specify a different existing contact.
func ssmcontacts_DeleteContact(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.DeleteContactInput{
		// ContactId: *string, // Required
	}

	if len(_ssmcontactsContactId) > 0 {
		input.ContactId = aws.String(_ssmcontactsContactId)
	}

	if resp, err := client.DeleteContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// To stop receiving engagements on a contact channel, you can delete the channel
// from a contact. Deleting the contact channel does not remove it from the
// contact's engagement plan, but the stage that includes the channel will be
// ignored. If you delete the only contact channel for a contact, you'll no longer
// be able to engage that contact during an incident.
func ssmcontacts_DeleteContactChannel(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.DeleteContactChannelInput{
		// ContactChannelId: *string, // Required
	}

	if len(_ssmcontactsContactChannelId) > 0 {
		input.ContactChannelId = aws.String(_ssmcontactsContactChannelId)
	}

	if resp, err := client.DeleteContactChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a rotation from the system. If a rotation belongs to more than one
// on-call schedule, this operation deletes it from all of them.
func ssmcontacts_DeleteRotation(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.DeleteRotationInput{
		// RotationId: *string, // Required
	}

	if len(_ssmcontactsRotationId) > 0 {
		input.RotationId = aws.String(_ssmcontactsRotationId)
	}

	if resp, err := client.DeleteRotation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing override for an on-call rotation.
func ssmcontacts_DeleteRotationOverride(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.DeleteRotationOverrideInput{
		// RotationId: *string, // Required
		// RotationOverrideId: *string, // Required
	}

	if len(_ssmcontactsRotationId) > 0 {
		input.RotationId = aws.String(_ssmcontactsRotationId)
	}
	if len(_ssmcontactsRotationOverrideId) > 0 {
		input.RotationOverrideId = aws.String(_ssmcontactsRotationOverrideId)
	}

	if resp, err := client.DeleteRotationOverride(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Incident Manager uses engagements to engage contacts and escalation plans
// during an incident. Use this command to describe the engagement that occurred
// during an incident.
func ssmcontacts_DescribeEngagement(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.DescribeEngagementInput{
		// EngagementId: *string, // Required
	}

	if len(_ssmcontactsEngagementId) > 0 {
		input.EngagementId = aws.String(_ssmcontactsEngagementId)
	}

	if resp, err := client.DescribeEngagement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists details of the engagement to a contact channel.
func ssmcontacts_DescribePage(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.DescribePageInput{
		// PageId: *string, // Required
	}

	if len(_ssmcontactsPageId) > 0 {
		input.PageId = aws.String(_ssmcontactsPageId)
	}

	if resp, err := client.DescribePage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the specified contact or escalation plan.
func ssmcontacts_GetContact(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.GetContactInput{
		// ContactId: *string, // Required
	}

	if len(_ssmcontactsContactId) > 0 {
		input.ContactId = aws.String(_ssmcontactsContactId)
	}

	if resp, err := client.GetContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List details about a specific contact channel.
func ssmcontacts_GetContactChannel(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.GetContactChannelInput{
		// ContactChannelId: *string, // Required
	}

	if len(_ssmcontactsContactChannelId) > 0 {
		input.ContactChannelId = aws.String(_ssmcontactsContactChannelId)
	}

	if resp, err := client.GetContactChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the resource policies attached to the specified contact or escalation
// plan.
func ssmcontacts_GetContactPolicy(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.GetContactPolicyInput{
		// ContactArn: *string, // Required
	}

	if len(_ssmcontactsContactArn) > 0 {
		input.ContactArn = aws.String(_ssmcontactsContactArn)
	}

	if resp, err := client.GetContactPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an on-call rotation.
func ssmcontacts_GetRotation(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.GetRotationInput{
		// RotationId: *string, // Required
	}

	if len(_ssmcontactsRotationId) > 0 {
		input.RotationId = aws.String(_ssmcontactsRotationId)
	}

	if resp, err := client.GetRotation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an override to an on-call rotation.
func ssmcontacts_GetRotationOverride(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.GetRotationOverrideInput{
		// RotationId: *string, // Required
		// RotationOverrideId: *string, // Required
	}

	if len(_ssmcontactsRotationId) > 0 {
		input.RotationId = aws.String(_ssmcontactsRotationId)
	}
	if len(_ssmcontactsRotationOverrideId) > 0 {
		input.RotationOverrideId = aws.String(_ssmcontactsRotationOverrideId)
	}

	if resp, err := client.GetRotationOverride(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all contact channels for the specified contact.
func ssmcontacts_ListContactChannels(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.ListContactChannelsInput{
		// ContactId: *string, // Required
	}

	if len(_ssmcontactsContactId) > 0 {
		input.ContactId = aws.String(_ssmcontactsContactId)
	}
	if len(_ssmcontactsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmcontactsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsNextToken) > 0 {
		input.NextToken = aws.String(_ssmcontactsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListContactChannels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssmcontacts.ListContactChannelsOutput
	p := ssmcontacts.NewListContactChannelsPaginator(client, input)
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

// Lists all contacts and escalation plans in Incident Manager.
func ssmcontacts_ListContacts(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.ListContactsInput{}

	if len(_ssmcontactsAliasPrefix) > 0 {
		input.AliasPrefix = aws.String(_ssmcontactsAliasPrefix)
	}
	if len(_ssmcontactsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmcontactsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsNextToken) > 0 {
		input.NextToken = aws.String(_ssmcontactsNextToken)
	}
	if len(_ssmcontactsType) > 0 {
		if err := assignInputField(input, "Type", _ssmcontactsType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListContacts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssmcontacts.ListContactsOutput
	p := ssmcontacts.NewListContactsPaginator(client, input)
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

// Lists all engagements that have happened in an incident.
func ssmcontacts_ListEngagements(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.ListEngagementsInput{}

	if len(_ssmcontactsIncidentId) > 0 {
		input.IncidentId = aws.String(_ssmcontactsIncidentId)
	}
	if len(_ssmcontactsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmcontactsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsNextToken) > 0 {
		input.NextToken = aws.String(_ssmcontactsNextToken)
	}
	if len(_ssmcontactsTimeRangeValue) > 0 {
		if err := assignInputField(input, "TimeRangeValue", _ssmcontactsTimeRangeValue); err != nil {
			log.Errorf("invalid --time-range-value: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListEngagements(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssmcontacts.ListEngagementsOutput
	p := ssmcontacts.NewListEngagementsPaginator(client, input)
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

// Lists all of the engagements to contact channels that have been acknowledged.
func ssmcontacts_ListPageReceipts(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.ListPageReceiptsInput{
		// PageId: *string, // Required
	}

	if len(_ssmcontactsPageId) > 0 {
		input.PageId = aws.String(_ssmcontactsPageId)
	}
	if len(_ssmcontactsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmcontactsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsNextToken) > 0 {
		input.NextToken = aws.String(_ssmcontactsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPageReceipts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssmcontacts.ListPageReceiptsOutput
	p := ssmcontacts.NewListPageReceiptsPaginator(client, input)
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

// Returns the resolution path of an engagement. For example, the escalation plan
// engaged in an incident might target an on-call schedule that includes several
// contacts in a rotation, but just one contact on-call when the incident starts.
// The resolution path indicates the hierarchy of escalation plan > on-call
// schedule > contact.
func ssmcontacts_ListPageResolutions(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.ListPageResolutionsInput{
		// PageId: *string, // Required
	}

	if len(_ssmcontactsPageId) > 0 {
		input.PageId = aws.String(_ssmcontactsPageId)
	}
	if len(_ssmcontactsNextToken) > 0 {
		input.NextToken = aws.String(_ssmcontactsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPageResolutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssmcontacts.ListPageResolutionsOutput
	p := ssmcontacts.NewListPageResolutionsPaginator(client, input)
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

// Lists the engagements to a contact's contact channels.
func ssmcontacts_ListPagesByContact(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.ListPagesByContactInput{
		// ContactId: *string, // Required
	}

	if len(_ssmcontactsContactId) > 0 {
		input.ContactId = aws.String(_ssmcontactsContactId)
	}
	if len(_ssmcontactsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmcontactsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsNextToken) > 0 {
		input.NextToken = aws.String(_ssmcontactsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPagesByContact(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssmcontacts.ListPagesByContactOutput
	p := ssmcontacts.NewListPagesByContactPaginator(client, input)
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

// Lists the engagements to contact channels that occurred by engaging a contact.
func ssmcontacts_ListPagesByEngagement(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.ListPagesByEngagementInput{
		// EngagementId: *string, // Required
	}

	if len(_ssmcontactsEngagementId) > 0 {
		input.EngagementId = aws.String(_ssmcontactsEngagementId)
	}
	if len(_ssmcontactsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmcontactsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsNextToken) > 0 {
		input.NextToken = aws.String(_ssmcontactsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPagesByEngagement(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssmcontacts.ListPagesByEngagementOutput
	p := ssmcontacts.NewListPagesByEngagementPaginator(client, input)
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

// Returns a list of shifts based on rotation configuration parameters.
// The Incident Manager primarily uses this operation to populate the Preview
// calendar. It is not typically run by end users.
func ssmcontacts_ListPreviewRotationShifts(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.ListPreviewRotationShiftsInput{
		// EndTime: *time.Time, // Required
		// Members: []string, // Required
		// Recurrence: *types.RecurrenceSettings, // Required
		// TimeZoneId: *string, // Required
	}

	if len(_ssmcontactsEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _ssmcontactsEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsMembers) > 0 {
		input.Members = append([]string(nil), _ssmcontactsMembers...)
	}
	if len(_ssmcontactsRecurrence) > 0 {
		if err := assignInputField(input, "Recurrence", _ssmcontactsRecurrence); err != nil {
			log.Errorf("invalid --recurrence: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsTimeZoneId) > 0 {
		input.TimeZoneId = aws.String(_ssmcontactsTimeZoneId)
	}
	if len(_ssmcontactsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmcontactsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsNextToken) > 0 {
		input.NextToken = aws.String(_ssmcontactsNextToken)
	}
	if len(_ssmcontactsOverrides) > 0 {
		if err := assignInputField(input, "Overrides", _ssmcontactsOverrides); err != nil {
			log.Errorf("invalid --overrides: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsRotationStartTime) > 0 {
		if err := assignInputField(input, "RotationStartTime", _ssmcontactsRotationStartTime); err != nil {
			log.Errorf("invalid --rotation-start-time: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _ssmcontactsStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListPreviewRotationShifts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssmcontacts.ListPreviewRotationShiftsOutput
	p := ssmcontacts.NewListPreviewRotationShiftsPaginator(client, input)
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

// Retrieves a list of overrides currently specified for an on-call rotation.
func ssmcontacts_ListRotationOverrides(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.ListRotationOverridesInput{
		// EndTime: *time.Time, // Required
		// RotationId: *string, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_ssmcontactsEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _ssmcontactsEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsRotationId) > 0 {
		input.RotationId = aws.String(_ssmcontactsRotationId)
	}
	if len(_ssmcontactsStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _ssmcontactsStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmcontactsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsNextToken) > 0 {
		input.NextToken = aws.String(_ssmcontactsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRotationOverrides(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssmcontacts.ListRotationOverridesOutput
	p := ssmcontacts.NewListRotationOverridesPaginator(client, input)
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

// Returns a list of shifts generated by an existing rotation in the system.
func ssmcontacts_ListRotationShifts(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.ListRotationShiftsInput{
		// EndTime: *time.Time, // Required
		// RotationId: *string, // Required
	}

	if len(_ssmcontactsEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _ssmcontactsEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsRotationId) > 0 {
		input.RotationId = aws.String(_ssmcontactsRotationId)
	}
	if len(_ssmcontactsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmcontactsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsNextToken) > 0 {
		input.NextToken = aws.String(_ssmcontactsNextToken)
	}
	if len(_ssmcontactsStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _ssmcontactsStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListRotationShifts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssmcontacts.ListRotationShiftsOutput
	p := ssmcontacts.NewListRotationShiftsPaginator(client, input)
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

// Retrieves a list of on-call rotations.
func ssmcontacts_ListRotations(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.ListRotationsInput{}

	if len(_ssmcontactsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssmcontactsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsNextToken) > 0 {
		input.NextToken = aws.String(_ssmcontactsNextToken)
	}
	if len(_ssmcontactsRotationNamePrefix) > 0 {
		input.RotationNamePrefix = aws.String(_ssmcontactsRotationNamePrefix)
	}

	if disablePaginator() {
		if resp, err := client.ListRotations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ssmcontacts.ListRotationsOutput
	p := ssmcontacts.NewListRotationsPaginator(client, input)
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

// Lists the tags of a contact, escalation plan, rotation, or on-call schedule.
func ssmcontacts_ListTagsForResource(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_ssmcontactsResourceARN) > 0 {
		input.ResourceARN = aws.String(_ssmcontactsResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a resource policy to the specified contact or escalation plan. The
// resource policy is used to share the contact or escalation plan using Resource
// Access Manager (RAM). For more information about cross-account sharing, see [Setting up cross-account functionality].
//
// [Setting up cross-account functionality]: https://docs.aws.amazon.com/incident-manager/latest/userguide/xa.html
func ssmcontacts_PutContactPolicy(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.PutContactPolicyInput{
		// ContactArn: *string, // Required
		// Policy: *string, // Required
	}

	if len(_ssmcontactsContactArn) > 0 {
		input.ContactArn = aws.String(_ssmcontactsContactArn)
	}
	if len(_ssmcontactsPolicy) > 0 {
		input.Policy = aws.String(_ssmcontactsPolicy)
	}

	if resp, err := client.PutContactPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends an activation code to a contact channel. The contact can use this code to
// activate the contact channel in the console or with the ActivateChannel
// operation. Incident Manager can't engage a contact channel until it has been
// activated.
func ssmcontacts_SendActivationCode(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.SendActivationCodeInput{
		// ContactChannelId: *string, // Required
	}

	if len(_ssmcontactsContactChannelId) > 0 {
		input.ContactChannelId = aws.String(_ssmcontactsContactChannelId)
	}

	if resp, err := client.SendActivationCode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an engagement to a contact or escalation plan. The engagement engages
// each contact specified in the incident.
func ssmcontacts_StartEngagement(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.StartEngagementInput{
		// ContactId: *string, // Required
		// Content: *string, // Required
		// Sender: *string, // Required
		// Subject: *string, // Required
	}

	if len(_ssmcontactsContactId) > 0 {
		input.ContactId = aws.String(_ssmcontactsContactId)
	}
	if len(_ssmcontactsContent) > 0 {
		input.Content = aws.String(_ssmcontactsContent)
	}
	if len(_ssmcontactsSender) > 0 {
		input.Sender = aws.String(_ssmcontactsSender)
	}
	if len(_ssmcontactsSubject) > 0 {
		input.Subject = aws.String(_ssmcontactsSubject)
	}
	if len(_ssmcontactsIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_ssmcontactsIdempotencyToken)
	}
	if len(_ssmcontactsIncidentId) > 0 {
		input.IncidentId = aws.String(_ssmcontactsIncidentId)
	}
	if len(_ssmcontactsPublicContent) > 0 {
		input.PublicContent = aws.String(_ssmcontactsPublicContent)
	}
	if len(_ssmcontactsPublicSubject) > 0 {
		input.PublicSubject = aws.String(_ssmcontactsPublicSubject)
	}

	if resp, err := client.StartEngagement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops an engagement before it finishes the final stage of the escalation plan
// or engagement plan. Further contacts aren't engaged.
func ssmcontacts_StopEngagement(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.StopEngagementInput{
		// EngagementId: *string, // Required
	}

	if len(_ssmcontactsEngagementId) > 0 {
		input.EngagementId = aws.String(_ssmcontactsEngagementId)
	}
	if len(_ssmcontactsReason) > 0 {
		input.Reason = aws.String(_ssmcontactsReason)
	}

	if resp, err := client.StopEngagement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tags a contact or escalation plan. You can tag only contacts and escalation
// plans in the first region of your replication set.
func ssmcontacts_TagResource(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_ssmcontactsResourceARN) > 0 {
		input.ResourceARN = aws.String(_ssmcontactsResourceARN)
	}
	if len(_ssmcontactsTags) > 0 {
		if err := assignInputField(input, "Tags", _ssmcontactsTags); err != nil {
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

// Removes tags from the specified resource.
func ssmcontacts_UntagResource(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_ssmcontactsResourceARN) > 0 {
		input.ResourceARN = aws.String(_ssmcontactsResourceARN)
	}
	if len(_ssmcontactsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _ssmcontactsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the contact or escalation plan specified.
func ssmcontacts_UpdateContact(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.UpdateContactInput{
		// ContactId: *string, // Required
	}

	if len(_ssmcontactsContactId) > 0 {
		input.ContactId = aws.String(_ssmcontactsContactId)
	}
	if len(_ssmcontactsDisplayName) > 0 {
		input.DisplayName = aws.String(_ssmcontactsDisplayName)
	}
	if len(_ssmcontactsPlan) > 0 {
		if err := assignInputField(input, "Plan", _ssmcontactsPlan); err != nil {
			log.Errorf("invalid --plan: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a contact's contact channel.
func ssmcontacts_UpdateContactChannel(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.UpdateContactChannelInput{
		// ContactChannelId: *string, // Required
	}

	if len(_ssmcontactsContactChannelId) > 0 {
		input.ContactChannelId = aws.String(_ssmcontactsContactChannelId)
	}
	if len(_ssmcontactsDeliveryAddress) > 0 {
		if err := assignInputField(input, "DeliveryAddress", _ssmcontactsDeliveryAddress); err != nil {
			log.Errorf("invalid --delivery-address: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsName) > 0 {
		input.Name = aws.String(_ssmcontactsName)
	}

	if resp, err := client.UpdateContactChannel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the information specified for an on-call rotation.
func ssmcontacts_UpdateRotation(cfg aws.Config, client *ssmcontacts.Client) {
	input := &ssmcontacts.UpdateRotationInput{
		// Recurrence: *types.RecurrenceSettings, // Required
		// RotationId: *string, // Required
	}

	if len(_ssmcontactsRecurrence) > 0 {
		if err := assignInputField(input, "Recurrence", _ssmcontactsRecurrence); err != nil {
			log.Errorf("invalid --recurrence: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsRotationId) > 0 {
		input.RotationId = aws.String(_ssmcontactsRotationId)
	}
	if len(_ssmcontactsContactIds) > 0 {
		input.ContactIds = append([]string(nil), _ssmcontactsContactIds...)
	}
	if len(_ssmcontactsStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _ssmcontactsStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_ssmcontactsTimeZoneId) > 0 {
		input.TimeZoneId = aws.String(_ssmcontactsTimeZoneId)
	}

	if resp, err := client.UpdateRotation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_ssmcontactsCmd)
	_ssmcontactsCmd.Flags().SortFlags = false

	_ssmcontactsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_ssmcontactsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_ssmcontactsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsAcceptCode, "accept-code", "", "", "Accept Code")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsAcceptCodeValidation, "accept-code-validation", "", "", "Accept Code Validation")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsAcceptType, "accept-type", "", "", "Accept Type")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsActivationCode, "activation-code", "", "", "Activation Code")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsAlias, "alias", "", "", "Alias")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsAliasPrefix, "alias-prefix", "", "", "Alias Prefix")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsContactArn, "contact-arn", "", "", "Contact ARN")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsContactChannelId, "contact-channel-id", "", "", "Contact Channel ID")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsContactId, "contact-id", "", "", "Contact ID")
	_ssmcontactsCmd.Flags().StringSliceVarP(&_ssmcontactsContactIds, "contact-ids", "", nil, "Contact Ids")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsContent, "content", "", "", "Content")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsDeferActivation, "defer-activation", "", "", "Defer Activation")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsDeliveryAddress, "delivery-address", "", "", "Delivery Address")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsDisplayName, "display-name", "", "", "Display Name")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsEndTime, "end-time", "", "", "End Time")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsEngagementId, "engagement-id", "", "", "Engagement ID")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsIdempotencyToken, "idempotency-token", "", "", "Idempotency Token")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsIncidentId, "incident-id", "", "", "Incident ID")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsMaxResults, "max-results", "", "", "Max Results")
	_ssmcontactsCmd.Flags().StringSliceVarP(&_ssmcontactsMembers, "members", "", nil, "Members")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsName, "name", "", "", "Name")
	_ssmcontactsCmd.Flags().StringSliceVarP(&_ssmcontactsNewContactIds, "new-contact-ids", "", nil, "New Contact Ids")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsNextToken, "next-token", "", "", "Next Token")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsNote, "note", "", "", "Note")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsOverrides, "overrides", "", "", "Overrides")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsPageId, "page-id", "", "", "Page ID")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsPlan, "plan", "", "", "Plan")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsPolicy, "policy", "", "", "Policy")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsPublicContent, "public-content", "", "", "Public Content")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsPublicSubject, "public-subject", "", "", "Public Subject")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsReason, "reason", "", "", "Reason")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsRecurrence, "recurrence", "", "", "Recurrence")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsResourceARN, "resource-arn", "", "", "Resource ARN")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsRotationId, "rotation-id", "", "", "Rotation ID")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsRotationNamePrefix, "rotation-name-prefix", "", "", "Rotation Name Prefix")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsRotationOverrideId, "rotation-override-id", "", "", "Rotation Override ID")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsRotationStartTime, "rotation-start-time", "", "", "Rotation Start Time")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsSender, "sender", "", "", "Sender")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsStartTime, "start-time", "", "", "Start Time")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsSubject, "subject", "", "", "Subject")
	_ssmcontactsCmd.Flags().StringSliceVarP(&_ssmcontactsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsTags, "tags", "", "", "Tags")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsTimeRangeValue, "time-range-value", "", "", "Time Range Value")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsTimeZoneId, "time-zone-id", "", "", "Time Zone ID")
	_ssmcontactsCmd.Flags().StringVarP(&_ssmcontactsType, "type", "", "", "Type")

	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsAcceptPage, "accept-page", "", false, "Accept Page")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsActivateContactChannel, "activate-contact-channel", "", false, "Activate Contact Channel")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsCreateContact, "create-contact", "", false, "Create Contact")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsCreateContactChannel, "create-contact-channel", "", false, "Create Contact Channel")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsCreateRotation, "create-rotation", "", false, "Create Rotation")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsCreateRotationOverride, "create-rotation-override", "", false, "Create Rotation Override")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsDeactivateContactChannel, "deactivate-contact-channel", "", false, "Deactivate Contact Channel")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsDeleteContact, "delete-contact", "", false, "Delete Contact")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsDeleteContactChannel, "delete-contact-channel", "", false, "Delete Contact Channel")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsDeleteRotation, "delete-rotation", "", false, "Delete Rotation")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsDeleteRotationOverride, "delete-rotation-override", "", false, "Delete Rotation Override")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsDescribeEngagement, "describe-engagement", "", false, "Describe Engagement")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsDescribePage, "describe-page", "", false, "Describe Page")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsGetContact, "get-contact", "", false, "Get Contact")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsGetContactChannel, "get-contact-channel", "", false, "Get Contact Channel")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsGetContactPolicy, "get-contact-policy", "", false, "Get Contact Policy")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsGetRotation, "get-rotation", "", false, "Get Rotation")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsGetRotationOverride, "get-rotation-override", "", false, "Get Rotation Override")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsListContactChannels, "list-contact-channels", "", false, "List Contact Channels")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsListContacts, "list-contacts", "", false, "List Contacts")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsListEngagements, "list-engagements", "", false, "List Engagements")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsListPageReceipts, "list-page-receipts", "", false, "List Page Receipts")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsListPageResolutions, "list-page-resolutions", "", false, "List Page Resolutions")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsListPagesByContact, "list-pages-by-contact", "", false, "List Pages By Contact")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsListPagesByEngagement, "list-pages-by-engagement", "", false, "List Pages By Engagement")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsListPreviewRotationShifts, "list-preview-rotation-shifts", "", false, "List Preview Rotation Shifts")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsListRotationOverrides, "list-rotation-overrides", "", false, "List Rotation Overrides")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsListRotationShifts, "list-rotation-shifts", "", false, "List Rotation Shifts")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsListRotations, "list-rotations", "", false, "List Rotations")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsPutContactPolicy, "put-contact-policy", "", false, "Put Contact Policy")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsSendActivationCode, "send-activation-code", "", false, "Send Activation Code")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsStartEngagement, "start-engagement", "", false, "Start Engagement")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsStopEngagement, "stop-engagement", "", false, "Stop Engagement")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsTagResource, "tag-resource", "", false, "Tag Resource")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsUntagResource, "untag-resource", "", false, "Untag Resource")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsUpdateContact, "update-contact", "", false, "Update Contact")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsUpdateContactChannel, "update-contact-channel", "", false, "Update Contact Channel")
	_ssmcontactsCmd.Flags().BoolVarP(&_ssmcontactsUpdateRotation, "update-rotation", "", false, "Update Rotation")

}
