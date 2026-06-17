package cmd

func Execute(args []string) error {
	if p := _schedulerCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_schedulerCmd.Name()}, args...))
		return p.Execute()
	}
	_schedulerCmd.SetArgs(args)
	return _schedulerCmd.Execute()
}
