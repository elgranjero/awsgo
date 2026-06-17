package tests

import (
	"os"
	"testing"
)

type parityCase struct {
	Name string
	Args []string
}

func TestParityLiveReadOnlyHighValue(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping live parity in -short mode")
	}
	if v := getenvDefault("AWSGO_PARITY_LIVE", ""); v != "1" {
		t.Skip("set AWSGO_PARITY_LIVE=1 to run live parity tests")
	}

	awsBin := awsBinPath()
	requireBinary(t, awsBin)
	awsgoBin := awsgoBinPath(t)
	requireBinary(t, awsgoBin)

	cases := []parityCase{
		{Name: "sts-get-caller-identity", Args: []string{"sts", "get-caller-identity", "--output", "json"}},
		{Name: "iam-get-account-summary", Args: []string{"iam", "get-account-summary", "--output", "json"}},
		{Name: "iam-list-account-aliases", Args: []string{"iam", "list-account-aliases", "--output", "json"}},
		{Name: "ec2-describe-regions", Args: []string{"ec2", "describe-regions", "--all-regions", "--output", "json"}},
		{Name: "s3-list-buckets", Args: []string{"s3", "list-buckets", "--output", "json"}},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			awsRes := runCLI(t, awsBin, tc.Args, nil)
			awsgoRes := runCLI(t, awsgoBin, tc.Args, nil)

			if awsRes.ExitCode != awsgoRes.ExitCode {
				t.Fatalf("exit mismatch for %s\naws=%d stderr=%s\nawsgo=%d stderr=%s", tc.Name, awsRes.ExitCode, awsRes.Stderr, awsgoRes.ExitCode, awsgoRes.Stderr)
			}
			if awsRes.ExitCode != 0 {
				if !sameErrorClass(awsRes, awsgoRes) {
					t.Fatalf("error-class mismatch for %s\naws: %s\nawsgo: %s", tc.Name, awsRes.Stderr, awsgoRes.Stderr)
				}
				return
			}

			awsJSON := canonicalJSON(t, awsRes.Stdout)
			awsgoJSON := canonicalJSON(t, awsgoRes.Stdout)
			if awsJSON != awsgoJSON {
				t.Fatalf("json mismatch for %s\naws:   %s\nawsgo: %s", tc.Name, awsJSON, awsgoJSON)
			}
		})
	}
}

func getenvDefault(k, def string) string {
	v := def
	if x, ok := os.LookupEnv(k); ok {
		v = x
	}
	return v
}
