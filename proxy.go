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
	Scheme   string        `sql:"scheme,notnull"      json:"-"`
	CreateAt time.Time     `sql:"create_at,notnull"   json:"-"`
	UpdateAt time.Time     `sql:"update_at,notnull"   json:"-"`
	Response time.Duration `sql:"response,notnull"    json:"-"`
}

// ProxyGetAll - get all proxies
func (a *ADB) ProxyGetAll() []Proxy {
	var proxies []Proxy
	err := a.
		db.
		Model(&proxies).
		Select()
	chkErr("ProxyGetAll Select", err)
	return proxies
}

// ProxyGetAllScheme - get all proxies by scheme
func (a *ADB) ProxyGetAllScheme(v string) []Proxy {
	var proxies []Proxy
	err := a.
		db.
		Model(&proxies).
		Where("scheme = ?", v).
		Select()
	chkErr("ProxyGetAll Select", err)
	return proxies
}

// ProxyGetAllOld - get all old proxies
func (a *ADB) ProxyGetAllOld() []Proxy {
	var proxies []Proxy
	err := a.
		db.
		Model(&proxies).
		Where("work = true OR update_at < NOW() - (INTERVAL '3 days') * checks").
		Select()
	chkErr("ProxyGetAllOld Select", err)
	return proxies
}

// ProxyGetAllWorking - get all working proxies
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

// ProxyGetAllWorkingScheme - get all working proxies by scheme
func (a *ADB) ProxyGetAllWorkingScheme(v string) []Proxy {
	var proxies []Proxy
	err := a.
		db.
		Model(&proxies).
		Where("work = true ANS scheme = ?", v).
		Select()
	chkErr("ProxyGetAllWorking Select", err)
	return proxies
}

// ProxyGetAllAnonymous - get all anonimous proxies
func (a *ADB) ProxyGetAllAnonymous() []Proxy {
	var proxies []Proxy
	err := a.
		db.
		Model(&proxies).
		Where("anon = true").
		Select()
	chkErr("ProxyGetAllAnonymous Select", err)
	return proxies
}

// ProxyGetAllAnonymousScheme - get all anonimous proxies by scheme
func (a *ADB) ProxyGetAllAnonymousScheme(v string) []Proxy {
	var proxies []Proxy
	err := a.
		db.
		Model(&proxies).
		Where("anon = true AND scheme = ?", v).
		Select()
	chkErr("ProxyGetAllAnonymous Select", err)
	return proxies
}

// ProxyGetUniqueHosts - gel all unique hosts
func (a *ADB) ProxyGetUniqueHosts() []string {
	var hosts []string
	_, err := a.
		db.
		Query(&hosts, "SELECT DISTINCT host FROM proxies")
	chkErr("ProxyGetUniqueHosts Query", err)
	return hosts
}

// ProxyGetFrequentlyUsedPorts - get 10 frequently used ports
func (a *ADB) ProxyGetFrequentlyUsedPorts() []string {
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
	chkErr("ProxyGetFrequentlyUsedPorts Query", err)
	return ports
}

// ProxyInsert - insert new proxy
func (a *ADB) ProxyInsert(p Proxy) {
	_, err := a.
		db.
		Model(&p).
		Insert(&p)
	chkErr("ProxyInsert", err)
}

// ProxyUpdate - update existing proxy
func (a *ADB) ProxyUpdate(p Proxy) {
	_, err := a.
		db.
		Model(&p).
		Where("hostname = ?", p.Hostname).
		Update(&p)
	chkErr("ProxyUpdate", err)
}
