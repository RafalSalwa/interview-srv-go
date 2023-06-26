package repository

import (
	"github.com/RafalSalwa/interview-app-srv/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	log *logger.Logger
	db  *mongo.Client
}

func NewMongoRepository(db *mongo.Client, log *logger.Logger) *Mongo {
	return &Mongo{log: log, db: db}
}
