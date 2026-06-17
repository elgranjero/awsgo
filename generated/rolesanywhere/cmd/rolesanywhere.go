package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rolesanywhere"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rolesanywhereCmd represents the rolesanywhere command
var _rolesanywhereCmd = &cobra.Command{
	Use:   "rolesanywhere",
	Short: "AWS rolesanywhere CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := rolesanywhere.NewFromConfig(cfg)
		if _rolesanywhereCreateProfile {
			rolesanywhere_CreateProfile(cfg, client)
			return
		}
		if _rolesanywhereCreateTrustAnchor {
			rolesanywhere_CreateTrustAnchor(cfg, client)
			return
		}
		if _rolesanywhereDeleteAttributeMapping {
			rolesanywhere_DeleteAttributeMapping(cfg, client)
			return
		}
		if _rolesanywhereDeleteCrl {
			rolesanywhere_DeleteCrl(cfg, client)
			return
		}
		if _rolesanywhereDeleteProfile {
			rolesanywhere_DeleteProfile(cfg, client)
			return
		}
		if _rolesanywhereDeleteTrustAnchor {
			rolesanywhere_DeleteTrustAnchor(cfg, client)
			return
		}
		if _rolesanywhereDisableCrl {
			rolesanywhere_DisableCrl(cfg, client)
			return
		}
		if _rolesanywhereDisableProfile {
			rolesanywhere_DisableProfile(cfg, client)
			return
		}
		if _rolesanywhereDisableTrustAnchor {
			rolesanywhere_DisableTrustAnchor(cfg, client)
			return
		}
		if _rolesanywhereEnableCrl {
			rolesanywhere_EnableCrl(cfg, client)
			return
		}
		if _rolesanywhereEnableProfile {
			rolesanywhere_EnableProfile(cfg, client)
			return
		}
		if _rolesanywhereEnableTrustAnchor {
			rolesanywhere_EnableTrustAnchor(cfg, client)
			return
		}
		if _rolesanywhereGetCrl {
			rolesanywhere_GetCrl(cfg, client)
			return
		}
		if _rolesanywhereGetProfile {
			rolesanywhere_GetProfile(cfg, client)
			return
		}
		if _rolesanywhereGetSubject {
			rolesanywhere_GetSubject(cfg, client)
			return
		}
		if _rolesanywhereGetTrustAnchor {
			rolesanywhere_GetTrustAnchor(cfg, client)
			return
		}
		if _rolesanywhereImportCrl {
			rolesanywhere_ImportCrl(cfg, client)
			return
		}
		if _rolesanywhereListCrls {
			rolesanywhere_ListCrls(cfg, client)
			return
		}
		if _rolesanywhereListProfiles {
			rolesanywhere_ListProfiles(cfg, client)
			return
		}
		if _rolesanywhereListSubjects {
			rolesanywhere_ListSubjects(cfg, client)
			return
		}
		if _rolesanywhereListTagsForResource {
			rolesanywhere_ListTagsForResource(cfg, client)
			return
		}
		if _rolesanywhereListTrustAnchors {
			rolesanywhere_ListTrustAnchors(cfg, client)
			return
		}
		if _rolesanywherePutAttributeMapping {
			rolesanywhere_PutAttributeMapping(cfg, client)
			return
		}
		if _rolesanywherePutNotificationSettings {
			rolesanywhere_PutNotificationSettings(cfg, client)
			return
		}
		if _rolesanywhereResetNotificationSettings {
			rolesanywhere_ResetNotificationSettings(cfg, client)
			return
		}
		if _rolesanywhereTagResource {
			rolesanywhere_TagResource(cfg, client)
			return
		}
		if _rolesanywhereUntagResource {
			rolesanywhere_UntagResource(cfg, client)
			return
		}
		if _rolesanywhereUpdateCrl {
			rolesanywhere_UpdateCrl(cfg, client)
			return
		}
		if _rolesanywhereUpdateProfile {
			rolesanywhere_UpdateProfile(cfg, client)
			return
		}
		if _rolesanywhereUpdateTrustAnchor {
			rolesanywhere_UpdateTrustAnchor(cfg, client)
			return
		}

	},
}

