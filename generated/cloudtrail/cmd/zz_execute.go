package cmd

func Execute(args []string) error {
	if p := _cloudtrailCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_cloudtrailCmd.Name()}, args...))
		return p.Execute()
	}
	_cloudtrailCmd.SetArgs(args)
	return _cloudtrailCmd.Execute()
}
