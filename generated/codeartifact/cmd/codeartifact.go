package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// codeartifactCmd represents the codeartifact command
var _codeartifactCmd = &cobra.Command{
	Use:   "codeartifact",
	Short: "AWS codeartifact CLI",
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
		client := codeartifact.NewFromConfig(cfg)
		if _codeartifactAssociateExternalConnection {
			codeartifact_AssociateExternalConnection(cfg, client)
			return
		}
		if _codeartifactCopyPackageVersions {
			codeartifact_CopyPackageVersions(cfg, client)
			return
		}
		if _codeartifactCreateDomain {
			codeartifact_CreateDomain(cfg, client)
			return
		}
		if _codeartifactCreatePackageGroup {
			codeartifact_CreatePackageGroup(cfg, client)
			return
		}
		if _codeartifactCreateRepository {
			codeartifact_CreateRepository(cfg, client)
			return
		}
		if _codeartifactDeleteDomain {
			codeartifact_DeleteDomain(cfg, client)
			return
		}
		if _codeartifactDeleteDomainPermissionsPolicy {
			codeartifact_DeleteDomainPermissionsPolicy(cfg, client)
			return
		}
		if _codeartifactDeletePackage {
			codeartifact_DeletePackage(cfg, client)
			return
		}
		if _codeartifactDeletePackageGroup {
			codeartifact_DeletePackageGroup(cfg, client)
			return
		}
		if _codeartifactDeletePackageVersions {
			codeartifact_DeletePackageVersions(cfg, client)
			return
		}
		if _codeartifactDeleteRepository {
			codeartifact_DeleteRepository(cfg, client)
			return
		}
		if _codeartifactDeleteRepositoryPermissionsPolicy {
			codeartifact_DeleteRepositoryPermissionsPolicy(cfg, client)
			return
		}
		if _codeartifactDescribeDomain {
			codeartifact_DescribeDomain(cfg, client)
			return
		}
		if _codeartifactDescribePackage {
			codeartifact_DescribePackage(cfg, client)
			return
		}
		if _codeartifactDescribePackageGroup {
			codeartifact_DescribePackageGroup(cfg, client)
			return
		}
		if _codeartifactDescribePackageVersion {
			codeartifact_DescribePackageVersion(cfg, client)
			return
		}
		if _codeartifactDescribeRepository {
			codeartifact_DescribeRepository(cfg, client)
			return
		}
		if _codeartifactDisassociateExternalConnection {
			codeartifact_DisassociateExternalConnection(cfg, client)
			return
		}
		if _codeartifactDisposePackageVersions {
			codeartifact_DisposePackageVersions(cfg, client)
			return
		}
		if _codeartifactGetAssociatedPackageGroup {
			codeartifact_GetAssociatedPackageGroup(cfg, client)
			return
		}
		if _codeartifactGetAuthorizationToken {
			codeartifact_GetAuthorizationToken(cfg, client)
			return
		}
		if _codeartifactGetDomainPermissionsPolicy {
			codeartifact_GetDomainPermissionsPolicy(cfg, client)
			return
		}
		if _codeartifactGetPackageVersionAsset {
			codeartifact_GetPackageVersionAsset(cfg, client)
			return
		}
		if _codeartifactGetPackageVersionReadme {
			codeartifact_GetPackageVersionReadme(cfg, client)
			return
		}
		if _codeartifactGetRepositoryEndpoint {
			codeartifact_GetRepositoryEndpoint(cfg, client)
			return
		}
		if _codeartifactGetRepositoryPermissionsPolicy {
			codeartifact_GetRepositoryPermissionsPolicy(cfg, client)
			return
		}
		if _codeartifactListAllowedRepositoriesForGroup {
			codeartifact_ListAllowedRepositoriesForGroup(cfg, client)
			return
		}
		if _codeartifactListAssociatedPackages {
			codeartifact_ListAssociatedPackages(cfg, client)
			return
		}
		if _codeartifactListDomains {
			codeartifact_ListDomains(cfg, client)
			return
		}
		if _codeartifactListPackageGroups {
			codeartifact_ListPackageGroups(cfg, client)
			return
		}
		if _codeartifactListPackageVersionAssets {
			codeartifact_ListPackageVersionAssets(cfg, client)
			return
		}
		if _codeartifactListPackageVersionDependencies {
			codeartifact_ListPackageVersionDependencies(cfg, client)
			return
		}
		if _codeartifactListPackageVersions {
			codeartifact_ListPackageVersions(cfg, client)
			return
		}
		if _codeartifactListPackages {
			codeartifact_ListPackages(cfg, client)
			return
		}
		if _codeartifactListRepositories {
			codeartifact_ListRepositories(cfg, client)
			return
		}
		if _codeartifactListRepositoriesInDomain {
			codeartifact_ListRepositoriesInDomain(cfg, client)
			return
		}
		if _codeartifactListSubPackageGroups {
			codeartifact_ListSubPackageGroups(cfg, client)
			return
		}
		if _codeartifactListTagsForResource {
			codeartifact_ListTagsForResource(cfg, client)
			return
		}
		if _codeartifactPublishPackageVersion {
			codeartifact_PublishPackageVersion(cfg, client)
			return
		}
		if _codeartifactPutDomainPermissionsPolicy {
			codeartifact_PutDomainPermissionsPolicy(cfg, client)
			return
		}
		if _codeartifactPutPackageOriginConfiguration {
			codeartifact_PutPackageOriginConfiguration(cfg, client)
			return
		}
		if _codeartifactPutRepositoryPermissionsPolicy {
			codeartifact_PutRepositoryPermissionsPolicy(cfg, client)
			return
		}
		if _codeartifactTagResource {
			codeartifact_TagResource(cfg, client)
			return
		}
		if _codeartifactUntagResource {
			codeartifact_UntagResource(cfg, client)
			return
		}
		if _codeartifactUpdatePackageGroup {
			codeartifact_UpdatePackageGroup(cfg, client)
			return
		}
		if _codeartifactUpdatePackageGroupOriginConfiguration {
			codeartifact_UpdatePackageGroupOriginConfiguration(cfg, client)
			return
		}
		if _codeartifactUpdatePackageVersionsStatus {
			codeartifact_UpdatePackageVersionsStatus(cfg, client)
			return
		}
		if _codeartifactUpdateRepository {
			codeartifact_UpdateRepository(cfg, client)
			return
		}

	},
}

