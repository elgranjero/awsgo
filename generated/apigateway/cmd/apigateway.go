package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// apigatewayCmd represents the apigateway command
var _apigatewayCmd = &cobra.Command{
	Use:   "apigateway",
	Short: "AWS apigateway CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := apigateway.NewFromConfig(cfg)
		if _apigatewayCreateApiKey {
			apigateway_CreateApiKey(cfg, client)
			return
		}
		if _apigatewayCreateAuthorizer {
			apigateway_CreateAuthorizer(cfg, client)
			return
		}
		if _apigatewayCreateBasePathMapping {
			apigateway_CreateBasePathMapping(cfg, client)
			return
		}
		if _apigatewayCreateDeployment {
			apigateway_CreateDeployment(cfg, client)
			return
		}
		if _apigatewayCreateDocumentationPart {
			apigateway_CreateDocumentationPart(cfg, client)
			return
		}
		if _apigatewayCreateDocumentationVersion {
			apigateway_CreateDocumentationVersion(cfg, client)
			return
		}
		if _apigatewayCreateDomainName {
			apigateway_CreateDomainName(cfg, client)
			return
		}
		if _apigatewayCreateDomainNameAccessAssociation {
			apigateway_CreateDomainNameAccessAssociation(cfg, client)
			return
		}
		if _apigatewayCreateModel {
			apigateway_CreateModel(cfg, client)
			return
		}
		if _apigatewayCreateRequestValidator {
			apigateway_CreateRequestValidator(cfg, client)
			return
		}
		if _apigatewayCreateResource {
			apigateway_CreateResource(cfg, client)
			return
		}
		if _apigatewayCreateRestApi {
			apigateway_CreateRestApi(cfg, client)
			return
		}
		if _apigatewayCreateStage {
			apigateway_CreateStage(cfg, client)
			return
		}
		if _apigatewayCreateUsagePlan {
			apigateway_CreateUsagePlan(cfg, client)
			return
		}
		if _apigatewayCreateUsagePlanKey {
			apigateway_CreateUsagePlanKey(cfg, client)
			return
		}
		if _apigatewayCreateVpcLink {
			apigateway_CreateVpcLink(cfg, client)
			return
		}
		if _apigatewayDeleteApiKey {
			apigateway_DeleteApiKey(cfg, client)
			return
		}
		if _apigatewayDeleteAuthorizer {
			apigateway_DeleteAuthorizer(cfg, client)
			return
		}
		if _apigatewayDeleteBasePathMapping {
			apigateway_DeleteBasePathMapping(cfg, client)
			return
		}
		if _apigatewayDeleteClientCertificate {
			apigateway_DeleteClientCertificate(cfg, client)
			return
		}
		if _apigatewayDeleteDeployment {
			apigateway_DeleteDeployment(cfg, client)
			return
		}
		if _apigatewayDeleteDocumentationPart {
			apigateway_DeleteDocumentationPart(cfg, client)
			return
		}
		if _apigatewayDeleteDocumentationVersion {
			apigateway_DeleteDocumentationVersion(cfg, client)
			return
		}
		if _apigatewayDeleteDomainName {
			apigateway_DeleteDomainName(cfg, client)
			return
		}
		if _apigatewayDeleteDomainNameAccessAssociation {
			apigateway_DeleteDomainNameAccessAssociation(cfg, client)
			return
		}
		if _apigatewayDeleteGatewayResponse {
			apigateway_DeleteGatewayResponse(cfg, client)
			return
		}
		if _apigatewayDeleteIntegration {
			apigateway_DeleteIntegration(cfg, client)
			return
		}
		if _apigatewayDeleteIntegrationResponse {
			apigateway_DeleteIntegrationResponse(cfg, client)
			return
		}
		if _apigatewayDeleteMethod {
			apigateway_DeleteMethod(cfg, client)
			return
		}
		if _apigatewayDeleteMethodResponse {
			apigateway_DeleteMethodResponse(cfg, client)
			return
		}
		if _apigatewayDeleteModel {
			apigateway_DeleteModel(cfg, client)
			return
		}
		if _apigatewayDeleteRequestValidator {
			apigateway_DeleteRequestValidator(cfg, client)
			return
		}
		if _apigatewayDeleteResource {
			apigateway_DeleteResource(cfg, client)
			return
		}
		if _apigatewayDeleteRestApi {
			apigateway_DeleteRestApi(cfg, client)
			return
		}
		if _apigatewayDeleteStage {
			apigateway_DeleteStage(cfg, client)
			return
		}
		if _apigatewayDeleteUsagePlan {
			apigateway_DeleteUsagePlan(cfg, client)
			return
		}
		if _apigatewayDeleteUsagePlanKey {
			apigateway_DeleteUsagePlanKey(cfg, client)
			return
		}
		if _apigatewayDeleteVpcLink {
			apigateway_DeleteVpcLink(cfg, client)
			return
		}
		if _apigatewayFlushStageAuthorizersCache {
			apigateway_FlushStageAuthorizersCache(cfg, client)
			return
		}
		if _apigatewayFlushStageCache {
			apigateway_FlushStageCache(cfg, client)
			return
		}
		if _apigatewayGenerateClientCertificate {
			apigateway_GenerateClientCertificate(cfg, client)
			return
		}
		if _apigatewayGetAccount {
			apigateway_GetAccount(cfg, client)
			return
		}
		if _apigatewayGetApiKey {
			apigateway_GetApiKey(cfg, client)
			return
		}
		if _apigatewayGetApiKeys {
			apigateway_GetApiKeys(cfg, client)
			return
		}
		if _apigatewayGetAuthorizer {
			apigateway_GetAuthorizer(cfg, client)
			return
		}
		if _apigatewayGetAuthorizers {
			apigateway_GetAuthorizers(cfg, client)
			return
		}
		if _apigatewayGetBasePathMapping {
			apigateway_GetBasePathMapping(cfg, client)
			return
		}
		if _apigatewayGetBasePathMappings {
			apigateway_GetBasePathMappings(cfg, client)
			return
		}
		if _apigatewayGetClientCertificate {
			apigateway_GetClientCertificate(cfg, client)
			return
		}
		if _apigatewayGetClientCertificates {
			apigateway_GetClientCertificates(cfg, client)
			return
		}
		if _apigatewayGetDeployment {
			apigateway_GetDeployment(cfg, client)
			return
		}
		if _apigatewayGetDeployments {
			apigateway_GetDeployments(cfg, client)
			return
		}
		if _apigatewayGetDocumentationPart {
			apigateway_GetDocumentationPart(cfg, client)
			return
		}
		if _apigatewayGetDocumentationParts {
			apigateway_GetDocumentationParts(cfg, client)
			return
		}
		if _apigatewayGetDocumentationVersion {
			apigateway_GetDocumentationVersion(cfg, client)
			return
		}
		if _apigatewayGetDocumentationVersions {
			apigateway_GetDocumentationVersions(cfg, client)
			return
		}
		if _apigatewayGetDomainName {
			apigateway_GetDomainName(cfg, client)
			return
		}
		if _apigatewayGetDomainNameAccessAssociations {
			apigateway_GetDomainNameAccessAssociations(cfg, client)
			return
		}
		if _apigatewayGetDomainNames {
			apigateway_GetDomainNames(cfg, client)
			return
		}
		if _apigatewayGetExport {
			apigateway_GetExport(cfg, client)
			return
		}
		if _apigatewayGetGatewayResponse {
			apigateway_GetGatewayResponse(cfg, client)
			return
		}
		if _apigatewayGetGatewayResponses {
			apigateway_GetGatewayResponses(cfg, client)
			return
		}
		if _apigatewayGetIntegration {
			apigateway_GetIntegration(cfg, client)
			return
		}
		if _apigatewayGetIntegrationResponse {
			apigateway_GetIntegrationResponse(cfg, client)
			return
		}
		if _apigatewayGetMethod {
			apigateway_GetMethod(cfg, client)
			return
		}
		if _apigatewayGetMethodResponse {
			apigateway_GetMethodResponse(cfg, client)
			return
		}
		if _apigatewayGetModel {
			apigateway_GetModel(cfg, client)
			return
		}
		if _apigatewayGetModelTemplate {
			apigateway_GetModelTemplate(cfg, client)
			return
		}
		if _apigatewayGetModels {
			apigateway_GetModels(cfg, client)
			return
		}
		if _apigatewayGetRequestValidator {
			apigateway_GetRequestValidator(cfg, client)
			return
		}
		if _apigatewayGetRequestValidators {
			apigateway_GetRequestValidators(cfg, client)
			return
		}
		if _apigatewayGetResource {
			apigateway_GetResource(cfg, client)
			return
		}
		if _apigatewayGetResources {
			apigateway_GetResources(cfg, client)
			return
		}
		if _apigatewayGetRestApi {
			apigateway_GetRestApi(cfg, client)
			return
		}
		if _apigatewayGetRestApis {
			apigateway_GetRestApis(cfg, client)
			return
		}
		if _apigatewayGetSdk {
			apigateway_GetSdk(cfg, client)
			return
		}
		if _apigatewayGetSdkType {
			apigateway_GetSdkType(cfg, client)
			return
		}
		if _apigatewayGetSdkTypes {
			apigateway_GetSdkTypes(cfg, client)
			return
		}
		if _apigatewayGetStage {
			apigateway_GetStage(cfg, client)
			return
		}
		if _apigatewayGetStages {
			apigateway_GetStages(cfg, client)
			return
		}
		if _apigatewayGetTags {
			apigateway_GetTags(cfg, client)
			return
		}
		if _apigatewayGetUsage {
			apigateway_GetUsage(cfg, client)
			return
		}
		if _apigatewayGetUsagePlan {
			apigateway_GetUsagePlan(cfg, client)
			return
		}
		if _apigatewayGetUsagePlanKey {
			apigateway_GetUsagePlanKey(cfg, client)
			return
		}
		if _apigatewayGetUsagePlanKeys {
			apigateway_GetUsagePlanKeys(cfg, client)
			return
		}
		if _apigatewayGetUsagePlans {
			apigateway_GetUsagePlans(cfg, client)
			return
		}
		if _apigatewayGetVpcLink {
			apigateway_GetVpcLink(cfg, client)
			return
		}
		if _apigatewayGetVpcLinks {
			apigateway_GetVpcLinks(cfg, client)
			return
		}
		if _apigatewayImportApiKeys {
			apigateway_ImportApiKeys(cfg, client)
			return
		}
		if _apigatewayImportDocumentationParts {
			apigateway_ImportDocumentationParts(cfg, client)
			return
		}
		if _apigatewayImportRestApi {
			apigateway_ImportRestApi(cfg, client)
			return
		}
		if _apigatewayPutGatewayResponse {
			apigateway_PutGatewayResponse(cfg, client)
			return
		}
		if _apigatewayPutIntegration {
			apigateway_PutIntegration(cfg, client)
			return
		}
		if _apigatewayPutIntegrationResponse {
			apigateway_PutIntegrationResponse(cfg, client)
			return
		}
		if _apigatewayPutMethod {
			apigateway_PutMethod(cfg, client)
			return
		}
		if _apigatewayPutMethodResponse {
			apigateway_PutMethodResponse(cfg, client)
			return
		}
		if _apigatewayPutRestApi {
			apigateway_PutRestApi(cfg, client)
			return
		}
		if _apigatewayRejectDomainNameAccessAssociation {
			apigateway_RejectDomainNameAccessAssociation(cfg, client)
			return
		}
		if _apigatewayTagResource {
			apigateway_TagResource(cfg, client)
			return
		}
		if _apigatewayTestInvokeAuthorizer {
			apigateway_TestInvokeAuthorizer(cfg, client)
			return
		}
		if _apigatewayTestInvokeMethod {
			apigateway_TestInvokeMethod(cfg, client)
			return
		}
		if _apigatewayUntagResource {
			apigateway_UntagResource(cfg, client)
			return
		}
		if _apigatewayUpdateAccount {
			apigateway_UpdateAccount(cfg, client)
			return
		}
		if _apigatewayUpdateApiKey {
			apigateway_UpdateApiKey(cfg, client)
			return
		}
		if _apigatewayUpdateAuthorizer {
			apigateway_UpdateAuthorizer(cfg, client)
			return
		}
		if _apigatewayUpdateBasePathMapping {
			apigateway_UpdateBasePathMapping(cfg, client)
			return
		}
		if _apigatewayUpdateClientCertificate {
			apigateway_UpdateClientCertificate(cfg, client)
			return
		}
		if _apigatewayUpdateDeployment {
			apigateway_UpdateDeployment(cfg, client)
			return
		}
		if _apigatewayUpdateDocumentationPart {
			apigateway_UpdateDocumentationPart(cfg, client)
			return
		}
		if _apigatewayUpdateDocumentationVersion {
			apigateway_UpdateDocumentationVersion(cfg, client)
			return
		}
		if _apigatewayUpdateDomainName {
			apigateway_UpdateDomainName(cfg, client)
			return
		}
		if _apigatewayUpdateGatewayResponse {
			apigateway_UpdateGatewayResponse(cfg, client)
			return
		}
		if _apigatewayUpdateIntegration {
			apigateway_UpdateIntegration(cfg, client)
			return
		}
		if _apigatewayUpdateIntegrationResponse {
			apigateway_UpdateIntegrationResponse(cfg, client)
			return
		}
		if _apigatewayUpdateMethod {
			apigateway_UpdateMethod(cfg, client)
			return
		}
		if _apigatewayUpdateMethodResponse {
			apigateway_UpdateMethodResponse(cfg, client)
			return
		}
		if _apigatewayUpdateModel {
			apigateway_UpdateModel(cfg, client)
			return
		}
		if _apigatewayUpdateRequestValidator {
			apigateway_UpdateRequestValidator(cfg, client)
			return
		}
		if _apigatewayUpdateResource {
			apigateway_UpdateResource(cfg, client)
			return
		}
		if _apigatewayUpdateRestApi {
			apigateway_UpdateRestApi(cfg, client)
			return
		}
		if _apigatewayUpdateStage {
			apigateway_UpdateStage(cfg, client)
			return
		}
		if _apigatewayUpdateUsage {
			apigateway_UpdateUsage(cfg, client)
			return
		}
		if _apigatewayUpdateUsagePlan {
			apigateway_UpdateUsagePlan(cfg, client)
			return
		}
		if _apigatewayUpdateVpcLink {
			apigateway_UpdateVpcLink(cfg, client)
			return
		}

	},
}

