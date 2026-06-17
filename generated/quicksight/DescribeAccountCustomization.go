package quicksight

// DescribeAccountCustomization is generated as a reference stub.
// Executable command wiring lives under cmd/quicksight.go.
//
// Describes the customizations associated with the provided Amazon Web Services
// account and Amazon Quick Sight namespace. The Quick Sight console evaluates
// which customizations to apply by running this API operation with the Resolved
// flag included.
//
// To determine what customizations display when you run this command, it can help
// to visualize the relationship of the entities involved.
//
// - Amazon Web Services account - The Amazon Web Services account exists at the
// top of the hierarchy. It has the potential to use all of the Amazon Web Services
// Regions and Amazon Web Services Services. When you subscribe to Quick Sight, you
// choose one Amazon Web Services Region to use as your home Region. That's where
// your free SPICE capacity is located. You can use Quick Sight in any supported
// Amazon Web Services Region.
//
// - Amazon Web Services Region - You can sign in to Quick Sight in any Amazon
// Web Services Region. If you have a user directory, it resides in us-east-1,
// which is US East (N. Virginia). Generally speaking, these users have access to
// Quick Sight in any Amazon Web Services Region, unless they are constrained to a
// namespace.
//
// To run the command in a different Amazon Web Services Region, you change your
//
// Region settings. If you're using the CLI, you can use one of the following
// options:
//
// - Use [command line options].
//
// - Use [named profiles].
//
// - Run aws configure to change your default Amazon Web Services Region. Use
// Enter to key the same settings for your keys. For more information, see [Configuring the CLI].
//
// - Namespace - A Quick Sight namespace is a partition that contains users and
// assets (data sources, datasets, dashboards, and so on). To access assets that
// are in a specific namespace, users and groups must also be part of the same
// namespace. People who share a namespace are completely isolated from users and
// assets in other namespaces, even if they are in the same Amazon Web Services
// account and Amazon Web Services Region.
//
// - Applied customizations - Quick Sight customizations can apply to an Amazon
// Web Services account or to a namespace. Settings that you apply to a namespace
// override settings that you apply to an Amazon Web Services account.
//
// [named profiles]: https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-profiles.html
// [Configuring the CLI]: https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html
// [command line options]: https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-options.html
