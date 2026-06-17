package shield

// EnableApplicationLayerAutomaticResponse is generated as a reference stub.
// Executable command wiring lives under cmd/shield.go.
//
// Enable the Shield Advanced automatic application layer DDoS mitigation for the
// protected resource.
//
// This feature is available for Amazon CloudFront distributions and Application
// Load Balancers only.
//
// This causes Shield Advanced to create, verify, and apply WAF rules for DDoS
// attacks that it detects for the resource. Shield Advanced applies the rules in a
// Shield rule group inside the web ACL that you've associated with the resource.
// For information about how automatic mitigation works and the requirements for
// using it, see [Shield Advanced automatic application layer DDoS mitigation].
//
// Don't use this action to make changes to automatic mitigation settings when
// it's already enabled for a resource. Instead, use UpdateApplicationLayerAutomaticResponse.
//
// To use this feature, you must associate a web ACL with the protected resource.
// The web ACL must be created using the latest version of WAF (v2). You can
// associate the web ACL through the Shield Advanced console at [https://console.aws.amazon.com/wafv2/shieldv2#/]. For more
// information, see [Getting Started with Shield Advanced]. You can also associate the web ACL to the resource through
// the WAF console or the WAF API, but you must manage Shield Advanced automatic
// mitigation through Shield Advanced. For information about WAF, see [WAF Developer Guide].
//
// [https://console.aws.amazon.com/wafv2/shieldv2#/]: https://console.aws.amazon.com/wafv2/shieldv2#/
// [Getting Started with Shield Advanced]: https://docs.aws.amazon.com/waf/latest/developerguide/getting-started-ddos.html
// [WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
// [Shield Advanced automatic application layer DDoS mitigation]: https://docs.aws.amazon.com/waf/latest/developerguide/ddos-advanced-automatic-app-layer-response.html