var (
	_rolesanywhereCreateProfile             bool
	_rolesanywhereCreateTrustAnchor         bool
	_rolesanywhereDeleteAttributeMapping    bool
	_rolesanywhereDeleteCrl                 bool
	_rolesanywhereDeleteProfile             bool
	_rolesanywhereDeleteTrustAnchor         bool
	_rolesanywhereDisableCrl                bool
	_rolesanywhereDisableProfile            bool
	_rolesanywhereDisableTrustAnchor        bool
	_rolesanywhereEnableCrl                 bool
	_rolesanywhereEnableProfile             bool
	_rolesanywhereEnableTrustAnchor         bool
	_rolesanywhereGetCrl                    bool
	_rolesanywhereGetProfile                bool
	_rolesanywhereGetSubject                bool
	_rolesanywhereGetTrustAnchor            bool
	_rolesanywhereImportCrl                 bool
	_rolesanywhereListCrls                  bool
	_rolesanywhereListProfiles              bool
	_rolesanywhereListSubjects              bool
	_rolesanywhereListTagsForResource       bool
	_rolesanywhereListTrustAnchors          bool
	_rolesanywherePutAttributeMapping       bool
	_rolesanywherePutNotificationSettings   bool
	_rolesanywhereResetNotificationSettings bool
	_rolesanywhereTagResource               bool
	_rolesanywhereUntagResource             bool
	_rolesanywhereUpdateCrl                 bool
	_rolesanywhereUpdateProfile             bool
	_rolesanywhereUpdateTrustAnchor         bool

	_rolesanywhereAcceptRoleSessionName     string
	_rolesanywhereCertificateField          string
	_rolesanywhereCrlData                   string
	_rolesanywhereCrlId                     string
	_rolesanywhereDurationSeconds           string
	_rolesanywhereEnabled                   string
	_rolesanywhereManagedPolicyArns         []string
	_rolesanywhereMappingRules              string
	_rolesanywhereName                      string
	_rolesanywhereNextToken                 string
	_rolesanywhereNotificationSettingKeys   string
	_rolesanywhereNotificationSettings      string
	_rolesanywherePageSize                  string
	_rolesanywhereProfileId                 string
	_rolesanywhereRequireInstanceProperties string
	_rolesanywhereResourceArn               string
	_rolesanywhereRoleArns                  []string
	_rolesanywhereSessionPolicy             string
	_rolesanywhereSource                    string
	_rolesanywhereSpecifiers                []string
	_rolesanywhereSubjectId                 string
	_rolesanywhereTagKeys                   []string
	_rolesanywhereTags                      string
	_rolesanywhereTrustAnchorArn            string
	_rolesanywhereTrustAnchorId             string
)

