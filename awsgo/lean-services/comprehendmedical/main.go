package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/comprehendmedical"
)

var fields_describe_entities_detection_v2_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_describe_icd10cm_inference_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_describe_phi_detection_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_describe_rx_norm_inference_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_describe_snomedct_inference_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_detect_entities = []leanruntime.Field{
	{Name: "Text", Flag: "text", Type: "*string", Required: true},
}

var fields_detect_entities_v2 = []leanruntime.Field{
	{Name: "Text", Flag: "text", Type: "*string", Required: true},
}

var fields_detect_phi = []leanruntime.Field{
	{Name: "Text", Flag: "text", Type: "*string", Required: true},
}

var fields_infer_icd10cm = []leanruntime.Field{
	{Name: "Text", Flag: "text", Type: "*string", Required: true},
}

var fields_infer_rx_norm = []leanruntime.Field{
	{Name: "Text", Flag: "text", Type: "*string", Required: true},
}

var fields_infer_snomedct = []leanruntime.Field{
	{Name: "Text", Flag: "text", Type: "*string", Required: true},
}

var fields_list_entities_detection_v2_jobs = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ComprehendMedicalAsyncJobFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_icd10cm_inference_jobs = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ComprehendMedicalAsyncJobFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_phi_detection_jobs = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ComprehendMedicalAsyncJobFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_rx_norm_inference_jobs = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ComprehendMedicalAsyncJobFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_snomedct_inference_jobs = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ComprehendMedicalAsyncJobFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_start_entities_detection_v2_job = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: true},
	{Name: "InputDataConfig", Flag: "input-data-config", Type: "*types.InputDataConfig", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: false},
	{Name: "KMSKey", Flag: "kms-key", Type: "*string", Required: false},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "OutputDataConfig", Flag: "output-data-config", Type: "*types.OutputDataConfig", Required: true},
}

var fields_start_icd10cm_inference_job = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: true},
	{Name: "InputDataConfig", Flag: "input-data-config", Type: "*types.InputDataConfig", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: false},
	{Name: "KMSKey", Flag: "kms-key", Type: "*string", Required: false},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "OutputDataConfig", Flag: "output-data-config", Type: "*types.OutputDataConfig", Required: true},
}

var fields_start_phi_detection_job = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: true},
	{Name: "InputDataConfig", Flag: "input-data-config", Type: "*types.InputDataConfig", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: false},
	{Name: "KMSKey", Flag: "kms-key", Type: "*string", Required: false},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "OutputDataConfig", Flag: "output-data-config", Type: "*types.OutputDataConfig", Required: true},
}

var fields_start_rx_norm_inference_job = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: true},
	{Name: "InputDataConfig", Flag: "input-data-config", Type: "*types.InputDataConfig", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: false},
	{Name: "KMSKey", Flag: "kms-key", Type: "*string", Required: false},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "OutputDataConfig", Flag: "output-data-config", Type: "*types.OutputDataConfig", Required: true},
}

var fields_start_snomedct_inference_job = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: true},
	{Name: "InputDataConfig", Flag: "input-data-config", Type: "*types.InputDataConfig", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: false},
	{Name: "KMSKey", Flag: "kms-key", Type: "*string", Required: false},
	{Name: "LanguageCode", Flag: "language-code", Type: "types.LanguageCode", Required: true},
	{Name: "OutputDataConfig", Flag: "output-data-config", Type: "*types.OutputDataConfig", Required: true},
}

