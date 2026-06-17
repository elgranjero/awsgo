package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ecrpublicCmd represents the ecrpublic command
var _ecrpublicCmd = &cobra.Command{
	Use:   "ecrpublic",
	Short: "AWS ecrpublic CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := ecrpublic.NewFromConfig(cfg)
		if _ecrpublicBatchCheckLayerAvailability {
			ecrpublic_BatchCheckLayerAvailability(cfg, client)
			return
		}
		if _ecrpublicBatchDeleteImage {
			ecrpublic_BatchDeleteImage(cfg, client)
			return
		}
		if _ecrpublicCompleteLayerUpload {
			ecrpublic_CompleteLayerUpload(cfg, client)
			return
		}
		if _ecrpublicCreateRepository {
			ecrpublic_CreateRepository(cfg, client)
			return
		}
		if _ecrpublicDeleteRepository {
			ecrpublic_DeleteRepository(cfg, client)
			return
		}
		if _ecrpublicDeleteRepositoryPolicy {
			ecrpublic_DeleteRepositoryPolicy(cfg, client)
			return
		}
		if _ecrpublicDescribeImageTags {
			ecrpublic_DescribeImageTags(cfg, client)
			return
		}
		if _ecrpublicDescribeImages {
			ecrpublic_DescribeImages(cfg, client)
			return
		}
		if _ecrpublicDescribeRegistries {
			ecrpublic_DescribeRegistries(cfg, client)
			return
		}
		if _ecrpublicDescribeRepositories {
			ecrpublic_DescribeRepositories(cfg, client)
			return
		}
		if _ecrpublicGetAuthorizationToken {
			ecrpublic_GetAuthorizationToken(cfg, client)
			return
		}
		if _ecrpublicGetRegistryCatalogData {
			ecrpublic_GetRegistryCatalogData(cfg, client)
			return
		}
		if _ecrpublicGetRepositoryCatalogData {
			ecrpublic_GetRepositoryCatalogData(cfg, client)
			return
		}
		if _ecrpublicGetRepositoryPolicy {
			ecrpublic_GetRepositoryPolicy(cfg, client)
			return
		}
		if _ecrpublicInitiateLayerUpload {
			ecrpublic_InitiateLayerUpload(cfg, client)
			return
		}
		if _ecrpublicListTagsForResource {
			ecrpublic_ListTagsForResource(cfg, client)
			return
		}
		if _ecrpublicPutImage {
			ecrpublic_PutImage(cfg, client)
			return
		}
		if _ecrpublicPutRegistryCatalogData {
			ecrpublic_PutRegistryCatalogData(cfg, client)
			return
		}
		if _ecrpublicPutRepositoryCatalogData {
			ecrpublic_PutRepositoryCatalogData(cfg, client)
			return
		}
		if _ecrpublicSetRepositoryPolicy {
			ecrpublic_SetRepositoryPolicy(cfg, client)
			return
		}
		if _ecrpublicTagResource {
			ecrpublic_TagResource(cfg, client)
			return
		}
		if _ecrpublicUntagResource {
			ecrpublic_UntagResource(cfg, client)
			return
		}
		if _ecrpublicUploadLayerPart {
			ecrpublic_UploadLayerPart(cfg, client)
			return
		}

	},
}

