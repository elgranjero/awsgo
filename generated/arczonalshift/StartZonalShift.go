package arczonalshift

// StartZonalShift is generated as a reference stub.
// Executable command wiring lives under cmd/arczonalshift.go.
//
// You start a zonal shift to temporarily move load balancer traffic away from an
// Availability Zone in an Amazon Web Services Region, to help your application
// recover immediately, for example, from a developer's bad code deployment or from
// an Amazon Web Services infrastructure failure in a single Availability Zone. You
// can start a zonal shift in ARC only for managed resources in your Amazon Web
// Services account in an Amazon Web Services Region. Resources are automatically
// registered with ARC by Amazon Web Services services.
//
// Amazon Application Recovery Controller currently supports enabling the
// following resources for zonal shift and zonal autoshift:
//
// [Amazon EC2 Auto Scaling groups]
//
// [Amazon Elastic Kubernetes Service]
//
// [Application Load Balancer]
//
// [Network Load Balancer]
//
// When you start a zonal shift, traffic for the resource is no longer routed to
// the Availability Zone. The zonal shift is created immediately in ARC. However,
// it can take a short time, typically up to a few minutes, for existing,
// in-progress connections in the Availability Zone to complete.
//
// For more information, see [Zonal shift] in the Amazon Application Recovery Controller
// Developer Guide.
//
// [Amazon EC2 Auto Scaling groups]: https://docs.aws.amazon.com/r53recovery/latest/dg/arc-zonal-shift.resource-types.ec2-auto-scaling-groups.html
// [Amazon Elastic Kubernetes Service]: https://docs.aws.amazon.com/r53recovery/latest/dg/arc-zonal-shift.resource-types.eks.html
// [Application Load Balancer]: https://docs.aws.amazon.com/r53recovery/latest/dg/arc-zonal-shift.resource-types.app-load-balancers.html
// [Network Load Balancer]: https://docs.aws.amazon.com/r53recovery/latest/dg/arc-zonal-shift.resource-types.network-load-balancers.html
// [Zonal shift]: https://docs.aws.amazon.com/r53recovery/latest/dg/arc-zonal-shift.html
