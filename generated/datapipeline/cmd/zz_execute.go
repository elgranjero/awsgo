package cmd

func Execute(args []string) error {
	if p := _datapipelineCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_datapipelineCmd.Name()}, args...))
		return p.Execute()
	}
	_datapipelineCmd.SetArgs(args)
	return _datapipelineCmd.Execute()
}
