package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmeetings"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// chimesdkmeetingsCmd represents the chimesdkmeetings command
var _chimesdkmeetingsCmd = &cobra.Command{
	Use:   "chimesdkmeetings",
	Short: "AWS chimesdkmeetings CLI",
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
		client := chimesdkmeetings.NewFromConfig(cfg)
		if _chimesdkmeetingsBatchCreateAttendee {
			chimesdkmeetings_BatchCreateAttendee(cfg, client)
			return
		}
		if _chimesdkmeetingsBatchUpdateAttendeeCapabilitiesExcept {
			chimesdkmeetings_BatchUpdateAttendeeCapabilitiesExcept(cfg, client)
			return
		}
		if _chimesdkmeetingsCreateAttendee {
			chimesdkmeetings_CreateAttendee(cfg, client)
			return
		}
		if _chimesdkmeetingsCreateMeeting {
			chimesdkmeetings_CreateMeeting(cfg, client)
			return
		}
		if _chimesdkmeetingsCreateMeetingWithAttendees {
			chimesdkmeetings_CreateMeetingWithAttendees(cfg, client)
			return
		}
		if _chimesdkmeetingsDeleteAttendee {
			chimesdkmeetings_DeleteAttendee(cfg, client)
			return
		}
		if _chimesdkmeetingsDeleteMeeting {
			chimesdkmeetings_DeleteMeeting(cfg, client)
			return
		}
		if _chimesdkmeetingsGetAttendee {
			chimesdkmeetings_GetAttendee(cfg, client)
			return
		}
		if _chimesdkmeetingsGetMeeting {
			chimesdkmeetings_GetMeeting(cfg, client)
			return
		}
		if _chimesdkmeetingsListAttendees {
			chimesdkmeetings_ListAttendees(cfg, client)
			return
		}
		if _chimesdkmeetingsListTagsForResource {
			chimesdkmeetings_ListTagsForResource(cfg, client)
			return
		}
		if _chimesdkmeetingsStartMeetingTranscription {
			chimesdkmeetings_StartMeetingTranscription(cfg, client)
			return
		}
		if _chimesdkmeetingsStopMeetingTranscription {
			chimesdkmeetings_StopMeetingTranscription(cfg, client)
			return
		}
		if _chimesdkmeetingsTagResource {
			chimesdkmeetings_TagResource(cfg, client)
			return
		}
		if _chimesdkmeetingsUntagResource {
			chimesdkmeetings_UntagResource(cfg, client)
			return
		}
		if _chimesdkmeetingsUpdateAttendeeCapabilities {
			chimesdkmeetings_UpdateAttendeeCapabilities(cfg, client)
			return
		}

	},
}

var (
	_chimesdkmeetingsBatchCreateAttendee                   bool
	_chimesdkmeetingsBatchUpdateAttendeeCapabilitiesExcept bool
	_chimesdkmeetingsCreateAttendee                        bool
	_chimesdkmeetingsCreateMeeting                         bool
	_chimesdkmeetingsCreateMeetingWithAttendees            bool
	_chimesdkmeetingsDeleteAttendee                        bool
	_chimesdkmeetingsDeleteMeeting                         bool
	_chimesdkmeetingsGetAttendee                           bool
	_chimesdkmeetingsGetMeeting                            bool
	_chimesdkmeetingsListAttendees                         bool
	_chimesdkmeetingsListTagsForResource                   bool
	_chimesdkmeetingsStartMeetingTranscription             bool
	_chimesdkmeetingsStopMeetingTranscription              bool
	_chimesdkmeetingsTagResource                           bool
	_chimesdkmeetingsUntagResource                         bool
	_chimesdkmeetingsUpdateAttendeeCapabilities            bool

	_chimesdkmeetingsAttendeeId                 string
	_chimesdkmeetingsAttendees                  string
	_chimesdkmeetingsCapabilities               string
	_chimesdkmeetingsClientRequestToken         string
	_chimesdkmeetingsExcludedAttendeeIds        string
	_chimesdkmeetingsExternalMeetingId          string
	_chimesdkmeetingsExternalUserId             string
	_chimesdkmeetingsMaxResults                 string
	_chimesdkmeetingsMediaPlacementNetworkType  string
	_chimesdkmeetingsMediaRegion                string
	_chimesdkmeetingsMeetingFeatures            string
	_chimesdkmeetingsMeetingHostId              string
	_chimesdkmeetingsMeetingId                  string
	_chimesdkmeetingsNextToken                  string
	_chimesdkmeetingsNotificationsConfiguration string
	_chimesdkmeetingsPrimaryMeetingId           string
	_chimesdkmeetingsResourceARN                string
	_chimesdkmeetingsTagKeys                    []string
	_chimesdkmeetingsTags                       string
	_chimesdkmeetingsTenantIds                  []string
	_chimesdkmeetingsTranscriptionConfiguration string
)

