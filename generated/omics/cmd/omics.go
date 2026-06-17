package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/omics"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// omicsCmd represents the omics command
var _omicsCmd = &cobra.Command{
	Use:   "omics",
	Short: "AWS omics CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := omics.NewFromConfig(cfg)
		if _omicsAbortMultipartReadSetUpload {
			omics_AbortMultipartReadSetUpload(cfg, client)
			return
		}
		if _omicsAcceptShare {
			omics_AcceptShare(cfg, client)
			return
		}
		if _omicsBatchDeleteReadSet {
			omics_BatchDeleteReadSet(cfg, client)
			return
		}
		if _omicsCancelAnnotationImportJob {
			omics_CancelAnnotationImportJob(cfg, client)
			return
		}
		if _omicsCancelRun {
			omics_CancelRun(cfg, client)
			return
		}
		if _omicsCancelVariantImportJob {
			omics_CancelVariantImportJob(cfg, client)
			return
		}
		if _omicsCompleteMultipartReadSetUpload {
			omics_CompleteMultipartReadSetUpload(cfg, client)
			return
		}
		if _omicsCreateAnnotationStore {
			omics_CreateAnnotationStore(cfg, client)
			return
		}
		if _omicsCreateAnnotationStoreVersion {
			omics_CreateAnnotationStoreVersion(cfg, client)
			return
		}
		if _omicsCreateMultipartReadSetUpload {
			omics_CreateMultipartReadSetUpload(cfg, client)
			return
		}
		if _omicsCreateReferenceStore {
			omics_CreateReferenceStore(cfg, client)
			return
		}
		if _omicsCreateRunCache {
			omics_CreateRunCache(cfg, client)
			return
		}
		if _omicsCreateRunGroup {
			omics_CreateRunGroup(cfg, client)
			return
		}
		if _omicsCreateSequenceStore {
			omics_CreateSequenceStore(cfg, client)
			return
		}
		if _omicsCreateShare {
			omics_CreateShare(cfg, client)
			return
		}
		if _omicsCreateVariantStore {
			omics_CreateVariantStore(cfg, client)
			return
		}
		if _omicsCreateWorkflow {
			omics_CreateWorkflow(cfg, client)
			return
		}
		if _omicsCreateWorkflowVersion {
			omics_CreateWorkflowVersion(cfg, client)
			return
		}
		if _omicsDeleteAnnotationStore {
			omics_DeleteAnnotationStore(cfg, client)
			return
		}
		if _omicsDeleteAnnotationStoreVersions {
			omics_DeleteAnnotationStoreVersions(cfg, client)
			return
		}
		if _omicsDeleteReference {
			omics_DeleteReference(cfg, client)
			return
		}
		if _omicsDeleteReferenceStore {
			omics_DeleteReferenceStore(cfg, client)
			return
		}
		if _omicsDeleteRun {
			omics_DeleteRun(cfg, client)
			return
		}
		if _omicsDeleteRunCache {
			omics_DeleteRunCache(cfg, client)
			return
		}
		if _omicsDeleteRunGroup {
			omics_DeleteRunGroup(cfg, client)
			return
		}
		if _omicsDeleteS3AccessPolicy {
			omics_DeleteS3AccessPolicy(cfg, client)
			return
		}
		if _omicsDeleteSequenceStore {
			omics_DeleteSequenceStore(cfg, client)
			return
		}
		if _omicsDeleteShare {
			omics_DeleteShare(cfg, client)
			return
		}
		if _omicsDeleteVariantStore {
			omics_DeleteVariantStore(cfg, client)
			return
		}
		if _omicsDeleteWorkflow {
			omics_DeleteWorkflow(cfg, client)
			return
		}
		if _omicsDeleteWorkflowVersion {
			omics_DeleteWorkflowVersion(cfg, client)
			return
		}
		if _omicsGetAnnotationImportJob {
			omics_GetAnnotationImportJob(cfg, client)
			return
		}
		if _omicsGetAnnotationStore {
			omics_GetAnnotationStore(cfg, client)
			return
		}
		if _omicsGetAnnotationStoreVersion {
			omics_GetAnnotationStoreVersion(cfg, client)
			return
		}
		if _omicsGetReadSet {
			omics_GetReadSet(cfg, client)
			return
		}
		if _omicsGetReadSetActivationJob {
			omics_GetReadSetActivationJob(cfg, client)
			return
		}
		if _omicsGetReadSetExportJob {
			omics_GetReadSetExportJob(cfg, client)
			return
		}
		if _omicsGetReadSetImportJob {
			omics_GetReadSetImportJob(cfg, client)
			return
		}
		if _omicsGetReadSetMetadata {
			omics_GetReadSetMetadata(cfg, client)
			return
		}
		if _omicsGetReference {
			omics_GetReference(cfg, client)
			return
		}
		if _omicsGetReferenceImportJob {
			omics_GetReferenceImportJob(cfg, client)
			return
		}
		if _omicsGetReferenceMetadata {
			omics_GetReferenceMetadata(cfg, client)
			return
		}
		if _omicsGetReferenceStore {
			omics_GetReferenceStore(cfg, client)
			return
		}
		if _omicsGetRun {
			omics_GetRun(cfg, client)
			return
		}
		if _omicsGetRunCache {
			omics_GetRunCache(cfg, client)
			return
		}
		if _omicsGetRunGroup {
			omics_GetRunGroup(cfg, client)
			return
		}
		if _omicsGetRunTask {
			omics_GetRunTask(cfg, client)
			return
		}
		if _omicsGetS3AccessPolicy {
			omics_GetS3AccessPolicy(cfg, client)
			return
		}
		if _omicsGetSequenceStore {
			omics_GetSequenceStore(cfg, client)
			return
		}
		if _omicsGetShare {
			omics_GetShare(cfg, client)
			return
		}
		if _omicsGetVariantImportJob {
			omics_GetVariantImportJob(cfg, client)
			return
		}
		if _omicsGetVariantStore {
			omics_GetVariantStore(cfg, client)
			return
		}
		if _omicsGetWorkflow {
			omics_GetWorkflow(cfg, client)
			return
		}
		if _omicsGetWorkflowVersion {
			omics_GetWorkflowVersion(cfg, client)
			return
		}
		if _omicsListAnnotationImportJobs {
			omics_ListAnnotationImportJobs(cfg, client)
			return
		}
		if _omicsListAnnotationStoreVersions {
			omics_ListAnnotationStoreVersions(cfg, client)
			return
		}
		if _omicsListAnnotationStores {
			omics_ListAnnotationStores(cfg, client)
			return
		}
		if _omicsListMultipartReadSetUploads {
			omics_ListMultipartReadSetUploads(cfg, client)
			return
		}
		if _omicsListReadSetActivationJobs {
			omics_ListReadSetActivationJobs(cfg, client)
			return
		}
		if _omicsListReadSetExportJobs {
			omics_ListReadSetExportJobs(cfg, client)
			return
		}
		if _omicsListReadSetImportJobs {
			omics_ListReadSetImportJobs(cfg, client)
			return
		}
		if _omicsListReadSetUploadParts {
			omics_ListReadSetUploadParts(cfg, client)
			return
		}
		if _omicsListReadSets {
			omics_ListReadSets(cfg, client)
			return
		}
		if _omicsListReferenceImportJobs {
			omics_ListReferenceImportJobs(cfg, client)
			return
		}
		if _omicsListReferenceStores {
			omics_ListReferenceStores(cfg, client)
			return
		}
		if _omicsListReferences {
			omics_ListReferences(cfg, client)
			return
		}
		if _omicsListRunCaches {
			omics_ListRunCaches(cfg, client)
			return
		}
		if _omicsListRunGroups {
			omics_ListRunGroups(cfg, client)
			return
		}
		if _omicsListRunTasks {
			omics_ListRunTasks(cfg, client)
			return
		}
		if _omicsListRuns {
			omics_ListRuns(cfg, client)
			return
		}
		if _omicsListSequenceStores {
			omics_ListSequenceStores(cfg, client)
			return
		}
		if _omicsListShares {
			omics_ListShares(cfg, client)
			return
		}
		if _omicsListTagsForResource {
			omics_ListTagsForResource(cfg, client)
			return
		}
		if _omicsListVariantImportJobs {
			omics_ListVariantImportJobs(cfg, client)
			return
		}
		if _omicsListVariantStores {
			omics_ListVariantStores(cfg, client)
			return
		}
		if _omicsListWorkflowVersions {
			omics_ListWorkflowVersions(cfg, client)
			return
		}
		if _omicsListWorkflows {
			omics_ListWorkflows(cfg, client)
			return
		}
		if _omicsPutS3AccessPolicy {
			omics_PutS3AccessPolicy(cfg, client)
			return
		}
		if _omicsStartAnnotationImportJob {
			omics_StartAnnotationImportJob(cfg, client)
			return
		}
		if _omicsStartReadSetActivationJob {
			omics_StartReadSetActivationJob(cfg, client)
			return
		}
		if _omicsStartReadSetExportJob {
			omics_StartReadSetExportJob(cfg, client)
			return
		}
		if _omicsStartReadSetImportJob {
			omics_StartReadSetImportJob(cfg, client)
			return
		}
		if _omicsStartReferenceImportJob {
			omics_StartReferenceImportJob(cfg, client)
			return
		}
		if _omicsStartRun {
			omics_StartRun(cfg, client)
			return
		}
		if _omicsStartVariantImportJob {
			omics_StartVariantImportJob(cfg, client)
			return
		}
		if _omicsTagResource {
			omics_TagResource(cfg, client)
			return
		}
		if _omicsUntagResource {
			omics_UntagResource(cfg, client)
			return
		}
		if _omicsUpdateAnnotationStore {
			omics_UpdateAnnotationStore(cfg, client)
			return
		}
		if _omicsUpdateAnnotationStoreVersion {
			omics_UpdateAnnotationStoreVersion(cfg, client)
			return
		}
		if _omicsUpdateRunCache {
			omics_UpdateRunCache(cfg, client)
			return
		}
		if _omicsUpdateRunGroup {
			omics_UpdateRunGroup(cfg, client)
			return
		}
		if _omicsUpdateSequenceStore {
			omics_UpdateSequenceStore(cfg, client)
			return
		}
		if _omicsUpdateVariantStore {
			omics_UpdateVariantStore(cfg, client)
			return
		}
		if _omicsUpdateWorkflow {
			omics_UpdateWorkflow(cfg, client)
			return
		}
		if _omicsUpdateWorkflowVersion {
			omics_UpdateWorkflowVersion(cfg, client)
			return
		}
		if _omicsUploadReadSetPart {
			omics_UploadReadSetPart(cfg, client)
			return
		}

	},
}

var (
	_omicsAbortMultipartReadSetUpload    bool
	_omicsAcceptShare                    bool
	_omicsBatchDeleteReadSet             bool
	_omicsCancelAnnotationImportJob      bool
	_omicsCancelRun                      bool
	_omicsCancelVariantImportJob         bool
	_omicsCompleteMultipartReadSetUpload bool
	_omicsCreateAnnotationStore          bool
	_omicsCreateAnnotationStoreVersion   bool
	_omicsCreateMultipartReadSetUpload   bool
	_omicsCreateReferenceStore           bool
	_omicsCreateRunCache                 bool
	_omicsCreateRunGroup                 bool
	_omicsCreateSequenceStore            bool
	_omicsCreateShare                    bool
	_omicsCreateVariantStore             bool
	_omicsCreateWorkflow                 bool
	_omicsCreateWorkflowVersion          bool
	_omicsDeleteAnnotationStore          bool
	_omicsDeleteAnnotationStoreVersions  bool
	_omicsDeleteReference                bool
	_omicsDeleteReferenceStore           bool
	_omicsDeleteRun                      bool
	_omicsDeleteRunCache                 bool
	_omicsDeleteRunGroup                 bool
	_omicsDeleteS3AccessPolicy           bool
	_omicsDeleteSequenceStore            bool
	_omicsDeleteShare                    bool
	_omicsDeleteVariantStore             bool
	_omicsDeleteWorkflow                 bool
	_omicsDeleteWorkflowVersion          bool
	_omicsGetAnnotationImportJob         bool
	_omicsGetAnnotationStore             bool
	_omicsGetAnnotationStoreVersion      bool
	_omicsGetReadSet                     bool
	_omicsGetReadSetActivationJob        bool
	_omicsGetReadSetExportJob            bool
	_omicsGetReadSetImportJob            bool
	_omicsGetReadSetMetadata             bool
	_omicsGetReference                   bool
	_omicsGetReferenceImportJob          bool
	_omicsGetReferenceMetadata           bool
	_omicsGetReferenceStore              bool
	_omicsGetRun                         bool
	_omicsGetRunCache                    bool
	_omicsGetRunGroup                    bool
	_omicsGetRunTask                     bool
	_omicsGetS3AccessPolicy              bool
	_omicsGetSequenceStore               bool
	_omicsGetShare                       bool
	_omicsGetVariantImportJob            bool
	_omicsGetVariantStore                bool
	_omicsGetWorkflow                    bool
	_omicsGetWorkflowVersion             bool
	_omicsListAnnotationImportJobs       bool
	_omicsListAnnotationStoreVersions    bool
	_omicsListAnnotationStores           bool
	_omicsListMultipartReadSetUploads    bool
	_omicsListReadSetActivationJobs      bool
	_omicsListReadSetExportJobs          bool
	_omicsListReadSetImportJobs          bool
	_omicsListReadSetUploadParts         bool
	_omicsListReadSets                   bool
	_omicsListReferenceImportJobs        bool
	_omicsListReferenceStores            bool
	_omicsListReferences                 bool
	_omicsListRunCaches                  bool
	_omicsListRunGroups                  bool
	_omicsListRunTasks                   bool
	_omicsListRuns                       bool
	_omicsListSequenceStores             bool
	_omicsListShares                     bool
	_omicsListTagsForResource            bool
	_omicsListVariantImportJobs          bool
	_omicsListVariantStores              bool
	_omicsListWorkflowVersions           bool
	_omicsListWorkflows                  bool
	_omicsPutS3AccessPolicy              bool
	_omicsStartAnnotationImportJob       bool
	_omicsStartReadSetActivationJob      bool
	_omicsStartReadSetExportJob          bool
	_omicsStartReadSetImportJob          bool
	_omicsStartReferenceImportJob        bool
	_omicsStartRun                       bool
	_omicsStartVariantImportJob          bool
	_omicsTagResource                    bool
	_omicsUntagResource                  bool
	_omicsUpdateAnnotationStore          bool
	_omicsUpdateAnnotationStoreVersion   bool
	_omicsUpdateRunCache                 bool
	_omicsUpdateRunGroup                 bool
	_omicsUpdateSequenceStore            bool
	_omicsUpdateVariantStore             bool
	_omicsUpdateWorkflow                 bool
	_omicsUpdateWorkflowVersion          bool
	_omicsUploadReadSetPart              bool

	_omicsAccelerators            string
	_omicsAnnotationFields        string
	_omicsCacheBehavior           string
	_omicsCacheBucketOwnerId      string
	_omicsCacheId                 string
	_omicsCacheS3Location         string
	_omicsClientToken             string
	_omicsContainerRegistryMap    string
	_omicsContainerRegistryMapUri string
	_omicsDefinitionRepository    string
	_omicsDefinitionUri           string
	_omicsDefinitionZip           string
	_omicsDescription             string
	_omicsDestination             string
	_omicsDestinationName         string
	_omicsEngine                  string
	_omicsETagAlgorithmFamily     string
	_omicsExport                  string
	_omicsFallbackLocation        string
	_omicsFile                    string
	_omicsFilter                  string
	_omicsForce                   string
	_omicsFormatOptions           string
	_omicsGeneratedFrom           string
	_omicsId                      string
	_omicsIds                     []string
	_omicsItems                   string
	_omicsJobId                   string
	_omicsLogLevel                string
	_omicsMain                    string
	_omicsMaxCpus                 string
	_omicsMaxDuration             string
	_omicsMaxGpus                 string
	_omicsMaxResults              string
	_omicsMaxRuns                 string
	_omicsName                    string
	_omicsNextToken               string
	_omicsOutputUri               string
	_omicsParameterTemplate       string
	_omicsParameterTemplatePath   string
	_omicsParameters              string
	_omicsPartNumber              string
	_omicsPartSource              string
	_omicsParts                   string
	_omicsPayload                 string
	_omicsPrincipalSubscriber     string
	_omicsPriority                string
	_omicsPropagatedSetLevelTags  []string
	_omicsRange                   string
	_omicsReadmeMarkdown          string
	_omicsReadmePath              string
	_omicsReadmeUri               string
	_omicsReference               string
	_omicsReferenceArn            string
	_omicsReferenceStoreId        string
	_omicsRequestId               string
	_omicsResourceArn             string
	_omicsResourceOwner           string
	_omicsRetentionMode           string
	_omicsRoleArn                 string
	_omicsRunGroupId              string
	_omicsRunId                   string
	_omicsRunLeftNormalization    string
	_omicsS3AccessConfig          string
	_omicsS3AccessPointArn        string
	_omicsS3AccessPolicy          string
	_omicsSampleId                string
	_omicsSequenceStoreId         string
	_omicsShareId                 string
	_omicsShareName               string
	_omicsSourceFileType          string
	_omicsSources                 string
	_omicsSseConfig               string
	_omicsStartingToken           string
	_omicsStatus                  string
	_omicsStorageCapacity         string
	_omicsStorageType             string
	_omicsStoreFormat             string
	_omicsStoreOptions            string
	_omicsSubjectId               string
	_omicsTagKeys                 []string
	_omicsTags                    string
	_omicsTaskId                  string
	_omicsType                    string
	_omicsUploadId                string
	_omicsVersionName             string
	_omicsVersionOptions          string
	_omicsVersions                []string
	_omicsWorkflowBucketOwnerId   string
	_omicsWorkflowId              string
	_omicsWorkflowOwnerId         string
	_omicsWorkflowType            string
	_omicsWorkflowVersionName     string
)

