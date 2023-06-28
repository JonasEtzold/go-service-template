package test

import (
	"fmt"
	"os"
	"testing"

	"gitlab.com/JonasEtzold/go-service-template/internal/pkg/config"
	"gitlab.com/JonasEtzold/go-service-template/internal/pkg/db"
	models "gitlab.com/JonasEtzold/go-service-template/internal/pkg/models/example"
	"gitlab.com/JonasEtzold/go-service-template/internal/pkg/persistence"
	"go.uber.org/zap"
)

var exampleTest models.Example

func TestMain(t *testing.M) {
	setup()
	code := t.Run()
	os.Exit(code)
}
func setup() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	config.Setup("./test.env", logger)
	db.Setup(logger)
	db.Get().Exec("DELETE FROM example")
}

func TestAddExample(t *testing.T) {
	example := models.Example{
		Name: "John",
		Text: "Doe",
	}
	s := persistence.GetExampleRepository()
	if err := s.Add(&example); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	exampleTest = example
}

func TestGetAllExamples(t *testing.T) {
	s := persistence.GetExampleRepository()
	if _, err := s.All(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestGetExampleById(t *testing.T) {
	s := persistence.GetExampleRepository()
	if _, err := s.Get(fmt.Sprint(exampleTest.ID)); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
