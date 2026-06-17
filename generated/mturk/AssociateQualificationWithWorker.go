package mturk

// AssociateQualificationWithWorker is generated as a reference stub.
// Executable command wiring lives under cmd/mturk.go.
//
// The AssociateQualificationWithWorker operation gives a Worker a Qualification.
//
// AssociateQualificationWithWorker does not require that the Worker submit a
// Qualification request. It gives the Qualification directly to the Worker.
//
// You can only assign a Qualification of a Qualification type that you created
// (using the CreateQualificationType operation).
//
// Note: AssociateQualificationWithWorker does not affect any pending
// Qualification requests for the Qualification by the Worker. If you assign a
// Qualification to a Worker, then later grant a Qualification request made by the
// Worker, the granting of the request may modify the Qualification score. To
// resolve a pending Qualification request without affecting the Qualification the
// Worker already has, reject the request with the RejectQualificationRequest
// operation.
