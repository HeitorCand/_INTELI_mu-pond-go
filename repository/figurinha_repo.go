package repository

import (
	"mu-pond-go/domain"

	"gorm.io/gorm"
)

type FigurinhaRepository interface {
	Create(f *domain.Figurinha) error
	FindAll(tipo *domain.FigurinhaTipo, posicao *domain.FigurinhaPosicao) ([]domain.Figurinha, error)
	FindByID(id uint) (*domain.Figurinha, error)
	Update(f *domain.Figurinha) error
	Delete(id uint) error
}

type gormFigurinhaRepository struct {
	db *gorm.DB
}

func NewFigurinhaRepository(db *gorm.DB) FigurinhaRepository {
	return &gormFigurinhaRepository{db: db}
}

func (r *gormFigurinhaRepository) Create(f *domain.Figurinha) error {
	return r.db.Create(f).Error
}

func (r *gormFigurinhaRepository) FindAll(tipo *domain.FigurinhaTipo, posicao *domain.FigurinhaPosicao) ([]domain.Figurinha, error) {
	query := r.db.Order("created_at DESC")
	if tipo != nil {
		query = query.Where("tipo = ?", *tipo)
	}
	if posicao != nil {
		query = query.Where("posicao = ?", *posicao)
	}
	var figurinhas []domain.Figurinha
	if err := query.Find(&figurinhas).Error; err != nil {
		return nil, err
	}
	return figurinhas, nil
}

func (r *gormFigurinhaRepository) FindByID(id uint) (*domain.Figurinha, error) {
	var f domain.Figurinha
	if err := r.db.First(&f, id).Error; err != nil {
		return nil, err
	}
	return &f, nil
}

func (r *gormFigurinhaRepository) Update(f *domain.Figurinha) error {
	return r.db.Save(f).Error
}

func (r *gormFigurinhaRepository) Delete(id uint) error {
	result := r.db.Delete(&domain.Figurinha{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
