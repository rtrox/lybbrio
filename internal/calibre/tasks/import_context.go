package tasks

import (
	"context"
	"fmt"
	"lybbrio/internal/ent/schema/ksuid"
	"strings"
)

type importOutputCtxKey string

const importOutputKey importOutputCtxKey = "importOutput"

var ignorableFilenames []string = []string{
	"cover",
	"metadata",
}

type importContext struct {
	visitedAuthors    map[int64]ksuid.ID
	visitedTags       map[int64]ksuid.ID
	visitedPublishers map[int64]ksuid.ID
	visitedLanguages  map[int64]ksuid.ID
	visitedSeries     map[int64]ksuid.ID
	failedBooks       []string
}

func newImportContext() *importContext {
	return &importContext{
		visitedAuthors:    make(map[int64]ksuid.ID),
		visitedTags:       make(map[int64]ksuid.ID),
		visitedPublishers: make(map[int64]ksuid.ID),
		visitedLanguages:  make(map[int64]ksuid.ID),
		visitedSeries:     make(map[int64]ksuid.ID),
	}
}

func (c *importContext) String() string {
	var ret strings.Builder
	if len(c.failedBooks) > 0 {
		ret.WriteString("Failed Books:\n")
		for _, book := range c.failedBooks {
			ret.WriteString(fmt.Sprintf("\t%s\n", book))
		}
	}
	return ret.String()
}

func (c *importContext) AddFailedBook(book string) {
	c.failedBooks = append(c.failedBooks, book)
}

func (c *importContext) AuthorVisited(id int64) (ksuid.ID, bool) {
	ret, ok := c.visitedAuthors[id]
	return ret, ok
}

func (c *importContext) AddAuthorVisited(id int64, ksuid ksuid.ID) {
	c.visitedAuthors[id] = ksuid
}

func (c *importContext) TagVisited(id int64) (ksuid.ID, bool) {
	ret, ok := c.visitedTags[id]
	return ret, ok
}

func (c *importContext) AddTagVisited(id int64, ksuid ksuid.ID) {
	c.visitedTags[id] = ksuid
}

func (c *importContext) PublisherVisited(id int64) (ksuid.ID, bool) {
	ret, ok := c.visitedPublishers[id]
	return ret, ok
}

func (c *importContext) AddPublisherVisited(id int64, ksuid ksuid.ID) {
	c.visitedPublishers[id] = ksuid
}

func (c *importContext) LanguageVisited(id int64) (ksuid.ID, bool) {
	ret, ok := c.visitedLanguages[id]
	return ret, ok
}

func (c *importContext) AddLanguageVisited(id int64, ksuid ksuid.ID) {
	c.visitedLanguages[id] = ksuid
}

func (c *importContext) SeriesVisited(id int64) (ksuid.ID, bool) {
	ret, ok := c.visitedSeries[id]
	return ret, ok
}

func (c *importContext) AddSeriesVisited(id int64, ksuid ksuid.ID) {
	c.visitedSeries[id] = ksuid
}

func importContextFrom(ctx context.Context) *importContext {
	output := ctx.Value(importOutputKey)
	if output == nil {
		return nil
	}
	return output.(*importContext)
}

func importContextTo(ctx context.Context, output *importContext) context.Context {
	return context.WithValue(ctx, importOutputKey, output)
}
