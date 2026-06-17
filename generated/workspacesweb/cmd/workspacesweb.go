package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/workspacesweb"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// workspaceswebCmd represents the workspacesweb command
var _workspaceswebCmd = &cobra.Command{
	Use:   "workspacesweb",
	Short: "AWS workspacesweb CLI",
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
		client := workspacesweb.NewFromConfig(cfg)
		if _workspaceswebAssociateBrowserSettings {
			workspacesweb_AssociateBrowserSettings(cfg, client)
			return
		}
		if _workspaceswebAssociateDataProtectionSettings {
			workspacesweb_AssociateDataProtectionSettings(cfg, client)
			return
		}
		if _workspaceswebAssociateIpAccessSettings {
			workspacesweb_AssociateIpAccessSettings(cfg, client)
			return
		}
		if _workspaceswebAssociateNetworkSettings {
			workspacesweb_AssociateNetworkSettings(cfg, client)
			return
		}
		if _workspaceswebAssociateSessionLogger {
			workspacesweb_AssociateSessionLogger(cfg, client)
			return
		}
		if _workspaceswebAssociateTrustStore {
			workspacesweb_AssociateTrustStore(cfg, client)
			return
		}
		if _workspaceswebAssociateUserAccessLoggingSettings {
			workspacesweb_AssociateUserAccessLoggingSettings(cfg, client)
			return
		}
		if _workspaceswebAssociateUserSettings {
			workspacesweb_AssociateUserSettings(cfg, client)
			return
		}
		if _workspaceswebCreateBrowserSettings {
			workspacesweb_CreateBrowserSettings(cfg, client)
			return
		}
		if _workspaceswebCreateDataProtectionSettings {
			workspacesweb_CreateDataProtectionSettings(cfg, client)
			return
		}
		if _workspaceswebCreateIdentityProvider {
			workspacesweb_CreateIdentityProvider(cfg, client)
			return
		}
		if _workspaceswebCreateIpAccessSettings {
			workspacesweb_CreateIpAccessSettings(cfg, client)
			return
		}
		if _workspaceswebCreateNetworkSettings {
			workspacesweb_CreateNetworkSettings(cfg, client)
			return
		}
		if _workspaceswebCreatePortal {
			workspacesweb_CreatePortal(cfg, client)
			return
		}
		if _workspaceswebCreateSessionLogger {
			workspacesweb_CreateSessionLogger(cfg, client)
			return
		}
		if _workspaceswebCreateTrustStore {
			workspacesweb_CreateTrustStore(cfg, client)
			return
		}
		if _workspaceswebCreateUserAccessLoggingSettings {
			workspacesweb_CreateUserAccessLoggingSettings(cfg, client)
			return
		}
		if _workspaceswebCreateUserSettings {
			workspacesweb_CreateUserSettings(cfg, client)
			return
		}
		if _workspaceswebDeleteBrowserSettings {
			workspacesweb_DeleteBrowserSettings(cfg, client)
			return
		}
		if _workspaceswebDeleteDataProtectionSettings {
			workspacesweb_DeleteDataProtectionSettings(cfg, client)
			return
		}
		if _workspaceswebDeleteIdentityProvider {
			workspacesweb_DeleteIdentityProvider(cfg, client)
			return
		}
		if _workspaceswebDeleteIpAccessSettings {
			workspacesweb_DeleteIpAccessSettings(cfg, client)
			return
		}
		if _workspaceswebDeleteNetworkSettings {
			workspacesweb_DeleteNetworkSettings(cfg, client)
			return
		}
		if _workspaceswebDeletePortal {
			workspacesweb_DeletePortal(cfg, client)
			return
		}
		if _workspaceswebDeleteSessionLogger {
			workspacesweb_DeleteSessionLogger(cfg, client)
			return
		}
		if _workspaceswebDeleteTrustStore {
			workspacesweb_DeleteTrustStore(cfg, client)
			return
		}
		if _workspaceswebDeleteUserAccessLoggingSettings {
			workspacesweb_DeleteUserAccessLoggingSettings(cfg, client)
			return
		}
		if _workspaceswebDeleteUserSettings {
			workspacesweb_DeleteUserSettings(cfg, client)
			return
		}
		if _workspaceswebDisassociateBrowserSettings {
			workspacesweb_DisassociateBrowserSettings(cfg, client)
			return
		}
		if _workspaceswebDisassociateDataProtectionSettings {
			workspacesweb_DisassociateDataProtectionSettings(cfg, client)
			return
		}
		if _workspaceswebDisassociateIpAccessSettings {
			workspacesweb_DisassociateIpAccessSettings(cfg, client)
			return
		}
		if _workspaceswebDisassociateNetworkSettings {
			workspacesweb_DisassociateNetworkSettings(cfg, client)
			return
		}
		if _workspaceswebDisassociateSessionLogger {
			workspacesweb_DisassociateSessionLogger(cfg, client)
			return
		}
		if _workspaceswebDisassociateTrustStore {
			workspacesweb_DisassociateTrustStore(cfg, client)
			return
		}
		if _workspaceswebDisassociateUserAccessLoggingSettings {
			workspacesweb_DisassociateUserAccessLoggingSettings(cfg, client)
			return
		}
		if _workspaceswebDisassociateUserSettings {
			workspacesweb_DisassociateUserSettings(cfg, client)
			return
		}
		if _workspaceswebExpireSession {
			workspacesweb_ExpireSession(cfg, client)
			return
		}
		if _workspaceswebGetBrowserSettings {
			workspacesweb_GetBrowserSettings(cfg, client)
			return
		}
		if _workspaceswebGetDataProtectionSettings {
			workspacesweb_GetDataProtectionSettings(cfg, client)
			return
		}
		if _workspaceswebGetIdentityProvider {
			workspacesweb_GetIdentityProvider(cfg, client)
			return
		}
		if _workspaceswebGetIpAccessSettings {
			workspacesweb_GetIpAccessSettings(cfg, client)
			return
		}
		if _workspaceswebGetNetworkSettings {
			workspacesweb_GetNetworkSettings(cfg, client)
			return
		}
		if _workspaceswebGetPortal {
			workspacesweb_GetPortal(cfg, client)
			return
		}
		if _workspaceswebGetPortalServiceProviderMetadata {
			workspacesweb_GetPortalServiceProviderMetadata(cfg, client)
			return
		}
		if _workspaceswebGetSession {
			workspacesweb_GetSession(cfg, client)
			return
		}
		if _workspaceswebGetSessionLogger {
			workspacesweb_GetSessionLogger(cfg, client)
			return
		}
		if _workspaceswebGetTrustStore {
			workspacesweb_GetTrustStore(cfg, client)
			return
		}
		if _workspaceswebGetTrustStoreCertificate {
			workspacesweb_GetTrustStoreCertificate(cfg, client)
			return
		}
		if _workspaceswebGetUserAccessLoggingSettings {
			workspacesweb_GetUserAccessLoggingSettings(cfg, client)
			return
		}
		if _workspaceswebGetUserSettings {
			workspacesweb_GetUserSettings(cfg, client)
			return
		}
		if _workspaceswebListBrowserSettings {
			workspacesweb_ListBrowserSettings(cfg, client)
			return
		}
		if _workspaceswebListDataProtectionSettings {
			workspacesweb_ListDataProtectionSettings(cfg, client)
			return
		}
		if _workspaceswebListIdentityProviders {
			workspacesweb_ListIdentityProviders(cfg, client)
			return
		}
		if _workspaceswebListIpAccessSettings {
			workspacesweb_ListIpAccessSettings(cfg, client)
			return
		}
		if _workspaceswebListNetworkSettings {
			workspacesweb_ListNetworkSettings(cfg, client)
			return
		}
		if _workspaceswebListPortals {
			workspacesweb_ListPortals(cfg, client)
			return
		}
		if _workspaceswebListSessionLoggers {
			workspacesweb_ListSessionLoggers(cfg, client)
			return
		}
		if _workspaceswebListSessions {
			workspacesweb_ListSessions(cfg, client)
			return
		}
		if _workspaceswebListTagsForResource {
			workspacesweb_ListTagsForResource(cfg, client)
			return
		}
		if _workspaceswebListTrustStoreCertificates {
			workspacesweb_ListTrustStoreCertificates(cfg, client)
			return
		}
		if _workspaceswebListTrustStores {
			workspacesweb_ListTrustStores(cfg, client)
			return
		}
		if _workspaceswebListUserAccessLoggingSettings {
			workspacesweb_ListUserAccessLoggingSettings(cfg, client)
			return
		}
		if _workspaceswebListUserSettings {
			workspacesweb_ListUserSettings(cfg, client)
			return
		}
		if _workspaceswebTagResource {
			workspacesweb_TagResource(cfg, client)
			return
		}
		if _workspaceswebUntagResource {
			workspacesweb_UntagResource(cfg, client)
			return
		}
		if _workspaceswebUpdateBrowserSettings {
			workspacesweb_UpdateBrowserSettings(cfg, client)
			return
		}
		if _workspaceswebUpdateDataProtectionSettings {
			workspacesweb_UpdateDataProtectionSettings(cfg, client)
			return
		}
		if _workspaceswebUpdateIdentityProvider {
			workspacesweb_UpdateIdentityProvider(cfg, client)
			return
		}
		if _workspaceswebUpdateIpAccessSettings {
			workspacesweb_UpdateIpAccessSettings(cfg, client)
			return
		}
		if _workspaceswebUpdateNetworkSettings {
			workspacesweb_UpdateNetworkSettings(cfg, client)
			return
		}
		if _workspaceswebUpdatePortal {
			workspacesweb_UpdatePortal(cfg, client)
			return
		}
		if _workspaceswebUpdateSessionLogger {
			workspacesweb_UpdateSessionLogger(cfg, client)
			return
		}
		if _workspaceswebUpdateTrustStore {
			workspacesweb_UpdateTrustStore(cfg, client)
			return
		}
		if _workspaceswebUpdateUserAccessLoggingSettings {
			workspacesweb_UpdateUserAccessLoggingSettings(cfg, client)
			return
		}
		if _workspaceswebUpdateUserSettings {
			workspacesweb_UpdateUserSettings(cfg, client)
			return
		}

	},
}

