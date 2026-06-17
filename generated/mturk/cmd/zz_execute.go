package cmd

func Execute(args []string) error {
	if p := _mturkCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_mturkCmd.Name()}, args...))
		return p.Execute()
	}
	_mturkCmd.SetArgs(args)
	return _mturkCmd.Execute()
}
