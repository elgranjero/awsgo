package cmd

func Execute(args []string) error {
	if p := _inspector2Cmd.Parent(); p != nil {
		p.SetArgs(append([]string{_inspector2Cmd.Name()}, args...))
		return p.Execute()
	}
	_inspector2Cmd.SetArgs(args)
	return _inspector2Cmd.Execute()
}
