package sagemaker

// CreateNotebookInstance is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Creates an SageMaker AI notebook instance. A notebook instance is a machine
// learning (ML) compute instance running on a Jupyter notebook.
//
// In a CreateNotebookInstance request, specify the type of ML compute instance
// that you want to run. SageMaker AI launches the instance, installs common
// libraries that you can use to explore datasets for model training, and attaches
// an ML storage volume to the notebook instance.
//
// SageMaker AI also provides a set of example notebooks. Each notebook
// demonstrates how to use SageMaker AI with a specific algorithm or with a machine
// learning framework.
//
// After receiving the request, SageMaker AI does the following:
//
// - Creates a network interface in the SageMaker AI VPC.
//
// - (Option) If you specified SubnetId , SageMaker AI creates a network
// interface in your own VPC, which is inferred from the subnet ID that you provide
// in the input. When creating this network interface, SageMaker AI attaches the
// security group that you specified in the request to the network interface that
// it creates in your VPC.
//
// - Launches an EC2 instance of the type specified in the request in the
// SageMaker AI VPC. If you specified SubnetId of your VPC, SageMaker AI
// specifies both network interfaces when launching this instance. This enables
// inbound traffic from your own VPC to the notebook instance, assuming that the
// security groups allow it.
//
// After creating the notebook instance, SageMaker AI returns its Amazon Resource
// Name (ARN). You can't change the name of a notebook instance after you create
// it.
//
// After SageMaker AI creates the notebook instance, you can connect to the
// Jupyter server and work in Jupyter notebooks. For example, you can write code to
// explore a dataset that you can use for model training, train a model, host
// models by creating SageMaker AI endpoints, and validate hosted models.
//
// For more information, see [How It Works].
//
// [How It Works]: https://docs.aws.amazon.com/sagemaker/latest/dg/how-it-works.html
