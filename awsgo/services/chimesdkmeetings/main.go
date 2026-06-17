package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/chimesdkmeetings/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"batch-create-attendee", "batch-update-attendee-capabilities-except", "create-attendee", "create-meeting", "create-meeting-with-attendees", "delete-attendee", "delete-meeting", "get-attendee", "get-meeting", "list-attendees", "list-tags-for-resource", "start-meeting-transcription", "stop-meeting-transcription", "tag-resource", "untag-resource", "update-attendee-capabilities"},
		OperationSet: map[string]bool{"batch-create-attendee": true, "batch-update-attendee-capabilities-except": true, "create-attendee": true, "create-meeting": true, "create-meeting-with-attendees": true, "delete-attendee": true, "delete-meeting": true, "get-attendee": true, "get-meeting": true, "list-attendees": true, "list-tags-for-resource": true, "start-meeting-transcription": true, "stop-meeting-transcription": true, "tag-resource": true, "untag-resource": true, "update-attendee-capabilities": true},
		OperationInputs: map[string][]string{
			"batch-create-attendee":                     {"Attendees", "MeetingId"},
			"batch-update-attendee-capabilities-except": {"Capabilities", "ExcludedAttendeeIds", "MeetingId"},
			"create-attendee":                           {"Capabilities", "ExternalUserId", "MeetingId"},
			"create-meeting":                            {"ClientRequestToken", "ExternalMeetingId", "MediaPlacementNetworkType", "MediaRegion", "MeetingFeatures", "MeetingHostId", "NotificationsConfiguration", "PrimaryMeetingId", "Tags", "TenantIds"},
			"create-meeting-with-attendees":             {"Attendees", "ClientRequestToken", "ExternalMeetingId", "MediaPlacementNetworkType", "MediaRegion", "MeetingFeatures", "MeetingHostId", "NotificationsConfiguration", "PrimaryMeetingId", "Tags", "TenantIds"},
			"delete-attendee":                           {"AttendeeId", "MeetingId"},
			"delete-meeting":                            {"MeetingId"},
			"get-attendee":                              {"AttendeeId", "MeetingId"},
			"get-meeting":                               {"MeetingId"},
			"list-attendees":                            {"MaxResults", "MeetingId", "NextToken"},
			"list-tags-for-resource":                    {"ResourceARN"},
			"start-meeting-transcription":               {"MeetingId", "TranscriptionConfiguration"},
			"stop-meeting-transcription":                {"MeetingId"},
			"tag-resource":                              {"ResourceARN", "Tags"},
			"untag-resource":                            {"ResourceARN", "TagKeys"},
			"update-attendee-capabilities":              {"AttendeeId", "Capabilities", "MeetingId"},
		},
		OperationInputTypes: map[string]map[string]string{
			"batch-create-attendee":                     {"Attendees": "[]types.CreateAttendeeRequestItem", "MeetingId": "*string"},
			"batch-update-attendee-capabilities-except": {"Capabilities": "*types.AttendeeCapabilities", "ExcludedAttendeeIds": "[]types.AttendeeIdItem", "MeetingId": "*string"},
			"create-attendee":                           {"Capabilities": "*types.AttendeeCapabilities", "ExternalUserId": "*string", "MeetingId": "*string"},
			"create-meeting":                            {"ClientRequestToken": "*string", "ExternalMeetingId": "*string", "MediaPlacementNetworkType": "types.MediaPlacementNetworkType", "MediaRegion": "*string", "MeetingFeatures": "*types.MeetingFeaturesConfiguration", "MeetingHostId": "*string", "NotificationsConfiguration": "*types.NotificationsConfiguration", "PrimaryMeetingId": "*string", "Tags": "[]types.Tag", "TenantIds": "[]string"},
			"create-meeting-with-attendees":             {"Attendees": "[]types.CreateAttendeeRequestItem", "ClientRequestToken": "*string", "ExternalMeetingId": "*string", "MediaPlacementNetworkType": "types.MediaPlacementNetworkType", "MediaRegion": "*string", "MeetingFeatures": "*types.MeetingFeaturesConfiguration", "MeetingHostId": "*string", "NotificationsConfiguration": "*types.NotificationsConfiguration", "PrimaryMeetingId": "*string", "Tags": "[]types.Tag", "TenantIds": "[]string"},
			"delete-attendee":                           {"AttendeeId": "*string", "MeetingId": "*string"},
			"delete-meeting":                            {"MeetingId": "*string"},
			"get-attendee":                              {"AttendeeId": "*string", "MeetingId": "*string"},
			"get-meeting":                               {"MeetingId": "*string"},
			"list-attendees":                            {"MaxResults": "*int32", "MeetingId": "*string", "NextToken": "*string"},
			"list-tags-for-resource":                    {"ResourceARN": "*string"},
			"start-meeting-transcription":               {"MeetingId": "*string", "TranscriptionConfiguration": "*types.TranscriptionConfiguration"},
			"stop-meeting-transcription":                {"MeetingId": "*string"},
			"tag-resource":                              {"ResourceARN": "*string", "Tags": "[]types.Tag"},
			"untag-resource":                            {"ResourceARN": "*string", "TagKeys": "[]string"},
			"update-attendee-capabilities":              {"AttendeeId": "*string", "Capabilities": "*types.AttendeeCapabilities", "MeetingId": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"batch-create-attendee":                     {"Attendees", "MeetingId"},
			"batch-update-attendee-capabilities-except": {"Capabilities", "ExcludedAttendeeIds", "MeetingId"},
			"create-attendee":                           {"ExternalUserId", "MeetingId"},
			"create-meeting":                            {"ClientRequestToken", "ExternalMeetingId", "MediaRegion"},
			"create-meeting-with-attendees":             {"Attendees", "ClientRequestToken", "ExternalMeetingId", "MediaRegion"},
			"delete-attendee":                           {"AttendeeId", "MeetingId"},
			"delete-meeting":                            {"MeetingId"},
			"get-attendee":                              {"AttendeeId", "MeetingId"},
			"get-meeting":                               {"MeetingId"},
			"list-attendees":                            {"MeetingId"},
			"list-tags-for-resource":                    {"ResourceARN"},
			"start-meeting-transcription":               {"MeetingId", "TranscriptionConfiguration"},
			"stop-meeting-transcription":                {"MeetingId"},
			"tag-resource":                              {"ResourceARN", "Tags"},
			"untag-resource":                            {"ResourceARN", "TagKeys"},
			"update-attendee-capabilities":              {"AttendeeId", "Capabilities", "MeetingId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("chimesdkmeetings", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
