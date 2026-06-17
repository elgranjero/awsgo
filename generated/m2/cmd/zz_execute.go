package cmd

func Execute(args []string) error {
	if p := _m2Cmd.Parent(); p != nil {
		p.SetArgs(append([]string{_m2Cmd.Name()}, args...))
		return p.Execute()
	}
	_m2Cmd.SetArgs(args)
	return _m2Cmd.Execute()
}