// Stops a multipart read set upload into a sequence store and returns a response
// with no body if the operation is successful. To confirm that a multipart read
// set upload has been stopped, use the ListMultipartReadSetUploads API operation
// to view all active multipart read set uploads.
func omics_AbortMultipartReadSetUpload(cfg aws.Config, client *omics.Client) {
	input := &omics.AbortMultipartReadSetUploadInput{
		// SequenceStoreId: *string, // Required
		// UploadId: *string, // Required
	}

	if len(_omicsSequenceStoreId) > 0 {
		input.SequenceStoreId = aws.String(_omicsSequenceStoreId)
	}
	if len(_omicsUploadId) > 0 {
		input.UploadId = aws.String(_omicsUploadId)
	}

	if resp, err := client.AbortMultipartReadSetUpload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Accept a resource share request.
func omics_AcceptShare(cfg aws.Config, client *omics.Client) {
	input := &omics.AcceptShareInput{
		// ShareId: *string, // Required
	}

	if len(_omicsShareId) > 0 {
		input.ShareId = aws.String(_omicsShareId)
	}

	if resp, err := client.AcceptShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes one or more read sets. If the operation is successful, it returns a
// response with no body. If there is an error with deleting one of the read sets,
// the operation returns an error list. If the operation successfully deletes only
// a subset of files, it will return an error list for the remaining files that
// fail to be deleted. There is a limit of 100 read sets that can be deleted in
// each BatchDeleteReadSet API call.
func omics_BatchDeleteReadSet(cfg aws.Config, client *omics.Client) {
	input := &omics.BatchDeleteReadSetInput{
		// Ids: []string, // Required
		// SequenceStoreId: *string, // Required
	}

	if len(_omicsIds) > 0 {
		input.Ids = append([]string(nil), _omicsIds...)
	}
	if len(_omicsSequenceStoreId) > 0 {
		input.SequenceStoreId = aws.String(_omicsSequenceStoreId)
	}

	if resp, err := client.BatchDeleteReadSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Amazon Web Services HealthOmics variant stores and annotation stores will no
// longer be open to new customers starting November 7, 2025. If you would like to
// use variant stores or annotation stores, sign up prior to that date. Existing
// customers can continue to use the service as normal. For more information, see [Amazon Web Services HealthOmics variant store and annotation store availability change].
//
// Cancels an annotation import job.
//
// [Amazon Web Services HealthOmics variant store and annotation store availability change]: https://docs.aws.amazon.com/omics/latest/dev/variant-store-availability-change.html
func omics_CancelAnnotationImportJob(cfg aws.Config, client *omics.Client) {
	input := &omics.CancelAnnotationImportJobInput{
		// JobId: *string, // Required
	}

	if len(_omicsJobId) > 0 {
		input.JobId = aws.String(_omicsJobId)
	}

	if resp, err := client.CancelAnnotationImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a run using its ID and returns a response with no body if the operation
// is successful. To confirm that the run has been cancelled, use the ListRuns API
// operation to check that it is no longer listed.
func omics_CancelRun(cfg aws.Config, client *omics.Client) {
	input := &omics.CancelRunInput{
		// Id: *string, // Required
	}

	if len(_omicsId) > 0 {
		input.Id = aws.String(_omicsId)
	}

	if resp, err := client.CancelRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Amazon Web Services HealthOmics variant stores and annotation stores will no
// longer be open to new customers starting November 7, 2025. If you would like to
// use variant stores or annotation stores, sign up prior to that date. Existing
// customers can continue to use the service as normal. For more information, see [Amazon Web Services HealthOmics variant store and annotation store availability change].
//
// Cancels a variant import job.
//
// [Amazon Web Services HealthOmics variant store and annotation store availability change]: https://docs.aws.amazon.com/omics/latest/dev/variant-store-availability-change.html
func omics_CancelVariantImportJob(cfg aws.Config, client *omics.Client) {
	input := &omics.CancelVariantImportJobInput{
		// JobId: *string, // Required
	}

	if len(_omicsJobId) > 0 {
		input.JobId = aws.String(_omicsJobId)
	}

	if resp, err := client.CancelVariantImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Completes a multipart read set upload into a sequence store after you have
// initiated the upload process with CreateMultipartReadSetUpload and uploaded all
// read set parts using UploadReadSetPart . You must specify the parts you uploaded
// using the parts parameter. If the operation is successful, it returns the read
// set ID(s) of the uploaded read set(s).
//
// For more information, see [Direct upload to a sequence store] in the Amazon Web Services HealthOmics User Guide.
//
// [Direct upload to a sequence store]: https://docs.aws.amazon.com/omics/latest/dev/synchronous-uploads.html
func omics_CompleteMultipartReadSetUpload(cfg aws.Config, client *omics.Client) {
	input := &omics.CompleteMultipartReadSetUploadInput{
		// Parts: []types.CompleteReadSetUploadPartListItem, // Required
		// SequenceStoreId: *string, // Required
		// UploadId: *string, // Required
	}

	if len(_omicsParts) > 0 {
		if err := assignInputField(input, "Parts", _omicsParts); err != nil {
			log.Errorf("invalid --parts: %s", err.Error())
			return
		}
	}
	if len(_omicsSequenceStoreId) > 0 {
		input.SequenceStoreId = aws.String(_omicsSequenceStoreId)
	}
	if len(_omicsUploadId) > 0 {
		input.UploadId = aws.String(_omicsUploadId)
	}

	if resp, err := client.CompleteMultipartReadSetUpload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Amazon Web Services HealthOmics variant stores and annotation stores will no
// longer be open to new customers starting November 7, 2025. If you would like to
// use variant stores or annotation stores, sign up prior to that date. Existing
// customers can continue to use the service as normal. For more information, see [Amazon Web Services HealthOmics variant store and annotation store availability change].
//
// Creates an annotation store.
//
// [Amazon Web Services HealthOmics variant store and annotation store availability change]: https://docs.aws.amazon.com/omics/latest/dev/variant-store-availability-change.html
func omics_CreateAnnotationStore(cfg aws.Config, client *omics.Client) {
	input := &omics.CreateAnnotationStoreInput{
		// StoreFormat: types.StoreFormat, // Required
	}

	if len(_omicsStoreFormat) > 0 {
		if err := assignInputField(input, "StoreFormat", _omicsStoreFormat); err != nil {
			log.Errorf("invalid --store-format: %s", err.Error())
			return
		}
	}
	if len(_omicsDescription) > 0 {
		input.Description = aws.String(_omicsDescription)
	}
	if len(_omicsName) > 0 {
		input.Name = aws.String(_omicsName)
	}
	if len(_omicsReference) > 0 {
		if err := assignInputField(input, "Reference", _omicsReference); err != nil {
			log.Errorf("invalid --reference: %s", err.Error())
			return
		}
	}
	if len(_omicsSseConfig) > 0 {
		if err := assignInputField(input, "SseConfig", _omicsSseConfig); err != nil {
			log.Errorf("invalid --sse-config: %s", err.Error())
			return
		}
	}
	if len(_omicsStoreOptions) > 0 {
		if err := assignInputField(input, "StoreOptions", _omicsStoreOptions); err != nil {
			log.Errorf("invalid --store-options: %s", err.Error())
			return
		}
	}
	if len(_omicsTags) > 0 {
		if err := assignInputField(input, "Tags", _omicsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_omicsVersionName) > 0 {
		input.VersionName = aws.String(_omicsVersionName)
	}

	if resp, err := client.CreateAnnotationStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new version of an annotation store.
func omics_CreateAnnotationStoreVersion(cfg aws.Config, client *omics.Client) {
	input := &omics.CreateAnnotationStoreVersionInput{
		// Name: *string, // Required
		// VersionName: *string, // Required
	}

	if len(_omicsName) > 0 {
		input.Name = aws.String(_omicsName)
	}
	if len(_omicsVersionName) > 0 {
		input.VersionName = aws.String(_omicsVersionName)
	}
	if len(_omicsDescription) > 0 {
		input.Description = aws.String(_omicsDescription)
	}
	if len(_omicsTags) > 0 {
		if err := assignInputField(input, "Tags", _omicsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_omicsVersionOptions) > 0 {
		if err := assignInputField(input, "VersionOptions", _omicsVersionOptions); err != nil {
			log.Errorf("invalid --version-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAnnotationStoreVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates a multipart read set upload for uploading partitioned source files
// into a sequence store. You can directly import source files from an EC2 instance
// and other local compute, or from an S3 bucket. To separate these source files
// into parts, use the split operation. Each part cannot be larger than 100 MB. If
// the operation is successful, it provides an uploadId which is required by the
// UploadReadSetPart API operation to upload parts into a sequence store.
//
// To continue uploading a multipart read set into your sequence store, you must
// use the UploadReadSetPart API operation to upload each part individually
// following the steps below:
//
// - Specify the uploadId obtained from the previous call to
// CreateMultipartReadSetUpload .
//
// - Upload parts for that uploadId .
//
// When you have finished uploading parts, use the CompleteMultipartReadSetUpload
// API to complete the multipart read set upload and to retrieve the final read set
// IDs in the response.
//
// To learn more about creating parts and the split operation, see [Direct upload to a sequence store] in the Amazon
// Web Services HealthOmics User Guide.
//
// [Direct upload to a sequence store]: https://docs.aws.amazon.com/omics/latest/dev/synchronous-uploads.html
func omics_CreateMultipartReadSetUpload(cfg aws.Config, client *omics.Client) {
	input := &omics.CreateMultipartReadSetUploadInput{
		// Name: *string, // Required
		// SampleId: *string, // Required
		// SequenceStoreId: *string, // Required
		// SourceFileType: types.FileType, // Required
		// SubjectId: *string, // Required
	}

	if len(_omicsName) > 0 {
		input.Name = aws.String(_omicsName)
	}
	if len(_omicsSampleId) > 0 {
		input.SampleId = aws.String(_omicsSampleId)
	}
	if len(_omicsSequenceStoreId) > 0 {
		input.SequenceStoreId = aws.String(_omicsSequenceStoreId)
	}
	if len(_omicsSourceFileType) > 0 {
		if err := assignInputField(input, "SourceFileType", _omicsSourceFileType); err != nil {
			log.Errorf("invalid --source-file-type: %s", err.Error())
			return
		}
	}
	if len(_omicsSubjectId) > 0 {
		input.SubjectId = aws.String(_omicsSubjectId)
	}
	if len(_omicsClientToken) > 0 {
		input.ClientToken = aws.String(_omicsClientToken)
	}
	if len(_omicsDescription) > 0 {
		input.Description = aws.String(_omicsDescription)
	}
	if len(_omicsGeneratedFrom) > 0 {
		input.GeneratedFrom = aws.String(_omicsGeneratedFrom)
	}
	if len(_omicsReferenceArn) > 0 {
		input.ReferenceArn = aws.String(_omicsReferenceArn)
	}
	if len(_omicsTags) > 0 {
		if err := assignInputField(input, "Tags", _omicsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMultipartReadSetUpload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a reference store and returns metadata in JSON format. Reference stores
// are used to store reference genomes in FASTA format. A reference store is
// created when the first reference genome is imported. To import additional
// reference genomes from an Amazon S3 bucket, use the StartReferenceImportJob API
// operation.
//
// For more information, see [Creating a HealthOmics reference store] in the Amazon Web Services HealthOmics User Guide.
//
// [Creating a HealthOmics reference store]: https://docs.aws.amazon.com/omics/latest/dev/create-reference-store.html
func omics_CreateReferenceStore(cfg aws.Config, client *omics.Client) {
	input := &omics.CreateReferenceStoreInput{
		// Name: *string, // Required
	}

	if len(_omicsName) > 0 {
		input.Name = aws.String(_omicsName)
	}
	if len(_omicsClientToken) > 0 {
		input.ClientToken = aws.String(_omicsClientToken)
	}
	if len(_omicsDescription) > 0 {
		input.Description = aws.String(_omicsDescription)
	}
	if len(_omicsSseConfig) > 0 {
		if err := assignInputField(input, "SseConfig", _omicsSseConfig); err != nil {
			log.Errorf("invalid --sse-config: %s", err.Error())
			return
		}
	}
	if len(_omicsTags) > 0 {
		if err := assignInputField(input, "Tags", _omicsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateReferenceStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a run cache to store and reference task outputs from completed private
// runs. Specify an Amazon S3 location where Amazon Web Services HealthOmics saves
// the cached data. This data must be immediately accessible and not in an archived
// state. You can save intermediate task files to a run cache if they are declared
// as task outputs in the workflow definition file.
//
// For more information, see [Call caching] and [Creating a run cache] in the Amazon Web Services HealthOmics User
// Guide.
//
// [Call caching]: https://docs.aws.amazon.com/omics/latest/dev/workflows-call-caching.html
// [Creating a run cache]: https://docs.aws.amazon.com/omics/latest/dev/workflow-cache-create.html
func omics_CreateRunCache(cfg aws.Config, client *omics.Client) {
	input := &omics.CreateRunCacheInput{
		// CacheS3Location: *string, // Required
		// RequestId: *string, // Required
	}

	if len(_omicsCacheS3Location) > 0 {
		input.CacheS3Location = aws.String(_omicsCacheS3Location)
	}
	if len(_omicsRequestId) > 0 {
		input.RequestId = aws.String(_omicsRequestId)
	}
	if len(_omicsCacheBehavior) > 0 {
		if err := assignInputField(input, "CacheBehavior", _omicsCacheBehavior); err != nil {
			log.Errorf("invalid --cache-behavior: %s", err.Error())
			return
		}
	}
	if len(_omicsCacheBucketOwnerId) > 0 {
		input.CacheBucketOwnerId = aws.String(_omicsCacheBucketOwnerId)
	}
	if len(_omicsDescription) > 0 {
		input.Description = aws.String(_omicsDescription)
	}
	if len(_omicsName) > 0 {
		input.Name = aws.String(_omicsName)
	}
	if len(_omicsTags) > 0 {
		if err := assignInputField(input, "Tags", _omicsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRunCache(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a run group to limit the compute resources for the runs that are added
// to the group. Returns an ARN, ID, and tags for the run group.
func omics_CreateRunGroup(cfg aws.Config, client *omics.Client) {
	input := &omics.CreateRunGroupInput{
		// RequestId: *string, // Required
	}

	if len(_omicsRequestId) > 0 {
		input.RequestId = aws.String(_omicsRequestId)
	}
	if len(_omicsMaxCpus) > 0 {
		if err := assignInputField(input, "MaxCpus", _omicsMaxCpus); err != nil {
			log.Errorf("invalid --max-cpus: %s", err.Error())
			return
		}
	}
	if len(_omicsMaxDuration) > 0 {
		if err := assignInputField(input, "MaxDuration", _omicsMaxDuration); err != nil {
			log.Errorf("invalid --max-duration: %s", err.Error())
			return
		}
	}
	if len(_omicsMaxGpus) > 0 {
		if err := assignInputField(input, "MaxGpus", _omicsMaxGpus); err != nil {
			log.Errorf("invalid --max-gpus: %s", err.Error())
			return
		}
	}
	if len(_omicsMaxRuns) > 0 {
		if err := assignInputField(input, "MaxRuns", _omicsMaxRuns); err != nil {
			log.Errorf("invalid --max-runs: %s", err.Error())
			return
		}
	}
	if len(_omicsName) > 0 {
		input.Name = aws.String(_omicsName)
	}
	if len(_omicsTags) > 0 {
		if err := assignInputField(input, "Tags", _omicsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRunGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a sequence store and returns its metadata. Sequence stores are used to
// store sequence data files called read sets that are saved in FASTQ, BAM, uBAM,
// or CRAM formats. For aligned formats (BAM and CRAM), a sequence store can only
// use one reference genome. For unaligned formats (FASTQ and uBAM), a reference
// genome is not required. You can create multiple sequence stores per region per
// account.
//
// The following are optional parameters you can specify for your sequence store:
//
// - Use s3AccessConfig to configure your sequence store with S3 access logs
// (recommended).
//
// - Use sseConfig to define your own KMS key for encryption.
//
// - Use eTagAlgorithmFamily to define which algorithm to use for the HealthOmics
// eTag on objects.
//
// - Use fallbackLocation to define a backup location for storing files that have
// failed a direct upload.
//
// - Use propagatedSetLevelTags to configure tags that propagate to all objects
// in your store.
//
// For more information, see [Creating a HealthOmics sequence store] in the Amazon Web Services HealthOmics User Guide.
//
// [Creating a HealthOmics sequence store]: https://docs.aws.amazon.com/omics/latest/dev/create-sequence-store.html
func omics_CreateSequenceStore(cfg aws.Config, client *omics.Client) {
	input := &omics.CreateSequenceStoreInput{
		// Name: *string, // Required
	}

	if len(_omicsName) > 0 {
		input.Name = aws.String(_omicsName)
	}
	if len(_omicsClientToken) > 0 {
		input.ClientToken = aws.String(_omicsClientToken)
	}
	if len(_omicsDescription) > 0 {
		input.Description = aws.String(_omicsDescription)
	}
	if len(_omicsETagAlgorithmFamily) > 0 {
		if err := assignInputField(input, "ETagAlgorithmFamily", _omicsETagAlgorithmFamily); err != nil {
			log.Errorf("invalid --etag-algorithm-family: %s", err.Error())
			return
		}
	}
	if len(_omicsFallbackLocation) > 0 {
		input.FallbackLocation = aws.String(_omicsFallbackLocation)
	}
	if len(_omicsPropagatedSetLevelTags) > 0 {
		input.PropagatedSetLevelTags = append([]string(nil), _omicsPropagatedSetLevelTags...)
	}
	if len(_omicsS3AccessConfig) > 0 {
		if err := assignInputField(input, "S3AccessConfig", _omicsS3AccessConfig); err != nil {
			log.Errorf("invalid --s3-access-config: %s", err.Error())
			return
		}
	}
	if len(_omicsSseConfig) > 0 {
		if err := assignInputField(input, "SseConfig", _omicsSseConfig); err != nil {
			log.Errorf("invalid --sse-config: %s", err.Error())
			return
		}
	}
	if len(_omicsTags) > 0 {
		if err := assignInputField(input, "Tags", _omicsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSequenceStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a cross-account shared resource. The resource owner makes an offer to
// share the resource with the principal subscriber (an AWS user with a different
// account than the resource owner).
//
// The following resources support cross-account sharing:
//
// - HealthOmics variant stores
//
// - HealthOmics annotation stores
//
// - Private workflows
func omics_CreateShare(cfg aws.Config, client *omics.Client) {
	input := &omics.CreateShareInput{
		// PrincipalSubscriber: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_omicsPrincipalSubscriber) > 0 {
		input.PrincipalSubscriber = aws.String(_omicsPrincipalSubscriber)
	}
	if len(_omicsResourceArn) > 0 {
		input.ResourceArn = aws.String(_omicsResourceArn)
	}
	if len(_omicsShareName) > 0 {
		input.ShareName = aws.String(_omicsShareName)
	}

	if resp, err := client.CreateShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Amazon Web Services HealthOmics variant stores and annotation stores will no
// longer be open to new customers starting November 7, 2025. If you would like to
// use variant stores or annotation stores, sign up prior to that date. Existing
// customers can continue to use the service as normal. For more information, see [Amazon Web Services HealthOmics variant store and annotation store availability change].
//
// Creates a variant store.
//
// [Amazon Web Services HealthOmics variant store and annotation store availability change]: https://docs.aws.amazon.com/omics/latest/dev/variant-store-availability-change.html
func omics_CreateVariantStore(cfg aws.Config, client *omics.Client) {
	input := &omics.CreateVariantStoreInput{
		// Reference: types.ReferenceItem, // Required
	}

	if len(_omicsReference) > 0 {
		if err := assignInputField(input, "Reference", _omicsReference); err != nil {
			log.Errorf("invalid --reference: %s", err.Error())
			return
		}
	}
	if len(_omicsDescription) > 0 {
		input.Description = aws.String(_omicsDescription)
	}
	if len(_omicsName) > 0 {
		input.Name = aws.String(_omicsName)
	}
	if len(_omicsSseConfig) > 0 {
		if err := assignInputField(input, "SseConfig", _omicsSseConfig); err != nil {
			log.Errorf("invalid --sse-config: %s", err.Error())
			return
		}
	}
	if len(_omicsTags) > 0 {
		if err := assignInputField(input, "Tags", _omicsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateVariantStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a private workflow. Before you create a private workflow, you must
// create and configure these required resources:
//
// - Workflow definition file: A workflow definition file written in WDL,
// Nextflow, or CWL. The workflow definition specifies the inputs and outputs for
// runs that use the workflow. It also includes specifications for the runs and run
// tasks for your workflow, including compute and memory requirements. The workflow
// definition file must be in .zip format. For more information, see [Workflow definition files]in Amazon
// Web Services HealthOmics.
//
// - You can use Amazon Q CLI to build and validate your workflow definition
// files in WDL, Nextflow, and CWL. For more information, see [Example prompts for Amazon Q CLI]and the [Amazon Web Services HealthOmics Agentic generative AI tutorial]on GitHub.
//
// - (Optional) Parameter template file: A parameter template file written in
// JSON. Create the file to define the run parameters, or Amazon Web Services
// HealthOmics generates the parameter template for you. For more information, see [Parameter template files for HealthOmics workflows]
// .
//
// - ECR container images: Create container images for the workflow in a private
// ECR repository, or synchronize images from a supported upstream registry with
// your Amazon ECR private repository.
//
// - (Optional) Sentieon licenses: Request a Sentieon license to use the
// Sentieon software in private workflows.
//
// For more information, see [Creating or updating a private workflow in Amazon Web Services HealthOmics] in the Amazon Web Services HealthOmics User Guide.
//
// [Example prompts for Amazon Q CLI]: https://docs.aws.amazon.com/omics/latest/dev/getting-started.html#omics-q-prompts
// [Workflow definition files]: https://docs.aws.amazon.com/omics/latest/dev/workflow-definition-files.html
// [Parameter template files for HealthOmics workflows]: https://docs.aws.amazon.com/omics/latest/dev/parameter-templates.html
// [Creating or updating a private workflow in Amazon Web Services HealthOmics]: https://docs.aws.amazon.com/omics/latest/dev/creating-private-workflows.html
// [Amazon Web Services HealthOmics Agentic generative AI tutorial]: https://github.com/aws-samples/aws-healthomics-tutorials/tree/main/generative-ai
func omics_CreateWorkflow(cfg aws.Config, client *omics.Client) {
	input := &omics.CreateWorkflowInput{
		// RequestId: *string, // Required
	}

	if len(_omicsRequestId) > 0 {
		input.RequestId = aws.String(_omicsRequestId)
	}
	if len(_omicsAccelerators) > 0 {
		if err := assignInputField(input, "Accelerators", _omicsAccelerators); err != nil {
			log.Errorf("invalid --accelerators: %s", err.Error())
			return
		}
	}
	if len(_omicsContainerRegistryMap) > 0 {
		if err := assignInputField(input, "ContainerRegistryMap", _omicsContainerRegistryMap); err != nil {
			log.Errorf("invalid --container-registry-map: %s", err.Error())
			return
		}
	}
	if len(_omicsContainerRegistryMapUri) > 0 {
		input.ContainerRegistryMapUri = aws.String(_omicsContainerRegistryMapUri)
	}
	if len(_omicsDefinitionRepository) > 0 {
		if err := assignInputField(input, "DefinitionRepository", _omicsDefinitionRepository); err != nil {
			log.Errorf("invalid --definition-repository: %s", err.Error())
			return
		}
	}
	if len(_omicsDefinitionUri) > 0 {
		input.DefinitionUri = aws.String(_omicsDefinitionUri)
	}
	if len(_omicsDefinitionZip) > 0 {
		if err := assignInputField(input, "DefinitionZip", _omicsDefinitionZip); err != nil {
			log.Errorf("invalid --definition-zip: %s", err.Error())
			return
		}
	}
	if len(_omicsDescription) > 0 {
		input.Description = aws.String(_omicsDescription)
	}
	if len(_omicsEngine) > 0 {
		if err := assignInputField(input, "Engine", _omicsEngine); err != nil {
			log.Errorf("invalid --engine: %s", err.Error())
			return
		}
	}
	if len(_omicsMain) > 0 {
		input.Main = aws.String(_omicsMain)
	}
	if len(_omicsName) > 0 {
		input.Name = aws.String(_omicsName)
	}
	if len(_omicsParameterTemplate) > 0 {
		if err := assignInputField(input, "ParameterTemplate", _omicsParameterTemplate); err != nil {
			log.Errorf("invalid --parameter-template: %s", err.Error())
			return
		}
	}
	if len(_omicsParameterTemplatePath) > 0 {
		input.ParameterTemplatePath = aws.String(_omicsParameterTemplatePath)
	}
	if len(_omicsReadmeMarkdown) > 0 {
		input.ReadmeMarkdown = aws.String(_omicsReadmeMarkdown)
	}
	if len(_omicsReadmePath) > 0 {
		input.ReadmePath = aws.String(_omicsReadmePath)
	}
	if len(_omicsReadmeUri) > 0 {
		input.ReadmeUri = aws.String(_omicsReadmeUri)
	}
	if len(_omicsStorageCapacity) > 0 {
		if err := assignInputField(input, "StorageCapacity", _omicsStorageCapacity); err != nil {
			log.Errorf("invalid --storage-capacity: %s", err.Error())
			return
		}
	}
	if len(_omicsStorageType) > 0 {
		if err := assignInputField(input, "StorageType", _omicsStorageType); err != nil {
			log.Errorf("invalid --storage-type: %s", err.Error())
			return
		}
	}
	if len(_omicsTags) > 0 {
		if err := assignInputField(input, "Tags", _omicsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_omicsWorkflowBucketOwnerId) > 0 {
		input.WorkflowBucketOwnerId = aws.String(_omicsWorkflowBucketOwnerId)
	}

	if resp, err := client.CreateWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new workflow version for the workflow that you specify with the
// workflowId parameter.
//
// When you create a new version of a workflow, you need to specify the
// configuration for the new version. It doesn't inherit any configuration values
// from the workflow.
//
// Provide a version name that is unique for this workflow. You cannot change the
// name after HealthOmics creates the version.
//
// Don't include any personally identifiable information (PII) in the version
// name. Version names appear in the workflow version ARN.
//
// For more information, see [Workflow versioning in Amazon Web Services HealthOmics] in the Amazon Web Services HealthOmics User Guide.
//
// [Workflow versioning in Amazon Web Services HealthOmics]: https://docs.aws.amazon.com/omics/latest/dev/workflow-versions.html
func omics_CreateWorkflowVersion(cfg aws.Config, client *omics.Client) {
	input := &omics.CreateWorkflowVersionInput{
		// RequestId: *string, // Required
		// VersionName: *string, // Required
		// WorkflowId: *string, // Required
	}

	if len(_omicsRequestId) > 0 {
		input.RequestId = aws.String(_omicsRequestId)
	}
	if len(_omicsVersionName) > 0 {
		input.VersionName = aws.String(_omicsVersionName)
	}
	if len(_omicsWorkflowId) > 0 {
		input.WorkflowId = aws.String(_omicsWorkflowId)
	}
	if len(_omicsAccelerators) > 0 {
		if err := assignInputField(input, "Accelerators", _omicsAccelerators); err != nil {
			log.Errorf("invalid --accelerators: %s", err.Error())
			return
		}
	}
	if len(_omicsContainerRegistryMap) > 0 {
		if err := assignInputField(input, "ContainerRegistryMap", _omicsContainerRegistryMap); err != nil {
			log.Errorf("invalid --container-registry-map: %s", err.Error())
			return
		}
	}
	if len(_omicsContainerRegistryMapUri) > 0 {
		input.ContainerRegistryMapUri = aws.String(_omicsContainerRegistryMapUri)
	}
	if len(_omicsDefinitionRepository) > 0 {
		if err := assignInputField(input, "DefinitionRepository", _omicsDefinitionRepository); err != nil {
			log.Errorf("invalid --definition-repository: %s", err.Error())
			return
		}
	}
	if len(_omicsDefinitionUri) > 0 {
		input.DefinitionUri = aws.String(_omicsDefinitionUri)
	}
	if len(_omicsDefinitionZip) > 0 {
		if err := assignInputField(input, "DefinitionZip", _omicsDefinitionZip); err != nil {
			log.Errorf("invalid --definition-zip: %s", err.Error())
			return
		}
	}
	if len(_omicsDescription) > 0 {
		input.Description = aws.String(_omicsDescription)
	}
	if len(_omicsEngine) > 0 {
		if err := assignInputField(input, "Engine", _omicsEngine); err != nil {
			log.Errorf("invalid --engine: %s", err.Error())
			return
		}
	}
	if len(_omicsMain) > 0 {
		input.Main = aws.String(_omicsMain)
	}
	if len(_omicsParameterTemplate) > 0 {
		if err := assignInputField(input, "ParameterTemplate", _omicsParameterTemplate); err != nil {
			log.Errorf("invalid --parameter-template: %s", err.Error())
			return
		}
	}
	if len(_omicsParameterTemplatePath) > 0 {
		input.ParameterTemplatePath = aws.String(_omicsParameterTemplatePath)
	}
	if len(_omicsReadmeMarkdown) > 0 {
		input.ReadmeMarkdown = aws.String(_omicsReadmeMarkdown)
	}
	if len(_omicsReadmePath) > 0 {
		input.ReadmePath = aws.String(_omicsReadmePath)
	}
	if len(_omicsReadmeUri) > 0 {
		input.ReadmeUri = aws.String(_omicsReadmeUri)
	}
	if len(_omicsStorageCapacity) > 0 {
		if err := assignInputField(input, "StorageCapacity", _omicsStorageCapacity); err != nil {
			log.Errorf("invalid --storage-capacity: %s", err.Error())
			return
		}
	}
	if len(_omicsStorageType) > 0 {
		if err := assignInputField(input, "StorageType", _omicsStorageType); err != nil {
			log.Errorf("invalid --storage-type: %s", err.Error())
			return
		}
	}
	if len(_omicsTags) > 0 {
		if err := assignInputField(input, "Tags", _omicsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_omicsWorkflowBucketOwnerId) > 0 {
		input.WorkflowBucketOwnerId = aws.String(_omicsWorkflowBucketOwnerId)
	}

	if resp, err := client.CreateWorkflowVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Amazon Web Services HealthOmics variant stores and annotation stores will no
// longer be open to new customers starting November 7, 2025. If you would like to
// use variant stores or annotation stores, sign up prior to that date. Existing
// customers can continue to use the service as normal. For more information, see [Amazon Web Services HealthOmics variant store and annotation store availability change].
//
// Deletes an annotation store.
//
// [Amazon Web Services HealthOmics variant store and annotation store availability change]: https://docs.aws.amazon.com/omics/latest/dev/variant-store-availability-change.html
func omics_DeleteAnnotationStore(cfg aws.Config, client *omics.Client) {
	input := &omics.DeleteAnnotationStoreInput{
		// Name: *string, // Required
	}

	if len(_omicsName) > 0 {
		input.Name = aws.String(_omicsName)
	}
	if len(_omicsForce) > 0 {
		if err := assignInputField(input, "Force", _omicsForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteAnnotationStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes one or multiple versions of an annotation store.
func omics_DeleteAnnotationStoreVersions(cfg aws.Config, client *omics.Client) {
	input := &omics.DeleteAnnotationStoreVersionsInput{
		// Name: *string, // Required
		// Versions: []string, // Required
	}

	if len(_omicsName) > 0 {
		input.Name = aws.String(_omicsName)
	}
	if len(_omicsVersions) > 0 {
		input.Versions = append([]string(nil), _omicsVersions...)
	}
	if len(_omicsForce) > 0 {
		if err := assignInputField(input, "Force", _omicsForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteAnnotationStoreVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a reference genome and returns a response with no body if the operation
// is successful. The read set associated with the reference genome must first be
// deleted before deleting the reference genome. After the reference genome is
// deleted, you can delete the reference store using the DeleteReferenceStore API
// operation.
//
// For more information, see [Deleting HealthOmics reference and sequence stores] in the Amazon Web Services HealthOmics User Guide.
//
// [Deleting HealthOmics reference and sequence stores]: https://docs.aws.amazon.com/omics/latest/dev/deleting-reference-and-sequence-stores.html
func omics_DeleteReference(cfg aws.Config, client *omics.Client) {
	input := &omics.DeleteReferenceInput{
		// Id: *string, // Required
		// ReferenceStoreId: *string, // Required
	}

	if len(_omicsId) > 0 {
		input.Id = aws.String(_omicsId)
	}
	if len(_omicsReferenceStoreId) > 0 {
		input.ReferenceStoreId = aws.String(_omicsReferenceStoreId)
	}

	if resp, err := client.DeleteReference(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a reference store and returns a response with no body if the operation
// is successful. You can only delete a reference store when it does not contain
// any reference genomes. To empty a reference store, use DeleteReference .
//
// For more information about your workflow status, see [Deleting HealthOmics reference and sequence stores] in the Amazon Web
// Services HealthOmics User Guide.
//
// [Deleting HealthOmics reference and sequence stores]: https://docs.aws.amazon.com/omics/latest/dev/deleting-reference-and-sequence-stores.html
func omics_DeleteReferenceStore(cfg aws.Config, client *omics.Client) {
	input := &omics.DeleteReferenceStoreInput{
		// Id: *string, // Required
	}

	if len(_omicsId) > 0 {
		input.Id = aws.String(_omicsId)
	}

	if resp, err := client.DeleteReferenceStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a run and returns a response with no body if the operation is
// successful. You can only delete a run that has reached a COMPLETED , FAILED , or
// CANCELLED stage. A completed run has delivered an output, or was cancelled and
// resulted in no output. When you delete a run, only the metadata associated with
// the run is deleted. The run outputs remain in Amazon S3 and logs remain in
// CloudWatch.
//
// To verify that the workflow is deleted:
//
// - Use ListRuns to confirm the workflow no longer appears in the list.
//
// - Use GetRun to verify the workflow cannot be found.
func omics_DeleteRun(cfg aws.Config, client *omics.Client) {
	input := &omics.DeleteRunInput{
		// Id: *string, // Required
	}

	if len(_omicsId) > 0 {
		input.Id = aws.String(_omicsId)
	}

	if resp, err := client.DeleteRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a run cache and returns a response with no body if the operation is
// successful. This action removes the cache metadata stored in the service
// account, but does not delete the data in Amazon S3. You can access the cache
// data in Amazon S3, for inspection or to troubleshoot issues. You can remove old
// cache data using standard S3 Delete operations.
//
// For more information, see [Deleting a run cache] in the Amazon Web Services HealthOmics User Guide.
//
// [Deleting a run cache]: https://docs.aws.amazon.com/omics/latest/dev/workflow-cache-delete.html
func omics_DeleteRunCache(cfg aws.Config, client *omics.Client) {
	input := &omics.DeleteRunCacheInput{
		// Id: *string, // Required
	}

	if len(_omicsId) > 0 {
		input.Id = aws.String(_omicsId)
	}

	if resp, err := client.DeleteRunCache(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a run group and returns a response with no body if the operation is
// successful.
//
// To verify that the run group is deleted:
//
// - Use ListRunGroups to confirm the workflow no longer appears in the list.
//
// - Use GetRunGroup to verify the workflow cannot be found.
func omics_DeleteRunGroup(cfg aws.Config, client *omics.Client) {
	input := &omics.DeleteRunGroupInput{
		// Id: *string, // Required
	}

	if len(_omicsId) > 0 {
		input.Id = aws.String(_omicsId)
	}

	if resp, err := client.DeleteRunGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an access policy for the specified store.
func omics_DeleteS3AccessPolicy(cfg aws.Config, client *omics.Client) {
	input := &omics.DeleteS3AccessPolicyInput{
		// S3AccessPointArn: *string, // Required
	}

	if len(_omicsS3AccessPointArn) > 0 {
		input.S3AccessPointArn = aws.String(_omicsS3AccessPointArn)
	}

	if resp, err := client.DeleteS3AccessPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a sequence store and returns a response with no body if the operation
// is successful. You can only delete a sequence store when it does not contain any
// read sets.
//
// Use the BatchDeleteReadSet API operation to ensure that all read sets in the
// sequence store are deleted. When a sequence store is deleted, all tags
// associated with the store are also deleted.
//
// For more information, see [Deleting HealthOmics reference and sequence stores] in the Amazon Web Services HealthOmics User Guide.
//
// [Deleting HealthOmics reference and sequence stores]: https://docs.aws.amazon.com/omics/latest/dev/deleting-reference-and-sequence-stores.html
func omics_DeleteSequenceStore(cfg aws.Config, client *omics.Client) {
	input := &omics.DeleteSequenceStoreInput{
		// Id: *string, // Required
	}

	if len(_omicsId) > 0 {
		input.Id = aws.String(_omicsId)
	}

	if resp, err := client.DeleteSequenceStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a resource share. If you are the resource owner, the subscriber will no
// longer have access to the shared resource. If you are the subscriber, this
// operation deletes your access to the share.
func omics_DeleteShare(cfg aws.Config, client *omics.Client) {
	input := &omics.DeleteShareInput{
		// ShareId: *string, // Required
	}

	if len(_omicsShareId) > 0 {
		input.ShareId = aws.String(_omicsShareId)
	}

	if resp, err := client.DeleteShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Amazon Web Services HealthOmics variant stores and annotation stores will no
// longer be open to new customers starting November 7, 2025. If you would like to
// use variant stores or annotation stores, sign up prior to that date. Existing
// customers can continue to use the service as normal. For more information, see [Amazon Web Services HealthOmics variant store and annotation store availability change].
//
// Deletes a variant store.
//
// [Amazon Web Services HealthOmics variant store and annotation store availability change]: https://docs.aws.amazon.com/omics/latest/dev/variant-store-availability-change.html
func omics_DeleteVariantStore(cfg aws.Config, client *omics.Client) {
	input := &omics.DeleteVariantStoreInput{
		// Name: *string, // Required
	}

	if len(_omicsName) > 0 {
		input.Name = aws.String(_omicsName)
	}
	if len(_omicsForce) > 0 {
		if err := assignInputField(input, "Force", _omicsForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteVariantStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a workflow by specifying its ID. This operation returns a response with
// no body if the deletion is successful.
//
// To verify that the workflow is deleted:
//
// - Use ListWorkflows to confirm the workflow no longer appears in the list.
//
// - Use GetWorkflow to verify the workflow cannot be found.
func omics_DeleteWorkflow(cfg aws.Config, client *omics.Client) {
	input := &omics.DeleteWorkflowInput{
		// Id: *string, // Required
	}

	if len(_omicsId) > 0 {
		input.Id = aws.String(_omicsId)
	}

	if resp, err := client.DeleteWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a workflow version. Deleting a workflow version doesn't affect any
// ongoing runs that are using the workflow version.
//
// For more information, see [Workflow versioning in Amazon Web Services HealthOmics] in the Amazon Web Services HealthOmics User Guide.
//
// [Workflow versioning in Amazon Web Services HealthOmics]: https://docs.aws.amazon.com/omics/latest/dev/workflow-versions.html
func omics_DeleteWorkflowVersion(cfg aws.Config, client *omics.Client) {
	input := &omics.DeleteWorkflowVersionInput{
		// VersionName: *string, // Required
		// WorkflowId: *string, // Required
	}

	if len(_omicsVersionName) > 0 {
		input.VersionName = aws.String(_omicsVersionName)
	}
	if len(_omicsWorkflowId) > 0 {
		input.WorkflowId = aws.String(_omicsWorkflowId)
	}

	if resp, err := client.DeleteWorkflowVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Amazon Web Services HealthOmics variant stores and annotation stores will no
// longer be open to new customers starting November 7, 2025. If you would like to
// use variant stores or annotation stores, sign up prior to that date. Existing
// customers can continue to use the service as normal. For more information, see [Amazon Web Services HealthOmics variant store and annotation store availability change].
//
// Gets information about an annotation import job.
//
// [Amazon Web Services HealthOmics variant store and annotation store availability change]: https://docs.aws.amazon.com/omics/latest/dev/variant-store-availability-change.html
func omics_GetAnnotationImportJob(cfg aws.Config, client *omics.Client) {
	input := &omics.GetAnnotationImportJobInput{
		// JobId: *string, // Required
	}

	if len(_omicsJobId) > 0 {
		input.JobId = aws.String(_omicsJobId)
	}

	if resp, err := client.GetAnnotationImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Amazon Web Services HealthOmics variant stores and annotation stores will no
// longer be open to new customers starting November 7, 2025. If you would like to
// use variant stores or annotation stores, sign up prior to that date. Existing
// customers can continue to use the service as normal. For more information, see [Amazon Web Services HealthOmics variant store and annotation store availability change].
//
// Gets information about an annotation store.
//
// [Amazon Web Services HealthOmics variant store and annotation store availability change]: https://docs.aws.amazon.com/omics/latest/dev/variant-store-availability-change.html
func omics_GetAnnotationStore(cfg aws.Config, client *omics.Client) {
	input := &omics.GetAnnotationStoreInput{
		// Name: *string, // Required
	}

	if len(_omicsName) > 0 {
		input.Name = aws.String(_omicsName)
	}

	if resp, err := client.GetAnnotationStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the metadata for an annotation store version.
func omics_GetAnnotationStoreVersion(cfg aws.Config, client *omics.Client) {
	input := &omics.GetAnnotationStoreVersionInput{
		// Name: *string, // Required
		// VersionName: *string, // Required
	}

	if len(_omicsName) > 0 {
		input.Name = aws.String(_omicsName)
	}
	if len(_omicsVersionName) > 0 {
		input.VersionName = aws.String(_omicsVersionName)
	}

	if resp, err := client.GetAnnotationStoreVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information from parts of a read set and returns the read
// set in the same format that it was uploaded. You must have read sets uploaded to
// your sequence store in order to run this operation.
func omics_GetReadSet(cfg aws.Config, client *omics.Client) {
	input := &omics.GetReadSetInput{
		// Id: *string, // Required
		// PartNumber: *int32, // Required
		// SequenceStoreId: *string, // Required
	}

	if len(_omicsId) > 0 {
		input.Id = aws.String(_omicsId)
	}
	if len(_omicsPartNumber) > 0 {
		if err := assignInputField(input, "PartNumber", _omicsPartNumber); err != nil {
			log.Errorf("invalid --part-number: %s", err.Error())
			return
		}
	}
	if len(_omicsSequenceStoreId) > 0 {
		input.SequenceStoreId = aws.String(_omicsSequenceStoreId)
	}
	if len(_omicsFile) > 0 {
		if err := assignInputField(input, "File", _omicsFile); err != nil {
			log.Errorf("invalid --file: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetReadSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns detailed information about the status of a read set activation job in
// JSON format.
func omics_GetReadSetActivationJob(cfg aws.Config, client *omics.Client) {
	input := &omics.GetReadSetActivationJobInput{
		// Id: *string, // Required
		// SequenceStoreId: *string, // Required
	}

	if len(_omicsId) > 0 {
		input.Id = aws.String(_omicsId)
	}
	if len(_omicsSequenceStoreId) > 0 {
		input.SequenceStoreId = aws.String(_omicsSequenceStoreId)
	}

	if resp, err := client.GetReadSetActivationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves status information about a read set export job and returns the data
// in JSON format. Use this operation to actively monitor the progress of an export
// job.
func omics_GetReadSetExportJob(cfg aws.Config, client *omics.Client) {
	input := &omics.GetReadSetExportJobInput{
		// Id: *string, // Required
		// SequenceStoreId: *string, // Required
	}

	if len(_omicsId) > 0 {
		input.Id = aws.String(_omicsId)
	}
	if len(_omicsSequenceStoreId) > 0 {
		input.SequenceStoreId = aws.String(_omicsSequenceStoreId)
	}

	if resp, err := client.GetReadSetExportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets detailed and status information about a read set import job and returns
// the data in JSON format.
func omics_GetReadSetImportJob(cfg aws.Config, client *omics.Client) {
	input := &omics.GetReadSetImportJobInput{
		// Id: *string, // Required
		// SequenceStoreId: *string, // Required
	}

	if len(_omicsId) > 0 {
		input.Id = aws.String(_omicsId)
	}
	if len(_omicsSequenceStoreId) > 0 {
		input.SequenceStoreId = aws.String(_omicsSequenceStoreId)
	}

	if resp, err := client.GetReadSetImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the metadata for a read set from a sequence store in JSON format.
// This operation does not return tags. To retrieve the list of tags for a read
// set, use the ListTagsForResource API operation.
func omics_GetReadSetMetadata(cfg aws.Config, client *omics.Client) {
	input := &omics.GetReadSetMetadataInput{
		// Id: *string, // Required
		// SequenceStoreId: *string, // Required
	}

	if len(_omicsId) > 0 {
		input.Id = aws.String(_omicsId)
	}
	if len(_omicsSequenceStoreId) > 0 {
		input.SequenceStoreId = aws.String(_omicsSequenceStoreId)
	}

	if resp, err := client.GetReadSetMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Downloads parts of data from a reference genome and returns the reference file
// in the same format that it was uploaded.
//
// For more information, see [Creating a HealthOmics reference store] in the Amazon Web Services HealthOmics User Guide.
//
// [Creating a HealthOmics reference store]: https://docs.aws.amazon.com/omics/latest/dev/create-reference-store.html
func omics_GetReference(cfg aws.Config, client *omics.Client) {
	input := &omics.GetReferenceInput{
		// Id: *string, // Required
		// PartNumber: *int32, // Required
		// ReferenceStoreId: *string, // Required
	}

	if len(_omicsId) > 0 {
		input.Id = aws.String(_omicsId)
	}
	if len(_omicsPartNumber) > 0 {
		if err := assignInputField(input, "PartNumber", _omicsPartNumber); err != nil {
			log.Errorf("invalid --part-number: %s", err.Error())
			return
		}
	}
	if len(_omicsReferenceStoreId) > 0 {
		input.ReferenceStoreId = aws.String(_omicsReferenceStoreId)
	}
	if len(_omicsFile) > 0 {
		if err := assignInputField(input, "File", _omicsFile); err != nil {
			log.Errorf("invalid --file: %s", err.Error())
			return
		}
	}
	if len(_omicsRange) > 0 {
		input.Range = aws.String(_omicsRange)
	}

	if resp, err := client.GetReference(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Monitors the status of a reference import job. This operation can be called
// after calling the StartReferenceImportJob operation.
func omics_GetReferenceImportJob(cfg aws.Config, client *omics.Client) {
	input := &omics.GetReferenceImportJobInput{
		// Id: *string, // Required
		// ReferenceStoreId: *string, // Required
	}

	if len(_omicsId) > 0 {
		input.Id = aws.String(_omicsId)
	}
	if len(_omicsReferenceStoreId) > 0 {
		input.ReferenceStoreId = aws.String(_omicsReferenceStoreId)
	}

	if resp, err := client.GetReferenceImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves metadata for a reference genome. This operation returns the number of
// parts, part size, and MD5 of an entire file. This operation does not return
// tags. To retrieve the list of tags for a read set, use the ListTagsForResource
// API operation.
func omics_GetReferenceMetadata(cfg aws.Config, client *omics.Client) {
	input := &omics.GetReferenceMetadataInput{
		// Id: *string, // Required
		// ReferenceStoreId: *string, // Required
	}

	if len(_omicsId) > 0 {
		input.Id = aws.String(_omicsId)
	}
	if len(_omicsReferenceStoreId) > 0 {
		input.ReferenceStoreId = aws.String(_omicsReferenceStoreId)
	}

	if resp, err := client.GetReferenceMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a reference store.
func omics_GetReferenceStore(cfg aws.Config, client *omics.Client) {
	input := &omics.GetReferenceStoreInput{
		// Id: *string, // Required
	}

	if len(_omicsId) > 0 {
		input.Id = aws.String(_omicsId)
	}

	if resp, err := client.GetReferenceStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets detailed information about a specific run using its ID.
// Amazon Web Services HealthOmics stores a configurable number of runs, as
// determined by service limits, that are available to the console and API. If
// GetRun does not return the requested run, you can find all run logs in the
// CloudWatch logs. For more information about viewing the run logs, see [CloudWatch logs]in the
// Amazon Web Services HealthOmics User Guide.
//
// [CloudWatch logs]: https://docs.aws.amazon.com/omics/latest/dev/monitoring-cloudwatch-logs.html
func omics_GetRun(cfg aws.Config, client *omics.Client) {
	input := &omics.GetRunInput{
		// Id: *string, // Required
	}

	if len(_omicsId) > 0 {
		input.Id = aws.String(_omicsId)
	}
	if len(_omicsExport) > 0 {
		if err := assignInputField(input, "Export", _omicsExport); err != nil {
			log.Errorf("invalid --export: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about the specified run cache using its ID.
// For more information, see [Call caching for Amazon Web Services HealthOmics runs] in the Amazon Web Services HealthOmics User Guide.
//
// [Call caching for Amazon Web Services HealthOmics runs]: https://docs.aws.amazon.com/omics/latest/dev/workflows-call-caching.html
func omics_GetRunCache(cfg aws.Config, client *omics.Client) {
	input := &omics.GetRunCacheInput{
		// Id: *string, // Required
	}

	if len(_omicsId) > 0 {
		input.Id = aws.String(_omicsId)
	}

	if resp, err := client.GetRunCache(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a run group and returns its metadata.
func omics_GetRunGroup(cfg aws.Config, client *omics.Client) {
	input := &omics.GetRunGroupInput{
		// Id: *string, // Required
	}

	if len(_omicsId) > 0 {
		input.Id = aws.String(_omicsId)
	}

	if resp, err := client.GetRunGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets detailed information about a run task using its ID.
func omics_GetRunTask(cfg aws.Config, client *omics.Client) {
	input := &omics.GetRunTaskInput{
		// Id: *string, // Required
		// TaskId: *string, // Required
	}

	if len(_omicsId) > 0 {
		input.Id = aws.String(_omicsId)
	}
	if len(_omicsTaskId) > 0 {
		input.TaskId = aws.String(_omicsTaskId)
	}

	if resp, err := client.GetRunTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about an access policy on a given store.
func omics_GetS3AccessPolicy(cfg aws.Config, client *omics.Client) {
	input := &omics.GetS3AccessPolicyInput{
		// S3AccessPointArn: *string, // Required
	}

	if len(_omicsS3AccessPointArn) > 0 {
		input.S3AccessPointArn = aws.String(_omicsS3AccessPointArn)
	}

	if resp, err := client.GetS3AccessPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves metadata for a sequence store using its ID and returns it in JSON
// format.
func omics_GetSequenceStore(cfg aws.Config, client *omics.Client) {
	input := &omics.GetSequenceStoreInput{
		// Id: *string, // Required
	}

	if len(_omicsId) > 0 {
		input.Id = aws.String(_omicsId)
	}

	if resp, err := client.GetSequenceStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the metadata for the specified resource share.
func omics_GetShare(cfg aws.Config, client *omics.Client) {
	input := &omics.GetShareInput{
		// ShareId: *string, // Required
	}

	if len(_omicsShareId) > 0 {
		input.ShareId = aws.String(_omicsShareId)
	}

	if resp, err := client.GetShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Amazon Web Services HealthOmics variant stores and annotation stores will no
// longer be open to new customers starting November 7, 2025. If you would like to
// use variant stores or annotation stores, sign up prior to that date. Existing
// customers can continue to use the service as normal. For more information, see [Amazon Web Services HealthOmics variant store and annotation store availability change].
//
// Gets information about a variant import job.
//
// [Amazon Web Services HealthOmics variant store and annotation store availability change]: https://docs.aws.amazon.com/omics/latest/dev/variant-store-availability-change.html
func omics_GetVariantImportJob(cfg aws.Config, client *omics.Client) {
	input := &omics.GetVariantImportJobInput{
		// JobId: *string, // Required
	}

	if len(_omicsJobId) > 0 {
		input.JobId = aws.String(_omicsJobId)
	}

	if resp, err := client.GetVariantImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Amazon Web Services HealthOmics variant stores and annotation stores will no
// longer be open to new customers starting November 7, 2025. If you would like to
// use variant stores or annotation stores, sign up prior to that date. Existing
// customers can continue to use the service as normal. For more information, see [Amazon Web Services HealthOmics variant store and annotation store availability change].
//
// Gets information about a variant store.
//
// [Amazon Web Services HealthOmics variant store and annotation store availability change]: https://docs.aws.amazon.com/omics/latest/dev/variant-store-availability-change.html
func omics_GetVariantStore(cfg aws.Config, client *omics.Client) {
	input := &omics.GetVariantStoreInput{
		// Name: *string, // Required
	}

	if len(_omicsName) > 0 {
		input.Name = aws.String(_omicsName)
	}

	if resp, err := client.GetVariantStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets all information about a workflow using its ID.
// If a workflow is shared with you, you cannot export the workflow.
//
// For more information about your workflow status, see [Verify the workflow status] in the Amazon Web
// Services HealthOmics User Guide.
//
// [Verify the workflow status]: https://docs.aws.amazon.com/omics/latest/dev/using-get-workflow.html
func omics_GetWorkflow(cfg aws.Config, client *omics.Client) {
	input := &omics.GetWorkflowInput{
		// Id: *string, // Required
	}

	if len(_omicsId) > 0 {
		input.Id = aws.String(_omicsId)
	}
	if len(_omicsExport) > 0 {
		if err := assignInputField(input, "Export", _omicsExport); err != nil {
			log.Errorf("invalid --export: %s", err.Error())
			return
		}
	}
	if len(_omicsType) > 0 {
		if err := assignInputField(input, "Type", _omicsType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_omicsWorkflowOwnerId) > 0 {
		input.WorkflowOwnerId = aws.String(_omicsWorkflowOwnerId)
	}

	if resp, err := client.GetWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a workflow version. For more information, see [Workflow versioning in Amazon Web Services HealthOmics] in the
// Amazon Web Services HealthOmics User Guide.
//
// [Workflow versioning in Amazon Web Services HealthOmics]: https://docs.aws.amazon.com/omics/latest/dev/workflow-versions.html
func omics_GetWorkflowVersion(cfg aws.Config, client *omics.Client) {
	input := &omics.GetWorkflowVersionInput{
		// VersionName: *string, // Required
		// WorkflowId: *string, // Required
	}

	if len(_omicsVersionName) > 0 {
		input.VersionName = aws.String(_omicsVersionName)
	}
	if len(_omicsWorkflowId) > 0 {
		input.WorkflowId = aws.String(_omicsWorkflowId)
	}
	if len(_omicsExport) > 0 {
		if err := assignInputField(input, "Export", _omicsExport); err != nil {
			log.Errorf("invalid --export: %s", err.Error())
			return
		}
	}
	if len(_omicsType) > 0 {
		if err := assignInputField(input, "Type", _omicsType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_omicsWorkflowOwnerId) > 0 {
		input.WorkflowOwnerId = aws.String(_omicsWorkflowOwnerId)
	}

	if resp, err := client.GetWorkflowVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Amazon Web Services HealthOmics variant stores and annotation stores will no
// longer be open to new customers starting November 7, 2025. If you would like to
// use variant stores or annotation stores, sign up prior to that date. Existing
// customers can continue to use the service as normal. For more information, see [Amazon Web Services HealthOmics variant store and annotation store availability change].
//
// Retrieves a list of annotation import jobs.
//
// [Amazon Web Services HealthOmics variant store and annotation store availability change]: https://docs.aws.amazon.com/omics/latest/dev/variant-store-availability-change.html
func omics_ListAnnotationImportJobs(cfg aws.Config, client *omics.Client) {
	input := &omics.ListAnnotationImportJobsInput{}

	if len(_omicsFilter) > 0 {
		if err := assignInputField(input, "Filter", _omicsFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_omicsIds) > 0 {
		input.Ids = append([]string(nil), _omicsIds...)
	}
	if len(_omicsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _omicsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_omicsNextToken) > 0 {
		input.NextToken = aws.String(_omicsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAnnotationImportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*omics.ListAnnotationImportJobsOutput
	p := omics.NewListAnnotationImportJobsPaginator(client, input)
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

// Lists the versions of an annotation store.
func omics_ListAnnotationStoreVersions(cfg aws.Config, client *omics.Client) {
	input := &omics.ListAnnotationStoreVersionsInput{
		// Name: *string, // Required
	}

	if len(_omicsName) > 0 {
		input.Name = aws.String(_omicsName)
	}
	if len(_omicsFilter) > 0 {
		if err := assignInputField(input, "Filter", _omicsFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_omicsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _omicsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_omicsNextToken) > 0 {
		input.NextToken = aws.String(_omicsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAnnotationStoreVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*omics.ListAnnotationStoreVersionsOutput
	p := omics.NewListAnnotationStoreVersionsPaginator(client, input)
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

// Amazon Web Services HealthOmics variant stores and annotation stores will no
// longer be open to new customers starting November 7, 2025. If you would like to
// use variant stores or annotation stores, sign up prior to that date. Existing
// customers can continue to use the service as normal. For more information, see [Amazon Web Services HealthOmics variant store and annotation store availability change].
//
// Retrieves a list of annotation stores.
//
// [Amazon Web Services HealthOmics variant store and annotation store availability change]: https://docs.aws.amazon.com/omics/latest/dev/variant-store-availability-change.html
func omics_ListAnnotationStores(cfg aws.Config, client *omics.Client) {
	input := &omics.ListAnnotationStoresInput{}

	if len(_omicsFilter) > 0 {
		if err := assignInputField(input, "Filter", _omicsFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_omicsIds) > 0 {
		input.Ids = append([]string(nil), _omicsIds...)
	}
	if len(_omicsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _omicsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_omicsNextToken) > 0 {
		input.NextToken = aws.String(_omicsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAnnotationStores(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*omics.ListAnnotationStoresOutput
	p := omics.NewListAnnotationStoresPaginator(client, input)
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

// Lists in-progress multipart read set uploads for a sequence store and returns
// it in a JSON formatted output. Multipart read set uploads are initiated by the
// CreateMultipartReadSetUploads API operation. This operation returns a response
// with no body when the upload is complete.
func omics_ListMultipartReadSetUploads(cfg aws.Config, client *omics.Client) {
	input := &omics.ListMultipartReadSetUploadsInput{
		// SequenceStoreId: *string, // Required
	}

	if len(_omicsSequenceStoreId) > 0 {
		input.SequenceStoreId = aws.String(_omicsSequenceStoreId)
	}
	if len(_omicsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _omicsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_omicsNextToken) > 0 {
		input.NextToken = aws.String(_omicsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMultipartReadSetUploads(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*omics.ListMultipartReadSetUploadsOutput
	p := omics.NewListMultipartReadSetUploadsPaginator(client, input)
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

// Retrieves a list of read set activation jobs and returns the metadata in a JSON
// formatted output. To extract metadata from a read set activation job, use the
// GetReadSetActivationJob API operation.
func omics_ListReadSetActivationJobs(cfg aws.Config, client *omics.Client) {
	input := &omics.ListReadSetActivationJobsInput{
		// SequenceStoreId: *string, // Required
	}

	if len(_omicsSequenceStoreId) > 0 {
		input.SequenceStoreId = aws.String(_omicsSequenceStoreId)
	}
	if len(_omicsFilter) > 0 {
		if err := assignInputField(input, "Filter", _omicsFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_omicsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _omicsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_omicsNextToken) > 0 {
		input.NextToken = aws.String(_omicsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListReadSetActivationJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*omics.ListReadSetActivationJobsOutput
	p := omics.NewListReadSetActivationJobsPaginator(client, input)
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

// Retrieves a list of read set export jobs in a JSON formatted response. This API
// operation is used to check the status of a read set export job initiated by the
// StartReadSetExportJob API operation.
func omics_ListReadSetExportJobs(cfg aws.Config, client *omics.Client) {
	input := &omics.ListReadSetExportJobsInput{
		// SequenceStoreId: *string, // Required
	}

	if len(_omicsSequenceStoreId) > 0 {
		input.SequenceStoreId = aws.String(_omicsSequenceStoreId)
	}
	if len(_omicsFilter) > 0 {
		if err := assignInputField(input, "Filter", _omicsFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_omicsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _omicsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_omicsNextToken) > 0 {
		input.NextToken = aws.String(_omicsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListReadSetExportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*omics.ListReadSetExportJobsOutput
	p := omics.NewListReadSetExportJobsPaginator(client, input)
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

// Retrieves a list of read set import jobs and returns the data in JSON format.
func omics_ListReadSetImportJobs(cfg aws.Config, client *omics.Client) {
	input := &omics.ListReadSetImportJobsInput{
		// SequenceStoreId: *string, // Required
	}

	if len(_omicsSequenceStoreId) > 0 {
		input.SequenceStoreId = aws.String(_omicsSequenceStoreId)
	}
	if len(_omicsFilter) > 0 {
		if err := assignInputField(input, "Filter", _omicsFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_omicsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _omicsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_omicsNextToken) > 0 {
		input.NextToken = aws.String(_omicsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListReadSetImportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*omics.ListReadSetImportJobsOutput
	p := omics.NewListReadSetImportJobsPaginator(client, input)
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

// Lists all parts in a multipart read set upload for a sequence store and returns
// the metadata in a JSON formatted output.
func omics_ListReadSetUploadParts(cfg aws.Config, client *omics.Client) {
	input := &omics.ListReadSetUploadPartsInput{
		// PartSource: types.ReadSetPartSource, // Required
		// SequenceStoreId: *string, // Required
		// UploadId: *string, // Required
	}

	if len(_omicsPartSource) > 0 {
		if err := assignInputField(input, "PartSource", _omicsPartSource); err != nil {
			log.Errorf("invalid --part-source: %s", err.Error())
			return
		}
	}
	if len(_omicsSequenceStoreId) > 0 {
		input.SequenceStoreId = aws.String(_omicsSequenceStoreId)
	}
	if len(_omicsUploadId) > 0 {
		input.UploadId = aws.String(_omicsUploadId)
	}
	if len(_omicsFilter) > 0 {
		if err := assignInputField(input, "Filter", _omicsFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_omicsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _omicsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_omicsNextToken) > 0 {
		input.NextToken = aws.String(_omicsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListReadSetUploadParts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*omics.ListReadSetUploadPartsOutput
	p := omics.NewListReadSetUploadPartsPaginator(client, input)
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

// Retrieves a list of read sets from a sequence store ID and returns the metadata
// in JSON format.
func omics_ListReadSets(cfg aws.Config, client *omics.Client) {
	input := &omics.ListReadSetsInput{
		// SequenceStoreId: *string, // Required
	}

	if len(_omicsSequenceStoreId) > 0 {
		input.SequenceStoreId = aws.String(_omicsSequenceStoreId)
	}
	if len(_omicsFilter) > 0 {
		if err := assignInputField(input, "Filter", _omicsFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_omicsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _omicsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_omicsNextToken) > 0 {
		input.NextToken = aws.String(_omicsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListReadSets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*omics.ListReadSetsOutput
	p := omics.NewListReadSetsPaginator(client, input)
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

// Retrieves the metadata of one or more reference import jobs for a reference
// store.
func omics_ListReferenceImportJobs(cfg aws.Config, client *omics.Client) {
	input := &omics.ListReferenceImportJobsInput{
		// ReferenceStoreId: *string, // Required
	}

	if len(_omicsReferenceStoreId) > 0 {
		input.ReferenceStoreId = aws.String(_omicsReferenceStoreId)
	}
	if len(_omicsFilter) > 0 {
		if err := assignInputField(input, "Filter", _omicsFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_omicsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _omicsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_omicsNextToken) > 0 {
		input.NextToken = aws.String(_omicsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListReferenceImportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*omics.ListReferenceImportJobsOutput
	p := omics.NewListReferenceImportJobsPaginator(client, input)
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

// Retrieves a list of reference stores linked to your account and returns their
// metadata in JSON format.
//
// For more information, see [Creating a reference store] in the Amazon Web Services HealthOmics User Guide.
//
// [Creating a reference store]: https://docs.aws.amazon.com/omics/latest/dev/create-reference-store.html
func omics_ListReferenceStores(cfg aws.Config, client *omics.Client) {
	input := &omics.ListReferenceStoresInput{}

	if len(_omicsFilter) > 0 {
		if err := assignInputField(input, "Filter", _omicsFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_omicsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _omicsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_omicsNextToken) > 0 {
		input.NextToken = aws.String(_omicsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListReferenceStores(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*omics.ListReferenceStoresOutput
	p := omics.NewListReferenceStoresPaginator(client, input)
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

// Retrieves the metadata of one or more reference genomes in a reference store.
// For more information, see [Creating a reference store] in the Amazon Web Services HealthOmics User Guide.
//
// [Creating a reference store]: https://docs.aws.amazon.com/omics/latest/dev/create-reference-store.html
func omics_ListReferences(cfg aws.Config, client *omics.Client) {
	input := &omics.ListReferencesInput{
		// ReferenceStoreId: *string, // Required
	}

	if len(_omicsReferenceStoreId) > 0 {
		input.ReferenceStoreId = aws.String(_omicsReferenceStoreId)
	}
	if len(_omicsFilter) > 0 {
		if err := assignInputField(input, "Filter", _omicsFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_omicsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _omicsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_omicsNextToken) > 0 {
		input.NextToken = aws.String(_omicsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListReferences(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*omics.ListReferencesOutput
	p := omics.NewListReferencesPaginator(client, input)
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

// Retrieves a list of your run caches and the metadata for each cache.
func omics_ListRunCaches(cfg aws.Config, client *omics.Client) {
	input := &omics.ListRunCachesInput{}

	if len(_omicsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _omicsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_omicsStartingToken) > 0 {
		input.StartingToken = aws.String(_omicsStartingToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRunCaches(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*omics.ListRunCachesOutput
	p := omics.NewListRunCachesPaginator(client, input)
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

// Retrieves a list of all run groups and returns the metadata for each run group.
func omics_ListRunGroups(cfg aws.Config, client *omics.Client) {
	input := &omics.ListRunGroupsInput{}

	if len(_omicsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _omicsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_omicsName) > 0 {
		input.Name = aws.String(_omicsName)
	}
	if len(_omicsStartingToken) > 0 {
		input.StartingToken = aws.String(_omicsStartingToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRunGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*omics.ListRunGroupsOutput
	p := omics.NewListRunGroupsPaginator(client, input)
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

// Returns a list of tasks and status information within their specified run. Use
// this operation to monitor runs and to identify which specific tasks have failed.
func omics_ListRunTasks(cfg aws.Config, client *omics.Client) {
	input := &omics.ListRunTasksInput{
		// Id: *string, // Required
	}

	if len(_omicsId) > 0 {
		input.Id = aws.String(_omicsId)
	}
	if len(_omicsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _omicsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_omicsStartingToken) > 0 {
		input.StartingToken = aws.String(_omicsStartingToken)
	}
	if len(_omicsStatus) > 0 {
		if err := assignInputField(input, "Status", _omicsStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListRunTasks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*omics.ListRunTasksOutput
	p := omics.NewListRunTasksPaginator(client, input)
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

// Retrieves a list of runs and returns each run's metadata and status.
// Amazon Web Services HealthOmics stores a configurable number of runs, as
// determined by service limits, that are available to the console and API. If the
// ListRuns response doesn't include specific runs that you expected, you can find
// all run logs in the CloudWatch logs. For more information about viewing the run
// logs, see [CloudWatch logs]in the Amazon Web Services HealthOmics User Guide.
//
// [CloudWatch logs]: https://docs.aws.amazon.com/omics/latest/dev/monitoring-cloudwatch-logs.html
func omics_ListRuns(cfg aws.Config, client *omics.Client) {
	input := &omics.ListRunsInput{}

	if len(_omicsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _omicsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_omicsName) > 0 {
		input.Name = aws.String(_omicsName)
	}
	if len(_omicsRunGroupId) > 0 {
		input.RunGroupId = aws.String(_omicsRunGroupId)
	}
	if len(_omicsStartingToken) > 0 {
		input.StartingToken = aws.String(_omicsStartingToken)
	}
	if len(_omicsStatus) > 0 {
		if err := assignInputField(input, "Status", _omicsStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListRuns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*omics.ListRunsOutput
	p := omics.NewListRunsPaginator(client, input)
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

// Retrieves a list of sequence stores and returns each sequence store's metadata.
// For more information, see [Creating a HealthOmics sequence store] in the Amazon Web Services HealthOmics User Guide.
//
// [Creating a HealthOmics sequence store]: https://docs.aws.amazon.com/omics/latest/dev/create-sequence-store.html
func omics_ListSequenceStores(cfg aws.Config, client *omics.Client) {
	input := &omics.ListSequenceStoresInput{}

	if len(_omicsFilter) > 0 {
		if err := assignInputField(input, "Filter", _omicsFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_omicsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _omicsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_omicsNextToken) > 0 {
		input.NextToken = aws.String(_omicsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSequenceStores(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*omics.ListSequenceStoresOutput
	p := omics.NewListSequenceStoresPaginator(client, input)
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

// Retrieves the resource shares associated with an account. Use the filter
// parameter to retrieve a specific subset of the shares.
func omics_ListShares(cfg aws.Config, client *omics.Client) {
	input := &omics.ListSharesInput{
		// ResourceOwner: types.ResourceOwner, // Required
	}

	if len(_omicsResourceOwner) > 0 {
		if err := assignInputField(input, "ResourceOwner", _omicsResourceOwner); err != nil {
			log.Errorf("invalid --resource-owner: %s", err.Error())
			return
		}
	}
	if len(_omicsFilter) > 0 {
		if err := assignInputField(input, "Filter", _omicsFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_omicsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _omicsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_omicsNextToken) > 0 {
		input.NextToken = aws.String(_omicsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListShares(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*omics.ListSharesOutput
	p := omics.NewListSharesPaginator(client, input)
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

// Retrieves a list of tags for a resource.
func omics_ListTagsForResource(cfg aws.Config, client *omics.Client) {
	input := &omics.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_omicsResourceArn) > 0 {
		input.ResourceArn = aws.String(_omicsResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Amazon Web Services HealthOmics variant stores and annotation stores will no
// longer be open to new customers starting November 7, 2025. If you would like to
// use variant stores or annotation stores, sign up prior to that date. Existing
// customers can continue to use the service as normal. For more information, see [Amazon Web Services HealthOmics variant store and annotation store availability change].
//
// Retrieves a list of variant import jobs.
//
// [Amazon Web Services HealthOmics variant store and annotation store availability change]: https://docs.aws.amazon.com/omics/latest/dev/variant-store-availability-change.html
func omics_ListVariantImportJobs(cfg aws.Config, client *omics.Client) {
	input := &omics.ListVariantImportJobsInput{}

	if len(_omicsFilter) > 0 {
		if err := assignInputField(input, "Filter", _omicsFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_omicsIds) > 0 {
		input.Ids = append([]string(nil), _omicsIds...)
	}
	if len(_omicsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _omicsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_omicsNextToken) > 0 {
		input.NextToken = aws.String(_omicsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListVariantImportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*omics.ListVariantImportJobsOutput
	p := omics.NewListVariantImportJobsPaginator(client, input)
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

// Amazon Web Services HealthOmics variant stores and annotation stores will no
// longer be open to new customers starting November 7, 2025. If you would like to
// use variant stores or annotation stores, sign up prior to that date. Existing
// customers can continue to use the service as normal. For more information, see [Amazon Web Services HealthOmics variant store and annotation store availability change].
//
// Retrieves a list of variant stores.
//
// [Amazon Web Services HealthOmics variant store and annotation store availability change]: https://docs.aws.amazon.com/omics/latest/dev/variant-store-availability-change.html
func omics_ListVariantStores(cfg aws.Config, client *omics.Client) {
	input := &omics.ListVariantStoresInput{}

	if len(_omicsFilter) > 0 {
		if err := assignInputField(input, "Filter", _omicsFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_omicsIds) > 0 {
		input.Ids = append([]string(nil), _omicsIds...)
	}
	if len(_omicsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _omicsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_omicsNextToken) > 0 {
		input.NextToken = aws.String(_omicsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListVariantStores(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*omics.ListVariantStoresOutput
	p := omics.NewListVariantStoresPaginator(client, input)
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

// Lists the workflow versions for the specified workflow. For more information,
// see [Workflow versioning in Amazon Web Services HealthOmics]in the Amazon Web Services HealthOmics User Guide.
//
// [Workflow versioning in Amazon Web Services HealthOmics]: https://docs.aws.amazon.com/omics/latest/dev/workflow-versions.html
func omics_ListWorkflowVersions(cfg aws.Config, client *omics.Client) {
	input := &omics.ListWorkflowVersionsInput{
		// WorkflowId: *string, // Required
	}

	if len(_omicsWorkflowId) > 0 {
		input.WorkflowId = aws.String(_omicsWorkflowId)
	}
	if len(_omicsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _omicsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_omicsStartingToken) > 0 {
		input.StartingToken = aws.String(_omicsStartingToken)
	}
	if len(_omicsType) > 0 {
		if err := assignInputField(input, "Type", _omicsType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_omicsWorkflowOwnerId) > 0 {
		input.WorkflowOwnerId = aws.String(_omicsWorkflowOwnerId)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkflowVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*omics.ListWorkflowVersionsOutput
	p := omics.NewListWorkflowVersionsPaginator(client, input)
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

// Retrieves a list of existing workflows. You can filter for specific workflows
// by their name and type. Using the type parameter, specify PRIVATE to retrieve a
// list of private workflows or specify READY2RUN for a list of all Ready2Run
// workflows. If you do not specify the type of workflow, this operation returns a
// list of existing workflows.
func omics_ListWorkflows(cfg aws.Config, client *omics.Client) {
	input := &omics.ListWorkflowsInput{}

	if len(_omicsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _omicsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_omicsName) > 0 {
		input.Name = aws.String(_omicsName)
	}
	if len(_omicsStartingToken) > 0 {
		input.StartingToken = aws.String(_omicsStartingToken)
	}
	if len(_omicsType) > 0 {
		if err := assignInputField(input, "Type", _omicsType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
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

	var results []*omics.ListWorkflowsOutput
	p := omics.NewListWorkflowsPaginator(client, input)
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

// Adds an access policy to the specified store.
func omics_PutS3AccessPolicy(cfg aws.Config, client *omics.Client) {
	input := &omics.PutS3AccessPolicyInput{
		// S3AccessPointArn: *string, // Required
		// S3AccessPolicy: *string, // Required
	}

	if len(_omicsS3AccessPointArn) > 0 {
		input.S3AccessPointArn = aws.String(_omicsS3AccessPointArn)
	}
	if len(_omicsS3AccessPolicy) > 0 {
		input.S3AccessPolicy = aws.String(_omicsS3AccessPolicy)
	}

	if resp, err := client.PutS3AccessPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Amazon Web Services HealthOmics variant stores and annotation stores will no
// longer be open to new customers starting November 7, 2025. If you would like to
// use variant stores or annotation stores, sign up prior to that date. Existing
// customers can continue to use the service as normal. For more information, see [Amazon Web Services HealthOmics variant store and annotation store availability change].
//
// Starts an annotation import job.
//
// [Amazon Web Services HealthOmics variant store and annotation store availability change]: https://docs.aws.amazon.com/omics/latest/dev/variant-store-availability-change.html
func omics_StartAnnotationImportJob(cfg aws.Config, client *omics.Client) {
	input := &omics.StartAnnotationImportJobInput{
		// DestinationName: *string, // Required
		// Items: []types.AnnotationImportItemSource, // Required
		// RoleArn: *string, // Required
	}

	if len(_omicsDestinationName) > 0 {
		input.DestinationName = aws.String(_omicsDestinationName)
	}
	if len(_omicsItems) > 0 {
		if err := assignInputField(input, "Items", _omicsItems); err != nil {
			log.Errorf("invalid --items: %s", err.Error())
			return
		}
	}
	if len(_omicsRoleArn) > 0 {
		input.RoleArn = aws.String(_omicsRoleArn)
	}
	if len(_omicsAnnotationFields) > 0 {
		if err := assignInputField(input, "AnnotationFields", _omicsAnnotationFields); err != nil {
			log.Errorf("invalid --annotation-fields: %s", err.Error())
			return
		}
	}
	if len(_omicsFormatOptions) > 0 {
		if err := assignInputField(input, "FormatOptions", _omicsFormatOptions); err != nil {
			log.Errorf("invalid --format-options: %s", err.Error())
			return
		}
	}
	if len(_omicsRunLeftNormalization) > 0 {
		if err := assignInputField(input, "RunLeftNormalization", _omicsRunLeftNormalization); err != nil {
			log.Errorf("invalid --run-left-normalization: %s", err.Error())
			return
		}
	}
	if len(_omicsVersionName) > 0 {
		input.VersionName = aws.String(_omicsVersionName)
	}

	if resp, err := client.StartAnnotationImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Activates an archived read set and returns its metadata in a JSON formatted
// output. AWS HealthOmics automatically archives unused read sets after 30 days.
// To monitor the status of your read set activation job, use the
// GetReadSetActivationJob operation.
//
// To learn more, see [Activating read sets] in the Amazon Web Services HealthOmics User Guide.
//
// [Activating read sets]: https://docs.aws.amazon.com/omics/latest/dev/activating-read-sets.html
func omics_StartReadSetActivationJob(cfg aws.Config, client *omics.Client) {
	input := &omics.StartReadSetActivationJobInput{
		// SequenceStoreId: *string, // Required
		// Sources: []types.StartReadSetActivationJobSourceItem, // Required
	}

	if len(_omicsSequenceStoreId) > 0 {
		input.SequenceStoreId = aws.String(_omicsSequenceStoreId)
	}
	if len(_omicsSources) > 0 {
		if err := assignInputField(input, "Sources", _omicsSources); err != nil {
			log.Errorf("invalid --sources: %s", err.Error())
			return
		}
	}
	if len(_omicsClientToken) > 0 {
		input.ClientToken = aws.String(_omicsClientToken)
	}

	if resp, err := client.StartReadSetActivationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a read set export job. When the export job is finished, the read set is
// exported to an Amazon S3 bucket which can be retrieved using the
// GetReadSetExportJob API operation.
//
// To monitor the status of the export job, use the ListReadSetExportJobs API
// operation.
func omics_StartReadSetExportJob(cfg aws.Config, client *omics.Client) {
	input := &omics.StartReadSetExportJobInput{
		// Destination: *string, // Required
		// RoleArn: *string, // Required
		// SequenceStoreId: *string, // Required
		// Sources: []types.ExportReadSet, // Required
	}

	if len(_omicsDestination) > 0 {
		input.Destination = aws.String(_omicsDestination)
	}
	if len(_omicsRoleArn) > 0 {
		input.RoleArn = aws.String(_omicsRoleArn)
	}
	if len(_omicsSequenceStoreId) > 0 {
		input.SequenceStoreId = aws.String(_omicsSequenceStoreId)
	}
	if len(_omicsSources) > 0 {
		if err := assignInputField(input, "Sources", _omicsSources); err != nil {
			log.Errorf("invalid --sources: %s", err.Error())
			return
		}
	}
	if len(_omicsClientToken) > 0 {
		input.ClientToken = aws.String(_omicsClientToken)
	}

	if resp, err := client.StartReadSetExportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports a read set from the sequence store. Read set import jobs support a
// maximum of 100 read sets of different types. Monitor the progress of your read
// set import job by calling the GetReadSetImportJob API operation.
func omics_StartReadSetImportJob(cfg aws.Config, client *omics.Client) {
	input := &omics.StartReadSetImportJobInput{
		// RoleArn: *string, // Required
		// SequenceStoreId: *string, // Required
		// Sources: []types.StartReadSetImportJobSourceItem, // Required
	}

	if len(_omicsRoleArn) > 0 {
		input.RoleArn = aws.String(_omicsRoleArn)
	}
	if len(_omicsSequenceStoreId) > 0 {
		input.SequenceStoreId = aws.String(_omicsSequenceStoreId)
	}
	if len(_omicsSources) > 0 {
		if err := assignInputField(input, "Sources", _omicsSources); err != nil {
			log.Errorf("invalid --sources: %s", err.Error())
			return
		}
	}
	if len(_omicsClientToken) > 0 {
		input.ClientToken = aws.String(_omicsClientToken)
	}

	if resp, err := client.StartReadSetImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports a reference genome from Amazon S3 into a specified reference store. You
// can have multiple reference genomes in a reference store. You can only import
// reference genomes one at a time into each reference store. Monitor the status of
// your reference import job by using the GetReferenceImportJob API operation.
func omics_StartReferenceImportJob(cfg aws.Config, client *omics.Client) {
	input := &omics.StartReferenceImportJobInput{
		// ReferenceStoreId: *string, // Required
		// RoleArn: *string, // Required
		// Sources: []types.StartReferenceImportJobSourceItem, // Required
	}

	if len(_omicsReferenceStoreId) > 0 {
		input.ReferenceStoreId = aws.String(_omicsReferenceStoreId)
	}
	if len(_omicsRoleArn) > 0 {
		input.RoleArn = aws.String(_omicsRoleArn)
	}
	if len(_omicsSources) > 0 {
		if err := assignInputField(input, "Sources", _omicsSources); err != nil {
			log.Errorf("invalid --sources: %s", err.Error())
			return
		}
	}
	if len(_omicsClientToken) > 0 {
		input.ClientToken = aws.String(_omicsClientToken)
	}

	if resp, err := client.StartReferenceImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a new run and returns details about the run, or duplicates an existing
// run. A run is a single invocation of a workflow. If you provide request IDs,
// Amazon Web Services HealthOmics identifies duplicate requests and starts the run
// only once. Monitor the progress of the run by calling the GetRun API operation.
//
// To start a new run, the following inputs are required:
//
// - A service role ARN ( roleArn ).
//
// - The run's workflow ID ( workflowId , not the uuid or runId ).
//
// - An Amazon S3 location ( outputUri ) where the run outputs will be saved.
//
// - All required workflow parameters ( parameter ), which can include optional
// parameters from the parameter template. The run cannot include any parameters
// that are not defined in the parameter template. To see all possible parameters,
// use the GetRun API operation.
//
// - For runs with a STATIC (default) storage type, specify the required storage
// capacity (in gibibytes). A storage capacity value is not required for runs that
// use DYNAMIC storage.
//
// StartRun can also duplicate an existing run using the run's default values. You
// can modify these default values and/or add other optional inputs. To duplicate a
// run, the following inputs are required:
//
// - A service role ARN ( roleArn ).
//
// - The ID of the run to duplicate ( runId ).
//
// - An Amazon S3 location where the run outputs will be saved ( outputUri ).
//
// To learn more about the optional parameters for StartRun , see [Starting a run] in the Amazon
// Web Services HealthOmics User Guide.
//
// Use the retentionMode input to control how long the metadata for each run is
// stored in CloudWatch. There are two retention modes:
//
// - Specify REMOVE to automatically remove the oldest runs when you reach the
// maximum service retention limit for runs. It is recommended that you use the
// REMOVE mode to initiate major run requests so that your runs do not fail when
// you reach the limit.
//
// - The retentionMode is set to the RETAIN mode by default, which allows you to
// manually remove runs after reaching the maximum service retention limit. Under
// this setting, you cannot create additional runs until you remove the excess
// runs.
//
// To learn more about the retention modes, see [Run retention mode] in the Amazon Web Services
// HealthOmics User Guide.
//
// You can use Amazon Q CLI to analyze run logs and make performance optimization
// recommendations. To get started, see the [Amazon Web Services HealthOmics MCP server]on GitHub.
//
// [Starting a run]: https://docs.aws.amazon.com/omics/latest/dev/starting-a-run.html
// [Amazon Web Services HealthOmics MCP server]: https://github.com/awslabs/mcp/tree/main/src/aws-healthomics-mcp-server
// [Run retention mode]: https://docs.aws.amazon.com/omics/latest/dev/run-retention.html
func omics_StartRun(cfg aws.Config, client *omics.Client) {
	input := &omics.StartRunInput{
		// OutputUri: *string, // Required
		// RequestId: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_omicsOutputUri) > 0 {
		input.OutputUri = aws.String(_omicsOutputUri)
	}
	if len(_omicsRequestId) > 0 {
		input.RequestId = aws.String(_omicsRequestId)
	}
	if len(_omicsRoleArn) > 0 {
		input.RoleArn = aws.String(_omicsRoleArn)
	}
	if len(_omicsCacheBehavior) > 0 {
		if err := assignInputField(input, "CacheBehavior", _omicsCacheBehavior); err != nil {
			log.Errorf("invalid --cache-behavior: %s", err.Error())
			return
		}
	}
	if len(_omicsCacheId) > 0 {
		input.CacheId = aws.String(_omicsCacheId)
	}
	if len(_omicsLogLevel) > 0 {
		if err := assignInputField(input, "LogLevel", _omicsLogLevel); err != nil {
			log.Errorf("invalid --log-level: %s", err.Error())
			return
		}
	}
	if len(_omicsName) > 0 {
		input.Name = aws.String(_omicsName)
	}
	if len(_omicsParameters) > 0 {
		if err := assignInputField(input, "Parameters", _omicsParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_omicsPriority) > 0 {
		if err := assignInputField(input, "Priority", _omicsPriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}
	if len(_omicsRetentionMode) > 0 {
		if err := assignInputField(input, "RetentionMode", _omicsRetentionMode); err != nil {
			log.Errorf("invalid --retention-mode: %s", err.Error())
			return
		}
	}
	if len(_omicsRunGroupId) > 0 {
		input.RunGroupId = aws.String(_omicsRunGroupId)
	}
	if len(_omicsRunId) > 0 {
		input.RunId = aws.String(_omicsRunId)
	}
	if len(_omicsStorageCapacity) > 0 {
		if err := assignInputField(input, "StorageCapacity", _omicsStorageCapacity); err != nil {
			log.Errorf("invalid --storage-capacity: %s", err.Error())
			return
		}
	}
	if len(_omicsStorageType) > 0 {
		if err := assignInputField(input, "StorageType", _omicsStorageType); err != nil {
			log.Errorf("invalid --storage-type: %s", err.Error())
			return
		}
	}
	if len(_omicsTags) > 0 {
		if err := assignInputField(input, "Tags", _omicsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_omicsWorkflowId) > 0 {
		input.WorkflowId = aws.String(_omicsWorkflowId)
	}
	if len(_omicsWorkflowOwnerId) > 0 {
		input.WorkflowOwnerId = aws.String(_omicsWorkflowOwnerId)
	}
	if len(_omicsWorkflowType) > 0 {
		if err := assignInputField(input, "WorkflowType", _omicsWorkflowType); err != nil {
			log.Errorf("invalid --workflow-type: %s", err.Error())
			return
		}
	}
	if len(_omicsWorkflowVersionName) > 0 {
		input.WorkflowVersionName = aws.String(_omicsWorkflowVersionName)
	}

	if resp, err := client.StartRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Amazon Web Services HealthOmics variant stores and annotation stores will no
// longer be open to new customers starting November 7, 2025. If you would like to
// use variant stores or annotation stores, sign up prior to that date. Existing
// customers can continue to use the service as normal. For more information, see [Amazon Web Services HealthOmics variant store and annotation store availability change].
//
// Starts a variant import job.
//
// [Amazon Web Services HealthOmics variant store and annotation store availability change]: https://docs.aws.amazon.com/omics/latest/dev/variant-store-availability-change.html
func omics_StartVariantImportJob(cfg aws.Config, client *omics.Client) {
	input := &omics.StartVariantImportJobInput{
		// DestinationName: *string, // Required
		// Items: []types.VariantImportItemSource, // Required
		// RoleArn: *string, // Required
	}

	if len(_omicsDestinationName) > 0 {
		input.DestinationName = aws.String(_omicsDestinationName)
	}
	if len(_omicsItems) > 0 {
		if err := assignInputField(input, "Items", _omicsItems); err != nil {
			log.Errorf("invalid --items: %s", err.Error())
			return
		}
	}
	if len(_omicsRoleArn) > 0 {
		input.RoleArn = aws.String(_omicsRoleArn)
	}
	if len(_omicsAnnotationFields) > 0 {
		if err := assignInputField(input, "AnnotationFields", _omicsAnnotationFields); err != nil {
			log.Errorf("invalid --annotation-fields: %s", err.Error())
			return
		}
	}
	if len(_omicsRunLeftNormalization) > 0 {
		if err := assignInputField(input, "RunLeftNormalization", _omicsRunLeftNormalization); err != nil {
			log.Errorf("invalid --run-left-normalization: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartVariantImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tags a resource.
func omics_TagResource(cfg aws.Config, client *omics.Client) {
	input := &omics.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_omicsResourceArn) > 0 {
		input.ResourceArn = aws.String(_omicsResourceArn)
	}
	if len(_omicsTags) > 0 {
		if err := assignInputField(input, "Tags", _omicsTags); err != nil {
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

// Removes tags from a resource.
func omics_UntagResource(cfg aws.Config, client *omics.Client) {
	input := &omics.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_omicsResourceArn) > 0 {
		input.ResourceArn = aws.String(_omicsResourceArn)
	}
	if len(_omicsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _omicsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Amazon Web Services HealthOmics variant stores and annotation stores will no
// longer be open to new customers starting November 7, 2025. If you would like to
// use variant stores or annotation stores, sign up prior to that date. Existing
// customers can continue to use the service as normal. For more information, see [Amazon Web Services HealthOmics variant store and annotation store availability change].
//
// Updates an annotation store.
//
// [Amazon Web Services HealthOmics variant store and annotation store availability change]: https://docs.aws.amazon.com/omics/latest/dev/variant-store-availability-change.html
func omics_UpdateAnnotationStore(cfg aws.Config, client *omics.Client) {
	input := &omics.UpdateAnnotationStoreInput{
		// Name: *string, // Required
	}

	if len(_omicsName) > 0 {
		input.Name = aws.String(_omicsName)
	}
	if len(_omicsDescription) > 0 {
		input.Description = aws.String(_omicsDescription)
	}

	if resp, err := client.UpdateAnnotationStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the description of an annotation store version.
func omics_UpdateAnnotationStoreVersion(cfg aws.Config, client *omics.Client) {
	input := &omics.UpdateAnnotationStoreVersionInput{
		// Name: *string, // Required
		// VersionName: *string, // Required
	}

	if len(_omicsName) > 0 {
		input.Name = aws.String(_omicsName)
	}
	if len(_omicsVersionName) > 0 {
		input.VersionName = aws.String(_omicsVersionName)
	}
	if len(_omicsDescription) > 0 {
		input.Description = aws.String(_omicsDescription)
	}

	if resp, err := client.UpdateAnnotationStoreVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a run cache using its ID and returns a response with no body if the
// operation is successful. You can update the run cache description, name, or the
// default run cache behavior with CACHE_ON_FAILURE or CACHE_ALWAYS . To confirm
// that your run cache settings have been properly updated, use the GetRunCache
// API operation.
//
// For more information, see [How call caching works] in the Amazon Web Services HealthOmics User Guide.
//
// [How call caching works]: https://docs.aws.amazon.com/omics/latest/dev/how-run-cache.html
func omics_UpdateRunCache(cfg aws.Config, client *omics.Client) {
	input := &omics.UpdateRunCacheInput{
		// Id: *string, // Required
	}

	if len(_omicsId) > 0 {
		input.Id = aws.String(_omicsId)
	}
	if len(_omicsCacheBehavior) > 0 {
		if err := assignInputField(input, "CacheBehavior", _omicsCacheBehavior); err != nil {
			log.Errorf("invalid --cache-behavior: %s", err.Error())
			return
		}
	}
	if len(_omicsDescription) > 0 {
		input.Description = aws.String(_omicsDescription)
	}
	if len(_omicsName) > 0 {
		input.Name = aws.String(_omicsName)
	}

	if resp, err := client.UpdateRunCache(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the settings of a run group and returns a response with no body if the
// operation is successful.
//
// You can update the following settings with UpdateRunGroup :
//
// - Maximum number of CPUs
//
// - Run time (measured in minutes)
//
// - Number of GPUs
//
// - Number of concurrent runs
//
// - Group name
//
// To confirm that the settings have been successfully updated, use the
// ListRunGroups or GetRunGroup API operations to verify that the desired changes
// have been made.
func omics_UpdateRunGroup(cfg aws.Config, client *omics.Client) {
	input := &omics.UpdateRunGroupInput{
		// Id: *string, // Required
	}

	if len(_omicsId) > 0 {
		input.Id = aws.String(_omicsId)
	}
	if len(_omicsMaxCpus) > 0 {
		if err := assignInputField(input, "MaxCpus", _omicsMaxCpus); err != nil {
			log.Errorf("invalid --max-cpus: %s", err.Error())
			return
		}
	}
	if len(_omicsMaxDuration) > 0 {
		if err := assignInputField(input, "MaxDuration", _omicsMaxDuration); err != nil {
			log.Errorf("invalid --max-duration: %s", err.Error())
			return
		}
	}
	if len(_omicsMaxGpus) > 0 {
		if err := assignInputField(input, "MaxGpus", _omicsMaxGpus); err != nil {
			log.Errorf("invalid --max-gpus: %s", err.Error())
			return
		}
	}
	if len(_omicsMaxRuns) > 0 {
		if err := assignInputField(input, "MaxRuns", _omicsMaxRuns); err != nil {
			log.Errorf("invalid --max-runs: %s", err.Error())
			return
		}
	}
	if len(_omicsName) > 0 {
		input.Name = aws.String(_omicsName)
	}

	if resp, err := client.UpdateRunGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update one or more parameters for the sequence store.
func omics_UpdateSequenceStore(cfg aws.Config, client *omics.Client) {
	input := &omics.UpdateSequenceStoreInput{
		// Id: *string, // Required
	}

	if len(_omicsId) > 0 {
		input.Id = aws.String(_omicsId)
	}
	if len(_omicsClientToken) > 0 {
		input.ClientToken = aws.String(_omicsClientToken)
	}
	if len(_omicsDescription) > 0 {
		input.Description = aws.String(_omicsDescription)
	}
	if len(_omicsFallbackLocation) > 0 {
		input.FallbackLocation = aws.String(_omicsFallbackLocation)
	}
	if len(_omicsName) > 0 {
		input.Name = aws.String(_omicsName)
	}
	if len(_omicsPropagatedSetLevelTags) > 0 {
		input.PropagatedSetLevelTags = append([]string(nil), _omicsPropagatedSetLevelTags...)
	}
	if len(_omicsS3AccessConfig) > 0 {
		if err := assignInputField(input, "S3AccessConfig", _omicsS3AccessConfig); err != nil {
			log.Errorf("invalid --s3-access-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSequenceStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Amazon Web Services HealthOmics variant stores and annotation stores will no
// longer be open to new customers starting November 7, 2025. If you would like to
// use variant stores or annotation stores, sign up prior to that date. Existing
// customers can continue to use the service as normal. For more information, see [Amazon Web Services HealthOmics variant store and annotation store availability change].
//
// Updates a variant store.
//
// [Amazon Web Services HealthOmics variant store and annotation store availability change]: https://docs.aws.amazon.com/omics/latest/dev/variant-store-availability-change.html
func omics_UpdateVariantStore(cfg aws.Config, client *omics.Client) {
	input := &omics.UpdateVariantStoreInput{
		// Name: *string, // Required
	}

	if len(_omicsName) > 0 {
		input.Name = aws.String(_omicsName)
	}
	if len(_omicsDescription) > 0 {
		input.Description = aws.String(_omicsDescription)
	}

	if resp, err := client.UpdateVariantStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates information about a workflow.
// You can update the following workflow information:
//
// - Name
//
// - Description
//
// - Default storage type
//
// - Default storage capacity (with workflow ID)
//
// This operation returns a response with no body if the operation is successful.
// You can check the workflow updates by calling the GetWorkflow API operation.
//
// For more information, see [Update a private workflow] in the Amazon Web Services HealthOmics User Guide.
//
// [Update a private workflow]: https://docs.aws.amazon.com/omics/latest/dev/update-private-workflow.html
func omics_UpdateWorkflow(cfg aws.Config, client *omics.Client) {
	input := &omics.UpdateWorkflowInput{
		// Id: *string, // Required
	}

	if len(_omicsId) > 0 {
		input.Id = aws.String(_omicsId)
	}
	if len(_omicsDescription) > 0 {
		input.Description = aws.String(_omicsDescription)
	}
	if len(_omicsName) > 0 {
		input.Name = aws.String(_omicsName)
	}
	if len(_omicsReadmeMarkdown) > 0 {
		input.ReadmeMarkdown = aws.String(_omicsReadmeMarkdown)
	}
	if len(_omicsStorageCapacity) > 0 {
		if err := assignInputField(input, "StorageCapacity", _omicsStorageCapacity); err != nil {
			log.Errorf("invalid --storage-capacity: %s", err.Error())
			return
		}
	}
	if len(_omicsStorageType) > 0 {
		if err := assignInputField(input, "StorageType", _omicsStorageType); err != nil {
			log.Errorf("invalid --storage-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates information about the workflow version. For more information, see [Workflow versioning in Amazon Web Services HealthOmics] in
// the Amazon Web Services HealthOmics User Guide.
//
// [Workflow versioning in Amazon Web Services HealthOmics]: https://docs.aws.amazon.com/omics/latest/dev/workflow-versions.html
func omics_UpdateWorkflowVersion(cfg aws.Config, client *omics.Client) {
	input := &omics.UpdateWorkflowVersionInput{
		// VersionName: *string, // Required
		// WorkflowId: *string, // Required
	}

	if len(_omicsVersionName) > 0 {
		input.VersionName = aws.String(_omicsVersionName)
	}
	if len(_omicsWorkflowId) > 0 {
		input.WorkflowId = aws.String(_omicsWorkflowId)
	}
	if len(_omicsDescription) > 0 {
		input.Description = aws.String(_omicsDescription)
	}
	if len(_omicsReadmeMarkdown) > 0 {
		input.ReadmeMarkdown = aws.String(_omicsReadmeMarkdown)
	}
	if len(_omicsStorageCapacity) > 0 {
		if err := assignInputField(input, "StorageCapacity", _omicsStorageCapacity); err != nil {
			log.Errorf("invalid --storage-capacity: %s", err.Error())
			return
		}
	}
	if len(_omicsStorageType) > 0 {
		if err := assignInputField(input, "StorageType", _omicsStorageType); err != nil {
			log.Errorf("invalid --storage-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateWorkflowVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Uploads a specific part of a read set into a sequence store. When you a upload
// a read set part with a part number that already exists, the new part replaces
// the existing one. This operation returns a JSON formatted response containing a
// string identifier that is used to confirm that parts are being added to the
// intended upload.
//
// For more information, see [Direct upload to a sequence store] in the Amazon Web Services HealthOmics User Guide.
//
// [Direct upload to a sequence store]: https://docs.aws.amazon.com/omics/latest/dev/synchronous-uploads.html
func omics_UploadReadSetPart(cfg aws.Config, client *omics.Client) {
	input := &omics.UploadReadSetPartInput{
		// PartNumber: *int32, // Required
		// PartSource: types.ReadSetPartSource, // Required
		// Payload: io.Reader, // Required
		// SequenceStoreId: *string, // Required
		// UploadId: *string, // Required
	}

	if len(_omicsPartNumber) > 0 {
		if err := assignInputField(input, "PartNumber", _omicsPartNumber); err != nil {
			log.Errorf("invalid --part-number: %s", err.Error())
			return
		}
	}
	if len(_omicsPartSource) > 0 {
		if err := assignInputField(input, "PartSource", _omicsPartSource); err != nil {
			log.Errorf("invalid --part-source: %s", err.Error())
			return
		}
	}
	if len(_omicsPayload) > 0 {
		if err := assignInputField(input, "Payload", _omicsPayload); err != nil {
			log.Errorf("invalid --payload: %s", err.Error())
			return
		}
	}
	if len(_omicsSequenceStoreId) > 0 {
		input.SequenceStoreId = aws.String(_omicsSequenceStoreId)
	}
	if len(_omicsUploadId) > 0 {
		input.UploadId = aws.String(_omicsUploadId)
	}

	if resp, err := client.UploadReadSetPart(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_omicsCmd)
	_omicsCmd.Flags().SortFlags = false

	_omicsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_omicsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_omicsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_omicsCmd.Flags().StringVarP(&_omicsAccelerators, "accelerators", "", "", "Accelerators")
	_omicsCmd.Flags().StringVarP(&_omicsAnnotationFields, "annotation-fields", "", "", "Annotation Fields")
	_omicsCmd.Flags().StringVarP(&_omicsCacheBehavior, "cache-behavior", "", "", "Cache Behavior")
	_omicsCmd.Flags().StringVarP(&_omicsCacheBucketOwnerId, "cache-bucket-owner-id", "", "", "Cache Bucket Owner ID")
	_omicsCmd.Flags().StringVarP(&_omicsCacheId, "cache-id", "", "", "Cache ID")
	_omicsCmd.Flags().StringVarP(&_omicsCacheS3Location, "cache-s3-location", "", "", "Cache S3 Location")
	_omicsCmd.Flags().StringVarP(&_omicsClientToken, "client-token", "", "", "Client Token")
	_omicsCmd.Flags().StringVarP(&_omicsContainerRegistryMap, "container-registry-map", "", "", "Container Registry Map")
	_omicsCmd.Flags().StringVarP(&_omicsContainerRegistryMapUri, "container-registry-map-uri", "", "", "Container Registry Map URI")
	_omicsCmd.Flags().StringVarP(&_omicsDefinitionRepository, "definition-repository", "", "", "Definition Repository")
	_omicsCmd.Flags().StringVarP(&_omicsDefinitionUri, "definition-uri", "", "", "Definition URI")
	_omicsCmd.Flags().StringVarP(&_omicsDefinitionZip, "definition-zip", "", "", "Definition Zip")
	_omicsCmd.Flags().StringVarP(&_omicsDescription, "description", "", "", "Description")
	_omicsCmd.Flags().StringVarP(&_omicsDestination, "destination", "", "", "Destination")
	_omicsCmd.Flags().StringVarP(&_omicsDestinationName, "destination-name", "", "", "Destination Name")
	_omicsCmd.Flags().StringVarP(&_omicsEngine, "engine", "", "", "Engine")
	_omicsCmd.Flags().StringVarP(&_omicsETagAlgorithmFamily, "etag-algorithm-family", "", "", "Etag Algorithm Family")
	_omicsCmd.Flags().StringVarP(&_omicsExport, "export", "", "", "Export")
	_omicsCmd.Flags().StringVarP(&_omicsFallbackLocation, "fallback-location", "", "", "Fallback Location")
	_omicsCmd.Flags().StringVarP(&_omicsFile, "file", "", "", "File")
	_omicsCmd.Flags().StringVarP(&_omicsFilter, "filter", "", "", "Filter")
	_omicsCmd.Flags().StringVarP(&_omicsForce, "force", "", "", "Force")
	_omicsCmd.Flags().StringVarP(&_omicsFormatOptions, "format-options", "", "", "Format Options")
	_omicsCmd.Flags().StringVarP(&_omicsGeneratedFrom, "generated-from", "", "", "Generated From")
	_omicsCmd.Flags().StringVarP(&_omicsId, "id", "", "", "ID")
	_omicsCmd.Flags().StringSliceVarP(&_omicsIds, "ids", "", nil, "Ids")
	_omicsCmd.Flags().StringVarP(&_omicsItems, "items", "", "", "Items")
	_omicsCmd.Flags().StringVarP(&_omicsJobId, "job-id", "", "", "Job ID")
	_omicsCmd.Flags().StringVarP(&_omicsLogLevel, "log-level", "", "", "Log Level")
	_omicsCmd.Flags().StringVarP(&_omicsMain, "main", "", "", "Main")
	_omicsCmd.Flags().StringVarP(&_omicsMaxCpus, "max-cpus", "", "", "Max Cpus")
	_omicsCmd.Flags().StringVarP(&_omicsMaxDuration, "max-duration", "", "", "Max Duration")
	_omicsCmd.Flags().StringVarP(&_omicsMaxGpus, "max-gpus", "", "", "Max Gpus")
	_omicsCmd.Flags().StringVarP(&_omicsMaxResults, "max-results", "", "", "Max Results")
	_omicsCmd.Flags().StringVarP(&_omicsMaxRuns, "max-runs", "", "", "Max Runs")
	_omicsCmd.Flags().StringVarP(&_omicsName, "name", "", "", "Name")
	_omicsCmd.Flags().StringVarP(&_omicsNextToken, "next-token", "", "", "Next Token")
	_omicsCmd.Flags().StringVarP(&_omicsOutputUri, "output-uri", "", "", "Output URI")
	_omicsCmd.Flags().StringVarP(&_omicsParameterTemplate, "parameter-template", "", "", "Parameter Template")
	_omicsCmd.Flags().StringVarP(&_omicsParameterTemplatePath, "parameter-template-path", "", "", "Parameter Template Path")
	_omicsCmd.Flags().StringVarP(&_omicsParameters, "parameters", "", "", "Parameters")
	_omicsCmd.Flags().StringVarP(&_omicsPartNumber, "part-number", "", "", "Part Number")
	_omicsCmd.Flags().StringVarP(&_omicsPartSource, "part-source", "", "", "Part Source")
	_omicsCmd.Flags().StringVarP(&_omicsParts, "parts", "", "", "Parts")
	_omicsCmd.Flags().StringVarP(&_omicsPayload, "payload", "", "", "Payload")
	_omicsCmd.Flags().StringVarP(&_omicsPrincipalSubscriber, "principal-subscriber", "", "", "Principal Subscriber")
	_omicsCmd.Flags().StringVarP(&_omicsPriority, "priority", "", "", "Priority")
	_omicsCmd.Flags().StringSliceVarP(&_omicsPropagatedSetLevelTags, "propagated-set-level-tags", "", nil, "Propagated Set Level Tags")
	_omicsCmd.Flags().StringVarP(&_omicsRange, "range", "", "", "Range")
	_omicsCmd.Flags().StringVarP(&_omicsReadmeMarkdown, "readme-markdown", "", "", "Readme Markdown")
	_omicsCmd.Flags().StringVarP(&_omicsReadmePath, "readme-path", "", "", "Readme Path")
	_omicsCmd.Flags().StringVarP(&_omicsReadmeUri, "readme-uri", "", "", "Readme URI")
	_omicsCmd.Flags().StringVarP(&_omicsReference, "reference", "", "", "Reference")
	_omicsCmd.Flags().StringVarP(&_omicsReferenceArn, "reference-arn", "", "", "Reference ARN")
	_omicsCmd.Flags().StringVarP(&_omicsReferenceStoreId, "reference-store-id", "", "", "Reference Store ID")
	_omicsCmd.Flags().StringVarP(&_omicsRequestId, "request-id", "", "", "Request ID")
	_omicsCmd.Flags().StringVarP(&_omicsResourceArn, "resource-arn", "", "", "Resource ARN")
	_omicsCmd.Flags().StringVarP(&_omicsResourceOwner, "resource-owner", "", "", "Resource Owner")
	_omicsCmd.Flags().StringVarP(&_omicsRetentionMode, "retention-mode", "", "", "Retention Mode")
	_omicsCmd.Flags().StringVarP(&_omicsRoleArn, "role-arn", "", "", "Role ARN")
	_omicsCmd.Flags().StringVarP(&_omicsRunGroupId, "run-group-id", "", "", "Run Group ID")
	_omicsCmd.Flags().StringVarP(&_omicsRunId, "run-id", "", "", "Run ID")
	_omicsCmd.Flags().StringVarP(&_omicsRunLeftNormalization, "run-left-normalization", "", "", "Run Left Normalization")
	_omicsCmd.Flags().StringVarP(&_omicsS3AccessConfig, "s3-access-config", "", "", "S3 Access Config")
	_omicsCmd.Flags().StringVarP(&_omicsS3AccessPointArn, "s3-access-point-arn", "", "", "S3 Access Point ARN")
	_omicsCmd.Flags().StringVarP(&_omicsS3AccessPolicy, "s3-access-policy", "", "", "S3 Access Policy")
	_omicsCmd.Flags().StringVarP(&_omicsSampleId, "sample-id", "", "", "Sample ID")
	_omicsCmd.Flags().StringVarP(&_omicsSequenceStoreId, "sequence-store-id", "", "", "Sequence Store ID")
	_omicsCmd.Flags().StringVarP(&_omicsShareId, "share-id", "", "", "Share ID")
	_omicsCmd.Flags().StringVarP(&_omicsShareName, "share-name", "", "", "Share Name")
	_omicsCmd.Flags().StringVarP(&_omicsSourceFileType, "source-file-type", "", "", "Source File Type")
	_omicsCmd.Flags().StringVarP(&_omicsSources, "sources", "", "", "Sources")
	_omicsCmd.Flags().StringVarP(&_omicsSseConfig, "sse-config", "", "", "SSE Config")
	_omicsCmd.Flags().StringVarP(&_omicsStartingToken, "starting-token", "", "", "Starting Token")
	_omicsCmd.Flags().StringVarP(&_omicsStatus, "status", "", "", "Status")
	_omicsCmd.Flags().StringVarP(&_omicsStorageCapacity, "storage-capacity", "", "", "Storage Capacity")
	_omicsCmd.Flags().StringVarP(&_omicsStorageType, "storage-type", "", "", "Storage Type")
	_omicsCmd.Flags().StringVarP(&_omicsStoreFormat, "store-format", "", "", "Store Format")
	_omicsCmd.Flags().StringVarP(&_omicsStoreOptions, "store-options", "", "", "Store Options")
	_omicsCmd.Flags().StringVarP(&_omicsSubjectId, "subject-id", "", "", "Subject ID")
	_omicsCmd.Flags().StringSliceVarP(&_omicsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_omicsCmd.Flags().StringVarP(&_omicsTags, "tags", "", "", "Tags")
	_omicsCmd.Flags().StringVarP(&_omicsTaskId, "task-id", "", "", "Task ID")
	_omicsCmd.Flags().StringVarP(&_omicsType, "type", "", "", "Type")
	_omicsCmd.Flags().StringVarP(&_omicsUploadId, "upload-id", "", "", "Upload ID")
	_omicsCmd.Flags().StringVarP(&_omicsVersionName, "version-name", "", "", "Version Name")
	_omicsCmd.Flags().StringVarP(&_omicsVersionOptions, "version-options", "", "", "Version Options")
	_omicsCmd.Flags().StringSliceVarP(&_omicsVersions, "versions", "", nil, "Versions")
	_omicsCmd.Flags().StringVarP(&_omicsWorkflowBucketOwnerId, "workflow-bucket-owner-id", "", "", "Workflow Bucket Owner ID")
	_omicsCmd.Flags().StringVarP(&_omicsWorkflowId, "workflow-id", "", "", "Workflow ID")
	_omicsCmd.Flags().StringVarP(&_omicsWorkflowOwnerId, "workflow-owner-id", "", "", "Workflow Owner ID")
	_omicsCmd.Flags().StringVarP(&_omicsWorkflowType, "workflow-type", "", "", "Workflow Type")
	_omicsCmd.Flags().StringVarP(&_omicsWorkflowVersionName, "workflow-version-name", "", "", "Workflow Version Name")

	_omicsCmd.Flags().BoolVarP(&_omicsAbortMultipartReadSetUpload, "abort-multipart-read-set-upload", "", false, "Abort Multipart Read Set Upload")
	_omicsCmd.Flags().BoolVarP(&_omicsAcceptShare, "accept-share", "", false, "Accept Share")
	_omicsCmd.Flags().BoolVarP(&_omicsBatchDeleteReadSet, "batch-delete-read-set", "", false, "Batch Delete Read Set")
	_omicsCmd.Flags().BoolVarP(&_omicsCancelAnnotationImportJob, "cancel-annotation-import-job", "", false, "Cancel Annotation Import Job")
	_omicsCmd.Flags().BoolVarP(&_omicsCancelRun, "cancel-run", "", false, "Cancel Run")
	_omicsCmd.Flags().BoolVarP(&_omicsCancelVariantImportJob, "cancel-variant-import-job", "", false, "Cancel Variant Import Job")
	_omicsCmd.Flags().BoolVarP(&_omicsCompleteMultipartReadSetUpload, "complete-multipart-read-set-upload", "", false, "Complete Multipart Read Set Upload")
	_omicsCmd.Flags().BoolVarP(&_omicsCreateAnnotationStore, "create-annotation-store", "", false, "Create Annotation Store")
	_omicsCmd.Flags().BoolVarP(&_omicsCreateAnnotationStoreVersion, "create-annotation-store-version", "", false, "Create Annotation Store Version")
	_omicsCmd.Flags().BoolVarP(&_omicsCreateMultipartReadSetUpload, "create-multipart-read-set-upload", "", false, "Create Multipart Read Set Upload")
	_omicsCmd.Flags().BoolVarP(&_omicsCreateReferenceStore, "create-reference-store", "", false, "Create Reference Store")
	_omicsCmd.Flags().BoolVarP(&_omicsCreateRunCache, "create-run-cache", "", false, "Create Run Cache")
	_omicsCmd.Flags().BoolVarP(&_omicsCreateRunGroup, "create-run-group", "", false, "Create Run Group")
	_omicsCmd.Flags().BoolVarP(&_omicsCreateSequenceStore, "create-sequence-store", "", false, "Create Sequence Store")
	_omicsCmd.Flags().BoolVarP(&_omicsCreateShare, "create-share", "", false, "Create Share")
	_omicsCmd.Flags().BoolVarP(&_omicsCreateVariantStore, "create-variant-store", "", false, "Create Variant Store")
	_omicsCmd.Flags().BoolVarP(&_omicsCreateWorkflow, "create-workflow", "", false, "Create Workflow")
	_omicsCmd.Flags().BoolVarP(&_omicsCreateWorkflowVersion, "create-workflow-version", "", false, "Create Workflow Version")
	_omicsCmd.Flags().BoolVarP(&_omicsDeleteAnnotationStore, "delete-annotation-store", "", false, "Delete Annotation Store")
	_omicsCmd.Flags().BoolVarP(&_omicsDeleteAnnotationStoreVersions, "delete-annotation-store-versions", "", false, "Delete Annotation Store Versions")
	_omicsCmd.Flags().BoolVarP(&_omicsDeleteReference, "delete-reference", "", false, "Delete Reference")
	_omicsCmd.Flags().BoolVarP(&_omicsDeleteReferenceStore, "delete-reference-store", "", false, "Delete Reference Store")
	_omicsCmd.Flags().BoolVarP(&_omicsDeleteRun, "delete-run", "", false, "Delete Run")
	_omicsCmd.Flags().BoolVarP(&_omicsDeleteRunCache, "delete-run-cache", "", false, "Delete Run Cache")
	_omicsCmd.Flags().BoolVarP(&_omicsDeleteRunGroup, "delete-run-group", "", false, "Delete Run Group")
	_omicsCmd.Flags().BoolVarP(&_omicsDeleteS3AccessPolicy, "delete-s3-access-policy", "", false, "Delete S3 Access Policy")
	_omicsCmd.Flags().BoolVarP(&_omicsDeleteSequenceStore, "delete-sequence-store", "", false, "Delete Sequence Store")
	_omicsCmd.Flags().BoolVarP(&_omicsDeleteShare, "delete-share", "", false, "Delete Share")
	_omicsCmd.Flags().BoolVarP(&_omicsDeleteVariantStore, "delete-variant-store", "", false, "Delete Variant Store")
	_omicsCmd.Flags().BoolVarP(&_omicsDeleteWorkflow, "delete-workflow", "", false, "Delete Workflow")
	_omicsCmd.Flags().BoolVarP(&_omicsDeleteWorkflowVersion, "delete-workflow-version", "", false, "Delete Workflow Version")
	_omicsCmd.Flags().BoolVarP(&_omicsGetAnnotationImportJob, "get-annotation-import-job", "", false, "Get Annotation Import Job")
	_omicsCmd.Flags().BoolVarP(&_omicsGetAnnotationStore, "get-annotation-store", "", false, "Get Annotation Store")
	_omicsCmd.Flags().BoolVarP(&_omicsGetAnnotationStoreVersion, "get-annotation-store-version", "", false, "Get Annotation Store Version")
	_omicsCmd.Flags().BoolVarP(&_omicsGetReadSet, "get-read-set", "", false, "Get Read Set")
	_omicsCmd.Flags().BoolVarP(&_omicsGetReadSetActivationJob, "get-read-set-activation-job", "", false, "Get Read Set Activation Job")
	_omicsCmd.Flags().BoolVarP(&_omicsGetReadSetExportJob, "get-read-set-export-job", "", false, "Get Read Set Export Job")
	_omicsCmd.Flags().BoolVarP(&_omicsGetReadSetImportJob, "get-read-set-import-job", "", false, "Get Read Set Import Job")
	_omicsCmd.Flags().BoolVarP(&_omicsGetReadSetMetadata, "get-read-set-metadata", "", false, "Get Read Set Metadata")
	_omicsCmd.Flags().BoolVarP(&_omicsGetReference, "get-reference", "", false, "Get Reference")
	_omicsCmd.Flags().BoolVarP(&_omicsGetReferenceImportJob, "get-reference-import-job", "", false, "Get Reference Import Job")
	_omicsCmd.Flags().BoolVarP(&_omicsGetReferenceMetadata, "get-reference-metadata", "", false, "Get Reference Metadata")
	_omicsCmd.Flags().BoolVarP(&_omicsGetReferenceStore, "get-reference-store", "", false, "Get Reference Store")
	_omicsCmd.Flags().BoolVarP(&_omicsGetRun, "get-run", "", false, "Get Run")
	_omicsCmd.Flags().BoolVarP(&_omicsGetRunCache, "get-run-cache", "", false, "Get Run Cache")
	_omicsCmd.Flags().BoolVarP(&_omicsGetRunGroup, "get-run-group", "", false, "Get Run Group")
	_omicsCmd.Flags().BoolVarP(&_omicsGetRunTask, "get-run-task", "", false, "Get Run Task")
	_omicsCmd.Flags().BoolVarP(&_omicsGetS3AccessPolicy, "get-s3-access-policy", "", false, "Get S3 Access Policy")
	_omicsCmd.Flags().BoolVarP(&_omicsGetSequenceStore, "get-sequence-store", "", false, "Get Sequence Store")
	_omicsCmd.Flags().BoolVarP(&_omicsGetShare, "get-share", "", false, "Get Share")
	_omicsCmd.Flags().BoolVarP(&_omicsGetVariantImportJob, "get-variant-import-job", "", false, "Get Variant Import Job")
	_omicsCmd.Flags().BoolVarP(&_omicsGetVariantStore, "get-variant-store", "", false, "Get Variant Store")
	_omicsCmd.Flags().BoolVarP(&_omicsGetWorkflow, "get-workflow", "", false, "Get Workflow")
	_omicsCmd.Flags().BoolVarP(&_omicsGetWorkflowVersion, "get-workflow-version", "", false, "Get Workflow Version")
	_omicsCmd.Flags().BoolVarP(&_omicsListAnnotationImportJobs, "list-annotation-import-jobs", "", false, "List Annotation Import Jobs")
	_omicsCmd.Flags().BoolVarP(&_omicsListAnnotationStoreVersions, "list-annotation-store-versions", "", false, "List Annotation Store Versions")
	_omicsCmd.Flags().BoolVarP(&_omicsListAnnotationStores, "list-annotation-stores", "", false, "List Annotation Stores")
	_omicsCmd.Flags().BoolVarP(&_omicsListMultipartReadSetUploads, "list-multipart-read-set-uploads", "", false, "List Multipart Read Set Uploads")
	_omicsCmd.Flags().BoolVarP(&_omicsListReadSetActivationJobs, "list-read-set-activation-jobs", "", false, "List Read Set Activation Jobs")
	_omicsCmd.Flags().BoolVarP(&_omicsListReadSetExportJobs, "list-read-set-export-jobs", "", false, "List Read Set Export Jobs")
	_omicsCmd.Flags().BoolVarP(&_omicsListReadSetImportJobs, "list-read-set-import-jobs", "", false, "List Read Set Import Jobs")
	_omicsCmd.Flags().BoolVarP(&_omicsListReadSetUploadParts, "list-read-set-upload-parts", "", false, "List Read Set Upload Parts")
	_omicsCmd.Flags().BoolVarP(&_omicsListReadSets, "list-read-sets", "", false, "List Read Sets")
	_omicsCmd.Flags().BoolVarP(&_omicsListReferenceImportJobs, "list-reference-import-jobs", "", false, "List Reference Import Jobs")
	_omicsCmd.Flags().BoolVarP(&_omicsListReferenceStores, "list-reference-stores", "", false, "List Reference Stores")
	_omicsCmd.Flags().BoolVarP(&_omicsListReferences, "list-references", "", false, "List References")
	_omicsCmd.Flags().BoolVarP(&_omicsListRunCaches, "list-run-caches", "", false, "List Run Caches")
	_omicsCmd.Flags().BoolVarP(&_omicsListRunGroups, "list-run-groups", "", false, "List Run Groups")
	_omicsCmd.Flags().BoolVarP(&_omicsListRunTasks, "list-run-tasks", "", false, "List Run Tasks")
	_omicsCmd.Flags().BoolVarP(&_omicsListRuns, "list-runs", "", false, "List Runs")
	_omicsCmd.Flags().BoolVarP(&_omicsListSequenceStores, "list-sequence-stores", "", false, "List Sequence Stores")
	_omicsCmd.Flags().BoolVarP(&_omicsListShares, "list-shares", "", false, "List Shares")
	_omicsCmd.Flags().BoolVarP(&_omicsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_omicsCmd.Flags().BoolVarP(&_omicsListVariantImportJobs, "list-variant-import-jobs", "", false, "List Variant Import Jobs")
	_omicsCmd.Flags().BoolVarP(&_omicsListVariantStores, "list-variant-stores", "", false, "List Variant Stores")
	_omicsCmd.Flags().BoolVarP(&_omicsListWorkflowVersions, "list-workflow-versions", "", false, "List Workflow Versions")
	_omicsCmd.Flags().BoolVarP(&_omicsListWorkflows, "list-workflows", "", false, "List Workflows")
	_omicsCmd.Flags().BoolVarP(&_omicsPutS3AccessPolicy, "put-s3-access-policy", "", false, "Put S3 Access Policy")
	_omicsCmd.Flags().BoolVarP(&_omicsStartAnnotationImportJob, "start-annotation-import-job", "", false, "Start Annotation Import Job")
	_omicsCmd.Flags().BoolVarP(&_omicsStartReadSetActivationJob, "start-read-set-activation-job", "", false, "Start Read Set Activation Job")
	_omicsCmd.Flags().BoolVarP(&_omicsStartReadSetExportJob, "start-read-set-export-job", "", false, "Start Read Set Export Job")
	_omicsCmd.Flags().BoolVarP(&_omicsStartReadSetImportJob, "start-read-set-import-job", "", false, "Start Read Set Import Job")
	_omicsCmd.Flags().BoolVarP(&_omicsStartReferenceImportJob, "start-reference-import-job", "", false, "Start Reference Import Job")
	_omicsCmd.Flags().BoolVarP(&_omicsStartRun, "start-run", "", false, "Start Run")
	_omicsCmd.Flags().BoolVarP(&_omicsStartVariantImportJob, "start-variant-import-job", "", false, "Start Variant Import Job")
	_omicsCmd.Flags().BoolVarP(&_omicsTagResource, "tag-resource", "", false, "Tag Resource")
	_omicsCmd.Flags().BoolVarP(&_omicsUntagResource, "untag-resource", "", false, "Untag Resource")
	_omicsCmd.Flags().BoolVarP(&_omicsUpdateAnnotationStore, "update-annotation-store", "", false, "Update Annotation Store")
	_omicsCmd.Flags().BoolVarP(&_omicsUpdateAnnotationStoreVersion, "update-annotation-store-version", "", false, "Update Annotation Store Version")
	_omicsCmd.Flags().BoolVarP(&_omicsUpdateRunCache, "update-run-cache", "", false, "Update Run Cache")
	_omicsCmd.Flags().BoolVarP(&_omicsUpdateRunGroup, "update-run-group", "", false, "Update Run Group")
	_omicsCmd.Flags().BoolVarP(&_omicsUpdateSequenceStore, "update-sequence-store", "", false, "Update Sequence Store")
	_omicsCmd.Flags().BoolVarP(&_omicsUpdateVariantStore, "update-variant-store", "", false, "Update Variant Store")
	_omicsCmd.Flags().BoolVarP(&_omicsUpdateWorkflow, "update-workflow", "", false, "Update Workflow")
	_omicsCmd.Flags().BoolVarP(&_omicsUpdateWorkflowVersion, "update-workflow-version", "", false, "Update Workflow Version")
	_omicsCmd.Flags().BoolVarP(&_omicsUploadReadSetPart, "upload-read-set-part", "", false, "Upload Read Set Part")

}
