package service

import (
	"errors"
	"mu-pond-go/domain"
	"mu-pond-go/repository"

	"gorm.io/gorm"
)

var (
	ErrFigurinhaNotFound = errors.New("figurinha não encontrado")
	ErrInvalidTipo       = errors.New("tipo inválido")
	ErrInvalidPosicao    = errors.New("posicao inválida")
	ErrCampoObrigatorio  = errors.New("todos os campos são obrigatórios")
)

type FigurinhaService interface {
	Create(req domain.CreateFigurinhaRequest) (*domain.Figurinha, error)
	List(tipo, posicao string) ([]domain.Figurinha, error)
	GetByID(id uint) (*domain.Figurinha, error)
	Update(id uint, req domain.UpdateFigurinhaRequest) (*domain.Figurinha, error)
	Delete(id uint) error
}

type figurinhaService struct {
	repo repository.FigurinhaRepository
}

func NewFigurinhaService(repo repository.FigurinhaRepository) FigurinhaService {
	return &figurinhaService{repo: repo}
}

func (s *figurinhaService) Create(req domain.CreateFigurinhaRequest) (*domain.Figurinha, error) {
	if req.Numero == "" || string(req.Tipo) == "" || string(req.Posicao) == "" {
		return nil, ErrCampoObrigatorio
	}
	if !req.Tipo.IsValid() {
		return nil, ErrInvalidTipo
	}
	if !req.Posicao.IsValid() {
		return nil, ErrInvalidPosicao
	}

	f := &domain.Figurinha{
		Numero:  req.Numero,
		Tipo:    req.Tipo,
		Posicao: req.Posicao,
	}

	if err := s.repo.Create(f); err != nil {
		return nil, err
	}
	return f, nil
}

func (s *figurinhaService) List(tipo, posicao string) ([]domain.Figurinha, error) {
	var tipoPtr *domain.FigurinhaTipo
	var posicaoPtr *domain.FigurinhaPosicao

	if tipo != "" {
		t := domain.FigurinhaTipo(tipo)
		if !t.IsValid() {
			return nil, ErrInvalidTipo
		}
		tipoPtr = &t
	}

	if posicao != "" {
		p := domain.FigurinhaPosicao(posicao)
		if !p.IsValid() {
			return nil, ErrInvalidPosicao
		}
		posicaoPtr = &p
	}

	return s.repo.FindAll(tipoPtr, posicaoPtr)
}

func (s *figurinhaService) GetByID(id uint) (*domain.Figurinha, error) {
	f, err := s.repo.FindByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrFigurinhaNotFound
	}
	return f, err
}

func (s *figurinhaService) Update(id uint, req domain.UpdateFigurinhaRequest) (*domain.Figurinha, error) {
	f, err := s.repo.FindByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrFigurinhaNotFound
	}
	if err != nil {
		return nil, err
	}

	if req.Tipo != nil {
		if !req.Tipo.IsValid() {
			return nil, ErrInvalidTipo
		}
		f.Tipo = *req.Tipo
	}
	if req.Posicao != nil {
		if !req.Posicao.IsValid() {
			return nil, ErrInvalidPosicao
		}
		f.Posicao = *req.Posicao
	}
	if req.Numero != nil {
		f.Numero = *req.Numero
	}

	if err := s.repo.Update(f); err != nil {
		return nil, err
	}
	return f, nil
}

func (s *figurinhaService) Delete(id uint) error {
	err := s.repo.Delete(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrFigurinhaNotFound
	}
	return err
}
