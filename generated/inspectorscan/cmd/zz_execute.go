package cmd

func Execute(args []string) error {
	if p := _inspectorscanCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_inspectorscanCmd.Name()}, args...))
		return p.Execute()
	}
	_inspectorscanCmd.SetArgs(args)
	return _inspectorscanCmd.Execute()
}
