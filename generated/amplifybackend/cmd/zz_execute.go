package cmd

func Execute(args []string) error {
	if p := _amplifybackendCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_amplifybackendCmd.Name()}, args...))
		return p.Execute()
	}
	_amplifybackendCmd.SetArgs(args)
	return _amplifybackendCmd.Execute()
}
