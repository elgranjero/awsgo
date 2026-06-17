package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/notificationscontacts"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// notificationscontactsCmd represents the notificationscontacts command
var _notificationscontactsCmd = &cobra.Command{
	Use:   "notificationscontacts",
	Short: "AWS notificationscontacts CLI",
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
		client := notificationscontacts.NewFromConfig(cfg)
		if _notificationscontactsActivateEmailContact {
			notificationscontacts_ActivateEmailContact(cfg, client)
			return
		}
		if _notificationscontactsCreateEmailContact {
			notificationscontacts_CreateEmailContact(cfg, client)
			return
		}
		if _notificationscontactsDeleteEmailContact {
			notificationscontacts_DeleteEmailContact(cfg, client)
			return
		}
		if _notificationscontactsGetEmailContact {
			notificationscontacts_GetEmailContact(cfg, client)
			return
		}
		if _notificationscontactsListEmailContacts {
			notificationscontacts_ListEmailContacts(cfg, client)
			return
		}
		if _notificationscontactsListTagsForResource {
			notificationscontacts_ListTagsForResource(cfg, client)
			return
		}
		if _notificationscontactsSendActivationCode {
			notificationscontacts_SendActivationCode(cfg, client)
			return
		}
		if _notificationscontactsTagResource {
			notificationscontacts_TagResource(cfg, client)
			return
		}
		if _notificationscontactsUntagResource {
			notificationscontacts_UntagResource(cfg, client)
			return
		}

	},
}

var (
	_notificationscontactsActivateEmailContact bool
	_notificationscontactsCreateEmailContact   bool
	_notificationscontactsDeleteEmailContact   bool
	_notificationscontactsGetEmailContact      bool
	_notificationscontactsListEmailContacts    bool
	_notificationscontactsListTagsForResource  bool
	_notificationscontactsSendActivationCode   bool
	_notificationscontactsTagResource          bool
	_notificationscontactsUntagResource        bool

	_notificationscontactsArn          string
	_notificationscontactsCode         string
	_notificationscontactsEmailAddress string
	_notificationscontactsMaxResults   string
	_notificationscontactsName         string
	_notificationscontactsNextToken    string
	_notificationscontactsTagKeys      []string
	_notificationscontactsTags         string
)

