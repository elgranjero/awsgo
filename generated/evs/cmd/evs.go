package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/evs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// evsCmd represents the evs command
var _evsCmd = &cobra.Command{
	Use:   "evs",
	Short: "AWS evs CLI",
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
		client := evs.NewFromConfig(cfg)
		if _evsAssociateEipToVlan {
			evs_AssociateEipToVlan(cfg, client)
			return
		}
		if _evsCreateEnvironment {
			evs_CreateEnvironment(cfg, client)
			return
		}
		if _evsCreateEnvironmentHost {
			evs_CreateEnvironmentHost(cfg, client)
			return
		}
		if _evsDeleteEnvironment {
			evs_DeleteEnvironment(cfg, client)
			return
		}
		if _evsDeleteEnvironmentHost {
			evs_DeleteEnvironmentHost(cfg, client)
			return
		}
		if _evsDisassociateEipFromVlan {
			evs_DisassociateEipFromVlan(cfg, client)
			return
		}
		if _evsGetEnvironment {
			evs_GetEnvironment(cfg, client)
			return
		}
		if _evsGetVersions {
			evs_GetVersions(cfg, client)
			return
		}
		if _evsListEnvironmentHosts {
			evs_ListEnvironmentHosts(cfg, client)
			return
		}
		if _evsListEnvironmentVlans {
			evs_ListEnvironmentVlans(cfg, client)
			return
		}
		if _evsListEnvironments {
			evs_ListEnvironments(cfg, client)
			return
		}
		if _evsListTagsForResource {
			evs_ListTagsForResource(cfg, client)
			return
		}
		if _evsTagResource {
			evs_TagResource(cfg, client)
			return
		}
		if _evsUntagResource {
			evs_UntagResource(cfg, client)
			return
		}

	},
}

var (
	_evsAssociateEipToVlan      bool
	_evsCreateEnvironment       bool
	_evsCreateEnvironmentHost   bool
	_evsDeleteEnvironment       bool
	_evsDeleteEnvironmentHost   bool
	_evsDisassociateEipFromVlan bool
	_evsGetEnvironment          bool
	_evsGetVersions             bool
	_evsListEnvironmentHosts    bool
	_evsListEnvironmentVlans    bool
	_evsListEnvironments        bool
	_evsListTagsForResource     bool
	_evsTagResource             bool
	_evsUntagResource           bool

	_evsAllocationId                string
	_evsAssociationId               string
	_evsClientToken                 string
	_evsConnectivityInfo            string
	_evsEnvironmentId               string
	_evsEnvironmentName             string
	_evsEsxVersion                  string
	_evsHost                        string
	_evsHostName                    string
	_evsHosts                       string
	_evsInitialVlans                string
	_evsKmsKeyId                    string
	_evsLicenseInfo                 string
	_evsMaxResults                  string
	_evsNextToken                   string
	_evsResourceArn                 string
	_evsServiceAccessSecurityGroups string
	_evsServiceAccessSubnetId       string
	_evsSiteId                      string
	_evsState                       string
	_evsTagKeys                     []string
	_evsTags                        string
	_evsTermsAccepted               string
	_evsVcfHostnames                string
	_evsVcfVersion                  string
	_evsVlanName                    string
	_evsVpcId                       string
)

// Associates an Elastic IP address with a public HCX VLAN. This operation is only
// allowed for public HCX VLANs at this time.
func evs_AssociateEipToVlan(cfg aws.Config, client *evs.Client) {
	input := &evs.AssociateEipToVlanInput{
		// AllocationId: *string, // Required
		// EnvironmentId: *string, // Required
		// VlanName: *string, // Required
	}

	if len(_evsAllocationId) > 0 {
		input.AllocationId = aws.String(_evsAllocationId)
	}
	if len(_evsEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_evsEnvironmentId)
	}
	if len(_evsVlanName) > 0 {
		input.VlanName = aws.String(_evsVlanName)
	}
	if len(_evsClientToken) > 0 {
		input.ClientToken = aws.String(_evsClientToken)
	}

	if resp, err := client.AssociateEipToVlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon EVS environment that runs VCF software, such as SDDC Manager,
