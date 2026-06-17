package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// wellarchitectedCmd represents the wellarchitected command
var _wellarchitectedCmd = &cobra.Command{
	Use:   "wellarchitected",
	Short: "AWS wellarchitected CLI",
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
		client := wellarchitected.NewFromConfig(cfg)
		if _wellarchitectedAssociateLenses {
			wellarchitected_AssociateLenses(cfg, client)
			return
		}
		if _wellarchitectedAssociateProfiles {
			wellarchitected_AssociateProfiles(cfg, client)
			return
		}
		if _wellarchitectedCreateLensShare {
			wellarchitected_CreateLensShare(cfg, client)
			return
		}
		if _wellarchitectedCreateLensVersion {
			wellarchitected_CreateLensVersion(cfg, client)
			return
		}
		if _wellarchitectedCreateMilestone {
			wellarchitected_CreateMilestone(cfg, client)
			return
		}
		if _wellarchitectedCreateProfile {
			wellarchitected_CreateProfile(cfg, client)
			return
		}
		if _wellarchitectedCreateProfileShare {
			wellarchitected_CreateProfileShare(cfg, client)
			return
		}
		if _wellarchitectedCreateReviewTemplate {
			wellarchitected_CreateReviewTemplate(cfg, client)
			return
		}
		if _wellarchitectedCreateTemplateShare {
			wellarchitected_CreateTemplateShare(cfg, client)
			return
		}
		if _wellarchitectedCreateWorkload {
			wellarchitected_CreateWorkload(cfg, client)
			return
		}
		if _wellarchitectedCreateWorkloadShare {
			wellarchitected_CreateWorkloadShare(cfg, client)
			return
		}
		if _wellarchitectedDeleteLens {
			wellarchitected_DeleteLens(cfg, client)
			return
		}
		if _wellarchitectedDeleteLensShare {
			wellarchitected_DeleteLensShare(cfg, client)
			return
		}
		if _wellarchitectedDeleteProfile {
			wellarchitected_DeleteProfile(cfg, client)
			return
		}
		if _wellarchitectedDeleteProfileShare {
			wellarchitected_DeleteProfileShare(cfg, client)
			return
		}
		if _wellarchitectedDeleteReviewTemplate {
			wellarchitected_DeleteReviewTemplate(cfg, client)
			return
		}
		if _wellarchitectedDeleteTemplateShare {
			wellarchitected_DeleteTemplateShare(cfg, client)
			return
		}
		if _wellarchitectedDeleteWorkload {
			wellarchitected_DeleteWorkload(cfg, client)
			return
		}
		if _wellarchitectedDeleteWorkloadShare {
			wellarchitected_DeleteWorkloadShare(cfg, client)
			return
		}
		if _wellarchitectedDisassociateLenses {
			wellarchitected_DisassociateLenses(cfg, client)
			return
		}
		if _wellarchitectedDisassociateProfiles {
			wellarchitected_DisassociateProfiles(cfg, client)
			return
		}
		if _wellarchitectedExportLens {
			wellarchitected_ExportLens(cfg, client)
			return
		}
		if _wellarchitectedGetAnswer {
			wellarchitected_GetAnswer(cfg, client)
			return
		}
		if _wellarchitectedGetConsolidatedReport {
			wellarchitected_GetConsolidatedReport(cfg, client)
			return
		}
		if _wellarchitectedGetGlobalSettings {
			wellarchitected_GetGlobalSettings(cfg, client)
			return
		}
		if _wellarchitectedGetLens {
			wellarchitected_GetLens(cfg, client)
			return
		}
		if _wellarchitectedGetLensReview {
			wellarchitected_GetLensReview(cfg, client)
			return
		}
		if _wellarchitectedGetLensReviewReport {
			wellarchitected_GetLensReviewReport(cfg, client)
			return
		}
		if _wellarchitectedGetLensVersionDifference {
			wellarchitected_GetLensVersionDifference(cfg, client)
			return
		}
		if _wellarchitectedGetMilestone {
			wellarchitected_GetMilestone(cfg, client)
			return
		}
		if _wellarchitectedGetProfile {
			wellarchitected_GetProfile(cfg, client)
			return
		}
		if _wellarchitectedGetProfileTemplate {
			wellarchitected_GetProfileTemplate(cfg, client)
			return
		}
		if _wellarchitectedGetReviewTemplate {
			wellarchitected_GetReviewTemplate(cfg, client)
			return
		}
		if _wellarchitectedGetReviewTemplateAnswer {
			wellarchitected_GetReviewTemplateAnswer(cfg, client)
			return
		}
		if _wellarchitectedGetReviewTemplateLensReview {
			wellarchitected_GetReviewTemplateLensReview(cfg, client)
			return
		}
		if _wellarchitectedGetWorkload {
			wellarchitected_GetWorkload(cfg, client)
			return
		}
		if _wellarchitectedImportLens {
			wellarchitected_ImportLens(cfg, client)
			return
		}
		if _wellarchitectedListAnswers {
			wellarchitected_ListAnswers(cfg, client)
			return
		}
		if _wellarchitectedListCheckDetails {
			wellarchitected_ListCheckDetails(cfg, client)
			return
		}
		if _wellarchitectedListCheckSummaries {
			wellarchitected_ListCheckSummaries(cfg, client)
			return
		}
		if _wellarchitectedListLensReviewImprovements {
			wellarchitected_ListLensReviewImprovements(cfg, client)
			return
		}
		if _wellarchitectedListLensReviews {
			wellarchitected_ListLensReviews(cfg, client)
			return
		}
		if _wellarchitectedListLensShares {
			wellarchitected_ListLensShares(cfg, client)
			return
		}
		if _wellarchitectedListLenses {
			wellarchitected_ListLenses(cfg, client)
			return
		}
		if _wellarchitectedListMilestones {
			wellarchitected_ListMilestones(cfg, client)
			return
		}
		if _wellarchitectedListNotifications {
			wellarchitected_ListNotifications(cfg, client)
			return
		}
		if _wellarchitectedListProfileNotifications {
			wellarchitected_ListProfileNotifications(cfg, client)
			return
		}
		if _wellarchitectedListProfileShares {
			wellarchitected_ListProfileShares(cfg, client)
			return
		}
		if _wellarchitectedListProfiles {
			wellarchitected_ListProfiles(cfg, client)
			return
		}
		if _wellarchitectedListReviewTemplateAnswers {
			wellarchitected_ListReviewTemplateAnswers(cfg, client)
			return
		}
		if _wellarchitectedListReviewTemplates {
			wellarchitected_ListReviewTemplates(cfg, client)
			return
		}
		if _wellarchitectedListShareInvitations {
			wellarchitected_ListShareInvitations(cfg, client)
			return
		}
		if _wellarchitectedListTagsForResource {
			wellarchitected_ListTagsForResource(cfg, client)
			return
		}
		if _wellarchitectedListTemplateShares {
			wellarchitected_ListTemplateShares(cfg, client)
			return
		}
		if _wellarchitectedListWorkloadShares {
			wellarchitected_ListWorkloadShares(cfg, client)
			return
		}
		if _wellarchitectedListWorkloads {
			wellarchitected_ListWorkloads(cfg, client)
			return
		}
		if _wellarchitectedTagResource {
			wellarchitected_TagResource(cfg, client)
			return
		}
		if _wellarchitectedUntagResource {
			wellarchitected_UntagResource(cfg, client)
			return
		}
		if _wellarchitectedUpdateAnswer {
			wellarchitected_UpdateAnswer(cfg, client)
			return
		}
		if _wellarchitectedUpdateGlobalSettings {
			wellarchitected_UpdateGlobalSettings(cfg, client)
			return
		}
		if _wellarchitectedUpdateIntegration {
			wellarchitected_UpdateIntegration(cfg, client)
			return
		}
		if _wellarchitectedUpdateLensReview {
			wellarchitected_UpdateLensReview(cfg, client)
			return
		}
		if _wellarchitectedUpdateProfile {
			wellarchitected_UpdateProfile(cfg, client)
			return
		}
		if _wellarchitectedUpdateReviewTemplate {
			wellarchitected_UpdateReviewTemplate(cfg, client)
			return
		}
		if _wellarchitectedUpdateReviewTemplateAnswer {
			wellarchitected_UpdateReviewTemplateAnswer(cfg, client)
			return
		}
		if _wellarchitectedUpdateReviewTemplateLensReview {
			wellarchitected_UpdateReviewTemplateLensReview(cfg, client)
			return
		}
		if _wellarchitectedUpdateShareInvitation {
			wellarchitected_UpdateShareInvitation(cfg, client)
			return
		}
		if _wellarchitectedUpdateWorkload {
			wellarchitected_UpdateWorkload(cfg, client)
			return
		}
		if _wellarchitectedUpdateWorkloadShare {
			wellarchitected_UpdateWorkloadShare(cfg, client)
			return
		}
		if _wellarchitectedUpgradeLensReview {
			wellarchitected_UpgradeLensReview(cfg, client)
			return
		}
		if _wellarchitectedUpgradeProfileVersion {
			wellarchitected_UpgradeProfileVersion(cfg, client)
			return
		}
		if _wellarchitectedUpgradeReviewTemplateLensReview {
			wellarchitected_UpgradeReviewTemplateLensReview(cfg, client)
			return
		}

	},
}

