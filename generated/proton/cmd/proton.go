package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/proton"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// protonCmd represents the proton command
var _protonCmd = &cobra.Command{
	Use:   "proton",
	Short: "AWS proton CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := proton.NewFromConfig(cfg)
		if _protonAcceptEnvironmentAccountConnection {
			proton_AcceptEnvironmentAccountConnection(cfg, client)
			return
		}
		if _protonCancelComponentDeployment {
			proton_CancelComponentDeployment(cfg, client)
			return
		}
		if _protonCancelEnvironmentDeployment {
			proton_CancelEnvironmentDeployment(cfg, client)
			return
		}
		if _protonCancelServiceInstanceDeployment {
			proton_CancelServiceInstanceDeployment(cfg, client)
			return
		}
		if _protonCancelServicePipelineDeployment {
			proton_CancelServicePipelineDeployment(cfg, client)
			return
		}
		if _protonCreateComponent {
			proton_CreateComponent(cfg, client)
			return
		}
		if _protonCreateEnvironment {
			proton_CreateEnvironment(cfg, client)
			return
		}
		if _protonCreateEnvironmentAccountConnection {
			proton_CreateEnvironmentAccountConnection(cfg, client)
			return
		}
		if _protonCreateEnvironmentTemplate {
			proton_CreateEnvironmentTemplate(cfg, client)
			return
		}
		if _protonCreateEnvironmentTemplateVersion {
			proton_CreateEnvironmentTemplateVersion(cfg, client)
			return
		}
		if _protonCreateRepository {
			proton_CreateRepository(cfg, client)
			return
		}
		if _protonCreateService {
			proton_CreateService(cfg, client)
			return
		}
		if _protonCreateServiceInstance {
			proton_CreateServiceInstance(cfg, client)
			return
		}
		if _protonCreateServiceSyncConfig {
			proton_CreateServiceSyncConfig(cfg, client)
			return
		}
		if _protonCreateServiceTemplate {
			proton_CreateServiceTemplate(cfg, client)
			return
		}
		if _protonCreateServiceTemplateVersion {
			proton_CreateServiceTemplateVersion(cfg, client)
			return
		}
		if _protonCreateTemplateSyncConfig {
			proton_CreateTemplateSyncConfig(cfg, client)
			return
		}
		if _protonDeleteComponent {
			proton_DeleteComponent(cfg, client)
			return
		}
		if _protonDeleteDeployment {
			proton_DeleteDeployment(cfg, client)
			return
		}
		if _protonDeleteEnvironment {
			proton_DeleteEnvironment(cfg, client)
			return
		}
		if _protonDeleteEnvironmentAccountConnection {
			proton_DeleteEnvironmentAccountConnection(cfg, client)
			return
		}
		if _protonDeleteEnvironmentTemplate {
			proton_DeleteEnvironmentTemplate(cfg, client)
			return
		}
		if _protonDeleteEnvironmentTemplateVersion {
			proton_DeleteEnvironmentTemplateVersion(cfg, client)
			return
		}
		if _protonDeleteRepository {
			proton_DeleteRepository(cfg, client)
			return
		}
		if _protonDeleteService {
			proton_DeleteService(cfg, client)
			return
		}
		if _protonDeleteServiceSyncConfig {
			proton_DeleteServiceSyncConfig(cfg, client)
			return
		}
		if _protonDeleteServiceTemplate {
			proton_DeleteServiceTemplate(cfg, client)
			return
		}
		if _protonDeleteServiceTemplateVersion {
			proton_DeleteServiceTemplateVersion(cfg, client)
			return
		}
		if _protonDeleteTemplateSyncConfig {
			proton_DeleteTemplateSyncConfig(cfg, client)
			return
		}
		if _protonGetAccountSettings {
			proton_GetAccountSettings(cfg, client)
			return
		}
		if _protonGetComponent {
			proton_GetComponent(cfg, client)
			return
		}
		if _protonGetDeployment {
			proton_GetDeployment(cfg, client)
			return
		}
		if _protonGetEnvironment {
			proton_GetEnvironment(cfg, client)
			return
		}
		if _protonGetEnvironmentAccountConnection {
			proton_GetEnvironmentAccountConnection(cfg, client)
			return
		}
		if _protonGetEnvironmentTemplate {
			proton_GetEnvironmentTemplate(cfg, client)
			return
		}
		if _protonGetEnvironmentTemplateVersion {
			proton_GetEnvironmentTemplateVersion(cfg, client)
			return
		}
		if _protonGetRepository {
			proton_GetRepository(cfg, client)
			return
		}
		if _protonGetRepositorySyncStatus {
			proton_GetRepositorySyncStatus(cfg, client)
			return
		}
		if _protonGetResourcesSummary {
			proton_GetResourcesSummary(cfg, client)
			return
		}
		if _protonGetService {
			proton_GetService(cfg, client)
			return
		}
		if _protonGetServiceInstance {
			proton_GetServiceInstance(cfg, client)
			return
		}
		if _protonGetServiceInstanceSyncStatus {
			proton_GetServiceInstanceSyncStatus(cfg, client)
			return
		}
		if _protonGetServiceSyncBlockerSummary {
			proton_GetServiceSyncBlockerSummary(cfg, client)
			return
		}
		if _protonGetServiceSyncConfig {
			proton_GetServiceSyncConfig(cfg, client)
			return
		}
		if _protonGetServiceTemplate {
			proton_GetServiceTemplate(cfg, client)
			return
		}
		if _protonGetServiceTemplateVersion {
			proton_GetServiceTemplateVersion(cfg, client)
			return
		}
		if _protonGetTemplateSyncConfig {
			proton_GetTemplateSyncConfig(cfg, client)
			return
		}
		if _protonGetTemplateSyncStatus {
			proton_GetTemplateSyncStatus(cfg, client)
			return
		}
		if _protonListComponentOutputs {
			proton_ListComponentOutputs(cfg, client)
			return
		}
		if _protonListComponentProvisionedResources {
			proton_ListComponentProvisionedResources(cfg, client)
			return
		}
		if _protonListComponents {
			proton_ListComponents(cfg, client)
			return
		}
		if _protonListDeployments {
			proton_ListDeployments(cfg, client)
			return
		}
		if _protonListEnvironmentAccountConnections {
			proton_ListEnvironmentAccountConnections(cfg, client)
			return
		}
		if _protonListEnvironmentOutputs {
			proton_ListEnvironmentOutputs(cfg, client)
			return
		}
		if _protonListEnvironmentProvisionedResources {
			proton_ListEnvironmentProvisionedResources(cfg, client)
			return
		}
		if _protonListEnvironmentTemplateVersions {
			proton_ListEnvironmentTemplateVersions(cfg, client)
			return
		}
		if _protonListEnvironmentTemplates {
			proton_ListEnvironmentTemplates(cfg, client)
			return
		}
		if _protonListEnvironments {
			proton_ListEnvironments(cfg, client)
			return
		}
		if _protonListRepositories {
			proton_ListRepositories(cfg, client)
			return
		}
		if _protonListRepositorySyncDefinitions {
			proton_ListRepositorySyncDefinitions(cfg, client)
			return
		}
		if _protonListServiceInstanceOutputs {
			proton_ListServiceInstanceOutputs(cfg, client)
			return
		}
		if _protonListServiceInstanceProvisionedResources {
			proton_ListServiceInstanceProvisionedResources(cfg, client)
			return
		}
		if _protonListServiceInstances {
			proton_ListServiceInstances(cfg, client)
			return
		}
		if _protonListServicePipelineOutputs {
			proton_ListServicePipelineOutputs(cfg, client)
			return
		}
		if _protonListServicePipelineProvisionedResources {
			proton_ListServicePipelineProvisionedResources(cfg, client)
			return
		}
		if _protonListServiceTemplateVersions {
			proton_ListServiceTemplateVersions(cfg, client)
			return
		}
		if _protonListServiceTemplates {
			proton_ListServiceTemplates(cfg, client)
			return
		}
		if _protonListServices {
			proton_ListServices(cfg, client)
			return
		}
		if _protonListTagsForResource {
			proton_ListTagsForResource(cfg, client)
			return
		}
		if _protonNotifyResourceDeploymentStatusChange {
			proton_NotifyResourceDeploymentStatusChange(cfg, client)
			return
		}
		if _protonRejectEnvironmentAccountConnection {
			proton_RejectEnvironmentAccountConnection(cfg, client)
			return
		}
		if _protonTagResource {
			proton_TagResource(cfg, client)
			return
		}
		if _protonUntagResource {
			proton_UntagResource(cfg, client)
			return
		}
		if _protonUpdateAccountSettings {
			proton_UpdateAccountSettings(cfg, client)
			return
		}
		if _protonUpdateComponent {
			proton_UpdateComponent(cfg, client)
			return
		}
		if _protonUpdateEnvironment {
			proton_UpdateEnvironment(cfg, client)
			return
		}
		if _protonUpdateEnvironmentAccountConnection {
			proton_UpdateEnvironmentAccountConnection(cfg, client)
			return
		}
		if _protonUpdateEnvironmentTemplate {
			proton_UpdateEnvironmentTemplate(cfg, client)
			return
		}
		if _protonUpdateEnvironmentTemplateVersion {
			proton_UpdateEnvironmentTemplateVersion(cfg, client)
			return
		}
		if _protonUpdateService {
			proton_UpdateService(cfg, client)
			return
		}
		if _protonUpdateServiceInstance {
			proton_UpdateServiceInstance(cfg, client)
			return
		}
		if _protonUpdateServicePipeline {
			proton_UpdateServicePipeline(cfg, client)
			return
		}
		if _protonUpdateServiceSyncBlocker {
			proton_UpdateServiceSyncBlocker(cfg, client)
			return
		}
		if _protonUpdateServiceSyncConfig {
			proton_UpdateServiceSyncConfig(cfg, client)
			return
		}
		if _protonUpdateServiceTemplate {
			proton_UpdateServiceTemplate(cfg, client)
			return
		}
		if _protonUpdateServiceTemplateVersion {
			proton_UpdateServiceTemplateVersion(cfg, client)
			return
		}
		if _protonUpdateTemplateSyncConfig {
			proton_UpdateTemplateSyncConfig(cfg, client)
			return
		}

	},
}

var (
	_protonAcceptEnvironmentAccountConnection      bool
	_protonCancelComponentDeployment               bool
	_protonCancelEnvironmentDeployment             bool
	_protonCancelServiceInstanceDeployment         bool
	_protonCancelServicePipelineDeployment         bool
	_protonCreateComponent                         bool
	_protonCreateEnvironment                       bool
	_protonCreateEnvironmentAccountConnection      bool
	_protonCreateEnvironmentTemplate               bool
	_protonCreateEnvironmentTemplateVersion        bool
	_protonCreateRepository                        bool
	_protonCreateService                           bool
	_protonCreateServiceInstance                   bool
	_protonCreateServiceSyncConfig                 bool
	_protonCreateServiceTemplate                   bool
	_protonCreateServiceTemplateVersion            bool
	_protonCreateTemplateSyncConfig                bool
	_protonDeleteComponent                         bool
	_protonDeleteDeployment                        bool
	_protonDeleteEnvironment                       bool
	_protonDeleteEnvironmentAccountConnection      bool
	_protonDeleteEnvironmentTemplate               bool
	_protonDeleteEnvironmentTemplateVersion        bool
	_protonDeleteRepository                        bool
	_protonDeleteService                           bool
	_protonDeleteServiceSyncConfig                 bool
	_protonDeleteServiceTemplate                   bool
	_protonDeleteServiceTemplateVersion            bool
	_protonDeleteTemplateSyncConfig                bool
	_protonGetAccountSettings                      bool
	_protonGetComponent                            bool
	_protonGetDeployment                           bool
	_protonGetEnvironment                          bool
	_protonGetEnvironmentAccountConnection         bool
	_protonGetEnvironmentTemplate                  bool
	_protonGetEnvironmentTemplateVersion           bool
	_protonGetRepository                           bool
	_protonGetRepositorySyncStatus                 bool
	_protonGetResourcesSummary                     bool
	_protonGetService                              bool
	_protonGetServiceInstance                      bool
	_protonGetServiceInstanceSyncStatus            bool
	_protonGetServiceSyncBlockerSummary            bool
	_protonGetServiceSyncConfig                    bool
	_protonGetServiceTemplate                      bool
	_protonGetServiceTemplateVersion               bool
	_protonGetTemplateSyncConfig                   bool
	_protonGetTemplateSyncStatus                   bool
	_protonListComponentOutputs                    bool
	_protonListComponentProvisionedResources       bool
	_protonListComponents                          bool
	_protonListDeployments                         bool
	_protonListEnvironmentAccountConnections       bool
	_protonListEnvironmentOutputs                  bool
	_protonListEnvironmentProvisionedResources     bool
	_protonListEnvironmentTemplateVersions         bool
	_protonListEnvironmentTemplates                bool
	_protonListEnvironments                        bool
	_protonListRepositories                        bool
	_protonListRepositorySyncDefinitions           bool
	_protonListServiceInstanceOutputs              bool
	_protonListServiceInstanceProvisionedResources bool
	_protonListServiceInstances                    bool
	_protonListServicePipelineOutputs              bool
	_protonListServicePipelineProvisionedResources bool
	_protonListServiceTemplateVersions             bool
	_protonListServiceTemplates                    bool
	_protonListServices                            bool
	_protonListTagsForResource                     bool
	_protonNotifyResourceDeploymentStatusChange    bool
	_protonRejectEnvironmentAccountConnection      bool
	_protonTagResource                             bool
	_protonUntagResource                           bool
	_protonUpdateAccountSettings                   bool
	_protonUpdateComponent                         bool
	_protonUpdateEnvironment                       bool
	_protonUpdateEnvironmentAccountConnection      bool
	_protonUpdateEnvironmentTemplate               bool
	_protonUpdateEnvironmentTemplateVersion        bool
	_protonUpdateService                           bool
	_protonUpdateServiceInstance                   bool
	_protonUpdateServicePipeline                   bool
	_protonUpdateServiceSyncBlocker                bool
	_protonUpdateServiceSyncConfig                 bool
	_protonUpdateServiceTemplate                   bool
	_protonUpdateServiceTemplateVersion            bool
	_protonUpdateTemplateSyncConfig                bool

	_protonBranch                               string
	_protonBranchName                           string
	_protonClientToken                          string
	_protonCodebuildRoleArn                     string
	_protonCompatibleEnvironmentTemplates       string
	_protonComponentName                        string
	_protonComponentRoleArn                     string
	_protonConnectionArn                        string
	_protonDeletePipelineProvisioningRepository string
	_protonDeploymentId                         string
	_protonDeploymentType                       string
	_protonDescription                          string
	_protonDisplayName                          string
	_protonEncryptionKey                        string
	_protonEnvironmentAccountConnectionId       string
	_protonEnvironmentName                      string
	_protonEnvironmentTemplates                 string
	_protonFilePath                             string
	_protonFilters                              string
	_protonId                                   string
	_protonMajorVersion                         string
	_protonManagementAccountId                  string
	_protonManifest                             string
	_protonMaxResults                           string
	_protonMinorVersion                         string
	_protonName                                 string
	_protonNextToken                            string
	_protonOutputs                              string
	_protonPipelineCodebuildRoleArn             string
	_protonPipelineProvisioning                 string
	_protonPipelineProvisioningRepository       string
	_protonPipelineServiceRoleArn               string
	_protonProtonServiceRoleArn                 string
	_protonProvider                             string
	_protonProvisioning                         string
	_protonProvisioningRepository               string
	_protonRepositoryConnectionArn              string
	_protonRepositoryId                         string
	_protonRepositoryName                       string
	_protonRepositoryProvider                   string
	_protonRequestedBy                          string
	_protonResolvedReason                       string
	_protonResourceArn                          string
	_protonRoleArn                              string
	_protonServiceInstanceName                  string
	_protonServiceName                          string
	_protonServiceSpec                          string
	_protonSortBy                               string
	_protonSortOrder                            string
	_protonSource                               string
	_protonSpec                                 string
	_protonStatus                               string
	_protonStatusMessage                        string
	_protonStatuses                             string
	_protonSubdirectory                         string
	_protonSupportedComponentSources            string
	_protonSyncType                             string
	_protonTagKeys                              []string
	_protonTags                                 string
	_protonTemplateFile                         string
	_protonTemplateMajorVersion                 string
	_protonTemplateMinorVersion                 string
	_protonTemplateName                         string
	_protonTemplateType                         string
	_protonTemplateVersion                      string
)

