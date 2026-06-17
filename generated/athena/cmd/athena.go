package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// athenaCmd represents the athena command
var _athenaCmd = &cobra.Command{
	Use:   "athena",
	Short: "AWS athena CLI",
	Run: func(cmd *cobra.Command, args []string) {
		_awsOutput = resolveAWSOutput(_awsProfile, cmd.Flags().Changed("output"))
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := athena.NewFromConfig(cfg)
		if _athenaBatchGetNamedQuery {
			athena_BatchGetNamedQuery(cfg, client)
			return
		}
		if _athenaBatchGetPreparedStatement {
			athena_BatchGetPreparedStatement(cfg, client)
			return
		}
		if _athenaBatchGetQueryExecution {
			athena_BatchGetQueryExecution(cfg, client)
			return
		}
		if _athenaCancelCapacityReservation {
			athena_CancelCapacityReservation(cfg, client)
			return
		}
		if _athenaCreateCapacityReservation {
			athena_CreateCapacityReservation(cfg, client)
			return
		}
		if _athenaCreateDataCatalog {
			athena_CreateDataCatalog(cfg, client)
			return
		}
		if _athenaCreateNamedQuery {
			athena_CreateNamedQuery(cfg, client)
			return
		}
		if _athenaCreateNotebook {
			athena_CreateNotebook(cfg, client)
			return
		}
		if _athenaCreatePreparedStatement {
			athena_CreatePreparedStatement(cfg, client)
			return
		}
		if _athenaCreatePresignedNotebookUrl {
			athena_CreatePresignedNotebookUrl(cfg, client)
			return
		}
		if _athenaCreateWorkGroup {
			athena_CreateWorkGroup(cfg, client)
			return
		}
		if _athenaDeleteCapacityReservation {
			athena_DeleteCapacityReservation(cfg, client)
			return
		}
		if _athenaDeleteDataCatalog {
			athena_DeleteDataCatalog(cfg, client)
			return
		}
		if _athenaDeleteNamedQuery {
			athena_DeleteNamedQuery(cfg, client)
			return
		}
		if _athenaDeleteNotebook {
			athena_DeleteNotebook(cfg, client)
			return
		}
		if _athenaDeletePreparedStatement {
			athena_DeletePreparedStatement(cfg, client)
			return
		}
		if _athenaDeleteWorkGroup {
			athena_DeleteWorkGroup(cfg, client)
			return
		}
		if _athenaExportNotebook {
			athena_ExportNotebook(cfg, client)
			return
		}
		if _athenaGetCalculationExecution {
			athena_GetCalculationExecution(cfg, client)
			return
		}
		if _athenaGetCalculationExecutionCode {
			athena_GetCalculationExecutionCode(cfg, client)
			return
		}
		if _athenaGetCalculationExecutionStatus {
			athena_GetCalculationExecutionStatus(cfg, client)
			return
		}
		if _athenaGetCapacityAssignmentConfiguration {
			athena_GetCapacityAssignmentConfiguration(cfg, client)
			return
		}
		if _athenaGetCapacityReservation {
			athena_GetCapacityReservation(cfg, client)
			return
		}
		if _athenaGetDataCatalog {
			athena_GetDataCatalog(cfg, client)
			return
		}
		if _athenaGetDatabase {
			athena_GetDatabase(cfg, client)
			return
		}
		if _athenaGetNamedQuery {
			athena_GetNamedQuery(cfg, client)
			return
		}
		if _athenaGetNotebookMetadata {
			athena_GetNotebookMetadata(cfg, client)
			return
		}
		if _athenaGetPreparedStatement {
			athena_GetPreparedStatement(cfg, client)
			return
		}
		if _athenaGetQueryExecution {
			athena_GetQueryExecution(cfg, client)
			return
		}
		if _athenaGetQueryResults {
			athena_GetQueryResults(cfg, client)
			return
		}
		if _athenaGetQueryRuntimeStatistics {
			athena_GetQueryRuntimeStatistics(cfg, client)
			return
		}
		if _athenaGetResourceDashboard {
			athena_GetResourceDashboard(cfg, client)
			return
		}
		if _athenaGetSession {
			athena_GetSession(cfg, client)
			return
		}
		if _athenaGetSessionEndpoint {
			athena_GetSessionEndpoint(cfg, client)
			return
		}
		if _athenaGetSessionStatus {
			athena_GetSessionStatus(cfg, client)
			return
		}
		if _athenaGetTableMetadata {
			athena_GetTableMetadata(cfg, client)
			return
		}
		if _athenaGetWorkGroup {
			athena_GetWorkGroup(cfg, client)
			return
		}
		if _athenaImportNotebook {
			athena_ImportNotebook(cfg, client)
			return
		}
		if _athenaListApplicationDPUSizes {
			athena_ListApplicationDPUSizes(cfg, client)
			return
		}
		if _athenaListCalculationExecutions {
			athena_ListCalculationExecutions(cfg, client)
			return
		}
		if _athenaListCapacityReservations {
			athena_ListCapacityReservations(cfg, client)
			return
		}
		if _athenaListDataCatalogs {
			athena_ListDataCatalogs(cfg, client)
			return
		}
		if _athenaListDatabases {
			athena_ListDatabases(cfg, client)
			return
		}
		if _athenaListEngineVersions {
			athena_ListEngineVersions(cfg, client)
			return
		}
		if _athenaListExecutors {
			athena_ListExecutors(cfg, client)
			return
		}
		if _athenaListNamedQueries {
			athena_ListNamedQueries(cfg, client)
			return
		}
		if _athenaListNotebookMetadata {
			athena_ListNotebookMetadata(cfg, client)
			return
		}
		if _athenaListNotebookSessions {
			athena_ListNotebookSessions(cfg, client)
			return
		}
		if _athenaListPreparedStatements {
			athena_ListPreparedStatements(cfg, client)
			return
		}
		if _athenaListQueryExecutions {
			athena_ListQueryExecutions(cfg, client)
			return
		}
		if _athenaListSessions {
			athena_ListSessions(cfg, client)
			return
		}
		if _athenaListTableMetadata {
			athena_ListTableMetadata(cfg, client)
			return
		}
		if _athenaListTagsForResource {
			athena_ListTagsForResource(cfg, client)
			return
		}
		if _athenaListWorkGroups {
			athena_ListWorkGroups(cfg, client)
			return
		}
		if _athenaPutCapacityAssignmentConfiguration {
			athena_PutCapacityAssignmentConfiguration(cfg, client)
			return
		}
		if _athenaStartCalculationExecution {
			athena_StartCalculationExecution(cfg, client)
			return
		}
		if _athenaStartQueryExecution {
			athena_StartQueryExecution(cfg, client)
			return
		}
		if _athenaStartSession {
			athena_StartSession(cfg, client)
			return
		}
		if _athenaStopCalculationExecution {
			athena_StopCalculationExecution(cfg, client)
			return
		}
		if _athenaStopQueryExecution {
			athena_StopQueryExecution(cfg, client)
			return
		}
		if _athenaTagResource {
			athena_TagResource(cfg, client)
			return
		}
		if _athenaTerminateSession {
			athena_TerminateSession(cfg, client)
			return
		}
		if _athenaUntagResource {
			athena_UntagResource(cfg, client)
			return
		}
		if _athenaUpdateCapacityReservation {
			athena_UpdateCapacityReservation(cfg, client)
			return
		}
		if _athenaUpdateDataCatalog {
			athena_UpdateDataCatalog(cfg, client)
			return
		}
		if _athenaUpdateNamedQuery {
			athena_UpdateNamedQuery(cfg, client)
			return
		}
		if _athenaUpdateNotebook {
			athena_UpdateNotebook(cfg, client)
			return
		}
		if _athenaUpdateNotebookMetadata {
			athena_UpdateNotebookMetadata(cfg, client)
			return
		}
		if _athenaUpdatePreparedStatement {
			athena_UpdatePreparedStatement(cfg, client)
			return
		}
		if _athenaUpdateWorkGroup {
			athena_UpdateWorkGroup(cfg, client)
			return
		}

	},
}