var (
	_apigatewayCreateApiKey                      bool
	_apigatewayCreateAuthorizer                  bool
	_apigatewayCreateBasePathMapping             bool
	_apigatewayCreateDeployment                  bool
	_apigatewayCreateDocumentationPart           bool
	_apigatewayCreateDocumentationVersion        bool
	_apigatewayCreateDomainName                  bool
	_apigatewayCreateDomainNameAccessAssociation bool
	_apigatewayCreateModel                       bool
	_apigatewayCreateRequestValidator            bool
	_apigatewayCreateResource                    bool
	_apigatewayCreateRestApi                     bool
	_apigatewayCreateStage                       bool
	_apigatewayCreateUsagePlan                   bool
	_apigatewayCreateUsagePlanKey                bool
	_apigatewayCreateVpcLink                     bool
	_apigatewayDeleteApiKey                      bool
	_apigatewayDeleteAuthorizer                  bool
	_apigatewayDeleteBasePathMapping             bool
	_apigatewayDeleteClientCertificate           bool
	_apigatewayDeleteDeployment                  bool
	_apigatewayDeleteDocumentationPart           bool
	_apigatewayDeleteDocumentationVersion        bool
	_apigatewayDeleteDomainName                  bool
	_apigatewayDeleteDomainNameAccessAssociation bool
	_apigatewayDeleteGatewayResponse             bool
	_apigatewayDeleteIntegration                 bool
	_apigatewayDeleteIntegrationResponse         bool
	_apigatewayDeleteMethod                      bool
	_apigatewayDeleteMethodResponse              bool
	_apigatewayDeleteModel                       bool
	_apigatewayDeleteRequestValidator            bool
	_apigatewayDeleteResource                    bool
	_apigatewayDeleteRestApi                     bool
	_apigatewayDeleteStage                       bool
	_apigatewayDeleteUsagePlan                   bool
	_apigatewayDeleteUsagePlanKey                bool
	_apigatewayDeleteVpcLink                     bool
	_apigatewayFlushStageAuthorizersCache        bool
	_apigatewayFlushStageCache                   bool
	_apigatewayGenerateClientCertificate         bool
	_apigatewayGetAccount                        bool
	_apigatewayGetApiKey                         bool
	_apigatewayGetApiKeys                        bool
	_apigatewayGetAuthorizer                     bool
	_apigatewayGetAuthorizers                    bool
	_apigatewayGetBasePathMapping                bool
	_apigatewayGetBasePathMappings               bool
	_apigatewayGetClientCertificate              bool
	_apigatewayGetClientCertificates             bool
	_apigatewayGetDeployment                     bool
	_apigatewayGetDeployments                    bool
	_apigatewayGetDocumentationPart              bool
	_apigatewayGetDocumentationParts             bool
	_apigatewayGetDocumentationVersion           bool
	_apigatewayGetDocumentationVersions          bool
	_apigatewayGetDomainName                     bool
	_apigatewayGetDomainNameAccessAssociations   bool
	_apigatewayGetDomainNames                    bool
	_apigatewayGetExport                         bool
	_apigatewayGetGatewayResponse                bool
	_apigatewayGetGatewayResponses               bool
	_apigatewayGetIntegration                    bool
	_apigatewayGetIntegrationResponse            bool
	_apigatewayGetMethod                         bool
	_apigatewayGetMethodResponse                 bool
	_apigatewayGetModel                          bool
	_apigatewayGetModelTemplate                  bool
	_apigatewayGetModels                         bool
	_apigatewayGetRequestValidator               bool
	_apigatewayGetRequestValidators              bool
	_apigatewayGetResource                       bool
	_apigatewayGetResources                      bool
	_apigatewayGetRestApi                        bool
	_apigatewayGetRestApis                       bool
	_apigatewayGetSdk                            bool
	_apigatewayGetSdkType                        bool
	_apigatewayGetSdkTypes                       bool
	_apigatewayGetStage                          bool
	_apigatewayGetStages                         bool
	_apigatewayGetTags                           bool
	_apigatewayGetUsage                          bool
	_apigatewayGetUsagePlan                      bool
	_apigatewayGetUsagePlanKey                   bool
	_apigatewayGetUsagePlanKeys                  bool
	_apigatewayGetUsagePlans                     bool
	_apigatewayGetVpcLink                        bool
	_apigatewayGetVpcLinks                       bool
	_apigatewayImportApiKeys                     bool
	_apigatewayImportDocumentationParts          bool
	_apigatewayImportRestApi                     bool
	_apigatewayPutGatewayResponse                bool
	_apigatewayPutIntegration                    bool
	_apigatewayPutIntegrationResponse            bool
	_apigatewayPutMethod                         bool
	_apigatewayPutMethodResponse                 bool
	_apigatewayPutRestApi                        bool
	_apigatewayRejectDomainNameAccessAssociation bool
	_apigatewayTagResource                       bool
	_apigatewayTestInvokeAuthorizer              bool
	_apigatewayTestInvokeMethod                  bool
	_apigatewayUntagResource                     bool
	_apigatewayUpdateAccount                     bool
	_apigatewayUpdateApiKey                      bool
	_apigatewayUpdateAuthorizer                  bool
	_apigatewayUpdateBasePathMapping             bool
	_apigatewayUpdateClientCertificate           bool
	_apigatewayUpdateDeployment                  bool
	_apigatewayUpdateDocumentationPart           bool
	_apigatewayUpdateDocumentationVersion        bool
	_apigatewayUpdateDomainName                  bool
	_apigatewayUpdateGatewayResponse             bool
	_apigatewayUpdateIntegration                 bool
	_apigatewayUpdateIntegrationResponse         bool
	_apigatewayUpdateMethod                      bool
	_apigatewayUpdateMethodResponse              bool
	_apigatewayUpdateModel                       bool
	_apigatewayUpdateRequestValidator            bool
	_apigatewayUpdateResource                    bool
	_apigatewayUpdateRestApi                     bool
	_apigatewayUpdateStage                       bool
	_apigatewayUpdateUsage                       bool
	_apigatewayUpdateUsagePlan                   bool
	_apigatewayUpdateVpcLink                     bool

	_apigatewayAccepts                             string
	_apigatewayAccessAssociationSource             string
	_apigatewayAccessAssociationSourceType         string
	_apigatewayAdditionalContext                   string
	_apigatewayApiKey                              string
	_apigatewayApiKeyRequired                      string
	_apigatewayApiKeySource                        string
	_apigatewayApiStages                           string
	_apigatewayAuthType                            string
	_apigatewayAuthorizationScopes                 []string
	_apigatewayAuthorizationType                   string
	_apigatewayAuthorizerCredentials               string
	_apigatewayAuthorizerId                        string
	_apigatewayAuthorizerResultTtlInSeconds        string
	_apigatewayAuthorizerUri                       string
	_apigatewayBasePath                            string
	_apigatewayBinaryMediaTypes                    []string
	_apigatewayBody                                string
	_apigatewayCacheClusterEnabled                 string
	_apigatewayCacheClusterSize                    string
	_apigatewayCacheKeyParameters                  []string
	_apigatewayCacheNamespace                      string
	_apigatewayCanarySettings                      string
	_apigatewayCertificateArn                      string
	_apigatewayCertificateBody                     string
	_apigatewayCertificateChain                    string
	_apigatewayCertificateName                     string
	_apigatewayCertificatePrivateKey               string
	_apigatewayClientCertificateId                 string
	_apigatewayCloneFrom                           string
	_apigatewayConnectionId                        string
	_apigatewayConnectionType                      string
	_apigatewayContentHandling                     string
	_apigatewayContentType                         string
	_apigatewayCredentials                         string
	_apigatewayCustomerId                          string
	_apigatewayDeploymentId                        string
	_apigatewayDescription                         string
	_apigatewayDisableExecuteApiEndpoint           string
	_apigatewayDocumentationPartId                 string
	_apigatewayDocumentationVersion                string
	_apigatewayDomainName                          string
	_apigatewayDomainNameAccessAssociationArn      string
	_apigatewayDomainNameArn                       string
	_apigatewayDomainNameId                        string
	_apigatewayEmbed                               []string
	_apigatewayEnabled                             string
	_apigatewayEndDate                             string
	_apigatewayEndpointAccessMode                  string
	_apigatewayEndpointConfiguration               string
	_apigatewayExportType                          string
	_apigatewayFailOnWarnings                      string
	_apigatewayFlatten                             string
	_apigatewayFormat                              string
	_apigatewayGenerateDistinctId                  string
	_apigatewayHeaders                             string
	_apigatewayHttpMethod                          string
	_apigatewayId                                  string
	_apigatewayIdentitySource                      string
	_apigatewayIdentityValidationExpression        string
	_apigatewayIncludeValue                        string
	_apigatewayIncludeValues                       string
	_apigatewayIntegrationHttpMethod               string
	_apigatewayIntegrationTarget                   string
	_apigatewayKeyId                               string
	_apigatewayKeyType                             string
	_apigatewayLimit                               string
	_apigatewayLocation                            string
	_apigatewayLocationStatus                      string
	_apigatewayMinimumCompressionSize              string
	_apigatewayMode                                string
	_apigatewayModelName                           string
	_apigatewayMultiValueHeaders                   string
	_apigatewayMutualTlsAuthentication             string
	_apigatewayName                                string
	_apigatewayNameQuery                           string
	_apigatewayOperationName                       string
	_apigatewayOwnershipVerificationCertificateArn string
	_apigatewayParameters                          string
	_apigatewayParentId                            string
	_apigatewayPassthroughBehavior                 string
	_apigatewayPatchOperations                     string
	_apigatewayPath                                string
	_apigatewayPathPart                            string
	_apigatewayPathWithQueryString                 string
	_apigatewayPolicy                              string
	_apigatewayPosition                            string
	_apigatewayProperties                          string
	_apigatewayProviderARNs                        []string
	_apigatewayQuota                               string
	_apigatewayRegionalCertificateArn              string
	_apigatewayRegionalCertificateName             string
	_apigatewayRequestModels                       string
	_apigatewayRequestParameters                   string
	_apigatewayRequestTemplates                    string
	_apigatewayRequestValidatorId                  string
	_apigatewayResourceArn                         string
	_apigatewayResourceId                          string
	_apigatewayResourceOwner                       string
	_apigatewayResponseModels                      string
	_apigatewayResponseParameters                  string
	_apigatewayResponseTemplates                   string
	_apigatewayResponseTransferMode                string
	_apigatewayResponseType                        string
	_apigatewayRestApiId                           string
	_apigatewayRoutingMode                         string
	_apigatewaySchema                              string
	_apigatewaySdkType                             string
	_apigatewaySecurityPolicy                      string
	_apigatewaySelectionPattern                    string
	_apigatewayStage                               string
	_apigatewayStageDescription                    string
	_apigatewayStageKeys                           string
	_apigatewayStageName                           string
	_apigatewayStageVariables                      string
	_apigatewayStartDate                           string
	_apigatewayStatusCode                          string
	_apigatewayTagKeys                             []string
	_apigatewayTags                                string
	_apigatewayTargetArns                          []string
	_apigatewayThrottle                            string
	_apigatewayTimeoutInMillis                     string
	_apigatewayTlsConfig                           string
	_apigatewayTracingEnabled                      string
	_apigatewayType                                string
	_apigatewayUri                                 string
	_apigatewayUsagePlanId                         string
	_apigatewayValidateRequestBody                 string
	_apigatewayValidateRequestParameters           string
	_apigatewayValue                               string
	_apigatewayVariables                           string
	_apigatewayVersion                             string
	_apigatewayVpcLinkId                           string
)

