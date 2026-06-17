package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ecrCmd represents the ecr command
var _ecrCmd = &cobra.Command{
	Use:   "ecr",
	Short: "AWS ecr CLI",
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
		client := ecr.NewFromConfig(cfg)
		if _ecrBatchCheckLayerAvailability {
			ecr_BatchCheckLayerAvailability(cfg, client)
			return
		}
		if _ecrBatchDeleteImage {
			ecr_BatchDeleteImage(cfg, client)
			return
		}
		if _ecrBatchGetImage {
			ecr_BatchGetImage(cfg, client)
			return
		}
		if _ecrBatchGetRepositoryScanningConfiguration {
			ecr_BatchGetRepositoryScanningConfiguration(cfg, client)
			return
		}
		if _ecrCompleteLayerUpload {
			ecr_CompleteLayerUpload(cfg, client)
			return
		}
		if _ecrCreatePullThroughCacheRule {
			ecr_CreatePullThroughCacheRule(cfg, client)
			return
		}
		if _ecrCreateRepository {
			ecr_CreateRepository(cfg, client)
			return
		}
		if _ecrCreateRepositoryCreationTemplate {
			ecr_CreateRepositoryCreationTemplate(cfg, client)
			return
		}
		if _ecrDeleteLifecyclePolicy {
			ecr_DeleteLifecyclePolicy(cfg, client)
			return
		}
		if _ecrDeletePullThroughCacheRule {
			ecr_DeletePullThroughCacheRule(cfg, client)
			return
		}
		if _ecrDeleteRegistryPolicy {
			ecr_DeleteRegistryPolicy(cfg, client)
			return
		}
		if _ecrDeleteRepository {
			ecr_DeleteRepository(cfg, client)
			return
		}
		if _ecrDeleteRepositoryCreationTemplate {
			ecr_DeleteRepositoryCreationTemplate(cfg, client)
			return
		}
		if _ecrDeleteRepositoryPolicy {
			ecr_DeleteRepositoryPolicy(cfg, client)
			return
		}
		if _ecrDeleteSigningConfiguration {
			ecr_DeleteSigningConfiguration(cfg, client)
			return
		}
		if _ecrDeregisterPullTimeUpdateExclusion {
			ecr_DeregisterPullTimeUpdateExclusion(cfg, client)
			return
		}
		if _ecrDescribeImageReplicationStatus {
			ecr_DescribeImageReplicationStatus(cfg, client)
			return
		}
		if _ecrDescribeImageScanFindings {
			ecr_DescribeImageScanFindings(cfg, client)
			return
		}
		if _ecrDescribeImageSigningStatus {
			ecr_DescribeImageSigningStatus(cfg, client)
			return
		}
		if _ecrDescribeImages {
			ecr_DescribeImages(cfg, client)
			return
		}
		if _ecrDescribePullThroughCacheRules {
			ecr_DescribePullThroughCacheRules(cfg, client)
			return
		}
		if _ecrDescribeRegistry {
			ecr_DescribeRegistry(cfg, client)
			return
		}
		if _ecrDescribeRepositories {
			ecr_DescribeRepositories(cfg, client)
			return
		}
		if _ecrDescribeRepositoryCreationTemplates {
			ecr_DescribeRepositoryCreationTemplates(cfg, client)
			return
		}
		if _ecrGetAccountSetting {
			ecr_GetAccountSetting(cfg, client)
			return
		}
		if _ecrGetAuthorizationToken {
			ecr_GetAuthorizationToken(cfg, client)
			return
		}
		if _ecrGetDownloadUrlForLayer {
			ecr_GetDownloadUrlForLayer(cfg, client)
			return
		}
		if _ecrGetLifecyclePolicy {
			ecr_GetLifecyclePolicy(cfg, client)
			return
		}
		if _ecrGetLifecyclePolicyPreview {
			ecr_GetLifecyclePolicyPreview(cfg, client)
			return
		}
		if _ecrGetRegistryPolicy {
			ecr_GetRegistryPolicy(cfg, client)
			return
		}
		if _ecrGetRegistryScanningConfiguration {
			ecr_GetRegistryScanningConfiguration(cfg, client)
			return
		}
		if _ecrGetRepositoryPolicy {
			ecr_GetRepositoryPolicy(cfg, client)
			return
		}
		if _ecrGetSigningConfiguration {
			ecr_GetSigningConfiguration(cfg, client)
			return
		}
		if _ecrInitiateLayerUpload {
			ecr_InitiateLayerUpload(cfg, client)
			return
		}
		if _ecrListImageReferrers {
			ecr_ListImageReferrers(cfg, client)
			return
		}
		if _ecrListImages {
			ecr_ListImages(cfg, client)
			return
		}
		if _ecrListPullTimeUpdateExclusions {
			ecr_ListPullTimeUpdateExclusions(cfg, client)
			return
		}
		if _ecrListTagsForResource {
			ecr_ListTagsForResource(cfg, client)
			return
		}
		if _ecrPutAccountSetting {
			ecr_PutAccountSetting(cfg, client)
			return
		}
		if _ecrPutImage {
			ecr_PutImage(cfg, client)
			return
		}
		if _ecrPutImageScanningConfiguration {
			ecr_PutImageScanningConfiguration(cfg, client)
			return
		}
		if _ecrPutImageTagMutability {
			ecr_PutImageTagMutability(cfg, client)
			return
		}
		if _ecrPutLifecyclePolicy {
			ecr_PutLifecyclePolicy(cfg, client)
			return
		}
		if _ecrPutRegistryPolicy {
			ecr_PutRegistryPolicy(cfg, client)
			return
		}
		if _ecrPutRegistryScanningConfiguration {
			ecr_PutRegistryScanningConfiguration(cfg, client)
			return
		}
		if _ecrPutReplicationConfiguration {
			ecr_PutReplicationConfiguration(cfg, client)
			return
		}
		if _ecrPutSigningConfiguration {
			ecr_PutSigningConfiguration(cfg, client)
			return
		}
		if _ecrRegisterPullTimeUpdateExclusion {
			ecr_RegisterPullTimeUpdateExclusion(cfg, client)
			return
		}
		if _ecrSetRepositoryPolicy {
			ecr_SetRepositoryPolicy(cfg, client)
			return
		}
		if _ecrStartImageScan {
			ecr_StartImageScan(cfg, client)
			return
		}
		if _ecrStartLifecyclePolicyPreview {
			ecr_StartLifecyclePolicyPreview(cfg, client)
			return
		}
		if _ecrTagResource {
			ecr_TagResource(cfg, client)
			return
		}
		if _ecrUntagResource {
			ecr_UntagResource(cfg, client)
			return
		}
		if _ecrUpdateImageStorageClass {
			ecr_UpdateImageStorageClass(cfg, client)
			return
		}
		if _ecrUpdatePullThroughCacheRule {
			ecr_UpdatePullThroughCacheRule(cfg, client)
			return
		}
		if _ecrUpdateRepositoryCreationTemplate {
			ecr_UpdateRepositoryCreationTemplate(cfg, client)
			return
		}
		if _ecrUploadLayerPart {
			ecr_UploadLayerPart(cfg, client)
			return
		}
		if _ecrValidatePullThroughCacheRule {
			ecr_ValidatePullThroughCacheRule(cfg, client)
			return
		}

	},
}

