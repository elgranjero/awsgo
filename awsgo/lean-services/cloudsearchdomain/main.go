package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/cloudsearchdomain"
)

var fields_search = []leanruntime.Field{
	{Name: "Cursor", Flag: "cursor", Type: "*string", Required: false},
	{Name: "Expr", Flag: "expr", Type: "*string", Required: false},
	{Name: "Facet", Flag: "facet", Type: "*string", Required: false},
	{Name: "FilterQuery", Flag: "filter-query", Type: "*string", Required: false},
	{Name: "Highlight", Flag: "highlight", Type: "*string", Required: false},
	{Name: "Partial", Flag: "partial", Type: "bool", Required: false},
	{Name: "Query", Flag: "query", Type: "*string", Required: true},
	{Name: "QueryOptions", Flag: "query-options", Type: "*string", Required: false},
	{Name: "QueryParser", Flag: "query-parser", Type: "types.QueryParser", Required: false},
	{Name: "Return", Flag: "return", Type: "*string", Required: false},
	{Name: "Size", Flag: "size", Type: "int64", Required: false},
	{Name: "Sort", Flag: "sort", Type: "*string", Required: false},
	{Name: "Start", Flag: "start", Type: "int64", Required: false},
	{Name: "Stats", Flag: "stats", Type: "*string", Required: false},
}

var fields_suggest = []leanruntime.Field{
	{Name: "Query", Flag: "query", Type: "*string", Required: true},
	{Name: "Size", Flag: "size", Type: "int64", Required: false},
	{Name: "Suggester", Flag: "suggester", Type: "*string", Required: true},
}

var fields_upload_documents = []leanruntime.Field{
	{Name: "ContentType", Flag: "content-type", Type: "types.ContentType", Required: true},
	{Name: "Documents", Flag: "documents", Type: "io.Reader", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"search": {
			Name:   "search",
			Fields: fields_search,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_search, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Search(ctx, input)
			},
		},
		"suggest": {
			Name:   "suggest",
			Fields: fields_suggest,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SuggestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_suggest, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Suggest(ctx, input)
			},
		},
		"upload-documents": {
			Name:   "upload-documents",
			Fields: fields_upload_documents,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UploadDocumentsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_upload_documents, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UploadDocuments(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("cloudsearchdomain", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
