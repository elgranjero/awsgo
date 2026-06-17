package cmd

func Execute(args []string) error {
	if p := _workspacesthinclientCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_workspacesthinclientCmd.Name()}, args...))
		return p.Execute()
	}
	_workspacesthinclientCmd.SetArgs(args)
	return _workspacesthinclientCmd.Execute()
}
