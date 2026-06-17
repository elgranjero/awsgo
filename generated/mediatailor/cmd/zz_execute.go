package cmd

func Execute(args []string) error {
	if p := _mediatailorCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_mediatailorCmd.Name()}, args...))
		return p.Execute()
	}
	_mediatailorCmd.SetArgs(args)
	return _mediatailorCmd.Execute()
}
