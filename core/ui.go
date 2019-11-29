package core

import (
	"docker.io/go-docker/api/types"
	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
)

// UI represents the UI interactions
type UI struct {
	logger *logrus.Logger
}

// NewUI creates a new instance of UI
func NewUI(logger *logrus.Logger) *UI {
	return &UI{logger}
}

type displayedContainer struct {
	ShortID       string
	ImageName     string
	CurrentStatus string
	State         string
}

// GetChooseContainer prompts a new select with the containers
func (ui *UI) GetChooseContainer(containers []types.Container) *types.Container {
	items := []displayedContainer{}
	templates := &promptui.SelectTemplates{
		Active:   `{{ printf ">" | blue }} {{ .ShortID | bold }} ({{ .ImageName | bold }})`,
		Inactive: "  {{ .ShortID }} ({{ .ImageName }})",
		Selected: " {{ .ShortID | bold }} {{ .ImageName | bold | blue }} selected",
		Details: `
		--------- Container ----------
		{{ "ID:" | faint }}	{{ .ShortID }}
		{{ "Image:" | faint }}	{{ .ImageName }}
		{{ "Status:" | faint }}	{{ .CurrentStatus }}
		{{ "State:" | faint }}	{{ .State }}`,
	}
	for _, container := range containers {
		item := displayedContainer{
			container.ID[:12],
			container.Image,
			container.Status,
			container.State,
		}
		items = append(items, item)
	}
	prompt := promptui.Select{
		Label:     "Select Container",
		Items:     items,
		Templates: templates,
	}

	index, _, err := prompt.Run()
	if err != nil {
		ui.logger.Errorf("Prompt failed %v", err)
		return nil
	}
	return &containers[index]
}

// GetCommandSelect shows a new command select
func (ui *UI) GetCommandSelect() string {
	prompt := promptui.Select{
		Label: "Select Commnad",
		Items: []string{"exec", "logs", "stop"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		ui.logger.Errorf("Select containers dialog failed | %s", err)
	}
	return result
}
