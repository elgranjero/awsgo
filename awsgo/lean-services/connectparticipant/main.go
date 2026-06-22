package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/connectparticipant"
)

var fields_cancel_participant_authentication = []leanruntime.Field{
	{Name: "ConnectionToken", Flag: "connection-token", Type: "*string", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_complete_attachment_upload = []leanruntime.Field{
	{Name: "AttachmentIds", Flag: "attachment-ids", Type: "[]string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "ConnectionToken", Flag: "connection-token", Type: "*string", Required: true},
}

var fields_create_participant_connection = []leanruntime.Field{
	{Name: "ConnectParticipant", Flag: "connect-participant", Type: "*bool", Required: false},
	{Name: "ParticipantToken", Flag: "participant-token", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "[]types.ConnectionType", Required: false},
}

var fields_describe_view = []leanruntime.Field{
	{Name: "ConnectionToken", Flag: "connection-token", Type: "*string", Required: true},
	{Name: "ViewToken", Flag: "view-token", Type: "*string", Required: true},
}

var fields_disconnect_participant = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ConnectionToken", Flag: "connection-token", Type: "*string", Required: true},
}

var fields_get_attachment = []leanruntime.Field{
	{Name: "AttachmentId", Flag: "attachment-id", Type: "*string", Required: true},
	{Name: "ConnectionToken", Flag: "connection-token", Type: "*string", Required: true},
	{Name: "UrlExpiryInSeconds", Flag: "url-expiry-in-seconds", Type: "*int32", Required: false},
}

var fields_get_authentication_url = []leanruntime.Field{
	{Name: "ConnectionToken", Flag: "connection-token", Type: "*string", Required: true},
	{Name: "RedirectUri", Flag: "redirect-uri", Type: "*string", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_get_transcript = []leanruntime.Field{
	{Name: "ConnectionToken", Flag: "connection-token", Type: "*string", Required: true},
	{Name: "ContactId", Flag: "contact-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ScanDirection", Flag: "scan-direction", Type: "types.ScanDirection", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortKey", Required: false},
	{Name: "StartPosition", Flag: "start-position", Type: "*types.StartPosition", Required: false},
}

var fields_send_event = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ConnectionToken", Flag: "connection-token", Type: "*string", Required: true},
	{Name: "Content", Flag: "content", Type: "*string", Required: false},
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: true},
}

var fields_send_message = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "ConnectionToken", Flag: "connection-token", Type: "*string", Required: true},
	{Name: "Content", Flag: "content", Type: "*string", Required: true},
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: true},
}

var fields_start_attachment_upload = []leanruntime.Field{
	{Name: "AttachmentName", Flag: "attachment-name", Type: "*string", Required: true},
	{Name: "AttachmentSizeInBytes", Flag: "attachment-size-in-bytes", Type: "int64", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "ConnectionToken", Flag: "connection-token", Type: "*string", Required: true},
	{Name: "ContentType", Flag: "content-type", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"cancel-participant-authentication": {
			Name:   "cancel-participant-authentication",
			Fields: fields_cancel_participant_authentication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelParticipantAuthenticationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_participant_authentication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelParticipantAuthentication(ctx, input)
			},
		},
		"complete-attachment-upload": {
			Name:   "complete-attachment-upload",
			Fields: fields_complete_attachment_upload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CompleteAttachmentUploadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_complete_attachment_upload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CompleteAttachmentUpload(ctx, input)
			},
		},
		"create-participant-connection": {
			Name:   "create-participant-connection",
			Fields: fields_create_participant_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateParticipantConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_participant_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateParticipantConnection(ctx, input)
			},
		},
		"describe-view": {
			Name:   "describe-view",
			Fields: fields_describe_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeViewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_view, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeView(ctx, input)
			},
		},
		"disconnect-participant": {
			Name:   "disconnect-participant",
			Fields: fields_disconnect_participant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisconnectParticipantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disconnect_participant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisconnectParticipant(ctx, input)
			},
		},
		"get-attachment": {
			Name:   "get-attachment",
			Fields: fields_get_attachment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAttachmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_attachment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAttachment(ctx, input)
			},
		},
		"get-authentication-url": {
			Name:   "get-authentication-url",
			Fields: fields_get_authentication_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAuthenticationUrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_authentication_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAuthenticationUrl(ctx, input)
			},
		},
		"get-transcript": {
			Name:   "get-transcript",
			Fields: fields_get_transcript,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTranscriptInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_transcript, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetTranscript(ctx, input)
				}
				var results []*svc.GetTranscriptOutput
				p := svc.NewGetTranscriptPaginator(client, input)
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
		"send-event": {
			Name:   "send-event",
			Fields: fields_send_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendEvent(ctx, input)
			},
		},
		"send-message": {
			Name:   "send-message",
			Fields: fields_send_message,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendMessageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_message, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendMessage(ctx, input)
			},
		},
		"start-attachment-upload": {
			Name:   "start-attachment-upload",
			Fields: fields_start_attachment_upload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartAttachmentUploadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_attachment_upload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartAttachmentUpload(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("connectparticipant", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
