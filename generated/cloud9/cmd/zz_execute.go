package cmd

func Execute(args []string) error {
	if p := _cloud9Cmd.Parent(); p != nil {
		p.SetArgs(append([]string{_cloud9Cmd.Name()}, args...))
		return p.Execute()
	}
	_cloud9Cmd.SetArgs(args)
	return _cloud9Cmd.Execute()
}
