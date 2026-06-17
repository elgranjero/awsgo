package cmd

func Execute(args []string) error {
	if p := _qbusinessCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_qbusinessCmd.Name()}, args...))
		return p.Execute()
	}
	_qbusinessCmd.SetArgs(args)
	return _qbusinessCmd.Execute()
}
