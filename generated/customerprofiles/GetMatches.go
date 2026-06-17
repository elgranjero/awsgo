package customerprofiles

// GetMatches is generated as a reference stub.
// Executable command wiring lives under cmd/customerprofiles.go.
//
// Before calling this API, use [CreateDomain] or [UpdateDomain] to enable identity resolution: set Matching
// to true.
//
// GetMatches returns potentially matching profiles, based on the results of the
// latest run of a machine learning process.
//
// The process of matching duplicate profiles. If Matching = true , Amazon Connect
// Customer Profiles starts a weekly batch process called Identity Resolution Job.
// If you do not specify a date and time for Identity Resolution Job to run, by
// default it runs every Saturday at 12AM UTC to detect duplicate profiles in your
// domains.
//
// After the Identity Resolution Job completes, use the [GetMatches] API to return and review
// the results. Or, if you have configured ExportingConfig in the MatchingRequest ,
// you can download the results from S3.
//
// Amazon Connect uses the following profile attributes to identify matches:
//
// - PhoneNumber
//
// - HomePhoneNumber
//
// - BusinessPhoneNumber
//
// - MobilePhoneNumber
//
// - EmailAddress
//
// - PersonalEmailAddress
//
// - BusinessEmailAddress
//
// - FullName
//
// For example, two or more profiles—with spelling mistakes such as John Doe and
// Jhn Doe, or different casing email addresses such as JOHN_DOE(at)ANYCOMPANY.COM and
// johndoe(at)anycompany.com, or different phone number formats such as 555-010-0000
// and +1-555-010-0000—can be detected as belonging to the same customer John Doe
// and merged into a unified profile.
//
// [GetMatches]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_GetMatches.html
// [UpdateDomain]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_UpdateDomain.html
// [CreateDomain]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_CreateDomain.html
