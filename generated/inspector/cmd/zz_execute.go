package cmd

func Execute(args []string) error {
	if p := _inspectorCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_inspectorCmd.Name()}, args...))
		return p.Execute()
	}
	_inspectorCmd.SetArgs(args)
	return _inspectorCmd.Execute()
}
