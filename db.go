package main

import "gorm.io/gorm"

type DBClient interface {
	DBMigrate() error
}

type Client struct {
	db *gorm.DB
}

func NewClient() (DBClient, error) {

}
