package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rekognitionCmd represents the rekognition command
var _rekognitionCmd = &cobra.Command{
	Use:   "rekognition",
	Short: "AWS rekognition CLI",
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
		client := rekognition.NewFromConfig(cfg)
		if _rekognitionAssociateFaces {
			rekognition_AssociateFaces(cfg, client)
			return
		}
		if _rekognitionCompareFaces {
			rekognition_CompareFaces(cfg, client)
			return
		}
		if _rekognitionCopyProjectVersion {
			rekognition_CopyProjectVersion(cfg, client)
			return
		}
		if _rekognitionCreateCollection {
			rekognition_CreateCollection(cfg, client)
			return
		}
		if _rekognitionCreateDataset {
			rekognition_CreateDataset(cfg, client)
			return
		}
		if _rekognitionCreateFaceLivenessSession {
			rekognition_CreateFaceLivenessSession(cfg, client)
			return
		}
		if _rekognitionCreateProject {
			rekognition_CreateProject(cfg, client)
			return
		}
		if _rekognitionCreateProjectVersion {
			rekognition_CreateProjectVersion(cfg, client)
			return
		}
		if _rekognitionCreateStreamProcessor {
			rekognition_CreateStreamProcessor(cfg, client)
			return
		}
		if _rekognitionCreateUser {
			rekognition_CreateUser(cfg, client)
			return
		}
		if _rekognitionDeleteCollection {
			rekognition_DeleteCollection(cfg, client)
			return
		}
		if _rekognitionDeleteDataset {
			rekognition_DeleteDataset(cfg, client)
			return
		}
		if _rekognitionDeleteFaces {
			rekognition_DeleteFaces(cfg, client)
			return
		}
		if _rekognitionDeleteProject {
			rekognition_DeleteProject(cfg, client)
			return
		}
		if _rekognitionDeleteProjectPolicy {
			rekognition_DeleteProjectPolicy(cfg, client)
			return
		}
		if _rekognitionDeleteProjectVersion {
			rekognition_DeleteProjectVersion(cfg, client)
			return
		}
		if _rekognitionDeleteStreamProcessor {
			rekognition_DeleteStreamProcessor(cfg, client)
			return
		}
		if _rekognitionDeleteUser {
			rekognition_DeleteUser(cfg, client)
			return
		}
		if _rekognitionDescribeCollection {
			rekognition_DescribeCollection(cfg, client)
			return
		}
		if _rekognitionDescribeDataset {
			rekognition_DescribeDataset(cfg, client)
			return
		}
		if _rekognitionDescribeProjectVersions {
			rekognition_DescribeProjectVersions(cfg, client)
			return
		}
		if _rekognitionDescribeProjects {
			rekognition_DescribeProjects(cfg, client)
			return
		}
		if _rekognitionDescribeStreamProcessor {
			rekognition_DescribeStreamProcessor(cfg, client)
			return
		}
		if _rekognitionDetectCustomLabels {
			rekognition_DetectCustomLabels(cfg, client)
			return
		}
		if _rekognitionDetectFaces {
			rekognition_DetectFaces(cfg, client)
			return
		}
		if _rekognitionDetectLabels {
			rekognition_DetectLabels(cfg, client)
			return
		}
		if _rekognitionDetectModerationLabels {
			rekognition_DetectModerationLabels(cfg, client)
			return
		}
		if _rekognitionDetectProtectiveEquipment {
			rekognition_DetectProtectiveEquipment(cfg, client)
			return
		}
		if _rekognitionDetectText {
			rekognition_DetectText(cfg, client)
			return
		}
		if _rekognitionDisassociateFaces {
			rekognition_DisassociateFaces(cfg, client)
			return
		}
		if _rekognitionDistributeDatasetEntries {
			rekognition_DistributeDatasetEntries(cfg, client)
			return
		}
		if _rekognitionGetCelebrityInfo {
			rekognition_GetCelebrityInfo(cfg, client)
			return
		}
		if _rekognitionGetCelebrityRecognition {
			rekognition_GetCelebrityRecognition(cfg, client)
			return
		}
		if _rekognitionGetContentModeration {
			rekognition_GetContentModeration(cfg, client)
			return
		}
		if _rekognitionGetFaceDetection {
			rekognition_GetFaceDetection(cfg, client)
			return
		}
		if _rekognitionGetFaceLivenessSessionResults {
			rekognition_GetFaceLivenessSessionResults(cfg, client)
			return
		}
		if _rekognitionGetFaceSearch {
			rekognition_GetFaceSearch(cfg, client)
			return
		}
		if _rekognitionGetLabelDetection {
			rekognition_GetLabelDetection(cfg, client)
			return
		}
		if _rekognitionGetMediaAnalysisJob {
			rekognition_GetMediaAnalysisJob(cfg, client)
			return
		}
		if _rekognitionGetPersonTracking {
			rekognition_GetPersonTracking(cfg, client)
			return
		}
		if _rekognitionGetSegmentDetection {
			rekognition_GetSegmentDetection(cfg, client)
			return
		}
		if _rekognitionGetTextDetection {
			rekognition_GetTextDetection(cfg, client)
			return
		}
		if _rekognitionIndexFaces {
			rekognition_IndexFaces(cfg, client)
			return
		}
		if _rekognitionListCollections {
			rekognition_ListCollections(cfg, client)
			return
		}
		if _rekognitionListDatasetEntries {
			rekognition_ListDatasetEntries(cfg, client)
			return
		}
		if _rekognitionListDatasetLabels {
			rekognition_ListDatasetLabels(cfg, client)
			return
		}
		if _rekognitionListFaces {
			rekognition_ListFaces(cfg, client)
			return
		}
		if _rekognitionListMediaAnalysisJobs {
			rekognition_ListMediaAnalysisJobs(cfg, client)
			return
		}
		if _rekognitionListProjectPolicies {
			rekognition_ListProjectPolicies(cfg, client)
			return
		}
		if _rekognitionListStreamProcessors {
			rekognition_ListStreamProcessors(cfg, client)
			return
		}
		if _rekognitionListTagsForResource {
			rekognition_ListTagsForResource(cfg, client)
			return
		}
		if _rekognitionListUsers {
			rekognition_ListUsers(cfg, client)
			return
		}
		if _rekognitionPutProjectPolicy {
			rekognition_PutProjectPolicy(cfg, client)
			return
		}
		if _rekognitionRecognizeCelebrities {
			rekognition_RecognizeCelebrities(cfg, client)
			return
		}
		if _rekognitionSearchFaces {
			rekognition_SearchFaces(cfg, client)
			return
		}
		if _rekognitionSearchFacesByImage {
			rekognition_SearchFacesByImage(cfg, client)
			return
		}
		if _rekognitionSearchUsers {
			rekognition_SearchUsers(cfg, client)
			return
		}
		if _rekognitionSearchUsersByImage {
			rekognition_SearchUsersByImage(cfg, client)
			return
		}
		if _rekognitionStartCelebrityRecognition {
			rekognition_StartCelebrityRecognition(cfg, client)
			return
		}
		if _rekognitionStartContentModeration {
			rekognition_StartContentModeration(cfg, client)
			return
		}
		if _rekognitionStartFaceDetection {
			rekognition_StartFaceDetection(cfg, client)
			return
		}
		if _rekognitionStartFaceSearch {
			rekognition_StartFaceSearch(cfg, client)
			return
		}
		if _rekognitionStartLabelDetection {
			rekognition_StartLabelDetection(cfg, client)
			return
		}
		if _rekognitionStartMediaAnalysisJob {
			rekognition_StartMediaAnalysisJob(cfg, client)
			return
		}
		if _rekognitionStartPersonTracking {
			rekognition_StartPersonTracking(cfg, client)
			return
		}
		if _rekognitionStartProjectVersion {
			rekognition_StartProjectVersion(cfg, client)
			return
		}
		if _rekognitionStartSegmentDetection {
			rekognition_StartSegmentDetection(cfg, client)
			return
		}
		if _rekognitionStartStreamProcessor {
			rekognition_StartStreamProcessor(cfg, client)
			return
		}
		if _rekognitionStartTextDetection {
			rekognition_StartTextDetection(cfg, client)
			return
		}
		if _rekognitionStopProjectVersion {
			rekognition_StopProjectVersion(cfg, client)
			return
		}
		if _rekognitionStopStreamProcessor {
			rekognition_StopStreamProcessor(cfg, client)
			return
		}
		if _rekognitionTagResource {
			rekognition_TagResource(cfg, client)
			return
		}
		if _rekognitionUntagResource {
			rekognition_UntagResource(cfg, client)
			return
		}
		if _rekognitionUpdateDatasetEntries {
			rekognition_UpdateDatasetEntries(cfg, client)
			return
		}
		if _rekognitionUpdateStreamProcessor {
			rekognition_UpdateStreamProcessor(cfg, client)
			return
		}

	},
}

var (
	_rekognitionAssociateFaces                bool
	_rekognitionCompareFaces                  bool
	_rekognitionCopyProjectVersion            bool
	_rekognitionCreateCollection              bool
	_rekognitionCreateDataset                 bool
	_rekognitionCreateFaceLivenessSession     bool
	_rekognitionCreateProject                 bool
	_rekognitionCreateProjectVersion          bool
	_rekognitionCreateStreamProcessor         bool
	_rekognitionCreateUser                    bool
	_rekognitionDeleteCollection              bool
	_rekognitionDeleteDataset                 bool
	_rekognitionDeleteFaces                   bool
	_rekognitionDeleteProject                 bool
	_rekognitionDeleteProjectPolicy           bool
	_rekognitionDeleteProjectVersion          bool
	_rekognitionDeleteStreamProcessor         bool
	_rekognitionDeleteUser                    bool
	_rekognitionDescribeCollection            bool
	_rekognitionDescribeDataset               bool
	_rekognitionDescribeProjectVersions       bool
	_rekognitionDescribeProjects              bool
	_rekognitionDescribeStreamProcessor       bool
	_rekognitionDetectCustomLabels            bool
	_rekognitionDetectFaces                   bool
	_rekognitionDetectLabels                  bool
	_rekognitionDetectModerationLabels        bool
	_rekognitionDetectProtectiveEquipment     bool
	_rekognitionDetectText                    bool
	_rekognitionDisassociateFaces             bool
	_rekognitionDistributeDatasetEntries      bool
	_rekognitionGetCelebrityInfo              bool
	_rekognitionGetCelebrityRecognition       bool
	_rekognitionGetContentModeration          bool
	_rekognitionGetFaceDetection              bool
	_rekognitionGetFaceLivenessSessionResults bool
	_rekognitionGetFaceSearch                 bool
	_rekognitionGetLabelDetection             bool
	_rekognitionGetMediaAnalysisJob           bool
	_rekognitionGetPersonTracking             bool
	_rekognitionGetSegmentDetection           bool
	_rekognitionGetTextDetection              bool
	_rekognitionIndexFaces                    bool
	_rekognitionListCollections               bool
	_rekognitionListDatasetEntries            bool
	_rekognitionListDatasetLabels             bool
	_rekognitionListFaces                     bool
	_rekognitionListMediaAnalysisJobs         bool
	_rekognitionListProjectPolicies           bool
	_rekognitionListStreamProcessors          bool
	_rekognitionListTagsForResource           bool
	_rekognitionListUsers                     bool
	_rekognitionPutProjectPolicy              bool
	_rekognitionRecognizeCelebrities          bool
	_rekognitionSearchFaces                   bool
	_rekognitionSearchFacesByImage            bool
	_rekognitionSearchUsers                   bool
	_rekognitionSearchUsersByImage            bool
	_rekognitionStartCelebrityRecognition     bool
	_rekognitionStartContentModeration        bool
	_rekognitionStartFaceDetection            bool
	_rekognitionStartFaceSearch               bool
	_rekognitionStartLabelDetection           bool
	_rekognitionStartMediaAnalysisJob         bool
	_rekognitionStartPersonTracking           bool
	_rekognitionStartProjectVersion           bool
	_rekognitionStartSegmentDetection         bool
	_rekognitionStartStreamProcessor          bool
	_rekognitionStartTextDetection            bool
	_rekognitionStopProjectVersion            bool
	_rekognitionStopStreamProcessor           bool
	_rekognitionTagResource                   bool
	_rekognitionUntagResource                 bool
	_rekognitionUpdateDatasetEntries          bool
	_rekognitionUpdateStreamProcessor         bool

	_rekognitionAggregateBy                    string
	_rekognitionAttributes                     string
	_rekognitionAutoUpdate                     string
	_rekognitionChanges                        string
	_rekognitionClientRequestToken             string
	_rekognitionCollectionId                   string
	_rekognitionContainsLabels                 []string
	_rekognitionDataSharingPreference          string
	_rekognitionDataSharingPreferenceForUpdate string
	_rekognitionDatasetArn                     string
	_rekognitionDatasetSource                  string
	_rekognitionDatasetType                    string
	_rekognitionDatasets                       string
	_rekognitionDestinationProjectArn          string
	_rekognitionDetectionAttributes            string
	_rekognitionExternalImageId                string
	_rekognitionFaceAttributes                 string
	_rekognitionFaceId                         string
	_rekognitionFaceIds                        []string
	_rekognitionFaceMatchThreshold             string
	_rekognitionFeature                        string
	_rekognitionFeatureConfig                  string
	_rekognitionFeatures                       string
	_rekognitionFilters                        string
	_rekognitionHasErrors                      string
	_rekognitionHumanLoopConfig                string
	_rekognitionId                             string
	_rekognitionImage                          string
	_rekognitionInput                          string
	_rekognitionJobId                          string
	_rekognitionJobName                        string
	_rekognitionJobTag                         string
	_rekognitionKmsKeyId                       string
	_rekognitionLabeled                        string
	_rekognitionMaxFaces                       string
	_rekognitionMaxInferenceUnits              string
	_rekognitionMaxLabels                      string
	_rekognitionMaxResults                     string
	_rekognitionMaxUsers                       string
	_rekognitionMinConfidence                  string
	_rekognitionMinInferenceUnits              string
	_rekognitionName                           string
	_rekognitionNextToken                      string
	_rekognitionNotificationChannel            string
	_rekognitionOperationsConfig               string
	_rekognitionOutputConfig                   string
	_rekognitionParametersToDelete             string
	_rekognitionPolicyDocument                 string
	_rekognitionPolicyName                     string
	_rekognitionPolicyRevisionId               string
	_rekognitionProjectArn                     string
	_rekognitionProjectName                    string
	_rekognitionProjectNames                   []string
	_rekognitionProjectVersion                 string
	_rekognitionProjectVersionArn              string
	_rekognitionQualityFilter                  string
	_rekognitionRegionsOfInterest              string
	_rekognitionRegionsOfInterestForUpdate     string
	_rekognitionResourceArn                    string
	_rekognitionRoleArn                        string
	_rekognitionSegmentTypes                   string
	_rekognitionSessionId                      string
	_rekognitionSettings                       string
	_rekognitionSettingsForUpdate              string
	_rekognitionSimilarityThreshold            string
	_rekognitionSortBy                         string
	_rekognitionSourceImage                    string
	_rekognitionSourceProjectArn               string
	_rekognitionSourceProjectVersionArn        string
	_rekognitionSourceRefContains              string
	_rekognitionStartSelector                  string
	_rekognitionStopSelector                   string
	_rekognitionSummarizationAttributes        string
	_rekognitionTagKeys                        []string
	_rekognitionTags                           string
	_rekognitionTargetImage                    string
	_rekognitionTestingData                    string
	_rekognitionTrainingData                   string
	_rekognitionUserId                         string
	_rekognitionUserMatchThreshold             string
	_rekognitionVersionDescription             string
	_rekognitionVersionName                    string
	_rekognitionVersionNames                   []string
	_rekognitionVideo                          string
)

