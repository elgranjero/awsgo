package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/rekognition"
)

var fields_associate_faces = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "CollectionId", Flag: "collection-id", Type: "*string", Required: true},
	{Name: "FaceIds", Flag: "face-ids", Type: "[]string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
	{Name: "UserMatchThreshold", Flag: "user-match-threshold", Type: "*float32", Required: false},
}

var fields_compare_faces = []leanruntime.Field{
	{Name: "QualityFilter", Flag: "quality-filter", Type: "types.QualityFilter", Required: false},
	{Name: "SimilarityThreshold", Flag: "similarity-threshold", Type: "*float32", Required: false},
	{Name: "SourceImage", Flag: "source-image", Type: "*types.Image", Required: true},
	{Name: "TargetImage", Flag: "target-image", Type: "*types.Image", Required: true},
}

var fields_copy_project_version = []leanruntime.Field{
	{Name: "DestinationProjectArn", Flag: "destination-project-arn", Type: "*string", Required: true},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "OutputConfig", Flag: "output-config", Type: "*types.OutputConfig", Required: true},
	{Name: "SourceProjectArn", Flag: "source-project-arn", Type: "*string", Required: true},
	{Name: "SourceProjectVersionArn", Flag: "source-project-version-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VersionName", Flag: "version-name", Type: "*string", Required: true},
}

