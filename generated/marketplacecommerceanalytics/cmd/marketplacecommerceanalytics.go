package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/marketplacecommerceanalytics"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// marketplacecommerceanalyticsCmd represents the marketplacecommerceanalytics command
var _marketplacecommerceanalyticsCmd = &cobra.Command{
	Use:   "marketplacecommerceanalytics",
	Short: "AWS marketplacecommerceanalytics CLI",
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
		client := marketplacecommerceanalytics.NewFromConfig(cfg)
		if _marketplacecommerceanalyticsGenerateDataSet {
			marketplacecommerceanalytics_GenerateDataSet(cfg, client)
			return
		}
		if _marketplacecommerceanalyticsStartSupportDataExport {
			marketplacecommerceanalytics_StartSupportDataExport(cfg, client)
			return
		}

	},
}

var (
	_marketplacecommerceanalyticsGenerateDataSet        bool
	_marketplacecommerceanalyticsStartSupportDataExport bool

	_marketplacecommerceanalyticsCustomerDefinedValues   string
	_marketplacecommerceanalyticsDataSetPublicationDate  string
	_marketplacecommerceanalyticsDataSetType             string
	_marketplacecommerceanalyticsDestinationS3BucketName string
	_marketplacecommerceanalyticsDestinationS3Prefix     string
	_marketplacecommerceanalyticsFromDate                string
	_marketplacecommerceanalyticsRoleNameArn             string
	_marketplacecommerceanalyticsSnsTopicArn             string
)

