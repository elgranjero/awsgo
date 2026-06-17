package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// route53recoverycontrolconfigCmd represents the route53recoverycontrolconfig command
var _route53recoverycontrolconfigCmd = &cobra.Command{
	Use:   "route53recoverycontrolconfig",
	Short: "AWS route53recoverycontrolconfig CLI",
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
		client := route53recoverycontrolconfig.NewFromConfig(cfg)
		if _route53recoverycontrolconfigCreateCluster {
			route53recoverycontrolconfig_CreateCluster(cfg, client)
			return
		}
		if _route53recoverycontrolconfigCreateControlPanel {
			route53recoverycontrolconfig_CreateControlPanel(cfg, client)
			return
		}
		if _route53recoverycontrolconfigCreateRoutingControl {
			route53recoverycontrolconfig_CreateRoutingControl(cfg, client)
			return
		}
		if _route53recoverycontrolconfigCreateSafetyRule {
			route53recoverycontrolconfig_CreateSafetyRule(cfg, client)
			return
		}
		if _route53recoverycontrolconfigDeleteCluster {
			route53recoverycontrolconfig_DeleteCluster(cfg, client)
			return
		}
		if _route53recoverycontrolconfigDeleteControlPanel {
			route53recoverycontrolconfig_DeleteControlPanel(cfg, client)
			return
		}
		if _route53recoverycontrolconfigDeleteRoutingControl {
			route53recoverycontrolconfig_DeleteRoutingControl(cfg, client)
			return
		}
		if _route53recoverycontrolconfigDeleteSafetyRule {
			route53recoverycontrolconfig_DeleteSafetyRule(cfg, client)
			return
		}
		if _route53recoverycontrolconfigDescribeCluster {
			route53recoverycontrolconfig_DescribeCluster(cfg, client)
			return
		}
		if _route53recoverycontrolconfigDescribeControlPanel {
			route53recoverycontrolconfig_DescribeControlPanel(cfg, client)
			return
		}
		if _route53recoverycontrolconfigDescribeRoutingControl {
			route53recoverycontrolconfig_DescribeRoutingControl(cfg, client)
			return
		}
		if _route53recoverycontrolconfigDescribeSafetyRule {
			route53recoverycontrolconfig_DescribeSafetyRule(cfg, client)
			return
		}
		if _route53recoverycontrolconfigGetResourcePolicy {
			route53recoverycontrolconfig_GetResourcePolicy(cfg, client)
			return
		}
		if _route53recoverycontrolconfigListAssociatedRoute53HealthChecks {
			route53recoverycontrolconfig_ListAssociatedRoute53HealthChecks(cfg, client)
			return
		}
		if _route53recoverycontrolconfigListClusters {
			route53recoverycontrolconfig_ListClusters(cfg, client)
			return
		}
		if _route53recoverycontrolconfigListControlPanels {
			route53recoverycontrolconfig_ListControlPanels(cfg, client)
			return
		}
		if _route53recoverycontrolconfigListRoutingControls {
			route53recoverycontrolconfig_ListRoutingControls(cfg, client)
			return
		}
		if _route53recoverycontrolconfigListSafetyRules {
			route53recoverycontrolconfig_ListSafetyRules(cfg, client)
			return
		}
		if _route53recoverycontrolconfigListTagsForResource {
			route53recoverycontrolconfig_ListTagsForResource(cfg, client)
			return
		}
		if _route53recoverycontrolconfigTagResource {
			route53recoverycontrolconfig_TagResource(cfg, client)
			return
		}
		if _route53recoverycontrolconfigUntagResource {
			route53recoverycontrolconfig_UntagResource(cfg, client)
			return
		}
		if _route53recoverycontrolconfigUpdateCluster {
			route53recoverycontrolconfig_UpdateCluster(cfg, client)
			return
		}
		if _route53recoverycontrolconfigUpdateControlPanel {
			route53recoverycontrolconfig_UpdateControlPanel(cfg, client)
			return
		}
		if _route53recoverycontrolconfigUpdateRoutingControl {
			route53recoverycontrolconfig_UpdateRoutingControl(cfg, client)
			return
		}
		if _route53recoverycontrolconfigUpdateSafetyRule {
			route53recoverycontrolconfig_UpdateSafetyRule(cfg, client)
			return
		}

	},
}

