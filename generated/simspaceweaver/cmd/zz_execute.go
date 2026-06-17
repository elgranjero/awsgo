package cmd

func Execute(args []string) error {
	if p := _simspaceweaverCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_simspaceweaverCmd.Name()}, args...))
		return p.Execute()
	}
	_simspaceweaverCmd.SetArgs(args)
	return _simspaceweaverCmd.Execute()
}
