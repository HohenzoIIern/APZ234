package server

import (
	"database/sql"
)

type Server struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	Cpu_count      int    `json:"cpu_count"`
	TotalDiskSpace int64  `json:"totalDiskSpace"`
}

type Disk struct {
	Id        int64 `json:"id"`
	Space     int64 `json:"space"`
	Server_id *int  `json:"server_id"`
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) ListServers() ([]*Server, error) {
	rows, err := s.Db.Query("SELECT q.server_id, name, cpu_count, coalesce(sum(space), 0)space FROM server q left join disk w on q.server_id=w.server_id group by q.server_id, name, cpu_count LIMIT 200")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*Server
	for rows.Next() {
		var c Server
		if err := rows.Scan(&c.Id, &c.Name, &c.Cpu_count, &c.TotalDiskSpace); err != nil {
			return nil, err
		}
		res = append(res, &c)
	}
	if res == nil {
		res = make([]*Server, 0)
	}
	return res, nil
}

func (s *Store) ListDisk() ([]*Disk, error) {
	rows, err := s.Db.Query("SELECT disk_id, space, server_id FROM disk LIMIT 200")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*Disk
	for rows.Next() {
		var c Disk
		if err := rows.Scan(&c.Id, &c.Space, &c.Server_id); err != nil {
			return nil, err
		}
		res = append(res, &c)
	}
	if res == nil {
		res = make([]*Disk, 0)
	}
	return res, nil
}

func (s *Store) CreateServer(c *Server) error {
	err := s.Db.QueryRow("INSERT INTO Server (name, cpu_count) VALUES ($1, $2) returning server_id", c.Name, c.Cpu_count).Scan(&c.Id)
	return err
}

func (s *Store) CreateDisk(c *Disk) error {
	err := s.Db.QueryRow("INSERT INTO Disk (Space) VALUES ($1) returning disk_id", c.Space).Scan(&c.Id)
	return err
}

func (s *Store) AddDiskToServer(server_id int, disk_id int) (Server, error) {
	var c Server
	if _, err := s.Db.Exec("update disk set server_id = (select Server_id from server where Server_id=$1) where disk_id=$2 returning server_id", server_id, disk_id); err != nil {
		return c, err
	}
	err := s.Db.QueryRow(`select server_id, name, cpu_count, coalesce((select sum(space)
	from disk where q.Server_id=Server_id), 0)space from server q where server_id=$1`, server_id).Scan(&c.Id, &c.Name, &c.Cpu_count, &c.TotalDiskSpace)
	return c, err
}
