package cmd

func Execute(args []string) error {
	if p := _wafregionalCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_wafregionalCmd.Name()}, args...))
		return p.Execute()
	}
	_wafregionalCmd.SetArgs(args)
	return _wafregionalCmd.Execute()
}
