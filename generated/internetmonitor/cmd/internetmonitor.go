package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/internetmonitor"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// internetmonitorCmd represents the internetmonitor command
var _internetmonitorCmd = &cobra.Command{
	Use:   "internetmonitor",
	Short: "AWS internetmonitor CLI",
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
		client := internetmonitor.NewFromConfig(cfg)
		if _internetmonitorCreateMonitor {
			internetmonitor_CreateMonitor(cfg, client)
			return
		}
		if _internetmonitorDeleteMonitor {
			internetmonitor_DeleteMonitor(cfg, client)
			return
		}
		if _internetmonitorGetHealthEvent {
			internetmonitor_GetHealthEvent(cfg, client)
			return
		}
		if _internetmonitorGetInternetEvent {
			internetmonitor_GetInternetEvent(cfg, client)
			return
		}
		if _internetmonitorGetMonitor {
			internetmonitor_GetMonitor(cfg, client)
			return
		}
		if _internetmonitorGetQueryResults {
			internetmonitor_GetQueryResults(cfg, client)
			return
		}
		if _internetmonitorGetQueryStatus {
			internetmonitor_GetQueryStatus(cfg, client)
			return
		}
		if _internetmonitorListHealthEvents {
			internetmonitor_ListHealthEvents(cfg, client)
			return
		}
		if _internetmonitorListInternetEvents {
			internetmonitor_ListInternetEvents(cfg, client)
			return
		}
		if _internetmonitorListMonitors {
			internetmonitor_ListMonitors(cfg, client)
			return
		}
		if _internetmonitorListTagsForResource {
			internetmonitor_ListTagsForResource(cfg, client)
			return
		}
		if _internetmonitorStartQuery {
			internetmonitor_StartQuery(cfg, client)
			return
		}
		if _internetmonitorStopQuery {
			internetmonitor_StopQuery(cfg, client)
			return
		}
		if _internetmonitorTagResource {
			internetmonitor_TagResource(cfg, client)
			return
		}
		if _internetmonitorUntagResource {
			internetmonitor_UntagResource(cfg, client)
			return
		}
		if _internetmonitorUpdateMonitor {
			internetmonitor_UpdateMonitor(cfg, client)
			return
		}

	},
}

var (
	_internetmonitorCreateMonitor       bool
	_internetmonitorDeleteMonitor       bool
	_internetmonitorGetHealthEvent      bool
	_internetmonitorGetInternetEvent    bool
	_internetmonitorGetMonitor          bool
	_internetmonitorGetQueryResults     bool
	_internetmonitorGetQueryStatus      bool
	_internetmonitorListHealthEvents    bool
	_internetmonitorListInternetEvents  bool
	_internetmonitorListMonitors        bool
	_internetmonitorListTagsForResource bool
	_internetmonitorStartQuery          bool
	_internetmonitorStopQuery           bool
	_internetmonitorTagResource         bool
	_internetmonitorUntagResource       bool
	_internetmonitorUpdateMonitor       bool

	_internetmonitorClientToken                     string
	_internetmonitorEndTime                         string
	_internetmonitorEventId                         string
	_internetmonitorEventStatus                     string
	_internetmonitorEventType                       string
	_internetmonitorFilterParameters                string
	_internetmonitorHealthEventsConfig              string
	_internetmonitorIncludeLinkedAccounts           string
	_internetmonitorInternetMeasurementsLogDelivery string
	_internetmonitorLinkedAccountId                 string
	_internetmonitorMaxCityNetworksToMonitor        string
	_internetmonitorMaxResults                      string
	_internetmonitorMonitorName                     string
	_internetmonitorMonitorStatus                   string
	_internetmonitorNextToken                       string
	_internetmonitorQueryId                         string
	_internetmonitorQueryType                       string
	_internetmonitorResourceArn                     string
	_internetmonitorResources                       []string
	_internetmonitorResourcesToAdd                  []string
	_internetmonitorResourcesToRemove               []string
	_internetmonitorStartTime                       string
	_internetmonitorStatus                          string
	_internetmonitorTagKeys                         []string
	_internetmonitorTags                            string
	_internetmonitorTrafficPercentageToMonitor      string
)

