package cmd

func Execute(args []string) error {
	if p := _neptuneCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_neptuneCmd.Name()}, args...))
		return p.Execute()
	}
	_neptuneCmd.SetArgs(args)
	return _neptuneCmd.Execute()
}
