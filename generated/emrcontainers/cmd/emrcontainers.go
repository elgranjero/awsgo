package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/emrcontainers"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// emrcontainersCmd represents the emrcontainers command
var _emrcontainersCmd = &cobra.Command{
	Use:   "emrcontainers",
	Short: "AWS emrcontainers CLI",
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
		client := emrcontainers.NewFromConfig(cfg)
		if _emrcontainersCancelJobRun {
			emrcontainers_CancelJobRun(cfg, client)
			return
		}
		if _emrcontainersCreateJobTemplate {
			emrcontainers_CreateJobTemplate(cfg, client)
			return
		}
		if _emrcontainersCreateManagedEndpoint {
			emrcontainers_CreateManagedEndpoint(cfg, client)
			return
		}
		if _emrcontainersCreateSecurityConfiguration {
			emrcontainers_CreateSecurityConfiguration(cfg, client)
			return
		}
		if _emrcontainersCreateVirtualCluster {
			emrcontainers_CreateVirtualCluster(cfg, client)
			return
		}
		if _emrcontainersDeleteJobTemplate {
			emrcontainers_DeleteJobTemplate(cfg, client)
			return
		}
		if _emrcontainersDeleteManagedEndpoint {
			emrcontainers_DeleteManagedEndpoint(cfg, client)
			return
		}
		if _emrcontainersDeleteVirtualCluster {
			emrcontainers_DeleteVirtualCluster(cfg, client)
			return
		}
		if _emrcontainersDescribeJobRun {
			emrcontainers_DescribeJobRun(cfg, client)
			return
		}
		if _emrcontainersDescribeJobTemplate {
			emrcontainers_DescribeJobTemplate(cfg, client)
			return
		}
		if _emrcontainersDescribeManagedEndpoint {
			emrcontainers_DescribeManagedEndpoint(cfg, client)
			return
		}
		if _emrcontainersDescribeSecurityConfiguration {
			emrcontainers_DescribeSecurityConfiguration(cfg, client)
			return
		}
		if _emrcontainersDescribeVirtualCluster {
			emrcontainers_DescribeVirtualCluster(cfg, client)
			return
		}
		if _emrcontainersGetManagedEndpointSessionCredentials {
			emrcontainers_GetManagedEndpointSessionCredentials(cfg, client)
			return
		}
		if _emrcontainersListJobRuns {
			emrcontainers_ListJobRuns(cfg, client)
			return
		}
		if _emrcontainersListJobTemplates {
			emrcontainers_ListJobTemplates(cfg, client)
			return
		}
		if _emrcontainersListManagedEndpoints {
			emrcontainers_ListManagedEndpoints(cfg, client)
			return
		}
		if _emrcontainersListSecurityConfigurations {
			emrcontainers_ListSecurityConfigurations(cfg, client)
			return
		}
		if _emrcontainersListTagsForResource {
			emrcontainers_ListTagsForResource(cfg, client)
			return
		}
		if _emrcontainersListVirtualClusters {
			emrcontainers_ListVirtualClusters(cfg, client)
			return
		}
		if _emrcontainersStartJobRun {
			emrcontainers_StartJobRun(cfg, client)
			return
		}
		if _emrcontainersTagResource {
			emrcontainers_TagResource(cfg, client)
			return
		}
		if _emrcontainersUntagResource {
			emrcontainers_UntagResource(cfg, client)
			return
		}

	},
}

var (
	_emrcontainersCancelJobRun                         bool
	_emrcontainersCreateJobTemplate                    bool
	_emrcontainersCreateManagedEndpoint                bool
	_emrcontainersCreateSecurityConfiguration          bool
	_emrcontainersCreateVirtualCluster                 bool
	_emrcontainersDeleteJobTemplate                    bool
	_emrcontainersDeleteManagedEndpoint                bool
	_emrcontainersDeleteVirtualCluster                 bool
	_emrcontainersDescribeJobRun                       bool
	_emrcontainersDescribeJobTemplate                  bool
	_emrcontainersDescribeManagedEndpoint              bool
	_emrcontainersDescribeSecurityConfiguration        bool
	_emrcontainersDescribeVirtualCluster               bool
	_emrcontainersGetManagedEndpointSessionCredentials bool
	_emrcontainersListJobRuns                          bool
	_emrcontainersListJobTemplates                     bool
	_emrcontainersListManagedEndpoints                 bool
	_emrcontainersListSecurityConfigurations           bool
	_emrcontainersListTagsForResource                  bool
	_emrcontainersListVirtualClusters                  bool
	_emrcontainersStartJobRun                          bool
	_emrcontainersTagResource                          bool
	_emrcontainersUntagResource                        bool

	_emrcontainersCertificateArn            string
	_emrcontainersClientToken               string
	_emrcontainersConfigurationOverrides    string
	_emrcontainersContainerProvider         string
	_emrcontainersContainerProviderId       string
	_emrcontainersContainerProviderType     string
	_emrcontainersCreatedAfter              string
	_emrcontainersCreatedBefore             string
	_emrcontainersCredentialType            string
	_emrcontainersDurationInSeconds         string
	_emrcontainersEksAccessEntryIntegrated  string
	_emrcontainersEndpointIdentifier        string
	_emrcontainersExecutionRoleArn          string
	_emrcontainersId                        string
	_emrcontainersJobDriver                 string
	_emrcontainersJobTemplateData           string
	_emrcontainersJobTemplateId             string
	_emrcontainersJobTemplateParameters     string
	_emrcontainersKmsKeyArn                 string
	_emrcontainersLogContext                string
	_emrcontainersMaxResults                string
	_emrcontainersName                      string
	_emrcontainersNextToken                 string
	_emrcontainersReleaseLabel              string
	_emrcontainersResourceArn               string
	_emrcontainersRetryPolicyConfiguration  string
	_emrcontainersSecurityConfigurationData string
	_emrcontainersSecurityConfigurationId   string
	_emrcontainersStates                    string
	_emrcontainersTagKeys                   []string
	_emrcontainersTags                      string
	_emrcontainersType                      string
	_emrcontainersTypes                     []string
	_emrcontainersVirtualClusterId          string
	_emrcontainersVirtualClusterIdentifier  string
)

