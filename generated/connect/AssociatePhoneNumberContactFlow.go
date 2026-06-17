package connect

// AssociatePhoneNumberContactFlow is generated as a reference stub.
// Executable command wiring lives under cmd/connect.go.
//
// Associates a flow with a phone number claimed to your Amazon Connect instance.
//
// If the number is claimed to a traffic distribution group, and you are calling
// this API using an instance in the Amazon Web Services Region where the traffic
// distribution group was created, you can use either a full phone number ARN or
// UUID value for the PhoneNumberId URI request parameter. However, if the number
// is claimed to a traffic distribution group and you are calling this API using an
// instance in the alternate Amazon Web Services Region associated with the traffic
// distribution group, you must provide a full phone number ARN. If a UUID is
// provided in this scenario, you will receive a ResourceNotFoundException .
