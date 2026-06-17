package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/snowball"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// snowballCmd represents the snowball command
var _snowballCmd = &cobra.Command{
	Use:   "snowball",
	Short: "AWS snowball CLI",
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
		client := snowball.NewFromConfig(cfg)
		if _snowballCancelCluster {
			snowball_CancelCluster(cfg, client)
			return
		}
		if _snowballCancelJob {
			snowball_CancelJob(cfg, client)
			return
		}
		if _snowballCreateAddress {
			snowball_CreateAddress(cfg, client)
			return
		}
		if _snowballCreateCluster {
			snowball_CreateCluster(cfg, client)
			return
		}
		if _snowballCreateJob {
			snowball_CreateJob(cfg, client)
			return
		}
		if _snowballCreateLongTermPricing {
			snowball_CreateLongTermPricing(cfg, client)
			return
		}
		if _snowballCreateReturnShippingLabel {
			snowball_CreateReturnShippingLabel(cfg, client)
			return
		}
		if _snowballDescribeAddress {
			snowball_DescribeAddress(cfg, client)
			return
		}
		if _snowballDescribeAddresses {
			snowball_DescribeAddresses(cfg, client)
			return
		}
		if _snowballDescribeCluster {
			snowball_DescribeCluster(cfg, client)
			return
		}
		if _snowballDescribeJob {
			snowball_DescribeJob(cfg, client)
			return
		}
		if _snowballDescribeReturnShippingLabel {
			snowball_DescribeReturnShippingLabel(cfg, client)
			return
		}
		if _snowballGetJobManifest {
			snowball_GetJobManifest(cfg, client)
			return
		}
		if _snowballGetJobUnlockCode {
			snowball_GetJobUnlockCode(cfg, client)
			return
		}
		if _snowballGetSnowballUsage {
			snowball_GetSnowballUsage(cfg, client)
			return
		}
		if _snowballGetSoftwareUpdates {
			snowball_GetSoftwareUpdates(cfg, client)
			return
		}
		if _snowballListClusterJobs {
			snowball_ListClusterJobs(cfg, client)
			return
		}
		if _snowballListClusters {
			snowball_ListClusters(cfg, client)
			return
		}
		if _snowballListCompatibleImages {
			snowball_ListCompatibleImages(cfg, client)
			return
		}
		if _snowballListJobs {
			snowball_ListJobs(cfg, client)
			return
		}
		if _snowballListLongTermPricing {
			snowball_ListLongTermPricing(cfg, client)
			return
		}
		if _snowballListPickupLocations {
			snowball_ListPickupLocations(cfg, client)
			return
		}
		if _snowballListServiceVersions {
			snowball_ListServiceVersions(cfg, client)
			return
		}
		if _snowballUpdateCluster {
			snowball_UpdateCluster(cfg, client)
			return
		}
		if _snowballUpdateJob {
			snowball_UpdateJob(cfg, client)
			return
		}
		if _snowballUpdateJobShipmentState {
			snowball_UpdateJobShipmentState(cfg, client)
			return
		}
		if _snowballUpdateLongTermPricing {
			snowball_UpdateLongTermPricing(cfg, client)
			return
		}

	},
}

var (
	_snowballCancelCluster               bool
	_snowballCancelJob                   bool
	_snowballCreateAddress               bool
	_snowballCreateCluster               bool
	_snowballCreateJob                   bool
	_snowballCreateLongTermPricing       bool
	_snowballCreateReturnShippingLabel   bool
	_snowballDescribeAddress             bool
	_snowballDescribeAddresses           bool
	_snowballDescribeCluster             bool
	_snowballDescribeJob                 bool
	_snowballDescribeReturnShippingLabel bool
	_snowballGetJobManifest              bool
	_snowballGetJobUnlockCode            bool
	_snowballGetSnowballUsage            bool
	_snowballGetSoftwareUpdates          bool
	_snowballListClusterJobs             bool
	_snowballListClusters                bool
	_snowballListCompatibleImages        bool
	_snowballListJobs                    bool
	_snowballListLongTermPricing         bool
	_snowballListPickupLocations         bool
	_snowballListServiceVersions         bool
	_snowballUpdateCluster               bool
	_snowballUpdateJob                   bool
	_snowballUpdateJobShipmentState      bool
	_snowballUpdateLongTermPricing       bool

	_snowballAddress                      string
	_snowballAddressId                    string
	_snowballClusterId                    string
	_snowballDependentServices            string
	_snowballDescription                  string
	_snowballDeviceConfiguration          string
	_snowballForceCreateJobs              string
	_snowballForwardingAddressId          string
	_snowballImpactLevel                  string
	_snowballInitialClusterSize           string
	_snowballIsLongTermPricingAutoRenew   string
	_snowballJobId                        string
	_snowballJobType                      string
	_snowballKmsKeyARN                    string
	_snowballLongTermPricingId            string
	_snowballLongTermPricingIds           []string
	_snowballLongTermPricingType          string
	_snowballMaxResults                   string
	_snowballNextToken                    string
	_snowballNotification                 string
	_snowballOnDeviceServiceConfiguration string
	_snowballPickupDetails                string
	_snowballRemoteManagement             string
	_snowballReplacementJob               string
	_snowballResources                    string
	_snowballRoleARN                      string
	_snowballServiceName                  string
	_snowballShipmentState                string
	_snowballShippingOption               string
	_snowballSnowballCapacityPreference   string
	_snowballSnowballType                 string
	_snowballTaxDocuments                 string
)

