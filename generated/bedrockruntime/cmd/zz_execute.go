package cmd

func Execute(args []string) error {
	if p := _bedrockruntimeCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_bedrockruntimeCmd.Name()}, args...))
		return p.Execute()
	}
	_bedrockruntimeCmd.SetArgs(args)
	return _bedrockruntimeCmd.Execute()
}
