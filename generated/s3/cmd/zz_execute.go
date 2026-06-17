package cmd

func Execute(args []string) error {
	if p := _s3Cmd.Parent(); p != nil {
		p.SetArgs(append([]string{_s3Cmd.Name()}, args...))
		return p.Execute()
	}
	_s3Cmd.SetArgs(args)
	return _s3Cmd.Execute()
}
