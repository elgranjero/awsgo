package cloudhsm

// CreateHsm is generated as a reference stub.
// Executable command wiring lives under cmd/cloudhsm.go.
//
// This is documentation for AWS CloudHSM Classic. For more information, see [AWS CloudHSM Classic FAQs], the [AWS CloudHSM Classic User Guide]
// , and the [AWS CloudHSM Classic API Reference].
//
// For information about the current version of AWS CloudHSM, see [AWS CloudHSM], the [AWS CloudHSM User Guide], and the [AWS CloudHSM API Reference].
//
// Creates an uninitialized HSM instance.
//
// There is an upfront fee charged for each HSM instance that you create with the
// CreateHsm operation. If you accidentally provision an HSM and want to request a
// refund, delete the instance using the DeleteHsmoperation, go to the [AWS Support Center], create a new case,
// and select Account and Billing Support.
//
// It can take up to 20 minutes to create and provision an HSM. You can monitor
// the status of the HSM with the DescribeHsmoperation. The HSM is ready to be initialized
// when the status changes to RUNNING .
//
// Deprecated: This API is deprecated.
//
// [AWS CloudHSM User Guide]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/
// [AWS CloudHSM Classic FAQs]: http://aws.amazon.com/cloudhsm/faqs-classic/
// [AWS CloudHSM]: http://aws.amazon.com/cloudhsm/
// [AWS CloudHSM API Reference]: https://docs.aws.amazon.com/cloudhsm/latest/APIReference/
// [AWS CloudHSM Classic User Guide]: https://docs.aws.amazon.com/cloudhsm/classic/userguide/
// [AWS CloudHSM Classic API Reference]: https://docs.aws.amazon.com/cloudhsm/classic/APIReference/
// [AWS Support Center]: https://console.aws.amazon.com/support/home
