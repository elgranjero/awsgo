package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// route53recoveryreadinessCmd represents the route53recoveryreadiness command
var _route53recoveryreadinessCmd = &cobra.Command{
	Use:   "route53recoveryreadiness",
	Short: "AWS route53recoveryreadiness CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := route53recoveryreadiness.NewFromConfig(cfg)
		if _route53recoveryreadinessCreateCell {
			route53recoveryreadiness_CreateCell(cfg, client)
			return
		}
		if _route53recoveryreadinessCreateCrossAccountAuthorization {
			route53recoveryreadiness_CreateCrossAccountAuthorization(cfg, client)
			return
		}
		if _route53recoveryreadinessCreateReadinessCheck {
			route53recoveryreadiness_CreateReadinessCheck(cfg, client)
			return
		}
		if _route53recoveryreadinessCreateRecoveryGroup {
			route53recoveryreadiness_CreateRecoveryGroup(cfg, client)
			return
		}
		if _route53recoveryreadinessCreateResourceSet {
			route53recoveryreadiness_CreateResourceSet(cfg, client)
			return
		}
		if _route53recoveryreadinessDeleteCell {
			route53recoveryreadiness_DeleteCell(cfg, client)
			return
		}
		if _route53recoveryreadinessDeleteCrossAccountAuthorization {
			route53recoveryreadiness_DeleteCrossAccountAuthorization(cfg, client)
			return
		}
		if _route53recoveryreadinessDeleteReadinessCheck {
			route53recoveryreadiness_DeleteReadinessCheck(cfg, client)
			return
		}
		if _route53recoveryreadinessDeleteRecoveryGroup {
			route53recoveryreadiness_DeleteRecoveryGroup(cfg, client)
			return
		}
		if _route53recoveryreadinessDeleteResourceSet {
			route53recoveryreadiness_DeleteResourceSet(cfg, client)
			return
		}
		if _route53recoveryreadinessGetArchitectureRecommendations {
			route53recoveryreadiness_GetArchitectureRecommendations(cfg, client)
			return
		}
		if _route53recoveryreadinessGetCell {
			route53recoveryreadiness_GetCell(cfg, client)
			return
		}
		if _route53recoveryreadinessGetCellReadinessSummary {
			route53recoveryreadiness_GetCellReadinessSummary(cfg, client)
			return
		}
		if _route53recoveryreadinessGetReadinessCheck {
			route53recoveryreadiness_GetReadinessCheck(cfg, client)
			return
		}
		if _route53recoveryreadinessGetReadinessCheckResourceStatus {
			route53recoveryreadiness_GetReadinessCheckResourceStatus(cfg, client)
			return
		}
		if _route53recoveryreadinessGetReadinessCheckStatus {
			route53recoveryreadiness_GetReadinessCheckStatus(cfg, client)
			return
		}
		if _route53recoveryreadinessGetRecoveryGroup {
			route53recoveryreadiness_GetRecoveryGroup(cfg, client)
			return
		}
		if _route53recoveryreadinessGetRecoveryGroupReadinessSummary {
			route53recoveryreadiness_GetRecoveryGroupReadinessSummary(cfg, client)
			return
		}
		if _route53recoveryreadinessGetResourceSet {
			route53recoveryreadiness_GetResourceSet(cfg, client)
			return
		}
		if _route53recoveryreadinessListCells {
			route53recoveryreadiness_ListCells(cfg, client)
			return
		}
		if _route53recoveryreadinessListCrossAccountAuthorizations {
			route53recoveryreadiness_ListCrossAccountAuthorizations(cfg, client)
			return
		}
		if _route53recoveryreadinessListReadinessChecks {
			route53recoveryreadiness_ListReadinessChecks(cfg, client)
			return
		}
		if _route53recoveryreadinessListRecoveryGroups {
			route53recoveryreadiness_ListRecoveryGroups(cfg, client)
			return
		}
		if _route53recoveryreadinessListResourceSets {
			route53recoveryreadiness_ListResourceSets(cfg, client)
			return
		}
		if _route53recoveryreadinessListRules {
			route53recoveryreadiness_ListRules(cfg, client)
			return
		}
		if _route53recoveryreadinessListTagsForResources {
			route53recoveryreadiness_ListTagsForResources(cfg, client)
			return
		}
		if _route53recoveryreadinessTagResource {
			route53recoveryreadiness_TagResource(cfg, client)
			return
		}
		if _route53recoveryreadinessUntagResource {
			route53recoveryreadiness_UntagResource(cfg, client)
			return
		}
		if _route53recoveryreadinessUpdateCell {
			route53recoveryreadiness_UpdateCell(cfg, client)
			return
		}
		if _route53recoveryreadinessUpdateReadinessCheck {
			route53recoveryreadiness_UpdateReadinessCheck(cfg, client)
			return
		}
		if _route53recoveryreadinessUpdateRecoveryGroup {
			route53recoveryreadiness_UpdateRecoveryGroup(cfg, client)
			return
		}
		if _route53recoveryreadinessUpdateResourceSet {
			route53recoveryreadiness_UpdateResourceSet(cfg, client)
			return
		}

	},
}

