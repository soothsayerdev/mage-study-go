package exmphttp

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// "Base de dados" da taverna (simples na memoria)
var aventureiros = []Aventureiro{
	{ID: 1, Nome: "Aragorn", Classe: "Guerreiro", Level: 15, HP: 100},
	{ID: 2, Nome: "Gandalf", Classe: "Mago", Level: 50, HP: 80},
}

var proximoID = 3

func main() {
	// Criando nossa "taverna" (servidor echo)
	taverna := echo.New()

	// Middleware = "Guardas da taverna" que verificam todos que entram
	taverna.Use(middleware.Logger())  // Guarda que anota quem entra/sai
	taverna.Use(middleware.Recover()) // Guarda que evita que a taverna quebre

	// Definindo as mesas da taverna (rotas/endpoints)

	// Mesa do "Registro de Aventureiros" - Listar todos
	taverna.GET("/aventureiros", listarAventureiros)

	// Mesa do "Consultor" - Buscar aventureiro expecifico
	taverna.GET("aventureiros/:id", buscarAventureiro)

	// Mesa do "Recrutador" - Registrar novo aventureiro
	taverna.POST("/aventureiros", criarAventureiro)

	// Mesa do "Treinador" - Atualizar aventureiro existente
	taverna.PUT("/aventureiros/:id", atualizarAventureiro)

	// Mesa do "Cemiterio" - Remover aventureiro
	taverna.DELETE("/aventureiros/:id", removerAventureiro)

	// Abrindo a taverna na porta 8080
	taverna.Logger.Fatal(taverna.Start(":8080"))
}

// Handler = "NPC especializado" que atende pedidos expecificos

// NPC "Escriba" - Lista todos os aventureiros registrados
func listarAventureiros(c echo.Context) error {
	// Retorna a lista completa como se fosse um "pergaminho"
	return c.JSON(http.StatusOK, aventureiros)
}

// NPC "Oraculo" - Encontra um aventureiro especifico
func buscarAventureiro(c echo.Context) error {
	// Pega o "numero de indentificação" do aventureiro
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Se o ID não for valido, retorna "pergaminho ilegivel"
		return c.JSON(http.StatusBadRequest, map[string]string{
			"erro": "ID invalido",
		})
	}

	// Procura o aventureiro nos registros
	for _, aventureiro := range aventureiros {
		if aventureiro.ID == id {
			// Encontrou! Retorna as informações
			return c.JSON(http.StatusOK, aventureiro)
		}
	}

	// Não encontrou = "Aventureiro desaparecido"
	return c.JSON(http.StatusNotFound, map[string]string{
		"erro": "Aventureiro não encontrado",
	})
}

// NPC "Recrutador" - Registra novo aventureiro na guild
func criarAventureiro(c echo.Context) error {
	var novoAventureiro Aventureiro

	// Le o "formulario de inscrição" enviado pelo cliente
	if err := c.Bind(&novoAventureiro); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"erro": "Dados invalidos no formulario",
		})
	}

	// Atribui um novo "numero de registro"
	novoAventureiro.ID = proximoID
	proximoID++

	// Adiciona aos registros da guilda
	aventureiros = append(aventureiros, novoAventureiro)

	// Retorna confirmação com os dados do novo aventureiro
	return c.JSON(http.StatusCreated, novoAventureiro)
}

// NPC "Treinador" - Atualiza informações de aventureiro existente
func atualizarAventureiro(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"erro": "ID invalido",
		})
	}

	// Procura o aventureiro para treinar
	for i, aventureiro := range aventureiros {
		if aventureiro.ID == id {
			var dadosAtualizados Aventureiro

			// Le as "melhorias" enviadas
			if err := c.Bind(&dadosAtualizados); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{
					"erro": "Dados de atualização invalidos",
				})
			}

			// Mantem o ID original e atualiza o resto
			dadosAtualizados.ID = aventureiro.ID
			aventureiros[i] = dadosAtualizados

			return c.JSON(http.StatusOK, dadosAtualizados)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"erro": "Aventureiro não encontrado para atualização",
	})
}

// NPC "Coveiro" - Remove aventureiro dos registros
func removerAventureiro(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"erro": "ID invalido",
		})
	}

	// Procura e remove o aventureiro
	for i, aventureiro := range aventureiros {
		if aventureiro.ID == id {
			// Remove da lista (como aposentadoria forçadakkkkkkk)
			aventureiros = append(aventureiros[:i], aventureiros[i+1:]...)

			return c.JSON(http.StatusOK, map[string]string{
				"mensagem": "Aventureiro removido com sucesso",
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"erro": "Aventureiro não encontrado para remoção",
	})
}

// Componentes Principais
// Echo Framework = Motor da taverna
// Handlers = NPCs especializados
// Middleware = Guardas e sistemas de segurança
// Routes = Caminhos para diferentes serviços
// JSON = Linguagem universal de comunicação
// HTTP Status Codes = Códigos de resposta (200=sucesso, 404=não encontrado, etc.)
