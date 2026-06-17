package cmd

func Execute(args []string) error {
	if p := _mediapackagevodCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_mediapackagevodCmd.Name()}, args...))
		return p.Execute()
	}
	_mediapackagevodCmd.SetArgs(args)
	return _mediapackagevodCmd.Execute()
}
