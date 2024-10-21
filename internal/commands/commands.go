package commands

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/rcarvalho-pb/cli-todo/internal/models"
)

type CmdFlags struct {
	Add          string
	Del          int
	Edit         string
	Toggle       int
	List         bool
	ListArgument string
}

func NewCmdFlags() *CmdFlags {
	cmdFlags := CmdFlags{}

	flag.StringVar(&cmdFlags.Add, "add", "", "add a new task")
	flag.StringVar(&cmdFlags.Edit, "edit", "", "edit a task title by the task id. id:new_title")
	flag.IntVar(&cmdFlags.Del, "del", 0, "Specify a task by id to be deleted")
	flag.IntVar(&cmdFlags.Toggle, "toggle", 0, "Specify a task by id to be toggled")
	flag.BoolVar(&cmdFlags.List, "list", false, "list all unfinished tasks | -a list all tasks | -f list all finished tasks")

	listCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd.StringVar(&cmdFlags.ListArgument, "f", "", "")
	listCmd.StringVar(&cmdFlags.ListArgument, "a", "", "")

	flag.Parse()

	return &cmdFlags
}

func (cf *CmdFlags) Execute(task *models.Task) {
	switch {
	case cf.List:
		fmt.Println(os.Args)

	case cf.Add != "":
		fmt.Println(task.AddTask(cf.Add))

	case cf.Edit != "":
		parts := strings.Split(cf.Edit, ":")
		id, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(task.UpdateTask(id, parts[1]))

	case cf.Del > 0:
		fmt.Println(task.DeleteTaks(int64(cf.Del)))

	case cf.Toggle > 0:
		fmt.Println(task.ToggleTask(int64(cf.Toggle)))

	default:
		fmt.Println("Unknown command")
	}
}

func listAll(task *models.Task) {
	tasks, err := task.GetAllTasks()
	if err != nil {
		log.Fatal("Error getting all tasks")
	}

	task.Print(tasks)
}

func listAllFinished(task *models.Task) {
	tasks, err := task.GetAllFinishedTasks()
	if err != nil {
		log.Fatal("Error getting all unfinished tasks")
	}

	task.Print(tasks)
}

func listAllUnfinished(task *models.Task) {
	tasks, err := task.GetUnfinishedTasks()
	if err != nil {
		log.Fatal("Error getting all unfinished tasks")
	}

	task.Print(tasks)
}
