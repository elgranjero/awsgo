package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// imagebuilderCmd represents the imagebuilder command
var _imagebuilderCmd = &cobra.Command{
	Use:   "imagebuilder",
	Short: "AWS imagebuilder CLI",
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
		client := imagebuilder.NewFromConfig(cfg)
		if _imagebuilderCancelImageCreation {
			imagebuilder_CancelImageCreation(cfg, client)
			return
		}
		if _imagebuilderCancelLifecycleExecution {
			imagebuilder_CancelLifecycleExecution(cfg, client)
			return
		}
		if _imagebuilderCreateComponent {
			imagebuilder_CreateComponent(cfg, client)
			return
		}
		if _imagebuilderCreateContainerRecipe {
			imagebuilder_CreateContainerRecipe(cfg, client)
			return
		}
		if _imagebuilderCreateDistributionConfiguration {
			imagebuilder_CreateDistributionConfiguration(cfg, client)
			return
		}
		if _imagebuilderCreateImage {
			imagebuilder_CreateImage(cfg, client)
			return
		}
		if _imagebuilderCreateImagePipeline {
			imagebuilder_CreateImagePipeline(cfg, client)
			return
		}
		if _imagebuilderCreateImageRecipe {
			imagebuilder_CreateImageRecipe(cfg, client)
			return
		}
		if _imagebuilderCreateInfrastructureConfiguration {
			imagebuilder_CreateInfrastructureConfiguration(cfg, client)
			return
		}
		if _imagebuilderCreateLifecyclePolicy {
			imagebuilder_CreateLifecyclePolicy(cfg, client)
			return
		}
		if _imagebuilderCreateWorkflow {
			imagebuilder_CreateWorkflow(cfg, client)
			return
		}
		if _imagebuilderDeleteComponent {
			imagebuilder_DeleteComponent(cfg, client)
			return
		}
		if _imagebuilderDeleteContainerRecipe {
			imagebuilder_DeleteContainerRecipe(cfg, client)
			return
		}
		if _imagebuilderDeleteDistributionConfiguration {
			imagebuilder_DeleteDistributionConfiguration(cfg, client)
			return
		}
		if _imagebuilderDeleteImage {
			imagebuilder_DeleteImage(cfg, client)
			return
		}
		if _imagebuilderDeleteImagePipeline {
			imagebuilder_DeleteImagePipeline(cfg, client)
			return
		}
		if _imagebuilderDeleteImageRecipe {
			imagebuilder_DeleteImageRecipe(cfg, client)
			return
		}
		if _imagebuilderDeleteInfrastructureConfiguration {
			imagebuilder_DeleteInfrastructureConfiguration(cfg, client)
			return
		}
		if _imagebuilderDeleteLifecyclePolicy {
			imagebuilder_DeleteLifecyclePolicy(cfg, client)
			return
		}
		if _imagebuilderDeleteWorkflow {
			imagebuilder_DeleteWorkflow(cfg, client)
			return
		}
		if _imagebuilderDistributeImage {
			imagebuilder_DistributeImage(cfg, client)
			return
		}
		if _imagebuilderGetComponent {
			imagebuilder_GetComponent(cfg, client)
			return
		}
		if _imagebuilderGetComponentPolicy {
			imagebuilder_GetComponentPolicy(cfg, client)
			return
		}
		if _imagebuilderGetContainerRecipe {
			imagebuilder_GetContainerRecipe(cfg, client)
			return
		}
		if _imagebuilderGetContainerRecipePolicy {
			imagebuilder_GetContainerRecipePolicy(cfg, client)
			return
		}
		if _imagebuilderGetDistributionConfiguration {
			imagebuilder_GetDistributionConfiguration(cfg, client)
			return
		}
		if _imagebuilderGetImage {
			imagebuilder_GetImage(cfg, client)
			return
		}
		if _imagebuilderGetImagePipeline {
			imagebuilder_GetImagePipeline(cfg, client)
			return
		}
		if _imagebuilderGetImagePolicy {
			imagebuilder_GetImagePolicy(cfg, client)
			return
		}
		if _imagebuilderGetImageRecipe {
			imagebuilder_GetImageRecipe(cfg, client)
			return
		}
		if _imagebuilderGetImageRecipePolicy {
			imagebuilder_GetImageRecipePolicy(cfg, client)
			return
		}
		if _imagebuilderGetInfrastructureConfiguration {
			imagebuilder_GetInfrastructureConfiguration(cfg, client)
			return
		}
		if _imagebuilderGetLifecycleExecution {
			imagebuilder_GetLifecycleExecution(cfg, client)
			return
		}
		if _imagebuilderGetLifecyclePolicy {
			imagebuilder_GetLifecyclePolicy(cfg, client)
			return
		}
		if _imagebuilderGetMarketplaceResource {
			imagebuilder_GetMarketplaceResource(cfg, client)
			return
		}
		if _imagebuilderGetWorkflow {
			imagebuilder_GetWorkflow(cfg, client)
			return
		}
		if _imagebuilderGetWorkflowExecution {
			imagebuilder_GetWorkflowExecution(cfg, client)
			return
		}
		if _imagebuilderGetWorkflowStepExecution {
			imagebuilder_GetWorkflowStepExecution(cfg, client)
			return
		}
		if _imagebuilderImportComponent {
			imagebuilder_ImportComponent(cfg, client)
			return
		}
		if _imagebuilderImportDiskImage {
			imagebuilder_ImportDiskImage(cfg, client)
			return
		}
		if _imagebuilderImportVmImage {
			imagebuilder_ImportVmImage(cfg, client)
			return
		}
		if _imagebuilderListComponentBuildVersions {
			imagebuilder_ListComponentBuildVersions(cfg, client)
			return
		}
		if _imagebuilderListComponents {
			imagebuilder_ListComponents(cfg, client)
			return
		}
		if _imagebuilderListContainerRecipes {
			imagebuilder_ListContainerRecipes(cfg, client)
			return
		}
		if _imagebuilderListDistributionConfigurations {
			imagebuilder_ListDistributionConfigurations(cfg, client)
			return
		}
		if _imagebuilderListImageBuildVersions {
			imagebuilder_ListImageBuildVersions(cfg, client)
			return
		}
		if _imagebuilderListImagePackages {
			imagebuilder_ListImagePackages(cfg, client)
			return
		}
		if _imagebuilderListImagePipelineImages {
			imagebuilder_ListImagePipelineImages(cfg, client)
			return
		}
		if _imagebuilderListImagePipelines {
			imagebuilder_ListImagePipelines(cfg, client)
			return
		}
		if _imagebuilderListImageRecipes {
			imagebuilder_ListImageRecipes(cfg, client)
			return
		}
		if _imagebuilderListImageScanFindingAggregations {
			imagebuilder_ListImageScanFindingAggregations(cfg, client)
			return
		}
		if _imagebuilderListImageScanFindings {
			imagebuilder_ListImageScanFindings(cfg, client)
			return
		}
		if _imagebuilderListImages {
			imagebuilder_ListImages(cfg, client)
			return
		}
		if _imagebuilderListInfrastructureConfigurations {
			imagebuilder_ListInfrastructureConfigurations(cfg, client)
			return
		}
		if _imagebuilderListLifecycleExecutionResources {
			imagebuilder_ListLifecycleExecutionResources(cfg, client)
			return
		}
		if _imagebuilderListLifecycleExecutions {
			imagebuilder_ListLifecycleExecutions(cfg, client)
			return
		}
		if _imagebuilderListLifecyclePolicies {
			imagebuilder_ListLifecyclePolicies(cfg, client)
			return
		}
		if _imagebuilderListTagsForResource {
			imagebuilder_ListTagsForResource(cfg, client)
			return
		}
		if _imagebuilderListWaitingWorkflowSteps {
			imagebuilder_ListWaitingWorkflowSteps(cfg, client)
			return
		}
		if _imagebuilderListWorkflowBuildVersions {
			imagebuilder_ListWorkflowBuildVersions(cfg, client)
			return
		}
		if _imagebuilderListWorkflowExecutions {
			imagebuilder_ListWorkflowExecutions(cfg, client)
			return
		}
		if _imagebuilderListWorkflowStepExecutions {
			imagebuilder_ListWorkflowStepExecutions(cfg, client)
			return
		}
		if _imagebuilderListWorkflows {
			imagebuilder_ListWorkflows(cfg, client)
			return
		}
		if _imagebuilderPutComponentPolicy {
			imagebuilder_PutComponentPolicy(cfg, client)
			return
		}
		if _imagebuilderPutContainerRecipePolicy {
			imagebuilder_PutContainerRecipePolicy(cfg, client)
			return
		}
		if _imagebuilderPutImagePolicy {
			imagebuilder_PutImagePolicy(cfg, client)
			return
		}
		if _imagebuilderPutImageRecipePolicy {
			imagebuilder_PutImageRecipePolicy(cfg, client)
			return
		}
		if _imagebuilderRetryImage {
			imagebuilder_RetryImage(cfg, client)
			return
		}
		if _imagebuilderSendWorkflowStepAction {
			imagebuilder_SendWorkflowStepAction(cfg, client)
			return
		}
		if _imagebuilderStartImagePipelineExecution {
			imagebuilder_StartImagePipelineExecution(cfg, client)
			return
		}
		if _imagebuilderStartResourceStateUpdate {
			imagebuilder_StartResourceStateUpdate(cfg, client)
			return
		}
		if _imagebuilderTagResource {
			imagebuilder_TagResource(cfg, client)
			return
		}
		if _imagebuilderUntagResource {
			imagebuilder_UntagResource(cfg, client)
			return
		}
		if _imagebuilderUpdateDistributionConfiguration {
			imagebuilder_UpdateDistributionConfiguration(cfg, client)
			return
		}
		if _imagebuilderUpdateImagePipeline {
			imagebuilder_UpdateImagePipeline(cfg, client)
			return
		}
		if _imagebuilderUpdateInfrastructureConfiguration {
			imagebuilder_UpdateInfrastructureConfiguration(cfg, client)
			return
		}
		if _imagebuilderUpdateLifecyclePolicy {
			imagebuilder_UpdateLifecyclePolicy(cfg, client)
			return
		}

	},
}

