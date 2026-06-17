package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// apigatewayv2Cmd represents the apigatewayv2 command
var _apigatewayv2Cmd = &cobra.Command{
	Use:   "apigatewayv2",
	Short: "AWS apigatewayv2 CLI",
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
		client := apigatewayv2.NewFromConfig(cfg)
		if _apigatewayv2CreateApi {
			apigatewayv2_CreateApi(cfg, client)
			return
		}
		if _apigatewayv2CreateApiMapping {
			apigatewayv2_CreateApiMapping(cfg, client)
			return
		}
		if _apigatewayv2CreateAuthorizer {
			apigatewayv2_CreateAuthorizer(cfg, client)
			return
		}
		if _apigatewayv2CreateDeployment {
			apigatewayv2_CreateDeployment(cfg, client)
			return
		}
		if _apigatewayv2CreateDomainName {
			apigatewayv2_CreateDomainName(cfg, client)
			return
		}
		if _apigatewayv2CreateIntegration {
			apigatewayv2_CreateIntegration(cfg, client)
			return
		}
		if _apigatewayv2CreateIntegrationResponse {
			apigatewayv2_CreateIntegrationResponse(cfg, client)
			return
		}
		if _apigatewayv2CreateModel {
			apigatewayv2_CreateModel(cfg, client)
			return
		}
		if _apigatewayv2CreatePortal {
			apigatewayv2_CreatePortal(cfg, client)
			return
		}
		if _apigatewayv2CreatePortalProduct {
			apigatewayv2_CreatePortalProduct(cfg, client)
			return
		}
		if _apigatewayv2CreateProductPage {
			apigatewayv2_CreateProductPage(cfg, client)
			return
		}
		if _apigatewayv2CreateProductRestEndpointPage {
			apigatewayv2_CreateProductRestEndpointPage(cfg, client)
			return
		}
		if _apigatewayv2CreateRoute {
			apigatewayv2_CreateRoute(cfg, client)
			return
		}
		if _apigatewayv2CreateRouteResponse {
			apigatewayv2_CreateRouteResponse(cfg, client)
			return
		}
		if _apigatewayv2CreateRoutingRule {
			apigatewayv2_CreateRoutingRule(cfg, client)
			return
		}
		if _apigatewayv2CreateStage {
			apigatewayv2_CreateStage(cfg, client)
			return
		}
		if _apigatewayv2CreateVpcLink {
			apigatewayv2_CreateVpcLink(cfg, client)
			return
		}
		if _apigatewayv2DeleteAccessLogSettings {
			apigatewayv2_DeleteAccessLogSettings(cfg, client)
			return
		}
		if _apigatewayv2DeleteApi {
			apigatewayv2_DeleteApi(cfg, client)
			return
		}
		if _apigatewayv2DeleteApiMapping {
			apigatewayv2_DeleteApiMapping(cfg, client)
			return
		}
		if _apigatewayv2DeleteAuthorizer {
			apigatewayv2_DeleteAuthorizer(cfg, client)
			return
		}
		if _apigatewayv2DeleteCorsConfiguration {
			apigatewayv2_DeleteCorsConfiguration(cfg, client)
			return
		}
		if _apigatewayv2DeleteDeployment {
			apigatewayv2_DeleteDeployment(cfg, client)
			return
		}
		if _apigatewayv2DeleteDomainName {
			apigatewayv2_DeleteDomainName(cfg, client)
			return
		}
		if _apigatewayv2DeleteIntegration {
			apigatewayv2_DeleteIntegration(cfg, client)
			return
		}
		if _apigatewayv2DeleteIntegrationResponse {
			apigatewayv2_DeleteIntegrationResponse(cfg, client)
			return
		}
		if _apigatewayv2DeleteModel {
			apigatewayv2_DeleteModel(cfg, client)
			return
		}
		if _apigatewayv2DeletePortal {
			apigatewayv2_DeletePortal(cfg, client)
			return
		}
		if _apigatewayv2DeletePortalProduct {
			apigatewayv2_DeletePortalProduct(cfg, client)
			return
		}
		if _apigatewayv2DeletePortalProductSharingPolicy {
			apigatewayv2_DeletePortalProductSharingPolicy(cfg, client)
			return
		}
		if _apigatewayv2DeleteProductPage {
			apigatewayv2_DeleteProductPage(cfg, client)
			return
		}
		if _apigatewayv2DeleteProductRestEndpointPage {
			apigatewayv2_DeleteProductRestEndpointPage(cfg, client)
			return
		}
		if _apigatewayv2DeleteRoute {
			apigatewayv2_DeleteRoute(cfg, client)
			return
		}
		if _apigatewayv2DeleteRouteRequestParameter {
			apigatewayv2_DeleteRouteRequestParameter(cfg, client)
			return
		}
		if _apigatewayv2DeleteRouteResponse {
			apigatewayv2_DeleteRouteResponse(cfg, client)
			return
		}
		if _apigatewayv2DeleteRouteSettings {
			apigatewayv2_DeleteRouteSettings(cfg, client)
			return
		}
		if _apigatewayv2DeleteRoutingRule {
			apigatewayv2_DeleteRoutingRule(cfg, client)
			return
		}
		if _apigatewayv2DeleteStage {
			apigatewayv2_DeleteStage(cfg, client)
			return
		}
		if _apigatewayv2DeleteVpcLink {
			apigatewayv2_DeleteVpcLink(cfg, client)
			return
		}
		if _apigatewayv2DisablePortal {
			apigatewayv2_DisablePortal(cfg, client)
			return
		}
		if _apigatewayv2ExportApi {
			apigatewayv2_ExportApi(cfg, client)
			return
		}
		if _apigatewayv2GetApi {
			apigatewayv2_GetApi(cfg, client)
			return
		}
		if _apigatewayv2GetApiMapping {
			apigatewayv2_GetApiMapping(cfg, client)
			return
		}
		if _apigatewayv2GetApiMappings {
			apigatewayv2_GetApiMappings(cfg, client)
			return
		}
		if _apigatewayv2GetApis {
			apigatewayv2_GetApis(cfg, client)
			return
		}
		if _apigatewayv2GetAuthorizer {
			apigatewayv2_GetAuthorizer(cfg, client)
			return
		}
		if _apigatewayv2GetAuthorizers {
			apigatewayv2_GetAuthorizers(cfg, client)
			return
		}
		if _apigatewayv2GetDeployment {
			apigatewayv2_GetDeployment(cfg, client)
			return
		}
		if _apigatewayv2GetDeployments {
			apigatewayv2_GetDeployments(cfg, client)
			return
		}
		if _apigatewayv2GetDomainName {
			apigatewayv2_GetDomainName(cfg, client)
			return
		}
		if _apigatewayv2GetDomainNames {
			apigatewayv2_GetDomainNames(cfg, client)
			return
		}
		if _apigatewayv2GetIntegration {
			apigatewayv2_GetIntegration(cfg, client)
			return
		}
		if _apigatewayv2GetIntegrationResponse {
			apigatewayv2_GetIntegrationResponse(cfg, client)
			return
		}
		if _apigatewayv2GetIntegrationResponses {
			apigatewayv2_GetIntegrationResponses(cfg, client)
			return
		}
		if _apigatewayv2GetIntegrations {
			apigatewayv2_GetIntegrations(cfg, client)
			return
		}
		if _apigatewayv2GetModel {
			apigatewayv2_GetModel(cfg, client)
			return
		}
		if _apigatewayv2GetModelTemplate {
			apigatewayv2_GetModelTemplate(cfg, client)
			return
		}
		if _apigatewayv2GetModels {
			apigatewayv2_GetModels(cfg, client)
			return
		}
		if _apigatewayv2GetPortal {
			apigatewayv2_GetPortal(cfg, client)
			return
		}
		if _apigatewayv2GetPortalProduct {
			apigatewayv2_GetPortalProduct(cfg, client)
			return
		}
		if _apigatewayv2GetPortalProductSharingPolicy {
			apigatewayv2_GetPortalProductSharingPolicy(cfg, client)
			return
		}
		if _apigatewayv2GetProductPage {
			apigatewayv2_GetProductPage(cfg, client)
			return
		}
		if _apigatewayv2GetProductRestEndpointPage {
			apigatewayv2_GetProductRestEndpointPage(cfg, client)
			return
		}
		if _apigatewayv2GetRoute {
			apigatewayv2_GetRoute(cfg, client)
			return
		}
		if _apigatewayv2GetRouteResponse {
			apigatewayv2_GetRouteResponse(cfg, client)
			return
		}
		if _apigatewayv2GetRouteResponses {
			apigatewayv2_GetRouteResponses(cfg, client)
			return
		}
		if _apigatewayv2GetRoutes {
			apigatewayv2_GetRoutes(cfg, client)
			return
		}
		if _apigatewayv2GetRoutingRule {
			apigatewayv2_GetRoutingRule(cfg, client)
			return
		}
		if _apigatewayv2GetStage {
			apigatewayv2_GetStage(cfg, client)
			return
		}
		if _apigatewayv2GetStages {
			apigatewayv2_GetStages(cfg, client)
			return
		}
		if _apigatewayv2GetTags {
			apigatewayv2_GetTags(cfg, client)
			return
		}
		if _apigatewayv2GetVpcLink {
			apigatewayv2_GetVpcLink(cfg, client)
			return
		}
		if _apigatewayv2GetVpcLinks {
			apigatewayv2_GetVpcLinks(cfg, client)
			return
		}
		if _apigatewayv2ImportApi {
			apigatewayv2_ImportApi(cfg, client)
			return
		}
		if _apigatewayv2ListPortalProducts {
			apigatewayv2_ListPortalProducts(cfg, client)
			return
		}
		if _apigatewayv2ListPortals {
			apigatewayv2_ListPortals(cfg, client)
			return
		}
		if _apigatewayv2ListProductPages {
			apigatewayv2_ListProductPages(cfg, client)
			return
		}
		if _apigatewayv2ListProductRestEndpointPages {
			apigatewayv2_ListProductRestEndpointPages(cfg, client)
			return
		}
		if _apigatewayv2ListRoutingRules {
			apigatewayv2_ListRoutingRules(cfg, client)
			return
		}
		if _apigatewayv2PreviewPortal {
			apigatewayv2_PreviewPortal(cfg, client)
			return
		}
		if _apigatewayv2PublishPortal {
			apigatewayv2_PublishPortal(cfg, client)
			return
		}
		if _apigatewayv2PutPortalProductSharingPolicy {
			apigatewayv2_PutPortalProductSharingPolicy(cfg, client)
			return
		}
		if _apigatewayv2PutRoutingRule {
			apigatewayv2_PutRoutingRule(cfg, client)
			return
		}
		if _apigatewayv2ReimportApi {
			apigatewayv2_ReimportApi(cfg, client)
			return
		}
		if _apigatewayv2ResetAuthorizersCache {
			apigatewayv2_ResetAuthorizersCache(cfg, client)
			return
		}
		if _apigatewayv2TagResource {
			apigatewayv2_TagResource(cfg, client)
			return
		}
		if _apigatewayv2UntagResource {
			apigatewayv2_UntagResource(cfg, client)
			return
		}
		if _apigatewayv2UpdateApi {
			apigatewayv2_UpdateApi(cfg, client)
			return
		}
		if _apigatewayv2UpdateApiMapping {
			apigatewayv2_UpdateApiMapping(cfg, client)
			return
		}
		if _apigatewayv2UpdateAuthorizer {
			apigatewayv2_UpdateAuthorizer(cfg, client)
			return
		}
		if _apigatewayv2UpdateDeployment {
			apigatewayv2_UpdateDeployment(cfg, client)
			return
		}
		if _apigatewayv2UpdateDomainName {
			apigatewayv2_UpdateDomainName(cfg, client)
			return
		}
		if _apigatewayv2UpdateIntegration {
			apigatewayv2_UpdateIntegration(cfg, client)
			return
		}
		if _apigatewayv2UpdateIntegrationResponse {
			apigatewayv2_UpdateIntegrationResponse(cfg, client)
			return
		}
		if _apigatewayv2UpdateModel {
			apigatewayv2_UpdateModel(cfg, client)
			return
		}
		if _apigatewayv2UpdatePortal {
			apigatewayv2_UpdatePortal(cfg, client)
			return
		}
		if _apigatewayv2UpdatePortalProduct {
			apigatewayv2_UpdatePortalProduct(cfg, client)
			return
		}
		if _apigatewayv2UpdateProductPage {
			apigatewayv2_UpdateProductPage(cfg, client)
			return
		}
		if _apigatewayv2UpdateProductRestEndpointPage {
			apigatewayv2_UpdateProductRestEndpointPage(cfg, client)
			return
		}
		if _apigatewayv2UpdateRoute {
			apigatewayv2_UpdateRoute(cfg, client)
			return
		}
		if _apigatewayv2UpdateRouteResponse {
			apigatewayv2_UpdateRouteResponse(cfg, client)
			return
		}
		if _apigatewayv2UpdateStage {
			apigatewayv2_UpdateStage(cfg, client)
			return
		}
		if _apigatewayv2UpdateVpcLink {
			apigatewayv2_UpdateVpcLink(cfg, client)
			return
		}

	},
}

