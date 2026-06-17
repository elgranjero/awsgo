package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pcaconnectorscep"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// pcaconnectorscepCmd represents the pcaconnectorscep command
var _pcaconnectorscepCmd = &cobra.Command{
	Use:   "pcaconnectorscep",
	Short: "AWS pcaconnectorscep CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := pcaconnectorscep.NewFromConfig(cfg)
		if _pcaconnectorscepCreateChallenge {
			pcaconnectorscep_CreateChallenge(cfg, client)
			return
		}
		if _pcaconnectorscepCreateConnector {
			pcaconnectorscep_CreateConnector(cfg, client)
			return
		}
		if _pcaconnectorscepDeleteChallenge {
			pcaconnectorscep_DeleteChallenge(cfg, client)
			return
		}
		if _pcaconnectorscepDeleteConnector {
			pcaconnectorscep_DeleteConnector(cfg, client)
			return
		}
		if _pcaconnectorscepGetChallengeMetadata {
			pcaconnectorscep_GetChallengeMetadata(cfg, client)
			return
		}
		if _pcaconnectorscepGetChallengePassword {
			pcaconnectorscep_GetChallengePassword(cfg, client)
			return
		}
		if _pcaconnectorscepGetConnector {
			pcaconnectorscep_GetConnector(cfg, client)
			return
		}
		if _pcaconnectorscepListChallengeMetadata {
			pcaconnectorscep_ListChallengeMetadata(cfg, client)
			return
		}
		if _pcaconnectorscepListConnectors {
			pcaconnectorscep_ListConnectors(cfg, client)
			return
		}
		if _pcaconnectorscepListTagsForResource {
			pcaconnectorscep_ListTagsForResource(cfg, client)
			return
		}
		if _pcaconnectorscepTagResource {
			pcaconnectorscep_TagResource(cfg, client)
			return
		}
		if _pcaconnectorscepUntagResource {
			pcaconnectorscep_UntagResource(cfg, client)
			return
		}

	},
}

var (
	_pcaconnectorscepCreateChallenge       bool
	_pcaconnectorscepCreateConnector       bool
	_pcaconnectorscepDeleteChallenge       bool
	_pcaconnectorscepDeleteConnector       bool
	_pcaconnectorscepGetChallengeMetadata  bool
	_pcaconnectorscepGetChallengePassword  bool
	_pcaconnectorscepGetConnector          bool
	_pcaconnectorscepListChallengeMetadata bool
	_pcaconnectorscepListConnectors        bool
	_pcaconnectorscepListTagsForResource   bool
	_pcaconnectorscepTagResource           bool
	_pcaconnectorscepUntagResource         bool

	_pcaconnectorscepCertificateAuthorityArn string
	_pcaconnectorscepChallengeArn            string
	_pcaconnectorscepClientToken             string
	_pcaconnectorscepConnectorArn            string
	_pcaconnectorscepMaxResults              string
	_pcaconnectorscepMobileDeviceManagement  string
	_pcaconnectorscepNextToken               string
	_pcaconnectorscepResourceArn             string
	_pcaconnectorscepTagKeys                 []string
	_pcaconnectorscepTags                    string
	_pcaconnectorscepVpcEndpointId           string
)

