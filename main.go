package databaser

import (
	"github.com/tarkistaa/databaser/collections"
	"github.com/tarkistaa/databaser/variables"
	"github.com/tarkistaa/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var err error

func SetHost(path string) error {
	variables.DatabaseHost = path
	variables.Client, err = mongo.NewClient(options.Client().ApplyURI(variables.DatabaseHost))
	if err != nil {
		return err
	}

	err = connect()
	if err != nil {
		return err
	}

	err = ping()

	collections.Subject.Init()
	collections.Test.Init()
	collections.User.Init()
	collections.Question.Init()
	return err
}

func connect() error {
	return variables.Client.Connect(variables.Ctx)
}

func ping() error {
	return variables.Client.Ping(variables.Ctx, readpref.Primary())
}

func CreateSubject(subject models.Subject) error {
	return collections.Subject.Create(subject)
}

func ReadSubjects() ([]models.Subject, error) {
	return collections.Subject.Read()
}

func ReadSubjectById(id primitive.ObjectID) (models.Subject, error) {
	return collections.Subject.ReadById(id)
}

func UpdateSubject(subject models.Subject) error {
	return collections.Subject.Update(subject)
}

func DeleteSubjects() error {
	return collections.Subject.Delete()
}

func DeleteSubject(subject models.Subject) error {
	return collections.Subject.DeleteById(subject.Id)
}

func CreateUser(user models.User) error {
	return collections.User.Create(user)
}

func ReadUsers() ([]models.User, error) {
	return collections.User.Read()
}

func ReadUserById(id primitive.ObjectID) (models.User, error) {
	return collections.User.ReadById(id)
}

func UpdateUser(user models.User) error {
	return collections.User.Update(user)
}

func DeleteUsers() error {
	return collections.User.Delete()
}

func DeleteUser(user models.User) error {
	return collections.User.DeleteById(user.Id)
}

func CreateTest(test models.Test) error {
	return collections.Test.Create(test)
}

func ReadTests() ([]models.Test, error) {
	return collections.Test.Read()
}

func ReadTestById(id primitive.ObjectID) (models.Test, error) {
	return collections.Test.ReadById(id)
}

func UpdateTest(test models.Test) error {
	return collections.Test.Update(test)
}

func DeleteTests() error {
	return collections.Test.Delete()
}

func DeleteTest(test models.Test) error {
	return collections.Test.DeleteById(test.Id)
}

func CreateQuestion(question models.Question) error {
	return collections.Question.Create(question)
}

func ReadQuestions() ([]models.Question, error) {
	return collections.Question.Read()
}

func ReadQuestionById(id primitive.ObjectID) (models.Question, error) {
	return collections.Question.ReadById(id)
}

func UpdateQuestion(question models.Question) error {
	return collections.Question.Update(question)
}

func DeleteQuestions() error {
	return collections.Question.Delete()
}

func DeleteQuestion(question models.Question) error {
	return collections.Question.DeleteById(question.Id)
}
