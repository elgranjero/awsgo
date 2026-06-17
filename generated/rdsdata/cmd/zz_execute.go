package cmd

func Execute(args []string) error {
	if p := _rdsdataCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_rdsdataCmd.Name()}, args...))
		return p.Execute()
	}
	_rdsdataCmd.SetArgs(args)
	return _rdsdataCmd.Execute()
}
