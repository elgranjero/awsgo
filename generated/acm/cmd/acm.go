package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// acmCmd represents the acm command
var _acmCmd = &cobra.Command{
	Use:   "acm",
	Short: "AWS acm CLI",
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
		client := acm.NewFromConfig(cfg)
		if _acmAddTagsToCertificate {
			acm_AddTagsToCertificate(cfg, client)
			return
		}
		if _acmDeleteCertificate {
			acm_DeleteCertificate(cfg, client)
			return
		}
		if _acmDescribeCertificate {
			acm_DescribeCertificate(cfg, client)
			return
		}
		if _acmExportCertificate {
			acm_ExportCertificate(cfg, client)
			return
		}
		if _acmGetAccountConfiguration {
			acm_GetAccountConfiguration(cfg, client)
			return
		}
		if _acmGetCertificate {
			acm_GetCertificate(cfg, client)
			return
		}
		if _acmImportCertificate {
			acm_ImportCertificate(cfg, client)
			return
		}
		if _acmListCertificates {
			acm_ListCertificates(cfg, client)
			return
		}
		if _acmListTagsForCertificate {
			acm_ListTagsForCertificate(cfg, client)
			return
		}
		if _acmPutAccountConfiguration {
			acm_PutAccountConfiguration(cfg, client)
			return
		}
		if _acmRemoveTagsFromCertificate {
			acm_RemoveTagsFromCertificate(cfg, client)
			return
		}
		if _acmRenewCertificate {
			acm_RenewCertificate(cfg, client)
			return
		}
		if _acmRequestCertificate {
			acm_RequestCertificate(cfg, client)
			return
		}
		if _acmResendValidationEmail {
			acm_ResendValidationEmail(cfg, client)
			return
		}
		if _acmRevokeCertificate {
			acm_RevokeCertificate(cfg, client)
			return
		}
		if _acmUpdateCertificateOptions {
			acm_UpdateCertificateOptions(cfg, client)
			return
		}

	},
}

var (
	_acmAddTagsToCertificate      bool
	_acmDeleteCertificate         bool
	_acmDescribeCertificate       bool
	_acmExportCertificate         bool
	_acmGetAccountConfiguration   bool
	_acmGetCertificate            bool
	_acmImportCertificate         bool
	_acmListCertificates          bool
	_acmListTagsForCertificate    bool
	_acmPutAccountConfiguration   bool
	_acmRemoveTagsFromCertificate bool
	_acmRenewCertificate          bool
	_acmRequestCertificate        bool
	_acmResendValidationEmail     bool
	_acmRevokeCertificate         bool
	_acmUpdateCertificateOptions  bool

	_acmCertificate             string
	_acmCertificateArn          string
	_acmCertificateAuthorityArn string
	_acmCertificateChain        string
	_acmCertificateStatuses     string
	_acmDomain                  string
	_acmDomainName              string
	_acmDomainValidationOptions string
	_acmExpiryEvents            string
	_acmIdempotencyToken        string
	_acmIncludes                string
	_acmKeyAlgorithm            string
	_acmManagedBy               string
	_acmMaxItems                string
	_acmNextToken               string
	_acmOptions                 string
	_acmPassphrase              string
	_acmPrivateKey              string
	_acmRevocationReason        string
	_acmSortBy                  string
	_acmSortOrder               string
	_acmSubjectAlternativeNames []string
	_acmTags                    string
	_acmValidationDomain        string
	_acmValidationMethod        string
)

