package cmd

func Execute(args []string) error {
	if p := _sesv2Cmd.Parent(); p != nil {
		p.SetArgs(append([]string{_sesv2Cmd.Name()}, args...))
		return p.Execute()
	}
	_sesv2Cmd.SetArgs(args)
	return _sesv2Cmd.Execute()
}