var (
	_route53recoverycontrolconfigCreateCluster                     bool
	_route53recoverycontrolconfigCreateControlPanel                bool
	_route53recoverycontrolconfigCreateRoutingControl              bool
	_route53recoverycontrolconfigCreateSafetyRule                  bool
	_route53recoverycontrolconfigDeleteCluster                     bool
	_route53recoverycontrolconfigDeleteControlPanel                bool
	_route53recoverycontrolconfigDeleteRoutingControl              bool
	_route53recoverycontrolconfigDeleteSafetyRule                  bool
	_route53recoverycontrolconfigDescribeCluster                   bool
	_route53recoverycontrolconfigDescribeControlPanel              bool
	_route53recoverycontrolconfigDescribeRoutingControl            bool
	_route53recoverycontrolconfigDescribeSafetyRule                bool
	_route53recoverycontrolconfigGetResourcePolicy                 bool
	_route53recoverycontrolconfigListAssociatedRoute53HealthChecks bool
	_route53recoverycontrolconfigListClusters                      bool
	_route53recoverycontrolconfigListControlPanels                 bool
	_route53recoverycontrolconfigListRoutingControls               bool
	_route53recoverycontrolconfigListSafetyRules                   bool
	_route53recoverycontrolconfigListTagsForResource               bool
	_route53recoverycontrolconfigTagResource                       bool
	_route53recoverycontrolconfigUntagResource                     bool
	_route53recoverycontrolconfigUpdateCluster                     bool
	_route53recoverycontrolconfigUpdateControlPanel                bool
	_route53recoverycontrolconfigUpdateRoutingControl              bool
	_route53recoverycontrolconfigUpdateSafetyRule                  bool

	_route53recoverycontrolconfigAssertionRule       string
	_route53recoverycontrolconfigAssertionRuleUpdate string
	_route53recoverycontrolconfigClientToken         string
	_route53recoverycontrolconfigClusterArn          string
	_route53recoverycontrolconfigClusterName         string
	_route53recoverycontrolconfigControlPanelArn     string
	_route53recoverycontrolconfigControlPanelName    string
	_route53recoverycontrolconfigGatingRule          string
	_route53recoverycontrolconfigGatingRuleUpdate    string
	_route53recoverycontrolconfigMaxResults          string
	_route53recoverycontrolconfigNetworkType         string
	_route53recoverycontrolconfigNextToken           string
	_route53recoverycontrolconfigResourceArn         string
	_route53recoverycontrolconfigRoutingControlArn   string
	_route53recoverycontrolconfigRoutingControlName  string
	_route53recoverycontrolconfigSafetyRuleArn       string
	_route53recoverycontrolconfigTagKeys             []string
	_route53recoverycontrolconfigTags                string
)

