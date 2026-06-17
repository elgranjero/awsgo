package cmd

func Execute(args []string) error {
	if p := _observabilityadminCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_observabilityadminCmd.Name()}, args...))
		return p.Execute()
	}
	_observabilityadminCmd.SetArgs(args)
	return _observabilityadminCmd.Execute()
}
