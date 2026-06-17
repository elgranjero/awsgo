package tests

import (
	"strings"
	"testing"
)

func TestParityHighValueServicesPresentInManifest(t *testing.T) {
	services := loadManifestV2(t)
	seen := map[string]bool{}
	for _, svc := range services {
		seen[svc.Name] = true
	}

	highValue := []string{
		"acm", "autoscaling", "cloudtrail", "cloudwatch", "configservice",
		"costexplorer", "dynamodb", "ec2", "ecr", "ecs", "efs",
		"eventbridge", "guardduty", "iam", "kms", "lambda",
		"rds", "route53", "s3", "secretsmanager", "securityhub",
		"sns", "sqs", "ssm", "sts", "wafv2",
	}

	for _, svc := range highValue {
		if !seen[svc] {
			t.Fatalf("high-value service %q missing from awsgo manifest", svc)
		}
	}
}

func TestParityHelpFlagNamesAcronymsAndRequiredOptional(t *testing.T) {
	bin := awsgoBinPath(t)
	requireBinary(t, bin)

	out := runCLI(t, bin, []string{"s3", "upload-part", "--help"}, nil)
	if out.ExitCode != 0 {
		t.Fatalf("%s failed: %s", commandString(bin, []string{"s3", "upload-part", "--help"}), out.Stderr)
	}
	mustContainAll(t, out.Stdout, []string{
		"--checksum-crc32",
		"--checksum-crc32c",
		"--checksum-crc64nvme",
		"--checksum-sha1",
		"--checksum-sha256",
		"--content-md5",
		"--sse-customer-key-md5",
	})

	out2 := runCLI(t, bin, []string{"wafv2", "create-web-acl", "--help"}, nil)
	if out2.ExitCode != 0 {
		t.Fatalf("%s failed: %s", commandString(bin, []string{"wafv2", "create-web-acl", "--help"}), out2.Stderr)
	}
	mustContainAll(t, out2.Stdout, []string{
		"--on-source-ddos-protection-config",
		"Required",
		"Optional",
	})
}

func TestParityHelpHasOperationFlagsSection(t *testing.T) {
	bin := awsgoBinPath(t)
	requireBinary(t, bin)

	out := runCLI(t, bin, []string{"iam", "list-roles", "--help"}, nil)
	if out.ExitCode != 0 {
		t.Fatalf("%s failed: %s", commandString(bin, []string{"iam", "list-roles", "--help"}), out.Stderr)
	}
	if !strings.Contains(out.Stdout, "Operation Flags:") {
		t.Fatalf("expected Operation Flags section in help output\n%s", out.Stdout)
	}
	if !strings.Contains(out.Stdout, "--max-items") {
		t.Fatalf("expected --max-items operation flag in help output\n%s", out.Stdout)
	}
}
