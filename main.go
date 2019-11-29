package main

import (
	"context"
	"os"

	docker "docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"github.com/Roverr/dkr/core"
	"github.com/sirupsen/logrus"
)

func main() {
	argsWithoutProg := os.Args[1:]
	logger := logrus.New()
	cli, err := docker.NewEnvClient()
	if err != nil {
		logger.Fatal(err)
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		logger.Fatal(err)
	}
	if len(containers) == 0 {
		logger.Info("No containers running!")
		return
	}
	ui := core.NewUI(logger)
	manager := core.NewManager(logger)
	if len(os.Args) == 1 {
		result := ui.GetChooseMainOption()
		if result == "" {
			return
		}
		manager.RunCmd(result, "")
		targetContainer := ui.GetChooseContainer(containers)
		command := ui.GetCommandSelect()
		manager.RunCmd(command, targetContainer.ID)
	}
	if len(os.Args) == 2 {
		switch argsWithoutProg[0] {
		case "exec":
			targetContainer := ui.GetChooseContainer(containers)
			manager.RunCmd("exec", targetContainer.ID)
		case "logs":
			targetContainer := ui.GetChooseContainer(containers)
			manager.RunCmd("logs", targetContainer.ID)
		case "stop":
			targetContainer := ui.GetChooseContainer(containers)
			manager.RunCmd("stop", targetContainer.ID)
		}
	}
}
