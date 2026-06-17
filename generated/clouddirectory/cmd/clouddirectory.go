package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// clouddirectoryCmd represents the clouddirectory command
var _clouddirectoryCmd = &cobra.Command{
	Use:   "clouddirectory",
	Short: "AWS clouddirectory CLI",
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
		client := clouddirectory.NewFromConfig(cfg)
		if _clouddirectoryAddFacetToObject {
			clouddirectory_AddFacetToObject(cfg, client)
			return
		}
		if _clouddirectoryApplySchema {
			clouddirectory_ApplySchema(cfg, client)
			return
		}
		if _clouddirectoryAttachObject {
			clouddirectory_AttachObject(cfg, client)
			return
		}
		if _clouddirectoryAttachPolicy {
			clouddirectory_AttachPolicy(cfg, client)
			return
		}
		if _clouddirectoryAttachToIndex {
			clouddirectory_AttachToIndex(cfg, client)
			return
		}
		if _clouddirectoryAttachTypedLink {
			clouddirectory_AttachTypedLink(cfg, client)
			return
		}
		if _clouddirectoryBatchRead {
			clouddirectory_BatchRead(cfg, client)
			return
		}
		if _clouddirectoryBatchWrite {
			clouddirectory_BatchWrite(cfg, client)
			return
		}
		if _clouddirectoryCreateDirectory {
			clouddirectory_CreateDirectory(cfg, client)
			return
		}
		if _clouddirectoryCreateFacet {
			clouddirectory_CreateFacet(cfg, client)
			return
		}
		if _clouddirectoryCreateIndex {
			clouddirectory_CreateIndex(cfg, client)
			return
		}
		if _clouddirectoryCreateObject {
			clouddirectory_CreateObject(cfg, client)
			return
		}
		if _clouddirectoryCreateSchema {
			clouddirectory_CreateSchema(cfg, client)
			return
		}
		if _clouddirectoryCreateTypedLinkFacet {
			clouddirectory_CreateTypedLinkFacet(cfg, client)
			return
		}
		if _clouddirectoryDeleteDirectory {
			clouddirectory_DeleteDirectory(cfg, client)
			return
		}
		if _clouddirectoryDeleteFacet {
			clouddirectory_DeleteFacet(cfg, client)
			return
		}
		if _clouddirectoryDeleteObject {
			clouddirectory_DeleteObject(cfg, client)
			return
		}
		if _clouddirectoryDeleteSchema {
			clouddirectory_DeleteSchema(cfg, client)
			return
		}
		if _clouddirectoryDeleteTypedLinkFacet {
			clouddirectory_DeleteTypedLinkFacet(cfg, client)
			return
		}
		if _clouddirectoryDetachFromIndex {
			clouddirectory_DetachFromIndex(cfg, client)
			return
		}
		if _clouddirectoryDetachObject {
			clouddirectory_DetachObject(cfg, client)
			return
		}
		if _clouddirectoryDetachPolicy {
			clouddirectory_DetachPolicy(cfg, client)
			return
		}
		if _clouddirectoryDetachTypedLink {
			clouddirectory_DetachTypedLink(cfg, client)
			return
		}
		if _clouddirectoryDisableDirectory {
			clouddirectory_DisableDirectory(cfg, client)
			return
		}
		if _clouddirectoryEnableDirectory {
			clouddirectory_EnableDirectory(cfg, client)
			return
		}
		if _clouddirectoryGetAppliedSchemaVersion {
			clouddirectory_GetAppliedSchemaVersion(cfg, client)
			return
		}
		if _clouddirectoryGetDirectory {
			clouddirectory_GetDirectory(cfg, client)
			return
		}
		if _clouddirectoryGetFacet {
			clouddirectory_GetFacet(cfg, client)
			return
		}
		if _clouddirectoryGetLinkAttributes {
			clouddirectory_GetLinkAttributes(cfg, client)
			return
		}
		if _clouddirectoryGetObjectAttributes {
			clouddirectory_GetObjectAttributes(cfg, client)
			return
		}
		if _clouddirectoryGetObjectInformation {
			clouddirectory_GetObjectInformation(cfg, client)
			return
		}
		if _clouddirectoryGetSchemaAsJson {
			clouddirectory_GetSchemaAsJson(cfg, client)
			return
		}
		if _clouddirectoryGetTypedLinkFacetInformation {
			clouddirectory_GetTypedLinkFacetInformation(cfg, client)
			return
		}
		if _clouddirectoryListAppliedSchemaArns {
			clouddirectory_ListAppliedSchemaArns(cfg, client)
			return
		}
		if _clouddirectoryListAttachedIndices {
			clouddirectory_ListAttachedIndices(cfg, client)
			return
		}
		if _clouddirectoryListDevelopmentSchemaArns {
			clouddirectory_ListDevelopmentSchemaArns(cfg, client)
			return
		}
		if _clouddirectoryListDirectories {
			clouddirectory_ListDirectories(cfg, client)
			return
		}
		if _clouddirectoryListFacetAttributes {
			clouddirectory_ListFacetAttributes(cfg, client)
			return
		}
		if _clouddirectoryListFacetNames {
			clouddirectory_ListFacetNames(cfg, client)
			return
		}
		if _clouddirectoryListIncomingTypedLinks {
			clouddirectory_ListIncomingTypedLinks(cfg, client)
			return
		}
		if _clouddirectoryListIndex {
			clouddirectory_ListIndex(cfg, client)
			return
		}
		if _clouddirectoryListManagedSchemaArns {
			clouddirectory_ListManagedSchemaArns(cfg, client)
			return
		}
		if _clouddirectoryListObjectAttributes {
			clouddirectory_ListObjectAttributes(cfg, client)
			return
		}
		if _clouddirectoryListObjectChildren {
			clouddirectory_ListObjectChildren(cfg, client)
			return
		}
		if _clouddirectoryListObjectParentPaths {
			clouddirectory_ListObjectParentPaths(cfg, client)
			return
		}
		if _clouddirectoryListObjectParents {
			clouddirectory_ListObjectParents(cfg, client)
			return
		}
		if _clouddirectoryListObjectPolicies {
			clouddirectory_ListObjectPolicies(cfg, client)
			return
		}
		if _clouddirectoryListOutgoingTypedLinks {
			clouddirectory_ListOutgoingTypedLinks(cfg, client)
			return
		}
		if _clouddirectoryListPolicyAttachments {
			clouddirectory_ListPolicyAttachments(cfg, client)
			return
		}
		if _clouddirectoryListPublishedSchemaArns {
			clouddirectory_ListPublishedSchemaArns(cfg, client)
			return
		}
		if _clouddirectoryListTagsForResource {
			clouddirectory_ListTagsForResource(cfg, client)
			return
		}
		if _clouddirectoryListTypedLinkFacetAttributes {
			clouddirectory_ListTypedLinkFacetAttributes(cfg, client)
			return
		}
		if _clouddirectoryListTypedLinkFacetNames {
			clouddirectory_ListTypedLinkFacetNames(cfg, client)
			return
		}
		if _clouddirectoryLookupPolicy {
			clouddirectory_LookupPolicy(cfg, client)
			return
		}
		if _clouddirectoryPublishSchema {
			clouddirectory_PublishSchema(cfg, client)
			return
		}
		if _clouddirectoryPutSchemaFromJson {
			clouddirectory_PutSchemaFromJson(cfg, client)
			return
		}
		if _clouddirectoryRemoveFacetFromObject {
			clouddirectory_RemoveFacetFromObject(cfg, client)
			return
		}
		if _clouddirectoryTagResource {
			clouddirectory_TagResource(cfg, client)
			return
		}
		if _clouddirectoryUntagResource {
			clouddirectory_UntagResource(cfg, client)
			return
		}
		if _clouddirectoryUpdateFacet {
			clouddirectory_UpdateFacet(cfg, client)
			return
		}
		if _clouddirectoryUpdateLinkAttributes {
			clouddirectory_UpdateLinkAttributes(cfg, client)
			return
		}
		if _clouddirectoryUpdateObjectAttributes {
			clouddirectory_UpdateObjectAttributes(cfg, client)
			return
		}
		if _clouddirectoryUpdateSchema {
			clouddirectory_UpdateSchema(cfg, client)
			return
		}
		if _clouddirectoryUpdateTypedLinkFacet {
			clouddirectory_UpdateTypedLinkFacet(cfg, client)
			return
		}
		if _clouddirectoryUpgradeAppliedSchema {
			clouddirectory_UpgradeAppliedSchema(cfg, client)
			return
		}
		if _clouddirectoryUpgradePublishedSchema {
			clouddirectory_UpgradePublishedSchema(cfg, client)
			return
		}

	},
}

