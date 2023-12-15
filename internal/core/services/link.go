package services

import (
	"context"
	"fmt"

	"github.com/Furkan-Gulsen/golang-url-shortener/internal/core/domain"
	"github.com/Furkan-Gulsen/golang-url-shortener/internal/core/ports"
)

type LinkService struct {
	db    ports.LinkDB
	cache ports.Cache
}

func NewLinkDomain(d ports.LinkDB, c ports.Cache) *LinkService {
	return &LinkService{db: d, cache: c}
}

func (service *LinkService) GetAllLinksFromDB(ctx context.Context) (*[]domain.Link, error) {
	links, err := service.db.GetAllLinks(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all short URLs: %w", err)
	}
	return &links, nil
}

func (service *LinkService) GetOriginalURL(ctx context.Context, shortLinkKey string) (*string, error) {
	link, err := service.cache.Get(ctx, shortLinkKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get short URL for identifier '%s': %w", shortLinkKey, err)
	}
	return &link, nil
}

func (service *LinkService) Create(ctx context.Context, link domain.Link) error {
	if err := service.cache.Set(ctx, link.Id, link.OriginalURL); err != nil {
		return fmt.Errorf("failed to set short URL for identifier '%s': %w", link.Id, err)
	}
	if err := service.db.CreateLink(ctx, link); err != nil {
		return fmt.Errorf("failed to create short URL: %w", err)
	}
	return nil
}

func (service *LinkService) Delete(ctx context.Context, short string) error {
	if err := service.db.DeleteLink(ctx, short); err != nil {
		return fmt.Errorf("failed to delete short URL for identifier '%s': %w", short, err)
	}
	if err := service.cache.Delete(ctx, short); err != nil {
		return fmt.Errorf("failed to delete short URL for identifier '%s': %w", short, err)
	}
	return nil
}