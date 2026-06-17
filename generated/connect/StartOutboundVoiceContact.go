package connect

// StartOutboundVoiceContact is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Places an outbound call to a contact, and then initiates the flow. It performs
// the actions in the flow that's specified (in ContactFlowId ).
//
// Agents do not initiate the outbound API, which means that they do not dial the
// contact. If the flow places an outbound call to a contact, and then puts the
// contact in queue, the call is then routed to the agent, like any other inbound
// case.
//
// Dialing timeout for this operation can be configured with the
// “RingTimeoutInSeconds” parameter. If not specified, the default dialing timeout
// will be 60 seconds which means if the call is not connected within 60 seconds,
// it fails.
//
// UK numbers with a 447 prefix are not allowed by default. Before you can dial
// these UK mobile numbers, you must submit a service quota increase request. For
// more information, see [Amazon Connect Service Quotas]in the Amazon Connect Administrator Guide.
//
// Campaign calls are not allowed by default. Before you can make a call with
// TrafficType = CAMPAIGN , you must submit a service quota increase request to the
// quota [Amazon Connect campaigns].
//
// For Preview dialing mode, only the Amazon Connect outbound campaigns service
// principal is allowed to assume a role in your account and call this API with
// OutboundStrategy.
//
// [Amazon Connect Service Quotas]: https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html
// [Amazon Connect campaigns]: https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html#outbound-communications-quotas