var (
	_athenaBatchGetNamedQuery                 bool
	_athenaBatchGetPreparedStatement          bool
	_athenaBatchGetQueryExecution             bool
	_athenaCancelCapacityReservation          bool
	_athenaCreateCapacityReservation          bool
	_athenaCreateDataCatalog                  bool
	_athenaCreateNamedQuery                   bool
	_athenaCreateNotebook                     bool
	_athenaCreatePreparedStatement            bool
	_athenaCreatePresignedNotebookUrl         bool
	_athenaCreateWorkGroup                    bool
	_athenaDeleteCapacityReservation          bool
	_athenaDeleteDataCatalog                  bool
	_athenaDeleteNamedQuery                   bool
	_athenaDeleteNotebook                     bool
	_athenaDeletePreparedStatement            bool
	_athenaDeleteWorkGroup                    bool
	_athenaExportNotebook                     bool
	_athenaGetCalculationExecution            bool
	_athenaGetCalculationExecutionCode        bool
	_athenaGetCalculationExecutionStatus      bool
	_athenaGetCapacityAssignmentConfiguration bool
	_athenaGetCapacityReservation             bool
	_athenaGetDataCatalog                     bool
	_athenaGetDatabase                        bool
	_athenaGetNamedQuery                      bool
	_athenaGetNotebookMetadata                bool
	_athenaGetPreparedStatement               bool
	_athenaGetQueryExecution                  bool
	_athenaGetQueryResults                    bool
	_athenaGetQueryRuntimeStatistics          bool
	_athenaGetResourceDashboard               bool
	_athenaGetSession                         bool
	_athenaGetSessionEndpoint                 bool
	_athenaGetSessionStatus                   bool
	_athenaGetTableMetadata                   bool
	_athenaGetWorkGroup                       bool
	_athenaImportNotebook                     bool
	_athenaListApplicationDPUSizes            bool
	_athenaListCalculationExecutions          bool
	_athenaListCapacityReservations           bool
	_athenaListDataCatalogs                   bool
	_athenaListDatabases                      bool
	_athenaListEngineVersions                 bool
	_athenaListExecutors                      bool
	_athenaListNamedQueries                   bool
	_athenaListNotebookMetadata               bool
	_athenaListNotebookSessions               bool
	_athenaListPreparedStatements             bool
	_athenaListQueryExecutions                bool
	_athenaListSessions                       bool
	_athenaListTableMetadata                  bool
	_athenaListTagsForResource                bool
	_athenaListWorkGroups                     bool
	_athenaPutCapacityAssignmentConfiguration bool
	_athenaStartCalculationExecution          bool
	_athenaStartQueryExecution                bool
	_athenaStartSession                       bool
	_athenaStopCalculationExecution           bool
	_athenaStopQueryExecution                 bool
	_athenaTagResource                        bool
	_athenaTerminateSession                   bool
	_athenaUntagResource                      bool
	_athenaUpdateCapacityReservation          bool
	_athenaUpdateDataCatalog                  bool
	_athenaUpdateNamedQuery                   bool
	_athenaUpdateNotebook                     bool
	_athenaUpdateNotebookMetadata             bool
	_athenaUpdatePreparedStatement            bool
	_athenaUpdateWorkGroup                    bool

	_athenaCalculationConfiguration    string
	_athenaCalculationExecutionId      string
	_athenaCapacityAssignments         string
	_athenaCapacityReservationName     string
	_athenaCatalogName                 string
	_athenaClientRequestToken          string
	_athenaCodeBlock                   string
	_athenaConfiguration               string
	_athenaConfigurationUpdates        string
	_athenaCopyWorkGroupTags           string
	_athenaDatabase                    string
	_athenaDatabaseName                string
	_athenaDeleteCatalogOnly           string
	_athenaDescription                 string
	_athenaEngineConfiguration         string
	_athenaExecutionParameters         []string
	_athenaExecutionRole               string
	_athenaExecutorStateFilter         string
	_athenaExpression                  string
	_athenaFilters                     string
	_athenaMaxResults                  string
	_athenaMonitoringConfiguration     string
	_athenaName                        string
	_athenaNamedQueryId                string
	_athenaNamedQueryIds               []string
	_athenaNextToken                   string
	_athenaNotebookId                  string
	_athenaNotebookS3LocationUri       string
	_athenaNotebookVersion             string
	_athenaParameters                  string
	_athenaPayload                     string
	_athenaPreparedStatementNames      []string
	_athenaQueryExecutionContext       string
	_athenaQueryExecutionId            string
	_athenaQueryExecutionIds           []string
	_athenaQueryResultType             string
	_athenaQueryStatement              string
	_athenaQueryString                 string
	_athenaRecursiveDeleteOption       string
	_athenaResourceARN                 string
	_athenaResultConfiguration         string
	_athenaResultReuseConfiguration    string
	_athenaSessionId                   string
	_athenaSessionIdleTimeoutInMinutes string
	_athenaState                       string
	_athenaStateFilter                 string
	_athenaStatementName               string
	_athenaTableName                   string
	_athenaTagKeys                     []string
	_athenaTags                        string
	_athenaTargetDpus                  string
	_athenaType                        string
	_athenaWorkGroup                   string
)

