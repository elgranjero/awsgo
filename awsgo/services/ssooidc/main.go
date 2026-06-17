package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/ssooidc/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-token", "create-token-with-iam", "register-client", "start-device-authorization"},
		OperationSet: map[string]bool{"create-token": true, "create-token-with-iam": true, "register-client": true, "start-device-authorization": true},
		OperationInputs: map[string][]string{
			"create-token":               {"ClientId", "ClientSecret", "Code", "CodeVerifier", "DeviceCode", "GrantType", "RedirectUri", "RefreshToken", "Scope"},
			"create-token-with-iam":      {"Assertion", "ClientId", "Code", "CodeVerifier", "GrantType", "RedirectUri", "RefreshToken", "RequestedTokenType", "Scope", "SubjectToken", "SubjectTokenType"},
			"register-client":            {"ClientName", "ClientType", "EntitledApplicationArn", "GrantTypes", "IssuerUrl", "RedirectUris", "Scopes"},
			"start-device-authorization": {"ClientId", "ClientSecret", "StartUrl"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-token":               {"ClientId": "*string", "ClientSecret": "*string", "Code": "*string", "CodeVerifier": "*string", "DeviceCode": "*string", "GrantType": "*string", "RedirectUri": "*string", "RefreshToken": "*string", "Scope": "[]string"},
			"create-token-with-iam":      {"Assertion": "*string", "ClientId": "*string", "Code": "*string", "CodeVerifier": "*string", "GrantType": "*string", "RedirectUri": "*string", "RefreshToken": "*string", "RequestedTokenType": "*string", "Scope": "[]string", "SubjectToken": "*string", "SubjectTokenType": "*string"},
			"register-client":            {"ClientName": "*string", "ClientType": "*string", "EntitledApplicationArn": "*string", "GrantTypes": "[]string", "IssuerUrl": "*string", "RedirectUris": "[]string", "Scopes": "[]string"},
			"start-device-authorization": {"ClientId": "*string", "ClientSecret": "*string", "StartUrl": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"create-token":               {"ClientId", "ClientSecret", "GrantType"},
			"create-token-with-iam":      {"ClientId", "GrantType"},
			"register-client":            {"ClientName", "ClientType"},
			"start-device-authorization": {"ClientId", "ClientSecret", "StartUrl"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("ssooidc", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