// Associates one or more faces with an existing UserID. Takes an array of FaceIds
// . Each FaceId that are present in the FaceIds list is associated with the
// provided UserID. The number of FaceIds that can be used as input in a single
// request is limited to 100.
//
// Note that the total number of faces that can be associated with a single UserID
// is also limited to 100. Once a UserID has 100 faces associated with it, no
// additional faces can be added. If more API calls are made after the limit is
// reached, a ServiceQuotaExceededException will result.
//
// The UserMatchThreshold parameter specifies the minimum user match confidence
// required for the face to be associated with a UserID that has at least one
// FaceID already associated. This ensures that the FaceIds are associated with
// the right UserID. The value ranges from 0-100 and default value is 75.
//
// If successful, an array of AssociatedFace objects containing the associated
// FaceIds is returned. If a given face is already associated with the given UserID
// , it will be ignored and will not be returned in the response. If a given face
// is already associated to a different UserID , isn't found in the collection,
// doesn’t meet the UserMatchThreshold , or there are already 100 faces associated
// with the UserID , it will be returned as part of an array of
// UnsuccessfulFaceAssociations.
//
// The UserStatus reflects the status of an operation which updates a UserID
// representation with a list of given faces. The UserStatus can be:
//
// - ACTIVE - All associations or disassociations of FaceID(s) for a UserID are
// complete.
//
// - CREATED - A UserID has been created, but has no FaceID(s) associated with
// it.
//
// - UPDATING - A UserID is being updated and there are current associations or
// disassociations of FaceID(s) taking place.
func rekognition_AssociateFaces(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.AssociateFacesInput{
		// CollectionId: *string, // Required
		// FaceIds: []string, // Required
		// UserId: *string, // Required
	}

	if len(_rekognitionCollectionId) > 0 {
		input.CollectionId = aws.String(_rekognitionCollectionId)
	}
	if len(_rekognitionFaceIds) > 0 {
		input.FaceIds = append([]string(nil), _rekognitionFaceIds...)
	}
	if len(_rekognitionUserId) > 0 {
		input.UserId = aws.String(_rekognitionUserId)
	}
	if len(_rekognitionClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_rekognitionClientRequestToken)
	}
	if len(_rekognitionUserMatchThreshold) > 0 {
		if err := assignInputField(input, "UserMatchThreshold", _rekognitionUserMatchThreshold); err != nil {
			log.Errorf("invalid --user-match-threshold: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateFaces(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Compares a face in the source input image with each of the 100 largest faces
// detected in the target input image.
//
// If the source image contains multiple faces, the service detects the largest
// face and compares it with each face detected in the target image.
//
// CompareFaces uses machine learning algorithms, which are probabilistic. A false
// negative is an incorrect prediction that a face in the target image has a low
// similarity confidence score when compared to the face in the source image. To
// reduce the probability of false negatives, we recommend that you compare the
// target image against multiple source images. If you plan to use CompareFaces to
// make a decision that impacts an individual's rights, privacy, or access to
// services, we recommend that you pass the result to a human for review and
// further validation before taking action.
//
// You pass the input and target images either as base64-encoded image bytes or as
// references to images in an Amazon S3 bucket. If you use the AWS CLI to call
// Amazon Rekognition operations, passing image bytes isn't supported. The image
// must be formatted as a PNG or JPEG file.
//
// In response, the operation returns an array of face matches ordered by
// similarity score in descending order. For each face match, the response provides
// a bounding box of the face, facial landmarks, pose details (pitch, roll, and
// yaw), quality (brightness and sharpness), and confidence value (indicating the
// level of confidence that the bounding box contains a face). The response also
// provides a similarity score, which indicates how closely the faces match.
//
// By default, only faces with a similarity score of greater than or equal to 80%
// are returned in the response. You can change this value by specifying the
// SimilarityThreshold parameter.
//
// CompareFaces also returns an array of faces that don't match the source image.
// For each face, it returns a bounding box, confidence value, landmarks, pose
// details, and quality. The response also returns information about the face in
// the source image, including the bounding box of the face and confidence value.
//
// The QualityFilter input parameter allows you to filter out detected faces that
// don’t meet a required quality bar. The quality bar is based on a variety of
// common use cases. Use QualityFilter to set the quality bar by specifying LOW ,
// MEDIUM , or HIGH . If you do not want to filter detected faces, specify NONE .
// The default value is NONE .
//
// If the image doesn't contain Exif metadata, CompareFaces returns orientation
// information for the source and target images. Use these values to display the
// images with the correct image orientation.
//
// If no faces are detected in the source or target images, CompareFaces returns
// an InvalidParameterException error.
//
// This is a stateless API operation. That is, data returned by this operation
// doesn't persist.
//
// For an example, see Comparing Faces in Images in the Amazon Rekognition
// Developer Guide.
//
// This operation requires permissions to perform the rekognition:CompareFaces
// action.
func rekognition_CompareFaces(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.CompareFacesInput{
		// SourceImage: *types.Image, // Required
		// TargetImage: *types.Image, // Required
	}

	if len(_rekognitionSourceImage) > 0 {
		if err := assignInputField(input, "SourceImage", _rekognitionSourceImage); err != nil {
			log.Errorf("invalid --source-image: %s", err.Error())
			return
		}
	}
	if len(_rekognitionTargetImage) > 0 {
		if err := assignInputField(input, "TargetImage", _rekognitionTargetImage); err != nil {
			log.Errorf("invalid --target-image: %s", err.Error())
			return
		}
	}
	if len(_rekognitionQualityFilter) > 0 {
		if err := assignInputField(input, "QualityFilter", _rekognitionQualityFilter); err != nil {
			log.Errorf("invalid --quality-filter: %s", err.Error())
			return
		}
	}
	if len(_rekognitionSimilarityThreshold) > 0 {
		if err := assignInputField(input, "SimilarityThreshold", _rekognitionSimilarityThreshold); err != nil {
			log.Errorf("invalid --similarity-threshold: %s", err.Error())
			return
		}
	}

	if resp, err := client.CompareFaces(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation applies only to Amazon Rekognition Custom Labels.
// Copies a version of an Amazon Rekognition Custom Labels model from a source
// project to a destination project. The source and destination projects can be in
// different AWS accounts but must be in the same AWS Region. You can't copy a
// model to another AWS service.
//
// To copy a model version to a different AWS account, you need to create a
// resource-based policy known as a project policy. You attach the project policy
// to the source project by calling PutProjectPolicy. The project policy gives permission to copy
// the model version from a trusting AWS account to a trusted account.
//
// For more information creating and attaching a project policy, see Attaching a
// project policy (SDK) in the Amazon Rekognition Custom Labels Developer Guide.
//
// If you are copying a model version to a project in the same AWS account, you
// don't need to create a project policy.
//
// Copying project versions is supported only for Custom Labels models.
//
// To copy a model, the destination project, source project, and source model
// version must already exist.
//
// Copying a model version takes a while to complete. To get the current status,
// call DescribeProjectVersionsand check the value of Status in the ProjectVersionDescription object. The copy operation has
// finished when the value of Status is COPYING_COMPLETED .
//
// This operation requires permissions to perform the
// rekognition:CopyProjectVersion action.
func rekognition_CopyProjectVersion(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.CopyProjectVersionInput{
		// DestinationProjectArn: *string, // Required
		// OutputConfig: *types.OutputConfig, // Required
		// SourceProjectArn: *string, // Required
		// SourceProjectVersionArn: *string, // Required
		// VersionName: *string, // Required
	}

	if len(_rekognitionDestinationProjectArn) > 0 {
		input.DestinationProjectArn = aws.String(_rekognitionDestinationProjectArn)
	}
	if len(_rekognitionOutputConfig) > 0 {
		if err := assignInputField(input, "OutputConfig", _rekognitionOutputConfig); err != nil {
			log.Errorf("invalid --output-config: %s", err.Error())
			return
		}
	}
	if len(_rekognitionSourceProjectArn) > 0 {
		input.SourceProjectArn = aws.String(_rekognitionSourceProjectArn)
	}
	if len(_rekognitionSourceProjectVersionArn) > 0 {
		input.SourceProjectVersionArn = aws.String(_rekognitionSourceProjectVersionArn)
	}
	if len(_rekognitionVersionName) > 0 {
		input.VersionName = aws.String(_rekognitionVersionName)
	}
	if len(_rekognitionKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_rekognitionKmsKeyId)
	}
	if len(_rekognitionTags) > 0 {
		if err := assignInputField(input, "Tags", _rekognitionTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CopyProjectVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a collection in an AWS Region. You can add faces to the collection
// using the IndexFacesoperation.
//
// For example, you might create collections, one for each of your application
// users. A user can then index faces using the IndexFaces operation and persist
// results in a specific collection. Then, a user can search the collection for
// faces in the user-specific container.
//
// When you create a collection, it is associated with the latest version of the
// face model version.
//
// Collection names are case-sensitive.
//
// This operation requires permissions to perform the rekognition:CreateCollection
// action. If you want to tag your collection, you also require permission to
// perform the rekognition:TagResource operation.
func rekognition_CreateCollection(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.CreateCollectionInput{
		// CollectionId: *string, // Required
	}

	if len(_rekognitionCollectionId) > 0 {
		input.CollectionId = aws.String(_rekognitionCollectionId)
	}
	if len(_rekognitionTags) > 0 {
		if err := assignInputField(input, "Tags", _rekognitionTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCollection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation applies only to Amazon Rekognition Custom Labels.
// Creates a new Amazon Rekognition Custom Labels dataset. You can create a
// dataset by using an Amazon Sagemaker format manifest file or by copying an
// existing Amazon Rekognition Custom Labels dataset.
//
// To create a training dataset for a project, specify TRAIN for the value of
// DatasetType . To create the test dataset for a project, specify TEST for the
// value of DatasetType .
//
// The response from CreateDataset is the Amazon Resource Name (ARN) for the
// dataset. Creating a dataset takes a while to complete. Use DescribeDatasetto check the current
// status. The dataset created successfully if the value of Status is
// CREATE_COMPLETE .
//
// To check if any non-terminal errors occurred, call ListDatasetEntries and check for the presence
// of errors lists in the JSON Lines.
//
// Dataset creation fails if a terminal error occurs ( Status = CREATE_FAILED ).
// Currently, you can't access the terminal error information.
//
// For more information, see Creating dataset in the Amazon Rekognition Custom
// Labels Developer Guide.
//
// This operation requires permissions to perform the rekognition:CreateDataset
// action. If you want to copy an existing dataset, you also require permission to
// perform the rekognition:ListDatasetEntries action.
func rekognition_CreateDataset(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.CreateDatasetInput{
		// DatasetType: types.DatasetType, // Required
		// ProjectArn: *string, // Required
	}

	if len(_rekognitionDatasetType) > 0 {
		if err := assignInputField(input, "DatasetType", _rekognitionDatasetType); err != nil {
			log.Errorf("invalid --dataset-type: %s", err.Error())
			return
		}
	}
	if len(_rekognitionProjectArn) > 0 {
		input.ProjectArn = aws.String(_rekognitionProjectArn)
	}
	if len(_rekognitionDatasetSource) > 0 {
		if err := assignInputField(input, "DatasetSource", _rekognitionDatasetSource); err != nil {
			log.Errorf("invalid --dataset-source: %s", err.Error())
			return
		}
	}
	if len(_rekognitionTags) > 0 {
		if err := assignInputField(input, "Tags", _rekognitionTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API operation initiates a Face Liveness session. It returns a SessionId ,
// which you can use to start streaming Face Liveness video and get the results for
// a Face Liveness session.
//
// You can use the OutputConfig option in the Settings parameter to provide an
// Amazon S3 bucket location. The Amazon S3 bucket stores reference images and
// audit images. If no Amazon S3 bucket is defined, raw bytes are sent instead.
//
// You can use AuditImagesLimit to limit the number of audit images returned when
// GetFaceLivenessSessionResults is called. This number is between 0 and 4. By
// default, it is set to 0. The limit is best effort and based on the duration of
// the selfie-video.
func rekognition_CreateFaceLivenessSession(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.CreateFaceLivenessSessionInput{}

	if len(_rekognitionClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_rekognitionClientRequestToken)
	}
	if len(_rekognitionKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_rekognitionKmsKeyId)
	}
	if len(_rekognitionSettings) > 0 {
		if err := assignInputField(input, "Settings", _rekognitionSettings); err != nil {
			log.Errorf("invalid --settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFaceLivenessSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Amazon Rekognition project. A project is a group of resources
// (datasets, model versions) that you use to create and manage a Amazon
// Rekognition Custom Labels Model or custom adapter. You can specify a feature to
// create the project with, if no feature is specified then Custom Labels is used
// by default. For adapters, you can also choose whether or not to have the project
// auto update by using the AutoUpdate argument. This operation requires
// permissions to perform the rekognition:CreateProject action.
func rekognition_CreateProject(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.CreateProjectInput{
		// ProjectName: *string, // Required
	}

	if len(_rekognitionProjectName) > 0 {
		input.ProjectName = aws.String(_rekognitionProjectName)
	}
	if len(_rekognitionAutoUpdate) > 0 {
		if err := assignInputField(input, "AutoUpdate", _rekognitionAutoUpdate); err != nil {
			log.Errorf("invalid --auto-update: %s", err.Error())
			return
		}
	}
	if len(_rekognitionFeature) > 0 {
		if err := assignInputField(input, "Feature", _rekognitionFeature); err != nil {
			log.Errorf("invalid --feature: %s", err.Error())
			return
		}
	}
	if len(_rekognitionTags) > 0 {
		if err := assignInputField(input, "Tags", _rekognitionTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new version of Amazon Rekognition project (like a Custom Labels model
// or a custom adapter) and begins training. Models and adapters are managed as
// part of a Rekognition project. The response from CreateProjectVersion is an
// Amazon Resource Name (ARN) for the project version.
//
// The FeatureConfig operation argument allows you to configure specific model or
// adapter settings. You can provide a description to the project version by using
// the VersionDescription argment. Training can take a while to complete. You can
// get the current status by calling DescribeProjectVersions. Training completed successfully if the
// value of the Status field is TRAINING_COMPLETED . Once training has successfully
// completed, call DescribeProjectVersionsto get the training results and evaluate the model.
//
// This operation requires permissions to perform the
// rekognition:CreateProjectVersion action.
//
// The following applies only to projects with Amazon Rekognition Custom Labels as
// the chosen feature:
//
// You can train a model in a project that doesn't have associated datasets by
// specifying manifest files in the TrainingData and TestingData fields.
//
// If you open the console after training a model with manifest files, Amazon
// Rekognition Custom Labels creates the datasets for you using the most recent
// manifest files. You can no longer train a model version for the project by
// specifying manifest files.
//
// Instead of training with a project without associated datasets, we recommend
// that you use the manifest files to create training and test datasets for the
// project.
func rekognition_CreateProjectVersion(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.CreateProjectVersionInput{
		// OutputConfig: *types.OutputConfig, // Required
		// ProjectArn: *string, // Required
		// VersionName: *string, // Required
	}

	if len(_rekognitionOutputConfig) > 0 {
		if err := assignInputField(input, "OutputConfig", _rekognitionOutputConfig); err != nil {
			log.Errorf("invalid --output-config: %s", err.Error())
			return
		}
	}
	if len(_rekognitionProjectArn) > 0 {
		input.ProjectArn = aws.String(_rekognitionProjectArn)
	}
	if len(_rekognitionVersionName) > 0 {
		input.VersionName = aws.String(_rekognitionVersionName)
	}
	if len(_rekognitionFeatureConfig) > 0 {
		if err := assignInputField(input, "FeatureConfig", _rekognitionFeatureConfig); err != nil {
			log.Errorf("invalid --feature-config: %s", err.Error())
			return
		}
	}
	if len(_rekognitionKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_rekognitionKmsKeyId)
	}
	if len(_rekognitionTags) > 0 {
		if err := assignInputField(input, "Tags", _rekognitionTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_rekognitionTestingData) > 0 {
		if err := assignInputField(input, "TestingData", _rekognitionTestingData); err != nil {
			log.Errorf("invalid --testing-data: %s", err.Error())
			return
		}
	}
	if len(_rekognitionTrainingData) > 0 {
		if err := assignInputField(input, "TrainingData", _rekognitionTrainingData); err != nil {
			log.Errorf("invalid --training-data: %s", err.Error())
			return
		}
	}
	if len(_rekognitionVersionDescription) > 0 {
		input.VersionDescription = aws.String(_rekognitionVersionDescription)
	}

	if resp, err := client.CreateProjectVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Rekognition stream processor that you can use to detect and
// recognize faces or to detect labels in a streaming video.
//
// Amazon Rekognition Video is a consumer of live video from Amazon Kinesis Video
// Streams. There are two different settings for stream processors in Amazon
// Rekognition: detecting faces and detecting labels.
//
// - If you are creating a stream processor for detecting faces, you provide as
// input a Kinesis video stream ( Input ) and a Kinesis data stream ( Output )
// stream for receiving the output. You must use the FaceSearch option in
// Settings , specifying the collection that contains the faces you want to
// recognize. After you have finished analyzing a streaming video, use StopStreamProcessorto stop
// processing.
//
// - If you are creating a stream processor to detect labels, you provide as
// input a Kinesis video stream ( Input ), Amazon S3 bucket information ( Output
// ), and an Amazon SNS topic ARN ( NotificationChannel ). You can also provide a
// KMS key ID to encrypt the data sent to your Amazon S3 bucket. You specify what
// you want to detect by using the ConnectedHome option in settings, and
// selecting one of the following: PERSON , PET , PACKAGE , ALL You can also
// specify where in the frame you want Amazon Rekognition to monitor with
// RegionsOfInterest . When you run the StartStreamProcessoroperation on a label detection stream
// processor, you input start and stop information to determine the length of the
// processing time.
//
// Use Name to assign an identifier for the stream processor. You use Name to
// manage the stream processor. For example, you can start processing the source
// video by calling StartStreamProcessorwith the Name field.
//
// This operation requires permissions to perform the
// rekognition:CreateStreamProcessor action. If you want to tag your stream
// processor, you also require permission to perform the rekognition:TagResource
// operation.
func rekognition_CreateStreamProcessor(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.CreateStreamProcessorInput{
		// Input: *types.StreamProcessorInput, // Required
		// Name: *string, // Required
		// Output: *types.StreamProcessorOutput, // Required
		// RoleArn: *string, // Required
		// Settings: *types.StreamProcessorSettings, // Required
	}

	if len(_rekognitionInput) > 0 {
		if err := assignInputField(input, "Input", _rekognitionInput); err != nil {
			log.Errorf("invalid --input: %s", err.Error())
			return
		}
	}
	if len(_rekognitionName) > 0 {
		input.Name = aws.String(_rekognitionName)
	}
	if len(_rekognitionRoleArn) > 0 {
		input.RoleArn = aws.String(_rekognitionRoleArn)
	}
	if len(_rekognitionSettings) > 0 {
		if err := assignInputField(input, "Settings", _rekognitionSettings); err != nil {
			log.Errorf("invalid --settings: %s", err.Error())
			return
		}
	}
	if len(_rekognitionDataSharingPreference) > 0 {
		if err := assignInputField(input, "DataSharingPreference", _rekognitionDataSharingPreference); err != nil {
			log.Errorf("invalid --data-sharing-preference: %s", err.Error())
			return
		}
	}
	if len(_rekognitionKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_rekognitionKmsKeyId)
	}
	if len(_rekognitionNotificationChannel) > 0 {
		if err := assignInputField(input, "NotificationChannel", _rekognitionNotificationChannel); err != nil {
			log.Errorf("invalid --notification-channel: %s", err.Error())
			return
		}
	}
	if len(_rekognitionRegionsOfInterest) > 0 {
		if err := assignInputField(input, "RegionsOfInterest", _rekognitionRegionsOfInterest); err != nil {
			log.Errorf("invalid --regions-of-interest: %s", err.Error())
			return
		}
	}
	if len(_rekognitionTags) > 0 {
		if err := assignInputField(input, "Tags", _rekognitionTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateStreamProcessor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new User within a collection specified by CollectionId . Takes UserId
// as a parameter, which is a user provided ID which should be unique within the
// collection. The provided UserId will alias the system generated UUID to make
// the UserId more user friendly.
//
// Uses a ClientToken , an idempotency token that ensures a call to CreateUser
// completes only once. If the value is not supplied, the AWS SDK generates an
// idempotency token for the requests. This prevents retries after a network error
// results from making multiple CreateUser calls.
func rekognition_CreateUser(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.CreateUserInput{
		// CollectionId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_rekognitionCollectionId) > 0 {
		input.CollectionId = aws.String(_rekognitionCollectionId)
	}
	if len(_rekognitionUserId) > 0 {
		input.UserId = aws.String(_rekognitionUserId)
	}
	if len(_rekognitionClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_rekognitionClientRequestToken)
	}

	if resp, err := client.CreateUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified collection. Note that this operation removes all faces in
// the collection. For an example, see [Deleting a collection].
//
// This operation requires permissions to perform the rekognition:DeleteCollection
// action.
//
// [Deleting a collection]: https://docs.aws.amazon.com/rekognition/latest/dg/delete-collection-procedure.html
func rekognition_DeleteCollection(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.DeleteCollectionInput{
		// CollectionId: *string, // Required
	}

	if len(_rekognitionCollectionId) > 0 {
		input.CollectionId = aws.String(_rekognitionCollectionId)
	}

	if resp, err := client.DeleteCollection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation applies only to Amazon Rekognition Custom Labels.
// Deletes an existing Amazon Rekognition Custom Labels dataset. Deleting a
// dataset might take while. Use DescribeDatasetto check the current status. The dataset is still
// deleting if the value of Status is DELETE_IN_PROGRESS . If you try to access the
// dataset after it is deleted, you get a ResourceNotFoundException exception.
//
// You can't delete a dataset while it is creating ( Status = CREATE_IN_PROGRESS )
// or if the dataset is updating ( Status = UPDATE_IN_PROGRESS ).
//
// This operation requires permissions to perform the rekognition:DeleteDataset
// action.
func rekognition_DeleteDataset(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.DeleteDatasetInput{
		// DatasetArn: *string, // Required
	}

	if len(_rekognitionDatasetArn) > 0 {
		input.DatasetArn = aws.String(_rekognitionDatasetArn)
	}

	if resp, err := client.DeleteDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes faces from a collection. You specify a collection ID and an array of
// face IDs to remove from the collection.
//
// This operation requires permissions to perform the rekognition:DeleteFaces
// action.
func rekognition_DeleteFaces(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.DeleteFacesInput{
		// CollectionId: *string, // Required
		// FaceIds: []string, // Required
	}

	if len(_rekognitionCollectionId) > 0 {
		input.CollectionId = aws.String(_rekognitionCollectionId)
	}
	if len(_rekognitionFaceIds) > 0 {
		input.FaceIds = append([]string(nil), _rekognitionFaceIds...)
	}

	if resp, err := client.DeleteFaces(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Amazon Rekognition project. To delete a project you must first delete
// all models or adapters associated with the project. To delete a model or
// adapter, see DeleteProjectVersion.
//
// DeleteProject is an asynchronous operation. To check if the project is deleted,
// call DescribeProjects. The project is deleted when the project no longer appears in the
// response. Be aware that deleting a given project will also delete any
// ProjectPolicies associated with that project.
//
// This operation requires permissions to perform the rekognition:DeleteProject
// action.
func rekognition_DeleteProject(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.DeleteProjectInput{
		// ProjectArn: *string, // Required
	}

	if len(_rekognitionProjectArn) > 0 {
		input.ProjectArn = aws.String(_rekognitionProjectArn)
	}

	if resp, err := client.DeleteProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation applies only to Amazon Rekognition Custom Labels.
// Deletes an existing project policy.
//
// To get a list of project policies attached to a project, call ListProjectPolicies. To attach a
// project policy to a project, call PutProjectPolicy.
//
// This operation requires permissions to perform the
// rekognition:DeleteProjectPolicy action.
func rekognition_DeleteProjectPolicy(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.DeleteProjectPolicyInput{
		// PolicyName: *string, // Required
		// ProjectArn: *string, // Required
	}

	if len(_rekognitionPolicyName) > 0 {
		input.PolicyName = aws.String(_rekognitionPolicyName)
	}
	if len(_rekognitionProjectArn) > 0 {
		input.ProjectArn = aws.String(_rekognitionProjectArn)
	}
	if len(_rekognitionPolicyRevisionId) > 0 {
		input.PolicyRevisionId = aws.String(_rekognitionPolicyRevisionId)
	}

	if resp, err := client.DeleteProjectPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Rekognition project model or project version, like a Amazon
// Rekognition Custom Labels model or a custom adapter.
//
// You can't delete a project version if it is running or if it is training. To
// check the status of a project version, use the Status field returned from DescribeProjectVersions. To
// stop a project version call StopProjectVersion. If the project version is training, wait until it
// finishes.
//
// This operation requires permissions to perform the
// rekognition:DeleteProjectVersion action.
func rekognition_DeleteProjectVersion(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.DeleteProjectVersionInput{
		// ProjectVersionArn: *string, // Required
	}

	if len(_rekognitionProjectVersionArn) > 0 {
		input.ProjectVersionArn = aws.String(_rekognitionProjectVersionArn)
	}

	if resp, err := client.DeleteProjectVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the stream processor identified by Name . You assign the value for Name
// when you create the stream processor with CreateStreamProcessor. You might not be able to use the
// same name for a stream processor for a few seconds after calling
// DeleteStreamProcessor .
func rekognition_DeleteStreamProcessor(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.DeleteStreamProcessorInput{
		// Name: *string, // Required
	}

	if len(_rekognitionName) > 0 {
		input.Name = aws.String(_rekognitionName)
	}

	if resp, err := client.DeleteStreamProcessor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified UserID within the collection. Faces that are associated
// with the UserID are disassociated from the UserID before deleting the specified
// UserID. If the specified Collection or UserID is already deleted or not found,
// a ResourceNotFoundException will be thrown. If the action is successful with a
// 200 response, an empty HTTP body is returned.
func rekognition_DeleteUser(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.DeleteUserInput{
		// CollectionId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_rekognitionCollectionId) > 0 {
		input.CollectionId = aws.String(_rekognitionCollectionId)
	}
	if len(_rekognitionUserId) > 0 {
		input.UserId = aws.String(_rekognitionUserId)
	}
	if len(_rekognitionClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_rekognitionClientRequestToken)
	}

	if resp, err := client.DeleteUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified collection. You can use DescribeCollection to get
// information, such as the number of faces indexed into a collection and the
// version of the model used by the collection for face detection.
//
// For more information, see Describing a Collection in the Amazon Rekognition
// Developer Guide.
func rekognition_DescribeCollection(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.DescribeCollectionInput{
		// CollectionId: *string, // Required
	}

	if len(_rekognitionCollectionId) > 0 {
		input.CollectionId = aws.String(_rekognitionCollectionId)
	}

	if resp, err := client.DescribeCollection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation applies only to Amazon Rekognition Custom Labels.
// Describes an Amazon Rekognition Custom Labels dataset. You can get information
// such as the current status of a dataset and statistics about the images and
// labels in a dataset.
//
// This operation requires permissions to perform the rekognition:DescribeDataset
// action.
func rekognition_DescribeDataset(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.DescribeDatasetInput{
		// DatasetArn: *string, // Required
	}

	if len(_rekognitionDatasetArn) > 0 {
		input.DatasetArn = aws.String(_rekognitionDatasetArn)
	}

	if resp, err := client.DescribeDataset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists and describes the versions of an Amazon Rekognition project. You can
// specify up to 10 model or adapter versions in ProjectVersionArns . If you don't
// specify a value, descriptions for all model/adapter versions in the project are
// returned.
//
// This operation requires permissions to perform the
// rekognition:DescribeProjectVersions action.
func rekognition_DescribeProjectVersions(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.DescribeProjectVersionsInput{
		// ProjectArn: *string, // Required
	}

	if len(_rekognitionProjectArn) > 0 {
		input.ProjectArn = aws.String(_rekognitionProjectArn)
	}
	if len(_rekognitionMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _rekognitionMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_rekognitionNextToken) > 0 {
		input.NextToken = aws.String(_rekognitionNextToken)
	}
	if len(_rekognitionVersionNames) > 0 {
		input.VersionNames = append([]string(nil), _rekognitionVersionNames...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeProjectVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rekognition.DescribeProjectVersionsOutput
	p := rekognition.NewDescribeProjectVersionsPaginator(client, input)
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

// Gets information about your Rekognition projects.
// This operation requires permissions to perform the rekognition:DescribeProjects
// action.
func rekognition_DescribeProjects(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.DescribeProjectsInput{}

	if len(_rekognitionFeatures) > 0 {
		if err := assignInputField(input, "Features", _rekognitionFeatures); err != nil {
			log.Errorf("invalid --features: %s", err.Error())
			return
		}
	}
	if len(_rekognitionMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _rekognitionMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_rekognitionNextToken) > 0 {
		input.NextToken = aws.String(_rekognitionNextToken)
	}
	if len(_rekognitionProjectNames) > 0 {
		input.ProjectNames = append([]string(nil), _rekognitionProjectNames...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeProjects(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rekognition.DescribeProjectsOutput
	p := rekognition.NewDescribeProjectsPaginator(client, input)
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

// Provides information about a stream processor created by CreateStreamProcessor. You can get
// information about the input and output streams, the input parameters for the
// face recognition being performed, and the current status of the stream
// processor.
func rekognition_DescribeStreamProcessor(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.DescribeStreamProcessorInput{
		// Name: *string, // Required
	}

	if len(_rekognitionName) > 0 {
		input.Name = aws.String(_rekognitionName)
	}

	if resp, err := client.DescribeStreamProcessor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation applies only to Amazon Rekognition Custom Labels.
// Detects custom labels in a supplied image by using an Amazon Rekognition Custom
// Labels model.
//
// You specify which version of a model version to use by using the
// ProjectVersionArn input parameter.
//
// You pass the input image as base64-encoded image bytes or as a reference to an
// image in an Amazon S3 bucket. If you use the AWS CLI to call Amazon Rekognition
// operations, passing image bytes is not supported. The image must be either a PNG
// or JPEG formatted file.
//
// For each object that the model version detects on an image, the API returns a (
// CustomLabel ) object in an array ( CustomLabels ). Each CustomLabel object
// provides the label name ( Name ), the level of confidence that the image
// contains the object ( Confidence ), and object location information, if it
// exists, for the label on the image ( Geometry ).
//
// To filter labels that are returned, specify a value for MinConfidence .
// DetectCustomLabelsLabels only returns labels with a confidence that's higher
// than the specified value.
//
// The value of MinConfidence maps to the assumed threshold values created during
// training. For more information, see Assumed threshold in the Amazon Rekognition
// Custom Labels Developer Guide. Amazon Rekognition Custom Labels metrics
// expresses an assumed threshold as a floating point value between 0-1. The range
// of MinConfidence normalizes the threshold value to a percentage value (0-100).
// Confidence responses from DetectCustomLabels are also returned as a percentage.
// You can use MinConfidence to change the precision and recall or your model. For
// more information, see Analyzing an image in the Amazon Rekognition Custom Labels
// Developer Guide.
//
// If you don't specify a value for MinConfidence , DetectCustomLabels returns
// labels based on the assumed threshold of each label.
//
// This is a stateless API operation. That is, the operation does not persist any
// data.
//
// This operation requires permissions to perform the
// rekognition:DetectCustomLabels action.
//
// For more information, see Analyzing an image in the Amazon Rekognition Custom
// Labels Developer Guide.
func rekognition_DetectCustomLabels(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.DetectCustomLabelsInput{
		// Image: *types.Image, // Required
		// ProjectVersionArn: *string, // Required
	}

	if len(_rekognitionImage) > 0 {
		if err := assignInputField(input, "Image", _rekognitionImage); err != nil {
			log.Errorf("invalid --image: %s", err.Error())
			return
		}
	}
	if len(_rekognitionProjectVersionArn) > 0 {
		input.ProjectVersionArn = aws.String(_rekognitionProjectVersionArn)
	}
	if len(_rekognitionMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _rekognitionMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_rekognitionMinConfidence) > 0 {
		if err := assignInputField(input, "MinConfidence", _rekognitionMinConfidence); err != nil {
			log.Errorf("invalid --min-confidence: %s", err.Error())
			return
		}
	}

	if resp, err := client.DetectCustomLabels(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detects faces within an image that is provided as input.
// DetectFaces detects the 100 largest faces in the image. For each face detected,
// the operation returns face details. These details include a bounding box of the
// face, a confidence value (that the bounding box contains a face), and a fixed
// set of attributes such as facial landmarks (for example, coordinates of eye and
// mouth), pose, presence of facial occlusion, and so on.
//
// The face-detection algorithm is most effective on frontal faces. For
// non-frontal or obscured faces, the algorithm might not detect the faces or might
// detect faces with lower confidence.
//
// You pass the input image either as base64-encoded image bytes or as a reference
// to an image in an Amazon S3 bucket. If you use the AWS CLI to call Amazon
// Rekognition operations, passing image bytes is not supported. The image must be
// either a PNG or JPEG formatted file.
//
// This is a stateless API operation. That is, the operation does not persist any
// data.
//
// This operation requires permissions to perform the rekognition:DetectFaces
// action.
func rekognition_DetectFaces(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.DetectFacesInput{
		// Image: *types.Image, // Required
	}

	if len(_rekognitionImage) > 0 {
		if err := assignInputField(input, "Image", _rekognitionImage); err != nil {
			log.Errorf("invalid --image: %s", err.Error())
			return
		}
	}
	if len(_rekognitionAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _rekognitionAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}

	if resp, err := client.DetectFaces(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detects instances of real-world entities within an image (JPEG or PNG) provided
// as input. This includes objects like flower, tree, and table; events like
// wedding, graduation, and birthday party; and concepts like landscape, evening,
// and nature.
//
// For an example, see Analyzing images stored in an Amazon S3 bucket in the
// Amazon Rekognition Developer Guide.
//
// You pass the input image as base64-encoded image bytes or as a reference to an
// image in an Amazon S3 bucket. If you use the AWS CLI to call Amazon Rekognition
// operations, passing image bytes is not supported. The image must be either a PNG
// or JPEG formatted file.
//
// # Optional Parameters
//
// You can specify one or both of the GENERAL_LABELS and IMAGE_PROPERTIES feature
// types when calling the DetectLabels API. Including GENERAL_LABELS will ensure
// the response includes the labels detected in the input image, while including
// IMAGE_PROPERTIES will ensure the response includes information about the image
// quality and color.
//
// When using GENERAL_LABELS and/or IMAGE_PROPERTIES you can provide filtering
// criteria to the Settings parameter. You can filter with sets of individual
// labels or with label categories. You can specify inclusive filters, exclusive
// filters, or a combination of inclusive and exclusive filters. For more
// information on filtering see [Detecting Labels in an Image].
//
// When getting labels, you can specify MinConfidence to control the confidence
// threshold for the labels returned. The default is 55%. You can also add the
// MaxLabels parameter to limit the number of labels returned. The default and
// upper limit is 1000 labels. These arguments are only valid when supplying
// GENERAL_LABELS as a feature type.
//
// # Response Elements
//
// For each object, scene, and concept the API returns one or more labels. The API
// returns the following types of information about labels:
//
// - Name - The name of the detected label.
//
// - Confidence - The level of confidence in the label assigned to a detected
// object.
//
// - Parents - The ancestor labels for a detected label. DetectLabels returns a
// hierarchical taxonomy of detected labels. For example, a detected car might be
// assigned the label car. The label car has two parent labels: Vehicle (its
// parent) and Transportation (its grandparent). The response includes the all
// ancestors for a label, where every ancestor is a unique label. In the previous
// example, Car, Vehicle, and Transportation are returned as unique labels in the
// response.
//
// - Aliases - Possible Aliases for the label.
//
// - Categories - The label categories that the detected label belongs to.
//
// - BoundingBox — Bounding boxes are described for all instances of detected
// common object labels, returned in an array of Instance objects. An Instance
// object contains a BoundingBox object, describing the location of the label on
// the input image. It also includes the confidence for the accuracy of the
// detected bounding box.
//
// The API returns the following information regarding the image, as part of the
// ImageProperties structure:
//
// - Quality - Information about the Sharpness, Brightness, and Contrast of the
// input image, scored between 0 to 100. Image quality is returned for the entire
// image, as well as the background and the foreground.
//
// - Dominant Color - An array of the dominant colors in the image.
//
// - Foreground - Information about the sharpness, brightness, and dominant
// colors of the input image’s foreground.
//
// - Background - Information about the sharpness, brightness, and dominant
// colors of the input image’s background.
//
// The list of returned labels will include at least one label for every detected
// object, along with information about that label. In the following example,
// suppose the input image has a lighthouse, the sea, and a rock. The response
// includes all three labels, one for each object, as well as the confidence in the
// label:
//
// {Name: lighthouse, Confidence: 98.4629}
//
// {Name: rock,Confidence: 79.2097}
//
// {Name: sea,Confidence: 75.061}
//
// The list of labels can include multiple labels for the same object. For
// example, if the input image shows a flower (for example, a tulip), the operation
// might return the following three labels.
//
// {Name: flower,Confidence: 99.0562}
//
// {Name: plant,Confidence: 99.0562}
//
// {Name: tulip,Confidence: 99.0562}
//
// In this example, the detection algorithm more precisely identifies the flower
// as a tulip.
//
// If the object detected is a person, the operation doesn't provide the same
// facial details that the DetectFacesoperation provides.
//
// This is a stateless API operation that doesn't return any data.
//
// This operation requires permissions to perform the rekognition:DetectLabels
// action.
//
// [Detecting Labels in an Image]: https://docs.aws.amazon.com/rekognition/latest/dg/labels-detect-labels-image.html
func rekognition_DetectLabels(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.DetectLabelsInput{
		// Image: *types.Image, // Required
	}

	if len(_rekognitionImage) > 0 {
		if err := assignInputField(input, "Image", _rekognitionImage); err != nil {
			log.Errorf("invalid --image: %s", err.Error())
			return
		}
	}
	if len(_rekognitionFeatures) > 0 {
		if err := assignInputField(input, "Features", _rekognitionFeatures); err != nil {
			log.Errorf("invalid --features: %s", err.Error())
			return
		}
	}
	if len(_rekognitionMaxLabels) > 0 {
		if err := assignInputField(input, "MaxLabels", _rekognitionMaxLabels); err != nil {
			log.Errorf("invalid --max-labels: %s", err.Error())
			return
		}
	}
	if len(_rekognitionMinConfidence) > 0 {
		if err := assignInputField(input, "MinConfidence", _rekognitionMinConfidence); err != nil {
			log.Errorf("invalid --min-confidence: %s", err.Error())
			return
		}
	}
	if len(_rekognitionSettings) > 0 {
		if err := assignInputField(input, "Settings", _rekognitionSettings); err != nil {
			log.Errorf("invalid --settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.DetectLabels(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detects unsafe content in a specified JPEG or PNG format image. Use
// DetectModerationLabels to moderate images depending on your requirements. For
// example, you might want to filter images that contain nudity, but not images
// containing suggestive content.
//
// To filter images, use the labels returned by DetectModerationLabels to
// determine which types of content are appropriate.
//
// For information about moderation labels, see Detecting Unsafe Content in the
// Amazon Rekognition Developer Guide.
//
// You pass the input image either as base64-encoded image bytes or as a reference
// to an image in an Amazon S3 bucket. If you use the AWS CLI to call Amazon
// Rekognition operations, passing image bytes is not supported. The image must be
// either a PNG or JPEG formatted file.
//
// You can specify an adapter to use when retrieving label predictions by
// providing a ProjectVersionArn to the ProjectVersion argument.
func rekognition_DetectModerationLabels(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.DetectModerationLabelsInput{
		// Image: *types.Image, // Required
	}

	if len(_rekognitionImage) > 0 {
		if err := assignInputField(input, "Image", _rekognitionImage); err != nil {
			log.Errorf("invalid --image: %s", err.Error())
			return
		}
	}
	if len(_rekognitionHumanLoopConfig) > 0 {
		if err := assignInputField(input, "HumanLoopConfig", _rekognitionHumanLoopConfig); err != nil {
			log.Errorf("invalid --human-loop-config: %s", err.Error())
			return
		}
	}
	if len(_rekognitionMinConfidence) > 0 {
		if err := assignInputField(input, "MinConfidence", _rekognitionMinConfidence); err != nil {
			log.Errorf("invalid --min-confidence: %s", err.Error())
			return
		}
	}
	if len(_rekognitionProjectVersion) > 0 {
		input.ProjectVersion = aws.String(_rekognitionProjectVersion)
	}

	if resp, err := client.DetectModerationLabels(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detects Personal Protective Equipment (PPE) worn by people detected in an
// image. Amazon Rekognition can detect the following types of PPE.
//
// - Face cover
//
// - Hand cover
//
// - Head cover
//
// You pass the input image as base64-encoded image bytes or as a reference to an
// image in an Amazon S3 bucket. The image must be either a PNG or JPG formatted
// file.
//
// DetectProtectiveEquipment detects PPE worn by up to 15 persons detected in an
// image.
//
// For each person detected in the image the API returns an array of body parts
// (face, head, left-hand, right-hand). For each body part, an array of detected
// items of PPE is returned, including an indicator of whether or not the PPE
// covers the body part. The API returns the confidence it has in each detection
// (person, PPE, body part and body part coverage). It also returns a bounding box
// (BoundingBox ) for each detected person and each detected item of PPE.
//
// You can optionally request a summary of detected PPE items with the
// SummarizationAttributes input parameter. The summary provides the following
// information.
//
// - The persons detected as wearing all of the types of PPE that you specify.
//
// - The persons detected as not wearing all of the types PPE that you specify.
//
// - The persons detected where PPE adornment could not be determined.
//
// This is a stateless API operation. That is, the operation does not persist any
// data.
//
// This operation requires permissions to perform the
// rekognition:DetectProtectiveEquipment action.
func rekognition_DetectProtectiveEquipment(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.DetectProtectiveEquipmentInput{
		// Image: *types.Image, // Required
	}

	if len(_rekognitionImage) > 0 {
		if err := assignInputField(input, "Image", _rekognitionImage); err != nil {
			log.Errorf("invalid --image: %s", err.Error())
			return
		}
	}
	if len(_rekognitionSummarizationAttributes) > 0 {
		if err := assignInputField(input, "SummarizationAttributes", _rekognitionSummarizationAttributes); err != nil {
			log.Errorf("invalid --summarization-attributes: %s", err.Error())
			return
		}
	}

	if resp, err := client.DetectProtectiveEquipment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detects text in the input image and converts it into machine-readable text.
// Pass the input image as base64-encoded image bytes or as a reference to an
// image in an Amazon S3 bucket. If you use the AWS CLI to call Amazon Rekognition
// operations, you must pass it as a reference to an image in an Amazon S3 bucket.
// For the AWS CLI, passing image bytes is not supported. The image must be either
// a .png or .jpeg formatted file.
//
// The DetectText operation returns text in an array of TextDetection elements, TextDetections .
// Each TextDetection element provides information about a single word or line of
// text that was detected in the image.
//
// A word is one or more script characters that are not separated by spaces.
// DetectText can detect up to 100 words in an image.
//
// A line is a string of equally spaced words. A line isn't necessarily a complete
// sentence. For example, a driver's license number is detected as a line. A line
// ends when there is no aligned text after it. Also, a line ends when there is a
// large gap between words, relative to the length of the words. This means,
// depending on the gap between words, Amazon Rekognition may detect multiple lines
// in text aligned in the same direction. Periods don't represent the end of a
// line. If a sentence spans multiple lines, the DetectText operation returns
// multiple lines.
//
// To determine whether a TextDetection element is a line of text or a word, use
// the TextDetection object Type field.
//
// To be detected, text must be within +/- 90 degrees orientation of the
// horizontal axis.
//
// For more information, see Detecting text in the Amazon Rekognition Developer
// Guide.
func rekognition_DetectText(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.DetectTextInput{
		// Image: *types.Image, // Required
	}

	if len(_rekognitionImage) > 0 {
		if err := assignInputField(input, "Image", _rekognitionImage); err != nil {
			log.Errorf("invalid --image: %s", err.Error())
			return
		}
	}
	if len(_rekognitionFilters) > 0 {
		if err := assignInputField(input, "Filters", _rekognitionFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}

	if resp, err := client.DetectText(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the association between a Face supplied in an array of FaceIds and the
// User. If the User is not present already, then a ResourceNotFound exception is
// thrown. If successful, an array of faces that are disassociated from the User is
// returned. If a given face is already disassociated from the given UserID, it
// will be ignored and not be returned in the response. If a given face is already
// associated with a different User or not found in the collection it will be
// returned as part of UnsuccessfulDisassociations . You can remove 1 - 100 face
// IDs from a user at one time.
func rekognition_DisassociateFaces(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.DisassociateFacesInput{
		// CollectionId: *string, // Required
		// FaceIds: []string, // Required
		// UserId: *string, // Required
	}

	if len(_rekognitionCollectionId) > 0 {
		input.CollectionId = aws.String(_rekognitionCollectionId)
	}
	if len(_rekognitionFaceIds) > 0 {
		input.FaceIds = append([]string(nil), _rekognitionFaceIds...)
	}
	if len(_rekognitionUserId) > 0 {
		input.UserId = aws.String(_rekognitionUserId)
	}
	if len(_rekognitionClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_rekognitionClientRequestToken)
	}

	if resp, err := client.DisassociateFaces(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation applies only to Amazon Rekognition Custom Labels.
// Distributes the entries (images) in a training dataset across the training
// dataset and the test dataset for a project. DistributeDatasetEntries moves 20%
// of the training dataset images to the test dataset. An entry is a JSON Line that
// describes an image.
//
// You supply the Amazon Resource Names (ARN) of a project's training dataset and
// test dataset. The training dataset must contain the images that you want to
// split. The test dataset must be empty. The datasets must belong to the same
// project. To create training and test datasets for a project, call CreateDataset.
//
// Distributing a dataset takes a while to complete. To check the status call
// DescribeDataset . The operation is complete when the Status field for the
// training dataset and the test dataset is UPDATE_COMPLETE . If the dataset split
// fails, the value of Status is UPDATE_FAILED .
//
// This operation requires permissions to perform the
// rekognition:DistributeDatasetEntries action.
func rekognition_DistributeDatasetEntries(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.DistributeDatasetEntriesInput{
		// Datasets: []types.DistributeDataset, // Required
	}

	if len(_rekognitionDatasets) > 0 {
		if err := assignInputField(input, "Datasets", _rekognitionDatasets); err != nil {
			log.Errorf("invalid --datasets: %s", err.Error())
			return
		}
	}

	if resp, err := client.DistributeDatasetEntries(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the name and additional information about a celebrity based on their
// Amazon Rekognition ID. The additional information is returned as an array of
// URLs. If there is no additional information about the celebrity, this list is
// empty.
//
// For more information, see Getting information about a celebrity in the Amazon
// Rekognition Developer Guide.
//
// This operation requires permissions to perform the rekognition:GetCelebrityInfo
// action.
func rekognition_GetCelebrityInfo(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.GetCelebrityInfoInput{
		// Id: *string, // Required
	}

	if len(_rekognitionId) > 0 {
		input.Id = aws.String(_rekognitionId)
	}

	if resp, err := client.GetCelebrityInfo(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the celebrity recognition results for a Amazon Rekognition Video analysis
// started by StartCelebrityRecognition.
//
// Celebrity recognition in a video is an asynchronous operation. Analysis is
// started by a call to StartCelebrityRecognitionwhich returns a job identifier ( JobId ).
//
// When the celebrity recognition operation finishes, Amazon Rekognition Video
// publishes a completion status to the Amazon Simple Notification Service topic
// registered in the initial call to StartCelebrityRecognition . To get the results
// of the celebrity recognition analysis, first check that the status value
// published to the Amazon SNS topic is SUCCEEDED . If so, call
// GetCelebrityDetection and pass the job identifier ( JobId ) from the initial
// call to StartCelebrityDetection .
//
// For more information, see Working With Stored Videos in the Amazon Rekognition
// Developer Guide.
//
// GetCelebrityRecognition returns detected celebrities and the time(s) they are
// detected in an array ( Celebrities ) of CelebrityRecognition objects. Each CelebrityRecognition
// contains information about the celebrity in a CelebrityDetailobject and the time, Timestamp ,
// the celebrity was detected. This CelebrityDetailobject stores information about the detected
// celebrity's face attributes, a face bounding box, known gender, the celebrity's
// name, and a confidence estimate.
//
// GetCelebrityRecognition only returns the default facial attributes ( BoundingBox
// , Confidence , Landmarks , Pose , and Quality ). The BoundingBox field only
// applies to the detected face instance. The other facial attributes listed in the
// Face object of the following response syntax are not returned. For more
// information, see FaceDetail in the Amazon Rekognition Developer Guide.
//
// By default, the Celebrities array is sorted by time (milliseconds from the
// start of the video). You can also sort the array by celebrity by specifying the
// value ID in the SortBy input parameter.
//
// The CelebrityDetail object includes the celebrity identifer and additional
// information urls. If you don't store the additional information urls, you can
// get them later by calling GetCelebrityInfowith the celebrity identifer.
//
// No information is returned for faces not recognized as celebrities.
//
// Use MaxResults parameter to limit the number of labels returned. If there are
// more results than specified in MaxResults , the value of NextToken in the
// operation response contains a pagination token for getting the next set of
// results. To get the next page of results, call GetCelebrityDetection and
// populate the NextToken request parameter with the token value returned from the
// previous call to GetCelebrityRecognition .
func rekognition_GetCelebrityRecognition(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.GetCelebrityRecognitionInput{
		// JobId: *string, // Required
	}

	if len(_rekognitionJobId) > 0 {
		input.JobId = aws.String(_rekognitionJobId)
	}
	if len(_rekognitionMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _rekognitionMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_rekognitionNextToken) > 0 {
		input.NextToken = aws.String(_rekognitionNextToken)
	}
	if len(_rekognitionSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _rekognitionSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetCelebrityRecognition(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rekognition.GetCelebrityRecognitionOutput
	p := rekognition.NewGetCelebrityRecognitionPaginator(client, input)
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

// Gets the inappropriate, unwanted, or offensive content analysis results for a
// Amazon Rekognition Video analysis started by StartContentModeration. For a list of moderation labels
// in Amazon Rekognition, see [Using the image and video moderation APIs].
//
// Amazon Rekognition Video inappropriate or offensive content detection in a
// stored video is an asynchronous operation. You start analysis by calling StartContentModerationwhich
// returns a job identifier ( JobId ). When analysis finishes, Amazon Rekognition
// Video publishes a completion status to the Amazon Simple Notification Service
// topic registered in the initial call to StartContentModeration . To get the
// results of the content analysis, first check that the status value published to
// the Amazon SNS topic is SUCCEEDED . If so, call GetContentModeration and pass
// the job identifier ( JobId ) from the initial call to StartContentModeration .
//
// For more information, see Working with Stored Videos in the Amazon Rekognition
// Devlopers Guide.
//
// GetContentModeration returns detected inappropriate, unwanted, or offensive
// content moderation labels, and the time they are detected, in an array,
// ModerationLabels , of ContentModerationDetection objects.
//
// By default, the moderated labels are returned sorted by time, in milliseconds
// from the start of the video. You can also sort them by moderated label by
// specifying NAME for the SortBy input parameter.
//
// Since video analysis can return a large number of results, use the MaxResults
// parameter to limit the number of labels returned in a single call to
// GetContentModeration . If there are more results than specified in MaxResults ,
// the value of NextToken in the operation response contains a pagination token
// for getting the next set of results. To get the next page of results, call
// GetContentModeration and populate the NextToken request parameter with the
// value of NextToken returned from the previous call to GetContentModeration .
//
// For more information, see moderating content in the Amazon Rekognition
// Developer Guide.
//
// [Using the image and video moderation APIs]: https://docs.aws.amazon.com/rekognition/latest/dg/moderation.html#moderation-api
func rekognition_GetContentModeration(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.GetContentModerationInput{
		// JobId: *string, // Required
	}

	if len(_rekognitionJobId) > 0 {
		input.JobId = aws.String(_rekognitionJobId)
	}
	if len(_rekognitionAggregateBy) > 0 {
		if err := assignInputField(input, "AggregateBy", _rekognitionAggregateBy); err != nil {
			log.Errorf("invalid --aggregate-by: %s", err.Error())
			return
		}
	}
	if len(_rekognitionMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _rekognitionMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_rekognitionNextToken) > 0 {
		input.NextToken = aws.String(_rekognitionNextToken)
	}
	if len(_rekognitionSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _rekognitionSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetContentModeration(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rekognition.GetContentModerationOutput
	p := rekognition.NewGetContentModerationPaginator(client, input)
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

// Gets face detection results for a Amazon Rekognition Video analysis started by StartFaceDetection.
// Face detection with Amazon Rekognition Video is an asynchronous operation. You
// start face detection by calling StartFaceDetectionwhich returns a job identifier ( JobId ). When
// the face detection operation finishes, Amazon Rekognition Video publishes a
// completion status to the Amazon Simple Notification Service topic registered in
// the initial call to StartFaceDetection . To get the results of the face
// detection operation, first check that the status value published to the Amazon
// SNS topic is SUCCEEDED . If so, call GetFaceDetection and pass the job identifier ( JobId ) from
// the initial call to StartFaceDetection .
//
// GetFaceDetection returns an array of detected faces ( Faces ) sorted by the time
// the faces were detected.
//
// Use MaxResults parameter to limit the number of labels returned. If there are
// more results than specified in MaxResults , the value of NextToken in the
// operation response contains a pagination token for getting the next set of
// results. To get the next page of results, call GetFaceDetection and populate
// the NextToken request parameter with the token value returned from the previous
// call to GetFaceDetection .
//
// Note that for the GetFaceDetection operation, the returned values for
// FaceOccluded and EyeDirection will always be "null".
func rekognition_GetFaceDetection(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.GetFaceDetectionInput{
		// JobId: *string, // Required
	}

	if len(_rekognitionJobId) > 0 {
		input.JobId = aws.String(_rekognitionJobId)
	}
	if len(_rekognitionMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _rekognitionMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_rekognitionNextToken) > 0 {
		input.NextToken = aws.String(_rekognitionNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetFaceDetection(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rekognition.GetFaceDetectionOutput
	p := rekognition.NewGetFaceDetectionPaginator(client, input)
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

// Retrieves the results of a specific Face Liveness session. It requires the
// sessionId as input, which was created using CreateFaceLivenessSession . Returns
// the corresponding Face Liveness confidence score, a reference image that
// includes a face bounding box, and audit images that also contain face bounding
// boxes. The Face Liveness confidence score ranges from 0 to 100.
//
// The number of audit images returned by GetFaceLivenessSessionResults is defined
// by the AuditImagesLimit paramater when calling CreateFaceLivenessSession .
// Reference images are always returned when possible.
func rekognition_GetFaceLivenessSessionResults(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.GetFaceLivenessSessionResultsInput{
		// SessionId: *string, // Required
	}

	if len(_rekognitionSessionId) > 0 {
		input.SessionId = aws.String(_rekognitionSessionId)
	}

	if resp, err := client.GetFaceLivenessSessionResults(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the face search results for Amazon Rekognition Video face search started
// by StartFaceSearch. The search returns faces in a collection that match the faces of persons
// detected in a video. It also includes the time(s) that faces are matched in the
// video.
//
// Face search in a video is an asynchronous operation. You start face search by
// calling to StartFaceSearchwhich returns a job identifier ( JobId ). When the search operation
// finishes, Amazon Rekognition Video publishes a completion status to the Amazon
// Simple Notification Service topic registered in the initial call to
// StartFaceSearch . To get the search results, first check that the status value
// published to the Amazon SNS topic is SUCCEEDED . If so, call GetFaceSearch and
// pass the job identifier ( JobId ) from the initial call to StartFaceSearch .
//
// For more information, see Searching Faces in a Collection in the Amazon
// Rekognition Developer Guide.
//
// The search results are retured in an array, Persons , of PersonMatch objects. Each
// PersonMatch element contains details about the matching faces in the input
// collection, person information (facial attributes, bounding boxes, and person
// identifer) for the matched person, and the time the person was matched in the
// video.
//
// GetFaceSearch only returns the default facial attributes ( BoundingBox ,
// Confidence , Landmarks , Pose , and Quality ). The other facial attributes
// listed in the Face object of the following response syntax are not returned.
// For more information, see FaceDetail in the Amazon Rekognition Developer Guide.
//
// By default, the Persons array is sorted by the time, in milliseconds from the
// start of the video, persons are matched. You can also sort by persons by
// specifying INDEX for the SORTBY input parameter.
func rekognition_GetFaceSearch(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.GetFaceSearchInput{
		// JobId: *string, // Required
	}

	if len(_rekognitionJobId) > 0 {
		input.JobId = aws.String(_rekognitionJobId)
	}
	if len(_rekognitionMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _rekognitionMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_rekognitionNextToken) > 0 {
		input.NextToken = aws.String(_rekognitionNextToken)
	}
	if len(_rekognitionSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _rekognitionSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetFaceSearch(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rekognition.GetFaceSearchOutput
	p := rekognition.NewGetFaceSearchPaginator(client, input)
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

// Gets the label detection results of a Amazon Rekognition Video analysis started
// by StartLabelDetection.
//
// The label detection operation is started by a call to StartLabelDetection which returns a job
// identifier ( JobId ). When the label detection operation finishes, Amazon
// Rekognition publishes a completion status to the Amazon Simple Notification
// Service topic registered in the initial call to StartlabelDetection .
//
// To get the results of the label detection operation, first check that the
// status value published to the Amazon SNS topic is SUCCEEDED . If so, call GetLabelDetection and
// pass the job identifier ( JobId ) from the initial call to StartLabelDetection .
//
// GetLabelDetection returns an array of detected labels ( Labels ) sorted by the
// time the labels were detected. You can also sort by the label name by specifying
// NAME for the SortBy input parameter. If there is no NAME specified, the default
// sort is by timestamp.
//
// You can select how results are aggregated by using the AggregateBy input
// parameter. The default aggregation method is TIMESTAMPS . You can also aggregate
// by SEGMENTS , which aggregates all instances of labels detected in a given
// segment.
//
// The returned Labels array may include the following attributes:
//
// - Name - The name of the detected label.
//
// - Confidence - The level of confidence in the label assigned to a detected
// object.
//
// - Parents - The ancestor labels for a detected label. GetLabelDetection
// returns a hierarchical taxonomy of detected labels. For example, a detected car
// might be assigned the label car. The label car has two parent labels: Vehicle
// (its parent) and Transportation (its grandparent). The response includes the all
// ancestors for a label, where every ancestor is a unique label. In the previous
// example, Car, Vehicle, and Transportation are returned as unique labels in the
// response.
//
// - Aliases - Possible Aliases for the label.
//
// - Categories - The label categories that the detected label belongs to.
//
// - BoundingBox — Bounding boxes are described for all instances of detected
// common object labels, returned in an array of Instance objects. An Instance
// object contains a BoundingBox object, describing the location of the label on
// the input image. It also includes the confidence for the accuracy of the
// detected bounding box.
//
// - Timestamp - Time, in milliseconds from the start of the video, that the
// label was detected. For aggregation by SEGMENTS , the StartTimestampMillis ,
// EndTimestampMillis , and DurationMillis structures are what define a segment.
// Although the “Timestamp” structure is still returned with each label, its value
// is set to be the same as StartTimestampMillis .
//
// Timestamp and Bounding box information are returned for detected Instances,
// only if aggregation is done by TIMESTAMPS . If aggregating by SEGMENTS ,
// information about detected instances isn’t returned.
//
// The version of the label model used for the detection is also returned.
//
// Note DominantColors isn't returned for Instances , although it is shown as part
// of the response in the sample seen below.
//
// Use MaxResults parameter to limit the number of labels returned. If there are
// more results than specified in MaxResults , the value of NextToken in the
// operation response contains a pagination token for getting the next set of
// results. To get the next page of results, call GetlabelDetection and populate
// the NextToken request parameter with the token value returned from the previous
// call to GetLabelDetection .
//
// If you are retrieving results while using the Amazon Simple Notification
// Service, note that you will receive an "ERROR" notification if the job
// encounters an issue.
func rekognition_GetLabelDetection(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.GetLabelDetectionInput{
		// JobId: *string, // Required
	}

	if len(_rekognitionJobId) > 0 {
		input.JobId = aws.String(_rekognitionJobId)
	}
	if len(_rekognitionAggregateBy) > 0 {
		if err := assignInputField(input, "AggregateBy", _rekognitionAggregateBy); err != nil {
			log.Errorf("invalid --aggregate-by: %s", err.Error())
			return
		}
	}
	if len(_rekognitionMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _rekognitionMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_rekognitionNextToken) > 0 {
		input.NextToken = aws.String(_rekognitionNextToken)
	}
	if len(_rekognitionSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _rekognitionSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetLabelDetection(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rekognition.GetLabelDetectionOutput
	p := rekognition.NewGetLabelDetectionPaginator(client, input)
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

// Retrieves the results for a given media analysis job. Takes a JobId returned by
// StartMediaAnalysisJob.
func rekognition_GetMediaAnalysisJob(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.GetMediaAnalysisJobInput{
		// JobId: *string, // Required
	}

	if len(_rekognitionJobId) > 0 {
		input.JobId = aws.String(_rekognitionJobId)
	}

	if resp, err := client.GetMediaAnalysisJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// End of support notice: On October 31, 2025, AWS will discontinue support for
// Amazon Rekognition People Pathing. After October 31, 2025, you will no longer be
// able to use the Rekognition People Pathing capability. For more information,
// visit this [blog post].
//
// Gets the path tracking results of a Amazon Rekognition Video analysis started
// by StartPersonTracking.
//
// The person path tracking operation is started by a call to StartPersonTracking
// which returns a job identifier ( JobId ). When the operation finishes, Amazon
// Rekognition Video publishes a completion status to the Amazon Simple
// Notification Service topic registered in the initial call to StartPersonTracking
// .
//
// To get the results of the person path tracking operation, first check that the
// status value published to the Amazon SNS topic is SUCCEEDED . If so, call GetPersonTracking and
// pass the job identifier ( JobId ) from the initial call to StartPersonTracking .
//
// GetPersonTracking returns an array, Persons , of tracked persons and the time(s)
// their paths were tracked in the video.
//
// GetPersonTracking only returns the default facial attributes ( BoundingBox ,
// Confidence , Landmarks , Pose , and Quality ). The other facial attributes
// listed in the Face object of the following response syntax are not returned.
//
// For more information, see FaceDetail in the Amazon Rekognition Developer Guide.
//
// By default, the array is sorted by the time(s) a person's path is tracked in
// the video. You can sort by tracked persons by specifying INDEX for the SortBy
// input parameter.
//
// Use the MaxResults parameter to limit the number of items returned. If there
// are more results than specified in MaxResults , the value of NextToken in the
// operation response contains a pagination token for getting the next set of
// results. To get the next page of results, call GetPersonTracking and populate
// the NextToken request parameter with the token value returned from the previous
// call to GetPersonTracking .
//
// [blog post]: https://aws.amazon.com/blogs/machine-learning/transitioning-from-amazon-rekognition-people-pathing-exploring-other-alternatives/
func rekognition_GetPersonTracking(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.GetPersonTrackingInput{
		// JobId: *string, // Required
	}

	if len(_rekognitionJobId) > 0 {
		input.JobId = aws.String(_rekognitionJobId)
	}
	if len(_rekognitionMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _rekognitionMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_rekognitionNextToken) > 0 {
		input.NextToken = aws.String(_rekognitionNextToken)
	}
	if len(_rekognitionSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _rekognitionSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetPersonTracking(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rekognition.GetPersonTrackingOutput
	p := rekognition.NewGetPersonTrackingPaginator(client, input)
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

// Gets the segment detection results of a Amazon Rekognition Video analysis
// started by StartSegmentDetection.
//
// Segment detection with Amazon Rekognition Video is an asynchronous operation.
// You start segment detection by calling StartSegmentDetectionwhich returns a job identifier ( JobId ).
// When the segment detection operation finishes, Amazon Rekognition publishes a
// completion status to the Amazon Simple Notification Service topic registered in
// the initial call to StartSegmentDetection . To get the results of the segment
// detection operation, first check that the status value published to the Amazon
// SNS topic is SUCCEEDED . if so, call GetSegmentDetection and pass the job
// identifier ( JobId ) from the initial call of StartSegmentDetection .
//
// GetSegmentDetection returns detected segments in an array ( Segments ) of SegmentDetection
// objects. Segments is sorted by the segment types specified in the SegmentTypes
// input parameter of StartSegmentDetection . Each element of the array includes
// the detected segment, the precentage confidence in the acuracy of the detected
// segment, the type of the segment, and the frame in which the segment was
// detected.
//
// Use SelectedSegmentTypes to find out the type of segment detection requested in
// the call to StartSegmentDetection .
//
// Use the MaxResults parameter to limit the number of segment detections
// returned. If there are more results than specified in MaxResults , the value of
// NextToken in the operation response contains a pagination token for getting the
// next set of results. To get the next page of results, call GetSegmentDetection
// and populate the NextToken request parameter with the token value returned from
// the previous call to GetSegmentDetection .
//
// For more information, see Detecting video segments in stored video in the
// Amazon Rekognition Developer Guide.
func rekognition_GetSegmentDetection(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.GetSegmentDetectionInput{
		// JobId: *string, // Required
	}

	if len(_rekognitionJobId) > 0 {
		input.JobId = aws.String(_rekognitionJobId)
	}
	if len(_rekognitionMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _rekognitionMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_rekognitionNextToken) > 0 {
		input.NextToken = aws.String(_rekognitionNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetSegmentDetection(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rekognition.GetSegmentDetectionOutput
	p := rekognition.NewGetSegmentDetectionPaginator(client, input)
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

// Gets the text detection results of a Amazon Rekognition Video analysis started
// by StartTextDetection.
//
// Text detection with Amazon Rekognition Video is an asynchronous operation. You
// start text detection by calling StartTextDetectionwhich returns a job identifier ( JobId ) When
// the text detection operation finishes, Amazon Rekognition publishes a completion
// status to the Amazon Simple Notification Service topic registered in the initial
// call to StartTextDetection . To get the results of the text detection operation,
// first check that the status value published to the Amazon SNS topic is SUCCEEDED
// . if so, call GetTextDetection and pass the job identifier ( JobId ) from the
// initial call of StartLabelDetection .
//
// GetTextDetection returns an array of detected text ( TextDetections ) sorted by
// the time the text was detected, up to 100 words per frame of video.
//
// Each element of the array includes the detected text, the precentage confidence
// in the acuracy of the detected text, the time the text was detected, bounding
// box information for where the text was located, and unique identifiers for words
// and their lines.
//
// Use MaxResults parameter to limit the number of text detections returned. If
// there are more results than specified in MaxResults , the value of NextToken in
// the operation response contains a pagination token for getting the next set of
// results. To get the next page of results, call GetTextDetection and populate
// the NextToken request parameter with the token value returned from the previous
// call to GetTextDetection .
func rekognition_GetTextDetection(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.GetTextDetectionInput{
		// JobId: *string, // Required
	}

	if len(_rekognitionJobId) > 0 {
		input.JobId = aws.String(_rekognitionJobId)
	}
	if len(_rekognitionMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _rekognitionMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_rekognitionNextToken) > 0 {
		input.NextToken = aws.String(_rekognitionNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetTextDetection(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rekognition.GetTextDetectionOutput
	p := rekognition.NewGetTextDetectionPaginator(client, input)
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

// Detects faces in the input image and adds them to the specified collection.
// Amazon Rekognition doesn't save the actual faces that are detected. Instead,
// the underlying detection algorithm first detects the faces in the input image.
// For each face, the algorithm extracts facial features into a feature vector, and
// stores it in the backend database. Amazon Rekognition uses feature vectors when
// it performs face match and search operations using the SearchFacesand SearchFacesByImage operations.
//
// For more information, see Adding faces to a collection in the Amazon
// Rekognition Developer Guide.
//
// To get the number of faces in a collection, call DescribeCollection.
//
// If you're using version 1.0 of the face detection model, IndexFaces indexes the
// 15 largest faces in the input image. Later versions of the face detection model
// index the 100 largest faces in the input image.
//
// If you're using version 4 or later of the face model, image orientation
// information is not returned in the OrientationCorrection field.
//
// To determine which version of the model you're using, call DescribeCollection and supply the
// collection ID. You can also get the model version from the value of
// FaceModelVersion in the response from IndexFaces
//
// For more information, see Model Versioning in the Amazon Rekognition Developer
// Guide.
//
// If you provide the optional ExternalImageId for the input image you provided,
// Amazon Rekognition associates this ID with all faces that it detects. When you
// call the ListFacesoperation, the response returns the external ID. You can use this
// external image ID to create a client-side index to associate the faces with each
// image. You can then use the index to find all faces in an image.
//
// You can specify the maximum number of faces to index with the MaxFaces input
// parameter. This is useful when you want to index the largest faces in an image
// and don't want to index smaller faces, such as those belonging to people
// standing in the background.
//
// The QualityFilter input parameter allows you to filter out detected faces that
// don’t meet a required quality bar. The quality bar is based on a variety of
// common use cases. By default, IndexFaces chooses the quality bar that's used to
// filter faces. You can also explicitly choose the quality bar. Use QualityFilter
// , to set the quality bar by specifying LOW , MEDIUM , or HIGH . If you do not
// want to filter detected faces, specify NONE .
//
// To use quality filtering, you need a collection associated with version 3 of
// the face model or higher. To get the version of the face model associated with a
// collection, call DescribeCollection.
//
// Information about faces detected in an image, but not indexed, is returned in
// an array of UnindexedFaceobjects, UnindexedFaces . Faces aren't indexed for reasons such as:
//
// - The number of faces detected exceeds the value of the MaxFaces request
// parameter.
//
// - The face is too small compared to the image dimensions.
//
// - The face is too blurry.
//
// - The image is too dark.
//
// - The face has an extreme pose.
//
// - The face doesn’t have enough detail to be suitable for face search.
//
// In response, the IndexFaces operation returns an array of metadata for all
// detected faces, FaceRecords . This includes:
//
// - The bounding box, BoundingBox , of the detected face.
//
// - A confidence value, Confidence , which indicates the confidence that the
// bounding box contains a face.
//
// - A face ID, FaceId , assigned by the service for each face that's detected
// and stored.
//
// - An image ID, ImageId , assigned by the service for the input image.
//
// If you request ALL or specific facial attributes (e.g., FACE_OCCLUDED ) by using
// the detectionAttributes parameter, Amazon Rekognition returns detailed facial
// attributes, such as facial landmarks (for example, location of eye and mouth),
// facial occlusion, and other facial attributes.
//
// If you provide the same image, specify the same collection, and use the same
// external ID in the IndexFaces operation, Amazon Rekognition doesn't save
// duplicate face metadata.
//
// The input image is passed either as base64-encoded image bytes, or as a
// reference to an image in an Amazon S3 bucket. If you use the AWS CLI to call
// Amazon Rekognition operations, passing image bytes isn't supported. The image
// must be formatted as a PNG or JPEG file.
//
// This operation requires permissions to perform the rekognition:IndexFaces
// action.
func rekognition_IndexFaces(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.IndexFacesInput{
		// CollectionId: *string, // Required
		// Image: *types.Image, // Required
	}

	if len(_rekognitionCollectionId) > 0 {
		input.CollectionId = aws.String(_rekognitionCollectionId)
	}
	if len(_rekognitionImage) > 0 {
		if err := assignInputField(input, "Image", _rekognitionImage); err != nil {
			log.Errorf("invalid --image: %s", err.Error())
			return
		}
	}
	if len(_rekognitionDetectionAttributes) > 0 {
		if err := assignInputField(input, "DetectionAttributes", _rekognitionDetectionAttributes); err != nil {
			log.Errorf("invalid --detection-attributes: %s", err.Error())
			return
		}
	}
	if len(_rekognitionExternalImageId) > 0 {
		input.ExternalImageId = aws.String(_rekognitionExternalImageId)
	}
	if len(_rekognitionMaxFaces) > 0 {
		if err := assignInputField(input, "MaxFaces", _rekognitionMaxFaces); err != nil {
			log.Errorf("invalid --max-faces: %s", err.Error())
			return
		}
	}
	if len(_rekognitionQualityFilter) > 0 {
		if err := assignInputField(input, "QualityFilter", _rekognitionQualityFilter); err != nil {
			log.Errorf("invalid --quality-filter: %s", err.Error())
			return
		}
	}

	if resp, err := client.IndexFaces(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns list of collection IDs in your account. If the result is truncated, the
// response also provides a NextToken that you can use in the subsequent request
// to fetch the next set of collection IDs.
//
// For an example, see Listing collections in the Amazon Rekognition Developer
// Guide.
//
// This operation requires permissions to perform the rekognition:ListCollections
// action.
func rekognition_ListCollections(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.ListCollectionsInput{}

	if len(_rekognitionMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _rekognitionMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_rekognitionNextToken) > 0 {
		input.NextToken = aws.String(_rekognitionNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCollections(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rekognition.ListCollectionsOutput
	p := rekognition.NewListCollectionsPaginator(client, input)
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

// This operation applies only to Amazon Rekognition Custom Labels.
// Lists the entries (images) within a dataset. An entry is a JSON Line that
// contains the information for a single image, including the image location,
// assigned labels, and object location bounding boxes. For more information, see [Creating a manifest file].
//
// JSON Lines in the response include information about non-terminal errors found
// in the dataset. Non terminal errors are reported in errors lists within each
// JSON Line. The same information is reported in the training and testing
// validation result manifests that Amazon Rekognition Custom Labels creates during
// model training.
//
// You can filter the response in variety of ways, such as choosing which labels
// to return and returning JSON Lines created after a specific date.
//
// This operation requires permissions to perform the
// rekognition:ListDatasetEntries action.
//
// [Creating a manifest file]: https://docs.aws.amazon.com/rekognition/latest/customlabels-dg/md-manifest-files.html
func rekognition_ListDatasetEntries(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.ListDatasetEntriesInput{
		// DatasetArn: *string, // Required
	}

	if len(_rekognitionDatasetArn) > 0 {
		input.DatasetArn = aws.String(_rekognitionDatasetArn)
	}
	if len(_rekognitionContainsLabels) > 0 {
		input.ContainsLabels = append([]string(nil), _rekognitionContainsLabels...)
	}
	if len(_rekognitionHasErrors) > 0 {
		if err := assignInputField(input, "HasErrors", _rekognitionHasErrors); err != nil {
			log.Errorf("invalid --has-errors: %s", err.Error())
			return
		}
	}
	if len(_rekognitionLabeled) > 0 {
		if err := assignInputField(input, "Labeled", _rekognitionLabeled); err != nil {
			log.Errorf("invalid --labeled: %s", err.Error())
			return
		}
	}
	if len(_rekognitionMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _rekognitionMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_rekognitionNextToken) > 0 {
		input.NextToken = aws.String(_rekognitionNextToken)
	}
	if len(_rekognitionSourceRefContains) > 0 {
		input.SourceRefContains = aws.String(_rekognitionSourceRefContains)
	}

	if disablePaginator() {
		if resp, err := client.ListDatasetEntries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rekognition.ListDatasetEntriesOutput
	p := rekognition.NewListDatasetEntriesPaginator(client, input)
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

// This operation applies only to Amazon Rekognition Custom Labels.
// Lists the labels in a dataset. Amazon Rekognition Custom Labels uses labels to
// describe images. For more information, see [Labeling images].
//
// Lists the labels in a dataset. Amazon Rekognition Custom Labels uses labels to
// describe images. For more information, see Labeling images in the Amazon
// Rekognition Custom Labels Developer Guide.
//
// [Labeling images]: https://docs.aws.amazon.com/rekognition/latest/customlabels-dg/md-labeling-images.html
func rekognition_ListDatasetLabels(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.ListDatasetLabelsInput{
		// DatasetArn: *string, // Required
	}

	if len(_rekognitionDatasetArn) > 0 {
		input.DatasetArn = aws.String(_rekognitionDatasetArn)
	}
	if len(_rekognitionMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _rekognitionMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_rekognitionNextToken) > 0 {
		input.NextToken = aws.String(_rekognitionNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDatasetLabels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rekognition.ListDatasetLabelsOutput
	p := rekognition.NewListDatasetLabelsPaginator(client, input)
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

// Returns metadata for faces in the specified collection. This metadata includes
// information such as the bounding box coordinates, the confidence (that the
// bounding box contains a face), and face ID. For an example, see Listing Faces in
// a Collection in the Amazon Rekognition Developer Guide.
//
// This operation requires permissions to perform the rekognition:ListFaces action.
func rekognition_ListFaces(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.ListFacesInput{
		// CollectionId: *string, // Required
	}

	if len(_rekognitionCollectionId) > 0 {
		input.CollectionId = aws.String(_rekognitionCollectionId)
	}
	if len(_rekognitionFaceIds) > 0 {
		input.FaceIds = append([]string(nil), _rekognitionFaceIds...)
	}
	if len(_rekognitionMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _rekognitionMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_rekognitionNextToken) > 0 {
		input.NextToken = aws.String(_rekognitionNextToken)
	}
	if len(_rekognitionUserId) > 0 {
		input.UserId = aws.String(_rekognitionUserId)
	}

	if disablePaginator() {
		if resp, err := client.ListFaces(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rekognition.ListFacesOutput
	p := rekognition.NewListFacesPaginator(client, input)
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

// Returns a list of media analysis jobs. Results are sorted by CreationTimestamp
// in descending order.
func rekognition_ListMediaAnalysisJobs(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.ListMediaAnalysisJobsInput{}

	if len(_rekognitionMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _rekognitionMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_rekognitionNextToken) > 0 {
		input.NextToken = aws.String(_rekognitionNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMediaAnalysisJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rekognition.ListMediaAnalysisJobsOutput
	p := rekognition.NewListMediaAnalysisJobsPaginator(client, input)
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

// This operation applies only to Amazon Rekognition Custom Labels.
// Gets a list of the project policies attached to a project.
//
// To attach a project policy to a project, call PutProjectPolicy. To remove a project policy from
// a project, call DeleteProjectPolicy.
//
// This operation requires permissions to perform the
// rekognition:ListProjectPolicies action.
func rekognition_ListProjectPolicies(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.ListProjectPoliciesInput{
		// ProjectArn: *string, // Required
	}

	if len(_rekognitionProjectArn) > 0 {
		input.ProjectArn = aws.String(_rekognitionProjectArn)
	}
	if len(_rekognitionMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _rekognitionMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_rekognitionNextToken) > 0 {
		input.NextToken = aws.String(_rekognitionNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListProjectPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rekognition.ListProjectPoliciesOutput
	p := rekognition.NewListProjectPoliciesPaginator(client, input)
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

// Gets a list of stream processors that you have created with CreateStreamProcessor.
func rekognition_ListStreamProcessors(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.ListStreamProcessorsInput{}

	if len(_rekognitionMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _rekognitionMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_rekognitionNextToken) > 0 {
		input.NextToken = aws.String(_rekognitionNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListStreamProcessors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rekognition.ListStreamProcessorsOutput
	p := rekognition.NewListStreamProcessorsPaginator(client, input)
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

// Returns a list of tags in an Amazon Rekognition collection, stream processor,
// or Custom Labels model.
//
// This operation requires permissions to perform the
// rekognition:ListTagsForResource action.
func rekognition_ListTagsForResource(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_rekognitionResourceArn) > 0 {
		input.ResourceArn = aws.String(_rekognitionResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns metadata of the User such as UserID in the specified collection.
// Anonymous User (to reserve faces without any identity) is not returned as part
// of this request. The results are sorted by system generated primary key ID. If
// the response is truncated, NextToken is returned in the response that can be
// used in the subsequent request to retrieve the next set of identities.
func rekognition_ListUsers(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.ListUsersInput{
		// CollectionId: *string, // Required
	}

	if len(_rekognitionCollectionId) > 0 {
		input.CollectionId = aws.String(_rekognitionCollectionId)
	}
	if len(_rekognitionMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _rekognitionMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_rekognitionNextToken) > 0 {
		input.NextToken = aws.String(_rekognitionNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListUsers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*rekognition.ListUsersOutput
	p := rekognition.NewListUsersPaginator(client, input)
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

// This operation applies only to Amazon Rekognition Custom Labels.
// Attaches a project policy to a Amazon Rekognition Custom Labels project in a
// trusting AWS account. A project policy specifies that a trusted AWS account can
// copy a model version from a trusting AWS account to a project in the trusted AWS
// account. To copy a model version you use the CopyProjectVersionoperation. Only applies to Custom
// Labels projects.
//
// For more information about the format of a project policy document, see
// Attaching a project policy (SDK) in the Amazon Rekognition Custom Labels
// Developer Guide.
//
// The response from PutProjectPolicy is a revision ID for the project policy. You
// can attach multiple project policies to a project. You can also update an
// existing project policy by specifying the policy revision ID of the existing
// policy.
//
// To remove a project policy from a project, call DeleteProjectPolicy. To get a list of project
// policies attached to a project, call ListProjectPolicies.
//
// You copy a model version by calling CopyProjectVersion.
//
// This operation requires permissions to perform the rekognition:PutProjectPolicy
// action.
func rekognition_PutProjectPolicy(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.PutProjectPolicyInput{
		// PolicyDocument: *string, // Required
		// PolicyName: *string, // Required
		// ProjectArn: *string, // Required
	}

	if len(_rekognitionPolicyDocument) > 0 {
		input.PolicyDocument = aws.String(_rekognitionPolicyDocument)
	}
	if len(_rekognitionPolicyName) > 0 {
		input.PolicyName = aws.String(_rekognitionPolicyName)
	}
	if len(_rekognitionProjectArn) > 0 {
		input.ProjectArn = aws.String(_rekognitionProjectArn)
	}
	if len(_rekognitionPolicyRevisionId) > 0 {
		input.PolicyRevisionId = aws.String(_rekognitionPolicyRevisionId)
	}

	if resp, err := client.PutProjectPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an array of celebrities recognized in the input image. For more
// information, see Recognizing celebrities in the Amazon Rekognition Developer
// Guide.
//
// RecognizeCelebrities returns the 64 largest faces in the image. It lists the
// recognized celebrities in the CelebrityFaces array and any unrecognized faces
// in the UnrecognizedFaces array. RecognizeCelebrities doesn't return celebrities
// whose faces aren't among the largest 64 faces in the image.
//
// For each celebrity recognized, RecognizeCelebrities returns a Celebrity object.
// The Celebrity object contains the celebrity name, ID, URL links to additional
// information, match confidence, and a ComparedFace object that you can use to
// locate the celebrity's face on the image.
//
// Amazon Rekognition doesn't retain information about which images a celebrity
// has been recognized in. Your application must store this information and use the
// Celebrity ID property as a unique identifier for the celebrity. If you don't
// store the celebrity name or additional information URLs returned by
// RecognizeCelebrities , you will need the ID to identify the celebrity in a call
// to the GetCelebrityInfooperation.
//
// You pass the input image either as base64-encoded image bytes or as a reference
// to an image in an Amazon S3 bucket. If you use the AWS CLI to call Amazon
// Rekognition operations, passing image bytes is not supported. The image must be
// either a PNG or JPEG formatted file.
//
// For an example, see Recognizing celebrities in an image in the Amazon
// Rekognition Developer Guide.
//
// This operation requires permissions to perform the
// rekognition:RecognizeCelebrities operation.
func rekognition_RecognizeCelebrities(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.RecognizeCelebritiesInput{
		// Image: *types.Image, // Required
	}

	if len(_rekognitionImage) > 0 {
		if err := assignInputField(input, "Image", _rekognitionImage); err != nil {
			log.Errorf("invalid --image: %s", err.Error())
			return
		}
	}

	if resp, err := client.RecognizeCelebrities(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// For a given input face ID, searches for matching faces in the collection the
// face belongs to. You get a face ID when you add a face to the collection using
// the IndexFacesoperation. The operation compares the features of the input face with faces
// in the specified collection.
//
// You can also search faces without indexing faces by using the SearchFacesByImage
// operation.
//
// The operation response returns an array of faces that match, ordered by
// similarity score with the highest similarity first. More specifically, it is an
// array of metadata for each face match that is found. Along with the metadata,
// the response also includes a confidence value for each face match, indicating
// the confidence that the specific face matches the input face.
//
// For an example, see Searching for a face using its face ID in the Amazon
// Rekognition Developer Guide.
//
// This operation requires permissions to perform the rekognition:SearchFaces
// action.
func rekognition_SearchFaces(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.SearchFacesInput{
		// CollectionId: *string, // Required
		// FaceId: *string, // Required
	}

	if len(_rekognitionCollectionId) > 0 {
		input.CollectionId = aws.String(_rekognitionCollectionId)
	}
	if len(_rekognitionFaceId) > 0 {
		input.FaceId = aws.String(_rekognitionFaceId)
	}
	if len(_rekognitionFaceMatchThreshold) > 0 {
		if err := assignInputField(input, "FaceMatchThreshold", _rekognitionFaceMatchThreshold); err != nil {
			log.Errorf("invalid --face-match-threshold: %s", err.Error())
			return
		}
	}
	if len(_rekognitionMaxFaces) > 0 {
		if err := assignInputField(input, "MaxFaces", _rekognitionMaxFaces); err != nil {
			log.Errorf("invalid --max-faces: %s", err.Error())
			return
		}
	}

	if resp, err := client.SearchFaces(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// For a given input image, first detects the largest face in the image, and then
// searches the specified collection for matching faces. The operation compares the
// features of the input face with faces in the specified collection.
//
// To search for all faces in an input image, you might first call the IndexFaces operation,
// and then use the face IDs returned in subsequent calls to the SearchFacesoperation.
//
// You can also call the DetectFaces operation and use the bounding boxes in the
// response to make face crops, which then you can pass in to the
// SearchFacesByImage operation.
//
// You pass the input image either as base64-encoded image bytes or as a reference
// to an image in an Amazon S3 bucket. If you use the AWS CLI to call Amazon
// Rekognition operations, passing image bytes is not supported. The image must be
// either a PNG or JPEG formatted file.
//
// The response returns an array of faces that match, ordered by similarity score
// with the highest similarity first. More specifically, it is an array of metadata
// for each face match found. Along with the metadata, the response also includes a
// similarity indicating how similar the face is to the input face. In the
// response, the operation also returns the bounding box (and a confidence level
// that the bounding box contains a face) of the face that Amazon Rekognition used
// for the input image.
//
// If no faces are detected in the input image, SearchFacesByImage returns an
// InvalidParameterException error.
//
// For an example, Searching for a Face Using an Image in the Amazon Rekognition
// Developer Guide.
//
// The QualityFilter input parameter allows you to filter out detected faces that
// don’t meet a required quality bar. The quality bar is based on a variety of
// common use cases. Use QualityFilter to set the quality bar for filtering by
// specifying LOW , MEDIUM , or HIGH . If you do not want to filter detected faces,
// specify NONE . The default value is NONE .
//
// To use quality filtering, you need a collection associated with version 3 of
// the face model or higher. To get the version of the face model associated with a
// collection, call DescribeCollection.
//
// This operation requires permissions to perform the
// rekognition:SearchFacesByImage action.
func rekognition_SearchFacesByImage(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.SearchFacesByImageInput{
		// CollectionId: *string, // Required
		// Image: *types.Image, // Required
	}

	if len(_rekognitionCollectionId) > 0 {
		input.CollectionId = aws.String(_rekognitionCollectionId)
	}
	if len(_rekognitionImage) > 0 {
		if err := assignInputField(input, "Image", _rekognitionImage); err != nil {
			log.Errorf("invalid --image: %s", err.Error())
			return
		}
	}
	if len(_rekognitionFaceMatchThreshold) > 0 {
		if err := assignInputField(input, "FaceMatchThreshold", _rekognitionFaceMatchThreshold); err != nil {
			log.Errorf("invalid --face-match-threshold: %s", err.Error())
			return
		}
	}
	if len(_rekognitionMaxFaces) > 0 {
		if err := assignInputField(input, "MaxFaces", _rekognitionMaxFaces); err != nil {
			log.Errorf("invalid --max-faces: %s", err.Error())
			return
		}
	}
	if len(_rekognitionQualityFilter) > 0 {
		if err := assignInputField(input, "QualityFilter", _rekognitionQualityFilter); err != nil {
			log.Errorf("invalid --quality-filter: %s", err.Error())
			return
		}
	}

	if resp, err := client.SearchFacesByImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Searches for UserIDs within a collection based on a FaceId or UserId . This API
// can be used to find the closest UserID (with a highest similarity) to associate
// a face. The request must be provided with either FaceId or UserId . The
// operation returns an array of UserID that match the FaceId or UserId , ordered
// by similarity score with the highest similarity first.
func rekognition_SearchUsers(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.SearchUsersInput{
		// CollectionId: *string, // Required
	}

	if len(_rekognitionCollectionId) > 0 {
		input.CollectionId = aws.String(_rekognitionCollectionId)
	}
	if len(_rekognitionFaceId) > 0 {
		input.FaceId = aws.String(_rekognitionFaceId)
	}
	if len(_rekognitionMaxUsers) > 0 {
		if err := assignInputField(input, "MaxUsers", _rekognitionMaxUsers); err != nil {
			log.Errorf("invalid --max-users: %s", err.Error())
			return
		}
	}
	if len(_rekognitionUserId) > 0 {
		input.UserId = aws.String(_rekognitionUserId)
	}
	if len(_rekognitionUserMatchThreshold) > 0 {
		if err := assignInputField(input, "UserMatchThreshold", _rekognitionUserMatchThreshold); err != nil {
			log.Errorf("invalid --user-match-threshold: %s", err.Error())
			return
		}
	}

	if resp, err := client.SearchUsers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Searches for UserIDs using a supplied image. It first detects the largest face
// in the image, and then searches a specified collection for matching UserIDs.
//
// The operation returns an array of UserIDs that match the face in the supplied
// image, ordered by similarity score with the highest similarity first. It also
// returns a bounding box for the face found in the input image.
//
// Information about faces detected in the supplied image, but not used for the
// search, is returned in an array of UnsearchedFace objects. If no valid face is
// detected in the image, the response will contain an empty UserMatches list and
// no SearchedFace object.
func rekognition_SearchUsersByImage(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.SearchUsersByImageInput{
		// CollectionId: *string, // Required
		// Image: *types.Image, // Required
	}

	if len(_rekognitionCollectionId) > 0 {
		input.CollectionId = aws.String(_rekognitionCollectionId)
	}
	if len(_rekognitionImage) > 0 {
		if err := assignInputField(input, "Image", _rekognitionImage); err != nil {
			log.Errorf("invalid --image: %s", err.Error())
			return
		}
	}
	if len(_rekognitionMaxUsers) > 0 {
		if err := assignInputField(input, "MaxUsers", _rekognitionMaxUsers); err != nil {
			log.Errorf("invalid --max-users: %s", err.Error())
			return
		}
	}
	if len(_rekognitionQualityFilter) > 0 {
		if err := assignInputField(input, "QualityFilter", _rekognitionQualityFilter); err != nil {
			log.Errorf("invalid --quality-filter: %s", err.Error())
			return
		}
	}
	if len(_rekognitionUserMatchThreshold) > 0 {
		if err := assignInputField(input, "UserMatchThreshold", _rekognitionUserMatchThreshold); err != nil {
			log.Errorf("invalid --user-match-threshold: %s", err.Error())
			return
		}
	}

	if resp, err := client.SearchUsersByImage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts asynchronous recognition of celebrities in a stored video.
// Amazon Rekognition Video can detect celebrities in a video must be stored in an
// Amazon S3 bucket. Use Videoto specify the bucket name and the filename of the video.
// StartCelebrityRecognition returns a job identifier ( JobId ) which you use to
// get the results of the analysis. When celebrity recognition analysis is
// finished, Amazon Rekognition Video publishes a completion status to the Amazon
// Simple Notification Service topic that you specify in NotificationChannel . To
// get the results of the celebrity recognition analysis, first check that the
// status value published to the Amazon SNS topic is SUCCEEDED . If so, call GetCelebrityRecognition and
// pass the job identifier ( JobId ) from the initial call to
// StartCelebrityRecognition .
//
// For more information, see Recognizing celebrities in the Amazon Rekognition
// Developer Guide.
func rekognition_StartCelebrityRecognition(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.StartCelebrityRecognitionInput{
		// Video: *types.Video, // Required
	}

	if len(_rekognitionVideo) > 0 {
		if err := assignInputField(input, "Video", _rekognitionVideo); err != nil {
			log.Errorf("invalid --video: %s", err.Error())
			return
		}
	}
	if len(_rekognitionClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_rekognitionClientRequestToken)
	}
	if len(_rekognitionJobTag) > 0 {
		input.JobTag = aws.String(_rekognitionJobTag)
	}
	if len(_rekognitionNotificationChannel) > 0 {
		if err := assignInputField(input, "NotificationChannel", _rekognitionNotificationChannel); err != nil {
			log.Errorf("invalid --notification-channel: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartCelebrityRecognition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts asynchronous detection of inappropriate, unwanted, or offensive content
// in a stored video. For a list of moderation labels in Amazon Rekognition, see [Using the image and video moderation APIs].
//
// Amazon Rekognition Video can moderate content in a video stored in an Amazon S3
// bucket. Use Videoto specify the bucket name and the filename of the video.
// StartContentModeration returns a job identifier ( JobId ) which you use to get
// the results of the analysis. When content analysis is finished, Amazon
// Rekognition Video publishes a completion status to the Amazon Simple
// Notification Service topic that you specify in NotificationChannel .
//
// To get the results of the content analysis, first check that the status value
// published to the Amazon SNS topic is SUCCEEDED . If so, call GetContentModeration and pass the job
// identifier ( JobId ) from the initial call to StartContentModeration .
//
// For more information, see Moderating content in the Amazon Rekognition
// Developer Guide.
//
// [Using the image and video moderation APIs]: https://docs.aws.amazon.com/rekognition/latest/dg/moderation.html#moderation-api
func rekognition_StartContentModeration(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.StartContentModerationInput{
		// Video: *types.Video, // Required
	}

	if len(_rekognitionVideo) > 0 {
		if err := assignInputField(input, "Video", _rekognitionVideo); err != nil {
			log.Errorf("invalid --video: %s", err.Error())
			return
		}
	}
	if len(_rekognitionClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_rekognitionClientRequestToken)
	}
	if len(_rekognitionJobTag) > 0 {
		input.JobTag = aws.String(_rekognitionJobTag)
	}
	if len(_rekognitionMinConfidence) > 0 {
		if err := assignInputField(input, "MinConfidence", _rekognitionMinConfidence); err != nil {
			log.Errorf("invalid --min-confidence: %s", err.Error())
			return
		}
	}
	if len(_rekognitionNotificationChannel) > 0 {
		if err := assignInputField(input, "NotificationChannel", _rekognitionNotificationChannel); err != nil {
			log.Errorf("invalid --notification-channel: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartContentModeration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts asynchronous detection of faces in a stored video.
// Amazon Rekognition Video can detect faces in a video stored in an Amazon S3
// bucket. Use Videoto specify the bucket name and the filename of the video.
// StartFaceDetection returns a job identifier ( JobId ) that you use to get the
// results of the operation. When face detection is finished, Amazon Rekognition
// Video publishes a completion status to the Amazon Simple Notification Service
// topic that you specify in NotificationChannel . To get the results of the face
// detection operation, first check that the status value published to the Amazon
// SNS topic is SUCCEEDED . If so, call GetFaceDetection and pass the job identifier ( JobId ) from
// the initial call to StartFaceDetection .
//
// For more information, see Detecting faces in a stored video in the Amazon
// Rekognition Developer Guide.
func rekognition_StartFaceDetection(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.StartFaceDetectionInput{
		// Video: *types.Video, // Required
	}

	if len(_rekognitionVideo) > 0 {
		if err := assignInputField(input, "Video", _rekognitionVideo); err != nil {
			log.Errorf("invalid --video: %s", err.Error())
			return
		}
	}
	if len(_rekognitionClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_rekognitionClientRequestToken)
	}
	if len(_rekognitionFaceAttributes) > 0 {
		if err := assignInputField(input, "FaceAttributes", _rekognitionFaceAttributes); err != nil {
			log.Errorf("invalid --face-attributes: %s", err.Error())
			return
		}
	}
	if len(_rekognitionJobTag) > 0 {
		input.JobTag = aws.String(_rekognitionJobTag)
	}
	if len(_rekognitionNotificationChannel) > 0 {
		if err := assignInputField(input, "NotificationChannel", _rekognitionNotificationChannel); err != nil {
			log.Errorf("invalid --notification-channel: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartFaceDetection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the asynchronous search for faces in a collection that match the faces
// of persons detected in a stored video.
//
// The video must be stored in an Amazon S3 bucket. Use Video to specify the bucket
// name and the filename of the video. StartFaceSearch returns a job identifier (
// JobId ) which you use to get the search results once the search has completed.
// When searching is finished, Amazon Rekognition Video publishes a completion
// status to the Amazon Simple Notification Service topic that you specify in
// NotificationChannel . To get the search results, first check that the status
// value published to the Amazon SNS topic is SUCCEEDED . If so, call GetFaceSearch and pass
// the job identifier ( JobId ) from the initial call to StartFaceSearch . For more
// information, see [Searching stored videos for faces].
//
// [Searching stored videos for faces]: https://docs.aws.amazon.com/rekognition/latest/dg/procedure-person-search-videos.html
func rekognition_StartFaceSearch(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.StartFaceSearchInput{
		// CollectionId: *string, // Required
		// Video: *types.Video, // Required
	}

	if len(_rekognitionCollectionId) > 0 {
		input.CollectionId = aws.String(_rekognitionCollectionId)
	}
	if len(_rekognitionVideo) > 0 {
		if err := assignInputField(input, "Video", _rekognitionVideo); err != nil {
			log.Errorf("invalid --video: %s", err.Error())
			return
		}
	}
	if len(_rekognitionClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_rekognitionClientRequestToken)
	}
	if len(_rekognitionFaceMatchThreshold) > 0 {
		if err := assignInputField(input, "FaceMatchThreshold", _rekognitionFaceMatchThreshold); err != nil {
			log.Errorf("invalid --face-match-threshold: %s", err.Error())
			return
		}
	}
	if len(_rekognitionJobTag) > 0 {
		input.JobTag = aws.String(_rekognitionJobTag)
	}
	if len(_rekognitionNotificationChannel) > 0 {
		if err := assignInputField(input, "NotificationChannel", _rekognitionNotificationChannel); err != nil {
			log.Errorf("invalid --notification-channel: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartFaceSearch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts asynchronous detection of labels in a stored video.
// Amazon Rekognition Video can detect labels in a video. Labels are instances of
// real-world entities. This includes objects like flower, tree, and table; events
// like wedding, graduation, and birthday party; concepts like landscape, evening,
// and nature; and activities like a person getting out of a car or a person
// skiing.
//
// The video must be stored in an Amazon S3 bucket. Use Video to specify the bucket
// name and the filename of the video. StartLabelDetection returns a job
// identifier ( JobId ) which you use to get the results of the operation. When
// label detection is finished, Amazon Rekognition Video publishes a completion
// status to the Amazon Simple Notification Service topic that you specify in
// NotificationChannel .
//
// To get the results of the label detection operation, first check that the
// status value published to the Amazon SNS topic is SUCCEEDED . If so, call GetLabelDetection and
// pass the job identifier ( JobId ) from the initial call to StartLabelDetection .
//
// # Optional Parameters
//
// StartLabelDetection has the GENERAL_LABELS Feature applied by default. This
// feature allows you to provide filtering criteria to the Settings parameter. You
// can filter with sets of individual labels or with label categories. You can
// specify inclusive filters, exclusive filters, or a combination of inclusive and
// exclusive filters. For more information on filtering, see [Detecting labels in a video].
//
// You can specify MinConfidence to control the confidence threshold for the
// labels returned. The default is 50.
//
// [Detecting labels in a video]: https://docs.aws.amazon.com/rekognition/latest/dg/labels-detecting-labels-video.html
func rekognition_StartLabelDetection(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.StartLabelDetectionInput{
		// Video: *types.Video, // Required
	}

	if len(_rekognitionVideo) > 0 {
		if err := assignInputField(input, "Video", _rekognitionVideo); err != nil {
			log.Errorf("invalid --video: %s", err.Error())
			return
		}
	}
	if len(_rekognitionClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_rekognitionClientRequestToken)
	}
	if len(_rekognitionFeatures) > 0 {
		if err := assignInputField(input, "Features", _rekognitionFeatures); err != nil {
			log.Errorf("invalid --features: %s", err.Error())
			return
		}
	}
	if len(_rekognitionJobTag) > 0 {
		input.JobTag = aws.String(_rekognitionJobTag)
	}
	if len(_rekognitionMinConfidence) > 0 {
		if err := assignInputField(input, "MinConfidence", _rekognitionMinConfidence); err != nil {
			log.Errorf("invalid --min-confidence: %s", err.Error())
			return
		}
	}
	if len(_rekognitionNotificationChannel) > 0 {
		if err := assignInputField(input, "NotificationChannel", _rekognitionNotificationChannel); err != nil {
			log.Errorf("invalid --notification-channel: %s", err.Error())
			return
		}
	}
	if len(_rekognitionSettings) > 0 {
		if err := assignInputField(input, "Settings", _rekognitionSettings); err != nil {
			log.Errorf("invalid --settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartLabelDetection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates a new media analysis job. Accepts a manifest file in an Amazon S3
// bucket. The output is a manifest file and a summary of the manifest stored in
// the Amazon S3 bucket.
func rekognition_StartMediaAnalysisJob(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.StartMediaAnalysisJobInput{
		// Input: *types.MediaAnalysisInput, // Required
		// OperationsConfig: *types.MediaAnalysisOperationsConfig, // Required
		// OutputConfig: *types.MediaAnalysisOutputConfig, // Required
	}

	if len(_rekognitionInput) > 0 {
		if err := assignInputField(input, "Input", _rekognitionInput); err != nil {
			log.Errorf("invalid --input: %s", err.Error())
			return
		}
	}
	if len(_rekognitionOperationsConfig) > 0 {
		if err := assignInputField(input, "OperationsConfig", _rekognitionOperationsConfig); err != nil {
			log.Errorf("invalid --operations-config: %s", err.Error())
			return
		}
	}
	if len(_rekognitionOutputConfig) > 0 {
		if err := assignInputField(input, "OutputConfig", _rekognitionOutputConfig); err != nil {
			log.Errorf("invalid --output-config: %s", err.Error())
			return
		}
	}
	if len(_rekognitionClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_rekognitionClientRequestToken)
	}
	if len(_rekognitionJobName) > 0 {
		input.JobName = aws.String(_rekognitionJobName)
	}
	if len(_rekognitionKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_rekognitionKmsKeyId)
	}

	if resp, err := client.StartMediaAnalysisJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// End of support notice: On October 31, 2025, AWS will discontinue support for
// Amazon Rekognition People Pathing. After October 31, 2025, you will no longer be
// able to use the Rekognition People Pathing capability. For more information,
// visit this [blog post].
//
// Starts the asynchronous tracking of a person's path in a stored video.
//
// Amazon Rekognition Video can track the path of people in a video stored in an
// Amazon S3 bucket. Use Videoto specify the bucket name and the filename of the video.
// StartPersonTracking returns a job identifier ( JobId ) which you use to get the
// results of the operation. When label detection is finished, Amazon Rekognition
// publishes a completion status to the Amazon Simple Notification Service topic
// that you specify in NotificationChannel .
//
// To get the results of the person detection operation, first check that the
// status value published to the Amazon SNS topic is SUCCEEDED . If so, call GetPersonTracking and
// pass the job identifier ( JobId ) from the initial call to StartPersonTracking .
//
// [blog post]: https://aws.amazon.com/blogs/machine-learning/transitioning-from-amazon-rekognition-people-pathing-exploring-other-alternatives/
func rekognition_StartPersonTracking(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.StartPersonTrackingInput{
		// Video: *types.Video, // Required
	}

	if len(_rekognitionVideo) > 0 {
		if err := assignInputField(input, "Video", _rekognitionVideo); err != nil {
			log.Errorf("invalid --video: %s", err.Error())
			return
		}
	}
	if len(_rekognitionClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_rekognitionClientRequestToken)
	}
	if len(_rekognitionJobTag) > 0 {
		input.JobTag = aws.String(_rekognitionJobTag)
	}
	if len(_rekognitionNotificationChannel) > 0 {
		if err := assignInputField(input, "NotificationChannel", _rekognitionNotificationChannel); err != nil {
			log.Errorf("invalid --notification-channel: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartPersonTracking(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation applies only to Amazon Rekognition Custom Labels.
// Starts the running of the version of a model. Starting a model takes a while to
// complete. To check the current state of the model, use DescribeProjectVersions.
//
// Once the model is running, you can detect custom labels in new images by
// calling DetectCustomLabels.
//
// You are charged for the amount of time that the model is running. To stop a
// running model, call StopProjectVersion.
//
// This operation requires permissions to perform the
// rekognition:StartProjectVersion action.
func rekognition_StartProjectVersion(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.StartProjectVersionInput{
		// MinInferenceUnits: *int32, // Required
		// ProjectVersionArn: *string, // Required
	}

	if len(_rekognitionMinInferenceUnits) > 0 {
		if err := assignInputField(input, "MinInferenceUnits", _rekognitionMinInferenceUnits); err != nil {
			log.Errorf("invalid --min-inference-units: %s", err.Error())
			return
		}
	}
	if len(_rekognitionProjectVersionArn) > 0 {
		input.ProjectVersionArn = aws.String(_rekognitionProjectVersionArn)
	}
	if len(_rekognitionMaxInferenceUnits) > 0 {
		if err := assignInputField(input, "MaxInferenceUnits", _rekognitionMaxInferenceUnits); err != nil {
			log.Errorf("invalid --max-inference-units: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartProjectVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts asynchronous detection of segment detection in a stored video.
// Amazon Rekognition Video can detect segments in a video stored in an Amazon S3
// bucket. Use Videoto specify the bucket name and the filename of the video.
// StartSegmentDetection returns a job identifier ( JobId ) which you use to get
// the results of the operation. When segment detection is finished, Amazon
// Rekognition Video publishes a completion status to the Amazon Simple
// Notification Service topic that you specify in NotificationChannel .
//
// You can use the Filters (StartSegmentDetectionFilters ) input parameter to specify the minimum detection
// confidence returned in the response. Within Filters , use ShotFilter (StartShotDetectionFilter ) to
// filter detected shots. Use TechnicalCueFilter (StartTechnicalCueDetectionFilter ) to filter technical cues.
//
// To get the results of the segment detection operation, first check that the
// status value published to the Amazon SNS topic is SUCCEEDED . if so, call GetSegmentDetection and
// pass the job identifier ( JobId ) from the initial call to StartSegmentDetection
// .
//
// For more information, see Detecting video segments in stored video in the
// Amazon Rekognition Developer Guide.
func rekognition_StartSegmentDetection(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.StartSegmentDetectionInput{
		// SegmentTypes: []types.SegmentType, // Required
		// Video: *types.Video, // Required
	}

	if len(_rekognitionSegmentTypes) > 0 {
		if err := assignInputField(input, "SegmentTypes", _rekognitionSegmentTypes); err != nil {
			log.Errorf("invalid --segment-types: %s", err.Error())
			return
		}
	}
	if len(_rekognitionVideo) > 0 {
		if err := assignInputField(input, "Video", _rekognitionVideo); err != nil {
			log.Errorf("invalid --video: %s", err.Error())
			return
		}
	}
	if len(_rekognitionClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_rekognitionClientRequestToken)
	}
	if len(_rekognitionFilters) > 0 {
		if err := assignInputField(input, "Filters", _rekognitionFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rekognitionJobTag) > 0 {
		input.JobTag = aws.String(_rekognitionJobTag)
	}
	if len(_rekognitionNotificationChannel) > 0 {
		if err := assignInputField(input, "NotificationChannel", _rekognitionNotificationChannel); err != nil {
			log.Errorf("invalid --notification-channel: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartSegmentDetection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts processing a stream processor. You create a stream processor by calling CreateStreamProcessor
// . To tell StartStreamProcessor which stream processor to start, use the value
// of the Name field specified in the call to CreateStreamProcessor .
//
// If you are using a label detection stream processor to detect labels, you need
// to provide a Start selector and a Stop selector to determine the length of the
// stream processing time.
func rekognition_StartStreamProcessor(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.StartStreamProcessorInput{
		// Name: *string, // Required
	}

	if len(_rekognitionName) > 0 {
		input.Name = aws.String(_rekognitionName)
	}
	if len(_rekognitionStartSelector) > 0 {
		if err := assignInputField(input, "StartSelector", _rekognitionStartSelector); err != nil {
			log.Errorf("invalid --start-selector: %s", err.Error())
			return
		}
	}
	if len(_rekognitionStopSelector) > 0 {
		if err := assignInputField(input, "StopSelector", _rekognitionStopSelector); err != nil {
			log.Errorf("invalid --stop-selector: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartStreamProcessor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts asynchronous detection of text in a stored video.
// Amazon Rekognition Video can detect text in a video stored in an Amazon S3
// bucket. Use Videoto specify the bucket name and the filename of the video.
// StartTextDetection returns a job identifier ( JobId ) which you use to get the
// results of the operation. When text detection is finished, Amazon Rekognition
// Video publishes a completion status to the Amazon Simple Notification Service
// topic that you specify in NotificationChannel .
//
// To get the results of the text detection operation, first check that the status
// value published to the Amazon SNS topic is SUCCEEDED . if so, call GetTextDetection and pass
// the job identifier ( JobId ) from the initial call to StartTextDetection .
func rekognition_StartTextDetection(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.StartTextDetectionInput{
		// Video: *types.Video, // Required
	}

	if len(_rekognitionVideo) > 0 {
		if err := assignInputField(input, "Video", _rekognitionVideo); err != nil {
			log.Errorf("invalid --video: %s", err.Error())
			return
		}
	}
	if len(_rekognitionClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_rekognitionClientRequestToken)
	}
	if len(_rekognitionFilters) > 0 {
		if err := assignInputField(input, "Filters", _rekognitionFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_rekognitionJobTag) > 0 {
		input.JobTag = aws.String(_rekognitionJobTag)
	}
	if len(_rekognitionNotificationChannel) > 0 {
		if err := assignInputField(input, "NotificationChannel", _rekognitionNotificationChannel); err != nil {
			log.Errorf("invalid --notification-channel: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartTextDetection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation applies only to Amazon Rekognition Custom Labels.
// Stops a running model. The operation might take a while to complete. To check
// the current status, call DescribeProjectVersions. Only applies to Custom Labels projects.
//
// This operation requires permissions to perform the
// rekognition:StopProjectVersion action.
func rekognition_StopProjectVersion(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.StopProjectVersionInput{
		// ProjectVersionArn: *string, // Required
	}

	if len(_rekognitionProjectVersionArn) > 0 {
		input.ProjectVersionArn = aws.String(_rekognitionProjectVersionArn)
	}

	if resp, err := client.StopProjectVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a running stream processor that was created by CreateStreamProcessor.
func rekognition_StopStreamProcessor(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.StopStreamProcessorInput{
		// Name: *string, // Required
	}

	if len(_rekognitionName) > 0 {
		input.Name = aws.String(_rekognitionName)
	}

	if resp, err := client.StopStreamProcessor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more key-value tags to an Amazon Rekognition collection, stream
// processor, or Custom Labels model. For more information, see [Tagging AWS Resources].
//
// This operation requires permissions to perform the rekognition:TagResource
// action.
//
// [Tagging AWS Resources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
func rekognition_TagResource(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_rekognitionResourceArn) > 0 {
		input.ResourceArn = aws.String(_rekognitionResourceArn)
	}
	if len(_rekognitionTags) > 0 {
		if err := assignInputField(input, "Tags", _rekognitionTags); err != nil {
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

// Removes one or more tags from an Amazon Rekognition collection, stream
// processor, or Custom Labels model.
//
// This operation requires permissions to perform the rekognition:UntagResource
// action.
func rekognition_UntagResource(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_rekognitionResourceArn) > 0 {
		input.ResourceArn = aws.String(_rekognitionResourceArn)
	}
	if len(_rekognitionTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _rekognitionTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation applies only to Amazon Rekognition Custom Labels.
// Adds or updates one or more entries (images) in a dataset. An entry is a JSON
// Line which contains the information for a single image, including the image
// location, assigned labels, and object location bounding boxes. For more
// information, see Image-Level labels in manifest files and Object localization in
// manifest files in the Amazon Rekognition Custom Labels Developer Guide.
//
// If the source-ref field in the JSON line references an existing image, the
// existing image in the dataset is updated. If source-ref field doesn't reference
// an existing image, the image is added as a new image to the dataset.
//
// You specify the changes that you want to make in the Changes input parameter.
// There isn't a limit to the number JSON Lines that you can change, but the size
// of Changes must be less than 5MB.
//
// UpdateDatasetEntries returns immediatly, but the dataset update might take a
// while to complete. Use DescribeDatasetto check the current status. The dataset updated
// successfully if the value of Status is UPDATE_COMPLETE .
//
// To check if any non-terminal errors occured, call ListDatasetEntries and check for the presence
// of errors lists in the JSON Lines.
//
// Dataset update fails if a terminal error occurs ( Status = UPDATE_FAILED ).
// Currently, you can't access the terminal error information from the Amazon
// Rekognition Custom Labels SDK.
//
// This operation requires permissions to perform the
// rekognition:UpdateDatasetEntries action.
func rekognition_UpdateDatasetEntries(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.UpdateDatasetEntriesInput{
		// Changes: *types.DatasetChanges, // Required
		// DatasetArn: *string, // Required
	}

	if len(_rekognitionChanges) > 0 {
		if err := assignInputField(input, "Changes", _rekognitionChanges); err != nil {
			log.Errorf("invalid --changes: %s", err.Error())
			return
		}
	}
	if len(_rekognitionDatasetArn) > 0 {
		input.DatasetArn = aws.String(_rekognitionDatasetArn)
	}

	if resp, err := client.UpdateDatasetEntries(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows you to update a stream processor. You can change some settings and
// regions of interest and delete certain parameters.
func rekognition_UpdateStreamProcessor(cfg aws.Config, client *rekognition.Client) {
	input := &rekognition.UpdateStreamProcessorInput{
		// Name: *string, // Required
	}

	if len(_rekognitionName) > 0 {
		input.Name = aws.String(_rekognitionName)
	}
	if len(_rekognitionDataSharingPreferenceForUpdate) > 0 {
		if err := assignInputField(input, "DataSharingPreferenceForUpdate", _rekognitionDataSharingPreferenceForUpdate); err != nil {
			log.Errorf("invalid --data-sharing-preference-for-update: %s", err.Error())
			return
		}
	}
	if len(_rekognitionParametersToDelete) > 0 {
		if err := assignInputField(input, "ParametersToDelete", _rekognitionParametersToDelete); err != nil {
			log.Errorf("invalid --parameters-to-delete: %s", err.Error())
			return
		}
	}
	if len(_rekognitionRegionsOfInterestForUpdate) > 0 {
		if err := assignInputField(input, "RegionsOfInterestForUpdate", _rekognitionRegionsOfInterestForUpdate); err != nil {
			log.Errorf("invalid --regions-of-interest-for-update: %s", err.Error())
			return
		}
	}
	if len(_rekognitionSettingsForUpdate) > 0 {
		if err := assignInputField(input, "SettingsForUpdate", _rekognitionSettingsForUpdate); err != nil {
			log.Errorf("invalid --settings-for-update: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateStreamProcessor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_rekognitionCmd)
	_rekognitionCmd.Flags().SortFlags = false

	_rekognitionCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_rekognitionCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_rekognitionCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_rekognitionCmd.Flags().StringVarP(&_rekognitionAggregateBy, "aggregate-by", "", "", "Aggregate By")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionAttributes, "attributes", "", "", "Attributes")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionAutoUpdate, "auto-update", "", "", "Auto Update")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionChanges, "changes", "", "", "Changes")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionCollectionId, "collection-id", "", "", "Collection ID")
	_rekognitionCmd.Flags().StringSliceVarP(&_rekognitionContainsLabels, "contains-labels", "", nil, "Contains Labels")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionDataSharingPreference, "data-sharing-preference", "", "", "Data Sharing Preference")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionDataSharingPreferenceForUpdate, "data-sharing-preference-for-update", "", "", "Data Sharing Preference For Update")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionDatasetArn, "dataset-arn", "", "", "Dataset ARN")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionDatasetSource, "dataset-source", "", "", "Dataset Source")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionDatasetType, "dataset-type", "", "", "Dataset Type")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionDatasets, "datasets", "", "", "Datasets")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionDestinationProjectArn, "destination-project-arn", "", "", "Destination Project ARN")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionDetectionAttributes, "detection-attributes", "", "", "Detection Attributes")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionExternalImageId, "external-image-id", "", "", "External Image ID")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionFaceAttributes, "face-attributes", "", "", "Face Attributes")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionFaceId, "face-id", "", "", "Face ID")
	_rekognitionCmd.Flags().StringSliceVarP(&_rekognitionFaceIds, "face-ids", "", nil, "Face Ids")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionFaceMatchThreshold, "face-match-threshold", "", "", "Face Match Threshold")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionFeature, "feature", "", "", "Feature")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionFeatureConfig, "feature-config", "", "", "Feature Config")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionFeatures, "features", "", "", "Features")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionFilters, "filters", "", "", "Filters")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionHasErrors, "has-errors", "", "", "Has Errors")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionHumanLoopConfig, "human-loop-config", "", "", "Human Loop Config")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionId, "id", "", "", "ID")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionImage, "image", "", "", "Image")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionInput, "input", "", "", "Input")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionJobId, "job-id", "", "", "Job ID")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionJobName, "job-name", "", "", "Job Name")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionJobTag, "job-tag", "", "", "Job Tag")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionLabeled, "labeled", "", "", "Labeled")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionMaxFaces, "max-faces", "", "", "Max Faces")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionMaxInferenceUnits, "max-inference-units", "", "", "Max Inference Units")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionMaxLabels, "max-labels", "", "", "Max Labels")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionMaxResults, "max-results", "", "", "Max Results")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionMaxUsers, "max-users", "", "", "Max Users")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionMinConfidence, "min-confidence", "", "", "Min Confidence")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionMinInferenceUnits, "min-inference-units", "", "", "Min Inference Units")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionName, "name", "", "", "Name")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionNextToken, "next-token", "", "", "Next Token")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionNotificationChannel, "notification-channel", "", "", "Notification Channel")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionOperationsConfig, "operations-config", "", "", "Operations Config")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionOutputConfig, "output-config", "", "", "Output Config")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionParametersToDelete, "parameters-to-delete", "", "", "Parameters To Delete")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionPolicyDocument, "policy-document", "", "", "Policy Document")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionPolicyName, "policy-name", "", "", "Policy Name")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionPolicyRevisionId, "policy-revision-id", "", "", "Policy Revision ID")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionProjectArn, "project-arn", "", "", "Project ARN")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionProjectName, "project-name", "", "", "Project Name")
	_rekognitionCmd.Flags().StringSliceVarP(&_rekognitionProjectNames, "project-names", "", nil, "Project Names")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionProjectVersion, "project-version", "", "", "Project Version")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionProjectVersionArn, "project-version-arn", "", "", "Project Version ARN")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionQualityFilter, "quality-filter", "", "", "Quality Filter")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionRegionsOfInterest, "regions-of-interest", "", "", "Regions Of Interest")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionRegionsOfInterestForUpdate, "regions-of-interest-for-update", "", "", "Regions Of Interest For Update")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionResourceArn, "resource-arn", "", "", "Resource ARN")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionRoleArn, "role-arn", "", "", "Role ARN")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionSegmentTypes, "segment-types", "", "", "Segment Types")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionSessionId, "session-id", "", "", "Session ID")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionSettings, "settings", "", "", "Settings")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionSettingsForUpdate, "settings-for-update", "", "", "Settings For Update")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionSimilarityThreshold, "similarity-threshold", "", "", "Similarity Threshold")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionSortBy, "sort-by", "", "", "Sort By")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionSourceImage, "source-image", "", "", "Source Image")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionSourceProjectArn, "source-project-arn", "", "", "Source Project ARN")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionSourceProjectVersionArn, "source-project-version-arn", "", "", "Source Project Version ARN")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionSourceRefContains, "source-ref-contains", "", "", "Source Ref Contains")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionStartSelector, "start-selector", "", "", "Start Selector")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionStopSelector, "stop-selector", "", "", "Stop Selector")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionSummarizationAttributes, "summarization-attributes", "", "", "Summarization Attributes")
	_rekognitionCmd.Flags().StringSliceVarP(&_rekognitionTagKeys, "tag-keys", "", nil, "Tag Keys")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionTags, "tags", "", "", "Tags")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionTargetImage, "target-image", "", "", "Target Image")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionTestingData, "testing-data", "", "", "Testing Data")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionTrainingData, "training-data", "", "", "Training Data")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionUserId, "user-id", "", "", "User ID")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionUserMatchThreshold, "user-match-threshold", "", "", "User Match Threshold")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionVersionDescription, "version-description", "", "", "Version Description")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionVersionName, "version-name", "", "", "Version Name")
	_rekognitionCmd.Flags().StringSliceVarP(&_rekognitionVersionNames, "version-names", "", nil, "Version Names")
	_rekognitionCmd.Flags().StringVarP(&_rekognitionVideo, "video", "", "", "Video")

	_rekognitionCmd.Flags().BoolVarP(&_rekognitionAssociateFaces, "associate-faces", "", false, "Associate Faces")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionCompareFaces, "compare-faces", "", false, "Compare Faces")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionCopyProjectVersion, "copy-project-version", "", false, "Copy Project Version")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionCreateCollection, "create-collection", "", false, "Create Collection")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionCreateDataset, "create-dataset", "", false, "Create Dataset")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionCreateFaceLivenessSession, "create-face-liveness-session", "", false, "Create Face Liveness Session")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionCreateProject, "create-project", "", false, "Create Project")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionCreateProjectVersion, "create-project-version", "", false, "Create Project Version")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionCreateStreamProcessor, "create-stream-processor", "", false, "Create Stream Processor")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionCreateUser, "create-user", "", false, "Create User")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionDeleteCollection, "delete-collection", "", false, "Delete Collection")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionDeleteDataset, "delete-dataset", "", false, "Delete Dataset")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionDeleteFaces, "delete-faces", "", false, "Delete Faces")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionDeleteProject, "delete-project", "", false, "Delete Project")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionDeleteProjectPolicy, "delete-project-policy", "", false, "Delete Project Policy")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionDeleteProjectVersion, "delete-project-version", "", false, "Delete Project Version")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionDeleteStreamProcessor, "delete-stream-processor", "", false, "Delete Stream Processor")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionDeleteUser, "delete-user", "", false, "Delete User")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionDescribeCollection, "describe-collection", "", false, "Describe Collection")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionDescribeDataset, "describe-dataset", "", false, "Describe Dataset")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionDescribeProjectVersions, "describe-project-versions", "", false, "Describe Project Versions")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionDescribeProjects, "describe-projects", "", false, "Describe Projects")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionDescribeStreamProcessor, "describe-stream-processor", "", false, "Describe Stream Processor")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionDetectCustomLabels, "detect-custom-labels", "", false, "Detect Custom Labels")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionDetectFaces, "detect-faces", "", false, "Detect Faces")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionDetectLabels, "detect-labels", "", false, "Detect Labels")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionDetectModerationLabels, "detect-moderation-labels", "", false, "Detect Moderation Labels")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionDetectProtectiveEquipment, "detect-protective-equipment", "", false, "Detect Protective Equipment")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionDetectText, "detect-text", "", false, "Detect Text")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionDisassociateFaces, "disassociate-faces", "", false, "Disassociate Faces")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionDistributeDatasetEntries, "distribute-dataset-entries", "", false, "Distribute Dataset Entries")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionGetCelebrityInfo, "get-celebrity-info", "", false, "Get Celebrity Info")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionGetCelebrityRecognition, "get-celebrity-recognition", "", false, "Get Celebrity Recognition")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionGetContentModeration, "get-content-moderation", "", false, "Get Content Moderation")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionGetFaceDetection, "get-face-detection", "", false, "Get Face Detection")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionGetFaceLivenessSessionResults, "get-face-liveness-session-results", "", false, "Get Face Liveness Session Results")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionGetFaceSearch, "get-face-search", "", false, "Get Face Search")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionGetLabelDetection, "get-label-detection", "", false, "Get Label Detection")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionGetMediaAnalysisJob, "get-media-analysis-job", "", false, "Get Media Analysis Job")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionGetPersonTracking, "get-person-tracking", "", false, "Get Person Tracking")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionGetSegmentDetection, "get-segment-detection", "", false, "Get Segment Detection")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionGetTextDetection, "get-text-detection", "", false, "Get Text Detection")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionIndexFaces, "index-faces", "", false, "Index Faces")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionListCollections, "list-collections", "", false, "List Collections")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionListDatasetEntries, "list-dataset-entries", "", false, "List Dataset Entries")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionListDatasetLabels, "list-dataset-labels", "", false, "List Dataset Labels")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionListFaces, "list-faces", "", false, "List Faces")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionListMediaAnalysisJobs, "list-media-analysis-jobs", "", false, "List Media Analysis Jobs")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionListProjectPolicies, "list-project-policies", "", false, "List Project Policies")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionListStreamProcessors, "list-stream-processors", "", false, "List Stream Processors")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionListUsers, "list-users", "", false, "List Users")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionPutProjectPolicy, "put-project-policy", "", false, "Put Project Policy")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionRecognizeCelebrities, "recognize-celebrities", "", false, "Recognize Celebrities")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionSearchFaces, "search-faces", "", false, "Search Faces")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionSearchFacesByImage, "search-faces-by-image", "", false, "Search Faces By Image")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionSearchUsers, "search-users", "", false, "Search Users")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionSearchUsersByImage, "search-users-by-image", "", false, "Search Users By Image")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionStartCelebrityRecognition, "start-celebrity-recognition", "", false, "Start Celebrity Recognition")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionStartContentModeration, "start-content-moderation", "", false, "Start Content Moderation")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionStartFaceDetection, "start-face-detection", "", false, "Start Face Detection")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionStartFaceSearch, "start-face-search", "", false, "Start Face Search")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionStartLabelDetection, "start-label-detection", "", false, "Start Label Detection")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionStartMediaAnalysisJob, "start-media-analysis-job", "", false, "Start Media Analysis Job")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionStartPersonTracking, "start-person-tracking", "", false, "Start Person Tracking")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionStartProjectVersion, "start-project-version", "", false, "Start Project Version")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionStartSegmentDetection, "start-segment-detection", "", false, "Start Segment Detection")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionStartStreamProcessor, "start-stream-processor", "", false, "Start Stream Processor")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionStartTextDetection, "start-text-detection", "", false, "Start Text Detection")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionStopProjectVersion, "stop-project-version", "", false, "Stop Project Version")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionStopStreamProcessor, "stop-stream-processor", "", false, "Stop Stream Processor")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionTagResource, "tag-resource", "", false, "Tag Resource")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionUntagResource, "untag-resource", "", false, "Untag Resource")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionUpdateDatasetEntries, "update-dataset-entries", "", false, "Update Dataset Entries")
	_rekognitionCmd.Flags().BoolVarP(&_rekognitionUpdateStreamProcessor, "update-stream-processor", "", false, "Update Stream Processor")

}
