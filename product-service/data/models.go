package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func New(mongoClient *mongo.Client) Models {
	client = mongoClient
	return Models{
		Product: Product{},
	}
}

type Models struct {
	Product Product
}

type Product struct {
	ID          string    `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string    `bson:"name" json:"name"`
	Description string    `bson:"description" json:"description"`
	Price       float64   `bson:"price" json:"price"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}

func (p *Product) Insert(product Product) error {
	collection := client.Database("review-rating").Collection("products")

	_, err := collection.InsertOne(context.TODO(), Product{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		log.Println("Error inserting into products:", err)
		return err
	}

	return nil
}

func (p *Product) All() ([]*Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("review-rating").Collection("products")

	opts := options.Find()
	opts.SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		log.Println("Finding all products error:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []*Product

	for cursor.Next(ctx) {
		var item Product
		err := cursor.Decode(&item)
		if err != nil {
			log.Println("Error decoding product into slice:", err)
			return nil, err
		}
		products = append(products, &item)
	}

	return products, nil
}

func (p *Product) GetOne(id string) (*Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("review-rating").Collection("products")

	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var product Product
	err = collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *Product) DropCollection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("review-rating").Collection("products")

	if err := collection.Drop(ctx); err != nil {
		return err
	}
	return nil
}

func (p *Product) Update() (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("review-rating").Collection("products")

	docID, err := primitive.ObjectIDFromHex(p.ID)
	if err != nil {
		return nil, err
	}

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": docID},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "name", Value: p.Name},
				{Key: "description", Value: p.Description},
				{Key: "price", Value: p.Price},
				{Key: "updated_at", Value: time.Now()},
			}},
		},
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}
