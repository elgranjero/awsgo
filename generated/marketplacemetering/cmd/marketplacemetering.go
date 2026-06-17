package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/marketplacemetering"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// marketplacemeteringCmd represents the marketplacemetering command
var _marketplacemeteringCmd = &cobra.Command{
	Use:   "marketplacemetering",
	Short: "AWS marketplacemetering CLI",
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
		client := marketplacemetering.NewFromConfig(cfg)
		if _marketplacemeteringBatchMeterUsage {
			marketplacemetering_BatchMeterUsage(cfg, client)
			return
		}
		if _marketplacemeteringMeterUsage {
			marketplacemetering_MeterUsage(cfg, client)
			return
		}
		if _marketplacemeteringRegisterUsage {
			marketplacemetering_RegisterUsage(cfg, client)
			return
		}
		if _marketplacemeteringResolveCustomer {
			marketplacemetering_ResolveCustomer(cfg, client)
			return
		}

	},
}

var (
	_marketplacemeteringBatchMeterUsage bool
	_marketplacemeteringMeterUsage      bool
	_marketplacemeteringRegisterUsage   bool
	_marketplacemeteringResolveCustomer bool

	_marketplacemeteringClientToken       string
	_marketplacemeteringDryRun            string
	_marketplacemeteringNonce             string
	_marketplacemeteringProductCode       string
	_marketplacemeteringPublicKeyVersion  string
	_marketplacemeteringRegistrationToken string
	_marketplacemeteringTimestamp         string
	_marketplacemeteringUsageAllocations  string
	_marketplacemeteringUsageDimension    string
	_marketplacemeteringUsageQuantity     string
	_marketplacemeteringUsageRecords      string
)

