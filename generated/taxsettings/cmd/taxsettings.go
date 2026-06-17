package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/taxsettings"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// taxsettingsCmd represents the taxsettings command
var _taxsettingsCmd = &cobra.Command{
	Use:   "taxsettings",
	Short: "AWS taxsettings CLI",
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
		client := taxsettings.NewFromConfig(cfg)
		if _taxsettingsBatchDeleteTaxRegistration {
			taxsettings_BatchDeleteTaxRegistration(cfg, client)
			return
		}
		if _taxsettingsBatchGetTaxExemptions {
			taxsettings_BatchGetTaxExemptions(cfg, client)
			return
		}
		if _taxsettingsBatchPutTaxRegistration {
			taxsettings_BatchPutTaxRegistration(cfg, client)
			return
		}
		if _taxsettingsDeleteSupplementalTaxRegistration {
			taxsettings_DeleteSupplementalTaxRegistration(cfg, client)
			return
		}
		if _taxsettingsDeleteTaxRegistration {
			taxsettings_DeleteTaxRegistration(cfg, client)
			return
		}
		if _taxsettingsGetTaxExemptionTypes {
			taxsettings_GetTaxExemptionTypes(cfg, client)
			return
		}
		if _taxsettingsGetTaxInheritance {
			taxsettings_GetTaxInheritance(cfg, client)
			return
		}
		if _taxsettingsGetTaxRegistration {
			taxsettings_GetTaxRegistration(cfg, client)
			return
		}
		if _taxsettingsGetTaxRegistrationDocument {
			taxsettings_GetTaxRegistrationDocument(cfg, client)
			return
		}
		if _taxsettingsListSupplementalTaxRegistrations {
			taxsettings_ListSupplementalTaxRegistrations(cfg, client)
			return
		}
		if _taxsettingsListTaxExemptions {
			taxsettings_ListTaxExemptions(cfg, client)
			return
		}
		if _taxsettingsListTaxRegistrations {
			taxsettings_ListTaxRegistrations(cfg, client)
			return
		}
		if _taxsettingsPutSupplementalTaxRegistration {
			taxsettings_PutSupplementalTaxRegistration(cfg, client)
			return
		}
		if _taxsettingsPutTaxExemption {
			taxsettings_PutTaxExemption(cfg, client)
			return
		}
		if _taxsettingsPutTaxInheritance {
			taxsettings_PutTaxInheritance(cfg, client)
			return
		}
		if _taxsettingsPutTaxRegistration {
			taxsettings_PutTaxRegistration(cfg, client)
			return
		}

	},
}

var (
	_taxsettingsBatchDeleteTaxRegistration        bool
	_taxsettingsBatchGetTaxExemptions             bool
	_taxsettingsBatchPutTaxRegistration           bool
	_taxsettingsDeleteSupplementalTaxRegistration bool
	_taxsettingsDeleteTaxRegistration             bool
	_taxsettingsGetTaxExemptionTypes              bool
	_taxsettingsGetTaxInheritance                 bool
	_taxsettingsGetTaxRegistration                bool
	_taxsettingsGetTaxRegistrationDocument        bool
	_taxsettingsListSupplementalTaxRegistrations  bool
	_taxsettingsListTaxExemptions                 bool
	_taxsettingsListTaxRegistrations              bool
	_taxsettingsPutSupplementalTaxRegistration    bool
	_taxsettingsPutTaxExemption                   bool
	_taxsettingsPutTaxInheritance                 bool
	_taxsettingsPutTaxRegistration                bool

	_taxsettingsAccountId             string
	_taxsettingsAccountIds            []string
	_taxsettingsAuthority             string
	_taxsettingsAuthorityId           string
	_taxsettingsDestinationS3Location string
	_taxsettingsExemptionCertificate  string
	_taxsettingsExemptionType         string
	_taxsettingsHeritageStatus        string
	_taxsettingsMaxResults            string
	_taxsettingsNextToken             string
	_taxsettingsTaxDocumentMetadata   string
	_taxsettingsTaxRegistrationEntry  string
)

