package collections

import (
	"github.com/tarkistaa/databaser/variables"
	"github.com/tarkistaa/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuestionCollection struct {
	collection *mongo.Collection
}

var Question = new(QuestionCollection)

func (c *QuestionCollection) Init() {
	c.collection = variables.Client.Database(variables.DATABASE_NAME).Collection(variables.QUESTIONS_COLLECTION)
}

func (c *QuestionCollection) Create(question models.Question) error {
	_, err := c.collection.InsertOne(variables.Ctx, question)
	return err
}

func (c *QuestionCollection) Read() ([]models.Question, error) {
	var questions []models.Question

	cursor, err := c.collection.Find(variables.Ctx, bson.D{})
	if err != nil {
		return questions, err
	}

	for cursor.Next(variables.Ctx) {
		var question models.Question
		err := cursor.Decode(&question)
		if err != nil {
			return questions, err
		}

		questions = append(questions, question)
	}

	if err := cursor.Err(); err != nil {
		return questions, nil
	}

	err = cursor.Close(variables.Ctx)

	return questions, err
}

func (c *QuestionCollection) ReadById(id primitive.ObjectID) (models.Question, error) {
	var filter = bson.D{primitive.E{Key: "_id", Value: id}}
	var question models.Question

	err := c.collection.FindOne(variables.Ctx, filter).Decode(&question)
	if err != nil {
		return question, err
	}

	return question, err
}

func (c *QuestionCollection) Update(question models.Question) error {
	filter := bson.D{{"_id", question.Id}}
	update := bson.D{
		{"$set", bson.D{
			{"testId", question.TestId},
			{"question", question.Question},
			{"options", question.Options},
		}},
	}

	_, err := c.collection.UpdateOne(variables.Ctx, filter, update)
	return err
}

func (c *QuestionCollection) DeleteById(id primitive.ObjectID) error {
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	_, err := c.collection.DeleteOne(variables.Ctx, filter)
	return err
}

func (c *QuestionCollection) Delete() error {
	_, err := c.collection.DeleteMany(variables.Ctx, bson.D{})
	return err
}
