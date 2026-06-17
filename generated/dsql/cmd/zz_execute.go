package cmd

func Execute(args []string) error {
	if p := _dsqlCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_dsqlCmd.Name()}, args...))
		return p.Execute()
	}
	_dsqlCmd.SetArgs(args)
	return _dsqlCmd.Execute()
}
