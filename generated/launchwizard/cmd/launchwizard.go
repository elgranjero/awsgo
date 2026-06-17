package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/launchwizard"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// launchwizardCmd represents the launchwizard command
var _launchwizardCmd = &cobra.Command{
	Use:   "launchwizard",
	Short: "AWS launchwizard CLI",
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
		client := launchwizard.NewFromConfig(cfg)
		if _launchwizardCreateDeployment {
			launchwizard_CreateDeployment(cfg, client)
			return
		}
		if _launchwizardDeleteDeployment {
			launchwizard_DeleteDeployment(cfg, client)
			return
		}
		if _launchwizardGetDeployment {
			launchwizard_GetDeployment(cfg, client)
			return
		}
		if _launchwizardGetDeploymentPatternVersion {
			launchwizard_GetDeploymentPatternVersion(cfg, client)
			return
		}
		if _launchwizardGetWorkload {
			launchwizard_GetWorkload(cfg, client)
			return
		}
		if _launchwizardGetWorkloadDeploymentPattern {
			launchwizard_GetWorkloadDeploymentPattern(cfg, client)
			return
		}
		if _launchwizardListDeploymentEvents {
			launchwizard_ListDeploymentEvents(cfg, client)
			return
		}
		if _launchwizardListDeploymentPatternVersions {
			launchwizard_ListDeploymentPatternVersions(cfg, client)
			return
		}
		if _launchwizardListDeployments {
			launchwizard_ListDeployments(cfg, client)
			return
		}
		if _launchwizardListTagsForResource {
			launchwizard_ListTagsForResource(cfg, client)
			return
		}
		if _launchwizardListWorkloadDeploymentPatterns {
			launchwizard_ListWorkloadDeploymentPatterns(cfg, client)
			return
		}
		if _launchwizardListWorkloads {
			launchwizard_ListWorkloads(cfg, client)
			return
		}
		if _launchwizardTagResource {
			launchwizard_TagResource(cfg, client)
			return
		}
		if _launchwizardUntagResource {
			launchwizard_UntagResource(cfg, client)
			return
		}
		if _launchwizardUpdateDeployment {
			launchwizard_UpdateDeployment(cfg, client)
			return
		}

	},
}

var (
	_launchwizardCreateDeployment               bool
	_launchwizardDeleteDeployment               bool
	_launchwizardGetDeployment                  bool
	_launchwizardGetDeploymentPatternVersion    bool
	_launchwizardGetWorkload                    bool
	_launchwizardGetWorkloadDeploymentPattern   bool
	_launchwizardListDeploymentEvents           bool
	_launchwizardListDeploymentPatternVersions  bool
	_launchwizardListDeployments                bool
	_launchwizardListTagsForResource            bool
	_launchwizardListWorkloadDeploymentPatterns bool
	_launchwizardListWorkloads                  bool
	_launchwizardTagResource                    bool
	_launchwizardUntagResource                  bool
	_launchwizardUpdateDeployment               bool

	_launchwizardDeploymentId                 string
	_launchwizardDeploymentPatternName        string
	_launchwizardDeploymentPatternVersionName string
	_launchwizardDryRun                       string
	_launchwizardFilters                      string
	_launchwizardForce                        string
	_launchwizardMaxResults                   string
	_launchwizardName                         string
	_launchwizardNextToken                    string
	_launchwizardResourceArn                  string
	_launchwizardSpecifications               string
	_launchwizardTagKeys                      []string
	_launchwizardTags                         string
	_launchwizardWorkloadName                 string
	_launchwizardWorkloadVersionName          string
)

