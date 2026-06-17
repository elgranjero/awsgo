package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemakermetrics"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// sagemakermetricsCmd represents the sagemakermetrics command
var _sagemakermetricsCmd = &cobra.Command{
	Use:   "sagemakermetrics",
	Short: "AWS sagemakermetrics CLI",
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
		client := sagemakermetrics.NewFromConfig(cfg)
		if _sagemakermetricsBatchGetMetrics {
			sagemakermetrics_BatchGetMetrics(cfg, client)
			return
		}
		if _sagemakermetricsBatchPutMetrics {
			sagemakermetrics_BatchPutMetrics(cfg, client)
			return
		}

	},
}

var (
	_sagemakermetricsBatchGetMetrics bool
	_sagemakermetricsBatchPutMetrics bool

	_sagemakermetricsMetricData         string
	_sagemakermetricsMetricQueries      string
	_sagemakermetricsTrialComponentName string
)

// Used to retrieve training metrics from SageMaker.
func sagemakermetrics_BatchGetMetrics(cfg aws.Config, client *sagemakermetrics.Client) {
	input := &sagemakermetrics.BatchGetMetricsInput{
		// MetricQueries: []types.MetricQuery, // Required
	}

	if len(_sagemakermetricsMetricQueries) > 0 {
		if err := assignInputField(input, "MetricQueries", _sagemakermetricsMetricQueries); err != nil {
			log.Errorf("invalid --metric-queries: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchGetMetrics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Used to ingest training metrics into SageMaker. These metrics can be visualized
// in SageMaker Studio.
func sagemakermetrics_BatchPutMetrics(cfg aws.Config, client *sagemakermetrics.Client) {
	input := &sagemakermetrics.BatchPutMetricsInput{
		// MetricData: []types.RawMetricData, // Required
		// TrialComponentName: *string, // Required
	}

	if len(_sagemakermetricsMetricData) > 0 {
		if err := assignInputField(input, "MetricData", _sagemakermetricsMetricData); err != nil {
			log.Errorf("invalid --metric-data: %s", err.Error())
			return
		}
	}
	if len(_sagemakermetricsTrialComponentName) > 0 {
		input.TrialComponentName = aws.String(_sagemakermetricsTrialComponentName)
	}

	if resp, err := client.BatchPutMetrics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_sagemakermetricsCmd)
	_sagemakermetricsCmd.Flags().SortFlags = false

	_sagemakermetricsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_sagemakermetricsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_sagemakermetricsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_sagemakermetricsCmd.Flags().StringVarP(&_sagemakermetricsMetricData, "metric-data", "", "", "Metric Data")
	_sagemakermetricsCmd.Flags().StringVarP(&_sagemakermetricsMetricQueries, "metric-queries", "", "", "Metric Queries")
	_sagemakermetricsCmd.Flags().StringVarP(&_sagemakermetricsTrialComponentName, "trial-component-name", "", "", "Trial Component Name")

	_sagemakermetricsCmd.Flags().BoolVarP(&_sagemakermetricsBatchGetMetrics, "batch-get-metrics", "", false, "Batch Get Metrics")
	_sagemakermetricsCmd.Flags().BoolVarP(&_sagemakermetricsBatchPutMetrics, "batch-put-metrics", "", false, "Batch Put Metrics")

}
