package cmd

func Execute(args []string) error {
	if p := _controltowerCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_controltowerCmd.Name()}, args...))
		return p.Execute()
	}
	_controltowerCmd.SetArgs(args)
	return _controltowerCmd.Execute()
}
