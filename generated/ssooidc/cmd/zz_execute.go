package cmd

func Execute(args []string) error {
	if p := _ssooidcCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_ssooidcCmd.Name()}, args...))
		return p.Execute()
	}
	_ssooidcCmd.SetArgs(args)
	return _ssooidcCmd.Execute()
}
