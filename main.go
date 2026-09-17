package main

import (
	b "bufio"
	f "fmt"
	system "os"
	"slices"
	"strconv"
)

type Task struct {
	ID     int
	Task   string
	Stats  string
	Active bool
}

var tasks = []Task{}

func main() {
	firtTask()
	firtTask()
	firtTask()
	firtTask()
	firtTask()
	for isRunning := true; isRunning == true; {
		headQuestions()
		text := scanCli("Escolha o que quer fazer: ")
		switch text {
		case "1":
			createTask()
		case "2":
			updateTask()
		case "3":
			deleteTask()
		case "4":
			readByStatsTasks()
		}
	}
}

func headQuestions() {
	f.Println("")
	f.Println("[1] - Criar task")
	f.Println("[2] - Atualizar task")
	f.Println("[3] - Excluir task")
	f.Println("[4] - Listar tasks")
	f.Println("")
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
	defTask.ID = len(tasks) + 1
	defTask.Task = "Primeira Tarefa"
	defTask.Stats = "a fazer"
	defTask.Active = true
	tasks = append(tasks, defTask)
}

func createTask() {
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
	newTask.ID = len(tasks) + 1
	newTask.Task = taskName // + " [" + newTask.Stats + "]"
	newTask.Stats = "a fazer"
	newTask.Active = true
	tasks = append(tasks, newTask)

	f.Println("Task", newTask.Task, "criada ID:", newTask.ID)
}

func readAllTasks() {
	f.Println("")
	f.Println("---------TASKS---------")
	for _, task := range tasks {
		f.Println("[" + strconv.Itoa(task.ID) + "] " + task.Task + " [" + task.Stats + "]")
	}
	f.Println("-----------------------")
	f.Println("")
}

func readByStatsTasks() {
	f.Println("")
	f.Println("[1] - a fazer")
	f.Println("[2] - em andamento")
	f.Println("[3] - concluido")
	f.Println("[4] - todas")
	f.Println("")
	statsFilter := scanCli("Pelo o que você deseja filtar: ")
	stats := ""
	switch statsFilter {
	case "1":
		stats = "a fazer"
	case "2":
		stats = "em andamento"
	case "3":
		stats = "concluido"
	case "4":
		stats = "todas"
	}
	f.Println("")
	f.Println("---------TASKS---------")
	for _, task := range tasks {
		if task.Stats == stats {
			f.Println("[" + strconv.Itoa(task.ID) + "] " + task.Task + " [" + task.Stats + "]")
		} else if stats == "todas" {
			f.Println("[" + strconv.Itoa(task.ID) + "] " + task.Task + " [" + task.Stats + "]")
		}
	}
	f.Println("-----------------------")
	f.Println("")
}

func updateTask() {
	idTask := scanCli("Qual task deve ser atualizada ID: ")
	id, err := strconv.Atoi(idTask)

	f.Println("")
	f.Println("[1] - Nome da Task (", tasks[id-1].Task, ")")
	f.Println("[2] - Status da Task (", tasks[id-1].Stats, ")")
	f.Println("")
	option := scanCli("O que vamos alterar: ")

	switch option {
	case "1":
		newName := scanCli("Digite o novo nome para a task: ")
		tasks[id-1].Task = newName
		f.Println("Task atualizada!")
		readAllTasks()
	case "2":
		f.Println("")
		f.Println("[1] - a fazer")
		f.Println("[2] - em andamento")
		f.Println("[3] - concluido")
		f.Println("")
		newStats := scanCli("Escolha o novo Status: ")
		switch newStats {
		case "1":
			tasks[id-1].Stats = "a fazer"
		case "2":
			tasks[id-1].Stats = "em andamento"
		case "3":
			tasks[id-1].Stats = "concluido"
		}
		f.Println("Task atualizada!")
		readAllTasks()
	}

	if err != nil {
	}
}

func deleteTask() {
	readAllTasks()

	deleteID := scanCli("Escolha o item que deseja excluir pelo ID: ")
	id, err := strconv.Atoi(deleteID)
	tasks = slices.Delete(tasks, id-1, id)

	f.Println("Item excluido com sucesso!")

	readAllTasks()

	if err != nil {
	}
}