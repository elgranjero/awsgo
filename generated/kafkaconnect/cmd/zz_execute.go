package cmd

func Execute(args []string) error {
	if p := _kafkaconnectCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_kafkaconnectCmd.Name()}, args...))
		return p.Execute()
	}
	_kafkaconnectCmd.SetArgs(args)
	return _kafkaconnectCmd.Execute()
}
