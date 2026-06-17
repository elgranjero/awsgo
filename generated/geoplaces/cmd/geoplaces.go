package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/geoplaces"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// geoplacesCmd represents the geoplaces command
var _geoplacesCmd = &cobra.Command{
	Use:   "geoplaces",
	Short: "AWS geoplaces CLI",
	Run: func(cmd *cobra.Command, args []string) {
		_awsOutput = resolveAWSOutput(_awsProfile, cmd.Flags().Changed("output"))
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := geoplaces.NewFromConfig(cfg)
		if _geoplacesAutocomplete {
			geoplaces_Autocomplete(cfg, client)
			return
		}
		if _geoplacesGeocode {
			geoplaces_Geocode(cfg, client)
			return
		}
		if _geoplacesGetPlace {
			geoplaces_GetPlace(cfg, client)
			return
		}
		if _geoplacesReverseGeocode {
			geoplaces_ReverseGeocode(cfg, client)
			return
		}
		if _geoplacesSearchNearby {
			geoplaces_SearchNearby(cfg, client)
			return
		}
		if _geoplacesSearchText {
			geoplaces_SearchText(cfg, client)
			return
		}
		if _geoplacesSuggest {
			geoplaces_Suggest(cfg, client)
			return
		}

	},
}

var (
	_geoplacesAutocomplete   bool
	_geoplacesGeocode        bool
	_geoplacesGetPlace       bool
	_geoplacesReverseGeocode bool
	_geoplacesSearchNearby   bool
	_geoplacesSearchText     bool
	_geoplacesSuggest        bool

	_geoplacesAdditionalFeatures  string
	_geoplacesBiasPosition        string
	_geoplacesFilter              string
	_geoplacesHeading             string
	_geoplacesIntendedUse         string
	_geoplacesKey                 string
	_geoplacesLanguage            string
	_geoplacesMaxQueryRefinements string
	_geoplacesMaxResults          string
	_geoplacesNextToken           string
	_geoplacesPlaceId             string
	_geoplacesPoliticalView       string
	_geoplacesPostalCodeMode      string
	_geoplacesQueryComponents     string
	_geoplacesQueryId             string
	_geoplacesQueryPosition       string
	_geoplacesQueryRadius         string
	_geoplacesQueryText           string
)

