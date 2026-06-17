package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pcaconnectorad"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// pcaconnectoradCmd represents the pcaconnectorad command
var _pcaconnectoradCmd = &cobra.Command{
	Use:   "pcaconnectorad",
	Short: "AWS pcaconnectorad CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := pcaconnectorad.NewFromConfig(cfg)
		if _pcaconnectoradCreateConnector {
			pcaconnectorad_CreateConnector(cfg, client)
			return
		}
		if _pcaconnectoradCreateDirectoryRegistration {
			pcaconnectorad_CreateDirectoryRegistration(cfg, client)
			return
		}
		if _pcaconnectoradCreateServicePrincipalName {
			pcaconnectorad_CreateServicePrincipalName(cfg, client)
			return
		}
		if _pcaconnectoradCreateTemplate {
			pcaconnectorad_CreateTemplate(cfg, client)
			return
		}
		if _pcaconnectoradCreateTemplateGroupAccessControlEntry {
			pcaconnectorad_CreateTemplateGroupAccessControlEntry(cfg, client)
			return
		}
		if _pcaconnectoradDeleteConnector {
			pcaconnectorad_DeleteConnector(cfg, client)
			return
		}
		if _pcaconnectoradDeleteDirectoryRegistration {
			pcaconnectorad_DeleteDirectoryRegistration(cfg, client)
			return
		}
		if _pcaconnectoradDeleteServicePrincipalName {
			pcaconnectorad_DeleteServicePrincipalName(cfg, client)
			return
		}
		if _pcaconnectoradDeleteTemplate {
			pcaconnectorad_DeleteTemplate(cfg, client)
			return
		}
		if _pcaconnectoradDeleteTemplateGroupAccessControlEntry {
			pcaconnectorad_DeleteTemplateGroupAccessControlEntry(cfg, client)
			return
		}
		if _pcaconnectoradGetConnector {
			pcaconnectorad_GetConnector(cfg, client)
			return
		}
		if _pcaconnectoradGetDirectoryRegistration {
			pcaconnectorad_GetDirectoryRegistration(cfg, client)
			return
		}
		if _pcaconnectoradGetServicePrincipalName {
			pcaconnectorad_GetServicePrincipalName(cfg, client)
			return
		}
		if _pcaconnectoradGetTemplate {
			pcaconnectorad_GetTemplate(cfg, client)
			return
		}
		if _pcaconnectoradGetTemplateGroupAccessControlEntry {
			pcaconnectorad_GetTemplateGroupAccessControlEntry(cfg, client)
			return
		}
		if _pcaconnectoradListConnectors {
			pcaconnectorad_ListConnectors(cfg, client)
			return
		}
		if _pcaconnectoradListDirectoryRegistrations {
			pcaconnectorad_ListDirectoryRegistrations(cfg, client)
			return
		}
		if _pcaconnectoradListServicePrincipalNames {
			pcaconnectorad_ListServicePrincipalNames(cfg, client)
			return
		}
		if _pcaconnectoradListTagsForResource {
			pcaconnectorad_ListTagsForResource(cfg, client)
			return
		}
		if _pcaconnectoradListTemplateGroupAccessControlEntries {
			pcaconnectorad_ListTemplateGroupAccessControlEntries(cfg, client)
			return
		}
		if _pcaconnectoradListTemplates {
			pcaconnectorad_ListTemplates(cfg, client)
			return
		}
		if _pcaconnectoradTagResource {
			pcaconnectorad_TagResource(cfg, client)
			return
		}
		if _pcaconnectoradUntagResource {
			pcaconnectorad_UntagResource(cfg, client)
			return
		}
		if _pcaconnectoradUpdateTemplate {
			pcaconnectorad_UpdateTemplate(cfg, client)
			return
		}
		if _pcaconnectoradUpdateTemplateGroupAccessControlEntry {
			pcaconnectorad_UpdateTemplateGroupAccessControlEntry(cfg, client)
			return
		}

	},
}

