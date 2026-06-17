package machinelearning

// CreateDataSourceFromRDS is generated as a reference stub.
// Executable command wiring lives under cmd/machinelearning.go.
//
// Creates a DataSource object from an [Amazon Relational Database Service] (Amazon RDS). A DataSource references data
// that can be used to perform CreateMLModel , CreateEvaluation , or
// CreateBatchPrediction operations.
//
// CreateDataSourceFromRDS is an asynchronous operation. In response to
// CreateDataSourceFromRDS , Amazon Machine Learning (Amazon ML) immediately
// returns and sets the DataSource status to PENDING . After the DataSource is
// created and ready for use, Amazon ML sets the Status parameter to COMPLETED .
// DataSource in the COMPLETED or PENDING state can be used only to perform
// >CreateMLModel >, CreateEvaluation , or CreateBatchPrediction operations.
//
// If Amazon ML cannot accept the input source, it sets the Status parameter to
// FAILED and includes an error message in the Message attribute of the
// GetDataSource operation response.
//
// [Amazon Relational Database Service]: http://aws.amazon.com/rds/
