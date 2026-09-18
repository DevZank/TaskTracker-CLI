# TaskTracker CLI

Gerenciador de tarefas via linha de comando, escrito em Go. Armazena as tarefas em um arquivo JSON local, sem depender de bibliotecas externas — só a standard library.

Projeto baseado no desafio [Task Tracker](https://roadmap.sh/projects/task-tracker) do [roadmap.sh](https://roadmap.sh).

## Funcionalidades

- Criar, atualizar e excluir tarefas
- Marcar uma tarefa como `to-do`, `in-progress` ou `done`
- Listar todas as tarefas
- Filtrar tarefas por status

## Requisitos do desafio

- [x] Adicionar, atualizar e excluir tarefas
- [x] Marcar uma tarefa como em andamento ou concluída
- [x] Listar todas as tarefas
- [x] Listar tarefas concluídas
- [x] Listar tarefas não concluídas
- [x] Listar tarefas em andamento

**Restrições seguidas:**

- [x] Linguagem: Go
- [x] Entrada do usuário via linha de comando
- [x] Persistência em arquivo JSON no diretório atual
- [x] Arquivo JSON criado automaticamente caso não exista
- [x] Uso apenas do módulo nativo de arquivos (`os`, sem libs externas)
- [x] Tratamento de erros e casos extremos

## Como rodar

```bash
git clone https://github.com/DevZank/TaskTracker-CLI.git
cd TaskTracker-CLI
go run main.go
```

O programa abre um menu interativo no terminal — não usa flags/subcomandos, as ações são escolhidas por número.

## Estrutura de uma tarefa

Cada tarefa é salva no `data.json` com os seguintes campos:

```json
{
  "id": 1,
  "description": "Estudar Go",
  "stats": "to-do",
  "createdAt": "2026-09-17T18:20:59Z",
  "updateAt": "2026-09-17T18:20:59Z"
}
```

## Sobre o desenvolvimento

Esse foi meu primeiro projeto mais "completo" em Go. A base do projeto (CRUD funcional, menu no terminal, estrutura das tarefas) eu escrevi sozinho, pesquisando pontualmente no Google conforme surgiam dúvidas específicas (ex: como funciona `slices.Delete`, como formatar strings em Go).

Depois de ter a v1 funcionando, faltava a persistência em JSON — parte em que usei o Claude como apoio pra entender melhor `encoding/json` (Marshal/Unmarshal) e revisar o código em busca de bugs. Em vez de pedir a correção pronta, pedia pra ele listar os problemas e explicar a causa, e eu mesmo implementava a solução depois.

Isso ajudou a encontrar alguns bugs ao longo das iterações, entre eles:

- **IDs duplicados após exclusão** — gerar o próximo ID com base em `len(lista)` quebra assim que você remove um item do meio, porque o tamanho da lista não corresponde mais ao maior ID já usado. Resolvido buscando o maior `ID` existente na lista, em vez de depender do tamanho dela.
- **Índice de exclusão calculado por `id - 1`** — funciona só enquanto os IDs coincidem com a posição na lista. Depois de uma primeira exclusão, isso deixa de ser verdade. A correção foi buscar o índice correto iterando e comparando o `ID` de cada item.
- **Erros de leitura/parse do JSON sendo ignorados** — várias funções logavam o erro mas continuavam a execução mesmo assim, o que podia sobrescrever o `data.json` com dados incompletos caso o arquivo estivesse corrompido. Adicionado `return` logo após cada tratamento de erro relevante.
- **`slices.Delete` mal compreendido no início** — o segundo e terceiro argumentos definem um intervalo `[i:j)` (o `j` é exclusivo), não "o índice a remover". Isso gerava confusão ao tentar remover um único item.

Achei essa abordagem (pedir só a direção do bug, sem a solução) bem mais efetiva pra realmente entender *por que* algo dava errado, em vez de só copiar um código corrigido e seguir em frente.

## Licença

Livre para uso e estudo.