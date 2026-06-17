package databasemigrationservice

// BatchStartRecommendations is generated as a reference stub.
// Executable command wiring lives under cmd/databasemigrationservice.go.
//
// End of support notice: On May 20, 2026, Amazon Web Services will end support
//
// for Amazon Web Services DMS Fleet Advisor;. After May 20, 2026, you will no
// longer be able to access the Amazon Web Services DMS Fleet Advisor; console or
// Amazon Web Services DMS Fleet Advisor; resources. For more information, see [Amazon Web Services DMS Fleet Advisor end of support].
//
// Starts the analysis of up to 20 source databases to recommend target engines
// for each source database. This is a batch version of [StartRecommendations].
//
// The result of analysis of each source database is reported individually in the
// response. Because the batch request can result in a combination of successful
// and unsuccessful actions, you should check for batch errors even when the call
// returns an HTTP status code of 200 .
//
// [Amazon Web Services DMS Fleet Advisor end of support]: https://docs.aws.amazon.com/dms/latest/userguide/dms_fleet.advisor-end-of-support.html
// [StartRecommendations]: https://docs.aws.amazon.com/dms/latest/APIReference/API_StartRecommendations.html
