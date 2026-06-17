package cmd

func Execute(args []string) error {
	if p := _ssmguiconnectCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_ssmguiconnectCmd.Name()}, args...))
		return p.Execute()
	}
	_ssmguiconnectCmd.SetArgs(args)
	return _ssmguiconnectCmd.Execute()
}
