package cmd

func Execute(args []string) error {
	if p := _frauddetectorCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_frauddetectorCmd.Name()}, args...))
		return p.Execute()
	}
	_frauddetectorCmd.SetArgs(args)
	return _frauddetectorCmd.Execute()
}
