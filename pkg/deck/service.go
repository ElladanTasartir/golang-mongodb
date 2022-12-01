package deck

import "github.com/ElladanTasartir/golang-mongodb/pkg/storage"

type DeckService struct {
	deckStorage *storage.DeckStorage
}

func CreateDeckService(deckStorage *storage.DeckStorage) *DeckService {
	return &DeckService{
		deckStorage,
	}
}

func (service *DeckService) RetrieveAllDecks() (*[]storage.Deck, error) {
	return service.deckStorage.GetDecks()
}

func (service *DeckService) RetrieveADeckById(id string) (*storage.Deck, error) {
	return service.deckStorage.GetOneDeck(id)
}

func (service *DeckService) CreateNewDeck(deck *storage.Deck) (*storage.Deck, error) {
	return service.deckStorage.CreateNewDeck(deck)
}
