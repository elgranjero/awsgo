package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/gameliftstreams"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// gameliftstreamsCmd represents the gameliftstreams command
var _gameliftstreamsCmd = &cobra.Command{
	Use:   "gameliftstreams",
	Short: "AWS gameliftstreams CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := gameliftstreams.NewFromConfig(cfg)
		if _gameliftstreamsAddStreamGroupLocations {
			gameliftstreams_AddStreamGroupLocations(cfg, client)
			return
		}
		if _gameliftstreamsAssociateApplications {
			gameliftstreams_AssociateApplications(cfg, client)
			return
		}
		if _gameliftstreamsCreateApplication {
			gameliftstreams_CreateApplication(cfg, client)
			return
		}
		if _gameliftstreamsCreateStreamGroup {
			gameliftstreams_CreateStreamGroup(cfg, client)
			return
		}
		if _gameliftstreamsCreateStreamSessionConnection {
			gameliftstreams_CreateStreamSessionConnection(cfg, client)
			return
		}
		if _gameliftstreamsDeleteApplication {
			gameliftstreams_DeleteApplication(cfg, client)
			return
		}
		if _gameliftstreamsDeleteStreamGroup {
			gameliftstreams_DeleteStreamGroup(cfg, client)
			return
		}
		if _gameliftstreamsDisassociateApplications {
			gameliftstreams_DisassociateApplications(cfg, client)
			return
		}
		if _gameliftstreamsExportStreamSessionFiles {
			gameliftstreams_ExportStreamSessionFiles(cfg, client)
			return
		}
		if _gameliftstreamsGetApplication {
			gameliftstreams_GetApplication(cfg, client)
			return
		}
		if _gameliftstreamsGetStreamGroup {
			gameliftstreams_GetStreamGroup(cfg, client)
			return
		}
		if _gameliftstreamsGetStreamSession {
			gameliftstreams_GetStreamSession(cfg, client)
			return
		}
		if _gameliftstreamsListApplications {
			gameliftstreams_ListApplications(cfg, client)
			return
		}
		if _gameliftstreamsListStreamGroups {
			gameliftstreams_ListStreamGroups(cfg, client)
			return
		}
		if _gameliftstreamsListStreamSessions {
			gameliftstreams_ListStreamSessions(cfg, client)
			return
		}
		if _gameliftstreamsListStreamSessionsByAccount {
			gameliftstreams_ListStreamSessionsByAccount(cfg, client)
			return
		}
		if _gameliftstreamsListTagsForResource {
			gameliftstreams_ListTagsForResource(cfg, client)
			return
		}
		if _gameliftstreamsRemoveStreamGroupLocations {
			gameliftstreams_RemoveStreamGroupLocations(cfg, client)
			return
		}
		if _gameliftstreamsStartStreamSession {
			gameliftstreams_StartStreamSession(cfg, client)
			return
		}
		if _gameliftstreamsTagResource {
			gameliftstreams_TagResource(cfg, client)
			return
		}
		if _gameliftstreamsTerminateStreamSession {
			gameliftstreams_TerminateStreamSession(cfg, client)
			return
		}
		if _gameliftstreamsUntagResource {
			gameliftstreams_UntagResource(cfg, client)
			return
		}
		if _gameliftstreamsUpdateApplication {
			gameliftstreams_UpdateApplication(cfg, client)
			return
		}
		if _gameliftstreamsUpdateStreamGroup {
			gameliftstreams_UpdateStreamGroup(cfg, client)
			return
		}

	},
}

var (
	_gameliftstreamsAddStreamGroupLocations       bool
	_gameliftstreamsAssociateApplications         bool
	_gameliftstreamsCreateApplication             bool
	_gameliftstreamsCreateStreamGroup             bool
	_gameliftstreamsCreateStreamSessionConnection bool
	_gameliftstreamsDeleteApplication             bool
	_gameliftstreamsDeleteStreamGroup             bool
	_gameliftstreamsDisassociateApplications      bool
	_gameliftstreamsExportStreamSessionFiles      bool
	_gameliftstreamsGetApplication                bool
	_gameliftstreamsGetStreamGroup                bool
	_gameliftstreamsGetStreamSession              bool
	_gameliftstreamsListApplications              bool
	_gameliftstreamsListStreamGroups              bool
	_gameliftstreamsListStreamSessions            bool
	_gameliftstreamsListStreamSessionsByAccount   bool
	_gameliftstreamsListTagsForResource           bool
	_gameliftstreamsRemoveStreamGroupLocations    bool
	_gameliftstreamsStartStreamSession            bool
	_gameliftstreamsTagResource                   bool
	_gameliftstreamsTerminateStreamSession        bool
	_gameliftstreamsUntagResource                 bool
	_gameliftstreamsUpdateApplication             bool
	_gameliftstreamsUpdateStreamGroup             bool

	_gameliftstreamsAdditionalEnvironmentVariables string
	_gameliftstreamsAdditionalLaunchArgs           []string
	_gameliftstreamsApplicationIdentifier          string
	_gameliftstreamsApplicationIdentifiers         []string
	_gameliftstreamsApplicationLogOutputUri        string
	_gameliftstreamsApplicationLogPaths            []string
	_gameliftstreamsApplicationSourceUri           string
	_gameliftstreamsClientToken                    string
	_gameliftstreamsConnectionTimeoutSeconds       string
	_gameliftstreamsDefaultApplicationIdentifier   string
	_gameliftstreamsDescription                    string
	_gameliftstreamsExecutablePath                 string
	_gameliftstreamsExportFilesStatus              string
	_gameliftstreamsIdentifier                     string
	_gameliftstreamsLocationConfigurations         string
	_gameliftstreamsLocations                      []string
	_gameliftstreamsMaxResults                     string
	_gameliftstreamsNextToken                      string
	_gameliftstreamsOutputUri                      string
	_gameliftstreamsPerformanceStatsConfiguration  string
	_gameliftstreamsProtocol                       string
	_gameliftstreamsResourceArn                    string
	_gameliftstreamsRuntimeEnvironment             string
	_gameliftstreamsSessionLengthSeconds           string
	_gameliftstreamsSignalRequest                  string
	_gameliftstreamsStatus                         string
	_gameliftstreamsStreamClass                    string
	_gameliftstreamsStreamSessionIdentifier        string
	_gameliftstreamsTagKeys                        []string
	_gameliftstreamsTags                           string
	_gameliftstreamsUserId                         string
)

