package models

import (
	"errors"
	"strings"
	"time"
)

type Publication struct {
	Id             uint64    `json:"id,omitempty"`
	Title          string    `json:"title,omitempty"`
	Content        string    `json:"content,omitempty"`
	AuthorId       uint64    `json:"authorId,omitempty"`
	AuthorNickName string    `json:"authorNickName,omitempty"`
	Likes          uint64    `json:"likes"`
	CreatedAt      time.Time `json:"createdAt,omitempty"`
}

func (p *Publication) validate() error {
	if p.Title == "" {
		return errors.New("title is required")
	}

	if p.Content == "" {
		return errors.New("content is required")
	}

	return nil
}

func (p *Publication) format() {
	p.Title = strings.TrimSpace(p.Title)
	p.Content = strings.TrimSpace(p.Content)
}

func (p *Publication) Prepare() error {
	if err := p.validate(); err != nil {
		return err
	}

	p.format()
	return nil
}
