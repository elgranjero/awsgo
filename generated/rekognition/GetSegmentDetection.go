package rekognition

// GetSegmentDetection is generated as a reference stub.
// Executable command wiring lives under cmd/rekognition.go.
//
// Gets the segment detection results of a Amazon Rekognition Video analysis
// started by StartSegmentDetection.
//
// Segment detection with Amazon Rekognition Video is an asynchronous operation.
// You start segment detection by calling StartSegmentDetectionwhich returns a job identifier ( JobId ).
// When the segment detection operation finishes, Amazon Rekognition publishes a
// completion status to the Amazon Simple Notification Service topic registered in
// the initial call to StartSegmentDetection . To get the results of the segment
// detection operation, first check that the status value published to the Amazon
// SNS topic is SUCCEEDED . if so, call GetSegmentDetection and pass the job
// identifier ( JobId ) from the initial call of StartSegmentDetection .
//
// GetSegmentDetection returns detected segments in an array ( Segments ) of SegmentDetection
// objects. Segments is sorted by the segment types specified in the SegmentTypes
// input parameter of StartSegmentDetection . Each element of the array includes
// the detected segment, the precentage confidence in the acuracy of the detected
// segment, the type of the segment, and the frame in which the segment was
// detected.
//
// Use SelectedSegmentTypes to find out the type of segment detection requested in
// the call to StartSegmentDetection .
//
// Use the MaxResults parameter to limit the number of segment detections
// returned. If there are more results than specified in MaxResults , the value of
// NextToken in the operation response contains a pagination token for getting the
// next set of results. To get the next page of results, call GetSegmentDetection
// and populate the NextToken request parameter with the token value returned from
// the previous call to GetSegmentDetection .
//
// For more information, see Detecting video segments in stored video in the
// Amazon Rekognition Developer Guide.
