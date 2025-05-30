package exmpecho

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

// Struct User - como um formulario de cadastro
type User struct {
	ID    int    `json:"id"`    // Campo ID: numero inteiro, será "id" no JSON
	Name  string `json:"name"`  // Campo Name: texto, será name no JSON
	Email string `json:"email"` // Campo Email: texto, será "email" no JSON
}

// As tags `json:...` dizem como o campo aparece quando convertido para JSON

// "Banco de dados" em memoria - como uma agenda telefonica
var users = []User{
	{ID: 1, Name: "BK", Email: "bk@gmail.com"},       // Usuario 1 pré-cadastrado
	{ID: 2, Name: "Diogo", Email: "diogo@gmail.com"}, // Usuario 2 pré-cadastrado
}

// Variavel global que simula um banco de dados (dados ficam na memoria)

func main() {
	// Criar o "garçom" (servidor Echo)
	e := echo.New()
	// echo.New() cria uma nova instancia do servidor Echo

	// Middleware - como um porteiro que registra quem entra
	e.Use(middleware.Logger())
	// Logger: registra todas as requisiçoes no console (quem acesou, quando etc)

	e.Use(middleware.Recover())
	// Recover: se o servidor "quebrar", ele se recupera automaticamente

	// Rotas - como placas de direção em um shopping
	e.GET("/users", getUsers)          // Rota GET: quando alguem acessa "/users", chama getUsers
	e.GET("/users/:id", getUser)       // Rota GET: "/users/123" chama getUser(123 vira parametro)
	e.POST("/users", createUser)       // Rota POST: para criar usuario, chama createUser
	e.PUT("/users/:id", updateUser)    // Rota PUT: para atualizar usuario, chama updateUser
	e.DELETE("/users/:id", deleteUser) // Rota DELETE: para deletar usuario, chama deleteUser

	// Iniciar servidor na porta 8080
	e.Logger.Fatal(e.Start(":8080"))
	// Start(":8080"): liga o servidor na porta 8080
	// Logger.Fatal(): se der erro ao iniciar, para o programa e mostra o erro

}

// GET /users - Como mostrar toda a agenda
func getUsers(c echo.Context) error {
	// c = contexto da requisição (como um envelope com todas as informações)
	// error = tipo de retorna (se deu certo ou errado)

	return c.JSON(http.StatusOK, users)
	// c.JSON(): converte 'users' para JSON e envia como resposta
	// http.StatusOK: codigo 200 (sucess)
	// users: nossa lista de usuarios

}

// GET /users/:id - Como procurar um contata especifico
func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// c.Param("id"): pega o valor do :id da URL (ex: /users/123 -> "123")
	// strconv.Atoi(): converte string "123" para numero 123
	// _ = ignora o erro da conversao (nao é boa pratica, mas simplifica)

	for _, user := range users {
		// for range: percorre cada usuario na lista
		// _ = ignora o indice, so queremos o usuario

		if user.ID == id {
			// Se o ID do usuario atual é igual ao ID procurado
			return c.JSON(http.StatusOK, user)
			// Retorna o usuario encontrado em JSON com codigo 200
		}
	}

	// Se chegou aqui, nao encontrou o usuario
	return c.JSON(http.StatusNotFound, map[string]string{
		"message": "Usuario não encontrado",
	})
	// Retorna erro 404 com mensagem explicativa
	// map[string]string: cria um objeto JSON {"message": "..."}
}

// POST /users - Como adicionar um novo contato
func createUser(c echo.Context) error {
	user := new(User)
	// new(User): cria um novo usuario vazio (como um formulario em branco)

	// Bind - como preencher automaticamente um formulario
	if err := c.Bind(user); err != nil {
		// c.Bind(user): pega o JSON da requisiçãi e preenche o struct user
		// if err != nil: se deu erro na conversão

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Dados invalidos",
		})
		// Retorna erro 400 (Bad Request) se os dados estão malformados
	}

	// Gerar novo ID (como um numero sequencial)
	user.ID = len(users) + 1
	// len(users): conta quantos usuarios existem
	// +1: proximo numero disponivel (se tem 2 usuarios, novo ID sera 3)

	users = append(users, *user)
	// append(): adiciona o novo usuario a lista
	// *user: desreferencia o ponteiro (pega o valor real do usuario)

	return c.JSON(http.StatusCreated, user)
	// Retorna o usuario criado com codigo 201 (Created)
}

// PUT /users/:id - Como atualizar um contato existente
func updateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// Pega o ID da URL e converte para numero

	for i, user := range users {
		// i = indice do usuario na lista (0, 1, 2...)
		// user = o usuario atual

		if user.ID == id {
			// Se encontrou o usuario com o ID procurado

			updateUser := new(User)
			// Cria um novo struct para receber os dados atualizados

			if err := c.Bind(updateUser); err != nil {
				// Tenta converter o JSON da requisição para o struct

				return c.JSON(http.StatusBadRequest, map[string]string{
					"error": "Dados invalidos",
				})
				// Se deu erro, retorna 400
			}

			updateUser.ID = id
			// Garante que o ID nao mude (mantem o ID original)

			users[i] = *updateUser
			// Substitui o usuario antigo pelo atualizado na posição i

			return c.JSON(http.StatusOK, updateUser)
			// Retorna o usuario atualizado com codigo 200
		}
	}

	// Se não encontrou o usuario
	return c.JSON(http.StatusNotFound, map[string]string{
		"message": "Usuario nao encontrado",
	})
	// Retorna erro 404
}

// DELETE /users/:id - Como remover um contato
func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// Pega o ID da URL

	for i, user := range users {
		// Percorre a lista procurando o usuario

		if user.ID == id {
			// Se encontrou o usuario

			// Remove o usuario do slice
			users = append(users[:i], users[i+1:]...)
			// users[:i]: todos elementos ANTES da posição i
			// users[i +1:]: todos os elementos DEPOIS da posição i
			// append(): junta as duas partes, "pulando" o elemento i
			// Exemplo: [A,B,C,D] removendo posição 1 -> [A] + [C,D] = [A,C,D]

			return c.JSON(http.StatusOK, map[string]string{
				"message": "Usuario deletado com sucesso",
			})
			// Retorna confirmação com codigo 200
		}
	}

	// Se nçao encontrou o usuario
	return c.JSON(http.StatusNotFound, map[string]string{
		"message": "Usuario nao encontrado",
	})
	// Retorna erro 404
}

// Resumo dos Conceitos Principais:
// Context (c): Carrega informações da requisição HTTP
// Bind(): Converte JSON da requisição para struct Go
// Param(): Extrai parâmetros da URL (:id)
// JSON(): Converte dados Go para JSON e envia como resposta
// Slice operations: Manipulação de listas (append, [:i], [i+1:])
// HTTP Status Codes: 200 (OK), 201 (Created), 400 (Bad Request), 404 (Not Found)
