package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3outposts"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// s3outpostsCmd represents the s3outposts command
var _s3outpostsCmd = &cobra.Command{
	Use:   "s3outposts",
	Short: "AWS s3outposts CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := s3outposts.NewFromConfig(cfg)
		if _s3outpostsCreateEndpoint {
			s3outposts_CreateEndpoint(cfg, client)
			return
		}
		if _s3outpostsDeleteEndpoint {
			s3outposts_DeleteEndpoint(cfg, client)
			return
		}
		if _s3outpostsListEndpoints {
			s3outposts_ListEndpoints(cfg, client)
			return
		}
		if _s3outpostsListOutpostsWithS3 {
			s3outposts_ListOutpostsWithS3(cfg, client)
			return
		}
		if _s3outpostsListSharedEndpoints {
			s3outposts_ListSharedEndpoints(cfg, client)
			return
		}

	},
}

var (
	_s3outpostsCreateEndpoint      bool
	_s3outpostsDeleteEndpoint      bool
	_s3outpostsListEndpoints       bool
	_s3outpostsListOutpostsWithS3  bool
	_s3outpostsListSharedEndpoints bool

	_s3outpostsAccessType            string
	_s3outpostsCustomerOwnedIpv4Pool string
	_s3outpostsEndpointId            string
	_s3outpostsMaxResults            string
	_s3outpostsNextToken             string
	_s3outpostsOutpostId             string
	_s3outpostsSecurityGroupId       string
	_s3outpostsSubnetId              string
)

