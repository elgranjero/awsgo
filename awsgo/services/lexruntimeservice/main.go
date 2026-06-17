package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/lexruntimeservice/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"delete-session", "get-session", "post-content", "post-text", "put-session"},
		OperationSet: map[string]bool{"delete-session": true, "get-session": true, "post-content": true, "post-text": true, "put-session": true},
		OperationInputs: map[string][]string{
			"delete-session": {"BotAlias", "BotName", "UserId"},
			"get-session":    {"BotAlias", "BotName", "CheckpointLabelFilter", "UserId"},
			"post-content":   {"Accept", "ActiveContexts", "BotAlias", "BotName", "ContentType", "InputStream", "RequestAttributes", "SessionAttributes", "UserId"},
			"post-text":      {"ActiveContexts", "BotAlias", "BotName", "InputText", "RequestAttributes", "SessionAttributes", "UserId"},
			"put-session":    {"Accept", "ActiveContexts", "BotAlias", "BotName", "DialogAction", "RecentIntentSummaryView", "SessionAttributes", "UserId"},
		},
		OperationInputTypes: map[string]map[string]string{
			"delete-session": {"BotAlias": "*string", "BotName": "*string", "UserId": "*string"},
			"get-session":    {"BotAlias": "*string", "BotName": "*string", "CheckpointLabelFilter": "*string", "UserId": "*string"},
			"post-content":   {"Accept": "*string", "ActiveContexts": "*string", "BotAlias": "*string", "BotName": "*string", "ContentType": "*string", "InputStream": "io.Reader", "RequestAttributes": "*string", "SessionAttributes": "*string", "UserId": "*string"},
			"post-text":      {"ActiveContexts": "[]types.ActiveContext", "BotAlias": "*string", "BotName": "*string", "InputText": "*string", "RequestAttributes": "map[string]string", "SessionAttributes": "map[string]string", "UserId": "*string"},
			"put-session":    {"Accept": "*string", "ActiveContexts": "[]types.ActiveContext", "BotAlias": "*string", "BotName": "*string", "DialogAction": "*types.DialogAction", "RecentIntentSummaryView": "[]types.IntentSummary", "SessionAttributes": "map[string]string", "UserId": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"delete-session": {"BotAlias", "BotName", "UserId"},
			"get-session":    {"BotAlias", "BotName", "UserId"},
			"post-content":   {"BotAlias", "BotName", "ContentType", "InputStream", "UserId"},
			"post-text":      {"BotAlias", "BotName", "InputText", "UserId"},
			"put-session":    {"BotAlias", "BotName", "UserId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("lexruntimeservice", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
