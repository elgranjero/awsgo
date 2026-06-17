package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appsync"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// appsyncCmd represents the appsync command
var _appsyncCmd = &cobra.Command{
	Use:   "appsync",
	Short: "AWS appsync CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := appsync.NewFromConfig(cfg)
		if _appsyncAssociateApi {
			appsync_AssociateApi(cfg, client)
			return
		}
		if _appsyncAssociateMergedGraphqlApi {
			appsync_AssociateMergedGraphqlApi(cfg, client)
			return
		}
		if _appsyncAssociateSourceGraphqlApi {
			appsync_AssociateSourceGraphqlApi(cfg, client)
			return
		}
		if _appsyncCreateApi {
			appsync_CreateApi(cfg, client)
			return
		}
		if _appsyncCreateApiCache {
			appsync_CreateApiCache(cfg, client)
			return
		}
		if _appsyncCreateApiKey {
			appsync_CreateApiKey(cfg, client)
			return
		}
		if _appsyncCreateChannelNamespace {
			appsync_CreateChannelNamespace(cfg, client)
			return
		}
		if _appsyncCreateDataSource {
			appsync_CreateDataSource(cfg, client)
			return
		}
		if _appsyncCreateDomainName {
			appsync_CreateDomainName(cfg, client)
			return
		}
		if _appsyncCreateFunction {
			appsync_CreateFunction(cfg, client)
			return
		}
		if _appsyncCreateGraphqlApi {
			appsync_CreateGraphqlApi(cfg, client)
			return
		}
		if _appsyncCreateResolver {
			appsync_CreateResolver(cfg, client)
			return
		}
		if _appsyncCreateType {
			appsync_CreateType(cfg, client)
			return
		}
		if _appsyncDeleteApi {
			appsync_DeleteApi(cfg, client)
			return
		}
		if _appsyncDeleteApiCache {
			appsync_DeleteApiCache(cfg, client)
			return
		}
		if _appsyncDeleteApiKey {
			appsync_DeleteApiKey(cfg, client)
			return
		}
		if _appsyncDeleteChannelNamespace {
			appsync_DeleteChannelNamespace(cfg, client)
			return
		}
		if _appsyncDeleteDataSource {
			appsync_DeleteDataSource(cfg, client)
			return
		}
		if _appsyncDeleteDomainName {
			appsync_DeleteDomainName(cfg, client)
			return
		}
		if _appsyncDeleteFunction {
			appsync_DeleteFunction(cfg, client)
			return
		}
		if _appsyncDeleteGraphqlApi {
			appsync_DeleteGraphqlApi(cfg, client)
			return
		}
		if _appsyncDeleteResolver {
			appsync_DeleteResolver(cfg, client)
			return
		}
		if _appsyncDeleteType {
			appsync_DeleteType(cfg, client)
			return
		}
		if _appsyncDisassociateApi {
			appsync_DisassociateApi(cfg, client)
			return
		}
		if _appsyncDisassociateMergedGraphqlApi {
			appsync_DisassociateMergedGraphqlApi(cfg, client)
			return
		}
		if _appsyncDisassociateSourceGraphqlApi {
			appsync_DisassociateSourceGraphqlApi(cfg, client)
			return
		}
		if _appsyncEvaluateCode {
			appsync_EvaluateCode(cfg, client)
			return
		}
		if _appsyncEvaluateMappingTemplate {
			appsync_EvaluateMappingTemplate(cfg, client)
			return
		}
		if _appsyncFlushApiCache {
			appsync_FlushApiCache(cfg, client)
			return
		}
		if _appsyncGetApi {
			appsync_GetApi(cfg, client)
			return
		}
		if _appsyncGetApiAssociation {
			appsync_GetApiAssociation(cfg, client)
			return
		}
		if _appsyncGetApiCache {
			appsync_GetApiCache(cfg, client)
			return
		}
		if _appsyncGetChannelNamespace {
			appsync_GetChannelNamespace(cfg, client)
			return
		}
		if _appsyncGetDataSource {
			appsync_GetDataSource(cfg, client)
			return
		}
		if _appsyncGetDataSourceIntrospection {
			appsync_GetDataSourceIntrospection(cfg, client)
			return
		}
		if _appsyncGetDomainName {
			appsync_GetDomainName(cfg, client)
			return
		}
		if _appsyncGetFunction {
			appsync_GetFunction(cfg, client)
			return
		}
		if _appsyncGetGraphqlApi {
			appsync_GetGraphqlApi(cfg, client)
			return
		}
		if _appsyncGetGraphqlApiEnvironmentVariables {
			appsync_GetGraphqlApiEnvironmentVariables(cfg, client)
			return
		}
		if _appsyncGetIntrospectionSchema {
			appsync_GetIntrospectionSchema(cfg, client)
			return
		}
		if _appsyncGetResolver {
			appsync_GetResolver(cfg, client)
			return
		}
		if _appsyncGetSchemaCreationStatus {
			appsync_GetSchemaCreationStatus(cfg, client)
			return
		}
		if _appsyncGetSourceApiAssociation {
			appsync_GetSourceApiAssociation(cfg, client)
			return
		}
		if _appsyncGetType {
			appsync_GetType(cfg, client)
			return
		}
		if _appsyncListApiKeys {
			appsync_ListApiKeys(cfg, client)
			return
		}
		if _appsyncListApis {
			appsync_ListApis(cfg, client)
			return
		}
		if _appsyncListChannelNamespaces {
			appsync_ListChannelNamespaces(cfg, client)
			return
		}
		if _appsyncListDataSources {
			appsync_ListDataSources(cfg, client)
			return
		}
		if _appsyncListDomainNames {
			appsync_ListDomainNames(cfg, client)
			return
		}
		if _appsyncListFunctions {
			appsync_ListFunctions(cfg, client)
			return
		}
		if _appsyncListGraphqlApis {
			appsync_ListGraphqlApis(cfg, client)
			return
		}
		if _appsyncListResolvers {
			appsync_ListResolvers(cfg, client)
			return
		}
		if _appsyncListResolversByFunction {
			appsync_ListResolversByFunction(cfg, client)
			return
		}
		if _appsyncListSourceApiAssociations {
			appsync_ListSourceApiAssociations(cfg, client)
			return
		}
		if _appsyncListTagsForResource {
			appsync_ListTagsForResource(cfg, client)
			return
		}
		if _appsyncListTypes {
			appsync_ListTypes(cfg, client)
			return
		}
		if _appsyncListTypesByAssociation {
			appsync_ListTypesByAssociation(cfg, client)
			return
		}
		if _appsyncPutGraphqlApiEnvironmentVariables {
			appsync_PutGraphqlApiEnvironmentVariables(cfg, client)
			return
		}
		if _appsyncStartDataSourceIntrospection {
			appsync_StartDataSourceIntrospection(cfg, client)
			return
		}
		if _appsyncStartSchemaCreation {
			appsync_StartSchemaCreation(cfg, client)
			return
		}
		if _appsyncStartSchemaMerge {
			appsync_StartSchemaMerge(cfg, client)
			return
		}
		if _appsyncTagResource {
			appsync_TagResource(cfg, client)
			return
		}
		if _appsyncUntagResource {
			appsync_UntagResource(cfg, client)
			return
		}
		if _appsyncUpdateApi {
			appsync_UpdateApi(cfg, client)
			return
		}
		if _appsyncUpdateApiCache {
			appsync_UpdateApiCache(cfg, client)
			return
		}
		if _appsyncUpdateApiKey {
			appsync_UpdateApiKey(cfg, client)
			return
		}
		if _appsyncUpdateChannelNamespace {
			appsync_UpdateChannelNamespace(cfg, client)
			return
		}
		if _appsyncUpdateDataSource {
			appsync_UpdateDataSource(cfg, client)
			return
		}
		if _appsyncUpdateDomainName {
			appsync_UpdateDomainName(cfg, client)
			return
		}
		if _appsyncUpdateFunction {
			appsync_UpdateFunction(cfg, client)
			return
		}
		if _appsyncUpdateGraphqlApi {
			appsync_UpdateGraphqlApi(cfg, client)
			return
		}
		if _appsyncUpdateResolver {
			appsync_UpdateResolver(cfg, client)
			return
		}
		if _appsyncUpdateSourceApiAssociation {
			appsync_UpdateSourceApiAssociation(cfg, client)
			return
		}
		if _appsyncUpdateType {
			appsync_UpdateType(cfg, client)
			return
		}

	},
}

var (
	_appsyncAssociateApi                      bool
	_appsyncAssociateMergedGraphqlApi         bool
	_appsyncAssociateSourceGraphqlApi         bool
	_appsyncCreateApi                         bool
	_appsyncCreateApiCache                    bool
	_appsyncCreateApiKey                      bool
	_appsyncCreateChannelNamespace            bool
	_appsyncCreateDataSource                  bool
	_appsyncCreateDomainName                  bool
	_appsyncCreateFunction                    bool
	_appsyncCreateGraphqlApi                  bool
	_appsyncCreateResolver                    bool
	_appsyncCreateType                        bool
	_appsyncDeleteApi                         bool
	_appsyncDeleteApiCache                    bool
	_appsyncDeleteApiKey                      bool
	_appsyncDeleteChannelNamespace            bool
	_appsyncDeleteDataSource                  bool
	_appsyncDeleteDomainName                  bool
	_appsyncDeleteFunction                    bool
	_appsyncDeleteGraphqlApi                  bool
	_appsyncDeleteResolver                    bool
	_appsyncDeleteType                        bool
	_appsyncDisassociateApi                   bool
	_appsyncDisassociateMergedGraphqlApi      bool
	_appsyncDisassociateSourceGraphqlApi      bool
	_appsyncEvaluateCode                      bool
	_appsyncEvaluateMappingTemplate           bool
	_appsyncFlushApiCache                     bool
	_appsyncGetApi                            bool
	_appsyncGetApiAssociation                 bool
	_appsyncGetApiCache                       bool
	_appsyncGetChannelNamespace               bool
	_appsyncGetDataSource                     bool
	_appsyncGetDataSourceIntrospection        bool
	_appsyncGetDomainName                     bool
	_appsyncGetFunction                       bool
	_appsyncGetGraphqlApi                     bool
	_appsyncGetGraphqlApiEnvironmentVariables bool
	_appsyncGetIntrospectionSchema            bool
	_appsyncGetResolver                       bool
	_appsyncGetSchemaCreationStatus           bool
	_appsyncGetSourceApiAssociation           bool
	_appsyncGetType                           bool
	_appsyncListApiKeys                       bool
	_appsyncListApis                          bool
	_appsyncListChannelNamespaces             bool
	_appsyncListDataSources                   bool
	_appsyncListDomainNames                   bool
	_appsyncListFunctions                     bool
	_appsyncListGraphqlApis                   bool
	_appsyncListResolvers                     bool
	_appsyncListResolversByFunction           bool
	_appsyncListSourceApiAssociations         bool
	_appsyncListTagsForResource               bool
	_appsyncListTypes                         bool
	_appsyncListTypesByAssociation            bool
	_appsyncPutGraphqlApiEnvironmentVariables bool
	_appsyncStartDataSourceIntrospection      bool
	_appsyncStartSchemaCreation               bool
	_appsyncStartSchemaMerge                  bool
	_appsyncTagResource                       bool
	_appsyncUntagResource                     bool
	_appsyncUpdateApi                         bool
	_appsyncUpdateApiCache                    bool
	_appsyncUpdateApiKey                      bool
	_appsyncUpdateChannelNamespace            bool
	_appsyncUpdateDataSource                  bool
	_appsyncUpdateDomainName                  bool
	_appsyncUpdateFunction                    bool
	_appsyncUpdateGraphqlApi                  bool
	_appsyncUpdateResolver                    bool
	_appsyncUpdateSourceApiAssociation        bool
	_appsyncUpdateType                        bool

	_appsyncAdditionalAuthenticationProviders string
	_appsyncApiCachingBehavior                string
	_appsyncApiId                             string
	_appsyncApiType                           string
	_appsyncAssociationId                     string
	_appsyncAtRestEncryptionEnabled           string
	_appsyncAuthenticationType                string
	_appsyncCachingConfig                     string
	_appsyncCertificateArn                    string
	_appsyncCode                              string
	_appsyncCodeHandlers                      string
	_appsyncContext                           string
	_appsyncDataSourceName                    string
	_appsyncDefinition                        string
	_appsyncDescription                       string
	_appsyncDomainName                        string
	_appsyncDynamodbConfig                    string
	_appsyncElasticsearchConfig               string
	_appsyncEnhancedMetricsConfig             string
	_appsyncEnvironmentVariables              string
	_appsyncEventBridgeConfig                 string
	_appsyncEventConfig                       string
	_appsyncExpires                           string
	_appsyncFieldName                         string
	_appsyncFormat                            string
	_appsyncFunction                          string
	_appsyncFunctionId                        string
	_appsyncFunctionVersion                   string
	_appsyncHandlerConfigs                    string
	_appsyncHealthMetricsConfig               string
	_appsyncHttpConfig                        string
	_appsyncId                                string
	_appsyncIncludeDirectives                 string
	_appsyncIncludeModelsSDL                  string
	_appsyncIntrospectionConfig               string
	_appsyncIntrospectionId                   string
	_appsyncKind                              string
	_appsyncLambdaAuthorizerConfig            string
	_appsyncLambdaConfig                      string
	_appsyncLogConfig                         string
	_appsyncMaxBatchSize                      string
	_appsyncMaxResults                        string
	_appsyncMergedApiExecutionRoleArn         string
	_appsyncMergedApiIdentifier               string
	_appsyncMetricsConfig                     string
	_appsyncName                              string
	_appsyncNextToken                         string
	_appsyncOpenIDConnectConfig               string
	_appsyncOpenSearchServiceConfig           string
	_appsyncOwner                             string
	_appsyncOwnerContact                      string
	_appsyncPipelineConfig                    string
	_appsyncPublishAuthModes                  string
	_appsyncQueryDepthLimit                   string
	_appsyncRdsDataApiConfig                  string
	_appsyncRelationalDatabaseConfig          string
	_appsyncRequestMappingTemplate            string
	_appsyncResolverCountLimit                string
	_appsyncResourceArn                       string
	_appsyncResponseMappingTemplate           string
	_appsyncRuntime                           string
	_appsyncServiceRoleArn                    string
	_appsyncSourceApiAssociationConfig        string
	_appsyncSourceApiIdentifier               string
	_appsyncSubscribeAuthModes                string
	_appsyncSyncConfig                        string
	_appsyncTagKeys                           []string
	_appsyncTags                              string
	_appsyncTemplate                          string
	_appsyncTransitEncryptionEnabled          string
	_appsyncTtl                               string
	_appsyncType                              string
	_appsyncTypeName                          string
	_appsyncUserPoolConfig                    string
	_appsyncVisibility                        string
	_appsyncXrayEnabled                       string
)

