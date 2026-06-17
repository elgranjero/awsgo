package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53recoverycluster"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// route53recoveryclusterCmd represents the route53recoverycluster command
var _route53recoveryclusterCmd = &cobra.Command{
	Use:   "route53recoverycluster",
	Short: "AWS route53recoverycluster CLI",
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
		client := route53recoverycluster.NewFromConfig(cfg)
		if _route53recoveryclusterGetRoutingControlState {
			route53recoverycluster_GetRoutingControlState(cfg, client)
			return
		}
		if _route53recoveryclusterListRoutingControls {
			route53recoverycluster_ListRoutingControls(cfg, client)
			return
		}
		if _route53recoveryclusterUpdateRoutingControlState {
			route53recoverycluster_UpdateRoutingControlState(cfg, client)
			return
		}
		if _route53recoveryclusterUpdateRoutingControlStates {
			route53recoverycluster_UpdateRoutingControlStates(cfg, client)
			return
		}

	},
}

var (
	_route53recoveryclusterGetRoutingControlState     bool
	_route53recoveryclusterListRoutingControls        bool
	_route53recoveryclusterUpdateRoutingControlState  bool
	_route53recoveryclusterUpdateRoutingControlStates bool

	_route53recoveryclusterControlPanelArn                  string
	_route53recoveryclusterMaxResults                       string
	_route53recoveryclusterNextToken                        string
	_route53recoveryclusterRoutingControlArn                string
	_route53recoveryclusterRoutingControlState              string
	_route53recoveryclusterSafetyRulesToOverride            []string
	_route53recoveryclusterUpdateRoutingControlStateEntries string
)

