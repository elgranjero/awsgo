package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/networkflowmonitor"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// networkflowmonitorCmd represents the networkflowmonitor command
var _networkflowmonitorCmd = &cobra.Command{
	Use:   "networkflowmonitor",
	Short: "AWS networkflowmonitor CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := networkflowmonitor.NewFromConfig(cfg)
		if _networkflowmonitorCreateMonitor {
			networkflowmonitor_CreateMonitor(cfg, client)
			return
		}
		if _networkflowmonitorCreateScope {
			networkflowmonitor_CreateScope(cfg, client)
			return
		}
		if _networkflowmonitorDeleteMonitor {
			networkflowmonitor_DeleteMonitor(cfg, client)
			return
		}
		if _networkflowmonitorDeleteScope {
			networkflowmonitor_DeleteScope(cfg, client)
			return
		}
		if _networkflowmonitorGetMonitor {
			networkflowmonitor_GetMonitor(cfg, client)
			return
		}
		if _networkflowmonitorGetQueryResultsMonitorTopContributors {
			networkflowmonitor_GetQueryResultsMonitorTopContributors(cfg, client)
			return
		}
		if _networkflowmonitorGetQueryResultsWorkloadInsightsTopContributors {
			networkflowmonitor_GetQueryResultsWorkloadInsightsTopContributors(cfg, client)
			return
		}
		if _networkflowmonitorGetQueryResultsWorkloadInsightsTopContributorsData {
			networkflowmonitor_GetQueryResultsWorkloadInsightsTopContributorsData(cfg, client)
			return
		}
		if _networkflowmonitorGetQueryStatusMonitorTopContributors {
			networkflowmonitor_GetQueryStatusMonitorTopContributors(cfg, client)
			return
		}
		if _networkflowmonitorGetQueryStatusWorkloadInsightsTopContributors {
			networkflowmonitor_GetQueryStatusWorkloadInsightsTopContributors(cfg, client)
			return
		}
		if _networkflowmonitorGetQueryStatusWorkloadInsightsTopContributorsData {
			networkflowmonitor_GetQueryStatusWorkloadInsightsTopContributorsData(cfg, client)
			return
		}
		if _networkflowmonitorGetScope {
			networkflowmonitor_GetScope(cfg, client)
			return
		}
		if _networkflowmonitorListMonitors {
			networkflowmonitor_ListMonitors(cfg, client)
			return
		}
		if _networkflowmonitorListScopes {
			networkflowmonitor_ListScopes(cfg, client)
			return
		}
		if _networkflowmonitorListTagsForResource {
			networkflowmonitor_ListTagsForResource(cfg, client)
			return
		}
		if _networkflowmonitorStartQueryMonitorTopContributors {
			networkflowmonitor_StartQueryMonitorTopContributors(cfg, client)
			return
		}
		if _networkflowmonitorStartQueryWorkloadInsightsTopContributors {
			networkflowmonitor_StartQueryWorkloadInsightsTopContributors(cfg, client)
			return
		}
		if _networkflowmonitorStartQueryWorkloadInsightsTopContributorsData {
			networkflowmonitor_StartQueryWorkloadInsightsTopContributorsData(cfg, client)
			return
		}
		if _networkflowmonitorStopQueryMonitorTopContributors {
			networkflowmonitor_StopQueryMonitorTopContributors(cfg, client)
			return
		}
		if _networkflowmonitorStopQueryWorkloadInsightsTopContributors {
			networkflowmonitor_StopQueryWorkloadInsightsTopContributors(cfg, client)
			return
		}
		if _networkflowmonitorStopQueryWorkloadInsightsTopContributorsData {
			networkflowmonitor_StopQueryWorkloadInsightsTopContributorsData(cfg, client)
			return
		}
		if _networkflowmonitorTagResource {
			networkflowmonitor_TagResource(cfg, client)
			return
		}
		if _networkflowmonitorUntagResource {
			networkflowmonitor_UntagResource(cfg, client)
			return
		}
		if _networkflowmonitorUpdateMonitor {
			networkflowmonitor_UpdateMonitor(cfg, client)
			return
		}
		if _networkflowmonitorUpdateScope {
			networkflowmonitor_UpdateScope(cfg, client)
			return
		}

	},
}

