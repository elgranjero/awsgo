package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/geoplaces"
)

var fields_autocomplete = []leanruntime.Field{
	{Name: "AdditionalFeatures", Flag: "additional-features", Type: "[]types.AutocompleteAdditionalFeature", Required: false},
	{Name: "BiasPosition", Flag: "bias-position", Type: "[]float64", Required: false},
	{Name: "Filter", Flag: "filter", Type: "*types.AutocompleteFilter", Required: false},
	{Name: "IntendedUse", Flag: "intended-use", Type: "types.AutocompleteIntendedUse", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: false},
	{Name: "Language", Flag: "language", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "PoliticalView", Flag: "political-view", Type: "*string", Required: false},
	{Name: "PostalCodeMode", Flag: "postal-code-mode", Type: "types.PostalCodeMode", Required: false},
	{Name: "QueryText", Flag: "query-text", Type: "*string", Required: true},
}

var fields_geocode = []leanruntime.Field{
	{Name: "AdditionalFeatures", Flag: "additional-features", Type: "[]types.GeocodeAdditionalFeature", Required: false},
	{Name: "BiasPosition", Flag: "bias-position", Type: "[]float64", Required: false},
	{Name: "Filter", Flag: "filter", Type: "*types.GeocodeFilter", Required: false},
	{Name: "IntendedUse", Flag: "intended-use", Type: "types.GeocodeIntendedUse", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: false},
	{Name: "Language", Flag: "language", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "PoliticalView", Flag: "political-view", Type: "*string", Required: false},
	{Name: "QueryComponents", Flag: "query-components", Type: "*types.GeocodeQueryComponents", Required: false},
	{Name: "QueryText", Flag: "query-text", Type: "*string", Required: false},
}

var fields_get_place = []leanruntime.Field{
	{Name: "AdditionalFeatures", Flag: "additional-features", Type: "[]types.GetPlaceAdditionalFeature", Required: false},
	{Name: "IntendedUse", Flag: "intended-use", Type: "types.GetPlaceIntendedUse", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: false},
	{Name: "Language", Flag: "language", Type: "*string", Required: false},
	{Name: "PlaceId", Flag: "place-id", Type: "*string", Required: true},
	{Name: "PoliticalView", Flag: "political-view", Type: "*string", Required: false},
}

var fields_reverse_geocode = []leanruntime.Field{
	{Name: "AdditionalFeatures", Flag: "additional-features", Type: "[]types.ReverseGeocodeAdditionalFeature", Required: false},
	{Name: "Filter", Flag: "filter", Type: "*types.ReverseGeocodeFilter", Required: false},
	{Name: "Heading", Flag: "heading", Type: "float64", Required: false},
	{Name: "IntendedUse", Flag: "intended-use", Type: "types.ReverseGeocodeIntendedUse", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: false},
	{Name: "Language", Flag: "language", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "PoliticalView", Flag: "political-view", Type: "*string", Required: false},
	{Name: "QueryPosition", Flag: "query-position", Type: "[]float64", Required: true},
	{Name: "QueryRadius", Flag: "query-radius", Type: "*int64", Required: false},
}

var fields_search_nearby = []leanruntime.Field{
	{Name: "AdditionalFeatures", Flag: "additional-features", Type: "[]types.SearchNearbyAdditionalFeature", Required: false},
	{Name: "Filter", Flag: "filter", Type: "*types.SearchNearbyFilter", Required: false},
	{Name: "IntendedUse", Flag: "intended-use", Type: "types.SearchNearbyIntendedUse", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: false},
	{Name: "Language", Flag: "language", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PoliticalView", Flag: "political-view", Type: "*string", Required: false},
	{Name: "QueryPosition", Flag: "query-position", Type: "[]float64", Required: true},
	{Name: "QueryRadius", Flag: "query-radius", Type: "*int64", Required: false},
}

var fields_search_text = []leanruntime.Field{
	{Name: "AdditionalFeatures", Flag: "additional-features", Type: "[]types.SearchTextAdditionalFeature", Required: false},
	{Name: "BiasPosition", Flag: "bias-position", Type: "[]float64", Required: false},
	{Name: "Filter", Flag: "filter", Type: "*types.SearchTextFilter", Required: false},
	{Name: "IntendedUse", Flag: "intended-use", Type: "types.SearchTextIntendedUse", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: false},
	{Name: "Language", Flag: "language", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PoliticalView", Flag: "political-view", Type: "*string", Required: false},
	{Name: "QueryId", Flag: "query-id", Type: "*string", Required: false},
	{Name: "QueryText", Flag: "query-text", Type: "*string", Required: false},
}

var fields_suggest = []leanruntime.Field{
	{Name: "AdditionalFeatures", Flag: "additional-features", Type: "[]types.SuggestAdditionalFeature", Required: false},
	{Name: "BiasPosition", Flag: "bias-position", Type: "[]float64", Required: false},
	{Name: "Filter", Flag: "filter", Type: "*types.SuggestFilter", Required: false},
	{Name: "IntendedUse", Flag: "intended-use", Type: "types.SuggestIntendedUse", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: false},
	{Name: "Language", Flag: "language", Type: "*string", Required: false},
	{Name: "MaxQueryRefinements", Flag: "max-query-refinements", Type: "*int32", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "PoliticalView", Flag: "political-view", Type: "*string", Required: false},
	{Name: "QueryText", Flag: "query-text", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"autocomplete": {
			Name:   "autocomplete",
			Fields: fields_autocomplete,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AutocompleteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_autocomplete, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Autocomplete(ctx, input)
			},
		},
		"geocode": {
			Name:   "geocode",
			Fields: fields_geocode,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GeocodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_geocode, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Geocode(ctx, input)
			},
		},
		"get-place": {
			Name:   "get-place",
			Fields: fields_get_place,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPlaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_place, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPlace(ctx, input)
			},
		},
		"reverse-geocode": {
			Name:   "reverse-geocode",
			Fields: fields_reverse_geocode,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReverseGeocodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reverse_geocode, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReverseGeocode(ctx, input)
			},
		},
		"search-nearby": {
			Name:   "search-nearby",
			Fields: fields_search_nearby,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchNearbyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_search_nearby, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SearchNearby(ctx, input)
			},
		},
		"search-text": {
			Name:   "search-text",
			Fields: fields_search_text,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchTextInput{}
				if _, err := leanruntime.ApplyInput(input, fields_search_text, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SearchText(ctx, input)
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
	}
	if err := leanruntime.Execute("geoplaces", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
