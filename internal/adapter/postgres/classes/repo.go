package classes

import (
	"errors"
	"school_schedule_2/internal/config"
)

var NotFoundError = errors.New("classes not found")

type Repo struct {
	cluster *config.Cluster
}

func NewRepo(cluster *config.Cluster) *Repo {
	return &Repo{
		cluster: cluster,
	}
}
