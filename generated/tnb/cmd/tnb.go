package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/tnb"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// tnbCmd represents the tnb command
var _tnbCmd = &cobra.Command{
	Use:   "tnb",
	Short: "AWS tnb CLI",
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
		client := tnb.NewFromConfig(cfg)
		if _tnbCancelSolNetworkOperation {
			tnb_CancelSolNetworkOperation(cfg, client)
			return
		}
		if _tnbCreateSolFunctionPackage {
			tnb_CreateSolFunctionPackage(cfg, client)
			return
		}
		if _tnbCreateSolNetworkInstance {
			tnb_CreateSolNetworkInstance(cfg, client)
			return
		}
		if _tnbCreateSolNetworkPackage {
			tnb_CreateSolNetworkPackage(cfg, client)
			return
		}
		if _tnbDeleteSolFunctionPackage {
			tnb_DeleteSolFunctionPackage(cfg, client)
			return
		}
		if _tnbDeleteSolNetworkInstance {
			tnb_DeleteSolNetworkInstance(cfg, client)
			return
		}
		if _tnbDeleteSolNetworkPackage {
			tnb_DeleteSolNetworkPackage(cfg, client)
			return
		}
		if _tnbGetSolFunctionInstance {
			tnb_GetSolFunctionInstance(cfg, client)
			return
		}
		if _tnbGetSolFunctionPackage {
			tnb_GetSolFunctionPackage(cfg, client)
			return
		}
		if _tnbGetSolFunctionPackageContent {
			tnb_GetSolFunctionPackageContent(cfg, client)
			return
		}
		if _tnbGetSolFunctionPackageDescriptor {
			tnb_GetSolFunctionPackageDescriptor(cfg, client)
			return
		}
		if _tnbGetSolNetworkInstance {
			tnb_GetSolNetworkInstance(cfg, client)
			return
		}
		if _tnbGetSolNetworkOperation {
			tnb_GetSolNetworkOperation(cfg, client)
			return
		}
		if _tnbGetSolNetworkPackage {
			tnb_GetSolNetworkPackage(cfg, client)
			return
		}
		if _tnbGetSolNetworkPackageContent {
			tnb_GetSolNetworkPackageContent(cfg, client)
			return
		}
		if _tnbGetSolNetworkPackageDescriptor {
			tnb_GetSolNetworkPackageDescriptor(cfg, client)
			return
		}
		if _tnbInstantiateSolNetworkInstance {
			tnb_InstantiateSolNetworkInstance(cfg, client)
			return
		}
		if _tnbListSolFunctionInstances {
			tnb_ListSolFunctionInstances(cfg, client)
			return
		}
		if _tnbListSolFunctionPackages {
			tnb_ListSolFunctionPackages(cfg, client)
			return
		}
		if _tnbListSolNetworkInstances {
			tnb_ListSolNetworkInstances(cfg, client)
			return
		}
		if _tnbListSolNetworkOperations {
			tnb_ListSolNetworkOperations(cfg, client)
			return
		}
		if _tnbListSolNetworkPackages {
			tnb_ListSolNetworkPackages(cfg, client)
			return
		}
		if _tnbListTagsForResource {
			tnb_ListTagsForResource(cfg, client)
			return
		}
		if _tnbPutSolFunctionPackageContent {
			tnb_PutSolFunctionPackageContent(cfg, client)
			return
		}
		if _tnbPutSolNetworkPackageContent {
			tnb_PutSolNetworkPackageContent(cfg, client)
			return
		}
		if _tnbTagResource {
			tnb_TagResource(cfg, client)
			return
		}
		if _tnbTerminateSolNetworkInstance {
			tnb_TerminateSolNetworkInstance(cfg, client)
			return
		}
		if _tnbUntagResource {
			tnb_UntagResource(cfg, client)
			return
		}
		if _tnbUpdateSolFunctionPackage {
			tnb_UpdateSolFunctionPackage(cfg, client)
			return
		}
		if _tnbUpdateSolNetworkInstance {
			tnb_UpdateSolNetworkInstance(cfg, client)
			return
		}
		if _tnbUpdateSolNetworkPackage {
			tnb_UpdateSolNetworkPackage(cfg, client)
			return
		}
		if _tnbValidateSolFunctionPackageContent {
			tnb_ValidateSolFunctionPackageContent(cfg, client)
			return
		}
		if _tnbValidateSolNetworkPackageContent {
			tnb_ValidateSolNetworkPackageContent(cfg, client)
			return
		}

	},
}

