package bedrockagent

// CreateKnowledgeBase is generated as a reference stub.
// Executable command wiring lives under cmd/bedrockagent.go.
//
// Creates a knowledge base. A knowledge base contains your data sources so that
// Large Language Models (LLMs) can use your data. To create a knowledge base, you
// must first set up your data sources and configure a supported vector store. For
// more information, see [Set up a knowledge base].
//
// If you prefer to let Amazon Bedrock create and manage a vector store for you in
// Amazon OpenSearch Service, use the console. For more information, see [Create a knowledge base].
//
// - Provide the name and an optional description .
//
// - Provide the Amazon Resource Name (ARN) with permissions to create a
// knowledge base in the roleArn field.
//
// - Provide the embedding model to use in the embeddingModelArn field in the
// knowledgeBaseConfiguration object.
//
// - Provide the configuration for your vector store in the storageConfiguration
// object.
//
// - For an Amazon OpenSearch Service database, use the
// opensearchServerlessConfiguration object. For more information, see [Create a vector store in Amazon OpenSearch Service].
//
// - For an Amazon Aurora database, use the RdsConfiguration object. For more
// information, see [Create a vector store in Amazon Aurora].
//
// - For a Pinecone database, use the pineconeConfiguration object. For more
// information, see [Create a vector store in Pinecone].
//
// - For a Redis Enterprise Cloud database, use the
// redisEnterpriseCloudConfiguration object. For more information, see [Create a vector store in Redis Enterprise Cloud].
//
// [Create a knowledge base]: https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-create
// [Create a vector store in Amazon OpenSearch Service]: https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-setup-oss.html
// [Create a vector store in Redis Enterprise Cloud]: https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-setup-redis.html
// [Set up a knowledge base]: https://docs.aws.amazon.com/bedrock/latest/userguide/knowlege-base-prereq.html
// [Create a vector store in Amazon Aurora]: https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-setup-rds.html
// [Create a vector store in Pinecone]: https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-setup-pinecone.html
