package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dlm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// dlmCmd represents the dlm command
var _dlmCmd = &cobra.Command{
	Use:   "dlm",
	Short: "AWS dlm CLI",
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
		client := dlm.NewFromConfig(cfg)
		if _dlmCreateLifecyclePolicy {
			dlm_CreateLifecyclePolicy(cfg, client)
			return
		}
		if _dlmDeleteLifecyclePolicy {
			dlm_DeleteLifecyclePolicy(cfg, client)
			return
		}
		if _dlmGetLifecyclePolicies {
			dlm_GetLifecyclePolicies(cfg, client)
			return
		}
		if _dlmGetLifecyclePolicy {
			dlm_GetLifecyclePolicy(cfg, client)
			return
		}
		if _dlmListTagsForResource {
			dlm_ListTagsForResource(cfg, client)
			return
		}
		if _dlmTagResource {
			dlm_TagResource(cfg, client)
			return
		}
		if _dlmUntagResource {
			dlm_UntagResource(cfg, client)
			return
		}
		if _dlmUpdateLifecyclePolicy {
			dlm_UpdateLifecyclePolicy(cfg, client)
			return
		}

	},
}

var (
	_dlmCreateLifecyclePolicy bool
	_dlmDeleteLifecyclePolicy bool
	_dlmGetLifecyclePolicies  bool
	_dlmGetLifecyclePolicy    bool
	_dlmListTagsForResource   bool
	_dlmTagResource           bool
	_dlmUntagResource         bool
	_dlmUpdateLifecyclePolicy bool

	_dlmCopyTags               string
	_dlmCreateInterval         string
	_dlmCrossRegionCopyTargets string
	_dlmDefaultPolicy          string
	_dlmDefaultPolicyType      string
	_dlmDescription            string
	_dlmExclusions             string
	_dlmExecutionRoleArn       string
	_dlmExtendDeletion         string
	_dlmPolicyDetails          string
	_dlmPolicyId               string
	_dlmPolicyIds              []string
	_dlmResourceArn            string
	_dlmResourceTypes          string
	_dlmRetainInterval         string
	_dlmState                  string
	_dlmTagKeys                []string
	_dlmTags                   string
	_dlmTagsToAdd              []string
	_dlmTargetTags             []string
)

