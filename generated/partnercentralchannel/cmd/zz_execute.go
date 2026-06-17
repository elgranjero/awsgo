package cmd

func Execute(args []string) error {
	if p := _partnercentralchannelCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_partnercentralchannelCmd.Name()}, args...))
		return p.Execute()
	}
	_partnercentralchannelCmd.SetArgs(args)
	return _partnercentralchannelCmd.Execute()
}
