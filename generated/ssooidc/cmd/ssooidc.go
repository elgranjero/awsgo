package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssooidc"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ssooidcCmd represents the ssooidc command
var _ssooidcCmd = &cobra.Command{
	Use:   "ssooidc",
	Short: "AWS ssooidc CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := ssooidc.NewFromConfig(cfg)
		if _ssooidcCreateToken {
			ssooidc_CreateToken(cfg, client)
			return
		}
		if _ssooidcCreateTokenWithIAM {
			ssooidc_CreateTokenWithIAM(cfg, client)
			return
		}
		if _ssooidcRegisterClient {
			ssooidc_RegisterClient(cfg, client)
			return
		}
		if _ssooidcStartDeviceAuthorization {
			ssooidc_StartDeviceAuthorization(cfg, client)
			return
		}

	},
}

var (
	_ssooidcCreateToken              bool
	_ssooidcCreateTokenWithIAM       bool
	_ssooidcRegisterClient           bool
	_ssooidcStartDeviceAuthorization bool

	_ssooidcAssertion              string
	_ssooidcClientId               string
	_ssooidcClientName             string
	_ssooidcClientSecret           string
	_ssooidcClientType             string
	_ssooidcCode                   string
	_ssooidcCodeVerifier           string
	_ssooidcDeviceCode             string
	_ssooidcEntitledApplicationArn string
	_ssooidcGrantType              string
	_ssooidcGrantTypes             []string
	_ssooidcIssuerUrl              string
	_ssooidcRedirectUri            string
	_ssooidcRedirectUris           []string
	_ssooidcRefreshToken           string
	_ssooidcRequestedTokenType     string
	_ssooidcScope                  []string
	_ssooidcScopes                 []string
	_ssooidcStartUrl               string
	_ssooidcSubjectToken           string
	_ssooidcSubjectTokenType       string
)

