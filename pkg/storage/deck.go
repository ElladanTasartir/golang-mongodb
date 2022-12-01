package storage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	apperrors "github.com/ElladanTasartir/golang-mongodb/pkg/errors"
)

type Card struct {
	Name   string `json:"name" binding:"required"`
	Stars  int    `json:"stars" binding:"required"`
	Effect string `json:"effect" binding:"required"`
}

type Deck struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name        string             `json:"name"`
	Theme       string             `json:"theme"`
	Cards       []Card             `json:"cards"`
	FusionCards []Card             `json:"fusion_cards" bson:"fusion_cards,omitempty"`
}

func (d *Deck) MarshalJSON() ([]byte, error) {
	if d.Cards == nil {
		d.Cards = make([]Card, 0)
	}

	if d.FusionCards == nil {
		d.FusionCards = make([]Card, 0)
	}

	type newDeck Deck

	return json.Marshal((*newDeck)(d))
}

const deckCollection = "decks"

type DeckStorage struct {
	collection *mongo.Collection
}

func CreateDeckStorage(client *DbClient) *DeckStorage {
	return &DeckStorage{
		collection: client.database.Collection(deckCollection),
	}
}

func (deckStorage *DeckStorage) CreateNewDeck(deck *Deck) (*Deck, error) {
	deck.Id = primitive.NewObjectID()

	_, err := deckStorage.collection.InsertOne(context.TODO(), deck)
	if err != nil {
		fmt.Println(err.Error())
		return deck, err
	}

	return deck, nil
}

func (deckStorage *DeckStorage) GetDecks() (*[]Deck, error) {
	var decks []Deck = []Deck{}
	cursor, queryError := deckStorage.collection.Find(context.TODO(), bson.D{})
	if queryError != nil {
		fmt.Println(queryError.Error())
		return &decks, errors.New("There was an error with getting decks data")
	}

	if err := cursor.All(context.TODO(), &decks); err != nil {
		fmt.Println(err.Error())
		return &decks, errors.New("There was an error converting decks data")
	}

	return &decks, nil
}

func (deckStorage *DeckStorage) GetOneDeck(id string) (*Deck, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.D{primitive.E{
		Key:   "_id",
		Value: objectId,
	}}

	var deck Deck
	err = deckStorage.collection.FindOne(context.TODO(), filter).Decode(&deck)
	if err != nil {
		fmt.Print("entrei no erro")
		if err == mongo.ErrNoDocuments {
			return nil, &apperrors.NotFound{
				Code: 404, Entity: "Deck",
			}
		}

		return nil, err
	}

	return &deck, nil
}
