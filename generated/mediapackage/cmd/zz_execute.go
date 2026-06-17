package cmd

func Execute(args []string) error {
	if p := _mediapackageCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_mediapackageCmd.Name()}, args...))
		return p.Execute()
	}
	_mediapackageCmd.SetArgs(args)
	return _mediapackageCmd.Execute()
}
