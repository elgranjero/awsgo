package omics

// CreateWorkflow is generated as a reference stub.
// Executable command wiring lives under cmd/omics.go.
//
// Creates a private workflow. Before you create a private workflow, you must
// create and configure these required resources:
//
// - Workflow definition file: A workflow definition file written in WDL,
// Nextflow, or CWL. The workflow definition specifies the inputs and outputs for
// runs that use the workflow. It also includes specifications for the runs and run
// tasks for your workflow, including compute and memory requirements. The workflow
// definition file must be in .zip format. For more information, see [Workflow definition files]in Amazon
// Web Services HealthOmics.
//
// - You can use Amazon Q CLI to build and validate your workflow definition
// files in WDL, Nextflow, and CWL. For more information, see [Example prompts for Amazon Q CLI]and the [Amazon Web Services HealthOmics Agentic generative AI tutorial]on GitHub.
//
// - (Optional) Parameter template file: A parameter template file written in
// JSON. Create the file to define the run parameters, or Amazon Web Services
// HealthOmics generates the parameter template for you. For more information, see [Parameter template files for HealthOmics workflows]
// .
//
// - ECR container images: Create container images for the workflow in a private
// ECR repository, or synchronize images from a supported upstream registry with
// your Amazon ECR private repository.
//
// - (Optional) Sentieon licenses: Request a Sentieon license to use the
// Sentieon software in private workflows.
//
// For more information, see [Creating or updating a private workflow in Amazon Web Services HealthOmics] in the Amazon Web Services HealthOmics User Guide.
//
// [Example prompts for Amazon Q CLI]: https://docs.aws.amazon.com/omics/latest/dev/getting-started.html#omics-q-prompts
// [Workflow definition files]: https://docs.aws.amazon.com/omics/latest/dev/workflow-definition-files.html
// [Parameter template files for HealthOmics workflows]: https://docs.aws.amazon.com/omics/latest/dev/parameter-templates.html
// [Creating or updating a private workflow in Amazon Web Services HealthOmics]: https://docs.aws.amazon.com/omics/latest/dev/creating-private-workflows.html
// [Amazon Web Services HealthOmics Agentic generative AI tutorial]: https://github.com/aws-samples/aws-healthomics-tutorials/tree/main/generative-ai
