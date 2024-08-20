package storage

import "path"

type Config struct {
	FileName string
	FilePath string
}

func (c *Config) fullPath() string {
	return path.Join(c.FilePath, c.FileName)
}
