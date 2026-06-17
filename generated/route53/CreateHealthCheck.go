package route53

// CreateHealthCheck is generated as a reference stub.
// Executable command wiring lives under cmd/route53.go.
//
// Creates a new health check.
//
// For information about adding health checks to resource record sets, see [HealthCheckId] in [ChangeResourceRecordSets].
//
// # ELB Load Balancers
//
// If you're registering EC2 instances with an Elastic Load Balancing (ELB) load
// balancer, do not create Amazon Route 53 health checks for the EC2 instances.
// When you register an EC2 instance with a load balancer, you configure settings
// for an ELB health check, which performs a similar function to a Route 53 health
// check.
//
// # Private Hosted Zones
//
// You can associate health checks with failover resource record sets in a private
// hosted zone. Note the following:
//
// - Route 53 health checkers are outside the VPC. To check the health of an
// endpoint within a VPC by IP address, you must assign a public IP address to the
// instance in the VPC.
//
// - You can configure a health checker to check the health of an external
// resource that the instance relies on, such as a database server.
//
// - You can create a CloudWatch metric, associate an alarm with the metric, and
// then create a health check that is based on the state of the alarm. For example,
// you might create a CloudWatch metric that checks the status of the Amazon EC2
// StatusCheckFailed metric, add an alarm to the metric, and then create a health
// check that is based on the state of the alarm. For information about creating
// CloudWatch metrics and alarms by using the CloudWatch console, see the [Amazon CloudWatch User Guide].
//
// [HealthCheckId]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_ResourceRecordSet.html#Route53-Type-ResourceRecordSet-HealthCheckId
// [ChangeResourceRecordSets]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_ChangeResourceRecordSets.html
// [Amazon CloudWatch User Guide]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/WhatIsCloudWatch.html
