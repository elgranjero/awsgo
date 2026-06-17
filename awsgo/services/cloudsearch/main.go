package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/cloudsearch/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"build-suggesters", "create-domain", "define-analysis-scheme", "define-expression", "define-index-field", "define-suggester", "delete-analysis-scheme", "delete-domain", "delete-expression", "delete-index-field", "delete-suggester", "describe-analysis-schemes", "describe-availability-options", "describe-domain-endpoint-options", "describe-domains", "describe-expressions", "describe-index-fields", "describe-scaling-parameters", "describe-service-access-policies", "describe-suggesters", "index-documents", "list-domain-names", "update-availability-options", "update-domain-endpoint-options", "update-scaling-parameters", "update-service-access-policies"},
		OperationSet: map[string]bool{"build-suggesters": true, "create-domain": true, "define-analysis-scheme": true, "define-expression": true, "define-index-field": true, "define-suggester": true, "delete-analysis-scheme": true, "delete-domain": true, "delete-expression": true, "delete-index-field": true, "delete-suggester": true, "describe-analysis-schemes": true, "describe-availability-options": true, "describe-domain-endpoint-options": true, "describe-domains": true, "describe-expressions": true, "describe-index-fields": true, "describe-scaling-parameters": true, "describe-service-access-policies": true, "describe-suggesters": true, "index-documents": true, "list-domain-names": true, "update-availability-options": true, "update-domain-endpoint-options": true, "update-scaling-parameters": true, "update-service-access-policies": true},
		OperationInputs: map[string][]string{
			"build-suggesters":                 {"DomainName"},
			"create-domain":                    {"DomainName"},
			"define-analysis-scheme":           {"AnalysisScheme", "DomainName"},
			"define-expression":                {"DomainName", "Expression"},
			"define-index-field":               {"DomainName", "IndexField"},
			"define-suggester":                 {"DomainName", "Suggester"},
			"delete-analysis-scheme":           {"AnalysisSchemeName", "DomainName"},
			"delete-domain":                    {"DomainName"},
			"delete-expression":                {"DomainName", "ExpressionName"},
			"delete-index-field":               {"DomainName", "IndexFieldName"},
			"delete-suggester":                 {"DomainName", "SuggesterName"},
			"describe-analysis-schemes":        {"AnalysisSchemeNames", "Deployed", "DomainName"},
			"describe-availability-options":    {"Deployed", "DomainName"},
			"describe-domain-endpoint-options": {"Deployed", "DomainName"},
			"describe-domains":                 {"DomainNames"},
			"describe-expressions":             {"Deployed", "DomainName", "ExpressionNames"},
			"describe-index-fields":            {"Deployed", "DomainName", "FieldNames"},
			"describe-scaling-parameters":      {"DomainName"},
			"describe-service-access-policies": {"Deployed", "DomainName"},
			"describe-suggesters":              {"Deployed", "DomainName", "SuggesterNames"},
			"index-documents":                  {"DomainName"},
			"list-domain-names":                {},
			"update-availability-options":      {"DomainName", "MultiAZ"},
			"update-domain-endpoint-options":   {"DomainEndpointOptions", "DomainName"},
			"update-scaling-parameters":        {"DomainName", "ScalingParameters"},
			"update-service-access-policies":   {"AccessPolicies", "DomainName"},
		},
		OperationInputTypes: map[string]map[string]string{
			"build-suggesters":                 {"DomainName": "*string"},
			"create-domain":                    {"DomainName": "*string"},
			"define-analysis-scheme":           {"AnalysisScheme": "*types.AnalysisScheme", "DomainName": "*string"},
			"define-expression":                {"DomainName": "*string", "Expression": "*types.Expression"},
			"define-index-field":               {"DomainName": "*string", "IndexField": "*types.IndexField"},
			"define-suggester":                 {"DomainName": "*string", "Suggester": "*types.Suggester"},
			"delete-analysis-scheme":           {"AnalysisSchemeName": "*string", "DomainName": "*string"},
			"delete-domain":                    {"DomainName": "*string"},
			"delete-expression":                {"DomainName": "*string", "ExpressionName": "*string"},
			"delete-index-field":               {"DomainName": "*string", "IndexFieldName": "*string"},
			"delete-suggester":                 {"DomainName": "*string", "SuggesterName": "*string"},
			"describe-analysis-schemes":        {"AnalysisSchemeNames": "[]string", "Deployed": "*bool", "DomainName": "*string"},
			"describe-availability-options":    {"Deployed": "*bool", "DomainName": "*string"},
			"describe-domain-endpoint-options": {"Deployed": "*bool", "DomainName": "*string"},
			"describe-domains":                 {"DomainNames": "[]string"},
			"describe-expressions":             {"Deployed": "*bool", "DomainName": "*string", "ExpressionNames": "[]string"},
			"describe-index-fields":            {"Deployed": "*bool", "DomainName": "*string", "FieldNames": "[]string"},
			"describe-scaling-parameters":      {"DomainName": "*string"},
			"describe-service-access-policies": {"Deployed": "*bool", "DomainName": "*string"},
			"describe-suggesters":              {"Deployed": "*bool", "DomainName": "*string", "SuggesterNames": "[]string"},
			"index-documents":                  {"DomainName": "*string"},
			"list-domain-names":                {},
			"update-availability-options":      {"DomainName": "*string", "MultiAZ": "*bool"},
			"update-domain-endpoint-options":   {"DomainEndpointOptions": "*types.DomainEndpointOptions", "DomainName": "*string"},
			"update-scaling-parameters":        {"DomainName": "*string", "ScalingParameters": "*types.ScalingParameters"},
			"update-service-access-policies":   {"AccessPolicies": "*string", "DomainName": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"build-suggesters":                 {"DomainName"},
			"create-domain":                    {"DomainName"},
			"define-analysis-scheme":           {"AnalysisScheme", "DomainName"},
			"define-expression":                {"DomainName", "Expression"},
			"define-index-field":               {"DomainName", "IndexField"},
			"define-suggester":                 {"DomainName", "Suggester"},
			"delete-analysis-scheme":           {"AnalysisSchemeName", "DomainName"},
			"delete-domain":                    {"DomainName"},
			"delete-expression":                {"DomainName", "ExpressionName"},
			"delete-index-field":               {"DomainName", "IndexFieldName"},
			"delete-suggester":                 {"DomainName", "SuggesterName"},
			"describe-analysis-schemes":        {"DomainName"},
			"describe-availability-options":    {"DomainName"},
			"describe-domain-endpoint-options": {"DomainName"},
			"describe-domains":                 {},
			"describe-expressions":             {"DomainName"},
			"describe-index-fields":            {"DomainName"},
			"describe-scaling-parameters":      {"DomainName"},
			"describe-service-access-policies": {"DomainName"},
			"describe-suggesters":              {"DomainName"},
			"index-documents":                  {"DomainName"},
			"list-domain-names":                {},
			"update-availability-options":      {"DomainName", "MultiAZ"},
			"update-domain-endpoint-options":   {"DomainEndpointOptions", "DomainName"},
			"update-scaling-parameters":        {"DomainName", "ScalingParameters"},
			"update-service-access-policies":   {"AccessPolicies", "DomainName"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("cloudsearch", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
