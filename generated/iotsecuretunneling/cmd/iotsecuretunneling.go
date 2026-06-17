package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotsecuretunneling"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// iotsecuretunnelingCmd represents the iotsecuretunneling command
var _iotsecuretunnelingCmd = &cobra.Command{
	Use:   "iotsecuretunneling",
	Short: "AWS iotsecuretunneling CLI",
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
		client := iotsecuretunneling.NewFromConfig(cfg)
		if _iotsecuretunnelingCloseTunnel {
			iotsecuretunneling_CloseTunnel(cfg, client)
			return
		}
		if _iotsecuretunnelingDescribeTunnel {
			iotsecuretunneling_DescribeTunnel(cfg, client)
			return
		}
		if _iotsecuretunnelingListTagsForResource {
			iotsecuretunneling_ListTagsForResource(cfg, client)
			return
		}
		if _iotsecuretunnelingListTunnels {
			iotsecuretunneling_ListTunnels(cfg, client)
			return
		}
		if _iotsecuretunnelingOpenTunnel {
			iotsecuretunneling_OpenTunnel(cfg, client)
			return
		}
		if _iotsecuretunnelingRotateTunnelAccessToken {
			iotsecuretunneling_RotateTunnelAccessToken(cfg, client)
			return
		}
		if _iotsecuretunnelingTagResource {
			iotsecuretunneling_TagResource(cfg, client)
			return
		}
		if _iotsecuretunnelingUntagResource {
			iotsecuretunneling_UntagResource(cfg, client)
			return
		}

	},
}

var (
	_iotsecuretunnelingCloseTunnel             bool
	_iotsecuretunnelingDescribeTunnel          bool
	_iotsecuretunnelingListTagsForResource     bool
	_iotsecuretunnelingListTunnels             bool
	_iotsecuretunnelingOpenTunnel              bool
	_iotsecuretunnelingRotateTunnelAccessToken bool
	_iotsecuretunnelingTagResource             bool
	_iotsecuretunnelingUntagResource           bool

	_iotsecuretunnelingClientMode        string
	_iotsecuretunnelingDelete            string
	_iotsecuretunnelingDescription       string
	_iotsecuretunnelingDestinationConfig string
	_iotsecuretunnelingMaxResults        string
	_iotsecuretunnelingNextToken         string
	_iotsecuretunnelingResourceArn       string
	_iotsecuretunnelingTagKeys           []string
	_iotsecuretunnelingTags              string
	_iotsecuretunnelingThingName         string
	_iotsecuretunnelingTimeoutConfig     string
	_iotsecuretunnelingTunnelId          string
)

// Closes a tunnel identified by the unique tunnel id. When a CloseTunnel request
// is received, we close the WebSocket connections between the client and proxy
// server so no data can be transmitted.
//
// Requires permission to access the [CloseTunnel] action.
//
// [CloseTunnel]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iotsecuretunneling_CloseTunnel(cfg aws.Config, client *iotsecuretunneling.Client) {
	input := &iotsecuretunneling.CloseTunnelInput{
		// TunnelId: *string, // Required
	}

	if len(_iotsecuretunnelingTunnelId) > 0 {
		input.TunnelId = aws.String(_iotsecuretunnelingTunnelId)
	}
	if len(_iotsecuretunnelingDelete) > 0 {
		if err := assignInputField(input, "Delete", _iotsecuretunnelingDelete); err != nil {
			log.Errorf("invalid --delete: %s", err.Error())
			return
		}
	}

	if resp, err := client.CloseTunnel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a tunnel identified by the unique tunnel id.
// Requires permission to access the [DescribeTunnel] action.
//
// [DescribeTunnel]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iotsecuretunneling_DescribeTunnel(cfg aws.Config, client *iotsecuretunneling.Client) {
	input := &iotsecuretunneling.DescribeTunnelInput{
		// TunnelId: *string, // Required
	}

	if len(_iotsecuretunnelingTunnelId) > 0 {
		input.TunnelId = aws.String(_iotsecuretunnelingTunnelId)
	}

	if resp, err := client.DescribeTunnel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the tags for the specified resource.
