package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/socialmessaging"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// socialmessagingCmd represents the socialmessaging command
var _socialmessagingCmd = &cobra.Command{
	Use:   "socialmessaging",
	Short: "AWS socialmessaging CLI",
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
		client := socialmessaging.NewFromConfig(cfg)
		if _socialmessagingAssociateWhatsAppBusinessAccount {
			socialmessaging_AssociateWhatsAppBusinessAccount(cfg, client)
			return
		}
		if _socialmessagingCreateWhatsAppMessageTemplate {
			socialmessaging_CreateWhatsAppMessageTemplate(cfg, client)
			return
		}
		if _socialmessagingCreateWhatsAppMessageTemplateFromLibrary {
			socialmessaging_CreateWhatsAppMessageTemplateFromLibrary(cfg, client)
			return
		}
		if _socialmessagingCreateWhatsAppMessageTemplateMedia {
			socialmessaging_CreateWhatsAppMessageTemplateMedia(cfg, client)
			return
		}
		if _socialmessagingDeleteWhatsAppMessageMedia {
			socialmessaging_DeleteWhatsAppMessageMedia(cfg, client)
			return
		}
		if _socialmessagingDeleteWhatsAppMessageTemplate {
			socialmessaging_DeleteWhatsAppMessageTemplate(cfg, client)
			return
		}
		if _socialmessagingDisassociateWhatsAppBusinessAccount {
			socialmessaging_DisassociateWhatsAppBusinessAccount(cfg, client)
			return
		}
		if _socialmessagingGetLinkedWhatsAppBusinessAccount {
			socialmessaging_GetLinkedWhatsAppBusinessAccount(cfg, client)
			return
		}
		if _socialmessagingGetLinkedWhatsAppBusinessAccountPhoneNumber {
			socialmessaging_GetLinkedWhatsAppBusinessAccountPhoneNumber(cfg, client)
			return
		}
		if _socialmessagingGetWhatsAppMessageMedia {
			socialmessaging_GetWhatsAppMessageMedia(cfg, client)
			return
		}
		if _socialmessagingGetWhatsAppMessageTemplate {
			socialmessaging_GetWhatsAppMessageTemplate(cfg, client)
			return
		}
		if _socialmessagingListLinkedWhatsAppBusinessAccounts {
			socialmessaging_ListLinkedWhatsAppBusinessAccounts(cfg, client)
			return
		}
		if _socialmessagingListTagsForResource {
			socialmessaging_ListTagsForResource(cfg, client)
			return
		}
		if _socialmessagingListWhatsAppMessageTemplates {
			socialmessaging_ListWhatsAppMessageTemplates(cfg, client)
			return
		}
		if _socialmessagingListWhatsAppTemplateLibrary {
			socialmessaging_ListWhatsAppTemplateLibrary(cfg, client)
			return
		}
		if _socialmessagingPostWhatsAppMessageMedia {
			socialmessaging_PostWhatsAppMessageMedia(cfg, client)
			return
		}
		if _socialmessagingPutWhatsAppBusinessAccountEventDestinations {
			socialmessaging_PutWhatsAppBusinessAccountEventDestinations(cfg, client)
			return
		}
		if _socialmessagingSendWhatsAppMessage {
			socialmessaging_SendWhatsAppMessage(cfg, client)
			return
		}
		if _socialmessagingTagResource {
			socialmessaging_TagResource(cfg, client)
			return
		}
		if _socialmessagingUntagResource {
			socialmessaging_UntagResource(cfg, client)
			return
		}
		if _socialmessagingUpdateWhatsAppMessageTemplate {
			socialmessaging_UpdateWhatsAppMessageTemplate(cfg, client)
			return
		}

	},
}

