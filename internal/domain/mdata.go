package domain

import (
	"context"
	"time"
)

type MangaDataRepo interface {
	Store(ctx context.Context, mdata *MangaData) (*MangaData, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, mdata *MangaData) (*MangaData, error)
	ListMangaData(ctx context.Context, apiKey string) ([]MangaData, error)
	GetMangaDataByApiKey(ctx context.Context, apiKey string) (*MangaData, error)
}

type MangaData struct {
	ID                      int           `json:"id"`
	BackupManga             []Manga       `json:"backupManga"`
	BackupCategories        []Category    `json:"backupCategories"`
	BackupBrokenSources     []interface{} `json:"backupBrokenSources"`
	BackupSources           []Sources     `json:"backupSources"`
	BackupPreferences       []interface{} `json:"backupPreferences"`
	BackupSourcePreferences []interface{} `json:"backupSourcePreferences"`
	UserApiKey              *APIKey       `json:"user_api_key,omitempty"`
	CreatedAt               time.Time     `json:"created_at"`
	UpdatedAt               time.Time     `json:"updated_at"`
}

type Manga struct {
	Source         int64      `json:"source"`
	URL            string     `json:"url"`
	Title          string     `json:"title"`
	Artist         string     `json:"artist"`
	Author         string     `json:"author"`
	Description    string     `json:"description"`
	Genre          []string   `json:"genre"`
	Status         int        `json:"status"`
	ThumbnailURL   string     `json:"thumbnailUrl"`
	DateAdded      int64      `json:"dateAdded"`
	Viewer         int        `json:"viewer"`
	Chapters       []Chapter  `json:"chapters"`
	Tracking       []Tracking `json:"tracking"`
	Favorite       bool       `json:"favorite"`
	Categories     []int      `json:"categories"`
	ViewerFlags    int        `json:"viewer_flags"`
	History        []History  `json:"history"`
	LastModifiedAt int64      `json:"lastModifiedAt"`
}

type Chapter struct {
	URL            string  `json:"url"`
	Name           string  `json:"name"`
	Scanlator      string  `json:"scanlator"`
	Read           bool    `json:"read"`
	Bookmark       bool    `json:"bookmark"`
	LastPageRead   int64   `json:"lastPageRead"`
	DateFetch      int64   `json:"dateFetch"`
	DateUpload     int64   `json:"dateUpload"`
	ChapterNumber  float64 `json:"chapterNumber"`
	SourceOrder    int     `json:"sourceOrder"`
	LastModifiedAt int64   `json:"lastModifiedAt"`
}

type Tracking struct {
	SyncId              *int     `json:"syncId,omitempty"`
	LibraryId           *int64   `json:"libraryId,omitempty"`
	MediaIdInt          *int     `json:"mediaIdInt,omitempty"`
	TrackingUrl         *string  `json:"trackingUrl,omitempty"`
	Title               *string  `json:"title,omitempty"`
	LastChapterRead     *float64 `json:"lastChapterRead,omitempty"`
	TotalChapters       *int     `json:"totalChapters,omitempty"`
	Score               *float64 `json:"score,omitempty"`
	Status              *int     `json:"status,omitempty"`
	StartedReadingDate  *int64   `json:"startedReadingDate,omitempty"`
	FinishedReadingDate *int64   `json:"finishedReadingDate,omitempty"`
	MediaId             *int64   `json:"mediaId,omitempty"`
}

type History struct {
	URL          string `json:"url"`
	LastRead     int64  `json:"lastRead"`
	ReadDuration int64  `json:"readDuration"`
}

type Sources struct {
	Name     string `json:"name"`
	SourceID int64  `json:"sourceId"`
}

type Category struct {
	Name  string `json:"name"`
	Flags int    `json:"flags"`
	Order int    `json:"order"`
}
