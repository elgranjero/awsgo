package cmd

func Execute(args []string) error {
	if p := _kafkaCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_kafkaCmd.Name()}, args...))
		return p.Execute()
	}
	_kafkaCmd.SetArgs(args)
	return _kafkaCmd.Execute()
}