// Returns the details of a single named query or a list of up to 50 queries,
// which you provide as an array of query ID strings. Requires you to have access
// to the workgroup in which the queries were saved. Use ListNamedQueriesInputto get the list of named
// query IDs in the specified workgroup. If information could not be retrieved for
// a submitted query ID, information about the query ID submitted is listed under UnprocessedNamedQueryId
// . Named queries differ from executed queries. Use BatchGetQueryExecutionInputto get details about each
// unique query execution, and ListQueryExecutionsInputto get a list of query execution IDs.
func athena_BatchGetNamedQuery(cfg aws.Config, client *athena.Client) {
	input := &athena.BatchGetNamedQueryInput{
		// NamedQueryIds: []string, // Required
	}

	if len(_athenaNamedQueryIds) > 0 {
		input.NamedQueryIds = append([]string(nil), _athenaNamedQueryIds...)
	}

	if resp, err := client.BatchGetNamedQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the details of a single prepared statement or a list of up to 256
// prepared statements for the array of prepared statement names that you provide.
// Requires you to have access to the workgroup to which the prepared statements
// belong. If a prepared statement cannot be retrieved for the name specified, the
// statement is listed in UnprocessedPreparedStatementNames .
func athena_BatchGetPreparedStatement(cfg aws.Config, client *athena.Client) {
	input := &athena.BatchGetPreparedStatementInput{
		// PreparedStatementNames: []string, // Required
		// WorkGroup: *string, // Required
	}

	if len(_athenaPreparedStatementNames) > 0 {
		input.PreparedStatementNames = append([]string(nil), _athenaPreparedStatementNames...)
	}
	if len(_athenaWorkGroup) > 0 {
		input.WorkGroup = aws.String(_athenaWorkGroup)
	}

	if resp, err := client.BatchGetPreparedStatement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the details of a single query execution or a list of up to 50 query
// executions, which you provide as an array of query execution ID strings.
// Requires you to have access to the workgroup in which the queries ran. To get a
// list of query execution IDs, use ListQueryExecutionsInput$WorkGroup. Query executions differ from named (saved)
// queries. Use BatchGetNamedQueryInputto get details about named queries.
func athena_BatchGetQueryExecution(cfg aws.Config, client *athena.Client) {
	input := &athena.BatchGetQueryExecutionInput{
		// QueryExecutionIds: []string, // Required
	}

	if len(_athenaQueryExecutionIds) > 0 {
		input.QueryExecutionIds = append([]string(nil), _athenaQueryExecutionIds...)
	}

	if resp, err := client.BatchGetQueryExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels the capacity reservation with the specified name. Cancelled
// reservations remain in your account and will be deleted 45 days after
// cancellation. During the 45 days, you cannot re-purpose or reuse a reservation
// that has been cancelled, but you can refer to its tags and view it for
// historical reference.
func athena_CancelCapacityReservation(cfg aws.Config, client *athena.Client) {
	input := &athena.CancelCapacityReservationInput{
		// Name: *string, // Required
	}

	if len(_athenaName) > 0 {
		input.Name = aws.String(_athenaName)
	}

	if resp, err := client.CancelCapacityReservation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a capacity reservation with the specified name and number of requested
// data processing units.
func athena_CreateCapacityReservation(cfg aws.Config, client *athena.Client) {
	input := &athena.CreateCapacityReservationInput{
		// Name: *string, // Required
		// TargetDpus: *int32, // Required
	}

	if len(_athenaName) > 0 {
		input.Name = aws.String(_athenaName)
	}
	if len(_athenaTargetDpus) > 0 {
		if err := assignInputField(input, "TargetDpus", _athenaTargetDpus); err != nil {
			log.Errorf("invalid --target-dpus: %s", err.Error())
			return
		}
	}
	if len(_athenaTags) > 0 {
		if err := assignInputField(input, "Tags", _athenaTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCapacityReservation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates (registers) a data catalog with the specified name and properties.
// Catalogs created are visible to all users of the same Amazon Web Services
// account.
//
// For a FEDERATED catalog, this API operation creates the following resources.
//
// - CFN Stack Name with a maximum length of 128 characters and prefix
// athenafederatedcatalog-CATALOG_NAME_SANITIZED with length 23 characters.
//
// - Lambda Function Name with a maximum length of 64 characters and prefix
// athenafederatedcatalog_CATALOG_NAME_SANITIZED with length 23 characters.
//
// - Glue Connection Name with a maximum length of 255 characters and a prefix
// athenafederatedcatalog_CATALOG_NAME_SANITIZED with length 23 characters.
func athena_CreateDataCatalog(cfg aws.Config, client *athena.Client) {
	input := &athena.CreateDataCatalogInput{
		// Name: *string, // Required
		// Type: types.DataCatalogType, // Required
	}

	if len(_athenaName) > 0 {
		input.Name = aws.String(_athenaName)
	}
	if len(_athenaType) > 0 {
		if err := assignInputField(input, "Type", _athenaType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_athenaDescription) > 0 {
		input.Description = aws.String(_athenaDescription)
	}
	if len(_athenaParameters) > 0 {
		if err := assignInputField(input, "Parameters", _athenaParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_athenaTags) > 0 {
		if err := assignInputField(input, "Tags", _athenaTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataCatalog(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a named query in the specified workgroup. Requires that you have access
// to the workgroup.
func athena_CreateNamedQuery(cfg aws.Config, client *athena.Client) {
	input := &athena.CreateNamedQueryInput{
		// Database: *string, // Required
		// Name: *string, // Required
		// QueryString: *string, // Required
	}

	if len(_athenaDatabase) > 0 {
		input.Database = aws.String(_athenaDatabase)
	}
	if len(_athenaName) > 0 {
		input.Name = aws.String(_athenaName)
	}
	if len(_athenaQueryString) > 0 {
		input.QueryString = aws.String(_athenaQueryString)
	}
	if len(_athenaClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_athenaClientRequestToken)
	}
	if len(_athenaDescription) > 0 {
		input.Description = aws.String(_athenaDescription)
	}
	if len(_athenaWorkGroup) > 0 {
		input.WorkGroup = aws.String(_athenaWorkGroup)
	}

	if resp, err := client.CreateNamedQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an empty ipynb file in the specified Apache Spark enabled workgroup.
// Throws an error if a file in the workgroup with the same name already exists.
func athena_CreateNotebook(cfg aws.Config, client *athena.Client) {
	input := &athena.CreateNotebookInput{
		// Name: *string, // Required
		// WorkGroup: *string, // Required
	}

	if len(_athenaName) > 0 {
		input.Name = aws.String(_athenaName)
	}
	if len(_athenaWorkGroup) > 0 {
		input.WorkGroup = aws.String(_athenaWorkGroup)
	}
	if len(_athenaClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_athenaClientRequestToken)
	}

	if resp, err := client.CreateNotebook(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a prepared statement for use with SQL queries in Athena.
func athena_CreatePreparedStatement(cfg aws.Config, client *athena.Client) {
	input := &athena.CreatePreparedStatementInput{
		// QueryStatement: *string, // Required
		// StatementName: *string, // Required
		// WorkGroup: *string, // Required
	}

	if len(_athenaQueryStatement) > 0 {
		input.QueryStatement = aws.String(_athenaQueryStatement)
	}
	if len(_athenaStatementName) > 0 {
		input.StatementName = aws.String(_athenaStatementName)
	}
	if len(_athenaWorkGroup) > 0 {
		input.WorkGroup = aws.String(_athenaWorkGroup)
	}
	if len(_athenaDescription) > 0 {
		input.Description = aws.String(_athenaDescription)
	}

	if resp, err := client.CreatePreparedStatement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an authentication token and the URL at which the notebook can be accessed.
// During programmatic access, CreatePresignedNotebookUrl must be called every 10
// minutes to refresh the authentication token. For information about granting
// programmatic access, see [Grant programmatic access].
//
// [Grant programmatic access]: https://docs.aws.amazon.com/athena/latest/ug/setting-up.html#setting-up-grant-programmatic-access
func athena_CreatePresignedNotebookUrl(cfg aws.Config, client *athena.Client) {
	input := &athena.CreatePresignedNotebookUrlInput{
		// SessionId: *string, // Required
	}

	if len(_athenaSessionId) > 0 {
		input.SessionId = aws.String(_athenaSessionId)
	}

	if resp, err := client.CreatePresignedNotebookUrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a workgroup with the specified name. A workgroup can be an Apache Spark
// enabled workgroup or an Athena SQL workgroup.
func athena_CreateWorkGroup(cfg aws.Config, client *athena.Client) {
	input := &athena.CreateWorkGroupInput{
		// Name: *string, // Required
	}

	if len(_athenaName) > 0 {
		input.Name = aws.String(_athenaName)
	}
	if len(_athenaConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _athenaConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_athenaDescription) > 0 {
		input.Description = aws.String(_athenaDescription)
	}
	if len(_athenaTags) > 0 {
		if err := assignInputField(input, "Tags", _athenaTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWorkGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a cancelled capacity reservation. A reservation must be cancelled
// before it can be deleted. A deleted reservation is immediately removed from your
// account and can no longer be referenced, including by its ARN. A deleted
// reservation cannot be called by GetCapacityReservation , and deleted
// reservations do not appear in the output of ListCapacityReservations .
func athena_DeleteCapacityReservation(cfg aws.Config, client *athena.Client) {
	input := &athena.DeleteCapacityReservationInput{
		// Name: *string, // Required
	}

	if len(_athenaName) > 0 {
		input.Name = aws.String(_athenaName)
	}

	if resp, err := client.DeleteCapacityReservation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a data catalog.
func athena_DeleteDataCatalog(cfg aws.Config, client *athena.Client) {
	input := &athena.DeleteDataCatalogInput{
		// Name: *string, // Required
	}

	if len(_athenaName) > 0 {
		input.Name = aws.String(_athenaName)
	}
	if len(_athenaDeleteCatalogOnly) > 0 {
		if err := assignInputField(input, "DeleteCatalogOnly", _athenaDeleteCatalogOnly); err != nil {
			log.Errorf("invalid --delete-catalog-only: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteDataCatalog(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the named query if you have access to the workgroup in which the query
// was saved.
func athena_DeleteNamedQuery(cfg aws.Config, client *athena.Client) {
	input := &athena.DeleteNamedQueryInput{
		// NamedQueryId: *string, // Required
	}

	if len(_athenaNamedQueryId) > 0 {
		input.NamedQueryId = aws.String(_athenaNamedQueryId)
	}

	if resp, err := client.DeleteNamedQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified notebook.
func athena_DeleteNotebook(cfg aws.Config, client *athena.Client) {
	input := &athena.DeleteNotebookInput{
		// NotebookId: *string, // Required
	}

	if len(_athenaNotebookId) > 0 {
		input.NotebookId = aws.String(_athenaNotebookId)
	}

	if resp, err := client.DeleteNotebook(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the prepared statement with the specified name from the specified
// workgroup.
func athena_DeletePreparedStatement(cfg aws.Config, client *athena.Client) {
	input := &athena.DeletePreparedStatementInput{
		// StatementName: *string, // Required
		// WorkGroup: *string, // Required
	}

	if len(_athenaStatementName) > 0 {
		input.StatementName = aws.String(_athenaStatementName)
	}
	if len(_athenaWorkGroup) > 0 {
		input.WorkGroup = aws.String(_athenaWorkGroup)
	}

	if resp, err := client.DeletePreparedStatement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the workgroup with the specified name. The primary workgroup cannot be
// deleted.
func athena_DeleteWorkGroup(cfg aws.Config, client *athena.Client) {
	input := &athena.DeleteWorkGroupInput{
		// WorkGroup: *string, // Required
	}

	if len(_athenaWorkGroup) > 0 {
		input.WorkGroup = aws.String(_athenaWorkGroup)
	}
	if len(_athenaRecursiveDeleteOption) > 0 {
		if err := assignInputField(input, "RecursiveDeleteOption", _athenaRecursiveDeleteOption); err != nil {
			log.Errorf("invalid --recursive-delete-option: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteWorkGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Exports the specified notebook and its metadata.
func athena_ExportNotebook(cfg aws.Config, client *athena.Client) {
	input := &athena.ExportNotebookInput{
		// NotebookId: *string, // Required
	}

	if len(_athenaNotebookId) > 0 {
		input.NotebookId = aws.String(_athenaNotebookId)
	}

	if resp, err := client.ExportNotebook(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a previously submitted calculation execution.
func athena_GetCalculationExecution(cfg aws.Config, client *athena.Client) {
	input := &athena.GetCalculationExecutionInput{
		// CalculationExecutionId: *string, // Required
	}

	if len(_athenaCalculationExecutionId) > 0 {
		input.CalculationExecutionId = aws.String(_athenaCalculationExecutionId)
	}

	if resp, err := client.GetCalculationExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the unencrypted code that was executed for the calculation.
func athena_GetCalculationExecutionCode(cfg aws.Config, client *athena.Client) {
	input := &athena.GetCalculationExecutionCodeInput{
		// CalculationExecutionId: *string, // Required
	}

	if len(_athenaCalculationExecutionId) > 0 {
		input.CalculationExecutionId = aws.String(_athenaCalculationExecutionId)
	}

	if resp, err := client.GetCalculationExecutionCode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the status of a current calculation.
func athena_GetCalculationExecutionStatus(cfg aws.Config, client *athena.Client) {
	input := &athena.GetCalculationExecutionStatusInput{
		// CalculationExecutionId: *string, // Required
	}

	if len(_athenaCalculationExecutionId) > 0 {
		input.CalculationExecutionId = aws.String(_athenaCalculationExecutionId)
	}

	if resp, err := client.GetCalculationExecutionStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the capacity assignment configuration for a capacity reservation, if one
// exists.
func athena_GetCapacityAssignmentConfiguration(cfg aws.Config, client *athena.Client) {
	input := &athena.GetCapacityAssignmentConfigurationInput{
		// CapacityReservationName: *string, // Required
	}

	if len(_athenaCapacityReservationName) > 0 {
		input.CapacityReservationName = aws.String(_athenaCapacityReservationName)
	}

	if resp, err := client.GetCapacityAssignmentConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the capacity reservation with the specified name.
func athena_GetCapacityReservation(cfg aws.Config, client *athena.Client) {
	input := &athena.GetCapacityReservationInput{
		// Name: *string, // Required
	}

	if len(_athenaName) > 0 {
		input.Name = aws.String(_athenaName)
	}

	if resp, err := client.GetCapacityReservation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the specified data catalog.
func athena_GetDataCatalog(cfg aws.Config, client *athena.Client) {
	input := &athena.GetDataCatalogInput{
		// Name: *string, // Required
	}

	if len(_athenaName) > 0 {
		input.Name = aws.String(_athenaName)
	}
	if len(_athenaWorkGroup) > 0 {
		input.WorkGroup = aws.String(_athenaWorkGroup)
	}

	if resp, err := client.GetDataCatalog(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a database object for the specified database and data catalog.
func athena_GetDatabase(cfg aws.Config, client *athena.Client) {
	input := &athena.GetDatabaseInput{
		// CatalogName: *string, // Required
		// DatabaseName: *string, // Required
	}

	if len(_athenaCatalogName) > 0 {
		input.CatalogName = aws.String(_athenaCatalogName)
	}
	if len(_athenaDatabaseName) > 0 {
		input.DatabaseName = aws.String(_athenaDatabaseName)
	}
	if len(_athenaWorkGroup) > 0 {
		input.WorkGroup = aws.String(_athenaWorkGroup)
	}

	if resp, err := client.GetDatabase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a single query. Requires that you have access to the
// workgroup in which the query was saved.
func athena_GetNamedQuery(cfg aws.Config, client *athena.Client) {
	input := &athena.GetNamedQueryInput{
		// NamedQueryId: *string, // Required
	}

	if len(_athenaNamedQueryId) > 0 {
		input.NamedQueryId = aws.String(_athenaNamedQueryId)
	}

	if resp, err := client.GetNamedQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves notebook metadata for the specified notebook ID.
func athena_GetNotebookMetadata(cfg aws.Config, client *athena.Client) {
	input := &athena.GetNotebookMetadataInput{
		// NotebookId: *string, // Required
	}

	if len(_athenaNotebookId) > 0 {
		input.NotebookId = aws.String(_athenaNotebookId)
	}

	if resp, err := client.GetNotebookMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the prepared statement with the specified name from the specified
// workgroup.
func athena_GetPreparedStatement(cfg aws.Config, client *athena.Client) {
	input := &athena.GetPreparedStatementInput{
		// StatementName: *string, // Required
		// WorkGroup: *string, // Required
	}

	if len(_athenaStatementName) > 0 {
		input.StatementName = aws.String(_athenaStatementName)
	}
	if len(_athenaWorkGroup) > 0 {
		input.WorkGroup = aws.String(_athenaWorkGroup)
	}

	if resp, err := client.GetPreparedStatement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a single execution of a query if you have access to
// the workgroup in which the query ran. Each time a query executes, information
// about the query execution is saved with a unique ID.
func athena_GetQueryExecution(cfg aws.Config, client *athena.Client) {
	input := &athena.GetQueryExecutionInput{
		// QueryExecutionId: *string, // Required
	}

	if len(_athenaQueryExecutionId) > 0 {
		input.QueryExecutionId = aws.String(_athenaQueryExecutionId)
	}

	if resp, err := client.GetQueryExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Streams the results of a single query execution specified by QueryExecutionId
// from the Athena query results location in Amazon S3. For more information, see [Working with query results, recent queries, and output files]
// in the Amazon Athena User Guide. This request does not execute the query but
// returns results. Use StartQueryExecutionto run a query.
//
// To stream query results successfully, the IAM principal with permission to call
// GetQueryResults also must have permissions to the Amazon S3 GetObject action
// for the Athena query results location.
//
// IAM principals with permission to the Amazon S3 GetObject action for the query
// results location are able to retrieve query results from Amazon S3 even if
// permission to the GetQueryResults action is denied. To restrict user or role
// access, ensure that Amazon S3 permissions to the Athena query location are
// denied.
//
// [Working with query results, recent queries, and output files]: https://docs.aws.amazon.com/athena/latest/ug/querying.html
func athena_GetQueryResults(cfg aws.Config, client *athena.Client) {
	input := &athena.GetQueryResultsInput{
		// QueryExecutionId: *string, // Required
	}

	if len(_athenaQueryExecutionId) > 0 {
		input.QueryExecutionId = aws.String(_athenaQueryExecutionId)
	}
	if len(_athenaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _athenaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_athenaNextToken) > 0 {
		input.NextToken = aws.String(_athenaNextToken)
	}
	if len(_athenaQueryResultType) > 0 {
		if err := assignInputField(input, "QueryResultType", _athenaQueryResultType); err != nil {
			log.Errorf("invalid --query-result-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetQueryResults(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*athena.GetQueryResultsOutput
	p := athena.NewGetQueryResultsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns query execution runtime statistics related to a single execution of a
// query if you have access to the workgroup in which the query ran. Statistics
// from the Timeline section of the response object are available as soon as QueryExecutionStatus$State is
// in a SUCCEEDED or FAILED state. The remaining non-timeline statistics in the
// response (like stage-level input and output row count and data size) are updated
// asynchronously and may not be available immediately after a query completes or,
// in some cases, may not be returned. The non-timeline statistics are also not
// included when a query has row-level filters defined in Lake Formation.
func athena_GetQueryRuntimeStatistics(cfg aws.Config, client *athena.Client) {
	input := &athena.GetQueryRuntimeStatisticsInput{
		// QueryExecutionId: *string, // Required
	}

	if len(_athenaQueryExecutionId) > 0 {
		input.QueryExecutionId = aws.String(_athenaQueryExecutionId)
	}

	if resp, err := client.GetQueryRuntimeStatistics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the Live UI/Persistence UI for a session.
func athena_GetResourceDashboard(cfg aws.Config, client *athena.Client) {
	input := &athena.GetResourceDashboardInput{
		// ResourceARN: *string, // Required
	}

	if len(_athenaResourceARN) > 0 {
		input.ResourceARN = aws.String(_athenaResourceARN)
	}

	if resp, err := client.GetResourceDashboard(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the full details of a previously created session, including the session
// status and configuration.
func athena_GetSession(cfg aws.Config, client *athena.Client) {
	input := &athena.GetSessionInput{
		// SessionId: *string, // Required
	}

	if len(_athenaSessionId) > 0 {
		input.SessionId = aws.String(_athenaSessionId)
	}

	if resp, err := client.GetSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a connection endpoint and authentication token for a given session Id.
func athena_GetSessionEndpoint(cfg aws.Config, client *athena.Client) {
	input := &athena.GetSessionEndpointInput{
		// SessionId: *string, // Required
	}

	if len(_athenaSessionId) > 0 {
		input.SessionId = aws.String(_athenaSessionId)
	}

	if resp, err := client.GetSessionEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the current status of a session.
func athena_GetSessionStatus(cfg aws.Config, client *athena.Client) {
	input := &athena.GetSessionStatusInput{
		// SessionId: *string, // Required
	}

	if len(_athenaSessionId) > 0 {
		input.SessionId = aws.String(_athenaSessionId)
	}

	if resp, err := client.GetSessionStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns table metadata for the specified catalog, database, and table.
func athena_GetTableMetadata(cfg aws.Config, client *athena.Client) {
	input := &athena.GetTableMetadataInput{
		// CatalogName: *string, // Required
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_athenaCatalogName) > 0 {
		input.CatalogName = aws.String(_athenaCatalogName)
	}
	if len(_athenaDatabaseName) > 0 {
		input.DatabaseName = aws.String(_athenaDatabaseName)
	}
	if len(_athenaTableName) > 0 {
		input.TableName = aws.String(_athenaTableName)
	}
	if len(_athenaWorkGroup) > 0 {
		input.WorkGroup = aws.String(_athenaWorkGroup)
	}

	if resp, err := client.GetTableMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the workgroup with the specified name.
func athena_GetWorkGroup(cfg aws.Config, client *athena.Client) {
	input := &athena.GetWorkGroupInput{
		// WorkGroup: *string, // Required
	}

	if len(_athenaWorkGroup) > 0 {
		input.WorkGroup = aws.String(_athenaWorkGroup)
	}

	if resp, err := client.GetWorkGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports a single ipynb file to a Spark enabled workgroup. To import the
// notebook, the request must specify a value for either Payload or
// NoteBookS3LocationUri . If neither is specified or both are specified, an
// InvalidRequestException occurs. The maximum file size that can be imported is 10
// megabytes. If an ipynb file with the same name already exists in the workgroup,
// throws an error.
func athena_ImportNotebook(cfg aws.Config, client *athena.Client) {
	input := &athena.ImportNotebookInput{
		// Name: *string, // Required
		// Type: types.NotebookType, // Required
		// WorkGroup: *string, // Required
	}

	if len(_athenaName) > 0 {
		input.Name = aws.String(_athenaName)
	}
	if len(_athenaType) > 0 {
		if err := assignInputField(input, "Type", _athenaType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_athenaWorkGroup) > 0 {
		input.WorkGroup = aws.String(_athenaWorkGroup)
	}
	if len(_athenaClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_athenaClientRequestToken)
	}
	if len(_athenaNotebookS3LocationUri) > 0 {
		input.NotebookS3LocationUri = aws.String(_athenaNotebookS3LocationUri)
	}
	if len(_athenaPayload) > 0 {
		input.Payload = aws.String(_athenaPayload)
	}

	if resp, err := client.ImportNotebook(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the supported DPU sizes for the supported application runtimes (for
// example, Athena notebook version 1 ).
func athena_ListApplicationDPUSizes(cfg aws.Config, client *athena.Client) {
	input := &athena.ListApplicationDPUSizesInput{}

	if len(_athenaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _athenaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_athenaNextToken) > 0 {
		input.NextToken = aws.String(_athenaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApplicationDPUSizes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*athena.ListApplicationDPUSizesOutput
	p := athena.NewListApplicationDPUSizesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists the calculations that have been submitted to a session in descending
// order. Newer calculations are listed first; older calculations are listed later.
func athena_ListCalculationExecutions(cfg aws.Config, client *athena.Client) {
	input := &athena.ListCalculationExecutionsInput{
		// SessionId: *string, // Required
	}

	if len(_athenaSessionId) > 0 {
		input.SessionId = aws.String(_athenaSessionId)
	}
	if len(_athenaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _athenaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_athenaNextToken) > 0 {
		input.NextToken = aws.String(_athenaNextToken)
	}
	if len(_athenaStateFilter) > 0 {
		if err := assignInputField(input, "StateFilter", _athenaStateFilter); err != nil {
			log.Errorf("invalid --state-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCalculationExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*athena.ListCalculationExecutionsOutput
	p := athena.NewListCalculationExecutionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists the capacity reservations for the current account.
func athena_ListCapacityReservations(cfg aws.Config, client *athena.Client) {
	input := &athena.ListCapacityReservationsInput{}

	if len(_athenaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _athenaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_athenaNextToken) > 0 {
		input.NextToken = aws.String(_athenaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCapacityReservations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*athena.ListCapacityReservationsOutput
	p := athena.NewListCapacityReservationsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists the data catalogs in the current Amazon Web Services account.
// In the Athena console, data catalogs are listed as "data sources" on the Data
// sources page under the Data source name column.
func athena_ListDataCatalogs(cfg aws.Config, client *athena.Client) {
	input := &athena.ListDataCatalogsInput{}

	if len(_athenaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _athenaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_athenaNextToken) > 0 {
		input.NextToken = aws.String(_athenaNextToken)
	}
	if len(_athenaWorkGroup) > 0 {
		input.WorkGroup = aws.String(_athenaWorkGroup)
	}

	if disablePaginator() {
		if resp, err := client.ListDataCatalogs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*athena.ListDataCatalogsOutput
	p := athena.NewListDataCatalogsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists the databases in the specified data catalog.
func athena_ListDatabases(cfg aws.Config, client *athena.Client) {
	input := &athena.ListDatabasesInput{
		// CatalogName: *string, // Required
	}

	if len(_athenaCatalogName) > 0 {
		input.CatalogName = aws.String(_athenaCatalogName)
	}
	if len(_athenaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _athenaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_athenaNextToken) > 0 {
		input.NextToken = aws.String(_athenaNextToken)
	}
	if len(_athenaWorkGroup) > 0 {
		input.WorkGroup = aws.String(_athenaWorkGroup)
	}

	if disablePaginator() {
		if resp, err := client.ListDatabases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*athena.ListDatabasesOutput
	p := athena.NewListDatabasesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of engine versions that are available to choose from, including
// the Auto option.
func athena_ListEngineVersions(cfg aws.Config, client *athena.Client) {
	input := &athena.ListEngineVersionsInput{}

	if len(_athenaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _athenaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_athenaNextToken) > 0 {
		input.NextToken = aws.String(_athenaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEngineVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*athena.ListEngineVersionsOutput
	p := athena.NewListEngineVersionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists, in descending order, the executors that joined a session. Newer
// executors are listed first; older executors are listed later. The result can be
// optionally filtered by state.
func athena_ListExecutors(cfg aws.Config, client *athena.Client) {
	input := &athena.ListExecutorsInput{
		// SessionId: *string, // Required
	}

	if len(_athenaSessionId) > 0 {
		input.SessionId = aws.String(_athenaSessionId)
	}
	if len(_athenaExecutorStateFilter) > 0 {
		if err := assignInputField(input, "ExecutorStateFilter", _athenaExecutorStateFilter); err != nil {
			log.Errorf("invalid --executor-state-filter: %s", err.Error())
			return
		}
	}
	if len(_athenaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _athenaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_athenaNextToken) > 0 {
		input.NextToken = aws.String(_athenaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListExecutors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*athena.ListExecutorsOutput
	p := athena.NewListExecutorsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Provides a list of available query IDs only for queries saved in the specified
// workgroup. Requires that you have access to the specified workgroup. If a
// workgroup is not specified, lists the saved queries for the primary workgroup.
func athena_ListNamedQueries(cfg aws.Config, client *athena.Client) {
	input := &athena.ListNamedQueriesInput{}

	if len(_athenaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _athenaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_athenaNextToken) > 0 {
		input.NextToken = aws.String(_athenaNextToken)
	}
	if len(_athenaWorkGroup) > 0 {
		input.WorkGroup = aws.String(_athenaWorkGroup)
	}

	if disablePaginator() {
		if resp, err := client.ListNamedQueries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*athena.ListNamedQueriesOutput
	p := athena.NewListNamedQueriesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Displays the notebook files for the specified workgroup in paginated format.
func athena_ListNotebookMetadata(cfg aws.Config, client *athena.Client) {
	input := &athena.ListNotebookMetadataInput{
		// WorkGroup: *string, // Required
	}

	if len(_athenaWorkGroup) > 0 {
		input.WorkGroup = aws.String(_athenaWorkGroup)
	}
	if len(_athenaFilters) > 0 {
		if err := assignInputField(input, "Filters", _athenaFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_athenaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _athenaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_athenaNextToken) > 0 {
		input.NextToken = aws.String(_athenaNextToken)
	}

	if resp, err := client.ListNotebookMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists, in descending order, the sessions that have been created in a notebook
// that are in an active state like CREATING , CREATED , IDLE or BUSY . Newer
// sessions are listed first; older sessions are listed later.
func athena_ListNotebookSessions(cfg aws.Config, client *athena.Client) {
	input := &athena.ListNotebookSessionsInput{
		// NotebookId: *string, // Required
	}

	if len(_athenaNotebookId) > 0 {
		input.NotebookId = aws.String(_athenaNotebookId)
	}
	if len(_athenaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _athenaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_athenaNextToken) > 0 {
		input.NextToken = aws.String(_athenaNextToken)
	}

	if resp, err := client.ListNotebookSessions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the prepared statements in the specified workgroup.
func athena_ListPreparedStatements(cfg aws.Config, client *athena.Client) {
	input := &athena.ListPreparedStatementsInput{
		// WorkGroup: *string, // Required
	}

	if len(_athenaWorkGroup) > 0 {
		input.WorkGroup = aws.String(_athenaWorkGroup)
	}
	if len(_athenaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _athenaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_athenaNextToken) > 0 {
		input.NextToken = aws.String(_athenaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPreparedStatements(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*athena.ListPreparedStatementsOutput
	p := athena.NewListPreparedStatementsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Provides a list of available query execution IDs for the queries in the
// specified workgroup. Athena keeps a query history for 45 days. If a workgroup is
// not specified, returns a list of query execution IDs for the primary workgroup.
// Requires you to have access to the workgroup in which the queries ran.
func athena_ListQueryExecutions(cfg aws.Config, client *athena.Client) {
	input := &athena.ListQueryExecutionsInput{}

	if len(_athenaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _athenaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_athenaNextToken) > 0 {
		input.NextToken = aws.String(_athenaNextToken)
	}
	if len(_athenaWorkGroup) > 0 {
		input.WorkGroup = aws.String(_athenaWorkGroup)
	}

	if disablePaginator() {
		if resp, err := client.ListQueryExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*athena.ListQueryExecutionsOutput
	p := athena.NewListQueryExecutionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists the sessions in a workgroup that are in an active state like CREATING ,
// CREATED , IDLE , or BUSY . Newer sessions are listed first; older sessions are
// listed later.
func athena_ListSessions(cfg aws.Config, client *athena.Client) {
	input := &athena.ListSessionsInput{
		// WorkGroup: *string, // Required
	}

	if len(_athenaWorkGroup) > 0 {
		input.WorkGroup = aws.String(_athenaWorkGroup)
	}
	if len(_athenaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _athenaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_athenaNextToken) > 0 {
		input.NextToken = aws.String(_athenaNextToken)
	}
	if len(_athenaStateFilter) > 0 {
		if err := assignInputField(input, "StateFilter", _athenaStateFilter); err != nil {
			log.Errorf("invalid --state-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSessions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*athena.ListSessionsOutput
	p := athena.NewListSessionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists the metadata for the tables in the specified data catalog database.
func athena_ListTableMetadata(cfg aws.Config, client *athena.Client) {
	input := &athena.ListTableMetadataInput{
		// CatalogName: *string, // Required
		// DatabaseName: *string, // Required
	}

	if len(_athenaCatalogName) > 0 {
		input.CatalogName = aws.String(_athenaCatalogName)
	}
	if len(_athenaDatabaseName) > 0 {
		input.DatabaseName = aws.String(_athenaDatabaseName)
	}
	if len(_athenaExpression) > 0 {
		input.Expression = aws.String(_athenaExpression)
	}
	if len(_athenaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _athenaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_athenaNextToken) > 0 {
		input.NextToken = aws.String(_athenaNextToken)
	}
	if len(_athenaWorkGroup) > 0 {
		input.WorkGroup = aws.String(_athenaWorkGroup)
	}

	if disablePaginator() {
		if resp, err := client.ListTableMetadata(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*athena.ListTableMetadataOutput
	p := athena.NewListTableMetadataPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists the tags associated with an Athena resource.
func athena_ListTagsForResource(cfg aws.Config, client *athena.Client) {
	input := &athena.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_athenaResourceARN) > 0 {
		input.ResourceARN = aws.String(_athenaResourceARN)
	}
	if len(_athenaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _athenaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_athenaNextToken) > 0 {
		input.NextToken = aws.String(_athenaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*athena.ListTagsForResourceOutput
	p := athena.NewListTagsForResourcePaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists available workgroups for the account.
func athena_ListWorkGroups(cfg aws.Config, client *athena.Client) {
	input := &athena.ListWorkGroupsInput{}

	if len(_athenaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _athenaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_athenaNextToken) > 0 {
		input.NextToken = aws.String(_athenaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*athena.ListWorkGroupsOutput
	p := athena.NewListWorkGroupsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Puts a new capacity assignment configuration for a specified capacity
// reservation. If a capacity assignment configuration already exists for the
// capacity reservation, replaces the existing capacity assignment configuration.
func athena_PutCapacityAssignmentConfiguration(cfg aws.Config, client *athena.Client) {
	input := &athena.PutCapacityAssignmentConfigurationInput{
		// CapacityAssignments: []types.CapacityAssignment, // Required
		// CapacityReservationName: *string, // Required
	}

	if len(_athenaCapacityAssignments) > 0 {
		if err := assignInputField(input, "CapacityAssignments", _athenaCapacityAssignments); err != nil {
			log.Errorf("invalid --capacity-assignments: %s", err.Error())
			return
		}
	}
	if len(_athenaCapacityReservationName) > 0 {
		input.CapacityReservationName = aws.String(_athenaCapacityReservationName)
	}

	if resp, err := client.PutCapacityAssignmentConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Submits calculations for execution within a session. You can supply the code to
// run as an inline code block within the request.
//
// The request syntax requires the StartCalculationExecutionRequest$CodeBlock parameter or the CalculationConfiguration$CodeBlock parameter, but not both.
// Because CalculationConfiguration$CodeBlockis deprecated, use the StartCalculationExecutionRequest$CodeBlock parameter instead.
func athena_StartCalculationExecution(cfg aws.Config, client *athena.Client) {
	input := &athena.StartCalculationExecutionInput{
		// SessionId: *string, // Required
	}

	if len(_athenaSessionId) > 0 {
		input.SessionId = aws.String(_athenaSessionId)
	}
	if len(_athenaCalculationConfiguration) > 0 {
		if err := assignInputField(input, "CalculationConfiguration", _athenaCalculationConfiguration); err != nil {
			log.Errorf("invalid --calculation-configuration: %s", err.Error())
			return
		}
	}
	if len(_athenaClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_athenaClientRequestToken)
	}
	if len(_athenaCodeBlock) > 0 {
		input.CodeBlock = aws.String(_athenaCodeBlock)
	}
	if len(_athenaDescription) > 0 {
		input.Description = aws.String(_athenaDescription)
	}

	if resp, err := client.StartCalculationExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Runs the SQL query statements contained in the Query . Requires you to have
// access to the workgroup in which the query ran. Running queries against an
// external catalog requires GetDataCatalogpermission to the catalog. For code samples using the
// Amazon Web Services SDK for Java, see [Examples and Code Samples]in the Amazon Athena User Guide.
//
// [Examples and Code Samples]: http://docs.aws.amazon.com/athena/latest/ug/code-samples.html
func athena_StartQueryExecution(cfg aws.Config, client *athena.Client) {
	input := &athena.StartQueryExecutionInput{
		// QueryString: *string, // Required
	}

	if len(_athenaQueryString) > 0 {
		input.QueryString = aws.String(_athenaQueryString)
	}
	if len(_athenaClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_athenaClientRequestToken)
	}
	if len(_athenaEngineConfiguration) > 0 {
		if err := assignInputField(input, "EngineConfiguration", _athenaEngineConfiguration); err != nil {
			log.Errorf("invalid --engine-configuration: %s", err.Error())
			return
		}
	}
	if len(_athenaExecutionParameters) > 0 {
		input.ExecutionParameters = append([]string(nil), _athenaExecutionParameters...)
	}
	if len(_athenaQueryExecutionContext) > 0 {
		if err := assignInputField(input, "QueryExecutionContext", _athenaQueryExecutionContext); err != nil {
			log.Errorf("invalid --query-execution-context: %s", err.Error())
			return
		}
	}
	if len(_athenaResultConfiguration) > 0 {
		if err := assignInputField(input, "ResultConfiguration", _athenaResultConfiguration); err != nil {
			log.Errorf("invalid --result-configuration: %s", err.Error())
			return
		}
	}
	if len(_athenaResultReuseConfiguration) > 0 {
		if err := assignInputField(input, "ResultReuseConfiguration", _athenaResultReuseConfiguration); err != nil {
			log.Errorf("invalid --result-reuse-configuration: %s", err.Error())
			return
		}
	}
	if len(_athenaWorkGroup) > 0 {
		input.WorkGroup = aws.String(_athenaWorkGroup)
	}

	if resp, err := client.StartQueryExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a session for running calculations within a workgroup. The session is
// ready when it reaches an IDLE state.
func athena_StartSession(cfg aws.Config, client *athena.Client) {
	input := &athena.StartSessionInput{
		// EngineConfiguration: *types.EngineConfiguration, // Required
		// WorkGroup: *string, // Required
	}

	if len(_athenaEngineConfiguration) > 0 {
		if err := assignInputField(input, "EngineConfiguration", _athenaEngineConfiguration); err != nil {
			log.Errorf("invalid --engine-configuration: %s", err.Error())
			return
		}
	}
	if len(_athenaWorkGroup) > 0 {
		input.WorkGroup = aws.String(_athenaWorkGroup)
	}
	if len(_athenaClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_athenaClientRequestToken)
	}
	if len(_athenaCopyWorkGroupTags) > 0 {
		if err := assignInputField(input, "CopyWorkGroupTags", _athenaCopyWorkGroupTags); err != nil {
			log.Errorf("invalid --copy-work-group-tags: %s", err.Error())
			return
		}
	}
	if len(_athenaDescription) > 0 {
		input.Description = aws.String(_athenaDescription)
	}
	if len(_athenaExecutionRole) > 0 {
		input.ExecutionRole = aws.String(_athenaExecutionRole)
	}
	if len(_athenaMonitoringConfiguration) > 0 {
		if err := assignInputField(input, "MonitoringConfiguration", _athenaMonitoringConfiguration); err != nil {
			log.Errorf("invalid --monitoring-configuration: %s", err.Error())
			return
		}
	}
	if len(_athenaNotebookVersion) > 0 {
		input.NotebookVersion = aws.String(_athenaNotebookVersion)
	}
	if len(_athenaSessionIdleTimeoutInMinutes) > 0 {
		if err := assignInputField(input, "SessionIdleTimeoutInMinutes", _athenaSessionIdleTimeoutInMinutes); err != nil {
			log.Errorf("invalid --session-idle-timeout-in-minutes: %s", err.Error())
			return
		}
	}
	if len(_athenaTags) > 0 {
		if err := assignInputField(input, "Tags", _athenaTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Requests the cancellation of a calculation. A StopCalculationExecution call on
// a calculation that is already in a terminal state (for example, STOPPED , FAILED
// , or COMPLETED ) succeeds but has no effect.
//
// Cancelling a calculation is done on a best effort basis. If a calculation
// cannot be cancelled, you can be charged for its completion. If you are concerned
// about being charged for a calculation that cannot be cancelled, consider
// terminating the session in which the calculation is running.
func athena_StopCalculationExecution(cfg aws.Config, client *athena.Client) {
	input := &athena.StopCalculationExecutionInput{
		// CalculationExecutionId: *string, // Required
	}

	if len(_athenaCalculationExecutionId) > 0 {
		input.CalculationExecutionId = aws.String(_athenaCalculationExecutionId)
	}

	if resp, err := client.StopCalculationExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops a query execution. Requires you to have access to the workgroup in which
// the query ran.
func athena_StopQueryExecution(cfg aws.Config, client *athena.Client) {
	input := &athena.StopQueryExecutionInput{
		// QueryExecutionId: *string, // Required
	}

	if len(_athenaQueryExecutionId) > 0 {
		input.QueryExecutionId = aws.String(_athenaQueryExecutionId)
	}

	if resp, err := client.StopQueryExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more tags to an Athena resource. A tag is a label that you assign
// to a resource. Each tag consists of a key and an optional value, both of which
// you define. For example, you can use tags to categorize Athena workgroups, data
// catalogs, or capacity reservations by purpose, owner, or environment. Use a
// consistent set of tag keys to make it easier to search and filter the resources
// in your account. For best practices, see [Tagging Best Practices]. Tag keys can be from 1 to 128 UTF-8
// Unicode characters, and tag values can be from 0 to 256 UTF-8 Unicode
// characters. Tags can use letters and numbers representable in UTF-8, and the
// following characters: + - = . _ : / (at). Tag keys and values are case-sensitive.
// Tag keys must be unique per resource. If you specify more than one tag, separate
// them by commas.
//
// [Tagging Best Practices]: https://docs.aws.amazon.com/whitepapers/latest/tagging-best-practices/tagging-best-practices.html
func athena_TagResource(cfg aws.Config, client *athena.Client) {
	input := &athena.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_athenaResourceARN) > 0 {
		input.ResourceARN = aws.String(_athenaResourceARN)
	}
	if len(_athenaTags) > 0 {
		if err := assignInputField(input, "Tags", _athenaTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Terminates an active session. A TerminateSession call on a session that is
// already inactive (for example, in a FAILED , TERMINATED or TERMINATING state)
// succeeds but has no effect. Calculations running in the session when
// TerminateSession is called are forcefully stopped, but may display as FAILED
// instead of STOPPED .
func athena_TerminateSession(cfg aws.Config, client *athena.Client) {
	input := &athena.TerminateSessionInput{
		// SessionId: *string, // Required
	}

	if len(_athenaSessionId) > 0 {
		input.SessionId = aws.String(_athenaSessionId)
	}

	if resp, err := client.TerminateSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes one or more tags from an Athena resource.
func athena_UntagResource(cfg aws.Config, client *athena.Client) {
	input := &athena.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_athenaResourceARN) > 0 {
		input.ResourceARN = aws.String(_athenaResourceARN)
	}
	if len(_athenaTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _athenaTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the number of requested data processing units for the capacity
// reservation with the specified name.
func athena_UpdateCapacityReservation(cfg aws.Config, client *athena.Client) {
	input := &athena.UpdateCapacityReservationInput{
		// Name: *string, // Required
		// TargetDpus: *int32, // Required
	}

	if len(_athenaName) > 0 {
		input.Name = aws.String(_athenaName)
	}
	if len(_athenaTargetDpus) > 0 {
		if err := assignInputField(input, "TargetDpus", _athenaTargetDpus); err != nil {
			log.Errorf("invalid --target-dpus: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCapacityReservation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the data catalog that has the specified name.
func athena_UpdateDataCatalog(cfg aws.Config, client *athena.Client) {
	input := &athena.UpdateDataCatalogInput{
		// Name: *string, // Required
		// Type: types.DataCatalogType, // Required
	}

	if len(_athenaName) > 0 {
		input.Name = aws.String(_athenaName)
	}
	if len(_athenaType) > 0 {
		if err := assignInputField(input, "Type", _athenaType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_athenaDescription) > 0 {
		input.Description = aws.String(_athenaDescription)
	}
	if len(_athenaParameters) > 0 {
		if err := assignInputField(input, "Parameters", _athenaParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDataCatalog(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a NamedQuery object. The database or workgroup cannot be updated.
func athena_UpdateNamedQuery(cfg aws.Config, client *athena.Client) {
	input := &athena.UpdateNamedQueryInput{
		// Name: *string, // Required
		// NamedQueryId: *string, // Required
		// QueryString: *string, // Required
	}

	if len(_athenaName) > 0 {
		input.Name = aws.String(_athenaName)
	}
	if len(_athenaNamedQueryId) > 0 {
		input.NamedQueryId = aws.String(_athenaNamedQueryId)
	}
	if len(_athenaQueryString) > 0 {
		input.QueryString = aws.String(_athenaQueryString)
	}
	if len(_athenaDescription) > 0 {
		input.Description = aws.String(_athenaDescription)
	}

	if resp, err := client.UpdateNamedQuery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the contents of a Spark notebook.
func athena_UpdateNotebook(cfg aws.Config, client *athena.Client) {
	input := &athena.UpdateNotebookInput{
		// NotebookId: *string, // Required
		// Payload: *string, // Required
		// Type: types.NotebookType, // Required
	}

	if len(_athenaNotebookId) > 0 {
		input.NotebookId = aws.String(_athenaNotebookId)
	}
	if len(_athenaPayload) > 0 {
		input.Payload = aws.String(_athenaPayload)
	}
	if len(_athenaType) > 0 {
		if err := assignInputField(input, "Type", _athenaType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_athenaClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_athenaClientRequestToken)
	}
	if len(_athenaSessionId) > 0 {
		input.SessionId = aws.String(_athenaSessionId)
	}

	if resp, err := client.UpdateNotebook(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the metadata for a notebook.
func athena_UpdateNotebookMetadata(cfg aws.Config, client *athena.Client) {
	input := &athena.UpdateNotebookMetadataInput{
		// Name: *string, // Required
		// NotebookId: *string, // Required
	}

	if len(_athenaName) > 0 {
		input.Name = aws.String(_athenaName)
	}
	if len(_athenaNotebookId) > 0 {
		input.NotebookId = aws.String(_athenaNotebookId)
	}
	if len(_athenaClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_athenaClientRequestToken)
	}

	if resp, err := client.UpdateNotebookMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a prepared statement.
func athena_UpdatePreparedStatement(cfg aws.Config, client *athena.Client) {
	input := &athena.UpdatePreparedStatementInput{
		// QueryStatement: *string, // Required
		// StatementName: *string, // Required
		// WorkGroup: *string, // Required
	}

	if len(_athenaQueryStatement) > 0 {
		input.QueryStatement = aws.String(_athenaQueryStatement)
	}
	if len(_athenaStatementName) > 0 {
		input.StatementName = aws.String(_athenaStatementName)
	}
	if len(_athenaWorkGroup) > 0 {
		input.WorkGroup = aws.String(_athenaWorkGroup)
	}
	if len(_athenaDescription) > 0 {
		input.Description = aws.String(_athenaDescription)
	}

	if resp, err := client.UpdatePreparedStatement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the workgroup with the specified name. The workgroup's name cannot be
// changed. Only ConfigurationUpdates can be specified.
func athena_UpdateWorkGroup(cfg aws.Config, client *athena.Client) {
	input := &athena.UpdateWorkGroupInput{
		// WorkGroup: *string, // Required
	}

	if len(_athenaWorkGroup) > 0 {
		input.WorkGroup = aws.String(_athenaWorkGroup)
	}
	if len(_athenaConfigurationUpdates) > 0 {
		if err := assignInputField(input, "ConfigurationUpdates", _athenaConfigurationUpdates); err != nil {
			log.Errorf("invalid --configuration-updates: %s", err.Error())
			return
		}
	}
	if len(_athenaDescription) > 0 {
		input.Description = aws.String(_athenaDescription)
	}
	if len(_athenaState) > 0 {
		if err := assignInputField(input, "State", _athenaState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateWorkGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_athenaCmd)
	_athenaCmd.Flags().SortFlags = false

	_athenaCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_athenaCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_athenaCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_athenaCmd.Flags().StringVarP(&_athenaCalculationConfiguration, "calculation-configuration", "", "", "Calculation Configuration")
	_athenaCmd.Flags().StringVarP(&_athenaCalculationExecutionId, "calculation-execution-id", "", "", "Calculation Execution ID")
	_athenaCmd.Flags().StringVarP(&_athenaCapacityAssignments, "capacity-assignments", "", "", "Capacity Assignments")
	_athenaCmd.Flags().StringVarP(&_athenaCapacityReservationName, "capacity-reservation-name", "", "", "Capacity Reservation Name")
	_athenaCmd.Flags().StringVarP(&_athenaCatalogName, "catalog-name", "", "", "Catalog Name")
	_athenaCmd.Flags().StringVarP(&_athenaClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_athenaCmd.Flags().StringVarP(&_athenaCodeBlock, "code-block", "", "", "Code Block")
	_athenaCmd.Flags().StringVarP(&_athenaConfiguration, "configuration", "", "", "Configuration")
	_athenaCmd.Flags().StringVarP(&_athenaConfigurationUpdates, "configuration-updates", "", "", "Configuration Updates")
	_athenaCmd.Flags().StringVarP(&_athenaCopyWorkGroupTags, "copy-work-group-tags", "", "", "Copy Work Group Tags")
	_athenaCmd.Flags().StringVarP(&_athenaDatabase, "database", "", "", "Database")
	_athenaCmd.Flags().StringVarP(&_athenaDatabaseName, "database-name", "", "", "Database Name")
	_athenaCmd.Flags().StringVarP(&_athenaDeleteCatalogOnly, "delete-catalog-only", "", "", "Delete Catalog Only")
	_athenaCmd.Flags().StringVarP(&_athenaDescription, "description", "", "", "Description")
	_athenaCmd.Flags().StringVarP(&_athenaEngineConfiguration, "engine-configuration", "", "", "Engine Configuration")
	_athenaCmd.Flags().StringSliceVarP(&_athenaExecutionParameters, "execution-parameters", "", nil, "Execution Parameters")
	_athenaCmd.Flags().StringVarP(&_athenaExecutionRole, "execution-role", "", "", "Execution Role")
	_athenaCmd.Flags().StringVarP(&_athenaExecutorStateFilter, "executor-state-filter", "", "", "Executor State Filter")
	_athenaCmd.Flags().StringVarP(&_athenaExpression, "expression", "", "", "Expression")
	_athenaCmd.Flags().StringVarP(&_athenaFilters, "filters", "", "", "Filters")
	_athenaCmd.Flags().StringVarP(&_athenaMaxResults, "max-results", "", "", "Max Results")
	_athenaCmd.Flags().StringVarP(&_athenaMonitoringConfiguration, "monitoring-configuration", "", "", "Monitoring Configuration")
	_athenaCmd.Flags().StringVarP(&_athenaName, "name", "", "", "Name")
	_athenaCmd.Flags().StringVarP(&_athenaNamedQueryId, "named-query-id", "", "", "Named Query ID")
	_athenaCmd.Flags().StringSliceVarP(&_athenaNamedQueryIds, "named-query-ids", "", nil, "Named Query Ids")
	_athenaCmd.Flags().StringVarP(&_athenaNextToken, "next-token", "", "", "Next Token")
	_athenaCmd.Flags().StringVarP(&_athenaNotebookId, "notebook-id", "", "", "Notebook ID")
	_athenaCmd.Flags().StringVarP(&_athenaNotebookS3LocationUri, "notebook-s3-location-uri", "", "", "Notebook S3 Location URI")
	_athenaCmd.Flags().StringVarP(&_athenaNotebookVersion, "notebook-version", "", "", "Notebook Version")
	_athenaCmd.Flags().StringVarP(&_athenaParameters, "parameters", "", "", "Parameters")
	_athenaCmd.Flags().StringVarP(&_athenaPayload, "payload", "", "", "Payload")
	_athenaCmd.Flags().StringSliceVarP(&_athenaPreparedStatementNames, "prepared-statement-names", "", nil, "Prepared Statement Names")
	_athenaCmd.Flags().StringVarP(&_athenaQueryExecutionContext, "query-execution-context", "", "", "Query Execution Context")
	_athenaCmd.Flags().StringVarP(&_athenaQueryExecutionId, "query-execution-id", "", "", "Query Execution ID")
	_athenaCmd.Flags().StringSliceVarP(&_athenaQueryExecutionIds, "query-execution-ids", "", nil, "Query Execution Ids")
	_athenaCmd.Flags().StringVarP(&_athenaQueryResultType, "query-result-type", "", "", "Query Result Type")
	_athenaCmd.Flags().StringVarP(&_athenaQueryStatement, "query-statement", "", "", "Query Statement")
	_athenaCmd.Flags().StringVarP(&_athenaQueryString, "query-string", "", "", "Query String")
	_athenaCmd.Flags().StringVarP(&_athenaRecursiveDeleteOption, "recursive-delete-option", "", "", "Recursive Delete Option")
	_athenaCmd.Flags().StringVarP(&_athenaResourceARN, "resource-arn", "", "", "Resource ARN")
	_athenaCmd.Flags().StringVarP(&_athenaResultConfiguration, "result-configuration", "", "", "Result Configuration")
	_athenaCmd.Flags().StringVarP(&_athenaResultReuseConfiguration, "result-reuse-configuration", "", "", "Result Reuse Configuration")
	_athenaCmd.Flags().StringVarP(&_athenaSessionId, "session-id", "", "", "Session ID")
	_athenaCmd.Flags().StringVarP(&_athenaSessionIdleTimeoutInMinutes, "session-idle-timeout-in-minutes", "", "", "Session Idle Timeout In Minutes")
	_athenaCmd.Flags().StringVarP(&_athenaState, "state", "", "", "State")
	_athenaCmd.Flags().StringVarP(&_athenaStateFilter, "state-filter", "", "", "State Filter")
	_athenaCmd.Flags().StringVarP(&_athenaStatementName, "statement-name", "", "", "Statement Name")
	_athenaCmd.Flags().StringVarP(&_athenaTableName, "table-name", "", "", "Table Name")
	_athenaCmd.Flags().StringSliceVarP(&_athenaTagKeys, "tag-keys", "", nil, "Tag Keys")
	_athenaCmd.Flags().StringVarP(&_athenaTags, "tags", "", "", "Tags")
	_athenaCmd.Flags().StringVarP(&_athenaTargetDpus, "target-dpus", "", "", "Target Dpus")
	_athenaCmd.Flags().StringVarP(&_athenaType, "type", "", "", "Type")
	_athenaCmd.Flags().StringVarP(&_athenaWorkGroup, "work-group", "", "", "Work Group")

	_athenaCmd.Flags().BoolVarP(&_athenaBatchGetNamedQuery, "batch-get-named-query", "", false, "Batch Get Named Query")
	_athenaCmd.Flags().BoolVarP(&_athenaBatchGetPreparedStatement, "batch-get-prepared-statement", "", false, "Batch Get Prepared Statement")
	_athenaCmd.Flags().BoolVarP(&_athenaBatchGetQueryExecution, "batch-get-query-execution", "", false, "Batch Get Query Execution")
	_athenaCmd.Flags().BoolVarP(&_athenaCancelCapacityReservation, "cancel-capacity-reservation", "", false, "Cancel Capacity Reservation")
	_athenaCmd.Flags().BoolVarP(&_athenaCreateCapacityReservation, "create-capacity-reservation", "", false, "Create Capacity Reservation")
	_athenaCmd.Flags().BoolVarP(&_athenaCreateDataCatalog, "create-data-catalog", "", false, "Create Data Catalog")
	_athenaCmd.Flags().BoolVarP(&_athenaCreateNamedQuery, "create-named-query", "", false, "Create Named Query")
	_athenaCmd.Flags().BoolVarP(&_athenaCreateNotebook, "create-notebook", "", false, "Create Notebook")
	_athenaCmd.Flags().BoolVarP(&_athenaCreatePreparedStatement, "create-prepared-statement", "", false, "Create Prepared Statement")
	_athenaCmd.Flags().BoolVarP(&_athenaCreatePresignedNotebookUrl, "create-presigned-notebook-url", "", false, "Create Presigned Notebook URL")
	_athenaCmd.Flags().BoolVarP(&_athenaCreateWorkGroup, "create-work-group", "", false, "Create Work Group")
	_athenaCmd.Flags().BoolVarP(&_athenaDeleteCapacityReservation, "delete-capacity-reservation", "", false, "Delete Capacity Reservation")
	_athenaCmd.Flags().BoolVarP(&_athenaDeleteDataCatalog, "delete-data-catalog", "", false, "Delete Data Catalog")
	_athenaCmd.Flags().BoolVarP(&_athenaDeleteNamedQuery, "delete-named-query", "", false, "Delete Named Query")
	_athenaCmd.Flags().BoolVarP(&_athenaDeleteNotebook, "delete-notebook", "", false, "Delete Notebook")
	_athenaCmd.Flags().BoolVarP(&_athenaDeletePreparedStatement, "delete-prepared-statement", "", false, "Delete Prepared Statement")
	_athenaCmd.Flags().BoolVarP(&_athenaDeleteWorkGroup, "delete-work-group", "", false, "Delete Work Group")
	_athenaCmd.Flags().BoolVarP(&_athenaExportNotebook, "export-notebook", "", false, "Export Notebook")
	_athenaCmd.Flags().BoolVarP(&_athenaGetCalculationExecution, "get-calculation-execution", "", false, "Get Calculation Execution")
	_athenaCmd.Flags().BoolVarP(&_athenaGetCalculationExecutionCode, "get-calculation-execution-code", "", false, "Get Calculation Execution Code")
	_athenaCmd.Flags().BoolVarP(&_athenaGetCalculationExecutionStatus, "get-calculation-execution-status", "", false, "Get Calculation Execution Status")
	_athenaCmd.Flags().BoolVarP(&_athenaGetCapacityAssignmentConfiguration, "get-capacity-assignment-configuration", "", false, "Get Capacity Assignment Configuration")
	_athenaCmd.Flags().BoolVarP(&_athenaGetCapacityReservation, "get-capacity-reservation", "", false, "Get Capacity Reservation")
	_athenaCmd.Flags().BoolVarP(&_athenaGetDataCatalog, "get-data-catalog", "", false, "Get Data Catalog")
	_athenaCmd.Flags().BoolVarP(&_athenaGetDatabase, "get-database", "", false, "Get Database")
	_athenaCmd.Flags().BoolVarP(&_athenaGetNamedQuery, "get-named-query", "", false, "Get Named Query")
	_athenaCmd.Flags().BoolVarP(&_athenaGetNotebookMetadata, "get-notebook-metadata", "", false, "Get Notebook Metadata")
	_athenaCmd.Flags().BoolVarP(&_athenaGetPreparedStatement, "get-prepared-statement", "", false, "Get Prepared Statement")
	_athenaCmd.Flags().BoolVarP(&_athenaGetQueryExecution, "get-query-execution", "", false, "Get Query Execution")
	_athenaCmd.Flags().BoolVarP(&_athenaGetQueryResults, "get-query-results", "", false, "Get Query Results")
	_athenaCmd.Flags().BoolVarP(&_athenaGetQueryRuntimeStatistics, "get-query-runtime-statistics", "", false, "Get Query Runtime Statistics")
	_athenaCmd.Flags().BoolVarP(&_athenaGetResourceDashboard, "get-resource-dashboard", "", false, "Get Resource Dashboard")
	_athenaCmd.Flags().BoolVarP(&_athenaGetSession, "get-session", "", false, "Get Session")
	_athenaCmd.Flags().BoolVarP(&_athenaGetSessionEndpoint, "get-session-endpoint", "", false, "Get Session Endpoint")
	_athenaCmd.Flags().BoolVarP(&_athenaGetSessionStatus, "get-session-status", "", false, "Get Session Status")
	_athenaCmd.Flags().BoolVarP(&_athenaGetTableMetadata, "get-table-metadata", "", false, "Get Table Metadata")
	_athenaCmd.Flags().BoolVarP(&_athenaGetWorkGroup, "get-work-group", "", false, "Get Work Group")
	_athenaCmd.Flags().BoolVarP(&_athenaImportNotebook, "import-notebook", "", false, "Import Notebook")
	_athenaCmd.Flags().BoolVarP(&_athenaListApplicationDPUSizes, "list-application-dpu-sizes", "", false, "List Application Dpu Sizes")
	_athenaCmd.Flags().BoolVarP(&_athenaListCalculationExecutions, "list-calculation-executions", "", false, "List Calculation Executions")
	_athenaCmd.Flags().BoolVarP(&_athenaListCapacityReservations, "list-capacity-reservations", "", false, "List Capacity Reservations")
	_athenaCmd.Flags().BoolVarP(&_athenaListDataCatalogs, "list-data-catalogs", "", false, "List Data Catalogs")
	_athenaCmd.Flags().BoolVarP(&_athenaListDatabases, "list-databases", "", false, "List Databases")
	_athenaCmd.Flags().BoolVarP(&_athenaListEngineVersions, "list-engine-versions", "", false, "List Engine Versions")
	_athenaCmd.Flags().BoolVarP(&_athenaListExecutors, "list-executors", "", false, "List Executors")
	_athenaCmd.Flags().BoolVarP(&_athenaListNamedQueries, "list-named-queries", "", false, "List Named Queries")
	_athenaCmd.Flags().BoolVarP(&_athenaListNotebookMetadata, "list-notebook-metadata", "", false, "List Notebook Metadata")
	_athenaCmd.Flags().BoolVarP(&_athenaListNotebookSessions, "list-notebook-sessions", "", false, "List Notebook Sessions")
	_athenaCmd.Flags().BoolVarP(&_athenaListPreparedStatements, "list-prepared-statements", "", false, "List Prepared Statements")
	_athenaCmd.Flags().BoolVarP(&_athenaListQueryExecutions, "list-query-executions", "", false, "List Query Executions")
	_athenaCmd.Flags().BoolVarP(&_athenaListSessions, "list-sessions", "", false, "List Sessions")
	_athenaCmd.Flags().BoolVarP(&_athenaListTableMetadata, "list-table-metadata", "", false, "List Table Metadata")
	_athenaCmd.Flags().BoolVarP(&_athenaListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_athenaCmd.Flags().BoolVarP(&_athenaListWorkGroups, "list-work-groups", "", false, "List Work Groups")
	_athenaCmd.Flags().BoolVarP(&_athenaPutCapacityAssignmentConfiguration, "put-capacity-assignment-configuration", "", false, "Put Capacity Assignment Configuration")
	_athenaCmd.Flags().BoolVarP(&_athenaStartCalculationExecution, "start-calculation-execution", "", false, "Start Calculation Execution")
	_athenaCmd.Flags().BoolVarP(&_athenaStartQueryExecution, "start-query-execution", "", false, "Start Query Execution")
	_athenaCmd.Flags().BoolVarP(&_athenaStartSession, "start-session", "", false, "Start Session")
	_athenaCmd.Flags().BoolVarP(&_athenaStopCalculationExecution, "stop-calculation-execution", "", false, "Stop Calculation Execution")
	_athenaCmd.Flags().BoolVarP(&_athenaStopQueryExecution, "stop-query-execution", "", false, "Stop Query Execution")
	_athenaCmd.Flags().BoolVarP(&_athenaTagResource, "tag-resource", "", false, "Tag Resource")
	_athenaCmd.Flags().BoolVarP(&_athenaTerminateSession, "terminate-session", "", false, "Terminate Session")
	_athenaCmd.Flags().BoolVarP(&_athenaUntagResource, "untag-resource", "", false, "Untag Resource")
	_athenaCmd.Flags().BoolVarP(&_athenaUpdateCapacityReservation, "update-capacity-reservation", "", false, "Update Capacity Reservation")
	_athenaCmd.Flags().BoolVarP(&_athenaUpdateDataCatalog, "update-data-catalog", "", false, "Update Data Catalog")
	_athenaCmd.Flags().BoolVarP(&_athenaUpdateNamedQuery, "update-named-query", "", false, "Update Named Query")
	_athenaCmd.Flags().BoolVarP(&_athenaUpdateNotebook, "update-notebook", "", false, "Update Notebook")
	_athenaCmd.Flags().BoolVarP(&_athenaUpdateNotebookMetadata, "update-notebook-metadata", "", false, "Update Notebook Metadata")
	_athenaCmd.Flags().BoolVarP(&_athenaUpdatePreparedStatement, "update-prepared-statement", "", false, "Update Prepared Statement")
	_athenaCmd.Flags().BoolVarP(&_athenaUpdateWorkGroup, "update-work-group", "", false, "Update Work Group")

}
