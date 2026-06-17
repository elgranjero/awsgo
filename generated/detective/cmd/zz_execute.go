package cmd

func Execute(args []string) error {
	if p := _detectiveCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_detectiveCmd.Name()}, args...))
		return p.Execute()
	}
	_detectiveCmd.SetArgs(args)
	return _detectiveCmd.Execute()
}
