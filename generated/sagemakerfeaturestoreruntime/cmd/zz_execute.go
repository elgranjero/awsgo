package cmd

func Execute(args []string) error {
	if p := _sagemakerfeaturestoreruntimeCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_sagemakerfeaturestoreruntimeCmd.Name()}, args...))
		return p.Execute()
	}
	_sagemakerfeaturestoreruntimeCmd.SetArgs(args)
	return _sagemakerfeaturestoreruntimeCmd.Execute()
}
