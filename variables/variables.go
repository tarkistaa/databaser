package variables

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var (
	Ctx, _       = context.WithTimeout(context.Background(), 10*time.Second)
	DatabaseHost string
	Client       *mongo.Client
)
