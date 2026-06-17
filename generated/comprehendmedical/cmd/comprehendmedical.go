package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/comprehendmedical"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// comprehendmedicalCmd represents the comprehendmedical command
var _comprehendmedicalCmd = &cobra.Command{
	Use:   "comprehendmedical",
	Short: "AWS comprehendmedical CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := comprehendmedical.NewFromConfig(cfg)
		if _comprehendmedicalDescribeEntitiesDetectionV2Job {
			comprehendmedical_DescribeEntitiesDetectionV2Job(cfg, client)
			return
		}
		if _comprehendmedicalDescribeICD10CMInferenceJob {
			comprehendmedical_DescribeICD10CMInferenceJob(cfg, client)
			return
		}
		if _comprehendmedicalDescribePHIDetectionJob {
			comprehendmedical_DescribePHIDetectionJob(cfg, client)
			return
		}
		if _comprehendmedicalDescribeRxNormInferenceJob {
			comprehendmedical_DescribeRxNormInferenceJob(cfg, client)
			return
		}
		if _comprehendmedicalDescribeSNOMEDCTInferenceJob {
			comprehendmedical_DescribeSNOMEDCTInferenceJob(cfg, client)
			return
		}
		if _comprehendmedicalDetectEntities {
			comprehendmedical_DetectEntities(cfg, client)
			return
		}
		if _comprehendmedicalDetectEntitiesV2 {
			comprehendmedical_DetectEntitiesV2(cfg, client)
			return
		}
		if _comprehendmedicalDetectPHI {
			comprehendmedical_DetectPHI(cfg, client)
			return
		}
		if _comprehendmedicalInferICD10CM {
			comprehendmedical_InferICD10CM(cfg, client)
			return
		}
		if _comprehendmedicalInferRxNorm {
			comprehendmedical_InferRxNorm(cfg, client)
			return
		}
		if _comprehendmedicalInferSNOMEDCT {
			comprehendmedical_InferSNOMEDCT(cfg, client)
			return
		}
		if _comprehendmedicalListEntitiesDetectionV2Jobs {
			comprehendmedical_ListEntitiesDetectionV2Jobs(cfg, client)
			return
		}
		if _comprehendmedicalListICD10CMInferenceJobs {
			comprehendmedical_ListICD10CMInferenceJobs(cfg, client)
			return
		}
		if _comprehendmedicalListPHIDetectionJobs {
			comprehendmedical_ListPHIDetectionJobs(cfg, client)
			return
		}
		if _comprehendmedicalListRxNormInferenceJobs {
			comprehendmedical_ListRxNormInferenceJobs(cfg, client)
			return
		}
		if _comprehendmedicalListSNOMEDCTInferenceJobs {
			comprehendmedical_ListSNOMEDCTInferenceJobs(cfg, client)
			return
		}
		if _comprehendmedicalStartEntitiesDetectionV2Job {
			comprehendmedical_StartEntitiesDetectionV2Job(cfg, client)
			return
		}
		if _comprehendmedicalStartICD10CMInferenceJob {
			comprehendmedical_StartICD10CMInferenceJob(cfg, client)
			return
		}
		if _comprehendmedicalStartPHIDetectionJob {
			comprehendmedical_StartPHIDetectionJob(cfg, client)
			return
		}
		if _comprehendmedicalStartRxNormInferenceJob {
			comprehendmedical_StartRxNormInferenceJob(cfg, client)
			return
		}
		if _comprehendmedicalStartSNOMEDCTInferenceJob {
			comprehendmedical_StartSNOMEDCTInferenceJob(cfg, client)
			return
		}
		if _comprehendmedicalStopEntitiesDetectionV2Job {
			comprehendmedical_StopEntitiesDetectionV2Job(cfg, client)
			return
		}
		if _comprehendmedicalStopICD10CMInferenceJob {
			comprehendmedical_StopICD10CMInferenceJob(cfg, client)
			return
		}
		if _comprehendmedicalStopPHIDetectionJob {
			comprehendmedical_StopPHIDetectionJob(cfg, client)
			return
		}
		if _comprehendmedicalStopRxNormInferenceJob {
			comprehendmedical_StopRxNormInferenceJob(cfg, client)
			return
		}
		if _comprehendmedicalStopSNOMEDCTInferenceJob {
			comprehendmedical_StopSNOMEDCTInferenceJob(cfg, client)
			return
		}

	},
}

