package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearchdomain"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cloudsearchdomainCmd represents the cloudsearchdomain command
var _cloudsearchdomainCmd = &cobra.Command{
	Use:   "cloudsearchdomain",
	Short: "AWS cloudsearchdomain CLI",
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
		client := cloudsearchdomain.NewFromConfig(cfg)
		if _cloudsearchdomainSearch {
			cloudsearchdomain_Search(cfg, client)
			return
		}
		if _cloudsearchdomainSuggest {
			cloudsearchdomain_Suggest(cfg, client)
			return
		}
		if _cloudsearchdomainUploadDocuments {
			cloudsearchdomain_UploadDocuments(cfg, client)
			return
		}

	},
}

var (
	_cloudsearchdomainSearch          bool
	_cloudsearchdomainSuggest         bool
	_cloudsearchdomainUploadDocuments bool

	_cloudsearchdomainContentType  string
	_cloudsearchdomainCursor       string
	_cloudsearchdomainDocuments    string
	_cloudsearchdomainExpr         string
	_cloudsearchdomainFacet        string
	_cloudsearchdomainFilterQuery  string
	_cloudsearchdomainHighlight    string
	_cloudsearchdomainPartial      string
	_cloudsearchdomainQuery        string
	_cloudsearchdomainQueryOptions string
	_cloudsearchdomainQueryParser  string
	_cloudsearchdomainReturn       string
	_cloudsearchdomainSize         string
	_cloudsearchdomainSort         string
	_cloudsearchdomainStart        string
	_cloudsearchdomainStats        string
	_cloudsearchdomainSuggester    string
)

