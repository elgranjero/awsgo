package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/cloudsearch"
)

var fields_build_suggesters = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_create_domain = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_define_analysis_scheme = []leanruntime.Field{
	{Name: "AnalysisScheme", Flag: "analysis-scheme", Type: "*types.AnalysisScheme", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_define_expression = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "Expression", Flag: "expression", Type: "*types.Expression", Required: true},
}

var fields_define_index_field = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "IndexField", Flag: "index-field", Type: "*types.IndexField", Required: true},
}

var fields_define_suggester = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "Suggester", Flag: "suggester", Type: "*types.Suggester", Required: true},
}

var fields_delete_analysis_scheme = []leanruntime.Field{
	{Name: "AnalysisSchemeName", Flag: "analysis-scheme-name", Type: "*string", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_delete_domain = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_delete_expression = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "ExpressionName", Flag: "expression-name", Type: "*string", Required: true},
}

var fields_delete_index_field = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "IndexFieldName", Flag: "index-field-name", Type: "*string", Required: true},
}

var fields_delete_suggester = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "SuggesterName", Flag: "suggester-name", Type: "*string", Required: true},
}

var fields_describe_analysis_schemes = []leanruntime.Field{
	{Name: "AnalysisSchemeNames", Flag: "analysis-scheme-names", Type: "[]string", Required: false},
	{Name: "Deployed", Flag: "deployed", Type: "*bool", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_describe_availability_options = []leanruntime.Field{
	{Name: "Deployed", Flag: "deployed", Type: "*bool", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_describe_domain_endpoint_options = []leanruntime.Field{
	{Name: "Deployed", Flag: "deployed", Type: "*bool", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_describe_domains = []leanruntime.Field{
	{Name: "DomainNames", Flag: "domain-names", Type: "[]string", Required: false},
}

var fields_describe_expressions = []leanruntime.Field{
	{Name: "Deployed", Flag: "deployed", Type: "*bool", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "ExpressionNames", Flag: "expression-names", Type: "[]string", Required: false},
}

var fields_describe_index_fields = []leanruntime.Field{
	{Name: "Deployed", Flag: "deployed", Type: "*bool", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "FieldNames", Flag: "field-names", Type: "[]string", Required: false},
}

var fields_describe_scaling_parameters = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_describe_service_access_policies = []leanruntime.Field{
	{Name: "Deployed", Flag: "deployed", Type: "*bool", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_describe_suggesters = []leanruntime.Field{
	{Name: "Deployed", Flag: "deployed", Type: "*bool", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "SuggesterNames", Flag: "suggester-names", Type: "[]string", Required: false},
}

var fields_index_documents = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_list_domain_names = []leanruntime.Field{}

var fields_update_availability_options = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MultiAZ", Flag: "multi-az", Type: "*bool", Required: true},
}

var fields_update_domain_endpoint_options = []leanruntime.Field{
	{Name: "DomainEndpointOptions", Flag: "domain-endpoint-options", Type: "*types.DomainEndpointOptions", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_update_scaling_parameters = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "ScalingParameters", Flag: "scaling-parameters", Type: "*types.ScalingParameters", Required: true},
}

var fields_update_service_access_policies = []leanruntime.Field{
	{Name: "AccessPolicies", Flag: "access-policies", Type: "*string", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"build-suggesters": {
			Name:   "build-suggesters",
			Fields: fields_build_suggesters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BuildSuggestersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_build_suggesters, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BuildSuggesters(ctx, input)
			},
		},
		"create-domain": {
			Name:   "create-domain",
			Fields: fields_create_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDomain(ctx, input)
			},
		},
		"define-analysis-scheme": {
			Name:   "define-analysis-scheme",
			Fields: fields_define_analysis_scheme,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DefineAnalysisSchemeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_define_analysis_scheme, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DefineAnalysisScheme(ctx, input)
			},
		},
		"define-expression": {
			Name:   "define-expression",
			Fields: fields_define_expression,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DefineExpressionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_define_expression, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DefineExpression(ctx, input)
			},
		},
		"define-index-field": {
			Name:   "define-index-field",
			Fields: fields_define_index_field,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DefineIndexFieldInput{}
				if _, err := leanruntime.ApplyInput(input, fields_define_index_field, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DefineIndexField(ctx, input)
			},
		},
		"define-suggester": {
			Name:   "define-suggester",
			Fields: fields_define_suggester,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DefineSuggesterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_define_suggester, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DefineSuggester(ctx, input)
			},
		},
		"delete-analysis-scheme": {
			Name:   "delete-analysis-scheme",
			Fields: fields_delete_analysis_scheme,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAnalysisSchemeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_analysis_scheme, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAnalysisScheme(ctx, input)
			},
		},
		"delete-domain": {
			Name:   "delete-domain",
			Fields: fields_delete_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDomain(ctx, input)
			},
		},
		"delete-expression": {
			Name:   "delete-expression",
			Fields: fields_delete_expression,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteExpressionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_expression, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteExpression(ctx, input)
			},
		},
		"delete-index-field": {
			Name:   "delete-index-field",
			Fields: fields_delete_index_field,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIndexFieldInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_index_field, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIndexField(ctx, input)
			},
		},
		"delete-suggester": {
			Name:   "delete-suggester",
			Fields: fields_delete_suggester,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSuggesterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_suggester, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSuggester(ctx, input)
			},
		},
		"describe-analysis-schemes": {
			Name:   "describe-analysis-schemes",
			Fields: fields_describe_analysis_schemes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAnalysisSchemesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_analysis_schemes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAnalysisSchemes(ctx, input)
			},
		},
		"describe-availability-options": {
			Name:   "describe-availability-options",
			Fields: fields_describe_availability_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAvailabilityOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_availability_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAvailabilityOptions(ctx, input)
			},
		},
		"describe-domain-endpoint-options": {
			Name:   "describe-domain-endpoint-options",
			Fields: fields_describe_domain_endpoint_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDomainEndpointOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_domain_endpoint_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDomainEndpointOptions(ctx, input)
			},
		},
		"describe-domains": {
			Name:   "describe-domains",
			Fields: fields_describe_domains,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDomainsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_domains, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDomains(ctx, input)
			},
		},
		"describe-expressions": {
			Name:   "describe-expressions",
			Fields: fields_describe_expressions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeExpressionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_expressions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeExpressions(ctx, input)
			},
		},
		"describe-index-fields": {
			Name:   "describe-index-fields",
			Fields: fields_describe_index_fields,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIndexFieldsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_index_fields, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeIndexFields(ctx, input)
			},
		},
		"describe-scaling-parameters": {
			Name:   "describe-scaling-parameters",
			Fields: fields_describe_scaling_parameters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeScalingParametersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_scaling_parameters, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeScalingParameters(ctx, input)
			},
		},
		"describe-service-access-policies": {
			Name:   "describe-service-access-policies",
			Fields: fields_describe_service_access_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeServiceAccessPoliciesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_service_access_policies, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeServiceAccessPolicies(ctx, input)
			},
		},
		"describe-suggesters": {
			Name:   "describe-suggesters",
			Fields: fields_describe_suggesters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSuggestersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_suggesters, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSuggesters(ctx, input)
			},
		},
		"index-documents": {
			Name:   "index-documents",
			Fields: fields_index_documents,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.IndexDocumentsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_index_documents, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.IndexDocuments(ctx, input)
			},
		},
		"list-domain-names": {
			Name:   "list-domain-names",
			Fields: fields_list_domain_names,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDomainNamesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_domain_names, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDomainNames(ctx, input)
			},
		},
		"update-availability-options": {
			Name:   "update-availability-options",
			Fields: fields_update_availability_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAvailabilityOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_availability_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAvailabilityOptions(ctx, input)
			},
		},
		"update-domain-endpoint-options": {
			Name:   "update-domain-endpoint-options",
			Fields: fields_update_domain_endpoint_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDomainEndpointOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_domain_endpoint_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDomainEndpointOptions(ctx, input)
			},
		},
		"update-scaling-parameters": {
			Name:   "update-scaling-parameters",
			Fields: fields_update_scaling_parameters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateScalingParametersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_scaling_parameters, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateScalingParameters(ctx, input)
			},
		},
		"update-service-access-policies": {
			Name:   "update-service-access-policies",
			Fields: fields_update_service_access_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateServiceAccessPoliciesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_service_access_policies, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateServiceAccessPolicies(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("cloudsearch", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
