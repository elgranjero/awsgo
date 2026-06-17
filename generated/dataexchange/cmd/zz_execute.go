package cmd

func Execute(args []string) error {
	if p := _dataexchangeCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_dataexchangeCmd.Name()}, args...))
		return p.Execute()
	}
	_dataexchangeCmd.SetArgs(args)
	return _dataexchangeCmd.Execute()
}
