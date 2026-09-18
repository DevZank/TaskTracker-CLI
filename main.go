package main

import (
	b "bufio"
	"encoding/json"
	f "fmt"
	"os"
	system "os"
	"slices"
	"strconv"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Stats       string    `json:"stats"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdateAt    time.Time `json:"updateAt"`
}

func main() {
	for isRunning := true; isRunning == true; {
		headQuestions()
		text := scanCli("Escolha o que quer fazer: ")
		switch text {
		case "1":
			description := scanCli("Nome da Tarefa: ")
			createTasks(description) // executa a função addTasks passando como parametro a nova lista
		case "2":
			updateTask()
		case "3":
			deleteTask()
		case "4":
			readAllTasks()
		case "5":
			readTasksByStats()
		case "6":
			isRunning = false
		default:
			f.Println("🛑  Essa opção não existe!")
		}

	}
}

func loadTasks() ([]Task, error) { // faz com que a função retorne dois valores, a lista e o possivel erro
	var tasks []Task // cria uma variavel do tipo lista estruturado em Task

	data, err := os.ReadFile("data.json") // data recebe o Json do arquivo data.json

	if err != nil { // catch de erros
		if os.IsNotExist(err) {
			return tasks, nil // Verifica se o arquivo existe, se não, ele retorna a lista
		}
		return nil, err
	}

	err = json.Unmarshal(data, &tasks) // err executa a função Unmarshal que transforma os dados do .json no tipo []Task | & na frente significa "aqui está o endereço na memória de tasks"

	if err != nil { // catch de erros
		return nil, err
	}

	return tasks, nil // retorna a lista task para a main e o possivel erro
}

func createTasks(description string) error {
	tasksList, err := loadTasks() // variavel tasks recebe a lista já existente no arquivo .json

	if err != nil { // catch de error
		f.Println("")
		f.Println("🛑  Erro ao carregar: ", err)
		return nil
	}

	dataNow := time.Now()

	biggerID := 0

	for _, task := range tasksList {
		if task.ID > biggerID {
			biggerID = task.ID
		}
	}

	newID := biggerID + 1

	newTask := Task{ // cria uma tarefa nova
		ID:          newID,
		Description: description,
		Stats:       "todo",
		CreatedAt:   dataNow,
		UpdateAt:    dataNow,
	}

	tasksList = append(tasksList, newTask) // append adiciona newTask no final da lista tasks, só que na memória, nada foi salvo em arquivo ainda.

	data, err := json.MarshalIndent(tasksList, "", " ") // pega a lista inteira (com a tarefa nova incluída) e transforma em texto JSON formatado

	if err != nil { // catch de erros
		f.Println("")
		f.Println("🛑  Ocorreu um erro:", err)
		return nil
	}

	f.Println("")
	f.Println("✅  Task [", newTask.Description, "] criada! [ ID", newTask.ID, "]")

	return os.WriteFile("data.json", data, 0644) // escreve esse texto no arquivo, substituindo tudo que tinha antes — por isso era essencial ter carregado a lista completa primeiro (loadTasks)
}

func readTasksByStats() error {
	tasksList, err := loadTasks()

	if err != nil { // catch de error
		f.Println("")
		f.Println("🛑  Erro ao carregar: ", err)
		return nil
	}

	stats := ""

	f.Println("")
	f.Println("[1] - to-do")
	f.Println("[2] - in-progress")
	f.Println("[3] - done")

	statsFilter := scanCli("Pelo o que você deseja filtrar: ")

	switch statsFilter {
	case "1":
		stats = "to-do"
	case "2":
		stats = "in-progress"
	case "3":
		stats = "done"
	default:
		f.Println("🛑  Essa opção não existe!")
	}
	f.Println("")
	f.Println("---------TASKS---------")
	for _, task := range tasksList {
		if task.Stats == stats {
			f.Println("[" + strconv.Itoa(task.ID) + "] " + task.Description + " [" + task.Stats + "]")
		}
	}
	f.Println("-----------------------")

	return nil
}

func updateTask() error {
	tasksList, err := loadTasks()

	if err != nil { // catch de error
		f.Println("")
		f.Println("🛑  Erro ao carregar: ", err)
		return nil
	}

	dataNow := time.Now()

	readAllTasks()

	idTask := scanCli("Qual task deve ser atualizada? ID: ")
	id, err := strconv.Atoi(idTask)

	if err != nil { // catch de error
		f.Println("")
		f.Println("🛑  Erro ao converter id: ", err)
		return nil
	}

	index := -1
	for i, task := range tasksList {
		if task.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		f.Println("")
		f.Println("🛑  Nenhuma task encontrada com esse ID")
		return nil
	}

	f.Println("")
	f.Println("[1] - Nome da Task (", tasksList[index].Description, ")")
	f.Println("[2] - Status da Task (", tasksList[index].Stats, ")")
	option := scanCli("Selecione o que deseja alterar: ")

	switch option {
	case "1":
		newName := scanCli("Digite o novo nome para a task: ")
		tasksList[index].Description = newName
		tasksList[index].UpdateAt = dataNow
		f.Println("✅  Task atualizada!")
	case "2":
		f.Println("")
		f.Println("[1] - to-do")
		f.Println("[2] - in-progress")
		f.Println("[3] - done")
		f.Println("")
		newStats := scanCli("Escolha o novo Status: ")
		switch newStats {
		case "1":
			tasksList[index].Stats = "to-do"
		case "2":
			tasksList[index].Stats = "in-progress"
		case "3":
			tasksList[index].Stats = "done"
		default:
			f.Println("🛑  Essa opção não existe!")
		}
		f.Println("Task atualizada!")
		tasksList[index].UpdateAt = dataNow
	default:
		f.Println("🛑  Essa opção não existe!")
	}

	data, err := json.MarshalIndent(tasksList, "", " ")

	if err != nil {
		f.Println("")
		f.Println("🛑  Ocorreu um erro:", err)
		return nil
	}

	return os.WriteFile("data.json", data, 0644)
}

func deleteTask() error {
	tasksList, err := loadTasks()

	if err != nil { // catch de error
		f.Println("")
		f.Println("🛑  Erro ao carregar: ", err)
		return nil
	}

	readAllTasks()

	deleteID := scanCli("Escolha o item que deseja excluir pelo ID: ")
	id, err := strconv.Atoi(deleteID)

	if err != nil { // catch de error
		f.Println("")
		f.Println("🛑  Erro ao converter id: ", err)
		return nil
	}

	index := -1
	for i, task := range tasksList {
		if task.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		f.Println("")
		f.Println("🛑  Nenhuma task encontrada com esse ID")
		return nil
	}

	tasksList = slices.Delete(tasksList, index, index+1)

	f.Println("✅  Item excluido com sucesso!")

	if err != nil {
		f.Println("")
		f.Println("🛑  Ocorreu um erro: ", err)
		return nil
	}

	data, err := json.MarshalIndent(tasksList, "", " ") // pega a lista inteira (com a tarefa nova incluída) e transforma em texto JSON formatado

	if err != nil { // catch de erros
		f.Println("")
		f.Println("🛑  Ocorreu um erro:", err)
		return nil
	}

	return os.WriteFile("data.json", data, 0644)
}

func readAllTasks() error {
	tasksList, err := loadTasks()

	if err != nil { // catch de error
		f.Println("")
		f.Println("🛑  Erro ao carregar: ", err)
		return nil
	}

	f.Println("")
	f.Println("---------TASKS---------")
	for _, task := range tasksList {
		f.Println("[" + strconv.Itoa(task.ID) + "] " + task.Description + " [" + task.Stats + "]")
	}
	f.Println("-----------------------")

	return nil
}

func scanCli(message string) string {
	text := ""
	scanner := b.NewScanner(system.Stdin)
	f.Println("")
	f.Print("↪️   ", message)

	if scanner.Scan() {
		text = scanner.Text()
	}

	return text
}

func headQuestions() {
	f.Println("")
	f.Println("[1] - Criar task")
	f.Println("[2] - Atualizar task")
	f.Println("[3] - Excluir task")
	f.Println("[4] - Listar todas as tasks")
	f.Println("[5] - Filtrar e listar as tasks")
	f.Println("[6] - Sair")
}
