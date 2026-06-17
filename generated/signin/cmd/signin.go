package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/signin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// signinCmd represents the signin command
var _signinCmd = &cobra.Command{
	Use:   "signin",
	Short: "AWS signin CLI",
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
		client := signin.NewFromConfig(cfg)
		if _signinCreateOAuth2Token {
			signin_CreateOAuth2Token(cfg, client)
			return
		}

	},
}

var (
	_signinCreateOAuth2Token bool

	_signinTokenInput string
)

// CreateOAuth2Token API
// Path: /v1/token Request Method: POST Content-Type: application/json or
// application/x-www-form-urlencoded
//
// This API implements OAuth 2.0 flows for AWS Sign-In CLI clients, supporting
// both:
//
// - Authorization code redemption (grant_type=authorization_code) - NOT
// idempotent
// - Token refresh (grant_type=refresh_token) - Idempotent within token validity
// window
//
// The operation behavior is determined by the grant_type parameter in the request
// body:
//
// Authorization Code Flow (NOT Idempotent):
//
// - JSON or form-encoded body with client_id, grant_type=authorization_code,
// code, redirect_uri, code_verifier
// - Returns access_token, token_type, expires_in, refresh_token, and id_token
// - Each authorization code can only be used ONCE for security (prevents replay
// attacks)
//
// Token Refresh Flow (Idempotent):
//
// - JSON or form-encoded body with client_id, grant_type=refresh_token,
// refresh_token
// - Returns access_token, token_type, expires_in, and refresh_token (no
// id_token)
// - Multiple calls with same refresh_token return consistent results within
// validity window
//
// Authentication and authorization:
//
// - Confidential clients: sigv4 signing required with signin:ExchangeToken
// permissions
// - CLI clients (public): authn/authz skipped based on client_id & grant_type
//
// Note: This operation cannot be marked as (at)idempotent because it handles both
// idempotent (token refresh) and non-idempotent (auth code redemption) flows in a
// single endpoint.
func signin_CreateOAuth2Token(cfg aws.Config, client *signin.Client) {
	input := &signin.CreateOAuth2TokenInput{
		// TokenInput: *types.CreateOAuth2TokenRequestBody, // Required
	}

	if len(_signinTokenInput) > 0 {
		if err := assignInputField(input, "TokenInput", _signinTokenInput); err != nil {
			log.Errorf("invalid --token-input: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateOAuth2Token(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_signinCmd)
	_signinCmd.Flags().SortFlags = false

	_signinCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_signinCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_signinCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_signinCmd.Flags().StringVarP(&_signinTokenInput, "token-input", "", "", "Token Input")

	_signinCmd.Flags().BoolVarP(&_signinCreateOAuth2Token, "create-oauth2-token", "", false, "Create Oauth2 Token")

}
