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
func GetAll() ([]Proxy, error) {
	var proxies []Proxy
	rows, err := pool.Query(context.Background(), `
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
		errmsg("GetAll Query", err)
		return proxies, err
	}
	for rows.Next() {
		var proxy Proxy
		err := rows.Scan(&proxy.IsWork, &proxy.IsAnon, &proxy.Checks, &proxy.Hostname,
			&proxy.Host, &proxy.Port, &proxy.Scheme, &proxy.CreateAt, &proxy.UpdateAt,
			&proxy.Response)
		if err != nil {
			errmsg("GetAll Scan", err)
			return proxies, err
		}
		proxies = append(proxies, proxy)
	}
	return proxies, rows.Err()
}

// // ProxyGetByID - get proxy by id
// func GetByID(id int64) (Proxy, error) {
// 	var proxy Proxy
// 	err := pool.QueryRow(context.Background(), `
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

// GetAllCount - get count of proxy
func GetAllCount() int64 {
	var count int64
	err := pool.QueryRow(context.Background(), `
		SELECT
			COUNT(*)
		FROM
			proxies
	`).Scan(&count)
	if err != nil {
		errmsg("GetAllCount QueryRow", err)
	}
	return count
}

// GetAllWorkCount - get count of working proxy
func GetAllWorkCount() int64 {
	var count int64
	err := pool.QueryRow(context.Background(), `
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

// GetAllAnonymousCount - get count of anonymous proxy
func GetAllAnonymousCount() int64 {
	var count int64
	err := pool.QueryRow(context.Background(), `
		SELECT
			COUNT(*)
		FROM
			proxies
		WHERE
			work = TRUE AND anon = TRUE
	`).Scan(&count)
	if err != nil {
		errmsg("GetAllAnonymousCount QueryRow", err)
	}
	return count
}

// GetAllSchemeCount - get count of proxies by scheme
func GetAllSchemeCount(scheme string) int64 {
	var count int64
	err := pool.QueryRow(context.Background(), `
		SELECT
			COUNT(*)
		FROM
			proxies
		WHERE
			scheme = $1
	`, scheme).Scan(&count)
	if err != nil {
		errmsg("GetAllSchemeCount QueryRow", err)
	}
	return count
}

// GetAllOldCount - get count of all old proxies
func GetAllOldCount() int64 {
	var count int64
	err := pool.QueryRow(context.Background(), `
		SELECT
			COUNT(*)
		FROM
			proxies
		WHERE
			work = true OR update_at < NOW() - (INTERVAL '3 days') * checks"
	`).Scan(&count)
	if err != nil {
		errmsg("GetAllOldCount QueryRow", err)
	}
	return count
}

// GetAllWorkingSchemeCount - get count of working proxies by scheme
func GetAllWorkingSchemeCount(scheme string) int64 {
	var count int64
	err := pool.QueryRow(context.Background(), `
		SELECT
			COUNT(*)
		FROM
			proxies
		WHERE
			work = true AND scheme = $1
	`, scheme).Scan(&count)
	if err != nil {
		errmsg("GetAllWorkingSchemeCount QueryRow", err)
	}
	return count
}

// GetAllAnonymousSchemeCount - get count of anonymous proxies by scheme
func GetAllAnonymousSchemeCount(scheme string) int64 {
	var count int64
	err := pool.QueryRow(context.Background(), `
		SELECT
			COUNT(*)
		FROM
			proxies
		WHERE
			anon = true AND scheme = $1	
	`, scheme).Scan(&count)
	if err != nil {
		errmsg("GetAllAnonymousSchemeCount QueryRow", err)
	}
	return count
}

// // ProxyGetAllScheme - get all proxies by scheme
// func (a *ADB) ProxyGetAllScheme(v string) ([]Proxy, error) {
// 	var proxies []Proxy
// 	err := a.
// 		db.
// 		Model(&proxies).
// 		Where("scheme = ?", v).
// 		Select()
// 	return proxies, err
// }

// GetAllOld - get all old proxies
func GetAllOld() ([]string, error) {
	var proxies []string
	rows, err := pool.Query(context.Background(), `
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
func GetAllWorking() ([]string, error) {
	var proxies []string
	rows, err := pool.Query(context.Background(), `
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
func GetAllWorkingScheme(scheme string) ([]string, error) {
	var proxies []string
	rows, err := pool.Query(context.Background(), `
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
func GetAllAnonymous() ([]string, error) {
	var proxies []string
	rows, err := pool.Query(context.Background(), `
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
func GetAllAnonymousScheme(scheme string) ([]string, error) {
	var proxies []string
	rows, err := pool.Query(context.Background(), `
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
func GetUniqueHosts() ([]string, error) {
	var hosts []string
	rows, err := pool.Query(context.Background(), `
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
func GetFrequentlyUsedPorts() ([]int64, error) {
	var ports []int64
	rows, err := pool.Query(context.Background(), `
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
func Insert(p *Proxy) error {
	_, err := pool.Exec(context.Background(), `
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
func Update(p *Proxy) error {
	_, err := pool.Exec(context.Background(), `
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
func GetRandomWorking(n int) ([]string, error) {
	var proxies []string
	rows, err := pool.Query(context.Background(), `
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
