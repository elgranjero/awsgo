package ec2

// DescribeScheduledInstanceAvailability is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Finds available schedules that meet the specified criteria.
//
// You can search for an available schedule no more than 3 months in advance. You
// must meet the minimum required duration of 1,200 hours per year. For example,
// the minimum daily schedule is 4 hours, the minimum weekly schedule is 24 hours,
// and the minimum monthly schedule is 100 hours.
//
// After you find a schedule that meets your needs, call PurchaseScheduledInstances to purchase Scheduled
// Instances with that schedule.
