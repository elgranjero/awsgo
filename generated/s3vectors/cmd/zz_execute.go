package cmd

func Execute(args []string) error {
	if p := _s3vectorsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_s3vectorsCmd.Name()}, args...))
		return p.Execute()
	}
	_s3vectorsCmd.SetArgs(args)
	return _s3vectorsCmd.Execute()
}