var (
	_route53recoveryreadinessCreateCell                       bool
	_route53recoveryreadinessCreateCrossAccountAuthorization  bool
	_route53recoveryreadinessCreateReadinessCheck             bool
	_route53recoveryreadinessCreateRecoveryGroup              bool
	_route53recoveryreadinessCreateResourceSet                bool
	_route53recoveryreadinessDeleteCell                       bool
	_route53recoveryreadinessDeleteCrossAccountAuthorization  bool
	_route53recoveryreadinessDeleteReadinessCheck             bool
	_route53recoveryreadinessDeleteRecoveryGroup              bool
	_route53recoveryreadinessDeleteResourceSet                bool
	_route53recoveryreadinessGetArchitectureRecommendations   bool
	_route53recoveryreadinessGetCell                          bool
	_route53recoveryreadinessGetCellReadinessSummary          bool
	_route53recoveryreadinessGetReadinessCheck                bool
	_route53recoveryreadinessGetReadinessCheckResourceStatus  bool
	_route53recoveryreadinessGetReadinessCheckStatus          bool
	_route53recoveryreadinessGetRecoveryGroup                 bool
	_route53recoveryreadinessGetRecoveryGroupReadinessSummary bool
	_route53recoveryreadinessGetResourceSet                   bool
	_route53recoveryreadinessListCells                        bool
	_route53recoveryreadinessListCrossAccountAuthorizations   bool
	_route53recoveryreadinessListReadinessChecks              bool
	_route53recoveryreadinessListRecoveryGroups               bool
	_route53recoveryreadinessListResourceSets                 bool
	_route53recoveryreadinessListRules                        bool
	_route53recoveryreadinessListTagsForResources             bool
	_route53recoveryreadinessTagResource                      bool
	_route53recoveryreadinessUntagResource                    bool
	_route53recoveryreadinessUpdateCell                       bool
	_route53recoveryreadinessUpdateReadinessCheck             bool
	_route53recoveryreadinessUpdateRecoveryGroup              bool
	_route53recoveryreadinessUpdateResourceSet                bool

	_route53recoveryreadinessCellName                  string
	_route53recoveryreadinessCells                     []string
	_route53recoveryreadinessCrossAccountAuthorization string
	_route53recoveryreadinessMaxResults                string
	_route53recoveryreadinessNextToken                 string
	_route53recoveryreadinessReadinessCheckName        string
	_route53recoveryreadinessRecoveryGroupName         string
	_route53recoveryreadinessResourceArn               string
	_route53recoveryreadinessResourceIdentifier        string
	_route53recoveryreadinessResourceSetName           string
	_route53recoveryreadinessResourceSetType           string
	_route53recoveryreadinessResourceType              string
	_route53recoveryreadinessResources                 string
	_route53recoveryreadinessTagKeys                   []string
	_route53recoveryreadinessTags                      string
)