var fields_create_collection = []leanruntime.Field{
	{Name: "CollectionId", Flag: "collection-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_dataset = []leanruntime.Field{
	{Name: "DatasetSource", Flag: "dataset-source", Type: "*types.DatasetSource", Required: false},
	{Name: "DatasetType", Flag: "dataset-type", Type: "types.DatasetType", Required: true},
	{Name: "ProjectArn", Flag: "project-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_face_liveness_session = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "Settings", Flag: "settings", Type: "*types.CreateFaceLivenessSessionRequestSettings", Required: false},
}

var fields_create_project = []leanruntime.Field{
	{Name: "AutoUpdate", Flag: "auto-update", Type: "types.ProjectAutoUpdate", Required: false},
	{Name: "Feature", Flag: "feature", Type: "types.CustomizationFeature", Required: false},
	{Name: "ProjectName", Flag: "project-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_project_version = []leanruntime.Field{
	{Name: "FeatureConfig", Flag: "feature-config", Type: "*types.CustomizationFeatureConfig", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "OutputConfig", Flag: "output-config", Type: "*types.OutputConfig", Required: true},
	{Name: "ProjectArn", Flag: "project-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TestingData", Flag: "testing-data", Type: "*types.TestingData", Required: false},
	{Name: "TrainingData", Flag: "training-data", Type: "*types.TrainingData", Required: false},
	{Name: "VersionDescription", Flag: "version-description", Type: "*string", Required: false},
	{Name: "VersionName", Flag: "version-name", Type: "*string", Required: true},
}

var fields_create_stream_processor = []leanruntime.Field{
	{Name: "DataSharingPreference", Flag: "data-sharing-preference", Type: "*types.StreamProcessorDataSharingPreference", Required: false},
	{Name: "Input", Flag: "input", Type: "*types.StreamProcessorInput", Required: true},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NotificationChannel", Flag: "notification-channel", Type: "*types.StreamProcessorNotificationChannel", Required: false},
	{Name: "Output", Flag: "output", Type: "*types.StreamProcessorOutput", Required: true},
	{Name: "RegionsOfInterest", Flag: "regions-of-interest", Type: "[]types.RegionOfInterest", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Settings", Flag: "settings", Type: "*types.StreamProcessorSettings", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_user = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "CollectionId", Flag: "collection-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_delete_collection = []leanruntime.Field{
	{Name: "CollectionId", Flag: "collection-id", Type: "*string", Required: true},
}

var fields_delete_dataset = []leanruntime.Field{
	{Name: "DatasetArn", Flag: "dataset-arn", Type: "*string", Required: true},
}

var fields_delete_faces = []leanruntime.Field{
	{Name: "CollectionId", Flag: "collection-id", Type: "*string", Required: true},
	{Name: "FaceIds", Flag: "face-ids", Type: "[]string", Required: true},
}

var fields_delete_project = []leanruntime.Field{
	{Name: "ProjectArn", Flag: "project-arn", Type: "*string", Required: true},
}

var fields_delete_project_policy = []leanruntime.Field{
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "PolicyRevisionId", Flag: "policy-revision-id", Type: "*string", Required: false},
	{Name: "ProjectArn", Flag: "project-arn", Type: "*string", Required: true},
}

var fields_delete_project_version = []leanruntime.Field{
	{Name: "ProjectVersionArn", Flag: "project-version-arn", Type: "*string", Required: true},
}

var fields_delete_stream_processor = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_user = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "CollectionId", Flag: "collection-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_describe_collection = []leanruntime.Field{
	{Name: "CollectionId", Flag: "collection-id", Type: "*string", Required: true},
}

var fields_describe_dataset = []leanruntime.Field{
	{Name: "DatasetArn", Flag: "dataset-arn", Type: "*string", Required: true},
}

var fields_describe_project_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProjectArn", Flag: "project-arn", Type: "*string", Required: true},
	{Name: "VersionNames", Flag: "version-names", Type: "[]string", Required: false},
}

var fields_describe_projects = []leanruntime.Field{
	{Name: "Features", Flag: "features", Type: "[]types.CustomizationFeature", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProjectNames", Flag: "project-names", Type: "[]string", Required: false},
}

var fields_describe_stream_processor = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_detect_custom_labels = []leanruntime.Field{
	{Name: "Image", Flag: "image", Type: "*types.Image", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MinConfidence", Flag: "min-confidence", Type: "*float32", Required: false},
	{Name: "ProjectVersionArn", Flag: "project-version-arn", Type: "*string", Required: true},
}

var fields_detect_faces = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "[]types.Attribute", Required: false},
	{Name: "Image", Flag: "image", Type: "*types.Image", Required: true},
}

var fields_detect_labels = []leanruntime.Field{
	{Name: "Features", Flag: "features", Type: "[]types.DetectLabelsFeatureName", Required: false},
	{Name: "Image", Flag: "image", Type: "*types.Image", Required: true},
	{Name: "MaxLabels", Flag: "max-labels", Type: "*int32", Required: false},
	{Name: "MinConfidence", Flag: "min-confidence", Type: "*float32", Required: false},
	{Name: "Settings", Flag: "settings", Type: "*types.DetectLabelsSettings", Required: false},
}

var fields_detect_moderation_labels = []leanruntime.Field{
	{Name: "HumanLoopConfig", Flag: "human-loop-config", Type: "*types.HumanLoopConfig", Required: false},
	{Name: "Image", Flag: "image", Type: "*types.Image", Required: true},
	{Name: "MinConfidence", Flag: "min-confidence", Type: "*float32", Required: false},
	{Name: "ProjectVersion", Flag: "project-version", Type: "*string", Required: false},
}

var fields_detect_protective_equipment = []leanruntime.Field{
	{Name: "Image", Flag: "image", Type: "*types.Image", Required: true},
	{Name: "SummarizationAttributes", Flag: "summarization-attributes", Type: "*types.ProtectiveEquipmentSummarizationAttributes", Required: false},
}

var fields_detect_text = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.DetectTextFilters", Required: false},
	{Name: "Image", Flag: "image", Type: "*types.Image", Required: true},
}

var fields_disassociate_faces = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "CollectionId", Flag: "collection-id", Type: "*string", Required: true},
	{Name: "FaceIds", Flag: "face-ids", Type: "[]string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_distribute_dataset_entries = []leanruntime.Field{
	{Name: "Datasets", Flag: "datasets", Type: "[]types.DistributeDataset", Required: true},
}

var fields_get_celebrity_info = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_celebrity_recognition = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.CelebrityRecognitionSortBy", Required: false},
}

var fields_get_content_moderation = []leanruntime.Field{
	{Name: "AggregateBy", Flag: "aggregate-by", Type: "types.ContentModerationAggregateBy", Required: false},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ContentModerationSortBy", Required: false},
}

var fields_get_face_detection = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_face_liveness_session_results = []leanruntime.Field{
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_get_face_search = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.FaceSearchSortBy", Required: false},
}

var fields_get_label_detection = []leanruntime.Field{
	{Name: "AggregateBy", Flag: "aggregate-by", Type: "types.LabelDetectionAggregateBy", Required: false},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.LabelDetectionSortBy", Required: false},
}

