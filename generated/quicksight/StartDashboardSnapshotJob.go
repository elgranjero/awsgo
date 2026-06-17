package quicksight

// StartDashboardSnapshotJob is generated as a reference stub.
// Executable command wiring lives under cmd/quicksight.go.
//
// Starts an asynchronous job that generates a snapshot of a dashboard's output.
// You can request one or several of the following format configurations in each
// API call.
//
// - 1 PDF
//
// - 1 Excel workbook that includes up to 5 table or pivot table visuals
//
// - 5 CSVs from table or pivot table visuals
//
// Exporting CSV, Excel, or Pixel Perfect PDF reports requires Pixel Perfect
// Report Add-on.
//
// The status of a submitted job can be polled with the
// DescribeDashboardSnapshotJob API. When you call the DescribeDashboardSnapshotJob
// API, check the JobStatus field in the response. Once the job reaches a COMPLETED
// or FAILED status, use the DescribeDashboardSnapshotJobResult API to obtain the
// URLs for the generated files. If the job fails, the
// DescribeDashboardSnapshotJobResult API returns detailed information about the
// error that occurred.
//
// # StartDashboardSnapshotJob API throttling
//
// Quick Sight utilizes API throttling to create a more consistent user experience
// within a time span for customers when they call the StartDashboardSnapshotJob .
// By default, 12 jobs can run simlutaneously in one Amazon Web Services account
// and users can submit up 10 API requests per second before an account is
// throttled. If an overwhelming number of API requests are made by the same user
// in a short period of time, Quick Sight throttles the API calls to maintin an
// optimal experience and reliability for all Quick Sight users.
//
// # Common throttling scenarios
//
// The following list provides information about the most commin throttling
// scenarios that can occur.
//
// - A large number of SnapshotExport API jobs are running simultaneously on an
// Amazon Web Services account. When a new StartDashboardSnapshotJob is created
// and there are already 12 jobs with the RUNNING status, the new job request
// fails and returns a LimitExceededException error. Wait for a current job to
// comlpete before you resubmit the new job.
//
// - A large number of API requests are submitted on an Amazon Web Services
// account. When a user makes more than 10 API calls to the Quick Sight API in one
// second, a ThrottlingException is returned.
//
// If your use case requires a higher throttling limit, contact your account admin
// or [Amazon Web ServicesSupport]to explore options to tailor a more optimal expereince for your account.
//
// # Best practices to handle throttling
//
// If your use case projects high levels of API traffic, try to reduce the degree
// of frequency and parallelism of API calls as much as you can to avoid
// throttling. You can also perform a timing test to calculate an estimate for the
// total processing time of your projected load that stays within the throttling
// limits of the Quick Sight APIs. For example, if your projected traffic is 100
// snapshot jobs before 12:00 PM per day, start 12 jobs in parallel and measure the
// amount of time it takes to proccess all 12 jobs. Once you obtain the result,
// multiply the duration by 9, for example (12 minutes * 9 = 108 minutes) . Use the
// new result to determine the latest time at which the jobs need to be started to
// meet your target deadline.
//
// The time that it takes to process a job can be impacted by the following
// factors:
//
// - The dataset type (Direct Query or SPICE).
//
// - The size of the dataset.
//
// - The complexity of the calculated fields that are used in the dashboard.
//
// - The number of visuals that are on a sheet.
//
// - The types of visuals that are on the sheet.
//
// - The number of formats and snapshots that are requested in the job
// configuration.
//
// - The size of the generated snapshots.
//
// # Registered user support
//
// You can generate snapshots for registered Quick Sight users by using the
// Snapshot Job APIs with [identity-enhanced IAM role session credentials]. This approach allows you to create snapshots on behalf
// of specific Quick Sight users while respecting their row-level security (RLS),
// column-level security (CLS), dynamic default parameters and dashboard
// parameter/filter settings.
//
// To generate snapshots for registered Quick Sight users, you need to:
//
// - Obtain identity-enhanced IAM role session credentials from Amazon Web
// Services Security Token Service (STS).
//
// - Use these credentials to call the Snapshot Job APIs.
//
// Identity-enhanced credentials are credentials that contain information about
// the end user (e.g., registered Quick Sight user).
//
// If your Quick Sight users are backed by [Amazon Web Services Identity Center], then you need to set up a [trusted token issuer]. Then,
// getting identity-enhanced IAM credentials for a Quick Sight user will look like
// the following:
//
// - Authenticate user with your OIDC compliant Identity Provider. You should
// get auth tokens back.
//
// - Use the OIDC API, [CreateTokenWithIAM], to exchange auth tokens to IAM tokens. One of the
// resulted tokens will be identity token.
//
// - Call STS AssumeRole API as you normally would, but provide an extra
// ProvidedContexts parameter in the API request. The list of contexts must have
// a single trusted context assertion. The ProviderArn should be
// arn:aws:iam::aws:contextProvider/IdentityCenter while ContextAssertion will be
// the identity token you received in response from CreateTokenWithIAM
//
// For more details, see [IdC documentation on Identity-enhanced IAM role sessions].
//
// To obtain Identity-enhanced credentials for Quick Sight native users, IAM
// federated users, or Active Directory users, follow the steps below:
//
// - Call Quick Sight [GetIdentityContext API]to get identity token.
//
// - Call STS AssumeRole API as you normally would, but provide extra
// ProvidedContexts parameter in the API request. The list of contexts must have
// a single trusted context assertion. The ProviderArn should be
// arn:aws:iam::aws:contextProvider/QuickSight while ContextAssertion will be the
// identity token you received in response from GetIdentityContext
//
// After obtaining the identity-enhanced IAM role session credentials, you can use
// them to start a job, describe the job and describe job result. You can use the
// same credentials as long as they haven't expired. All API requests made with
// these credentials are considered to be made by the impersonated Quick Sight
// user.
//
// When using identity-enhanced session credentials, set the UserConfiguration
// request attribute to null. Otherwise, the request will be invalid.
//
// # Possible error scenarios
//
// The request fails with an Access Denied error in the following scenarios:
//
// - The credentials have expired.
//
// - The impersonated Quick Sight user doesn't have access to the specified
// dashboard.
//
// - The impersonated Quick Sight user is restricted from exporting data in the
// selected formats. For more information about export restrictions, see [Customizing access to Amazon Quick Sight capabilities].
//
// [GetIdentityContext API]: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_GetIdentityContext.html
// [identity-enhanced IAM role session credentials]: https://docs.aws.amazon.com/singlesignon/latest/userguide/trustedidentitypropagation-identity-enhanced-iam-role-sessions.html
// [CreateTokenWithIAM]: https://docs.aws.amazon.com/singlesignon/latest/OIDCAPIReference/API_CreateTokenWithIAM.html
// [IdC documentation on Identity-enhanced IAM role sessions]: https://docs.aws.amazon.com/singlesignon/latest/userguide/trustedidentitypropagation-identity-enhanced-iam-role-sessions.html
// [trusted token issuer]: https://docs.aws.amazon.com/singlesignon/latest/userguide/setuptrustedtokenissuer.html
// [Customizing access to Amazon Quick Sight capabilities]: https://docs.aws.amazon.com/quicksuite/latest/userguide/create-custom-permisions-profile.html
// [Amazon Web ServicesSupport]: http://aws.amazon.com/contact-us/
// [Amazon Web Services Identity Center]: https://docs.aws.amazon.com/singlesignon/latest/userguide/what-is.html
