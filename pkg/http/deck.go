package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ElladanTasartir/golang-mongodb/pkg/deck"
	apperrors "github.com/ElladanTasartir/golang-mongodb/pkg/errors"
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
	s.httpServer.GET("/decks/:id", router.getDeckById)
	s.httpServer.POST("/decks", router.createDecks)
}

func (r *DeckRouter) createDecks(ctx *gin.Context) {
	var body CreateDeckBody
	if err := ctx.ShouldBind(&body); err != nil {
		fmt.Println(body)
		fmt.Println(err)
		SendResponse(ctx, &Response{
			Code: http.StatusBadRequest,
			Body: gin.H{
				"message": "Error validating body of request",
			},
		})
		return
	}

	newDeck := storage.Deck{
		Name:        body.Name,
		Theme:       body.Theme,
		Cards:       body.Cards,
		FusionCards: body.FusionCards,
	}

	deck, err := r.service.CreateNewDeck(&newDeck)
	if err != nil {
		SendResponse(ctx, &Response{
			Code: http.StatusUnprocessableEntity,
			Body: gin.H{
				"message": err.Error(),
			},
		})
		return
	}

	SendResponse(ctx, &Response{
		Code: http.StatusCreated,
		Body: deck,
	})
	return
}

func (r *DeckRouter) getDecks(ctx *gin.Context) {
	decks, err := r.service.RetrieveAllDecks()
	if err != nil {
		SendResponse(ctx, &Response{
			Code: http.StatusBadRequest,
			Body: gin.H{
				"message": err.Error(),
			},
		})
		return
	}

	SendResponse(ctx, &Response{
		Code: http.StatusOK,
		Body: decks,
	})
	return
}

func (r *DeckRouter) getDeckById(ctx *gin.Context) {
	id, sent := ctx.Params.Get("id")
	if !sent {
		SendResponse(ctx, &Response{
			Code: http.StatusBadRequest,
			Body: gin.H{
				"message": "Deck ID must be sent",
			},
		})
		return
	}

	deck, err := r.service.RetrieveADeckById(id)
	if err != nil {
		var appError *apperrors.NotFound
		if errors.As(err, &appError) {
			SendResponse(ctx, &Response{
				Code: http.StatusNotFound,
				Body: gin.H{
					"message": err.Error(),
				},
			})
			return
		}

		SendResponse(ctx, &Response{
			Code: http.StatusBadRequest,
			Body: gin.H{
				"message": err.Error(),
			},
		})
		return
	}

	SendResponse(ctx, &Response{
		Code: http.StatusOK,
		Body: deck,
	})
	return
}
