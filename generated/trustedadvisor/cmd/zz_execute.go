package cmd

func Execute(args []string) error {
	if p := _trustedadvisorCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_trustedadvisorCmd.Name()}, args...))
		return p.Execute()
	}
	_trustedadvisorCmd.SetArgs(args)
	return _trustedadvisorCmd.Execute()
}