var fields_get_media_analysis_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_get_person_tracking = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.PersonTrackingSortBy", Required: false},
}

var fields_get_segment_detection = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_text_detection = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_index_faces = []leanruntime.Field{
	{Name: "CollectionId", Flag: "collection-id", Type: "*string", Required: true},
	{Name: "DetectionAttributes", Flag: "detection-attributes", Type: "[]types.Attribute", Required: false},
	{Name: "ExternalImageId", Flag: "external-image-id", Type: "*string", Required: false},
	{Name: "Image", Flag: "image", Type: "*types.Image", Required: true},
	{Name: "MaxFaces", Flag: "max-faces", Type: "*int32", Required: false},
	{Name: "QualityFilter", Flag: "quality-filter", Type: "types.QualityFilter", Required: false},
}

var fields_list_collections = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_dataset_entries = []leanruntime.Field{
	{Name: "ContainsLabels", Flag: "contains-labels", Type: "[]string", Required: false},
	{Name: "DatasetArn", Flag: "dataset-arn", Type: "*string", Required: true},
	{Name: "HasErrors", Flag: "has-errors", Type: "*bool", Required: false},
	{Name: "Labeled", Flag: "labeled", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SourceRefContains", Flag: "source-ref-contains", Type: "*string", Required: false},
}

var fields_list_dataset_labels = []leanruntime.Field{
	{Name: "DatasetArn", Flag: "dataset-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_faces = []leanruntime.Field{
	{Name: "CollectionId", Flag: "collection-id", Type: "*string", Required: true},
	{Name: "FaceIds", Flag: "face-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
}

var fields_list_media_analysis_jobs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_project_policies = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProjectArn", Flag: "project-arn", Type: "*string", Required: true},
}

var fields_list_stream_processors = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_users = []leanruntime.Field{
	{Name: "CollectionId", Flag: "collection-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_put_project_policy = []leanruntime.Field{
	{Name: "PolicyDocument", Flag: "policy-document", Type: "*string", Required: true},
	{Name: "PolicyName", Flag: "policy-name", Type: "*string", Required: true},
	{Name: "PolicyRevisionId", Flag: "policy-revision-id", Type: "*string", Required: false},
	{Name: "ProjectArn", Flag: "project-arn", Type: "*string", Required: true},
}

var fields_recognize_celebrities = []leanruntime.Field{
	{Name: "Image", Flag: "image", Type: "*types.Image", Required: true},
}

var fields_search_faces = []leanruntime.Field{
	{Name: "CollectionId", Flag: "collection-id", Type: "*string", Required: true},
	{Name: "FaceId", Flag: "face-id", Type: "*string", Required: true},
	{Name: "FaceMatchThreshold", Flag: "face-match-threshold", Type: "*float32", Required: false},
	{Name: "MaxFaces", Flag: "max-faces", Type: "*int32", Required: false},
}

var fields_search_faces_by_image = []leanruntime.Field{
	{Name: "CollectionId", Flag: "collection-id", Type: "*string", Required: true},
	{Name: "FaceMatchThreshold", Flag: "face-match-threshold", Type: "*float32", Required: false},
	{Name: "Image", Flag: "image", Type: "*types.Image", Required: true},
	{Name: "MaxFaces", Flag: "max-faces", Type: "*int32", Required: false},
	{Name: "QualityFilter", Flag: "quality-filter", Type: "types.QualityFilter", Required: false},
}

var fields_search_users = []leanruntime.Field{
	{Name: "CollectionId", Flag: "collection-id", Type: "*string", Required: true},
	{Name: "FaceId", Flag: "face-id", Type: "*string", Required: false},
	{Name: "MaxUsers", Flag: "max-users", Type: "*int32", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
	{Name: "UserMatchThreshold", Flag: "user-match-threshold", Type: "*float32", Required: false},
}

var fields_search_users_by_image = []leanruntime.Field{
	{Name: "CollectionId", Flag: "collection-id", Type: "*string", Required: true},
	{Name: "Image", Flag: "image", Type: "*types.Image", Required: true},
	{Name: "MaxUsers", Flag: "max-users", Type: "*int32", Required: false},
	{Name: "QualityFilter", Flag: "quality-filter", Type: "types.QualityFilter", Required: false},
	{Name: "UserMatchThreshold", Flag: "user-match-threshold", Type: "*float32", Required: false},
}

var fields_start_celebrity_recognition = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "JobTag", Flag: "job-tag", Type: "*string", Required: false},
	{Name: "NotificationChannel", Flag: "notification-channel", Type: "*types.NotificationChannel", Required: false},
	{Name: "Video", Flag: "video", Type: "*types.Video", Required: true},
}

var fields_start_content_moderation = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "JobTag", Flag: "job-tag", Type: "*string", Required: false},
	{Name: "MinConfidence", Flag: "min-confidence", Type: "*float32", Required: false},
	{Name: "NotificationChannel", Flag: "notification-channel", Type: "*types.NotificationChannel", Required: false},
	{Name: "Video", Flag: "video", Type: "*types.Video", Required: true},
}