// Get the state for a routing control. A routing control is a simple on/off
// switch that you can use to route traffic to cells. When a routing control state
// is set to ON, traffic flows to a cell. When the state is set to OFF, traffic
// does not flow.
//
// Before you can create a routing control, you must first create a cluster, and
// then host the control in a control panel on the cluster. For more information,
// see [Create routing control structures]in the Amazon Route 53 Application Recovery Controller Developer Guide. You
// access one of the endpoints for the cluster to get or update the routing control
// state to redirect traffic for your application.
//
// You must specify Regional endpoints when you work with API cluster operations
// to get or update routing control states in Route 53 ARC.
//
// To see a code example for getting a routing control state, including accessing
// Regional cluster endpoints in sequence, see [API examples]in the Amazon Route 53 Application
// Recovery Controller Developer Guide.
//
// Learn more about working with routing controls in the following topics in the
// Amazon Route 53 Application Recovery Controller Developer Guide:
//
// [Viewing and updating routing control states]
//
// [Working with routing controls in Route 53 ARC]
//
// [API examples]: https://docs.aws.amazon.com/r53recovery/latest/dg/service_code_examples_actions.html
// [Working with routing controls in Route 53 ARC]: https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.html
// [Create routing control structures]: https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.create.html
// [Viewing and updating routing control states]: https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.update.html
func route53recoverycluster_GetRoutingControlState(cfg aws.Config, client *route53recoverycluster.Client) {
	input := &route53recoverycluster.GetRoutingControlStateInput{
		// RoutingControlArn: *string, // Required
	}

	if len(_route53recoveryclusterRoutingControlArn) > 0 {
		input.RoutingControlArn = aws.String(_route53recoveryclusterRoutingControlArn)
	}

	if resp, err := client.GetRoutingControlState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List routing control names and Amazon Resource Names (ARNs), as well as the
// routing control state for each routing control, along with the control panel
// name and control panel ARN for the routing controls. If you specify a control
// panel ARN, this call lists the routing controls in the control panel. Otherwise,
// it lists all the routing controls in the cluster.
//
// A routing control is a simple on/off switch in Route 53 ARC that you can use to
// route traffic to cells. When a routing control state is set to ON, traffic flows
// to a cell. When the state is set to OFF, traffic does not flow.
//
// Before you can create a routing control, you must first create a cluster, and
// then host the control in a control panel on the cluster. For more information,
// see [Create routing control structures]in the Amazon Route 53 Application Recovery Controller Developer Guide. You
// access one of the endpoints for the cluster to get or update the routing control
// state to redirect traffic for your application.
//
// You must specify Regional endpoints when you work with API cluster operations
// to use this API operation to list routing controls in Route 53 ARC.
//
// Learn more about working with routing controls in the following topics in the
// Amazon Route 53 Application Recovery Controller Developer Guide:
//
// [Viewing and updating routing control states]
//
// [Working with routing controls in Route 53 ARC]
//
// [Working with routing controls in Route 53 ARC]: https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.html
// [Create routing control structures]: https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.create.html
// [Viewing and updating routing control states]: https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.update.html
func route53recoverycluster_ListRoutingControls(cfg aws.Config, client *route53recoverycluster.Client) {
	input := &route53recoverycluster.ListRoutingControlsInput{}

	if len(_route53recoveryclusterControlPanelArn) > 0 {
		input.ControlPanelArn = aws.String(_route53recoveryclusterControlPanelArn)
	}
	if len(_route53recoveryclusterMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53recoveryclusterMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53recoveryclusterNextToken) > 0 {
		input.NextToken = aws.String(_route53recoveryclusterNextToken)
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

	var results []*route53recoverycluster.ListRoutingControlsOutput
	p := route53recoverycluster.NewListRoutingControlsPaginator(client, input)
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

// Set the state of the routing control to reroute traffic. You can set the value
// to ON or OFF. When the state is ON, traffic flows to a cell. When the state is
// OFF, traffic does not flow.
//
// With Route 53 ARC, you can add safety rules for routing controls, which are
// safeguards for routing control state updates that help prevent unexpected
// outcomes, like fail open traffic routing. However, there are scenarios when you
// might want to bypass the routing control safeguards that are enforced with
// safety rules that you've configured. For example, you might want to fail over
// quickly for disaster recovery, and one or more safety rules might be
// unexpectedly preventing you from updating a routing control state to reroute
// traffic. In a "break glass" scenario like this, you can override one or more
// safety rules to change a routing control state and fail over your application.
//
// The SafetyRulesToOverride property enables you override one or more safety
// rules and update routing control states. For more information, see [Override safety rules to reroute traffic]in the
// Amazon Route 53 Application Recovery Controller Developer Guide.
//
// You must specify Regional endpoints when you work with API cluster operations
// to get or update routing control states in Route 53 ARC.
//
// To see a code example for getting a routing control state, including accessing
// Regional cluster endpoints in sequence, see [API examples]in the Amazon Route 53 Application
// Recovery Controller Developer Guide.
//
// [Viewing and updating routing control states]
//
// [Working with routing controls overall]
//
// [API examples]: https://docs.aws.amazon.com/r53recovery/latest/dg/service_code_examples_actions.html
// [Viewing and updating routing control states]: https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.update.html
// [Override safety rules to reroute traffic]: https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.override-safety-rule.html
// [Working with routing controls overall]: https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.html
func route53recoverycluster_UpdateRoutingControlState(cfg aws.Config, client *route53recoverycluster.Client) {
	input := &route53recoverycluster.UpdateRoutingControlStateInput{
		// RoutingControlArn: *string, // Required
		// RoutingControlState: types.RoutingControlState, // Required
	}

	if len(_route53recoveryclusterRoutingControlArn) > 0 {
		input.RoutingControlArn = aws.String(_route53recoveryclusterRoutingControlArn)
	}
	if len(_route53recoveryclusterRoutingControlState) > 0 {
		if err := assignInputField(input, "RoutingControlState", _route53recoveryclusterRoutingControlState); err != nil {
			log.Errorf("invalid --routing-control-state: %s", err.Error())
			return
		}
	}
	if len(_route53recoveryclusterSafetyRulesToOverride) > 0 {
		input.SafetyRulesToOverride = append([]string(nil), _route53recoveryclusterSafetyRulesToOverride...)
	}

	if resp, err := client.UpdateRoutingControlState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Set multiple routing control states. You can set the value for each state to be
// ON or OFF. When the state is ON, traffic flows to a cell. When it's OFF, traffic
// does not flow.
//
// With Route 53 ARC, you can add safety rules for routing controls, which are
// safeguards for routing control state updates that help prevent unexpected
// outcomes, like fail open traffic routing. However, there are scenarios when you
// might want to bypass the routing control safeguards that are enforced with
// safety rules that you've configured. For example, you might want to fail over
// quickly for disaster recovery, and one or more safety rules might be
// unexpectedly preventing you from updating a routing control state to reroute
// traffic. In a "break glass" scenario like this, you can override one or more
// safety rules to change a routing control state and fail over your application.
//
// The SafetyRulesToOverride property enables you override one or more safety
// rules and update routing control states. For more information, see [Override safety rules to reroute traffic]in the
// Amazon Route 53 Application Recovery Controller Developer Guide.
//
// You must specify Regional endpoints when you work with API cluster operations
// to get or update routing control states in Route 53 ARC.
//
// To see a code example for getting a routing control state, including accessing
// Regional cluster endpoints in sequence, see [API examples]in the Amazon Route 53 Application
// Recovery Controller Developer Guide.
//
// [Viewing and updating routing control states]
//
// [Working with routing controls overall]
//
// [API examples]: https://docs.aws.amazon.com/r53recovery/latest/dg/service_code_examples_actions.html
// [Viewing and updating routing control states]: https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.update.html
// [Override safety rules to reroute traffic]: https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.override-safety-rule.html
// [Working with routing controls overall]: https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.html
func route53recoverycluster_UpdateRoutingControlStates(cfg aws.Config, client *route53recoverycluster.Client) {
	input := &route53recoverycluster.UpdateRoutingControlStatesInput{
		// UpdateRoutingControlStateEntries: []types.UpdateRoutingControlStateEntry, // Required
	}

	if len(_route53recoveryclusterUpdateRoutingControlStateEntries) > 0 {
		if err := assignInputField(input, "UpdateRoutingControlStateEntries", _route53recoveryclusterUpdateRoutingControlStateEntries); err != nil {
			log.Errorf("invalid --update-routing-control-state-entries: %s", err.Error())
			return
		}
	}
	if len(_route53recoveryclusterSafetyRulesToOverride) > 0 {
		input.SafetyRulesToOverride = append([]string(nil), _route53recoveryclusterSafetyRulesToOverride...)
	}

	if resp, err := client.UpdateRoutingControlStates(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_route53recoveryclusterCmd)
	_route53recoveryclusterCmd.Flags().SortFlags = false

	_route53recoveryclusterCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_route53recoveryclusterCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_route53recoveryclusterCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_route53recoveryclusterCmd.Flags().StringVarP(&_route53recoveryclusterControlPanelArn, "control-panel-arn", "", "", "Control Panel ARN")
	_route53recoveryclusterCmd.Flags().StringVarP(&_route53recoveryclusterMaxResults, "max-results", "", "", "Max Results")
	_route53recoveryclusterCmd.Flags().StringVarP(&_route53recoveryclusterNextToken, "next-token", "", "", "Next Token")
	_route53recoveryclusterCmd.Flags().StringVarP(&_route53recoveryclusterRoutingControlArn, "routing-control-arn", "", "", "Routing Control ARN")
	_route53recoveryclusterCmd.Flags().StringVarP(&_route53recoveryclusterRoutingControlState, "routing-control-state", "", "", "Routing Control State")
	_route53recoveryclusterCmd.Flags().StringSliceVarP(&_route53recoveryclusterSafetyRulesToOverride, "safety-rules-to-override", "", nil, "Safety Rules To Override")
	_route53recoveryclusterCmd.Flags().StringVarP(&_route53recoveryclusterUpdateRoutingControlStateEntries, "update-routing-control-state-entries", "", "", "Update Routing Control State Entries")

	_route53recoveryclusterCmd.Flags().BoolVarP(&_route53recoveryclusterGetRoutingControlState, "get-routing-control-state", "", false, "Get Routing Control State")
	_route53recoveryclusterCmd.Flags().BoolVarP(&_route53recoveryclusterListRoutingControls, "list-routing-controls", "", false, "List Routing Controls")
	_route53recoveryclusterCmd.Flags().BoolVarP(&_route53recoveryclusterUpdateRoutingControlState, "update-routing-control-state", "", false, "Update Routing Control State")
	_route53recoveryclusterCmd.Flags().BoolVarP(&_route53recoveryclusterUpdateRoutingControlStates, "update-routing-control-states", "", false, "Update Routing Control States")

}
