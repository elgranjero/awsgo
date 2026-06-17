package cmd

func Execute(args []string) error {
	if p := _appconfigCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_appconfigCmd.Name()}, args...))
		return p.Execute()
	}
	_appconfigCmd.SetArgs(args)
	return _appconfigCmd.Execute()
}
