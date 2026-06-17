package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/simspaceweaver"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// simspaceweaverCmd represents the simspaceweaver command
var _simspaceweaverCmd = &cobra.Command{
	Use:   "simspaceweaver",
	Short: "AWS simspaceweaver CLI",
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
		client := simspaceweaver.NewFromConfig(cfg)
		if _simspaceweaverCreateSnapshot {
			simspaceweaver_CreateSnapshot(cfg, client)
			return
		}
		if _simspaceweaverDeleteApp {
			simspaceweaver_DeleteApp(cfg, client)
			return
		}
		if _simspaceweaverDeleteSimulation {
			simspaceweaver_DeleteSimulation(cfg, client)
			return
		}
		if _simspaceweaverDescribeApp {
			simspaceweaver_DescribeApp(cfg, client)
			return
		}
		if _simspaceweaverDescribeSimulation {
			simspaceweaver_DescribeSimulation(cfg, client)
			return
		}
		if _simspaceweaverListApps {
			simspaceweaver_ListApps(cfg, client)
			return
		}
		if _simspaceweaverListSimulations {
			simspaceweaver_ListSimulations(cfg, client)
			return
		}
		if _simspaceweaverListTagsForResource {
			simspaceweaver_ListTagsForResource(cfg, client)
			return
		}
		if _simspaceweaverStartApp {
			simspaceweaver_StartApp(cfg, client)
			return
		}
		if _simspaceweaverStartClock {
			simspaceweaver_StartClock(cfg, client)
			return
		}
		if _simspaceweaverStartSimulation {
			simspaceweaver_StartSimulation(cfg, client)
			return
		}
		if _simspaceweaverStopApp {
			simspaceweaver_StopApp(cfg, client)
			return
		}
		if _simspaceweaverStopClock {
			simspaceweaver_StopClock(cfg, client)
			return
		}
		if _simspaceweaverStopSimulation {
			simspaceweaver_StopSimulation(cfg, client)
			return
		}
		if _simspaceweaverTagResource {
			simspaceweaver_TagResource(cfg, client)
			return
		}
		if _simspaceweaverUntagResource {
			simspaceweaver_UntagResource(cfg, client)
			return
		}

	},
}

var (
	_simspaceweaverCreateSnapshot      bool
	_simspaceweaverDeleteApp           bool
	_simspaceweaverDeleteSimulation    bool
	_simspaceweaverDescribeApp         bool
	_simspaceweaverDescribeSimulation  bool
	_simspaceweaverListApps            bool
	_simspaceweaverListSimulations     bool
	_simspaceweaverListTagsForResource bool
	_simspaceweaverStartApp            bool
	_simspaceweaverStartClock          bool
	_simspaceweaverStartSimulation     bool
	_simspaceweaverStopApp             bool
	_simspaceweaverStopClock           bool
	_simspaceweaverStopSimulation      bool
	_simspaceweaverTagResource         bool
	_simspaceweaverUntagResource       bool

	_simspaceweaverApp                string
	_simspaceweaverClientToken        string
	_simspaceweaverDescription        string
	_simspaceweaverDestination        string
	_simspaceweaverDomain             string
	_simspaceweaverLaunchOverrides    string
	_simspaceweaverMaxResults         string
	_simspaceweaverMaximumDuration    string
	_simspaceweaverName               string
	_simspaceweaverNextToken          string
	_simspaceweaverResourceArn        string
	_simspaceweaverRoleArn            string
	_simspaceweaverSchemaS3Location   string
	_simspaceweaverSimulation         string
	_simspaceweaverSnapshotS3Location string
	_simspaceweaverTagKeys            []string
	_simspaceweaverTags               string
)

