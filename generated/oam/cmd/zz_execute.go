package cmd

func Execute(args []string) error {
	if p := _oamCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_oamCmd.Name()}, args...))
		return p.Execute()
	}
	_oamCmd.SetArgs(args)
	return _oamCmd.Execute()
}
