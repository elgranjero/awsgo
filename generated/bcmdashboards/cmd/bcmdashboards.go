package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bcmdashboards"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// bcmdashboardsCmd represents the bcmdashboards command
var _bcmdashboardsCmd = &cobra.Command{
	Use:   "bcmdashboards",
	Short: "AWS bcmdashboards CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := bcmdashboards.NewFromConfig(cfg)
		if _bcmdashboardsCreateDashboard {
			bcmdashboards_CreateDashboard(cfg, client)
			return
		}
		if _bcmdashboardsDeleteDashboard {
			bcmdashboards_DeleteDashboard(cfg, client)
			return
		}
		if _bcmdashboardsGetDashboard {
			bcmdashboards_GetDashboard(cfg, client)
			return
		}
		if _bcmdashboardsGetResourcePolicy {
			bcmdashboards_GetResourcePolicy(cfg, client)
			return
		}
		if _bcmdashboardsListDashboards {
			bcmdashboards_ListDashboards(cfg, client)
			return
		}
		if _bcmdashboardsListTagsForResource {
			bcmdashboards_ListTagsForResource(cfg, client)
			return
		}
		if _bcmdashboardsTagResource {
			bcmdashboards_TagResource(cfg, client)
			return
		}
		if _bcmdashboardsUntagResource {
			bcmdashboards_UntagResource(cfg, client)
			return
		}
		if _bcmdashboardsUpdateDashboard {
			bcmdashboards_UpdateDashboard(cfg, client)
			return
		}

	},
}

var (
	_bcmdashboardsCreateDashboard     bool
	_bcmdashboardsDeleteDashboard     bool
	_bcmdashboardsGetDashboard        bool
	_bcmdashboardsGetResourcePolicy   bool
	_bcmdashboardsListDashboards      bool
	_bcmdashboardsListTagsForResource bool
	_bcmdashboardsTagResource         bool
	_bcmdashboardsUntagResource       bool
	_bcmdashboardsUpdateDashboard     bool

	_bcmdashboardsArn             string
	_bcmdashboardsDescription     string
	_bcmdashboardsMaxResults      string
	_bcmdashboardsName            string
	_bcmdashboardsNextToken       string
	_bcmdashboardsResourceArn     string
	_bcmdashboardsResourceTagKeys []string
	_bcmdashboardsResourceTags    string
	_bcmdashboardsWidgets         string
)

