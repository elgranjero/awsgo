package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eksauth"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// eksauthCmd represents the eksauth command
var _eksauthCmd = &cobra.Command{
	Use:   "eksauth",
	Short: "AWS eksauth CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := eksauth.NewFromConfig(cfg)
		if _eksauthAssumeRoleForPodIdentity {
			eksauth_AssumeRoleForPodIdentity(cfg, client)
			return
		}

	},
}

var (
	_eksauthAssumeRoleForPodIdentity bool

	_eksauthClusterName string
	_eksauthToken       string
)

// The Amazon EKS Auth API and the AssumeRoleForPodIdentity action are only used
// by the EKS Pod Identity Agent.
//
// We recommend that applications use the Amazon Web Services SDKs to connect to
// Amazon Web Services services; if credentials from an EKS Pod Identity
// association are available in the pod, the latest versions of the SDKs use them
// automatically.
func eksauth_AssumeRoleForPodIdentity(cfg aws.Config, client *eksauth.Client) {
	input := &eksauth.AssumeRoleForPodIdentityInput{
		// ClusterName: *string, // Required
		// Token: *string, // Required
	}

	if len(_eksauthClusterName) > 0 {
		input.ClusterName = aws.String(_eksauthClusterName)
	}
	if len(_eksauthToken) > 0 {
		input.Token = aws.String(_eksauthToken)
	}

	if resp, err := client.AssumeRoleForPodIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_eksauthCmd)
	_eksauthCmd.Flags().SortFlags = false

	_eksauthCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_eksauthCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_eksauthCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_eksauthCmd.Flags().StringVarP(&_eksauthClusterName, "cluster-name", "", "", "Cluster Name")
	_eksauthCmd.Flags().StringVarP(&_eksauthToken, "token", "", "", "Token")

	_eksauthCmd.Flags().BoolVarP(&_eksauthAssumeRoleForPodIdentity, "assume-role-for-pod-identity", "", false, "Assume Role For Pod Identity")

}
