package machinelearning

// CreateMLModel is generated as a reference stub.
// Executable command wiring lives under cmd/machinelearning.go.
//
// Creates a new MLModel using the DataSource and the recipe as information
// sources.
//
// An MLModel is nearly immutable. Users can update only the MLModelName and the
// ScoreThreshold in an MLModel without creating a new MLModel .
//
// CreateMLModel is an asynchronous operation. In response to CreateMLModel ,
// Amazon Machine Learning (Amazon ML) immediately returns and sets the MLModel
// status to PENDING . After the MLModel has been created and ready is for use,
// Amazon ML sets the status to COMPLETED .
//
// You can use the GetMLModel operation to check the progress of the MLModel
// during the creation operation.
//
// CreateMLModel requires a DataSource with computed statistics, which can be
// created by setting ComputeStatistics to true in CreateDataSourceFromRDS ,
// CreateDataSourceFromS3 , or CreateDataSourceFromRedshift operations.