var (
	_pcaconnectoradCreateConnector                       bool
	_pcaconnectoradCreateDirectoryRegistration           bool
	_pcaconnectoradCreateServicePrincipalName            bool
	_pcaconnectoradCreateTemplate                        bool
	_pcaconnectoradCreateTemplateGroupAccessControlEntry bool
	_pcaconnectoradDeleteConnector                       bool
	_pcaconnectoradDeleteDirectoryRegistration           bool
	_pcaconnectoradDeleteServicePrincipalName            bool
	_pcaconnectoradDeleteTemplate                        bool
	_pcaconnectoradDeleteTemplateGroupAccessControlEntry bool
	_pcaconnectoradGetConnector                          bool
	_pcaconnectoradGetDirectoryRegistration              bool
	_pcaconnectoradGetServicePrincipalName               bool
	_pcaconnectoradGetTemplate                           bool
	_pcaconnectoradGetTemplateGroupAccessControlEntry    bool
	_pcaconnectoradListConnectors                        bool
	_pcaconnectoradListDirectoryRegistrations            bool
	_pcaconnectoradListServicePrincipalNames             bool
	_pcaconnectoradListTagsForResource                   bool
	_pcaconnectoradListTemplateGroupAccessControlEntries bool
	_pcaconnectoradListTemplates                         bool
	_pcaconnectoradTagResource                           bool
	_pcaconnectoradUntagResource                         bool
	_pcaconnectoradUpdateTemplate                        bool
	_pcaconnectoradUpdateTemplateGroupAccessControlEntry bool

	_pcaconnectoradAccessRights                  string
	_pcaconnectoradCertificateAuthorityArn       string
	_pcaconnectoradClientToken                   string
	_pcaconnectoradConnectorArn                  string
	_pcaconnectoradDefinition                    string
	_pcaconnectoradDirectoryId                   string
	_pcaconnectoradDirectoryRegistrationArn      string
	_pcaconnectoradGroupDisplayName              string
	_pcaconnectoradGroupSecurityIdentifier       string
	_pcaconnectoradMaxResults                    string
	_pcaconnectoradName                          string
	_pcaconnectoradNextToken                     string
	_pcaconnectoradReenrollAllCertificateHolders string
	_pcaconnectoradResourceArn                   string
	_pcaconnectoradTagKeys                       []string
	_pcaconnectoradTags                          string
	_pcaconnectoradTemplateArn                   string
	_pcaconnectoradVpcInformation                string
)