var (
	_imagebuilderCancelImageCreation               bool
	_imagebuilderCancelLifecycleExecution          bool
	_imagebuilderCreateComponent                   bool
	_imagebuilderCreateContainerRecipe             bool
	_imagebuilderCreateDistributionConfiguration   bool
	_imagebuilderCreateImage                       bool
	_imagebuilderCreateImagePipeline               bool
	_imagebuilderCreateImageRecipe                 bool
	_imagebuilderCreateInfrastructureConfiguration bool
	_imagebuilderCreateLifecyclePolicy             bool
	_imagebuilderCreateWorkflow                    bool
	_imagebuilderDeleteComponent                   bool
	_imagebuilderDeleteContainerRecipe             bool
	_imagebuilderDeleteDistributionConfiguration   bool
	_imagebuilderDeleteImage                       bool
	_imagebuilderDeleteImagePipeline               bool
	_imagebuilderDeleteImageRecipe                 bool
	_imagebuilderDeleteInfrastructureConfiguration bool
	_imagebuilderDeleteLifecyclePolicy             bool
	_imagebuilderDeleteWorkflow                    bool
	_imagebuilderDistributeImage                   bool
	_imagebuilderGetComponent                      bool
	_imagebuilderGetComponentPolicy                bool
	_imagebuilderGetContainerRecipe                bool
	_imagebuilderGetContainerRecipePolicy          bool
	_imagebuilderGetDistributionConfiguration      bool
	_imagebuilderGetImage                          bool
	_imagebuilderGetImagePipeline                  bool
	_imagebuilderGetImagePolicy                    bool
	_imagebuilderGetImageRecipe                    bool
	_imagebuilderGetImageRecipePolicy              bool
	_imagebuilderGetInfrastructureConfiguration    bool
	_imagebuilderGetLifecycleExecution             bool
	_imagebuilderGetLifecyclePolicy                bool
	_imagebuilderGetMarketplaceResource            bool
	_imagebuilderGetWorkflow                       bool
	_imagebuilderGetWorkflowExecution              bool
	_imagebuilderGetWorkflowStepExecution          bool
	_imagebuilderImportComponent                   bool
	_imagebuilderImportDiskImage                   bool
	_imagebuilderImportVmImage                     bool
	_imagebuilderListComponentBuildVersions        bool
	_imagebuilderListComponents                    bool
	_imagebuilderListContainerRecipes              bool
	_imagebuilderListDistributionConfigurations    bool
	_imagebuilderListImageBuildVersions            bool
	_imagebuilderListImagePackages                 bool
	_imagebuilderListImagePipelineImages           bool
	_imagebuilderListImagePipelines                bool
	_imagebuilderListImageRecipes                  bool
	_imagebuilderListImageScanFindingAggregations  bool
	_imagebuilderListImageScanFindings             bool
	_imagebuilderListImages                        bool
	_imagebuilderListInfrastructureConfigurations  bool
	_imagebuilderListLifecycleExecutionResources   bool
	_imagebuilderListLifecycleExecutions           bool
	_imagebuilderListLifecyclePolicies             bool
	_imagebuilderListTagsForResource               bool
	_imagebuilderListWaitingWorkflowSteps          bool
	_imagebuilderListWorkflowBuildVersions         bool
	_imagebuilderListWorkflowExecutions            bool
	_imagebuilderListWorkflowStepExecutions        bool
	_imagebuilderListWorkflows                     bool
	_imagebuilderPutComponentPolicy                bool
	_imagebuilderPutContainerRecipePolicy          bool
	_imagebuilderPutImagePolicy                    bool
	_imagebuilderPutImageRecipePolicy              bool
	_imagebuilderRetryImage                        bool
	_imagebuilderSendWorkflowStepAction            bool
	_imagebuilderStartImagePipelineExecution       bool
	_imagebuilderStartResourceStateUpdate          bool
	_imagebuilderTagResource                       bool
	_imagebuilderUntagResource                     bool
	_imagebuilderUpdateDistributionConfiguration   bool
	_imagebuilderUpdateImagePipeline               bool
	_imagebuilderUpdateInfrastructureConfiguration bool
	_imagebuilderUpdateLifecyclePolicy             bool

	_imagebuilderAction                          string
	_imagebuilderAdditionalInstanceConfiguration string
	_imagebuilderAmiTags                         string
	_imagebuilderBlockDeviceMappings             string
	_imagebuilderByName                          string
	_imagebuilderChangeDescription               string
	_imagebuilderClientToken                     string
	_imagebuilderComponentArn                    string
	_imagebuilderComponentBuildVersionArn        string
	_imagebuilderComponentVersionArn             string
	_imagebuilderComponents                      string
	_imagebuilderContainerRecipeArn              string
	_imagebuilderContainerType                   string
	_imagebuilderData                            string
	_imagebuilderDescription                     string
	_imagebuilderDistributionConfigurationArn    string
	_imagebuilderDistributions                   string
	_imagebuilderDockerfileTemplateData          string
	_imagebuilderDockerfileTemplateUri           string
	_imagebuilderDryRun                          string
	_imagebuilderEnhancedImageMetadataEnabled    string
	_imagebuilderExclusionRules                  string
	_imagebuilderExecutionRole                   string
	_imagebuilderFilter                          string
	_imagebuilderFilters                         string
	_imagebuilderFormat                          string
	_imagebuilderImageArn                        string
	_imagebuilderImageBuildVersionArn            string
	_imagebuilderImageOsVersionOverride          string
	_imagebuilderImagePipelineArn                string
	_imagebuilderImageRecipeArn                  string
	_imagebuilderImageScanningConfiguration      string
	_imagebuilderImageTestsConfiguration         string
	_imagebuilderImageVersionArn                 string
	_imagebuilderIncludeDeprecated               string
	_imagebuilderIncludeResources                string
	_imagebuilderInfrastructureConfigurationArn  string
	_imagebuilderInstanceConfiguration           string
	_imagebuilderInstanceMetadataOptions         string
	_imagebuilderInstanceProfileName             string
	_imagebuilderInstanceTypes                   []string
	_imagebuilderKeyPair                         string
	_imagebuilderKmsKeyId                        string
	_imagebuilderLifecycleExecutionId            string
	_imagebuilderLifecyclePolicyArn              string
	_imagebuilderLogging                         string
	_imagebuilderLoggingConfiguration            string
	_imagebuilderMaxResults                      string
	_imagebuilderName                            string
	_imagebuilderNextToken                       string
	_imagebuilderOsVersion                       string
	_imagebuilderOwner                           string
	_imagebuilderParentImage                     string
	_imagebuilderParentResourceId                string
	_imagebuilderPlacement                       string
	_imagebuilderPlatform                        string
	_imagebuilderPlatformOverride                string
	_imagebuilderPolicy                          string
	_imagebuilderPolicyDetails                   string
	_imagebuilderReason                          string
	_imagebuilderResourceArn                     string
	_imagebuilderResourceLocation                string
	_imagebuilderResourceSelection               string
	_imagebuilderResourceTags                    string
	_imagebuilderResourceType                    string
	_imagebuilderSchedule                        string
	_imagebuilderSecurityGroupIds                []string
	_imagebuilderSemanticVersion                 string
	_imagebuilderSnsTopicArn                     string
	_imagebuilderSourceImage                     string
	_imagebuilderState                           string
	_imagebuilderStatus                          string
	_imagebuilderStepExecutionId                 string
	_imagebuilderSubnetId                        string
	_imagebuilderSupportedOsVersions             []string
	_imagebuilderTagKeys                         []string
	_imagebuilderTags                            string
	_imagebuilderTargetRepository                string
	_imagebuilderTerminateInstanceOnFailure      string
	_imagebuilderType                            string
	_imagebuilderUpdateAt                        string
	_imagebuilderUri                             string
	_imagebuilderVmImportTaskId                  string
	_imagebuilderWorkflowBuildVersionArn         string
	_imagebuilderWorkflowExecutionId             string
	_imagebuilderWorkflowVersionArn              string
	_imagebuilderWorkflows                       string
	_imagebuilderWorkingDirectory                string
)

