package cmd

func Execute(args []string) error {
	if p := _transferCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_transferCmd.Name()}, args...))
		return p.Execute()
	}
	_transferCmd.SetArgs(args)
	return _transferCmd.Execute()
}
