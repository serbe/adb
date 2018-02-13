package adb

import (
	"time"
)

// Proxy - proxy unit
type Proxy struct {
	Insert   bool          `sql:"-"                   json:"-"`
	Update   bool          `sql:"-"                   json:"-"`
	IsWork   bool          `sql:"work,notnull"        json:"-"`
	IsAnon   bool          `sql:"anon,notnull"        json:"-"`
	Checks   int           `sql:"checks,notnull"      json:"-"`
	Hostname string        `sql:"hostname,pk,notnull" json:"hostname"`
	Host     string        `sql:"host,notnull"        json:"-"`
	Port     string        `sql:"port,notnull"        json:"-"`
	CreateAt time.Time     `sql:"create_at,notnull"   json:"-"`
	UpdateAt time.Time     `sql:"update_at,notnull"   json:"-"`
	Response time.Duration `sql:"response,notnull"    json:"-"`
}

func (a *ADB) ProxyGetAll() []Proxy {
	var proxies []Proxy
	err := a.
		db.
		Model(&proxies).
		Select()
	chkErr("ProxyGetAll Select", err)
	return proxies
}

func (a *ADB) ProxyGetAllOld() []Proxy {
	var proxies []Proxy
	err := a.
		db.
		Model(&proxies).
		Where("update_at < NOW() - (INTERVAL '3 days') * checks").
		Select()
	chkErr("ProxyGetAllOld Select", err)
	return proxies
}

func (a *ADB) ProxyGetAllWorking() []Proxy {
	var proxies []Proxy
	err := a.
		db.
		Model(&proxies).
		Where("work = true").
		Select()
	chkErr("ProxyGetAllWorking Select", err)
	return proxies
}

func (a *ADB) ProxyGetAllAnonimous() []Proxy {
	var proxies []Proxy
	err := a.
		db.
		Model(&proxies).
		Where("anon = true").
		Select()
	chkErr("ProxyGetAllAnonimous Select", err)
	return proxies
}

func (a *ADB) ProxyGetUniqueHosts() []string {
	var hosts []string
	_, err := a.
		db.
		Query(&hosts, "SELECT DISTINCT host FROM proxies")
	chkErr("ProxyGetUniqueHosts Query", err)
	return hosts
}

func (a *ADB) ProxyGetFequentlyUsedPorts() []string {
	var ports []string
	_, err := a.
		db.
		Query(&ports, `
			SELECT
				port
			FROM
				proxies
			GROUP BY
				port
			ORDER BY
				count(port) DESC
			LIMIT 10
		`)
	chkErr("ProxyGetFequentlyUsedPorts Query", err)
	return ports
}

func (a *ADB) ProxyCreate(p Proxy) error {
	_, err := a.
		db.
		Model(&p).
		Insert(&p)
	chkErr("ProxyCreate Insert", err)
	return err
}

func (a *ADB) ProxyUpdate(p Proxy) error {
	_, err := a.
		db.
		Model(&p).
		Where("hostname = ?", p.Hostname).
		Update(&p)
	chkErr("ProxyUpdate Update", err)
	return err
}
