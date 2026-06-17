package cmd

func Execute(args []string) error {
	if p := _ecrCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_ecrCmd.Name()}, args...))
		return p.Execute()
	}
	_ecrCmd.SetArgs(args)
	return _ecrCmd.Execute()
}
