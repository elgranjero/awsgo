package ssm

// ResetServiceSetting is generated as a reference stub.
// Executable command wiring lives under cmd/ssm.go.
//
// ServiceSetting is an account-level setting for an Amazon Web Services service.
// This setting defines how a user interacts with or uses a service or a feature of
// a service. For example, if an Amazon Web Services service charges money to the
// account based on feature or service usage, then the Amazon Web Services service
// team might create a default setting of "false". This means the user can't use
// this feature unless they change the setting to "true" and intentionally opt in
// for a paid feature.
//
// Services map a SettingId object to a setting value. Amazon Web Services
// services teams define the default value for a SettingId . You can't create a new
// SettingId , but you can overwrite the default value if you have the
// ssm:UpdateServiceSetting permission for the setting. Use the GetServiceSetting API operation to
// view the current value. Use the UpdateServiceSettingAPI operation to change the default setting.
//
// Reset the service setting for the account to the default value as provisioned
// by the Amazon Web Services service team.
