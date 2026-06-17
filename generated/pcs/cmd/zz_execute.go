package cmd

func Execute(args []string) error {
	if p := _pcsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_pcsCmd.Name()}, args...))
		return p.Execute()
	}
	_pcsCmd.SetArgs(args)
	return _pcsCmd.Execute()
}
