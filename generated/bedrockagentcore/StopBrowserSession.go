package bedrockagentcore

// StopBrowserSession is generated as a reference stub.
// Executable command wiring lives under cmd/bedrockagentcore.go.
//
// Terminates an active browser session in Amazon Bedrock AgentCore. This
// operation stops the session, releases associated resources, and makes the
// session unavailable for further use.
//
// To stop a browser session, you must specify both the browser identifier and the
// session ID. Once stopped, a session cannot be restarted; you must create a new
// session using StartBrowserSession .
//
// The following operations are related to StopBrowserSession :
//
// [StartBrowserSession]
//
// [GetBrowserSession]
//
// [GetBrowserSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_GetBrowserSession.html
// [StartBrowserSession]: https://docs.aws.amazon.com/bedrock-agentcore/latest/APIReference/API_StartBrowserSession.html