func iotsecuretunneling_ListTagsForResource(cfg aws.Config, client *iotsecuretunneling.Client) {
	input := &iotsecuretunneling.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_iotsecuretunnelingResourceArn) > 0 {
		input.ResourceArn = aws.String(_iotsecuretunnelingResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List all tunnels for an Amazon Web Services account. Tunnels are listed by
// creation time in descending order, newer tunnels will be listed before older
// tunnels.
//
// Requires permission to access the [ListTunnels] action.
//
// [ListTunnels]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iotsecuretunneling_ListTunnels(cfg aws.Config, client *iotsecuretunneling.Client) {
	input := &iotsecuretunneling.ListTunnelsInput{}

	if len(_iotsecuretunnelingMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotsecuretunnelingMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotsecuretunnelingNextToken) > 0 {
		input.NextToken = aws.String(_iotsecuretunnelingNextToken)
	}
	if len(_iotsecuretunnelingThingName) > 0 {
		input.ThingName = aws.String(_iotsecuretunnelingThingName)
	}

	if disablePaginator() {
		if resp, err := client.ListTunnels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotsecuretunneling.ListTunnelsOutput
	p := iotsecuretunneling.NewListTunnelsPaginator(client, input)
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

// Creates a new tunnel, and returns two client access tokens for clients to use
// to connect to the IoT Secure Tunneling proxy server.
//
// Requires permission to access the [OpenTunnel] action.
//
// [OpenTunnel]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iotsecuretunneling_OpenTunnel(cfg aws.Config, client *iotsecuretunneling.Client) {
	input := &iotsecuretunneling.OpenTunnelInput{}

	if len(_iotsecuretunnelingDescription) > 0 {
		input.Description = aws.String(_iotsecuretunnelingDescription)
	}
	if len(_iotsecuretunnelingDestinationConfig) > 0 {
		if err := assignInputField(input, "DestinationConfig", _iotsecuretunnelingDestinationConfig); err != nil {
			log.Errorf("invalid --destination-config: %s", err.Error())
			return
		}
	}
	if len(_iotsecuretunnelingTags) > 0 {
		if err := assignInputField(input, "Tags", _iotsecuretunnelingTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_iotsecuretunnelingTimeoutConfig) > 0 {
		if err := assignInputField(input, "TimeoutConfig", _iotsecuretunnelingTimeoutConfig); err != nil {
			log.Errorf("invalid --timeout-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.OpenTunnel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Revokes the current client access token (CAT) and returns new CAT for clients
// to use when reconnecting to secure tunneling to access the same tunnel.
//
// Requires permission to access the [RotateTunnelAccessToken] action.
//
// Rotating the CAT doesn't extend the tunnel duration. For example, say the
// tunnel duration is 12 hours and the tunnel has already been open for 4 hours.
// When you rotate the access tokens, the new tokens that are generated can only be
// used for the remaining 8 hours.
//
// [RotateTunnelAccessToken]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iotsecuretunneling_RotateTunnelAccessToken(cfg aws.Config, client *iotsecuretunneling.Client) {
	input := &iotsecuretunneling.RotateTunnelAccessTokenInput{
		// ClientMode: types.ClientMode, // Required
		// TunnelId: *string, // Required
	}

	if len(_iotsecuretunnelingClientMode) > 0 {
		if err := assignInputField(input, "ClientMode", _iotsecuretunnelingClientMode); err != nil {
			log.Errorf("invalid --client-mode: %s", err.Error())
			return
		}
	}
	if len(_iotsecuretunnelingTunnelId) > 0 {
		input.TunnelId = aws.String(_iotsecuretunnelingTunnelId)
	}
	if len(_iotsecuretunnelingDestinationConfig) > 0 {
		if err := assignInputField(input, "DestinationConfig", _iotsecuretunnelingDestinationConfig); err != nil {
			log.Errorf("invalid --destination-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.RotateTunnelAccessToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A resource tag.
func iotsecuretunneling_TagResource(cfg aws.Config, client *iotsecuretunneling.Client) {
	input := &iotsecuretunneling.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_iotsecuretunnelingResourceArn) > 0 {
		input.ResourceArn = aws.String(_iotsecuretunnelingResourceArn)
	}
	if len(_iotsecuretunnelingTags) > 0 {
		if err := assignInputField(input, "Tags", _iotsecuretunnelingTags); err != nil {
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

// Removes a tag from a resource.
func iotsecuretunneling_UntagResource(cfg aws.Config, client *iotsecuretunneling.Client) {
	input := &iotsecuretunneling.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_iotsecuretunnelingResourceArn) > 0 {
		input.ResourceArn = aws.String(_iotsecuretunnelingResourceArn)
	}
	if len(_iotsecuretunnelingTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _iotsecuretunnelingTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_iotsecuretunnelingCmd)
	_iotsecuretunnelingCmd.Flags().SortFlags = false

	_iotsecuretunnelingCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_iotsecuretunnelingCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_iotsecuretunnelingCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_iotsecuretunnelingCmd.Flags().StringVarP(&_iotsecuretunnelingClientMode, "client-mode", "", "", "Client Mode")
	_iotsecuretunnelingCmd.Flags().StringVarP(&_iotsecuretunnelingDelete, "delete", "", "", "Delete")
	_iotsecuretunnelingCmd.Flags().StringVarP(&_iotsecuretunnelingDescription, "description", "", "", "Description")
	_iotsecuretunnelingCmd.Flags().StringVarP(&_iotsecuretunnelingDestinationConfig, "destination-config", "", "", "Destination Config")
	_iotsecuretunnelingCmd.Flags().StringVarP(&_iotsecuretunnelingMaxResults, "max-results", "", "", "Max Results")
	_iotsecuretunnelingCmd.Flags().StringVarP(&_iotsecuretunnelingNextToken, "next-token", "", "", "Next Token")
	_iotsecuretunnelingCmd.Flags().StringVarP(&_iotsecuretunnelingResourceArn, "resource-arn", "", "", "Resource ARN")
	_iotsecuretunnelingCmd.Flags().StringSliceVarP(&_iotsecuretunnelingTagKeys, "tag-keys", "", nil, "Tag Keys")
	_iotsecuretunnelingCmd.Flags().StringVarP(&_iotsecuretunnelingTags, "tags", "", "", "Tags")
	_iotsecuretunnelingCmd.Flags().StringVarP(&_iotsecuretunnelingThingName, "thing-name", "", "", "Thing Name")
	_iotsecuretunnelingCmd.Flags().StringVarP(&_iotsecuretunnelingTimeoutConfig, "timeout-config", "", "", "Timeout Config")
	_iotsecuretunnelingCmd.Flags().StringVarP(&_iotsecuretunnelingTunnelId, "tunnel-id", "", "", "Tunnel ID")

	_iotsecuretunnelingCmd.Flags().BoolVarP(&_iotsecuretunnelingCloseTunnel, "close-tunnel", "", false, "Close Tunnel")
	_iotsecuretunnelingCmd.Flags().BoolVarP(&_iotsecuretunnelingDescribeTunnel, "describe-tunnel", "", false, "Describe Tunnel")
	_iotsecuretunnelingCmd.Flags().BoolVarP(&_iotsecuretunnelingListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_iotsecuretunnelingCmd.Flags().BoolVarP(&_iotsecuretunnelingListTunnels, "list-tunnels", "", false, "List Tunnels")
	_iotsecuretunnelingCmd.Flags().BoolVarP(&_iotsecuretunnelingOpenTunnel, "open-tunnel", "", false, "Open Tunnel")
	_iotsecuretunnelingCmd.Flags().BoolVarP(&_iotsecuretunnelingRotateTunnelAccessToken, "rotate-tunnel-access-token", "", false, "Rotate Tunnel Access Token")
	_iotsecuretunnelingCmd.Flags().BoolVarP(&_iotsecuretunnelingTagResource, "tag-resource", "", false, "Tag Resource")
	_iotsecuretunnelingCmd.Flags().BoolVarP(&_iotsecuretunnelingUntagResource, "untag-resource", "", false, "Untag Resource")

}
