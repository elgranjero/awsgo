package cmd

func Execute(args []string) error {
	if p := _lexmodelsv2Cmd.Parent(); p != nil {
		p.SetArgs(append([]string{_lexmodelsv2Cmd.Name()}, args...))
		return p.Execute()
	}
	_lexmodelsv2Cmd.SetArgs(args)
	return _lexmodelsv2Cmd.Execute()
}
