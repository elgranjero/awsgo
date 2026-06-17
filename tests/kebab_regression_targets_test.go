package tests

import (
	"strings"
	"testing"
)

type helpRegressionCase struct {
	Service        string
	Operation      string
	ExpectedFlags  []string
	ForbiddenFlags []string
}

func TestKebabRegressionTargets(t *testing.T) {
	bin := awsgoBinPath(t)
	requireBinary(t, bin)

	cases := []helpRegressionCase{
		{
			Service:   "cognitoidentityprovider",
			Operation: "create-user-pool-client",
			ExpectedFlags: []string{
				"--allowed-oauth-flows",
				"--allowed-oauth-flows-user-pool-client",
				"--callback-urls",
				"--logout-urls",
			},
			ForbiddenFlags: []string{
				"--allowed-o-auth-flows",
				"--callback-ur-ls",
				"--logout-ur-ls",
			},
		},
		{
			Service:   "lexmodelsv2",
			Operation: "create-intent",
			ExpectedFlags: []string{
				"--qna-intent-configuration",
			},
			ForbiddenFlags: []string{
				"--qn-a-intent-configuration",
			},
		},
		{
			Service:   "location",
			Operation: "search-place-index-for-text",
			ExpectedFlags: []string{
				"--filter-bbox",
			},
			ForbiddenFlags: []string{
				"--filter-b-box",
			},
		},
		{
			Service:   "opensearch",
			Operation: "list-instance-type-details",
			ExpectedFlags: []string{
				"--retrieve-azs",
			},
			ForbiddenFlags: []string{
				"--retrieve-a-zs",
			},
		},
		{
			Service:   "odb",
			Operation: "create-cloud-vm-cluster",
			ExpectedFlags: []string{
				"--data-storage-size-in-tbs",
				"--db-node-storage-size-in-gbs",
				"--memory-size-in-gbs",
			},
			ForbiddenFlags: []string{
				"--data-storage-size-in-t-bs",
				"--db-node-storage-size-in-g-bs",
				"--memory-size-in-g-bs",
			},
		},
		{
			Service:   "quicksight",
			Operation: "predict-qa-results",
			ExpectedFlags: []string{
				"--include-quicksight-q-index",
			},
			ForbiddenFlags: []string{
				"--include-quick-sight-q-index",
				"--include-quick-sight-qindex",
				"--include-quicksight-qindex",
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Service+"_"+tc.Operation, func(t *testing.T) {
			out := runCLI(t, bin, []string{tc.Service, tc.Operation, "--help"}, nil)
			if out.ExitCode != 0 {
				t.Fatalf("%s failed: %s", commandString(bin, []string{tc.Service, tc.Operation, "--help"}), out.Stderr)
			}
			mustContainAll(t, out.Stdout, tc.ExpectedFlags)
			for _, bad := range tc.ForbiddenFlags {
				if strings.Contains(out.Stdout, bad) {
					t.Fatalf("found forbidden flag %q in help output\n%s", bad, out.Stdout)
				}
			}
		})
	}
}