var fields_stop_entities_detection_v2_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_stop_icd10cm_inference_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_stop_phi_detection_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_stop_rx_norm_inference_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_stop_snomedct_inference_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"describe-entities-detection-v2-job": {
			Name:   "describe-entities-detection-v2-job",
			Fields: fields_describe_entities_detection_v2_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEntitiesDetectionV2JobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_entities_detection_v2_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEntitiesDetectionV2Job(ctx, input)
			},
		},
		"describe-icd10cm-inference-job": {
			Name:   "describe-icd10cm-inference-job",
			Fields: fields_describe_icd10cm_inference_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeICD10CMInferenceJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_icd10cm_inference_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeICD10CMInferenceJob(ctx, input)
			},
		},
		"describe-phi-detection-job": {
			Name:   "describe-phi-detection-job",
			Fields: fields_describe_phi_detection_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePHIDetectionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_phi_detection_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePHIDetectionJob(ctx, input)
			},
		},
		"describe-rx-norm-inference-job": {
			Name:   "describe-rx-norm-inference-job",
			Fields: fields_describe_rx_norm_inference_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRxNormInferenceJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_rx_norm_inference_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRxNormInferenceJob(ctx, input)
			},
		},
		"describe-snomedct-inference-job": {
			Name:   "describe-snomedct-inference-job",
			Fields: fields_describe_snomedct_inference_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSNOMEDCTInferenceJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_snomedct_inference_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSNOMEDCTInferenceJob(ctx, input)
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
		"detect-entities-v2": {
			Name:   "detect-entities-v2",
			Fields: fields_detect_entities_v2,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetectEntitiesV2Input{}
				if _, err := leanruntime.ApplyInput(input, fields_detect_entities_v2, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetectEntitiesV2(ctx, input)
			},
		},
		"detect-phi": {
			Name:   "detect-phi",
			Fields: fields_detect_phi,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetectPHIInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detect_phi, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetectPHI(ctx, input)
			},
		},
		"infer-icd10cm": {
			Name:   "infer-icd10cm",
			Fields: fields_infer_icd10cm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InferICD10CMInput{}
				if _, err := leanruntime.ApplyInput(input, fields_infer_icd10cm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InferICD10CM(ctx, input)
			},
		},
		"infer-rx-norm": {
			Name:   "infer-rx-norm",
			Fields: fields_infer_rx_norm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InferRxNormInput{}
				if _, err := leanruntime.ApplyInput(input, fields_infer_rx_norm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InferRxNorm(ctx, input)
			},
		},
		"infer-snomedct": {
			Name:   "infer-snomedct",
			Fields: fields_infer_snomedct,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InferSNOMEDCTInput{}
				if _, err := leanruntime.ApplyInput(input, fields_infer_snomedct, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InferSNOMEDCT(ctx, input)
			},
		},
		"list-entities-detection-v2-jobs": {
			Name:   "list-entities-detection-v2-jobs",
			Fields: fields_list_entities_detection_v2_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEntitiesDetectionV2JobsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_entities_detection_v2_jobs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListEntitiesDetectionV2Jobs(ctx, input)
			},
		},
		"list-icd10cm-inference-jobs": {
			Name:   "list-icd10cm-inference-jobs",
			Fields: fields_list_icd10cm_inference_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListICD10CMInferenceJobsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_icd10cm_inference_jobs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListICD10CMInferenceJobs(ctx, input)
			},
		},
		"list-phi-detection-jobs": {
			Name:   "list-phi-detection-jobs",
			Fields: fields_list_phi_detection_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPHIDetectionJobsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_phi_detection_jobs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListPHIDetectionJobs(ctx, input)
			},
		},
		"list-rx-norm-inference-jobs": {
			Name:   "list-rx-norm-inference-jobs",
			Fields: fields_list_rx_norm_inference_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRxNormInferenceJobsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_rx_norm_inference_jobs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListRxNormInferenceJobs(ctx, input)
			},
		},
		"list-snomedct-inference-jobs": {
			Name:   "list-snomedct-inference-jobs",
			Fields: fields_list_snomedct_inference_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSNOMEDCTInferenceJobsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_snomedct_inference_jobs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListSNOMEDCTInferenceJobs(ctx, input)
			},
		},
		"start-entities-detection-v2-job": {
			Name:   "start-entities-detection-v2-job",
			Fields: fields_start_entities_detection_v2_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartEntitiesDetectionV2JobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_entities_detection_v2_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartEntitiesDetectionV2Job(ctx, input)
			},
		},
		"start-icd10cm-inference-job": {
			Name:   "start-icd10cm-inference-job",
			Fields: fields_start_icd10cm_inference_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartICD10CMInferenceJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_icd10cm_inference_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartICD10CMInferenceJob(ctx, input)
			},
		},
		"start-phi-detection-job": {
			Name:   "start-phi-detection-job",
			Fields: fields_start_phi_detection_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartPHIDetectionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_phi_detection_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartPHIDetectionJob(ctx, input)
			},
		},
		"start-rx-norm-inference-job": {
			Name:   "start-rx-norm-inference-job",
			Fields: fields_start_rx_norm_inference_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartRxNormInferenceJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_rx_norm_inference_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartRxNormInferenceJob(ctx, input)
			},
		},
		"start-snomedct-inference-job": {
			Name:   "start-snomedct-inference-job",
			Fields: fields_start_snomedct_inference_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSNOMEDCTInferenceJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_snomedct_inference_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSNOMEDCTInferenceJob(ctx, input)
			},
		},
		"stop-entities-detection-v2-job": {
			Name:   "stop-entities-detection-v2-job",
			Fields: fields_stop_entities_detection_v2_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopEntitiesDetectionV2JobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_entities_detection_v2_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopEntitiesDetectionV2Job(ctx, input)
			},
		},
		"stop-icd10cm-inference-job": {
			Name:   "stop-icd10cm-inference-job",
			Fields: fields_stop_icd10cm_inference_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopICD10CMInferenceJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_icd10cm_inference_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopICD10CMInferenceJob(ctx, input)
			},
		},
		"stop-phi-detection-job": {
			Name:   "stop-phi-detection-job",
			Fields: fields_stop_phi_detection_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopPHIDetectionJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_phi_detection_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopPHIDetectionJob(ctx, input)
			},
		},
		"stop-rx-norm-inference-job": {
			Name:   "stop-rx-norm-inference-job",
			Fields: fields_stop_rx_norm_inference_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopRxNormInferenceJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_rx_norm_inference_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopRxNormInferenceJob(ctx, input)
			},
		},
		"stop-snomedct-inference-job": {
			Name:   "stop-snomedct-inference-job",
			Fields: fields_stop_snomedct_inference_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopSNOMEDCTInferenceJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_snomedct_inference_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopSNOMEDCTInferenceJob(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("comprehendmedical", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
