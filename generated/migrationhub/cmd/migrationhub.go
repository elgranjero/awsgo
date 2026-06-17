package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// migrationhubCmd represents the migrationhub command
var _migrationhubCmd = &cobra.Command{
	Use:   "migrationhub",
	Short: "AWS migrationhub CLI",
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
		client := migrationhub.NewFromConfig(cfg)
		if _migrationhubAssociateCreatedArtifact {
			migrationhub_AssociateCreatedArtifact(cfg, client)
			return
		}
		if _migrationhubAssociateDiscoveredResource {
			migrationhub_AssociateDiscoveredResource(cfg, client)
			return
		}
		if _migrationhubAssociateSourceResource {
			migrationhub_AssociateSourceResource(cfg, client)
			return
		}
		if _migrationhubCreateProgressUpdateStream {
			migrationhub_CreateProgressUpdateStream(cfg, client)
			return
		}
		if _migrationhubDeleteProgressUpdateStream {
			migrationhub_DeleteProgressUpdateStream(cfg, client)
			return
		}
		if _migrationhubDescribeApplicationState {
			migrationhub_DescribeApplicationState(cfg, client)
			return
		}
		if _migrationhubDescribeMigrationTask {
			migrationhub_DescribeMigrationTask(cfg, client)
			return
		}
		if _migrationhubDisassociateCreatedArtifact {
			migrationhub_DisassociateCreatedArtifact(cfg, client)
			return
		}
		if _migrationhubDisassociateDiscoveredResource {
			migrationhub_DisassociateDiscoveredResource(cfg, client)
			return
		}
		if _migrationhubDisassociateSourceResource {
			migrationhub_DisassociateSourceResource(cfg, client)
			return
		}
		if _migrationhubImportMigrationTask {
			migrationhub_ImportMigrationTask(cfg, client)
			return
		}
		if _migrationhubListApplicationStates {
			migrationhub_ListApplicationStates(cfg, client)
			return
		}
		if _migrationhubListCreatedArtifacts {
			migrationhub_ListCreatedArtifacts(cfg, client)
			return
		}
		if _migrationhubListDiscoveredResources {
			migrationhub_ListDiscoveredResources(cfg, client)
			return
		}
		if _migrationhubListMigrationTaskUpdates {
			migrationhub_ListMigrationTaskUpdates(cfg, client)
			return
		}
		if _migrationhubListMigrationTasks {
			migrationhub_ListMigrationTasks(cfg, client)
			return
		}
		if _migrationhubListProgressUpdateStreams {
			migrationhub_ListProgressUpdateStreams(cfg, client)
			return
		}
		if _migrationhubListSourceResources {
			migrationhub_ListSourceResources(cfg, client)
			return
		}
		if _migrationhubNotifyApplicationState {
			migrationhub_NotifyApplicationState(cfg, client)
			return
		}
		if _migrationhubNotifyMigrationTaskState {
			migrationhub_NotifyMigrationTaskState(cfg, client)
			return
		}
		if _migrationhubPutResourceAttributes {
			migrationhub_PutResourceAttributes(cfg, client)
			return
		}

	},
}

var (
	_migrationhubAssociateCreatedArtifact       bool
	_migrationhubAssociateDiscoveredResource    bool
	_migrationhubAssociateSourceResource        bool
	_migrationhubCreateProgressUpdateStream     bool
	_migrationhubDeleteProgressUpdateStream     bool
	_migrationhubDescribeApplicationState       bool
	_migrationhubDescribeMigrationTask          bool
	_migrationhubDisassociateCreatedArtifact    bool
	_migrationhubDisassociateDiscoveredResource bool
	_migrationhubDisassociateSourceResource     bool
	_migrationhubImportMigrationTask            bool
	_migrationhubListApplicationStates          bool
	_migrationhubListCreatedArtifacts           bool
	_migrationhubListDiscoveredResources        bool
	_migrationhubListMigrationTaskUpdates       bool
	_migrationhubListMigrationTasks             bool
	_migrationhubListProgressUpdateStreams      bool
	_migrationhubListSourceResources            bool
	_migrationhubNotifyApplicationState         bool
	_migrationhubNotifyMigrationTaskState       bool
	_migrationhubPutResourceAttributes          bool

	_migrationhubApplicationId            string
	_migrationhubApplicationIds           []string
	_migrationhubConfigurationId          string
	_migrationhubCreatedArtifact          string
	_migrationhubCreatedArtifactName      string
	_migrationhubDiscoveredResource       string
	_migrationhubDryRun                   string
	_migrationhubMaxResults               string
	_migrationhubMigrationTaskName        string
	_migrationhubNextToken                string
	_migrationhubNextUpdateSeconds        string
	_migrationhubProgressUpdateStream     string
	_migrationhubProgressUpdateStreamName string
	_migrationhubResourceAttributeList    string
	_migrationhubResourceName             string
	_migrationhubSourceResource           string
	_migrationhubSourceResourceName       string
	_migrationhubStatus                   string
	_migrationhubTask                     string
	_migrationhubUpdateDateTime           string
)