// Creates a cell in an account.
func route53recoveryreadiness_CreateCell(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.CreateCellInput{
		// CellName: *string, // Required
	}

	if len(_route53recoveryreadinessCellName) > 0 {
		input.CellName = aws.String(_route53recoveryreadinessCellName)
	}
	if len(_route53recoveryreadinessCells) > 0 {
		input.Cells = append([]string(nil), _route53recoveryreadinessCells...)
	}
	if len(_route53recoveryreadinessTags) > 0 {
		if err := assignInputField(input, "Tags", _route53recoveryreadinessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCell(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a cross-account readiness authorization. This lets you authorize
// another account to work with Route 53 Application Recovery Controller, for
// example, to check the readiness status of resources in a separate account.
func route53recoveryreadiness_CreateCrossAccountAuthorization(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.CreateCrossAccountAuthorizationInput{
		// CrossAccountAuthorization: *string, // Required
	}

	if len(_route53recoveryreadinessCrossAccountAuthorization) > 0 {
		input.CrossAccountAuthorization = aws.String(_route53recoveryreadinessCrossAccountAuthorization)
	}

	if resp, err := client.CreateCrossAccountAuthorization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a readiness check in an account. A readiness check monitors a resource
// set in your application, such as a set of Amazon Aurora instances, that
// Application Recovery Controller is auditing recovery readiness for. The audits
// run once every minute on every resource that's associated with a readiness
// check.
func route53recoveryreadiness_CreateReadinessCheck(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.CreateReadinessCheckInput{
		// ReadinessCheckName: *string, // Required
		// ResourceSetName: *string, // Required
	}

	if len(_route53recoveryreadinessReadinessCheckName) > 0 {
		input.ReadinessCheckName = aws.String(_route53recoveryreadinessReadinessCheckName)
	}
	if len(_route53recoveryreadinessResourceSetName) > 0 {
		input.ResourceSetName = aws.String(_route53recoveryreadinessResourceSetName)
	}
	if len(_route53recoveryreadinessTags) > 0 {
		if err := assignInputField(input, "Tags", _route53recoveryreadinessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateReadinessCheck(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a recovery group in an account. A recovery group corresponds to an
// application and includes a list of the cells that make up the application.
func route53recoveryreadiness_CreateRecoveryGroup(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.CreateRecoveryGroupInput{
		// RecoveryGroupName: *string, // Required
	}

	if len(_route53recoveryreadinessRecoveryGroupName) > 0 {
		input.RecoveryGroupName = aws.String(_route53recoveryreadinessRecoveryGroupName)
	}
	if len(_route53recoveryreadinessCells) > 0 {
		input.Cells = append([]string(nil), _route53recoveryreadinessCells...)
	}
	if len(_route53recoveryreadinessTags) > 0 {
		if err := assignInputField(input, "Tags", _route53recoveryreadinessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRecoveryGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a resource set. A resource set is a set of resources of one type that
// span multiple cells. You can associate a resource set with a readiness check to
// monitor the resources for failover readiness.
func route53recoveryreadiness_CreateResourceSet(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.CreateResourceSetInput{
		// ResourceSetName: *string, // Required
		// ResourceSetType: *string, // Required
		// Resources: []types.Resource, // Required
	}

	if len(_route53recoveryreadinessResourceSetName) > 0 {
		input.ResourceSetName = aws.String(_route53recoveryreadinessResourceSetName)
	}
	if len(_route53recoveryreadinessResourceSetType) > 0 {
		input.ResourceSetType = aws.String(_route53recoveryreadinessResourceSetType)
	}
	if len(_route53recoveryreadinessResources) > 0 {
		if err := assignInputField(input, "Resources", _route53recoveryreadinessResources); err != nil {
			log.Errorf("invalid --resources: %s", err.Error())
			return
		}
	}
	if len(_route53recoveryreadinessTags) > 0 {
		if err := assignInputField(input, "Tags", _route53recoveryreadinessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateResourceSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a cell. When successful, the response code is 204, with no response body.
func route53recoveryreadiness_DeleteCell(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.DeleteCellInput{
		// CellName: *string, // Required
	}

	if len(_route53recoveryreadinessCellName) > 0 {
		input.CellName = aws.String(_route53recoveryreadinessCellName)
	}

	if resp, err := client.DeleteCell(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes cross account readiness authorization.
func route53recoveryreadiness_DeleteCrossAccountAuthorization(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.DeleteCrossAccountAuthorizationInput{
		// CrossAccountAuthorization: *string, // Required
	}

	if len(_route53recoveryreadinessCrossAccountAuthorization) > 0 {
		input.CrossAccountAuthorization = aws.String(_route53recoveryreadinessCrossAccountAuthorization)
	}

	if resp, err := client.DeleteCrossAccountAuthorization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a readiness check.
func route53recoveryreadiness_DeleteReadinessCheck(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.DeleteReadinessCheckInput{
		// ReadinessCheckName: *string, // Required
	}

	if len(_route53recoveryreadinessReadinessCheckName) > 0 {
		input.ReadinessCheckName = aws.String(_route53recoveryreadinessReadinessCheckName)
	}

	if resp, err := client.DeleteReadinessCheck(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a recovery group.
func route53recoveryreadiness_DeleteRecoveryGroup(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.DeleteRecoveryGroupInput{
		// RecoveryGroupName: *string, // Required
	}

	if len(_route53recoveryreadinessRecoveryGroupName) > 0 {
		input.RecoveryGroupName = aws.String(_route53recoveryreadinessRecoveryGroupName)
	}

	if resp, err := client.DeleteRecoveryGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a resource set.
func route53recoveryreadiness_DeleteResourceSet(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.DeleteResourceSetInput{
		// ResourceSetName: *string, // Required
	}

	if len(_route53recoveryreadinessResourceSetName) > 0 {
		input.ResourceSetName = aws.String(_route53recoveryreadinessResourceSetName)
	}

	if resp, err := client.DeleteResourceSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets recommendations about architecture designs for improving resiliency for an
// application, based on a recovery group.
func route53recoveryreadiness_GetArchitectureRecommendations(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.GetArchitectureRecommendationsInput{
		// RecoveryGroupName: *string, // Required
	}

	if len(_route53recoveryreadinessRecoveryGroupName) > 0 {
		input.RecoveryGroupName = aws.String(_route53recoveryreadinessRecoveryGroupName)
	}
	if len(_route53recoveryreadinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53recoveryreadinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53recoveryreadinessNextToken) > 0 {
		input.NextToken = aws.String(_route53recoveryreadinessNextToken)
	}

	if resp, err := client.GetArchitectureRecommendations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a cell including cell name, cell Amazon Resource Name
// (ARN), ARNs of nested cells for this cell, and a list of those cell ARNs with
// their associated recovery group ARNs.
func route53recoveryreadiness_GetCell(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.GetCellInput{
		// CellName: *string, // Required
	}

	if len(_route53recoveryreadinessCellName) > 0 {
		input.CellName = aws.String(_route53recoveryreadinessCellName)
	}

	if resp, err := client.GetCell(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets readiness for a cell. Aggregates the readiness of all the resources that
// are associated with the cell into a single value.
func route53recoveryreadiness_GetCellReadinessSummary(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.GetCellReadinessSummaryInput{
		// CellName: *string, // Required
	}

	if len(_route53recoveryreadinessCellName) > 0 {
		input.CellName = aws.String(_route53recoveryreadinessCellName)
	}
	if len(_route53recoveryreadinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53recoveryreadinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53recoveryreadinessNextToken) > 0 {
		input.NextToken = aws.String(_route53recoveryreadinessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetCellReadinessSummary(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53recoveryreadiness.GetCellReadinessSummaryOutput
	p := route53recoveryreadiness.NewGetCellReadinessSummaryPaginator(client, input)
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

// Gets details about a readiness check.
func route53recoveryreadiness_GetReadinessCheck(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.GetReadinessCheckInput{
		// ReadinessCheckName: *string, // Required
	}

	if len(_route53recoveryreadinessReadinessCheckName) > 0 {
		input.ReadinessCheckName = aws.String(_route53recoveryreadinessReadinessCheckName)
	}

	if resp, err := client.GetReadinessCheck(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets individual readiness status for a readiness check. To see the overall
// readiness status for a recovery group, that considers the readiness status for
// all the readiness checks in the recovery group, use
// GetRecoveryGroupReadinessSummary.
func route53recoveryreadiness_GetReadinessCheckResourceStatus(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.GetReadinessCheckResourceStatusInput{
		// ReadinessCheckName: *string, // Required
		// ResourceIdentifier: *string, // Required
	}

	if len(_route53recoveryreadinessReadinessCheckName) > 0 {
		input.ReadinessCheckName = aws.String(_route53recoveryreadinessReadinessCheckName)
	}
	if len(_route53recoveryreadinessResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_route53recoveryreadinessResourceIdentifier)
	}
	if len(_route53recoveryreadinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53recoveryreadinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53recoveryreadinessNextToken) > 0 {
		input.NextToken = aws.String(_route53recoveryreadinessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetReadinessCheckResourceStatus(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53recoveryreadiness.GetReadinessCheckResourceStatusOutput
	p := route53recoveryreadiness.NewGetReadinessCheckResourceStatusPaginator(client, input)
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

// Gets the readiness status for an individual readiness check. To see the overall
// readiness status for a recovery group, that considers the readiness status for
// all the readiness checks in a recovery group, use
// GetRecoveryGroupReadinessSummary.
func route53recoveryreadiness_GetReadinessCheckStatus(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.GetReadinessCheckStatusInput{
		// ReadinessCheckName: *string, // Required
	}

	if len(_route53recoveryreadinessReadinessCheckName) > 0 {
		input.ReadinessCheckName = aws.String(_route53recoveryreadinessReadinessCheckName)
	}
	if len(_route53recoveryreadinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53recoveryreadinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53recoveryreadinessNextToken) > 0 {
		input.NextToken = aws.String(_route53recoveryreadinessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetReadinessCheckStatus(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53recoveryreadiness.GetReadinessCheckStatusOutput
	p := route53recoveryreadiness.NewGetReadinessCheckStatusPaginator(client, input)
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

// Gets details about a recovery group, including a list of the cells that are
// included in it.
func route53recoveryreadiness_GetRecoveryGroup(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.GetRecoveryGroupInput{
		// RecoveryGroupName: *string, // Required
	}

	if len(_route53recoveryreadinessRecoveryGroupName) > 0 {
		input.RecoveryGroupName = aws.String(_route53recoveryreadinessRecoveryGroupName)
	}

	if resp, err := client.GetRecoveryGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays a summary of information about a recovery group's readiness status.
// Includes the readiness checks for resources in the recovery group and the
// readiness status of each one.
func route53recoveryreadiness_GetRecoveryGroupReadinessSummary(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.GetRecoveryGroupReadinessSummaryInput{
		// RecoveryGroupName: *string, // Required
	}

	if len(_route53recoveryreadinessRecoveryGroupName) > 0 {
		input.RecoveryGroupName = aws.String(_route53recoveryreadinessRecoveryGroupName)
	}
	if len(_route53recoveryreadinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53recoveryreadinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53recoveryreadinessNextToken) > 0 {
		input.NextToken = aws.String(_route53recoveryreadinessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetRecoveryGroupReadinessSummary(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53recoveryreadiness.GetRecoveryGroupReadinessSummaryOutput
	p := route53recoveryreadiness.NewGetRecoveryGroupReadinessSummaryPaginator(client, input)
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

// Displays the details about a resource set, including a list of the resources in
// the set.
func route53recoveryreadiness_GetResourceSet(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.GetResourceSetInput{
		// ResourceSetName: *string, // Required
	}

	if len(_route53recoveryreadinessResourceSetName) > 0 {
		input.ResourceSetName = aws.String(_route53recoveryreadinessResourceSetName)
	}

	if resp, err := client.GetResourceSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the cells for an account.
func route53recoveryreadiness_ListCells(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.ListCellsInput{}

	if len(_route53recoveryreadinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53recoveryreadinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53recoveryreadinessNextToken) > 0 {
		input.NextToken = aws.String(_route53recoveryreadinessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCells(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53recoveryreadiness.ListCellsOutput
	p := route53recoveryreadiness.NewListCellsPaginator(client, input)
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

// Lists the cross-account readiness authorizations that are in place for an
// account.
func route53recoveryreadiness_ListCrossAccountAuthorizations(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.ListCrossAccountAuthorizationsInput{}

	if len(_route53recoveryreadinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53recoveryreadinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53recoveryreadinessNextToken) > 0 {
		input.NextToken = aws.String(_route53recoveryreadinessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCrossAccountAuthorizations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53recoveryreadiness.ListCrossAccountAuthorizationsOutput
	p := route53recoveryreadiness.NewListCrossAccountAuthorizationsPaginator(client, input)
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

// Lists the readiness checks for an account.
func route53recoveryreadiness_ListReadinessChecks(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.ListReadinessChecksInput{}

	if len(_route53recoveryreadinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53recoveryreadinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53recoveryreadinessNextToken) > 0 {
		input.NextToken = aws.String(_route53recoveryreadinessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListReadinessChecks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53recoveryreadiness.ListReadinessChecksOutput
	p := route53recoveryreadiness.NewListReadinessChecksPaginator(client, input)
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

// Lists the recovery groups in an account.
func route53recoveryreadiness_ListRecoveryGroups(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.ListRecoveryGroupsInput{}

	if len(_route53recoveryreadinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53recoveryreadinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53recoveryreadinessNextToken) > 0 {
		input.NextToken = aws.String(_route53recoveryreadinessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRecoveryGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53recoveryreadiness.ListRecoveryGroupsOutput
	p := route53recoveryreadiness.NewListRecoveryGroupsPaginator(client, input)
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

// Lists the resource sets in an account.
func route53recoveryreadiness_ListResourceSets(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.ListResourceSetsInput{}

	if len(_route53recoveryreadinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53recoveryreadinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53recoveryreadinessNextToken) > 0 {
		input.NextToken = aws.String(_route53recoveryreadinessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListResourceSets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53recoveryreadiness.ListResourceSetsOutput
	p := route53recoveryreadiness.NewListResourceSetsPaginator(client, input)
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

// Lists all readiness rules, or lists the readiness rules for a specific resource
// type.
func route53recoveryreadiness_ListRules(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.ListRulesInput{}

	if len(_route53recoveryreadinessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53recoveryreadinessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53recoveryreadinessNextToken) > 0 {
		input.NextToken = aws.String(_route53recoveryreadinessNextToken)
	}
	if len(_route53recoveryreadinessResourceType) > 0 {
		input.ResourceType = aws.String(_route53recoveryreadinessResourceType)
	}

	if disablePaginator() {
		if resp, err := client.ListRules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53recoveryreadiness.ListRulesOutput
	p := route53recoveryreadiness.NewListRulesPaginator(client, input)
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

// Lists the tags for a resource.
func route53recoveryreadiness_ListTagsForResources(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.ListTagsForResourcesInput{
		// ResourceArn: *string, // Required
	}

	if len(_route53recoveryreadinessResourceArn) > 0 {
		input.ResourceArn = aws.String(_route53recoveryreadinessResourceArn)
	}

	if resp, err := client.ListTagsForResources(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a tag to a resource.
func route53recoveryreadiness_TagResource(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_route53recoveryreadinessResourceArn) > 0 {
		input.ResourceArn = aws.String(_route53recoveryreadinessResourceArn)
	}
	if len(_route53recoveryreadinessTags) > 0 {
		if err := assignInputField(input, "Tags", _route53recoveryreadinessTags); err != nil {
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

// Removes a tag from a resource.
func route53recoveryreadiness_UntagResource(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_route53recoveryreadinessResourceArn) > 0 {
		input.ResourceArn = aws.String(_route53recoveryreadinessResourceArn)
	}
	if len(_route53recoveryreadinessTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _route53recoveryreadinessTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a cell to replace the list of nested cells with a new list of nested
// cells.
func route53recoveryreadiness_UpdateCell(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.UpdateCellInput{
		// CellName: *string, // Required
		// Cells: []string, // Required
	}

	if len(_route53recoveryreadinessCellName) > 0 {
		input.CellName = aws.String(_route53recoveryreadinessCellName)
	}
	if len(_route53recoveryreadinessCells) > 0 {
		input.Cells = append([]string(nil), _route53recoveryreadinessCells...)
	}

	if resp, err := client.UpdateCell(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a readiness check.
func route53recoveryreadiness_UpdateReadinessCheck(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.UpdateReadinessCheckInput{
		// ReadinessCheckName: *string, // Required
		// ResourceSetName: *string, // Required
	}

	if len(_route53recoveryreadinessReadinessCheckName) > 0 {
		input.ReadinessCheckName = aws.String(_route53recoveryreadinessReadinessCheckName)
	}
	if len(_route53recoveryreadinessResourceSetName) > 0 {
		input.ResourceSetName = aws.String(_route53recoveryreadinessResourceSetName)
	}

	if resp, err := client.UpdateReadinessCheck(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a recovery group.
func route53recoveryreadiness_UpdateRecoveryGroup(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.UpdateRecoveryGroupInput{
		// Cells: []string, // Required
		// RecoveryGroupName: *string, // Required
	}

	if len(_route53recoveryreadinessCells) > 0 {
		input.Cells = append([]string(nil), _route53recoveryreadinessCells...)
	}
	if len(_route53recoveryreadinessRecoveryGroupName) > 0 {
		input.RecoveryGroupName = aws.String(_route53recoveryreadinessRecoveryGroupName)
	}

	if resp, err := client.UpdateRecoveryGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a resource set.
func route53recoveryreadiness_UpdateResourceSet(cfg aws.Config, client *route53recoveryreadiness.Client) {
	input := &route53recoveryreadiness.UpdateResourceSetInput{
		// ResourceSetName: *string, // Required
		// ResourceSetType: *string, // Required
		// Resources: []types.Resource, // Required
	}

	if len(_route53recoveryreadinessResourceSetName) > 0 {
		input.ResourceSetName = aws.String(_route53recoveryreadinessResourceSetName)
	}
	if len(_route53recoveryreadinessResourceSetType) > 0 {
		input.ResourceSetType = aws.String(_route53recoveryreadinessResourceSetType)
	}
	if len(_route53recoveryreadinessResources) > 0 {
		if err := assignInputField(input, "Resources", _route53recoveryreadinessResources); err != nil {
			log.Errorf("invalid --resources: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateResourceSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_route53recoveryreadinessCmd)
	_route53recoveryreadinessCmd.Flags().SortFlags = false

	_route53recoveryreadinessCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_route53recoveryreadinessCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_route53recoveryreadinessCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_route53recoveryreadinessCmd.Flags().StringVarP(&_route53recoveryreadinessCellName, "cell-name", "", "", "Cell Name")
	_route53recoveryreadinessCmd.Flags().StringSliceVarP(&_route53recoveryreadinessCells, "cells", "", nil, "Cells")
	_route53recoveryreadinessCmd.Flags().StringVarP(&_route53recoveryreadinessCrossAccountAuthorization, "cross-account-authorization", "", "", "Cross Account Authorization")
	_route53recoveryreadinessCmd.Flags().StringVarP(&_route53recoveryreadinessMaxResults, "max-results", "", "", "Max Results")
	_route53recoveryreadinessCmd.Flags().StringVarP(&_route53recoveryreadinessNextToken, "next-token", "", "", "Next Token")
	_route53recoveryreadinessCmd.Flags().StringVarP(&_route53recoveryreadinessReadinessCheckName, "readiness-check-name", "", "", "Readiness Check Name")
	_route53recoveryreadinessCmd.Flags().StringVarP(&_route53recoveryreadinessRecoveryGroupName, "recovery-group-name", "", "", "Recovery Group Name")
	_route53recoveryreadinessCmd.Flags().StringVarP(&_route53recoveryreadinessResourceArn, "resource-arn", "", "", "Resource ARN")
	_route53recoveryreadinessCmd.Flags().StringVarP(&_route53recoveryreadinessResourceIdentifier, "resource-identifier", "", "", "Resource Identifier")
	_route53recoveryreadinessCmd.Flags().StringVarP(&_route53recoveryreadinessResourceSetName, "resource-set-name", "", "", "Resource Set Name")
	_route53recoveryreadinessCmd.Flags().StringVarP(&_route53recoveryreadinessResourceSetType, "resource-set-type", "", "", "Resource Set Type")
	_route53recoveryreadinessCmd.Flags().StringVarP(&_route53recoveryreadinessResourceType, "resource-type", "", "", "Resource Type")
	_route53recoveryreadinessCmd.Flags().StringVarP(&_route53recoveryreadinessResources, "resources", "", "", "Resources")
	_route53recoveryreadinessCmd.Flags().StringSliceVarP(&_route53recoveryreadinessTagKeys, "tag-keys", "", nil, "Tag Keys")
	_route53recoveryreadinessCmd.Flags().StringVarP(&_route53recoveryreadinessTags, "tags", "", "", "Tags")

	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessCreateCell, "create-cell", "", false, "Create Cell")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessCreateCrossAccountAuthorization, "create-cross-account-authorization", "", false, "Create Cross Account Authorization")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessCreateReadinessCheck, "create-readiness-check", "", false, "Create Readiness Check")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessCreateRecoveryGroup, "create-recovery-group", "", false, "Create Recovery Group")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessCreateResourceSet, "create-resource-set", "", false, "Create Resource Set")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessDeleteCell, "delete-cell", "", false, "Delete Cell")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessDeleteCrossAccountAuthorization, "delete-cross-account-authorization", "", false, "Delete Cross Account Authorization")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessDeleteReadinessCheck, "delete-readiness-check", "", false, "Delete Readiness Check")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessDeleteRecoveryGroup, "delete-recovery-group", "", false, "Delete Recovery Group")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessDeleteResourceSet, "delete-resource-set", "", false, "Delete Resource Set")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessGetArchitectureRecommendations, "get-architecture-recommendations", "", false, "Get Architecture Recommendations")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessGetCell, "get-cell", "", false, "Get Cell")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessGetCellReadinessSummary, "get-cell-readiness-summary", "", false, "Get Cell Readiness Summary")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessGetReadinessCheck, "get-readiness-check", "", false, "Get Readiness Check")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessGetReadinessCheckResourceStatus, "get-readiness-check-resource-status", "", false, "Get Readiness Check Resource Status")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessGetReadinessCheckStatus, "get-readiness-check-status", "", false, "Get Readiness Check Status")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessGetRecoveryGroup, "get-recovery-group", "", false, "Get Recovery Group")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessGetRecoveryGroupReadinessSummary, "get-recovery-group-readiness-summary", "", false, "Get Recovery Group Readiness Summary")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessGetResourceSet, "get-resource-set", "", false, "Get Resource Set")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessListCells, "list-cells", "", false, "List Cells")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessListCrossAccountAuthorizations, "list-cross-account-authorizations", "", false, "List Cross Account Authorizations")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessListReadinessChecks, "list-readiness-checks", "", false, "List Readiness Checks")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessListRecoveryGroups, "list-recovery-groups", "", false, "List Recovery Groups")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessListResourceSets, "list-resource-sets", "", false, "List Resource Sets")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessListRules, "list-rules", "", false, "List Rules")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessListTagsForResources, "list-tags-for-resources", "", false, "List Tags For Resources")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessTagResource, "tag-resource", "", false, "Tag Resource")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessUntagResource, "untag-resource", "", false, "Untag Resource")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessUpdateCell, "update-cell", "", false, "Update Cell")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessUpdateReadinessCheck, "update-readiness-check", "", false, "Update Readiness Check")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessUpdateRecoveryGroup, "update-recovery-group", "", false, "Update Recovery Group")
	_route53recoveryreadinessCmd.Flags().BoolVarP(&_route53recoveryreadinessUpdateResourceSet, "update-resource-set", "", false, "Update Resource Set")

}
