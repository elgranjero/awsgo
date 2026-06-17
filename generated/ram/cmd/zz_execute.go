package cmd

func Execute(args []string) error {
	if p := _ramCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_ramCmd.Name()}, args...))
		return p.Execute()
	}
	_ramCmd.SetArgs(args)
	return _ramCmd.Execute()
}
