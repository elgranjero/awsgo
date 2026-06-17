package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/acmpca"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// acmpcaCmd represents the acmpca command
var _acmpcaCmd = &cobra.Command{
	Use:   "acmpca",
	Short: "AWS acmpca CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := acmpca.NewFromConfig(cfg)
		if _acmpcaCreateCertificateAuthority {
			acmpca_CreateCertificateAuthority(cfg, client)
			return
		}
		if _acmpcaCreateCertificateAuthorityAuditReport {
			acmpca_CreateCertificateAuthorityAuditReport(cfg, client)
			return
		}
		if _acmpcaCreatePermission {
			acmpca_CreatePermission(cfg, client)
			return
		}
		if _acmpcaDeleteCertificateAuthority {
			acmpca_DeleteCertificateAuthority(cfg, client)
			return
		}
		if _acmpcaDeletePermission {
			acmpca_DeletePermission(cfg, client)
			return
		}
		if _acmpcaDeletePolicy {
			acmpca_DeletePolicy(cfg, client)
			return
		}
		if _acmpcaDescribeCertificateAuthority {
			acmpca_DescribeCertificateAuthority(cfg, client)
			return
		}
		if _acmpcaDescribeCertificateAuthorityAuditReport {
			acmpca_DescribeCertificateAuthorityAuditReport(cfg, client)
			return
		}
		if _acmpcaGetCertificate {
			acmpca_GetCertificate(cfg, client)
			return
		}
		if _acmpcaGetCertificateAuthorityCertificate {
			acmpca_GetCertificateAuthorityCertificate(cfg, client)
			return
		}
		if _acmpcaGetCertificateAuthorityCsr {
			acmpca_GetCertificateAuthorityCsr(cfg, client)
			return
		}
		if _acmpcaGetPolicy {
			acmpca_GetPolicy(cfg, client)
			return
		}
		if _acmpcaImportCertificateAuthorityCertificate {
			acmpca_ImportCertificateAuthorityCertificate(cfg, client)
			return
		}
		if _acmpcaIssueCertificate {
			acmpca_IssueCertificate(cfg, client)
			return
		}
		if _acmpcaListCertificateAuthorities {
			acmpca_ListCertificateAuthorities(cfg, client)
			return
		}
		if _acmpcaListPermissions {
			acmpca_ListPermissions(cfg, client)
			return
		}
		if _acmpcaListTags {
			acmpca_ListTags(cfg, client)
			return
		}
		if _acmpcaPutPolicy {
			acmpca_PutPolicy(cfg, client)
			return
		}
		if _acmpcaRestoreCertificateAuthority {
			acmpca_RestoreCertificateAuthority(cfg, client)
			return
		}
		if _acmpcaRevokeCertificate {
			acmpca_RevokeCertificate(cfg, client)
			return
		}
		if _acmpcaTagCertificateAuthority {
			acmpca_TagCertificateAuthority(cfg, client)
			return
		}
		if _acmpcaUntagCertificateAuthority {
			acmpca_UntagCertificateAuthority(cfg, client)
			return
		}
		if _acmpcaUpdateCertificateAuthority {
			acmpca_UpdateCertificateAuthority(cfg, client)
			return
		}

	},
}

var (
	_acmpcaCreateCertificateAuthority              bool
	_acmpcaCreateCertificateAuthorityAuditReport   bool
	_acmpcaCreatePermission                        bool
	_acmpcaDeleteCertificateAuthority              bool
	_acmpcaDeletePermission                        bool
	_acmpcaDeletePolicy                            bool
	_acmpcaDescribeCertificateAuthority            bool
	_acmpcaDescribeCertificateAuthorityAuditReport bool
	_acmpcaGetCertificate                          bool
	_acmpcaGetCertificateAuthorityCertificate      bool
	_acmpcaGetCertificateAuthorityCsr              bool
	_acmpcaGetPolicy                               bool
	_acmpcaImportCertificateAuthorityCertificate   bool
	_acmpcaIssueCertificate                        bool
	_acmpcaListCertificateAuthorities              bool
	_acmpcaListPermissions                         bool
	_acmpcaListTags                                bool
	_acmpcaPutPolicy                               bool
	_acmpcaRestoreCertificateAuthority             bool
	_acmpcaRevokeCertificate                       bool
	_acmpcaTagCertificateAuthority                 bool
	_acmpcaUntagCertificateAuthority               bool
	_acmpcaUpdateCertificateAuthority              bool

	_acmpcaActions                           string
	_acmpcaApiPassthrough                    string
	_acmpcaAuditReportId                     string
	_acmpcaAuditReportResponseFormat         string
	_acmpcaCertificate                       string
	_acmpcaCertificateArn                    string
	_acmpcaCertificateAuthorityArn           string
	_acmpcaCertificateAuthorityConfiguration string
	_acmpcaCertificateAuthorityType          string
	_acmpcaCertificateChain                  string
	_acmpcaCertificateSerial                 string
	_acmpcaCsr                               string
	_acmpcaIdempotencyToken                  string
	_acmpcaKeyStorageSecurityStandard        string
	_acmpcaMaxResults                        string
	_acmpcaNextToken                         string
	_acmpcaPermanentDeletionTimeInDays       string
	_acmpcaPolicy                            string
	_acmpcaPrincipal                         string
	_acmpcaResourceArn                       string
	_acmpcaResourceOwner                     string
	_acmpcaRevocationConfiguration           string
	_acmpcaRevocationReason                  string
	_acmpcaS3BucketName                      string
	_acmpcaSigningAlgorithm                  string
	_acmpcaSourceAccount                     string
	_acmpcaStatus                            string
	_acmpcaTags                              string
	_acmpcaTemplateArn                       string
	_acmpcaUsageMode                         string
	_acmpcaValidity                          string
	_acmpcaValidityNotBefore                 string
)

