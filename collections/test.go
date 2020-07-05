package collections

import (
	"github.com/tarkistaa/databaser/variables"
	"github.com/tarkistaa/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TestCollection struct {
	collection *mongo.Collection
}

var Test = new(TestCollection)

func (c *TestCollection) Init() {
	c.collection = variables.Client.Database(variables.DATABASE_NAME).Collection(variables.TESTS_COLLECTION)
}

func (c *TestCollection) Create(test models.Test) error {
	_, err := c.collection.InsertOne(variables.Ctx, test)
	return err
}

func (c *TestCollection) Read() ([]models.Test, error) {
	var tests []models.Test

	cursor, err := c.collection.Find(variables.Ctx, bson.D{})
	if err != nil {
		return tests, err
	}

	for cursor.Next(variables.Ctx) {
		var test models.Test
		err := cursor.Decode(&test)
		if err != nil {
			return tests, err
		}

		tests = append(tests, test)
	}

	if err := cursor.Err(); err != nil {
		return tests, nil
	}

	err = cursor.Close(variables.Ctx)

	return tests, err
}

func (c *TestCollection) ReadById(id primitive.ObjectID) (models.Test, error) {
	var filter = bson.D{primitive.E{Key: "_id", Value: id}}
	var test models.Test

	err := c.collection.FindOne(variables.Ctx, filter).Decode(&test)
	if err != nil {
		return test, err
	}

	return test, err
}

func (c *TestCollection) Update(test models.Test) error {
	filter := bson.D{{"_id", test.Id}}
	update := bson.D{
		{"$set", bson.D{
			{"name", test.Name},
			{"authorId", test.AuthorId},
			{"subjectId", test.SubjectId},
			{"themeId", test.ThemeId},
			{"questionIds", test.QuestionIds},
			{"timeLimited", test.TimeLimited},
			{"timeLimit", test.TimeLimit},
		}},
	}

	_, err := c.collection.UpdateOne(variables.Ctx, filter, update)
	return err
}

func (c *TestCollection) DeleteById(id primitive.ObjectID) error {
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	_, err := c.collection.DeleteOne(variables.Ctx, filter)
	return err
}

func (c *TestCollection) Delete() error {
	_, err := c.collection.DeleteMany(variables.Ctx, bson.D{})
	return err
}
