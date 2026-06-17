package machinelearning

// CreateDataSourceFromRedshift is generated as a reference stub.
// Executable command wiring lives under cmd/machinelearning.go.
//
// Creates a DataSource from a database hosted on an Amazon Redshift cluster. A
// DataSource references data that can be used to perform either CreateMLModel ,
// CreateEvaluation , or CreateBatchPrediction operations.
//
// CreateDataSourceFromRedshift is an asynchronous operation. In response to
// CreateDataSourceFromRedshift , Amazon Machine Learning (Amazon ML) immediately
// returns and sets the DataSource status to PENDING . After the DataSource is
// created and ready for use, Amazon ML sets the Status parameter to COMPLETED .
// DataSource in COMPLETED or PENDING states can be used to perform only
// CreateMLModel , CreateEvaluation , or CreateBatchPrediction operations.
//
// If Amazon ML can't accept the input source, it sets the Status parameter to
// FAILED and includes an error message in the Message attribute of the
// GetDataSource operation response.
//
// The observations should be contained in the database hosted on an Amazon
// Redshift cluster and should be specified by a SelectSqlQuery query. Amazon ML
// executes an Unload command in Amazon Redshift to transfer the result set of the
// SelectSqlQuery query to S3StagingLocation .
//
// After the DataSource has been created, it's ready for use in evaluations and
// batch predictions. If you plan to use the DataSource to train an MLModel , the
// DataSource also requires a recipe. A recipe describes how each input variable
// will be used in training an MLModel . Will the variable be included or excluded
// from training? Will the variable be manipulated; for example, will it be
// combined with another variable or will it be split apart into word combinations?
// The recipe provides answers to these questions.
//
// You can't change an existing datasource, but you can copy and modify the
// settings from an existing Amazon Redshift datasource to create a new datasource.
// To do so, call GetDataSource for an existing datasource and copy the values to
// a CreateDataSource call. Change the settings that you want to change and make
// sure that all required fields have the appropriate values.
