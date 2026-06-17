package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspectorscan"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// inspectorscanCmd represents the inspectorscan command
var _inspectorscanCmd = &cobra.Command{
	Use:   "inspectorscan",
	Short: "AWS inspectorscan CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := inspectorscan.NewFromConfig(cfg)
		if _inspectorscanScanSbom {
			inspectorscan_ScanSbom(cfg, client)
			return
		}

	},
}

var (
	_inspectorscanScanSbom bool

	_inspectorscanOutputFormat string
	_inspectorscanSbom         string
)

// Scans a provided CycloneDX 1.5 SBOM and reports on any vulnerabilities
// discovered in that SBOM. You can generate compatible SBOMs for your resources
// using the [Amazon Inspector SBOM generator].
//
// The output of this action reports NVD and CVSS scores when NVD and CVSS scores
// are available. Because the output reports both scores, you might notice a
// discrepency between them. However, you can triage the severity of either score
// depending on the vendor of your choosing.
//
// [Amazon Inspector SBOM generator]: https://docs.aws.amazon.com/inspector/latest/user/sbom-generator.html
func inspectorscan_ScanSbom(cfg aws.Config, client *inspectorscan.Client) {
	input := &inspectorscan.ScanSbomInput{
		// Sbom: document.Interface, // Required
	}

	if len(_inspectorscanSbom) > 0 {
		if err := assignInputField(input, "Sbom", _inspectorscanSbom); err != nil {
			log.Errorf("invalid --sbom: %s", err.Error())
			return
		}
	}
	if len(_inspectorscanOutputFormat) > 0 {
		if err := assignInputField(input, "OutputFormat", _inspectorscanOutputFormat); err != nil {
			log.Errorf("invalid --output-format: %s", err.Error())
			return
		}
	}

	if resp, err := client.ScanSbom(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_inspectorscanCmd)
	_inspectorscanCmd.Flags().SortFlags = false

	_inspectorscanCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_inspectorscanCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_inspectorscanCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_inspectorscanCmd.Flags().StringVarP(&_inspectorscanOutputFormat, "output-format", "", "", "Output Format")
	_inspectorscanCmd.Flags().StringVarP(&_inspectorscanSbom, "sbom", "", "", "Sbom")

	_inspectorscanCmd.Flags().BoolVarP(&_inspectorscanScanSbom, "scan-sbom", "", false, "Scan Sbom")

}
