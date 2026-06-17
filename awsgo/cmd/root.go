package cmd

import (
	"sort"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:           "aws",
		Aliases:       []string{"awsgo"},
		Short:         "Dynamic AWS-style CLI over generated SDK commands",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	outputFormat string
	awsProfile   string
	awsRegion    string
	queryExpr    string
	verbose      bool
	inputJSON    string
	inputFile    string
	outputSet    bool
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&outputFormat, "output", "", "Output format: json|yaml|text|table|csv|markdown|html")
	rootCmd.PersistentFlags().StringVar(&awsProfile, "profile", "", "AWS shared config profile")
	rootCmd.PersistentFlags().StringVar(&awsRegion, "region", "", "AWS region")
	rootCmd.PersistentFlags().StringVar(&queryExpr, "query", "", "Apply a JMESPath query to JSON output")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Print forwarded SDK command arguments and runtime diagnostics")
	rootCmd.PersistentFlags().StringVar(&inputJSON, "input-json", "", "Inline JSON object merged into SDK input shape")
	rootCmd.PersistentFlags().StringVar(&inputJSON, "cli-input-json", "", "AWS CLI compatible input JSON payload")
	rootCmd.PersistentFlags().StringVar(&inputFile, "input-file", "", "Path to JSON file merged into SDK input shape")

	serviceNames := make([]string, 0, len(serviceRegistry))
	for name := range serviceRegistry {
		serviceNames = append(serviceNames, name)
	}
	sort.Strings(serviceNames)

	for _, serviceName := range serviceNames {
		serviceMeta := serviceRegistry[serviceName]
		svc := serviceName
		operations := append([]string(nil), serviceMeta.Operations...)
		sort.Strings(operations)

		serviceCmd := &cobra.Command{
			Use:   svc,
			Short: "AWS " + svc + " operations",
		}

		for _, op := range operations {
			opName := op
			opCmd := &cobra.Command{
				Use:                opName,
				DisableFlagParsing: true,
				RunE: func(cmd *cobra.Command, args []string) error {
					return runServiceOperation(svc, opName, args)
				},
			}
			serviceCmd.AddCommand(opCmd)
		}

		rootCmd.AddCommand(serviceCmd)
	}
}
