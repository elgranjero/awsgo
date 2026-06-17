package cmd

func Execute(args []string) error {
	if p := _datazoneCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_datazoneCmd.Name()}, args...))
		return p.Execute()
	}
	_datazoneCmd.SetArgs(args)
	return _datazoneCmd.Execute()
}
