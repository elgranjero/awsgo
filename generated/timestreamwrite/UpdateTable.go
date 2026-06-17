package timestreamwrite

// UpdateTable is generated as a reference stub.
// Executable command wiring lives under cmd/timestreamwrite.go.
//
// Modifies the retention duration of the memory store and magnetic store for your
// Timestream table. Note that the change in retention duration takes effect
// immediately. For example, if the retention period of the memory store was
// initially set to 2 hours and then changed to 24 hours, the memory store will be
// capable of holding 24 hours of data, but will be populated with 24 hours of data
// 22 hours after this change was made. Timestream does not retrieve data from the
// magnetic store to populate the memory store.
//
// See [code sample] for details.
//
// [code sample]: https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.update-table.html