var (
	_workspaceswebAssociateBrowserSettings              bool
	_workspaceswebAssociateDataProtectionSettings       bool
	_workspaceswebAssociateIpAccessSettings             bool
	_workspaceswebAssociateNetworkSettings              bool
	_workspaceswebAssociateSessionLogger                bool
	_workspaceswebAssociateTrustStore                   bool
	_workspaceswebAssociateUserAccessLoggingSettings    bool
	_workspaceswebAssociateUserSettings                 bool
	_workspaceswebCreateBrowserSettings                 bool
	_workspaceswebCreateDataProtectionSettings          bool
	_workspaceswebCreateIdentityProvider                bool
	_workspaceswebCreateIpAccessSettings                bool
	_workspaceswebCreateNetworkSettings                 bool
	_workspaceswebCreatePortal                          bool
	_workspaceswebCreateSessionLogger                   bool
	_workspaceswebCreateTrustStore                      bool
	_workspaceswebCreateUserAccessLoggingSettings       bool
	_workspaceswebCreateUserSettings                    bool
	_workspaceswebDeleteBrowserSettings                 bool
	_workspaceswebDeleteDataProtectionSettings          bool
	_workspaceswebDeleteIdentityProvider                bool
	_workspaceswebDeleteIpAccessSettings                bool
	_workspaceswebDeleteNetworkSettings                 bool
	_workspaceswebDeletePortal                          bool
	_workspaceswebDeleteSessionLogger                   bool
	_workspaceswebDeleteTrustStore                      bool
	_workspaceswebDeleteUserAccessLoggingSettings       bool
	_workspaceswebDeleteUserSettings                    bool
	_workspaceswebDisassociateBrowserSettings           bool
	_workspaceswebDisassociateDataProtectionSettings    bool
	_workspaceswebDisassociateIpAccessSettings          bool
	_workspaceswebDisassociateNetworkSettings           bool
	_workspaceswebDisassociateSessionLogger             bool
	_workspaceswebDisassociateTrustStore                bool
	_workspaceswebDisassociateUserAccessLoggingSettings bool
	_workspaceswebDisassociateUserSettings              bool
	_workspaceswebExpireSession                         bool
	_workspaceswebGetBrowserSettings                    bool
	_workspaceswebGetDataProtectionSettings             bool
	_workspaceswebGetIdentityProvider                   bool
	_workspaceswebGetIpAccessSettings                   bool
	_workspaceswebGetNetworkSettings                    bool
	_workspaceswebGetPortal                             bool
	_workspaceswebGetPortalServiceProviderMetadata      bool
	_workspaceswebGetSession                            bool
	_workspaceswebGetSessionLogger                      bool
	_workspaceswebGetTrustStore                         bool
	_workspaceswebGetTrustStoreCertificate              bool
	_workspaceswebGetUserAccessLoggingSettings          bool
	_workspaceswebGetUserSettings                       bool
	_workspaceswebListBrowserSettings                   bool
	_workspaceswebListDataProtectionSettings            bool
	_workspaceswebListIdentityProviders                 bool
	_workspaceswebListIpAccessSettings                  bool
	_workspaceswebListNetworkSettings                   bool
	_workspaceswebListPortals                           bool
	_workspaceswebListSessionLoggers                    bool
	_workspaceswebListSessions                          bool
	_workspaceswebListTagsForResource                   bool
	_workspaceswebListTrustStoreCertificates            bool
	_workspaceswebListTrustStores                       bool
	_workspaceswebListUserAccessLoggingSettings         bool
	_workspaceswebListUserSettings                      bool
	_workspaceswebTagResource                           bool
	_workspaceswebUntagResource                         bool
	_workspaceswebUpdateBrowserSettings                 bool
	_workspaceswebUpdateDataProtectionSettings          bool
	_workspaceswebUpdateIdentityProvider                bool
	_workspaceswebUpdateIpAccessSettings                bool
	_workspaceswebUpdateNetworkSettings                 bool
	_workspaceswebUpdatePortal                          bool
	_workspaceswebUpdateSessionLogger                   bool
	_workspaceswebUpdateTrustStore                      bool
	_workspaceswebUpdateUserAccessLoggingSettings       bool
	_workspaceswebUpdateUserSettings                    bool

	_workspaceswebAdditionalEncryptionContext        string
	_workspaceswebAuthenticationType                 string
	_workspaceswebBrandingConfigurationInput         string
	_workspaceswebBrowserPolicy                      string
	_workspaceswebBrowserSettingsArn                 string
	_workspaceswebCertificateList                    string
	_workspaceswebCertificatesToAdd                  string
	_workspaceswebCertificatesToDelete               []string
	_workspaceswebClientToken                        string
	_workspaceswebCookieSynchronizationConfiguration string
	_workspaceswebCopyAllowed                        string
	_workspaceswebCustomerManagedKey                 string
	_workspaceswebDataProtectionSettingsArn          string
	_workspaceswebDeepLinkAllowed                    string
	_workspaceswebDescription                        string
	_workspaceswebDisconnectTimeoutInMinutes         string
	_workspaceswebDisplayName                        string
	_workspaceswebDownloadAllowed                    string
	_workspaceswebEventFilter                        string
	_workspaceswebIdentityProviderArn                string
	_workspaceswebIdentityProviderDetails            string
	_workspaceswebIdentityProviderName               string
	_workspaceswebIdentityProviderType               string
	_workspaceswebIdleDisconnectTimeoutInMinutes     string
	_workspaceswebInlineRedactionConfiguration       string
	_workspaceswebInstanceType                       string
	_workspaceswebIpAccessSettingsArn                string
	_workspaceswebIpRules                            string
	_workspaceswebKinesisStreamArn                   string
	_workspaceswebLogConfiguration                   string
	_workspaceswebMaxConcurrentSessions              string
	_workspaceswebMaxResults                         string
	_workspaceswebNetworkSettingsArn                 string
	_workspaceswebNextToken                          string
	_workspaceswebPasteAllowed                       string
	_workspaceswebPortalArn                          string
	_workspaceswebPortalCustomDomain                 string
	_workspaceswebPortalId                           string
	_workspaceswebPrintAllowed                       string
	_workspaceswebResourceArn                        string
	_workspaceswebSecurityGroupIds                   []string
	_workspaceswebSessionId                          string
	_workspaceswebSessionLoggerArn                   string
	_workspaceswebSortBy                             string
	_workspaceswebStatus                             string
	_workspaceswebSubnetIds                          []string
	_workspaceswebTagKeys                            []string
	_workspaceswebTags                               string
	_workspaceswebThumbprint                         string
	_workspaceswebToolbarConfiguration               string
	_workspaceswebTrustStoreArn                      string
	_workspaceswebUploadAllowed                      string
	_workspaceswebUserAccessLoggingSettingsArn       string
	_workspaceswebUserSettingsArn                    string
	_workspaceswebUsername                           string
	_workspaceswebVpcId                              string
	_workspaceswebWebAuthnAllowed                    string
	_workspaceswebWebContentFilteringPolicy          string
)

// Associates a browser settings resource with a web portal.
func workspacesweb_AssociateBrowserSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.AssociateBrowserSettingsInput{
		// BrowserSettingsArn: *string, // Required
		// PortalArn: *string, // Required
	}

	if len(_workspaceswebBrowserSettingsArn) > 0 {
		input.BrowserSettingsArn = aws.String(_workspaceswebBrowserSettingsArn)
	}
	if len(_workspaceswebPortalArn) > 0 {
		input.PortalArn = aws.String(_workspaceswebPortalArn)
	}

	if resp, err := client.AssociateBrowserSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a data protection settings resource with a web portal.
func workspacesweb_AssociateDataProtectionSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.AssociateDataProtectionSettingsInput{
		// DataProtectionSettingsArn: *string, // Required
		// PortalArn: *string, // Required
	}

	if len(_workspaceswebDataProtectionSettingsArn) > 0 {
		input.DataProtectionSettingsArn = aws.String(_workspaceswebDataProtectionSettingsArn)
	}
	if len(_workspaceswebPortalArn) > 0 {
		input.PortalArn = aws.String(_workspaceswebPortalArn)
	}

	if resp, err := client.AssociateDataProtectionSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates an IP access settings resource with a web portal.
func workspacesweb_AssociateIpAccessSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.AssociateIpAccessSettingsInput{
		// IpAccessSettingsArn: *string, // Required
		// PortalArn: *string, // Required
	}

	if len(_workspaceswebIpAccessSettingsArn) > 0 {
		input.IpAccessSettingsArn = aws.String(_workspaceswebIpAccessSettingsArn)
	}
	if len(_workspaceswebPortalArn) > 0 {
		input.PortalArn = aws.String(_workspaceswebPortalArn)
	}

	if resp, err := client.AssociateIpAccessSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a network settings resource with a web portal.
func workspacesweb_AssociateNetworkSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.AssociateNetworkSettingsInput{
		// NetworkSettingsArn: *string, // Required
		// PortalArn: *string, // Required
	}

	if len(_workspaceswebNetworkSettingsArn) > 0 {
		input.NetworkSettingsArn = aws.String(_workspaceswebNetworkSettingsArn)
	}
	if len(_workspaceswebPortalArn) > 0 {
		input.PortalArn = aws.String(_workspaceswebPortalArn)
	}

	if resp, err := client.AssociateNetworkSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a session logger with a portal.
