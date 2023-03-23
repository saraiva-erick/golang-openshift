package main

import (
    "fmt"
    "net/http"
    "time"

    "github.com/labstack/echo/v4"
)

//By: Erick Saraiva

func main() {
    // cria uma nova instância do servidor Echo
    e := echo.New()

    // define a rota para a página inicial
    e.GET("/", func(c echo.Context) error {
        // Obter a data e hora atual
        now := time.Now()

        // Calcular quantos dias restam até o final do ano
        daysLeft := time.Date(now.Year(), 12, 31, 0, 0, 0, 0, time.UTC).YearDay() - now.YearDay()

        // cria uma string com a resposta
        response := fmt.Sprintf("Hoje é %s, %s. Já passaram %d dias neste ano e restam %d dias até o final do ano.",
            now.Weekday().String(), now.Format("02/01/2006"), now.YearDay(), daysLeft)

        // retorna a resposta como uma resposta HTTP
        return c.String(http.StatusOK, response)
    })

    // inicia o servidor na porta 8080
    e.Logger.Fatal(e.Start(":8080"))
}