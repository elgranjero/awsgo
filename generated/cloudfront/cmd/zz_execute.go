package cmd

func Execute(args []string) error {
	if p := _cloudfrontCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_cloudfrontCmd.Name()}, args...))
		return p.Execute()
	}
	_cloudfrontCmd.SetArgs(args)
	return _cloudfrontCmd.Execute()
}
