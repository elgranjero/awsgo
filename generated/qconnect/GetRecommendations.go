package qconnect

// GetRecommendations is generated as a reference stub.
// Executable command wiring lives under cmd/qconnect.go.
//
// This API will be discontinued starting June 1, 2024. To receive generative
// responses after March 1, 2024, you will need to create a new Assistant in the
// Amazon Connect console and integrate the Amazon Q in Connect JavaScript library
// (amazon-q-connectjs) into your applications.
//
// Retrieves recommendations for the specified session. To avoid retrieving the
// same recommendations in subsequent calls, use [NotifyRecommendationsReceived]. This API supports long-polling
// behavior with the waitTimeSeconds parameter. Short poll is the default behavior
// and only returns recommendations already available. To perform a manual query
// against an assistant, use [QueryAssistant].
//
// Deprecated: GetRecommendations API will be discontinued starting June 1, 2024.
// To receive generative responses after March 1, 2024 you will need to create a
// new Assistant in the Connect console and integrate the Amazon Q in Connect
// JavaScript library (amazon-q-connectjs) into your applications.
//
// [QueryAssistant]: https://docs.aws.amazon.com/amazon-q-connect/latest/APIReference/API_QueryAssistant.html
// [NotifyRecommendationsReceived]: https://docs.aws.amazon.com/amazon-q-connect/latest/APIReference/API_NotifyRecommendationsReceived.html
