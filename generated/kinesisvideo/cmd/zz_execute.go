package cmd

func Execute(args []string) error {
	if p := _kinesisvideoCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_kinesisvideoCmd.Name()}, args...))
		return p.Execute()
	}
	_kinesisvideoCmd.SetArgs(args)
	return _kinesisvideoCmd.Execute()
}
