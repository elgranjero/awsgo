package cmd

func Execute(args []string) error {
	if p := _wafv2Cmd.Parent(); p != nil {
		p.SetArgs(append([]string{_wafv2Cmd.Name()}, args...))
		return p.Execute()
	}
	_wafv2Cmd.SetArgs(args)
	return _wafv2Cmd.Execute()
}