var (
	_ecrpublicBatchCheckLayerAvailability bool
	_ecrpublicBatchDeleteImage            bool
	_ecrpublicCompleteLayerUpload         bool
	_ecrpublicCreateRepository            bool
	_ecrpublicDeleteRepository            bool
	_ecrpublicDeleteRepositoryPolicy      bool
	_ecrpublicDescribeImageTags           bool
	_ecrpublicDescribeImages              bool
	_ecrpublicDescribeRegistries          bool
	_ecrpublicDescribeRepositories        bool
	_ecrpublicGetAuthorizationToken       bool
	_ecrpublicGetRegistryCatalogData      bool
	_ecrpublicGetRepositoryCatalogData    bool
	_ecrpublicGetRepositoryPolicy         bool
	_ecrpublicInitiateLayerUpload         bool
	_ecrpublicListTagsForResource         bool
	_ecrpublicPutImage                    bool
	_ecrpublicPutRegistryCatalogData      bool
	_ecrpublicPutRepositoryCatalogData    bool
	_ecrpublicSetRepositoryPolicy         bool
	_ecrpublicTagResource                 bool
	_ecrpublicUntagResource               bool
	_ecrpublicUploadLayerPart             bool

	_ecrpublicCatalogData            string
	_ecrpublicDisplayName            string
	_ecrpublicForce                  string
	_ecrpublicImageDigest            string
	_ecrpublicImageIds               string
	_ecrpublicImageManifest          string
	_ecrpublicImageManifestMediaType string
	_ecrpublicImageTag               string
	_ecrpublicLayerDigests           []string
	_ecrpublicLayerPartBlob          string
	_ecrpublicMaxResults             string
	_ecrpublicNextToken              string
	_ecrpublicPartFirstByte          string
	_ecrpublicPartLastByte           string
	_ecrpublicPolicyText             string
	_ecrpublicRegistryId             string
	_ecrpublicRepositoryName         string
	_ecrpublicRepositoryNames        []string
	_ecrpublicResourceArn            string
	_ecrpublicTagKeys                []string
	_ecrpublicTags                   string
	_ecrpublicUploadId               string
)

// Checks the availability of one or more image layers that are within a
// repository in a public registry. When an image is pushed to a repository, each
// image layer is checked to verify if it has been uploaded before. If it has been
// uploaded, then the image layer is skipped.
//
// This operation is used by the Amazon ECR proxy and is not generally used by
// customers for pulling and pushing images. In most cases, you should use the
// docker CLI to pull, tag, and push images.
func ecrpublic_BatchCheckLayerAvailability(cfg aws.Config, client *ecrpublic.Client) {
	input := &ecrpublic.BatchCheckLayerAvailabilityInput{
		// LayerDigests: []string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_ecrpublicLayerDigests) > 0 {
		input.LayerDigests = append([]string(nil), _ecrpublicLayerDigests...)
	}
	if len(_ecrpublicRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrpublicRepositoryName)
	}
	if len(_ecrpublicRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrpublicRegistryId)
	}

	if resp, err := client.BatchCheckLayerAvailability(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a list of specified images that are within a repository in a public
// registry. Images are specified with either an imageTag or imageDigest .
//
// You can remove a tag from an image by specifying the image's tag in your
// request. When you remove the last tag from an image, the image is deleted from
// your repository.
//
// You can completely delete an image (and all of its tags) by specifying the
// digest of the image in your request.
func ecrpublic_BatchDeleteImage(cfg aws.Config, client *ecrpublic.Client) {
	input := &ecrpublic.BatchDeleteImageInput{
		// ImageIds: []types.ImageIdentifier, // Required
		// RepositoryName: *string, // Required
	}

	if len(_ecrpublicImageIds) > 0 {
		if err := assignInputField(input, "ImageIds", _ecrpublicImageIds); err != nil {
			log.Errorf("invalid --image-ids: %s", err.Error())
			return
		}
	}
	if len(_ecrpublicRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrpublicRepositoryName)
	}
	if len(_ecrpublicRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrpublicRegistryId)
	}

	if resp, err := client.BatchDeleteImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Informs Amazon ECR that the image layer upload is complete for a specified
// public registry, repository name, and upload ID. You can optionally provide a
// sha256 digest of the image layer for data validation purposes.
//
// When an image is pushed, the CompleteLayerUpload API is called once for each
// new image layer to verify that the upload is complete.
//
// This operation is used by the Amazon ECR proxy and is not generally used by
// customers for pulling and pushing images. In most cases, you should use the
// docker CLI to pull, tag, and push images.
func ecrpublic_CompleteLayerUpload(cfg aws.Config, client *ecrpublic.Client) {
	input := &ecrpublic.CompleteLayerUploadInput{
		// LayerDigests: []string, // Required
		// RepositoryName: *string, // Required
		// UploadId: *string, // Required
	}

	if len(_ecrpublicLayerDigests) > 0 {
		input.LayerDigests = append([]string(nil), _ecrpublicLayerDigests...)
	}
	if len(_ecrpublicRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrpublicRepositoryName)
	}
	if len(_ecrpublicUploadId) > 0 {
		input.UploadId = aws.String(_ecrpublicUploadId)
	}
	if len(_ecrpublicRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrpublicRegistryId)
	}

	if resp, err := client.CompleteLayerUpload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a repository in a public registry. For more information, see [Amazon ECR repositories] in the
