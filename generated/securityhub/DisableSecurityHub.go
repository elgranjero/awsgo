package securityhub

// DisableSecurityHub is generated as a reference stub.
// Executable command wiring lives under cmd/securityhub.go.
//
// Disables Security Hub CSPM in your account only in the current Amazon Web
// Services Region. To disable Security Hub CSPM in all Regions, you must submit
// one request per Region where you have enabled Security Hub CSPM.
//
// You can't disable Security Hub CSPM in an account that is currently the
// Security Hub CSPM administrator.
//
// When you disable Security Hub CSPM, your existing findings and insights and any
// Security Hub CSPM configuration settings are deleted after 90 days and cannot be
// recovered. Any standards that were enabled are disabled, and your administrator
// and member account associations are removed.
//
// If you want to save your existing findings, you must export them before you
// disable Security Hub CSPM.
