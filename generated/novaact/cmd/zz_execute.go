package cmd

func Execute(args []string) error {
	if p := _novaactCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_novaactCmd.Name()}, args...))
		return p.Execute()
	}
	_novaactCmd.SetArgs(args)
	return _novaactCmd.Execute()
}