var (
	_apigatewayv2CreateApi                        bool
	_apigatewayv2CreateApiMapping                 bool
	_apigatewayv2CreateAuthorizer                 bool
	_apigatewayv2CreateDeployment                 bool
	_apigatewayv2CreateDomainName                 bool
	_apigatewayv2CreateIntegration                bool
	_apigatewayv2CreateIntegrationResponse        bool
	_apigatewayv2CreateModel                      bool
	_apigatewayv2CreatePortal                     bool
	_apigatewayv2CreatePortalProduct              bool
	_apigatewayv2CreateProductPage                bool
	_apigatewayv2CreateProductRestEndpointPage    bool
	_apigatewayv2CreateRoute                      bool
	_apigatewayv2CreateRouteResponse              bool
	_apigatewayv2CreateRoutingRule                bool
	_apigatewayv2CreateStage                      bool
	_apigatewayv2CreateVpcLink                    bool
	_apigatewayv2DeleteAccessLogSettings          bool
	_apigatewayv2DeleteApi                        bool
	_apigatewayv2DeleteApiMapping                 bool
	_apigatewayv2DeleteAuthorizer                 bool
	_apigatewayv2DeleteCorsConfiguration          bool
	_apigatewayv2DeleteDeployment                 bool
	_apigatewayv2DeleteDomainName                 bool
	_apigatewayv2DeleteIntegration                bool
	_apigatewayv2DeleteIntegrationResponse        bool
	_apigatewayv2DeleteModel                      bool
	_apigatewayv2DeletePortal                     bool
	_apigatewayv2DeletePortalProduct              bool
	_apigatewayv2DeletePortalProductSharingPolicy bool
	_apigatewayv2DeleteProductPage                bool
	_apigatewayv2DeleteProductRestEndpointPage    bool
	_apigatewayv2DeleteRoute                      bool
	_apigatewayv2DeleteRouteRequestParameter      bool
	_apigatewayv2DeleteRouteResponse              bool
	_apigatewayv2DeleteRouteSettings              bool
	_apigatewayv2DeleteRoutingRule                bool
	_apigatewayv2DeleteStage                      bool
	_apigatewayv2DeleteVpcLink                    bool
	_apigatewayv2DisablePortal                    bool
	_apigatewayv2ExportApi                        bool
	_apigatewayv2GetApi                           bool
	_apigatewayv2GetApiMapping                    bool
	_apigatewayv2GetApiMappings                   bool
	_apigatewayv2GetApis                          bool
	_apigatewayv2GetAuthorizer                    bool
	_apigatewayv2GetAuthorizers                   bool
	_apigatewayv2GetDeployment                    bool
	_apigatewayv2GetDeployments                   bool
	_apigatewayv2GetDomainName                    bool
	_apigatewayv2GetDomainNames                   bool
	_apigatewayv2GetIntegration                   bool
	_apigatewayv2GetIntegrationResponse           bool
	_apigatewayv2GetIntegrationResponses          bool
	_apigatewayv2GetIntegrations                  bool
	_apigatewayv2GetModel                         bool
	_apigatewayv2GetModelTemplate                 bool
	_apigatewayv2GetModels                        bool
	_apigatewayv2GetPortal                        bool
	_apigatewayv2GetPortalProduct                 bool
	_apigatewayv2GetPortalProductSharingPolicy    bool
	_apigatewayv2GetProductPage                   bool
	_apigatewayv2GetProductRestEndpointPage       bool
	_apigatewayv2GetRoute                         bool
	_apigatewayv2GetRouteResponse                 bool
	_apigatewayv2GetRouteResponses                bool
	_apigatewayv2GetRoutes                        bool
	_apigatewayv2GetRoutingRule                   bool
	_apigatewayv2GetStage                         bool
	_apigatewayv2GetStages                        bool
	_apigatewayv2GetTags                          bool
	_apigatewayv2GetVpcLink                       bool
	_apigatewayv2GetVpcLinks                      bool
	_apigatewayv2ImportApi                        bool
	_apigatewayv2ListPortalProducts               bool
	_apigatewayv2ListPortals                      bool
	_apigatewayv2ListProductPages                 bool
	_apigatewayv2ListProductRestEndpointPages     bool
	_apigatewayv2ListRoutingRules                 bool
	_apigatewayv2PreviewPortal                    bool
	_apigatewayv2PublishPortal                    bool
	_apigatewayv2PutPortalProductSharingPolicy    bool
	_apigatewayv2PutRoutingRule                   bool
	_apigatewayv2ReimportApi                      bool
	_apigatewayv2ResetAuthorizersCache            bool
	_apigatewayv2TagResource                      bool
	_apigatewayv2UntagResource                    bool
	_apigatewayv2UpdateApi                        bool
	_apigatewayv2UpdateApiMapping                 bool
	_apigatewayv2UpdateAuthorizer                 bool
	_apigatewayv2UpdateDeployment                 bool
	_apigatewayv2UpdateDomainName                 bool
	_apigatewayv2UpdateIntegration                bool
	_apigatewayv2UpdateIntegrationResponse        bool
	_apigatewayv2UpdateModel                      bool
	_apigatewayv2UpdatePortal                     bool
	_apigatewayv2UpdatePortalProduct              bool
	_apigatewayv2UpdateProductPage                bool
	_apigatewayv2UpdateProductRestEndpointPage    bool
	_apigatewayv2UpdateRoute                      bool
	_apigatewayv2UpdateRouteResponse              bool
	_apigatewayv2UpdateStage                      bool
	_apigatewayv2UpdateVpcLink                    bool

	_apigatewayv2AccessLogSettings                string
	_apigatewayv2Actions                          string
	_apigatewayv2ApiId                            string
	_apigatewayv2ApiKeyRequired                   string
	_apigatewayv2ApiKeySelectionExpression        string
	_apigatewayv2ApiMappingId                     string
	_apigatewayv2ApiMappingKey                    string
	_apigatewayv2Authorization                    string
	_apigatewayv2AuthorizationScopes              []string
	_apigatewayv2AuthorizationType                string
	_apigatewayv2AuthorizerCredentialsArn         string
	_apigatewayv2AuthorizerId                     string
	_apigatewayv2AuthorizerPayloadFormatVersion   string
	_apigatewayv2AuthorizerResultTtlInSeconds     string
	_apigatewayv2AuthorizerType                   string
	_apigatewayv2AuthorizerUri                    string
	_apigatewayv2AutoDeploy                       string
	_apigatewayv2Basepath                         string
	_apigatewayv2Body                             string
	_apigatewayv2ClientCertificateId              string
	_apigatewayv2Conditions                       string
	_apigatewayv2ConnectionId                     string
	_apigatewayv2ConnectionType                   string
	_apigatewayv2ContentHandlingStrategy          string
	_apigatewayv2ContentType                      string
	_apigatewayv2CorsConfiguration                string
	_apigatewayv2CredentialsArn                   string
	_apigatewayv2DefaultRouteSettings             string
	_apigatewayv2DeploymentId                     string
	_apigatewayv2Description                      string
	_apigatewayv2DisableExecuteApiEndpoint        string
	_apigatewayv2DisableSchemaValidation          string
	_apigatewayv2DisplayContent                   string
	_apigatewayv2DisplayName                      string
	_apigatewayv2DisplayOrder                     string
	_apigatewayv2DomainName                       string
	_apigatewayv2DomainNameConfigurations         string
	_apigatewayv2DomainNameId                     string
	_apigatewayv2EnableSimpleResponses            string
	_apigatewayv2EndpointConfiguration            string
	_apigatewayv2ExportVersion                    string
	_apigatewayv2FailOnWarnings                   string
	_apigatewayv2IdentitySource                   []string
	_apigatewayv2IdentityValidationExpression     string
	_apigatewayv2IncludeExtensions                string
	_apigatewayv2IncludeRawDisplayContent         string
	_apigatewayv2IncludedPortalProductArns        []string
	_apigatewayv2IntegrationId                    string
	_apigatewayv2IntegrationMethod                string
	_apigatewayv2IntegrationResponseId            string
	_apigatewayv2IntegrationResponseKey           string
	_apigatewayv2IntegrationSubtype               string
	_apigatewayv2IntegrationType                  string
	_apigatewayv2IntegrationUri                   string
	_apigatewayv2IpAddressType                    string
	_apigatewayv2JwtConfiguration                 string
	_apigatewayv2LogoUri                          string
	_apigatewayv2MaxResults                       string
	_apigatewayv2ModelId                          string
	_apigatewayv2ModelSelectionExpression         string
	_apigatewayv2MutualTlsAuthentication          string
	_apigatewayv2Name                             string
	_apigatewayv2NextToken                        string
	_apigatewayv2OperationName                    string
	_apigatewayv2OutputType                       string
	_apigatewayv2PassthroughBehavior              string
	_apigatewayv2PayloadFormatVersion             string
	_apigatewayv2PolicyDocument                   string
	_apigatewayv2PortalContent                    string
	_apigatewayv2PortalId                         string
	_apigatewayv2PortalProductId                  string
	_apigatewayv2Priority                         string
	_apigatewayv2ProductPageId                    string
	_apigatewayv2ProductRestEndpointPageId        string
	_apigatewayv2ProtocolType                     string
	_apigatewayv2RequestModels                    string
	_apigatewayv2RequestParameterKey              string
	_apigatewayv2RequestParameters                string
	_apigatewayv2RequestTemplates                 string
	_apigatewayv2ResourceArn                      string
	_apigatewayv2ResourceOwner                    string
	_apigatewayv2ResourceOwnerAccountId           string
	_apigatewayv2ResponseModels                   string
	_apigatewayv2ResponseParameters               string
	_apigatewayv2ResponseTemplates                string
	_apigatewayv2RestEndpointIdentifier           string
	_apigatewayv2RouteId                          string
	_apigatewayv2RouteKey                         string
	_apigatewayv2RouteResponseId                  string
	_apigatewayv2RouteResponseKey                 string
	_apigatewayv2RouteResponseSelectionExpression string
	_apigatewayv2RouteSelectionExpression         string
	_apigatewayv2RouteSettings                    string
	_apigatewayv2RoutingMode                      string
	_apigatewayv2RoutingRuleId                    string
	_apigatewayv2RumAppMonitorName                string
	_apigatewayv2Schema                           string
	_apigatewayv2SecurityGroupIds                 []string
	_apigatewayv2Specification                    string
	_apigatewayv2Stage                            string
	_apigatewayv2StageName                        string
	_apigatewayv2StageVariables                   string
	_apigatewayv2SubnetIds                        []string
	_apigatewayv2TagKeys                          []string
	_apigatewayv2Tags                             string
	_apigatewayv2Target                           string
	_apigatewayv2TemplateSelectionExpression      string
	_apigatewayv2TimeoutInMillis                  string
	_apigatewayv2TlsConfig                        string
	_apigatewayv2TryItState                       string
	_apigatewayv2Version                          string
	_apigatewayv2VpcLinkId                        string
)

