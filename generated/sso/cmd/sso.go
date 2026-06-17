package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sso"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ssoCmd represents the sso command
var _ssoCmd = &cobra.Command{
	Use:   "sso",
	Short: "AWS sso CLI",
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
		client := sso.NewFromConfig(cfg)
		if _ssoGetRoleCredentials {
			sso_GetRoleCredentials(cfg, client)
			return
		}
		if _ssoListAccountRoles {
			sso_ListAccountRoles(cfg, client)
			return
		}
		if _ssoListAccounts {
			sso_ListAccounts(cfg, client)
			return
		}
		if _ssoLogout {
			sso_Logout(cfg, client)
			return
		}

	},
}

var (
	_ssoGetRoleCredentials bool
	_ssoListAccountRoles   bool
	_ssoListAccounts       bool
	_ssoLogout             bool

	_ssoAccessToken string
	_ssoAccountId   string
	_ssoMaxResults  string
	_ssoNextToken   string
	_ssoRoleName    string
)

// Returns the STS short-term credentials for a given role name that is assigned
// to the user.
func sso_GetRoleCredentials(cfg aws.Config, client *sso.Client) {
	input := &sso.GetRoleCredentialsInput{
		// AccessToken: *string, // Required
		// AccountId: *string, // Required
		// RoleName: *string, // Required
	}

	if len(_ssoAccessToken) > 0 {
		input.AccessToken = aws.String(_ssoAccessToken)
	}
	if len(_ssoAccountId) > 0 {
		input.AccountId = aws.String(_ssoAccountId)
	}
	if len(_ssoRoleName) > 0 {
		input.RoleName = aws.String(_ssoRoleName)
	}

	if resp, err := client.GetRoleCredentials(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all roles that are assigned to the user for a given AWS account.
func sso_ListAccountRoles(cfg aws.Config, client *sso.Client) {
	input := &sso.ListAccountRolesInput{
		// AccessToken: *string, // Required
		// AccountId: *string, // Required
	}

	if len(_ssoAccessToken) > 0 {
		input.AccessToken = aws.String(_ssoAccessToken)
	}
	if len(_ssoAccountId) > 0 {
		input.AccountId = aws.String(_ssoAccountId)
	}
	if len(_ssoMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssoMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssoNextToken) > 0 {
		input.NextToken = aws.String(_ssoNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccountRoles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sso.ListAccountRolesOutput
	p := sso.NewListAccountRolesPaginator(client, input)
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

// Lists all AWS accounts assigned to the user. These AWS accounts are assigned by
// the administrator of the account. For more information, see [Assign User Access]in the IAM Identity
// Center User Guide. This operation returns a paginated response.
//
// [Assign User Access]: https://docs.aws.amazon.com/singlesignon/latest/userguide/useraccess.html#assignusers
func sso_ListAccounts(cfg aws.Config, client *sso.Client) {
	input := &sso.ListAccountsInput{
		// AccessToken: *string, // Required
	}

	if len(_ssoAccessToken) > 0 {
		input.AccessToken = aws.String(_ssoAccessToken)
	}
	if len(_ssoMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ssoMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ssoNextToken) > 0 {
		input.NextToken = aws.String(_ssoNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccounts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sso.ListAccountsOutput
	p := sso.NewListAccountsPaginator(client, input)
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

// Removes the locally stored SSO tokens from the client-side cache and sends an
// API call to the IAM Identity Center service to invalidate the corresponding
// server-side IAM Identity Center sign in session.
//
// If a user uses IAM Identity Center to access the AWS CLI, the user’s IAM
// Identity Center sign in session is used to obtain an IAM session, as specified
// in the corresponding IAM Identity Center permission set. More specifically, IAM
// Identity Center assumes an IAM role in the target account on behalf of the user,
// and the corresponding temporary AWS credentials are returned to the client.
//
// After user logout, any existing IAM role sessions that were created by using
// IAM Identity Center permission sets continue based on the duration configured in
// the permission set. For more information, see [User authentications]in the IAM Identity Center User
// Guide.
//
// [User authentications]: https://docs.aws.amazon.com/singlesignon/latest/userguide/authconcept.html
func sso_Logout(cfg aws.Config, client *sso.Client) {
	input := &sso.LogoutInput{
		// AccessToken: *string, // Required
	}

	if len(_ssoAccessToken) > 0 {
		input.AccessToken = aws.String(_ssoAccessToken)
	}

	if resp, err := client.Logout(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_ssoCmd)
	_ssoCmd.Flags().SortFlags = false

	_ssoCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_ssoCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_ssoCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_ssoCmd.Flags().StringVarP(&_ssoAccessToken, "access-token", "", "", "Access Token")
	_ssoCmd.Flags().StringVarP(&_ssoAccountId, "account-id", "", "", "Account ID")
	_ssoCmd.Flags().StringVarP(&_ssoMaxResults, "max-results", "", "", "Max Results")
	_ssoCmd.Flags().StringVarP(&_ssoNextToken, "next-token", "", "", "Next Token")
	_ssoCmd.Flags().StringVarP(&_ssoRoleName, "role-name", "", "", "Role Name")

	_ssoCmd.Flags().BoolVarP(&_ssoGetRoleCredentials, "get-role-credentials", "", false, "Get Role Credentials")
	_ssoCmd.Flags().BoolVarP(&_ssoListAccountRoles, "list-account-roles", "", false, "List Account Roles")
	_ssoCmd.Flags().BoolVarP(&_ssoListAccounts, "list-accounts", "", false, "List Accounts")
	_ssoCmd.Flags().BoolVarP(&_ssoLogout, "logout", "", false, "Logout")

}
