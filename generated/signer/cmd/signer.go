package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/signer"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// signerCmd represents the signer command
var _signerCmd = &cobra.Command{
	Use:   "signer",
	Short: "AWS signer CLI",
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
		client := signer.NewFromConfig(cfg)
		if _signerAddProfilePermission {
			signer_AddProfilePermission(cfg, client)
			return
		}
		if _signerCancelSigningProfile {
			signer_CancelSigningProfile(cfg, client)
			return
		}
		if _signerDescribeSigningJob {
			signer_DescribeSigningJob(cfg, client)
			return
		}
		if _signerGetRevocationStatus {
			signer_GetRevocationStatus(cfg, client)
			return
		}
		if _signerGetSigningPlatform {
			signer_GetSigningPlatform(cfg, client)
			return
		}
		if _signerGetSigningProfile {
			signer_GetSigningProfile(cfg, client)
			return
		}
		if _signerListProfilePermissions {
			signer_ListProfilePermissions(cfg, client)
			return
		}
		if _signerListSigningJobs {
			signer_ListSigningJobs(cfg, client)
			return
		}
		if _signerListSigningPlatforms {
			signer_ListSigningPlatforms(cfg, client)
			return
		}
		if _signerListSigningProfiles {
			signer_ListSigningProfiles(cfg, client)
			return
		}
		if _signerListTagsForResource {
			signer_ListTagsForResource(cfg, client)
			return
		}
		if _signerPutSigningProfile {
			signer_PutSigningProfile(cfg, client)
			return
		}
		if _signerRemoveProfilePermission {
			signer_RemoveProfilePermission(cfg, client)
			return
		}
		if _signerRevokeSignature {
			signer_RevokeSignature(cfg, client)
			return
		}
		if _signerRevokeSigningProfile {
			signer_RevokeSigningProfile(cfg, client)
			return
		}
		if _signerSignPayload {
			signer_SignPayload(cfg, client)
			return
		}
		if _signerStartSigningJob {
			signer_StartSigningJob(cfg, client)
			return
		}
		if _signerTagResource {
			signer_TagResource(cfg, client)
			return
		}
		if _signerUntagResource {
			signer_UntagResource(cfg, client)
			return
		}

	},
}

var (
	_signerAddProfilePermission    bool
	_signerCancelSigningProfile    bool
	_signerDescribeSigningJob      bool
	_signerGetRevocationStatus     bool
	_signerGetSigningPlatform      bool
	_signerGetSigningProfile       bool
	_signerListProfilePermissions  bool
	_signerListSigningJobs         bool
	_signerListSigningPlatforms    bool
	_signerListSigningProfiles     bool
	_signerListTagsForResource     bool
	_signerPutSigningProfile       bool
	_signerRemoveProfilePermission bool
	_signerRevokeSignature         bool
	_signerRevokeSigningProfile    bool
	_signerSignPayload             bool
	_signerStartSigningJob         bool
	_signerTagResource             bool
	_signerUntagResource           bool

	_signerAction                  string
	_signerCategory                string
	_signerCertificateHashes       []string
	_signerClientRequestToken      string
	_signerDestination             string
	_signerEffectiveTime           string
	_signerIncludeCanceled         string
	_signerIsRevoked               string
	_signerJobArn                  string
	_signerJobId                   string
	_signerJobInvoker              string
	_signerJobOwner                string
	_signerMaxResults              string
	_signerNextToken               string
	_signerOverrides               string
	_signerPartner                 string
	_signerPayload                 string
	_signerPayloadFormat           string
	_signerPlatformId              string
	_signerPrincipal               string
	_signerProfileName             string
	_signerProfileOwner            string
	_signerProfileVersion          string
	_signerProfileVersionArn       string
	_signerReason                  string
	_signerRequestedBy             string
	_signerResourceArn             string
	_signerRevisionId              string
	_signerSignatureExpiresAfter   string
	_signerSignatureExpiresBefore  string
	_signerSignatureTimestamp      string
	_signerSignatureValidityPeriod string
	_signerSigningMaterial         string
	_signerSigningParameters       string
	_signerSource                  string
	_signerStatementId             string
	_signerStatus                  string
	_signerStatuses                string
	_signerTagKeys                 []string
	_signerTags                    string
	_signerTarget                  string
)

