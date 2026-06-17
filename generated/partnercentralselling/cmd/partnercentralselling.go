package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/partnercentralselling"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// partnercentralsellingCmd represents the partnercentralselling command
var _partnercentralsellingCmd = &cobra.Command{
	Use:   "partnercentralselling",
	Short: "AWS partnercentralselling CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := partnercentralselling.NewFromConfig(cfg)
		if _partnercentralsellingAcceptEngagementInvitation {
			partnercentralselling_AcceptEngagementInvitation(cfg, client)
			return
		}
		if _partnercentralsellingAssignOpportunity {
			partnercentralselling_AssignOpportunity(cfg, client)
			return
		}
		if _partnercentralsellingAssociateOpportunity {
			partnercentralselling_AssociateOpportunity(cfg, client)
			return
		}
		if _partnercentralsellingCreateEngagement {
			partnercentralselling_CreateEngagement(cfg, client)
			return
		}
		if _partnercentralsellingCreateEngagementContext {
			partnercentralselling_CreateEngagementContext(cfg, client)
			return
		}
		if _partnercentralsellingCreateEngagementInvitation {
			partnercentralselling_CreateEngagementInvitation(cfg, client)
			return
		}
		if _partnercentralsellingCreateOpportunity {
			partnercentralselling_CreateOpportunity(cfg, client)
			return
		}
		if _partnercentralsellingCreateResourceSnapshot {
			partnercentralselling_CreateResourceSnapshot(cfg, client)
			return
		}
		if _partnercentralsellingCreateResourceSnapshotJob {
			partnercentralselling_CreateResourceSnapshotJob(cfg, client)
			return
		}
		if _partnercentralsellingDeleteResourceSnapshotJob {
			partnercentralselling_DeleteResourceSnapshotJob(cfg, client)
			return
		}
		if _partnercentralsellingDisassociateOpportunity {
			partnercentralselling_DisassociateOpportunity(cfg, client)
			return
		}
		if _partnercentralsellingGetAwsOpportunitySummary {
			partnercentralselling_GetAwsOpportunitySummary(cfg, client)
			return
		}
		if _partnercentralsellingGetEngagement {
			partnercentralselling_GetEngagement(cfg, client)
			return
		}
		if _partnercentralsellingGetEngagementInvitation {
			partnercentralselling_GetEngagementInvitation(cfg, client)
			return
		}
		if _partnercentralsellingGetOpportunity {
			partnercentralselling_GetOpportunity(cfg, client)
			return
		}
		if _partnercentralsellingGetResourceSnapshot {
			partnercentralselling_GetResourceSnapshot(cfg, client)
			return
		}
		if _partnercentralsellingGetResourceSnapshotJob {
			partnercentralselling_GetResourceSnapshotJob(cfg, client)
			return
		}
		if _partnercentralsellingGetSellingSystemSettings {
			partnercentralselling_GetSellingSystemSettings(cfg, client)
			return
		}
		if _partnercentralsellingListEngagementByAcceptingInvitationTasks {
			partnercentralselling_ListEngagementByAcceptingInvitationTasks(cfg, client)
			return
		}
		if _partnercentralsellingListEngagementFromOpportunityTasks {
			partnercentralselling_ListEngagementFromOpportunityTasks(cfg, client)
			return
		}
		if _partnercentralsellingListEngagementInvitations {
			partnercentralselling_ListEngagementInvitations(cfg, client)
			return
		}
		if _partnercentralsellingListEngagementMembers {
			partnercentralselling_ListEngagementMembers(cfg, client)
			return
		}
		if _partnercentralsellingListEngagementResourceAssociations {
			partnercentralselling_ListEngagementResourceAssociations(cfg, client)
			return
		}
		if _partnercentralsellingListEngagements {
			partnercentralselling_ListEngagements(cfg, client)
			return
		}
		if _partnercentralsellingListOpportunities {
			partnercentralselling_ListOpportunities(cfg, client)
			return
		}
		if _partnercentralsellingListOpportunityFromEngagementTasks {
			partnercentralselling_ListOpportunityFromEngagementTasks(cfg, client)
			return
		}
		if _partnercentralsellingListResourceSnapshotJobs {
			partnercentralselling_ListResourceSnapshotJobs(cfg, client)
			return
		}
		if _partnercentralsellingListResourceSnapshots {
			partnercentralselling_ListResourceSnapshots(cfg, client)
			return
		}
		if _partnercentralsellingListSolutions {
			partnercentralselling_ListSolutions(cfg, client)
			return
		}
		if _partnercentralsellingListTagsForResource {
			partnercentralselling_ListTagsForResource(cfg, client)
			return
		}
		if _partnercentralsellingPutSellingSystemSettings {
			partnercentralselling_PutSellingSystemSettings(cfg, client)
			return
		}
		if _partnercentralsellingRejectEngagementInvitation {
			partnercentralselling_RejectEngagementInvitation(cfg, client)
			return
		}
		if _partnercentralsellingStartEngagementByAcceptingInvitationTask {
			partnercentralselling_StartEngagementByAcceptingInvitationTask(cfg, client)
			return
		}
		if _partnercentralsellingStartEngagementFromOpportunityTask {
			partnercentralselling_StartEngagementFromOpportunityTask(cfg, client)
			return
		}
		if _partnercentralsellingStartOpportunityFromEngagementTask {
			partnercentralselling_StartOpportunityFromEngagementTask(cfg, client)
			return
		}
		if _partnercentralsellingStartResourceSnapshotJob {
			partnercentralselling_StartResourceSnapshotJob(cfg, client)
			return
		}
		if _partnercentralsellingStopResourceSnapshotJob {
			partnercentralselling_StopResourceSnapshotJob(cfg, client)
			return
		}
		if _partnercentralsellingSubmitOpportunity {
			partnercentralselling_SubmitOpportunity(cfg, client)
			return
		}
		if _partnercentralsellingTagResource {
			partnercentralselling_TagResource(cfg, client)
			return
		}
		if _partnercentralsellingUntagResource {
			partnercentralselling_UntagResource(cfg, client)
			return
		}
		if _partnercentralsellingUpdateEngagementContext {
			partnercentralselling_UpdateEngagementContext(cfg, client)
			return
		}
		if _partnercentralsellingUpdateOpportunity {
			partnercentralselling_UpdateOpportunity(cfg, client)
			return
		}

	},
}