var (
	_codeartifactAssociateExternalConnection           bool
	_codeartifactCopyPackageVersions                   bool
	_codeartifactCreateDomain                          bool
	_codeartifactCreatePackageGroup                    bool
	_codeartifactCreateRepository                      bool
	_codeartifactDeleteDomain                          bool
	_codeartifactDeleteDomainPermissionsPolicy         bool
	_codeartifactDeletePackage                         bool
	_codeartifactDeletePackageGroup                    bool
	_codeartifactDeletePackageVersions                 bool
	_codeartifactDeleteRepository                      bool
	_codeartifactDeleteRepositoryPermissionsPolicy     bool
	_codeartifactDescribeDomain                        bool
	_codeartifactDescribePackage                       bool
	_codeartifactDescribePackageGroup                  bool
	_codeartifactDescribePackageVersion                bool
	_codeartifactDescribeRepository                    bool
	_codeartifactDisassociateExternalConnection        bool
	_codeartifactDisposePackageVersions                bool
	_codeartifactGetAssociatedPackageGroup             bool
	_codeartifactGetAuthorizationToken                 bool
	_codeartifactGetDomainPermissionsPolicy            bool
	_codeartifactGetPackageVersionAsset                bool
	_codeartifactGetPackageVersionReadme               bool
	_codeartifactGetRepositoryEndpoint                 bool
	_codeartifactGetRepositoryPermissionsPolicy        bool
	_codeartifactListAllowedRepositoriesForGroup       bool
	_codeartifactListAssociatedPackages                bool
	_codeartifactListDomains                           bool
	_codeartifactListPackageGroups                     bool
	_codeartifactListPackageVersionAssets              bool
	_codeartifactListPackageVersionDependencies        bool
	_codeartifactListPackageVersions                   bool
	_codeartifactListPackages                          bool
	_codeartifactListRepositories                      bool
	_codeartifactListRepositoriesInDomain              bool
	_codeartifactListSubPackageGroups                  bool
	_codeartifactListTagsForResource                   bool
	_codeartifactPublishPackageVersion                 bool
	_codeartifactPutDomainPermissionsPolicy            bool
	_codeartifactPutPackageOriginConfiguration         bool
	_codeartifactPutRepositoryPermissionsPolicy        bool
	_codeartifactTagResource                           bool
	_codeartifactUntagResource                         bool
	_codeartifactUpdatePackageGroup                    bool
	_codeartifactUpdatePackageGroupOriginConfiguration bool
	_codeartifactUpdatePackageVersionsStatus           bool
	_codeartifactUpdateRepository                      bool

	_codeartifactAddAllowedRepositories    string
	_codeartifactAdministratorAccount      string
	_codeartifactAllowOverwrite            string
	_codeartifactAsset                     string
	_codeartifactAssetContent              string
	_codeartifactAssetName                 string
	_codeartifactAssetSHA256               string
	_codeartifactContactInfo               string
	_codeartifactDescription               string
	_codeartifactDestinationRepository     string
	_codeartifactDomain                    string
	_codeartifactDomainOwner               string
	_codeartifactDurationSeconds           string
	_codeartifactEncryptionKey             string
	_codeartifactEndpointType              string
	_codeartifactExpectedStatus            string
	_codeartifactExternalConnection        string
	_codeartifactFormat                    string
	_codeartifactIncludeFromUpstream       string
	_codeartifactMaxResults                string
	_codeartifactNamespace                 string
	_codeartifactNextToken                 string
	_codeartifactOriginRestrictionType     string
	_codeartifactOriginType                string
	_codeartifactPackage                   string
	_codeartifactPackageGroup              string
	_codeartifactPackagePrefix             string
	_codeartifactPackageVersion            string
	_codeartifactPackageVersionRevision    string
	_codeartifactPolicyDocument            string
	_codeartifactPolicyRevision            string
	_codeartifactPrefix                    string
	_codeartifactPreview                   string
	_codeartifactPublish                   string
	_codeartifactRemoveAllowedRepositories string
	_codeartifactRepository                string
	_codeartifactRepositoryPrefix          string
	_codeartifactResourceArn               string
	_codeartifactRestrictions              string
	_codeartifactSortBy                    string
	_codeartifactSourceRepository          string
	_codeartifactStatus                    string
	_codeartifactTagKeys                   []string
	_codeartifactTags                      string
	_codeartifactTargetStatus              string
	_codeartifactUnfinished                string
	_codeartifactUpstream                  string
	_codeartifactUpstreams                 string
	_codeartifactVersionRevisions          string
	_codeartifactVersions                  []string
)

