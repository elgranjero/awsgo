package bedrockagentcorecontrol

// StartPolicyGeneration is generated as a reference stub.
// Executable command wiring lives under cmd/bedrockagentcorecontrol.go.
//
// Initiates the AI-powered generation of Cedar policies from natural language
// descriptions within the AgentCore Policy system. This feature enables both
// technical and non-technical users to create policies by describing their
// authorization requirements in plain English, which is then automatically
// translated into formal Cedar policy statements. The generation process analyzes
// the natural language input along with the Gateway's tool context to produce
// validated policy options. Generated policy assets are automatically deleted
// after 7 days, so you should review and create policies from the generated assets
// within this timeframe. Once created, policies are permanent and not subject to
// this expiration. Generated policies should be reviewed and tested in log-only
// mode before deploying to production. Use this when you want to describe policy
// intent naturally rather than learning Cedar syntax, though generated policies
// may require refinement for complex scenarios.
