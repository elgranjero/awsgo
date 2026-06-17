package amp

// CreateScraper is generated as a reference stub.
// Executable command wiring lives under cmd/amp.go.
//
// The CreateScraper operation creates a scraper to collect metrics. A scraper
// pulls metrics from Prometheus-compatible sources and sends them to your Amazon
// Managed Service for Prometheus workspace. You can configure scrapers to collect
// metrics from Amazon EKS clusters, Amazon MSK clusters, or from VPC-based sources
// that support DNS-based service discovery. Scrapers are flexible, and can be
// configured to control what metrics are collected, the frequency of collection,
// what transformations are applied to the metrics, and more.
//
// An IAM role will be created for you that Amazon Managed Service for Prometheus
// uses to access the metrics in your source. You must configure this role with a
// policy that allows it to scrape metrics from your source. For Amazon EKS
// sources, see [Configuring your Amazon EKS cluster]in the Amazon Managed Service for Prometheus User Guide.
//
// The scrapeConfiguration parameter contains the base-64 encoded YAML
// configuration for the scraper.
//
// When creating a scraper, the service creates a Network Interface in each
// Availability Zone that are passed into CreateScraper through subnets. These
// network interfaces are used to connect to your source within the VPC for
// scraping metrics.
//
// For more information about collectors, including what metrics are collected,
// and how to configure the scraper, see [Using an Amazon Web Services managed collector]in the Amazon Managed Service for
// Prometheus User Guide.
//
// [Using an Amazon Web Services managed collector]: https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-collector-how-to.html
// [Configuring your Amazon EKS cluster]: https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-collector-how-to.html#AMP-collector-eks-setup
