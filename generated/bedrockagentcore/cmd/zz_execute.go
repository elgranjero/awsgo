package cmd

func Execute(args []string) error {
	if p := _bedrockagentcoreCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_bedrockagentcoreCmd.Name()}, args...))
		return p.Execute()
	}
	_bedrockagentcoreCmd.SetArgs(args)
	return _bedrockagentcoreCmd.Execute()
}
