package cmd

func Execute(args []string) error {
	if p := _aiopsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_aiopsCmd.Name()}, args...))
		return p.Execute()
	}
	_aiopsCmd.SetArgs(args)
	return _aiopsCmd.Execute()
}
