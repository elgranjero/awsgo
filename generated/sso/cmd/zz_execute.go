package cmd

func Execute(args []string) error {
	if p := _ssoCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_ssoCmd.Name()}, args...))
		return p.Execute()
	}
	_ssoCmd.SetArgs(args)
	return _ssoCmd.Execute()
}
