package clouddirectory

// ListObjectParentPaths is generated as a reference stub.
// Executable command wiring lives under cmd/clouddirectory.go.
//
// Retrieves all available parent paths for any object type such as node, leaf
// node, policy node, and index node objects. For more information about objects,
// see [Directory Structure].
//
// Use this API to evaluate all parents for an object. The call returns all
// objects from the root of the directory up to the requested object. The API
// returns the number of paths based on user-defined MaxResults , in case there are
// multiple paths to the parent. The order of the paths and nodes returned is
// consistent among multiple API calls unless the objects are deleted or moved.
// Paths not leading to the directory root are ignored from the target object.
//
// [Directory Structure]: https://docs.aws.amazon.com/clouddirectory/latest/developerguide/key_concepts_directorystructure.html
