package cmd

func Execute(args []string) error {
	if p := _codestarnotificationsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_codestarnotificationsCmd.Name()}, args...))
		return p.Execute()
	}
	_codestarnotificationsCmd.SetArgs(args)
	return _codestarnotificationsCmd.Execute()
}
