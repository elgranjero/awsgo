package mturk

// DeleteQualificationType is generated as a reference stub.
// Executable command wiring lives under cmd/mturk.go.
//
// The DeleteQualificationType deletes a Qualification type and deletes any HIT
//
// types that are associated with the Qualification type.
//
// This operation does not revoke Qualifications already assigned to Workers
// because the Qualifications might be needed for active HITs. If there are any
// pending requests for the Qualification type, Amazon Mechanical Turk rejects
// those requests. After you delete a Qualification type, you can no longer use it
// to create HITs or HIT types.
//
// DeleteQualificationType must wait for all the HITs that use the deleted
// Qualification type to be deleted before completing. It may take up to 48 hours
// before DeleteQualificationType completes and the unique name of the
// Qualification type is available for reuse with CreateQualificationType.
