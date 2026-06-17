package cmd

func Execute(args []string) error {
	if p := _personalizeeventsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_personalizeeventsCmd.Name()}, args...))
		return p.Execute()
	}
	_personalizeeventsCmd.SetArgs(args)
	return _personalizeeventsCmd.Execute()
}
