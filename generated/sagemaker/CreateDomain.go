package sagemaker

// CreateDomain is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Creates a Domain . A domain consists of an associated Amazon Elastic File System
// volume, a list of authorized users, and a variety of security, application,
// policy, and Amazon Virtual Private Cloud (VPC) configurations. Users within a
// domain can share notebook files and other artifacts with each other.
//
// # EFS storage
//
// When a domain is created, an EFS volume is created for use by all of the users
// within the domain. Each user receives a private home directory within the EFS
// volume for notebooks, Git repositories, and data files.
//
// SageMaker AI uses the Amazon Web Services Key Management Service (Amazon Web
// Services KMS) to encrypt the EFS volume attached to the domain with an Amazon
// Web Services managed key by default. For more control, you can specify a
// customer managed key. For more information, see [Protect Data at Rest Using Encryption].
//
// # VPC configuration
//
// All traffic between the domain and the Amazon EFS volume is through the
// specified VPC and subnets. For other traffic, you can specify the
// AppNetworkAccessType parameter. AppNetworkAccessType corresponds to the network
// access type that you choose when you onboard to the domain. The following
// options are available:
//
// - PublicInternetOnly - Non-EFS traffic goes through a VPC managed by Amazon
// SageMaker AI, which allows internet access. This is the default value.
//
// - VpcOnly - All traffic is through the specified VPC and subnets. Internet
// access is disabled by default. To allow internet access, you must specify a NAT
// gateway.
//
// When internet access is disabled, you won't be able to run a Amazon SageMaker
//
// AI Studio notebook or to train or host models unless your VPC has an interface
// endpoint to the SageMaker AI API and runtime or a NAT gateway and your security
// groups allow outbound connections.
//
// NFS traffic over TCP on port 2049 needs to be allowed in both inbound and
// outbound rules in order to launch a Amazon SageMaker AI Studio app successfully.
//
// For more information, see [Connect Amazon SageMaker AI Studio Notebooks to Resources in a VPC].
//
// [Connect Amazon SageMaker AI Studio Notebooks to Resources in a VPC]: https://docs.aws.amazon.com/sagemaker/latest/dg/studio-notebooks-and-internet-access.html
// [Protect Data at Rest Using Encryption]: https://docs.aws.amazon.com/sagemaker/latest/dg/encryption-at-rest.html
