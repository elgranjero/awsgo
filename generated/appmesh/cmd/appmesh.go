package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// appmeshCmd represents the appmesh command
var _appmeshCmd = &cobra.Command{
	Use:   "appmesh",
	Short: "AWS appmesh CLI",
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
		client := appmesh.NewFromConfig(cfg)
		if _appmeshCreateGatewayRoute {
			appmesh_CreateGatewayRoute(cfg, client)
			return
		}
		if _appmeshCreateMesh {
			appmesh_CreateMesh(cfg, client)
			return
		}
		if _appmeshCreateRoute {
			appmesh_CreateRoute(cfg, client)
			return
		}
		if _appmeshCreateVirtualGateway {
			appmesh_CreateVirtualGateway(cfg, client)
			return
		}
		if _appmeshCreateVirtualNode {
			appmesh_CreateVirtualNode(cfg, client)
			return
		}
		if _appmeshCreateVirtualRouter {
			appmesh_CreateVirtualRouter(cfg, client)
			return
		}
		if _appmeshCreateVirtualService {
			appmesh_CreateVirtualService(cfg, client)
			return
		}
		if _appmeshDeleteGatewayRoute {
			appmesh_DeleteGatewayRoute(cfg, client)
			return
		}
		if _appmeshDeleteMesh {
			appmesh_DeleteMesh(cfg, client)
			return
		}
		if _appmeshDeleteRoute {
			appmesh_DeleteRoute(cfg, client)
			return
		}
		if _appmeshDeleteVirtualGateway {
			appmesh_DeleteVirtualGateway(cfg, client)
			return
		}
		if _appmeshDeleteVirtualNode {
			appmesh_DeleteVirtualNode(cfg, client)
			return
		}
		if _appmeshDeleteVirtualRouter {
			appmesh_DeleteVirtualRouter(cfg, client)
			return
		}
		if _appmeshDeleteVirtualService {
			appmesh_DeleteVirtualService(cfg, client)
			return
		}
		if _appmeshDescribeGatewayRoute {
			appmesh_DescribeGatewayRoute(cfg, client)
			return
		}
		if _appmeshDescribeMesh {
			appmesh_DescribeMesh(cfg, client)
			return
		}
		if _appmeshDescribeRoute {
			appmesh_DescribeRoute(cfg, client)
			return
		}
		if _appmeshDescribeVirtualGateway {
			appmesh_DescribeVirtualGateway(cfg, client)
			return
		}
		if _appmeshDescribeVirtualNode {
			appmesh_DescribeVirtualNode(cfg, client)
			return
		}
		if _appmeshDescribeVirtualRouter {
			appmesh_DescribeVirtualRouter(cfg, client)
			return
		}
		if _appmeshDescribeVirtualService {
			appmesh_DescribeVirtualService(cfg, client)
			return
		}
		if _appmeshListGatewayRoutes {
			appmesh_ListGatewayRoutes(cfg, client)
			return
		}
		if _appmeshListMeshes {
			appmesh_ListMeshes(cfg, client)
			return
		}
		if _appmeshListRoutes {
			appmesh_ListRoutes(cfg, client)
			return
		}
		if _appmeshListTagsForResource {
			appmesh_ListTagsForResource(cfg, client)
			return
		}
		if _appmeshListVirtualGateways {
			appmesh_ListVirtualGateways(cfg, client)
			return
		}
		if _appmeshListVirtualNodes {
			appmesh_ListVirtualNodes(cfg, client)
			return
		}
		if _appmeshListVirtualRouters {
			appmesh_ListVirtualRouters(cfg, client)
			return
		}
		if _appmeshListVirtualServices {
			appmesh_ListVirtualServices(cfg, client)
			return
		}
		if _appmeshTagResource {
			appmesh_TagResource(cfg, client)
			return
		}
		if _appmeshUntagResource {
			appmesh_UntagResource(cfg, client)
			return
		}
		if _appmeshUpdateGatewayRoute {
			appmesh_UpdateGatewayRoute(cfg, client)
			return
		}
		if _appmeshUpdateMesh {
			appmesh_UpdateMesh(cfg, client)
			return
		}
		if _appmeshUpdateRoute {
			appmesh_UpdateRoute(cfg, client)
			return
		}
		if _appmeshUpdateVirtualGateway {
			appmesh_UpdateVirtualGateway(cfg, client)
			return
		}
		if _appmeshUpdateVirtualNode {
			appmesh_UpdateVirtualNode(cfg, client)
			return
		}
		if _appmeshUpdateVirtualRouter {
			appmesh_UpdateVirtualRouter(cfg, client)
			return
		}
		if _appmeshUpdateVirtualService {
			appmesh_UpdateVirtualService(cfg, client)
			return
		}

	},
}

