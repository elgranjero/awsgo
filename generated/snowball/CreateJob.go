package snowball

// CreateJob is generated as a reference stub.
// Executable command wiring lives under cmd/snowball.go.
//
// Creates a job to import or export data between Amazon S3 and your on-premises
// data center. Your Amazon Web Services account must have the right trust policies
// and permissions in place to create a job for a Snow device. If you're creating a
// job for a node in a cluster, you only need to provide the clusterId value; the
// other job attributes are inherited from the cluster.
//
// Only the Snowball; Edge device type is supported when ordering clustered jobs.
//
// The device capacity is optional.
//
// Availability of device types differ by Amazon Web Services Region. For more
// information about Region availability, see [Amazon Web Services Regional Services].
//
// Snow Family devices and their capacities.
//
// - Device type: SNC1_SSD
//
// - Capacity: T14
//
// - Description: Snowcone
//
// - Device type: SNC1_HDD
//
// - Capacity: T8
//
// - Description: Snowcone
//
// - Device type: EDGE_S
//
// - Capacity: T98
//
// - Description: Snowball Edge Storage Optimized for data transfer only
//
// - Device type: EDGE_CG
//
// - Capacity: T42
//
// - Description: Snowball Edge Compute Optimized with GPU
//
// - Device type: EDGE_C
//
// - Capacity: T42
//
// - Description: Snowball Edge Compute Optimized without GPU
//
// - Device type: EDGE
//
// - Capacity: T100
//
// - Description: Snowball Edge Storage Optimized with EC2 Compute
//
// This device is replaced with T98.
//
// - Device type: STANDARD
//
// - Capacity: T50
//
// - Description: Original Snowball device
//
// This device is only available in the Ningxia, Beijing, and Singapore Amazon Web
//
// Services Region
//
// - Device type: STANDARD
//
// - Capacity: T80
//
// - Description: Original Snowball device
//
// This device is only available in the Ningxia, Beijing, and Singapore Amazon Web
//
// Services Region.
//
// - Snow Family device type: RACK_5U_C
//
// - Capacity: T13
//
// - Description: Snowblade.
//
// - Device type: V3_5S
//
// - Capacity: T240
//
// - Description: Snowball Edge Storage Optimized 210TB
//
// [Amazon Web Services Regional Services]: https://aws.amazon.com/about-aws/global-infrastructure/regional-product-services/?p=ngi&loc=4
