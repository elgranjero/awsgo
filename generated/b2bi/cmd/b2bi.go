package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/b2bi"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// b2biCmd represents the b2bi command
var _b2biCmd = &cobra.Command{
	Use:   "b2bi",
	Short: "AWS b2bi CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := b2bi.NewFromConfig(cfg)
		if _b2biCreateCapability {
			b2bi_CreateCapability(cfg, client)
			return
		}
		if _b2biCreatePartnership {
			b2bi_CreatePartnership(cfg, client)
			return
		}
		if _b2biCreateProfile {
			b2bi_CreateProfile(cfg, client)
			return
		}
		if _b2biCreateStarterMappingTemplate {
			b2bi_CreateStarterMappingTemplate(cfg, client)
			return
		}
		if _b2biCreateTransformer {
			b2bi_CreateTransformer(cfg, client)
			return
		}
		if _b2biDeleteCapability {
			b2bi_DeleteCapability(cfg, client)
			return
		}
		if _b2biDeletePartnership {
			b2bi_DeletePartnership(cfg, client)
			return
		}
		if _b2biDeleteProfile {
			b2bi_DeleteProfile(cfg, client)
			return
		}
		if _b2biDeleteTransformer {
			b2bi_DeleteTransformer(cfg, client)
			return
		}
		if _b2biGenerateMapping {
			b2bi_GenerateMapping(cfg, client)
			return
		}
		if _b2biGetCapability {
			b2bi_GetCapability(cfg, client)
			return
		}
		if _b2biGetPartnership {
			b2bi_GetPartnership(cfg, client)
			return
		}
		if _b2biGetProfile {
			b2bi_GetProfile(cfg, client)
			return
		}
		if _b2biGetTransformer {
			b2bi_GetTransformer(cfg, client)
			return
		}
		if _b2biGetTransformerJob {
			b2bi_GetTransformerJob(cfg, client)
			return
		}
		if _b2biListCapabilities {
			b2bi_ListCapabilities(cfg, client)
			return
		}
		if _b2biListPartnerships {
			b2bi_ListPartnerships(cfg, client)
			return
		}
		if _b2biListProfiles {
			b2bi_ListProfiles(cfg, client)
			return
		}
		if _b2biListTagsForResource {
			b2bi_ListTagsForResource(cfg, client)
			return
		}
		if _b2biListTransformers {
			b2bi_ListTransformers(cfg, client)
			return
		}
		if _b2biStartTransformerJob {
			b2bi_StartTransformerJob(cfg, client)
			return
		}
		if _b2biTagResource {
			b2bi_TagResource(cfg, client)
			return
		}
		if _b2biTestConversion {
			b2bi_TestConversion(cfg, client)
			return
		}
		if _b2biTestMapping {
			b2bi_TestMapping(cfg, client)
			return
		}
		if _b2biTestParsing {
			b2bi_TestParsing(cfg, client)
			return
		}
		if _b2biUntagResource {
			b2bi_UntagResource(cfg, client)
			return
		}
		if _b2biUpdateCapability {
			b2bi_UpdateCapability(cfg, client)
			return
		}
		if _b2biUpdatePartnership {
			b2bi_UpdatePartnership(cfg, client)
			return
		}
		if _b2biUpdateProfile {
			b2bi_UpdateProfile(cfg, client)
			return
		}
		if _b2biUpdateTransformer {
			b2bi_UpdateTransformer(cfg, client)
			return
		}

	},
}

var (
	_b2biCreateCapability             bool
	_b2biCreatePartnership            bool
	_b2biCreateProfile                bool
	_b2biCreateStarterMappingTemplate bool
	_b2biCreateTransformer            bool
	_b2biDeleteCapability             bool
	_b2biDeletePartnership            bool
	_b2biDeleteProfile                bool
	_b2biDeleteTransformer            bool
	_b2biGenerateMapping              bool
	_b2biGetCapability                bool
	_b2biGetPartnership               bool
	_b2biGetProfile                   bool
	_b2biGetTransformer               bool
	_b2biGetTransformerJob            bool
	_b2biListCapabilities             bool
	_b2biListPartnerships             bool
	_b2biListProfiles                 bool
	_b2biListTagsForResource          bool
	_b2biListTransformers             bool
	_b2biStartTransformerJob          bool
	_b2biTagResource                  bool
	_b2biTestConversion               bool
	_b2biTestMapping                  bool
	_b2biTestParsing                  bool
	_b2biUntagResource                bool
	_b2biUpdateCapability             bool
	_b2biUpdatePartnership            bool
	_b2biUpdateProfile                bool
	_b2biUpdateTransformer            bool

	_b2biAdvancedOptions       string
	_b2biBusinessName          string
	_b2biCapabilities          []string
	_b2biCapabilityId          string
	_b2biCapabilityOptions     string
	_b2biClientToken           string
	_b2biConfiguration         string
	_b2biEdiType               string
	_b2biEmail                 string
	_b2biFileFormat            string
	_b2biInputConversion       string
	_b2biInputFile             string
	_b2biInputFileContent      string
	_b2biInstructionsDocuments string
	_b2biLogging               string
	_b2biMapping               string
	_b2biMappingTemplate       string
	_b2biMappingType           string
	_b2biMaxResults            string
	_b2biName                  string
	_b2biNextToken             string
	_b2biOutputConversion      string
	_b2biOutputFileContent     string
	_b2biOutputLocation        string
	_b2biOutputSampleLocation  string
	_b2biPartnershipId         string
	_b2biPhone                 string
	_b2biProfileId             string
	_b2biResourceARN           string
	_b2biSampleDocument        string
	_b2biSampleDocuments       string
	_b2biSource                string
	_b2biStatus                string
	_b2biTagKeys               []string
	_b2biTags                  string
	_b2biTarget                string
	_b2biTemplateDetails       string
	_b2biTransformerId         string
	_b2biTransformerJobId      string
	_b2biType                  string
)

