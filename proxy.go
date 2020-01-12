package adb

import (
	"context"
	"time"
)

// Proxy - proxy unit
type Proxy struct {
	IsWork   bool          `sql:"work,notnull"        json:"-"`
	IsAnon   bool          `sql:"anon,notnull"        json:"-"`
	Checks   int           `sql:"checks,notnull"      json:"-"`
	Hostname string        `sql:"hostname,pk,notnull" json:"hostname"`
	Host     string        `sql:"host,notnull"        json:"-"`
	Port     int           `sql:"port,notnull"        json:"-"`
	Scheme   string        `sql:"scheme,notnull"      json:"-"`
	CreateAt time.Time     `sql:"create_at,notnull"   json:"-"`
	UpdateAt time.Time     `sql:"update_at,notnull"   json:"-"`
	Response time.Duration `sql:"response,notnull"    json:"-"`
}

// GetAll - get all proxies
func (db *DB) GetAll() ([]Proxy, error) {
	var proxies []Proxy
	rows, err := db.Pool.Query(context.Background(), `
		SELECT
			work,
			anon,
			checks,
			hostname,
			host,
			port,
			scheme,
			create_at,
			update_at,
			response
		FROM
			proxies
	`)
	if err != nil {
		errmsg("GetProxyAll Query", err)
		return proxies, err
	}
	for rows.Next() {
		var proxy Proxy
		err := rows.Scan(&proxy.IsWork, &proxy.IsAnon, &proxy.Checks, &proxy.Hostname,
			&proxy.Host, &proxy.Port, &proxy.Scheme, &proxy.CreateAt, &proxy.UpdateAt,
			&proxy.Response)
		if err != nil {
			errmsg("GetProxyAll Scan", err)
			return proxies, err
		}
		proxies = append(proxies, proxy)
	}
	return proxies, rows.Err()
}

// // ProxyGetByID - get proxy by id
// func (db *DB) GetByID(id int64) (Proxy, error) {
// 	var proxy Proxy
// 	err := db.Pool.QueryRow(context.Background(), `
// 		SELECT
// 			work,
// 			anon,
// 			checks,
// 			hostname,
// 			host,
// 			port,
// 			scheme,
// 			create_at,
// 			update_at,
// 			response
// 		FROM
// 			proxies
// 		WHERE

// 	`)
// 	return proxy, err
// }

// GetCountAll - get count of proxy
func (db *DB) GetCountAll() int64 {
	var count int64
	err := db.Pool.QueryRow(context.Background(), `
		SELECT
			COUNT(*)
		FROM
			proxies
	`).Scan(&count)
	if err != nil {
		errmsg("GetCountAll QueryRow", err)
	}
	return count
}

// GetCountAllWork - get count of working proxy
func (db *DB) GetCountAllWork() int64 {
	var count int64
	err := db.Pool.QueryRow(context.Background(), `
		SELECT
			COUNT(*)
		FROM
			proxies
		WHERE
			work = true
	`).Scan(&count)
	if err != nil {
		errmsg("GetAllWorkCount QueryRow", err)
	}
	return count
}

// GetCountAllAnonymous - get count of anonymous proxy
func (db *DB) GetCountAllAnonymous() int64 {
	var count int64
	err := db.Pool.QueryRow(context.Background(), `
		SELECT
			COUNT(*)
		FROM
			proxies
		WHERE
			work = TRUE AND anon = TRUE
	`).Scan(&count)
	if err != nil {
		errmsg("GetCountAllAnonymous QueryRow", err)
	}
	return count
}

// GetCountAllScheme - get count of proxies by scheme
func (db *DB) GetCountAllScheme(scheme string) int64 {
	var count int64
	err := db.Pool.QueryRow(context.Background(), `
		SELECT
			COUNT(*)
		FROM
			proxies
		WHERE
			scheme = $1
	`, scheme).Scan(&count)
	if err != nil {
		errmsg("GetCountAllScheme QueryRow", err)
	}
	return count
}

// GetCountAllOld - get count of all old proxies
func (db *DB) GetCountAllOld() int64 {
	var count int64
	err := db.Pool.QueryRow(context.Background(), `
		SELECT
			COUNT(*)
		FROM
			proxies
		WHERE
			work = true OR update_at < NOW() - (INTERVAL '3 days') * checks"
	`).Scan(&count)
	if err != nil {
		errmsg("GetCountAllOld QueryRow", err)
	}
	return count
}

// GetCountAllWorkingScheme - get count of working proxies by scheme
func (db *DB) GetCountAllWorkingScheme(scheme string) int64 {
	var count int64
	err := db.Pool.QueryRow(context.Background(), `
		SELECT
			COUNT(*)
		FROM
			proxies
		WHERE
			work = true AND scheme = $1
	`, scheme).Scan(&count)
	if err != nil {
		errmsg("GetCountAllWorkingScheme QueryRow", err)
	}
	return count
}

