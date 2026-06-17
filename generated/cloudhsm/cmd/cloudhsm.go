package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cloudhsmCmd represents the cloudhsm command
var _cloudhsmCmd = &cobra.Command{
	Use:   "cloudhsm",
	Short: "AWS cloudhsm CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := cloudhsm.NewFromConfig(cfg)
		if _cloudhsmAddTagsToResource {
			cloudhsm_AddTagsToResource(cfg, client)
			return
		}
		if _cloudhsmCreateHapg {
			cloudhsm_CreateHapg(cfg, client)
			return
		}
		if _cloudhsmCreateHsm {
			cloudhsm_CreateHsm(cfg, client)
			return
		}
		if _cloudhsmCreateLunaClient {
			cloudhsm_CreateLunaClient(cfg, client)
			return
		}
		if _cloudhsmDeleteHapg {
			cloudhsm_DeleteHapg(cfg, client)
			return
		}
		if _cloudhsmDeleteHsm {
			cloudhsm_DeleteHsm(cfg, client)
			return
		}
		if _cloudhsmDeleteLunaClient {
			cloudhsm_DeleteLunaClient(cfg, client)
			return
		}
		if _cloudhsmDescribeHapg {
			cloudhsm_DescribeHapg(cfg, client)
			return
		}
		if _cloudhsmDescribeHsm {
			cloudhsm_DescribeHsm(cfg, client)
			return
		}
		if _cloudhsmDescribeLunaClient {
			cloudhsm_DescribeLunaClient(cfg, client)
			return
		}
		if _cloudhsmGetConfig {
			cloudhsm_GetConfig(cfg, client)
			return
		}
		if _cloudhsmListAvailableZones {
			cloudhsm_ListAvailableZones(cfg, client)
			return
		}
		if _cloudhsmListHapgs {
			cloudhsm_ListHapgs(cfg, client)
			return
		}
		if _cloudhsmListHsms {
			cloudhsm_ListHsms(cfg, client)
			return
		}
		if _cloudhsmListLunaClients {
			cloudhsm_ListLunaClients(cfg, client)
			return
		}
		if _cloudhsmListTagsForResource {
			cloudhsm_ListTagsForResource(cfg, client)
			return
		}
		if _cloudhsmModifyHapg {
			cloudhsm_ModifyHapg(cfg, client)
			return
		}
		if _cloudhsmModifyHsm {
			cloudhsm_ModifyHsm(cfg, client)
			return
		}
		if _cloudhsmModifyLunaClient {
			cloudhsm_ModifyLunaClient(cfg, client)
			return
		}
		if _cloudhsmRemoveTagsFromResource {
			cloudhsm_RemoveTagsFromResource(cfg, client)
			return
		}

	},
}

var (
	_cloudhsmAddTagsToResource      bool
	_cloudhsmCreateHapg             bool
	_cloudhsmCreateHsm              bool
	_cloudhsmCreateLunaClient       bool
	_cloudhsmDeleteHapg             bool
	_cloudhsmDeleteHsm              bool
	_cloudhsmDeleteLunaClient       bool
	_cloudhsmDescribeHapg           bool
	_cloudhsmDescribeHsm            bool
	_cloudhsmDescribeLunaClient     bool
	_cloudhsmGetConfig              bool
	_cloudhsmListAvailableZones     bool
	_cloudhsmListHapgs              bool
	_cloudhsmListHsms               bool
	_cloudhsmListLunaClients        bool
	_cloudhsmListTagsForResource    bool
	_cloudhsmModifyHapg             bool
	_cloudhsmModifyHsm              bool
	_cloudhsmModifyLunaClient       bool
	_cloudhsmRemoveTagsFromResource bool

	_cloudhsmCertificate            string
	_cloudhsmCertificateFingerprint string
	_cloudhsmClientArn              string
	_cloudhsmClientToken            string
	_cloudhsmClientVersion          string
	_cloudhsmEniIp                  string
	_cloudhsmExternalId             string
	_cloudhsmHapgArn                string
	_cloudhsmHapgList               []string
	_cloudhsmHsmArn                 string
	_cloudhsmHsmSerialNumber        string
	_cloudhsmIamRoleArn             string
	_cloudhsmLabel                  string
	_cloudhsmNextToken              string
	_cloudhsmPartitionSerialList    []string
	_cloudhsmResourceArn            string
	_cloudhsmSshKey                 string
	_cloudhsmSubnetId               string
	_cloudhsmSubscriptionType       string
	_cloudhsmSyslogIp               string
	_cloudhsmTagKeyList             []string
	_cloudhsmTagList                string
)