var (
	_wellarchitectedAssociateLenses                 bool
	_wellarchitectedAssociateProfiles               bool
	_wellarchitectedCreateLensShare                 bool
	_wellarchitectedCreateLensVersion               bool
	_wellarchitectedCreateMilestone                 bool
	_wellarchitectedCreateProfile                   bool
	_wellarchitectedCreateProfileShare              bool
	_wellarchitectedCreateReviewTemplate            bool
	_wellarchitectedCreateTemplateShare             bool
	_wellarchitectedCreateWorkload                  bool
	_wellarchitectedCreateWorkloadShare             bool
	_wellarchitectedDeleteLens                      bool
	_wellarchitectedDeleteLensShare                 bool
	_wellarchitectedDeleteProfile                   bool
	_wellarchitectedDeleteProfileShare              bool
	_wellarchitectedDeleteReviewTemplate            bool
	_wellarchitectedDeleteTemplateShare             bool
	_wellarchitectedDeleteWorkload                  bool
	_wellarchitectedDeleteWorkloadShare             bool
	_wellarchitectedDisassociateLenses              bool
	_wellarchitectedDisassociateProfiles            bool
	_wellarchitectedExportLens                      bool
	_wellarchitectedGetAnswer                       bool
	_wellarchitectedGetConsolidatedReport           bool
	_wellarchitectedGetGlobalSettings               bool
	_wellarchitectedGetLens                         bool
	_wellarchitectedGetLensReview                   bool
	_wellarchitectedGetLensReviewReport             bool
	_wellarchitectedGetLensVersionDifference        bool
	_wellarchitectedGetMilestone                    bool
	_wellarchitectedGetProfile                      bool
	_wellarchitectedGetProfileTemplate              bool
	_wellarchitectedGetReviewTemplate               bool
	_wellarchitectedGetReviewTemplateAnswer         bool
	_wellarchitectedGetReviewTemplateLensReview     bool
	_wellarchitectedGetWorkload                     bool
	_wellarchitectedImportLens                      bool
	_wellarchitectedListAnswers                     bool
	_wellarchitectedListCheckDetails                bool
	_wellarchitectedListCheckSummaries              bool
	_wellarchitectedListLensReviewImprovements      bool
	_wellarchitectedListLensReviews                 bool
	_wellarchitectedListLensShares                  bool
	_wellarchitectedListLenses                      bool
	_wellarchitectedListMilestones                  bool
	_wellarchitectedListNotifications               bool
	_wellarchitectedListProfileNotifications        bool
	_wellarchitectedListProfileShares               bool
	_wellarchitectedListProfiles                    bool
	_wellarchitectedListReviewTemplateAnswers       bool
	_wellarchitectedListReviewTemplates             bool
	_wellarchitectedListShareInvitations            bool
	_wellarchitectedListTagsForResource             bool
	_wellarchitectedListTemplateShares              bool
	_wellarchitectedListWorkloadShares              bool
	_wellarchitectedListWorkloads                   bool
	_wellarchitectedTagResource                     bool
	_wellarchitectedUntagResource                   bool
	_wellarchitectedUpdateAnswer                    bool
	_wellarchitectedUpdateGlobalSettings            bool
	_wellarchitectedUpdateIntegration               bool
	_wellarchitectedUpdateLensReview                bool
	_wellarchitectedUpdateProfile                   bool
	_wellarchitectedUpdateReviewTemplate            bool
	_wellarchitectedUpdateReviewTemplateAnswer      bool
	_wellarchitectedUpdateReviewTemplateLensReview  bool
	_wellarchitectedUpdateShareInvitation           bool
	_wellarchitectedUpdateWorkload                  bool
	_wellarchitectedUpdateWorkloadShare             bool
	_wellarchitectedUpgradeLensReview               bool
	_wellarchitectedUpgradeProfileVersion           bool
	_wellarchitectedUpgradeReviewTemplateLensReview bool

	_wellarchitectedAccountIds                      []string
	_wellarchitectedApplications                    []string
	_wellarchitectedArchitecturalDesign             string
	_wellarchitectedAwsRegions                      []string
	_wellarchitectedBaseLensVersion                 string
	_wellarchitectedChoiceId                        string
	_wellarchitectedChoiceUpdates                   string
	_wellarchitectedClientRequestToken              string
	_wellarchitectedDescription                     string
	_wellarchitectedDiscoveryConfig                 string
	_wellarchitectedDiscoveryIntegrationStatus      string
	_wellarchitectedEnvironment                     string
	_wellarchitectedFormat                          string
	_wellarchitectedImprovementStatus               string
	_wellarchitectedIncludeSharedResources          string
	_wellarchitectedIndustry                        string
	_wellarchitectedIndustryType                    string
	_wellarchitectedIntegratingService              string
	_wellarchitectedIsApplicable                    string
	_wellarchitectedIsMajorVersion                  string
	_wellarchitectedIsReviewOwnerUpdateAcknowledged string
	_wellarchitectedJiraConfiguration               string
	_wellarchitectedJSONString                      string
	_wellarchitectedLensAlias                       string
	_wellarchitectedLensAliases                     []string
	_wellarchitectedLensArn                         string
	_wellarchitectedLensName                        string
	_wellarchitectedLensNamePrefix                  string
	_wellarchitectedLensNotes                       string
	_wellarchitectedLensStatus                      string
	_wellarchitectedLensType                        string
	_wellarchitectedLensVersion                     string
	_wellarchitectedLenses                          []string
	_wellarchitectedLensesToAssociate               []string
	_wellarchitectedLensesToDisassociate            []string
	_wellarchitectedMaxResults                      string
	_wellarchitectedMilestoneName                   string
	_wellarchitectedMilestoneNumber                 string
	_wellarchitectedNextToken                       string
	_wellarchitectedNonAwsRegions                   []string
	_wellarchitectedNotes                           string
	_wellarchitectedOrganizationSharingStatus       string
	_wellarchitectedPermissionType                  string
	_wellarchitectedPillarId                        string
	_wellarchitectedPillarNotes                     string
	_wellarchitectedPillarPriorities                []string
	_wellarchitectedProfileArn                      string
	_wellarchitectedProfileArns                     []string
	_wellarchitectedProfileDescription              string
	_wellarchitectedProfileName                     string
	_wellarchitectedProfileNamePrefix               string
	_wellarchitectedProfileOwnerType                string
	_wellarchitectedProfileQuestions                string
	_wellarchitectedProfileVersion                  string
	_wellarchitectedQuestionId                      string
	_wellarchitectedQuestionPriority                string
	_wellarchitectedReason                          string
	_wellarchitectedResourceArn                     string
	_wellarchitectedReviewOwner                     string
	_wellarchitectedReviewTemplateArns              []string
	_wellarchitectedSelectedChoices                 []string
	_wellarchitectedShareId                         string
	_wellarchitectedShareInvitationAction           string
	_wellarchitectedShareInvitationId               string
	_wellarchitectedShareResourceType               string
	_wellarchitectedSharedWith                      string
	_wellarchitectedSharedWithPrefix                string
	_wellarchitectedStatus                          string
	_wellarchitectedTagKeys                         []string
	_wellarchitectedTags                            string
	_wellarchitectedTargetLensVersion               string
	_wellarchitectedTemplateArn                     string
	_wellarchitectedTemplateName                    string
	_wellarchitectedTemplateNamePrefix              string
	_wellarchitectedWorkloadArn                     string
	_wellarchitectedWorkloadId                      string
	_wellarchitectedWorkloadName                    string
	_wellarchitectedWorkloadNamePrefix              string
)

