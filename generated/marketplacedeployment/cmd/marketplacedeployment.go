package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/marketplacedeployment"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// marketplacedeploymentCmd represents the marketplacedeployment command
var _marketplacedeploymentCmd = &cobra.Command{
	Use:   "marketplacedeployment",
	Short: "AWS marketplacedeployment CLI",
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
		client := marketplacedeployment.NewFromConfig(cfg)
		if _marketplacedeploymentListTagsForResource {
			marketplacedeployment_ListTagsForResource(cfg, client)
			return
		}
		if _marketplacedeploymentPutDeploymentParameter {
			marketplacedeployment_PutDeploymentParameter(cfg, client)
			return
		}
		if _marketplacedeploymentTagResource {
			marketplacedeployment_TagResource(cfg, client)
			return
		}
		if _marketplacedeploymentUntagResource {
			marketplacedeployment_UntagResource(cfg, client)
			return
		}

	},
}

var (
	_marketplacedeploymentListTagsForResource    bool
	_marketplacedeploymentPutDeploymentParameter bool
	_marketplacedeploymentTagResource            bool
	_marketplacedeploymentUntagResource          bool

	_marketplacedeploymentAgreementId         string
	_marketplacedeploymentCatalog             string
	_marketplacedeploymentClientToken         string
	_marketplacedeploymentDeploymentParameter string
	_marketplacedeploymentExpirationDate      string
	_marketplacedeploymentProductId           string
	_marketplacedeploymentResourceArn         string
	_marketplacedeploymentTagKeys             []string
	_marketplacedeploymentTags                string
)

// Lists all tags that have been added to a deployment parameter resource.
func marketplacedeployment_ListTagsForResource(cfg aws.Config, client *marketplacedeployment.Client) {
	input := &marketplacedeployment.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_marketplacedeploymentResourceArn) > 0 {
		input.ResourceArn = aws.String(_marketplacedeploymentResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a deployment parameter and is targeted by catalog and
// agreementId .
func marketplacedeployment_PutDeploymentParameter(cfg aws.Config, client *marketplacedeployment.Client) {
	input := &marketplacedeployment.PutDeploymentParameterInput{
		// AgreementId: *string, // Required
		// Catalog: *string, // Required
		// DeploymentParameter: *types.DeploymentParameterInput, // Required
		// ProductId: *string, // Required
	}

	if len(_marketplacedeploymentAgreementId) > 0 {
		input.AgreementId = aws.String(_marketplacedeploymentAgreementId)
	}
	if len(_marketplacedeploymentCatalog) > 0 {
		input.Catalog = aws.String(_marketplacedeploymentCatalog)
	}
	if len(_marketplacedeploymentDeploymentParameter) > 0 {
		if err := assignInputField(input, "DeploymentParameter", _marketplacedeploymentDeploymentParameter); err != nil {
			log.Errorf("invalid --deployment-parameter: %s", err.Error())
			return
		}
	}
	if len(_marketplacedeploymentProductId) > 0 {
		input.ProductId = aws.String(_marketplacedeploymentProductId)
	}
	if len(_marketplacedeploymentClientToken) > 0 {
		input.ClientToken = aws.String(_marketplacedeploymentClientToken)
	}
	if len(_marketplacedeploymentExpirationDate) > 0 {
		if err := assignInputField(input, "ExpirationDate", _marketplacedeploymentExpirationDate); err != nil {
			log.Errorf("invalid --expiration-date: %s", err.Error())
			return
		}
	}
	if len(_marketplacedeploymentTags) > 0 {
		if err := assignInputField(input, "Tags", _marketplacedeploymentTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutDeploymentParameter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tags a resource.
func marketplacedeployment_TagResource(cfg aws.Config, client *marketplacedeployment.Client) {
	input := &marketplacedeployment.TagResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_marketplacedeploymentResourceArn) > 0 {
		input.ResourceArn = aws.String(_marketplacedeploymentResourceArn)
	}
	if len(_marketplacedeploymentTags) > 0 {
		if err := assignInputField(input, "Tags", _marketplacedeploymentTags); err != nil {
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

// Removes a tag or list of tags from a resource.
func marketplacedeployment_UntagResource(cfg aws.Config, client *marketplacedeployment.Client) {
	input := &marketplacedeployment.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_marketplacedeploymentResourceArn) > 0 {
		input.ResourceArn = aws.String(_marketplacedeploymentResourceArn)
	}
	if len(_marketplacedeploymentTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _marketplacedeploymentTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_marketplacedeploymentCmd)
	_marketplacedeploymentCmd.Flags().SortFlags = false

	_marketplacedeploymentCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_marketplacedeploymentCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_marketplacedeploymentCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_marketplacedeploymentCmd.Flags().StringVarP(&_marketplacedeploymentAgreementId, "agreement-id", "", "", "Agreement ID")
	_marketplacedeploymentCmd.Flags().StringVarP(&_marketplacedeploymentCatalog, "catalog", "", "", "Catalog")
	_marketplacedeploymentCmd.Flags().StringVarP(&_marketplacedeploymentClientToken, "client-token", "", "", "Client Token")
	_marketplacedeploymentCmd.Flags().StringVarP(&_marketplacedeploymentDeploymentParameter, "deployment-parameter", "", "", "Deployment Parameter")
	_marketplacedeploymentCmd.Flags().StringVarP(&_marketplacedeploymentExpirationDate, "expiration-date", "", "", "Expiration Date")
	_marketplacedeploymentCmd.Flags().StringVarP(&_marketplacedeploymentProductId, "product-id", "", "", "Product ID")
	_marketplacedeploymentCmd.Flags().StringVarP(&_marketplacedeploymentResourceArn, "resource-arn", "", "", "Resource ARN")
	_marketplacedeploymentCmd.Flags().StringSliceVarP(&_marketplacedeploymentTagKeys, "tag-keys", "", nil, "Tag Keys")
	_marketplacedeploymentCmd.Flags().StringVarP(&_marketplacedeploymentTags, "tags", "", "", "Tags")

	_marketplacedeploymentCmd.Flags().BoolVarP(&_marketplacedeploymentListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_marketplacedeploymentCmd.Flags().BoolVarP(&_marketplacedeploymentPutDeploymentParameter, "put-deployment-parameter", "", false, "Put Deployment Parameter")
	_marketplacedeploymentCmd.Flags().BoolVarP(&_marketplacedeploymentTagResource, "tag-resource", "", false, "Tag Resource")
	_marketplacedeploymentCmd.Flags().BoolVarP(&_marketplacedeploymentUntagResource, "untag-resource", "", false, "Untag Resource")

}