var (
	_tnbCancelSolNetworkOperation         bool
	_tnbCreateSolFunctionPackage          bool
	_tnbCreateSolNetworkInstance          bool
	_tnbCreateSolNetworkPackage           bool
	_tnbDeleteSolFunctionPackage          bool
	_tnbDeleteSolNetworkInstance          bool
	_tnbDeleteSolNetworkPackage           bool
	_tnbGetSolFunctionInstance            bool
	_tnbGetSolFunctionPackage             bool
	_tnbGetSolFunctionPackageContent      bool
	_tnbGetSolFunctionPackageDescriptor   bool
	_tnbGetSolNetworkInstance             bool
	_tnbGetSolNetworkOperation            bool
	_tnbGetSolNetworkPackage              bool
	_tnbGetSolNetworkPackageContent       bool
	_tnbGetSolNetworkPackageDescriptor    bool
	_tnbInstantiateSolNetworkInstance     bool
	_tnbListSolFunctionInstances          bool
	_tnbListSolFunctionPackages           bool
	_tnbListSolNetworkInstances           bool
	_tnbListSolNetworkOperations          bool
	_tnbListSolNetworkPackages            bool
	_tnbListTagsForResource               bool
	_tnbPutSolFunctionPackageContent      bool
	_tnbPutSolNetworkPackageContent       bool
	_tnbTagResource                       bool
	_tnbTerminateSolNetworkInstance       bool
	_tnbUntagResource                     bool
	_tnbUpdateSolFunctionPackage          bool
	_tnbUpdateSolNetworkInstance          bool
	_tnbUpdateSolNetworkPackage           bool
	_tnbValidateSolFunctionPackageContent bool
	_tnbValidateSolNetworkPackageContent  bool

	_tnbAccept                string
	_tnbAdditionalParamsForNs string
	_tnbContentType           string
	_tnbDryRun                string
	_tnbFile                  string
	_tnbMaxResults            string
	_tnbModifyVnfInfoData     string
	_tnbNextToken             string
	_tnbNsDescription         string
	_tnbNsInstanceId          string
	_tnbNsLcmOpOccId          string
	_tnbNsName                string
	_tnbNsdInfoId             string
	_tnbNsdOperationalState   string
	_tnbOperationalState      string
	_tnbResourceArn           string
	_tnbTagKeys               []string
	_tnbTags                  string
	_tnbUpdateNs              string
	_tnbUpdateType            string
	_tnbVnfInstanceId         string
	_tnbVnfPkgId              string
)