// Creates a connector between Amazon Web Services Private CA and an Active
// Directory. You must specify the private CA, directory ID, and security groups.
func pcaconnectorad_CreateConnector(cfg aws.Config, client *pcaconnectorad.Client) {
	input := &pcaconnectorad.CreateConnectorInput{
		// CertificateAuthorityArn: *string, // Required
		// DirectoryId: *string, // Required
		// VpcInformation: *types.VpcInformation, // Required
	}

	if len(_pcaconnectoradCertificateAuthorityArn) > 0 {
		input.CertificateAuthorityArn = aws.String(_pcaconnectoradCertificateAuthorityArn)
	}
	if len(_pcaconnectoradDirectoryId) > 0 {
		input.DirectoryId = aws.String(_pcaconnectoradDirectoryId)
	}
	if len(_pcaconnectoradVpcInformation) > 0 {
		if err := assignInputField(input, "VpcInformation", _pcaconnectoradVpcInformation); err != nil {
			log.Errorf("invalid --vpc-information: %s", err.Error())
			return
		}
	}
	if len(_pcaconnectoradClientToken) > 0 {
		input.ClientToken = aws.String(_pcaconnectoradClientToken)
	}
	if len(_pcaconnectoradTags) > 0 {
		if err := assignInputField(input, "Tags", _pcaconnectoradTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a directory registration that authorizes communication between Amazon
// Web Services Private CA and an Active Directory
func pcaconnectorad_CreateDirectoryRegistration(cfg aws.Config, client *pcaconnectorad.Client) {
	input := &pcaconnectorad.CreateDirectoryRegistrationInput{
		// DirectoryId: *string, // Required
	}

	if len(_pcaconnectoradDirectoryId) > 0 {
		input.DirectoryId = aws.String(_pcaconnectoradDirectoryId)
	}
	if len(_pcaconnectoradClientToken) > 0 {
		input.ClientToken = aws.String(_pcaconnectoradClientToken)
	}
	if len(_pcaconnectoradTags) > 0 {
		if err := assignInputField(input, "Tags", _pcaconnectoradTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDirectoryRegistration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a service principal name (SPN) for the service account in Active
// Directory. Kerberos authentication uses SPNs to associate a service instance
// with a service sign-in account.
func pcaconnectorad_CreateServicePrincipalName(cfg aws.Config, client *pcaconnectorad.Client) {
	input := &pcaconnectorad.CreateServicePrincipalNameInput{
		// ConnectorArn: *string, // Required
		// DirectoryRegistrationArn: *string, // Required
	}

	if len(_pcaconnectoradConnectorArn) > 0 {
		input.ConnectorArn = aws.String(_pcaconnectoradConnectorArn)
	}
	if len(_pcaconnectoradDirectoryRegistrationArn) > 0 {
		input.DirectoryRegistrationArn = aws.String(_pcaconnectoradDirectoryRegistrationArn)
	}
	if len(_pcaconnectoradClientToken) > 0 {
		input.ClientToken = aws.String(_pcaconnectoradClientToken)
	}

	if resp, err := client.CreateServicePrincipalName(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Active Directory compatible certificate template. The connectors
// issues certificates using these templates based on the requester’s Active
// Directory group membership.
func pcaconnectorad_CreateTemplate(cfg aws.Config, client *pcaconnectorad.Client) {
	input := &pcaconnectorad.CreateTemplateInput{
		// ConnectorArn: *string, // Required
		// Definition: types.TemplateDefinition, // Required
		// Name: *string, // Required
	}

	if len(_pcaconnectoradConnectorArn) > 0 {
		input.ConnectorArn = aws.String(_pcaconnectoradConnectorArn)
	}
	if len(_pcaconnectoradDefinition) > 0 {
		if err := assignInputField(input, "Definition", _pcaconnectoradDefinition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}
	if len(_pcaconnectoradName) > 0 {
		input.Name = aws.String(_pcaconnectoradName)
	}
	if len(_pcaconnectoradClientToken) > 0 {
		input.ClientToken = aws.String(_pcaconnectoradClientToken)
	}
	if len(_pcaconnectoradTags) > 0 {
		if err := assignInputField(input, "Tags", _pcaconnectoradTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a group access control entry. Allow or deny Active Directory groups from
// enrolling and/or autoenrolling with the template based on the group security
// identifiers (SIDs).
func pcaconnectorad_CreateTemplateGroupAccessControlEntry(cfg aws.Config, client *pcaconnectorad.Client) {
	input := &pcaconnectorad.CreateTemplateGroupAccessControlEntryInput{
		// AccessRights: *types.AccessRights, // Required
		// GroupDisplayName: *string, // Required
		// GroupSecurityIdentifier: *string, // Required
		// TemplateArn: *string, // Required
	}

	if len(_pcaconnectoradAccessRights) > 0 {
		if err := assignInputField(input, "AccessRights", _pcaconnectoradAccessRights); err != nil {
			log.Errorf("invalid --access-rights: %s", err.Error())
			return
		}
	}
	if len(_pcaconnectoradGroupDisplayName) > 0 {
		input.GroupDisplayName = aws.String(_pcaconnectoradGroupDisplayName)
	}
	if len(_pcaconnectoradGroupSecurityIdentifier) > 0 {
		input.GroupSecurityIdentifier = aws.String(_pcaconnectoradGroupSecurityIdentifier)
	}
	if len(_pcaconnectoradTemplateArn) > 0 {
		input.TemplateArn = aws.String(_pcaconnectoradTemplateArn)
	}
	if len(_pcaconnectoradClientToken) > 0 {
		input.ClientToken = aws.String(_pcaconnectoradClientToken)
	}

	if resp, err := client.CreateTemplateGroupAccessControlEntry(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a connector for Active Directory. You must provide the Amazon Resource
// Name (ARN) of the connector that you want to delete. You can find the ARN by
// calling the [https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_ListConnectors]action. Deleting a connector does not deregister your directory
// with Amazon Web Services Private CA. You can deregister your directory by
// calling the [https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_DeleteDirectoryRegistration]action.
//
// [https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_ListConnectors]: https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_ListConnectors
// [https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_DeleteDirectoryRegistration]: https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_DeleteDirectoryRegistration
func pcaconnectorad_DeleteConnector(cfg aws.Config, client *pcaconnectorad.Client) {
	input := &pcaconnectorad.DeleteConnectorInput{
		// ConnectorArn: *string, // Required
	}

	if len(_pcaconnectoradConnectorArn) > 0 {
		input.ConnectorArn = aws.String(_pcaconnectoradConnectorArn)
	}

	if resp, err := client.DeleteConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a directory registration. Deleting a directory registration
// deauthorizes Amazon Web Services Private CA with the directory.
func pcaconnectorad_DeleteDirectoryRegistration(cfg aws.Config, client *pcaconnectorad.Client) {
	input := &pcaconnectorad.DeleteDirectoryRegistrationInput{
		// DirectoryRegistrationArn: *string, // Required
	}

	if len(_pcaconnectoradDirectoryRegistrationArn) > 0 {
		input.DirectoryRegistrationArn = aws.String(_pcaconnectoradDirectoryRegistrationArn)
	}

	if resp, err := client.DeleteDirectoryRegistration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the service principal name (SPN) used by a connector to authenticate
// with your Active Directory.
func pcaconnectorad_DeleteServicePrincipalName(cfg aws.Config, client *pcaconnectorad.Client) {
	input := &pcaconnectorad.DeleteServicePrincipalNameInput{
		// ConnectorArn: *string, // Required
		// DirectoryRegistrationArn: *string, // Required
	}

	if len(_pcaconnectoradConnectorArn) > 0 {
		input.ConnectorArn = aws.String(_pcaconnectoradConnectorArn)
	}
	if len(_pcaconnectoradDirectoryRegistrationArn) > 0 {
		input.DirectoryRegistrationArn = aws.String(_pcaconnectoradDirectoryRegistrationArn)
	}

	if resp, err := client.DeleteServicePrincipalName(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a template. Certificates issued using the template are still valid
// until they are revoked or expired.
func pcaconnectorad_DeleteTemplate(cfg aws.Config, client *pcaconnectorad.Client) {
	input := &pcaconnectorad.DeleteTemplateInput{
		// TemplateArn: *string, // Required
	}

	if len(_pcaconnectoradTemplateArn) > 0 {
		input.TemplateArn = aws.String(_pcaconnectoradTemplateArn)
	}

	if resp, err := client.DeleteTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a group access control entry.
func pcaconnectorad_DeleteTemplateGroupAccessControlEntry(cfg aws.Config, client *pcaconnectorad.Client) {
	input := &pcaconnectorad.DeleteTemplateGroupAccessControlEntryInput{
		// GroupSecurityIdentifier: *string, // Required
		// TemplateArn: *string, // Required
	}

	if len(_pcaconnectoradGroupSecurityIdentifier) > 0 {
		input.GroupSecurityIdentifier = aws.String(_pcaconnectoradGroupSecurityIdentifier)
	}
	if len(_pcaconnectoradTemplateArn) > 0 {
		input.TemplateArn = aws.String(_pcaconnectoradTemplateArn)
	}

	if resp, err := client.DeleteTemplateGroupAccessControlEntry(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists information about your connector. You specify the connector on input by
// its ARN (Amazon Resource Name).
func pcaconnectorad_GetConnector(cfg aws.Config, client *pcaconnectorad.Client) {
	input := &pcaconnectorad.GetConnectorInput{
		// ConnectorArn: *string, // Required
	}

	if len(_pcaconnectoradConnectorArn) > 0 {
		input.ConnectorArn = aws.String(_pcaconnectoradConnectorArn)
	}

	if resp, err := client.GetConnector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A structure that contains information about your directory registration.
func pcaconnectorad_GetDirectoryRegistration(cfg aws.Config, client *pcaconnectorad.Client) {
	input := &pcaconnectorad.GetDirectoryRegistrationInput{
		// DirectoryRegistrationArn: *string, // Required
	}

	if len(_pcaconnectoradDirectoryRegistrationArn) > 0 {
		input.DirectoryRegistrationArn = aws.String(_pcaconnectoradDirectoryRegistrationArn)
	}

	if resp, err := client.GetDirectoryRegistration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the service principal name that the connector uses to authenticate with
// Active Directory.
func pcaconnectorad_GetServicePrincipalName(cfg aws.Config, client *pcaconnectorad.Client) {
	input := &pcaconnectorad.GetServicePrincipalNameInput{
		// ConnectorArn: *string, // Required
		// DirectoryRegistrationArn: *string, // Required
	}

	if len(_pcaconnectoradConnectorArn) > 0 {
		input.ConnectorArn = aws.String(_pcaconnectoradConnectorArn)
	}
	if len(_pcaconnectoradDirectoryRegistrationArn) > 0 {
		input.DirectoryRegistrationArn = aws.String(_pcaconnectoradDirectoryRegistrationArn)
	}

	if resp, err := client.GetServicePrincipalName(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a certificate template that the connector uses to issue certificates
// from a private CA.
func pcaconnectorad_GetTemplate(cfg aws.Config, client *pcaconnectorad.Client) {
	input := &pcaconnectorad.GetTemplateInput{
		// TemplateArn: *string, // Required
	}

	if len(_pcaconnectoradTemplateArn) > 0 {
		input.TemplateArn = aws.String(_pcaconnectoradTemplateArn)
	}

	if resp, err := client.GetTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the group access control entries for a template.
func pcaconnectorad_GetTemplateGroupAccessControlEntry(cfg aws.Config, client *pcaconnectorad.Client) {
	input := &pcaconnectorad.GetTemplateGroupAccessControlEntryInput{
		// GroupSecurityIdentifier: *string, // Required
		// TemplateArn: *string, // Required
	}

	if len(_pcaconnectoradGroupSecurityIdentifier) > 0 {
		input.GroupSecurityIdentifier = aws.String(_pcaconnectoradGroupSecurityIdentifier)
	}
	if len(_pcaconnectoradTemplateArn) > 0 {
		input.TemplateArn = aws.String(_pcaconnectoradTemplateArn)
	}

	if resp, err := client.GetTemplateGroupAccessControlEntry(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the connectors that you created by using the [https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateConnector] action.
//
// [https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateConnector]: https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateConnector
func pcaconnectorad_ListConnectors(cfg aws.Config, client *pcaconnectorad.Client) {
	input := &pcaconnectorad.ListConnectorsInput{}

	if len(_pcaconnectoradMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pcaconnectoradMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pcaconnectoradNextToken) > 0 {
		input.NextToken = aws.String(_pcaconnectoradNextToken)
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

	var results []*pcaconnectorad.ListConnectorsOutput
	p := pcaconnectorad.NewListConnectorsPaginator(client, input)
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

// Lists the directory registrations that you created by using the [https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateDirectoryRegistration] action.
//
// [https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateDirectoryRegistration]: https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateDirectoryRegistration
func pcaconnectorad_ListDirectoryRegistrations(cfg aws.Config, client *pcaconnectorad.Client) {
	input := &pcaconnectorad.ListDirectoryRegistrationsInput{}

	if len(_pcaconnectoradMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pcaconnectoradMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pcaconnectoradNextToken) > 0 {
		input.NextToken = aws.String(_pcaconnectoradNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDirectoryRegistrations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pcaconnectorad.ListDirectoryRegistrationsOutput
	p := pcaconnectorad.NewListDirectoryRegistrationsPaginator(client, input)
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

// Lists the service principal names that the connector uses to authenticate with
// Active Directory.
func pcaconnectorad_ListServicePrincipalNames(cfg aws.Config, client *pcaconnectorad.Client) {
	input := &pcaconnectorad.ListServicePrincipalNamesInput{
		// DirectoryRegistrationArn: *string, // Required
	}

	if len(_pcaconnectoradDirectoryRegistrationArn) > 0 {
		input.DirectoryRegistrationArn = aws.String(_pcaconnectoradDirectoryRegistrationArn)
	}
	if len(_pcaconnectoradMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pcaconnectoradMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pcaconnectoradNextToken) > 0 {
		input.NextToken = aws.String(_pcaconnectoradNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListServicePrincipalNames(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pcaconnectorad.ListServicePrincipalNamesOutput
	p := pcaconnectorad.NewListServicePrincipalNamesPaginator(client, input)
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

// Lists the tags, if any, that are associated with your resource.
func pcaconnectorad_ListTagsForResource(cfg aws.Config, client *pcaconnectorad.Client) {
	input := &pcaconnectorad.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_pcaconnectoradResourceArn) > 0 {
		input.ResourceArn = aws.String(_pcaconnectoradResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists group access control entries you created.
func pcaconnectorad_ListTemplateGroupAccessControlEntries(cfg aws.Config, client *pcaconnectorad.Client) {
	input := &pcaconnectorad.ListTemplateGroupAccessControlEntriesInput{
		// TemplateArn: *string, // Required
	}

	if len(_pcaconnectoradTemplateArn) > 0 {
		input.TemplateArn = aws.String(_pcaconnectoradTemplateArn)
	}
	if len(_pcaconnectoradMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pcaconnectoradMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pcaconnectoradNextToken) > 0 {
		input.NextToken = aws.String(_pcaconnectoradNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTemplateGroupAccessControlEntries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pcaconnectorad.ListTemplateGroupAccessControlEntriesOutput
	p := pcaconnectorad.NewListTemplateGroupAccessControlEntriesPaginator(client, input)
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

// Lists the templates, if any, that are associated with a connector.
func pcaconnectorad_ListTemplates(cfg aws.Config, client *pcaconnectorad.Client) {
	input := &pcaconnectorad.ListTemplatesInput{
		// ConnectorArn: *string, // Required
	}

	if len(_pcaconnectoradConnectorArn) > 0 {
		input.ConnectorArn = aws.String(_pcaconnectoradConnectorArn)
	}
	if len(_pcaconnectoradMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pcaconnectoradMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pcaconnectoradNextToken) > 0 {
		input.NextToken = aws.String(_pcaconnectoradNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pcaconnectorad.ListTemplatesOutput
	p := pcaconnectorad.NewListTemplatesPaginator(client, input)
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

// Adds one or more tags to your resource.
func pcaconnectorad_TagResource(cfg aws.Config, client *pcaconnectorad.Client) {
	input := &pcaconnectorad.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_pcaconnectoradResourceArn) > 0 {
		input.ResourceArn = aws.String(_pcaconnectoradResourceArn)
	}
	if len(_pcaconnectoradTags) > 0 {
		if err := assignInputField(input, "Tags", _pcaconnectoradTags); err != nil {
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
func pcaconnectorad_UntagResource(cfg aws.Config, client *pcaconnectorad.Client) {
	input := &pcaconnectorad.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_pcaconnectoradResourceArn) > 0 {
		input.ResourceArn = aws.String(_pcaconnectoradResourceArn)
	}
	if len(_pcaconnectoradTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _pcaconnectoradTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update template configuration to define the information included in
// certificates.
func pcaconnectorad_UpdateTemplate(cfg aws.Config, client *pcaconnectorad.Client) {
	input := &pcaconnectorad.UpdateTemplateInput{
		// TemplateArn: *string, // Required
	}

	if len(_pcaconnectoradTemplateArn) > 0 {
		input.TemplateArn = aws.String(_pcaconnectoradTemplateArn)
	}
	if len(_pcaconnectoradDefinition) > 0 {
		if err := assignInputField(input, "Definition", _pcaconnectoradDefinition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}
	if len(_pcaconnectoradReenrollAllCertificateHolders) > 0 {
		if err := assignInputField(input, "ReenrollAllCertificateHolders", _pcaconnectoradReenrollAllCertificateHolders); err != nil {
			log.Errorf("invalid --reenroll-all-certificate-holders: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a group access control entry you created using [CreateTemplateGroupAccessControlEntry].
//
// [CreateTemplateGroupAccessControlEntry]: https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateTemplateGroupAccessControlEntry.html
func pcaconnectorad_UpdateTemplateGroupAccessControlEntry(cfg aws.Config, client *pcaconnectorad.Client) {
	input := &pcaconnectorad.UpdateTemplateGroupAccessControlEntryInput{
		// GroupSecurityIdentifier: *string, // Required
		// TemplateArn: *string, // Required
	}

	if len(_pcaconnectoradGroupSecurityIdentifier) > 0 {
		input.GroupSecurityIdentifier = aws.String(_pcaconnectoradGroupSecurityIdentifier)
	}
	if len(_pcaconnectoradTemplateArn) > 0 {
		input.TemplateArn = aws.String(_pcaconnectoradTemplateArn)
	}
	if len(_pcaconnectoradAccessRights) > 0 {
		if err := assignInputField(input, "AccessRights", _pcaconnectoradAccessRights); err != nil {
			log.Errorf("invalid --access-rights: %s", err.Error())
			return
		}
	}
	if len(_pcaconnectoradGroupDisplayName) > 0 {
		input.GroupDisplayName = aws.String(_pcaconnectoradGroupDisplayName)
	}

	if resp, err := client.UpdateTemplateGroupAccessControlEntry(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_pcaconnectoradCmd)
	_pcaconnectoradCmd.Flags().SortFlags = false

	_pcaconnectoradCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_pcaconnectoradCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_pcaconnectoradCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_pcaconnectoradCmd.Flags().StringVarP(&_pcaconnectoradAccessRights, "access-rights", "", "", "Access Rights")
	_pcaconnectoradCmd.Flags().StringVarP(&_pcaconnectoradCertificateAuthorityArn, "certificate-authority-arn", "", "", "Certificate Authority ARN")
	_pcaconnectoradCmd.Flags().StringVarP(&_pcaconnectoradClientToken, "client-token", "", "", "Client Token")
	_pcaconnectoradCmd.Flags().StringVarP(&_pcaconnectoradConnectorArn, "connector-arn", "", "", "Connector ARN")
	_pcaconnectoradCmd.Flags().StringVarP(&_pcaconnectoradDefinition, "definition", "", "", "Definition")
	_pcaconnectoradCmd.Flags().StringVarP(&_pcaconnectoradDirectoryId, "directory-id", "", "", "Directory ID")
	_pcaconnectoradCmd.Flags().StringVarP(&_pcaconnectoradDirectoryRegistrationArn, "directory-registration-arn", "", "", "Directory Registration ARN")
	_pcaconnectoradCmd.Flags().StringVarP(&_pcaconnectoradGroupDisplayName, "group-display-name", "", "", "Group Display Name")
	_pcaconnectoradCmd.Flags().StringVarP(&_pcaconnectoradGroupSecurityIdentifier, "group-security-identifier", "", "", "Group Security Identifier")
	_pcaconnectoradCmd.Flags().StringVarP(&_pcaconnectoradMaxResults, "max-results", "", "", "Max Results")
	_pcaconnectoradCmd.Flags().StringVarP(&_pcaconnectoradName, "name", "", "", "Name")
	_pcaconnectoradCmd.Flags().StringVarP(&_pcaconnectoradNextToken, "next-token", "", "", "Next Token")
	_pcaconnectoradCmd.Flags().StringVarP(&_pcaconnectoradReenrollAllCertificateHolders, "reenroll-all-certificate-holders", "", "", "Reenroll All Certificate Holders")
	_pcaconnectoradCmd.Flags().StringVarP(&_pcaconnectoradResourceArn, "resource-arn", "", "", "Resource ARN")
	_pcaconnectoradCmd.Flags().StringSliceVarP(&_pcaconnectoradTagKeys, "tag-keys", "", nil, "Tag Keys")
	_pcaconnectoradCmd.Flags().StringVarP(&_pcaconnectoradTags, "tags", "", "", "Tags")
	_pcaconnectoradCmd.Flags().StringVarP(&_pcaconnectoradTemplateArn, "template-arn", "", "", "Template ARN")
	_pcaconnectoradCmd.Flags().StringVarP(&_pcaconnectoradVpcInformation, "vpc-information", "", "", "VPC Information")

	_pcaconnectoradCmd.Flags().BoolVarP(&_pcaconnectoradCreateConnector, "create-connector", "", false, "Create Connector")
	_pcaconnectoradCmd.Flags().BoolVarP(&_pcaconnectoradCreateDirectoryRegistration, "create-directory-registration", "", false, "Create Directory Registration")
	_pcaconnectoradCmd.Flags().BoolVarP(&_pcaconnectoradCreateServicePrincipalName, "create-service-principal-name", "", false, "Create Service Principal Name")
	_pcaconnectoradCmd.Flags().BoolVarP(&_pcaconnectoradCreateTemplate, "create-template", "", false, "Create Template")
	_pcaconnectoradCmd.Flags().BoolVarP(&_pcaconnectoradCreateTemplateGroupAccessControlEntry, "create-template-group-access-control-entry", "", false, "Create Template Group Access Control Entry")
	_pcaconnectoradCmd.Flags().BoolVarP(&_pcaconnectoradDeleteConnector, "delete-connector", "", false, "Delete Connector")
	_pcaconnectoradCmd.Flags().BoolVarP(&_pcaconnectoradDeleteDirectoryRegistration, "delete-directory-registration", "", false, "Delete Directory Registration")
	_pcaconnectoradCmd.Flags().BoolVarP(&_pcaconnectoradDeleteServicePrincipalName, "delete-service-principal-name", "", false, "Delete Service Principal Name")
	_pcaconnectoradCmd.Flags().BoolVarP(&_pcaconnectoradDeleteTemplate, "delete-template", "", false, "Delete Template")
	_pcaconnectoradCmd.Flags().BoolVarP(&_pcaconnectoradDeleteTemplateGroupAccessControlEntry, "delete-template-group-access-control-entry", "", false, "Delete Template Group Access Control Entry")
	_pcaconnectoradCmd.Flags().BoolVarP(&_pcaconnectoradGetConnector, "get-connector", "", false, "Get Connector")
	_pcaconnectoradCmd.Flags().BoolVarP(&_pcaconnectoradGetDirectoryRegistration, "get-directory-registration", "", false, "Get Directory Registration")
	_pcaconnectoradCmd.Flags().BoolVarP(&_pcaconnectoradGetServicePrincipalName, "get-service-principal-name", "", false, "Get Service Principal Name")
	_pcaconnectoradCmd.Flags().BoolVarP(&_pcaconnectoradGetTemplate, "get-template", "", false, "Get Template")
	_pcaconnectoradCmd.Flags().BoolVarP(&_pcaconnectoradGetTemplateGroupAccessControlEntry, "get-template-group-access-control-entry", "", false, "Get Template Group Access Control Entry")
	_pcaconnectoradCmd.Flags().BoolVarP(&_pcaconnectoradListConnectors, "list-connectors", "", false, "List Connectors")
	_pcaconnectoradCmd.Flags().BoolVarP(&_pcaconnectoradListDirectoryRegistrations, "list-directory-registrations", "", false, "List Directory Registrations")
	_pcaconnectoradCmd.Flags().BoolVarP(&_pcaconnectoradListServicePrincipalNames, "list-service-principal-names", "", false, "List Service Principal Names")
	_pcaconnectoradCmd.Flags().BoolVarP(&_pcaconnectoradListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_pcaconnectoradCmd.Flags().BoolVarP(&_pcaconnectoradListTemplateGroupAccessControlEntries, "list-template-group-access-control-entries", "", false, "List Template Group Access Control Entries")
	_pcaconnectoradCmd.Flags().BoolVarP(&_pcaconnectoradListTemplates, "list-templates", "", false, "List Templates")
	_pcaconnectoradCmd.Flags().BoolVarP(&_pcaconnectoradTagResource, "tag-resource", "", false, "Tag Resource")
	_pcaconnectoradCmd.Flags().BoolVarP(&_pcaconnectoradUntagResource, "untag-resource", "", false, "Untag Resource")
	_pcaconnectoradCmd.Flags().BoolVarP(&_pcaconnectoradUpdateTemplate, "update-template", "", false, "Update Template")
	_pcaconnectoradCmd.Flags().BoolVarP(&_pcaconnectoradUpdateTemplateGroupAccessControlEntry, "update-template-group-access-control-entry", "", false, "Update Template Group Access Control Entry")

}
