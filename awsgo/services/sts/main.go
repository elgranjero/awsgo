package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/sts/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"assume-role", "assume-role-with-saml", "assume-role-with-web-identity", "assume-root", "decode-authorization-message", "get-access-key-info", "get-caller-identity", "get-delegated-access-token", "get-federation-token", "get-session-token", "get-web-identity-token"},
		OperationSet: map[string]bool{"assume-role": true, "assume-role-with-saml": true, "assume-role-with-web-identity": true, "assume-root": true, "decode-authorization-message": true, "get-access-key-info": true, "get-caller-identity": true, "get-delegated-access-token": true, "get-federation-token": true, "get-session-token": true, "get-web-identity-token": true},
		OperationInputs: map[string][]string{
			"assume-role":                   {"DurationSeconds", "ExternalId", "Policy", "PolicyArns", "ProvidedContexts", "RoleArn", "RoleSessionName", "SerialNumber", "SourceIdentity", "Tags", "TokenCode", "TransitiveTagKeys"},
			"assume-role-with-saml":         {"DurationSeconds", "Policy", "PolicyArns", "PrincipalArn", "RoleArn", "SAMLAssertion"},
			"assume-role-with-web-identity": {"DurationSeconds", "Policy", "PolicyArns", "ProviderId", "RoleArn", "RoleSessionName", "WebIdentityToken"},
			"assume-root":                   {"DurationSeconds", "TargetPrincipal", "TaskPolicyArn"},
			"decode-authorization-message":  {"EncodedMessage"},
			"get-access-key-info":           {"AccessKeyId"},
			"get-caller-identity":           {},
			"get-delegated-access-token":    {"TradeInToken"},
			"get-federation-token":          {"DurationSeconds", "Name", "Policy", "PolicyArns", "Tags"},
			"get-session-token":             {"DurationSeconds", "SerialNumber", "TokenCode"},
			"get-web-identity-token":        {"Audience", "DurationSeconds", "SigningAlgorithm", "Tags"},
		},
		OperationInputTypes: map[string]map[string]string{
			"assume-role":                   {"DurationSeconds": "*int32", "ExternalId": "*string", "Policy": "*string", "PolicyArns": "[]types.PolicyDescriptorType", "ProvidedContexts": "[]types.ProvidedContext", "RoleArn": "*string", "RoleSessionName": "*string", "SerialNumber": "*string", "SourceIdentity": "*string", "Tags": "[]types.Tag", "TokenCode": "*string", "TransitiveTagKeys": "[]string"},
			"assume-role-with-saml":         {"DurationSeconds": "*int32", "Policy": "*string", "PolicyArns": "[]types.PolicyDescriptorType", "PrincipalArn": "*string", "RoleArn": "*string", "SAMLAssertion": "*string"},
			"assume-role-with-web-identity": {"DurationSeconds": "*int32", "Policy": "*string", "PolicyArns": "[]types.PolicyDescriptorType", "ProviderId": "*string", "RoleArn": "*string", "RoleSessionName": "*string", "WebIdentityToken": "*string"},
			"assume-root":                   {"DurationSeconds": "*int32", "TargetPrincipal": "*string", "TaskPolicyArn": "*types.PolicyDescriptorType"},
			"decode-authorization-message":  {"EncodedMessage": "*string"},
			"get-access-key-info":           {"AccessKeyId": "*string"},
			"get-caller-identity":           {},
			"get-delegated-access-token":    {"TradeInToken": "*string"},
			"get-federation-token":          {"DurationSeconds": "*int32", "Name": "*string", "Policy": "*string", "PolicyArns": "[]types.PolicyDescriptorType", "Tags": "[]types.Tag"},
			"get-session-token":             {"DurationSeconds": "*int32", "SerialNumber": "*string", "TokenCode": "*string"},
			"get-web-identity-token":        {"Audience": "[]string", "DurationSeconds": "*int32", "SigningAlgorithm": "*string", "Tags": "[]types.Tag"},
		},
		OperationInputRequired: map[string][]string{
			"assume-role":                   {"RoleArn", "RoleSessionName"},
			"assume-role-with-saml":         {"PrincipalArn", "RoleArn", "SAMLAssertion"},
			"assume-role-with-web-identity": {"RoleArn", "RoleSessionName", "WebIdentityToken"},
			"assume-root":                   {"TargetPrincipal", "TaskPolicyArn"},
			"decode-authorization-message":  {"EncodedMessage"},
			"get-access-key-info":           {"AccessKeyId"},
			"get-caller-identity":           {},
			"get-delegated-access-token":    {"TradeInToken"},
			"get-federation-token":          {"Name"},
			"get-session-token":             {},
			"get-web-identity-token":        {"Audience", "SigningAlgorithm"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("sts", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
