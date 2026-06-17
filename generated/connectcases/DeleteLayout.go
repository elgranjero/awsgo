package connectcases

// DeleteLayout is generated as a reference stub.
// Executable command wiring lives under cmd/connectcases.go.
//
// Deletes a layout from a cases template. You can delete up to 100 layouts per
// domain.
//
// After a layout is deleted:
//
// - You can still retrieve the layout by calling GetLayout .
//
// - You cannot update a deleted layout by calling UpdateLayout ; it throws a
// ValidationException .
//
// - Deleted layouts are not included in the ListLayouts response.
