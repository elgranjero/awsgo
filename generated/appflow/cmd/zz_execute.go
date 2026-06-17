package cmd

func Execute(args []string) error {
	if p := _appflowCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_appflowCmd.Name()}, args...))
		return p.Execute()
	}
	_appflowCmd.SetArgs(args)
	return _appflowCmd.Execute()
}
