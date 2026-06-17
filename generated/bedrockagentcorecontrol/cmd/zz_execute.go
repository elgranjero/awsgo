package cmd

func Execute(args []string) error {
	if p := _bedrockagentcorecontrolCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_bedrockagentcorecontrolCmd.Name()}, args...))
		return p.Execute()
	}
	_bedrockagentcorecontrolCmd.SetArgs(args)
	return _bedrockagentcorecontrolCmd.Execute()
}
