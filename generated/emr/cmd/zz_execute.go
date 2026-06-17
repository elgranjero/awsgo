package cmd

func Execute(args []string) error {
	if p := _emrCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_emrCmd.Name()}, args...))
		return p.Execute()
	}
	_emrCmd.SetArgs(args)
	return _emrCmd.Execute()
}
