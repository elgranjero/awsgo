package marketplacemetering

// ResolveCustomer is generated as a reference stub.
// Executable command wiring lives under cmd/marketplacemetering.go.
//
// ResolveCustomer is called by a SaaS application during the registration
// process. When a buyer visits your website during the registration process, the
// buyer submits a registration token through their browser. The registration token
// is resolved through this API to obtain a CustomerIdentifier along with the
// CustomerAWSAccountId , ProductCode , and LicenseArn .
//
// To successfully resolve the token, the API must be called from the account that
// was used to publish the SaaS application. For an example of using
// ResolveCustomer , see [ResolveCustomer code example] in the Amazon Web Services Marketplace Seller Guide.
//
// Permission is required for this operation. Your IAM role or user performing
// this operation requires a policy to allow the aws-marketplace:ResolveCustomer
// action. For more information, see [Actions, resources, and condition keys for Amazon Web Services Marketplace Metering Service]in the Service Authorization Reference.
//
// For Amazon Web Services Regions that support ResolveCustomer , see [ResolveCustomer Region support].
//
// [ResolveCustomer code example]: https://docs.aws.amazon.com/marketplace/latest/userguide/saas-code-examples.html#saas-resolvecustomer-example
// [Actions, resources, and condition keys for Amazon Web Services Marketplace Metering Service]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsmarketplacemeteringservice.html
// [ResolveCustomer Region support]: https://docs.aws.amazon.com/marketplace/latest/APIReference/metering-regions.html#resolvecustomer-region-support
