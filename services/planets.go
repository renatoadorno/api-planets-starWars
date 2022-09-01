package services

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"planets.api/models"
)

type PlanetInterface interface {
	Insert(models.Planets) (models.Planets, error)
	Get(string) (models.Planets, error)
	GetByName(string) (models.Planets, error)
  Delete(name string) error
}

type PlanetClient struct {
	ctx  context.Context
	coll *mongo.Collection
}

func NewClient(coll *mongo.Collection, ctx context.Context) *PlanetClient {
	return &PlanetClient{
		ctx:  ctx,
		coll: coll,
	}
}

func (c *PlanetClient) Insert(planet models.Planets) (models.Planets, error) {
	planets := models.Planets{}

	res, err := c.coll.InsertOne(c.ctx, planet)

	if err != nil {
		return planets, err
	}

	id := res.InsertedID.(primitive.ObjectID).Hex()
	return c.Get(id)
}

func (c *PlanetClient) Get(id string) (models.Planets, error) {
	planet := models.Planets{}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return planet, err
	}

	err = c.coll.FindOne(c.ctx, bson.M{"_id": _id}).Decode(&planet)

	if err != nil {
		return planet, err
	}

	return planet, nil
}

func (c *PlanetClient) GetByName(name string) (models.Planets, error) {
	var planet models.Planets

	err := c.coll.FindOne(c.ctx, bson.M{"name": name}).Decode(planet)

  if err != nil {
		return planet, err
	}

	return planet, err
}

func (c *PlanetClient) Delete(name string) error {
	filter := bson.D{primitive.E{Key: "name", Value: name}}
	result, _ := c.coll.DeleteOne(c.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no planet found for delete")
	}
	return nil
}
