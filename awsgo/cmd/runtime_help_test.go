package cmd

import (
	"strings"
	"testing"
)

func TestKnownInputFlagHelpShowsOperationFlags(t *testing.T) {
	lines := knownInputFlagHelp("iam", "list-roles")
	if len(lines) == 0 {
		t.Fatal("expected operation flags for iam list-roles")
	}
	joined := strings.Join(lines, "\n")
	if !strings.Contains(joined, "--max-items") {
		t.Fatalf("expected --max-items in help output, got: %s", joined)
	}
}

func TestKnownInputFlagHelpShowsRequiredAndOptional(t *testing.T) {
	const svcName = "__help_test_service__"
	serviceRegistry[svcName] = serviceDef{
		OperationInputs: map[string][]string{
			"op": {"Name", "Marker"},
		},
		OperationInputTypes: map[string]map[string]string{
			"op": {
				"Name":   "*string",
				"Marker": "*string",
			},
		},
		OperationInputRequired: map[string][]string{
			"op": {"Name"},
		},
	}
	defer delete(serviceRegistry, svcName)

	lines := knownInputFlagHelp(svcName, "op")
	joined := strings.Join(lines, "\n")
	if !strings.Contains(joined, "--name") || !strings.Contains(joined, "Required") {
		t.Fatalf("expected required marker for --name, got: %s", joined)
	}
	if !strings.Contains(joined, "--marker") || !strings.Contains(joined, "Optional") {
		t.Fatalf("expected optional marker for --marker, got: %s", joined)
	}
}

func TestConsumeOperationFlagsSkipsReservedNames(t *testing.T) {
	svc := serviceDef{
		OperationInputTypes: map[string]map[string]string{
			"op": {
				"Region":   "*string",
				"MaxItems": "*int32",
			},
		},
	}
	overrides, kept, err := consumeOperationFlags(svc, "op", []string{"--region", "us-east-1", "--max-items", "2"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(overrides) != 1 {
		t.Fatalf("expected exactly one override, got %#v", overrides)
	}
	if _, ok := overrides["MaxItems"]; !ok {
		t.Fatalf("expected MaxItems override, got %#v", overrides)
	}
	if len(kept) != 2 || kept[0] != "--region" || kept[1] != "us-east-1" {
		t.Fatalf("expected reserved --region to be kept, got %#v", kept)
	}
}

func TestToKebabAcronymAndDigits(t *testing.T) {
	cases := map[string]string{
		"ChecksumCRC32":                   "checksum-crc32",
		"ChecksumCRC32C":                  "checksum-crc32c",
		"ChecksumCRC64NVME":               "checksum-crc64nvme",
		"ChecksumSHA1":                    "checksum-sha1",
		"ChecksumSHA256":                  "checksum-sha256",
		"ContentMD5":                      "content-md5",
		"SSECustomerKeyMD5":               "sse-customer-key-md5",
		"OnSourceDDoSProtectionConfig":    "on-source-ddos-protection-config",
		"AssignIpv6Addresses":             "assign-ipv6-addresses",
		"AllowedOAuthFlows":               "allowed-oauth-flows",
		"CallbackURLs":                    "callback-urls",
		"LogoutURLs":                      "logout-urls",
		"TopicARNs":                       "topic-arns",
		"IncludeQuickSightQIndex":         "include-quicksight-q-index",
		"QnAIntentConfiguration":          "qna-intent-configuration",
		"FilterBBox":                      "filter-bbox",
		"RetrieveAZs":                     "retrieve-azs",
		"DataStorageSizeInTBs":            "data-storage-size-in-tbs",
		"MemoryPerOracleComputeUnitInGBs": "memory-per-oracle-compute-unit-in-gbs",
	}
	for in, want := range cases {
		if got := toKebab(in); got != want {
			t.Fatalf("%s => %s, want %s", in, got, want)
		}
	}
}

func TestKnownInputFlagHelpAlignment(t *testing.T) {
	const svcName = "__help_align_service__"
	serviceRegistry[svcName] = serviceDef{
		OperationInputs: map[string][]string{
			"op": {"Name", "OnSourceDDoSProtectionConfig", "Scope"},
		},
		OperationInputTypes: map[string]map[string]string{
			"op": {
				"Name":                         "*string",
				"OnSourceDDoSProtectionConfig": "*types.OnSourceDDoSProtectionConfig",
				"Scope":                        "types.Scope",
			},
		},
		OperationInputRequired: map[string][]string{
			"op": {"Name", "Scope"},
		},
	}
	defer delete(serviceRegistry, svcName)

	lines := knownInputFlagHelp(svcName, "op")
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines, got %d: %#v", len(lines), lines)
	}
	reqCol := -1
	for _, ln := range lines {
		iOpt := strings.Index(ln, "  Optional")
		iReq := strings.Index(ln, "  Required")
		i := iOpt
		if iReq >= 0 {
			i = iReq
		}
		if i < 0 {
			t.Fatalf("missing required/optional marker: %q", ln)
		}
		if reqCol == -1 {
			reqCol = i
		}
		if i != reqCol {
			t.Fatalf("misaligned required/optional columns: %q", lines)
		}
	}
}
