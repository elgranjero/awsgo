package bedrockagentcore

// SaveBrowserSessionProfile is generated as a reference stub.
// Executable command wiring lives under cmd/bedrockagentcore.go.
//
// Saves the current state of a browser session as a reusable profile in Amazon
// Bedrock AgentCore. A browser profile captures persistent browser data such as
// cookies and local storage from an active session, enabling you to reuse this
// data in future browser sessions.
//
// To save a browser session profile, you must specify the profile identifier,
// browser identifier, and session ID. The session must be active when saving the
// profile. Once saved, the profile can be used with the StartBrowserSession
// operation to initialize new sessions with the stored browser state.
//
// Browser profiles are useful for scenarios that require persistent
// authentication, maintaining user preferences across sessions, or continuing
// tasks that depend on previously stored browser data.
//
// The following operations are related to SaveBrowserSessionProfile :
//
// [StartBrowserSession]
//
// [GetBrowserSession]
//
// [GetBrowserSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_GetBrowserSession.html
// [StartBrowserSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_StartBrowserSession.html
