package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/polly"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// pollyCmd represents the polly command
var _pollyCmd = &cobra.Command{
	Use:   "polly",
	Short: "AWS polly CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := polly.NewFromConfig(cfg)
		if _pollyDeleteLexicon {
			polly_DeleteLexicon(cfg, client)
			return
		}
		if _pollyDescribeVoices {
			polly_DescribeVoices(cfg, client)
			return
		}
		if _pollyGetLexicon {
			polly_GetLexicon(cfg, client)
			return
		}
		if _pollyGetSpeechSynthesisTask {
			polly_GetSpeechSynthesisTask(cfg, client)
			return
		}
		if _pollyListLexicons {
			polly_ListLexicons(cfg, client)
			return
		}
		if _pollyListSpeechSynthesisTasks {
			polly_ListSpeechSynthesisTasks(cfg, client)
			return
		}
		if _pollyPutLexicon {
			polly_PutLexicon(cfg, client)
			return
		}
		if _pollyStartSpeechSynthesisTask {
			polly_StartSpeechSynthesisTask(cfg, client)
			return
		}
		if _pollySynthesizeSpeech {
			polly_SynthesizeSpeech(cfg, client)
			return
		}

	},
}

var (
	_pollyDeleteLexicon            bool
	_pollyDescribeVoices           bool
	_pollyGetLexicon               bool
	_pollyGetSpeechSynthesisTask   bool
	_pollyListLexicons             bool
	_pollyListSpeechSynthesisTasks bool
	_pollyPutLexicon               bool
	_pollyStartSpeechSynthesisTask bool
	_pollySynthesizeSpeech         bool

	_pollyContent                        string
	_pollyEngine                         string
	_pollyIncludeAdditionalLanguageCodes string
	_pollyLanguageCode                   string
	_pollyLexiconNames                   []string
	_pollyMaxResults                     string
	_pollyName                           string
	_pollyNextToken                      string
	_pollyOutputFormat                   string
	_pollyOutputS3BucketName             string
	_pollyOutputS3KeyPrefix              string
	_pollySampleRate                     string
	_pollySnsTopicArn                    string
	_pollySpeechMarkTypes                string
	_pollyStatus                         string
	_pollyTaskId                         string
	_pollyText                           string
	_pollyTextType                       string
	_pollyVoiceId                        string
)

