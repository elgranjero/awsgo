package cloudtrail

// PutEventSelectors is generated as a reference stub.
// Executable command wiring lives under cmd/cloudtrail.go.
//
// Configures event selectors (also referred to as basic event selectors) or
// advanced event selectors for your trail. You can use either
// AdvancedEventSelectors or EventSelectors , but not both. If you apply
// AdvancedEventSelectors to a trail, any existing EventSelectors are overwritten.
//
// You can use AdvancedEventSelectors to log management events, data events for
// all resource types, and network activity events.
//
// You can use EventSelectors to log management events and data events for the
// following resource types:
//
// - AWS::DynamoDB::Table
//
// - AWS::Lambda::Function
//
// - AWS::S3::Object
//
// You can't use EventSelectors to log network activity events.
//
// If you want your trail to log Insights events, be sure the event selector or
// advanced event selector enables logging of the Insights event types you want
// configured for your trail. For more information about logging Insights events,
// see [Working with CloudTrail Insights]in the CloudTrail User Guide. By default, trails created without specific
// event selectors are configured to log all read and write management events, and
// no data events or network activity events.
//
// When an event occurs in your account, CloudTrail evaluates the event selectors
// or advanced event selectors in all trails. For each trail, if the event matches
// any event selector, the trail processes and logs the event. If the event doesn't
// match any event selector, the trail doesn't log the event.
//
// Example
//
// - You create an event selector for a trail and specify that you want to log
// write-only events.
//
// - The EC2 GetConsoleOutput and RunInstances API operations occur in your
// account.
//
// - CloudTrail evaluates whether the events match your event selectors.
//
// - The RunInstances is a write-only event and it matches your event selector.
// The trail logs the event.
//
// - The GetConsoleOutput is a read-only event that doesn't match your event
// selector. The trail doesn't log the event.
//
// The PutEventSelectors operation must be called from the Region in which the
// trail was created; otherwise, an InvalidHomeRegionException exception is thrown.
//
// You can configure up to five event selectors for each trail.
//
// You can add advanced event selectors, and conditions for your advanced event
// selectors, up to a maximum of 500 values for all conditions and selectors on a
// trail. For more information, see [Logging management events], [Logging data events], [Logging network activity events], and [Quotas in CloudTrail] in the CloudTrail User Guide.
//
// [Logging network activity events]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-network-events-with-cloudtrail.html
// [Logging management events]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-management-events-with-cloudtrail.html
// [Working with CloudTrail Insights]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-insights-events-with-cloudtrail.html
// [Quotas in CloudTrail]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html
// [Logging data events]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html
