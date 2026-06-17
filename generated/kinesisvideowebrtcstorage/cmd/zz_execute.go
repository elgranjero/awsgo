package cmd

func Execute(args []string) error {
	if p := _kinesisvideowebrtcstorageCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_kinesisvideowebrtcstorageCmd.Name()}, args...))
		return p.Execute()
	}
	_kinesisvideowebrtcstorageCmd.SetArgs(args)
	return _kinesisvideowebrtcstorageCmd.Execute()
}
