package cmd

func Execute(args []string) error {
	if p := _mediapackagev2Cmd.Parent(); p != nil {
		p.SetArgs(append([]string{_mediapackagev2Cmd.Name()}, args...))
		return p.Execute()
	}
	_mediapackagev2Cmd.SetArgs(args)
	return _mediapackagev2Cmd.Execute()
}
