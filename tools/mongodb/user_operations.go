package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"user-service/internal/user"
)

func (db *MongoDB) Upsert(user *user.User) error {
	opts := options.Update().SetUpsert(true)

	if _, updateErr := db.GetUserCollection().UpdateByID(CTX, user.ID, bson.M{"$set": user}, opts); updateErr != nil {
		return updateErr
	}

	return nil
}

func (db *MongoDB) IsEmailExists(email string) (bool, error) {
	count, err := db.GetUserCollection().CountDocuments(CTX, bson.M{"email": email})
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (db *MongoDB) IsNicknameExists(nickname string) (bool, error) {
	count, err := db.GetUserCollection().CountDocuments(CTX, bson.M{"nickname": nickname})
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (db *MongoDB) DeleteUserByID(id primitive.ObjectID) error {
	if _, err := db.GetUserCollection().DeleteOne(CTX, bson.M{"_id": id}); err != nil {
		return err
	}

	return nil
}

func (db *MongoDB) GetUserByID(id primitive.ObjectID) (*user.User, error) {
	var user user.User

	if err := db.GetUserCollection().FindOne(CTX, bson.M{"_id": id}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *MongoDB) GetUsers() ([]*user.User, error) {
	var users []*user.User

	cursor, err := db.GetUserCollection().Find(CTX, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(CTX, &users); err != nil {
		return nil, err
	}

	return users, nil
}
