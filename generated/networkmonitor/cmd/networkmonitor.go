package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/networkmonitor"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// networkmonitorCmd represents the networkmonitor command
var _networkmonitorCmd = &cobra.Command{
	Use:   "networkmonitor",
	Short: "AWS networkmonitor CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := networkmonitor.NewFromConfig(cfg)
		if _networkmonitorCreateMonitor {
			networkmonitor_CreateMonitor(cfg, client)
			return
		}
		if _networkmonitorCreateProbe {
			networkmonitor_CreateProbe(cfg, client)
			return
		}
		if _networkmonitorDeleteMonitor {
			networkmonitor_DeleteMonitor(cfg, client)
			return
		}
		if _networkmonitorDeleteProbe {
			networkmonitor_DeleteProbe(cfg, client)
			return
		}
		if _networkmonitorGetMonitor {
			networkmonitor_GetMonitor(cfg, client)
			return
		}
		if _networkmonitorGetProbe {
			networkmonitor_GetProbe(cfg, client)
			return
		}
		if _networkmonitorListMonitors {
			networkmonitor_ListMonitors(cfg, client)
			return
		}
		if _networkmonitorListTagsForResource {
			networkmonitor_ListTagsForResource(cfg, client)
			return
		}
		if _networkmonitorTagResource {
			networkmonitor_TagResource(cfg, client)
			return
		}
		if _networkmonitorUntagResource {
			networkmonitor_UntagResource(cfg, client)
			return
		}
		if _networkmonitorUpdateMonitor {
			networkmonitor_UpdateMonitor(cfg, client)
			return
		}
		if _networkmonitorUpdateProbe {
			networkmonitor_UpdateProbe(cfg, client)
			return
		}

	},
}

var (
	_networkmonitorCreateMonitor       bool
	_networkmonitorCreateProbe         bool
	_networkmonitorDeleteMonitor       bool
	_networkmonitorDeleteProbe         bool
	_networkmonitorGetMonitor          bool
	_networkmonitorGetProbe            bool
	_networkmonitorListMonitors        bool
	_networkmonitorListTagsForResource bool
	_networkmonitorTagResource         bool
	_networkmonitorUntagResource       bool
	_networkmonitorUpdateMonitor       bool
	_networkmonitorUpdateProbe         bool

	_networkmonitorAggregationPeriod string
	_networkmonitorClientToken       string
	_networkmonitorDestination       string
	_networkmonitorDestinationPort   string
	_networkmonitorMaxResults        string
	_networkmonitorMonitorName       string
	_networkmonitorNextToken         string
	_networkmonitorPacketSize        string
	_networkmonitorProbe             string
	_networkmonitorProbeId           string
	_networkmonitorProbes            string
	_networkmonitorProtocol          string
	_networkmonitorResourceArn       string
	_networkmonitorState             string
	_networkmonitorTagKeys           []string
	_networkmonitorTags              string
)

