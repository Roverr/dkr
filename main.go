package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	docker "docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"github.com/Roverr/dkr/core"
	"github.com/sirupsen/logrus"
)

func main() {
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

	systemCh := make(chan os.Signal, 1)
	signal.Notify(systemCh, os.Interrupt, syscall.SIGTERM)

	done := make(chan bool, 1)
	go func() {
		startUI(ui, manager, containers)
		done <- true
	}()
	select {
	case <-done:
		os.Exit(0)
	case <-systemCh:
		os.Exit(0)
	}
}

func startUI(ui *core.UI, manager *core.Manager, containers []types.Container) {
	argsWithoutProg := os.Args[1:]
	if len(os.Args) == 1 {
		result := ui.GetChooseMainOption()
		if result == "" {
			return
		}
		manager.RunCmd(result, "")
		targetContainer := ui.GetChooseContainer(containers)
		if targetContainer == nil {
			return
		}
		command := ui.GetCommandSelect()
		if command == "" {
			return
		}
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
