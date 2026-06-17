package cmd

func Execute(args []string) error {
	if p := _customerprofilesCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_customerprofilesCmd.Name()}, args...))
		return p.Execute()
	}
	_customerprofilesCmd.SetArgs(args)
	return _customerprofilesCmd.Execute()
}