// GetCountAllAnonymousScheme - get count of anonymous proxies by scheme
func (db *DB) GetCountAllAnonymousScheme(scheme string) int64 {
	var count int64
	err := db.Pool.QueryRow(context.Background(), `
		SELECT
			COUNT(*)
		FROM
			proxies
		WHERE
			anon = true AND scheme = $1	
	`, scheme).Scan(&count)
	if err != nil {
		errmsg("GetCountAllAnonymousScheme QueryRow", err)
	}
	return count
}

// GetAllScheme - get all proxies by scheme
func (db *DB) GetAllScheme(scheme string) ([]string, error) {
	var proxies []string
	rows, err := db.Pool.Query(context.Background(), `
		SELECT
			hostname
		FROM
			proxies
		WHERE
			scheme = $1	
	`, scheme)
	if err != nil {
		errmsg("GetAllScheme QueryRow", err)
	}
	for rows.Next() {
		var proxy string
		err := rows.Scan(&proxy)
		if err != nil {
			errmsg("GetAllScheme Scan", err)
			return proxies, err
		}
		proxies = append(proxies, proxy)
	}
	return proxies, rows.Err()
}

// GetAllOld - get all old proxies
func (db *DB) GetAllOld() ([]string, error) {
	var proxies []string
	rows, err := db.Pool.Query(context.Background(), `
		SELECT
			hostname
		FROM
			proxies
		WHERE
			work = true OR update_at < NOW() - (INTERVAL '3 days') * checks
	`)
	if err != nil {
		errmsg("GetAllOld Query", err)
		return proxies, err
	}
	for rows.Next() {
		var proxy string
		err := rows.Scan(&proxy)
		if err != nil {
			errmsg("GetAllOld Scan", err)
			return proxies, err
		}
		proxies = append(proxies, proxy)
	}
	return proxies, rows.Err()
}

// GetAllWorking - get all working proxies
func (db *DB) GetAllWorking() ([]string, error) {
	var proxies []string
	rows, err := db.Pool.Query(context.Background(), `
		SELECT
			hostname
		FROM
			proxies
		WHERE
			work = true
	`)
	if err != nil {
		errmsg("GetAllWorking Query", err)
		return proxies, err
	}
	for rows.Next() {
		var proxy string
		err := rows.Scan(&proxy)
		if err != nil {
			errmsg("GetAllWorking Scan", err)
			return proxies, err
		}
		proxies = append(proxies, proxy)
	}
	return proxies, rows.Err()
}

// GetAllWorkingScheme - get all working proxies by scheme
func (db *DB) GetAllWorkingScheme(scheme string) ([]string, error) {
	var proxies []string
	rows, err := db.Pool.Query(context.Background(), `
		SELECT
			hostname
		FROM
			proxies
		WHERE
			work = true ANS scheme = $1
	`, scheme)
	if err != nil {
		errmsg("GetAllWorkingScheme Query", err)
		return proxies, err
	}
	for rows.Next() {
		var proxy string
		err := rows.Scan(&proxy)
		if err != nil {
			errmsg("GetAllWorkingScheme Scan", err)
			return proxies, err
		}
		proxies = append(proxies, proxy)
	}
	return proxies, rows.Err()
}

// GetAllAnonymous - get all anonymous proxies
func (db *DB) GetAllAnonymous() ([]string, error) {
	var proxies []string
	rows, err := db.Pool.Query(context.Background(), `
		SELECT
			hostname
		FROM
			proxies
		WHERE
			anon = true
	`)
	if err != nil {
		errmsg("GetAllAnonymous Query", err)
		return proxies, err
	}
	for rows.Next() {
		var proxy string
		err := rows.Scan(&proxy)
		if err != nil {
			errmsg("GetAllAnonymous Scan", err)
			return proxies, err
		}
		proxies = append(proxies, proxy)
	}
	return proxies, rows.Err()
}

// GetAllAnonymousScheme - get all anonymous proxies by scheme
func (db *DB) GetAllAnonymousScheme(scheme string) ([]string, error) {
	var proxies []string
	rows, err := db.Pool.Query(context.Background(), `
		SELECT
			hostname
		FROM
			proxies
		WHERE
			anon = true AND scheme = $1
	`, scheme)
	if err != nil {
		errmsg("GetAllAnonymousScheme Query", err)
		return proxies, err
	}
	for rows.Next() {
		var proxy string
		err := rows.Scan(&proxy)
		if err != nil {
			errmsg("GetAllAnonymousScheme Scan", err)
			return proxies, err
		}
		proxies = append(proxies, proxy)
	}
	return proxies, rows.Err()
}

