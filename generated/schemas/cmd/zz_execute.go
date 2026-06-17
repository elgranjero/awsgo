package cmd

func Execute(args []string) error {
	if p := _schemasCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_schemasCmd.Name()}, args...))
		return p.Execute()
	}
	_schemasCmd.SetArgs(args)
	return _schemasCmd.Execute()
}