// Creates an Amazon Data Lifecycle Manager lifecycle policy. Amazon Data
// Lifecycle Manager supports the following policy types:
//
// - Custom EBS snapshot policy
//
// - Custom EBS-backed AMI policy
//
// - Cross-account copy event policy
//
// - Default policy for EBS snapshots
//
// - Default policy for EBS-backed AMIs
//
// For more information, see [Default policies vs custom policies].
//
// If you create a default policy, you can specify the request parameters either
// in the request body, or in the PolicyDetails request structure, but not both.
//
// [Default policies vs custom policies]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/policy-differences.html
func dlm_CreateLifecyclePolicy(cfg aws.Config, client *dlm.Client) {
	input := &dlm.CreateLifecyclePolicyInput{
		// Description: *string, // Required
		// ExecutionRoleArn: *string, // Required
		// State: types.SettablePolicyStateValues, // Required
	}

	if len(_dlmDescription) > 0 {
		input.Description = aws.String(_dlmDescription)
	}
	if len(_dlmExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_dlmExecutionRoleArn)
	}
	if len(_dlmState) > 0 {
		if err := assignInputField(input, "State", _dlmState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}
	if len(_dlmCopyTags) > 0 {
		if err := assignInputField(input, "CopyTags", _dlmCopyTags); err != nil {
			log.Errorf("invalid --copy-tags: %s", err.Error())
			return
		}
	}
	if len(_dlmCreateInterval) > 0 {
		if err := assignInputField(input, "CreateInterval", _dlmCreateInterval); err != nil {
			log.Errorf("invalid --create-interval: %s", err.Error())
			return
		}
	}
	if len(_dlmCrossRegionCopyTargets) > 0 {
		if err := assignInputField(input, "CrossRegionCopyTargets", _dlmCrossRegionCopyTargets); err != nil {
			log.Errorf("invalid --cross-region-copy-targets: %s", err.Error())
			return
		}
	}
	if len(_dlmDefaultPolicy) > 0 {
		if err := assignInputField(input, "DefaultPolicy", _dlmDefaultPolicy); err != nil {
			log.Errorf("invalid --default-policy: %s", err.Error())
			return
		}
	}
	if len(_dlmExclusions) > 0 {
		if err := assignInputField(input, "Exclusions", _dlmExclusions); err != nil {
			log.Errorf("invalid --exclusions: %s", err.Error())
			return
		}
	}
	if len(_dlmExtendDeletion) > 0 {
		if err := assignInputField(input, "ExtendDeletion", _dlmExtendDeletion); err != nil {
			log.Errorf("invalid --extend-deletion: %s", err.Error())
			return
		}
	}
	if len(_dlmPolicyDetails) > 0 {
		if err := assignInputField(input, "PolicyDetails", _dlmPolicyDetails); err != nil {
			log.Errorf("invalid --policy-details: %s", err.Error())
			return
		}
	}
	if len(_dlmRetainInterval) > 0 {
		if err := assignInputField(input, "RetainInterval", _dlmRetainInterval); err != nil {
			log.Errorf("invalid --retain-interval: %s", err.Error())
			return
		}
	}
	if len(_dlmTags) > 0 {
		if err := assignInputField(input, "Tags", _dlmTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLifecyclePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified lifecycle policy and halts the automated operations that
// the policy specified.
//
// For more information about deleting a policy, see [Delete lifecycle policies].
//
// [Delete lifecycle policies]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/view-modify-delete.html#delete
func dlm_DeleteLifecyclePolicy(cfg aws.Config, client *dlm.Client) {
	input := &dlm.DeleteLifecyclePolicyInput{
		// PolicyId: *string, // Required
	}

	if len(_dlmPolicyId) > 0 {
		input.PolicyId = aws.String(_dlmPolicyId)
	}

	if resp, err := client.DeleteLifecyclePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets summary information about all or the specified data lifecycle policies.
// To get complete information about a policy, use [GetLifecyclePolicy].
//
// [GetLifecyclePolicy]: https://docs.aws.amazon.com/dlm/latest/APIReference/API_GetLifecyclePolicy.html
func dlm_GetLifecyclePolicies(cfg aws.Config, client *dlm.Client) {
	input := &dlm.GetLifecyclePoliciesInput{}

	if len(_dlmDefaultPolicyType) > 0 {
		if err := assignInputField(input, "DefaultPolicyType", _dlmDefaultPolicyType); err != nil {
			log.Errorf("invalid --default-policy-type: %s", err.Error())
			return
		}
	}
	if len(_dlmPolicyIds) > 0 {
		input.PolicyIds = append([]string(nil), _dlmPolicyIds...)
	}
	if len(_dlmResourceTypes) > 0 {
		if err := assignInputField(input, "ResourceTypes", _dlmResourceTypes); err != nil {
			log.Errorf("invalid --resource-types: %s", err.Error())
			return
		}
	}
	if len(_dlmState) > 0 {
		if err := assignInputField(input, "State", _dlmState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}
	if len(_dlmTagsToAdd) > 0 {
		input.TagsToAdd = append([]string(nil), _dlmTagsToAdd...)
	}
	if len(_dlmTargetTags) > 0 {
		input.TargetTags = append([]string(nil), _dlmTargetTags...)
	}

	if resp, err := client.GetLifecyclePolicies(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets detailed information about the specified lifecycle policy.
func dlm_GetLifecyclePolicy(cfg aws.Config, client *dlm.Client) {
	input := &dlm.GetLifecyclePolicyInput{
		// PolicyId: *string, // Required
	}

	if len(_dlmPolicyId) > 0 {
		input.PolicyId = aws.String(_dlmPolicyId)
	}

	if resp, err := client.GetLifecyclePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the tags for the specified resource.
func dlm_ListTagsForResource(cfg aws.Config, client *dlm.Client) {
	input := &dlm.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_dlmResourceArn) > 0 {
		input.ResourceArn = aws.String(_dlmResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds the specified tags to the specified resource.
func dlm_TagResource(cfg aws.Config, client *dlm.Client) {
	input := &dlm.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_dlmResourceArn) > 0 {
		input.ResourceArn = aws.String(_dlmResourceArn)
	}
	if len(_dlmTags) > 0 {
		if err := assignInputField(input, "Tags", _dlmTags); err != nil {
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

// Removes the specified tags from the specified resource.
func dlm_UntagResource(cfg aws.Config, client *dlm.Client) {
	input := &dlm.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_dlmResourceArn) > 0 {
		input.ResourceArn = aws.String(_dlmResourceArn)
	}
	if len(_dlmTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _dlmTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified lifecycle policy.
// For more information about updating a policy, see [Modify lifecycle policies].
//
// [Modify lifecycle policies]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/view-modify-delete.html#modify
func dlm_UpdateLifecyclePolicy(cfg aws.Config, client *dlm.Client) {
	input := &dlm.UpdateLifecyclePolicyInput{
		// PolicyId: *string, // Required
	}

	if len(_dlmPolicyId) > 0 {
		input.PolicyId = aws.String(_dlmPolicyId)
	}
	if len(_dlmCopyTags) > 0 {
		if err := assignInputField(input, "CopyTags", _dlmCopyTags); err != nil {
			log.Errorf("invalid --copy-tags: %s", err.Error())
			return
		}
	}
	if len(_dlmCreateInterval) > 0 {
		if err := assignInputField(input, "CreateInterval", _dlmCreateInterval); err != nil {
			log.Errorf("invalid --create-interval: %s", err.Error())
			return
		}
	}
	if len(_dlmCrossRegionCopyTargets) > 0 {
		if err := assignInputField(input, "CrossRegionCopyTargets", _dlmCrossRegionCopyTargets); err != nil {
			log.Errorf("invalid --cross-region-copy-targets: %s", err.Error())
			return
		}
	}
	if len(_dlmDescription) > 0 {
		input.Description = aws.String(_dlmDescription)
	}
	if len(_dlmExclusions) > 0 {
		if err := assignInputField(input, "Exclusions", _dlmExclusions); err != nil {
			log.Errorf("invalid --exclusions: %s", err.Error())
			return
		}
	}
	if len(_dlmExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_dlmExecutionRoleArn)
	}
	if len(_dlmExtendDeletion) > 0 {
		if err := assignInputField(input, "ExtendDeletion", _dlmExtendDeletion); err != nil {
			log.Errorf("invalid --extend-deletion: %s", err.Error())
			return
		}
	}
	if len(_dlmPolicyDetails) > 0 {
		if err := assignInputField(input, "PolicyDetails", _dlmPolicyDetails); err != nil {
			log.Errorf("invalid --policy-details: %s", err.Error())
			return
		}
	}
	if len(_dlmRetainInterval) > 0 {
		if err := assignInputField(input, "RetainInterval", _dlmRetainInterval); err != nil {
			log.Errorf("invalid --retain-interval: %s", err.Error())
			return
		}
	}
	if len(_dlmState) > 0 {
		if err := assignInputField(input, "State", _dlmState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateLifecyclePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_dlmCmd)
	_dlmCmd.Flags().SortFlags = false

	_dlmCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_dlmCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_dlmCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_dlmCmd.Flags().StringVarP(&_dlmCopyTags, "copy-tags", "", "", "Copy Tags")
	_dlmCmd.Flags().StringVarP(&_dlmCreateInterval, "create-interval", "", "", "Create Interval")
	_dlmCmd.Flags().StringVarP(&_dlmCrossRegionCopyTargets, "cross-region-copy-targets", "", "", "Cross Region Copy Targets")
	_dlmCmd.Flags().StringVarP(&_dlmDefaultPolicy, "default-policy", "", "", "Default Policy")
	_dlmCmd.Flags().StringVarP(&_dlmDefaultPolicyType, "default-policy-type", "", "", "Default Policy Type")
	_dlmCmd.Flags().StringVarP(&_dlmDescription, "description", "", "", "Description")
	_dlmCmd.Flags().StringVarP(&_dlmExclusions, "exclusions", "", "", "Exclusions")
	_dlmCmd.Flags().StringVarP(&_dlmExecutionRoleArn, "execution-role-arn", "", "", "Execution Role ARN")
	_dlmCmd.Flags().StringVarP(&_dlmExtendDeletion, "extend-deletion", "", "", "Extend Deletion")
	_dlmCmd.Flags().StringVarP(&_dlmPolicyDetails, "policy-details", "", "", "Policy Details")
	_dlmCmd.Flags().StringVarP(&_dlmPolicyId, "policy-id", "", "", "Policy ID")
	_dlmCmd.Flags().StringSliceVarP(&_dlmPolicyIds, "policy-ids", "", nil, "Policy Ids")
	_dlmCmd.Flags().StringVarP(&_dlmResourceArn, "resource-arn", "", "", "Resource ARN")
	_dlmCmd.Flags().StringVarP(&_dlmResourceTypes, "resource-types", "", "", "Resource Types")
	_dlmCmd.Flags().StringVarP(&_dlmRetainInterval, "retain-interval", "", "", "Retain Interval")
	_dlmCmd.Flags().StringVarP(&_dlmState, "state", "", "", "State")
	_dlmCmd.Flags().StringSliceVarP(&_dlmTagKeys, "tag-keys", "", nil, "Tag Keys")
	_dlmCmd.Flags().StringVarP(&_dlmTags, "tags", "", "", "Tags")
	_dlmCmd.Flags().StringSliceVarP(&_dlmTagsToAdd, "tags-to-add", "", nil, "Tags To Add")
	_dlmCmd.Flags().StringSliceVarP(&_dlmTargetTags, "target-tags", "", nil, "Target Tags")

	_dlmCmd.Flags().BoolVarP(&_dlmCreateLifecyclePolicy, "create-lifecycle-policy", "", false, "Create Lifecycle Policy")
	_dlmCmd.Flags().BoolVarP(&_dlmDeleteLifecyclePolicy, "delete-lifecycle-policy", "", false, "Delete Lifecycle Policy")
	_dlmCmd.Flags().BoolVarP(&_dlmGetLifecyclePolicies, "get-lifecycle-policies", "", false, "Get Lifecycle Policies")
	_dlmCmd.Flags().BoolVarP(&_dlmGetLifecyclePolicy, "get-lifecycle-policy", "", false, "Get Lifecycle Policy")
	_dlmCmd.Flags().BoolVarP(&_dlmListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_dlmCmd.Flags().BoolVarP(&_dlmTagResource, "tag-resource", "", false, "Tag Resource")
	_dlmCmd.Flags().BoolVarP(&_dlmUntagResource, "untag-resource", "", false, "Untag Resource")
	_dlmCmd.Flags().BoolVarP(&_dlmUpdateLifecyclePolicy, "update-lifecycle-policy", "", false, "Update Lifecycle Policy")

}