// Creates a new dashboard that can contain multiple widgets displaying cost and
// usage data. You can add custom widgets or use predefined widgets, arranging them
// in your preferred layout.
func bcmdashboards_CreateDashboard(cfg aws.Config, client *bcmdashboards.Client) {
	input := &bcmdashboards.CreateDashboardInput{
		// Name: *string, // Required
		// Widgets: []types.Widget, // Required
	}

	if len(_bcmdashboardsName) > 0 {
		input.Name = aws.String(_bcmdashboardsName)
	}
	if len(_bcmdashboardsWidgets) > 0 {
		if err := assignInputField(input, "Widgets", _bcmdashboardsWidgets); err != nil {
			log.Errorf("invalid --widgets: %s", err.Error())
			return
		}
	}
	if len(_bcmdashboardsDescription) > 0 {
		input.Description = aws.String(_bcmdashboardsDescription)
	}
	if len(_bcmdashboardsResourceTags) > 0 {
		if err := assignInputField(input, "ResourceTags", _bcmdashboardsResourceTags); err != nil {
			log.Errorf("invalid --resource-tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDashboard(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified dashboard. This action cannot be undone.
func bcmdashboards_DeleteDashboard(cfg aws.Config, client *bcmdashboards.Client) {
	input := &bcmdashboards.DeleteDashboardInput{
		// Arn: *string, // Required
	}

	if len(_bcmdashboardsArn) > 0 {
		input.Arn = aws.String(_bcmdashboardsArn)
	}

	if resp, err := client.DeleteDashboard(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the configuration and metadata of a specified dashboard, including
// its widgets and layout settings.
func bcmdashboards_GetDashboard(cfg aws.Config, client *bcmdashboards.Client) {
	input := &bcmdashboards.GetDashboardInput{
		// Arn: *string, // Required
	}

	if len(_bcmdashboardsArn) > 0 {
		input.Arn = aws.String(_bcmdashboardsArn)
	}

	if resp, err := client.GetDashboard(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the resource-based policy attached to a dashboard, showing sharing
// configurations and permissions.
func bcmdashboards_GetResourcePolicy(cfg aws.Config, client *bcmdashboards.Client) {
	input := &bcmdashboards.GetResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_bcmdashboardsResourceArn) > 0 {
		input.ResourceArn = aws.String(_bcmdashboardsResourceArn)
	}

	if resp, err := client.GetResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of all dashboards in your account.
func bcmdashboards_ListDashboards(cfg aws.Config, client *bcmdashboards.Client) {
	input := &bcmdashboards.ListDashboardsInput{}

	if len(_bcmdashboardsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bcmdashboardsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bcmdashboardsNextToken) > 0 {
		input.NextToken = aws.String(_bcmdashboardsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDashboards(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bcmdashboards.ListDashboardsOutput
	p := bcmdashboards.NewListDashboardsPaginator(client, input)
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

// Returns a list of all tags associated with a specified dashboard resource.
func bcmdashboards_ListTagsForResource(cfg aws.Config, client *bcmdashboards.Client) {
	input := &bcmdashboards.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_bcmdashboardsResourceArn) > 0 {
		input.ResourceArn = aws.String(_bcmdashboardsResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates tags for a specified dashboard resource.
func bcmdashboards_TagResource(cfg aws.Config, client *bcmdashboards.Client) {
	input := &bcmdashboards.TagResourceInput{
		// ResourceArn: *string, // Required
		// ResourceTags: []types.ResourceTag, // Required
	}

	if len(_bcmdashboardsResourceArn) > 0 {
		input.ResourceArn = aws.String(_bcmdashboardsResourceArn)
	}
	if len(_bcmdashboardsResourceTags) > 0 {
		if err := assignInputField(input, "ResourceTags", _bcmdashboardsResourceTags); err != nil {
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

// Removes specified tags from a dashboard resource.
func bcmdashboards_UntagResource(cfg aws.Config, client *bcmdashboards.Client) {
	input := &bcmdashboards.UntagResourceInput{
		// ResourceArn: *string, // Required
		// ResourceTagKeys: []string, // Required
	}

	if len(_bcmdashboardsResourceArn) > 0 {
		input.ResourceArn = aws.String(_bcmdashboardsResourceArn)
	}
	if len(_bcmdashboardsResourceTagKeys) > 0 {
		input.ResourceTagKeys = append([]string(nil), _bcmdashboardsResourceTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing dashboard's properties, including its name, description,
// and widget configurations.
func bcmdashboards_UpdateDashboard(cfg aws.Config, client *bcmdashboards.Client) {
	input := &bcmdashboards.UpdateDashboardInput{
		// Arn: *string, // Required
	}

	if len(_bcmdashboardsArn) > 0 {
		input.Arn = aws.String(_bcmdashboardsArn)
	}
	if len(_bcmdashboardsDescription) > 0 {
		input.Description = aws.String(_bcmdashboardsDescription)
	}
	if len(_bcmdashboardsName) > 0 {
		input.Name = aws.String(_bcmdashboardsName)
	}
	if len(_bcmdashboardsWidgets) > 0 {
		if err := assignInputField(input, "Widgets", _bcmdashboardsWidgets); err != nil {
			log.Errorf("invalid --widgets: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDashboard(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_bcmdashboardsCmd)
	_bcmdashboardsCmd.Flags().SortFlags = false

	_bcmdashboardsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_bcmdashboardsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_bcmdashboardsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_bcmdashboardsCmd.Flags().StringVarP(&_bcmdashboardsArn, "arn", "", "", "ARN")
	_bcmdashboardsCmd.Flags().StringVarP(&_bcmdashboardsDescription, "description", "", "", "Description")
	_bcmdashboardsCmd.Flags().StringVarP(&_bcmdashboardsMaxResults, "max-results", "", "", "Max Results")
	_bcmdashboardsCmd.Flags().StringVarP(&_bcmdashboardsName, "name", "", "", "Name")
	_bcmdashboardsCmd.Flags().StringVarP(&_bcmdashboardsNextToken, "next-token", "", "", "Next Token")
	_bcmdashboardsCmd.Flags().StringVarP(&_bcmdashboardsResourceArn, "resource-arn", "", "", "Resource ARN")
	_bcmdashboardsCmd.Flags().StringSliceVarP(&_bcmdashboardsResourceTagKeys, "resource-tag-keys", "", nil, "Resource Tag Keys")
	_bcmdashboardsCmd.Flags().StringVarP(&_bcmdashboardsResourceTags, "resource-tags", "", "", "Resource Tags")
	_bcmdashboardsCmd.Flags().StringVarP(&_bcmdashboardsWidgets, "widgets", "", "", "Widgets")

	_bcmdashboardsCmd.Flags().BoolVarP(&_bcmdashboardsCreateDashboard, "create-dashboard", "", false, "Create Dashboard")
	_bcmdashboardsCmd.Flags().BoolVarP(&_bcmdashboardsDeleteDashboard, "delete-dashboard", "", false, "Delete Dashboard")
	_bcmdashboardsCmd.Flags().BoolVarP(&_bcmdashboardsGetDashboard, "get-dashboard", "", false, "Get Dashboard")
	_bcmdashboardsCmd.Flags().BoolVarP(&_bcmdashboardsGetResourcePolicy, "get-resource-policy", "", false, "Get Resource Policy")
	_bcmdashboardsCmd.Flags().BoolVarP(&_bcmdashboardsListDashboards, "list-dashboards", "", false, "List Dashboards")
	_bcmdashboardsCmd.Flags().BoolVarP(&_bcmdashboardsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_bcmdashboardsCmd.Flags().BoolVarP(&_bcmdashboardsTagResource, "tag-resource", "", false, "Tag Resource")
	_bcmdashboardsCmd.Flags().BoolVarP(&_bcmdashboardsUntagResource, "untag-resource", "", false, "Untag Resource")
	_bcmdashboardsCmd.Flags().BoolVarP(&_bcmdashboardsUpdateDashboard, "update-dashboard", "", false, "Update Dashboard")

}
