package cmd

func Execute(args []string) error {
	if p := _wafCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_wafCmd.Name()}, args...))
		return p.Execute()
	}
	_wafCmd.SetArgs(args)
	return _wafCmd.Execute()
}
