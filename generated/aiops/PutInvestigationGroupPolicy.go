package aiops

// PutInvestigationGroupPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/aiops.go.
//
// Creates an IAM resource policy and assigns it to the specified investigation
// group.
//
// If you create your investigation group with CreateInvestigationGroup and you
// want to enable CloudWatch alarms to create investigations and add events to
// investigations, you must use this operation to create a policy similar to this
// example.
//
// { "Version": "2008-10-17", "Statement": [ { "Effect": "Allow", "Principal": {
// "Service": "aiops.alarms.cloudwatch.amazonaws.com" }, "Action": [
// "aiops:CreateInvestigation", "aiops:CreateInvestigationEvent" ], "Resource":
// "*", "Condition": { "StringEquals": { "aws:SourceAccount": "account-id" },
// "ArnLike": { "aws:SourceArn": "arn:aws:cloudwatch:region:account-id:alarm:*" } }
// } ] }
