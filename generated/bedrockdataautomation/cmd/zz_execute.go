package cmd

func Execute(args []string) error {
	if p := _bedrockdataautomationCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_bedrockdataautomationCmd.Name()}, args...))
		return p.Execute()
	}
	_bedrockdataautomationCmd.SetArgs(args)
	return _bedrockdataautomationCmd.Execute()
}
