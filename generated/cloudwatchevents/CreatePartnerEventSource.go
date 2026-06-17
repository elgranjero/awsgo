package cloudwatchevents

// CreatePartnerEventSource is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatchevents.go.
//
// Called by an SaaS partner to create a partner event source. This operation is
// not used by Amazon Web Services customers.
//
// Each partner event source can be used by one Amazon Web Services account to
// create a matching partner event bus in that Amazon Web Services account. A SaaS
// partner must create one partner event source for each Amazon Web Services
// account that wants to receive those event types.
//
// A partner event source creates events based on resources within the SaaS
// partner's service or application.
//
// An Amazon Web Services account that creates a partner event bus that matches
// the partner event source can use that event bus to receive events from the
// partner, and then process them using Amazon Web Services Events rules and
// targets.
//
// Partner event source names follow this format:
//
// partner_name/event_namespace/event_name
//
// partner_name is determined during partner registration and identifies the
// partner to Amazon Web Services customers. event_namespace is determined by the
// partner and is a way for the partner to categorize their events. event_name is
// determined by the partner, and should uniquely identify an event-generating
// resource within the partner system. The combination of event_namespace and
// event_name should help Amazon Web Services customers decide whether to create an
// event bus to receive these events.