// Creates an endpoint and associates it with the specified Outpost.
// It can take up to 5 minutes for this action to finish.
//
// Related actions include:
//
// [DeleteEndpoint]
//
// [ListEndpoints]
//
// [ListEndpoints]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3outposts_ListEndpoints.html
// [DeleteEndpoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3outposts_DeleteEndpoint.html
func s3outposts_CreateEndpoint(cfg aws.Config, client *s3outposts.Client) {
	input := &s3outposts.CreateEndpointInput{
		// OutpostId: *string, // Required
		// SecurityGroupId: *string, // Required
		// SubnetId: *string, // Required
	}

	if len(_s3outpostsOutpostId) > 0 {
		input.OutpostId = aws.String(_s3outpostsOutpostId)
	}
	if len(_s3outpostsSecurityGroupId) > 0 {
		input.SecurityGroupId = aws.String(_s3outpostsSecurityGroupId)
	}
	if len(_s3outpostsSubnetId) > 0 {
		input.SubnetId = aws.String(_s3outpostsSubnetId)
	}
	if len(_s3outpostsAccessType) > 0 {
		if err := assignInputField(input, "AccessType", _s3outpostsAccessType); err != nil {
			log.Errorf("invalid --access-type: %s", err.Error())
			return
		}
	}
	if len(_s3outpostsCustomerOwnedIpv4Pool) > 0 {
		input.CustomerOwnedIpv4Pool = aws.String(_s3outpostsCustomerOwnedIpv4Pool)
	}

	if resp, err := client.CreateEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an endpoint.
// It can take up to 5 minutes for this action to finish.
//
// Related actions include:
//
// [CreateEndpoint]
//
// [ListEndpoints]
//
// [ListEndpoints]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3outposts_ListEndpoints.html
// [CreateEndpoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3outposts_CreateEndpoint.html
func s3outposts_DeleteEndpoint(cfg aws.Config, client *s3outposts.Client) {
	input := &s3outposts.DeleteEndpointInput{
		// EndpointId: *string, // Required
		// OutpostId: *string, // Required
	}

	if len(_s3outpostsEndpointId) > 0 {
		input.EndpointId = aws.String(_s3outpostsEndpointId)
	}
	if len(_s3outpostsOutpostId) > 0 {
		input.OutpostId = aws.String(_s3outpostsOutpostId)
	}

	if resp, err := client.DeleteEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists endpoints associated with the specified Outpost.
// Related actions include:
//
// [CreateEndpoint]
//
// [DeleteEndpoint]
//
// [CreateEndpoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3outposts_CreateEndpoint.html
// [DeleteEndpoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3outposts_DeleteEndpoint.html
func s3outposts_ListEndpoints(cfg aws.Config, client *s3outposts.Client) {
	input := &s3outposts.ListEndpointsInput{}

	if len(_s3outpostsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _s3outpostsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_s3outpostsNextToken) > 0 {
		input.NextToken = aws.String(_s3outpostsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEndpoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*s3outposts.ListEndpointsOutput
	p := s3outposts.NewListEndpointsPaginator(client, input)
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

// Lists the Outposts with S3 on Outposts capacity for your Amazon Web Services
// account. Includes S3 on Outposts that you have access to as the Outposts owner,
// or as a shared user from Resource Access Manager (RAM).
func s3outposts_ListOutpostsWithS3(cfg aws.Config, client *s3outposts.Client) {
	input := &s3outposts.ListOutpostsWithS3Input{}

	if len(_s3outpostsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _s3outpostsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_s3outpostsNextToken) > 0 {
		input.NextToken = aws.String(_s3outpostsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListOutpostsWithS3(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*s3outposts.ListOutpostsWithS3Output
	p := s3outposts.NewListOutpostsWithS3Paginator(client, input)
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

// Lists all endpoints associated with an Outpost that has been shared by Amazon
// Web Services Resource Access Manager (RAM).
//
// Related actions include:
//
// [CreateEndpoint]
//
// [DeleteEndpoint]
//
// [CreateEndpoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3outposts_CreateEndpoint.html
// [DeleteEndpoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3outposts_DeleteEndpoint.html
func s3outposts_ListSharedEndpoints(cfg aws.Config, client *s3outposts.Client) {
	input := &s3outposts.ListSharedEndpointsInput{
		// OutpostId: *string, // Required
	}

	if len(_s3outpostsOutpostId) > 0 {
		input.OutpostId = aws.String(_s3outpostsOutpostId)
	}
	if len(_s3outpostsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _s3outpostsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_s3outpostsNextToken) > 0 {
		input.NextToken = aws.String(_s3outpostsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSharedEndpoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*s3outposts.ListSharedEndpointsOutput
	p := s3outposts.NewListSharedEndpointsPaginator(client, input)
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

func init() {
	_rootCmd.AddCommand(_s3outpostsCmd)
	_s3outpostsCmd.Flags().SortFlags = false

	_s3outpostsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_s3outpostsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_s3outpostsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_s3outpostsCmd.Flags().StringVarP(&_s3outpostsAccessType, "access-type", "", "", "Access Type")
	_s3outpostsCmd.Flags().StringVarP(&_s3outpostsCustomerOwnedIpv4Pool, "customer-owned-ipv4-pool", "", "", "Customer Owned IPV4 Pool")
	_s3outpostsCmd.Flags().StringVarP(&_s3outpostsEndpointId, "endpoint-id", "", "", "Endpoint ID")
	_s3outpostsCmd.Flags().StringVarP(&_s3outpostsMaxResults, "max-results", "", "", "Max Results")
	_s3outpostsCmd.Flags().StringVarP(&_s3outpostsNextToken, "next-token", "", "", "Next Token")
	_s3outpostsCmd.Flags().StringVarP(&_s3outpostsOutpostId, "outpost-id", "", "", "Outpost ID")
	_s3outpostsCmd.Flags().StringVarP(&_s3outpostsSecurityGroupId, "security-group-id", "", "", "Security Group ID")
	_s3outpostsCmd.Flags().StringVarP(&_s3outpostsSubnetId, "subnet-id", "", "", "Subnet ID")

	_s3outpostsCmd.Flags().BoolVarP(&_s3outpostsCreateEndpoint, "create-endpoint", "", false, "Create Endpoint")
	_s3outpostsCmd.Flags().BoolVarP(&_s3outpostsDeleteEndpoint, "delete-endpoint", "", false, "Delete Endpoint")
	_s3outpostsCmd.Flags().BoolVarP(&_s3outpostsListEndpoints, "list-endpoints", "", false, "List Endpoints")
	_s3outpostsCmd.Flags().BoolVarP(&_s3outpostsListOutpostsWithS3, "list-outposts-with-s3", "", false, "List Outposts With S3")
	_s3outpostsCmd.Flags().BoolVarP(&_s3outpostsListSharedEndpoints, "list-shared-endpoints", "", false, "List Shared Endpoints")

}
