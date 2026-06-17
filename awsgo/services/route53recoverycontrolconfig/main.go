package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/route53recoverycontrolconfig/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-cluster", "create-control-panel", "create-routing-control", "create-safety-rule", "delete-cluster", "delete-control-panel", "delete-routing-control", "delete-safety-rule", "describe-cluster", "describe-control-panel", "describe-routing-control", "describe-safety-rule", "get-resource-policy", "list-associated-route53-health-checks", "list-clusters", "list-control-panels", "list-routing-controls", "list-safety-rules", "list-tags-for-resource", "tag-resource", "untag-resource", "update-cluster", "update-control-panel", "update-routing-control", "update-safety-rule"},
		OperationSet: map[string]bool{"create-cluster": true, "create-control-panel": true, "create-routing-control": true, "create-safety-rule": true, "delete-cluster": true, "delete-control-panel": true, "delete-routing-control": true, "delete-safety-rule": true, "describe-cluster": true, "describe-control-panel": true, "describe-routing-control": true, "describe-safety-rule": true, "get-resource-policy": true, "list-associated-route53-health-checks": true, "list-clusters": true, "list-control-panels": true, "list-routing-controls": true, "list-safety-rules": true, "list-tags-for-resource": true, "tag-resource": true, "untag-resource": true, "update-cluster": true, "update-control-panel": true, "update-routing-control": true, "update-safety-rule": true},
		OperationInputs: map[string][]string{
			"create-cluster":                        {"ClientToken", "ClusterName", "NetworkType", "Tags"},
			"create-control-panel":                  {"ClientToken", "ClusterArn", "ControlPanelName", "Tags"},
			"create-routing-control":                {"ClientToken", "ClusterArn", "ControlPanelArn", "RoutingControlName"},
			"create-safety-rule":                    {"AssertionRule", "ClientToken", "GatingRule", "Tags"},
			"delete-cluster":                        {"ClusterArn"},
			"delete-control-panel":                  {"ControlPanelArn"},
			"delete-routing-control":                {"RoutingControlArn"},
			"delete-safety-rule":                    {"SafetyRuleArn"},
			"describe-cluster":                      {"ClusterArn"},
			"describe-control-panel":                {"ControlPanelArn"},
			"describe-routing-control":              {"RoutingControlArn"},
			"describe-safety-rule":                  {"SafetyRuleArn"},
			"get-resource-policy":                   {"ResourceArn"},
			"list-associated-route53-health-checks": {"MaxResults", "NextToken", "RoutingControlArn"},
			"list-clusters":                         {"MaxResults", "NextToken"},
			"list-control-panels":                   {"ClusterArn", "MaxResults", "NextToken"},
			"list-routing-controls":                 {"ControlPanelArn", "MaxResults", "NextToken"},
			"list-safety-rules":                     {"ControlPanelArn", "MaxResults", "NextToken"},
			"list-tags-for-resource":                {"ResourceArn"},
			"tag-resource":                          {"ResourceArn", "Tags"},
			"untag-resource":                        {"ResourceArn", "TagKeys"},
			"update-cluster":                        {"ClusterArn", "NetworkType"},
			"update-control-panel":                  {"ControlPanelArn", "ControlPanelName"},
			"update-routing-control":                {"RoutingControlArn", "RoutingControlName"},
			"update-safety-rule":                    {"AssertionRuleUpdate", "GatingRuleUpdate"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-cluster":                        {"ClientToken": "*string", "ClusterName": "*string", "NetworkType": "types.NetworkType", "Tags": "map[string]string"},
			"create-control-panel":                  {"ClientToken": "*string", "ClusterArn": "*string", "ControlPanelName": "*string", "Tags": "map[string]string"},
			"create-routing-control":                {"ClientToken": "*string", "ClusterArn": "*string", "ControlPanelArn": "*string", "RoutingControlName": "*string"},
			"create-safety-rule":                    {"AssertionRule": "*types.NewAssertionRule", "ClientToken": "*string", "GatingRule": "*types.NewGatingRule", "Tags": "map[string]string"},
			"delete-cluster":                        {"ClusterArn": "*string"},
			"delete-control-panel":                  {"ControlPanelArn": "*string"},
			"delete-routing-control":                {"RoutingControlArn": "*string"},
			"delete-safety-rule":                    {"SafetyRuleArn": "*string"},
			"describe-cluster":                      {"ClusterArn": "*string"},
			"describe-control-panel":                {"ControlPanelArn": "*string"},
			"describe-routing-control":              {"RoutingControlArn": "*string"},
			"describe-safety-rule":                  {"SafetyRuleArn": "*string"},
			"get-resource-policy":                   {"ResourceArn": "*string"},
			"list-associated-route53-health-checks": {"MaxResults": "*int32", "NextToken": "*string", "RoutingControlArn": "*string"},
			"list-clusters":                         {"MaxResults": "*int32", "NextToken": "*string"},
			"list-control-panels":                   {"ClusterArn": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-routing-controls":                 {"ControlPanelArn": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-safety-rules":                     {"ControlPanelArn": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":                {"ResourceArn": "*string"},
			"tag-resource":                          {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                        {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-cluster":                        {"ClusterArn": "*string", "NetworkType": "types.NetworkType"},
			"update-control-panel":                  {"ControlPanelArn": "*string", "ControlPanelName": "*string"},
			"update-routing-control":                {"RoutingControlArn": "*string", "RoutingControlName": "*string"},
			"update-safety-rule":                    {"AssertionRuleUpdate": "*types.AssertionRuleUpdate", "GatingRuleUpdate": "*types.GatingRuleUpdate"},
		},
		OperationInputRequired: map[string][]string{
			"create-cluster":                        {"ClusterName"},
			"create-control-panel":                  {"ClusterArn", "ControlPanelName"},
			"create-routing-control":                {"ClusterArn", "RoutingControlName"},
			"create-safety-rule":                    {},
			"delete-cluster":                        {"ClusterArn"},
			"delete-control-panel":                  {"ControlPanelArn"},
			"delete-routing-control":                {"RoutingControlArn"},
			"delete-safety-rule":                    {"SafetyRuleArn"},
			"describe-cluster":                      {"ClusterArn"},
			"describe-control-panel":                {"ControlPanelArn"},
			"describe-routing-control":              {"RoutingControlArn"},
			"describe-safety-rule":                  {"SafetyRuleArn"},
			"get-resource-policy":                   {"ResourceArn"},
			"list-associated-route53-health-checks": {"RoutingControlArn"},
			"list-clusters":                         {},
			"list-control-panels":                   {},
			"list-routing-controls":                 {"ControlPanelArn"},
			"list-safety-rules":                     {"ControlPanelArn"},
			"list-tags-for-resource":                {"ResourceArn"},
			"tag-resource":                          {"ResourceArn", "Tags"},
			"untag-resource":                        {"ResourceArn", "TagKeys"},
			"update-cluster":                        {"ClusterArn", "NetworkType"},
			"update-control-panel":                  {"ControlPanelArn", "ControlPanelName"},
			"update-routing-control":                {"RoutingControlArn", "RoutingControlName"},
			"update-safety-rule":                    {},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("route53recoverycontrolconfig", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
