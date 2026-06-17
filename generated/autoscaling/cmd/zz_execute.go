package cmd

func Execute(args []string) error {
	if p := _autoscalingCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_autoscalingCmd.Name()}, args...))
		return p.Execute()
	}
	_autoscalingCmd.SetArgs(args)
	return _autoscalingCmd.Execute()
}