// Instantiates a capability based on the specified parameters. A trading
// capability contains the information required to transform incoming EDI documents
// into JSON or XML outputs.
func b2bi_CreateCapability(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.CreateCapabilityInput{
		// Configuration: types.CapabilityConfiguration, // Required
		// Name: *string, // Required
		// Type: types.CapabilityType, // Required
	}

	if len(_b2biConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _b2biConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_b2biName) > 0 {
		input.Name = aws.String(_b2biName)
	}
	if len(_b2biType) > 0 {
		if err := assignInputField(input, "Type", _b2biType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_b2biClientToken) > 0 {
		input.ClientToken = aws.String(_b2biClientToken)
	}
	if len(_b2biInstructionsDocuments) > 0 {
		if err := assignInputField(input, "InstructionsDocuments", _b2biInstructionsDocuments); err != nil {
			log.Errorf("invalid --instructions-documents: %s", err.Error())
			return
		}
	}
	if len(_b2biTags) > 0 {
		if err := assignInputField(input, "Tags", _b2biTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCapability(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a partnership between a customer and a trading partner, based on the
// supplied parameters. A partnership represents the connection between you and
// your trading partner. It ties together a profile and one or more trading
// capabilities.
func b2bi_CreatePartnership(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.CreatePartnershipInput{
		// Capabilities: []string, // Required
		// Email: *string, // Required
		// Name: *string, // Required
		// ProfileId: *string, // Required
	}

	if len(_b2biCapabilities) > 0 {
		input.Capabilities = append([]string(nil), _b2biCapabilities...)
	}
	if len(_b2biEmail) > 0 {
		input.Email = aws.String(_b2biEmail)
	}
	if len(_b2biName) > 0 {
		input.Name = aws.String(_b2biName)
	}
	if len(_b2biProfileId) > 0 {
		input.ProfileId = aws.String(_b2biProfileId)
	}
	if len(_b2biCapabilityOptions) > 0 {
		if err := assignInputField(input, "CapabilityOptions", _b2biCapabilityOptions); err != nil {
			log.Errorf("invalid --capability-options: %s", err.Error())
			return
		}
	}
	if len(_b2biClientToken) > 0 {
		input.ClientToken = aws.String(_b2biClientToken)
	}
	if len(_b2biPhone) > 0 {
		input.Phone = aws.String(_b2biPhone)
	}
	if len(_b2biTags) > 0 {
		if err := assignInputField(input, "Tags", _b2biTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePartnership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a customer profile. You can have up to five customer profiles, each
// representing a distinct private network. A profile is the mechanism used to
// create the concept of a private network.
func b2bi_CreateProfile(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.CreateProfileInput{
		// BusinessName: *string, // Required
		// Logging: types.Logging, // Required
		// Name: *string, // Required
		// Phone: *string, // Required
	}

	if len(_b2biBusinessName) > 0 {
		input.BusinessName = aws.String(_b2biBusinessName)
	}
	if len(_b2biLogging) > 0 {
		if err := assignInputField(input, "Logging", _b2biLogging); err != nil {
			log.Errorf("invalid --logging: %s", err.Error())
			return
		}
	}
	if len(_b2biName) > 0 {
		input.Name = aws.String(_b2biName)
	}
	if len(_b2biPhone) > 0 {
		input.Phone = aws.String(_b2biPhone)
	}
	if len(_b2biClientToken) > 0 {
		input.ClientToken = aws.String(_b2biClientToken)
	}
	if len(_b2biEmail) > 0 {
		input.Email = aws.String(_b2biEmail)
	}
	if len(_b2biTags) > 0 {
		if err := assignInputField(input, "Tags", _b2biTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Amazon Web Services B2B Data Interchange uses a mapping template in JSONata or
// XSLT format to transform a customer input file into a JSON or XML file that can
// be converted to EDI.
//
// If you provide a sample EDI file with the same structure as the EDI files that
// you wish to generate, then the service can generate a mapping template. The
// starter template contains placeholder values which you can replace with JSONata
// or XSLT expressions to take data from your input file and insert it into the
// JSON or XML file that is used to generate the EDI.
//
// If you do not provide a sample EDI file, then the service can generate a
// mapping template based on the EDI settings in the templateDetails parameter.
//
// Currently, we only support generating a template that can generate the input to
// produce an Outbound X12 EDI file.
func b2bi_CreateStarterMappingTemplate(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.CreateStarterMappingTemplateInput{
		// MappingType: types.MappingType, // Required
		// TemplateDetails: types.TemplateDetails, // Required
	}

	if len(_b2biMappingType) > 0 {
		if err := assignInputField(input, "MappingType", _b2biMappingType); err != nil {
			log.Errorf("invalid --mapping-type: %s", err.Error())
			return
		}
	}
	if len(_b2biTemplateDetails) > 0 {
		if err := assignInputField(input, "TemplateDetails", _b2biTemplateDetails); err != nil {
			log.Errorf("invalid --template-details: %s", err.Error())
			return
		}
	}
	if len(_b2biOutputSampleLocation) > 0 {
		if err := assignInputField(input, "OutputSampleLocation", _b2biOutputSampleLocation); err != nil {
			log.Errorf("invalid --output-sample-location: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateStarterMappingTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a transformer. Amazon Web Services B2B Data Interchange currently
// supports two scenarios:
//
// - Inbound EDI: the Amazon Web Services customer receives an EDI file from
// their trading partner. Amazon Web Services B2B Data Interchange converts this
// EDI file into a JSON or XML file with a service-defined structure. A mapping
// template provided by the customer, in JSONata or XSLT format, is optionally
// applied to this file to produce a JSON or XML file with the structure the
// customer requires.
//
// - Outbound EDI: the Amazon Web Services customer has a JSON or XML file
// containing data that they wish to use in an EDI file. A mapping template,
// provided by the customer (in either JSONata or XSLT format) is applied to this
// file to generate a JSON or XML file in the service-defined structure. This file
// is then converted to an EDI file.
//
// The following fields are provided for backwards compatibility only: fileFormat ,
// mappingTemplate , ediType , and sampleDocument .
//
// - Use the mapping data type in place of mappingTemplate and fileFormat
//
// - Use the sampleDocuments data type in place of sampleDocument
//
// - Use either the inputConversion or outputConversion in place of ediType
func b2bi_CreateTransformer(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.CreateTransformerInput{
		// Name: *string, // Required
	}

	if len(_b2biName) > 0 {
		input.Name = aws.String(_b2biName)
	}
	if len(_b2biClientToken) > 0 {
		input.ClientToken = aws.String(_b2biClientToken)
	}
	if len(_b2biEdiType) > 0 {
		if err := assignInputField(input, "EdiType", _b2biEdiType); err != nil {
			log.Errorf("invalid --edi-type: %s", err.Error())
			return
		}
	}
	if len(_b2biFileFormat) > 0 {
		if err := assignInputField(input, "FileFormat", _b2biFileFormat); err != nil {
			log.Errorf("invalid --file-format: %s", err.Error())
			return
		}
	}
	if len(_b2biInputConversion) > 0 {
		if err := assignInputField(input, "InputConversion", _b2biInputConversion); err != nil {
			log.Errorf("invalid --input-conversion: %s", err.Error())
			return
		}
	}
	if len(_b2biMapping) > 0 {
		if err := assignInputField(input, "Mapping", _b2biMapping); err != nil {
			log.Errorf("invalid --mapping: %s", err.Error())
			return
		}
	}
	if len(_b2biMappingTemplate) > 0 {
		input.MappingTemplate = aws.String(_b2biMappingTemplate)
	}
	if len(_b2biOutputConversion) > 0 {
		if err := assignInputField(input, "OutputConversion", _b2biOutputConversion); err != nil {
			log.Errorf("invalid --output-conversion: %s", err.Error())
			return
		}
	}
	if len(_b2biSampleDocument) > 0 {
		input.SampleDocument = aws.String(_b2biSampleDocument)
	}
	if len(_b2biSampleDocuments) > 0 {
		if err := assignInputField(input, "SampleDocuments", _b2biSampleDocuments); err != nil {
			log.Errorf("invalid --sample-documents: %s", err.Error())
			return
		}
	}
	if len(_b2biTags) > 0 {
		if err := assignInputField(input, "Tags", _b2biTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTransformer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified capability. A trading capability contains the information
// required to transform incoming EDI documents into JSON or XML outputs.
func b2bi_DeleteCapability(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.DeleteCapabilityInput{
		// CapabilityId: *string, // Required
	}

	if len(_b2biCapabilityId) > 0 {
		input.CapabilityId = aws.String(_b2biCapabilityId)
	}

	if resp, err := client.DeleteCapability(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified partnership. A partnership represents the connection
// between you and your trading partner. It ties together a profile and one or more
// trading capabilities.
func b2bi_DeletePartnership(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.DeletePartnershipInput{
		// PartnershipId: *string, // Required
	}

	if len(_b2biPartnershipId) > 0 {
		input.PartnershipId = aws.String(_b2biPartnershipId)
	}

	if resp, err := client.DeletePartnership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified profile. A profile is the mechanism used to create the
// concept of a private network.
func b2bi_DeleteProfile(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.DeleteProfileInput{
		// ProfileId: *string, // Required
	}

	if len(_b2biProfileId) > 0 {
		input.ProfileId = aws.String(_b2biProfileId)
	}

	if resp, err := client.DeleteProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified transformer. A transformer can take an EDI file as input
// and transform it into a JSON-or XML-formatted document. Alternatively, a
// transformer can take a JSON-or XML-formatted document as input and transform it
// into an EDI file.
func b2bi_DeleteTransformer(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.DeleteTransformerInput{
		// TransformerId: *string, // Required
	}

	if len(_b2biTransformerId) > 0 {
		input.TransformerId = aws.String(_b2biTransformerId)
	}

	if resp, err := client.DeleteTransformer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Takes sample input and output documents and uses Amazon Bedrock to generate a
// mapping automatically. Depending on the accuracy and other factors, you can then
// edit the mapping for your needs.
//
// Before you can use the AI-assisted feature for Amazon Web Services B2B Data
// Interchange you must enable models in Amazon Bedrock. For details, see [AI-assisted template mapping prerequisites]in the
// Amazon Web Services B2B Data Interchange User guide.
//
// To generate a mapping, perform the following steps:
//
// - Start with an X12 EDI document to use as the input.
//
// - Call TestMapping using your EDI document.
//
// - Use the output from the TestMapping operation as either input or output for
// your GenerateMapping call, along with your sample file.
//
// [AI-assisted template mapping prerequisites]: https://docs.aws.amazon.com/b2bi/latest/userguide/ai-assisted-mapping.html#ai-assist-prereq
func b2bi_GenerateMapping(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.GenerateMappingInput{
		// InputFileContent: *string, // Required
		// MappingType: types.MappingType, // Required
		// OutputFileContent: *string, // Required
	}

	if len(_b2biInputFileContent) > 0 {
		input.InputFileContent = aws.String(_b2biInputFileContent)
	}
	if len(_b2biMappingType) > 0 {
		if err := assignInputField(input, "MappingType", _b2biMappingType); err != nil {
			log.Errorf("invalid --mapping-type: %s", err.Error())
			return
		}
	}
	if len(_b2biOutputFileContent) > 0 {
		input.OutputFileContent = aws.String(_b2biOutputFileContent)
	}

	if resp, err := client.GenerateMapping(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details for the specified capability. A trading capability
// contains the information required to transform incoming EDI documents into JSON
// or XML outputs.
func b2bi_GetCapability(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.GetCapabilityInput{
		// CapabilityId: *string, // Required
	}

	if len(_b2biCapabilityId) > 0 {
		input.CapabilityId = aws.String(_b2biCapabilityId)
	}

	if resp, err := client.GetCapability(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details for a partnership, based on the partner and profile IDs
// specified. A partnership represents the connection between you and your trading
// partner. It ties together a profile and one or more trading capabilities.
func b2bi_GetPartnership(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.GetPartnershipInput{
		// PartnershipId: *string, // Required
	}

	if len(_b2biPartnershipId) > 0 {
		input.PartnershipId = aws.String(_b2biPartnershipId)
	}

	if resp, err := client.GetPartnership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details for the profile specified by the profile ID. A profile is
// the mechanism used to create the concept of a private network.
func b2bi_GetProfile(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.GetProfileInput{
		// ProfileId: *string, // Required
	}

	if len(_b2biProfileId) > 0 {
		input.ProfileId = aws.String(_b2biProfileId)
	}

	if resp, err := client.GetProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details for the transformer specified by the transformer ID. A
// transformer can take an EDI file as input and transform it into a JSON-or
// XML-formatted document. Alternatively, a transformer can take a JSON-or
// XML-formatted document as input and transform it into an EDI file.
func b2bi_GetTransformer(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.GetTransformerInput{
		// TransformerId: *string, // Required
	}

	if len(_b2biTransformerId) > 0 {
		input.TransformerId = aws.String(_b2biTransformerId)
	}

	if resp, err := client.GetTransformer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the details of the transformer run, based on the Transformer job ID.
// If 30 days have elapsed since your transformer job was started, the system
// deletes it. So, if you run GetTransformerJob and supply a transformerId and
// transformerJobId for a job that was started more than 30 days previously, you
// receive a 404 response.
func b2bi_GetTransformerJob(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.GetTransformerJobInput{
		// TransformerId: *string, // Required
		// TransformerJobId: *string, // Required
	}

	if len(_b2biTransformerId) > 0 {
		input.TransformerId = aws.String(_b2biTransformerId)
	}
	if len(_b2biTransformerJobId) > 0 {
		input.TransformerJobId = aws.String(_b2biTransformerJobId)
	}

	if resp, err := client.GetTransformerJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the capabilities associated with your Amazon Web Services account for
// your current or specified region. A trading capability contains the information
// required to transform incoming EDI documents into JSON or XML outputs.
func b2bi_ListCapabilities(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.ListCapabilitiesInput{}

	if len(_b2biMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _b2biMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_b2biNextToken) > 0 {
		input.NextToken = aws.String(_b2biNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCapabilities(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*b2bi.ListCapabilitiesOutput
	p := b2bi.NewListCapabilitiesPaginator(client, input)
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

// Lists the partnerships associated with your Amazon Web Services account for
// your current or specified region. A partnership represents the connection
// between you and your trading partner. It ties together a profile and one or more
// trading capabilities.
func b2bi_ListPartnerships(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.ListPartnershipsInput{}

	if len(_b2biMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _b2biMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_b2biNextToken) > 0 {
		input.NextToken = aws.String(_b2biNextToken)
	}
	if len(_b2biProfileId) > 0 {
		input.ProfileId = aws.String(_b2biProfileId)
	}

	if disablePaginator() {
		if resp, err := client.ListPartnerships(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*b2bi.ListPartnershipsOutput
	p := b2bi.NewListPartnershipsPaginator(client, input)
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

// Lists the profiles associated with your Amazon Web Services account for your
// current or specified region. A profile is the mechanism used to create the
// concept of a private network.
func b2bi_ListProfiles(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.ListProfilesInput{}

	if len(_b2biMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _b2biMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_b2biNextToken) > 0 {
		input.NextToken = aws.String(_b2biNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*b2bi.ListProfilesOutput
	p := b2bi.NewListProfilesPaginator(client, input)
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

// Lists all of the tags associated with the Amazon Resource Name (ARN) that you
// specify. The resource can be a capability, partnership, profile, or transformer.
func b2bi_ListTagsForResource(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_b2biResourceARN) > 0 {
		input.ResourceARN = aws.String(_b2biResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the available transformers. A transformer can take an EDI file as input
// and transform it into a JSON-or XML-formatted document. Alternatively, a
// transformer can take a JSON-or XML-formatted document as input and transform it
// into an EDI file.
func b2bi_ListTransformers(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.ListTransformersInput{}

	if len(_b2biMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _b2biMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_b2biNextToken) > 0 {
		input.NextToken = aws.String(_b2biNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTransformers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*b2bi.ListTransformersOutput
	p := b2bi.NewListTransformersPaginator(client, input)
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

// Runs a job, using a transformer, to parse input EDI (electronic data
// interchange) file into the output structures used by Amazon Web Services B2B
// Data Interchange.
//
// If you only want to transform EDI (electronic data interchange) documents, you
// don't need to create profiles, partnerships or capabilities. Just create and
// configure a transformer, and then run the StartTransformerJob API to process
// your files.
//
// The system stores transformer jobs for 30 days. During that period, you can run [GetTransformerJob]
// and supply its transformerId and transformerJobId to return details of the job.
//
// [GetTransformerJob]: https://docs.aws.amazon.com/b2bi/latest/APIReference/API_GetTransformerJob.html
func b2bi_StartTransformerJob(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.StartTransformerJobInput{
		// InputFile: *types.S3Location, // Required
		// OutputLocation: *types.S3Location, // Required
		// TransformerId: *string, // Required
	}

	if len(_b2biInputFile) > 0 {
		if err := assignInputField(input, "InputFile", _b2biInputFile); err != nil {
			log.Errorf("invalid --input-file: %s", err.Error())
			return
		}
	}
	if len(_b2biOutputLocation) > 0 {
		if err := assignInputField(input, "OutputLocation", _b2biOutputLocation); err != nil {
			log.Errorf("invalid --output-location: %s", err.Error())
			return
		}
	}
	if len(_b2biTransformerId) > 0 {
		input.TransformerId = aws.String(_b2biTransformerId)
	}
	if len(_b2biClientToken) > 0 {
		input.ClientToken = aws.String(_b2biClientToken)
	}

	if resp, err := client.StartTransformerJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches a key-value pair to a resource, as identified by its Amazon Resource
// Name (ARN). Resources are capability, partnership, profile, transformers and
// other entities.
//
// There is no response returned from this call.
func b2bi_TagResource(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_b2biResourceARN) > 0 {
		input.ResourceARN = aws.String(_b2biResourceARN)
	}
	if len(_b2biTags) > 0 {
		if err := assignInputField(input, "Tags", _b2biTags); err != nil {
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

// This operation mimics the latter half of a typical Outbound EDI request. It
// takes an input JSON/XML in the B2Bi shape as input, converts it to an X12 EDI
// string, and return that string.
func b2bi_TestConversion(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.TestConversionInput{
		// Source: *types.ConversionSource, // Required
		// Target: *types.ConversionTarget, // Required
	}

	if len(_b2biSource) > 0 {
		if err := assignInputField(input, "Source", _b2biSource); err != nil {
			log.Errorf("invalid --source: %s", err.Error())
			return
		}
	}
	if len(_b2biTarget) > 0 {
		if err := assignInputField(input, "Target", _b2biTarget); err != nil {
			log.Errorf("invalid --target: %s", err.Error())
			return
		}
	}

	if resp, err := client.TestConversion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Maps the input file according to the provided template file. The API call
// downloads the file contents from the Amazon S3 location, and passes the contents
// in as a string, to the inputFileContent parameter.
func b2bi_TestMapping(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.TestMappingInput{
		// FileFormat: types.FileFormat, // Required
		// InputFileContent: *string, // Required
		// MappingTemplate: *string, // Required
	}

	if len(_b2biFileFormat) > 0 {
		if err := assignInputField(input, "FileFormat", _b2biFileFormat); err != nil {
			log.Errorf("invalid --file-format: %s", err.Error())
			return
		}
	}
	if len(_b2biInputFileContent) > 0 {
		input.InputFileContent = aws.String(_b2biInputFileContent)
	}
	if len(_b2biMappingTemplate) > 0 {
		input.MappingTemplate = aws.String(_b2biMappingTemplate)
	}

	if resp, err := client.TestMapping(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Parses the input EDI (electronic data interchange) file. The input file has a
// file size limit of 250 KB.
func b2bi_TestParsing(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.TestParsingInput{
		// EdiType: types.EdiType, // Required
		// FileFormat: types.FileFormat, // Required
		// InputFile: *types.S3Location, // Required
	}

	if len(_b2biEdiType) > 0 {
		if err := assignInputField(input, "EdiType", _b2biEdiType); err != nil {
			log.Errorf("invalid --edi-type: %s", err.Error())
			return
		}
	}
	if len(_b2biFileFormat) > 0 {
		if err := assignInputField(input, "FileFormat", _b2biFileFormat); err != nil {
			log.Errorf("invalid --file-format: %s", err.Error())
			return
		}
	}
	if len(_b2biInputFile) > 0 {
		if err := assignInputField(input, "InputFile", _b2biInputFile); err != nil {
			log.Errorf("invalid --input-file: %s", err.Error())
			return
		}
	}
	if len(_b2biAdvancedOptions) > 0 {
		if err := assignInputField(input, "AdvancedOptions", _b2biAdvancedOptions); err != nil {
			log.Errorf("invalid --advanced-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.TestParsing(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detaches a key-value pair from the specified resource, as identified by its
// Amazon Resource Name (ARN). Resources are capability, partnership, profile,
// transformers and other entities.
func b2bi_UntagResource(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_b2biResourceARN) > 0 {
		input.ResourceARN = aws.String(_b2biResourceARN)
	}
	if len(_b2biTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _b2biTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates some of the parameters for a capability, based on the specified
// parameters. A trading capability contains the information required to transform
// incoming EDI documents into JSON or XML outputs.
func b2bi_UpdateCapability(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.UpdateCapabilityInput{
		// CapabilityId: *string, // Required
	}

	if len(_b2biCapabilityId) > 0 {
		input.CapabilityId = aws.String(_b2biCapabilityId)
	}
	if len(_b2biConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _b2biConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_b2biInstructionsDocuments) > 0 {
		if err := assignInputField(input, "InstructionsDocuments", _b2biInstructionsDocuments); err != nil {
			log.Errorf("invalid --instructions-documents: %s", err.Error())
			return
		}
	}
	if len(_b2biName) > 0 {
		input.Name = aws.String(_b2biName)
	}

	if resp, err := client.UpdateCapability(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates some of the parameters for a partnership between a customer and trading
// partner. A partnership represents the connection between you and your trading
// partner. It ties together a profile and one or more trading capabilities.
func b2bi_UpdatePartnership(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.UpdatePartnershipInput{
		// PartnershipId: *string, // Required
	}

	if len(_b2biPartnershipId) > 0 {
		input.PartnershipId = aws.String(_b2biPartnershipId)
	}
	if len(_b2biCapabilities) > 0 {
		input.Capabilities = append([]string(nil), _b2biCapabilities...)
	}
	if len(_b2biCapabilityOptions) > 0 {
		if err := assignInputField(input, "CapabilityOptions", _b2biCapabilityOptions); err != nil {
			log.Errorf("invalid --capability-options: %s", err.Error())
			return
		}
	}
	if len(_b2biName) > 0 {
		input.Name = aws.String(_b2biName)
	}

	if resp, err := client.UpdatePartnership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified parameters for a profile. A profile is the mechanism used
// to create the concept of a private network.
func b2bi_UpdateProfile(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.UpdateProfileInput{
		// ProfileId: *string, // Required
	}

	if len(_b2biProfileId) > 0 {
		input.ProfileId = aws.String(_b2biProfileId)
	}
	if len(_b2biBusinessName) > 0 {
		input.BusinessName = aws.String(_b2biBusinessName)
	}
	if len(_b2biEmail) > 0 {
		input.Email = aws.String(_b2biEmail)
	}
	if len(_b2biName) > 0 {
		input.Name = aws.String(_b2biName)
	}
	if len(_b2biPhone) > 0 {
		input.Phone = aws.String(_b2biPhone)
	}

	if resp, err := client.UpdateProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified parameters for a transformer. A transformer can take an
// EDI file as input and transform it into a JSON-or XML-formatted document.
// Alternatively, a transformer can take a JSON-or XML-formatted document as input
// and transform it into an EDI file.
func b2bi_UpdateTransformer(cfg aws.Config, client *b2bi.Client) {
	input := &b2bi.UpdateTransformerInput{
		// TransformerId: *string, // Required
	}

	if len(_b2biTransformerId) > 0 {
		input.TransformerId = aws.String(_b2biTransformerId)
	}
	if len(_b2biEdiType) > 0 {
		if err := assignInputField(input, "EdiType", _b2biEdiType); err != nil {
			log.Errorf("invalid --edi-type: %s", err.Error())
			return
		}
	}
	if len(_b2biFileFormat) > 0 {
		if err := assignInputField(input, "FileFormat", _b2biFileFormat); err != nil {
			log.Errorf("invalid --file-format: %s", err.Error())
			return
		}
	}
	if len(_b2biInputConversion) > 0 {
		if err := assignInputField(input, "InputConversion", _b2biInputConversion); err != nil {
			log.Errorf("invalid --input-conversion: %s", err.Error())
			return
		}
	}
	if len(_b2biMapping) > 0 {
		if err := assignInputField(input, "Mapping", _b2biMapping); err != nil {
			log.Errorf("invalid --mapping: %s", err.Error())
			return
		}
	}
	if len(_b2biMappingTemplate) > 0 {
		input.MappingTemplate = aws.String(_b2biMappingTemplate)
	}
	if len(_b2biName) > 0 {
		input.Name = aws.String(_b2biName)
	}
	if len(_b2biOutputConversion) > 0 {
		if err := assignInputField(input, "OutputConversion", _b2biOutputConversion); err != nil {
			log.Errorf("invalid --output-conversion: %s", err.Error())
			return
		}
	}
	if len(_b2biSampleDocument) > 0 {
		input.SampleDocument = aws.String(_b2biSampleDocument)
	}
	if len(_b2biSampleDocuments) > 0 {
		if err := assignInputField(input, "SampleDocuments", _b2biSampleDocuments); err != nil {
			log.Errorf("invalid --sample-documents: %s", err.Error())
			return
		}
	}
	if len(_b2biStatus) > 0 {
		if err := assignInputField(input, "Status", _b2biStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTransformer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_b2biCmd)
	_b2biCmd.Flags().SortFlags = false

	_b2biCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_b2biCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_b2biCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_b2biCmd.Flags().StringVarP(&_b2biAdvancedOptions, "advanced-options", "", "", "Advanced Options")
	_b2biCmd.Flags().StringVarP(&_b2biBusinessName, "business-name", "", "", "Business Name")
	_b2biCmd.Flags().StringSliceVarP(&_b2biCapabilities, "capabilities", "", nil, "Capabilities")
	_b2biCmd.Flags().StringVarP(&_b2biCapabilityId, "capability-id", "", "", "Capability ID")
	_b2biCmd.Flags().StringVarP(&_b2biCapabilityOptions, "capability-options", "", "", "Capability Options")
	_b2biCmd.Flags().StringVarP(&_b2biClientToken, "client-token", "", "", "Client Token")
	_b2biCmd.Flags().StringVarP(&_b2biConfiguration, "configuration", "", "", "Configuration")
	_b2biCmd.Flags().StringVarP(&_b2biEdiType, "edi-type", "", "", "Edi Type")
	_b2biCmd.Flags().StringVarP(&_b2biEmail, "email", "", "", "Email")
	_b2biCmd.Flags().StringVarP(&_b2biFileFormat, "file-format", "", "", "File Format")
	_b2biCmd.Flags().StringVarP(&_b2biInputConversion, "input-conversion", "", "", "Input Conversion")
	_b2biCmd.Flags().StringVarP(&_b2biInputFile, "input-file", "", "", "Input File")
	_b2biCmd.Flags().StringVarP(&_b2biInputFileContent, "input-file-content", "", "", "Input File Content")
	_b2biCmd.Flags().StringVarP(&_b2biInstructionsDocuments, "instructions-documents", "", "", "Instructions Documents")
	_b2biCmd.Flags().StringVarP(&_b2biLogging, "logging", "", "", "Logging")
	_b2biCmd.Flags().StringVarP(&_b2biMapping, "mapping", "", "", "Mapping")
	_b2biCmd.Flags().StringVarP(&_b2biMappingTemplate, "mapping-template", "", "", "Mapping Template")
	_b2biCmd.Flags().StringVarP(&_b2biMappingType, "mapping-type", "", "", "Mapping Type")
	_b2biCmd.Flags().StringVarP(&_b2biMaxResults, "max-results", "", "", "Max Results")
	_b2biCmd.Flags().StringVarP(&_b2biName, "name", "", "", "Name")
	_b2biCmd.Flags().StringVarP(&_b2biNextToken, "next-token", "", "", "Next Token")
	_b2biCmd.Flags().StringVarP(&_b2biOutputConversion, "output-conversion", "", "", "Output Conversion")
	_b2biCmd.Flags().StringVarP(&_b2biOutputFileContent, "output-file-content", "", "", "Output File Content")
	_b2biCmd.Flags().StringVarP(&_b2biOutputLocation, "output-location", "", "", "Output Location")
	_b2biCmd.Flags().StringVarP(&_b2biOutputSampleLocation, "output-sample-location", "", "", "Output Sample Location")
	_b2biCmd.Flags().StringVarP(&_b2biPartnershipId, "partnership-id", "", "", "Partnership ID")
	_b2biCmd.Flags().StringVarP(&_b2biPhone, "phone", "", "", "Phone")
	_b2biCmd.Flags().StringVarP(&_b2biProfileId, "profile-id", "", "", "Profile ID")
	_b2biCmd.Flags().StringVarP(&_b2biResourceARN, "resource-arn", "", "", "Resource ARN")
	_b2biCmd.Flags().StringVarP(&_b2biSampleDocument, "sample-document", "", "", "Sample Document")
	_b2biCmd.Flags().StringVarP(&_b2biSampleDocuments, "sample-documents", "", "", "Sample Documents")
	_b2biCmd.Flags().StringVarP(&_b2biSource, "source", "", "", "Source")
	_b2biCmd.Flags().StringVarP(&_b2biStatus, "status", "", "", "Status")
	_b2biCmd.Flags().StringSliceVarP(&_b2biTagKeys, "tag-keys", "", nil, "Tag Keys")
	_b2biCmd.Flags().StringVarP(&_b2biTags, "tags", "", "", "Tags")
	_b2biCmd.Flags().StringVarP(&_b2biTarget, "target", "", "", "Target")
	_b2biCmd.Flags().StringVarP(&_b2biTemplateDetails, "template-details", "", "", "Template Details")
	_b2biCmd.Flags().StringVarP(&_b2biTransformerId, "transformer-id", "", "", "Transformer ID")
	_b2biCmd.Flags().StringVarP(&_b2biTransformerJobId, "transformer-job-id", "", "", "Transformer Job ID")
	_b2biCmd.Flags().StringVarP(&_b2biType, "type", "", "", "Type")

	_b2biCmd.Flags().BoolVarP(&_b2biCreateCapability, "create-capability", "", false, "Create Capability")
	_b2biCmd.Flags().BoolVarP(&_b2biCreatePartnership, "create-partnership", "", false, "Create Partnership")
	_b2biCmd.Flags().BoolVarP(&_b2biCreateProfile, "create-profile", "", false, "Create Profile")
	_b2biCmd.Flags().BoolVarP(&_b2biCreateStarterMappingTemplate, "create-starter-mapping-template", "", false, "Create Starter Mapping Template")
	_b2biCmd.Flags().BoolVarP(&_b2biCreateTransformer, "create-transformer", "", false, "Create Transformer")
	_b2biCmd.Flags().BoolVarP(&_b2biDeleteCapability, "delete-capability", "", false, "Delete Capability")
	_b2biCmd.Flags().BoolVarP(&_b2biDeletePartnership, "delete-partnership", "", false, "Delete Partnership")
	_b2biCmd.Flags().BoolVarP(&_b2biDeleteProfile, "delete-profile", "", false, "Delete Profile")
	_b2biCmd.Flags().BoolVarP(&_b2biDeleteTransformer, "delete-transformer", "", false, "Delete Transformer")
	_b2biCmd.Flags().BoolVarP(&_b2biGenerateMapping, "generate-mapping", "", false, "Generate Mapping")
	_b2biCmd.Flags().BoolVarP(&_b2biGetCapability, "get-capability", "", false, "Get Capability")
	_b2biCmd.Flags().BoolVarP(&_b2biGetPartnership, "get-partnership", "", false, "Get Partnership")
	_b2biCmd.Flags().BoolVarP(&_b2biGetProfile, "get-profile", "", false, "Get Profile")
	_b2biCmd.Flags().BoolVarP(&_b2biGetTransformer, "get-transformer", "", false, "Get Transformer")
	_b2biCmd.Flags().BoolVarP(&_b2biGetTransformerJob, "get-transformer-job", "", false, "Get Transformer Job")
	_b2biCmd.Flags().BoolVarP(&_b2biListCapabilities, "list-capabilities", "", false, "List Capabilities")
	_b2biCmd.Flags().BoolVarP(&_b2biListPartnerships, "list-partnerships", "", false, "List Partnerships")
	_b2biCmd.Flags().BoolVarP(&_b2biListProfiles, "list-profiles", "", false, "List Profiles")
	_b2biCmd.Flags().BoolVarP(&_b2biListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_b2biCmd.Flags().BoolVarP(&_b2biListTransformers, "list-transformers", "", false, "List Transformers")
	_b2biCmd.Flags().BoolVarP(&_b2biStartTransformerJob, "start-transformer-job", "", false, "Start Transformer Job")
	_b2biCmd.Flags().BoolVarP(&_b2biTagResource, "tag-resource", "", false, "Tag Resource")
	_b2biCmd.Flags().BoolVarP(&_b2biTestConversion, "test-conversion", "", false, "Test Conversion")
	_b2biCmd.Flags().BoolVarP(&_b2biTestMapping, "test-mapping", "", false, "Test Mapping")
	_b2biCmd.Flags().BoolVarP(&_b2biTestParsing, "test-parsing", "", false, "Test Parsing")
	_b2biCmd.Flags().BoolVarP(&_b2biUntagResource, "untag-resource", "", false, "Untag Resource")
	_b2biCmd.Flags().BoolVarP(&_b2biUpdateCapability, "update-capability", "", false, "Update Capability")
	_b2biCmd.Flags().BoolVarP(&_b2biUpdatePartnership, "update-partnership", "", false, "Update Partnership")
	_b2biCmd.Flags().BoolVarP(&_b2biUpdateProfile, "update-profile", "", false, "Update Profile")
	_b2biCmd.Flags().BoolVarP(&_b2biUpdateTransformer, "update-transformer", "", false, "Update Transformer")

}
