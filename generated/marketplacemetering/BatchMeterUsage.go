package marketplacemetering

// BatchMeterUsage is generated as a reference stub.
// Executable command wiring lives under cmd/marketplacemetering.go.
//
// Amazon Web Services Marketplace is introducing Concurrent Agreements, enabling
// buyers to make multiple purchases per Amazon Web Services account. Starting June
// 1, 2026, new SaaS products must use CustomerAWSAccountId (instead of
// CustomerIdentifier ), LicenseArn (instead of ProductCode ) to support this
// feature. Existing integrations will continue to work. Review the new integration
// for Concurrent Agreements [here].
//
// To post metering records for customers, SaaS applications call BatchMeterUsage ,
// which is used for metering SaaS flexible consumption pricing (FCP). Identical
// requests are idempotent and can be retried with the same records or a subset of
// records. Each BatchMeterUsage request is for only one product. If you want to
// meter usage for multiple products, you must make multiple BatchMeterUsage calls.
//
// Usage records should be submitted in quick succession following a recorded
// event. Usage records aren't accepted 6 hours or more after an event.
//
// BatchMeterUsage can process up to 25 UsageRecords at a time, and each request
// must be less than 1 MB in size. Optionally, you can have multiple usage
// allocations for usage data that's split into buckets according to predefined
// tags.
//
// BatchMeterUsage returns a list of UsageRecordResult objects, which have each
// UsageRecord . It also returns a list of UnprocessedRecords , which indicate
// errors on the service side that should be retried.
//
// For Amazon Web Services Regions that support BatchMeterUsage , see [BatchMeterUsage Region support].
//
// For an example of BatchMeterUsage , see [BatchMeterUsage code example] in the Amazon Web Services Marketplace
// Seller Guide.
//
// [here]: https://catalog.workshops.aws/mpseller/en-US/saas/integration-for-concurrent-agreements
// [BatchMeterUsage code example]: https://docs.aws.amazon.com/marketplace/latest/userguide/saas-code-examples.html#saas-batchmeterusage-example
// [BatchMeterUsage Region support]: https://docs.aws.amazon.com/marketplace/latest/APIReference/metering-regions.html#batchmeterusage-region-support
