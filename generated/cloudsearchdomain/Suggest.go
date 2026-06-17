package cloudsearchdomain

// Suggest is generated as a reference stub.
// Executable command wiring lives under cmd/cloudsearchdomain.go.
//
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
