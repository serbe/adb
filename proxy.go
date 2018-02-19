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
func (a *ADB) ProxyGetAll() ([]Proxy, error) {
	var proxies []Proxy
	err := a.
		db.
		Model(&proxies).
		Select()
	return proxies, err
}

// ProxyGetAllScheme - get all proxies by scheme
func (a *ADB) ProxyGetAllScheme(v string) ([]Proxy, error) {
	var proxies []Proxy
	err := a.
		db.
		Model(&proxies).
		Where("scheme = ?", v).
		Select()
	return proxies, err
}

// ProxyGetAllOld - get all old proxies
func (a *ADB) ProxyGetAllOld() ([]Proxy, error) {
	var proxies []Proxy
	err := a.
		db.
		Model(&proxies).
		Where("work = true OR update_at < NOW() - (INTERVAL '3 days') * checks").
		Select()
	return proxies, err
}

// ProxyGetAllWorking - get all working proxies
func (a *ADB) ProxyGetAllWorking() ([]Proxy, error) {
	var proxies []Proxy
	err := a.
		db.
		Model(&proxies).
		Where("work = true").
		Select()
	return proxies, err
}

// ProxyGetAllWorkingScheme - get all working proxies by scheme
func (a *ADB) ProxyGetAllWorkingScheme(v string) ([]Proxy, error) {
	var proxies []Proxy
	err := a.
		db.
		Model(&proxies).
		Where("work = true ANS scheme = ?", v).
		Select()
	return proxies, err
}

// ProxyGetAllAnonymous - get all anonymous proxies
func (a *ADB) ProxyGetAllAnonymous() ([]Proxy, error) {
	var proxies []Proxy
	err := a.
		db.
		Model(&proxies).
		Where("anon = true").
		Select()
	return proxies, err
}

// ProxyGetAllAnonymousScheme - get all anonymous proxies by scheme
func (a *ADB) ProxyGetAllAnonymousScheme(v string) ([]Proxy, error) {
	var proxies []Proxy
	err := a.
		db.
		Model(&proxies).
		Where("anon = true AND scheme = ?", v).
		Select()
	return proxies, err
}

// ProxyGetUniqueHosts - gel all unique hosts
func (a *ADB) ProxyGetUniqueHosts() ([]string, error) {
	var hosts []string
	_, err := a.
		db.
		Query(&hosts, "SELECT DISTINCT host FROM proxies")
	return hosts, err
}

// ProxyGetFrequentlyUsedPorts - get 10 frequently used ports
func (a *ADB) ProxyGetFrequentlyUsedPorts() ([]string, error) {
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
	return ports, err
}

// ProxyInsert - insert new proxy
func (a *ADB) ProxyInsert(p Proxy) error {
	_, err := a.
		db.
		Model(&p).
		Insert(&p)
	return err
}

// ProxyUpdate - update existing proxy
func (a *ADB) ProxyUpdate(p Proxy) error {
	_, err := a.
		db.
		Model(&p).
		Where("hostname = ?", p.Hostname).
		Update(&p)
	return err
}
