package cmd

func Execute(args []string) error {
	if p := _autoscalingplansCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_autoscalingplansCmd.Name()}, args...))
		return p.Execute()
	}
	_autoscalingplansCmd.SetArgs(args)
	return _autoscalingplansCmd.Execute()
}
