package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/signerdata"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// signerdataCmd represents the signerdata command
var _signerdataCmd = &cobra.Command{
	Use:   "signerdata",
	Short: "AWS signerdata CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := signerdata.NewFromConfig(cfg)
		if _signerdataGetRevocationStatus {
			signerdata_GetRevocationStatus(cfg, client)
			return
		}

	},
}

var (
	_signerdataGetRevocationStatus bool

	_signerdataCertificateHashes  []string
	_signerdataJobArn             string
	_signerdataPlatformId         string
	_signerdataProfileVersionArn  string
	_signerdataSignatureTimestamp string
)

// Retrieves the revocation status for a signed artifact by checking if the
// signing profile, job, or certificate has been revoked.
func signerdata_GetRevocationStatus(cfg aws.Config, client *signerdata.Client) {
	input := &signerdata.GetRevocationStatusInput{
		// CertificateHashes: []string, // Required
		// JobArn: *string, // Required
		// PlatformId: *string, // Required
		// ProfileVersionArn: *string, // Required
		// SignatureTimestamp: *time.Time, // Required
	}

	if len(_signerdataCertificateHashes) > 0 {
		input.CertificateHashes = append([]string(nil), _signerdataCertificateHashes...)
	}
	if len(_signerdataJobArn) > 0 {
		input.JobArn = aws.String(_signerdataJobArn)
	}
	if len(_signerdataPlatformId) > 0 {
		input.PlatformId = aws.String(_signerdataPlatformId)
	}
	if len(_signerdataProfileVersionArn) > 0 {
		input.ProfileVersionArn = aws.String(_signerdataProfileVersionArn)
	}
	if len(_signerdataSignatureTimestamp) > 0 {
		if err := assignInputField(input, "SignatureTimestamp", _signerdataSignatureTimestamp); err != nil {
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

func init() {
	_rootCmd.AddCommand(_signerdataCmd)
	_signerdataCmd.Flags().SortFlags = false

	_signerdataCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_signerdataCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_signerdataCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_signerdataCmd.Flags().StringSliceVarP(&_signerdataCertificateHashes, "certificate-hashes", "", nil, "Certificate Hashes")
	_signerdataCmd.Flags().StringVarP(&_signerdataJobArn, "job-arn", "", "", "Job ARN")
	_signerdataCmd.Flags().StringVarP(&_signerdataPlatformId, "platform-id", "", "", "Platform ID")
	_signerdataCmd.Flags().StringVarP(&_signerdataProfileVersionArn, "profile-version-arn", "", "", "Profile Version ARN")
	_signerdataCmd.Flags().StringVarP(&_signerdataSignatureTimestamp, "signature-timestamp", "", "", "Signature Timestamp")

	_signerdataCmd.Flags().BoolVarP(&_signerdataGetRevocationStatus, "get-revocation-status", "", false, "Get Revocation Status")

}
