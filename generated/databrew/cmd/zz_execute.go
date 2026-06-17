package cmd

func Execute(args []string) error {
	if p := _databrewCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_databrewCmd.Name()}, args...))
		return p.Execute()
	}
	_databrewCmd.SetArgs(args)
	return _databrewCmd.Execute()
}
