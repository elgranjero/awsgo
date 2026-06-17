package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/account"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// accountCmd represents the account command
var _accountCmd = &cobra.Command{
	Use:   "account",
	Short: "AWS account CLI",
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
		client := account.NewFromConfig(cfg)
		if _accountAcceptPrimaryEmailUpdate {
			account_AcceptPrimaryEmailUpdate(cfg, client)
			return
		}
		if _accountDeleteAlternateContact {
			account_DeleteAlternateContact(cfg, client)
			return
		}
		if _accountDisableRegion {
			account_DisableRegion(cfg, client)
			return
		}
		if _accountEnableRegion {
			account_EnableRegion(cfg, client)
			return
		}
		if _accountGetAccountInformation {
			account_GetAccountInformation(cfg, client)
			return
		}
		if _accountGetAlternateContact {
			account_GetAlternateContact(cfg, client)
			return
		}
		if _accountGetContactInformation {
			account_GetContactInformation(cfg, client)
			return
		}
		if _accountGetGovCloudAccountInformation {
			account_GetGovCloudAccountInformation(cfg, client)
			return
		}
		if _accountGetPrimaryEmail {
			account_GetPrimaryEmail(cfg, client)
			return
		}
		if _accountGetRegionOptStatus {
			account_GetRegionOptStatus(cfg, client)
			return
		}
		if _accountListRegions {
			account_ListRegions(cfg, client)
			return
		}
		if _accountPutAccountName {
			account_PutAccountName(cfg, client)
			return
		}
		if _accountPutAlternateContact {
			account_PutAlternateContact(cfg, client)
			return
		}
		if _accountPutContactInformation {
			account_PutContactInformation(cfg, client)
			return
		}
		if _accountStartPrimaryEmailUpdate {
			account_StartPrimaryEmailUpdate(cfg, client)
			return
		}

	},
}

var (
	_accountAcceptPrimaryEmailUpdate      bool
	_accountDeleteAlternateContact        bool
	_accountDisableRegion                 bool
	_accountEnableRegion                  bool
	_accountGetAccountInformation         bool
	_accountGetAlternateContact           bool
	_accountGetContactInformation         bool
	_accountGetGovCloudAccountInformation bool
	_accountGetPrimaryEmail               bool
	_accountGetRegionOptStatus            bool
	_accountListRegions                   bool
	_accountPutAccountName                bool
	_accountPutAlternateContact           bool
	_accountPutContactInformation         bool
	_accountStartPrimaryEmailUpdate       bool

	_accountAccountId               string
	_accountAccountName             string
	_accountAlternateContactType    string
	_accountContactInformation      string
	_accountEmailAddress            string
	_accountMaxResults              string
	_accountName                    string
	_accountNextToken               string
	_accountOtp                     string
	_accountPhoneNumber             string
	_accountPrimaryEmail            string
	_accountRegionName              string
	_accountRegionOptStatusContains string
	_accountStandardAccountId       string
	_accountTitle                   string
)