// This is documentation for AWS CloudHSM Classic. For more information, see [AWS CloudHSM Classic FAQs], the [AWS CloudHSM Classic User Guide]
// , and the [AWS CloudHSM Classic API Reference].
//
// For information about the current version of AWS CloudHSM, see [AWS CloudHSM], the [AWS CloudHSM User Guide], and the [AWS CloudHSM API Reference].
//
// Adds or overwrites one or more tags for the specified AWS CloudHSM resource.
//
// Each tag consists of a key and a value. Tag keys must be unique to each
// resource.
//
// Deprecated: This API is deprecated.
//
// [AWS CloudHSM User Guide]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/
// [AWS CloudHSM]: http://aws.amazon.com/cloudhsm/
// [AWS CloudHSM API Reference]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/
// [AWS CloudHSM Classic API Reference]: https://docs.aws.amazon.com/cloudhsm/classic/APIReference/
// [AWS CloudHSM Classic FAQs]: http://aws.amazon.com/cloudhsm/faqs-classic/
// [AWS CloudHSM Classic User Guide]: https://docs.aws.amazon.com/cloudhsm/classic/userguide/
func cloudhsm_AddTagsToResource(cfg aws.Config, client *cloudhsm.Client) {
	input := &cloudhsm.AddTagsToResourceInput{
		// ResourceArn: *string, // Required
		// TagList: []types.Tag, // Required
	}

	if len(_cloudhsmResourceArn) > 0 {
		input.ResourceArn = aws.String(_cloudhsmResourceArn)
	}
	if len(_cloudhsmTagList) > 0 {
		if err := assignInputField(input, "TagList", _cloudhsmTagList); err != nil {
			log.Errorf("invalid --tag-list: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddTagsToResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is documentation for AWS CloudHSM Classic. For more information, see [AWS CloudHSM Classic FAQs], the [AWS CloudHSM Classic User Guide]
// , and the [AWS CloudHSM Classic API Reference].
//
// For information about the current version of AWS CloudHSM, see [AWS CloudHSM], the [AWS CloudHSM User Guide], and the [AWS CloudHSM API Reference].
//
// Creates a high-availability partition group. A high-availability partition
// group is a group of partitions that spans multiple physical HSMs.
//
// Deprecated: This API is deprecated.
//
// [AWS CloudHSM User Guide]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/
// [AWS CloudHSM]: http://aws.amazon.com/cloudhsm/
// [AWS CloudHSM API Reference]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/
// [AWS CloudHSM Classic API Reference]: https://docs.aws.amazon.com/cloudhsm/classic/APIReference/
// [AWS CloudHSM Classic FAQs]: http://aws.amazon.com/cloudhsm/faqs-classic/
// [AWS CloudHSM Classic User Guide]: https://docs.aws.amazon.com/cloudhsm/classic/userguide/
func cloudhsm_CreateHapg(cfg aws.Config, client *cloudhsm.Client) {
	input := &cloudhsm.CreateHapgInput{
		// Label: *string, // Required
	}

	if len(_cloudhsmLabel) > 0 {
		input.Label = aws.String(_cloudhsmLabel)
	}

	if resp, err := client.CreateHapg(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is documentation for AWS CloudHSM Classic. For more information, see [AWS CloudHSM Classic FAQs], the [AWS CloudHSM Classic User Guide]
// , and the [AWS CloudHSM Classic API Reference].
//
// For information about the current version of AWS CloudHSM, see [AWS CloudHSM], the [AWS CloudHSM User Guide], and the [AWS CloudHSM API Reference].
//
// Creates an uninitialized HSM instance.
//
// There is an upfront fee charged for each HSM instance that you create with the
// CreateHsm operation. If you accidentally provision an HSM and want to request a
// refund, delete the instance using the DeleteHsmoperation, go to the [AWS Support Center], create a new case,
// and select Account and Billing Support.
//
// It can take up to 20 minutes to create and provision an HSM. You can monitor
// the status of the HSM with the DescribeHsmoperation. The HSM is ready to be initialized
// when the status changes to RUNNING .
//
// Deprecated: This API is deprecated.
//
// [AWS CloudHSM User Guide]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/
// [AWS CloudHSM]: http://aws.amazon.com/cloudhsm/
// [AWS CloudHSM API Reference]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/
// [AWS CloudHSM Classic API Reference]: https://docs.aws.amazon.com/cloudhsm/classic/APIReference/
// [AWS Support Center]: https://console.aws.amazon.com/support/home
// [AWS CloudHSM Classic FAQs]: http://aws.amazon.com/cloudhsm/faqs-classic/
// [AWS CloudHSM Classic User Guide]: https://docs.aws.amazon.com/cloudhsm/classic/userguide/
func cloudhsm_CreateHsm(cfg aws.Config, client *cloudhsm.Client) {
	input := &cloudhsm.CreateHsmInput{
		// IamRoleArn: *string, // Required
		// SshKey: *string, // Required
		// SubnetId: *string, // Required
		// SubscriptionType: types.SubscriptionType, // Required
	}

	if len(_cloudhsmIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_cloudhsmIamRoleArn)
	}
	if len(_cloudhsmSshKey) > 0 {
		input.SshKey = aws.String(_cloudhsmSshKey)
	}
	if len(_cloudhsmSubnetId) > 0 {
		input.SubnetId = aws.String(_cloudhsmSubnetId)
	}
	if len(_cloudhsmSubscriptionType) > 0 {
		if err := assignInputField(input, "SubscriptionType", _cloudhsmSubscriptionType); err != nil {
			log.Errorf("invalid --subscription-type: %s", err.Error())
			return
		}
	}
	if len(_cloudhsmClientToken) > 0 {
		input.ClientToken = aws.String(_cloudhsmClientToken)
	}
	if len(_cloudhsmEniIp) > 0 {
		input.EniIp = aws.String(_cloudhsmEniIp)
	}
	if len(_cloudhsmExternalId) > 0 {
		input.ExternalId = aws.String(_cloudhsmExternalId)
	}
	if len(_cloudhsmSyslogIp) > 0 {
		input.SyslogIp = aws.String(_cloudhsmSyslogIp)
	}

	if resp, err := client.CreateHsm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is documentation for AWS CloudHSM Classic. For more information, see [AWS CloudHSM Classic FAQs], the [AWS CloudHSM Classic User Guide]
// , and the [AWS CloudHSM Classic API Reference].
//
// For information about the current version of AWS CloudHSM, see [AWS CloudHSM], the [AWS CloudHSM User Guide], and the [AWS CloudHSM API Reference].
//
// Creates an HSM client.
//
// Deprecated: This API is deprecated.
//
// [AWS CloudHSM User Guide]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/
// [AWS CloudHSM]: http://aws.amazon.com/cloudhsm/
// [AWS CloudHSM API Reference]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/
// [AWS CloudHSM Classic API Reference]: https://docs.aws.amazon.com/cloudhsm/classic/APIReference/
// [AWS CloudHSM Classic FAQs]: http://aws.amazon.com/cloudhsm/faqs-classic/
// [AWS CloudHSM Classic User Guide]: https://docs.aws.amazon.com/cloudhsm/classic/userguide/
func cloudhsm_CreateLunaClient(cfg aws.Config, client *cloudhsm.Client) {
	input := &cloudhsm.CreateLunaClientInput{
		// Certificate: *string, // Required
	}

	if len(_cloudhsmCertificate) > 0 {
		input.Certificate = aws.String(_cloudhsmCertificate)
	}
	if len(_cloudhsmLabel) > 0 {
		input.Label = aws.String(_cloudhsmLabel)
	}

	if resp, err := client.CreateLunaClient(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is documentation for AWS CloudHSM Classic. For more information, see [AWS CloudHSM Classic FAQs], the [AWS CloudHSM Classic User Guide]
// , and the [AWS CloudHSM Classic API Reference].
//
// For information about the current version of AWS CloudHSM, see [AWS CloudHSM], the [AWS CloudHSM User Guide], and the [AWS CloudHSM API Reference].
//
// Deletes a high-availability partition group.
//
// Deprecated: This API is deprecated.
//
// [AWS CloudHSM User Guide]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/
// [AWS CloudHSM]: http://aws.amazon.com/cloudhsm/
// [AWS CloudHSM API Reference]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/
// [AWS CloudHSM Classic API Reference]: https://docs.aws.amazon.com/cloudhsm/classic/APIReference/
// [AWS CloudHSM Classic FAQs]: http://aws.amazon.com/cloudhsm/faqs-classic/
// [AWS CloudHSM Classic User Guide]: https://docs.aws.amazon.com/cloudhsm/classic/userguide/
func cloudhsm_DeleteHapg(cfg aws.Config, client *cloudhsm.Client) {
	input := &cloudhsm.DeleteHapgInput{
		// HapgArn: *string, // Required
	}

	if len(_cloudhsmHapgArn) > 0 {
		input.HapgArn = aws.String(_cloudhsmHapgArn)
	}

	if resp, err := client.DeleteHapg(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is documentation for AWS CloudHSM Classic. For more information, see [AWS CloudHSM Classic FAQs], the [AWS CloudHSM Classic User Guide]
// , and the [AWS CloudHSM Classic API Reference].
//
// For information about the current version of AWS CloudHSM, see [AWS CloudHSM], the [AWS CloudHSM User Guide], and the [AWS CloudHSM API Reference].
//
// Deletes an HSM. After completion, this operation cannot be undone and your key
// material cannot be recovered.
//
// Deprecated: This API is deprecated.
//
// [AWS CloudHSM User Guide]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/
// [AWS CloudHSM]: http://aws.amazon.com/cloudhsm/
// [AWS CloudHSM API Reference]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/
// [AWS CloudHSM Classic API Reference]: https://docs.aws.amazon.com/cloudhsm/classic/APIReference/
// [AWS CloudHSM Classic FAQs]: http://aws.amazon.com/cloudhsm/faqs-classic/
// [AWS CloudHSM Classic User Guide]: https://docs.aws.amazon.com/cloudhsm/classic/userguide/
func cloudhsm_DeleteHsm(cfg aws.Config, client *cloudhsm.Client) {
	input := &cloudhsm.DeleteHsmInput{
		// HsmArn: *string, // Required
	}

	if len(_cloudhsmHsmArn) > 0 {
		input.HsmArn = aws.String(_cloudhsmHsmArn)
	}

	if resp, err := client.DeleteHsm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is documentation for AWS CloudHSM Classic. For more information, see [AWS CloudHSM Classic FAQs], the [AWS CloudHSM Classic User Guide]
// , and the [AWS CloudHSM Classic API Reference].
//
// For information about the current version of AWS CloudHSM, see [AWS CloudHSM], the [AWS CloudHSM User Guide], and the [AWS CloudHSM API Reference].
//
// Deletes a client.
//
// Deprecated: This API is deprecated.
//
// [AWS CloudHSM User Guide]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/
// [AWS CloudHSM]: http://aws.amazon.com/cloudhsm/
// [AWS CloudHSM API Reference]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/
// [AWS CloudHSM Classic API Reference]: https://docs.aws.amazon.com/cloudhsm/classic/APIReference/
// [AWS CloudHSM Classic FAQs]: http://aws.amazon.com/cloudhsm/faqs-classic/
// [AWS CloudHSM Classic User Guide]: https://docs.aws.amazon.com/cloudhsm/classic/userguide/
func cloudhsm_DeleteLunaClient(cfg aws.Config, client *cloudhsm.Client) {
	input := &cloudhsm.DeleteLunaClientInput{
		// ClientArn: *string, // Required
	}

	if len(_cloudhsmClientArn) > 0 {
		input.ClientArn = aws.String(_cloudhsmClientArn)
	}

	if resp, err := client.DeleteLunaClient(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is documentation for AWS CloudHSM Classic. For more information, see [AWS CloudHSM Classic FAQs], the [AWS CloudHSM Classic User Guide]
// , and the [AWS CloudHSM Classic API Reference].
//
// For information about the current version of AWS CloudHSM, see [AWS CloudHSM], the [AWS CloudHSM User Guide], and the [AWS CloudHSM API Reference].
//
// Retrieves information about a high-availability partition group.
//
// Deprecated: This API is deprecated.
//
// [AWS CloudHSM User Guide]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/
// [AWS CloudHSM]: http://aws.amazon.com/cloudhsm/
// [AWS CloudHSM API Reference]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/
// [AWS CloudHSM Classic API Reference]: https://docs.aws.amazon.com/cloudhsm/classic/APIReference/
// [AWS CloudHSM Classic FAQs]: http://aws.amazon.com/cloudhsm/faqs-classic/
// [AWS CloudHSM Classic User Guide]: https://docs.aws.amazon.com/cloudhsm/classic/userguide/
func cloudhsm_DescribeHapg(cfg aws.Config, client *cloudhsm.Client) {
	input := &cloudhsm.DescribeHapgInput{
		// HapgArn: *string, // Required
	}

	if len(_cloudhsmHapgArn) > 0 {
		input.HapgArn = aws.String(_cloudhsmHapgArn)
	}

	if resp, err := client.DescribeHapg(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is documentation for AWS CloudHSM Classic. For more information, see [AWS CloudHSM Classic FAQs], the [AWS CloudHSM Classic User Guide]
// , and the [AWS CloudHSM Classic API Reference].
//
// For information about the current version of AWS CloudHSM, see [AWS CloudHSM], the [AWS CloudHSM User Guide], and the [AWS CloudHSM API Reference].
//
// Retrieves information about an HSM. You can identify the HSM by its ARN or its
// serial number.
//
// Deprecated: This API is deprecated.
//
// [AWS CloudHSM User Guide]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/
// [AWS CloudHSM]: http://aws.amazon.com/cloudhsm/
// [AWS CloudHSM API Reference]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/
// [AWS CloudHSM Classic API Reference]: https://docs.aws.amazon.com/cloudhsm/classic/APIReference/
// [AWS CloudHSM Classic FAQs]: http://aws.amazon.com/cloudhsm/faqs-classic/
// [AWS CloudHSM Classic User Guide]: https://docs.aws.amazon.com/cloudhsm/classic/userguide/
func cloudhsm_DescribeHsm(cfg aws.Config, client *cloudhsm.Client) {
	input := &cloudhsm.DescribeHsmInput{}

	if len(_cloudhsmHsmArn) > 0 {
		input.HsmArn = aws.String(_cloudhsmHsmArn)
	}
	if len(_cloudhsmHsmSerialNumber) > 0 {
		input.HsmSerialNumber = aws.String(_cloudhsmHsmSerialNumber)
	}

	if resp, err := client.DescribeHsm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is documentation for AWS CloudHSM Classic. For more information, see [AWS CloudHSM Classic FAQs], the [AWS CloudHSM Classic User Guide]
// , and the [AWS CloudHSM Classic API Reference].
//
// For information about the current version of AWS CloudHSM, see [AWS CloudHSM], the [AWS CloudHSM User Guide], and the [AWS CloudHSM API Reference].
//
// Retrieves information about an HSM client.
//
// Deprecated: This API is deprecated.
//
// [AWS CloudHSM User Guide]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/
// [AWS CloudHSM]: http://aws.amazon.com/cloudhsm/
// [AWS CloudHSM API Reference]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/
// [AWS CloudHSM Classic API Reference]: https://docs.aws.amazon.com/cloudhsm/classic/APIReference/
// [AWS CloudHSM Classic FAQs]: http://aws.amazon.com/cloudhsm/faqs-classic/
// [AWS CloudHSM Classic User Guide]: https://docs.aws.amazon.com/cloudhsm/classic/userguide/
func cloudhsm_DescribeLunaClient(cfg aws.Config, client *cloudhsm.Client) {
	input := &cloudhsm.DescribeLunaClientInput{}

	if len(_cloudhsmCertificateFingerprint) > 0 {
		input.CertificateFingerprint = aws.String(_cloudhsmCertificateFingerprint)
	}
	if len(_cloudhsmClientArn) > 0 {
		input.ClientArn = aws.String(_cloudhsmClientArn)
	}

	if resp, err := client.DescribeLunaClient(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is documentation for AWS CloudHSM Classic. For more information, see [AWS CloudHSM Classic FAQs], the [AWS CloudHSM Classic User Guide]
// , and the [AWS CloudHSM Classic API Reference].
//
// For information about the current version of AWS CloudHSM, see [AWS CloudHSM], the [AWS CloudHSM User Guide], and the [AWS CloudHSM API Reference].
//
// Gets the configuration files necessary to connect to all high availability
// partition groups the client is associated with.
//
// Deprecated: This API is deprecated.
//
// [AWS CloudHSM User Guide]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/
// [AWS CloudHSM]: http://aws.amazon.com/cloudhsm/
// [AWS CloudHSM API Reference]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/
// [AWS CloudHSM Classic API Reference]: https://docs.aws.amazon.com/cloudhsm/classic/APIReference/
// [AWS CloudHSM Classic FAQs]: http://aws.amazon.com/cloudhsm/faqs-classic/
// [AWS CloudHSM Classic User Guide]: https://docs.aws.amazon.com/cloudhsm/classic/userguide/
func cloudhsm_GetConfig(cfg aws.Config, client *cloudhsm.Client) {
	input := &cloudhsm.GetConfigInput{
		// ClientArn: *string, // Required
		// ClientVersion: types.ClientVersion, // Required
		// HapgList: []string, // Required
	}

	if len(_cloudhsmClientArn) > 0 {
		input.ClientArn = aws.String(_cloudhsmClientArn)
	}
	if len(_cloudhsmClientVersion) > 0 {
		if err := assignInputField(input, "ClientVersion", _cloudhsmClientVersion); err != nil {
			log.Errorf("invalid --client-version: %s", err.Error())
			return
		}
	}
	if len(_cloudhsmHapgList) > 0 {
		input.HapgList = append([]string(nil), _cloudhsmHapgList...)
	}

	if resp, err := client.GetConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is documentation for AWS CloudHSM Classic. For more information, see [AWS CloudHSM Classic FAQs], the [AWS CloudHSM Classic User Guide]
// , and the [AWS CloudHSM Classic API Reference].
//
// For information about the current version of AWS CloudHSM, see [AWS CloudHSM], the [AWS CloudHSM User Guide], and the [AWS CloudHSM API Reference].
//
// Lists the Availability Zones that have available AWS CloudHSM capacity.
//
// Deprecated: This API is deprecated.
//
// [AWS CloudHSM User Guide]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/
// [AWS CloudHSM]: http://aws.amazon.com/cloudhsm/
// [AWS CloudHSM API Reference]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/
// [AWS CloudHSM Classic API Reference]: https://docs.aws.amazon.com/cloudhsm/classic/APIReference/
// [AWS CloudHSM Classic FAQs]: http://aws.amazon.com/cloudhsm/faqs-classic/
// [AWS CloudHSM Classic User Guide]: https://docs.aws.amazon.com/cloudhsm/classic/userguide/
func cloudhsm_ListAvailableZones(cfg aws.Config, client *cloudhsm.Client) {
	input := &cloudhsm.ListAvailableZonesInput{}

	if resp, err := client.ListAvailableZones(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is documentation for AWS CloudHSM Classic. For more information, see [AWS CloudHSM Classic FAQs], the [AWS CloudHSM Classic User Guide]
// , and the [AWS CloudHSM Classic API Reference].
//
// For information about the current version of AWS CloudHSM, see [AWS CloudHSM], the [AWS CloudHSM User Guide], and the [AWS CloudHSM API Reference].
//
// Lists the high-availability partition groups for the account.
//
// This operation supports pagination with the use of the NextToken member. If
// more results are available, the NextToken member of the response contains a
// token that you pass in the next call to ListHapgs to retrieve the next set of
// items.
//
// Deprecated: This API is deprecated.
//
// [AWS CloudHSM User Guide]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/
// [AWS CloudHSM]: http://aws.amazon.com/cloudhsm/
// [AWS CloudHSM API Reference]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/
// [AWS CloudHSM Classic API Reference]: https://docs.aws.amazon.com/cloudhsm/classic/APIReference/
// [AWS CloudHSM Classic FAQs]: http://aws.amazon.com/cloudhsm/faqs-classic/
// [AWS CloudHSM Classic User Guide]: https://docs.aws.amazon.com/cloudhsm/classic/userguide/
func cloudhsm_ListHapgs(cfg aws.Config, client *cloudhsm.Client) {
	input := &cloudhsm.ListHapgsInput{}

	if len(_cloudhsmNextToken) > 0 {
		input.NextToken = aws.String(_cloudhsmNextToken)
	}

	if resp, err := client.ListHapgs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is documentation for AWS CloudHSM Classic. For more information, see [AWS CloudHSM Classic FAQs], the [AWS CloudHSM Classic User Guide]
// , and the [AWS CloudHSM Classic API Reference].
//
// For information about the current version of AWS CloudHSM, see [AWS CloudHSM], the [AWS CloudHSM User Guide], and the [AWS CloudHSM API Reference].
//
// Retrieves the identifiers of all of the HSMs provisioned for the current
// customer.
//
// This operation supports pagination with the use of the NextToken member. If
// more results are available, the NextToken member of the response contains a
// token that you pass in the next call to ListHsms to retrieve the next set of
// items.
//
// Deprecated: This API is deprecated.
//
// [AWS CloudHSM User Guide]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/
// [AWS CloudHSM]: http://aws.amazon.com/cloudhsm/
// [AWS CloudHSM API Reference]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/
// [AWS CloudHSM Classic API Reference]: https://docs.aws.amazon.com/cloudhsm/classic/APIReference/
// [AWS CloudHSM Classic FAQs]: http://aws.amazon.com/cloudhsm/faqs-classic/
// [AWS CloudHSM Classic User Guide]: https://docs.aws.amazon.com/cloudhsm/classic/userguide/
func cloudhsm_ListHsms(cfg aws.Config, client *cloudhsm.Client) {
	input := &cloudhsm.ListHsmsInput{}

	if len(_cloudhsmNextToken) > 0 {
		input.NextToken = aws.String(_cloudhsmNextToken)
	}

	if resp, err := client.ListHsms(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is documentation for AWS CloudHSM Classic. For more information, see [AWS CloudHSM Classic FAQs], the [AWS CloudHSM Classic User Guide]
// , and the [AWS CloudHSM Classic API Reference].
//
// For information about the current version of AWS CloudHSM, see [AWS CloudHSM], the [AWS CloudHSM User Guide], and the [AWS CloudHSM API Reference].
//
// Lists all of the clients.
//
// This operation supports pagination with the use of the NextToken member. If
// more results are available, the NextToken member of the response contains a
// token that you pass in the next call to ListLunaClients to retrieve the next
// set of items.
//
// Deprecated: This API is deprecated.
//
// [AWS CloudHSM User Guide]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/
// [AWS CloudHSM]: http://aws.amazon.com/cloudhsm/
// [AWS CloudHSM API Reference]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/
// [AWS CloudHSM Classic API Reference]: https://docs.aws.amazon.com/cloudhsm/classic/APIReference/
// [AWS CloudHSM Classic FAQs]: http://aws.amazon.com/cloudhsm/faqs-classic/
// [AWS CloudHSM Classic User Guide]: https://docs.aws.amazon.com/cloudhsm/classic/userguide/
func cloudhsm_ListLunaClients(cfg aws.Config, client *cloudhsm.Client) {
	input := &cloudhsm.ListLunaClientsInput{}

	if len(_cloudhsmNextToken) > 0 {
		input.NextToken = aws.String(_cloudhsmNextToken)
	}

	if resp, err := client.ListLunaClients(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is documentation for AWS CloudHSM Classic. For more information, see [AWS CloudHSM Classic FAQs], the [AWS CloudHSM Classic User Guide]
// , and the [AWS CloudHSM Classic API Reference].
//
// For information about the current version of AWS CloudHSM, see [AWS CloudHSM], the [AWS CloudHSM User Guide], and the [AWS CloudHSM API Reference].
//
// Returns a list of all tags for the specified AWS CloudHSM resource.
//
// Deprecated: This API is deprecated.
//
// [AWS CloudHSM User Guide]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/
// [AWS CloudHSM]: http://aws.amazon.com/cloudhsm/
// [AWS CloudHSM API Reference]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/
// [AWS CloudHSM Classic API Reference]: https://docs.aws.amazon.com/cloudhsm/classic/APIReference/
// [AWS CloudHSM Classic FAQs]: http://aws.amazon.com/cloudhsm/faqs-classic/
// [AWS CloudHSM Classic User Guide]: https://docs.aws.amazon.com/cloudhsm/classic/userguide/
func cloudhsm_ListTagsForResource(cfg aws.Config, client *cloudhsm.Client) {
	input := &cloudhsm.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_cloudhsmResourceArn) > 0 {
		input.ResourceArn = aws.String(_cloudhsmResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is documentation for AWS CloudHSM Classic. For more information, see [AWS CloudHSM Classic FAQs], the [AWS CloudHSM Classic User Guide]
// , and the [AWS CloudHSM Classic API Reference].
//
// For information about the current version of AWS CloudHSM, see [AWS CloudHSM], the [AWS CloudHSM User Guide], and the [AWS CloudHSM API Reference].
//
// Modifies an existing high-availability partition group.
//
// Deprecated: This API is deprecated.
//
// [AWS CloudHSM User Guide]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/
// [AWS CloudHSM]: http://aws.amazon.com/cloudhsm/
// [AWS CloudHSM API Reference]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/
// [AWS CloudHSM Classic API Reference]: https://docs.aws.amazon.com/cloudhsm/classic/APIReference/
// [AWS CloudHSM Classic FAQs]: http://aws.amazon.com/cloudhsm/faqs-classic/
// [AWS CloudHSM Classic User Guide]: https://docs.aws.amazon.com/cloudhsm/classic/userguide/
func cloudhsm_ModifyHapg(cfg aws.Config, client *cloudhsm.Client) {
	input := &cloudhsm.ModifyHapgInput{
		// HapgArn: *string, // Required
	}

	if len(_cloudhsmHapgArn) > 0 {
		input.HapgArn = aws.String(_cloudhsmHapgArn)
	}
	if len(_cloudhsmLabel) > 0 {
		input.Label = aws.String(_cloudhsmLabel)
	}
	if len(_cloudhsmPartitionSerialList) > 0 {
		input.PartitionSerialList = append([]string(nil), _cloudhsmPartitionSerialList...)
	}

	if resp, err := client.ModifyHapg(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is documentation for AWS CloudHSM Classic. For more information, see [AWS CloudHSM Classic FAQs], the [AWS CloudHSM Classic User Guide]
// , and the [AWS CloudHSM Classic API Reference].
//
// For information about the current version of AWS CloudHSM, see [AWS CloudHSM], the [AWS CloudHSM User Guide], and the [AWS CloudHSM API Reference].
//
// Modifies an HSM.
//
// This operation can result in the HSM being offline for up to 15 minutes while
// the AWS CloudHSM service is reconfigured. If you are modifying a production HSM,
// you should ensure that your AWS CloudHSM service is configured for high
// availability, and consider executing this operation during a maintenance window.
//
// Deprecated: This API is deprecated.
//
// [AWS CloudHSM User Guide]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/
// [AWS CloudHSM]: http://aws.amazon.com/cloudhsm/
// [AWS CloudHSM API Reference]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/
// [AWS CloudHSM Classic API Reference]: https://docs.aws.amazon.com/cloudhsm/classic/APIReference/
// [AWS CloudHSM Classic FAQs]: http://aws.amazon.com/cloudhsm/faqs-classic/
// [AWS CloudHSM Classic User Guide]: https://docs.aws.amazon.com/cloudhsm/classic/userguide/
func cloudhsm_ModifyHsm(cfg aws.Config, client *cloudhsm.Client) {
	input := &cloudhsm.ModifyHsmInput{
		// HsmArn: *string, // Required
	}

	if len(_cloudhsmHsmArn) > 0 {
		input.HsmArn = aws.String(_cloudhsmHsmArn)
	}
	if len(_cloudhsmEniIp) > 0 {
		input.EniIp = aws.String(_cloudhsmEniIp)
	}
	if len(_cloudhsmExternalId) > 0 {
		input.ExternalId = aws.String(_cloudhsmExternalId)
	}
	if len(_cloudhsmIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_cloudhsmIamRoleArn)
	}
	if len(_cloudhsmSubnetId) > 0 {
		input.SubnetId = aws.String(_cloudhsmSubnetId)
	}
	if len(_cloudhsmSyslogIp) > 0 {
		input.SyslogIp = aws.String(_cloudhsmSyslogIp)
	}

	if resp, err := client.ModifyHsm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is documentation for AWS CloudHSM Classic. For more information, see [AWS CloudHSM Classic FAQs], the [AWS CloudHSM Classic User Guide]
// , and the [AWS CloudHSM Classic API Reference].
//
// For information about the current version of AWS CloudHSM, see [AWS CloudHSM], the [AWS CloudHSM User Guide], and the [AWS CloudHSM API Reference].
//
// Modifies the certificate used by the client.
//
// This action can potentially start a workflow to install the new certificate on
// the client's HSMs.
//
// Deprecated: This API is deprecated.
//
// [AWS CloudHSM User Guide]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/
// [AWS CloudHSM]: http://aws.amazon.com/cloudhsm/
// [AWS CloudHSM API Reference]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/
// [AWS CloudHSM Classic API Reference]: https://docs.aws.amazon.com/cloudhsm/classic/APIReference/
// [AWS CloudHSM Classic FAQs]: http://aws.amazon.com/cloudhsm/faqs-classic/
// [AWS CloudHSM Classic User Guide]: https://docs.aws.amazon.com/cloudhsm/classic/userguide/
func cloudhsm_ModifyLunaClient(cfg aws.Config, client *cloudhsm.Client) {
	input := &cloudhsm.ModifyLunaClientInput{
		// Certificate: *string, // Required
		// ClientArn: *string, // Required
	}

	if len(_cloudhsmCertificate) > 0 {
		input.Certificate = aws.String(_cloudhsmCertificate)
	}
	if len(_cloudhsmClientArn) > 0 {
		input.ClientArn = aws.String(_cloudhsmClientArn)
	}

	if resp, err := client.ModifyLunaClient(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is documentation for AWS CloudHSM Classic. For more information, see [AWS CloudHSM Classic FAQs], the [AWS CloudHSM Classic User Guide]
// , and the [AWS CloudHSM Classic API Reference].
//
// For information about the current version of AWS CloudHSM, see [AWS CloudHSM], the [AWS CloudHSM User Guide], and the [AWS CloudHSM API Reference].
//
// Removes one or more tags from the specified AWS CloudHSM resource.
//
// To remove a tag, specify only the tag key to remove (not the value). To
// overwrite the value for an existing tag, use AddTagsToResource.
//
// Deprecated: This API is deprecated.
//
// [AWS CloudHSM User Guide]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/
// [AWS CloudHSM]: http://aws.amazon.com/cloudhsm/
// [AWS CloudHSM API Reference]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/
// [AWS CloudHSM Classic API Reference]: https://docs.aws.amazon.com/cloudhsm/classic/APIReference/
// [AWS CloudHSM Classic FAQs]: http://aws.amazon.com/cloudhsm/faqs-classic/
// [AWS CloudHSM Classic User Guide]: https://docs.aws.amazon.com/cloudhsm/classic/userguide/
func cloudhsm_RemoveTagsFromResource(cfg aws.Config, client *cloudhsm.Client) {
	input := &cloudhsm.RemoveTagsFromResourceInput{
		// ResourceArn: *string, // Required
		// TagKeyList: []string, // Required
	}

	if len(_cloudhsmResourceArn) > 0 {
		input.ResourceArn = aws.String(_cloudhsmResourceArn)
	}
	if len(_cloudhsmTagKeyList) > 0 {
		input.TagKeyList = append([]string(nil), _cloudhsmTagKeyList...)
	}

	if resp, err := client.RemoveTagsFromResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_cloudhsmCmd)
	_cloudhsmCmd.Flags().SortFlags = false

	_cloudhsmCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_cloudhsmCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_cloudhsmCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_cloudhsmCmd.Flags().StringVarP(&_cloudhsmCertificate, "certificate", "", "", "Certificate")
	_cloudhsmCmd.Flags().StringVarP(&_cloudhsmCertificateFingerprint, "certificate-fingerprint", "", "", "Certificate Fingerprint")
	_cloudhsmCmd.Flags().StringVarP(&_cloudhsmClientArn, "client-arn", "", "", "Client ARN")
	_cloudhsmCmd.Flags().StringVarP(&_cloudhsmClientToken, "client-token", "", "", "Client Token")
	_cloudhsmCmd.Flags().StringVarP(&_cloudhsmClientVersion, "client-version", "", "", "Client Version")
	_cloudhsmCmd.Flags().StringVarP(&_cloudhsmEniIp, "eni-ip", "", "", "Eni IP")
	_cloudhsmCmd.Flags().StringVarP(&_cloudhsmExternalId, "external-id", "", "", "External ID")
	_cloudhsmCmd.Flags().StringVarP(&_cloudhsmHapgArn, "hapg-arn", "", "", "Hapg ARN")
	_cloudhsmCmd.Flags().StringSliceVarP(&_cloudhsmHapgList, "hapg-list", "", nil, "Hapg List")
	_cloudhsmCmd.Flags().StringVarP(&_cloudhsmHsmArn, "hsm-arn", "", "", "Hsm ARN")
	_cloudhsmCmd.Flags().StringVarP(&_cloudhsmHsmSerialNumber, "hsm-serial-number", "", "", "Hsm Serial Number")
	_cloudhsmCmd.Flags().StringVarP(&_cloudhsmIamRoleArn, "iam-role-arn", "", "", "IAM Role ARN")
	_cloudhsmCmd.Flags().StringVarP(&_cloudhsmLabel, "label", "", "", "Label")
	_cloudhsmCmd.Flags().StringVarP(&_cloudhsmNextToken, "next-token", "", "", "Next Token")
	_cloudhsmCmd.Flags().StringSliceVarP(&_cloudhsmPartitionSerialList, "partition-serial-list", "", nil, "Partition Serial List")
	_cloudhsmCmd.Flags().StringVarP(&_cloudhsmResourceArn, "resource-arn", "", "", "Resource ARN")
	_cloudhsmCmd.Flags().StringVarP(&_cloudhsmSshKey, "ssh-key", "", "", "SSH Key")
	_cloudhsmCmd.Flags().StringVarP(&_cloudhsmSubnetId, "subnet-id", "", "", "Subnet ID")
	_cloudhsmCmd.Flags().StringVarP(&_cloudhsmSubscriptionType, "subscription-type", "", "", "Subscription Type")
	_cloudhsmCmd.Flags().StringVarP(&_cloudhsmSyslogIp, "syslog-ip", "", "", "Syslog IP")
	_cloudhsmCmd.Flags().StringSliceVarP(&_cloudhsmTagKeyList, "tag-key-list", "", nil, "Tag Key List")
	_cloudhsmCmd.Flags().StringVarP(&_cloudhsmTagList, "tag-list", "", "", "Tag List")

	_cloudhsmCmd.Flags().BoolVarP(&_cloudhsmAddTagsToResource, "add-tags-to-resource", "", false, "Add Tags To Resource")
	_cloudhsmCmd.Flags().BoolVarP(&_cloudhsmCreateHapg, "create-hapg", "", false, "Create Hapg")
	_cloudhsmCmd.Flags().BoolVarP(&_cloudhsmCreateHsm, "create-hsm", "", false, "Create Hsm")
	_cloudhsmCmd.Flags().BoolVarP(&_cloudhsmCreateLunaClient, "create-luna-client", "", false, "Create Luna Client")
	_cloudhsmCmd.Flags().BoolVarP(&_cloudhsmDeleteHapg, "delete-hapg", "", false, "Delete Hapg")
	_cloudhsmCmd.Flags().BoolVarP(&_cloudhsmDeleteHsm, "delete-hsm", "", false, "Delete Hsm")
	_cloudhsmCmd.Flags().BoolVarP(&_cloudhsmDeleteLunaClient, "delete-luna-client", "", false, "Delete Luna Client")
	_cloudhsmCmd.Flags().BoolVarP(&_cloudhsmDescribeHapg, "describe-hapg", "", false, "Describe Hapg")
	_cloudhsmCmd.Flags().BoolVarP(&_cloudhsmDescribeHsm, "describe-hsm", "", false, "Describe Hsm")
	_cloudhsmCmd.Flags().BoolVarP(&_cloudhsmDescribeLunaClient, "describe-luna-client", "", false, "Describe Luna Client")
	_cloudhsmCmd.Flags().BoolVarP(&_cloudhsmGetConfig, "get-config", "", false, "Get Config")
	_cloudhsmCmd.Flags().BoolVarP(&_cloudhsmListAvailableZones, "list-available-zones", "", false, "List Available Zones")
	_cloudhsmCmd.Flags().BoolVarP(&_cloudhsmListHapgs, "list-hapgs", "", false, "List Hapgs")
	_cloudhsmCmd.Flags().BoolVarP(&_cloudhsmListHsms, "list-hsms", "", false, "List Hsms")
	_cloudhsmCmd.Flags().BoolVarP(&_cloudhsmListLunaClients, "list-luna-clients", "", false, "List Luna Clients")
	_cloudhsmCmd.Flags().BoolVarP(&_cloudhsmListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_cloudhsmCmd.Flags().BoolVarP(&_cloudhsmModifyHapg, "modify-hapg", "", false, "Modify Hapg")
	_cloudhsmCmd.Flags().BoolVarP(&_cloudhsmModifyHsm, "modify-hsm", "", false, "Modify Hsm")
	_cloudhsmCmd.Flags().BoolVarP(&_cloudhsmModifyLunaClient, "modify-luna-client", "", false, "Modify Luna Client")
	_cloudhsmCmd.Flags().BoolVarP(&_cloudhsmRemoveTagsFromResource, "remove-tags-from-resource", "", false, "Remove Tags From Resource")

}