// Associate a lens to a workload.
// Up to 10 lenses can be associated with a workload in a single API operation. A
// maximum of 20 lenses can be associated with a workload.
//
// # Disclaimer
//
// By accessing and/or applying custom lenses created by another Amazon Web
// Services user or account, you acknowledge that custom lenses created by other
// users and shared with you are Third Party Content as defined in the Amazon Web
// Services Customer Agreement.
func wellarchitected_AssociateLenses(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.AssociateLensesInput{
		// LensAliases: []string, // Required
		// WorkloadId: *string, // Required
	}

	if len(_wellarchitectedLensAliases) > 0 {
		input.LensAliases = append([]string(nil), _wellarchitectedLensAliases...)
	}
	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}

	if resp, err := client.AssociateLenses(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associate a profile with a workload.
func wellarchitected_AssociateProfiles(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.AssociateProfilesInput{
		// ProfileArns: []string, // Required
		// WorkloadId: *string, // Required
	}

	if len(_wellarchitectedProfileArns) > 0 {
		input.ProfileArns = append([]string(nil), _wellarchitectedProfileArns...)
	}
	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}

	if resp, err := client.AssociateProfiles(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a lens share.
// The owner of a lens can share it with other Amazon Web Services accounts,
// users, an organization, and organizational units (OUs) in the same Amazon Web
// Services Region. Lenses provided by Amazon Web Services (Amazon Web Services
// Official Content) cannot be shared.
//
// Shared access to a lens is not removed until the lens invitation is deleted.
//
// If you share a lens with an organization or OU, all accounts in the
// organization or OU are granted access to the lens.
//
// For more information, see [Sharing a custom lens] in the Well-Architected Tool User Guide.
//
// # Disclaimer
//
// By sharing your custom lenses with other Amazon Web Services accounts, you
// acknowledge that Amazon Web Services will make your custom lenses available to
// those other accounts. Those other accounts may continue to access and use your
// shared custom lenses even if you delete the custom lenses from your own Amazon
// Web Services account or terminate your Amazon Web Services account.
//
// [Sharing a custom lens]: https://docs.aws.amazon.com/wellarchitected/latest/userguide/lenses-sharing.html
func wellarchitected_CreateLensShare(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.CreateLensShareInput{
		// ClientRequestToken: *string, // Required
		// LensAlias: *string, // Required
		// SharedWith: *string, // Required
	}

	if len(_wellarchitectedClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_wellarchitectedClientRequestToken)
	}
	if len(_wellarchitectedLensAlias) > 0 {
		input.LensAlias = aws.String(_wellarchitectedLensAlias)
	}
	if len(_wellarchitectedSharedWith) > 0 {
		input.SharedWith = aws.String(_wellarchitectedSharedWith)
	}

	if resp, err := client.CreateLensShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a new lens version.
// A lens can have up to 100 versions.
//
// Use this operation to publish a new lens version after you have imported a
// lens. The LensAlias is used to identify the lens to be published. The owner of
// a lens can share the lens with other Amazon Web Services accounts and users in
// the same Amazon Web Services Region. Only the owner of a lens can delete it.
func wellarchitected_CreateLensVersion(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.CreateLensVersionInput{
		// ClientRequestToken: *string, // Required
		// LensAlias: *string, // Required
		// LensVersion: *string, // Required
	}

	if len(_wellarchitectedClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_wellarchitectedClientRequestToken)
	}
	if len(_wellarchitectedLensAlias) > 0 {
		input.LensAlias = aws.String(_wellarchitectedLensAlias)
	}
	if len(_wellarchitectedLensVersion) > 0 {
		input.LensVersion = aws.String(_wellarchitectedLensVersion)
	}
	if len(_wellarchitectedIsMajorVersion) > 0 {
		if err := assignInputField(input, "IsMajorVersion", _wellarchitectedIsMajorVersion); err != nil {
			log.Errorf("invalid --is-major-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLensVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a milestone for an existing workload.
func wellarchitected_CreateMilestone(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.CreateMilestoneInput{
		// ClientRequestToken: *string, // Required
		// MilestoneName: *string, // Required
		// WorkloadId: *string, // Required
	}

	if len(_wellarchitectedClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_wellarchitectedClientRequestToken)
	}
	if len(_wellarchitectedMilestoneName) > 0 {
		input.MilestoneName = aws.String(_wellarchitectedMilestoneName)
	}
	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}

	if resp, err := client.CreateMilestone(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a profile.
func wellarchitected_CreateProfile(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.CreateProfileInput{
		// ClientRequestToken: *string, // Required
		// ProfileDescription: *string, // Required
		// ProfileName: *string, // Required
		// ProfileQuestions: []types.ProfileQuestionUpdate, // Required
	}

	if len(_wellarchitectedClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_wellarchitectedClientRequestToken)
	}
	if len(_wellarchitectedProfileDescription) > 0 {
		input.ProfileDescription = aws.String(_wellarchitectedProfileDescription)
	}
	if len(_wellarchitectedProfileName) > 0 {
		input.ProfileName = aws.String(_wellarchitectedProfileName)
	}
	if len(_wellarchitectedProfileQuestions) > 0 {
		if err := assignInputField(input, "ProfileQuestions", _wellarchitectedProfileQuestions); err != nil {
			log.Errorf("invalid --profile-questions: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedTags) > 0 {
		if err := assignInputField(input, "Tags", _wellarchitectedTags); err != nil {
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

// Create a profile share.
func wellarchitected_CreateProfileShare(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.CreateProfileShareInput{
		// ClientRequestToken: *string, // Required
		// ProfileArn: *string, // Required
		// SharedWith: *string, // Required
	}

	if len(_wellarchitectedClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_wellarchitectedClientRequestToken)
	}
	if len(_wellarchitectedProfileArn) > 0 {
		input.ProfileArn = aws.String(_wellarchitectedProfileArn)
	}
	if len(_wellarchitectedSharedWith) > 0 {
		input.SharedWith = aws.String(_wellarchitectedSharedWith)
	}

	if resp, err := client.CreateProfileShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a review template.
// # Disclaimer
//
// Do not include or gather personal identifiable information (PII) of end users
// or other identifiable individuals in or via your review templates. If your
// review template or those shared with you and used in your account do include or
// collect PII you are responsible for: ensuring that the included PII is processed
// in accordance with applicable law, providing adequate privacy notices, and
// obtaining necessary consents for processing such data.
func wellarchitected_CreateReviewTemplate(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.CreateReviewTemplateInput{
		// ClientRequestToken: *string, // Required
		// Description: *string, // Required
		// Lenses: []string, // Required
		// TemplateName: *string, // Required
	}

	if len(_wellarchitectedClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_wellarchitectedClientRequestToken)
	}
	if len(_wellarchitectedDescription) > 0 {
		input.Description = aws.String(_wellarchitectedDescription)
	}
	if len(_wellarchitectedLenses) > 0 {
		input.Lenses = append([]string(nil), _wellarchitectedLenses...)
	}
	if len(_wellarchitectedTemplateName) > 0 {
		input.TemplateName = aws.String(_wellarchitectedTemplateName)
	}
	if len(_wellarchitectedNotes) > 0 {
		input.Notes = aws.String(_wellarchitectedNotes)
	}
	if len(_wellarchitectedTags) > 0 {
		if err := assignInputField(input, "Tags", _wellarchitectedTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateReviewTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a review template share.
// The owner of a review template can share it with other Amazon Web Services
// accounts, users, an organization, and organizational units (OUs) in the same
// Amazon Web Services Region.
//
// Shared access to a review template is not removed until the review template
// share invitation is deleted.
//
// If you share a review template with an organization or OU, all accounts in the
// organization or OU are granted access to the review template.
//
// # Disclaimer
//
// By sharing your review template with other Amazon Web Services accounts, you
// acknowledge that Amazon Web Services will make your review template available to
// those other accounts.
func wellarchitected_CreateTemplateShare(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.CreateTemplateShareInput{
		// ClientRequestToken: *string, // Required
		// SharedWith: *string, // Required
		// TemplateArn: *string, // Required
	}

	if len(_wellarchitectedClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_wellarchitectedClientRequestToken)
	}
	if len(_wellarchitectedSharedWith) > 0 {
		input.SharedWith = aws.String(_wellarchitectedSharedWith)
	}
	if len(_wellarchitectedTemplateArn) > 0 {
		input.TemplateArn = aws.String(_wellarchitectedTemplateArn)
	}

	if resp, err := client.CreateTemplateShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a new workload.
// The owner of a workload can share the workload with other Amazon Web Services
// accounts, users, an organization, and organizational units (OUs) in the same
// Amazon Web Services Region. Only the owner of a workload can delete it.
//
// For more information, see [Defining a Workload] in the Well-Architected Tool User Guide.
//
// Either AwsRegions , NonAwsRegions , or both must be specified when creating a
// workload.
//
// You also must specify ReviewOwner , even though the parameter is listed as not
// being required in the following section.
//
// When creating a workload using a review template, you must have the following
// IAM permissions:
//
// - wellarchitected:GetReviewTemplate
//
// - wellarchitected:GetReviewTemplateAnswer
//
// - wellarchitected:ListReviewTemplateAnswers
//
// - wellarchitected:GetReviewTemplateLensReview
//
// [Defining a Workload]: https://docs.aws.amazon.com/wellarchitected/latest/userguide/define-workload.html
func wellarchitected_CreateWorkload(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.CreateWorkloadInput{
		// ClientRequestToken: *string, // Required
		// Description: *string, // Required
		// Environment: types.WorkloadEnvironment, // Required
		// Lenses: []string, // Required
		// WorkloadName: *string, // Required
	}

	if len(_wellarchitectedClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_wellarchitectedClientRequestToken)
	}
	if len(_wellarchitectedDescription) > 0 {
		input.Description = aws.String(_wellarchitectedDescription)
	}
	if len(_wellarchitectedEnvironment) > 0 {
		if err := assignInputField(input, "Environment", _wellarchitectedEnvironment); err != nil {
			log.Errorf("invalid --environment: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedLenses) > 0 {
		input.Lenses = append([]string(nil), _wellarchitectedLenses...)
	}
	if len(_wellarchitectedWorkloadName) > 0 {
		input.WorkloadName = aws.String(_wellarchitectedWorkloadName)
	}
	if len(_wellarchitectedAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _wellarchitectedAccountIds...)
	}
	if len(_wellarchitectedApplications) > 0 {
		input.Applications = append([]string(nil), _wellarchitectedApplications...)
	}
	if len(_wellarchitectedArchitecturalDesign) > 0 {
		input.ArchitecturalDesign = aws.String(_wellarchitectedArchitecturalDesign)
	}
	if len(_wellarchitectedAwsRegions) > 0 {
		input.AwsRegions = append([]string(nil), _wellarchitectedAwsRegions...)
	}
	if len(_wellarchitectedDiscoveryConfig) > 0 {
		if err := assignInputField(input, "DiscoveryConfig", _wellarchitectedDiscoveryConfig); err != nil {
			log.Errorf("invalid --discovery-config: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedIndustry) > 0 {
		input.Industry = aws.String(_wellarchitectedIndustry)
	}
	if len(_wellarchitectedIndustryType) > 0 {
		input.IndustryType = aws.String(_wellarchitectedIndustryType)
	}
	if len(_wellarchitectedJiraConfiguration) > 0 {
		if err := assignInputField(input, "JiraConfiguration", _wellarchitectedJiraConfiguration); err != nil {
			log.Errorf("invalid --jira-configuration: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedNonAwsRegions) > 0 {
		input.NonAwsRegions = append([]string(nil), _wellarchitectedNonAwsRegions...)
	}
	if len(_wellarchitectedNotes) > 0 {
		input.Notes = aws.String(_wellarchitectedNotes)
	}
	if len(_wellarchitectedPillarPriorities) > 0 {
		input.PillarPriorities = append([]string(nil), _wellarchitectedPillarPriorities...)
	}
	if len(_wellarchitectedProfileArns) > 0 {
		input.ProfileArns = append([]string(nil), _wellarchitectedProfileArns...)
	}
	if len(_wellarchitectedReviewOwner) > 0 {
		input.ReviewOwner = aws.String(_wellarchitectedReviewOwner)
	}
	if len(_wellarchitectedReviewTemplateArns) > 0 {
		input.ReviewTemplateArns = append([]string(nil), _wellarchitectedReviewTemplateArns...)
	}
	if len(_wellarchitectedTags) > 0 {
		if err := assignInputField(input, "Tags", _wellarchitectedTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWorkload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a workload share.
// The owner of a workload can share it with other Amazon Web Services accounts
// and users in the same Amazon Web Services Region. Shared access to a workload is
// not removed until the workload invitation is deleted.
//
// If you share a workload with an organization or OU, all accounts in the
// organization or OU are granted access to the workload.
//
// For more information, see [Sharing a workload] in the Well-Architected Tool User Guide.
//
// [Sharing a workload]: https://docs.aws.amazon.com/wellarchitected/latest/userguide/workloads-sharing.html
func wellarchitected_CreateWorkloadShare(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.CreateWorkloadShareInput{
		// ClientRequestToken: *string, // Required
		// PermissionType: types.PermissionType, // Required
		// SharedWith: *string, // Required
		// WorkloadId: *string, // Required
	}

	if len(_wellarchitectedClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_wellarchitectedClientRequestToken)
	}
	if len(_wellarchitectedPermissionType) > 0 {
		if err := assignInputField(input, "PermissionType", _wellarchitectedPermissionType); err != nil {
			log.Errorf("invalid --permission-type: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedSharedWith) > 0 {
		input.SharedWith = aws.String(_wellarchitectedSharedWith)
	}
	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}

	if resp, err := client.CreateWorkloadShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an existing lens.
// Only the owner of a lens can delete it. After the lens is deleted, Amazon Web
// Services accounts and users that you shared the lens with can continue to use
// it, but they will no longer be able to apply it to new workloads.
//
// # Disclaimer
//
// By sharing your custom lenses with other Amazon Web Services accounts, you
// acknowledge that Amazon Web Services will make your custom lenses available to
// those other accounts. Those other accounts may continue to access and use your
// shared custom lenses even if you delete the custom lenses from your own Amazon
// Web Services account or terminate your Amazon Web Services account.
func wellarchitected_DeleteLens(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.DeleteLensInput{
		// ClientRequestToken: *string, // Required
		// LensAlias: *string, // Required
		// LensStatus: types.LensStatusType, // Required
	}

	if len(_wellarchitectedClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_wellarchitectedClientRequestToken)
	}
	if len(_wellarchitectedLensAlias) > 0 {
		input.LensAlias = aws.String(_wellarchitectedLensAlias)
	}
	if len(_wellarchitectedLensStatus) > 0 {
		if err := assignInputField(input, "LensStatus", _wellarchitectedLensStatus); err != nil {
			log.Errorf("invalid --lens-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteLens(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a lens share.
// After the lens share is deleted, Amazon Web Services accounts, users,
// organizations, and organizational units (OUs) that you shared the lens with can
// continue to use it, but they will no longer be able to apply it to new
// workloads.
//
// # Disclaimer
//
// By sharing your custom lenses with other Amazon Web Services accounts, you
// acknowledge that Amazon Web Services will make your custom lenses available to
// those other accounts. Those other accounts may continue to access and use your
// shared custom lenses even if you delete the custom lenses from your own Amazon
// Web Services account or terminate your Amazon Web Services account.
func wellarchitected_DeleteLensShare(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.DeleteLensShareInput{
		// ClientRequestToken: *string, // Required
		// LensAlias: *string, // Required
		// ShareId: *string, // Required
	}

	if len(_wellarchitectedClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_wellarchitectedClientRequestToken)
	}
	if len(_wellarchitectedLensAlias) > 0 {
		input.LensAlias = aws.String(_wellarchitectedLensAlias)
	}
	if len(_wellarchitectedShareId) > 0 {
		input.ShareId = aws.String(_wellarchitectedShareId)
	}

	if resp, err := client.DeleteLensShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a profile.
// # Disclaimer
//
// By sharing your profile with other Amazon Web Services accounts, you
// acknowledge that Amazon Web Services will make your profile available to those
// other accounts. Those other accounts may continue to access and use your shared
// profile even if you delete the profile from your own Amazon Web Services account
// or terminate your Amazon Web Services account.
func wellarchitected_DeleteProfile(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.DeleteProfileInput{
		// ClientRequestToken: *string, // Required
		// ProfileArn: *string, // Required
	}

	if len(_wellarchitectedClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_wellarchitectedClientRequestToken)
	}
	if len(_wellarchitectedProfileArn) > 0 {
		input.ProfileArn = aws.String(_wellarchitectedProfileArn)
	}

	if resp, err := client.DeleteProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a profile share.
func wellarchitected_DeleteProfileShare(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.DeleteProfileShareInput{
		// ClientRequestToken: *string, // Required
		// ProfileArn: *string, // Required
		// ShareId: *string, // Required
	}

	if len(_wellarchitectedClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_wellarchitectedClientRequestToken)
	}
	if len(_wellarchitectedProfileArn) > 0 {
		input.ProfileArn = aws.String(_wellarchitectedProfileArn)
	}
	if len(_wellarchitectedShareId) > 0 {
		input.ShareId = aws.String(_wellarchitectedShareId)
	}

	if resp, err := client.DeleteProfileShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a review template.
// Only the owner of a review template can delete it.
//
// After the review template is deleted, Amazon Web Services accounts, users,
// organizations, and organizational units (OUs) that you shared the review
// template with will no longer be able to apply it to new workloads.
func wellarchitected_DeleteReviewTemplate(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.DeleteReviewTemplateInput{
		// ClientRequestToken: *string, // Required
		// TemplateArn: *string, // Required
	}

	if len(_wellarchitectedClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_wellarchitectedClientRequestToken)
	}
	if len(_wellarchitectedTemplateArn) > 0 {
		input.TemplateArn = aws.String(_wellarchitectedTemplateArn)
	}

	if resp, err := client.DeleteReviewTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a review template share.
// After the review template share is deleted, Amazon Web Services accounts,
// users, organizations, and organizational units (OUs) that you shared the review
// template with will no longer be able to apply it to new workloads.
func wellarchitected_DeleteTemplateShare(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.DeleteTemplateShareInput{
		// ClientRequestToken: *string, // Required
		// ShareId: *string, // Required
		// TemplateArn: *string, // Required
	}

	if len(_wellarchitectedClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_wellarchitectedClientRequestToken)
	}
	if len(_wellarchitectedShareId) > 0 {
		input.ShareId = aws.String(_wellarchitectedShareId)
	}
	if len(_wellarchitectedTemplateArn) > 0 {
		input.TemplateArn = aws.String(_wellarchitectedTemplateArn)
	}

	if resp, err := client.DeleteTemplateShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an existing workload.
func wellarchitected_DeleteWorkload(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.DeleteWorkloadInput{
		// ClientRequestToken: *string, // Required
		// WorkloadId: *string, // Required
	}

	if len(_wellarchitectedClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_wellarchitectedClientRequestToken)
	}
	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}

	if resp, err := client.DeleteWorkload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a workload share.
func wellarchitected_DeleteWorkloadShare(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.DeleteWorkloadShareInput{
		// ClientRequestToken: *string, // Required
		// ShareId: *string, // Required
		// WorkloadId: *string, // Required
	}

	if len(_wellarchitectedClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_wellarchitectedClientRequestToken)
	}
	if len(_wellarchitectedShareId) > 0 {
		input.ShareId = aws.String(_wellarchitectedShareId)
	}
	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}

	if resp, err := client.DeleteWorkloadShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociate a lens from a workload.
// Up to 10 lenses can be disassociated from a workload in a single API operation.
//
// The Amazon Web Services Well-Architected Framework lens ( wellarchitected )
// cannot be removed from a workload.
func wellarchitected_DisassociateLenses(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.DisassociateLensesInput{
		// LensAliases: []string, // Required
		// WorkloadId: *string, // Required
	}

	if len(_wellarchitectedLensAliases) > 0 {
		input.LensAliases = append([]string(nil), _wellarchitectedLensAliases...)
	}
	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}

	if resp, err := client.DisassociateLenses(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociate a profile from a workload.
func wellarchitected_DisassociateProfiles(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.DisassociateProfilesInput{
		// ProfileArns: []string, // Required
		// WorkloadId: *string, // Required
	}

	if len(_wellarchitectedProfileArns) > 0 {
		input.ProfileArns = append([]string(nil), _wellarchitectedProfileArns...)
	}
	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}

	if resp, err := client.DisassociateProfiles(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Export an existing lens.
// Only the owner of a lens can export it. Lenses provided by Amazon Web Services
// (Amazon Web Services Official Content) cannot be exported.
//
// Lenses are defined in JSON. For more information, see [JSON format specification] in the Well-Architected
// Tool User Guide.
//
// # Disclaimer
//
// Do not include or gather personal identifiable information (PII) of end users
// or other identifiable individuals in or via your custom lenses. If your custom
// lens or those shared with you and used in your account do include or collect PII
// you are responsible for: ensuring that the included PII is processed in
// accordance with applicable law, providing adequate privacy notices, and
// obtaining necessary consents for processing such data.
//
// [JSON format specification]: https://docs.aws.amazon.com/wellarchitected/latest/userguide/lenses-format-specification.html
func wellarchitected_ExportLens(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.ExportLensInput{
		// LensAlias: *string, // Required
	}

	if len(_wellarchitectedLensAlias) > 0 {
		input.LensAlias = aws.String(_wellarchitectedLensAlias)
	}
	if len(_wellarchitectedLensVersion) > 0 {
		input.LensVersion = aws.String(_wellarchitectedLensVersion)
	}

	if resp, err := client.ExportLens(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the answer to a specific question in a workload review.
func wellarchitected_GetAnswer(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.GetAnswerInput{
		// LensAlias: *string, // Required
		// QuestionId: *string, // Required
		// WorkloadId: *string, // Required
	}

	if len(_wellarchitectedLensAlias) > 0 {
		input.LensAlias = aws.String(_wellarchitectedLensAlias)
	}
	if len(_wellarchitectedQuestionId) > 0 {
		input.QuestionId = aws.String(_wellarchitectedQuestionId)
	}
	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}
	if len(_wellarchitectedMilestoneNumber) > 0 {
		if err := assignInputField(input, "MilestoneNumber", _wellarchitectedMilestoneNumber); err != nil {
			log.Errorf("invalid --milestone-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetAnswer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get a consolidated report of your workloads.
// You can optionally choose to include workloads that have been shared with you.
func wellarchitected_GetConsolidatedReport(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.GetConsolidatedReportInput{
		// Format: types.ReportFormat, // Required
	}

	if len(_wellarchitectedFormat) > 0 {
		if err := assignInputField(input, "Format", _wellarchitectedFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedIncludeSharedResources) > 0 {
		if err := assignInputField(input, "IncludeSharedResources", _wellarchitectedIncludeSharedResources); err != nil {
			log.Errorf("invalid --include-shared-resources: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wellarchitectedMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedNextToken) > 0 {
		input.NextToken = aws.String(_wellarchitectedNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetConsolidatedReport(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*wellarchitected.GetConsolidatedReportOutput
	p := wellarchitected.NewGetConsolidatedReportPaginator(client, input)
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

// Global settings for all workloads.
func wellarchitected_GetGlobalSettings(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.GetGlobalSettingsInput{}

	if resp, err := client.GetGlobalSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get an existing lens.
func wellarchitected_GetLens(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.GetLensInput{
		// LensAlias: *string, // Required
	}

	if len(_wellarchitectedLensAlias) > 0 {
		input.LensAlias = aws.String(_wellarchitectedLensAlias)
	}
	if len(_wellarchitectedLensVersion) > 0 {
		input.LensVersion = aws.String(_wellarchitectedLensVersion)
	}

	if resp, err := client.GetLens(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get lens review.
func wellarchitected_GetLensReview(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.GetLensReviewInput{
		// LensAlias: *string, // Required
		// WorkloadId: *string, // Required
	}

	if len(_wellarchitectedLensAlias) > 0 {
		input.LensAlias = aws.String(_wellarchitectedLensAlias)
	}
	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}
	if len(_wellarchitectedMilestoneNumber) > 0 {
		if err := assignInputField(input, "MilestoneNumber", _wellarchitectedMilestoneNumber); err != nil {
			log.Errorf("invalid --milestone-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetLensReview(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get lens review report.
func wellarchitected_GetLensReviewReport(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.GetLensReviewReportInput{
		// LensAlias: *string, // Required
		// WorkloadId: *string, // Required
	}

	if len(_wellarchitectedLensAlias) > 0 {
		input.LensAlias = aws.String(_wellarchitectedLensAlias)
	}
	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}
	if len(_wellarchitectedMilestoneNumber) > 0 {
		if err := assignInputField(input, "MilestoneNumber", _wellarchitectedMilestoneNumber); err != nil {
			log.Errorf("invalid --milestone-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetLensReviewReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get lens version differences.
func wellarchitected_GetLensVersionDifference(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.GetLensVersionDifferenceInput{
		// LensAlias: *string, // Required
	}

	if len(_wellarchitectedLensAlias) > 0 {
		input.LensAlias = aws.String(_wellarchitectedLensAlias)
	}
	if len(_wellarchitectedBaseLensVersion) > 0 {
		input.BaseLensVersion = aws.String(_wellarchitectedBaseLensVersion)
	}
	if len(_wellarchitectedTargetLensVersion) > 0 {
		input.TargetLensVersion = aws.String(_wellarchitectedTargetLensVersion)
	}

	if resp, err := client.GetLensVersionDifference(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get a milestone for an existing workload.
func wellarchitected_GetMilestone(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.GetMilestoneInput{
		// MilestoneNumber: *int32, // Required
		// WorkloadId: *string, // Required
	}

	if len(_wellarchitectedMilestoneNumber) > 0 {
		if err := assignInputField(input, "MilestoneNumber", _wellarchitectedMilestoneNumber); err != nil {
			log.Errorf("invalid --milestone-number: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}

	if resp, err := client.GetMilestone(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get profile information.
func wellarchitected_GetProfile(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.GetProfileInput{
		// ProfileArn: *string, // Required
	}

	if len(_wellarchitectedProfileArn) > 0 {
		input.ProfileArn = aws.String(_wellarchitectedProfileArn)
	}
	if len(_wellarchitectedProfileVersion) > 0 {
		input.ProfileVersion = aws.String(_wellarchitectedProfileVersion)
	}

	if resp, err := client.GetProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get profile template.
func wellarchitected_GetProfileTemplate(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.GetProfileTemplateInput{}

	if resp, err := client.GetProfileTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get review template.
func wellarchitected_GetReviewTemplate(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.GetReviewTemplateInput{
		// TemplateArn: *string, // Required
	}

	if len(_wellarchitectedTemplateArn) > 0 {
		input.TemplateArn = aws.String(_wellarchitectedTemplateArn)
	}

	if resp, err := client.GetReviewTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get review template answer.
func wellarchitected_GetReviewTemplateAnswer(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.GetReviewTemplateAnswerInput{
		// LensAlias: *string, // Required
		// QuestionId: *string, // Required
		// TemplateArn: *string, // Required
	}

	if len(_wellarchitectedLensAlias) > 0 {
		input.LensAlias = aws.String(_wellarchitectedLensAlias)
	}
	if len(_wellarchitectedQuestionId) > 0 {
		input.QuestionId = aws.String(_wellarchitectedQuestionId)
	}
	if len(_wellarchitectedTemplateArn) > 0 {
		input.TemplateArn = aws.String(_wellarchitectedTemplateArn)
	}

	if resp, err := client.GetReviewTemplateAnswer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get a lens review associated with a review template.
func wellarchitected_GetReviewTemplateLensReview(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.GetReviewTemplateLensReviewInput{
		// LensAlias: *string, // Required
		// TemplateArn: *string, // Required
	}

	if len(_wellarchitectedLensAlias) > 0 {
		input.LensAlias = aws.String(_wellarchitectedLensAlias)
	}
	if len(_wellarchitectedTemplateArn) > 0 {
		input.TemplateArn = aws.String(_wellarchitectedTemplateArn)
	}

	if resp, err := client.GetReviewTemplateLensReview(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get an existing workload.
func wellarchitected_GetWorkload(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.GetWorkloadInput{
		// WorkloadId: *string, // Required
	}

	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}

	if resp, err := client.GetWorkload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Import a new custom lens or update an existing custom lens.
// To update an existing custom lens, specify its ARN as the LensAlias . If no ARN
// is specified, a new custom lens is created.
//
// The new or updated lens will have a status of DRAFT . The lens cannot be applied
// to workloads or shared with other Amazon Web Services accounts until it's
// published with CreateLensVersion.
//
// Lenses are defined in JSON. For more information, see [JSON format specification] in the Well-Architected
// Tool User Guide.
//
// A custom lens cannot exceed 500 KB in size.
//
// # Disclaimer
//
// Do not include or gather personal identifiable information (PII) of end users
// or other identifiable individuals in or via your custom lenses. If your custom
// lens or those shared with you and used in your account do include or collect PII
// you are responsible for: ensuring that the included PII is processed in
// accordance with applicable law, providing adequate privacy notices, and
// obtaining necessary consents for processing such data.
//
// [JSON format specification]: https://docs.aws.amazon.com/wellarchitected/latest/userguide/lenses-format-specification.html
func wellarchitected_ImportLens(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.ImportLensInput{
		// ClientRequestToken: *string, // Required
		// JSONString: *string, // Required
	}

	if len(_wellarchitectedClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_wellarchitectedClientRequestToken)
	}
	if len(_wellarchitectedJSONString) > 0 {
		input.JSONString = aws.String(_wellarchitectedJSONString)
	}
	if len(_wellarchitectedLensAlias) > 0 {
		input.LensAlias = aws.String(_wellarchitectedLensAlias)
	}
	if len(_wellarchitectedTags) > 0 {
		if err := assignInputField(input, "Tags", _wellarchitectedTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportLens(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List of answers for a particular workload and lens.
func wellarchitected_ListAnswers(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.ListAnswersInput{
		// LensAlias: *string, // Required
		// WorkloadId: *string, // Required
	}

	if len(_wellarchitectedLensAlias) > 0 {
		input.LensAlias = aws.String(_wellarchitectedLensAlias)
	}
	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}
	if len(_wellarchitectedMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wellarchitectedMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedMilestoneNumber) > 0 {
		if err := assignInputField(input, "MilestoneNumber", _wellarchitectedMilestoneNumber); err != nil {
			log.Errorf("invalid --milestone-number: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedNextToken) > 0 {
		input.NextToken = aws.String(_wellarchitectedNextToken)
	}
	if len(_wellarchitectedPillarId) > 0 {
		input.PillarId = aws.String(_wellarchitectedPillarId)
	}
	if len(_wellarchitectedQuestionPriority) > 0 {
		if err := assignInputField(input, "QuestionPriority", _wellarchitectedQuestionPriority); err != nil {
			log.Errorf("invalid --question-priority: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAnswers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*wellarchitected.ListAnswersOutput
	p := wellarchitected.NewListAnswersPaginator(client, input)
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

// List of Trusted Advisor check details by account related to the workload.
func wellarchitected_ListCheckDetails(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.ListCheckDetailsInput{
		// ChoiceId: *string, // Required
		// LensArn: *string, // Required
		// PillarId: *string, // Required
		// QuestionId: *string, // Required
		// WorkloadId: *string, // Required
	}

	if len(_wellarchitectedChoiceId) > 0 {
		input.ChoiceId = aws.String(_wellarchitectedChoiceId)
	}
	if len(_wellarchitectedLensArn) > 0 {
		input.LensArn = aws.String(_wellarchitectedLensArn)
	}
	if len(_wellarchitectedPillarId) > 0 {
		input.PillarId = aws.String(_wellarchitectedPillarId)
	}
	if len(_wellarchitectedQuestionId) > 0 {
		input.QuestionId = aws.String(_wellarchitectedQuestionId)
	}
	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}
	if len(_wellarchitectedMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wellarchitectedMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedNextToken) > 0 {
		input.NextToken = aws.String(_wellarchitectedNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCheckDetails(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*wellarchitected.ListCheckDetailsOutput
	p := wellarchitected.NewListCheckDetailsPaginator(client, input)
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

// List of Trusted Advisor checks summarized for all accounts related to the
// workload.
func wellarchitected_ListCheckSummaries(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.ListCheckSummariesInput{
		// ChoiceId: *string, // Required
		// LensArn: *string, // Required
		// PillarId: *string, // Required
		// QuestionId: *string, // Required
		// WorkloadId: *string, // Required
	}

	if len(_wellarchitectedChoiceId) > 0 {
		input.ChoiceId = aws.String(_wellarchitectedChoiceId)
	}
	if len(_wellarchitectedLensArn) > 0 {
		input.LensArn = aws.String(_wellarchitectedLensArn)
	}
	if len(_wellarchitectedPillarId) > 0 {
		input.PillarId = aws.String(_wellarchitectedPillarId)
	}
	if len(_wellarchitectedQuestionId) > 0 {
		input.QuestionId = aws.String(_wellarchitectedQuestionId)
	}
	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}
	if len(_wellarchitectedMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wellarchitectedMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedNextToken) > 0 {
		input.NextToken = aws.String(_wellarchitectedNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCheckSummaries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*wellarchitected.ListCheckSummariesOutput
	p := wellarchitected.NewListCheckSummariesPaginator(client, input)
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

// List the improvements of a particular lens review.
func wellarchitected_ListLensReviewImprovements(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.ListLensReviewImprovementsInput{
		// LensAlias: *string, // Required
		// WorkloadId: *string, // Required
	}

	if len(_wellarchitectedLensAlias) > 0 {
		input.LensAlias = aws.String(_wellarchitectedLensAlias)
	}
	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}
	if len(_wellarchitectedMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wellarchitectedMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedMilestoneNumber) > 0 {
		if err := assignInputField(input, "MilestoneNumber", _wellarchitectedMilestoneNumber); err != nil {
			log.Errorf("invalid --milestone-number: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedNextToken) > 0 {
		input.NextToken = aws.String(_wellarchitectedNextToken)
	}
	if len(_wellarchitectedPillarId) > 0 {
		input.PillarId = aws.String(_wellarchitectedPillarId)
	}
	if len(_wellarchitectedQuestionPriority) > 0 {
		if err := assignInputField(input, "QuestionPriority", _wellarchitectedQuestionPriority); err != nil {
			log.Errorf("invalid --question-priority: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListLensReviewImprovements(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*wellarchitected.ListLensReviewImprovementsOutput
	p := wellarchitected.NewListLensReviewImprovementsPaginator(client, input)
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

// List lens reviews for a particular workload.
func wellarchitected_ListLensReviews(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.ListLensReviewsInput{
		// WorkloadId: *string, // Required
	}

	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}
	if len(_wellarchitectedMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wellarchitectedMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedMilestoneNumber) > 0 {
		if err := assignInputField(input, "MilestoneNumber", _wellarchitectedMilestoneNumber); err != nil {
			log.Errorf("invalid --milestone-number: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedNextToken) > 0 {
		input.NextToken = aws.String(_wellarchitectedNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLensReviews(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*wellarchitected.ListLensReviewsOutput
	p := wellarchitected.NewListLensReviewsPaginator(client, input)
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

// List the lens shares associated with the lens.
func wellarchitected_ListLensShares(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.ListLensSharesInput{
		// LensAlias: *string, // Required
	}

	if len(_wellarchitectedLensAlias) > 0 {
		input.LensAlias = aws.String(_wellarchitectedLensAlias)
	}
	if len(_wellarchitectedMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wellarchitectedMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedNextToken) > 0 {
		input.NextToken = aws.String(_wellarchitectedNextToken)
	}
	if len(_wellarchitectedSharedWithPrefix) > 0 {
		input.SharedWithPrefix = aws.String(_wellarchitectedSharedWithPrefix)
	}
	if len(_wellarchitectedStatus) > 0 {
		if err := assignInputField(input, "Status", _wellarchitectedStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListLensShares(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*wellarchitected.ListLensSharesOutput
	p := wellarchitected.NewListLensSharesPaginator(client, input)
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

// List the available lenses.
func wellarchitected_ListLenses(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.ListLensesInput{}

	if len(_wellarchitectedLensName) > 0 {
		input.LensName = aws.String(_wellarchitectedLensName)
	}
	if len(_wellarchitectedLensStatus) > 0 {
		if err := assignInputField(input, "LensStatus", _wellarchitectedLensStatus); err != nil {
			log.Errorf("invalid --lens-status: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedLensType) > 0 {
		if err := assignInputField(input, "LensType", _wellarchitectedLensType); err != nil {
			log.Errorf("invalid --lens-type: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wellarchitectedMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedNextToken) > 0 {
		input.NextToken = aws.String(_wellarchitectedNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLenses(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*wellarchitected.ListLensesOutput
	p := wellarchitected.NewListLensesPaginator(client, input)
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

// List all milestones for an existing workload.
func wellarchitected_ListMilestones(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.ListMilestonesInput{
		// WorkloadId: *string, // Required
	}

	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}
	if len(_wellarchitectedMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wellarchitectedMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedNextToken) > 0 {
		input.NextToken = aws.String(_wellarchitectedNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMilestones(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*wellarchitected.ListMilestonesOutput
	p := wellarchitected.NewListMilestonesPaginator(client, input)
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

// List lens notifications.
func wellarchitected_ListNotifications(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.ListNotificationsInput{}

	if len(_wellarchitectedMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wellarchitectedMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedNextToken) > 0 {
		input.NextToken = aws.String(_wellarchitectedNextToken)
	}
	if len(_wellarchitectedResourceArn) > 0 {
		input.ResourceArn = aws.String(_wellarchitectedResourceArn)
	}
	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}

	if disablePaginator() {
		if resp, err := client.ListNotifications(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*wellarchitected.ListNotificationsOutput
	p := wellarchitected.NewListNotificationsPaginator(client, input)
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

// List profile notifications.
func wellarchitected_ListProfileNotifications(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.ListProfileNotificationsInput{}

	if len(_wellarchitectedMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wellarchitectedMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedNextToken) > 0 {
		input.NextToken = aws.String(_wellarchitectedNextToken)
	}
	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}

	if disablePaginator() {
		if resp, err := client.ListProfileNotifications(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*wellarchitected.ListProfileNotificationsOutput
	p := wellarchitected.NewListProfileNotificationsPaginator(client, input)
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

// List profile shares.
func wellarchitected_ListProfileShares(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.ListProfileSharesInput{
		// ProfileArn: *string, // Required
	}

	if len(_wellarchitectedProfileArn) > 0 {
		input.ProfileArn = aws.String(_wellarchitectedProfileArn)
	}
	if len(_wellarchitectedMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wellarchitectedMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedNextToken) > 0 {
		input.NextToken = aws.String(_wellarchitectedNextToken)
	}
	if len(_wellarchitectedSharedWithPrefix) > 0 {
		input.SharedWithPrefix = aws.String(_wellarchitectedSharedWithPrefix)
	}
	if len(_wellarchitectedStatus) > 0 {
		if err := assignInputField(input, "Status", _wellarchitectedStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListProfileShares(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*wellarchitected.ListProfileSharesOutput
	p := wellarchitected.NewListProfileSharesPaginator(client, input)
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

// List profiles.
func wellarchitected_ListProfiles(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.ListProfilesInput{}

	if len(_wellarchitectedMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wellarchitectedMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedNextToken) > 0 {
		input.NextToken = aws.String(_wellarchitectedNextToken)
	}
	if len(_wellarchitectedProfileNamePrefix) > 0 {
		input.ProfileNamePrefix = aws.String(_wellarchitectedProfileNamePrefix)
	}
	if len(_wellarchitectedProfileOwnerType) > 0 {
		if err := assignInputField(input, "ProfileOwnerType", _wellarchitectedProfileOwnerType); err != nil {
			log.Errorf("invalid --profile-owner-type: %s", err.Error())
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

	var results []*wellarchitected.ListProfilesOutput
	p := wellarchitected.NewListProfilesPaginator(client, input)
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

// List the answers of a review template.
func wellarchitected_ListReviewTemplateAnswers(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.ListReviewTemplateAnswersInput{
		// LensAlias: *string, // Required
		// TemplateArn: *string, // Required
	}

	if len(_wellarchitectedLensAlias) > 0 {
		input.LensAlias = aws.String(_wellarchitectedLensAlias)
	}
	if len(_wellarchitectedTemplateArn) > 0 {
		input.TemplateArn = aws.String(_wellarchitectedTemplateArn)
	}
	if len(_wellarchitectedMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wellarchitectedMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedNextToken) > 0 {
		input.NextToken = aws.String(_wellarchitectedNextToken)
	}
	if len(_wellarchitectedPillarId) > 0 {
		input.PillarId = aws.String(_wellarchitectedPillarId)
	}

	if disablePaginator() {
		if resp, err := client.ListReviewTemplateAnswers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*wellarchitected.ListReviewTemplateAnswersOutput
	p := wellarchitected.NewListReviewTemplateAnswersPaginator(client, input)
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

// List review templates.
func wellarchitected_ListReviewTemplates(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.ListReviewTemplatesInput{}

	if len(_wellarchitectedMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wellarchitectedMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedNextToken) > 0 {
		input.NextToken = aws.String(_wellarchitectedNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListReviewTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*wellarchitected.ListReviewTemplatesOutput
	p := wellarchitected.NewListReviewTemplatesPaginator(client, input)
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

// List the share invitations.
// WorkloadNamePrefix , LensNamePrefix , ProfileNamePrefix , and TemplateNamePrefix
// are mutually exclusive. Use the parameter that matches your ShareResourceType .
func wellarchitected_ListShareInvitations(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.ListShareInvitationsInput{}

	if len(_wellarchitectedLensNamePrefix) > 0 {
		input.LensNamePrefix = aws.String(_wellarchitectedLensNamePrefix)
	}
	if len(_wellarchitectedMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wellarchitectedMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedNextToken) > 0 {
		input.NextToken = aws.String(_wellarchitectedNextToken)
	}
	if len(_wellarchitectedProfileNamePrefix) > 0 {
		input.ProfileNamePrefix = aws.String(_wellarchitectedProfileNamePrefix)
	}
	if len(_wellarchitectedShareResourceType) > 0 {
		if err := assignInputField(input, "ShareResourceType", _wellarchitectedShareResourceType); err != nil {
			log.Errorf("invalid --share-resource-type: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedTemplateNamePrefix) > 0 {
		input.TemplateNamePrefix = aws.String(_wellarchitectedTemplateNamePrefix)
	}
	if len(_wellarchitectedWorkloadNamePrefix) > 0 {
		input.WorkloadNamePrefix = aws.String(_wellarchitectedWorkloadNamePrefix)
	}

	if disablePaginator() {
		if resp, err := client.ListShareInvitations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*wellarchitected.ListShareInvitationsOutput
	p := wellarchitected.NewListShareInvitationsPaginator(client, input)
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

// List the tags for a resource.
// The WorkloadArn parameter can be a workload ARN, a custom lens ARN, a profile
// ARN, or review template ARN.
func wellarchitected_ListTagsForResource(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.ListTagsForResourceInput{
		// WorkloadArn: *string, // Required
	}

	if len(_wellarchitectedWorkloadArn) > 0 {
		input.WorkloadArn = aws.String(_wellarchitectedWorkloadArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List review template shares.
func wellarchitected_ListTemplateShares(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.ListTemplateSharesInput{
		// TemplateArn: *string, // Required
	}

	if len(_wellarchitectedTemplateArn) > 0 {
		input.TemplateArn = aws.String(_wellarchitectedTemplateArn)
	}
	if len(_wellarchitectedMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wellarchitectedMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedNextToken) > 0 {
		input.NextToken = aws.String(_wellarchitectedNextToken)
	}
	if len(_wellarchitectedSharedWithPrefix) > 0 {
		input.SharedWithPrefix = aws.String(_wellarchitectedSharedWithPrefix)
	}
	if len(_wellarchitectedStatus) > 0 {
		if err := assignInputField(input, "Status", _wellarchitectedStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTemplateShares(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*wellarchitected.ListTemplateSharesOutput
	p := wellarchitected.NewListTemplateSharesPaginator(client, input)
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

// List the workload shares associated with the workload.
func wellarchitected_ListWorkloadShares(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.ListWorkloadSharesInput{
		// WorkloadId: *string, // Required
	}

	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}
	if len(_wellarchitectedMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wellarchitectedMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedNextToken) > 0 {
		input.NextToken = aws.String(_wellarchitectedNextToken)
	}
	if len(_wellarchitectedSharedWithPrefix) > 0 {
		input.SharedWithPrefix = aws.String(_wellarchitectedSharedWithPrefix)
	}
	if len(_wellarchitectedStatus) > 0 {
		if err := assignInputField(input, "Status", _wellarchitectedStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListWorkloadShares(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*wellarchitected.ListWorkloadSharesOutput
	p := wellarchitected.NewListWorkloadSharesPaginator(client, input)
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

// Paginated list of workloads.
func wellarchitected_ListWorkloads(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.ListWorkloadsInput{}

	if len(_wellarchitectedMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _wellarchitectedMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedNextToken) > 0 {
		input.NextToken = aws.String(_wellarchitectedNextToken)
	}
	if len(_wellarchitectedWorkloadNamePrefix) > 0 {
		input.WorkloadNamePrefix = aws.String(_wellarchitectedWorkloadNamePrefix)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkloads(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*wellarchitected.ListWorkloadsOutput
	p := wellarchitected.NewListWorkloadsPaginator(client, input)
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

// Adds one or more tags to the specified resource.
// The WorkloadArn parameter can be a workload ARN, a custom lens ARN, a profile
// ARN, or review template ARN.
func wellarchitected_TagResource(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.TagResourceInput{
		// Tags: map[string]string, // Required
		// WorkloadArn: *string, // Required
	}

	if len(_wellarchitectedTags) > 0 {
		if err := assignInputField(input, "Tags", _wellarchitectedTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedWorkloadArn) > 0 {
		input.WorkloadArn = aws.String(_wellarchitectedWorkloadArn)
	}

	if resp, err := client.TagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes specified tags from a resource.
// The WorkloadArn parameter can be a workload ARN, a custom lens ARN, a profile
// ARN, or review template ARN.
//
// To specify multiple tags, use separate tagKeys parameters, for example:
//
// DELETE /tags/WorkloadArn?tagKeys=key1&tagKeys=key2
func wellarchitected_UntagResource(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.UntagResourceInput{
		// TagKeys: []string, // Required
		// WorkloadArn: *string, // Required
	}

	if len(_wellarchitectedTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _wellarchitectedTagKeys...)
	}
	if len(_wellarchitectedWorkloadArn) > 0 {
		input.WorkloadArn = aws.String(_wellarchitectedWorkloadArn)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the answer to a specific question in a workload review.
func wellarchitected_UpdateAnswer(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.UpdateAnswerInput{
		// LensAlias: *string, // Required
		// QuestionId: *string, // Required
		// WorkloadId: *string, // Required
	}

	if len(_wellarchitectedLensAlias) > 0 {
		input.LensAlias = aws.String(_wellarchitectedLensAlias)
	}
	if len(_wellarchitectedQuestionId) > 0 {
		input.QuestionId = aws.String(_wellarchitectedQuestionId)
	}
	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}
	if len(_wellarchitectedChoiceUpdates) > 0 {
		if err := assignInputField(input, "ChoiceUpdates", _wellarchitectedChoiceUpdates); err != nil {
			log.Errorf("invalid --choice-updates: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedIsApplicable) > 0 {
		if err := assignInputField(input, "IsApplicable", _wellarchitectedIsApplicable); err != nil {
			log.Errorf("invalid --is-applicable: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedNotes) > 0 {
		input.Notes = aws.String(_wellarchitectedNotes)
	}
	if len(_wellarchitectedReason) > 0 {
		if err := assignInputField(input, "Reason", _wellarchitectedReason); err != nil {
			log.Errorf("invalid --reason: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedSelectedChoices) > 0 {
		input.SelectedChoices = append([]string(nil), _wellarchitectedSelectedChoices...)
	}

	if resp, err := client.UpdateAnswer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update whether the Amazon Web Services account is opted into organization
// sharing and discovery integration features.
func wellarchitected_UpdateGlobalSettings(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.UpdateGlobalSettingsInput{}

	if len(_wellarchitectedDiscoveryIntegrationStatus) > 0 {
		if err := assignInputField(input, "DiscoveryIntegrationStatus", _wellarchitectedDiscoveryIntegrationStatus); err != nil {
			log.Errorf("invalid --discovery-integration-status: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedJiraConfiguration) > 0 {
		if err := assignInputField(input, "JiraConfiguration", _wellarchitectedJiraConfiguration); err != nil {
			log.Errorf("invalid --jira-configuration: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedOrganizationSharingStatus) > 0 {
		if err := assignInputField(input, "OrganizationSharingStatus", _wellarchitectedOrganizationSharingStatus); err != nil {
			log.Errorf("invalid --organization-sharing-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateGlobalSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update integration features.
func wellarchitected_UpdateIntegration(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.UpdateIntegrationInput{
		// ClientRequestToken: *string, // Required
		// IntegratingService: types.IntegratingService, // Required
		// WorkloadId: *string, // Required
	}

	if len(_wellarchitectedClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_wellarchitectedClientRequestToken)
	}
	if len(_wellarchitectedIntegratingService) > 0 {
		if err := assignInputField(input, "IntegratingService", _wellarchitectedIntegratingService); err != nil {
			log.Errorf("invalid --integrating-service: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}

	if resp, err := client.UpdateIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update lens review for a particular workload.
func wellarchitected_UpdateLensReview(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.UpdateLensReviewInput{
		// LensAlias: *string, // Required
		// WorkloadId: *string, // Required
	}

	if len(_wellarchitectedLensAlias) > 0 {
		input.LensAlias = aws.String(_wellarchitectedLensAlias)
	}
	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}
	if len(_wellarchitectedJiraConfiguration) > 0 {
		if err := assignInputField(input, "JiraConfiguration", _wellarchitectedJiraConfiguration); err != nil {
			log.Errorf("invalid --jira-configuration: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedLensNotes) > 0 {
		input.LensNotes = aws.String(_wellarchitectedLensNotes)
	}
	if len(_wellarchitectedPillarNotes) > 0 {
		if err := assignInputField(input, "PillarNotes", _wellarchitectedPillarNotes); err != nil {
			log.Errorf("invalid --pillar-notes: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateLensReview(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a profile.
func wellarchitected_UpdateProfile(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.UpdateProfileInput{
		// ProfileArn: *string, // Required
	}

	if len(_wellarchitectedProfileArn) > 0 {
		input.ProfileArn = aws.String(_wellarchitectedProfileArn)
	}
	if len(_wellarchitectedProfileDescription) > 0 {
		input.ProfileDescription = aws.String(_wellarchitectedProfileDescription)
	}
	if len(_wellarchitectedProfileQuestions) > 0 {
		if err := assignInputField(input, "ProfileQuestions", _wellarchitectedProfileQuestions); err != nil {
			log.Errorf("invalid --profile-questions: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a review template.
func wellarchitected_UpdateReviewTemplate(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.UpdateReviewTemplateInput{
		// TemplateArn: *string, // Required
	}

	if len(_wellarchitectedTemplateArn) > 0 {
		input.TemplateArn = aws.String(_wellarchitectedTemplateArn)
	}
	if len(_wellarchitectedDescription) > 0 {
		input.Description = aws.String(_wellarchitectedDescription)
	}
	if len(_wellarchitectedLensesToAssociate) > 0 {
		input.LensesToAssociate = append([]string(nil), _wellarchitectedLensesToAssociate...)
	}
	if len(_wellarchitectedLensesToDisassociate) > 0 {
		input.LensesToDisassociate = append([]string(nil), _wellarchitectedLensesToDisassociate...)
	}
	if len(_wellarchitectedNotes) > 0 {
		input.Notes = aws.String(_wellarchitectedNotes)
	}
	if len(_wellarchitectedTemplateName) > 0 {
		input.TemplateName = aws.String(_wellarchitectedTemplateName)
	}

	if resp, err := client.UpdateReviewTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a review template answer.
func wellarchitected_UpdateReviewTemplateAnswer(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.UpdateReviewTemplateAnswerInput{
		// LensAlias: *string, // Required
		// QuestionId: *string, // Required
		// TemplateArn: *string, // Required
	}

	if len(_wellarchitectedLensAlias) > 0 {
		input.LensAlias = aws.String(_wellarchitectedLensAlias)
	}
	if len(_wellarchitectedQuestionId) > 0 {
		input.QuestionId = aws.String(_wellarchitectedQuestionId)
	}
	if len(_wellarchitectedTemplateArn) > 0 {
		input.TemplateArn = aws.String(_wellarchitectedTemplateArn)
	}
	if len(_wellarchitectedChoiceUpdates) > 0 {
		if err := assignInputField(input, "ChoiceUpdates", _wellarchitectedChoiceUpdates); err != nil {
			log.Errorf("invalid --choice-updates: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedIsApplicable) > 0 {
		if err := assignInputField(input, "IsApplicable", _wellarchitectedIsApplicable); err != nil {
			log.Errorf("invalid --is-applicable: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedNotes) > 0 {
		input.Notes = aws.String(_wellarchitectedNotes)
	}
	if len(_wellarchitectedReason) > 0 {
		if err := assignInputField(input, "Reason", _wellarchitectedReason); err != nil {
			log.Errorf("invalid --reason: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedSelectedChoices) > 0 {
		input.SelectedChoices = append([]string(nil), _wellarchitectedSelectedChoices...)
	}

	if resp, err := client.UpdateReviewTemplateAnswer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a lens review associated with a review template.
func wellarchitected_UpdateReviewTemplateLensReview(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.UpdateReviewTemplateLensReviewInput{
		// LensAlias: *string, // Required
		// TemplateArn: *string, // Required
	}

	if len(_wellarchitectedLensAlias) > 0 {
		input.LensAlias = aws.String(_wellarchitectedLensAlias)
	}
	if len(_wellarchitectedTemplateArn) > 0 {
		input.TemplateArn = aws.String(_wellarchitectedTemplateArn)
	}
	if len(_wellarchitectedLensNotes) > 0 {
		input.LensNotes = aws.String(_wellarchitectedLensNotes)
	}
	if len(_wellarchitectedPillarNotes) > 0 {
		if err := assignInputField(input, "PillarNotes", _wellarchitectedPillarNotes); err != nil {
			log.Errorf("invalid --pillar-notes: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateReviewTemplateLensReview(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a workload or custom lens share invitation.
// This API operation can be called independently of any resource. Previous
// documentation implied that a workload ARN must be specified.
func wellarchitected_UpdateShareInvitation(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.UpdateShareInvitationInput{
		// ShareInvitationAction: types.ShareInvitationAction, // Required
		// ShareInvitationId: *string, // Required
	}

	if len(_wellarchitectedShareInvitationAction) > 0 {
		if err := assignInputField(input, "ShareInvitationAction", _wellarchitectedShareInvitationAction); err != nil {
			log.Errorf("invalid --share-invitation-action: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedShareInvitationId) > 0 {
		input.ShareInvitationId = aws.String(_wellarchitectedShareInvitationId)
	}

	if resp, err := client.UpdateShareInvitation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update an existing workload.
func wellarchitected_UpdateWorkload(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.UpdateWorkloadInput{
		// WorkloadId: *string, // Required
	}

	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}
	if len(_wellarchitectedAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _wellarchitectedAccountIds...)
	}
	if len(_wellarchitectedApplications) > 0 {
		input.Applications = append([]string(nil), _wellarchitectedApplications...)
	}
	if len(_wellarchitectedArchitecturalDesign) > 0 {
		input.ArchitecturalDesign = aws.String(_wellarchitectedArchitecturalDesign)
	}
	if len(_wellarchitectedAwsRegions) > 0 {
		input.AwsRegions = append([]string(nil), _wellarchitectedAwsRegions...)
	}
	if len(_wellarchitectedDescription) > 0 {
		input.Description = aws.String(_wellarchitectedDescription)
	}
	if len(_wellarchitectedDiscoveryConfig) > 0 {
		if err := assignInputField(input, "DiscoveryConfig", _wellarchitectedDiscoveryConfig); err != nil {
			log.Errorf("invalid --discovery-config: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedEnvironment) > 0 {
		if err := assignInputField(input, "Environment", _wellarchitectedEnvironment); err != nil {
			log.Errorf("invalid --environment: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedImprovementStatus) > 0 {
		if err := assignInputField(input, "ImprovementStatus", _wellarchitectedImprovementStatus); err != nil {
			log.Errorf("invalid --improvement-status: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedIndustry) > 0 {
		input.Industry = aws.String(_wellarchitectedIndustry)
	}
	if len(_wellarchitectedIndustryType) > 0 {
		input.IndustryType = aws.String(_wellarchitectedIndustryType)
	}
	if len(_wellarchitectedIsReviewOwnerUpdateAcknowledged) > 0 {
		if err := assignInputField(input, "IsReviewOwnerUpdateAcknowledged", _wellarchitectedIsReviewOwnerUpdateAcknowledged); err != nil {
			log.Errorf("invalid --is-review-owner-update-acknowledged: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedJiraConfiguration) > 0 {
		if err := assignInputField(input, "JiraConfiguration", _wellarchitectedJiraConfiguration); err != nil {
			log.Errorf("invalid --jira-configuration: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedNonAwsRegions) > 0 {
		input.NonAwsRegions = append([]string(nil), _wellarchitectedNonAwsRegions...)
	}
	if len(_wellarchitectedNotes) > 0 {
		input.Notes = aws.String(_wellarchitectedNotes)
	}
	if len(_wellarchitectedPillarPriorities) > 0 {
		input.PillarPriorities = append([]string(nil), _wellarchitectedPillarPriorities...)
	}
	if len(_wellarchitectedReviewOwner) > 0 {
		input.ReviewOwner = aws.String(_wellarchitectedReviewOwner)
	}
	if len(_wellarchitectedWorkloadName) > 0 {
		input.WorkloadName = aws.String(_wellarchitectedWorkloadName)
	}

	if resp, err := client.UpdateWorkload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a workload share.
func wellarchitected_UpdateWorkloadShare(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.UpdateWorkloadShareInput{
		// PermissionType: types.PermissionType, // Required
		// ShareId: *string, // Required
		// WorkloadId: *string, // Required
	}

	if len(_wellarchitectedPermissionType) > 0 {
		if err := assignInputField(input, "PermissionType", _wellarchitectedPermissionType); err != nil {
			log.Errorf("invalid --permission-type: %s", err.Error())
			return
		}
	}
	if len(_wellarchitectedShareId) > 0 {
		input.ShareId = aws.String(_wellarchitectedShareId)
	}
	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}

	if resp, err := client.UpdateWorkloadShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Upgrade lens review for a particular workload.
func wellarchitected_UpgradeLensReview(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.UpgradeLensReviewInput{
		// LensAlias: *string, // Required
		// MilestoneName: *string, // Required
		// WorkloadId: *string, // Required
	}

	if len(_wellarchitectedLensAlias) > 0 {
		input.LensAlias = aws.String(_wellarchitectedLensAlias)
	}
	if len(_wellarchitectedMilestoneName) > 0 {
		input.MilestoneName = aws.String(_wellarchitectedMilestoneName)
	}
	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}
	if len(_wellarchitectedClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_wellarchitectedClientRequestToken)
	}

	if resp, err := client.UpgradeLensReview(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Upgrade a profile.
func wellarchitected_UpgradeProfileVersion(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.UpgradeProfileVersionInput{
		// ProfileArn: *string, // Required
		// WorkloadId: *string, // Required
	}

	if len(_wellarchitectedProfileArn) > 0 {
		input.ProfileArn = aws.String(_wellarchitectedProfileArn)
	}
	if len(_wellarchitectedWorkloadId) > 0 {
		input.WorkloadId = aws.String(_wellarchitectedWorkloadId)
	}
	if len(_wellarchitectedClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_wellarchitectedClientRequestToken)
	}
	if len(_wellarchitectedMilestoneName) > 0 {
		input.MilestoneName = aws.String(_wellarchitectedMilestoneName)
	}

	if resp, err := client.UpgradeProfileVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Upgrade the lens review of a review template.
func wellarchitected_UpgradeReviewTemplateLensReview(cfg aws.Config, client *wellarchitected.Client) {
	input := &wellarchitected.UpgradeReviewTemplateLensReviewInput{
		// LensAlias: *string, // Required
		// TemplateArn: *string, // Required
	}

	if len(_wellarchitectedLensAlias) > 0 {
		input.LensAlias = aws.String(_wellarchitectedLensAlias)
	}
	if len(_wellarchitectedTemplateArn) > 0 {
		input.TemplateArn = aws.String(_wellarchitectedTemplateArn)
	}
	if len(_wellarchitectedClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_wellarchitectedClientRequestToken)
	}

	if resp, err := client.UpgradeReviewTemplateLensReview(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_wellarchitectedCmd)
	_wellarchitectedCmd.Flags().SortFlags = false

	_wellarchitectedCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_wellarchitectedCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_wellarchitectedCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_wellarchitectedCmd.Flags().StringSliceVarP(&_wellarchitectedAccountIds, "account-ids", "", nil, "Account Ids")
	_wellarchitectedCmd.Flags().StringSliceVarP(&_wellarchitectedApplications, "applications", "", nil, "Applications")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedArchitecturalDesign, "architectural-design", "", "", "Architectural Design")
	_wellarchitectedCmd.Flags().StringSliceVarP(&_wellarchitectedAwsRegions, "aws-regions", "", nil, "AWS Regions")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedBaseLensVersion, "base-lens-version", "", "", "Base Lens Version")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedChoiceId, "choice-id", "", "", "Choice ID")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedChoiceUpdates, "choice-updates", "", "", "Choice Updates")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedDescription, "description", "", "", "Description")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedDiscoveryConfig, "discovery-config", "", "", "Discovery Config")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedDiscoveryIntegrationStatus, "discovery-integration-status", "", "", "Discovery Integration Status")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedEnvironment, "environment", "", "", "Environment")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedFormat, "format", "", "", "Format")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedImprovementStatus, "improvement-status", "", "", "Improvement Status")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedIncludeSharedResources, "include-shared-resources", "", "", "Include Shared Resources")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedIndustry, "industry", "", "", "Industry")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedIndustryType, "industry-type", "", "", "Industry Type")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedIntegratingService, "integrating-service", "", "", "Integrating Service")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedIsApplicable, "is-applicable", "", "", "Is Applicable")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedIsMajorVersion, "is-major-version", "", "", "Is Major Version")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedIsReviewOwnerUpdateAcknowledged, "is-review-owner-update-acknowledged", "", "", "Is Review Owner Update Acknowledged")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedJiraConfiguration, "jira-configuration", "", "", "Jira Configuration")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedJSONString, "json-string", "", "", "JSON String")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedLensAlias, "lens-alias", "", "", "Lens Alias")
	_wellarchitectedCmd.Flags().StringSliceVarP(&_wellarchitectedLensAliases, "lens-aliases", "", nil, "Lens Aliases")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedLensArn, "lens-arn", "", "", "Lens ARN")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedLensName, "lens-name", "", "", "Lens Name")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedLensNamePrefix, "lens-name-prefix", "", "", "Lens Name Prefix")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedLensNotes, "lens-notes", "", "", "Lens Notes")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedLensStatus, "lens-status", "", "", "Lens Status")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedLensType, "lens-type", "", "", "Lens Type")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedLensVersion, "lens-version", "", "", "Lens Version")
	_wellarchitectedCmd.Flags().StringSliceVarP(&_wellarchitectedLenses, "lenses", "", nil, "Lenses")
	_wellarchitectedCmd.Flags().StringSliceVarP(&_wellarchitectedLensesToAssociate, "lenses-to-associate", "", nil, "Lenses To Associate")
	_wellarchitectedCmd.Flags().StringSliceVarP(&_wellarchitectedLensesToDisassociate, "lenses-to-disassociate", "", nil, "Lenses To Disassociate")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedMaxResults, "max-results", "", "", "Max Results")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedMilestoneName, "milestone-name", "", "", "Milestone Name")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedMilestoneNumber, "milestone-number", "", "", "Milestone Number")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedNextToken, "next-token", "", "", "Next Token")
	_wellarchitectedCmd.Flags().StringSliceVarP(&_wellarchitectedNonAwsRegions, "non-aws-regions", "", nil, "Non AWS Regions")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedNotes, "notes", "", "", "Notes")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedOrganizationSharingStatus, "organization-sharing-status", "", "", "Organization Sharing Status")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedPermissionType, "permission-type", "", "", "Permission Type")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedPillarId, "pillar-id", "", "", "Pillar ID")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedPillarNotes, "pillar-notes", "", "", "Pillar Notes")
	_wellarchitectedCmd.Flags().StringSliceVarP(&_wellarchitectedPillarPriorities, "pillar-priorities", "", nil, "Pillar Priorities")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedProfileArn, "profile-arn", "", "", "Profile ARN")
	_wellarchitectedCmd.Flags().StringSliceVarP(&_wellarchitectedProfileArns, "profile-arns", "", nil, "Profile Arns")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedProfileDescription, "profile-description", "", "", "Profile Description")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedProfileName, "profile-name", "", "", "Profile Name")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedProfileNamePrefix, "profile-name-prefix", "", "", "Profile Name Prefix")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedProfileOwnerType, "profile-owner-type", "", "", "Profile Owner Type")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedProfileQuestions, "profile-questions", "", "", "Profile Questions")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedProfileVersion, "profile-version", "", "", "Profile Version")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedQuestionId, "question-id", "", "", "Question ID")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedQuestionPriority, "question-priority", "", "", "Question Priority")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedReason, "reason", "", "", "Reason")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedResourceArn, "resource-arn", "", "", "Resource ARN")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedReviewOwner, "review-owner", "", "", "Review Owner")
	_wellarchitectedCmd.Flags().StringSliceVarP(&_wellarchitectedReviewTemplateArns, "review-template-arns", "", nil, "Review Template Arns")
	_wellarchitectedCmd.Flags().StringSliceVarP(&_wellarchitectedSelectedChoices, "selected-choices", "", nil, "Selected Choices")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedShareId, "share-id", "", "", "Share ID")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedShareInvitationAction, "share-invitation-action", "", "", "Share Invitation Action")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedShareInvitationId, "share-invitation-id", "", "", "Share Invitation ID")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedShareResourceType, "share-resource-type", "", "", "Share Resource Type")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedSharedWith, "shared-with", "", "", "Shared With")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedSharedWithPrefix, "shared-with-prefix", "", "", "Shared With Prefix")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedStatus, "status", "", "", "Status")
	_wellarchitectedCmd.Flags().StringSliceVarP(&_wellarchitectedTagKeys, "tag-keys", "", nil, "Tag Keys")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedTags, "tags", "", "", "Tags")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedTargetLensVersion, "target-lens-version", "", "", "Target Lens Version")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedTemplateArn, "template-arn", "", "", "Template ARN")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedTemplateName, "template-name", "", "", "Template Name")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedTemplateNamePrefix, "template-name-prefix", "", "", "Template Name Prefix")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedWorkloadArn, "workload-arn", "", "", "Workload ARN")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedWorkloadId, "workload-id", "", "", "Workload ID")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedWorkloadName, "workload-name", "", "", "Workload Name")
	_wellarchitectedCmd.Flags().StringVarP(&_wellarchitectedWorkloadNamePrefix, "workload-name-prefix", "", "", "Workload Name Prefix")

	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedAssociateLenses, "associate-lenses", "", false, "Associate Lenses")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedAssociateProfiles, "associate-profiles", "", false, "Associate Profiles")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedCreateLensShare, "create-lens-share", "", false, "Create Lens Share")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedCreateLensVersion, "create-lens-version", "", false, "Create Lens Version")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedCreateMilestone, "create-milestone", "", false, "Create Milestone")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedCreateProfile, "create-profile", "", false, "Create Profile")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedCreateProfileShare, "create-profile-share", "", false, "Create Profile Share")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedCreateReviewTemplate, "create-review-template", "", false, "Create Review Template")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedCreateTemplateShare, "create-template-share", "", false, "Create Template Share")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedCreateWorkload, "create-workload", "", false, "Create Workload")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedCreateWorkloadShare, "create-workload-share", "", false, "Create Workload Share")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedDeleteLens, "delete-lens", "", false, "Delete Lens")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedDeleteLensShare, "delete-lens-share", "", false, "Delete Lens Share")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedDeleteProfile, "delete-profile", "", false, "Delete Profile")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedDeleteProfileShare, "delete-profile-share", "", false, "Delete Profile Share")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedDeleteReviewTemplate, "delete-review-template", "", false, "Delete Review Template")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedDeleteTemplateShare, "delete-template-share", "", false, "Delete Template Share")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedDeleteWorkload, "delete-workload", "", false, "Delete Workload")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedDeleteWorkloadShare, "delete-workload-share", "", false, "Delete Workload Share")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedDisassociateLenses, "disassociate-lenses", "", false, "Disassociate Lenses")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedDisassociateProfiles, "disassociate-profiles", "", false, "Disassociate Profiles")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedExportLens, "export-lens", "", false, "Export Lens")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedGetAnswer, "get-answer", "", false, "Get Answer")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedGetConsolidatedReport, "get-consolidated-report", "", false, "Get Consolidated Report")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedGetGlobalSettings, "get-global-settings", "", false, "Get Global Settings")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedGetLens, "get-lens", "", false, "Get Lens")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedGetLensReview, "get-lens-review", "", false, "Get Lens Review")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedGetLensReviewReport, "get-lens-review-report", "", false, "Get Lens Review Report")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedGetLensVersionDifference, "get-lens-version-difference", "", false, "Get Lens Version Difference")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedGetMilestone, "get-milestone", "", false, "Get Milestone")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedGetProfile, "get-profile", "", false, "Get Profile")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedGetProfileTemplate, "get-profile-template", "", false, "Get Profile Template")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedGetReviewTemplate, "get-review-template", "", false, "Get Review Template")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedGetReviewTemplateAnswer, "get-review-template-answer", "", false, "Get Review Template Answer")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedGetReviewTemplateLensReview, "get-review-template-lens-review", "", false, "Get Review Template Lens Review")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedGetWorkload, "get-workload", "", false, "Get Workload")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedImportLens, "import-lens", "", false, "Import Lens")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedListAnswers, "list-answers", "", false, "List Answers")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedListCheckDetails, "list-check-details", "", false, "List Check Details")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedListCheckSummaries, "list-check-summaries", "", false, "List Check Summaries")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedListLensReviewImprovements, "list-lens-review-improvements", "", false, "List Lens Review Improvements")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedListLensReviews, "list-lens-reviews", "", false, "List Lens Reviews")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedListLensShares, "list-lens-shares", "", false, "List Lens Shares")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedListLenses, "list-lenses", "", false, "List Lenses")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedListMilestones, "list-milestones", "", false, "List Milestones")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedListNotifications, "list-notifications", "", false, "List Notifications")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedListProfileNotifications, "list-profile-notifications", "", false, "List Profile Notifications")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedListProfileShares, "list-profile-shares", "", false, "List Profile Shares")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedListProfiles, "list-profiles", "", false, "List Profiles")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedListReviewTemplateAnswers, "list-review-template-answers", "", false, "List Review Template Answers")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedListReviewTemplates, "list-review-templates", "", false, "List Review Templates")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedListShareInvitations, "list-share-invitations", "", false, "List Share Invitations")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedListTemplateShares, "list-template-shares", "", false, "List Template Shares")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedListWorkloadShares, "list-workload-shares", "", false, "List Workload Shares")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedListWorkloads, "list-workloads", "", false, "List Workloads")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedTagResource, "tag-resource", "", false, "Tag Resource")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedUntagResource, "untag-resource", "", false, "Untag Resource")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedUpdateAnswer, "update-answer", "", false, "Update Answer")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedUpdateGlobalSettings, "update-global-settings", "", false, "Update Global Settings")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedUpdateIntegration, "update-integration", "", false, "Update Integration")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedUpdateLensReview, "update-lens-review", "", false, "Update Lens Review")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedUpdateProfile, "update-profile", "", false, "Update Profile")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedUpdateReviewTemplate, "update-review-template", "", false, "Update Review Template")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedUpdateReviewTemplateAnswer, "update-review-template-answer", "", false, "Update Review Template Answer")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedUpdateReviewTemplateLensReview, "update-review-template-lens-review", "", false, "Update Review Template Lens Review")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedUpdateShareInvitation, "update-share-invitation", "", false, "Update Share Invitation")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedUpdateWorkload, "update-workload", "", false, "Update Workload")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedUpdateWorkloadShare, "update-workload-share", "", false, "Update Workload Share")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedUpgradeLensReview, "upgrade-lens-review", "", false, "Upgrade Lens Review")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedUpgradeProfileVersion, "upgrade-profile-version", "", false, "Upgrade Profile Version")
	_wellarchitectedCmd.Flags().BoolVarP(&_wellarchitectedUpgradeReviewTemplateLensReview, "upgrade-review-template-lens-review", "", false, "Upgrade Review Template Lens Review")

}
