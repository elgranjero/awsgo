package cmd

func Execute(args []string) error {
	if p := _cognitosyncCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_cognitosyncCmd.Name()}, args...))
		return p.Execute()
	}
	_cognitosyncCmd.SetArgs(args)
	return _cognitosyncCmd.Execute()
}
