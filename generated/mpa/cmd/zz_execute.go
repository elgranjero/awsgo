package cmd

func Execute(args []string) error {
	if p := _mpaCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_mpaCmd.Name()}, args...))
		return p.Execute()
	}
	_mpaCmd.SetArgs(args)
	return _mpaCmd.Execute()
}
