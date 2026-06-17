package elasticsearchservice

// DeleteElasticsearchServiceRole is generated as a reference stub.
// Executable command wiring lives under cmd/elasticsearchservice.go.
//
// Deletes the service-linked role that Elasticsearch Service uses to manage and
// maintain VPC domains. Role deletion will fail if any existing VPC domains use
// the role. You must delete any such Elasticsearch domains before deleting the
// role. See [Deleting Elasticsearch Service Role]in VPC Endpoints for Amazon Elasticsearch Service Domains.
//
// [Deleting Elasticsearch Service Role]: http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-enabling-slr
