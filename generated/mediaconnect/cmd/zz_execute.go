package cmd

func Execute(args []string) error {
	if p := _mediaconnectCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_mediaconnectCmd.Name()}, args...))
		return p.Execute()
	}
	_mediaconnectCmd.SetArgs(args)
	return _mediaconnectCmd.Execute()
}
