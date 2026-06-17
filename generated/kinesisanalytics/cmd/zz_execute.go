package cmd

func Execute(args []string) error {
	if p := _kinesisanalyticsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_kinesisanalyticsCmd.Name()}, args...))
		return p.Execute()
	}
	_kinesisanalyticsCmd.SetArgs(args)
	return _kinesisanalyticsCmd.Execute()
}
