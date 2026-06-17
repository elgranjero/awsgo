package cmd

func Execute(args []string) error {
	if p := _mediaconvertCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_mediaconvertCmd.Name()}, args...))
		return p.Execute()
	}
	_mediaconvertCmd.SetArgs(args)
	return _mediaconvertCmd.Execute()
}