var (
	_socialmessagingAssociateWhatsAppBusinessAccount            bool
	_socialmessagingCreateWhatsAppMessageTemplate               bool
	_socialmessagingCreateWhatsAppMessageTemplateFromLibrary    bool
	_socialmessagingCreateWhatsAppMessageTemplateMedia          bool
	_socialmessagingDeleteWhatsAppMessageMedia                  bool
	_socialmessagingDeleteWhatsAppMessageTemplate               bool
	_socialmessagingDisassociateWhatsAppBusinessAccount         bool
	_socialmessagingGetLinkedWhatsAppBusinessAccount            bool
	_socialmessagingGetLinkedWhatsAppBusinessAccountPhoneNumber bool
	_socialmessagingGetWhatsAppMessageMedia                     bool
	_socialmessagingGetWhatsAppMessageTemplate                  bool
	_socialmessagingListLinkedWhatsAppBusinessAccounts          bool
	_socialmessagingListTagsForResource                         bool
	_socialmessagingListWhatsAppMessageTemplates                bool
	_socialmessagingListWhatsAppTemplateLibrary                 bool
	_socialmessagingPostWhatsAppMessageMedia                    bool
	_socialmessagingPutWhatsAppBusinessAccountEventDestinations bool
	_socialmessagingSendWhatsAppMessage                         bool
	_socialmessagingTagResource                                 bool
	_socialmessagingUntagResource                               bool
	_socialmessagingUpdateWhatsAppMessageTemplate               bool

	_socialmessagingCtaUrlLinkTrackingOptedOut string
	_socialmessagingDeleteAllLanguages         string
	_socialmessagingDestinationS3File          string
	_socialmessagingDestinationS3PresignedUrl  string
	_socialmessagingEventDestinations          string
	_socialmessagingFilters                    string
	_socialmessagingId                         string
	_socialmessagingMaxResults                 string
	_socialmessagingMediaId                    string
	_socialmessagingMessage                    string
	_socialmessagingMetaApiVersion             string
	_socialmessagingMetaLibraryTemplate        string
	_socialmessagingMetaTemplateId             string
	_socialmessagingMetadataOnly               string
	_socialmessagingNextToken                  string
	_socialmessagingOriginationPhoneNumberId   string
	_socialmessagingParameterFormat            string
	_socialmessagingResourceArn                string
	_socialmessagingSetupFinalization          string
	_socialmessagingSignupCallback             string
	_socialmessagingSourceS3File               string
	_socialmessagingSourceS3PresignedUrl       string
	_socialmessagingTagKeys                    []string
	_socialmessagingTags                       string
	_socialmessagingTemplateCategory           string
	_socialmessagingTemplateComponents         string
	_socialmessagingTemplateDefinition         string
	_socialmessagingTemplateName               string
)

