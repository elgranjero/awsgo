package cmd

func Execute(args []string) error {
	if p := _cloudhsmCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_cloudhsmCmd.Name()}, args...))
		return p.Execute()
	}
	_cloudhsmCmd.SetArgs(args)
	return _cloudhsmCmd.Execute()
}
