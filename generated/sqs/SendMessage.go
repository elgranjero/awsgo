package sqs

// SendMessage is generated as a reference stub.
// Executable command wiring lives under cmd/sqs.go.
//
// Delivers a message to the specified queue.
//
// A message can include only XML, JSON, and unformatted text. The following
// Unicode characters are allowed. For more information, see the [W3C specification for characters].
//
// #x9 | #xA | #xD | #x20 to #xD7FF | #xE000 to #xFFFD | #x10000 to #x10FFFF
//
// If a message contains characters outside the allowed set, Amazon SQS rejects
// the message and returns an InvalidMessageContents error. Ensure that your
// message body includes only valid characters to avoid this exception.
//
// [W3C specification for characters]: http://www.w3.org/TR/REC-xml/#charsets
