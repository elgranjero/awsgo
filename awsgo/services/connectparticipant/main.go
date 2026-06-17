package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/connectparticipant/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"cancel-participant-authentication", "complete-attachment-upload", "create-participant-connection", "describe-view", "disconnect-participant", "get-attachment", "get-authentication-url", "get-transcript", "send-event", "send-message", "start-attachment-upload"},
		OperationSet: map[string]bool{"cancel-participant-authentication": true, "complete-attachment-upload": true, "create-participant-connection": true, "describe-view": true, "disconnect-participant": true, "get-attachment": true, "get-authentication-url": true, "get-transcript": true, "send-event": true, "send-message": true, "start-attachment-upload": true},
		OperationInputs: map[string][]string{
			"cancel-participant-authentication": {"ConnectionToken", "SessionId"},
			"complete-attachment-upload":        {"AttachmentIds", "ClientToken", "ConnectionToken"},
			"create-participant-connection":     {"ConnectParticipant", "ParticipantToken", "Type"},
			"describe-view":                     {"ConnectionToken", "ViewToken"},
			"disconnect-participant":            {"ClientToken", "ConnectionToken"},
			"get-attachment":                    {"AttachmentId", "ConnectionToken", "UrlExpiryInSeconds"},
			"get-authentication-url":            {"ConnectionToken", "RedirectUri", "SessionId"},
			"get-transcript":                    {"ConnectionToken", "ContactId", "MaxResults", "NextToken", "ScanDirection", "SortOrder", "StartPosition"},
			"send-event":                        {"ClientToken", "ConnectionToken", "Content", "ContentType"},
			"send-message":                      {"ClientToken", "ConnectionToken", "Content", "ContentType"},
			"start-attachment-upload":           {"AttachmentName", "AttachmentSizeInBytes", "ClientToken", "ConnectionToken", "ContentType"},
		},
		OperationInputTypes: map[string]map[string]string{
			"cancel-participant-authentication": {"ConnectionToken": "*string", "SessionId": "*string"},
			"complete-attachment-upload":        {"AttachmentIds": "[]string", "ClientToken": "*string", "ConnectionToken": "*string"},
			"create-participant-connection":     {"ConnectParticipant": "*bool", "ParticipantToken": "*string", "Type": "[]types.ConnectionType"},
			"describe-view":                     {"ConnectionToken": "*string", "ViewToken": "*string"},
			"disconnect-participant":            {"ClientToken": "*string", "ConnectionToken": "*string"},
			"get-attachment":                    {"AttachmentId": "*string", "ConnectionToken": "*string", "UrlExpiryInSeconds": "*int32"},
			"get-authentication-url":            {"ConnectionToken": "*string", "RedirectUri": "*string", "SessionId": "*string"},
			"get-transcript":                    {"ConnectionToken": "*string", "ContactId": "*string", "MaxResults": "*int32", "NextToken": "*string", "ScanDirection": "types.ScanDirection", "SortOrder": "types.SortKey", "StartPosition": "*types.StartPosition"},
			"send-event":                        {"ClientToken": "*string", "ConnectionToken": "*string", "Content": "*string", "ContentType": "*string"},
			"send-message":                      {"ClientToken": "*string", "ConnectionToken": "*string", "Content": "*string", "ContentType": "*string"},
			"start-attachment-upload":           {"AttachmentName": "*string", "AttachmentSizeInBytes": "int64", "ClientToken": "*string", "ConnectionToken": "*string", "ContentType": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"cancel-participant-authentication": {"ConnectionToken", "SessionId"},
			"complete-attachment-upload":        {"AttachmentIds", "ClientToken", "ConnectionToken"},
			"create-participant-connection":     {"ParticipantToken"},
			"describe-view":                     {"ConnectionToken", "ViewToken"},
			"disconnect-participant":            {"ConnectionToken"},
			"get-attachment":                    {"AttachmentId", "ConnectionToken"},
			"get-authentication-url":            {"ConnectionToken", "RedirectUri", "SessionId"},
			"get-transcript":                    {"ConnectionToken"},
			"send-event":                        {"ConnectionToken", "ContentType"},
			"send-message":                      {"ConnectionToken", "Content", "ContentType"},
			"start-attachment-upload":           {"AttachmentName", "AttachmentSizeInBytes", "ClientToken", "ConnectionToken", "ContentType"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("connectparticipant", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
