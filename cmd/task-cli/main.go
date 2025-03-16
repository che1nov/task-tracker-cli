package main

import (
	"os"
	"task-tracker/intrernal/cli"
	"task-tracker/intrernal/repository"
	"task-tracker/intrernal/usecase"
)

func main() {
	repo := repository.NewTaskRepository("tasks.json")
	usecase := usecase.NewTaskUsecase(*repo)
	cli := cli.NewCLI(usecase)

	cli.Run(os.Args)
}
