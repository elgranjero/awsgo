package rds

// ModifyCurrentDBClusterCapacity is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Set the capacity of an Aurora Serverless v1 DB cluster to a specific value.
//
// Aurora Serverless v1 scales seamlessly based on the workload on the DB cluster.
// In some cases, the capacity might not scale fast enough to meet a sudden change
// in workload, such as a large number of new transactions. Call
// ModifyCurrentDBClusterCapacity to set the capacity explicitly.
//
// After this call sets the DB cluster capacity, Aurora Serverless v1 can
// automatically scale the DB cluster based on the cooldown period for scaling up
// and the cooldown period for scaling down.
//
// For more information about Aurora Serverless v1, see [Using Amazon Aurora Serverless v1] in the Amazon Aurora User
// Guide.
//
// If you call ModifyCurrentDBClusterCapacity with the default TimeoutAction ,
// connections that prevent Aurora Serverless v1 from finding a scaling point might
// be dropped. For more information about scaling points, see [Autoscaling for Aurora Serverless v1]in the Amazon Aurora
// User Guide.
//
// This operation only applies to Aurora Serverless v1 DB clusters.
//
// [Autoscaling for Aurora Serverless v1]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.how-it-works.html#aurora-serverless.how-it-works.auto-scaling
// [Using Amazon Aurora Serverless v1]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.html
