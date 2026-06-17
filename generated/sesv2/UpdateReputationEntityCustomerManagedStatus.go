package sesv2

// UpdateReputationEntityCustomerManagedStatus is generated as a reference stub.
// Executable command wiring lives under cmd/sesv2.go.
//
// Update the customer-managed sending status for a reputation entity. This allows
// you to enable, disable, or reinstate sending for the entity.
//
// The customer-managed status works in conjunction with the Amazon Web Services
// Amazon SES-managed status to determine the overall sending capability. When you
// update the customer-managed status, the Amazon Web Services Amazon SES-managed
// status remains unchanged. If Amazon Web Services Amazon SES has disabled the
// entity, it will not be allowed to send regardless of the customer-managed status
// setting. When you reinstate an entity through the customer-managed status, it
// can continue sending only if the Amazon Web Services Amazon SES-managed status
// also permits sending, even if there are active reputation findings, until the
// findings are resolved or new violations occur.
