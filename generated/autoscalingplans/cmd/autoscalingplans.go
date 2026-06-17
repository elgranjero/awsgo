package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscalingplans"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// autoscalingplansCmd represents the autoscalingplans command
var _autoscalingplansCmd = &cobra.Command{
	Use:   "autoscalingplans",
	Short: "AWS autoscalingplans CLI",
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
		client := autoscalingplans.NewFromConfig(cfg)
		if _autoscalingplansCreateScalingPlan {
			autoscalingplans_CreateScalingPlan(cfg, client)
			return
		}
		if _autoscalingplansDeleteScalingPlan {
			autoscalingplans_DeleteScalingPlan(cfg, client)
			return
		}
		if _autoscalingplansDescribeScalingPlanResources {
			autoscalingplans_DescribeScalingPlanResources(cfg, client)
			return
		}
		if _autoscalingplansDescribeScalingPlans {
			autoscalingplans_DescribeScalingPlans(cfg, client)
			return
		}
		if _autoscalingplansGetScalingPlanResourceForecastData {
			autoscalingplans_GetScalingPlanResourceForecastData(cfg, client)
			return
		}
		if _autoscalingplansUpdateScalingPlan {
			autoscalingplans_UpdateScalingPlan(cfg, client)
			return
		}

	},
}

var (
	_autoscalingplansCreateScalingPlan                  bool
	_autoscalingplansDeleteScalingPlan                  bool
	_autoscalingplansDescribeScalingPlanResources       bool
	_autoscalingplansDescribeScalingPlans               bool
	_autoscalingplansGetScalingPlanResourceForecastData bool
	_autoscalingplansUpdateScalingPlan                  bool

	_autoscalingplansApplicationSource   string
	_autoscalingplansApplicationSources  string
	_autoscalingplansEndTime             string
	_autoscalingplansForecastDataType    string
	_autoscalingplansMaxResults          string
	_autoscalingplansNextToken           string
	_autoscalingplansResourceId          string
	_autoscalingplansScalableDimension   string
	_autoscalingplansScalingInstructions string
	_autoscalingplansScalingPlanName     string
	_autoscalingplansScalingPlanNames    []string
	_autoscalingplansScalingPlanVersion  string
	_autoscalingplansServiceNamespace    string
	_autoscalingplansStartTime           string
)