var fields_start_face_detection = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "FaceAttributes", Flag: "face-attributes", Type: "types.FaceAttributes", Required: false},
	{Name: "JobTag", Flag: "job-tag", Type: "*string", Required: false},
	{Name: "NotificationChannel", Flag: "notification-channel", Type: "*types.NotificationChannel", Required: false},
	{Name: "Video", Flag: "video", Type: "*types.Video", Required: true},
}

var fields_start_face_search = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "CollectionId", Flag: "collection-id", Type: "*string", Required: true},
	{Name: "FaceMatchThreshold", Flag: "face-match-threshold", Type: "*float32", Required: false},
	{Name: "JobTag", Flag: "job-tag", Type: "*string", Required: false},
	{Name: "NotificationChannel", Flag: "notification-channel", Type: "*types.NotificationChannel", Required: false},
	{Name: "Video", Flag: "video", Type: "*types.Video", Required: true},
}

var fields_start_label_detection = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Features", Flag: "features", Type: "[]types.LabelDetectionFeatureName", Required: false},
	{Name: "JobTag", Flag: "job-tag", Type: "*string", Required: false},
	{Name: "MinConfidence", Flag: "min-confidence", Type: "*float32", Required: false},
	{Name: "NotificationChannel", Flag: "notification-channel", Type: "*types.NotificationChannel", Required: false},
	{Name: "Settings", Flag: "settings", Type: "*types.LabelDetectionSettings", Required: false},
	{Name: "Video", Flag: "video", Type: "*types.Video", Required: true},
}

var fields_start_media_analysis_job = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Input", Flag: "input", Type: "*types.MediaAnalysisInput", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "OperationsConfig", Flag: "operations-config", Type: "*types.MediaAnalysisOperationsConfig", Required: true},
	{Name: "OutputConfig", Flag: "output-config", Type: "*types.MediaAnalysisOutputConfig", Required: true},
}

var fields_start_person_tracking = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "JobTag", Flag: "job-tag", Type: "*string", Required: false},
	{Name: "NotificationChannel", Flag: "notification-channel", Type: "*types.NotificationChannel", Required: false},
	{Name: "Video", Flag: "video", Type: "*types.Video", Required: true},
}

var fields_start_project_version = []leanruntime.Field{
	{Name: "MaxInferenceUnits", Flag: "max-inference-units", Type: "*int32", Required: false},
	{Name: "MinInferenceUnits", Flag: "min-inference-units", Type: "*int32", Required: true},
	{Name: "ProjectVersionArn", Flag: "project-version-arn", Type: "*string", Required: true},
}

var fields_start_segment_detection = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "*types.StartSegmentDetectionFilters", Required: false},
	{Name: "JobTag", Flag: "job-tag", Type: "*string", Required: false},
	{Name: "NotificationChannel", Flag: "notification-channel", Type: "*types.NotificationChannel", Required: false},
	{Name: "SegmentTypes", Flag: "segment-types", Type: "[]types.SegmentType", Required: true},
	{Name: "Video", Flag: "video", Type: "*types.Video", Required: true},
}

var fields_start_stream_processor = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "StartSelector", Flag: "start-selector", Type: "*types.StreamProcessingStartSelector", Required: false},
	{Name: "StopSelector", Flag: "stop-selector", Type: "*types.StreamProcessingStopSelector", Required: false},
}