// Retrieves a list of documents that match the specified search criteria. How you
// specify the search criteria depends on which query parser you use. Amazon
// CloudSearch supports four query parsers:
//
// - simple : search all text and text-array fields for the specified string.
// Search for phrases, individual terms, and prefixes.
// - structured : search specific fields, construct compound queries using
// Boolean operators, and use advanced features such as term boosting and proximity
// searching.
// - lucene : specify search criteria using the Apache Lucene query parser syntax.
// - dismax : specify search criteria using the simplified subset of the Apache
// Lucene query parser syntax defined by the DisMax query parser.
//
// For more information, see [Searching Your Data] in the Amazon CloudSearch Developer Guide.
//
// The endpoint for submitting Search requests is domain-specific. You submit
// search requests to a domain's search endpoint. To get the search endpoint for
// your domain, use the Amazon CloudSearch configuration service DescribeDomains
// action. A domain's endpoints are also displayed on the domain dashboard in the
// Amazon CloudSearch console.
//
// [Searching Your Data]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/searching.html
func cloudsearchdomain_Search(cfg aws.Config, client *cloudsearchdomain.Client) {
	input := &cloudsearchdomain.SearchInput{
		// Query: *string, // Required
	}

	if len(_cloudsearchdomainQuery) > 0 {
		input.Query = aws.String(_cloudsearchdomainQuery)
	}
	if len(_cloudsearchdomainCursor) > 0 {
		input.Cursor = aws.String(_cloudsearchdomainCursor)
	}
	if len(_cloudsearchdomainExpr) > 0 {
		input.Expr = aws.String(_cloudsearchdomainExpr)
	}
	if len(_cloudsearchdomainFacet) > 0 {
		input.Facet = aws.String(_cloudsearchdomainFacet)
	}
	if len(_cloudsearchdomainFilterQuery) > 0 {
		input.FilterQuery = aws.String(_cloudsearchdomainFilterQuery)
	}
	if len(_cloudsearchdomainHighlight) > 0 {
		input.Highlight = aws.String(_cloudsearchdomainHighlight)
	}
	if len(_cloudsearchdomainPartial) > 0 {
		if err := assignInputField(input, "Partial", _cloudsearchdomainPartial); err != nil {
			log.Errorf("invalid --partial: %s", err.Error())
			return
		}
	}
	if len(_cloudsearchdomainQueryOptions) > 0 {
		input.QueryOptions = aws.String(_cloudsearchdomainQueryOptions)
	}
	if len(_cloudsearchdomainQueryParser) > 0 {
		if err := assignInputField(input, "QueryParser", _cloudsearchdomainQueryParser); err != nil {
			log.Errorf("invalid --query-parser: %s", err.Error())
			return
		}
	}
	if len(_cloudsearchdomainReturn) > 0 {
		input.Return = aws.String(_cloudsearchdomainReturn)
	}
	if len(_cloudsearchdomainSize) > 0 {
		if err := assignInputField(input, "Size", _cloudsearchdomainSize); err != nil {
			log.Errorf("invalid --size: %s", err.Error())
			return
		}
	}
	if len(_cloudsearchdomainSort) > 0 {
		input.Sort = aws.String(_cloudsearchdomainSort)
	}
	if len(_cloudsearchdomainStart) > 0 {
		if err := assignInputField(input, "Start", _cloudsearchdomainStart); err != nil {
			log.Errorf("invalid --start: %s", err.Error())
			return
		}
	}
	if len(_cloudsearchdomainStats) > 0 {
		input.Stats = aws.String(_cloudsearchdomainStats)
	}

	if resp, err := client.Search(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves autocomplete suggestions for a partial query string. You can use
// suggestions enable you to display likely matches before users finish typing. In
// Amazon CloudSearch, suggestions are based on the contents of a particular text
// field. When you request suggestions, Amazon CloudSearch finds all of the
// documents whose values in the suggester field start with the specified query
// string. The beginning of the field must match the query string to be considered
// a match.
//
// For more information about configuring suggesters and retrieving suggestions,
// see [Getting Suggestions]in the Amazon CloudSearch Developer Guide.
//
// The endpoint for submitting Suggest requests is domain-specific. You submit
// suggest requests to a domain's search endpoint. To get the search endpoint for
// your domain, use the Amazon CloudSearch configuration service DescribeDomains
// action. A domain's endpoints are also displayed on the domain dashboard in the
// Amazon CloudSearch console.
//
// [Getting Suggestions]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/getting-suggestions.html
func cloudsearchdomain_Suggest(cfg aws.Config, client *cloudsearchdomain.Client) {
	input := &cloudsearchdomain.SuggestInput{
		// Query: *string, // Required
		// Suggester: *string, // Required
	}

	if len(_cloudsearchdomainQuery) > 0 {
		input.Query = aws.String(_cloudsearchdomainQuery)
	}
	if len(_cloudsearchdomainSuggester) > 0 {
		input.Suggester = aws.String(_cloudsearchdomainSuggester)
	}
	if len(_cloudsearchdomainSize) > 0 {
		if err := assignInputField(input, "Size", _cloudsearchdomainSize); err != nil {
			log.Errorf("invalid --size: %s", err.Error())
			return
		}
	}

	if resp, err := client.Suggest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Posts a batch of documents to a search domain for indexing. A document batch is
// a collection of add and delete operations that represent the documents you want
// to add, update, or delete from your domain. Batches can be described in either
// JSON or XML. Each item that you want Amazon CloudSearch to return as a search
// result (such as a product) is represented as a document. Every document has a
// unique ID and one or more fields that contain the data that you want to search
// and return in results. Individual documents cannot contain more than 1 MB of
// data. The entire batch cannot exceed 5 MB. To get the best possible upload
// performance, group add and delete operations in batches that are close the 5 MB
// limit. Submitting a large volume of single-document batches can overload a
// domain's document service.
//
// The endpoint for submitting UploadDocuments requests is domain-specific. To get
// the document endpoint for your domain, use the Amazon CloudSearch configuration
// service DescribeDomains action. A domain's endpoints are also displayed on the
// domain dashboard in the Amazon CloudSearch console.
//
// For more information about formatting your data for Amazon CloudSearch, see [Preparing Your Data] in
// the Amazon CloudSearch Developer Guide. For more information about uploading
// data for indexing, see [Uploading Data]in the Amazon CloudSearch Developer Guide.
//
// [Preparing Your Data]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/preparing-data.html
// [Uploading Data]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/uploading-data.html
func cloudsearchdomain_UploadDocuments(cfg aws.Config, client *cloudsearchdomain.Client) {
	input := &cloudsearchdomain.UploadDocumentsInput{
		// ContentType: types.ContentType, // Required
		// Documents: io.Reader, // Required
	}

	if len(_cloudsearchdomainContentType) > 0 {
		if err := assignInputField(input, "ContentType", _cloudsearchdomainContentType); err != nil {
			log.Errorf("invalid --content-type: %s", err.Error())
			return
		}
	}
	if len(_cloudsearchdomainDocuments) > 0 {
		if err := assignInputField(input, "Documents", _cloudsearchdomainDocuments); err != nil {
			log.Errorf("invalid --documents: %s", err.Error())
			return
		}
	}

	if resp, err := client.UploadDocuments(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_cloudsearchdomainCmd)
	_cloudsearchdomainCmd.Flags().SortFlags = false

	_cloudsearchdomainCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_cloudsearchdomainCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_cloudsearchdomainCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_cloudsearchdomainCmd.Flags().StringVarP(&_cloudsearchdomainContentType, "content-type", "", "", "Content Type")
	_cloudsearchdomainCmd.Flags().StringVarP(&_cloudsearchdomainCursor, "cursor", "", "", "Cursor")
	_cloudsearchdomainCmd.Flags().StringVarP(&_cloudsearchdomainDocuments, "documents", "", "", "Documents")
	_cloudsearchdomainCmd.Flags().StringVarP(&_cloudsearchdomainExpr, "expr", "", "", "Expr")
	_cloudsearchdomainCmd.Flags().StringVarP(&_cloudsearchdomainFacet, "facet", "", "", "Facet")
	_cloudsearchdomainCmd.Flags().StringVarP(&_cloudsearchdomainFilterQuery, "filter-query", "", "", "Filter Query")
	_cloudsearchdomainCmd.Flags().StringVarP(&_cloudsearchdomainHighlight, "highlight", "", "", "Highlight")
	_cloudsearchdomainCmd.Flags().StringVarP(&_cloudsearchdomainPartial, "partial", "", "", "Partial")
	_cloudsearchdomainCmd.Flags().StringVarP(&_cloudsearchdomainQuery, "query", "", "", "Query")
	_cloudsearchdomainCmd.Flags().StringVarP(&_cloudsearchdomainQueryOptions, "query-options", "", "", "Query Options")
	_cloudsearchdomainCmd.Flags().StringVarP(&_cloudsearchdomainQueryParser, "query-parser", "", "", "Query Parser")
	_cloudsearchdomainCmd.Flags().StringVarP(&_cloudsearchdomainReturn, "return", "", "", "Return")
	_cloudsearchdomainCmd.Flags().StringVarP(&_cloudsearchdomainSize, "size", "", "", "Size")
	_cloudsearchdomainCmd.Flags().StringVarP(&_cloudsearchdomainSort, "sort", "", "", "Sort")
	_cloudsearchdomainCmd.Flags().StringVarP(&_cloudsearchdomainStart, "start", "", "", "Start")
	_cloudsearchdomainCmd.Flags().StringVarP(&_cloudsearchdomainStats, "stats", "", "", "Stats")
	_cloudsearchdomainCmd.Flags().StringVarP(&_cloudsearchdomainSuggester, "suggester", "", "", "Suggester")

	_cloudsearchdomainCmd.Flags().BoolVarP(&_cloudsearchdomainSearch, "search", "", false, "Search")
	_cloudsearchdomainCmd.Flags().BoolVarP(&_cloudsearchdomainSuggest, "suggest", "", false, "Suggest")
	_cloudsearchdomainCmd.Flags().BoolVarP(&_cloudsearchdomainUploadDocuments, "upload-documents", "", false, "Upload Documents")

}
