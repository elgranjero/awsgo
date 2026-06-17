package connectparticipant

// CreateParticipantConnection is generated as a reference stub.
// Executable command wiring lives under cmd/connectparticipant.go.
//
// Creates the participant's connection.
//
// For security recommendations, see [Amazon Connect Chat security best practices].
//
// For WebRTC security recommendations, see [Amazon Connect WebRTC security best practices].
//
// ParticipantToken is used for invoking this API instead of ConnectionToken .
//
// The participant token is valid for the lifetime of the participant – until they
// are part of a contact. For WebRTC participants, if they leave or are
// disconnected for 60 seconds, a new participant needs to be created using the [CreateParticipant]
// API.
//
// For WEBSOCKET Type:
//
// The response URL for has a connect expiry timeout of 100s. Clients must
// manually connect to the returned websocket URL and subscribe to the desired
// topic.
//
// For chat, you need to publish the following on the established websocket
// connection:
//
// {"topic":"aws/subscribe","content":{"topics":["aws/chat"]}}
//
// Upon websocket URL expiry, as specified in the response ConnectionExpiry
// parameter, clients need to call this API again to obtain a new websocket URL and
// perform the same steps as before.
//
// The expiry time for the connection token is different than the
// ChatDurationInMinutes . Expiry time for the connection token is 1 day.
//
// For WEBRTC_CONNECTION Type:
//
// The response includes connection data required for the client application to
// join the call using the Amazon Chime SDK client libraries. The WebRTCConnection
// response contains Meeting and Attendee information needed to establish the media
// connection.
//
// The attendee join token in WebRTCConnection response is valid for the lifetime
// of the participant in the call. If a participant leaves or is disconnected for
// 60 seconds, their participant credentials will no longer be valid, and a new
// participant will need to be created to rejoin the call.
//
// Message streaming support: This API can also be used together with the [StartContactStreaming] API to
// create a participant connection for chat contacts that are not using a
// websocket. For more information about message streaming, [Enable real-time chat message streaming]in the Amazon Connect
// Administrator Guide.
//
// Multi-user web, in-app, video calling support:
//
// For WebRTC calls, this API is used in conjunction with the CreateParticipant
// API to enable multi-party calling. The StartWebRTCContact API creates the
// initial contact and routes it to an agent, while CreateParticipant adds
// additional participants to the ongoing call. For more information about
// multi-party WebRTC calls, see [Enable multi-user web, in-app, and video calling]in the Amazon Connect Administrator Guide.
//
// Feature specifications: For information about feature specifications, such as
// the allowed number of open websocket connections per participant or maximum
// number of WebRTC participants, see [Feature specifications]in the Amazon Connect Administrator Guide.
//
// The Amazon Connect Participant Service APIs do not use [Signature Version 4 authentication].
//
// [Feature specifications]: https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html#feature-limits
// [StartContactStreaming]: https://docs.aws.amazon.com/connect/latest/APIReference/API_StartContactStreaming.html
// [CreateParticipant]: https://docs.aws.amazon.com/connect/latest/APIReference/API_CreateParticipant.html
// [Amazon Connect WebRTC security best practices]: https://docs.aws.amazon.com/connect/latest/adminguide/security-best-practices.html#bp-webrtc-security
// [Enable real-time chat message streaming]: https://docs.aws.amazon.com/connect/latest/adminguide/chat-message-streaming.html
// [Signature Version 4 authentication]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
// [Enable multi-user web, in-app, and video calling]: https://docs.aws.amazon.com/connect/latest/adminguide/enable-multiuser-inapp.html
// [Amazon Connect Chat security best practices]: https://docs.aws.amazon.com/connect/latest/adminguide/security-best-practices.html#bp-security-chat