// Creates a snapshot of the specified simulation. A snapshot is a file that
// contains simulation state data at a specific time. The state data saved in a
// snapshot includes entity data from the State Fabric, the simulation
// configuration specified in the schema, and the clock tick number. You can use
// the snapshot to initialize a new simulation. For more information about
// snapshots, see [Snapshots]in the SimSpace Weaver User Guide.
//
// You specify a Destination when you create a snapshot. The Destination is the
// name of an Amazon S3 bucket and an optional ObjectKeyPrefix . The
// ObjectKeyPrefix is usually the name of a folder in the bucket. SimSpace Weaver
// creates a snapshot folder inside the Destination and places the snapshot file
// there.
//
// The snapshot file is an Amazon S3 object. It has an object key with the form:
// object-key-prefix/snapshot/simulation-name-YYMMdd-HHmm-ss.zip , where:
//
// - YY is the 2-digit year
//
// - MM is the 2-digit month
//
// - dd is the 2-digit day of the month
//
// - HH is the 2-digit hour (24-hour clock)
//
// - mm is the 2-digit minutes
//
// - ss is the 2-digit seconds
//
// [Snapshots]: https://docs.aws.amazon.com/simspaceweaver/latest/userguide/working-with_snapshots.html
func simspaceweaver_CreateSnapshot(cfg aws.Config, client *simspaceweaver.Client) {
	input := &simspaceweaver.CreateSnapshotInput{
		// Destination: *types.S3Destination, // Required
		// Simulation: *string, // Required
	}

	if len(_simspaceweaverDestination) > 0 {
		if err := assignInputField(input, "Destination", _simspaceweaverDestination); err != nil {
			log.Errorf("invalid --destination: %s", err.Error())
			return
		}
	}
	if len(_simspaceweaverSimulation) > 0 {
		input.Simulation = aws.String(_simspaceweaverSimulation)
	}

	if resp, err := client.CreateSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the instance of the given custom app.
func simspaceweaver_DeleteApp(cfg aws.Config, client *simspaceweaver.Client) {
	input := &simspaceweaver.DeleteAppInput{
		// App: *string, // Required
		// Domain: *string, // Required
		// Simulation: *string, // Required
	}

	if len(_simspaceweaverApp) > 0 {
		input.App = aws.String(_simspaceweaverApp)
	}
	if len(_simspaceweaverDomain) > 0 {
		input.Domain = aws.String(_simspaceweaverDomain)
	}
	if len(_simspaceweaverSimulation) > 0 {
		input.Simulation = aws.String(_simspaceweaverSimulation)
	}

	if resp, err := client.DeleteApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes all SimSpace Weaver resources assigned to the given simulation.
// Your simulation uses resources in other Amazon Web Services. This API operation
// doesn't delete resources in other Amazon Web Services.
func simspaceweaver_DeleteSimulation(cfg aws.Config, client *simspaceweaver.Client) {
	input := &simspaceweaver.DeleteSimulationInput{
		// Simulation: *string, // Required
	}

	if len(_simspaceweaverSimulation) > 0 {
		input.Simulation = aws.String(_simspaceweaverSimulation)
	}

	if resp, err := client.DeleteSimulation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the state of the given custom app.
func simspaceweaver_DescribeApp(cfg aws.Config, client *simspaceweaver.Client) {
	input := &simspaceweaver.DescribeAppInput{
		// App: *string, // Required
		// Domain: *string, // Required
		// Simulation: *string, // Required
	}

	if len(_simspaceweaverApp) > 0 {
		input.App = aws.String(_simspaceweaverApp)
	}
	if len(_simspaceweaverDomain) > 0 {
		input.Domain = aws.String(_simspaceweaverDomain)
	}
	if len(_simspaceweaverSimulation) > 0 {
		input.Simulation = aws.String(_simspaceweaverSimulation)
	}

	if resp, err := client.DescribeApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the current state of the given simulation.
func simspaceweaver_DescribeSimulation(cfg aws.Config, client *simspaceweaver.Client) {
	input := &simspaceweaver.DescribeSimulationInput{
		// Simulation: *string, // Required
	}

	if len(_simspaceweaverSimulation) > 0 {
		input.Simulation = aws.String(_simspaceweaverSimulation)
	}

	if resp, err := client.DescribeSimulation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all custom apps or service apps for the given simulation and domain.
func simspaceweaver_ListApps(cfg aws.Config, client *simspaceweaver.Client) {
	input := &simspaceweaver.ListAppsInput{
		// Simulation: *string, // Required
	}

	if len(_simspaceweaverSimulation) > 0 {
		input.Simulation = aws.String(_simspaceweaverSimulation)
	}
	if len(_simspaceweaverDomain) > 0 {
		input.Domain = aws.String(_simspaceweaverDomain)
	}
	if len(_simspaceweaverMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _simspaceweaverMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_simspaceweaverNextToken) > 0 {
		input.NextToken = aws.String(_simspaceweaverNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApps(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*simspaceweaver.ListAppsOutput
	p := simspaceweaver.NewListAppsPaginator(client, input)
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

// Lists the SimSpace Weaver simulations in the Amazon Web Services account used
// to make the API call.
func simspaceweaver_ListSimulations(cfg aws.Config, client *simspaceweaver.Client) {
	input := &simspaceweaver.ListSimulationsInput{}

	if len(_simspaceweaverMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _simspaceweaverMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_simspaceweaverNextToken) > 0 {
		input.NextToken = aws.String(_simspaceweaverNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSimulations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*simspaceweaver.ListSimulationsOutput
	p := simspaceweaver.NewListSimulationsPaginator(client, input)
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

// Lists all tags on a SimSpace Weaver resource.
func simspaceweaver_ListTagsForResource(cfg aws.Config, client *simspaceweaver.Client) {
	input := &simspaceweaver.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_simspaceweaverResourceArn) > 0 {
		input.ResourceArn = aws.String(_simspaceweaverResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a custom app with the configuration specified in the simulation schema.
func simspaceweaver_StartApp(cfg aws.Config, client *simspaceweaver.Client) {
	input := &simspaceweaver.StartAppInput{
		// Domain: *string, // Required
		// Name: *string, // Required
		// Simulation: *string, // Required
	}

	if len(_simspaceweaverDomain) > 0 {
		input.Domain = aws.String(_simspaceweaverDomain)
	}
	if len(_simspaceweaverName) > 0 {
		input.Name = aws.String(_simspaceweaverName)
	}
	if len(_simspaceweaverSimulation) > 0 {
		input.Simulation = aws.String(_simspaceweaverSimulation)
	}
	if len(_simspaceweaverClientToken) > 0 {
		input.ClientToken = aws.String(_simspaceweaverClientToken)
	}
	if len(_simspaceweaverDescription) > 0 {
		input.Description = aws.String(_simspaceweaverDescription)
	}
	if len(_simspaceweaverLaunchOverrides) > 0 {
		if err := assignInputField(input, "LaunchOverrides", _simspaceweaverLaunchOverrides); err != nil {
			log.Errorf("invalid --launch-overrides: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the simulation clock.
func simspaceweaver_StartClock(cfg aws.Config, client *simspaceweaver.Client) {
	input := &simspaceweaver.StartClockInput{
		// Simulation: *string, // Required
	}

	if len(_simspaceweaverSimulation) > 0 {
		input.Simulation = aws.String(_simspaceweaverSimulation)
	}

	if resp, err := client.StartClock(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a simulation with the given name. You must choose to start your
// simulation from a schema or from a snapshot. For more information about the
// schema, see the [schema reference]in the SimSpace Weaver User Guide. For more information about
// snapshots, see [Snapshots]in the SimSpace Weaver User Guide.
//
// [schema reference]: https://docs.aws.amazon.com/simspaceweaver/latest/userguide/schema-reference.html
// [Snapshots]: https://docs.aws.amazon.com/simspaceweaver/latest/userguide/working-with_snapshots.html
func simspaceweaver_StartSimulation(cfg aws.Config, client *simspaceweaver.Client) {
	input := &simspaceweaver.StartSimulationInput{
		// Name: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_simspaceweaverName) > 0 {
		input.Name = aws.String(_simspaceweaverName)
	}
	if len(_simspaceweaverRoleArn) > 0 {
		input.RoleArn = aws.String(_simspaceweaverRoleArn)
	}
	if len(_simspaceweaverClientToken) > 0 {
		input.ClientToken = aws.String(_simspaceweaverClientToken)
	}
	if len(_simspaceweaverDescription) > 0 {
		input.Description = aws.String(_simspaceweaverDescription)
	}
	if len(_simspaceweaverMaximumDuration) > 0 {
		input.MaximumDuration = aws.String(_simspaceweaverMaximumDuration)
	}
	if len(_simspaceweaverSchemaS3Location) > 0 {
		if err := assignInputField(input, "SchemaS3Location", _simspaceweaverSchemaS3Location); err != nil {
			log.Errorf("invalid --schema-s3-location: %s", err.Error())
			return
		}
	}
	if len(_simspaceweaverSnapshotS3Location) > 0 {
		if err := assignInputField(input, "SnapshotS3Location", _simspaceweaverSnapshotS3Location); err != nil {
			log.Errorf("invalid --snapshot-s3-location: %s", err.Error())
			return
		}
	}
	if len(_simspaceweaverTags) > 0 {
		if err := assignInputField(input, "Tags", _simspaceweaverTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartSimulation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the given custom app and shuts down all of its allocated compute
// resources.
func simspaceweaver_StopApp(cfg aws.Config, client *simspaceweaver.Client) {
	input := &simspaceweaver.StopAppInput{
		// App: *string, // Required
		// Domain: *string, // Required
		// Simulation: *string, // Required
	}

	if len(_simspaceweaverApp) > 0 {
		input.App = aws.String(_simspaceweaverApp)
	}
	if len(_simspaceweaverDomain) > 0 {
		input.Domain = aws.String(_simspaceweaverDomain)
	}
	if len(_simspaceweaverSimulation) > 0 {
		input.Simulation = aws.String(_simspaceweaverSimulation)
	}

	if resp, err := client.StopApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the simulation clock.
func simspaceweaver_StopClock(cfg aws.Config, client *simspaceweaver.Client) {
	input := &simspaceweaver.StopClockInput{
		// Simulation: *string, // Required
	}

	if len(_simspaceweaverSimulation) > 0 {
		input.Simulation = aws.String(_simspaceweaverSimulation)
	}

	if resp, err := client.StopClock(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the given simulation.
// You can't restart a simulation after you stop it. If you want to restart a
// simulation, then you must stop it, delete it, and start a new instance of it.
func simspaceweaver_StopSimulation(cfg aws.Config, client *simspaceweaver.Client) {
	input := &simspaceweaver.StopSimulationInput{
		// Simulation: *string, // Required
	}

	if len(_simspaceweaverSimulation) > 0 {
		input.Simulation = aws.String(_simspaceweaverSimulation)
	}

	if resp, err := client.StopSimulation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds tags to a SimSpace Weaver resource. For more information about tags, see [Tagging Amazon Web Services resources]
// in the Amazon Web Services General Reference.
//
// [Tagging Amazon Web Services resources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
func simspaceweaver_TagResource(cfg aws.Config, client *simspaceweaver.Client) {
	input := &simspaceweaver.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_simspaceweaverResourceArn) > 0 {
		input.ResourceArn = aws.String(_simspaceweaverResourceArn)
	}
	if len(_simspaceweaverTags) > 0 {
		if err := assignInputField(input, "Tags", _simspaceweaverTags); err != nil {
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

// Removes tags from a SimSpace Weaver resource. For more information about tags,
// see [Tagging Amazon Web Services resources]in the Amazon Web Services General Reference.
//
// [Tagging Amazon Web Services resources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
func simspaceweaver_UntagResource(cfg aws.Config, client *simspaceweaver.Client) {
	input := &simspaceweaver.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_simspaceweaverResourceArn) > 0 {
		input.ResourceArn = aws.String(_simspaceweaverResourceArn)
	}
	if len(_simspaceweaverTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _simspaceweaverTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_simspaceweaverCmd)
	_simspaceweaverCmd.Flags().SortFlags = false

	_simspaceweaverCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_simspaceweaverCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_simspaceweaverCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_simspaceweaverCmd.Flags().StringVarP(&_simspaceweaverApp, "app", "", "", "App")
	_simspaceweaverCmd.Flags().StringVarP(&_simspaceweaverClientToken, "client-token", "", "", "Client Token")
	_simspaceweaverCmd.Flags().StringVarP(&_simspaceweaverDescription, "description", "", "", "Description")
	_simspaceweaverCmd.Flags().StringVarP(&_simspaceweaverDestination, "destination", "", "", "Destination")
	_simspaceweaverCmd.Flags().StringVarP(&_simspaceweaverDomain, "domain", "", "", "Domain")
	_simspaceweaverCmd.Flags().StringVarP(&_simspaceweaverLaunchOverrides, "launch-overrides", "", "", "Launch Overrides")
	_simspaceweaverCmd.Flags().StringVarP(&_simspaceweaverMaxResults, "max-results", "", "", "Max Results")
	_simspaceweaverCmd.Flags().StringVarP(&_simspaceweaverMaximumDuration, "maximum-duration", "", "", "Maximum Duration")
	_simspaceweaverCmd.Flags().StringVarP(&_simspaceweaverName, "name", "", "", "Name")
	_simspaceweaverCmd.Flags().StringVarP(&_simspaceweaverNextToken, "next-token", "", "", "Next Token")
	_simspaceweaverCmd.Flags().StringVarP(&_simspaceweaverResourceArn, "resource-arn", "", "", "Resource ARN")
	_simspaceweaverCmd.Flags().StringVarP(&_simspaceweaverRoleArn, "role-arn", "", "", "Role ARN")
	_simspaceweaverCmd.Flags().StringVarP(&_simspaceweaverSchemaS3Location, "schema-s3-location", "", "", "Schema S3 Location")
	_simspaceweaverCmd.Flags().StringVarP(&_simspaceweaverSimulation, "simulation", "", "", "Simulation")
	_simspaceweaverCmd.Flags().StringVarP(&_simspaceweaverSnapshotS3Location, "snapshot-s3-location", "", "", "Snapshot S3 Location")
	_simspaceweaverCmd.Flags().StringSliceVarP(&_simspaceweaverTagKeys, "tag-keys", "", nil, "Tag Keys")
	_simspaceweaverCmd.Flags().StringVarP(&_simspaceweaverTags, "tags", "", "", "Tags")

	_simspaceweaverCmd.Flags().BoolVarP(&_simspaceweaverCreateSnapshot, "create-snapshot", "", false, "Create Snapshot")
	_simspaceweaverCmd.Flags().BoolVarP(&_simspaceweaverDeleteApp, "delete-app", "", false, "Delete App")
	_simspaceweaverCmd.Flags().BoolVarP(&_simspaceweaverDeleteSimulation, "delete-simulation", "", false, "Delete Simulation")
	_simspaceweaverCmd.Flags().BoolVarP(&_simspaceweaverDescribeApp, "describe-app", "", false, "Describe App")
	_simspaceweaverCmd.Flags().BoolVarP(&_simspaceweaverDescribeSimulation, "describe-simulation", "", false, "Describe Simulation")
	_simspaceweaverCmd.Flags().BoolVarP(&_simspaceweaverListApps, "list-apps", "", false, "List Apps")
	_simspaceweaverCmd.Flags().BoolVarP(&_simspaceweaverListSimulations, "list-simulations", "", false, "List Simulations")
	_simspaceweaverCmd.Flags().BoolVarP(&_simspaceweaverListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_simspaceweaverCmd.Flags().BoolVarP(&_simspaceweaverStartApp, "start-app", "", false, "Start App")
	_simspaceweaverCmd.Flags().BoolVarP(&_simspaceweaverStartClock, "start-clock", "", false, "Start Clock")
	_simspaceweaverCmd.Flags().BoolVarP(&_simspaceweaverStartSimulation, "start-simulation", "", false, "Start Simulation")
	_simspaceweaverCmd.Flags().BoolVarP(&_simspaceweaverStopApp, "stop-app", "", false, "Stop App")
	_simspaceweaverCmd.Flags().BoolVarP(&_simspaceweaverStopClock, "stop-clock", "", false, "Stop Clock")
	_simspaceweaverCmd.Flags().BoolVarP(&_simspaceweaverStopSimulation, "stop-simulation", "", false, "Stop Simulation")
	_simspaceweaverCmd.Flags().BoolVarP(&_simspaceweaverTagResource, "tag-resource", "", false, "Tag Resource")
	_simspaceweaverCmd.Flags().BoolVarP(&_simspaceweaverUntagResource, "untag-resource", "", false, "Untag Resource")

}
