package workspaces

// ImportClientBranding is generated as a reference stub.
// Executable command wiring lives under cmd/workspaces.go.
//
// Imports client branding. Client branding allows you to customize your
// WorkSpace's client login portal. You can tailor your login portal company logo,
// the support email address, support link, link to reset password, and a custom
// message for users trying to sign in.
//
// After you import client branding, the default branding experience for the
// specified platform type is replaced with the imported experience
//
// - You must specify at least one platform type when importing client branding.
//
// - You can import up to 6 MB of data with each request. If your request
// exceeds this limit, you can import client branding for different platform types
// using separate requests.
//
// - In each platform type, the SupportEmail and SupportLink parameters are
// mutually exclusive. You can specify only one parameter for each platform type,
// but not both.
//
// - Imported data can take up to a minute to appear in the WorkSpaces client.
