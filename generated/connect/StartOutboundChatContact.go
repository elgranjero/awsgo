package connect

// StartOutboundChatContact is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Initiates a new outbound SMS or WhatsApp contact to a customer. Response of
// this API provides the ContactId of the outbound SMS or WhatsApp contact created.
//
// SourceEndpoint only supports Endpoints with CONNECT_PHONENUMBER_ARN as Type and
// DestinationEndpoint only supports Endpoints with TELEPHONE_NUMBER as Type.
// ContactFlowId initiates the flow to manage the new contact created.
//
// This API can be used to initiate outbound SMS or WhatsApp contacts for an
// agent, or it can also deflect an ongoing contact to an outbound SMS or WhatsApp
// contact by using the [StartOutboundChatContact]Flow Action.
//
// For more information about using SMS or WhatsApp in Amazon Connect, see the
// following topics in the Amazon Connect Administrator Guide:
//
// [Set up SMS messaging]
//
// [Request an SMS-enabled phone number through Amazon Web Services End User Messaging SMS]
//
// [Set up WhatsApp Business messaging]
//
// [Set up SMS messaging]: https://docs.aws.amazon.com/connect/latest/adminguide/setup-sms-messaging.html
// [Request an SMS-enabled phone number through Amazon Web Services End User Messaging SMS]: https://docs.aws.amazon.com/connect/latest/adminguide/sms-number.html
// [Set up WhatsApp Business messaging]: https://docs.aws.amazon.com/connect/latest/adminguide/whatsapp-integration.html
// [StartOutboundChatContact]: https://docs.aws.amazon.com/connect/latest/APIReference/API_StartOutboundChatContact.html
