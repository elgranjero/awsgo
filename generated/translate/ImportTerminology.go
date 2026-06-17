package translate

// ImportTerminology is generated as a reference stub.
// Executable command wiring lives under cmd/translate.go.
//
// Creates or updates a custom terminology, depending on whether one already
// exists for the given terminology name. Importing a terminology with the same
// name as an existing one will merge the terminologies based on the chosen merge
// strategy. The only supported merge strategy is OVERWRITE, where the imported
// terminology overwrites the existing terminology of the same name.
//
// If you import a terminology that overwrites an existing one, the new
// terminology takes up to 10 minutes to fully propagate. After that, translations
// have access to the new terminology.
