package fly

import (
	"log"
	"os"
	"os/exec"
	"sort"

	"github.com/AlecAivazis/survey/v2"
)

type Answer struct {
	ServerName string
}

type Fly struct {
	source *ConfigReader
}

func (fly *Fly) GetQuestions() []*survey.Question {
	var options []string
	serverList := fly.source.GetServerList()
	sort.Slice(serverList, func(i int, j int) bool { return serverList[i].Name < serverList[j].Name })
	for _, server := range serverList {
		options = append(options, server.Name)
	}

	return []*survey.Question{
		{
			Name: "serverName",
			Prompt: &survey.Select{
				Message:  "Choose a server: ",
				Options:  options,
				Default:  options[0],
				PageSize: 10,
				Filter:   FuzzySearch,
			},
		},
	}
}

func (fly *Fly) GetCommand(answer *Answer) *exec.Cmd {
	server := fly.source.GetServer(answer.ServerName)
	command := exec.Command("sshpass", "-p", server.Password, "ssh", server.UserName+"@"+server.Host)
	return command
}

func (fly *Fly) RepairCommand(answer *Answer) *exec.Cmd {
	server := fly.source.GetServer(answer.ServerName)
	command := exec.Command("ssh", server.UserName+"@"+server.Host)
	return command
}

func (fly *Fly) Ask() *Answer {
	answer := Answer{}
	if err := survey.Ask(fly.GetQuestions(), &answer); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return &answer
}

func (fly *Fly) Run(answer *Answer) {
	command := fly.GetCommand(answer)
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	if err := command.Run(); err != nil {
		log.Fatal(err)
	}
}

func NewFly() *Fly {
	return &Fly{source: NewConfigReader()}
}
