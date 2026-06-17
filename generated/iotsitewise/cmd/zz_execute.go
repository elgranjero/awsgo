package cmd

func Execute(args []string) error {
	if p := _iotsitewiseCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_iotsitewiseCmd.Name()}, args...))
		return p.Execute()
	}
	_iotsitewiseCmd.SetArgs(args)
	return _iotsitewiseCmd.Execute()
}