// Adds an existing external connection to a repository. One external connection
// is allowed per repository.
//
// A repository can have one or more upstream repositories, or an external
// connection.
func codeartifact_AssociateExternalConnection(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.AssociateExternalConnectionInput{
		// Domain: *string, // Required
		// ExternalConnection: *string, // Required
		// Repository: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactExternalConnection) > 0 {
		input.ExternalConnection = aws.String(_codeartifactExternalConnection)
	}
	if len(_codeartifactRepository) > 0 {
		input.Repository = aws.String(_codeartifactRepository)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}

	if resp, err := client.AssociateExternalConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Copies package versions from one repository to another repository in the same
// domain.
//
// You must specify versions or versionRevisions . You cannot specify both.
func codeartifact_CopyPackageVersions(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.CopyPackageVersionsInput{
		// DestinationRepository: *string, // Required
		// Domain: *string, // Required
		// Format: types.PackageFormat, // Required
		// Package: *string, // Required
		// SourceRepository: *string, // Required
	}

	if len(_codeartifactDestinationRepository) > 0 {
		input.DestinationRepository = aws.String(_codeartifactDestinationRepository)
	}
	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactFormat) > 0 {
		if err := assignInputField(input, "Format", _codeartifactFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_codeartifactPackage) > 0 {
		input.Package = aws.String(_codeartifactPackage)
	}
	if len(_codeartifactSourceRepository) > 0 {
		input.SourceRepository = aws.String(_codeartifactSourceRepository)
	}
	if len(_codeartifactAllowOverwrite) > 0 {
		if err := assignInputField(input, "AllowOverwrite", _codeartifactAllowOverwrite); err != nil {
			log.Errorf("invalid --allow-overwrite: %s", err.Error())
			return
		}
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactIncludeFromUpstream) > 0 {
		if err := assignInputField(input, "IncludeFromUpstream", _codeartifactIncludeFromUpstream); err != nil {
			log.Errorf("invalid --include-from-upstream: %s", err.Error())
			return
		}
	}
	if len(_codeartifactNamespace) > 0 {
		input.Namespace = aws.String(_codeartifactNamespace)
	}
	if len(_codeartifactVersionRevisions) > 0 {
		if err := assignInputField(input, "VersionRevisions", _codeartifactVersionRevisions); err != nil {
			log.Errorf("invalid --version-revisions: %s", err.Error())
			return
		}
	}
	if len(_codeartifactVersions) > 0 {
		input.Versions = append([]string(nil), _codeartifactVersions...)
	}

	if resp, err := client.CopyPackageVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a domain. CodeArtifact domains make it easier to manage multiple
// repositories across an organization. You can use a domain to apply permissions
// across many repositories owned by different Amazon Web Services accounts. An
// asset is stored only once in a domain, even if it's in multiple repositories.
//
// Although you can have multiple domains, we recommend a single production domain
// that contains all published artifacts so that your development teams can find
// and share packages. You can use a second pre-production domain to test changes
// to the production domain configuration.
func codeartifact_CreateDomain(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.CreateDomainInput{
		// Domain: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactEncryptionKey) > 0 {
		input.EncryptionKey = aws.String(_codeartifactEncryptionKey)
	}
	if len(_codeartifactTags) > 0 {
		if err := assignInputField(input, "Tags", _codeartifactTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a package group. For more information about creating package groups,
// including example CLI commands, see [Create a package group]in the CodeArtifact User Guide.
//
// [Create a package group]: https://docs.aws.amazon.com/codeartifact/latest/ug/create-package-group.html
func codeartifact_CreatePackageGroup(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.CreatePackageGroupInput{
		// Domain: *string, // Required
		// PackageGroup: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactPackageGroup) > 0 {
		input.PackageGroup = aws.String(_codeartifactPackageGroup)
	}
	if len(_codeartifactContactInfo) > 0 {
		input.ContactInfo = aws.String(_codeartifactContactInfo)
	}
	if len(_codeartifactDescription) > 0 {
		input.Description = aws.String(_codeartifactDescription)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactTags) > 0 {
		if err := assignInputField(input, "Tags", _codeartifactTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePackageGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a repository.
func codeartifact_CreateRepository(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.CreateRepositoryInput{
		// Domain: *string, // Required
		// Repository: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactRepository) > 0 {
		input.Repository = aws.String(_codeartifactRepository)
	}
	if len(_codeartifactDescription) > 0 {
		input.Description = aws.String(_codeartifactDescription)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactTags) > 0 {
		if err := assignInputField(input, "Tags", _codeartifactTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_codeartifactUpstreams) > 0 {
		if err := assignInputField(input, "Upstreams", _codeartifactUpstreams); err != nil {
			log.Errorf("invalid --upstreams: %s", err.Error())
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

// Deletes a domain. You cannot delete a domain that contains repositories. If
// you want to delete a domain with repositories, first delete its repositories.
func codeartifact_DeleteDomain(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.DeleteDomainInput{
		// Domain: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}

	if resp, err := client.DeleteDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the resource policy set on a domain.
func codeartifact_DeleteDomainPermissionsPolicy(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.DeleteDomainPermissionsPolicyInput{
		// Domain: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactPolicyRevision) > 0 {
		input.PolicyRevision = aws.String(_codeartifactPolicyRevision)
	}

	if resp, err := client.DeleteDomainPermissionsPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a package and all associated package versions. A deleted package cannot
// be restored. To delete one or more package versions, use the [DeletePackageVersions]API.
//
// [DeletePackageVersions]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_DeletePackageVersions.html
func codeartifact_DeletePackage(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.DeletePackageInput{
		// Domain: *string, // Required
		// Format: types.PackageFormat, // Required
		// Package: *string, // Required
		// Repository: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactFormat) > 0 {
		if err := assignInputField(input, "Format", _codeartifactFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_codeartifactPackage) > 0 {
		input.Package = aws.String(_codeartifactPackage)
	}
	if len(_codeartifactRepository) > 0 {
		input.Repository = aws.String(_codeartifactRepository)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactNamespace) > 0 {
		input.Namespace = aws.String(_codeartifactNamespace)
	}

	if resp, err := client.DeletePackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a package group. Deleting a package group does not delete packages or
// package versions associated with the package group. When a package group is
// deleted, the direct child package groups will become children of the package
// group's direct parent package group. Therefore, if any of the child groups are
// inheriting any settings from the parent, those settings could change.
func codeartifact_DeletePackageGroup(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.DeletePackageGroupInput{
		// Domain: *string, // Required
		// PackageGroup: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactPackageGroup) > 0 {
		input.PackageGroup = aws.String(_codeartifactPackageGroup)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}

	if resp, err := client.DeletePackageGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes one or more versions of a package. A deleted package version cannot be
// restored in your repository. If you want to remove a package version from your
// repository and be able to restore it later, set its status to Archived .
// Archived packages cannot be downloaded from a repository and don't show up with
// list package APIs (for example, [ListPackageVersions]), but you can restore them using [UpdatePackageVersionsStatus].
//
// [ListPackageVersions]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_ListPackageVersions.html
// [UpdatePackageVersionsStatus]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_UpdatePackageVersionsStatus.html
func codeartifact_DeletePackageVersions(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.DeletePackageVersionsInput{
		// Domain: *string, // Required
		// Format: types.PackageFormat, // Required
		// Package: *string, // Required
		// Repository: *string, // Required
		// Versions: []string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactFormat) > 0 {
		if err := assignInputField(input, "Format", _codeartifactFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_codeartifactPackage) > 0 {
		input.Package = aws.String(_codeartifactPackage)
	}
	if len(_codeartifactRepository) > 0 {
		input.Repository = aws.String(_codeartifactRepository)
	}
	if len(_codeartifactVersions) > 0 {
		input.Versions = append([]string(nil), _codeartifactVersions...)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactExpectedStatus) > 0 {
		if err := assignInputField(input, "ExpectedStatus", _codeartifactExpectedStatus); err != nil {
			log.Errorf("invalid --expected-status: %s", err.Error())
			return
		}
	}
	if len(_codeartifactNamespace) > 0 {
		input.Namespace = aws.String(_codeartifactNamespace)
	}

	if resp, err := client.DeletePackageVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a repository.
func codeartifact_DeleteRepository(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.DeleteRepositoryInput{
		// Domain: *string, // Required
		// Repository: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactRepository) > 0 {
		input.Repository = aws.String(_codeartifactRepository)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}

	if resp, err := client.DeleteRepository(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the resource policy that is set on a repository. After a resource
// policy is deleted, the permissions allowed and denied by the deleted policy are
// removed. The effect of deleting a resource policy might not be immediate.
//
// Use DeleteRepositoryPermissionsPolicy with caution. After a policy is deleted,
// Amazon Web Services users, roles, and accounts lose permissions to perform the
// repository actions granted by the deleted policy.
func codeartifact_DeleteRepositoryPermissionsPolicy(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.DeleteRepositoryPermissionsPolicyInput{
		// Domain: *string, // Required
		// Repository: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactRepository) > 0 {
		input.Repository = aws.String(_codeartifactRepository)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactPolicyRevision) > 0 {
		input.PolicyRevision = aws.String(_codeartifactPolicyRevision)
	}

	if resp, err := client.DeleteRepositoryPermissionsPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a [DomainDescription] object that contains information about the requested domain.
//
// [DomainDescription]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_DomainDescription.html
func codeartifact_DescribeDomain(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.DescribeDomainInput{
		// Domain: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}

	if resp, err := client.DescribeDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a [PackageDescription] object that contains information about the requested package.
//
// [PackageDescription]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageDescription.html
func codeartifact_DescribePackage(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.DescribePackageInput{
		// Domain: *string, // Required
		// Format: types.PackageFormat, // Required
		// Package: *string, // Required
		// Repository: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactFormat) > 0 {
		if err := assignInputField(input, "Format", _codeartifactFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_codeartifactPackage) > 0 {
		input.Package = aws.String(_codeartifactPackage)
	}
	if len(_codeartifactRepository) > 0 {
		input.Repository = aws.String(_codeartifactRepository)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactNamespace) > 0 {
		input.Namespace = aws.String(_codeartifactNamespace)
	}

	if resp, err := client.DescribePackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a [PackageGroupDescription] object that contains information about the requested package group.
//
// [PackageGroupDescription]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageGroupDescription.html
func codeartifact_DescribePackageGroup(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.DescribePackageGroupInput{
		// Domain: *string, // Required
		// PackageGroup: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactPackageGroup) > 0 {
		input.PackageGroup = aws.String(_codeartifactPackageGroup)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}

	if resp, err := client.DescribePackageGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a [PackageVersionDescription] object that contains information about the requested package
// version.
//
// [PackageVersionDescription]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageVersionDescription.html
func codeartifact_DescribePackageVersion(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.DescribePackageVersionInput{
		// Domain: *string, // Required
		// Format: types.PackageFormat, // Required
		// Package: *string, // Required
		// PackageVersion: *string, // Required
		// Repository: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactFormat) > 0 {
		if err := assignInputField(input, "Format", _codeartifactFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_codeartifactPackage) > 0 {
		input.Package = aws.String(_codeartifactPackage)
	}
	if len(_codeartifactPackageVersion) > 0 {
		input.PackageVersion = aws.String(_codeartifactPackageVersion)
	}
	if len(_codeartifactRepository) > 0 {
		input.Repository = aws.String(_codeartifactRepository)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactNamespace) > 0 {
		input.Namespace = aws.String(_codeartifactNamespace)
	}

	if resp, err := client.DescribePackageVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a RepositoryDescription object that contains detailed information
// about the requested repository.
func codeartifact_DescribeRepository(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.DescribeRepositoryInput{
		// Domain: *string, // Required
		// Repository: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactRepository) > 0 {
		input.Repository = aws.String(_codeartifactRepository)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}

	if resp, err := client.DescribeRepository(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes an existing external connection from a repository.
func codeartifact_DisassociateExternalConnection(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.DisassociateExternalConnectionInput{
		// Domain: *string, // Required
		// ExternalConnection: *string, // Required
		// Repository: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactExternalConnection) > 0 {
		input.ExternalConnection = aws.String(_codeartifactExternalConnection)
	}
	if len(_codeartifactRepository) > 0 {
		input.Repository = aws.String(_codeartifactRepository)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}

	if resp, err := client.DisassociateExternalConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the assets in package versions and sets the package versions' status
// to Disposed . A disposed package version cannot be restored in your repository
// because its assets are deleted.
//
// To view all disposed package versions in a repository, use [ListPackageVersions] and set the [status]
// parameter to Disposed .
//
// To view information about a disposed package version, use [DescribePackageVersion].
//
// [DescribePackageVersion]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_DescribePackageVersion.html
// [ListPackageVersions]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_ListPackageVersions.html
// [status]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_ListPackageVersions.html#API_ListPackageVersions_RequestSyntax
func codeartifact_DisposePackageVersions(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.DisposePackageVersionsInput{
		// Domain: *string, // Required
		// Format: types.PackageFormat, // Required
		// Package: *string, // Required
		// Repository: *string, // Required
		// Versions: []string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactFormat) > 0 {
		if err := assignInputField(input, "Format", _codeartifactFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_codeartifactPackage) > 0 {
		input.Package = aws.String(_codeartifactPackage)
	}
	if len(_codeartifactRepository) > 0 {
		input.Repository = aws.String(_codeartifactRepository)
	}
	if len(_codeartifactVersions) > 0 {
		input.Versions = append([]string(nil), _codeartifactVersions...)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactExpectedStatus) > 0 {
		if err := assignInputField(input, "ExpectedStatus", _codeartifactExpectedStatus); err != nil {
			log.Errorf("invalid --expected-status: %s", err.Error())
			return
		}
	}
	if len(_codeartifactNamespace) > 0 {
		input.Namespace = aws.String(_codeartifactNamespace)
	}
	if len(_codeartifactVersionRevisions) > 0 {
		if err := assignInputField(input, "VersionRevisions", _codeartifactVersionRevisions); err != nil {
			log.Errorf("invalid --version-revisions: %s", err.Error())
			return
		}
	}

	if resp, err := client.DisposePackageVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the most closely associated package group to the specified package.
// This API does not require that the package exist in any repository in the
// domain. As such, GetAssociatedPackageGroup can be used to see which package
// group's origin configuration applies to a package before that package is in a
// repository. This can be helpful to check if public packages are blocked without
// ingesting them.
//
// For information package group association and matching, see [Package group definition syntax and matching behavior] in the
// CodeArtifact User Guide.
//
// [Package group definition syntax and matching behavior]: https://docs.aws.amazon.com/codeartifact/latest/ug/package-group-definition-syntax-matching-behavior.html
func codeartifact_GetAssociatedPackageGroup(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.GetAssociatedPackageGroupInput{
		// Domain: *string, // Required
		// Format: types.PackageFormat, // Required
		// Package: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactFormat) > 0 {
		if err := assignInputField(input, "Format", _codeartifactFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_codeartifactPackage) > 0 {
		input.Package = aws.String(_codeartifactPackage)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactNamespace) > 0 {
		input.Namespace = aws.String(_codeartifactNamespace)
	}

	if resp, err := client.GetAssociatedPackageGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates a temporary authorization token for accessing repositories in the
// domain. This API requires the codeartifact:GetAuthorizationToken and
// sts:GetServiceBearerToken permissions. For more information about authorization
// tokens, see [CodeArtifact authentication and tokens].
//
// CodeArtifact authorization tokens are valid for a period of 12 hours when
// created with the login command. You can call login periodically to refresh the
// token. When you create an authorization token with the GetAuthorizationToken
// API, you can set a custom authorization period, up to a maximum of 12 hours,
// with the durationSeconds parameter.
//
// The authorization period begins after login or GetAuthorizationToken is called.
// If login or GetAuthorizationToken is called while assuming a role, the token
// lifetime is independent of the maximum session duration of the role. For
// example, if you call sts assume-role and specify a session duration of 15
// minutes, then generate a CodeArtifact authorization token, the token will be
// valid for the full authorization period even though this is longer than the
// 15-minute session duration.
//
// See [Using IAM Roles] for more information on controlling session duration.
//
// [Using IAM Roles]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html
// [CodeArtifact authentication and tokens]: https://docs.aws.amazon.com/codeartifact/latest/ug/tokens-authentication.html
func codeartifact_GetAuthorizationToken(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.GetAuthorizationTokenInput{
		// Domain: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactDurationSeconds) > 0 {
		if err := assignInputField(input, "DurationSeconds", _codeartifactDurationSeconds); err != nil {
			log.Errorf("invalid --duration-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetAuthorizationToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the resource policy attached to the specified domain.
// The policy is a resource-based policy, not an identity-based policy. For more
// information, see [Identity-based policies and resource-based policies]in the IAM User Guide.
//
// [Identity-based policies and resource-based policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_identity-vs-resource.html
func codeartifact_GetDomainPermissionsPolicy(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.GetDomainPermissionsPolicyInput{
		// Domain: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}

	if resp, err := client.GetDomainPermissionsPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an asset (or file) that is in a package. For example, for a Maven
// package version, use GetPackageVersionAsset to download a JAR file, a POM file,
// or any other assets in the package version.
func codeartifact_GetPackageVersionAsset(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.GetPackageVersionAssetInput{
		// Asset: *string, // Required
		// Domain: *string, // Required
		// Format: types.PackageFormat, // Required
		// Package: *string, // Required
		// PackageVersion: *string, // Required
		// Repository: *string, // Required
	}

	if len(_codeartifactAsset) > 0 {
		input.Asset = aws.String(_codeartifactAsset)
	}
	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactFormat) > 0 {
		if err := assignInputField(input, "Format", _codeartifactFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_codeartifactPackage) > 0 {
		input.Package = aws.String(_codeartifactPackage)
	}
	if len(_codeartifactPackageVersion) > 0 {
		input.PackageVersion = aws.String(_codeartifactPackageVersion)
	}
	if len(_codeartifactRepository) > 0 {
		input.Repository = aws.String(_codeartifactRepository)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactNamespace) > 0 {
		input.Namespace = aws.String(_codeartifactNamespace)
	}
	if len(_codeartifactPackageVersionRevision) > 0 {
		input.PackageVersionRevision = aws.String(_codeartifactPackageVersionRevision)
	}

	if resp, err := client.GetPackageVersionAsset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the readme file or descriptive text for a package version.
// The returned text might contain formatting. For example, it might contain
// formatting for Markdown or reStructuredText.
func codeartifact_GetPackageVersionReadme(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.GetPackageVersionReadmeInput{
		// Domain: *string, // Required
		// Format: types.PackageFormat, // Required
		// Package: *string, // Required
		// PackageVersion: *string, // Required
		// Repository: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactFormat) > 0 {
		if err := assignInputField(input, "Format", _codeartifactFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_codeartifactPackage) > 0 {
		input.Package = aws.String(_codeartifactPackage)
	}
	if len(_codeartifactPackageVersion) > 0 {
		input.PackageVersion = aws.String(_codeartifactPackageVersion)
	}
	if len(_codeartifactRepository) > 0 {
		input.Repository = aws.String(_codeartifactRepository)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactNamespace) > 0 {
		input.Namespace = aws.String(_codeartifactNamespace)
	}

	if resp, err := client.GetPackageVersionReadme(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the endpoint of a repository for a specific package format. A
// repository has one endpoint for each package format:
//
// - cargo
//
// - generic
//
// - maven
//
// - npm
//
// - nuget
//
// - pypi
//
// - ruby
//
// - swift
func codeartifact_GetRepositoryEndpoint(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.GetRepositoryEndpointInput{
		// Domain: *string, // Required
		// Format: types.PackageFormat, // Required
		// Repository: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactFormat) > 0 {
		if err := assignInputField(input, "Format", _codeartifactFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_codeartifactRepository) > 0 {
		input.Repository = aws.String(_codeartifactRepository)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactEndpointType) > 0 {
		if err := assignInputField(input, "EndpointType", _codeartifactEndpointType); err != nil {
			log.Errorf("invalid --endpoint-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetRepositoryEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the resource policy that is set on a repository.
func codeartifact_GetRepositoryPermissionsPolicy(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.GetRepositoryPermissionsPolicyInput{
		// Domain: *string, // Required
		// Repository: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactRepository) > 0 {
		input.Repository = aws.String(_codeartifactRepository)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}

	if resp, err := client.GetRepositoryPermissionsPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the repositories in the added repositories list of the specified
// restriction type for a package group. For more information about restriction
// types and added repository lists, see [Package group origin controls]in the CodeArtifact User Guide.
//
// [Package group origin controls]: https://docs.aws.amazon.com/codeartifact/latest/ug/package-group-origin-controls.html
func codeartifact_ListAllowedRepositoriesForGroup(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.ListAllowedRepositoriesForGroupInput{
		// Domain: *string, // Required
		// OriginRestrictionType: types.PackageGroupOriginRestrictionType, // Required
		// PackageGroup: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactOriginRestrictionType) > 0 {
		if err := assignInputField(input, "OriginRestrictionType", _codeartifactOriginRestrictionType); err != nil {
			log.Errorf("invalid --origin-restriction-type: %s", err.Error())
			return
		}
	}
	if len(_codeartifactPackageGroup) > 0 {
		input.PackageGroup = aws.String(_codeartifactPackageGroup)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codeartifactMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codeartifactNextToken) > 0 {
		input.NextToken = aws.String(_codeartifactNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAllowedRepositoriesForGroup(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codeartifact.ListAllowedRepositoriesForGroupOutput
	p := codeartifact.NewListAllowedRepositoriesForGroupPaginator(client, input)
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

// Returns a list of packages associated with the requested package group. For
// information package group association and matching, see [Package group definition syntax and matching behavior]in the CodeArtifact
// User Guide.
//
// [Package group definition syntax and matching behavior]: https://docs.aws.amazon.com/codeartifact/latest/ug/package-group-definition-syntax-matching-behavior.html
func codeartifact_ListAssociatedPackages(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.ListAssociatedPackagesInput{
		// Domain: *string, // Required
		// PackageGroup: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactPackageGroup) > 0 {
		input.PackageGroup = aws.String(_codeartifactPackageGroup)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codeartifactMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codeartifactNextToken) > 0 {
		input.NextToken = aws.String(_codeartifactNextToken)
	}
	if len(_codeartifactPreview) > 0 {
		if err := assignInputField(input, "Preview", _codeartifactPreview); err != nil {
			log.Errorf("invalid --preview: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAssociatedPackages(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codeartifact.ListAssociatedPackagesOutput
	p := codeartifact.NewListAssociatedPackagesPaginator(client, input)
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

// Returns a list of [DomainSummary] objects for all domains owned by the Amazon Web Services
// account that makes this call. Each returned DomainSummary object contains
// information about a domain.
//
// [DomainSummary]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageVersionDescription.html
func codeartifact_ListDomains(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.ListDomainsInput{}

	if len(_codeartifactMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codeartifactMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codeartifactNextToken) > 0 {
		input.NextToken = aws.String(_codeartifactNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDomains(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codeartifact.ListDomainsOutput
	p := codeartifact.NewListDomainsPaginator(client, input)
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

// Returns a list of package groups in the requested domain.
func codeartifact_ListPackageGroups(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.ListPackageGroupsInput{
		// Domain: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codeartifactMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codeartifactNextToken) > 0 {
		input.NextToken = aws.String(_codeartifactNextToken)
	}
	if len(_codeartifactPrefix) > 0 {
		input.Prefix = aws.String(_codeartifactPrefix)
	}

	if disablePaginator() {
		if resp, err := client.ListPackageGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codeartifact.ListPackageGroupsOutput
	p := codeartifact.NewListPackageGroupsPaginator(client, input)
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

// Returns a list of [AssetSummary] objects for assets in a package version.
//
// [AssetSummary]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_AssetSummary.html
func codeartifact_ListPackageVersionAssets(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.ListPackageVersionAssetsInput{
		// Domain: *string, // Required
		// Format: types.PackageFormat, // Required
		// Package: *string, // Required
		// PackageVersion: *string, // Required
		// Repository: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactFormat) > 0 {
		if err := assignInputField(input, "Format", _codeartifactFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_codeartifactPackage) > 0 {
		input.Package = aws.String(_codeartifactPackage)
	}
	if len(_codeartifactPackageVersion) > 0 {
		input.PackageVersion = aws.String(_codeartifactPackageVersion)
	}
	if len(_codeartifactRepository) > 0 {
		input.Repository = aws.String(_codeartifactRepository)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codeartifactMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codeartifactNamespace) > 0 {
		input.Namespace = aws.String(_codeartifactNamespace)
	}
	if len(_codeartifactNextToken) > 0 {
		input.NextToken = aws.String(_codeartifactNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPackageVersionAssets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codeartifact.ListPackageVersionAssetsOutput
	p := codeartifact.NewListPackageVersionAssetsPaginator(client, input)
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

// Returns the direct dependencies for a package version. The dependencies are
// returned as [PackageDependency]objects. CodeArtifact extracts the dependencies for a package
// version from the metadata file for the package format (for example, the
// package.json file for npm packages and the pom.xml file for Maven). Any package
// version dependencies that are not listed in the configuration file are not
// returned.
//
// [PackageDependency]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageDependency.html
func codeartifact_ListPackageVersionDependencies(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.ListPackageVersionDependenciesInput{
		// Domain: *string, // Required
		// Format: types.PackageFormat, // Required
		// Package: *string, // Required
		// PackageVersion: *string, // Required
		// Repository: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactFormat) > 0 {
		if err := assignInputField(input, "Format", _codeartifactFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_codeartifactPackage) > 0 {
		input.Package = aws.String(_codeartifactPackage)
	}
	if len(_codeartifactPackageVersion) > 0 {
		input.PackageVersion = aws.String(_codeartifactPackageVersion)
	}
	if len(_codeartifactRepository) > 0 {
		input.Repository = aws.String(_codeartifactRepository)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactNamespace) > 0 {
		input.Namespace = aws.String(_codeartifactNamespace)
	}
	if len(_codeartifactNextToken) > 0 {
		input.NextToken = aws.String(_codeartifactNextToken)
	}

	if resp, err := client.ListPackageVersionDependencies(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of [PackageVersionSummary] objects for package versions in a repository that match the
// request parameters. Package versions of all statuses will be returned by default
// when calling list-package-versions with no --status parameter.
//
// [PackageVersionSummary]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageVersionSummary.html
func codeartifact_ListPackageVersions(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.ListPackageVersionsInput{
		// Domain: *string, // Required
		// Format: types.PackageFormat, // Required
		// Package: *string, // Required
		// Repository: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactFormat) > 0 {
		if err := assignInputField(input, "Format", _codeartifactFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_codeartifactPackage) > 0 {
		input.Package = aws.String(_codeartifactPackage)
	}
	if len(_codeartifactRepository) > 0 {
		input.Repository = aws.String(_codeartifactRepository)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codeartifactMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codeartifactNamespace) > 0 {
		input.Namespace = aws.String(_codeartifactNamespace)
	}
	if len(_codeartifactNextToken) > 0 {
		input.NextToken = aws.String(_codeartifactNextToken)
	}
	if len(_codeartifactOriginType) > 0 {
		if err := assignInputField(input, "OriginType", _codeartifactOriginType); err != nil {
			log.Errorf("invalid --origin-type: %s", err.Error())
			return
		}
	}
	if len(_codeartifactSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _codeartifactSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_codeartifactStatus) > 0 {
		if err := assignInputField(input, "Status", _codeartifactStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListPackageVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codeartifact.ListPackageVersionsOutput
	p := codeartifact.NewListPackageVersionsPaginator(client, input)
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

// Returns a list of [PackageSummary] objects for packages in a repository that match the request
// parameters.
//
// [PackageSummary]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageSummary.html
func codeartifact_ListPackages(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.ListPackagesInput{
		// Domain: *string, // Required
		// Repository: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactRepository) > 0 {
		input.Repository = aws.String(_codeartifactRepository)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactFormat) > 0 {
		if err := assignInputField(input, "Format", _codeartifactFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_codeartifactMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codeartifactMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codeartifactNamespace) > 0 {
		input.Namespace = aws.String(_codeartifactNamespace)
	}
	if len(_codeartifactNextToken) > 0 {
		input.NextToken = aws.String(_codeartifactNextToken)
	}
	if len(_codeartifactPackagePrefix) > 0 {
		input.PackagePrefix = aws.String(_codeartifactPackagePrefix)
	}
	if len(_codeartifactPublish) > 0 {
		if err := assignInputField(input, "Publish", _codeartifactPublish); err != nil {
			log.Errorf("invalid --publish: %s", err.Error())
			return
		}
	}
	if len(_codeartifactUpstream) > 0 {
		if err := assignInputField(input, "Upstream", _codeartifactUpstream); err != nil {
			log.Errorf("invalid --upstream: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListPackages(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codeartifact.ListPackagesOutput
	p := codeartifact.NewListPackagesPaginator(client, input)
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

// Returns a list of [RepositorySummary] objects. Each RepositorySummary contains information about
// a repository in the specified Amazon Web Services account and that matches the
// input parameters.
//
// [RepositorySummary]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_RepositorySummary.html
func codeartifact_ListRepositories(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.ListRepositoriesInput{}

	if len(_codeartifactMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codeartifactMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codeartifactNextToken) > 0 {
		input.NextToken = aws.String(_codeartifactNextToken)
	}
	if len(_codeartifactRepositoryPrefix) > 0 {
		input.RepositoryPrefix = aws.String(_codeartifactRepositoryPrefix)
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

	var results []*codeartifact.ListRepositoriesOutput
	p := codeartifact.NewListRepositoriesPaginator(client, input)
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

// Returns a list of [RepositorySummary] objects. Each RepositorySummary contains information about
// a repository in the specified domain and that matches the input parameters.
//
// [RepositorySummary]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_RepositorySummary.html
func codeartifact_ListRepositoriesInDomain(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.ListRepositoriesInDomainInput{
		// Domain: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactAdministratorAccount) > 0 {
		input.AdministratorAccount = aws.String(_codeartifactAdministratorAccount)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codeartifactMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codeartifactNextToken) > 0 {
		input.NextToken = aws.String(_codeartifactNextToken)
	}
	if len(_codeartifactRepositoryPrefix) > 0 {
		input.RepositoryPrefix = aws.String(_codeartifactRepositoryPrefix)
	}

	if disablePaginator() {
		if resp, err := client.ListRepositoriesInDomain(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codeartifact.ListRepositoriesInDomainOutput
	p := codeartifact.NewListRepositoriesInDomainPaginator(client, input)
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

// Returns a list of direct children of the specified package group.
// For information package group hierarchy, see [Package group definition syntax and matching behavior] in the CodeArtifact User Guide.
//
// [Package group definition syntax and matching behavior]: https://docs.aws.amazon.com/codeartifact/latest/ug/package-group-definition-syntax-matching-behavior.html
func codeartifact_ListSubPackageGroups(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.ListSubPackageGroupsInput{
		// Domain: *string, // Required
		// PackageGroup: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactPackageGroup) > 0 {
		input.PackageGroup = aws.String(_codeartifactPackageGroup)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codeartifactMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codeartifactNextToken) > 0 {
		input.NextToken = aws.String(_codeartifactNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSubPackageGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codeartifact.ListSubPackageGroupsOutput
	p := codeartifact.NewListSubPackageGroupsPaginator(client, input)
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

// Gets information about Amazon Web Services tags for a specified Amazon Resource
// Name (ARN) in CodeArtifact.
func codeartifact_ListTagsForResource(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_codeartifactResourceArn) > 0 {
		input.ResourceArn = aws.String(_codeartifactResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new package version containing one or more assets (or files).
// The unfinished flag can be used to keep the package version in the Unfinished
// state until all of its assets have been uploaded (see [Package version status]in the CodeArtifact user
// guide). To set the package version’s status to Published , omit the unfinished
// flag when uploading the final asset, or set the status using [UpdatePackageVersionStatus]. Once a package
// version’s status is set to Published , it cannot change back to Unfinished .
//
// Only generic packages can be published using this API. For more information,
// see [Using generic packages]in the CodeArtifact User Guide.
//
// [Using generic packages]: https://docs.aws.amazon.com/codeartifact/latest/ug/using-generic.html
// [UpdatePackageVersionStatus]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_UpdatePackageVersionsStatus.html
// [Package version status]: https://docs.aws.amazon.com/codeartifact/latest/ug/packages-overview.html#package-version-status.html#package-version-status
func codeartifact_PublishPackageVersion(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.PublishPackageVersionInput{
		// AssetContent: io.Reader, // Required
		// AssetName: *string, // Required
		// AssetSHA256: *string, // Required
		// Domain: *string, // Required
		// Format: types.PackageFormat, // Required
		// Package: *string, // Required
		// PackageVersion: *string, // Required
		// Repository: *string, // Required
	}

	if len(_codeartifactAssetContent) > 0 {
		if err := assignInputField(input, "AssetContent", _codeartifactAssetContent); err != nil {
			log.Errorf("invalid --asset-content: %s", err.Error())
			return
		}
	}
	if len(_codeartifactAssetName) > 0 {
		input.AssetName = aws.String(_codeartifactAssetName)
	}
	if len(_codeartifactAssetSHA256) > 0 {
		input.AssetSHA256 = aws.String(_codeartifactAssetSHA256)
	}
	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactFormat) > 0 {
		if err := assignInputField(input, "Format", _codeartifactFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_codeartifactPackage) > 0 {
		input.Package = aws.String(_codeartifactPackage)
	}
	if len(_codeartifactPackageVersion) > 0 {
		input.PackageVersion = aws.String(_codeartifactPackageVersion)
	}
	if len(_codeartifactRepository) > 0 {
		input.Repository = aws.String(_codeartifactRepository)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactNamespace) > 0 {
		input.Namespace = aws.String(_codeartifactNamespace)
	}
	if len(_codeartifactUnfinished) > 0 {
		if err := assignInputField(input, "Unfinished", _codeartifactUnfinished); err != nil {
			log.Errorf("invalid --unfinished: %s", err.Error())
			return
		}
	}

	if resp, err := client.PublishPackageVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets a resource policy on a domain that specifies permissions to access it.
// When you call PutDomainPermissionsPolicy , the resource policy on the domain is
// ignored when evaluting permissions. This ensures that the owner of a domain
// cannot lock themselves out of the domain, which would prevent them from being
// able to update the resource policy.
func codeartifact_PutDomainPermissionsPolicy(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.PutDomainPermissionsPolicyInput{
		// Domain: *string, // Required
		// PolicyDocument: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactPolicyDocument) > 0 {
		input.PolicyDocument = aws.String(_codeartifactPolicyDocument)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactPolicyRevision) > 0 {
		input.PolicyRevision = aws.String(_codeartifactPolicyRevision)
	}

	if resp, err := client.PutDomainPermissionsPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the package origin configuration for a package.
// The package origin configuration determines how new versions of a package can
// be added to a repository. You can allow or block direct publishing of new
// package versions, or ingestion and retaining of new package versions from an
// external connection or upstream source. For more information about package
// origin controls and configuration, see [Editing package origin controls]in the CodeArtifact User Guide.
//
// PutPackageOriginConfiguration can be called on a package that doesn't yet exist
// in the repository. When called on a package that does not exist, a package is
// created in the repository with no versions and the requested restrictions are
// set on the package. This can be used to preemptively block ingesting or
// retaining any versions from external connections or upstream repositories, or to
// block publishing any versions of the package into the repository before
// connecting any package managers or publishers to the repository.
//
// [Editing package origin controls]: https://docs.aws.amazon.com/codeartifact/latest/ug/package-origin-controls.html
func codeartifact_PutPackageOriginConfiguration(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.PutPackageOriginConfigurationInput{
		// Domain: *string, // Required
		// Format: types.PackageFormat, // Required
		// Package: *string, // Required
		// Repository: *string, // Required
		// Restrictions: *types.PackageOriginRestrictions, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactFormat) > 0 {
		if err := assignInputField(input, "Format", _codeartifactFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_codeartifactPackage) > 0 {
		input.Package = aws.String(_codeartifactPackage)
	}
	if len(_codeartifactRepository) > 0 {
		input.Repository = aws.String(_codeartifactRepository)
	}
	if len(_codeartifactRestrictions) > 0 {
		if err := assignInputField(input, "Restrictions", _codeartifactRestrictions); err != nil {
			log.Errorf("invalid --restrictions: %s", err.Error())
			return
		}
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactNamespace) > 0 {
		input.Namespace = aws.String(_codeartifactNamespace)
	}

	if resp, err := client.PutPackageOriginConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the resource policy on a repository that specifies permissions to access
// it.
//
// When you call PutRepositoryPermissionsPolicy , the resource policy on the
// repository is ignored when evaluting permissions. This ensures that the owner of
// a repository cannot lock themselves out of the repository, which would prevent
// them from being able to update the resource policy.
func codeartifact_PutRepositoryPermissionsPolicy(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.PutRepositoryPermissionsPolicyInput{
		// Domain: *string, // Required
		// PolicyDocument: *string, // Required
		// Repository: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactPolicyDocument) > 0 {
		input.PolicyDocument = aws.String(_codeartifactPolicyDocument)
	}
	if len(_codeartifactRepository) > 0 {
		input.Repository = aws.String(_codeartifactRepository)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactPolicyRevision) > 0 {
		input.PolicyRevision = aws.String(_codeartifactPolicyRevision)
	}

	if resp, err := client.PutRepositoryPermissionsPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates tags for a resource in CodeArtifact.
func codeartifact_TagResource(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_codeartifactResourceArn) > 0 {
		input.ResourceArn = aws.String(_codeartifactResourceArn)
	}
	if len(_codeartifactTags) > 0 {
		if err := assignInputField(input, "Tags", _codeartifactTags); err != nil {
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

// Removes tags from a resource in CodeArtifact.
func codeartifact_UntagResource(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_codeartifactResourceArn) > 0 {
		input.ResourceArn = aws.String(_codeartifactResourceArn)
	}
	if len(_codeartifactTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _codeartifactTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a package group. This API cannot be used to update a package group's
// origin configuration or pattern. To update a package group's origin
// configuration, use [UpdatePackageGroupOriginConfiguration].
//
// [UpdatePackageGroupOriginConfiguration]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_UpdatePackageGroupOriginConfiguration.html
func codeartifact_UpdatePackageGroup(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.UpdatePackageGroupInput{
		// Domain: *string, // Required
		// PackageGroup: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactPackageGroup) > 0 {
		input.PackageGroup = aws.String(_codeartifactPackageGroup)
	}
	if len(_codeartifactContactInfo) > 0 {
		input.ContactInfo = aws.String(_codeartifactContactInfo)
	}
	if len(_codeartifactDescription) > 0 {
		input.Description = aws.String(_codeartifactDescription)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}

	if resp, err := client.UpdatePackageGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the package origin configuration for a package group.
// The package origin configuration determines how new versions of a package can
// be added to a repository. You can allow or block direct publishing of new
// package versions, or ingestion and retaining of new package versions from an
// external connection or upstream source. For more information about package group
// origin controls and configuration, see [Package group origin controls]in the CodeArtifact User Guide.
//
// [Package group origin controls]: https://docs.aws.amazon.com/codeartifact/latest/ug/package-group-origin-controls.html
func codeartifact_UpdatePackageGroupOriginConfiguration(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.UpdatePackageGroupOriginConfigurationInput{
		// Domain: *string, // Required
		// PackageGroup: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactPackageGroup) > 0 {
		input.PackageGroup = aws.String(_codeartifactPackageGroup)
	}
	if len(_codeartifactAddAllowedRepositories) > 0 {
		if err := assignInputField(input, "AddAllowedRepositories", _codeartifactAddAllowedRepositories); err != nil {
			log.Errorf("invalid --add-allowed-repositories: %s", err.Error())
			return
		}
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactRemoveAllowedRepositories) > 0 {
		if err := assignInputField(input, "RemoveAllowedRepositories", _codeartifactRemoveAllowedRepositories); err != nil {
			log.Errorf("invalid --remove-allowed-repositories: %s", err.Error())
			return
		}
	}
	if len(_codeartifactRestrictions) > 0 {
		if err := assignInputField(input, "Restrictions", _codeartifactRestrictions); err != nil {
			log.Errorf("invalid --restrictions: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePackageGroupOriginConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the status of one or more versions of a package. Using
// UpdatePackageVersionsStatus , you can update the status of package versions to
// Archived , Published , or Unlisted . To set the status of a package version to
// Disposed , use [DisposePackageVersions].
//
// [DisposePackageVersions]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_DisposePackageVersions.html
func codeartifact_UpdatePackageVersionsStatus(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.UpdatePackageVersionsStatusInput{
		// Domain: *string, // Required
		// Format: types.PackageFormat, // Required
		// Package: *string, // Required
		// Repository: *string, // Required
		// TargetStatus: types.PackageVersionStatus, // Required
		// Versions: []string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactFormat) > 0 {
		if err := assignInputField(input, "Format", _codeartifactFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_codeartifactPackage) > 0 {
		input.Package = aws.String(_codeartifactPackage)
	}
	if len(_codeartifactRepository) > 0 {
		input.Repository = aws.String(_codeartifactRepository)
	}
	if len(_codeartifactTargetStatus) > 0 {
		if err := assignInputField(input, "TargetStatus", _codeartifactTargetStatus); err != nil {
			log.Errorf("invalid --target-status: %s", err.Error())
			return
		}
	}
	if len(_codeartifactVersions) > 0 {
		input.Versions = append([]string(nil), _codeartifactVersions...)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactExpectedStatus) > 0 {
		if err := assignInputField(input, "ExpectedStatus", _codeartifactExpectedStatus); err != nil {
			log.Errorf("invalid --expected-status: %s", err.Error())
			return
		}
	}
	if len(_codeartifactNamespace) > 0 {
		input.Namespace = aws.String(_codeartifactNamespace)
	}
	if len(_codeartifactVersionRevisions) > 0 {
		if err := assignInputField(input, "VersionRevisions", _codeartifactVersionRevisions); err != nil {
			log.Errorf("invalid --version-revisions: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePackageVersionsStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the properties of a repository.
func codeartifact_UpdateRepository(cfg aws.Config, client *codeartifact.Client) {
	input := &codeartifact.UpdateRepositoryInput{
		// Domain: *string, // Required
		// Repository: *string, // Required
	}

	if len(_codeartifactDomain) > 0 {
		input.Domain = aws.String(_codeartifactDomain)
	}
	if len(_codeartifactRepository) > 0 {
		input.Repository = aws.String(_codeartifactRepository)
	}
	if len(_codeartifactDescription) > 0 {
		input.Description = aws.String(_codeartifactDescription)
	}
	if len(_codeartifactDomainOwner) > 0 {
		input.DomainOwner = aws.String(_codeartifactDomainOwner)
	}
	if len(_codeartifactUpstreams) > 0 {
		if err := assignInputField(input, "Upstreams", _codeartifactUpstreams); err != nil {
			log.Errorf("invalid --upstreams: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRepository(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_codeartifactCmd)
	_codeartifactCmd.Flags().SortFlags = false

	_codeartifactCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_codeartifactCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_codeartifactCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_codeartifactCmd.Flags().StringVarP(&_codeartifactAddAllowedRepositories, "add-allowed-repositories", "", "", "Add Allowed Repositories")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactAdministratorAccount, "administrator-account", "", "", "Administrator Account")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactAllowOverwrite, "allow-overwrite", "", "", "Allow Overwrite")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactAsset, "asset", "", "", "Asset")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactAssetContent, "asset-content", "", "", "Asset Content")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactAssetName, "asset-name", "", "", "Asset Name")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactAssetSHA256, "asset-sha256", "", "", "Asset SHA256")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactContactInfo, "contact-info", "", "", "Contact Info")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactDescription, "description", "", "", "Description")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactDestinationRepository, "destination-repository", "", "", "Destination Repository")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactDomain, "domain", "", "", "Domain")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactDomainOwner, "domain-owner", "", "", "Domain Owner")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactDurationSeconds, "duration-seconds", "", "", "Duration Seconds")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactEncryptionKey, "encryption-key", "", "", "Encryption Key")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactEndpointType, "endpoint-type", "", "", "Endpoint Type")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactExpectedStatus, "expected-status", "", "", "Expected Status")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactExternalConnection, "external-connection", "", "", "External Connection")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactFormat, "format", "", "", "Format")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactIncludeFromUpstream, "include-from-upstream", "", "", "Include From Upstream")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactMaxResults, "max-results", "", "", "Max Results")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactNamespace, "namespace", "", "", "Namespace")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactNextToken, "next-token", "", "", "Next Token")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactOriginRestrictionType, "origin-restriction-type", "", "", "Origin Restriction Type")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactOriginType, "origin-type", "", "", "Origin Type")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactPackage, "package", "", "", "Package")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactPackageGroup, "package-group", "", "", "Package Group")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactPackagePrefix, "package-prefix", "", "", "Package Prefix")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactPackageVersion, "package-version", "", "", "Package Version")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactPackageVersionRevision, "package-version-revision", "", "", "Package Version Revision")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactPolicyDocument, "policy-document", "", "", "Policy Document")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactPolicyRevision, "policy-revision", "", "", "Policy Revision")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactPrefix, "prefix", "", "", "Prefix")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactPreview, "preview", "", "", "Preview")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactPublish, "publish", "", "", "Publish")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactRemoveAllowedRepositories, "remove-allowed-repositories", "", "", "Remove Allowed Repositories")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactRepository, "repository", "", "", "Repository")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactRepositoryPrefix, "repository-prefix", "", "", "Repository Prefix")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactResourceArn, "resource-arn", "", "", "Resource ARN")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactRestrictions, "restrictions", "", "", "Restrictions")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactSortBy, "sort-by", "", "", "Sort By")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactSourceRepository, "source-repository", "", "", "Source Repository")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactStatus, "status", "", "", "Status")
	_codeartifactCmd.Flags().StringSliceVarP(&_codeartifactTagKeys, "tag-keys", "", nil, "Tag Keys")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactTags, "tags", "", "", "Tags")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactTargetStatus, "target-status", "", "", "Target Status")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactUnfinished, "unfinished", "", "", "Unfinished")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactUpstream, "upstream", "", "", "Upstream")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactUpstreams, "upstreams", "", "", "Upstreams")
	_codeartifactCmd.Flags().StringVarP(&_codeartifactVersionRevisions, "version-revisions", "", "", "Version Revisions")
	_codeartifactCmd.Flags().StringSliceVarP(&_codeartifactVersions, "versions", "", nil, "Versions")

	_codeartifactCmd.Flags().BoolVarP(&_codeartifactAssociateExternalConnection, "associate-external-connection", "", false, "Associate External Connection")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactCopyPackageVersions, "copy-package-versions", "", false, "Copy Package Versions")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactCreateDomain, "create-domain", "", false, "Create Domain")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactCreatePackageGroup, "create-package-group", "", false, "Create Package Group")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactCreateRepository, "create-repository", "", false, "Create Repository")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactDeleteDomain, "delete-domain", "", false, "Delete Domain")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactDeleteDomainPermissionsPolicy, "delete-domain-permissions-policy", "", false, "Delete Domain Permissions Policy")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactDeletePackage, "delete-package", "", false, "Delete Package")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactDeletePackageGroup, "delete-package-group", "", false, "Delete Package Group")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactDeletePackageVersions, "delete-package-versions", "", false, "Delete Package Versions")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactDeleteRepository, "delete-repository", "", false, "Delete Repository")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactDeleteRepositoryPermissionsPolicy, "delete-repository-permissions-policy", "", false, "Delete Repository Permissions Policy")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactDescribeDomain, "describe-domain", "", false, "Describe Domain")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactDescribePackage, "describe-package", "", false, "Describe Package")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactDescribePackageGroup, "describe-package-group", "", false, "Describe Package Group")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactDescribePackageVersion, "describe-package-version", "", false, "Describe Package Version")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactDescribeRepository, "describe-repository", "", false, "Describe Repository")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactDisassociateExternalConnection, "disassociate-external-connection", "", false, "Disassociate External Connection")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactDisposePackageVersions, "dispose-package-versions", "", false, "Dispose Package Versions")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactGetAssociatedPackageGroup, "get-associated-package-group", "", false, "Get Associated Package Group")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactGetAuthorizationToken, "get-authorization-token", "", false, "Get Authorization Token")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactGetDomainPermissionsPolicy, "get-domain-permissions-policy", "", false, "Get Domain Permissions Policy")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactGetPackageVersionAsset, "get-package-version-asset", "", false, "Get Package Version Asset")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactGetPackageVersionReadme, "get-package-version-readme", "", false, "Get Package Version Readme")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactGetRepositoryEndpoint, "get-repository-endpoint", "", false, "Get Repository Endpoint")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactGetRepositoryPermissionsPolicy, "get-repository-permissions-policy", "", false, "Get Repository Permissions Policy")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactListAllowedRepositoriesForGroup, "list-allowed-repositories-for-group", "", false, "List Allowed Repositories For Group")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactListAssociatedPackages, "list-associated-packages", "", false, "List Associated Packages")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactListDomains, "list-domains", "", false, "List Domains")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactListPackageGroups, "list-package-groups", "", false, "List Package Groups")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactListPackageVersionAssets, "list-package-version-assets", "", false, "List Package Version Assets")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactListPackageVersionDependencies, "list-package-version-dependencies", "", false, "List Package Version Dependencies")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactListPackageVersions, "list-package-versions", "", false, "List Package Versions")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactListPackages, "list-packages", "", false, "List Packages")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactListRepositories, "list-repositories", "", false, "List Repositories")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactListRepositoriesInDomain, "list-repositories-in-domain", "", false, "List Repositories In Domain")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactListSubPackageGroups, "list-sub-package-groups", "", false, "List Sub Package Groups")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactPublishPackageVersion, "publish-package-version", "", false, "Publish Package Version")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactPutDomainPermissionsPolicy, "put-domain-permissions-policy", "", false, "Put Domain Permissions Policy")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactPutPackageOriginConfiguration, "put-package-origin-configuration", "", false, "Put Package Origin Configuration")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactPutRepositoryPermissionsPolicy, "put-repository-permissions-policy", "", false, "Put Repository Permissions Policy")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactTagResource, "tag-resource", "", false, "Tag Resource")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactUntagResource, "untag-resource", "", false, "Untag Resource")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactUpdatePackageGroup, "update-package-group", "", false, "Update Package Group")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactUpdatePackageGroupOriginConfiguration, "update-package-group-origin-configuration", "", false, "Update Package Group Origin Configuration")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactUpdatePackageVersionsStatus, "update-package-versions-status", "", false, "Update Package Versions Status")
	_codeartifactCmd.Flags().BoolVarP(&_codeartifactUpdateRepository, "update-repository", "", false, "Update Repository")

}