// Maps an endpoint to your custom domain.
func appsync_AssociateApi(cfg aws.Config, client *appsync.Client) {
	input := &appsync.AssociateApiInput{
		// ApiId: *string, // Required
		// DomainName: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncDomainName) > 0 {
		input.DomainName = aws.String(_appsyncDomainName)
	}

	if resp, err := client.AssociateApi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an association between a Merged API and source API using the source
// API's identifier.
func appsync_AssociateMergedGraphqlApi(cfg aws.Config, client *appsync.Client) {
	input := &appsync.AssociateMergedGraphqlApiInput{
		// MergedApiIdentifier: *string, // Required
		// SourceApiIdentifier: *string, // Required
	}

	if len(_appsyncMergedApiIdentifier) > 0 {
		input.MergedApiIdentifier = aws.String(_appsyncMergedApiIdentifier)
	}
	if len(_appsyncSourceApiIdentifier) > 0 {
		input.SourceApiIdentifier = aws.String(_appsyncSourceApiIdentifier)
	}
	if len(_appsyncDescription) > 0 {
		input.Description = aws.String(_appsyncDescription)
	}
	if len(_appsyncSourceApiAssociationConfig) > 0 {
		if err := assignInputField(input, "SourceApiAssociationConfig", _appsyncSourceApiAssociationConfig); err != nil {
			log.Errorf("invalid --source-api-association-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateMergedGraphqlApi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an association between a Merged API and source API using the Merged
// API's identifier.
func appsync_AssociateSourceGraphqlApi(cfg aws.Config, client *appsync.Client) {
	input := &appsync.AssociateSourceGraphqlApiInput{
		// MergedApiIdentifier: *string, // Required
		// SourceApiIdentifier: *string, // Required
	}

	if len(_appsyncMergedApiIdentifier) > 0 {
		input.MergedApiIdentifier = aws.String(_appsyncMergedApiIdentifier)
	}
	if len(_appsyncSourceApiIdentifier) > 0 {
		input.SourceApiIdentifier = aws.String(_appsyncSourceApiIdentifier)
	}
	if len(_appsyncDescription) > 0 {
		input.Description = aws.String(_appsyncDescription)
	}
	if len(_appsyncSourceApiAssociationConfig) > 0 {
		if err := assignInputField(input, "SourceApiAssociationConfig", _appsyncSourceApiAssociationConfig); err != nil {
			log.Errorf("invalid --source-api-association-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateSourceGraphqlApi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Api object. Use this operation to create an AppSync API with your
// preferred configuration, such as an Event API that provides real-time message
// publishing and message subscriptions over WebSockets.
func appsync_CreateApi(cfg aws.Config, client *appsync.Client) {
	input := &appsync.CreateApiInput{
		// EventConfig: *types.EventConfig, // Required
		// Name: *string, // Required
	}

	if len(_appsyncEventConfig) > 0 {
		if err := assignInputField(input, "EventConfig", _appsyncEventConfig); err != nil {
			log.Errorf("invalid --event-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncName) > 0 {
		input.Name = aws.String(_appsyncName)
	}
	if len(_appsyncOwnerContact) > 0 {
		input.OwnerContact = aws.String(_appsyncOwnerContact)
	}
	if len(_appsyncTags) > 0 {
		if err := assignInputField(input, "Tags", _appsyncTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateApi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a cache for the GraphQL API.
func appsync_CreateApiCache(cfg aws.Config, client *appsync.Client) {
	input := &appsync.CreateApiCacheInput{
		// ApiCachingBehavior: types.ApiCachingBehavior, // Required
		// ApiId: *string, // Required
		// Ttl: int64, // Required
		// Type: types.ApiCacheType, // Required
	}

	if len(_appsyncApiCachingBehavior) > 0 {
		if err := assignInputField(input, "ApiCachingBehavior", _appsyncApiCachingBehavior); err != nil {
			log.Errorf("invalid --api-caching-behavior: %s", err.Error())
			return
		}
	}
	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncTtl) > 0 {
		if err := assignInputField(input, "Ttl", _appsyncTtl); err != nil {
			log.Errorf("invalid --ttl: %s", err.Error())
			return
		}
	}
	if len(_appsyncType) > 0 {
		if err := assignInputField(input, "Type", _appsyncType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_appsyncAtRestEncryptionEnabled) > 0 {
		if err := assignInputField(input, "AtRestEncryptionEnabled", _appsyncAtRestEncryptionEnabled); err != nil {
			log.Errorf("invalid --at-rest-encryption-enabled: %s", err.Error())
			return
		}
	}
	if len(_appsyncHealthMetricsConfig) > 0 {
		if err := assignInputField(input, "HealthMetricsConfig", _appsyncHealthMetricsConfig); err != nil {
			log.Errorf("invalid --health-metrics-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncTransitEncryptionEnabled) > 0 {
		if err := assignInputField(input, "TransitEncryptionEnabled", _appsyncTransitEncryptionEnabled); err != nil {
			log.Errorf("invalid --transit-encryption-enabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateApiCache(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a unique key that you can distribute to clients who invoke your API.
func appsync_CreateApiKey(cfg aws.Config, client *appsync.Client) {
	input := &appsync.CreateApiKeyInput{
		// ApiId: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncDescription) > 0 {
		input.Description = aws.String(_appsyncDescription)
	}
	if len(_appsyncExpires) > 0 {
		if err := assignInputField(input, "Expires", _appsyncExpires); err != nil {
			log.Errorf("invalid --expires: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateApiKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a ChannelNamespace for an Api .
func appsync_CreateChannelNamespace(cfg aws.Config, client *appsync.Client) {
	input := &appsync.CreateChannelNamespaceInput{
		// ApiId: *string, // Required
		// Name: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncName) > 0 {
		input.Name = aws.String(_appsyncName)
	}
	if len(_appsyncCodeHandlers) > 0 {
		input.CodeHandlers = aws.String(_appsyncCodeHandlers)
	}
	if len(_appsyncHandlerConfigs) > 0 {
		if err := assignInputField(input, "HandlerConfigs", _appsyncHandlerConfigs); err != nil {
			log.Errorf("invalid --handler-configs: %s", err.Error())
			return
		}
	}
	if len(_appsyncPublishAuthModes) > 0 {
		if err := assignInputField(input, "PublishAuthModes", _appsyncPublishAuthModes); err != nil {
			log.Errorf("invalid --publish-auth-modes: %s", err.Error())
			return
		}
	}
	if len(_appsyncSubscribeAuthModes) > 0 {
		if err := assignInputField(input, "SubscribeAuthModes", _appsyncSubscribeAuthModes); err != nil {
			log.Errorf("invalid --subscribe-auth-modes: %s", err.Error())
			return
		}
	}
	if len(_appsyncTags) > 0 {
		if err := assignInputField(input, "Tags", _appsyncTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateChannelNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a DataSource object.
func appsync_CreateDataSource(cfg aws.Config, client *appsync.Client) {
	input := &appsync.CreateDataSourceInput{
		// ApiId: *string, // Required
		// Name: *string, // Required
		// Type: types.DataSourceType, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncName) > 0 {
		input.Name = aws.String(_appsyncName)
	}
	if len(_appsyncType) > 0 {
		if err := assignInputField(input, "Type", _appsyncType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_appsyncDescription) > 0 {
		input.Description = aws.String(_appsyncDescription)
	}
	if len(_appsyncDynamodbConfig) > 0 {
		if err := assignInputField(input, "DynamodbConfig", _appsyncDynamodbConfig); err != nil {
			log.Errorf("invalid --dynamodb-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncElasticsearchConfig) > 0 {
		if err := assignInputField(input, "ElasticsearchConfig", _appsyncElasticsearchConfig); err != nil {
			log.Errorf("invalid --elasticsearch-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncEventBridgeConfig) > 0 {
		if err := assignInputField(input, "EventBridgeConfig", _appsyncEventBridgeConfig); err != nil {
			log.Errorf("invalid --event-bridge-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncHttpConfig) > 0 {
		if err := assignInputField(input, "HttpConfig", _appsyncHttpConfig); err != nil {
			log.Errorf("invalid --http-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncLambdaConfig) > 0 {
		if err := assignInputField(input, "LambdaConfig", _appsyncLambdaConfig); err != nil {
			log.Errorf("invalid --lambda-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncMetricsConfig) > 0 {
		if err := assignInputField(input, "MetricsConfig", _appsyncMetricsConfig); err != nil {
			log.Errorf("invalid --metrics-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncOpenSearchServiceConfig) > 0 {
		if err := assignInputField(input, "OpenSearchServiceConfig", _appsyncOpenSearchServiceConfig); err != nil {
			log.Errorf("invalid --open-search-service-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncRelationalDatabaseConfig) > 0 {
		if err := assignInputField(input, "RelationalDatabaseConfig", _appsyncRelationalDatabaseConfig); err != nil {
			log.Errorf("invalid --relational-database-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncServiceRoleArn) > 0 {
		input.ServiceRoleArn = aws.String(_appsyncServiceRoleArn)
	}

	if resp, err := client.CreateDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a custom DomainName object.
func appsync_CreateDomainName(cfg aws.Config, client *appsync.Client) {
	input := &appsync.CreateDomainNameInput{
		// CertificateArn: *string, // Required
		// DomainName: *string, // Required
	}

	if len(_appsyncCertificateArn) > 0 {
		input.CertificateArn = aws.String(_appsyncCertificateArn)
	}
	if len(_appsyncDomainName) > 0 {
		input.DomainName = aws.String(_appsyncDomainName)
	}
	if len(_appsyncDescription) > 0 {
		input.Description = aws.String(_appsyncDescription)
	}
	if len(_appsyncTags) > 0 {
		if err := assignInputField(input, "Tags", _appsyncTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDomainName(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Function object.
// A function is a reusable entity. You can use multiple functions to compose the
// resolver logic.
func appsync_CreateFunction(cfg aws.Config, client *appsync.Client) {
	input := &appsync.CreateFunctionInput{
		// ApiId: *string, // Required
		// DataSourceName: *string, // Required
		// Name: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncDataSourceName) > 0 {
		input.DataSourceName = aws.String(_appsyncDataSourceName)
	}
	if len(_appsyncName) > 0 {
		input.Name = aws.String(_appsyncName)
	}
	if len(_appsyncCode) > 0 {
		input.Code = aws.String(_appsyncCode)
	}
	if len(_appsyncDescription) > 0 {
		input.Description = aws.String(_appsyncDescription)
	}
	if len(_appsyncFunctionVersion) > 0 {
		input.FunctionVersion = aws.String(_appsyncFunctionVersion)
	}
	if len(_appsyncMaxBatchSize) > 0 {
		if err := assignInputField(input, "MaxBatchSize", _appsyncMaxBatchSize); err != nil {
			log.Errorf("invalid --max-batch-size: %s", err.Error())
			return
		}
	}
	if len(_appsyncRequestMappingTemplate) > 0 {
		input.RequestMappingTemplate = aws.String(_appsyncRequestMappingTemplate)
	}
	if len(_appsyncResponseMappingTemplate) > 0 {
		input.ResponseMappingTemplate = aws.String(_appsyncResponseMappingTemplate)
	}
	if len(_appsyncRuntime) > 0 {
		if err := assignInputField(input, "Runtime", _appsyncRuntime); err != nil {
			log.Errorf("invalid --runtime: %s", err.Error())
			return
		}
	}
	if len(_appsyncSyncConfig) > 0 {
		if err := assignInputField(input, "SyncConfig", _appsyncSyncConfig); err != nil {
			log.Errorf("invalid --sync-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFunction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a GraphqlApi object.
func appsync_CreateGraphqlApi(cfg aws.Config, client *appsync.Client) {
	input := &appsync.CreateGraphqlApiInput{
		// AuthenticationType: types.AuthenticationType, // Required
		// Name: *string, // Required
	}

	if len(_appsyncAuthenticationType) > 0 {
		if err := assignInputField(input, "AuthenticationType", _appsyncAuthenticationType); err != nil {
			log.Errorf("invalid --authentication-type: %s", err.Error())
			return
		}
	}
	if len(_appsyncName) > 0 {
		input.Name = aws.String(_appsyncName)
	}
	if len(_appsyncAdditionalAuthenticationProviders) > 0 {
		if err := assignInputField(input, "AdditionalAuthenticationProviders", _appsyncAdditionalAuthenticationProviders); err != nil {
			log.Errorf("invalid --additional-authentication-providers: %s", err.Error())
			return
		}
	}
	if len(_appsyncApiType) > 0 {
		if err := assignInputField(input, "ApiType", _appsyncApiType); err != nil {
			log.Errorf("invalid --api-type: %s", err.Error())
			return
		}
	}
	if len(_appsyncEnhancedMetricsConfig) > 0 {
		if err := assignInputField(input, "EnhancedMetricsConfig", _appsyncEnhancedMetricsConfig); err != nil {
			log.Errorf("invalid --enhanced-metrics-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncIntrospectionConfig) > 0 {
		if err := assignInputField(input, "IntrospectionConfig", _appsyncIntrospectionConfig); err != nil {
			log.Errorf("invalid --introspection-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncLambdaAuthorizerConfig) > 0 {
		if err := assignInputField(input, "LambdaAuthorizerConfig", _appsyncLambdaAuthorizerConfig); err != nil {
			log.Errorf("invalid --lambda-authorizer-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncLogConfig) > 0 {
		if err := assignInputField(input, "LogConfig", _appsyncLogConfig); err != nil {
			log.Errorf("invalid --log-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncMergedApiExecutionRoleArn) > 0 {
		input.MergedApiExecutionRoleArn = aws.String(_appsyncMergedApiExecutionRoleArn)
	}
	if len(_appsyncOpenIDConnectConfig) > 0 {
		if err := assignInputField(input, "OpenIDConnectConfig", _appsyncOpenIDConnectConfig); err != nil {
			log.Errorf("invalid --open-id-connect-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncOwnerContact) > 0 {
		input.OwnerContact = aws.String(_appsyncOwnerContact)
	}
	if len(_appsyncQueryDepthLimit) > 0 {
		if err := assignInputField(input, "QueryDepthLimit", _appsyncQueryDepthLimit); err != nil {
			log.Errorf("invalid --query-depth-limit: %s", err.Error())
			return
		}
	}
	if len(_appsyncResolverCountLimit) > 0 {
		if err := assignInputField(input, "ResolverCountLimit", _appsyncResolverCountLimit); err != nil {
			log.Errorf("invalid --resolver-count-limit: %s", err.Error())
			return
		}
	}
	if len(_appsyncTags) > 0 {
		if err := assignInputField(input, "Tags", _appsyncTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_appsyncUserPoolConfig) > 0 {
		if err := assignInputField(input, "UserPoolConfig", _appsyncUserPoolConfig); err != nil {
			log.Errorf("invalid --user-pool-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncVisibility) > 0 {
		if err := assignInputField(input, "Visibility", _appsyncVisibility); err != nil {
			log.Errorf("invalid --visibility: %s", err.Error())
			return
		}
	}
	if len(_appsyncXrayEnabled) > 0 {
		if err := assignInputField(input, "XrayEnabled", _appsyncXrayEnabled); err != nil {
			log.Errorf("invalid --xray-enabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGraphqlApi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Resolver object.
// A resolver converts incoming requests into a format that a data source can
// understand, and converts the data source's responses into GraphQL.
func appsync_CreateResolver(cfg aws.Config, client *appsync.Client) {
	input := &appsync.CreateResolverInput{
		// ApiId: *string, // Required
		// FieldName: *string, // Required
		// TypeName: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncFieldName) > 0 {
		input.FieldName = aws.String(_appsyncFieldName)
	}
	if len(_appsyncTypeName) > 0 {
		input.TypeName = aws.String(_appsyncTypeName)
	}
	if len(_appsyncCachingConfig) > 0 {
		if err := assignInputField(input, "CachingConfig", _appsyncCachingConfig); err != nil {
			log.Errorf("invalid --caching-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncCode) > 0 {
		input.Code = aws.String(_appsyncCode)
	}
	if len(_appsyncDataSourceName) > 0 {
		input.DataSourceName = aws.String(_appsyncDataSourceName)
	}
	if len(_appsyncKind) > 0 {
		if err := assignInputField(input, "Kind", _appsyncKind); err != nil {
			log.Errorf("invalid --kind: %s", err.Error())
			return
		}
	}
	if len(_appsyncMaxBatchSize) > 0 {
		if err := assignInputField(input, "MaxBatchSize", _appsyncMaxBatchSize); err != nil {
			log.Errorf("invalid --max-batch-size: %s", err.Error())
			return
		}
	}
	if len(_appsyncMetricsConfig) > 0 {
		if err := assignInputField(input, "MetricsConfig", _appsyncMetricsConfig); err != nil {
			log.Errorf("invalid --metrics-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncPipelineConfig) > 0 {
		if err := assignInputField(input, "PipelineConfig", _appsyncPipelineConfig); err != nil {
			log.Errorf("invalid --pipeline-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncRequestMappingTemplate) > 0 {
		input.RequestMappingTemplate = aws.String(_appsyncRequestMappingTemplate)
	}
	if len(_appsyncResponseMappingTemplate) > 0 {
		input.ResponseMappingTemplate = aws.String(_appsyncResponseMappingTemplate)
	}
	if len(_appsyncRuntime) > 0 {
		if err := assignInputField(input, "Runtime", _appsyncRuntime); err != nil {
			log.Errorf("invalid --runtime: %s", err.Error())
			return
		}
	}
	if len(_appsyncSyncConfig) > 0 {
		if err := assignInputField(input, "SyncConfig", _appsyncSyncConfig); err != nil {
			log.Errorf("invalid --sync-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateResolver(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Type object.
func appsync_CreateType(cfg aws.Config, client *appsync.Client) {
	input := &appsync.CreateTypeInput{
		// ApiId: *string, // Required
		// Definition: *string, // Required
		// Format: types.TypeDefinitionFormat, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncDefinition) > 0 {
		input.Definition = aws.String(_appsyncDefinition)
	}
	if len(_appsyncFormat) > 0 {
		if err := assignInputField(input, "Format", _appsyncFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Api object
func appsync_DeleteApi(cfg aws.Config, client *appsync.Client) {
	input := &appsync.DeleteApiInput{
		// ApiId: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}

	if resp, err := client.DeleteApi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an ApiCache object.
func appsync_DeleteApiCache(cfg aws.Config, client *appsync.Client) {
	input := &appsync.DeleteApiCacheInput{
		// ApiId: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}

	if resp, err := client.DeleteApiCache(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an API key.
func appsync_DeleteApiKey(cfg aws.Config, client *appsync.Client) {
	input := &appsync.DeleteApiKeyInput{
		// ApiId: *string, // Required
		// Id: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncId) > 0 {
		input.Id = aws.String(_appsyncId)
	}

	if resp, err := client.DeleteApiKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a ChannelNamespace .
func appsync_DeleteChannelNamespace(cfg aws.Config, client *appsync.Client) {
	input := &appsync.DeleteChannelNamespaceInput{
		// ApiId: *string, // Required
		// Name: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncName) > 0 {
		input.Name = aws.String(_appsyncName)
	}

	if resp, err := client.DeleteChannelNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a DataSource object.
func appsync_DeleteDataSource(cfg aws.Config, client *appsync.Client) {
	input := &appsync.DeleteDataSourceInput{
		// ApiId: *string, // Required
		// Name: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncName) > 0 {
		input.Name = aws.String(_appsyncName)
	}

	if resp, err := client.DeleteDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a custom DomainName object.
func appsync_DeleteDomainName(cfg aws.Config, client *appsync.Client) {
	input := &appsync.DeleteDomainNameInput{
		// DomainName: *string, // Required
	}

	if len(_appsyncDomainName) > 0 {
		input.DomainName = aws.String(_appsyncDomainName)
	}

	if resp, err := client.DeleteDomainName(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Function .
func appsync_DeleteFunction(cfg aws.Config, client *appsync.Client) {
	input := &appsync.DeleteFunctionInput{
		// ApiId: *string, // Required
		// FunctionId: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncFunctionId) > 0 {
		input.FunctionId = aws.String(_appsyncFunctionId)
	}

	if resp, err := client.DeleteFunction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a GraphqlApi object.
func appsync_DeleteGraphqlApi(cfg aws.Config, client *appsync.Client) {
	input := &appsync.DeleteGraphqlApiInput{
		// ApiId: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}

	if resp, err := client.DeleteGraphqlApi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Resolver object.
func appsync_DeleteResolver(cfg aws.Config, client *appsync.Client) {
	input := &appsync.DeleteResolverInput{
		// ApiId: *string, // Required
		// FieldName: *string, // Required
		// TypeName: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncFieldName) > 0 {
		input.FieldName = aws.String(_appsyncFieldName)
	}
	if len(_appsyncTypeName) > 0 {
		input.TypeName = aws.String(_appsyncTypeName)
	}

	if resp, err := client.DeleteResolver(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Type object.
func appsync_DeleteType(cfg aws.Config, client *appsync.Client) {
	input := &appsync.DeleteTypeInput{
		// ApiId: *string, // Required
		// TypeName: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncTypeName) > 0 {
		input.TypeName = aws.String(_appsyncTypeName)
	}

	if resp, err := client.DeleteType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes an ApiAssociation object from a custom domain.
func appsync_DisassociateApi(cfg aws.Config, client *appsync.Client) {
	input := &appsync.DisassociateApiInput{
		// DomainName: *string, // Required
	}

	if len(_appsyncDomainName) > 0 {
		input.DomainName = aws.String(_appsyncDomainName)
	}

	if resp, err := client.DisassociateApi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an association between a Merged API and source API using the source
// API's identifier and the association ID.
func appsync_DisassociateMergedGraphqlApi(cfg aws.Config, client *appsync.Client) {
	input := &appsync.DisassociateMergedGraphqlApiInput{
		// AssociationId: *string, // Required
		// SourceApiIdentifier: *string, // Required
	}

	if len(_appsyncAssociationId) > 0 {
		input.AssociationId = aws.String(_appsyncAssociationId)
	}
	if len(_appsyncSourceApiIdentifier) > 0 {
		input.SourceApiIdentifier = aws.String(_appsyncSourceApiIdentifier)
	}

	if resp, err := client.DisassociateMergedGraphqlApi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an association between a Merged API and source API using the Merged
// API's identifier and the association ID.
func appsync_DisassociateSourceGraphqlApi(cfg aws.Config, client *appsync.Client) {
	input := &appsync.DisassociateSourceGraphqlApiInput{
		// AssociationId: *string, // Required
		// MergedApiIdentifier: *string, // Required
	}

	if len(_appsyncAssociationId) > 0 {
		input.AssociationId = aws.String(_appsyncAssociationId)
	}
	if len(_appsyncMergedApiIdentifier) > 0 {
		input.MergedApiIdentifier = aws.String(_appsyncMergedApiIdentifier)
	}

	if resp, err := client.DisassociateSourceGraphqlApi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Evaluates the given code and returns the response. The code definition
// requirements depend on the specified runtime. For APPSYNC_JS runtimes, the code
// defines the request and response functions. The request function takes the
// incoming request after a GraphQL operation is parsed and converts it into a
// request configuration for the selected data source operation. The response
// function interprets responses from the data source and maps it to the shape of
// the GraphQL field output type.
func appsync_EvaluateCode(cfg aws.Config, client *appsync.Client) {
	input := &appsync.EvaluateCodeInput{
		// Code: *string, // Required
		// Context: *string, // Required
		// Runtime: *types.AppSyncRuntime, // Required
	}

	if len(_appsyncCode) > 0 {
		input.Code = aws.String(_appsyncCode)
	}
	if len(_appsyncContext) > 0 {
		input.Context = aws.String(_appsyncContext)
	}
	if len(_appsyncRuntime) > 0 {
		if err := assignInputField(input, "Runtime", _appsyncRuntime); err != nil {
			log.Errorf("invalid --runtime: %s", err.Error())
			return
		}
	}
	if len(_appsyncFunction) > 0 {
		input.Function = aws.String(_appsyncFunction)
	}

	if resp, err := client.EvaluateCode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Evaluates a given template and returns the response. The mapping template can
// be a request or response template.
//
// Request templates take the incoming request after a GraphQL operation is parsed
// and convert it into a request configuration for the selected data source
// operation. Response templates interpret responses from the data source and map
// it to the shape of the GraphQL field output type.
//
// Mapping templates are written in the Apache Velocity Template Language (VTL).
func appsync_EvaluateMappingTemplate(cfg aws.Config, client *appsync.Client) {
	input := &appsync.EvaluateMappingTemplateInput{
		// Context: *string, // Required
		// Template: *string, // Required
	}

	if len(_appsyncContext) > 0 {
		input.Context = aws.String(_appsyncContext)
	}
	if len(_appsyncTemplate) > 0 {
		input.Template = aws.String(_appsyncTemplate)
	}

	if resp, err := client.EvaluateMappingTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Flushes an ApiCache object.
func appsync_FlushApiCache(cfg aws.Config, client *appsync.Client) {
	input := &appsync.FlushApiCacheInput{
		// ApiId: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}

	if resp, err := client.FlushApiCache(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves an Api object.
func appsync_GetApi(cfg aws.Config, client *appsync.Client) {
	input := &appsync.GetApiInput{
		// ApiId: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}

	if resp, err := client.GetApi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves an ApiAssociation object.
func appsync_GetApiAssociation(cfg aws.Config, client *appsync.Client) {
	input := &appsync.GetApiAssociationInput{
		// DomainName: *string, // Required
	}

	if len(_appsyncDomainName) > 0 {
		input.DomainName = aws.String(_appsyncDomainName)
	}

	if resp, err := client.GetApiAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves an ApiCache object.
func appsync_GetApiCache(cfg aws.Config, client *appsync.Client) {
	input := &appsync.GetApiCacheInput{
		// ApiId: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}

	if resp, err := client.GetApiCache(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the channel namespace for a specified Api .
func appsync_GetChannelNamespace(cfg aws.Config, client *appsync.Client) {
	input := &appsync.GetChannelNamespaceInput{
		// ApiId: *string, // Required
		// Name: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncName) > 0 {
		input.Name = aws.String(_appsyncName)
	}

	if resp, err := client.GetChannelNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a DataSource object.
func appsync_GetDataSource(cfg aws.Config, client *appsync.Client) {
	input := &appsync.GetDataSourceInput{
		// ApiId: *string, // Required
		// Name: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncName) > 0 {
		input.Name = aws.String(_appsyncName)
	}

	if resp, err := client.GetDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the record of an existing introspection. If the retrieval is
// successful, the result of the instrospection will also be returned. If the
// retrieval fails the operation, an error message will be returned instead.
func appsync_GetDataSourceIntrospection(cfg aws.Config, client *appsync.Client) {
	input := &appsync.GetDataSourceIntrospectionInput{
		// IntrospectionId: *string, // Required
	}

	if len(_appsyncIntrospectionId) > 0 {
		input.IntrospectionId = aws.String(_appsyncIntrospectionId)
	}
	if len(_appsyncIncludeModelsSDL) > 0 {
		if err := assignInputField(input, "IncludeModelsSDL", _appsyncIncludeModelsSDL); err != nil {
			log.Errorf("invalid --include-models-sdl: %s", err.Error())
			return
		}
	}
	if len(_appsyncMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appsyncMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appsyncNextToken) > 0 {
		input.NextToken = aws.String(_appsyncNextToken)
	}

	if resp, err := client.GetDataSourceIntrospection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a custom DomainName object.
func appsync_GetDomainName(cfg aws.Config, client *appsync.Client) {
	input := &appsync.GetDomainNameInput{
		// DomainName: *string, // Required
	}

	if len(_appsyncDomainName) > 0 {
		input.DomainName = aws.String(_appsyncDomainName)
	}

	if resp, err := client.GetDomainName(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get a Function .
func appsync_GetFunction(cfg aws.Config, client *appsync.Client) {
	input := &appsync.GetFunctionInput{
		// ApiId: *string, // Required
		// FunctionId: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncFunctionId) > 0 {
		input.FunctionId = aws.String(_appsyncFunctionId)
	}

	if resp, err := client.GetFunction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a GraphqlApi object.
func appsync_GetGraphqlApi(cfg aws.Config, client *appsync.Client) {
	input := &appsync.GetGraphqlApiInput{
		// ApiId: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}

	if resp, err := client.GetGraphqlApi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the list of environmental variable key-value pairs associated with an
// API by its ID value.
func appsync_GetGraphqlApiEnvironmentVariables(cfg aws.Config, client *appsync.Client) {
	input := &appsync.GetGraphqlApiEnvironmentVariablesInput{
		// ApiId: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}

	if resp, err := client.GetGraphqlApiEnvironmentVariables(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the introspection schema for a GraphQL API.
func appsync_GetIntrospectionSchema(cfg aws.Config, client *appsync.Client) {
	input := &appsync.GetIntrospectionSchemaInput{
		// ApiId: *string, // Required
		// Format: types.OutputType, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncFormat) > 0 {
		if err := assignInputField(input, "Format", _appsyncFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_appsyncIncludeDirectives) > 0 {
		if err := assignInputField(input, "IncludeDirectives", _appsyncIncludeDirectives); err != nil {
			log.Errorf("invalid --include-directives: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetIntrospectionSchema(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a Resolver object.
func appsync_GetResolver(cfg aws.Config, client *appsync.Client) {
	input := &appsync.GetResolverInput{
		// ApiId: *string, // Required
		// FieldName: *string, // Required
		// TypeName: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncFieldName) > 0 {
		input.FieldName = aws.String(_appsyncFieldName)
	}
	if len(_appsyncTypeName) > 0 {
		input.TypeName = aws.String(_appsyncTypeName)
	}

	if resp, err := client.GetResolver(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the current status of a schema creation operation.
func appsync_GetSchemaCreationStatus(cfg aws.Config, client *appsync.Client) {
	input := &appsync.GetSchemaCreationStatusInput{
		// ApiId: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}

	if resp, err := client.GetSchemaCreationStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a SourceApiAssociation object.
func appsync_GetSourceApiAssociation(cfg aws.Config, client *appsync.Client) {
	input := &appsync.GetSourceApiAssociationInput{
		// AssociationId: *string, // Required
		// MergedApiIdentifier: *string, // Required
	}

	if len(_appsyncAssociationId) > 0 {
		input.AssociationId = aws.String(_appsyncAssociationId)
	}
	if len(_appsyncMergedApiIdentifier) > 0 {
		input.MergedApiIdentifier = aws.String(_appsyncMergedApiIdentifier)
	}

	if resp, err := client.GetSourceApiAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a Type object.
func appsync_GetType(cfg aws.Config, client *appsync.Client) {
	input := &appsync.GetTypeInput{
		// ApiId: *string, // Required
		// Format: types.TypeDefinitionFormat, // Required
		// TypeName: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncFormat) > 0 {
		if err := assignInputField(input, "Format", _appsyncFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_appsyncTypeName) > 0 {
		input.TypeName = aws.String(_appsyncTypeName)
	}

	if resp, err := client.GetType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the API keys for a given API.
// API keys are deleted automatically 60 days after they expire. However, they may
// still be included in the response until they have actually been deleted. You can
// safely call DeleteApiKey to manually delete a key before it's automatically
// deleted.
func appsync_ListApiKeys(cfg aws.Config, client *appsync.Client) {
	input := &appsync.ListApiKeysInput{
		// ApiId: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appsyncMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appsyncNextToken) > 0 {
		input.NextToken = aws.String(_appsyncNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApiKeys(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appsync.ListApiKeysOutput
	p := appsync.NewListApiKeysPaginator(client, input)
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

// Lists the APIs in your AppSync account.
// ListApis returns only the high level API details. For more detailed information
// about an API, use GetApi .
func appsync_ListApis(cfg aws.Config, client *appsync.Client) {
	input := &appsync.ListApisInput{}

	if len(_appsyncMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appsyncMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appsyncNextToken) > 0 {
		input.NextToken = aws.String(_appsyncNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApis(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appsync.ListApisOutput
	p := appsync.NewListApisPaginator(client, input)
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

// Lists the channel namespaces for a specified Api .
// ListChannelNamespaces returns only high level details for the channel
// namespace. To retrieve code handlers, use GetChannelNamespace .
func appsync_ListChannelNamespaces(cfg aws.Config, client *appsync.Client) {
	input := &appsync.ListChannelNamespacesInput{
		// ApiId: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appsyncMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appsyncNextToken) > 0 {
		input.NextToken = aws.String(_appsyncNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListChannelNamespaces(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appsync.ListChannelNamespacesOutput
	p := appsync.NewListChannelNamespacesPaginator(client, input)
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

// Lists the data sources for a given API.
func appsync_ListDataSources(cfg aws.Config, client *appsync.Client) {
	input := &appsync.ListDataSourcesInput{
		// ApiId: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appsyncMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appsyncNextToken) > 0 {
		input.NextToken = aws.String(_appsyncNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDataSources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appsync.ListDataSourcesOutput
	p := appsync.NewListDataSourcesPaginator(client, input)
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

// Lists multiple custom domain names.
func appsync_ListDomainNames(cfg aws.Config, client *appsync.Client) {
	input := &appsync.ListDomainNamesInput{}

	if len(_appsyncMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appsyncMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appsyncNextToken) > 0 {
		input.NextToken = aws.String(_appsyncNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDomainNames(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appsync.ListDomainNamesOutput
	p := appsync.NewListDomainNamesPaginator(client, input)
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

// List multiple functions.
func appsync_ListFunctions(cfg aws.Config, client *appsync.Client) {
	input := &appsync.ListFunctionsInput{
		// ApiId: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appsyncMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appsyncNextToken) > 0 {
		input.NextToken = aws.String(_appsyncNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFunctions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appsync.ListFunctionsOutput
	p := appsync.NewListFunctionsPaginator(client, input)
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

// Lists your GraphQL APIs.
func appsync_ListGraphqlApis(cfg aws.Config, client *appsync.Client) {
	input := &appsync.ListGraphqlApisInput{}

	if len(_appsyncApiType) > 0 {
		if err := assignInputField(input, "ApiType", _appsyncApiType); err != nil {
			log.Errorf("invalid --api-type: %s", err.Error())
			return
		}
	}
	if len(_appsyncMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appsyncMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appsyncNextToken) > 0 {
		input.NextToken = aws.String(_appsyncNextToken)
	}
	if len(_appsyncOwner) > 0 {
		if err := assignInputField(input, "Owner", _appsyncOwner); err != nil {
			log.Errorf("invalid --owner: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListGraphqlApis(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appsync.ListGraphqlApisOutput
	p := appsync.NewListGraphqlApisPaginator(client, input)
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

// Lists the resolvers for a given API and type.
func appsync_ListResolvers(cfg aws.Config, client *appsync.Client) {
	input := &appsync.ListResolversInput{
		// ApiId: *string, // Required
		// TypeName: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncTypeName) > 0 {
		input.TypeName = aws.String(_appsyncTypeName)
	}
	if len(_appsyncMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appsyncMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appsyncNextToken) > 0 {
		input.NextToken = aws.String(_appsyncNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListResolvers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appsync.ListResolversOutput
	p := appsync.NewListResolversPaginator(client, input)
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

// List the resolvers that are associated with a specific function.
func appsync_ListResolversByFunction(cfg aws.Config, client *appsync.Client) {
	input := &appsync.ListResolversByFunctionInput{
		// ApiId: *string, // Required
		// FunctionId: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncFunctionId) > 0 {
		input.FunctionId = aws.String(_appsyncFunctionId)
	}
	if len(_appsyncMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appsyncMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appsyncNextToken) > 0 {
		input.NextToken = aws.String(_appsyncNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListResolversByFunction(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appsync.ListResolversByFunctionOutput
	p := appsync.NewListResolversByFunctionPaginator(client, input)
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

// Lists the SourceApiAssociationSummary data.
func appsync_ListSourceApiAssociations(cfg aws.Config, client *appsync.Client) {
	input := &appsync.ListSourceApiAssociationsInput{
		// ApiId: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appsyncMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appsyncNextToken) > 0 {
		input.NextToken = aws.String(_appsyncNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSourceApiAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appsync.ListSourceApiAssociationsOutput
	p := appsync.NewListSourceApiAssociationsPaginator(client, input)
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

// Lists the tags for a resource.
func appsync_ListTagsForResource(cfg aws.Config, client *appsync.Client) {
	input := &appsync.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_appsyncResourceArn) > 0 {
		input.ResourceArn = aws.String(_appsyncResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the types for a given API.
func appsync_ListTypes(cfg aws.Config, client *appsync.Client) {
	input := &appsync.ListTypesInput{
		// ApiId: *string, // Required
		// Format: types.TypeDefinitionFormat, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncFormat) > 0 {
		if err := assignInputField(input, "Format", _appsyncFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_appsyncMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appsyncMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appsyncNextToken) > 0 {
		input.NextToken = aws.String(_appsyncNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appsync.ListTypesOutput
	p := appsync.NewListTypesPaginator(client, input)
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

// Lists Type objects by the source API association ID.
func appsync_ListTypesByAssociation(cfg aws.Config, client *appsync.Client) {
	input := &appsync.ListTypesByAssociationInput{
		// AssociationId: *string, // Required
		// Format: types.TypeDefinitionFormat, // Required
		// MergedApiIdentifier: *string, // Required
	}

	if len(_appsyncAssociationId) > 0 {
		input.AssociationId = aws.String(_appsyncAssociationId)
	}
	if len(_appsyncFormat) > 0 {
		if err := assignInputField(input, "Format", _appsyncFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_appsyncMergedApiIdentifier) > 0 {
		input.MergedApiIdentifier = aws.String(_appsyncMergedApiIdentifier)
	}
	if len(_appsyncMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _appsyncMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_appsyncNextToken) > 0 {
		input.NextToken = aws.String(_appsyncNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTypesByAssociation(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appsync.ListTypesByAssociationOutput
	p := appsync.NewListTypesByAssociationPaginator(client, input)
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

// Creates a list of environmental variables in an API by its ID value.
// When creating an environmental variable, it must follow the constraints below:
//
// - Both JavaScript and VTL templates support environmental variables.
//
// - Environmental variables are not evaluated before function invocation.
//
// - Environmental variables only support string values.
//
// - Any defined value in an environmental variable is considered a string
// literal and not expanded.
//
// - Variable evaluations should ideally be performed in the function code.
//
// When creating an environmental variable key-value pair, it must follow the
// additional constraints below:
//
// - Keys must begin with a letter.
//
// - Keys must be at least two characters long.
//
// - Keys can only contain letters, numbers, and the underscore character (_).
//
// - Values can be up to 512 characters long.
//
// - You can configure up to 50 key-value pairs in a GraphQL API.
//
// You can create a list of environmental variables by adding it to the
// environmentVariables payload as a list in the format
// {"key1":"value1","key2":"value2", …} . Note that each call of the
// PutGraphqlApiEnvironmentVariables action will result in the overwriting of the
// existing environmental variable list of that API. This means the existing
// environmental variables will be lost. To avoid this, you must include all
// existing and new environmental variables in the list each time you call this
// action.
func appsync_PutGraphqlApiEnvironmentVariables(cfg aws.Config, client *appsync.Client) {
	input := &appsync.PutGraphqlApiEnvironmentVariablesInput{
		// ApiId: *string, // Required
		// EnvironmentVariables: map[string]string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncEnvironmentVariables) > 0 {
		if err := assignInputField(input, "EnvironmentVariables", _appsyncEnvironmentVariables); err != nil {
			log.Errorf("invalid --environment-variables: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutGraphqlApiEnvironmentVariables(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new introspection. Returns the introspectionId of the new
// introspection after its creation.
func appsync_StartDataSourceIntrospection(cfg aws.Config, client *appsync.Client) {
	input := &appsync.StartDataSourceIntrospectionInput{}

	if len(_appsyncRdsDataApiConfig) > 0 {
		if err := assignInputField(input, "RdsDataApiConfig", _appsyncRdsDataApiConfig); err != nil {
			log.Errorf("invalid --rds-data-api-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartDataSourceIntrospection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a new schema to your GraphQL API.
// This operation is asynchronous. Use to determine when it has completed.
func appsync_StartSchemaCreation(cfg aws.Config, client *appsync.Client) {
	input := &appsync.StartSchemaCreationInput{
		// ApiId: *string, // Required
		// Definition: []byte, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncDefinition) > 0 {
		if err := assignInputField(input, "Definition", _appsyncDefinition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartSchemaCreation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates a merge operation. Returns a status that shows the result of the
// merge operation.
func appsync_StartSchemaMerge(cfg aws.Config, client *appsync.Client) {
	input := &appsync.StartSchemaMergeInput{
		// AssociationId: *string, // Required
		// MergedApiIdentifier: *string, // Required
	}

	if len(_appsyncAssociationId) > 0 {
		input.AssociationId = aws.String(_appsyncAssociationId)
	}
	if len(_appsyncMergedApiIdentifier) > 0 {
		input.MergedApiIdentifier = aws.String(_appsyncMergedApiIdentifier)
	}

	if resp, err := client.StartSchemaMerge(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tags a resource with user-supplied tags.
func appsync_TagResource(cfg aws.Config, client *appsync.Client) {
	input := &appsync.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_appsyncResourceArn) > 0 {
		input.ResourceArn = aws.String(_appsyncResourceArn)
	}
	if len(_appsyncTags) > 0 {
		if err := assignInputField(input, "Tags", _appsyncTags); err != nil {
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

// Untags a resource.
func appsync_UntagResource(cfg aws.Config, client *appsync.Client) {
	input := &appsync.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_appsyncResourceArn) > 0 {
		input.ResourceArn = aws.String(_appsyncResourceArn)
	}
	if len(_appsyncTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _appsyncTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an Api .
func appsync_UpdateApi(cfg aws.Config, client *appsync.Client) {
	input := &appsync.UpdateApiInput{
		// ApiId: *string, // Required
		// EventConfig: *types.EventConfig, // Required
		// Name: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncEventConfig) > 0 {
		if err := assignInputField(input, "EventConfig", _appsyncEventConfig); err != nil {
			log.Errorf("invalid --event-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncName) > 0 {
		input.Name = aws.String(_appsyncName)
	}
	if len(_appsyncOwnerContact) > 0 {
		input.OwnerContact = aws.String(_appsyncOwnerContact)
	}

	if resp, err := client.UpdateApi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the cache for the GraphQL API.
func appsync_UpdateApiCache(cfg aws.Config, client *appsync.Client) {
	input := &appsync.UpdateApiCacheInput{
		// ApiCachingBehavior: types.ApiCachingBehavior, // Required
		// ApiId: *string, // Required
		// Ttl: int64, // Required
		// Type: types.ApiCacheType, // Required
	}

	if len(_appsyncApiCachingBehavior) > 0 {
		if err := assignInputField(input, "ApiCachingBehavior", _appsyncApiCachingBehavior); err != nil {
			log.Errorf("invalid --api-caching-behavior: %s", err.Error())
			return
		}
	}
	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncTtl) > 0 {
		if err := assignInputField(input, "Ttl", _appsyncTtl); err != nil {
			log.Errorf("invalid --ttl: %s", err.Error())
			return
		}
	}
	if len(_appsyncType) > 0 {
		if err := assignInputField(input, "Type", _appsyncType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_appsyncHealthMetricsConfig) > 0 {
		if err := assignInputField(input, "HealthMetricsConfig", _appsyncHealthMetricsConfig); err != nil {
			log.Errorf("invalid --health-metrics-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateApiCache(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an API key. You can update the key as long as it's not deleted.
func appsync_UpdateApiKey(cfg aws.Config, client *appsync.Client) {
	input := &appsync.UpdateApiKeyInput{
		// ApiId: *string, // Required
		// Id: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncId) > 0 {
		input.Id = aws.String(_appsyncId)
	}
	if len(_appsyncDescription) > 0 {
		input.Description = aws.String(_appsyncDescription)
	}
	if len(_appsyncExpires) > 0 {
		if err := assignInputField(input, "Expires", _appsyncExpires); err != nil {
			log.Errorf("invalid --expires: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateApiKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a ChannelNamespace associated with an Api .
func appsync_UpdateChannelNamespace(cfg aws.Config, client *appsync.Client) {
	input := &appsync.UpdateChannelNamespaceInput{
		// ApiId: *string, // Required
		// Name: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncName) > 0 {
		input.Name = aws.String(_appsyncName)
	}
	if len(_appsyncCodeHandlers) > 0 {
		input.CodeHandlers = aws.String(_appsyncCodeHandlers)
	}
	if len(_appsyncHandlerConfigs) > 0 {
		if err := assignInputField(input, "HandlerConfigs", _appsyncHandlerConfigs); err != nil {
			log.Errorf("invalid --handler-configs: %s", err.Error())
			return
		}
	}
	if len(_appsyncPublishAuthModes) > 0 {
		if err := assignInputField(input, "PublishAuthModes", _appsyncPublishAuthModes); err != nil {
			log.Errorf("invalid --publish-auth-modes: %s", err.Error())
			return
		}
	}
	if len(_appsyncSubscribeAuthModes) > 0 {
		if err := assignInputField(input, "SubscribeAuthModes", _appsyncSubscribeAuthModes); err != nil {
			log.Errorf("invalid --subscribe-auth-modes: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateChannelNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a DataSource object.
func appsync_UpdateDataSource(cfg aws.Config, client *appsync.Client) {
	input := &appsync.UpdateDataSourceInput{
		// ApiId: *string, // Required
		// Name: *string, // Required
		// Type: types.DataSourceType, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncName) > 0 {
		input.Name = aws.String(_appsyncName)
	}
	if len(_appsyncType) > 0 {
		if err := assignInputField(input, "Type", _appsyncType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_appsyncDescription) > 0 {
		input.Description = aws.String(_appsyncDescription)
	}
	if len(_appsyncDynamodbConfig) > 0 {
		if err := assignInputField(input, "DynamodbConfig", _appsyncDynamodbConfig); err != nil {
			log.Errorf("invalid --dynamodb-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncElasticsearchConfig) > 0 {
		if err := assignInputField(input, "ElasticsearchConfig", _appsyncElasticsearchConfig); err != nil {
			log.Errorf("invalid --elasticsearch-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncEventBridgeConfig) > 0 {
		if err := assignInputField(input, "EventBridgeConfig", _appsyncEventBridgeConfig); err != nil {
			log.Errorf("invalid --event-bridge-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncHttpConfig) > 0 {
		if err := assignInputField(input, "HttpConfig", _appsyncHttpConfig); err != nil {
			log.Errorf("invalid --http-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncLambdaConfig) > 0 {
		if err := assignInputField(input, "LambdaConfig", _appsyncLambdaConfig); err != nil {
			log.Errorf("invalid --lambda-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncMetricsConfig) > 0 {
		if err := assignInputField(input, "MetricsConfig", _appsyncMetricsConfig); err != nil {
			log.Errorf("invalid --metrics-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncOpenSearchServiceConfig) > 0 {
		if err := assignInputField(input, "OpenSearchServiceConfig", _appsyncOpenSearchServiceConfig); err != nil {
			log.Errorf("invalid --open-search-service-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncRelationalDatabaseConfig) > 0 {
		if err := assignInputField(input, "RelationalDatabaseConfig", _appsyncRelationalDatabaseConfig); err != nil {
			log.Errorf("invalid --relational-database-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncServiceRoleArn) > 0 {
		input.ServiceRoleArn = aws.String(_appsyncServiceRoleArn)
	}

	if resp, err := client.UpdateDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a custom DomainName object.
func appsync_UpdateDomainName(cfg aws.Config, client *appsync.Client) {
	input := &appsync.UpdateDomainNameInput{
		// DomainName: *string, // Required
	}

	if len(_appsyncDomainName) > 0 {
		input.DomainName = aws.String(_appsyncDomainName)
	}
	if len(_appsyncDescription) > 0 {
		input.Description = aws.String(_appsyncDescription)
	}

	if resp, err := client.UpdateDomainName(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Function object.
func appsync_UpdateFunction(cfg aws.Config, client *appsync.Client) {
	input := &appsync.UpdateFunctionInput{
		// ApiId: *string, // Required
		// DataSourceName: *string, // Required
		// FunctionId: *string, // Required
		// Name: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncDataSourceName) > 0 {
		input.DataSourceName = aws.String(_appsyncDataSourceName)
	}
	if len(_appsyncFunctionId) > 0 {
		input.FunctionId = aws.String(_appsyncFunctionId)
	}
	if len(_appsyncName) > 0 {
		input.Name = aws.String(_appsyncName)
	}
	if len(_appsyncCode) > 0 {
		input.Code = aws.String(_appsyncCode)
	}
	if len(_appsyncDescription) > 0 {
		input.Description = aws.String(_appsyncDescription)
	}
	if len(_appsyncFunctionVersion) > 0 {
		input.FunctionVersion = aws.String(_appsyncFunctionVersion)
	}
	if len(_appsyncMaxBatchSize) > 0 {
		if err := assignInputField(input, "MaxBatchSize", _appsyncMaxBatchSize); err != nil {
			log.Errorf("invalid --max-batch-size: %s", err.Error())
			return
		}
	}
	if len(_appsyncRequestMappingTemplate) > 0 {
		input.RequestMappingTemplate = aws.String(_appsyncRequestMappingTemplate)
	}
	if len(_appsyncResponseMappingTemplate) > 0 {
		input.ResponseMappingTemplate = aws.String(_appsyncResponseMappingTemplate)
	}
	if len(_appsyncRuntime) > 0 {
		if err := assignInputField(input, "Runtime", _appsyncRuntime); err != nil {
			log.Errorf("invalid --runtime: %s", err.Error())
			return
		}
	}
	if len(_appsyncSyncConfig) > 0 {
		if err := assignInputField(input, "SyncConfig", _appsyncSyncConfig); err != nil {
			log.Errorf("invalid --sync-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFunction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a GraphqlApi object.
func appsync_UpdateGraphqlApi(cfg aws.Config, client *appsync.Client) {
	input := &appsync.UpdateGraphqlApiInput{
		// ApiId: *string, // Required
		// AuthenticationType: types.AuthenticationType, // Required
		// Name: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncAuthenticationType) > 0 {
		if err := assignInputField(input, "AuthenticationType", _appsyncAuthenticationType); err != nil {
			log.Errorf("invalid --authentication-type: %s", err.Error())
			return
		}
	}
	if len(_appsyncName) > 0 {
		input.Name = aws.String(_appsyncName)
	}
	if len(_appsyncAdditionalAuthenticationProviders) > 0 {
		if err := assignInputField(input, "AdditionalAuthenticationProviders", _appsyncAdditionalAuthenticationProviders); err != nil {
			log.Errorf("invalid --additional-authentication-providers: %s", err.Error())
			return
		}
	}
	if len(_appsyncEnhancedMetricsConfig) > 0 {
		if err := assignInputField(input, "EnhancedMetricsConfig", _appsyncEnhancedMetricsConfig); err != nil {
			log.Errorf("invalid --enhanced-metrics-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncIntrospectionConfig) > 0 {
		if err := assignInputField(input, "IntrospectionConfig", _appsyncIntrospectionConfig); err != nil {
			log.Errorf("invalid --introspection-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncLambdaAuthorizerConfig) > 0 {
		if err := assignInputField(input, "LambdaAuthorizerConfig", _appsyncLambdaAuthorizerConfig); err != nil {
			log.Errorf("invalid --lambda-authorizer-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncLogConfig) > 0 {
		if err := assignInputField(input, "LogConfig", _appsyncLogConfig); err != nil {
			log.Errorf("invalid --log-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncMergedApiExecutionRoleArn) > 0 {
		input.MergedApiExecutionRoleArn = aws.String(_appsyncMergedApiExecutionRoleArn)
	}
	if len(_appsyncOpenIDConnectConfig) > 0 {
		if err := assignInputField(input, "OpenIDConnectConfig", _appsyncOpenIDConnectConfig); err != nil {
			log.Errorf("invalid --open-id-connect-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncOwnerContact) > 0 {
		input.OwnerContact = aws.String(_appsyncOwnerContact)
	}
	if len(_appsyncQueryDepthLimit) > 0 {
		if err := assignInputField(input, "QueryDepthLimit", _appsyncQueryDepthLimit); err != nil {
			log.Errorf("invalid --query-depth-limit: %s", err.Error())
			return
		}
	}
	if len(_appsyncResolverCountLimit) > 0 {
		if err := assignInputField(input, "ResolverCountLimit", _appsyncResolverCountLimit); err != nil {
			log.Errorf("invalid --resolver-count-limit: %s", err.Error())
			return
		}
	}
	if len(_appsyncUserPoolConfig) > 0 {
		if err := assignInputField(input, "UserPoolConfig", _appsyncUserPoolConfig); err != nil {
			log.Errorf("invalid --user-pool-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncXrayEnabled) > 0 {
		if err := assignInputField(input, "XrayEnabled", _appsyncXrayEnabled); err != nil {
			log.Errorf("invalid --xray-enabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateGraphqlApi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Resolver object.
func appsync_UpdateResolver(cfg aws.Config, client *appsync.Client) {
	input := &appsync.UpdateResolverInput{
		// ApiId: *string, // Required
		// FieldName: *string, // Required
		// TypeName: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncFieldName) > 0 {
		input.FieldName = aws.String(_appsyncFieldName)
	}
	if len(_appsyncTypeName) > 0 {
		input.TypeName = aws.String(_appsyncTypeName)
	}
	if len(_appsyncCachingConfig) > 0 {
		if err := assignInputField(input, "CachingConfig", _appsyncCachingConfig); err != nil {
			log.Errorf("invalid --caching-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncCode) > 0 {
		input.Code = aws.String(_appsyncCode)
	}
	if len(_appsyncDataSourceName) > 0 {
		input.DataSourceName = aws.String(_appsyncDataSourceName)
	}
	if len(_appsyncKind) > 0 {
		if err := assignInputField(input, "Kind", _appsyncKind); err != nil {
			log.Errorf("invalid --kind: %s", err.Error())
			return
		}
	}
	if len(_appsyncMaxBatchSize) > 0 {
		if err := assignInputField(input, "MaxBatchSize", _appsyncMaxBatchSize); err != nil {
			log.Errorf("invalid --max-batch-size: %s", err.Error())
			return
		}
	}
	if len(_appsyncMetricsConfig) > 0 {
		if err := assignInputField(input, "MetricsConfig", _appsyncMetricsConfig); err != nil {
			log.Errorf("invalid --metrics-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncPipelineConfig) > 0 {
		if err := assignInputField(input, "PipelineConfig", _appsyncPipelineConfig); err != nil {
			log.Errorf("invalid --pipeline-config: %s", err.Error())
			return
		}
	}
	if len(_appsyncRequestMappingTemplate) > 0 {
		input.RequestMappingTemplate = aws.String(_appsyncRequestMappingTemplate)
	}
	if len(_appsyncResponseMappingTemplate) > 0 {
		input.ResponseMappingTemplate = aws.String(_appsyncResponseMappingTemplate)
	}
	if len(_appsyncRuntime) > 0 {
		if err := assignInputField(input, "Runtime", _appsyncRuntime); err != nil {
			log.Errorf("invalid --runtime: %s", err.Error())
			return
		}
	}
	if len(_appsyncSyncConfig) > 0 {
		if err := assignInputField(input, "SyncConfig", _appsyncSyncConfig); err != nil {
			log.Errorf("invalid --sync-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateResolver(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates some of the configuration choices of a particular source API
// association.
func appsync_UpdateSourceApiAssociation(cfg aws.Config, client *appsync.Client) {
	input := &appsync.UpdateSourceApiAssociationInput{
		// AssociationId: *string, // Required
		// MergedApiIdentifier: *string, // Required
	}

	if len(_appsyncAssociationId) > 0 {
		input.AssociationId = aws.String(_appsyncAssociationId)
	}
	if len(_appsyncMergedApiIdentifier) > 0 {
		input.MergedApiIdentifier = aws.String(_appsyncMergedApiIdentifier)
	}
	if len(_appsyncDescription) > 0 {
		input.Description = aws.String(_appsyncDescription)
	}
	if len(_appsyncSourceApiAssociationConfig) > 0 {
		if err := assignInputField(input, "SourceApiAssociationConfig", _appsyncSourceApiAssociationConfig); err != nil {
			log.Errorf("invalid --source-api-association-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSourceApiAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Type object.
func appsync_UpdateType(cfg aws.Config, client *appsync.Client) {
	input := &appsync.UpdateTypeInput{
		// ApiId: *string, // Required
		// Format: types.TypeDefinitionFormat, // Required
		// TypeName: *string, // Required
	}

	if len(_appsyncApiId) > 0 {
		input.ApiId = aws.String(_appsyncApiId)
	}
	if len(_appsyncFormat) > 0 {
		if err := assignInputField(input, "Format", _appsyncFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_appsyncTypeName) > 0 {
		input.TypeName = aws.String(_appsyncTypeName)
	}
	if len(_appsyncDefinition) > 0 {
		input.Definition = aws.String(_appsyncDefinition)
	}

	if resp, err := client.UpdateType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_appsyncCmd)
	_appsyncCmd.Flags().SortFlags = false

	_appsyncCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_appsyncCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_appsyncCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_appsyncCmd.Flags().StringVarP(&_appsyncAdditionalAuthenticationProviders, "additional-authentication-providers", "", "", "Additional Authentication Providers")
	_appsyncCmd.Flags().StringVarP(&_appsyncApiCachingBehavior, "api-caching-behavior", "", "", "API Caching Behavior")
	_appsyncCmd.Flags().StringVarP(&_appsyncApiId, "api-id", "", "", "API ID")
	_appsyncCmd.Flags().StringVarP(&_appsyncApiType, "api-type", "", "", "API Type")
	_appsyncCmd.Flags().StringVarP(&_appsyncAssociationId, "association-id", "", "", "Association ID")
	_appsyncCmd.Flags().StringVarP(&_appsyncAtRestEncryptionEnabled, "at-rest-encryption-enabled", "", "", "At Rest Encryption Enabled")
	_appsyncCmd.Flags().StringVarP(&_appsyncAuthenticationType, "authentication-type", "", "", "Authentication Type")
	_appsyncCmd.Flags().StringVarP(&_appsyncCachingConfig, "caching-config", "", "", "Caching Config")
	_appsyncCmd.Flags().StringVarP(&_appsyncCertificateArn, "certificate-arn", "", "", "Certificate ARN")
	_appsyncCmd.Flags().StringVarP(&_appsyncCode, "code", "", "", "Code")
	_appsyncCmd.Flags().StringVarP(&_appsyncCodeHandlers, "code-handlers", "", "", "Code Handlers")
	_appsyncCmd.Flags().StringVarP(&_appsyncContext, "context", "", "", "Context")
	_appsyncCmd.Flags().StringVarP(&_appsyncDataSourceName, "data-source-name", "", "", "Data Source Name")
	_appsyncCmd.Flags().StringVarP(&_appsyncDefinition, "definition", "", "", "Definition")
	_appsyncCmd.Flags().StringVarP(&_appsyncDescription, "description", "", "", "Description")
	_appsyncCmd.Flags().StringVarP(&_appsyncDomainName, "domain-name", "", "", "Domain Name")
	_appsyncCmd.Flags().StringVarP(&_appsyncDynamodbConfig, "dynamodb-config", "", "", "Dynamodb Config")
	_appsyncCmd.Flags().StringVarP(&_appsyncElasticsearchConfig, "elasticsearch-config", "", "", "Elasticsearch Config")
	_appsyncCmd.Flags().StringVarP(&_appsyncEnhancedMetricsConfig, "enhanced-metrics-config", "", "", "Enhanced Metrics Config")
	_appsyncCmd.Flags().StringVarP(&_appsyncEnvironmentVariables, "environment-variables", "", "", "Environment Variables")
	_appsyncCmd.Flags().StringVarP(&_appsyncEventBridgeConfig, "event-bridge-config", "", "", "Event Bridge Config")
	_appsyncCmd.Flags().StringVarP(&_appsyncEventConfig, "event-config", "", "", "Event Config")
	_appsyncCmd.Flags().StringVarP(&_appsyncExpires, "expires", "", "", "Expires")
	_appsyncCmd.Flags().StringVarP(&_appsyncFieldName, "field-name", "", "", "Field Name")
	_appsyncCmd.Flags().StringVarP(&_appsyncFormat, "format", "", "", "Format")
	_appsyncCmd.Flags().StringVarP(&_appsyncFunction, "function", "", "", "Function")
	_appsyncCmd.Flags().StringVarP(&_appsyncFunctionId, "function-id", "", "", "Function ID")
	_appsyncCmd.Flags().StringVarP(&_appsyncFunctionVersion, "function-version", "", "", "Function Version")
	_appsyncCmd.Flags().StringVarP(&_appsyncHandlerConfigs, "handler-configs", "", "", "Handler Configs")
	_appsyncCmd.Flags().StringVarP(&_appsyncHealthMetricsConfig, "health-metrics-config", "", "", "Health Metrics Config")
	_appsyncCmd.Flags().StringVarP(&_appsyncHttpConfig, "http-config", "", "", "HTTP Config")
	_appsyncCmd.Flags().StringVarP(&_appsyncId, "id", "", "", "ID")
	_appsyncCmd.Flags().StringVarP(&_appsyncIncludeDirectives, "include-directives", "", "", "Include Directives")
	_appsyncCmd.Flags().StringVarP(&_appsyncIncludeModelsSDL, "include-models-sdl", "", "", "Include Models Sdl")
	_appsyncCmd.Flags().StringVarP(&_appsyncIntrospectionConfig, "introspection-config", "", "", "Introspection Config")
	_appsyncCmd.Flags().StringVarP(&_appsyncIntrospectionId, "introspection-id", "", "", "Introspection ID")
	_appsyncCmd.Flags().StringVarP(&_appsyncKind, "kind", "", "", "Kind")
	_appsyncCmd.Flags().StringVarP(&_appsyncLambdaAuthorizerConfig, "lambda-authorizer-config", "", "", "Lambda Authorizer Config")
	_appsyncCmd.Flags().StringVarP(&_appsyncLambdaConfig, "lambda-config", "", "", "Lambda Config")
	_appsyncCmd.Flags().StringVarP(&_appsyncLogConfig, "log-config", "", "", "Log Config")
	_appsyncCmd.Flags().StringVarP(&_appsyncMaxBatchSize, "max-batch-size", "", "", "Max Batch Size")
	_appsyncCmd.Flags().StringVarP(&_appsyncMaxResults, "max-results", "", "", "Max Results")
	_appsyncCmd.Flags().StringVarP(&_appsyncMergedApiExecutionRoleArn, "merged-api-execution-role-arn", "", "", "Merged API Execution Role ARN")
	_appsyncCmd.Flags().StringVarP(&_appsyncMergedApiIdentifier, "merged-api-identifier", "", "", "Merged API Identifier")
	_appsyncCmd.Flags().StringVarP(&_appsyncMetricsConfig, "metrics-config", "", "", "Metrics Config")
	_appsyncCmd.Flags().StringVarP(&_appsyncName, "name", "", "", "Name")
	_appsyncCmd.Flags().StringVarP(&_appsyncNextToken, "next-token", "", "", "Next Token")
	_appsyncCmd.Flags().StringVarP(&_appsyncOpenIDConnectConfig, "open-id-connect-config", "", "", "Open ID Connect Config")
	_appsyncCmd.Flags().StringVarP(&_appsyncOpenSearchServiceConfig, "open-search-service-config", "", "", "Open Search Service Config")
	_appsyncCmd.Flags().StringVarP(&_appsyncOwner, "owner", "", "", "Owner")
	_appsyncCmd.Flags().StringVarP(&_appsyncOwnerContact, "owner-contact", "", "", "Owner Contact")
	_appsyncCmd.Flags().StringVarP(&_appsyncPipelineConfig, "pipeline-config", "", "", "Pipeline Config")
	_appsyncCmd.Flags().StringVarP(&_appsyncPublishAuthModes, "publish-auth-modes", "", "", "Publish Auth Modes")
	_appsyncCmd.Flags().StringVarP(&_appsyncQueryDepthLimit, "query-depth-limit", "", "", "Query Depth Limit")
	_appsyncCmd.Flags().StringVarP(&_appsyncRdsDataApiConfig, "rds-data-api-config", "", "", "RDS Data API Config")
	_appsyncCmd.Flags().StringVarP(&_appsyncRelationalDatabaseConfig, "relational-database-config", "", "", "Relational Database Config")
	_appsyncCmd.Flags().StringVarP(&_appsyncRequestMappingTemplate, "request-mapping-template", "", "", "Request Mapping Template")
	_appsyncCmd.Flags().StringVarP(&_appsyncResolverCountLimit, "resolver-count-limit", "", "", "Resolver Count Limit")
	_appsyncCmd.Flags().StringVarP(&_appsyncResourceArn, "resource-arn", "", "", "Resource ARN")
	_appsyncCmd.Flags().StringVarP(&_appsyncResponseMappingTemplate, "response-mapping-template", "", "", "Response Mapping Template")
	_appsyncCmd.Flags().StringVarP(&_appsyncRuntime, "runtime", "", "", "Runtime")
	_appsyncCmd.Flags().StringVarP(&_appsyncServiceRoleArn, "service-role-arn", "", "", "Service Role ARN")
	_appsyncCmd.Flags().StringVarP(&_appsyncSourceApiAssociationConfig, "source-api-association-config", "", "", "Source API Association Config")
	_appsyncCmd.Flags().StringVarP(&_appsyncSourceApiIdentifier, "source-api-identifier", "", "", "Source API Identifier")
	_appsyncCmd.Flags().StringVarP(&_appsyncSubscribeAuthModes, "subscribe-auth-modes", "", "", "Subscribe Auth Modes")
	_appsyncCmd.Flags().StringVarP(&_appsyncSyncConfig, "sync-config", "", "", "Sync Config")
	_appsyncCmd.Flags().StringSliceVarP(&_appsyncTagKeys, "tag-keys", "", nil, "Tag Keys")
	_appsyncCmd.Flags().StringVarP(&_appsyncTags, "tags", "", "", "Tags")
	_appsyncCmd.Flags().StringVarP(&_appsyncTemplate, "template", "", "", "Template")
	_appsyncCmd.Flags().StringVarP(&_appsyncTransitEncryptionEnabled, "transit-encryption-enabled", "", "", "Transit Encryption Enabled")
	_appsyncCmd.Flags().StringVarP(&_appsyncTtl, "ttl", "", "", "TTL")
	_appsyncCmd.Flags().StringVarP(&_appsyncType, "type", "", "", "Type")
	_appsyncCmd.Flags().StringVarP(&_appsyncTypeName, "type-name", "", "", "Type Name")
	_appsyncCmd.Flags().StringVarP(&_appsyncUserPoolConfig, "user-pool-config", "", "", "User Pool Config")
	_appsyncCmd.Flags().StringVarP(&_appsyncVisibility, "visibility", "", "", "Visibility")
	_appsyncCmd.Flags().StringVarP(&_appsyncXrayEnabled, "xray-enabled", "", "", "Xray Enabled")

	_appsyncCmd.Flags().BoolVarP(&_appsyncAssociateApi, "associate-api", "", false, "Associate API")
	_appsyncCmd.Flags().BoolVarP(&_appsyncAssociateMergedGraphqlApi, "associate-merged-graphql-api", "", false, "Associate Merged Graphql API")
	_appsyncCmd.Flags().BoolVarP(&_appsyncAssociateSourceGraphqlApi, "associate-source-graphql-api", "", false, "Associate Source Graphql API")
	_appsyncCmd.Flags().BoolVarP(&_appsyncCreateApi, "create-api", "", false, "Create API")
	_appsyncCmd.Flags().BoolVarP(&_appsyncCreateApiCache, "create-api-cache", "", false, "Create API Cache")
	_appsyncCmd.Flags().BoolVarP(&_appsyncCreateApiKey, "create-api-key", "", false, "Create API Key")
	_appsyncCmd.Flags().BoolVarP(&_appsyncCreateChannelNamespace, "create-channel-namespace", "", false, "Create Channel Namespace")
	_appsyncCmd.Flags().BoolVarP(&_appsyncCreateDataSource, "create-data-source", "", false, "Create Data Source")
	_appsyncCmd.Flags().BoolVarP(&_appsyncCreateDomainName, "create-domain-name", "", false, "Create Domain Name")
	_appsyncCmd.Flags().BoolVarP(&_appsyncCreateFunction, "create-function", "", false, "Create Function")
	_appsyncCmd.Flags().BoolVarP(&_appsyncCreateGraphqlApi, "create-graphql-api", "", false, "Create Graphql API")
	_appsyncCmd.Flags().BoolVarP(&_appsyncCreateResolver, "create-resolver", "", false, "Create Resolver")
	_appsyncCmd.Flags().BoolVarP(&_appsyncCreateType, "create-type", "", false, "Create Type")
	_appsyncCmd.Flags().BoolVarP(&_appsyncDeleteApi, "delete-api", "", false, "Delete API")
	_appsyncCmd.Flags().BoolVarP(&_appsyncDeleteApiCache, "delete-api-cache", "", false, "Delete API Cache")
	_appsyncCmd.Flags().BoolVarP(&_appsyncDeleteApiKey, "delete-api-key", "", false, "Delete API Key")
	_appsyncCmd.Flags().BoolVarP(&_appsyncDeleteChannelNamespace, "delete-channel-namespace", "", false, "Delete Channel Namespace")
	_appsyncCmd.Flags().BoolVarP(&_appsyncDeleteDataSource, "delete-data-source", "", false, "Delete Data Source")
	_appsyncCmd.Flags().BoolVarP(&_appsyncDeleteDomainName, "delete-domain-name", "", false, "Delete Domain Name")
	_appsyncCmd.Flags().BoolVarP(&_appsyncDeleteFunction, "delete-function", "", false, "Delete Function")
	_appsyncCmd.Flags().BoolVarP(&_appsyncDeleteGraphqlApi, "delete-graphql-api", "", false, "Delete Graphql API")
	_appsyncCmd.Flags().BoolVarP(&_appsyncDeleteResolver, "delete-resolver", "", false, "Delete Resolver")
	_appsyncCmd.Flags().BoolVarP(&_appsyncDeleteType, "delete-type", "", false, "Delete Type")
	_appsyncCmd.Flags().BoolVarP(&_appsyncDisassociateApi, "disassociate-api", "", false, "Disassociate API")
	_appsyncCmd.Flags().BoolVarP(&_appsyncDisassociateMergedGraphqlApi, "disassociate-merged-graphql-api", "", false, "Disassociate Merged Graphql API")
	_appsyncCmd.Flags().BoolVarP(&_appsyncDisassociateSourceGraphqlApi, "disassociate-source-graphql-api", "", false, "Disassociate Source Graphql API")
	_appsyncCmd.Flags().BoolVarP(&_appsyncEvaluateCode, "evaluate-code", "", false, "Evaluate Code")
	_appsyncCmd.Flags().BoolVarP(&_appsyncEvaluateMappingTemplate, "evaluate-mapping-template", "", false, "Evaluate Mapping Template")
	_appsyncCmd.Flags().BoolVarP(&_appsyncFlushApiCache, "flush-api-cache", "", false, "Flush API Cache")
	_appsyncCmd.Flags().BoolVarP(&_appsyncGetApi, "get-api", "", false, "Get API")
	_appsyncCmd.Flags().BoolVarP(&_appsyncGetApiAssociation, "get-api-association", "", false, "Get API Association")
	_appsyncCmd.Flags().BoolVarP(&_appsyncGetApiCache, "get-api-cache", "", false, "Get API Cache")
	_appsyncCmd.Flags().BoolVarP(&_appsyncGetChannelNamespace, "get-channel-namespace", "", false, "Get Channel Namespace")
	_appsyncCmd.Flags().BoolVarP(&_appsyncGetDataSource, "get-data-source", "", false, "Get Data Source")
	_appsyncCmd.Flags().BoolVarP(&_appsyncGetDataSourceIntrospection, "get-data-source-introspection", "", false, "Get Data Source Introspection")
	_appsyncCmd.Flags().BoolVarP(&_appsyncGetDomainName, "get-domain-name", "", false, "Get Domain Name")
	_appsyncCmd.Flags().BoolVarP(&_appsyncGetFunction, "get-function", "", false, "Get Function")
	_appsyncCmd.Flags().BoolVarP(&_appsyncGetGraphqlApi, "get-graphql-api", "", false, "Get Graphql API")
	_appsyncCmd.Flags().BoolVarP(&_appsyncGetGraphqlApiEnvironmentVariables, "get-graphql-api-environment-variables", "", false, "Get Graphql API Environment Variables")
	_appsyncCmd.Flags().BoolVarP(&_appsyncGetIntrospectionSchema, "get-introspection-schema", "", false, "Get Introspection Schema")
	_appsyncCmd.Flags().BoolVarP(&_appsyncGetResolver, "get-resolver", "", false, "Get Resolver")
	_appsyncCmd.Flags().BoolVarP(&_appsyncGetSchemaCreationStatus, "get-schema-creation-status", "", false, "Get Schema Creation Status")
	_appsyncCmd.Flags().BoolVarP(&_appsyncGetSourceApiAssociation, "get-source-api-association", "", false, "Get Source API Association")
	_appsyncCmd.Flags().BoolVarP(&_appsyncGetType, "get-type", "", false, "Get Type")
	_appsyncCmd.Flags().BoolVarP(&_appsyncListApiKeys, "list-api-keys", "", false, "List API Keys")
	_appsyncCmd.Flags().BoolVarP(&_appsyncListApis, "list-apis", "", false, "List Apis")
	_appsyncCmd.Flags().BoolVarP(&_appsyncListChannelNamespaces, "list-channel-namespaces", "", false, "List Channel Namespaces")
	_appsyncCmd.Flags().BoolVarP(&_appsyncListDataSources, "list-data-sources", "", false, "List Data Sources")
	_appsyncCmd.Flags().BoolVarP(&_appsyncListDomainNames, "list-domain-names", "", false, "List Domain Names")
	_appsyncCmd.Flags().BoolVarP(&_appsyncListFunctions, "list-functions", "", false, "List Functions")
	_appsyncCmd.Flags().BoolVarP(&_appsyncListGraphqlApis, "list-graphql-apis", "", false, "List Graphql Apis")
	_appsyncCmd.Flags().BoolVarP(&_appsyncListResolvers, "list-resolvers", "", false, "List Resolvers")
	_appsyncCmd.Flags().BoolVarP(&_appsyncListResolversByFunction, "list-resolvers-by-function", "", false, "List Resolvers By Function")
	_appsyncCmd.Flags().BoolVarP(&_appsyncListSourceApiAssociations, "list-source-api-associations", "", false, "List Source API Associations")
	_appsyncCmd.Flags().BoolVarP(&_appsyncListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_appsyncCmd.Flags().BoolVarP(&_appsyncListTypes, "list-types", "", false, "List Types")
	_appsyncCmd.Flags().BoolVarP(&_appsyncListTypesByAssociation, "list-types-by-association", "", false, "List Types By Association")
	_appsyncCmd.Flags().BoolVarP(&_appsyncPutGraphqlApiEnvironmentVariables, "put-graphql-api-environment-variables", "", false, "Put Graphql API Environment Variables")
	_appsyncCmd.Flags().BoolVarP(&_appsyncStartDataSourceIntrospection, "start-data-source-introspection", "", false, "Start Data Source Introspection")
	_appsyncCmd.Flags().BoolVarP(&_appsyncStartSchemaCreation, "start-schema-creation", "", false, "Start Schema Creation")
	_appsyncCmd.Flags().BoolVarP(&_appsyncStartSchemaMerge, "start-schema-merge", "", false, "Start Schema Merge")
	_appsyncCmd.Flags().BoolVarP(&_appsyncTagResource, "tag-resource", "", false, "Tag Resource")
	_appsyncCmd.Flags().BoolVarP(&_appsyncUntagResource, "untag-resource", "", false, "Untag Resource")
	_appsyncCmd.Flags().BoolVarP(&_appsyncUpdateApi, "update-api", "", false, "Update API")
	_appsyncCmd.Flags().BoolVarP(&_appsyncUpdateApiCache, "update-api-cache", "", false, "Update API Cache")
	_appsyncCmd.Flags().BoolVarP(&_appsyncUpdateApiKey, "update-api-key", "", false, "Update API Key")
	_appsyncCmd.Flags().BoolVarP(&_appsyncUpdateChannelNamespace, "update-channel-namespace", "", false, "Update Channel Namespace")
	_appsyncCmd.Flags().BoolVarP(&_appsyncUpdateDataSource, "update-data-source", "", false, "Update Data Source")
	_appsyncCmd.Flags().BoolVarP(&_appsyncUpdateDomainName, "update-domain-name", "", false, "Update Domain Name")
	_appsyncCmd.Flags().BoolVarP(&_appsyncUpdateFunction, "update-function", "", false, "Update Function")
	_appsyncCmd.Flags().BoolVarP(&_appsyncUpdateGraphqlApi, "update-graphql-api", "", false, "Update Graphql API")
	_appsyncCmd.Flags().BoolVarP(&_appsyncUpdateResolver, "update-resolver", "", false, "Update Resolver")
	_appsyncCmd.Flags().BoolVarP(&_appsyncUpdateSourceApiAssociation, "update-source-api-association", "", false, "Update Source API Association")
	_appsyncCmd.Flags().BoolVarP(&_appsyncUpdateType, "update-type", "", false, "Update Type")

}
