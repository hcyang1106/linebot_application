package repository

import (
	"github.com/hcyang1106/awesomeProject/config"
	"github.com/hcyang1106/awesomeProject/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Repository struct {
	DB     *mgo.Collection
	Config *config.Config
}

func NewRepository(config *config.Config) *Repository {
	session, _ := mgo.Dial(config.MongoDBAddress)
	db := session.DB(config.DbName)
	c := db.C(config.CollectionName)
	repo := &Repository{
		DB: c,
		Config: config,
	}
	return repo
}

func (r *Repository) CreateHistory(history *model.History) error {
	err := r.DB.Insert(history)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetHistoriesByName(name string) ([]*model.History, error) {
	histories := make([]*model.History, 0)
	err := r.DB.Find(bson.M{"name": name}).All(&histories)
	if err != nil {
		return nil, err
	}
	return histories, nil
}

func (r *Repository) FindOneHistoryByName(name string) (*model.History, error) {
	var history model.History
	err := r.DB.Find(bson.M{"name": name}).One(&history)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &history, nil
}