// Autocomplete completes potential places and addresses as the user types, based
// on the partial input. The API enhances the efficiency and accuracy of address by
// completing query based on a few entered keystrokes. It helps you by completing
// partial queries with valid address completion. Also, the API supports the
// filtering of results based on geographic location, country, or specific place
// types, and can be tailored using optional parameters like language and political
// views.
//
// For more information, see [Autocomplete] in the Amazon Location Service Developer Guide.
//
// [Autocomplete]: https://docs.aws.amazon.com/location/latest/developerguide/autocomplete.html
func geoplaces_Autocomplete(cfg aws.Config, client *geoplaces.Client) {
	input := &geoplaces.AutocompleteInput{
		// QueryText: *string, // Required
	}

	if len(_geoplacesQueryText) > 0 {
		input.QueryText = aws.String(_geoplacesQueryText)
	}
	if len(_geoplacesAdditionalFeatures) > 0 {
		if err := assignInputField(input, "AdditionalFeatures", _geoplacesAdditionalFeatures); err != nil {
			log.Errorf("invalid --additional-features: %s", err.Error())
			return
		}
	}
	if len(_geoplacesBiasPosition) > 0 {
		if err := assignInputField(input, "BiasPosition", _geoplacesBiasPosition); err != nil {
			log.Errorf("invalid --bias-position: %s", err.Error())
			return
		}
	}
	if len(_geoplacesFilter) > 0 {
		if err := assignInputField(input, "Filter", _geoplacesFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_geoplacesIntendedUse) > 0 {
		if err := assignInputField(input, "IntendedUse", _geoplacesIntendedUse); err != nil {
			log.Errorf("invalid --intended-use: %s", err.Error())
			return
		}
	}
	if len(_geoplacesKey) > 0 {
		input.Key = aws.String(_geoplacesKey)
	}
	if len(_geoplacesLanguage) > 0 {
		input.Language = aws.String(_geoplacesLanguage)
	}
	if len(_geoplacesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _geoplacesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_geoplacesPoliticalView) > 0 {
		input.PoliticalView = aws.String(_geoplacesPoliticalView)
	}
	if len(_geoplacesPostalCodeMode) > 0 {
		if err := assignInputField(input, "PostalCodeMode", _geoplacesPostalCodeMode); err != nil {
			log.Errorf("invalid --postal-code-mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.Autocomplete(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Geocode converts a textual address or place into geographic coordinates. You
// can obtain geographic coordinates, address component, and other related
// information. It supports flexible queries, including free-form text or
// structured queries with components like street names, postal codes, and regions.
// The Geocode API can also provide additional features such as time zone
// information and the inclusion of political views.
//
// For more information, see [Geocode] in the Amazon Location Service Developer Guide.
//
// [Geocode]: https://docs.aws.amazon.com/location/latest/developerguide/geocode.html
func geoplaces_Geocode(cfg aws.Config, client *geoplaces.Client) {
	input := &geoplaces.GeocodeInput{}

	if len(_geoplacesAdditionalFeatures) > 0 {
		if err := assignInputField(input, "AdditionalFeatures", _geoplacesAdditionalFeatures); err != nil {
			log.Errorf("invalid --additional-features: %s", err.Error())
			return
		}
	}
	if len(_geoplacesBiasPosition) > 0 {
		if err := assignInputField(input, "BiasPosition", _geoplacesBiasPosition); err != nil {
			log.Errorf("invalid --bias-position: %s", err.Error())
			return
		}
	}
	if len(_geoplacesFilter) > 0 {
		if err := assignInputField(input, "Filter", _geoplacesFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_geoplacesIntendedUse) > 0 {
		if err := assignInputField(input, "IntendedUse", _geoplacesIntendedUse); err != nil {
			log.Errorf("invalid --intended-use: %s", err.Error())
			return
		}
	}
	if len(_geoplacesKey) > 0 {
		input.Key = aws.String(_geoplacesKey)
	}
	if len(_geoplacesLanguage) > 0 {
		input.Language = aws.String(_geoplacesLanguage)
	}
	if len(_geoplacesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _geoplacesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_geoplacesPoliticalView) > 0 {
		input.PoliticalView = aws.String(_geoplacesPoliticalView)
	}
	if len(_geoplacesQueryComponents) > 0 {
		if err := assignInputField(input, "QueryComponents", _geoplacesQueryComponents); err != nil {
			log.Errorf("invalid --query-components: %s", err.Error())
			return
		}
	}
	if len(_geoplacesQueryText) > 0 {
		input.QueryText = aws.String(_geoplacesQueryText)
	}

	if resp, err := client.Geocode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// GetPlace finds a place by its unique ID. A PlaceId is returned by other place
// operations.
//
// For more information, see [GetPlace] in the Amazon Location Service Developer Guide.
//
// [GetPlace]: https://docs.aws.amazon.com/location/latest/developerguide/get-place.html
func geoplaces_GetPlace(cfg aws.Config, client *geoplaces.Client) {
	input := &geoplaces.GetPlaceInput{
		// PlaceId: *string, // Required
	}

	if len(_geoplacesPlaceId) > 0 {
		input.PlaceId = aws.String(_geoplacesPlaceId)
	}
	if len(_geoplacesAdditionalFeatures) > 0 {
		if err := assignInputField(input, "AdditionalFeatures", _geoplacesAdditionalFeatures); err != nil {
			log.Errorf("invalid --additional-features: %s", err.Error())
			return
		}
	}
	if len(_geoplacesIntendedUse) > 0 {
		if err := assignInputField(input, "IntendedUse", _geoplacesIntendedUse); err != nil {
			log.Errorf("invalid --intended-use: %s", err.Error())
			return
		}
	}
	if len(_geoplacesKey) > 0 {
		input.Key = aws.String(_geoplacesKey)
	}
	if len(_geoplacesLanguage) > 0 {
		input.Language = aws.String(_geoplacesLanguage)
	}
	if len(_geoplacesPoliticalView) > 0 {
		input.PoliticalView = aws.String(_geoplacesPoliticalView)
	}

	if resp, err := client.GetPlace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// ReverseGeocode converts geographic coordinates into a human-readable address or
// place. You can obtain address component, and other related information such as
// place type, category, street information. The Reverse Geocode API supports
// filtering to on place type so that you can refine result based on your need.
// Also, The Reverse Geocode API can also provide additional features such as time
// zone information and the inclusion of political views.
//
// For more information, see [Reverse Geocode] in the Amazon Location Service Developer Guide.
//
// [Reverse Geocode]: https://docs.aws.amazon.com/location/latest/developerguide/reverse-geocode.html
func geoplaces_ReverseGeocode(cfg aws.Config, client *geoplaces.Client) {
	input := &geoplaces.ReverseGeocodeInput{
		// QueryPosition: []float64, // Required
	}

	if len(_geoplacesQueryPosition) > 0 {
		if err := assignInputField(input, "QueryPosition", _geoplacesQueryPosition); err != nil {
			log.Errorf("invalid --query-position: %s", err.Error())
			return
		}
	}
	if len(_geoplacesAdditionalFeatures) > 0 {
		if err := assignInputField(input, "AdditionalFeatures", _geoplacesAdditionalFeatures); err != nil {
			log.Errorf("invalid --additional-features: %s", err.Error())
			return
		}
	}
	if len(_geoplacesFilter) > 0 {
		if err := assignInputField(input, "Filter", _geoplacesFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_geoplacesHeading) > 0 {
		if err := assignInputField(input, "Heading", _geoplacesHeading); err != nil {
			log.Errorf("invalid --heading: %s", err.Error())
			return
		}
	}
	if len(_geoplacesIntendedUse) > 0 {
		if err := assignInputField(input, "IntendedUse", _geoplacesIntendedUse); err != nil {
			log.Errorf("invalid --intended-use: %s", err.Error())
			return
		}
	}
	if len(_geoplacesKey) > 0 {
		input.Key = aws.String(_geoplacesKey)
	}
	if len(_geoplacesLanguage) > 0 {
		input.Language = aws.String(_geoplacesLanguage)
	}
	if len(_geoplacesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _geoplacesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_geoplacesPoliticalView) > 0 {
		input.PoliticalView = aws.String(_geoplacesPoliticalView)
	}
	if len(_geoplacesQueryRadius) > 0 {
		if err := assignInputField(input, "QueryRadius", _geoplacesQueryRadius); err != nil {
			log.Errorf("invalid --query-radius: %s", err.Error())
			return
		}
	}

	if resp, err := client.ReverseGeocode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// SearchNearby queries for points of interest within a radius from a central
// coordinates, returning place results with optional filters such as categories,
// business chains, food types and more. The API returns details such as a place
// name, address, phone, category, food type, contact, opening hours. Also, the API
// can return phonemes, time zones and more based on requested parameters.
//
// For more information, see [Search Nearby] in the Amazon Location Service Developer Guide.
//
// [Search Nearby]: https://docs.aws.amazon.com/location/latest/developerguide/search-nearby.html
func geoplaces_SearchNearby(cfg aws.Config, client *geoplaces.Client) {
	input := &geoplaces.SearchNearbyInput{
		// QueryPosition: []float64, // Required
	}

	if len(_geoplacesQueryPosition) > 0 {
		if err := assignInputField(input, "QueryPosition", _geoplacesQueryPosition); err != nil {
			log.Errorf("invalid --query-position: %s", err.Error())
			return
		}
	}
	if len(_geoplacesAdditionalFeatures) > 0 {
		if err := assignInputField(input, "AdditionalFeatures", _geoplacesAdditionalFeatures); err != nil {
			log.Errorf("invalid --additional-features: %s", err.Error())
			return
		}
	}
	if len(_geoplacesFilter) > 0 {
		if err := assignInputField(input, "Filter", _geoplacesFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_geoplacesIntendedUse) > 0 {
		if err := assignInputField(input, "IntendedUse", _geoplacesIntendedUse); err != nil {
			log.Errorf("invalid --intended-use: %s", err.Error())
			return
		}
	}
	if len(_geoplacesKey) > 0 {
		input.Key = aws.String(_geoplacesKey)
	}
	if len(_geoplacesLanguage) > 0 {
		input.Language = aws.String(_geoplacesLanguage)
	}
	if len(_geoplacesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _geoplacesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_geoplacesNextToken) > 0 {
		input.NextToken = aws.String(_geoplacesNextToken)
	}
	if len(_geoplacesPoliticalView) > 0 {
		input.PoliticalView = aws.String(_geoplacesPoliticalView)
	}
	if len(_geoplacesQueryRadius) > 0 {
		if err := assignInputField(input, "QueryRadius", _geoplacesQueryRadius); err != nil {
			log.Errorf("invalid --query-radius: %s", err.Error())
			return
		}
	}

	if resp, err := client.SearchNearby(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// SearchText searches for geocode and place information. You can then complete a
// follow-up query suggested from the Suggest API via a query id.
//
// For more information, see [Search Text] in the Amazon Location Service Developer Guide.
//
// [Search Text]: https://docs.aws.amazon.com/location/latest/developerguide/search-text.html
func geoplaces_SearchText(cfg aws.Config, client *geoplaces.Client) {
	input := &geoplaces.SearchTextInput{}

	if len(_geoplacesAdditionalFeatures) > 0 {
		if err := assignInputField(input, "AdditionalFeatures", _geoplacesAdditionalFeatures); err != nil {
			log.Errorf("invalid --additional-features: %s", err.Error())
			return
		}
	}
	if len(_geoplacesBiasPosition) > 0 {
		if err := assignInputField(input, "BiasPosition", _geoplacesBiasPosition); err != nil {
			log.Errorf("invalid --bias-position: %s", err.Error())
			return
		}
	}
	if len(_geoplacesFilter) > 0 {
		if err := assignInputField(input, "Filter", _geoplacesFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_geoplacesIntendedUse) > 0 {
		if err := assignInputField(input, "IntendedUse", _geoplacesIntendedUse); err != nil {
			log.Errorf("invalid --intended-use: %s", err.Error())
			return
		}
	}
	if len(_geoplacesKey) > 0 {
		input.Key = aws.String(_geoplacesKey)
	}
	if len(_geoplacesLanguage) > 0 {
		input.Language = aws.String(_geoplacesLanguage)
	}
	if len(_geoplacesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _geoplacesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_geoplacesNextToken) > 0 {
		input.NextToken = aws.String(_geoplacesNextToken)
	}
	if len(_geoplacesPoliticalView) > 0 {
		input.PoliticalView = aws.String(_geoplacesPoliticalView)
	}
	if len(_geoplacesQueryId) > 0 {
		input.QueryId = aws.String(_geoplacesQueryId)
	}
	if len(_geoplacesQueryText) > 0 {
		input.QueryText = aws.String(_geoplacesQueryText)
	}

	if resp, err := client.SearchText(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Suggest provides intelligent predictions or recommendations based on the user's
// input or context, such as relevant places, points of interest, query terms or
// search category. It is designed to help users find places or point of interests
// candidates or identify a follow on query based on incomplete or misspelled
// queries. It returns a list of possible matches or refinements that can be used
// to formulate a more accurate query. Users can select the most appropriate
// suggestion and use it for further searching. The API provides options for
// filtering results by location and other attributes, and allows for additional
// features like phonemes and timezones. The response includes refined query terms
// and detailed place information.
//
// For more information, see [Suggest] in the Amazon Location Service Developer Guide.
//
// [Suggest]: https://docs.aws.amazon.com/location/latest/developerguide/suggest.html
func geoplaces_Suggest(cfg aws.Config, client *geoplaces.Client) {
	input := &geoplaces.SuggestInput{
		// QueryText: *string, // Required
	}

	if len(_geoplacesQueryText) > 0 {
		input.QueryText = aws.String(_geoplacesQueryText)
	}
	if len(_geoplacesAdditionalFeatures) > 0 {
		if err := assignInputField(input, "AdditionalFeatures", _geoplacesAdditionalFeatures); err != nil {
			log.Errorf("invalid --additional-features: %s", err.Error())
			return
		}
	}
	if len(_geoplacesBiasPosition) > 0 {
		if err := assignInputField(input, "BiasPosition", _geoplacesBiasPosition); err != nil {
			log.Errorf("invalid --bias-position: %s", err.Error())
			return
		}
	}
	if len(_geoplacesFilter) > 0 {
		if err := assignInputField(input, "Filter", _geoplacesFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_geoplacesIntendedUse) > 0 {
		if err := assignInputField(input, "IntendedUse", _geoplacesIntendedUse); err != nil {
			log.Errorf("invalid --intended-use: %s", err.Error())
			return
		}
	}
	if len(_geoplacesKey) > 0 {
		input.Key = aws.String(_geoplacesKey)
	}
	if len(_geoplacesLanguage) > 0 {
		input.Language = aws.String(_geoplacesLanguage)
	}
	if len(_geoplacesMaxQueryRefinements) > 0 {
		if err := assignInputField(input, "MaxQueryRefinements", _geoplacesMaxQueryRefinements); err != nil {
			log.Errorf("invalid --max-query-refinements: %s", err.Error())
			return
		}
	}
	if len(_geoplacesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _geoplacesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_geoplacesPoliticalView) > 0 {
		input.PoliticalView = aws.String(_geoplacesPoliticalView)
	}

	if resp, err := client.Suggest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_geoplacesCmd)
	_geoplacesCmd.Flags().SortFlags = false

	_geoplacesCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_geoplacesCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_geoplacesCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_geoplacesCmd.Flags().StringVarP(&_geoplacesAdditionalFeatures, "additional-features", "", "", "Additional Features")
	_geoplacesCmd.Flags().StringVarP(&_geoplacesBiasPosition, "bias-position", "", "", "Bias Position")
	_geoplacesCmd.Flags().StringVarP(&_geoplacesFilter, "filter", "", "", "Filter")
	_geoplacesCmd.Flags().StringVarP(&_geoplacesHeading, "heading", "", "", "Heading")
	_geoplacesCmd.Flags().StringVarP(&_geoplacesIntendedUse, "intended-use", "", "", "Intended Use")
	_geoplacesCmd.Flags().StringVarP(&_geoplacesKey, "key", "", "", "Key")
	_geoplacesCmd.Flags().StringVarP(&_geoplacesLanguage, "language", "", "", "Language")
	_geoplacesCmd.Flags().StringVarP(&_geoplacesMaxQueryRefinements, "max-query-refinements", "", "", "Max Query Refinements")
	_geoplacesCmd.Flags().StringVarP(&_geoplacesMaxResults, "max-results", "", "", "Max Results")
	_geoplacesCmd.Flags().StringVarP(&_geoplacesNextToken, "next-token", "", "", "Next Token")
	_geoplacesCmd.Flags().StringVarP(&_geoplacesPlaceId, "place-id", "", "", "Place ID")
	_geoplacesCmd.Flags().StringVarP(&_geoplacesPoliticalView, "political-view", "", "", "Political View")
	_geoplacesCmd.Flags().StringVarP(&_geoplacesPostalCodeMode, "postal-code-mode", "", "", "Postal Code Mode")
	_geoplacesCmd.Flags().StringVarP(&_geoplacesQueryComponents, "query-components", "", "", "Query Components")
	_geoplacesCmd.Flags().StringVarP(&_geoplacesQueryId, "query-id", "", "", "Query ID")
	_geoplacesCmd.Flags().StringVarP(&_geoplacesQueryPosition, "query-position", "", "", "Query Position")
	_geoplacesCmd.Flags().StringVarP(&_geoplacesQueryRadius, "query-radius", "", "", "Query Radius")
	_geoplacesCmd.Flags().StringVarP(&_geoplacesQueryText, "query-text", "", "", "Query Text")

	_geoplacesCmd.Flags().BoolVarP(&_geoplacesAutocomplete, "autocomplete", "", false, "Autocomplete")
	_geoplacesCmd.Flags().BoolVarP(&_geoplacesGeocode, "geocode", "", false, "Geocode")
	_geoplacesCmd.Flags().BoolVarP(&_geoplacesGetPlace, "get-place", "", false, "Get Place")
	_geoplacesCmd.Flags().BoolVarP(&_geoplacesReverseGeocode, "reverse-geocode", "", false, "Reverse Geocode")
	_geoplacesCmd.Flags().BoolVarP(&_geoplacesSearchNearby, "search-nearby", "", false, "Search Nearby")
	_geoplacesCmd.Flags().BoolVarP(&_geoplacesSearchText, "search-text", "", false, "Search Text")
	_geoplacesCmd.Flags().BoolVarP(&_geoplacesSuggest, "suggest", "", false, "Suggest")

}