var (
	_ecrBatchCheckLayerAvailability             bool
	_ecrBatchDeleteImage                        bool
	_ecrBatchGetImage                           bool
	_ecrBatchGetRepositoryScanningConfiguration bool
	_ecrCompleteLayerUpload                     bool
	_ecrCreatePullThroughCacheRule              bool
	_ecrCreateRepository                        bool
	_ecrCreateRepositoryCreationTemplate        bool
	_ecrDeleteLifecyclePolicy                   bool
	_ecrDeletePullThroughCacheRule              bool
	_ecrDeleteRegistryPolicy                    bool
	_ecrDeleteRepository                        bool
	_ecrDeleteRepositoryCreationTemplate        bool
	_ecrDeleteRepositoryPolicy                  bool
	_ecrDeleteSigningConfiguration              bool
	_ecrDeregisterPullTimeUpdateExclusion       bool
	_ecrDescribeImageReplicationStatus          bool
	_ecrDescribeImageScanFindings               bool
	_ecrDescribeImageSigningStatus              bool
	_ecrDescribeImages                          bool
	_ecrDescribePullThroughCacheRules           bool
	_ecrDescribeRegistry                        bool
	_ecrDescribeRepositories                    bool
	_ecrDescribeRepositoryCreationTemplates     bool
	_ecrGetAccountSetting                       bool
	_ecrGetAuthorizationToken                   bool
	_ecrGetDownloadUrlForLayer                  bool
	_ecrGetLifecyclePolicy                      bool
	_ecrGetLifecyclePolicyPreview               bool
	_ecrGetRegistryPolicy                       bool
	_ecrGetRegistryScanningConfiguration        bool
	_ecrGetRepositoryPolicy                     bool
	_ecrGetSigningConfiguration                 bool
	_ecrInitiateLayerUpload                     bool
	_ecrListImageReferrers                      bool
	_ecrListImages                              bool
	_ecrListPullTimeUpdateExclusions            bool
	_ecrListTagsForResource                     bool
	_ecrPutAccountSetting                       bool
	_ecrPutImage                                bool
	_ecrPutImageScanningConfiguration           bool
	_ecrPutImageTagMutability                   bool
	_ecrPutLifecyclePolicy                      bool
	_ecrPutRegistryPolicy                       bool
	_ecrPutRegistryScanningConfiguration        bool
	_ecrPutReplicationConfiguration             bool
	_ecrPutSigningConfiguration                 bool
	_ecrRegisterPullTimeUpdateExclusion         bool
	_ecrSetRepositoryPolicy                     bool
	_ecrStartImageScan                          bool
	_ecrStartLifecyclePolicyPreview             bool
	_ecrTagResource                             bool
	_ecrUntagResource                           bool
	_ecrUpdateImageStorageClass                 bool
	_ecrUpdatePullThroughCacheRule              bool
	_ecrUpdateRepositoryCreationTemplate        bool
	_ecrUploadLayerPart                         bool
	_ecrValidatePullThroughCacheRule            bool

	_ecrAcceptedMediaTypes                 []string
	_ecrAppliedFor                         string
	_ecrCredentialArn                      string
	_ecrCustomRoleArn                      string
	_ecrDescription                        string
	_ecrEcrRepositoryPrefix                string
	_ecrEcrRepositoryPrefixes              []string
	_ecrEncryptionConfiguration            string
	_ecrFilter                             string
	_ecrForce                              string
	_ecrImageDigest                        string
	_ecrImageId                            string
	_ecrImageIds                           string
	_ecrImageManifest                      string
	_ecrImageManifestMediaType             string
	_ecrImageScanningConfiguration         string
	_ecrImageTag                           string
	_ecrImageTagMutability                 string
	_ecrImageTagMutabilityExclusionFilters string
	_ecrLayerDigest                        string
	_ecrLayerDigests                       []string
	_ecrLayerPartBlob                      string
	_ecrLifecyclePolicy                    string
	_ecrLifecyclePolicyText                string
	_ecrMaxResults                         string
	_ecrName                               string
	_ecrNextToken                          string
	_ecrPartFirstByte                      string
	_ecrPartLastByte                       string
	_ecrPolicyText                         string
	_ecrPrefix                             string
	_ecrPrefixes                           []string
	_ecrPrincipalArn                       string
	_ecrRegistryId                         string
	_ecrRegistryIds                        []string
	_ecrReplicationConfiguration           string
	_ecrRepositoryName                     string
	_ecrRepositoryNames                    []string
	_ecrRepositoryPolicy                   string
	_ecrResourceArn                        string
	_ecrResourceTags                       string
	_ecrRules                              string
	_ecrScanType                           string
	_ecrSigningConfiguration               string
	_ecrSubjectId                          string
	_ecrTagKeys                            []string
	_ecrTags                               string
	_ecrTargetStorageClass                 string
	_ecrUploadId                           string
	_ecrUpstreamRegistry                   string
	_ecrUpstreamRegistryUrl                string
	_ecrUpstreamRepositoryPrefix           string
	_ecrValue                              string
)