var fields_start_text_detection = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "*types.StartTextDetectionFilters", Required: false},
	{Name: "JobTag", Flag: "job-tag", Type: "*string", Required: false},
	{Name: "NotificationChannel", Flag: "notification-channel", Type: "*types.NotificationChannel", Required: false},
	{Name: "Video", Flag: "video", Type: "*types.Video", Required: true},
}

var fields_stop_project_version = []leanruntime.Field{
	{Name: "ProjectVersionArn", Flag: "project-version-arn", Type: "*string", Required: true},
}

var fields_stop_stream_processor = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_dataset_entries = []leanruntime.Field{
	{Name: "Changes", Flag: "changes", Type: "*types.DatasetChanges", Required: true},
	{Name: "DatasetArn", Flag: "dataset-arn", Type: "*string", Required: true},
}

var fields_update_stream_processor = []leanruntime.Field{
	{Name: "DataSharingPreferenceForUpdate", Flag: "data-sharing-preference-for-update", Type: "*types.StreamProcessorDataSharingPreference", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ParametersToDelete", Flag: "parameters-to-delete", Type: "[]types.StreamProcessorParameterToDelete", Required: false},
	{Name: "RegionsOfInterestForUpdate", Flag: "regions-of-interest-for-update", Type: "[]types.RegionOfInterest", Required: false},
	{Name: "SettingsForUpdate", Flag: "settings-for-update", Type: "*types.StreamProcessorSettingsForUpdate", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-faces": {
			Name:   "associate-faces",
			Fields: fields_associate_faces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateFacesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_faces, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateFaces(ctx, input)
			},
		},
		"compare-faces": {
			Name:   "compare-faces",
			Fields: fields_compare_faces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CompareFacesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_compare_faces, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CompareFaces(ctx, input)
			},
		},
		"copy-project-version": {
			Name:   "copy-project-version",
			Fields: fields_copy_project_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CopyProjectVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_copy_project_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CopyProjectVersion(ctx, input)
			},
		},
		"create-collection": {
			Name:   "create-collection",
			Fields: fields_create_collection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCollectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_collection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCollection(ctx, input)
			},
		},
		"create-dataset": {
			Name:   "create-dataset",
			Fields: fields_create_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataset(ctx, input)
			},
		},
		"create-face-liveness-session": {
			Name:   "create-face-liveness-session",
			Fields: fields_create_face_liveness_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFaceLivenessSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_face_liveness_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFaceLivenessSession(ctx, input)
			},
		},
		"create-project": {
			Name:   "create-project",
			Fields: fields_create_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProject(ctx, input)
			},
		},
		"create-project-version": {
			Name:   "create-project-version",
			Fields: fields_create_project_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProjectVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_project_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProjectVersion(ctx, input)
			},
		},
		"create-stream-processor": {
			Name:   "create-stream-processor",
			Fields: fields_create_stream_processor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStreamProcessorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_stream_processor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStreamProcessor(ctx, input)
			},
		},
		"create-user": {
			Name:   "create-user",
			Fields: fields_create_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUser(ctx, input)
			},
		},
		"delete-collection": {
			Name:   "delete-collection",
			Fields: fields_delete_collection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCollectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_collection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCollection(ctx, input)
			},
		},
		"delete-dataset": {
			Name:   "delete-dataset",
			Fields: fields_delete_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataset(ctx, input)
			},
		},
		"delete-faces": {
			Name:   "delete-faces",
			Fields: fields_delete_faces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFacesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_faces, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFaces(ctx, input)
			},
		},
		"delete-project": {
			Name:   "delete-project",
			Fields: fields_delete_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProject(ctx, input)
			},
		},
		"delete-project-policy": {
			Name:   "delete-project-policy",
			Fields: fields_delete_project_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProjectPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_project_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProjectPolicy(ctx, input)
			},
		},
		"delete-project-version": {
			Name:   "delete-project-version",
			Fields: fields_delete_project_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProjectVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_project_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProjectVersion(ctx, input)
			},
		},
		"delete-stream-processor": {
			Name:   "delete-stream-processor",
			Fields: fields_delete_stream_processor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteStreamProcessorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_stream_processor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteStreamProcessor(ctx, input)
			},
		},
		"delete-user": {
			Name:   "delete-user",
			Fields: fields_delete_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUser(ctx, input)
			},
		},
		"describe-collection": {
			Name:   "describe-collection",
			Fields: fields_describe_collection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCollectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_collection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCollection(ctx, input)
			},
		},
		"describe-dataset": {
			Name:   "describe-dataset",
			Fields: fields_describe_dataset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDatasetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_dataset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDataset(ctx, input)
			},
		},
		"describe-project-versions": {
			Name:   "describe-project-versions",
			Fields: fields_describe_project_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProjectVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_project_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeProjectVersions(ctx, input)
				}
				var results []*svc.DescribeProjectVersionsOutput
				p := svc.NewDescribeProjectVersionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-projects": {
			Name:   "describe-projects",
			Fields: fields_describe_projects,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProjectsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_projects, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeProjects(ctx, input)
				}
				var results []*svc.DescribeProjectsOutput
				p := svc.NewDescribeProjectsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-stream-processor": {
			Name:   "describe-stream-processor",
			Fields: fields_describe_stream_processor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStreamProcessorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_stream_processor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeStreamProcessor(ctx, input)
			},
		},
		"detect-custom-labels": {
			Name:   "detect-custom-labels",
			Fields: fields_detect_custom_labels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetectCustomLabelsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detect_custom_labels, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetectCustomLabels(ctx, input)
			},
		},
		"detect-faces": {
			Name:   "detect-faces",
			Fields: fields_detect_faces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetectFacesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detect_faces, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetectFaces(ctx, input)
			},
		},
		"detect-labels": {
			Name:   "detect-labels",
			Fields: fields_detect_labels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetectLabelsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detect_labels, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetectLabels(ctx, input)
			},
		},
		"detect-moderation-labels": {
			Name:   "detect-moderation-labels",
			Fields: fields_detect_moderation_labels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetectModerationLabelsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detect_moderation_labels, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetectModerationLabels(ctx, input)
			},
		},
		"detect-protective-equipment": {
			Name:   "detect-protective-equipment",
			Fields: fields_detect_protective_equipment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetectProtectiveEquipmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detect_protective_equipment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetectProtectiveEquipment(ctx, input)
			},
		},
		"detect-text": {
			Name:   "detect-text",
			Fields: fields_detect_text,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetectTextInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detect_text, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetectText(ctx, input)
			},
		},
		"disassociate-faces": {
			Name:   "disassociate-faces",
			Fields: fields_disassociate_faces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateFacesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_faces, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateFaces(ctx, input)
			},
		},
		"distribute-dataset-entries": {
			Name:   "distribute-dataset-entries",
			Fields: fields_distribute_dataset_entries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DistributeDatasetEntriesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_distribute_dataset_entries, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DistributeDatasetEntries(ctx, input)
			},
		},
		"get-celebrity-info": {
			Name:   "get-celebrity-info",
			Fields: fields_get_celebrity_info,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCelebrityInfoInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_celebrity_info, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCelebrityInfo(ctx, input)
			},
		},
		"get-celebrity-recognition": {
			Name:   "get-celebrity-recognition",
			Fields: fields_get_celebrity_recognition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCelebrityRecognitionInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_celebrity_recognition, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetCelebrityRecognition(ctx, input)
				}
				var results []*svc.GetCelebrityRecognitionOutput
				p := svc.NewGetCelebrityRecognitionPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-content-moderation": {
			Name:   "get-content-moderation",
			Fields: fields_get_content_moderation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetContentModerationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_content_moderation, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetContentModeration(ctx, input)
				}
				var results []*svc.GetContentModerationOutput
				p := svc.NewGetContentModerationPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-face-detection": {
			Name:   "get-face-detection",
			Fields: fields_get_face_detection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFaceDetectionInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_face_detection, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetFaceDetection(ctx, input)
				}
				var results []*svc.GetFaceDetectionOutput
				p := svc.NewGetFaceDetectionPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-face-liveness-session-results": {
			Name:   "get-face-liveness-session-results",
			Fields: fields_get_face_liveness_session_results,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFaceLivenessSessionResultsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_face_liveness_session_results, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFaceLivenessSessionResults(ctx, input)
			},
		},
		"get-face-search": {
			Name:   "get-face-search",
			Fields: fields_get_face_search,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFaceSearchInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_face_search, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetFaceSearch(ctx, input)
				}
				var results []*svc.GetFaceSearchOutput
				p := svc.NewGetFaceSearchPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-label-detection": {
			Name:   "get-label-detection",
			Fields: fields_get_label_detection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLabelDetectionInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_label_detection, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetLabelDetection(ctx, input)
				}
				var results []*svc.GetLabelDetectionOutput
				p := svc.NewGetLabelDetectionPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-media-analysis-job": {
			Name:   "get-media-analysis-job",
			Fields: fields_get_media_analysis_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMediaAnalysisJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_media_analysis_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMediaAnalysisJob(ctx, input)
			},
		},
		"get-person-tracking": {
			Name:   "get-person-tracking",
			Fields: fields_get_person_tracking,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPersonTrackingInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_person_tracking, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetPersonTracking(ctx, input)
				}
				var results []*svc.GetPersonTrackingOutput
				p := svc.NewGetPersonTrackingPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-segment-detection": {
			Name:   "get-segment-detection",
			Fields: fields_get_segment_detection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSegmentDetectionInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_segment_detection, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetSegmentDetection(ctx, input)
				}
				var results []*svc.GetSegmentDetectionOutput
				p := svc.NewGetSegmentDetectionPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-text-detection": {
			Name:   "get-text-detection",
			Fields: fields_get_text_detection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTextDetectionInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_text_detection, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetTextDetection(ctx, input)
				}
				var results []*svc.GetTextDetectionOutput
				p := svc.NewGetTextDetectionPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"index-faces": {
			Name:   "index-faces",
			Fields: fields_index_faces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.IndexFacesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_index_faces, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.IndexFaces(ctx, input)
			},
		},
		"list-collections": {
			Name:   "list-collections",
			Fields: fields_list_collections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCollectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_collections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCollections(ctx, input)
				}
				var results []*svc.ListCollectionsOutput
				p := svc.NewListCollectionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-dataset-entries": {
			Name:   "list-dataset-entries",
			Fields: fields_list_dataset_entries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDatasetEntriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_dataset_entries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDatasetEntries(ctx, input)
				}
				var results []*svc.ListDatasetEntriesOutput
				p := svc.NewListDatasetEntriesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-dataset-labels": {
			Name:   "list-dataset-labels",
			Fields: fields_list_dataset_labels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDatasetLabelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_dataset_labels, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDatasetLabels(ctx, input)
				}
				var results []*svc.ListDatasetLabelsOutput
				p := svc.NewListDatasetLabelsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-faces": {
			Name:   "list-faces",
			Fields: fields_list_faces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFacesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_faces, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFaces(ctx, input)
				}
				var results []*svc.ListFacesOutput
				p := svc.NewListFacesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-media-analysis-jobs": {
			Name:   "list-media-analysis-jobs",
			Fields: fields_list_media_analysis_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMediaAnalysisJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_media_analysis_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMediaAnalysisJobs(ctx, input)
				}
				var results []*svc.ListMediaAnalysisJobsOutput
				p := svc.NewListMediaAnalysisJobsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-project-policies": {
			Name:   "list-project-policies",
			Fields: fields_list_project_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProjectPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_project_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProjectPolicies(ctx, input)
				}
				var results []*svc.ListProjectPoliciesOutput
				p := svc.NewListProjectPoliciesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-stream-processors": {
			Name:   "list-stream-processors",
			Fields: fields_list_stream_processors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStreamProcessorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_stream_processors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStreamProcessors(ctx, input)
				}
				var results []*svc.ListStreamProcessorsOutput
				p := svc.NewListStreamProcessorsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-tags-for-resource": {
			Name:   "list-tags-for-resource",
			Fields: fields_list_tags_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForResource(ctx, input)
			},
		},
		"list-users": {
			Name:   "list-users",
			Fields: fields_list_users,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUsersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_users, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUsers(ctx, input)
				}
				var results []*svc.ListUsersOutput
				p := svc.NewListUsersPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"put-project-policy": {
			Name:   "put-project-policy",
			Fields: fields_put_project_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutProjectPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_project_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutProjectPolicy(ctx, input)
			},
		},
		"recognize-celebrities": {
			Name:   "recognize-celebrities",
			Fields: fields_recognize_celebrities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RecognizeCelebritiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_recognize_celebrities, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RecognizeCelebrities(ctx, input)
			},
		},
		"search-faces": {
			Name:   "search-faces",
			Fields: fields_search_faces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchFacesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_search_faces, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SearchFaces(ctx, input)
			},
		},
		"search-faces-by-image": {
			Name:   "search-faces-by-image",
			Fields: fields_search_faces_by_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchFacesByImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_search_faces_by_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SearchFacesByImage(ctx, input)
			},
		},
		"search-users": {
			Name:   "search-users",
			Fields: fields_search_users,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchUsersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_search_users, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SearchUsers(ctx, input)
			},
		},
		"search-users-by-image": {
			Name:   "search-users-by-image",
			Fields: fields_search_users_by_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchUsersByImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_search_users_by_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SearchUsersByImage(ctx, input)
			},
		},
		"start-celebrity-recognition": {
			Name:   "start-celebrity-recognition",
			Fields: fields_start_celebrity_recognition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartCelebrityRecognitionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_celebrity_recognition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartCelebrityRecognition(ctx, input)
			},
		},
		"start-content-moderation": {
			Name:   "start-content-moderation",
			Fields: fields_start_content_moderation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartContentModerationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_content_moderation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartContentModeration(ctx, input)
			},
		},
		"start-face-detection": {
			Name:   "start-face-detection",
			Fields: fields_start_face_detection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartFaceDetectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_face_detection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartFaceDetection(ctx, input)
			},
		},
		"start-face-search": {
			Name:   "start-face-search",
			Fields: fields_start_face_search,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartFaceSearchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_face_search, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartFaceSearch(ctx, input)
			},
		},
		"start-label-detection": {
			Name:   "start-label-detection",
			Fields: fields_start_label_detection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartLabelDetectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_label_detection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartLabelDetection(ctx, input)
			},
		},
		"start-media-analysis-job": {
			Name:   "start-media-analysis-job",
			Fields: fields_start_media_analysis_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMediaAnalysisJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_media_analysis_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMediaAnalysisJob(ctx, input)
			},
		},
		"start-person-tracking": {
			Name:   "start-person-tracking",
			Fields: fields_start_person_tracking,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartPersonTrackingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_person_tracking, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartPersonTracking(ctx, input)
			},
		},
		"start-project-version": {
			Name:   "start-project-version",
			Fields: fields_start_project_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartProjectVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_project_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartProjectVersion(ctx, input)
			},
		},
		"start-segment-detection": {
			Name:   "start-segment-detection",
			Fields: fields_start_segment_detection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSegmentDetectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_segment_detection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSegmentDetection(ctx, input)
			},
		},
		"start-stream-processor": {
			Name:   "start-stream-processor",
			Fields: fields_start_stream_processor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartStreamProcessorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_stream_processor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartStreamProcessor(ctx, input)
			},
		},
		"start-text-detection": {
			Name:   "start-text-detection",
			Fields: fields_start_text_detection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartTextDetectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_text_detection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartTextDetection(ctx, input)
			},
		},
		"stop-project-version": {
			Name:   "stop-project-version",
			Fields: fields_stop_project_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopProjectVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_project_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopProjectVersion(ctx, input)
			},
		},
		"stop-stream-processor": {
			Name:   "stop-stream-processor",
			Fields: fields_stop_stream_processor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopStreamProcessorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_stream_processor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopStreamProcessor(ctx, input)
			},
		},
		"tag-resource": {
			Name:   "tag-resource",
			Fields: fields_tag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagResource(ctx, input)
			},
		},
		"untag-resource": {
			Name:   "untag-resource",
			Fields: fields_untag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagResource(ctx, input)
			},
		},
		"update-dataset-entries": {
			Name:   "update-dataset-entries",
			Fields: fields_update_dataset_entries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDatasetEntriesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_dataset_entries, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDatasetEntries(ctx, input)
			},
		},
		"update-stream-processor": {
			Name:   "update-stream-processor",
			Fields: fields_update_stream_processor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateStreamProcessorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_stream_processor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateStreamProcessor(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("rekognition", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
