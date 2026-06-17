package cmd

func Execute(args []string) error {
	if p := _s3controlCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_s3controlCmd.Name()}, args...))
		return p.Execute()
	}
	_s3controlCmd.SetArgs(args)
	return _s3controlCmd.Execute()
}