// Create a new cluster. A cluster is a set of redundant Regional endpoints
// against which you can run API calls to update or get the state of one or more
// routing controls. Each cluster has a name, status, Amazon Resource Name (ARN),
// and an array of the five cluster endpoints (one for each supported Amazon Web
// Services Region) that you can use with API calls to the cluster data plane.
func route53recoverycontrolconfig_CreateCluster(cfg aws.Config, client *route53recoverycontrolconfig.Client) {
	input := &route53recoverycontrolconfig.CreateClusterInput{
		// ClusterName: *string, // Required
	}

	if len(_route53recoverycontrolconfigClusterName) > 0 {
		input.ClusterName = aws.String(_route53recoverycontrolconfigClusterName)
	}
	if len(_route53recoverycontrolconfigClientToken) > 0 {
		input.ClientToken = aws.String(_route53recoverycontrolconfigClientToken)
	}
	if len(_route53recoverycontrolconfigNetworkType) > 0 {
		if err := assignInputField(input, "NetworkType", _route53recoverycontrolconfigNetworkType); err != nil {
			log.Errorf("invalid --network-type: %s", err.Error())
			return
		}
	}
	if len(_route53recoverycontrolconfigTags) > 0 {
		if err := assignInputField(input, "Tags", _route53recoverycontrolconfigTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new control panel. A control panel represents a group of routing
// controls that can be changed together in a single transaction. You can use a
// control panel to centrally view the operational status of applications across
// your organization, and trigger multi-app failovers in a single transaction, for
// example, to fail over an Availability Zone or Amazon Web Services Region.
func route53recoverycontrolconfig_CreateControlPanel(cfg aws.Config, client *route53recoverycontrolconfig.Client) {
	input := &route53recoverycontrolconfig.CreateControlPanelInput{
		// ClusterArn: *string, // Required
		// ControlPanelName: *string, // Required
	}

	if len(_route53recoverycontrolconfigClusterArn) > 0 {
		input.ClusterArn = aws.String(_route53recoverycontrolconfigClusterArn)
	}
	if len(_route53recoverycontrolconfigControlPanelName) > 0 {
		input.ControlPanelName = aws.String(_route53recoverycontrolconfigControlPanelName)
	}
	if len(_route53recoverycontrolconfigClientToken) > 0 {
		input.ClientToken = aws.String(_route53recoverycontrolconfigClientToken)
	}
	if len(_route53recoverycontrolconfigTags) > 0 {
		if err := assignInputField(input, "Tags", _route53recoverycontrolconfigTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateControlPanel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new routing control.
// A routing control has one of two states: ON and OFF. You can map the routing
// control state to the state of an Amazon Route 53 health check, which can be used
// to control traffic routing.
//
// To get or update the routing control state, see the Recovery Cluster (data
// plane) API actions for Amazon Route 53 Application Recovery Controller.
func route53recoverycontrolconfig_CreateRoutingControl(cfg aws.Config, client *route53recoverycontrolconfig.Client) {
	input := &route53recoverycontrolconfig.CreateRoutingControlInput{
		// ClusterArn: *string, // Required
		// RoutingControlName: *string, // Required
	}

	if len(_route53recoverycontrolconfigClusterArn) > 0 {
		input.ClusterArn = aws.String(_route53recoverycontrolconfigClusterArn)
	}
	if len(_route53recoverycontrolconfigRoutingControlName) > 0 {
		input.RoutingControlName = aws.String(_route53recoverycontrolconfigRoutingControlName)
	}
	if len(_route53recoverycontrolconfigClientToken) > 0 {
		input.ClientToken = aws.String(_route53recoverycontrolconfigClientToken)
	}
	if len(_route53recoverycontrolconfigControlPanelArn) > 0 {
		input.ControlPanelArn = aws.String(_route53recoverycontrolconfigControlPanelArn)
	}

	if resp, err := client.CreateRoutingControl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a safety rule in a control panel. Safety rules let you add safeguards
// around changing routing control states, and for enabling and disabling routing
// controls, to help prevent unexpected outcomes.
//
// There are two types of safety rules: assertion rules and gating rules.
//
// Assertion rule: An assertion rule enforces that, when you change a routing
// control state, that a certain criteria is met. For example, the criteria might
// be that at least one routing control state is On after the transaction so that
// traffic continues to flow to at least one cell for the application. This ensures
// that you avoid a fail-open scenario.
//
// Gating rule: A gating rule lets you configure a gating routing control as an
// overall "on/off" switch for a group of routing controls. Or, you can configure
// more complex gating scenarios, for example by configuring multiple gating
// routing controls.
//
// For more information, see [Safety rules] in the Amazon Route 53 Application Recovery
// Controller Developer Guide.
//
// [Safety rules]: https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.safety-rules.html
func route53recoverycontrolconfig_CreateSafetyRule(cfg aws.Config, client *route53recoverycontrolconfig.Client) {
	input := &route53recoverycontrolconfig.CreateSafetyRuleInput{}

	if len(_route53recoverycontrolconfigAssertionRule) > 0 {
		if err := assignInputField(input, "AssertionRule", _route53recoverycontrolconfigAssertionRule); err != nil {
			log.Errorf("invalid --assertion-rule: %s", err.Error())
			return
		}
	}
	if len(_route53recoverycontrolconfigClientToken) > 0 {
		input.ClientToken = aws.String(_route53recoverycontrolconfigClientToken)
	}
	if len(_route53recoverycontrolconfigGatingRule) > 0 {
		if err := assignInputField(input, "GatingRule", _route53recoverycontrolconfigGatingRule); err != nil {
			log.Errorf("invalid --gating-rule: %s", err.Error())
			return
		}
	}
	if len(_route53recoverycontrolconfigTags) > 0 {
		if err := assignInputField(input, "Tags", _route53recoverycontrolconfigTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSafetyRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a cluster.
func route53recoverycontrolconfig_DeleteCluster(cfg aws.Config, client *route53recoverycontrolconfig.Client) {
	input := &route53recoverycontrolconfig.DeleteClusterInput{
		// ClusterArn: *string, // Required
	}

	if len(_route53recoverycontrolconfigClusterArn) > 0 {
		input.ClusterArn = aws.String(_route53recoverycontrolconfigClusterArn)
	}

	if resp, err := client.DeleteCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a control panel.
func route53recoverycontrolconfig_DeleteControlPanel(cfg aws.Config, client *route53recoverycontrolconfig.Client) {
	input := &route53recoverycontrolconfig.DeleteControlPanelInput{
		// ControlPanelArn: *string, // Required
	}

	if len(_route53recoverycontrolconfigControlPanelArn) > 0 {
		input.ControlPanelArn = aws.String(_route53recoverycontrolconfigControlPanelArn)
	}

	if resp, err := client.DeleteControlPanel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a routing control.
func route53recoverycontrolconfig_DeleteRoutingControl(cfg aws.Config, client *route53recoverycontrolconfig.Client) {
	input := &route53recoverycontrolconfig.DeleteRoutingControlInput{
		// RoutingControlArn: *string, // Required
	}

	if len(_route53recoverycontrolconfigRoutingControlArn) > 0 {
		input.RoutingControlArn = aws.String(_route53recoverycontrolconfigRoutingControlArn)
	}

	if resp, err := client.DeleteRoutingControl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a safety rule.
// />
func route53recoverycontrolconfig_DeleteSafetyRule(cfg aws.Config, client *route53recoverycontrolconfig.Client) {
	input := &route53recoverycontrolconfig.DeleteSafetyRuleInput{
		// SafetyRuleArn: *string, // Required
	}

	if len(_route53recoverycontrolconfigSafetyRuleArn) > 0 {
		input.SafetyRuleArn = aws.String(_route53recoverycontrolconfigSafetyRuleArn)
	}

	if resp, err := client.DeleteSafetyRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Display the details about a cluster. The response includes the cluster name,
// endpoints, status, and Amazon Resource Name (ARN).
func route53recoverycontrolconfig_DescribeCluster(cfg aws.Config, client *route53recoverycontrolconfig.Client) {
	input := &route53recoverycontrolconfig.DescribeClusterInput{
		// ClusterArn: *string, // Required
	}

	if len(_route53recoverycontrolconfigClusterArn) > 0 {
		input.ClusterArn = aws.String(_route53recoverycontrolconfigClusterArn)
	}

	if resp, err := client.DescribeCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays details about a control panel.
func route53recoverycontrolconfig_DescribeControlPanel(cfg aws.Config, client *route53recoverycontrolconfig.Client) {
	input := &route53recoverycontrolconfig.DescribeControlPanelInput{
		// ControlPanelArn: *string, // Required
	}

	if len(_route53recoverycontrolconfigControlPanelArn) > 0 {
		input.ControlPanelArn = aws.String(_route53recoverycontrolconfigControlPanelArn)
	}

	if resp, err := client.DescribeControlPanel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays details about a routing control. A routing control has one of two
// states: ON and OFF. You can map the routing control state to the state of an
// Amazon Route 53 health check, which can be used to control routing.
//
// To get or update the routing control state, see the Recovery Cluster (data
// plane) API actions for Amazon Route 53 Application Recovery Controller.
func route53recoverycontrolconfig_DescribeRoutingControl(cfg aws.Config, client *route53recoverycontrolconfig.Client) {
	input := &route53recoverycontrolconfig.DescribeRoutingControlInput{
		// RoutingControlArn: *string, // Required
	}

	if len(_route53recoverycontrolconfigRoutingControlArn) > 0 {
		input.RoutingControlArn = aws.String(_route53recoverycontrolconfigRoutingControlArn)
	}

	if resp, err := client.DescribeRoutingControl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a safety rule.
func route53recoverycontrolconfig_DescribeSafetyRule(cfg aws.Config, client *route53recoverycontrolconfig.Client) {
	input := &route53recoverycontrolconfig.DescribeSafetyRuleInput{
		// SafetyRuleArn: *string, // Required
	}

	if len(_route53recoverycontrolconfigSafetyRuleArn) > 0 {
		input.SafetyRuleArn = aws.String(_route53recoverycontrolconfigSafetyRuleArn)
	}

	if resp, err := client.DescribeSafetyRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get information about the resource policy for a cluster.
func route53recoverycontrolconfig_GetResourcePolicy(cfg aws.Config, client *route53recoverycontrolconfig.Client) {
	input := &route53recoverycontrolconfig.GetResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_route53recoverycontrolconfigResourceArn) > 0 {
		input.ResourceArn = aws.String(_route53recoverycontrolconfigResourceArn)
	}

	if resp, err := client.GetResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an array of all Amazon Route 53 health checks associated with a
// specific routing control.
func route53recoverycontrolconfig_ListAssociatedRoute53HealthChecks(cfg aws.Config, client *route53recoverycontrolconfig.Client) {
	input := &route53recoverycontrolconfig.ListAssociatedRoute53HealthChecksInput{
		// RoutingControlArn: *string, // Required
	}

	if len(_route53recoverycontrolconfigRoutingControlArn) > 0 {
		input.RoutingControlArn = aws.String(_route53recoverycontrolconfigRoutingControlArn)
	}
	if len(_route53recoverycontrolconfigMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53recoverycontrolconfigMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53recoverycontrolconfigNextToken) > 0 {
		input.NextToken = aws.String(_route53recoverycontrolconfigNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssociatedRoute53HealthChecks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53recoverycontrolconfig.ListAssociatedRoute53HealthChecksOutput
	p := route53recoverycontrolconfig.NewListAssociatedRoute53HealthChecksPaginator(client, input)
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

// Returns an array of all the clusters in an account.
func route53recoverycontrolconfig_ListClusters(cfg aws.Config, client *route53recoverycontrolconfig.Client) {
	input := &route53recoverycontrolconfig.ListClustersInput{}

	if len(_route53recoverycontrolconfigMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53recoverycontrolconfigMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53recoverycontrolconfigNextToken) > 0 {
		input.NextToken = aws.String(_route53recoverycontrolconfigNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListClusters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53recoverycontrolconfig.ListClustersOutput
	p := route53recoverycontrolconfig.NewListClustersPaginator(client, input)
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

// Returns an array of control panels in an account or in a cluster.
func route53recoverycontrolconfig_ListControlPanels(cfg aws.Config, client *route53recoverycontrolconfig.Client) {
	input := &route53recoverycontrolconfig.ListControlPanelsInput{}

	if len(_route53recoverycontrolconfigClusterArn) > 0 {
		input.ClusterArn = aws.String(_route53recoverycontrolconfigClusterArn)
	}
	if len(_route53recoverycontrolconfigMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53recoverycontrolconfigMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53recoverycontrolconfigNextToken) > 0 {
		input.NextToken = aws.String(_route53recoverycontrolconfigNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListControlPanels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53recoverycontrolconfig.ListControlPanelsOutput
	p := route53recoverycontrolconfig.NewListControlPanelsPaginator(client, input)
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

// Returns an array of routing controls for a control panel. A routing control is
// an Amazon Route 53 Application Recovery Controller construct that has one of two
// states: ON and OFF. You can map the routing control state to the state of an
// Amazon Route 53 health check, which can be used to control routing.
func route53recoverycontrolconfig_ListRoutingControls(cfg aws.Config, client *route53recoverycontrolconfig.Client) {
	input := &route53recoverycontrolconfig.ListRoutingControlsInput{
		// ControlPanelArn: *string, // Required
	}

	if len(_route53recoverycontrolconfigControlPanelArn) > 0 {
		input.ControlPanelArn = aws.String(_route53recoverycontrolconfigControlPanelArn)
	}
	if len(_route53recoverycontrolconfigMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53recoverycontrolconfigMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53recoverycontrolconfigNextToken) > 0 {
		input.NextToken = aws.String(_route53recoverycontrolconfigNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRoutingControls(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53recoverycontrolconfig.ListRoutingControlsOutput
	p := route53recoverycontrolconfig.NewListRoutingControlsPaginator(client, input)
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

// List the safety rules (the assertion rules and gating rules) that you've
// defined for the routing controls in a control panel.
func route53recoverycontrolconfig_ListSafetyRules(cfg aws.Config, client *route53recoverycontrolconfig.Client) {
	input := &route53recoverycontrolconfig.ListSafetyRulesInput{
		// ControlPanelArn: *string, // Required
	}

	if len(_route53recoverycontrolconfigControlPanelArn) > 0 {
		input.ControlPanelArn = aws.String(_route53recoverycontrolconfigControlPanelArn)
	}
	if len(_route53recoverycontrolconfigMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53recoverycontrolconfigMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53recoverycontrolconfigNextToken) > 0 {
		input.NextToken = aws.String(_route53recoverycontrolconfigNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSafetyRules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53recoverycontrolconfig.ListSafetyRulesOutput
	p := route53recoverycontrolconfig.NewListSafetyRulesPaginator(client, input)
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
func route53recoverycontrolconfig_ListTagsForResource(cfg aws.Config, client *route53recoverycontrolconfig.Client) {
	input := &route53recoverycontrolconfig.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_route53recoverycontrolconfigResourceArn) > 0 {
		input.ResourceArn = aws.String(_route53recoverycontrolconfigResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a tag to a resource.
func route53recoverycontrolconfig_TagResource(cfg aws.Config, client *route53recoverycontrolconfig.Client) {
	input := &route53recoverycontrolconfig.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_route53recoverycontrolconfigResourceArn) > 0 {
		input.ResourceArn = aws.String(_route53recoverycontrolconfigResourceArn)
	}
	if len(_route53recoverycontrolconfigTags) > 0 {
		if err := assignInputField(input, "Tags", _route53recoverycontrolconfigTags); err != nil {
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
func route53recoverycontrolconfig_UntagResource(cfg aws.Config, client *route53recoverycontrolconfig.Client) {
	input := &route53recoverycontrolconfig.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_route53recoverycontrolconfigResourceArn) > 0 {
		input.ResourceArn = aws.String(_route53recoverycontrolconfigResourceArn)
	}
	if len(_route53recoverycontrolconfigTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _route53recoverycontrolconfigTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing cluster. You can only update the network type of a cluster.
func route53recoverycontrolconfig_UpdateCluster(cfg aws.Config, client *route53recoverycontrolconfig.Client) {
	input := &route53recoverycontrolconfig.UpdateClusterInput{
		// ClusterArn: *string, // Required
		// NetworkType: types.NetworkType, // Required
	}

	if len(_route53recoverycontrolconfigClusterArn) > 0 {
		input.ClusterArn = aws.String(_route53recoverycontrolconfigClusterArn)
	}
	if len(_route53recoverycontrolconfigNetworkType) > 0 {
		if err := assignInputField(input, "NetworkType", _route53recoverycontrolconfigNetworkType); err != nil {
			log.Errorf("invalid --network-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a control panel. The only update you can make to a control panel is to
// change the name of the control panel.
func route53recoverycontrolconfig_UpdateControlPanel(cfg aws.Config, client *route53recoverycontrolconfig.Client) {
	input := &route53recoverycontrolconfig.UpdateControlPanelInput{
		// ControlPanelArn: *string, // Required
		// ControlPanelName: *string, // Required
	}

	if len(_route53recoverycontrolconfigControlPanelArn) > 0 {
		input.ControlPanelArn = aws.String(_route53recoverycontrolconfigControlPanelArn)
	}
	if len(_route53recoverycontrolconfigControlPanelName) > 0 {
		input.ControlPanelName = aws.String(_route53recoverycontrolconfigControlPanelName)
	}

	if resp, err := client.UpdateControlPanel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a routing control. You can only update the name of the routing control.
// To get or update the routing control state, see the Recovery Cluster (data
// plane) API actions for Amazon Route 53 Application Recovery Controller.
func route53recoverycontrolconfig_UpdateRoutingControl(cfg aws.Config, client *route53recoverycontrolconfig.Client) {
	input := &route53recoverycontrolconfig.UpdateRoutingControlInput{
		// RoutingControlArn: *string, // Required
		// RoutingControlName: *string, // Required
	}

	if len(_route53recoverycontrolconfigRoutingControlArn) > 0 {
		input.RoutingControlArn = aws.String(_route53recoverycontrolconfigRoutingControlArn)
	}
	if len(_route53recoverycontrolconfigRoutingControlName) > 0 {
		input.RoutingControlName = aws.String(_route53recoverycontrolconfigRoutingControlName)
	}

	if resp, err := client.UpdateRoutingControl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a safety rule (an assertion rule or gating rule). You can only update
// the name and the waiting period for a safety rule. To make other updates, delete
// the safety rule and create a new one.
func route53recoverycontrolconfig_UpdateSafetyRule(cfg aws.Config, client *route53recoverycontrolconfig.Client) {
	input := &route53recoverycontrolconfig.UpdateSafetyRuleInput{}

	if len(_route53recoverycontrolconfigAssertionRuleUpdate) > 0 {
		if err := assignInputField(input, "AssertionRuleUpdate", _route53recoverycontrolconfigAssertionRuleUpdate); err != nil {
			log.Errorf("invalid --assertion-rule-update: %s", err.Error())
			return
		}
	}
	if len(_route53recoverycontrolconfigGatingRuleUpdate) > 0 {
		if err := assignInputField(input, "GatingRuleUpdate", _route53recoverycontrolconfigGatingRuleUpdate); err != nil {
			log.Errorf("invalid --gating-rule-update: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSafetyRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_route53recoverycontrolconfigCmd)
	_route53recoverycontrolconfigCmd.Flags().SortFlags = false

	_route53recoverycontrolconfigCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_route53recoverycontrolconfigCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_route53recoverycontrolconfigCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_route53recoverycontrolconfigCmd.Flags().StringVarP(&_route53recoverycontrolconfigAssertionRule, "assertion-rule", "", "", "Assertion Rule")
	_route53recoverycontrolconfigCmd.Flags().StringVarP(&_route53recoverycontrolconfigAssertionRuleUpdate, "assertion-rule-update", "", "", "Assertion Rule Update")
	_route53recoverycontrolconfigCmd.Flags().StringVarP(&_route53recoverycontrolconfigClientToken, "client-token", "", "", "Client Token")
	_route53recoverycontrolconfigCmd.Flags().StringVarP(&_route53recoverycontrolconfigClusterArn, "cluster-arn", "", "", "Cluster ARN")
	_route53recoverycontrolconfigCmd.Flags().StringVarP(&_route53recoverycontrolconfigClusterName, "cluster-name", "", "", "Cluster Name")
	_route53recoverycontrolconfigCmd.Flags().StringVarP(&_route53recoverycontrolconfigControlPanelArn, "control-panel-arn", "", "", "Control Panel ARN")
	_route53recoverycontrolconfigCmd.Flags().StringVarP(&_route53recoverycontrolconfigControlPanelName, "control-panel-name", "", "", "Control Panel Name")
	_route53recoverycontrolconfigCmd.Flags().StringVarP(&_route53recoverycontrolconfigGatingRule, "gating-rule", "", "", "Gating Rule")
	_route53recoverycontrolconfigCmd.Flags().StringVarP(&_route53recoverycontrolconfigGatingRuleUpdate, "gating-rule-update", "", "", "Gating Rule Update")
	_route53recoverycontrolconfigCmd.Flags().StringVarP(&_route53recoverycontrolconfigMaxResults, "max-results", "", "", "Max Results")
	_route53recoverycontrolconfigCmd.Flags().StringVarP(&_route53recoverycontrolconfigNetworkType, "network-type", "", "", "Network Type")
	_route53recoverycontrolconfigCmd.Flags().StringVarP(&_route53recoverycontrolconfigNextToken, "next-token", "", "", "Next Token")
	_route53recoverycontrolconfigCmd.Flags().StringVarP(&_route53recoverycontrolconfigResourceArn, "resource-arn", "", "", "Resource ARN")
	_route53recoverycontrolconfigCmd.Flags().StringVarP(&_route53recoverycontrolconfigRoutingControlArn, "routing-control-arn", "", "", "Routing Control ARN")
	_route53recoverycontrolconfigCmd.Flags().StringVarP(&_route53recoverycontrolconfigRoutingControlName, "routing-control-name", "", "", "Routing Control Name")
	_route53recoverycontrolconfigCmd.Flags().StringVarP(&_route53recoverycontrolconfigSafetyRuleArn, "safety-rule-arn", "", "", "Safety Rule ARN")
	_route53recoverycontrolconfigCmd.Flags().StringSliceVarP(&_route53recoverycontrolconfigTagKeys, "tag-keys", "", nil, "Tag Keys")
	_route53recoverycontrolconfigCmd.Flags().StringVarP(&_route53recoverycontrolconfigTags, "tags", "", "", "Tags")

	_route53recoverycontrolconfigCmd.Flags().BoolVarP(&_route53recoverycontrolconfigCreateCluster, "create-cluster", "", false, "Create Cluster")
	_route53recoverycontrolconfigCmd.Flags().BoolVarP(&_route53recoverycontrolconfigCreateControlPanel, "create-control-panel", "", false, "Create Control Panel")
	_route53recoverycontrolconfigCmd.Flags().BoolVarP(&_route53recoverycontrolconfigCreateRoutingControl, "create-routing-control", "", false, "Create Routing Control")
	_route53recoverycontrolconfigCmd.Flags().BoolVarP(&_route53recoverycontrolconfigCreateSafetyRule, "create-safety-rule", "", false, "Create Safety Rule")
	_route53recoverycontrolconfigCmd.Flags().BoolVarP(&_route53recoverycontrolconfigDeleteCluster, "delete-cluster", "", false, "Delete Cluster")
	_route53recoverycontrolconfigCmd.Flags().BoolVarP(&_route53recoverycontrolconfigDeleteControlPanel, "delete-control-panel", "", false, "Delete Control Panel")
	_route53recoverycontrolconfigCmd.Flags().BoolVarP(&_route53recoverycontrolconfigDeleteRoutingControl, "delete-routing-control", "", false, "Delete Routing Control")
	_route53recoverycontrolconfigCmd.Flags().BoolVarP(&_route53recoverycontrolconfigDeleteSafetyRule, "delete-safety-rule", "", false, "Delete Safety Rule")
	_route53recoverycontrolconfigCmd.Flags().BoolVarP(&_route53recoverycontrolconfigDescribeCluster, "describe-cluster", "", false, "Describe Cluster")
	_route53recoverycontrolconfigCmd.Flags().BoolVarP(&_route53recoverycontrolconfigDescribeControlPanel, "describe-control-panel", "", false, "Describe Control Panel")
	_route53recoverycontrolconfigCmd.Flags().BoolVarP(&_route53recoverycontrolconfigDescribeRoutingControl, "describe-routing-control", "", false, "Describe Routing Control")
	_route53recoverycontrolconfigCmd.Flags().BoolVarP(&_route53recoverycontrolconfigDescribeSafetyRule, "describe-safety-rule", "", false, "Describe Safety Rule")
	_route53recoverycontrolconfigCmd.Flags().BoolVarP(&_route53recoverycontrolconfigGetResourcePolicy, "get-resource-policy", "", false, "Get Resource Policy")
	_route53recoverycontrolconfigCmd.Flags().BoolVarP(&_route53recoverycontrolconfigListAssociatedRoute53HealthChecks, "list-associated-route53-health-checks", "", false, "List Associated Route53 Health Checks")
	_route53recoverycontrolconfigCmd.Flags().BoolVarP(&_route53recoverycontrolconfigListClusters, "list-clusters", "", false, "List Clusters")
	_route53recoverycontrolconfigCmd.Flags().BoolVarP(&_route53recoverycontrolconfigListControlPanels, "list-control-panels", "", false, "List Control Panels")
	_route53recoverycontrolconfigCmd.Flags().BoolVarP(&_route53recoverycontrolconfigListRoutingControls, "list-routing-controls", "", false, "List Routing Controls")
	_route53recoverycontrolconfigCmd.Flags().BoolVarP(&_route53recoverycontrolconfigListSafetyRules, "list-safety-rules", "", false, "List Safety Rules")
	_route53recoverycontrolconfigCmd.Flags().BoolVarP(&_route53recoverycontrolconfigListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_route53recoverycontrolconfigCmd.Flags().BoolVarP(&_route53recoverycontrolconfigTagResource, "tag-resource", "", false, "Tag Resource")
	_route53recoverycontrolconfigCmd.Flags().BoolVarP(&_route53recoverycontrolconfigUntagResource, "untag-resource", "", false, "Untag Resource")
	_route53recoverycontrolconfigCmd.Flags().BoolVarP(&_route53recoverycontrolconfigUpdateCluster, "update-cluster", "", false, "Update Cluster")
	_route53recoverycontrolconfigCmd.Flags().BoolVarP(&_route53recoverycontrolconfigUpdateControlPanel, "update-control-panel", "", false, "Update Control Panel")
	_route53recoverycontrolconfigCmd.Flags().BoolVarP(&_route53recoverycontrolconfigUpdateRoutingControl, "update-routing-control", "", false, "Update Routing Control")
	_route53recoverycontrolconfigCmd.Flags().BoolVarP(&_route53recoverycontrolconfigUpdateSafetyRule, "update-safety-rule", "", false, "Update Safety Rule")

}
