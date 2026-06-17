package cmd

func Execute(args []string) error {
	if p := _medialiveCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_medialiveCmd.Name()}, args...))
		return p.Execute()
	}
	_medialiveCmd.SetArgs(args)
	return _medialiveCmd.Execute()
}