// Cancels a job run. A job run is a unit of work, such as a Spark jar, PySpark
// script, or SparkSQL query, that you submit to Amazon EMR on EKS.
func emrcontainers_CancelJobRun(cfg aws.Config, client *emrcontainers.Client) {
	input := &emrcontainers.CancelJobRunInput{
		// Id: *string, // Required
		// VirtualClusterId: *string, // Required
	}

	if len(_emrcontainersId) > 0 {
		input.Id = aws.String(_emrcontainersId)
	}
	if len(_emrcontainersVirtualClusterId) > 0 {
		input.VirtualClusterId = aws.String(_emrcontainersVirtualClusterId)
	}

	if resp, err := client.CancelJobRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a job template. Job template stores values of StartJobRun API request
// in a template and can be used to start a job run. Job template allows two use
// cases: avoid repeating recurring StartJobRun API request values, enforcing
// certain values in StartJobRun API request.
func emrcontainers_CreateJobTemplate(cfg aws.Config, client *emrcontainers.Client) {
	input := &emrcontainers.CreateJobTemplateInput{
		// ClientToken: *string, // Required
		// JobTemplateData: *types.JobTemplateData, // Required
		// Name: *string, // Required
	}

	if len(_emrcontainersClientToken) > 0 {
		input.ClientToken = aws.String(_emrcontainersClientToken)
	}
	if len(_emrcontainersJobTemplateData) > 0 {
		if err := assignInputField(input, "JobTemplateData", _emrcontainersJobTemplateData); err != nil {
			log.Errorf("invalid --job-template-data: %s", err.Error())
			return
		}
	}
	if len(_emrcontainersName) > 0 {
		input.Name = aws.String(_emrcontainersName)
	}
	if len(_emrcontainersKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_emrcontainersKmsKeyArn)
	}
	if len(_emrcontainersTags) > 0 {
		if err := assignInputField(input, "Tags", _emrcontainersTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateJobTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a managed endpoint. A managed endpoint is a gateway that connects
// Amazon EMR Studio to Amazon EMR on EKS so that Amazon EMR Studio can communicate
// with your virtual cluster.
func emrcontainers_CreateManagedEndpoint(cfg aws.Config, client *emrcontainers.Client) {
	input := &emrcontainers.CreateManagedEndpointInput{
		// ClientToken: *string, // Required
		// ExecutionRoleArn: *string, // Required
		// Name: *string, // Required
		// ReleaseLabel: *string, // Required
		// Type: *string, // Required
		// VirtualClusterId: *string, // Required
	}

	if len(_emrcontainersClientToken) > 0 {
		input.ClientToken = aws.String(_emrcontainersClientToken)
	}
	if len(_emrcontainersExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_emrcontainersExecutionRoleArn)
	}
	if len(_emrcontainersName) > 0 {
		input.Name = aws.String(_emrcontainersName)
	}
	if len(_emrcontainersReleaseLabel) > 0 {
		input.ReleaseLabel = aws.String(_emrcontainersReleaseLabel)
	}
	if len(_emrcontainersType) > 0 {
		input.Type = aws.String(_emrcontainersType)
	}
	if len(_emrcontainersVirtualClusterId) > 0 {
		input.VirtualClusterId = aws.String(_emrcontainersVirtualClusterId)
	}
	if len(_emrcontainersCertificateArn) > 0 {
		input.CertificateArn = aws.String(_emrcontainersCertificateArn)
	}
	if len(_emrcontainersConfigurationOverrides) > 0 {
		if err := assignInputField(input, "ConfigurationOverrides", _emrcontainersConfigurationOverrides); err != nil {
			log.Errorf("invalid --configuration-overrides: %s", err.Error())
			return
		}
	}
	if len(_emrcontainersTags) > 0 {
		if err := assignInputField(input, "Tags", _emrcontainersTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateManagedEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a security configuration. Security configurations in Amazon EMR on EKS
// are templates for different security setups. You can use security configurations
// to configure the Lake Formation integration setup. You can also create a
// security configuration to re-use a security setup each time you create a virtual
// cluster.
func emrcontainers_CreateSecurityConfiguration(cfg aws.Config, client *emrcontainers.Client) {
	input := &emrcontainers.CreateSecurityConfigurationInput{
		// ClientToken: *string, // Required
		// Name: *string, // Required
		// SecurityConfigurationData: *types.SecurityConfigurationData, // Required
	}

	if len(_emrcontainersClientToken) > 0 {
		input.ClientToken = aws.String(_emrcontainersClientToken)
	}
	if len(_emrcontainersName) > 0 {
		input.Name = aws.String(_emrcontainersName)
	}
	if len(_emrcontainersSecurityConfigurationData) > 0 {
		if err := assignInputField(input, "SecurityConfigurationData", _emrcontainersSecurityConfigurationData); err != nil {
			log.Errorf("invalid --security-configuration-data: %s", err.Error())
			return
		}
	}
	if len(_emrcontainersContainerProvider) > 0 {
		if err := assignInputField(input, "ContainerProvider", _emrcontainersContainerProvider); err != nil {
			log.Errorf("invalid --container-provider: %s", err.Error())
			return
		}
	}
	if len(_emrcontainersTags) > 0 {
		if err := assignInputField(input, "Tags", _emrcontainersTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSecurityConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a virtual cluster. Virtual cluster is a managed entity on Amazon EMR on
// EKS. You can create, describe, list and delete virtual clusters. They do not
// consume any additional resource in your system. A single virtual cluster maps to
// a single Kubernetes namespace. Given this relationship, you can model virtual
// clusters the same way you model Kubernetes namespaces to meet your requirements.
func emrcontainers_CreateVirtualCluster(cfg aws.Config, client *emrcontainers.Client) {
	input := &emrcontainers.CreateVirtualClusterInput{
		// ClientToken: *string, // Required
		// ContainerProvider: *types.ContainerProvider, // Required
		// Name: *string, // Required
	}

	if len(_emrcontainersClientToken) > 0 {
		input.ClientToken = aws.String(_emrcontainersClientToken)
	}
	if len(_emrcontainersContainerProvider) > 0 {
		if err := assignInputField(input, "ContainerProvider", _emrcontainersContainerProvider); err != nil {
			log.Errorf("invalid --container-provider: %s", err.Error())
			return
		}
	}
	if len(_emrcontainersName) > 0 {
		input.Name = aws.String(_emrcontainersName)
	}
	if len(_emrcontainersSecurityConfigurationId) > 0 {
		input.SecurityConfigurationId = aws.String(_emrcontainersSecurityConfigurationId)
	}
	if len(_emrcontainersTags) > 0 {
		if err := assignInputField(input, "Tags", _emrcontainersTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateVirtualCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a job template. Job template stores values of StartJobRun API request
// in a template and can be used to start a job run. Job template allows two use
// cases: avoid repeating recurring StartJobRun API request values, enforcing
// certain values in StartJobRun API request.
func emrcontainers_DeleteJobTemplate(cfg aws.Config, client *emrcontainers.Client) {
	input := &emrcontainers.DeleteJobTemplateInput{
		// Id: *string, // Required
	}

	if len(_emrcontainersId) > 0 {
		input.Id = aws.String(_emrcontainersId)
	}

	if resp, err := client.DeleteJobTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a managed endpoint. A managed endpoint is a gateway that connects
// Amazon EMR Studio to Amazon EMR on EKS so that Amazon EMR Studio can communicate
// with your virtual cluster.
func emrcontainers_DeleteManagedEndpoint(cfg aws.Config, client *emrcontainers.Client) {
	input := &emrcontainers.DeleteManagedEndpointInput{
		// Id: *string, // Required
		// VirtualClusterId: *string, // Required
	}

	if len(_emrcontainersId) > 0 {
		input.Id = aws.String(_emrcontainersId)
	}
	if len(_emrcontainersVirtualClusterId) > 0 {
		input.VirtualClusterId = aws.String(_emrcontainersVirtualClusterId)
	}

	if resp, err := client.DeleteManagedEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a virtual cluster. Virtual cluster is a managed entity on Amazon EMR on
// EKS. You can create, describe, list and delete virtual clusters. They do not
// consume any additional resource in your system. A single virtual cluster maps to
// a single Kubernetes namespace. Given this relationship, you can model virtual
// clusters the same way you model Kubernetes namespaces to meet your requirements.
func emrcontainers_DeleteVirtualCluster(cfg aws.Config, client *emrcontainers.Client) {
	input := &emrcontainers.DeleteVirtualClusterInput{
		// Id: *string, // Required
	}

	if len(_emrcontainersId) > 0 {
		input.Id = aws.String(_emrcontainersId)
	}

	if resp, err := client.DeleteVirtualCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays detailed information about a job run. A job run is a unit of work,
// such as a Spark jar, PySpark script, or SparkSQL query, that you submit to
// Amazon EMR on EKS.
func emrcontainers_DescribeJobRun(cfg aws.Config, client *emrcontainers.Client) {
	input := &emrcontainers.DescribeJobRunInput{
		// Id: *string, // Required
		// VirtualClusterId: *string, // Required
	}

	if len(_emrcontainersId) > 0 {
		input.Id = aws.String(_emrcontainersId)
	}
	if len(_emrcontainersVirtualClusterId) > 0 {
		input.VirtualClusterId = aws.String(_emrcontainersVirtualClusterId)
	}

	if resp, err := client.DescribeJobRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays detailed information about a specified job template. Job template
// stores values of StartJobRun API request in a template and can be used to start
// a job run. Job template allows two use cases: avoid repeating recurring
// StartJobRun API request values, enforcing certain values in StartJobRun API
// request.
func emrcontainers_DescribeJobTemplate(cfg aws.Config, client *emrcontainers.Client) {
	input := &emrcontainers.DescribeJobTemplateInput{
		// Id: *string, // Required
	}

	if len(_emrcontainersId) > 0 {
		input.Id = aws.String(_emrcontainersId)
	}

	if resp, err := client.DescribeJobTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays detailed information about a managed endpoint. A managed endpoint is a
// gateway that connects Amazon EMR Studio to Amazon EMR on EKS so that Amazon EMR
// Studio can communicate with your virtual cluster.
func emrcontainers_DescribeManagedEndpoint(cfg aws.Config, client *emrcontainers.Client) {
	input := &emrcontainers.DescribeManagedEndpointInput{
		// Id: *string, // Required
		// VirtualClusterId: *string, // Required
	}

	if len(_emrcontainersId) > 0 {
		input.Id = aws.String(_emrcontainersId)
	}
	if len(_emrcontainersVirtualClusterId) > 0 {
		input.VirtualClusterId = aws.String(_emrcontainersVirtualClusterId)
	}

	if resp, err := client.DescribeManagedEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays detailed information about a specified security configuration.
// Security configurations in Amazon EMR on EKS are templates for different
// security setups. You can use security configurations to configure the Lake
// Formation integration setup. You can also create a security configuration to
// re-use a security setup each time you create a virtual cluster.
func emrcontainers_DescribeSecurityConfiguration(cfg aws.Config, client *emrcontainers.Client) {
	input := &emrcontainers.DescribeSecurityConfigurationInput{
		// Id: *string, // Required
	}

	if len(_emrcontainersId) > 0 {
		input.Id = aws.String(_emrcontainersId)
	}

	if resp, err := client.DescribeSecurityConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays detailed information about a specified virtual cluster. Virtual
// cluster is a managed entity on Amazon EMR on EKS. You can create, describe, list
// and delete virtual clusters. They do not consume any additional resource in your
// system. A single virtual cluster maps to a single Kubernetes namespace. Given
// this relationship, you can model virtual clusters the same way you model
// Kubernetes namespaces to meet your requirements.
func emrcontainers_DescribeVirtualCluster(cfg aws.Config, client *emrcontainers.Client) {
	input := &emrcontainers.DescribeVirtualClusterInput{
		// Id: *string, // Required
	}

	if len(_emrcontainersId) > 0 {
		input.Id = aws.String(_emrcontainersId)
	}

	if resp, err := client.DescribeVirtualCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generate a session token to connect to a managed endpoint.
func emrcontainers_GetManagedEndpointSessionCredentials(cfg aws.Config, client *emrcontainers.Client) {
	input := &emrcontainers.GetManagedEndpointSessionCredentialsInput{
		// CredentialType: *string, // Required
		// EndpointIdentifier: *string, // Required
		// ExecutionRoleArn: *string, // Required
		// VirtualClusterIdentifier: *string, // Required
	}

	if len(_emrcontainersCredentialType) > 0 {
		input.CredentialType = aws.String(_emrcontainersCredentialType)
	}
	if len(_emrcontainersEndpointIdentifier) > 0 {
		input.EndpointIdentifier = aws.String(_emrcontainersEndpointIdentifier)
	}
	if len(_emrcontainersExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_emrcontainersExecutionRoleArn)
	}
	if len(_emrcontainersVirtualClusterIdentifier) > 0 {
		input.VirtualClusterIdentifier = aws.String(_emrcontainersVirtualClusterIdentifier)
	}
	if len(_emrcontainersClientToken) > 0 {
		input.ClientToken = aws.String(_emrcontainersClientToken)
	}
	if len(_emrcontainersDurationInSeconds) > 0 {
		if err := assignInputField(input, "DurationInSeconds", _emrcontainersDurationInSeconds); err != nil {
			log.Errorf("invalid --duration-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_emrcontainersLogContext) > 0 {
		input.LogContext = aws.String(_emrcontainersLogContext)
	}

	if resp, err := client.GetManagedEndpointSessionCredentials(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists job runs based on a set of parameters. A job run is a unit of work, such
// as a Spark jar, PySpark script, or SparkSQL query, that you submit to Amazon EMR
// on EKS.
func emrcontainers_ListJobRuns(cfg aws.Config, client *emrcontainers.Client) {
	input := &emrcontainers.ListJobRunsInput{
		// VirtualClusterId: *string, // Required
	}

	if len(_emrcontainersVirtualClusterId) > 0 {
		input.VirtualClusterId = aws.String(_emrcontainersVirtualClusterId)
	}
	if len(_emrcontainersCreatedAfter) > 0 {
		if err := assignInputField(input, "CreatedAfter", _emrcontainersCreatedAfter); err != nil {
			log.Errorf("invalid --created-after: %s", err.Error())
			return
		}
	}
	if len(_emrcontainersCreatedBefore) > 0 {
		if err := assignInputField(input, "CreatedBefore", _emrcontainersCreatedBefore); err != nil {
			log.Errorf("invalid --created-before: %s", err.Error())
			return
		}
	}
	if len(_emrcontainersMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _emrcontainersMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_emrcontainersName) > 0 {
		input.Name = aws.String(_emrcontainersName)
	}
	if len(_emrcontainersNextToken) > 0 {
		input.NextToken = aws.String(_emrcontainersNextToken)
	}
	if len(_emrcontainersStates) > 0 {
		if err := assignInputField(input, "States", _emrcontainersStates); err != nil {
			log.Errorf("invalid --states: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListJobRuns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*emrcontainers.ListJobRunsOutput
	p := emrcontainers.NewListJobRunsPaginator(client, input)
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

// Lists job templates based on a set of parameters. Job template stores values of
// StartJobRun API request in a template and can be used to start a job run. Job
// template allows two use cases: avoid repeating recurring StartJobRun API request
// values, enforcing certain values in StartJobRun API request.
func emrcontainers_ListJobTemplates(cfg aws.Config, client *emrcontainers.Client) {
	input := &emrcontainers.ListJobTemplatesInput{}

	if len(_emrcontainersCreatedAfter) > 0 {
		if err := assignInputField(input, "CreatedAfter", _emrcontainersCreatedAfter); err != nil {
			log.Errorf("invalid --created-after: %s", err.Error())
			return
		}
	}
	if len(_emrcontainersCreatedBefore) > 0 {
		if err := assignInputField(input, "CreatedBefore", _emrcontainersCreatedBefore); err != nil {
			log.Errorf("invalid --created-before: %s", err.Error())
			return
		}
	}
	if len(_emrcontainersMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _emrcontainersMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_emrcontainersNextToken) > 0 {
		input.NextToken = aws.String(_emrcontainersNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListJobTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*emrcontainers.ListJobTemplatesOutput
	p := emrcontainers.NewListJobTemplatesPaginator(client, input)
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

// Lists managed endpoints based on a set of parameters. A managed endpoint is a
// gateway that connects Amazon EMR Studio to Amazon EMR on EKS so that Amazon EMR
// Studio can communicate with your virtual cluster.
func emrcontainers_ListManagedEndpoints(cfg aws.Config, client *emrcontainers.Client) {
	input := &emrcontainers.ListManagedEndpointsInput{
		// VirtualClusterId: *string, // Required
	}

	if len(_emrcontainersVirtualClusterId) > 0 {
		input.VirtualClusterId = aws.String(_emrcontainersVirtualClusterId)
	}
	if len(_emrcontainersCreatedAfter) > 0 {
		if err := assignInputField(input, "CreatedAfter", _emrcontainersCreatedAfter); err != nil {
			log.Errorf("invalid --created-after: %s", err.Error())
			return
		}
	}
	if len(_emrcontainersCreatedBefore) > 0 {
		if err := assignInputField(input, "CreatedBefore", _emrcontainersCreatedBefore); err != nil {
			log.Errorf("invalid --created-before: %s", err.Error())
			return
		}
	}
	if len(_emrcontainersMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _emrcontainersMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_emrcontainersNextToken) > 0 {
		input.NextToken = aws.String(_emrcontainersNextToken)
	}
	if len(_emrcontainersStates) > 0 {
		if err := assignInputField(input, "States", _emrcontainersStates); err != nil {
			log.Errorf("invalid --states: %s", err.Error())
			return
		}
	}
	if len(_emrcontainersTypes) > 0 {
		input.Types = append([]string(nil), _emrcontainersTypes...)
	}

	if disablePaginator() {
		if resp, err := client.ListManagedEndpoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*emrcontainers.ListManagedEndpointsOutput
	p := emrcontainers.NewListManagedEndpointsPaginator(client, input)
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

// Lists security configurations based on a set of parameters. Security
// configurations in Amazon EMR on EKS are templates for different security setups.
// You can use security configurations to configure the Lake Formation integration
// setup. You can also create a security configuration to re-use a security setup
// each time you create a virtual cluster.
func emrcontainers_ListSecurityConfigurations(cfg aws.Config, client *emrcontainers.Client) {
	input := &emrcontainers.ListSecurityConfigurationsInput{}

	if len(_emrcontainersCreatedAfter) > 0 {
		if err := assignInputField(input, "CreatedAfter", _emrcontainersCreatedAfter); err != nil {
			log.Errorf("invalid --created-after: %s", err.Error())
			return
		}
	}
	if len(_emrcontainersCreatedBefore) > 0 {
		if err := assignInputField(input, "CreatedBefore", _emrcontainersCreatedBefore); err != nil {
			log.Errorf("invalid --created-before: %s", err.Error())
			return
		}
	}
	if len(_emrcontainersMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _emrcontainersMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_emrcontainersNextToken) > 0 {
		input.NextToken = aws.String(_emrcontainersNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSecurityConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*emrcontainers.ListSecurityConfigurationsOutput
	p := emrcontainers.NewListSecurityConfigurationsPaginator(client, input)
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

// Lists the tags assigned to the resources.
func emrcontainers_ListTagsForResource(cfg aws.Config, client *emrcontainers.Client) {
	input := &emrcontainers.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_emrcontainersResourceArn) > 0 {
		input.ResourceArn = aws.String(_emrcontainersResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists information about the specified virtual cluster. Virtual cluster is a
// managed entity on Amazon EMR on EKS. You can create, describe, list and delete
// virtual clusters. They do not consume any additional resource in your system. A
// single virtual cluster maps to a single Kubernetes namespace. Given this
// relationship, you can model virtual clusters the same way you model Kubernetes
// namespaces to meet your requirements.
func emrcontainers_ListVirtualClusters(cfg aws.Config, client *emrcontainers.Client) {
	input := &emrcontainers.ListVirtualClustersInput{}

	if len(_emrcontainersContainerProviderId) > 0 {
		input.ContainerProviderId = aws.String(_emrcontainersContainerProviderId)
	}
	if len(_emrcontainersContainerProviderType) > 0 {
		if err := assignInputField(input, "ContainerProviderType", _emrcontainersContainerProviderType); err != nil {
			log.Errorf("invalid --container-provider-type: %s", err.Error())
			return
		}
	}
	if len(_emrcontainersCreatedAfter) > 0 {
		if err := assignInputField(input, "CreatedAfter", _emrcontainersCreatedAfter); err != nil {
			log.Errorf("invalid --created-after: %s", err.Error())
			return
		}
	}
	if len(_emrcontainersCreatedBefore) > 0 {
		if err := assignInputField(input, "CreatedBefore", _emrcontainersCreatedBefore); err != nil {
			log.Errorf("invalid --created-before: %s", err.Error())
			return
		}
	}
	if len(_emrcontainersEksAccessEntryIntegrated) > 0 {
		if err := assignInputField(input, "EksAccessEntryIntegrated", _emrcontainersEksAccessEntryIntegrated); err != nil {
			log.Errorf("invalid --eks-access-entry-integrated: %s", err.Error())
			return
		}
	}
	if len(_emrcontainersMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _emrcontainersMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_emrcontainersNextToken) > 0 {
		input.NextToken = aws.String(_emrcontainersNextToken)
	}
	if len(_emrcontainersStates) > 0 {
		if err := assignInputField(input, "States", _emrcontainersStates); err != nil {
			log.Errorf("invalid --states: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListVirtualClusters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*emrcontainers.ListVirtualClustersOutput
	p := emrcontainers.NewListVirtualClustersPaginator(client, input)
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

// Starts a job run. A job run is a unit of work, such as a Spark jar, PySpark
// script, or SparkSQL query, that you submit to Amazon EMR on EKS.
func emrcontainers_StartJobRun(cfg aws.Config, client *emrcontainers.Client) {
	input := &emrcontainers.StartJobRunInput{
		// ClientToken: *string, // Required
		// VirtualClusterId: *string, // Required
	}

	if len(_emrcontainersClientToken) > 0 {
		input.ClientToken = aws.String(_emrcontainersClientToken)
	}
	if len(_emrcontainersVirtualClusterId) > 0 {
		input.VirtualClusterId = aws.String(_emrcontainersVirtualClusterId)
	}
	if len(_emrcontainersConfigurationOverrides) > 0 {
		if err := assignInputField(input, "ConfigurationOverrides", _emrcontainersConfigurationOverrides); err != nil {
			log.Errorf("invalid --configuration-overrides: %s", err.Error())
			return
		}
	}
	if len(_emrcontainersExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_emrcontainersExecutionRoleArn)
	}
	if len(_emrcontainersJobDriver) > 0 {
		if err := assignInputField(input, "JobDriver", _emrcontainersJobDriver); err != nil {
			log.Errorf("invalid --job-driver: %s", err.Error())
			return
		}
	}
	if len(_emrcontainersJobTemplateId) > 0 {
		input.JobTemplateId = aws.String(_emrcontainersJobTemplateId)
	}
	if len(_emrcontainersJobTemplateParameters) > 0 {
		if err := assignInputField(input, "JobTemplateParameters", _emrcontainersJobTemplateParameters); err != nil {
			log.Errorf("invalid --job-template-parameters: %s", err.Error())
			return
		}
	}
	if len(_emrcontainersName) > 0 {
		input.Name = aws.String(_emrcontainersName)
	}
	if len(_emrcontainersReleaseLabel) > 0 {
		input.ReleaseLabel = aws.String(_emrcontainersReleaseLabel)
	}
	if len(_emrcontainersRetryPolicyConfiguration) > 0 {
		if err := assignInputField(input, "RetryPolicyConfiguration", _emrcontainersRetryPolicyConfiguration); err != nil {
			log.Errorf("invalid --retry-policy-configuration: %s", err.Error())
			return
		}
	}
	if len(_emrcontainersTags) > 0 {
		if err := assignInputField(input, "Tags", _emrcontainersTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartJobRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns tags to resources. A tag is a label that you assign to an Amazon Web
// Services resource. Each tag consists of a key and an optional value, both of
// which you define. Tags enable you to categorize your Amazon Web Services
// resources by attributes such as purpose, owner, or environment. When you have
// many resources of the same type, you can quickly identify a specific resource
// based on the tags you've assigned to it. For example, you can define a set of
// tags for your Amazon EMR on EKS clusters to help you track each cluster's owner
// and stack level. We recommend that you devise a consistent set of tag keys for
// each resource type. You can then search and filter the resources based on the
// tags that you add.
func emrcontainers_TagResource(cfg aws.Config, client *emrcontainers.Client) {
	input := &emrcontainers.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_emrcontainersResourceArn) > 0 {
		input.ResourceArn = aws.String(_emrcontainersResourceArn)
	}
	if len(_emrcontainersTags) > 0 {
		if err := assignInputField(input, "Tags", _emrcontainersTags); err != nil {
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

// Removes tags from resources.
func emrcontainers_UntagResource(cfg aws.Config, client *emrcontainers.Client) {
	input := &emrcontainers.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_emrcontainersResourceArn) > 0 {
		input.ResourceArn = aws.String(_emrcontainersResourceArn)
	}
	if len(_emrcontainersTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _emrcontainersTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_emrcontainersCmd)
	_emrcontainersCmd.Flags().SortFlags = false

	_emrcontainersCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_emrcontainersCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_emrcontainersCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersCertificateArn, "certificate-arn", "", "", "Certificate ARN")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersClientToken, "client-token", "", "", "Client Token")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersConfigurationOverrides, "configuration-overrides", "", "", "Configuration Overrides")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersContainerProvider, "container-provider", "", "", "Container Provider")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersContainerProviderId, "container-provider-id", "", "", "Container Provider ID")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersContainerProviderType, "container-provider-type", "", "", "Container Provider Type")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersCreatedAfter, "created-after", "", "", "Created After")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersCreatedBefore, "created-before", "", "", "Created Before")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersCredentialType, "credential-type", "", "", "Credential Type")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersDurationInSeconds, "duration-in-seconds", "", "", "Duration In Seconds")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersEksAccessEntryIntegrated, "eks-access-entry-integrated", "", "", "Eks Access Entry Integrated")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersEndpointIdentifier, "endpoint-identifier", "", "", "Endpoint Identifier")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersExecutionRoleArn, "execution-role-arn", "", "", "Execution Role ARN")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersId, "id", "", "", "ID")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersJobDriver, "job-driver", "", "", "Job Driver")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersJobTemplateData, "job-template-data", "", "", "Job Template Data")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersJobTemplateId, "job-template-id", "", "", "Job Template ID")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersJobTemplateParameters, "job-template-parameters", "", "", "Job Template Parameters")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersKmsKeyArn, "kms-key-arn", "", "", "KMS Key ARN")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersLogContext, "log-context", "", "", "Log Context")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersMaxResults, "max-results", "", "", "Max Results")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersName, "name", "", "", "Name")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersNextToken, "next-token", "", "", "Next Token")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersReleaseLabel, "release-label", "", "", "Release Label")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersResourceArn, "resource-arn", "", "", "Resource ARN")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersRetryPolicyConfiguration, "retry-policy-configuration", "", "", "Retry Policy Configuration")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersSecurityConfigurationData, "security-configuration-data", "", "", "Security Configuration Data")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersSecurityConfigurationId, "security-configuration-id", "", "", "Security Configuration ID")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersStates, "states", "", "", "States")
	_emrcontainersCmd.Flags().StringSliceVarP(&_emrcontainersTagKeys, "tag-keys", "", nil, "Tag Keys")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersTags, "tags", "", "", "Tags")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersType, "type", "", "", "Type")
	_emrcontainersCmd.Flags().StringSliceVarP(&_emrcontainersTypes, "types", "", nil, "Types")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersVirtualClusterId, "virtual-cluster-id", "", "", "Virtual Cluster ID")
	_emrcontainersCmd.Flags().StringVarP(&_emrcontainersVirtualClusterIdentifier, "virtual-cluster-identifier", "", "", "Virtual Cluster Identifier")

	_emrcontainersCmd.Flags().BoolVarP(&_emrcontainersCancelJobRun, "cancel-job-run", "", false, "Cancel Job Run")
	_emrcontainersCmd.Flags().BoolVarP(&_emrcontainersCreateJobTemplate, "create-job-template", "", false, "Create Job Template")
	_emrcontainersCmd.Flags().BoolVarP(&_emrcontainersCreateManagedEndpoint, "create-managed-endpoint", "", false, "Create Managed Endpoint")
	_emrcontainersCmd.Flags().BoolVarP(&_emrcontainersCreateSecurityConfiguration, "create-security-configuration", "", false, "Create Security Configuration")
	_emrcontainersCmd.Flags().BoolVarP(&_emrcontainersCreateVirtualCluster, "create-virtual-cluster", "", false, "Create Virtual Cluster")
	_emrcontainersCmd.Flags().BoolVarP(&_emrcontainersDeleteJobTemplate, "delete-job-template", "", false, "Delete Job Template")
	_emrcontainersCmd.Flags().BoolVarP(&_emrcontainersDeleteManagedEndpoint, "delete-managed-endpoint", "", false, "Delete Managed Endpoint")
	_emrcontainersCmd.Flags().BoolVarP(&_emrcontainersDeleteVirtualCluster, "delete-virtual-cluster", "", false, "Delete Virtual Cluster")
	_emrcontainersCmd.Flags().BoolVarP(&_emrcontainersDescribeJobRun, "describe-job-run", "", false, "Describe Job Run")
	_emrcontainersCmd.Flags().BoolVarP(&_emrcontainersDescribeJobTemplate, "describe-job-template", "", false, "Describe Job Template")
	_emrcontainersCmd.Flags().BoolVarP(&_emrcontainersDescribeManagedEndpoint, "describe-managed-endpoint", "", false, "Describe Managed Endpoint")
	_emrcontainersCmd.Flags().BoolVarP(&_emrcontainersDescribeSecurityConfiguration, "describe-security-configuration", "", false, "Describe Security Configuration")
	_emrcontainersCmd.Flags().BoolVarP(&_emrcontainersDescribeVirtualCluster, "describe-virtual-cluster", "", false, "Describe Virtual Cluster")
	_emrcontainersCmd.Flags().BoolVarP(&_emrcontainersGetManagedEndpointSessionCredentials, "get-managed-endpoint-session-credentials", "", false, "Get Managed Endpoint Session Credentials")
	_emrcontainersCmd.Flags().BoolVarP(&_emrcontainersListJobRuns, "list-job-runs", "", false, "List Job Runs")
	_emrcontainersCmd.Flags().BoolVarP(&_emrcontainersListJobTemplates, "list-job-templates", "", false, "List Job Templates")
	_emrcontainersCmd.Flags().BoolVarP(&_emrcontainersListManagedEndpoints, "list-managed-endpoints", "", false, "List Managed Endpoints")
	_emrcontainersCmd.Flags().BoolVarP(&_emrcontainersListSecurityConfigurations, "list-security-configurations", "", false, "List Security Configurations")
	_emrcontainersCmd.Flags().BoolVarP(&_emrcontainersListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_emrcontainersCmd.Flags().BoolVarP(&_emrcontainersListVirtualClusters, "list-virtual-clusters", "", false, "List Virtual Clusters")
	_emrcontainersCmd.Flags().BoolVarP(&_emrcontainersStartJobRun, "start-job-run", "", false, "Start Job Run")
	_emrcontainersCmd.Flags().BoolVarP(&_emrcontainersTagResource, "tag-resource", "", false, "Tag Resource")
	_emrcontainersCmd.Flags().BoolVarP(&_emrcontainersUntagResource, "untag-resource", "", false, "Untag Resource")

}