var (
	_clouddirectoryAddFacetToObject             bool
	_clouddirectoryApplySchema                  bool
	_clouddirectoryAttachObject                 bool
	_clouddirectoryAttachPolicy                 bool
	_clouddirectoryAttachToIndex                bool
	_clouddirectoryAttachTypedLink              bool
	_clouddirectoryBatchRead                    bool
	_clouddirectoryBatchWrite                   bool
	_clouddirectoryCreateDirectory              bool
	_clouddirectoryCreateFacet                  bool
	_clouddirectoryCreateIndex                  bool
	_clouddirectoryCreateObject                 bool
	_clouddirectoryCreateSchema                 bool
	_clouddirectoryCreateTypedLinkFacet         bool
	_clouddirectoryDeleteDirectory              bool
	_clouddirectoryDeleteFacet                  bool
	_clouddirectoryDeleteObject                 bool
	_clouddirectoryDeleteSchema                 bool
	_clouddirectoryDeleteTypedLinkFacet         bool
	_clouddirectoryDetachFromIndex              bool
	_clouddirectoryDetachObject                 bool
	_clouddirectoryDetachPolicy                 bool
	_clouddirectoryDetachTypedLink              bool
	_clouddirectoryDisableDirectory             bool
	_clouddirectoryEnableDirectory              bool
	_clouddirectoryGetAppliedSchemaVersion      bool
	_clouddirectoryGetDirectory                 bool
	_clouddirectoryGetFacet                     bool
	_clouddirectoryGetLinkAttributes            bool
	_clouddirectoryGetObjectAttributes          bool
	_clouddirectoryGetObjectInformation         bool
	_clouddirectoryGetSchemaAsJson              bool
	_clouddirectoryGetTypedLinkFacetInformation bool
	_clouddirectoryListAppliedSchemaArns        bool
	_clouddirectoryListAttachedIndices          bool
	_clouddirectoryListDevelopmentSchemaArns    bool
	_clouddirectoryListDirectories              bool
	_clouddirectoryListFacetAttributes          bool
	_clouddirectoryListFacetNames               bool
	_clouddirectoryListIncomingTypedLinks       bool
	_clouddirectoryListIndex                    bool
	_clouddirectoryListManagedSchemaArns        bool
	_clouddirectoryListObjectAttributes         bool
	_clouddirectoryListObjectChildren           bool
	_clouddirectoryListObjectParentPaths        bool
	_clouddirectoryListObjectParents            bool
	_clouddirectoryListObjectPolicies           bool
	_clouddirectoryListOutgoingTypedLinks       bool
	_clouddirectoryListPolicyAttachments        bool
	_clouddirectoryListPublishedSchemaArns      bool
	_clouddirectoryListTagsForResource          bool
	_clouddirectoryListTypedLinkFacetAttributes bool
	_clouddirectoryListTypedLinkFacetNames      bool
	_clouddirectoryLookupPolicy                 bool
	_clouddirectoryPublishSchema                bool
	_clouddirectoryPutSchemaFromJson            bool
	_clouddirectoryRemoveFacetFromObject        bool
	_clouddirectoryTagResource                  bool
	_clouddirectoryUntagResource                bool
	_clouddirectoryUpdateFacet                  bool
	_clouddirectoryUpdateLinkAttributes         bool
	_clouddirectoryUpdateObjectAttributes       bool
	_clouddirectoryUpdateSchema                 bool
	_clouddirectoryUpdateTypedLinkFacet         bool
	_clouddirectoryUpgradeAppliedSchema         bool
	_clouddirectoryUpgradePublishedSchema       bool

	_clouddirectoryAttributeNames              []string
	_clouddirectoryAttributeUpdates            string
	_clouddirectoryAttributes                  string
	_clouddirectoryChildReference              string
	_clouddirectoryConsistencyLevel            string
	_clouddirectoryDevelopmentSchemaArn        string
	_clouddirectoryDirectoryArn                string
	_clouddirectoryDocument                    string
	_clouddirectoryDryRun                      string
	_clouddirectoryFacet                       string
	_clouddirectoryFacetFilter                 string
	_clouddirectoryFacetStyle                  string
	_clouddirectoryFilterAttributeRanges       string
	_clouddirectoryFilterTypedLink             string
	_clouddirectoryIdentityAttributeOrder      []string
	_clouddirectoryIncludeAllLinksToEachParent string
	_clouddirectoryIndexReference              string
	_clouddirectoryIsUnique                    string
	_clouddirectoryLinkName                    string
	_clouddirectoryMaxResults                  string
	_clouddirectoryMinorVersion                string
	_clouddirectoryName                        string
	_clouddirectoryNextToken                   string
	_clouddirectoryObjectAttributeList         string
	_clouddirectoryObjectReference             string
	_clouddirectoryObjectType                  string
	_clouddirectoryOperations                  string
	_clouddirectoryOrderedIndexedAttributeList string
	_clouddirectoryParentReference             string
	_clouddirectoryPolicyReference             string
	_clouddirectoryPublishedSchemaArn          string
	_clouddirectoryRangesOnIndexedValues       string
	_clouddirectoryResourceArn                 string
	_clouddirectorySchemaArn                   string
	_clouddirectorySchemaFacet                 string
	_clouddirectorySchemaFacets                string
	_clouddirectorySourceObjectReference       string
	_clouddirectoryState                       string
	_clouddirectoryTagKeys                     []string
	_clouddirectoryTags                        string
	_clouddirectoryTargetObjectReference       string
	_clouddirectoryTargetReference             string
	_clouddirectoryTypedLinkFacet              string
	_clouddirectoryTypedLinkSpecifier          string
	_clouddirectoryVersion                     string
)

