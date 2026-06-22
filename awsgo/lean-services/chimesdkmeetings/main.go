package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/chimesdkmeetings"
)

var fields_batch_create_attendee = []leanruntime.Field{
	{Name: "Attendees", Flag: "attendees", Type: "[]types.CreateAttendeeRequestItem", Required: true},
	{Name: "MeetingId", Flag: "meeting-id", Type: "*string", Required: true},
}

var fields_batch_update_attendee_capabilities_except = []leanruntime.Field{
	{Name: "Capabilities", Flag: "capabilities", Type: "*types.AttendeeCapabilities", Required: true},
	{Name: "ExcludedAttendeeIds", Flag: "excluded-attendee-ids", Type: "[]types.AttendeeIdItem", Required: true},
	{Name: "MeetingId", Flag: "meeting-id", Type: "*string", Required: true},
}

var fields_create_attendee = []leanruntime.Field{
	{Name: "Capabilities", Flag: "capabilities", Type: "*types.AttendeeCapabilities", Required: false},
	{Name: "ExternalUserId", Flag: "external-user-id", Type: "*string", Required: true},
	{Name: "MeetingId", Flag: "meeting-id", Type: "*string", Required: true},
}

var fields_create_meeting = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "ExternalMeetingId", Flag: "external-meeting-id", Type: "*string", Required: true},
	{Name: "MediaPlacementNetworkType", Flag: "media-placement-network-type", Type: "types.MediaPlacementNetworkType", Required: false},
	{Name: "MediaRegion", Flag: "media-region", Type: "*string", Required: true},
	{Name: "MeetingFeatures", Flag: "meeting-features", Type: "*types.MeetingFeaturesConfiguration", Required: false},
	{Name: "MeetingHostId", Flag: "meeting-host-id", Type: "*string", Required: false},
	{Name: "NotificationsConfiguration", Flag: "notifications-configuration", Type: "*types.NotificationsConfiguration", Required: false},
	{Name: "PrimaryMeetingId", Flag: "primary-meeting-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TenantIds", Flag: "tenant-ids", Type: "[]string", Required: false},
}

var fields_create_meeting_with_attendees = []leanruntime.Field{
	{Name: "Attendees", Flag: "attendees", Type: "[]types.CreateAttendeeRequestItem", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "ExternalMeetingId", Flag: "external-meeting-id", Type: "*string", Required: true},
	{Name: "MediaPlacementNetworkType", Flag: "media-placement-network-type", Type: "types.MediaPlacementNetworkType", Required: false},
	{Name: "MediaRegion", Flag: "media-region", Type: "*string", Required: true},
	{Name: "MeetingFeatures", Flag: "meeting-features", Type: "*types.MeetingFeaturesConfiguration", Required: false},
	{Name: "MeetingHostId", Flag: "meeting-host-id", Type: "*string", Required: false},
	{Name: "NotificationsConfiguration", Flag: "notifications-configuration", Type: "*types.NotificationsConfiguration", Required: false},
	{Name: "PrimaryMeetingId", Flag: "primary-meeting-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TenantIds", Flag: "tenant-ids", Type: "[]string", Required: false},
}

var fields_delete_attendee = []leanruntime.Field{
	{Name: "AttendeeId", Flag: "attendee-id", Type: "*string", Required: true},
	{Name: "MeetingId", Flag: "meeting-id", Type: "*string", Required: true},
}

var fields_delete_meeting = []leanruntime.Field{
	{Name: "MeetingId", Flag: "meeting-id", Type: "*string", Required: true},
}

var fields_get_attendee = []leanruntime.Field{
	{Name: "AttendeeId", Flag: "attendee-id", Type: "*string", Required: true},
	{Name: "MeetingId", Flag: "meeting-id", Type: "*string", Required: true},
}

var fields_get_meeting = []leanruntime.Field{
	{Name: "MeetingId", Flag: "meeting-id", Type: "*string", Required: true},
}

