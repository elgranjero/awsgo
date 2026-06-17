package cmd

func Execute(args []string) error {
	if p := _ioteventsdataCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_ioteventsdataCmd.Name()}, args...))
		return p.Execute()
	}
	_ioteventsdataCmd.SetArgs(args)
	return _ioteventsdataCmd.Execute()
}
