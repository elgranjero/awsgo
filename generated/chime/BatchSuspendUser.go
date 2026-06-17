package chime

// BatchSuspendUser is generated as a reference stub.
// Executable command wiring lives under cmd/chime.go.
//
// Suspends up to 50 users from a Team or EnterpriseLWA Amazon Chime account. For
// more information about different account types, see [Managing Your Amazon Chime Accounts]in the Amazon Chime
// Administration Guide.
//
// Users suspended from a Team account are disassociated from the account,but they
// can continue to use Amazon Chime as free users. To remove the suspension from
// suspended Team account users, invite them to the Team account again. You can
// use the InviteUsersaction to do so.
//
// Users suspended from an EnterpriseLWA account are immediately signed out of
// Amazon Chime and can no longer sign in. To remove the suspension from suspended
// EnterpriseLWA account users, use the BatchUnsuspendUser action.
//
// To sign out users without suspending them, use the LogoutUser action.
//
// [Managing Your Amazon Chime Accounts]: https://docs.aws.amazon.com/chime/latest/ag/manage-chime-account.html
