package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/manifoldco/promptui"
)

type App struct {
	Name        string
	Environment string
	Version     string
	GitTag      string
}

func main() {
	shaList := [][32]byte{}
	for i := 0; i < 5; i++ {
		shaList = append(shaList, sha256.Sum256([]byte(fmt.Sprint(time.Now().Unix()+int64(i)))))
	}

	apps := []App{
		{Name: "app1", Environment: "dev", Version: "1.3.0", GitTag: fmt.Sprintf("%x", shaList[0])},
		{Name: "app1", Environment: "sandbox", Version: "1.2.0", GitTag: fmt.Sprintf("%x", shaList[1])},
		{Name: "app1", Environment: "prod", Version: "1.1.0", GitTag: fmt.Sprintf("%x", shaList[2])},
		{Name: "app2", Environment: "dev", Version: "1.1.5", GitTag: fmt.Sprintf("%x", shaList[3])},
		{Name: "app2", Environment: "dev", Version: "1.1.5", GitTag: fmt.Sprintf("%x", shaList[4])},
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U000025B6 {{ .Name | cyan }} {{ .Version | white }} ({{ .Environment | red }})",
		Inactive: "  {{ .Name | cyan }} {{ .Version | white }} ({{ .Environment | red }})",
		Selected: "\U000025B6 {{ .Name | red | cyan }}",
		Details: `
--------- Deployment details ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Environment:" | faint }}	{{ .Environment }}
{{ "Version:" | faint }}	{{ .Version }}
{{ "GitTag:" | faint }}	{{ .GitTag }}`,
	}

	searcher := func(input string, index int) bool {
		app := apps[index]
		name := strings.Replace(strings.ToLower(app.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Select an application to deploy:",
		Items:     apps,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	confirmPrompt := promptui.Prompt{
		Label:     "Are you sure you want to deploy " + apps[i].Name + ", version " + apps[i].Version + " to " + apps[i].Environment + "?",
		IsConfirm: true,
	}

	_, err = confirmPrompt.Run()
	if err != nil {
		log.Fatal("You canceled the deployment")
	}

	log.Info("Running deployment...")
	time.Sleep(2 * time.Second)

	log.Info("You deployed ", apps[i].Name, " to ", apps[i].Environment, " at tag ", apps[i].GitTag)
}