// Creates a scaling plan.
func autoscalingplans_CreateScalingPlan(cfg aws.Config, client *autoscalingplans.Client) {
	input := &autoscalingplans.CreateScalingPlanInput{
		// ApplicationSource: *types.ApplicationSource, // Required
		// ScalingInstructions: []types.ScalingInstruction, // Required
		// ScalingPlanName: *string, // Required
	}

	if len(_autoscalingplansApplicationSource) > 0 {
		if err := assignInputField(input, "ApplicationSource", _autoscalingplansApplicationSource); err != nil {
			log.Errorf("invalid --application-source: %s", err.Error())
			return
		}
	}
	if len(_autoscalingplansScalingInstructions) > 0 {
		if err := assignInputField(input, "ScalingInstructions", _autoscalingplansScalingInstructions); err != nil {
			log.Errorf("invalid --scaling-instructions: %s", err.Error())
			return
		}
	}
	if len(_autoscalingplansScalingPlanName) > 0 {
		input.ScalingPlanName = aws.String(_autoscalingplansScalingPlanName)
	}

	if resp, err := client.CreateScalingPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified scaling plan.
// Deleting a scaling plan deletes the underlying ScalingInstruction for all of the scalable
// resources that are covered by the plan.
//
// If the plan has launched resources or has scaling activities in progress, you
// must delete those resources separately.
func autoscalingplans_DeleteScalingPlan(cfg aws.Config, client *autoscalingplans.Client) {
	input := &autoscalingplans.DeleteScalingPlanInput{
		// ScalingPlanName: *string, // Required
		// ScalingPlanVersion: *int64, // Required
	}

	if len(_autoscalingplansScalingPlanName) > 0 {
		input.ScalingPlanName = aws.String(_autoscalingplansScalingPlanName)
	}
	if len(_autoscalingplansScalingPlanVersion) > 0 {
		if err := assignInputField(input, "ScalingPlanVersion", _autoscalingplansScalingPlanVersion); err != nil {
			log.Errorf("invalid --scaling-plan-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteScalingPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the scalable resources in the specified scaling plan.
func autoscalingplans_DescribeScalingPlanResources(cfg aws.Config, client *autoscalingplans.Client) {
	input := &autoscalingplans.DescribeScalingPlanResourcesInput{
		// ScalingPlanName: *string, // Required
		// ScalingPlanVersion: *int64, // Required
	}

	if len(_autoscalingplansScalingPlanName) > 0 {
		input.ScalingPlanName = aws.String(_autoscalingplansScalingPlanName)
	}
	if len(_autoscalingplansScalingPlanVersion) > 0 {
		if err := assignInputField(input, "ScalingPlanVersion", _autoscalingplansScalingPlanVersion); err != nil {
			log.Errorf("invalid --scaling-plan-version: %s", err.Error())
			return
		}
	}
	if len(_autoscalingplansMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _autoscalingplansMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_autoscalingplansNextToken) > 0 {
		input.NextToken = aws.String(_autoscalingplansNextToken)
	}

	if resp, err := client.DescribeScalingPlanResources(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes one or more of your scaling plans.
func autoscalingplans_DescribeScalingPlans(cfg aws.Config, client *autoscalingplans.Client) {
	input := &autoscalingplans.DescribeScalingPlansInput{}

	if len(_autoscalingplansApplicationSources) > 0 {
		if err := assignInputField(input, "ApplicationSources", _autoscalingplansApplicationSources); err != nil {
			log.Errorf("invalid --application-sources: %s", err.Error())
			return
		}
	}
	if len(_autoscalingplansMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _autoscalingplansMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_autoscalingplansNextToken) > 0 {
		input.NextToken = aws.String(_autoscalingplansNextToken)
	}
	if len(_autoscalingplansScalingPlanNames) > 0 {
		input.ScalingPlanNames = append([]string(nil), _autoscalingplansScalingPlanNames...)
	}
	if len(_autoscalingplansScalingPlanVersion) > 0 {
		if err := assignInputField(input, "ScalingPlanVersion", _autoscalingplansScalingPlanVersion); err != nil {
			log.Errorf("invalid --scaling-plan-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeScalingPlans(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the forecast data for a scalable resource.
// Capacity forecasts are represented as predicted values, or data points, that
// are calculated using historical data points from a specified CloudWatch load
// metric. Data points are available for up to 56 days.
func autoscalingplans_GetScalingPlanResourceForecastData(cfg aws.Config, client *autoscalingplans.Client) {
	input := &autoscalingplans.GetScalingPlanResourceForecastDataInput{
		// EndTime: *time.Time, // Required
		// ForecastDataType: types.ForecastDataType, // Required
		// ResourceId: *string, // Required
		// ScalableDimension: types.ScalableDimension, // Required
		// ScalingPlanName: *string, // Required
		// ScalingPlanVersion: *int64, // Required
		// ServiceNamespace: types.ServiceNamespace, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_autoscalingplansEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _autoscalingplansEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_autoscalingplansForecastDataType) > 0 {
		if err := assignInputField(input, "ForecastDataType", _autoscalingplansForecastDataType); err != nil {
			log.Errorf("invalid --forecast-data-type: %s", err.Error())
			return
		}
	}
	if len(_autoscalingplansResourceId) > 0 {
		input.ResourceId = aws.String(_autoscalingplansResourceId)
	}
	if len(_autoscalingplansScalableDimension) > 0 {
		if err := assignInputField(input, "ScalableDimension", _autoscalingplansScalableDimension); err != nil {
			log.Errorf("invalid --scalable-dimension: %s", err.Error())
			return
		}
	}
	if len(_autoscalingplansScalingPlanName) > 0 {
		input.ScalingPlanName = aws.String(_autoscalingplansScalingPlanName)
	}
	if len(_autoscalingplansScalingPlanVersion) > 0 {
		if err := assignInputField(input, "ScalingPlanVersion", _autoscalingplansScalingPlanVersion); err != nil {
			log.Errorf("invalid --scaling-plan-version: %s", err.Error())
			return
		}
	}
	if len(_autoscalingplansServiceNamespace) > 0 {
		if err := assignInputField(input, "ServiceNamespace", _autoscalingplansServiceNamespace); err != nil {
			log.Errorf("invalid --service-namespace: %s", err.Error())
			return
		}
	}
	if len(_autoscalingplansStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _autoscalingplansStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetScalingPlanResourceForecastData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified scaling plan.
// You cannot update a scaling plan if it is in the process of being created,
// updated, or deleted.
func autoscalingplans_UpdateScalingPlan(cfg aws.Config, client *autoscalingplans.Client) {
	input := &autoscalingplans.UpdateScalingPlanInput{
		// ScalingPlanName: *string, // Required
		// ScalingPlanVersion: *int64, // Required
	}

	if len(_autoscalingplansScalingPlanName) > 0 {
		input.ScalingPlanName = aws.String(_autoscalingplansScalingPlanName)
	}
	if len(_autoscalingplansScalingPlanVersion) > 0 {
		if err := assignInputField(input, "ScalingPlanVersion", _autoscalingplansScalingPlanVersion); err != nil {
			log.Errorf("invalid --scaling-plan-version: %s", err.Error())
			return
		}
	}
	if len(_autoscalingplansApplicationSource) > 0 {
		if err := assignInputField(input, "ApplicationSource", _autoscalingplansApplicationSource); err != nil {
			log.Errorf("invalid --application-source: %s", err.Error())
			return
		}
	}
	if len(_autoscalingplansScalingInstructions) > 0 {
		if err := assignInputField(input, "ScalingInstructions", _autoscalingplansScalingInstructions); err != nil {
			log.Errorf("invalid --scaling-instructions: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateScalingPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_autoscalingplansCmd)
	_autoscalingplansCmd.Flags().SortFlags = false

	_autoscalingplansCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_autoscalingplansCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_autoscalingplansCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_autoscalingplansCmd.Flags().StringVarP(&_autoscalingplansApplicationSource, "application-source", "", "", "Application Source")
	_autoscalingplansCmd.Flags().StringVarP(&_autoscalingplansApplicationSources, "application-sources", "", "", "Application Sources")
	_autoscalingplansCmd.Flags().StringVarP(&_autoscalingplansEndTime, "end-time", "", "", "End Time")
	_autoscalingplansCmd.Flags().StringVarP(&_autoscalingplansForecastDataType, "forecast-data-type", "", "", "Forecast Data Type")
	_autoscalingplansCmd.Flags().StringVarP(&_autoscalingplansMaxResults, "max-results", "", "", "Max Results")
	_autoscalingplansCmd.Flags().StringVarP(&_autoscalingplansNextToken, "next-token", "", "", "Next Token")
	_autoscalingplansCmd.Flags().StringVarP(&_autoscalingplansResourceId, "resource-id", "", "", "Resource ID")
	_autoscalingplansCmd.Flags().StringVarP(&_autoscalingplansScalableDimension, "scalable-dimension", "", "", "Scalable Dimension")
	_autoscalingplansCmd.Flags().StringVarP(&_autoscalingplansScalingInstructions, "scaling-instructions", "", "", "Scaling Instructions")
	_autoscalingplansCmd.Flags().StringVarP(&_autoscalingplansScalingPlanName, "scaling-plan-name", "", "", "Scaling Plan Name")
	_autoscalingplansCmd.Flags().StringSliceVarP(&_autoscalingplansScalingPlanNames, "scaling-plan-names", "", nil, "Scaling Plan Names")
	_autoscalingplansCmd.Flags().StringVarP(&_autoscalingplansScalingPlanVersion, "scaling-plan-version", "", "", "Scaling Plan Version")
	_autoscalingplansCmd.Flags().StringVarP(&_autoscalingplansServiceNamespace, "service-namespace", "", "", "Service Namespace")
	_autoscalingplansCmd.Flags().StringVarP(&_autoscalingplansStartTime, "start-time", "", "", "Start Time")

	_autoscalingplansCmd.Flags().BoolVarP(&_autoscalingplansCreateScalingPlan, "create-scaling-plan", "", false, "Create Scaling Plan")
	_autoscalingplansCmd.Flags().BoolVarP(&_autoscalingplansDeleteScalingPlan, "delete-scaling-plan", "", false, "Delete Scaling Plan")
	_autoscalingplansCmd.Flags().BoolVarP(&_autoscalingplansDescribeScalingPlanResources, "describe-scaling-plan-resources", "", false, "Describe Scaling Plan Resources")
	_autoscalingplansCmd.Flags().BoolVarP(&_autoscalingplansDescribeScalingPlans, "describe-scaling-plans", "", false, "Describe Scaling Plans")
	_autoscalingplansCmd.Flags().BoolVarP(&_autoscalingplansGetScalingPlanResourceForecastData, "get-scaling-plan-resource-forecast-data", "", false, "Get Scaling Plan Resource Forecast Data")
	_autoscalingplansCmd.Flags().BoolVarP(&_autoscalingplansUpdateScalingPlan, "update-scaling-plan", "", false, "Update Scaling Plan")

}
