package securityhub

// EnableSecurityHub is generated as a reference stub.
// Executable command wiring lives under cmd/securityhub.go.
//
// Enables Security Hub CSPM for your account in the current Region or the Region
// you specify in the request.
//
// When you enable Security Hub CSPM, you grant to Security Hub CSPM the
// permissions necessary to gather findings from other services that are integrated
// with Security Hub CSPM.
//
// When you use the EnableSecurityHub operation to enable Security Hub CSPM, you
// also automatically enable the following standards:
//
// - Center for Internet Security (CIS) Amazon Web Services Foundations
// Benchmark v1.2.0
//
// - Amazon Web Services Foundational Security Best Practices
//
// Other standards are not automatically enabled.
//
// To opt out of automatically enabled standards, set EnableDefaultStandards to
// false .
//
// After you enable Security Hub CSPM, to enable a standard, use the
// BatchEnableStandards operation. To disable a standard, use the
// BatchDisableStandards operation.
//
// To learn more, see the [setup information] in the Security Hub CSPM User Guide.
//
// [setup information]: https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-settingup.html