// Creates an Api resource.
func apigatewayv2_CreateApi(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.CreateApiInput{
		// Name: *string, // Required
		// ProtocolType: types.ProtocolType, // Required
	}

	if len(_apigatewayv2Name) > 0 {
		input.Name = aws.String(_apigatewayv2Name)
	}
	if len(_apigatewayv2ProtocolType) > 0 {
		if err := assignInputField(input, "ProtocolType", _apigatewayv2ProtocolType); err != nil {
			log.Errorf("invalid --protocol-type: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2ApiKeySelectionExpression) > 0 {
		input.ApiKeySelectionExpression = aws.String(_apigatewayv2ApiKeySelectionExpression)
	}
	if len(_apigatewayv2CorsConfiguration) > 0 {
		if err := assignInputField(input, "CorsConfiguration", _apigatewayv2CorsConfiguration); err != nil {
			log.Errorf("invalid --cors-configuration: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2CredentialsArn) > 0 {
		input.CredentialsArn = aws.String(_apigatewayv2CredentialsArn)
	}
	if len(_apigatewayv2Description) > 0 {
		input.Description = aws.String(_apigatewayv2Description)
	}
	if len(_apigatewayv2DisableExecuteApiEndpoint) > 0 {
		if err := assignInputField(input, "DisableExecuteApiEndpoint", _apigatewayv2DisableExecuteApiEndpoint); err != nil {
			log.Errorf("invalid --disable-execute-api-endpoint: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2DisableSchemaValidation) > 0 {
		if err := assignInputField(input, "DisableSchemaValidation", _apigatewayv2DisableSchemaValidation); err != nil {
			log.Errorf("invalid --disable-schema-validation: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2IpAddressType) > 0 {
		if err := assignInputField(input, "IpAddressType", _apigatewayv2IpAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2RouteKey) > 0 {
		input.RouteKey = aws.String(_apigatewayv2RouteKey)
	}
	if len(_apigatewayv2RouteSelectionExpression) > 0 {
		input.RouteSelectionExpression = aws.String(_apigatewayv2RouteSelectionExpression)
	}
	if len(_apigatewayv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _apigatewayv2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2Target) > 0 {
		input.Target = aws.String(_apigatewayv2Target)
	}
	if len(_apigatewayv2Version) > 0 {
		input.Version = aws.String(_apigatewayv2Version)
	}

	if resp, err := client.CreateApi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an API mapping.
func apigatewayv2_CreateApiMapping(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.CreateApiMappingInput{
		// ApiId: *string, // Required
		// DomainName: *string, // Required
		// Stage: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2DomainName) > 0 {
		input.DomainName = aws.String(_apigatewayv2DomainName)
	}
	if len(_apigatewayv2Stage) > 0 {
		input.Stage = aws.String(_apigatewayv2Stage)
	}
	if len(_apigatewayv2ApiMappingKey) > 0 {
		input.ApiMappingKey = aws.String(_apigatewayv2ApiMappingKey)
	}

	if resp, err := client.CreateApiMapping(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Authorizer for an API.
func apigatewayv2_CreateAuthorizer(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.CreateAuthorizerInput{
		// ApiId: *string, // Required
		// AuthorizerType: types.AuthorizerType, // Required
		// IdentitySource: []string, // Required
		// Name: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2AuthorizerType) > 0 {
		if err := assignInputField(input, "AuthorizerType", _apigatewayv2AuthorizerType); err != nil {
			log.Errorf("invalid --authorizer-type: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2IdentitySource) > 0 {
		input.IdentitySource = append([]string(nil), _apigatewayv2IdentitySource...)
	}
	if len(_apigatewayv2Name) > 0 {
		input.Name = aws.String(_apigatewayv2Name)
	}
	if len(_apigatewayv2AuthorizerCredentialsArn) > 0 {
		input.AuthorizerCredentialsArn = aws.String(_apigatewayv2AuthorizerCredentialsArn)
	}
	if len(_apigatewayv2AuthorizerPayloadFormatVersion) > 0 {
		input.AuthorizerPayloadFormatVersion = aws.String(_apigatewayv2AuthorizerPayloadFormatVersion)
	}
	if len(_apigatewayv2AuthorizerResultTtlInSeconds) > 0 {
		if err := assignInputField(input, "AuthorizerResultTtlInSeconds", _apigatewayv2AuthorizerResultTtlInSeconds); err != nil {
			log.Errorf("invalid --authorizer-result-ttl-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2AuthorizerUri) > 0 {
		input.AuthorizerUri = aws.String(_apigatewayv2AuthorizerUri)
	}
	if len(_apigatewayv2EnableSimpleResponses) > 0 {
		if err := assignInputField(input, "EnableSimpleResponses", _apigatewayv2EnableSimpleResponses); err != nil {
			log.Errorf("invalid --enable-simple-responses: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2IdentityValidationExpression) > 0 {
		input.IdentityValidationExpression = aws.String(_apigatewayv2IdentityValidationExpression)
	}
	if len(_apigatewayv2JwtConfiguration) > 0 {
		if err := assignInputField(input, "JwtConfiguration", _apigatewayv2JwtConfiguration); err != nil {
			log.Errorf("invalid --jwt-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAuthorizer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Deployment for an API.
func apigatewayv2_CreateDeployment(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.CreateDeploymentInput{
		// ApiId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2Description) > 0 {
		input.Description = aws.String(_apigatewayv2Description)
	}
	if len(_apigatewayv2StageName) > 0 {
		input.StageName = aws.String(_apigatewayv2StageName)
	}

	if resp, err := client.CreateDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a domain name.
func apigatewayv2_CreateDomainName(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.CreateDomainNameInput{
		// DomainName: *string, // Required
	}

	if len(_apigatewayv2DomainName) > 0 {
		input.DomainName = aws.String(_apigatewayv2DomainName)
	}
	if len(_apigatewayv2DomainNameConfigurations) > 0 {
		if err := assignInputField(input, "DomainNameConfigurations", _apigatewayv2DomainNameConfigurations); err != nil {
			log.Errorf("invalid --domain-name-configurations: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2MutualTlsAuthentication) > 0 {
		if err := assignInputField(input, "MutualTlsAuthentication", _apigatewayv2MutualTlsAuthentication); err != nil {
			log.Errorf("invalid --mutual-tls-authentication: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2RoutingMode) > 0 {
		if err := assignInputField(input, "RoutingMode", _apigatewayv2RoutingMode); err != nil {
			log.Errorf("invalid --routing-mode: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _apigatewayv2Tags); err != nil {
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

// Creates an Integration.
func apigatewayv2_CreateIntegration(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.CreateIntegrationInput{
		// ApiId: *string, // Required
		// IntegrationType: types.IntegrationType, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2IntegrationType) > 0 {
		if err := assignInputField(input, "IntegrationType", _apigatewayv2IntegrationType); err != nil {
			log.Errorf("invalid --integration-type: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2ConnectionId) > 0 {
		input.ConnectionId = aws.String(_apigatewayv2ConnectionId)
	}
	if len(_apigatewayv2ConnectionType) > 0 {
		if err := assignInputField(input, "ConnectionType", _apigatewayv2ConnectionType); err != nil {
			log.Errorf("invalid --connection-type: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2ContentHandlingStrategy) > 0 {
		if err := assignInputField(input, "ContentHandlingStrategy", _apigatewayv2ContentHandlingStrategy); err != nil {
			log.Errorf("invalid --content-handling-strategy: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2CredentialsArn) > 0 {
		input.CredentialsArn = aws.String(_apigatewayv2CredentialsArn)
	}
	if len(_apigatewayv2Description) > 0 {
		input.Description = aws.String(_apigatewayv2Description)
	}
	if len(_apigatewayv2IntegrationMethod) > 0 {
		input.IntegrationMethod = aws.String(_apigatewayv2IntegrationMethod)
	}
	if len(_apigatewayv2IntegrationSubtype) > 0 {
		input.IntegrationSubtype = aws.String(_apigatewayv2IntegrationSubtype)
	}
	if len(_apigatewayv2IntegrationUri) > 0 {
		input.IntegrationUri = aws.String(_apigatewayv2IntegrationUri)
	}
	if len(_apigatewayv2PassthroughBehavior) > 0 {
		if err := assignInputField(input, "PassthroughBehavior", _apigatewayv2PassthroughBehavior); err != nil {
			log.Errorf("invalid --passthrough-behavior: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2PayloadFormatVersion) > 0 {
		input.PayloadFormatVersion = aws.String(_apigatewayv2PayloadFormatVersion)
	}
	if len(_apigatewayv2RequestParameters) > 0 {
		if err := assignInputField(input, "RequestParameters", _apigatewayv2RequestParameters); err != nil {
			log.Errorf("invalid --request-parameters: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2RequestTemplates) > 0 {
		if err := assignInputField(input, "RequestTemplates", _apigatewayv2RequestTemplates); err != nil {
			log.Errorf("invalid --request-templates: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2ResponseParameters) > 0 {
		if err := assignInputField(input, "ResponseParameters", _apigatewayv2ResponseParameters); err != nil {
			log.Errorf("invalid --response-parameters: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2TemplateSelectionExpression) > 0 {
		input.TemplateSelectionExpression = aws.String(_apigatewayv2TemplateSelectionExpression)
	}
	if len(_apigatewayv2TimeoutInMillis) > 0 {
		if err := assignInputField(input, "TimeoutInMillis", _apigatewayv2TimeoutInMillis); err != nil {
			log.Errorf("invalid --timeout-in-millis: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2TlsConfig) > 0 {
		if err := assignInputField(input, "TlsConfig", _apigatewayv2TlsConfig); err != nil {
			log.Errorf("invalid --tls-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an IntegrationResponses.
func apigatewayv2_CreateIntegrationResponse(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.CreateIntegrationResponseInput{
		// ApiId: *string, // Required
		// IntegrationId: *string, // Required
		// IntegrationResponseKey: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2IntegrationId) > 0 {
		input.IntegrationId = aws.String(_apigatewayv2IntegrationId)
	}
	if len(_apigatewayv2IntegrationResponseKey) > 0 {
		input.IntegrationResponseKey = aws.String(_apigatewayv2IntegrationResponseKey)
	}
	if len(_apigatewayv2ContentHandlingStrategy) > 0 {
		if err := assignInputField(input, "ContentHandlingStrategy", _apigatewayv2ContentHandlingStrategy); err != nil {
			log.Errorf("invalid --content-handling-strategy: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2ResponseParameters) > 0 {
		if err := assignInputField(input, "ResponseParameters", _apigatewayv2ResponseParameters); err != nil {
			log.Errorf("invalid --response-parameters: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2ResponseTemplates) > 0 {
		if err := assignInputField(input, "ResponseTemplates", _apigatewayv2ResponseTemplates); err != nil {
			log.Errorf("invalid --response-templates: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2TemplateSelectionExpression) > 0 {
		input.TemplateSelectionExpression = aws.String(_apigatewayv2TemplateSelectionExpression)
	}

	if resp, err := client.CreateIntegrationResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Model for an API.
func apigatewayv2_CreateModel(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.CreateModelInput{
		// ApiId: *string, // Required
		// Name: *string, // Required
		// Schema: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2Name) > 0 {
		input.Name = aws.String(_apigatewayv2Name)
	}
	if len(_apigatewayv2Schema) > 0 {
		input.Schema = aws.String(_apigatewayv2Schema)
	}
	if len(_apigatewayv2ContentType) > 0 {
		input.ContentType = aws.String(_apigatewayv2ContentType)
	}
	if len(_apigatewayv2Description) > 0 {
		input.Description = aws.String(_apigatewayv2Description)
	}

	if resp, err := client.CreateModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a portal.
func apigatewayv2_CreatePortal(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.CreatePortalInput{
		// Authorization: *types.Authorization, // Required
		// EndpointConfiguration: *types.EndpointConfigurationRequest, // Required
		// PortalContent: *types.PortalContent, // Required
	}

	if len(_apigatewayv2Authorization) > 0 {
		if err := assignInputField(input, "Authorization", _apigatewayv2Authorization); err != nil {
			log.Errorf("invalid --authorization: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2EndpointConfiguration) > 0 {
		if err := assignInputField(input, "EndpointConfiguration", _apigatewayv2EndpointConfiguration); err != nil {
			log.Errorf("invalid --endpoint-configuration: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2PortalContent) > 0 {
		if err := assignInputField(input, "PortalContent", _apigatewayv2PortalContent); err != nil {
			log.Errorf("invalid --portal-content: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2IncludedPortalProductArns) > 0 {
		input.IncludedPortalProductArns = append([]string(nil), _apigatewayv2IncludedPortalProductArns...)
	}
	if len(_apigatewayv2LogoUri) > 0 {
		input.LogoUri = aws.String(_apigatewayv2LogoUri)
	}
	if len(_apigatewayv2RumAppMonitorName) > 0 {
		input.RumAppMonitorName = aws.String(_apigatewayv2RumAppMonitorName)
	}
	if len(_apigatewayv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _apigatewayv2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePortal(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new portal product.
func apigatewayv2_CreatePortalProduct(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.CreatePortalProductInput{
		// DisplayName: *string, // Required
	}

	if len(_apigatewayv2DisplayName) > 0 {
		input.DisplayName = aws.String(_apigatewayv2DisplayName)
	}
	if len(_apigatewayv2Description) > 0 {
		input.Description = aws.String(_apigatewayv2Description)
	}
	if len(_apigatewayv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _apigatewayv2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePortalProduct(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new product page for a portal product.
func apigatewayv2_CreateProductPage(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.CreateProductPageInput{
		// DisplayContent: *types.DisplayContent, // Required
		// PortalProductId: *string, // Required
	}

	if len(_apigatewayv2DisplayContent) > 0 {
		if err := assignInputField(input, "DisplayContent", _apigatewayv2DisplayContent); err != nil {
			log.Errorf("invalid --display-content: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2PortalProductId) > 0 {
		input.PortalProductId = aws.String(_apigatewayv2PortalProductId)
	}

	if resp, err := client.CreateProductPage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a product REST endpoint page for a portal product.
func apigatewayv2_CreateProductRestEndpointPage(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.CreateProductRestEndpointPageInput{
		// PortalProductId: *string, // Required
		// RestEndpointIdentifier: *types.RestEndpointIdentifier, // Required
	}

	if len(_apigatewayv2PortalProductId) > 0 {
		input.PortalProductId = aws.String(_apigatewayv2PortalProductId)
	}
	if len(_apigatewayv2RestEndpointIdentifier) > 0 {
		if err := assignInputField(input, "RestEndpointIdentifier", _apigatewayv2RestEndpointIdentifier); err != nil {
			log.Errorf("invalid --rest-endpoint-identifier: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2DisplayContent) > 0 {
		if err := assignInputField(input, "DisplayContent", _apigatewayv2DisplayContent); err != nil {
			log.Errorf("invalid --display-content: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2TryItState) > 0 {
		if err := assignInputField(input, "TryItState", _apigatewayv2TryItState); err != nil {
			log.Errorf("invalid --try-it-state: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProductRestEndpointPage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Route for an API.
func apigatewayv2_CreateRoute(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.CreateRouteInput{
		// ApiId: *string, // Required
		// RouteKey: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2RouteKey) > 0 {
		input.RouteKey = aws.String(_apigatewayv2RouteKey)
	}
	if len(_apigatewayv2ApiKeyRequired) > 0 {
		if err := assignInputField(input, "ApiKeyRequired", _apigatewayv2ApiKeyRequired); err != nil {
			log.Errorf("invalid --api-key-required: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2AuthorizationScopes) > 0 {
		input.AuthorizationScopes = append([]string(nil), _apigatewayv2AuthorizationScopes...)
	}
	if len(_apigatewayv2AuthorizationType) > 0 {
		if err := assignInputField(input, "AuthorizationType", _apigatewayv2AuthorizationType); err != nil {
			log.Errorf("invalid --authorization-type: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2AuthorizerId) > 0 {
		input.AuthorizerId = aws.String(_apigatewayv2AuthorizerId)
	}
	if len(_apigatewayv2ModelSelectionExpression) > 0 {
		input.ModelSelectionExpression = aws.String(_apigatewayv2ModelSelectionExpression)
	}
	if len(_apigatewayv2OperationName) > 0 {
		input.OperationName = aws.String(_apigatewayv2OperationName)
	}
	if len(_apigatewayv2RequestModels) > 0 {
		if err := assignInputField(input, "RequestModels", _apigatewayv2RequestModels); err != nil {
			log.Errorf("invalid --request-models: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2RequestParameters) > 0 {
		if err := assignInputField(input, "RequestParameters", _apigatewayv2RequestParameters); err != nil {
			log.Errorf("invalid --request-parameters: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2RouteResponseSelectionExpression) > 0 {
		input.RouteResponseSelectionExpression = aws.String(_apigatewayv2RouteResponseSelectionExpression)
	}
	if len(_apigatewayv2Target) > 0 {
		input.Target = aws.String(_apigatewayv2Target)
	}

	if resp, err := client.CreateRoute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a RouteResponse for a Route.
func apigatewayv2_CreateRouteResponse(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.CreateRouteResponseInput{
		// ApiId: *string, // Required
		// RouteId: *string, // Required
		// RouteResponseKey: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2RouteId) > 0 {
		input.RouteId = aws.String(_apigatewayv2RouteId)
	}
	if len(_apigatewayv2RouteResponseKey) > 0 {
		input.RouteResponseKey = aws.String(_apigatewayv2RouteResponseKey)
	}
	if len(_apigatewayv2ModelSelectionExpression) > 0 {
		input.ModelSelectionExpression = aws.String(_apigatewayv2ModelSelectionExpression)
	}
	if len(_apigatewayv2ResponseModels) > 0 {
		if err := assignInputField(input, "ResponseModels", _apigatewayv2ResponseModels); err != nil {
			log.Errorf("invalid --response-models: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2ResponseParameters) > 0 {
		if err := assignInputField(input, "ResponseParameters", _apigatewayv2ResponseParameters); err != nil {
			log.Errorf("invalid --response-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRouteResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a RoutingRule.
func apigatewayv2_CreateRoutingRule(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.CreateRoutingRuleInput{
		// Actions: []types.RoutingRuleAction, // Required
		// Conditions: []types.RoutingRuleCondition, // Required
		// DomainName: *string, // Required
		// Priority: *int32, // Required
	}

	if len(_apigatewayv2Actions) > 0 {
		if err := assignInputField(input, "Actions", _apigatewayv2Actions); err != nil {
			log.Errorf("invalid --actions: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2Conditions) > 0 {
		if err := assignInputField(input, "Conditions", _apigatewayv2Conditions); err != nil {
			log.Errorf("invalid --conditions: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2DomainName) > 0 {
		input.DomainName = aws.String(_apigatewayv2DomainName)
	}
	if len(_apigatewayv2Priority) > 0 {
		if err := assignInputField(input, "Priority", _apigatewayv2Priority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2DomainNameId) > 0 {
		input.DomainNameId = aws.String(_apigatewayv2DomainNameId)
	}

	if resp, err := client.CreateRoutingRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Stage for an API.
func apigatewayv2_CreateStage(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.CreateStageInput{
		// ApiId: *string, // Required
		// StageName: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2StageName) > 0 {
		input.StageName = aws.String(_apigatewayv2StageName)
	}
	if len(_apigatewayv2AccessLogSettings) > 0 {
		if err := assignInputField(input, "AccessLogSettings", _apigatewayv2AccessLogSettings); err != nil {
			log.Errorf("invalid --access-log-settings: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2AutoDeploy) > 0 {
		if err := assignInputField(input, "AutoDeploy", _apigatewayv2AutoDeploy); err != nil {
			log.Errorf("invalid --auto-deploy: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2ClientCertificateId) > 0 {
		input.ClientCertificateId = aws.String(_apigatewayv2ClientCertificateId)
	}
	if len(_apigatewayv2DefaultRouteSettings) > 0 {
		if err := assignInputField(input, "DefaultRouteSettings", _apigatewayv2DefaultRouteSettings); err != nil {
			log.Errorf("invalid --default-route-settings: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2DeploymentId) > 0 {
		input.DeploymentId = aws.String(_apigatewayv2DeploymentId)
	}
	if len(_apigatewayv2Description) > 0 {
		input.Description = aws.String(_apigatewayv2Description)
	}
	if len(_apigatewayv2RouteSettings) > 0 {
		if err := assignInputField(input, "RouteSettings", _apigatewayv2RouteSettings); err != nil {
			log.Errorf("invalid --route-settings: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2StageVariables) > 0 {
		if err := assignInputField(input, "StageVariables", _apigatewayv2StageVariables); err != nil {
			log.Errorf("invalid --stage-variables: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _apigatewayv2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
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

// Creates a VPC link.
func apigatewayv2_CreateVpcLink(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.CreateVpcLinkInput{
		// Name: *string, // Required
		// SubnetIds: []string, // Required
	}

	if len(_apigatewayv2Name) > 0 {
		input.Name = aws.String(_apigatewayv2Name)
	}
	if len(_apigatewayv2SubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _apigatewayv2SubnetIds...)
	}
	if len(_apigatewayv2SecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _apigatewayv2SecurityGroupIds...)
	}
	if len(_apigatewayv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _apigatewayv2Tags); err != nil {
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

// Deletes the AccessLogSettings for a Stage. To disable access logging for a
// Stage, delete its AccessLogSettings.
func apigatewayv2_DeleteAccessLogSettings(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.DeleteAccessLogSettingsInput{
		// ApiId: *string, // Required
		// StageName: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2StageName) > 0 {
		input.StageName = aws.String(_apigatewayv2StageName)
	}

	if resp, err := client.DeleteAccessLogSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Api resource.
func apigatewayv2_DeleteApi(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.DeleteApiInput{
		// ApiId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}

	if resp, err := client.DeleteApi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an API mapping.
func apigatewayv2_DeleteApiMapping(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.DeleteApiMappingInput{
		// ApiMappingId: *string, // Required
		// DomainName: *string, // Required
	}

	if len(_apigatewayv2ApiMappingId) > 0 {
		input.ApiMappingId = aws.String(_apigatewayv2ApiMappingId)
	}
	if len(_apigatewayv2DomainName) > 0 {
		input.DomainName = aws.String(_apigatewayv2DomainName)
	}

	if resp, err := client.DeleteApiMapping(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Authorizer.
func apigatewayv2_DeleteAuthorizer(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.DeleteAuthorizerInput{
		// ApiId: *string, // Required
		// AuthorizerId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2AuthorizerId) > 0 {
		input.AuthorizerId = aws.String(_apigatewayv2AuthorizerId)
	}

	if resp, err := client.DeleteAuthorizer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a CORS configuration.
func apigatewayv2_DeleteCorsConfiguration(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.DeleteCorsConfigurationInput{
		// ApiId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}

	if resp, err := client.DeleteCorsConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Deployment.
func apigatewayv2_DeleteDeployment(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.DeleteDeploymentInput{
		// ApiId: *string, // Required
		// DeploymentId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2DeploymentId) > 0 {
		input.DeploymentId = aws.String(_apigatewayv2DeploymentId)
	}

	if resp, err := client.DeleteDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a domain name.
func apigatewayv2_DeleteDomainName(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.DeleteDomainNameInput{
		// DomainName: *string, // Required
	}

	if len(_apigatewayv2DomainName) > 0 {
		input.DomainName = aws.String(_apigatewayv2DomainName)
	}

	if resp, err := client.DeleteDomainName(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Integration.
func apigatewayv2_DeleteIntegration(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.DeleteIntegrationInput{
		// ApiId: *string, // Required
		// IntegrationId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2IntegrationId) > 0 {
		input.IntegrationId = aws.String(_apigatewayv2IntegrationId)
	}

	if resp, err := client.DeleteIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an IntegrationResponses.
func apigatewayv2_DeleteIntegrationResponse(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.DeleteIntegrationResponseInput{
		// ApiId: *string, // Required
		// IntegrationId: *string, // Required
		// IntegrationResponseId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2IntegrationId) > 0 {
		input.IntegrationId = aws.String(_apigatewayv2IntegrationId)
	}
	if len(_apigatewayv2IntegrationResponseId) > 0 {
		input.IntegrationResponseId = aws.String(_apigatewayv2IntegrationResponseId)
	}

	if resp, err := client.DeleteIntegrationResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Model.
func apigatewayv2_DeleteModel(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.DeleteModelInput{
		// ApiId: *string, // Required
		// ModelId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2ModelId) > 0 {
		input.ModelId = aws.String(_apigatewayv2ModelId)
	}

	if resp, err := client.DeleteModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a portal.
func apigatewayv2_DeletePortal(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.DeletePortalInput{
		// PortalId: *string, // Required
	}

	if len(_apigatewayv2PortalId) > 0 {
		input.PortalId = aws.String(_apigatewayv2PortalId)
	}

	if resp, err := client.DeletePortal(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a portal product.
func apigatewayv2_DeletePortalProduct(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.DeletePortalProductInput{
		// PortalProductId: *string, // Required
	}

	if len(_apigatewayv2PortalProductId) > 0 {
		input.PortalProductId = aws.String(_apigatewayv2PortalProductId)
	}

	if resp, err := client.DeletePortalProduct(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the sharing policy for a portal product.
func apigatewayv2_DeletePortalProductSharingPolicy(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.DeletePortalProductSharingPolicyInput{
		// PortalProductId: *string, // Required
	}

	if len(_apigatewayv2PortalProductId) > 0 {
		input.PortalProductId = aws.String(_apigatewayv2PortalProductId)
	}

	if resp, err := client.DeletePortalProductSharingPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a product page of a portal product.
func apigatewayv2_DeleteProductPage(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.DeleteProductPageInput{
		// PortalProductId: *string, // Required
		// ProductPageId: *string, // Required
	}

	if len(_apigatewayv2PortalProductId) > 0 {
		input.PortalProductId = aws.String(_apigatewayv2PortalProductId)
	}
	if len(_apigatewayv2ProductPageId) > 0 {
		input.ProductPageId = aws.String(_apigatewayv2ProductPageId)
	}

	if resp, err := client.DeleteProductPage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a product REST endpoint page.
func apigatewayv2_DeleteProductRestEndpointPage(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.DeleteProductRestEndpointPageInput{
		// PortalProductId: *string, // Required
		// ProductRestEndpointPageId: *string, // Required
	}

	if len(_apigatewayv2PortalProductId) > 0 {
		input.PortalProductId = aws.String(_apigatewayv2PortalProductId)
	}
	if len(_apigatewayv2ProductRestEndpointPageId) > 0 {
		input.ProductRestEndpointPageId = aws.String(_apigatewayv2ProductRestEndpointPageId)
	}

	if resp, err := client.DeleteProductRestEndpointPage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Route.
func apigatewayv2_DeleteRoute(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.DeleteRouteInput{
		// ApiId: *string, // Required
		// RouteId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2RouteId) > 0 {
		input.RouteId = aws.String(_apigatewayv2RouteId)
	}

	if resp, err := client.DeleteRoute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a route request parameter. Supported only for WebSocket APIs.
func apigatewayv2_DeleteRouteRequestParameter(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.DeleteRouteRequestParameterInput{
		// ApiId: *string, // Required
		// RequestParameterKey: *string, // Required
		// RouteId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2RequestParameterKey) > 0 {
		input.RequestParameterKey = aws.String(_apigatewayv2RequestParameterKey)
	}
	if len(_apigatewayv2RouteId) > 0 {
		input.RouteId = aws.String(_apigatewayv2RouteId)
	}

	if resp, err := client.DeleteRouteRequestParameter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a RouteResponse.
func apigatewayv2_DeleteRouteResponse(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.DeleteRouteResponseInput{
		// ApiId: *string, // Required
		// RouteId: *string, // Required
		// RouteResponseId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2RouteId) > 0 {
		input.RouteId = aws.String(_apigatewayv2RouteId)
	}
	if len(_apigatewayv2RouteResponseId) > 0 {
		input.RouteResponseId = aws.String(_apigatewayv2RouteResponseId)
	}

	if resp, err := client.DeleteRouteResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the RouteSettings for a stage.
func apigatewayv2_DeleteRouteSettings(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.DeleteRouteSettingsInput{
		// ApiId: *string, // Required
		// RouteKey: *string, // Required
		// StageName: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2RouteKey) > 0 {
		input.RouteKey = aws.String(_apigatewayv2RouteKey)
	}
	if len(_apigatewayv2StageName) > 0 {
		input.StageName = aws.String(_apigatewayv2StageName)
	}

	if resp, err := client.DeleteRouteSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a routing rule.
func apigatewayv2_DeleteRoutingRule(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.DeleteRoutingRuleInput{
		// DomainName: *string, // Required
		// RoutingRuleId: *string, // Required
	}

	if len(_apigatewayv2DomainName) > 0 {
		input.DomainName = aws.String(_apigatewayv2DomainName)
	}
	if len(_apigatewayv2RoutingRuleId) > 0 {
		input.RoutingRuleId = aws.String(_apigatewayv2RoutingRuleId)
	}
	if len(_apigatewayv2DomainNameId) > 0 {
		input.DomainNameId = aws.String(_apigatewayv2DomainNameId)
	}

	if resp, err := client.DeleteRoutingRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Stage.
func apigatewayv2_DeleteStage(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.DeleteStageInput{
		// ApiId: *string, // Required
		// StageName: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2StageName) > 0 {
		input.StageName = aws.String(_apigatewayv2StageName)
	}

	if resp, err := client.DeleteStage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a VPC link.
func apigatewayv2_DeleteVpcLink(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.DeleteVpcLinkInput{
		// VpcLinkId: *string, // Required
	}

	if len(_apigatewayv2VpcLinkId) > 0 {
		input.VpcLinkId = aws.String(_apigatewayv2VpcLinkId)
	}

	if resp, err := client.DeleteVpcLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the publication of a portal portal.
func apigatewayv2_DisablePortal(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.DisablePortalInput{
		// PortalId: *string, // Required
	}

	if len(_apigatewayv2PortalId) > 0 {
		input.PortalId = aws.String(_apigatewayv2PortalId)
	}

	if resp, err := client.DisablePortal(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func apigatewayv2_ExportApi(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.ExportApiInput{
		// ApiId: *string, // Required
		// OutputType: *string, // Required
		// Specification: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2OutputType) > 0 {
		input.OutputType = aws.String(_apigatewayv2OutputType)
	}
	if len(_apigatewayv2Specification) > 0 {
		input.Specification = aws.String(_apigatewayv2Specification)
	}
	if len(_apigatewayv2ExportVersion) > 0 {
		input.ExportVersion = aws.String(_apigatewayv2ExportVersion)
	}
	if len(_apigatewayv2IncludeExtensions) > 0 {
		if err := assignInputField(input, "IncludeExtensions", _apigatewayv2IncludeExtensions); err != nil {
			log.Errorf("invalid --include-extensions: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2StageName) > 0 {
		input.StageName = aws.String(_apigatewayv2StageName)
	}

	if resp, err := client.ExportApi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an Api resource.
func apigatewayv2_GetApi(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetApiInput{
		// ApiId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}

	if resp, err := client.GetApi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an API mapping.
func apigatewayv2_GetApiMapping(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetApiMappingInput{
		// ApiMappingId: *string, // Required
		// DomainName: *string, // Required
	}

	if len(_apigatewayv2ApiMappingId) > 0 {
		input.ApiMappingId = aws.String(_apigatewayv2ApiMappingId)
	}
	if len(_apigatewayv2DomainName) > 0 {
		input.DomainName = aws.String(_apigatewayv2DomainName)
	}

	if resp, err := client.GetApiMapping(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets API mappings.
func apigatewayv2_GetApiMappings(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetApiMappingsInput{
		// DomainName: *string, // Required
	}

	if len(_apigatewayv2DomainName) > 0 {
		input.DomainName = aws.String(_apigatewayv2DomainName)
	}
	if len(_apigatewayv2MaxResults) > 0 {
		input.MaxResults = aws.String(_apigatewayv2MaxResults)
	}
	if len(_apigatewayv2NextToken) > 0 {
		input.NextToken = aws.String(_apigatewayv2NextToken)
	}

	if resp, err := client.GetApiMappings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a collection of Api resources.
func apigatewayv2_GetApis(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetApisInput{}

	if len(_apigatewayv2MaxResults) > 0 {
		input.MaxResults = aws.String(_apigatewayv2MaxResults)
	}
	if len(_apigatewayv2NextToken) > 0 {
		input.NextToken = aws.String(_apigatewayv2NextToken)
	}

	if resp, err := client.GetApis(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an Authorizer.
func apigatewayv2_GetAuthorizer(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetAuthorizerInput{
		// ApiId: *string, // Required
		// AuthorizerId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2AuthorizerId) > 0 {
		input.AuthorizerId = aws.String(_apigatewayv2AuthorizerId)
	}

	if resp, err := client.GetAuthorizer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the Authorizers for an API.
func apigatewayv2_GetAuthorizers(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetAuthorizersInput{
		// ApiId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2MaxResults) > 0 {
		input.MaxResults = aws.String(_apigatewayv2MaxResults)
	}
	if len(_apigatewayv2NextToken) > 0 {
		input.NextToken = aws.String(_apigatewayv2NextToken)
	}

	if resp, err := client.GetAuthorizers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a Deployment.
func apigatewayv2_GetDeployment(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetDeploymentInput{
		// ApiId: *string, // Required
		// DeploymentId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2DeploymentId) > 0 {
		input.DeploymentId = aws.String(_apigatewayv2DeploymentId)
	}

	if resp, err := client.GetDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the Deployments for an API.
func apigatewayv2_GetDeployments(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetDeploymentsInput{
		// ApiId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2MaxResults) > 0 {
		input.MaxResults = aws.String(_apigatewayv2MaxResults)
	}
	if len(_apigatewayv2NextToken) > 0 {
		input.NextToken = aws.String(_apigatewayv2NextToken)
	}

	if resp, err := client.GetDeployments(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a domain name.
func apigatewayv2_GetDomainName(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetDomainNameInput{
		// DomainName: *string, // Required
	}

	if len(_apigatewayv2DomainName) > 0 {
		input.DomainName = aws.String(_apigatewayv2DomainName)
	}

	if resp, err := client.GetDomainName(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the domain names for an AWS account.
func apigatewayv2_GetDomainNames(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetDomainNamesInput{}

	if len(_apigatewayv2MaxResults) > 0 {
		input.MaxResults = aws.String(_apigatewayv2MaxResults)
	}
	if len(_apigatewayv2NextToken) > 0 {
		input.NextToken = aws.String(_apigatewayv2NextToken)
	}

	if resp, err := client.GetDomainNames(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an Integration.
func apigatewayv2_GetIntegration(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetIntegrationInput{
		// ApiId: *string, // Required
		// IntegrationId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2IntegrationId) > 0 {
		input.IntegrationId = aws.String(_apigatewayv2IntegrationId)
	}

	if resp, err := client.GetIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an IntegrationResponses.
func apigatewayv2_GetIntegrationResponse(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetIntegrationResponseInput{
		// ApiId: *string, // Required
		// IntegrationId: *string, // Required
		// IntegrationResponseId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2IntegrationId) > 0 {
		input.IntegrationId = aws.String(_apigatewayv2IntegrationId)
	}
	if len(_apigatewayv2IntegrationResponseId) > 0 {
		input.IntegrationResponseId = aws.String(_apigatewayv2IntegrationResponseId)
	}

	if resp, err := client.GetIntegrationResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the IntegrationResponses for an Integration.
func apigatewayv2_GetIntegrationResponses(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetIntegrationResponsesInput{
		// ApiId: *string, // Required
		// IntegrationId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2IntegrationId) > 0 {
		input.IntegrationId = aws.String(_apigatewayv2IntegrationId)
	}
	if len(_apigatewayv2MaxResults) > 0 {
		input.MaxResults = aws.String(_apigatewayv2MaxResults)
	}
	if len(_apigatewayv2NextToken) > 0 {
		input.NextToken = aws.String(_apigatewayv2NextToken)
	}

	if resp, err := client.GetIntegrationResponses(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the Integrations for an API.
func apigatewayv2_GetIntegrations(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetIntegrationsInput{
		// ApiId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2MaxResults) > 0 {
		input.MaxResults = aws.String(_apigatewayv2MaxResults)
	}
	if len(_apigatewayv2NextToken) > 0 {
		input.NextToken = aws.String(_apigatewayv2NextToken)
	}

	if resp, err := client.GetIntegrations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a Model.
func apigatewayv2_GetModel(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetModelInput{
		// ApiId: *string, // Required
		// ModelId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2ModelId) > 0 {
		input.ModelId = aws.String(_apigatewayv2ModelId)
	}

	if resp, err := client.GetModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a model template.
func apigatewayv2_GetModelTemplate(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetModelTemplateInput{
		// ApiId: *string, // Required
		// ModelId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2ModelId) > 0 {
		input.ModelId = aws.String(_apigatewayv2ModelId)
	}

	if resp, err := client.GetModelTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the Models for an API.
func apigatewayv2_GetModels(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetModelsInput{
		// ApiId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2MaxResults) > 0 {
		input.MaxResults = aws.String(_apigatewayv2MaxResults)
	}
	if len(_apigatewayv2NextToken) > 0 {
		input.NextToken = aws.String(_apigatewayv2NextToken)
	}

	if resp, err := client.GetModels(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a portal.
func apigatewayv2_GetPortal(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetPortalInput{
		// PortalId: *string, // Required
	}

	if len(_apigatewayv2PortalId) > 0 {
		input.PortalId = aws.String(_apigatewayv2PortalId)
	}

	if resp, err := client.GetPortal(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a portal product.
func apigatewayv2_GetPortalProduct(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetPortalProductInput{
		// PortalProductId: *string, // Required
	}

	if len(_apigatewayv2PortalProductId) > 0 {
		input.PortalProductId = aws.String(_apigatewayv2PortalProductId)
	}
	if len(_apigatewayv2ResourceOwnerAccountId) > 0 {
		input.ResourceOwnerAccountId = aws.String(_apigatewayv2ResourceOwnerAccountId)
	}

	if resp, err := client.GetPortalProduct(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the sharing policy for a portal product.
func apigatewayv2_GetPortalProductSharingPolicy(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetPortalProductSharingPolicyInput{
		// PortalProductId: *string, // Required
	}

	if len(_apigatewayv2PortalProductId) > 0 {
		input.PortalProductId = aws.String(_apigatewayv2PortalProductId)
	}

	if resp, err := client.GetPortalProductSharingPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a product page of a portal product.
func apigatewayv2_GetProductPage(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetProductPageInput{
		// PortalProductId: *string, // Required
		// ProductPageId: *string, // Required
	}

	if len(_apigatewayv2PortalProductId) > 0 {
		input.PortalProductId = aws.String(_apigatewayv2PortalProductId)
	}
	if len(_apigatewayv2ProductPageId) > 0 {
		input.ProductPageId = aws.String(_apigatewayv2ProductPageId)
	}
	if len(_apigatewayv2ResourceOwnerAccountId) > 0 {
		input.ResourceOwnerAccountId = aws.String(_apigatewayv2ResourceOwnerAccountId)
	}

	if resp, err := client.GetProductPage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a product REST endpoint page.
func apigatewayv2_GetProductRestEndpointPage(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetProductRestEndpointPageInput{
		// PortalProductId: *string, // Required
		// ProductRestEndpointPageId: *string, // Required
	}

	if len(_apigatewayv2PortalProductId) > 0 {
		input.PortalProductId = aws.String(_apigatewayv2PortalProductId)
	}
	if len(_apigatewayv2ProductRestEndpointPageId) > 0 {
		input.ProductRestEndpointPageId = aws.String(_apigatewayv2ProductRestEndpointPageId)
	}
	if len(_apigatewayv2IncludeRawDisplayContent) > 0 {
		input.IncludeRawDisplayContent = aws.String(_apigatewayv2IncludeRawDisplayContent)
	}
	if len(_apigatewayv2ResourceOwnerAccountId) > 0 {
		input.ResourceOwnerAccountId = aws.String(_apigatewayv2ResourceOwnerAccountId)
	}

	if resp, err := client.GetProductRestEndpointPage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a Route.
func apigatewayv2_GetRoute(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetRouteInput{
		// ApiId: *string, // Required
		// RouteId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2RouteId) > 0 {
		input.RouteId = aws.String(_apigatewayv2RouteId)
	}

	if resp, err := client.GetRoute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a RouteResponse.
func apigatewayv2_GetRouteResponse(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetRouteResponseInput{
		// ApiId: *string, // Required
		// RouteId: *string, // Required
		// RouteResponseId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2RouteId) > 0 {
		input.RouteId = aws.String(_apigatewayv2RouteId)
	}
	if len(_apigatewayv2RouteResponseId) > 0 {
		input.RouteResponseId = aws.String(_apigatewayv2RouteResponseId)
	}

	if resp, err := client.GetRouteResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the RouteResponses for a Route.
func apigatewayv2_GetRouteResponses(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetRouteResponsesInput{
		// ApiId: *string, // Required
		// RouteId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2RouteId) > 0 {
		input.RouteId = aws.String(_apigatewayv2RouteId)
	}
	if len(_apigatewayv2MaxResults) > 0 {
		input.MaxResults = aws.String(_apigatewayv2MaxResults)
	}
	if len(_apigatewayv2NextToken) > 0 {
		input.NextToken = aws.String(_apigatewayv2NextToken)
	}

	if resp, err := client.GetRouteResponses(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the Routes for an API.
func apigatewayv2_GetRoutes(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetRoutesInput{
		// ApiId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2MaxResults) > 0 {
		input.MaxResults = aws.String(_apigatewayv2MaxResults)
	}
	if len(_apigatewayv2NextToken) > 0 {
		input.NextToken = aws.String(_apigatewayv2NextToken)
	}

	if resp, err := client.GetRoutes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a routing rule.
func apigatewayv2_GetRoutingRule(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetRoutingRuleInput{
		// DomainName: *string, // Required
		// RoutingRuleId: *string, // Required
	}

	if len(_apigatewayv2DomainName) > 0 {
		input.DomainName = aws.String(_apigatewayv2DomainName)
	}
	if len(_apigatewayv2RoutingRuleId) > 0 {
		input.RoutingRuleId = aws.String(_apigatewayv2RoutingRuleId)
	}
	if len(_apigatewayv2DomainNameId) > 0 {
		input.DomainNameId = aws.String(_apigatewayv2DomainNameId)
	}

	if resp, err := client.GetRoutingRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a Stage.
func apigatewayv2_GetStage(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetStageInput{
		// ApiId: *string, // Required
		// StageName: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2StageName) > 0 {
		input.StageName = aws.String(_apigatewayv2StageName)
	}

	if resp, err := client.GetStage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the Stages for an API.
func apigatewayv2_GetStages(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetStagesInput{
		// ApiId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2MaxResults) > 0 {
		input.MaxResults = aws.String(_apigatewayv2MaxResults)
	}
	if len(_apigatewayv2NextToken) > 0 {
		input.NextToken = aws.String(_apigatewayv2NextToken)
	}

	if resp, err := client.GetStages(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a collection of Tag resources.
func apigatewayv2_GetTags(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetTagsInput{
		// ResourceArn: *string, // Required
	}

	if len(_apigatewayv2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_apigatewayv2ResourceArn)
	}

	if resp, err := client.GetTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a VPC link.
func apigatewayv2_GetVpcLink(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetVpcLinkInput{
		// VpcLinkId: *string, // Required
	}

	if len(_apigatewayv2VpcLinkId) > 0 {
		input.VpcLinkId = aws.String(_apigatewayv2VpcLinkId)
	}

	if resp, err := client.GetVpcLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a collection of VPC links.
func apigatewayv2_GetVpcLinks(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.GetVpcLinksInput{}

	if len(_apigatewayv2MaxResults) > 0 {
		input.MaxResults = aws.String(_apigatewayv2MaxResults)
	}
	if len(_apigatewayv2NextToken) > 0 {
		input.NextToken = aws.String(_apigatewayv2NextToken)
	}

	if resp, err := client.GetVpcLinks(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports an API.
func apigatewayv2_ImportApi(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.ImportApiInput{
		// Body: *string, // Required
	}

	if len(_apigatewayv2Body) > 0 {
		input.Body = aws.String(_apigatewayv2Body)
	}
	if len(_apigatewayv2Basepath) > 0 {
		input.Basepath = aws.String(_apigatewayv2Basepath)
	}
	if len(_apigatewayv2FailOnWarnings) > 0 {
		if err := assignInputField(input, "FailOnWarnings", _apigatewayv2FailOnWarnings); err != nil {
			log.Errorf("invalid --fail-on-warnings: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportApi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists portal products.
func apigatewayv2_ListPortalProducts(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.ListPortalProductsInput{}

	if len(_apigatewayv2MaxResults) > 0 {
		input.MaxResults = aws.String(_apigatewayv2MaxResults)
	}
	if len(_apigatewayv2NextToken) > 0 {
		input.NextToken = aws.String(_apigatewayv2NextToken)
	}
	if len(_apigatewayv2ResourceOwner) > 0 {
		input.ResourceOwner = aws.String(_apigatewayv2ResourceOwner)
	}

	if resp, err := client.ListPortalProducts(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists portals.
func apigatewayv2_ListPortals(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.ListPortalsInput{}

	if len(_apigatewayv2MaxResults) > 0 {
		input.MaxResults = aws.String(_apigatewayv2MaxResults)
	}
	if len(_apigatewayv2NextToken) > 0 {
		input.NextToken = aws.String(_apigatewayv2NextToken)
	}

	if resp, err := client.ListPortals(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the product pages for a portal product.
func apigatewayv2_ListProductPages(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.ListProductPagesInput{
		// PortalProductId: *string, // Required
	}

	if len(_apigatewayv2PortalProductId) > 0 {
		input.PortalProductId = aws.String(_apigatewayv2PortalProductId)
	}
	if len(_apigatewayv2MaxResults) > 0 {
		input.MaxResults = aws.String(_apigatewayv2MaxResults)
	}
	if len(_apigatewayv2NextToken) > 0 {
		input.NextToken = aws.String(_apigatewayv2NextToken)
	}
	if len(_apigatewayv2ResourceOwnerAccountId) > 0 {
		input.ResourceOwnerAccountId = aws.String(_apigatewayv2ResourceOwnerAccountId)
	}

	if resp, err := client.ListProductPages(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the product REST endpoint pages of a portal product.
func apigatewayv2_ListProductRestEndpointPages(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.ListProductRestEndpointPagesInput{
		// PortalProductId: *string, // Required
	}

	if len(_apigatewayv2PortalProductId) > 0 {
		input.PortalProductId = aws.String(_apigatewayv2PortalProductId)
	}
	if len(_apigatewayv2MaxResults) > 0 {
		input.MaxResults = aws.String(_apigatewayv2MaxResults)
	}
	if len(_apigatewayv2NextToken) > 0 {
		input.NextToken = aws.String(_apigatewayv2NextToken)
	}
	if len(_apigatewayv2ResourceOwnerAccountId) > 0 {
		input.ResourceOwnerAccountId = aws.String(_apigatewayv2ResourceOwnerAccountId)
	}

	if resp, err := client.ListProductRestEndpointPages(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists routing rules.
func apigatewayv2_ListRoutingRules(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.ListRoutingRulesInput{
		// DomainName: *string, // Required
	}

	if len(_apigatewayv2DomainName) > 0 {
		input.DomainName = aws.String(_apigatewayv2DomainName)
	}
	if len(_apigatewayv2DomainNameId) > 0 {
		input.DomainNameId = aws.String(_apigatewayv2DomainNameId)
	}
	if len(_apigatewayv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _apigatewayv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2NextToken) > 0 {
		input.NextToken = aws.String(_apigatewayv2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRoutingRules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*apigatewayv2.ListRoutingRulesOutput
	p := apigatewayv2.NewListRoutingRulesPaginator(client, input)
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

// Creates a portal preview.
func apigatewayv2_PreviewPortal(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.PreviewPortalInput{
		// PortalId: *string, // Required
	}

	if len(_apigatewayv2PortalId) > 0 {
		input.PortalId = aws.String(_apigatewayv2PortalId)
	}

	if resp, err := client.PreviewPortal(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Publishes a portal.
func apigatewayv2_PublishPortal(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.PublishPortalInput{
		// PortalId: *string, // Required
	}

	if len(_apigatewayv2PortalId) > 0 {
		input.PortalId = aws.String(_apigatewayv2PortalId)
	}
	if len(_apigatewayv2Description) > 0 {
		input.Description = aws.String(_apigatewayv2Description)
	}

	if resp, err := client.PublishPortal(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the sharing policy for a portal product.
func apigatewayv2_PutPortalProductSharingPolicy(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.PutPortalProductSharingPolicyInput{
		// PolicyDocument: *string, // Required
		// PortalProductId: *string, // Required
	}

	if len(_apigatewayv2PolicyDocument) > 0 {
		input.PolicyDocument = aws.String(_apigatewayv2PolicyDocument)
	}
	if len(_apigatewayv2PortalProductId) > 0 {
		input.PortalProductId = aws.String(_apigatewayv2PortalProductId)
	}

	if resp, err := client.PutPortalProductSharingPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a routing rule.
func apigatewayv2_PutRoutingRule(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.PutRoutingRuleInput{
		// Actions: []types.RoutingRuleAction, // Required
		// Conditions: []types.RoutingRuleCondition, // Required
		// DomainName: *string, // Required
		// Priority: *int32, // Required
		// RoutingRuleId: *string, // Required
	}

	if len(_apigatewayv2Actions) > 0 {
		if err := assignInputField(input, "Actions", _apigatewayv2Actions); err != nil {
			log.Errorf("invalid --actions: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2Conditions) > 0 {
		if err := assignInputField(input, "Conditions", _apigatewayv2Conditions); err != nil {
			log.Errorf("invalid --conditions: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2DomainName) > 0 {
		input.DomainName = aws.String(_apigatewayv2DomainName)
	}
	if len(_apigatewayv2Priority) > 0 {
		if err := assignInputField(input, "Priority", _apigatewayv2Priority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2RoutingRuleId) > 0 {
		input.RoutingRuleId = aws.String(_apigatewayv2RoutingRuleId)
	}
	if len(_apigatewayv2DomainNameId) > 0 {
		input.DomainNameId = aws.String(_apigatewayv2DomainNameId)
	}

	if resp, err := client.PutRoutingRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Puts an Api resource.
func apigatewayv2_ReimportApi(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.ReimportApiInput{
		// ApiId: *string, // Required
		// Body: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2Body) > 0 {
		input.Body = aws.String(_apigatewayv2Body)
	}
	if len(_apigatewayv2Basepath) > 0 {
		input.Basepath = aws.String(_apigatewayv2Basepath)
	}
	if len(_apigatewayv2FailOnWarnings) > 0 {
		if err := assignInputField(input, "FailOnWarnings", _apigatewayv2FailOnWarnings); err != nil {
			log.Errorf("invalid --fail-on-warnings: %s", err.Error())
			return
		}
	}

	if resp, err := client.ReimportApi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Resets all authorizer cache entries on a stage. Supported only for HTTP APIs.
func apigatewayv2_ResetAuthorizersCache(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.ResetAuthorizersCacheInput{
		// ApiId: *string, // Required
		// StageName: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2StageName) > 0 {
		input.StageName = aws.String(_apigatewayv2StageName)
	}

	if resp, err := client.ResetAuthorizersCache(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Tag resource to represent a tag.
func apigatewayv2_TagResource(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.TagResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_apigatewayv2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_apigatewayv2ResourceArn)
	}
	if len(_apigatewayv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _apigatewayv2Tags); err != nil {
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

// Deletes a Tag.
func apigatewayv2_UntagResource(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_apigatewayv2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_apigatewayv2ResourceArn)
	}
	if len(_apigatewayv2TagKeys) > 0 {
		input.TagKeys = append([]string(nil), _apigatewayv2TagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an Api resource.
func apigatewayv2_UpdateApi(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.UpdateApiInput{
		// ApiId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2ApiKeySelectionExpression) > 0 {
		input.ApiKeySelectionExpression = aws.String(_apigatewayv2ApiKeySelectionExpression)
	}
	if len(_apigatewayv2CorsConfiguration) > 0 {
		if err := assignInputField(input, "CorsConfiguration", _apigatewayv2CorsConfiguration); err != nil {
			log.Errorf("invalid --cors-configuration: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2CredentialsArn) > 0 {
		input.CredentialsArn = aws.String(_apigatewayv2CredentialsArn)
	}
	if len(_apigatewayv2Description) > 0 {
		input.Description = aws.String(_apigatewayv2Description)
	}
	if len(_apigatewayv2DisableExecuteApiEndpoint) > 0 {
		if err := assignInputField(input, "DisableExecuteApiEndpoint", _apigatewayv2DisableExecuteApiEndpoint); err != nil {
			log.Errorf("invalid --disable-execute-api-endpoint: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2DisableSchemaValidation) > 0 {
		if err := assignInputField(input, "DisableSchemaValidation", _apigatewayv2DisableSchemaValidation); err != nil {
			log.Errorf("invalid --disable-schema-validation: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2IpAddressType) > 0 {
		if err := assignInputField(input, "IpAddressType", _apigatewayv2IpAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2Name) > 0 {
		input.Name = aws.String(_apigatewayv2Name)
	}
	if len(_apigatewayv2RouteKey) > 0 {
		input.RouteKey = aws.String(_apigatewayv2RouteKey)
	}
	if len(_apigatewayv2RouteSelectionExpression) > 0 {
		input.RouteSelectionExpression = aws.String(_apigatewayv2RouteSelectionExpression)
	}
	if len(_apigatewayv2Target) > 0 {
		input.Target = aws.String(_apigatewayv2Target)
	}
	if len(_apigatewayv2Version) > 0 {
		input.Version = aws.String(_apigatewayv2Version)
	}

	if resp, err := client.UpdateApi(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The API mapping.
func apigatewayv2_UpdateApiMapping(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.UpdateApiMappingInput{
		// ApiId: *string, // Required
		// ApiMappingId: *string, // Required
		// DomainName: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2ApiMappingId) > 0 {
		input.ApiMappingId = aws.String(_apigatewayv2ApiMappingId)
	}
	if len(_apigatewayv2DomainName) > 0 {
		input.DomainName = aws.String(_apigatewayv2DomainName)
	}
	if len(_apigatewayv2ApiMappingKey) > 0 {
		input.ApiMappingKey = aws.String(_apigatewayv2ApiMappingKey)
	}
	if len(_apigatewayv2Stage) > 0 {
		input.Stage = aws.String(_apigatewayv2Stage)
	}

	if resp, err := client.UpdateApiMapping(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an Authorizer.
func apigatewayv2_UpdateAuthorizer(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.UpdateAuthorizerInput{
		// ApiId: *string, // Required
		// AuthorizerId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2AuthorizerId) > 0 {
		input.AuthorizerId = aws.String(_apigatewayv2AuthorizerId)
	}
	if len(_apigatewayv2AuthorizerCredentialsArn) > 0 {
		input.AuthorizerCredentialsArn = aws.String(_apigatewayv2AuthorizerCredentialsArn)
	}
	if len(_apigatewayv2AuthorizerPayloadFormatVersion) > 0 {
		input.AuthorizerPayloadFormatVersion = aws.String(_apigatewayv2AuthorizerPayloadFormatVersion)
	}
	if len(_apigatewayv2AuthorizerResultTtlInSeconds) > 0 {
		if err := assignInputField(input, "AuthorizerResultTtlInSeconds", _apigatewayv2AuthorizerResultTtlInSeconds); err != nil {
			log.Errorf("invalid --authorizer-result-ttl-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2AuthorizerType) > 0 {
		if err := assignInputField(input, "AuthorizerType", _apigatewayv2AuthorizerType); err != nil {
			log.Errorf("invalid --authorizer-type: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2AuthorizerUri) > 0 {
		input.AuthorizerUri = aws.String(_apigatewayv2AuthorizerUri)
	}
	if len(_apigatewayv2EnableSimpleResponses) > 0 {
		if err := assignInputField(input, "EnableSimpleResponses", _apigatewayv2EnableSimpleResponses); err != nil {
			log.Errorf("invalid --enable-simple-responses: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2IdentitySource) > 0 {
		input.IdentitySource = append([]string(nil), _apigatewayv2IdentitySource...)
	}
	if len(_apigatewayv2IdentityValidationExpression) > 0 {
		input.IdentityValidationExpression = aws.String(_apigatewayv2IdentityValidationExpression)
	}
	if len(_apigatewayv2JwtConfiguration) > 0 {
		if err := assignInputField(input, "JwtConfiguration", _apigatewayv2JwtConfiguration); err != nil {
			log.Errorf("invalid --jwt-configuration: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2Name) > 0 {
		input.Name = aws.String(_apigatewayv2Name)
	}

	if resp, err := client.UpdateAuthorizer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Deployment.
func apigatewayv2_UpdateDeployment(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.UpdateDeploymentInput{
		// ApiId: *string, // Required
		// DeploymentId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2DeploymentId) > 0 {
		input.DeploymentId = aws.String(_apigatewayv2DeploymentId)
	}
	if len(_apigatewayv2Description) > 0 {
		input.Description = aws.String(_apigatewayv2Description)
	}

	if resp, err := client.UpdateDeployment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a domain name.
func apigatewayv2_UpdateDomainName(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.UpdateDomainNameInput{
		// DomainName: *string, // Required
	}

	if len(_apigatewayv2DomainName) > 0 {
		input.DomainName = aws.String(_apigatewayv2DomainName)
	}
	if len(_apigatewayv2DomainNameConfigurations) > 0 {
		if err := assignInputField(input, "DomainNameConfigurations", _apigatewayv2DomainNameConfigurations); err != nil {
			log.Errorf("invalid --domain-name-configurations: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2MutualTlsAuthentication) > 0 {
		if err := assignInputField(input, "MutualTlsAuthentication", _apigatewayv2MutualTlsAuthentication); err != nil {
			log.Errorf("invalid --mutual-tls-authentication: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2RoutingMode) > 0 {
		if err := assignInputField(input, "RoutingMode", _apigatewayv2RoutingMode); err != nil {
			log.Errorf("invalid --routing-mode: %s", err.Error())
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

// Updates an Integration.
func apigatewayv2_UpdateIntegration(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.UpdateIntegrationInput{
		// ApiId: *string, // Required
		// IntegrationId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2IntegrationId) > 0 {
		input.IntegrationId = aws.String(_apigatewayv2IntegrationId)
	}
	if len(_apigatewayv2ConnectionId) > 0 {
		input.ConnectionId = aws.String(_apigatewayv2ConnectionId)
	}
	if len(_apigatewayv2ConnectionType) > 0 {
		if err := assignInputField(input, "ConnectionType", _apigatewayv2ConnectionType); err != nil {
			log.Errorf("invalid --connection-type: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2ContentHandlingStrategy) > 0 {
		if err := assignInputField(input, "ContentHandlingStrategy", _apigatewayv2ContentHandlingStrategy); err != nil {
			log.Errorf("invalid --content-handling-strategy: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2CredentialsArn) > 0 {
		input.CredentialsArn = aws.String(_apigatewayv2CredentialsArn)
	}
	if len(_apigatewayv2Description) > 0 {
		input.Description = aws.String(_apigatewayv2Description)
	}
	if len(_apigatewayv2IntegrationMethod) > 0 {
		input.IntegrationMethod = aws.String(_apigatewayv2IntegrationMethod)
	}
	if len(_apigatewayv2IntegrationSubtype) > 0 {
		input.IntegrationSubtype = aws.String(_apigatewayv2IntegrationSubtype)
	}
	if len(_apigatewayv2IntegrationType) > 0 {
		if err := assignInputField(input, "IntegrationType", _apigatewayv2IntegrationType); err != nil {
			log.Errorf("invalid --integration-type: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2IntegrationUri) > 0 {
		input.IntegrationUri = aws.String(_apigatewayv2IntegrationUri)
	}
	if len(_apigatewayv2PassthroughBehavior) > 0 {
		if err := assignInputField(input, "PassthroughBehavior", _apigatewayv2PassthroughBehavior); err != nil {
			log.Errorf("invalid --passthrough-behavior: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2PayloadFormatVersion) > 0 {
		input.PayloadFormatVersion = aws.String(_apigatewayv2PayloadFormatVersion)
	}
	if len(_apigatewayv2RequestParameters) > 0 {
		if err := assignInputField(input, "RequestParameters", _apigatewayv2RequestParameters); err != nil {
			log.Errorf("invalid --request-parameters: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2RequestTemplates) > 0 {
		if err := assignInputField(input, "RequestTemplates", _apigatewayv2RequestTemplates); err != nil {
			log.Errorf("invalid --request-templates: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2ResponseParameters) > 0 {
		if err := assignInputField(input, "ResponseParameters", _apigatewayv2ResponseParameters); err != nil {
			log.Errorf("invalid --response-parameters: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2TemplateSelectionExpression) > 0 {
		input.TemplateSelectionExpression = aws.String(_apigatewayv2TemplateSelectionExpression)
	}
	if len(_apigatewayv2TimeoutInMillis) > 0 {
		if err := assignInputField(input, "TimeoutInMillis", _apigatewayv2TimeoutInMillis); err != nil {
			log.Errorf("invalid --timeout-in-millis: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2TlsConfig) > 0 {
		if err := assignInputField(input, "TlsConfig", _apigatewayv2TlsConfig); err != nil {
			log.Errorf("invalid --tls-config: %s", err.Error())
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

// Updates an IntegrationResponses.
func apigatewayv2_UpdateIntegrationResponse(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.UpdateIntegrationResponseInput{
		// ApiId: *string, // Required
		// IntegrationId: *string, // Required
		// IntegrationResponseId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2IntegrationId) > 0 {
		input.IntegrationId = aws.String(_apigatewayv2IntegrationId)
	}
	if len(_apigatewayv2IntegrationResponseId) > 0 {
		input.IntegrationResponseId = aws.String(_apigatewayv2IntegrationResponseId)
	}
	if len(_apigatewayv2ContentHandlingStrategy) > 0 {
		if err := assignInputField(input, "ContentHandlingStrategy", _apigatewayv2ContentHandlingStrategy); err != nil {
			log.Errorf("invalid --content-handling-strategy: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2IntegrationResponseKey) > 0 {
		input.IntegrationResponseKey = aws.String(_apigatewayv2IntegrationResponseKey)
	}
	if len(_apigatewayv2ResponseParameters) > 0 {
		if err := assignInputField(input, "ResponseParameters", _apigatewayv2ResponseParameters); err != nil {
			log.Errorf("invalid --response-parameters: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2ResponseTemplates) > 0 {
		if err := assignInputField(input, "ResponseTemplates", _apigatewayv2ResponseTemplates); err != nil {
			log.Errorf("invalid --response-templates: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2TemplateSelectionExpression) > 0 {
		input.TemplateSelectionExpression = aws.String(_apigatewayv2TemplateSelectionExpression)
	}

	if resp, err := client.UpdateIntegrationResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Model.
func apigatewayv2_UpdateModel(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.UpdateModelInput{
		// ApiId: *string, // Required
		// ModelId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2ModelId) > 0 {
		input.ModelId = aws.String(_apigatewayv2ModelId)
	}
	if len(_apigatewayv2ContentType) > 0 {
		input.ContentType = aws.String(_apigatewayv2ContentType)
	}
	if len(_apigatewayv2Description) > 0 {
		input.Description = aws.String(_apigatewayv2Description)
	}
	if len(_apigatewayv2Name) > 0 {
		input.Name = aws.String(_apigatewayv2Name)
	}
	if len(_apigatewayv2Schema) > 0 {
		input.Schema = aws.String(_apigatewayv2Schema)
	}

	if resp, err := client.UpdateModel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a portal.
func apigatewayv2_UpdatePortal(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.UpdatePortalInput{
		// PortalId: *string, // Required
	}

	if len(_apigatewayv2PortalId) > 0 {
		input.PortalId = aws.String(_apigatewayv2PortalId)
	}
	if len(_apigatewayv2Authorization) > 0 {
		if err := assignInputField(input, "Authorization", _apigatewayv2Authorization); err != nil {
			log.Errorf("invalid --authorization: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2EndpointConfiguration) > 0 {
		if err := assignInputField(input, "EndpointConfiguration", _apigatewayv2EndpointConfiguration); err != nil {
			log.Errorf("invalid --endpoint-configuration: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2IncludedPortalProductArns) > 0 {
		input.IncludedPortalProductArns = append([]string(nil), _apigatewayv2IncludedPortalProductArns...)
	}
	if len(_apigatewayv2LogoUri) > 0 {
		input.LogoUri = aws.String(_apigatewayv2LogoUri)
	}
	if len(_apigatewayv2PortalContent) > 0 {
		if err := assignInputField(input, "PortalContent", _apigatewayv2PortalContent); err != nil {
			log.Errorf("invalid --portal-content: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2RumAppMonitorName) > 0 {
		input.RumAppMonitorName = aws.String(_apigatewayv2RumAppMonitorName)
	}

	if resp, err := client.UpdatePortal(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the portal product.
func apigatewayv2_UpdatePortalProduct(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.UpdatePortalProductInput{
		// PortalProductId: *string, // Required
	}

	if len(_apigatewayv2PortalProductId) > 0 {
		input.PortalProductId = aws.String(_apigatewayv2PortalProductId)
	}
	if len(_apigatewayv2Description) > 0 {
		input.Description = aws.String(_apigatewayv2Description)
	}
	if len(_apigatewayv2DisplayName) > 0 {
		input.DisplayName = aws.String(_apigatewayv2DisplayName)
	}
	if len(_apigatewayv2DisplayOrder) > 0 {
		if err := assignInputField(input, "DisplayOrder", _apigatewayv2DisplayOrder); err != nil {
			log.Errorf("invalid --display-order: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePortalProduct(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a product page of a portal product.
func apigatewayv2_UpdateProductPage(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.UpdateProductPageInput{
		// PortalProductId: *string, // Required
		// ProductPageId: *string, // Required
	}

	if len(_apigatewayv2PortalProductId) > 0 {
		input.PortalProductId = aws.String(_apigatewayv2PortalProductId)
	}
	if len(_apigatewayv2ProductPageId) > 0 {
		input.ProductPageId = aws.String(_apigatewayv2ProductPageId)
	}
	if len(_apigatewayv2DisplayContent) > 0 {
		if err := assignInputField(input, "DisplayContent", _apigatewayv2DisplayContent); err != nil {
			log.Errorf("invalid --display-content: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateProductPage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a product REST endpoint page.
func apigatewayv2_UpdateProductRestEndpointPage(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.UpdateProductRestEndpointPageInput{
		// PortalProductId: *string, // Required
		// ProductRestEndpointPageId: *string, // Required
	}

	if len(_apigatewayv2PortalProductId) > 0 {
		input.PortalProductId = aws.String(_apigatewayv2PortalProductId)
	}
	if len(_apigatewayv2ProductRestEndpointPageId) > 0 {
		input.ProductRestEndpointPageId = aws.String(_apigatewayv2ProductRestEndpointPageId)
	}
	if len(_apigatewayv2DisplayContent) > 0 {
		if err := assignInputField(input, "DisplayContent", _apigatewayv2DisplayContent); err != nil {
			log.Errorf("invalid --display-content: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2TryItState) > 0 {
		if err := assignInputField(input, "TryItState", _apigatewayv2TryItState); err != nil {
			log.Errorf("invalid --try-it-state: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateProductRestEndpointPage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Route.
func apigatewayv2_UpdateRoute(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.UpdateRouteInput{
		// ApiId: *string, // Required
		// RouteId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2RouteId) > 0 {
		input.RouteId = aws.String(_apigatewayv2RouteId)
	}
	if len(_apigatewayv2ApiKeyRequired) > 0 {
		if err := assignInputField(input, "ApiKeyRequired", _apigatewayv2ApiKeyRequired); err != nil {
			log.Errorf("invalid --api-key-required: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2AuthorizationScopes) > 0 {
		input.AuthorizationScopes = append([]string(nil), _apigatewayv2AuthorizationScopes...)
	}
	if len(_apigatewayv2AuthorizationType) > 0 {
		if err := assignInputField(input, "AuthorizationType", _apigatewayv2AuthorizationType); err != nil {
			log.Errorf("invalid --authorization-type: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2AuthorizerId) > 0 {
		input.AuthorizerId = aws.String(_apigatewayv2AuthorizerId)
	}
	if len(_apigatewayv2ModelSelectionExpression) > 0 {
		input.ModelSelectionExpression = aws.String(_apigatewayv2ModelSelectionExpression)
	}
	if len(_apigatewayv2OperationName) > 0 {
		input.OperationName = aws.String(_apigatewayv2OperationName)
	}
	if len(_apigatewayv2RequestModels) > 0 {
		if err := assignInputField(input, "RequestModels", _apigatewayv2RequestModels); err != nil {
			log.Errorf("invalid --request-models: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2RequestParameters) > 0 {
		if err := assignInputField(input, "RequestParameters", _apigatewayv2RequestParameters); err != nil {
			log.Errorf("invalid --request-parameters: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2RouteKey) > 0 {
		input.RouteKey = aws.String(_apigatewayv2RouteKey)
	}
	if len(_apigatewayv2RouteResponseSelectionExpression) > 0 {
		input.RouteResponseSelectionExpression = aws.String(_apigatewayv2RouteResponseSelectionExpression)
	}
	if len(_apigatewayv2Target) > 0 {
		input.Target = aws.String(_apigatewayv2Target)
	}

	if resp, err := client.UpdateRoute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a RouteResponse.
func apigatewayv2_UpdateRouteResponse(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.UpdateRouteResponseInput{
		// ApiId: *string, // Required
		// RouteId: *string, // Required
		// RouteResponseId: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2RouteId) > 0 {
		input.RouteId = aws.String(_apigatewayv2RouteId)
	}
	if len(_apigatewayv2RouteResponseId) > 0 {
		input.RouteResponseId = aws.String(_apigatewayv2RouteResponseId)
	}
	if len(_apigatewayv2ModelSelectionExpression) > 0 {
		input.ModelSelectionExpression = aws.String(_apigatewayv2ModelSelectionExpression)
	}
	if len(_apigatewayv2ResponseModels) > 0 {
		if err := assignInputField(input, "ResponseModels", _apigatewayv2ResponseModels); err != nil {
			log.Errorf("invalid --response-models: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2ResponseParameters) > 0 {
		if err := assignInputField(input, "ResponseParameters", _apigatewayv2ResponseParameters); err != nil {
			log.Errorf("invalid --response-parameters: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2RouteResponseKey) > 0 {
		input.RouteResponseKey = aws.String(_apigatewayv2RouteResponseKey)
	}

	if resp, err := client.UpdateRouteResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Stage.
func apigatewayv2_UpdateStage(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.UpdateStageInput{
		// ApiId: *string, // Required
		// StageName: *string, // Required
	}

	if len(_apigatewayv2ApiId) > 0 {
		input.ApiId = aws.String(_apigatewayv2ApiId)
	}
	if len(_apigatewayv2StageName) > 0 {
		input.StageName = aws.String(_apigatewayv2StageName)
	}
	if len(_apigatewayv2AccessLogSettings) > 0 {
		if err := assignInputField(input, "AccessLogSettings", _apigatewayv2AccessLogSettings); err != nil {
			log.Errorf("invalid --access-log-settings: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2AutoDeploy) > 0 {
		if err := assignInputField(input, "AutoDeploy", _apigatewayv2AutoDeploy); err != nil {
			log.Errorf("invalid --auto-deploy: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2ClientCertificateId) > 0 {
		input.ClientCertificateId = aws.String(_apigatewayv2ClientCertificateId)
	}
	if len(_apigatewayv2DefaultRouteSettings) > 0 {
		if err := assignInputField(input, "DefaultRouteSettings", _apigatewayv2DefaultRouteSettings); err != nil {
			log.Errorf("invalid --default-route-settings: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2DeploymentId) > 0 {
		input.DeploymentId = aws.String(_apigatewayv2DeploymentId)
	}
	if len(_apigatewayv2Description) > 0 {
		input.Description = aws.String(_apigatewayv2Description)
	}
	if len(_apigatewayv2RouteSettings) > 0 {
		if err := assignInputField(input, "RouteSettings", _apigatewayv2RouteSettings); err != nil {
			log.Errorf("invalid --route-settings: %s", err.Error())
			return
		}
	}
	if len(_apigatewayv2StageVariables) > 0 {
		if err := assignInputField(input, "StageVariables", _apigatewayv2StageVariables); err != nil {
			log.Errorf("invalid --stage-variables: %s", err.Error())
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

// Updates a VPC link.
func apigatewayv2_UpdateVpcLink(cfg aws.Config, client *apigatewayv2.Client) {
	input := &apigatewayv2.UpdateVpcLinkInput{
		// VpcLinkId: *string, // Required
	}

	if len(_apigatewayv2VpcLinkId) > 0 {
		input.VpcLinkId = aws.String(_apigatewayv2VpcLinkId)
	}
	if len(_apigatewayv2Name) > 0 {
		input.Name = aws.String(_apigatewayv2Name)
	}

	if resp, err := client.UpdateVpcLink(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_apigatewayv2Cmd)
	_apigatewayv2Cmd.Flags().SortFlags = false

	_apigatewayv2Cmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_apigatewayv2Cmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_apigatewayv2Cmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2AccessLogSettings, "access-log-settings", "", "", "Access Log Settings")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2Actions, "actions", "", "", "Actions")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2ApiId, "api-id", "", "", "API ID")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2ApiKeyRequired, "api-key-required", "", "", "API Key Required")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2ApiKeySelectionExpression, "api-key-selection-expression", "", "", "API Key Selection Expression")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2ApiMappingId, "api-mapping-id", "", "", "API Mapping ID")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2ApiMappingKey, "api-mapping-key", "", "", "API Mapping Key")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2Authorization, "authorization", "", "", "Authorization")
	_apigatewayv2Cmd.Flags().StringSliceVarP(&_apigatewayv2AuthorizationScopes, "authorization-scopes", "", nil, "Authorization Scopes")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2AuthorizationType, "authorization-type", "", "", "Authorization Type")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2AuthorizerCredentialsArn, "authorizer-credentials-arn", "", "", "Authorizer Credentials ARN")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2AuthorizerId, "authorizer-id", "", "", "Authorizer ID")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2AuthorizerPayloadFormatVersion, "authorizer-payload-format-version", "", "", "Authorizer Payload Format Version")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2AuthorizerResultTtlInSeconds, "authorizer-result-ttl-in-seconds", "", "", "Authorizer Result TTL In Seconds")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2AuthorizerType, "authorizer-type", "", "", "Authorizer Type")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2AuthorizerUri, "authorizer-uri", "", "", "Authorizer URI")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2AutoDeploy, "auto-deploy", "", "", "Auto Deploy")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2Basepath, "basepath", "", "", "Basepath")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2Body, "body", "", "", "Body")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2ClientCertificateId, "client-certificate-id", "", "", "Client Certificate ID")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2Conditions, "conditions", "", "", "Conditions")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2ConnectionId, "connection-id", "", "", "Connection ID")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2ConnectionType, "connection-type", "", "", "Connection Type")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2ContentHandlingStrategy, "content-handling-strategy", "", "", "Content Handling Strategy")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2ContentType, "content-type", "", "", "Content Type")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2CorsConfiguration, "cors-configuration", "", "", "Cors Configuration")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2CredentialsArn, "credentials-arn", "", "", "Credentials ARN")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2DefaultRouteSettings, "default-route-settings", "", "", "Default Route Settings")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2DeploymentId, "deployment-id", "", "", "Deployment ID")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2Description, "description", "", "", "Description")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2DisableExecuteApiEndpoint, "disable-execute-api-endpoint", "", "", "Disable Execute API Endpoint")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2DisableSchemaValidation, "disable-schema-validation", "", "", "Disable Schema Validation")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2DisplayContent, "display-content", "", "", "Display Content")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2DisplayName, "display-name", "", "", "Display Name")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2DisplayOrder, "display-order", "", "", "Display Order")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2DomainName, "domain-name", "", "", "Domain Name")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2DomainNameConfigurations, "domain-name-configurations", "", "", "Domain Name Configurations")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2DomainNameId, "domain-name-id", "", "", "Domain Name ID")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2EnableSimpleResponses, "enable-simple-responses", "", "", "Enable Simple Responses")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2EndpointConfiguration, "endpoint-configuration", "", "", "Endpoint Configuration")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2ExportVersion, "export-version", "", "", "Export Version")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2FailOnWarnings, "fail-on-warnings", "", "", "Fail On Warnings")
	_apigatewayv2Cmd.Flags().StringSliceVarP(&_apigatewayv2IdentitySource, "identity-source", "", nil, "Identity Source")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2IdentityValidationExpression, "identity-validation-expression", "", "", "Identity Validation Expression")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2IncludeExtensions, "include-extensions", "", "", "Include Extensions")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2IncludeRawDisplayContent, "include-raw-display-content", "", "", "Include Raw Display Content")
	_apigatewayv2Cmd.Flags().StringSliceVarP(&_apigatewayv2IncludedPortalProductArns, "included-portal-product-arns", "", nil, "Included Portal Product Arns")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2IntegrationId, "integration-id", "", "", "Integration ID")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2IntegrationMethod, "integration-method", "", "", "Integration Method")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2IntegrationResponseId, "integration-response-id", "", "", "Integration Response ID")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2IntegrationResponseKey, "integration-response-key", "", "", "Integration Response Key")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2IntegrationSubtype, "integration-subtype", "", "", "Integration Subtype")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2IntegrationType, "integration-type", "", "", "Integration Type")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2IntegrationUri, "integration-uri", "", "", "Integration URI")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2IpAddressType, "ip-address-type", "", "", "IP Address Type")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2JwtConfiguration, "jwt-configuration", "", "", "Jwt Configuration")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2LogoUri, "logo-uri", "", "", "Logo URI")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2MaxResults, "max-results", "", "", "Max Results")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2ModelId, "model-id", "", "", "Model ID")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2ModelSelectionExpression, "model-selection-expression", "", "", "Model Selection Expression")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2MutualTlsAuthentication, "mutual-tls-authentication", "", "", "Mutual TLS Authentication")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2Name, "name", "", "", "Name")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2NextToken, "next-token", "", "", "Next Token")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2OperationName, "operation-name", "", "", "Operation Name")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2OutputType, "output-type", "", "", "Output Type")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2PassthroughBehavior, "passthrough-behavior", "", "", "Passthrough Behavior")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2PayloadFormatVersion, "payload-format-version", "", "", "Payload Format Version")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2PolicyDocument, "policy-document", "", "", "Policy Document")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2PortalContent, "portal-content", "", "", "Portal Content")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2PortalId, "portal-id", "", "", "Portal ID")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2PortalProductId, "portal-product-id", "", "", "Portal Product ID")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2Priority, "priority", "", "", "Priority")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2ProductPageId, "product-page-id", "", "", "Product Page ID")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2ProductRestEndpointPageId, "product-rest-endpoint-page-id", "", "", "Product Rest Endpoint Page ID")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2ProtocolType, "protocol-type", "", "", "Protocol Type")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2RequestModels, "request-models", "", "", "Request Models")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2RequestParameterKey, "request-parameter-key", "", "", "Request Parameter Key")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2RequestParameters, "request-parameters", "", "", "Request Parameters")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2RequestTemplates, "request-templates", "", "", "Request Templates")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2ResourceArn, "resource-arn", "", "", "Resource ARN")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2ResourceOwner, "resource-owner", "", "", "Resource Owner")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2ResourceOwnerAccountId, "resource-owner-account-id", "", "", "Resource Owner Account ID")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2ResponseModels, "response-models", "", "", "Response Models")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2ResponseParameters, "response-parameters", "", "", "Response Parameters")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2ResponseTemplates, "response-templates", "", "", "Response Templates")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2RestEndpointIdentifier, "rest-endpoint-identifier", "", "", "Rest Endpoint Identifier")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2RouteId, "route-id", "", "", "Route ID")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2RouteKey, "route-key", "", "", "Route Key")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2RouteResponseId, "route-response-id", "", "", "Route Response ID")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2RouteResponseKey, "route-response-key", "", "", "Route Response Key")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2RouteResponseSelectionExpression, "route-response-selection-expression", "", "", "Route Response Selection Expression")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2RouteSelectionExpression, "route-selection-expression", "", "", "Route Selection Expression")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2RouteSettings, "route-settings", "", "", "Route Settings")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2RoutingMode, "routing-mode", "", "", "Routing Mode")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2RoutingRuleId, "routing-rule-id", "", "", "Routing Rule ID")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2RumAppMonitorName, "rum-app-monitor-name", "", "", "Rum App Monitor Name")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2Schema, "schema", "", "", "Schema")
	_apigatewayv2Cmd.Flags().StringSliceVarP(&_apigatewayv2SecurityGroupIds, "security-group-ids", "", nil, "Security Group Ids")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2Specification, "specification", "", "", "Specification")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2Stage, "stage", "", "", "Stage")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2StageName, "stage-name", "", "", "Stage Name")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2StageVariables, "stage-variables", "", "", "Stage Variables")
	_apigatewayv2Cmd.Flags().StringSliceVarP(&_apigatewayv2SubnetIds, "subnet-ids", "", nil, "Subnet Ids")
	_apigatewayv2Cmd.Flags().StringSliceVarP(&_apigatewayv2TagKeys, "tag-keys", "", nil, "Tag Keys")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2Tags, "tags", "", "", "Tags")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2Target, "target", "", "", "Target")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2TemplateSelectionExpression, "template-selection-expression", "", "", "Template Selection Expression")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2TimeoutInMillis, "timeout-in-millis", "", "", "Timeout In Millis")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2TlsConfig, "tls-config", "", "", "TLS Config")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2TryItState, "try-it-state", "", "", "Try It State")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2Version, "version", "", "", "Version")
	_apigatewayv2Cmd.Flags().StringVarP(&_apigatewayv2VpcLinkId, "vpc-link-id", "", "", "VPC Link ID")

	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2CreateApi, "create-api", "", false, "Create API")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2CreateApiMapping, "create-api-mapping", "", false, "Create API Mapping")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2CreateAuthorizer, "create-authorizer", "", false, "Create Authorizer")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2CreateDeployment, "create-deployment", "", false, "Create Deployment")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2CreateDomainName, "create-domain-name", "", false, "Create Domain Name")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2CreateIntegration, "create-integration", "", false, "Create Integration")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2CreateIntegrationResponse, "create-integration-response", "", false, "Create Integration Response")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2CreateModel, "create-model", "", false, "Create Model")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2CreatePortal, "create-portal", "", false, "Create Portal")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2CreatePortalProduct, "create-portal-product", "", false, "Create Portal Product")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2CreateProductPage, "create-product-page", "", false, "Create Product Page")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2CreateProductRestEndpointPage, "create-product-rest-endpoint-page", "", false, "Create Product Rest Endpoint Page")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2CreateRoute, "create-route", "", false, "Create Route")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2CreateRouteResponse, "create-route-response", "", false, "Create Route Response")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2CreateRoutingRule, "create-routing-rule", "", false, "Create Routing Rule")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2CreateStage, "create-stage", "", false, "Create Stage")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2CreateVpcLink, "create-vpc-link", "", false, "Create VPC Link")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2DeleteAccessLogSettings, "delete-access-log-settings", "", false, "Delete Access Log Settings")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2DeleteApi, "delete-api", "", false, "Delete API")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2DeleteApiMapping, "delete-api-mapping", "", false, "Delete API Mapping")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2DeleteAuthorizer, "delete-authorizer", "", false, "Delete Authorizer")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2DeleteCorsConfiguration, "delete-cors-configuration", "", false, "Delete Cors Configuration")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2DeleteDeployment, "delete-deployment", "", false, "Delete Deployment")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2DeleteDomainName, "delete-domain-name", "", false, "Delete Domain Name")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2DeleteIntegration, "delete-integration", "", false, "Delete Integration")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2DeleteIntegrationResponse, "delete-integration-response", "", false, "Delete Integration Response")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2DeleteModel, "delete-model", "", false, "Delete Model")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2DeletePortal, "delete-portal", "", false, "Delete Portal")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2DeletePortalProduct, "delete-portal-product", "", false, "Delete Portal Product")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2DeletePortalProductSharingPolicy, "delete-portal-product-sharing-policy", "", false, "Delete Portal Product Sharing Policy")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2DeleteProductPage, "delete-product-page", "", false, "Delete Product Page")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2DeleteProductRestEndpointPage, "delete-product-rest-endpoint-page", "", false, "Delete Product Rest Endpoint Page")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2DeleteRoute, "delete-route", "", false, "Delete Route")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2DeleteRouteRequestParameter, "delete-route-request-parameter", "", false, "Delete Route Request Parameter")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2DeleteRouteResponse, "delete-route-response", "", false, "Delete Route Response")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2DeleteRouteSettings, "delete-route-settings", "", false, "Delete Route Settings")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2DeleteRoutingRule, "delete-routing-rule", "", false, "Delete Routing Rule")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2DeleteStage, "delete-stage", "", false, "Delete Stage")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2DeleteVpcLink, "delete-vpc-link", "", false, "Delete VPC Link")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2DisablePortal, "disable-portal", "", false, "Disable Portal")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2ExportApi, "export-api", "", false, "Export API")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetApi, "get-api", "", false, "Get API")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetApiMapping, "get-api-mapping", "", false, "Get API Mapping")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetApiMappings, "get-api-mappings", "", false, "Get API Mappings")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetApis, "get-apis", "", false, "Get Apis")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetAuthorizer, "get-authorizer", "", false, "Get Authorizer")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetAuthorizers, "get-authorizers", "", false, "Get Authorizers")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetDeployment, "get-deployment", "", false, "Get Deployment")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetDeployments, "get-deployments", "", false, "Get Deployments")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetDomainName, "get-domain-name", "", false, "Get Domain Name")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetDomainNames, "get-domain-names", "", false, "Get Domain Names")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetIntegration, "get-integration", "", false, "Get Integration")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetIntegrationResponse, "get-integration-response", "", false, "Get Integration Response")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetIntegrationResponses, "get-integration-responses", "", false, "Get Integration Responses")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetIntegrations, "get-integrations", "", false, "Get Integrations")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetModel, "get-model", "", false, "Get Model")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetModelTemplate, "get-model-template", "", false, "Get Model Template")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetModels, "get-models", "", false, "Get Models")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetPortal, "get-portal", "", false, "Get Portal")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetPortalProduct, "get-portal-product", "", false, "Get Portal Product")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetPortalProductSharingPolicy, "get-portal-product-sharing-policy", "", false, "Get Portal Product Sharing Policy")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetProductPage, "get-product-page", "", false, "Get Product Page")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetProductRestEndpointPage, "get-product-rest-endpoint-page", "", false, "Get Product Rest Endpoint Page")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetRoute, "get-route", "", false, "Get Route")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetRouteResponse, "get-route-response", "", false, "Get Route Response")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetRouteResponses, "get-route-responses", "", false, "Get Route Responses")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetRoutes, "get-routes", "", false, "Get Routes")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetRoutingRule, "get-routing-rule", "", false, "Get Routing Rule")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetStage, "get-stage", "", false, "Get Stage")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetStages, "get-stages", "", false, "Get Stages")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetTags, "get-tags", "", false, "Get Tags")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetVpcLink, "get-vpc-link", "", false, "Get VPC Link")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2GetVpcLinks, "get-vpc-links", "", false, "Get VPC Links")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2ImportApi, "import-api", "", false, "Import API")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2ListPortalProducts, "list-portal-products", "", false, "List Portal Products")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2ListPortals, "list-portals", "", false, "List Portals")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2ListProductPages, "list-product-pages", "", false, "List Product Pages")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2ListProductRestEndpointPages, "list-product-rest-endpoint-pages", "", false, "List Product Rest Endpoint Pages")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2ListRoutingRules, "list-routing-rules", "", false, "List Routing Rules")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2PreviewPortal, "preview-portal", "", false, "Preview Portal")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2PublishPortal, "publish-portal", "", false, "Publish Portal")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2PutPortalProductSharingPolicy, "put-portal-product-sharing-policy", "", false, "Put Portal Product Sharing Policy")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2PutRoutingRule, "put-routing-rule", "", false, "Put Routing Rule")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2ReimportApi, "reimport-api", "", false, "Reimport API")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2ResetAuthorizersCache, "reset-authorizers-cache", "", false, "Reset Authorizers Cache")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2TagResource, "tag-resource", "", false, "Tag Resource")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2UntagResource, "untag-resource", "", false, "Untag Resource")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2UpdateApi, "update-api", "", false, "Update API")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2UpdateApiMapping, "update-api-mapping", "", false, "Update API Mapping")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2UpdateAuthorizer, "update-authorizer", "", false, "Update Authorizer")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2UpdateDeployment, "update-deployment", "", false, "Update Deployment")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2UpdateDomainName, "update-domain-name", "", false, "Update Domain Name")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2UpdateIntegration, "update-integration", "", false, "Update Integration")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2UpdateIntegrationResponse, "update-integration-response", "", false, "Update Integration Response")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2UpdateModel, "update-model", "", false, "Update Model")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2UpdatePortal, "update-portal", "", false, "Update Portal")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2UpdatePortalProduct, "update-portal-product", "", false, "Update Portal Product")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2UpdateProductPage, "update-product-page", "", false, "Update Product Page")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2UpdateProductRestEndpointPage, "update-product-rest-endpoint-page", "", false, "Update Product Rest Endpoint Page")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2UpdateRoute, "update-route", "", false, "Update Route")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2UpdateRouteResponse, "update-route-response", "", false, "Update Route Response")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2UpdateStage, "update-stage", "", false, "Update Stage")
	_apigatewayv2Cmd.Flags().BoolVarP(&_apigatewayv2UpdateVpcLink, "update-vpc-link", "", false, "Update VPC Link")

}
