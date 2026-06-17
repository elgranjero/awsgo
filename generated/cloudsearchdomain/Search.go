package cloudsearchdomain

// Search is generated as a reference stub.
// Executable command wiring lives under cmd/cloudsearchdomain.go.
//
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