// Cancels a cluster job. You can only cancel a cluster job while it's in the
// AwaitingQuorum status. You'll have at least an hour after creating a cluster job
// to cancel it.
func snowball_CancelCluster(cfg aws.Config, client *snowball.Client) {
	input := &snowball.CancelClusterInput{
		// ClusterId: *string, // Required
	}

	if len(_snowballClusterId) > 0 {
		input.ClusterId = aws.String(_snowballClusterId)
	}

	if resp, err := client.CancelCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels the specified job. You can only cancel a job before its JobState value
// changes to PreparingAppliance . Requesting the ListJobs or DescribeJob action
// returns a job's JobState as part of the response element data returned.
func snowball_CancelJob(cfg aws.Config, client *snowball.Client) {
	input := &snowball.CancelJobInput{
		// JobId: *string, // Required
	}

	if len(_snowballJobId) > 0 {
		input.JobId = aws.String(_snowballJobId)
	}

	if resp, err := client.CancelJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an address for a Snow device to be shipped to. In most regions,
// addresses are validated at the time of creation. The address you provide must be
// located within the serviceable area of your region. If the address is invalid or
// unsupported, then an exception is thrown. If providing an address as a JSON file
// through the cli-input-json option, include the full file path. For example,
// --cli-input-json file://create-address.json .
func snowball_CreateAddress(cfg aws.Config, client *snowball.Client) {
	input := &snowball.CreateAddressInput{
		// Address: *types.Address, // Required
	}

	if len(_snowballAddress) > 0 {
		if err := assignInputField(input, "Address", _snowballAddress); err != nil {
			log.Errorf("invalid --address: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAddress(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an empty cluster. Each cluster supports five nodes. You use the CreateJob action
// separately to create the jobs for each of these nodes. The cluster does not ship
// until these five node jobs have been created.
func snowball_CreateCluster(cfg aws.Config, client *snowball.Client) {
	input := &snowball.CreateClusterInput{
		// AddressId: *string, // Required
		// JobType: types.JobType, // Required
		// ShippingOption: types.ShippingOption, // Required
		// SnowballType: types.SnowballType, // Required
	}

	if len(_snowballAddressId) > 0 {
		input.AddressId = aws.String(_snowballAddressId)
	}
	if len(_snowballJobType) > 0 {
		if err := assignInputField(input, "JobType", _snowballJobType); err != nil {
			log.Errorf("invalid --job-type: %s", err.Error())
			return
		}
	}
	if len(_snowballShippingOption) > 0 {
		if err := assignInputField(input, "ShippingOption", _snowballShippingOption); err != nil {
			log.Errorf("invalid --shipping-option: %s", err.Error())
			return
		}
	}
	if len(_snowballSnowballType) > 0 {
		if err := assignInputField(input, "SnowballType", _snowballSnowballType); err != nil {
			log.Errorf("invalid --snowball-type: %s", err.Error())
			return
		}
	}
	if len(_snowballDescription) > 0 {
		input.Description = aws.String(_snowballDescription)
	}
	if len(_snowballForceCreateJobs) > 0 {
		if err := assignInputField(input, "ForceCreateJobs", _snowballForceCreateJobs); err != nil {
			log.Errorf("invalid --force-create-jobs: %s", err.Error())
			return
		}
	}
	if len(_snowballForwardingAddressId) > 0 {
		input.ForwardingAddressId = aws.String(_snowballForwardingAddressId)
	}
	if len(_snowballInitialClusterSize) > 0 {
		if err := assignInputField(input, "InitialClusterSize", _snowballInitialClusterSize); err != nil {
			log.Errorf("invalid --initial-cluster-size: %s", err.Error())
			return
		}
	}
	if len(_snowballKmsKeyARN) > 0 {
		input.KmsKeyARN = aws.String(_snowballKmsKeyARN)
	}
	if len(_snowballLongTermPricingIds) > 0 {
		input.LongTermPricingIds = append([]string(nil), _snowballLongTermPricingIds...)
	}
	if len(_snowballNotification) > 0 {
		if err := assignInputField(input, "Notification", _snowballNotification); err != nil {
			log.Errorf("invalid --notification: %s", err.Error())
			return
		}
	}
	if len(_snowballOnDeviceServiceConfiguration) > 0 {
		if err := assignInputField(input, "OnDeviceServiceConfiguration", _snowballOnDeviceServiceConfiguration); err != nil {
			log.Errorf("invalid --on-device-service-configuration: %s", err.Error())
			return
		}
	}
	if len(_snowballRemoteManagement) > 0 {
		if err := assignInputField(input, "RemoteManagement", _snowballRemoteManagement); err != nil {
			log.Errorf("invalid --remote-management: %s", err.Error())
			return
		}
	}
	if len(_snowballResources) > 0 {
		if err := assignInputField(input, "Resources", _snowballResources); err != nil {
			log.Errorf("invalid --resources: %s", err.Error())
			return
		}
	}
	if len(_snowballRoleARN) > 0 {
		input.RoleARN = aws.String(_snowballRoleARN)
	}
	if len(_snowballSnowballCapacityPreference) > 0 {
		if err := assignInputField(input, "SnowballCapacityPreference", _snowballSnowballCapacityPreference); err != nil {
			log.Errorf("invalid --snowball-capacity-preference: %s", err.Error())
			return
		}
	}
	if len(_snowballTaxDocuments) > 0 {
		if err := assignInputField(input, "TaxDocuments", _snowballTaxDocuments); err != nil {
			log.Errorf("invalid --tax-documents: %s", err.Error())
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

// Creates a job to import or export data between Amazon S3 and your on-premises
// data center. Your Amazon Web Services account must have the right trust policies
// and permissions in place to create a job for a Snow device. If you're creating a
// job for a node in a cluster, you only need to provide the clusterId value; the
// other job attributes are inherited from the cluster.
//
// Only the Snowball; Edge device type is supported when ordering clustered jobs.
//
// The device capacity is optional.
//
// Availability of device types differ by Amazon Web Services Region. For more
// information about Region availability, see [Amazon Web Services Regional Services].
//
// Snow Family devices and their capacities.
//
// - Device type: SNC1_SSD
//
// - Capacity: T14
//
// - Description: Snowcone
//
// - Device type: SNC1_HDD
//
// - Capacity: T8
//
// - Description: Snowcone
//
// - Device type: EDGE_S
//
// - Capacity: T98
//
// - Description: Snowball Edge Storage Optimized for data transfer only
//
// - Device type: EDGE_CG
//
// - Capacity: T42
//
// - Description: Snowball Edge Compute Optimized with GPU
//
// - Device type: EDGE_C
//
// - Capacity: T42
//
// - Description: Snowball Edge Compute Optimized without GPU
//
// - Device type: EDGE
//
// - Capacity: T100
//
// - Description: Snowball Edge Storage Optimized with EC2 Compute
//
// This device is replaced with T98.
//
// - Device type: STANDARD
//
// - Capacity: T50
//
// - Description: Original Snowball device
//
// # This device is only available in the Ningxia, Beijing, and Singapore Amazon Web
//
// # Services Region
//
// - Device type: STANDARD
//
// - Capacity: T80
//
// - Description: Original Snowball device
//
// # This device is only available in the Ningxia, Beijing, and Singapore Amazon Web
//
// Services Region.
//
// - Snow Family device type: RACK_5U_C
//
// - Capacity: T13
//
// - Description: Snowblade.
//
// - Device type: V3_5S
//
// - Capacity: T240
//
// - Description: Snowball Edge Storage Optimized 210TB
//
// [Amazon Web Services Regional Services]: https://aws.amazon.com/about-aws/global-infrastructure/regional-product-services/?p=ngi&loc=4
func snowball_CreateJob(cfg aws.Config, client *snowball.Client) {
	input := &snowball.CreateJobInput{}

	if len(_snowballAddressId) > 0 {
		input.AddressId = aws.String(_snowballAddressId)
	}
	if len(_snowballClusterId) > 0 {
		input.ClusterId = aws.String(_snowballClusterId)
	}
	if len(_snowballDescription) > 0 {
		input.Description = aws.String(_snowballDescription)
	}
	if len(_snowballDeviceConfiguration) > 0 {
		if err := assignInputField(input, "DeviceConfiguration", _snowballDeviceConfiguration); err != nil {
			log.Errorf("invalid --device-configuration: %s", err.Error())
			return
		}
	}
	if len(_snowballForwardingAddressId) > 0 {
		input.ForwardingAddressId = aws.String(_snowballForwardingAddressId)
	}
	if len(_snowballImpactLevel) > 0 {
		if err := assignInputField(input, "ImpactLevel", _snowballImpactLevel); err != nil {
			log.Errorf("invalid --impact-level: %s", err.Error())
			return
		}
	}
	if len(_snowballJobType) > 0 {
		if err := assignInputField(input, "JobType", _snowballJobType); err != nil {
			log.Errorf("invalid --job-type: %s", err.Error())
			return
		}
	}
	if len(_snowballKmsKeyARN) > 0 {
		input.KmsKeyARN = aws.String(_snowballKmsKeyARN)
	}
	if len(_snowballLongTermPricingId) > 0 {
		input.LongTermPricingId = aws.String(_snowballLongTermPricingId)
	}
	if len(_snowballNotification) > 0 {
		if err := assignInputField(input, "Notification", _snowballNotification); err != nil {
			log.Errorf("invalid --notification: %s", err.Error())
			return
		}
	}
	if len(_snowballOnDeviceServiceConfiguration) > 0 {
		if err := assignInputField(input, "OnDeviceServiceConfiguration", _snowballOnDeviceServiceConfiguration); err != nil {
			log.Errorf("invalid --on-device-service-configuration: %s", err.Error())
			return
		}
	}
	if len(_snowballPickupDetails) > 0 {
		if err := assignInputField(input, "PickupDetails", _snowballPickupDetails); err != nil {
			log.Errorf("invalid --pickup-details: %s", err.Error())
			return
		}
	}
	if len(_snowballRemoteManagement) > 0 {
		if err := assignInputField(input, "RemoteManagement", _snowballRemoteManagement); err != nil {
			log.Errorf("invalid --remote-management: %s", err.Error())
			return
		}
	}
	if len(_snowballResources) > 0 {
		if err := assignInputField(input, "Resources", _snowballResources); err != nil {
			log.Errorf("invalid --resources: %s", err.Error())
			return
		}
	}
	if len(_snowballRoleARN) > 0 {
		input.RoleARN = aws.String(_snowballRoleARN)
	}
	if len(_snowballShippingOption) > 0 {
		if err := assignInputField(input, "ShippingOption", _snowballShippingOption); err != nil {
			log.Errorf("invalid --shipping-option: %s", err.Error())
			return
		}
	}
	if len(_snowballSnowballCapacityPreference) > 0 {
		if err := assignInputField(input, "SnowballCapacityPreference", _snowballSnowballCapacityPreference); err != nil {
			log.Errorf("invalid --snowball-capacity-preference: %s", err.Error())
			return
		}
	}
	if len(_snowballSnowballType) > 0 {
		if err := assignInputField(input, "SnowballType", _snowballSnowballType); err != nil {
			log.Errorf("invalid --snowball-type: %s", err.Error())
			return
		}
	}
	if len(_snowballTaxDocuments) > 0 {
		if err := assignInputField(input, "TaxDocuments", _snowballTaxDocuments); err != nil {
			log.Errorf("invalid --tax-documents: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a job with the long-term usage option for a device. The long-term usage
// is a 1-year or 3-year long-term pricing type for the device. You are billed
// upfront, and Amazon Web Services provides discounts for long-term pricing.
func snowball_CreateLongTermPricing(cfg aws.Config, client *snowball.Client) {
	input := &snowball.CreateLongTermPricingInput{
		// LongTermPricingType: types.LongTermPricingType, // Required
		// SnowballType: types.SnowballType, // Required
	}

	if len(_snowballLongTermPricingType) > 0 {
		if err := assignInputField(input, "LongTermPricingType", _snowballLongTermPricingType); err != nil {
			log.Errorf("invalid --long-term-pricing-type: %s", err.Error())
			return
		}
	}
	if len(_snowballSnowballType) > 0 {
		if err := assignInputField(input, "SnowballType", _snowballSnowballType); err != nil {
			log.Errorf("invalid --snowball-type: %s", err.Error())
			return
		}
	}
	if len(_snowballIsLongTermPricingAutoRenew) > 0 {
		if err := assignInputField(input, "IsLongTermPricingAutoRenew", _snowballIsLongTermPricingAutoRenew); err != nil {
			log.Errorf("invalid --is-long-term-pricing-auto-renew: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLongTermPricing(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a shipping label that will be used to return the Snow device to Amazon
// Web Services.
func snowball_CreateReturnShippingLabel(cfg aws.Config, client *snowball.Client) {
	input := &snowball.CreateReturnShippingLabelInput{
		// JobId: *string, // Required
	}

	if len(_snowballJobId) > 0 {
		input.JobId = aws.String(_snowballJobId)
	}
	if len(_snowballShippingOption) > 0 {
		if err := assignInputField(input, "ShippingOption", _snowballShippingOption); err != nil {
			log.Errorf("invalid --shipping-option: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateReturnShippingLabel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Takes an AddressId and returns specific details about that address in the form
// of an Address object.
func snowball_DescribeAddress(cfg aws.Config, client *snowball.Client) {
	input := &snowball.DescribeAddressInput{
		// AddressId: *string, // Required
	}

	if len(_snowballAddressId) > 0 {
		input.AddressId = aws.String(_snowballAddressId)
	}

	if resp, err := client.DescribeAddress(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a specified number of ADDRESS objects. Calling this API in one of the
// US regions will return addresses from the list of all addresses associated with
// this account in all US regions.
func snowball_DescribeAddresses(cfg aws.Config, client *snowball.Client) {
	input := &snowball.DescribeAddressesInput{}

	if len(_snowballMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _snowballMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_snowballNextToken) > 0 {
		input.NextToken = aws.String(_snowballNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeAddresses(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*snowball.DescribeAddressesOutput
	p := snowball.NewDescribeAddressesPaginator(client, input)
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

// Returns information about a specific cluster including shipping information,
// cluster status, and other important metadata.
func snowball_DescribeCluster(cfg aws.Config, client *snowball.Client) {
	input := &snowball.DescribeClusterInput{
		// ClusterId: *string, // Required
	}

	if len(_snowballClusterId) > 0 {
		input.ClusterId = aws.String(_snowballClusterId)
	}

	if resp, err := client.DescribeCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specific job including shipping information, job
// status, and other important metadata.
func snowball_DescribeJob(cfg aws.Config, client *snowball.Client) {
	input := &snowball.DescribeJobInput{
		// JobId: *string, // Required
	}

	if len(_snowballJobId) > 0 {
		input.JobId = aws.String(_snowballJobId)
	}

	if resp, err := client.DescribeJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Information on the shipping label of a Snow device that is being returned to
// Amazon Web Services.
func snowball_DescribeReturnShippingLabel(cfg aws.Config, client *snowball.Client) {
	input := &snowball.DescribeReturnShippingLabelInput{
		// JobId: *string, // Required
	}

	if len(_snowballJobId) > 0 {
		input.JobId = aws.String(_snowballJobId)
	}

	if resp, err := client.DescribeReturnShippingLabel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a link to an Amazon S3 presigned URL for the manifest file associated
// with the specified JobId value. You can access the manifest file for up to 60
// minutes after this request has been made. To access the manifest file after 60
// minutes have passed, you'll have to make another call to the GetJobManifest
// action.
//
// The manifest is an encrypted file that you can download after your job enters
// the WithCustomer status. This is the only valid status for calling this API as
// the manifest and UnlockCode code value are used for securing your device and
// should only be used when you have the device. The manifest is decrypted by using
// the UnlockCode code value, when you pass both values to the Snow device through
// the Snowball client when the client is started for the first time.
//
// As a best practice, we recommend that you don't save a copy of an UnlockCode
// value in the same location as the manifest file for that job. Saving these
// separately helps prevent unauthorized parties from gaining access to the Snow
// device associated with that job.
//
// The credentials of a given job, including its manifest file and unlock code,
// expire 360 days after the job is created.
func snowball_GetJobManifest(cfg aws.Config, client *snowball.Client) {
	input := &snowball.GetJobManifestInput{
		// JobId: *string, // Required
	}

	if len(_snowballJobId) > 0 {
		input.JobId = aws.String(_snowballJobId)
	}

	if resp, err := client.GetJobManifest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the UnlockCode code value for the specified job. A particular UnlockCode
// value can be accessed for up to 360 days after the associated job has been
// created.
//
// The UnlockCode value is a 29-character code with 25 alphanumeric characters and
// 4 hyphens. This code is used to decrypt the manifest file when it is passed
// along with the manifest to the Snow device through the Snowball client when the
// client is started for the first time. The only valid status for calling this API
// is WithCustomer as the manifest and Unlock code values are used for securing
// your device and should only be used when you have the device.
//
// As a best practice, we recommend that you don't save a copy of the UnlockCode
// in the same location as the manifest file for that job. Saving these separately
// helps prevent unauthorized parties from gaining access to the Snow device
// associated with that job.
func snowball_GetJobUnlockCode(cfg aws.Config, client *snowball.Client) {
	input := &snowball.GetJobUnlockCodeInput{
		// JobId: *string, // Required
	}

	if len(_snowballJobId) > 0 {
		input.JobId = aws.String(_snowballJobId)
	}

	if resp, err := client.GetJobUnlockCode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the Snow Family service limit for your account, and
// also the number of Snow devices your account has in use.
//
// The default service limit for the number of Snow devices that you can have at
// one time is 1. If you want to increase your service limit, contact Amazon Web
// Services Support.
func snowball_GetSnowballUsage(cfg aws.Config, client *snowball.Client) {
	input := &snowball.GetSnowballUsageInput{}

	if resp, err := client.GetSnowballUsage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an Amazon S3 presigned URL for an update file associated with a
// specified JobId .
func snowball_GetSoftwareUpdates(cfg aws.Config, client *snowball.Client) {
	input := &snowball.GetSoftwareUpdatesInput{
		// JobId: *string, // Required
	}

	if len(_snowballJobId) > 0 {
		input.JobId = aws.String(_snowballJobId)
	}

	if resp, err := client.GetSoftwareUpdates(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an array of JobListEntry objects of the specified length. Each
// JobListEntry object is for a job in the specified cluster and contains a job's
// state, a job's ID, and other information.
func snowball_ListClusterJobs(cfg aws.Config, client *snowball.Client) {
	input := &snowball.ListClusterJobsInput{
		// ClusterId: *string, // Required
	}

	if len(_snowballClusterId) > 0 {
		input.ClusterId = aws.String(_snowballClusterId)
	}
	if len(_snowballMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _snowballMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_snowballNextToken) > 0 {
		input.NextToken = aws.String(_snowballNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListClusterJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*snowball.ListClusterJobsOutput
	p := snowball.NewListClusterJobsPaginator(client, input)
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

// Returns an array of ClusterListEntry objects of the specified length. Each
// ClusterListEntry object contains a cluster's state, a cluster's ID, and other
// important status information.
func snowball_ListClusters(cfg aws.Config, client *snowball.Client) {
	input := &snowball.ListClustersInput{}

	if len(_snowballMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _snowballMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_snowballNextToken) > 0 {
		input.NextToken = aws.String(_snowballNextToken)
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

	var results []*snowball.ListClustersOutput
	p := snowball.NewListClustersPaginator(client, input)
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

// This action returns a list of the different Amazon EC2-compatible Amazon
// Machine Images (AMIs) that are owned by your Amazon Web Services accountthat
// would be supported for use on a Snow device. Currently, supported AMIs are based
// on the Amazon Linux-2, Ubuntu 20.04 LTS - Focal, or Ubuntu 22.04 LTS - Jammy
// images, available on the Amazon Web Services Marketplace. Ubuntu 16.04 LTS -
// Xenial (HVM) images are no longer supported in the Market, but still supported
// for use on devices through Amazon EC2 VM Import/Export and running locally in
// AMIs.
func snowball_ListCompatibleImages(cfg aws.Config, client *snowball.Client) {
	input := &snowball.ListCompatibleImagesInput{}

	if len(_snowballMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _snowballMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_snowballNextToken) > 0 {
		input.NextToken = aws.String(_snowballNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCompatibleImages(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*snowball.ListCompatibleImagesOutput
	p := snowball.NewListCompatibleImagesPaginator(client, input)
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

// Returns an array of JobListEntry objects of the specified length. Each
// JobListEntry object contains a job's state, a job's ID, and a value that
// indicates whether the job is a job part, in the case of export jobs. Calling
// this API action in one of the US regions will return jobs from the list of all
// jobs associated with this account in all US regions.
func snowball_ListJobs(cfg aws.Config, client *snowball.Client) {
	input := &snowball.ListJobsInput{}

	if len(_snowballMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _snowballMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_snowballNextToken) > 0 {
		input.NextToken = aws.String(_snowballNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*snowball.ListJobsOutput
	p := snowball.NewListJobsPaginator(client, input)
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

// Lists all long-term pricing types.
func snowball_ListLongTermPricing(cfg aws.Config, client *snowball.Client) {
	input := &snowball.ListLongTermPricingInput{}

	if len(_snowballMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _snowballMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_snowballNextToken) > 0 {
		input.NextToken = aws.String(_snowballNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLongTermPricing(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*snowball.ListLongTermPricingOutput
	p := snowball.NewListLongTermPricingPaginator(client, input)
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

// A list of locations from which the customer can choose to pickup a device.
func snowball_ListPickupLocations(cfg aws.Config, client *snowball.Client) {
	input := &snowball.ListPickupLocationsInput{}

	if len(_snowballMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _snowballMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_snowballNextToken) > 0 {
		input.NextToken = aws.String(_snowballNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPickupLocations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*snowball.ListPickupLocationsOutput
	p := snowball.NewListPickupLocationsPaginator(client, input)
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

// Lists all supported versions for Snow on-device services. Returns an array of
// ServiceVersion object containing the supported versions for a particular service.
func snowball_ListServiceVersions(cfg aws.Config, client *snowball.Client) {
	input := &snowball.ListServiceVersionsInput{
		// ServiceName: types.ServiceName, // Required
	}

	if len(_snowballServiceName) > 0 {
		if err := assignInputField(input, "ServiceName", _snowballServiceName); err != nil {
			log.Errorf("invalid --service-name: %s", err.Error())
			return
		}
	}
	if len(_snowballDependentServices) > 0 {
		if err := assignInputField(input, "DependentServices", _snowballDependentServices); err != nil {
			log.Errorf("invalid --dependent-services: %s", err.Error())
			return
		}
	}
	if len(_snowballMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _snowballMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_snowballNextToken) > 0 {
		input.NextToken = aws.String(_snowballNextToken)
	}

	if resp, err := client.ListServiceVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// While a cluster's ClusterState value is in the AwaitingQuorum state, you can
// update some of the information associated with a cluster. Once the cluster
// changes to a different job state, usually 60 minutes after the cluster being
// created, this action is no longer available.
func snowball_UpdateCluster(cfg aws.Config, client *snowball.Client) {
	input := &snowball.UpdateClusterInput{
		// ClusterId: *string, // Required
	}

	if len(_snowballClusterId) > 0 {
		input.ClusterId = aws.String(_snowballClusterId)
	}
	if len(_snowballAddressId) > 0 {
		input.AddressId = aws.String(_snowballAddressId)
	}
	if len(_snowballDescription) > 0 {
		input.Description = aws.String(_snowballDescription)
	}
	if len(_snowballForwardingAddressId) > 0 {
		input.ForwardingAddressId = aws.String(_snowballForwardingAddressId)
	}
	if len(_snowballNotification) > 0 {
		if err := assignInputField(input, "Notification", _snowballNotification); err != nil {
			log.Errorf("invalid --notification: %s", err.Error())
			return
		}
	}
	if len(_snowballOnDeviceServiceConfiguration) > 0 {
		if err := assignInputField(input, "OnDeviceServiceConfiguration", _snowballOnDeviceServiceConfiguration); err != nil {
			log.Errorf("invalid --on-device-service-configuration: %s", err.Error())
			return
		}
	}
	if len(_snowballResources) > 0 {
		if err := assignInputField(input, "Resources", _snowballResources); err != nil {
			log.Errorf("invalid --resources: %s", err.Error())
			return
		}
	}
	if len(_snowballRoleARN) > 0 {
		input.RoleARN = aws.String(_snowballRoleARN)
	}
	if len(_snowballShippingOption) > 0 {
		if err := assignInputField(input, "ShippingOption", _snowballShippingOption); err != nil {
			log.Errorf("invalid --shipping-option: %s", err.Error())
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

// While a job's JobState value is New , you can update some of the information
// associated with a job. Once the job changes to a different job state, usually
// within 60 minutes of the job being created, this action is no longer available.
func snowball_UpdateJob(cfg aws.Config, client *snowball.Client) {
	input := &snowball.UpdateJobInput{
		// JobId: *string, // Required
	}

	if len(_snowballJobId) > 0 {
		input.JobId = aws.String(_snowballJobId)
	}
	if len(_snowballAddressId) > 0 {
		input.AddressId = aws.String(_snowballAddressId)
	}
	if len(_snowballDescription) > 0 {
		input.Description = aws.String(_snowballDescription)
	}
	if len(_snowballForwardingAddressId) > 0 {
		input.ForwardingAddressId = aws.String(_snowballForwardingAddressId)
	}
	if len(_snowballNotification) > 0 {
		if err := assignInputField(input, "Notification", _snowballNotification); err != nil {
			log.Errorf("invalid --notification: %s", err.Error())
			return
		}
	}
	if len(_snowballOnDeviceServiceConfiguration) > 0 {
		if err := assignInputField(input, "OnDeviceServiceConfiguration", _snowballOnDeviceServiceConfiguration); err != nil {
			log.Errorf("invalid --on-device-service-configuration: %s", err.Error())
			return
		}
	}
	if len(_snowballPickupDetails) > 0 {
		if err := assignInputField(input, "PickupDetails", _snowballPickupDetails); err != nil {
			log.Errorf("invalid --pickup-details: %s", err.Error())
			return
		}
	}
	if len(_snowballResources) > 0 {
		if err := assignInputField(input, "Resources", _snowballResources); err != nil {
			log.Errorf("invalid --resources: %s", err.Error())
			return
		}
	}
	if len(_snowballRoleARN) > 0 {
		input.RoleARN = aws.String(_snowballRoleARN)
	}
	if len(_snowballShippingOption) > 0 {
		if err := assignInputField(input, "ShippingOption", _snowballShippingOption); err != nil {
			log.Errorf("invalid --shipping-option: %s", err.Error())
			return
		}
	}
	if len(_snowballSnowballCapacityPreference) > 0 {
		if err := assignInputField(input, "SnowballCapacityPreference", _snowballSnowballCapacityPreference); err != nil {
			log.Errorf("invalid --snowball-capacity-preference: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the state when a shipment state changes to a different state.
func snowball_UpdateJobShipmentState(cfg aws.Config, client *snowball.Client) {
	input := &snowball.UpdateJobShipmentStateInput{
		// JobId: *string, // Required
		// ShipmentState: types.ShipmentState, // Required
	}

	if len(_snowballJobId) > 0 {
		input.JobId = aws.String(_snowballJobId)
	}
	if len(_snowballShipmentState) > 0 {
		if err := assignInputField(input, "ShipmentState", _snowballShipmentState); err != nil {
			log.Errorf("invalid --shipment-state: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateJobShipmentState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the long-term pricing type.
func snowball_UpdateLongTermPricing(cfg aws.Config, client *snowball.Client) {
	input := &snowball.UpdateLongTermPricingInput{
		// LongTermPricingId: *string, // Required
	}

	if len(_snowballLongTermPricingId) > 0 {
		input.LongTermPricingId = aws.String(_snowballLongTermPricingId)
	}
	if len(_snowballIsLongTermPricingAutoRenew) > 0 {
		if err := assignInputField(input, "IsLongTermPricingAutoRenew", _snowballIsLongTermPricingAutoRenew); err != nil {
			log.Errorf("invalid --is-long-term-pricing-auto-renew: %s", err.Error())
			return
		}
	}
	if len(_snowballReplacementJob) > 0 {
		input.ReplacementJob = aws.String(_snowballReplacementJob)
	}

	if resp, err := client.UpdateLongTermPricing(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_snowballCmd)
	_snowballCmd.Flags().SortFlags = false

	_snowballCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_snowballCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_snowballCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_snowballCmd.Flags().StringVarP(&_snowballAddress, "address", "", "", "Address")
	_snowballCmd.Flags().StringVarP(&_snowballAddressId, "address-id", "", "", "Address ID")
	_snowballCmd.Flags().StringVarP(&_snowballClusterId, "cluster-id", "", "", "Cluster ID")
	_snowballCmd.Flags().StringVarP(&_snowballDependentServices, "dependent-services", "", "", "Dependent Services")
	_snowballCmd.Flags().StringVarP(&_snowballDescription, "description", "", "", "Description")
	_snowballCmd.Flags().StringVarP(&_snowballDeviceConfiguration, "device-configuration", "", "", "Device Configuration")
	_snowballCmd.Flags().StringVarP(&_snowballForceCreateJobs, "force-create-jobs", "", "", "Force Create Jobs")
	_snowballCmd.Flags().StringVarP(&_snowballForwardingAddressId, "forwarding-address-id", "", "", "Forwarding Address ID")
	_snowballCmd.Flags().StringVarP(&_snowballImpactLevel, "impact-level", "", "", "Impact Level")
	_snowballCmd.Flags().StringVarP(&_snowballInitialClusterSize, "initial-cluster-size", "", "", "Initial Cluster Size")
	_snowballCmd.Flags().StringVarP(&_snowballIsLongTermPricingAutoRenew, "is-long-term-pricing-auto-renew", "", "", "Is Long Term Pricing Auto Renew")
	_snowballCmd.Flags().StringVarP(&_snowballJobId, "job-id", "", "", "Job ID")
	_snowballCmd.Flags().StringVarP(&_snowballJobType, "job-type", "", "", "Job Type")
	_snowballCmd.Flags().StringVarP(&_snowballKmsKeyARN, "kms-key-arn", "", "", "KMS Key ARN")
	_snowballCmd.Flags().StringVarP(&_snowballLongTermPricingId, "long-term-pricing-id", "", "", "Long Term Pricing ID")
	_snowballCmd.Flags().StringSliceVarP(&_snowballLongTermPricingIds, "long-term-pricing-ids", "", nil, "Long Term Pricing Ids")
	_snowballCmd.Flags().StringVarP(&_snowballLongTermPricingType, "long-term-pricing-type", "", "", "Long Term Pricing Type")
	_snowballCmd.Flags().StringVarP(&_snowballMaxResults, "max-results", "", "", "Max Results")
	_snowballCmd.Flags().StringVarP(&_snowballNextToken, "next-token", "", "", "Next Token")
	_snowballCmd.Flags().StringVarP(&_snowballNotification, "notification", "", "", "Notification")
	_snowballCmd.Flags().StringVarP(&_snowballOnDeviceServiceConfiguration, "on-device-service-configuration", "", "", "On Device Service Configuration")
	_snowballCmd.Flags().StringVarP(&_snowballPickupDetails, "pickup-details", "", "", "Pickup Details")
	_snowballCmd.Flags().StringVarP(&_snowballRemoteManagement, "remote-management", "", "", "Remote Management")
	_snowballCmd.Flags().StringVarP(&_snowballReplacementJob, "replacement-job", "", "", "Replacement Job")
	_snowballCmd.Flags().StringVarP(&_snowballResources, "resources", "", "", "Resources")
	_snowballCmd.Flags().StringVarP(&_snowballRoleARN, "role-arn", "", "", "Role ARN")
	_snowballCmd.Flags().StringVarP(&_snowballServiceName, "service-name", "", "", "Service Name")
	_snowballCmd.Flags().StringVarP(&_snowballShipmentState, "shipment-state", "", "", "Shipment State")
	_snowballCmd.Flags().StringVarP(&_snowballShippingOption, "shipping-option", "", "", "Shipping Option")
	_snowballCmd.Flags().StringVarP(&_snowballSnowballCapacityPreference, "snowball-capacity-preference", "", "", "Snowball Capacity Preference")
	_snowballCmd.Flags().StringVarP(&_snowballSnowballType, "snowball-type", "", "", "Snowball Type")
	_snowballCmd.Flags().StringVarP(&_snowballTaxDocuments, "tax-documents", "", "", "Tax Documents")

	_snowballCmd.Flags().BoolVarP(&_snowballCancelCluster, "cancel-cluster", "", false, "Cancel Cluster")
	_snowballCmd.Flags().BoolVarP(&_snowballCancelJob, "cancel-job", "", false, "Cancel Job")
	_snowballCmd.Flags().BoolVarP(&_snowballCreateAddress, "create-address", "", false, "Create Address")
	_snowballCmd.Flags().BoolVarP(&_snowballCreateCluster, "create-cluster", "", false, "Create Cluster")
	_snowballCmd.Flags().BoolVarP(&_snowballCreateJob, "create-job", "", false, "Create Job")
	_snowballCmd.Flags().BoolVarP(&_snowballCreateLongTermPricing, "create-long-term-pricing", "", false, "Create Long Term Pricing")
	_snowballCmd.Flags().BoolVarP(&_snowballCreateReturnShippingLabel, "create-return-shipping-label", "", false, "Create Return Shipping Label")
	_snowballCmd.Flags().BoolVarP(&_snowballDescribeAddress, "describe-address", "", false, "Describe Address")
	_snowballCmd.Flags().BoolVarP(&_snowballDescribeAddresses, "describe-addresses", "", false, "Describe Addresses")
	_snowballCmd.Flags().BoolVarP(&_snowballDescribeCluster, "describe-cluster", "", false, "Describe Cluster")
	_snowballCmd.Flags().BoolVarP(&_snowballDescribeJob, "describe-job", "", false, "Describe Job")
	_snowballCmd.Flags().BoolVarP(&_snowballDescribeReturnShippingLabel, "describe-return-shipping-label", "", false, "Describe Return Shipping Label")
	_snowballCmd.Flags().BoolVarP(&_snowballGetJobManifest, "get-job-manifest", "", false, "Get Job Manifest")
	_snowballCmd.Flags().BoolVarP(&_snowballGetJobUnlockCode, "get-job-unlock-code", "", false, "Get Job Unlock Code")
	_snowballCmd.Flags().BoolVarP(&_snowballGetSnowballUsage, "get-snowball-usage", "", false, "Get Snowball Usage")
	_snowballCmd.Flags().BoolVarP(&_snowballGetSoftwareUpdates, "get-software-updates", "", false, "Get Software Updates")
	_snowballCmd.Flags().BoolVarP(&_snowballListClusterJobs, "list-cluster-jobs", "", false, "List Cluster Jobs")
	_snowballCmd.Flags().BoolVarP(&_snowballListClusters, "list-clusters", "", false, "List Clusters")
	_snowballCmd.Flags().BoolVarP(&_snowballListCompatibleImages, "list-compatible-images", "", false, "List Compatible Images")
	_snowballCmd.Flags().BoolVarP(&_snowballListJobs, "list-jobs", "", false, "List Jobs")
	_snowballCmd.Flags().BoolVarP(&_snowballListLongTermPricing, "list-long-term-pricing", "", false, "List Long Term Pricing")
	_snowballCmd.Flags().BoolVarP(&_snowballListPickupLocations, "list-pickup-locations", "", false, "List Pickup Locations")
	_snowballCmd.Flags().BoolVarP(&_snowballListServiceVersions, "list-service-versions", "", false, "List Service Versions")
	_snowballCmd.Flags().BoolVarP(&_snowballUpdateCluster, "update-cluster", "", false, "Update Cluster")
	_snowballCmd.Flags().BoolVarP(&_snowballUpdateJob, "update-job", "", false, "Update Job")
	_snowballCmd.Flags().BoolVarP(&_snowballUpdateJobShipmentState, "update-job-shipment-state", "", false, "Update Job Shipment State")
	_snowballCmd.Flags().BoolVarP(&_snowballUpdateLongTermPricing, "update-long-term-pricing", "", false, "Update Long Term Pricing")

}
