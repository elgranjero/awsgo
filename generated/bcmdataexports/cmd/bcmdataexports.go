package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bcmdataexports"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// bcmdataexportsCmd represents the bcmdataexports command
var _bcmdataexportsCmd = &cobra.Command{
	Use:   "bcmdataexports",
	Short: "AWS bcmdataexports CLI",
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
		client := bcmdataexports.NewFromConfig(cfg)
		if _bcmdataexportsCreateExport {
			bcmdataexports_CreateExport(cfg, client)
			return
		}
		if _bcmdataexportsDeleteExport {
			bcmdataexports_DeleteExport(cfg, client)
			return
		}
		if _bcmdataexportsGetExecution {
			bcmdataexports_GetExecution(cfg, client)
			return
		}
		if _bcmdataexportsGetExport {
			bcmdataexports_GetExport(cfg, client)
			return
		}
		if _bcmdataexportsGetTable {
			bcmdataexports_GetTable(cfg, client)
			return
		}
		if _bcmdataexportsListExecutions {
			bcmdataexports_ListExecutions(cfg, client)
			return
		}
		if _bcmdataexportsListExports {
			bcmdataexports_ListExports(cfg, client)
			return
		}
		if _bcmdataexportsListTables {
			bcmdataexports_ListTables(cfg, client)
			return
		}
		if _bcmdataexportsListTagsForResource {
			bcmdataexports_ListTagsForResource(cfg, client)
			return
		}
		if _bcmdataexportsTagResource {
			bcmdataexports_TagResource(cfg, client)
			return
		}
		if _bcmdataexportsUntagResource {
			bcmdataexports_UntagResource(cfg, client)
			return
		}
		if _bcmdataexportsUpdateExport {
			bcmdataexports_UpdateExport(cfg, client)
			return
		}

	},
}

var (
	_bcmdataexportsCreateExport        bool
	_bcmdataexportsDeleteExport        bool
	_bcmdataexportsGetExecution        bool
	_bcmdataexportsGetExport           bool
	_bcmdataexportsGetTable            bool
	_bcmdataexportsListExecutions      bool
	_bcmdataexportsListExports         bool
	_bcmdataexportsListTables          bool
	_bcmdataexportsListTagsForResource bool
	_bcmdataexportsTagResource         bool
	_bcmdataexportsUntagResource       bool
	_bcmdataexportsUpdateExport        bool

	_bcmdataexportsExecutionId     string
	_bcmdataexportsExport          string
	_bcmdataexportsExportArn       string
	_bcmdataexportsMaxResults      string
	_bcmdataexportsNextToken       string
	_bcmdataexportsResourceArn     string
	_bcmdataexportsResourceTagKeys []string
	_bcmdataexportsResourceTags    string
	_bcmdataexportsTableName       string
	_bcmdataexportsTableProperties string
)

