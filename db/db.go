package db

import (
	"database/sql"

	go_ora "github.com/sijms/go-ora/v2"
)

type DbInfo struct {
	Username string
	Password string
}

type Produto struct {
	Produto  string `json:"produto"`
	Material string `json:"material"`
}
type Pessoas struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	AnoNascimento string `json:"anoNascimento"`
	Nivel         string `json:"nivel"`
}

func Conn() (*sql.DB, error) {
	serverInfo := DbInfo{
		Username: "ADMIN",
		Password: ".Hundecko20.",
	}
	urlOptions := map[string]string{
		"TRACE FILE": "trace.log",
		"SSL VERIFY": "FALSE",
	}
	connectString := "(description= (retry_count=20)(retry_delay=3)(address=(protocol=tcps)(port=1522)(host=adb.us-ashburn-1.oraclecloud.com))(connect_data=(service_name=g03605989f2ec5c_y12pb5ffee5jp7y5_medium.adb.oraclecloud.com))(security=(ssl_server_dn_match=yes)))"

	db, err := sql.Open("oracle", go_ora.BuildJDBC(serverInfo.Username, serverInfo.Password, connectString, urlOptions))

	if err != nil {
		// panic(err.Error())
		return nil, err
	}
	// defer db.Close()
	err = db.Ping()
	if err != nil {
		// log.Fatal(err)
		return nil, err
	}
	return db, nil
}

func GetPessoas(db *sql.DB) ([]Pessoas, error) {
	query := "select unique(NOME_FUNCIONARIO) from FEEDBACK"
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	var res []Pessoas
	for rows.Next() {
		var pessoa Pessoas
		if err := rows.Scan(
			&pessoa.Name,
		); err != nil {
			return nil, err
		}
		res = append(res, pessoa)
	}
	return res, nil
}

func GetProdutos(db *sql.DB) ([]Produto, error) {
	query := "select unique(PRODUTO) from FEEDBACK"
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	var res []Produto
	for rows.Next() {
		var produto Produto
		if err := rows.Scan(
			&produto.Produto,
		); err != nil {
			return nil, err
		}
		res = append(res, produto)
	}
	return res, nil
}