// Activates an email contact using an activation code. This code is in the
// activation email sent to the email address associated with this email contact.
func notificationscontacts_ActivateEmailContact(cfg aws.Config, client *notificationscontacts.Client) {
	input := &notificationscontacts.ActivateEmailContactInput{
		// Arn: *string, // Required
		// Code: *string, // Required
	}

	if len(_notificationscontactsArn) > 0 {
		input.Arn = aws.String(_notificationscontactsArn)
	}
	if len(_notificationscontactsCode) > 0 {
		input.Code = aws.String(_notificationscontactsCode)
	}

	if resp, err := client.ActivateEmailContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an email contact for the provided email address.
func notificationscontacts_CreateEmailContact(cfg aws.Config, client *notificationscontacts.Client) {
	input := &notificationscontacts.CreateEmailContactInput{
		// EmailAddress: *string, // Required
		// Name: *string, // Required
	}

	if len(_notificationscontactsEmailAddress) > 0 {
		input.EmailAddress = aws.String(_notificationscontactsEmailAddress)
	}
	if len(_notificationscontactsName) > 0 {
		input.Name = aws.String(_notificationscontactsName)
	}
	if len(_notificationscontactsTags) > 0 {
		if err := assignInputField(input, "Tags", _notificationscontactsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEmailContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an email contact.
// Deleting an email contact removes it from all associated notification
// configurations.
func notificationscontacts_DeleteEmailContact(cfg aws.Config, client *notificationscontacts.Client) {
	input := &notificationscontacts.DeleteEmailContactInput{
		// Arn: *string, // Required
	}

	if len(_notificationscontactsArn) > 0 {
		input.Arn = aws.String(_notificationscontactsArn)
	}

	if resp, err := client.DeleteEmailContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an email contact.
func notificationscontacts_GetEmailContact(cfg aws.Config, client *notificationscontacts.Client) {
	input := &notificationscontacts.GetEmailContactInput{
		// Arn: *string, // Required
	}

	if len(_notificationscontactsArn) > 0 {
		input.Arn = aws.String(_notificationscontactsArn)
	}

	if resp, err := client.GetEmailContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all email contacts created under the Account.
func notificationscontacts_ListEmailContacts(cfg aws.Config, client *notificationscontacts.Client) {
	input := &notificationscontacts.ListEmailContactsInput{}

	if len(_notificationscontactsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _notificationscontactsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_notificationscontactsNextToken) > 0 {
		input.NextToken = aws.String(_notificationscontactsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEmailContacts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*notificationscontacts.ListEmailContactsOutput
	p := notificationscontacts.NewListEmailContactsPaginator(client, input)
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

// Lists all of the tags associated with the Amazon Resource Name (ARN) that you
// specify. The resource can be a user, server, or role.
func notificationscontacts_ListTagsForResource(cfg aws.Config, client *notificationscontacts.Client) {
	input := &notificationscontacts.ListTagsForResourceInput{
		// Arn: *string, // Required
	}

	if len(_notificationscontactsArn) > 0 {
		input.Arn = aws.String(_notificationscontactsArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends an activation email to the email address associated with the specified
// email contact.
//
// It might take a few minutes for the activation email to arrive. If it doesn't
// arrive, check in your spam folder or try sending another activation email.
func notificationscontacts_SendActivationCode(cfg aws.Config, client *notificationscontacts.Client) {
	input := &notificationscontacts.SendActivationCodeInput{
		// Arn: *string, // Required
	}

	if len(_notificationscontactsArn) > 0 {
		input.Arn = aws.String(_notificationscontactsArn)
	}

	if resp, err := client.SendActivationCode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches a key-value pair to a resource, as identified by its Amazon Resource
// Name (ARN). Taggable resources in AWS User Notifications Contacts include email
// contacts.
func notificationscontacts_TagResource(cfg aws.Config, client *notificationscontacts.Client) {
	input := &notificationscontacts.TagResourceInput{
		// Arn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_notificationscontactsArn) > 0 {
		input.Arn = aws.String(_notificationscontactsArn)
	}
	if len(_notificationscontactsTags) > 0 {
		if err := assignInputField(input, "Tags", _notificationscontactsTags); err != nil {
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

// Detaches a key-value pair from a resource, as identified by its Amazon Resource
// Name (ARN). Taggable resources in AWS User Notifications Contacts include email
// contacts..
func notificationscontacts_UntagResource(cfg aws.Config, client *notificationscontacts.Client) {
	input := &notificationscontacts.UntagResourceInput{
		// Arn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_notificationscontactsArn) > 0 {
		input.Arn = aws.String(_notificationscontactsArn)
	}
	if len(_notificationscontactsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _notificationscontactsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_notificationscontactsCmd)
	_notificationscontactsCmd.Flags().SortFlags = false

	_notificationscontactsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_notificationscontactsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_notificationscontactsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_notificationscontactsCmd.Flags().StringVarP(&_notificationscontactsArn, "arn", "", "", "ARN")
	_notificationscontactsCmd.Flags().StringVarP(&_notificationscontactsCode, "code", "", "", "Code")
	_notificationscontactsCmd.Flags().StringVarP(&_notificationscontactsEmailAddress, "email-address", "", "", "Email Address")
	_notificationscontactsCmd.Flags().StringVarP(&_notificationscontactsMaxResults, "max-results", "", "", "Max Results")
	_notificationscontactsCmd.Flags().StringVarP(&_notificationscontactsName, "name", "", "", "Name")
	_notificationscontactsCmd.Flags().StringVarP(&_notificationscontactsNextToken, "next-token", "", "", "Next Token")
	_notificationscontactsCmd.Flags().StringSliceVarP(&_notificationscontactsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_notificationscontactsCmd.Flags().StringVarP(&_notificationscontactsTags, "tags", "", "", "Tags")

	_notificationscontactsCmd.Flags().BoolVarP(&_notificationscontactsActivateEmailContact, "activate-email-contact", "", false, "Activate Email Contact")
	_notificationscontactsCmd.Flags().BoolVarP(&_notificationscontactsCreateEmailContact, "create-email-contact", "", false, "Create Email Contact")
	_notificationscontactsCmd.Flags().BoolVarP(&_notificationscontactsDeleteEmailContact, "delete-email-contact", "", false, "Delete Email Contact")
	_notificationscontactsCmd.Flags().BoolVarP(&_notificationscontactsGetEmailContact, "get-email-contact", "", false, "Get Email Contact")
	_notificationscontactsCmd.Flags().BoolVarP(&_notificationscontactsListEmailContacts, "list-email-contacts", "", false, "List Email Contacts")
	_notificationscontactsCmd.Flags().BoolVarP(&_notificationscontactsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_notificationscontactsCmd.Flags().BoolVarP(&_notificationscontactsSendActivationCode, "send-activation-code", "", false, "Send Activation Code")
	_notificationscontactsCmd.Flags().BoolVarP(&_notificationscontactsTagResource, "tag-resource", "", false, "Tag Resource")
	_notificationscontactsCmd.Flags().BoolVarP(&_notificationscontactsUntagResource, "untag-resource", "", false, "Untag Resource")

}
