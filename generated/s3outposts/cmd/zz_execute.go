package cmd

func Execute(args []string) error {
	if p := _s3outpostsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_s3outpostsCmd.Name()}, args...))
		return p.Execute()
	}
	_s3outpostsCmd.SetArgs(args)
	return _s3outpostsCmd.Execute()
}
