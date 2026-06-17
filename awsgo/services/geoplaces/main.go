package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/geoplaces/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"autocomplete", "geocode", "get-place", "reverse-geocode", "search-nearby", "search-text", "suggest"},
		OperationSet: map[string]bool{"autocomplete": true, "geocode": true, "get-place": true, "reverse-geocode": true, "search-nearby": true, "search-text": true, "suggest": true},
		OperationInputs: map[string][]string{
			"autocomplete":    {"AdditionalFeatures", "BiasPosition", "Filter", "IntendedUse", "Key", "Language", "MaxResults", "PoliticalView", "PostalCodeMode", "QueryText"},
			"geocode":         {"AdditionalFeatures", "BiasPosition", "Filter", "IntendedUse", "Key", "Language", "MaxResults", "PoliticalView", "QueryComponents", "QueryText"},
			"get-place":       {"AdditionalFeatures", "IntendedUse", "Key", "Language", "PlaceId", "PoliticalView"},
			"reverse-geocode": {"AdditionalFeatures", "Filter", "Heading", "IntendedUse", "Key", "Language", "MaxResults", "PoliticalView", "QueryPosition", "QueryRadius"},
			"search-nearby":   {"AdditionalFeatures", "Filter", "IntendedUse", "Key", "Language", "MaxResults", "NextToken", "PoliticalView", "QueryPosition", "QueryRadius"},
			"search-text":     {"AdditionalFeatures", "BiasPosition", "Filter", "IntendedUse", "Key", "Language", "MaxResults", "NextToken", "PoliticalView", "QueryId", "QueryText"},
			"suggest":         {"AdditionalFeatures", "BiasPosition", "Filter", "IntendedUse", "Key", "Language", "MaxQueryRefinements", "MaxResults", "PoliticalView", "QueryText"},
		},
		OperationInputTypes: map[string]map[string]string{
			"autocomplete":    {"AdditionalFeatures": "[]types.AutocompleteAdditionalFeature", "BiasPosition": "[]float64", "Filter": "*types.AutocompleteFilter", "IntendedUse": "types.AutocompleteIntendedUse", "Key": "*string", "Language": "*string", "MaxResults": "*int32", "PoliticalView": "*string", "PostalCodeMode": "types.PostalCodeMode", "QueryText": "*string"},
			"geocode":         {"AdditionalFeatures": "[]types.GeocodeAdditionalFeature", "BiasPosition": "[]float64", "Filter": "*types.GeocodeFilter", "IntendedUse": "types.GeocodeIntendedUse", "Key": "*string", "Language": "*string", "MaxResults": "*int32", "PoliticalView": "*string", "QueryComponents": "*types.GeocodeQueryComponents", "QueryText": "*string"},
			"get-place":       {"AdditionalFeatures": "[]types.GetPlaceAdditionalFeature", "IntendedUse": "types.GetPlaceIntendedUse", "Key": "*string", "Language": "*string", "PlaceId": "*string", "PoliticalView": "*string"},
			"reverse-geocode": {"AdditionalFeatures": "[]types.ReverseGeocodeAdditionalFeature", "Filter": "*types.ReverseGeocodeFilter", "Heading": "float64", "IntendedUse": "types.ReverseGeocodeIntendedUse", "Key": "*string", "Language": "*string", "MaxResults": "*int32", "PoliticalView": "*string", "QueryPosition": "[]float64", "QueryRadius": "*int64"},
			"search-nearby":   {"AdditionalFeatures": "[]types.SearchNearbyAdditionalFeature", "Filter": "*types.SearchNearbyFilter", "IntendedUse": "types.SearchNearbyIntendedUse", "Key": "*string", "Language": "*string", "MaxResults": "*int32", "NextToken": "*string", "PoliticalView": "*string", "QueryPosition": "[]float64", "QueryRadius": "*int64"},
			"search-text":     {"AdditionalFeatures": "[]types.SearchTextAdditionalFeature", "BiasPosition": "[]float64", "Filter": "*types.SearchTextFilter", "IntendedUse": "types.SearchTextIntendedUse", "Key": "*string", "Language": "*string", "MaxResults": "*int32", "NextToken": "*string", "PoliticalView": "*string", "QueryId": "*string", "QueryText": "*string"},
			"suggest":         {"AdditionalFeatures": "[]types.SuggestAdditionalFeature", "BiasPosition": "[]float64", "Filter": "*types.SuggestFilter", "IntendedUse": "types.SuggestIntendedUse", "Key": "*string", "Language": "*string", "MaxQueryRefinements": "*int32", "MaxResults": "*int32", "PoliticalView": "*string", "QueryText": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"autocomplete":    {"QueryText"},
			"geocode":         {},
			"get-place":       {"PlaceId"},
			"reverse-geocode": {"QueryPosition"},
			"search-nearby":   {"QueryPosition"},
			"search-text":     {},
			"suggest":         {"QueryText"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("geoplaces", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
