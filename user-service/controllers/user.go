package controllers

import (
	"context"

	"github.com/jaredwarren/rx/user-service/models"
	proto "github.com/jaredwarren/rx/user-service/proto/user"
)

// User manages User CRUD
type User struct {
	// utils *utils.Utils
}

// Create ...
func (u *User) Create(context.Context, *proto.User) (*proto.Response, error) {
	return nil, nil
}

// Get ...
func (u *User) Get(context.Context, *proto.User) (*proto.Response, error) {
	return &proto.Response{
		User: &proto.User{
			Name: "Jared Warren",
		},
	}, nil
}

// GetAll ...
func (u *User) GetAll(context.Context, *proto.Request) (*proto.Response, error) {
	return nil, nil
	// sessionCopy := databases.Database.MgDbSession.Copy()
	// defer sessionCopy.Close()

	// // Get a collection to execute the query against.
	// collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	// var users []models.User
	// err := collection.Find(bson.M{}).All(&users)
	// return users, err
}

// Auth ...
func (u *User) Auth(context.Context, *proto.User) (*proto.Token, error) {
	return nil, nil
}

// ValidateToken ...
func (u *User) ValidateToken(context.Context, *proto.Token) (*proto.Token, error) {
	return nil, nil
}

// GetByID finds a User by its id
func (u *User) GetByID(id string) (models.User, error) {
	return models.User{}, nil
	// var err error
	// err = u.utils.ValidateObjectID(id)
	// if err != nil {
	// 	return models.User{}, err
	// }

	// sessionCopy := databases.Database.MgDbSession.Copy()
	// defer sessionCopy.Close()

	// // Get a collection to execute the query against.
	// collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	// var user models.User
	// err = collection.FindId(bson.ObjectIdHex(id)).One(&user)
	// return user, err
}

// DeleteByID finds a User by its id
func (u *User) DeleteByID(id string) error {
	return nil
	// var err error
	// err = u.utils.ValidateObjectID(id)
	// if err != nil {
	// 	return err
	// }

	// sessionCopy := databases.Database.MgDbSession.Copy()
	// defer sessionCopy.Close()

	// // Get a collection to execute the query against.
	// collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	// err = collection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	// return err
}

// Login User
func (u *User) Login(name string, password string) (models.User, error) {
	// sessionCopy := databases.Database.MgDbSession.Copy()
	// defer sessionCopy.Close()

	// // Get a collection to execute the query against.
	// collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	// var user models.User
	// err := collection.Find(bson.M{"$and": []bson.M{bson.M{"name": name}, bson.M{"password": password}}}).One(&user)
	// return user, err
	return models.User{}, nil
}

// Insert adds a new User into database'
func (u *User) Insert(user models.User) error {
	return nil
	// sessionCopy := databases.Database.MgDbSession.Copy()
	// defer sessionCopy.Close()

	// // Get a collection to execute the query against.
	// collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	// err := collection.Insert(&user)
	// return err
}

// Delete remove an existing User
func (u *User) Delete(user models.User) error {
	return nil
	// sessionCopy := databases.Database.MgDbSession.Copy()
	// defer sessionCopy.Close()

	// // Get a collection to execute the query against.
	// collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	// err := collection.Remove(&user)
	// return err
}

// Update modifies an existing User
func (u *User) Update(user models.User) error {
	return nil
	// sessionCopy := databases.Database.MgDbSession.Copy()
	// defer sessionCopy.Close()

	// // Get a collection to execute the query against.
	// collection := sessionCopy.DB(databases.Database.Databasename).C(common.ColUsers)

	// err := collection.UpdateId(user.ID, &user)
	// return err
}
