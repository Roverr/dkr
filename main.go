package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	docker "docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"github.com/manifoldco/promptui"
)

func main() {
	argsWithoutProg := os.Args[1:]
	for _, arg := range argsWithoutProg {
		fmt.Println(arg)
	}
	cli, err := docker.NewEnvClient()
	if err != nil {
		log.Fatal(err)
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) == 1 {
		targetContainer := getUserToChooseContainer(containers)

		prompt := promptui.Select{
			Label: "Select Commnad",
			Items: []string{"exec", "logs", "stop"},
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		runCommand(result, targetContainer.ID)
	}
	if len(os.Args) == 2 {
		switch argsWithoutProg[0] {
		case "exec":
			targetContainer := getUserToChooseContainer(containers)
			runCommand("exec", targetContainer.ID)
		case "logs":
			targetContainer := getUserToChooseContainer(containers)
			runCommand("logs", targetContainer.ID)
		case "stop":
			targetContainer := getUserToChooseContainer(containers)
			runCommand("stop", targetContainer.ID)
		}
	}
}

func getUserToChooseContainer(containers []types.Container) *types.Container {
	containerMap := map[string]*types.Container{}
	items := []string{}
	for _, container := range containers {
		item := fmt.Sprintf(
			"%s  %s  %s  %s",
			container.ID[:12],
			container.Image,
			container.Status,
			container.State,
		)
		containerMap[item] = &container
		items = append(items, item)
	}
	prompt := promptui.Select{
		Label: "Select Container",
		Items: items,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil
	}
	return containerMap[result]
}

func commandForOS(commands ...string) *exec.Cmd {
	cmd := exec.Command(commands[0], commands[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd
}

func runCommand(command, containerID string) {
	switch command {
	case "exec":
		options := []string{
			"/bin/bash",
			"/bin/zsh",
			"/bin/sh",
		}
		for _, bin := range options {
			err := commandForOS("docker", "exec", "-it", containerID, bin).Run()
			if err == nil {
				break
			}
		}
	case "logs":
		cmd := exec.Command("docker", "logs", containerID)
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	case "stop":
		cmd := exec.Command("docker", "stop", containerID)
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}
}
