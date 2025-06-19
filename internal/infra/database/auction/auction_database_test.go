package auction

import (
	"context"
	"fullcycle-auction_go/configuration/database/mongodb"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestIntegration_CreateAuction_ShouldUpdateStatusToCompleted(t *testing.T) {
	t.Setenv("MONGO_INITDB_ROOT_USERNAME", "admin")
	t.Setenv("MONGO_INITDB_ROOT_PASSWORD", "admin")
	t.Setenv("MONGODB_URL", "mongodb://admin:admin@localhost:27017/auctions?authSource=admin")
	t.Setenv("MONGODB_DB", "test_auction_db_"+uuid.New().String())
	t.Setenv("AUCTION_INTERVAL", "1s")

	ctx := context.Background()

	databaseConnection, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer databaseConnection.Client().Disconnect(ctx)

	// Adia a remoção de todo o banco de dados de teste para limpeza
	defer func() {
		if err := databaseConnection.Drop(ctx); err != nil {
			t.Fatalf("Falha ao remover o banco de dados de teste: %v", err)
		}
	}()
	// --- Fim da Configuração ---

	auctionRepo := NewAuctionRepository(databaseConnection)

	// Assume que `Active` (ou outro status inicial) e `New` (para Condition) são valores válidos.
	// O teste precisa de um status inicial para verificar a mudança.
	auction := &auction_entity.Auction{
		Id:          uuid.New().String(),
		ProductName: "Test Product",
		Category:    "Test Category",
		Description: "Test Description",
		Condition:   auction_entity.ProductCondition(0), // Ex: New
		Status:      auction_entity.AuctionStatus(1),    // Ex: Active
		Timestamp:   time.Now(),
	}

	// Chama o método que está sendo testado
	internalErr := auctionRepo.CreateAuction(ctx, auction)
	assert.Nil(t, internalErr)

	// Aguarda um tempo maior que o AUCTION_INTERVAL para a goroutine atualizar o status
	time.Sleep(2 * time.Second)

	// Busca o leilão no banco de dados para verificar seu status
	var resultAuction AuctionEntityMongo
	filter := bson.M{"_id": auction.Id}
	err = auctionRepo.Collection.FindOne(ctx, filter).Decode(&resultAuction)

	// Asserções
	assert.Nil(t, err, "Deveria encontrar o leilão criado")
	assert.Equal(t, auction_entity.Completed, resultAuction.Status, "O status do leilão deveria ser atualizado para Completed")
}
