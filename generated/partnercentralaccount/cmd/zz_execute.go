package cmd

func Execute(args []string) error {
	if p := _partnercentralaccountCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_partnercentralaccountCmd.Name()}, args...))
		return p.Execute()
	}
	_partnercentralaccountCmd.SetArgs(args)
	return _partnercentralaccountCmd.Execute()
}
