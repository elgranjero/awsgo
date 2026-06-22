package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/comprehend"
)

var fields_batch_detect_dominant_language = []leanruntime.Field{
	{Name: "TextList", Flag: "text-list", Type: "[]string", Required: true},
}

var fields_batch_detect_entities = []leanruntime.Field{
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "TextList", Flag: "text-list", Type: "[]string", Required: true},
}

var fields_batch_detect_key_phrases = []leanruntime.Field{
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "TextList", Flag: "text-list", Type: "[]string", Required: true},
}

var fields_batch_detect_sentiment = []leanruntime.Field{
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "TextList", Flag: "text-list", Type: "[]string", Required: true},
}

var fields_batch_detect_syntax = []leanruntime.Field{
	{Name: "LanguageCode", Flag: "language-code", Type: "types.SyntaxLanguageCode", Required: true},
	{Name: "TextList", Flag: "text-list", Type: "[]string", Required: true},
}

var fields_batch_detect_targeted_sentiment = []leanruntime.Field{
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "TextList", Flag: "text-list", Type: "[]string", Required: true},
}

var fields_classify_document = []leanruntime.Field{
	{Name: "Bytes", Flag: "bytes", Type: "[]byte", Required: false},
	{Name: "DocumentReaderConfig", Flag: "document-reader-config", Type: "*types.DocumentReaderConfig", Required: false},
	{Name: "EndpointArn", Flag: "endpoint-arn", Type: "*string", Required: true},
	{Name: "Text", Flag: "text", Type: "*string", Required: false},
}

var fields_contains_pii_entities = []leanruntime.Field{
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "Text", Flag: "text", Type: "*string", Required: true},
}

var fields_create_dataset = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DatasetName", Flag: "dataset-name", Type: "*string", Required: true},
	{Name: "DatasetType", Flag: "dataset-type", Type: "types.DatasetType", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FlywheelArn", Flag: "flywheel-arn", Type: "*string", Required: true},
	{Name: "InputDataConfig", Flag: "input-data-config", Type: "*types.DatasetInputDataConfig", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_document_classifier = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: true},
	{Name: "DocumentClassifierName", Flag: "document-classifier-name", Type: "*string", Required: true},
	{Name: "InputDataConfig", Flag: "input-data-config", Type: "*types.DocumentClassifierInputDataConfig", Required: true},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "Mode", Flag: "mode", Type: "types.DocumentClassifierMode", Required: false},
	{Name: "ModelKmsKeyId", Flag: "model-kms-key-id", Type: "*string", Required: false},
	{Name: "ModelPolicy", Flag: "model-policy", Type: "*string", Required: false},
	{Name: "OutputDataConfig", Flag: "output-data-config", Type: "*types.DocumentClassifierOutputDataConfig", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VersionName", Flag: "version-name", Type: "*string", Required: false},
	{Name: "VolumeKmsKeyId", Flag: "volume-kms-key-id", Type: "*string", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: false},
}

var fields_create_endpoint = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: false},
	{Name: "DesiredInferenceUnits", Flag: "desired-inference-units", Type: "*int32", Required: true},
	{Name: "EndpointName", Flag: "endpoint-name", Type: "*string", Required: true},
	{Name: "FlywheelArn", Flag: "flywheel-arn", Type: "*string", Required: false},
	{Name: "ModelArn", Flag: "model-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_entity_recognizer = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: true},
	{Name: "InputDataConfig", Flag: "input-data-config", Type: "*types.EntityRecognizerInputDataConfig", Required: true},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "ModelKmsKeyId", Flag: "model-kms-key-id", Type: "*string", Required: false},
	{Name: "ModelPolicy", Flag: "model-policy", Type: "*string", Required: false},
	{Name: "RecognizerName", Flag: "recognizer-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VersionName", Flag: "version-name", Type: "*string", Required: false},
	{Name: "VolumeKmsKeyId", Flag: "volume-kms-key-id", Type: "*string", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: false},
}

var fields_create_flywheel = []leanruntime.Field{
	{Name: "ActiveModelArn", Flag: "active-model-arn", Type: "*string", Required: false},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: true},
	{Name: "DataLakeS3Uri", Flag: "data-lake-s3-uri", Type: "*string", Required: true},
	{Name: "DataSecurityConfig", Flag: "data-security-config", Type: "*types.DataSecurityConfig", Required: false},
	{Name: "FlywheelName", Flag: "flywheel-name", Type: "*string", Required: true},
	{Name: "ModelType", Flag: "model-type", Type: "types.ModelType", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TaskConfig", Flag: "task-config", Type: "*types.TaskConfig", Required: false},
}

