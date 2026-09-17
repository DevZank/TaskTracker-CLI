package main

import (
	b "bufio"
	f "fmt"
	system "os"
	"strconv"
)

type Task struct {
	ID     int
	Task   string
	Stats  string
	Active bool
}

var tasks = []Task{}
var ID = 1

func main() {
	for isRunning := true; isRunning == true; {
		headQuestions()
		scanCli("Escolha o que quer fazer: ")
		isRunning = false
	}

	// firtTask()
	// createTask()
	// readTasks()
}

func headQuestions() {
	f.Println("[1] - Criar task")
	f.Println("[2] - Atualizar task")
	f.Println("[3] - Excluir task")
	f.Println("[4] - Listar tasks")
}

func scanCli(message string) string {
	text := ""
	scanner := b.NewScanner(system.Stdin)
	f.Print(message)

	if scanner.Scan() {
		text = scanner.Text()
	}

	return text
}

func firtTask() {
	defTask := Task{}
	defTask.ID = ID
	defTask.Task = "Primeira Tarefa"
	defTask.Stats = "a fazer"
	defTask.Active = true
	tasks = append(tasks, defTask)
}

func createTask() {
	ID++
	taskName := ""

	// Cria um scanner que lê da entrada padrão (teclado)
	scanner := b.NewScanner(system.Stdin)
	f.Print("Digite o nome da Task: ")

	// Lê a linha inteira até o Enter
	if scanner.Scan() {
		taskName = scanner.Text()
	}

	// f.Scanln(&taskReader)

	newTask := Task{}
	newTask.ID = ID
	newTask.Task = taskName // + " [" + newTask.Stats + "]"
	newTask.Stats = "a fazer"
	newTask.Active = true
	tasks = append(tasks, newTask)

	f.Println("Task", newTask.Task, "criada ID:", newTask.ID)
}

func readTasks() {
	f.Println("")
	f.Println("------")
	for _, task := range tasks {
		f.Println("[" + strconv.Itoa(task.ID) + "] " + task.Task + " [" + task.Stats + "]")
	}
}

func updateTask() {
	tasks[1].Task = "Atualizado"
}