// Creates a monitor between a source subnet and destination IP address. Within a
// monitor you'll create one or more probes that monitor network traffic between
// your source Amazon Web Services VPC subnets and your destination IP addresses.
// Each probe then aggregates and sends metrics to Amazon CloudWatch.
//
// You can also create a monitor with probes using this command. For each probe,
// you define the following:
//
// - source —The subnet IDs where the probes will be created.
//
// - destination — The target destination IP address for the probe.
//
// - destinationPort —Required only if the protocol is TCP .
//
// - protocol —The communication protocol between the source and destination.
// This will be either TCP or ICMP .
//
// - packetSize —The size of the packets. This must be a number between 56 and
// 8500 .
//
// - (Optional) tags —Key-value pairs created and assigned to the probe.
func networkmonitor_CreateMonitor(cfg aws.Config, client *networkmonitor.Client) {
	input := &networkmonitor.CreateMonitorInput{
		// MonitorName: *string, // Required
	}

	if len(_networkmonitorMonitorName) > 0 {
		input.MonitorName = aws.String(_networkmonitorMonitorName)
	}
	if len(_networkmonitorAggregationPeriod) > 0 {
		if err := assignInputField(input, "AggregationPeriod", _networkmonitorAggregationPeriod); err != nil {
			log.Errorf("invalid --aggregation-period: %s", err.Error())
			return
		}
	}
	if len(_networkmonitorClientToken) > 0 {
		input.ClientToken = aws.String(_networkmonitorClientToken)
	}
	if len(_networkmonitorProbes) > 0 {
		if err := assignInputField(input, "Probes", _networkmonitorProbes); err != nil {
			log.Errorf("invalid --probes: %s", err.Error())
			return
		}
	}
	if len(_networkmonitorTags) > 0 {
		if err := assignInputField(input, "Tags", _networkmonitorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMonitor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a probe within a monitor. Once you create a probe, and it begins
// monitoring your network traffic, you'll incur billing charges for that probe.
// This action requires the monitorName parameter. Run ListMonitors to get a list
// of monitor names. Note the name of the monitorName you want to create the probe
// for.
func networkmonitor_CreateProbe(cfg aws.Config, client *networkmonitor.Client) {
	input := &networkmonitor.CreateProbeInput{
		// MonitorName: *string, // Required
		// Probe: *types.ProbeInput, // Required
	}

	if len(_networkmonitorMonitorName) > 0 {
		input.MonitorName = aws.String(_networkmonitorMonitorName)
	}
	if len(_networkmonitorProbe) > 0 {
		if err := assignInputField(input, "Probe", _networkmonitorProbe); err != nil {
			log.Errorf("invalid --probe: %s", err.Error())
			return
		}
	}
	if len(_networkmonitorClientToken) > 0 {
		input.ClientToken = aws.String(_networkmonitorClientToken)
	}
	if len(_networkmonitorTags) > 0 {
		if err := assignInputField(input, "Tags", _networkmonitorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProbe(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified monitor.
// This action requires the monitorName parameter. Run ListMonitors to get a list
// of monitor names.
func networkmonitor_DeleteMonitor(cfg aws.Config, client *networkmonitor.Client) {
	input := &networkmonitor.DeleteMonitorInput{
		// MonitorName: *string, // Required
	}

	if len(_networkmonitorMonitorName) > 0 {
		input.MonitorName = aws.String(_networkmonitorMonitorName)
	}

	if resp, err := client.DeleteMonitor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified probe. Once a probe is deleted you'll no longer incur any
// billing fees for that probe.
//
// This action requires both the monitorName and probeId parameters. Run
// ListMonitors to get a list of monitor names. Run GetMonitor to get a list of
// probes and probe IDs. You can only delete a single probe at a time using this
// action.
func networkmonitor_DeleteProbe(cfg aws.Config, client *networkmonitor.Client) {
	input := &networkmonitor.DeleteProbeInput{
		// MonitorName: *string, // Required
		// ProbeId: *string, // Required
	}

	if len(_networkmonitorMonitorName) > 0 {
		input.MonitorName = aws.String(_networkmonitorMonitorName)
	}
	if len(_networkmonitorProbeId) > 0 {
		input.ProbeId = aws.String(_networkmonitorProbeId)
	}

	if resp, err := client.DeleteProbe(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details about a specific monitor.
// This action requires the monitorName parameter. Run ListMonitors to get a list
// of monitor names.
func networkmonitor_GetMonitor(cfg aws.Config, client *networkmonitor.Client) {
	input := &networkmonitor.GetMonitorInput{
		// MonitorName: *string, // Required
	}

	if len(_networkmonitorMonitorName) > 0 {
		input.MonitorName = aws.String(_networkmonitorMonitorName)
	}

	if resp, err := client.GetMonitor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the details about a probe. This action requires both the monitorName
// and probeId parameters. Run ListMonitors to get a list of monitor names. Run
// GetMonitor to get a list of probes and probe IDs.
func networkmonitor_GetProbe(cfg aws.Config, client *networkmonitor.Client) {
	input := &networkmonitor.GetProbeInput{
		// MonitorName: *string, // Required
		// ProbeId: *string, // Required
	}

	if len(_networkmonitorMonitorName) > 0 {
		input.MonitorName = aws.String(_networkmonitorMonitorName)
	}
	if len(_networkmonitorProbeId) > 0 {
		input.ProbeId = aws.String(_networkmonitorProbeId)
	}

	if resp, err := client.GetProbe(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of all of your monitors.
func networkmonitor_ListMonitors(cfg aws.Config, client *networkmonitor.Client) {
	input := &networkmonitor.ListMonitorsInput{}

	if len(_networkmonitorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkmonitorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkmonitorNextToken) > 0 {
		input.NextToken = aws.String(_networkmonitorNextToken)
	}
	if len(_networkmonitorState) > 0 {
		input.State = aws.String(_networkmonitorState)
	}

	if disablePaginator() {
		if resp, err := client.ListMonitors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkmonitor.ListMonitorsOutput
	p := networkmonitor.NewListMonitorsPaginator(client, input)
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

// Lists the tags assigned to this resource.
func networkmonitor_ListTagsForResource(cfg aws.Config, client *networkmonitor.Client) {
	input := &networkmonitor.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_networkmonitorResourceArn) > 0 {
		input.ResourceArn = aws.String(_networkmonitorResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds key-value pairs to a monitor or probe.
func networkmonitor_TagResource(cfg aws.Config, client *networkmonitor.Client) {
	input := &networkmonitor.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_networkmonitorResourceArn) > 0 {
		input.ResourceArn = aws.String(_networkmonitorResourceArn)
	}
	if len(_networkmonitorTags) > 0 {
		if err := assignInputField(input, "Tags", _networkmonitorTags); err != nil {
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

// Removes a key-value pair from a monitor or probe.
func networkmonitor_UntagResource(cfg aws.Config, client *networkmonitor.Client) {
	input := &networkmonitor.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_networkmonitorResourceArn) > 0 {
		input.ResourceArn = aws.String(_networkmonitorResourceArn)
	}
	if len(_networkmonitorTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _networkmonitorTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the aggregationPeriod for a monitor. Monitors support an
// aggregationPeriod of either 30 or 60 seconds. This action requires the
// monitorName and probeId parameter. Run ListMonitors to get a list of monitor
// names.
func networkmonitor_UpdateMonitor(cfg aws.Config, client *networkmonitor.Client) {
	input := &networkmonitor.UpdateMonitorInput{
		// AggregationPeriod: *int64, // Required
		// MonitorName: *string, // Required
	}

	if len(_networkmonitorAggregationPeriod) > 0 {
		if err := assignInputField(input, "AggregationPeriod", _networkmonitorAggregationPeriod); err != nil {
			log.Errorf("invalid --aggregation-period: %s", err.Error())
			return
		}
	}
	if len(_networkmonitorMonitorName) > 0 {
		input.MonitorName = aws.String(_networkmonitorMonitorName)
	}

	if resp, err := client.UpdateMonitor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a monitor probe. This action requires both the monitorName and probeId
// parameters. Run ListMonitors to get a list of monitor names. Run GetMonitor to
// get a list of probes and probe IDs.
//
// You can update the following para create a monitor with probes using this
// command. For each probe, you define the following:
//
// - state —The state of the probe.
//
// - destination — The target destination IP address for the probe.
//
// - destinationPort —Required only if the protocol is TCP .
//
// - protocol —The communication protocol between the source and destination.
// This will be either TCP or ICMP .
//
// - packetSize —The size of the packets. This must be a number between 56 and
// 8500 .
//
// - (Optional) tags —Key-value pairs created and assigned to the probe.
func networkmonitor_UpdateProbe(cfg aws.Config, client *networkmonitor.Client) {
	input := &networkmonitor.UpdateProbeInput{
		// MonitorName: *string, // Required
		// ProbeId: *string, // Required
	}

	if len(_networkmonitorMonitorName) > 0 {
		input.MonitorName = aws.String(_networkmonitorMonitorName)
	}
	if len(_networkmonitorProbeId) > 0 {
		input.ProbeId = aws.String(_networkmonitorProbeId)
	}
	if len(_networkmonitorDestination) > 0 {
		input.Destination = aws.String(_networkmonitorDestination)
	}
	if len(_networkmonitorDestinationPort) > 0 {
		if err := assignInputField(input, "DestinationPort", _networkmonitorDestinationPort); err != nil {
			log.Errorf("invalid --destination-port: %s", err.Error())
			return
		}
	}
	if len(_networkmonitorPacketSize) > 0 {
		if err := assignInputField(input, "PacketSize", _networkmonitorPacketSize); err != nil {
			log.Errorf("invalid --packet-size: %s", err.Error())
			return
		}
	}
	if len(_networkmonitorProtocol) > 0 {
		if err := assignInputField(input, "Protocol", _networkmonitorProtocol); err != nil {
			log.Errorf("invalid --protocol: %s", err.Error())
			return
		}
	}
	if len(_networkmonitorState) > 0 {
		if err := assignInputField(input, "State", _networkmonitorState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateProbe(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_networkmonitorCmd)
	_networkmonitorCmd.Flags().SortFlags = false

	_networkmonitorCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_networkmonitorCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_networkmonitorCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_networkmonitorCmd.Flags().StringVarP(&_networkmonitorAggregationPeriod, "aggregation-period", "", "", "Aggregation Period")
	_networkmonitorCmd.Flags().StringVarP(&_networkmonitorClientToken, "client-token", "", "", "Client Token")
	_networkmonitorCmd.Flags().StringVarP(&_networkmonitorDestination, "destination", "", "", "Destination")
	_networkmonitorCmd.Flags().StringVarP(&_networkmonitorDestinationPort, "destination-port", "", "", "Destination Port")
	_networkmonitorCmd.Flags().StringVarP(&_networkmonitorMaxResults, "max-results", "", "", "Max Results")
	_networkmonitorCmd.Flags().StringVarP(&_networkmonitorMonitorName, "monitor-name", "", "", "Monitor Name")
	_networkmonitorCmd.Flags().StringVarP(&_networkmonitorNextToken, "next-token", "", "", "Next Token")
	_networkmonitorCmd.Flags().StringVarP(&_networkmonitorPacketSize, "packet-size", "", "", "Packet Size")
	_networkmonitorCmd.Flags().StringVarP(&_networkmonitorProbe, "probe", "", "", "Probe")
	_networkmonitorCmd.Flags().StringVarP(&_networkmonitorProbeId, "probe-id", "", "", "Probe ID")
	_networkmonitorCmd.Flags().StringVarP(&_networkmonitorProbes, "probes", "", "", "Probes")
	_networkmonitorCmd.Flags().StringVarP(&_networkmonitorProtocol, "protocol", "", "", "Protocol")
	_networkmonitorCmd.Flags().StringVarP(&_networkmonitorResourceArn, "resource-arn", "", "", "Resource ARN")
	_networkmonitorCmd.Flags().StringVarP(&_networkmonitorState, "state", "", "", "State")
	_networkmonitorCmd.Flags().StringSliceVarP(&_networkmonitorTagKeys, "tag-keys", "", nil, "Tag Keys")
	_networkmonitorCmd.Flags().StringVarP(&_networkmonitorTags, "tags", "", "", "Tags")

	_networkmonitorCmd.Flags().BoolVarP(&_networkmonitorCreateMonitor, "create-monitor", "", false, "Create Monitor")
	_networkmonitorCmd.Flags().BoolVarP(&_networkmonitorCreateProbe, "create-probe", "", false, "Create Probe")
	_networkmonitorCmd.Flags().BoolVarP(&_networkmonitorDeleteMonitor, "delete-monitor", "", false, "Delete Monitor")
	_networkmonitorCmd.Flags().BoolVarP(&_networkmonitorDeleteProbe, "delete-probe", "", false, "Delete Probe")
	_networkmonitorCmd.Flags().BoolVarP(&_networkmonitorGetMonitor, "get-monitor", "", false, "Get Monitor")
	_networkmonitorCmd.Flags().BoolVarP(&_networkmonitorGetProbe, "get-probe", "", false, "Get Probe")
	_networkmonitorCmd.Flags().BoolVarP(&_networkmonitorListMonitors, "list-monitors", "", false, "List Monitors")
	_networkmonitorCmd.Flags().BoolVarP(&_networkmonitorListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_networkmonitorCmd.Flags().BoolVarP(&_networkmonitorTagResource, "tag-resource", "", false, "Tag Resource")
	_networkmonitorCmd.Flags().BoolVarP(&_networkmonitorUntagResource, "untag-resource", "", false, "Untag Resource")
	_networkmonitorCmd.Flags().BoolVarP(&_networkmonitorUpdateMonitor, "update-monitor", "", false, "Update Monitor")
	_networkmonitorCmd.Flags().BoolVarP(&_networkmonitorUpdateProbe, "update-probe", "", false, "Update Probe")

}
