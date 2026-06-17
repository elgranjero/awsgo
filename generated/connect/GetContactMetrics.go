package connect

// GetContactMetrics is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Retrieves contact metric data for a specified contact.
//
// # Use cases
//
// Following are common use cases for position in queue and estimated wait time:
//
// - Customer-Facing Wait Time Announcements - Display or announce the estimated
// wait time and position in queue to customers before or during their queue
// experience.
//
// - Callback Offerings - Offer customers a callback option when the estimated
// wait time or position in queue exceeds a defined threshold.
//
// - Queue Routing Decisions - Route incoming contacts to less congested queues
// by comparing estimated wait time and position in queue across multiple queues.
//
// - Self-Service Deflection - Redirect customers to self-service options like
// chatbots or FAQs when estimated wait time is high or position in queue is
// unfavorable.
//
// Important things to know
//
// - Metrics are only available while the contact is actively in queue.
//
// - For more information, see the [Position in queue]metric in the Amazon Connect Administrator
// Guide.
//
// Endpoints: See [Amazon Connect endpoints and quotas].
//
// [Position in queue]: https://docs.aws.amazon.com/connect/latest/adminguide/metrics-definitions.html
// [Amazon Connect endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/connect_region.html
