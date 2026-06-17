package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/cloudhsm/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"add-tags-to-resource", "create-hapg", "create-hsm", "create-luna-client", "delete-hapg", "delete-hsm", "delete-luna-client", "describe-hapg", "describe-hsm", "describe-luna-client", "get-config", "list-available-zones", "list-hapgs", "list-hsms", "list-luna-clients", "list-tags-for-resource", "modify-hapg", "modify-hsm", "modify-luna-client", "remove-tags-from-resource"},
		OperationSet: map[string]bool{"add-tags-to-resource": true, "create-hapg": true, "create-hsm": true, "create-luna-client": true, "delete-hapg": true, "delete-hsm": true, "delete-luna-client": true, "describe-hapg": true, "describe-hsm": true, "describe-luna-client": true, "get-config": true, "list-available-zones": true, "list-hapgs": true, "list-hsms": true, "list-luna-clients": true, "list-tags-for-resource": true, "modify-hapg": true, "modify-hsm": true, "modify-luna-client": true, "remove-tags-from-resource": true},
		OperationInputs: map[string][]string{
			"add-tags-to-resource":      {"ResourceArn", "TagList"},
			"create-hapg":               {"Label"},
			"create-hsm":                {"ClientToken", "EniIp", "ExternalId", "IamRoleArn", "SshKey", "SubnetId", "SubscriptionType", "SyslogIp"},
			"create-luna-client":        {"Certificate", "Label"},
			"delete-hapg":               {"HapgArn"},
			"delete-hsm":                {"HsmArn"},
			"delete-luna-client":        {"ClientArn"},
			"describe-hapg":             {"HapgArn"},
			"describe-hsm":              {"HsmArn", "HsmSerialNumber"},
			"describe-luna-client":      {"CertificateFingerprint", "ClientArn"},
			"get-config":                {"ClientArn", "ClientVersion", "HapgList"},
			"list-available-zones":      {},
			"list-hapgs":                {"NextToken"},
			"list-hsms":                 {"NextToken"},
			"list-luna-clients":         {"NextToken"},
			"list-tags-for-resource":    {"ResourceArn"},
			"modify-hapg":               {"HapgArn", "Label", "PartitionSerialList"},
			"modify-hsm":                {"EniIp", "ExternalId", "HsmArn", "IamRoleArn", "SubnetId", "SyslogIp"},
			"modify-luna-client":        {"Certificate", "ClientArn"},
			"remove-tags-from-resource": {"ResourceArn", "TagKeyList"},
		},
		OperationInputTypes: map[string]map[string]string{
			"add-tags-to-resource":      {"ResourceArn": "*string", "TagList": "[]types.Tag"},
			"create-hapg":               {"Label": "*string"},
			"create-hsm":                {"ClientToken": "*string", "EniIp": "*string", "ExternalId": "*string", "IamRoleArn": "*string", "SshKey": "*string", "SubnetId": "*string", "SubscriptionType": "types.SubscriptionType", "SyslogIp": "*string"},
			"create-luna-client":        {"Certificate": "*string", "Label": "*string"},
			"delete-hapg":               {"HapgArn": "*string"},
			"delete-hsm":                {"HsmArn": "*string"},
			"delete-luna-client":        {"ClientArn": "*string"},
			"describe-hapg":             {"HapgArn": "*string"},
			"describe-hsm":              {"HsmArn": "*string", "HsmSerialNumber": "*string"},
			"describe-luna-client":      {"CertificateFingerprint": "*string", "ClientArn": "*string"},
			"get-config":                {"ClientArn": "*string", "ClientVersion": "types.ClientVersion", "HapgList": "[]string"},
			"list-available-zones":      {},
			"list-hapgs":                {"NextToken": "*string"},
			"list-hsms":                 {"NextToken": "*string"},
			"list-luna-clients":         {"NextToken": "*string"},
			"list-tags-for-resource":    {"ResourceArn": "*string"},
			"modify-hapg":               {"HapgArn": "*string", "Label": "*string", "PartitionSerialList": "[]string"},
			"modify-hsm":                {"EniIp": "*string", "ExternalId": "*string", "HsmArn": "*string", "IamRoleArn": "*string", "SubnetId": "*string", "SyslogIp": "*string"},
			"modify-luna-client":        {"Certificate": "*string", "ClientArn": "*string"},
			"remove-tags-from-resource": {"ResourceArn": "*string", "TagKeyList": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"add-tags-to-resource":      {"ResourceArn", "TagList"},
			"create-hapg":               {"Label"},
			"create-hsm":                {"IamRoleArn", "SshKey", "SubnetId", "SubscriptionType"},
			"create-luna-client":        {"Certificate"},
			"delete-hapg":               {"HapgArn"},
			"delete-hsm":                {"HsmArn"},
			"delete-luna-client":        {"ClientArn"},
			"describe-hapg":             {"HapgArn"},
			"describe-hsm":              {},
			"describe-luna-client":      {},
			"get-config":                {"ClientArn", "ClientVersion", "HapgList"},
			"list-available-zones":      {},
			"list-hapgs":                {},
			"list-hsms":                 {},
			"list-luna-clients":         {},
			"list-tags-for-resource":    {"ResourceArn"},
			"modify-hapg":               {"HapgArn"},
			"modify-hsm":                {"HsmArn"},
			"modify-luna-client":        {"Certificate", "ClientArn"},
			"remove-tags-from-resource": {"ResourceArn", "TagKeyList"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("cloudhsm", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