// In a management account, an environment account connection request is accepted.
// When the environment account connection request is accepted, Proton can use the
// associated IAM role to provision environment infrastructure resources in the
// associated environment account.
//
// For more information, see [Environment account connections] in the Proton User guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Environment account connections]: https://docs.aws.amazon.com/proton/latest/userguide/ag-env-account-connections.html
func proton_AcceptEnvironmentAccountConnection(cfg aws.Config, client *proton.Client) {
	input := &proton.AcceptEnvironmentAccountConnectionInput{
		// Id: *string, // Required
	}

	if len(_protonId) > 0 {
		input.Id = aws.String(_protonId)
	}

	if resp, err := client.AcceptEnvironmentAccountConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attempts to cancel a component deployment (for a component that is in the
// IN_PROGRESS deployment status).
//
// For more information about components, see [Proton components] in the Proton User Guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Proton components]: https://docs.aws.amazon.com/proton/latest/userguide/ag-components.html
func proton_CancelComponentDeployment(cfg aws.Config, client *proton.Client) {
	input := &proton.CancelComponentDeploymentInput{
		// ComponentName: *string, // Required
	}

	if len(_protonComponentName) > 0 {
		input.ComponentName = aws.String(_protonComponentName)
	}

	if resp, err := client.CancelComponentDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attempts to cancel an environment deployment on an UpdateEnvironment action, if the deployment
// is IN_PROGRESS . For more information, see [Update an environment] in the Proton User guide.
//
// The following list includes potential cancellation scenarios.
//
// - If the cancellation attempt succeeds, the resulting deployment state is
// CANCELLED .
//
// - If the cancellation attempt fails, the resulting deployment state is FAILED .
//
// - If the current UpdateEnvironmentaction succeeds before the cancellation attempt starts, the
// resulting deployment state is SUCCEEDED and the cancellation attempt has no
// effect.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Update an environment]: https://docs.aws.amazon.com/proton/latest/userguide/ag-env-update.html
func proton_CancelEnvironmentDeployment(cfg aws.Config, client *proton.Client) {
	input := &proton.CancelEnvironmentDeploymentInput{
		// EnvironmentName: *string, // Required
	}

	if len(_protonEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_protonEnvironmentName)
	}

	if resp, err := client.CancelEnvironmentDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attempts to cancel a service instance deployment on an UpdateServiceInstance action, if the
// deployment is IN_PROGRESS . For more information, see [Update a service instance] in the Proton User guide.
//
// The following list includes potential cancellation scenarios.
//
// - If the cancellation attempt succeeds, the resulting deployment state is
// CANCELLED .
//
// - If the cancellation attempt fails, the resulting deployment state is FAILED .
//
// - If the current UpdateServiceInstanceaction succeeds before the cancellation attempt starts, the
// resulting deployment state is SUCCEEDED and the cancellation attempt has no
// effect.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Update a service instance]: https://docs.aws.amazon.com/proton/latest/userguide/ag-svc-instance-update.html
func proton_CancelServiceInstanceDeployment(cfg aws.Config, client *proton.Client) {
	input := &proton.CancelServiceInstanceDeploymentInput{
		// ServiceInstanceName: *string, // Required
		// ServiceName: *string, // Required
	}

	if len(_protonServiceInstanceName) > 0 {
		input.ServiceInstanceName = aws.String(_protonServiceInstanceName)
	}
	if len(_protonServiceName) > 0 {
		input.ServiceName = aws.String(_protonServiceName)
	}

	if resp, err := client.CancelServiceInstanceDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attempts to cancel a service pipeline deployment on an UpdateServicePipeline action, if the
// deployment is IN_PROGRESS . For more information, see [Update a service pipeline] in the Proton User guide.
//
// The following list includes potential cancellation scenarios.
//
// - If the cancellation attempt succeeds, the resulting deployment state is
// CANCELLED .
//
// - If the cancellation attempt fails, the resulting deployment state is FAILED .
//
// - If the current UpdateServicePipelineaction succeeds before the cancellation attempt starts, the
// resulting deployment state is SUCCEEDED and the cancellation attempt has no
// effect.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Update a service pipeline]: https://docs.aws.amazon.com/proton/latest/userguide/ag-svc-pipeline-update.html
func proton_CancelServicePipelineDeployment(cfg aws.Config, client *proton.Client) {
	input := &proton.CancelServicePipelineDeploymentInput{
		// ServiceName: *string, // Required
	}

	if len(_protonServiceName) > 0 {
		input.ServiceName = aws.String(_protonServiceName)
	}

	if resp, err := client.CancelServicePipelineDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create an Proton component. A component is an infrastructure extension for a
// service instance.
//
// For more information about components, see [Proton components] in the Proton User Guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Proton components]: https://docs.aws.amazon.com/proton/latest/userguide/ag-components.html
func proton_CreateComponent(cfg aws.Config, client *proton.Client) {
	input := &proton.CreateComponentInput{
		// Manifest: *string, // Required
		// Name: *string, // Required
		// TemplateFile: *string, // Required
	}

	if len(_protonManifest) > 0 {
		input.Manifest = aws.String(_protonManifest)
	}
	if len(_protonName) > 0 {
		input.Name = aws.String(_protonName)
	}
	if len(_protonTemplateFile) > 0 {
		input.TemplateFile = aws.String(_protonTemplateFile)
	}
	if len(_protonClientToken) > 0 {
		input.ClientToken = aws.String(_protonClientToken)
	}
	if len(_protonDescription) > 0 {
		input.Description = aws.String(_protonDescription)
	}
	if len(_protonEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_protonEnvironmentName)
	}
	if len(_protonServiceInstanceName) > 0 {
		input.ServiceInstanceName = aws.String(_protonServiceInstanceName)
	}
	if len(_protonServiceName) > 0 {
		input.ServiceName = aws.String(_protonServiceName)
	}
	if len(_protonServiceSpec) > 0 {
		input.ServiceSpec = aws.String(_protonServiceSpec)
	}
	if len(_protonTags) > 0 {
		if err := assignInputField(input, "Tags", _protonTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deploy a new environment. An Proton environment is created from an environment
// template that defines infrastructure and resources that can be shared across
// services.
//
// You can provision environments using the following methods:
//
// - Amazon Web Services-managed provisioning: Proton makes direct calls to
// provision your resources.
//
// - Self-managed provisioning: Proton makes pull requests on your repository to
// provide compiled infrastructure as code (IaC) files that your IaC engine uses to
// provision resources.
//
// For more information, see [Environments] and [Provisioning methods] in the Proton User Guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Environments]: https://docs.aws.amazon.com/proton/latest/userguide/ag-environments.html
// [Provisioning methods]: https://docs.aws.amazon.com/proton/latest/userguide/ag-works-prov-methods.html
func proton_CreateEnvironment(cfg aws.Config, client *proton.Client) {
	input := &proton.CreateEnvironmentInput{
		// Name: *string, // Required
		// Spec: *string, // Required
		// TemplateMajorVersion: *string, // Required
		// TemplateName: *string, // Required
	}

	if len(_protonName) > 0 {
		input.Name = aws.String(_protonName)
	}
	if len(_protonSpec) > 0 {
		input.Spec = aws.String(_protonSpec)
	}
	if len(_protonTemplateMajorVersion) > 0 {
		input.TemplateMajorVersion = aws.String(_protonTemplateMajorVersion)
	}
	if len(_protonTemplateName) > 0 {
		input.TemplateName = aws.String(_protonTemplateName)
	}
	if len(_protonCodebuildRoleArn) > 0 {
		input.CodebuildRoleArn = aws.String(_protonCodebuildRoleArn)
	}
	if len(_protonComponentRoleArn) > 0 {
		input.ComponentRoleArn = aws.String(_protonComponentRoleArn)
	}
	if len(_protonDescription) > 0 {
		input.Description = aws.String(_protonDescription)
	}
	if len(_protonEnvironmentAccountConnectionId) > 0 {
		input.EnvironmentAccountConnectionId = aws.String(_protonEnvironmentAccountConnectionId)
	}
	if len(_protonProtonServiceRoleArn) > 0 {
		input.ProtonServiceRoleArn = aws.String(_protonProtonServiceRoleArn)
	}
	if len(_protonProvisioningRepository) > 0 {
		if err := assignInputField(input, "ProvisioningRepository", _protonProvisioningRepository); err != nil {
			log.Errorf("invalid --provisioning-repository: %s", err.Error())
			return
		}
	}
	if len(_protonTags) > 0 {
		if err := assignInputField(input, "Tags", _protonTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_protonTemplateMinorVersion) > 0 {
		input.TemplateMinorVersion = aws.String(_protonTemplateMinorVersion)
	}

	if resp, err := client.CreateEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create an environment account connection in an environment account so that
// environment infrastructure resources can be provisioned in the environment
// account from a management account.
//
// An environment account connection is a secure bi-directional connection between
// a management account and an environment account that maintains authorization and
// permissions. For more information, see [Environment account connections]in the Proton User guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Environment account connections]: https://docs.aws.amazon.com/proton/latest/userguide/ag-env-account-connections.html
func proton_CreateEnvironmentAccountConnection(cfg aws.Config, client *proton.Client) {
	input := &proton.CreateEnvironmentAccountConnectionInput{
		// EnvironmentName: *string, // Required
		// ManagementAccountId: *string, // Required
	}

	if len(_protonEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_protonEnvironmentName)
	}
	if len(_protonManagementAccountId) > 0 {
		input.ManagementAccountId = aws.String(_protonManagementAccountId)
	}
	if len(_protonClientToken) > 0 {
		input.ClientToken = aws.String(_protonClientToken)
	}
	if len(_protonCodebuildRoleArn) > 0 {
		input.CodebuildRoleArn = aws.String(_protonCodebuildRoleArn)
	}
	if len(_protonComponentRoleArn) > 0 {
		input.ComponentRoleArn = aws.String(_protonComponentRoleArn)
	}
	if len(_protonRoleArn) > 0 {
		input.RoleArn = aws.String(_protonRoleArn)
	}
	if len(_protonTags) > 0 {
		if err := assignInputField(input, "Tags", _protonTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEnvironmentAccountConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create an environment template for Proton. For more information, see [Environment Templates] in the
// Proton User Guide.
//
// You can create an environment template in one of the two following ways:
//
// - Register and publish a standard environment template that instructs Proton
// to deploy and manage environment infrastructure.
//
// - Register and publish a customer managed environment template that connects
// Proton to your existing provisioned infrastructure that you manage. Proton
// doesn't manage your existing provisioned infrastructure. To create an
// environment template for customer provisioned and managed infrastructure,
// include the provisioning parameter and set the value to CUSTOMER_MANAGED . For
// more information, see [Register and publish an environment template]in the Proton User Guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Register and publish an environment template]: https://docs.aws.amazon.com/proton/latest/userguide/template-create.html
// [Environment Templates]: https://docs.aws.amazon.com/proton/latest/userguide/ag-templates.html
func proton_CreateEnvironmentTemplate(cfg aws.Config, client *proton.Client) {
	input := &proton.CreateEnvironmentTemplateInput{
		// Name: *string, // Required
	}

	if len(_protonName) > 0 {
		input.Name = aws.String(_protonName)
	}
	if len(_protonDescription) > 0 {
		input.Description = aws.String(_protonDescription)
	}
	if len(_protonDisplayName) > 0 {
		input.DisplayName = aws.String(_protonDisplayName)
	}
	if len(_protonEncryptionKey) > 0 {
		input.EncryptionKey = aws.String(_protonEncryptionKey)
	}
	if len(_protonProvisioning) > 0 {
		if err := assignInputField(input, "Provisioning", _protonProvisioning); err != nil {
			log.Errorf("invalid --provisioning: %s", err.Error())
			return
		}
	}
	if len(_protonTags) > 0 {
		if err := assignInputField(input, "Tags", _protonTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEnvironmentTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a new major or minor version of an environment template. A major version
// of an environment template is a version that isn't backwards compatible. A minor
// version of an environment template is a version that's backwards compatible
// within its major version.
//
// Deprecated: AWS Proton is not accepting new customers.
func proton_CreateEnvironmentTemplateVersion(cfg aws.Config, client *proton.Client) {
	input := &proton.CreateEnvironmentTemplateVersionInput{
		// Source: types.TemplateVersionSourceInput, // Required
		// TemplateName: *string, // Required
	}

	if len(_protonSource) > 0 {
		if err := assignInputField(input, "Source", _protonSource); err != nil {
			log.Errorf("invalid --source: %s", err.Error())
			return
		}
	}
	if len(_protonTemplateName) > 0 {
		input.TemplateName = aws.String(_protonTemplateName)
	}
	if len(_protonClientToken) > 0 {
		input.ClientToken = aws.String(_protonClientToken)
	}
	if len(_protonDescription) > 0 {
		input.Description = aws.String(_protonDescription)
	}
	if len(_protonMajorVersion) > 0 {
		input.MajorVersion = aws.String(_protonMajorVersion)
	}
	if len(_protonTags) > 0 {
		if err := assignInputField(input, "Tags", _protonTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEnvironmentTemplateVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create and register a link to a repository. Proton uses the link to repeatedly
// access the repository, to either push to it (self-managed provisioning) or pull
// from it (template sync). You can share a linked repository across multiple
// resources (like environments using self-managed provisioning, or synced
// templates). When you create a repository link, Proton creates a [service-linked role]for you.
//
// For more information, see [Self-managed provisioning], [Template bundles], and [Template sync configurations] in the Proton User Guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Self-managed provisioning]: https://docs.aws.amazon.com/proton/latest/userguide/ag-works-prov-methods.html#ag-works-prov-methods-self
// [Template sync configurations]: https://docs.aws.amazon.com/proton/latest/userguide/ag-template-sync-configs.html
// [Template bundles]: https://docs.aws.amazon.com/proton/latest/userguide/ag-template-authoring.html#ag-template-bundles
// [service-linked role]: https://docs.aws.amazon.com/proton/latest/userguide/using-service-linked-roles.html
func proton_CreateRepository(cfg aws.Config, client *proton.Client) {
	input := &proton.CreateRepositoryInput{
		// ConnectionArn: *string, // Required
		// Name: *string, // Required
		// Provider: types.RepositoryProvider, // Required
	}

	if len(_protonConnectionArn) > 0 {
		input.ConnectionArn = aws.String(_protonConnectionArn)
	}
	if len(_protonName) > 0 {
		input.Name = aws.String(_protonName)
	}
	if len(_protonProvider) > 0 {
		if err := assignInputField(input, "Provider", _protonProvider); err != nil {
			log.Errorf("invalid --provider: %s", err.Error())
			return
		}
	}
	if len(_protonEncryptionKey) > 0 {
		input.EncryptionKey = aws.String(_protonEncryptionKey)
	}
	if len(_protonTags) > 0 {
		if err := assignInputField(input, "Tags", _protonTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRepository(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create an Proton service. An Proton service is an instantiation of a service
// template and often includes several service instances and pipeline. For more
// information, see [Services]in the Proton User Guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Services]: https://docs.aws.amazon.com/proton/latest/userguide/ag-services.html
func proton_CreateService(cfg aws.Config, client *proton.Client) {
	input := &proton.CreateServiceInput{
		// Name: *string, // Required
		// Spec: *string, // Required
		// TemplateMajorVersion: *string, // Required
		// TemplateName: *string, // Required
	}

	if len(_protonName) > 0 {
		input.Name = aws.String(_protonName)
	}
	if len(_protonSpec) > 0 {
		input.Spec = aws.String(_protonSpec)
	}
	if len(_protonTemplateMajorVersion) > 0 {
		input.TemplateMajorVersion = aws.String(_protonTemplateMajorVersion)
	}
	if len(_protonTemplateName) > 0 {
		input.TemplateName = aws.String(_protonTemplateName)
	}
	if len(_protonBranchName) > 0 {
		input.BranchName = aws.String(_protonBranchName)
	}
	if len(_protonDescription) > 0 {
		input.Description = aws.String(_protonDescription)
	}
	if len(_protonRepositoryConnectionArn) > 0 {
		input.RepositoryConnectionArn = aws.String(_protonRepositoryConnectionArn)
	}
	if len(_protonRepositoryId) > 0 {
		input.RepositoryId = aws.String(_protonRepositoryId)
	}
	if len(_protonTags) > 0 {
		if err := assignInputField(input, "Tags", _protonTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_protonTemplateMinorVersion) > 0 {
		input.TemplateMinorVersion = aws.String(_protonTemplateMinorVersion)
	}

	if resp, err := client.CreateService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a service instance.
// Deprecated: AWS Proton is not accepting new customers.
func proton_CreateServiceInstance(cfg aws.Config, client *proton.Client) {
	input := &proton.CreateServiceInstanceInput{
		// Name: *string, // Required
		// ServiceName: *string, // Required
		// Spec: *string, // Required
	}

	if len(_protonName) > 0 {
		input.Name = aws.String(_protonName)
	}
	if len(_protonServiceName) > 0 {
		input.ServiceName = aws.String(_protonServiceName)
	}
	if len(_protonSpec) > 0 {
		input.Spec = aws.String(_protonSpec)
	}
	if len(_protonClientToken) > 0 {
		input.ClientToken = aws.String(_protonClientToken)
	}
	if len(_protonTags) > 0 {
		if err := assignInputField(input, "Tags", _protonTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_protonTemplateMajorVersion) > 0 {
		input.TemplateMajorVersion = aws.String(_protonTemplateMajorVersion)
	}
	if len(_protonTemplateMinorVersion) > 0 {
		input.TemplateMinorVersion = aws.String(_protonTemplateMinorVersion)
	}

	if resp, err := client.CreateServiceInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create the Proton Ops configuration file.
// Deprecated: AWS Proton is not accepting new customers.
func proton_CreateServiceSyncConfig(cfg aws.Config, client *proton.Client) {
	input := &proton.CreateServiceSyncConfigInput{
		// Branch: *string, // Required
		// FilePath: *string, // Required
		// RepositoryName: *string, // Required
		// RepositoryProvider: types.RepositoryProvider, // Required
		// ServiceName: *string, // Required
	}

	if len(_protonBranch) > 0 {
		input.Branch = aws.String(_protonBranch)
	}
	if len(_protonFilePath) > 0 {
		input.FilePath = aws.String(_protonFilePath)
	}
	if len(_protonRepositoryName) > 0 {
		input.RepositoryName = aws.String(_protonRepositoryName)
	}
	if len(_protonRepositoryProvider) > 0 {
		if err := assignInputField(input, "RepositoryProvider", _protonRepositoryProvider); err != nil {
			log.Errorf("invalid --repository-provider: %s", err.Error())
			return
		}
	}
	if len(_protonServiceName) > 0 {
		input.ServiceName = aws.String(_protonServiceName)
	}

	if resp, err := client.CreateServiceSyncConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a service template. The administrator creates a service template to
// define standardized infrastructure and an optional CI/CD service pipeline.
// Developers, in turn, select the service template from Proton. If the selected
// service template includes a service pipeline definition, they provide a link to
// their source code repository. Proton then deploys and manages the infrastructure
// defined by the selected service template. For more information, see [Proton templates]in the
// Proton User Guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Proton templates]: https://docs.aws.amazon.com/proton/latest/userguide/ag-templates.html
func proton_CreateServiceTemplate(cfg aws.Config, client *proton.Client) {
	input := &proton.CreateServiceTemplateInput{
		// Name: *string, // Required
	}

	if len(_protonName) > 0 {
		input.Name = aws.String(_protonName)
	}
	if len(_protonDescription) > 0 {
		input.Description = aws.String(_protonDescription)
	}
	if len(_protonDisplayName) > 0 {
		input.DisplayName = aws.String(_protonDisplayName)
	}
	if len(_protonEncryptionKey) > 0 {
		input.EncryptionKey = aws.String(_protonEncryptionKey)
	}
	if len(_protonPipelineProvisioning) > 0 {
		if err := assignInputField(input, "PipelineProvisioning", _protonPipelineProvisioning); err != nil {
			log.Errorf("invalid --pipeline-provisioning: %s", err.Error())
			return
		}
	}
	if len(_protonTags) > 0 {
		if err := assignInputField(input, "Tags", _protonTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateServiceTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a new major or minor version of a service template. A major version of a
// service template is a version that isn't backward compatible. A minor version of
// a service template is a version that's backward compatible within its major
// version.
//
// Deprecated: AWS Proton is not accepting new customers.
func proton_CreateServiceTemplateVersion(cfg aws.Config, client *proton.Client) {
	input := &proton.CreateServiceTemplateVersionInput{
		// CompatibleEnvironmentTemplates: []types.CompatibleEnvironmentTemplateInput, // Required
		// Source: types.TemplateVersionSourceInput, // Required
		// TemplateName: *string, // Required
	}

	if len(_protonCompatibleEnvironmentTemplates) > 0 {
		if err := assignInputField(input, "CompatibleEnvironmentTemplates", _protonCompatibleEnvironmentTemplates); err != nil {
			log.Errorf("invalid --compatible-environment-templates: %s", err.Error())
			return
		}
	}
	if len(_protonSource) > 0 {
		if err := assignInputField(input, "Source", _protonSource); err != nil {
			log.Errorf("invalid --source: %s", err.Error())
			return
		}
	}
	if len(_protonTemplateName) > 0 {
		input.TemplateName = aws.String(_protonTemplateName)
	}
	if len(_protonClientToken) > 0 {
		input.ClientToken = aws.String(_protonClientToken)
	}
	if len(_protonDescription) > 0 {
		input.Description = aws.String(_protonDescription)
	}
	if len(_protonMajorVersion) > 0 {
		input.MajorVersion = aws.String(_protonMajorVersion)
	}
	if len(_protonSupportedComponentSources) > 0 {
		if err := assignInputField(input, "SupportedComponentSources", _protonSupportedComponentSources); err != nil {
			log.Errorf("invalid --supported-component-sources: %s", err.Error())
			return
		}
	}
	if len(_protonTags) > 0 {
		if err := assignInputField(input, "Tags", _protonTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateServiceTemplateVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Set up a template to create new template versions automatically by tracking a
// linked repository. A linked repository is a repository that has been registered
// with Proton. For more information, see CreateRepository.
//
// When a commit is pushed to your linked repository, Proton checks for changes to
// your repository template bundles. If it detects a template bundle change, a new
// major or minor version of its template is created, if the version doesn’t
// already exist. For more information, see [Template sync configurations]in the Proton User Guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Template sync configurations]: https://docs.aws.amazon.com/proton/latest/userguide/ag-template-sync-configs.html
func proton_CreateTemplateSyncConfig(cfg aws.Config, client *proton.Client) {
	input := &proton.CreateTemplateSyncConfigInput{
		// Branch: *string, // Required
		// RepositoryName: *string, // Required
		// RepositoryProvider: types.RepositoryProvider, // Required
		// TemplateName: *string, // Required
		// TemplateType: types.TemplateType, // Required
	}

	if len(_protonBranch) > 0 {
		input.Branch = aws.String(_protonBranch)
	}
	if len(_protonRepositoryName) > 0 {
		input.RepositoryName = aws.String(_protonRepositoryName)
	}
	if len(_protonRepositoryProvider) > 0 {
		if err := assignInputField(input, "RepositoryProvider", _protonRepositoryProvider); err != nil {
			log.Errorf("invalid --repository-provider: %s", err.Error())
			return
		}
	}
	if len(_protonTemplateName) > 0 {
		input.TemplateName = aws.String(_protonTemplateName)
	}
	if len(_protonTemplateType) > 0 {
		if err := assignInputField(input, "TemplateType", _protonTemplateType); err != nil {
			log.Errorf("invalid --template-type: %s", err.Error())
			return
		}
	}
	if len(_protonSubdirectory) > 0 {
		input.Subdirectory = aws.String(_protonSubdirectory)
	}

	if resp, err := client.CreateTemplateSyncConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an Proton component resource.
// For more information about components, see [Proton components] in the Proton User Guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Proton components]: https://docs.aws.amazon.com/proton/latest/userguide/ag-components.html
func proton_DeleteComponent(cfg aws.Config, client *proton.Client) {
	input := &proton.DeleteComponentInput{
		// Name: *string, // Required
	}

	if len(_protonName) > 0 {
		input.Name = aws.String(_protonName)
	}

	if resp, err := client.DeleteComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete the deployment.
// Deprecated: AWS Proton is not accepting new customers.
func proton_DeleteDeployment(cfg aws.Config, client *proton.Client) {
	input := &proton.DeleteDeploymentInput{
		// Id: *string, // Required
	}

	if len(_protonId) > 0 {
		input.Id = aws.String(_protonId)
	}

	if resp, err := client.DeleteDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an environment.
// Deprecated: AWS Proton is not accepting new customers.
func proton_DeleteEnvironment(cfg aws.Config, client *proton.Client) {
	input := &proton.DeleteEnvironmentInput{
		// Name: *string, // Required
	}

	if len(_protonName) > 0 {
		input.Name = aws.String(_protonName)
	}

	if resp, err := client.DeleteEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// In an environment account, delete an environment account connection.
// After you delete an environment account connection that’s in use by an Proton
// environment, Proton can’t manage the environment infrastructure resources until
// a new environment account connection is accepted for the environment account and
// associated environment. You're responsible for cleaning up provisioned resources
// that remain without an environment connection.
//
// For more information, see [Environment account connections] in the Proton User guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Environment account connections]: https://docs.aws.amazon.com/proton/latest/userguide/ag-env-account-connections.html
func proton_DeleteEnvironmentAccountConnection(cfg aws.Config, client *proton.Client) {
	input := &proton.DeleteEnvironmentAccountConnectionInput{
		// Id: *string, // Required
	}

	if len(_protonId) > 0 {
		input.Id = aws.String(_protonId)
	}

	if resp, err := client.DeleteEnvironmentAccountConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// If no other major or minor versions of an environment template exist, delete
// the environment template.
//
// Deprecated: AWS Proton is not accepting new customers.
func proton_DeleteEnvironmentTemplate(cfg aws.Config, client *proton.Client) {
	input := &proton.DeleteEnvironmentTemplateInput{
		// Name: *string, // Required
	}

	if len(_protonName) > 0 {
		input.Name = aws.String(_protonName)
	}

	if resp, err := client.DeleteEnvironmentTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// If no other minor versions of an environment template exist, delete a major
// version of the environment template if it's not the Recommended version. Delete
// the Recommended version of the environment template if no other major versions
// or minor versions of the environment template exist. A major version of an
// environment template is a version that's not backward compatible.
//
// Delete a minor version of an environment template if it isn't the Recommended
// version. Delete a Recommended minor version of the environment template if no
// other minor versions of the environment template exist. A minor version of an
// environment template is a version that's backward compatible.
//
// Deprecated: AWS Proton is not accepting new customers.
func proton_DeleteEnvironmentTemplateVersion(cfg aws.Config, client *proton.Client) {
	input := &proton.DeleteEnvironmentTemplateVersionInput{
		// MajorVersion: *string, // Required
		// MinorVersion: *string, // Required
		// TemplateName: *string, // Required
	}

	if len(_protonMajorVersion) > 0 {
		input.MajorVersion = aws.String(_protonMajorVersion)
	}
	if len(_protonMinorVersion) > 0 {
		input.MinorVersion = aws.String(_protonMinorVersion)
	}
	if len(_protonTemplateName) > 0 {
		input.TemplateName = aws.String(_protonTemplateName)
	}

	if resp, err := client.DeleteEnvironmentTemplateVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// De-register and unlink your repository.
// Deprecated: AWS Proton is not accepting new customers.
func proton_DeleteRepository(cfg aws.Config, client *proton.Client) {
	input := &proton.DeleteRepositoryInput{
		// Name: *string, // Required
		// Provider: types.RepositoryProvider, // Required
	}

	if len(_protonName) > 0 {
		input.Name = aws.String(_protonName)
	}
	if len(_protonProvider) > 0 {
		if err := assignInputField(input, "Provider", _protonProvider); err != nil {
			log.Errorf("invalid --provider: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteRepository(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a service, with its instances and pipeline.
// You can't delete a service if it has any service instances that have components
// attached to them.
//
// For more information about components, see [Proton components] in the Proton User Guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Proton components]: https://docs.aws.amazon.com/proton/latest/userguide/ag-components.html
func proton_DeleteService(cfg aws.Config, client *proton.Client) {
	input := &proton.DeleteServiceInput{
		// Name: *string, // Required
	}

	if len(_protonName) > 0 {
		input.Name = aws.String(_protonName)
	}

	if resp, err := client.DeleteService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete the Proton Ops file.
// Deprecated: AWS Proton is not accepting new customers.
func proton_DeleteServiceSyncConfig(cfg aws.Config, client *proton.Client) {
	input := &proton.DeleteServiceSyncConfigInput{
		// ServiceName: *string, // Required
	}

	if len(_protonServiceName) > 0 {
		input.ServiceName = aws.String(_protonServiceName)
	}

	if resp, err := client.DeleteServiceSyncConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// If no other major or minor versions of the service template exist, delete the
// service template.
//
// Deprecated: AWS Proton is not accepting new customers.
func proton_DeleteServiceTemplate(cfg aws.Config, client *proton.Client) {
	input := &proton.DeleteServiceTemplateInput{
		// Name: *string, // Required
	}

	if len(_protonName) > 0 {
		input.Name = aws.String(_protonName)
	}

	if resp, err := client.DeleteServiceTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// If no other minor versions of a service template exist, delete a major version
// of the service template if it's not the Recommended version. Delete the
// Recommended version of the service template if no other major versions or minor
// versions of the service template exist. A major version of a service template is
// a version that isn't backwards compatible.
//
// Delete a minor version of a service template if it's not the Recommended
// version. Delete a Recommended minor version of the service template if no other
// minor versions of the service template exist. A minor version of a service
// template is a version that's backwards compatible.
//
// Deprecated: AWS Proton is not accepting new customers.
func proton_DeleteServiceTemplateVersion(cfg aws.Config, client *proton.Client) {
	input := &proton.DeleteServiceTemplateVersionInput{
		// MajorVersion: *string, // Required
		// MinorVersion: *string, // Required
		// TemplateName: *string, // Required
	}

	if len(_protonMajorVersion) > 0 {
		input.MajorVersion = aws.String(_protonMajorVersion)
	}
	if len(_protonMinorVersion) > 0 {
		input.MinorVersion = aws.String(_protonMinorVersion)
	}
	if len(_protonTemplateName) > 0 {
		input.TemplateName = aws.String(_protonTemplateName)
	}

	if resp, err := client.DeleteServiceTemplateVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a template sync configuration.
// Deprecated: AWS Proton is not accepting new customers.
func proton_DeleteTemplateSyncConfig(cfg aws.Config, client *proton.Client) {
	input := &proton.DeleteTemplateSyncConfigInput{
		// TemplateName: *string, // Required
		// TemplateType: types.TemplateType, // Required
	}

	if len(_protonTemplateName) > 0 {
		input.TemplateName = aws.String(_protonTemplateName)
	}
	if len(_protonTemplateType) > 0 {
		if err := assignInputField(input, "TemplateType", _protonTemplateType); err != nil {
			log.Errorf("invalid --template-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteTemplateSyncConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get detail data for Proton account-wide settings.
// Deprecated: AWS Proton is not accepting new customers.
func proton_GetAccountSettings(cfg aws.Config, client *proton.Client) {
	input := &proton.GetAccountSettingsInput{}

	if resp, err := client.GetAccountSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get detailed data for a component.
// For more information about components, see [Proton components] in the Proton User Guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Proton components]: https://docs.aws.amazon.com/proton/latest/userguide/ag-components.html
func proton_GetComponent(cfg aws.Config, client *proton.Client) {
	input := &proton.GetComponentInput{
		// Name: *string, // Required
	}

	if len(_protonName) > 0 {
		input.Name = aws.String(_protonName)
	}

	if resp, err := client.GetComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get detailed data for a deployment.
// Deprecated: AWS Proton is not accepting new customers.
func proton_GetDeployment(cfg aws.Config, client *proton.Client) {
	input := &proton.GetDeploymentInput{
		// Id: *string, // Required
	}

	if len(_protonId) > 0 {
		input.Id = aws.String(_protonId)
	}
	if len(_protonComponentName) > 0 {
		input.ComponentName = aws.String(_protonComponentName)
	}
	if len(_protonEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_protonEnvironmentName)
	}
	if len(_protonServiceInstanceName) > 0 {
		input.ServiceInstanceName = aws.String(_protonServiceInstanceName)
	}
	if len(_protonServiceName) > 0 {
		input.ServiceName = aws.String(_protonServiceName)
	}

	if resp, err := client.GetDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get detailed data for an environment.
// Deprecated: AWS Proton is not accepting new customers.
func proton_GetEnvironment(cfg aws.Config, client *proton.Client) {
	input := &proton.GetEnvironmentInput{
		// Name: *string, // Required
	}

	if len(_protonName) > 0 {
		input.Name = aws.String(_protonName)
	}

	if resp, err := client.GetEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// In an environment account, get the detailed data for an environment account
// connection.
//
// For more information, see [Environment account connections] in the Proton User guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Environment account connections]: https://docs.aws.amazon.com/proton/latest/userguide/ag-env-account-connections.html
func proton_GetEnvironmentAccountConnection(cfg aws.Config, client *proton.Client) {
	input := &proton.GetEnvironmentAccountConnectionInput{
		// Id: *string, // Required
	}

	if len(_protonId) > 0 {
		input.Id = aws.String(_protonId)
	}

	if resp, err := client.GetEnvironmentAccountConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get detailed data for an environment template.
// Deprecated: AWS Proton is not accepting new customers.
func proton_GetEnvironmentTemplate(cfg aws.Config, client *proton.Client) {
	input := &proton.GetEnvironmentTemplateInput{
		// Name: *string, // Required
	}

	if len(_protonName) > 0 {
		input.Name = aws.String(_protonName)
	}

	if resp, err := client.GetEnvironmentTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get detailed data for a major or minor version of an environment template.
// Deprecated: AWS Proton is not accepting new customers.
func proton_GetEnvironmentTemplateVersion(cfg aws.Config, client *proton.Client) {
	input := &proton.GetEnvironmentTemplateVersionInput{
		// MajorVersion: *string, // Required
		// MinorVersion: *string, // Required
		// TemplateName: *string, // Required
	}

	if len(_protonMajorVersion) > 0 {
		input.MajorVersion = aws.String(_protonMajorVersion)
	}
	if len(_protonMinorVersion) > 0 {
		input.MinorVersion = aws.String(_protonMinorVersion)
	}
	if len(_protonTemplateName) > 0 {
		input.TemplateName = aws.String(_protonTemplateName)
	}

	if resp, err := client.GetEnvironmentTemplateVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get detail data for a linked repository.
// Deprecated: AWS Proton is not accepting new customers.
func proton_GetRepository(cfg aws.Config, client *proton.Client) {
	input := &proton.GetRepositoryInput{
		// Name: *string, // Required
		// Provider: types.RepositoryProvider, // Required
	}

	if len(_protonName) > 0 {
		input.Name = aws.String(_protonName)
	}
	if len(_protonProvider) > 0 {
		if err := assignInputField(input, "Provider", _protonProvider); err != nil {
			log.Errorf("invalid --provider: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetRepository(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the sync status of a repository used for Proton template sync. For more
// information about template sync, see .
//
// A repository sync status isn't tied to the Proton Repository resource (or any
// other Proton resource). Therefore, tags on an Proton Repository resource have no
// effect on this action. Specifically, you can't use these tags to control access
// to this action using Attribute-based access control (ABAC).
//
// For more information about ABAC, see [ABAC] in the Proton User Guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [ABAC]: https://docs.aws.amazon.com/proton/latest/userguide/security_iam_service-with-iam.html#security_iam_service-with-iam-tags
func proton_GetRepositorySyncStatus(cfg aws.Config, client *proton.Client) {
	input := &proton.GetRepositorySyncStatusInput{
		// Branch: *string, // Required
		// RepositoryName: *string, // Required
		// RepositoryProvider: types.RepositoryProvider, // Required
		// SyncType: types.SyncType, // Required
	}

	if len(_protonBranch) > 0 {
		input.Branch = aws.String(_protonBranch)
	}
	if len(_protonRepositoryName) > 0 {
		input.RepositoryName = aws.String(_protonRepositoryName)
	}
	if len(_protonRepositoryProvider) > 0 {
		if err := assignInputField(input, "RepositoryProvider", _protonRepositoryProvider); err != nil {
			log.Errorf("invalid --repository-provider: %s", err.Error())
			return
		}
	}
	if len(_protonSyncType) > 0 {
		if err := assignInputField(input, "SyncType", _protonSyncType); err != nil {
			log.Errorf("invalid --sync-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetRepositorySyncStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get counts of Proton resources.
// For infrastructure-provisioning resources (environments, services, service
// instances, pipelines), the action returns staleness counts. A resource is stale
// when it's behind the recommended version of the Proton template that it uses and
// it needs an update to become current.
//
// The action returns staleness counts (counts of resources that are up-to-date,
// behind a template major version, or behind a template minor version), the total
// number of resources, and the number of resources that are in a failed state,
// grouped by resource type. Components, environments, and service templates return
// less information - see the components , environments , and serviceTemplates
// field descriptions.
//
// For context, the action also returns the total number of each type of Proton
// template in the Amazon Web Services account.
//
// For more information, see [Proton dashboard] in the Proton User Guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Proton dashboard]: https://docs.aws.amazon.com/proton/latest/userguide/monitoring-dashboard.html
func proton_GetResourcesSummary(cfg aws.Config, client *proton.Client) {
	input := &proton.GetResourcesSummaryInput{}

	if resp, err := client.GetResourcesSummary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get detailed data for a service.
// Deprecated: AWS Proton is not accepting new customers.
func proton_GetService(cfg aws.Config, client *proton.Client) {
	input := &proton.GetServiceInput{
		// Name: *string, // Required
	}

	if len(_protonName) > 0 {
		input.Name = aws.String(_protonName)
	}

	if resp, err := client.GetService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get detailed data for a service instance. A service instance is an
// instantiation of service template and it runs in a specific environment.
//
// Deprecated: AWS Proton is not accepting new customers.
func proton_GetServiceInstance(cfg aws.Config, client *proton.Client) {
	input := &proton.GetServiceInstanceInput{
		// Name: *string, // Required
		// ServiceName: *string, // Required
	}

	if len(_protonName) > 0 {
		input.Name = aws.String(_protonName)
	}
	if len(_protonServiceName) > 0 {
		input.ServiceName = aws.String(_protonServiceName)
	}

	if resp, err := client.GetServiceInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the status of the synced service instance.
// Deprecated: AWS Proton is not accepting new customers.
func proton_GetServiceInstanceSyncStatus(cfg aws.Config, client *proton.Client) {
	input := &proton.GetServiceInstanceSyncStatusInput{
		// ServiceInstanceName: *string, // Required
		// ServiceName: *string, // Required
	}

	if len(_protonServiceInstanceName) > 0 {
		input.ServiceInstanceName = aws.String(_protonServiceInstanceName)
	}
	if len(_protonServiceName) > 0 {
		input.ServiceName = aws.String(_protonServiceName)
	}

	if resp, err := client.GetServiceInstanceSyncStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get detailed data for the service sync blocker summary.
// Deprecated: AWS Proton is not accepting new customers.
func proton_GetServiceSyncBlockerSummary(cfg aws.Config, client *proton.Client) {
	input := &proton.GetServiceSyncBlockerSummaryInput{
		// ServiceName: *string, // Required
	}

	if len(_protonServiceName) > 0 {
		input.ServiceName = aws.String(_protonServiceName)
	}
	if len(_protonServiceInstanceName) > 0 {
		input.ServiceInstanceName = aws.String(_protonServiceInstanceName)
	}

	if resp, err := client.GetServiceSyncBlockerSummary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get detailed information for the service sync configuration.
// Deprecated: AWS Proton is not accepting new customers.
func proton_GetServiceSyncConfig(cfg aws.Config, client *proton.Client) {
	input := &proton.GetServiceSyncConfigInput{
		// ServiceName: *string, // Required
	}

	if len(_protonServiceName) > 0 {
		input.ServiceName = aws.String(_protonServiceName)
	}

	if resp, err := client.GetServiceSyncConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get detailed data for a service template.
// Deprecated: AWS Proton is not accepting new customers.
func proton_GetServiceTemplate(cfg aws.Config, client *proton.Client) {
	input := &proton.GetServiceTemplateInput{
		// Name: *string, // Required
	}

	if len(_protonName) > 0 {
		input.Name = aws.String(_protonName)
	}

	if resp, err := client.GetServiceTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get detailed data for a major or minor version of a service template.
// Deprecated: AWS Proton is not accepting new customers.
func proton_GetServiceTemplateVersion(cfg aws.Config, client *proton.Client) {
	input := &proton.GetServiceTemplateVersionInput{
		// MajorVersion: *string, // Required
		// MinorVersion: *string, // Required
		// TemplateName: *string, // Required
	}

	if len(_protonMajorVersion) > 0 {
		input.MajorVersion = aws.String(_protonMajorVersion)
	}
	if len(_protonMinorVersion) > 0 {
		input.MinorVersion = aws.String(_protonMinorVersion)
	}
	if len(_protonTemplateName) > 0 {
		input.TemplateName = aws.String(_protonTemplateName)
	}

	if resp, err := client.GetServiceTemplateVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get detail data for a template sync configuration.
// Deprecated: AWS Proton is not accepting new customers.
func proton_GetTemplateSyncConfig(cfg aws.Config, client *proton.Client) {
	input := &proton.GetTemplateSyncConfigInput{
		// TemplateName: *string, // Required
		// TemplateType: types.TemplateType, // Required
	}

	if len(_protonTemplateName) > 0 {
		input.TemplateName = aws.String(_protonTemplateName)
	}
	if len(_protonTemplateType) > 0 {
		if err := assignInputField(input, "TemplateType", _protonTemplateType); err != nil {
			log.Errorf("invalid --template-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetTemplateSyncConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the status of a template sync.
// Deprecated: AWS Proton is not accepting new customers.
func proton_GetTemplateSyncStatus(cfg aws.Config, client *proton.Client) {
	input := &proton.GetTemplateSyncStatusInput{
		// TemplateName: *string, // Required
		// TemplateType: types.TemplateType, // Required
		// TemplateVersion: *string, // Required
	}

	if len(_protonTemplateName) > 0 {
		input.TemplateName = aws.String(_protonTemplateName)
	}
	if len(_protonTemplateType) > 0 {
		if err := assignInputField(input, "TemplateType", _protonTemplateType); err != nil {
			log.Errorf("invalid --template-type: %s", err.Error())
			return
		}
	}
	if len(_protonTemplateVersion) > 0 {
		input.TemplateVersion = aws.String(_protonTemplateVersion)
	}

	if resp, err := client.GetTemplateSyncStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get a list of component Infrastructure as Code (IaC) outputs.
// For more information about components, see [Proton components] in the Proton User Guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Proton components]: https://docs.aws.amazon.com/proton/latest/userguide/ag-components.html
func proton_ListComponentOutputs(cfg aws.Config, client *proton.Client) {
	input := &proton.ListComponentOutputsInput{
		// ComponentName: *string, // Required
	}

	if len(_protonComponentName) > 0 {
		input.ComponentName = aws.String(_protonComponentName)
	}
	if len(_protonDeploymentId) > 0 {
		input.DeploymentId = aws.String(_protonDeploymentId)
	}
	if len(_protonNextToken) > 0 {
		input.NextToken = aws.String(_protonNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListComponentOutputs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*proton.ListComponentOutputsOutput
	p := proton.NewListComponentOutputsPaginator(client, input)
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

// List provisioned resources for a component with details.
// For more information about components, see [Proton components] in the Proton User Guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Proton components]: https://docs.aws.amazon.com/proton/latest/userguide/ag-components.html
func proton_ListComponentProvisionedResources(cfg aws.Config, client *proton.Client) {
	input := &proton.ListComponentProvisionedResourcesInput{
		// ComponentName: *string, // Required
	}

	if len(_protonComponentName) > 0 {
		input.ComponentName = aws.String(_protonComponentName)
	}
	if len(_protonNextToken) > 0 {
		input.NextToken = aws.String(_protonNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListComponentProvisionedResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*proton.ListComponentProvisionedResourcesOutput
	p := proton.NewListComponentProvisionedResourcesPaginator(client, input)
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

// List components with summary data. You can filter the result list by
// environment, service, or a single service instance.
//
// For more information about components, see [Proton components] in the Proton User Guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Proton components]: https://docs.aws.amazon.com/proton/latest/userguide/ag-components.html
func proton_ListComponents(cfg aws.Config, client *proton.Client) {
	input := &proton.ListComponentsInput{}

	if len(_protonEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_protonEnvironmentName)
	}
	if len(_protonMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _protonMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_protonNextToken) > 0 {
		input.NextToken = aws.String(_protonNextToken)
	}
	if len(_protonServiceInstanceName) > 0 {
		input.ServiceInstanceName = aws.String(_protonServiceInstanceName)
	}
	if len(_protonServiceName) > 0 {
		input.ServiceName = aws.String(_protonServiceName)
	}

	if disablePaginator() {
		if resp, err := client.ListComponents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*proton.ListComponentsOutput
	p := proton.NewListComponentsPaginator(client, input)
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

// List deployments. You can filter the result list by environment, service, or a
// single service instance.
//
// Deprecated: AWS Proton is not accepting new customers.
func proton_ListDeployments(cfg aws.Config, client *proton.Client) {
	input := &proton.ListDeploymentsInput{}

	if len(_protonComponentName) > 0 {
		input.ComponentName = aws.String(_protonComponentName)
	}
	if len(_protonEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_protonEnvironmentName)
	}
	if len(_protonMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _protonMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_protonNextToken) > 0 {
		input.NextToken = aws.String(_protonNextToken)
	}
	if len(_protonServiceInstanceName) > 0 {
		input.ServiceInstanceName = aws.String(_protonServiceInstanceName)
	}
	if len(_protonServiceName) > 0 {
		input.ServiceName = aws.String(_protonServiceName)
	}

	if disablePaginator() {
		if resp, err := client.ListDeployments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*proton.ListDeploymentsOutput
	p := proton.NewListDeploymentsPaginator(client, input)
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

// View a list of environment account connections.
// For more information, see [Environment account connections] in the Proton User guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Environment account connections]: https://docs.aws.amazon.com/proton/latest/userguide/ag-env-account-connections.html
func proton_ListEnvironmentAccountConnections(cfg aws.Config, client *proton.Client) {
	input := &proton.ListEnvironmentAccountConnectionsInput{
		// RequestedBy: types.EnvironmentAccountConnectionRequesterAccountType, // Required
	}

	if len(_protonRequestedBy) > 0 {
		if err := assignInputField(input, "RequestedBy", _protonRequestedBy); err != nil {
			log.Errorf("invalid --requested-by: %s", err.Error())
			return
		}
	}
	if len(_protonEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_protonEnvironmentName)
	}
	if len(_protonMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _protonMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_protonNextToken) > 0 {
		input.NextToken = aws.String(_protonNextToken)
	}
	if len(_protonStatuses) > 0 {
		if err := assignInputField(input, "Statuses", _protonStatuses); err != nil {
			log.Errorf("invalid --statuses: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListEnvironmentAccountConnections(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*proton.ListEnvironmentAccountConnectionsOutput
	p := proton.NewListEnvironmentAccountConnectionsPaginator(client, input)
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

// List the infrastructure as code outputs for your environment.
// Deprecated: AWS Proton is not accepting new customers.
func proton_ListEnvironmentOutputs(cfg aws.Config, client *proton.Client) {
	input := &proton.ListEnvironmentOutputsInput{
		// EnvironmentName: *string, // Required
	}

	if len(_protonEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_protonEnvironmentName)
	}
	if len(_protonDeploymentId) > 0 {
		input.DeploymentId = aws.String(_protonDeploymentId)
	}
	if len(_protonNextToken) > 0 {
		input.NextToken = aws.String(_protonNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEnvironmentOutputs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*proton.ListEnvironmentOutputsOutput
	p := proton.NewListEnvironmentOutputsPaginator(client, input)
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

// List the provisioned resources for your environment.
// Deprecated: AWS Proton is not accepting new customers.
func proton_ListEnvironmentProvisionedResources(cfg aws.Config, client *proton.Client) {
	input := &proton.ListEnvironmentProvisionedResourcesInput{
		// EnvironmentName: *string, // Required
	}

	if len(_protonEnvironmentName) > 0 {
		input.EnvironmentName = aws.String(_protonEnvironmentName)
	}
	if len(_protonNextToken) > 0 {
		input.NextToken = aws.String(_protonNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEnvironmentProvisionedResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*proton.ListEnvironmentProvisionedResourcesOutput
	p := proton.NewListEnvironmentProvisionedResourcesPaginator(client, input)
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

// List major or minor versions of an environment template with detail data.
// Deprecated: AWS Proton is not accepting new customers.
func proton_ListEnvironmentTemplateVersions(cfg aws.Config, client *proton.Client) {
	input := &proton.ListEnvironmentTemplateVersionsInput{
		// TemplateName: *string, // Required
	}

	if len(_protonTemplateName) > 0 {
		input.TemplateName = aws.String(_protonTemplateName)
	}
	if len(_protonMajorVersion) > 0 {
		input.MajorVersion = aws.String(_protonMajorVersion)
	}
	if len(_protonMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _protonMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_protonNextToken) > 0 {
		input.NextToken = aws.String(_protonNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEnvironmentTemplateVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*proton.ListEnvironmentTemplateVersionsOutput
	p := proton.NewListEnvironmentTemplateVersionsPaginator(client, input)
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

// List environment templates.
// Deprecated: AWS Proton is not accepting new customers.
func proton_ListEnvironmentTemplates(cfg aws.Config, client *proton.Client) {
	input := &proton.ListEnvironmentTemplatesInput{}

	if len(_protonMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _protonMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_protonNextToken) > 0 {
		input.NextToken = aws.String(_protonNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEnvironmentTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*proton.ListEnvironmentTemplatesOutput
	p := proton.NewListEnvironmentTemplatesPaginator(client, input)
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

// List environments with detail data summaries.
// Deprecated: AWS Proton is not accepting new customers.
func proton_ListEnvironments(cfg aws.Config, client *proton.Client) {
	input := &proton.ListEnvironmentsInput{}

	if len(_protonEnvironmentTemplates) > 0 {
		if err := assignInputField(input, "EnvironmentTemplates", _protonEnvironmentTemplates); err != nil {
			log.Errorf("invalid --environment-templates: %s", err.Error())
			return
		}
	}
	if len(_protonMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _protonMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_protonNextToken) > 0 {
		input.NextToken = aws.String(_protonNextToken)
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

	var results []*proton.ListEnvironmentsOutput
	p := proton.NewListEnvironmentsPaginator(client, input)
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

// List linked repositories with detail data.
// Deprecated: AWS Proton is not accepting new customers.
func proton_ListRepositories(cfg aws.Config, client *proton.Client) {
	input := &proton.ListRepositoriesInput{}

	if len(_protonMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _protonMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_protonNextToken) > 0 {
		input.NextToken = aws.String(_protonNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRepositories(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*proton.ListRepositoriesOutput
	p := proton.NewListRepositoriesPaginator(client, input)
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

// List repository sync definitions with detail data.
// Deprecated: AWS Proton is not accepting new customers.
func proton_ListRepositorySyncDefinitions(cfg aws.Config, client *proton.Client) {
	input := &proton.ListRepositorySyncDefinitionsInput{
		// RepositoryName: *string, // Required
		// RepositoryProvider: types.RepositoryProvider, // Required
		// SyncType: types.SyncType, // Required
	}

	if len(_protonRepositoryName) > 0 {
		input.RepositoryName = aws.String(_protonRepositoryName)
	}
	if len(_protonRepositoryProvider) > 0 {
		if err := assignInputField(input, "RepositoryProvider", _protonRepositoryProvider); err != nil {
			log.Errorf("invalid --repository-provider: %s", err.Error())
			return
		}
	}
	if len(_protonSyncType) > 0 {
		if err := assignInputField(input, "SyncType", _protonSyncType); err != nil {
			log.Errorf("invalid --sync-type: %s", err.Error())
			return
		}
	}
	if len(_protonNextToken) > 0 {
		input.NextToken = aws.String(_protonNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRepositorySyncDefinitions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*proton.ListRepositorySyncDefinitionsOutput
	p := proton.NewListRepositorySyncDefinitionsPaginator(client, input)
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

// Get a list service of instance Infrastructure as Code (IaC) outputs.
// Deprecated: AWS Proton is not accepting new customers.
func proton_ListServiceInstanceOutputs(cfg aws.Config, client *proton.Client) {
	input := &proton.ListServiceInstanceOutputsInput{
		// ServiceInstanceName: *string, // Required
		// ServiceName: *string, // Required
	}

	if len(_protonServiceInstanceName) > 0 {
		input.ServiceInstanceName = aws.String(_protonServiceInstanceName)
	}
	if len(_protonServiceName) > 0 {
		input.ServiceName = aws.String(_protonServiceName)
	}
	if len(_protonDeploymentId) > 0 {
		input.DeploymentId = aws.String(_protonDeploymentId)
	}
	if len(_protonNextToken) > 0 {
		input.NextToken = aws.String(_protonNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListServiceInstanceOutputs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*proton.ListServiceInstanceOutputsOutput
	p := proton.NewListServiceInstanceOutputsPaginator(client, input)
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

// List provisioned resources for a service instance with details.
// Deprecated: AWS Proton is not accepting new customers.
func proton_ListServiceInstanceProvisionedResources(cfg aws.Config, client *proton.Client) {
	input := &proton.ListServiceInstanceProvisionedResourcesInput{
		// ServiceInstanceName: *string, // Required
		// ServiceName: *string, // Required
	}

	if len(_protonServiceInstanceName) > 0 {
		input.ServiceInstanceName = aws.String(_protonServiceInstanceName)
	}
	if len(_protonServiceName) > 0 {
		input.ServiceName = aws.String(_protonServiceName)
	}
	if len(_protonNextToken) > 0 {
		input.NextToken = aws.String(_protonNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListServiceInstanceProvisionedResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*proton.ListServiceInstanceProvisionedResourcesOutput
	p := proton.NewListServiceInstanceProvisionedResourcesPaginator(client, input)
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

// List service instances with summary data. This action lists service instances
// of all services in the Amazon Web Services account.
//
// Deprecated: AWS Proton is not accepting new customers.
func proton_ListServiceInstances(cfg aws.Config, client *proton.Client) {
	input := &proton.ListServiceInstancesInput{}

	if len(_protonFilters) > 0 {
		if err := assignInputField(input, "Filters", _protonFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_protonMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _protonMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_protonNextToken) > 0 {
		input.NextToken = aws.String(_protonNextToken)
	}
	if len(_protonServiceName) > 0 {
		input.ServiceName = aws.String(_protonServiceName)
	}
	if len(_protonSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _protonSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_protonSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _protonSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListServiceInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*proton.ListServiceInstancesOutput
	p := proton.NewListServiceInstancesPaginator(client, input)
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

// Get a list of service pipeline Infrastructure as Code (IaC) outputs.
// Deprecated: AWS Proton is not accepting new customers.
func proton_ListServicePipelineOutputs(cfg aws.Config, client *proton.Client) {
	input := &proton.ListServicePipelineOutputsInput{
		// ServiceName: *string, // Required
	}

	if len(_protonServiceName) > 0 {
		input.ServiceName = aws.String(_protonServiceName)
	}
	if len(_protonDeploymentId) > 0 {
		input.DeploymentId = aws.String(_protonDeploymentId)
	}
	if len(_protonNextToken) > 0 {
		input.NextToken = aws.String(_protonNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListServicePipelineOutputs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*proton.ListServicePipelineOutputsOutput
	p := proton.NewListServicePipelineOutputsPaginator(client, input)
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

// List provisioned resources for a service and pipeline with details.
// Deprecated: AWS Proton is not accepting new customers.
func proton_ListServicePipelineProvisionedResources(cfg aws.Config, client *proton.Client) {
	input := &proton.ListServicePipelineProvisionedResourcesInput{
		// ServiceName: *string, // Required
	}

	if len(_protonServiceName) > 0 {
		input.ServiceName = aws.String(_protonServiceName)
	}
	if len(_protonNextToken) > 0 {
		input.NextToken = aws.String(_protonNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListServicePipelineProvisionedResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*proton.ListServicePipelineProvisionedResourcesOutput
	p := proton.NewListServicePipelineProvisionedResourcesPaginator(client, input)
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

// List major or minor versions of a service template with detail data.
// Deprecated: AWS Proton is not accepting new customers.
func proton_ListServiceTemplateVersions(cfg aws.Config, client *proton.Client) {
	input := &proton.ListServiceTemplateVersionsInput{
		// TemplateName: *string, // Required
	}

	if len(_protonTemplateName) > 0 {
		input.TemplateName = aws.String(_protonTemplateName)
	}
	if len(_protonMajorVersion) > 0 {
		input.MajorVersion = aws.String(_protonMajorVersion)
	}
	if len(_protonMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _protonMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_protonNextToken) > 0 {
		input.NextToken = aws.String(_protonNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListServiceTemplateVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*proton.ListServiceTemplateVersionsOutput
	p := proton.NewListServiceTemplateVersionsPaginator(client, input)
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

// List service templates with detail data.
// Deprecated: AWS Proton is not accepting new customers.
func proton_ListServiceTemplates(cfg aws.Config, client *proton.Client) {
	input := &proton.ListServiceTemplatesInput{}

	if len(_protonMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _protonMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_protonNextToken) > 0 {
		input.NextToken = aws.String(_protonNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListServiceTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*proton.ListServiceTemplatesOutput
	p := proton.NewListServiceTemplatesPaginator(client, input)
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

// List services with summaries of detail data.
// Deprecated: AWS Proton is not accepting new customers.
func proton_ListServices(cfg aws.Config, client *proton.Client) {
	input := &proton.ListServicesInput{}

	if len(_protonMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _protonMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_protonNextToken) > 0 {
		input.NextToken = aws.String(_protonNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListServices(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*proton.ListServicesOutput
	p := proton.NewListServicesPaginator(client, input)
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

// List tags for a resource. For more information, see [Proton resources and tagging] in the Proton User Guide.
// Deprecated: AWS Proton is not accepting new customers.
//
// [Proton resources and tagging]: https://docs.aws.amazon.com/proton/latest/userguide/resources.html
func proton_ListTagsForResource(cfg aws.Config, client *proton.Client) {
	input := &proton.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_protonResourceArn) > 0 {
		input.ResourceArn = aws.String(_protonResourceArn)
	}
	if len(_protonMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _protonMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_protonNextToken) > 0 {
		input.NextToken = aws.String(_protonNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*proton.ListTagsForResourceOutput
	p := proton.NewListTagsForResourcePaginator(client, input)
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

// Notify Proton of status changes to a provisioned resource when you use
// self-managed provisioning.
//
// For more information, see [Self-managed provisioning] in the Proton User Guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Self-managed provisioning]: https://docs.aws.amazon.com/proton/latest/userguide/ag-works-prov-methods.html#ag-works-prov-methods-self
func proton_NotifyResourceDeploymentStatusChange(cfg aws.Config, client *proton.Client) {
	input := &proton.NotifyResourceDeploymentStatusChangeInput{
		// ResourceArn: *string, // Required
	}

	if len(_protonResourceArn) > 0 {
		input.ResourceArn = aws.String(_protonResourceArn)
	}
	if len(_protonDeploymentId) > 0 {
		input.DeploymentId = aws.String(_protonDeploymentId)
	}
	if len(_protonOutputs) > 0 {
		if err := assignInputField(input, "Outputs", _protonOutputs); err != nil {
			log.Errorf("invalid --outputs: %s", err.Error())
			return
		}
	}
	if len(_protonStatus) > 0 {
		if err := assignInputField(input, "Status", _protonStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_protonStatusMessage) > 0 {
		input.StatusMessage = aws.String(_protonStatusMessage)
	}

	if resp, err := client.NotifyResourceDeploymentStatusChange(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// In a management account, reject an environment account connection from another
// environment account.
//
// After you reject an environment account connection request, you can't accept or
// use the rejected environment account connection.
//
// You can’t reject an environment account connection that's connected to an
// environment.
//
// For more information, see [Environment account connections] in the Proton User guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Environment account connections]: https://docs.aws.amazon.com/proton/latest/userguide/ag-env-account-connections.html
func proton_RejectEnvironmentAccountConnection(cfg aws.Config, client *proton.Client) {
	input := &proton.RejectEnvironmentAccountConnectionInput{
		// Id: *string, // Required
	}

	if len(_protonId) > 0 {
		input.Id = aws.String(_protonId)
	}

	if resp, err := client.RejectEnvironmentAccountConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tag a resource. A tag is a key-value pair of metadata that you associate with
// an Proton resource.
//
// For more information, see [Proton resources and tagging] in the Proton User Guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Proton resources and tagging]: https://docs.aws.amazon.com/proton/latest/userguide/resources.html
func proton_TagResource(cfg aws.Config, client *proton.Client) {
	input := &proton.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_protonResourceArn) > 0 {
		input.ResourceArn = aws.String(_protonResourceArn)
	}
	if len(_protonTags) > 0 {
		if err := assignInputField(input, "Tags", _protonTags); err != nil {
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

// Remove a customer tag from a resource. A tag is a key-value pair of metadata
// associated with an Proton resource.
//
// For more information, see [Proton resources and tagging] in the Proton User Guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Proton resources and tagging]: https://docs.aws.amazon.com/proton/latest/userguide/resources.html
func proton_UntagResource(cfg aws.Config, client *proton.Client) {
	input := &proton.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_protonResourceArn) > 0 {
		input.ResourceArn = aws.String(_protonResourceArn)
	}
	if len(_protonTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _protonTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update Proton settings that are used for multiple services in the Amazon Web
// Services account.
//
// Deprecated: AWS Proton is not accepting new customers.
func proton_UpdateAccountSettings(cfg aws.Config, client *proton.Client) {
	input := &proton.UpdateAccountSettingsInput{}

	if len(_protonDeletePipelineProvisioningRepository) > 0 {
		if err := assignInputField(input, "DeletePipelineProvisioningRepository", _protonDeletePipelineProvisioningRepository); err != nil {
			log.Errorf("invalid --delete-pipeline-provisioning-repository: %s", err.Error())
			return
		}
	}
	if len(_protonPipelineCodebuildRoleArn) > 0 {
		input.PipelineCodebuildRoleArn = aws.String(_protonPipelineCodebuildRoleArn)
	}
	if len(_protonPipelineProvisioningRepository) > 0 {
		if err := assignInputField(input, "PipelineProvisioningRepository", _protonPipelineProvisioningRepository); err != nil {
			log.Errorf("invalid --pipeline-provisioning-repository: %s", err.Error())
			return
		}
	}
	if len(_protonPipelineServiceRoleArn) > 0 {
		input.PipelineServiceRoleArn = aws.String(_protonPipelineServiceRoleArn)
	}

	if resp, err := client.UpdateAccountSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a component.
// There are a few modes for updating a component. The deploymentType field
// defines the mode.
//
// You can't update a component while its deployment status, or the deployment
// status of a service instance attached to it, is IN_PROGRESS .
//
// For more information about components, see [Proton components] in the Proton User Guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Proton components]: https://docs.aws.amazon.com/proton/latest/userguide/ag-components.html
func proton_UpdateComponent(cfg aws.Config, client *proton.Client) {
	input := &proton.UpdateComponentInput{
		// DeploymentType: types.ComponentDeploymentUpdateType, // Required
		// Name: *string, // Required
	}

	if len(_protonDeploymentType) > 0 {
		if err := assignInputField(input, "DeploymentType", _protonDeploymentType); err != nil {
			log.Errorf("invalid --deployment-type: %s", err.Error())
			return
		}
	}
	if len(_protonName) > 0 {
		input.Name = aws.String(_protonName)
	}
	if len(_protonClientToken) > 0 {
		input.ClientToken = aws.String(_protonClientToken)
	}
	if len(_protonDescription) > 0 {
		input.Description = aws.String(_protonDescription)
	}
	if len(_protonServiceInstanceName) > 0 {
		input.ServiceInstanceName = aws.String(_protonServiceInstanceName)
	}
	if len(_protonServiceName) > 0 {
		input.ServiceName = aws.String(_protonServiceName)
	}
	if len(_protonServiceSpec) > 0 {
		input.ServiceSpec = aws.String(_protonServiceSpec)
	}
	if len(_protonTemplateFile) > 0 {
		input.TemplateFile = aws.String(_protonTemplateFile)
	}

	if resp, err := client.UpdateComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update an environment.
// If the environment is associated with an environment account connection, don't
// update or include the protonServiceRoleArn and provisioningRepository parameter
// to update or connect to an environment account connection.
//
// You can only update to a new environment account connection if that connection
// was created in the same environment account that the current environment account
// connection was created in. The account connection must also be associated with
// the current environment.
//
// If the environment isn't associated with an environment account connection,
// don't update or include the environmentAccountConnectionId parameter. You can't
// update or connect the environment to an environment account connection if it
// isn't already associated with an environment connection.
//
// You can update either the environmentAccountConnectionId or protonServiceRoleArn
// parameter and value. You can’t update both.
//
// If the environment was configured for Amazon Web Services-managed provisioning,
// omit the provisioningRepository parameter.
//
// If the environment was configured for self-managed provisioning, specify the
// provisioningRepository parameter and omit the protonServiceRoleArn and
// environmentAccountConnectionId parameters.
//
// For more information, see [Environments] and [Provisioning methods] in the Proton User Guide.
//
// There are four modes for updating an environment. The deploymentType field
// defines the mode.
//
// # NONE
//
// In this mode, a deployment doesn't occur. Only the requested metadata
// parameters are updated.
//
// CURRENT_VERSION
//
// In this mode, the environment is deployed and updated with the new spec that
// you provide. Only requested parameters are updated. Don’t include minor or major
// version parameters when you use this deployment-type .
//
// MINOR_VERSION
//
// In this mode, the environment is deployed and updated with the published,
// recommended (latest) minor version of the current major version in use, by
// default. You can also specify a different minor version of the current major
// version in use.
//
// MAJOR_VERSION
//
// In this mode, the environment is deployed and updated with the published,
// recommended (latest) major and minor version of the current template, by
// default. You can also specify a different major version that's higher than the
// major version in use and a minor version.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Environments]: https://docs.aws.amazon.com/proton/latest/userguide/ag-environments.html
// [Provisioning methods]: https://docs.aws.amazon.com/proton/latest/userguide/ag-works-prov-methods.html
func proton_UpdateEnvironment(cfg aws.Config, client *proton.Client) {
	input := &proton.UpdateEnvironmentInput{
		// DeploymentType: types.DeploymentUpdateType, // Required
		// Name: *string, // Required
	}

	if len(_protonDeploymentType) > 0 {
		if err := assignInputField(input, "DeploymentType", _protonDeploymentType); err != nil {
			log.Errorf("invalid --deployment-type: %s", err.Error())
			return
		}
	}
	if len(_protonName) > 0 {
		input.Name = aws.String(_protonName)
	}
	if len(_protonCodebuildRoleArn) > 0 {
		input.CodebuildRoleArn = aws.String(_protonCodebuildRoleArn)
	}
	if len(_protonComponentRoleArn) > 0 {
		input.ComponentRoleArn = aws.String(_protonComponentRoleArn)
	}
	if len(_protonDescription) > 0 {
		input.Description = aws.String(_protonDescription)
	}
	if len(_protonEnvironmentAccountConnectionId) > 0 {
		input.EnvironmentAccountConnectionId = aws.String(_protonEnvironmentAccountConnectionId)
	}
	if len(_protonProtonServiceRoleArn) > 0 {
		input.ProtonServiceRoleArn = aws.String(_protonProtonServiceRoleArn)
	}
	if len(_protonProvisioningRepository) > 0 {
		if err := assignInputField(input, "ProvisioningRepository", _protonProvisioningRepository); err != nil {
			log.Errorf("invalid --provisioning-repository: %s", err.Error())
			return
		}
	}
	if len(_protonSpec) > 0 {
		input.Spec = aws.String(_protonSpec)
	}
	if len(_protonTemplateMajorVersion) > 0 {
		input.TemplateMajorVersion = aws.String(_protonTemplateMajorVersion)
	}
	if len(_protonTemplateMinorVersion) > 0 {
		input.TemplateMinorVersion = aws.String(_protonTemplateMinorVersion)
	}

	if resp, err := client.UpdateEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// In an environment account, update an environment account connection to use a
// new IAM role.
//
// For more information, see [Environment account connections] in the Proton User guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Environment account connections]: https://docs.aws.amazon.com/proton/latest/userguide/ag-env-account-connections.html
func proton_UpdateEnvironmentAccountConnection(cfg aws.Config, client *proton.Client) {
	input := &proton.UpdateEnvironmentAccountConnectionInput{
		// Id: *string, // Required
	}

	if len(_protonId) > 0 {
		input.Id = aws.String(_protonId)
	}
	if len(_protonCodebuildRoleArn) > 0 {
		input.CodebuildRoleArn = aws.String(_protonCodebuildRoleArn)
	}
	if len(_protonComponentRoleArn) > 0 {
		input.ComponentRoleArn = aws.String(_protonComponentRoleArn)
	}
	if len(_protonRoleArn) > 0 {
		input.RoleArn = aws.String(_protonRoleArn)
	}

	if resp, err := client.UpdateEnvironmentAccountConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update an environment template.
// Deprecated: AWS Proton is not accepting new customers.
func proton_UpdateEnvironmentTemplate(cfg aws.Config, client *proton.Client) {
	input := &proton.UpdateEnvironmentTemplateInput{
		// Name: *string, // Required
	}

	if len(_protonName) > 0 {
		input.Name = aws.String(_protonName)
	}
	if len(_protonDescription) > 0 {
		input.Description = aws.String(_protonDescription)
	}
	if len(_protonDisplayName) > 0 {
		input.DisplayName = aws.String(_protonDisplayName)
	}

	if resp, err := client.UpdateEnvironmentTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a major or minor version of an environment template.
// Deprecated: AWS Proton is not accepting new customers.
func proton_UpdateEnvironmentTemplateVersion(cfg aws.Config, client *proton.Client) {
	input := &proton.UpdateEnvironmentTemplateVersionInput{
		// MajorVersion: *string, // Required
		// MinorVersion: *string, // Required
		// TemplateName: *string, // Required
	}

	if len(_protonMajorVersion) > 0 {
		input.MajorVersion = aws.String(_protonMajorVersion)
	}
	if len(_protonMinorVersion) > 0 {
		input.MinorVersion = aws.String(_protonMinorVersion)
	}
	if len(_protonTemplateName) > 0 {
		input.TemplateName = aws.String(_protonTemplateName)
	}
	if len(_protonDescription) > 0 {
		input.Description = aws.String(_protonDescription)
	}
	if len(_protonStatus) > 0 {
		if err := assignInputField(input, "Status", _protonStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEnvironmentTemplateVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Edit a service description or use a spec to add and delete service instances.
// Existing service instances and the service pipeline can't be edited using this
// API. They can only be deleted.
//
// Use the description parameter to modify the description.
//
// Edit the spec parameter to add or delete instances.
//
// You can't delete a service instance (remove it from the spec) if it has an
// attached component.
//
// For more information about components, see [Proton components] in the Proton User Guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Proton components]: https://docs.aws.amazon.com/proton/latest/userguide/ag-components.html
func proton_UpdateService(cfg aws.Config, client *proton.Client) {
	input := &proton.UpdateServiceInput{
		// Name: *string, // Required
	}

	if len(_protonName) > 0 {
		input.Name = aws.String(_protonName)
	}
	if len(_protonDescription) > 0 {
		input.Description = aws.String(_protonDescription)
	}
	if len(_protonSpec) > 0 {
		input.Spec = aws.String(_protonSpec)
	}

	if resp, err := client.UpdateService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a service instance.
// There are a few modes for updating a service instance. The deploymentType field
// defines the mode.
//
// You can't update a service instance while its deployment status, or the
// deployment status of a component attached to it, is IN_PROGRESS .
//
// For more information about components, see [Proton components] in the Proton User Guide.
//
// Deprecated: AWS Proton is not accepting new customers.
//
// [Proton components]: https://docs.aws.amazon.com/proton/latest/userguide/ag-components.html
func proton_UpdateServiceInstance(cfg aws.Config, client *proton.Client) {
	input := &proton.UpdateServiceInstanceInput{
		// DeploymentType: types.DeploymentUpdateType, // Required
		// Name: *string, // Required
		// ServiceName: *string, // Required
	}

	if len(_protonDeploymentType) > 0 {
		if err := assignInputField(input, "DeploymentType", _protonDeploymentType); err != nil {
			log.Errorf("invalid --deployment-type: %s", err.Error())
			return
		}
	}
	if len(_protonName) > 0 {
		input.Name = aws.String(_protonName)
	}
	if len(_protonServiceName) > 0 {
		input.ServiceName = aws.String(_protonServiceName)
	}
	if len(_protonClientToken) > 0 {
		input.ClientToken = aws.String(_protonClientToken)
	}
	if len(_protonSpec) > 0 {
		input.Spec = aws.String(_protonSpec)
	}
	if len(_protonTemplateMajorVersion) > 0 {
		input.TemplateMajorVersion = aws.String(_protonTemplateMajorVersion)
	}
	if len(_protonTemplateMinorVersion) > 0 {
		input.TemplateMinorVersion = aws.String(_protonTemplateMinorVersion)
	}

	if resp, err := client.UpdateServiceInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the service pipeline.
// There are four modes for updating a service pipeline. The deploymentType field
// defines the mode.
//
// # NONE
//
// In this mode, a deployment doesn't occur. Only the requested metadata
// parameters are updated.
//
// CURRENT_VERSION
//
// In this mode, the service pipeline is deployed and updated with the new spec
// that you provide. Only requested parameters are updated. Don’t include major or
// minor version parameters when you use this deployment-type .
//
// MINOR_VERSION
//
// In this mode, the service pipeline is deployed and updated with the published,
// recommended (latest) minor version of the current major version in use, by
// default. You can specify a different minor version of the current major version
// in use.
//
// MAJOR_VERSION
//
// In this mode, the service pipeline is deployed and updated with the published,
// recommended (latest) major and minor version of the current template by default.
// You can specify a different major version that's higher than the major version
// in use and a minor version.
//
// Deprecated: AWS Proton is not accepting new customers.
func proton_UpdateServicePipeline(cfg aws.Config, client *proton.Client) {
	input := &proton.UpdateServicePipelineInput{
		// DeploymentType: types.DeploymentUpdateType, // Required
		// ServiceName: *string, // Required
		// Spec: *string, // Required
	}

	if len(_protonDeploymentType) > 0 {
		if err := assignInputField(input, "DeploymentType", _protonDeploymentType); err != nil {
			log.Errorf("invalid --deployment-type: %s", err.Error())
			return
		}
	}
	if len(_protonServiceName) > 0 {
		input.ServiceName = aws.String(_protonServiceName)
	}
	if len(_protonSpec) > 0 {
		input.Spec = aws.String(_protonSpec)
	}
	if len(_protonTemplateMajorVersion) > 0 {
		input.TemplateMajorVersion = aws.String(_protonTemplateMajorVersion)
	}
	if len(_protonTemplateMinorVersion) > 0 {
		input.TemplateMinorVersion = aws.String(_protonTemplateMinorVersion)
	}

	if resp, err := client.UpdateServicePipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the service sync blocker by resolving it.
// Deprecated: AWS Proton is not accepting new customers.
func proton_UpdateServiceSyncBlocker(cfg aws.Config, client *proton.Client) {
	input := &proton.UpdateServiceSyncBlockerInput{
		// Id: *string, // Required
		// ResolvedReason: *string, // Required
	}

	if len(_protonId) > 0 {
		input.Id = aws.String(_protonId)
	}
	if len(_protonResolvedReason) > 0 {
		input.ResolvedReason = aws.String(_protonResolvedReason)
	}

	if resp, err := client.UpdateServiceSyncBlocker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the Proton Ops config file.
// Deprecated: AWS Proton is not accepting new customers.
func proton_UpdateServiceSyncConfig(cfg aws.Config, client *proton.Client) {
	input := &proton.UpdateServiceSyncConfigInput{
		// Branch: *string, // Required
		// FilePath: *string, // Required
		// RepositoryName: *string, // Required
		// RepositoryProvider: types.RepositoryProvider, // Required
		// ServiceName: *string, // Required
	}

	if len(_protonBranch) > 0 {
		input.Branch = aws.String(_protonBranch)
	}
	if len(_protonFilePath) > 0 {
		input.FilePath = aws.String(_protonFilePath)
	}
	if len(_protonRepositoryName) > 0 {
		input.RepositoryName = aws.String(_protonRepositoryName)
	}
	if len(_protonRepositoryProvider) > 0 {
		if err := assignInputField(input, "RepositoryProvider", _protonRepositoryProvider); err != nil {
			log.Errorf("invalid --repository-provider: %s", err.Error())
			return
		}
	}
	if len(_protonServiceName) > 0 {
		input.ServiceName = aws.String(_protonServiceName)
	}

	if resp, err := client.UpdateServiceSyncConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a service template.
// Deprecated: AWS Proton is not accepting new customers.
func proton_UpdateServiceTemplate(cfg aws.Config, client *proton.Client) {
	input := &proton.UpdateServiceTemplateInput{
		// Name: *string, // Required
	}

	if len(_protonName) > 0 {
		input.Name = aws.String(_protonName)
	}
	if len(_protonDescription) > 0 {
		input.Description = aws.String(_protonDescription)
	}
	if len(_protonDisplayName) > 0 {
		input.DisplayName = aws.String(_protonDisplayName)
	}

	if resp, err := client.UpdateServiceTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a major or minor version of a service template.
// Deprecated: AWS Proton is not accepting new customers.
func proton_UpdateServiceTemplateVersion(cfg aws.Config, client *proton.Client) {
	input := &proton.UpdateServiceTemplateVersionInput{
		// MajorVersion: *string, // Required
		// MinorVersion: *string, // Required
		// TemplateName: *string, // Required
	}

	if len(_protonMajorVersion) > 0 {
		input.MajorVersion = aws.String(_protonMajorVersion)
	}
	if len(_protonMinorVersion) > 0 {
		input.MinorVersion = aws.String(_protonMinorVersion)
	}
	if len(_protonTemplateName) > 0 {
		input.TemplateName = aws.String(_protonTemplateName)
	}
	if len(_protonCompatibleEnvironmentTemplates) > 0 {
		if err := assignInputField(input, "CompatibleEnvironmentTemplates", _protonCompatibleEnvironmentTemplates); err != nil {
			log.Errorf("invalid --compatible-environment-templates: %s", err.Error())
			return
		}
	}
	if len(_protonDescription) > 0 {
		input.Description = aws.String(_protonDescription)
	}
	if len(_protonStatus) > 0 {
		if err := assignInputField(input, "Status", _protonStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_protonSupportedComponentSources) > 0 {
		if err := assignInputField(input, "SupportedComponentSources", _protonSupportedComponentSources); err != nil {
			log.Errorf("invalid --supported-component-sources: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateServiceTemplateVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update template sync configuration parameters, except for the templateName and
// templateType . Repository details (branch, name, and provider) should be of a
// linked repository. A linked repository is a repository that has been registered
// with Proton. For more information, see CreateRepository.
//
// Deprecated: AWS Proton is not accepting new customers.
func proton_UpdateTemplateSyncConfig(cfg aws.Config, client *proton.Client) {
	input := &proton.UpdateTemplateSyncConfigInput{
		// Branch: *string, // Required
		// RepositoryName: *string, // Required
		// RepositoryProvider: types.RepositoryProvider, // Required
		// TemplateName: *string, // Required
		// TemplateType: types.TemplateType, // Required
	}

	if len(_protonBranch) > 0 {
		input.Branch = aws.String(_protonBranch)
	}
	if len(_protonRepositoryName) > 0 {
		input.RepositoryName = aws.String(_protonRepositoryName)
	}
	if len(_protonRepositoryProvider) > 0 {
		if err := assignInputField(input, "RepositoryProvider", _protonRepositoryProvider); err != nil {
			log.Errorf("invalid --repository-provider: %s", err.Error())
			return
		}
	}
	if len(_protonTemplateName) > 0 {
		input.TemplateName = aws.String(_protonTemplateName)
	}
	if len(_protonTemplateType) > 0 {
		if err := assignInputField(input, "TemplateType", _protonTemplateType); err != nil {
			log.Errorf("invalid --template-type: %s", err.Error())
			return
		}
	}
	if len(_protonSubdirectory) > 0 {
		input.Subdirectory = aws.String(_protonSubdirectory)
	}

	if resp, err := client.UpdateTemplateSyncConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_protonCmd)
	_protonCmd.Flags().SortFlags = false

	_protonCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_protonCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_protonCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_protonCmd.Flags().StringVarP(&_protonBranch, "branch", "", "", "Branch")
	_protonCmd.Flags().StringVarP(&_protonBranchName, "branch-name", "", "", "Branch Name")
	_protonCmd.Flags().StringVarP(&_protonClientToken, "client-token", "", "", "Client Token")
	_protonCmd.Flags().StringVarP(&_protonCodebuildRoleArn, "codebuild-role-arn", "", "", "Codebuild Role ARN")
	_protonCmd.Flags().StringVarP(&_protonCompatibleEnvironmentTemplates, "compatible-environment-templates", "", "", "Compatible Environment Templates")
	_protonCmd.Flags().StringVarP(&_protonComponentName, "component-name", "", "", "Component Name")
	_protonCmd.Flags().StringVarP(&_protonComponentRoleArn, "component-role-arn", "", "", "Component Role ARN")
	_protonCmd.Flags().StringVarP(&_protonConnectionArn, "connection-arn", "", "", "Connection ARN")
	_protonCmd.Flags().StringVarP(&_protonDeletePipelineProvisioningRepository, "delete-pipeline-provisioning-repository", "", "", "Delete Pipeline Provisioning Repository")
	_protonCmd.Flags().StringVarP(&_protonDeploymentId, "deployment-id", "", "", "Deployment ID")
	_protonCmd.Flags().StringVarP(&_protonDeploymentType, "deployment-type", "", "", "Deployment Type")
	_protonCmd.Flags().StringVarP(&_protonDescription, "description", "", "", "Description")
	_protonCmd.Flags().StringVarP(&_protonDisplayName, "display-name", "", "", "Display Name")
	_protonCmd.Flags().StringVarP(&_protonEncryptionKey, "encryption-key", "", "", "Encryption Key")
	_protonCmd.Flags().StringVarP(&_protonEnvironmentAccountConnectionId, "environment-account-connection-id", "", "", "Environment Account Connection ID")
	_protonCmd.Flags().StringVarP(&_protonEnvironmentName, "environment-name", "", "", "Environment Name")
	_protonCmd.Flags().StringVarP(&_protonEnvironmentTemplates, "environment-templates", "", "", "Environment Templates")
	_protonCmd.Flags().StringVarP(&_protonFilePath, "file-path", "", "", "File Path")
	_protonCmd.Flags().StringVarP(&_protonFilters, "filters", "", "", "Filters")
	_protonCmd.Flags().StringVarP(&_protonId, "id", "", "", "ID")
	_protonCmd.Flags().StringVarP(&_protonMajorVersion, "major-version", "", "", "Major Version")
	_protonCmd.Flags().StringVarP(&_protonManagementAccountId, "management-account-id", "", "", "Management Account ID")
	_protonCmd.Flags().StringVarP(&_protonManifest, "manifest", "", "", "Manifest")
	_protonCmd.Flags().StringVarP(&_protonMaxResults, "max-results", "", "", "Max Results")
	_protonCmd.Flags().StringVarP(&_protonMinorVersion, "minor-version", "", "", "Minor Version")
	_protonCmd.Flags().StringVarP(&_protonName, "name", "", "", "Name")
	_protonCmd.Flags().StringVarP(&_protonNextToken, "next-token", "", "", "Next Token")
	_protonCmd.Flags().StringVarP(&_protonOutputs, "outputs", "", "", "Outputs")
	_protonCmd.Flags().StringVarP(&_protonPipelineCodebuildRoleArn, "pipeline-codebuild-role-arn", "", "", "Pipeline Codebuild Role ARN")
	_protonCmd.Flags().StringVarP(&_protonPipelineProvisioning, "pipeline-provisioning", "", "", "Pipeline Provisioning")
	_protonCmd.Flags().StringVarP(&_protonPipelineProvisioningRepository, "pipeline-provisioning-repository", "", "", "Pipeline Provisioning Repository")
	_protonCmd.Flags().StringVarP(&_protonPipelineServiceRoleArn, "pipeline-service-role-arn", "", "", "Pipeline Service Role ARN")
	_protonCmd.Flags().StringVarP(&_protonProtonServiceRoleArn, "proton-service-role-arn", "", "", "Proton Service Role ARN")
	_protonCmd.Flags().StringVarP(&_protonProvider, "provider", "", "", "Provider")
	_protonCmd.Flags().StringVarP(&_protonProvisioning, "provisioning", "", "", "Provisioning")
	_protonCmd.Flags().StringVarP(&_protonProvisioningRepository, "provisioning-repository", "", "", "Provisioning Repository")
	_protonCmd.Flags().StringVarP(&_protonRepositoryConnectionArn, "repository-connection-arn", "", "", "Repository Connection ARN")
	_protonCmd.Flags().StringVarP(&_protonRepositoryId, "repository-id", "", "", "Repository ID")
	_protonCmd.Flags().StringVarP(&_protonRepositoryName, "repository-name", "", "", "Repository Name")
	_protonCmd.Flags().StringVarP(&_protonRepositoryProvider, "repository-provider", "", "", "Repository Provider")
	_protonCmd.Flags().StringVarP(&_protonRequestedBy, "requested-by", "", "", "Requested By")
	_protonCmd.Flags().StringVarP(&_protonResolvedReason, "resolved-reason", "", "", "Resolved Reason")
	_protonCmd.Flags().StringVarP(&_protonResourceArn, "resource-arn", "", "", "Resource ARN")
	_protonCmd.Flags().StringVarP(&_protonRoleArn, "role-arn", "", "", "Role ARN")
	_protonCmd.Flags().StringVarP(&_protonServiceInstanceName, "service-instance-name", "", "", "Service Instance Name")
	_protonCmd.Flags().StringVarP(&_protonServiceName, "service-name", "", "", "Service Name")
	_protonCmd.Flags().StringVarP(&_protonServiceSpec, "service-spec", "", "", "Service Spec")
	_protonCmd.Flags().StringVarP(&_protonSortBy, "sort-by", "", "", "Sort By")
	_protonCmd.Flags().StringVarP(&_protonSortOrder, "sort-order", "", "", "Sort Order")
	_protonCmd.Flags().StringVarP(&_protonSource, "source", "", "", "Source")
	_protonCmd.Flags().StringVarP(&_protonSpec, "spec", "", "", "Spec")
	_protonCmd.Flags().StringVarP(&_protonStatus, "status", "", "", "Status")
	_protonCmd.Flags().StringVarP(&_protonStatusMessage, "status-message", "", "", "Status Message")
	_protonCmd.Flags().StringVarP(&_protonStatuses, "statuses", "", "", "Statuses")
	_protonCmd.Flags().StringVarP(&_protonSubdirectory, "subdirectory", "", "", "Subdirectory")
	_protonCmd.Flags().StringVarP(&_protonSupportedComponentSources, "supported-component-sources", "", "", "Supported Component Sources")
	_protonCmd.Flags().StringVarP(&_protonSyncType, "sync-type", "", "", "Sync Type")
	_protonCmd.Flags().StringSliceVarP(&_protonTagKeys, "tag-keys", "", nil, "Tag Keys")
	_protonCmd.Flags().StringVarP(&_protonTags, "tags", "", "", "Tags")
	_protonCmd.Flags().StringVarP(&_protonTemplateFile, "template-file", "", "", "Template File")
	_protonCmd.Flags().StringVarP(&_protonTemplateMajorVersion, "template-major-version", "", "", "Template Major Version")
	_protonCmd.Flags().StringVarP(&_protonTemplateMinorVersion, "template-minor-version", "", "", "Template Minor Version")
	_protonCmd.Flags().StringVarP(&_protonTemplateName, "template-name", "", "", "Template Name")
	_protonCmd.Flags().StringVarP(&_protonTemplateType, "template-type", "", "", "Template Type")
	_protonCmd.Flags().StringVarP(&_protonTemplateVersion, "template-version", "", "", "Template Version")

	_protonCmd.Flags().BoolVarP(&_protonAcceptEnvironmentAccountConnection, "accept-environment-account-connection", "", false, "Accept Environment Account Connection")
	_protonCmd.Flags().BoolVarP(&_protonCancelComponentDeployment, "cancel-component-deployment", "", false, "Cancel Component Deployment")
	_protonCmd.Flags().BoolVarP(&_protonCancelEnvironmentDeployment, "cancel-environment-deployment", "", false, "Cancel Environment Deployment")
	_protonCmd.Flags().BoolVarP(&_protonCancelServiceInstanceDeployment, "cancel-service-instance-deployment", "", false, "Cancel Service Instance Deployment")
	_protonCmd.Flags().BoolVarP(&_protonCancelServicePipelineDeployment, "cancel-service-pipeline-deployment", "", false, "Cancel Service Pipeline Deployment")
	_protonCmd.Flags().BoolVarP(&_protonCreateComponent, "create-component", "", false, "Create Component")
	_protonCmd.Flags().BoolVarP(&_protonCreateEnvironment, "create-environment", "", false, "Create Environment")
	_protonCmd.Flags().BoolVarP(&_protonCreateEnvironmentAccountConnection, "create-environment-account-connection", "", false, "Create Environment Account Connection")
	_protonCmd.Flags().BoolVarP(&_protonCreateEnvironmentTemplate, "create-environment-template", "", false, "Create Environment Template")
	_protonCmd.Flags().BoolVarP(&_protonCreateEnvironmentTemplateVersion, "create-environment-template-version", "", false, "Create Environment Template Version")
	_protonCmd.Flags().BoolVarP(&_protonCreateRepository, "create-repository", "", false, "Create Repository")
	_protonCmd.Flags().BoolVarP(&_protonCreateService, "create-service", "", false, "Create Service")
	_protonCmd.Flags().BoolVarP(&_protonCreateServiceInstance, "create-service-instance", "", false, "Create Service Instance")
	_protonCmd.Flags().BoolVarP(&_protonCreateServiceSyncConfig, "create-service-sync-config", "", false, "Create Service Sync Config")
	_protonCmd.Flags().BoolVarP(&_protonCreateServiceTemplate, "create-service-template", "", false, "Create Service Template")
	_protonCmd.Flags().BoolVarP(&_protonCreateServiceTemplateVersion, "create-service-template-version", "", false, "Create Service Template Version")
	_protonCmd.Flags().BoolVarP(&_protonCreateTemplateSyncConfig, "create-template-sync-config", "", false, "Create Template Sync Config")
	_protonCmd.Flags().BoolVarP(&_protonDeleteComponent, "delete-component", "", false, "Delete Component")
	_protonCmd.Flags().BoolVarP(&_protonDeleteDeployment, "delete-deployment", "", false, "Delete Deployment")
	_protonCmd.Flags().BoolVarP(&_protonDeleteEnvironment, "delete-environment", "", false, "Delete Environment")
	_protonCmd.Flags().BoolVarP(&_protonDeleteEnvironmentAccountConnection, "delete-environment-account-connection", "", false, "Delete Environment Account Connection")
	_protonCmd.Flags().BoolVarP(&_protonDeleteEnvironmentTemplate, "delete-environment-template", "", false, "Delete Environment Template")
	_protonCmd.Flags().BoolVarP(&_protonDeleteEnvironmentTemplateVersion, "delete-environment-template-version", "", false, "Delete Environment Template Version")
	_protonCmd.Flags().BoolVarP(&_protonDeleteRepository, "delete-repository", "", false, "Delete Repository")
	_protonCmd.Flags().BoolVarP(&_protonDeleteService, "delete-service", "", false, "Delete Service")
	_protonCmd.Flags().BoolVarP(&_protonDeleteServiceSyncConfig, "delete-service-sync-config", "", false, "Delete Service Sync Config")
	_protonCmd.Flags().BoolVarP(&_protonDeleteServiceTemplate, "delete-service-template", "", false, "Delete Service Template")
	_protonCmd.Flags().BoolVarP(&_protonDeleteServiceTemplateVersion, "delete-service-template-version", "", false, "Delete Service Template Version")
	_protonCmd.Flags().BoolVarP(&_protonDeleteTemplateSyncConfig, "delete-template-sync-config", "", false, "Delete Template Sync Config")
	_protonCmd.Flags().BoolVarP(&_protonGetAccountSettings, "get-account-settings", "", false, "Get Account Settings")
	_protonCmd.Flags().BoolVarP(&_protonGetComponent, "get-component", "", false, "Get Component")
	_protonCmd.Flags().BoolVarP(&_protonGetDeployment, "get-deployment", "", false, "Get Deployment")
	_protonCmd.Flags().BoolVarP(&_protonGetEnvironment, "get-environment", "", false, "Get Environment")
	_protonCmd.Flags().BoolVarP(&_protonGetEnvironmentAccountConnection, "get-environment-account-connection", "", false, "Get Environment Account Connection")
	_protonCmd.Flags().BoolVarP(&_protonGetEnvironmentTemplate, "get-environment-template", "", false, "Get Environment Template")
	_protonCmd.Flags().BoolVarP(&_protonGetEnvironmentTemplateVersion, "get-environment-template-version", "", false, "Get Environment Template Version")
	_protonCmd.Flags().BoolVarP(&_protonGetRepository, "get-repository", "", false, "Get Repository")
	_protonCmd.Flags().BoolVarP(&_protonGetRepositorySyncStatus, "get-repository-sync-status", "", false, "Get Repository Sync Status")
	_protonCmd.Flags().BoolVarP(&_protonGetResourcesSummary, "get-resources-summary", "", false, "Get Resources Summary")
	_protonCmd.Flags().BoolVarP(&_protonGetService, "get-service", "", false, "Get Service")
	_protonCmd.Flags().BoolVarP(&_protonGetServiceInstance, "get-service-instance", "", false, "Get Service Instance")
	_protonCmd.Flags().BoolVarP(&_protonGetServiceInstanceSyncStatus, "get-service-instance-sync-status", "", false, "Get Service Instance Sync Status")
	_protonCmd.Flags().BoolVarP(&_protonGetServiceSyncBlockerSummary, "get-service-sync-blocker-summary", "", false, "Get Service Sync Blocker Summary")
	_protonCmd.Flags().BoolVarP(&_protonGetServiceSyncConfig, "get-service-sync-config", "", false, "Get Service Sync Config")
	_protonCmd.Flags().BoolVarP(&_protonGetServiceTemplate, "get-service-template", "", false, "Get Service Template")
	_protonCmd.Flags().BoolVarP(&_protonGetServiceTemplateVersion, "get-service-template-version", "", false, "Get Service Template Version")
	_protonCmd.Flags().BoolVarP(&_protonGetTemplateSyncConfig, "get-template-sync-config", "", false, "Get Template Sync Config")
	_protonCmd.Flags().BoolVarP(&_protonGetTemplateSyncStatus, "get-template-sync-status", "", false, "Get Template Sync Status")
	_protonCmd.Flags().BoolVarP(&_protonListComponentOutputs, "list-component-outputs", "", false, "List Component Outputs")
	_protonCmd.Flags().BoolVarP(&_protonListComponentProvisionedResources, "list-component-provisioned-resources", "", false, "List Component Provisioned Resources")
	_protonCmd.Flags().BoolVarP(&_protonListComponents, "list-components", "", false, "List Components")
	_protonCmd.Flags().BoolVarP(&_protonListDeployments, "list-deployments", "", false, "List Deployments")
	_protonCmd.Flags().BoolVarP(&_protonListEnvironmentAccountConnections, "list-environment-account-connections", "", false, "List Environment Account Connections")
	_protonCmd.Flags().BoolVarP(&_protonListEnvironmentOutputs, "list-environment-outputs", "", false, "List Environment Outputs")
	_protonCmd.Flags().BoolVarP(&_protonListEnvironmentProvisionedResources, "list-environment-provisioned-resources", "", false, "List Environment Provisioned Resources")
	_protonCmd.Flags().BoolVarP(&_protonListEnvironmentTemplateVersions, "list-environment-template-versions", "", false, "List Environment Template Versions")
	_protonCmd.Flags().BoolVarP(&_protonListEnvironmentTemplates, "list-environment-templates", "", false, "List Environment Templates")
	_protonCmd.Flags().BoolVarP(&_protonListEnvironments, "list-environments", "", false, "List Environments")
	_protonCmd.Flags().BoolVarP(&_protonListRepositories, "list-repositories", "", false, "List Repositories")
	_protonCmd.Flags().BoolVarP(&_protonListRepositorySyncDefinitions, "list-repository-sync-definitions", "", false, "List Repository Sync Definitions")
	_protonCmd.Flags().BoolVarP(&_protonListServiceInstanceOutputs, "list-service-instance-outputs", "", false, "List Service Instance Outputs")
	_protonCmd.Flags().BoolVarP(&_protonListServiceInstanceProvisionedResources, "list-service-instance-provisioned-resources", "", false, "List Service Instance Provisioned Resources")
	_protonCmd.Flags().BoolVarP(&_protonListServiceInstances, "list-service-instances", "", false, "List Service Instances")
	_protonCmd.Flags().BoolVarP(&_protonListServicePipelineOutputs, "list-service-pipeline-outputs", "", false, "List Service Pipeline Outputs")
	_protonCmd.Flags().BoolVarP(&_protonListServicePipelineProvisionedResources, "list-service-pipeline-provisioned-resources", "", false, "List Service Pipeline Provisioned Resources")
	_protonCmd.Flags().BoolVarP(&_protonListServiceTemplateVersions, "list-service-template-versions", "", false, "List Service Template Versions")
	_protonCmd.Flags().BoolVarP(&_protonListServiceTemplates, "list-service-templates", "", false, "List Service Templates")
	_protonCmd.Flags().BoolVarP(&_protonListServices, "list-services", "", false, "List Services")
	_protonCmd.Flags().BoolVarP(&_protonListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_protonCmd.Flags().BoolVarP(&_protonNotifyResourceDeploymentStatusChange, "notify-resource-deployment-status-change", "", false, "Notify Resource Deployment Status Change")
	_protonCmd.Flags().BoolVarP(&_protonRejectEnvironmentAccountConnection, "reject-environment-account-connection", "", false, "Reject Environment Account Connection")
	_protonCmd.Flags().BoolVarP(&_protonTagResource, "tag-resource", "", false, "Tag Resource")
	_protonCmd.Flags().BoolVarP(&_protonUntagResource, "untag-resource", "", false, "Untag Resource")
	_protonCmd.Flags().BoolVarP(&_protonUpdateAccountSettings, "update-account-settings", "", false, "Update Account Settings")
	_protonCmd.Flags().BoolVarP(&_protonUpdateComponent, "update-component", "", false, "Update Component")
	_protonCmd.Flags().BoolVarP(&_protonUpdateEnvironment, "update-environment", "", false, "Update Environment")
	_protonCmd.Flags().BoolVarP(&_protonUpdateEnvironmentAccountConnection, "update-environment-account-connection", "", false, "Update Environment Account Connection")
	_protonCmd.Flags().BoolVarP(&_protonUpdateEnvironmentTemplate, "update-environment-template", "", false, "Update Environment Template")
	_protonCmd.Flags().BoolVarP(&_protonUpdateEnvironmentTemplateVersion, "update-environment-template-version", "", false, "Update Environment Template Version")
	_protonCmd.Flags().BoolVarP(&_protonUpdateService, "update-service", "", false, "Update Service")
	_protonCmd.Flags().BoolVarP(&_protonUpdateServiceInstance, "update-service-instance", "", false, "Update Service Instance")
	_protonCmd.Flags().BoolVarP(&_protonUpdateServicePipeline, "update-service-pipeline", "", false, "Update Service Pipeline")
	_protonCmd.Flags().BoolVarP(&_protonUpdateServiceSyncBlocker, "update-service-sync-blocker", "", false, "Update Service Sync Blocker")
	_protonCmd.Flags().BoolVarP(&_protonUpdateServiceSyncConfig, "update-service-sync-config", "", false, "Update Service Sync Config")
	_protonCmd.Flags().BoolVarP(&_protonUpdateServiceTemplate, "update-service-template", "", false, "Update Service Template")
	_protonCmd.Flags().BoolVarP(&_protonUpdateServiceTemplateVersion, "update-service-template-version", "", false, "Update Service Template Version")
	_protonCmd.Flags().BoolVarP(&_protonUpdateTemplateSyncConfig, "update-template-sync-config", "", false, "Update Template Sync Config")

}
