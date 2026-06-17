package connect

// ListRoutingProfileManualAssignmentQueues is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Lists the manual assignment queues associated with a routing profile.
//
// # Use cases
//
// Following are common uses cases for this API:
//
// - This API returns list of queues where contacts can be manually assigned or
// picked by an agent who has access to the Worklist app. The user can additionally
// filter on queues, if they have access to those queues (otherwise a invalid
// request exception will be thrown).
//
// For information about how manual contact assignment works in the agent
//
// workspace, see the [Access the Worklist app in the Amazon Connect agent workspace]in the Amazon Connect Administrator Guide.
//
// Important things to know
//
// - This API only returns the manual assignment queues associated with a
// routing profile. Use the ListRoutingProfileQueues API to list the auto
// assignment queues for the routing profile.
//
// Endpoints: See [Amazon Connect endpoints and quotas].
//
// [Amazon Connect endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/connect_region.html
//
// [Access the Worklist app in the Amazon Connect agent workspace]: https://docs.aws.amazon.com/connect/latest/adminguide/worklist-app.html