// GetUniqueHosts - gel all unique proxy
func (db *DB) GetUniqueHosts() ([]string, error) {
	var hosts []string
	rows, err := db.Pool.Query(context.Background(), `
		SELECT
			DISTINCT host
		FROM
			proxies
		WHERE
			work = true
	`)
	if err != nil {
		errmsg("GetUniqueHosts Query", err)
		return hosts, err
	}
	for rows.Next() {
		var host string
		err := rows.Scan(&host)
		if err != nil {
			errmsg("GetUniqueHosts Scan", err)
			return hosts, err
		}
		hosts = append(hosts, host)
	}
	return hosts, rows.Err()
}

// GetFrequentlyUsedPorts - get 20 frequently used ports
func (db *DB) GetFrequentlyUsedPorts() ([]int64, error) {
	var ports []int64
	rows, err := db.Pool.Query(context.Background(), `
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
	if err != nil {
		errmsg("GetFrequentlyUsedPorts Query", err)
		return ports, err
	}
	for rows.Next() {
		var port int64
		err := rows.Scan(&port)
		if err != nil {
			errmsg("GetFrequentlyUsedPorts Scan", err)
			return ports, err
		}
		ports = append(ports, port)
	}
	return ports, rows.Err()
}

// Insert - insert new proxy
func (db *DB) Insert(p *Proxy) error {
	_, err := db.Pool.Exec(context.Background(), `
		INSERT INTO
			proxies (
				work,
				anon,
				checks,
				hostname,
				host,
				port,
				scheme,
				create_at,
				update_at,
				response
			)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`, p.IsWork, p.IsAnon, p.Checks, p.Hostname, p.Host, p.Port, p.Scheme, p.CreateAt, p.UpdateAt, p.Response)
	return err
}

// Update - update existing proxy
func (db *DB) Update(p *Proxy) error {
	_, err := db.Pool.Exec(context.Background(), `
		UPDATE
			proxies
		SET
			work = $2,
			anon = $3,
			checks = $4,
			host = $5,
			port = $6,
			scheme = $7,
			create_at = $8,
			update_at = $9,
			response = &10
		WHERE
			hostname = $1
	`, p.Hostname, p.IsWork, p.IsAnon, p.Checks, p.Host, p.Port, p.Scheme, p.CreateAt, p.UpdateAt, p.Response)
	return err
}

// GetRandomWorking - get n random working proxies
func (db *DB) GetRandomWorking(n int) ([]string, error) {
	var proxies []string
	rows, err := db.Pool.Query(context.Background(), `
		SELECT
			hostname
		FROM
			proxies
		WHERE
			work = true
		ORDER BY
			random()
		LIMIT
			$1
	`, n)
	if err != nil {
		errmsg("GetRandomWorking Query", err)
		return proxies, err
	}
	for rows.Next() {
		var proxy string
		err := rows.Scan(&proxy)
		if err != nil {
			errmsg("GetRandomWorking Scan", err)
			return proxies, err
		}
		proxies = append(proxies, proxy)
	}
	return proxies, rows.Err()
}

// // ProxyGetRandomAnonymous - get n random anonymous proxies
// func (a *ADB) ProxyGetRandomAnonymous(n int) ([]string, error) {
// 	var proxies []string
// 	_, err := a.
// 		db.
// 		Query(&proxies, `
// 			SELECT
// 				hostname
// 			FROM
// 				proxies
// 			WHERE
// 				work = true AND anon = true
// 			ORDER BY
// 				random()
// 			LIMIT
// 				?
// 		`, n)
// 	return proxies, err
// }

// // CheckNotExists - check list of hostnames with not exist in base
// func (a *ADB) CheckNotExists(s []string) ([]string, error) {
// 	var (
// 		proxies []string
// 		err     error
// 		j       int64
// 	)
// 	if len(s) == 0 {
// 		return proxies, errors.New("Empty input")
// 	}
// 	var mapS = make(map[string]bool)
// 	for i := range s {
// 		mapS[s[i]] = true
// 	}
// 	count := a.ProxyGetAllCount()
// 	for j = 0; j < count/100000+1; j++ {
// 		var pr []string
// 		_, err = a.
// 			db.
// 			Query(&pr, `
// 				SELECT
// 					hostname
// 				FROM
// 					proxies
// 				ORDER BY
// 					id
// 				OFFSET
// 					?
// 				LIMIT
// 					100000
// 			`, j*100000)
// 		for i := range pr {
// 			delete(mapS, pr[i])
// 		}
// 	}
// 	for k := range mapS {
// 		proxies = append(proxies, k)
// 	}
// 	return proxies, err
// }
