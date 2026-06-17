package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/paymentcryptography/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"add-key-replication-regions", "create-alias", "create-key", "delete-alias", "delete-key", "disable-default-key-replication-regions", "enable-default-key-replication-regions", "export-key", "get-alias", "get-certificate-signing-request", "get-default-key-replication-regions", "get-key", "get-parameters-for-export", "get-parameters-for-import", "get-public-key-certificate", "import-key", "list-aliases", "list-keys", "list-tags-for-resource", "remove-key-replication-regions", "restore-key", "start-key-usage", "stop-key-usage", "tag-resource", "untag-resource", "update-alias"},
		OperationSet: map[string]bool{"add-key-replication-regions": true, "create-alias": true, "create-key": true, "delete-alias": true, "delete-key": true, "disable-default-key-replication-regions": true, "enable-default-key-replication-regions": true, "export-key": true, "get-alias": true, "get-certificate-signing-request": true, "get-default-key-replication-regions": true, "get-key": true, "get-parameters-for-export": true, "get-parameters-for-import": true, "get-public-key-certificate": true, "import-key": true, "list-aliases": true, "list-keys": true, "list-tags-for-resource": true, "remove-key-replication-regions": true, "restore-key": true, "start-key-usage": true, "stop-key-usage": true, "tag-resource": true, "untag-resource": true, "update-alias": true},
		OperationInputs: map[string][]string{
			"add-key-replication-regions": {"KeyIdentifier", "ReplicationRegions"},
			"create-alias":                {"AliasName", "KeyArn"},
			"create-key":                  {"DeriveKeyUsage", "Enabled", "Exportable", "KeyAttributes", "KeyCheckValueAlgorithm", "ReplicationRegions", "Tags"},
			"delete-alias":                {"AliasName"},
			"delete-key":                  {"DeleteKeyInDays", "KeyIdentifier"},
			"disable-default-key-replication-regions": {"ReplicationRegions"},
			"enable-default-key-replication-regions":  {"ReplicationRegions"},
			"export-key":                              {"ExportAttributes", "ExportKeyIdentifier", "KeyMaterial"},
			"get-alias":                               {"AliasName"},
			"get-certificate-signing-request":         {"CertificateSubject", "KeyIdentifier", "SigningAlgorithm"},
			"get-default-key-replication-regions":     {},
			"get-key":                                 {"KeyIdentifier"},
			"get-parameters-for-export":               {"KeyMaterialType", "SigningKeyAlgorithm"},
			"get-parameters-for-import":               {"KeyMaterialType", "WrappingKeyAlgorithm"},
			"get-public-key-certificate":              {"KeyIdentifier"},
			"import-key":                              {"Enabled", "KeyCheckValueAlgorithm", "KeyMaterial", "ReplicationRegions", "Tags"},
			"list-aliases":                            {"KeyArn", "MaxResults", "NextToken"},
			"list-keys":                               {"KeyState", "MaxResults", "NextToken"},
			"list-tags-for-resource":                  {"MaxResults", "NextToken", "ResourceArn"},
			"remove-key-replication-regions":          {"KeyIdentifier", "ReplicationRegions"},
			"restore-key":                             {"KeyIdentifier"},
			"start-key-usage":                         {"KeyIdentifier"},
			"stop-key-usage":                          {"KeyIdentifier"},
			"tag-resource":                            {"ResourceArn", "Tags"},
			"untag-resource":                          {"ResourceArn", "TagKeys"},
			"update-alias":                            {"AliasName", "KeyArn"},
		},
		OperationInputTypes: map[string]map[string]string{
			"add-key-replication-regions": {"KeyIdentifier": "*string", "ReplicationRegions": "[]string"},
			"create-alias":                {"AliasName": "*string", "KeyArn": "*string"},
			"create-key":                  {"DeriveKeyUsage": "types.DeriveKeyUsage", "Enabled": "*bool", "Exportable": "*bool", "KeyAttributes": "*types.KeyAttributes", "KeyCheckValueAlgorithm": "types.KeyCheckValueAlgorithm", "ReplicationRegions": "[]string", "Tags": "[]types.Tag"},
			"delete-alias":                {"AliasName": "*string"},
			"delete-key":                  {"DeleteKeyInDays": "*int32", "KeyIdentifier": "*string"},
			"disable-default-key-replication-regions": {"ReplicationRegions": "[]string"},
			"enable-default-key-replication-regions":  {"ReplicationRegions": "[]string"},
			"export-key":                              {"ExportAttributes": "*types.ExportAttributes", "ExportKeyIdentifier": "*string", "KeyMaterial": "types.ExportKeyMaterial"},
			"get-alias":                               {"AliasName": "*string"},
			"get-certificate-signing-request":         {"CertificateSubject": "*types.CertificateSubjectType", "KeyIdentifier": "*string", "SigningAlgorithm": "types.SigningAlgorithmType"},
			"get-default-key-replication-regions":     {},
			"get-key":                                 {"KeyIdentifier": "*string"},
			"get-parameters-for-export":               {"KeyMaterialType": "types.KeyMaterialType", "SigningKeyAlgorithm": "types.KeyAlgorithm"},
			"get-parameters-for-import":               {"KeyMaterialType": "types.KeyMaterialType", "WrappingKeyAlgorithm": "types.KeyAlgorithm"},
			"get-public-key-certificate":              {"KeyIdentifier": "*string"},
			"import-key":                              {"Enabled": "*bool", "KeyCheckValueAlgorithm": "types.KeyCheckValueAlgorithm", "KeyMaterial": "types.ImportKeyMaterial", "ReplicationRegions": "[]string", "Tags": "[]types.Tag"},
			"list-aliases":                            {"KeyArn": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-keys":                               {"KeyState": "types.KeyState", "MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":                  {"MaxResults": "*int32", "NextToken": "*string", "ResourceArn": "*string"},
			"remove-key-replication-regions":          {"KeyIdentifier": "*string", "ReplicationRegions": "[]string"},
			"restore-key":                             {"KeyIdentifier": "*string"},
			"start-key-usage":                         {"KeyIdentifier": "*string"},
			"stop-key-usage":                          {"KeyIdentifier": "*string"},
			"tag-resource":                            {"ResourceArn": "*string", "Tags": "[]types.Tag"},
			"untag-resource":                          {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-alias":                            {"AliasName": "*string", "KeyArn": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"add-key-replication-regions": {"KeyIdentifier", "ReplicationRegions"},
			"create-alias":                {"AliasName"},
			"create-key":                  {"Exportable", "KeyAttributes"},
			"delete-alias":                {"AliasName"},
			"delete-key":                  {"KeyIdentifier"},
			"disable-default-key-replication-regions": {"ReplicationRegions"},
			"enable-default-key-replication-regions":  {"ReplicationRegions"},
			"export-key":                              {"ExportKeyIdentifier", "KeyMaterial"},
			"get-alias":                               {"AliasName"},
			"get-certificate-signing-request":         {"CertificateSubject", "KeyIdentifier", "SigningAlgorithm"},
			"get-default-key-replication-regions":     {},
			"get-key":                                 {"KeyIdentifier"},
			"get-parameters-for-export":               {"KeyMaterialType", "SigningKeyAlgorithm"},
			"get-parameters-for-import":               {"KeyMaterialType", "WrappingKeyAlgorithm"},
			"get-public-key-certificate":              {"KeyIdentifier"},
			"import-key":                              {"KeyMaterial"},
			"list-aliases":                            {},
			"list-keys":                               {},
			"list-tags-for-resource":                  {"ResourceArn"},
			"remove-key-replication-regions":          {"KeyIdentifier", "ReplicationRegions"},
			"restore-key":                             {"KeyIdentifier"},
			"start-key-usage":                         {"KeyIdentifier"},
			"stop-key-usage":                          {"KeyIdentifier"},
			"tag-resource":                            {"ResourceArn", "Tags"},
			"untag-resource":                          {"ResourceArn", "TagKeys"},
			"update-alias":                            {"AliasName"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("paymentcryptography", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
