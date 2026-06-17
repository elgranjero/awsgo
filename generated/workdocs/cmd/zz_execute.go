package cmd

func Execute(args []string) error {
	if p := _workdocsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_workdocsCmd.Name()}, args...))
		return p.Execute()
	}
	_workdocsCmd.SetArgs(args)
	return _workdocsCmd.Execute()
}
