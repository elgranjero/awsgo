package cmd

func Execute(args []string) error {
	if p := _omicsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_omicsCmd.Name()}, args...))
		return p.Execute()
	}
	_omicsCmd.SetArgs(args)
	return _omicsCmd.Execute()
}