// Creates a deployment for the given workload. Deployments created by this
// operation are not available in the Launch Wizard console to use the Clone
// deployment action on.
func launchwizard_CreateDeployment(cfg aws.Config, client *launchwizard.Client) {
	input := &launchwizard.CreateDeploymentInput{
		// DeploymentPatternName: *string, // Required
		// Name: *string, // Required
		// Specifications: map[string]string, // Required
		// WorkloadName: *string, // Required
	}

	if len(_launchwizardDeploymentPatternName) > 0 {
		input.DeploymentPatternName = aws.String(_launchwizardDeploymentPatternName)
	}
	if len(_launchwizardName) > 0 {
		input.Name = aws.String(_launchwizardName)
	}
	if len(_launchwizardSpecifications) > 0 {
		if err := assignInputField(input, "Specifications", _launchwizardSpecifications); err != nil {
			log.Errorf("invalid --specifications: %s", err.Error())
			return
		}
	}
	if len(_launchwizardWorkloadName) > 0 {
		input.WorkloadName = aws.String(_launchwizardWorkloadName)
	}
	if len(_launchwizardDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _launchwizardDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_launchwizardTags) > 0 {
		if err := assignInputField(input, "Tags", _launchwizardTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a deployment.
func launchwizard_DeleteDeployment(cfg aws.Config, client *launchwizard.Client) {
	input := &launchwizard.DeleteDeploymentInput{
		// DeploymentId: *string, // Required
	}

	if len(_launchwizardDeploymentId) > 0 {
		input.DeploymentId = aws.String(_launchwizardDeploymentId)
	}

	if resp, err := client.DeleteDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the deployment.
func launchwizard_GetDeployment(cfg aws.Config, client *launchwizard.Client) {
	input := &launchwizard.GetDeploymentInput{
		// DeploymentId: *string, // Required
	}

	if len(_launchwizardDeploymentId) > 0 {
		input.DeploymentId = aws.String(_launchwizardDeploymentId)
	}

	if resp, err := client.GetDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a deployment pattern version.
func launchwizard_GetDeploymentPatternVersion(cfg aws.Config, client *launchwizard.Client) {
	input := &launchwizard.GetDeploymentPatternVersionInput{
		// DeploymentPatternName: *string, // Required
		// DeploymentPatternVersionName: *string, // Required
		// WorkloadName: *string, // Required
	}

	if len(_launchwizardDeploymentPatternName) > 0 {
		input.DeploymentPatternName = aws.String(_launchwizardDeploymentPatternName)
	}
	if len(_launchwizardDeploymentPatternVersionName) > 0 {
		input.DeploymentPatternVersionName = aws.String(_launchwizardDeploymentPatternVersionName)
	}
	if len(_launchwizardWorkloadName) > 0 {
		input.WorkloadName = aws.String(_launchwizardWorkloadName)
	}

	if resp, err := client.GetDeploymentPatternVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a workload.
func launchwizard_GetWorkload(cfg aws.Config, client *launchwizard.Client) {
	input := &launchwizard.GetWorkloadInput{
		// WorkloadName: *string, // Required
	}

	if len(_launchwizardWorkloadName) > 0 {
		input.WorkloadName = aws.String(_launchwizardWorkloadName)
	}

	if resp, err := client.GetWorkload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details for a given workload and deployment pattern, including the
// available specifications. You can use the [ListWorkloads]operation to discover the available
// workload names and the [ListWorkloadDeploymentPatterns]operation to discover the available deployment pattern
// names of a given workload.
//
// [ListWorkloadDeploymentPatterns]: https://docs.aws.amazon.com/launchwizard/latest/APIReference/API_ListWorkloadDeploymentPatterns.html
// [ListWorkloads]: https://docs.aws.amazon.com/launchwizard/latest/APIReference/API_ListWorkloads.html
func launchwizard_GetWorkloadDeploymentPattern(cfg aws.Config, client *launchwizard.Client) {
	input := &launchwizard.GetWorkloadDeploymentPatternInput{
		// DeploymentPatternName: *string, // Required
		// WorkloadName: *string, // Required
	}

	if len(_launchwizardDeploymentPatternName) > 0 {
		input.DeploymentPatternName = aws.String(_launchwizardDeploymentPatternName)
	}
	if len(_launchwizardWorkloadName) > 0 {
		input.WorkloadName = aws.String(_launchwizardWorkloadName)
	}

	if resp, err := client.GetWorkloadDeploymentPattern(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the events of a deployment.
func launchwizard_ListDeploymentEvents(cfg aws.Config, client *launchwizard.Client) {
	input := &launchwizard.ListDeploymentEventsInput{
		// DeploymentId: *string, // Required
	}

	if len(_launchwizardDeploymentId) > 0 {
		input.DeploymentId = aws.String(_launchwizardDeploymentId)
	}
	if len(_launchwizardMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _launchwizardMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_launchwizardNextToken) > 0 {
		input.NextToken = aws.String(_launchwizardNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDeploymentEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*launchwizard.ListDeploymentEventsOutput
	p := launchwizard.NewListDeploymentEventsPaginator(client, input)
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

// Lists the deployment pattern versions.
func launchwizard_ListDeploymentPatternVersions(cfg aws.Config, client *launchwizard.Client) {
	input := &launchwizard.ListDeploymentPatternVersionsInput{
		// DeploymentPatternName: *string, // Required
		// WorkloadName: *string, // Required
	}

	if len(_launchwizardDeploymentPatternName) > 0 {
		input.DeploymentPatternName = aws.String(_launchwizardDeploymentPatternName)
	}
	if len(_launchwizardWorkloadName) > 0 {
		input.WorkloadName = aws.String(_launchwizardWorkloadName)
	}
	if len(_launchwizardFilters) > 0 {
		if err := assignInputField(input, "Filters", _launchwizardFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_launchwizardMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _launchwizardMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_launchwizardNextToken) > 0 {
		input.NextToken = aws.String(_launchwizardNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDeploymentPatternVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*launchwizard.ListDeploymentPatternVersionsOutput
	p := launchwizard.NewListDeploymentPatternVersionsPaginator(client, input)
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

// Lists the deployments that have been created.
func launchwizard_ListDeployments(cfg aws.Config, client *launchwizard.Client) {
	input := &launchwizard.ListDeploymentsInput{}

	if len(_launchwizardFilters) > 0 {
		if err := assignInputField(input, "Filters", _launchwizardFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_launchwizardMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _launchwizardMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_launchwizardNextToken) > 0 {
		input.NextToken = aws.String(_launchwizardNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDeployments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*launchwizard.ListDeploymentsOutput
	p := launchwizard.NewListDeploymentsPaginator(client, input)
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

// Lists the tags associated with a specified resource.
func launchwizard_ListTagsForResource(cfg aws.Config, client *launchwizard.Client) {
	input := &launchwizard.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_launchwizardResourceArn) > 0 {
		input.ResourceArn = aws.String(_launchwizardResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the workload deployment patterns for a given workload name. You can use
// the [ListWorkloads]operation to discover the available workload names.
//
// [ListWorkloads]: https://docs.aws.amazon.com/launchwizard/latest/APIReference/API_ListWorkloads.html
func launchwizard_ListWorkloadDeploymentPatterns(cfg aws.Config, client *launchwizard.Client) {
	input := &launchwizard.ListWorkloadDeploymentPatternsInput{
		// WorkloadName: *string, // Required
	}

	if len(_launchwizardWorkloadName) > 0 {
		input.WorkloadName = aws.String(_launchwizardWorkloadName)
	}
	if len(_launchwizardMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _launchwizardMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_launchwizardNextToken) > 0 {
		input.NextToken = aws.String(_launchwizardNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkloadDeploymentPatterns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*launchwizard.ListWorkloadDeploymentPatternsOutput
	p := launchwizard.NewListWorkloadDeploymentPatternsPaginator(client, input)
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

// Lists the available workload names. You can use the [ListWorkloadDeploymentPatterns] operation to discover the
// available deployment patterns for a given workload.
//
// [ListWorkloadDeploymentPatterns]: https://docs.aws.amazon.com/launchwizard/latest/APIReference/API_ListWorkloadDeploymentPatterns.html
func launchwizard_ListWorkloads(cfg aws.Config, client *launchwizard.Client) {
	input := &launchwizard.ListWorkloadsInput{}

	if len(_launchwizardMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _launchwizardMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_launchwizardNextToken) > 0 {
		input.NextToken = aws.String(_launchwizardNextToken)
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

	var results []*launchwizard.ListWorkloadsOutput
	p := launchwizard.NewListWorkloadsPaginator(client, input)
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

// Adds the specified tags to the given resource.
func launchwizard_TagResource(cfg aws.Config, client *launchwizard.Client) {
	input := &launchwizard.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_launchwizardResourceArn) > 0 {
		input.ResourceArn = aws.String(_launchwizardResourceArn)
	}
	if len(_launchwizardTags) > 0 {
		if err := assignInputField(input, "Tags", _launchwizardTags); err != nil {
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

// Removes the specified tags from the given resource.
func launchwizard_UntagResource(cfg aws.Config, client *launchwizard.Client) {
	input := &launchwizard.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_launchwizardResourceArn) > 0 {
		input.ResourceArn = aws.String(_launchwizardResourceArn)
	}
	if len(_launchwizardTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _launchwizardTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a deployment.
func launchwizard_UpdateDeployment(cfg aws.Config, client *launchwizard.Client) {
	input := &launchwizard.UpdateDeploymentInput{
		// DeploymentId: *string, // Required
		// Specifications: map[string]string, // Required
	}

	if len(_launchwizardDeploymentId) > 0 {
		input.DeploymentId = aws.String(_launchwizardDeploymentId)
	}
	if len(_launchwizardSpecifications) > 0 {
		if err := assignInputField(input, "Specifications", _launchwizardSpecifications); err != nil {
			log.Errorf("invalid --specifications: %s", err.Error())
			return
		}
	}
	if len(_launchwizardDeploymentPatternVersionName) > 0 {
		input.DeploymentPatternVersionName = aws.String(_launchwizardDeploymentPatternVersionName)
	}
	if len(_launchwizardDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _launchwizardDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_launchwizardForce) > 0 {
		if err := assignInputField(input, "Force", _launchwizardForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}
	if len(_launchwizardWorkloadVersionName) > 0 {
		input.WorkloadVersionName = aws.String(_launchwizardWorkloadVersionName)
	}

	if resp, err := client.UpdateDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_launchwizardCmd)
	_launchwizardCmd.Flags().SortFlags = false

	_launchwizardCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_launchwizardCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_launchwizardCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_launchwizardCmd.Flags().StringVarP(&_launchwizardDeploymentId, "deployment-id", "", "", "Deployment ID")
	_launchwizardCmd.Flags().StringVarP(&_launchwizardDeploymentPatternName, "deployment-pattern-name", "", "", "Deployment Pattern Name")
	_launchwizardCmd.Flags().StringVarP(&_launchwizardDeploymentPatternVersionName, "deployment-pattern-version-name", "", "", "Deployment Pattern Version Name")
	_launchwizardCmd.Flags().StringVarP(&_launchwizardDryRun, "dry-run", "", "", "Dry Run")
	_launchwizardCmd.Flags().StringVarP(&_launchwizardFilters, "filters", "", "", "Filters")
	_launchwizardCmd.Flags().StringVarP(&_launchwizardForce, "force", "", "", "Force")
	_launchwizardCmd.Flags().StringVarP(&_launchwizardMaxResults, "max-results", "", "", "Max Results")
	_launchwizardCmd.Flags().StringVarP(&_launchwizardName, "name", "", "", "Name")
	_launchwizardCmd.Flags().StringVarP(&_launchwizardNextToken, "next-token", "", "", "Next Token")
	_launchwizardCmd.Flags().StringVarP(&_launchwizardResourceArn, "resource-arn", "", "", "Resource ARN")
	_launchwizardCmd.Flags().StringVarP(&_launchwizardSpecifications, "specifications", "", "", "Specifications")
	_launchwizardCmd.Flags().StringSliceVarP(&_launchwizardTagKeys, "tag-keys", "", nil, "Tag Keys")
	_launchwizardCmd.Flags().StringVarP(&_launchwizardTags, "tags", "", "", "Tags")
	_launchwizardCmd.Flags().StringVarP(&_launchwizardWorkloadName, "workload-name", "", "", "Workload Name")
	_launchwizardCmd.Flags().StringVarP(&_launchwizardWorkloadVersionName, "workload-version-name", "", "", "Workload Version Name")

	_launchwizardCmd.Flags().BoolVarP(&_launchwizardCreateDeployment, "create-deployment", "", false, "Create Deployment")
	_launchwizardCmd.Flags().BoolVarP(&_launchwizardDeleteDeployment, "delete-deployment", "", false, "Delete Deployment")
	_launchwizardCmd.Flags().BoolVarP(&_launchwizardGetDeployment, "get-deployment", "", false, "Get Deployment")
	_launchwizardCmd.Flags().BoolVarP(&_launchwizardGetDeploymentPatternVersion, "get-deployment-pattern-version", "", false, "Get Deployment Pattern Version")
	_launchwizardCmd.Flags().BoolVarP(&_launchwizardGetWorkload, "get-workload", "", false, "Get Workload")
	_launchwizardCmd.Flags().BoolVarP(&_launchwizardGetWorkloadDeploymentPattern, "get-workload-deployment-pattern", "", false, "Get Workload Deployment Pattern")
	_launchwizardCmd.Flags().BoolVarP(&_launchwizardListDeploymentEvents, "list-deployment-events", "", false, "List Deployment Events")
	_launchwizardCmd.Flags().BoolVarP(&_launchwizardListDeploymentPatternVersions, "list-deployment-pattern-versions", "", false, "List Deployment Pattern Versions")
	_launchwizardCmd.Flags().BoolVarP(&_launchwizardListDeployments, "list-deployments", "", false, "List Deployments")
	_launchwizardCmd.Flags().BoolVarP(&_launchwizardListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_launchwizardCmd.Flags().BoolVarP(&_launchwizardListWorkloadDeploymentPatterns, "list-workload-deployment-patterns", "", false, "List Workload Deployment Patterns")
	_launchwizardCmd.Flags().BoolVarP(&_launchwizardListWorkloads, "list-workloads", "", false, "List Workloads")
	_launchwizardCmd.Flags().BoolVarP(&_launchwizardTagResource, "tag-resource", "", false, "Tag Resource")
	_launchwizardCmd.Flags().BoolVarP(&_launchwizardUntagResource, "untag-resource", "", false, "Untag Resource")
	_launchwizardCmd.Flags().BoolVarP(&_launchwizardUpdateDeployment, "update-deployment", "", false, "Update Deployment")

}