// Adds one or more tags to an ACM certificate. Tags are labels that you can use
// to identify and organize your Amazon Web Services resources. Each tag consists
// of a key and an optional value . You specify the certificate on input by its
// Amazon Resource Name (ARN). You specify the tag by using a key-value pair.
//
// You can apply a tag to just one certificate if you want to identify a specific
// characteristic of that certificate, or you can apply the same tag to multiple
// certificates if you want to filter for a common relationship among those
// certificates. Similarly, you can apply the same tag to multiple resources if you
// want to specify a relationship among those resources. For example, you can add
// the same tag to an ACM certificate and an Elastic Load Balancing load balancer
// to indicate that they are both used by the same website. For more information,
// see [Tagging ACM certificates].
//
// To remove one or more tags, use the RemoveTagsFromCertificate action. To view all of the tags that have
// been applied to the certificate, use the ListTagsForCertificateaction.
//
// [Tagging ACM certificates]: https://docs.aws.amazon.com/acm/latest/userguide/tags.html
func acm_AddTagsToCertificate(cfg aws.Config, client *acm.Client) {
	input := &acm.AddTagsToCertificateInput{
		// CertificateArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_acmCertificateArn) > 0 {
		input.CertificateArn = aws.String(_acmCertificateArn)
	}
	if len(_acmTags) > 0 {
		if err := assignInputField(input, "Tags", _acmTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddTagsToCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a certificate and its associated private key. If this action succeeds,
// the certificate no longer appears in the list that can be displayed by calling
// the ListCertificatesaction or be retrieved by calling the GetCertificate action. The certificate will not be
// available for use by Amazon Web Services services integrated with ACM.
//
// You cannot delete an ACM certificate that is being used by another Amazon Web
// Services service. To delete a certificate that is in use, the certificate
// association must first be removed.
func acm_DeleteCertificate(cfg aws.Config, client *acm.Client) {
	input := &acm.DeleteCertificateInput{
		// CertificateArn: *string, // Required
	}

	if len(_acmCertificateArn) > 0 {
		input.CertificateArn = aws.String(_acmCertificateArn)
	}

	if resp, err := client.DeleteCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns detailed metadata about the specified ACM certificate.
// If you have just created a certificate using the RequestCertificate action,
// there is a delay of several seconds before you can retrieve information about
// it.
func acm_DescribeCertificate(cfg aws.Config, client *acm.Client) {
	input := &acm.DescribeCertificateInput{
		// CertificateArn: *string, // Required
	}

	if len(_acmCertificateArn) > 0 {
		input.CertificateArn = aws.String(_acmCertificateArn)
	}

	if resp, err := client.DescribeCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Exports a private certificate issued by a private certificate authority (CA) or
// public certificate for use anywhere. The exported file contains the certificate,
// the certificate chain, and the encrypted private key associated with the public
// key that is embedded in the certificate. For security, you must assign a
// passphrase for the private key when exporting it.
//
// For information about exporting and formatting a certificate using the ACM
// console or CLI, see [Export a private certificate]and [Export a public certificate].
//
// [Export a public certificate]: https://docs.aws.amazon.com/acm/latest/userguide/export-public-certificate
// [Export a private certificate]: https://docs.aws.amazon.com/acm/latest/userguide/export-private.html
func acm_ExportCertificate(cfg aws.Config, client *acm.Client) {
	input := &acm.ExportCertificateInput{
		// CertificateArn: *string, // Required
		// Passphrase: []byte, // Required
	}

	if len(_acmCertificateArn) > 0 {
		input.CertificateArn = aws.String(_acmCertificateArn)
	}
	if len(_acmPassphrase) > 0 {
		if err := assignInputField(input, "Passphrase", _acmPassphrase); err != nil {
			log.Errorf("invalid --passphrase: %s", err.Error())
			return
		}
	}

	if resp, err := client.ExportCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the account configuration options associated with an Amazon Web
// Services account.
func acm_GetAccountConfiguration(cfg aws.Config, client *acm.Client) {
	input := &acm.GetAccountConfigurationInput{}

	if resp, err := client.GetAccountConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a certificate and its certificate chain. The certificate may be
// either a public or private certificate issued using the ACM RequestCertificate
// action, or a certificate imported into ACM using the ImportCertificate action.
// The chain consists of the certificate of the issuing CA and the intermediate
// certificates of any other subordinate CAs. All of the certificates are base64
// encoded. You can use [OpenSSL]to decode the certificates and inspect individual fields.
//
// [OpenSSL]: https://wiki.openssl.org/index.php/Command_Line_Utilities
func acm_GetCertificate(cfg aws.Config, client *acm.Client) {
	input := &acm.GetCertificateInput{
		// CertificateArn: *string, // Required
	}

	if len(_acmCertificateArn) > 0 {
		input.CertificateArn = aws.String(_acmCertificateArn)
	}

	if resp, err := client.GetCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports a certificate into Certificate Manager (ACM) to use with services that
// are integrated with ACM. Note that [integrated services]allow only certificate types and keys they
// support to be associated with their resources. Further, their support differs
// depending on whether the certificate is imported into IAM or into ACM. For more
// information, see the documentation for each service. For more information about
// importing certificates into ACM, see [Importing Certificates]in the Certificate Manager User Guide.
//
// ACM does not provide [managed renewal] for certificates that you import.
//
// Note the following guidelines when importing third party certificates:
//
// - You must enter the private key that matches the certificate you are
// importing.
//
// - The private key must be unencrypted. You cannot import a private key that
// is protected by a password or a passphrase.
//
// - The private key must be no larger than 5 KB (5,120 bytes).
//
// - The certificate, private key, and certificate chain must be PEM-encoded.
//
// - The current time must be between the Not Before and Not After certificate
// fields.
//
// - The Issuer field must not be empty.
//
// - The OCSP authority URL, if present, must not exceed 1000 characters.
//
// - To import a new certificate, omit the CertificateArn argument. Include this
// argument only when you want to replace a previously imported certificate.
//
// - When you import a certificate by using the CLI, you must specify the
// certificate, the certificate chain, and the private key by their file names
// preceded by fileb:// . For example, you can specify a certificate saved in the
// C:\temp folder as fileb://C:\temp\certificate_to_import.pem . If you are
// making an HTTP or HTTPS Query request, include these arguments as BLOBs.
//
// - When you import a certificate by using an SDK, you must specify the
// certificate, the certificate chain, and the private key files in the manner
// required by the programming language you're using.
//
// - The cryptographic algorithm of an imported certificate must match the
// algorithm of the signing CA. For example, if the signing CA key type is RSA,
// then the certificate key type must also be RSA.
//
// This operation returns the [Amazon Resource Name (ARN)] of the imported certificate.
//
// [Importing Certificates]: https://docs.aws.amazon.com/acm/latest/userguide/import-certificate.html
// [integrated services]: https://docs.aws.amazon.com/acm/latest/userguide/acm-services.html
// [Amazon Resource Name (ARN)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
// [managed renewal]: https://docs.aws.amazon.com/acm/latest/userguide/acm-renewal.html
func acm_ImportCertificate(cfg aws.Config, client *acm.Client) {
	input := &acm.ImportCertificateInput{
		// Certificate: []byte, // Required
		// PrivateKey: []byte, // Required
	}

	if len(_acmCertificate) > 0 {
		if err := assignInputField(input, "Certificate", _acmCertificate); err != nil {
			log.Errorf("invalid --certificate: %s", err.Error())
			return
		}
	}
	if len(_acmPrivateKey) > 0 {
		if err := assignInputField(input, "PrivateKey", _acmPrivateKey); err != nil {
			log.Errorf("invalid --private-key: %s", err.Error())
			return
		}
	}
	if len(_acmCertificateArn) > 0 {
		input.CertificateArn = aws.String(_acmCertificateArn)
	}
	if len(_acmCertificateChain) > 0 {
		if err := assignInputField(input, "CertificateChain", _acmCertificateChain); err != nil {
			log.Errorf("invalid --certificate-chain: %s", err.Error())
			return
		}
	}
	if len(_acmTags) > 0 {
		if err := assignInputField(input, "Tags", _acmTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of certificate ARNs and domain names. You can request that
// only certificates that match a specific status be listed. You can also filter by
// specific attributes of the certificate. Default filtering returns only RSA_2048
// certificates. For more information, see Filters.
func acm_ListCertificates(cfg aws.Config, client *acm.Client) {
	input := &acm.ListCertificatesInput{}

	if len(_acmCertificateStatuses) > 0 {
		if err := assignInputField(input, "CertificateStatuses", _acmCertificateStatuses); err != nil {
			log.Errorf("invalid --certificate-statuses: %s", err.Error())
			return
		}
	}
	if len(_acmIncludes) > 0 {
		if err := assignInputField(input, "Includes", _acmIncludes); err != nil {
			log.Errorf("invalid --includes: %s", err.Error())
			return
		}
	}
	if len(_acmMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _acmMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_acmNextToken) > 0 {
		input.NextToken = aws.String(_acmNextToken)
	}
	if len(_acmSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _acmSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_acmSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _acmSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCertificates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*acm.ListCertificatesOutput
	p := acm.NewListCertificatesPaginator(client, input)
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

// Lists the tags that have been applied to the ACM certificate. Use the
// certificate's Amazon Resource Name (ARN) to specify the certificate. To add a
// tag to an ACM certificate, use the AddTagsToCertificateaction. To delete a tag, use the RemoveTagsFromCertificate action.
func acm_ListTagsForCertificate(cfg aws.Config, client *acm.Client) {
	input := &acm.ListTagsForCertificateInput{
		// CertificateArn: *string, // Required
	}

	if len(_acmCertificateArn) > 0 {
		input.CertificateArn = aws.String(_acmCertificateArn)
	}

	if resp, err := client.ListTagsForCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or modifies account-level configurations in ACM.
// The supported configuration option is DaysBeforeExpiry . This option specifies
// the number of days prior to certificate expiration when ACM starts generating
// EventBridge events. ACM sends one event per day per certificate until the
// certificate expires. By default, accounts receive events starting 45 days before
// certificate expiration.
func acm_PutAccountConfiguration(cfg aws.Config, client *acm.Client) {
	input := &acm.PutAccountConfigurationInput{
		// IdempotencyToken: *string, // Required
	}

	if len(_acmIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_acmIdempotencyToken)
	}
	if len(_acmExpiryEvents) > 0 {
		if err := assignInputField(input, "ExpiryEvents", _acmExpiryEvents); err != nil {
			log.Errorf("invalid --expiry-events: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutAccountConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Remove one or more tags from an ACM certificate. A tag consists of a key-value
// pair. If you do not specify the value portion of the tag when calling this
// function, the tag will be removed regardless of value. If you specify a value,
// the tag is removed only if it is associated with the specified value.
//
// To add tags to a certificate, use the AddTagsToCertificate action. To view all of the tags that
// have been applied to a specific ACM certificate, use the ListTagsForCertificateaction.
func acm_RemoveTagsFromCertificate(cfg aws.Config, client *acm.Client) {
	input := &acm.RemoveTagsFromCertificateInput{
		// CertificateArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_acmCertificateArn) > 0 {
		input.CertificateArn = aws.String(_acmCertificateArn)
	}
	if len(_acmTags) > 0 {
		if err := assignInputField(input, "Tags", _acmTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.RemoveTagsFromCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Renews an [eligible ACM certificate]. In order to renew your Amazon Web Services Private CA certificates
// with ACM, you must first [grant the ACM service principal permission to do so]. For more information, see [Testing Managed Renewal] in the ACM User Guide.
//
// [Testing Managed Renewal]: https://docs.aws.amazon.com/acm/latest/userguide/manual-renewal.html
// [grant the ACM service principal permission to do so]: https://docs.aws.amazon.com/privateca/latest/userguide/PcaPermissions.html
// [eligible ACM certificate]: https://docs.aws.amazon.com/acm/latest/userguide/managed-renewal.html
func acm_RenewCertificate(cfg aws.Config, client *acm.Client) {
	input := &acm.RenewCertificateInput{
		// CertificateArn: *string, // Required
	}

	if len(_acmCertificateArn) > 0 {
		input.CertificateArn = aws.String(_acmCertificateArn)
	}

	if resp, err := client.RenewCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Requests an ACM certificate for use with other Amazon Web Services services. To
// request an ACM certificate, you must specify a fully qualified domain name
// (FQDN) in the DomainName parameter. You can also specify additional FQDNs in
// the SubjectAlternativeNames parameter.
//
// If you are requesting a private certificate, domain validation is not required.
// If you are requesting a public certificate, each domain name that you specify
// must be validated to verify that you own or control the domain. You can use [DNS validation]or [email validation]
// . We recommend that you use DNS validation.
//
// ACM behavior differs from the [RFC 6125] specification of the certificate validation
// process. ACM first checks for a Subject Alternative Name, and, if it finds one,
// ignores the common name (CN).
//
// After successful completion of the RequestCertificate action, there is a delay
// of several seconds before you can retrieve information about the new
// certificate.
//
// [email validation]: https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-validate-email.html
// [RFC 6125]: https://datatracker.ietf.org/doc/html/rfc6125#appendix-B.2
// [DNS validation]: https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-validate-dns.html
func acm_RequestCertificate(cfg aws.Config, client *acm.Client) {
	input := &acm.RequestCertificateInput{
		// DomainName: *string, // Required
	}

	if len(_acmDomainName) > 0 {
		input.DomainName = aws.String(_acmDomainName)
	}
	if len(_acmCertificateAuthorityArn) > 0 {
		input.CertificateAuthorityArn = aws.String(_acmCertificateAuthorityArn)
	}
	if len(_acmDomainValidationOptions) > 0 {
		if err := assignInputField(input, "DomainValidationOptions", _acmDomainValidationOptions); err != nil {
			log.Errorf("invalid --domain-validation-options: %s", err.Error())
			return
		}
	}
	if len(_acmIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_acmIdempotencyToken)
	}
	if len(_acmKeyAlgorithm) > 0 {
		if err := assignInputField(input, "KeyAlgorithm", _acmKeyAlgorithm); err != nil {
			log.Errorf("invalid --key-algorithm: %s", err.Error())
			return
		}
	}
	if len(_acmManagedBy) > 0 {
		if err := assignInputField(input, "ManagedBy", _acmManagedBy); err != nil {
			log.Errorf("invalid --managed-by: %s", err.Error())
			return
		}
	}
	if len(_acmOptions) > 0 {
		if err := assignInputField(input, "Options", _acmOptions); err != nil {
			log.Errorf("invalid --options: %s", err.Error())
			return
		}
	}
	if len(_acmSubjectAlternativeNames) > 0 {
		input.SubjectAlternativeNames = append([]string(nil), _acmSubjectAlternativeNames...)
	}
	if len(_acmTags) > 0 {
		if err := assignInputField(input, "Tags", _acmTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_acmValidationMethod) > 0 {
		if err := assignInputField(input, "ValidationMethod", _acmValidationMethod); err != nil {
			log.Errorf("invalid --validation-method: %s", err.Error())
			return
		}
	}

	if resp, err := client.RequestCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Resends the email that requests domain ownership validation. The domain owner
// or an authorized representative must approve the ACM certificate before it can
// be issued. The certificate can be approved by clicking a link in the mail to
// navigate to the Amazon certificate approval website and then clicking I Approve.
// However, the validation email can be blocked by spam filters. Therefore, if you
// do not receive the original mail, you can request that the mail be resent within
// 72 hours of requesting the ACM certificate. If more than 72 hours have elapsed
// since your original request or since your last attempt to resend validation
// mail, you must request a new certificate. For more information about setting up
// your contact email addresses, see [Configure Email for your Domain].
//
// [Configure Email for your Domain]: https://docs.aws.amazon.com/acm/latest/userguide/setup-email.html
func acm_ResendValidationEmail(cfg aws.Config, client *acm.Client) {
	input := &acm.ResendValidationEmailInput{
		// CertificateArn: *string, // Required
		// Domain: *string, // Required
		// ValidationDomain: *string, // Required
	}

	if len(_acmCertificateArn) > 0 {
		input.CertificateArn = aws.String(_acmCertificateArn)
	}
	if len(_acmDomain) > 0 {
		input.Domain = aws.String(_acmDomain)
	}
	if len(_acmValidationDomain) > 0 {
		input.ValidationDomain = aws.String(_acmValidationDomain)
	}

	if resp, err := client.ResendValidationEmail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Revokes a public ACM certificate. You can only revoke certificates that have
// been previously exported.
func acm_RevokeCertificate(cfg aws.Config, client *acm.Client) {
	input := &acm.RevokeCertificateInput{
		// CertificateArn: *string, // Required
		// RevocationReason: types.RevocationReason, // Required
	}

	if len(_acmCertificateArn) > 0 {
		input.CertificateArn = aws.String(_acmCertificateArn)
	}
	if len(_acmRevocationReason) > 0 {
		if err := assignInputField(input, "RevocationReason", _acmRevocationReason); err != nil {
			log.Errorf("invalid --revocation-reason: %s", err.Error())
			return
		}
	}

	if resp, err := client.RevokeCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a certificate. You can use this function to specify whether to opt in
// to or out of recording your certificate in a certificate transparency log and
// exporting. For more information, see [Opting Out of Certificate Transparency Logging]and [Certificate Manager Exportable Managed Certificates].
//
// [Opting Out of Certificate Transparency Logging]: https://docs.aws.amazon.com/acm/latest/userguide/acm-bestpractices.html#best-practices-transparency
// [Certificate Manager Exportable Managed Certificates]: https://docs.aws.amazon.com/acm/latest/userguide/acm-exportable-certificates.html
func acm_UpdateCertificateOptions(cfg aws.Config, client *acm.Client) {
	input := &acm.UpdateCertificateOptionsInput{
		// CertificateArn: *string, // Required
		// Options: *types.CertificateOptions, // Required
	}

	if len(_acmCertificateArn) > 0 {
		input.CertificateArn = aws.String(_acmCertificateArn)
	}
	if len(_acmOptions) > 0 {
		if err := assignInputField(input, "Options", _acmOptions); err != nil {
			log.Errorf("invalid --options: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCertificateOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_acmCmd)
	_acmCmd.Flags().SortFlags = false

	_acmCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_acmCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_acmCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_acmCmd.Flags().StringVarP(&_acmCertificate, "certificate", "", "", "Certificate")
	_acmCmd.Flags().StringVarP(&_acmCertificateArn, "certificate-arn", "", "", "Certificate ARN")
	_acmCmd.Flags().StringVarP(&_acmCertificateAuthorityArn, "certificate-authority-arn", "", "", "Certificate Authority ARN")
	_acmCmd.Flags().StringVarP(&_acmCertificateChain, "certificate-chain", "", "", "Certificate Chain")
	_acmCmd.Flags().StringVarP(&_acmCertificateStatuses, "certificate-statuses", "", "", "Certificate Statuses")
	_acmCmd.Flags().StringVarP(&_acmDomain, "domain", "", "", "Domain")
	_acmCmd.Flags().StringVarP(&_acmDomainName, "domain-name", "", "", "Domain Name")
	_acmCmd.Flags().StringVarP(&_acmDomainValidationOptions, "domain-validation-options", "", "", "Domain Validation Options")
	_acmCmd.Flags().StringVarP(&_acmExpiryEvents, "expiry-events", "", "", "Expiry Events")
	_acmCmd.Flags().StringVarP(&_acmIdempotencyToken, "idempotency-token", "", "", "Idempotency Token")
	_acmCmd.Flags().StringVarP(&_acmIncludes, "includes", "", "", "Includes")
	_acmCmd.Flags().StringVarP(&_acmKeyAlgorithm, "key-algorithm", "", "", "Key Algorithm")
	_acmCmd.Flags().StringVarP(&_acmManagedBy, "managed-by", "", "", "Managed By")
	_acmCmd.Flags().StringVarP(&_acmMaxItems, "max-items", "", "", "Max Items")
	_acmCmd.Flags().StringVarP(&_acmNextToken, "next-token", "", "", "Next Token")
	_acmCmd.Flags().StringVarP(&_acmOptions, "options", "", "", "Options")
	_acmCmd.Flags().StringVarP(&_acmPassphrase, "passphrase", "", "", "Passphrase")
	_acmCmd.Flags().StringVarP(&_acmPrivateKey, "private-key", "", "", "Private Key")
	_acmCmd.Flags().StringVarP(&_acmRevocationReason, "revocation-reason", "", "", "Revocation Reason")
	_acmCmd.Flags().StringVarP(&_acmSortBy, "sort-by", "", "", "Sort By")
	_acmCmd.Flags().StringVarP(&_acmSortOrder, "sort-order", "", "", "Sort Order")
	_acmCmd.Flags().StringSliceVarP(&_acmSubjectAlternativeNames, "subject-alternative-names", "", nil, "Subject Alternative Names")
	_acmCmd.Flags().StringVarP(&_acmTags, "tags", "", "", "Tags")
	_acmCmd.Flags().StringVarP(&_acmValidationDomain, "validation-domain", "", "", "Validation Domain")
	_acmCmd.Flags().StringVarP(&_acmValidationMethod, "validation-method", "", "", "Validation Method")

	_acmCmd.Flags().BoolVarP(&_acmAddTagsToCertificate, "add-tags-to-certificate", "", false, "Add Tags To Certificate")
	_acmCmd.Flags().BoolVarP(&_acmDeleteCertificate, "delete-certificate", "", false, "Delete Certificate")
	_acmCmd.Flags().BoolVarP(&_acmDescribeCertificate, "describe-certificate", "", false, "Describe Certificate")
	_acmCmd.Flags().BoolVarP(&_acmExportCertificate, "export-certificate", "", false, "Export Certificate")
	_acmCmd.Flags().BoolVarP(&_acmGetAccountConfiguration, "get-account-configuration", "", false, "Get Account Configuration")
	_acmCmd.Flags().BoolVarP(&_acmGetCertificate, "get-certificate", "", false, "Get Certificate")
	_acmCmd.Flags().BoolVarP(&_acmImportCertificate, "import-certificate", "", false, "Import Certificate")
	_acmCmd.Flags().BoolVarP(&_acmListCertificates, "list-certificates", "", false, "List Certificates")
	_acmCmd.Flags().BoolVarP(&_acmListTagsForCertificate, "list-tags-for-certificate", "", false, "List Tags For Certificate")
	_acmCmd.Flags().BoolVarP(&_acmPutAccountConfiguration, "put-account-configuration", "", false, "Put Account Configuration")
	_acmCmd.Flags().BoolVarP(&_acmRemoveTagsFromCertificate, "remove-tags-from-certificate", "", false, "Remove Tags From Certificate")
	_acmCmd.Flags().BoolVarP(&_acmRenewCertificate, "renew-certificate", "", false, "Renew Certificate")
	_acmCmd.Flags().BoolVarP(&_acmRequestCertificate, "request-certificate", "", false, "Request Certificate")
	_acmCmd.Flags().BoolVarP(&_acmResendValidationEmail, "resend-validation-email", "", false, "Resend Validation Email")
	_acmCmd.Flags().BoolVarP(&_acmRevokeCertificate, "revoke-certificate", "", false, "Revoke Certificate")
	_acmCmd.Flags().BoolVarP(&_acmUpdateCertificateOptions, "update-certificate-options", "", false, "Update Certificate Options")

}
