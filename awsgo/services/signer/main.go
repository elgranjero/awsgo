package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/signer/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"add-profile-permission", "cancel-signing-profile", "describe-signing-job", "get-revocation-status", "get-signing-platform", "get-signing-profile", "list-profile-permissions", "list-signing-jobs", "list-signing-platforms", "list-signing-profiles", "list-tags-for-resource", "put-signing-profile", "remove-profile-permission", "revoke-signature", "revoke-signing-profile", "sign-payload", "start-signing-job", "tag-resource", "untag-resource"},
		OperationSet: map[string]bool{"add-profile-permission": true, "cancel-signing-profile": true, "describe-signing-job": true, "get-revocation-status": true, "get-signing-platform": true, "get-signing-profile": true, "list-profile-permissions": true, "list-signing-jobs": true, "list-signing-platforms": true, "list-signing-profiles": true, "list-tags-for-resource": true, "put-signing-profile": true, "remove-profile-permission": true, "revoke-signature": true, "revoke-signing-profile": true, "sign-payload": true, "start-signing-job": true, "tag-resource": true, "untag-resource": true},
		OperationInputs: map[string][]string{
			"add-profile-permission":    {"Action", "Principal", "ProfileName", "ProfileVersion", "RevisionId", "StatementId"},
			"cancel-signing-profile":    {"ProfileName"},
			"describe-signing-job":      {"JobId"},
			"get-revocation-status":     {"CertificateHashes", "JobArn", "PlatformId", "ProfileVersionArn", "SignatureTimestamp"},
			"get-signing-platform":      {"PlatformId"},
			"get-signing-profile":       {"ProfileName", "ProfileOwner"},
			"list-profile-permissions":  {"NextToken", "ProfileName"},
			"list-signing-jobs":         {"IsRevoked", "JobInvoker", "MaxResults", "NextToken", "PlatformId", "RequestedBy", "SignatureExpiresAfter", "SignatureExpiresBefore", "Status"},
			"list-signing-platforms":    {"Category", "MaxResults", "NextToken", "Partner", "Target"},
			"list-signing-profiles":     {"IncludeCanceled", "MaxResults", "NextToken", "PlatformId", "Statuses"},
			"list-tags-for-resource":    {"ResourceArn"},
			"put-signing-profile":       {"Overrides", "PlatformId", "ProfileName", "SignatureValidityPeriod", "SigningMaterial", "SigningParameters", "Tags"},
			"remove-profile-permission": {"ProfileName", "RevisionId", "StatementId"},
			"revoke-signature":          {"JobId", "JobOwner", "Reason"},
			"revoke-signing-profile":    {"EffectiveTime", "ProfileName", "ProfileVersion", "Reason"},
			"sign-payload":              {"Payload", "PayloadFormat", "ProfileName", "ProfileOwner"},
			"start-signing-job":         {"ClientRequestToken", "Destination", "ProfileName", "ProfileOwner", "Source"},
			"tag-resource":              {"ResourceArn", "Tags"},
			"untag-resource":            {"ResourceArn", "TagKeys"},
		},
		OperationInputTypes: map[string]map[string]string{
			"add-profile-permission":    {"Action": "*string", "Principal": "*string", "ProfileName": "*string", "ProfileVersion": "*string", "RevisionId": "*string", "StatementId": "*string"},
			"cancel-signing-profile":    {"ProfileName": "*string"},
			"describe-signing-job":      {"JobId": "*string"},
			"get-revocation-status":     {"CertificateHashes": "[]string", "JobArn": "*string", "PlatformId": "*string", "ProfileVersionArn": "*string", "SignatureTimestamp": "*time.Time"},
			"get-signing-platform":      {"PlatformId": "*string"},
			"get-signing-profile":       {"ProfileName": "*string", "ProfileOwner": "*string"},
			"list-profile-permissions":  {"NextToken": "*string", "ProfileName": "*string"},
			"list-signing-jobs":         {"IsRevoked": "bool", "JobInvoker": "*string", "MaxResults": "*int32", "NextToken": "*string", "PlatformId": "*string", "RequestedBy": "*string", "SignatureExpiresAfter": "*time.Time", "SignatureExpiresBefore": "*time.Time", "Status": "types.SigningStatus"},
			"list-signing-platforms":    {"Category": "*string", "MaxResults": "*int32", "NextToken": "*string", "Partner": "*string", "Target": "*string"},
			"list-signing-profiles":     {"IncludeCanceled": "bool", "MaxResults": "*int32", "NextToken": "*string", "PlatformId": "*string", "Statuses": "[]types.SigningProfileStatus"},
			"list-tags-for-resource":    {"ResourceArn": "*string"},
			"put-signing-profile":       {"Overrides": "*types.SigningPlatformOverrides", "PlatformId": "*string", "ProfileName": "*string", "SignatureValidityPeriod": "*types.SignatureValidityPeriod", "SigningMaterial": "*types.SigningMaterial", "SigningParameters": "map[string]string", "Tags": "map[string]string"},
			"remove-profile-permission": {"ProfileName": "*string", "RevisionId": "*string", "StatementId": "*string"},
			"revoke-signature":          {"JobId": "*string", "JobOwner": "*string", "Reason": "*string"},
			"revoke-signing-profile":    {"EffectiveTime": "*time.Time", "ProfileName": "*string", "ProfileVersion": "*string", "Reason": "*string"},
			"sign-payload":              {"Payload": "[]byte", "PayloadFormat": "*string", "ProfileName": "*string", "ProfileOwner": "*string"},
			"start-signing-job":         {"ClientRequestToken": "*string", "Destination": "*types.Destination", "ProfileName": "*string", "ProfileOwner": "*string", "Source": "*types.Source"},
			"tag-resource":              {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":            {"ResourceArn": "*string", "TagKeys": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"add-profile-permission":    {"Action", "Principal", "ProfileName", "StatementId"},
			"cancel-signing-profile":    {"ProfileName"},
			"describe-signing-job":      {"JobId"},
			"get-revocation-status":     {"CertificateHashes", "JobArn", "PlatformId", "ProfileVersionArn", "SignatureTimestamp"},
			"get-signing-platform":      {"PlatformId"},
			"get-signing-profile":       {"ProfileName"},
			"list-profile-permissions":  {"ProfileName"},
			"list-signing-jobs":         {},
			"list-signing-platforms":    {},
			"list-signing-profiles":     {},
			"list-tags-for-resource":    {"ResourceArn"},
			"put-signing-profile":       {"PlatformId", "ProfileName"},
			"remove-profile-permission": {"ProfileName", "RevisionId", "StatementId"},
			"revoke-signature":          {"JobId", "Reason"},
			"revoke-signing-profile":    {"EffectiveTime", "ProfileName", "ProfileVersion", "Reason"},
			"sign-payload":              {"Payload", "PayloadFormat", "ProfileName"},
			"start-signing-job":         {"ClientRequestToken", "Destination", "ProfileName", "Source"},
			"tag-resource":              {"ResourceArn", "Tags"},
			"untag-resource":            {"ResourceArn", "TagKeys"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("signer", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
