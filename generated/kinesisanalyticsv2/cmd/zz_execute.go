package cmd

func Execute(args []string) error {
	if p := _kinesisanalyticsv2Cmd.Parent(); p != nil {
		p.SetArgs(append([]string{_kinesisanalyticsv2Cmd.Name()}, args...))
		return p.Execute()
	}
	_kinesisanalyticsv2Cmd.SetArgs(args)
	return _kinesisanalyticsv2Cmd.Execute()
}
