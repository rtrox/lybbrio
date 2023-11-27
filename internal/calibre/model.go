package calibre

import "time"

type Author struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Sort  string `json:"sort"`
	Link  string `json:"link"`
	Books []Book `json:"books,omitempty" gorm:"many2many:books_authors_link;foreignKey:id;joinForeignKey:author;References:ID;JoinReferences:book"`
}

type Book struct {
	ID           int64        `json:"id"`
	Title        string       `json:"title"`
	Sort         string       `json:"sort"`
	Timestamp    *time.Time   `json:"timestamp"`
	PubDate      *time.Time   `json:"pub_date" gorm:"column:pubdate"`
	SeriesIndex  *float64     `json:"series_index"`
	AuthorSort   string       `json:"author_sort"`
	ISBN         string       `json:"isbn"`
	LCCN         string       `json:"lccn"`
	Path         string       `json:"path"`
	Flags        int64        `json:"flags"`
	UUID         string       `json:"uuid"`
	HasCover     bool         `json:"has_cover"`
	LastModified time.Time    `json:"last_modified"`
	Authors      []Author     `json:"authors" gorm:"many2many:books_authors_link;foreignKey:id;joinForeignKey:book;References:ID;JoinReferences:author"`
	Tags         []Tag        `json:"tags" gorm:"many2many:books_tags_link;foreignKey:id;joinForeignKey:book;References:ID;JoinReferences:tag"`
	Identifiers  []Identifier `json:"identifiers" gorm:"foreignKey:book"`
	Publisher    []Publisher  `json:"publisher" gorm:"many2many:books_publishers_link;foreignKey:id;joinForeignKey:book;References:ID;JoinReferences:publisher"`
	Comments     Comment      `json:"comments" gorm:"foreignKey:book"`
	Languages    []Language   `json:"languages" gorm:"many2many:books_languages_link;foreignKey:id;joinForeignKey:book;References:ID;JoinReferences:lang_code"`
}

type Identifier struct {
	ID   int64  `json:"id"`
	Book int64  `json:"book"`
	Type string `json:"type"`
	Val  string `json:"val"`
}

type Tag struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Books []Book `json:"books,omitempty" gorm:"many2many:books_tags_link;foreignKey:id;joinForeignKey:tag;References:ID;JoinReferences:book"`
}

type Publisher struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Books []Book `json:"books,omitempty" gorm:"many2many:books_publishers_link;foreignKey:id;joinForeignKey:publisher;References:ID;JoinReferences:book"`
}

type Comment struct {
	ID   int64  `json:"id"`
	Book int64  `json:"book"`
	Text string `json:"text"`
}

type Language struct {
	ID       int64  `json:"id"`
	Book     int64  `json:"book"`
	LangCode string `json:"lang"`
}

type Series struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Sort      string `json:"sort"`
	BookCount int64  `json:"book_count,omitempty"`
	Books     []Book `json:"books,omitempty" gorm:"many2many:books_series_link;foreignKey:id;joinForeignKey:series;References:ID;JoinReferences:book"`
}