// Deletes tax registration for multiple accounts in batch. This can be used to
// delete tax registrations for up to five accounts in one batch.
//
// This API operation can't be used to delete your tax registration in Brazil. Use
// the [Payment preferences]page in the Billing and Cost Management console instead.
//
// [Payment preferences]: https://console.aws.amazon.com/billing/home#/paymentpreferences/paymentmethods
func taxsettings_BatchDeleteTaxRegistration(cfg aws.Config, client *taxsettings.Client) {
	input := &taxsettings.BatchDeleteTaxRegistrationInput{
		// AccountIds: []string, // Required
	}

	if len(_taxsettingsAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _taxsettingsAccountIds...)
	}

	if resp, err := client.BatchDeleteTaxRegistration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the active tax exemptions for a given list of accounts. The IAM action is
// tax:GetExemptions .
func taxsettings_BatchGetTaxExemptions(cfg aws.Config, client *taxsettings.Client) {
	input := &taxsettings.BatchGetTaxExemptionsInput{
		// AccountIds: []string, // Required
	}

	if len(_taxsettingsAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _taxsettingsAccountIds...)
	}

	if resp, err := client.BatchGetTaxExemptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates tax registration for multiple accounts in batch. This can be
// used to add or update tax registrations for up to five accounts in one batch.
// You can't set a TRN if there's a pending TRN. You'll need to delete the pending
// TRN first.
//
// To call this API operation for specific countries, see the following
// country-specific requirements.
//
// # Bangladesh
//
// - You must specify the tax registration certificate document in the
// taxRegistrationDocuments field of the VerificationDetails object.
//
// # Brazil
//
// - You must complete the tax registration process in the [Payment preferences]page in the Billing
// and Cost Management console. After your TRN and billing address are verified,
// you can call this API operation.
//
// - For Amazon Web Services accounts created through Organizations, you can
// call this API operation when you don't have a billing address.
//
// # Georgia
//
// - The valid personType values are Physical Person and Business .
//
// # Indonesia
//
// - PutTaxRegistration : The use of this operation to submit tax information is
// subject to the [Amazon Web Services service terms]. By submitting, you’re providing consent for Amazon Web
// Services to validate NIK, NPWP, and NITKU data, provided by you with the
// Directorate General of Taxes of Indonesia in accordance with the Minister of
// Finance Regulation (PMK) Number 112/PMK.03/2022.
//
// - BatchPutTaxRegistration : The use of this operation to submit tax
// information is subject to the [Amazon Web Services service terms]. By submitting, you’re providing consent for
// Amazon Web Services to validate NIK, NPWP, and NITKU data, provided by you with
// the Directorate General of Taxes of Indonesia in accordance with the Minister of
// Finance Regulation (PMK) Number 112/PMK.03/2022, through our third-party partner
// PT Achilles Advanced Management (OnlinePajak).
//
// - You must specify the taxRegistrationNumberType in the
// indonesiaAdditionalInfo field of the additionalTaxInformation object.
//
// - If you specify decisionNumber , you must specify the
// ppnExceptionDesignationCode in the indonesiaAdditionalInfo field of the
// additionalTaxInformation object. If the taxRegistrationNumberType is set to
// NPWP or NITKU, valid values for ppnExceptionDesignationCode are either 01 , 02
// , 03 , 07 , or 08 .
//
// # For other taxRegistrationNumberType values, ppnExceptionDesignationCode must be
//
// either 01 , 07 , or 08 .
//
// - If ppnExceptionDesignationCode is 07 , you must specify the decisionNumber
// in the indonesiaAdditionalInfo field of the additionalTaxInformation object.
//
// # Kenya
//
// - You must specify the personType in the kenyaAdditionalInfo field of the
// additionalTaxInformation object.
//
// - If the personType is Physical Person , you must specify the tax registration
// certificate document in the taxRegistrationDocuments field of the
// VerificationDetails object.
//
// # Malaysia
//
// - The sector valid values are Business and Individual .
//
// - RegistrationType valid values are NRIC for individual, and TIN and sales and
// service tax (SST) for Business.
//
// - For individual, you can specify the taxInformationNumber in
// MalaysiaAdditionalInfo with NRIC type, and a valid MyKad or NRIC number.
//
// - For business, you must specify a businessRegistrationNumber in
// MalaysiaAdditionalInfo with a TIN type and tax identification number.
//
// - For business resellers, you must specify a businessRegistrationNumber and
// taxInformationNumber in MalaysiaAdditionalInfo with a sales and service tax
// (SST) type and a valid SST number.
//
// - For business resellers with service codes, you must specify
// businessRegistrationNumber , taxInformationNumber , and distinct
// serviceTaxCodes in MalaysiaAdditionalInfo with a SST type and valid sales and
// service tax (SST) number. By using this API operation, Amazon Web Services
// registers your self-declaration that you’re an authorized business reseller
// registered with the Royal Malaysia Customs Department (RMCD), and have a valid
// SST number.
//
// - Amazon Web Services reserves the right to seek additional information
// and/or take other actions to support your self-declaration as appropriate.
//
// - Amazon Web Services is currently registered under the following service tax
// codes. You must include at least one of the service tax codes in the service tax
// code strings to declare yourself as an authorized registered business reseller.
//
// Taxable service and service tax codes:
//
// # Consultancy - 9907061674
//
// # Training or coaching service - 9907071685
//
// # IT service - 9907101676
//
// # Digital services and electronic medium - 9907121690
//
// # Nepal
//
// - The sector valid values are Business and Individual .
//
// # Saudi Arabia
//
// - For address , you must specify addressLine3 .
//
// # South Korea
//
// - You must specify the certifiedEmailId and legalName in the
// TaxRegistrationEntry object. Use Korean characters for legalName .
//
// - You must specify the businessRepresentativeName , itemOfBusiness , and
// lineOfBusiness in the southKoreaAdditionalInfo field of the
// additionalTaxInformation object. Use Korean characters for these fields.
//
// - You must specify the tax registration certificate document in the
// taxRegistrationDocuments field of the VerificationDetails object.
//
// - For the address object, use Korean characters for addressLine1 ,
// addressLine2 city , postalCode , and stateOrRegion .
//
// # Spain
//
// - You must specify the registrationType in the spainAdditionalInfo field of
// the additionalTaxInformation object.
//
// - If the registrationType is Local , you must specify the tax registration
// certificate document in the taxRegistrationDocuments field of the
// VerificationDetails object.
//
// # Turkey
//
// - You must specify the sector in the taxRegistrationEntry object.
//
// - If your sector is Business , Individual , or Government :
//
// - Specify the taxOffice . If your sector is Individual , don't enter this
// value.
//
// - (Optional) Specify the kepEmailId . If your sector is Individual , don't
// enter this value.
//
// - Note: In the Tax Settings page of the Billing console, Government appears as
// Public institutions
//
// - If your sector is Business and you're subject to KDV tax, you must specify
// your industry in the industries field.
//
// - For address , you must specify districtOrCounty .
//
// # Ukraine
//
// - The sector valid values are Business and Individual .
//
// [Amazon Web Services service terms]: http://aws.amazon.com/service-terms/
// [Payment preferences]: https://console.aws.amazon.com/billing/home#/paymentpreferences/paymentmethods
func taxsettings_BatchPutTaxRegistration(cfg aws.Config, client *taxsettings.Client) {
	input := &taxsettings.BatchPutTaxRegistrationInput{
		// AccountIds: []string, // Required
		// TaxRegistrationEntry: *types.TaxRegistrationEntry, // Required
	}

	if len(_taxsettingsAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _taxsettingsAccountIds...)
	}
	if len(_taxsettingsTaxRegistrationEntry) > 0 {
		if err := assignInputField(input, "TaxRegistrationEntry", _taxsettingsTaxRegistrationEntry); err != nil {
			log.Errorf("invalid --tax-registration-entry: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchPutTaxRegistration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a supplemental tax registration for a single account.
func taxsettings_DeleteSupplementalTaxRegistration(cfg aws.Config, client *taxsettings.Client) {
	input := &taxsettings.DeleteSupplementalTaxRegistrationInput{
		// AuthorityId: *string, // Required
	}

	if len(_taxsettingsAuthorityId) > 0 {
		input.AuthorityId = aws.String(_taxsettingsAuthorityId)
	}

	if resp, err := client.DeleteSupplementalTaxRegistration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes tax registration for a single account.
// This API operation can't be used to delete your tax registration in Brazil. Use
// the [Payment preferences]page in the Billing and Cost Management console instead.
//
// [Payment preferences]: https://console.aws.amazon.com/billing/home#/paymentpreferences/paymentmethods
func taxsettings_DeleteTaxRegistration(cfg aws.Config, client *taxsettings.Client) {
	input := &taxsettings.DeleteTaxRegistrationInput{}

	if len(_taxsettingsAccountId) > 0 {
		input.AccountId = aws.String(_taxsettingsAccountId)
	}

	if resp, err := client.DeleteTaxRegistration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get supported tax exemption types. The IAM action is tax:GetExemptions .
func taxsettings_GetTaxExemptionTypes(cfg aws.Config, client *taxsettings.Client) {
	input := &taxsettings.GetTaxExemptionTypesInput{}

	if resp, err := client.GetTaxExemptionTypes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The get account tax inheritance status.
func taxsettings_GetTaxInheritance(cfg aws.Config, client *taxsettings.Client) {
	input := &taxsettings.GetTaxInheritanceInput{}

	if resp, err := client.GetTaxInheritance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves tax registration for a single account.
func taxsettings_GetTaxRegistration(cfg aws.Config, client *taxsettings.Client) {
	input := &taxsettings.GetTaxRegistrationInput{}

	if len(_taxsettingsAccountId) > 0 {
		input.AccountId = aws.String(_taxsettingsAccountId)
	}

	if resp, err := client.GetTaxRegistration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Downloads your tax documents to the Amazon S3 bucket that you specify in your
// request.
func taxsettings_GetTaxRegistrationDocument(cfg aws.Config, client *taxsettings.Client) {
	input := &taxsettings.GetTaxRegistrationDocumentInput{
		// TaxDocumentMetadata: *types.TaxDocumentMetadata, // Required
	}

	if len(_taxsettingsTaxDocumentMetadata) > 0 {
		if err := assignInputField(input, "TaxDocumentMetadata", _taxsettingsTaxDocumentMetadata); err != nil {
			log.Errorf("invalid --tax-document-metadata: %s", err.Error())
			return
		}
	}
	if len(_taxsettingsDestinationS3Location) > 0 {
		if err := assignInputField(input, "DestinationS3Location", _taxsettingsDestinationS3Location); err != nil {
			log.Errorf("invalid --destination-s3-location: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetTaxRegistrationDocument(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves supplemental tax registrations for a single account.
func taxsettings_ListSupplementalTaxRegistrations(cfg aws.Config, client *taxsettings.Client) {
	input := &taxsettings.ListSupplementalTaxRegistrationsInput{}

	if len(_taxsettingsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _taxsettingsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_taxsettingsNextToken) > 0 {
		input.NextToken = aws.String(_taxsettingsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSupplementalTaxRegistrations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*taxsettings.ListSupplementalTaxRegistrationsOutput
	p := taxsettings.NewListSupplementalTaxRegistrationsPaginator(client, input)
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

// Retrieves the tax exemption of accounts listed in a consolidated billing
// family. The IAM action is tax:GetExemptions .
func taxsettings_ListTaxExemptions(cfg aws.Config, client *taxsettings.Client) {
	input := &taxsettings.ListTaxExemptionsInput{}

	if len(_taxsettingsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _taxsettingsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_taxsettingsNextToken) > 0 {
		input.NextToken = aws.String(_taxsettingsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTaxExemptions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*taxsettings.ListTaxExemptionsOutput
	p := taxsettings.NewListTaxExemptionsPaginator(client, input)
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

// Retrieves the tax registration of accounts listed in a consolidated billing
// family. This can be used to retrieve up to 100 accounts' tax registrations in
// one call (default 50).
func taxsettings_ListTaxRegistrations(cfg aws.Config, client *taxsettings.Client) {
	input := &taxsettings.ListTaxRegistrationsInput{}

	if len(_taxsettingsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _taxsettingsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_taxsettingsNextToken) > 0 {
		input.NextToken = aws.String(_taxsettingsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTaxRegistrations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*taxsettings.ListTaxRegistrationsOutput
	p := taxsettings.NewListTaxRegistrationsPaginator(client, input)
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

// Stores supplemental tax registration for a single account.
func taxsettings_PutSupplementalTaxRegistration(cfg aws.Config, client *taxsettings.Client) {
	input := &taxsettings.PutSupplementalTaxRegistrationInput{
		// TaxRegistrationEntry: *types.SupplementalTaxRegistrationEntry, // Required
	}

	if len(_taxsettingsTaxRegistrationEntry) > 0 {
		if err := assignInputField(input, "TaxRegistrationEntry", _taxsettingsTaxRegistrationEntry); err != nil {
			log.Errorf("invalid --tax-registration-entry: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutSupplementalTaxRegistration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds the tax exemption for a single account or all accounts listed in a
// consolidated billing family. The IAM action is tax:UpdateExemptions .
func taxsettings_PutTaxExemption(cfg aws.Config, client *taxsettings.Client) {
	input := &taxsettings.PutTaxExemptionInput{
		// AccountIds: []string, // Required
		// Authority: *types.Authority, // Required
		// ExemptionCertificate: *types.ExemptionCertificate, // Required
		// ExemptionType: *string, // Required
	}

	if len(_taxsettingsAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _taxsettingsAccountIds...)
	}
	if len(_taxsettingsAuthority) > 0 {
		if err := assignInputField(input, "Authority", _taxsettingsAuthority); err != nil {
			log.Errorf("invalid --authority: %s", err.Error())
			return
		}
	}
	if len(_taxsettingsExemptionCertificate) > 0 {
		if err := assignInputField(input, "ExemptionCertificate", _taxsettingsExemptionCertificate); err != nil {
			log.Errorf("invalid --exemption-certificate: %s", err.Error())
			return
		}
	}
	if len(_taxsettingsExemptionType) > 0 {
		input.ExemptionType = aws.String(_taxsettingsExemptionType)
	}

	if resp, err := client.PutTaxExemption(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The updated tax inheritance status.
func taxsettings_PutTaxInheritance(cfg aws.Config, client *taxsettings.Client) {
	input := &taxsettings.PutTaxInheritanceInput{}

	if len(_taxsettingsHeritageStatus) > 0 {
		if err := assignInputField(input, "HeritageStatus", _taxsettingsHeritageStatus); err != nil {
			log.Errorf("invalid --heritage-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutTaxInheritance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates tax registration for a single account. You can't set a TRN if
// there's a pending TRN. You'll need to delete the pending TRN first.
//
// To call this API operation for specific countries, see the following
// country-specific requirements.
//
// # Bangladesh
//
// - You must specify the tax registration certificate document in the
// taxRegistrationDocuments field of the VerificationDetails object.
//
// # Brazil
//
// - You must complete the tax registration process in the [Payment preferences]page in the Billing
// and Cost Management console. After your TRN and billing address are verified,
// you can call this API operation.
//
// - For Amazon Web Services accounts created through Organizations, you can
// call this API operation when you don't have a billing address.
//
// # Georgia
//
// - The valid personType values are Physical Person and Business .
//
// # Indonesia
//
// - PutTaxRegistration : The use of this operation to submit tax information is
// subject to the [Amazon Web Services service terms]. By submitting, you’re providing consent for Amazon Web
// Services to validate NIK, NPWP, and NITKU data, provided by you with the
// Directorate General of Taxes of Indonesia in accordance with the Minister of
// Finance Regulation (PMK) Number 112/PMK.03/2022.
//
// - BatchPutTaxRegistration : The use of this operation to submit tax
// information is subject to the [Amazon Web Services service terms]. By submitting, you’re providing consent for
// Amazon Web Services to validate NIK, NPWP, and NITKU data, provided by you with
// the Directorate General of Taxes of Indonesia in accordance with the Minister of
// Finance Regulation (PMK) Number 112/PMK.03/2022, through our third-party partner
// PT Achilles Advanced Management (OnlinePajak).
//
// - You must specify the taxRegistrationNumberType in the
// indonesiaAdditionalInfo field of the additionalTaxInformation object.
//
// - If you specify decisionNumber , you must specify the
// ppnExceptionDesignationCode in the indonesiaAdditionalInfo field of the
// additionalTaxInformation object. If the taxRegistrationNumberType is set to
// NPWP or NITKU, valid values for ppnExceptionDesignationCode are either 01 , 02
// , 03 , 07 , or 08 .
//
// # For other taxRegistrationNumberType values, ppnExceptionDesignationCode must be
//
// either 01 , 07 , or 08 .
//
// - If ppnExceptionDesignationCode is 07 , you must specify the decisionNumber
// in the indonesiaAdditionalInfo field of the additionalTaxInformation object.
//
// # Kenya
//
// - You must specify the personType in the kenyaAdditionalInfo field of the
// additionalTaxInformation object.
//
// - If the personType is Physical Person , you must specify the tax registration
// certificate document in the taxRegistrationDocuments field of the
// VerificationDetails object.
//
// # Malaysia
//
// - The sector valid values are Business and Individual .
//
// - RegistrationType valid values are NRIC for individual, and TIN and sales and
// service tax (SST) for Business.
//
// - For individual, you can specify the taxInformationNumber in
// MalaysiaAdditionalInfo with NRIC type, and a valid MyKad or NRIC number.
//
// - For business, you must specify a businessRegistrationNumber in
// MalaysiaAdditionalInfo with a TIN type and tax identification number.
//
// - For business resellers, you must specify a businessRegistrationNumber and
// taxInformationNumber in MalaysiaAdditionalInfo with a sales and service tax
// (SST) type and a valid SST number.
//
// - For business resellers with service codes, you must specify
// businessRegistrationNumber , taxInformationNumber , and distinct
// serviceTaxCodes in MalaysiaAdditionalInfo with a SST type and valid sales and
// service tax (SST) number. By using this API operation, Amazon Web Services
// registers your self-declaration that you’re an authorized business reseller
// registered with the Royal Malaysia Customs Department (RMCD), and have a valid
// SST number.
//
// - Amazon Web Services reserves the right to seek additional information
// and/or take other actions to support your self-declaration as appropriate.
//
// - Amazon Web Services is currently registered under the following service tax
// codes. You must include at least one of the service tax codes in the service tax
// code strings to declare yourself as an authorized registered business reseller.
//
// Taxable service and service tax codes:
//
// # Consultancy - 9907061674
//
// # Training or coaching service - 9907071685
//
// # IT service - 9907101676
//
// # Digital services and electronic medium - 9907121690
//
// # Nepal
//
// - The sector valid values are Business and Individual .
//
// # Saudi Arabia
//
// - For address , you must specify addressLine3 .
//
// # South Korea
//
// - You must specify the certifiedEmailId and legalName in the
// TaxRegistrationEntry object. Use Korean characters for legalName .
//
// - You must specify the businessRepresentativeName , itemOfBusiness , and
// lineOfBusiness in the southKoreaAdditionalInfo field of the
// additionalTaxInformation object. Use Korean characters for these fields.
//
// - You must specify the tax registration certificate document in the
// taxRegistrationDocuments field of the VerificationDetails object.
//
// - For the address object, use Korean characters for addressLine1 ,
// addressLine2 city , postalCode , and stateOrRegion .
//
// # Spain
//
// - You must specify the registrationType in the spainAdditionalInfo field of
// the additionalTaxInformation object.
//
// - If the registrationType is Local , you must specify the tax registration
// certificate document in the taxRegistrationDocuments field of the
// VerificationDetails object.
//
// # Turkey
//
// - You must specify the sector in the taxRegistrationEntry object.
//
// - If your sector is Business , Individual , or Government :
//
// - Specify the taxOffice . If your sector is Individual , don't enter this
// value.
//
// - (Optional) Specify the kepEmailId . If your sector is Individual , don't
// enter this value.
//
// - Note: In the Tax Settings page of the Billing console, Government appears as
// Public institutions
//
// - If your sector is Business and you're subject to KDV tax, you must specify
// your industry in the industries field.
//
// - For address , you must specify districtOrCounty .
//
// # Ukraine
//
// - The sector valid values are Business and Individual .
//
// [Amazon Web Services service terms]: http://aws.amazon.com/service-terms/
// [Payment preferences]: https://console.aws.amazon.com/billing/home#/paymentpreferences/paymentmethods
func taxsettings_PutTaxRegistration(cfg aws.Config, client *taxsettings.Client) {
	input := &taxsettings.PutTaxRegistrationInput{
		// TaxRegistrationEntry: *types.TaxRegistrationEntry, // Required
	}

	if len(_taxsettingsTaxRegistrationEntry) > 0 {
		if err := assignInputField(input, "TaxRegistrationEntry", _taxsettingsTaxRegistrationEntry); err != nil {
			log.Errorf("invalid --tax-registration-entry: %s", err.Error())
			return
		}
	}
	if len(_taxsettingsAccountId) > 0 {
		input.AccountId = aws.String(_taxsettingsAccountId)
	}

	if resp, err := client.PutTaxRegistration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_taxsettingsCmd)
	_taxsettingsCmd.Flags().SortFlags = false

	_taxsettingsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_taxsettingsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_taxsettingsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_taxsettingsCmd.Flags().StringVarP(&_taxsettingsAccountId, "account-id", "", "", "Account ID")
	_taxsettingsCmd.Flags().StringSliceVarP(&_taxsettingsAccountIds, "account-ids", "", nil, "Account Ids")
	_taxsettingsCmd.Flags().StringVarP(&_taxsettingsAuthority, "authority", "", "", "Authority")
	_taxsettingsCmd.Flags().StringVarP(&_taxsettingsAuthorityId, "authority-id", "", "", "Authority ID")
	_taxsettingsCmd.Flags().StringVarP(&_taxsettingsDestinationS3Location, "destination-s3-location", "", "", "Destination S3 Location")
	_taxsettingsCmd.Flags().StringVarP(&_taxsettingsExemptionCertificate, "exemption-certificate", "", "", "Exemption Certificate")
	_taxsettingsCmd.Flags().StringVarP(&_taxsettingsExemptionType, "exemption-type", "", "", "Exemption Type")
	_taxsettingsCmd.Flags().StringVarP(&_taxsettingsHeritageStatus, "heritage-status", "", "", "Heritage Status")
	_taxsettingsCmd.Flags().StringVarP(&_taxsettingsMaxResults, "max-results", "", "", "Max Results")
	_taxsettingsCmd.Flags().StringVarP(&_taxsettingsNextToken, "next-token", "", "", "Next Token")
	_taxsettingsCmd.Flags().StringVarP(&_taxsettingsTaxDocumentMetadata, "tax-document-metadata", "", "", "Tax Document Metadata")
	_taxsettingsCmd.Flags().StringVarP(&_taxsettingsTaxRegistrationEntry, "tax-registration-entry", "", "", "Tax Registration Entry")

	_taxsettingsCmd.Flags().BoolVarP(&_taxsettingsBatchDeleteTaxRegistration, "batch-delete-tax-registration", "", false, "Batch Delete Tax Registration")
	_taxsettingsCmd.Flags().BoolVarP(&_taxsettingsBatchGetTaxExemptions, "batch-get-tax-exemptions", "", false, "Batch Get Tax Exemptions")
	_taxsettingsCmd.Flags().BoolVarP(&_taxsettingsBatchPutTaxRegistration, "batch-put-tax-registration", "", false, "Batch Put Tax Registration")
	_taxsettingsCmd.Flags().BoolVarP(&_taxsettingsDeleteSupplementalTaxRegistration, "delete-supplemental-tax-registration", "", false, "Delete Supplemental Tax Registration")
	_taxsettingsCmd.Flags().BoolVarP(&_taxsettingsDeleteTaxRegistration, "delete-tax-registration", "", false, "Delete Tax Registration")
	_taxsettingsCmd.Flags().BoolVarP(&_taxsettingsGetTaxExemptionTypes, "get-tax-exemption-types", "", false, "Get Tax Exemption Types")
	_taxsettingsCmd.Flags().BoolVarP(&_taxsettingsGetTaxInheritance, "get-tax-inheritance", "", false, "Get Tax Inheritance")
	_taxsettingsCmd.Flags().BoolVarP(&_taxsettingsGetTaxRegistration, "get-tax-registration", "", false, "Get Tax Registration")
	_taxsettingsCmd.Flags().BoolVarP(&_taxsettingsGetTaxRegistrationDocument, "get-tax-registration-document", "", false, "Get Tax Registration Document")
	_taxsettingsCmd.Flags().BoolVarP(&_taxsettingsListSupplementalTaxRegistrations, "list-supplemental-tax-registrations", "", false, "List Supplemental Tax Registrations")
	_taxsettingsCmd.Flags().BoolVarP(&_taxsettingsListTaxExemptions, "list-tax-exemptions", "", false, "List Tax Exemptions")
	_taxsettingsCmd.Flags().BoolVarP(&_taxsettingsListTaxRegistrations, "list-tax-registrations", "", false, "List Tax Registrations")
	_taxsettingsCmd.Flags().BoolVarP(&_taxsettingsPutSupplementalTaxRegistration, "put-supplemental-tax-registration", "", false, "Put Supplemental Tax Registration")
	_taxsettingsCmd.Flags().BoolVarP(&_taxsettingsPutTaxExemption, "put-tax-exemption", "", false, "Put Tax Exemption")
	_taxsettingsCmd.Flags().BoolVarP(&_taxsettingsPutTaxInheritance, "put-tax-inheritance", "", false, "Put Tax Inheritance")
	_taxsettingsCmd.Flags().BoolVarP(&_taxsettingsPutTaxRegistration, "put-tax-registration", "", false, "Put Tax Registration")

}
