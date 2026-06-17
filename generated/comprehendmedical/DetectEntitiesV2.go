package comprehendmedical

// DetectEntitiesV2 is generated as a reference stub.
// Executable command wiring lives under cmd/comprehendmedical.go.
//
// Inspects the clinical text for a variety of medical entities and returns
// specific information about them such as entity category, location, and
// confidence score on that information. Amazon Comprehend Medical only detects
// medical entities in English language texts.
//
// The DetectEntitiesV2 operation replaces the DetectEntities operation. This new action uses a
// different model for determining the entities in your medical text and changes
// the way that some entities are returned in the output. You should use the
// DetectEntitiesV2 operation in all new applications.
//
// The DetectEntitiesV2 operation returns the Acuity and Direction entities as
// attributes instead of types.
