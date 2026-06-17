package backup

// CreateRestoreTestingSelection is generated as a reference stub.
// Executable command wiring lives under cmd/backup.go.
//
// This request can be sent after CreateRestoreTestingPlan request returns
// successfully. This is the second part of creating a resource testing plan, and
// it must be completed sequentially.
//
// This consists of RestoreTestingSelectionName , ProtectedResourceType , and one
// of the following:
//
// - ProtectedResourceArns
//
// - ProtectedResourceConditions
//
// Each protected resource type can have one single value.
//
// A restore testing selection can include a wildcard value ("*") for
// ProtectedResourceArns along with ProtectedResourceConditions . Alternatively,
// you can include up to 30 specific protected resource ARNs in
// ProtectedResourceArns .
//
// Cannot select by both protected resource types AND specific ARNs. Request will
// fail if both are included.