var fields_delete_document_classifier = []leanruntime.Field{
	{Name: "DocumentClassifierArn", Flag: "document-classifier-arn", Type: "*string", Required: true},
}

var fields_delete_endpoint = []leanruntime.Field{
	{Name: "EndpointArn", Flag: "endpoint-arn", Type: "*string", Required: true},
}

var fields_delete_entity_recognizer = []leanruntime.Field{
	{Name: "EntityRecognizerArn", Flag: "entity-recognizer-arn", Type: "*string", Required: true},
}

var fields_delete_flywheel = []leanruntime.Field{
	{Name: "FlywheelArn", Flag: "flywheel-arn", Type: "*string", Required: true},
}

var fields_delete_resource_policy = []leanruntime.Field{
	{Name: "PolicyRevisionId", Flag: "policy-revision-id", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_describe_dataset = []leanruntime.Field{
	{Name: "DatasetArn", Flag: "dataset-arn", Type: "*string", Required: true},
}

var fields_describe_document_classification_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_describe_document_classifier = []leanruntime.Field{
	{Name: "DocumentClassifierArn", Flag: "document-classifier-arn", Type: "*string", Required: true},
}

var fields_describe_dominant_language_detection_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_describe_endpoint = []leanruntime.Field{
	{Name: "EndpointArn", Flag: "endpoint-arn", Type: "*string", Required: true},
}

var fields_describe_entities_detection_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_describe_entity_recognizer = []leanruntime.Field{
	{Name: "EntityRecognizerArn", Flag: "entity-recognizer-arn", Type: "*string", Required: true},
}

var fields_describe_events_detection_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_describe_flywheel = []leanruntime.Field{
	{Name: "FlywheelArn", Flag: "flywheel-arn", Type: "*string", Required: true},
}

var fields_describe_flywheel_iteration = []leanruntime.Field{
	{Name: "FlywheelArn", Flag: "flywheel-arn", Type: "*string", Required: true},
	{Name: "FlywheelIterationId", Flag: "flywheel-iteration-id", Type: "*string", Required: true},
}

var fields_describe_key_phrases_detection_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_describe_pii_entities_detection_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_describe_resource_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_describe_sentiment_detection_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_describe_targeted_sentiment_detection_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_describe_topics_detection_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_detect_dominant_language = []leanruntime.Field{
	{Name: "Text", Flag: "text", Type: "*string", Required: true},
}

var fields_detect_entities = []leanruntime.Field{
	{Name: "Bytes", Flag: "bytes", Type: "[]byte", Required: false},
	{Name: "DocumentReaderConfig", Flag: "document-reader-config", Type: "*types.DocumentReaderConfig", Required: false},
	{Name: "EndpointArn", Flag: "endpoint-arn", Type: "*string", Required: false},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: false},
	{Name: "Text", Flag: "text", Type: "*string", Required: false},
}

var fields_detect_key_phrases = []leanruntime.Field{
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "Text", Flag: "text", Type: "*string", Required: true},
}

var fields_detect_pii_entities = []leanruntime.Field{
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "Text", Flag: "text", Type: "*string", Required: true},
}

var fields_detect_sentiment = []leanruntime.Field{
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "Text", Flag: "text", Type: "*string", Required: true},
}

var fields_detect_syntax = []leanruntime.Field{
	{Name: "LanguageCode", Flag: "language-code", Type: "types.SyntaxLanguageCode", Required: true},
	{Name: "Text", Flag: "text", Type: "*string", Required: true},
}

var fields_detect_targeted_sentiment = []leanruntime.Field{
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "Text", Flag: "text", Type: "*string", Required: true},
}

var fields_detect_toxic_content = []leanruntime.Field{
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "TextSegments", Flag: "text-segments", Type: "[]types.TextSegment", Required: true},
}

var fields_import_model = []leanruntime.Field{
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: false},
	{Name: "ModelKmsKeyId", Flag: "model-kms-key-id", Type: "*string", Required: false},
	{Name: "ModelName", Flag: "model-name", Type: "*string", Required: false},
	{Name: "SourceModelArn", Flag: "source-model-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VersionName", Flag: "version-name", Type: "*string", Required: false},
}