// Creates a data export and specifies the data query, the delivery preference,
// and any optional resource tags.
//
// A DataQuery consists of both a QueryStatement and TableConfigurations .
//
// The QueryStatement is an SQL statement. Data Exports only supports a limited
// subset of the SQL syntax. For more information on the SQL syntax that is
// supported, see [Data query]. To view the available tables and columns, see the [Data Exports table dictionary].
//
// The TableConfigurations is a collection of specified TableProperties for the
// table being queried in the QueryStatement . TableProperties are additional
// configurations you can provide to change the data and schema of a table. Each
// table can have different TableProperties. However, tables are not required to
// have any TableProperties. Each table property has a default value that it
// assumes if not specified. For more information on table configurations, see [Data query].
// To view the table properties available for each table, see the [Data Exports table dictionary]or use the
// ListTables API to get a response of all tables and their available properties.
//
// [Data Exports table dictionary]: https://docs.aws.amazon.com/cur/latest/userguide/de-table-dictionary.html
// [Data query]: https://docs.aws.amazon.com/cur/latest/userguide/de-data-query.html
func bcmdataexports_CreateExport(cfg aws.Config, client *bcmdataexports.Client) {
	input := &bcmdataexports.CreateExportInput{
		// Export: *types.Export, // Required
	}

	if len(_bcmdataexportsExport) > 0 {
		if err := assignInputField(input, "Export", _bcmdataexportsExport); err != nil {
			log.Errorf("invalid --export: %s", err.Error())
			return
		}
	}
	if len(_bcmdataexportsResourceTags) > 0 {
		if err := assignInputField(input, "ResourceTags", _bcmdataexportsResourceTags); err != nil {
			log.Errorf("invalid --resource-tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing data export.
func bcmdataexports_DeleteExport(cfg aws.Config, client *bcmdataexports.Client) {
	input := &bcmdataexports.DeleteExportInput{
		// ExportArn: *string, // Required
	}

	if len(_bcmdataexportsExportArn) > 0 {
		input.ExportArn = aws.String(_bcmdataexportsExportArn)
	}

	if resp, err := client.DeleteExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Exports data based on the source data update.
func bcmdataexports_GetExecution(cfg aws.Config, client *bcmdataexports.Client) {
	input := &bcmdataexports.GetExecutionInput{
		// ExecutionId: *string, // Required
		// ExportArn: *string, // Required
	}

	if len(_bcmdataexportsExecutionId) > 0 {
		input.ExecutionId = aws.String(_bcmdataexportsExecutionId)
	}
	if len(_bcmdataexportsExportArn) > 0 {
		input.ExportArn = aws.String(_bcmdataexportsExportArn)
	}

	if resp, err := client.GetExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Views the definition of an existing data export.
func bcmdataexports_GetExport(cfg aws.Config, client *bcmdataexports.Client) {
	input := &bcmdataexports.GetExportInput{
		// ExportArn: *string, // Required
	}

	if len(_bcmdataexportsExportArn) > 0 {
		input.ExportArn = aws.String(_bcmdataexportsExportArn)
	}

	if resp, err := client.GetExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the metadata for the specified table and table properties. This
// includes the list of columns in the table schema, their data types, and column
// descriptions.
func bcmdataexports_GetTable(cfg aws.Config, client *bcmdataexports.Client) {
	input := &bcmdataexports.GetTableInput{
		// TableName: *string, // Required
	}

	if len(_bcmdataexportsTableName) > 0 {
		input.TableName = aws.String(_bcmdataexportsTableName)
	}
	if len(_bcmdataexportsTableProperties) > 0 {
		if err := assignInputField(input, "TableProperties", _bcmdataexportsTableProperties); err != nil {
			log.Errorf("invalid --table-properties: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the historical executions for the export.
func bcmdataexports_ListExecutions(cfg aws.Config, client *bcmdataexports.Client) {
	input := &bcmdataexports.ListExecutionsInput{
		// ExportArn: *string, // Required
	}

	if len(_bcmdataexportsExportArn) > 0 {
		input.ExportArn = aws.String(_bcmdataexportsExportArn)
	}
	if len(_bcmdataexportsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bcmdataexportsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bcmdataexportsNextToken) > 0 {
		input.NextToken = aws.String(_bcmdataexportsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bcmdataexports.ListExecutionsOutput
	p := bcmdataexports.NewListExecutionsPaginator(client, input)
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

// Lists all data export definitions.
func bcmdataexports_ListExports(cfg aws.Config, client *bcmdataexports.Client) {
	input := &bcmdataexports.ListExportsInput{}

	if len(_bcmdataexportsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bcmdataexportsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bcmdataexportsNextToken) > 0 {
		input.NextToken = aws.String(_bcmdataexportsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListExports(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bcmdataexports.ListExportsOutput
	p := bcmdataexports.NewListExportsPaginator(client, input)
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

// Lists all available tables in data exports.
func bcmdataexports_ListTables(cfg aws.Config, client *bcmdataexports.Client) {
	input := &bcmdataexports.ListTablesInput{}

	if len(_bcmdataexportsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bcmdataexportsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bcmdataexportsNextToken) > 0 {
		input.NextToken = aws.String(_bcmdataexportsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTables(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*bcmdataexports.ListTablesOutput
	p := bcmdataexports.NewListTablesPaginator(client, input)
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

// List tags associated with an existing data export.
func bcmdataexports_ListTagsForResource(cfg aws.Config, client *bcmdataexports.Client) {
	input := &bcmdataexports.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_bcmdataexportsResourceArn) > 0 {
		input.ResourceArn = aws.String(_bcmdataexportsResourceArn)
	}
	if len(_bcmdataexportsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _bcmdataexportsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_bcmdataexportsNextToken) > 0 {
		input.NextToken = aws.String(_bcmdataexportsNextToken)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds tags for an existing data export definition.
func bcmdataexports_TagResource(cfg aws.Config, client *bcmdataexports.Client) {
	input := &bcmdataexports.TagResourceInput{
		// ResourceArn: *string, // Required
		// ResourceTags: []types.ResourceTag, // Required
	}

	if len(_bcmdataexportsResourceArn) > 0 {
		input.ResourceArn = aws.String(_bcmdataexportsResourceArn)
	}
	if len(_bcmdataexportsResourceTags) > 0 {
		if err := assignInputField(input, "ResourceTags", _bcmdataexportsResourceTags); err != nil {
			log.Errorf("invalid --resource-tags: %s", err.Error())
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

// Deletes tags associated with an existing data export definition.
func bcmdataexports_UntagResource(cfg aws.Config, client *bcmdataexports.Client) {
	input := &bcmdataexports.UntagResourceInput{
		// ResourceArn: *string, // Required
		// ResourceTagKeys: []string, // Required
	}

	if len(_bcmdataexportsResourceArn) > 0 {
		input.ResourceArn = aws.String(_bcmdataexportsResourceArn)
	}
	if len(_bcmdataexportsResourceTagKeys) > 0 {
		input.ResourceTagKeys = append([]string(nil), _bcmdataexportsResourceTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing data export by overwriting all export parameters. All
// export parameters must be provided in the UpdateExport request.
func bcmdataexports_UpdateExport(cfg aws.Config, client *bcmdataexports.Client) {
	input := &bcmdataexports.UpdateExportInput{
		// Export: *types.Export, // Required
		// ExportArn: *string, // Required
	}

	if len(_bcmdataexportsExport) > 0 {
		if err := assignInputField(input, "Export", _bcmdataexportsExport); err != nil {
			log.Errorf("invalid --export: %s", err.Error())
			return
		}
	}
	if len(_bcmdataexportsExportArn) > 0 {
		input.ExportArn = aws.String(_bcmdataexportsExportArn)
	}

	if resp, err := client.UpdateExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_bcmdataexportsCmd)
	_bcmdataexportsCmd.Flags().SortFlags = false

	_bcmdataexportsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_bcmdataexportsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_bcmdataexportsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_bcmdataexportsCmd.Flags().StringVarP(&_bcmdataexportsExecutionId, "execution-id", "", "", "Execution ID")
	_bcmdataexportsCmd.Flags().StringVarP(&_bcmdataexportsExport, "export", "", "", "Export")
	_bcmdataexportsCmd.Flags().StringVarP(&_bcmdataexportsExportArn, "export-arn", "", "", "Export ARN")
	_bcmdataexportsCmd.Flags().StringVarP(&_bcmdataexportsMaxResults, "max-results", "", "", "Max Results")
	_bcmdataexportsCmd.Flags().StringVarP(&_bcmdataexportsNextToken, "next-token", "", "", "Next Token")
	_bcmdataexportsCmd.Flags().StringVarP(&_bcmdataexportsResourceArn, "resource-arn", "", "", "Resource ARN")
	_bcmdataexportsCmd.Flags().StringSliceVarP(&_bcmdataexportsResourceTagKeys, "resource-tag-keys", "", nil, "Resource Tag Keys")
	_bcmdataexportsCmd.Flags().StringVarP(&_bcmdataexportsResourceTags, "resource-tags", "", "", "Resource Tags")
	_bcmdataexportsCmd.Flags().StringVarP(&_bcmdataexportsTableName, "table-name", "", "", "Table Name")
	_bcmdataexportsCmd.Flags().StringVarP(&_bcmdataexportsTableProperties, "table-properties", "", "", "Table Properties")

	_bcmdataexportsCmd.Flags().BoolVarP(&_bcmdataexportsCreateExport, "create-export", "", false, "Create Export")
	_bcmdataexportsCmd.Flags().BoolVarP(&_bcmdataexportsDeleteExport, "delete-export", "", false, "Delete Export")
	_bcmdataexportsCmd.Flags().BoolVarP(&_bcmdataexportsGetExecution, "get-execution", "", false, "Get Execution")
	_bcmdataexportsCmd.Flags().BoolVarP(&_bcmdataexportsGetExport, "get-export", "", false, "Get Export")
	_bcmdataexportsCmd.Flags().BoolVarP(&_bcmdataexportsGetTable, "get-table", "", false, "Get Table")
	_bcmdataexportsCmd.Flags().BoolVarP(&_bcmdataexportsListExecutions, "list-executions", "", false, "List Executions")
	_bcmdataexportsCmd.Flags().BoolVarP(&_bcmdataexportsListExports, "list-exports", "", false, "List Exports")
	_bcmdataexportsCmd.Flags().BoolVarP(&_bcmdataexportsListTables, "list-tables", "", false, "List Tables")
	_bcmdataexportsCmd.Flags().BoolVarP(&_bcmdataexportsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_bcmdataexportsCmd.Flags().BoolVarP(&_bcmdataexportsTagResource, "tag-resource", "", false, "Tag Resource")
	_bcmdataexportsCmd.Flags().BoolVarP(&_bcmdataexportsUntagResource, "untag-resource", "", false, "Untag Resource")
	_bcmdataexportsCmd.Flags().BoolVarP(&_bcmdataexportsUpdateExport, "update-export", "", false, "Update Export")

}