// Checks the availability of one or more image layers in a repository.
// When an image is pushed to a repository, each image layer is checked to verify
// if it has been uploaded before. If it has been uploaded, then the image layer is
// skipped.
//
// This operation is used by the Amazon ECR proxy and is not generally used by
// customers for pulling and pushing images. In most cases, you should use the
// docker CLI to pull, tag, and push images.
func ecr_BatchCheckLayerAvailability(cfg aws.Config, client *ecr.Client) {
	input := &ecr.BatchCheckLayerAvailabilityInput{
		// LayerDigests: []string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_ecrLayerDigests) > 0 {
		input.LayerDigests = append([]string(nil), _ecrLayerDigests...)
	}
	if len(_ecrRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrRepositoryName)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if resp, err := client.BatchCheckLayerAvailability(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a list of specified images within a repository. Images are specified
// with either an imageTag or imageDigest .
//
// You can remove a tag from an image by specifying the image's tag in your
// request. When you remove the last tag from an image, the image is deleted from
// your repository.
//
// You can completely delete an image (and all of its tags) by specifying the
// image's digest in your request.
func ecr_BatchDeleteImage(cfg aws.Config, client *ecr.Client) {
	input := &ecr.BatchDeleteImageInput{
		// ImageIds: []types.ImageIdentifier, // Required
		// RepositoryName: *string, // Required
	}

	if len(_ecrImageIds) > 0 {
		if err := assignInputField(input, "ImageIds", _ecrImageIds); err != nil {
			log.Errorf("invalid --image-ids: %s", err.Error())
			return
		}
	}
	if len(_ecrRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrRepositoryName)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if resp, err := client.BatchDeleteImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets detailed information for an image. Images are specified with either an
// imageTag or imageDigest .
//
// When an image is pulled, the BatchGetImage API is called once to retrieve the
// image manifest.
func ecr_BatchGetImage(cfg aws.Config, client *ecr.Client) {
	input := &ecr.BatchGetImageInput{
		// ImageIds: []types.ImageIdentifier, // Required
		// RepositoryName: *string, // Required
	}

	if len(_ecrImageIds) > 0 {
		if err := assignInputField(input, "ImageIds", _ecrImageIds); err != nil {
			log.Errorf("invalid --image-ids: %s", err.Error())
			return
		}
	}
	if len(_ecrRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrRepositoryName)
	}
	if len(_ecrAcceptedMediaTypes) > 0 {
		input.AcceptedMediaTypes = append([]string(nil), _ecrAcceptedMediaTypes...)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if resp, err := client.BatchGetImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the scanning configuration for one or more repositories.
func ecr_BatchGetRepositoryScanningConfiguration(cfg aws.Config, client *ecr.Client) {
	input := &ecr.BatchGetRepositoryScanningConfigurationInput{
		// RepositoryNames: []string, // Required
	}

	if len(_ecrRepositoryNames) > 0 {
		input.RepositoryNames = append([]string(nil), _ecrRepositoryNames...)
	}

	if resp, err := client.BatchGetRepositoryScanningConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Informs Amazon ECR that the image layer upload has completed for a specified
// registry, repository name, and upload ID. You can optionally provide a sha256
// digest of the image layer for data validation purposes.
//
// When an image is pushed, the CompleteLayerUpload API is called once per each
// new image layer to verify that the upload has completed.
//
// This operation is used by the Amazon ECR proxy and is not generally used by
// customers for pulling and pushing images. In most cases, you should use the
// docker CLI to pull, tag, and push images.
func ecr_CompleteLayerUpload(cfg aws.Config, client *ecr.Client) {
	input := &ecr.CompleteLayerUploadInput{
		// LayerDigests: []string, // Required
		// RepositoryName: *string, // Required
		// UploadId: *string, // Required
	}

	if len(_ecrLayerDigests) > 0 {
		input.LayerDigests = append([]string(nil), _ecrLayerDigests...)
	}
	if len(_ecrRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrRepositoryName)
	}
	if len(_ecrUploadId) > 0 {
		input.UploadId = aws.String(_ecrUploadId)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if resp, err := client.CompleteLayerUpload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a pull through cache rule. A pull through cache rule provides a way to
// cache images from an upstream registry source in your Amazon ECR private
// registry. For more information, see [Using pull through cache rules]in the Amazon Elastic Container Registry
// User Guide.
//
// [Using pull through cache rules]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/pull-through-cache.html
func ecr_CreatePullThroughCacheRule(cfg aws.Config, client *ecr.Client) {
	input := &ecr.CreatePullThroughCacheRuleInput{
		// EcrRepositoryPrefix: *string, // Required
		// UpstreamRegistryUrl: *string, // Required
	}

	if len(_ecrEcrRepositoryPrefix) > 0 {
		input.EcrRepositoryPrefix = aws.String(_ecrEcrRepositoryPrefix)
	}
	if len(_ecrUpstreamRegistryUrl) > 0 {
		input.UpstreamRegistryUrl = aws.String(_ecrUpstreamRegistryUrl)
	}
	if len(_ecrCredentialArn) > 0 {
		input.CredentialArn = aws.String(_ecrCredentialArn)
	}
	if len(_ecrCustomRoleArn) > 0 {
		input.CustomRoleArn = aws.String(_ecrCustomRoleArn)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}
	if len(_ecrUpstreamRegistry) > 0 {
		if err := assignInputField(input, "UpstreamRegistry", _ecrUpstreamRegistry); err != nil {
			log.Errorf("invalid --upstream-registry: %s", err.Error())
			return
		}
	}
	if len(_ecrUpstreamRepositoryPrefix) > 0 {
		input.UpstreamRepositoryPrefix = aws.String(_ecrUpstreamRepositoryPrefix)
	}

	if resp, err := client.CreatePullThroughCacheRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a repository. For more information, see [Amazon ECR repositories] in the Amazon Elastic
// Container Registry User Guide.
//
// [Amazon ECR repositories]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/Repositories.html
func ecr_CreateRepository(cfg aws.Config, client *ecr.Client) {
	input := &ecr.CreateRepositoryInput{
		// RepositoryName: *string, // Required
	}

	if len(_ecrRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrRepositoryName)
	}
	if len(_ecrEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _ecrEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_ecrImageScanningConfiguration) > 0 {
		if err := assignInputField(input, "ImageScanningConfiguration", _ecrImageScanningConfiguration); err != nil {
			log.Errorf("invalid --image-scanning-configuration: %s", err.Error())
			return
		}
	}
	if len(_ecrImageTagMutability) > 0 {
		if err := assignInputField(input, "ImageTagMutability", _ecrImageTagMutability); err != nil {
			log.Errorf("invalid --image-tag-mutability: %s", err.Error())
			return
		}
	}
	if len(_ecrImageTagMutabilityExclusionFilters) > 0 {
		if err := assignInputField(input, "ImageTagMutabilityExclusionFilters", _ecrImageTagMutabilityExclusionFilters); err != nil {
			log.Errorf("invalid --image-tag-mutability-exclusion-filters: %s", err.Error())
			return
		}
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}
	if len(_ecrTags) > 0 {
		if err := assignInputField(input, "Tags", _ecrTags); err != nil {
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

// Creates a repository creation template. This template is used to define the
// settings for repositories created by Amazon ECR on your behalf. For example,
// repositories created through pull through cache actions. For more information,
// see [Private repository creation templates]in the Amazon Elastic Container Registry User Guide.
//
// [Private repository creation templates]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/repository-creation-templates.html
func ecr_CreateRepositoryCreationTemplate(cfg aws.Config, client *ecr.Client) {
	input := &ecr.CreateRepositoryCreationTemplateInput{
		// AppliedFor: []types.RCTAppliedFor, // Required
		// Prefix: *string, // Required
	}

	if len(_ecrAppliedFor) > 0 {
		if err := assignInputField(input, "AppliedFor", _ecrAppliedFor); err != nil {
			log.Errorf("invalid --applied-for: %s", err.Error())
			return
		}
	}
	if len(_ecrPrefix) > 0 {
		input.Prefix = aws.String(_ecrPrefix)
	}
	if len(_ecrCustomRoleArn) > 0 {
		input.CustomRoleArn = aws.String(_ecrCustomRoleArn)
	}
	if len(_ecrDescription) > 0 {
		input.Description = aws.String(_ecrDescription)
	}
	if len(_ecrEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _ecrEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_ecrImageTagMutability) > 0 {
		if err := assignInputField(input, "ImageTagMutability", _ecrImageTagMutability); err != nil {
			log.Errorf("invalid --image-tag-mutability: %s", err.Error())
			return
		}
	}
	if len(_ecrImageTagMutabilityExclusionFilters) > 0 {
		if err := assignInputField(input, "ImageTagMutabilityExclusionFilters", _ecrImageTagMutabilityExclusionFilters); err != nil {
			log.Errorf("invalid --image-tag-mutability-exclusion-filters: %s", err.Error())
			return
		}
	}
	if len(_ecrLifecyclePolicy) > 0 {
		input.LifecyclePolicy = aws.String(_ecrLifecyclePolicy)
	}
	if len(_ecrRepositoryPolicy) > 0 {
		input.RepositoryPolicy = aws.String(_ecrRepositoryPolicy)
	}
	if len(_ecrResourceTags) > 0 {
		if err := assignInputField(input, "ResourceTags", _ecrResourceTags); err != nil {
			log.Errorf("invalid --resource-tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRepositoryCreationTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the lifecycle policy associated with the specified repository.
func ecr_DeleteLifecyclePolicy(cfg aws.Config, client *ecr.Client) {
	input := &ecr.DeleteLifecyclePolicyInput{
		// RepositoryName: *string, // Required
	}

	if len(_ecrRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrRepositoryName)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if resp, err := client.DeleteLifecyclePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a pull through cache rule.
func ecr_DeletePullThroughCacheRule(cfg aws.Config, client *ecr.Client) {
	input := &ecr.DeletePullThroughCacheRuleInput{
		// EcrRepositoryPrefix: *string, // Required
	}

	if len(_ecrEcrRepositoryPrefix) > 0 {
		input.EcrRepositoryPrefix = aws.String(_ecrEcrRepositoryPrefix)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if resp, err := client.DeletePullThroughCacheRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the registry permissions policy.
func ecr_DeleteRegistryPolicy(cfg aws.Config, client *ecr.Client) {
	input := &ecr.DeleteRegistryPolicyInput{}

	if resp, err := client.DeleteRegistryPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a repository. If the repository isn't empty, you must either delete the
// contents of the repository or use the force option to delete the repository and
// have Amazon ECR delete all of its contents on your behalf.
func ecr_DeleteRepository(cfg aws.Config, client *ecr.Client) {
	input := &ecr.DeleteRepositoryInput{
		// RepositoryName: *string, // Required
	}

	if len(_ecrRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrRepositoryName)
	}
	if len(_ecrForce) > 0 {
		if err := assignInputField(input, "Force", _ecrForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if resp, err := client.DeleteRepository(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a repository creation template.
func ecr_DeleteRepositoryCreationTemplate(cfg aws.Config, client *ecr.Client) {
	input := &ecr.DeleteRepositoryCreationTemplateInput{
		// Prefix: *string, // Required
	}

	if len(_ecrPrefix) > 0 {
		input.Prefix = aws.String(_ecrPrefix)
	}

	if resp, err := client.DeleteRepositoryCreationTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the repository policy associated with the specified repository.
func ecr_DeleteRepositoryPolicy(cfg aws.Config, client *ecr.Client) {
	input := &ecr.DeleteRepositoryPolicyInput{
		// RepositoryName: *string, // Required
	}

	if len(_ecrRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrRepositoryName)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if resp, err := client.DeleteRepositoryPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the registry's signing configuration. Images pushed after deletion of
// the signing configuration will no longer be automatically signed.
//
// For more information, see [Managed signing] in the Amazon Elastic Container Registry User Guide.
//
// Deleting the signing configuration does not affect existing image signatures.
//
// [Managed signing]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/managed-signing.html
func ecr_DeleteSigningConfiguration(cfg aws.Config, client *ecr.Client) {
	input := &ecr.DeleteSigningConfigurationInput{}

	if resp, err := client.DeleteSigningConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a principal from the pull time update exclusion list for a registry.
// Once removed, Amazon ECR will resume updating the pull time if the specified
// principal pulls an image.
func ecr_DeregisterPullTimeUpdateExclusion(cfg aws.Config, client *ecr.Client) {
	input := &ecr.DeregisterPullTimeUpdateExclusionInput{
		// PrincipalArn: *string, // Required
	}

	if len(_ecrPrincipalArn) > 0 {
		input.PrincipalArn = aws.String(_ecrPrincipalArn)
	}

	if resp, err := client.DeregisterPullTimeUpdateExclusion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the replication status for a specified image.
func ecr_DescribeImageReplicationStatus(cfg aws.Config, client *ecr.Client) {
	input := &ecr.DescribeImageReplicationStatusInput{
		// ImageId: *types.ImageIdentifier, // Required
		// RepositoryName: *string, // Required
	}

	if len(_ecrImageId) > 0 {
		if err := assignInputField(input, "ImageId", _ecrImageId); err != nil {
			log.Errorf("invalid --image-id: %s", err.Error())
			return
		}
	}
	if len(_ecrRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrRepositoryName)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if resp, err := client.DescribeImageReplicationStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the scan findings for the specified image.
func ecr_DescribeImageScanFindings(cfg aws.Config, client *ecr.Client) {
	input := &ecr.DescribeImageScanFindingsInput{
		// ImageId: *types.ImageIdentifier, // Required
		// RepositoryName: *string, // Required
	}

	if len(_ecrImageId) > 0 {
		if err := assignInputField(input, "ImageId", _ecrImageId); err != nil {
			log.Errorf("invalid --image-id: %s", err.Error())
			return
		}
	}
	if len(_ecrRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrRepositoryName)
	}
	if len(_ecrMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ecrMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ecrNextToken) > 0 {
		input.NextToken = aws.String(_ecrNextToken)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if disablePaginator() {
		if resp, err := client.DescribeImageScanFindings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ecr.DescribeImageScanFindingsOutput
	p := ecr.NewDescribeImageScanFindingsPaginator(client, input)
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

// Returns the signing status for a specified image. If the image matched signing
// rules that reference different signing profiles, a status is returned for each
// profile.
//
// For more information, see [Managed signing] in the Amazon Elastic Container Registry User Guide.
//
// [Managed signing]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/managed-signing.html
func ecr_DescribeImageSigningStatus(cfg aws.Config, client *ecr.Client) {
	input := &ecr.DescribeImageSigningStatusInput{
		// ImageId: *types.ImageIdentifier, // Required
		// RepositoryName: *string, // Required
	}

	if len(_ecrImageId) > 0 {
		if err := assignInputField(input, "ImageId", _ecrImageId); err != nil {
			log.Errorf("invalid --image-id: %s", err.Error())
			return
		}
	}
	if len(_ecrRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrRepositoryName)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if resp, err := client.DescribeImageSigningStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns metadata about the images in a repository.
// Starting with Docker version 1.9, the Docker client compresses image layers
// before pushing them to a V2 Docker registry. The output of the docker images
// command shows the uncompressed image size. Therefore, Docker might return a
// larger image than the image shown in the Amazon Web Services Management Console.
//
// The new version of Amazon ECR Basic Scanning doesn't use the ImageDetail$imageScanFindingsSummary and ImageDetail$imageScanStatus attributes
// from the API response to return scan results. Use the DescribeImageScanFindingsAPI instead. For more
// information about Amazon Web Services native basic scanning, see [Scan images for software vulnerabilities in Amazon ECR].
//
// [Scan images for software vulnerabilities in Amazon ECR]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning.html
func ecr_DescribeImages(cfg aws.Config, client *ecr.Client) {
	input := &ecr.DescribeImagesInput{
		// RepositoryName: *string, // Required
	}

	if len(_ecrRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrRepositoryName)
	}
	if len(_ecrFilter) > 0 {
		if err := assignInputField(input, "Filter", _ecrFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_ecrImageIds) > 0 {
		if err := assignInputField(input, "ImageIds", _ecrImageIds); err != nil {
			log.Errorf("invalid --image-ids: %s", err.Error())
			return
		}
	}
	if len(_ecrMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ecrMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ecrNextToken) > 0 {
		input.NextToken = aws.String(_ecrNextToken)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if disablePaginator() {
		if resp, err := client.DescribeImages(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ecr.DescribeImagesOutput
	p := ecr.NewDescribeImagesPaginator(client, input)
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

// Returns the pull through cache rules for a registry.
func ecr_DescribePullThroughCacheRules(cfg aws.Config, client *ecr.Client) {
	input := &ecr.DescribePullThroughCacheRulesInput{}

	if len(_ecrEcrRepositoryPrefixes) > 0 {
		input.EcrRepositoryPrefixes = append([]string(nil), _ecrEcrRepositoryPrefixes...)
	}
	if len(_ecrMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ecrMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ecrNextToken) > 0 {
		input.NextToken = aws.String(_ecrNextToken)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if disablePaginator() {
		if resp, err := client.DescribePullThroughCacheRules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ecr.DescribePullThroughCacheRulesOutput
	p := ecr.NewDescribePullThroughCacheRulesPaginator(client, input)
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

// Describes the settings for a registry. The replication configuration for a
// repository can be created or updated with the PutReplicationConfigurationAPI action.
func ecr_DescribeRegistry(cfg aws.Config, client *ecr.Client) {
	input := &ecr.DescribeRegistryInput{}

	if resp, err := client.DescribeRegistry(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes image repositories in a registry.
func ecr_DescribeRepositories(cfg aws.Config, client *ecr.Client) {
	input := &ecr.DescribeRepositoriesInput{}

	if len(_ecrMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ecrMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ecrNextToken) > 0 {
		input.NextToken = aws.String(_ecrNextToken)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}
	if len(_ecrRepositoryNames) > 0 {
		input.RepositoryNames = append([]string(nil), _ecrRepositoryNames...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeRepositories(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ecr.DescribeRepositoriesOutput
	p := ecr.NewDescribeRepositoriesPaginator(client, input)
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

// Returns details about the repository creation templates in a registry. The
// prefixes request parameter can be used to return the details for a specific
// repository creation template.
func ecr_DescribeRepositoryCreationTemplates(cfg aws.Config, client *ecr.Client) {
	input := &ecr.DescribeRepositoryCreationTemplatesInput{}

	if len(_ecrMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ecrMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ecrNextToken) > 0 {
		input.NextToken = aws.String(_ecrNextToken)
	}
	if len(_ecrPrefixes) > 0 {
		input.Prefixes = append([]string(nil), _ecrPrefixes...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeRepositoryCreationTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ecr.DescribeRepositoryCreationTemplatesOutput
	p := ecr.NewDescribeRepositoryCreationTemplatesPaginator(client, input)
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

// Retrieves the account setting value for the specified setting name.
func ecr_GetAccountSetting(cfg aws.Config, client *ecr.Client) {
	input := &ecr.GetAccountSettingInput{
		// Name: *string, // Required
	}

	if len(_ecrName) > 0 {
		input.Name = aws.String(_ecrName)
	}

	if resp, err := client.GetAccountSetting(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves an authorization token. An authorization token represents your IAM
// authentication credentials and can be used to access any Amazon ECR registry
// that your IAM principal has access to. The authorization token is valid for 12
// hours.
//
// The authorizationToken returned is a base64 encoded string that can be decoded
// and used in a docker login command to authenticate to a registry. The CLI
// offers an get-login-password command that simplifies the login process. For
// more information, see [Registry authentication]in the Amazon Elastic Container Registry User Guide.
//
// [Registry authentication]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/Registries.html#registry_auth
func ecr_GetAuthorizationToken(cfg aws.Config, client *ecr.Client) {
	input := &ecr.GetAuthorizationTokenInput{}

	if len(_ecrRegistryIds) > 0 {
		input.RegistryIds = append([]string(nil), _ecrRegistryIds...)
	}

	if resp, err := client.GetAuthorizationToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the pre-signed Amazon S3 download URL corresponding to an image
// layer. You can only get URLs for image layers that are referenced in an image.
//
// When an image is pulled, the GetDownloadUrlForLayer API is called once per
// image layer that is not already cached.
//
// This operation is used by the Amazon ECR proxy and is not generally used by
// customers for pulling and pushing images. In most cases, you should use the
// docker CLI to pull, tag, and push images.
func ecr_GetDownloadUrlForLayer(cfg aws.Config, client *ecr.Client) {
	input := &ecr.GetDownloadUrlForLayerInput{
		// LayerDigest: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_ecrLayerDigest) > 0 {
		input.LayerDigest = aws.String(_ecrLayerDigest)
	}
	if len(_ecrRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrRepositoryName)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if resp, err := client.GetDownloadUrlForLayer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the lifecycle policy for the specified repository.
func ecr_GetLifecyclePolicy(cfg aws.Config, client *ecr.Client) {
	input := &ecr.GetLifecyclePolicyInput{
		// RepositoryName: *string, // Required
	}

	if len(_ecrRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrRepositoryName)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if resp, err := client.GetLifecyclePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the results of the lifecycle policy preview request for the specified
// repository.
func ecr_GetLifecyclePolicyPreview(cfg aws.Config, client *ecr.Client) {
	input := &ecr.GetLifecyclePolicyPreviewInput{
		// RepositoryName: *string, // Required
	}

	if len(_ecrRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrRepositoryName)
	}
	if len(_ecrFilter) > 0 {
		if err := assignInputField(input, "Filter", _ecrFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_ecrImageIds) > 0 {
		if err := assignInputField(input, "ImageIds", _ecrImageIds); err != nil {
			log.Errorf("invalid --image-ids: %s", err.Error())
			return
		}
	}
	if len(_ecrMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ecrMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ecrNextToken) > 0 {
		input.NextToken = aws.String(_ecrNextToken)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if disablePaginator() {
		if resp, err := client.GetLifecyclePolicyPreview(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ecr.GetLifecyclePolicyPreviewOutput
	p := ecr.NewGetLifecyclePolicyPreviewPaginator(client, input)
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

// Retrieves the permissions policy for a registry.
func ecr_GetRegistryPolicy(cfg aws.Config, client *ecr.Client) {
	input := &ecr.GetRegistryPolicyInput{}

	if resp, err := client.GetRegistryPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the scanning configuration for a registry.
func ecr_GetRegistryScanningConfiguration(cfg aws.Config, client *ecr.Client) {
	input := &ecr.GetRegistryScanningConfigurationInput{}

	if resp, err := client.GetRegistryScanningConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the repository policy for the specified repository.
func ecr_GetRepositoryPolicy(cfg aws.Config, client *ecr.Client) {
	input := &ecr.GetRepositoryPolicyInput{
		// RepositoryName: *string, // Required
	}

	if len(_ecrRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrRepositoryName)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if resp, err := client.GetRepositoryPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the registry's signing configuration, which defines rules for
// automatically signing images using Amazon Web Services Signer.
//
// For more information, see [Managed signing] in the Amazon Elastic Container Registry User Guide.
//
// [Managed signing]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/managed-signing.html
func ecr_GetSigningConfiguration(cfg aws.Config, client *ecr.Client) {
	input := &ecr.GetSigningConfigurationInput{}

	if resp, err := client.GetSigningConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Notifies Amazon ECR that you intend to upload an image layer.
// When an image is pushed, the InitiateLayerUpload API is called once per image
// layer that has not already been uploaded. Whether or not an image layer has been
// uploaded is determined by the BatchCheckLayerAvailability API action.
//
// This operation is used by the Amazon ECR proxy and is not generally used by
// customers for pulling and pushing images. In most cases, you should use the
// docker CLI to pull, tag, and push images.
func ecr_InitiateLayerUpload(cfg aws.Config, client *ecr.Client) {
	input := &ecr.InitiateLayerUploadInput{
		// RepositoryName: *string, // Required
	}

	if len(_ecrRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrRepositoryName)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if resp, err := client.InitiateLayerUpload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the artifacts associated with a specified subject image.
// The IAM principal invoking this operation must have the ecr:BatchGetImage
// permission.
func ecr_ListImageReferrers(cfg aws.Config, client *ecr.Client) {
	input := &ecr.ListImageReferrersInput{
		// RepositoryName: *string, // Required
		// SubjectId: *types.SubjectIdentifier, // Required
	}

	if len(_ecrRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrRepositoryName)
	}
	if len(_ecrSubjectId) > 0 {
		if err := assignInputField(input, "SubjectId", _ecrSubjectId); err != nil {
			log.Errorf("invalid --subject-id: %s", err.Error())
			return
		}
	}
	if len(_ecrFilter) > 0 {
		if err := assignInputField(input, "Filter", _ecrFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_ecrMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ecrMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ecrNextToken) > 0 {
		input.NextToken = aws.String(_ecrNextToken)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if resp, err := client.ListImageReferrers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the image IDs for the specified repository.
// You can filter images based on whether or not they are tagged by using the
// tagStatus filter and specifying either TAGGED , UNTAGGED or ANY . For example,
// you can filter your results to return only UNTAGGED images and then pipe that
// result to a BatchDeleteImageoperation to delete them. Or, you can filter your results to return
// only TAGGED images to list all of the tags in your repository.
func ecr_ListImages(cfg aws.Config, client *ecr.Client) {
	input := &ecr.ListImagesInput{
		// RepositoryName: *string, // Required
	}

	if len(_ecrRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrRepositoryName)
	}
	if len(_ecrFilter) > 0 {
		if err := assignInputField(input, "Filter", _ecrFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_ecrMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ecrMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ecrNextToken) > 0 {
		input.NextToken = aws.String(_ecrNextToken)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
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

	var results []*ecr.ListImagesOutput
	p := ecr.NewListImagesPaginator(client, input)
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

// Lists the IAM principals that are excluded from having their image pull times
// recorded.
func ecr_ListPullTimeUpdateExclusions(cfg aws.Config, client *ecr.Client) {
	input := &ecr.ListPullTimeUpdateExclusionsInput{}

	if len(_ecrMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ecrMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ecrNextToken) > 0 {
		input.NextToken = aws.String(_ecrNextToken)
	}

	if resp, err := client.ListPullTimeUpdateExclusions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List the tags for an Amazon ECR resource.
func ecr_ListTagsForResource(cfg aws.Config, client *ecr.Client) {
	input := &ecr.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_ecrResourceArn) > 0 {
		input.ResourceArn = aws.String(_ecrResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows you to change the basic scan type version or registry policy scope.
func ecr_PutAccountSetting(cfg aws.Config, client *ecr.Client) {
	input := &ecr.PutAccountSettingInput{
		// Name: *string, // Required
		// Value: *string, // Required
	}

	if len(_ecrName) > 0 {
		input.Name = aws.String(_ecrName)
	}
	if len(_ecrValue) > 0 {
		input.Value = aws.String(_ecrValue)
	}

	if resp, err := client.PutAccountSetting(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates the image manifest and tags associated with an image.
// When an image is pushed and all new image layers have been uploaded, the
// PutImage API is called once to create or update the image manifest and the tags
// associated with the image.
//
// This operation is used by the Amazon ECR proxy and is not generally used by
// customers for pulling and pushing images. In most cases, you should use the
// docker CLI to pull, tag, and push images.
func ecr_PutImage(cfg aws.Config, client *ecr.Client) {
	input := &ecr.PutImageInput{
		// ImageManifest: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_ecrImageManifest) > 0 {
		input.ImageManifest = aws.String(_ecrImageManifest)
	}
	if len(_ecrRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrRepositoryName)
	}
	if len(_ecrImageDigest) > 0 {
		input.ImageDigest = aws.String(_ecrImageDigest)
	}
	if len(_ecrImageManifestMediaType) > 0 {
		input.ImageManifestMediaType = aws.String(_ecrImageManifestMediaType)
	}
	if len(_ecrImageTag) > 0 {
		input.ImageTag = aws.String(_ecrImageTag)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if resp, err := client.PutImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The PutImageScanningConfiguration API is being deprecated, in favor of
// specifying the image scanning configuration at the registry level. For more
// information, see PutRegistryScanningConfiguration.
//
// Updates the image scanning configuration for the specified repository.
func ecr_PutImageScanningConfiguration(cfg aws.Config, client *ecr.Client) {
	input := &ecr.PutImageScanningConfigurationInput{
		// ImageScanningConfiguration: *types.ImageScanningConfiguration, // Required
		// RepositoryName: *string, // Required
	}

	if len(_ecrImageScanningConfiguration) > 0 {
		if err := assignInputField(input, "ImageScanningConfiguration", _ecrImageScanningConfiguration); err != nil {
			log.Errorf("invalid --image-scanning-configuration: %s", err.Error())
			return
		}
	}
	if len(_ecrRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrRepositoryName)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if resp, err := client.PutImageScanningConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the image tag mutability settings for the specified repository. For
// more information, see [Image tag mutability]in the Amazon Elastic Container Registry User Guide.
//
// [Image tag mutability]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-tag-mutability.html
func ecr_PutImageTagMutability(cfg aws.Config, client *ecr.Client) {
	input := &ecr.PutImageTagMutabilityInput{
		// ImageTagMutability: types.ImageTagMutability, // Required
		// RepositoryName: *string, // Required
	}

	if len(_ecrImageTagMutability) > 0 {
		if err := assignInputField(input, "ImageTagMutability", _ecrImageTagMutability); err != nil {
			log.Errorf("invalid --image-tag-mutability: %s", err.Error())
			return
		}
	}
	if len(_ecrRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrRepositoryName)
	}
	if len(_ecrImageTagMutabilityExclusionFilters) > 0 {
		if err := assignInputField(input, "ImageTagMutabilityExclusionFilters", _ecrImageTagMutabilityExclusionFilters); err != nil {
			log.Errorf("invalid --image-tag-mutability-exclusion-filters: %s", err.Error())
			return
		}
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if resp, err := client.PutImageTagMutability(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates the lifecycle policy for the specified repository. For more
// information, see [Lifecycle policy template].
//
// [Lifecycle policy template]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html
func ecr_PutLifecyclePolicy(cfg aws.Config, client *ecr.Client) {
	input := &ecr.PutLifecyclePolicyInput{
		// LifecyclePolicyText: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_ecrLifecyclePolicyText) > 0 {
		input.LifecyclePolicyText = aws.String(_ecrLifecyclePolicyText)
	}
	if len(_ecrRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrRepositoryName)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if resp, err := client.PutLifecyclePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates the permissions policy for your registry.
// A registry policy is used to specify permissions for another Amazon Web
// Services account and is used when configuring cross-account replication. For
// more information, see [Registry permissions]in the Amazon Elastic Container Registry User Guide.
//
// [Registry permissions]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/registry-permissions.html
func ecr_PutRegistryPolicy(cfg aws.Config, client *ecr.Client) {
	input := &ecr.PutRegistryPolicyInput{
		// PolicyText: *string, // Required
	}

	if len(_ecrPolicyText) > 0 {
		input.PolicyText = aws.String(_ecrPolicyText)
	}

	if resp, err := client.PutRegistryPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates the scanning configuration for your private registry.
func ecr_PutRegistryScanningConfiguration(cfg aws.Config, client *ecr.Client) {
	input := &ecr.PutRegistryScanningConfigurationInput{}

	if len(_ecrRules) > 0 {
		if err := assignInputField(input, "Rules", _ecrRules); err != nil {
			log.Errorf("invalid --rules: %s", err.Error())
			return
		}
	}
	if len(_ecrScanType) > 0 {
		if err := assignInputField(input, "ScanType", _ecrScanType); err != nil {
			log.Errorf("invalid --scan-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutRegistryScanningConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates the replication configuration for a registry. The existing
// replication configuration for a repository can be retrieved with the DescribeRegistryAPI
// action. The first time the PutReplicationConfiguration API is called, a
// service-linked IAM role is created in your account for the replication process.
// For more information, see [Using service-linked roles for Amazon ECR]in the Amazon Elastic Container Registry User Guide.
// For more information on the custom role for replication, see [Creating an IAM role for replication].
//
// When configuring cross-account replication, the destination account must grant
// the source account permission to replicate. This permission is controlled using
// a registry permissions policy. For more information, see PutRegistryPolicy.
//
// [Creating an IAM role for replication]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/replication-creation-templates.html#roles-creatingrole-user-console
// [Using service-linked roles for Amazon ECR]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/using-service-linked-roles.html
func ecr_PutReplicationConfiguration(cfg aws.Config, client *ecr.Client) {
	input := &ecr.PutReplicationConfigurationInput{
		// ReplicationConfiguration: *types.ReplicationConfiguration, // Required
	}

	if len(_ecrReplicationConfiguration) > 0 {
		if err := assignInputField(input, "ReplicationConfiguration", _ecrReplicationConfiguration); err != nil {
			log.Errorf("invalid --replication-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutReplicationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates the registry's signing configuration, which defines rules
// for automatically signing images with Amazon Web Services Signer.
//
// For more information, see [Managed signing] in the Amazon Elastic Container Registry User Guide.
//
// To successfully generate a signature, the IAM principal pushing images must
// have permission to sign payloads with the Amazon Web Services Signer signing
// profile referenced in the signing configuration.
//
// [Managed signing]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/managed-signing.html
func ecr_PutSigningConfiguration(cfg aws.Config, client *ecr.Client) {
	input := &ecr.PutSigningConfigurationInput{
		// SigningConfiguration: *types.SigningConfiguration, // Required
	}

	if len(_ecrSigningConfiguration) > 0 {
		if err := assignInputField(input, "SigningConfiguration", _ecrSigningConfiguration); err != nil {
			log.Errorf("invalid --signing-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutSigningConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds an IAM principal to the pull time update exclusion list for a registry.
// Amazon ECR will not record the pull time if an excluded principal pulls an
// image.
func ecr_RegisterPullTimeUpdateExclusion(cfg aws.Config, client *ecr.Client) {
	input := &ecr.RegisterPullTimeUpdateExclusionInput{
		// PrincipalArn: *string, // Required
	}

	if len(_ecrPrincipalArn) > 0 {
		input.PrincipalArn = aws.String(_ecrPrincipalArn)
	}

	if resp, err := client.RegisterPullTimeUpdateExclusion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies a repository policy to the specified repository to control access
// permissions. For more information, see [Amazon ECR Repository policies]in the Amazon Elastic Container Registry
// User Guide.
//
// [Amazon ECR Repository policies]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/repository-policies.html
func ecr_SetRepositoryPolicy(cfg aws.Config, client *ecr.Client) {
	input := &ecr.SetRepositoryPolicyInput{
		// PolicyText: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_ecrPolicyText) > 0 {
		input.PolicyText = aws.String(_ecrPolicyText)
	}
	if len(_ecrRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrRepositoryName)
	}
	if len(_ecrForce) > 0 {
		if err := assignInputField(input, "Force", _ecrForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if resp, err := client.SetRepositoryPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a basic image vulnerability scan.
// A basic image scan can only be started once per 24 hours on an individual
// image. This limit includes if an image was scanned on initial push. You can
// start up to 100,000 basic scans per 24 hours. This limit includes both scans on
// initial push and scans initiated by the StartImageScan API. For more
// information, see [Basic scanning]in the Amazon Elastic Container Registry User Guide.
//
// [Basic scanning]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning-basic.html
func ecr_StartImageScan(cfg aws.Config, client *ecr.Client) {
	input := &ecr.StartImageScanInput{
		// ImageId: *types.ImageIdentifier, // Required
		// RepositoryName: *string, // Required
	}

	if len(_ecrImageId) > 0 {
		if err := assignInputField(input, "ImageId", _ecrImageId); err != nil {
			log.Errorf("invalid --image-id: %s", err.Error())
			return
		}
	}
	if len(_ecrRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrRepositoryName)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if resp, err := client.StartImageScan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a preview of a lifecycle policy for the specified repository. This
// allows you to see the results before associating the lifecycle policy with the
// repository.
func ecr_StartLifecyclePolicyPreview(cfg aws.Config, client *ecr.Client) {
	input := &ecr.StartLifecyclePolicyPreviewInput{
		// RepositoryName: *string, // Required
	}

	if len(_ecrRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrRepositoryName)
	}
	if len(_ecrLifecyclePolicyText) > 0 {
		input.LifecyclePolicyText = aws.String(_ecrLifecyclePolicyText)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if resp, err := client.StartLifecyclePolicyPreview(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds specified tags to a resource with the specified ARN. Existing tags on a
// resource are not changed if they are not specified in the request parameters.
func ecr_TagResource(cfg aws.Config, client *ecr.Client) {
	input := &ecr.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_ecrResourceArn) > 0 {
		input.ResourceArn = aws.String(_ecrResourceArn)
	}
	if len(_ecrTags) > 0 {
		if err := assignInputField(input, "Tags", _ecrTags); err != nil {
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

// Deletes specified tags from a resource.
func ecr_UntagResource(cfg aws.Config, client *ecr.Client) {
	input := &ecr.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_ecrResourceArn) > 0 {
		input.ResourceArn = aws.String(_ecrResourceArn)
	}
	if len(_ecrTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _ecrTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Transitions an image between storage classes. You can transition images from
// Amazon ECR standard storage class to Amazon ECR archival storage class for
// long-term storage, or restore archived images back to Amazon ECR standard.
func ecr_UpdateImageStorageClass(cfg aws.Config, client *ecr.Client) {
	input := &ecr.UpdateImageStorageClassInput{
		// ImageId: *types.ImageIdentifier, // Required
		// RepositoryName: *string, // Required
		// TargetStorageClass: types.TargetStorageClass, // Required
	}

	if len(_ecrImageId) > 0 {
		if err := assignInputField(input, "ImageId", _ecrImageId); err != nil {
			log.Errorf("invalid --image-id: %s", err.Error())
			return
		}
	}
	if len(_ecrRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrRepositoryName)
	}
	if len(_ecrTargetStorageClass) > 0 {
		if err := assignInputField(input, "TargetStorageClass", _ecrTargetStorageClass); err != nil {
			log.Errorf("invalid --target-storage-class: %s", err.Error())
			return
		}
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if resp, err := client.UpdateImageStorageClass(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing pull through cache rule.
func ecr_UpdatePullThroughCacheRule(cfg aws.Config, client *ecr.Client) {
	input := &ecr.UpdatePullThroughCacheRuleInput{
		// EcrRepositoryPrefix: *string, // Required
	}

	if len(_ecrEcrRepositoryPrefix) > 0 {
		input.EcrRepositoryPrefix = aws.String(_ecrEcrRepositoryPrefix)
	}
	if len(_ecrCredentialArn) > 0 {
		input.CredentialArn = aws.String(_ecrCredentialArn)
	}
	if len(_ecrCustomRoleArn) > 0 {
		input.CustomRoleArn = aws.String(_ecrCustomRoleArn)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if resp, err := client.UpdatePullThroughCacheRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing repository creation template.
func ecr_UpdateRepositoryCreationTemplate(cfg aws.Config, client *ecr.Client) {
	input := &ecr.UpdateRepositoryCreationTemplateInput{
		// Prefix: *string, // Required
	}

	if len(_ecrPrefix) > 0 {
		input.Prefix = aws.String(_ecrPrefix)
	}
	if len(_ecrAppliedFor) > 0 {
		if err := assignInputField(input, "AppliedFor", _ecrAppliedFor); err != nil {
			log.Errorf("invalid --applied-for: %s", err.Error())
			return
		}
	}
	if len(_ecrCustomRoleArn) > 0 {
		input.CustomRoleArn = aws.String(_ecrCustomRoleArn)
	}
	if len(_ecrDescription) > 0 {
		input.Description = aws.String(_ecrDescription)
	}
	if len(_ecrEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _ecrEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_ecrImageTagMutability) > 0 {
		if err := assignInputField(input, "ImageTagMutability", _ecrImageTagMutability); err != nil {
			log.Errorf("invalid --image-tag-mutability: %s", err.Error())
			return
		}
	}
	if len(_ecrImageTagMutabilityExclusionFilters) > 0 {
		if err := assignInputField(input, "ImageTagMutabilityExclusionFilters", _ecrImageTagMutabilityExclusionFilters); err != nil {
			log.Errorf("invalid --image-tag-mutability-exclusion-filters: %s", err.Error())
			return
		}
	}
	if len(_ecrLifecyclePolicy) > 0 {
		input.LifecyclePolicy = aws.String(_ecrLifecyclePolicy)
	}
	if len(_ecrRepositoryPolicy) > 0 {
		input.RepositoryPolicy = aws.String(_ecrRepositoryPolicy)
	}
	if len(_ecrResourceTags) > 0 {
		if err := assignInputField(input, "ResourceTags", _ecrResourceTags); err != nil {
			log.Errorf("invalid --resource-tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRepositoryCreationTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Uploads an image layer part to Amazon ECR.
// When an image is pushed, each new image layer is uploaded in parts. The maximum
// size of each image layer part can be 20971520 bytes (or about 20MB). The
// UploadLayerPart API is called once per each new image layer part.
//
// This operation is used by the Amazon ECR proxy and is not generally used by
// customers for pulling and pushing images. In most cases, you should use the
// docker CLI to pull, tag, and push images.
func ecr_UploadLayerPart(cfg aws.Config, client *ecr.Client) {
	input := &ecr.UploadLayerPartInput{
		// LayerPartBlob: []byte, // Required
		// PartFirstByte: *int64, // Required
		// PartLastByte: *int64, // Required
		// RepositoryName: *string, // Required
		// UploadId: *string, // Required
	}

	if len(_ecrLayerPartBlob) > 0 {
		if err := assignInputField(input, "LayerPartBlob", _ecrLayerPartBlob); err != nil {
			log.Errorf("invalid --layer-part-blob: %s", err.Error())
			return
		}
	}
	if len(_ecrPartFirstByte) > 0 {
		if err := assignInputField(input, "PartFirstByte", _ecrPartFirstByte); err != nil {
			log.Errorf("invalid --part-first-byte: %s", err.Error())
			return
		}
	}
	if len(_ecrPartLastByte) > 0 {
		if err := assignInputField(input, "PartLastByte", _ecrPartLastByte); err != nil {
			log.Errorf("invalid --part-last-byte: %s", err.Error())
			return
		}
	}
	if len(_ecrRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrRepositoryName)
	}
	if len(_ecrUploadId) > 0 {
		input.UploadId = aws.String(_ecrUploadId)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if resp, err := client.UploadLayerPart(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Validates an existing pull through cache rule for an upstream registry that
// requires authentication. This will retrieve the contents of the Amazon Web
// Services Secrets Manager secret, verify the syntax, and then validate that
// authentication to the upstream registry is successful.
func ecr_ValidatePullThroughCacheRule(cfg aws.Config, client *ecr.Client) {
	input := &ecr.ValidatePullThroughCacheRuleInput{
		// EcrRepositoryPrefix: *string, // Required
	}

	if len(_ecrEcrRepositoryPrefix) > 0 {
		input.EcrRepositoryPrefix = aws.String(_ecrEcrRepositoryPrefix)
	}
	if len(_ecrRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrRegistryId)
	}

	if resp, err := client.ValidatePullThroughCacheRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_ecrCmd)
	_ecrCmd.Flags().SortFlags = false

	_ecrCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_ecrCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_ecrCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_ecrCmd.Flags().StringSliceVarP(&_ecrAcceptedMediaTypes, "accepted-media-types", "", nil, "Accepted Media Types")
	_ecrCmd.Flags().StringVarP(&_ecrAppliedFor, "applied-for", "", "", "Applied For")
	_ecrCmd.Flags().StringVarP(&_ecrCredentialArn, "credential-arn", "", "", "Credential ARN")
	_ecrCmd.Flags().StringVarP(&_ecrCustomRoleArn, "custom-role-arn", "", "", "Custom Role ARN")
	_ecrCmd.Flags().StringVarP(&_ecrDescription, "description", "", "", "Description")
	_ecrCmd.Flags().StringVarP(&_ecrEcrRepositoryPrefix, "ecr-repository-prefix", "", "", "ECR Repository Prefix")
	_ecrCmd.Flags().StringSliceVarP(&_ecrEcrRepositoryPrefixes, "ecr-repository-prefixes", "", nil, "ECR Repository Prefixes")
	_ecrCmd.Flags().StringVarP(&_ecrEncryptionConfiguration, "encryption-configuration", "", "", "Encryption Configuration")
	_ecrCmd.Flags().StringVarP(&_ecrFilter, "filter", "", "", "Filter")
	_ecrCmd.Flags().StringVarP(&_ecrForce, "force", "", "", "Force")
	_ecrCmd.Flags().StringVarP(&_ecrImageDigest, "image-digest", "", "", "Image Digest")
	_ecrCmd.Flags().StringVarP(&_ecrImageId, "image-id", "", "", "Image ID")
	_ecrCmd.Flags().StringVarP(&_ecrImageIds, "image-ids", "", "", "Image Ids")
	_ecrCmd.Flags().StringVarP(&_ecrImageManifest, "image-manifest", "", "", "Image Manifest")
	_ecrCmd.Flags().StringVarP(&_ecrImageManifestMediaType, "image-manifest-media-type", "", "", "Image Manifest Media Type")
	_ecrCmd.Flags().StringVarP(&_ecrImageScanningConfiguration, "image-scanning-configuration", "", "", "Image Scanning Configuration")
	_ecrCmd.Flags().StringVarP(&_ecrImageTag, "image-tag", "", "", "Image Tag")
	_ecrCmd.Flags().StringVarP(&_ecrImageTagMutability, "image-tag-mutability", "", "", "Image Tag Mutability")
	_ecrCmd.Flags().StringVarP(&_ecrImageTagMutabilityExclusionFilters, "image-tag-mutability-exclusion-filters", "", "", "Image Tag Mutability Exclusion Filters")
	_ecrCmd.Flags().StringVarP(&_ecrLayerDigest, "layer-digest", "", "", "Layer Digest")
	_ecrCmd.Flags().StringSliceVarP(&_ecrLayerDigests, "layer-digests", "", nil, "Layer Digests")
	_ecrCmd.Flags().StringVarP(&_ecrLayerPartBlob, "layer-part-blob", "", "", "Layer Part Blob")
	_ecrCmd.Flags().StringVarP(&_ecrLifecyclePolicy, "lifecycle-policy", "", "", "Lifecycle Policy")
	_ecrCmd.Flags().StringVarP(&_ecrLifecyclePolicyText, "lifecycle-policy-text", "", "", "Lifecycle Policy Text")
	_ecrCmd.Flags().StringVarP(&_ecrMaxResults, "max-results", "", "", "Max Results")
	_ecrCmd.Flags().StringVarP(&_ecrName, "name", "", "", "Name")
	_ecrCmd.Flags().StringVarP(&_ecrNextToken, "next-token", "", "", "Next Token")
	_ecrCmd.Flags().StringVarP(&_ecrPartFirstByte, "part-first-byte", "", "", "Part First Byte")
	_ecrCmd.Flags().StringVarP(&_ecrPartLastByte, "part-last-byte", "", "", "Part Last Byte")
	_ecrCmd.Flags().StringVarP(&_ecrPolicyText, "policy-text", "", "", "Policy Text")
	_ecrCmd.Flags().StringVarP(&_ecrPrefix, "prefix", "", "", "Prefix")
	_ecrCmd.Flags().StringSliceVarP(&_ecrPrefixes, "prefixes", "", nil, "Prefixes")
	_ecrCmd.Flags().StringVarP(&_ecrPrincipalArn, "principal-arn", "", "", "Principal ARN")
	_ecrCmd.Flags().StringVarP(&_ecrRegistryId, "registry-id", "", "", "Registry ID")
	_ecrCmd.Flags().StringSliceVarP(&_ecrRegistryIds, "registry-ids", "", nil, "Registry Ids")
	_ecrCmd.Flags().StringVarP(&_ecrReplicationConfiguration, "replication-configuration", "", "", "Replication Configuration")
	_ecrCmd.Flags().StringVarP(&_ecrRepositoryName, "repository-name", "", "", "Repository Name")
	_ecrCmd.Flags().StringSliceVarP(&_ecrRepositoryNames, "repository-names", "", nil, "Repository Names")
	_ecrCmd.Flags().StringVarP(&_ecrRepositoryPolicy, "repository-policy", "", "", "Repository Policy")
	_ecrCmd.Flags().StringVarP(&_ecrResourceArn, "resource-arn", "", "", "Resource ARN")
	_ecrCmd.Flags().StringVarP(&_ecrResourceTags, "resource-tags", "", "", "Resource Tags")
	_ecrCmd.Flags().StringVarP(&_ecrRules, "rules", "", "", "Rules")
	_ecrCmd.Flags().StringVarP(&_ecrScanType, "scan-type", "", "", "Scan Type")
	_ecrCmd.Flags().StringVarP(&_ecrSigningConfiguration, "signing-configuration", "", "", "Signing Configuration")
	_ecrCmd.Flags().StringVarP(&_ecrSubjectId, "subject-id", "", "", "Subject ID")
	_ecrCmd.Flags().StringSliceVarP(&_ecrTagKeys, "tag-keys", "", nil, "Tag Keys")
	_ecrCmd.Flags().StringVarP(&_ecrTags, "tags", "", "", "Tags")
	_ecrCmd.Flags().StringVarP(&_ecrTargetStorageClass, "target-storage-class", "", "", "Target Storage Class")
	_ecrCmd.Flags().StringVarP(&_ecrUploadId, "upload-id", "", "", "Upload ID")
	_ecrCmd.Flags().StringVarP(&_ecrUpstreamRegistry, "upstream-registry", "", "", "Upstream Registry")
	_ecrCmd.Flags().StringVarP(&_ecrUpstreamRegistryUrl, "upstream-registry-url", "", "", "Upstream Registry URL")
	_ecrCmd.Flags().StringVarP(&_ecrUpstreamRepositoryPrefix, "upstream-repository-prefix", "", "", "Upstream Repository Prefix")
	_ecrCmd.Flags().StringVarP(&_ecrValue, "value", "", "", "Value")

	_ecrCmd.Flags().BoolVarP(&_ecrBatchCheckLayerAvailability, "batch-check-layer-availability", "", false, "Batch Check Layer Availability")
	_ecrCmd.Flags().BoolVarP(&_ecrBatchDeleteImage, "batch-delete-image", "", false, "Batch Delete Image")
	_ecrCmd.Flags().BoolVarP(&_ecrBatchGetImage, "batch-get-image", "", false, "Batch Get Image")
	_ecrCmd.Flags().BoolVarP(&_ecrBatchGetRepositoryScanningConfiguration, "batch-get-repository-scanning-configuration", "", false, "Batch Get Repository Scanning Configuration")
	_ecrCmd.Flags().BoolVarP(&_ecrCompleteLayerUpload, "complete-layer-upload", "", false, "Complete Layer Upload")
	_ecrCmd.Flags().BoolVarP(&_ecrCreatePullThroughCacheRule, "create-pull-through-cache-rule", "", false, "Create Pull Through Cache Rule")
	_ecrCmd.Flags().BoolVarP(&_ecrCreateRepository, "create-repository", "", false, "Create Repository")
	_ecrCmd.Flags().BoolVarP(&_ecrCreateRepositoryCreationTemplate, "create-repository-creation-template", "", false, "Create Repository Creation Template")
	_ecrCmd.Flags().BoolVarP(&_ecrDeleteLifecyclePolicy, "delete-lifecycle-policy", "", false, "Delete Lifecycle Policy")
	_ecrCmd.Flags().BoolVarP(&_ecrDeletePullThroughCacheRule, "delete-pull-through-cache-rule", "", false, "Delete Pull Through Cache Rule")
	_ecrCmd.Flags().BoolVarP(&_ecrDeleteRegistryPolicy, "delete-registry-policy", "", false, "Delete Registry Policy")
	_ecrCmd.Flags().BoolVarP(&_ecrDeleteRepository, "delete-repository", "", false, "Delete Repository")
	_ecrCmd.Flags().BoolVarP(&_ecrDeleteRepositoryCreationTemplate, "delete-repository-creation-template", "", false, "Delete Repository Creation Template")
	_ecrCmd.Flags().BoolVarP(&_ecrDeleteRepositoryPolicy, "delete-repository-policy", "", false, "Delete Repository Policy")
	_ecrCmd.Flags().BoolVarP(&_ecrDeleteSigningConfiguration, "delete-signing-configuration", "", false, "Delete Signing Configuration")
	_ecrCmd.Flags().BoolVarP(&_ecrDeregisterPullTimeUpdateExclusion, "deregister-pull-time-update-exclusion", "", false, "Deregister Pull Time Update Exclusion")
	_ecrCmd.Flags().BoolVarP(&_ecrDescribeImageReplicationStatus, "describe-image-replication-status", "", false, "Describe Image Replication Status")
	_ecrCmd.Flags().BoolVarP(&_ecrDescribeImageScanFindings, "describe-image-scan-findings", "", false, "Describe Image Scan Findings")
	_ecrCmd.Flags().BoolVarP(&_ecrDescribeImageSigningStatus, "describe-image-signing-status", "", false, "Describe Image Signing Status")
	_ecrCmd.Flags().BoolVarP(&_ecrDescribeImages, "describe-images", "", false, "Describe Images")
	_ecrCmd.Flags().BoolVarP(&_ecrDescribePullThroughCacheRules, "describe-pull-through-cache-rules", "", false, "Describe Pull Through Cache Rules")
	_ecrCmd.Flags().BoolVarP(&_ecrDescribeRegistry, "describe-registry", "", false, "Describe Registry")
	_ecrCmd.Flags().BoolVarP(&_ecrDescribeRepositories, "describe-repositories", "", false, "Describe Repositories")
	_ecrCmd.Flags().BoolVarP(&_ecrDescribeRepositoryCreationTemplates, "describe-repository-creation-templates", "", false, "Describe Repository Creation Templates")
	_ecrCmd.Flags().BoolVarP(&_ecrGetAccountSetting, "get-account-setting", "", false, "Get Account Setting")
	_ecrCmd.Flags().BoolVarP(&_ecrGetAuthorizationToken, "get-authorization-token", "", false, "Get Authorization Token")
	_ecrCmd.Flags().BoolVarP(&_ecrGetDownloadUrlForLayer, "get-download-url-for-layer", "", false, "Get Download URL For Layer")
	_ecrCmd.Flags().BoolVarP(&_ecrGetLifecyclePolicy, "get-lifecycle-policy", "", false, "Get Lifecycle Policy")
	_ecrCmd.Flags().BoolVarP(&_ecrGetLifecyclePolicyPreview, "get-lifecycle-policy-preview", "", false, "Get Lifecycle Policy Preview")
	_ecrCmd.Flags().BoolVarP(&_ecrGetRegistryPolicy, "get-registry-policy", "", false, "Get Registry Policy")
	_ecrCmd.Flags().BoolVarP(&_ecrGetRegistryScanningConfiguration, "get-registry-scanning-configuration", "", false, "Get Registry Scanning Configuration")
	_ecrCmd.Flags().BoolVarP(&_ecrGetRepositoryPolicy, "get-repository-policy", "", false, "Get Repository Policy")
	_ecrCmd.Flags().BoolVarP(&_ecrGetSigningConfiguration, "get-signing-configuration", "", false, "Get Signing Configuration")
	_ecrCmd.Flags().BoolVarP(&_ecrInitiateLayerUpload, "initiate-layer-upload", "", false, "Initiate Layer Upload")
	_ecrCmd.Flags().BoolVarP(&_ecrListImageReferrers, "list-image-referrers", "", false, "List Image Referrers")
	_ecrCmd.Flags().BoolVarP(&_ecrListImages, "list-images", "", false, "List Images")
	_ecrCmd.Flags().BoolVarP(&_ecrListPullTimeUpdateExclusions, "list-pull-time-update-exclusions", "", false, "List Pull Time Update Exclusions")
	_ecrCmd.Flags().BoolVarP(&_ecrListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_ecrCmd.Flags().BoolVarP(&_ecrPutAccountSetting, "put-account-setting", "", false, "Put Account Setting")
	_ecrCmd.Flags().BoolVarP(&_ecrPutImage, "put-image", "", false, "Put Image")
	_ecrCmd.Flags().BoolVarP(&_ecrPutImageScanningConfiguration, "put-image-scanning-configuration", "", false, "Put Image Scanning Configuration")
	_ecrCmd.Flags().BoolVarP(&_ecrPutImageTagMutability, "put-image-tag-mutability", "", false, "Put Image Tag Mutability")
	_ecrCmd.Flags().BoolVarP(&_ecrPutLifecyclePolicy, "put-lifecycle-policy", "", false, "Put Lifecycle Policy")
	_ecrCmd.Flags().BoolVarP(&_ecrPutRegistryPolicy, "put-registry-policy", "", false, "Put Registry Policy")
	_ecrCmd.Flags().BoolVarP(&_ecrPutRegistryScanningConfiguration, "put-registry-scanning-configuration", "", false, "Put Registry Scanning Configuration")
	_ecrCmd.Flags().BoolVarP(&_ecrPutReplicationConfiguration, "put-replication-configuration", "", false, "Put Replication Configuration")
	_ecrCmd.Flags().BoolVarP(&_ecrPutSigningConfiguration, "put-signing-configuration", "", false, "Put Signing Configuration")
	_ecrCmd.Flags().BoolVarP(&_ecrRegisterPullTimeUpdateExclusion, "register-pull-time-update-exclusion", "", false, "Register Pull Time Update Exclusion")
	_ecrCmd.Flags().BoolVarP(&_ecrSetRepositoryPolicy, "set-repository-policy", "", false, "Set Repository Policy")
	_ecrCmd.Flags().BoolVarP(&_ecrStartImageScan, "start-image-scan", "", false, "Start Image Scan")
	_ecrCmd.Flags().BoolVarP(&_ecrStartLifecyclePolicyPreview, "start-lifecycle-policy-preview", "", false, "Start Lifecycle Policy Preview")
	_ecrCmd.Flags().BoolVarP(&_ecrTagResource, "tag-resource", "", false, "Tag Resource")
	_ecrCmd.Flags().BoolVarP(&_ecrUntagResource, "untag-resource", "", false, "Untag Resource")
	_ecrCmd.Flags().BoolVarP(&_ecrUpdateImageStorageClass, "update-image-storage-class", "", false, "Update Image Storage Class")
	_ecrCmd.Flags().BoolVarP(&_ecrUpdatePullThroughCacheRule, "update-pull-through-cache-rule", "", false, "Update Pull Through Cache Rule")
	_ecrCmd.Flags().BoolVarP(&_ecrUpdateRepositoryCreationTemplate, "update-repository-creation-template", "", false, "Update Repository Creation Template")
	_ecrCmd.Flags().BoolVarP(&_ecrUploadLayerPart, "upload-layer-part", "", false, "Upload Layer Part")
	_ecrCmd.Flags().BoolVarP(&_ecrValidatePullThroughCacheRule, "validate-pull-through-cache-rule", "", false, "Validate Pull Through Cache Rule")

}
