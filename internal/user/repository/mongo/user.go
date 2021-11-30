package mongo

import (
	"context"
	"errors"
	"fmt"

	models "github.com/Alexander272/games-library/internal/user/models"
	"github.com/Alexander272/games-library/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	db *mongo.Collection
}

func NewUserRepo(db *mongo.Database, collection string) *UserRepo {
	return &UserRepo{
		db: db.Collection(collection),
	}
}

func (r *UserRepo) Create(ctx context.Context, user models.User) (id string, err error) {
	res, err := r.db.InsertOne(ctx, user)
	if err != nil {
		return id, fmt.Errorf("failed to execute query. error: %w", err)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return id, fmt.Errorf("failed to convert objectid")
	}
	logger.Tracef("Created document with oid %s.\n", oid)
	return oid.Hex(), nil
}

func (r *UserRepo) GetAll(ctx context.Context) (users []models.User, err error) {
	filter := bson.M{}
	cur, err := r.db.Find(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return users, models.ErrUserNotFound
		}
		return users, fmt.Errorf("failed to execute query. error: %w", err)
	}
	if err := cur.All(ctx, &users); err != nil {
		return users, fmt.Errorf("failed to decode document. error: %w", err)
	}

	return users, nil
}

func (r *UserRepo) GetById(ctx context.Context, userId string) (user models.User, err error) {
	oid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return user, fmt.Errorf("failed to convert hex to objectid. error: %w", err)
	}

	filter := bson.M{"_id": oid}
	res := r.db.FindOne(ctx, filter)
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return user, models.ErrUserNotFound
		}
		return user, fmt.Errorf("failed to execute query. error: %w", res.Err())
	}
	if err := res.Decode(&user); err != nil {
		return user, fmt.Errorf("failed to decode document. error: %w", err)
	}

	return user, nil
}

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (user models.User, err error) {
	filter := bson.M{"email": email}
	res := r.db.FindOne(ctx, filter)
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return user, models.ErrUserNotFound
		}
		return user, fmt.Errorf("failed to execute query. error: %w", res.Err())
	}
	if err := res.Decode(&user); err != nil {
		return user, fmt.Errorf("failed to decode document. error: %w", err)
	}

	return user, nil
}

func (r *UserRepo) Update(ctx context.Context, user models.User) error {
	oid, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return fmt.Errorf("failed to convert hex to objectid. error: %w", err)
	}

	filter := bson.M{"_id": oid}
	userByte, err := bson.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal document. error: %w", err)
	}

	var updateObj bson.M
	if err := bson.Unmarshal(userByte, &updateObj); err != nil {
		return fmt.Errorf("failed to unmarshal document. error: %w", err)
	}

	delete(updateObj, "_id")
	update := bson.M{"$set": updateObj}

	res, err := r.db.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	if res.MatchedCount == 0 {
		return models.ErrUserNotFound
	}

	logger.Tracef("Matched %v documents and updated %v documents.\n", res.MatchedCount, res.ModifiedCount)
	return nil
}

func (r *UserRepo) Remove(ctx context.Context, userId string) error {
	oid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return fmt.Errorf("failed to convert hex to objectid. error: %w", err)
	}

	filter := bson.M{"_id": oid}
	res, err := r.db.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %w", err)
	}
	if res.DeletedCount == 0 {
		return models.ErrUserNotFound
	}

	logger.Tracef("Delete %v documents.\n", res.DeletedCount)
	return nil
}
