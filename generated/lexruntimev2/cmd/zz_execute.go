package cmd

func Execute(args []string) error {
	if p := _lexruntimev2Cmd.Parent(); p != nil {
		p.SetArgs(append([]string{_lexruntimev2Cmd.Name()}, args...))
		return p.Execute()
	}
	_lexruntimev2Cmd.SetArgs(args)
	return _lexruntimev2Cmd.Execute()
}
