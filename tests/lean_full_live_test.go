package tests

import "testing"

func TestLeanFullLiveReadOnlyParity(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping live lean/full parity in -short mode")
	}
	if v := getenvDefault("AWSGO_LEAN_FULL_LIVE", ""); v != "1" {
		t.Skip("set AWSGO_LEAN_FULL_LIVE=1 to run live lean/full parity tests")
	}

	awsgoBin := awsgoBinPath(t)
	requireBinary(t, awsgoBin)

	cases := []parityCase{
		{Name: "sts-get-caller-identity", Args: []string{"sts", "get-caller-identity", "--output", "json"}},
		{Name: "ec2-describe-instances", Args: []string{"ec2", "describe-instances", "--max-results", "5", "--output", "json"}},
		{Name: "ec2-describe-volumes", Args: []string{"ec2", "describe-volumes", "--max-results", "5", "--output", "json"}},
		{Name: "s3-list-buckets", Args: []string{"s3", "list-buckets", "--output", "json"}},
	}

	baseEnv := map[string]string{
		"AWSGO_DISABLE_LEAN":      "",
		"AWSGO_DISABLE_PAGINATOR": "1",
		"AWS_DEFAULT_OUTPUT":      "json",
	}
	fullEnv := map[string]string{
		"AWSGO_DISABLE_LEAN":      "1",
		"AWSGO_DISABLE_PAGINATOR": "1",
		"AWS_DEFAULT_OUTPUT":      "json",
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			leanRes := runCLI(t, awsgoBin, tc.Args, baseEnv)
			fullRes := runCLI(t, awsgoBin, tc.Args, fullEnv)

			if leanRes.ExitCode != fullRes.ExitCode {
				t.Fatalf("exit mismatch for %s\nlean=%d stderr=%s\nfull=%d stderr=%s", tc.Name, leanRes.ExitCode, leanRes.Stderr, fullRes.ExitCode, fullRes.Stderr)
			}
			if leanRes.ExitCode != 0 {
				if !sameErrorClass(leanRes, fullRes) {
					t.Fatalf("error-class mismatch for %s\nlean: %s\nfull: %s", tc.Name, leanRes.Stderr, fullRes.Stderr)
				}
				return
			}

			leanJSON := canonicalJSON(t, leanRes.Stdout)
			fullJSON := canonicalJSON(t, fullRes.Stdout)
			if leanJSON != fullJSON {
				t.Fatalf("json mismatch for %s\nlean: %s\nfull: %s", tc.Name, leanJSON, fullJSON)
			}
		})
	}
}
