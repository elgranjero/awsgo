package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/billing"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// billingCmd represents the billing command
var _billingCmd = &cobra.Command{
	Use:   "billing",
	Short: "AWS billing CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := billing.NewFromConfig(cfg)
		if _billingAssociateSourceViews {
			billing_AssociateSourceViews(cfg, client)
			return
		}
		if _billingCreateBillingView {
			billing_CreateBillingView(cfg, client)
			return
		}
		if _billingDeleteBillingView {
			billing_DeleteBillingView(cfg, client)
			return
		}
		if _billingDisassociateSourceViews {
			billing_DisassociateSourceViews(cfg, client)
			return
		}
		if _billingGetBillingView {
			billing_GetBillingView(cfg, client)
			return
		}
		if _billingGetResourcePolicy {
			billing_GetResourcePolicy(cfg, client)
			return
		}
		if _billingListBillingViews {
			billing_ListBillingViews(cfg, client)
			return
		}
		if _billingListSourceViewsForBillingView {
			billing_ListSourceViewsForBillingView(cfg, client)
			return
		}
		if _billingListTagsForResource {
			billing_ListTagsForResource(cfg, client)
			return
		}
		if _billingTagResource {
			billing_TagResource(cfg, client)
			return
		}
		if _billingUntagResource {
			billing_UntagResource(cfg, client)
			return
		}
		if _billingUpdateBillingView {
			billing_UpdateBillingView(cfg, client)
			return
		}

	},
}

var (
	_billingAssociateSourceViews          bool
	_billingCreateBillingView             bool
	_billingDeleteBillingView             bool
	_billingDisassociateSourceViews       bool
	_billingGetBillingView                bool
	_billingGetResourcePolicy             bool
	_billingListBillingViews              bool
	_billingListSourceViewsForBillingView bool
	_billingListTagsForResource           bool
	_billingTagResource                   bool
	_billingUntagResource                 bool
	_billingUpdateBillingView             bool

	_billingActiveTimeRange      string
	_billingArn                  string
	_billingArns                 []string
	_billingBillingViewTypes     string
	_billingClientToken          string
	_billingDataFilterExpression string
	_billingDescription          string
	_billingForce                string
	_billingMaxResults           string
	_billingName                 string
	_billingNames                string
	_billingNextToken            string
	_billingOwnerAccountId       string
	_billingResourceArn          string
	_billingResourceTagKeys      []string
	_billingResourceTags         string
	_billingSourceAccountId      string
	_billingSourceViews          []string
)