var fields_list_datasets = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.DatasetFilter", Required: false},
	{Name: "FlywheelArn", Flag: "flywheel-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_document_classification_jobs = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.DocumentClassificationJobFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_document_classifier_summaries = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_document_classifiers = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.DocumentClassifierFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_dominant_language_detection_jobs = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.DominantLanguageDetectionJobFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_endpoints = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.EndpointFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_entities_detection_jobs = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.EntitiesDetectionJobFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_entity_recognizer_summaries = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_entity_recognizers = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.EntityRecognizerFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_events_detection_jobs = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.EventsDetectionJobFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_flywheel_iteration_history = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.FlywheelIterationFilter", Required: false},
	{Name: "FlywheelArn", Flag: "flywheel-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_flywheels = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.FlywheelFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_key_phrases_detection_jobs = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.KeyPhrasesDetectionJobFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_pii_entities_detection_jobs = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.PiiEntitiesDetectionJobFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_sentiment_detection_jobs = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.SentimentDetectionJobFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_targeted_sentiment_detection_jobs = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.TargetedSentimentDetectionJobFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_topics_detection_jobs = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.TopicsDetectionJobFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_put_resource_policy = []leanruntime.Field{
	{Name: "PolicyRevisionId", Flag: "policy-revision-id", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "ResourcePolicy", Flag: "resource-policy", Type: "*string", Required: true},
}

var fields_start_document_classification_job = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: true},
	{Name: "DocumentClassifierArn", Flag: "document-classifier-arn", Type: "*string", Required: false},
	{Name: "FlywheelArn", Flag: "flywheel-arn", Type: "*string", Required: false},
	{Name: "InputDataConfig", Flag: "input-data-config", Type: "*types.InputDataConfig", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: false},
	{Name: "OutputDataConfig", Flag: "output-data-config", Type: "*types.OutputDataConfig", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VolumeKmsKeyId", Flag: "volume-kms-key-id", Type: "*string", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: false},
}

var fields_start_dominant_language_detection_job = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: true},
	{Name: "InputDataConfig", Flag: "input-data-config", Type: "*types.InputDataConfig", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: false},
	{Name: "OutputDataConfig", Flag: "output-data-config", Type: "*types.OutputDataConfig", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VolumeKmsKeyId", Flag: "volume-kms-key-id", Type: "*string", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: false},
}

var fields_start_entities_detection_job = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: true},
	{Name: "EntityRecognizerArn", Flag: "entity-recognizer-arn", Type: "*string", Required: false},
	{Name: "FlywheelArn", Flag: "flywheel-arn", Type: "*string", Required: false},
	{Name: "InputDataConfig", Flag: "input-data-config", Type: "*types.InputDataConfig", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: false},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "OutputDataConfig", Flag: "output-data-config", Type: "*types.OutputDataConfig", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VolumeKmsKeyId", Flag: "volume-kms-key-id", Type: "*string", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: false},
}

var fields_start_events_detection_job = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: true},
	{Name: "InputDataConfig", Flag: "input-data-config", Type: "*types.InputDataConfig", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: false},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "OutputDataConfig", Flag: "output-data-config", Type: "*types.OutputDataConfig", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetEventTypes", Flag: "target-event-types", Type: "[]string", Required: true},
}

var fields_start_flywheel_iteration = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "FlywheelArn", Flag: "flywheel-arn", Type: "*string", Required: true},
}

var fields_start_key_phrases_detection_job = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: true},
	{Name: "InputDataConfig", Flag: "input-data-config", Type: "*types.InputDataConfig", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: false},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "OutputDataConfig", Flag: "output-data-config", Type: "*types.OutputDataConfig", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VolumeKmsKeyId", Flag: "volume-kms-key-id", Type: "*string", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: false},
}

var fields_start_pii_entities_detection_job = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: true},
	{Name: "InputDataConfig", Flag: "input-data-config", Type: "*types.InputDataConfig", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: false},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "Mode", Flag: "mode", Type: "types.PiiEntitiesDetectionMode", Required: true},
	{Name: "OutputDataConfig", Flag: "output-data-config", Type: "*types.OutputDataConfig", Required: true},
	{Name: "RedactionConfig", Flag: "redaction-config", Type: "*types.RedactionConfig", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_start_sentiment_detection_job = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: true},
	{Name: "InputDataConfig", Flag: "input-data-config", Type: "*types.InputDataConfig", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: false},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "OutputDataConfig", Flag: "output-data-config", Type: "*types.OutputDataConfig", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VolumeKmsKeyId", Flag: "volume-kms-key-id", Type: "*string", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: false},
}

var fields_start_targeted_sentiment_detection_job = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: true},
	{Name: "InputDataConfig", Flag: "input-data-config", Type: "*types.InputDataConfig", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: false},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "OutputDataConfig", Flag: "output-data-config", Type: "*types.OutputDataConfig", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VolumeKmsKeyId", Flag: "volume-kms-key-id", Type: "*string", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: false},
}