// Creates a monitor in Amazon CloudWatch Internet Monitor. A monitor is built
// based on information from the application resources that you add: VPCs, Network
// Load Balancers (NLBs), Amazon CloudFront distributions, and Amazon WorkSpaces
// directories. Internet Monitor then publishes internet measurements from Amazon
// Web Services that are specific to the city-networks. That is, the locations and
// ASNs (typically internet service providers or ISPs), where clients access your
// application. For more information, see [Using Amazon CloudWatch Internet Monitor]in the Amazon CloudWatch User Guide.
//
// When you create a monitor, you choose the percentage of traffic that you want
// to monitor. You can also set a maximum limit for the number of city-networks
// where client traffic is monitored, that caps the total traffic that Internet
// Monitor monitors. A city-network maximum is the limit of city-networks, but you
// only pay for the number of city-networks that are actually monitored. You can
// update your monitor at any time to change the percentage of traffic to monitor
// or the city-networks maximum. For more information, see [Choosing a city-network maximum value]in the Amazon
// CloudWatch User Guide.
//
// [Using Amazon CloudWatch Internet Monitor]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-InternetMonitor.html
// [Choosing a city-network maximum value]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/IMCityNetworksMaximum.html
func internetmonitor_CreateMonitor(cfg aws.Config, client *internetmonitor.Client) {
	input := &internetmonitor.CreateMonitorInput{
		// MonitorName: *string, // Required
	}

	if len(_internetmonitorMonitorName) > 0 {
		input.MonitorName = aws.String(_internetmonitorMonitorName)
	}
	if len(_internetmonitorClientToken) > 0 {
		input.ClientToken = aws.String(_internetmonitorClientToken)
	}
	if len(_internetmonitorHealthEventsConfig) > 0 {
		if err := assignInputField(input, "HealthEventsConfig", _internetmonitorHealthEventsConfig); err != nil {
			log.Errorf("invalid --health-events-config: %s", err.Error())
			return
		}
	}
	if len(_internetmonitorInternetMeasurementsLogDelivery) > 0 {
		if err := assignInputField(input, "InternetMeasurementsLogDelivery", _internetmonitorInternetMeasurementsLogDelivery); err != nil {
			log.Errorf("invalid --internet-measurements-log-delivery: %s", err.Error())
			return
		}
	}
	if len(_internetmonitorMaxCityNetworksToMonitor) > 0 {
		if err := assignInputField(input, "MaxCityNetworksToMonitor", _internetmonitorMaxCityNetworksToMonitor); err != nil {
			log.Errorf("invalid --max-city-networks-to-monitor: %s", err.Error())
			return
		}
	}
	if len(_internetmonitorResources) > 0 {
		input.Resources = append([]string(nil), _internetmonitorResources...)
	}
	if len(_internetmonitorTags) > 0 {
		if err := assignInputField(input, "Tags", _internetmonitorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_internetmonitorTrafficPercentageToMonitor) > 0 {
		if err := assignInputField(input, "TrafficPercentageToMonitor", _internetmonitorTrafficPercentageToMonitor); err != nil {
			log.Errorf("invalid --traffic-percentage-to-monitor: %s", err.Error())
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

// Deletes a monitor in Amazon CloudWatch Internet Monitor.
func internetmonitor_DeleteMonitor(cfg aws.Config, client *internetmonitor.Client) {
	input := &internetmonitor.DeleteMonitorInput{
		// MonitorName: *string, // Required
	}

	if len(_internetmonitorMonitorName) > 0 {
		input.MonitorName = aws.String(_internetmonitorMonitorName)
	}

	if resp, err := client.DeleteMonitor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information that Amazon CloudWatch Internet Monitor has created and stored
// about a health event for a specified monitor. This information includes the
// impacted locations, and all the information related to the event, by location.
//
// The information returned includes the impact on performance, availability, and
// round-trip time, information about the network providers (ASNs), the event type,
// and so on.
//
// Information rolled up at the global traffic level is also returned, including
// the impact type and total traffic impact.
func internetmonitor_GetHealthEvent(cfg aws.Config, client *internetmonitor.Client) {
	input := &internetmonitor.GetHealthEventInput{
		// EventId: *string, // Required
		// MonitorName: *string, // Required
	}

	if len(_internetmonitorEventId) > 0 {
		input.EventId = aws.String(_internetmonitorEventId)
	}
	if len(_internetmonitorMonitorName) > 0 {
		input.MonitorName = aws.String(_internetmonitorMonitorName)
	}
	if len(_internetmonitorLinkedAccountId) > 0 {
		input.LinkedAccountId = aws.String(_internetmonitorLinkedAccountId)
	}

	if resp, err := client.GetHealthEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information that Amazon CloudWatch Internet Monitor has generated about an
// internet event. Internet Monitor displays information about recent global health
// events, called internet events, on a global outages map that is available to all
// Amazon Web Services customers.
//
// The information returned here includes the impacted location, when the event
// started and (if the event is over) ended, the type of event ( PERFORMANCE or
// AVAILABILITY ), and the status ( ACTIVE or RESOLVED ).
func internetmonitor_GetInternetEvent(cfg aws.Config, client *internetmonitor.Client) {
	input := &internetmonitor.GetInternetEventInput{
		// EventId: *string, // Required
	}

	if len(_internetmonitorEventId) > 0 {
		input.EventId = aws.String(_internetmonitorEventId)
	}

	if resp, err := client.GetInternetEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a monitor in Amazon CloudWatch Internet Monitor based on
// a monitor name. The information returned includes the Amazon Resource Name
// (ARN), create time, modified time, resources included in the monitor, and status
// information.
func internetmonitor_GetMonitor(cfg aws.Config, client *internetmonitor.Client) {
	input := &internetmonitor.GetMonitorInput{
		// MonitorName: *string, // Required
	}

	if len(_internetmonitorMonitorName) > 0 {
		input.MonitorName = aws.String(_internetmonitorMonitorName)
	}
	if len(_internetmonitorLinkedAccountId) > 0 {
		input.LinkedAccountId = aws.String(_internetmonitorLinkedAccountId)
	}

	if resp, err := client.GetMonitor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Return the data for a query with the Amazon CloudWatch Internet Monitor query
// interface. Specify the query that you want to return results for by providing a
// QueryId and a monitor name.
//
// For more information about using the query interface, including examples, see [Using the Amazon CloudWatch Internet Monitor query interface]
// in the Amazon CloudWatch Internet Monitor User Guide.
//
// [Using the Amazon CloudWatch Internet Monitor query interface]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-IM-view-cw-tools-cwim-query.html
func internetmonitor_GetQueryResults(cfg aws.Config, client *internetmonitor.Client) {
	input := &internetmonitor.GetQueryResultsInput{
		// MonitorName: *string, // Required
		// QueryId: *string, // Required
	}

	if len(_internetmonitorMonitorName) > 0 {
		input.MonitorName = aws.String(_internetmonitorMonitorName)
	}
	if len(_internetmonitorQueryId) > 0 {
		input.QueryId = aws.String(_internetmonitorQueryId)
	}
	if len(_internetmonitorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _internetmonitorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_internetmonitorNextToken) > 0 {
		input.NextToken = aws.String(_internetmonitorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetQueryResults(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*internetmonitor.GetQueryResultsOutput
	p := internetmonitor.NewGetQueryResultsPaginator(client, input)
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

// Returns the current status of a query for the Amazon CloudWatch Internet
// Monitor query interface, for a specified query ID and monitor. When you run a
// query, check the status to make sure that the query has SUCCEEDED before you
// review the results.
//
// - QUEUED : The query is scheduled to run.
//
// - RUNNING : The query is in progress but not complete.
//
// - SUCCEEDED : The query completed sucessfully.
//
// - FAILED : The query failed due to an error.
//
// - CANCELED : The query was canceled.
func internetmonitor_GetQueryStatus(cfg aws.Config, client *internetmonitor.Client) {
	input := &internetmonitor.GetQueryStatusInput{
		// MonitorName: *string, // Required
		// QueryId: *string, // Required
	}

	if len(_internetmonitorMonitorName) > 0 {
		input.MonitorName = aws.String(_internetmonitorMonitorName)
	}
	if len(_internetmonitorQueryId) > 0 {
		input.QueryId = aws.String(_internetmonitorQueryId)
	}

	if resp, err := client.GetQueryStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all health events for a monitor in Amazon CloudWatch Internet Monitor.
// Returns information for health events including the event start and end times,
// and the status.
//
// Health events that have start times during the time frame that is requested are
// not included in the list of health events.
func internetmonitor_ListHealthEvents(cfg aws.Config, client *internetmonitor.Client) {
	input := &internetmonitor.ListHealthEventsInput{
		// MonitorName: *string, // Required
	}

	if len(_internetmonitorMonitorName) > 0 {
		input.MonitorName = aws.String(_internetmonitorMonitorName)
	}
	if len(_internetmonitorEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _internetmonitorEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_internetmonitorEventStatus) > 0 {
		if err := assignInputField(input, "EventStatus", _internetmonitorEventStatus); err != nil {
			log.Errorf("invalid --event-status: %s", err.Error())
			return
		}
	}
	if len(_internetmonitorLinkedAccountId) > 0 {
		input.LinkedAccountId = aws.String(_internetmonitorLinkedAccountId)
	}
	if len(_internetmonitorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _internetmonitorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_internetmonitorNextToken) > 0 {
		input.NextToken = aws.String(_internetmonitorNextToken)
	}
	if len(_internetmonitorStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _internetmonitorStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListHealthEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*internetmonitor.ListHealthEventsOutput
	p := internetmonitor.NewListHealthEventsPaginator(client, input)
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

// Lists internet events that cause performance or availability issues for client
// locations. Amazon CloudWatch Internet Monitor displays information about recent
// global health events, called internet events, on a global outages map that is
// available to all Amazon Web Services customers.
//
// You can constrain the list of internet events returned by providing a start
// time and end time to define a total time frame for events you want to list. Both
// start time and end time specify the time when an event started. End time is
// optional. If you don't include it, the default end time is the current time.
//
// You can also limit the events returned to a specific status ( ACTIVE or RESOLVED
// ) or type ( PERFORMANCE or AVAILABILITY ).
func internetmonitor_ListInternetEvents(cfg aws.Config, client *internetmonitor.Client) {
	input := &internetmonitor.ListInternetEventsInput{}

	if len(_internetmonitorEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _internetmonitorEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_internetmonitorEventStatus) > 0 {
		input.EventStatus = aws.String(_internetmonitorEventStatus)
	}
	if len(_internetmonitorEventType) > 0 {
		input.EventType = aws.String(_internetmonitorEventType)
	}
	if len(_internetmonitorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _internetmonitorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_internetmonitorNextToken) > 0 {
		input.NextToken = aws.String(_internetmonitorNextToken)
	}
	if len(_internetmonitorStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _internetmonitorStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListInternetEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*internetmonitor.ListInternetEventsOutput
	p := internetmonitor.NewListInternetEventsPaginator(client, input)
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

// Lists all of your monitors for Amazon CloudWatch Internet Monitor and their
// statuses, along with the Amazon Resource Name (ARN) and name of each monitor.
func internetmonitor_ListMonitors(cfg aws.Config, client *internetmonitor.Client) {
	input := &internetmonitor.ListMonitorsInput{}

	if len(_internetmonitorIncludeLinkedAccounts) > 0 {
		if err := assignInputField(input, "IncludeLinkedAccounts", _internetmonitorIncludeLinkedAccounts); err != nil {
			log.Errorf("invalid --include-linked-accounts: %s", err.Error())
			return
		}
	}
	if len(_internetmonitorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _internetmonitorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_internetmonitorMonitorStatus) > 0 {
		input.MonitorStatus = aws.String(_internetmonitorMonitorStatus)
	}
	if len(_internetmonitorNextToken) > 0 {
		input.NextToken = aws.String(_internetmonitorNextToken)
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

	var results []*internetmonitor.ListMonitorsOutput
	p := internetmonitor.NewListMonitorsPaginator(client, input)
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

// Lists the tags for a resource. Tags are supported only for monitors in Amazon
// CloudWatch Internet Monitor.
func internetmonitor_ListTagsForResource(cfg aws.Config, client *internetmonitor.Client) {
	input := &internetmonitor.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_internetmonitorResourceArn) > 0 {
		input.ResourceArn = aws.String(_internetmonitorResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Start a query to return data for a specific query type for the Amazon
// CloudWatch Internet Monitor query interface. Specify a time period for the data
// that you want returned by using StartTime and EndTime . You filter the query
// results to return by providing parameters that you specify with FilterParameters
// .
//
// For more information about using the query interface, including examples, see [Using the Amazon CloudWatch Internet Monitor query interface]
// in the Amazon CloudWatch Internet Monitor User Guide.
//
// [Using the Amazon CloudWatch Internet Monitor query interface]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-IM-view-cw-tools-cwim-query.html
func internetmonitor_StartQuery(cfg aws.Config, client *internetmonitor.Client) {
	input := &internetmonitor.StartQueryInput{
		// EndTime: *time.Time, // Required
		// MonitorName: *string, // Required
		// QueryType: types.QueryType, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_internetmonitorEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _internetmonitorEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_internetmonitorMonitorName) > 0 {
		input.MonitorName = aws.String(_internetmonitorMonitorName)
	}
	if len(_internetmonitorQueryType) > 0 {
		if err := assignInputField(input, "QueryType", _internetmonitorQueryType); err != nil {
			log.Errorf("invalid --query-type: %s", err.Error())
			return
		}
	}
	if len(_internetmonitorStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _internetmonitorStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_internetmonitorFilterParameters) > 0 {
		if err := assignInputField(input, "FilterParameters", _internetmonitorFilterParameters); err != nil {
			log.Errorf("invalid --filter-parameters: %s", err.Error())
			return
		}
	}
	if len(_internetmonitorLinkedAccountId) > 0 {
		input.LinkedAccountId = aws.String(_internetmonitorLinkedAccountId)
	}

	if resp, err := client.StartQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stop a query that is progress for a specific monitor.
func internetmonitor_StopQuery(cfg aws.Config, client *internetmonitor.Client) {
	input := &internetmonitor.StopQueryInput{
		// MonitorName: *string, // Required
		// QueryId: *string, // Required
	}

	if len(_internetmonitorMonitorName) > 0 {
		input.MonitorName = aws.String(_internetmonitorMonitorName)
	}
	if len(_internetmonitorQueryId) > 0 {
		input.QueryId = aws.String(_internetmonitorQueryId)
	}

	if resp, err := client.StopQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a tag to a resource. Tags are supported only for monitors in Amazon
// CloudWatch Internet Monitor. You can add a maximum of 50 tags in Internet
// Monitor.
//
// A minimum of one tag is required for this call. It returns an error if you use
// the TagResource request with 0 tags.
func internetmonitor_TagResource(cfg aws.Config, client *internetmonitor.Client) {
	input := &internetmonitor.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_internetmonitorResourceArn) > 0 {
		input.ResourceArn = aws.String(_internetmonitorResourceArn)
	}
	if len(_internetmonitorTags) > 0 {
		if err := assignInputField(input, "Tags", _internetmonitorTags); err != nil {
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
func internetmonitor_UntagResource(cfg aws.Config, client *internetmonitor.Client) {
	input := &internetmonitor.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_internetmonitorResourceArn) > 0 {
		input.ResourceArn = aws.String(_internetmonitorResourceArn)
	}
	if len(_internetmonitorTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _internetmonitorTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a monitor. You can update a monitor to change the percentage of traffic
// to monitor or the maximum number of city-networks (locations and ASNs), to add
// or remove resources, or to change the status of the monitor. Note that you can't
// change the name of a monitor.
//
// The city-network maximum that you choose is the limit, but you only pay for the
// number of city-networks that are actually monitored. For more information, see [Choosing a city-network maximum value]
// in the Amazon CloudWatch User Guide.
//
// [Choosing a city-network maximum value]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/IMCityNetworksMaximum.html
func internetmonitor_UpdateMonitor(cfg aws.Config, client *internetmonitor.Client) {
	input := &internetmonitor.UpdateMonitorInput{
		// MonitorName: *string, // Required
	}

	if len(_internetmonitorMonitorName) > 0 {
		input.MonitorName = aws.String(_internetmonitorMonitorName)
	}
	if len(_internetmonitorClientToken) > 0 {
		input.ClientToken = aws.String(_internetmonitorClientToken)
	}
	if len(_internetmonitorHealthEventsConfig) > 0 {
		if err := assignInputField(input, "HealthEventsConfig", _internetmonitorHealthEventsConfig); err != nil {
			log.Errorf("invalid --health-events-config: %s", err.Error())
			return
		}
	}
	if len(_internetmonitorInternetMeasurementsLogDelivery) > 0 {
		if err := assignInputField(input, "InternetMeasurementsLogDelivery", _internetmonitorInternetMeasurementsLogDelivery); err != nil {
			log.Errorf("invalid --internet-measurements-log-delivery: %s", err.Error())
			return
		}
	}
	if len(_internetmonitorMaxCityNetworksToMonitor) > 0 {
		if err := assignInputField(input, "MaxCityNetworksToMonitor", _internetmonitorMaxCityNetworksToMonitor); err != nil {
			log.Errorf("invalid --max-city-networks-to-monitor: %s", err.Error())
			return
		}
	}
	if len(_internetmonitorResourcesToAdd) > 0 {
		input.ResourcesToAdd = append([]string(nil), _internetmonitorResourcesToAdd...)
	}
	if len(_internetmonitorResourcesToRemove) > 0 {
		input.ResourcesToRemove = append([]string(nil), _internetmonitorResourcesToRemove...)
	}
	if len(_internetmonitorStatus) > 0 {
		if err := assignInputField(input, "Status", _internetmonitorStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_internetmonitorTrafficPercentageToMonitor) > 0 {
		if err := assignInputField(input, "TrafficPercentageToMonitor", _internetmonitorTrafficPercentageToMonitor); err != nil {
			log.Errorf("invalid --traffic-percentage-to-monitor: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMonitor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_internetmonitorCmd)
	_internetmonitorCmd.Flags().SortFlags = false

	_internetmonitorCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_internetmonitorCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_internetmonitorCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_internetmonitorCmd.Flags().StringVarP(&_internetmonitorClientToken, "client-token", "", "", "Client Token")
	_internetmonitorCmd.Flags().StringVarP(&_internetmonitorEndTime, "end-time", "", "", "End Time")
	_internetmonitorCmd.Flags().StringVarP(&_internetmonitorEventId, "event-id", "", "", "Event ID")
	_internetmonitorCmd.Flags().StringVarP(&_internetmonitorEventStatus, "event-status", "", "", "Event Status")
	_internetmonitorCmd.Flags().StringVarP(&_internetmonitorEventType, "event-type", "", "", "Event Type")
	_internetmonitorCmd.Flags().StringVarP(&_internetmonitorFilterParameters, "filter-parameters", "", "", "Filter Parameters")
	_internetmonitorCmd.Flags().StringVarP(&_internetmonitorHealthEventsConfig, "health-events-config", "", "", "Health Events Config")
	_internetmonitorCmd.Flags().StringVarP(&_internetmonitorIncludeLinkedAccounts, "include-linked-accounts", "", "", "Include Linked Accounts")
	_internetmonitorCmd.Flags().StringVarP(&_internetmonitorInternetMeasurementsLogDelivery, "internet-measurements-log-delivery", "", "", "Internet Measurements Log Delivery")
	_internetmonitorCmd.Flags().StringVarP(&_internetmonitorLinkedAccountId, "linked-account-id", "", "", "Linked Account ID")
	_internetmonitorCmd.Flags().StringVarP(&_internetmonitorMaxCityNetworksToMonitor, "max-city-networks-to-monitor", "", "", "Max City Networks To Monitor")
	_internetmonitorCmd.Flags().StringVarP(&_internetmonitorMaxResults, "max-results", "", "", "Max Results")
	_internetmonitorCmd.Flags().StringVarP(&_internetmonitorMonitorName, "monitor-name", "", "", "Monitor Name")
	_internetmonitorCmd.Flags().StringVarP(&_internetmonitorMonitorStatus, "monitor-status", "", "", "Monitor Status")
	_internetmonitorCmd.Flags().StringVarP(&_internetmonitorNextToken, "next-token", "", "", "Next Token")
	_internetmonitorCmd.Flags().StringVarP(&_internetmonitorQueryId, "query-id", "", "", "Query ID")
	_internetmonitorCmd.Flags().StringVarP(&_internetmonitorQueryType, "query-type", "", "", "Query Type")
	_internetmonitorCmd.Flags().StringVarP(&_internetmonitorResourceArn, "resource-arn", "", "", "Resource ARN")
	_internetmonitorCmd.Flags().StringSliceVarP(&_internetmonitorResources, "resources", "", nil, "Resources")
	_internetmonitorCmd.Flags().StringSliceVarP(&_internetmonitorResourcesToAdd, "resources-to-add", "", nil, "Resources To Add")
	_internetmonitorCmd.Flags().StringSliceVarP(&_internetmonitorResourcesToRemove, "resources-to-remove", "", nil, "Resources To Remove")
	_internetmonitorCmd.Flags().StringVarP(&_internetmonitorStartTime, "start-time", "", "", "Start Time")
	_internetmonitorCmd.Flags().StringVarP(&_internetmonitorStatus, "status", "", "", "Status")
	_internetmonitorCmd.Flags().StringSliceVarP(&_internetmonitorTagKeys, "tag-keys", "", nil, "Tag Keys")
	_internetmonitorCmd.Flags().StringVarP(&_internetmonitorTags, "tags", "", "", "Tags")
	_internetmonitorCmd.Flags().StringVarP(&_internetmonitorTrafficPercentageToMonitor, "traffic-percentage-to-monitor", "", "", "Traffic Percentage To Monitor")

	_internetmonitorCmd.Flags().BoolVarP(&_internetmonitorCreateMonitor, "create-monitor", "", false, "Create Monitor")
	_internetmonitorCmd.Flags().BoolVarP(&_internetmonitorDeleteMonitor, "delete-monitor", "", false, "Delete Monitor")
	_internetmonitorCmd.Flags().BoolVarP(&_internetmonitorGetHealthEvent, "get-health-event", "", false, "Get Health Event")
	_internetmonitorCmd.Flags().BoolVarP(&_internetmonitorGetInternetEvent, "get-internet-event", "", false, "Get Internet Event")
	_internetmonitorCmd.Flags().BoolVarP(&_internetmonitorGetMonitor, "get-monitor", "", false, "Get Monitor")
	_internetmonitorCmd.Flags().BoolVarP(&_internetmonitorGetQueryResults, "get-query-results", "", false, "Get Query Results")
	_internetmonitorCmd.Flags().BoolVarP(&_internetmonitorGetQueryStatus, "get-query-status", "", false, "Get Query Status")
	_internetmonitorCmd.Flags().BoolVarP(&_internetmonitorListHealthEvents, "list-health-events", "", false, "List Health Events")
	_internetmonitorCmd.Flags().BoolVarP(&_internetmonitorListInternetEvents, "list-internet-events", "", false, "List Internet Events")
	_internetmonitorCmd.Flags().BoolVarP(&_internetmonitorListMonitors, "list-monitors", "", false, "List Monitors")
	_internetmonitorCmd.Flags().BoolVarP(&_internetmonitorListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_internetmonitorCmd.Flags().BoolVarP(&_internetmonitorStartQuery, "start-query", "", false, "Start Query")
	_internetmonitorCmd.Flags().BoolVarP(&_internetmonitorStopQuery, "stop-query", "", false, "Stop Query")
	_internetmonitorCmd.Flags().BoolVarP(&_internetmonitorTagResource, "tag-resource", "", false, "Tag Resource")
	_internetmonitorCmd.Flags().BoolVarP(&_internetmonitorUntagResource, "untag-resource", "", false, "Untag Resource")
	_internetmonitorCmd.Flags().BoolVarP(&_internetmonitorUpdateMonitor, "update-monitor", "", false, "Update Monitor")

}
