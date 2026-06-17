package partnercentralselling

// AssociateOpportunity is generated as a reference stub.
// Executable command wiring lives under cmd/partnercentralselling.go.
//
// Enables you to create a formal association between an Opportunity and various
// related entities, enriching the context and details of the opportunity for
// better collaboration and decision making. You can associate an opportunity with
// the following entity types:
//
// - Partner Solution: A software product or consulting practice created and
// delivered by Partners. Partner Solutions help customers address business
// challenges using Amazon Web Services services.
//
// - Amazon Web Services Products: Amazon Web Services offers many products and
// services that provide scalable, reliable, and cost-effective infrastructure
// solutions. For the latest list of Amazon Web Services products, see [Amazon Web Services products].
//
// - Amazon Web Services Marketplace private offer: Allows Amazon Web Services
// Marketplace sellers to extend custom pricing and terms to individual Amazon Web
// Services customers. Sellers can negotiate custom prices, payment schedules, and
// end user license terms through private offers, enabling Amazon Web Services
// customers to acquire software solutions tailored to their specific needs. For
// more information, see [Private offers in Amazon Web Services Marketplace].
//
// To obtain identifiers for these entities, use the following methods:
//
// - Solution: Use the ListSolutions operation.
//
// - AWS Products: For the latest list of Amazon Web Services products, see [Amazon Web Services products].
//
// - Amazon Web Services Marketplace private offer: Use the [Using the Amazon Web Services Marketplace Catalog API]to list entities.
// Specifically, use the ListEntities operation to retrieve a list of private
// offers. The request returns the details of available private offers. For more
// information, see [ListEntities].
//
// [Private offers in Amazon Web Services Marketplace]: https://docs.aws.amazon.com/marketplace/latest/buyerguide/buyer-private-offers.html
// [Using the Amazon Web Services Marketplace Catalog API]: https://docs.aws.amazon.com/marketplace/latest/APIReference/catalog-apis.html
// [ListEntities]: https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/API_ListEntities.html
// [Amazon Web Services products]: https://github.com/aws-samples/partner-crm-integration-samples/blob/main/resources/aws_products.json