// Creates and returns access and refresh tokens for clients that are
// authenticated using client secrets. The access token can be used to fetch
// short-lived credentials for the assigned AWS accounts or to access application
// APIs using bearer authentication.
func ssooidc_CreateToken(cfg aws.Config, client *ssooidc.Client) {
	input := &ssooidc.CreateTokenInput{
		// ClientId: *string, // Required
		// ClientSecret: *string, // Required
		// GrantType: *string, // Required
	}

	if len(_ssooidcClientId) > 0 {
		input.ClientId = aws.String(_ssooidcClientId)
	}
	if len(_ssooidcClientSecret) > 0 {
		input.ClientSecret = aws.String(_ssooidcClientSecret)
	}
	if len(_ssooidcGrantType) > 0 {
		input.GrantType = aws.String(_ssooidcGrantType)
	}
	if len(_ssooidcCode) > 0 {
		input.Code = aws.String(_ssooidcCode)
	}
	if len(_ssooidcCodeVerifier) > 0 {
		input.CodeVerifier = aws.String(_ssooidcCodeVerifier)
	}
	if len(_ssooidcDeviceCode) > 0 {
		input.DeviceCode = aws.String(_ssooidcDeviceCode)
	}
	if len(_ssooidcRedirectUri) > 0 {
		input.RedirectUri = aws.String(_ssooidcRedirectUri)
	}
	if len(_ssooidcRefreshToken) > 0 {
		input.RefreshToken = aws.String(_ssooidcRefreshToken)
	}
	if len(_ssooidcScope) > 0 {
		input.Scope = append([]string(nil), _ssooidcScope...)
	}

	if resp, err := client.CreateToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates and returns access and refresh tokens for authorized client
// applications that are authenticated using any IAM entity, such as a service role
// or user. These tokens might contain defined scopes that specify permissions such
// as read:profile or write:data . Through downscoping, you can use the scopes
// parameter to request tokens with reduced permissions compared to the original
// client application's permissions or, if applicable, the refresh token's scopes.
// The access token can be used to fetch short-lived credentials for the assigned
// Amazon Web Services accounts or to access application APIs using bearer
// authentication.
//
// This API is used with Signature Version 4. For more information, see [Amazon Web Services Signature Version 4 for API Requests].
//
// [Amazon Web Services Signature Version 4 for API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_sigv.html
func ssooidc_CreateTokenWithIAM(cfg aws.Config, client *ssooidc.Client) {
	input := &ssooidc.CreateTokenWithIAMInput{
		// ClientId: *string, // Required
		// GrantType: *string, // Required
	}

	if len(_ssooidcClientId) > 0 {
		input.ClientId = aws.String(_ssooidcClientId)
	}
	if len(_ssooidcGrantType) > 0 {
		input.GrantType = aws.String(_ssooidcGrantType)
	}
	if len(_ssooidcAssertion) > 0 {
		input.Assertion = aws.String(_ssooidcAssertion)
	}
	if len(_ssooidcCode) > 0 {
		input.Code = aws.String(_ssooidcCode)
	}
	if len(_ssooidcCodeVerifier) > 0 {
		input.CodeVerifier = aws.String(_ssooidcCodeVerifier)
	}
	if len(_ssooidcRedirectUri) > 0 {
		input.RedirectUri = aws.String(_ssooidcRedirectUri)
	}
	if len(_ssooidcRefreshToken) > 0 {
		input.RefreshToken = aws.String(_ssooidcRefreshToken)
	}
	if len(_ssooidcRequestedTokenType) > 0 {
		input.RequestedTokenType = aws.String(_ssooidcRequestedTokenType)
	}
	if len(_ssooidcScope) > 0 {
		input.Scope = append([]string(nil), _ssooidcScope...)
	}
	if len(_ssooidcSubjectToken) > 0 {
		input.SubjectToken = aws.String(_ssooidcSubjectToken)
	}
	if len(_ssooidcSubjectTokenType) > 0 {
		input.SubjectTokenType = aws.String(_ssooidcSubjectTokenType)
	}

	if resp, err := client.CreateTokenWithIAM(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers a public client with IAM Identity Center. This allows clients to
// perform authorization using the authorization code grant with Proof Key for Code
// Exchange (PKCE) or the device code grant.
func ssooidc_RegisterClient(cfg aws.Config, client *ssooidc.Client) {
	input := &ssooidc.RegisterClientInput{
		// ClientName: *string, // Required
		// ClientType: *string, // Required
	}

	if len(_ssooidcClientName) > 0 {
		input.ClientName = aws.String(_ssooidcClientName)
	}
	if len(_ssooidcClientType) > 0 {
		input.ClientType = aws.String(_ssooidcClientType)
	}
	if len(_ssooidcEntitledApplicationArn) > 0 {
		input.EntitledApplicationArn = aws.String(_ssooidcEntitledApplicationArn)
	}
	if len(_ssooidcGrantTypes) > 0 {
		input.GrantTypes = append([]string(nil), _ssooidcGrantTypes...)
	}
	if len(_ssooidcIssuerUrl) > 0 {
		input.IssuerUrl = aws.String(_ssooidcIssuerUrl)
	}
	if len(_ssooidcRedirectUris) > 0 {
		input.RedirectUris = append([]string(nil), _ssooidcRedirectUris...)
	}
	if len(_ssooidcScopes) > 0 {
		input.Scopes = append([]string(nil), _ssooidcScopes...)
	}

	if resp, err := client.RegisterClient(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates device authorization by requesting a pair of verification codes from
// the authorization service.
func ssooidc_StartDeviceAuthorization(cfg aws.Config, client *ssooidc.Client) {
	input := &ssooidc.StartDeviceAuthorizationInput{
		// ClientId: *string, // Required
		// ClientSecret: *string, // Required
		// StartUrl: *string, // Required
	}

	if len(_ssooidcClientId) > 0 {
		input.ClientId = aws.String(_ssooidcClientId)
	}
	if len(_ssooidcClientSecret) > 0 {
		input.ClientSecret = aws.String(_ssooidcClientSecret)
	}
	if len(_ssooidcStartUrl) > 0 {
		input.StartUrl = aws.String(_ssooidcStartUrl)
	}

	if resp, err := client.StartDeviceAuthorization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_ssooidcCmd)
	_ssooidcCmd.Flags().SortFlags = false

	_ssooidcCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_ssooidcCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_ssooidcCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_ssooidcCmd.Flags().StringVarP(&_ssooidcAssertion, "assertion", "", "", "Assertion")
	_ssooidcCmd.Flags().StringVarP(&_ssooidcClientId, "client-id", "", "", "Client ID")
	_ssooidcCmd.Flags().StringVarP(&_ssooidcClientName, "client-name", "", "", "Client Name")
	_ssooidcCmd.Flags().StringVarP(&_ssooidcClientSecret, "client-secret", "", "", "Client Secret")
	_ssooidcCmd.Flags().StringVarP(&_ssooidcClientType, "client-type", "", "", "Client Type")
	_ssooidcCmd.Flags().StringVarP(&_ssooidcCode, "code", "", "", "Code")
	_ssooidcCmd.Flags().StringVarP(&_ssooidcCodeVerifier, "code-verifier", "", "", "Code Verifier")
	_ssooidcCmd.Flags().StringVarP(&_ssooidcDeviceCode, "device-code", "", "", "Device Code")
	_ssooidcCmd.Flags().StringVarP(&_ssooidcEntitledApplicationArn, "entitled-application-arn", "", "", "Entitled Application ARN")
	_ssooidcCmd.Flags().StringVarP(&_ssooidcGrantType, "grant-type", "", "", "Grant Type")
	_ssooidcCmd.Flags().StringSliceVarP(&_ssooidcGrantTypes, "grant-types", "", nil, "Grant Types")
	_ssooidcCmd.Flags().StringVarP(&_ssooidcIssuerUrl, "issuer-url", "", "", "Issuer URL")
	_ssooidcCmd.Flags().StringVarP(&_ssooidcRedirectUri, "redirect-uri", "", "", "Redirect URI")
	_ssooidcCmd.Flags().StringSliceVarP(&_ssooidcRedirectUris, "redirect-uris", "", nil, "Redirect Uris")
	_ssooidcCmd.Flags().StringVarP(&_ssooidcRefreshToken, "refresh-token", "", "", "Refresh Token")
	_ssooidcCmd.Flags().StringVarP(&_ssooidcRequestedTokenType, "requested-token-type", "", "", "Requested Token Type")
	_ssooidcCmd.Flags().StringSliceVarP(&_ssooidcScope, "scope", "", nil, "Scope")
	_ssooidcCmd.Flags().StringSliceVarP(&_ssooidcScopes, "scopes", "", nil, "Scopes")
	_ssooidcCmd.Flags().StringVarP(&_ssooidcStartUrl, "start-url", "", "", "Start URL")
	_ssooidcCmd.Flags().StringVarP(&_ssooidcSubjectToken, "subject-token", "", "", "Subject Token")
	_ssooidcCmd.Flags().StringVarP(&_ssooidcSubjectTokenType, "subject-token-type", "", "", "Subject Token Type")

	_ssooidcCmd.Flags().BoolVarP(&_ssooidcCreateToken, "create-token", "", false, "Create Token")
	_ssooidcCmd.Flags().BoolVarP(&_ssooidcCreateTokenWithIAM, "create-token-with-iam", "", false, "Create Token With IAM")
	_ssooidcCmd.Flags().BoolVarP(&_ssooidcRegisterClient, "register-client", "", false, "Register Client")
	_ssooidcCmd.Flags().BoolVarP(&_ssooidcStartDeviceAuthorization, "start-device-authorization", "", false, "Start Device Authorization")

}