// Deletes the specified pronunciation lexicon stored in an Amazon Web Services
// Region. A lexicon which has been deleted is not available for speech synthesis,
// nor is it possible to retrieve it using either the GetLexicon or ListLexicon
// APIs.
//
// For more information, see [Managing Lexicons].
//
// [Managing Lexicons]: https://docs.aws.amazon.com/polly/latest/dg/managing-lexicons.html
func polly_DeleteLexicon(cfg aws.Config, client *polly.Client) {
	input := &polly.DeleteLexiconInput{
		// Name: *string, // Required
	}

	if len(_pollyName) > 0 {
		input.Name = aws.String(_pollyName)
	}

	if resp, err := client.DeleteLexicon(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the list of voices that are available for use when requesting speech
// synthesis. Each voice speaks a specified language, is either male or female, and
// is identified by an ID, which is the ASCII version of the voice name.
//
// When synthesizing speech ( SynthesizeSpeech ), you provide the voice ID for the
// voice you want from the list of voices returned by DescribeVoices .
//
// For example, you want your news reader application to read news in a specific
// language, but giving a user the option to choose the voice. Using the
// DescribeVoices operation you can provide the user with a list of available
// voices to select from.
//
// You can optionally specify a language code to filter the available voices. For
// example, if you specify en-US , the operation returns a list of all available US
// English voices.
//
// This operation requires permissions to perform the polly:DescribeVoices action.
func polly_DescribeVoices(cfg aws.Config, client *polly.Client) {
	input := &polly.DescribeVoicesInput{}

	if len(_pollyEngine) > 0 {
		if err := assignInputField(input, "Engine", _pollyEngine); err != nil {
			log.Errorf("invalid --engine: %s", err.Error())
			return
		}
	}
	if len(_pollyIncludeAdditionalLanguageCodes) > 0 {
		if err := assignInputField(input, "IncludeAdditionalLanguageCodes", _pollyIncludeAdditionalLanguageCodes); err != nil {
			log.Errorf("invalid --include-additional-language-codes: %s", err.Error())
			return
		}
	}
	if len(_pollyLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _pollyLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_pollyNextToken) > 0 {
		input.NextToken = aws.String(_pollyNextToken)
	}

	if resp, err := client.DescribeVoices(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the content of the specified pronunciation lexicon stored in an Amazon
// Web Services Region. For more information, see [Managing Lexicons].
//
// [Managing Lexicons]: https://docs.aws.amazon.com/polly/latest/dg/managing-lexicons.html
func polly_GetLexicon(cfg aws.Config, client *polly.Client) {
	input := &polly.GetLexiconInput{
		// Name: *string, // Required
	}

	if len(_pollyName) > 0 {
		input.Name = aws.String(_pollyName)
	}

	if resp, err := client.GetLexicon(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a specific SpeechSynthesisTask object based on its TaskID. This
// object contains information about the given speech synthesis task, including the
// status of the task, and a link to the S3 bucket containing the output of the
// task.
func polly_GetSpeechSynthesisTask(cfg aws.Config, client *polly.Client) {
	input := &polly.GetSpeechSynthesisTaskInput{
		// TaskId: *string, // Required
	}

	if len(_pollyTaskId) > 0 {
		input.TaskId = aws.String(_pollyTaskId)
	}

	if resp, err := client.GetSpeechSynthesisTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of pronunciation lexicons stored in an Amazon Web Services
// Region. For more information, see [Managing Lexicons].
//
// [Managing Lexicons]: https://docs.aws.amazon.com/polly/latest/dg/managing-lexicons.html
func polly_ListLexicons(cfg aws.Config, client *polly.Client) {
	input := &polly.ListLexiconsInput{}

	if len(_pollyNextToken) > 0 {
		input.NextToken = aws.String(_pollyNextToken)
	}

	if resp, err := client.ListLexicons(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of SpeechSynthesisTask objects ordered by their creation date.
// This operation can filter the tasks by their status, for example, allowing users
// to list only tasks that are completed.
func polly_ListSpeechSynthesisTasks(cfg aws.Config, client *polly.Client) {
	input := &polly.ListSpeechSynthesisTasksInput{}

	if len(_pollyMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pollyMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pollyNextToken) > 0 {
		input.NextToken = aws.String(_pollyNextToken)
	}
	if len(_pollyStatus) > 0 {
		if err := assignInputField(input, "Status", _pollyStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSpeechSynthesisTasks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*polly.ListSpeechSynthesisTasksOutput
	p := polly.NewListSpeechSynthesisTasksPaginator(client, input)
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

// Stores a pronunciation lexicon in an Amazon Web Services Region. If a lexicon
// with the same name already exists in the region, it is overwritten by the new
// lexicon. Lexicon operations have eventual consistency, therefore, it might take
// some time before the lexicon is available to the SynthesizeSpeech operation.
//
// For more information, see [Managing Lexicons].
//
// [Managing Lexicons]: https://docs.aws.amazon.com/polly/latest/dg/managing-lexicons.html
func polly_PutLexicon(cfg aws.Config, client *polly.Client) {
	input := &polly.PutLexiconInput{
		// Content: *string, // Required
		// Name: *string, // Required
	}

	if len(_pollyContent) > 0 {
		input.Content = aws.String(_pollyContent)
	}
	if len(_pollyName) > 0 {
		input.Name = aws.String(_pollyName)
	}

	if resp, err := client.PutLexicon(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows the creation of an asynchronous synthesis task, by starting a new
// SpeechSynthesisTask . This operation requires all the standard information
// needed for speech synthesis, plus the name of an Amazon S3 bucket for the
// service to store the output of the synthesis task and two optional parameters (
// OutputS3KeyPrefix and SnsTopicArn ). Once the synthesis task is created, this
// operation will return a SpeechSynthesisTask object, which will include an
// identifier of this task as well as the current status. The SpeechSynthesisTask
// object is available for 72 hours after starting the asynchronous synthesis task.
func polly_StartSpeechSynthesisTask(cfg aws.Config, client *polly.Client) {
	input := &polly.StartSpeechSynthesisTaskInput{
		// OutputFormat: types.OutputFormat, // Required
		// OutputS3BucketName: *string, // Required
		// Text: *string, // Required
		// VoiceId: types.VoiceId, // Required
	}

	if len(_pollyOutputFormat) > 0 {
		if err := assignInputField(input, "OutputFormat", _pollyOutputFormat); err != nil {
			log.Errorf("invalid --output-format: %s", err.Error())
			return
		}
	}
	if len(_pollyOutputS3BucketName) > 0 {
		input.OutputS3BucketName = aws.String(_pollyOutputS3BucketName)
	}
	if len(_pollyText) > 0 {
		input.Text = aws.String(_pollyText)
	}
	if len(_pollyVoiceId) > 0 {
		if err := assignInputField(input, "VoiceId", _pollyVoiceId); err != nil {
			log.Errorf("invalid --voice-id: %s", err.Error())
			return
		}
	}
	if len(_pollyEngine) > 0 {
		if err := assignInputField(input, "Engine", _pollyEngine); err != nil {
			log.Errorf("invalid --engine: %s", err.Error())
			return
		}
	}
	if len(_pollyLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _pollyLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_pollyLexiconNames) > 0 {
		input.LexiconNames = append([]string(nil), _pollyLexiconNames...)
	}
	if len(_pollyOutputS3KeyPrefix) > 0 {
		input.OutputS3KeyPrefix = aws.String(_pollyOutputS3KeyPrefix)
	}
	if len(_pollySampleRate) > 0 {
		input.SampleRate = aws.String(_pollySampleRate)
	}
	if len(_pollySnsTopicArn) > 0 {
		input.SnsTopicArn = aws.String(_pollySnsTopicArn)
	}
	if len(_pollySpeechMarkTypes) > 0 {
		if err := assignInputField(input, "SpeechMarkTypes", _pollySpeechMarkTypes); err != nil {
			log.Errorf("invalid --speech-mark-types: %s", err.Error())
			return
		}
	}
	if len(_pollyTextType) > 0 {
		if err := assignInputField(input, "TextType", _pollyTextType); err != nil {
			log.Errorf("invalid --text-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartSpeechSynthesisTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Synthesizes UTF-8 input, plain text or SSML, to a stream of bytes. SSML input
// must be valid, well-formed SSML. Some alphabets might not be available with all
// the voices (for example, Cyrillic might not be read at all by English voices)
// unless phoneme mapping is used. For more information, see [How it Works].
//
// [How it Works]: https://docs.aws.amazon.com/polly/latest/dg/how-text-to-speech-works.html
func polly_SynthesizeSpeech(cfg aws.Config, client *polly.Client) {
	input := &polly.SynthesizeSpeechInput{
		// OutputFormat: types.OutputFormat, // Required
		// Text: *string, // Required
		// VoiceId: types.VoiceId, // Required
	}

	if len(_pollyOutputFormat) > 0 {
		if err := assignInputField(input, "OutputFormat", _pollyOutputFormat); err != nil {
			log.Errorf("invalid --output-format: %s", err.Error())
			return
		}
	}
	if len(_pollyText) > 0 {
		input.Text = aws.String(_pollyText)
	}
	if len(_pollyVoiceId) > 0 {
		if err := assignInputField(input, "VoiceId", _pollyVoiceId); err != nil {
			log.Errorf("invalid --voice-id: %s", err.Error())
			return
		}
	}
	if len(_pollyEngine) > 0 {
		if err := assignInputField(input, "Engine", _pollyEngine); err != nil {
			log.Errorf("invalid --engine: %s", err.Error())
			return
		}
	}
	if len(_pollyLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _pollyLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_pollyLexiconNames) > 0 {
		input.LexiconNames = append([]string(nil), _pollyLexiconNames...)
	}
	if len(_pollySampleRate) > 0 {
		input.SampleRate = aws.String(_pollySampleRate)
	}
	if len(_pollySpeechMarkTypes) > 0 {
		if err := assignInputField(input, "SpeechMarkTypes", _pollySpeechMarkTypes); err != nil {
			log.Errorf("invalid --speech-mark-types: %s", err.Error())
			return
		}
	}
	if len(_pollyTextType) > 0 {
		if err := assignInputField(input, "TextType", _pollyTextType); err != nil {
			log.Errorf("invalid --text-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.SynthesizeSpeech(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_pollyCmd)
	_pollyCmd.Flags().SortFlags = false

	_pollyCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_pollyCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_pollyCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_pollyCmd.Flags().StringVarP(&_pollyContent, "content", "", "", "Content")
	_pollyCmd.Flags().StringVarP(&_pollyEngine, "engine", "", "", "Engine")
	_pollyCmd.Flags().StringVarP(&_pollyIncludeAdditionalLanguageCodes, "include-additional-language-codes", "", "", "Include Additional Language Codes")
	_pollyCmd.Flags().StringVarP(&_pollyLanguageCode, "language-code", "", "", "Language Code")
	_pollyCmd.Flags().StringSliceVarP(&_pollyLexiconNames, "lexicon-names", "", nil, "Lexicon Names")
	_pollyCmd.Flags().StringVarP(&_pollyMaxResults, "max-results", "", "", "Max Results")
	_pollyCmd.Flags().StringVarP(&_pollyName, "name", "", "", "Name")
	_pollyCmd.Flags().StringVarP(&_pollyNextToken, "next-token", "", "", "Next Token")
	_pollyCmd.Flags().StringVarP(&_pollyOutputFormat, "output-format", "", "", "Output Format")
	_pollyCmd.Flags().StringVarP(&_pollyOutputS3BucketName, "output-s3-bucket-name", "", "", "Output S3 Bucket Name")
	_pollyCmd.Flags().StringVarP(&_pollyOutputS3KeyPrefix, "output-s3-key-prefix", "", "", "Output S3 Key Prefix")
	_pollyCmd.Flags().StringVarP(&_pollySampleRate, "sample-rate", "", "", "Sample Rate")
	_pollyCmd.Flags().StringVarP(&_pollySnsTopicArn, "sns-topic-arn", "", "", "SNS Topic ARN")
	_pollyCmd.Flags().StringVarP(&_pollySpeechMarkTypes, "speech-mark-types", "", "", "Speech Mark Types")
	_pollyCmd.Flags().StringVarP(&_pollyStatus, "status", "", "", "Status")
	_pollyCmd.Flags().StringVarP(&_pollyTaskId, "task-id", "", "", "Task ID")
	_pollyCmd.Flags().StringVarP(&_pollyText, "text", "", "", "Text")
	_pollyCmd.Flags().StringVarP(&_pollyTextType, "text-type", "", "", "Text Type")
	_pollyCmd.Flags().StringVarP(&_pollyVoiceId, "voice-id", "", "", "Voice ID")

	_pollyCmd.Flags().BoolVarP(&_pollyDeleteLexicon, "delete-lexicon", "", false, "Delete Lexicon")
	_pollyCmd.Flags().BoolVarP(&_pollyDescribeVoices, "describe-voices", "", false, "Describe Voices")
	_pollyCmd.Flags().BoolVarP(&_pollyGetLexicon, "get-lexicon", "", false, "Get Lexicon")
	_pollyCmd.Flags().BoolVarP(&_pollyGetSpeechSynthesisTask, "get-speech-synthesis-task", "", false, "Get Speech Synthesis Task")
	_pollyCmd.Flags().BoolVarP(&_pollyListLexicons, "list-lexicons", "", false, "List Lexicons")
	_pollyCmd.Flags().BoolVarP(&_pollyListSpeechSynthesisTasks, "list-speech-synthesis-tasks", "", false, "List Speech Synthesis Tasks")
	_pollyCmd.Flags().BoolVarP(&_pollyPutLexicon, "put-lexicon", "", false, "Put Lexicon")
	_pollyCmd.Flags().BoolVarP(&_pollyStartSpeechSynthesisTask, "start-speech-synthesis-task", "", false, "Start Speech Synthesis Task")
	_pollyCmd.Flags().BoolVarP(&_pollySynthesizeSpeech, "synthesize-speech", "", false, "Synthesize Speech")

}
