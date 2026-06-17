package cmd

func Execute(args []string) error {
	if p := _bedrockagentCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_bedrockagentCmd.Name()}, args...))
		return p.Execute()
	}
	_bedrockagentCmd.SetArgs(args)
	return _bedrockagentCmd.Execute()
}
