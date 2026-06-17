package cmd

func Execute(args []string) error {
	if p := _cloudsearchCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_cloudsearchCmd.Name()}, args...))
		return p.Execute()
	}
	_cloudsearchCmd.SetArgs(args)
	return _cloudsearchCmd.Execute()
}
