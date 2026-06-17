package cmd

func Execute(args []string) error {
	if p := _lightsailCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_lightsailCmd.Name()}, args...))
		return p.Execute()
	}
	_lightsailCmd.SetArgs(args)
	return _lightsailCmd.Execute()
}