var (
	_partnercentralsellingAcceptEngagementInvitation               bool
	_partnercentralsellingAssignOpportunity                        bool
	_partnercentralsellingAssociateOpportunity                     bool
	_partnercentralsellingCreateEngagement                         bool
	_partnercentralsellingCreateEngagementContext                  bool
	_partnercentralsellingCreateEngagementInvitation               bool
	_partnercentralsellingCreateOpportunity                        bool
	_partnercentralsellingCreateResourceSnapshot                   bool
	_partnercentralsellingCreateResourceSnapshotJob                bool
	_partnercentralsellingDeleteResourceSnapshotJob                bool
	_partnercentralsellingDisassociateOpportunity                  bool
	_partnercentralsellingGetAwsOpportunitySummary                 bool
	_partnercentralsellingGetEngagement                            bool
	_partnercentralsellingGetEngagementInvitation                  bool
	_partnercentralsellingGetOpportunity                           bool
	_partnercentralsellingGetResourceSnapshot                      bool
	_partnercentralsellingGetResourceSnapshotJob                   bool
	_partnercentralsellingGetSellingSystemSettings                 bool
	_partnercentralsellingListEngagementByAcceptingInvitationTasks bool
	_partnercentralsellingListEngagementFromOpportunityTasks       bool
	_partnercentralsellingListEngagementInvitations                bool
	_partnercentralsellingListEngagementMembers                    bool
	_partnercentralsellingListEngagementResourceAssociations       bool
	_partnercentralsellingListEngagements                          bool
	_partnercentralsellingListOpportunities                        bool
	_partnercentralsellingListOpportunityFromEngagementTasks       bool
	_partnercentralsellingListResourceSnapshotJobs                 bool
	_partnercentralsellingListResourceSnapshots                    bool
	_partnercentralsellingListSolutions                            bool
	_partnercentralsellingListTagsForResource                      bool
	_partnercentralsellingPutSellingSystemSettings                 bool
	_partnercentralsellingRejectEngagementInvitation               bool
	_partnercentralsellingStartEngagementByAcceptingInvitationTask bool
	_partnercentralsellingStartEngagementFromOpportunityTask       bool
	_partnercentralsellingStartOpportunityFromEngagementTask       bool
	_partnercentralsellingStartResourceSnapshotJob                 bool
	_partnercentralsellingStopResourceSnapshotJob                  bool
	_partnercentralsellingSubmitOpportunity                        bool
	_partnercentralsellingTagResource                              bool
	_partnercentralsellingUntagResource                            bool
	_partnercentralsellingUpdateEngagementContext                  bool
	_partnercentralsellingUpdateOpportunity                        bool

	_partnercentralsellingAssignee                           string
	_partnercentralsellingAwsSubmission                      string
	_partnercentralsellingCatalog                            string
	_partnercentralsellingCategory                           []string
	_partnercentralsellingClientToken                        string
	_partnercentralsellingContextIdentifier                  string
	_partnercentralsellingContextTypes                       string
	_partnercentralsellingContexts                           string
	_partnercentralsellingCreatedBy                          string
	_partnercentralsellingCreatedDate                        string
	_partnercentralsellingCustomer                           string
	_partnercentralsellingCustomerCompanyName                []string
	_partnercentralsellingDescription                        string
	_partnercentralsellingEngagementIdentifier               string
	_partnercentralsellingEngagementInvitationIdentifier     []string
	_partnercentralsellingEngagementLastModifiedAt           string
	_partnercentralsellingExcludeContextTypes                string
	_partnercentralsellingExcludeCreatedBy                   []string
	_partnercentralsellingIdentifier                         string
	_partnercentralsellingInvitation                         string
	_partnercentralsellingInvolvementType                    string
	_partnercentralsellingLastModifiedDate                   string
	_partnercentralsellingLifeCycle                          string
	_partnercentralsellingLifeCycleReviewStatus              string
	_partnercentralsellingLifeCycleStage                     string
	_partnercentralsellingMarketing                          string
	_partnercentralsellingMaxResults                         string
	_partnercentralsellingNationalSecurity                   string
	_partnercentralsellingNextToken                          string
	_partnercentralsellingOpportunityIdentifier              []string
	_partnercentralsellingOpportunityTeam                    string
	_partnercentralsellingOpportunityType                    string
	_partnercentralsellingOrigin                             string
	_partnercentralsellingParticipantType                    string
	_partnercentralsellingPartnerOpportunityIdentifier       string
	_partnercentralsellingPayload                            string
	_partnercentralsellingPayloadType                        string
	_partnercentralsellingPrimaryNeedsFromAws                string
	_partnercentralsellingProject                            string
	_partnercentralsellingRejectionReason                    string
	_partnercentralsellingRelatedEntityIdentifier            string
	_partnercentralsellingRelatedEntityType                  string
	_partnercentralsellingRelatedOpportunityIdentifier       string
	_partnercentralsellingResourceArn                        string
	_partnercentralsellingResourceIdentifier                 string
	_partnercentralsellingResourceSnapshotJobIdentifier      string
	_partnercentralsellingResourceSnapshotJobRoleIdentifier  string
	_partnercentralsellingResourceSnapshotTemplateIdentifier string
	_partnercentralsellingResourceType                       string
	_partnercentralsellingRevision                           string
	_partnercentralsellingSenderAwsAccountId                 []string
	_partnercentralsellingSoftwareRevenue                    string
	_partnercentralsellingSort                               string
	_partnercentralsellingStatus                             string
	_partnercentralsellingTagKeys                            []string
	_partnercentralsellingTags                               string
	_partnercentralsellingTargetCloseDate                    string
	_partnercentralsellingTaskIdentifier                     []string
	_partnercentralsellingTaskStatus                         string
	_partnercentralsellingTitle                              string
	_partnercentralsellingType                               string
	_partnercentralsellingVisibility                         string
)

// Use the AcceptEngagementInvitation action to accept an engagement invitation
// shared by AWS. Accepting the invitation indicates your willingness to
// participate in the engagement, granting you access to all engagement-related
// data.
func partnercentralselling_AcceptEngagementInvitation(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.AcceptEngagementInvitationInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralsellingIdentifier)
	}

	if resp, err := client.AcceptEngagementInvitation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables you to reassign an existing Opportunity to another user within your
