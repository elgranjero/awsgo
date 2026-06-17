package cmd

func Execute(args []string) error {
	if p := _protonCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_protonCmd.Name()}, args...))
		return p.Execute()
	}
	_protonCmd.SetArgs(args)
	return _protonCmd.Execute()
}