// For general-purpose connectors. Creates a challenge password for the specified
// connector. The SCEP protocol uses a challenge password to authenticate a request
// before issuing a certificate from a certificate authority (CA). Your SCEP
// clients include the challenge password as part of their certificate request to
// Connector for SCEP. To retrieve the connector Amazon Resource Names (ARNs) for
// the connectors in your account, call [ListConnectors].
//
// To create additional challenge passwords for the connector, call CreateChallenge
// again. We recommend frequently rotating your challenge passwords.
//
// [ListConnectors]: https://docs.aws.amazon.com/pca-connector-scep/latest/APIReference/API_ListConnectors.html
func pcaconnectorscep_CreateChallenge(cfg aws.Config, client *pcaconnectorscep.Client) {
	input := &pcaconnectorscep.CreateChallengeInput{
		// ConnectorArn: *string, // Required
	}

	if len(_pcaconnectorscepConnectorArn) > 0 {
		input.ConnectorArn = aws.String(_pcaconnectorscepConnectorArn)
	}
	if len(_pcaconnectorscepClientToken) > 0 {
		input.ClientToken = aws.String(_pcaconnectorscepClientToken)
	}
	if len(_pcaconnectorscepTags) > 0 {
		if err := assignInputField(input, "Tags", _pcaconnectorscepTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateChallenge(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a SCEP connector. A SCEP connector links Amazon Web Services Private
// Certificate Authority to your SCEP-compatible devices and mobile device
// management (MDM) systems. Before you create a connector, you must complete a set
// of prerequisites, including creation of a private certificate authority (CA) to
// use with this connector. For more information, see [Connector for SCEP prerequisites].
//
// [Connector for SCEP prerequisites]: https://docs.aws.amazon.com/privateca/latest/userguide/scep-connector.htmlconnector-for-scep-prerequisites.html
func pcaconnectorscep_CreateConnector(cfg aws.Config, client *pcaconnectorscep.Client) {
	input := &pcaconnectorscep.CreateConnectorInput{
		// CertificateAuthorityArn: *string, // Required
	}

	if len(_pcaconnectorscepCertificateAuthorityArn) > 0 {
		input.CertificateAuthorityArn = aws.String(_pcaconnectorscepCertificateAuthorityArn)
	}
	if len(_pcaconnectorscepClientToken) > 0 {
		input.ClientToken = aws.String(_pcaconnectorscepClientToken)
	}
	if len(_pcaconnectorscepMobileDeviceManagement) > 0 {
		if err := assignInputField(input, "MobileDeviceManagement", _pcaconnectorscepMobileDeviceManagement); err != nil {
			log.Errorf("invalid --mobile-device-management: %s", err.Error())
			return
		}
	}
	if len(_pcaconnectorscepTags) > 0 {
		if err := assignInputField(input, "Tags", _pcaconnectorscepTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_pcaconnectorscepVpcEndpointId) > 0 {
		input.VpcEndpointId = aws.String(_pcaconnectorscepVpcEndpointId)
	}

	if resp, err := client.CreateConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified [Challenge].
//
// [Challenge]: https://docs.aws.amazon.com/pca-connector-scep/latest/APIReference/API_Challenge.html
func pcaconnectorscep_DeleteChallenge(cfg aws.Config, client *pcaconnectorscep.Client) {
	input := &pcaconnectorscep.DeleteChallengeInput{
		// ChallengeArn: *string, // Required
	}

	if len(_pcaconnectorscepChallengeArn) > 0 {
		input.ChallengeArn = aws.String(_pcaconnectorscepChallengeArn)
	}

	if resp, err := client.DeleteChallenge(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified [Connector]. This operation also deletes any challenges associated
// with the connector.
//
// [Connector]: https://docs.aws.amazon.com/pca-connector-scep/latest/APIReference/API_Connector.html
func pcaconnectorscep_DeleteConnector(cfg aws.Config, client *pcaconnectorscep.Client) {
	input := &pcaconnectorscep.DeleteConnectorInput{
		// ConnectorArn: *string, // Required
	}

	if len(_pcaconnectorscepConnectorArn) > 0 {
		input.ConnectorArn = aws.String(_pcaconnectorscepConnectorArn)
	}

	if resp, err := client.DeleteConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the metadata for the specified [Challenge].
//
// [Challenge]: https://docs.aws.amazon.com/pca-connector-scep/latest/APIReference/API_Challenge.html
func pcaconnectorscep_GetChallengeMetadata(cfg aws.Config, client *pcaconnectorscep.Client) {
	input := &pcaconnectorscep.GetChallengeMetadataInput{
		// ChallengeArn: *string, // Required
	}

	if len(_pcaconnectorscepChallengeArn) > 0 {
		input.ChallengeArn = aws.String(_pcaconnectorscepChallengeArn)
	}

	if resp, err := client.GetChallengeMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the challenge password for the specified [Challenge].
//
// [Challenge]: https://docs.aws.amazon.com/pca-connector-scep/latest/APIReference/API_Challenge.html
func pcaconnectorscep_GetChallengePassword(cfg aws.Config, client *pcaconnectorscep.Client) {
	input := &pcaconnectorscep.GetChallengePasswordInput{
		// ChallengeArn: *string, // Required
	}

	if len(_pcaconnectorscepChallengeArn) > 0 {
		input.ChallengeArn = aws.String(_pcaconnectorscepChallengeArn)
	}

	if resp, err := client.GetChallengePassword(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details about the specified [Connector]. Calling this action returns important
// details about the connector, such as the public SCEP URL where your clients can
// request certificates.
//
// [Connector]: https://docs.aws.amazon.com/pca-connector-scep/latest/APIReference/API_Connector.html
func pcaconnectorscep_GetConnector(cfg aws.Config, client *pcaconnectorscep.Client) {
	input := &pcaconnectorscep.GetConnectorInput{
		// ConnectorArn: *string, // Required
	}

	if len(_pcaconnectorscepConnectorArn) > 0 {
		input.ConnectorArn = aws.String(_pcaconnectorscepConnectorArn)
	}

	if resp, err := client.GetConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the challenge metadata for the specified ARN.
func pcaconnectorscep_ListChallengeMetadata(cfg aws.Config, client *pcaconnectorscep.Client) {
	input := &pcaconnectorscep.ListChallengeMetadataInput{
		// ConnectorArn: *string, // Required
	}

	if len(_pcaconnectorscepConnectorArn) > 0 {
		input.ConnectorArn = aws.String(_pcaconnectorscepConnectorArn)
	}
	if len(_pcaconnectorscepMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pcaconnectorscepMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pcaconnectorscepNextToken) > 0 {
		input.NextToken = aws.String(_pcaconnectorscepNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListChallengeMetadata(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pcaconnectorscep.ListChallengeMetadataOutput
	p := pcaconnectorscep.NewListChallengeMetadataPaginator(client, input)
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

// Lists the connectors belonging to your Amazon Web Services account.
func pcaconnectorscep_ListConnectors(cfg aws.Config, client *pcaconnectorscep.Client) {
	input := &pcaconnectorscep.ListConnectorsInput{}

	if len(_pcaconnectorscepMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pcaconnectorscepMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pcaconnectorscepNextToken) > 0 {
		input.NextToken = aws.String(_pcaconnectorscepNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListConnectors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pcaconnectorscep.ListConnectorsOutput
	p := pcaconnectorscep.NewListConnectorsPaginator(client, input)
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

// Retrieves the tags associated with the specified resource. Tags are key-value
// pairs that you can use to categorize and manage your resources, for purposes
// like billing. For example, you might set the tag key to "customer" and the value
// to the customer name or ID. You can specify one or more tags to add to each
// Amazon Web Services resource, up to 50 tags for a resource.
func pcaconnectorscep_ListTagsForResource(cfg aws.Config, client *pcaconnectorscep.Client) {
	input := &pcaconnectorscep.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_pcaconnectorscepResourceArn) > 0 {
		input.ResourceArn = aws.String(_pcaconnectorscepResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more tags to your resource.
func pcaconnectorscep_TagResource(cfg aws.Config, client *pcaconnectorscep.Client) {
	input := &pcaconnectorscep.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_pcaconnectorscepResourceArn) > 0 {
		input.ResourceArn = aws.String(_pcaconnectorscepResourceArn)
	}
	if len(_pcaconnectorscepTags) > 0 {
		if err := assignInputField(input, "Tags", _pcaconnectorscepTags); err != nil {
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

// Removes one or more tags from your resource.
func pcaconnectorscep_UntagResource(cfg aws.Config, client *pcaconnectorscep.Client) {
	input := &pcaconnectorscep.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_pcaconnectorscepResourceArn) > 0 {
		input.ResourceArn = aws.String(_pcaconnectorscepResourceArn)
	}
	if len(_pcaconnectorscepTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _pcaconnectorscepTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_pcaconnectorscepCmd)
	_pcaconnectorscepCmd.Flags().SortFlags = false

	_pcaconnectorscepCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_pcaconnectorscepCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_pcaconnectorscepCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_pcaconnectorscepCmd.Flags().StringVarP(&_pcaconnectorscepCertificateAuthorityArn, "certificate-authority-arn", "", "", "Certificate Authority ARN")
	_pcaconnectorscepCmd.Flags().StringVarP(&_pcaconnectorscepChallengeArn, "challenge-arn", "", "", "Challenge ARN")
	_pcaconnectorscepCmd.Flags().StringVarP(&_pcaconnectorscepClientToken, "client-token", "", "", "Client Token")
	_pcaconnectorscepCmd.Flags().StringVarP(&_pcaconnectorscepConnectorArn, "connector-arn", "", "", "Connector ARN")
	_pcaconnectorscepCmd.Flags().StringVarP(&_pcaconnectorscepMaxResults, "max-results", "", "", "Max Results")
	_pcaconnectorscepCmd.Flags().StringVarP(&_pcaconnectorscepMobileDeviceManagement, "mobile-device-management", "", "", "Mobile Device Management")
	_pcaconnectorscepCmd.Flags().StringVarP(&_pcaconnectorscepNextToken, "next-token", "", "", "Next Token")
	_pcaconnectorscepCmd.Flags().StringVarP(&_pcaconnectorscepResourceArn, "resource-arn", "", "", "Resource ARN")
	_pcaconnectorscepCmd.Flags().StringSliceVarP(&_pcaconnectorscepTagKeys, "tag-keys", "", nil, "Tag Keys")
	_pcaconnectorscepCmd.Flags().StringVarP(&_pcaconnectorscepTags, "tags", "", "", "Tags")
	_pcaconnectorscepCmd.Flags().StringVarP(&_pcaconnectorscepVpcEndpointId, "vpc-endpoint-id", "", "", "VPC Endpoint ID")

	_pcaconnectorscepCmd.Flags().BoolVarP(&_pcaconnectorscepCreateChallenge, "create-challenge", "", false, "Create Challenge")
	_pcaconnectorscepCmd.Flags().BoolVarP(&_pcaconnectorscepCreateConnector, "create-connector", "", false, "Create Connector")
	_pcaconnectorscepCmd.Flags().BoolVarP(&_pcaconnectorscepDeleteChallenge, "delete-challenge", "", false, "Delete Challenge")
	_pcaconnectorscepCmd.Flags().BoolVarP(&_pcaconnectorscepDeleteConnector, "delete-connector", "", false, "Delete Connector")
	_pcaconnectorscepCmd.Flags().BoolVarP(&_pcaconnectorscepGetChallengeMetadata, "get-challenge-metadata", "", false, "Get Challenge Metadata")
	_pcaconnectorscepCmd.Flags().BoolVarP(&_pcaconnectorscepGetChallengePassword, "get-challenge-password", "", false, "Get Challenge Password")
	_pcaconnectorscepCmd.Flags().BoolVarP(&_pcaconnectorscepGetConnector, "get-connector", "", false, "Get Connector")
	_pcaconnectorscepCmd.Flags().BoolVarP(&_pcaconnectorscepListChallengeMetadata, "list-challenge-metadata", "", false, "List Challenge Metadata")
	_pcaconnectorscepCmd.Flags().BoolVarP(&_pcaconnectorscepListConnectors, "list-connectors", "", false, "List Connectors")
	_pcaconnectorscepCmd.Flags().BoolVarP(&_pcaconnectorscepListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_pcaconnectorscepCmd.Flags().BoolVarP(&_pcaconnectorscepTagResource, "tag-resource", "", false, "Tag Resource")
	_pcaconnectorscepCmd.Flags().BoolVarP(&_pcaconnectorscepUntagResource, "untag-resource", "", false, "Untag Resource")

}