// Associates one or more source billing views with an existing billing view.
// This allows creating aggregate billing views that combine data from multiple
// sources.
func billing_AssociateSourceViews(cfg aws.Config, client *billing.Client) {
	input := &billing.AssociateSourceViewsInput{
		// Arn: *string, // Required
		// SourceViews: []string, // Required
	}

	if len(_billingArn) > 0 {
		input.Arn = aws.String(_billingArn)
	}
	if len(_billingSourceViews) > 0 {
		input.SourceViews = append([]string(nil), _billingSourceViews...)
	}

	if resp, err := client.AssociateSourceViews(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a billing view with the specified billing view attributes.
func billing_CreateBillingView(cfg aws.Config, client *billing.Client) {
	input := &billing.CreateBillingViewInput{
		// Name: *string, // Required
		// SourceViews: []string, // Required
	}

	if len(_billingName) > 0 {
		input.Name = aws.String(_billingName)
	}
	if len(_billingSourceViews) > 0 {
		input.SourceViews = append([]string(nil), _billingSourceViews...)
	}
	if len(_billingClientToken) > 0 {
		input.ClientToken = aws.String(_billingClientToken)
	}
	if len(_billingDataFilterExpression) > 0 {
		if err := assignInputField(input, "DataFilterExpression", _billingDataFilterExpression); err != nil {
			log.Errorf("invalid --data-filter-expression: %s", err.Error())
			return
		}
	}
	if len(_billingDescription) > 0 {
		input.Description = aws.String(_billingDescription)
	}
	if len(_billingResourceTags) > 0 {
		if err := assignInputField(input, "ResourceTags", _billingResourceTags); err != nil {
			log.Errorf("invalid --resource-tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBillingView(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified billing view.
func billing_DeleteBillingView(cfg aws.Config, client *billing.Client) {
	input := &billing.DeleteBillingViewInput{
		// Arn: *string, // Required
	}

	if len(_billingArn) > 0 {
		input.Arn = aws.String(_billingArn)
	}
	if len(_billingForce) > 0 {
		if err := assignInputField(input, "Force", _billingForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteBillingView(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the association between one or more source billing views and an
// existing billing view. This allows modifying the composition of aggregate
// billing views.
func billing_DisassociateSourceViews(cfg aws.Config, client *billing.Client) {
	input := &billing.DisassociateSourceViewsInput{
		// Arn: *string, // Required
		// SourceViews: []string, // Required
	}

	if len(_billingArn) > 0 {
		input.Arn = aws.String(_billingArn)
	}
	if len(_billingSourceViews) > 0 {
		input.SourceViews = append([]string(nil), _billingSourceViews...)
	}

	if resp, err := client.DisassociateSourceViews(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the metadata associated to the specified billing view ARN.
func billing_GetBillingView(cfg aws.Config, client *billing.Client) {
	input := &billing.GetBillingViewInput{
		// Arn: *string, // Required
	}

	if len(_billingArn) > 0 {
		input.Arn = aws.String(_billingArn)
	}

	if resp, err := client.GetBillingView(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the resource-based policy document attached to the resource in JSON
// format.
func billing_GetResourcePolicy(cfg aws.Config, client *billing.Client) {
	input := &billing.GetResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_billingResourceArn) > 0 {
		input.ResourceArn = aws.String(_billingResourceArn)
	}

	if resp, err := client.GetResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the billing views available for a given time period.
// Every Amazon Web Services account has a unique PRIMARY billing view that
// represents the billing data available by default. Accounts that use Billing
// Conductor also have BILLING_GROUP billing views representing pro forma costs
// associated with each created billing group.
func billing_ListBillingViews(cfg aws.Config, client *billing.Client) {
	input := &billing.ListBillingViewsInput{}

	if len(_billingActiveTimeRange) > 0 {
		if err := assignInputField(input, "ActiveTimeRange", _billingActiveTimeRange); err != nil {
			log.Errorf("invalid --active-time-range: %s", err.Error())
			return
		}
	}
	if len(_billingArns) > 0 {
		input.Arns = append([]string(nil), _billingArns...)
	}
	if len(_billingBillingViewTypes) > 0 {
		if err := assignInputField(input, "BillingViewTypes", _billingBillingViewTypes); err != nil {
			log.Errorf("invalid --billing-view-types: %s", err.Error())
			return
		}
	}
	if len(_billingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _billingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_billingNames) > 0 {
		if err := assignInputField(input, "Names", _billingNames); err != nil {
			log.Errorf("invalid --names: %s", err.Error())
			return
		}
	}
	if len(_billingNextToken) > 0 {
		input.NextToken = aws.String(_billingNextToken)
	}
	if len(_billingOwnerAccountId) > 0 {
		input.OwnerAccountId = aws.String(_billingOwnerAccountId)
	}
	if len(_billingSourceAccountId) > 0 {
		input.SourceAccountId = aws.String(_billingSourceAccountId)
	}

	if disablePaginator() {
		if resp, err := client.ListBillingViews(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*billing.ListBillingViewsOutput
	p := billing.NewListBillingViewsPaginator(client, input)
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

// Lists the source views (managed Amazon Web Services billing views) associated
// with the billing view.
func billing_ListSourceViewsForBillingView(cfg aws.Config, client *billing.Client) {
	input := &billing.ListSourceViewsForBillingViewInput{
		// Arn: *string, // Required
	}

	if len(_billingArn) > 0 {
		input.Arn = aws.String(_billingArn)
	}
	if len(_billingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _billingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_billingNextToken) > 0 {
		input.NextToken = aws.String(_billingNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSourceViewsForBillingView(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*billing.ListSourceViewsForBillingViewOutput
	p := billing.NewListSourceViewsForBillingViewPaginator(client, input)
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

// Lists tags associated with the billing view resource.
func billing_ListTagsForResource(cfg aws.Config, client *billing.Client) {
	input := &billing.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_billingResourceArn) > 0 {
		input.ResourceArn = aws.String(_billingResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// An API operation for adding one or more tags (key-value pairs) to a resource.
func billing_TagResource(cfg aws.Config, client *billing.Client) {
	input := &billing.TagResourceInput{
		// ResourceArn: *string, // Required
		// ResourceTags: []types.ResourceTag, // Required
	}

	if len(_billingResourceArn) > 0 {
		input.ResourceArn = aws.String(_billingResourceArn)
	}
	if len(_billingResourceTags) > 0 {
		if err := assignInputField(input, "ResourceTags", _billingResourceTags); err != nil {
			log.Errorf("invalid --resource-tags: %s", err.Error())
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

// Removes one or more tags from a resource. Specify only tag keys in your
// request. Don't specify the value.
func billing_UntagResource(cfg aws.Config, client *billing.Client) {
	input := &billing.UntagResourceInput{
		// ResourceArn: *string, // Required
		// ResourceTagKeys: []string, // Required
	}

	if len(_billingResourceArn) > 0 {
		input.ResourceArn = aws.String(_billingResourceArn)
	}
	if len(_billingResourceTagKeys) > 0 {
		input.ResourceTagKeys = append([]string(nil), _billingResourceTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// An API to update the attributes of the billing view.
func billing_UpdateBillingView(cfg aws.Config, client *billing.Client) {
	input := &billing.UpdateBillingViewInput{
		// Arn: *string, // Required
	}

	if len(_billingArn) > 0 {
		input.Arn = aws.String(_billingArn)
	}
	if len(_billingDataFilterExpression) > 0 {
		if err := assignInputField(input, "DataFilterExpression", _billingDataFilterExpression); err != nil {
			log.Errorf("invalid --data-filter-expression: %s", err.Error())
			return
		}
	}
	if len(_billingDescription) > 0 {
		input.Description = aws.String(_billingDescription)
	}
	if len(_billingName) > 0 {
		input.Name = aws.String(_billingName)
	}

	if resp, err := client.UpdateBillingView(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_billingCmd)
	_billingCmd.Flags().SortFlags = false

	_billingCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_billingCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_billingCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_billingCmd.Flags().StringVarP(&_billingActiveTimeRange, "active-time-range", "", "", "Active Time Range")
	_billingCmd.Flags().StringVarP(&_billingArn, "arn", "", "", "ARN")
	_billingCmd.Flags().StringSliceVarP(&_billingArns, "arns", "", nil, "Arns")
	_billingCmd.Flags().StringVarP(&_billingBillingViewTypes, "billing-view-types", "", "", "Billing View Types")
	_billingCmd.Flags().StringVarP(&_billingClientToken, "client-token", "", "", "Client Token")
	_billingCmd.Flags().StringVarP(&_billingDataFilterExpression, "data-filter-expression", "", "", "Data Filter Expression")
	_billingCmd.Flags().StringVarP(&_billingDescription, "description", "", "", "Description")
	_billingCmd.Flags().StringVarP(&_billingForce, "force", "", "", "Force")
	_billingCmd.Flags().StringVarP(&_billingMaxResults, "max-results", "", "", "Max Results")
	_billingCmd.Flags().StringVarP(&_billingName, "name", "", "", "Name")
	_billingCmd.Flags().StringVarP(&_billingNames, "names", "", "", "Names")
	_billingCmd.Flags().StringVarP(&_billingNextToken, "next-token", "", "", "Next Token")
	_billingCmd.Flags().StringVarP(&_billingOwnerAccountId, "owner-account-id", "", "", "Owner Account ID")
	_billingCmd.Flags().StringVarP(&_billingResourceArn, "resource-arn", "", "", "Resource ARN")
	_billingCmd.Flags().StringSliceVarP(&_billingResourceTagKeys, "resource-tag-keys", "", nil, "Resource Tag Keys")
	_billingCmd.Flags().StringVarP(&_billingResourceTags, "resource-tags", "", "", "Resource Tags")
	_billingCmd.Flags().StringVarP(&_billingSourceAccountId, "source-account-id", "", "", "Source Account ID")
	_billingCmd.Flags().StringSliceVarP(&_billingSourceViews, "source-views", "", nil, "Source Views")

	_billingCmd.Flags().BoolVarP(&_billingAssociateSourceViews, "associate-source-views", "", false, "Associate Source Views")
	_billingCmd.Flags().BoolVarP(&_billingCreateBillingView, "create-billing-view", "", false, "Create Billing View")
	_billingCmd.Flags().BoolVarP(&_billingDeleteBillingView, "delete-billing-view", "", false, "Delete Billing View")
	_billingCmd.Flags().BoolVarP(&_billingDisassociateSourceViews, "disassociate-source-views", "", false, "Disassociate Source Views")
	_billingCmd.Flags().BoolVarP(&_billingGetBillingView, "get-billing-view", "", false, "Get Billing View")
	_billingCmd.Flags().BoolVarP(&_billingGetResourcePolicy, "get-resource-policy", "", false, "Get Resource Policy")
	_billingCmd.Flags().BoolVarP(&_billingListBillingViews, "list-billing-views", "", false, "List Billing Views")
	_billingCmd.Flags().BoolVarP(&_billingListSourceViewsForBillingView, "list-source-views-for-billing-view", "", false, "List Source Views For Billing View")
	_billingCmd.Flags().BoolVarP(&_billingListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_billingCmd.Flags().BoolVarP(&_billingTagResource, "tag-resource", "", false, "Tag Resource")
	_billingCmd.Flags().BoolVarP(&_billingUntagResource, "untag-resource", "", false, "Untag Resource")
	_billingCmd.Flags().BoolVarP(&_billingUpdateBillingView, "update-billing-view", "", false, "Update Billing View")

}
