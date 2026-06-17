package firehose

// DeleteDeliveryStream is generated as a reference stub.
// Executable command wiring lives under cmd/firehose.go.
//
// Deletes a Firehose stream and its data.
//
// You can delete a Firehose stream only if it is in one of the following states:
// ACTIVE , DELETING , CREATING_FAILED , or DELETING_FAILED . You can't delete a
// Firehose stream that is in the CREATING state. To check the state of a Firehose
// stream, use DescribeDeliveryStream.
//
// DeleteDeliveryStream is an asynchronous API. When an API request to
// DeleteDeliveryStream succeeds, the Firehose stream is marked for deletion, and
// it goes into the DELETING state.While the Firehose stream is in the DELETING
// state, the service might continue to accept records, but it doesn't make any
// guarantees with respect to delivering the data. Therefore, as a best practice,
// first stop any applications that are sending records before you delete a
// Firehose stream.
//
// Removal of a Firehose stream that is in the DELETING state is a low priority
// operation for the service. A stream may remain in the DELETING state for
// several minutes. Therefore, as a best practice, applications should not wait for
// streams in the DELETING state to be removed.