// Add locations that can host stream sessions. To add a location, the stream
// group must be in ACTIVE status. You configure locations and their corresponding
// capacity for each stream group. Creating a stream group in a location that's
// nearest to your end users can help minimize latency and improve quality.
//
// This operation provisions stream capacity at the specified locations. By
// default, all locations have 1 or 2 capacity, depending on the stream class
// option: 2 for 'High' and 1 for 'Ultra' and 'Win2022'. This operation also copies
// the content files of all associated applications to an internal S3 bucket at
// each location. This allows Amazon GameLift Streams to host performant stream
// sessions.
func gameliftstreams_AddStreamGroupLocations(cfg aws.Config, client *gameliftstreams.Client) {
	input := &gameliftstreams.AddStreamGroupLocationsInput{
		// Identifier: *string, // Required
		// LocationConfigurations: []types.LocationConfiguration, // Required
	}

	if len(_gameliftstreamsIdentifier) > 0 {
		input.Identifier = aws.String(_gameliftstreamsIdentifier)
	}
	if len(_gameliftstreamsLocationConfigurations) > 0 {
		if err := assignInputField(input, "LocationConfigurations", _gameliftstreamsLocationConfigurations); err != nil {
			log.Errorf("invalid --location-configurations: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddStreamGroupLocations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// When you associate, or link, an application with a stream group, then Amazon
// GameLift Streams can launch the application using the stream group's allocated
// compute resources. The stream group must be in ACTIVE status. You can reverse
// this action by using [DisassociateApplications].
//
// If a stream group does not already have a linked application, Amazon GameLift
// Streams will automatically assign the first application provided in
// ApplicationIdentifiers as the default.
//
// [DisassociateApplications]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_DisassociateApplications.html
func gameliftstreams_AssociateApplications(cfg aws.Config, client *gameliftstreams.Client) {
	input := &gameliftstreams.AssociateApplicationsInput{
		// ApplicationIdentifiers: []string, // Required
		// Identifier: *string, // Required
	}

	if len(_gameliftstreamsApplicationIdentifiers) > 0 {
		input.ApplicationIdentifiers = append([]string(nil), _gameliftstreamsApplicationIdentifiers...)
	}
	if len(_gameliftstreamsIdentifier) > 0 {
		input.Identifier = aws.String(_gameliftstreamsIdentifier)
	}

	if resp, err := client.AssociateApplications(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an application resource in Amazon GameLift Streams, which specifies the
// application content you want to stream, such as a game build or other software,
// and configures the settings to run it.
//
// Before you create an application, upload your application content files to an
// Amazon Simple Storage Service (Amazon S3) bucket. For more information, see
// Getting Started in the Amazon GameLift Streams Developer Guide.
//
// Make sure that your files in the Amazon S3 bucket are the correct version you
// want to use. If you change the files at a later time, you will need to create a
// new Amazon GameLift Streams application.
//
// If the request is successful, Amazon GameLift Streams begins to create an
// application and sets the status to INITIALIZED . When an application reaches
// READY status, you can use the application to set up stream groups and start
// streams. To track application status, call [GetApplication].
//
// [GetApplication]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_GetApplication.html
func gameliftstreams_CreateApplication(cfg aws.Config, client *gameliftstreams.Client) {
	input := &gameliftstreams.CreateApplicationInput{
		// ApplicationSourceUri: *string, // Required
		// Description: *string, // Required
		// ExecutablePath: *string, // Required
		// RuntimeEnvironment: *types.RuntimeEnvironment, // Required
	}

	if len(_gameliftstreamsApplicationSourceUri) > 0 {
		input.ApplicationSourceUri = aws.String(_gameliftstreamsApplicationSourceUri)
	}
	if len(_gameliftstreamsDescription) > 0 {
		input.Description = aws.String(_gameliftstreamsDescription)
	}
	if len(_gameliftstreamsExecutablePath) > 0 {
		input.ExecutablePath = aws.String(_gameliftstreamsExecutablePath)
	}
	if len(_gameliftstreamsRuntimeEnvironment) > 0 {
		if err := assignInputField(input, "RuntimeEnvironment", _gameliftstreamsRuntimeEnvironment); err != nil {
			log.Errorf("invalid --runtime-environment: %s", err.Error())
			return
		}
	}
	if len(_gameliftstreamsApplicationLogOutputUri) > 0 {
		input.ApplicationLogOutputUri = aws.String(_gameliftstreamsApplicationLogOutputUri)
	}
	if len(_gameliftstreamsApplicationLogPaths) > 0 {
		input.ApplicationLogPaths = append([]string(nil), _gameliftstreamsApplicationLogPaths...)
	}
	if len(_gameliftstreamsClientToken) > 0 {
		input.ClientToken = aws.String(_gameliftstreamsClientToken)
	}
	if len(_gameliftstreamsTags) > 0 {
		if err := assignInputField(input, "Tags", _gameliftstreamsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stream groups manage how Amazon GameLift Streams allocates resources and
// handles concurrent streams, allowing you to effectively manage capacity and
// costs. Within a stream group, you specify an application to stream, streaming
// locations and their capacity, and the stream class you want to use when
// streaming applications to your end-users. A stream class defines the hardware
// configuration of the compute resources that Amazon GameLift Streams will use
// when streaming, such as the CPU, GPU, and memory.
//
// Stream capacity represents the number of concurrent streams that can be active
// at a time. You set stream capacity per location, per stream group. The following
// capacity settings are available:
//
// - Always-on capacity: This setting, if non-zero, indicates minimum streaming
// capacity which is allocated to you and is never released back to the service.
// You pay for this base level of capacity at all times, whether used or idle.
//
// - Maximum capacity: This indicates the maximum capacity that the service can
// allocate for you. Newly created streams may take a few minutes to start.
// Capacity is released back to the service when idle. You pay for capacity that is
// allocated to you until it is released.
//
// - Target-idle capacity: This indicates idle capacity which the service
// pre-allocates and holds for you in anticipation of future activity. This helps
// to insulate your users from capacity-allocation delays. You pay for capacity
// which is held in this intentional idle state.
//
// Values for capacity must be whole number multiples of the tenancy value of the
// stream group's stream class.
//
// To adjust the capacity of any ACTIVE stream group, call [UpdateStreamGroup].
//
// If the CreateStreamGroup request is successful, Amazon GameLift Streams assigns
// a unique ID to the stream group resource and sets the status to ACTIVATING . It
// can take a few minutes for Amazon GameLift Streams to finish creating the stream
// group while it searches for unallocated compute resources and provisions them.
// When complete, the stream group status will be ACTIVE and you can start stream
// sessions by using [StartStreamSession]. To check the stream group's status, call [GetStreamGroup].
//
// Stream groups should be recreated every 3-4 weeks to pick up important service
// updates and fixes. Stream groups that are older than 180 days can no longer be
// updated with new application associations. Stream groups expire when they are
// 365 days old, at which point they can no longer stream sessions. The exact
// expiration date is indicated by the date value in the ExpiresAt field.
//
// [GetStreamGroup]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_GetStreamGroup.html
// [UpdateStreamGroup]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_UpdateStreamGroup.html
// [StartStreamSession]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_StartStreamSession.html
func gameliftstreams_CreateStreamGroup(cfg aws.Config, client *gameliftstreams.Client) {
	input := &gameliftstreams.CreateStreamGroupInput{
		// Description: *string, // Required
		// StreamClass: types.StreamClass, // Required
	}

	if len(_gameliftstreamsDescription) > 0 {
		input.Description = aws.String(_gameliftstreamsDescription)
	}
	if len(_gameliftstreamsStreamClass) > 0 {
		if err := assignInputField(input, "StreamClass", _gameliftstreamsStreamClass); err != nil {
			log.Errorf("invalid --stream-class: %s", err.Error())
			return
		}
	}
	if len(_gameliftstreamsClientToken) > 0 {
		input.ClientToken = aws.String(_gameliftstreamsClientToken)
	}
	if len(_gameliftstreamsDefaultApplicationIdentifier) > 0 {
		input.DefaultApplicationIdentifier = aws.String(_gameliftstreamsDefaultApplicationIdentifier)
	}
	if len(_gameliftstreamsLocationConfigurations) > 0 {
		if err := assignInputField(input, "LocationConfigurations", _gameliftstreamsLocationConfigurations); err != nil {
			log.Errorf("invalid --location-configurations: %s", err.Error())
			return
		}
	}
	if len(_gameliftstreamsTags) > 0 {
		if err := assignInputField(input, "Tags", _gameliftstreamsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateStreamGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables clients to reconnect to a stream session while preserving all session
// state and data in the disconnected session. This reconnection process can be
// initiated when a stream session is in either PENDING_CLIENT_RECONNECTION or
// ACTIVE status. The process works as follows:
//
// - Initial disconnect:
//
// - When a client disconnects or loses connection, the stream session
// transitions from CONNECTED to PENDING_CLIENT_RECONNECTION
//
// - Reconnection time window:
//
// - Clients have ConnectionTimeoutSeconds (defined in [StartStreamSession]) to reconnect before
// session termination
//
// - Your backend server must call CreateStreamSessionConnection to initiate
// reconnection
//
// - Session transitions to RECONNECTING status
//
// - Reconnection completion:
//
// - On successful CreateStreamSessionConnection, session status changes to
// ACTIVE
//
// - Provide the new connection information to the requesting client
//
// - Client must establish connection within ConnectionTimeoutSeconds
//
// - Session terminates automatically if client fails to connect in time
//
// For more information about the stream session lifecycle, see [Stream sessions] in the Amazon
// GameLift Streams Developer Guide.
//
// To begin re-connecting to an existing stream session, specify the stream group
// ID and stream session ID that you want to reconnect to, and the signal request
// to use with the stream.
//
// [Stream sessions]: https://docs.aws.amazon.com/gameliftstreams/latest/developerguide/stream-sessions.html
// [StartStreamSession]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_StartStreamSession.html
func gameliftstreams_CreateStreamSessionConnection(cfg aws.Config, client *gameliftstreams.Client) {
	input := &gameliftstreams.CreateStreamSessionConnectionInput{
		// Identifier: *string, // Required
		// SignalRequest: *string, // Required
		// StreamSessionIdentifier: *string, // Required
	}

	if len(_gameliftstreamsIdentifier) > 0 {
		input.Identifier = aws.String(_gameliftstreamsIdentifier)
	}
	if len(_gameliftstreamsSignalRequest) > 0 {
		input.SignalRequest = aws.String(_gameliftstreamsSignalRequest)
	}
	if len(_gameliftstreamsStreamSessionIdentifier) > 0 {
		input.StreamSessionIdentifier = aws.String(_gameliftstreamsStreamSessionIdentifier)
	}
	if len(_gameliftstreamsClientToken) > 0 {
		input.ClientToken = aws.String(_gameliftstreamsClientToken)
	}

	if resp, err := client.CreateStreamSessionConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Permanently deletes an Amazon GameLift Streams application resource. This also
// deletes the application content files stored with Amazon GameLift Streams.
// However, this does not delete the original files that you uploaded to your
// Amazon S3 bucket; you can delete these any time after Amazon GameLift Streams
// creates an application, which is the only time Amazon GameLift Streams accesses
// your Amazon S3 bucket.
//
// You can only delete an application that meets the following conditions:
//
// - The application is in READY or ERROR status. You cannot delete an
// application that's in PROCESSING or INITIALIZED status.
//
// - The application is not the default application of any stream groups. You
// must first delete the stream group by using [DeleteStreamGroup].
//
// - The application is not linked to any stream groups. You must first unlink
// the stream group by using [DisassociateApplications].
//
// - An application is not streaming in any ongoing stream session. You must
// wait until the client ends the stream session or call [TerminateStreamSession]to end the stream.
//
// If any active stream groups exist for this application, this request returns a
// ValidationException .
//
// [DisassociateApplications]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_DisassociateApplications.html
// [TerminateStreamSession]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_TerminateStreamSession.html
// [DeleteStreamGroup]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_DeleteStreamGroup.html
func gameliftstreams_DeleteApplication(cfg aws.Config, client *gameliftstreams.Client) {
	input := &gameliftstreams.DeleteApplicationInput{
		// Identifier: *string, // Required
	}

	if len(_gameliftstreamsIdentifier) > 0 {
		input.Identifier = aws.String(_gameliftstreamsIdentifier)
	}

	if resp, err := client.DeleteApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Permanently deletes all compute resources and information related to a stream
// group. To delete a stream group, specify the unique stream group identifier.
// During the deletion process, the stream group's status is DELETING . This
// operation stops streams in progress and prevents new streams from starting. As a
// best practice, before deleting the stream group, call [ListStreamSessions]to check for streams in
// progress and take action to stop them. When you delete a stream group, any
// application associations referring to that stream group are automatically
// removed.
//
// [ListStreamSessions]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_ListStreamSessions.html
func gameliftstreams_DeleteStreamGroup(cfg aws.Config, client *gameliftstreams.Client) {
	input := &gameliftstreams.DeleteStreamGroupInput{
		// Identifier: *string, // Required
	}

	if len(_gameliftstreamsIdentifier) > 0 {
		input.Identifier = aws.String(_gameliftstreamsIdentifier)
	}

	if resp, err := client.DeleteStreamGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// When you disassociate, or unlink, an application from a stream group, you can
// no longer stream this application by using that stream group's allocated compute
// resources. Any streams in process will continue until they terminate, which
// helps avoid interrupting an end-user's stream. Amazon GameLift Streams will not
// initiate new streams in the stream group using the disassociated application.
// The disassociate action does not affect the stream capacity of a stream group.
// To disassociate an application, the stream group must be in ACTIVE status.
//
// If you disassociate the default application, Amazon GameLift Streams will
// automatically choose a new default application from the remaining associated
// applications. To change which application is the default application, call [UpdateStreamGroup]and
// specify a new DefaultApplicationIdentifier .
//
// [UpdateStreamGroup]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_UpdateStreamGroup.html
func gameliftstreams_DisassociateApplications(cfg aws.Config, client *gameliftstreams.Client) {
	input := &gameliftstreams.DisassociateApplicationsInput{
		// ApplicationIdentifiers: []string, // Required
		// Identifier: *string, // Required
	}

	if len(_gameliftstreamsApplicationIdentifiers) > 0 {
		input.ApplicationIdentifiers = append([]string(nil), _gameliftstreamsApplicationIdentifiers...)
	}
	if len(_gameliftstreamsIdentifier) > 0 {
		input.Identifier = aws.String(_gameliftstreamsIdentifier)
	}

	if resp, err := client.DisassociateApplications(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Export the files that your application modifies or generates in a stream
// session, which can help you debug or verify your application. When your
// application runs, it generates output files such as logs, diagnostic
// information, crash dumps, save files, user data, screenshots, and so on. The
// files can be defined by the engine or frameworks that your application uses, or
// information that you've programmed your application to output.
//
// You can only call this action on a stream session that is in progress,
// specifically in one of the following statuses ACTIVE , CONNECTED ,
// PENDING_CLIENT_RECONNECTION , and RECONNECTING . You must provide an Amazon
// Simple Storage Service (Amazon S3) bucket to store the files in. When the
// session ends, Amazon GameLift Streams produces a compressed folder that contains
// all of the files and directories that were modified or created by the
// application during the stream session. AWS uses your security credentials to
// authenticate and authorize access to your Amazon S3 bucket.
//
// Amazon GameLift Streams collects the following generated and modified files.
// Find them in the corresponding folders in the .zip archive.
//
// - application/ : The folder where your application or game is stored.
//
// - profile/ : The user profile folder.
//
// - temp/ : The system temp folder.
//
// To verify the status of the exported files, use GetStreamSession.
//
// To delete the files, delete the object in the S3 bucket.
func gameliftstreams_ExportStreamSessionFiles(cfg aws.Config, client *gameliftstreams.Client) {
	input := &gameliftstreams.ExportStreamSessionFilesInput{
		// Identifier: *string, // Required
		// OutputUri: *string, // Required
		// StreamSessionIdentifier: *string, // Required
	}

	if len(_gameliftstreamsIdentifier) > 0 {
		input.Identifier = aws.String(_gameliftstreamsIdentifier)
	}
	if len(_gameliftstreamsOutputUri) > 0 {
		input.OutputUri = aws.String(_gameliftstreamsOutputUri)
	}
	if len(_gameliftstreamsStreamSessionIdentifier) > 0 {
		input.StreamSessionIdentifier = aws.String(_gameliftstreamsStreamSessionIdentifier)
	}

	if resp, err := client.ExportStreamSessionFiles(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves properties for an Amazon GameLift Streams application resource.
// Specify the ID of the application that you want to retrieve. If the operation is
// successful, it returns properties for the requested application.
func gameliftstreams_GetApplication(cfg aws.Config, client *gameliftstreams.Client) {
	input := &gameliftstreams.GetApplicationInput{
		// Identifier: *string, // Required
	}

	if len(_gameliftstreamsIdentifier) > 0 {
		input.Identifier = aws.String(_gameliftstreamsIdentifier)
	}

	if resp, err := client.GetApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves properties for a Amazon GameLift Streams stream group resource.
// Specify the ID of the stream group that you want to retrieve. If the operation
// is successful, it returns properties for the requested stream group.
func gameliftstreams_GetStreamGroup(cfg aws.Config, client *gameliftstreams.Client) {
	input := &gameliftstreams.GetStreamGroupInput{
		// Identifier: *string, // Required
	}

	if len(_gameliftstreamsIdentifier) > 0 {
		input.Identifier = aws.String(_gameliftstreamsIdentifier)
	}

	if resp, err := client.GetStreamGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves properties for a Amazon GameLift Streams stream session resource.
// Specify the Amazon Resource Name (ARN) of the stream session that you want to
// retrieve and its stream group ARN. If the operation is successful, it returns
// properties for the requested resource.
func gameliftstreams_GetStreamSession(cfg aws.Config, client *gameliftstreams.Client) {
	input := &gameliftstreams.GetStreamSessionInput{
		// Identifier: *string, // Required
		// StreamSessionIdentifier: *string, // Required
	}

	if len(_gameliftstreamsIdentifier) > 0 {
		input.Identifier = aws.String(_gameliftstreamsIdentifier)
	}
	if len(_gameliftstreamsStreamSessionIdentifier) > 0 {
		input.StreamSessionIdentifier = aws.String(_gameliftstreamsStreamSessionIdentifier)
	}

	if resp, err := client.GetStreamSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of all Amazon GameLift Streams applications that are
// associated with the Amazon Web Services account in use. This operation returns
// applications in all statuses, in no particular order. You can paginate the
// results as needed.
func gameliftstreams_ListApplications(cfg aws.Config, client *gameliftstreams.Client) {
	input := &gameliftstreams.ListApplicationsInput{}

	if len(_gameliftstreamsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _gameliftstreamsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_gameliftstreamsNextToken) > 0 {
		input.NextToken = aws.String(_gameliftstreamsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApplications(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*gameliftstreams.ListApplicationsOutput
	p := gameliftstreams.NewListApplicationsPaginator(client, input)
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

// Retrieves a list of all Amazon GameLift Streams stream groups that are
// associated with the Amazon Web Services account in use. This operation returns
// stream groups in all statuses, in no particular order. You can paginate the
// results as needed.
func gameliftstreams_ListStreamGroups(cfg aws.Config, client *gameliftstreams.Client) {
	input := &gameliftstreams.ListStreamGroupsInput{}

	if len(_gameliftstreamsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _gameliftstreamsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_gameliftstreamsNextToken) > 0 {
		input.NextToken = aws.String(_gameliftstreamsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListStreamGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*gameliftstreams.ListStreamGroupsOutput
	p := gameliftstreams.NewListStreamGroupsPaginator(client, input)
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

// Retrieves a list of Amazon GameLift Streams stream sessions that a stream group
// is hosting.
//
// To retrieve stream sessions, specify the stream group, and optionally filter by
// stream session status. You can paginate the results as needed.
//
// This operation returns the requested stream sessions in no particular order.
func gameliftstreams_ListStreamSessions(cfg aws.Config, client *gameliftstreams.Client) {
	input := &gameliftstreams.ListStreamSessionsInput{
		// Identifier: *string, // Required
	}

	if len(_gameliftstreamsIdentifier) > 0 {
		input.Identifier = aws.String(_gameliftstreamsIdentifier)
	}
	if len(_gameliftstreamsExportFilesStatus) > 0 {
		if err := assignInputField(input, "ExportFilesStatus", _gameliftstreamsExportFilesStatus); err != nil {
			log.Errorf("invalid --export-files-status: %s", err.Error())
			return
		}
	}
	if len(_gameliftstreamsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _gameliftstreamsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_gameliftstreamsNextToken) > 0 {
		input.NextToken = aws.String(_gameliftstreamsNextToken)
	}
	if len(_gameliftstreamsStatus) > 0 {
		if err := assignInputField(input, "Status", _gameliftstreamsStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListStreamSessions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*gameliftstreams.ListStreamSessionsOutput
	p := gameliftstreams.NewListStreamSessionsPaginator(client, input)
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

// Retrieves a list of Amazon GameLift Streams stream sessions that this user
// account has access to.
//
// In the returned list of stream sessions, the ExportFilesMetadata property only
// shows the Status value. To get the OutpurUri and StatusReason values, use [GetStreamSession].
//
// We don't recommend using this operation to regularly check stream session
// statuses because it's costly. Instead, to check status updates for a specific
// stream session, use [GetStreamSession].
//
// [GetStreamSession]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_GetStreamSession.html
func gameliftstreams_ListStreamSessionsByAccount(cfg aws.Config, client *gameliftstreams.Client) {
	input := &gameliftstreams.ListStreamSessionsByAccountInput{}

	if len(_gameliftstreamsExportFilesStatus) > 0 {
		if err := assignInputField(input, "ExportFilesStatus", _gameliftstreamsExportFilesStatus); err != nil {
			log.Errorf("invalid --export-files-status: %s", err.Error())
			return
		}
	}
	if len(_gameliftstreamsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _gameliftstreamsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_gameliftstreamsNextToken) > 0 {
		input.NextToken = aws.String(_gameliftstreamsNextToken)
	}
	if len(_gameliftstreamsStatus) > 0 {
		if err := assignInputField(input, "Status", _gameliftstreamsStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListStreamSessionsByAccount(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*gameliftstreams.ListStreamSessionsByAccountOutput
	p := gameliftstreams.NewListStreamSessionsByAccountPaginator(client, input)
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

// Retrieves all tags assigned to a Amazon GameLift Streams resource. To list tags
// for a resource, specify the ARN value for the resource.
//
// # Learn more
//
// [Tagging Amazon Web Services Resources]in the Amazon Web Services General Reference
//
// [Amazon Web Services Tagging Strategies]
//
// [Tagging Amazon Web Services Resources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
// [Amazon Web Services Tagging Strategies]: http://aws.amazon.com/answers/account-management/aws-tagging-strategies/
func gameliftstreams_ListTagsForResource(cfg aws.Config, client *gameliftstreams.Client) {
	input := &gameliftstreams.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_gameliftstreamsResourceArn) > 0 {
		input.ResourceArn = aws.String(_gameliftstreamsResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a set of remote locations from this stream group. To remove a
// location, the stream group must be in ACTIVE status. When you remove a
// location, Amazon GameLift Streams releases allocated compute resources in that
// location. Stream sessions can no longer start from removed locations in a stream
// group. Amazon GameLift Streams also deletes the content files of all associated
// applications that were in Amazon GameLift Streams's internal Amazon S3 bucket at
// this location.
//
// You cannot remove the Amazon Web Services Region location where you initially
// created this stream group, known as the primary location. However, you can set
// the stream capacity to zero to avoid incurring costs for allocated compute
// resources in that location.
func gameliftstreams_RemoveStreamGroupLocations(cfg aws.Config, client *gameliftstreams.Client) {
	input := &gameliftstreams.RemoveStreamGroupLocationsInput{
		// Identifier: *string, // Required
		// Locations: []string, // Required
	}

	if len(_gameliftstreamsIdentifier) > 0 {
		input.Identifier = aws.String(_gameliftstreamsIdentifier)
	}
	if len(_gameliftstreamsLocations) > 0 {
		input.Locations = append([]string(nil), _gameliftstreamsLocations...)
	}

	if resp, err := client.RemoveStreamGroupLocations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action initiates a new stream session and outputs connection information
// that clients can use to access the stream. A stream session refers to an
// instance of a stream that Amazon GameLift Streams transmits from the server to
// the end-user. A stream session runs on a compute resource that a stream group
// has allocated. The start stream session process works as follows:
//
// - Prerequisites:
//
// - You must have a stream group in ACTIVE status
//
// - You must have idle or on-demand capacity in a stream group in the location
// you want to stream from
//
// - You must have at least one application associated to the stream group (use [AssociateApplications]
// if needed)
//
// - Start stream request:
//
// - Your backend server calls StartStreamSession to initiate connection
//
// - Amazon GameLift Streams creates the stream session resource, assigns an
// Amazon Resource Name (ARN) value, and begins searching for available stream
// capacity to run the stream
//
// - Session transitions to ACTIVATING status
//
// - Placement completion:
//
// - If Amazon GameLift Streams is successful in finding capacity for the
// stream, the stream session status changes to ACTIVE status and
// StartStreamSession returns stream connection information
//
// - If Amazon GameLift Streams was not successful in finding capacity within
// the placement timeout period (defined according to the capacity type and
// platform type), the stream session status changes to ERROR status and
// StartStreamSession returns a StatusReason of placementTimeout
//
// - Connection completion:
//
// - Provide the new connection information to the requesting client
//
// - Client must establish connection within ConnectionTimeoutSeconds (specified
// in StartStreamSession parameters)
//
// - Session terminates automatically if client fails to connect in time
//
// For more information about the stream session lifecycle, see [Stream sessions] in the Amazon
// GameLift Streams Developer Guide.
//
// Timeouts to be aware of that affect a stream session:
//
// - Placement timeout: The amount of time that Amazon GameLift Streams has to
// find capacity for a stream request. Placement timeout varies based on the
// capacity type used to fulfill your stream request:
//
// - Always-on capacity: 75 seconds
//
// - On-demand capacity:
//
// - Linux/Proton runtimes: 90 seconds
//
// - Windows runtime: 10 minutes
//
// - Connection timeout: The amount of time that Amazon GameLift Streams waits
// for a client to connect to a stream session in ACTIVE status, or reconnect to
// a stream session in PENDING_CLIENT_RECONNECTION status, the latter of which
// occurs when a client disconnects or loses connection from a stream session. If
// no client connects before the timeout, Amazon GameLift Streams terminates the
// stream session. This value is specified by ConnectionTimeoutSeconds in the
// StartStreamSession parameters.
//
// - Idle timeout: A stream session will be terminated if no user input has been
// received for 60 minutes.
//
// - Maximum session length: A stream session will be terminated after this
// amount of time has elapsed since it started, regardless of any existing client
// connections. This value is specified by SessionLengthSeconds in the
// StartStreamSession parameters.
//
// To start a new stream session, specify a stream group ID and application ID,
// along with the transport protocol and signal request to use with the stream
// session.
//
// For stream groups that have multiple locations, provide a set of locations
// ordered by priority using a Locations parameter. Amazon GameLift Streams will
// start a single stream session in the next available location. An application
// must be finished replicating to a remote location before the remote location can
// host a stream.
//
// To reconnect to a stream session after a client disconnects or loses
// connection, use [CreateStreamSessionConnection].
//
// [Stream sessions]: https://docs.aws.amazon.com/gameliftstreams/latest/developerguide/stream-sessions.html
// [AssociateApplications]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_AssociateApplications.html
// [CreateStreamSessionConnection]: https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_CreateStreamSessionConnection.html
func gameliftstreams_StartStreamSession(cfg aws.Config, client *gameliftstreams.Client) {
	input := &gameliftstreams.StartStreamSessionInput{
		// ApplicationIdentifier: *string, // Required
		// Identifier: *string, // Required
		// Protocol: types.Protocol, // Required
		// SignalRequest: *string, // Required
	}

	if len(_gameliftstreamsApplicationIdentifier) > 0 {
		input.ApplicationIdentifier = aws.String(_gameliftstreamsApplicationIdentifier)
	}
	if len(_gameliftstreamsIdentifier) > 0 {
		input.Identifier = aws.String(_gameliftstreamsIdentifier)
	}
	if len(_gameliftstreamsProtocol) > 0 {
		if err := assignInputField(input, "Protocol", _gameliftstreamsProtocol); err != nil {
			log.Errorf("invalid --protocol: %s", err.Error())
			return
		}
	}
	if len(_gameliftstreamsSignalRequest) > 0 {
		input.SignalRequest = aws.String(_gameliftstreamsSignalRequest)
	}
	if len(_gameliftstreamsAdditionalEnvironmentVariables) > 0 {
		if err := assignInputField(input, "AdditionalEnvironmentVariables", _gameliftstreamsAdditionalEnvironmentVariables); err != nil {
			log.Errorf("invalid --additional-environment-variables: %s", err.Error())
			return
		}
	}
	if len(_gameliftstreamsAdditionalLaunchArgs) > 0 {
		input.AdditionalLaunchArgs = append([]string(nil), _gameliftstreamsAdditionalLaunchArgs...)
	}
	if len(_gameliftstreamsClientToken) > 0 {
		input.ClientToken = aws.String(_gameliftstreamsClientToken)
	}
	if len(_gameliftstreamsConnectionTimeoutSeconds) > 0 {
		if err := assignInputField(input, "ConnectionTimeoutSeconds", _gameliftstreamsConnectionTimeoutSeconds); err != nil {
			log.Errorf("invalid --connection-timeout-seconds: %s", err.Error())
			return
		}
	}
	if len(_gameliftstreamsDescription) > 0 {
		input.Description = aws.String(_gameliftstreamsDescription)
	}
	if len(_gameliftstreamsLocations) > 0 {
		input.Locations = append([]string(nil), _gameliftstreamsLocations...)
	}
	if len(_gameliftstreamsPerformanceStatsConfiguration) > 0 {
		if err := assignInputField(input, "PerformanceStatsConfiguration", _gameliftstreamsPerformanceStatsConfiguration); err != nil {
			log.Errorf("invalid --performance-stats-configuration: %s", err.Error())
			return
		}
	}
	if len(_gameliftstreamsSessionLengthSeconds) > 0 {
		if err := assignInputField(input, "SessionLengthSeconds", _gameliftstreamsSessionLengthSeconds); err != nil {
			log.Errorf("invalid --session-length-seconds: %s", err.Error())
			return
		}
	}
	if len(_gameliftstreamsUserId) > 0 {
		input.UserId = aws.String(_gameliftstreamsUserId)
	}

	if resp, err := client.StartStreamSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns one or more tags to a Amazon GameLift Streams resource. Use tags to
// organize Amazon Web Services resources for a range of purposes. You can assign
// tags to the following Amazon GameLift Streams resource types:
//
// - Application
//
// - StreamGroup
//
// # Learn more
//
// [Tagging Amazon Web Services Resources]in the Amazon Web Services General Reference
//
// [Amazon Web Services Tagging Strategies]
//
// [Tagging Amazon Web Services Resources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
// [Amazon Web Services Tagging Strategies]: http://aws.amazon.com/answers/account-management/aws-tagging-strategies/
func gameliftstreams_TagResource(cfg aws.Config, client *gameliftstreams.Client) {
	input := &gameliftstreams.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_gameliftstreamsResourceArn) > 0 {
		input.ResourceArn = aws.String(_gameliftstreamsResourceArn)
	}
	if len(_gameliftstreamsTags) > 0 {
		if err := assignInputField(input, "Tags", _gameliftstreamsTags); err != nil {
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

// Permanently terminates an active stream session. When called, the stream
// session status changes to TERMINATING . You can terminate a stream session in
// any status except ACTIVATING . If the stream session is in ACTIVATING status,
// an exception is thrown.
func gameliftstreams_TerminateStreamSession(cfg aws.Config, client *gameliftstreams.Client) {
	input := &gameliftstreams.TerminateStreamSessionInput{
		// Identifier: *string, // Required
		// StreamSessionIdentifier: *string, // Required
	}

	if len(_gameliftstreamsIdentifier) > 0 {
		input.Identifier = aws.String(_gameliftstreamsIdentifier)
	}
	if len(_gameliftstreamsStreamSessionIdentifier) > 0 {
		input.StreamSessionIdentifier = aws.String(_gameliftstreamsStreamSessionIdentifier)
	}

	if resp, err := client.TerminateStreamSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes one or more tags from a Amazon GameLift Streams resource. To remove
// tags, specify the Amazon GameLift Streams resource and a list of one or more
// tags to remove.
func gameliftstreams_UntagResource(cfg aws.Config, client *gameliftstreams.Client) {
	input := &gameliftstreams.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_gameliftstreamsResourceArn) > 0 {
		input.ResourceArn = aws.String(_gameliftstreamsResourceArn)
	}
	if len(_gameliftstreamsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _gameliftstreamsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the mutable configuration settings for a Amazon GameLift Streams
// application resource. You can change the Description , ApplicationLogOutputUri ,
// and ApplicationLogPaths .
//
// To update application settings, specify the application ID and provide the new
// values. If the operation is successful, it returns the complete updated set of
// settings for the application.
func gameliftstreams_UpdateApplication(cfg aws.Config, client *gameliftstreams.Client) {
	input := &gameliftstreams.UpdateApplicationInput{
		// Identifier: *string, // Required
	}

	if len(_gameliftstreamsIdentifier) > 0 {
		input.Identifier = aws.String(_gameliftstreamsIdentifier)
	}
	if len(_gameliftstreamsApplicationLogOutputUri) > 0 {
		input.ApplicationLogOutputUri = aws.String(_gameliftstreamsApplicationLogOutputUri)
	}
	if len(_gameliftstreamsApplicationLogPaths) > 0 {
		input.ApplicationLogPaths = append([]string(nil), _gameliftstreamsApplicationLogPaths...)
	}
	if len(_gameliftstreamsDescription) > 0 {
		input.Description = aws.String(_gameliftstreamsDescription)
	}

	if resp, err := client.UpdateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration settings for an Amazon GameLift Streams stream group
// resource. To update a stream group, it must be in ACTIVE status. You can change
// the description, the set of locations, and the requested capacity of a stream
// group per location. If you want to change the stream class, create a new stream
// group.
//
// Stream capacity represents the number of concurrent streams that can be active
// at a time. You set stream capacity per location, per stream group. The following
// capacity settings are available:
//
// - Always-on capacity: This setting, if non-zero, indicates minimum streaming
// capacity which is allocated to you and is never released back to the service.
// You pay for this base level of capacity at all times, whether used or idle.
//
// - Maximum capacity: This indicates the maximum capacity that the service can
// allocate for you. Newly created streams may take a few minutes to start.
// Capacity is released back to the service when idle. You pay for capacity that is
// allocated to you until it is released.
//
// - Target-idle capacity: This indicates idle capacity which the service
// pre-allocates and holds for you in anticipation of future activity. This helps
// to insulate your users from capacity-allocation delays. You pay for capacity
// which is held in this intentional idle state.
//
// Values for capacity must be whole number multiples of the tenancy value of the
// stream group's stream class.
//
// To update a stream group, specify the stream group's Amazon Resource Name (ARN)
// and provide the new values. If the request is successful, Amazon GameLift
// Streams returns the complete updated metadata for the stream group. Expired
// stream groups cannot be updated.
func gameliftstreams_UpdateStreamGroup(cfg aws.Config, client *gameliftstreams.Client) {
	input := &gameliftstreams.UpdateStreamGroupInput{
		// Identifier: *string, // Required
	}

	if len(_gameliftstreamsIdentifier) > 0 {
		input.Identifier = aws.String(_gameliftstreamsIdentifier)
	}
	if len(_gameliftstreamsDefaultApplicationIdentifier) > 0 {
		input.DefaultApplicationIdentifier = aws.String(_gameliftstreamsDefaultApplicationIdentifier)
	}
	if len(_gameliftstreamsDescription) > 0 {
		input.Description = aws.String(_gameliftstreamsDescription)
	}
	if len(_gameliftstreamsLocationConfigurations) > 0 {
		if err := assignInputField(input, "LocationConfigurations", _gameliftstreamsLocationConfigurations); err != nil {
			log.Errorf("invalid --location-configurations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateStreamGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_gameliftstreamsCmd)
	_gameliftstreamsCmd.Flags().SortFlags = false

	_gameliftstreamsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_gameliftstreamsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_gameliftstreamsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_gameliftstreamsCmd.Flags().StringVarP(&_gameliftstreamsAdditionalEnvironmentVariables, "additional-environment-variables", "", "", "Additional Environment Variables")
	_gameliftstreamsCmd.Flags().StringSliceVarP(&_gameliftstreamsAdditionalLaunchArgs, "additional-launch-args", "", nil, "Additional Launch Args")
	_gameliftstreamsCmd.Flags().StringVarP(&_gameliftstreamsApplicationIdentifier, "application-identifier", "", "", "Application Identifier")
	_gameliftstreamsCmd.Flags().StringSliceVarP(&_gameliftstreamsApplicationIdentifiers, "application-identifiers", "", nil, "Application Identifiers")
	_gameliftstreamsCmd.Flags().StringVarP(&_gameliftstreamsApplicationLogOutputUri, "application-log-output-uri", "", "", "Application Log Output URI")
	_gameliftstreamsCmd.Flags().StringSliceVarP(&_gameliftstreamsApplicationLogPaths, "application-log-paths", "", nil, "Application Log Paths")
	_gameliftstreamsCmd.Flags().StringVarP(&_gameliftstreamsApplicationSourceUri, "application-source-uri", "", "", "Application Source URI")
	_gameliftstreamsCmd.Flags().StringVarP(&_gameliftstreamsClientToken, "client-token", "", "", "Client Token")
	_gameliftstreamsCmd.Flags().StringVarP(&_gameliftstreamsConnectionTimeoutSeconds, "connection-timeout-seconds", "", "", "Connection Timeout Seconds")
	_gameliftstreamsCmd.Flags().StringVarP(&_gameliftstreamsDefaultApplicationIdentifier, "default-application-identifier", "", "", "Default Application Identifier")
	_gameliftstreamsCmd.Flags().StringVarP(&_gameliftstreamsDescription, "description", "", "", "Description")
	_gameliftstreamsCmd.Flags().StringVarP(&_gameliftstreamsExecutablePath, "executable-path", "", "", "Executable Path")
	_gameliftstreamsCmd.Flags().StringVarP(&_gameliftstreamsExportFilesStatus, "export-files-status", "", "", "Export Files Status")
	_gameliftstreamsCmd.Flags().StringVarP(&_gameliftstreamsIdentifier, "identifier", "", "", "Identifier")
	_gameliftstreamsCmd.Flags().StringVarP(&_gameliftstreamsLocationConfigurations, "location-configurations", "", "", "Location Configurations")
	_gameliftstreamsCmd.Flags().StringSliceVarP(&_gameliftstreamsLocations, "locations", "", nil, "Locations")
	_gameliftstreamsCmd.Flags().StringVarP(&_gameliftstreamsMaxResults, "max-results", "", "", "Max Results")
	_gameliftstreamsCmd.Flags().StringVarP(&_gameliftstreamsNextToken, "next-token", "", "", "Next Token")
	_gameliftstreamsCmd.Flags().StringVarP(&_gameliftstreamsOutputUri, "output-uri", "", "", "Output URI")
	_gameliftstreamsCmd.Flags().StringVarP(&_gameliftstreamsPerformanceStatsConfiguration, "performance-stats-configuration", "", "", "Performance Stats Configuration")
	_gameliftstreamsCmd.Flags().StringVarP(&_gameliftstreamsProtocol, "protocol", "", "", "Protocol")
	_gameliftstreamsCmd.Flags().StringVarP(&_gameliftstreamsResourceArn, "resource-arn", "", "", "Resource ARN")
	_gameliftstreamsCmd.Flags().StringVarP(&_gameliftstreamsRuntimeEnvironment, "runtime-environment", "", "", "Runtime Environment")
	_gameliftstreamsCmd.Flags().StringVarP(&_gameliftstreamsSessionLengthSeconds, "session-length-seconds", "", "", "Session Length Seconds")
	_gameliftstreamsCmd.Flags().StringVarP(&_gameliftstreamsSignalRequest, "signal-request", "", "", "Signal Request")
	_gameliftstreamsCmd.Flags().StringVarP(&_gameliftstreamsStatus, "status", "", "", "Status")
	_gameliftstreamsCmd.Flags().StringVarP(&_gameliftstreamsStreamClass, "stream-class", "", "", "Stream Class")
	_gameliftstreamsCmd.Flags().StringVarP(&_gameliftstreamsStreamSessionIdentifier, "stream-session-identifier", "", "", "Stream Session Identifier")
	_gameliftstreamsCmd.Flags().StringSliceVarP(&_gameliftstreamsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_gameliftstreamsCmd.Flags().StringVarP(&_gameliftstreamsTags, "tags", "", "", "Tags")
	_gameliftstreamsCmd.Flags().StringVarP(&_gameliftstreamsUserId, "user-id", "", "", "User ID")

	_gameliftstreamsCmd.Flags().BoolVarP(&_gameliftstreamsAddStreamGroupLocations, "add-stream-group-locations", "", false, "Add Stream Group Locations")
	_gameliftstreamsCmd.Flags().BoolVarP(&_gameliftstreamsAssociateApplications, "associate-applications", "", false, "Associate Applications")
	_gameliftstreamsCmd.Flags().BoolVarP(&_gameliftstreamsCreateApplication, "create-application", "", false, "Create Application")
	_gameliftstreamsCmd.Flags().BoolVarP(&_gameliftstreamsCreateStreamGroup, "create-stream-group", "", false, "Create Stream Group")
	_gameliftstreamsCmd.Flags().BoolVarP(&_gameliftstreamsCreateStreamSessionConnection, "create-stream-session-connection", "", false, "Create Stream Session Connection")
	_gameliftstreamsCmd.Flags().BoolVarP(&_gameliftstreamsDeleteApplication, "delete-application", "", false, "Delete Application")
	_gameliftstreamsCmd.Flags().BoolVarP(&_gameliftstreamsDeleteStreamGroup, "delete-stream-group", "", false, "Delete Stream Group")
	_gameliftstreamsCmd.Flags().BoolVarP(&_gameliftstreamsDisassociateApplications, "disassociate-applications", "", false, "Disassociate Applications")
	_gameliftstreamsCmd.Flags().BoolVarP(&_gameliftstreamsExportStreamSessionFiles, "export-stream-session-files", "", false, "Export Stream Session Files")
	_gameliftstreamsCmd.Flags().BoolVarP(&_gameliftstreamsGetApplication, "get-application", "", false, "Get Application")
	_gameliftstreamsCmd.Flags().BoolVarP(&_gameliftstreamsGetStreamGroup, "get-stream-group", "", false, "Get Stream Group")
	_gameliftstreamsCmd.Flags().BoolVarP(&_gameliftstreamsGetStreamSession, "get-stream-session", "", false, "Get Stream Session")
	_gameliftstreamsCmd.Flags().BoolVarP(&_gameliftstreamsListApplications, "list-applications", "", false, "List Applications")
	_gameliftstreamsCmd.Flags().BoolVarP(&_gameliftstreamsListStreamGroups, "list-stream-groups", "", false, "List Stream Groups")
	_gameliftstreamsCmd.Flags().BoolVarP(&_gameliftstreamsListStreamSessions, "list-stream-sessions", "", false, "List Stream Sessions")
	_gameliftstreamsCmd.Flags().BoolVarP(&_gameliftstreamsListStreamSessionsByAccount, "list-stream-sessions-by-account", "", false, "List Stream Sessions By Account")
	_gameliftstreamsCmd.Flags().BoolVarP(&_gameliftstreamsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_gameliftstreamsCmd.Flags().BoolVarP(&_gameliftstreamsRemoveStreamGroupLocations, "remove-stream-group-locations", "", false, "Remove Stream Group Locations")
	_gameliftstreamsCmd.Flags().BoolVarP(&_gameliftstreamsStartStreamSession, "start-stream-session", "", false, "Start Stream Session")
	_gameliftstreamsCmd.Flags().BoolVarP(&_gameliftstreamsTagResource, "tag-resource", "", false, "Tag Resource")
	_gameliftstreamsCmd.Flags().BoolVarP(&_gameliftstreamsTerminateStreamSession, "terminate-stream-session", "", false, "Terminate Stream Session")
	_gameliftstreamsCmd.Flags().BoolVarP(&_gameliftstreamsUntagResource, "untag-resource", "", false, "Untag Resource")
	_gameliftstreamsCmd.Flags().BoolVarP(&_gameliftstreamsUpdateApplication, "update-application", "", false, "Update Application")
	_gameliftstreamsCmd.Flags().BoolVarP(&_gameliftstreamsUpdateStreamGroup, "update-stream-group", "", false, "Update Stream Group")

}
