package cmd

func Execute(args []string) error {
	if p := _emrcontainersCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_emrcontainersCmd.Name()}, args...))
		return p.Execute()
	}
	_emrcontainersCmd.SetArgs(args)
	return _emrcontainersCmd.Execute()
}
