package cmd

func Execute(args []string) error {
	if p := _cloudhsmv2Cmd.Parent(); p != nil {
		p.SetArgs(append([]string{_cloudhsmv2Cmd.Name()}, args...))
		return p.Execute()
	}
	_cloudhsmv2Cmd.SetArgs(args)
	return _cloudhsmv2Cmd.Execute()
}