// Given a data set type and data set publication date, asynchronously publishes
// the requested data set to the specified S3 bucket and notifies the specified SNS
// topic once the data is available. Returns a unique request identifier that can
// be used to correlate requests with notifications from the SNS topic. Data sets
// will be published in comma-separated values (CSV) format with the file name
// {data_set_type}_YYYY-MM-DD.csv. If a file with the same name already exists
// (e.g. if the same data set is requested twice), the original file will be
// overwritten by the new file. Requires a Role with an attached permissions policy
// providing Allow permissions for the following actions: s3:PutObject,
// s3:GetBucketLocation, sns:GetTopicAttributes, sns:Publish, iam:GetRolePolicy.
func marketplacecommerceanalytics_GenerateDataSet(cfg aws.Config, client *marketplacecommerceanalytics.Client) {
	input := &marketplacecommerceanalytics.GenerateDataSetInput{
		// DataSetPublicationDate: *time.Time, // Required
		// DataSetType: types.DataSetType, // Required
		// DestinationS3BucketName: *string, // Required
		// RoleNameArn: *string, // Required
		// SnsTopicArn: *string, // Required
	}

	if len(_marketplacecommerceanalyticsDataSetPublicationDate) > 0 {
		if err := assignInputField(input, "DataSetPublicationDate", _marketplacecommerceanalyticsDataSetPublicationDate); err != nil {
			log.Errorf("invalid --data-set-publication-date: %s", err.Error())
			return
		}
	}
	if len(_marketplacecommerceanalyticsDataSetType) > 0 {
		if err := assignInputField(input, "DataSetType", _marketplacecommerceanalyticsDataSetType); err != nil {
			log.Errorf("invalid --data-set-type: %s", err.Error())
			return
		}
	}
	if len(_marketplacecommerceanalyticsDestinationS3BucketName) > 0 {
		input.DestinationS3BucketName = aws.String(_marketplacecommerceanalyticsDestinationS3BucketName)
	}
	if len(_marketplacecommerceanalyticsRoleNameArn) > 0 {
		input.RoleNameArn = aws.String(_marketplacecommerceanalyticsRoleNameArn)
	}
	if len(_marketplacecommerceanalyticsSnsTopicArn) > 0 {
		input.SnsTopicArn = aws.String(_marketplacecommerceanalyticsSnsTopicArn)
	}
	if len(_marketplacecommerceanalyticsCustomerDefinedValues) > 0 {
		if err := assignInputField(input, "CustomerDefinedValues", _marketplacecommerceanalyticsCustomerDefinedValues); err != nil {
			log.Errorf("invalid --customer-defined-values: %s", err.Error())
			return
		}
	}
	if len(_marketplacecommerceanalyticsDestinationS3Prefix) > 0 {
		input.DestinationS3Prefix = aws.String(_marketplacecommerceanalyticsDestinationS3Prefix)
	}

	if resp, err := client.GenerateDataSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This target has been deprecated. Given a data set type and a from date,
// asynchronously publishes the requested customer support data to the specified S3
// bucket and notifies the specified SNS topic once the data is available. Returns
// a unique request identifier that can be used to correlate requests with
// notifications from the SNS topic. Data sets will be published in comma-separated
// values (CSV) format with the file name
// {data_set_type}_YYYY-MM-DD'T'HH-mm-ss'Z'.csv. If a file with the same name
// already exists (e.g. if the same data set is requested twice), the original file
// will be overwritten by the new file. Requires a Role with an attached
// permissions policy providing Allow permissions for the following actions:
// s3:PutObject, s3:GetBucketLocation, sns:GetTopicAttributes, sns:Publish,
// iam:GetRolePolicy.
//
// Deprecated: This target has been deprecated. As of December 2022 Product
// Support Connection is no longer supported.
func marketplacecommerceanalytics_StartSupportDataExport(cfg aws.Config, client *marketplacecommerceanalytics.Client) {
	input := &marketplacecommerceanalytics.StartSupportDataExportInput{
		// DataSetType: types.SupportDataSetType, // Required
		// DestinationS3BucketName: *string, // Required
		// FromDate: *time.Time, // Required
		// RoleNameArn: *string, // Required
		// SnsTopicArn: *string, // Required
	}

	if len(_marketplacecommerceanalyticsDataSetType) > 0 {
		if err := assignInputField(input, "DataSetType", _marketplacecommerceanalyticsDataSetType); err != nil {
			log.Errorf("invalid --data-set-type: %s", err.Error())
			return
		}
	}
	if len(_marketplacecommerceanalyticsDestinationS3BucketName) > 0 {
		input.DestinationS3BucketName = aws.String(_marketplacecommerceanalyticsDestinationS3BucketName)
	}
	if len(_marketplacecommerceanalyticsFromDate) > 0 {
		if err := assignInputField(input, "FromDate", _marketplacecommerceanalyticsFromDate); err != nil {
			log.Errorf("invalid --from-date: %s", err.Error())
			return
		}
	}
	if len(_marketplacecommerceanalyticsRoleNameArn) > 0 {
		input.RoleNameArn = aws.String(_marketplacecommerceanalyticsRoleNameArn)
	}
	if len(_marketplacecommerceanalyticsSnsTopicArn) > 0 {
		input.SnsTopicArn = aws.String(_marketplacecommerceanalyticsSnsTopicArn)
	}
	if len(_marketplacecommerceanalyticsCustomerDefinedValues) > 0 {
		if err := assignInputField(input, "CustomerDefinedValues", _marketplacecommerceanalyticsCustomerDefinedValues); err != nil {
			log.Errorf("invalid --customer-defined-values: %s", err.Error())
			return
		}
	}
	if len(_marketplacecommerceanalyticsDestinationS3Prefix) > 0 {
		input.DestinationS3Prefix = aws.String(_marketplacecommerceanalyticsDestinationS3Prefix)
	}

	if resp, err := client.StartSupportDataExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_marketplacecommerceanalyticsCmd)
	_marketplacecommerceanalyticsCmd.Flags().SortFlags = false

	_marketplacecommerceanalyticsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_marketplacecommerceanalyticsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_marketplacecommerceanalyticsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_marketplacecommerceanalyticsCmd.Flags().StringVarP(&_marketplacecommerceanalyticsCustomerDefinedValues, "customer-defined-values", "", "", "Customer Defined Values")
	_marketplacecommerceanalyticsCmd.Flags().StringVarP(&_marketplacecommerceanalyticsDataSetPublicationDate, "data-set-publication-date", "", "", "Data Set Publication Date")
	_marketplacecommerceanalyticsCmd.Flags().StringVarP(&_marketplacecommerceanalyticsDataSetType, "data-set-type", "", "", "Data Set Type")
	_marketplacecommerceanalyticsCmd.Flags().StringVarP(&_marketplacecommerceanalyticsDestinationS3BucketName, "destination-s3-bucket-name", "", "", "Destination S3 Bucket Name")
	_marketplacecommerceanalyticsCmd.Flags().StringVarP(&_marketplacecommerceanalyticsDestinationS3Prefix, "destination-s3-prefix", "", "", "Destination S3 Prefix")
	_marketplacecommerceanalyticsCmd.Flags().StringVarP(&_marketplacecommerceanalyticsFromDate, "from-date", "", "", "From Date")
	_marketplacecommerceanalyticsCmd.Flags().StringVarP(&_marketplacecommerceanalyticsRoleNameArn, "role-name-arn", "", "", "Role Name ARN")
	_marketplacecommerceanalyticsCmd.Flags().StringVarP(&_marketplacecommerceanalyticsSnsTopicArn, "sns-topic-arn", "", "", "SNS Topic ARN")

	_marketplacecommerceanalyticsCmd.Flags().BoolVarP(&_marketplacecommerceanalyticsGenerateDataSet, "generate-data-set", "", false, "Generate Data Set")
	_marketplacecommerceanalyticsCmd.Flags().BoolVarP(&_marketplacecommerceanalyticsStartSupportDataExport, "start-support-data-export", "", false, "Start Support Data Export")

}
