package db

import (
	"database/sql"
	"errors"

	"go.uber.org/zap"

	_ "github.com/lib/pq"
)
var ErrNotFound = errors.New("not found")
// object that works w data base
type DbClient struct {
	db     *sql.DB 
	logger *zap.Logger
}

type Song struct {
	Title string `json:"title"`
	Author string `json:"author"`
	SpotifyLink string `json:"spotify_link"`
}

func NewDbClient(psqlInfo string, logger *zap.Logger) *DbClient {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.Error("failed to init connection", zap.Error(err))
		return nil
	}
	err = db.Ping()
	if err != nil {
		logger.Error("failed to ping database", zap.Error(err))
		return nil
	}

	client := &DbClient{
		db:     db,
		logger: logger,
	}
	return client
}

func (c *DbClient) GetSongs() ([]Song, error) {
	sqlStatement := "SELECT * FROM public.song"
	
	rows, err := c.db.Query(sqlStatement)
	if err != nil {
		return nil,err
	}
	defer rows.Close()
	res := make([]Song, 0)
	for rows.Next() {
		var s Song
		
		if err := rows.Scan(&s.Title, &s.Author, &s.SpotifyLink); err != nil {
            return nil, err
        } 
		res = append(res, s)
	}
	return res, nil
}

func (c *DbClient) Shutdown() {
	c.db.Close()
}
