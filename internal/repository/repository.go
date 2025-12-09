package repository

import (
	"context"

	"github.com/s-dartau/notes/internal/model"
)

type NoteRepository interface {
	Create(ctx context.Context, n model.Note) error
	GetAll(ctx context.Context)
	GetById(ctx context.Context)
	Update(ctx context.Context)
	Delete(ctx context.Context)
}
