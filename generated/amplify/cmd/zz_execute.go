package cmd

func Execute(args []string) error {
	if p := _amplifyCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_amplifyCmd.Name()}, args...))
		return p.Execute()
	}
	_amplifyCmd.SetArgs(args)
	return _amplifyCmd.Execute()
}
