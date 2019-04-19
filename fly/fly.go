package fly

import (
	"gopkg.in/AlecAivazis/survey.v1"
	"log"
	"os"
	"os/exec"
)

type Answer struct {
	ServerName string
}

type Fly struct {
	DataBase *DB
}

func (fly *Fly) GetQuestions() []*survey.Question {
	var options []string
	serverList := fly.DataBase.GetServerList()
	for _, server := range serverList {
		options = append(options, server.Name)
	}

	return []*survey.Question {
		{
			Name: "serverName",
			Prompt: &survey.Select{
				Message: "Choose a server",
				Options: options,
				Default: options[0],
			},
		},
	}
}

func (fly *Fly) GetCommand(answer *Answer) *exec.Cmd {
	server := fly.DataBase.GetServer(answer.ServerName)
	// log.Printf("sshpass -p %s ssh %s@%s\n", server.Password, server.UserName, server.Host)
	command := exec.Command("sshpass", "-p", server.Password, "ssh", server.UserName +"@" + server.Host)
	return command
}

func (fly *Fly) Fly() {
	fly.DataBase.Start()
}

func (fly *Fly) Crash() {
	fly.DataBase.close()
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
	log.Print(answer.ServerName)
	command := fly.GetCommand(answer)
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	_ = command.Run()
}

func (fly *Fly) UpdateServer(server Server) {
	otherServer := fly.DataBase.getDefault()
	server.Merge(&otherServer)
	fly.DataBase.UpdateServer(server)
}

func (fly *Fly) UpdateDefault(server Server) {
	fly.DataBase.UpdateDefault(server)
}

func NewFly() *Fly {
	return &Fly{DataBase: NewDB()}
}