// Associates a created artifact of an AWS cloud resource, the target receiving
// the migration, with the migration task performed by a migration tool. This API
// has the following traits:
//
// - Migration tools can call the AssociateCreatedArtifact operation to indicate
// which AWS artifact is associated with a migration task.
//
// - The created artifact name must be provided in ARN (Amazon Resource Name)
// format which will contain information about type and region; for example:
// arn:aws:ec2:us-east-1:488216288981:image/ami-6d0ba87b .
//
// - Examples of the AWS resource behind the created artifact are, AMI's, EC2
// instance, or DMS endpoint, etc.
func migrationhub_AssociateCreatedArtifact(cfg aws.Config, client *migrationhub.Client) {
	input := &migrationhub.AssociateCreatedArtifactInput{
		// CreatedArtifact: *types.CreatedArtifact, // Required
		// MigrationTaskName: *string, // Required
		// ProgressUpdateStream: *string, // Required
	}

	if len(_migrationhubCreatedArtifact) > 0 {
		if err := assignInputField(input, "CreatedArtifact", _migrationhubCreatedArtifact); err != nil {
			log.Errorf("invalid --created-artifact: %s", err.Error())
			return
		}
	}
	if len(_migrationhubMigrationTaskName) > 0 {
		input.MigrationTaskName = aws.String(_migrationhubMigrationTaskName)
	}
	if len(_migrationhubProgressUpdateStream) > 0 {
		input.ProgressUpdateStream = aws.String(_migrationhubProgressUpdateStream)
	}
	if len(_migrationhubDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _migrationhubDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateCreatedArtifact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a discovered resource ID from Application Discovery Service with a
// migration task.
func migrationhub_AssociateDiscoveredResource(cfg aws.Config, client *migrationhub.Client) {
	input := &migrationhub.AssociateDiscoveredResourceInput{
		// DiscoveredResource: *types.DiscoveredResource, // Required
		// MigrationTaskName: *string, // Required
		// ProgressUpdateStream: *string, // Required
	}

	if len(_migrationhubDiscoveredResource) > 0 {
		if err := assignInputField(input, "DiscoveredResource", _migrationhubDiscoveredResource); err != nil {
			log.Errorf("invalid --discovered-resource: %s", err.Error())
			return
		}
	}
	if len(_migrationhubMigrationTaskName) > 0 {
		input.MigrationTaskName = aws.String(_migrationhubMigrationTaskName)
	}
	if len(_migrationhubProgressUpdateStream) > 0 {
		input.ProgressUpdateStream = aws.String(_migrationhubProgressUpdateStream)
	}
	if len(_migrationhubDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _migrationhubDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateDiscoveredResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a source resource with a migration task. For example, the source
// resource can be a source server, an application, or a migration wave.
func migrationhub_AssociateSourceResource(cfg aws.Config, client *migrationhub.Client) {
	input := &migrationhub.AssociateSourceResourceInput{
		// MigrationTaskName: *string, // Required
		// ProgressUpdateStream: *string, // Required
		// SourceResource: *types.SourceResource, // Required
	}

	if len(_migrationhubMigrationTaskName) > 0 {
		input.MigrationTaskName = aws.String(_migrationhubMigrationTaskName)
	}
	if len(_migrationhubProgressUpdateStream) > 0 {
		input.ProgressUpdateStream = aws.String(_migrationhubProgressUpdateStream)
	}
	if len(_migrationhubSourceResource) > 0 {
		if err := assignInputField(input, "SourceResource", _migrationhubSourceResource); err != nil {
			log.Errorf("invalid --source-resource: %s", err.Error())
			return
		}
	}
	if len(_migrationhubDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _migrationhubDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateSourceResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a progress update stream which is an AWS resource used for access
// control as well as a namespace for migration task names that is implicitly
// linked to your AWS account. It must uniquely identify the migration tool as it
// is used for all updates made by the tool; however, it does not need to be unique
// for each AWS account because it is scoped to the AWS account.
func migrationhub_CreateProgressUpdateStream(cfg aws.Config, client *migrationhub.Client) {
	input := &migrationhub.CreateProgressUpdateStreamInput{
		// ProgressUpdateStreamName: *string, // Required
	}

	if len(_migrationhubProgressUpdateStreamName) > 0 {
		input.ProgressUpdateStreamName = aws.String(_migrationhubProgressUpdateStreamName)
	}
	if len(_migrationhubDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _migrationhubDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProgressUpdateStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a progress update stream, including all of its tasks, which was
// previously created as an AWS resource used for access control. This API has the
// following traits:
//
// - The only parameter needed for DeleteProgressUpdateStream is the stream name
// (same as a CreateProgressUpdateStream call).
//
// - The call will return, and a background process will asynchronously delete
// the stream and all of its resources (tasks, associated resources, resource
// attributes, created artifacts).
//
// - If the stream takes time to be deleted, it might still show up on a
// ListProgressUpdateStreams call.
//
// - CreateProgressUpdateStream , ImportMigrationTask , NotifyMigrationTaskState
// , and all Associate[*] APIs related to the tasks belonging to the stream will
// throw "InvalidInputException" if the stream of the same name is in the process
// of being deleted.
//
// - Once the stream and all of its resources are deleted,
// CreateProgressUpdateStream for a stream of the same name will succeed, and
// that stream will be an entirely new logical resource (without any resources
// associated with the old stream).
func migrationhub_DeleteProgressUpdateStream(cfg aws.Config, client *migrationhub.Client) {
	input := &migrationhub.DeleteProgressUpdateStreamInput{
		// ProgressUpdateStreamName: *string, // Required
	}

	if len(_migrationhubProgressUpdateStreamName) > 0 {
		input.ProgressUpdateStreamName = aws.String(_migrationhubProgressUpdateStreamName)
	}
	if len(_migrationhubDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _migrationhubDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteProgressUpdateStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the migration status of an application.
func migrationhub_DescribeApplicationState(cfg aws.Config, client *migrationhub.Client) {
	input := &migrationhub.DescribeApplicationStateInput{
		// ApplicationId: *string, // Required
	}

	if len(_migrationhubApplicationId) > 0 {
		input.ApplicationId = aws.String(_migrationhubApplicationId)
	}

	if resp, err := client.DescribeApplicationState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of all attributes associated with a specific migration task.
func migrationhub_DescribeMigrationTask(cfg aws.Config, client *migrationhub.Client) {
	input := &migrationhub.DescribeMigrationTaskInput{
		// MigrationTaskName: *string, // Required
		// ProgressUpdateStream: *string, // Required
	}

	if len(_migrationhubMigrationTaskName) > 0 {
		input.MigrationTaskName = aws.String(_migrationhubMigrationTaskName)
	}
	if len(_migrationhubProgressUpdateStream) > 0 {
		input.ProgressUpdateStream = aws.String(_migrationhubProgressUpdateStream)
	}

	if resp, err := client.DescribeMigrationTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a created artifact of an AWS resource with a migration task
// performed by a migration tool that was previously associated. This API has the
// following traits:
//
// - A migration user can call the DisassociateCreatedArtifacts operation to
// disassociate a created AWS Artifact from a migration task.
//
// - The created artifact name must be provided in ARN (Amazon Resource Name)
// format which will contain information about type and region; for example:
// arn:aws:ec2:us-east-1:488216288981:image/ami-6d0ba87b .
//
// - Examples of the AWS resource behind the created artifact are, AMI's, EC2
// instance, or RDS instance, etc.
func migrationhub_DisassociateCreatedArtifact(cfg aws.Config, client *migrationhub.Client) {
	input := &migrationhub.DisassociateCreatedArtifactInput{
		// CreatedArtifactName: *string, // Required
		// MigrationTaskName: *string, // Required
		// ProgressUpdateStream: *string, // Required
	}

	if len(_migrationhubCreatedArtifactName) > 0 {
		input.CreatedArtifactName = aws.String(_migrationhubCreatedArtifactName)
	}
	if len(_migrationhubMigrationTaskName) > 0 {
		input.MigrationTaskName = aws.String(_migrationhubMigrationTaskName)
	}
	if len(_migrationhubProgressUpdateStream) > 0 {
		input.ProgressUpdateStream = aws.String(_migrationhubProgressUpdateStream)
	}
	if len(_migrationhubDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _migrationhubDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}

	if resp, err := client.DisassociateCreatedArtifact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociate an Application Discovery Service discovered resource from a
// migration task.
func migrationhub_DisassociateDiscoveredResource(cfg aws.Config, client *migrationhub.Client) {
	input := &migrationhub.DisassociateDiscoveredResourceInput{
		// ConfigurationId: *string, // Required
		// MigrationTaskName: *string, // Required
		// ProgressUpdateStream: *string, // Required
	}

	if len(_migrationhubConfigurationId) > 0 {
		input.ConfigurationId = aws.String(_migrationhubConfigurationId)
	}
	if len(_migrationhubMigrationTaskName) > 0 {
		input.MigrationTaskName = aws.String(_migrationhubMigrationTaskName)
	}
	if len(_migrationhubProgressUpdateStream) > 0 {
		input.ProgressUpdateStream = aws.String(_migrationhubProgressUpdateStream)
	}
	if len(_migrationhubDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _migrationhubDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}

	if resp, err := client.DisassociateDiscoveredResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the association between a source resource and a migration task.
func migrationhub_DisassociateSourceResource(cfg aws.Config, client *migrationhub.Client) {
	input := &migrationhub.DisassociateSourceResourceInput{
		// MigrationTaskName: *string, // Required
		// ProgressUpdateStream: *string, // Required
		// SourceResourceName: *string, // Required
	}

	if len(_migrationhubMigrationTaskName) > 0 {
		input.MigrationTaskName = aws.String(_migrationhubMigrationTaskName)
	}
	if len(_migrationhubProgressUpdateStream) > 0 {
		input.ProgressUpdateStream = aws.String(_migrationhubProgressUpdateStream)
	}
	if len(_migrationhubSourceResourceName) > 0 {
		input.SourceResourceName = aws.String(_migrationhubSourceResourceName)
	}
	if len(_migrationhubDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _migrationhubDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}

	if resp, err := client.DisassociateSourceResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers a new migration task which represents a server, database, etc., being
// migrated to AWS by a migration tool.
//
// This API is a prerequisite to calling the NotifyMigrationTaskState API as the
// migration tool must first register the migration task with Migration Hub.
func migrationhub_ImportMigrationTask(cfg aws.Config, client *migrationhub.Client) {
	input := &migrationhub.ImportMigrationTaskInput{
		// MigrationTaskName: *string, // Required
		// ProgressUpdateStream: *string, // Required
	}

	if len(_migrationhubMigrationTaskName) > 0 {
		input.MigrationTaskName = aws.String(_migrationhubMigrationTaskName)
	}
	if len(_migrationhubProgressUpdateStream) > 0 {
		input.ProgressUpdateStream = aws.String(_migrationhubProgressUpdateStream)
	}
	if len(_migrationhubDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _migrationhubDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportMigrationTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the migration statuses for your applications. If you use the optional
// ApplicationIds parameter, only the migration statuses for those applications
// will be returned.
func migrationhub_ListApplicationStates(cfg aws.Config, client *migrationhub.Client) {
	input := &migrationhub.ListApplicationStatesInput{}

	if len(_migrationhubApplicationIds) > 0 {
		input.ApplicationIds = append([]string(nil), _migrationhubApplicationIds...)
	}
	if len(_migrationhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _migrationhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_migrationhubNextToken) > 0 {
		input.NextToken = aws.String(_migrationhubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApplicationStates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*migrationhub.ListApplicationStatesOutput
	p := migrationhub.NewListApplicationStatesPaginator(client, input)
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

// Lists the created artifacts attached to a given migration task in an update
// stream. This API has the following traits:
//
// - Gets the list of the created artifacts while migration is taking place.
//
// - Shows the artifacts created by the migration tool that was associated by
// the AssociateCreatedArtifact API.
//
// - Lists created artifacts in a paginated interface.
func migrationhub_ListCreatedArtifacts(cfg aws.Config, client *migrationhub.Client) {
	input := &migrationhub.ListCreatedArtifactsInput{
		// MigrationTaskName: *string, // Required
		// ProgressUpdateStream: *string, // Required
	}

	if len(_migrationhubMigrationTaskName) > 0 {
		input.MigrationTaskName = aws.String(_migrationhubMigrationTaskName)
	}
	if len(_migrationhubProgressUpdateStream) > 0 {
		input.ProgressUpdateStream = aws.String(_migrationhubProgressUpdateStream)
	}
	if len(_migrationhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _migrationhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_migrationhubNextToken) > 0 {
		input.NextToken = aws.String(_migrationhubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCreatedArtifacts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*migrationhub.ListCreatedArtifactsOutput
	p := migrationhub.NewListCreatedArtifactsPaginator(client, input)
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

// Lists discovered resources associated with the given MigrationTask .
func migrationhub_ListDiscoveredResources(cfg aws.Config, client *migrationhub.Client) {
	input := &migrationhub.ListDiscoveredResourcesInput{
		// MigrationTaskName: *string, // Required
		// ProgressUpdateStream: *string, // Required
	}

	if len(_migrationhubMigrationTaskName) > 0 {
		input.MigrationTaskName = aws.String(_migrationhubMigrationTaskName)
	}
	if len(_migrationhubProgressUpdateStream) > 0 {
		input.ProgressUpdateStream = aws.String(_migrationhubProgressUpdateStream)
	}
	if len(_migrationhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _migrationhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_migrationhubNextToken) > 0 {
		input.NextToken = aws.String(_migrationhubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDiscoveredResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*migrationhub.ListDiscoveredResourcesOutput
	p := migrationhub.NewListDiscoveredResourcesPaginator(client, input)
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

// This is a paginated API that returns all the migration-task states for the
// specified MigrationTaskName and ProgressUpdateStream .
func migrationhub_ListMigrationTaskUpdates(cfg aws.Config, client *migrationhub.Client) {
	input := &migrationhub.ListMigrationTaskUpdatesInput{
		// MigrationTaskName: *string, // Required
		// ProgressUpdateStream: *string, // Required
	}

	if len(_migrationhubMigrationTaskName) > 0 {
		input.MigrationTaskName = aws.String(_migrationhubMigrationTaskName)
	}
	if len(_migrationhubProgressUpdateStream) > 0 {
		input.ProgressUpdateStream = aws.String(_migrationhubProgressUpdateStream)
	}
	if len(_migrationhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _migrationhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_migrationhubNextToken) > 0 {
		input.NextToken = aws.String(_migrationhubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMigrationTaskUpdates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*migrationhub.ListMigrationTaskUpdatesOutput
	p := migrationhub.NewListMigrationTaskUpdatesPaginator(client, input)
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

// Lists all, or filtered by resource name, migration tasks associated with the
// user account making this call. This API has the following traits:
//
// - Can show a summary list of the most recent migration tasks.
//
// - Can show a summary list of migration tasks associated with a given
// discovered resource.
//
// - Lists migration tasks in a paginated interface.
func migrationhub_ListMigrationTasks(cfg aws.Config, client *migrationhub.Client) {
	input := &migrationhub.ListMigrationTasksInput{}

	if len(_migrationhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _migrationhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_migrationhubNextToken) > 0 {
		input.NextToken = aws.String(_migrationhubNextToken)
	}
	if len(_migrationhubResourceName) > 0 {
		input.ResourceName = aws.String(_migrationhubResourceName)
	}

	if disablePaginator() {
		if resp, err := client.ListMigrationTasks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*migrationhub.ListMigrationTasksOutput
	p := migrationhub.NewListMigrationTasksPaginator(client, input)
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

// Lists progress update streams associated with the user account making this call.
func migrationhub_ListProgressUpdateStreams(cfg aws.Config, client *migrationhub.Client) {
	input := &migrationhub.ListProgressUpdateStreamsInput{}

	if len(_migrationhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _migrationhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_migrationhubNextToken) > 0 {
		input.NextToken = aws.String(_migrationhubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListProgressUpdateStreams(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*migrationhub.ListProgressUpdateStreamsOutput
	p := migrationhub.NewListProgressUpdateStreamsPaginator(client, input)
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

// Lists all the source resource that are associated with the specified
// MigrationTaskName and ProgressUpdateStream .
func migrationhub_ListSourceResources(cfg aws.Config, client *migrationhub.Client) {
	input := &migrationhub.ListSourceResourcesInput{
		// MigrationTaskName: *string, // Required
		// ProgressUpdateStream: *string, // Required
	}

	if len(_migrationhubMigrationTaskName) > 0 {
		input.MigrationTaskName = aws.String(_migrationhubMigrationTaskName)
	}
	if len(_migrationhubProgressUpdateStream) > 0 {
		input.ProgressUpdateStream = aws.String(_migrationhubProgressUpdateStream)
	}
	if len(_migrationhubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _migrationhubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_migrationhubNextToken) > 0 {
		input.NextToken = aws.String(_migrationhubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSourceResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*migrationhub.ListSourceResourcesOutput
	p := migrationhub.NewListSourceResourcesPaginator(client, input)
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

// Sets the migration state of an application. For a given application identified
// by the value passed to ApplicationId , its status is set or updated by passing
// one of three values to Status : NOT_STARTED | IN_PROGRESS | COMPLETED .
func migrationhub_NotifyApplicationState(cfg aws.Config, client *migrationhub.Client) {
	input := &migrationhub.NotifyApplicationStateInput{
		// ApplicationId: *string, // Required
		// Status: types.ApplicationStatus, // Required
	}

	if len(_migrationhubApplicationId) > 0 {
		input.ApplicationId = aws.String(_migrationhubApplicationId)
	}
	if len(_migrationhubStatus) > 0 {
		if err := assignInputField(input, "Status", _migrationhubStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_migrationhubDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _migrationhubDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_migrationhubUpdateDateTime) > 0 {
		if err := assignInputField(input, "UpdateDateTime", _migrationhubUpdateDateTime); err != nil {
			log.Errorf("invalid --update-date-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.NotifyApplicationState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Notifies Migration Hub of the current status, progress, or other detail
// regarding a migration task. This API has the following traits:
//
// - Migration tools will call the NotifyMigrationTaskState API to share the
// latest progress and status.
//
// - MigrationTaskName is used for addressing updates to the correct target.
//
// - ProgressUpdateStream is used for access control and to provide a namespace
// for each migration tool.
func migrationhub_NotifyMigrationTaskState(cfg aws.Config, client *migrationhub.Client) {
	input := &migrationhub.NotifyMigrationTaskStateInput{
		// MigrationTaskName: *string, // Required
		// NextUpdateSeconds: int32, // Required
		// ProgressUpdateStream: *string, // Required
		// Task: *types.Task, // Required
		// UpdateDateTime: *time.Time, // Required
	}

	if len(_migrationhubMigrationTaskName) > 0 {
		input.MigrationTaskName = aws.String(_migrationhubMigrationTaskName)
	}
	if len(_migrationhubNextUpdateSeconds) > 0 {
		if err := assignInputField(input, "NextUpdateSeconds", _migrationhubNextUpdateSeconds); err != nil {
			log.Errorf("invalid --next-update-seconds: %s", err.Error())
			return
		}
	}
	if len(_migrationhubProgressUpdateStream) > 0 {
		input.ProgressUpdateStream = aws.String(_migrationhubProgressUpdateStream)
	}
	if len(_migrationhubTask) > 0 {
		if err := assignInputField(input, "Task", _migrationhubTask); err != nil {
			log.Errorf("invalid --task: %s", err.Error())
			return
		}
	}
	if len(_migrationhubUpdateDateTime) > 0 {
		if err := assignInputField(input, "UpdateDateTime", _migrationhubUpdateDateTime); err != nil {
			log.Errorf("invalid --update-date-time: %s", err.Error())
			return
		}
	}
	if len(_migrationhubDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _migrationhubDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}

	if resp, err := client.NotifyMigrationTaskState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides identifying details of the resource being migrated so that it can be
// associated in the Application Discovery Service repository. This association
// occurs asynchronously after PutResourceAttributes returns.
//
// - Keep in mind that subsequent calls to PutResourceAttributes will override
// previously stored attributes. For example, if it is first called with a MAC
// address, but later, it is desired to add an IP address, it will then be required
// to call it with both the IP and MAC addresses to prevent overriding the MAC
// address.
//
// - Note the instructions regarding the special use case of the [ResourceAttributeList]
// ResourceAttributeList parameter when specifying any "VM" related value.
//
// Because this is an asynchronous call, it will always return 200, whether an
// association occurs or not. To confirm if an association was found based on the
// provided details, call ListDiscoveredResources .
//
// [ResourceAttributeList]: https://docs.aws.amazon.com/migrationhub/latest/ug/API_PutResourceAttributes.html#migrationhub-PutResourceAttributes-request-ResourceAttributeList
func migrationhub_PutResourceAttributes(cfg aws.Config, client *migrationhub.Client) {
	input := &migrationhub.PutResourceAttributesInput{
		// MigrationTaskName: *string, // Required
		// ProgressUpdateStream: *string, // Required
		// ResourceAttributeList: []types.ResourceAttribute, // Required
	}

	if len(_migrationhubMigrationTaskName) > 0 {
		input.MigrationTaskName = aws.String(_migrationhubMigrationTaskName)
	}
	if len(_migrationhubProgressUpdateStream) > 0 {
		input.ProgressUpdateStream = aws.String(_migrationhubProgressUpdateStream)
	}
	if len(_migrationhubResourceAttributeList) > 0 {
		if err := assignInputField(input, "ResourceAttributeList", _migrationhubResourceAttributeList); err != nil {
			log.Errorf("invalid --resource-attribute-list: %s", err.Error())
			return
		}
	}
	if len(_migrationhubDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _migrationhubDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutResourceAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_migrationhubCmd)
	_migrationhubCmd.Flags().SortFlags = false

	_migrationhubCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_migrationhubCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_migrationhubCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_migrationhubCmd.Flags().StringVarP(&_migrationhubApplicationId, "application-id", "", "", "Application ID")
	_migrationhubCmd.Flags().StringSliceVarP(&_migrationhubApplicationIds, "application-ids", "", nil, "Application Ids")
	_migrationhubCmd.Flags().StringVarP(&_migrationhubConfigurationId, "configuration-id", "", "", "Configuration ID")
	_migrationhubCmd.Flags().StringVarP(&_migrationhubCreatedArtifact, "created-artifact", "", "", "Created Artifact")
	_migrationhubCmd.Flags().StringVarP(&_migrationhubCreatedArtifactName, "created-artifact-name", "", "", "Created Artifact Name")
	_migrationhubCmd.Flags().StringVarP(&_migrationhubDiscoveredResource, "discovered-resource", "", "", "Discovered Resource")
	_migrationhubCmd.Flags().StringVarP(&_migrationhubDryRun, "dry-run", "", "", "Dry Run")
	_migrationhubCmd.Flags().StringVarP(&_migrationhubMaxResults, "max-results", "", "", "Max Results")
	_migrationhubCmd.Flags().StringVarP(&_migrationhubMigrationTaskName, "migration-task-name", "", "", "Migration Task Name")
	_migrationhubCmd.Flags().StringVarP(&_migrationhubNextToken, "next-token", "", "", "Next Token")
	_migrationhubCmd.Flags().StringVarP(&_migrationhubNextUpdateSeconds, "next-update-seconds", "", "", "Next Update Seconds")
	_migrationhubCmd.Flags().StringVarP(&_migrationhubProgressUpdateStream, "progress-update-stream", "", "", "Progress Update Stream")
	_migrationhubCmd.Flags().StringVarP(&_migrationhubProgressUpdateStreamName, "progress-update-stream-name", "", "", "Progress Update Stream Name")
	_migrationhubCmd.Flags().StringVarP(&_migrationhubResourceAttributeList, "resource-attribute-list", "", "", "Resource Attribute List")
	_migrationhubCmd.Flags().StringVarP(&_migrationhubResourceName, "resource-name", "", "", "Resource Name")
	_migrationhubCmd.Flags().StringVarP(&_migrationhubSourceResource, "source-resource", "", "", "Source Resource")
	_migrationhubCmd.Flags().StringVarP(&_migrationhubSourceResourceName, "source-resource-name", "", "", "Source Resource Name")
	_migrationhubCmd.Flags().StringVarP(&_migrationhubStatus, "status", "", "", "Status")
	_migrationhubCmd.Flags().StringVarP(&_migrationhubTask, "task", "", "", "Task")
	_migrationhubCmd.Flags().StringVarP(&_migrationhubUpdateDateTime, "update-date-time", "", "", "Update Date Time")

	_migrationhubCmd.Flags().BoolVarP(&_migrationhubAssociateCreatedArtifact, "associate-created-artifact", "", false, "Associate Created Artifact")
	_migrationhubCmd.Flags().BoolVarP(&_migrationhubAssociateDiscoveredResource, "associate-discovered-resource", "", false, "Associate Discovered Resource")
	_migrationhubCmd.Flags().BoolVarP(&_migrationhubAssociateSourceResource, "associate-source-resource", "", false, "Associate Source Resource")
	_migrationhubCmd.Flags().BoolVarP(&_migrationhubCreateProgressUpdateStream, "create-progress-update-stream", "", false, "Create Progress Update Stream")
	_migrationhubCmd.Flags().BoolVarP(&_migrationhubDeleteProgressUpdateStream, "delete-progress-update-stream", "", false, "Delete Progress Update Stream")
	_migrationhubCmd.Flags().BoolVarP(&_migrationhubDescribeApplicationState, "describe-application-state", "", false, "Describe Application State")
	_migrationhubCmd.Flags().BoolVarP(&_migrationhubDescribeMigrationTask, "describe-migration-task", "", false, "Describe Migration Task")
	_migrationhubCmd.Flags().BoolVarP(&_migrationhubDisassociateCreatedArtifact, "disassociate-created-artifact", "", false, "Disassociate Created Artifact")
	_migrationhubCmd.Flags().BoolVarP(&_migrationhubDisassociateDiscoveredResource, "disassociate-discovered-resource", "", false, "Disassociate Discovered Resource")
	_migrationhubCmd.Flags().BoolVarP(&_migrationhubDisassociateSourceResource, "disassociate-source-resource", "", false, "Disassociate Source Resource")
	_migrationhubCmd.Flags().BoolVarP(&_migrationhubImportMigrationTask, "import-migration-task", "", false, "Import Migration Task")
	_migrationhubCmd.Flags().BoolVarP(&_migrationhubListApplicationStates, "list-application-states", "", false, "List Application States")
	_migrationhubCmd.Flags().BoolVarP(&_migrationhubListCreatedArtifacts, "list-created-artifacts", "", false, "List Created Artifacts")
	_migrationhubCmd.Flags().BoolVarP(&_migrationhubListDiscoveredResources, "list-discovered-resources", "", false, "List Discovered Resources")
	_migrationhubCmd.Flags().BoolVarP(&_migrationhubListMigrationTaskUpdates, "list-migration-task-updates", "", false, "List Migration Task Updates")
	_migrationhubCmd.Flags().BoolVarP(&_migrationhubListMigrationTasks, "list-migration-tasks", "", false, "List Migration Tasks")
	_migrationhubCmd.Flags().BoolVarP(&_migrationhubListProgressUpdateStreams, "list-progress-update-streams", "", false, "List Progress Update Streams")
	_migrationhubCmd.Flags().BoolVarP(&_migrationhubListSourceResources, "list-source-resources", "", false, "List Source Resources")
	_migrationhubCmd.Flags().BoolVarP(&_migrationhubNotifyApplicationState, "notify-application-state", "", false, "Notify Application State")
	_migrationhubCmd.Flags().BoolVarP(&_migrationhubNotifyMigrationTaskState, "notify-migration-task-state", "", false, "Notify Migration Task State")
	_migrationhubCmd.Flags().BoolVarP(&_migrationhubPutResourceAttributes, "put-resource-attributes", "", false, "Put Resource Attributes")

}
