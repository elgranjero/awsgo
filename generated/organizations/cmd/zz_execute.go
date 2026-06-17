package cmd

func Execute(args []string) error {
	if p := _organizationsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_organizationsCmd.Name()}, args...))
		return p.Execute()
	}
	_organizationsCmd.SetArgs(args)
	return _organizationsCmd.Execute()
}
