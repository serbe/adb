package adb

import (
	"time"
)

// Proxy - proxy unit
type Proxy struct {
	Insert   bool          `sql:"-"           json:"-"`
	Update   bool          `sql:"-"           json:"-"`
	IsWork   bool          `sql:"work"        json:"-"`
	IsAnon   bool          `sql:"anon"        json:"-"`
	Checks   int           `sql:"checks"      json:"-"`
	Hostname string        `sql:"hostname,pk" json:"hostname"`
	Host     string        `sql:"host"        json:"-"`
	Port     string        `sql:"port"        json:"-"`
	CreateAt time.Time     `sql:"create_at"   json:"-"`
	UpdateAt time.Time     `sql:"update_at"   json:"-"`
	Response time.Duration `sql:"response"    json:"-"`
}

func (a *ADB) ProxyGetAll() []Proxy {
	var proxies []Proxy
	err := a.
		db.
		Model(&proxies).
		Select()
	chkErr("ProxyGetAll select", err)
	return proxies
}

func (a *ADB) ProxyGetAllOld() []Proxy {
	var proxies []Proxy
	err := a.
		db.
		Model(&proxies).
		Where("update_at < NOW() - (INTERVAL '3 days') * checks").
		Select()
	chkErr("ProxyGetAllOld select", err)
	return proxies
}

func (a *ADB) ProxyGetAllWorking() []Proxy {
	var proxies []Proxy
	err := a.
		db.
		Model(&proxies).
		Where("work = true").
		Select()
	chkErr("ProxyGetAllWorking select", err)
	return proxies
}

func (a *ADB) ProxyGetAllAnonimous() []Proxy {
	var proxies []Proxy
	err := a.
		db.
		Model(&proxies).
		Where("anon = true").
		Select()
	chkErr("ProxyGetAllAnonimous select", err)
	return proxies
}