var (
	_comprehendmedicalDescribeEntitiesDetectionV2Job bool
	_comprehendmedicalDescribeICD10CMInferenceJob    bool
	_comprehendmedicalDescribePHIDetectionJob        bool
	_comprehendmedicalDescribeRxNormInferenceJob     bool
	_comprehendmedicalDescribeSNOMEDCTInferenceJob   bool
	_comprehendmedicalDetectEntities                 bool
	_comprehendmedicalDetectEntitiesV2               bool
	_comprehendmedicalDetectPHI                      bool
	_comprehendmedicalInferICD10CM                   bool
	_comprehendmedicalInferRxNorm                    bool
	_comprehendmedicalInferSNOMEDCT                  bool
	_comprehendmedicalListEntitiesDetectionV2Jobs    bool
	_comprehendmedicalListICD10CMInferenceJobs       bool
	_comprehendmedicalListPHIDetectionJobs           bool
	_comprehendmedicalListRxNormInferenceJobs        bool
	_comprehendmedicalListSNOMEDCTInferenceJobs      bool
	_comprehendmedicalStartEntitiesDetectionV2Job    bool
	_comprehendmedicalStartICD10CMInferenceJob       bool
	_comprehendmedicalStartPHIDetectionJob           bool
	_comprehendmedicalStartRxNormInferenceJob        bool
	_comprehendmedicalStartSNOMEDCTInferenceJob      bool
	_comprehendmedicalStopEntitiesDetectionV2Job     bool
	_comprehendmedicalStopICD10CMInferenceJob        bool
	_comprehendmedicalStopPHIDetectionJob            bool
	_comprehendmedicalStopRxNormInferenceJob         bool
	_comprehendmedicalStopSNOMEDCTInferenceJob       bool

	_comprehendmedicalClientRequestToken string
	_comprehendmedicalDataAccessRoleArn  string
	_comprehendmedicalFilter             string
	_comprehendmedicalInputDataConfig    string
	_comprehendmedicalJobId              string
	_comprehendmedicalJobName            string
	_comprehendmedicalKMSKey             string
	_comprehendmedicalLanguageCode       string
	_comprehendmedicalMaxResults         string
	_comprehendmedicalNextToken          string
	_comprehendmedicalOutputDataConfig   string
	_comprehendmedicalText               string
)

