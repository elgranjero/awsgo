package cmd

func Execute(args []string) error {
	if p := _firehoseCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_firehoseCmd.Name()}, args...))
		return p.Execute()
	}
	_firehoseCmd.SetArgs(args)
	return _firehoseCmd.Execute()
}