func workspacesweb_AssociateSessionLogger(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.AssociateSessionLoggerInput{
		// PortalArn: *string, // Required
		// SessionLoggerArn: *string, // Required
	}

	if len(_workspaceswebPortalArn) > 0 {
		input.PortalArn = aws.String(_workspaceswebPortalArn)
	}
	if len(_workspaceswebSessionLoggerArn) > 0 {
		input.SessionLoggerArn = aws.String(_workspaceswebSessionLoggerArn)
	}

	if resp, err := client.AssociateSessionLogger(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a trust store with a web portal.
func workspacesweb_AssociateTrustStore(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.AssociateTrustStoreInput{
		// PortalArn: *string, // Required
		// TrustStoreArn: *string, // Required
	}

	if len(_workspaceswebPortalArn) > 0 {
		input.PortalArn = aws.String(_workspaceswebPortalArn)
	}
	if len(_workspaceswebTrustStoreArn) > 0 {
		input.TrustStoreArn = aws.String(_workspaceswebTrustStoreArn)
	}

	if resp, err := client.AssociateTrustStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a user access logging settings resource with a web portal.
func workspacesweb_AssociateUserAccessLoggingSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.AssociateUserAccessLoggingSettingsInput{
		// PortalArn: *string, // Required
		// UserAccessLoggingSettingsArn: *string, // Required
	}

	if len(_workspaceswebPortalArn) > 0 {
		input.PortalArn = aws.String(_workspaceswebPortalArn)
	}
	if len(_workspaceswebUserAccessLoggingSettingsArn) > 0 {
		input.UserAccessLoggingSettingsArn = aws.String(_workspaceswebUserAccessLoggingSettingsArn)
	}

	if resp, err := client.AssociateUserAccessLoggingSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a user settings resource with a web portal.
func workspacesweb_AssociateUserSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.AssociateUserSettingsInput{
		// PortalArn: *string, // Required
		// UserSettingsArn: *string, // Required
	}

	if len(_workspaceswebPortalArn) > 0 {
		input.PortalArn = aws.String(_workspaceswebPortalArn)
	}
	if len(_workspaceswebUserSettingsArn) > 0 {
		input.UserSettingsArn = aws.String(_workspaceswebUserSettingsArn)
	}

	if resp, err := client.AssociateUserSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a browser settings resource that can be associated with a web portal.
// Once associated with a web portal, browser settings control how the browser will
// behave once a user starts a streaming session for the web portal.
func workspacesweb_CreateBrowserSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.CreateBrowserSettingsInput{}

	if len(_workspaceswebAdditionalEncryptionContext) > 0 {
		if err := assignInputField(input, "AdditionalEncryptionContext", _workspaceswebAdditionalEncryptionContext); err != nil {
			log.Errorf("invalid --additional-encryption-context: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebBrowserPolicy) > 0 {
		input.BrowserPolicy = aws.String(_workspaceswebBrowserPolicy)
	}
	if len(_workspaceswebClientToken) > 0 {
		input.ClientToken = aws.String(_workspaceswebClientToken)
	}
	if len(_workspaceswebCustomerManagedKey) > 0 {
		input.CustomerManagedKey = aws.String(_workspaceswebCustomerManagedKey)
	}
	if len(_workspaceswebTags) > 0 {
		if err := assignInputField(input, "Tags", _workspaceswebTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebWebContentFilteringPolicy) > 0 {
		if err := assignInputField(input, "WebContentFilteringPolicy", _workspaceswebWebContentFilteringPolicy); err != nil {
			log.Errorf("invalid --web-content-filtering-policy: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBrowserSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a data protection settings resource that can be associated with a web
// portal.
func workspacesweb_CreateDataProtectionSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.CreateDataProtectionSettingsInput{}

	if len(_workspaceswebAdditionalEncryptionContext) > 0 {
		if err := assignInputField(input, "AdditionalEncryptionContext", _workspaceswebAdditionalEncryptionContext); err != nil {
			log.Errorf("invalid --additional-encryption-context: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebClientToken) > 0 {
		input.ClientToken = aws.String(_workspaceswebClientToken)
	}
	if len(_workspaceswebCustomerManagedKey) > 0 {
		input.CustomerManagedKey = aws.String(_workspaceswebCustomerManagedKey)
	}
	if len(_workspaceswebDescription) > 0 {
		input.Description = aws.String(_workspaceswebDescription)
	}
	if len(_workspaceswebDisplayName) > 0 {
		input.DisplayName = aws.String(_workspaceswebDisplayName)
	}
	if len(_workspaceswebInlineRedactionConfiguration) > 0 {
		if err := assignInputField(input, "InlineRedactionConfiguration", _workspaceswebInlineRedactionConfiguration); err != nil {
			log.Errorf("invalid --inline-redaction-configuration: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebTags) > 0 {
		if err := assignInputField(input, "Tags", _workspaceswebTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataProtectionSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an identity provider resource that is then associated with a web portal.
func workspacesweb_CreateIdentityProvider(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.CreateIdentityProviderInput{
		// IdentityProviderDetails: map[string]string, // Required
		// IdentityProviderName: *string, // Required
		// IdentityProviderType: types.IdentityProviderType, // Required
		// PortalArn: *string, // Required
	}

	if len(_workspaceswebIdentityProviderDetails) > 0 {
		if err := assignInputField(input, "IdentityProviderDetails", _workspaceswebIdentityProviderDetails); err != nil {
			log.Errorf("invalid --identity-provider-details: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebIdentityProviderName) > 0 {
		input.IdentityProviderName = aws.String(_workspaceswebIdentityProviderName)
	}
	if len(_workspaceswebIdentityProviderType) > 0 {
		if err := assignInputField(input, "IdentityProviderType", _workspaceswebIdentityProviderType); err != nil {
			log.Errorf("invalid --identity-provider-type: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebPortalArn) > 0 {
		input.PortalArn = aws.String(_workspaceswebPortalArn)
	}
	if len(_workspaceswebClientToken) > 0 {
		input.ClientToken = aws.String(_workspaceswebClientToken)
	}
	if len(_workspaceswebTags) > 0 {
		if err := assignInputField(input, "Tags", _workspaceswebTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateIdentityProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an IP access settings resource that can be associated with a web portal.
func workspacesweb_CreateIpAccessSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.CreateIpAccessSettingsInput{
		// IpRules: []types.IpRule, // Required
	}

	if len(_workspaceswebIpRules) > 0 {
		if err := assignInputField(input, "IpRules", _workspaceswebIpRules); err != nil {
			log.Errorf("invalid --ip-rules: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebAdditionalEncryptionContext) > 0 {
		if err := assignInputField(input, "AdditionalEncryptionContext", _workspaceswebAdditionalEncryptionContext); err != nil {
			log.Errorf("invalid --additional-encryption-context: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebClientToken) > 0 {
		input.ClientToken = aws.String(_workspaceswebClientToken)
	}
	if len(_workspaceswebCustomerManagedKey) > 0 {
		input.CustomerManagedKey = aws.String(_workspaceswebCustomerManagedKey)
	}
	if len(_workspaceswebDescription) > 0 {
		input.Description = aws.String(_workspaceswebDescription)
	}
	if len(_workspaceswebDisplayName) > 0 {
		input.DisplayName = aws.String(_workspaceswebDisplayName)
	}
	if len(_workspaceswebTags) > 0 {
		if err := assignInputField(input, "Tags", _workspaceswebTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateIpAccessSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a network settings resource that can be associated with a web portal.
// Once associated with a web portal, network settings define how streaming
// instances will connect with your specified VPC.
func workspacesweb_CreateNetworkSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.CreateNetworkSettingsInput{
		// SecurityGroupIds: []string, // Required
		// SubnetIds: []string, // Required
		// VpcId: *string, // Required
	}

	if len(_workspaceswebSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _workspaceswebSecurityGroupIds...)
	}
	if len(_workspaceswebSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _workspaceswebSubnetIds...)
	}
	if len(_workspaceswebVpcId) > 0 {
		input.VpcId = aws.String(_workspaceswebVpcId)
	}
	if len(_workspaceswebClientToken) > 0 {
		input.ClientToken = aws.String(_workspaceswebClientToken)
	}
	if len(_workspaceswebTags) > 0 {
		if err := assignInputField(input, "Tags", _workspaceswebTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateNetworkSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a web portal.
func workspacesweb_CreatePortal(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.CreatePortalInput{}

	if len(_workspaceswebAdditionalEncryptionContext) > 0 {
		if err := assignInputField(input, "AdditionalEncryptionContext", _workspaceswebAdditionalEncryptionContext); err != nil {
			log.Errorf("invalid --additional-encryption-context: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebAuthenticationType) > 0 {
		if err := assignInputField(input, "AuthenticationType", _workspaceswebAuthenticationType); err != nil {
			log.Errorf("invalid --authentication-type: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebClientToken) > 0 {
		input.ClientToken = aws.String(_workspaceswebClientToken)
	}
	if len(_workspaceswebCustomerManagedKey) > 0 {
		input.CustomerManagedKey = aws.String(_workspaceswebCustomerManagedKey)
	}
	if len(_workspaceswebDisplayName) > 0 {
		input.DisplayName = aws.String(_workspaceswebDisplayName)
	}
	if len(_workspaceswebInstanceType) > 0 {
		if err := assignInputField(input, "InstanceType", _workspaceswebInstanceType); err != nil {
			log.Errorf("invalid --instance-type: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebMaxConcurrentSessions) > 0 {
		if err := assignInputField(input, "MaxConcurrentSessions", _workspaceswebMaxConcurrentSessions); err != nil {
			log.Errorf("invalid --max-concurrent-sessions: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebPortalCustomDomain) > 0 {
		input.PortalCustomDomain = aws.String(_workspaceswebPortalCustomDomain)
	}
	if len(_workspaceswebTags) > 0 {
		if err := assignInputField(input, "Tags", _workspaceswebTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePortal(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a session logger.
func workspacesweb_CreateSessionLogger(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.CreateSessionLoggerInput{
		// EventFilter: types.EventFilter, // Required
		// LogConfiguration: *types.LogConfiguration, // Required
	}

	if len(_workspaceswebEventFilter) > 0 {
		if err := assignInputField(input, "EventFilter", _workspaceswebEventFilter); err != nil {
			log.Errorf("invalid --event-filter: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebLogConfiguration) > 0 {
		if err := assignInputField(input, "LogConfiguration", _workspaceswebLogConfiguration); err != nil {
			log.Errorf("invalid --log-configuration: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebAdditionalEncryptionContext) > 0 {
		if err := assignInputField(input, "AdditionalEncryptionContext", _workspaceswebAdditionalEncryptionContext); err != nil {
			log.Errorf("invalid --additional-encryption-context: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebClientToken) > 0 {
		input.ClientToken = aws.String(_workspaceswebClientToken)
	}
	if len(_workspaceswebCustomerManagedKey) > 0 {
		input.CustomerManagedKey = aws.String(_workspaceswebCustomerManagedKey)
	}
	if len(_workspaceswebDisplayName) > 0 {
		input.DisplayName = aws.String(_workspaceswebDisplayName)
	}
	if len(_workspaceswebTags) > 0 {
		if err := assignInputField(input, "Tags", _workspaceswebTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSessionLogger(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a trust store that can be associated with a web portal. A trust store
// contains certificate authority (CA) certificates. Once associated with a web
// portal, the browser in a streaming session will recognize certificates that have
// been issued using any of the CAs in the trust store. If your organization has
// internal websites that use certificates issued by private CAs, you should add
// the private CA certificate to the trust store.
func workspacesweb_CreateTrustStore(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.CreateTrustStoreInput{
		// CertificateList: [][]byte, // Required
	}

	if len(_workspaceswebCertificateList) > 0 {
		if err := assignInputField(input, "CertificateList", _workspaceswebCertificateList); err != nil {
			log.Errorf("invalid --certificate-list: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebClientToken) > 0 {
		input.ClientToken = aws.String(_workspaceswebClientToken)
	}
	if len(_workspaceswebTags) > 0 {
		if err := assignInputField(input, "Tags", _workspaceswebTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTrustStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a user access logging settings resource that can be associated with a
// web portal.
func workspacesweb_CreateUserAccessLoggingSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.CreateUserAccessLoggingSettingsInput{
		// KinesisStreamArn: *string, // Required
	}

	if len(_workspaceswebKinesisStreamArn) > 0 {
		input.KinesisStreamArn = aws.String(_workspaceswebKinesisStreamArn)
	}
	if len(_workspaceswebClientToken) > 0 {
		input.ClientToken = aws.String(_workspaceswebClientToken)
	}
	if len(_workspaceswebTags) > 0 {
		if err := assignInputField(input, "Tags", _workspaceswebTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateUserAccessLoggingSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a user settings resource that can be associated with a web portal. Once
// associated with a web portal, user settings control how users can transfer data
// between a streaming session and the their local devices.
func workspacesweb_CreateUserSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.CreateUserSettingsInput{
		// CopyAllowed: types.EnabledType, // Required
		// DownloadAllowed: types.EnabledType, // Required
		// PasteAllowed: types.EnabledType, // Required
		// PrintAllowed: types.EnabledType, // Required
		// UploadAllowed: types.EnabledType, // Required
	}

	if len(_workspaceswebCopyAllowed) > 0 {
		if err := assignInputField(input, "CopyAllowed", _workspaceswebCopyAllowed); err != nil {
			log.Errorf("invalid --copy-allowed: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebDownloadAllowed) > 0 {
		if err := assignInputField(input, "DownloadAllowed", _workspaceswebDownloadAllowed); err != nil {
			log.Errorf("invalid --download-allowed: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebPasteAllowed) > 0 {
		if err := assignInputField(input, "PasteAllowed", _workspaceswebPasteAllowed); err != nil {
			log.Errorf("invalid --paste-allowed: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebPrintAllowed) > 0 {
		if err := assignInputField(input, "PrintAllowed", _workspaceswebPrintAllowed); err != nil {
			log.Errorf("invalid --print-allowed: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebUploadAllowed) > 0 {
		if err := assignInputField(input, "UploadAllowed", _workspaceswebUploadAllowed); err != nil {
			log.Errorf("invalid --upload-allowed: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebAdditionalEncryptionContext) > 0 {
		if err := assignInputField(input, "AdditionalEncryptionContext", _workspaceswebAdditionalEncryptionContext); err != nil {
			log.Errorf("invalid --additional-encryption-context: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebBrandingConfigurationInput) > 0 {
		if err := assignInputField(input, "BrandingConfigurationInput", _workspaceswebBrandingConfigurationInput); err != nil {
			log.Errorf("invalid --branding-configuration-input: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebClientToken) > 0 {
		input.ClientToken = aws.String(_workspaceswebClientToken)
	}
	if len(_workspaceswebCookieSynchronizationConfiguration) > 0 {
		if err := assignInputField(input, "CookieSynchronizationConfiguration", _workspaceswebCookieSynchronizationConfiguration); err != nil {
			log.Errorf("invalid --cookie-synchronization-configuration: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebCustomerManagedKey) > 0 {
		input.CustomerManagedKey = aws.String(_workspaceswebCustomerManagedKey)
	}
	if len(_workspaceswebDeepLinkAllowed) > 0 {
		if err := assignInputField(input, "DeepLinkAllowed", _workspaceswebDeepLinkAllowed); err != nil {
			log.Errorf("invalid --deep-link-allowed: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebDisconnectTimeoutInMinutes) > 0 {
		if err := assignInputField(input, "DisconnectTimeoutInMinutes", _workspaceswebDisconnectTimeoutInMinutes); err != nil {
			log.Errorf("invalid --disconnect-timeout-in-minutes: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebIdleDisconnectTimeoutInMinutes) > 0 {
		if err := assignInputField(input, "IdleDisconnectTimeoutInMinutes", _workspaceswebIdleDisconnectTimeoutInMinutes); err != nil {
			log.Errorf("invalid --idle-disconnect-timeout-in-minutes: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebTags) > 0 {
		if err := assignInputField(input, "Tags", _workspaceswebTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebToolbarConfiguration) > 0 {
		if err := assignInputField(input, "ToolbarConfiguration", _workspaceswebToolbarConfiguration); err != nil {
			log.Errorf("invalid --toolbar-configuration: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebWebAuthnAllowed) > 0 {
		if err := assignInputField(input, "WebAuthnAllowed", _workspaceswebWebAuthnAllowed); err != nil {
			log.Errorf("invalid --web-authn-allowed: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateUserSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes browser settings.
func workspacesweb_DeleteBrowserSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.DeleteBrowserSettingsInput{
		// BrowserSettingsArn: *string, // Required
	}

	if len(_workspaceswebBrowserSettingsArn) > 0 {
		input.BrowserSettingsArn = aws.String(_workspaceswebBrowserSettingsArn)
	}

	if resp, err := client.DeleteBrowserSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes data protection settings.
func workspacesweb_DeleteDataProtectionSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.DeleteDataProtectionSettingsInput{
		// DataProtectionSettingsArn: *string, // Required
	}

	if len(_workspaceswebDataProtectionSettingsArn) > 0 {
		input.DataProtectionSettingsArn = aws.String(_workspaceswebDataProtectionSettingsArn)
	}

	if resp, err := client.DeleteDataProtectionSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the identity provider.
func workspacesweb_DeleteIdentityProvider(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.DeleteIdentityProviderInput{
		// IdentityProviderArn: *string, // Required
	}

	if len(_workspaceswebIdentityProviderArn) > 0 {
		input.IdentityProviderArn = aws.String(_workspaceswebIdentityProviderArn)
	}

	if resp, err := client.DeleteIdentityProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes IP access settings.
func workspacesweb_DeleteIpAccessSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.DeleteIpAccessSettingsInput{
		// IpAccessSettingsArn: *string, // Required
	}

	if len(_workspaceswebIpAccessSettingsArn) > 0 {
		input.IpAccessSettingsArn = aws.String(_workspaceswebIpAccessSettingsArn)
	}

	if resp, err := client.DeleteIpAccessSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes network settings.
func workspacesweb_DeleteNetworkSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.DeleteNetworkSettingsInput{
		// NetworkSettingsArn: *string, // Required
	}

	if len(_workspaceswebNetworkSettingsArn) > 0 {
		input.NetworkSettingsArn = aws.String(_workspaceswebNetworkSettingsArn)
	}

	if resp, err := client.DeleteNetworkSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a web portal.
func workspacesweb_DeletePortal(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.DeletePortalInput{
		// PortalArn: *string, // Required
	}

	if len(_workspaceswebPortalArn) > 0 {
		input.PortalArn = aws.String(_workspaceswebPortalArn)
	}

	if resp, err := client.DeletePortal(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a session logger resource.
func workspacesweb_DeleteSessionLogger(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.DeleteSessionLoggerInput{
		// SessionLoggerArn: *string, // Required
	}

	if len(_workspaceswebSessionLoggerArn) > 0 {
		input.SessionLoggerArn = aws.String(_workspaceswebSessionLoggerArn)
	}

	if resp, err := client.DeleteSessionLogger(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the trust store.
func workspacesweb_DeleteTrustStore(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.DeleteTrustStoreInput{
		// TrustStoreArn: *string, // Required
	}

	if len(_workspaceswebTrustStoreArn) > 0 {
		input.TrustStoreArn = aws.String(_workspaceswebTrustStoreArn)
	}

	if resp, err := client.DeleteTrustStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes user access logging settings.
func workspacesweb_DeleteUserAccessLoggingSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.DeleteUserAccessLoggingSettingsInput{
		// UserAccessLoggingSettingsArn: *string, // Required
	}

	if len(_workspaceswebUserAccessLoggingSettingsArn) > 0 {
		input.UserAccessLoggingSettingsArn = aws.String(_workspaceswebUserAccessLoggingSettingsArn)
	}

	if resp, err := client.DeleteUserAccessLoggingSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes user settings.
func workspacesweb_DeleteUserSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.DeleteUserSettingsInput{
		// UserSettingsArn: *string, // Required
	}

	if len(_workspaceswebUserSettingsArn) > 0 {
		input.UserSettingsArn = aws.String(_workspaceswebUserSettingsArn)
	}

	if resp, err := client.DeleteUserSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates browser settings from a web portal.
func workspacesweb_DisassociateBrowserSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.DisassociateBrowserSettingsInput{
		// PortalArn: *string, // Required
	}

	if len(_workspaceswebPortalArn) > 0 {
		input.PortalArn = aws.String(_workspaceswebPortalArn)
	}

	if resp, err := client.DisassociateBrowserSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates data protection settings from a web portal.
func workspacesweb_DisassociateDataProtectionSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.DisassociateDataProtectionSettingsInput{
		// PortalArn: *string, // Required
	}

	if len(_workspaceswebPortalArn) > 0 {
		input.PortalArn = aws.String(_workspaceswebPortalArn)
	}

	if resp, err := client.DisassociateDataProtectionSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates IP access settings from a web portal.
func workspacesweb_DisassociateIpAccessSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.DisassociateIpAccessSettingsInput{
		// PortalArn: *string, // Required
	}

	if len(_workspaceswebPortalArn) > 0 {
		input.PortalArn = aws.String(_workspaceswebPortalArn)
	}

	if resp, err := client.DisassociateIpAccessSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates network settings from a web portal.
func workspacesweb_DisassociateNetworkSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.DisassociateNetworkSettingsInput{
		// PortalArn: *string, // Required
	}

	if len(_workspaceswebPortalArn) > 0 {
		input.PortalArn = aws.String(_workspaceswebPortalArn)
	}

	if resp, err := client.DisassociateNetworkSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a session logger from a portal.
func workspacesweb_DisassociateSessionLogger(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.DisassociateSessionLoggerInput{
		// PortalArn: *string, // Required
	}

	if len(_workspaceswebPortalArn) > 0 {
		input.PortalArn = aws.String(_workspaceswebPortalArn)
	}

	if resp, err := client.DisassociateSessionLogger(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a trust store from a web portal.
func workspacesweb_DisassociateTrustStore(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.DisassociateTrustStoreInput{
		// PortalArn: *string, // Required
	}

	if len(_workspaceswebPortalArn) > 0 {
		input.PortalArn = aws.String(_workspaceswebPortalArn)
	}

	if resp, err := client.DisassociateTrustStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates user access logging settings from a web portal.
func workspacesweb_DisassociateUserAccessLoggingSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.DisassociateUserAccessLoggingSettingsInput{
		// PortalArn: *string, // Required
	}

	if len(_workspaceswebPortalArn) > 0 {
		input.PortalArn = aws.String(_workspaceswebPortalArn)
	}

	if resp, err := client.DisassociateUserAccessLoggingSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates user settings from a web portal.
func workspacesweb_DisassociateUserSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.DisassociateUserSettingsInput{
		// PortalArn: *string, // Required
	}

	if len(_workspaceswebPortalArn) > 0 {
		input.PortalArn = aws.String(_workspaceswebPortalArn)
	}

	if resp, err := client.DisassociateUserSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Expires an active secure browser session.
func workspacesweb_ExpireSession(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.ExpireSessionInput{
		// PortalId: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_workspaceswebPortalId) > 0 {
		input.PortalId = aws.String(_workspaceswebPortalId)
	}
	if len(_workspaceswebSessionId) > 0 {
		input.SessionId = aws.String(_workspaceswebSessionId)
	}

	if resp, err := client.ExpireSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets browser settings.
func workspacesweb_GetBrowserSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.GetBrowserSettingsInput{
		// BrowserSettingsArn: *string, // Required
	}

	if len(_workspaceswebBrowserSettingsArn) > 0 {
		input.BrowserSettingsArn = aws.String(_workspaceswebBrowserSettingsArn)
	}

	if resp, err := client.GetBrowserSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the data protection settings.
func workspacesweb_GetDataProtectionSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.GetDataProtectionSettingsInput{
		// DataProtectionSettingsArn: *string, // Required
	}

	if len(_workspaceswebDataProtectionSettingsArn) > 0 {
		input.DataProtectionSettingsArn = aws.String(_workspaceswebDataProtectionSettingsArn)
	}

	if resp, err := client.GetDataProtectionSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the identity provider.
func workspacesweb_GetIdentityProvider(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.GetIdentityProviderInput{
		// IdentityProviderArn: *string, // Required
	}

	if len(_workspaceswebIdentityProviderArn) > 0 {
		input.IdentityProviderArn = aws.String(_workspaceswebIdentityProviderArn)
	}

	if resp, err := client.GetIdentityProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the IP access settings.
func workspacesweb_GetIpAccessSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.GetIpAccessSettingsInput{
		// IpAccessSettingsArn: *string, // Required
	}

	if len(_workspaceswebIpAccessSettingsArn) > 0 {
		input.IpAccessSettingsArn = aws.String(_workspaceswebIpAccessSettingsArn)
	}

	if resp, err := client.GetIpAccessSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the network settings.
func workspacesweb_GetNetworkSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.GetNetworkSettingsInput{
		// NetworkSettingsArn: *string, // Required
	}

	if len(_workspaceswebNetworkSettingsArn) > 0 {
		input.NetworkSettingsArn = aws.String(_workspaceswebNetworkSettingsArn)
	}

	if resp, err := client.GetNetworkSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the web portal.
func workspacesweb_GetPortal(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.GetPortalInput{
		// PortalArn: *string, // Required
	}

	if len(_workspaceswebPortalArn) > 0 {
		input.PortalArn = aws.String(_workspaceswebPortalArn)
	}

	if resp, err := client.GetPortal(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the service provider metadata.
func workspacesweb_GetPortalServiceProviderMetadata(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.GetPortalServiceProviderMetadataInput{
		// PortalArn: *string, // Required
	}

	if len(_workspaceswebPortalArn) > 0 {
		input.PortalArn = aws.String(_workspaceswebPortalArn)
	}

	if resp, err := client.GetPortalServiceProviderMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information for a secure browser session.
func workspacesweb_GetSession(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.GetSessionInput{
		// PortalId: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_workspaceswebPortalId) > 0 {
		input.PortalId = aws.String(_workspaceswebPortalId)
	}
	if len(_workspaceswebSessionId) > 0 {
		input.SessionId = aws.String(_workspaceswebSessionId)
	}

	if resp, err := client.GetSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details about a specific session logger resource.
func workspacesweb_GetSessionLogger(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.GetSessionLoggerInput{
		// SessionLoggerArn: *string, // Required
	}

	if len(_workspaceswebSessionLoggerArn) > 0 {
		input.SessionLoggerArn = aws.String(_workspaceswebSessionLoggerArn)
	}

	if resp, err := client.GetSessionLogger(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the trust store.
func workspacesweb_GetTrustStore(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.GetTrustStoreInput{
		// TrustStoreArn: *string, // Required
	}

	if len(_workspaceswebTrustStoreArn) > 0 {
		input.TrustStoreArn = aws.String(_workspaceswebTrustStoreArn)
	}

	if resp, err := client.GetTrustStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the trust store certificate.
func workspacesweb_GetTrustStoreCertificate(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.GetTrustStoreCertificateInput{
		// Thumbprint: *string, // Required
		// TrustStoreArn: *string, // Required
	}

	if len(_workspaceswebThumbprint) > 0 {
		input.Thumbprint = aws.String(_workspaceswebThumbprint)
	}
	if len(_workspaceswebTrustStoreArn) > 0 {
		input.TrustStoreArn = aws.String(_workspaceswebTrustStoreArn)
	}

	if resp, err := client.GetTrustStoreCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets user access logging settings.
func workspacesweb_GetUserAccessLoggingSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.GetUserAccessLoggingSettingsInput{
		// UserAccessLoggingSettingsArn: *string, // Required
	}

	if len(_workspaceswebUserAccessLoggingSettingsArn) > 0 {
		input.UserAccessLoggingSettingsArn = aws.String(_workspaceswebUserAccessLoggingSettingsArn)
	}

	if resp, err := client.GetUserAccessLoggingSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets user settings.
func workspacesweb_GetUserSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.GetUserSettingsInput{
		// UserSettingsArn: *string, // Required
	}

	if len(_workspaceswebUserSettingsArn) > 0 {
		input.UserSettingsArn = aws.String(_workspaceswebUserSettingsArn)
	}

	if resp, err := client.GetUserSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of browser settings.
func workspacesweb_ListBrowserSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.ListBrowserSettingsInput{}

	if len(_workspaceswebMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workspaceswebMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebNextToken) > 0 {
		input.NextToken = aws.String(_workspaceswebNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBrowserSettings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workspacesweb.ListBrowserSettingsOutput
	p := workspacesweb.NewListBrowserSettingsPaginator(client, input)
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

// Retrieves a list of data protection settings.
func workspacesweb_ListDataProtectionSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.ListDataProtectionSettingsInput{}

	if len(_workspaceswebMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workspaceswebMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebNextToken) > 0 {
		input.NextToken = aws.String(_workspaceswebNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDataProtectionSettings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workspacesweb.ListDataProtectionSettingsOutput
	p := workspacesweb.NewListDataProtectionSettingsPaginator(client, input)
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

// Retrieves a list of identity providers for a specific web portal.
func workspacesweb_ListIdentityProviders(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.ListIdentityProvidersInput{
		// PortalArn: *string, // Required
	}

	if len(_workspaceswebPortalArn) > 0 {
		input.PortalArn = aws.String(_workspaceswebPortalArn)
	}
	if len(_workspaceswebMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workspaceswebMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebNextToken) > 0 {
		input.NextToken = aws.String(_workspaceswebNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIdentityProviders(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workspacesweb.ListIdentityProvidersOutput
	p := workspacesweb.NewListIdentityProvidersPaginator(client, input)
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

// Retrieves a list of IP access settings.
func workspacesweb_ListIpAccessSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.ListIpAccessSettingsInput{}

	if len(_workspaceswebMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workspaceswebMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebNextToken) > 0 {
		input.NextToken = aws.String(_workspaceswebNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIpAccessSettings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workspacesweb.ListIpAccessSettingsOutput
	p := workspacesweb.NewListIpAccessSettingsPaginator(client, input)
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

// Retrieves a list of network settings.
func workspacesweb_ListNetworkSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.ListNetworkSettingsInput{}

	if len(_workspaceswebMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workspaceswebMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebNextToken) > 0 {
		input.NextToken = aws.String(_workspaceswebNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListNetworkSettings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workspacesweb.ListNetworkSettingsOutput
	p := workspacesweb.NewListNetworkSettingsPaginator(client, input)
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

// Retrieves a list or web portals.
func workspacesweb_ListPortals(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.ListPortalsInput{}

	if len(_workspaceswebMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workspaceswebMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebNextToken) > 0 {
		input.NextToken = aws.String(_workspaceswebNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPortals(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workspacesweb.ListPortalsOutput
	p := workspacesweb.NewListPortalsPaginator(client, input)
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

// Lists all available session logger resources.
func workspacesweb_ListSessionLoggers(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.ListSessionLoggersInput{}

	if len(_workspaceswebMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workspaceswebMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebNextToken) > 0 {
		input.NextToken = aws.String(_workspaceswebNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSessionLoggers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workspacesweb.ListSessionLoggersOutput
	p := workspacesweb.NewListSessionLoggersPaginator(client, input)
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

// Lists information for multiple secure browser sessions from a specific portal.
func workspacesweb_ListSessions(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.ListSessionsInput{
		// PortalId: *string, // Required
	}

	if len(_workspaceswebPortalId) > 0 {
		input.PortalId = aws.String(_workspaceswebPortalId)
	}
	if len(_workspaceswebMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workspaceswebMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebNextToken) > 0 {
		input.NextToken = aws.String(_workspaceswebNextToken)
	}
	if len(_workspaceswebSessionId) > 0 {
		input.SessionId = aws.String(_workspaceswebSessionId)
	}
	if len(_workspaceswebSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _workspaceswebSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebStatus) > 0 {
		if err := assignInputField(input, "Status", _workspaceswebStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebUsername) > 0 {
		input.Username = aws.String(_workspaceswebUsername)
	}

	if disablePaginator() {
		if resp, err := client.ListSessions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workspacesweb.ListSessionsOutput
	p := workspacesweb.NewListSessionsPaginator(client, input)
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

// Retrieves a list of tags for a resource.
func workspacesweb_ListTagsForResource(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_workspaceswebResourceArn) > 0 {
		input.ResourceArn = aws.String(_workspaceswebResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of trust store certificates.
func workspacesweb_ListTrustStoreCertificates(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.ListTrustStoreCertificatesInput{
		// TrustStoreArn: *string, // Required
	}

	if len(_workspaceswebTrustStoreArn) > 0 {
		input.TrustStoreArn = aws.String(_workspaceswebTrustStoreArn)
	}
	if len(_workspaceswebMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workspaceswebMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebNextToken) > 0 {
		input.NextToken = aws.String(_workspaceswebNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTrustStoreCertificates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workspacesweb.ListTrustStoreCertificatesOutput
	p := workspacesweb.NewListTrustStoreCertificatesPaginator(client, input)
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

// Retrieves a list of trust stores.
func workspacesweb_ListTrustStores(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.ListTrustStoresInput{}

	if len(_workspaceswebMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workspaceswebMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebNextToken) > 0 {
		input.NextToken = aws.String(_workspaceswebNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTrustStores(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workspacesweb.ListTrustStoresOutput
	p := workspacesweb.NewListTrustStoresPaginator(client, input)
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

// Retrieves a list of user access logging settings.
func workspacesweb_ListUserAccessLoggingSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.ListUserAccessLoggingSettingsInput{}

	if len(_workspaceswebMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workspaceswebMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebNextToken) > 0 {
		input.NextToken = aws.String(_workspaceswebNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListUserAccessLoggingSettings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workspacesweb.ListUserAccessLoggingSettingsOutput
	p := workspacesweb.NewListUserAccessLoggingSettingsPaginator(client, input)
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

// Retrieves a list of user settings.
func workspacesweb_ListUserSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.ListUserSettingsInput{}

	if len(_workspaceswebMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _workspaceswebMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebNextToken) > 0 {
		input.NextToken = aws.String(_workspaceswebNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListUserSettings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*workspacesweb.ListUserSettingsOutput
	p := workspacesweb.NewListUserSettingsPaginator(client, input)
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

// Adds or overwrites one or more tags for the specified resource.
func workspacesweb_TagResource(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_workspaceswebResourceArn) > 0 {
		input.ResourceArn = aws.String(_workspaceswebResourceArn)
	}
	if len(_workspaceswebTags) > 0 {
		if err := assignInputField(input, "Tags", _workspaceswebTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebClientToken) > 0 {
		input.ClientToken = aws.String(_workspaceswebClientToken)
	}

	if resp, err := client.TagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes one or more tags from the specified resource.
func workspacesweb_UntagResource(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_workspaceswebResourceArn) > 0 {
		input.ResourceArn = aws.String(_workspaceswebResourceArn)
	}
	if len(_workspaceswebTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _workspaceswebTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates browser settings.
func workspacesweb_UpdateBrowserSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.UpdateBrowserSettingsInput{
		// BrowserSettingsArn: *string, // Required
	}

	if len(_workspaceswebBrowserSettingsArn) > 0 {
		input.BrowserSettingsArn = aws.String(_workspaceswebBrowserSettingsArn)
	}
	if len(_workspaceswebBrowserPolicy) > 0 {
		input.BrowserPolicy = aws.String(_workspaceswebBrowserPolicy)
	}
	if len(_workspaceswebClientToken) > 0 {
		input.ClientToken = aws.String(_workspaceswebClientToken)
	}
	if len(_workspaceswebWebContentFilteringPolicy) > 0 {
		if err := assignInputField(input, "WebContentFilteringPolicy", _workspaceswebWebContentFilteringPolicy); err != nil {
			log.Errorf("invalid --web-content-filtering-policy: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateBrowserSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates data protection settings.
func workspacesweb_UpdateDataProtectionSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.UpdateDataProtectionSettingsInput{
		// DataProtectionSettingsArn: *string, // Required
	}

	if len(_workspaceswebDataProtectionSettingsArn) > 0 {
		input.DataProtectionSettingsArn = aws.String(_workspaceswebDataProtectionSettingsArn)
	}
	if len(_workspaceswebClientToken) > 0 {
		input.ClientToken = aws.String(_workspaceswebClientToken)
	}
	if len(_workspaceswebDescription) > 0 {
		input.Description = aws.String(_workspaceswebDescription)
	}
	if len(_workspaceswebDisplayName) > 0 {
		input.DisplayName = aws.String(_workspaceswebDisplayName)
	}
	if len(_workspaceswebInlineRedactionConfiguration) > 0 {
		if err := assignInputField(input, "InlineRedactionConfiguration", _workspaceswebInlineRedactionConfiguration); err != nil {
			log.Errorf("invalid --inline-redaction-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDataProtectionSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the identity provider.
func workspacesweb_UpdateIdentityProvider(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.UpdateIdentityProviderInput{
		// IdentityProviderArn: *string, // Required
	}

	if len(_workspaceswebIdentityProviderArn) > 0 {
		input.IdentityProviderArn = aws.String(_workspaceswebIdentityProviderArn)
	}
	if len(_workspaceswebClientToken) > 0 {
		input.ClientToken = aws.String(_workspaceswebClientToken)
	}
	if len(_workspaceswebIdentityProviderDetails) > 0 {
		if err := assignInputField(input, "IdentityProviderDetails", _workspaceswebIdentityProviderDetails); err != nil {
			log.Errorf("invalid --identity-provider-details: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebIdentityProviderName) > 0 {
		input.IdentityProviderName = aws.String(_workspaceswebIdentityProviderName)
	}
	if len(_workspaceswebIdentityProviderType) > 0 {
		if err := assignInputField(input, "IdentityProviderType", _workspaceswebIdentityProviderType); err != nil {
			log.Errorf("invalid --identity-provider-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateIdentityProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates IP access settings.
func workspacesweb_UpdateIpAccessSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.UpdateIpAccessSettingsInput{
		// IpAccessSettingsArn: *string, // Required
	}

	if len(_workspaceswebIpAccessSettingsArn) > 0 {
		input.IpAccessSettingsArn = aws.String(_workspaceswebIpAccessSettingsArn)
	}
	if len(_workspaceswebClientToken) > 0 {
		input.ClientToken = aws.String(_workspaceswebClientToken)
	}
	if len(_workspaceswebDescription) > 0 {
		input.Description = aws.String(_workspaceswebDescription)
	}
	if len(_workspaceswebDisplayName) > 0 {
		input.DisplayName = aws.String(_workspaceswebDisplayName)
	}
	if len(_workspaceswebIpRules) > 0 {
		if err := assignInputField(input, "IpRules", _workspaceswebIpRules); err != nil {
			log.Errorf("invalid --ip-rules: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateIpAccessSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates network settings.
func workspacesweb_UpdateNetworkSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.UpdateNetworkSettingsInput{
		// NetworkSettingsArn: *string, // Required
	}

	if len(_workspaceswebNetworkSettingsArn) > 0 {
		input.NetworkSettingsArn = aws.String(_workspaceswebNetworkSettingsArn)
	}
	if len(_workspaceswebClientToken) > 0 {
		input.ClientToken = aws.String(_workspaceswebClientToken)
	}
	if len(_workspaceswebSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _workspaceswebSecurityGroupIds...)
	}
	if len(_workspaceswebSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _workspaceswebSubnetIds...)
	}
	if len(_workspaceswebVpcId) > 0 {
		input.VpcId = aws.String(_workspaceswebVpcId)
	}

	if resp, err := client.UpdateNetworkSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a web portal.
func workspacesweb_UpdatePortal(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.UpdatePortalInput{
		// PortalArn: *string, // Required
	}

	if len(_workspaceswebPortalArn) > 0 {
		input.PortalArn = aws.String(_workspaceswebPortalArn)
	}
	if len(_workspaceswebAuthenticationType) > 0 {
		if err := assignInputField(input, "AuthenticationType", _workspaceswebAuthenticationType); err != nil {
			log.Errorf("invalid --authentication-type: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebDisplayName) > 0 {
		input.DisplayName = aws.String(_workspaceswebDisplayName)
	}
	if len(_workspaceswebInstanceType) > 0 {
		if err := assignInputField(input, "InstanceType", _workspaceswebInstanceType); err != nil {
			log.Errorf("invalid --instance-type: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebMaxConcurrentSessions) > 0 {
		if err := assignInputField(input, "MaxConcurrentSessions", _workspaceswebMaxConcurrentSessions); err != nil {
			log.Errorf("invalid --max-concurrent-sessions: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebPortalCustomDomain) > 0 {
		input.PortalCustomDomain = aws.String(_workspaceswebPortalCustomDomain)
	}

	if resp, err := client.UpdatePortal(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the details of a session logger.
func workspacesweb_UpdateSessionLogger(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.UpdateSessionLoggerInput{
		// SessionLoggerArn: *string, // Required
	}

	if len(_workspaceswebSessionLoggerArn) > 0 {
		input.SessionLoggerArn = aws.String(_workspaceswebSessionLoggerArn)
	}
	if len(_workspaceswebDisplayName) > 0 {
		input.DisplayName = aws.String(_workspaceswebDisplayName)
	}
	if len(_workspaceswebEventFilter) > 0 {
		if err := assignInputField(input, "EventFilter", _workspaceswebEventFilter); err != nil {
			log.Errorf("invalid --event-filter: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebLogConfiguration) > 0 {
		if err := assignInputField(input, "LogConfiguration", _workspaceswebLogConfiguration); err != nil {
			log.Errorf("invalid --log-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSessionLogger(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the trust store.
func workspacesweb_UpdateTrustStore(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.UpdateTrustStoreInput{
		// TrustStoreArn: *string, // Required
	}

	if len(_workspaceswebTrustStoreArn) > 0 {
		input.TrustStoreArn = aws.String(_workspaceswebTrustStoreArn)
	}
	if len(_workspaceswebCertificatesToAdd) > 0 {
		if err := assignInputField(input, "CertificatesToAdd", _workspaceswebCertificatesToAdd); err != nil {
			log.Errorf("invalid --certificates-to-add: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebCertificatesToDelete) > 0 {
		input.CertificatesToDelete = append([]string(nil), _workspaceswebCertificatesToDelete...)
	}
	if len(_workspaceswebClientToken) > 0 {
		input.ClientToken = aws.String(_workspaceswebClientToken)
	}

	if resp, err := client.UpdateTrustStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the user access logging settings.
func workspacesweb_UpdateUserAccessLoggingSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.UpdateUserAccessLoggingSettingsInput{
		// UserAccessLoggingSettingsArn: *string, // Required
	}

	if len(_workspaceswebUserAccessLoggingSettingsArn) > 0 {
		input.UserAccessLoggingSettingsArn = aws.String(_workspaceswebUserAccessLoggingSettingsArn)
	}
	if len(_workspaceswebClientToken) > 0 {
		input.ClientToken = aws.String(_workspaceswebClientToken)
	}
	if len(_workspaceswebKinesisStreamArn) > 0 {
		input.KinesisStreamArn = aws.String(_workspaceswebKinesisStreamArn)
	}

	if resp, err := client.UpdateUserAccessLoggingSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the user settings.
func workspacesweb_UpdateUserSettings(cfg aws.Config, client *workspacesweb.Client) {
	input := &workspacesweb.UpdateUserSettingsInput{
		// UserSettingsArn: *string, // Required
	}

	if len(_workspaceswebUserSettingsArn) > 0 {
		input.UserSettingsArn = aws.String(_workspaceswebUserSettingsArn)
	}
	if len(_workspaceswebBrandingConfigurationInput) > 0 {
		if err := assignInputField(input, "BrandingConfigurationInput", _workspaceswebBrandingConfigurationInput); err != nil {
			log.Errorf("invalid --branding-configuration-input: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebClientToken) > 0 {
		input.ClientToken = aws.String(_workspaceswebClientToken)
	}
	if len(_workspaceswebCookieSynchronizationConfiguration) > 0 {
		if err := assignInputField(input, "CookieSynchronizationConfiguration", _workspaceswebCookieSynchronizationConfiguration); err != nil {
			log.Errorf("invalid --cookie-synchronization-configuration: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebCopyAllowed) > 0 {
		if err := assignInputField(input, "CopyAllowed", _workspaceswebCopyAllowed); err != nil {
			log.Errorf("invalid --copy-allowed: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebDeepLinkAllowed) > 0 {
		if err := assignInputField(input, "DeepLinkAllowed", _workspaceswebDeepLinkAllowed); err != nil {
			log.Errorf("invalid --deep-link-allowed: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebDisconnectTimeoutInMinutes) > 0 {
		if err := assignInputField(input, "DisconnectTimeoutInMinutes", _workspaceswebDisconnectTimeoutInMinutes); err != nil {
			log.Errorf("invalid --disconnect-timeout-in-minutes: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebDownloadAllowed) > 0 {
		if err := assignInputField(input, "DownloadAllowed", _workspaceswebDownloadAllowed); err != nil {
			log.Errorf("invalid --download-allowed: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebIdleDisconnectTimeoutInMinutes) > 0 {
		if err := assignInputField(input, "IdleDisconnectTimeoutInMinutes", _workspaceswebIdleDisconnectTimeoutInMinutes); err != nil {
			log.Errorf("invalid --idle-disconnect-timeout-in-minutes: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebPasteAllowed) > 0 {
		if err := assignInputField(input, "PasteAllowed", _workspaceswebPasteAllowed); err != nil {
			log.Errorf("invalid --paste-allowed: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebPrintAllowed) > 0 {
		if err := assignInputField(input, "PrintAllowed", _workspaceswebPrintAllowed); err != nil {
			log.Errorf("invalid --print-allowed: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebToolbarConfiguration) > 0 {
		if err := assignInputField(input, "ToolbarConfiguration", _workspaceswebToolbarConfiguration); err != nil {
			log.Errorf("invalid --toolbar-configuration: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebUploadAllowed) > 0 {
		if err := assignInputField(input, "UploadAllowed", _workspaceswebUploadAllowed); err != nil {
			log.Errorf("invalid --upload-allowed: %s", err.Error())
			return
		}
	}
	if len(_workspaceswebWebAuthnAllowed) > 0 {
		if err := assignInputField(input, "WebAuthnAllowed", _workspaceswebWebAuthnAllowed); err != nil {
			log.Errorf("invalid --web-authn-allowed: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateUserSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_workspaceswebCmd)
	_workspaceswebCmd.Flags().SortFlags = false

	_workspaceswebCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_workspaceswebCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_workspaceswebCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebAdditionalEncryptionContext, "additional-encryption-context", "", "", "Additional Encryption Context")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebAuthenticationType, "authentication-type", "", "", "Authentication Type")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebBrandingConfigurationInput, "branding-configuration-input", "", "", "Branding Configuration Input")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebBrowserPolicy, "browser-policy", "", "", "Browser Policy")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebBrowserSettingsArn, "browser-settings-arn", "", "", "Browser Settings ARN")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebCertificateList, "certificate-list", "", "", "Certificate List")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebCertificatesToAdd, "certificates-to-add", "", "", "Certificates To Add")
	_workspaceswebCmd.Flags().StringSliceVarP(&_workspaceswebCertificatesToDelete, "certificates-to-delete", "", nil, "Certificates To Delete")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebClientToken, "client-token", "", "", "Client Token")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebCookieSynchronizationConfiguration, "cookie-synchronization-configuration", "", "", "Cookie Synchronization Configuration")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebCopyAllowed, "copy-allowed", "", "", "Copy Allowed")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebCustomerManagedKey, "customer-managed-key", "", "", "Customer Managed Key")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebDataProtectionSettingsArn, "data-protection-settings-arn", "", "", "Data Protection Settings ARN")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebDeepLinkAllowed, "deep-link-allowed", "", "", "Deep Link Allowed")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebDescription, "description", "", "", "Description")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebDisconnectTimeoutInMinutes, "disconnect-timeout-in-minutes", "", "", "Disconnect Timeout In Minutes")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebDisplayName, "display-name", "", "", "Display Name")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebDownloadAllowed, "download-allowed", "", "", "Download Allowed")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebEventFilter, "event-filter", "", "", "Event Filter")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebIdentityProviderArn, "identity-provider-arn", "", "", "Identity Provider ARN")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebIdentityProviderDetails, "identity-provider-details", "", "", "Identity Provider Details")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebIdentityProviderName, "identity-provider-name", "", "", "Identity Provider Name")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebIdentityProviderType, "identity-provider-type", "", "", "Identity Provider Type")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebIdleDisconnectTimeoutInMinutes, "idle-disconnect-timeout-in-minutes", "", "", "Idle Disconnect Timeout In Minutes")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebInlineRedactionConfiguration, "inline-redaction-configuration", "", "", "Inline Redaction Configuration")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebInstanceType, "instance-type", "", "", "Instance Type")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebIpAccessSettingsArn, "ip-access-settings-arn", "", "", "IP Access Settings ARN")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebIpRules, "ip-rules", "", "", "IP Rules")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebKinesisStreamArn, "kinesis-stream-arn", "", "", "Kinesis Stream ARN")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebLogConfiguration, "log-configuration", "", "", "Log Configuration")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebMaxConcurrentSessions, "max-concurrent-sessions", "", "", "Max Concurrent Sessions")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebMaxResults, "max-results", "", "", "Max Results")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebNetworkSettingsArn, "network-settings-arn", "", "", "Network Settings ARN")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebNextToken, "next-token", "", "", "Next Token")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebPasteAllowed, "paste-allowed", "", "", "Paste Allowed")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebPortalArn, "portal-arn", "", "", "Portal ARN")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebPortalCustomDomain, "portal-custom-domain", "", "", "Portal Custom Domain")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebPortalId, "portal-id", "", "", "Portal ID")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebPrintAllowed, "print-allowed", "", "", "Print Allowed")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebResourceArn, "resource-arn", "", "", "Resource ARN")
	_workspaceswebCmd.Flags().StringSliceVarP(&_workspaceswebSecurityGroupIds, "security-group-ids", "", nil, "Security Group Ids")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebSessionId, "session-id", "", "", "Session ID")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebSessionLoggerArn, "session-logger-arn", "", "", "Session Logger ARN")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebSortBy, "sort-by", "", "", "Sort By")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebStatus, "status", "", "", "Status")
	_workspaceswebCmd.Flags().StringSliceVarP(&_workspaceswebSubnetIds, "subnet-ids", "", nil, "Subnet Ids")
	_workspaceswebCmd.Flags().StringSliceVarP(&_workspaceswebTagKeys, "tag-keys", "", nil, "Tag Keys")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebTags, "tags", "", "", "Tags")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebThumbprint, "thumbprint", "", "", "Thumbprint")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebToolbarConfiguration, "toolbar-configuration", "", "", "Toolbar Configuration")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebTrustStoreArn, "trust-store-arn", "", "", "Trust Store ARN")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebUploadAllowed, "upload-allowed", "", "", "Upload Allowed")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebUserAccessLoggingSettingsArn, "user-access-logging-settings-arn", "", "", "User Access Logging Settings ARN")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebUserSettingsArn, "user-settings-arn", "", "", "User Settings ARN")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebUsername, "username", "", "", "Username")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebVpcId, "vpc-id", "", "", "VPC ID")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebWebAuthnAllowed, "web-authn-allowed", "", "", "Web Authn Allowed")
	_workspaceswebCmd.Flags().StringVarP(&_workspaceswebWebContentFilteringPolicy, "web-content-filtering-policy", "", "", "Web Content Filtering Policy")

	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebAssociateBrowserSettings, "associate-browser-settings", "", false, "Associate Browser Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebAssociateDataProtectionSettings, "associate-data-protection-settings", "", false, "Associate Data Protection Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebAssociateIpAccessSettings, "associate-ip-access-settings", "", false, "Associate IP Access Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebAssociateNetworkSettings, "associate-network-settings", "", false, "Associate Network Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebAssociateSessionLogger, "associate-session-logger", "", false, "Associate Session Logger")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebAssociateTrustStore, "associate-trust-store", "", false, "Associate Trust Store")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebAssociateUserAccessLoggingSettings, "associate-user-access-logging-settings", "", false, "Associate User Access Logging Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebAssociateUserSettings, "associate-user-settings", "", false, "Associate User Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebCreateBrowserSettings, "create-browser-settings", "", false, "Create Browser Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebCreateDataProtectionSettings, "create-data-protection-settings", "", false, "Create Data Protection Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebCreateIdentityProvider, "create-identity-provider", "", false, "Create Identity Provider")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebCreateIpAccessSettings, "create-ip-access-settings", "", false, "Create IP Access Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebCreateNetworkSettings, "create-network-settings", "", false, "Create Network Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebCreatePortal, "create-portal", "", false, "Create Portal")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebCreateSessionLogger, "create-session-logger", "", false, "Create Session Logger")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebCreateTrustStore, "create-trust-store", "", false, "Create Trust Store")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebCreateUserAccessLoggingSettings, "create-user-access-logging-settings", "", false, "Create User Access Logging Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebCreateUserSettings, "create-user-settings", "", false, "Create User Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebDeleteBrowserSettings, "delete-browser-settings", "", false, "Delete Browser Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebDeleteDataProtectionSettings, "delete-data-protection-settings", "", false, "Delete Data Protection Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebDeleteIdentityProvider, "delete-identity-provider", "", false, "Delete Identity Provider")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebDeleteIpAccessSettings, "delete-ip-access-settings", "", false, "Delete IP Access Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebDeleteNetworkSettings, "delete-network-settings", "", false, "Delete Network Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebDeletePortal, "delete-portal", "", false, "Delete Portal")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebDeleteSessionLogger, "delete-session-logger", "", false, "Delete Session Logger")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebDeleteTrustStore, "delete-trust-store", "", false, "Delete Trust Store")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebDeleteUserAccessLoggingSettings, "delete-user-access-logging-settings", "", false, "Delete User Access Logging Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebDeleteUserSettings, "delete-user-settings", "", false, "Delete User Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebDisassociateBrowserSettings, "disassociate-browser-settings", "", false, "Disassociate Browser Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebDisassociateDataProtectionSettings, "disassociate-data-protection-settings", "", false, "Disassociate Data Protection Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebDisassociateIpAccessSettings, "disassociate-ip-access-settings", "", false, "Disassociate IP Access Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebDisassociateNetworkSettings, "disassociate-network-settings", "", false, "Disassociate Network Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebDisassociateSessionLogger, "disassociate-session-logger", "", false, "Disassociate Session Logger")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebDisassociateTrustStore, "disassociate-trust-store", "", false, "Disassociate Trust Store")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebDisassociateUserAccessLoggingSettings, "disassociate-user-access-logging-settings", "", false, "Disassociate User Access Logging Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebDisassociateUserSettings, "disassociate-user-settings", "", false, "Disassociate User Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebExpireSession, "expire-session", "", false, "Expire Session")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebGetBrowserSettings, "get-browser-settings", "", false, "Get Browser Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebGetDataProtectionSettings, "get-data-protection-settings", "", false, "Get Data Protection Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebGetIdentityProvider, "get-identity-provider", "", false, "Get Identity Provider")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebGetIpAccessSettings, "get-ip-access-settings", "", false, "Get IP Access Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebGetNetworkSettings, "get-network-settings", "", false, "Get Network Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebGetPortal, "get-portal", "", false, "Get Portal")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebGetPortalServiceProviderMetadata, "get-portal-service-provider-metadata", "", false, "Get Portal Service Provider Metadata")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebGetSession, "get-session", "", false, "Get Session")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebGetSessionLogger, "get-session-logger", "", false, "Get Session Logger")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebGetTrustStore, "get-trust-store", "", false, "Get Trust Store")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebGetTrustStoreCertificate, "get-trust-store-certificate", "", false, "Get Trust Store Certificate")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebGetUserAccessLoggingSettings, "get-user-access-logging-settings", "", false, "Get User Access Logging Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebGetUserSettings, "get-user-settings", "", false, "Get User Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebListBrowserSettings, "list-browser-settings", "", false, "List Browser Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebListDataProtectionSettings, "list-data-protection-settings", "", false, "List Data Protection Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebListIdentityProviders, "list-identity-providers", "", false, "List Identity Providers")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebListIpAccessSettings, "list-ip-access-settings", "", false, "List IP Access Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebListNetworkSettings, "list-network-settings", "", false, "List Network Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebListPortals, "list-portals", "", false, "List Portals")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebListSessionLoggers, "list-session-loggers", "", false, "List Session Loggers")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebListSessions, "list-sessions", "", false, "List Sessions")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebListTrustStoreCertificates, "list-trust-store-certificates", "", false, "List Trust Store Certificates")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebListTrustStores, "list-trust-stores", "", false, "List Trust Stores")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebListUserAccessLoggingSettings, "list-user-access-logging-settings", "", false, "List User Access Logging Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebListUserSettings, "list-user-settings", "", false, "List User Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebTagResource, "tag-resource", "", false, "Tag Resource")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebUntagResource, "untag-resource", "", false, "Untag Resource")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebUpdateBrowserSettings, "update-browser-settings", "", false, "Update Browser Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebUpdateDataProtectionSettings, "update-data-protection-settings", "", false, "Update Data Protection Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebUpdateIdentityProvider, "update-identity-provider", "", false, "Update Identity Provider")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebUpdateIpAccessSettings, "update-ip-access-settings", "", false, "Update IP Access Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebUpdateNetworkSettings, "update-network-settings", "", false, "Update Network Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebUpdatePortal, "update-portal", "", false, "Update Portal")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebUpdateSessionLogger, "update-session-logger", "", false, "Update Session Logger")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebUpdateTrustStore, "update-trust-store", "", false, "Update Trust Store")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebUpdateUserAccessLoggingSettings, "update-user-access-logging-settings", "", false, "Update User Access Logging Settings")
	_workspaceswebCmd.Flags().BoolVarP(&_workspaceswebUpdateUserSettings, "update-user-settings", "", false, "Update User Settings")

}