var (
	_appmeshCreateGatewayRoute     bool
	_appmeshCreateMesh             bool
	_appmeshCreateRoute            bool
	_appmeshCreateVirtualGateway   bool
	_appmeshCreateVirtualNode      bool
	_appmeshCreateVirtualRouter    bool
	_appmeshCreateVirtualService   bool
	_appmeshDeleteGatewayRoute     bool
	_appmeshDeleteMesh             bool
	_appmeshDeleteRoute            bool
	_appmeshDeleteVirtualGateway   bool
	_appmeshDeleteVirtualNode      bool
	_appmeshDeleteVirtualRouter    bool
	_appmeshDeleteVirtualService   bool
	_appmeshDescribeGatewayRoute   bool
	_appmeshDescribeMesh           bool
	_appmeshDescribeRoute          bool
	_appmeshDescribeVirtualGateway bool
	_appmeshDescribeVirtualNode    bool
	_appmeshDescribeVirtualRouter  bool
	_appmeshDescribeVirtualService bool
	_appmeshListGatewayRoutes      bool
	_appmeshListMeshes             bool
	_appmeshListRoutes             bool
	_appmeshListTagsForResource    bool
	_appmeshListVirtualGateways    bool
	_appmeshListVirtualNodes       bool
	_appmeshListVirtualRouters     bool
	_appmeshListVirtualServices    bool
	_appmeshTagResource            bool
	_appmeshUntagResource          bool
	_appmeshUpdateGatewayRoute     bool
	_appmeshUpdateMesh             bool
	_appmeshUpdateRoute            bool
	_appmeshUpdateVirtualGateway   bool
	_appmeshUpdateVirtualNode      bool
	_appmeshUpdateVirtualRouter    bool
	_appmeshUpdateVirtualService   bool

	_appmeshClientToken        string
	_appmeshGatewayRouteName   string
	_appmeshLimit              string
	_appmeshMeshName           string
	_appmeshMeshOwner          string
	_appmeshNextToken          string
	_appmeshResourceArn        string
	_appmeshRouteName          string
	_appmeshSpec               string
	_appmeshTagKeys            []string
	_appmeshTags               string
	_appmeshVirtualGatewayName string
	_appmeshVirtualNodeName    string
	_appmeshVirtualRouterName  string
	_appmeshVirtualServiceName string
)

