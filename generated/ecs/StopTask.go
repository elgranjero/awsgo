package ecs

// StopTask is generated as a reference stub.
// Executable command wiring lives under cmd/ecs.go.
//
// Stops a running task. Any tags associated with the task will be deleted.
//
// When you call StopTask on a task, the equivalent of docker stop is issued to
// the containers running in the task. This results in a stop signal value and a
// default 30-second timeout, after which the SIGKILL value is sent and the
// containers are forcibly stopped. This signal can be defined in your container
// image with the STOPSIGNAL instruction and will default to SIGTERM . If the
// container handles the SIGTERM value gracefully and exits within 30 seconds from
// receiving it, no SIGKILL value is sent.
//
// For Windows containers, POSIX signals do not work and runtime stops the
// container by sending a CTRL_SHUTDOWN_EVENT . For more information, see [Unable to react to graceful shutdown of (Windows) container #25982] on
// GitHub.
//
// The default 30-second timeout can be configured on the Amazon ECS container
// agent with the ECS_CONTAINER_STOP_TIMEOUT variable. For more information, see [Amazon ECS Container Agent Configuration]
// in the Amazon Elastic Container Service Developer Guide.
//
// [Unable to react to graceful shutdown of (Windows) container #25982]: https://github.com/moby/moby/issues/25982
// [Amazon ECS Container Agent Configuration]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-config.html