// Accepts the request that originated from StartPrimaryEmailUpdate to update the primary email address
// (also known as the root user email address) for the specified account.
func account_AcceptPrimaryEmailUpdate(cfg aws.Config, client *account.Client) {
	input := &account.AcceptPrimaryEmailUpdateInput{
		// AccountId: *string, // Required
		// Otp: *string, // Required
		// PrimaryEmail: *string, // Required
	}

	if len(_accountAccountId) > 0 {
		input.AccountId = aws.String(_accountAccountId)
	}
	if len(_accountOtp) > 0 {
		input.Otp = aws.String(_accountOtp)
	}
	if len(_accountPrimaryEmail) > 0 {
		input.PrimaryEmail = aws.String(_accountPrimaryEmail)
	}

	if resp, err := client.AcceptPrimaryEmailUpdate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified alternate contact from an Amazon Web Services account.
// For complete details about how to use the alternate contact operations, see [Update the alternate contacts for your Amazon Web Services account].
//
// Before you can update the alternate contact information for an Amazon Web
// Services account that is managed by Organizations, you must first enable
// integration between Amazon Web Services Account Management and Organizations.
// For more information, see [Enable trusted access for Amazon Web Services Account Management].
//
// [Enable trusted access for Amazon Web Services Account Management]: https://docs.aws.amazon.com/accounts/latest/reference/using-orgs-trusted-access.html
// [Update the alternate contacts for your Amazon Web Services account]: https://docs.aws.amazon.com/accounts/latest/reference/manage-acct-update-contact-alternate.html
func account_DeleteAlternateContact(cfg aws.Config, client *account.Client) {
	input := &account.DeleteAlternateContactInput{
		// AlternateContactType: types.AlternateContactType, // Required
	}

	if len(_accountAlternateContactType) > 0 {
		if err := assignInputField(input, "AlternateContactType", _accountAlternateContactType); err != nil {
			log.Errorf("invalid --alternate-contact-type: %s", err.Error())
			return
		}
	}
	if len(_accountAccountId) > 0 {
		input.AccountId = aws.String(_accountAccountId)
	}

	if resp, err := client.DeleteAlternateContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables (opts-out) a particular Region for an account.
// The act of disabling a Region will remove all IAM access to any resources that
// reside in that Region.
func account_DisableRegion(cfg aws.Config, client *account.Client) {
	input := &account.DisableRegionInput{
		// RegionName: *string, // Required
	}

	if len(_accountRegionName) > 0 {
		input.RegionName = aws.String(_accountRegionName)
	}
	if len(_accountAccountId) > 0 {
		input.AccountId = aws.String(_accountAccountId)
	}

	if resp, err := client.DisableRegion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables (opts-in) a particular Region for an account.
func account_EnableRegion(cfg aws.Config, client *account.Client) {
	input := &account.EnableRegionInput{
		// RegionName: *string, // Required
	}

	if len(_accountRegionName) > 0 {
		input.RegionName = aws.String(_accountRegionName)
	}
	if len(_accountAccountId) > 0 {
		input.AccountId = aws.String(_accountAccountId)
	}

	if resp, err := client.EnableRegion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the specified account including its account name,
// account ID, and account creation date and time. To use this API, an IAM user or
// role must have the account:GetAccountInformation IAM permission.
func account_GetAccountInformation(cfg aws.Config, client *account.Client) {
	input := &account.GetAccountInformationInput{}

	if len(_accountAccountId) > 0 {
		input.AccountId = aws.String(_accountAccountId)
	}

	if resp, err := client.GetAccountInformation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the specified alternate contact attached to an Amazon Web Services
// account.
//
// For complete details about how to use the alternate contact operations, see [Update the alternate contacts for your Amazon Web Services account].
//
// Before you can update the alternate contact information for an Amazon Web
// Services account that is managed by Organizations, you must first enable
// integration between Amazon Web Services Account Management and Organizations.
// For more information, see [Enable trusted access for Amazon Web Services Account Management].
//
// [Enable trusted access for Amazon Web Services Account Management]: https://docs.aws.amazon.com/accounts/latest/reference/using-orgs-trusted-access.html
// [Update the alternate contacts for your Amazon Web Services account]: https://docs.aws.amazon.com/accounts/latest/reference/manage-acct-update-contact-alternate.html
func account_GetAlternateContact(cfg aws.Config, client *account.Client) {
	input := &account.GetAlternateContactInput{
		// AlternateContactType: types.AlternateContactType, // Required
	}

	if len(_accountAlternateContactType) > 0 {
		if err := assignInputField(input, "AlternateContactType", _accountAlternateContactType); err != nil {
			log.Errorf("invalid --alternate-contact-type: %s", err.Error())
			return
		}
	}
	if len(_accountAccountId) > 0 {
		input.AccountId = aws.String(_accountAccountId)
	}

	if resp, err := client.GetAlternateContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the primary contact information of an Amazon Web Services account.
// For complete details about how to use the primary contact operations, see [Update the primary contact for your Amazon Web Services account].
//
// [Update the primary contact for your Amazon Web Services account]: https://docs.aws.amazon.com/accounts/latest/reference/manage-acct-update-contact-primary.html
func account_GetContactInformation(cfg aws.Config, client *account.Client) {
	input := &account.GetContactInformationInput{}

	if len(_accountAccountId) > 0 {
		input.AccountId = aws.String(_accountAccountId)
	}

	if resp, err := client.GetContactInformation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the GovCloud account linked to the specified
// standard account (if it exists) including the GovCloud account ID and state. To
// use this API, an IAM user or role must have the
// account:GetGovCloudAccountInformation IAM permission.
func account_GetGovCloudAccountInformation(cfg aws.Config, client *account.Client) {
	input := &account.GetGovCloudAccountInformationInput{}

	if len(_accountStandardAccountId) > 0 {
		input.StandardAccountId = aws.String(_accountStandardAccountId)
	}

	if resp, err := client.GetGovCloudAccountInformation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the primary email address for the specified account.
func account_GetPrimaryEmail(cfg aws.Config, client *account.Client) {
	input := &account.GetPrimaryEmailInput{
		// AccountId: *string, // Required
	}

	if len(_accountAccountId) > 0 {
		input.AccountId = aws.String(_accountAccountId)
	}

	if resp, err := client.GetPrimaryEmail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the opt-in status of a particular Region.
func account_GetRegionOptStatus(cfg aws.Config, client *account.Client) {
	input := &account.GetRegionOptStatusInput{
		// RegionName: *string, // Required
	}

	if len(_accountRegionName) > 0 {
		input.RegionName = aws.String(_accountRegionName)
	}
	if len(_accountAccountId) > 0 {
		input.AccountId = aws.String(_accountAccountId)
	}

	if resp, err := client.GetRegionOptStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the Regions for a given account and their respective opt-in statuses.
// Optionally, this list can be filtered by the region-opt-status-contains
// parameter.
func account_ListRegions(cfg aws.Config, client *account.Client) {
	input := &account.ListRegionsInput{}

	if len(_accountAccountId) > 0 {
		input.AccountId = aws.String(_accountAccountId)
	}
	if len(_accountMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _accountMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_accountNextToken) > 0 {
		input.NextToken = aws.String(_accountNextToken)
	}
	if len(_accountRegionOptStatusContains) > 0 {
		if err := assignInputField(input, "RegionOptStatusContains", _accountRegionOptStatusContains); err != nil {
			log.Errorf("invalid --region-opt-status-contains: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListRegions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*account.ListRegionsOutput
	p := account.NewListRegionsPaginator(client, input)
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

// Updates the account name of the specified account. To use this API, IAM
// principals must have the account:PutAccountName IAM permission.
func account_PutAccountName(cfg aws.Config, client *account.Client) {
	input := &account.PutAccountNameInput{
		// AccountName: *string, // Required
	}

	if len(_accountAccountName) > 0 {
		input.AccountName = aws.String(_accountAccountName)
	}
	if len(_accountAccountId) > 0 {
		input.AccountId = aws.String(_accountAccountId)
	}

	if resp, err := client.PutAccountName(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the specified alternate contact attached to an Amazon Web Services
// account.
//
// For complete details about how to use the alternate contact operations, see [Update the alternate contacts for your Amazon Web Services account].
//
// Before you can update the alternate contact information for an Amazon Web
// Services account that is managed by Organizations, you must first enable
// integration between Amazon Web Services Account Management and Organizations.
// For more information, see [Enable trusted access for Amazon Web Services Account Management].
//
// [Enable trusted access for Amazon Web Services Account Management]: https://docs.aws.amazon.com/accounts/latest/reference/using-orgs-trusted-access.html
// [Update the alternate contacts for your Amazon Web Services account]: https://docs.aws.amazon.com/accounts/latest/reference/manage-acct-update-contact-alternate.html
func account_PutAlternateContact(cfg aws.Config, client *account.Client) {
	input := &account.PutAlternateContactInput{
		// AlternateContactType: types.AlternateContactType, // Required
		// EmailAddress: *string, // Required
		// Name: *string, // Required
		// PhoneNumber: *string, // Required
		// Title: *string, // Required
	}

	if len(_accountAlternateContactType) > 0 {
		if err := assignInputField(input, "AlternateContactType", _accountAlternateContactType); err != nil {
			log.Errorf("invalid --alternate-contact-type: %s", err.Error())
			return
		}
	}
	if len(_accountEmailAddress) > 0 {
		input.EmailAddress = aws.String(_accountEmailAddress)
	}
	if len(_accountName) > 0 {
		input.Name = aws.String(_accountName)
	}
	if len(_accountPhoneNumber) > 0 {
		input.PhoneNumber = aws.String(_accountPhoneNumber)
	}
	if len(_accountTitle) > 0 {
		input.Title = aws.String(_accountTitle)
	}
	if len(_accountAccountId) > 0 {
		input.AccountId = aws.String(_accountAccountId)
	}

	if resp, err := client.PutAlternateContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the primary contact information of an Amazon Web Services account.
// For complete details about how to use the primary contact operations, see [Update the primary contact for your Amazon Web Services account].
//
// [Update the primary contact for your Amazon Web Services account]: https://docs.aws.amazon.com/accounts/latest/reference/manage-acct-update-contact-primary.html
func account_PutContactInformation(cfg aws.Config, client *account.Client) {
	input := &account.PutContactInformationInput{
		// ContactInformation: *types.ContactInformation, // Required
	}

	if len(_accountContactInformation) > 0 {
		if err := assignInputField(input, "ContactInformation", _accountContactInformation); err != nil {
			log.Errorf("invalid --contact-information: %s", err.Error())
			return
		}
	}
	if len(_accountAccountId) > 0 {
		input.AccountId = aws.String(_accountAccountId)
	}

	if resp, err := client.PutContactInformation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the process to update the primary email address for the specified
// account.
func account_StartPrimaryEmailUpdate(cfg aws.Config, client *account.Client) {
	input := &account.StartPrimaryEmailUpdateInput{
		// AccountId: *string, // Required
		// PrimaryEmail: *string, // Required
	}

	if len(_accountAccountId) > 0 {
		input.AccountId = aws.String(_accountAccountId)
	}
	if len(_accountPrimaryEmail) > 0 {
		input.PrimaryEmail = aws.String(_accountPrimaryEmail)
	}

	if resp, err := client.StartPrimaryEmailUpdate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_accountCmd)
	_accountCmd.Flags().SortFlags = false

	_accountCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_accountCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_accountCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_accountCmd.Flags().StringVarP(&_accountAccountId, "account-id", "", "", "Account ID")
	_accountCmd.Flags().StringVarP(&_accountAccountName, "account-name", "", "", "Account Name")
	_accountCmd.Flags().StringVarP(&_accountAlternateContactType, "alternate-contact-type", "", "", "Alternate Contact Type")
	_accountCmd.Flags().StringVarP(&_accountContactInformation, "contact-information", "", "", "Contact Information")
	_accountCmd.Flags().StringVarP(&_accountEmailAddress, "email-address", "", "", "Email Address")
	_accountCmd.Flags().StringVarP(&_accountMaxResults, "max-results", "", "", "Max Results")
	_accountCmd.Flags().StringVarP(&_accountName, "name", "", "", "Name")
	_accountCmd.Flags().StringVarP(&_accountNextToken, "next-token", "", "", "Next Token")
	_accountCmd.Flags().StringVarP(&_accountOtp, "otp", "", "", "Otp")
	_accountCmd.Flags().StringVarP(&_accountPhoneNumber, "phone-number", "", "", "Phone Number")
	_accountCmd.Flags().StringVarP(&_accountPrimaryEmail, "primary-email", "", "", "Primary Email")
	_accountCmd.Flags().StringVarP(&_accountRegionName, "region-name", "", "", "Region Name")
	_accountCmd.Flags().StringVarP(&_accountRegionOptStatusContains, "region-opt-status-contains", "", "", "Region Opt Status Contains")
	_accountCmd.Flags().StringVarP(&_accountStandardAccountId, "standard-account-id", "", "", "Standard Account ID")
	_accountCmd.Flags().StringVarP(&_accountTitle, "title", "", "", "Title")

	_accountCmd.Flags().BoolVarP(&_accountAcceptPrimaryEmailUpdate, "accept-primary-email-update", "", false, "Accept Primary Email Update")
	_accountCmd.Flags().BoolVarP(&_accountDeleteAlternateContact, "delete-alternate-contact", "", false, "Delete Alternate Contact")
	_accountCmd.Flags().BoolVarP(&_accountDisableRegion, "disable-region", "", false, "Disable Region")
	_accountCmd.Flags().BoolVarP(&_accountEnableRegion, "enable-region", "", false, "Enable Region")
	_accountCmd.Flags().BoolVarP(&_accountGetAccountInformation, "get-account-information", "", false, "Get Account Information")
	_accountCmd.Flags().BoolVarP(&_accountGetAlternateContact, "get-alternate-contact", "", false, "Get Alternate Contact")
	_accountCmd.Flags().BoolVarP(&_accountGetContactInformation, "get-contact-information", "", false, "Get Contact Information")
	_accountCmd.Flags().BoolVarP(&_accountGetGovCloudAccountInformation, "get-gov-cloud-account-information", "", false, "Get Gov Cloud Account Information")
	_accountCmd.Flags().BoolVarP(&_accountGetPrimaryEmail, "get-primary-email", "", false, "Get Primary Email")
	_accountCmd.Flags().BoolVarP(&_accountGetRegionOptStatus, "get-region-opt-status", "", false, "Get Region Opt Status")
	_accountCmd.Flags().BoolVarP(&_accountListRegions, "list-regions", "", false, "List Regions")
	_accountCmd.Flags().BoolVarP(&_accountPutAccountName, "put-account-name", "", false, "Put Account Name")
	_accountCmd.Flags().BoolVarP(&_accountPutAlternateContact, "put-alternate-contact", "", false, "Put Alternate Contact")
	_accountCmd.Flags().BoolVarP(&_accountPutContactInformation, "put-contact-information", "", false, "Put Contact Information")
	_accountCmd.Flags().BoolVarP(&_accountStartPrimaryEmailUpdate, "start-primary-email-update", "", false, "Start Primary Email Update")

}
