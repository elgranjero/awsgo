package cmd

func Execute(args []string) error {
	if p := _deadlineCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_deadlineCmd.Name()}, args...))
		return p.Execute()
	}
	_deadlineCmd.SetArgs(args)
	return _deadlineCmd.Execute()
}
