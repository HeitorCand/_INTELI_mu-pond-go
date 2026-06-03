package domain

import "time"

type FigurinhaTipo string

const (
	TipoComum         FigurinhaTipo = "comum"
	TipoBrilhante     FigurinhaTipo = "brilhante"
	TipoLegendsOuro   FigurinhaTipo = "legends_ouro"
	TipoLegendsBronze FigurinhaTipo = "legends_bronze"
)

func (t FigurinhaTipo) IsValid() bool {
	switch t {
	case TipoComum, TipoBrilhante, TipoLegendsOuro, TipoLegendsBronze:
		return true
	}
	return false
}

type FigurinhaPosicao string

const (
	PosicaoGoleiro      FigurinhaPosicao = "Goleiro"
	PosicaoZagueiro     FigurinhaPosicao = "Zagueiro"
	PosicaoMeioCampista FigurinhaPosicao = "Meio-campista"
	PosicaoAtacante     FigurinhaPosicao = "Atacante"
)

func (p FigurinhaPosicao) IsValid() bool {
	switch p {
	case PosicaoGoleiro, PosicaoZagueiro, PosicaoMeioCampista, PosicaoAtacante:
		return true
	}
	return false
}

type Figurinha struct {
	ID        uint             `json:"id" gorm:"primaryKey"`
	Numero    string           `json:"numero"`
	Tipo      FigurinhaTipo    `json:"tipo"`
	Posicao   FigurinhaPosicao `json:"posicao"`
	UpdatedAt time.Time        `json:"updated_at"`
	CreatedAt time.Time        `json:"created_at"`
}

type CreateFigurinhaRequest struct {
	Numero  string           `json:"numero" binding:"required"`
	Tipo    FigurinhaTipo    `json:"tipo" binding:"required"`
	Posicao FigurinhaPosicao `json:"posicao" binding:"required"`
}

type UpdateFigurinhaRequest struct {
	Numero  *string           `json:"numero"`
	Tipo    *FigurinhaTipo    `json:"tipo"`
	Posicao *FigurinhaPosicao `json:"posicao"`
}