// Partner Central account. The specified user receives the opportunity, and it
// appears on their Partner Central dashboard, allowing them to take necessary
// actions or proceed with the opportunity.
//
// This is useful for distributing opportunities to the appropriate team members
// or departments within your organization, ensuring that each opportunity is
// handled by the right person. By default, the opportunity owner is the one who
// creates it. Currently, there's no API to enumerate the list of available users.
func partnercentralselling_AssignOpportunity(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.AssignOpportunityInput{
		// Assignee: *types.AssigneeContact, // Required
		// Catalog: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralsellingAssignee) > 0 {
		if err := assignInputField(input, "Assignee", _partnercentralsellingAssignee); err != nil {
			log.Errorf("invalid --assignee: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralsellingIdentifier)
	}

	if resp, err := client.AssignOpportunity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables you to create a formal association between an Opportunity and various
// related entities, enriching the context and details of the opportunity for
// better collaboration and decision making. You can associate an opportunity with
// the following entity types:
//
// - Partner Solution: A software product or consulting practice created and
// delivered by Partners. Partner Solutions help customers address business
// challenges using Amazon Web Services services.
//
// - Amazon Web Services Products: Amazon Web Services offers many products and
// services that provide scalable, reliable, and cost-effective infrastructure
// solutions. For the latest list of Amazon Web Services products, see [Amazon Web Services products].
//
// - Amazon Web Services Marketplace private offer: Allows Amazon Web Services
// Marketplace sellers to extend custom pricing and terms to individual Amazon Web
// Services customers. Sellers can negotiate custom prices, payment schedules, and
// end user license terms through private offers, enabling Amazon Web Services
// customers to acquire software solutions tailored to their specific needs. For
// more information, see [Private offers in Amazon Web Services Marketplace].
//
// To obtain identifiers for these entities, use the following methods:
//
// - Solution: Use the ListSolutions operation.
//
// - AWS Products: For the latest list of Amazon Web Services products, see [Amazon Web Services products].
//
// - Amazon Web Services Marketplace private offer: Use the [Using the Amazon Web Services Marketplace Catalog API]to list entities.
// Specifically, use the ListEntities operation to retrieve a list of private
// offers. The request returns the details of available private offers. For more
// information, see [ListEntities].
//
// [Private offers in Amazon Web Services Marketplace]: https://docs.aws.amazon.com/marketplace/latest/buyerguide/buyer-private-offers.html
// [Using the Amazon Web Services Marketplace Catalog API]: https://docs.aws.amazon.com/marketplace/latest/APIReference/catalog-apis.html
// [ListEntities]: https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/API_ListEntities.html
// [Amazon Web Services products]: https://github.com/aws-samples/partner-crm-integration-samples/blob/main/resources/aws_products.json
func partnercentralselling_AssociateOpportunity(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.AssociateOpportunityInput{
		// Catalog: *string, // Required
		// OpportunityIdentifier: *string, // Required
		// RelatedEntityIdentifier: *string, // Required
		// RelatedEntityType: types.RelatedEntityType, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingOpportunityIdentifier) > 0 {
		input.OpportunityIdentifier = aws.String(_partnercentralsellingOpportunityIdentifier[0])
	}
	if len(_partnercentralsellingRelatedEntityIdentifier) > 0 {
		input.RelatedEntityIdentifier = aws.String(_partnercentralsellingRelatedEntityIdentifier)
	}
	if len(_partnercentralsellingRelatedEntityType) > 0 {
		if err := assignInputField(input, "RelatedEntityType", _partnercentralsellingRelatedEntityType); err != nil {
			log.Errorf("invalid --related-entity-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateOpportunity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The CreateEngagement action allows you to create an Engagement , which serves as
// a collaborative space between different parties such as AWS Partners and AWS
// Sellers. This action automatically adds the caller's AWS account as an active
// member of the newly created Engagement .
func partnercentralselling_CreateEngagement(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.CreateEngagementInput{
		// Catalog: *string, // Required
		// ClientToken: *string, // Required
		// Description: *string, // Required
		// Title: *string, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralsellingClientToken)
	}
	if len(_partnercentralsellingDescription) > 0 {
		input.Description = aws.String(_partnercentralsellingDescription)
	}
	if len(_partnercentralsellingTitle) > 0 {
		input.Title = aws.String(_partnercentralsellingTitle)
	}
	if len(_partnercentralsellingContexts) > 0 {
		if err := assignInputField(input, "Contexts", _partnercentralsellingContexts); err != nil {
			log.Errorf("invalid --contexts: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEngagement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new context within an existing engagement. This action allows you to
// add contextual information such as customer projects or documents to an
// engagement, providing additional details that help facilitate collaboration
// between engagement members.
func partnercentralselling_CreateEngagementContext(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.CreateEngagementContextInput{
		// Catalog: *string, // Required
		// ClientToken: *string, // Required
		// EngagementIdentifier: *string, // Required
		// Payload: types.EngagementContextPayload, // Required
		// Type: types.EngagementContextType, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralsellingClientToken)
	}
	if len(_partnercentralsellingEngagementIdentifier) > 0 {
		input.EngagementIdentifier = aws.String(_partnercentralsellingEngagementIdentifier)
	}
	if len(_partnercentralsellingPayload) > 0 {
		if err := assignInputField(input, "Payload", _partnercentralsellingPayload); err != nil {
			log.Errorf("invalid --payload: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingType) > 0 {
		if err := assignInputField(input, "Type", _partnercentralsellingType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEngagementContext(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action creates an invitation from a sender to a single receiver to join
// an engagement.
func partnercentralselling_CreateEngagementInvitation(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.CreateEngagementInvitationInput{
		// Catalog: *string, // Required
		// ClientToken: *string, // Required
		// EngagementIdentifier: *string, // Required
		// Invitation: *types.Invitation, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralsellingClientToken)
	}
	if len(_partnercentralsellingEngagementIdentifier) > 0 {
		input.EngagementIdentifier = aws.String(_partnercentralsellingEngagementIdentifier)
	}
	if len(_partnercentralsellingInvitation) > 0 {
		if err := assignInputField(input, "Invitation", _partnercentralsellingInvitation); err != nil {
			log.Errorf("invalid --invitation: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEngagementInvitation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Opportunity record in Partner Central. Use this operation to create
// a potential business opportunity for submission to Amazon Web Services. Creating
// an opportunity sets Lifecycle.ReviewStatus to Pending Submission .
//
// To submit an opportunity, follow these steps:
//
// - To create the opportunity, use CreateOpportunity .
//
// - To associate a solution with the opportunity, use AssociateOpportunity .
//
// - To start the engagement with AWS, use StartEngagementFromOpportunity .
//
// After submission, you can't edit the opportunity until the review is complete.
// But opportunities in the Pending Submission state must have complete details.
// You can update the opportunity while it's in the Pending Submission state.
//
// There's a set of mandatory fields to create opportunities, but consider
// providing optional fields to enrich the opportunity record.
func partnercentralselling_CreateOpportunity(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.CreateOpportunityInput{
		// Catalog: *string, // Required
		// ClientToken: *string, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralsellingClientToken)
	}
	if len(_partnercentralsellingCustomer) > 0 {
		if err := assignInputField(input, "Customer", _partnercentralsellingCustomer); err != nil {
			log.Errorf("invalid --customer: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingLifeCycle) > 0 {
		if err := assignInputField(input, "LifeCycle", _partnercentralsellingLifeCycle); err != nil {
			log.Errorf("invalid --life-cycle: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingMarketing) > 0 {
		if err := assignInputField(input, "Marketing", _partnercentralsellingMarketing); err != nil {
			log.Errorf("invalid --marketing: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingNationalSecurity) > 0 {
		if err := assignInputField(input, "NationalSecurity", _partnercentralsellingNationalSecurity); err != nil {
			log.Errorf("invalid --national-security: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingOpportunityTeam) > 0 {
		if err := assignInputField(input, "OpportunityTeam", _partnercentralsellingOpportunityTeam); err != nil {
			log.Errorf("invalid --opportunity-team: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingOpportunityType) > 0 {
		if err := assignInputField(input, "OpportunityType", _partnercentralsellingOpportunityType); err != nil {
			log.Errorf("invalid --opportunity-type: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingOrigin) > 0 {
		if err := assignInputField(input, "Origin", _partnercentralsellingOrigin); err != nil {
			log.Errorf("invalid --origin: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingPartnerOpportunityIdentifier) > 0 {
		input.PartnerOpportunityIdentifier = aws.String(_partnercentralsellingPartnerOpportunityIdentifier)
	}
	if len(_partnercentralsellingPrimaryNeedsFromAws) > 0 {
		if err := assignInputField(input, "PrimaryNeedsFromAws", _partnercentralsellingPrimaryNeedsFromAws); err != nil {
			log.Errorf("invalid --primary-needs-from-aws: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingProject) > 0 {
		if err := assignInputField(input, "Project", _partnercentralsellingProject); err != nil {
			log.Errorf("invalid --project: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingSoftwareRevenue) > 0 {
		if err := assignInputField(input, "SoftwareRevenue", _partnercentralsellingSoftwareRevenue); err != nil {
			log.Errorf("invalid --software-revenue: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingTags) > 0 {
		if err := assignInputField(input, "Tags", _partnercentralsellingTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateOpportunity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action allows you to create an immutable snapshot of a specific resource,
// such as an opportunity, within the context of an engagement. The snapshot
// captures a subset of the resource's data based on the schema defined by the
// provided template.
func partnercentralselling_CreateResourceSnapshot(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.CreateResourceSnapshotInput{
		// Catalog: *string, // Required
		// ClientToken: *string, // Required
		// EngagementIdentifier: *string, // Required
		// ResourceIdentifier: *string, // Required
		// ResourceSnapshotTemplateIdentifier: *string, // Required
		// ResourceType: types.ResourceType, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralsellingClientToken)
	}
	if len(_partnercentralsellingEngagementIdentifier) > 0 {
		input.EngagementIdentifier = aws.String(_partnercentralsellingEngagementIdentifier)
	}
	if len(_partnercentralsellingResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_partnercentralsellingResourceIdentifier)
	}
	if len(_partnercentralsellingResourceSnapshotTemplateIdentifier) > 0 {
		input.ResourceSnapshotTemplateIdentifier = aws.String(_partnercentralsellingResourceSnapshotTemplateIdentifier)
	}
	if len(_partnercentralsellingResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _partnercentralsellingResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateResourceSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this action to create a job to generate a snapshot of the specified
// resource within an engagement. It initiates an asynchronous process to create a
// resource snapshot. The job creates a new snapshot only if the resource state has
// changed, adhering to the same access control and immutability rules as direct
// snapshot creation.
func partnercentralselling_CreateResourceSnapshotJob(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.CreateResourceSnapshotJobInput{
		// Catalog: *string, // Required
		// ClientToken: *string, // Required
		// EngagementIdentifier: *string, // Required
		// ResourceIdentifier: *string, // Required
		// ResourceSnapshotTemplateIdentifier: *string, // Required
		// ResourceType: types.ResourceType, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralsellingClientToken)
	}
	if len(_partnercentralsellingEngagementIdentifier) > 0 {
		input.EngagementIdentifier = aws.String(_partnercentralsellingEngagementIdentifier)
	}
	if len(_partnercentralsellingResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_partnercentralsellingResourceIdentifier)
	}
	if len(_partnercentralsellingResourceSnapshotTemplateIdentifier) > 0 {
		input.ResourceSnapshotTemplateIdentifier = aws.String(_partnercentralsellingResourceSnapshotTemplateIdentifier)
	}
	if len(_partnercentralsellingResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _partnercentralsellingResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingTags) > 0 {
		if err := assignInputField(input, "Tags", _partnercentralsellingTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateResourceSnapshotJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this action to deletes a previously created resource snapshot job. The job
// must be in a stopped state before it can be deleted.
func partnercentralselling_DeleteResourceSnapshotJob(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.DeleteResourceSnapshotJobInput{
		// Catalog: *string, // Required
		// ResourceSnapshotJobIdentifier: *string, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingResourceSnapshotJobIdentifier) > 0 {
		input.ResourceSnapshotJobIdentifier = aws.String(_partnercentralsellingResourceSnapshotJobIdentifier)
	}

	if resp, err := client.DeleteResourceSnapshotJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows you to remove an existing association between an Opportunity and related
// entities, such as a Partner Solution, Amazon Web Services product, or an Amazon
// Web Services Marketplace offer. This operation is the counterpart to
// AssociateOpportunity , and it provides flexibility to manage associations as
// business needs change.
//
// Use this operation to update the associations of an Opportunity due to changes
// in the related entities, or if an association was made in error. Ensuring
// accurate associations helps maintain clarity and accuracy to track and manage
// business opportunities. When you replace an entity, first attach the new entity
// and then disassociate the one to be removed, especially if it's the last
// remaining entity that's required.
func partnercentralselling_DisassociateOpportunity(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.DisassociateOpportunityInput{
		// Catalog: *string, // Required
		// OpportunityIdentifier: *string, // Required
		// RelatedEntityIdentifier: *string, // Required
		// RelatedEntityType: types.RelatedEntityType, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingOpportunityIdentifier) > 0 {
		input.OpportunityIdentifier = aws.String(_partnercentralsellingOpportunityIdentifier[0])
	}
	if len(_partnercentralsellingRelatedEntityIdentifier) > 0 {
		input.RelatedEntityIdentifier = aws.String(_partnercentralsellingRelatedEntityIdentifier)
	}
	if len(_partnercentralsellingRelatedEntityType) > 0 {
		if err := assignInputField(input, "RelatedEntityType", _partnercentralsellingRelatedEntityType); err != nil {
			log.Errorf("invalid --related-entity-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DisassociateOpportunity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a summary of an AWS Opportunity. This summary includes high-level
// details about the opportunity sourced from AWS, such as lifecycle information,
// customer details, and involvement type. It is useful for tracking updates on the
// AWS opportunity corresponding to an opportunity in the partner's account.
func partnercentralselling_GetAwsOpportunitySummary(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.GetAwsOpportunitySummaryInput{
		// Catalog: *string, // Required
		// RelatedOpportunityIdentifier: *string, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingRelatedOpportunityIdentifier) > 0 {
		input.RelatedOpportunityIdentifier = aws.String(_partnercentralsellingRelatedOpportunityIdentifier)
	}

	if resp, err := client.GetAwsOpportunitySummary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this action to retrieve the engagement record for a given
// EngagementIdentifier .
func partnercentralselling_GetEngagement(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.GetEngagementInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralsellingIdentifier)
	}

	if resp, err := client.GetEngagement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details of an engagement invitation shared by AWS with a partner.
// The information includes aspects such as customer, project details, and
// lifecycle information. To connect an engagement invitation with an opportunity,
// match the invitation’s Payload.Project.Title with opportunity Project.Title .
func partnercentralselling_GetEngagementInvitation(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.GetEngagementInvitationInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralsellingIdentifier)
	}

	if resp, err := client.GetEngagementInvitation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Fetches the Opportunity record from Partner Central by a given Identifier .
// Use the ListOpportunities action or the event notification (from Amazon
// EventBridge) to obtain this identifier.
func partnercentralselling_GetOpportunity(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.GetOpportunityInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralsellingIdentifier)
	}

	if resp, err := client.GetOpportunity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this action to retrieve a specific snapshot record.
func partnercentralselling_GetResourceSnapshot(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.GetResourceSnapshotInput{
		// Catalog: *string, // Required
		// EngagementIdentifier: *string, // Required
		// ResourceIdentifier: *string, // Required
		// ResourceSnapshotTemplateIdentifier: *string, // Required
		// ResourceType: types.ResourceType, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingEngagementIdentifier) > 0 {
		input.EngagementIdentifier = aws.String(_partnercentralsellingEngagementIdentifier)
	}
	if len(_partnercentralsellingResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_partnercentralsellingResourceIdentifier)
	}
	if len(_partnercentralsellingResourceSnapshotTemplateIdentifier) > 0 {
		input.ResourceSnapshotTemplateIdentifier = aws.String(_partnercentralsellingResourceSnapshotTemplateIdentifier)
	}
	if len(_partnercentralsellingResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _partnercentralsellingResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingRevision) > 0 {
		if err := assignInputField(input, "Revision", _partnercentralsellingRevision); err != nil {
			log.Errorf("invalid --revision: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetResourceSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this action to retrieves information about a specific resource snapshot job.
func partnercentralselling_GetResourceSnapshotJob(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.GetResourceSnapshotJobInput{
		// Catalog: *string, // Required
		// ResourceSnapshotJobIdentifier: *string, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingResourceSnapshotJobIdentifier) > 0 {
		input.ResourceSnapshotJobIdentifier = aws.String(_partnercentralsellingResourceSnapshotJobIdentifier)
	}

	if resp, err := client.GetResourceSnapshotJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the currently set system settings, which include the IAM Role used
// for resource snapshot jobs.
func partnercentralselling_GetSellingSystemSettings(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.GetSellingSystemSettingsInput{
		// Catalog: *string, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}

	if resp, err := client.GetSellingSystemSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all in-progress, completed, or failed
// StartEngagementByAcceptingInvitationTask tasks that were initiated by the
// caller's account.
func partnercentralselling_ListEngagementByAcceptingInvitationTasks(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.ListEngagementByAcceptingInvitationTasksInput{
		// Catalog: *string, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingEngagementInvitationIdentifier) > 0 {
		input.EngagementInvitationIdentifier = append([]string(nil), _partnercentralsellingEngagementInvitationIdentifier...)
	}
	if len(_partnercentralsellingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _partnercentralsellingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingNextToken) > 0 {
		input.NextToken = aws.String(_partnercentralsellingNextToken)
	}
	if len(_partnercentralsellingOpportunityIdentifier) > 0 {
		input.OpportunityIdentifier = append([]string(nil), _partnercentralsellingOpportunityIdentifier...)
	}
	if len(_partnercentralsellingSort) > 0 {
		if err := assignInputField(input, "Sort", _partnercentralsellingSort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingTaskIdentifier) > 0 {
		input.TaskIdentifier = append([]string(nil), _partnercentralsellingTaskIdentifier...)
	}
	if len(_partnercentralsellingTaskStatus) > 0 {
		if err := assignInputField(input, "TaskStatus", _partnercentralsellingTaskStatus); err != nil {
			log.Errorf("invalid --task-status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListEngagementByAcceptingInvitationTasks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*partnercentralselling.ListEngagementByAcceptingInvitationTasksOutput
	p := partnercentralselling.NewListEngagementByAcceptingInvitationTasksPaginator(client, input)
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

// Lists all in-progress, completed, or failed EngagementFromOpportunity tasks
// that were initiated by the caller's account.
func partnercentralselling_ListEngagementFromOpportunityTasks(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.ListEngagementFromOpportunityTasksInput{
		// Catalog: *string, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingEngagementIdentifier) > 0 {
		input.EngagementIdentifier = []string{_partnercentralsellingEngagementIdentifier}
	}
	if len(_partnercentralsellingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _partnercentralsellingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingNextToken) > 0 {
		input.NextToken = aws.String(_partnercentralsellingNextToken)
	}
	if len(_partnercentralsellingOpportunityIdentifier) > 0 {
		input.OpportunityIdentifier = append([]string(nil), _partnercentralsellingOpportunityIdentifier...)
	}
	if len(_partnercentralsellingSort) > 0 {
		if err := assignInputField(input, "Sort", _partnercentralsellingSort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingTaskIdentifier) > 0 {
		input.TaskIdentifier = append([]string(nil), _partnercentralsellingTaskIdentifier...)
	}
	if len(_partnercentralsellingTaskStatus) > 0 {
		if err := assignInputField(input, "TaskStatus", _partnercentralsellingTaskStatus); err != nil {
			log.Errorf("invalid --task-status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListEngagementFromOpportunityTasks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*partnercentralselling.ListEngagementFromOpportunityTasksOutput
	p := partnercentralselling.NewListEngagementFromOpportunityTasksPaginator(client, input)
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

// Retrieves a list of engagement invitations sent to the partner. This allows
// partners to view all pending or past engagement invitations, helping them track
// opportunities shared by AWS.
func partnercentralselling_ListEngagementInvitations(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.ListEngagementInvitationsInput{
		// Catalog: *string, // Required
		// ParticipantType: types.ParticipantType, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingParticipantType) > 0 {
		if err := assignInputField(input, "ParticipantType", _partnercentralsellingParticipantType); err != nil {
			log.Errorf("invalid --participant-type: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingEngagementIdentifier) > 0 {
		input.EngagementIdentifier = []string{_partnercentralsellingEngagementIdentifier}
	}
	if len(_partnercentralsellingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _partnercentralsellingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingNextToken) > 0 {
		input.NextToken = aws.String(_partnercentralsellingNextToken)
	}
	if len(_partnercentralsellingPayloadType) > 0 {
		if err := assignInputField(input, "PayloadType", _partnercentralsellingPayloadType); err != nil {
			log.Errorf("invalid --payload-type: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingSenderAwsAccountId) > 0 {
		input.SenderAwsAccountId = append([]string(nil), _partnercentralsellingSenderAwsAccountId...)
	}
	if len(_partnercentralsellingSort) > 0 {
		if err := assignInputField(input, "Sort", _partnercentralsellingSort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingStatus) > 0 {
		if err := assignInputField(input, "Status", _partnercentralsellingStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListEngagementInvitations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*partnercentralselling.ListEngagementInvitationsOutput
	p := partnercentralselling.NewListEngagementInvitationsPaginator(client, input)
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

// Retrieves the details of member partners in an Engagement. This operation can
// only be invoked by members of the Engagement. The ListEngagementMembers
// operation allows you to fetch information about the members of a specific
// Engagement. This action is restricted to members of the Engagement being
// queried.
func partnercentralselling_ListEngagementMembers(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.ListEngagementMembersInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralsellingIdentifier)
	}
	if len(_partnercentralsellingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _partnercentralsellingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingNextToken) > 0 {
		input.NextToken = aws.String(_partnercentralsellingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEngagementMembers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*partnercentralselling.ListEngagementMembersOutput
	p := partnercentralselling.NewListEngagementMembersPaginator(client, input)
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

// Lists the associations between resources and engagements where the caller is a
// member and has at least one snapshot in the engagement.
func partnercentralselling_ListEngagementResourceAssociations(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.ListEngagementResourceAssociationsInput{
		// Catalog: *string, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingCreatedBy) > 0 {
		input.CreatedBy = aws.String(_partnercentralsellingCreatedBy)
	}
	if len(_partnercentralsellingEngagementIdentifier) > 0 {
		input.EngagementIdentifier = aws.String(_partnercentralsellingEngagementIdentifier)
	}
	if len(_partnercentralsellingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _partnercentralsellingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingNextToken) > 0 {
		input.NextToken = aws.String(_partnercentralsellingNextToken)
	}
	if len(_partnercentralsellingResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_partnercentralsellingResourceIdentifier)
	}
	if len(_partnercentralsellingResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _partnercentralsellingResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListEngagementResourceAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*partnercentralselling.ListEngagementResourceAssociationsOutput
	p := partnercentralselling.NewListEngagementResourceAssociationsPaginator(client, input)
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

// This action allows users to retrieve a list of Engagement records from Partner
// Central. This action can be used to manage and track various engagements across
// different stages of the partner selling process.
func partnercentralselling_ListEngagements(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.ListEngagementsInput{
		// Catalog: *string, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingContextTypes) > 0 {
		if err := assignInputField(input, "ContextTypes", _partnercentralsellingContextTypes); err != nil {
			log.Errorf("invalid --context-types: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingCreatedBy) > 0 {
		input.CreatedBy = []string{_partnercentralsellingCreatedBy}
	}
	if len(_partnercentralsellingEngagementIdentifier) > 0 {
		input.EngagementIdentifier = []string{_partnercentralsellingEngagementIdentifier}
	}
	if len(_partnercentralsellingExcludeContextTypes) > 0 {
		if err := assignInputField(input, "ExcludeContextTypes", _partnercentralsellingExcludeContextTypes); err != nil {
			log.Errorf("invalid --exclude-context-types: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingExcludeCreatedBy) > 0 {
		input.ExcludeCreatedBy = append([]string(nil), _partnercentralsellingExcludeCreatedBy...)
	}
	if len(_partnercentralsellingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _partnercentralsellingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingNextToken) > 0 {
		input.NextToken = aws.String(_partnercentralsellingNextToken)
	}
	if len(_partnercentralsellingSort) > 0 {
		if err := assignInputField(input, "Sort", _partnercentralsellingSort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
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

	var results []*partnercentralselling.ListEngagementsOutput
	p := partnercentralselling.NewListEngagementsPaginator(client, input)
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

// This request accepts a list of filters that retrieve opportunity subsets as
// well as sort options. This feature is available to partners from [Partner Central]using the
// ListOpportunities API action.
//
// To synchronize your system with Amazon Web Services, list only the
// opportunities that were newly created or updated. We recommend you rely on
// events emitted by the service into your Amazon Web Services account’s Amazon
// EventBridge default event bus. You can also use the ListOpportunities action.
//
// We recommend the following approach:
//
// - Find the latest LastModifiedDate that you stored, and only use the values
// that came from Amazon Web Services. Don’t use values generated by your system.
//
// - When you send a ListOpportunities request, submit the date in ISO 8601
// format in the AfterLastModifiedDate filter.
//
// - Amazon Web Services only returns opportunities created or updated on or
// after that date and time. Use NextToken to iterate over all pages.
//
// [Partner Central]: https://partnercentral.awspartner.com/
func partnercentralselling_ListOpportunities(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.ListOpportunitiesInput{
		// Catalog: *string, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingCreatedDate) > 0 {
		if err := assignInputField(input, "CreatedDate", _partnercentralsellingCreatedDate); err != nil {
			log.Errorf("invalid --created-date: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingCustomerCompanyName) > 0 {
		input.CustomerCompanyName = append([]string(nil), _partnercentralsellingCustomerCompanyName...)
	}
	if len(_partnercentralsellingIdentifier) > 0 {
		input.Identifier = []string{_partnercentralsellingIdentifier}
	}
	if len(_partnercentralsellingLastModifiedDate) > 0 {
		if err := assignInputField(input, "LastModifiedDate", _partnercentralsellingLastModifiedDate); err != nil {
			log.Errorf("invalid --last-modified-date: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingLifeCycleReviewStatus) > 0 {
		if err := assignInputField(input, "LifeCycleReviewStatus", _partnercentralsellingLifeCycleReviewStatus); err != nil {
			log.Errorf("invalid --life-cycle-review-status: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingLifeCycleStage) > 0 {
		if err := assignInputField(input, "LifeCycleStage", _partnercentralsellingLifeCycleStage); err != nil {
			log.Errorf("invalid --life-cycle-stage: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _partnercentralsellingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingNextToken) > 0 {
		input.NextToken = aws.String(_partnercentralsellingNextToken)
	}
	if len(_partnercentralsellingSort) > 0 {
		if err := assignInputField(input, "Sort", _partnercentralsellingSort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingTargetCloseDate) > 0 {
		if err := assignInputField(input, "TargetCloseDate", _partnercentralsellingTargetCloseDate); err != nil {
			log.Errorf("invalid --target-close-date: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListOpportunities(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*partnercentralselling.ListOpportunitiesOutput
	p := partnercentralselling.NewListOpportunitiesPaginator(client, input)
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

// Lists all in-progress, completed, or failed opportunity creation tasks from
// engagements that were initiated by the caller's account.
func partnercentralselling_ListOpportunityFromEngagementTasks(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.ListOpportunityFromEngagementTasksInput{
		// Catalog: *string, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingContextIdentifier) > 0 {
		input.ContextIdentifier = []string{_partnercentralsellingContextIdentifier}
	}
	if len(_partnercentralsellingEngagementIdentifier) > 0 {
		input.EngagementIdentifier = []string{_partnercentralsellingEngagementIdentifier}
	}
	if len(_partnercentralsellingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _partnercentralsellingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingNextToken) > 0 {
		input.NextToken = aws.String(_partnercentralsellingNextToken)
	}
	if len(_partnercentralsellingOpportunityIdentifier) > 0 {
		input.OpportunityIdentifier = append([]string(nil), _partnercentralsellingOpportunityIdentifier...)
	}
	if len(_partnercentralsellingSort) > 0 {
		if err := assignInputField(input, "Sort", _partnercentralsellingSort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingTaskIdentifier) > 0 {
		input.TaskIdentifier = append([]string(nil), _partnercentralsellingTaskIdentifier...)
	}
	if len(_partnercentralsellingTaskStatus) > 0 {
		if err := assignInputField(input, "TaskStatus", _partnercentralsellingTaskStatus); err != nil {
			log.Errorf("invalid --task-status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListOpportunityFromEngagementTasks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*partnercentralselling.ListOpportunityFromEngagementTasksOutput
	p := partnercentralselling.NewListOpportunityFromEngagementTasksPaginator(client, input)
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

// Lists resource snapshot jobs owned by the customer. This operation supports
// various filtering scenarios, including listing all jobs owned by the caller,
// jobs for a specific engagement, jobs with a specific status, or any combination
// of these filters.
func partnercentralselling_ListResourceSnapshotJobs(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.ListResourceSnapshotJobsInput{
		// Catalog: *string, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingEngagementIdentifier) > 0 {
		input.EngagementIdentifier = aws.String(_partnercentralsellingEngagementIdentifier)
	}
	if len(_partnercentralsellingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _partnercentralsellingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingNextToken) > 0 {
		input.NextToken = aws.String(_partnercentralsellingNextToken)
	}
	if len(_partnercentralsellingSort) > 0 {
		if err := assignInputField(input, "Sort", _partnercentralsellingSort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingStatus) > 0 {
		if err := assignInputField(input, "Status", _partnercentralsellingStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListResourceSnapshotJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*partnercentralselling.ListResourceSnapshotJobsOutput
	p := partnercentralselling.NewListResourceSnapshotJobsPaginator(client, input)
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

// Retrieves a list of resource view snapshots based on specified criteria. This
// operation supports various use cases, including:
//
// - Fetching all snapshots associated with an engagement.
//
// - Retrieving snapshots of a specific resource type within an engagement.
//
// - Obtaining snapshots for a particular resource using a specified template.
//
// - Accessing the latest snapshot of a resource within an engagement.
//
// - Filtering snapshots by resource owner.
func partnercentralselling_ListResourceSnapshots(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.ListResourceSnapshotsInput{
		// Catalog: *string, // Required
		// EngagementIdentifier: *string, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingEngagementIdentifier) > 0 {
		input.EngagementIdentifier = aws.String(_partnercentralsellingEngagementIdentifier)
	}
	if len(_partnercentralsellingCreatedBy) > 0 {
		input.CreatedBy = aws.String(_partnercentralsellingCreatedBy)
	}
	if len(_partnercentralsellingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _partnercentralsellingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingNextToken) > 0 {
		input.NextToken = aws.String(_partnercentralsellingNextToken)
	}
	if len(_partnercentralsellingResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_partnercentralsellingResourceIdentifier)
	}
	if len(_partnercentralsellingResourceSnapshotTemplateIdentifier) > 0 {
		input.ResourceSnapshotTemplateIdentifier = aws.String(_partnercentralsellingResourceSnapshotTemplateIdentifier)
	}
	if len(_partnercentralsellingResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _partnercentralsellingResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListResourceSnapshots(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*partnercentralselling.ListResourceSnapshotsOutput
	p := partnercentralselling.NewListResourceSnapshotsPaginator(client, input)
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

// Retrieves a list of Partner Solutions that the partner registered on Partner
// Central. This API is used to generate a list of solutions that an end user
// selects from for association with an opportunity.
func partnercentralselling_ListSolutions(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.ListSolutionsInput{
		// Catalog: *string, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingCategory) > 0 {
		input.Category = append([]string(nil), _partnercentralsellingCategory...)
	}
	if len(_partnercentralsellingIdentifier) > 0 {
		input.Identifier = []string{_partnercentralsellingIdentifier}
	}
	if len(_partnercentralsellingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _partnercentralsellingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingNextToken) > 0 {
		input.NextToken = aws.String(_partnercentralsellingNextToken)
	}
	if len(_partnercentralsellingSort) > 0 {
		if err := assignInputField(input, "Sort", _partnercentralsellingSort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingStatus) > 0 {
		if err := assignInputField(input, "Status", _partnercentralsellingStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSolutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*partnercentralselling.ListSolutionsOutput
	p := partnercentralselling.NewListSolutionsPaginator(client, input)
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

// Returns a list of tags for a resource.
func partnercentralselling_ListTagsForResource(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_partnercentralsellingResourceArn) > 0 {
		input.ResourceArn = aws.String(_partnercentralsellingResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the currently set system settings, which include the IAM Role used for
// resource snapshot jobs.
func partnercentralselling_PutSellingSystemSettings(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.PutSellingSystemSettingsInput{
		// Catalog: *string, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingResourceSnapshotJobRoleIdentifier) > 0 {
		input.ResourceSnapshotJobRoleIdentifier = aws.String(_partnercentralsellingResourceSnapshotJobRoleIdentifier)
	}

	if resp, err := client.PutSellingSystemSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action rejects an EngagementInvitation that AWS shared. Rejecting an
// invitation indicates that the partner doesn't want to pursue the opportunity,
// and all related data will become inaccessible thereafter.
func partnercentralselling_RejectEngagementInvitation(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.RejectEngagementInvitationInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralsellingIdentifier)
	}
	if len(_partnercentralsellingRejectionReason) > 0 {
		input.RejectionReason = aws.String(_partnercentralsellingRejectionReason)
	}

	if resp, err := client.RejectEngagementInvitation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action starts the engagement by accepting an EngagementInvitation . The
// task is asynchronous and involves the following steps: accepting the invitation,
// creating an opportunity in the partner’s account from the AWS opportunity, and
// copying details for tracking. When completed, an Opportunity Created event is
// generated, indicating that the opportunity has been successfully created in the
// partner's account.
func partnercentralselling_StartEngagementByAcceptingInvitationTask(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.StartEngagementByAcceptingInvitationTaskInput{
		// Catalog: *string, // Required
		// ClientToken: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralsellingClientToken)
	}
	if len(_partnercentralsellingIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralsellingIdentifier)
	}
	if len(_partnercentralsellingTags) > 0 {
		if err := assignInputField(input, "Tags", _partnercentralsellingTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartEngagementByAcceptingInvitationTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Similar to StartEngagementByAcceptingInvitationTask , this action is
// asynchronous and performs multiple steps before completion. This action
// orchestrates a comprehensive workflow that combines multiple API operations into
// a single task to create and initiate an engagement from an existing opportunity.
// It automatically executes a sequence of operations including GetOpportunity ,
// CreateEngagement (if it doesn't exist), CreateResourceSnapshot ,
// CreateResourceSnapshotJob , CreateEngagementInvitation (if not already
// invited/accepted), and SubmitOpportunity .
func partnercentralselling_StartEngagementFromOpportunityTask(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.StartEngagementFromOpportunityTaskInput{
		// AwsSubmission: *types.AwsSubmission, // Required
		// Catalog: *string, // Required
		// ClientToken: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralsellingAwsSubmission) > 0 {
		if err := assignInputField(input, "AwsSubmission", _partnercentralsellingAwsSubmission); err != nil {
			log.Errorf("invalid --aws-submission: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralsellingClientToken)
	}
	if len(_partnercentralsellingIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralsellingIdentifier)
	}
	if len(_partnercentralsellingTags) > 0 {
		if err := assignInputField(input, "Tags", _partnercentralsellingTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartEngagementFromOpportunityTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action creates an opportunity from an existing engagement context. The
// task is asynchronous and orchestrates the process of converting engagement
// contextual information into a structured opportunity record within the partner's
// account.
func partnercentralselling_StartOpportunityFromEngagementTask(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.StartOpportunityFromEngagementTaskInput{
		// Catalog: *string, // Required
		// ClientToken: *string, // Required
		// ContextIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralsellingClientToken)
	}
	if len(_partnercentralsellingContextIdentifier) > 0 {
		input.ContextIdentifier = aws.String(_partnercentralsellingContextIdentifier)
	}
	if len(_partnercentralsellingIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralsellingIdentifier)
	}
	if len(_partnercentralsellingTags) > 0 {
		if err := assignInputField(input, "Tags", _partnercentralsellingTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartOpportunityFromEngagementTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a resource snapshot job that has been previously created.
func partnercentralselling_StartResourceSnapshotJob(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.StartResourceSnapshotJobInput{
		// Catalog: *string, // Required
		// ResourceSnapshotJobIdentifier: *string, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingResourceSnapshotJobIdentifier) > 0 {
		input.ResourceSnapshotJobIdentifier = aws.String(_partnercentralsellingResourceSnapshotJobIdentifier)
	}

	if resp, err := client.StartResourceSnapshotJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a resource snapshot job. The job must be started prior to being stopped.
func partnercentralselling_StopResourceSnapshotJob(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.StopResourceSnapshotJobInput{
		// Catalog: *string, // Required
		// ResourceSnapshotJobIdentifier: *string, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingResourceSnapshotJobIdentifier) > 0 {
		input.ResourceSnapshotJobIdentifier = aws.String(_partnercentralsellingResourceSnapshotJobIdentifier)
	}

	if resp, err := client.StopResourceSnapshotJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this action to submit an Opportunity that was previously created by partner
// for AWS review. After you perform this action, the Opportunity becomes
// non-editable until it is reviewed by AWS and has LifeCycle.ReviewStatus  as
// either Approved or Action Required .
func partnercentralselling_SubmitOpportunity(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.SubmitOpportunityInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
		// InvolvementType: types.SalesInvolvementType, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralsellingIdentifier)
	}
	if len(_partnercentralsellingInvolvementType) > 0 {
		if err := assignInputField(input, "InvolvementType", _partnercentralsellingInvolvementType); err != nil {
			log.Errorf("invalid --involvement-type: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingVisibility) > 0 {
		if err := assignInputField(input, "Visibility", _partnercentralsellingVisibility); err != nil {
			log.Errorf("invalid --visibility: %s", err.Error())
			return
		}
	}

	if resp, err := client.SubmitOpportunity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns one or more tags (key-value pairs) to the specified resource.
func partnercentralselling_TagResource(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_partnercentralsellingResourceArn) > 0 {
		input.ResourceArn = aws.String(_partnercentralsellingResourceArn)
	}
	if len(_partnercentralsellingTags) > 0 {
		if err := assignInputField(input, "Tags", _partnercentralsellingTags); err != nil {
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

// Removes a tag or tags from a resource.
func partnercentralselling_UntagResource(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_partnercentralsellingResourceArn) > 0 {
		input.ResourceArn = aws.String(_partnercentralsellingResourceArn)
	}
	if len(_partnercentralsellingTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _partnercentralsellingTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the context information for an existing engagement with new or modified
// data.
func partnercentralselling_UpdateEngagementContext(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.UpdateEngagementContextInput{
		// Catalog: *string, // Required
		// ContextIdentifier: *string, // Required
		// EngagementIdentifier: *string, // Required
		// EngagementLastModifiedAt: *time.Time, // Required
		// Payload: types.UpdateEngagementContextPayload, // Required
		// Type: types.EngagementContextType, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingContextIdentifier) > 0 {
		input.ContextIdentifier = aws.String(_partnercentralsellingContextIdentifier)
	}
	if len(_partnercentralsellingEngagementIdentifier) > 0 {
		input.EngagementIdentifier = aws.String(_partnercentralsellingEngagementIdentifier)
	}
	if len(_partnercentralsellingEngagementLastModifiedAt) > 0 {
		if err := assignInputField(input, "EngagementLastModifiedAt", _partnercentralsellingEngagementLastModifiedAt); err != nil {
			log.Errorf("invalid --engagement-last-modified-at: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingPayload) > 0 {
		if err := assignInputField(input, "Payload", _partnercentralsellingPayload); err != nil {
			log.Errorf("invalid --payload: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingType) > 0 {
		if err := assignInputField(input, "Type", _partnercentralsellingType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEngagementContext(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the Opportunity record identified by a given Identifier . This operation
// allows you to modify the details of an existing opportunity to reflect the
// latest information and progress. Use this action to keep the opportunity record
// up-to-date and accurate.
//
// When you perform updates, include the entire payload with each request. If any
// field is omitted, the API assumes that the field is set to null . The best
// practice is to always perform a GetOpportunity to retrieve the latest values,
// then send the complete payload with the updated values to be changed.
func partnercentralselling_UpdateOpportunity(cfg aws.Config, client *partnercentralselling.Client) {
	input := &partnercentralselling.UpdateOpportunityInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
		// LastModifiedDate: *time.Time, // Required
	}

	if len(_partnercentralsellingCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralsellingCatalog)
	}
	if len(_partnercentralsellingIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralsellingIdentifier)
	}
	if len(_partnercentralsellingLastModifiedDate) > 0 {
		if err := assignInputField(input, "LastModifiedDate", _partnercentralsellingLastModifiedDate); err != nil {
			log.Errorf("invalid --last-modified-date: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingCustomer) > 0 {
		if err := assignInputField(input, "Customer", _partnercentralsellingCustomer); err != nil {
			log.Errorf("invalid --customer: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingLifeCycle) > 0 {
		if err := assignInputField(input, "LifeCycle", _partnercentralsellingLifeCycle); err != nil {
			log.Errorf("invalid --life-cycle: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingMarketing) > 0 {
		if err := assignInputField(input, "Marketing", _partnercentralsellingMarketing); err != nil {
			log.Errorf("invalid --marketing: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingNationalSecurity) > 0 {
		if err := assignInputField(input, "NationalSecurity", _partnercentralsellingNationalSecurity); err != nil {
			log.Errorf("invalid --national-security: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingOpportunityType) > 0 {
		if err := assignInputField(input, "OpportunityType", _partnercentralsellingOpportunityType); err != nil {
			log.Errorf("invalid --opportunity-type: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingPartnerOpportunityIdentifier) > 0 {
		input.PartnerOpportunityIdentifier = aws.String(_partnercentralsellingPartnerOpportunityIdentifier)
	}
	if len(_partnercentralsellingPrimaryNeedsFromAws) > 0 {
		if err := assignInputField(input, "PrimaryNeedsFromAws", _partnercentralsellingPrimaryNeedsFromAws); err != nil {
			log.Errorf("invalid --primary-needs-from-aws: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingProject) > 0 {
		if err := assignInputField(input, "Project", _partnercentralsellingProject); err != nil {
			log.Errorf("invalid --project: %s", err.Error())
			return
		}
	}
	if len(_partnercentralsellingSoftwareRevenue) > 0 {
		if err := assignInputField(input, "SoftwareRevenue", _partnercentralsellingSoftwareRevenue); err != nil {
			log.Errorf("invalid --software-revenue: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateOpportunity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_partnercentralsellingCmd)
	_partnercentralsellingCmd.Flags().SortFlags = false

	_partnercentralsellingCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_partnercentralsellingCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_partnercentralsellingCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingAssignee, "assignee", "", "", "Assignee")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingAwsSubmission, "aws-submission", "", "", "AWS Submission")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingCatalog, "catalog", "", "", "Catalog")
	_partnercentralsellingCmd.Flags().StringSliceVarP(&_partnercentralsellingCategory, "category", "", nil, "Category")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingClientToken, "client-token", "", "", "Client Token")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingContextIdentifier, "context-identifier", "", "", "Context Identifier")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingContextTypes, "context-types", "", "", "Context Types")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingContexts, "contexts", "", "", "Contexts")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingCreatedBy, "created-by", "", "", "Created By")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingCreatedDate, "created-date", "", "", "Created Date")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingCustomer, "customer", "", "", "Customer")
	_partnercentralsellingCmd.Flags().StringSliceVarP(&_partnercentralsellingCustomerCompanyName, "customer-company-name", "", nil, "Customer Company Name")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingDescription, "description", "", "", "Description")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingEngagementIdentifier, "engagement-identifier", "", "", "Engagement Identifier")
	_partnercentralsellingCmd.Flags().StringSliceVarP(&_partnercentralsellingEngagementInvitationIdentifier, "engagement-invitation-identifier", "", nil, "Engagement Invitation Identifier")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingEngagementLastModifiedAt, "engagement-last-modified-at", "", "", "Engagement Last Modified At")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingExcludeContextTypes, "exclude-context-types", "", "", "Exclude Context Types")
	_partnercentralsellingCmd.Flags().StringSliceVarP(&_partnercentralsellingExcludeCreatedBy, "exclude-created-by", "", nil, "Exclude Created By")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingIdentifier, "identifier", "", "", "Identifier")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingInvitation, "invitation", "", "", "Invitation")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingInvolvementType, "involvement-type", "", "", "Involvement Type")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingLastModifiedDate, "last-modified-date", "", "", "Last Modified Date")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingLifeCycle, "life-cycle", "", "", "Life Cycle")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingLifeCycleReviewStatus, "life-cycle-review-status", "", "", "Life Cycle Review Status")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingLifeCycleStage, "life-cycle-stage", "", "", "Life Cycle Stage")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingMarketing, "marketing", "", "", "Marketing")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingMaxResults, "max-results", "", "", "Max Results")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingNationalSecurity, "national-security", "", "", "National Security")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingNextToken, "next-token", "", "", "Next Token")
	_partnercentralsellingCmd.Flags().StringSliceVarP(&_partnercentralsellingOpportunityIdentifier, "opportunity-identifier", "", nil, "Opportunity Identifier")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingOpportunityTeam, "opportunity-team", "", "", "Opportunity Team")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingOpportunityType, "opportunity-type", "", "", "Opportunity Type")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingOrigin, "origin", "", "", "Origin")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingParticipantType, "participant-type", "", "", "Participant Type")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingPartnerOpportunityIdentifier, "partner-opportunity-identifier", "", "", "Partner Opportunity Identifier")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingPayload, "payload", "", "", "Payload")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingPayloadType, "payload-type", "", "", "Payload Type")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingPrimaryNeedsFromAws, "primary-needs-from-aws", "", "", "Primary Needs From AWS")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingProject, "project", "", "", "Project")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingRejectionReason, "rejection-reason", "", "", "Rejection Reason")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingRelatedEntityIdentifier, "related-entity-identifier", "", "", "Related Entity Identifier")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingRelatedEntityType, "related-entity-type", "", "", "Related Entity Type")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingRelatedOpportunityIdentifier, "related-opportunity-identifier", "", "", "Related Opportunity Identifier")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingResourceArn, "resource-arn", "", "", "Resource ARN")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingResourceIdentifier, "resource-identifier", "", "", "Resource Identifier")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingResourceSnapshotJobIdentifier, "resource-snapshot-job-identifier", "", "", "Resource Snapshot Job Identifier")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingResourceSnapshotJobRoleIdentifier, "resource-snapshot-job-role-identifier", "", "", "Resource Snapshot Job Role Identifier")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingResourceSnapshotTemplateIdentifier, "resource-snapshot-template-identifier", "", "", "Resource Snapshot Template Identifier")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingResourceType, "resource-type", "", "", "Resource Type")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingRevision, "revision", "", "", "Revision")
	_partnercentralsellingCmd.Flags().StringSliceVarP(&_partnercentralsellingSenderAwsAccountId, "sender-aws-account-id", "", nil, "Sender AWS Account ID")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingSoftwareRevenue, "software-revenue", "", "", "Software Revenue")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingSort, "sort", "", "", "Sort")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingStatus, "status", "", "", "Status")
	_partnercentralsellingCmd.Flags().StringSliceVarP(&_partnercentralsellingTagKeys, "tag-keys", "", nil, "Tag Keys")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingTags, "tags", "", "", "Tags")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingTargetCloseDate, "target-close-date", "", "", "Target Close Date")
	_partnercentralsellingCmd.Flags().StringSliceVarP(&_partnercentralsellingTaskIdentifier, "task-identifier", "", nil, "Task Identifier")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingTaskStatus, "task-status", "", "", "Task Status")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingTitle, "title", "", "", "Title")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingType, "type", "", "", "Type")
	_partnercentralsellingCmd.Flags().StringVarP(&_partnercentralsellingVisibility, "visibility", "", "", "Visibility")

	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingAcceptEngagementInvitation, "accept-engagement-invitation", "", false, "Accept Engagement Invitation")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingAssignOpportunity, "assign-opportunity", "", false, "Assign Opportunity")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingAssociateOpportunity, "associate-opportunity", "", false, "Associate Opportunity")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingCreateEngagement, "create-engagement", "", false, "Create Engagement")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingCreateEngagementContext, "create-engagement-context", "", false, "Create Engagement Context")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingCreateEngagementInvitation, "create-engagement-invitation", "", false, "Create Engagement Invitation")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingCreateOpportunity, "create-opportunity", "", false, "Create Opportunity")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingCreateResourceSnapshot, "create-resource-snapshot", "", false, "Create Resource Snapshot")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingCreateResourceSnapshotJob, "create-resource-snapshot-job", "", false, "Create Resource Snapshot Job")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingDeleteResourceSnapshotJob, "delete-resource-snapshot-job", "", false, "Delete Resource Snapshot Job")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingDisassociateOpportunity, "disassociate-opportunity", "", false, "Disassociate Opportunity")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingGetAwsOpportunitySummary, "get-aws-opportunity-summary", "", false, "Get AWS Opportunity Summary")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingGetEngagement, "get-engagement", "", false, "Get Engagement")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingGetEngagementInvitation, "get-engagement-invitation", "", false, "Get Engagement Invitation")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingGetOpportunity, "get-opportunity", "", false, "Get Opportunity")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingGetResourceSnapshot, "get-resource-snapshot", "", false, "Get Resource Snapshot")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingGetResourceSnapshotJob, "get-resource-snapshot-job", "", false, "Get Resource Snapshot Job")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingGetSellingSystemSettings, "get-selling-system-settings", "", false, "Get Selling System Settings")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingListEngagementByAcceptingInvitationTasks, "list-engagement-by-accepting-invitation-tasks", "", false, "List Engagement By Accepting Invitation Tasks")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingListEngagementFromOpportunityTasks, "list-engagement-from-opportunity-tasks", "", false, "List Engagement From Opportunity Tasks")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingListEngagementInvitations, "list-engagement-invitations", "", false, "List Engagement Invitations")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingListEngagementMembers, "list-engagement-members", "", false, "List Engagement Members")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingListEngagementResourceAssociations, "list-engagement-resource-associations", "", false, "List Engagement Resource Associations")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingListEngagements, "list-engagements", "", false, "List Engagements")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingListOpportunities, "list-opportunities", "", false, "List Opportunities")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingListOpportunityFromEngagementTasks, "list-opportunity-from-engagement-tasks", "", false, "List Opportunity From Engagement Tasks")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingListResourceSnapshotJobs, "list-resource-snapshot-jobs", "", false, "List Resource Snapshot Jobs")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingListResourceSnapshots, "list-resource-snapshots", "", false, "List Resource Snapshots")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingListSolutions, "list-solutions", "", false, "List Solutions")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingPutSellingSystemSettings, "put-selling-system-settings", "", false, "Put Selling System Settings")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingRejectEngagementInvitation, "reject-engagement-invitation", "", false, "Reject Engagement Invitation")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingStartEngagementByAcceptingInvitationTask, "start-engagement-by-accepting-invitation-task", "", false, "Start Engagement By Accepting Invitation Task")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingStartEngagementFromOpportunityTask, "start-engagement-from-opportunity-task", "", false, "Start Engagement From Opportunity Task")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingStartOpportunityFromEngagementTask, "start-opportunity-from-engagement-task", "", false, "Start Opportunity From Engagement Task")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingStartResourceSnapshotJob, "start-resource-snapshot-job", "", false, "Start Resource Snapshot Job")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingStopResourceSnapshotJob, "stop-resource-snapshot-job", "", false, "Stop Resource Snapshot Job")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingSubmitOpportunity, "submit-opportunity", "", false, "Submit Opportunity")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingTagResource, "tag-resource", "", false, "Tag Resource")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingUntagResource, "untag-resource", "", false, "Untag Resource")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingUpdateEngagementContext, "update-engagement-context", "", false, "Update Engagement Context")
	_partnercentralsellingCmd.Flags().BoolVarP(&_partnercentralsellingUpdateOpportunity, "update-opportunity", "", false, "Update Opportunity")

}
