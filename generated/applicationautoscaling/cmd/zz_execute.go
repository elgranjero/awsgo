package cmd

func Execute(args []string) error {
	if p := _applicationautoscalingCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_applicationautoscalingCmd.Name()}, args...))
		return p.Execute()
	}
	_applicationautoscalingCmd.SetArgs(args)
	return _applicationautoscalingCmd.Execute()
}
