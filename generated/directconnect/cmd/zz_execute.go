package cmd

func Execute(args []string) error {
	if p := _directconnectCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_directconnectCmd.Name()}, args...))
		return p.Execute()
	}
	_directconnectCmd.SetArgs(args)
	return _directconnectCmd.Execute()
}
