package cmd

func Execute(args []string) error {
	if p := _kinesisCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_kinesisCmd.Name()}, args...))
		return p.Execute()
	}
	_kinesisCmd.SetArgs(args)
	return _kinesisCmd.Execute()
}