// Create an ApiKey resource.
func apigateway_CreateApiKey(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.CreateApiKeyInput{}

	if len(_apigatewayCustomerId) > 0 {
		input.CustomerId = aws.String(_apigatewayCustomerId)
	}
	if len(_apigatewayDescription) > 0 {
		input.Description = aws.String(_apigatewayDescription)
	}
	if len(_apigatewayEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _apigatewayEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_apigatewayGenerateDistinctId) > 0 {
		if err := assignInputField(input, "GenerateDistinctId", _apigatewayGenerateDistinctId); err != nil {
			log.Errorf("invalid --generate-distinct-id: %s", err.Error())
			return
		}
	}
	if len(_apigatewayName) > 0 {
		input.Name = aws.String(_apigatewayName)
	}
	if len(_apigatewayStageKeys) > 0 {
		if err := assignInputField(input, "StageKeys", _apigatewayStageKeys); err != nil {
			log.Errorf("invalid --stage-keys: %s", err.Error())
			return
		}
	}
	if len(_apigatewayTags) > 0 {
		if err := assignInputField(input, "Tags", _apigatewayTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_apigatewayValue) > 0 {
		input.Value = aws.String(_apigatewayValue)
	}

	if resp, err := client.CreateApiKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a new Authorizer resource to an existing RestApi resource.
func apigateway_CreateAuthorizer(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.CreateAuthorizerInput{
		// Name: *string, // Required
		// RestApiId: *string, // Required
		// Type: types.AuthorizerType, // Required
	}

	if len(_apigatewayName) > 0 {
		input.Name = aws.String(_apigatewayName)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayType) > 0 {
		if err := assignInputField(input, "Type", _apigatewayType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_apigatewayAuthType) > 0 {
		input.AuthType = aws.String(_apigatewayAuthType)
	}
	if len(_apigatewayAuthorizerCredentials) > 0 {
		input.AuthorizerCredentials = aws.String(_apigatewayAuthorizerCredentials)
	}
	if len(_apigatewayAuthorizerResultTtlInSeconds) > 0 {
		if err := assignInputField(input, "AuthorizerResultTtlInSeconds", _apigatewayAuthorizerResultTtlInSeconds); err != nil {
			log.Errorf("invalid --authorizer-result-ttl-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_apigatewayAuthorizerUri) > 0 {
		input.AuthorizerUri = aws.String(_apigatewayAuthorizerUri)
	}
	if len(_apigatewayIdentitySource) > 0 {
		input.IdentitySource = aws.String(_apigatewayIdentitySource)
	}
	if len(_apigatewayIdentityValidationExpression) > 0 {
		input.IdentityValidationExpression = aws.String(_apigatewayIdentityValidationExpression)
	}
	if len(_apigatewayProviderARNs) > 0 {
		input.ProviderARNs = append([]string(nil), _apigatewayProviderARNs...)
	}

	if resp, err := client.CreateAuthorizer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new BasePathMapping resource.
func apigateway_CreateBasePathMapping(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.CreateBasePathMappingInput{
		// DomainName: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayDomainName) > 0 {
		input.DomainName = aws.String(_apigatewayDomainName)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayBasePath) > 0 {
		input.BasePath = aws.String(_apigatewayBasePath)
	}
	if len(_apigatewayDomainNameId) > 0 {
		input.DomainNameId = aws.String(_apigatewayDomainNameId)
	}
	if len(_apigatewayStage) > 0 {
		input.Stage = aws.String(_apigatewayStage)
	}

	if resp, err := client.CreateBasePathMapping(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Deployment resource, which makes a specified RestApi callable over
// the internet.
func apigateway_CreateDeployment(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.CreateDeploymentInput{
		// RestApiId: *string, // Required
	}

	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayCacheClusterEnabled) > 0 {
		if err := assignInputField(input, "CacheClusterEnabled", _apigatewayCacheClusterEnabled); err != nil {
			log.Errorf("invalid --cache-cluster-enabled: %s", err.Error())
			return
		}
	}
	if len(_apigatewayCacheClusterSize) > 0 {
		if err := assignInputField(input, "CacheClusterSize", _apigatewayCacheClusterSize); err != nil {
			log.Errorf("invalid --cache-cluster-size: %s", err.Error())
			return
		}
	}
	if len(_apigatewayCanarySettings) > 0 {
		if err := assignInputField(input, "CanarySettings", _apigatewayCanarySettings); err != nil {
			log.Errorf("invalid --canary-settings: %s", err.Error())
			return
		}
	}
	if len(_apigatewayDescription) > 0 {
		input.Description = aws.String(_apigatewayDescription)
	}
	if len(_apigatewayStageDescription) > 0 {
		input.StageDescription = aws.String(_apigatewayStageDescription)
	}
	if len(_apigatewayStageName) > 0 {
		input.StageName = aws.String(_apigatewayStageName)
	}
	if len(_apigatewayTracingEnabled) > 0 {
		if err := assignInputField(input, "TracingEnabled", _apigatewayTracingEnabled); err != nil {
			log.Errorf("invalid --tracing-enabled: %s", err.Error())
			return
		}
	}
	if len(_apigatewayVariables) > 0 {
		if err := assignInputField(input, "Variables", _apigatewayVariables); err != nil {
			log.Errorf("invalid --variables: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a documentation part.
func apigateway_CreateDocumentationPart(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.CreateDocumentationPartInput{
		// Location: *types.DocumentationPartLocation, // Required
		// Properties: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayLocation) > 0 {
		if err := assignInputField(input, "Location", _apigatewayLocation); err != nil {
			log.Errorf("invalid --location: %s", err.Error())
			return
		}
	}
	if len(_apigatewayProperties) > 0 {
		input.Properties = aws.String(_apigatewayProperties)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}

	if resp, err := client.CreateDocumentationPart(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a documentation version
func apigateway_CreateDocumentationVersion(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.CreateDocumentationVersionInput{
		// DocumentationVersion: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayDocumentationVersion) > 0 {
		input.DocumentationVersion = aws.String(_apigatewayDocumentationVersion)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayDescription) > 0 {
		input.Description = aws.String(_apigatewayDescription)
	}
	if len(_apigatewayStageName) > 0 {
		input.StageName = aws.String(_apigatewayStageName)
	}

	if resp, err := client.CreateDocumentationVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new domain name.
func apigateway_CreateDomainName(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.CreateDomainNameInput{
		// DomainName: *string, // Required
	}

	if len(_apigatewayDomainName) > 0 {
		input.DomainName = aws.String(_apigatewayDomainName)
	}
	if len(_apigatewayCertificateArn) > 0 {
		input.CertificateArn = aws.String(_apigatewayCertificateArn)
	}
	if len(_apigatewayCertificateBody) > 0 {
		input.CertificateBody = aws.String(_apigatewayCertificateBody)
	}
	if len(_apigatewayCertificateChain) > 0 {
		input.CertificateChain = aws.String(_apigatewayCertificateChain)
	}
	if len(_apigatewayCertificateName) > 0 {
		input.CertificateName = aws.String(_apigatewayCertificateName)
	}
	if len(_apigatewayCertificatePrivateKey) > 0 {
		input.CertificatePrivateKey = aws.String(_apigatewayCertificatePrivateKey)
	}
	if len(_apigatewayEndpointAccessMode) > 0 {
		if err := assignInputField(input, "EndpointAccessMode", _apigatewayEndpointAccessMode); err != nil {
			log.Errorf("invalid --endpoint-access-mode: %s", err.Error())
			return
		}
	}
	if len(_apigatewayEndpointConfiguration) > 0 {
		if err := assignInputField(input, "EndpointConfiguration", _apigatewayEndpointConfiguration); err != nil {
			log.Errorf("invalid --endpoint-configuration: %s", err.Error())
			return
		}
	}
	if len(_apigatewayMutualTlsAuthentication) > 0 {
		if err := assignInputField(input, "MutualTlsAuthentication", _apigatewayMutualTlsAuthentication); err != nil {
			log.Errorf("invalid --mutual-tls-authentication: %s", err.Error())
			return
		}
	}
	if len(_apigatewayOwnershipVerificationCertificateArn) > 0 {
		input.OwnershipVerificationCertificateArn = aws.String(_apigatewayOwnershipVerificationCertificateArn)
	}
	if len(_apigatewayPolicy) > 0 {
		input.Policy = aws.String(_apigatewayPolicy)
	}
	if len(_apigatewayRegionalCertificateArn) > 0 {
		input.RegionalCertificateArn = aws.String(_apigatewayRegionalCertificateArn)
	}
	if len(_apigatewayRegionalCertificateName) > 0 {
		input.RegionalCertificateName = aws.String(_apigatewayRegionalCertificateName)
	}
	if len(_apigatewayRoutingMode) > 0 {
		if err := assignInputField(input, "RoutingMode", _apigatewayRoutingMode); err != nil {
			log.Errorf("invalid --routing-mode: %s", err.Error())
			return
		}
	}
	if len(_apigatewaySecurityPolicy) > 0 {
		if err := assignInputField(input, "SecurityPolicy", _apigatewaySecurityPolicy); err != nil {
			log.Errorf("invalid --security-policy: %s", err.Error())
			return
		}
	}
	if len(_apigatewayTags) > 0 {
		if err := assignInputField(input, "Tags", _apigatewayTags); err != nil {
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

// Creates a domain name access association resource between an access
// association source and a private custom domain name.
func apigateway_CreateDomainNameAccessAssociation(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.CreateDomainNameAccessAssociationInput{
		// AccessAssociationSource: *string, // Required
		// AccessAssociationSourceType: types.AccessAssociationSourceType, // Required
		// DomainNameArn: *string, // Required
	}

	if len(_apigatewayAccessAssociationSource) > 0 {
		input.AccessAssociationSource = aws.String(_apigatewayAccessAssociationSource)
	}
	if len(_apigatewayAccessAssociationSourceType) > 0 {
		if err := assignInputField(input, "AccessAssociationSourceType", _apigatewayAccessAssociationSourceType); err != nil {
			log.Errorf("invalid --access-association-source-type: %s", err.Error())
			return
		}
	}
	if len(_apigatewayDomainNameArn) > 0 {
		input.DomainNameArn = aws.String(_apigatewayDomainNameArn)
	}
	if len(_apigatewayTags) > 0 {
		if err := assignInputField(input, "Tags", _apigatewayTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDomainNameAccessAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a new Model resource to an existing RestApi resource.
func apigateway_CreateModel(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.CreateModelInput{
		// ContentType: *string, // Required
		// Name: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayContentType) > 0 {
		input.ContentType = aws.String(_apigatewayContentType)
	}
	if len(_apigatewayName) > 0 {
		input.Name = aws.String(_apigatewayName)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayDescription) > 0 {
		input.Description = aws.String(_apigatewayDescription)
	}
	if len(_apigatewaySchema) > 0 {
		input.Schema = aws.String(_apigatewaySchema)
	}

	if resp, err := client.CreateModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a RequestValidator of a given RestApi.
func apigateway_CreateRequestValidator(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.CreateRequestValidatorInput{
		// RestApiId: *string, // Required
	}

	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayName) > 0 {
		input.Name = aws.String(_apigatewayName)
	}
	if len(_apigatewayValidateRequestBody) > 0 {
		if err := assignInputField(input, "ValidateRequestBody", _apigatewayValidateRequestBody); err != nil {
			log.Errorf("invalid --validate-request-body: %s", err.Error())
			return
		}
	}
	if len(_apigatewayValidateRequestParameters) > 0 {
		if err := assignInputField(input, "ValidateRequestParameters", _apigatewayValidateRequestParameters); err != nil {
			log.Errorf("invalid --validate-request-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRequestValidator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Resource resource.
func apigateway_CreateResource(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.CreateResourceInput{
		// ParentId: *string, // Required
		// PathPart: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayParentId) > 0 {
		input.ParentId = aws.String(_apigatewayParentId)
	}
	if len(_apigatewayPathPart) > 0 {
		input.PathPart = aws.String(_apigatewayPathPart)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}

	if resp, err := client.CreateResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new RestApi resource.
func apigateway_CreateRestApi(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.CreateRestApiInput{
		// Name: *string, // Required
	}

	if len(_apigatewayName) > 0 {
		input.Name = aws.String(_apigatewayName)
	}
	if len(_apigatewayApiKeySource) > 0 {
		if err := assignInputField(input, "ApiKeySource", _apigatewayApiKeySource); err != nil {
			log.Errorf("invalid --api-key-source: %s", err.Error())
			return
		}
	}
	if len(_apigatewayBinaryMediaTypes) > 0 {
		input.BinaryMediaTypes = append([]string(nil), _apigatewayBinaryMediaTypes...)
	}
	if len(_apigatewayCloneFrom) > 0 {
		input.CloneFrom = aws.String(_apigatewayCloneFrom)
	}
	if len(_apigatewayDescription) > 0 {
		input.Description = aws.String(_apigatewayDescription)
	}
	if len(_apigatewayDisableExecuteApiEndpoint) > 0 {
		if err := assignInputField(input, "DisableExecuteApiEndpoint", _apigatewayDisableExecuteApiEndpoint); err != nil {
			log.Errorf("invalid --disable-execute-api-endpoint: %s", err.Error())
			return
		}
	}
	if len(_apigatewayEndpointAccessMode) > 0 {
		if err := assignInputField(input, "EndpointAccessMode", _apigatewayEndpointAccessMode); err != nil {
			log.Errorf("invalid --endpoint-access-mode: %s", err.Error())
			return
		}
	}
	if len(_apigatewayEndpointConfiguration) > 0 {
		if err := assignInputField(input, "EndpointConfiguration", _apigatewayEndpointConfiguration); err != nil {
			log.Errorf("invalid --endpoint-configuration: %s", err.Error())
			return
		}
	}
	if len(_apigatewayMinimumCompressionSize) > 0 {
		if err := assignInputField(input, "MinimumCompressionSize", _apigatewayMinimumCompressionSize); err != nil {
			log.Errorf("invalid --minimum-compression-size: %s", err.Error())
			return
		}
	}
	if len(_apigatewayPolicy) > 0 {
		input.Policy = aws.String(_apigatewayPolicy)
	}
	if len(_apigatewaySecurityPolicy) > 0 {
		if err := assignInputField(input, "SecurityPolicy", _apigatewaySecurityPolicy); err != nil {
			log.Errorf("invalid --security-policy: %s", err.Error())
			return
		}
	}
	if len(_apigatewayTags) > 0 {
		if err := assignInputField(input, "Tags", _apigatewayTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_apigatewayVersion) > 0 {
		input.Version = aws.String(_apigatewayVersion)
	}

	if resp, err := client.CreateRestApi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Stage resource that references a pre-existing Deployment for the
// API.
func apigateway_CreateStage(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.CreateStageInput{
		// DeploymentId: *string, // Required
		// RestApiId: *string, // Required
		// StageName: *string, // Required
	}

	if len(_apigatewayDeploymentId) > 0 {
		input.DeploymentId = aws.String(_apigatewayDeploymentId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayStageName) > 0 {
		input.StageName = aws.String(_apigatewayStageName)
	}
	if len(_apigatewayCacheClusterEnabled) > 0 {
		if err := assignInputField(input, "CacheClusterEnabled", _apigatewayCacheClusterEnabled); err != nil {
			log.Errorf("invalid --cache-cluster-enabled: %s", err.Error())
			return
		}
	}
	if len(_apigatewayCacheClusterSize) > 0 {
		if err := assignInputField(input, "CacheClusterSize", _apigatewayCacheClusterSize); err != nil {
			log.Errorf("invalid --cache-cluster-size: %s", err.Error())
			return
		}
	}
	if len(_apigatewayCanarySettings) > 0 {
		if err := assignInputField(input, "CanarySettings", _apigatewayCanarySettings); err != nil {
			log.Errorf("invalid --canary-settings: %s", err.Error())
			return
		}
	}
	if len(_apigatewayDescription) > 0 {
		input.Description = aws.String(_apigatewayDescription)
	}
	if len(_apigatewayDocumentationVersion) > 0 {
		input.DocumentationVersion = aws.String(_apigatewayDocumentationVersion)
	}
	if len(_apigatewayTags) > 0 {
		if err := assignInputField(input, "Tags", _apigatewayTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_apigatewayTracingEnabled) > 0 {
		if err := assignInputField(input, "TracingEnabled", _apigatewayTracingEnabled); err != nil {
			log.Errorf("invalid --tracing-enabled: %s", err.Error())
			return
		}
	}
	if len(_apigatewayVariables) > 0 {
		if err := assignInputField(input, "Variables", _apigatewayVariables); err != nil {
			log.Errorf("invalid --variables: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateStage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a usage plan with the throttle and quota limits, as well as the
// associated API stages, specified in the payload.
func apigateway_CreateUsagePlan(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.CreateUsagePlanInput{
		// Name: *string, // Required
	}

	if len(_apigatewayName) > 0 {
		input.Name = aws.String(_apigatewayName)
	}
	if len(_apigatewayApiStages) > 0 {
		if err := assignInputField(input, "ApiStages", _apigatewayApiStages); err != nil {
			log.Errorf("invalid --api-stages: %s", err.Error())
			return
		}
	}
	if len(_apigatewayDescription) > 0 {
		input.Description = aws.String(_apigatewayDescription)
	}
	if len(_apigatewayQuota) > 0 {
		if err := assignInputField(input, "Quota", _apigatewayQuota); err != nil {
			log.Errorf("invalid --quota: %s", err.Error())
			return
		}
	}
	if len(_apigatewayTags) > 0 {
		if err := assignInputField(input, "Tags", _apigatewayTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_apigatewayThrottle) > 0 {
		if err := assignInputField(input, "Throttle", _apigatewayThrottle); err != nil {
			log.Errorf("invalid --throttle: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateUsagePlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a usage plan key for adding an existing API key to a usage plan.
func apigateway_CreateUsagePlanKey(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.CreateUsagePlanKeyInput{
		// KeyId: *string, // Required
		// KeyType: *string, // Required
		// UsagePlanId: *string, // Required
	}

	if len(_apigatewayKeyId) > 0 {
		input.KeyId = aws.String(_apigatewayKeyId)
	}
	if len(_apigatewayKeyType) > 0 {
		input.KeyType = aws.String(_apigatewayKeyType)
	}
	if len(_apigatewayUsagePlanId) > 0 {
		input.UsagePlanId = aws.String(_apigatewayUsagePlanId)
	}

	if resp, err := client.CreateUsagePlanKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a VPC link, under the caller's account in a selected region, in an
// asynchronous operation that typically takes 2-4 minutes to complete and become
// operational. The caller must have permissions to create and update VPC Endpoint
// services.
func apigateway_CreateVpcLink(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.CreateVpcLinkInput{
		// Name: *string, // Required
		// TargetArns: []string, // Required
	}

	if len(_apigatewayName) > 0 {
		input.Name = aws.String(_apigatewayName)
	}
	if len(_apigatewayTargetArns) > 0 {
		input.TargetArns = append([]string(nil), _apigatewayTargetArns...)
	}
	if len(_apigatewayDescription) > 0 {
		input.Description = aws.String(_apigatewayDescription)
	}
	if len(_apigatewayTags) > 0 {
		if err := assignInputField(input, "Tags", _apigatewayTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateVpcLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the ApiKey resource.
func apigateway_DeleteApiKey(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.DeleteApiKeyInput{
		// ApiKey: *string, // Required
	}

	if len(_apigatewayApiKey) > 0 {
		input.ApiKey = aws.String(_apigatewayApiKey)
	}

	if resp, err := client.DeleteApiKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing Authorizer resource.
func apigateway_DeleteAuthorizer(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.DeleteAuthorizerInput{
		// AuthorizerId: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayAuthorizerId) > 0 {
		input.AuthorizerId = aws.String(_apigatewayAuthorizerId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}

	if resp, err := client.DeleteAuthorizer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the BasePathMapping resource.
func apigateway_DeleteBasePathMapping(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.DeleteBasePathMappingInput{
		// BasePath: *string, // Required
		// DomainName: *string, // Required
	}

	if len(_apigatewayBasePath) > 0 {
		input.BasePath = aws.String(_apigatewayBasePath)
	}
	if len(_apigatewayDomainName) > 0 {
		input.DomainName = aws.String(_apigatewayDomainName)
	}
	if len(_apigatewayDomainNameId) > 0 {
		input.DomainNameId = aws.String(_apigatewayDomainNameId)
	}

	if resp, err := client.DeleteBasePathMapping(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the ClientCertificate resource.
func apigateway_DeleteClientCertificate(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.DeleteClientCertificateInput{
		// ClientCertificateId: *string, // Required
	}

	if len(_apigatewayClientCertificateId) > 0 {
		input.ClientCertificateId = aws.String(_apigatewayClientCertificateId)
	}

	if resp, err := client.DeleteClientCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Deployment resource. Deleting a deployment will only succeed if there
// are no Stage resources associated with it.
func apigateway_DeleteDeployment(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.DeleteDeploymentInput{
		// DeploymentId: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayDeploymentId) > 0 {
		input.DeploymentId = aws.String(_apigatewayDeploymentId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}

	if resp, err := client.DeleteDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a documentation part
func apigateway_DeleteDocumentationPart(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.DeleteDocumentationPartInput{
		// DocumentationPartId: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayDocumentationPartId) > 0 {
		input.DocumentationPartId = aws.String(_apigatewayDocumentationPartId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}

	if resp, err := client.DeleteDocumentationPart(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a documentation version.
func apigateway_DeleteDocumentationVersion(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.DeleteDocumentationVersionInput{
		// DocumentationVersion: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayDocumentationVersion) > 0 {
		input.DocumentationVersion = aws.String(_apigatewayDocumentationVersion)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}

	if resp, err := client.DeleteDocumentationVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the DomainName resource.
func apigateway_DeleteDomainName(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.DeleteDomainNameInput{
		// DomainName: *string, // Required
	}

	if len(_apigatewayDomainName) > 0 {
		input.DomainName = aws.String(_apigatewayDomainName)
	}
	if len(_apigatewayDomainNameId) > 0 {
		input.DomainNameId = aws.String(_apigatewayDomainNameId)
	}

	if resp, err := client.DeleteDomainName(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the DomainNameAccessAssociation resource.
// Only the AWS account that created the DomainNameAccessAssociation resource can
// delete it. To stop an access association source in another AWS account from
// accessing your private custom domain name, use the
// RejectDomainNameAccessAssociation operation.
func apigateway_DeleteDomainNameAccessAssociation(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.DeleteDomainNameAccessAssociationInput{
		// DomainNameAccessAssociationArn: *string, // Required
	}

	if len(_apigatewayDomainNameAccessAssociationArn) > 0 {
		input.DomainNameAccessAssociationArn = aws.String(_apigatewayDomainNameAccessAssociationArn)
	}

	if resp, err := client.DeleteDomainNameAccessAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Clears any customization of a GatewayResponse of a specified response type on
// the given RestApi and resets it with the default settings.
func apigateway_DeleteGatewayResponse(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.DeleteGatewayResponseInput{
		// ResponseType: types.GatewayResponseType, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayResponseType) > 0 {
		if err := assignInputField(input, "ResponseType", _apigatewayResponseType); err != nil {
			log.Errorf("invalid --response-type: %s", err.Error())
			return
		}
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}

	if resp, err := client.DeleteGatewayResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Represents a delete integration.
func apigateway_DeleteIntegration(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.DeleteIntegrationInput{
		// HttpMethod: *string, // Required
		// ResourceId: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayHttpMethod) > 0 {
		input.HttpMethod = aws.String(_apigatewayHttpMethod)
	}
	if len(_apigatewayResourceId) > 0 {
		input.ResourceId = aws.String(_apigatewayResourceId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}

	if resp, err := client.DeleteIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Represents a delete integration response.
func apigateway_DeleteIntegrationResponse(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.DeleteIntegrationResponseInput{
		// HttpMethod: *string, // Required
		// ResourceId: *string, // Required
		// RestApiId: *string, // Required
		// StatusCode: *string, // Required
	}

	if len(_apigatewayHttpMethod) > 0 {
		input.HttpMethod = aws.String(_apigatewayHttpMethod)
	}
	if len(_apigatewayResourceId) > 0 {
		input.ResourceId = aws.String(_apigatewayResourceId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayStatusCode) > 0 {
		input.StatusCode = aws.String(_apigatewayStatusCode)
	}

	if resp, err := client.DeleteIntegrationResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing Method resource.
func apigateway_DeleteMethod(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.DeleteMethodInput{
		// HttpMethod: *string, // Required
		// ResourceId: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayHttpMethod) > 0 {
		input.HttpMethod = aws.String(_apigatewayHttpMethod)
	}
	if len(_apigatewayResourceId) > 0 {
		input.ResourceId = aws.String(_apigatewayResourceId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}

	if resp, err := client.DeleteMethod(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing MethodResponse resource.
func apigateway_DeleteMethodResponse(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.DeleteMethodResponseInput{
		// HttpMethod: *string, // Required
		// ResourceId: *string, // Required
		// RestApiId: *string, // Required
		// StatusCode: *string, // Required
	}

	if len(_apigatewayHttpMethod) > 0 {
		input.HttpMethod = aws.String(_apigatewayHttpMethod)
	}
	if len(_apigatewayResourceId) > 0 {
		input.ResourceId = aws.String(_apigatewayResourceId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayStatusCode) > 0 {
		input.StatusCode = aws.String(_apigatewayStatusCode)
	}

	if resp, err := client.DeleteMethodResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a model.
func apigateway_DeleteModel(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.DeleteModelInput{
		// ModelName: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayModelName) > 0 {
		input.ModelName = aws.String(_apigatewayModelName)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}

	if resp, err := client.DeleteModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a RequestValidator of a given RestApi.
func apigateway_DeleteRequestValidator(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.DeleteRequestValidatorInput{
		// RequestValidatorId: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayRequestValidatorId) > 0 {
		input.RequestValidatorId = aws.String(_apigatewayRequestValidatorId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}

	if resp, err := client.DeleteRequestValidator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Resource resource.
func apigateway_DeleteResource(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.DeleteResourceInput{
		// ResourceId: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayResourceId) > 0 {
		input.ResourceId = aws.String(_apigatewayResourceId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}

	if resp, err := client.DeleteResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified API.
func apigateway_DeleteRestApi(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.DeleteRestApiInput{
		// RestApiId: *string, // Required
	}

	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}

	if resp, err := client.DeleteRestApi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Stage resource.
func apigateway_DeleteStage(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.DeleteStageInput{
		// RestApiId: *string, // Required
		// StageName: *string, // Required
	}

	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayStageName) > 0 {
		input.StageName = aws.String(_apigatewayStageName)
	}

	if resp, err := client.DeleteStage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a usage plan of a given plan Id.
func apigateway_DeleteUsagePlan(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.DeleteUsagePlanInput{
		// UsagePlanId: *string, // Required
	}

	if len(_apigatewayUsagePlanId) > 0 {
		input.UsagePlanId = aws.String(_apigatewayUsagePlanId)
	}

	if resp, err := client.DeleteUsagePlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a usage plan key and remove the underlying API key from the associated
// usage plan.
func apigateway_DeleteUsagePlanKey(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.DeleteUsagePlanKeyInput{
		// KeyId: *string, // Required
		// UsagePlanId: *string, // Required
	}

	if len(_apigatewayKeyId) > 0 {
		input.KeyId = aws.String(_apigatewayKeyId)
	}
	if len(_apigatewayUsagePlanId) > 0 {
		input.UsagePlanId = aws.String(_apigatewayUsagePlanId)
	}

	if resp, err := client.DeleteUsagePlanKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing VpcLink of a specified identifier.
func apigateway_DeleteVpcLink(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.DeleteVpcLinkInput{
		// VpcLinkId: *string, // Required
	}

	if len(_apigatewayVpcLinkId) > 0 {
		input.VpcLinkId = aws.String(_apigatewayVpcLinkId)
	}

	if resp, err := client.DeleteVpcLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Flushes all authorizer cache entries on a stage.
func apigateway_FlushStageAuthorizersCache(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.FlushStageAuthorizersCacheInput{
		// RestApiId: *string, // Required
		// StageName: *string, // Required
	}

	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayStageName) > 0 {
		input.StageName = aws.String(_apigatewayStageName)
	}

	if resp, err := client.FlushStageAuthorizersCache(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Flushes a stage's cache.
func apigateway_FlushStageCache(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.FlushStageCacheInput{
		// RestApiId: *string, // Required
		// StageName: *string, // Required
	}

	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayStageName) > 0 {
		input.StageName = aws.String(_apigatewayStageName)
	}

	if resp, err := client.FlushStageCache(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates a ClientCertificate resource.
func apigateway_GenerateClientCertificate(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GenerateClientCertificateInput{}

	if len(_apigatewayDescription) > 0 {
		input.Description = aws.String(_apigatewayDescription)
	}
	if len(_apigatewayTags) > 0 {
		if err := assignInputField(input, "Tags", _apigatewayTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.GenerateClientCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the current Account resource.
func apigateway_GetAccount(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetAccountInput{}

	if resp, err := client.GetAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the current ApiKey resource.
func apigateway_GetApiKey(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetApiKeyInput{
		// ApiKey: *string, // Required
	}

	if len(_apigatewayApiKey) > 0 {
		input.ApiKey = aws.String(_apigatewayApiKey)
	}
	if len(_apigatewayIncludeValue) > 0 {
		if err := assignInputField(input, "IncludeValue", _apigatewayIncludeValue); err != nil {
			log.Errorf("invalid --include-value: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetApiKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the current ApiKeys resource.
func apigateway_GetApiKeys(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetApiKeysInput{}

	if len(_apigatewayCustomerId) > 0 {
		input.CustomerId = aws.String(_apigatewayCustomerId)
	}
	if len(_apigatewayIncludeValues) > 0 {
		if err := assignInputField(input, "IncludeValues", _apigatewayIncludeValues); err != nil {
			log.Errorf("invalid --include-values: %s", err.Error())
			return
		}
	}
	if len(_apigatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _apigatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_apigatewayNameQuery) > 0 {
		input.NameQuery = aws.String(_apigatewayNameQuery)
	}
	if len(_apigatewayPosition) > 0 {
		input.Position = aws.String(_apigatewayPosition)
	}

	if disablePaginator() {
		if resp, err := client.GetApiKeys(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*apigateway.GetApiKeysOutput
	p := apigateway.NewGetApiKeysPaginator(client, input)
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

// Describe an existing Authorizer resource.
func apigateway_GetAuthorizer(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetAuthorizerInput{
		// AuthorizerId: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayAuthorizerId) > 0 {
		input.AuthorizerId = aws.String(_apigatewayAuthorizerId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}

	if resp, err := client.GetAuthorizer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describe an existing Authorizers resource.
func apigateway_GetAuthorizers(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetAuthorizersInput{
		// RestApiId: *string, // Required
	}

	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _apigatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_apigatewayPosition) > 0 {
		input.Position = aws.String(_apigatewayPosition)
	}

	if resp, err := client.GetAuthorizers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describe a BasePathMapping resource.
func apigateway_GetBasePathMapping(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetBasePathMappingInput{
		// BasePath: *string, // Required
		// DomainName: *string, // Required
	}

	if len(_apigatewayBasePath) > 0 {
		input.BasePath = aws.String(_apigatewayBasePath)
	}
	if len(_apigatewayDomainName) > 0 {
		input.DomainName = aws.String(_apigatewayDomainName)
	}
	if len(_apigatewayDomainNameId) > 0 {
		input.DomainNameId = aws.String(_apigatewayDomainNameId)
	}

	if resp, err := client.GetBasePathMapping(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Represents a collection of BasePathMapping resources.
func apigateway_GetBasePathMappings(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetBasePathMappingsInput{
		// DomainName: *string, // Required
	}

	if len(_apigatewayDomainName) > 0 {
		input.DomainName = aws.String(_apigatewayDomainName)
	}
	if len(_apigatewayDomainNameId) > 0 {
		input.DomainNameId = aws.String(_apigatewayDomainNameId)
	}
	if len(_apigatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _apigatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_apigatewayPosition) > 0 {
		input.Position = aws.String(_apigatewayPosition)
	}

	if disablePaginator() {
		if resp, err := client.GetBasePathMappings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*apigateway.GetBasePathMappingsOutput
	p := apigateway.NewGetBasePathMappingsPaginator(client, input)
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

// Gets information about the current ClientCertificate resource.
func apigateway_GetClientCertificate(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetClientCertificateInput{
		// ClientCertificateId: *string, // Required
	}

	if len(_apigatewayClientCertificateId) > 0 {
		input.ClientCertificateId = aws.String(_apigatewayClientCertificateId)
	}

	if resp, err := client.GetClientCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a collection of ClientCertificate resources.
func apigateway_GetClientCertificates(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetClientCertificatesInput{}

	if len(_apigatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _apigatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_apigatewayPosition) > 0 {
		input.Position = aws.String(_apigatewayPosition)
	}

	if disablePaginator() {
		if resp, err := client.GetClientCertificates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*apigateway.GetClientCertificatesOutput
	p := apigateway.NewGetClientCertificatesPaginator(client, input)
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

// Gets information about a Deployment resource.
func apigateway_GetDeployment(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetDeploymentInput{
		// DeploymentId: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayDeploymentId) > 0 {
		input.DeploymentId = aws.String(_apigatewayDeploymentId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayEmbed) > 0 {
		input.Embed = append([]string(nil), _apigatewayEmbed...)
	}

	if resp, err := client.GetDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a Deployments collection.
func apigateway_GetDeployments(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetDeploymentsInput{
		// RestApiId: *string, // Required
	}

	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _apigatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_apigatewayPosition) > 0 {
		input.Position = aws.String(_apigatewayPosition)
	}

	if disablePaginator() {
		if resp, err := client.GetDeployments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*apigateway.GetDeploymentsOutput
	p := apigateway.NewGetDeploymentsPaginator(client, input)
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

// Gets a documentation part.
func apigateway_GetDocumentationPart(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetDocumentationPartInput{
		// DocumentationPartId: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayDocumentationPartId) > 0 {
		input.DocumentationPartId = aws.String(_apigatewayDocumentationPartId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}

	if resp, err := client.GetDocumentationPart(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets documentation parts.
func apigateway_GetDocumentationParts(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetDocumentationPartsInput{
		// RestApiId: *string, // Required
	}

	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _apigatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_apigatewayLocationStatus) > 0 {
		if err := assignInputField(input, "LocationStatus", _apigatewayLocationStatus); err != nil {
			log.Errorf("invalid --location-status: %s", err.Error())
			return
		}
	}
	if len(_apigatewayNameQuery) > 0 {
		input.NameQuery = aws.String(_apigatewayNameQuery)
	}
	if len(_apigatewayPath) > 0 {
		input.Path = aws.String(_apigatewayPath)
	}
	if len(_apigatewayPosition) > 0 {
		input.Position = aws.String(_apigatewayPosition)
	}
	if len(_apigatewayType) > 0 {
		if err := assignInputField(input, "Type", _apigatewayType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetDocumentationParts(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a documentation version.
func apigateway_GetDocumentationVersion(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetDocumentationVersionInput{
		// DocumentationVersion: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayDocumentationVersion) > 0 {
		input.DocumentationVersion = aws.String(_apigatewayDocumentationVersion)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}

	if resp, err := client.GetDocumentationVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets documentation versions.
func apigateway_GetDocumentationVersions(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetDocumentationVersionsInput{
		// RestApiId: *string, // Required
	}

	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _apigatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_apigatewayPosition) > 0 {
		input.Position = aws.String(_apigatewayPosition)
	}

	if resp, err := client.GetDocumentationVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Represents a domain name that is contained in a simpler, more intuitive URL
// that can be called.
func apigateway_GetDomainName(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetDomainNameInput{
		// DomainName: *string, // Required
	}

	if len(_apigatewayDomainName) > 0 {
		input.DomainName = aws.String(_apigatewayDomainName)
	}
	if len(_apigatewayDomainNameId) > 0 {
		input.DomainNameId = aws.String(_apigatewayDomainNameId)
	}

	if resp, err := client.GetDomainName(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Represents a collection on DomainNameAccessAssociations resources.
func apigateway_GetDomainNameAccessAssociations(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetDomainNameAccessAssociationsInput{}

	if len(_apigatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _apigatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_apigatewayPosition) > 0 {
		input.Position = aws.String(_apigatewayPosition)
	}
	if len(_apigatewayResourceOwner) > 0 {
		if err := assignInputField(input, "ResourceOwner", _apigatewayResourceOwner); err != nil {
			log.Errorf("invalid --resource-owner: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetDomainNameAccessAssociations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Represents a collection of DomainName resources.
func apigateway_GetDomainNames(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetDomainNamesInput{}

	if len(_apigatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _apigatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_apigatewayPosition) > 0 {
		input.Position = aws.String(_apigatewayPosition)
	}
	if len(_apigatewayResourceOwner) > 0 {
		if err := assignInputField(input, "ResourceOwner", _apigatewayResourceOwner); err != nil {
			log.Errorf("invalid --resource-owner: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetDomainNames(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*apigateway.GetDomainNamesOutput
	p := apigateway.NewGetDomainNamesPaginator(client, input)
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

// Exports a deployed version of a RestApi in a specified format.
func apigateway_GetExport(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetExportInput{
		// ExportType: *string, // Required
		// RestApiId: *string, // Required
		// StageName: *string, // Required
	}

	if len(_apigatewayExportType) > 0 {
		input.ExportType = aws.String(_apigatewayExportType)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayStageName) > 0 {
		input.StageName = aws.String(_apigatewayStageName)
	}
	if len(_apigatewayAccepts) > 0 {
		input.Accepts = aws.String(_apigatewayAccepts)
	}
	if len(_apigatewayParameters) > 0 {
		if err := assignInputField(input, "Parameters", _apigatewayParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a GatewayResponse of a specified response type on the given RestApi.
func apigateway_GetGatewayResponse(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetGatewayResponseInput{
		// ResponseType: types.GatewayResponseType, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayResponseType) > 0 {
		if err := assignInputField(input, "ResponseType", _apigatewayResponseType); err != nil {
			log.Errorf("invalid --response-type: %s", err.Error())
			return
		}
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}

	if resp, err := client.GetGatewayResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the GatewayResponses collection on the given RestApi. If an API developer
// has not added any definitions for gateway responses, the result will be the API
// Gateway-generated default GatewayResponses collection for the supported response
// types.
func apigateway_GetGatewayResponses(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetGatewayResponsesInput{
		// RestApiId: *string, // Required
	}

	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _apigatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_apigatewayPosition) > 0 {
		input.Position = aws.String(_apigatewayPosition)
	}

	if resp, err := client.GetGatewayResponses(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the integration settings.
func apigateway_GetIntegration(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetIntegrationInput{
		// HttpMethod: *string, // Required
		// ResourceId: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayHttpMethod) > 0 {
		input.HttpMethod = aws.String(_apigatewayHttpMethod)
	}
	if len(_apigatewayResourceId) > 0 {
		input.ResourceId = aws.String(_apigatewayResourceId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}

	if resp, err := client.GetIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Represents a get integration response.
func apigateway_GetIntegrationResponse(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetIntegrationResponseInput{
		// HttpMethod: *string, // Required
		// ResourceId: *string, // Required
		// RestApiId: *string, // Required
		// StatusCode: *string, // Required
	}

	if len(_apigatewayHttpMethod) > 0 {
		input.HttpMethod = aws.String(_apigatewayHttpMethod)
	}
	if len(_apigatewayResourceId) > 0 {
		input.ResourceId = aws.String(_apigatewayResourceId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayStatusCode) > 0 {
		input.StatusCode = aws.String(_apigatewayStatusCode)
	}

	if resp, err := client.GetIntegrationResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describe an existing Method resource.
func apigateway_GetMethod(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetMethodInput{
		// HttpMethod: *string, // Required
		// ResourceId: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayHttpMethod) > 0 {
		input.HttpMethod = aws.String(_apigatewayHttpMethod)
	}
	if len(_apigatewayResourceId) > 0 {
		input.ResourceId = aws.String(_apigatewayResourceId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}

	if resp, err := client.GetMethod(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a MethodResponse resource.
func apigateway_GetMethodResponse(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetMethodResponseInput{
		// HttpMethod: *string, // Required
		// ResourceId: *string, // Required
		// RestApiId: *string, // Required
		// StatusCode: *string, // Required
	}

	if len(_apigatewayHttpMethod) > 0 {
		input.HttpMethod = aws.String(_apigatewayHttpMethod)
	}
	if len(_apigatewayResourceId) > 0 {
		input.ResourceId = aws.String(_apigatewayResourceId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayStatusCode) > 0 {
		input.StatusCode = aws.String(_apigatewayStatusCode)
	}

	if resp, err := client.GetMethodResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an existing model defined for a RestApi resource.
func apigateway_GetModel(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetModelInput{
		// ModelName: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayModelName) > 0 {
		input.ModelName = aws.String(_apigatewayModelName)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayFlatten) > 0 {
		if err := assignInputField(input, "Flatten", _apigatewayFlatten); err != nil {
			log.Errorf("invalid --flatten: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates a sample mapping template that can be used to transform a payload
// into the structure of a model.
func apigateway_GetModelTemplate(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetModelTemplateInput{
		// ModelName: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayModelName) > 0 {
		input.ModelName = aws.String(_apigatewayModelName)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}

	if resp, err := client.GetModelTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes existing Models defined for a RestApi resource.
func apigateway_GetModels(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetModelsInput{
		// RestApiId: *string, // Required
	}

	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _apigatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_apigatewayPosition) > 0 {
		input.Position = aws.String(_apigatewayPosition)
	}

	if disablePaginator() {
		if resp, err := client.GetModels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*apigateway.GetModelsOutput
	p := apigateway.NewGetModelsPaginator(client, input)
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

// Gets a RequestValidator of a given RestApi.
func apigateway_GetRequestValidator(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetRequestValidatorInput{
		// RequestValidatorId: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayRequestValidatorId) > 0 {
		input.RequestValidatorId = aws.String(_apigatewayRequestValidatorId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}

	if resp, err := client.GetRequestValidator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the RequestValidators collection of a given RestApi.
func apigateway_GetRequestValidators(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetRequestValidatorsInput{
		// RestApiId: *string, // Required
	}

	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _apigatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_apigatewayPosition) > 0 {
		input.Position = aws.String(_apigatewayPosition)
	}

	if resp, err := client.GetRequestValidators(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists information about a resource.
func apigateway_GetResource(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetResourceInput{
		// ResourceId: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayResourceId) > 0 {
		input.ResourceId = aws.String(_apigatewayResourceId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayEmbed) > 0 {
		input.Embed = append([]string(nil), _apigatewayEmbed...)
	}

	if resp, err := client.GetResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists information about a collection of Resource resources.
func apigateway_GetResources(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetResourcesInput{
		// RestApiId: *string, // Required
	}

	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayEmbed) > 0 {
		input.Embed = append([]string(nil), _apigatewayEmbed...)
	}
	if len(_apigatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _apigatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_apigatewayPosition) > 0 {
		input.Position = aws.String(_apigatewayPosition)
	}

	if disablePaginator() {
		if resp, err := client.GetResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*apigateway.GetResourcesOutput
	p := apigateway.NewGetResourcesPaginator(client, input)
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

// Lists the RestApi resource in the collection.
func apigateway_GetRestApi(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetRestApiInput{
		// RestApiId: *string, // Required
	}

	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}

	if resp, err := client.GetRestApi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the RestApis resources for your collection.
func apigateway_GetRestApis(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetRestApisInput{}

	if len(_apigatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _apigatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_apigatewayPosition) > 0 {
		input.Position = aws.String(_apigatewayPosition)
	}

	if disablePaginator() {
		if resp, err := client.GetRestApis(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*apigateway.GetRestApisOutput
	p := apigateway.NewGetRestApisPaginator(client, input)
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

// Generates a client SDK for a RestApi and Stage.
func apigateway_GetSdk(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetSdkInput{
		// RestApiId: *string, // Required
		// SdkType: *string, // Required
		// StageName: *string, // Required
	}

	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewaySdkType) > 0 {
		input.SdkType = aws.String(_apigatewaySdkType)
	}
	if len(_apigatewayStageName) > 0 {
		input.StageName = aws.String(_apigatewayStageName)
	}
	if len(_apigatewayParameters) > 0 {
		if err := assignInputField(input, "Parameters", _apigatewayParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetSdk(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an SDK type.
func apigateway_GetSdkType(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetSdkTypeInput{
		// Id: *string, // Required
	}

	if len(_apigatewayId) > 0 {
		input.Id = aws.String(_apigatewayId)
	}

	if resp, err := client.GetSdkType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets SDK types
func apigateway_GetSdkTypes(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetSdkTypesInput{}

	if len(_apigatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _apigatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_apigatewayPosition) > 0 {
		input.Position = aws.String(_apigatewayPosition)
	}

	if resp, err := client.GetSdkTypes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a Stage resource.
func apigateway_GetStage(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetStageInput{
		// RestApiId: *string, // Required
		// StageName: *string, // Required
	}

	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayStageName) > 0 {
		input.StageName = aws.String(_apigatewayStageName)
	}

	if resp, err := client.GetStage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about one or more Stage resources.
func apigateway_GetStages(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetStagesInput{
		// RestApiId: *string, // Required
	}

	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayDeploymentId) > 0 {
		input.DeploymentId = aws.String(_apigatewayDeploymentId)
	}

	if resp, err := client.GetStages(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the Tags collection for a given resource.
func apigateway_GetTags(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetTagsInput{
		// ResourceArn: *string, // Required
	}

	if len(_apigatewayResourceArn) > 0 {
		input.ResourceArn = aws.String(_apigatewayResourceArn)
	}
	if len(_apigatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _apigatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_apigatewayPosition) > 0 {
		input.Position = aws.String(_apigatewayPosition)
	}

	if resp, err := client.GetTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the usage data of a usage plan in a specified time interval.
func apigateway_GetUsage(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetUsageInput{
		// EndDate: *string, // Required
		// StartDate: *string, // Required
		// UsagePlanId: *string, // Required
	}

	if len(_apigatewayEndDate) > 0 {
		input.EndDate = aws.String(_apigatewayEndDate)
	}
	if len(_apigatewayStartDate) > 0 {
		input.StartDate = aws.String(_apigatewayStartDate)
	}
	if len(_apigatewayUsagePlanId) > 0 {
		input.UsagePlanId = aws.String(_apigatewayUsagePlanId)
	}
	if len(_apigatewayKeyId) > 0 {
		input.KeyId = aws.String(_apigatewayKeyId)
	}
	if len(_apigatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _apigatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_apigatewayPosition) > 0 {
		input.Position = aws.String(_apigatewayPosition)
	}

	if disablePaginator() {
		if resp, err := client.GetUsage(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*apigateway.GetUsageOutput
	p := apigateway.NewGetUsagePaginator(client, input)
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

// Gets a usage plan of a given plan identifier.
func apigateway_GetUsagePlan(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetUsagePlanInput{
		// UsagePlanId: *string, // Required
	}

	if len(_apigatewayUsagePlanId) > 0 {
		input.UsagePlanId = aws.String(_apigatewayUsagePlanId)
	}

	if resp, err := client.GetUsagePlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a usage plan key of a given key identifier.
func apigateway_GetUsagePlanKey(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetUsagePlanKeyInput{
		// KeyId: *string, // Required
		// UsagePlanId: *string, // Required
	}

	if len(_apigatewayKeyId) > 0 {
		input.KeyId = aws.String(_apigatewayKeyId)
	}
	if len(_apigatewayUsagePlanId) > 0 {
		input.UsagePlanId = aws.String(_apigatewayUsagePlanId)
	}

	if resp, err := client.GetUsagePlanKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets all the usage plan keys representing the API keys added to a specified
// usage plan.
func apigateway_GetUsagePlanKeys(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetUsagePlanKeysInput{
		// UsagePlanId: *string, // Required
	}

	if len(_apigatewayUsagePlanId) > 0 {
		input.UsagePlanId = aws.String(_apigatewayUsagePlanId)
	}
	if len(_apigatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _apigatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_apigatewayNameQuery) > 0 {
		input.NameQuery = aws.String(_apigatewayNameQuery)
	}
	if len(_apigatewayPosition) > 0 {
		input.Position = aws.String(_apigatewayPosition)
	}

	if disablePaginator() {
		if resp, err := client.GetUsagePlanKeys(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*apigateway.GetUsagePlanKeysOutput
	p := apigateway.NewGetUsagePlanKeysPaginator(client, input)
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

// Gets all the usage plans of the caller's account.
func apigateway_GetUsagePlans(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetUsagePlansInput{}

	if len(_apigatewayKeyId) > 0 {
		input.KeyId = aws.String(_apigatewayKeyId)
	}
	if len(_apigatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _apigatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_apigatewayPosition) > 0 {
		input.Position = aws.String(_apigatewayPosition)
	}

	if disablePaginator() {
		if resp, err := client.GetUsagePlans(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*apigateway.GetUsagePlansOutput
	p := apigateway.NewGetUsagePlansPaginator(client, input)
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

// Gets a specified VPC link under the caller's account in a region.
func apigateway_GetVpcLink(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetVpcLinkInput{
		// VpcLinkId: *string, // Required
	}

	if len(_apigatewayVpcLinkId) > 0 {
		input.VpcLinkId = aws.String(_apigatewayVpcLinkId)
	}

	if resp, err := client.GetVpcLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the VpcLinks collection under the caller's account in a selected region.
func apigateway_GetVpcLinks(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.GetVpcLinksInput{}

	if len(_apigatewayLimit) > 0 {
		if err := assignInputField(input, "Limit", _apigatewayLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_apigatewayPosition) > 0 {
		input.Position = aws.String(_apigatewayPosition)
	}

	if disablePaginator() {
		if resp, err := client.GetVpcLinks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*apigateway.GetVpcLinksOutput
	p := apigateway.NewGetVpcLinksPaginator(client, input)
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

// Import API keys from an external source, such as a CSV-formatted file.
func apigateway_ImportApiKeys(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.ImportApiKeysInput{
		// Body: []byte, // Required
		// Format: types.ApiKeysFormat, // Required
	}

	if len(_apigatewayBody) > 0 {
		if err := assignInputField(input, "Body", _apigatewayBody); err != nil {
			log.Errorf("invalid --body: %s", err.Error())
			return
		}
	}
	if len(_apigatewayFormat) > 0 {
		if err := assignInputField(input, "Format", _apigatewayFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_apigatewayFailOnWarnings) > 0 {
		if err := assignInputField(input, "FailOnWarnings", _apigatewayFailOnWarnings); err != nil {
			log.Errorf("invalid --fail-on-warnings: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportApiKeys(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports documentation parts
func apigateway_ImportDocumentationParts(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.ImportDocumentationPartsInput{
		// Body: []byte, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayBody) > 0 {
		if err := assignInputField(input, "Body", _apigatewayBody); err != nil {
			log.Errorf("invalid --body: %s", err.Error())
			return
		}
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayFailOnWarnings) > 0 {
		if err := assignInputField(input, "FailOnWarnings", _apigatewayFailOnWarnings); err != nil {
			log.Errorf("invalid --fail-on-warnings: %s", err.Error())
			return
		}
	}
	if len(_apigatewayMode) > 0 {
		if err := assignInputField(input, "Mode", _apigatewayMode); err != nil {
			log.Errorf("invalid --mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportDocumentationParts(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A feature of the API Gateway control service for creating a new API from an
// external API definition file.
func apigateway_ImportRestApi(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.ImportRestApiInput{
		// Body: []byte, // Required
	}

	if len(_apigatewayBody) > 0 {
		if err := assignInputField(input, "Body", _apigatewayBody); err != nil {
			log.Errorf("invalid --body: %s", err.Error())
			return
		}
	}
	if len(_apigatewayFailOnWarnings) > 0 {
		if err := assignInputField(input, "FailOnWarnings", _apigatewayFailOnWarnings); err != nil {
			log.Errorf("invalid --fail-on-warnings: %s", err.Error())
			return
		}
	}
	if len(_apigatewayParameters) > 0 {
		if err := assignInputField(input, "Parameters", _apigatewayParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportRestApi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a customization of a GatewayResponse of a specified response type and
// status code on the given RestApi.
func apigateway_PutGatewayResponse(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.PutGatewayResponseInput{
		// ResponseType: types.GatewayResponseType, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayResponseType) > 0 {
		if err := assignInputField(input, "ResponseType", _apigatewayResponseType); err != nil {
			log.Errorf("invalid --response-type: %s", err.Error())
			return
		}
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayResponseParameters) > 0 {
		if err := assignInputField(input, "ResponseParameters", _apigatewayResponseParameters); err != nil {
			log.Errorf("invalid --response-parameters: %s", err.Error())
			return
		}
	}
	if len(_apigatewayResponseTemplates) > 0 {
		if err := assignInputField(input, "ResponseTemplates", _apigatewayResponseTemplates); err != nil {
			log.Errorf("invalid --response-templates: %s", err.Error())
			return
		}
	}
	if len(_apigatewayStatusCode) > 0 {
		input.StatusCode = aws.String(_apigatewayStatusCode)
	}

	if resp, err := client.PutGatewayResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets up a method's integration.
func apigateway_PutIntegration(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.PutIntegrationInput{
		// HttpMethod: *string, // Required
		// ResourceId: *string, // Required
		// RestApiId: *string, // Required
		// Type: types.IntegrationType, // Required
	}

	if len(_apigatewayHttpMethod) > 0 {
		input.HttpMethod = aws.String(_apigatewayHttpMethod)
	}
	if len(_apigatewayResourceId) > 0 {
		input.ResourceId = aws.String(_apigatewayResourceId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayType) > 0 {
		if err := assignInputField(input, "Type", _apigatewayType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_apigatewayCacheKeyParameters) > 0 {
		input.CacheKeyParameters = append([]string(nil), _apigatewayCacheKeyParameters...)
	}
	if len(_apigatewayCacheNamespace) > 0 {
		input.CacheNamespace = aws.String(_apigatewayCacheNamespace)
	}
	if len(_apigatewayConnectionId) > 0 {
		input.ConnectionId = aws.String(_apigatewayConnectionId)
	}
	if len(_apigatewayConnectionType) > 0 {
		if err := assignInputField(input, "ConnectionType", _apigatewayConnectionType); err != nil {
			log.Errorf("invalid --connection-type: %s", err.Error())
			return
		}
	}
	if len(_apigatewayContentHandling) > 0 {
		if err := assignInputField(input, "ContentHandling", _apigatewayContentHandling); err != nil {
			log.Errorf("invalid --content-handling: %s", err.Error())
			return
		}
	}
	if len(_apigatewayCredentials) > 0 {
		input.Credentials = aws.String(_apigatewayCredentials)
	}
	if len(_apigatewayIntegrationHttpMethod) > 0 {
		input.IntegrationHttpMethod = aws.String(_apigatewayIntegrationHttpMethod)
	}
	if len(_apigatewayIntegrationTarget) > 0 {
		input.IntegrationTarget = aws.String(_apigatewayIntegrationTarget)
	}
	if len(_apigatewayPassthroughBehavior) > 0 {
		input.PassthroughBehavior = aws.String(_apigatewayPassthroughBehavior)
	}
	if len(_apigatewayRequestParameters) > 0 {
		if err := assignInputField(input, "RequestParameters", _apigatewayRequestParameters); err != nil {
			log.Errorf("invalid --request-parameters: %s", err.Error())
			return
		}
	}
	if len(_apigatewayRequestTemplates) > 0 {
		if err := assignInputField(input, "RequestTemplates", _apigatewayRequestTemplates); err != nil {
			log.Errorf("invalid --request-templates: %s", err.Error())
			return
		}
	}
	if len(_apigatewayResponseTransferMode) > 0 {
		if err := assignInputField(input, "ResponseTransferMode", _apigatewayResponseTransferMode); err != nil {
			log.Errorf("invalid --response-transfer-mode: %s", err.Error())
			return
		}
	}
	if len(_apigatewayTimeoutInMillis) > 0 {
		if err := assignInputField(input, "TimeoutInMillis", _apigatewayTimeoutInMillis); err != nil {
			log.Errorf("invalid --timeout-in-millis: %s", err.Error())
			return
		}
	}
	if len(_apigatewayTlsConfig) > 0 {
		if err := assignInputField(input, "TlsConfig", _apigatewayTlsConfig); err != nil {
			log.Errorf("invalid --tls-config: %s", err.Error())
			return
		}
	}
	if len(_apigatewayUri) > 0 {
		input.Uri = aws.String(_apigatewayUri)
	}

	if resp, err := client.PutIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Represents a put integration.
func apigateway_PutIntegrationResponse(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.PutIntegrationResponseInput{
		// HttpMethod: *string, // Required
		// ResourceId: *string, // Required
		// RestApiId: *string, // Required
		// StatusCode: *string, // Required
	}

	if len(_apigatewayHttpMethod) > 0 {
		input.HttpMethod = aws.String(_apigatewayHttpMethod)
	}
	if len(_apigatewayResourceId) > 0 {
		input.ResourceId = aws.String(_apigatewayResourceId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayStatusCode) > 0 {
		input.StatusCode = aws.String(_apigatewayStatusCode)
	}
	if len(_apigatewayContentHandling) > 0 {
		if err := assignInputField(input, "ContentHandling", _apigatewayContentHandling); err != nil {
			log.Errorf("invalid --content-handling: %s", err.Error())
			return
		}
	}
	if len(_apigatewayResponseParameters) > 0 {
		if err := assignInputField(input, "ResponseParameters", _apigatewayResponseParameters); err != nil {
			log.Errorf("invalid --response-parameters: %s", err.Error())
			return
		}
	}
	if len(_apigatewayResponseTemplates) > 0 {
		if err := assignInputField(input, "ResponseTemplates", _apigatewayResponseTemplates); err != nil {
			log.Errorf("invalid --response-templates: %s", err.Error())
			return
		}
	}
	if len(_apigatewaySelectionPattern) > 0 {
		input.SelectionPattern = aws.String(_apigatewaySelectionPattern)
	}

	if resp, err := client.PutIntegrationResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Add a method to an existing Resource resource.
func apigateway_PutMethod(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.PutMethodInput{
		// AuthorizationType: *string, // Required
		// HttpMethod: *string, // Required
		// ResourceId: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayAuthorizationType) > 0 {
		input.AuthorizationType = aws.String(_apigatewayAuthorizationType)
	}
	if len(_apigatewayHttpMethod) > 0 {
		input.HttpMethod = aws.String(_apigatewayHttpMethod)
	}
	if len(_apigatewayResourceId) > 0 {
		input.ResourceId = aws.String(_apigatewayResourceId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayApiKeyRequired) > 0 {
		if err := assignInputField(input, "ApiKeyRequired", _apigatewayApiKeyRequired); err != nil {
			log.Errorf("invalid --api-key-required: %s", err.Error())
			return
		}
	}
	if len(_apigatewayAuthorizationScopes) > 0 {
		input.AuthorizationScopes = append([]string(nil), _apigatewayAuthorizationScopes...)
	}
	if len(_apigatewayAuthorizerId) > 0 {
		input.AuthorizerId = aws.String(_apigatewayAuthorizerId)
	}
	if len(_apigatewayOperationName) > 0 {
		input.OperationName = aws.String(_apigatewayOperationName)
	}
	if len(_apigatewayRequestModels) > 0 {
		if err := assignInputField(input, "RequestModels", _apigatewayRequestModels); err != nil {
			log.Errorf("invalid --request-models: %s", err.Error())
			return
		}
	}
	if len(_apigatewayRequestParameters) > 0 {
		if err := assignInputField(input, "RequestParameters", _apigatewayRequestParameters); err != nil {
			log.Errorf("invalid --request-parameters: %s", err.Error())
			return
		}
	}
	if len(_apigatewayRequestValidatorId) > 0 {
		input.RequestValidatorId = aws.String(_apigatewayRequestValidatorId)
	}

	if resp, err := client.PutMethod(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a MethodResponse to an existing Method resource.
func apigateway_PutMethodResponse(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.PutMethodResponseInput{
		// HttpMethod: *string, // Required
		// ResourceId: *string, // Required
		// RestApiId: *string, // Required
		// StatusCode: *string, // Required
	}

	if len(_apigatewayHttpMethod) > 0 {
		input.HttpMethod = aws.String(_apigatewayHttpMethod)
	}
	if len(_apigatewayResourceId) > 0 {
		input.ResourceId = aws.String(_apigatewayResourceId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayStatusCode) > 0 {
		input.StatusCode = aws.String(_apigatewayStatusCode)
	}
	if len(_apigatewayResponseModels) > 0 {
		if err := assignInputField(input, "ResponseModels", _apigatewayResponseModels); err != nil {
			log.Errorf("invalid --response-models: %s", err.Error())
			return
		}
	}
	if len(_apigatewayResponseParameters) > 0 {
		if err := assignInputField(input, "ResponseParameters", _apigatewayResponseParameters); err != nil {
			log.Errorf("invalid --response-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutMethodResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A feature of the API Gateway control service for updating an existing API with
// an input of external API definitions. The update can take the form of merging
// the supplied definition into the existing API or overwriting the existing API.
func apigateway_PutRestApi(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.PutRestApiInput{
		// Body: []byte, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayBody) > 0 {
		if err := assignInputField(input, "Body", _apigatewayBody); err != nil {
			log.Errorf("invalid --body: %s", err.Error())
			return
		}
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayFailOnWarnings) > 0 {
		if err := assignInputField(input, "FailOnWarnings", _apigatewayFailOnWarnings); err != nil {
			log.Errorf("invalid --fail-on-warnings: %s", err.Error())
			return
		}
	}
	if len(_apigatewayMode) > 0 {
		if err := assignInputField(input, "Mode", _apigatewayMode); err != nil {
			log.Errorf("invalid --mode: %s", err.Error())
			return
		}
	}
	if len(_apigatewayParameters) > 0 {
		if err := assignInputField(input, "Parameters", _apigatewayParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutRestApi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Rejects a domain name access association with a private custom domain name.
// To reject a domain name access association with an access association source in
// another AWS account, use this operation. To remove a domain name access
// association with an access association source in your own account, use the
// DeleteDomainNameAccessAssociation operation.
func apigateway_RejectDomainNameAccessAssociation(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.RejectDomainNameAccessAssociationInput{
		// DomainNameAccessAssociationArn: *string, // Required
		// DomainNameArn: *string, // Required
	}

	if len(_apigatewayDomainNameAccessAssociationArn) > 0 {
		input.DomainNameAccessAssociationArn = aws.String(_apigatewayDomainNameAccessAssociationArn)
	}
	if len(_apigatewayDomainNameArn) > 0 {
		input.DomainNameArn = aws.String(_apigatewayDomainNameArn)
	}

	if resp, err := client.RejectDomainNameAccessAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates a tag on a given resource.
func apigateway_TagResource(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_apigatewayResourceArn) > 0 {
		input.ResourceArn = aws.String(_apigatewayResourceArn)
	}
	if len(_apigatewayTags) > 0 {
		if err := assignInputField(input, "Tags", _apigatewayTags); err != nil {
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

// Simulate the execution of an Authorizer in your RestApi with headers,
// parameters, and an incoming request body.
func apigateway_TestInvokeAuthorizer(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.TestInvokeAuthorizerInput{
		// AuthorizerId: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayAuthorizerId) > 0 {
		input.AuthorizerId = aws.String(_apigatewayAuthorizerId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayAdditionalContext) > 0 {
		if err := assignInputField(input, "AdditionalContext", _apigatewayAdditionalContext); err != nil {
			log.Errorf("invalid --additional-context: %s", err.Error())
			return
		}
	}
	if len(_apigatewayBody) > 0 {
		input.Body = aws.String(_apigatewayBody)
	}
	if len(_apigatewayHeaders) > 0 {
		if err := assignInputField(input, "Headers", _apigatewayHeaders); err != nil {
			log.Errorf("invalid --headers: %s", err.Error())
			return
		}
	}
	if len(_apigatewayMultiValueHeaders) > 0 {
		if err := assignInputField(input, "MultiValueHeaders", _apigatewayMultiValueHeaders); err != nil {
			log.Errorf("invalid --multi-value-headers: %s", err.Error())
			return
		}
	}
	if len(_apigatewayPathWithQueryString) > 0 {
		input.PathWithQueryString = aws.String(_apigatewayPathWithQueryString)
	}
	if len(_apigatewayStageVariables) > 0 {
		if err := assignInputField(input, "StageVariables", _apigatewayStageVariables); err != nil {
			log.Errorf("invalid --stage-variables: %s", err.Error())
			return
		}
	}

	if resp, err := client.TestInvokeAuthorizer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Simulate the invocation of a Method in your RestApi with headers, parameters,
// and an incoming request body.
func apigateway_TestInvokeMethod(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.TestInvokeMethodInput{
		// HttpMethod: *string, // Required
		// ResourceId: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayHttpMethod) > 0 {
		input.HttpMethod = aws.String(_apigatewayHttpMethod)
	}
	if len(_apigatewayResourceId) > 0 {
		input.ResourceId = aws.String(_apigatewayResourceId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayBody) > 0 {
		input.Body = aws.String(_apigatewayBody)
	}
	if len(_apigatewayClientCertificateId) > 0 {
		input.ClientCertificateId = aws.String(_apigatewayClientCertificateId)
	}
	if len(_apigatewayHeaders) > 0 {
		if err := assignInputField(input, "Headers", _apigatewayHeaders); err != nil {
			log.Errorf("invalid --headers: %s", err.Error())
			return
		}
	}
	if len(_apigatewayMultiValueHeaders) > 0 {
		if err := assignInputField(input, "MultiValueHeaders", _apigatewayMultiValueHeaders); err != nil {
			log.Errorf("invalid --multi-value-headers: %s", err.Error())
			return
		}
	}
	if len(_apigatewayPathWithQueryString) > 0 {
		input.PathWithQueryString = aws.String(_apigatewayPathWithQueryString)
	}
	if len(_apigatewayStageVariables) > 0 {
		if err := assignInputField(input, "StageVariables", _apigatewayStageVariables); err != nil {
			log.Errorf("invalid --stage-variables: %s", err.Error())
			return
		}
	}

	if resp, err := client.TestInvokeMethod(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a tag from a given resource.
func apigateway_UntagResource(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_apigatewayResourceArn) > 0 {
		input.ResourceArn = aws.String(_apigatewayResourceArn)
	}
	if len(_apigatewayTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _apigatewayTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes information about the current Account resource.
func apigateway_UpdateAccount(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.UpdateAccountInput{}

	if len(_apigatewayPatchOperations) > 0 {
		if err := assignInputField(input, "PatchOperations", _apigatewayPatchOperations); err != nil {
			log.Errorf("invalid --patch-operations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes information about an ApiKey resource.
func apigateway_UpdateApiKey(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.UpdateApiKeyInput{
		// ApiKey: *string, // Required
	}

	if len(_apigatewayApiKey) > 0 {
		input.ApiKey = aws.String(_apigatewayApiKey)
	}
	if len(_apigatewayPatchOperations) > 0 {
		if err := assignInputField(input, "PatchOperations", _apigatewayPatchOperations); err != nil {
			log.Errorf("invalid --patch-operations: %s", err.Error())
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

// Updates an existing Authorizer resource.
func apigateway_UpdateAuthorizer(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.UpdateAuthorizerInput{
		// AuthorizerId: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayAuthorizerId) > 0 {
		input.AuthorizerId = aws.String(_apigatewayAuthorizerId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayPatchOperations) > 0 {
		if err := assignInputField(input, "PatchOperations", _apigatewayPatchOperations); err != nil {
			log.Errorf("invalid --patch-operations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAuthorizer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes information about the BasePathMapping resource.
func apigateway_UpdateBasePathMapping(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.UpdateBasePathMappingInput{
		// BasePath: *string, // Required
		// DomainName: *string, // Required
	}

	if len(_apigatewayBasePath) > 0 {
		input.BasePath = aws.String(_apigatewayBasePath)
	}
	if len(_apigatewayDomainName) > 0 {
		input.DomainName = aws.String(_apigatewayDomainName)
	}
	if len(_apigatewayDomainNameId) > 0 {
		input.DomainNameId = aws.String(_apigatewayDomainNameId)
	}
	if len(_apigatewayPatchOperations) > 0 {
		if err := assignInputField(input, "PatchOperations", _apigatewayPatchOperations); err != nil {
			log.Errorf("invalid --patch-operations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateBasePathMapping(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes information about an ClientCertificate resource.
func apigateway_UpdateClientCertificate(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.UpdateClientCertificateInput{
		// ClientCertificateId: *string, // Required
	}

	if len(_apigatewayClientCertificateId) > 0 {
		input.ClientCertificateId = aws.String(_apigatewayClientCertificateId)
	}
	if len(_apigatewayPatchOperations) > 0 {
		if err := assignInputField(input, "PatchOperations", _apigatewayPatchOperations); err != nil {
			log.Errorf("invalid --patch-operations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateClientCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes information about a Deployment resource.
func apigateway_UpdateDeployment(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.UpdateDeploymentInput{
		// DeploymentId: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayDeploymentId) > 0 {
		input.DeploymentId = aws.String(_apigatewayDeploymentId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayPatchOperations) > 0 {
		if err := assignInputField(input, "PatchOperations", _apigatewayPatchOperations); err != nil {
			log.Errorf("invalid --patch-operations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a documentation part.
func apigateway_UpdateDocumentationPart(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.UpdateDocumentationPartInput{
		// DocumentationPartId: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayDocumentationPartId) > 0 {
		input.DocumentationPartId = aws.String(_apigatewayDocumentationPartId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayPatchOperations) > 0 {
		if err := assignInputField(input, "PatchOperations", _apigatewayPatchOperations); err != nil {
			log.Errorf("invalid --patch-operations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDocumentationPart(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a documentation version.
func apigateway_UpdateDocumentationVersion(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.UpdateDocumentationVersionInput{
		// DocumentationVersion: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayDocumentationVersion) > 0 {
		input.DocumentationVersion = aws.String(_apigatewayDocumentationVersion)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayPatchOperations) > 0 {
		if err := assignInputField(input, "PatchOperations", _apigatewayPatchOperations); err != nil {
			log.Errorf("invalid --patch-operations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDocumentationVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes information about the DomainName resource.
func apigateway_UpdateDomainName(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.UpdateDomainNameInput{
		// DomainName: *string, // Required
	}

	if len(_apigatewayDomainName) > 0 {
		input.DomainName = aws.String(_apigatewayDomainName)
	}
	if len(_apigatewayDomainNameId) > 0 {
		input.DomainNameId = aws.String(_apigatewayDomainNameId)
	}
	if len(_apigatewayPatchOperations) > 0 {
		if err := assignInputField(input, "PatchOperations", _apigatewayPatchOperations); err != nil {
			log.Errorf("invalid --patch-operations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDomainName(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a GatewayResponse of a specified response type on the given RestApi.
func apigateway_UpdateGatewayResponse(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.UpdateGatewayResponseInput{
		// ResponseType: types.GatewayResponseType, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayResponseType) > 0 {
		if err := assignInputField(input, "ResponseType", _apigatewayResponseType); err != nil {
			log.Errorf("invalid --response-type: %s", err.Error())
			return
		}
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayPatchOperations) > 0 {
		if err := assignInputField(input, "PatchOperations", _apigatewayPatchOperations); err != nil {
			log.Errorf("invalid --patch-operations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateGatewayResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Represents an update integration.
func apigateway_UpdateIntegration(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.UpdateIntegrationInput{
		// HttpMethod: *string, // Required
		// ResourceId: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayHttpMethod) > 0 {
		input.HttpMethod = aws.String(_apigatewayHttpMethod)
	}
	if len(_apigatewayResourceId) > 0 {
		input.ResourceId = aws.String(_apigatewayResourceId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayPatchOperations) > 0 {
		if err := assignInputField(input, "PatchOperations", _apigatewayPatchOperations); err != nil {
			log.Errorf("invalid --patch-operations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Represents an update integration response.
func apigateway_UpdateIntegrationResponse(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.UpdateIntegrationResponseInput{
		// HttpMethod: *string, // Required
		// ResourceId: *string, // Required
		// RestApiId: *string, // Required
		// StatusCode: *string, // Required
	}

	if len(_apigatewayHttpMethod) > 0 {
		input.HttpMethod = aws.String(_apigatewayHttpMethod)
	}
	if len(_apigatewayResourceId) > 0 {
		input.ResourceId = aws.String(_apigatewayResourceId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayStatusCode) > 0 {
		input.StatusCode = aws.String(_apigatewayStatusCode)
	}
	if len(_apigatewayPatchOperations) > 0 {
		if err := assignInputField(input, "PatchOperations", _apigatewayPatchOperations); err != nil {
			log.Errorf("invalid --patch-operations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateIntegrationResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing Method resource.
func apigateway_UpdateMethod(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.UpdateMethodInput{
		// HttpMethod: *string, // Required
		// ResourceId: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayHttpMethod) > 0 {
		input.HttpMethod = aws.String(_apigatewayHttpMethod)
	}
	if len(_apigatewayResourceId) > 0 {
		input.ResourceId = aws.String(_apigatewayResourceId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayPatchOperations) > 0 {
		if err := assignInputField(input, "PatchOperations", _apigatewayPatchOperations); err != nil {
			log.Errorf("invalid --patch-operations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMethod(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing MethodResponse resource.
func apigateway_UpdateMethodResponse(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.UpdateMethodResponseInput{
		// HttpMethod: *string, // Required
		// ResourceId: *string, // Required
		// RestApiId: *string, // Required
		// StatusCode: *string, // Required
	}

	if len(_apigatewayHttpMethod) > 0 {
		input.HttpMethod = aws.String(_apigatewayHttpMethod)
	}
	if len(_apigatewayResourceId) > 0 {
		input.ResourceId = aws.String(_apigatewayResourceId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayStatusCode) > 0 {
		input.StatusCode = aws.String(_apigatewayStatusCode)
	}
	if len(_apigatewayPatchOperations) > 0 {
		if err := assignInputField(input, "PatchOperations", _apigatewayPatchOperations); err != nil {
			log.Errorf("invalid --patch-operations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMethodResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes information about a model. The maximum size of the model is 400 KB.
func apigateway_UpdateModel(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.UpdateModelInput{
		// ModelName: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayModelName) > 0 {
		input.ModelName = aws.String(_apigatewayModelName)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayPatchOperations) > 0 {
		if err := assignInputField(input, "PatchOperations", _apigatewayPatchOperations); err != nil {
			log.Errorf("invalid --patch-operations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a RequestValidator of a given RestApi.
func apigateway_UpdateRequestValidator(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.UpdateRequestValidatorInput{
		// RequestValidatorId: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayRequestValidatorId) > 0 {
		input.RequestValidatorId = aws.String(_apigatewayRequestValidatorId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayPatchOperations) > 0 {
		if err := assignInputField(input, "PatchOperations", _apigatewayPatchOperations); err != nil {
			log.Errorf("invalid --patch-operations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRequestValidator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes information about a Resource resource.
func apigateway_UpdateResource(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.UpdateResourceInput{
		// ResourceId: *string, // Required
		// RestApiId: *string, // Required
	}

	if len(_apigatewayResourceId) > 0 {
		input.ResourceId = aws.String(_apigatewayResourceId)
	}
	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayPatchOperations) > 0 {
		if err := assignInputField(input, "PatchOperations", _apigatewayPatchOperations); err != nil {
			log.Errorf("invalid --patch-operations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes information about the specified API.
func apigateway_UpdateRestApi(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.UpdateRestApiInput{
		// RestApiId: *string, // Required
	}

	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayPatchOperations) > 0 {
		if err := assignInputField(input, "PatchOperations", _apigatewayPatchOperations); err != nil {
			log.Errorf("invalid --patch-operations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRestApi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes information about a Stage resource.
func apigateway_UpdateStage(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.UpdateStageInput{
		// RestApiId: *string, // Required
		// StageName: *string, // Required
	}

	if len(_apigatewayRestApiId) > 0 {
		input.RestApiId = aws.String(_apigatewayRestApiId)
	}
	if len(_apigatewayStageName) > 0 {
		input.StageName = aws.String(_apigatewayStageName)
	}
	if len(_apigatewayPatchOperations) > 0 {
		if err := assignInputField(input, "PatchOperations", _apigatewayPatchOperations); err != nil {
			log.Errorf("invalid --patch-operations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateStage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Grants a temporary extension to the remaining quota of a usage plan associated
// with a specified API key.
func apigateway_UpdateUsage(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.UpdateUsageInput{
		// KeyId: *string, // Required
		// UsagePlanId: *string, // Required
	}

	if len(_apigatewayKeyId) > 0 {
		input.KeyId = aws.String(_apigatewayKeyId)
	}
	if len(_apigatewayUsagePlanId) > 0 {
		input.UsagePlanId = aws.String(_apigatewayUsagePlanId)
	}
	if len(_apigatewayPatchOperations) > 0 {
		if err := assignInputField(input, "PatchOperations", _apigatewayPatchOperations); err != nil {
			log.Errorf("invalid --patch-operations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateUsage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a usage plan of a given plan Id.
func apigateway_UpdateUsagePlan(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.UpdateUsagePlanInput{
		// UsagePlanId: *string, // Required
	}

	if len(_apigatewayUsagePlanId) > 0 {
		input.UsagePlanId = aws.String(_apigatewayUsagePlanId)
	}
	if len(_apigatewayPatchOperations) > 0 {
		if err := assignInputField(input, "PatchOperations", _apigatewayPatchOperations); err != nil {
			log.Errorf("invalid --patch-operations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateUsagePlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing VpcLink of a specified identifier.
func apigateway_UpdateVpcLink(cfg aws.Config, client *apigateway.Client) {
	input := &apigateway.UpdateVpcLinkInput{
		// VpcLinkId: *string, // Required
	}

	if len(_apigatewayVpcLinkId) > 0 {
		input.VpcLinkId = aws.String(_apigatewayVpcLinkId)
	}
	if len(_apigatewayPatchOperations) > 0 {
		if err := assignInputField(input, "PatchOperations", _apigatewayPatchOperations); err != nil {
			log.Errorf("invalid --patch-operations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateVpcLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_apigatewayCmd)
	_apigatewayCmd.Flags().SortFlags = false

	_apigatewayCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_apigatewayCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_apigatewayCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_apigatewayCmd.Flags().StringVarP(&_apigatewayAccepts, "accepts", "", "", "Accepts")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayAccessAssociationSource, "access-association-source", "", "", "Access Association Source")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayAccessAssociationSourceType, "access-association-source-type", "", "", "Access Association Source Type")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayAdditionalContext, "additional-context", "", "", "Additional Context")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayApiKey, "api-key", "", "", "API Key")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayApiKeyRequired, "api-key-required", "", "", "API Key Required")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayApiKeySource, "api-key-source", "", "", "API Key Source")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayApiStages, "api-stages", "", "", "API Stages")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayAuthType, "auth-type", "", "", "Auth Type")
	_apigatewayCmd.Flags().StringSliceVarP(&_apigatewayAuthorizationScopes, "authorization-scopes", "", nil, "Authorization Scopes")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayAuthorizationType, "authorization-type", "", "", "Authorization Type")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayAuthorizerCredentials, "authorizer-credentials", "", "", "Authorizer Credentials")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayAuthorizerId, "authorizer-id", "", "", "Authorizer ID")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayAuthorizerResultTtlInSeconds, "authorizer-result-ttl-in-seconds", "", "", "Authorizer Result TTL In Seconds")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayAuthorizerUri, "authorizer-uri", "", "", "Authorizer URI")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayBasePath, "base-path", "", "", "Base Path")
	_apigatewayCmd.Flags().StringSliceVarP(&_apigatewayBinaryMediaTypes, "binary-media-types", "", nil, "Binary Media Types")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayBody, "body", "", "", "Body")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayCacheClusterEnabled, "cache-cluster-enabled", "", "", "Cache Cluster Enabled")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayCacheClusterSize, "cache-cluster-size", "", "", "Cache Cluster Size")
	_apigatewayCmd.Flags().StringSliceVarP(&_apigatewayCacheKeyParameters, "cache-key-parameters", "", nil, "Cache Key Parameters")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayCacheNamespace, "cache-namespace", "", "", "Cache Namespace")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayCanarySettings, "canary-settings", "", "", "Canary Settings")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayCertificateArn, "certificate-arn", "", "", "Certificate ARN")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayCertificateBody, "certificate-body", "", "", "Certificate Body")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayCertificateChain, "certificate-chain", "", "", "Certificate Chain")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayCertificateName, "certificate-name", "", "", "Certificate Name")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayCertificatePrivateKey, "certificate-private-key", "", "", "Certificate Private Key")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayClientCertificateId, "client-certificate-id", "", "", "Client Certificate ID")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayCloneFrom, "clone-from", "", "", "Clone From")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayConnectionId, "connection-id", "", "", "Connection ID")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayConnectionType, "connection-type", "", "", "Connection Type")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayContentHandling, "content-handling", "", "", "Content Handling")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayContentType, "content-type", "", "", "Content Type")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayCredentials, "credentials", "", "", "Credentials")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayCustomerId, "customer-id", "", "", "Customer ID")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayDeploymentId, "deployment-id", "", "", "Deployment ID")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayDescription, "description", "", "", "Description")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayDisableExecuteApiEndpoint, "disable-execute-api-endpoint", "", "", "Disable Execute API Endpoint")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayDocumentationPartId, "documentation-part-id", "", "", "Documentation Part ID")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayDocumentationVersion, "documentation-version", "", "", "Documentation Version")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayDomainName, "domain-name", "", "", "Domain Name")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayDomainNameAccessAssociationArn, "domain-name-access-association-arn", "", "", "Domain Name Access Association ARN")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayDomainNameArn, "domain-name-arn", "", "", "Domain Name ARN")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayDomainNameId, "domain-name-id", "", "", "Domain Name ID")
	_apigatewayCmd.Flags().StringSliceVarP(&_apigatewayEmbed, "embed", "", nil, "Embed")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayEnabled, "enabled", "", "", "Enabled")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayEndDate, "end-date", "", "", "End Date")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayEndpointAccessMode, "endpoint-access-mode", "", "", "Endpoint Access Mode")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayEndpointConfiguration, "endpoint-configuration", "", "", "Endpoint Configuration")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayExportType, "export-type", "", "", "Export Type")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayFailOnWarnings, "fail-on-warnings", "", "", "Fail On Warnings")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayFlatten, "flatten", "", "", "Flatten")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayFormat, "format", "", "", "Format")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayGenerateDistinctId, "generate-distinct-id", "", "", "Generate Distinct ID")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayHeaders, "headers", "", "", "Headers")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayHttpMethod, "http-method", "", "", "HTTP Method")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayId, "id", "", "", "ID")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayIdentitySource, "identity-source", "", "", "Identity Source")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayIdentityValidationExpression, "identity-validation-expression", "", "", "Identity Validation Expression")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayIncludeValue, "include-value", "", "", "Include Value")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayIncludeValues, "include-values", "", "", "Include Values")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayIntegrationHttpMethod, "integration-http-method", "", "", "Integration HTTP Method")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayIntegrationTarget, "integration-target", "", "", "Integration Target")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayKeyId, "key-id", "", "", "Key ID")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayKeyType, "key-type", "", "", "Key Type")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayLimit, "limit", "", "", "Limit")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayLocation, "location", "", "", "Location")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayLocationStatus, "location-status", "", "", "Location Status")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayMinimumCompressionSize, "minimum-compression-size", "", "", "Minimum Compression Size")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayMode, "mode", "", "", "Mode")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayModelName, "model-name", "", "", "Model Name")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayMultiValueHeaders, "multi-value-headers", "", "", "Multi Value Headers")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayMutualTlsAuthentication, "mutual-tls-authentication", "", "", "Mutual TLS Authentication")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayName, "name", "", "", "Name")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayNameQuery, "name-query", "", "", "Name Query")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayOperationName, "operation-name", "", "", "Operation Name")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayOwnershipVerificationCertificateArn, "ownership-verification-certificate-arn", "", "", "Ownership Verification Certificate ARN")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayParameters, "parameters", "", "", "Parameters")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayParentId, "parent-id", "", "", "Parent ID")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayPassthroughBehavior, "passthrough-behavior", "", "", "Passthrough Behavior")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayPatchOperations, "patch-operations", "", "", "Patch Operations")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayPath, "path", "", "", "Path")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayPathPart, "path-part", "", "", "Path Part")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayPathWithQueryString, "path-with-query-string", "", "", "Path With Query String")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayPolicy, "policy", "", "", "Policy")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayPosition, "position", "", "", "Position")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayProperties, "properties", "", "", "Properties")
	_apigatewayCmd.Flags().StringSliceVarP(&_apigatewayProviderARNs, "provider-arns", "", nil, "Provider Arns")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayQuota, "quota", "", "", "Quota")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayRegionalCertificateArn, "regional-certificate-arn", "", "", "Regional Certificate ARN")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayRegionalCertificateName, "regional-certificate-name", "", "", "Regional Certificate Name")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayRequestModels, "request-models", "", "", "Request Models")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayRequestParameters, "request-parameters", "", "", "Request Parameters")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayRequestTemplates, "request-templates", "", "", "Request Templates")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayRequestValidatorId, "request-validator-id", "", "", "Request Validator ID")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayResourceArn, "resource-arn", "", "", "Resource ARN")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayResourceId, "resource-id", "", "", "Resource ID")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayResourceOwner, "resource-owner", "", "", "Resource Owner")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayResponseModels, "response-models", "", "", "Response Models")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayResponseParameters, "response-parameters", "", "", "Response Parameters")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayResponseTemplates, "response-templates", "", "", "Response Templates")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayResponseTransferMode, "response-transfer-mode", "", "", "Response Transfer Mode")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayResponseType, "response-type", "", "", "Response Type")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayRestApiId, "rest-api-id", "", "", "Rest API ID")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayRoutingMode, "routing-mode", "", "", "Routing Mode")
	_apigatewayCmd.Flags().StringVarP(&_apigatewaySchema, "schema", "", "", "Schema")
	_apigatewayCmd.Flags().StringVarP(&_apigatewaySdkType, "sdk-type", "", "", "Sdk Type")
	_apigatewayCmd.Flags().StringVarP(&_apigatewaySecurityPolicy, "security-policy", "", "", "Security Policy")
	_apigatewayCmd.Flags().StringVarP(&_apigatewaySelectionPattern, "selection-pattern", "", "", "Selection Pattern")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayStage, "stage", "", "", "Stage")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayStageDescription, "stage-description", "", "", "Stage Description")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayStageKeys, "stage-keys", "", "", "Stage Keys")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayStageName, "stage-name", "", "", "Stage Name")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayStageVariables, "stage-variables", "", "", "Stage Variables")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayStartDate, "start-date", "", "", "Start Date")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayStatusCode, "status-code", "", "", "Status Code")
	_apigatewayCmd.Flags().StringSliceVarP(&_apigatewayTagKeys, "tag-keys", "", nil, "Tag Keys")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayTags, "tags", "", "", "Tags")
	_apigatewayCmd.Flags().StringSliceVarP(&_apigatewayTargetArns, "target-arns", "", nil, "Target Arns")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayThrottle, "throttle", "", "", "Throttle")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayTimeoutInMillis, "timeout-in-millis", "", "", "Timeout In Millis")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayTlsConfig, "tls-config", "", "", "TLS Config")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayTracingEnabled, "tracing-enabled", "", "", "Tracing Enabled")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayType, "type", "", "", "Type")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayUri, "uri", "", "", "URI")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayUsagePlanId, "usage-plan-id", "", "", "Usage Plan ID")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayValidateRequestBody, "validate-request-body", "", "", "Validate Request Body")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayValidateRequestParameters, "validate-request-parameters", "", "", "Validate Request Parameters")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayValue, "value", "", "", "Value")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayVariables, "variables", "", "", "Variables")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayVersion, "version", "", "", "Version")
	_apigatewayCmd.Flags().StringVarP(&_apigatewayVpcLinkId, "vpc-link-id", "", "", "VPC Link ID")

	_apigatewayCmd.Flags().BoolVarP(&_apigatewayCreateApiKey, "create-api-key", "", false, "Create API Key")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayCreateAuthorizer, "create-authorizer", "", false, "Create Authorizer")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayCreateBasePathMapping, "create-base-path-mapping", "", false, "Create Base Path Mapping")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayCreateDeployment, "create-deployment", "", false, "Create Deployment")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayCreateDocumentationPart, "create-documentation-part", "", false, "Create Documentation Part")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayCreateDocumentationVersion, "create-documentation-version", "", false, "Create Documentation Version")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayCreateDomainName, "create-domain-name", "", false, "Create Domain Name")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayCreateDomainNameAccessAssociation, "create-domain-name-access-association", "", false, "Create Domain Name Access Association")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayCreateModel, "create-model", "", false, "Create Model")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayCreateRequestValidator, "create-request-validator", "", false, "Create Request Validator")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayCreateResource, "create-resource", "", false, "Create Resource")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayCreateRestApi, "create-rest-api", "", false, "Create Rest API")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayCreateStage, "create-stage", "", false, "Create Stage")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayCreateUsagePlan, "create-usage-plan", "", false, "Create Usage Plan")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayCreateUsagePlanKey, "create-usage-plan-key", "", false, "Create Usage Plan Key")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayCreateVpcLink, "create-vpc-link", "", false, "Create VPC Link")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayDeleteApiKey, "delete-api-key", "", false, "Delete API Key")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayDeleteAuthorizer, "delete-authorizer", "", false, "Delete Authorizer")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayDeleteBasePathMapping, "delete-base-path-mapping", "", false, "Delete Base Path Mapping")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayDeleteClientCertificate, "delete-client-certificate", "", false, "Delete Client Certificate")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayDeleteDeployment, "delete-deployment", "", false, "Delete Deployment")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayDeleteDocumentationPart, "delete-documentation-part", "", false, "Delete Documentation Part")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayDeleteDocumentationVersion, "delete-documentation-version", "", false, "Delete Documentation Version")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayDeleteDomainName, "delete-domain-name", "", false, "Delete Domain Name")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayDeleteDomainNameAccessAssociation, "delete-domain-name-access-association", "", false, "Delete Domain Name Access Association")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayDeleteGatewayResponse, "delete-gateway-response", "", false, "Delete Gateway Response")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayDeleteIntegration, "delete-integration", "", false, "Delete Integration")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayDeleteIntegrationResponse, "delete-integration-response", "", false, "Delete Integration Response")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayDeleteMethod, "delete-method", "", false, "Delete Method")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayDeleteMethodResponse, "delete-method-response", "", false, "Delete Method Response")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayDeleteModel, "delete-model", "", false, "Delete Model")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayDeleteRequestValidator, "delete-request-validator", "", false, "Delete Request Validator")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayDeleteResource, "delete-resource", "", false, "Delete Resource")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayDeleteRestApi, "delete-rest-api", "", false, "Delete Rest API")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayDeleteStage, "delete-stage", "", false, "Delete Stage")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayDeleteUsagePlan, "delete-usage-plan", "", false, "Delete Usage Plan")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayDeleteUsagePlanKey, "delete-usage-plan-key", "", false, "Delete Usage Plan Key")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayDeleteVpcLink, "delete-vpc-link", "", false, "Delete VPC Link")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayFlushStageAuthorizersCache, "flush-stage-authorizers-cache", "", false, "Flush Stage Authorizers Cache")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayFlushStageCache, "flush-stage-cache", "", false, "Flush Stage Cache")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGenerateClientCertificate, "generate-client-certificate", "", false, "Generate Client Certificate")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetAccount, "get-account", "", false, "Get Account")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetApiKey, "get-api-key", "", false, "Get API Key")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetApiKeys, "get-api-keys", "", false, "Get API Keys")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetAuthorizer, "get-authorizer", "", false, "Get Authorizer")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetAuthorizers, "get-authorizers", "", false, "Get Authorizers")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetBasePathMapping, "get-base-path-mapping", "", false, "Get Base Path Mapping")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetBasePathMappings, "get-base-path-mappings", "", false, "Get Base Path Mappings")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetClientCertificate, "get-client-certificate", "", false, "Get Client Certificate")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetClientCertificates, "get-client-certificates", "", false, "Get Client Certificates")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetDeployment, "get-deployment", "", false, "Get Deployment")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetDeployments, "get-deployments", "", false, "Get Deployments")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetDocumentationPart, "get-documentation-part", "", false, "Get Documentation Part")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetDocumentationParts, "get-documentation-parts", "", false, "Get Documentation Parts")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetDocumentationVersion, "get-documentation-version", "", false, "Get Documentation Version")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetDocumentationVersions, "get-documentation-versions", "", false, "Get Documentation Versions")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetDomainName, "get-domain-name", "", false, "Get Domain Name")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetDomainNameAccessAssociations, "get-domain-name-access-associations", "", false, "Get Domain Name Access Associations")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetDomainNames, "get-domain-names", "", false, "Get Domain Names")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetExport, "get-export", "", false, "Get Export")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetGatewayResponse, "get-gateway-response", "", false, "Get Gateway Response")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetGatewayResponses, "get-gateway-responses", "", false, "Get Gateway Responses")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetIntegration, "get-integration", "", false, "Get Integration")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetIntegrationResponse, "get-integration-response", "", false, "Get Integration Response")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetMethod, "get-method", "", false, "Get Method")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetMethodResponse, "get-method-response", "", false, "Get Method Response")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetModel, "get-model", "", false, "Get Model")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetModelTemplate, "get-model-template", "", false, "Get Model Template")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetModels, "get-models", "", false, "Get Models")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetRequestValidator, "get-request-validator", "", false, "Get Request Validator")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetRequestValidators, "get-request-validators", "", false, "Get Request Validators")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetResource, "get-resource", "", false, "Get Resource")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetResources, "get-resources", "", false, "Get Resources")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetRestApi, "get-rest-api", "", false, "Get Rest API")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetRestApis, "get-rest-apis", "", false, "Get Rest Apis")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetSdk, "get-sdk", "", false, "Get Sdk")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetSdkType, "get-sdk-type", "", false, "Get Sdk Type")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetSdkTypes, "get-sdk-types", "", false, "Get Sdk Types")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetStage, "get-stage", "", false, "Get Stage")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetStages, "get-stages", "", false, "Get Stages")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetTags, "get-tags", "", false, "Get Tags")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetUsage, "get-usage", "", false, "Get Usage")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetUsagePlan, "get-usage-plan", "", false, "Get Usage Plan")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetUsagePlanKey, "get-usage-plan-key", "", false, "Get Usage Plan Key")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetUsagePlanKeys, "get-usage-plan-keys", "", false, "Get Usage Plan Keys")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetUsagePlans, "get-usage-plans", "", false, "Get Usage Plans")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetVpcLink, "get-vpc-link", "", false, "Get VPC Link")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayGetVpcLinks, "get-vpc-links", "", false, "Get VPC Links")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayImportApiKeys, "import-api-keys", "", false, "Import API Keys")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayImportDocumentationParts, "import-documentation-parts", "", false, "Import Documentation Parts")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayImportRestApi, "import-rest-api", "", false, "Import Rest API")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayPutGatewayResponse, "put-gateway-response", "", false, "Put Gateway Response")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayPutIntegration, "put-integration", "", false, "Put Integration")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayPutIntegrationResponse, "put-integration-response", "", false, "Put Integration Response")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayPutMethod, "put-method", "", false, "Put Method")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayPutMethodResponse, "put-method-response", "", false, "Put Method Response")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayPutRestApi, "put-rest-api", "", false, "Put Rest API")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayRejectDomainNameAccessAssociation, "reject-domain-name-access-association", "", false, "Reject Domain Name Access Association")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayTagResource, "tag-resource", "", false, "Tag Resource")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayTestInvokeAuthorizer, "test-invoke-authorizer", "", false, "Test Invoke Authorizer")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayTestInvokeMethod, "test-invoke-method", "", false, "Test Invoke Method")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayUntagResource, "untag-resource", "", false, "Untag Resource")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayUpdateAccount, "update-account", "", false, "Update Account")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayUpdateApiKey, "update-api-key", "", false, "Update API Key")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayUpdateAuthorizer, "update-authorizer", "", false, "Update Authorizer")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayUpdateBasePathMapping, "update-base-path-mapping", "", false, "Update Base Path Mapping")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayUpdateClientCertificate, "update-client-certificate", "", false, "Update Client Certificate")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayUpdateDeployment, "update-deployment", "", false, "Update Deployment")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayUpdateDocumentationPart, "update-documentation-part", "", false, "Update Documentation Part")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayUpdateDocumentationVersion, "update-documentation-version", "", false, "Update Documentation Version")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayUpdateDomainName, "update-domain-name", "", false, "Update Domain Name")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayUpdateGatewayResponse, "update-gateway-response", "", false, "Update Gateway Response")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayUpdateIntegration, "update-integration", "", false, "Update Integration")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayUpdateIntegrationResponse, "update-integration-response", "", false, "Update Integration Response")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayUpdateMethod, "update-method", "", false, "Update Method")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayUpdateMethodResponse, "update-method-response", "", false, "Update Method Response")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayUpdateModel, "update-model", "", false, "Update Model")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayUpdateRequestValidator, "update-request-validator", "", false, "Update Request Validator")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayUpdateResource, "update-resource", "", false, "Update Resource")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayUpdateRestApi, "update-rest-api", "", false, "Update Rest API")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayUpdateStage, "update-stage", "", false, "Update Stage")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayUpdateUsage, "update-usage", "", false, "Update Usage")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayUpdateUsagePlan, "update-usage-plan", "", false, "Update Usage Plan")
	_apigatewayCmd.Flags().BoolVarP(&_apigatewayUpdateVpcLink, "update-vpc-link", "", false, "Update VPC Link")

}