// Amazon Elastic Container Registry User Guide.
//
// [Amazon ECR repositories]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/Repositories.html
func ecrpublic_CreateRepository(cfg aws.Config, client *ecrpublic.Client) {
	input := &ecrpublic.CreateRepositoryInput{
		// RepositoryName: *string, // Required
	}

	if len(_ecrpublicRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrpublicRepositoryName)
	}
	if len(_ecrpublicCatalogData) > 0 {
		if err := assignInputField(input, "CatalogData", _ecrpublicCatalogData); err != nil {
			log.Errorf("invalid --catalog-data: %s", err.Error())
			return
		}
	}
	if len(_ecrpublicTags) > 0 {
		if err := assignInputField(input, "Tags", _ecrpublicTags); err != nil {
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

// Deletes a repository in a public registry. If the repository contains images,
// you must either manually delete all images in the repository or use the force
// option. This option deletes all images on your behalf before deleting the
// repository.
func ecrpublic_DeleteRepository(cfg aws.Config, client *ecrpublic.Client) {
	input := &ecrpublic.DeleteRepositoryInput{
		// RepositoryName: *string, // Required
	}

	if len(_ecrpublicRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrpublicRepositoryName)
	}
	if len(_ecrpublicForce) > 0 {
		if err := assignInputField(input, "Force", _ecrpublicForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}
	if len(_ecrpublicRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrpublicRegistryId)
	}

	if resp, err := client.DeleteRepository(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the repository policy that's associated with the specified repository.
func ecrpublic_DeleteRepositoryPolicy(cfg aws.Config, client *ecrpublic.Client) {
	input := &ecrpublic.DeleteRepositoryPolicyInput{
		// RepositoryName: *string, // Required
	}

	if len(_ecrpublicRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrpublicRepositoryName)
	}
	if len(_ecrpublicRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrpublicRegistryId)
	}

	if resp, err := client.DeleteRepositoryPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the image tag details for a repository in a public registry.
func ecrpublic_DescribeImageTags(cfg aws.Config, client *ecrpublic.Client) {
	input := &ecrpublic.DescribeImageTagsInput{
		// RepositoryName: *string, // Required
	}

	if len(_ecrpublicRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrpublicRepositoryName)
	}
	if len(_ecrpublicMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ecrpublicMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ecrpublicNextToken) > 0 {
		input.NextToken = aws.String(_ecrpublicNextToken)
	}
	if len(_ecrpublicRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrpublicRegistryId)
	}

	if disablePaginator() {
		if resp, err := client.DescribeImageTags(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ecrpublic.DescribeImageTagsOutput
	p := ecrpublic.NewDescribeImageTagsPaginator(client, input)
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

// Returns metadata that's related to the images in a repository in a public
// registry.
//
// Beginning with Docker version 1.9, the Docker client compresses image layers
// before pushing them to a V2 Docker registry. The output of the docker images
// command shows the uncompressed image size. Therefore, it might return a larger
// image size than the image sizes that are returned by DescribeImages.
func ecrpublic_DescribeImages(cfg aws.Config, client *ecrpublic.Client) {
	input := &ecrpublic.DescribeImagesInput{
		// RepositoryName: *string, // Required
	}

	if len(_ecrpublicRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrpublicRepositoryName)
	}
	if len(_ecrpublicImageIds) > 0 {
		if err := assignInputField(input, "ImageIds", _ecrpublicImageIds); err != nil {
			log.Errorf("invalid --image-ids: %s", err.Error())
			return
		}
	}
	if len(_ecrpublicMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ecrpublicMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ecrpublicNextToken) > 0 {
		input.NextToken = aws.String(_ecrpublicNextToken)
	}
	if len(_ecrpublicRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrpublicRegistryId)
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

	var results []*ecrpublic.DescribeImagesOutput
	p := ecrpublic.NewDescribeImagesPaginator(client, input)
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

// Returns details for a public registry.
func ecrpublic_DescribeRegistries(cfg aws.Config, client *ecrpublic.Client) {
	input := &ecrpublic.DescribeRegistriesInput{}

	if len(_ecrpublicMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ecrpublicMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ecrpublicNextToken) > 0 {
		input.NextToken = aws.String(_ecrpublicNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeRegistries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ecrpublic.DescribeRegistriesOutput
	p := ecrpublic.NewDescribeRegistriesPaginator(client, input)
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

// Describes repositories that are in a public registry.
func ecrpublic_DescribeRepositories(cfg aws.Config, client *ecrpublic.Client) {
	input := &ecrpublic.DescribeRepositoriesInput{}

	if len(_ecrpublicMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ecrpublicMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ecrpublicNextToken) > 0 {
		input.NextToken = aws.String(_ecrpublicNextToken)
	}
	if len(_ecrpublicRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrpublicRegistryId)
	}
	if len(_ecrpublicRepositoryNames) > 0 {
		input.RepositoryNames = append([]string(nil), _ecrpublicRepositoryNames...)
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

	var results []*ecrpublic.DescribeRepositoriesOutput
	p := ecrpublic.NewDescribeRepositoriesPaginator(client, input)
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

// Retrieves an authorization token. An authorization token represents your IAM
// authentication credentials. You can use it to access any Amazon ECR registry
// that your IAM principal has access to. The authorization token is valid for 12
// hours. This API requires the ecr-public:GetAuthorizationToken and
// sts:GetServiceBearerToken permissions.
func ecrpublic_GetAuthorizationToken(cfg aws.Config, client *ecrpublic.Client) {
	input := &ecrpublic.GetAuthorizationTokenInput{}

	if resp, err := client.GetAuthorizationToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves catalog metadata for a public registry.
func ecrpublic_GetRegistryCatalogData(cfg aws.Config, client *ecrpublic.Client) {
	input := &ecrpublic.GetRegistryCatalogDataInput{}

	if resp, err := client.GetRegistryCatalogData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve catalog metadata for a repository in a public registry. This metadata
// is displayed publicly in the Amazon ECR Public Gallery.
func ecrpublic_GetRepositoryCatalogData(cfg aws.Config, client *ecrpublic.Client) {
	input := &ecrpublic.GetRepositoryCatalogDataInput{
		// RepositoryName: *string, // Required
	}

	if len(_ecrpublicRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrpublicRepositoryName)
	}
	if len(_ecrpublicRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrpublicRegistryId)
	}

	if resp, err := client.GetRepositoryCatalogData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the repository policy for the specified repository.
func ecrpublic_GetRepositoryPolicy(cfg aws.Config, client *ecrpublic.Client) {
	input := &ecrpublic.GetRepositoryPolicyInput{
		// RepositoryName: *string, // Required
	}

	if len(_ecrpublicRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrpublicRepositoryName)
	}
	if len(_ecrpublicRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrpublicRegistryId)
	}

	if resp, err := client.GetRepositoryPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Notifies Amazon ECR that you intend to upload an image layer.
// When an image is pushed, the InitiateLayerUpload API is called once for each
// image layer that hasn't already been uploaded. Whether an image layer uploads is
// determined by the BatchCheckLayerAvailability API action.
//
// This operation is used by the Amazon ECR proxy and is not generally used by
// customers for pulling and pushing images. In most cases, you should use the
// docker CLI to pull, tag, and push images.
func ecrpublic_InitiateLayerUpload(cfg aws.Config, client *ecrpublic.Client) {
	input := &ecrpublic.InitiateLayerUploadInput{
		// RepositoryName: *string, // Required
	}

	if len(_ecrpublicRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrpublicRepositoryName)
	}
	if len(_ecrpublicRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrpublicRegistryId)
	}

	if resp, err := client.InitiateLayerUpload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List the tags for an Amazon ECR Public resource.
func ecrpublic_ListTagsForResource(cfg aws.Config, client *ecrpublic.Client) {
	input := &ecrpublic.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_ecrpublicResourceArn) > 0 {
		input.ResourceArn = aws.String(_ecrpublicResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates the image manifest and tags that are associated with an
// image.
//
// When an image is pushed and all new image layers have been uploaded, the
// PutImage API is called once to create or update the image manifest and the tags
// that are associated with the image.
//
// This operation is used by the Amazon ECR proxy and is not generally used by
// customers for pulling and pushing images. In most cases, you should use the
// docker CLI to pull, tag, and push images.
func ecrpublic_PutImage(cfg aws.Config, client *ecrpublic.Client) {
	input := &ecrpublic.PutImageInput{
		// ImageManifest: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_ecrpublicImageManifest) > 0 {
		input.ImageManifest = aws.String(_ecrpublicImageManifest)
	}
	if len(_ecrpublicRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrpublicRepositoryName)
	}
	if len(_ecrpublicImageDigest) > 0 {
		input.ImageDigest = aws.String(_ecrpublicImageDigest)
	}
	if len(_ecrpublicImageManifestMediaType) > 0 {
		input.ImageManifestMediaType = aws.String(_ecrpublicImageManifestMediaType)
	}
	if len(_ecrpublicImageTag) > 0 {
		input.ImageTag = aws.String(_ecrpublicImageTag)
	}
	if len(_ecrpublicRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrpublicRegistryId)
	}

	if resp, err := client.PutImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create or update the catalog data for a public registry.
func ecrpublic_PutRegistryCatalogData(cfg aws.Config, client *ecrpublic.Client) {
	input := &ecrpublic.PutRegistryCatalogDataInput{}

	if len(_ecrpublicDisplayName) > 0 {
		input.DisplayName = aws.String(_ecrpublicDisplayName)
	}

	if resp, err := client.PutRegistryCatalogData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates the catalog data for a repository in a public registry.
func ecrpublic_PutRepositoryCatalogData(cfg aws.Config, client *ecrpublic.Client) {
	input := &ecrpublic.PutRepositoryCatalogDataInput{
		// CatalogData: *types.RepositoryCatalogDataInput, // Required
		// RepositoryName: *string, // Required
	}

	if len(_ecrpublicCatalogData) > 0 {
		if err := assignInputField(input, "CatalogData", _ecrpublicCatalogData); err != nil {
			log.Errorf("invalid --catalog-data: %s", err.Error())
			return
		}
	}
	if len(_ecrpublicRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrpublicRepositoryName)
	}
	if len(_ecrpublicRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrpublicRegistryId)
	}

	if resp, err := client.PutRepositoryCatalogData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies a repository policy to the specified public repository to control
// access permissions. For more information, see [Amazon ECR Repository Policies]in the Amazon Elastic Container
// Registry User Guide.
//
// [Amazon ECR Repository Policies]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/repository-policies.html
func ecrpublic_SetRepositoryPolicy(cfg aws.Config, client *ecrpublic.Client) {
	input := &ecrpublic.SetRepositoryPolicyInput{
		// PolicyText: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_ecrpublicPolicyText) > 0 {
		input.PolicyText = aws.String(_ecrpublicPolicyText)
	}
	if len(_ecrpublicRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrpublicRepositoryName)
	}
	if len(_ecrpublicForce) > 0 {
		if err := assignInputField(input, "Force", _ecrpublicForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}
	if len(_ecrpublicRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrpublicRegistryId)
	}

	if resp, err := client.SetRepositoryPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified tags to a resource with the specified resourceArn . If
// existing tags on a resource aren't specified in the request parameters, they
// aren't changed. When a resource is deleted, the tags associated with that
// resource are also deleted.
func ecrpublic_TagResource(cfg aws.Config, client *ecrpublic.Client) {
	input := &ecrpublic.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_ecrpublicResourceArn) > 0 {
		input.ResourceArn = aws.String(_ecrpublicResourceArn)
	}
	if len(_ecrpublicTags) > 0 {
		if err := assignInputField(input, "Tags", _ecrpublicTags); err != nil {
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
func ecrpublic_UntagResource(cfg aws.Config, client *ecrpublic.Client) {
	input := &ecrpublic.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_ecrpublicResourceArn) > 0 {
		input.ResourceArn = aws.String(_ecrpublicResourceArn)
	}
	if len(_ecrpublicTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _ecrpublicTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Uploads an image layer part to Amazon ECR.
// When an image is pushed, each new image layer is uploaded in parts. The maximum
// size of each image layer part can be 20971520 bytes (about 20MB). The
// UploadLayerPart API is called once for each new image layer part.
//
// This operation is used by the Amazon ECR proxy and is not generally used by
// customers for pulling and pushing images. In most cases, you should use the
// docker CLI to pull, tag, and push images.
func ecrpublic_UploadLayerPart(cfg aws.Config, client *ecrpublic.Client) {
	input := &ecrpublic.UploadLayerPartInput{
		// LayerPartBlob: []byte, // Required
		// PartFirstByte: *int64, // Required
		// PartLastByte: *int64, // Required
		// RepositoryName: *string, // Required
		// UploadId: *string, // Required
	}

	if len(_ecrpublicLayerPartBlob) > 0 {
		if err := assignInputField(input, "LayerPartBlob", _ecrpublicLayerPartBlob); err != nil {
			log.Errorf("invalid --layer-part-blob: %s", err.Error())
			return
		}
	}
	if len(_ecrpublicPartFirstByte) > 0 {
		if err := assignInputField(input, "PartFirstByte", _ecrpublicPartFirstByte); err != nil {
			log.Errorf("invalid --part-first-byte: %s", err.Error())
			return
		}
	}
	if len(_ecrpublicPartLastByte) > 0 {
		if err := assignInputField(input, "PartLastByte", _ecrpublicPartLastByte); err != nil {
			log.Errorf("invalid --part-last-byte: %s", err.Error())
			return
		}
	}
	if len(_ecrpublicRepositoryName) > 0 {
		input.RepositoryName = aws.String(_ecrpublicRepositoryName)
	}
	if len(_ecrpublicUploadId) > 0 {
		input.UploadId = aws.String(_ecrpublicUploadId)
	}
	if len(_ecrpublicRegistryId) > 0 {
		input.RegistryId = aws.String(_ecrpublicRegistryId)
	}

	if resp, err := client.UploadLayerPart(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_ecrpublicCmd)
	_ecrpublicCmd.Flags().SortFlags = false

	_ecrpublicCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_ecrpublicCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_ecrpublicCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_ecrpublicCmd.Flags().StringVarP(&_ecrpublicCatalogData, "catalog-data", "", "", "Catalog Data")
	_ecrpublicCmd.Flags().StringVarP(&_ecrpublicDisplayName, "display-name", "", "", "Display Name")
	_ecrpublicCmd.Flags().StringVarP(&_ecrpublicForce, "force", "", "", "Force")
	_ecrpublicCmd.Flags().StringVarP(&_ecrpublicImageDigest, "image-digest", "", "", "Image Digest")
	_ecrpublicCmd.Flags().StringVarP(&_ecrpublicImageIds, "image-ids", "", "", "Image Ids")
	_ecrpublicCmd.Flags().StringVarP(&_ecrpublicImageManifest, "image-manifest", "", "", "Image Manifest")
	_ecrpublicCmd.Flags().StringVarP(&_ecrpublicImageManifestMediaType, "image-manifest-media-type", "", "", "Image Manifest Media Type")
	_ecrpublicCmd.Flags().StringVarP(&_ecrpublicImageTag, "image-tag", "", "", "Image Tag")
	_ecrpublicCmd.Flags().StringSliceVarP(&_ecrpublicLayerDigests, "layer-digests", "", nil, "Layer Digests")
	_ecrpublicCmd.Flags().StringVarP(&_ecrpublicLayerPartBlob, "layer-part-blob", "", "", "Layer Part Blob")
	_ecrpublicCmd.Flags().StringVarP(&_ecrpublicMaxResults, "max-results", "", "", "Max Results")
	_ecrpublicCmd.Flags().StringVarP(&_ecrpublicNextToken, "next-token", "", "", "Next Token")
	_ecrpublicCmd.Flags().StringVarP(&_ecrpublicPartFirstByte, "part-first-byte", "", "", "Part First Byte")
	_ecrpublicCmd.Flags().StringVarP(&_ecrpublicPartLastByte, "part-last-byte", "", "", "Part Last Byte")
	_ecrpublicCmd.Flags().StringVarP(&_ecrpublicPolicyText, "policy-text", "", "", "Policy Text")
	_ecrpublicCmd.Flags().StringVarP(&_ecrpublicRegistryId, "registry-id", "", "", "Registry ID")
	_ecrpublicCmd.Flags().StringVarP(&_ecrpublicRepositoryName, "repository-name", "", "", "Repository Name")
	_ecrpublicCmd.Flags().StringSliceVarP(&_ecrpublicRepositoryNames, "repository-names", "", nil, "Repository Names")
	_ecrpublicCmd.Flags().StringVarP(&_ecrpublicResourceArn, "resource-arn", "", "", "Resource ARN")
	_ecrpublicCmd.Flags().StringSliceVarP(&_ecrpublicTagKeys, "tag-keys", "", nil, "Tag Keys")
	_ecrpublicCmd.Flags().StringVarP(&_ecrpublicTags, "tags", "", "", "Tags")
	_ecrpublicCmd.Flags().StringVarP(&_ecrpublicUploadId, "upload-id", "", "", "Upload ID")

	_ecrpublicCmd.Flags().BoolVarP(&_ecrpublicBatchCheckLayerAvailability, "batch-check-layer-availability", "", false, "Batch Check Layer Availability")
	_ecrpublicCmd.Flags().BoolVarP(&_ecrpublicBatchDeleteImage, "batch-delete-image", "", false, "Batch Delete Image")
	_ecrpublicCmd.Flags().BoolVarP(&_ecrpublicCompleteLayerUpload, "complete-layer-upload", "", false, "Complete Layer Upload")
	_ecrpublicCmd.Flags().BoolVarP(&_ecrpublicCreateRepository, "create-repository", "", false, "Create Repository")
	_ecrpublicCmd.Flags().BoolVarP(&_ecrpublicDeleteRepository, "delete-repository", "", false, "Delete Repository")
	_ecrpublicCmd.Flags().BoolVarP(&_ecrpublicDeleteRepositoryPolicy, "delete-repository-policy", "", false, "Delete Repository Policy")
	_ecrpublicCmd.Flags().BoolVarP(&_ecrpublicDescribeImageTags, "describe-image-tags", "", false, "Describe Image Tags")
	_ecrpublicCmd.Flags().BoolVarP(&_ecrpublicDescribeImages, "describe-images", "", false, "Describe Images")
	_ecrpublicCmd.Flags().BoolVarP(&_ecrpublicDescribeRegistries, "describe-registries", "", false, "Describe Registries")
	_ecrpublicCmd.Flags().BoolVarP(&_ecrpublicDescribeRepositories, "describe-repositories", "", false, "Describe Repositories")
	_ecrpublicCmd.Flags().BoolVarP(&_ecrpublicGetAuthorizationToken, "get-authorization-token", "", false, "Get Authorization Token")
	_ecrpublicCmd.Flags().BoolVarP(&_ecrpublicGetRegistryCatalogData, "get-registry-catalog-data", "", false, "Get Registry Catalog Data")
	_ecrpublicCmd.Flags().BoolVarP(&_ecrpublicGetRepositoryCatalogData, "get-repository-catalog-data", "", false, "Get Repository Catalog Data")
	_ecrpublicCmd.Flags().BoolVarP(&_ecrpublicGetRepositoryPolicy, "get-repository-policy", "", false, "Get Repository Policy")
	_ecrpublicCmd.Flags().BoolVarP(&_ecrpublicInitiateLayerUpload, "initiate-layer-upload", "", false, "Initiate Layer Upload")
	_ecrpublicCmd.Flags().BoolVarP(&_ecrpublicListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_ecrpublicCmd.Flags().BoolVarP(&_ecrpublicPutImage, "put-image", "", false, "Put Image")
	_ecrpublicCmd.Flags().BoolVarP(&_ecrpublicPutRegistryCatalogData, "put-registry-catalog-data", "", false, "Put Registry Catalog Data")
	_ecrpublicCmd.Flags().BoolVarP(&_ecrpublicPutRepositoryCatalogData, "put-repository-catalog-data", "", false, "Put Repository Catalog Data")
	_ecrpublicCmd.Flags().BoolVarP(&_ecrpublicSetRepositoryPolicy, "set-repository-policy", "", false, "Set Repository Policy")
	_ecrpublicCmd.Flags().BoolVarP(&_ecrpublicTagResource, "tag-resource", "", false, "Tag Resource")
	_ecrpublicCmd.Flags().BoolVarP(&_ecrpublicUntagResource, "untag-resource", "", false, "Untag Resource")
	_ecrpublicCmd.Flags().BoolVarP(&_ecrpublicUploadLayerPart, "upload-layer-part", "", false, "Upload Layer Part")

}
