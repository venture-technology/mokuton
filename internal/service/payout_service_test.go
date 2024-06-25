package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/venture-technology/vtx-payout-microservice-user/config"
	"github.com/venture-technology/vtx-payout-microservice-user/internal/repository"
	"github.com/venture-technology/vtx-payout-microservice-user/models"
)

func setup(t *testing.T) (*sql.DB, *PayoutService) {

	t.Helper()

	config, err := config.Load("../../config/config.yaml")
	if err != nil {
		t.Fatalf("falha ao carregar a configuração: %v", err)
	}

	db, err := sql.Open("postgres", newPostgres(config.Database))
	if err != nil {
		t.Fatalf("falha ao conectar ao banco de dados: %v", err)
	}

	payoutRepository := repository.NewPayoutRepository(db)

	payoutService := NewPayoutService(payoutRepository)

	return db, payoutService

}

func newPostgres(dbConfig config.Database) string {
	return "user=" + dbConfig.User +
		" password=" + dbConfig.Password +
		" dbname=" + dbConfig.Name +
		" host=" + dbConfig.Host +
		" port=" + dbConfig.Port +
		" sslmode=disable"
}

func mockPayout() models.Payout {

	payoutData := map[string]interface{}{
		"amount":   100,
		"currency": "USD",
	}

	payout, err := json.Marshal(payoutData)
	if err != nil {
		panic(err)
	}

	return models.Payout{
		Payout: payout,
		Route:  "vtx-payout-mock-service/payout/test",
	}

}

func TestCreatePayout(t *testing.T) {

	db, payoutService := setup(t)
	defer db.Close()

	payoutMock := mockPayout()

	err := payoutService.CreatePayout(context.Background(), &payoutMock)

	if err != nil {
		t.Errorf("Erro ao criar payout: %v", err)
	}

}

func TestReadPayout(t *testing.T) {

	db, payoutService := setup(t)
	defer db.Close()

	var inputId string
	fmt.Println("Digite o UUID: ")
	fmt.Scanln(&inputId)

	payout, err := payoutService.ReadPayout(context.Background(), &inputId)

	if err != nil {
		t.Errorf("Erro ao fazer leitura da escola: %v", err.Error())
	}

	if reflect.TypeOf(payout) != reflect.TypeOf(models.Payout{}) {
		t.Errorf("Não foi retornado o payout: %v", err.Error())
	}

}

func TestDeletePayout(t *testing.T) {

	db, payoutService := setup(t)
	defer db.Close()

	var inputId string
	fmt.Println("Digite o UUID: ")
	fmt.Scanln(&inputId)

	err := payoutService.DeletePayout(context.Background(), &inputId)

	if err != nil {
		t.Errorf("Erro ao deletar escola: %v", err.Error())
	}

}
