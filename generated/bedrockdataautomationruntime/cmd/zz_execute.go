package cmd

func Execute(args []string) error {
	if p := _bedrockdataautomationruntimeCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_bedrockdataautomationruntimeCmd.Name()}, args...))
		return p.Execute()
	}
	_bedrockdataautomationruntimeCmd.SetArgs(args)
	return _bedrockdataautomationruntimeCmd.Execute()
}