// Gets the properties associated with a medical entities detection job. Use this
// operation to get the status of a detection job.
func comprehendmedical_DescribeEntitiesDetectionV2Job(cfg aws.Config, client *comprehendmedical.Client) {
	input := &comprehendmedical.DescribeEntitiesDetectionV2JobInput{
		// JobId: *string, // Required
	}

	if len(_comprehendmedicalJobId) > 0 {
		input.JobId = aws.String(_comprehendmedicalJobId)
	}

	if resp, err := client.DescribeEntitiesDetectionV2Job(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the properties associated with an InferICD10CM job. Use this operation to
// get the status of an inference job.
func comprehendmedical_DescribeICD10CMInferenceJob(cfg aws.Config, client *comprehendmedical.Client) {
	input := &comprehendmedical.DescribeICD10CMInferenceJobInput{
		// JobId: *string, // Required
	}

	if len(_comprehendmedicalJobId) > 0 {
		input.JobId = aws.String(_comprehendmedicalJobId)
	}

	if resp, err := client.DescribeICD10CMInferenceJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the properties associated with a protected health information (PHI)
// detection job. Use this operation to get the status of a detection job.
func comprehendmedical_DescribePHIDetectionJob(cfg aws.Config, client *comprehendmedical.Client) {
	input := &comprehendmedical.DescribePHIDetectionJobInput{
		// JobId: *string, // Required
	}

	if len(_comprehendmedicalJobId) > 0 {
		input.JobId = aws.String(_comprehendmedicalJobId)
	}

	if resp, err := client.DescribePHIDetectionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the properties associated with an InferRxNorm job. Use this operation to
// get the status of an inference job.
func comprehendmedical_DescribeRxNormInferenceJob(cfg aws.Config, client *comprehendmedical.Client) {
	input := &comprehendmedical.DescribeRxNormInferenceJobInput{
		// JobId: *string, // Required
	}

	if len(_comprehendmedicalJobId) > 0 {
		input.JobId = aws.String(_comprehendmedicalJobId)
	}

	if resp, err := client.DescribeRxNormInferenceJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the properties associated with an InferSNOMEDCT job. Use this operation
// to get the status of an inference job.
func comprehendmedical_DescribeSNOMEDCTInferenceJob(cfg aws.Config, client *comprehendmedical.Client) {
	input := &comprehendmedical.DescribeSNOMEDCTInferenceJobInput{
		// JobId: *string, // Required
	}

	if len(_comprehendmedicalJobId) > 0 {
		input.JobId = aws.String(_comprehendmedicalJobId)
	}

	if resp, err := client.DescribeSNOMEDCTInferenceJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The DetectEntities operation is deprecated. You should use the DetectEntitiesV2 operation
// instead.
//
// Inspects the clinical text for a variety of medical entities and returns
// specific information about them such as entity category, location, and
// confidence score on that information.
//
// Deprecated: This operation is deprecated, use DetectEntitiesV2 instead.
func comprehendmedical_DetectEntities(cfg aws.Config, client *comprehendmedical.Client) {
	input := &comprehendmedical.DetectEntitiesInput{
		// Text: *string, // Required
	}

	if len(_comprehendmedicalText) > 0 {
		input.Text = aws.String(_comprehendmedicalText)
	}

	if resp, err := client.DetectEntities(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Inspects the clinical text for a variety of medical entities and returns
// specific information about them such as entity category, location, and
// confidence score on that information. Amazon Comprehend Medical only detects
// medical entities in English language texts.
//
// The DetectEntitiesV2 operation replaces the DetectEntities operation. This new action uses a
// different model for determining the entities in your medical text and changes
// the way that some entities are returned in the output. You should use the
// DetectEntitiesV2 operation in all new applications.
//
// The DetectEntitiesV2 operation returns the Acuity and Direction entities as
// attributes instead of types.
func comprehendmedical_DetectEntitiesV2(cfg aws.Config, client *comprehendmedical.Client) {
	input := &comprehendmedical.DetectEntitiesV2Input{
		// Text: *string, // Required
	}

	if len(_comprehendmedicalText) > 0 {
		input.Text = aws.String(_comprehendmedicalText)
	}

	if resp, err := client.DetectEntitiesV2(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Inspects the clinical text for protected health information (PHI) entities and
// returns the entity category, location, and confidence score for each entity.
// Amazon Comprehend Medical only detects entities in English language texts.
func comprehendmedical_DetectPHI(cfg aws.Config, client *comprehendmedical.Client) {
	input := &comprehendmedical.DetectPHIInput{
		// Text: *string, // Required
	}

	if len(_comprehendmedicalText) > 0 {
		input.Text = aws.String(_comprehendmedicalText)
	}

	if resp, err := client.DetectPHI(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// InferICD10CM detects medical conditions as entities listed in a patient record
// and links those entities to normalized concept identifiers in the ICD-10-CM
// knowledge base from the Centers for Disease Control. Amazon Comprehend Medical
// only detects medical entities in English language texts.
func comprehendmedical_InferICD10CM(cfg aws.Config, client *comprehendmedical.Client) {
	input := &comprehendmedical.InferICD10CMInput{
		// Text: *string, // Required
	}

	if len(_comprehendmedicalText) > 0 {
		input.Text = aws.String(_comprehendmedicalText)
	}

	if resp, err := client.InferICD10CM(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// InferRxNorm detects medications as entities listed in a patient record and
// links to the normalized concept identifiers in the RxNorm database from the
// National Library of Medicine. Amazon Comprehend Medical only detects medical
// entities in English language texts.
func comprehendmedical_InferRxNorm(cfg aws.Config, client *comprehendmedical.Client) {
	input := &comprehendmedical.InferRxNormInput{
		// Text: *string, // Required
	}

	if len(_comprehendmedicalText) > 0 {
		input.Text = aws.String(_comprehendmedicalText)
	}

	if resp, err := client.InferRxNorm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// InferSNOMEDCT detects possible medical concepts as entities and links them to
// codes from the Systematized Nomenclature of Medicine, Clinical Terms (SNOMED-CT)
// ontology
func comprehendmedical_InferSNOMEDCT(cfg aws.Config, client *comprehendmedical.Client) {
	input := &comprehendmedical.InferSNOMEDCTInput{
		// Text: *string, // Required
	}

	if len(_comprehendmedicalText) > 0 {
		input.Text = aws.String(_comprehendmedicalText)
	}

	if resp, err := client.InferSNOMEDCT(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of medical entity detection jobs that you have submitted.
func comprehendmedical_ListEntitiesDetectionV2Jobs(cfg aws.Config, client *comprehendmedical.Client) {
	input := &comprehendmedical.ListEntitiesDetectionV2JobsInput{}

	if len(_comprehendmedicalFilter) > 0 {
		if err := assignInputField(input, "Filter", _comprehendmedicalFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_comprehendmedicalMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _comprehendmedicalMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_comprehendmedicalNextToken) > 0 {
		input.NextToken = aws.String(_comprehendmedicalNextToken)
	}

	if resp, err := client.ListEntitiesDetectionV2Jobs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of InferICD10CM jobs that you have submitted.
func comprehendmedical_ListICD10CMInferenceJobs(cfg aws.Config, client *comprehendmedical.Client) {
	input := &comprehendmedical.ListICD10CMInferenceJobsInput{}

	if len(_comprehendmedicalFilter) > 0 {
		if err := assignInputField(input, "Filter", _comprehendmedicalFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_comprehendmedicalMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _comprehendmedicalMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_comprehendmedicalNextToken) > 0 {
		input.NextToken = aws.String(_comprehendmedicalNextToken)
	}

	if resp, err := client.ListICD10CMInferenceJobs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of protected health information (PHI) detection jobs you have
// submitted.
func comprehendmedical_ListPHIDetectionJobs(cfg aws.Config, client *comprehendmedical.Client) {
	input := &comprehendmedical.ListPHIDetectionJobsInput{}

	if len(_comprehendmedicalFilter) > 0 {
		if err := assignInputField(input, "Filter", _comprehendmedicalFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_comprehendmedicalMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _comprehendmedicalMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_comprehendmedicalNextToken) > 0 {
		input.NextToken = aws.String(_comprehendmedicalNextToken)
	}

	if resp, err := client.ListPHIDetectionJobs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of InferRxNorm jobs that you have submitted.
func comprehendmedical_ListRxNormInferenceJobs(cfg aws.Config, client *comprehendmedical.Client) {
	input := &comprehendmedical.ListRxNormInferenceJobsInput{}

	if len(_comprehendmedicalFilter) > 0 {
		if err := assignInputField(input, "Filter", _comprehendmedicalFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_comprehendmedicalMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _comprehendmedicalMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_comprehendmedicalNextToken) > 0 {
		input.NextToken = aws.String(_comprehendmedicalNextToken)
	}

	if resp, err := client.ListRxNormInferenceJobs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of InferSNOMEDCT jobs a user has submitted.
func comprehendmedical_ListSNOMEDCTInferenceJobs(cfg aws.Config, client *comprehendmedical.Client) {
	input := &comprehendmedical.ListSNOMEDCTInferenceJobsInput{}

	if len(_comprehendmedicalFilter) > 0 {
		if err := assignInputField(input, "Filter", _comprehendmedicalFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_comprehendmedicalMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _comprehendmedicalMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_comprehendmedicalNextToken) > 0 {
		input.NextToken = aws.String(_comprehendmedicalNextToken)
	}

	if resp, err := client.ListSNOMEDCTInferenceJobs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an asynchronous medical entity detection job for a collection of
// documents. Use the DescribeEntitiesDetectionV2Job operation to track the status
// of a job.
func comprehendmedical_StartEntitiesDetectionV2Job(cfg aws.Config, client *comprehendmedical.Client) {
	input := &comprehendmedical.StartEntitiesDetectionV2JobInput{
		// DataAccessRoleArn: *string, // Required
		// InputDataConfig: *types.InputDataConfig, // Required
		// LanguageCode: types.LanguageCode, // Required
		// OutputDataConfig: *types.OutputDataConfig, // Required
	}

	if len(_comprehendmedicalDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_comprehendmedicalDataAccessRoleArn)
	}
	if len(_comprehendmedicalInputDataConfig) > 0 {
		if err := assignInputField(input, "InputDataConfig", _comprehendmedicalInputDataConfig); err != nil {
			log.Errorf("invalid --input-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendmedicalLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _comprehendmedicalLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_comprehendmedicalOutputDataConfig) > 0 {
		if err := assignInputField(input, "OutputDataConfig", _comprehendmedicalOutputDataConfig); err != nil {
			log.Errorf("invalid --output-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendmedicalClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_comprehendmedicalClientRequestToken)
	}
	if len(_comprehendmedicalJobName) > 0 {
		input.JobName = aws.String(_comprehendmedicalJobName)
	}
	if len(_comprehendmedicalKMSKey) > 0 {
		input.KMSKey = aws.String(_comprehendmedicalKMSKey)
	}

	if resp, err := client.StartEntitiesDetectionV2Job(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an asynchronous job to detect medical conditions and link them to the
// ICD-10-CM ontology. Use the DescribeICD10CMInferenceJob operation to track the
// status of a job.
func comprehendmedical_StartICD10CMInferenceJob(cfg aws.Config, client *comprehendmedical.Client) {
	input := &comprehendmedical.StartICD10CMInferenceJobInput{
		// DataAccessRoleArn: *string, // Required
		// InputDataConfig: *types.InputDataConfig, // Required
		// LanguageCode: types.LanguageCode, // Required
		// OutputDataConfig: *types.OutputDataConfig, // Required
	}

	if len(_comprehendmedicalDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_comprehendmedicalDataAccessRoleArn)
	}
	if len(_comprehendmedicalInputDataConfig) > 0 {
		if err := assignInputField(input, "InputDataConfig", _comprehendmedicalInputDataConfig); err != nil {
			log.Errorf("invalid --input-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendmedicalLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _comprehendmedicalLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_comprehendmedicalOutputDataConfig) > 0 {
		if err := assignInputField(input, "OutputDataConfig", _comprehendmedicalOutputDataConfig); err != nil {
			log.Errorf("invalid --output-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendmedicalClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_comprehendmedicalClientRequestToken)
	}
	if len(_comprehendmedicalJobName) > 0 {
		input.JobName = aws.String(_comprehendmedicalJobName)
	}
	if len(_comprehendmedicalKMSKey) > 0 {
		input.KMSKey = aws.String(_comprehendmedicalKMSKey)
	}

	if resp, err := client.StartICD10CMInferenceJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an asynchronous job to detect protected health information (PHI). Use
// the DescribePHIDetectionJob operation to track the status of a job.
func comprehendmedical_StartPHIDetectionJob(cfg aws.Config, client *comprehendmedical.Client) {
	input := &comprehendmedical.StartPHIDetectionJobInput{
		// DataAccessRoleArn: *string, // Required
		// InputDataConfig: *types.InputDataConfig, // Required
		// LanguageCode: types.LanguageCode, // Required
		// OutputDataConfig: *types.OutputDataConfig, // Required
	}

	if len(_comprehendmedicalDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_comprehendmedicalDataAccessRoleArn)
	}
	if len(_comprehendmedicalInputDataConfig) > 0 {
		if err := assignInputField(input, "InputDataConfig", _comprehendmedicalInputDataConfig); err != nil {
			log.Errorf("invalid --input-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendmedicalLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _comprehendmedicalLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_comprehendmedicalOutputDataConfig) > 0 {
		if err := assignInputField(input, "OutputDataConfig", _comprehendmedicalOutputDataConfig); err != nil {
			log.Errorf("invalid --output-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendmedicalClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_comprehendmedicalClientRequestToken)
	}
	if len(_comprehendmedicalJobName) > 0 {
		input.JobName = aws.String(_comprehendmedicalJobName)
	}
	if len(_comprehendmedicalKMSKey) > 0 {
		input.KMSKey = aws.String(_comprehendmedicalKMSKey)
	}

	if resp, err := client.StartPHIDetectionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an asynchronous job to detect medication entities and link them to the
// RxNorm ontology. Use the DescribeRxNormInferenceJob operation to track the
// status of a job.
func comprehendmedical_StartRxNormInferenceJob(cfg aws.Config, client *comprehendmedical.Client) {
	input := &comprehendmedical.StartRxNormInferenceJobInput{
		// DataAccessRoleArn: *string, // Required
		// InputDataConfig: *types.InputDataConfig, // Required
		// LanguageCode: types.LanguageCode, // Required
		// OutputDataConfig: *types.OutputDataConfig, // Required
	}

	if len(_comprehendmedicalDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_comprehendmedicalDataAccessRoleArn)
	}
	if len(_comprehendmedicalInputDataConfig) > 0 {
		if err := assignInputField(input, "InputDataConfig", _comprehendmedicalInputDataConfig); err != nil {
			log.Errorf("invalid --input-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendmedicalLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _comprehendmedicalLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_comprehendmedicalOutputDataConfig) > 0 {
		if err := assignInputField(input, "OutputDataConfig", _comprehendmedicalOutputDataConfig); err != nil {
			log.Errorf("invalid --output-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendmedicalClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_comprehendmedicalClientRequestToken)
	}
	if len(_comprehendmedicalJobName) > 0 {
		input.JobName = aws.String(_comprehendmedicalJobName)
	}
	if len(_comprehendmedicalKMSKey) > 0 {
		input.KMSKey = aws.String(_comprehendmedicalKMSKey)
	}

	if resp, err := client.StartRxNormInferenceJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an asynchronous job to detect medical concepts and link them to the
// SNOMED-CT ontology. Use the DescribeSNOMEDCTInferenceJob operation to track the
// status of a job.
func comprehendmedical_StartSNOMEDCTInferenceJob(cfg aws.Config, client *comprehendmedical.Client) {
	input := &comprehendmedical.StartSNOMEDCTInferenceJobInput{
		// DataAccessRoleArn: *string, // Required
		// InputDataConfig: *types.InputDataConfig, // Required
		// LanguageCode: types.LanguageCode, // Required
		// OutputDataConfig: *types.OutputDataConfig, // Required
	}

	if len(_comprehendmedicalDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_comprehendmedicalDataAccessRoleArn)
	}
	if len(_comprehendmedicalInputDataConfig) > 0 {
		if err := assignInputField(input, "InputDataConfig", _comprehendmedicalInputDataConfig); err != nil {
			log.Errorf("invalid --input-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendmedicalLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _comprehendmedicalLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_comprehendmedicalOutputDataConfig) > 0 {
		if err := assignInputField(input, "OutputDataConfig", _comprehendmedicalOutputDataConfig); err != nil {
			log.Errorf("invalid --output-data-config: %s", err.Error())
			return
		}
	}
	if len(_comprehendmedicalClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_comprehendmedicalClientRequestToken)
	}
	if len(_comprehendmedicalJobName) > 0 {
		input.JobName = aws.String(_comprehendmedicalJobName)
	}
	if len(_comprehendmedicalKMSKey) > 0 {
		input.KMSKey = aws.String(_comprehendmedicalKMSKey)
	}

	if resp, err := client.StartSNOMEDCTInferenceJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a medical entities detection job in progress.
func comprehendmedical_StopEntitiesDetectionV2Job(cfg aws.Config, client *comprehendmedical.Client) {
	input := &comprehendmedical.StopEntitiesDetectionV2JobInput{
		// JobId: *string, // Required
	}

	if len(_comprehendmedicalJobId) > 0 {
		input.JobId = aws.String(_comprehendmedicalJobId)
	}

	if resp, err := client.StopEntitiesDetectionV2Job(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops an InferICD10CM inference job in progress.
func comprehendmedical_StopICD10CMInferenceJob(cfg aws.Config, client *comprehendmedical.Client) {
	input := &comprehendmedical.StopICD10CMInferenceJobInput{
		// JobId: *string, // Required
	}

	if len(_comprehendmedicalJobId) > 0 {
		input.JobId = aws.String(_comprehendmedicalJobId)
	}

	if resp, err := client.StopICD10CMInferenceJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a protected health information (PHI) detection job in progress.
func comprehendmedical_StopPHIDetectionJob(cfg aws.Config, client *comprehendmedical.Client) {
	input := &comprehendmedical.StopPHIDetectionJobInput{
		// JobId: *string, // Required
	}

	if len(_comprehendmedicalJobId) > 0 {
		input.JobId = aws.String(_comprehendmedicalJobId)
	}

	if resp, err := client.StopPHIDetectionJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops an InferRxNorm inference job in progress.
func comprehendmedical_StopRxNormInferenceJob(cfg aws.Config, client *comprehendmedical.Client) {
	input := &comprehendmedical.StopRxNormInferenceJobInput{
		// JobId: *string, // Required
	}

	if len(_comprehendmedicalJobId) > 0 {
		input.JobId = aws.String(_comprehendmedicalJobId)
	}

	if resp, err := client.StopRxNormInferenceJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops an InferSNOMEDCT inference job in progress.
func comprehendmedical_StopSNOMEDCTInferenceJob(cfg aws.Config, client *comprehendmedical.Client) {
	input := &comprehendmedical.StopSNOMEDCTInferenceJobInput{
		// JobId: *string, // Required
	}

	if len(_comprehendmedicalJobId) > 0 {
		input.JobId = aws.String(_comprehendmedicalJobId)
	}

	if resp, err := client.StopSNOMEDCTInferenceJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_comprehendmedicalCmd)
	_comprehendmedicalCmd.Flags().SortFlags = false

	_comprehendmedicalCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_comprehendmedicalCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_comprehendmedicalCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_comprehendmedicalCmd.Flags().StringVarP(&_comprehendmedicalClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_comprehendmedicalCmd.Flags().StringVarP(&_comprehendmedicalDataAccessRoleArn, "data-access-role-arn", "", "", "Data Access Role ARN")
	_comprehendmedicalCmd.Flags().StringVarP(&_comprehendmedicalFilter, "filter", "", "", "Filter")
	_comprehendmedicalCmd.Flags().StringVarP(&_comprehendmedicalInputDataConfig, "input-data-config", "", "", "Input Data Config")
	_comprehendmedicalCmd.Flags().StringVarP(&_comprehendmedicalJobId, "job-id", "", "", "Job ID")
	_comprehendmedicalCmd.Flags().StringVarP(&_comprehendmedicalJobName, "job-name", "", "", "Job Name")
	_comprehendmedicalCmd.Flags().StringVarP(&_comprehendmedicalKMSKey, "kms-key", "", "", "KMS Key")
	_comprehendmedicalCmd.Flags().StringVarP(&_comprehendmedicalLanguageCode, "language-code", "", "", "Language Code")
	_comprehendmedicalCmd.Flags().StringVarP(&_comprehendmedicalMaxResults, "max-results", "", "", "Max Results")
	_comprehendmedicalCmd.Flags().StringVarP(&_comprehendmedicalNextToken, "next-token", "", "", "Next Token")
	_comprehendmedicalCmd.Flags().StringVarP(&_comprehendmedicalOutputDataConfig, "output-data-config", "", "", "Output Data Config")
	_comprehendmedicalCmd.Flags().StringVarP(&_comprehendmedicalText, "text", "", "", "Text")

	_comprehendmedicalCmd.Flags().BoolVarP(&_comprehendmedicalDescribeEntitiesDetectionV2Job, "describe-entities-detection-v2-job", "", false, "Describe Entities Detection V2 Job")
	_comprehendmedicalCmd.Flags().BoolVarP(&_comprehendmedicalDescribeICD10CMInferenceJob, "describe-icd10cm-inference-job", "", false, "Describe Icd10cm Inference Job")
	_comprehendmedicalCmd.Flags().BoolVarP(&_comprehendmedicalDescribePHIDetectionJob, "describe-phi-detection-job", "", false, "Describe Phi Detection Job")
	_comprehendmedicalCmd.Flags().BoolVarP(&_comprehendmedicalDescribeRxNormInferenceJob, "describe-rx-norm-inference-job", "", false, "Describe Rx Norm Inference Job")
	_comprehendmedicalCmd.Flags().BoolVarP(&_comprehendmedicalDescribeSNOMEDCTInferenceJob, "describe-snomedct-inference-job", "", false, "Describe Snomedct Inference Job")
	_comprehendmedicalCmd.Flags().BoolVarP(&_comprehendmedicalDetectEntities, "detect-entities", "", false, "Detect Entities")
	_comprehendmedicalCmd.Flags().BoolVarP(&_comprehendmedicalDetectEntitiesV2, "detect-entities-v2", "", false, "Detect Entities V2")
	_comprehendmedicalCmd.Flags().BoolVarP(&_comprehendmedicalDetectPHI, "detect-phi", "", false, "Detect Phi")
	_comprehendmedicalCmd.Flags().BoolVarP(&_comprehendmedicalInferICD10CM, "infer-icd10cm", "", false, "Infer Icd10cm")
	_comprehendmedicalCmd.Flags().BoolVarP(&_comprehendmedicalInferRxNorm, "infer-rx-norm", "", false, "Infer Rx Norm")
	_comprehendmedicalCmd.Flags().BoolVarP(&_comprehendmedicalInferSNOMEDCT, "infer-snomedct", "", false, "Infer Snomedct")
	_comprehendmedicalCmd.Flags().BoolVarP(&_comprehendmedicalListEntitiesDetectionV2Jobs, "list-entities-detection-v2-jobs", "", false, "List Entities Detection V2 Jobs")
	_comprehendmedicalCmd.Flags().BoolVarP(&_comprehendmedicalListICD10CMInferenceJobs, "list-icd10cm-inference-jobs", "", false, "List Icd10cm Inference Jobs")
	_comprehendmedicalCmd.Flags().BoolVarP(&_comprehendmedicalListPHIDetectionJobs, "list-phi-detection-jobs", "", false, "List Phi Detection Jobs")
	_comprehendmedicalCmd.Flags().BoolVarP(&_comprehendmedicalListRxNormInferenceJobs, "list-rx-norm-inference-jobs", "", false, "List Rx Norm Inference Jobs")
	_comprehendmedicalCmd.Flags().BoolVarP(&_comprehendmedicalListSNOMEDCTInferenceJobs, "list-snomedct-inference-jobs", "", false, "List Snomedct Inference Jobs")
	_comprehendmedicalCmd.Flags().BoolVarP(&_comprehendmedicalStartEntitiesDetectionV2Job, "start-entities-detection-v2-job", "", false, "Start Entities Detection V2 Job")
	_comprehendmedicalCmd.Flags().BoolVarP(&_comprehendmedicalStartICD10CMInferenceJob, "start-icd10cm-inference-job", "", false, "Start Icd10cm Inference Job")
	_comprehendmedicalCmd.Flags().BoolVarP(&_comprehendmedicalStartPHIDetectionJob, "start-phi-detection-job", "", false, "Start Phi Detection Job")
	_comprehendmedicalCmd.Flags().BoolVarP(&_comprehendmedicalStartRxNormInferenceJob, "start-rx-norm-inference-job", "", false, "Start Rx Norm Inference Job")
	_comprehendmedicalCmd.Flags().BoolVarP(&_comprehendmedicalStartSNOMEDCTInferenceJob, "start-snomedct-inference-job", "", false, "Start Snomedct Inference Job")
	_comprehendmedicalCmd.Flags().BoolVarP(&_comprehendmedicalStopEntitiesDetectionV2Job, "stop-entities-detection-v2-job", "", false, "Stop Entities Detection V2 Job")
	_comprehendmedicalCmd.Flags().BoolVarP(&_comprehendmedicalStopICD10CMInferenceJob, "stop-icd10cm-inference-job", "", false, "Stop Icd10cm Inference Job")
	_comprehendmedicalCmd.Flags().BoolVarP(&_comprehendmedicalStopPHIDetectionJob, "stop-phi-detection-job", "", false, "Stop Phi Detection Job")
	_comprehendmedicalCmd.Flags().BoolVarP(&_comprehendmedicalStopRxNormInferenceJob, "stop-rx-norm-inference-job", "", false, "Stop Rx Norm Inference Job")
	_comprehendmedicalCmd.Flags().BoolVarP(&_comprehendmedicalStopSNOMEDCTInferenceJob, "stop-snomedct-inference-job", "", false, "Stop Snomedct Inference Job")

}