// Creates a profile, a list of the roles that Roles Anywhere service is trusted
// to assume. You use profiles to intersect permissions with IAM managed policies.
//
// Required permissions: rolesanywhere:CreateProfile .
func rolesanywhere_CreateProfile(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.CreateProfileInput{
		// Name: *string, // Required
		// RoleArns: []string, // Required
	}

	if len(_rolesanywhereName) > 0 {
		input.Name = aws.String(_rolesanywhereName)
	}
	if len(_rolesanywhereRoleArns) > 0 {
		input.RoleArns = append([]string(nil), _rolesanywhereRoleArns...)
	}
	if len(_rolesanywhereAcceptRoleSessionName) > 0 {
		if err := assignInputField(input, "AcceptRoleSessionName", _rolesanywhereAcceptRoleSessionName); err != nil {
			log.Errorf("invalid --accept-role-session-name: %s", err.Error())
			return
		}
	}
	if len(_rolesanywhereDurationSeconds) > 0 {
		if err := assignInputField(input, "DurationSeconds", _rolesanywhereDurationSeconds); err != nil {
			log.Errorf("invalid --duration-seconds: %s", err.Error())
			return
		}
	}
	if len(_rolesanywhereEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _rolesanywhereEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_rolesanywhereManagedPolicyArns) > 0 {
		input.ManagedPolicyArns = append([]string(nil), _rolesanywhereManagedPolicyArns...)
	}
	if len(_rolesanywhereRequireInstanceProperties) > 0 {
		if err := assignInputField(input, "RequireInstanceProperties", _rolesanywhereRequireInstanceProperties); err != nil {
			log.Errorf("invalid --require-instance-properties: %s", err.Error())
			return
		}
	}
	if len(_rolesanywhereSessionPolicy) > 0 {
		input.SessionPolicy = aws.String(_rolesanywhereSessionPolicy)
	}
	if len(_rolesanywhereTags) > 0 {
		if err := assignInputField(input, "Tags", _rolesanywhereTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a trust anchor to establish trust between IAM Roles Anywhere and your
// certificate authority (CA). You can define a trust anchor as a reference to an
// Private Certificate Authority (Private CA) or by uploading a CA certificate.
// Your Amazon Web Services workloads can authenticate with the trust anchor using
// certificates issued by the CA in exchange for temporary Amazon Web Services
// credentials.
//
// Required permissions: rolesanywhere:CreateTrustAnchor .
func rolesanywhere_CreateTrustAnchor(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.CreateTrustAnchorInput{
		// Name: *string, // Required
		// Source: *types.Source, // Required
	}

	if len(_rolesanywhereName) > 0 {
		input.Name = aws.String(_rolesanywhereName)
	}
	if len(_rolesanywhereSource) > 0 {
		if err := assignInputField(input, "Source", _rolesanywhereSource); err != nil {
			log.Errorf("invalid --source: %s", err.Error())
			return
		}
	}
	if len(_rolesanywhereEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _rolesanywhereEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_rolesanywhereNotificationSettings) > 0 {
		if err := assignInputField(input, "NotificationSettings", _rolesanywhereNotificationSettings); err != nil {
			log.Errorf("invalid --notification-settings: %s", err.Error())
			return
		}
	}
	if len(_rolesanywhereTags) > 0 {
		if err := assignInputField(input, "Tags", _rolesanywhereTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTrustAnchor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an entry from the attribute mapping rules enforced by a given profile.
func rolesanywhere_DeleteAttributeMapping(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.DeleteAttributeMappingInput{
		// CertificateField: types.CertificateField, // Required
		// ProfileId: *string, // Required
	}

	if len(_rolesanywhereCertificateField) > 0 {
		if err := assignInputField(input, "CertificateField", _rolesanywhereCertificateField); err != nil {
			log.Errorf("invalid --certificate-field: %s", err.Error())
			return
		}
	}
	if len(_rolesanywhereProfileId) > 0 {
		input.ProfileId = aws.String(_rolesanywhereProfileId)
	}
	if len(_rolesanywhereSpecifiers) > 0 {
		input.Specifiers = append([]string(nil), _rolesanywhereSpecifiers...)
	}

	if resp, err := client.DeleteAttributeMapping(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a certificate revocation list (CRL).
// Required permissions: rolesanywhere:DeleteCrl .
func rolesanywhere_DeleteCrl(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.DeleteCrlInput{
		// CrlId: *string, // Required
	}

	if len(_rolesanywhereCrlId) > 0 {
		input.CrlId = aws.String(_rolesanywhereCrlId)
	}

	if resp, err := client.DeleteCrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a profile.
// Required permissions: rolesanywhere:DeleteProfile .
func rolesanywhere_DeleteProfile(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.DeleteProfileInput{
		// ProfileId: *string, // Required
	}

	if len(_rolesanywhereProfileId) > 0 {
		input.ProfileId = aws.String(_rolesanywhereProfileId)
	}

	if resp, err := client.DeleteProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a trust anchor.
// Required permissions: rolesanywhere:DeleteTrustAnchor .
func rolesanywhere_DeleteTrustAnchor(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.DeleteTrustAnchorInput{
		// TrustAnchorId: *string, // Required
	}

	if len(_rolesanywhereTrustAnchorId) > 0 {
		input.TrustAnchorId = aws.String(_rolesanywhereTrustAnchorId)
	}

	if resp, err := client.DeleteTrustAnchor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables a certificate revocation list (CRL).
// Required permissions: rolesanywhere:DisableCrl .
func rolesanywhere_DisableCrl(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.DisableCrlInput{
		// CrlId: *string, // Required
	}

	if len(_rolesanywhereCrlId) > 0 {
		input.CrlId = aws.String(_rolesanywhereCrlId)
	}

	if resp, err := client.DisableCrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables a profile. When disabled, temporary credential requests with this
// profile fail.
//
// Required permissions: rolesanywhere:DisableProfile .
func rolesanywhere_DisableProfile(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.DisableProfileInput{
		// ProfileId: *string, // Required
	}

	if len(_rolesanywhereProfileId) > 0 {
		input.ProfileId = aws.String(_rolesanywhereProfileId)
	}

	if resp, err := client.DisableProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables a trust anchor. When disabled, temporary credential requests
// specifying this trust anchor are unauthorized.
//
// Required permissions: rolesanywhere:DisableTrustAnchor .
func rolesanywhere_DisableTrustAnchor(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.DisableTrustAnchorInput{
		// TrustAnchorId: *string, // Required
	}

	if len(_rolesanywhereTrustAnchorId) > 0 {
		input.TrustAnchorId = aws.String(_rolesanywhereTrustAnchorId)
	}

	if resp, err := client.DisableTrustAnchor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables a certificate revocation list (CRL). When enabled, certificates stored
// in the CRL are unauthorized to receive session credentials.
//
// Required permissions: rolesanywhere:EnableCrl .
func rolesanywhere_EnableCrl(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.EnableCrlInput{
		// CrlId: *string, // Required
	}

	if len(_rolesanywhereCrlId) > 0 {
		input.CrlId = aws.String(_rolesanywhereCrlId)
	}

	if resp, err := client.EnableCrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables temporary credential requests for a profile.
// Required permissions: rolesanywhere:EnableProfile .
func rolesanywhere_EnableProfile(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.EnableProfileInput{
		// ProfileId: *string, // Required
	}

	if len(_rolesanywhereProfileId) > 0 {
		input.ProfileId = aws.String(_rolesanywhereProfileId)
	}

	if resp, err := client.EnableProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables a trust anchor. When enabled, certificates in the trust anchor chain
// are authorized for trust validation.
//
// Required permissions: rolesanywhere:EnableTrustAnchor .
func rolesanywhere_EnableTrustAnchor(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.EnableTrustAnchorInput{
		// TrustAnchorId: *string, // Required
	}

	if len(_rolesanywhereTrustAnchorId) > 0 {
		input.TrustAnchorId = aws.String(_rolesanywhereTrustAnchorId)
	}

	if resp, err := client.EnableTrustAnchor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a certificate revocation list (CRL).
// Required permissions: rolesanywhere:GetCrl .
func rolesanywhere_GetCrl(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.GetCrlInput{
		// CrlId: *string, // Required
	}

	if len(_rolesanywhereCrlId) > 0 {
		input.CrlId = aws.String(_rolesanywhereCrlId)
	}

	if resp, err := client.GetCrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a profile.
// Required permissions: rolesanywhere:GetProfile .
func rolesanywhere_GetProfile(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.GetProfileInput{
		// ProfileId: *string, // Required
	}

	if len(_rolesanywhereProfileId) > 0 {
		input.ProfileId = aws.String(_rolesanywhereProfileId)
	}

	if resp, err := client.GetProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a subject, which associates a certificate identity with authentication
// attempts. The subject stores auditing information such as the status of the last
// authentication attempt, the certificate data used in the attempt, and the last
// time the associated identity attempted authentication.
//
// Required permissions: rolesanywhere:GetSubject .
func rolesanywhere_GetSubject(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.GetSubjectInput{
		// SubjectId: *string, // Required
	}

	if len(_rolesanywhereSubjectId) > 0 {
		input.SubjectId = aws.String(_rolesanywhereSubjectId)
	}

	if resp, err := client.GetSubject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a trust anchor.
// Required permissions: rolesanywhere:GetTrustAnchor .
func rolesanywhere_GetTrustAnchor(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.GetTrustAnchorInput{
		// TrustAnchorId: *string, // Required
	}

	if len(_rolesanywhereTrustAnchorId) > 0 {
		input.TrustAnchorId = aws.String(_rolesanywhereTrustAnchorId)
	}

	if resp, err := client.GetTrustAnchor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports the certificate revocation list (CRL). A CRL is a list of certificates
// that have been revoked by the issuing certificate Authority (CA).In order to be
// properly imported, a CRL must be in PEM format. IAM Roles Anywhere validates
// against the CRL before issuing credentials.
//
// Required permissions: rolesanywhere:ImportCrl .
func rolesanywhere_ImportCrl(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.ImportCrlInput{
		// CrlData: []byte, // Required
		// Name: *string, // Required
		// TrustAnchorArn: *string, // Required
	}

	if len(_rolesanywhereCrlData) > 0 {
		if err := assignInputField(input, "CrlData", _rolesanywhereCrlData); err != nil {
			log.Errorf("invalid --crl-data: %s", err.Error())
			return
		}
	}
	if len(_rolesanywhereName) > 0 {
		input.Name = aws.String(_rolesanywhereName)
	}
	if len(_rolesanywhereTrustAnchorArn) > 0 {
		input.TrustAnchorArn = aws.String(_rolesanywhereTrustAnchorArn)
	}
	if len(_rolesanywhereEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _rolesanywhereEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_rolesanywhereTags) > 0 {
		if err := assignInputField(input, "Tags", _rolesanywhereTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportCrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all certificate revocation lists (CRL) in the authenticated account and
// Amazon Web Services Region.
//
// Required permissions: rolesanywhere:ListCrls .
func rolesanywhere_ListCrls(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.ListCrlsInput{}

	if len(_rolesanywhereNextToken) > 0 {
		input.NextToken = aws.String(_rolesanywhereNextToken)
	}
	if len(_rolesanywherePageSize) > 0 {
		if err := assignInputField(input, "PageSize", _rolesanywherePageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCrls(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rolesanywhere.ListCrlsOutput
	p := rolesanywhere.NewListCrlsPaginator(client, input)
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

// Lists all profiles in the authenticated account and Amazon Web Services Region.
// Required permissions: rolesanywhere:ListProfiles .
func rolesanywhere_ListProfiles(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.ListProfilesInput{}

	if len(_rolesanywhereNextToken) > 0 {
		input.NextToken = aws.String(_rolesanywhereNextToken)
	}
	if len(_rolesanywherePageSize) > 0 {
		if err := assignInputField(input, "PageSize", _rolesanywherePageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rolesanywhere.ListProfilesOutput
	p := rolesanywhere.NewListProfilesPaginator(client, input)
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

// Lists the subjects in the authenticated account and Amazon Web Services Region.
// Required permissions: rolesanywhere:ListSubjects .
func rolesanywhere_ListSubjects(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.ListSubjectsInput{}

	if len(_rolesanywhereNextToken) > 0 {
		input.NextToken = aws.String(_rolesanywhereNextToken)
	}
	if len(_rolesanywherePageSize) > 0 {
		if err := assignInputField(input, "PageSize", _rolesanywherePageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSubjects(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rolesanywhere.ListSubjectsOutput
	p := rolesanywhere.NewListSubjectsPaginator(client, input)
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

// Lists the tags attached to the resource.
// Required permissions: rolesanywhere:ListTagsForResource .
func rolesanywhere_ListTagsForResource(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_rolesanywhereResourceArn) > 0 {
		input.ResourceArn = aws.String(_rolesanywhereResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the trust anchors in the authenticated account and Amazon Web Services
// Region.
//
// Required permissions: rolesanywhere:ListTrustAnchors .
func rolesanywhere_ListTrustAnchors(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.ListTrustAnchorsInput{}

	if len(_rolesanywhereNextToken) > 0 {
		input.NextToken = aws.String(_rolesanywhereNextToken)
	}
	if len(_rolesanywherePageSize) > 0 {
		if err := assignInputField(input, "PageSize", _rolesanywherePageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTrustAnchors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rolesanywhere.ListTrustAnchorsOutput
	p := rolesanywhere.NewListTrustAnchorsPaginator(client, input)
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

// Put an entry in the attribute mapping rules that will be enforced by a given
// profile. A mapping specifies a certificate field and one or more specifiers that
// have contextual meanings.
func rolesanywhere_PutAttributeMapping(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.PutAttributeMappingInput{
		// CertificateField: types.CertificateField, // Required
		// MappingRules: []types.MappingRule, // Required
		// ProfileId: *string, // Required
	}

	if len(_rolesanywhereCertificateField) > 0 {
		if err := assignInputField(input, "CertificateField", _rolesanywhereCertificateField); err != nil {
			log.Errorf("invalid --certificate-field: %s", err.Error())
			return
		}
	}
	if len(_rolesanywhereMappingRules) > 0 {
		if err := assignInputField(input, "MappingRules", _rolesanywhereMappingRules); err != nil {
			log.Errorf("invalid --mapping-rules: %s", err.Error())
			return
		}
	}
	if len(_rolesanywhereProfileId) > 0 {
		input.ProfileId = aws.String(_rolesanywhereProfileId)
	}

	if resp, err := client.PutAttributeMapping(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches a list of notification settings to a trust anchor.
// A notification setting includes information such as event name, threshold,
// status of the notification setting, and the channel to notify.
//
// Required permissions: rolesanywhere:PutNotificationSettings .
func rolesanywhere_PutNotificationSettings(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.PutNotificationSettingsInput{
		// NotificationSettings: []types.NotificationSetting, // Required
		// TrustAnchorId: *string, // Required
	}

	if len(_rolesanywhereNotificationSettings) > 0 {
		if err := assignInputField(input, "NotificationSettings", _rolesanywhereNotificationSettings); err != nil {
			log.Errorf("invalid --notification-settings: %s", err.Error())
			return
		}
	}
	if len(_rolesanywhereTrustAnchorId) > 0 {
		input.TrustAnchorId = aws.String(_rolesanywhereTrustAnchorId)
	}

	if resp, err := client.PutNotificationSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Resets the custom notification setting to IAM Roles Anywhere default setting.
// Required permissions: rolesanywhere:ResetNotificationSettings .
func rolesanywhere_ResetNotificationSettings(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.ResetNotificationSettingsInput{
		// NotificationSettingKeys: []types.NotificationSettingKey, // Required
		// TrustAnchorId: *string, // Required
	}

	if len(_rolesanywhereNotificationSettingKeys) > 0 {
		if err := assignInputField(input, "NotificationSettingKeys", _rolesanywhereNotificationSettingKeys); err != nil {
			log.Errorf("invalid --notification-setting-keys: %s", err.Error())
			return
		}
	}
	if len(_rolesanywhereTrustAnchorId) > 0 {
		input.TrustAnchorId = aws.String(_rolesanywhereTrustAnchorId)
	}

	if resp, err := client.ResetNotificationSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches tags to a resource.
// Required permissions: rolesanywhere:TagResource .
func rolesanywhere_TagResource(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_rolesanywhereResourceArn) > 0 {
		input.ResourceArn = aws.String(_rolesanywhereResourceArn)
	}
	if len(_rolesanywhereTags) > 0 {
		if err := assignInputField(input, "Tags", _rolesanywhereTags); err != nil {
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

// Removes tags from the resource.
// Required permissions: rolesanywhere:UntagResource .
func rolesanywhere_UntagResource(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_rolesanywhereResourceArn) > 0 {
		input.ResourceArn = aws.String(_rolesanywhereResourceArn)
	}
	if len(_rolesanywhereTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _rolesanywhereTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the certificate revocation list (CRL). A CRL is a list of certificates
// that have been revoked by the issuing certificate authority (CA). IAM Roles
// Anywhere validates against the CRL before issuing credentials.
//
// Required permissions: rolesanywhere:UpdateCrl .
func rolesanywhere_UpdateCrl(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.UpdateCrlInput{
		// CrlId: *string, // Required
	}

	if len(_rolesanywhereCrlId) > 0 {
		input.CrlId = aws.String(_rolesanywhereCrlId)
	}
	if len(_rolesanywhereCrlData) > 0 {
		if err := assignInputField(input, "CrlData", _rolesanywhereCrlData); err != nil {
			log.Errorf("invalid --crl-data: %s", err.Error())
			return
		}
	}
	if len(_rolesanywhereName) > 0 {
		input.Name = aws.String(_rolesanywhereName)
	}

	if resp, err := client.UpdateCrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a profile, a list of the roles that IAM Roles Anywhere service is
// trusted to assume. You use profiles to intersect permissions with IAM managed
// policies.
//
// Required permissions: rolesanywhere:UpdateProfile .
func rolesanywhere_UpdateProfile(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.UpdateProfileInput{
		// ProfileId: *string, // Required
	}

	if len(_rolesanywhereProfileId) > 0 {
		input.ProfileId = aws.String(_rolesanywhereProfileId)
	}
	if len(_rolesanywhereAcceptRoleSessionName) > 0 {
		if err := assignInputField(input, "AcceptRoleSessionName", _rolesanywhereAcceptRoleSessionName); err != nil {
			log.Errorf("invalid --accept-role-session-name: %s", err.Error())
			return
		}
	}
	if len(_rolesanywhereDurationSeconds) > 0 {
		if err := assignInputField(input, "DurationSeconds", _rolesanywhereDurationSeconds); err != nil {
			log.Errorf("invalid --duration-seconds: %s", err.Error())
			return
		}
	}
	if len(_rolesanywhereManagedPolicyArns) > 0 {
		input.ManagedPolicyArns = append([]string(nil), _rolesanywhereManagedPolicyArns...)
	}
	if len(_rolesanywhereName) > 0 {
		input.Name = aws.String(_rolesanywhereName)
	}
	if len(_rolesanywhereRoleArns) > 0 {
		input.RoleArns = append([]string(nil), _rolesanywhereRoleArns...)
	}
	if len(_rolesanywhereSessionPolicy) > 0 {
		input.SessionPolicy = aws.String(_rolesanywhereSessionPolicy)
	}

	if resp, err := client.UpdateProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a trust anchor. You establish trust between IAM Roles Anywhere and your
// certificate authority (CA) by configuring a trust anchor. You can define a trust
// anchor as a reference to an Private Certificate Authority (Private CA) or by
// uploading a CA certificate. Your Amazon Web Services workloads can authenticate
// with the trust anchor using certificates issued by the CA in exchange for
// temporary Amazon Web Services credentials.
//
// Required permissions: rolesanywhere:UpdateTrustAnchor .
func rolesanywhere_UpdateTrustAnchor(cfg aws.Config, client *rolesanywhere.Client) {
	input := &rolesanywhere.UpdateTrustAnchorInput{
		// TrustAnchorId: *string, // Required
	}

	if len(_rolesanywhereTrustAnchorId) > 0 {
		input.TrustAnchorId = aws.String(_rolesanywhereTrustAnchorId)
	}
	if len(_rolesanywhereName) > 0 {
		input.Name = aws.String(_rolesanywhereName)
	}
	if len(_rolesanywhereSource) > 0 {
		if err := assignInputField(input, "Source", _rolesanywhereSource); err != nil {
			log.Errorf("invalid --source: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTrustAnchor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_rolesanywhereCmd)
	_rolesanywhereCmd.Flags().SortFlags = false

	_rolesanywhereCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_rolesanywhereCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_rolesanywhereCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_rolesanywhereCmd.Flags().StringVarP(&_rolesanywhereAcceptRoleSessionName, "accept-role-session-name", "", "", "Accept Role Session Name")
	_rolesanywhereCmd.Flags().StringVarP(&_rolesanywhereCertificateField, "certificate-field", "", "", "Certificate Field")
	_rolesanywhereCmd.Flags().StringVarP(&_rolesanywhereCrlData, "crl-data", "", "", "Crl Data")
	_rolesanywhereCmd.Flags().StringVarP(&_rolesanywhereCrlId, "crl-id", "", "", "Crl ID")
	_rolesanywhereCmd.Flags().StringVarP(&_rolesanywhereDurationSeconds, "duration-seconds", "", "", "Duration Seconds")
	_rolesanywhereCmd.Flags().StringVarP(&_rolesanywhereEnabled, "enabled", "", "", "Enabled")
	_rolesanywhereCmd.Flags().StringSliceVarP(&_rolesanywhereManagedPolicyArns, "managed-policy-arns", "", nil, "Managed Policy Arns")
	_rolesanywhereCmd.Flags().StringVarP(&_rolesanywhereMappingRules, "mapping-rules", "", "", "Mapping Rules")
	_rolesanywhereCmd.Flags().StringVarP(&_rolesanywhereName, "name", "", "", "Name")
	_rolesanywhereCmd.Flags().StringVarP(&_rolesanywhereNextToken, "next-token", "", "", "Next Token")
	_rolesanywhereCmd.Flags().StringVarP(&_rolesanywhereNotificationSettingKeys, "notification-setting-keys", "", "", "Notification Setting Keys")
	_rolesanywhereCmd.Flags().StringVarP(&_rolesanywhereNotificationSettings, "notification-settings", "", "", "Notification Settings")
	_rolesanywhereCmd.Flags().StringVarP(&_rolesanywherePageSize, "page-size", "", "", "Page Size")
	_rolesanywhereCmd.Flags().StringVarP(&_rolesanywhereProfileId, "profile-id", "", "", "Profile ID")
	_rolesanywhereCmd.Flags().StringVarP(&_rolesanywhereRequireInstanceProperties, "require-instance-properties", "", "", "Require Instance Properties")
	_rolesanywhereCmd.Flags().StringVarP(&_rolesanywhereResourceArn, "resource-arn", "", "", "Resource ARN")
	_rolesanywhereCmd.Flags().StringSliceVarP(&_rolesanywhereRoleArns, "role-arns", "", nil, "Role Arns")
	_rolesanywhereCmd.Flags().StringVarP(&_rolesanywhereSessionPolicy, "session-policy", "", "", "Session Policy")
	_rolesanywhereCmd.Flags().StringVarP(&_rolesanywhereSource, "source", "", "", "Source")
	_rolesanywhereCmd.Flags().StringSliceVarP(&_rolesanywhereSpecifiers, "specifiers", "", nil, "Specifiers")
	_rolesanywhereCmd.Flags().StringVarP(&_rolesanywhereSubjectId, "subject-id", "", "", "Subject ID")
	_rolesanywhereCmd.Flags().StringSliceVarP(&_rolesanywhereTagKeys, "tag-keys", "", nil, "Tag Keys")
	_rolesanywhereCmd.Flags().StringVarP(&_rolesanywhereTags, "tags", "", "", "Tags")
	_rolesanywhereCmd.Flags().StringVarP(&_rolesanywhereTrustAnchorArn, "trust-anchor-arn", "", "", "Trust Anchor ARN")
	_rolesanywhereCmd.Flags().StringVarP(&_rolesanywhereTrustAnchorId, "trust-anchor-id", "", "", "Trust Anchor ID")

	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywhereCreateProfile, "create-profile", "", false, "Create Profile")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywhereCreateTrustAnchor, "create-trust-anchor", "", false, "Create Trust Anchor")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywhereDeleteAttributeMapping, "delete-attribute-mapping", "", false, "Delete Attribute Mapping")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywhereDeleteCrl, "delete-crl", "", false, "Delete Crl")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywhereDeleteProfile, "delete-profile", "", false, "Delete Profile")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywhereDeleteTrustAnchor, "delete-trust-anchor", "", false, "Delete Trust Anchor")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywhereDisableCrl, "disable-crl", "", false, "Disable Crl")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywhereDisableProfile, "disable-profile", "", false, "Disable Profile")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywhereDisableTrustAnchor, "disable-trust-anchor", "", false, "Disable Trust Anchor")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywhereEnableCrl, "enable-crl", "", false, "Enable Crl")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywhereEnableProfile, "enable-profile", "", false, "Enable Profile")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywhereEnableTrustAnchor, "enable-trust-anchor", "", false, "Enable Trust Anchor")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywhereGetCrl, "get-crl", "", false, "Get Crl")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywhereGetProfile, "get-profile", "", false, "Get Profile")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywhereGetSubject, "get-subject", "", false, "Get Subject")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywhereGetTrustAnchor, "get-trust-anchor", "", false, "Get Trust Anchor")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywhereImportCrl, "import-crl", "", false, "Import Crl")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywhereListCrls, "list-crls", "", false, "List Crls")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywhereListProfiles, "list-profiles", "", false, "List Profiles")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywhereListSubjects, "list-subjects", "", false, "List Subjects")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywhereListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywhereListTrustAnchors, "list-trust-anchors", "", false, "List Trust Anchors")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywherePutAttributeMapping, "put-attribute-mapping", "", false, "Put Attribute Mapping")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywherePutNotificationSettings, "put-notification-settings", "", false, "Put Notification Settings")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywhereResetNotificationSettings, "reset-notification-settings", "", false, "Reset Notification Settings")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywhereTagResource, "tag-resource", "", false, "Tag Resource")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywhereUntagResource, "untag-resource", "", false, "Untag Resource")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywhereUpdateCrl, "update-crl", "", false, "Update Crl")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywhereUpdateProfile, "update-profile", "", false, "Update Profile")
	_rolesanywhereCmd.Flags().BoolVarP(&_rolesanywhereUpdateTrustAnchor, "update-trust-anchor", "", false, "Update Trust Anchor")

}