// Adds a new Facet to an object. An object can have more than one facet applied on it.
func clouddirectory_AddFacetToObject(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.AddFacetToObjectInput{
		// DirectoryArn: *string, // Required
		// ObjectReference: *types.ObjectReference, // Required
		// SchemaFacet: *types.SchemaFacet, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryObjectReference) > 0 {
		if err := assignInputField(input, "ObjectReference", _clouddirectoryObjectReference); err != nil {
			log.Errorf("invalid --object-reference: %s", err.Error())
			return
		}
	}
	if len(_clouddirectorySchemaFacet) > 0 {
		if err := assignInputField(input, "SchemaFacet", _clouddirectorySchemaFacet); err != nil {
			log.Errorf("invalid --schema-facet: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryObjectAttributeList) > 0 {
		if err := assignInputField(input, "ObjectAttributeList", _clouddirectoryObjectAttributeList); err != nil {
			log.Errorf("invalid --object-attribute-list: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddFacetToObject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Copies the input published schema, at the specified version, into the Directory with the
// same name and version as that of the published schema.
func clouddirectory_ApplySchema(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.ApplySchemaInput{
		// DirectoryArn: *string, // Required
		// PublishedSchemaArn: *string, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryPublishedSchemaArn) > 0 {
		input.PublishedSchemaArn = aws.String(_clouddirectoryPublishedSchemaArn)
	}

	if resp, err := client.ApplySchema(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches an existing object to another object. An object can be accessed in two
// ways:
//
// - Using the path
//
// - Using ObjectIdentifier
func clouddirectory_AttachObject(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.AttachObjectInput{
		// ChildReference: *types.ObjectReference, // Required
		// DirectoryArn: *string, // Required
		// LinkName: *string, // Required
		// ParentReference: *types.ObjectReference, // Required
	}

	if len(_clouddirectoryChildReference) > 0 {
		if err := assignInputField(input, "ChildReference", _clouddirectoryChildReference); err != nil {
			log.Errorf("invalid --child-reference: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryLinkName) > 0 {
		input.LinkName = aws.String(_clouddirectoryLinkName)
	}
	if len(_clouddirectoryParentReference) > 0 {
		if err := assignInputField(input, "ParentReference", _clouddirectoryParentReference); err != nil {
			log.Errorf("invalid --parent-reference: %s", err.Error())
			return
		}
	}

	if resp, err := client.AttachObject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches a policy object to a regular object. An object can have a limited
// number of attached policies.
func clouddirectory_AttachPolicy(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.AttachPolicyInput{
		// DirectoryArn: *string, // Required
		// ObjectReference: *types.ObjectReference, // Required
		// PolicyReference: *types.ObjectReference, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryObjectReference) > 0 {
		if err := assignInputField(input, "ObjectReference", _clouddirectoryObjectReference); err != nil {
			log.Errorf("invalid --object-reference: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryPolicyReference) > 0 {
		if err := assignInputField(input, "PolicyReference", _clouddirectoryPolicyReference); err != nil {
			log.Errorf("invalid --policy-reference: %s", err.Error())
			return
		}
	}

	if resp, err := client.AttachPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches the specified object to the specified index.
func clouddirectory_AttachToIndex(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.AttachToIndexInput{
		// DirectoryArn: *string, // Required
		// IndexReference: *types.ObjectReference, // Required
		// TargetReference: *types.ObjectReference, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryIndexReference) > 0 {
		if err := assignInputField(input, "IndexReference", _clouddirectoryIndexReference); err != nil {
			log.Errorf("invalid --index-reference: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryTargetReference) > 0 {
		if err := assignInputField(input, "TargetReference", _clouddirectoryTargetReference); err != nil {
			log.Errorf("invalid --target-reference: %s", err.Error())
			return
		}
	}

	if resp, err := client.AttachToIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches a typed link to a specified source and target object. For more
// information, see [Typed Links].
//
// [Typed Links]: https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink
func clouddirectory_AttachTypedLink(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.AttachTypedLinkInput{
		// Attributes: []types.AttributeNameAndValue, // Required
		// DirectoryArn: *string, // Required
		// SourceObjectReference: *types.ObjectReference, // Required
		// TargetObjectReference: *types.ObjectReference, // Required
		// TypedLinkFacet: *types.TypedLinkSchemaAndFacetName, // Required
	}

	if len(_clouddirectoryAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _clouddirectoryAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectorySourceObjectReference) > 0 {
		if err := assignInputField(input, "SourceObjectReference", _clouddirectorySourceObjectReference); err != nil {
			log.Errorf("invalid --source-object-reference: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryTargetObjectReference) > 0 {
		if err := assignInputField(input, "TargetObjectReference", _clouddirectoryTargetObjectReference); err != nil {
			log.Errorf("invalid --target-object-reference: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryTypedLinkFacet) > 0 {
		if err := assignInputField(input, "TypedLinkFacet", _clouddirectoryTypedLinkFacet); err != nil {
			log.Errorf("invalid --typed-link-facet: %s", err.Error())
			return
		}
	}

	if resp, err := client.AttachTypedLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Performs all the read operations in a batch.
func clouddirectory_BatchRead(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.BatchReadInput{
		// DirectoryArn: *string, // Required
		// Operations: []types.BatchReadOperation, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryOperations) > 0 {
		if err := assignInputField(input, "Operations", _clouddirectoryOperations); err != nil {
			log.Errorf("invalid --operations: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryConsistencyLevel) > 0 {
		if err := assignInputField(input, "ConsistencyLevel", _clouddirectoryConsistencyLevel); err != nil {
			log.Errorf("invalid --consistency-level: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchRead(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Performs all the write operations in a batch. Either all the operations succeed
// or none.
func clouddirectory_BatchWrite(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.BatchWriteInput{
		// DirectoryArn: *string, // Required
		// Operations: []types.BatchWriteOperation, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryOperations) > 0 {
		if err := assignInputField(input, "Operations", _clouddirectoryOperations); err != nil {
			log.Errorf("invalid --operations: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchWrite(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Directory by copying the published schema into the directory. A directory
// cannot be created without a schema.
//
// You can also quickly create a directory using a managed schema, called the
// QuickStartSchema . For more information, see [Managed Schema] in the Amazon Cloud Directory
// Developer Guide.
//
// [Managed Schema]: https://docs.aws.amazon.com/clouddirectory/latest/developerguide/schemas_managed.html
func clouddirectory_CreateDirectory(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.CreateDirectoryInput{
		// Name: *string, // Required
		// SchemaArn: *string, // Required
	}

	if len(_clouddirectoryName) > 0 {
		input.Name = aws.String(_clouddirectoryName)
	}
	if len(_clouddirectorySchemaArn) > 0 {
		input.SchemaArn = aws.String(_clouddirectorySchemaArn)
	}

	if resp, err := client.CreateDirectory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Facet in a schema. Facet creation is allowed only in development or
// applied schemas.
func clouddirectory_CreateFacet(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.CreateFacetInput{
		// Name: *string, // Required
		// SchemaArn: *string, // Required
	}

	if len(_clouddirectoryName) > 0 {
		input.Name = aws.String(_clouddirectoryName)
	}
	if len(_clouddirectorySchemaArn) > 0 {
		input.SchemaArn = aws.String(_clouddirectorySchemaArn)
	}
	if len(_clouddirectoryAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _clouddirectoryAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryFacetStyle) > 0 {
		if err := assignInputField(input, "FacetStyle", _clouddirectoryFacetStyle); err != nil {
			log.Errorf("invalid --facet-style: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryObjectType) > 0 {
		if err := assignInputField(input, "ObjectType", _clouddirectoryObjectType); err != nil {
			log.Errorf("invalid --object-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFacet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an index object. See [Indexing and search] for more information.
//
// [Indexing and search]: https://docs.aws.amazon.com/clouddirectory/latest/developerguide/indexing_search.html
func clouddirectory_CreateIndex(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.CreateIndexInput{
		// DirectoryArn: *string, // Required
		// IsUnique: bool, // Required
		// OrderedIndexedAttributeList: []types.AttributeKey, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryIsUnique) > 0 {
		if err := assignInputField(input, "IsUnique", _clouddirectoryIsUnique); err != nil {
			log.Errorf("invalid --is-unique: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryOrderedIndexedAttributeList) > 0 {
		if err := assignInputField(input, "OrderedIndexedAttributeList", _clouddirectoryOrderedIndexedAttributeList); err != nil {
			log.Errorf("invalid --ordered-indexed-attribute-list: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryLinkName) > 0 {
		input.LinkName = aws.String(_clouddirectoryLinkName)
	}
	if len(_clouddirectoryParentReference) > 0 {
		if err := assignInputField(input, "ParentReference", _clouddirectoryParentReference); err != nil {
			log.Errorf("invalid --parent-reference: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an object in a Directory. Additionally attaches the object to a parent, if a
// parent reference and LinkName is specified. An object is simply a collection of Facet
// attributes. You can also use this API call to create a policy object, if the
// facet from which you create the object is a policy facet.
func clouddirectory_CreateObject(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.CreateObjectInput{
		// DirectoryArn: *string, // Required
		// SchemaFacets: []types.SchemaFacet, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectorySchemaFacets) > 0 {
		if err := assignInputField(input, "SchemaFacets", _clouddirectorySchemaFacets); err != nil {
			log.Errorf("invalid --schema-facets: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryLinkName) > 0 {
		input.LinkName = aws.String(_clouddirectoryLinkName)
	}
	if len(_clouddirectoryObjectAttributeList) > 0 {
		if err := assignInputField(input, "ObjectAttributeList", _clouddirectoryObjectAttributeList); err != nil {
			log.Errorf("invalid --object-attribute-list: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryParentReference) > 0 {
		if err := assignInputField(input, "ParentReference", _clouddirectoryParentReference); err != nil {
			log.Errorf("invalid --parent-reference: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateObject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new schema in a development state. A schema can exist in three phases:
// - Development: This is a mutable phase of the schema. All new schemas are in
// the development phase. Once the schema is finalized, it can be published.
//
// - Published: Published schemas are immutable and have a version associated
// with them.
//
// - Applied: Applied schemas are mutable in a way that allows you to add new
// schema facets. You can also add new, nonrequired attributes to existing schema
// facets. You can apply only published schemas to directories.
func clouddirectory_CreateSchema(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.CreateSchemaInput{
		// Name: *string, // Required
	}

	if len(_clouddirectoryName) > 0 {
		input.Name = aws.String(_clouddirectoryName)
	}

	if resp, err := client.CreateSchema(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a TypedLinkFacet. For more information, see [Typed Links].
//
// [Typed Links]: https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink
func clouddirectory_CreateTypedLinkFacet(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.CreateTypedLinkFacetInput{
		// Facet: *types.TypedLinkFacet, // Required
		// SchemaArn: *string, // Required
	}

	if len(_clouddirectoryFacet) > 0 {
		if err := assignInputField(input, "Facet", _clouddirectoryFacet); err != nil {
			log.Errorf("invalid --facet: %s", err.Error())
			return
		}
	}
	if len(_clouddirectorySchemaArn) > 0 {
		input.SchemaArn = aws.String(_clouddirectorySchemaArn)
	}

	if resp, err := client.CreateTypedLinkFacet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a directory. Only disabled directories can be deleted. A deleted
// directory cannot be undone. Exercise extreme caution when deleting directories.
func clouddirectory_DeleteDirectory(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.DeleteDirectoryInput{
		// DirectoryArn: *string, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}

	if resp, err := client.DeleteDirectory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a given Facet. All attributes and Rules that are associated with the facet will
// be deleted. Only development schema facets are allowed deletion.
func clouddirectory_DeleteFacet(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.DeleteFacetInput{
		// Name: *string, // Required
		// SchemaArn: *string, // Required
	}

	if len(_clouddirectoryName) > 0 {
		input.Name = aws.String(_clouddirectoryName)
	}
	if len(_clouddirectorySchemaArn) > 0 {
		input.SchemaArn = aws.String(_clouddirectorySchemaArn)
	}

	if resp, err := client.DeleteFacet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an object and its associated attributes. Only objects with no children
// and no parents can be deleted. The maximum number of attributes that can be
// deleted during an object deletion is 30. For more information, see [Amazon Cloud Directory Limits].
//
// [Amazon Cloud Directory Limits]: https://docs.aws.amazon.com/clouddirectory/latest/developerguide/limits.html
func clouddirectory_DeleteObject(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.DeleteObjectInput{
		// DirectoryArn: *string, // Required
		// ObjectReference: *types.ObjectReference, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryObjectReference) > 0 {
		if err := assignInputField(input, "ObjectReference", _clouddirectoryObjectReference); err != nil {
			log.Errorf("invalid --object-reference: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteObject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a given schema. Schemas in a development and published state can only
// be deleted.
func clouddirectory_DeleteSchema(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.DeleteSchemaInput{
		// SchemaArn: *string, // Required
	}

	if len(_clouddirectorySchemaArn) > 0 {
		input.SchemaArn = aws.String(_clouddirectorySchemaArn)
	}

	if resp, err := client.DeleteSchema(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a TypedLinkFacet. For more information, see [Typed Links].
//
// [Typed Links]: https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink
func clouddirectory_DeleteTypedLinkFacet(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.DeleteTypedLinkFacetInput{
		// Name: *string, // Required
		// SchemaArn: *string, // Required
	}

	if len(_clouddirectoryName) > 0 {
		input.Name = aws.String(_clouddirectoryName)
	}
	if len(_clouddirectorySchemaArn) > 0 {
		input.SchemaArn = aws.String(_clouddirectorySchemaArn)
	}

	if resp, err := client.DeleteTypedLinkFacet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detaches the specified object from the specified index.
func clouddirectory_DetachFromIndex(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.DetachFromIndexInput{
		// DirectoryArn: *string, // Required
		// IndexReference: *types.ObjectReference, // Required
		// TargetReference: *types.ObjectReference, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryIndexReference) > 0 {
		if err := assignInputField(input, "IndexReference", _clouddirectoryIndexReference); err != nil {
			log.Errorf("invalid --index-reference: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryTargetReference) > 0 {
		if err := assignInputField(input, "TargetReference", _clouddirectoryTargetReference); err != nil {
			log.Errorf("invalid --target-reference: %s", err.Error())
			return
		}
	}

	if resp, err := client.DetachFromIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detaches a given object from the parent object. The object that is to be
// detached from the parent is specified by the link name.
func clouddirectory_DetachObject(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.DetachObjectInput{
		// DirectoryArn: *string, // Required
		// LinkName: *string, // Required
		// ParentReference: *types.ObjectReference, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryLinkName) > 0 {
		input.LinkName = aws.String(_clouddirectoryLinkName)
	}
	if len(_clouddirectoryParentReference) > 0 {
		if err := assignInputField(input, "ParentReference", _clouddirectoryParentReference); err != nil {
			log.Errorf("invalid --parent-reference: %s", err.Error())
			return
		}
	}

	if resp, err := client.DetachObject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detaches a policy from an object.
func clouddirectory_DetachPolicy(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.DetachPolicyInput{
		// DirectoryArn: *string, // Required
		// ObjectReference: *types.ObjectReference, // Required
		// PolicyReference: *types.ObjectReference, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryObjectReference) > 0 {
		if err := assignInputField(input, "ObjectReference", _clouddirectoryObjectReference); err != nil {
			log.Errorf("invalid --object-reference: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryPolicyReference) > 0 {
		if err := assignInputField(input, "PolicyReference", _clouddirectoryPolicyReference); err != nil {
			log.Errorf("invalid --policy-reference: %s", err.Error())
			return
		}
	}

	if resp, err := client.DetachPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detaches a typed link from a specified source and target object. For more
// information, see [Typed Links].
//
// [Typed Links]: https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink
func clouddirectory_DetachTypedLink(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.DetachTypedLinkInput{
		// DirectoryArn: *string, // Required
		// TypedLinkSpecifier: *types.TypedLinkSpecifier, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryTypedLinkSpecifier) > 0 {
		if err := assignInputField(input, "TypedLinkSpecifier", _clouddirectoryTypedLinkSpecifier); err != nil {
			log.Errorf("invalid --typed-link-specifier: %s", err.Error())
			return
		}
	}

	if resp, err := client.DetachTypedLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables the specified directory. Disabled directories cannot be read or
// written to. Only enabled directories can be disabled. Disabled directories may
// be reenabled.
func clouddirectory_DisableDirectory(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.DisableDirectoryInput{
		// DirectoryArn: *string, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}

	if resp, err := client.DisableDirectory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the specified directory. Only disabled directories can be enabled. Once
// enabled, the directory can then be read and written to.
func clouddirectory_EnableDirectory(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.EnableDirectoryInput{
		// DirectoryArn: *string, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}

	if resp, err := client.EnableDirectory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns current applied schema version ARN, including the minor version in use.
func clouddirectory_GetAppliedSchemaVersion(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.GetAppliedSchemaVersionInput{
		// SchemaArn: *string, // Required
	}

	if len(_clouddirectorySchemaArn) > 0 {
		input.SchemaArn = aws.String(_clouddirectorySchemaArn)
	}

	if resp, err := client.GetAppliedSchemaVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves metadata about a directory.
func clouddirectory_GetDirectory(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.GetDirectoryInput{
		// DirectoryArn: *string, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}

	if resp, err := client.GetDirectory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details of the Facet, such as facet name, attributes, Rules, or ObjectType . You can
// call this on all kinds of schema facets -- published, development, or applied.
func clouddirectory_GetFacet(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.GetFacetInput{
		// Name: *string, // Required
		// SchemaArn: *string, // Required
	}

	if len(_clouddirectoryName) > 0 {
		input.Name = aws.String(_clouddirectoryName)
	}
	if len(_clouddirectorySchemaArn) > 0 {
		input.SchemaArn = aws.String(_clouddirectorySchemaArn)
	}

	if resp, err := client.GetFacet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves attributes that are associated with a typed link.
func clouddirectory_GetLinkAttributes(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.GetLinkAttributesInput{
		// AttributeNames: []string, // Required
		// DirectoryArn: *string, // Required
		// TypedLinkSpecifier: *types.TypedLinkSpecifier, // Required
	}

	if len(_clouddirectoryAttributeNames) > 0 {
		input.AttributeNames = append([]string(nil), _clouddirectoryAttributeNames...)
	}
	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryTypedLinkSpecifier) > 0 {
		if err := assignInputField(input, "TypedLinkSpecifier", _clouddirectoryTypedLinkSpecifier); err != nil {
			log.Errorf("invalid --typed-link-specifier: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryConsistencyLevel) > 0 {
		if err := assignInputField(input, "ConsistencyLevel", _clouddirectoryConsistencyLevel); err != nil {
			log.Errorf("invalid --consistency-level: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetLinkAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves attributes within a facet that are associated with an object.
func clouddirectory_GetObjectAttributes(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.GetObjectAttributesInput{
		// AttributeNames: []string, // Required
		// DirectoryArn: *string, // Required
		// ObjectReference: *types.ObjectReference, // Required
		// SchemaFacet: *types.SchemaFacet, // Required
	}

	if len(_clouddirectoryAttributeNames) > 0 {
		input.AttributeNames = append([]string(nil), _clouddirectoryAttributeNames...)
	}
	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryObjectReference) > 0 {
		if err := assignInputField(input, "ObjectReference", _clouddirectoryObjectReference); err != nil {
			log.Errorf("invalid --object-reference: %s", err.Error())
			return
		}
	}
	if len(_clouddirectorySchemaFacet) > 0 {
		if err := assignInputField(input, "SchemaFacet", _clouddirectorySchemaFacet); err != nil {
			log.Errorf("invalid --schema-facet: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryConsistencyLevel) > 0 {
		if err := assignInputField(input, "ConsistencyLevel", _clouddirectoryConsistencyLevel); err != nil {
			log.Errorf("invalid --consistency-level: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetObjectAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves metadata about an object.
func clouddirectory_GetObjectInformation(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.GetObjectInformationInput{
		// DirectoryArn: *string, // Required
		// ObjectReference: *types.ObjectReference, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryObjectReference) > 0 {
		if err := assignInputField(input, "ObjectReference", _clouddirectoryObjectReference); err != nil {
			log.Errorf("invalid --object-reference: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryConsistencyLevel) > 0 {
		if err := assignInputField(input, "ConsistencyLevel", _clouddirectoryConsistencyLevel); err != nil {
			log.Errorf("invalid --consistency-level: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetObjectInformation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a JSON representation of the schema. See [JSON Schema Format] for more information.
//
// [JSON Schema Format]: https://docs.aws.amazon.com/clouddirectory/latest/developerguide/schemas_jsonformat.html#schemas_json
func clouddirectory_GetSchemaAsJson(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.GetSchemaAsJsonInput{
		// SchemaArn: *string, // Required
	}

	if len(_clouddirectorySchemaArn) > 0 {
		input.SchemaArn = aws.String(_clouddirectorySchemaArn)
	}

	if resp, err := client.GetSchemaAsJson(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the identity attribute order for a specific TypedLinkFacet. For more information, see [Typed Links]
// .
//
// [Typed Links]: https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink
func clouddirectory_GetTypedLinkFacetInformation(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.GetTypedLinkFacetInformationInput{
		// Name: *string, // Required
		// SchemaArn: *string, // Required
	}

	if len(_clouddirectoryName) > 0 {
		input.Name = aws.String(_clouddirectoryName)
	}
	if len(_clouddirectorySchemaArn) > 0 {
		input.SchemaArn = aws.String(_clouddirectorySchemaArn)
	}

	if resp, err := client.GetTypedLinkFacetInformation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists schema major versions applied to a directory. If SchemaArn is provided,
// lists the minor version.
func clouddirectory_ListAppliedSchemaArns(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.ListAppliedSchemaArnsInput{
		// DirectoryArn: *string, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _clouddirectoryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryNextToken) > 0 {
		input.NextToken = aws.String(_clouddirectoryNextToken)
	}
	if len(_clouddirectorySchemaArn) > 0 {
		input.SchemaArn = aws.String(_clouddirectorySchemaArn)
	}

	if disablePaginator() {
		if resp, err := client.ListAppliedSchemaArns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*clouddirectory.ListAppliedSchemaArnsOutput
	p := clouddirectory.NewListAppliedSchemaArnsPaginator(client, input)
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

// Lists indices attached to the specified object.
func clouddirectory_ListAttachedIndices(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.ListAttachedIndicesInput{
		// DirectoryArn: *string, // Required
		// TargetReference: *types.ObjectReference, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryTargetReference) > 0 {
		if err := assignInputField(input, "TargetReference", _clouddirectoryTargetReference); err != nil {
			log.Errorf("invalid --target-reference: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryConsistencyLevel) > 0 {
		if err := assignInputField(input, "ConsistencyLevel", _clouddirectoryConsistencyLevel); err != nil {
			log.Errorf("invalid --consistency-level: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _clouddirectoryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryNextToken) > 0 {
		input.NextToken = aws.String(_clouddirectoryNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAttachedIndices(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*clouddirectory.ListAttachedIndicesOutput
	p := clouddirectory.NewListAttachedIndicesPaginator(client, input)
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

// Retrieves each Amazon Resource Name (ARN) of schemas in the development state.
func clouddirectory_ListDevelopmentSchemaArns(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.ListDevelopmentSchemaArnsInput{}

	if len(_clouddirectoryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _clouddirectoryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryNextToken) > 0 {
		input.NextToken = aws.String(_clouddirectoryNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDevelopmentSchemaArns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*clouddirectory.ListDevelopmentSchemaArnsOutput
	p := clouddirectory.NewListDevelopmentSchemaArnsPaginator(client, input)
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

// Lists directories created within an account.
func clouddirectory_ListDirectories(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.ListDirectoriesInput{}

	if len(_clouddirectoryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _clouddirectoryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryNextToken) > 0 {
		input.NextToken = aws.String(_clouddirectoryNextToken)
	}
	if len(_clouddirectoryState) > 0 {
		if err := assignInputField(input, "State", _clouddirectoryState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDirectories(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*clouddirectory.ListDirectoriesOutput
	p := clouddirectory.NewListDirectoriesPaginator(client, input)
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

// Retrieves attributes attached to the facet.
func clouddirectory_ListFacetAttributes(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.ListFacetAttributesInput{
		// Name: *string, // Required
		// SchemaArn: *string, // Required
	}

	if len(_clouddirectoryName) > 0 {
		input.Name = aws.String(_clouddirectoryName)
	}
	if len(_clouddirectorySchemaArn) > 0 {
		input.SchemaArn = aws.String(_clouddirectorySchemaArn)
	}
	if len(_clouddirectoryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _clouddirectoryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryNextToken) > 0 {
		input.NextToken = aws.String(_clouddirectoryNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFacetAttributes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*clouddirectory.ListFacetAttributesOutput
	p := clouddirectory.NewListFacetAttributesPaginator(client, input)
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

// Retrieves the names of facets that exist in a schema.
func clouddirectory_ListFacetNames(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.ListFacetNamesInput{
		// SchemaArn: *string, // Required
	}

	if len(_clouddirectorySchemaArn) > 0 {
		input.SchemaArn = aws.String(_clouddirectorySchemaArn)
	}
	if len(_clouddirectoryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _clouddirectoryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryNextToken) > 0 {
		input.NextToken = aws.String(_clouddirectoryNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFacetNames(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*clouddirectory.ListFacetNamesOutput
	p := clouddirectory.NewListFacetNamesPaginator(client, input)
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

// Returns a paginated list of all the incoming TypedLinkSpecifier information for an object. It
// also supports filtering by typed link facet and identity attributes. For more
// information, see [Typed Links].
//
// [Typed Links]: https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink
func clouddirectory_ListIncomingTypedLinks(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.ListIncomingTypedLinksInput{
		// DirectoryArn: *string, // Required
		// ObjectReference: *types.ObjectReference, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryObjectReference) > 0 {
		if err := assignInputField(input, "ObjectReference", _clouddirectoryObjectReference); err != nil {
			log.Errorf("invalid --object-reference: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryConsistencyLevel) > 0 {
		if err := assignInputField(input, "ConsistencyLevel", _clouddirectoryConsistencyLevel); err != nil {
			log.Errorf("invalid --consistency-level: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryFilterAttributeRanges) > 0 {
		if err := assignInputField(input, "FilterAttributeRanges", _clouddirectoryFilterAttributeRanges); err != nil {
			log.Errorf("invalid --filter-attribute-ranges: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryFilterTypedLink) > 0 {
		if err := assignInputField(input, "FilterTypedLink", _clouddirectoryFilterTypedLink); err != nil {
			log.Errorf("invalid --filter-typed-link: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _clouddirectoryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryNextToken) > 0 {
		input.NextToken = aws.String(_clouddirectoryNextToken)
	}

	if resp, err := client.ListIncomingTypedLinks(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists objects attached to the specified index.
func clouddirectory_ListIndex(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.ListIndexInput{
		// DirectoryArn: *string, // Required
		// IndexReference: *types.ObjectReference, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryIndexReference) > 0 {
		if err := assignInputField(input, "IndexReference", _clouddirectoryIndexReference); err != nil {
			log.Errorf("invalid --index-reference: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryConsistencyLevel) > 0 {
		if err := assignInputField(input, "ConsistencyLevel", _clouddirectoryConsistencyLevel); err != nil {
			log.Errorf("invalid --consistency-level: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _clouddirectoryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryNextToken) > 0 {
		input.NextToken = aws.String(_clouddirectoryNextToken)
	}
	if len(_clouddirectoryRangesOnIndexedValues) > 0 {
		if err := assignInputField(input, "RangesOnIndexedValues", _clouddirectoryRangesOnIndexedValues); err != nil {
			log.Errorf("invalid --ranges-on-indexed-values: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListIndex(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*clouddirectory.ListIndexOutput
	p := clouddirectory.NewListIndexPaginator(client, input)
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

// Lists the major version families of each managed schema. If a major version ARN
// is provided as SchemaArn, the minor version revisions in that family are listed
// instead.
func clouddirectory_ListManagedSchemaArns(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.ListManagedSchemaArnsInput{}

	if len(_clouddirectoryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _clouddirectoryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryNextToken) > 0 {
		input.NextToken = aws.String(_clouddirectoryNextToken)
	}
	if len(_clouddirectorySchemaArn) > 0 {
		input.SchemaArn = aws.String(_clouddirectorySchemaArn)
	}

	if disablePaginator() {
		if resp, err := client.ListManagedSchemaArns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*clouddirectory.ListManagedSchemaArnsOutput
	p := clouddirectory.NewListManagedSchemaArnsPaginator(client, input)
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

// Lists all attributes that are associated with an object.
func clouddirectory_ListObjectAttributes(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.ListObjectAttributesInput{
		// DirectoryArn: *string, // Required
		// ObjectReference: *types.ObjectReference, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryObjectReference) > 0 {
		if err := assignInputField(input, "ObjectReference", _clouddirectoryObjectReference); err != nil {
			log.Errorf("invalid --object-reference: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryConsistencyLevel) > 0 {
		if err := assignInputField(input, "ConsistencyLevel", _clouddirectoryConsistencyLevel); err != nil {
			log.Errorf("invalid --consistency-level: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryFacetFilter) > 0 {
		if err := assignInputField(input, "FacetFilter", _clouddirectoryFacetFilter); err != nil {
			log.Errorf("invalid --facet-filter: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _clouddirectoryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryNextToken) > 0 {
		input.NextToken = aws.String(_clouddirectoryNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListObjectAttributes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*clouddirectory.ListObjectAttributesOutput
	p := clouddirectory.NewListObjectAttributesPaginator(client, input)
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

// Returns a paginated list of child objects that are associated with a given
// object.
func clouddirectory_ListObjectChildren(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.ListObjectChildrenInput{
		// DirectoryArn: *string, // Required
		// ObjectReference: *types.ObjectReference, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryObjectReference) > 0 {
		if err := assignInputField(input, "ObjectReference", _clouddirectoryObjectReference); err != nil {
			log.Errorf("invalid --object-reference: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryConsistencyLevel) > 0 {
		if err := assignInputField(input, "ConsistencyLevel", _clouddirectoryConsistencyLevel); err != nil {
			log.Errorf("invalid --consistency-level: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _clouddirectoryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryNextToken) > 0 {
		input.NextToken = aws.String(_clouddirectoryNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListObjectChildren(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*clouddirectory.ListObjectChildrenOutput
	p := clouddirectory.NewListObjectChildrenPaginator(client, input)
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

// Retrieves all available parent paths for any object type such as node, leaf
// node, policy node, and index node objects. For more information about objects,
// see [Directory Structure].
//
// Use this API to evaluate all parents for an object. The call returns all
// objects from the root of the directory up to the requested object. The API
// returns the number of paths based on user-defined MaxResults , in case there are
// multiple paths to the parent. The order of the paths and nodes returned is
// consistent among multiple API calls unless the objects are deleted or moved.
// Paths not leading to the directory root are ignored from the target object.
//
// [Directory Structure]: https://docs.aws.amazon.com/clouddirectory/latest/developerguide/key_concepts_directorystructure.html
func clouddirectory_ListObjectParentPaths(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.ListObjectParentPathsInput{
		// DirectoryArn: *string, // Required
		// ObjectReference: *types.ObjectReference, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryObjectReference) > 0 {
		if err := assignInputField(input, "ObjectReference", _clouddirectoryObjectReference); err != nil {
			log.Errorf("invalid --object-reference: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _clouddirectoryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryNextToken) > 0 {
		input.NextToken = aws.String(_clouddirectoryNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListObjectParentPaths(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*clouddirectory.ListObjectParentPathsOutput
	p := clouddirectory.NewListObjectParentPathsPaginator(client, input)
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

// Lists parent objects that are associated with a given object in pagination
// fashion.
func clouddirectory_ListObjectParents(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.ListObjectParentsInput{
		// DirectoryArn: *string, // Required
		// ObjectReference: *types.ObjectReference, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryObjectReference) > 0 {
		if err := assignInputField(input, "ObjectReference", _clouddirectoryObjectReference); err != nil {
			log.Errorf("invalid --object-reference: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryConsistencyLevel) > 0 {
		if err := assignInputField(input, "ConsistencyLevel", _clouddirectoryConsistencyLevel); err != nil {
			log.Errorf("invalid --consistency-level: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryIncludeAllLinksToEachParent) > 0 {
		if err := assignInputField(input, "IncludeAllLinksToEachParent", _clouddirectoryIncludeAllLinksToEachParent); err != nil {
			log.Errorf("invalid --include-all-links-to-each-parent: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _clouddirectoryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryNextToken) > 0 {
		input.NextToken = aws.String(_clouddirectoryNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListObjectParents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*clouddirectory.ListObjectParentsOutput
	p := clouddirectory.NewListObjectParentsPaginator(client, input)
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

// Returns policies attached to an object in pagination fashion.
func clouddirectory_ListObjectPolicies(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.ListObjectPoliciesInput{
		// DirectoryArn: *string, // Required
		// ObjectReference: *types.ObjectReference, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryObjectReference) > 0 {
		if err := assignInputField(input, "ObjectReference", _clouddirectoryObjectReference); err != nil {
			log.Errorf("invalid --object-reference: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryConsistencyLevel) > 0 {
		if err := assignInputField(input, "ConsistencyLevel", _clouddirectoryConsistencyLevel); err != nil {
			log.Errorf("invalid --consistency-level: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _clouddirectoryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryNextToken) > 0 {
		input.NextToken = aws.String(_clouddirectoryNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListObjectPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*clouddirectory.ListObjectPoliciesOutput
	p := clouddirectory.NewListObjectPoliciesPaginator(client, input)
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

// Returns a paginated list of all the outgoing TypedLinkSpecifier information for an object. It
// also supports filtering by typed link facet and identity attributes. For more
// information, see [Typed Links].
//
// [Typed Links]: https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink
func clouddirectory_ListOutgoingTypedLinks(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.ListOutgoingTypedLinksInput{
		// DirectoryArn: *string, // Required
		// ObjectReference: *types.ObjectReference, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryObjectReference) > 0 {
		if err := assignInputField(input, "ObjectReference", _clouddirectoryObjectReference); err != nil {
			log.Errorf("invalid --object-reference: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryConsistencyLevel) > 0 {
		if err := assignInputField(input, "ConsistencyLevel", _clouddirectoryConsistencyLevel); err != nil {
			log.Errorf("invalid --consistency-level: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryFilterAttributeRanges) > 0 {
		if err := assignInputField(input, "FilterAttributeRanges", _clouddirectoryFilterAttributeRanges); err != nil {
			log.Errorf("invalid --filter-attribute-ranges: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryFilterTypedLink) > 0 {
		if err := assignInputField(input, "FilterTypedLink", _clouddirectoryFilterTypedLink); err != nil {
			log.Errorf("invalid --filter-typed-link: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _clouddirectoryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryNextToken) > 0 {
		input.NextToken = aws.String(_clouddirectoryNextToken)
	}

	if resp, err := client.ListOutgoingTypedLinks(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns all of the ObjectIdentifiers to which a given policy is attached.
func clouddirectory_ListPolicyAttachments(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.ListPolicyAttachmentsInput{
		// DirectoryArn: *string, // Required
		// PolicyReference: *types.ObjectReference, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryPolicyReference) > 0 {
		if err := assignInputField(input, "PolicyReference", _clouddirectoryPolicyReference); err != nil {
			log.Errorf("invalid --policy-reference: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryConsistencyLevel) > 0 {
		if err := assignInputField(input, "ConsistencyLevel", _clouddirectoryConsistencyLevel); err != nil {
			log.Errorf("invalid --consistency-level: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _clouddirectoryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryNextToken) > 0 {
		input.NextToken = aws.String(_clouddirectoryNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPolicyAttachments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*clouddirectory.ListPolicyAttachmentsOutput
	p := clouddirectory.NewListPolicyAttachmentsPaginator(client, input)
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

// Lists the major version families of each published schema. If a major version
// ARN is provided as SchemaArn , the minor version revisions in that family are
// listed instead.
func clouddirectory_ListPublishedSchemaArns(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.ListPublishedSchemaArnsInput{}

	if len(_clouddirectoryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _clouddirectoryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryNextToken) > 0 {
		input.NextToken = aws.String(_clouddirectoryNextToken)
	}
	if len(_clouddirectorySchemaArn) > 0 {
		input.SchemaArn = aws.String(_clouddirectorySchemaArn)
	}

	if disablePaginator() {
		if resp, err := client.ListPublishedSchemaArns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*clouddirectory.ListPublishedSchemaArnsOutput
	p := clouddirectory.NewListPublishedSchemaArnsPaginator(client, input)
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

// Returns tags for a resource. Tagging is currently supported only for
// directories with a limit of 50 tags per directory. All 50 tags are returned for
// a given directory with this API call.
func clouddirectory_ListTagsForResource(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_clouddirectoryResourceArn) > 0 {
		input.ResourceArn = aws.String(_clouddirectoryResourceArn)
	}
	if len(_clouddirectoryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _clouddirectoryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryNextToken) > 0 {
		input.NextToken = aws.String(_clouddirectoryNextToken)
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

	var results []*clouddirectory.ListTagsForResourceOutput
	p := clouddirectory.NewListTagsForResourcePaginator(client, input)
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

// Returns a paginated list of all attribute definitions for a particular TypedLinkFacet. For
// more information, see [Typed Links].
//
// [Typed Links]: https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink
func clouddirectory_ListTypedLinkFacetAttributes(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.ListTypedLinkFacetAttributesInput{
		// Name: *string, // Required
		// SchemaArn: *string, // Required
	}

	if len(_clouddirectoryName) > 0 {
		input.Name = aws.String(_clouddirectoryName)
	}
	if len(_clouddirectorySchemaArn) > 0 {
		input.SchemaArn = aws.String(_clouddirectorySchemaArn)
	}
	if len(_clouddirectoryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _clouddirectoryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryNextToken) > 0 {
		input.NextToken = aws.String(_clouddirectoryNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTypedLinkFacetAttributes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*clouddirectory.ListTypedLinkFacetAttributesOutput
	p := clouddirectory.NewListTypedLinkFacetAttributesPaginator(client, input)
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

// Returns a paginated list of TypedLink facet names for a particular schema. For
// more information, see [Typed Links].
//
// [Typed Links]: https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink
func clouddirectory_ListTypedLinkFacetNames(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.ListTypedLinkFacetNamesInput{
		// SchemaArn: *string, // Required
	}

	if len(_clouddirectorySchemaArn) > 0 {
		input.SchemaArn = aws.String(_clouddirectorySchemaArn)
	}
	if len(_clouddirectoryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _clouddirectoryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryNextToken) > 0 {
		input.NextToken = aws.String(_clouddirectoryNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTypedLinkFacetNames(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*clouddirectory.ListTypedLinkFacetNamesOutput
	p := clouddirectory.NewListTypedLinkFacetNamesPaginator(client, input)
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

// Lists all policies from the root of the Directory to the object specified. If there are
// no policies present, an empty list is returned. If policies are present, and if
// some objects don't have the policies attached, it returns the ObjectIdentifier
// for such objects. If policies are present, it returns ObjectIdentifier ,
// policyId , and policyType . Paths that don't lead to the root from the target
// object are ignored. For more information, see [Policies].
//
// [Policies]: https://docs.aws.amazon.com/clouddirectory/latest/developerguide/key_concepts_directory.html#key_concepts_policies
func clouddirectory_LookupPolicy(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.LookupPolicyInput{
		// DirectoryArn: *string, // Required
		// ObjectReference: *types.ObjectReference, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryObjectReference) > 0 {
		if err := assignInputField(input, "ObjectReference", _clouddirectoryObjectReference); err != nil {
			log.Errorf("invalid --object-reference: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _clouddirectoryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryNextToken) > 0 {
		input.NextToken = aws.String(_clouddirectoryNextToken)
	}

	if disablePaginator() {
		if resp, err := client.LookupPolicy(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*clouddirectory.LookupPolicyOutput
	p := clouddirectory.NewLookupPolicyPaginator(client, input)
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

// Publishes a development schema with a major version and a recommended minor
// version.
func clouddirectory_PublishSchema(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.PublishSchemaInput{
		// DevelopmentSchemaArn: *string, // Required
		// Version: *string, // Required
	}

	if len(_clouddirectoryDevelopmentSchemaArn) > 0 {
		input.DevelopmentSchemaArn = aws.String(_clouddirectoryDevelopmentSchemaArn)
	}
	if len(_clouddirectoryVersion) > 0 {
		input.Version = aws.String(_clouddirectoryVersion)
	}
	if len(_clouddirectoryMinorVersion) > 0 {
		input.MinorVersion = aws.String(_clouddirectoryMinorVersion)
	}
	if len(_clouddirectoryName) > 0 {
		input.Name = aws.String(_clouddirectoryName)
	}

	if resp, err := client.PublishSchema(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows a schema to be updated using JSON upload. Only available for development
// schemas. See [JSON Schema Format]for more information.
//
// [JSON Schema Format]: https://docs.aws.amazon.com/clouddirectory/latest/developerguide/schemas_jsonformat.html#schemas_json
func clouddirectory_PutSchemaFromJson(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.PutSchemaFromJsonInput{
		// Document: *string, // Required
		// SchemaArn: *string, // Required
	}

	if len(_clouddirectoryDocument) > 0 {
		input.Document = aws.String(_clouddirectoryDocument)
	}
	if len(_clouddirectorySchemaArn) > 0 {
		input.SchemaArn = aws.String(_clouddirectorySchemaArn)
	}

	if resp, err := client.PutSchemaFromJson(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified facet from the specified object.
func clouddirectory_RemoveFacetFromObject(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.RemoveFacetFromObjectInput{
		// DirectoryArn: *string, // Required
		// ObjectReference: *types.ObjectReference, // Required
		// SchemaFacet: *types.SchemaFacet, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryObjectReference) > 0 {
		if err := assignInputField(input, "ObjectReference", _clouddirectoryObjectReference); err != nil {
			log.Errorf("invalid --object-reference: %s", err.Error())
			return
		}
	}
	if len(_clouddirectorySchemaFacet) > 0 {
		if err := assignInputField(input, "SchemaFacet", _clouddirectorySchemaFacet); err != nil {
			log.Errorf("invalid --schema-facet: %s", err.Error())
			return
		}
	}

	if resp, err := client.RemoveFacetFromObject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// An API operation for adding tags to a resource.
func clouddirectory_TagResource(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_clouddirectoryResourceArn) > 0 {
		input.ResourceArn = aws.String(_clouddirectoryResourceArn)
	}
	if len(_clouddirectoryTags) > 0 {
		if err := assignInputField(input, "Tags", _clouddirectoryTags); err != nil {
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

// An API operation for removing tags from a resource.
func clouddirectory_UntagResource(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_clouddirectoryResourceArn) > 0 {
		input.ResourceArn = aws.String(_clouddirectoryResourceArn)
	}
	if len(_clouddirectoryTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _clouddirectoryTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Does the following:
// - Adds new Attributes , Rules , or ObjectTypes .
//
// - Updates existing Attributes , Rules , or ObjectTypes .
//
// - Deletes existing Attributes , Rules , or ObjectTypes .
func clouddirectory_UpdateFacet(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.UpdateFacetInput{
		// Name: *string, // Required
		// SchemaArn: *string, // Required
	}

	if len(_clouddirectoryName) > 0 {
		input.Name = aws.String(_clouddirectoryName)
	}
	if len(_clouddirectorySchemaArn) > 0 {
		input.SchemaArn = aws.String(_clouddirectorySchemaArn)
	}
	if len(_clouddirectoryAttributeUpdates) > 0 {
		if err := assignInputField(input, "AttributeUpdates", _clouddirectoryAttributeUpdates); err != nil {
			log.Errorf("invalid --attribute-updates: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryObjectType) > 0 {
		if err := assignInputField(input, "ObjectType", _clouddirectoryObjectType); err != nil {
			log.Errorf("invalid --object-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFacet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a given typed link’s attributes. Attributes to be updated must not
// contribute to the typed link’s identity, as defined by its
// IdentityAttributeOrder .
func clouddirectory_UpdateLinkAttributes(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.UpdateLinkAttributesInput{
		// AttributeUpdates: []types.LinkAttributeUpdate, // Required
		// DirectoryArn: *string, // Required
		// TypedLinkSpecifier: *types.TypedLinkSpecifier, // Required
	}

	if len(_clouddirectoryAttributeUpdates) > 0 {
		if err := assignInputField(input, "AttributeUpdates", _clouddirectoryAttributeUpdates); err != nil {
			log.Errorf("invalid --attribute-updates: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryTypedLinkSpecifier) > 0 {
		if err := assignInputField(input, "TypedLinkSpecifier", _clouddirectoryTypedLinkSpecifier); err != nil {
			log.Errorf("invalid --typed-link-specifier: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateLinkAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a given object's attributes.
func clouddirectory_UpdateObjectAttributes(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.UpdateObjectAttributesInput{
		// AttributeUpdates: []types.ObjectAttributeUpdate, // Required
		// DirectoryArn: *string, // Required
		// ObjectReference: *types.ObjectReference, // Required
	}

	if len(_clouddirectoryAttributeUpdates) > 0 {
		if err := assignInputField(input, "AttributeUpdates", _clouddirectoryAttributeUpdates); err != nil {
			log.Errorf("invalid --attribute-updates: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryObjectReference) > 0 {
		if err := assignInputField(input, "ObjectReference", _clouddirectoryObjectReference); err != nil {
			log.Errorf("invalid --object-reference: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateObjectAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the schema name with a new name. Only development schema names can be
// updated.
func clouddirectory_UpdateSchema(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.UpdateSchemaInput{
		// Name: *string, // Required
		// SchemaArn: *string, // Required
	}

	if len(_clouddirectoryName) > 0 {
		input.Name = aws.String(_clouddirectoryName)
	}
	if len(_clouddirectorySchemaArn) > 0 {
		input.SchemaArn = aws.String(_clouddirectorySchemaArn)
	}

	if resp, err := client.UpdateSchema(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a TypedLinkFacet. For more information, see [Typed Links].
//
// [Typed Links]: https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink
func clouddirectory_UpdateTypedLinkFacet(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.UpdateTypedLinkFacetInput{
		// AttributeUpdates: []types.TypedLinkFacetAttributeUpdate, // Required
		// IdentityAttributeOrder: []string, // Required
		// Name: *string, // Required
		// SchemaArn: *string, // Required
	}

	if len(_clouddirectoryAttributeUpdates) > 0 {
		if err := assignInputField(input, "AttributeUpdates", _clouddirectoryAttributeUpdates); err != nil {
			log.Errorf("invalid --attribute-updates: %s", err.Error())
			return
		}
	}
	if len(_clouddirectoryIdentityAttributeOrder) > 0 {
		input.IdentityAttributeOrder = append([]string(nil), _clouddirectoryIdentityAttributeOrder...)
	}
	if len(_clouddirectoryName) > 0 {
		input.Name = aws.String(_clouddirectoryName)
	}
	if len(_clouddirectorySchemaArn) > 0 {
		input.SchemaArn = aws.String(_clouddirectorySchemaArn)
	}

	if resp, err := client.UpdateTypedLinkFacet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Upgrades a single directory in-place using the PublishedSchemaArn with schema
// updates found in MinorVersion . Backwards-compatible minor version upgrades are
// instantaneously available for readers on all objects in the directory. Note:
// This is a synchronous API call and upgrades only one schema on a given directory
// per call. To upgrade multiple directories from one schema, you would need to
// call this API on each directory.
func clouddirectory_UpgradeAppliedSchema(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.UpgradeAppliedSchemaInput{
		// DirectoryArn: *string, // Required
		// PublishedSchemaArn: *string, // Required
	}

	if len(_clouddirectoryDirectoryArn) > 0 {
		input.DirectoryArn = aws.String(_clouddirectoryDirectoryArn)
	}
	if len(_clouddirectoryPublishedSchemaArn) > 0 {
		input.PublishedSchemaArn = aws.String(_clouddirectoryPublishedSchemaArn)
	}
	if len(_clouddirectoryDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _clouddirectoryDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpgradeAppliedSchema(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Upgrades a published schema under a new minor version revision using the
// current contents of DevelopmentSchemaArn .
func clouddirectory_UpgradePublishedSchema(cfg aws.Config, client *clouddirectory.Client) {
	input := &clouddirectory.UpgradePublishedSchemaInput{
		// DevelopmentSchemaArn: *string, // Required
		// MinorVersion: *string, // Required
		// PublishedSchemaArn: *string, // Required
	}

	if len(_clouddirectoryDevelopmentSchemaArn) > 0 {
		input.DevelopmentSchemaArn = aws.String(_clouddirectoryDevelopmentSchemaArn)
	}
	if len(_clouddirectoryMinorVersion) > 0 {
		input.MinorVersion = aws.String(_clouddirectoryMinorVersion)
	}
	if len(_clouddirectoryPublishedSchemaArn) > 0 {
		input.PublishedSchemaArn = aws.String(_clouddirectoryPublishedSchemaArn)
	}
	if len(_clouddirectoryDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _clouddirectoryDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpgradePublishedSchema(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_clouddirectoryCmd)
	_clouddirectoryCmd.Flags().SortFlags = false

	_clouddirectoryCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_clouddirectoryCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_clouddirectoryCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_clouddirectoryCmd.Flags().StringSliceVarP(&_clouddirectoryAttributeNames, "attribute-names", "", nil, "Attribute Names")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryAttributeUpdates, "attribute-updates", "", "", "Attribute Updates")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryAttributes, "attributes", "", "", "Attributes")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryChildReference, "child-reference", "", "", "Child Reference")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryConsistencyLevel, "consistency-level", "", "", "Consistency Level")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryDevelopmentSchemaArn, "development-schema-arn", "", "", "Development Schema ARN")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryDirectoryArn, "directory-arn", "", "", "Directory ARN")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryDocument, "document", "", "", "Document")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryDryRun, "dry-run", "", "", "Dry Run")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryFacet, "facet", "", "", "Facet")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryFacetFilter, "facet-filter", "", "", "Facet Filter")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryFacetStyle, "facet-style", "", "", "Facet Style")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryFilterAttributeRanges, "filter-attribute-ranges", "", "", "Filter Attribute Ranges")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryFilterTypedLink, "filter-typed-link", "", "", "Filter Typed Link")
	_clouddirectoryCmd.Flags().StringSliceVarP(&_clouddirectoryIdentityAttributeOrder, "identity-attribute-order", "", nil, "Identity Attribute Order")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryIncludeAllLinksToEachParent, "include-all-links-to-each-parent", "", "", "Include All Links To Each Parent")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryIndexReference, "index-reference", "", "", "Index Reference")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryIsUnique, "is-unique", "", "", "Is Unique")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryLinkName, "link-name", "", "", "Link Name")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryMaxResults, "max-results", "", "", "Max Results")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryMinorVersion, "minor-version", "", "", "Minor Version")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryName, "name", "", "", "Name")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryNextToken, "next-token", "", "", "Next Token")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryObjectAttributeList, "object-attribute-list", "", "", "Object Attribute List")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryObjectReference, "object-reference", "", "", "Object Reference")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryObjectType, "object-type", "", "", "Object Type")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryOperations, "operations", "", "", "Operations")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryOrderedIndexedAttributeList, "ordered-indexed-attribute-list", "", "", "Ordered Indexed Attribute List")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryParentReference, "parent-reference", "", "", "Parent Reference")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryPolicyReference, "policy-reference", "", "", "Policy Reference")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryPublishedSchemaArn, "published-schema-arn", "", "", "Published Schema ARN")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryRangesOnIndexedValues, "ranges-on-indexed-values", "", "", "Ranges On Indexed Values")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryResourceArn, "resource-arn", "", "", "Resource ARN")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectorySchemaArn, "schema-arn", "", "", "Schema ARN")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectorySchemaFacet, "schema-facet", "", "", "Schema Facet")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectorySchemaFacets, "schema-facets", "", "", "Schema Facets")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectorySourceObjectReference, "source-object-reference", "", "", "Source Object Reference")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryState, "state", "", "", "State")
	_clouddirectoryCmd.Flags().StringSliceVarP(&_clouddirectoryTagKeys, "tag-keys", "", nil, "Tag Keys")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryTags, "tags", "", "", "Tags")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryTargetObjectReference, "target-object-reference", "", "", "Target Object Reference")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryTargetReference, "target-reference", "", "", "Target Reference")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryTypedLinkFacet, "typed-link-facet", "", "", "Typed Link Facet")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryTypedLinkSpecifier, "typed-link-specifier", "", "", "Typed Link Specifier")
	_clouddirectoryCmd.Flags().StringVarP(&_clouddirectoryVersion, "version", "", "", "Version")

	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryAddFacetToObject, "add-facet-to-object", "", false, "Add Facet To Object")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryApplySchema, "apply-schema", "", false, "Apply Schema")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryAttachObject, "attach-object", "", false, "Attach Object")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryAttachPolicy, "attach-policy", "", false, "Attach Policy")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryAttachToIndex, "attach-to-index", "", false, "Attach To Index")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryAttachTypedLink, "attach-typed-link", "", false, "Attach Typed Link")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryBatchRead, "batch-read", "", false, "Batch Read")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryBatchWrite, "batch-write", "", false, "Batch Write")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryCreateDirectory, "create-directory", "", false, "Create Directory")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryCreateFacet, "create-facet", "", false, "Create Facet")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryCreateIndex, "create-index", "", false, "Create Index")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryCreateObject, "create-object", "", false, "Create Object")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryCreateSchema, "create-schema", "", false, "Create Schema")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryCreateTypedLinkFacet, "create-typed-link-facet", "", false, "Create Typed Link Facet")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryDeleteDirectory, "delete-directory", "", false, "Delete Directory")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryDeleteFacet, "delete-facet", "", false, "Delete Facet")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryDeleteObject, "delete-object", "", false, "Delete Object")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryDeleteSchema, "delete-schema", "", false, "Delete Schema")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryDeleteTypedLinkFacet, "delete-typed-link-facet", "", false, "Delete Typed Link Facet")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryDetachFromIndex, "detach-from-index", "", false, "Detach From Index")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryDetachObject, "detach-object", "", false, "Detach Object")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryDetachPolicy, "detach-policy", "", false, "Detach Policy")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryDetachTypedLink, "detach-typed-link", "", false, "Detach Typed Link")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryDisableDirectory, "disable-directory", "", false, "Disable Directory")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryEnableDirectory, "enable-directory", "", false, "Enable Directory")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryGetAppliedSchemaVersion, "get-applied-schema-version", "", false, "Get Applied Schema Version")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryGetDirectory, "get-directory", "", false, "Get Directory")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryGetFacet, "get-facet", "", false, "Get Facet")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryGetLinkAttributes, "get-link-attributes", "", false, "Get Link Attributes")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryGetObjectAttributes, "get-object-attributes", "", false, "Get Object Attributes")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryGetObjectInformation, "get-object-information", "", false, "Get Object Information")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryGetSchemaAsJson, "get-schema-as-json", "", false, "Get Schema As JSON")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryGetTypedLinkFacetInformation, "get-typed-link-facet-information", "", false, "Get Typed Link Facet Information")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryListAppliedSchemaArns, "list-applied-schema-arns", "", false, "List Applied Schema Arns")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryListAttachedIndices, "list-attached-indices", "", false, "List Attached Indices")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryListDevelopmentSchemaArns, "list-development-schema-arns", "", false, "List Development Schema Arns")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryListDirectories, "list-directories", "", false, "List Directories")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryListFacetAttributes, "list-facet-attributes", "", false, "List Facet Attributes")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryListFacetNames, "list-facet-names", "", false, "List Facet Names")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryListIncomingTypedLinks, "list-incoming-typed-links", "", false, "List Incoming Typed Links")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryListIndex, "list-index", "", false, "List Index")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryListManagedSchemaArns, "list-managed-schema-arns", "", false, "List Managed Schema Arns")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryListObjectAttributes, "list-object-attributes", "", false, "List Object Attributes")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryListObjectChildren, "list-object-children", "", false, "List Object Children")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryListObjectParentPaths, "list-object-parent-paths", "", false, "List Object Parent Paths")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryListObjectParents, "list-object-parents", "", false, "List Object Parents")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryListObjectPolicies, "list-object-policies", "", false, "List Object Policies")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryListOutgoingTypedLinks, "list-outgoing-typed-links", "", false, "List Outgoing Typed Links")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryListPolicyAttachments, "list-policy-attachments", "", false, "List Policy Attachments")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryListPublishedSchemaArns, "list-published-schema-arns", "", false, "List Published Schema Arns")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryListTypedLinkFacetAttributes, "list-typed-link-facet-attributes", "", false, "List Typed Link Facet Attributes")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryListTypedLinkFacetNames, "list-typed-link-facet-names", "", false, "List Typed Link Facet Names")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryLookupPolicy, "lookup-policy", "", false, "Lookup Policy")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryPublishSchema, "publish-schema", "", false, "Publish Schema")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryPutSchemaFromJson, "put-schema-from-json", "", false, "Put Schema From JSON")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryRemoveFacetFromObject, "remove-facet-from-object", "", false, "Remove Facet From Object")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryTagResource, "tag-resource", "", false, "Tag Resource")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryUntagResource, "untag-resource", "", false, "Untag Resource")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryUpdateFacet, "update-facet", "", false, "Update Facet")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryUpdateLinkAttributes, "update-link-attributes", "", false, "Update Link Attributes")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryUpdateObjectAttributes, "update-object-attributes", "", false, "Update Object Attributes")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryUpdateSchema, "update-schema", "", false, "Update Schema")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryUpdateTypedLinkFacet, "update-typed-link-facet", "", false, "Update Typed Link Facet")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryUpgradeAppliedSchema, "upgrade-applied-schema", "", false, "Upgrade Applied Schema")
	_clouddirectoryCmd.Flags().BoolVarP(&_clouddirectoryUpgradePublishedSchema, "upgrade-published-schema", "", false, "Upgrade Published Schema")

}
