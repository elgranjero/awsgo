package waf

// CreateWebACLMigrationStack is generated as a reference stub.
// Executable command wiring lives under cmd/waf.go.
//
// Creates an AWS CloudFormation WAFV2 template for the specified web ACL in the
// specified Amazon S3 bucket. Then, in CloudFormation, you create a stack from the
// template, to create the web ACL and its resources in AWS WAFV2. Use this to
// migrate your AWS WAF Classic web ACL to the latest version of AWS WAF.
//
// This is part of a larger migration procedure for web ACLs from AWS WAF Classic
// to the latest version of AWS WAF. For the full procedure, including caveats and
// manual steps to complete the migration and switch over to the new web ACL, see [Migrating your AWS WAF Classic resources to AWS WAF]
// in the [AWS WAF Developer Guide].
//
// [Migrating your AWS WAF Classic resources to AWS WAF]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-migrating-from-classic.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
