package cmd

func Execute(args []string) error {
	if p := _greengrassv2Cmd.Parent(); p != nil {
		p.SetArgs(append([]string{_greengrassv2Cmd.Name()}, args...))
		return p.Execute()
	}
	_greengrassv2Cmd.SetArgs(args)
	return _greengrassv2Cmd.Execute()
}
