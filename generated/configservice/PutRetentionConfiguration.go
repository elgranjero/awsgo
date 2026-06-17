package configservice

// PutRetentionConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Creates and updates the retention configuration with details about retention
// period (number of days) that Config stores your historical information. The API
// creates the RetentionConfiguration object and names the object as default. When
// you have a RetentionConfiguration object named default, calling the API
// modifies the default object.
//
// Currently, Config supports only one retention configuration per region in your
// account.
