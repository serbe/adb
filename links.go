package adb

import (
	"time"
)

// Link - link unit
type Link struct {
	Insert   bool      `sql:"-"                   json:"-"`
	Update   bool      `sql:"-"                   json:"-"`
	Iterate  bool      `sql:"iterate,notnull"     json:"-"`
	Num      int64     `sql:"num,notnull"         json:"-"`
	Hostname string    `sql:"hostname,pk,notnull" json:"hostname"`
	UpdateAt time.Time `sql:"update_at,notnull"   json:"-"`
}

// LinksGetAll - get all links
func (a *ADB) LinksGetAll() ([]Link, error) {
	var links []Link
	err := a.
		db.
		Model(&links).
		Select()
	return links, err
}

// LinksGetAllIterate - get all iterate links
func (a *ADB) LinksGetAllIterate() ([]Link, error) {
	var links []Link
	err := a.
		db.
		Model(&links).
		Where("iterate = true").
		Select()
	return links, err
}

// LinksGetAllOld - get all old links
func (a *ADB) LinksGetAllOld() ([]Link, error) {
	var links []Link
	err := a.
		db.
		Model(&links).
		Where("iterate = true AND update_at < NOW() - (INTERVAL '1 hours')").
		Select()
	return links, err
}

// LinkInsert - insert new link
func (a *ADB) LinkInsert(l Link) error {
	_, err := a.db.Model(&l).Insert(&l)
	return err
}

// LinkUpdate - update existing link
func (a *ADB) LinkUpdate(l Link) error {
	_, err := a.db.Model(&l).Where("hostname = ?", l.Hostname).Update(&l)
	return err
}
