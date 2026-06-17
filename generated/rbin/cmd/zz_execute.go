package cmd

func Execute(args []string) error {
	if p := _rbinCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_rbinCmd.Name()}, args...))
		return p.Execute()
	}
	_rbinCmd.SetArgs(args)
	return _rbinCmd.Execute()
}