// Amazon Web Services Marketplace is introducing Concurrent Agreements, enabling
// buyers to make multiple purchases per Amazon Web Services account. Starting June
// 1, 2026, new SaaS products must use CustomerAWSAccountId (instead of
// CustomerIdentifier ), LicenseArn (instead of ProductCode ) to support this
// feature. Existing integrations will continue to work. Review the new integration
// for Concurrent Agreements [here].
//
// To post metering records for customers, SaaS applications call BatchMeterUsage ,
// which is used for metering SaaS flexible consumption pricing (FCP). Identical
// requests are idempotent and can be retried with the same records or a subset of
// records. Each BatchMeterUsage request is for only one product. If you want to
// meter usage for multiple products, you must make multiple BatchMeterUsage calls.
//
// Usage records should be submitted in quick succession following a recorded
// event. Usage records aren't accepted 6 hours or more after an event.
//
// BatchMeterUsage can process up to 25 UsageRecords at a time, and each request
// must be less than 1 MB in size. Optionally, you can have multiple usage
// allocations for usage data that's split into buckets according to predefined
// tags.
//
// BatchMeterUsage returns a list of UsageRecordResult objects, which have each
// UsageRecord . It also returns a list of UnprocessedRecords , which indicate
// errors on the service side that should be retried.
//
// For Amazon Web Services Regions that support BatchMeterUsage , see [BatchMeterUsage Region support].
//
// For an example of BatchMeterUsage , see [BatchMeterUsage code example] in the Amazon Web Services Marketplace
// Seller Guide.
//
// [here]: https://catalog.workshops.aws/mpseller/en-US/saas/integration-for-concurrent-agreements
// [BatchMeterUsage code example]: https://docs.aws.amazon.com/marketplace/latest/userguide/saas-code-examples.html#saas-batchmeterusage-example
// [BatchMeterUsage Region support]: https://docs.aws.amazon.com/marketplace/latest/APIReference/metering-regions.html#batchmeterusage-region-support
func marketplacemetering_BatchMeterUsage(cfg aws.Config, client *marketplacemetering.Client) {
	input := &marketplacemetering.BatchMeterUsageInput{
		// UsageRecords: []types.UsageRecord, // Required
	}

	if len(_marketplacemeteringUsageRecords) > 0 {
		if err := assignInputField(input, "UsageRecords", _marketplacemeteringUsageRecords); err != nil {
			log.Errorf("invalid --usage-records: %s", err.Error())
			return
		}
	}
	if len(_marketplacemeteringProductCode) > 0 {
		input.ProductCode = aws.String(_marketplacemeteringProductCode)
	}

	if resp, err := client.BatchMeterUsage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// As a seller, your software hosted in the buyer's Amazon Web Services account
// uses this API action to emit metering records directly to Amazon Web Services
// Marketplace. You must use the following buyer Amazon Web Services account
// credentials to sign the API request.
//
// - For Amazon EC2 deployments, your software must use the [IAM role for Amazon EC2]to sign the API call
// for MeterUsage API operation.
//
// - For Amazon EKS deployments, your software must use [IAM roles for service accounts (IRSA)]to sign the API call for
// the MeterUsage API operation. Using [EKS Pod Identity], the node role, or long-term access keys
// is not supported.
//
// - For Amazon ECS deployments, your software must use [Amazon ECS task IAM]role to sign the API
// call for the MeterUsage API operation. Using the node role or long-term access
// keys are not supported.
//
// - For Amazon Bedrock AgentCore Runtime deployments, your software must use
// the [AgentCore Runtime execution role]to sign the API call for the MeterUsage API operation. Long-term access
// keys are not supported.
//
// The handling of MeterUsage requests varies between Amazon Bedrock AgentCore
// Runtime and non-Amazon Bedrock AgentCore deployments.
//
// - For non-Amazon Bedrock AgentCore Runtime deployments, you can only report
// usage once per hour for each dimension. For AMI-based products, this is per
// dimension and per EC2 instance. For container products, this is per dimension
// and per ECS task or EKS pod. You can't modify values after they're recorded. If
// you report usage before a current hour ends, you will be unable to report
// additional usage until the next hour begins. The Timestamp request parameter
// is rounded down to the hour and used to enforce this once-per-hour rule for
// idempotency. For requests that are identical after the Timestamp is rounded
// down, the API is idempotent and returns the metering record ID.
//
// - For Amazon Bedrock AgentCore Runtime deployments, you can report usage
// multiple times per hour for the same dimension. You do not need to aggregate
// metering records by the hour. You must include an idempotency token in the
// ClientToken request parameter. If using an Amazon SDK or the Amazon Web
// Services CLI, you must use the latest version which automatically includes an
// idempotency token in the ClientToken request parameter so that the request is
// processed successfully. The Timestamp request parameter is not rounded down to
// the hour and is not used for duplicate validation. Requests with duplicate
// Timestamps are aggregated as long as the ClientToken is unique.
//
// If you submit records more than six hours after events occur, the records won't
// be accepted. The timestamp in your request determines when an event is recorded.
//
// You can optionally include multiple usage allocations, to provide customers
// with usage data split into buckets by tags that you define or allow the customer
// to define.
//
// For Amazon Web Services Regions that support MeterUsage , see [MeterUsage Region support for Amazon EC2] and [MeterUsage Region support for Amazon ECS and Amazon EKS].
//
// [MeterUsage Region support for Amazon ECS and Amazon EKS]: https://docs.aws.amazon.com/marketplace/latest/APIReference/metering-regions.html#meterusage-region-support-ecs-eks
// [Amazon ECS task IAM]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html
// [AgentCore Runtime execution role]: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/runtime-permissions.html#runtime-permissions-execution
// [IAM role for Amazon EC2]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/iam-roles-for-amazon-ec2.html
// [IAM roles for service accounts (IRSA)]: https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html
// [EKS Pod Identity]: https://docs.aws.amazon.com/eks/latest/userguide/pod-identities.html
// [MeterUsage Region support for Amazon EC2]: https://docs.aws.amazon.com/marketplace/latest/APIReference/metering-regions.html#meterusage-region-support-ec2
func marketplacemetering_MeterUsage(cfg aws.Config, client *marketplacemetering.Client) {
	input := &marketplacemetering.MeterUsageInput{
		// ProductCode: *string, // Required
		// Timestamp: *time.Time, // Required
		// UsageDimension: *string, // Required
	}

	if len(_marketplacemeteringProductCode) > 0 {
		input.ProductCode = aws.String(_marketplacemeteringProductCode)
	}
	if len(_marketplacemeteringTimestamp) > 0 {
		if err := assignInputField(input, "Timestamp", _marketplacemeteringTimestamp); err != nil {
			log.Errorf("invalid --timestamp: %s", err.Error())
			return
		}
	}
	if len(_marketplacemeteringUsageDimension) > 0 {
		input.UsageDimension = aws.String(_marketplacemeteringUsageDimension)
	}
	if len(_marketplacemeteringClientToken) > 0 {
		input.ClientToken = aws.String(_marketplacemeteringClientToken)
	}
	if len(_marketplacemeteringDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _marketplacemeteringDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_marketplacemeteringUsageAllocations) > 0 {
		if err := assignInputField(input, "UsageAllocations", _marketplacemeteringUsageAllocations); err != nil {
			log.Errorf("invalid --usage-allocations: %s", err.Error())
			return
		}
	}
	if len(_marketplacemeteringUsageQuantity) > 0 {
		if err := assignInputField(input, "UsageQuantity", _marketplacemeteringUsageQuantity); err != nil {
			log.Errorf("invalid --usage-quantity: %s", err.Error())
			return
		}
	}

	if resp, err := client.MeterUsage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Paid container software products sold through Amazon Web Services Marketplace
// must integrate with the Amazon Web Services Marketplace Metering Service and
// call the RegisterUsage operation for software entitlement and metering. Free
// and BYOL products for Amazon ECS or Amazon EKS aren't required to call
// RegisterUsage , but you may choose to do so if you would like to receive usage
// data in your seller reports. The sections below explain the behavior of
// RegisterUsage . RegisterUsage performs two primary functions: metering and
// entitlement.
//
// - Entitlement: RegisterUsage allows you to verify that the customer running
// your paid software is subscribed to your product on Amazon Web Services
// Marketplace, enabling you to guard against unauthorized use. Your container
// image that integrates with RegisterUsage is only required to guard against
// unauthorized use at container startup, as such a
// CustomerNotSubscribedException or PlatformNotSupportedException will only be
// thrown on the initial call to RegisterUsage . Subsequent calls from the same
// Amazon ECS task instance (e.g. task-id) or Amazon EKS pod will not throw a
// CustomerNotSubscribedException , even if the customer unsubscribes while the
// Amazon ECS task or Amazon EKS pod is still running.
//
// - Metering: RegisterUsage meters software use per ECS task, per hour, or per
// pod for Amazon EKS with usage prorated to the second. A minimum of 1 minute of
// usage applies to tasks that are short lived. For example, if a customer has a 10
// node Amazon ECS or Amazon EKS cluster and a service configured as a Daemon Set,
// then Amazon ECS or Amazon EKS will launch a task on all 10 cluster nodes and the
// customer will be charged for 10 tasks. Software metering is handled by the
// Amazon Web Services Marketplace metering control plane—your software is not
// required to perform metering-specific actions other than to call RegisterUsage
// to commence metering. The Amazon Web Services Marketplace metering control plane
// will also bill customers for running ECS tasks and Amazon EKS pods, regardless
// of the customer's subscription state, which removes the need for your software
// to run entitlement checks at runtime. For containers, RegisterUsage should be
// called immediately at launch. If you don’t register the container within the
// first 6 hours of the launch, Amazon Web Services Marketplace Metering Service
// doesn’t provide any metering guarantees for previous months. Metering will
// continue, however, for the current month forward until the container ends.
// RegisterUsage is for metering paid hourly container products.
//
// For Amazon Web Services Regions that support RegisterUsage , see [RegisterUsage Region support].
//
// [RegisterUsage Region support]: https://docs.aws.amazon.com/marketplace/latest/APIReference/metering-regions.html#registerusage-region-support
func marketplacemetering_RegisterUsage(cfg aws.Config, client *marketplacemetering.Client) {
	input := &marketplacemetering.RegisterUsageInput{
		// ProductCode: *string, // Required
		// PublicKeyVersion: *int32, // Required
	}

	if len(_marketplacemeteringProductCode) > 0 {
		input.ProductCode = aws.String(_marketplacemeteringProductCode)
	}
	if len(_marketplacemeteringPublicKeyVersion) > 0 {
		if err := assignInputField(input, "PublicKeyVersion", _marketplacemeteringPublicKeyVersion); err != nil {
			log.Errorf("invalid --public-key-version: %s", err.Error())
			return
		}
	}
	if len(_marketplacemeteringNonce) > 0 {
		input.Nonce = aws.String(_marketplacemeteringNonce)
	}

	if resp, err := client.RegisterUsage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// ResolveCustomer is called by a SaaS application during the registration
// process. When a buyer visits your website during the registration process, the
// buyer submits a registration token through their browser. The registration token
// is resolved through this API to obtain a CustomerIdentifier along with the
// CustomerAWSAccountId , ProductCode , and LicenseArn .
//
// To successfully resolve the token, the API must be called from the account that
// was used to publish the SaaS application. For an example of using
// ResolveCustomer , see [ResolveCustomer code example] in the Amazon Web Services Marketplace Seller Guide.
//
// Permission is required for this operation. Your IAM role or user performing
// this operation requires a policy to allow the aws-marketplace:ResolveCustomer
// action. For more information, see [Actions, resources, and condition keys for Amazon Web Services Marketplace Metering Service]in the Service Authorization Reference.
//
// For Amazon Web Services Regions that support ResolveCustomer , see [ResolveCustomer Region support].
//
// [ResolveCustomer code example]: https://docs.aws.amazon.com/marketplace/latest/userguide/saas-code-examples.html#saas-resolvecustomer-example
// [Actions, resources, and condition keys for Amazon Web Services Marketplace Metering Service]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsmarketplacemeteringservice.html
// [ResolveCustomer Region support]: https://docs.aws.amazon.com/marketplace/latest/APIReference/metering-regions.html#resolvecustomer-region-support
func marketplacemetering_ResolveCustomer(cfg aws.Config, client *marketplacemetering.Client) {
	input := &marketplacemetering.ResolveCustomerInput{
		// RegistrationToken: *string, // Required
	}

	if len(_marketplacemeteringRegistrationToken) > 0 {
		input.RegistrationToken = aws.String(_marketplacemeteringRegistrationToken)
	}

	if resp, err := client.ResolveCustomer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_marketplacemeteringCmd)
	_marketplacemeteringCmd.Flags().SortFlags = false

	_marketplacemeteringCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_marketplacemeteringCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_marketplacemeteringCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_marketplacemeteringCmd.Flags().StringVarP(&_marketplacemeteringClientToken, "client-token", "", "", "Client Token")
	_marketplacemeteringCmd.Flags().StringVarP(&_marketplacemeteringDryRun, "dry-run", "", "", "Dry Run")
	_marketplacemeteringCmd.Flags().StringVarP(&_marketplacemeteringNonce, "nonce", "", "", "Nonce")
	_marketplacemeteringCmd.Flags().StringVarP(&_marketplacemeteringProductCode, "product-code", "", "", "Product Code")
	_marketplacemeteringCmd.Flags().StringVarP(&_marketplacemeteringPublicKeyVersion, "public-key-version", "", "", "Public Key Version")
	_marketplacemeteringCmd.Flags().StringVarP(&_marketplacemeteringRegistrationToken, "registration-token", "", "", "Registration Token")
	_marketplacemeteringCmd.Flags().StringVarP(&_marketplacemeteringTimestamp, "timestamp", "", "", "Timestamp")
	_marketplacemeteringCmd.Flags().StringVarP(&_marketplacemeteringUsageAllocations, "usage-allocations", "", "", "Usage Allocations")
	_marketplacemeteringCmd.Flags().StringVarP(&_marketplacemeteringUsageDimension, "usage-dimension", "", "", "Usage Dimension")
	_marketplacemeteringCmd.Flags().StringVarP(&_marketplacemeteringUsageQuantity, "usage-quantity", "", "", "Usage Quantity")
	_marketplacemeteringCmd.Flags().StringVarP(&_marketplacemeteringUsageRecords, "usage-records", "", "", "Usage Records")

	_marketplacemeteringCmd.Flags().BoolVarP(&_marketplacemeteringBatchMeterUsage, "batch-meter-usage", "", false, "Batch Meter Usage")
	_marketplacemeteringCmd.Flags().BoolVarP(&_marketplacemeteringMeterUsage, "meter-usage", "", false, "Meter Usage")
	_marketplacemeteringCmd.Flags().BoolVarP(&_marketplacemeteringRegisterUsage, "register-usage", "", false, "Register Usage")
	_marketplacemeteringCmd.Flags().BoolVarP(&_marketplacemeteringResolveCustomer, "resolve-customer", "", false, "Resolve Customer")

}