var (
	_networkflowmonitorCreateMonitor                                      bool
	_networkflowmonitorCreateScope                                        bool
	_networkflowmonitorDeleteMonitor                                      bool
	_networkflowmonitorDeleteScope                                        bool
	_networkflowmonitorGetMonitor                                         bool
	_networkflowmonitorGetQueryResultsMonitorTopContributors              bool
	_networkflowmonitorGetQueryResultsWorkloadInsightsTopContributors     bool
	_networkflowmonitorGetQueryResultsWorkloadInsightsTopContributorsData bool
	_networkflowmonitorGetQueryStatusMonitorTopContributors               bool
	_networkflowmonitorGetQueryStatusWorkloadInsightsTopContributors      bool
	_networkflowmonitorGetQueryStatusWorkloadInsightsTopContributorsData  bool
	_networkflowmonitorGetScope                                           bool
	_networkflowmonitorListMonitors                                       bool
	_networkflowmonitorListScopes                                         bool
	_networkflowmonitorListTagsForResource                                bool
	_networkflowmonitorStartQueryMonitorTopContributors                   bool
	_networkflowmonitorStartQueryWorkloadInsightsTopContributors          bool
	_networkflowmonitorStartQueryWorkloadInsightsTopContributorsData      bool
	_networkflowmonitorStopQueryMonitorTopContributors                    bool
	_networkflowmonitorStopQueryWorkloadInsightsTopContributors           bool
	_networkflowmonitorStopQueryWorkloadInsightsTopContributorsData       bool
	_networkflowmonitorTagResource                                        bool
	_networkflowmonitorUntagResource                                      bool
	_networkflowmonitorUpdateMonitor                                      bool
	_networkflowmonitorUpdateScope                                        bool

	_networkflowmonitorClientToken             string
	_networkflowmonitorDestinationCategory     string
	_networkflowmonitorEndTime                 string
	_networkflowmonitorLimit                   string
	_networkflowmonitorLocalResources          string
	_networkflowmonitorLocalResourcesToAdd     string
	_networkflowmonitorLocalResourcesToRemove  string
	_networkflowmonitorMaxResults              string
	_networkflowmonitorMetricName              string
	_networkflowmonitorMonitorName             string
	_networkflowmonitorMonitorStatus           string
	_networkflowmonitorNextToken               string
	_networkflowmonitorQueryId                 string
	_networkflowmonitorRemoteResources         string
	_networkflowmonitorRemoteResourcesToAdd    string
	_networkflowmonitorRemoteResourcesToRemove string
	_networkflowmonitorResourceArn             string
	_networkflowmonitorResourcesToAdd          string
	_networkflowmonitorResourcesToDelete       string
	_networkflowmonitorScopeArn                string
	_networkflowmonitorScopeId                 string
	_networkflowmonitorStartTime               string
	_networkflowmonitorTagKeys                 []string
	_networkflowmonitorTags                    string
	_networkflowmonitorTargets                 string
)