var fields_start_topics_detection_job = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: true},
	{Name: "InputDataConfig", Flag: "input-data-config", Type: "*types.InputDataConfig", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: false},
	{Name: "NumberOfTopics", Flag: "number-of-topics", Type: "*int32", Required: false},
	{Name: "OutputDataConfig", Flag: "output-data-config", Type: "*types.OutputDataConfig", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VolumeKmsKeyId", Flag: "volume-kms-key-id", Type: "*string", Required: false},
	{Name: "VpcConfig", Flag: "vpc-config", Type: "*types.VpcConfig", Required: false},
}

var fields_stop_dominant_language_detection_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_stop_entities_detection_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_stop_events_detection_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_stop_key_phrases_detection_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_stop_pii_entities_detection_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_stop_sentiment_detection_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_stop_targeted_sentiment_detection_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_stop_training_document_classifier = []leanruntime.Field{
	{Name: "DocumentClassifierArn", Flag: "document-classifier-arn", Type: "*string", Required: true},
}

var fields_stop_training_entity_recognizer = []leanruntime.Field{
	{Name: "EntityRecognizerArn", Flag: "entity-recognizer-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_endpoint = []leanruntime.Field{
	{Name: "DesiredDataAccessRoleArn", Flag: "desired-data-access-role-arn", Type: "*string", Required: false},
	{Name: "DesiredInferenceUnits", Flag: "desired-inference-units", Type: "*int32", Required: false},
	{Name: "DesiredModelArn", Flag: "desired-model-arn", Type: "*string", Required: false},
	{Name: "EndpointArn", Flag: "endpoint-arn", Type: "*string", Required: true},
	{Name: "FlywheelArn", Flag: "flywheel-arn", Type: "*string", Required: false},
}

var fields_update_flywheel = []leanruntime.Field{
	{Name: "ActiveModelArn", Flag: "active-model-arn", Type: "*string", Required: false},
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: false},
	{Name: "DataSecurityConfig", Flag: "data-security-config", Type: "*types.UpdateDataSecurityConfig", Required: false},
	{Name: "FlywheelArn", Flag: "flywheel-arn", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-detect-dominant-language": {
			Name:   "batch-detect-dominant-language",
			Fields: fields_batch_detect_dominant_language,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDetectDominantLanguageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_detect_dominant_language, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDetectDominantLanguage(ctx, input)
			},
		},
		"batch-detect-entities": {
			Name:   "batch-detect-entities",
			Fields: fields_batch_detect_entities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDetectEntitiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_detect_entities, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDetectEntities(ctx, input)
			},
		},
		"batch-detect-key-phrases": {
			Name:   "batch-detect-key-phrases",
			Fields: fields_batch_detect_key_phrases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDetectKeyPhrasesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_detect_key_phrases, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDetectKeyPhrases(ctx, input)
			},
		},
		"batch-detect-sentiment": {
			Name:   "batch-detect-sentiment",
			Fields: fields_batch_detect_sentiment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDetectSentimentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_detect_sentiment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDetectSentiment(ctx, input)
			},
		},
		"batch-detect-syntax": {
			Name:   "batch-detect-syntax",
			Fields: fields_batch_detect_syntax,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDetectSyntaxInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_detect_syntax, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDetectSyntax(ctx, input)
			},
		},
		"batch-detect-targeted-sentiment": {
			Name:   "batch-detect-targeted-sentiment",
			Fields: fields_batch_detect_targeted_sentiment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDetectTargetedSentimentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_detect_targeted_sentiment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDetectTargetedSentiment(ctx, input)
			},
		},
		"classify-document": {
			Name:   "classify-document",
			Fields: fields_classify_document,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ClassifyDocumentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_classify_document, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ClassifyDocument(ctx, input)
			},
		},
		"contains-pii-entities": {
			Name:   "contains-pii-entities",
			Fields: fields_contains_pii_entities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ContainsPiiEntitiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_contains_pii_entities, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ContainsPiiEntities(ctx, input)
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
		"create-document-classifier": {
			Name:   "create-document-classifier",
			Fields: fields_create_document_classifier,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDocumentClassifierInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_document_classifier, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDocumentClassifier(ctx, input)
			},
		},
		"create-endpoint": {
			Name:   "create-endpoint",
			Fields: fields_create_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEndpoint(ctx, input)
			},
		},
		"create-entity-recognizer": {
			Name:   "create-entity-recognizer",
			Fields: fields_create_entity_recognizer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEntityRecognizerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_entity_recognizer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEntityRecognizer(ctx, input)
			},
		},
		"create-flywheel": {
			Name:   "create-flywheel",
			Fields: fields_create_flywheel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFlywheelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_flywheel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFlywheel(ctx, input)
			},
		},
		"delete-document-classifier": {
			Name:   "delete-document-classifier",
			Fields: fields_delete_document_classifier,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDocumentClassifierInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_document_classifier, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDocumentClassifier(ctx, input)
			},
		},
		"delete-endpoint": {
			Name:   "delete-endpoint",
			Fields: fields_delete_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEndpoint(ctx, input)
			},
		},
		"delete-entity-recognizer": {
			Name:   "delete-entity-recognizer",
			Fields: fields_delete_entity_recognizer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEntityRecognizerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_entity_recognizer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEntityRecognizer(ctx, input)
			},
		},
		"delete-flywheel": {
			Name:   "delete-flywheel",
			Fields: fields_delete_flywheel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFlywheelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_flywheel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFlywheel(ctx, input)
			},
		},
		"delete-resource-policy": {
			Name:   "delete-resource-policy",
			Fields: fields_delete_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourcePolicy(ctx, input)
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
		"describe-document-classification-job": {
			Name:   "describe-document-classification-job",
			Fields: fields_describe_document_classification_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDocumentClassificationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_document_classification_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDocumentClassificationJob(ctx, input)
			},
		},
		"describe-document-classifier": {
			Name:   "describe-document-classifier",
			Fields: fields_describe_document_classifier,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDocumentClassifierInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_document_classifier, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDocumentClassifier(ctx, input)
			},
		},
		"describe-dominant-language-detection-job": {
			Name:   "describe-dominant-language-detection-job",
			Fields: fields_describe_dominant_language_detection_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDominantLanguageDetectionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_dominant_language_detection_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDominantLanguageDetectionJob(ctx, input)
			},
		},
		"describe-endpoint": {
			Name:   "describe-endpoint",
			Fields: fields_describe_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEndpoint(ctx, input)
			},
		},
		"describe-entities-detection-job": {
			Name:   "describe-entities-detection-job",
			Fields: fields_describe_entities_detection_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEntitiesDetectionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_entities_detection_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEntitiesDetectionJob(ctx, input)
			},
		},
		"describe-entity-recognizer": {
			Name:   "describe-entity-recognizer",
			Fields: fields_describe_entity_recognizer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEntityRecognizerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_entity_recognizer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEntityRecognizer(ctx, input)
			},
		},
		"describe-events-detection-job": {
			Name:   "describe-events-detection-job",
			Fields: fields_describe_events_detection_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEventsDetectionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_events_detection_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEventsDetectionJob(ctx, input)
			},
		},
		"describe-flywheel": {
			Name:   "describe-flywheel",
			Fields: fields_describe_flywheel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFlywheelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_flywheel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFlywheel(ctx, input)
			},
		},
		"describe-flywheel-iteration": {
			Name:   "describe-flywheel-iteration",
			Fields: fields_describe_flywheel_iteration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFlywheelIterationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_flywheel_iteration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFlywheelIteration(ctx, input)
			},
		},
		"describe-key-phrases-detection-job": {
			Name:   "describe-key-phrases-detection-job",
			Fields: fields_describe_key_phrases_detection_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeKeyPhrasesDetectionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_key_phrases_detection_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeKeyPhrasesDetectionJob(ctx, input)
			},
		},
		"describe-pii-entities-detection-job": {
			Name:   "describe-pii-entities-detection-job",
			Fields: fields_describe_pii_entities_detection_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePiiEntitiesDetectionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_pii_entities_detection_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePiiEntitiesDetectionJob(ctx, input)
			},
		},
		"describe-resource-policy": {
			Name:   "describe-resource-policy",
			Fields: fields_describe_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeResourcePolicy(ctx, input)
			},
		},
		"describe-sentiment-detection-job": {
			Name:   "describe-sentiment-detection-job",
			Fields: fields_describe_sentiment_detection_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSentimentDetectionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_sentiment_detection_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSentimentDetectionJob(ctx, input)
			},
		},
		"describe-targeted-sentiment-detection-job": {
			Name:   "describe-targeted-sentiment-detection-job",
			Fields: fields_describe_targeted_sentiment_detection_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTargetedSentimentDetectionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_targeted_sentiment_detection_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTargetedSentimentDetectionJob(ctx, input)
			},
		},
		"describe-topics-detection-job": {
			Name:   "describe-topics-detection-job",
			Fields: fields_describe_topics_detection_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTopicsDetectionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_topics_detection_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTopicsDetectionJob(ctx, input)
			},
		},
		"detect-dominant-language": {
			Name:   "detect-dominant-language",
			Fields: fields_detect_dominant_language,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetectDominantLanguageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detect_dominant_language, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetectDominantLanguage(ctx, input)
			},
		},
		"detect-entities": {
			Name:   "detect-entities",
			Fields: fields_detect_entities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetectEntitiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detect_entities, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetectEntities(ctx, input)
			},
		},
		"detect-key-phrases": {
			Name:   "detect-key-phrases",
			Fields: fields_detect_key_phrases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetectKeyPhrasesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detect_key_phrases, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetectKeyPhrases(ctx, input)
			},
		},
		"detect-pii-entities": {
			Name:   "detect-pii-entities",
			Fields: fields_detect_pii_entities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetectPiiEntitiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detect_pii_entities, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetectPiiEntities(ctx, input)
			},
		},
		"detect-sentiment": {
			Name:   "detect-sentiment",
			Fields: fields_detect_sentiment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetectSentimentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detect_sentiment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetectSentiment(ctx, input)
			},
		},
		"detect-syntax": {
			Name:   "detect-syntax",
			Fields: fields_detect_syntax,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetectSyntaxInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detect_syntax, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetectSyntax(ctx, input)
			},
		},
		"detect-targeted-sentiment": {
			Name:   "detect-targeted-sentiment",
			Fields: fields_detect_targeted_sentiment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetectTargetedSentimentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detect_targeted_sentiment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetectTargetedSentiment(ctx, input)
			},
		},
		"detect-toxic-content": {
			Name:   "detect-toxic-content",
			Fields: fields_detect_toxic_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetectToxicContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detect_toxic_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetectToxicContent(ctx, input)
			},
		},
		"import-model": {
			Name:   "import-model",
			Fields: fields_import_model,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportModelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_model, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportModel(ctx, input)
			},
		},
		"list-datasets": {
			Name:   "list-datasets",
			Fields: fields_list_datasets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDatasetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_datasets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDatasets(ctx, input)
				}
				var results []*svc.ListDatasetsOutput
				p := svc.NewListDatasetsPaginator(client, input)
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
		"list-document-classification-jobs": {
			Name:   "list-document-classification-jobs",
			Fields: fields_list_document_classification_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDocumentClassificationJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_document_classification_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDocumentClassificationJobs(ctx, input)
				}
				var results []*svc.ListDocumentClassificationJobsOutput
				p := svc.NewListDocumentClassificationJobsPaginator(client, input)
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
		"list-document-classifier-summaries": {
			Name:   "list-document-classifier-summaries",
			Fields: fields_list_document_classifier_summaries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDocumentClassifierSummariesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_document_classifier_summaries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDocumentClassifierSummaries(ctx, input)
				}
				var results []*svc.ListDocumentClassifierSummariesOutput
				p := svc.NewListDocumentClassifierSummariesPaginator(client, input)
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
		"list-document-classifiers": {
			Name:   "list-document-classifiers",
			Fields: fields_list_document_classifiers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDocumentClassifiersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_document_classifiers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDocumentClassifiers(ctx, input)
				}
				var results []*svc.ListDocumentClassifiersOutput
				p := svc.NewListDocumentClassifiersPaginator(client, input)
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
		"list-dominant-language-detection-jobs": {
			Name:   "list-dominant-language-detection-jobs",
			Fields: fields_list_dominant_language_detection_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDominantLanguageDetectionJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_dominant_language_detection_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDominantLanguageDetectionJobs(ctx, input)
				}
				var results []*svc.ListDominantLanguageDetectionJobsOutput
				p := svc.NewListDominantLanguageDetectionJobsPaginator(client, input)
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
		"list-endpoints": {
			Name:   "list-endpoints",
			Fields: fields_list_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEndpointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_endpoints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEndpoints(ctx, input)
				}
				var results []*svc.ListEndpointsOutput
				p := svc.NewListEndpointsPaginator(client, input)
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
		"list-entities-detection-jobs": {
			Name:   "list-entities-detection-jobs",
			Fields: fields_list_entities_detection_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEntitiesDetectionJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_entities_detection_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEntitiesDetectionJobs(ctx, input)
				}
				var results []*svc.ListEntitiesDetectionJobsOutput
				p := svc.NewListEntitiesDetectionJobsPaginator(client, input)
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
		"list-entity-recognizer-summaries": {
			Name:   "list-entity-recognizer-summaries",
			Fields: fields_list_entity_recognizer_summaries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEntityRecognizerSummariesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_entity_recognizer_summaries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEntityRecognizerSummaries(ctx, input)
				}
				var results []*svc.ListEntityRecognizerSummariesOutput
				p := svc.NewListEntityRecognizerSummariesPaginator(client, input)
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
		"list-entity-recognizers": {
			Name:   "list-entity-recognizers",
			Fields: fields_list_entity_recognizers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEntityRecognizersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_entity_recognizers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEntityRecognizers(ctx, input)
				}
				var results []*svc.ListEntityRecognizersOutput
				p := svc.NewListEntityRecognizersPaginator(client, input)
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
		"list-events-detection-jobs": {
			Name:   "list-events-detection-jobs",
			Fields: fields_list_events_detection_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEventsDetectionJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_events_detection_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEventsDetectionJobs(ctx, input)
				}
				var results []*svc.ListEventsDetectionJobsOutput
				p := svc.NewListEventsDetectionJobsPaginator(client, input)
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
		"list-flywheel-iteration-history": {
			Name:   "list-flywheel-iteration-history",
			Fields: fields_list_flywheel_iteration_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFlywheelIterationHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_flywheel_iteration_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFlywheelIterationHistory(ctx, input)
				}
				var results []*svc.ListFlywheelIterationHistoryOutput
				p := svc.NewListFlywheelIterationHistoryPaginator(client, input)
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
		"list-flywheels": {
			Name:   "list-flywheels",
			Fields: fields_list_flywheels,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFlywheelsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_flywheels, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFlywheels(ctx, input)
				}
				var results []*svc.ListFlywheelsOutput
				p := svc.NewListFlywheelsPaginator(client, input)
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
		"list-key-phrases-detection-jobs": {
			Name:   "list-key-phrases-detection-jobs",
			Fields: fields_list_key_phrases_detection_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListKeyPhrasesDetectionJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_key_phrases_detection_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListKeyPhrasesDetectionJobs(ctx, input)
				}
				var results []*svc.ListKeyPhrasesDetectionJobsOutput
				p := svc.NewListKeyPhrasesDetectionJobsPaginator(client, input)
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
		"list-pii-entities-detection-jobs": {
			Name:   "list-pii-entities-detection-jobs",
			Fields: fields_list_pii_entities_detection_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPiiEntitiesDetectionJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_pii_entities_detection_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPiiEntitiesDetectionJobs(ctx, input)
				}
				var results []*svc.ListPiiEntitiesDetectionJobsOutput
				p := svc.NewListPiiEntitiesDetectionJobsPaginator(client, input)
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
		"list-sentiment-detection-jobs": {
			Name:   "list-sentiment-detection-jobs",
			Fields: fields_list_sentiment_detection_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSentimentDetectionJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sentiment_detection_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSentimentDetectionJobs(ctx, input)
				}
				var results []*svc.ListSentimentDetectionJobsOutput
				p := svc.NewListSentimentDetectionJobsPaginator(client, input)
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
		"list-targeted-sentiment-detection-jobs": {
			Name:   "list-targeted-sentiment-detection-jobs",
			Fields: fields_list_targeted_sentiment_detection_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTargetedSentimentDetectionJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_targeted_sentiment_detection_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTargetedSentimentDetectionJobs(ctx, input)
				}
				var results []*svc.ListTargetedSentimentDetectionJobsOutput
				p := svc.NewListTargetedSentimentDetectionJobsPaginator(client, input)
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
		"list-topics-detection-jobs": {
			Name:   "list-topics-detection-jobs",
			Fields: fields_list_topics_detection_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTopicsDetectionJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_topics_detection_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTopicsDetectionJobs(ctx, input)
				}
				var results []*svc.ListTopicsDetectionJobsOutput
				p := svc.NewListTopicsDetectionJobsPaginator(client, input)
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
		"put-resource-policy": {
			Name:   "put-resource-policy",
			Fields: fields_put_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutResourcePolicy(ctx, input)
			},
		},
		"start-document-classification-job": {
			Name:   "start-document-classification-job",
			Fields: fields_start_document_classification_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDocumentClassificationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_document_classification_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDocumentClassificationJob(ctx, input)
			},
		},
		"start-dominant-language-detection-job": {
			Name:   "start-dominant-language-detection-job",
			Fields: fields_start_dominant_language_detection_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDominantLanguageDetectionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_dominant_language_detection_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDominantLanguageDetectionJob(ctx, input)
			},
		},
		"start-entities-detection-job": {
			Name:   "start-entities-detection-job",
			Fields: fields_start_entities_detection_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartEntitiesDetectionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_entities_detection_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartEntitiesDetectionJob(ctx, input)
			},
		},
		"start-events-detection-job": {
			Name:   "start-events-detection-job",
			Fields: fields_start_events_detection_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartEventsDetectionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_events_detection_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartEventsDetectionJob(ctx, input)
			},
		},
		"start-flywheel-iteration": {
			Name:   "start-flywheel-iteration",
			Fields: fields_start_flywheel_iteration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartFlywheelIterationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_flywheel_iteration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartFlywheelIteration(ctx, input)
			},
		},
		"start-key-phrases-detection-job": {
			Name:   "start-key-phrases-detection-job",
			Fields: fields_start_key_phrases_detection_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartKeyPhrasesDetectionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_key_phrases_detection_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartKeyPhrasesDetectionJob(ctx, input)
			},
		},
		"start-pii-entities-detection-job": {
			Name:   "start-pii-entities-detection-job",
			Fields: fields_start_pii_entities_detection_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartPiiEntitiesDetectionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_pii_entities_detection_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartPiiEntitiesDetectionJob(ctx, input)
			},
		},
		"start-sentiment-detection-job": {
			Name:   "start-sentiment-detection-job",
			Fields: fields_start_sentiment_detection_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSentimentDetectionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_sentiment_detection_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSentimentDetectionJob(ctx, input)
			},
		},
		"start-targeted-sentiment-detection-job": {
			Name:   "start-targeted-sentiment-detection-job",
			Fields: fields_start_targeted_sentiment_detection_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartTargetedSentimentDetectionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_targeted_sentiment_detection_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartTargetedSentimentDetectionJob(ctx, input)
			},
		},
		"start-topics-detection-job": {
			Name:   "start-topics-detection-job",
			Fields: fields_start_topics_detection_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartTopicsDetectionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_topics_detection_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartTopicsDetectionJob(ctx, input)
			},
		},
		"stop-dominant-language-detection-job": {
			Name:   "stop-dominant-language-detection-job",
			Fields: fields_stop_dominant_language_detection_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopDominantLanguageDetectionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_dominant_language_detection_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopDominantLanguageDetectionJob(ctx, input)
			},
		},
		"stop-entities-detection-job": {
			Name:   "stop-entities-detection-job",
			Fields: fields_stop_entities_detection_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopEntitiesDetectionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_entities_detection_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopEntitiesDetectionJob(ctx, input)
			},
		},
		"stop-events-detection-job": {
			Name:   "stop-events-detection-job",
			Fields: fields_stop_events_detection_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopEventsDetectionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_events_detection_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopEventsDetectionJob(ctx, input)
			},
		},
		"stop-key-phrases-detection-job": {
			Name:   "stop-key-phrases-detection-job",
			Fields: fields_stop_key_phrases_detection_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopKeyPhrasesDetectionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_key_phrases_detection_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopKeyPhrasesDetectionJob(ctx, input)
			},
		},
		"stop-pii-entities-detection-job": {
			Name:   "stop-pii-entities-detection-job",
			Fields: fields_stop_pii_entities_detection_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopPiiEntitiesDetectionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_pii_entities_detection_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopPiiEntitiesDetectionJob(ctx, input)
			},
		},
		"stop-sentiment-detection-job": {
			Name:   "stop-sentiment-detection-job",
			Fields: fields_stop_sentiment_detection_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopSentimentDetectionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_sentiment_detection_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopSentimentDetectionJob(ctx, input)
			},
		},
		"stop-targeted-sentiment-detection-job": {
			Name:   "stop-targeted-sentiment-detection-job",
			Fields: fields_stop_targeted_sentiment_detection_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopTargetedSentimentDetectionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_targeted_sentiment_detection_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopTargetedSentimentDetectionJob(ctx, input)
			},
		},
		"stop-training-document-classifier": {
			Name:   "stop-training-document-classifier",
			Fields: fields_stop_training_document_classifier,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopTrainingDocumentClassifierInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_training_document_classifier, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopTrainingDocumentClassifier(ctx, input)
			},
		},
		"stop-training-entity-recognizer": {
			Name:   "stop-training-entity-recognizer",
			Fields: fields_stop_training_entity_recognizer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopTrainingEntityRecognizerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_training_entity_recognizer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopTrainingEntityRecognizer(ctx, input)
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
		"update-endpoint": {
			Name:   "update-endpoint",
			Fields: fields_update_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEndpoint(ctx, input)
			},
		},
		"update-flywheel": {
			Name:   "update-flywheel",
			Fields: fields_update_flywheel,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFlywheelInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_flywheel, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFlywheel(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("comprehend", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
