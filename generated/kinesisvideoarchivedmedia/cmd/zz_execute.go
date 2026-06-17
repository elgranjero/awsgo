package cmd

func Execute(args []string) error {
	if p := _kinesisvideoarchivedmediaCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_kinesisvideoarchivedmediaCmd.Name()}, args...))
		return p.Execute()
	}
	_kinesisvideoarchivedmediaCmd.SetArgs(args)
	return _kinesisvideoarchivedmediaCmd.Execute()
}
