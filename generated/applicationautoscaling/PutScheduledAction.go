package applicationautoscaling

// PutScheduledAction is generated as a reference stub.
// Executable command wiring lives under cmd/applicationautoscaling.go.
//
// Creates or updates a scheduled action for an Application Auto Scaling scalable
// target.
//
// Each scalable target is identified by a service namespace, resource ID, and
// scalable dimension. A scheduled action applies to the scalable target identified
// by those three attributes. You cannot create a scheduled action until you have
// registered the resource as a scalable target.
//
// When you specify start and end times with a recurring schedule using a cron
// expression or rates, they form the boundaries for when the recurring action
// starts and stops.
//
// To update a scheduled action, specify the parameters that you want to change.
// If you don't specify start and end times, the old values are deleted.
//
// For more information, see [Scheduled scaling] in the Application Auto Scaling User Guide.
//
// If a scalable target is deregistered, the scalable target is no longer
// available to run scheduled actions. Any scheduled actions that were specified
// for the scalable target are deleted.
//
// [Scheduled scaling]: https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-scheduled-scaling.html
