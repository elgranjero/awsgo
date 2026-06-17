package internetmonitor

// CreateMonitor is generated as a reference stub.
// Executable command wiring lives under cmd/internetmonitor.go.
//
// Creates a monitor in Amazon CloudWatch Internet Monitor. A monitor is built
// based on information from the application resources that you add: VPCs, Network
// Load Balancers (NLBs), Amazon CloudFront distributions, and Amazon WorkSpaces
// directories. Internet Monitor then publishes internet measurements from Amazon
// Web Services that are specific to the city-networks. That is, the locations and
// ASNs (typically internet service providers or ISPs), where clients access your
// application. For more information, see [Using Amazon CloudWatch Internet Monitor]in the Amazon CloudWatch User Guide.
//
// When you create a monitor, you choose the percentage of traffic that you want
// to monitor. You can also set a maximum limit for the number of city-networks
// where client traffic is monitored, that caps the total traffic that Internet
// Monitor monitors. A city-network maximum is the limit of city-networks, but you
// only pay for the number of city-networks that are actually monitored. You can
// update your monitor at any time to change the percentage of traffic to monitor
// or the city-networks maximum. For more information, see [Choosing a city-network maximum value]in the Amazon
// CloudWatch User Guide.
//
// [Using Amazon CloudWatch Internet Monitor]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-InternetMonitor.html
// [Choosing a city-network maximum value]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/IMCityNetworksMaximum.html
