package cmd

func Execute(args []string) error {
	if p := _iotsecuretunnelingCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_iotsecuretunnelingCmd.Name()}, args...))
		return p.Execute()
	}
	_iotsecuretunnelingCmd.SetArgs(args)
	return _iotsecuretunnelingCmd.Execute()
}
