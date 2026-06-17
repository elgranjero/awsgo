package cmd

func Execute(args []string) error {
	if p := _kinesisvideomediaCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_kinesisvideomediaCmd.Name()}, args...))
		return p.Execute()
	}
	_kinesisvideomediaCmd.SetArgs(args)
	return _kinesisvideomediaCmd.Execute()
}
