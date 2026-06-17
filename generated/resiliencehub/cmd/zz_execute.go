package cmd

func Execute(args []string) error {
	if p := _resiliencehubCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_resiliencehubCmd.Name()}, args...))
		return p.Execute()
	}
	_resiliencehubCmd.SetArgs(args)
	return _resiliencehubCmd.Execute()
}
