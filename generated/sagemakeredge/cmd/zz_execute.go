package cmd

func Execute(args []string) error {
	if p := _sagemakeredgeCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_sagemakeredgeCmd.Name()}, args...))
		return p.Execute()
	}
	_sagemakeredgeCmd.SetArgs(args)
	return _sagemakeredgeCmd.Execute()
}