// Creates up to 100 attendees for an active Amazon Chime SDK meeting. For more
// information about the Amazon Chime SDK, see [Using the Amazon Chime SDK]in the Amazon Chime Developer Guide.
//
// [Using the Amazon Chime SDK]: https://docs.aws.amazon.com/chime-sdk/latest/dg/meetings-sdk.html
func chimesdkmeetings_BatchCreateAttendee(cfg aws.Config, client *chimesdkmeetings.Client) {
	input := &chimesdkmeetings.BatchCreateAttendeeInput{
		// Attendees: []types.CreateAttendeeRequestItem, // Required
		// MeetingId: *string, // Required
	}

	if len(_chimesdkmeetingsAttendees) > 0 {
		if err := assignInputField(input, "Attendees", _chimesdkmeetingsAttendees); err != nil {
			log.Errorf("invalid --attendees: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmeetingsMeetingId) > 0 {
		input.MeetingId = aws.String(_chimesdkmeetingsMeetingId)
	}

	if resp, err := client.BatchCreateAttendee(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates AttendeeCapabilities except the capabilities listed in an
// ExcludedAttendeeIds table.
//
// You use the capabilities with a set of values that control what the
// capabilities can do, such as SendReceive data. For more information about those
// values, see .
//
// When using capabilities, be aware of these corner cases:
//
// - If you specify MeetingFeatures:Video:MaxResolution:None when you create a
// meeting, all API requests that include SendReceive , Send , or Receive for
// AttendeeCapabilities:Video will be rejected with ValidationError 400 .
//
// - If you specify MeetingFeatures:Content:MaxResolution:None when you create a
// meeting, all API requests that include SendReceive , Send , or Receive for
// AttendeeCapabilities:Content will be rejected with ValidationError 400 .
//
// - You can't set content capabilities to SendReceive or Receive unless you also
// set video capabilities to SendReceive or Receive . If you don't set the video
// capability to receive, the response will contain an HTTP 400 Bad Request status
// code. However, you can set your video capability to receive and you set your
// content capability to not receive.
//
// - If meeting features is defined as Video:MaxResolution:None but
// Content:MaxResolution is defined as something other than None and attendee
// capabilities are not defined in the API request, then the default attendee video
// capability is set to Receive and attendee content capability is set to
// SendReceive . This is because content SendReceive requires video to be at
// least Receive .
//
// - When you change an audio capability from None or Receive to Send or
// SendReceive , and if the attendee left their microphone unmuted, audio will
// flow from the attendee to the other meeting participants.
//
// - When you change a video or content capability from None or Receive to Send
// or SendReceive , and if the attendee turned on their video or content streams,
// remote attendees can receive those streams, but only after media renegotiation
// between the client and the Amazon Chime back-end server.
func chimesdkmeetings_BatchUpdateAttendeeCapabilitiesExcept(cfg aws.Config, client *chimesdkmeetings.Client) {
	input := &chimesdkmeetings.BatchUpdateAttendeeCapabilitiesExceptInput{
		// Capabilities: *types.AttendeeCapabilities, // Required
		// ExcludedAttendeeIds: []types.AttendeeIdItem, // Required
		// MeetingId: *string, // Required
	}

	if len(_chimesdkmeetingsCapabilities) > 0 {
		if err := assignInputField(input, "Capabilities", _chimesdkmeetingsCapabilities); err != nil {
			log.Errorf("invalid --capabilities: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmeetingsExcludedAttendeeIds) > 0 {
		if err := assignInputField(input, "ExcludedAttendeeIds", _chimesdkmeetingsExcludedAttendeeIds); err != nil {
			log.Errorf("invalid --excluded-attendee-ids: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmeetingsMeetingId) > 0 {
		input.MeetingId = aws.String(_chimesdkmeetingsMeetingId)
	}

	if resp, err := client.BatchUpdateAttendeeCapabilitiesExcept(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new attendee for an active Amazon Chime SDK meeting. For more
// information about the Amazon Chime SDK, see [Using the Amazon Chime SDK]in the Amazon Chime Developer
// Guide.
//
// [Using the Amazon Chime SDK]: https://docs.aws.amazon.com/chime-sdk/latest/dg/meetings-sdk.html
func chimesdkmeetings_CreateAttendee(cfg aws.Config, client *chimesdkmeetings.Client) {
	input := &chimesdkmeetings.CreateAttendeeInput{
		// ExternalUserId: *string, // Required
		// MeetingId: *string, // Required
	}

	if len(_chimesdkmeetingsExternalUserId) > 0 {
		input.ExternalUserId = aws.String(_chimesdkmeetingsExternalUserId)
	}
	if len(_chimesdkmeetingsMeetingId) > 0 {
		input.MeetingId = aws.String(_chimesdkmeetingsMeetingId)
	}
	if len(_chimesdkmeetingsCapabilities) > 0 {
		if err := assignInputField(input, "Capabilities", _chimesdkmeetingsCapabilities); err != nil {
			log.Errorf("invalid --capabilities: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAttendee(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Amazon Chime SDK meeting in the specified media Region with no
// initial attendees. For more information about specifying media Regions, see [Available Regions]and [Using meeting Regions]
// , both in the Amazon Chime SDK Developer Guide. For more information about the
// Amazon Chime SDK, see [Using the Amazon Chime SDK]in the Amazon Chime SDK Developer Guide.
//
// If you use this API in conjuction with the and APIs, and you don't specify the
// MeetingFeatures.Content.MaxResolution or MeetingFeatures.Video.MaxResolution
// parameters, the following defaults are used:
//
// - Content.MaxResolution: FHD
//
// - Video.MaxResolution: HD
//
// [Using meeting Regions]: https://docs.aws.amazon.com/chime-sdk/latest/dg/chime-sdk-meetings-regions.html
// [Available Regions]: https://docs.aws.amazon.com/chime-sdk/latest/dg/sdk-available-regions
// [Using the Amazon Chime SDK]: https://docs.aws.amazon.com/chime-sdk/latest/dg/meetings-sdk.html
func chimesdkmeetings_CreateMeeting(cfg aws.Config, client *chimesdkmeetings.Client) {
	input := &chimesdkmeetings.CreateMeetingInput{
		// ClientRequestToken: *string, // Required
		// ExternalMeetingId: *string, // Required
		// MediaRegion: *string, // Required
	}

	if len(_chimesdkmeetingsClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_chimesdkmeetingsClientRequestToken)
	}
	if len(_chimesdkmeetingsExternalMeetingId) > 0 {
		input.ExternalMeetingId = aws.String(_chimesdkmeetingsExternalMeetingId)
	}
	if len(_chimesdkmeetingsMediaRegion) > 0 {
		input.MediaRegion = aws.String(_chimesdkmeetingsMediaRegion)
	}
	if len(_chimesdkmeetingsMediaPlacementNetworkType) > 0 {
		if err := assignInputField(input, "MediaPlacementNetworkType", _chimesdkmeetingsMediaPlacementNetworkType); err != nil {
			log.Errorf("invalid --media-placement-network-type: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmeetingsMeetingFeatures) > 0 {
		if err := assignInputField(input, "MeetingFeatures", _chimesdkmeetingsMeetingFeatures); err != nil {
			log.Errorf("invalid --meeting-features: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmeetingsMeetingHostId) > 0 {
		input.MeetingHostId = aws.String(_chimesdkmeetingsMeetingHostId)
	}
	if len(_chimesdkmeetingsNotificationsConfiguration) > 0 {
		if err := assignInputField(input, "NotificationsConfiguration", _chimesdkmeetingsNotificationsConfiguration); err != nil {
			log.Errorf("invalid --notifications-configuration: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmeetingsPrimaryMeetingId) > 0 {
		input.PrimaryMeetingId = aws.String(_chimesdkmeetingsPrimaryMeetingId)
	}
	if len(_chimesdkmeetingsTags) > 0 {
		if err := assignInputField(input, "Tags", _chimesdkmeetingsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmeetingsTenantIds) > 0 {
		input.TenantIds = append([]string(nil), _chimesdkmeetingsTenantIds...)
	}

	if resp, err := client.CreateMeeting(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Amazon Chime SDK meeting in the specified media Region, with
// attendees. For more information about specifying media Regions, see [Available Regions]and [Using meeting Regions], both
// in the Amazon Chime SDK Developer Guide. For more information about the Amazon
// Chime SDK, see [Using the Amazon Chime SDK]in the Amazon Chime SDK Developer Guide.
//
// If you use this API in conjuction with the and APIs, and you don't specify the
// MeetingFeatures.Content.MaxResolution or MeetingFeatures.Video.MaxResolution
// parameters, the following defaults are used:
//
// - Content.MaxResolution: FHD
//
// - Video.MaxResolution: HD
//
// [Using meeting Regions]: https://docs.aws.amazon.com/chime-sdk/latest/dg/chime-sdk-meetings-regions.html
// [Available Regions]: https://docs.aws.amazon.com/chime-sdk/latest/dg/sdk-available-regions
// [Using the Amazon Chime SDK]: https://docs.aws.amazon.com/chime-sdk/latest/dg/meetings-sdk.html
func chimesdkmeetings_CreateMeetingWithAttendees(cfg aws.Config, client *chimesdkmeetings.Client) {
	input := &chimesdkmeetings.CreateMeetingWithAttendeesInput{
		// Attendees: []types.CreateAttendeeRequestItem, // Required
		// ClientRequestToken: *string, // Required
		// ExternalMeetingId: *string, // Required
		// MediaRegion: *string, // Required
	}

	if len(_chimesdkmeetingsAttendees) > 0 {
		if err := assignInputField(input, "Attendees", _chimesdkmeetingsAttendees); err != nil {
			log.Errorf("invalid --attendees: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmeetingsClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_chimesdkmeetingsClientRequestToken)
	}
	if len(_chimesdkmeetingsExternalMeetingId) > 0 {
		input.ExternalMeetingId = aws.String(_chimesdkmeetingsExternalMeetingId)
	}
	if len(_chimesdkmeetingsMediaRegion) > 0 {
		input.MediaRegion = aws.String(_chimesdkmeetingsMediaRegion)
	}
	if len(_chimesdkmeetingsMediaPlacementNetworkType) > 0 {
		if err := assignInputField(input, "MediaPlacementNetworkType", _chimesdkmeetingsMediaPlacementNetworkType); err != nil {
			log.Errorf("invalid --media-placement-network-type: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmeetingsMeetingFeatures) > 0 {
		if err := assignInputField(input, "MeetingFeatures", _chimesdkmeetingsMeetingFeatures); err != nil {
			log.Errorf("invalid --meeting-features: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmeetingsMeetingHostId) > 0 {
		input.MeetingHostId = aws.String(_chimesdkmeetingsMeetingHostId)
	}
	if len(_chimesdkmeetingsNotificationsConfiguration) > 0 {
		if err := assignInputField(input, "NotificationsConfiguration", _chimesdkmeetingsNotificationsConfiguration); err != nil {
			log.Errorf("invalid --notifications-configuration: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmeetingsPrimaryMeetingId) > 0 {
		input.PrimaryMeetingId = aws.String(_chimesdkmeetingsPrimaryMeetingId)
	}
	if len(_chimesdkmeetingsTags) > 0 {
		if err := assignInputField(input, "Tags", _chimesdkmeetingsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmeetingsTenantIds) > 0 {
		input.TenantIds = append([]string(nil), _chimesdkmeetingsTenantIds...)
	}

	if resp, err := client.CreateMeetingWithAttendees(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an attendee from the specified Amazon Chime SDK meeting and deletes
// their JoinToken . Attendees are automatically deleted when a Amazon Chime SDK
// meeting is deleted. For more information about the Amazon Chime SDK, see [Using the Amazon Chime SDK]in the
// Amazon Chime Developer Guide.
//
// [Using the Amazon Chime SDK]: https://docs.aws.amazon.com/chime-sdk/latest/dg/meetings-sdk.html
func chimesdkmeetings_DeleteAttendee(cfg aws.Config, client *chimesdkmeetings.Client) {
	input := &chimesdkmeetings.DeleteAttendeeInput{
		// AttendeeId: *string, // Required
		// MeetingId: *string, // Required
	}

	if len(_chimesdkmeetingsAttendeeId) > 0 {
		input.AttendeeId = aws.String(_chimesdkmeetingsAttendeeId)
	}
	if len(_chimesdkmeetingsMeetingId) > 0 {
		input.MeetingId = aws.String(_chimesdkmeetingsMeetingId)
	}

	if resp, err := client.DeleteAttendee(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified Amazon Chime SDK meeting. The operation deletes all
// attendees, disconnects all clients, and prevents new clients from joining the
// meeting. For more information about the Amazon Chime SDK, see [Using the Amazon Chime SDK]in the Amazon
// Chime Developer Guide.
//
// [Using the Amazon Chime SDK]: https://docs.aws.amazon.com/chime-sdk/latest/dg/meetings-sdk.html
func chimesdkmeetings_DeleteMeeting(cfg aws.Config, client *chimesdkmeetings.Client) {
	input := &chimesdkmeetings.DeleteMeetingInput{
		// MeetingId: *string, // Required
	}

	if len(_chimesdkmeetingsMeetingId) > 0 {
		input.MeetingId = aws.String(_chimesdkmeetingsMeetingId)
	}

	if resp, err := client.DeleteMeeting(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the Amazon Chime SDK attendee details for a specified meeting ID and
// attendee ID. For more information about the Amazon Chime SDK, see [Using the Amazon Chime SDK]in the Amazon
// Chime Developer Guide.
//
// [Using the Amazon Chime SDK]: https://docs.aws.amazon.com/chime-sdk/latest/dg/meetings-sdk.html
func chimesdkmeetings_GetAttendee(cfg aws.Config, client *chimesdkmeetings.Client) {
	input := &chimesdkmeetings.GetAttendeeInput{
		// AttendeeId: *string, // Required
		// MeetingId: *string, // Required
	}

	if len(_chimesdkmeetingsAttendeeId) > 0 {
		input.AttendeeId = aws.String(_chimesdkmeetingsAttendeeId)
	}
	if len(_chimesdkmeetingsMeetingId) > 0 {
		input.MeetingId = aws.String(_chimesdkmeetingsMeetingId)
	}

	if resp, err := client.GetAttendee(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the Amazon Chime SDK meeting details for the specified meeting ID. For
// more information about the Amazon Chime SDK, see [Using the Amazon Chime SDK]in the Amazon Chime Developer
// Guide.
//
// [Using the Amazon Chime SDK]: https://docs.aws.amazon.com/chime-sdk/latest/dg/meetings-sdk.html
func chimesdkmeetings_GetMeeting(cfg aws.Config, client *chimesdkmeetings.Client) {
	input := &chimesdkmeetings.GetMeetingInput{
		// MeetingId: *string, // Required
	}

	if len(_chimesdkmeetingsMeetingId) > 0 {
		input.MeetingId = aws.String(_chimesdkmeetingsMeetingId)
	}

	if resp, err := client.GetMeeting(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the attendees for the specified Amazon Chime SDK meeting. For more
// information about the Amazon Chime SDK, see [Using the Amazon Chime SDK]in the Amazon Chime Developer
// Guide.
//
// [Using the Amazon Chime SDK]: https://docs.aws.amazon.com/chime-sdk/latest/dg/meetings-sdk.html
func chimesdkmeetings_ListAttendees(cfg aws.Config, client *chimesdkmeetings.Client) {
	input := &chimesdkmeetings.ListAttendeesInput{
		// MeetingId: *string, // Required
	}

	if len(_chimesdkmeetingsMeetingId) > 0 {
		input.MeetingId = aws.String(_chimesdkmeetingsMeetingId)
	}
	if len(_chimesdkmeetingsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimesdkmeetingsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmeetingsNextToken) > 0 {
		input.NextToken = aws.String(_chimesdkmeetingsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAttendees(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chimesdkmeetings.ListAttendeesOutput
	p := chimesdkmeetings.NewListAttendeesPaginator(client, input)
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

// Returns a list of the tags available for the specified resource.
func chimesdkmeetings_ListTagsForResource(cfg aws.Config, client *chimesdkmeetings.Client) {
	input := &chimesdkmeetings.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_chimesdkmeetingsResourceARN) > 0 {
		input.ResourceARN = aws.String(_chimesdkmeetingsResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts transcription for the specified meetingId . For more information, refer
// to [Using Amazon Chime SDK live transcription]in the Amazon Chime SDK Developer Guide.
//
// If you specify an invalid configuration, a TranscriptFailed event will be sent
// with the contents of the BadRequestException generated by Amazon Transcribe.
// For more information on each parameter and which combinations are valid, refer
// to the [StartStreamTranscription]API in the Amazon Transcribe Developer Guide.
//
// By default, Amazon Transcribe may use and store audio content processed by the
// service to develop and improve Amazon Web Services AI/ML services as further
// described in section 50 of the [Amazon Web Services Service Terms]. Using Amazon Transcribe may be subject to
// federal and state laws or regulations regarding the recording or interception of
// electronic communications. It is your and your end users’ responsibility to
// comply with all applicable laws regarding the recording, including properly
// notifying all participants in a recorded session or communication that the
// session or communication is being recorded, and obtaining all necessary
// consents. You can opt out from Amazon Web Services using audio content to
// develop and improve AWS AI/ML services by configuring an AI services opt out
// policy using Amazon Web Services Organizations.
//
// [Amazon Web Services Service Terms]: https://aws.amazon.com/service-terms/
// [Using Amazon Chime SDK live transcription]: https://docs.aws.amazon.com/chime-sdk/latest/dg/meeting-transcription.html
// [StartStreamTranscription]: https://docs.aws.amazon.com/transcribe/latest/APIReference/API_streaming_StartStreamTranscription.html
func chimesdkmeetings_StartMeetingTranscription(cfg aws.Config, client *chimesdkmeetings.Client) {
	input := &chimesdkmeetings.StartMeetingTranscriptionInput{
		// MeetingId: *string, // Required
		// TranscriptionConfiguration: *types.TranscriptionConfiguration, // Required
	}

	if len(_chimesdkmeetingsMeetingId) > 0 {
		input.MeetingId = aws.String(_chimesdkmeetingsMeetingId)
	}
	if len(_chimesdkmeetingsTranscriptionConfiguration) > 0 {
		if err := assignInputField(input, "TranscriptionConfiguration", _chimesdkmeetingsTranscriptionConfiguration); err != nil {
			log.Errorf("invalid --transcription-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartMeetingTranscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops transcription for the specified meetingId . For more information, refer to [Using Amazon Chime SDK live transcription]
// in the Amazon Chime SDK Developer Guide.
//
// By default, Amazon Transcribe may use and store audio content processed by the
// service to develop and improve Amazon Web Services AI/ML services as further
// described in section 50 of the [Amazon Web Services Service Terms]. Using Amazon Transcribe may be subject to
// federal and state laws or regulations regarding the recording or interception of
// electronic communications. It is your and your end users’ responsibility to
// comply with all applicable laws regarding the recording, including properly
// notifying all participants in a recorded session or communication that the
// session or communication is being recorded, and obtaining all necessary
// consents. You can opt out from Amazon Web Services using audio content to
// develop and improve Amazon Web Services AI/ML services by configuring an AI
// services opt out policy using Amazon Web Services Organizations.
//
// [Amazon Web Services Service Terms]: https://aws.amazon.com/service-terms/
// [Using Amazon Chime SDK live transcription]: https://docs.aws.amazon.com/chime-sdk/latest/dg/meeting-transcription.html
func chimesdkmeetings_StopMeetingTranscription(cfg aws.Config, client *chimesdkmeetings.Client) {
	input := &chimesdkmeetings.StopMeetingTranscriptionInput{
		// MeetingId: *string, // Required
	}

	if len(_chimesdkmeetingsMeetingId) > 0 {
		input.MeetingId = aws.String(_chimesdkmeetingsMeetingId)
	}

	if resp, err := client.StopMeetingTranscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The resource that supports tags.
func chimesdkmeetings_TagResource(cfg aws.Config, client *chimesdkmeetings.Client) {
	input := &chimesdkmeetings.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_chimesdkmeetingsResourceARN) > 0 {
		input.ResourceARN = aws.String(_chimesdkmeetingsResourceARN)
	}
	if len(_chimesdkmeetingsTags) > 0 {
		if err := assignInputField(input, "Tags", _chimesdkmeetingsTags); err != nil {
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

// Removes the specified tags from the specified resources. When you specify a tag
// key, the action removes both that key and its associated value. The operation
// succeeds even if you attempt to remove tags from a resource that were already
// removed. Note the following:
//
// - To remove tags from a resource, you need the necessary permissions for the
// service that the resource belongs to as well as permissions for removing tags.
// For more information, see the documentation for the service whose resource you
// want to untag.
//
// - You can only tag resources that are located in the specified Amazon Web
// Services Region for the calling Amazon Web Services account.
//
// # Minimum permissions
//
// In addition to the tag:UntagResources permission required by this operation,
// you must also have the remove tags permission defined by the service that
// created the resource. For example, to remove the tags from an Amazon EC2
// instance using the UntagResources operation, you must have both of the
// following permissions:
//
// tag:UntagResource
//
// ChimeSDKMeetings:DeleteTags
func chimesdkmeetings_UntagResource(cfg aws.Config, client *chimesdkmeetings.Client) {
	input := &chimesdkmeetings.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_chimesdkmeetingsResourceARN) > 0 {
		input.ResourceARN = aws.String(_chimesdkmeetingsResourceARN)
	}
	if len(_chimesdkmeetingsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _chimesdkmeetingsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The capabilities that you want to update.
// You use the capabilities with a set of values that control what the
// capabilities can do, such as SendReceive data. For more information about those
// values, see .
//
// When using capabilities, be aware of these corner cases:
//
// - If you specify MeetingFeatures:Video:MaxResolution:None when you create a
// meeting, all API requests that include SendReceive , Send , or Receive for
// AttendeeCapabilities:Video will be rejected with ValidationError 400 .
//
// - If you specify MeetingFeatures:Content:MaxResolution:None when you create a
// meeting, all API requests that include SendReceive , Send , or Receive for
// AttendeeCapabilities:Content will be rejected with ValidationError 400 .
//
// - You can't set content capabilities to SendReceive or Receive unless you also
// set video capabilities to SendReceive or Receive . If you don't set the video
// capability to receive, the response will contain an HTTP 400 Bad Request status
// code. However, you can set your video capability to receive and you set your
// content capability to not receive.
//
// - If meeting features is defined as Video:MaxResolution:None but
// Content:MaxResolution is defined as something other than None and attendee
// capabilities are not defined in the API request, then the default attendee video
// capability is set to Receive and attendee content capability is set to
// SendReceive . This is because content SendReceive requires video to be at
// least Receive .
//
// - When you change an audio capability from None or Receive to Send or
// SendReceive , and if the attendee left their microphone unmuted, audio will
// flow from the attendee to the other meeting participants.
//
// - When you change a video or content capability from None or Receive to Send
// or SendReceive , and if the attendee turned on their video or content streams,
// remote attendees can receive those streams, but only after media renegotiation
// between the client and the Amazon Chime back-end server.
func chimesdkmeetings_UpdateAttendeeCapabilities(cfg aws.Config, client *chimesdkmeetings.Client) {
	input := &chimesdkmeetings.UpdateAttendeeCapabilitiesInput{
		// AttendeeId: *string, // Required
		// Capabilities: *types.AttendeeCapabilities, // Required
		// MeetingId: *string, // Required
	}

	if len(_chimesdkmeetingsAttendeeId) > 0 {
		input.AttendeeId = aws.String(_chimesdkmeetingsAttendeeId)
	}
	if len(_chimesdkmeetingsCapabilities) > 0 {
		if err := assignInputField(input, "Capabilities", _chimesdkmeetingsCapabilities); err != nil {
			log.Errorf("invalid --capabilities: %s", err.Error())
			return
		}
	}
	if len(_chimesdkmeetingsMeetingId) > 0 {
		input.MeetingId = aws.String(_chimesdkmeetingsMeetingId)
	}

	if resp, err := client.UpdateAttendeeCapabilities(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_chimesdkmeetingsCmd)
	_chimesdkmeetingsCmd.Flags().SortFlags = false

	_chimesdkmeetingsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_chimesdkmeetingsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_chimesdkmeetingsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_chimesdkmeetingsCmd.Flags().StringVarP(&_chimesdkmeetingsAttendeeId, "attendee-id", "", "", "Attendee ID")
	_chimesdkmeetingsCmd.Flags().StringVarP(&_chimesdkmeetingsAttendees, "attendees", "", "", "Attendees")
	_chimesdkmeetingsCmd.Flags().StringVarP(&_chimesdkmeetingsCapabilities, "capabilities", "", "", "Capabilities")
	_chimesdkmeetingsCmd.Flags().StringVarP(&_chimesdkmeetingsClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_chimesdkmeetingsCmd.Flags().StringVarP(&_chimesdkmeetingsExcludedAttendeeIds, "excluded-attendee-ids", "", "", "Excluded Attendee Ids")
	_chimesdkmeetingsCmd.Flags().StringVarP(&_chimesdkmeetingsExternalMeetingId, "external-meeting-id", "", "", "External Meeting ID")
	_chimesdkmeetingsCmd.Flags().StringVarP(&_chimesdkmeetingsExternalUserId, "external-user-id", "", "", "External User ID")
	_chimesdkmeetingsCmd.Flags().StringVarP(&_chimesdkmeetingsMaxResults, "max-results", "", "", "Max Results")
	_chimesdkmeetingsCmd.Flags().StringVarP(&_chimesdkmeetingsMediaPlacementNetworkType, "media-placement-network-type", "", "", "Media Placement Network Type")
	_chimesdkmeetingsCmd.Flags().StringVarP(&_chimesdkmeetingsMediaRegion, "media-region", "", "", "Media Region")
	_chimesdkmeetingsCmd.Flags().StringVarP(&_chimesdkmeetingsMeetingFeatures, "meeting-features", "", "", "Meeting Features")
	_chimesdkmeetingsCmd.Flags().StringVarP(&_chimesdkmeetingsMeetingHostId, "meeting-host-id", "", "", "Meeting Host ID")
	_chimesdkmeetingsCmd.Flags().StringVarP(&_chimesdkmeetingsMeetingId, "meeting-id", "", "", "Meeting ID")
	_chimesdkmeetingsCmd.Flags().StringVarP(&_chimesdkmeetingsNextToken, "next-token", "", "", "Next Token")
	_chimesdkmeetingsCmd.Flags().StringVarP(&_chimesdkmeetingsNotificationsConfiguration, "notifications-configuration", "", "", "Notifications Configuration")
	_chimesdkmeetingsCmd.Flags().StringVarP(&_chimesdkmeetingsPrimaryMeetingId, "primary-meeting-id", "", "", "Primary Meeting ID")
	_chimesdkmeetingsCmd.Flags().StringVarP(&_chimesdkmeetingsResourceARN, "resource-arn", "", "", "Resource ARN")
	_chimesdkmeetingsCmd.Flags().StringSliceVarP(&_chimesdkmeetingsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_chimesdkmeetingsCmd.Flags().StringVarP(&_chimesdkmeetingsTags, "tags", "", "", "Tags")
	_chimesdkmeetingsCmd.Flags().StringSliceVarP(&_chimesdkmeetingsTenantIds, "tenant-ids", "", nil, "Tenant Ids")
	_chimesdkmeetingsCmd.Flags().StringVarP(&_chimesdkmeetingsTranscriptionConfiguration, "transcription-configuration", "", "", "Transcription Configuration")

	_chimesdkmeetingsCmd.Flags().BoolVarP(&_chimesdkmeetingsBatchCreateAttendee, "batch-create-attendee", "", false, "Batch Create Attendee")
	_chimesdkmeetingsCmd.Flags().BoolVarP(&_chimesdkmeetingsBatchUpdateAttendeeCapabilitiesExcept, "batch-update-attendee-capabilities-except", "", false, "Batch Update Attendee Capabilities Except")
	_chimesdkmeetingsCmd.Flags().BoolVarP(&_chimesdkmeetingsCreateAttendee, "create-attendee", "", false, "Create Attendee")
	_chimesdkmeetingsCmd.Flags().BoolVarP(&_chimesdkmeetingsCreateMeeting, "create-meeting", "", false, "Create Meeting")
	_chimesdkmeetingsCmd.Flags().BoolVarP(&_chimesdkmeetingsCreateMeetingWithAttendees, "create-meeting-with-attendees", "", false, "Create Meeting With Attendees")
	_chimesdkmeetingsCmd.Flags().BoolVarP(&_chimesdkmeetingsDeleteAttendee, "delete-attendee", "", false, "Delete Attendee")
	_chimesdkmeetingsCmd.Flags().BoolVarP(&_chimesdkmeetingsDeleteMeeting, "delete-meeting", "", false, "Delete Meeting")
	_chimesdkmeetingsCmd.Flags().BoolVarP(&_chimesdkmeetingsGetAttendee, "get-attendee", "", false, "Get Attendee")
	_chimesdkmeetingsCmd.Flags().BoolVarP(&_chimesdkmeetingsGetMeeting, "get-meeting", "", false, "Get Meeting")
	_chimesdkmeetingsCmd.Flags().BoolVarP(&_chimesdkmeetingsListAttendees, "list-attendees", "", false, "List Attendees")
	_chimesdkmeetingsCmd.Flags().BoolVarP(&_chimesdkmeetingsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_chimesdkmeetingsCmd.Flags().BoolVarP(&_chimesdkmeetingsStartMeetingTranscription, "start-meeting-transcription", "", false, "Start Meeting Transcription")
	_chimesdkmeetingsCmd.Flags().BoolVarP(&_chimesdkmeetingsStopMeetingTranscription, "stop-meeting-transcription", "", false, "Stop Meeting Transcription")
	_chimesdkmeetingsCmd.Flags().BoolVarP(&_chimesdkmeetingsTagResource, "tag-resource", "", false, "Tag Resource")
	_chimesdkmeetingsCmd.Flags().BoolVarP(&_chimesdkmeetingsUntagResource, "untag-resource", "", false, "Untag Resource")
	_chimesdkmeetingsCmd.Flags().BoolVarP(&_chimesdkmeetingsUpdateAttendeeCapabilities, "update-attendee-capabilities", "", false, "Update Attendee Capabilities")

}