// CancelImageCreation cancels the creation of Image. This operation can only be
// used on images in a non-terminal state.
func imagebuilder_CancelImageCreation(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.CancelImageCreationInput{
		// ClientToken: *string, // Required
		// ImageBuildVersionArn: *string, // Required
	}

	if len(_imagebuilderClientToken) > 0 {
		input.ClientToken = aws.String(_imagebuilderClientToken)
	}
	if len(_imagebuilderImageBuildVersionArn) > 0 {
		input.ImageBuildVersionArn = aws.String(_imagebuilderImageBuildVersionArn)
	}

	if resp, err := client.CancelImageCreation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancel a specific image lifecycle policy runtime instance.
func imagebuilder_CancelLifecycleExecution(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.CancelLifecycleExecutionInput{
		// ClientToken: *string, // Required
		// LifecycleExecutionId: *string, // Required
	}

	if len(_imagebuilderClientToken) > 0 {
		input.ClientToken = aws.String(_imagebuilderClientToken)
	}
	if len(_imagebuilderLifecycleExecutionId) > 0 {
		input.LifecycleExecutionId = aws.String(_imagebuilderLifecycleExecutionId)
	}

	if resp, err := client.CancelLifecycleExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new component that can be used to build, validate, test, and assess
// your image. The component is based on a YAML document that you specify using
// exactly one of the following methods:
//
// - Inline, using the data property in the request body.
//
// - A URL that points to a YAML document file stored in Amazon S3, using the uri
// property in the request body.
func imagebuilder_CreateComponent(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.CreateComponentInput{
		// ClientToken: *string, // Required
		// Name: *string, // Required
		// Platform: types.Platform, // Required
		// SemanticVersion: *string, // Required
	}

	if len(_imagebuilderClientToken) > 0 {
		input.ClientToken = aws.String(_imagebuilderClientToken)
	}
	if len(_imagebuilderName) > 0 {
		input.Name = aws.String(_imagebuilderName)
	}
	if len(_imagebuilderPlatform) > 0 {
		if err := assignInputField(input, "Platform", _imagebuilderPlatform); err != nil {
			log.Errorf("invalid --platform: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderSemanticVersion) > 0 {
		input.SemanticVersion = aws.String(_imagebuilderSemanticVersion)
	}
	if len(_imagebuilderChangeDescription) > 0 {
		input.ChangeDescription = aws.String(_imagebuilderChangeDescription)
	}
	if len(_imagebuilderData) > 0 {
		input.Data = aws.String(_imagebuilderData)
	}
	if len(_imagebuilderDescription) > 0 {
		input.Description = aws.String(_imagebuilderDescription)
	}
	if len(_imagebuilderDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _imagebuilderDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_imagebuilderKmsKeyId)
	}
	if len(_imagebuilderSupportedOsVersions) > 0 {
		input.SupportedOsVersions = append([]string(nil), _imagebuilderSupportedOsVersions...)
	}
	if len(_imagebuilderTags) > 0 {
		if err := assignInputField(input, "Tags", _imagebuilderTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderUri) > 0 {
		input.Uri = aws.String(_imagebuilderUri)
	}

	if resp, err := client.CreateComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new container recipe. Container recipes define how images are
// configured, tested, and assessed.
func imagebuilder_CreateContainerRecipe(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.CreateContainerRecipeInput{
		// ClientToken: *string, // Required
		// ContainerType: types.ContainerType, // Required
		// Name: *string, // Required
		// ParentImage: *string, // Required
		// SemanticVersion: *string, // Required
		// TargetRepository: *types.TargetContainerRepository, // Required
	}

	if len(_imagebuilderClientToken) > 0 {
		input.ClientToken = aws.String(_imagebuilderClientToken)
	}
	if len(_imagebuilderContainerType) > 0 {
		if err := assignInputField(input, "ContainerType", _imagebuilderContainerType); err != nil {
			log.Errorf("invalid --container-type: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderName) > 0 {
		input.Name = aws.String(_imagebuilderName)
	}
	if len(_imagebuilderParentImage) > 0 {
		input.ParentImage = aws.String(_imagebuilderParentImage)
	}
	if len(_imagebuilderSemanticVersion) > 0 {
		input.SemanticVersion = aws.String(_imagebuilderSemanticVersion)
	}
	if len(_imagebuilderTargetRepository) > 0 {
		if err := assignInputField(input, "TargetRepository", _imagebuilderTargetRepository); err != nil {
			log.Errorf("invalid --target-repository: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderComponents) > 0 {
		if err := assignInputField(input, "Components", _imagebuilderComponents); err != nil {
			log.Errorf("invalid --components: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderDescription) > 0 {
		input.Description = aws.String(_imagebuilderDescription)
	}
	if len(_imagebuilderDockerfileTemplateData) > 0 {
		input.DockerfileTemplateData = aws.String(_imagebuilderDockerfileTemplateData)
	}
	if len(_imagebuilderDockerfileTemplateUri) > 0 {
		input.DockerfileTemplateUri = aws.String(_imagebuilderDockerfileTemplateUri)
	}
	if len(_imagebuilderImageOsVersionOverride) > 0 {
		input.ImageOsVersionOverride = aws.String(_imagebuilderImageOsVersionOverride)
	}
	if len(_imagebuilderInstanceConfiguration) > 0 {
		if err := assignInputField(input, "InstanceConfiguration", _imagebuilderInstanceConfiguration); err != nil {
			log.Errorf("invalid --instance-configuration: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_imagebuilderKmsKeyId)
	}
	if len(_imagebuilderPlatformOverride) > 0 {
		if err := assignInputField(input, "PlatformOverride", _imagebuilderPlatformOverride); err != nil {
			log.Errorf("invalid --platform-override: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderTags) > 0 {
		if err := assignInputField(input, "Tags", _imagebuilderTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderWorkingDirectory) > 0 {
		input.WorkingDirectory = aws.String(_imagebuilderWorkingDirectory)
	}

	if resp, err := client.CreateContainerRecipe(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new distribution configuration. Distribution configurations define
// and configure the outputs of your pipeline.
func imagebuilder_CreateDistributionConfiguration(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.CreateDistributionConfigurationInput{
		// ClientToken: *string, // Required
		// Distributions: []types.Distribution, // Required
		// Name: *string, // Required
	}

	if len(_imagebuilderClientToken) > 0 {
		input.ClientToken = aws.String(_imagebuilderClientToken)
	}
	if len(_imagebuilderDistributions) > 0 {
		if err := assignInputField(input, "Distributions", _imagebuilderDistributions); err != nil {
			log.Errorf("invalid --distributions: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderName) > 0 {
		input.Name = aws.String(_imagebuilderName)
	}
	if len(_imagebuilderDescription) > 0 {
		input.Description = aws.String(_imagebuilderDescription)
	}
	if len(_imagebuilderTags) > 0 {
		if err := assignInputField(input, "Tags", _imagebuilderTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDistributionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new image. This request will create a new image along with all of the
// configured output resources defined in the distribution configuration. You must
// specify exactly one recipe for your image, using either a ContainerRecipeArn or
// an ImageRecipeArn.
func imagebuilder_CreateImage(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.CreateImageInput{
		// ClientToken: *string, // Required
		// InfrastructureConfigurationArn: *string, // Required
	}

	if len(_imagebuilderClientToken) > 0 {
		input.ClientToken = aws.String(_imagebuilderClientToken)
	}
	if len(_imagebuilderInfrastructureConfigurationArn) > 0 {
		input.InfrastructureConfigurationArn = aws.String(_imagebuilderInfrastructureConfigurationArn)
	}
	if len(_imagebuilderContainerRecipeArn) > 0 {
		input.ContainerRecipeArn = aws.String(_imagebuilderContainerRecipeArn)
	}
	if len(_imagebuilderDistributionConfigurationArn) > 0 {
		input.DistributionConfigurationArn = aws.String(_imagebuilderDistributionConfigurationArn)
	}
	if len(_imagebuilderEnhancedImageMetadataEnabled) > 0 {
		if err := assignInputField(input, "EnhancedImageMetadataEnabled", _imagebuilderEnhancedImageMetadataEnabled); err != nil {
			log.Errorf("invalid --enhanced-image-metadata-enabled: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderExecutionRole) > 0 {
		input.ExecutionRole = aws.String(_imagebuilderExecutionRole)
	}
	if len(_imagebuilderImageRecipeArn) > 0 {
		input.ImageRecipeArn = aws.String(_imagebuilderImageRecipeArn)
	}
	if len(_imagebuilderImageScanningConfiguration) > 0 {
		if err := assignInputField(input, "ImageScanningConfiguration", _imagebuilderImageScanningConfiguration); err != nil {
			log.Errorf("invalid --image-scanning-configuration: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderImageTestsConfiguration) > 0 {
		if err := assignInputField(input, "ImageTestsConfiguration", _imagebuilderImageTestsConfiguration); err != nil {
			log.Errorf("invalid --image-tests-configuration: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderLoggingConfiguration) > 0 {
		if err := assignInputField(input, "LoggingConfiguration", _imagebuilderLoggingConfiguration); err != nil {
			log.Errorf("invalid --logging-configuration: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderTags) > 0 {
		if err := assignInputField(input, "Tags", _imagebuilderTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderWorkflows) > 0 {
		if err := assignInputField(input, "Workflows", _imagebuilderWorkflows); err != nil {
			log.Errorf("invalid --workflows: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new image pipeline. Image pipelines enable you to automate the
// creation and distribution of images.
func imagebuilder_CreateImagePipeline(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.CreateImagePipelineInput{
		// ClientToken: *string, // Required
		// InfrastructureConfigurationArn: *string, // Required
		// Name: *string, // Required
	}

	if len(_imagebuilderClientToken) > 0 {
		input.ClientToken = aws.String(_imagebuilderClientToken)
	}
	if len(_imagebuilderInfrastructureConfigurationArn) > 0 {
		input.InfrastructureConfigurationArn = aws.String(_imagebuilderInfrastructureConfigurationArn)
	}
	if len(_imagebuilderName) > 0 {
		input.Name = aws.String(_imagebuilderName)
	}
	if len(_imagebuilderContainerRecipeArn) > 0 {
		input.ContainerRecipeArn = aws.String(_imagebuilderContainerRecipeArn)
	}
	if len(_imagebuilderDescription) > 0 {
		input.Description = aws.String(_imagebuilderDescription)
	}
	if len(_imagebuilderDistributionConfigurationArn) > 0 {
		input.DistributionConfigurationArn = aws.String(_imagebuilderDistributionConfigurationArn)
	}
	if len(_imagebuilderEnhancedImageMetadataEnabled) > 0 {
		if err := assignInputField(input, "EnhancedImageMetadataEnabled", _imagebuilderEnhancedImageMetadataEnabled); err != nil {
			log.Errorf("invalid --enhanced-image-metadata-enabled: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderExecutionRole) > 0 {
		input.ExecutionRole = aws.String(_imagebuilderExecutionRole)
	}
	if len(_imagebuilderImageRecipeArn) > 0 {
		input.ImageRecipeArn = aws.String(_imagebuilderImageRecipeArn)
	}
	if len(_imagebuilderImageScanningConfiguration) > 0 {
		if err := assignInputField(input, "ImageScanningConfiguration", _imagebuilderImageScanningConfiguration); err != nil {
			log.Errorf("invalid --image-scanning-configuration: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderImageTestsConfiguration) > 0 {
		if err := assignInputField(input, "ImageTestsConfiguration", _imagebuilderImageTestsConfiguration); err != nil {
			log.Errorf("invalid --image-tests-configuration: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderLoggingConfiguration) > 0 {
		if err := assignInputField(input, "LoggingConfiguration", _imagebuilderLoggingConfiguration); err != nil {
			log.Errorf("invalid --logging-configuration: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderSchedule) > 0 {
		if err := assignInputField(input, "Schedule", _imagebuilderSchedule); err != nil {
			log.Errorf("invalid --schedule: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderStatus) > 0 {
		if err := assignInputField(input, "Status", _imagebuilderStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderTags) > 0 {
		if err := assignInputField(input, "Tags", _imagebuilderTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderWorkflows) > 0 {
		if err := assignInputField(input, "Workflows", _imagebuilderWorkflows); err != nil {
			log.Errorf("invalid --workflows: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateImagePipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new image recipe. Image recipes define how images are configured,
// tested, and assessed.
func imagebuilder_CreateImageRecipe(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.CreateImageRecipeInput{
		// ClientToken: *string, // Required
		// Name: *string, // Required
		// ParentImage: *string, // Required
		// SemanticVersion: *string, // Required
	}

	if len(_imagebuilderClientToken) > 0 {
		input.ClientToken = aws.String(_imagebuilderClientToken)
	}
	if len(_imagebuilderName) > 0 {
		input.Name = aws.String(_imagebuilderName)
	}
	if len(_imagebuilderParentImage) > 0 {
		input.ParentImage = aws.String(_imagebuilderParentImage)
	}
	if len(_imagebuilderSemanticVersion) > 0 {
		input.SemanticVersion = aws.String(_imagebuilderSemanticVersion)
	}
	if len(_imagebuilderAdditionalInstanceConfiguration) > 0 {
		if err := assignInputField(input, "AdditionalInstanceConfiguration", _imagebuilderAdditionalInstanceConfiguration); err != nil {
			log.Errorf("invalid --additional-instance-configuration: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderAmiTags) > 0 {
		if err := assignInputField(input, "AmiTags", _imagebuilderAmiTags); err != nil {
			log.Errorf("invalid --ami-tags: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderBlockDeviceMappings) > 0 {
		if err := assignInputField(input, "BlockDeviceMappings", _imagebuilderBlockDeviceMappings); err != nil {
			log.Errorf("invalid --block-device-mappings: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderComponents) > 0 {
		if err := assignInputField(input, "Components", _imagebuilderComponents); err != nil {
			log.Errorf("invalid --components: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderDescription) > 0 {
		input.Description = aws.String(_imagebuilderDescription)
	}
	if len(_imagebuilderTags) > 0 {
		if err := assignInputField(input, "Tags", _imagebuilderTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderWorkingDirectory) > 0 {
		input.WorkingDirectory = aws.String(_imagebuilderWorkingDirectory)
	}

	if resp, err := client.CreateImageRecipe(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new infrastructure configuration. An infrastructure configuration
// defines the environment in which your image will be built and tested.
func imagebuilder_CreateInfrastructureConfiguration(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.CreateInfrastructureConfigurationInput{
		// ClientToken: *string, // Required
		// InstanceProfileName: *string, // Required
		// Name: *string, // Required
	}

	if len(_imagebuilderClientToken) > 0 {
		input.ClientToken = aws.String(_imagebuilderClientToken)
	}
	if len(_imagebuilderInstanceProfileName) > 0 {
		input.InstanceProfileName = aws.String(_imagebuilderInstanceProfileName)
	}
	if len(_imagebuilderName) > 0 {
		input.Name = aws.String(_imagebuilderName)
	}
	if len(_imagebuilderDescription) > 0 {
		input.Description = aws.String(_imagebuilderDescription)
	}
	if len(_imagebuilderInstanceMetadataOptions) > 0 {
		if err := assignInputField(input, "InstanceMetadataOptions", _imagebuilderInstanceMetadataOptions); err != nil {
			log.Errorf("invalid --instance-metadata-options: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderInstanceTypes) > 0 {
		input.InstanceTypes = append([]string(nil), _imagebuilderInstanceTypes...)
	}
	if len(_imagebuilderKeyPair) > 0 {
		input.KeyPair = aws.String(_imagebuilderKeyPair)
	}
	if len(_imagebuilderLogging) > 0 {
		if err := assignInputField(input, "Logging", _imagebuilderLogging); err != nil {
			log.Errorf("invalid --logging: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderPlacement) > 0 {
		if err := assignInputField(input, "Placement", _imagebuilderPlacement); err != nil {
			log.Errorf("invalid --placement: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderResourceTags) > 0 {
		if err := assignInputField(input, "ResourceTags", _imagebuilderResourceTags); err != nil {
			log.Errorf("invalid --resource-tags: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _imagebuilderSecurityGroupIds...)
	}
	if len(_imagebuilderSnsTopicArn) > 0 {
		input.SnsTopicArn = aws.String(_imagebuilderSnsTopicArn)
	}
	if len(_imagebuilderSubnetId) > 0 {
		input.SubnetId = aws.String(_imagebuilderSubnetId)
	}
	if len(_imagebuilderTags) > 0 {
		if err := assignInputField(input, "Tags", _imagebuilderTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderTerminateInstanceOnFailure) > 0 {
		if err := assignInputField(input, "TerminateInstanceOnFailure", _imagebuilderTerminateInstanceOnFailure); err != nil {
			log.Errorf("invalid --terminate-instance-on-failure: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateInfrastructureConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a lifecycle policy resource.
func imagebuilder_CreateLifecyclePolicy(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.CreateLifecyclePolicyInput{
		// ClientToken: *string, // Required
		// ExecutionRole: *string, // Required
		// Name: *string, // Required
		// PolicyDetails: []types.LifecyclePolicyDetail, // Required
		// ResourceSelection: *types.LifecyclePolicyResourceSelection, // Required
		// ResourceType: types.LifecyclePolicyResourceType, // Required
	}

	if len(_imagebuilderClientToken) > 0 {
		input.ClientToken = aws.String(_imagebuilderClientToken)
	}
	if len(_imagebuilderExecutionRole) > 0 {
		input.ExecutionRole = aws.String(_imagebuilderExecutionRole)
	}
	if len(_imagebuilderName) > 0 {
		input.Name = aws.String(_imagebuilderName)
	}
	if len(_imagebuilderPolicyDetails) > 0 {
		if err := assignInputField(input, "PolicyDetails", _imagebuilderPolicyDetails); err != nil {
			log.Errorf("invalid --policy-details: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderResourceSelection) > 0 {
		if err := assignInputField(input, "ResourceSelection", _imagebuilderResourceSelection); err != nil {
			log.Errorf("invalid --resource-selection: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _imagebuilderResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderDescription) > 0 {
		input.Description = aws.String(_imagebuilderDescription)
	}
	if len(_imagebuilderStatus) > 0 {
		if err := assignInputField(input, "Status", _imagebuilderStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderTags) > 0 {
		if err := assignInputField(input, "Tags", _imagebuilderTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLifecyclePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a new workflow or a new version of an existing workflow.
func imagebuilder_CreateWorkflow(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.CreateWorkflowInput{
		// ClientToken: *string, // Required
		// Name: *string, // Required
		// SemanticVersion: *string, // Required
		// Type: types.WorkflowType, // Required
	}

	if len(_imagebuilderClientToken) > 0 {
		input.ClientToken = aws.String(_imagebuilderClientToken)
	}
	if len(_imagebuilderName) > 0 {
		input.Name = aws.String(_imagebuilderName)
	}
	if len(_imagebuilderSemanticVersion) > 0 {
		input.SemanticVersion = aws.String(_imagebuilderSemanticVersion)
	}
	if len(_imagebuilderType) > 0 {
		if err := assignInputField(input, "Type", _imagebuilderType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderChangeDescription) > 0 {
		input.ChangeDescription = aws.String(_imagebuilderChangeDescription)
	}
	if len(_imagebuilderData) > 0 {
		input.Data = aws.String(_imagebuilderData)
	}
	if len(_imagebuilderDescription) > 0 {
		input.Description = aws.String(_imagebuilderDescription)
	}
	if len(_imagebuilderDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _imagebuilderDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_imagebuilderKmsKeyId)
	}
	if len(_imagebuilderTags) > 0 {
		if err := assignInputField(input, "Tags", _imagebuilderTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderUri) > 0 {
		input.Uri = aws.String(_imagebuilderUri)
	}

	if resp, err := client.CreateWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a component build version.
func imagebuilder_DeleteComponent(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.DeleteComponentInput{
		// ComponentBuildVersionArn: *string, // Required
	}

	if len(_imagebuilderComponentBuildVersionArn) > 0 {
		input.ComponentBuildVersionArn = aws.String(_imagebuilderComponentBuildVersionArn)
	}

	if resp, err := client.DeleteComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a container recipe.
func imagebuilder_DeleteContainerRecipe(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.DeleteContainerRecipeInput{
		// ContainerRecipeArn: *string, // Required
	}

	if len(_imagebuilderContainerRecipeArn) > 0 {
		input.ContainerRecipeArn = aws.String(_imagebuilderContainerRecipeArn)
	}

	if resp, err := client.DeleteContainerRecipe(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a distribution configuration.
func imagebuilder_DeleteDistributionConfiguration(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.DeleteDistributionConfigurationInput{
		// DistributionConfigurationArn: *string, // Required
	}

	if len(_imagebuilderDistributionConfigurationArn) > 0 {
		input.DistributionConfigurationArn = aws.String(_imagebuilderDistributionConfigurationArn)
	}

	if resp, err := client.DeleteDistributionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Image Builder image resource. This does not delete any EC2 AMIs or
// ECR container images that are created during the image build process. You must
// clean those up separately, using the appropriate Amazon EC2 or Amazon ECR
// console actions, or API or CLI commands.
//
// - To deregister an EC2 Linux AMI, see [Deregister your Linux AMI]in the Amazon EC2 User Guide .
//
// - To deregister an EC2 Windows AMI, see [Deregister your Windows AMI]in the Amazon EC2 Windows Guide .
//
// - To delete a container image from Amazon ECR, see [Deleting an image]in the Amazon ECR User
// Guide.
//
// [Deregister your Linux AMI]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/deregister-ami.html
// [Deregister your Windows AMI]: https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/deregister-ami.html
// [Deleting an image]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/delete_image.html
func imagebuilder_DeleteImage(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.DeleteImageInput{
		// ImageBuildVersionArn: *string, // Required
	}

	if len(_imagebuilderImageBuildVersionArn) > 0 {
		input.ImageBuildVersionArn = aws.String(_imagebuilderImageBuildVersionArn)
	}

	if resp, err := client.DeleteImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an image pipeline.
func imagebuilder_DeleteImagePipeline(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.DeleteImagePipelineInput{
		// ImagePipelineArn: *string, // Required
	}

	if len(_imagebuilderImagePipelineArn) > 0 {
		input.ImagePipelineArn = aws.String(_imagebuilderImagePipelineArn)
	}

	if resp, err := client.DeleteImagePipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an image recipe.
func imagebuilder_DeleteImageRecipe(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.DeleteImageRecipeInput{
		// ImageRecipeArn: *string, // Required
	}

	if len(_imagebuilderImageRecipeArn) > 0 {
		input.ImageRecipeArn = aws.String(_imagebuilderImageRecipeArn)
	}

	if resp, err := client.DeleteImageRecipe(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an infrastructure configuration.
func imagebuilder_DeleteInfrastructureConfiguration(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.DeleteInfrastructureConfigurationInput{
		// InfrastructureConfigurationArn: *string, // Required
	}

	if len(_imagebuilderInfrastructureConfigurationArn) > 0 {
		input.InfrastructureConfigurationArn = aws.String(_imagebuilderInfrastructureConfigurationArn)
	}

	if resp, err := client.DeleteInfrastructureConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete the specified lifecycle policy resource.
func imagebuilder_DeleteLifecyclePolicy(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.DeleteLifecyclePolicyInput{
		// LifecyclePolicyArn: *string, // Required
	}

	if len(_imagebuilderLifecyclePolicyArn) > 0 {
		input.LifecyclePolicyArn = aws.String(_imagebuilderLifecyclePolicyArn)
	}

	if resp, err := client.DeleteLifecyclePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specific workflow resource.
func imagebuilder_DeleteWorkflow(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.DeleteWorkflowInput{
		// WorkflowBuildVersionArn: *string, // Required
	}

	if len(_imagebuilderWorkflowBuildVersionArn) > 0 {
		input.WorkflowBuildVersionArn = aws.String(_imagebuilderWorkflowBuildVersionArn)
	}

	if resp, err := client.DeleteWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// DistributeImage distributes existing AMIs to additional regions and accounts
// without rebuilding the image.
func imagebuilder_DistributeImage(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.DistributeImageInput{
		// ClientToken: *string, // Required
		// DistributionConfigurationArn: *string, // Required
		// ExecutionRole: *string, // Required
		// SourceImage: *string, // Required
	}

	if len(_imagebuilderClientToken) > 0 {
		input.ClientToken = aws.String(_imagebuilderClientToken)
	}
	if len(_imagebuilderDistributionConfigurationArn) > 0 {
		input.DistributionConfigurationArn = aws.String(_imagebuilderDistributionConfigurationArn)
	}
	if len(_imagebuilderExecutionRole) > 0 {
		input.ExecutionRole = aws.String(_imagebuilderExecutionRole)
	}
	if len(_imagebuilderSourceImage) > 0 {
		input.SourceImage = aws.String(_imagebuilderSourceImage)
	}
	if len(_imagebuilderLoggingConfiguration) > 0 {
		if err := assignInputField(input, "LoggingConfiguration", _imagebuilderLoggingConfiguration); err != nil {
			log.Errorf("invalid --logging-configuration: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderTags) > 0 {
		if err := assignInputField(input, "Tags", _imagebuilderTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.DistributeImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a component object.
func imagebuilder_GetComponent(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.GetComponentInput{
		// ComponentBuildVersionArn: *string, // Required
	}

	if len(_imagebuilderComponentBuildVersionArn) > 0 {
		input.ComponentBuildVersionArn = aws.String(_imagebuilderComponentBuildVersionArn)
	}

	if resp, err := client.GetComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a component policy.
func imagebuilder_GetComponentPolicy(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.GetComponentPolicyInput{
		// ComponentArn: *string, // Required
	}

	if len(_imagebuilderComponentArn) > 0 {
		input.ComponentArn = aws.String(_imagebuilderComponentArn)
	}

	if resp, err := client.GetComponentPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a container recipe.
func imagebuilder_GetContainerRecipe(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.GetContainerRecipeInput{
		// ContainerRecipeArn: *string, // Required
	}

	if len(_imagebuilderContainerRecipeArn) > 0 {
		input.ContainerRecipeArn = aws.String(_imagebuilderContainerRecipeArn)
	}

	if resp, err := client.GetContainerRecipe(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the policy for a container recipe.
func imagebuilder_GetContainerRecipePolicy(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.GetContainerRecipePolicyInput{
		// ContainerRecipeArn: *string, // Required
	}

	if len(_imagebuilderContainerRecipeArn) > 0 {
		input.ContainerRecipeArn = aws.String(_imagebuilderContainerRecipeArn)
	}

	if resp, err := client.GetContainerRecipePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a distribution configuration.
func imagebuilder_GetDistributionConfiguration(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.GetDistributionConfigurationInput{
		// DistributionConfigurationArn: *string, // Required
	}

	if len(_imagebuilderDistributionConfigurationArn) > 0 {
		input.DistributionConfigurationArn = aws.String(_imagebuilderDistributionConfigurationArn)
	}

	if resp, err := client.GetDistributionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an image.
func imagebuilder_GetImage(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.GetImageInput{
		// ImageBuildVersionArn: *string, // Required
	}

	if len(_imagebuilderImageBuildVersionArn) > 0 {
		input.ImageBuildVersionArn = aws.String(_imagebuilderImageBuildVersionArn)
	}

	if resp, err := client.GetImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an image pipeline.
func imagebuilder_GetImagePipeline(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.GetImagePipelineInput{
		// ImagePipelineArn: *string, // Required
	}

	if len(_imagebuilderImagePipelineArn) > 0 {
		input.ImagePipelineArn = aws.String(_imagebuilderImagePipelineArn)
	}

	if resp, err := client.GetImagePipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an image policy.
func imagebuilder_GetImagePolicy(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.GetImagePolicyInput{
		// ImageArn: *string, // Required
	}

	if len(_imagebuilderImageArn) > 0 {
		input.ImageArn = aws.String(_imagebuilderImageArn)
	}

	if resp, err := client.GetImagePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an image recipe.
func imagebuilder_GetImageRecipe(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.GetImageRecipeInput{
		// ImageRecipeArn: *string, // Required
	}

	if len(_imagebuilderImageRecipeArn) > 0 {
		input.ImageRecipeArn = aws.String(_imagebuilderImageRecipeArn)
	}

	if resp, err := client.GetImageRecipe(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an image recipe policy.
func imagebuilder_GetImageRecipePolicy(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.GetImageRecipePolicyInput{
		// ImageRecipeArn: *string, // Required
	}

	if len(_imagebuilderImageRecipeArn) > 0 {
		input.ImageRecipeArn = aws.String(_imagebuilderImageRecipeArn)
	}

	if resp, err := client.GetImageRecipePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an infrastructure configuration.
func imagebuilder_GetInfrastructureConfiguration(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.GetInfrastructureConfigurationInput{
		// InfrastructureConfigurationArn: *string, // Required
	}

	if len(_imagebuilderInfrastructureConfigurationArn) > 0 {
		input.InfrastructureConfigurationArn = aws.String(_imagebuilderInfrastructureConfigurationArn)
	}

	if resp, err := client.GetInfrastructureConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the runtime information that was logged for a specific runtime instance of
// the lifecycle policy.
func imagebuilder_GetLifecycleExecution(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.GetLifecycleExecutionInput{
		// LifecycleExecutionId: *string, // Required
	}

	if len(_imagebuilderLifecycleExecutionId) > 0 {
		input.LifecycleExecutionId = aws.String(_imagebuilderLifecycleExecutionId)
	}

	if resp, err := client.GetLifecycleExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get details for the specified image lifecycle policy.
func imagebuilder_GetLifecyclePolicy(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.GetLifecyclePolicyInput{
		// LifecyclePolicyArn: *string, // Required
	}

	if len(_imagebuilderLifecyclePolicyArn) > 0 {
		input.LifecyclePolicyArn = aws.String(_imagebuilderLifecyclePolicyArn)
	}

	if resp, err := client.GetLifecyclePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Verify the subscription and perform resource dependency checks on the requested
// Amazon Web Services Marketplace resource. For Amazon Web Services Marketplace
// components, the response contains fields to download the components and their
// artifacts.
func imagebuilder_GetMarketplaceResource(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.GetMarketplaceResourceInput{
		// ResourceArn: *string, // Required
		// ResourceType: types.MarketplaceResourceType, // Required
	}

	if len(_imagebuilderResourceArn) > 0 {
		input.ResourceArn = aws.String(_imagebuilderResourceArn)
	}
	if len(_imagebuilderResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _imagebuilderResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderResourceLocation) > 0 {
		input.ResourceLocation = aws.String(_imagebuilderResourceLocation)
	}

	if resp, err := client.GetMarketplaceResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get a workflow resource object.
func imagebuilder_GetWorkflow(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.GetWorkflowInput{
		// WorkflowBuildVersionArn: *string, // Required
	}

	if len(_imagebuilderWorkflowBuildVersionArn) > 0 {
		input.WorkflowBuildVersionArn = aws.String(_imagebuilderWorkflowBuildVersionArn)
	}

	if resp, err := client.GetWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the runtime information that was logged for a specific runtime instance of
// the workflow.
func imagebuilder_GetWorkflowExecution(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.GetWorkflowExecutionInput{
		// WorkflowExecutionId: *string, // Required
	}

	if len(_imagebuilderWorkflowExecutionId) > 0 {
		input.WorkflowExecutionId = aws.String(_imagebuilderWorkflowExecutionId)
	}

	if resp, err := client.GetWorkflowExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the runtime information that was logged for a specific runtime instance of
// the workflow step.
func imagebuilder_GetWorkflowStepExecution(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.GetWorkflowStepExecutionInput{
		// StepExecutionId: *string, // Required
	}

	if len(_imagebuilderStepExecutionId) > 0 {
		input.StepExecutionId = aws.String(_imagebuilderStepExecutionId)
	}

	if resp, err := client.GetWorkflowStepExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports a component and transforms its data into a component document.
func imagebuilder_ImportComponent(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.ImportComponentInput{
		// ClientToken: *string, // Required
		// Format: types.ComponentFormat, // Required
		// Name: *string, // Required
		// Platform: types.Platform, // Required
		// SemanticVersion: *string, // Required
		// Type: types.ComponentType, // Required
	}

	if len(_imagebuilderClientToken) > 0 {
		input.ClientToken = aws.String(_imagebuilderClientToken)
	}
	if len(_imagebuilderFormat) > 0 {
		if err := assignInputField(input, "Format", _imagebuilderFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderName) > 0 {
		input.Name = aws.String(_imagebuilderName)
	}
	if len(_imagebuilderPlatform) > 0 {
		if err := assignInputField(input, "Platform", _imagebuilderPlatform); err != nil {
			log.Errorf("invalid --platform: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderSemanticVersion) > 0 {
		input.SemanticVersion = aws.String(_imagebuilderSemanticVersion)
	}
	if len(_imagebuilderType) > 0 {
		if err := assignInputField(input, "Type", _imagebuilderType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderChangeDescription) > 0 {
		input.ChangeDescription = aws.String(_imagebuilderChangeDescription)
	}
	if len(_imagebuilderData) > 0 {
		input.Data = aws.String(_imagebuilderData)
	}
	if len(_imagebuilderDescription) > 0 {
		input.Description = aws.String(_imagebuilderDescription)
	}
	if len(_imagebuilderKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_imagebuilderKmsKeyId)
	}
	if len(_imagebuilderTags) > 0 {
		if err := assignInputField(input, "Tags", _imagebuilderTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderUri) > 0 {
		input.Uri = aws.String(_imagebuilderUri)
	}

	if resp, err := client.ImportComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Import a Windows operating system image from a verified Microsoft ISO disk
// file. The following disk images are supported:
//
// - Windows 11 Enterprise
func imagebuilder_ImportDiskImage(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.ImportDiskImageInput{
		// ClientToken: *string, // Required
		// InfrastructureConfigurationArn: *string, // Required
		// Name: *string, // Required
		// OsVersion: *string, // Required
		// Platform: *string, // Required
		// SemanticVersion: *string, // Required
		// Uri: *string, // Required
	}

	if len(_imagebuilderClientToken) > 0 {
		input.ClientToken = aws.String(_imagebuilderClientToken)
	}
	if len(_imagebuilderInfrastructureConfigurationArn) > 0 {
		input.InfrastructureConfigurationArn = aws.String(_imagebuilderInfrastructureConfigurationArn)
	}
	if len(_imagebuilderName) > 0 {
		input.Name = aws.String(_imagebuilderName)
	}
	if len(_imagebuilderOsVersion) > 0 {
		input.OsVersion = aws.String(_imagebuilderOsVersion)
	}
	if len(_imagebuilderPlatform) > 0 {
		input.Platform = aws.String(_imagebuilderPlatform)
	}
	if len(_imagebuilderSemanticVersion) > 0 {
		input.SemanticVersion = aws.String(_imagebuilderSemanticVersion)
	}
	if len(_imagebuilderUri) > 0 {
		input.Uri = aws.String(_imagebuilderUri)
	}
	if len(_imagebuilderDescription) > 0 {
		input.Description = aws.String(_imagebuilderDescription)
	}
	if len(_imagebuilderExecutionRole) > 0 {
		input.ExecutionRole = aws.String(_imagebuilderExecutionRole)
	}
	if len(_imagebuilderLoggingConfiguration) > 0 {
		if err := assignInputField(input, "LoggingConfiguration", _imagebuilderLoggingConfiguration); err != nil {
			log.Errorf("invalid --logging-configuration: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderTags) > 0 {
		if err := assignInputField(input, "Tags", _imagebuilderTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportDiskImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// When you export your virtual machine (VM) from its virtualization environment,
// that process creates a set of one or more disk container files that act as
// snapshots of your VM’s environment, settings, and data. The Amazon EC2 API [ImportImage]
// action uses those files to import your VM and create an AMI. To import using the
// CLI command, see [import-image]
//
// You can reference the task ID from the VM import to pull in the AMI that the
// import created as the base image for your Image Builder recipe.
//
// [ImportImage]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportImage.html
// [import-image]: https://docs.aws.amazon.com/cli/latest/reference/ec2/import-image.html
func imagebuilder_ImportVmImage(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.ImportVmImageInput{
		// ClientToken: *string, // Required
		// Name: *string, // Required
		// Platform: types.Platform, // Required
		// SemanticVersion: *string, // Required
		// VmImportTaskId: *string, // Required
	}

	if len(_imagebuilderClientToken) > 0 {
		input.ClientToken = aws.String(_imagebuilderClientToken)
	}
	if len(_imagebuilderName) > 0 {
		input.Name = aws.String(_imagebuilderName)
	}
	if len(_imagebuilderPlatform) > 0 {
		if err := assignInputField(input, "Platform", _imagebuilderPlatform); err != nil {
			log.Errorf("invalid --platform: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderSemanticVersion) > 0 {
		input.SemanticVersion = aws.String(_imagebuilderSemanticVersion)
	}
	if len(_imagebuilderVmImportTaskId) > 0 {
		input.VmImportTaskId = aws.String(_imagebuilderVmImportTaskId)
	}
	if len(_imagebuilderDescription) > 0 {
		input.Description = aws.String(_imagebuilderDescription)
	}
	if len(_imagebuilderLoggingConfiguration) > 0 {
		if err := assignInputField(input, "LoggingConfiguration", _imagebuilderLoggingConfiguration); err != nil {
			log.Errorf("invalid --logging-configuration: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderOsVersion) > 0 {
		input.OsVersion = aws.String(_imagebuilderOsVersion)
	}
	if len(_imagebuilderTags) > 0 {
		if err := assignInputField(input, "Tags", _imagebuilderTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportVmImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the list of component build versions for the specified component
// version Amazon Resource Name (ARN).
func imagebuilder_ListComponentBuildVersions(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.ListComponentBuildVersionsInput{}

	if len(_imagebuilderComponentVersionArn) > 0 {
		input.ComponentVersionArn = aws.String(_imagebuilderComponentVersionArn)
	}
	if len(_imagebuilderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _imagebuilderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderNextToken) > 0 {
		input.NextToken = aws.String(_imagebuilderNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListComponentBuildVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*imagebuilder.ListComponentBuildVersionsOutput
	p := imagebuilder.NewListComponentBuildVersionsPaginator(client, input)
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

// Returns the list of components that can be filtered by name, or by using the
// listed filters to streamline results. Newly created components can take up to
// two minutes to appear in the ListComponents API Results.
//
// The semantic version has four nodes: ../. You can assign values for the first
// three, and can filter on all of them.
//
// Filtering: With semantic versioning, you have the flexibility to use wildcards
// (x) to specify the most recent versions or nodes when selecting the base image
// or components for your recipe. When you use a wildcard in any node, all nodes to
// the right of the first wildcard must also be wildcards.
func imagebuilder_ListComponents(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.ListComponentsInput{}

	if len(_imagebuilderByName) > 0 {
		if err := assignInputField(input, "ByName", _imagebuilderByName); err != nil {
			log.Errorf("invalid --by-name: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderFilters) > 0 {
		if err := assignInputField(input, "Filters", _imagebuilderFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _imagebuilderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderNextToken) > 0 {
		input.NextToken = aws.String(_imagebuilderNextToken)
	}
	if len(_imagebuilderOwner) > 0 {
		if err := assignInputField(input, "Owner", _imagebuilderOwner); err != nil {
			log.Errorf("invalid --owner: %s", err.Error())
			return
		}
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

	var results []*imagebuilder.ListComponentsOutput
	p := imagebuilder.NewListComponentsPaginator(client, input)
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

// Returns a list of container recipes.
func imagebuilder_ListContainerRecipes(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.ListContainerRecipesInput{}

	if len(_imagebuilderFilters) > 0 {
		if err := assignInputField(input, "Filters", _imagebuilderFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _imagebuilderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderNextToken) > 0 {
		input.NextToken = aws.String(_imagebuilderNextToken)
	}
	if len(_imagebuilderOwner) > 0 {
		if err := assignInputField(input, "Owner", _imagebuilderOwner); err != nil {
			log.Errorf("invalid --owner: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListContainerRecipes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*imagebuilder.ListContainerRecipesOutput
	p := imagebuilder.NewListContainerRecipesPaginator(client, input)
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

// Returns a list of distribution configurations.
func imagebuilder_ListDistributionConfigurations(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.ListDistributionConfigurationsInput{}

	if len(_imagebuilderFilters) > 0 {
		if err := assignInputField(input, "Filters", _imagebuilderFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _imagebuilderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderNextToken) > 0 {
		input.NextToken = aws.String(_imagebuilderNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDistributionConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*imagebuilder.ListDistributionConfigurationsOutput
	p := imagebuilder.NewListDistributionConfigurationsPaginator(client, input)
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

// Returns a list of image build versions.
func imagebuilder_ListImageBuildVersions(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.ListImageBuildVersionsInput{}

	if len(_imagebuilderFilters) > 0 {
		if err := assignInputField(input, "Filters", _imagebuilderFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderImageVersionArn) > 0 {
		input.ImageVersionArn = aws.String(_imagebuilderImageVersionArn)
	}
	if len(_imagebuilderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _imagebuilderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderNextToken) > 0 {
		input.NextToken = aws.String(_imagebuilderNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListImageBuildVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*imagebuilder.ListImageBuildVersionsOutput
	p := imagebuilder.NewListImageBuildVersionsPaginator(client, input)
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

// List the Packages that are associated with an Image Build Version, as
// determined by Amazon Web Services Systems Manager Inventory at build time.
func imagebuilder_ListImagePackages(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.ListImagePackagesInput{
		// ImageBuildVersionArn: *string, // Required
	}

	if len(_imagebuilderImageBuildVersionArn) > 0 {
		input.ImageBuildVersionArn = aws.String(_imagebuilderImageBuildVersionArn)
	}
	if len(_imagebuilderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _imagebuilderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderNextToken) > 0 {
		input.NextToken = aws.String(_imagebuilderNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListImagePackages(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*imagebuilder.ListImagePackagesOutput
	p := imagebuilder.NewListImagePackagesPaginator(client, input)
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

// Returns a list of images created by the specified pipeline.
func imagebuilder_ListImagePipelineImages(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.ListImagePipelineImagesInput{
		// ImagePipelineArn: *string, // Required
	}

	if len(_imagebuilderImagePipelineArn) > 0 {
		input.ImagePipelineArn = aws.String(_imagebuilderImagePipelineArn)
	}
	if len(_imagebuilderFilters) > 0 {
		if err := assignInputField(input, "Filters", _imagebuilderFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _imagebuilderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderNextToken) > 0 {
		input.NextToken = aws.String(_imagebuilderNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListImagePipelineImages(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*imagebuilder.ListImagePipelineImagesOutput
	p := imagebuilder.NewListImagePipelineImagesPaginator(client, input)
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

// Returns a list of image pipelines.
func imagebuilder_ListImagePipelines(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.ListImagePipelinesInput{}

	if len(_imagebuilderFilters) > 0 {
		if err := assignInputField(input, "Filters", _imagebuilderFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _imagebuilderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderNextToken) > 0 {
		input.NextToken = aws.String(_imagebuilderNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListImagePipelines(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*imagebuilder.ListImagePipelinesOutput
	p := imagebuilder.NewListImagePipelinesPaginator(client, input)
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

// Returns a list of image recipes.
func imagebuilder_ListImageRecipes(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.ListImageRecipesInput{}

	if len(_imagebuilderFilters) > 0 {
		if err := assignInputField(input, "Filters", _imagebuilderFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _imagebuilderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderNextToken) > 0 {
		input.NextToken = aws.String(_imagebuilderNextToken)
	}
	if len(_imagebuilderOwner) > 0 {
		if err := assignInputField(input, "Owner", _imagebuilderOwner); err != nil {
			log.Errorf("invalid --owner: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListImageRecipes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*imagebuilder.ListImageRecipesOutput
	p := imagebuilder.NewListImageRecipesPaginator(client, input)
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

// Returns a list of image scan aggregations for your account. You can filter by
// the type of key that Image Builder uses to group results. For example, if you
// want to get a list of findings by severity level for one of your pipelines, you
// might specify your pipeline with the imagePipelineArn filter. If you don't
// specify a filter, Image Builder returns an aggregation for your account.
//
// To streamline results, you can use the following filters in your request:
//
// - accountId
//
// - imageBuildVersionArn
//
// - imagePipelineArn
//
// - vulnerabilityId
func imagebuilder_ListImageScanFindingAggregations(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.ListImageScanFindingAggregationsInput{}

	if len(_imagebuilderFilter) > 0 {
		if err := assignInputField(input, "Filter", _imagebuilderFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderNextToken) > 0 {
		input.NextToken = aws.String(_imagebuilderNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListImageScanFindingAggregations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*imagebuilder.ListImageScanFindingAggregationsOutput
	p := imagebuilder.NewListImageScanFindingAggregationsPaginator(client, input)
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

// Returns a list of image scan findings for your account.
func imagebuilder_ListImageScanFindings(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.ListImageScanFindingsInput{}

	if len(_imagebuilderFilters) > 0 {
		if err := assignInputField(input, "Filters", _imagebuilderFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _imagebuilderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderNextToken) > 0 {
		input.NextToken = aws.String(_imagebuilderNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListImageScanFindings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*imagebuilder.ListImageScanFindingsOutput
	p := imagebuilder.NewListImageScanFindingsPaginator(client, input)
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

// Returns the list of images that you have access to. Newly created images can
// take up to two minutes to appear in the ListImages API Results.
func imagebuilder_ListImages(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.ListImagesInput{}

	if len(_imagebuilderByName) > 0 {
		if err := assignInputField(input, "ByName", _imagebuilderByName); err != nil {
			log.Errorf("invalid --by-name: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderFilters) > 0 {
		if err := assignInputField(input, "Filters", _imagebuilderFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderIncludeDeprecated) > 0 {
		if err := assignInputField(input, "IncludeDeprecated", _imagebuilderIncludeDeprecated); err != nil {
			log.Errorf("invalid --include-deprecated: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _imagebuilderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderNextToken) > 0 {
		input.NextToken = aws.String(_imagebuilderNextToken)
	}
	if len(_imagebuilderOwner) > 0 {
		if err := assignInputField(input, "Owner", _imagebuilderOwner); err != nil {
			log.Errorf("invalid --owner: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListImages(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*imagebuilder.ListImagesOutput
	p := imagebuilder.NewListImagesPaginator(client, input)
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

// Returns a list of infrastructure configurations.
func imagebuilder_ListInfrastructureConfigurations(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.ListInfrastructureConfigurationsInput{}

	if len(_imagebuilderFilters) > 0 {
		if err := assignInputField(input, "Filters", _imagebuilderFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _imagebuilderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderNextToken) > 0 {
		input.NextToken = aws.String(_imagebuilderNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListInfrastructureConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*imagebuilder.ListInfrastructureConfigurationsOutput
	p := imagebuilder.NewListInfrastructureConfigurationsPaginator(client, input)
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

// List resources that the runtime instance of the image lifecycle identified for
// lifecycle actions.
func imagebuilder_ListLifecycleExecutionResources(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.ListLifecycleExecutionResourcesInput{
		// LifecycleExecutionId: *string, // Required
	}

	if len(_imagebuilderLifecycleExecutionId) > 0 {
		input.LifecycleExecutionId = aws.String(_imagebuilderLifecycleExecutionId)
	}
	if len(_imagebuilderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _imagebuilderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderNextToken) > 0 {
		input.NextToken = aws.String(_imagebuilderNextToken)
	}
	if len(_imagebuilderParentResourceId) > 0 {
		input.ParentResourceId = aws.String(_imagebuilderParentResourceId)
	}

	if disablePaginator() {
		if resp, err := client.ListLifecycleExecutionResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*imagebuilder.ListLifecycleExecutionResourcesOutput
	p := imagebuilder.NewListLifecycleExecutionResourcesPaginator(client, input)
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

// Get the lifecycle runtime history for the specified resource.
func imagebuilder_ListLifecycleExecutions(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.ListLifecycleExecutionsInput{
		// ResourceArn: *string, // Required
	}

	if len(_imagebuilderResourceArn) > 0 {
		input.ResourceArn = aws.String(_imagebuilderResourceArn)
	}
	if len(_imagebuilderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _imagebuilderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderNextToken) > 0 {
		input.NextToken = aws.String(_imagebuilderNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLifecycleExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*imagebuilder.ListLifecycleExecutionsOutput
	p := imagebuilder.NewListLifecycleExecutionsPaginator(client, input)
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

// Get a list of lifecycle policies in your Amazon Web Services account.
func imagebuilder_ListLifecyclePolicies(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.ListLifecyclePoliciesInput{}

	if len(_imagebuilderFilters) > 0 {
		if err := assignInputField(input, "Filters", _imagebuilderFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _imagebuilderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderNextToken) > 0 {
		input.NextToken = aws.String(_imagebuilderNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLifecyclePolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*imagebuilder.ListLifecyclePoliciesOutput
	p := imagebuilder.NewListLifecyclePoliciesPaginator(client, input)
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

// Returns the list of tags for the specified resource.
func imagebuilder_ListTagsForResource(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_imagebuilderResourceArn) > 0 {
		input.ResourceArn = aws.String(_imagebuilderResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get a list of workflow steps that are waiting for action for workflows in your
// Amazon Web Services account.
func imagebuilder_ListWaitingWorkflowSteps(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.ListWaitingWorkflowStepsInput{}

	if len(_imagebuilderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _imagebuilderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderNextToken) > 0 {
		input.NextToken = aws.String(_imagebuilderNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWaitingWorkflowSteps(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*imagebuilder.ListWaitingWorkflowStepsOutput
	p := imagebuilder.NewListWaitingWorkflowStepsPaginator(client, input)
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

// Returns a list of build versions for a specific workflow resource.
func imagebuilder_ListWorkflowBuildVersions(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.ListWorkflowBuildVersionsInput{}

	if len(_imagebuilderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _imagebuilderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderNextToken) > 0 {
		input.NextToken = aws.String(_imagebuilderNextToken)
	}
	if len(_imagebuilderWorkflowVersionArn) > 0 {
		input.WorkflowVersionArn = aws.String(_imagebuilderWorkflowVersionArn)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkflowBuildVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*imagebuilder.ListWorkflowBuildVersionsOutput
	p := imagebuilder.NewListWorkflowBuildVersionsPaginator(client, input)
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

// Returns a list of workflow runtime instance metadata objects for a specific
// image build version.
func imagebuilder_ListWorkflowExecutions(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.ListWorkflowExecutionsInput{
		// ImageBuildVersionArn: *string, // Required
	}

	if len(_imagebuilderImageBuildVersionArn) > 0 {
		input.ImageBuildVersionArn = aws.String(_imagebuilderImageBuildVersionArn)
	}
	if len(_imagebuilderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _imagebuilderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderNextToken) > 0 {
		input.NextToken = aws.String(_imagebuilderNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkflowExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*imagebuilder.ListWorkflowExecutionsOutput
	p := imagebuilder.NewListWorkflowExecutionsPaginator(client, input)
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

// Returns runtime data for each step in a runtime instance of the workflow that
// you specify in the request.
func imagebuilder_ListWorkflowStepExecutions(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.ListWorkflowStepExecutionsInput{
		// WorkflowExecutionId: *string, // Required
	}

	if len(_imagebuilderWorkflowExecutionId) > 0 {
		input.WorkflowExecutionId = aws.String(_imagebuilderWorkflowExecutionId)
	}
	if len(_imagebuilderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _imagebuilderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderNextToken) > 0 {
		input.NextToken = aws.String(_imagebuilderNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkflowStepExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*imagebuilder.ListWorkflowStepExecutionsOutput
	p := imagebuilder.NewListWorkflowStepExecutionsPaginator(client, input)
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

// Lists workflow build versions based on filtering parameters.
func imagebuilder_ListWorkflows(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.ListWorkflowsInput{}

	if len(_imagebuilderByName) > 0 {
		if err := assignInputField(input, "ByName", _imagebuilderByName); err != nil {
			log.Errorf("invalid --by-name: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderFilters) > 0 {
		if err := assignInputField(input, "Filters", _imagebuilderFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _imagebuilderMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderNextToken) > 0 {
		input.NextToken = aws.String(_imagebuilderNextToken)
	}
	if len(_imagebuilderOwner) > 0 {
		if err := assignInputField(input, "Owner", _imagebuilderOwner); err != nil {
			log.Errorf("invalid --owner: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListWorkflows(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*imagebuilder.ListWorkflowsOutput
	p := imagebuilder.NewListWorkflowsPaginator(client, input)
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

// Applies a policy to a component. We recommend that you call the RAM API [CreateResourceShare] to
// share resources. If you call the Image Builder API PutComponentPolicy , you must
// also call the RAM API [PromoteResourceShareCreatedFromPolicy]in order for the resource to be visible to all principals
// with whom the resource is shared.
//
// [PromoteResourceShareCreatedFromPolicy]: https://docs.aws.amazon.com/ram/latest/APIReference/API_PromoteResourceShareCreatedFromPolicy.html
// [CreateResourceShare]: https://docs.aws.amazon.com/ram/latest/APIReference/API_CreateResourceShare.html
func imagebuilder_PutComponentPolicy(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.PutComponentPolicyInput{
		// ComponentArn: *string, // Required
		// Policy: *string, // Required
	}

	if len(_imagebuilderComponentArn) > 0 {
		input.ComponentArn = aws.String(_imagebuilderComponentArn)
	}
	if len(_imagebuilderPolicy) > 0 {
		input.Policy = aws.String(_imagebuilderPolicy)
	}

	if resp, err := client.PutComponentPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies a policy to a container image. We recommend that you call the RAM API
// CreateResourceShare
// (https://docs.aws.amazon.com//ram/latest/APIReference/API_CreateResourceShare.html)
// to share resources. If you call the Image Builder API PutContainerImagePolicy ,
// you must also call the RAM API PromoteResourceShareCreatedFromPolicy
// (https://docs.aws.amazon.com//ram/latest/APIReference/API_PromoteResourceShareCreatedFromPolicy.html)
// in order for the resource to be visible to all principals with whom the resource
// is shared.
func imagebuilder_PutContainerRecipePolicy(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.PutContainerRecipePolicyInput{
		// ContainerRecipeArn: *string, // Required
		// Policy: *string, // Required
	}

	if len(_imagebuilderContainerRecipeArn) > 0 {
		input.ContainerRecipeArn = aws.String(_imagebuilderContainerRecipeArn)
	}
	if len(_imagebuilderPolicy) > 0 {
		input.Policy = aws.String(_imagebuilderPolicy)
	}

	if resp, err := client.PutContainerRecipePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies a policy to an image. We recommend that you call the RAM API [CreateResourceShare] to share
// resources. If you call the Image Builder API PutImagePolicy , you must also call
// the RAM API [PromoteResourceShareCreatedFromPolicy]in order for the resource to be visible to all principals with whom
// the resource is shared.
//
// [PromoteResourceShareCreatedFromPolicy]: https://docs.aws.amazon.com/ram/latest/APIReference/API_PromoteResourceShareCreatedFromPolicy.html
// [CreateResourceShare]: https://docs.aws.amazon.com/ram/latest/APIReference/API_CreateResourceShare.html
func imagebuilder_PutImagePolicy(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.PutImagePolicyInput{
		// ImageArn: *string, // Required
		// Policy: *string, // Required
	}

	if len(_imagebuilderImageArn) > 0 {
		input.ImageArn = aws.String(_imagebuilderImageArn)
	}
	if len(_imagebuilderPolicy) > 0 {
		input.Policy = aws.String(_imagebuilderPolicy)
	}

	if resp, err := client.PutImagePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies a policy to an image recipe. We recommend that you call the RAM API [CreateResourceShare] to
// share resources. If you call the Image Builder API PutImageRecipePolicy , you
// must also call the RAM API [PromoteResourceShareCreatedFromPolicy]in order for the resource to be visible to all
// principals with whom the resource is shared.
//
// [PromoteResourceShareCreatedFromPolicy]: https://docs.aws.amazon.com/ram/latest/APIReference/API_PromoteResourceShareCreatedFromPolicy.html
// [CreateResourceShare]: https://docs.aws.amazon.com/ram/latest/APIReference/API_CreateResourceShare.html
func imagebuilder_PutImageRecipePolicy(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.PutImageRecipePolicyInput{
		// ImageRecipeArn: *string, // Required
		// Policy: *string, // Required
	}

	if len(_imagebuilderImageRecipeArn) > 0 {
		input.ImageRecipeArn = aws.String(_imagebuilderImageRecipeArn)
	}
	if len(_imagebuilderPolicy) > 0 {
		input.Policy = aws.String(_imagebuilderPolicy)
	}

	if resp, err := client.PutImageRecipePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// RetryImage retries an image distribution without rebuilding the image.
func imagebuilder_RetryImage(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.RetryImageInput{
		// ClientToken: *string, // Required
		// ImageBuildVersionArn: *string, // Required
	}

	if len(_imagebuilderClientToken) > 0 {
		input.ClientToken = aws.String(_imagebuilderClientToken)
	}
	if len(_imagebuilderImageBuildVersionArn) > 0 {
		input.ImageBuildVersionArn = aws.String(_imagebuilderImageBuildVersionArn)
	}

	if resp, err := client.RetryImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Pauses or resumes image creation when the associated workflow runs a
// WaitForAction step.
func imagebuilder_SendWorkflowStepAction(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.SendWorkflowStepActionInput{
		// Action: types.WorkflowStepActionType, // Required
		// ClientToken: *string, // Required
		// ImageBuildVersionArn: *string, // Required
		// StepExecutionId: *string, // Required
	}

	if len(_imagebuilderAction) > 0 {
		if err := assignInputField(input, "Action", _imagebuilderAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderClientToken) > 0 {
		input.ClientToken = aws.String(_imagebuilderClientToken)
	}
	if len(_imagebuilderImageBuildVersionArn) > 0 {
		input.ImageBuildVersionArn = aws.String(_imagebuilderImageBuildVersionArn)
	}
	if len(_imagebuilderStepExecutionId) > 0 {
		input.StepExecutionId = aws.String(_imagebuilderStepExecutionId)
	}
	if len(_imagebuilderReason) > 0 {
		input.Reason = aws.String(_imagebuilderReason)
	}

	if resp, err := client.SendWorkflowStepAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Manually triggers a pipeline to create an image.
func imagebuilder_StartImagePipelineExecution(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.StartImagePipelineExecutionInput{
		// ClientToken: *string, // Required
		// ImagePipelineArn: *string, // Required
	}

	if len(_imagebuilderClientToken) > 0 {
		input.ClientToken = aws.String(_imagebuilderClientToken)
	}
	if len(_imagebuilderImagePipelineArn) > 0 {
		input.ImagePipelineArn = aws.String(_imagebuilderImagePipelineArn)
	}
	if len(_imagebuilderTags) > 0 {
		if err := assignInputField(input, "Tags", _imagebuilderTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartImagePipelineExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Begin asynchronous resource state update for lifecycle changes to the specified
// image resources.
func imagebuilder_StartResourceStateUpdate(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.StartResourceStateUpdateInput{
		// ClientToken: *string, // Required
		// ResourceArn: *string, // Required
		// State: *types.ResourceState, // Required
	}

	if len(_imagebuilderClientToken) > 0 {
		input.ClientToken = aws.String(_imagebuilderClientToken)
	}
	if len(_imagebuilderResourceArn) > 0 {
		input.ResourceArn = aws.String(_imagebuilderResourceArn)
	}
	if len(_imagebuilderState) > 0 {
		if err := assignInputField(input, "State", _imagebuilderState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderExclusionRules) > 0 {
		if err := assignInputField(input, "ExclusionRules", _imagebuilderExclusionRules); err != nil {
			log.Errorf("invalid --exclusion-rules: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderExecutionRole) > 0 {
		input.ExecutionRole = aws.String(_imagebuilderExecutionRole)
	}
	if len(_imagebuilderIncludeResources) > 0 {
		if err := assignInputField(input, "IncludeResources", _imagebuilderIncludeResources); err != nil {
			log.Errorf("invalid --include-resources: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderUpdateAt) > 0 {
		if err := assignInputField(input, "UpdateAt", _imagebuilderUpdateAt); err != nil {
			log.Errorf("invalid --update-at: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartResourceStateUpdate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a tag to a resource.
func imagebuilder_TagResource(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_imagebuilderResourceArn) > 0 {
		input.ResourceArn = aws.String(_imagebuilderResourceArn)
	}
	if len(_imagebuilderTags) > 0 {
		if err := assignInputField(input, "Tags", _imagebuilderTags); err != nil {
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

// Removes a tag from a resource.
func imagebuilder_UntagResource(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_imagebuilderResourceArn) > 0 {
		input.ResourceArn = aws.String(_imagebuilderResourceArn)
	}
	if len(_imagebuilderTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _imagebuilderTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a new distribution configuration. Distribution configurations define
// and configure the outputs of your pipeline.
func imagebuilder_UpdateDistributionConfiguration(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.UpdateDistributionConfigurationInput{
		// ClientToken: *string, // Required
		// DistributionConfigurationArn: *string, // Required
		// Distributions: []types.Distribution, // Required
	}

	if len(_imagebuilderClientToken) > 0 {
		input.ClientToken = aws.String(_imagebuilderClientToken)
	}
	if len(_imagebuilderDistributionConfigurationArn) > 0 {
		input.DistributionConfigurationArn = aws.String(_imagebuilderDistributionConfigurationArn)
	}
	if len(_imagebuilderDistributions) > 0 {
		if err := assignInputField(input, "Distributions", _imagebuilderDistributions); err != nil {
			log.Errorf("invalid --distributions: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderDescription) > 0 {
		input.Description = aws.String(_imagebuilderDescription)
	}

	if resp, err := client.UpdateDistributionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an image pipeline. Image pipelines enable you to automate the creation
// and distribution of images. You must specify exactly one recipe for your image,
// using either a containerRecipeArn or an imageRecipeArn .
//
// UpdateImagePipeline does not support selective updates for the pipeline. You
// must specify all of the required properties in the update request, not just the
// properties that have changed.
func imagebuilder_UpdateImagePipeline(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.UpdateImagePipelineInput{
		// ClientToken: *string, // Required
		// ImagePipelineArn: *string, // Required
		// InfrastructureConfigurationArn: *string, // Required
	}

	if len(_imagebuilderClientToken) > 0 {
		input.ClientToken = aws.String(_imagebuilderClientToken)
	}
	if len(_imagebuilderImagePipelineArn) > 0 {
		input.ImagePipelineArn = aws.String(_imagebuilderImagePipelineArn)
	}
	if len(_imagebuilderInfrastructureConfigurationArn) > 0 {
		input.InfrastructureConfigurationArn = aws.String(_imagebuilderInfrastructureConfigurationArn)
	}
	if len(_imagebuilderContainerRecipeArn) > 0 {
		input.ContainerRecipeArn = aws.String(_imagebuilderContainerRecipeArn)
	}
	if len(_imagebuilderDescription) > 0 {
		input.Description = aws.String(_imagebuilderDescription)
	}
	if len(_imagebuilderDistributionConfigurationArn) > 0 {
		input.DistributionConfigurationArn = aws.String(_imagebuilderDistributionConfigurationArn)
	}
	if len(_imagebuilderEnhancedImageMetadataEnabled) > 0 {
		if err := assignInputField(input, "EnhancedImageMetadataEnabled", _imagebuilderEnhancedImageMetadataEnabled); err != nil {
			log.Errorf("invalid --enhanced-image-metadata-enabled: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderExecutionRole) > 0 {
		input.ExecutionRole = aws.String(_imagebuilderExecutionRole)
	}
	if len(_imagebuilderImageRecipeArn) > 0 {
		input.ImageRecipeArn = aws.String(_imagebuilderImageRecipeArn)
	}
	if len(_imagebuilderImageScanningConfiguration) > 0 {
		if err := assignInputField(input, "ImageScanningConfiguration", _imagebuilderImageScanningConfiguration); err != nil {
			log.Errorf("invalid --image-scanning-configuration: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderImageTestsConfiguration) > 0 {
		if err := assignInputField(input, "ImageTestsConfiguration", _imagebuilderImageTestsConfiguration); err != nil {
			log.Errorf("invalid --image-tests-configuration: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderLoggingConfiguration) > 0 {
		if err := assignInputField(input, "LoggingConfiguration", _imagebuilderLoggingConfiguration); err != nil {
			log.Errorf("invalid --logging-configuration: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderSchedule) > 0 {
		if err := assignInputField(input, "Schedule", _imagebuilderSchedule); err != nil {
			log.Errorf("invalid --schedule: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderStatus) > 0 {
		if err := assignInputField(input, "Status", _imagebuilderStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderWorkflows) > 0 {
		if err := assignInputField(input, "Workflows", _imagebuilderWorkflows); err != nil {
			log.Errorf("invalid --workflows: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateImagePipeline(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a new infrastructure configuration. An infrastructure configuration
// defines the environment in which your image will be built and tested.
func imagebuilder_UpdateInfrastructureConfiguration(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.UpdateInfrastructureConfigurationInput{
		// ClientToken: *string, // Required
		// InfrastructureConfigurationArn: *string, // Required
		// InstanceProfileName: *string, // Required
	}

	if len(_imagebuilderClientToken) > 0 {
		input.ClientToken = aws.String(_imagebuilderClientToken)
	}
	if len(_imagebuilderInfrastructureConfigurationArn) > 0 {
		input.InfrastructureConfigurationArn = aws.String(_imagebuilderInfrastructureConfigurationArn)
	}
	if len(_imagebuilderInstanceProfileName) > 0 {
		input.InstanceProfileName = aws.String(_imagebuilderInstanceProfileName)
	}
	if len(_imagebuilderDescription) > 0 {
		input.Description = aws.String(_imagebuilderDescription)
	}
	if len(_imagebuilderInstanceMetadataOptions) > 0 {
		if err := assignInputField(input, "InstanceMetadataOptions", _imagebuilderInstanceMetadataOptions); err != nil {
			log.Errorf("invalid --instance-metadata-options: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderInstanceTypes) > 0 {
		input.InstanceTypes = append([]string(nil), _imagebuilderInstanceTypes...)
	}
	if len(_imagebuilderKeyPair) > 0 {
		input.KeyPair = aws.String(_imagebuilderKeyPair)
	}
	if len(_imagebuilderLogging) > 0 {
		if err := assignInputField(input, "Logging", _imagebuilderLogging); err != nil {
			log.Errorf("invalid --logging: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderPlacement) > 0 {
		if err := assignInputField(input, "Placement", _imagebuilderPlacement); err != nil {
			log.Errorf("invalid --placement: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderResourceTags) > 0 {
		if err := assignInputField(input, "ResourceTags", _imagebuilderResourceTags); err != nil {
			log.Errorf("invalid --resource-tags: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _imagebuilderSecurityGroupIds...)
	}
	if len(_imagebuilderSnsTopicArn) > 0 {
		input.SnsTopicArn = aws.String(_imagebuilderSnsTopicArn)
	}
	if len(_imagebuilderSubnetId) > 0 {
		input.SubnetId = aws.String(_imagebuilderSubnetId)
	}
	if len(_imagebuilderTerminateInstanceOnFailure) > 0 {
		if err := assignInputField(input, "TerminateInstanceOnFailure", _imagebuilderTerminateInstanceOnFailure); err != nil {
			log.Errorf("invalid --terminate-instance-on-failure: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateInfrastructureConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the specified lifecycle policy.
func imagebuilder_UpdateLifecyclePolicy(cfg aws.Config, client *imagebuilder.Client) {
	input := &imagebuilder.UpdateLifecyclePolicyInput{
		// ClientToken: *string, // Required
		// ExecutionRole: *string, // Required
		// LifecyclePolicyArn: *string, // Required
		// PolicyDetails: []types.LifecyclePolicyDetail, // Required
		// ResourceSelection: *types.LifecyclePolicyResourceSelection, // Required
		// ResourceType: types.LifecyclePolicyResourceType, // Required
	}

	if len(_imagebuilderClientToken) > 0 {
		input.ClientToken = aws.String(_imagebuilderClientToken)
	}
	if len(_imagebuilderExecutionRole) > 0 {
		input.ExecutionRole = aws.String(_imagebuilderExecutionRole)
	}
	if len(_imagebuilderLifecyclePolicyArn) > 0 {
		input.LifecyclePolicyArn = aws.String(_imagebuilderLifecyclePolicyArn)
	}
	if len(_imagebuilderPolicyDetails) > 0 {
		if err := assignInputField(input, "PolicyDetails", _imagebuilderPolicyDetails); err != nil {
			log.Errorf("invalid --policy-details: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderResourceSelection) > 0 {
		if err := assignInputField(input, "ResourceSelection", _imagebuilderResourceSelection); err != nil {
			log.Errorf("invalid --resource-selection: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _imagebuilderResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_imagebuilderDescription) > 0 {
		input.Description = aws.String(_imagebuilderDescription)
	}
	if len(_imagebuilderStatus) > 0 {
		if err := assignInputField(input, "Status", _imagebuilderStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateLifecyclePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_imagebuilderCmd)
	_imagebuilderCmd.Flags().SortFlags = false

	_imagebuilderCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_imagebuilderCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_imagebuilderCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderAction, "action", "", "", "Action")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderAdditionalInstanceConfiguration, "additional-instance-configuration", "", "", "Additional Instance Configuration")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderAmiTags, "ami-tags", "", "", "AMI Tags")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderBlockDeviceMappings, "block-device-mappings", "", "", "Block Device Mappings")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderByName, "by-name", "", "", "By Name")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderChangeDescription, "change-description", "", "", "Change Description")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderClientToken, "client-token", "", "", "Client Token")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderComponentArn, "component-arn", "", "", "Component ARN")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderComponentBuildVersionArn, "component-build-version-arn", "", "", "Component Build Version ARN")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderComponentVersionArn, "component-version-arn", "", "", "Component Version ARN")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderComponents, "components", "", "", "Components")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderContainerRecipeArn, "container-recipe-arn", "", "", "Container Recipe ARN")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderContainerType, "container-type", "", "", "Container Type")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderData, "data", "", "", "Data")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderDescription, "description", "", "", "Description")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderDistributionConfigurationArn, "distribution-configuration-arn", "", "", "Distribution Configuration ARN")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderDistributions, "distributions", "", "", "Distributions")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderDockerfileTemplateData, "dockerfile-template-data", "", "", "Dockerfile Template Data")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderDockerfileTemplateUri, "dockerfile-template-uri", "", "", "Dockerfile Template URI")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderDryRun, "dry-run", "", "", "Dry Run")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderEnhancedImageMetadataEnabled, "enhanced-image-metadata-enabled", "", "", "Enhanced Image Metadata Enabled")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderExclusionRules, "exclusion-rules", "", "", "Exclusion Rules")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderExecutionRole, "execution-role", "", "", "Execution Role")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderFilter, "filter", "", "", "Filter")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderFilters, "filters", "", "", "Filters")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderFormat, "format", "", "", "Format")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderImageArn, "image-arn", "", "", "Image ARN")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderImageBuildVersionArn, "image-build-version-arn", "", "", "Image Build Version ARN")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderImageOsVersionOverride, "image-os-version-override", "", "", "Image OS Version Override")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderImagePipelineArn, "image-pipeline-arn", "", "", "Image Pipeline ARN")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderImageRecipeArn, "image-recipe-arn", "", "", "Image Recipe ARN")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderImageScanningConfiguration, "image-scanning-configuration", "", "", "Image Scanning Configuration")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderImageTestsConfiguration, "image-tests-configuration", "", "", "Image Tests Configuration")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderImageVersionArn, "image-version-arn", "", "", "Image Version ARN")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderIncludeDeprecated, "include-deprecated", "", "", "Include Deprecated")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderIncludeResources, "include-resources", "", "", "Include Resources")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderInfrastructureConfigurationArn, "infrastructure-configuration-arn", "", "", "Infrastructure Configuration ARN")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderInstanceConfiguration, "instance-configuration", "", "", "Instance Configuration")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderInstanceMetadataOptions, "instance-metadata-options", "", "", "Instance Metadata Options")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderInstanceProfileName, "instance-profile-name", "", "", "Instance Profile Name")
	_imagebuilderCmd.Flags().StringSliceVarP(&_imagebuilderInstanceTypes, "instance-types", "", nil, "Instance Types")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderKeyPair, "key-pair", "", "", "Key Pair")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderLifecycleExecutionId, "lifecycle-execution-id", "", "", "Lifecycle Execution ID")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderLifecyclePolicyArn, "lifecycle-policy-arn", "", "", "Lifecycle Policy ARN")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderLogging, "logging", "", "", "Logging")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderLoggingConfiguration, "logging-configuration", "", "", "Logging Configuration")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderMaxResults, "max-results", "", "", "Max Results")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderName, "name", "", "", "Name")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderNextToken, "next-token", "", "", "Next Token")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderOsVersion, "os-version", "", "", "OS Version")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderOwner, "owner", "", "", "Owner")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderParentImage, "parent-image", "", "", "Parent Image")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderParentResourceId, "parent-resource-id", "", "", "Parent Resource ID")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderPlacement, "placement", "", "", "Placement")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderPlatform, "platform", "", "", "Platform")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderPlatformOverride, "platform-override", "", "", "Platform Override")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderPolicy, "policy", "", "", "Policy")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderPolicyDetails, "policy-details", "", "", "Policy Details")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderReason, "reason", "", "", "Reason")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderResourceArn, "resource-arn", "", "", "Resource ARN")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderResourceLocation, "resource-location", "", "", "Resource Location")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderResourceSelection, "resource-selection", "", "", "Resource Selection")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderResourceTags, "resource-tags", "", "", "Resource Tags")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderResourceType, "resource-type", "", "", "Resource Type")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderSchedule, "schedule", "", "", "Schedule")
	_imagebuilderCmd.Flags().StringSliceVarP(&_imagebuilderSecurityGroupIds, "security-group-ids", "", nil, "Security Group Ids")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderSemanticVersion, "semantic-version", "", "", "Semantic Version")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderSnsTopicArn, "sns-topic-arn", "", "", "SNS Topic ARN")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderSourceImage, "source-image", "", "", "Source Image")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderState, "state", "", "", "State")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderStatus, "status", "", "", "Status")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderStepExecutionId, "step-execution-id", "", "", "Step Execution ID")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderSubnetId, "subnet-id", "", "", "Subnet ID")
	_imagebuilderCmd.Flags().StringSliceVarP(&_imagebuilderSupportedOsVersions, "supported-os-versions", "", nil, "Supported OS Versions")
	_imagebuilderCmd.Flags().StringSliceVarP(&_imagebuilderTagKeys, "tag-keys", "", nil, "Tag Keys")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderTags, "tags", "", "", "Tags")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderTargetRepository, "target-repository", "", "", "Target Repository")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderTerminateInstanceOnFailure, "terminate-instance-on-failure", "", "", "Terminate Instance On Failure")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderType, "type", "", "", "Type")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderUpdateAt, "update-at", "", "", "Update At")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderUri, "uri", "", "", "URI")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderVmImportTaskId, "vm-import-task-id", "", "", "Vm Import Task ID")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderWorkflowBuildVersionArn, "workflow-build-version-arn", "", "", "Workflow Build Version ARN")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderWorkflowExecutionId, "workflow-execution-id", "", "", "Workflow Execution ID")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderWorkflowVersionArn, "workflow-version-arn", "", "", "Workflow Version ARN")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderWorkflows, "workflows", "", "", "Workflows")
	_imagebuilderCmd.Flags().StringVarP(&_imagebuilderWorkingDirectory, "working-directory", "", "", "Working Directory")

	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderCancelImageCreation, "cancel-image-creation", "", false, "Cancel Image Creation")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderCancelLifecycleExecution, "cancel-lifecycle-execution", "", false, "Cancel Lifecycle Execution")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderCreateComponent, "create-component", "", false, "Create Component")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderCreateContainerRecipe, "create-container-recipe", "", false, "Create Container Recipe")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderCreateDistributionConfiguration, "create-distribution-configuration", "", false, "Create Distribution Configuration")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderCreateImage, "create-image", "", false, "Create Image")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderCreateImagePipeline, "create-image-pipeline", "", false, "Create Image Pipeline")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderCreateImageRecipe, "create-image-recipe", "", false, "Create Image Recipe")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderCreateInfrastructureConfiguration, "create-infrastructure-configuration", "", false, "Create Infrastructure Configuration")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderCreateLifecyclePolicy, "create-lifecycle-policy", "", false, "Create Lifecycle Policy")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderCreateWorkflow, "create-workflow", "", false, "Create Workflow")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderDeleteComponent, "delete-component", "", false, "Delete Component")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderDeleteContainerRecipe, "delete-container-recipe", "", false, "Delete Container Recipe")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderDeleteDistributionConfiguration, "delete-distribution-configuration", "", false, "Delete Distribution Configuration")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderDeleteImage, "delete-image", "", false, "Delete Image")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderDeleteImagePipeline, "delete-image-pipeline", "", false, "Delete Image Pipeline")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderDeleteImageRecipe, "delete-image-recipe", "", false, "Delete Image Recipe")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderDeleteInfrastructureConfiguration, "delete-infrastructure-configuration", "", false, "Delete Infrastructure Configuration")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderDeleteLifecyclePolicy, "delete-lifecycle-policy", "", false, "Delete Lifecycle Policy")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderDeleteWorkflow, "delete-workflow", "", false, "Delete Workflow")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderDistributeImage, "distribute-image", "", false, "Distribute Image")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderGetComponent, "get-component", "", false, "Get Component")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderGetComponentPolicy, "get-component-policy", "", false, "Get Component Policy")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderGetContainerRecipe, "get-container-recipe", "", false, "Get Container Recipe")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderGetContainerRecipePolicy, "get-container-recipe-policy", "", false, "Get Container Recipe Policy")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderGetDistributionConfiguration, "get-distribution-configuration", "", false, "Get Distribution Configuration")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderGetImage, "get-image", "", false, "Get Image")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderGetImagePipeline, "get-image-pipeline", "", false, "Get Image Pipeline")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderGetImagePolicy, "get-image-policy", "", false, "Get Image Policy")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderGetImageRecipe, "get-image-recipe", "", false, "Get Image Recipe")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderGetImageRecipePolicy, "get-image-recipe-policy", "", false, "Get Image Recipe Policy")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderGetInfrastructureConfiguration, "get-infrastructure-configuration", "", false, "Get Infrastructure Configuration")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderGetLifecycleExecution, "get-lifecycle-execution", "", false, "Get Lifecycle Execution")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderGetLifecyclePolicy, "get-lifecycle-policy", "", false, "Get Lifecycle Policy")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderGetMarketplaceResource, "get-marketplace-resource", "", false, "Get Marketplace Resource")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderGetWorkflow, "get-workflow", "", false, "Get Workflow")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderGetWorkflowExecution, "get-workflow-execution", "", false, "Get Workflow Execution")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderGetWorkflowStepExecution, "get-workflow-step-execution", "", false, "Get Workflow Step Execution")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderImportComponent, "import-component", "", false, "Import Component")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderImportDiskImage, "import-disk-image", "", false, "Import Disk Image")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderImportVmImage, "import-vm-image", "", false, "Import Vm Image")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderListComponentBuildVersions, "list-component-build-versions", "", false, "List Component Build Versions")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderListComponents, "list-components", "", false, "List Components")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderListContainerRecipes, "list-container-recipes", "", false, "List Container Recipes")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderListDistributionConfigurations, "list-distribution-configurations", "", false, "List Distribution Configurations")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderListImageBuildVersions, "list-image-build-versions", "", false, "List Image Build Versions")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderListImagePackages, "list-image-packages", "", false, "List Image Packages")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderListImagePipelineImages, "list-image-pipeline-images", "", false, "List Image Pipeline Images")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderListImagePipelines, "list-image-pipelines", "", false, "List Image Pipelines")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderListImageRecipes, "list-image-recipes", "", false, "List Image Recipes")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderListImageScanFindingAggregations, "list-image-scan-finding-aggregations", "", false, "List Image Scan Finding Aggregations")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderListImageScanFindings, "list-image-scan-findings", "", false, "List Image Scan Findings")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderListImages, "list-images", "", false, "List Images")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderListInfrastructureConfigurations, "list-infrastructure-configurations", "", false, "List Infrastructure Configurations")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderListLifecycleExecutionResources, "list-lifecycle-execution-resources", "", false, "List Lifecycle Execution Resources")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderListLifecycleExecutions, "list-lifecycle-executions", "", false, "List Lifecycle Executions")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderListLifecyclePolicies, "list-lifecycle-policies", "", false, "List Lifecycle Policies")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderListWaitingWorkflowSteps, "list-waiting-workflow-steps", "", false, "List Waiting Workflow Steps")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderListWorkflowBuildVersions, "list-workflow-build-versions", "", false, "List Workflow Build Versions")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderListWorkflowExecutions, "list-workflow-executions", "", false, "List Workflow Executions")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderListWorkflowStepExecutions, "list-workflow-step-executions", "", false, "List Workflow Step Executions")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderListWorkflows, "list-workflows", "", false, "List Workflows")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderPutComponentPolicy, "put-component-policy", "", false, "Put Component Policy")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderPutContainerRecipePolicy, "put-container-recipe-policy", "", false, "Put Container Recipe Policy")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderPutImagePolicy, "put-image-policy", "", false, "Put Image Policy")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderPutImageRecipePolicy, "put-image-recipe-policy", "", false, "Put Image Recipe Policy")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderRetryImage, "retry-image", "", false, "Retry Image")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderSendWorkflowStepAction, "send-workflow-step-action", "", false, "Send Workflow Step Action")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderStartImagePipelineExecution, "start-image-pipeline-execution", "", false, "Start Image Pipeline Execution")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderStartResourceStateUpdate, "start-resource-state-update", "", false, "Start Resource State Update")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderTagResource, "tag-resource", "", false, "Tag Resource")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderUntagResource, "untag-resource", "", false, "Untag Resource")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderUpdateDistributionConfiguration, "update-distribution-configuration", "", false, "Update Distribution Configuration")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderUpdateImagePipeline, "update-image-pipeline", "", false, "Update Image Pipeline")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderUpdateInfrastructureConfiguration, "update-infrastructure-configuration", "", false, "Update Infrastructure Configuration")
	_imagebuilderCmd.Flags().BoolVarP(&_imagebuilderUpdateLifecyclePolicy, "update-lifecycle-policy", "", false, "Update Lifecycle Policy")

}
