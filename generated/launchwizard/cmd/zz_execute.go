package cmd

func Execute(args []string) error {
	if p := _launchwizardCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_launchwizardCmd.Name()}, args...))
		return p.Execute()
	}
	_launchwizardCmd.SetArgs(args)
	return _launchwizardCmd.Execute()
}
