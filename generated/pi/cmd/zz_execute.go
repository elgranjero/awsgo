package cmd

func Execute(args []string) error {
	if p := _piCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_piCmd.Name()}, args...))
		return p.Execute()
	}
	_piCmd.SetArgs(args)
	return _piCmd.Execute()
}
