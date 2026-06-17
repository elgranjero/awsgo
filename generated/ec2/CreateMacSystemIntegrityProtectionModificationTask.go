package ec2

// CreateMacSystemIntegrityProtectionModificationTask is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Creates a System Integrity Protection (SIP) modification task to configure the
// SIP settings for an x86 Mac instance or Apple silicon Mac instance. For more
// information, see [Configure SIP for Amazon EC2 instances]in the Amazon EC2 User Guide.
//
// When you configure the SIP settings for your instance, you can either enable or
// disable all SIP settings, or you can specify a custom SIP configuration that
// selectively enables or disables specific SIP settings.
//
// If you implement a custom configuration, [connect to the instance and verify the settings] to ensure that your requirements are
// properly implemented and functioning as intended.
//
// SIP configurations might change with macOS updates. We recommend that you
// review custom SIP settings after any macOS version upgrade to ensure continued
// compatibility and proper functionality of your security configurations.
//
// To enable or disable all SIP settings, use the
// MacSystemIntegrityProtectionStatus parameter only. For example, to enable all
// SIP settings, specify the following:
//
// - MacSystemIntegrityProtectionStatus=enabled
//
// To specify a custom configuration that selectively enables or disables specific
// SIP settings, use the MacSystemIntegrityProtectionStatus parameter to enable or
// disable all SIP settings, and then use the
// MacSystemIntegrityProtectionConfiguration parameter to specify exceptions. In
// this case, the exceptions you specify for
// MacSystemIntegrityProtectionConfiguration override the value you specify for
// MacSystemIntegrityProtectionStatus. For example, to enable all SIP settings,
// except NvramProtections , specify the following:
//
// - MacSystemIntegrityProtectionStatus=enabled
//
// - MacSystemIntegrityProtectionConfigurationRequest "NvramProtections=disabled"
//
// [Configure SIP for Amazon EC2 instances]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/mac-sip-settings.html#mac-sip-configure
// [connect to the instance and verify the settings]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/mac-sip-settings.html#mac-sip-check-settings
