package cmd

func Execute(args []string) error {
	if p := _bedrockCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_bedrockCmd.Name()}, args...))
		return p.Execute()
	}
	_bedrockCmd.SetArgs(args)
	return _bedrockCmd.Execute()
}