var fields_list_attendees = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MeetingId", Flag: "meeting-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_start_meeting_transcription = []leanruntime.Field{
	{Name: "MeetingId", Flag: "meeting-id", Type: "*string", Required: true},
	{Name: "TranscriptionConfiguration", Flag: "transcription-configuration", Type: "*types.TranscriptionConfiguration", Required: true},
}

var fields_stop_meeting_transcription = []leanruntime.Field{
	{Name: "MeetingId", Flag: "meeting-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_attendee_capabilities = []leanruntime.Field{
	{Name: "AttendeeId", Flag: "attendee-id", Type: "*string", Required: true},
	{Name: "Capabilities", Flag: "capabilities", Type: "*types.AttendeeCapabilities", Required: true},
	{Name: "MeetingId", Flag: "meeting-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-create-attendee": {
			Name:   "batch-create-attendee",
			Fields: fields_batch_create_attendee,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchCreateAttendeeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_create_attendee, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchCreateAttendee(ctx, input)
			},
		},
		"batch-update-attendee-capabilities-except": {
			Name:   "batch-update-attendee-capabilities-except",
			Fields: fields_batch_update_attendee_capabilities_except,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchUpdateAttendeeCapabilitiesExceptInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_update_attendee_capabilities_except, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchUpdateAttendeeCapabilitiesExcept(ctx, input)
			},
		},
		"create-attendee": {
			Name:   "create-attendee",
			Fields: fields_create_attendee,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAttendeeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_attendee, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAttendee(ctx, input)
			},
		},
		"create-meeting": {
			Name:   "create-meeting",
			Fields: fields_create_meeting,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMeetingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_meeting, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMeeting(ctx, input)
			},
		},
		"create-meeting-with-attendees": {
			Name:   "create-meeting-with-attendees",
			Fields: fields_create_meeting_with_attendees,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMeetingWithAttendeesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_meeting_with_attendees, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMeetingWithAttendees(ctx, input)
			},
		},
		"delete-attendee": {
			Name:   "delete-attendee",
			Fields: fields_delete_attendee,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAttendeeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_attendee, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAttendee(ctx, input)
			},
		},
		"delete-meeting": {
			Name:   "delete-meeting",
			Fields: fields_delete_meeting,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMeetingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_meeting, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMeeting(ctx, input)
			},
		},
		"get-attendee": {
			Name:   "get-attendee",
			Fields: fields_get_attendee,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAttendeeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_attendee, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAttendee(ctx, input)
			},
		},
		"get-meeting": {
			Name:   "get-meeting",
			Fields: fields_get_meeting,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMeetingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_meeting, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMeeting(ctx, input)
			},
		},
		"list-attendees": {
			Name:   "list-attendees",
			Fields: fields_list_attendees,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAttendeesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_attendees, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAttendees(ctx, input)
				}
				var results []*svc.ListAttendeesOutput
				p := svc.NewListAttendeesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-tags-for-resource": {
			Name:   "list-tags-for-resource",
			Fields: fields_list_tags_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForResource(ctx, input)
			},
		},
		"start-meeting-transcription": {
			Name:   "start-meeting-transcription",
			Fields: fields_start_meeting_transcription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMeetingTranscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_meeting_transcription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMeetingTranscription(ctx, input)
			},
		},
		"stop-meeting-transcription": {
			Name:   "stop-meeting-transcription",
			Fields: fields_stop_meeting_transcription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopMeetingTranscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_meeting_transcription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopMeetingTranscription(ctx, input)
			},
		},
		"tag-resource": {
			Name:   "tag-resource",
			Fields: fields_tag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagResource(ctx, input)
			},
		},
		"untag-resource": {
			Name:   "untag-resource",
			Fields: fields_untag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagResource(ctx, input)
			},
		},
		"update-attendee-capabilities": {
			Name:   "update-attendee-capabilities",
			Fields: fields_update_attendee_capabilities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAttendeeCapabilitiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_attendee_capabilities, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAttendeeCapabilities(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("chimesdkmeetings", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
