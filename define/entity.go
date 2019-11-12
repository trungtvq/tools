package define

import (
	"time"
)

type Video struct {
	ID        int       `json:"id" gorm:"primary_key,auto_increment,not null"`
	YoutubeID string    `json:"youtube_id" gorm:"not null"`
	Title     string    `json:"title"`
	Channel   string    `json:"channel"`
	Thumbnail string    `json:"thumbnail"`
	Duration  string    `json:"duration"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type User struct {
	ID        string    `json:"id" gorm:"primary_key,auto_increment,not null"`
	FirstName string    `json:"family_name"`
	LastName  string    `json:"given_name"`
	ImageURL  string    `json:"picture"`
	Status    int       `json:"status" gorm:"default:1"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Playlist struct {
	ID        uint32    `json:"id" gorm:"primary_key,auto_increment,not null"`
	UserID    string    `json:"-" gorm:"not null;index:user_id_idx"`
	Name      string    `json:"name" gorm:"not null"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type PlaylistVideo struct {
	ID         uint32    `json:"id" gorm:"primary_key,auto_increment,not null"`
	PlaylistID uint32    `json:"playlist_id" gorm:"not null"`
	VideoID    string    `json:"video_id" gorm:"not null"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}

type Favorited struct {
	ID        uint32    `json:"id" gorm:"primary_key,auto_increment,not null"`
	UserID    string    `json:"-" gorm:"not null"`
	VideoID   string    `json:"video_id" gorm:"not null"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Recently struct {
	ID        uint32    `json:"id" gorm:"primary_key,auto_increment,not null"`
	UserID    string    `json:"-" gorm:"not null"`
	VideoID   string    `json:"video_id" gorm:"not null"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Hashtag struct {
	ID        uint32    `json:"id" gorm:"primary_key,auto_increment,not null"`
	Name      string    `json:"name" gorm:"not null"`
	Value     string    `json:"value" gorm:"not null"`
	Type      string    `json:"type" gorm:"not null"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type InvalidSessionKey struct {
	ID         uint32 `json:"id" gorm:"primary_key,auto_increment,not null"`
	SessionKey string `json:"session_key" gorm:"not null"`
}
