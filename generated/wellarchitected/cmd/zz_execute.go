package cmd

func Execute(args []string) error {
	if p := _wellarchitectedCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_wellarchitectedCmd.Name()}, args...))
		return p.Execute()
	}
	_wellarchitectedCmd.SetArgs(args)
	return _wellarchitectedCmd.Execute()
}