// Create a monitor for specific network flows between local and remote resources,
// so that you can monitor network performance for one or several of your
// workloads. For each monitor, Network Flow Monitor publishes detailed end-to-end
// performance metrics and a network health indicator (NHI) that informs you
// whether there were Amazon Web Services network issues for one or more of the
// network flows tracked by a monitor, during a time period that you choose.
func networkflowmonitor_CreateMonitor(cfg aws.Config, client *networkflowmonitor.Client) {
	input := &networkflowmonitor.CreateMonitorInput{
		// LocalResources: []types.MonitorLocalResource, // Required
		// MonitorName: *string, // Required
		// ScopeArn: *string, // Required
	}

	if len(_networkflowmonitorLocalResources) > 0 {
		if err := assignInputField(input, "LocalResources", _networkflowmonitorLocalResources); err != nil {
			log.Errorf("invalid --local-resources: %s", err.Error())
			return
		}
	}
	if len(_networkflowmonitorMonitorName) > 0 {
		input.MonitorName = aws.String(_networkflowmonitorMonitorName)
	}
	if len(_networkflowmonitorScopeArn) > 0 {
		input.ScopeArn = aws.String(_networkflowmonitorScopeArn)
	}
	if len(_networkflowmonitorClientToken) > 0 {
		input.ClientToken = aws.String(_networkflowmonitorClientToken)
	}
	if len(_networkflowmonitorRemoteResources) > 0 {
		if err := assignInputField(input, "RemoteResources", _networkflowmonitorRemoteResources); err != nil {
			log.Errorf("invalid --remote-resources: %s", err.Error())
			return
		}
	}
	if len(_networkflowmonitorTags) > 0 {
		if err := assignInputField(input, "Tags", _networkflowmonitorTags); err != nil {
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

// In Network Flow Monitor, you specify a scope for the service to generate
// metrics for. By using the scope, Network Flow Monitor can generate a topology of
// all the resources to measure performance metrics for. When you create a scope,
// you enable permissions for Network Flow Monitor.
//
// A scope is a Region-account pair or multiple Region-account pairs. Network Flow
// Monitor uses your scope to determine all the resources (the topology) where
// Network Flow Monitor will gather network flow performance metrics for you. To
// provide performance metrics, Network Flow Monitor uses the data that is sent by
// the Network Flow Monitor agents you install on the resources.
//
// To define the Region-account pairs for your scope, the Network Flow Monitor API
// uses the following constucts, which allow for future flexibility in defining
// scopes:
//
// - Targets, which are arrays of targetResources.
//
// - Target resources, which are Region-targetIdentifier pairs.
//
// - Target identifiers, made up of a targetID (currently always an account ID)
// and a targetType (currently always an account).
func networkflowmonitor_CreateScope(cfg aws.Config, client *networkflowmonitor.Client) {
	input := &networkflowmonitor.CreateScopeInput{
		// Targets: []types.TargetResource, // Required
	}

	if len(_networkflowmonitorTargets) > 0 {
		if err := assignInputField(input, "Targets", _networkflowmonitorTargets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}
	if len(_networkflowmonitorClientToken) > 0 {
		input.ClientToken = aws.String(_networkflowmonitorClientToken)
	}
	if len(_networkflowmonitorTags) > 0 {
		if err := assignInputField(input, "Tags", _networkflowmonitorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateScope(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a monitor in Network Flow Monitor.
func networkflowmonitor_DeleteMonitor(cfg aws.Config, client *networkflowmonitor.Client) {
	input := &networkflowmonitor.DeleteMonitorInput{
		// MonitorName: *string, // Required
	}

	if len(_networkflowmonitorMonitorName) > 0 {
		input.MonitorName = aws.String(_networkflowmonitorMonitorName)
	}

	if resp, err := client.DeleteMonitor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a scope that has been defined.
func networkflowmonitor_DeleteScope(cfg aws.Config, client *networkflowmonitor.Client) {
	input := &networkflowmonitor.DeleteScopeInput{
		// ScopeId: *string, // Required
	}

	if len(_networkflowmonitorScopeId) > 0 {
		input.ScopeId = aws.String(_networkflowmonitorScopeId)
	}

	if resp, err := client.DeleteScope(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a monitor in Network Flow Monitor based on a monitor
// name. The information returned includes the Amazon Resource Name (ARN), create
// time, modified time, resources included in the monitor, and status information.
func networkflowmonitor_GetMonitor(cfg aws.Config, client *networkflowmonitor.Client) {
	input := &networkflowmonitor.GetMonitorInput{
		// MonitorName: *string, // Required
	}

	if len(_networkflowmonitorMonitorName) > 0 {
		input.MonitorName = aws.String(_networkflowmonitorMonitorName)
	}

	if resp, err := client.GetMonitor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Return the data for a query with the Network Flow Monitor query interface. You
// specify the query that you want to return results for by providing a query ID
// and a monitor name. This query returns the top contributors for a specific
// monitor.
//
// Create a query ID for this call by calling the corresponding API call to start
// the query, StartQueryMonitorTopContributors . Use the scope ID that was returned
// for your account by CreateScope .
//
// Top contributors in Network Flow Monitor are network flows with the highest
// values for a specific metric type. Top contributors can be across all workload
// insights, for a given scope, or for a specific monitor. Use the applicable call
// for the top contributors that you want to be returned.
func networkflowmonitor_GetQueryResultsMonitorTopContributors(cfg aws.Config, client *networkflowmonitor.Client) {
	input := &networkflowmonitor.GetQueryResultsMonitorTopContributorsInput{
		// MonitorName: *string, // Required
		// QueryId: *string, // Required
	}

	if len(_networkflowmonitorMonitorName) > 0 {
		input.MonitorName = aws.String(_networkflowmonitorMonitorName)
	}
	if len(_networkflowmonitorQueryId) > 0 {
		input.QueryId = aws.String(_networkflowmonitorQueryId)
	}
	if len(_networkflowmonitorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkflowmonitorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkflowmonitorNextToken) > 0 {
		input.NextToken = aws.String(_networkflowmonitorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetQueryResultsMonitorTopContributors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkflowmonitor.GetQueryResultsMonitorTopContributorsOutput
	p := networkflowmonitor.NewGetQueryResultsMonitorTopContributorsPaginator(client, input)
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

// Return the data for a query with the Network Flow Monitor query interface. You
// specify the query that you want to return results for by providing a query ID
// and a monitor name.
//
// This query returns the top contributors for a scope for workload insights.
// Workload insights provide a high level view of network flow performance data
// collected by agents. To return the data for the top contributors, see
// GetQueryResultsWorkloadInsightsTopContributorsData .
//
// Create a query ID for this call by calling the corresponding API call to start
// the query, StartQueryWorkloadInsightsTopContributors . Use the scope ID that was
// returned for your account by CreateScope .
//
// Top contributors in Network Flow Monitor are network flows with the highest
// values for a specific metric type. Top contributors can be across all workload
// insights, for a given scope, or for a specific monitor. Use the applicable call
// for the top contributors that you want to be returned.
func networkflowmonitor_GetQueryResultsWorkloadInsightsTopContributors(cfg aws.Config, client *networkflowmonitor.Client) {
	input := &networkflowmonitor.GetQueryResultsWorkloadInsightsTopContributorsInput{
		// QueryId: *string, // Required
		// ScopeId: *string, // Required
	}

	if len(_networkflowmonitorQueryId) > 0 {
		input.QueryId = aws.String(_networkflowmonitorQueryId)
	}
	if len(_networkflowmonitorScopeId) > 0 {
		input.ScopeId = aws.String(_networkflowmonitorScopeId)
	}
	if len(_networkflowmonitorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkflowmonitorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkflowmonitorNextToken) > 0 {
		input.NextToken = aws.String(_networkflowmonitorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetQueryResultsWorkloadInsightsTopContributors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkflowmonitor.GetQueryResultsWorkloadInsightsTopContributorsOutput
	p := networkflowmonitor.NewGetQueryResultsWorkloadInsightsTopContributorsPaginator(client, input)
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

// Return the data for a query with the Network Flow Monitor query interface.
// Specify the query that you want to return results for by providing a query ID
// and a scope ID.
//
// This query returns the data for top contributors for workload insights for a
// specific scope. Workload insights provide a high level view of network flow
// performance data collected by agents for a scope. To return just the top
// contributors, see GetQueryResultsWorkloadInsightsTopContributors .
//
// Create a query ID for this call by calling the corresponding API call to start
// the query, StartQueryWorkloadInsightsTopContributorsData . Use the scope ID that
// was returned for your account by CreateScope .
//
// Top contributors in Network Flow Monitor are network flows with the highest
// values for a specific metric type. Top contributors can be across all workload
// insights, for a given scope, or for a specific monitor. Use the applicable call
// for the top contributors that you want to be returned.
//
// The top contributor network flows overall are for a specific metric type, for
// example, the number of retransmissions.
func networkflowmonitor_GetQueryResultsWorkloadInsightsTopContributorsData(cfg aws.Config, client *networkflowmonitor.Client) {
	input := &networkflowmonitor.GetQueryResultsWorkloadInsightsTopContributorsDataInput{
		// QueryId: *string, // Required
		// ScopeId: *string, // Required
	}

	if len(_networkflowmonitorQueryId) > 0 {
		input.QueryId = aws.String(_networkflowmonitorQueryId)
	}
	if len(_networkflowmonitorScopeId) > 0 {
		input.ScopeId = aws.String(_networkflowmonitorScopeId)
	}
	if len(_networkflowmonitorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkflowmonitorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkflowmonitorNextToken) > 0 {
		input.NextToken = aws.String(_networkflowmonitorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetQueryResultsWorkloadInsightsTopContributorsData(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkflowmonitor.GetQueryResultsWorkloadInsightsTopContributorsDataOutput
	p := networkflowmonitor.NewGetQueryResultsWorkloadInsightsTopContributorsDataPaginator(client, input)
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

// Returns the current status of a query for the Network Flow Monitor query
// interface, for a specified query ID and monitor. This call returns the query
// status for the top contributors for a monitor.
//
// When you create a query, use this call to check the status of the query to make
// sure that it has has SUCCEEDED before you review the results. Use the same
// query ID that you used for the corresponding API call to start (create) the
// query, StartQueryMonitorTopContributors .
//
// When you run a query, use this call to check the status of the query to make
// sure that the query has SUCCEEDED before you review the results.
func networkflowmonitor_GetQueryStatusMonitorTopContributors(cfg aws.Config, client *networkflowmonitor.Client) {
	input := &networkflowmonitor.GetQueryStatusMonitorTopContributorsInput{
		// MonitorName: *string, // Required
		// QueryId: *string, // Required
	}

	if len(_networkflowmonitorMonitorName) > 0 {
		input.MonitorName = aws.String(_networkflowmonitorMonitorName)
	}
	if len(_networkflowmonitorQueryId) > 0 {
		input.QueryId = aws.String(_networkflowmonitorQueryId)
	}

	if resp, err := client.GetQueryStatusMonitorTopContributors(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Return the data for a query with the Network Flow Monitor query interface.
// Specify the query that you want to return results for by providing a query ID
// and a monitor name. This query returns the top contributors for workload
// insights.
//
// When you start a query, use this call to check the status of the query to make
// sure that it has has SUCCEEDED before you review the results. Use the same
// query ID that you used for the corresponding API call to start the query,
// StartQueryWorkloadInsightsTopContributors .
//
// Top contributors in Network Flow Monitor are network flows with the highest
// values for a specific metric type. Top contributors can be across all workload
// insights, for a given scope, or for a specific monitor. Use the applicable call
// for the top contributors that you want to be returned.
func networkflowmonitor_GetQueryStatusWorkloadInsightsTopContributors(cfg aws.Config, client *networkflowmonitor.Client) {
	input := &networkflowmonitor.GetQueryStatusWorkloadInsightsTopContributorsInput{
		// QueryId: *string, // Required
		// ScopeId: *string, // Required
	}

	if len(_networkflowmonitorQueryId) > 0 {
		input.QueryId = aws.String(_networkflowmonitorQueryId)
	}
	if len(_networkflowmonitorScopeId) > 0 {
		input.ScopeId = aws.String(_networkflowmonitorScopeId)
	}

	if resp, err := client.GetQueryStatusWorkloadInsightsTopContributors(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the current status of a query for the Network Flow Monitor query
// interface, for a specified query ID and monitor. This call returns the query
// status for the top contributors data for workload insights.
//
// When you start a query, use this call to check the status of the query to make
// sure that it has has SUCCEEDED before you review the results. Use the same
// query ID that you used for the corresponding API call to start the query,
// StartQueryWorkloadInsightsTopContributorsData .
//
// Top contributors in Network Flow Monitor are network flows with the highest
// values for a specific metric type. Top contributors can be across all workload
// insights, for a given scope, or for a specific monitor. Use the applicable call
// for the top contributors that you want to be returned.
//
// The top contributor network flows overall are for a specific metric type, for
// example, the number of retransmissions.
func networkflowmonitor_GetQueryStatusWorkloadInsightsTopContributorsData(cfg aws.Config, client *networkflowmonitor.Client) {
	input := &networkflowmonitor.GetQueryStatusWorkloadInsightsTopContributorsDataInput{
		// QueryId: *string, // Required
		// ScopeId: *string, // Required
	}

	if len(_networkflowmonitorQueryId) > 0 {
		input.QueryId = aws.String(_networkflowmonitorQueryId)
	}
	if len(_networkflowmonitorScopeId) > 0 {
		input.ScopeId = aws.String(_networkflowmonitorScopeId)
	}

	if resp, err := client.GetQueryStatusWorkloadInsightsTopContributorsData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a scope, including the name, status, tags, and target
// details. The scope in Network Flow Monitor is an account.
func networkflowmonitor_GetScope(cfg aws.Config, client *networkflowmonitor.Client) {
	input := &networkflowmonitor.GetScopeInput{
		// ScopeId: *string, // Required
	}

	if len(_networkflowmonitorScopeId) > 0 {
		input.ScopeId = aws.String(_networkflowmonitorScopeId)
	}

	if resp, err := client.GetScope(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List all monitors in an account. Optionally, you can list only monitors that
// have a specific status, by using the STATUS parameter.
func networkflowmonitor_ListMonitors(cfg aws.Config, client *networkflowmonitor.Client) {
	input := &networkflowmonitor.ListMonitorsInput{}

	if len(_networkflowmonitorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkflowmonitorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkflowmonitorMonitorStatus) > 0 {
		if err := assignInputField(input, "MonitorStatus", _networkflowmonitorMonitorStatus); err != nil {
			log.Errorf("invalid --monitor-status: %s", err.Error())
			return
		}
	}
	if len(_networkflowmonitorNextToken) > 0 {
		input.NextToken = aws.String(_networkflowmonitorNextToken)
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

	var results []*networkflowmonitor.ListMonitorsOutput
	p := networkflowmonitor.NewListMonitorsPaginator(client, input)
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

// List all the scopes for an account.
func networkflowmonitor_ListScopes(cfg aws.Config, client *networkflowmonitor.Client) {
	input := &networkflowmonitor.ListScopesInput{}

	if len(_networkflowmonitorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _networkflowmonitorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_networkflowmonitorNextToken) > 0 {
		input.NextToken = aws.String(_networkflowmonitorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListScopes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*networkflowmonitor.ListScopesOutput
	p := networkflowmonitor.NewListScopesPaginator(client, input)
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

// Returns all the tags for a resource.
func networkflowmonitor_ListTagsForResource(cfg aws.Config, client *networkflowmonitor.Client) {
	input := &networkflowmonitor.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_networkflowmonitorResourceArn) > 0 {
		input.ResourceArn = aws.String(_networkflowmonitorResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a query that you can use with the Network Flow Monitor query interface
// to return the top contributors for a monitor. Specify the monitor that you want
// to create the query for.
//
// The call returns a query ID that you can use with [GetQueryResultsMonitorTopContributors] to run the query and return
// the top contributors for a specific monitor.
//
// Top contributors in Network Flow Monitor are network flows with the highest
// values for a specific metric type. Top contributors can be across all workload
// insights, for a given scope, or for a specific monitor. Use the applicable APIs
// for the top contributors that you want to be returned.
//
// [GetQueryResultsMonitorTopContributors]: https://docs.aws.amazon.com/networkflowmonitor/2.0/APIReference/API_GetQueryResultsMonitorTopContributors.html
func networkflowmonitor_StartQueryMonitorTopContributors(cfg aws.Config, client *networkflowmonitor.Client) {
	input := &networkflowmonitor.StartQueryMonitorTopContributorsInput{
		// DestinationCategory: types.DestinationCategory, // Required
		// EndTime: *time.Time, // Required
		// MetricName: types.MonitorMetric, // Required
		// MonitorName: *string, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_networkflowmonitorDestinationCategory) > 0 {
		if err := assignInputField(input, "DestinationCategory", _networkflowmonitorDestinationCategory); err != nil {
			log.Errorf("invalid --destination-category: %s", err.Error())
			return
		}
	}
	if len(_networkflowmonitorEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _networkflowmonitorEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_networkflowmonitorMetricName) > 0 {
		if err := assignInputField(input, "MetricName", _networkflowmonitorMetricName); err != nil {
			log.Errorf("invalid --metric-name: %s", err.Error())
			return
		}
	}
	if len(_networkflowmonitorMonitorName) > 0 {
		input.MonitorName = aws.String(_networkflowmonitorMonitorName)
	}
	if len(_networkflowmonitorStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _networkflowmonitorStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_networkflowmonitorLimit) > 0 {
		if err := assignInputField(input, "Limit", _networkflowmonitorLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartQueryMonitorTopContributors(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a query with the Network Flow Monitor query interface that you can run
// to return workload insights top contributors. Specify the scope that you want to
// create a query for.
//
// The call returns a query ID that you can use with [GetQueryResultsWorkloadInsightsTopContributors] to run the query and return
// the top contributors for the workload insights for a scope.
//
// Top contributors in Network Flow Monitor are network flows with the highest
// values for a specific metric type. Top contributors can be across all workload
// insights, for a given scope, or for a specific monitor. Use the applicable APIs
// for the top contributors that you want to be returned.
//
// [GetQueryResultsWorkloadInsightsTopContributors]: https://docs.aws.amazon.com/networkflowmonitor/2.0/APIReference/API_GetQueryResultsWorkloadInsightsTopContributors.html
func networkflowmonitor_StartQueryWorkloadInsightsTopContributors(cfg aws.Config, client *networkflowmonitor.Client) {
	input := &networkflowmonitor.StartQueryWorkloadInsightsTopContributorsInput{
		// DestinationCategory: types.DestinationCategory, // Required
		// EndTime: *time.Time, // Required
		// MetricName: types.WorkloadInsightsMetric, // Required
		// ScopeId: *string, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_networkflowmonitorDestinationCategory) > 0 {
		if err := assignInputField(input, "DestinationCategory", _networkflowmonitorDestinationCategory); err != nil {
			log.Errorf("invalid --destination-category: %s", err.Error())
			return
		}
	}
	if len(_networkflowmonitorEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _networkflowmonitorEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_networkflowmonitorMetricName) > 0 {
		if err := assignInputField(input, "MetricName", _networkflowmonitorMetricName); err != nil {
			log.Errorf("invalid --metric-name: %s", err.Error())
			return
		}
	}
	if len(_networkflowmonitorScopeId) > 0 {
		input.ScopeId = aws.String(_networkflowmonitorScopeId)
	}
	if len(_networkflowmonitorStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _networkflowmonitorStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_networkflowmonitorLimit) > 0 {
		if err := assignInputField(input, "Limit", _networkflowmonitorLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartQueryWorkloadInsightsTopContributors(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a query with the Network Flow Monitor query interface that you can run
// to return data for workload insights top contributors. Specify the scope that
// you want to create a query for.
//
// The call returns a query ID that you can use with [GetQueryResultsWorkloadInsightsTopContributorsData] to run the query and return
// the data for the top contributors for the workload insights for a scope.
//
// Top contributors in Network Flow Monitor are network flows with the highest
// values for a specific metric type. Top contributors can be across all workload
// insights, for a given scope, or for a specific monitor. Use the applicable call
// for the top contributors that you want to be returned.
//
// [GetQueryResultsWorkloadInsightsTopContributorsData]: https://docs.aws.amazon.com/networkflowmonitor/2.0/APIReference/API_GetQueryResultsWorkloadInsightsTopContributorsData.html
func networkflowmonitor_StartQueryWorkloadInsightsTopContributorsData(cfg aws.Config, client *networkflowmonitor.Client) {
	input := &networkflowmonitor.StartQueryWorkloadInsightsTopContributorsDataInput{
		// DestinationCategory: types.DestinationCategory, // Required
		// EndTime: *time.Time, // Required
		// MetricName: types.WorkloadInsightsMetric, // Required
		// ScopeId: *string, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_networkflowmonitorDestinationCategory) > 0 {
		if err := assignInputField(input, "DestinationCategory", _networkflowmonitorDestinationCategory); err != nil {
			log.Errorf("invalid --destination-category: %s", err.Error())
			return
		}
	}
	if len(_networkflowmonitorEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _networkflowmonitorEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_networkflowmonitorMetricName) > 0 {
		if err := assignInputField(input, "MetricName", _networkflowmonitorMetricName); err != nil {
			log.Errorf("invalid --metric-name: %s", err.Error())
			return
		}
	}
	if len(_networkflowmonitorScopeId) > 0 {
		input.ScopeId = aws.String(_networkflowmonitorScopeId)
	}
	if len(_networkflowmonitorStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _networkflowmonitorStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartQueryWorkloadInsightsTopContributorsData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stop a top contributors query for a monitor. Specify the query that you want to
// stop by providing a query ID and a monitor name.
//
// Top contributors in Network Flow Monitor are network flows with the highest
// values for a specific metric type. Top contributors can be across all workload
// insights, for a given scope, or for a specific monitor. Use the applicable call
// for the top contributors that you want to be returned.
func networkflowmonitor_StopQueryMonitorTopContributors(cfg aws.Config, client *networkflowmonitor.Client) {
	input := &networkflowmonitor.StopQueryMonitorTopContributorsInput{
		// MonitorName: *string, // Required
		// QueryId: *string, // Required
	}

	if len(_networkflowmonitorMonitorName) > 0 {
		input.MonitorName = aws.String(_networkflowmonitorMonitorName)
	}
	if len(_networkflowmonitorQueryId) > 0 {
		input.QueryId = aws.String(_networkflowmonitorQueryId)
	}

	if resp, err := client.StopQueryMonitorTopContributors(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stop a top contributors query for workload insights. Specify the query that you
// want to stop by providing a query ID and a scope ID.
//
// Top contributors in Network Flow Monitor are network flows with the highest
// values for a specific metric type. Top contributors can be across all workload
// insights, for a given scope, or for a specific monitor. Use the applicable call
// for the top contributors that you want to be returned.
func networkflowmonitor_StopQueryWorkloadInsightsTopContributors(cfg aws.Config, client *networkflowmonitor.Client) {
	input := &networkflowmonitor.StopQueryWorkloadInsightsTopContributorsInput{
		// QueryId: *string, // Required
		// ScopeId: *string, // Required
	}

	if len(_networkflowmonitorQueryId) > 0 {
		input.QueryId = aws.String(_networkflowmonitorQueryId)
	}
	if len(_networkflowmonitorScopeId) > 0 {
		input.ScopeId = aws.String(_networkflowmonitorScopeId)
	}

	if resp, err := client.StopQueryWorkloadInsightsTopContributors(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stop a top contributors data query for workload insights. Specify the query
// that you want to stop by providing a query ID and a scope ID.
//
// Top contributors in Network Flow Monitor are network flows with the highest
// values for a specific metric type. Top contributors can be across all workload
// insights, for a given scope, or for a specific monitor. Use the applicable call
// for the top contributors that you want to be returned.
func networkflowmonitor_StopQueryWorkloadInsightsTopContributorsData(cfg aws.Config, client *networkflowmonitor.Client) {
	input := &networkflowmonitor.StopQueryWorkloadInsightsTopContributorsDataInput{
		// QueryId: *string, // Required
		// ScopeId: *string, // Required
	}

	if len(_networkflowmonitorQueryId) > 0 {
		input.QueryId = aws.String(_networkflowmonitorQueryId)
	}
	if len(_networkflowmonitorScopeId) > 0 {
		input.ScopeId = aws.String(_networkflowmonitorScopeId)
	}

	if resp, err := client.StopQueryWorkloadInsightsTopContributorsData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a tag to a resource.
func networkflowmonitor_TagResource(cfg aws.Config, client *networkflowmonitor.Client) {
	input := &networkflowmonitor.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_networkflowmonitorResourceArn) > 0 {
		input.ResourceArn = aws.String(_networkflowmonitorResourceArn)
	}
	if len(_networkflowmonitorTags) > 0 {
		if err := assignInputField(input, "Tags", _networkflowmonitorTags); err != nil {
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
func networkflowmonitor_UntagResource(cfg aws.Config, client *networkflowmonitor.Client) {
	input := &networkflowmonitor.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_networkflowmonitorResourceArn) > 0 {
		input.ResourceArn = aws.String(_networkflowmonitorResourceArn)
	}
	if len(_networkflowmonitorTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _networkflowmonitorTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a monitor to add or remove local or remote resources.
func networkflowmonitor_UpdateMonitor(cfg aws.Config, client *networkflowmonitor.Client) {
	input := &networkflowmonitor.UpdateMonitorInput{
		// MonitorName: *string, // Required
	}

	if len(_networkflowmonitorMonitorName) > 0 {
		input.MonitorName = aws.String(_networkflowmonitorMonitorName)
	}
	if len(_networkflowmonitorClientToken) > 0 {
		input.ClientToken = aws.String(_networkflowmonitorClientToken)
	}
	if len(_networkflowmonitorLocalResourcesToAdd) > 0 {
		if err := assignInputField(input, "LocalResourcesToAdd", _networkflowmonitorLocalResourcesToAdd); err != nil {
			log.Errorf("invalid --local-resources-to-add: %s", err.Error())
			return
		}
	}
	if len(_networkflowmonitorLocalResourcesToRemove) > 0 {
		if err := assignInputField(input, "LocalResourcesToRemove", _networkflowmonitorLocalResourcesToRemove); err != nil {
			log.Errorf("invalid --local-resources-to-remove: %s", err.Error())
			return
		}
	}
	if len(_networkflowmonitorRemoteResourcesToAdd) > 0 {
		if err := assignInputField(input, "RemoteResourcesToAdd", _networkflowmonitorRemoteResourcesToAdd); err != nil {
			log.Errorf("invalid --remote-resources-to-add: %s", err.Error())
			return
		}
	}
	if len(_networkflowmonitorRemoteResourcesToRemove) > 0 {
		if err := assignInputField(input, "RemoteResourcesToRemove", _networkflowmonitorRemoteResourcesToRemove); err != nil {
			log.Errorf("invalid --remote-resources-to-remove: %s", err.Error())
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

// Update a scope to add or remove resources that you want to be available for
// Network Flow Monitor to generate metrics for, when you have active agents on
// those resources sending metrics reports to the Network Flow Monitor backend.
func networkflowmonitor_UpdateScope(cfg aws.Config, client *networkflowmonitor.Client) {
	input := &networkflowmonitor.UpdateScopeInput{
		// ScopeId: *string, // Required
	}

	if len(_networkflowmonitorScopeId) > 0 {
		input.ScopeId = aws.String(_networkflowmonitorScopeId)
	}
	if len(_networkflowmonitorResourcesToAdd) > 0 {
		if err := assignInputField(input, "ResourcesToAdd", _networkflowmonitorResourcesToAdd); err != nil {
			log.Errorf("invalid --resources-to-add: %s", err.Error())
			return
		}
	}
	if len(_networkflowmonitorResourcesToDelete) > 0 {
		if err := assignInputField(input, "ResourcesToDelete", _networkflowmonitorResourcesToDelete); err != nil {
			log.Errorf("invalid --resources-to-delete: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateScope(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_networkflowmonitorCmd)
	_networkflowmonitorCmd.Flags().SortFlags = false

	_networkflowmonitorCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_networkflowmonitorCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_networkflowmonitorCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_networkflowmonitorCmd.Flags().StringVarP(&_networkflowmonitorClientToken, "client-token", "", "", "Client Token")
	_networkflowmonitorCmd.Flags().StringVarP(&_networkflowmonitorDestinationCategory, "destination-category", "", "", "Destination Category")
	_networkflowmonitorCmd.Flags().StringVarP(&_networkflowmonitorEndTime, "end-time", "", "", "End Time")
	_networkflowmonitorCmd.Flags().StringVarP(&_networkflowmonitorLimit, "limit", "", "", "Limit")
	_networkflowmonitorCmd.Flags().StringVarP(&_networkflowmonitorLocalResources, "local-resources", "", "", "Local Resources")
	_networkflowmonitorCmd.Flags().StringVarP(&_networkflowmonitorLocalResourcesToAdd, "local-resources-to-add", "", "", "Local Resources To Add")
	_networkflowmonitorCmd.Flags().StringVarP(&_networkflowmonitorLocalResourcesToRemove, "local-resources-to-remove", "", "", "Local Resources To Remove")
	_networkflowmonitorCmd.Flags().StringVarP(&_networkflowmonitorMaxResults, "max-results", "", "", "Max Results")
	_networkflowmonitorCmd.Flags().StringVarP(&_networkflowmonitorMetricName, "metric-name", "", "", "Metric Name")
	_networkflowmonitorCmd.Flags().StringVarP(&_networkflowmonitorMonitorName, "monitor-name", "", "", "Monitor Name")
	_networkflowmonitorCmd.Flags().StringVarP(&_networkflowmonitorMonitorStatus, "monitor-status", "", "", "Monitor Status")
	_networkflowmonitorCmd.Flags().StringVarP(&_networkflowmonitorNextToken, "next-token", "", "", "Next Token")
	_networkflowmonitorCmd.Flags().StringVarP(&_networkflowmonitorQueryId, "query-id", "", "", "Query ID")
	_networkflowmonitorCmd.Flags().StringVarP(&_networkflowmonitorRemoteResources, "remote-resources", "", "", "Remote Resources")
	_networkflowmonitorCmd.Flags().StringVarP(&_networkflowmonitorRemoteResourcesToAdd, "remote-resources-to-add", "", "", "Remote Resources To Add")
	_networkflowmonitorCmd.Flags().StringVarP(&_networkflowmonitorRemoteResourcesToRemove, "remote-resources-to-remove", "", "", "Remote Resources To Remove")
	_networkflowmonitorCmd.Flags().StringVarP(&_networkflowmonitorResourceArn, "resource-arn", "", "", "Resource ARN")
	_networkflowmonitorCmd.Flags().StringVarP(&_networkflowmonitorResourcesToAdd, "resources-to-add", "", "", "Resources To Add")
	_networkflowmonitorCmd.Flags().StringVarP(&_networkflowmonitorResourcesToDelete, "resources-to-delete", "", "", "Resources To Delete")
	_networkflowmonitorCmd.Flags().StringVarP(&_networkflowmonitorScopeArn, "scope-arn", "", "", "Scope ARN")
	_networkflowmonitorCmd.Flags().StringVarP(&_networkflowmonitorScopeId, "scope-id", "", "", "Scope ID")
	_networkflowmonitorCmd.Flags().StringVarP(&_networkflowmonitorStartTime, "start-time", "", "", "Start Time")
	_networkflowmonitorCmd.Flags().StringSliceVarP(&_networkflowmonitorTagKeys, "tag-keys", "", nil, "Tag Keys")
	_networkflowmonitorCmd.Flags().StringVarP(&_networkflowmonitorTags, "tags", "", "", "Tags")
	_networkflowmonitorCmd.Flags().StringVarP(&_networkflowmonitorTargets, "targets", "", "", "Targets")

	_networkflowmonitorCmd.Flags().BoolVarP(&_networkflowmonitorCreateMonitor, "create-monitor", "", false, "Create Monitor")
	_networkflowmonitorCmd.Flags().BoolVarP(&_networkflowmonitorCreateScope, "create-scope", "", false, "Create Scope")
	_networkflowmonitorCmd.Flags().BoolVarP(&_networkflowmonitorDeleteMonitor, "delete-monitor", "", false, "Delete Monitor")
	_networkflowmonitorCmd.Flags().BoolVarP(&_networkflowmonitorDeleteScope, "delete-scope", "", false, "Delete Scope")
	_networkflowmonitorCmd.Flags().BoolVarP(&_networkflowmonitorGetMonitor, "get-monitor", "", false, "Get Monitor")
	_networkflowmonitorCmd.Flags().BoolVarP(&_networkflowmonitorGetQueryResultsMonitorTopContributors, "get-query-results-monitor-top-contributors", "", false, "Get Query Results Monitor Top Contributors")
	_networkflowmonitorCmd.Flags().BoolVarP(&_networkflowmonitorGetQueryResultsWorkloadInsightsTopContributors, "get-query-results-workload-insights-top-contributors", "", false, "Get Query Results Workload Insights Top Contributors")
	_networkflowmonitorCmd.Flags().BoolVarP(&_networkflowmonitorGetQueryResultsWorkloadInsightsTopContributorsData, "get-query-results-workload-insights-top-contributors-data", "", false, "Get Query Results Workload Insights Top Contributors Data")
	_networkflowmonitorCmd.Flags().BoolVarP(&_networkflowmonitorGetQueryStatusMonitorTopContributors, "get-query-status-monitor-top-contributors", "", false, "Get Query Status Monitor Top Contributors")
	_networkflowmonitorCmd.Flags().BoolVarP(&_networkflowmonitorGetQueryStatusWorkloadInsightsTopContributors, "get-query-status-workload-insights-top-contributors", "", false, "Get Query Status Workload Insights Top Contributors")
	_networkflowmonitorCmd.Flags().BoolVarP(&_networkflowmonitorGetQueryStatusWorkloadInsightsTopContributorsData, "get-query-status-workload-insights-top-contributors-data", "", false, "Get Query Status Workload Insights Top Contributors Data")
	_networkflowmonitorCmd.Flags().BoolVarP(&_networkflowmonitorGetScope, "get-scope", "", false, "Get Scope")
	_networkflowmonitorCmd.Flags().BoolVarP(&_networkflowmonitorListMonitors, "list-monitors", "", false, "List Monitors")
	_networkflowmonitorCmd.Flags().BoolVarP(&_networkflowmonitorListScopes, "list-scopes", "", false, "List Scopes")
	_networkflowmonitorCmd.Flags().BoolVarP(&_networkflowmonitorListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_networkflowmonitorCmd.Flags().BoolVarP(&_networkflowmonitorStartQueryMonitorTopContributors, "start-query-monitor-top-contributors", "", false, "Start Query Monitor Top Contributors")
	_networkflowmonitorCmd.Flags().BoolVarP(&_networkflowmonitorStartQueryWorkloadInsightsTopContributors, "start-query-workload-insights-top-contributors", "", false, "Start Query Workload Insights Top Contributors")
	_networkflowmonitorCmd.Flags().BoolVarP(&_networkflowmonitorStartQueryWorkloadInsightsTopContributorsData, "start-query-workload-insights-top-contributors-data", "", false, "Start Query Workload Insights Top Contributors Data")
	_networkflowmonitorCmd.Flags().BoolVarP(&_networkflowmonitorStopQueryMonitorTopContributors, "stop-query-monitor-top-contributors", "", false, "Stop Query Monitor Top Contributors")
	_networkflowmonitorCmd.Flags().BoolVarP(&_networkflowmonitorStopQueryWorkloadInsightsTopContributors, "stop-query-workload-insights-top-contributors", "", false, "Stop Query Workload Insights Top Contributors")
	_networkflowmonitorCmd.Flags().BoolVarP(&_networkflowmonitorStopQueryWorkloadInsightsTopContributorsData, "stop-query-workload-insights-top-contributors-data", "", false, "Stop Query Workload Insights Top Contributors Data")
	_networkflowmonitorCmd.Flags().BoolVarP(&_networkflowmonitorTagResource, "tag-resource", "", false, "Tag Resource")
	_networkflowmonitorCmd.Flags().BoolVarP(&_networkflowmonitorUntagResource, "untag-resource", "", false, "Untag Resource")
	_networkflowmonitorCmd.Flags().BoolVarP(&_networkflowmonitorUpdateMonitor, "update-monitor", "", false, "Update Monitor")
	_networkflowmonitorCmd.Flags().BoolVarP(&_networkflowmonitorUpdateScope, "update-scope", "", false, "Update Scope")

}
