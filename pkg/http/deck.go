package http

import (
	"fmt"
	"net/http"

	"github.com/ElladanTasartir/golang-mongodb/pkg/deck"
	"github.com/ElladanTasartir/golang-mongodb/pkg/http/helpers"
	"github.com/ElladanTasartir/golang-mongodb/pkg/storage"
	"github.com/gin-gonic/gin"
)

type DeckRouter struct {
	service *deck.DeckService
}

type CreateDeckBody struct {
	Name        string         `json:"name" binding:"required"`
	Theme       string         `json:"theme" binding:"required"`
	Cards       []storage.Card `json:"cards" binding:"required"`
	FusionCards []storage.Card `json:"fusion_cards"`
}

func (s *Server) AddDecksEndpoints() {
	storage := storage.CreateDeckStorage(s.client)
	service := deck.CreateDeckService(storage)

	router := DeckRouter{
		service,
	}

	s.httpServer.GET("/decks", router.getDecks)
	s.httpServer.POST("/decks", router.createDecks)
}

func (r *DeckRouter) createDecks(ctx *gin.Context) {
	var body CreateDeckBody
	if err := ctx.ShouldBind(&body); err != nil {
		fmt.Println(body)
		fmt.Println(err)
		helpers.SendResponse(ctx, &helpers.Response{
			Code: http.StatusBadRequest,
			Body: gin.H{
				"message": "Error validating body of request",
			},
		})
	}

	newDeck := storage.Deck{
		Name:        body.Name,
		Theme:       body.Theme,
		Cards:       body.Cards,
		FusionCards: body.FusionCards,
	}

	deck, err := r.service.CreateNewDeck(&newDeck)
	if err != nil {
		helpers.SendResponse(ctx, &helpers.Response{
			Code: http.StatusUnprocessableEntity,
			Body: gin.H{
				"message": err.Error(),
			},
		})
	}

	helpers.SendResponse(ctx, &helpers.Response{
		Code: http.StatusCreated,
		Body: deck,
	})
}

func (r *DeckRouter) getDecks(ctx *gin.Context) {
	decks, err := r.service.RetrieveAllDecks()
	if err != nil {
		helpers.SendResponse(ctx, &helpers.Response{
			Code: http.StatusBadRequest,
			Body: gin.H{
				"message": err.Error(),
			},
		})
	}

	helpers.SendResponse(ctx, &helpers.Response{
		Code: http.StatusOK,
		Body: decks,
	})
}