// Cancels a network operation.
// A network operation is any operation that is done to your network, such as
// network instance instantiation or termination.
func tnb_CancelSolNetworkOperation(cfg aws.Config, client *tnb.Client) {
	input := &tnb.CancelSolNetworkOperationInput{
		// NsLcmOpOccId: *string, // Required
	}

	if len(_tnbNsLcmOpOccId) > 0 {
		input.NsLcmOpOccId = aws.String(_tnbNsLcmOpOccId)
	}

	if resp, err := client.CancelSolNetworkOperation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a function package.
// A function package is a .zip file in CSAR (Cloud Service Archive) format that
// contains a network function (an ETSI standard telecommunication application) and
// function package descriptor that uses the TOSCA standard to describe how the
// network functions should run on your network. For more information, see [Function packages]in the
// Amazon Web Services Telco Network Builder User Guide.
//
// Creating a function package is the first step for creating a network in AWS
// TNB. This request creates an empty container with an ID. The next step is to
// upload the actual CSAR zip file into that empty container. To upload function
// package content, see [PutSolFunctionPackageContent].
//
// [Function packages]: https://docs.aws.amazon.com/tnb/latest/ug/function-packages.html
// [PutSolFunctionPackageContent]: https://docs.aws.amazon.com/tnb/latest/APIReference/API_PutSolFunctionPackageContent.html
func tnb_CreateSolFunctionPackage(cfg aws.Config, client *tnb.Client) {
	input := &tnb.CreateSolFunctionPackageInput{}

	if len(_tnbTags) > 0 {
		if err := assignInputField(input, "Tags", _tnbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSolFunctionPackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a network instance.
// A network instance is a single network created in Amazon Web Services TNB that
// can be deployed and on which life-cycle operations (like terminate, update, and
// delete) can be performed. Creating a network instance is the third step after
// creating a network package. For more information about network instances, [Network instances]in
// the Amazon Web Services Telco Network Builder User Guide.
//
// Once you create a network instance, you can instantiate it. To instantiate a
// network, see [InstantiateSolNetworkInstance].
//
// [InstantiateSolNetworkInstance]: https://docs.aws.amazon.com/tnb/latest/APIReference/API_InstantiateSolNetworkInstance.html
// [Network instances]: https://docs.aws.amazon.com/tnb/latest/ug/network-instances.html
func tnb_CreateSolNetworkInstance(cfg aws.Config, client *tnb.Client) {
	input := &tnb.CreateSolNetworkInstanceInput{
		// NsName: *string, // Required
		// NsdInfoId: *string, // Required
	}

	if len(_tnbNsName) > 0 {
		input.NsName = aws.String(_tnbNsName)
	}
	if len(_tnbNsdInfoId) > 0 {
		input.NsdInfoId = aws.String(_tnbNsdInfoId)
	}
	if len(_tnbNsDescription) > 0 {
		input.NsDescription = aws.String(_tnbNsDescription)
	}
	if len(_tnbTags) > 0 {
		if err := assignInputField(input, "Tags", _tnbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSolNetworkInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a network package.
// A network package is a .zip file in CSAR (Cloud Service Archive) format defines
// the function packages you want to deploy and the Amazon Web Services
// infrastructure you want to deploy them on. For more information, see [Network instances]in the
// Amazon Web Services Telco Network Builder User Guide.
//
// A network package consists of a network service descriptor (NSD) file
// (required) and any additional files (optional), such as scripts specific to your
// needs. For example, if you have multiple function packages in your network
// package, you can use the NSD to define which network functions should run in
// certain VPCs, subnets, or EKS clusters.
//
// This request creates an empty network package container with an ID. Once you
// create a network package, you can upload the network package content using [PutSolNetworkPackageContent].
//
// [Network instances]: https://docs.aws.amazon.com/tnb/latest/ug/network-instances.html
// [PutSolNetworkPackageContent]: https://docs.aws.amazon.com/tnb/latest/APIReference/API_PutSolNetworkPackageContent.html
func tnb_CreateSolNetworkPackage(cfg aws.Config, client *tnb.Client) {
	input := &tnb.CreateSolNetworkPackageInput{}

	if len(_tnbTags) > 0 {
		if err := assignInputField(input, "Tags", _tnbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSolNetworkPackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a function package.
// A function package is a .zip file in CSAR (Cloud Service Archive) format that
// contains a network function (an ETSI standard telecommunication application) and
// function package descriptor that uses the TOSCA standard to describe how the
// network functions should run on your network.
//
// To delete a function package, the package must be in a disabled state. To
// disable a function package, see [UpdateSolFunctionPackage].
//
// [UpdateSolFunctionPackage]: https://docs.aws.amazon.com/tnb/latest/APIReference/API_UpdateSolFunctionPackage.html
func tnb_DeleteSolFunctionPackage(cfg aws.Config, client *tnb.Client) {
	input := &tnb.DeleteSolFunctionPackageInput{
		// VnfPkgId: *string, // Required
	}

	if len(_tnbVnfPkgId) > 0 {
		input.VnfPkgId = aws.String(_tnbVnfPkgId)
	}

	if resp, err := client.DeleteSolFunctionPackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a network instance.
// A network instance is a single network created in Amazon Web Services TNB that
// can be deployed and on which life-cycle operations (like terminate, update, and
// delete) can be performed.
//
// To delete a network instance, the instance must be in a stopped or terminated
// state. To terminate a network instance, see [TerminateSolNetworkInstance].
//
// [TerminateSolNetworkInstance]: https://docs.aws.amazon.com/tnb/latest/APIReference/API_TerminateSolNetworkInstance.html
func tnb_DeleteSolNetworkInstance(cfg aws.Config, client *tnb.Client) {
	input := &tnb.DeleteSolNetworkInstanceInput{
		// NsInstanceId: *string, // Required
	}

	if len(_tnbNsInstanceId) > 0 {
		input.NsInstanceId = aws.String(_tnbNsInstanceId)
	}

	if resp, err := client.DeleteSolNetworkInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes network package.
// A network package is a .zip file in CSAR (Cloud Service Archive) format defines
// the function packages you want to deploy and the Amazon Web Services
// infrastructure you want to deploy them on.
//
// To delete a network package, the package must be in a disable state. To disable
// a network package, see [UpdateSolNetworkPackage].
//
// [UpdateSolNetworkPackage]: https://docs.aws.amazon.com/tnb/latest/APIReference/API_UpdateSolNetworkPackage.html
func tnb_DeleteSolNetworkPackage(cfg aws.Config, client *tnb.Client) {
	input := &tnb.DeleteSolNetworkPackageInput{
		// NsdInfoId: *string, // Required
	}

	if len(_tnbNsdInfoId) > 0 {
		input.NsdInfoId = aws.String(_tnbNsdInfoId)
	}

	if resp, err := client.DeleteSolNetworkPackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the details of a network function instance, including the instantiation
// state and metadata from the function package descriptor in the network function
// package.
//
// A network function instance is a function in a function package .
func tnb_GetSolFunctionInstance(cfg aws.Config, client *tnb.Client) {
	input := &tnb.GetSolFunctionInstanceInput{
		// VnfInstanceId: *string, // Required
	}

	if len(_tnbVnfInstanceId) > 0 {
		input.VnfInstanceId = aws.String(_tnbVnfInstanceId)
	}

	if resp, err := client.GetSolFunctionInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the details of an individual function package, such as the operational
// state and whether the package is in use.
//
// A function package is a .zip file in CSAR (Cloud Service Archive) format that
// contains a network function (an ETSI standard telecommunication application) and
// function package descriptor that uses the TOSCA standard to describe how the
// network functions should run on your network..
func tnb_GetSolFunctionPackage(cfg aws.Config, client *tnb.Client) {
	input := &tnb.GetSolFunctionPackageInput{
		// VnfPkgId: *string, // Required
	}

	if len(_tnbVnfPkgId) > 0 {
		input.VnfPkgId = aws.String(_tnbVnfPkgId)
	}

	if resp, err := client.GetSolFunctionPackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the contents of a function package.
// A function package is a .zip file in CSAR (Cloud Service Archive) format that
// contains a network function (an ETSI standard telecommunication application) and
// function package descriptor that uses the TOSCA standard to describe how the
// network functions should run on your network.
func tnb_GetSolFunctionPackageContent(cfg aws.Config, client *tnb.Client) {
	input := &tnb.GetSolFunctionPackageContentInput{
		// Accept: types.PackageContentType, // Required
		// VnfPkgId: *string, // Required
	}

	if len(_tnbAccept) > 0 {
		if err := assignInputField(input, "Accept", _tnbAccept); err != nil {
			log.Errorf("invalid --accept: %s", err.Error())
			return
		}
	}
	if len(_tnbVnfPkgId) > 0 {
		input.VnfPkgId = aws.String(_tnbVnfPkgId)
	}

	if resp, err := client.GetSolFunctionPackageContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a function package descriptor in a function package.
// A function package descriptor is a .yaml file in a function package that uses
// the TOSCA standard to describe how the network function in the function package
// should run on your network.
//
// A function package is a .zip file in CSAR (Cloud Service Archive) format that
// contains a network function (an ETSI standard telecommunication application) and
// function package descriptor that uses the TOSCA standard to describe how the
// network functions should run on your network.
func tnb_GetSolFunctionPackageDescriptor(cfg aws.Config, client *tnb.Client) {
	input := &tnb.GetSolFunctionPackageDescriptorInput{
		// Accept: types.DescriptorContentType, // Required
		// VnfPkgId: *string, // Required
	}

	if len(_tnbAccept) > 0 {
		if err := assignInputField(input, "Accept", _tnbAccept); err != nil {
			log.Errorf("invalid --accept: %s", err.Error())
			return
		}
	}
	if len(_tnbVnfPkgId) > 0 {
		input.VnfPkgId = aws.String(_tnbVnfPkgId)
	}

	if resp, err := client.GetSolFunctionPackageDescriptor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the details of the network instance.
// A network instance is a single network created in Amazon Web Services TNB that
// can be deployed and on which life-cycle operations (like terminate, update, and
// delete) can be performed.
func tnb_GetSolNetworkInstance(cfg aws.Config, client *tnb.Client) {
	input := &tnb.GetSolNetworkInstanceInput{
		// NsInstanceId: *string, // Required
	}

	if len(_tnbNsInstanceId) > 0 {
		input.NsInstanceId = aws.String(_tnbNsInstanceId)
	}

	if resp, err := client.GetSolNetworkInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the details of a network operation, including the tasks involved in the
// network operation and the status of the tasks.
//
// A network operation is any operation that is done to your network, such as
// network instance instantiation or termination.
func tnb_GetSolNetworkOperation(cfg aws.Config, client *tnb.Client) {
	input := &tnb.GetSolNetworkOperationInput{
		// NsLcmOpOccId: *string, // Required
	}

	if len(_tnbNsLcmOpOccId) > 0 {
		input.NsLcmOpOccId = aws.String(_tnbNsLcmOpOccId)
	}

	if resp, err := client.GetSolNetworkOperation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the details of a network package.
// A network package is a .zip file in CSAR (Cloud Service Archive) format defines
// the function packages you want to deploy and the Amazon Web Services
// infrastructure you want to deploy them on.
func tnb_GetSolNetworkPackage(cfg aws.Config, client *tnb.Client) {
	input := &tnb.GetSolNetworkPackageInput{
		// NsdInfoId: *string, // Required
	}

	if len(_tnbNsdInfoId) > 0 {
		input.NsdInfoId = aws.String(_tnbNsdInfoId)
	}

	if resp, err := client.GetSolNetworkPackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the contents of a network package.
// A network package is a .zip file in CSAR (Cloud Service Archive) format defines
// the function packages you want to deploy and the Amazon Web Services
// infrastructure you want to deploy them on.
func tnb_GetSolNetworkPackageContent(cfg aws.Config, client *tnb.Client) {
	input := &tnb.GetSolNetworkPackageContentInput{
		// Accept: types.PackageContentType, // Required
		// NsdInfoId: *string, // Required
	}

	if len(_tnbAccept) > 0 {
		if err := assignInputField(input, "Accept", _tnbAccept); err != nil {
			log.Errorf("invalid --accept: %s", err.Error())
			return
		}
	}
	if len(_tnbNsdInfoId) > 0 {
		input.NsdInfoId = aws.String(_tnbNsdInfoId)
	}

	if resp, err := client.GetSolNetworkPackageContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the content of the network service descriptor.
// A network service descriptor is a .yaml file in a network package that uses the
// TOSCA standard to describe the network functions you want to deploy and the
// Amazon Web Services infrastructure you want to deploy the network functions on.
func tnb_GetSolNetworkPackageDescriptor(cfg aws.Config, client *tnb.Client) {
	input := &tnb.GetSolNetworkPackageDescriptorInput{
		// NsdInfoId: *string, // Required
	}

	if len(_tnbNsdInfoId) > 0 {
		input.NsdInfoId = aws.String(_tnbNsdInfoId)
	}

	if resp, err := client.GetSolNetworkPackageDescriptor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Instantiates a network instance.
// A network instance is a single network created in Amazon Web Services TNB that
// can be deployed and on which life-cycle operations (like terminate, update, and
// delete) can be performed.
//
// Before you can instantiate a network instance, you have to create a network
// instance. For more information, see [CreateSolNetworkInstance].
//
// [CreateSolNetworkInstance]: https://docs.aws.amazon.com/tnb/latest/APIReference/API_CreateSolNetworkInstance.html
func tnb_InstantiateSolNetworkInstance(cfg aws.Config, client *tnb.Client) {
	input := &tnb.InstantiateSolNetworkInstanceInput{
		// NsInstanceId: *string, // Required
	}

	if len(_tnbNsInstanceId) > 0 {
		input.NsInstanceId = aws.String(_tnbNsInstanceId)
	}
	if len(_tnbAdditionalParamsForNs) > 0 {
		if err := assignInputField(input, "AdditionalParamsForNs", _tnbAdditionalParamsForNs); err != nil {
			log.Errorf("invalid --additional-params-for-ns: %s", err.Error())
			return
		}
	}
	if len(_tnbDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _tnbDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_tnbTags) > 0 {
		if err := assignInputField(input, "Tags", _tnbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.InstantiateSolNetworkInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists network function instances.
// A network function instance is a function in a function package .
func tnb_ListSolFunctionInstances(cfg aws.Config, client *tnb.Client) {
	input := &tnb.ListSolFunctionInstancesInput{}

	if len(_tnbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _tnbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_tnbNextToken) > 0 {
		input.NextToken = aws.String(_tnbNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSolFunctionInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*tnb.ListSolFunctionInstancesOutput
	p := tnb.NewListSolFunctionInstancesPaginator(client, input)
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

// Lists information about function packages.
// A function package is a .zip file in CSAR (Cloud Service Archive) format that
// contains a network function (an ETSI standard telecommunication application) and
// function package descriptor that uses the TOSCA standard to describe how the
// network functions should run on your network.
func tnb_ListSolFunctionPackages(cfg aws.Config, client *tnb.Client) {
	input := &tnb.ListSolFunctionPackagesInput{}

	if len(_tnbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _tnbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_tnbNextToken) > 0 {
		input.NextToken = aws.String(_tnbNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSolFunctionPackages(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*tnb.ListSolFunctionPackagesOutput
	p := tnb.NewListSolFunctionPackagesPaginator(client, input)
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

// Lists your network instances.
// A network instance is a single network created in Amazon Web Services TNB that
// can be deployed and on which life-cycle operations (like terminate, update, and
// delete) can be performed.
func tnb_ListSolNetworkInstances(cfg aws.Config, client *tnb.Client) {
	input := &tnb.ListSolNetworkInstancesInput{}

	if len(_tnbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _tnbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_tnbNextToken) > 0 {
		input.NextToken = aws.String(_tnbNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSolNetworkInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*tnb.ListSolNetworkInstancesOutput
	p := tnb.NewListSolNetworkInstancesPaginator(client, input)
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

// Lists details for a network operation, including when the operation started and
// the status of the operation.
//
// A network operation is any operation that is done to your network, such as
// network instance instantiation or termination.
func tnb_ListSolNetworkOperations(cfg aws.Config, client *tnb.Client) {
	input := &tnb.ListSolNetworkOperationsInput{}

	if len(_tnbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _tnbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_tnbNextToken) > 0 {
		input.NextToken = aws.String(_tnbNextToken)
	}
	if len(_tnbNsInstanceId) > 0 {
		input.NsInstanceId = aws.String(_tnbNsInstanceId)
	}

	if disablePaginator() {
		if resp, err := client.ListSolNetworkOperations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*tnb.ListSolNetworkOperationsOutput
	p := tnb.NewListSolNetworkOperationsPaginator(client, input)
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

// Lists network packages.
// A network package is a .zip file in CSAR (Cloud Service Archive) format defines
// the function packages you want to deploy and the Amazon Web Services
// infrastructure you want to deploy them on.
func tnb_ListSolNetworkPackages(cfg aws.Config, client *tnb.Client) {
	input := &tnb.ListSolNetworkPackagesInput{}

	if len(_tnbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _tnbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_tnbNextToken) > 0 {
		input.NextToken = aws.String(_tnbNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSolNetworkPackages(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*tnb.ListSolNetworkPackagesOutput
	p := tnb.NewListSolNetworkPackagesPaginator(client, input)
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

// Lists tags for AWS TNB resources.
func tnb_ListTagsForResource(cfg aws.Config, client *tnb.Client) {
	input := &tnb.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_tnbResourceArn) > 0 {
		input.ResourceArn = aws.String(_tnbResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Uploads the contents of a function package.
// A function package is a .zip file in CSAR (Cloud Service Archive) format that
// contains a network function (an ETSI standard telecommunication application) and
// function package descriptor that uses the TOSCA standard to describe how the
// network functions should run on your network.
func tnb_PutSolFunctionPackageContent(cfg aws.Config, client *tnb.Client) {
	input := &tnb.PutSolFunctionPackageContentInput{
		// File: []byte, // Required
		// VnfPkgId: *string, // Required
	}

	if len(_tnbFile) > 0 {
		if err := assignInputField(input, "File", _tnbFile); err != nil {
			log.Errorf("invalid --file: %s", err.Error())
			return
		}
	}
	if len(_tnbVnfPkgId) > 0 {
		input.VnfPkgId = aws.String(_tnbVnfPkgId)
	}
	if len(_tnbContentType) > 0 {
		if err := assignInputField(input, "ContentType", _tnbContentType); err != nil {
			log.Errorf("invalid --content-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutSolFunctionPackageContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Uploads the contents of a network package.
// A network package is a .zip file in CSAR (Cloud Service Archive) format defines
// the function packages you want to deploy and the Amazon Web Services
// infrastructure you want to deploy them on.
func tnb_PutSolNetworkPackageContent(cfg aws.Config, client *tnb.Client) {
	input := &tnb.PutSolNetworkPackageContentInput{
		// File: []byte, // Required
		// NsdInfoId: *string, // Required
	}

	if len(_tnbFile) > 0 {
		if err := assignInputField(input, "File", _tnbFile); err != nil {
			log.Errorf("invalid --file: %s", err.Error())
			return
		}
	}
	if len(_tnbNsdInfoId) > 0 {
		input.NsdInfoId = aws.String(_tnbNsdInfoId)
	}
	if len(_tnbContentType) > 0 {
		if err := assignInputField(input, "ContentType", _tnbContentType); err != nil {
			log.Errorf("invalid --content-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutSolNetworkPackageContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tags an AWS TNB resource.
// A tag is a label that you assign to an Amazon Web Services resource. Each tag
// consists of a key and an optional value. You can use tags to search and filter
// your resources or track your Amazon Web Services costs.
func tnb_TagResource(cfg aws.Config, client *tnb.Client) {
	input := &tnb.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_tnbResourceArn) > 0 {
		input.ResourceArn = aws.String(_tnbResourceArn)
	}
	if len(_tnbTags) > 0 {
		if err := assignInputField(input, "Tags", _tnbTags); err != nil {
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

// Terminates a network instance.
// A network instance is a single network created in Amazon Web Services TNB that
// can be deployed and on which life-cycle operations (like terminate, update, and
// delete) can be performed.
//
// You must terminate a network instance before you can delete it.
func tnb_TerminateSolNetworkInstance(cfg aws.Config, client *tnb.Client) {
	input := &tnb.TerminateSolNetworkInstanceInput{
		// NsInstanceId: *string, // Required
	}

	if len(_tnbNsInstanceId) > 0 {
		input.NsInstanceId = aws.String(_tnbNsInstanceId)
	}
	if len(_tnbTags) > 0 {
		if err := assignInputField(input, "Tags", _tnbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TerminateSolNetworkInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Untags an AWS TNB resource.
// A tag is a label that you assign to an Amazon Web Services resource. Each tag
// consists of a key and an optional value. You can use tags to search and filter
// your resources or track your Amazon Web Services costs.
func tnb_UntagResource(cfg aws.Config, client *tnb.Client) {
	input := &tnb.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_tnbResourceArn) > 0 {
		input.ResourceArn = aws.String(_tnbResourceArn)
	}
	if len(_tnbTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _tnbTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the operational state of function package.
// A function package is a .zip file in CSAR (Cloud Service Archive) format that
// contains a network function (an ETSI standard telecommunication application) and
// function package descriptor that uses the TOSCA standard to describe how the
// network functions should run on your network.
func tnb_UpdateSolFunctionPackage(cfg aws.Config, client *tnb.Client) {
	input := &tnb.UpdateSolFunctionPackageInput{
		// OperationalState: types.OperationalState, // Required
		// VnfPkgId: *string, // Required
	}

	if len(_tnbOperationalState) > 0 {
		if err := assignInputField(input, "OperationalState", _tnbOperationalState); err != nil {
			log.Errorf("invalid --operational-state: %s", err.Error())
			return
		}
	}
	if len(_tnbVnfPkgId) > 0 {
		input.VnfPkgId = aws.String(_tnbVnfPkgId)
	}

	if resp, err := client.UpdateSolFunctionPackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a network instance.
// A network instance is a single network created in Amazon Web Services TNB that
// can be deployed and on which life-cycle operations (like terminate, update, and
// delete) can be performed.
//
// Choose the updateType parameter to target the necessary update of the network
// instance.
func tnb_UpdateSolNetworkInstance(cfg aws.Config, client *tnb.Client) {
	input := &tnb.UpdateSolNetworkInstanceInput{
		// NsInstanceId: *string, // Required
		// UpdateType: types.UpdateSolNetworkType, // Required
	}

	if len(_tnbNsInstanceId) > 0 {
		input.NsInstanceId = aws.String(_tnbNsInstanceId)
	}
	if len(_tnbUpdateType) > 0 {
		if err := assignInputField(input, "UpdateType", _tnbUpdateType); err != nil {
			log.Errorf("invalid --update-type: %s", err.Error())
			return
		}
	}
	if len(_tnbModifyVnfInfoData) > 0 {
		if err := assignInputField(input, "ModifyVnfInfoData", _tnbModifyVnfInfoData); err != nil {
			log.Errorf("invalid --modify-vnf-info-data: %s", err.Error())
			return
		}
	}
	if len(_tnbTags) > 0 {
		if err := assignInputField(input, "Tags", _tnbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_tnbUpdateNs) > 0 {
		if err := assignInputField(input, "UpdateNs", _tnbUpdateNs); err != nil {
			log.Errorf("invalid --update-ns: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSolNetworkInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the operational state of a network package.
// A network package is a .zip file in CSAR (Cloud Service Archive) format defines
// the function packages you want to deploy and the Amazon Web Services
// infrastructure you want to deploy them on.
//
// A network service descriptor is a .yaml file in a network package that uses the
// TOSCA standard to describe the network functions you want to deploy and the
// Amazon Web Services infrastructure you want to deploy the network functions on.
func tnb_UpdateSolNetworkPackage(cfg aws.Config, client *tnb.Client) {
	input := &tnb.UpdateSolNetworkPackageInput{
		// NsdInfoId: *string, // Required
		// NsdOperationalState: types.NsdOperationalState, // Required
	}

	if len(_tnbNsdInfoId) > 0 {
		input.NsdInfoId = aws.String(_tnbNsdInfoId)
	}
	if len(_tnbNsdOperationalState) > 0 {
		if err := assignInputField(input, "NsdOperationalState", _tnbNsdOperationalState); err != nil {
			log.Errorf("invalid --nsd-operational-state: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSolNetworkPackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Validates function package content. This can be used as a dry run before
// uploading function package content with [PutSolFunctionPackageContent].
//
// A function package is a .zip file in CSAR (Cloud Service Archive) format that
// contains a network function (an ETSI standard telecommunication application) and
// function package descriptor that uses the TOSCA standard to describe how the
// network functions should run on your network.
//
// [PutSolFunctionPackageContent]: https://docs.aws.amazon.com/tnb/latest/APIReference/API_PutSolFunctionPackageContent.html
func tnb_ValidateSolFunctionPackageContent(cfg aws.Config, client *tnb.Client) {
	input := &tnb.ValidateSolFunctionPackageContentInput{
		// File: []byte, // Required
		// VnfPkgId: *string, // Required
	}

	if len(_tnbFile) > 0 {
		if err := assignInputField(input, "File", _tnbFile); err != nil {
			log.Errorf("invalid --file: %s", err.Error())
			return
		}
	}
	if len(_tnbVnfPkgId) > 0 {
		input.VnfPkgId = aws.String(_tnbVnfPkgId)
	}
	if len(_tnbContentType) > 0 {
		if err := assignInputField(input, "ContentType", _tnbContentType); err != nil {
			log.Errorf("invalid --content-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.ValidateSolFunctionPackageContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Validates network package content. This can be used as a dry run before
// uploading network package content with [PutSolNetworkPackageContent].
//
// A network package is a .zip file in CSAR (Cloud Service Archive) format defines
// the function packages you want to deploy and the Amazon Web Services
// infrastructure you want to deploy them on.
//
// [PutSolNetworkPackageContent]: https://docs.aws.amazon.com/tnb/latest/APIReference/API_PutSolNetworkPackageContent.html
func tnb_ValidateSolNetworkPackageContent(cfg aws.Config, client *tnb.Client) {
	input := &tnb.ValidateSolNetworkPackageContentInput{
		// File: []byte, // Required
		// NsdInfoId: *string, // Required
	}

	if len(_tnbFile) > 0 {
		if err := assignInputField(input, "File", _tnbFile); err != nil {
			log.Errorf("invalid --file: %s", err.Error())
			return
		}
	}
	if len(_tnbNsdInfoId) > 0 {
		input.NsdInfoId = aws.String(_tnbNsdInfoId)
	}
	if len(_tnbContentType) > 0 {
		if err := assignInputField(input, "ContentType", _tnbContentType); err != nil {
			log.Errorf("invalid --content-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.ValidateSolNetworkPackageContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_tnbCmd)
	_tnbCmd.Flags().SortFlags = false

	_tnbCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_tnbCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_tnbCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_tnbCmd.Flags().StringVarP(&_tnbAccept, "accept", "", "", "Accept")
	_tnbCmd.Flags().StringVarP(&_tnbAdditionalParamsForNs, "additional-params-for-ns", "", "", "Additional Params For Ns")
	_tnbCmd.Flags().StringVarP(&_tnbContentType, "content-type", "", "", "Content Type")
	_tnbCmd.Flags().StringVarP(&_tnbDryRun, "dry-run", "", "", "Dry Run")
	_tnbCmd.Flags().StringVarP(&_tnbFile, "file", "", "", "File")
	_tnbCmd.Flags().StringVarP(&_tnbMaxResults, "max-results", "", "", "Max Results")
	_tnbCmd.Flags().StringVarP(&_tnbModifyVnfInfoData, "modify-vnf-info-data", "", "", "Modify Vnf Info Data")
	_tnbCmd.Flags().StringVarP(&_tnbNextToken, "next-token", "", "", "Next Token")
	_tnbCmd.Flags().StringVarP(&_tnbNsDescription, "ns-description", "", "", "Ns Description")
	_tnbCmd.Flags().StringVarP(&_tnbNsInstanceId, "ns-instance-id", "", "", "Ns Instance ID")
	_tnbCmd.Flags().StringVarP(&_tnbNsLcmOpOccId, "ns-lcm-op-occ-id", "", "", "Ns Lcm Op Occ ID")
	_tnbCmd.Flags().StringVarP(&_tnbNsName, "ns-name", "", "", "Ns Name")
	_tnbCmd.Flags().StringVarP(&_tnbNsdInfoId, "nsd-info-id", "", "", "Nsd Info ID")
	_tnbCmd.Flags().StringVarP(&_tnbNsdOperationalState, "nsd-operational-state", "", "", "Nsd Operational State")
	_tnbCmd.Flags().StringVarP(&_tnbOperationalState, "operational-state", "", "", "Operational State")
	_tnbCmd.Flags().StringVarP(&_tnbResourceArn, "resource-arn", "", "", "Resource ARN")
	_tnbCmd.Flags().StringSliceVarP(&_tnbTagKeys, "tag-keys", "", nil, "Tag Keys")
	_tnbCmd.Flags().StringVarP(&_tnbTags, "tags", "", "", "Tags")
	_tnbCmd.Flags().StringVarP(&_tnbUpdateNs, "update-ns", "", "", "Update Ns")
	_tnbCmd.Flags().StringVarP(&_tnbUpdateType, "update-type", "", "", "Update Type")
	_tnbCmd.Flags().StringVarP(&_tnbVnfInstanceId, "vnf-instance-id", "", "", "Vnf Instance ID")
	_tnbCmd.Flags().StringVarP(&_tnbVnfPkgId, "vnf-pkg-id", "", "", "Vnf Pkg ID")

	_tnbCmd.Flags().BoolVarP(&_tnbCancelSolNetworkOperation, "cancel-sol-network-operation", "", false, "Cancel Sol Network Operation")
	_tnbCmd.Flags().BoolVarP(&_tnbCreateSolFunctionPackage, "create-sol-function-package", "", false, "Create Sol Function Package")
	_tnbCmd.Flags().BoolVarP(&_tnbCreateSolNetworkInstance, "create-sol-network-instance", "", false, "Create Sol Network Instance")
	_tnbCmd.Flags().BoolVarP(&_tnbCreateSolNetworkPackage, "create-sol-network-package", "", false, "Create Sol Network Package")
	_tnbCmd.Flags().BoolVarP(&_tnbDeleteSolFunctionPackage, "delete-sol-function-package", "", false, "Delete Sol Function Package")
	_tnbCmd.Flags().BoolVarP(&_tnbDeleteSolNetworkInstance, "delete-sol-network-instance", "", false, "Delete Sol Network Instance")
	_tnbCmd.Flags().BoolVarP(&_tnbDeleteSolNetworkPackage, "delete-sol-network-package", "", false, "Delete Sol Network Package")
	_tnbCmd.Flags().BoolVarP(&_tnbGetSolFunctionInstance, "get-sol-function-instance", "", false, "Get Sol Function Instance")
	_tnbCmd.Flags().BoolVarP(&_tnbGetSolFunctionPackage, "get-sol-function-package", "", false, "Get Sol Function Package")
	_tnbCmd.Flags().BoolVarP(&_tnbGetSolFunctionPackageContent, "get-sol-function-package-content", "", false, "Get Sol Function Package Content")
	_tnbCmd.Flags().BoolVarP(&_tnbGetSolFunctionPackageDescriptor, "get-sol-function-package-descriptor", "", false, "Get Sol Function Package Descriptor")
	_tnbCmd.Flags().BoolVarP(&_tnbGetSolNetworkInstance, "get-sol-network-instance", "", false, "Get Sol Network Instance")
	_tnbCmd.Flags().BoolVarP(&_tnbGetSolNetworkOperation, "get-sol-network-operation", "", false, "Get Sol Network Operation")
	_tnbCmd.Flags().BoolVarP(&_tnbGetSolNetworkPackage, "get-sol-network-package", "", false, "Get Sol Network Package")
	_tnbCmd.Flags().BoolVarP(&_tnbGetSolNetworkPackageContent, "get-sol-network-package-content", "", false, "Get Sol Network Package Content")
	_tnbCmd.Flags().BoolVarP(&_tnbGetSolNetworkPackageDescriptor, "get-sol-network-package-descriptor", "", false, "Get Sol Network Package Descriptor")
	_tnbCmd.Flags().BoolVarP(&_tnbInstantiateSolNetworkInstance, "instantiate-sol-network-instance", "", false, "Instantiate Sol Network Instance")
	_tnbCmd.Flags().BoolVarP(&_tnbListSolFunctionInstances, "list-sol-function-instances", "", false, "List Sol Function Instances")
	_tnbCmd.Flags().BoolVarP(&_tnbListSolFunctionPackages, "list-sol-function-packages", "", false, "List Sol Function Packages")
	_tnbCmd.Flags().BoolVarP(&_tnbListSolNetworkInstances, "list-sol-network-instances", "", false, "List Sol Network Instances")
	_tnbCmd.Flags().BoolVarP(&_tnbListSolNetworkOperations, "list-sol-network-operations", "", false, "List Sol Network Operations")
	_tnbCmd.Flags().BoolVarP(&_tnbListSolNetworkPackages, "list-sol-network-packages", "", false, "List Sol Network Packages")
	_tnbCmd.Flags().BoolVarP(&_tnbListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_tnbCmd.Flags().BoolVarP(&_tnbPutSolFunctionPackageContent, "put-sol-function-package-content", "", false, "Put Sol Function Package Content")
	_tnbCmd.Flags().BoolVarP(&_tnbPutSolNetworkPackageContent, "put-sol-network-package-content", "", false, "Put Sol Network Package Content")
	_tnbCmd.Flags().BoolVarP(&_tnbTagResource, "tag-resource", "", false, "Tag Resource")
	_tnbCmd.Flags().BoolVarP(&_tnbTerminateSolNetworkInstance, "terminate-sol-network-instance", "", false, "Terminate Sol Network Instance")
	_tnbCmd.Flags().BoolVarP(&_tnbUntagResource, "untag-resource", "", false, "Untag Resource")
	_tnbCmd.Flags().BoolVarP(&_tnbUpdateSolFunctionPackage, "update-sol-function-package", "", false, "Update Sol Function Package")
	_tnbCmd.Flags().BoolVarP(&_tnbUpdateSolNetworkInstance, "update-sol-network-instance", "", false, "Update Sol Network Instance")
	_tnbCmd.Flags().BoolVarP(&_tnbUpdateSolNetworkPackage, "update-sol-network-package", "", false, "Update Sol Network Package")
	_tnbCmd.Flags().BoolVarP(&_tnbValidateSolFunctionPackageContent, "validate-sol-function-package-content", "", false, "Validate Sol Function Package Content")
	_tnbCmd.Flags().BoolVarP(&_tnbValidateSolNetworkPackageContent, "validate-sol-network-package-content", "", false, "Validate Sol Network Package Content")

}