// This is only used through the Amazon Web Services console during sign-up to
// associate your WhatsApp Business Account to your Amazon Web Services account.
func socialmessaging_AssociateWhatsAppBusinessAccount(cfg aws.Config, client *socialmessaging.Client) {
	input := &socialmessaging.AssociateWhatsAppBusinessAccountInput{}

	if len(_socialmessagingSetupFinalization) > 0 {
		if err := assignInputField(input, "SetupFinalization", _socialmessagingSetupFinalization); err != nil {
			log.Errorf("invalid --setup-finalization: %s", err.Error())
			return
		}
	}
	if len(_socialmessagingSignupCallback) > 0 {
		if err := assignInputField(input, "SignupCallback", _socialmessagingSignupCallback); err != nil {
			log.Errorf("invalid --signup-callback: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateWhatsAppBusinessAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new WhatsApp message template from a custom definition.
// Amazon Web Services End User Messaging Social does not store any WhatsApp
// message template content.
func socialmessaging_CreateWhatsAppMessageTemplate(cfg aws.Config, client *socialmessaging.Client) {
	input := &socialmessaging.CreateWhatsAppMessageTemplateInput{
		// Id: *string, // Required
		// TemplateDefinition: []byte, // Required
	}

	if len(_socialmessagingId) > 0 {
		input.Id = aws.String(_socialmessagingId)
	}
	if len(_socialmessagingTemplateDefinition) > 0 {
		if err := assignInputField(input, "TemplateDefinition", _socialmessagingTemplateDefinition); err != nil {
			log.Errorf("invalid --template-definition: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWhatsAppMessageTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new WhatsApp message template using a template from Meta's template
// library.
func socialmessaging_CreateWhatsAppMessageTemplateFromLibrary(cfg aws.Config, client *socialmessaging.Client) {
	input := &socialmessaging.CreateWhatsAppMessageTemplateFromLibraryInput{
		// Id: *string, // Required
		// MetaLibraryTemplate: *types.MetaLibraryTemplate, // Required
	}

	if len(_socialmessagingId) > 0 {
		input.Id = aws.String(_socialmessagingId)
	}
	if len(_socialmessagingMetaLibraryTemplate) > 0 {
		if err := assignInputField(input, "MetaLibraryTemplate", _socialmessagingMetaLibraryTemplate); err != nil {
			log.Errorf("invalid --meta-library-template: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWhatsAppMessageTemplateFromLibrary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Uploads media for use in a WhatsApp message template.
func socialmessaging_CreateWhatsAppMessageTemplateMedia(cfg aws.Config, client *socialmessaging.Client) {
	input := &socialmessaging.CreateWhatsAppMessageTemplateMediaInput{
		// Id: *string, // Required
	}

	if len(_socialmessagingId) > 0 {
		input.Id = aws.String(_socialmessagingId)
	}
	if len(_socialmessagingSourceS3File) > 0 {
		if err := assignInputField(input, "SourceS3File", _socialmessagingSourceS3File); err != nil {
			log.Errorf("invalid --source-s3-file: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWhatsAppMessageTemplateMedia(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a media object from the WhatsApp service. If the object is still in an
// Amazon S3 bucket you should delete it from there too.
func socialmessaging_DeleteWhatsAppMessageMedia(cfg aws.Config, client *socialmessaging.Client) {
	input := &socialmessaging.DeleteWhatsAppMessageMediaInput{
		// MediaId: *string, // Required
		// OriginationPhoneNumberId: *string, // Required
	}

	if len(_socialmessagingMediaId) > 0 {
		input.MediaId = aws.String(_socialmessagingMediaId)
	}
	if len(_socialmessagingOriginationPhoneNumberId) > 0 {
		input.OriginationPhoneNumberId = aws.String(_socialmessagingOriginationPhoneNumberId)
	}

	if resp, err := client.DeleteWhatsAppMessageMedia(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a WhatsApp message template.
func socialmessaging_DeleteWhatsAppMessageTemplate(cfg aws.Config, client *socialmessaging.Client) {
	input := &socialmessaging.DeleteWhatsAppMessageTemplateInput{
		// Id: *string, // Required
		// TemplateName: *string, // Required
	}

	if len(_socialmessagingId) > 0 {
		input.Id = aws.String(_socialmessagingId)
	}
	if len(_socialmessagingTemplateName) > 0 {
		input.TemplateName = aws.String(_socialmessagingTemplateName)
	}
	if len(_socialmessagingDeleteAllLanguages) > 0 {
		if err := assignInputField(input, "DeleteAllLanguages", _socialmessagingDeleteAllLanguages); err != nil {
			log.Errorf("invalid --delete-all-languages: %s", err.Error())
			return
		}
	}
	if len(_socialmessagingMetaTemplateId) > 0 {
		input.MetaTemplateId = aws.String(_socialmessagingMetaTemplateId)
	}

	if resp, err := client.DeleteWhatsAppMessageTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociate a WhatsApp Business Account (WABA) from your Amazon Web Services
// account.
func socialmessaging_DisassociateWhatsAppBusinessAccount(cfg aws.Config, client *socialmessaging.Client) {
	input := &socialmessaging.DisassociateWhatsAppBusinessAccountInput{
		// Id: *string, // Required
	}

	if len(_socialmessagingId) > 0 {
		input.Id = aws.String(_socialmessagingId)
	}

	if resp, err := client.DisassociateWhatsAppBusinessAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the details of your linked WhatsApp Business Account.
func socialmessaging_GetLinkedWhatsAppBusinessAccount(cfg aws.Config, client *socialmessaging.Client) {
	input := &socialmessaging.GetLinkedWhatsAppBusinessAccountInput{
		// Id: *string, // Required
	}

	if len(_socialmessagingId) > 0 {
		input.Id = aws.String(_socialmessagingId)
	}

	if resp, err := client.GetLinkedWhatsAppBusinessAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve the WABA account id and phone number details of a WhatsApp business
// account phone number.
func socialmessaging_GetLinkedWhatsAppBusinessAccountPhoneNumber(cfg aws.Config, client *socialmessaging.Client) {
	input := &socialmessaging.GetLinkedWhatsAppBusinessAccountPhoneNumberInput{
		// Id: *string, // Required
	}

	if len(_socialmessagingId) > 0 {
		input.Id = aws.String(_socialmessagingId)
	}

	if resp, err := client.GetLinkedWhatsAppBusinessAccountPhoneNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get a media file from the WhatsApp service. On successful completion the media
// file is retrieved from Meta and stored in the specified Amazon S3 bucket. Use
// either destinationS3File or destinationS3PresignedUrl for the destination. If
// both are used then an InvalidParameterException is returned.
func socialmessaging_GetWhatsAppMessageMedia(cfg aws.Config, client *socialmessaging.Client) {
	input := &socialmessaging.GetWhatsAppMessageMediaInput{
		// MediaId: *string, // Required
		// OriginationPhoneNumberId: *string, // Required
	}

	if len(_socialmessagingMediaId) > 0 {
		input.MediaId = aws.String(_socialmessagingMediaId)
	}
	if len(_socialmessagingOriginationPhoneNumberId) > 0 {
		input.OriginationPhoneNumberId = aws.String(_socialmessagingOriginationPhoneNumberId)
	}
	if len(_socialmessagingDestinationS3File) > 0 {
		if err := assignInputField(input, "DestinationS3File", _socialmessagingDestinationS3File); err != nil {
			log.Errorf("invalid --destination-s3-file: %s", err.Error())
			return
		}
	}
	if len(_socialmessagingDestinationS3PresignedUrl) > 0 {
		if err := assignInputField(input, "DestinationS3PresignedUrl", _socialmessagingDestinationS3PresignedUrl); err != nil {
			log.Errorf("invalid --destination-s3-presigned-url: %s", err.Error())
			return
		}
	}
	if len(_socialmessagingMetadataOnly) > 0 {
		if err := assignInputField(input, "MetadataOnly", _socialmessagingMetadataOnly); err != nil {
			log.Errorf("invalid --metadata-only: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetWhatsAppMessageMedia(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a specific WhatsApp message template.
func socialmessaging_GetWhatsAppMessageTemplate(cfg aws.Config, client *socialmessaging.Client) {
	input := &socialmessaging.GetWhatsAppMessageTemplateInput{
		// Id: *string, // Required
		// MetaTemplateId: *string, // Required
	}

	if len(_socialmessagingId) > 0 {
		input.Id = aws.String(_socialmessagingId)
	}
	if len(_socialmessagingMetaTemplateId) > 0 {
		input.MetaTemplateId = aws.String(_socialmessagingMetaTemplateId)
	}

	if resp, err := client.GetWhatsAppMessageTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List all WhatsApp Business Accounts linked to your Amazon Web Services account.
func socialmessaging_ListLinkedWhatsAppBusinessAccounts(cfg aws.Config, client *socialmessaging.Client) {
	input := &socialmessaging.ListLinkedWhatsAppBusinessAccountsInput{}

	if len(_socialmessagingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _socialmessagingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_socialmessagingNextToken) > 0 {
		input.NextToken = aws.String(_socialmessagingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLinkedWhatsAppBusinessAccounts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*socialmessaging.ListLinkedWhatsAppBusinessAccountsOutput
	p := socialmessaging.NewListLinkedWhatsAppBusinessAccountsPaginator(client, input)
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

// List all tags associated with a resource, such as a phone number or WABA.
func socialmessaging_ListTagsForResource(cfg aws.Config, client *socialmessaging.Client) {
	input := &socialmessaging.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_socialmessagingResourceArn) > 0 {
		input.ResourceArn = aws.String(_socialmessagingResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists WhatsApp message templates for a specific WhatsApp Business Account.
func socialmessaging_ListWhatsAppMessageTemplates(cfg aws.Config, client *socialmessaging.Client) {
	input := &socialmessaging.ListWhatsAppMessageTemplatesInput{
		// Id: *string, // Required
	}

	if len(_socialmessagingId) > 0 {
		input.Id = aws.String(_socialmessagingId)
	}
	if len(_socialmessagingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _socialmessagingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_socialmessagingNextToken) > 0 {
		input.NextToken = aws.String(_socialmessagingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWhatsAppMessageTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*socialmessaging.ListWhatsAppMessageTemplatesOutput
	p := socialmessaging.NewListWhatsAppMessageTemplatesPaginator(client, input)
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

// Lists templates available in Meta's template library for WhatsApp messaging.
func socialmessaging_ListWhatsAppTemplateLibrary(cfg aws.Config, client *socialmessaging.Client) {
	input := &socialmessaging.ListWhatsAppTemplateLibraryInput{
		// Id: *string, // Required
	}

	if len(_socialmessagingId) > 0 {
		input.Id = aws.String(_socialmessagingId)
	}
	if len(_socialmessagingFilters) > 0 {
		if err := assignInputField(input, "Filters", _socialmessagingFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_socialmessagingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _socialmessagingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_socialmessagingNextToken) > 0 {
		input.NextToken = aws.String(_socialmessagingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWhatsAppTemplateLibrary(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*socialmessaging.ListWhatsAppTemplateLibraryOutput
	p := socialmessaging.NewListWhatsAppTemplateLibraryPaginator(client, input)
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

// Upload a media file to the WhatsApp service. Only the specified
// originationPhoneNumberId has the permissions to send the media file when using [SendWhatsAppMessage]
// . You must use either sourceS3File or sourceS3PresignedUrl for the source. If
// both or neither are specified then an InvalidParameterException is returned.
//
// [SendWhatsAppMessage]: https://docs.aws.amazon.com/social-messaging/latest/APIReference/API_SendWhatsAppMessage.html
func socialmessaging_PostWhatsAppMessageMedia(cfg aws.Config, client *socialmessaging.Client) {
	input := &socialmessaging.PostWhatsAppMessageMediaInput{
		// OriginationPhoneNumberId: *string, // Required
	}

	if len(_socialmessagingOriginationPhoneNumberId) > 0 {
		input.OriginationPhoneNumberId = aws.String(_socialmessagingOriginationPhoneNumberId)
	}
	if len(_socialmessagingSourceS3File) > 0 {
		if err := assignInputField(input, "SourceS3File", _socialmessagingSourceS3File); err != nil {
			log.Errorf("invalid --source-s3-file: %s", err.Error())
			return
		}
	}
	if len(_socialmessagingSourceS3PresignedUrl) > 0 {
		if err := assignInputField(input, "SourceS3PresignedUrl", _socialmessagingSourceS3PresignedUrl); err != nil {
			log.Errorf("invalid --source-s3-presigned-url: %s", err.Error())
			return
		}
	}

	if resp, err := client.PostWhatsAppMessageMedia(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Add an event destination to log event data from WhatsApp for a WhatsApp
// Business Account (WABA). A WABA can only have one event destination at a time.
// All resources associated with the WABA use the same event destination.
func socialmessaging_PutWhatsAppBusinessAccountEventDestinations(cfg aws.Config, client *socialmessaging.Client) {
	input := &socialmessaging.PutWhatsAppBusinessAccountEventDestinationsInput{
		// EventDestinations: []types.WhatsAppBusinessAccountEventDestination, // Required
		// Id: *string, // Required
	}

	if len(_socialmessagingEventDestinations) > 0 {
		if err := assignInputField(input, "EventDestinations", _socialmessagingEventDestinations); err != nil {
			log.Errorf("invalid --event-destinations: %s", err.Error())
			return
		}
	}
	if len(_socialmessagingId) > 0 {
		input.Id = aws.String(_socialmessagingId)
	}

	if resp, err := client.PutWhatsAppBusinessAccountEventDestinations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Send a WhatsApp message. For examples of sending a message using the Amazon Web
// Services CLI, see [Sending messages]in the Amazon Web Services End User Messaging Social User
// Guide .
//
// [Sending messages]: https://docs.aws.amazon.com/social-messaging/latest/userguide/send-message.html
func socialmessaging_SendWhatsAppMessage(cfg aws.Config, client *socialmessaging.Client) {
	input := &socialmessaging.SendWhatsAppMessageInput{
		// Message: []byte, // Required
		// MetaApiVersion: *string, // Required
		// OriginationPhoneNumberId: *string, // Required
	}

	if len(_socialmessagingMessage) > 0 {
		if err := assignInputField(input, "Message", _socialmessagingMessage); err != nil {
			log.Errorf("invalid --message: %s", err.Error())
			return
		}
	}
	if len(_socialmessagingMetaApiVersion) > 0 {
		input.MetaApiVersion = aws.String(_socialmessagingMetaApiVersion)
	}
	if len(_socialmessagingOriginationPhoneNumberId) > 0 {
		input.OriginationPhoneNumberId = aws.String(_socialmessagingOriginationPhoneNumberId)
	}

	if resp, err := client.SendWhatsAppMessage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or overwrites only the specified tags for the specified resource. When you
// specify an existing tag key, the value is overwritten with the new value.
func socialmessaging_TagResource(cfg aws.Config, client *socialmessaging.Client) {
	input := &socialmessaging.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_socialmessagingResourceArn) > 0 {
		input.ResourceArn = aws.String(_socialmessagingResourceArn)
	}
	if len(_socialmessagingTags) > 0 {
		if err := assignInputField(input, "Tags", _socialmessagingTags); err != nil {
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

// Removes the specified tags from a resource.
func socialmessaging_UntagResource(cfg aws.Config, client *socialmessaging.Client) {
	input := &socialmessaging.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_socialmessagingResourceArn) > 0 {
		input.ResourceArn = aws.String(_socialmessagingResourceArn)
	}
	if len(_socialmessagingTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _socialmessagingTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing WhatsApp message template.
func socialmessaging_UpdateWhatsAppMessageTemplate(cfg aws.Config, client *socialmessaging.Client) {
	input := &socialmessaging.UpdateWhatsAppMessageTemplateInput{
		// Id: *string, // Required
		// MetaTemplateId: *string, // Required
	}

	if len(_socialmessagingId) > 0 {
		input.Id = aws.String(_socialmessagingId)
	}
	if len(_socialmessagingMetaTemplateId) > 0 {
		input.MetaTemplateId = aws.String(_socialmessagingMetaTemplateId)
	}
	if len(_socialmessagingCtaUrlLinkTrackingOptedOut) > 0 {
		if err := assignInputField(input, "CtaUrlLinkTrackingOptedOut", _socialmessagingCtaUrlLinkTrackingOptedOut); err != nil {
			log.Errorf("invalid --cta-url-link-tracking-opted-out: %s", err.Error())
			return
		}
	}
	if len(_socialmessagingParameterFormat) > 0 {
		input.ParameterFormat = aws.String(_socialmessagingParameterFormat)
	}
	if len(_socialmessagingTemplateCategory) > 0 {
		input.TemplateCategory = aws.String(_socialmessagingTemplateCategory)
	}
	if len(_socialmessagingTemplateComponents) > 0 {
		if err := assignInputField(input, "TemplateComponents", _socialmessagingTemplateComponents); err != nil {
			log.Errorf("invalid --template-components: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateWhatsAppMessageTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_socialmessagingCmd)
	_socialmessagingCmd.Flags().SortFlags = false

	_socialmessagingCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_socialmessagingCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_socialmessagingCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_socialmessagingCmd.Flags().StringVarP(&_socialmessagingCtaUrlLinkTrackingOptedOut, "cta-url-link-tracking-opted-out", "", "", "Cta URL Link Tracking Opted Out")
	_socialmessagingCmd.Flags().StringVarP(&_socialmessagingDeleteAllLanguages, "delete-all-languages", "", "", "Delete All Languages")
	_socialmessagingCmd.Flags().StringVarP(&_socialmessagingDestinationS3File, "destination-s3-file", "", "", "Destination S3 File")
	_socialmessagingCmd.Flags().StringVarP(&_socialmessagingDestinationS3PresignedUrl, "destination-s3-presigned-url", "", "", "Destination S3 Presigned URL")
	_socialmessagingCmd.Flags().StringVarP(&_socialmessagingEventDestinations, "event-destinations", "", "", "Event Destinations")
	_socialmessagingCmd.Flags().StringVarP(&_socialmessagingFilters, "filters", "", "", "Filters")
	_socialmessagingCmd.Flags().StringVarP(&_socialmessagingId, "id", "", "", "ID")
	_socialmessagingCmd.Flags().StringVarP(&_socialmessagingMaxResults, "max-results", "", "", "Max Results")
	_socialmessagingCmd.Flags().StringVarP(&_socialmessagingMediaId, "media-id", "", "", "Media ID")
	_socialmessagingCmd.Flags().StringVarP(&_socialmessagingMessage, "message", "", "", "Message")
	_socialmessagingCmd.Flags().StringVarP(&_socialmessagingMetaApiVersion, "meta-api-version", "", "", "Meta API Version")
	_socialmessagingCmd.Flags().StringVarP(&_socialmessagingMetaLibraryTemplate, "meta-library-template", "", "", "Meta Library Template")
	_socialmessagingCmd.Flags().StringVarP(&_socialmessagingMetaTemplateId, "meta-template-id", "", "", "Meta Template ID")
	_socialmessagingCmd.Flags().StringVarP(&_socialmessagingMetadataOnly, "metadata-only", "", "", "Metadata Only")
	_socialmessagingCmd.Flags().StringVarP(&_socialmessagingNextToken, "next-token", "", "", "Next Token")
	_socialmessagingCmd.Flags().StringVarP(&_socialmessagingOriginationPhoneNumberId, "origination-phone-number-id", "", "", "Origination Phone Number ID")
	_socialmessagingCmd.Flags().StringVarP(&_socialmessagingParameterFormat, "parameter-format", "", "", "Parameter Format")
	_socialmessagingCmd.Flags().StringVarP(&_socialmessagingResourceArn, "resource-arn", "", "", "Resource ARN")
	_socialmessagingCmd.Flags().StringVarP(&_socialmessagingSetupFinalization, "setup-finalization", "", "", "Setup Finalization")
	_socialmessagingCmd.Flags().StringVarP(&_socialmessagingSignupCallback, "signup-callback", "", "", "Signup Callback")
	_socialmessagingCmd.Flags().StringVarP(&_socialmessagingSourceS3File, "source-s3-file", "", "", "Source S3 File")
	_socialmessagingCmd.Flags().StringVarP(&_socialmessagingSourceS3PresignedUrl, "source-s3-presigned-url", "", "", "Source S3 Presigned URL")
	_socialmessagingCmd.Flags().StringSliceVarP(&_socialmessagingTagKeys, "tag-keys", "", nil, "Tag Keys")
	_socialmessagingCmd.Flags().StringVarP(&_socialmessagingTags, "tags", "", "", "Tags")
	_socialmessagingCmd.Flags().StringVarP(&_socialmessagingTemplateCategory, "template-category", "", "", "Template Category")
	_socialmessagingCmd.Flags().StringVarP(&_socialmessagingTemplateComponents, "template-components", "", "", "Template Components")
	_socialmessagingCmd.Flags().StringVarP(&_socialmessagingTemplateDefinition, "template-definition", "", "", "Template Definition")
	_socialmessagingCmd.Flags().StringVarP(&_socialmessagingTemplateName, "template-name", "", "", "Template Name")

	_socialmessagingCmd.Flags().BoolVarP(&_socialmessagingAssociateWhatsAppBusinessAccount, "associate-whats-app-business-account", "", false, "Associate Whats App Business Account")
	_socialmessagingCmd.Flags().BoolVarP(&_socialmessagingCreateWhatsAppMessageTemplate, "create-whats-app-message-template", "", false, "Create Whats App Message Template")
	_socialmessagingCmd.Flags().BoolVarP(&_socialmessagingCreateWhatsAppMessageTemplateFromLibrary, "create-whats-app-message-template-from-library", "", false, "Create Whats App Message Template From Library")
	_socialmessagingCmd.Flags().BoolVarP(&_socialmessagingCreateWhatsAppMessageTemplateMedia, "create-whats-app-message-template-media", "", false, "Create Whats App Message Template Media")
	_socialmessagingCmd.Flags().BoolVarP(&_socialmessagingDeleteWhatsAppMessageMedia, "delete-whats-app-message-media", "", false, "Delete Whats App Message Media")
	_socialmessagingCmd.Flags().BoolVarP(&_socialmessagingDeleteWhatsAppMessageTemplate, "delete-whats-app-message-template", "", false, "Delete Whats App Message Template")
	_socialmessagingCmd.Flags().BoolVarP(&_socialmessagingDisassociateWhatsAppBusinessAccount, "disassociate-whats-app-business-account", "", false, "Disassociate Whats App Business Account")
	_socialmessagingCmd.Flags().BoolVarP(&_socialmessagingGetLinkedWhatsAppBusinessAccount, "get-linked-whats-app-business-account", "", false, "Get Linked Whats App Business Account")
	_socialmessagingCmd.Flags().BoolVarP(&_socialmessagingGetLinkedWhatsAppBusinessAccountPhoneNumber, "get-linked-whats-app-business-account-phone-number", "", false, "Get Linked Whats App Business Account Phone Number")
	_socialmessagingCmd.Flags().BoolVarP(&_socialmessagingGetWhatsAppMessageMedia, "get-whats-app-message-media", "", false, "Get Whats App Message Media")
	_socialmessagingCmd.Flags().BoolVarP(&_socialmessagingGetWhatsAppMessageTemplate, "get-whats-app-message-template", "", false, "Get Whats App Message Template")
	_socialmessagingCmd.Flags().BoolVarP(&_socialmessagingListLinkedWhatsAppBusinessAccounts, "list-linked-whats-app-business-accounts", "", false, "List Linked Whats App Business Accounts")
	_socialmessagingCmd.Flags().BoolVarP(&_socialmessagingListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_socialmessagingCmd.Flags().BoolVarP(&_socialmessagingListWhatsAppMessageTemplates, "list-whats-app-message-templates", "", false, "List Whats App Message Templates")
	_socialmessagingCmd.Flags().BoolVarP(&_socialmessagingListWhatsAppTemplateLibrary, "list-whats-app-template-library", "", false, "List Whats App Template Library")
	_socialmessagingCmd.Flags().BoolVarP(&_socialmessagingPostWhatsAppMessageMedia, "post-whats-app-message-media", "", false, "Post Whats App Message Media")
	_socialmessagingCmd.Flags().BoolVarP(&_socialmessagingPutWhatsAppBusinessAccountEventDestinations, "put-whats-app-business-account-event-destinations", "", false, "Put Whats App Business Account Event Destinations")
	_socialmessagingCmd.Flags().BoolVarP(&_socialmessagingSendWhatsAppMessage, "send-whats-app-message", "", false, "Send Whats App Message")
	_socialmessagingCmd.Flags().BoolVarP(&_socialmessagingTagResource, "tag-resource", "", false, "Tag Resource")
	_socialmessagingCmd.Flags().BoolVarP(&_socialmessagingUntagResource, "untag-resource", "", false, "Untag Resource")
	_socialmessagingCmd.Flags().BoolVarP(&_socialmessagingUpdateWhatsAppMessageTemplate, "update-whats-app-message-template", "", false, "Update Whats App Message Template")

}
