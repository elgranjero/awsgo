package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/lexruntimev2/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"delete-session", "get-session", "put-session", "recognize-text", "recognize-utterance", "start-conversation"},
		OperationSet: map[string]bool{"delete-session": true, "get-session": true, "put-session": true, "recognize-text": true, "recognize-utterance": true, "start-conversation": true},
		OperationInputs: map[string][]string{
			"delete-session":      {"BotAliasId", "BotId", "LocaleId", "SessionId"},
			"get-session":         {"BotAliasId", "BotId", "LocaleId", "SessionId"},
			"put-session":         {"BotAliasId", "BotId", "LocaleId", "Messages", "RequestAttributes", "ResponseContentType", "SessionId", "SessionState"},
			"recognize-text":      {"BotAliasId", "BotId", "LocaleId", "RequestAttributes", "SessionId", "SessionState", "Text"},
			"recognize-utterance": {"BotAliasId", "BotId", "InputStream", "LocaleId", "RequestAttributes", "RequestContentType", "ResponseContentType", "SessionId", "SessionState"},
			"start-conversation":  {"BotAliasId", "BotId", "ConversationMode", "LocaleId", "SessionId"},
		},
		OperationInputTypes: map[string]map[string]string{
			"delete-session":      {"BotAliasId": "*string", "BotId": "*string", "LocaleId": "*string", "SessionId": "*string"},
			"get-session":         {"BotAliasId": "*string", "BotId": "*string", "LocaleId": "*string", "SessionId": "*string"},
			"put-session":         {"BotAliasId": "*string", "BotId": "*string", "LocaleId": "*string", "Messages": "[]types.Message", "RequestAttributes": "map[string]string", "ResponseContentType": "*string", "SessionId": "*string", "SessionState": "*types.SessionState"},
			"recognize-text":      {"BotAliasId": "*string", "BotId": "*string", "LocaleId": "*string", "RequestAttributes": "map[string]string", "SessionId": "*string", "SessionState": "*types.SessionState", "Text": "*string"},
			"recognize-utterance": {"BotAliasId": "*string", "BotId": "*string", "InputStream": "io.Reader", "LocaleId": "*string", "RequestAttributes": "*string", "RequestContentType": "*string", "ResponseContentType": "*string", "SessionId": "*string", "SessionState": "*string"},
			"start-conversation":  {"BotAliasId": "*string", "BotId": "*string", "ConversationMode": "types.ConversationMode", "LocaleId": "*string", "SessionId": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"delete-session":      {"BotAliasId", "BotId", "LocaleId", "SessionId"},
			"get-session":         {"BotAliasId", "BotId", "LocaleId", "SessionId"},
			"put-session":         {"BotAliasId", "BotId", "LocaleId", "SessionId", "SessionState"},
			"recognize-text":      {"BotAliasId", "BotId", "LocaleId", "SessionId", "Text"},
			"recognize-utterance": {"BotAliasId", "BotId", "LocaleId", "RequestContentType", "SessionId"},
			"start-conversation":  {"BotAliasId", "BotId", "LocaleId", "SessionId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("lexruntimev2", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