// Creates a root or subordinate private certificate authority (CA). You must
// specify the CA configuration, an optional configuration for Online Certificate
// Status Protocol (OCSP) and/or a certificate revocation list (CRL), the CA type,
// and an optional idempotency token to avoid accidental creation of multiple CAs.
// The CA configuration specifies the name of the algorithm and key size to be used
// to create the CA private key, the type of signing algorithm that the CA uses,
// and X.500 subject information. The OCSP configuration can optionally specify a
// custom URL for the OCSP responder. The CRL configuration specifies the CRL
// expiration period in days (the validity period of the CRL), the Amazon S3 bucket
// that will contain the CRL, and a CNAME alias for the S3 bucket that is included
// in certificates issued by the CA. If successful, this action returns the Amazon
// Resource Name (ARN) of the CA.
//
// Both Amazon Web Services Private CA and the IAM principal must have permission
// to write to the S3 bucket that you specify. If the IAM principal making the call
// does not have permission to write to the bucket, then an exception is thrown.
// For more information, see [Access policies for CRLs in Amazon S3].
//
// Amazon Web Services Private CA assets that are stored in Amazon S3 can be
// protected with encryption. For more information, see [Encrypting Your CRLs].
//
// [Access policies for CRLs in Amazon S3]: https://docs.aws.amazon.com/privateca/latest/userguide/crl-planning.html#s3-policies
// [Encrypting Your CRLs]: https://docs.aws.amazon.com/privateca/latest/userguide/crl-planning.html#crl-encryption
func acmpca_CreateCertificateAuthority(cfg aws.Config, client *acmpca.Client) {
	input := &acmpca.CreateCertificateAuthorityInput{
		// CertificateAuthorityConfiguration: *types.CertificateAuthorityConfiguration, // Required
		// CertificateAuthorityType: types.CertificateAuthorityType, // Required
	}

	if len(_acmpcaCertificateAuthorityConfiguration) > 0 {
		if err := assignInputField(input, "CertificateAuthorityConfiguration", _acmpcaCertificateAuthorityConfiguration); err != nil {
			log.Errorf("invalid --certificate-authority-configuration: %s", err.Error())
			return
		}
	}
	if len(_acmpcaCertificateAuthorityType) > 0 {
		if err := assignInputField(input, "CertificateAuthorityType", _acmpcaCertificateAuthorityType); err != nil {
			log.Errorf("invalid --certificate-authority-type: %s", err.Error())
			return
		}
	}
	if len(_acmpcaIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_acmpcaIdempotencyToken)
	}
	if len(_acmpcaKeyStorageSecurityStandard) > 0 {
		if err := assignInputField(input, "KeyStorageSecurityStandard", _acmpcaKeyStorageSecurityStandard); err != nil {
			log.Errorf("invalid --key-storage-security-standard: %s", err.Error())
			return
		}
	}
	if len(_acmpcaRevocationConfiguration) > 0 {
		if err := assignInputField(input, "RevocationConfiguration", _acmpcaRevocationConfiguration); err != nil {
			log.Errorf("invalid --revocation-configuration: %s", err.Error())
			return
		}
	}
	if len(_acmpcaTags) > 0 {
		if err := assignInputField(input, "Tags", _acmpcaTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_acmpcaUsageMode) > 0 {
		if err := assignInputField(input, "UsageMode", _acmpcaUsageMode); err != nil {
			log.Errorf("invalid --usage-mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCertificateAuthority(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an audit report that lists every time that your CA private key is used
// to issue a certificate. The [IssueCertificate]and [RevokeCertificate] actions use the private key.
//
// To save the audit report to your designated Amazon S3 bucket, you must create a
// bucket policy that grants Amazon Web Services Private CA permission to access
// and write to it. For an example policy, see [Prepare an Amazon S3 bucket for audit reports].
//
// Amazon Web Services Private CA assets that are stored in Amazon S3 can be
// protected with encryption. For more information, see [Encrypting Your Audit Reports].
//
// You can generate a maximum of one report every 30 minutes.
//
// [RevokeCertificate]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_RevokeCertificate.html
// [Encrypting Your Audit Reports]: https://docs.aws.amazon.com/privateca/latest/userguide/PcaAuditReport.html#audit-report-encryption
// [IssueCertificate]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_IssueCertificate.html
// [Prepare an Amazon S3 bucket for audit reports]: https://docs.aws.amazon.com/privateca/latest/userguide/PcaAuditReport.html#s3-access
func acmpca_CreateCertificateAuthorityAuditReport(cfg aws.Config, client *acmpca.Client) {
	input := &acmpca.CreateCertificateAuthorityAuditReportInput{
		// AuditReportResponseFormat: types.AuditReportResponseFormat, // Required
		// CertificateAuthorityArn: *string, // Required
		// S3BucketName: *string, // Required
	}

	if len(_acmpcaAuditReportResponseFormat) > 0 {
		if err := assignInputField(input, "AuditReportResponseFormat", _acmpcaAuditReportResponseFormat); err != nil {
			log.Errorf("invalid --audit-report-response-format: %s", err.Error())
			return
		}
	}
	if len(_acmpcaCertificateAuthorityArn) > 0 {
		input.CertificateAuthorityArn = aws.String(_acmpcaCertificateAuthorityArn)
	}
	if len(_acmpcaS3BucketName) > 0 {
		input.S3BucketName = aws.String(_acmpcaS3BucketName)
	}

	if resp, err := client.CreateCertificateAuthorityAuditReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Grants one or more permissions on a private CA to the Certificate Manager (ACM)
// service principal ( acm.amazonaws.com ). These permissions allow ACM to issue
// and renew ACM certificates that reside in the same Amazon Web Services account
// as the CA.
//
// You can list current permissions with the [ListPermissions] action and revoke them with the [DeletePermission]
// action.
//
// # About Permissions
//
// - If the private CA and the certificates it issues reside in the same
// account, you can use CreatePermission to grant permissions for ACM to carry
// out automatic certificate renewals.
//
// - For automatic certificate renewal to succeed, the ACM service principal
// needs permissions to create, retrieve, and list certificates.
//
// - If the private CA and the ACM certificates reside in different accounts,
// then permissions cannot be used to enable automatic renewals. Instead, the ACM
// certificate owner must set up a resource-based policy to enable cross-account
// issuance and renewals. For more information, see [Using a Resource Based Policy with Amazon Web Services Private CA].
//
// [Using a Resource Based Policy with Amazon Web Services Private CA]: https://docs.aws.amazon.com/privateca/latest/userguide/pca-rbp.html
// [ListPermissions]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_ListPermissions.html
// [DeletePermission]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_DeletePermission.html
func acmpca_CreatePermission(cfg aws.Config, client *acmpca.Client) {
	input := &acmpca.CreatePermissionInput{
		// Actions: []types.ActionType, // Required
		// CertificateAuthorityArn: *string, // Required
		// Principal: *string, // Required
	}

	if len(_acmpcaActions) > 0 {
		if err := assignInputField(input, "Actions", _acmpcaActions); err != nil {
			log.Errorf("invalid --actions: %s", err.Error())
			return
		}
	}
	if len(_acmpcaCertificateAuthorityArn) > 0 {
		input.CertificateAuthorityArn = aws.String(_acmpcaCertificateAuthorityArn)
	}
	if len(_acmpcaPrincipal) > 0 {
		input.Principal = aws.String(_acmpcaPrincipal)
	}
	if len(_acmpcaSourceAccount) > 0 {
		input.SourceAccount = aws.String(_acmpcaSourceAccount)
	}

	if resp, err := client.CreatePermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a private certificate authority (CA). You must provide the Amazon
// Resource Name (ARN) of the private CA that you want to delete. You can find the
// ARN by calling the [ListCertificateAuthorities]action.
//
// Deleting a CA will invalidate other CAs and certificates below it in your CA
// hierarchy.
//
// Before you can delete a CA that you have created and activated, you must
// disable it. To do this, call the [UpdateCertificateAuthority]action and set the CertificateAuthorityStatus
// parameter to DISABLED .
//
// Additionally, you can delete a CA if you are waiting for it to be created (that
// is, the status of the CA is CREATING ). You can also delete it if the CA has
// been created but you haven't yet imported the signed certificate into Amazon Web
// Services Private CA (that is, the status of the CA is PENDING_CERTIFICATE ).
//
// When you successfully call [DeleteCertificateAuthority], the CA's status changes to DELETED . However, the
// CA won't be permanently deleted until the restoration period has passed. By
// default, if you do not set the PermanentDeletionTimeInDays parameter, the CA
// remains restorable for 30 days. You can set the parameter from 7 to 30 days. The
// [DescribeCertificateAuthority]action returns the time remaining in the restoration window of a private CA in
// the DELETED state. To restore an eligible CA, call the [RestoreCertificateAuthority] action.
//
// A private CA can be deleted if it is in the PENDING_CERTIFICATE , CREATING ,
// EXPIRED , DISABLED , or FAILED state. To delete a CA in the ACTIVE state, you
// must first disable it, or else the delete request results in an exception. If
// you are deleting a private CA in the PENDING_CERTIFICATE or DISABLED state, you
// can set the length of its restoration period to 7-30 days. The default is 30.
// During this time, the status is set to DELETED and the CA can be restored. A
// private CA deleted in the CREATING or FAILED state has no assigned restoration
// period and cannot be restored.
//
// [ListCertificateAuthorities]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_ListCertificateAuthorities.html
// [RestoreCertificateAuthority]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_RestoreCertificateAuthority.html
// [UpdateCertificateAuthority]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_UpdateCertificateAuthority.html
// [DeleteCertificateAuthority]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_DeleteCertificateAuthority.html
// [DescribeCertificateAuthority]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_DescribeCertificateAuthority.html
func acmpca_DeleteCertificateAuthority(cfg aws.Config, client *acmpca.Client) {
	input := &acmpca.DeleteCertificateAuthorityInput{
		// CertificateAuthorityArn: *string, // Required
	}

	if len(_acmpcaCertificateAuthorityArn) > 0 {
		input.CertificateAuthorityArn = aws.String(_acmpcaCertificateAuthorityArn)
	}
	if len(_acmpcaPermanentDeletionTimeInDays) > 0 {
		if err := assignInputField(input, "PermanentDeletionTimeInDays", _acmpcaPermanentDeletionTimeInDays); err != nil {
			log.Errorf("invalid --permanent-deletion-time-in-days: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteCertificateAuthority(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Revokes permissions on a private CA granted to the Certificate Manager (ACM)
// service principal (acm.amazonaws.com).
//
// These permissions allow ACM to issue and renew ACM certificates that reside in
// the same Amazon Web Services account as the CA. If you revoke these permissions,
// ACM will no longer renew the affected certificates automatically.
//
// Permissions can be granted with the [CreatePermission] action and listed with the [ListPermissions] action.
//
// # About Permissions
//
// - If the private CA and the certificates it issues reside in the same
// account, you can use CreatePermission to grant permissions for ACM to carry
// out automatic certificate renewals.
//
// - For automatic certificate renewal to succeed, the ACM service principal
// needs permissions to create, retrieve, and list certificates.
//
// - If the private CA and the ACM certificates reside in different accounts,
// then permissions cannot be used to enable automatic renewals. Instead, the ACM
// certificate owner must set up a resource-based policy to enable cross-account
// issuance and renewals. For more information, see [Using a Resource Based Policy with Amazon Web Services Private CA].
//
// [Using a Resource Based Policy with Amazon Web Services Private CA]: https://docs.aws.amazon.com/privateca/latest/userguide/pca-rbp.html
// [CreatePermission]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreatePermission.html
// [ListPermissions]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_ListPermissions.html
func acmpca_DeletePermission(cfg aws.Config, client *acmpca.Client) {
	input := &acmpca.DeletePermissionInput{
		// CertificateAuthorityArn: *string, // Required
		// Principal: *string, // Required
	}

	if len(_acmpcaCertificateAuthorityArn) > 0 {
		input.CertificateAuthorityArn = aws.String(_acmpcaCertificateAuthorityArn)
	}
	if len(_acmpcaPrincipal) > 0 {
		input.Principal = aws.String(_acmpcaPrincipal)
	}
	if len(_acmpcaSourceAccount) > 0 {
		input.SourceAccount = aws.String(_acmpcaSourceAccount)
	}

	if resp, err := client.DeletePermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the resource-based policy attached to a private CA. Deletion will
// remove any access that the policy has granted. If there is no policy attached to
// the private CA, this action will return successful.
//
// If you delete a policy that was applied through Amazon Web Services Resource
// Access Manager (RAM), the CA will be removed from all shares in which it was
// included.
//
// The Certificate Manager Service Linked Role that the policy supports is not
// affected when you delete the policy.
//
// The current policy can be shown with [GetPolicy] and updated with [PutPolicy].
//
// # About Policies
//
// - A policy grants access on a private CA to an Amazon Web Services customer
// account, to Amazon Web Services Organizations, or to an Amazon Web Services
// Organizations unit. Policies are under the control of a CA administrator. For
// more information, see [Using a Resource Based Policy with Amazon Web Services Private CA].
//
// - A policy permits a user of Certificate Manager (ACM) to issue ACM
// certificates signed by a CA in another account.
//
// - For ACM to manage automatic renewal of these certificates, the ACM user
// must configure a Service Linked Role (SLR). The SLR allows the ACM service to
// assume the identity of the user, subject to confirmation against the Amazon Web
// Services Private CA policy. For more information, see [Using a Service Linked Role with ACM].
//
// - Updates made in Amazon Web Services Resource Manager (RAM) are reflected in
// policies. For more information, see [Attach a Policy for Cross-Account Access].
//
// [PutPolicy]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_PutPolicy.html
// [Using a Resource Based Policy with Amazon Web Services Private CA]: https://docs.aws.amazon.com/privateca/latest/userguide/pca-rbp.html
// [Using a Service Linked Role with ACM]: https://docs.aws.amazon.com/acm/latest/userguide/acm-slr.html
// [Attach a Policy for Cross-Account Access]: https://docs.aws.amazon.com/privateca/latest/userguide/pca-ram.html
// [GetPolicy]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_GetPolicy.html
func acmpca_DeletePolicy(cfg aws.Config, client *acmpca.Client) {
	input := &acmpca.DeletePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_acmpcaResourceArn) > 0 {
		input.ResourceArn = aws.String(_acmpcaResourceArn)
	}

	if resp, err := client.DeletePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists information about your private certificate authority (CA) or one that has
// been shared with you. You specify the private CA on input by its ARN (Amazon
// Resource Name). The output contains the status of your CA. This can be any of
// the following:
//
// - CREATING - Amazon Web Services Private CA is creating your private
// certificate authority.
//
// - PENDING_CERTIFICATE - The certificate is pending. You must use your Amazon
// Web Services Private CA-hosted or on-premises root or subordinate CA to sign
// your private CA CSR and then import it into Amazon Web Services Private CA.
//
// - ACTIVE - Your private CA is active.
//
// - DISABLED - Your private CA has been disabled.
//
// - EXPIRED - Your private CA certificate has expired.
//
// - FAILED - Your private CA has failed. Your CA can fail because of problems
// such a network outage or back-end Amazon Web Services failure or other errors. A
// failed CA can never return to the pending state. You must create a new CA.
//
// - DELETED - Your private CA is within the restoration period, after which it
// is permanently deleted. The length of time remaining in the CA's restoration
// period is also included in this action's output.
func acmpca_DescribeCertificateAuthority(cfg aws.Config, client *acmpca.Client) {
	input := &acmpca.DescribeCertificateAuthorityInput{
		// CertificateAuthorityArn: *string, // Required
	}

	if len(_acmpcaCertificateAuthorityArn) > 0 {
		input.CertificateAuthorityArn = aws.String(_acmpcaCertificateAuthorityArn)
	}

	if resp, err := client.DescribeCertificateAuthority(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists information about a specific audit report created by calling the [CreateCertificateAuthorityAuditReport] action.
// Audit information is created every time the certificate authority (CA) private
// key is used. The private key is used when you call the [IssueCertificate]action or the [RevokeCertificate] action.
//
// [RevokeCertificate]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_RevokeCertificate.html
// [IssueCertificate]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_IssueCertificate.html
// [CreateCertificateAuthorityAuditReport]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthorityAuditReport.html
func acmpca_DescribeCertificateAuthorityAuditReport(cfg aws.Config, client *acmpca.Client) {
	input := &acmpca.DescribeCertificateAuthorityAuditReportInput{
		// AuditReportId: *string, // Required
		// CertificateAuthorityArn: *string, // Required
	}

	if len(_acmpcaAuditReportId) > 0 {
		input.AuditReportId = aws.String(_acmpcaAuditReportId)
	}
	if len(_acmpcaCertificateAuthorityArn) > 0 {
		input.CertificateAuthorityArn = aws.String(_acmpcaCertificateAuthorityArn)
	}

	if resp, err := client.DescribeCertificateAuthorityAuditReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a certificate from your private CA or one that has been shared with
// you. The ARN of the certificate is returned when you call the [IssueCertificate]action. You must
// specify both the ARN of your private CA and the ARN of the issued certificate
// when calling the GetCertificate action. You can retrieve the certificate if it
// is in the ISSUED, EXPIRED, or REVOKED state. You can call the [CreateCertificateAuthorityAuditReport]action to create
// a report that contains information about all of the certificates issued and
// revoked by your private CA.
//
// [IssueCertificate]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_IssueCertificate.html
// [CreateCertificateAuthorityAuditReport]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthorityAuditReport.html
func acmpca_GetCertificate(cfg aws.Config, client *acmpca.Client) {
	input := &acmpca.GetCertificateInput{
		// CertificateArn: *string, // Required
		// CertificateAuthorityArn: *string, // Required
	}

	if len(_acmpcaCertificateArn) > 0 {
		input.CertificateArn = aws.String(_acmpcaCertificateArn)
	}
	if len(_acmpcaCertificateAuthorityArn) > 0 {
		input.CertificateAuthorityArn = aws.String(_acmpcaCertificateAuthorityArn)
	}

	if resp, err := client.GetCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the certificate and certificate chain for your private certificate
// authority (CA) or one that has been shared with you. Both the certificate and
// the chain are base64 PEM-encoded. The chain does not include the CA certificate.
// Each certificate in the chain signs the one before it.
func acmpca_GetCertificateAuthorityCertificate(cfg aws.Config, client *acmpca.Client) {
	input := &acmpca.GetCertificateAuthorityCertificateInput{
		// CertificateAuthorityArn: *string, // Required
	}

	if len(_acmpcaCertificateAuthorityArn) > 0 {
		input.CertificateAuthorityArn = aws.String(_acmpcaCertificateAuthorityArn)
	}

	if resp, err := client.GetCertificateAuthorityCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the certificate signing request (CSR) for your private certificate
// authority (CA). The CSR is created when you call the [CreateCertificateAuthority]action. Sign the CSR with
// your Amazon Web Services Private CA-hosted or on-premises root or subordinate
// CA. Then import the signed certificate back into Amazon Web Services Private CA
// by calling the [ImportCertificateAuthorityCertificate]action. The CSR is returned as a base64 PEM-encoded string.
//
// [ImportCertificateAuthorityCertificate]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_ImportCertificateAuthorityCertificate.html
// [CreateCertificateAuthority]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html
func acmpca_GetCertificateAuthorityCsr(cfg aws.Config, client *acmpca.Client) {
	input := &acmpca.GetCertificateAuthorityCsrInput{
		// CertificateAuthorityArn: *string, // Required
	}

	if len(_acmpcaCertificateAuthorityArn) > 0 {
		input.CertificateAuthorityArn = aws.String(_acmpcaCertificateAuthorityArn)
	}

	if resp, err := client.GetCertificateAuthorityCsr(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the resource-based policy attached to a private CA. If either the
// private CA resource or the policy cannot be found, this action returns a
// ResourceNotFoundException .
//
// The policy can be attached or updated with [PutPolicy] and removed with [DeletePolicy].
//
// # About Policies
//
// - A policy grants access on a private CA to an Amazon Web Services customer
// account, to Amazon Web Services Organizations, or to an Amazon Web Services
// Organizations unit. Policies are under the control of a CA administrator. For
// more information, see [Using a Resource Based Policy with Amazon Web Services Private CA].
//
// - A policy permits a user of Certificate Manager (ACM) to issue ACM
// certificates signed by a CA in another account.
//
// - For ACM to manage automatic renewal of these certificates, the ACM user
// must configure a Service Linked Role (SLR). The SLR allows the ACM service to
// assume the identity of the user, subject to confirmation against the Amazon Web
// Services Private CA policy. For more information, see [Using a Service Linked Role with ACM].
//
// - Updates made in Amazon Web Services Resource Manager (RAM) are reflected in
// policies. For more information, see [Attach a Policy for Cross-Account Access].
//
// [PutPolicy]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_PutPolicy.html
// [Using a Resource Based Policy with Amazon Web Services Private CA]: https://docs.aws.amazon.com/privateca/latest/userguide/pca-rbp.html
// [Using a Service Linked Role with ACM]: https://docs.aws.amazon.com/acm/latest/userguide/acm-slr.html
// [DeletePolicy]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_DeletePolicy.html
// [Attach a Policy for Cross-Account Access]: https://docs.aws.amazon.com/privateca/latest/userguide/pca-ram.html
func acmpca_GetPolicy(cfg aws.Config, client *acmpca.Client) {
	input := &acmpca.GetPolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_acmpcaResourceArn) > 0 {
		input.ResourceArn = aws.String(_acmpcaResourceArn)
	}

	if resp, err := client.GetPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports a signed private CA certificate into Amazon Web Services Private CA.
// This action is used when you are using a chain of trust whose root is located
// outside Amazon Web Services Private CA. Before you can call this action, the
// following preparations must in place:
//
// - In Amazon Web Services Private CA, call the [CreateCertificateAuthority]action to create the private CA
// that you plan to back with the imported certificate.
//
// - Call the [GetCertificateAuthorityCsr]action to generate a certificate signing request (CSR).
//
// - Sign the CSR using a root or intermediate CA hosted by either an
// on-premises PKI hierarchy or by a commercial CA.
//
// - Create a certificate chain and copy the signed certificate and the
// certificate chain to your working directory.
//
// Amazon Web Services Private CA supports three scenarios for installing a CA
// certificate:
//
// - Installing a certificate for a root CA hosted by Amazon Web Services
// Private CA.
//
// - Installing a subordinate CA certificate whose parent authority is hosted by
// Amazon Web Services Private CA.
//
// - Installing a subordinate CA certificate whose parent authority is
// externally hosted.
//
// The following additional requirements apply when you import a CA certificate.
//
// - Only a self-signed certificate can be imported as a root CA.
//
// - A self-signed certificate cannot be imported as a subordinate CA.
//
// - Your certificate chain must not include the private CA certificate that you
// are importing.
//
// - Your root CA must be the last certificate in your chain. The subordinate
// certificate, if any, that your root CA signed must be next to last. The
// subordinate certificate signed by the preceding subordinate CA must come next,
// and so on until your chain is built.
//
// - The chain must be PEM-encoded.
//
// - The maximum allowed size of a certificate is 32 KB.
//
// - The maximum allowed size of a certificate chain is 2 MB.
//
// # Enforcement of Critical Constraints
//
// Amazon Web Services Private CA allows the following extensions to be marked
// critical in the imported CA certificate or chain.
//
// - Authority key identifier
//
// - Basic constraints (must be marked critical)
//
// - Certificate policies
//
// - Extended key usage
//
// - Inhibit anyPolicy
//
// - Issuer alternative name
//
// - Key usage
//
// - Name constraints
//
// - Policy mappings
//
// - Subject alternative name
//
// - Subject directory attributes
//
// - Subject key identifier
//
// - Subject information access
//
// Amazon Web Services Private CA rejects the following extensions when they are
// marked critical in an imported CA certificate or chain.
//
// - Authority information access
//
// - CRL distribution points
//
// - Freshest CRL
//
// - Policy constraints
//
// Amazon Web Services Private Certificate Authority will also reject any other
// extension marked as critical not contained on the preceding list of allowed
// extensions.
//
// [GetCertificateAuthorityCsr]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_GetCertificateAuthorityCsr.html
// [CreateCertificateAuthority]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html
func acmpca_ImportCertificateAuthorityCertificate(cfg aws.Config, client *acmpca.Client) {
	input := &acmpca.ImportCertificateAuthorityCertificateInput{
		// Certificate: []byte, // Required
		// CertificateAuthorityArn: *string, // Required
	}

	if len(_acmpcaCertificate) > 0 {
		if err := assignInputField(input, "Certificate", _acmpcaCertificate); err != nil {
			log.Errorf("invalid --certificate: %s", err.Error())
			return
		}
	}
	if len(_acmpcaCertificateAuthorityArn) > 0 {
		input.CertificateAuthorityArn = aws.String(_acmpcaCertificateAuthorityArn)
	}
	if len(_acmpcaCertificateChain) > 0 {
		if err := assignInputField(input, "CertificateChain", _acmpcaCertificateChain); err != nil {
			log.Errorf("invalid --certificate-chain: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportCertificateAuthorityCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Uses your private certificate authority (CA), or one that has been shared with
// you, to issue a client certificate. This action returns the Amazon Resource Name
// (ARN) of the certificate. You can retrieve the certificate by calling the [GetCertificate]
// action and specifying the ARN.
//
// You cannot use the ACM ListCertificateAuthorities action to retrieve the ARNs
// of the certificates that you issue by using Amazon Web Services Private CA.
//
// [GetCertificate]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_GetCertificate.html
func acmpca_IssueCertificate(cfg aws.Config, client *acmpca.Client) {
	input := &acmpca.IssueCertificateInput{
		// CertificateAuthorityArn: *string, // Required
		// Csr: []byte, // Required
		// SigningAlgorithm: types.SigningAlgorithm, // Required
		// Validity: *types.Validity, // Required
	}

	if len(_acmpcaCertificateAuthorityArn) > 0 {
		input.CertificateAuthorityArn = aws.String(_acmpcaCertificateAuthorityArn)
	}
	if len(_acmpcaCsr) > 0 {
		if err := assignInputField(input, "Csr", _acmpcaCsr); err != nil {
			log.Errorf("invalid --csr: %s", err.Error())
			return
		}
	}
	if len(_acmpcaSigningAlgorithm) > 0 {
		if err := assignInputField(input, "SigningAlgorithm", _acmpcaSigningAlgorithm); err != nil {
			log.Errorf("invalid --signing-algorithm: %s", err.Error())
			return
		}
	}
	if len(_acmpcaValidity) > 0 {
		if err := assignInputField(input, "Validity", _acmpcaValidity); err != nil {
			log.Errorf("invalid --validity: %s", err.Error())
			return
		}
	}
	if len(_acmpcaApiPassthrough) > 0 {
		if err := assignInputField(input, "ApiPassthrough", _acmpcaApiPassthrough); err != nil {
			log.Errorf("invalid --api-passthrough: %s", err.Error())
			return
		}
	}
	if len(_acmpcaIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_acmpcaIdempotencyToken)
	}
	if len(_acmpcaTemplateArn) > 0 {
		input.TemplateArn = aws.String(_acmpcaTemplateArn)
	}
	if len(_acmpcaValidityNotBefore) > 0 {
		if err := assignInputField(input, "ValidityNotBefore", _acmpcaValidityNotBefore); err != nil {
			log.Errorf("invalid --validity-not-before: %s", err.Error())
			return
		}
	}

	if resp, err := client.IssueCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the private certificate authorities that you created by using the [CreateCertificateAuthority] action.
//
// [CreateCertificateAuthority]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html
func acmpca_ListCertificateAuthorities(cfg aws.Config, client *acmpca.Client) {
	input := &acmpca.ListCertificateAuthoritiesInput{}

	if len(_acmpcaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _acmpcaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_acmpcaNextToken) > 0 {
		input.NextToken = aws.String(_acmpcaNextToken)
	}
	if len(_acmpcaResourceOwner) > 0 {
		if err := assignInputField(input, "ResourceOwner", _acmpcaResourceOwner); err != nil {
			log.Errorf("invalid --resource-owner: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCertificateAuthorities(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*acmpca.ListCertificateAuthoritiesOutput
	p := acmpca.NewListCertificateAuthoritiesPaginator(client, input)
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

// List all permissions on a private CA, if any, granted to the Certificate
// Manager (ACM) service principal (acm.amazonaws.com).
//
// These permissions allow ACM to issue and renew ACM certificates that reside in
// the same Amazon Web Services account as the CA.
//
// Permissions can be granted with the [CreatePermission] action and revoked with the [DeletePermission] action.
//
// # About Permissions
//
// - If the private CA and the certificates it issues reside in the same
// account, you can use CreatePermission to grant permissions for ACM to carry
// out automatic certificate renewals.
//
// - For automatic certificate renewal to succeed, the ACM service principal
// needs permissions to create, retrieve, and list certificates.
//
// - If the private CA and the ACM certificates reside in different accounts,
// then permissions cannot be used to enable automatic renewals. Instead, the ACM
// certificate owner must set up a resource-based policy to enable cross-account
// issuance and renewals. For more information, see [Using a Resource Based Policy with Amazon Web Services Private CA].
//
// [Using a Resource Based Policy with Amazon Web Services Private CA]: https://docs.aws.amazon.com/privateca/latest/userguide/pca-rbp.html
// [CreatePermission]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreatePermission.html
// [DeletePermission]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_DeletePermission.html
func acmpca_ListPermissions(cfg aws.Config, client *acmpca.Client) {
	input := &acmpca.ListPermissionsInput{
		// CertificateAuthorityArn: *string, // Required
	}

	if len(_acmpcaCertificateAuthorityArn) > 0 {
		input.CertificateAuthorityArn = aws.String(_acmpcaCertificateAuthorityArn)
	}
	if len(_acmpcaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _acmpcaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_acmpcaNextToken) > 0 {
		input.NextToken = aws.String(_acmpcaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPermissions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*acmpca.ListPermissionsOutput
	p := acmpca.NewListPermissionsPaginator(client, input)
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

// Lists the tags, if any, that are associated with your private CA or one that
// has been shared with you. Tags are labels that you can use to identify and
// organize your CAs. Each tag consists of a key and an optional value. Call the [TagCertificateAuthority]
// action to add one or more tags to your CA. Call the [UntagCertificateAuthority]action to remove tags.
//
// [TagCertificateAuthority]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_TagCertificateAuthority.html
// [UntagCertificateAuthority]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_UntagCertificateAuthority.html
func acmpca_ListTags(cfg aws.Config, client *acmpca.Client) {
	input := &acmpca.ListTagsInput{
		// CertificateAuthorityArn: *string, // Required
	}

	if len(_acmpcaCertificateAuthorityArn) > 0 {
		input.CertificateAuthorityArn = aws.String(_acmpcaCertificateAuthorityArn)
	}
	if len(_acmpcaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _acmpcaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_acmpcaNextToken) > 0 {
		input.NextToken = aws.String(_acmpcaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTags(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*acmpca.ListTagsOutput
	p := acmpca.NewListTagsPaginator(client, input)
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

// Attaches a resource-based policy to a private CA.
// A policy can also be applied by sharing a private CA through Amazon Web
// Services Resource Access Manager (RAM). For more information, see [Attach a Policy for Cross-Account Access].
//
// The policy can be displayed with [GetPolicy] and removed with [DeletePolicy].
//
// # About Policies
//
// - A policy grants access on a private CA to an Amazon Web Services customer
// account, to Amazon Web Services Organizations, or to an Amazon Web Services
// Organizations unit. Policies are under the control of a CA administrator. For
// more information, see [Using a Resource Based Policy with Amazon Web Services Private CA].
//
// - A policy permits a user of Certificate Manager (ACM) to issue ACM
// certificates signed by a CA in another account.
//
// - For ACM to manage automatic renewal of these certificates, the ACM user
// must configure a Service Linked Role (SLR). The SLR allows the ACM service to
// assume the identity of the user, subject to confirmation against the Amazon Web
// Services Private CA policy. For more information, see [Using a Service Linked Role with ACM].
//
// - Updates made in Amazon Web Services Resource Manager (RAM) are reflected in
// policies. For more information, see [Attach a Policy for Cross-Account Access].
//
// [Using a Resource Based Policy with Amazon Web Services Private CA]: https://docs.aws.amazon.com/privateca/latest/userguide/pca-rbp.html
// [Using a Service Linked Role with ACM]: https://docs.aws.amazon.com/acm/latest/userguide/acm-slr.html
// [Attach a Policy for Cross-Account Access]: https://docs.aws.amazon.com/privateca/latest/userguide/pca-ram.html
// [DeletePolicy]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_DeletePolicy.html
// [GetPolicy]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_GetPolicy.html
func acmpca_PutPolicy(cfg aws.Config, client *acmpca.Client) {
	input := &acmpca.PutPolicyInput{
		// Policy: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_acmpcaPolicy) > 0 {
		input.Policy = aws.String(_acmpcaPolicy)
	}
	if len(_acmpcaResourceArn) > 0 {
		input.ResourceArn = aws.String(_acmpcaResourceArn)
	}

	if resp, err := client.PutPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restores a certificate authority (CA) that is in the DELETED state. You can
// restore a CA during the period that you defined in the
// PermanentDeletionTimeInDays parameter of the [DeleteCertificateAuthority]action. Currently, you can specify
// 7 to 30 days. If you did not specify a PermanentDeletionTimeInDays value, by
// default you can restore the CA at any time in a 30 day period. You can check the
// time remaining in the restoration period of a private CA in the DELETED state
// by calling the [DescribeCertificateAuthority]or [ListCertificateAuthorities] actions. The status of a restored CA is set to its
// pre-deletion status when the RestoreCertificateAuthority action returns. To
// change its status to ACTIVE , call the [UpdateCertificateAuthority] action. If the private CA was in the
// PENDING_CERTIFICATE state at deletion, you must use the [ImportCertificateAuthorityCertificate] action to import a
// certificate authority into the private CA before it can be activated. You cannot
// restore a CA after the restoration period has ended.
//
// [ListCertificateAuthorities]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_ListCertificateAuthorities.html
// [ImportCertificateAuthorityCertificate]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_ImportCertificateAuthorityCertificate.html
// [UpdateCertificateAuthority]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_UpdateCertificateAuthority.html
// [DeleteCertificateAuthority]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_DeleteCertificateAuthority.html
// [DescribeCertificateAuthority]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_DescribeCertificateAuthority.html
func acmpca_RestoreCertificateAuthority(cfg aws.Config, client *acmpca.Client) {
	input := &acmpca.RestoreCertificateAuthorityInput{
		// CertificateAuthorityArn: *string, // Required
	}

	if len(_acmpcaCertificateAuthorityArn) > 0 {
		input.CertificateAuthorityArn = aws.String(_acmpcaCertificateAuthorityArn)
	}

	if resp, err := client.RestoreCertificateAuthority(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Revokes a certificate that was issued inside Amazon Web Services Private CA. If
// you enable a certificate revocation list (CRL) when you create or update your
// private CA, information about the revoked certificates will be included in the
// CRL. Amazon Web Services Private CA writes the CRL to an S3 bucket that you
// specify. A CRL is typically updated approximately 30 minutes after a certificate
// is revoked. If for any reason the CRL update fails, Amazon Web Services Private
// CA attempts makes further attempts every 15 minutes. With Amazon CloudWatch, you
// can create alarms for the metrics CRLGenerated and MisconfiguredCRLBucket . For
// more information, see [Supported CloudWatch Metrics].
//
// Both Amazon Web Services Private CA and the IAM principal must have permission
// to write to the S3 bucket that you specify. If the IAM principal making the call
// does not have permission to write to the bucket, then an exception is thrown.
// For more information, see [Access policies for CRLs in Amazon S3].
//
// Amazon Web Services Private CA also writes revocation information to the audit
// report. For more information, see [CreateCertificateAuthorityAuditReport].
//
// You cannot revoke a root CA self-signed certificate.
//
// [Access policies for CRLs in Amazon S3]: https://docs.aws.amazon.com/privateca/latest/userguide/crl-planning.html#s3-policies
// [CreateCertificateAuthorityAuditReport]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthorityAuditReport.html
// [Supported CloudWatch Metrics]: https://docs.aws.amazon.com/privateca/latest/userguide/PcaCloudWatch.html
func acmpca_RevokeCertificate(cfg aws.Config, client *acmpca.Client) {
	input := &acmpca.RevokeCertificateInput{
		// CertificateAuthorityArn: *string, // Required
		// CertificateSerial: *string, // Required
		// RevocationReason: types.RevocationReason, // Required
	}

	if len(_acmpcaCertificateAuthorityArn) > 0 {
		input.CertificateAuthorityArn = aws.String(_acmpcaCertificateAuthorityArn)
	}
	if len(_acmpcaCertificateSerial) > 0 {
		input.CertificateSerial = aws.String(_acmpcaCertificateSerial)
	}
	if len(_acmpcaRevocationReason) > 0 {
		if err := assignInputField(input, "RevocationReason", _acmpcaRevocationReason); err != nil {
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

// Adds one or more tags to your private CA. Tags are labels that you can use to
// identify and organize your Amazon Web Services resources. Each tag consists of a
// key and an optional value. You specify the private CA on input by its Amazon
// Resource Name (ARN). You specify the tag by using a key-value pair. You can
// apply a tag to just one private CA if you want to identify a specific
// characteristic of that CA, or you can apply the same tag to multiple private CAs
// if you want to filter for a common relationship among those CAs. To remove one
// or more tags, use the [UntagCertificateAuthority]action. Call the [ListTags] action to see what tags are associated
// with your CA.
//
// To attach tags to a private CA during the creation procedure, a CA
// administrator must first associate an inline IAM policy with the
// CreateCertificateAuthority action and explicitly allow tagging. For more
// information, see [Attaching tags to a CA at the time of creation].
//
// [Attaching tags to a CA at the time of creation]: https://docs.aws.amazon.com/privateca/latest/userguide/auth-InlinePolicies.html#policy-tag-ca
// [UntagCertificateAuthority]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_UntagCertificateAuthority.html
// [ListTags]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_ListTags.html
func acmpca_TagCertificateAuthority(cfg aws.Config, client *acmpca.Client) {
	input := &acmpca.TagCertificateAuthorityInput{
		// CertificateAuthorityArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_acmpcaCertificateAuthorityArn) > 0 {
		input.CertificateAuthorityArn = aws.String(_acmpcaCertificateAuthorityArn)
	}
	if len(_acmpcaTags) > 0 {
		if err := assignInputField(input, "Tags", _acmpcaTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagCertificateAuthority(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Remove one or more tags from your private CA. A tag consists of a key-value
// pair. If you do not specify the value portion of the tag when calling this
// action, the tag will be removed regardless of value. If you specify a value, the
// tag is removed only if it is associated with the specified value. To add tags to
// a private CA, use the [TagCertificateAuthority]. Call the [ListTags] action to see what tags are associated with
// your CA.
//
// [TagCertificateAuthority]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_TagCertificateAuthority.html
// [ListTags]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_ListTags.html
func acmpca_UntagCertificateAuthority(cfg aws.Config, client *acmpca.Client) {
	input := &acmpca.UntagCertificateAuthorityInput{
		// CertificateAuthorityArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_acmpcaCertificateAuthorityArn) > 0 {
		input.CertificateAuthorityArn = aws.String(_acmpcaCertificateAuthorityArn)
	}
	if len(_acmpcaTags) > 0 {
		if err := assignInputField(input, "Tags", _acmpcaTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.UntagCertificateAuthority(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the status or configuration of a private certificate authority (CA).
// Your private CA must be in the ACTIVE or DISABLED state before you can update
// it. You can disable a private CA that is in the ACTIVE state or make a CA that
// is in the DISABLED state active again.
//
// Both Amazon Web Services Private CA and the IAM principal must have permission
// to write to the S3 bucket that you specify. If the IAM principal making the call
// does not have permission to write to the bucket, then an exception is thrown.
// For more information, see [Access policies for CRLs in Amazon S3].
//
// [Access policies for CRLs in Amazon S3]: https://docs.aws.amazon.com/privateca/latest/userguide/crl-planning.html#s3-policies
func acmpca_UpdateCertificateAuthority(cfg aws.Config, client *acmpca.Client) {
	input := &acmpca.UpdateCertificateAuthorityInput{
		// CertificateAuthorityArn: *string, // Required
	}

	if len(_acmpcaCertificateAuthorityArn) > 0 {
		input.CertificateAuthorityArn = aws.String(_acmpcaCertificateAuthorityArn)
	}
	if len(_acmpcaRevocationConfiguration) > 0 {
		if err := assignInputField(input, "RevocationConfiguration", _acmpcaRevocationConfiguration); err != nil {
			log.Errorf("invalid --revocation-configuration: %s", err.Error())
			return
		}
	}
	if len(_acmpcaStatus) > 0 {
		if err := assignInputField(input, "Status", _acmpcaStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCertificateAuthority(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_acmpcaCmd)
	_acmpcaCmd.Flags().SortFlags = false

	_acmpcaCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_acmpcaCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_acmpcaCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_acmpcaCmd.Flags().StringVarP(&_acmpcaActions, "actions", "", "", "Actions")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaApiPassthrough, "api-passthrough", "", "", "API Passthrough")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaAuditReportId, "audit-report-id", "", "", "Audit Report ID")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaAuditReportResponseFormat, "audit-report-response-format", "", "", "Audit Report Response Format")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaCertificate, "certificate", "", "", "Certificate")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaCertificateArn, "certificate-arn", "", "", "Certificate ARN")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaCertificateAuthorityArn, "certificate-authority-arn", "", "", "Certificate Authority ARN")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaCertificateAuthorityConfiguration, "certificate-authority-configuration", "", "", "Certificate Authority Configuration")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaCertificateAuthorityType, "certificate-authority-type", "", "", "Certificate Authority Type")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaCertificateChain, "certificate-chain", "", "", "Certificate Chain")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaCertificateSerial, "certificate-serial", "", "", "Certificate Serial")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaCsr, "csr", "", "", "Csr")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaIdempotencyToken, "idempotency-token", "", "", "Idempotency Token")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaKeyStorageSecurityStandard, "key-storage-security-standard", "", "", "Key Storage Security Standard")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaMaxResults, "max-results", "", "", "Max Results")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaNextToken, "next-token", "", "", "Next Token")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaPermanentDeletionTimeInDays, "permanent-deletion-time-in-days", "", "", "Permanent Deletion Time In Days")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaPolicy, "policy", "", "", "Policy")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaPrincipal, "principal", "", "", "Principal")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaResourceArn, "resource-arn", "", "", "Resource ARN")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaResourceOwner, "resource-owner", "", "", "Resource Owner")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaRevocationConfiguration, "revocation-configuration", "", "", "Revocation Configuration")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaRevocationReason, "revocation-reason", "", "", "Revocation Reason")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaS3BucketName, "s3-bucket-name", "", "", "S3 Bucket Name")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaSigningAlgorithm, "signing-algorithm", "", "", "Signing Algorithm")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaSourceAccount, "source-account", "", "", "Source Account")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaStatus, "status", "", "", "Status")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaTags, "tags", "", "", "Tags")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaTemplateArn, "template-arn", "", "", "Template ARN")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaUsageMode, "usage-mode", "", "", "Usage Mode")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaValidity, "validity", "", "", "Validity")
	_acmpcaCmd.Flags().StringVarP(&_acmpcaValidityNotBefore, "validity-not-before", "", "", "Validity Not Before")

	_acmpcaCmd.Flags().BoolVarP(&_acmpcaCreateCertificateAuthority, "create-certificate-authority", "", false, "Create Certificate Authority")
	_acmpcaCmd.Flags().BoolVarP(&_acmpcaCreateCertificateAuthorityAuditReport, "create-certificate-authority-audit-report", "", false, "Create Certificate Authority Audit Report")
	_acmpcaCmd.Flags().BoolVarP(&_acmpcaCreatePermission, "create-permission", "", false, "Create Permission")
	_acmpcaCmd.Flags().BoolVarP(&_acmpcaDeleteCertificateAuthority, "delete-certificate-authority", "", false, "Delete Certificate Authority")
	_acmpcaCmd.Flags().BoolVarP(&_acmpcaDeletePermission, "delete-permission", "", false, "Delete Permission")
	_acmpcaCmd.Flags().BoolVarP(&_acmpcaDeletePolicy, "delete-policy", "", false, "Delete Policy")
	_acmpcaCmd.Flags().BoolVarP(&_acmpcaDescribeCertificateAuthority, "describe-certificate-authority", "", false, "Describe Certificate Authority")
	_acmpcaCmd.Flags().BoolVarP(&_acmpcaDescribeCertificateAuthorityAuditReport, "describe-certificate-authority-audit-report", "", false, "Describe Certificate Authority Audit Report")
	_acmpcaCmd.Flags().BoolVarP(&_acmpcaGetCertificate, "get-certificate", "", false, "Get Certificate")
	_acmpcaCmd.Flags().BoolVarP(&_acmpcaGetCertificateAuthorityCertificate, "get-certificate-authority-certificate", "", false, "Get Certificate Authority Certificate")
	_acmpcaCmd.Flags().BoolVarP(&_acmpcaGetCertificateAuthorityCsr, "get-certificate-authority-csr", "", false, "Get Certificate Authority Csr")
	_acmpcaCmd.Flags().BoolVarP(&_acmpcaGetPolicy, "get-policy", "", false, "Get Policy")
	_acmpcaCmd.Flags().BoolVarP(&_acmpcaImportCertificateAuthorityCertificate, "import-certificate-authority-certificate", "", false, "Import Certificate Authority Certificate")
	_acmpcaCmd.Flags().BoolVarP(&_acmpcaIssueCertificate, "issue-certificate", "", false, "Issue Certificate")
	_acmpcaCmd.Flags().BoolVarP(&_acmpcaListCertificateAuthorities, "list-certificate-authorities", "", false, "List Certificate Authorities")
	_acmpcaCmd.Flags().BoolVarP(&_acmpcaListPermissions, "list-permissions", "", false, "List Permissions")
	_acmpcaCmd.Flags().BoolVarP(&_acmpcaListTags, "list-tags", "", false, "List Tags")
	_acmpcaCmd.Flags().BoolVarP(&_acmpcaPutPolicy, "put-policy", "", false, "Put Policy")
	_acmpcaCmd.Flags().BoolVarP(&_acmpcaRestoreCertificateAuthority, "restore-certificate-authority", "", false, "Restore Certificate Authority")
	_acmpcaCmd.Flags().BoolVarP(&_acmpcaRevokeCertificate, "revoke-certificate", "", false, "Revoke Certificate")
	_acmpcaCmd.Flags().BoolVarP(&_acmpcaTagCertificateAuthority, "tag-certificate-authority", "", false, "Tag Certificate Authority")
	_acmpcaCmd.Flags().BoolVarP(&_acmpcaUntagCertificateAuthority, "untag-certificate-authority", "", false, "Untag Certificate Authority")
	_acmpcaCmd.Flags().BoolVarP(&_acmpcaUpdateCertificateAuthority, "update-certificate-authority", "", false, "Update Certificate Authority")

}
