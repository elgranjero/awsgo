package ec2

// CreateCapacityReservation is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates a new Capacity Reservation with the specified attributes. Capacity
// Reservations enable you to reserve capacity for your Amazon EC2 instances in a
// specific Availability Zone for any duration.
//
// You can create a Capacity Reservation at any time, and you can choose when it
// starts. You can create a Capacity Reservation for immediate use or you can
// request a Capacity Reservation for a future date.
//
// For more information, see [Reserve compute capacity with On-Demand Capacity Reservations] in the Amazon EC2 User Guide.
//
// Your request to create a Capacity Reservation could fail if:
//
// - Amazon EC2 does not have sufficient capacity. In this case, try again at a
// later time, try in a different Availability Zone, or request a smaller Capacity
// Reservation. If your workload is flexible across instance types and sizes, try
// with different instance attributes.
//
// - The requested quantity exceeds your On-Demand Instance quota. In this case,
// increase your On-Demand Instance quota for the requested instance type and try
// again. For more information, see [Amazon EC2 Service Quotas]in the Amazon EC2 User Guide.
//
// [Amazon EC2 Service Quotas]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-resource-limits.html
// [Reserve compute capacity with On-Demand Capacity Reservations]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-capacity-reservations.html
