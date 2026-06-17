package cmd

func Execute(args []string) error {
	if p := _macie2Cmd.Parent(); p != nil {
		p.SetArgs(append([]string{_macie2Cmd.Name()}, args...))
		return p.Execute()
	}
	_macie2Cmd.SetArgs(args)
	return _macie2Cmd.Execute()
}
