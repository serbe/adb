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

// ProxyGetByID - get proxy by id
func (a *ADB) ProxyGetByID(id int64) (Proxy, error) {
	var proxy Proxy
	err := a.
		db.
		Model(&proxy).
		Where("id = ?", id).
		Select()
	return proxy, err
}

// ProxyGetAllCount - get count of proxy
func (a *ADB) ProxyGetAllCount() int64 {
	var proxies []Proxy
	c, _ := a.
		db.
		Model(&proxies).
		Count()
	return int64(c)
}

// ProxyGetAllWorkCount - get count of working proxy
func (a *ADB) ProxyGetAllWorkCount() int64 {
	var proxies []Proxy
	c, _ := a.
		db.
		Model(&proxies).
		Where("work = TRUE").
		Count()
	return int64(c)
}

// ProxyGetAllAnonymousCount - get count of anonimous proxy
func (a *ADB) ProxyGetAllAnonymousCount() int64 {
	var proxies []Proxy
	c, _ := a.
		db.
		Model(&proxies).
		Where("work = TRUE AND anon = TRUE").
		Count()
	return int64(c)
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

// ProxyGetFrequentlyUsedPorts - get 20 frequently used ports
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
			LIMIT 20
		`)
	return ports, err
}

// ProxyInsert - insert new proxy
func (a *ADB) ProxyInsert(p *Proxy) error {
	_, err := a.
		db.
		Model(p).
		Insert(p)
	return err
}

// ProxyUpdate - update existing proxy
func (a *ADB) ProxyUpdate(p *Proxy) error {
	_, err := a.
		db.
		Model(p).
		Where("hostname = ?", p.Hostname).
		Update(p)
	return err
}

// ProxyGetRandomWorking - get n random working proxies
func (a *ADB) ProxyGetRandomWorking(n int) ([]string, error) {
	var proxies []string
	_, err := a.
		db.
		Query(&proxies, `
			SELECT
				hostname
			FROM
				proxies
			WHERE
				work = true
			ORDER BY
				random()
			LIMIT
				?
		`, n)
	return proxies, err
}

// ProxyGetRandomAnonymous - get n random anonymous proxies
func (a *ADB) ProxyGetRandomAnonymous(n int) ([]string, error) {
	var proxies []string
	_, err := a.
		db.
		Query(&proxies, `
			SELECT
				hostname
			FROM
				proxies
			WHERE
				work = true AND anon = true
			ORDER BY
				random()
			LIMIT
				?
		`, n)
	return proxies, err
}
