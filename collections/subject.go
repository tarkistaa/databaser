package collections

import (
	"github.com/tarkistaa/databaser/variables"
	"github.com/tarkistaa/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SubjectCollection struct {
	collection *mongo.Collection
}

var Subject = new(SubjectCollection)

func (c *SubjectCollection) Init() {
	c.collection = variables.Client.Database(variables.DATABASE_NAME).Collection(variables.SUBJECTS_COLLECTION)
}

func (c *SubjectCollection) Create(subject models.Subject) error {
	_, err := c.collection.InsertOne(variables.Ctx, subject)
	return err
}

func (c *SubjectCollection) Read() ([]models.Subject, error) {
	var subjects []models.Subject

	cursor, err := c.collection.Find(variables.Ctx, bson.D{})
	if err != nil {
		return subjects, err
	}

	for cursor.Next(variables.Ctx) {
		var subject models.Subject
		err := cursor.Decode(&subject)
		if err != nil {
			return subjects, err
		}

		subjects = append(subjects, subject)
	}

	if err := cursor.Err(); err != nil {
		return subjects, nil
	}

	err = cursor.Close(variables.Ctx)

	return subjects, err
}

func (c *SubjectCollection) ReadById(id primitive.ObjectID) (models.Subject, error) {
	var filter = bson.D{primitive.E{Key: "_id", Value: id}}
	var subject models.Subject

	err := c.collection.FindOne(variables.Ctx, filter).Decode(&subject)
	if err != nil {
		return subject, err
	}

	return subject, err
}

func (c *SubjectCollection) Update(subject models.Subject) error {
	filter := bson.D{{"_id", subject.Id}}
	update := bson.D{
		{"$set", bson.D{
			{"name", subject.Name},
			{"authorId", subject.AuthorId},
			{"teacherIds", subject.TeacherIds},
			{"themes", subject.Name},
			{"testIds", subject.Name},
		}},
	}

	_, err := c.collection.UpdateOne(variables.Ctx, filter, update)
	return err
}

func (c *SubjectCollection) DeleteById(id primitive.ObjectID) error {
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	_, err := c.collection.DeleteOne(variables.Ctx, filter)
	return err
}

func (c *SubjectCollection) Delete() error {
	_, err := c.collection.DeleteMany(variables.Ctx, bson.D{})
	return err
}