// Creates a gateway route.
// A gateway route is attached to a virtual gateway and routes traffic to an
// existing virtual service. If a route matches a request, it can distribute
// traffic to a target virtual service.
//
// For more information about gateway routes, see [Gateway routes].
//
// [Gateway routes]: https://docs.aws.amazon.com/app-mesh/latest/userguide/gateway-routes.html
func appmesh_CreateGatewayRoute(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.CreateGatewayRouteInput{
		// GatewayRouteName: *string, // Required
		// MeshName: *string, // Required
		// Spec: *types.GatewayRouteSpec, // Required
		// VirtualGatewayName: *string, // Required
	}

	if len(_appmeshGatewayRouteName) > 0 {
		input.GatewayRouteName = aws.String(_appmeshGatewayRouteName)
	}
	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshSpec) > 0 {
		if err := assignInputField(input, "Spec", _appmeshSpec); err != nil {
			log.Errorf("invalid --spec: %s", err.Error())
			return
		}
	}
	if len(_appmeshVirtualGatewayName) > 0 {
		input.VirtualGatewayName = aws.String(_appmeshVirtualGatewayName)
	}
	if len(_appmeshClientToken) > 0 {
		input.ClientToken = aws.String(_appmeshClientToken)
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}
	if len(_appmeshTags) > 0 {
		if err := assignInputField(input, "Tags", _appmeshTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGatewayRoute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a service mesh.
// A service mesh is a logical boundary for network traffic between services that
// are represented by resources within the mesh. After you create your service
// mesh, you can create virtual services, virtual nodes, virtual routers, and
// routes to distribute traffic between the applications in your mesh.
//
// For more information about service meshes, see [Service meshes].
//
// [Service meshes]: https://docs.aws.amazon.com/app-mesh/latest/userguide/meshes.html
func appmesh_CreateMesh(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.CreateMeshInput{
		// MeshName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshClientToken) > 0 {
		input.ClientToken = aws.String(_appmeshClientToken)
	}
	if len(_appmeshSpec) > 0 {
		if err := assignInputField(input, "Spec", _appmeshSpec); err != nil {
			log.Errorf("invalid --spec: %s", err.Error())
			return
		}
	}
	if len(_appmeshTags) > 0 {
		if err := assignInputField(input, "Tags", _appmeshTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMesh(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a route that is associated with a virtual router.
// You can route several different protocols and define a retry policy for a
// route. Traffic can be routed to one or more virtual nodes.
//
// For more information about routes, see [Routes].
//
// [Routes]: https://docs.aws.amazon.com/app-mesh/latest/userguide/routes.html
func appmesh_CreateRoute(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.CreateRouteInput{
		// MeshName: *string, // Required
		// RouteName: *string, // Required
		// Spec: *types.RouteSpec, // Required
		// VirtualRouterName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshRouteName) > 0 {
		input.RouteName = aws.String(_appmeshRouteName)
	}
	if len(_appmeshSpec) > 0 {
		if err := assignInputField(input, "Spec", _appmeshSpec); err != nil {
			log.Errorf("invalid --spec: %s", err.Error())
			return
		}
	}
	if len(_appmeshVirtualRouterName) > 0 {
		input.VirtualRouterName = aws.String(_appmeshVirtualRouterName)
	}
	if len(_appmeshClientToken) > 0 {
		input.ClientToken = aws.String(_appmeshClientToken)
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}
	if len(_appmeshTags) > 0 {
		if err := assignInputField(input, "Tags", _appmeshTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRoute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a virtual gateway.
// A virtual gateway allows resources outside your mesh to communicate to
// resources that are inside your mesh. The virtual gateway represents an Envoy
// proxy running in an Amazon ECS task, in a Kubernetes service, or on an Amazon
// EC2 instance. Unlike a virtual node, which represents an Envoy running with an
// application, a virtual gateway represents Envoy deployed by itself.
//
// For more information about virtual gateways, see [Virtual gateways].
//
// [Virtual gateways]: https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_gateways.html
func appmesh_CreateVirtualGateway(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.CreateVirtualGatewayInput{
		// MeshName: *string, // Required
		// Spec: *types.VirtualGatewaySpec, // Required
		// VirtualGatewayName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshSpec) > 0 {
		if err := assignInputField(input, "Spec", _appmeshSpec); err != nil {
			log.Errorf("invalid --spec: %s", err.Error())
			return
		}
	}
	if len(_appmeshVirtualGatewayName) > 0 {
		input.VirtualGatewayName = aws.String(_appmeshVirtualGatewayName)
	}
	if len(_appmeshClientToken) > 0 {
		input.ClientToken = aws.String(_appmeshClientToken)
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}
	if len(_appmeshTags) > 0 {
		if err := assignInputField(input, "Tags", _appmeshTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateVirtualGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a virtual node within a service mesh.
// A virtual node acts as a logical pointer to a particular task group, such as an
// Amazon ECS service or a Kubernetes deployment. When you create a virtual node,
// you can specify the service discovery information for your task group, and
// whether the proxy running in a task group will communicate with other proxies
// using Transport Layer Security (TLS).
//
// You define a listener for any inbound traffic that your virtual node expects.
// Any virtual service that your virtual node expects to communicate to is
// specified as a backend .
//
// The response metadata for your new virtual node contains the arn that is
// associated with the virtual node. Set this value to the full ARN; for example,
// arn:aws:appmesh:us-west-2:123456789012:myMesh/default/virtualNode/myApp ) as the
// APPMESH_RESOURCE_ARN environment variable for your task group's Envoy proxy
// container in your task definition or pod spec. This is then mapped to the
// node.id and node.cluster Envoy parameters.
//
// By default, App Mesh uses the name of the resource you specified in
// APPMESH_RESOURCE_ARN when Envoy is referring to itself in metrics and traces.
// You can override this behavior by setting the APPMESH_RESOURCE_CLUSTER
// environment variable with your own name.
//
// For more information about virtual nodes, see [Virtual nodes]. You must be using 1.15.0 or
// later of the Envoy image when setting these variables. For more information
// aboutApp Mesh Envoy variables, see [Envoy image]in the App Mesh User Guide.
//
// [Virtual nodes]: https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_nodes.html
// [Envoy image]: https://docs.aws.amazon.com/app-mesh/latest/userguide/envoy.html
func appmesh_CreateVirtualNode(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.CreateVirtualNodeInput{
		// MeshName: *string, // Required
		// Spec: *types.VirtualNodeSpec, // Required
		// VirtualNodeName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshSpec) > 0 {
		if err := assignInputField(input, "Spec", _appmeshSpec); err != nil {
			log.Errorf("invalid --spec: %s", err.Error())
			return
		}
	}
	if len(_appmeshVirtualNodeName) > 0 {
		input.VirtualNodeName = aws.String(_appmeshVirtualNodeName)
	}
	if len(_appmeshClientToken) > 0 {
		input.ClientToken = aws.String(_appmeshClientToken)
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}
	if len(_appmeshTags) > 0 {
		if err := assignInputField(input, "Tags", _appmeshTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateVirtualNode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a virtual router within a service mesh.
// Specify a listener for any inbound traffic that your virtual router receives.
// Create a virtual router for each protocol and port that you need to route.
// Virtual routers handle traffic for one or more virtual services within your
// mesh. After you create your virtual router, create and associate routes for your
// virtual router that direct incoming requests to different virtual nodes.
//
// For more information about virtual routers, see [Virtual routers].
//
// [Virtual routers]: https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_routers.html
func appmesh_CreateVirtualRouter(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.CreateVirtualRouterInput{
		// MeshName: *string, // Required
		// Spec: *types.VirtualRouterSpec, // Required
		// VirtualRouterName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshSpec) > 0 {
		if err := assignInputField(input, "Spec", _appmeshSpec); err != nil {
			log.Errorf("invalid --spec: %s", err.Error())
			return
		}
	}
	if len(_appmeshVirtualRouterName) > 0 {
		input.VirtualRouterName = aws.String(_appmeshVirtualRouterName)
	}
	if len(_appmeshClientToken) > 0 {
		input.ClientToken = aws.String(_appmeshClientToken)
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}
	if len(_appmeshTags) > 0 {
		if err := assignInputField(input, "Tags", _appmeshTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateVirtualRouter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a virtual service within a service mesh.
// A virtual service is an abstraction of a real service that is provided by a
// virtual node directly or indirectly by means of a virtual router. Dependent
// services call your virtual service by its virtualServiceName , and those
// requests are routed to the virtual node or virtual router that is specified as
// the provider for the virtual service.
//
// For more information about virtual services, see [Virtual services].
//
// [Virtual services]: https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_services.html
func appmesh_CreateVirtualService(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.CreateVirtualServiceInput{
		// MeshName: *string, // Required
		// Spec: *types.VirtualServiceSpec, // Required
		// VirtualServiceName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshSpec) > 0 {
		if err := assignInputField(input, "Spec", _appmeshSpec); err != nil {
			log.Errorf("invalid --spec: %s", err.Error())
			return
		}
	}
	if len(_appmeshVirtualServiceName) > 0 {
		input.VirtualServiceName = aws.String(_appmeshVirtualServiceName)
	}
	if len(_appmeshClientToken) > 0 {
		input.ClientToken = aws.String(_appmeshClientToken)
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}
	if len(_appmeshTags) > 0 {
		if err := assignInputField(input, "Tags", _appmeshTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateVirtualService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing gateway route.
func appmesh_DeleteGatewayRoute(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.DeleteGatewayRouteInput{
		// GatewayRouteName: *string, // Required
		// MeshName: *string, // Required
		// VirtualGatewayName: *string, // Required
	}

	if len(_appmeshGatewayRouteName) > 0 {
		input.GatewayRouteName = aws.String(_appmeshGatewayRouteName)
	}
	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshVirtualGatewayName) > 0 {
		input.VirtualGatewayName = aws.String(_appmeshVirtualGatewayName)
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}

	if resp, err := client.DeleteGatewayRoute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing service mesh.
// You must delete all resources (virtual services, routes, virtual routers, and
// virtual nodes) in the service mesh before you can delete the mesh itself.
func appmesh_DeleteMesh(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.DeleteMeshInput{
		// MeshName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}

	if resp, err := client.DeleteMesh(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing route.
func appmesh_DeleteRoute(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.DeleteRouteInput{
		// MeshName: *string, // Required
		// RouteName: *string, // Required
		// VirtualRouterName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshRouteName) > 0 {
		input.RouteName = aws.String(_appmeshRouteName)
	}
	if len(_appmeshVirtualRouterName) > 0 {
		input.VirtualRouterName = aws.String(_appmeshVirtualRouterName)
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}

	if resp, err := client.DeleteRoute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing virtual gateway. You cannot delete a virtual gateway if any
// gateway routes are associated to it.
func appmesh_DeleteVirtualGateway(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.DeleteVirtualGatewayInput{
		// MeshName: *string, // Required
		// VirtualGatewayName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshVirtualGatewayName) > 0 {
		input.VirtualGatewayName = aws.String(_appmeshVirtualGatewayName)
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}

	if resp, err := client.DeleteVirtualGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing virtual node.
// You must delete any virtual services that list a virtual node as a service
// provider before you can delete the virtual node itself.
func appmesh_DeleteVirtualNode(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.DeleteVirtualNodeInput{
		// MeshName: *string, // Required
		// VirtualNodeName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshVirtualNodeName) > 0 {
		input.VirtualNodeName = aws.String(_appmeshVirtualNodeName)
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}

	if resp, err := client.DeleteVirtualNode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing virtual router.
// You must delete any routes associated with the virtual router before you can
// delete the router itself.
func appmesh_DeleteVirtualRouter(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.DeleteVirtualRouterInput{
		// MeshName: *string, // Required
		// VirtualRouterName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshVirtualRouterName) > 0 {
		input.VirtualRouterName = aws.String(_appmeshVirtualRouterName)
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}

	if resp, err := client.DeleteVirtualRouter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing virtual service.
func appmesh_DeleteVirtualService(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.DeleteVirtualServiceInput{
		// MeshName: *string, // Required
		// VirtualServiceName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshVirtualServiceName) > 0 {
		input.VirtualServiceName = aws.String(_appmeshVirtualServiceName)
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}

	if resp, err := client.DeleteVirtualService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an existing gateway route.
func appmesh_DescribeGatewayRoute(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.DescribeGatewayRouteInput{
		// GatewayRouteName: *string, // Required
		// MeshName: *string, // Required
		// VirtualGatewayName: *string, // Required
	}

	if len(_appmeshGatewayRouteName) > 0 {
		input.GatewayRouteName = aws.String(_appmeshGatewayRouteName)
	}
	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshVirtualGatewayName) > 0 {
		input.VirtualGatewayName = aws.String(_appmeshVirtualGatewayName)
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}

	if resp, err := client.DescribeGatewayRoute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an existing service mesh.
func appmesh_DescribeMesh(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.DescribeMeshInput{
		// MeshName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}

	if resp, err := client.DescribeMesh(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an existing route.
func appmesh_DescribeRoute(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.DescribeRouteInput{
		// MeshName: *string, // Required
		// RouteName: *string, // Required
		// VirtualRouterName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshRouteName) > 0 {
		input.RouteName = aws.String(_appmeshRouteName)
	}
	if len(_appmeshVirtualRouterName) > 0 {
		input.VirtualRouterName = aws.String(_appmeshVirtualRouterName)
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}

	if resp, err := client.DescribeRoute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an existing virtual gateway.
func appmesh_DescribeVirtualGateway(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.DescribeVirtualGatewayInput{
		// MeshName: *string, // Required
		// VirtualGatewayName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshVirtualGatewayName) > 0 {
		input.VirtualGatewayName = aws.String(_appmeshVirtualGatewayName)
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}

	if resp, err := client.DescribeVirtualGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an existing virtual node.
func appmesh_DescribeVirtualNode(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.DescribeVirtualNodeInput{
		// MeshName: *string, // Required
		// VirtualNodeName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshVirtualNodeName) > 0 {
		input.VirtualNodeName = aws.String(_appmeshVirtualNodeName)
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}

	if resp, err := client.DescribeVirtualNode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an existing virtual router.
func appmesh_DescribeVirtualRouter(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.DescribeVirtualRouterInput{
		// MeshName: *string, // Required
		// VirtualRouterName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshVirtualRouterName) > 0 {
		input.VirtualRouterName = aws.String(_appmeshVirtualRouterName)
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}

	if resp, err := client.DescribeVirtualRouter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an existing virtual service.
func appmesh_DescribeVirtualService(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.DescribeVirtualServiceInput{
		// MeshName: *string, // Required
		// VirtualServiceName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshVirtualServiceName) > 0 {
		input.VirtualServiceName = aws.String(_appmeshVirtualServiceName)
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}

	if resp, err := client.DescribeVirtualService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of existing gateway routes that are associated to a virtual
// gateway.
func appmesh_ListGatewayRoutes(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.ListGatewayRoutesInput{
		// MeshName: *string, // Required
		// VirtualGatewayName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshVirtualGatewayName) > 0 {
		input.VirtualGatewayName = aws.String(_appmeshVirtualGatewayName)
	}
	if len(_appmeshLimit) > 0 {
		if err := assignInputField(input, "Limit", _appmeshLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}
	if len(_appmeshNextToken) > 0 {
		input.NextToken = aws.String(_appmeshNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGatewayRoutes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appmesh.ListGatewayRoutesOutput
	p := appmesh.NewListGatewayRoutesPaginator(client, input)
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

// Returns a list of existing service meshes.
func appmesh_ListMeshes(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.ListMeshesInput{}

	if len(_appmeshLimit) > 0 {
		if err := assignInputField(input, "Limit", _appmeshLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_appmeshNextToken) > 0 {
		input.NextToken = aws.String(_appmeshNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMeshes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appmesh.ListMeshesOutput
	p := appmesh.NewListMeshesPaginator(client, input)
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

// Returns a list of existing routes in a service mesh.
func appmesh_ListRoutes(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.ListRoutesInput{
		// MeshName: *string, // Required
		// VirtualRouterName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshVirtualRouterName) > 0 {
		input.VirtualRouterName = aws.String(_appmeshVirtualRouterName)
	}
	if len(_appmeshLimit) > 0 {
		if err := assignInputField(input, "Limit", _appmeshLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}
	if len(_appmeshNextToken) > 0 {
		input.NextToken = aws.String(_appmeshNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRoutes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appmesh.ListRoutesOutput
	p := appmesh.NewListRoutesPaginator(client, input)
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

// List the tags for an App Mesh resource.
func appmesh_ListTagsForResource(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_appmeshResourceArn) > 0 {
		input.ResourceArn = aws.String(_appmeshResourceArn)
	}
	if len(_appmeshLimit) > 0 {
		if err := assignInputField(input, "Limit", _appmeshLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_appmeshNextToken) > 0 {
		input.NextToken = aws.String(_appmeshNextToken)
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

	var results []*appmesh.ListTagsForResourceOutput
	p := appmesh.NewListTagsForResourcePaginator(client, input)
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

// Returns a list of existing virtual gateways in a service mesh.
func appmesh_ListVirtualGateways(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.ListVirtualGatewaysInput{
		// MeshName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshLimit) > 0 {
		if err := assignInputField(input, "Limit", _appmeshLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}
	if len(_appmeshNextToken) > 0 {
		input.NextToken = aws.String(_appmeshNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListVirtualGateways(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appmesh.ListVirtualGatewaysOutput
	p := appmesh.NewListVirtualGatewaysPaginator(client, input)
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

// Returns a list of existing virtual nodes.
func appmesh_ListVirtualNodes(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.ListVirtualNodesInput{
		// MeshName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshLimit) > 0 {
		if err := assignInputField(input, "Limit", _appmeshLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}
	if len(_appmeshNextToken) > 0 {
		input.NextToken = aws.String(_appmeshNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListVirtualNodes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appmesh.ListVirtualNodesOutput
	p := appmesh.NewListVirtualNodesPaginator(client, input)
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

// Returns a list of existing virtual routers in a service mesh.
func appmesh_ListVirtualRouters(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.ListVirtualRoutersInput{
		// MeshName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshLimit) > 0 {
		if err := assignInputField(input, "Limit", _appmeshLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}
	if len(_appmeshNextToken) > 0 {
		input.NextToken = aws.String(_appmeshNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListVirtualRouters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appmesh.ListVirtualRoutersOutput
	p := appmesh.NewListVirtualRoutersPaginator(client, input)
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

// Returns a list of existing virtual services in a service mesh.
func appmesh_ListVirtualServices(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.ListVirtualServicesInput{
		// MeshName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshLimit) > 0 {
		if err := assignInputField(input, "Limit", _appmeshLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}
	if len(_appmeshNextToken) > 0 {
		input.NextToken = aws.String(_appmeshNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListVirtualServices(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*appmesh.ListVirtualServicesOutput
	p := appmesh.NewListVirtualServicesPaginator(client, input)
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

// Associates the specified tags to a resource with the specified resourceArn . If
// existing tags on a resource aren't specified in the request parameters, they
// aren't changed. When a resource is deleted, the tags associated with that
// resource are also deleted.
func appmesh_TagResource(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.TagRef, // Required
	}

	if len(_appmeshResourceArn) > 0 {
		input.ResourceArn = aws.String(_appmeshResourceArn)
	}
	if len(_appmeshTags) > 0 {
		if err := assignInputField(input, "Tags", _appmeshTags); err != nil {
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

// Deletes specified tags from a resource.
func appmesh_UntagResource(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_appmeshResourceArn) > 0 {
		input.ResourceArn = aws.String(_appmeshResourceArn)
	}
	if len(_appmeshTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _appmeshTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing gateway route that is associated to a specified virtual
// gateway in a service mesh.
func appmesh_UpdateGatewayRoute(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.UpdateGatewayRouteInput{
		// GatewayRouteName: *string, // Required
		// MeshName: *string, // Required
		// Spec: *types.GatewayRouteSpec, // Required
		// VirtualGatewayName: *string, // Required
	}

	if len(_appmeshGatewayRouteName) > 0 {
		input.GatewayRouteName = aws.String(_appmeshGatewayRouteName)
	}
	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshSpec) > 0 {
		if err := assignInputField(input, "Spec", _appmeshSpec); err != nil {
			log.Errorf("invalid --spec: %s", err.Error())
			return
		}
	}
	if len(_appmeshVirtualGatewayName) > 0 {
		input.VirtualGatewayName = aws.String(_appmeshVirtualGatewayName)
	}
	if len(_appmeshClientToken) > 0 {
		input.ClientToken = aws.String(_appmeshClientToken)
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}

	if resp, err := client.UpdateGatewayRoute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing service mesh.
func appmesh_UpdateMesh(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.UpdateMeshInput{
		// MeshName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshClientToken) > 0 {
		input.ClientToken = aws.String(_appmeshClientToken)
	}
	if len(_appmeshSpec) > 0 {
		if err := assignInputField(input, "Spec", _appmeshSpec); err != nil {
			log.Errorf("invalid --spec: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMesh(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing route for a specified service mesh and virtual router.
func appmesh_UpdateRoute(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.UpdateRouteInput{
		// MeshName: *string, // Required
		// RouteName: *string, // Required
		// Spec: *types.RouteSpec, // Required
		// VirtualRouterName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshRouteName) > 0 {
		input.RouteName = aws.String(_appmeshRouteName)
	}
	if len(_appmeshSpec) > 0 {
		if err := assignInputField(input, "Spec", _appmeshSpec); err != nil {
			log.Errorf("invalid --spec: %s", err.Error())
			return
		}
	}
	if len(_appmeshVirtualRouterName) > 0 {
		input.VirtualRouterName = aws.String(_appmeshVirtualRouterName)
	}
	if len(_appmeshClientToken) > 0 {
		input.ClientToken = aws.String(_appmeshClientToken)
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}

	if resp, err := client.UpdateRoute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing virtual gateway in a specified service mesh.
func appmesh_UpdateVirtualGateway(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.UpdateVirtualGatewayInput{
		// MeshName: *string, // Required
		// Spec: *types.VirtualGatewaySpec, // Required
		// VirtualGatewayName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshSpec) > 0 {
		if err := assignInputField(input, "Spec", _appmeshSpec); err != nil {
			log.Errorf("invalid --spec: %s", err.Error())
			return
		}
	}
	if len(_appmeshVirtualGatewayName) > 0 {
		input.VirtualGatewayName = aws.String(_appmeshVirtualGatewayName)
	}
	if len(_appmeshClientToken) > 0 {
		input.ClientToken = aws.String(_appmeshClientToken)
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}

	if resp, err := client.UpdateVirtualGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing virtual node in a specified service mesh.
func appmesh_UpdateVirtualNode(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.UpdateVirtualNodeInput{
		// MeshName: *string, // Required
		// Spec: *types.VirtualNodeSpec, // Required
		// VirtualNodeName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshSpec) > 0 {
		if err := assignInputField(input, "Spec", _appmeshSpec); err != nil {
			log.Errorf("invalid --spec: %s", err.Error())
			return
		}
	}
	if len(_appmeshVirtualNodeName) > 0 {
		input.VirtualNodeName = aws.String(_appmeshVirtualNodeName)
	}
	if len(_appmeshClientToken) > 0 {
		input.ClientToken = aws.String(_appmeshClientToken)
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}

	if resp, err := client.UpdateVirtualNode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing virtual router in a specified service mesh.
func appmesh_UpdateVirtualRouter(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.UpdateVirtualRouterInput{
		// MeshName: *string, // Required
		// Spec: *types.VirtualRouterSpec, // Required
		// VirtualRouterName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshSpec) > 0 {
		if err := assignInputField(input, "Spec", _appmeshSpec); err != nil {
			log.Errorf("invalid --spec: %s", err.Error())
			return
		}
	}
	if len(_appmeshVirtualRouterName) > 0 {
		input.VirtualRouterName = aws.String(_appmeshVirtualRouterName)
	}
	if len(_appmeshClientToken) > 0 {
		input.ClientToken = aws.String(_appmeshClientToken)
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}

	if resp, err := client.UpdateVirtualRouter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing virtual service in a specified service mesh.
func appmesh_UpdateVirtualService(cfg aws.Config, client *appmesh.Client) {
	input := &appmesh.UpdateVirtualServiceInput{
		// MeshName: *string, // Required
		// Spec: *types.VirtualServiceSpec, // Required
		// VirtualServiceName: *string, // Required
	}

	if len(_appmeshMeshName) > 0 {
		input.MeshName = aws.String(_appmeshMeshName)
	}
	if len(_appmeshSpec) > 0 {
		if err := assignInputField(input, "Spec", _appmeshSpec); err != nil {
			log.Errorf("invalid --spec: %s", err.Error())
			return
		}
	}
	if len(_appmeshVirtualServiceName) > 0 {
		input.VirtualServiceName = aws.String(_appmeshVirtualServiceName)
	}
	if len(_appmeshClientToken) > 0 {
		input.ClientToken = aws.String(_appmeshClientToken)
	}
	if len(_appmeshMeshOwner) > 0 {
		input.MeshOwner = aws.String(_appmeshMeshOwner)
	}

	if resp, err := client.UpdateVirtualService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_appmeshCmd)
	_appmeshCmd.Flags().SortFlags = false

	_appmeshCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_appmeshCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_appmeshCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_appmeshCmd.Flags().StringVarP(&_appmeshClientToken, "client-token", "", "", "Client Token")
	_appmeshCmd.Flags().StringVarP(&_appmeshGatewayRouteName, "gateway-route-name", "", "", "Gateway Route Name")
	_appmeshCmd.Flags().StringVarP(&_appmeshLimit, "limit", "", "", "Limit")
	_appmeshCmd.Flags().StringVarP(&_appmeshMeshName, "mesh-name", "", "", "Mesh Name")
	_appmeshCmd.Flags().StringVarP(&_appmeshMeshOwner, "mesh-owner", "", "", "Mesh Owner")
	_appmeshCmd.Flags().StringVarP(&_appmeshNextToken, "next-token", "", "", "Next Token")
	_appmeshCmd.Flags().StringVarP(&_appmeshResourceArn, "resource-arn", "", "", "Resource ARN")
	_appmeshCmd.Flags().StringVarP(&_appmeshRouteName, "route-name", "", "", "Route Name")
	_appmeshCmd.Flags().StringVarP(&_appmeshSpec, "spec", "", "", "Spec")
	_appmeshCmd.Flags().StringSliceVarP(&_appmeshTagKeys, "tag-keys", "", nil, "Tag Keys")
	_appmeshCmd.Flags().StringVarP(&_appmeshTags, "tags", "", "", "Tags")
	_appmeshCmd.Flags().StringVarP(&_appmeshVirtualGatewayName, "virtual-gateway-name", "", "", "Virtual Gateway Name")
	_appmeshCmd.Flags().StringVarP(&_appmeshVirtualNodeName, "virtual-node-name", "", "", "Virtual Node Name")
	_appmeshCmd.Flags().StringVarP(&_appmeshVirtualRouterName, "virtual-router-name", "", "", "Virtual Router Name")
	_appmeshCmd.Flags().StringVarP(&_appmeshVirtualServiceName, "virtual-service-name", "", "", "Virtual Service Name")

	_appmeshCmd.Flags().BoolVarP(&_appmeshCreateGatewayRoute, "create-gateway-route", "", false, "Create Gateway Route")
	_appmeshCmd.Flags().BoolVarP(&_appmeshCreateMesh, "create-mesh", "", false, "Create Mesh")
	_appmeshCmd.Flags().BoolVarP(&_appmeshCreateRoute, "create-route", "", false, "Create Route")
	_appmeshCmd.Flags().BoolVarP(&_appmeshCreateVirtualGateway, "create-virtual-gateway", "", false, "Create Virtual Gateway")
	_appmeshCmd.Flags().BoolVarP(&_appmeshCreateVirtualNode, "create-virtual-node", "", false, "Create Virtual Node")
	_appmeshCmd.Flags().BoolVarP(&_appmeshCreateVirtualRouter, "create-virtual-router", "", false, "Create Virtual Router")
	_appmeshCmd.Flags().BoolVarP(&_appmeshCreateVirtualService, "create-virtual-service", "", false, "Create Virtual Service")
	_appmeshCmd.Flags().BoolVarP(&_appmeshDeleteGatewayRoute, "delete-gateway-route", "", false, "Delete Gateway Route")
	_appmeshCmd.Flags().BoolVarP(&_appmeshDeleteMesh, "delete-mesh", "", false, "Delete Mesh")
	_appmeshCmd.Flags().BoolVarP(&_appmeshDeleteRoute, "delete-route", "", false, "Delete Route")
	_appmeshCmd.Flags().BoolVarP(&_appmeshDeleteVirtualGateway, "delete-virtual-gateway", "", false, "Delete Virtual Gateway")
	_appmeshCmd.Flags().BoolVarP(&_appmeshDeleteVirtualNode, "delete-virtual-node", "", false, "Delete Virtual Node")
	_appmeshCmd.Flags().BoolVarP(&_appmeshDeleteVirtualRouter, "delete-virtual-router", "", false, "Delete Virtual Router")
	_appmeshCmd.Flags().BoolVarP(&_appmeshDeleteVirtualService, "delete-virtual-service", "", false, "Delete Virtual Service")
	_appmeshCmd.Flags().BoolVarP(&_appmeshDescribeGatewayRoute, "describe-gateway-route", "", false, "Describe Gateway Route")
	_appmeshCmd.Flags().BoolVarP(&_appmeshDescribeMesh, "describe-mesh", "", false, "Describe Mesh")
	_appmeshCmd.Flags().BoolVarP(&_appmeshDescribeRoute, "describe-route", "", false, "Describe Route")
	_appmeshCmd.Flags().BoolVarP(&_appmeshDescribeVirtualGateway, "describe-virtual-gateway", "", false, "Describe Virtual Gateway")
	_appmeshCmd.Flags().BoolVarP(&_appmeshDescribeVirtualNode, "describe-virtual-node", "", false, "Describe Virtual Node")
	_appmeshCmd.Flags().BoolVarP(&_appmeshDescribeVirtualRouter, "describe-virtual-router", "", false, "Describe Virtual Router")
	_appmeshCmd.Flags().BoolVarP(&_appmeshDescribeVirtualService, "describe-virtual-service", "", false, "Describe Virtual Service")
	_appmeshCmd.Flags().BoolVarP(&_appmeshListGatewayRoutes, "list-gateway-routes", "", false, "List Gateway Routes")
	_appmeshCmd.Flags().BoolVarP(&_appmeshListMeshes, "list-meshes", "", false, "List Meshes")
	_appmeshCmd.Flags().BoolVarP(&_appmeshListRoutes, "list-routes", "", false, "List Routes")
	_appmeshCmd.Flags().BoolVarP(&_appmeshListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_appmeshCmd.Flags().BoolVarP(&_appmeshListVirtualGateways, "list-virtual-gateways", "", false, "List Virtual Gateways")
	_appmeshCmd.Flags().BoolVarP(&_appmeshListVirtualNodes, "list-virtual-nodes", "", false, "List Virtual Nodes")
	_appmeshCmd.Flags().BoolVarP(&_appmeshListVirtualRouters, "list-virtual-routers", "", false, "List Virtual Routers")
	_appmeshCmd.Flags().BoolVarP(&_appmeshListVirtualServices, "list-virtual-services", "", false, "List Virtual Services")
	_appmeshCmd.Flags().BoolVarP(&_appmeshTagResource, "tag-resource", "", false, "Tag Resource")
	_appmeshCmd.Flags().BoolVarP(&_appmeshUntagResource, "untag-resource", "", false, "Untag Resource")
	_appmeshCmd.Flags().BoolVarP(&_appmeshUpdateGatewayRoute, "update-gateway-route", "", false, "Update Gateway Route")
	_appmeshCmd.Flags().BoolVarP(&_appmeshUpdateMesh, "update-mesh", "", false, "Update Mesh")
	_appmeshCmd.Flags().BoolVarP(&_appmeshUpdateRoute, "update-route", "", false, "Update Route")
	_appmeshCmd.Flags().BoolVarP(&_appmeshUpdateVirtualGateway, "update-virtual-gateway", "", false, "Update Virtual Gateway")
	_appmeshCmd.Flags().BoolVarP(&_appmeshUpdateVirtualNode, "update-virtual-node", "", false, "Update Virtual Node")
	_appmeshCmd.Flags().BoolVarP(&_appmeshUpdateVirtualRouter, "update-virtual-router", "", false, "Update Virtual Router")
	_appmeshCmd.Flags().BoolVarP(&_appmeshUpdateVirtualService, "update-virtual-service", "", false, "Update Virtual Service")

}