// NSX Manager, and vCenter Server.
//
// During environment creation, Amazon EVS performs validations on DNS settings,
// provisions VLAN subnets and hosts, and deploys the supplied version of VCF.
//
// It can take several hours to create an environment. After the deployment
// completes, you can configure VCF in the vSphere user interface according to your
// needs.
//
// When creating a new environment, the default ESX version for the selected VCF
// version will be used, you cannot choose a specific ESX version in
// CreateEnvironment action. When a host has been added with a specific ESX
// version, it can only be upgraded using vCenter Lifecycle Manager.
//
// You cannot use the dedicatedHostId and placementGroupId parameters together in
// the same CreateEnvironment action. This results in a ValidationException
// response.
func evs_CreateEnvironment(cfg aws.Config, client *evs.Client) {
	input := &evs.CreateEnvironmentInput{
		// ConnectivityInfo: *types.ConnectivityInfo, // Required
		// Hosts: []types.HostInfoForCreate, // Required
		// InitialVlans: *types.InitialVlans, // Required
		// LicenseInfo: []types.LicenseInfo, // Required
		// ServiceAccessSubnetId: *string, // Required
		// SiteId: *string, // Required
		// TermsAccepted: *bool, // Required
		// VcfHostnames: *types.VcfHostnames, // Required
		// VcfVersion: types.VcfVersion, // Required
		// VpcId: *string, // Required
	}

	if len(_evsConnectivityInfo) > 0 {
		if err := assignInputField(input, "ConnectivityInfo", _evsConnectivityInfo); err != nil {
			log.Errorf("invalid --connectivity-info: %s", err.Error())
			return
		}
	}
	if len(_evsHosts) > 0 {
		if err := assignInputField(input, "Hosts", _evsHosts); err != nil {
			log.Errorf("invalid --hosts: %s", err.Error())
			return
		}
	}
	if len(_evsInitialVlans) > 0 {
		if err := assignInputField(input, "InitialVlans", _evsInitialVlans); err != nil {
			log.Errorf("invalid --initial-vlans: %s", err.Error())
			return
		}
	}
	if len(_evsLicenseInfo) > 0 {
		if err := assignInputField(input, "LicenseInfo", _evsLicenseInfo); err != nil {
			log.Errorf("invalid --license-info: %s", err.Error())
			return
		}
	}
	if len(_evsServiceAccessSubnetId) > 0 {
		input.ServiceAccessSubnetId = aws.String(_evsServiceAccessSubnetId)
	}
	if len(_evsSiteId) > 0 {
		input.SiteId = aws.String(_evsSiteId)
	}
	if len(_evsTermsAccepted) > 0 {
		if err := assignInputField(input, "TermsAccepted", _evsTermsAccepted); err != nil {
			log.Errorf("invalid --terms-accepted: %s", err.Error())
			return
		}
	}
	if len(_evsVcfHostnames) > 0 {
		if err := assignInputField(input, "VcfHostnames", _evsVcfHostnames); err != nil {
			log.Errorf("invalid --vcf-hostnames: %s", err.Error())
			return
		}
	}
	if len(_evsVcfVersion) > 0 {
		if err := assignInputField(input, "VcfVersion", _evsVcfVersion); err != nil {
			log.Errorf("invalid --vcf-version: %s", err.Error())
			return
		}
	}
	if len(_evsVpcId) > 0 {
		input.VpcId = aws.String(_evsVpcId)
	}
	if len(_evsClientToken) > 0 {
		input.ClientToken = aws.String(_evsClientToken)
	}
	if len(_evsEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_evsEnvironmentName)
	}
	if len(_evsKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_evsKmsKeyId)
	}
	if len(_evsServiceAccessSecurityGroups) > 0 {
		if err := assignInputField(input, "ServiceAccessSecurityGroups", _evsServiceAccessSecurityGroups); err != nil {
			log.Errorf("invalid --service-access-security-groups: %s", err.Error())
			return
		}
	}
	if len(_evsTags) > 0 {
		if err := assignInputField(input, "Tags", _evsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an ESX host and adds it to an Amazon EVS environment. Amazon EVS
// supports 4-16 hosts per environment.
//
// This action can only be used after the Amazon EVS environment is deployed.
//
// You can use the dedicatedHostId parameter to specify an Amazon EC2 Dedicated
// Host for ESX host creation.
//
// You can use the placementGroupId parameter to specify a cluster or partition
// placement group to launch EC2 instances into.
//
// If you don't specify an ESX version when adding hosts using
// CreateEnvironmentHost action, Amazon EVS automatically uses the default ESX
// version associated with your environment's VCF version. To find the default ESX
// version for a particular VCF version, use the GetVersions action.
//
// You cannot use the dedicatedHostId and placementGroupId parameters together in
// the same CreateEnvironmentHost action. This results in a ValidationException
// response.
func evs_CreateEnvironmentHost(cfg aws.Config, client *evs.Client) {
	input := &evs.CreateEnvironmentHostInput{
		// EnvironmentId: *string, // Required
		// Host: *types.HostInfoForCreate, // Required
	}

	if len(_evsEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_evsEnvironmentId)
	}
	if len(_evsHost) > 0 {
		if err := assignInputField(input, "Host", _evsHost); err != nil {
			log.Errorf("invalid --host: %s", err.Error())
			return
		}
	}
	if len(_evsClientToken) > 0 {
		input.ClientToken = aws.String(_evsClientToken)
	}
	if len(_evsEsxVersion) > 0 {
		input.EsxVersion = aws.String(_evsEsxVersion)
	}

	if resp, err := client.CreateEnvironmentHost(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon EVS environment.
// Amazon EVS environments will only be enabled for deletion once the hosts are
// deleted. You can delete hosts using the DeleteEnvironmentHost action.
//
// Environment deletion also deletes the associated Amazon EVS VLAN subnets and
// Amazon Web Services Secrets Manager secrets that Amazon EVS created. Amazon Web
// Services resources that you create are not deleted. These resources may continue
// to incur costs.
func evs_DeleteEnvironment(cfg aws.Config, client *evs.Client) {
	input := &evs.DeleteEnvironmentInput{
		// EnvironmentId: *string, // Required
	}

	if len(_evsEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_evsEnvironmentId)
	}
	if len(_evsClientToken) > 0 {
		input.ClientToken = aws.String(_evsClientToken)
	}

	if resp, err := client.DeleteEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a host from an Amazon EVS environment.
// Before deleting a host, you must unassign and decommission the host from within
// the SDDC Manager user interface. Not doing so could impact the availability of
// your virtual machines or result in data loss.
func evs_DeleteEnvironmentHost(cfg aws.Config, client *evs.Client) {
	input := &evs.DeleteEnvironmentHostInput{
		// EnvironmentId: *string, // Required
		// HostName: *string, // Required
	}

	if len(_evsEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_evsEnvironmentId)
	}
	if len(_evsHostName) > 0 {
		input.HostName = aws.String(_evsHostName)
	}
	if len(_evsClientToken) > 0 {
		input.ClientToken = aws.String(_evsClientToken)
	}

	if resp, err := client.DeleteEnvironmentHost(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates an Elastic IP address from a public HCX VLAN. This operation is
// only allowed for public HCX VLANs at this time.
func evs_DisassociateEipFromVlan(cfg aws.Config, client *evs.Client) {
	input := &evs.DisassociateEipFromVlanInput{
		// AssociationId: *string, // Required
		// EnvironmentId: *string, // Required
		// VlanName: *string, // Required
	}

	if len(_evsAssociationId) > 0 {
		input.AssociationId = aws.String(_evsAssociationId)
	}
	if len(_evsEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_evsEnvironmentId)
	}
	if len(_evsVlanName) > 0 {
		input.VlanName = aws.String(_evsVlanName)
	}
	if len(_evsClientToken) > 0 {
		input.ClientToken = aws.String(_evsClientToken)
	}

	if resp, err := client.DisassociateEipFromVlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a description of the specified environment.
func evs_GetEnvironment(cfg aws.Config, client *evs.Client) {
	input := &evs.GetEnvironmentInput{
		// EnvironmentId: *string, // Required
	}

	if len(_evsEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_evsEnvironmentId)
	}

	if resp, err := client.GetEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about VCF versions, ESX versions and EC2 instance types
// provided by Amazon EVS. For each VCF version, the response also includes the
// default ESX version and provided EC2 instance types.
func evs_GetVersions(cfg aws.Config, client *evs.Client) {
	input := &evs.GetVersionsInput{}

	if resp, err := client.GetVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List the hosts within an environment.
func evs_ListEnvironmentHosts(cfg aws.Config, client *evs.Client) {
	input := &evs.ListEnvironmentHostsInput{
		// EnvironmentId: *string, // Required
	}

	if len(_evsEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_evsEnvironmentId)
	}
	if len(_evsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _evsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_evsNextToken) > 0 {
		input.NextToken = aws.String(_evsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEnvironmentHosts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*evs.ListEnvironmentHostsOutput
	p := evs.NewListEnvironmentHostsPaginator(client, input)
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

// Lists environment VLANs that are associated with the specified environment.
func evs_ListEnvironmentVlans(cfg aws.Config, client *evs.Client) {
	input := &evs.ListEnvironmentVlansInput{
		// EnvironmentId: *string, // Required
	}

	if len(_evsEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_evsEnvironmentId)
	}
	if len(_evsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _evsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_evsNextToken) > 0 {
		input.NextToken = aws.String(_evsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEnvironmentVlans(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*evs.ListEnvironmentVlansOutput
	p := evs.NewListEnvironmentVlansPaginator(client, input)
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

// Lists the Amazon EVS environments in your Amazon Web Services account in the
// specified Amazon Web Services Region.
func evs_ListEnvironments(cfg aws.Config, client *evs.Client) {
	input := &evs.ListEnvironmentsInput{}

	if len(_evsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _evsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_evsNextToken) > 0 {
		input.NextToken = aws.String(_evsNextToken)
	}
	if len(_evsState) > 0 {
		if err := assignInputField(input, "State", _evsState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListEnvironments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*evs.ListEnvironmentsOutput
	p := evs.NewListEnvironmentsPaginator(client, input)
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

// Lists the tags for an Amazon EVS resource.
func evs_ListTagsForResource(cfg aws.Config, client *evs.Client) {
	input := &evs.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_evsResourceArn) > 0 {
		input.ResourceArn = aws.String(_evsResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified tags to an Amazon EVS resource with the specified
// resourceArn . If existing tags on a resource are not specified in the request
// parameters, they aren't changed. When a resource is deleted, the tags associated
// with that resource are also deleted. Tags that you create for Amazon EVS
// resources don't propagate to any other resources associated with the
// environment. For example, if you tag an environment with this operation, that
// tag doesn't automatically propagate to the VLAN subnets and hosts associated
// with the environment.
func evs_TagResource(cfg aws.Config, client *evs.Client) {
	input := &evs.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_evsResourceArn) > 0 {
		input.ResourceArn = aws.String(_evsResourceArn)
	}
	if len(_evsTags) > 0 {
		if err := assignInputField(input, "Tags", _evsTags); err != nil {
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

// Deletes specified tags from an Amazon EVS resource.
func evs_UntagResource(cfg aws.Config, client *evs.Client) {
	input := &evs.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_evsResourceArn) > 0 {
		input.ResourceArn = aws.String(_evsResourceArn)
	}
	if len(_evsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _evsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_evsCmd)
	_evsCmd.Flags().SortFlags = false

	_evsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_evsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_evsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_evsCmd.Flags().StringVarP(&_evsAllocationId, "allocation-id", "", "", "Allocation ID")
	_evsCmd.Flags().StringVarP(&_evsAssociationId, "association-id", "", "", "Association ID")
	_evsCmd.Flags().StringVarP(&_evsClientToken, "client-token", "", "", "Client Token")
	_evsCmd.Flags().StringVarP(&_evsConnectivityInfo, "connectivity-info", "", "", "Connectivity Info")
	_evsCmd.Flags().StringVarP(&_evsEnvironmentId, "environment-id", "", "", "Environment ID")
	_evsCmd.Flags().StringVarP(&_evsEnvironmentName, "environment-name", "", "", "Environment Name")
	_evsCmd.Flags().StringVarP(&_evsEsxVersion, "esx-version", "", "", "Esx Version")
	_evsCmd.Flags().StringVarP(&_evsHost, "host", "", "", "Host")
	_evsCmd.Flags().StringVarP(&_evsHostName, "host-name", "", "", "Host Name")
	_evsCmd.Flags().StringVarP(&_evsHosts, "hosts", "", "", "Hosts")
	_evsCmd.Flags().StringVarP(&_evsInitialVlans, "initial-vlans", "", "", "Initial Vlans")
	_evsCmd.Flags().StringVarP(&_evsKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_evsCmd.Flags().StringVarP(&_evsLicenseInfo, "license-info", "", "", "License Info")
	_evsCmd.Flags().StringVarP(&_evsMaxResults, "max-results", "", "", "Max Results")
	_evsCmd.Flags().StringVarP(&_evsNextToken, "next-token", "", "", "Next Token")
	_evsCmd.Flags().StringVarP(&_evsResourceArn, "resource-arn", "", "", "Resource ARN")
	_evsCmd.Flags().StringVarP(&_evsServiceAccessSecurityGroups, "service-access-security-groups", "", "", "Service Access Security Groups")
	_evsCmd.Flags().StringVarP(&_evsServiceAccessSubnetId, "service-access-subnet-id", "", "", "Service Access Subnet ID")
	_evsCmd.Flags().StringVarP(&_evsSiteId, "site-id", "", "", "Site ID")
	_evsCmd.Flags().StringVarP(&_evsState, "state", "", "", "State")
	_evsCmd.Flags().StringSliceVarP(&_evsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_evsCmd.Flags().StringVarP(&_evsTags, "tags", "", "", "Tags")
	_evsCmd.Flags().StringVarP(&_evsTermsAccepted, "terms-accepted", "", "", "Terms Accepted")
	_evsCmd.Flags().StringVarP(&_evsVcfHostnames, "vcf-hostnames", "", "", "Vcf Hostnames")
	_evsCmd.Flags().StringVarP(&_evsVcfVersion, "vcf-version", "", "", "Vcf Version")
	_evsCmd.Flags().StringVarP(&_evsVlanName, "vlan-name", "", "", "Vlan Name")
	_evsCmd.Flags().StringVarP(&_evsVpcId, "vpc-id", "", "", "VPC ID")

	_evsCmd.Flags().BoolVarP(&_evsAssociateEipToVlan, "associate-eip-to-vlan", "", false, "Associate EIP To Vlan")
	_evsCmd.Flags().BoolVarP(&_evsCreateEnvironment, "create-environment", "", false, "Create Environment")
	_evsCmd.Flags().BoolVarP(&_evsCreateEnvironmentHost, "create-environment-host", "", false, "Create Environment Host")
	_evsCmd.Flags().BoolVarP(&_evsDeleteEnvironment, "delete-environment", "", false, "Delete Environment")
	_evsCmd.Flags().BoolVarP(&_evsDeleteEnvironmentHost, "delete-environment-host", "", false, "Delete Environment Host")
	_evsCmd.Flags().BoolVarP(&_evsDisassociateEipFromVlan, "disassociate-eip-from-vlan", "", false, "Disassociate EIP From Vlan")
	_evsCmd.Flags().BoolVarP(&_evsGetEnvironment, "get-environment", "", false, "Get Environment")
	_evsCmd.Flags().BoolVarP(&_evsGetVersions, "get-versions", "", false, "Get Versions")
	_evsCmd.Flags().BoolVarP(&_evsListEnvironmentHosts, "list-environment-hosts", "", false, "List Environment Hosts")
	_evsCmd.Flags().BoolVarP(&_evsListEnvironmentVlans, "list-environment-vlans", "", false, "List Environment Vlans")
	_evsCmd.Flags().BoolVarP(&_evsListEnvironments, "list-environments", "", false, "List Environments")
	_evsCmd.Flags().BoolVarP(&_evsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_evsCmd.Flags().BoolVarP(&_evsTagResource, "tag-resource", "", false, "Tag Resource")
	_evsCmd.Flags().BoolVarP(&_evsUntagResource, "untag-resource", "", false, "Untag Resource")

}