// Adds cross-account permissions to a signing profile.
func signer_AddProfilePermission(cfg aws.Config, client *signer.Client) {
	input := &signer.AddProfilePermissionInput{
		// Action: *string, // Required
		// Principal: *string, // Required
		// ProfileName: *string, // Required
		// StatementId: *string, // Required
	}

	if len(_signerAction) > 0 {
		input.Action = aws.String(_signerAction)
	}
	if len(_signerPrincipal) > 0 {
		input.Principal = aws.String(_signerPrincipal)
	}
	if len(_signerProfileName) > 0 {
		input.ProfileName = aws.String(_signerProfileName)
	}
	if len(_signerStatementId) > 0 {
		input.StatementId = aws.String(_signerStatementId)
	}
	if len(_signerProfileVersion) > 0 {
		input.ProfileVersion = aws.String(_signerProfileVersion)
	}
	if len(_signerRevisionId) > 0 {
		input.RevisionId = aws.String(_signerRevisionId)
	}

	if resp, err := client.AddProfilePermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the state of an ACTIVE signing profile to CANCELED . A canceled profile
// is still viewable with the ListSigningProfiles operation, but it cannot perform
// new signing jobs. See [Data Retention]for more information on scheduled deletion of a canceled
// signing profile.
//
// [Data Retention]: https://docs.aws.amazon.com/signer/latest/developerguide/retention.html
func signer_CancelSigningProfile(cfg aws.Config, client *signer.Client) {
	input := &signer.CancelSigningProfileInput{
		// ProfileName: *string, // Required
	}

	if len(_signerProfileName) > 0 {
		input.ProfileName = aws.String(_signerProfileName)
	}

	if resp, err := client.CancelSigningProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specific code signing job. You specify the job by
// using the jobId value that is returned by the StartSigningJob operation.
func signer_DescribeSigningJob(cfg aws.Config, client *signer.Client) {
	input := &signer.DescribeSigningJobInput{
		// JobId: *string, // Required
	}

	if len(_signerJobId) > 0 {
		input.JobId = aws.String(_signerJobId)
	}

	if resp, err := client.DescribeSigningJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the revocation status of one or more of the signing profile, signing
// job, and signing certificate.
func signer_GetRevocationStatus(cfg aws.Config, client *signer.Client) {
	input := &signer.GetRevocationStatusInput{
		// CertificateHashes: []string, // Required
		// JobArn: *string, // Required
		// PlatformId: *string, // Required
		// ProfileVersionArn: *string, // Required
		// SignatureTimestamp: *time.Time, // Required
	}

	if len(_signerCertificateHashes) > 0 {
		input.CertificateHashes = append([]string(nil), _signerCertificateHashes...)
	}
	if len(_signerJobArn) > 0 {
		input.JobArn = aws.String(_signerJobArn)
	}
	if len(_signerPlatformId) > 0 {
		input.PlatformId = aws.String(_signerPlatformId)
	}
	if len(_signerProfileVersionArn) > 0 {
		input.ProfileVersionArn = aws.String(_signerProfileVersionArn)
	}
	if len(_signerSignatureTimestamp) > 0 {
		if err := assignInputField(input, "SignatureTimestamp", _signerSignatureTimestamp); err != nil {
			log.Errorf("invalid --signature-timestamp: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetRevocationStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information on a specific signing platform.
func signer_GetSigningPlatform(cfg aws.Config, client *signer.Client) {
	input := &signer.GetSigningPlatformInput{
		// PlatformId: *string, // Required
	}

	if len(_signerPlatformId) > 0 {
		input.PlatformId = aws.String(_signerPlatformId)
	}

	if resp, err := client.GetSigningPlatform(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information on a specific signing profile.
func signer_GetSigningProfile(cfg aws.Config, client *signer.Client) {
	input := &signer.GetSigningProfileInput{
		// ProfileName: *string, // Required
	}

	if len(_signerProfileName) > 0 {
		input.ProfileName = aws.String(_signerProfileName)
	}
	if len(_signerProfileOwner) > 0 {
		input.ProfileOwner = aws.String(_signerProfileOwner)
	}

	if resp, err := client.GetSigningProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the cross-account permissions associated with a signing profile.
func signer_ListProfilePermissions(cfg aws.Config, client *signer.Client) {
	input := &signer.ListProfilePermissionsInput{
		// ProfileName: *string, // Required
	}

	if len(_signerProfileName) > 0 {
		input.ProfileName = aws.String(_signerProfileName)
	}
	if len(_signerNextToken) > 0 {
		input.NextToken = aws.String(_signerNextToken)
	}

	if resp, err := client.ListProfilePermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all your signing jobs. You can use the maxResults parameter to limit the
// number of signing jobs that are returned in the response. If additional jobs
// remain to be listed, AWS Signer returns a nextToken value. Use this value in
// subsequent calls to ListSigningJobs to fetch the remaining values. You can
// continue calling ListSigningJobs with your maxResults parameter and with new
// values that Signer returns in the nextToken parameter until all of your signing
// jobs have been returned.
func signer_ListSigningJobs(cfg aws.Config, client *signer.Client) {
	input := &signer.ListSigningJobsInput{}

	if len(_signerIsRevoked) > 0 {
		if err := assignInputField(input, "IsRevoked", _signerIsRevoked); err != nil {
			log.Errorf("invalid --is-revoked: %s", err.Error())
			return
		}
	}
	if len(_signerJobInvoker) > 0 {
		input.JobInvoker = aws.String(_signerJobInvoker)
	}
	if len(_signerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _signerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_signerNextToken) > 0 {
		input.NextToken = aws.String(_signerNextToken)
	}
	if len(_signerPlatformId) > 0 {
		input.PlatformId = aws.String(_signerPlatformId)
	}
	if len(_signerRequestedBy) > 0 {
		input.RequestedBy = aws.String(_signerRequestedBy)
	}
	if len(_signerSignatureExpiresAfter) > 0 {
		if err := assignInputField(input, "SignatureExpiresAfter", _signerSignatureExpiresAfter); err != nil {
			log.Errorf("invalid --signature-expires-after: %s", err.Error())
			return
		}
	}
	if len(_signerSignatureExpiresBefore) > 0 {
		if err := assignInputField(input, "SignatureExpiresBefore", _signerSignatureExpiresBefore); err != nil {
			log.Errorf("invalid --signature-expires-before: %s", err.Error())
			return
		}
	}
	if len(_signerStatus) > 0 {
		if err := assignInputField(input, "Status", _signerStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSigningJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*signer.ListSigningJobsOutput
	p := signer.NewListSigningJobsPaginator(client, input)
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

// Lists all signing platforms available in AWS Signer that match the request
// parameters. If additional jobs remain to be listed, Signer returns a nextToken
// value. Use this value in subsequent calls to ListSigningJobs to fetch the
// remaining values. You can continue calling ListSigningJobs with your maxResults
// parameter and with new values that Signer returns in the nextToken parameter
// until all of your signing jobs have been returned.
func signer_ListSigningPlatforms(cfg aws.Config, client *signer.Client) {
	input := &signer.ListSigningPlatformsInput{}

	if len(_signerCategory) > 0 {
		input.Category = aws.String(_signerCategory)
	}
	if len(_signerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _signerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_signerNextToken) > 0 {
		input.NextToken = aws.String(_signerNextToken)
	}
	if len(_signerPartner) > 0 {
		input.Partner = aws.String(_signerPartner)
	}
	if len(_signerTarget) > 0 {
		input.Target = aws.String(_signerTarget)
	}

	if disablePaginator() {
		if resp, err := client.ListSigningPlatforms(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*signer.ListSigningPlatformsOutput
	p := signer.NewListSigningPlatformsPaginator(client, input)
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

// Lists all available signing profiles in your AWS account. Returns only profiles
// with an ACTIVE status unless the includeCanceled request field is set to true .
// If additional jobs remain to be listed, AWS Signer returns a nextToken value.
// Use this value in subsequent calls to ListSigningJobs to fetch the remaining
// values. You can continue calling ListSigningJobs with your maxResults parameter
// and with new values that Signer returns in the nextToken parameter until all of
// your signing jobs have been returned.
func signer_ListSigningProfiles(cfg aws.Config, client *signer.Client) {
	input := &signer.ListSigningProfilesInput{}

	if len(_signerIncludeCanceled) > 0 {
		if err := assignInputField(input, "IncludeCanceled", _signerIncludeCanceled); err != nil {
			log.Errorf("invalid --include-canceled: %s", err.Error())
			return
		}
	}
	if len(_signerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _signerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_signerNextToken) > 0 {
		input.NextToken = aws.String(_signerNextToken)
	}
	if len(_signerPlatformId) > 0 {
		input.PlatformId = aws.String(_signerPlatformId)
	}
	if len(_signerStatuses) > 0 {
		if err := assignInputField(input, "Statuses", _signerStatuses); err != nil {
			log.Errorf("invalid --statuses: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSigningProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*signer.ListSigningProfilesOutput
	p := signer.NewListSigningProfilesPaginator(client, input)
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

// Returns a list of the tags associated with a signing profile resource.
func signer_ListTagsForResource(cfg aws.Config, client *signer.Client) {
	input := &signer.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_signerResourceArn) > 0 {
		input.ResourceArn = aws.String(_signerResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a signing profile. A signing profile is a code-signing template that
// can be used to carry out a pre-defined signing job.
func signer_PutSigningProfile(cfg aws.Config, client *signer.Client) {
	input := &signer.PutSigningProfileInput{
		// PlatformId: *string, // Required
		// ProfileName: *string, // Required
	}

	if len(_signerPlatformId) > 0 {
		input.PlatformId = aws.String(_signerPlatformId)
	}
	if len(_signerProfileName) > 0 {
		input.ProfileName = aws.String(_signerProfileName)
	}
	if len(_signerOverrides) > 0 {
		if err := assignInputField(input, "Overrides", _signerOverrides); err != nil {
			log.Errorf("invalid --overrides: %s", err.Error())
			return
		}
	}
	if len(_signerSignatureValidityPeriod) > 0 {
		if err := assignInputField(input, "SignatureValidityPeriod", _signerSignatureValidityPeriod); err != nil {
			log.Errorf("invalid --signature-validity-period: %s", err.Error())
			return
		}
	}
	if len(_signerSigningMaterial) > 0 {
		if err := assignInputField(input, "SigningMaterial", _signerSigningMaterial); err != nil {
			log.Errorf("invalid --signing-material: %s", err.Error())
			return
		}
	}
	if len(_signerSigningParameters) > 0 {
		if err := assignInputField(input, "SigningParameters", _signerSigningParameters); err != nil {
			log.Errorf("invalid --signing-parameters: %s", err.Error())
			return
		}
	}
	if len(_signerTags) > 0 {
		if err := assignInputField(input, "Tags", _signerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutSigningProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes cross-account permissions from a signing profile.
func signer_RemoveProfilePermission(cfg aws.Config, client *signer.Client) {
	input := &signer.RemoveProfilePermissionInput{
		// ProfileName: *string, // Required
		// RevisionId: *string, // Required
		// StatementId: *string, // Required
	}

	if len(_signerProfileName) > 0 {
		input.ProfileName = aws.String(_signerProfileName)
	}
	if len(_signerRevisionId) > 0 {
		input.RevisionId = aws.String(_signerRevisionId)
	}
	if len(_signerStatementId) > 0 {
		input.StatementId = aws.String(_signerStatementId)
	}

	if resp, err := client.RemoveProfilePermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the state of a signing job to REVOKED . This indicates that the
// signature is no longer valid.
func signer_RevokeSignature(cfg aws.Config, client *signer.Client) {
	input := &signer.RevokeSignatureInput{
		// JobId: *string, // Required
		// Reason: *string, // Required
	}

	if len(_signerJobId) > 0 {
		input.JobId = aws.String(_signerJobId)
	}
	if len(_signerReason) > 0 {
		input.Reason = aws.String(_signerReason)
	}
	if len(_signerJobOwner) > 0 {
		input.JobOwner = aws.String(_signerJobOwner)
	}

	if resp, err := client.RevokeSignature(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the state of a signing profile to REVOKED . This indicates that
// signatures generated using the signing profile after an effective start date are
// no longer valid. A revoked profile is still viewable with the
// ListSigningProfiles operation, but it cannot perform new signing jobs. See [Data Retention] for
// more information on scheduled deletion of a revoked signing profile.
//
// [Data Retention]: https://docs.aws.amazon.com/signer/latest/developerguide/retention.html
func signer_RevokeSigningProfile(cfg aws.Config, client *signer.Client) {
	input := &signer.RevokeSigningProfileInput{
		// EffectiveTime: *time.Time, // Required
		// ProfileName: *string, // Required
		// ProfileVersion: *string, // Required
		// Reason: *string, // Required
	}

	if len(_signerEffectiveTime) > 0 {
		if err := assignInputField(input, "EffectiveTime", _signerEffectiveTime); err != nil {
			log.Errorf("invalid --effective-time: %s", err.Error())
			return
		}
	}
	if len(_signerProfileName) > 0 {
		input.ProfileName = aws.String(_signerProfileName)
	}
	if len(_signerProfileVersion) > 0 {
		input.ProfileVersion = aws.String(_signerProfileVersion)
	}
	if len(_signerReason) > 0 {
		input.Reason = aws.String(_signerReason)
	}

	if resp, err := client.RevokeSigningProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Signs a binary payload and returns a signature envelope.
func signer_SignPayload(cfg aws.Config, client *signer.Client) {
	input := &signer.SignPayloadInput{
		// Payload: []byte, // Required
		// PayloadFormat: *string, // Required
		// ProfileName: *string, // Required
	}

	if len(_signerPayload) > 0 {
		if err := assignInputField(input, "Payload", _signerPayload); err != nil {
			log.Errorf("invalid --payload: %s", err.Error())
			return
		}
	}
	if len(_signerPayloadFormat) > 0 {
		input.PayloadFormat = aws.String(_signerPayloadFormat)
	}
	if len(_signerProfileName) > 0 {
		input.ProfileName = aws.String(_signerProfileName)
	}
	if len(_signerProfileOwner) > 0 {
		input.ProfileOwner = aws.String(_signerProfileOwner)
	}

	if resp, err := client.SignPayload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates a signing job to be performed on the code provided. Signing jobs are
// viewable by the ListSigningJobs operation. Note the following requirements:
//
// - You must create an Amazon S3 source bucket. For more information, see [Creating a Bucket]in
// the Amazon S3 Getting Started Guide.
//
// - Your S3 source bucket must be version enabled.
//
// - You must create an S3 destination bucket. AWS Signer uses your S3
// destination bucket to write your signed code.
//
// - You specify the name of the source and destination buckets when calling the
// StartSigningJob operation.
//
// - You must ensure the S3 buckets are from the same Region as the signing
// profile. Cross-Region signing isn't supported.
//
// - You must also specify a request token that identifies your request to
// Signer.
//
// You can call the DescribeSigningJob and the ListSigningJobs actions after you call StartSigningJob .
//
// For a Java example that shows how to use this action, see [StartSigningJob].
//
// [Creating a Bucket]: http://docs.aws.amazon.com/AmazonS3/latest/gsg/CreatingABucket.html
// [StartSigningJob]: https://docs.aws.amazon.com/signer/latest/developerguide/api-startsigningjob.html
func signer_StartSigningJob(cfg aws.Config, client *signer.Client) {
	input := &signer.StartSigningJobInput{
		// ClientRequestToken: *string, // Required
		// Destination: *types.Destination, // Required
		// ProfileName: *string, // Required
		// Source: *types.Source, // Required
	}

	if len(_signerClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_signerClientRequestToken)
	}
	if len(_signerDestination) > 0 {
		if err := assignInputField(input, "Destination", _signerDestination); err != nil {
			log.Errorf("invalid --destination: %s", err.Error())
			return
		}
	}
	if len(_signerProfileName) > 0 {
		input.ProfileName = aws.String(_signerProfileName)
	}
	if len(_signerSource) > 0 {
		if err := assignInputField(input, "Source", _signerSource); err != nil {
			log.Errorf("invalid --source: %s", err.Error())
			return
		}
	}
	if len(_signerProfileOwner) > 0 {
		input.ProfileOwner = aws.String(_signerProfileOwner)
	}

	if resp, err := client.StartSigningJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more tags to a signing profile. Tags are labels that you can use to
// identify and organize your AWS resources. Each tag consists of a key and an
// optional value. To specify the signing profile, use its Amazon Resource Name
// (ARN). To specify the tag, use a key-value pair.
func signer_TagResource(cfg aws.Config, client *signer.Client) {
	input := &signer.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_signerResourceArn) > 0 {
		input.ResourceArn = aws.String(_signerResourceArn)
	}
	if len(_signerTags) > 0 {
		if err := assignInputField(input, "Tags", _signerTags); err != nil {
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

// Removes one or more tags from a signing profile. To remove the tags, specify a
// list of tag keys.
func signer_UntagResource(cfg aws.Config, client *signer.Client) {
	input := &signer.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_signerResourceArn) > 0 {
		input.ResourceArn = aws.String(_signerResourceArn)
	}
	if len(_signerTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _signerTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_signerCmd)
	_signerCmd.Flags().SortFlags = false

	_signerCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_signerCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_signerCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_signerCmd.Flags().StringVarP(&_signerAction, "action", "", "", "Action")
	_signerCmd.Flags().StringVarP(&_signerCategory, "category", "", "", "Category")
	_signerCmd.Flags().StringSliceVarP(&_signerCertificateHashes, "certificate-hashes", "", nil, "Certificate Hashes")
	_signerCmd.Flags().StringVarP(&_signerClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_signerCmd.Flags().StringVarP(&_signerDestination, "destination", "", "", "Destination")
	_signerCmd.Flags().StringVarP(&_signerEffectiveTime, "effective-time", "", "", "Effective Time")
	_signerCmd.Flags().StringVarP(&_signerIncludeCanceled, "include-canceled", "", "", "Include Canceled")
	_signerCmd.Flags().StringVarP(&_signerIsRevoked, "is-revoked", "", "", "Is Revoked")
	_signerCmd.Flags().StringVarP(&_signerJobArn, "job-arn", "", "", "Job ARN")
	_signerCmd.Flags().StringVarP(&_signerJobId, "job-id", "", "", "Job ID")
	_signerCmd.Flags().StringVarP(&_signerJobInvoker, "job-invoker", "", "", "Job Invoker")
	_signerCmd.Flags().StringVarP(&_signerJobOwner, "job-owner", "", "", "Job Owner")
	_signerCmd.Flags().StringVarP(&_signerMaxResults, "max-results", "", "", "Max Results")
	_signerCmd.Flags().StringVarP(&_signerNextToken, "next-token", "", "", "Next Token")
	_signerCmd.Flags().StringVarP(&_signerOverrides, "overrides", "", "", "Overrides")
	_signerCmd.Flags().StringVarP(&_signerPartner, "partner", "", "", "Partner")
	_signerCmd.Flags().StringVarP(&_signerPayload, "payload", "", "", "Payload")
	_signerCmd.Flags().StringVarP(&_signerPayloadFormat, "payload-format", "", "", "Payload Format")
	_signerCmd.Flags().StringVarP(&_signerPlatformId, "platform-id", "", "", "Platform ID")
	_signerCmd.Flags().StringVarP(&_signerPrincipal, "principal", "", "", "Principal")
	_signerCmd.Flags().StringVarP(&_signerProfileName, "profile-name", "", "", "Profile Name")
	_signerCmd.Flags().StringVarP(&_signerProfileOwner, "profile-owner", "", "", "Profile Owner")
	_signerCmd.Flags().StringVarP(&_signerProfileVersion, "profile-version", "", "", "Profile Version")
	_signerCmd.Flags().StringVarP(&_signerProfileVersionArn, "profile-version-arn", "", "", "Profile Version ARN")
	_signerCmd.Flags().StringVarP(&_signerReason, "reason", "", "", "Reason")
	_signerCmd.Flags().StringVarP(&_signerRequestedBy, "requested-by", "", "", "Requested By")
	_signerCmd.Flags().StringVarP(&_signerResourceArn, "resource-arn", "", "", "Resource ARN")
	_signerCmd.Flags().StringVarP(&_signerRevisionId, "revision-id", "", "", "Revision ID")
	_signerCmd.Flags().StringVarP(&_signerSignatureExpiresAfter, "signature-expires-after", "", "", "Signature Expires After")
	_signerCmd.Flags().StringVarP(&_signerSignatureExpiresBefore, "signature-expires-before", "", "", "Signature Expires Before")
	_signerCmd.Flags().StringVarP(&_signerSignatureTimestamp, "signature-timestamp", "", "", "Signature Timestamp")
	_signerCmd.Flags().StringVarP(&_signerSignatureValidityPeriod, "signature-validity-period", "", "", "Signature Validity Period")
	_signerCmd.Flags().StringVarP(&_signerSigningMaterial, "signing-material", "", "", "Signing Material")
	_signerCmd.Flags().StringVarP(&_signerSigningParameters, "signing-parameters", "", "", "Signing Parameters")
	_signerCmd.Flags().StringVarP(&_signerSource, "source", "", "", "Source")
	_signerCmd.Flags().StringVarP(&_signerStatementId, "statement-id", "", "", "Statement ID")
	_signerCmd.Flags().StringVarP(&_signerStatus, "status", "", "", "Status")
	_signerCmd.Flags().StringVarP(&_signerStatuses, "statuses", "", "", "Statuses")
	_signerCmd.Flags().StringSliceVarP(&_signerTagKeys, "tag-keys", "", nil, "Tag Keys")
	_signerCmd.Flags().StringVarP(&_signerTags, "tags", "", "", "Tags")
	_signerCmd.Flags().StringVarP(&_signerTarget, "target", "", "", "Target")

	_signerCmd.Flags().BoolVarP(&_signerAddProfilePermission, "add-profile-permission", "", false, "Add Profile Permission")
	_signerCmd.Flags().BoolVarP(&_signerCancelSigningProfile, "cancel-signing-profile", "", false, "Cancel Signing Profile")
	_signerCmd.Flags().BoolVarP(&_signerDescribeSigningJob, "describe-signing-job", "", false, "Describe Signing Job")
	_signerCmd.Flags().BoolVarP(&_signerGetRevocationStatus, "get-revocation-status", "", false, "Get Revocation Status")
	_signerCmd.Flags().BoolVarP(&_signerGetSigningPlatform, "get-signing-platform", "", false, "Get Signing Platform")
	_signerCmd.Flags().BoolVarP(&_signerGetSigningProfile, "get-signing-profile", "", false, "Get Signing Profile")
	_signerCmd.Flags().BoolVarP(&_signerListProfilePermissions, "list-profile-permissions", "", false, "List Profile Permissions")
	_signerCmd.Flags().BoolVarP(&_signerListSigningJobs, "list-signing-jobs", "", false, "List Signing Jobs")
	_signerCmd.Flags().BoolVarP(&_signerListSigningPlatforms, "list-signing-platforms", "", false, "List Signing Platforms")
	_signerCmd.Flags().BoolVarP(&_signerListSigningProfiles, "list-signing-profiles", "", false, "List Signing Profiles")
	_signerCmd.Flags().BoolVarP(&_signerListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_signerCmd.Flags().BoolVarP(&_signerPutSigningProfile, "put-signing-profile", "", false, "Put Signing Profile")
	_signerCmd.Flags().BoolVarP(&_signerRemoveProfilePermission, "remove-profile-permission", "", false, "Remove Profile Permission")
	_signerCmd.Flags().BoolVarP(&_signerRevokeSignature, "revoke-signature", "", false, "Revoke Signature")
	_signerCmd.Flags().BoolVarP(&_signerRevokeSigningProfile, "revoke-signing-profile", "", false, "Revoke Signing Profile")
	_signerCmd.Flags().BoolVarP(&_signerSignPayload, "sign-payload", "", false, "Sign Payload")
	_signerCmd.Flags().BoolVarP(&_signerStartSigningJob, "start-signing-job", "", false, "Start Signing Job")
	_signerCmd.Flags().BoolVarP(&_signerTagResource, "tag-resource", "", false, "Tag Resource")
	_signerCmd.Flags().BoolVarP(&_signerUntagResource, "untag-resource", "", false, "Untag Resource")

}
