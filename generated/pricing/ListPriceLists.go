package pricing

// ListPriceLists is generated as a reference stub.
// Executable command wiring lives under cmd/pricing.go.
//
// This feature is in preview release and is subject to change. Your use of
//
// Amazon Web Services Price List API is subject to the Beta Service Participation
// terms of the [Amazon Web Services Service Terms](Section 1.10).
//
// This returns a list of Price List references that the requester if authorized
// to view, given a ServiceCode , CurrencyCode , and an EffectiveDate . Use without
// a RegionCode filter to list Price List references from all available Amazon Web
// Services Regions. Use with a RegionCode filter to get the Price List reference
// that's specific to a specific Amazon Web Services Region. You can use the
// PriceListArn from the response to get your preferred Price List files through
// the [GetPriceListFileUrl]API.
//
// [Amazon Web Services Service Terms]: https://aws.amazon.com/service-terms/
// [GetPriceListFileUrl]: https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_pricing_GetPriceListFileUrl.html
