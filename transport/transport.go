package transport

import (
	"github.com/jmoiron/sqlx"
	"github.com/uchupx/bpjs-test-golang/config"
	"github.com/uchupx/bpjs-test-golang/data"
	"github.com/uchupx/bpjs-test-golang/data/mysql"
	"github.com/uchupx/bpjs-test-golang/database"
)

type Transport struct {
	mysqlConn             *sqlx.DB
	transactionRepository data.TransactionRepository
	transactionHandler    *TransactionHandler
}

func (t Transport) GetMySQLConn(conf *config.Config) *sqlx.DB {
	if t.mysqlConn == nil {
		mysqlConfig := database.Config{
			HostName: conf.Host,
			Username: conf.Username,
			Database: conf.Database,
			Password: conf.Password,
		}

		conn, err := database.NewConnection(mysqlConfig)
		if err != nil {
			panic(err)
		}

		t.mysqlConn = conn
	}

	return t.mysqlConn
}

func (t Transport) GetTransactionRepo(conf *config.Config) data.TransactionRepository {
	if t.transactionRepository == nil {
		repo := mysql.NewTransactionMysql(t.GetMySQLConn(conf))

		t.transactionRepository = repo
	}

	return t.transactionRepository
}

func (t Transport) GetTransactionHandler(conf *config.Config) *TransactionHandler {
	if t.transactionHandler == nil {
		handler := TransactionHandler{
			TransactionRepository: t.GetTransactionRepo(conf),
		}

		t.transactionHandler = &handler
	}

	return t.transactionHandler
}
