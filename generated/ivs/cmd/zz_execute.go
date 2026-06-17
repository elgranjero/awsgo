package cmd

func Execute(args []string) error {
	if p := _ivsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_ivsCmd.Name()}, args...))
		return p.Execute()
	}
	_ivsCmd.SetArgs(args)
	return _ivsCmd.Execute()
}
