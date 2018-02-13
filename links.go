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

func (a *ADB) LinksGetAll() []Link {
	var links []Link
	err := a.
		db.
		Model(&links).
		Select()
	chkErr("LinksGetAll select", err)
	return links
}

func (a *ADB) LinksGetAllIterate() []Link {
	var links []Link
	err := a.
		db.
		Model(&links).
		Where("iterate = true").
		Select()
	chkErr("LinksGetAllIterate select", err)
	return links
}

func (a *ADB) LinksGetAllOld() []Link {
	var links []Link
	err := a.
		db.
		Model(&links).
		Where("iterate = true AND update_at < NOW() - (INTERVAL '1 hours')").
		Select()
	chkErr("LinksGetAllOld select", err)
	return links
}

func (a *ADB) LinkCreate(l Link) error {
	_, err := a.db.Model(&l).Insert(&l)
	chkErr("LinkCreate Insert", err)
	return err
}

func (a *ADB) LinkUpdate(l Link) error {
	_, err := a.db.Model(&l).Where("hostname = ?", l.Hostname).Update(&l)
	chkErr("LinkUpdate Update", err)
	return err
}
