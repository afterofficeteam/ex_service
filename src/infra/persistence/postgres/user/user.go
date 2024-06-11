package user

import (
	dto "ex_service/src/app/dto/user"

	"log"

	"github.com/jmoiron/sqlx"

	"errors"
)

type UserRepository interface {
	Register(data *dto.RegisterReqDTO) (*dto.RegisterRespDTO, error)
	Login(data *dto.LoginReqDTO) (*dto.RegisterRespDTO, error)
}

const (
	Register = `INSERT INTO public.users (username) 
		values ($1) returning id, username`

	Login = `select u.id, u.username, w.id as ex_service_id 
	from public.users u
	JOIN public.ex_services w
	ON u.id = w.user_id
	where u.username = $1`

	CreateWallet = `INSERT INTO public.wallet (user_id) 
	values ($1) returning id as wallet_id`
)

var statement PreparedStatement

type PreparedStatement struct {
	login *sqlx.Stmt
}

type userRepo struct {
	Connection *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	repo := &userRepo{
		Connection: db,
	}
	InitPreparedStatement(repo)
	return repo
}

func (p *userRepo) Preparex(query string) *sqlx.Stmt {
	statement, err := p.Connection.Preparex(query)
	if err != nil {
		log.Fatalf("Failed to preparex query: %s. Error: %s", query, err.Error())
	}

	return statement
}

func InitPreparedStatement(m *userRepo) {
	statement = PreparedStatement{
		login: m.Preparex(Login),
	}
}

func (p *userRepo) Register(data *dto.RegisterReqDTO) (*dto.RegisterRespDTO, error) {
	var resultData dto.RegisterRespDTO

	tx, err := p.Connection.Beginx()
	if err != nil {
		log.Println("Failed Begin Tx Register  : ", err.Error())
		return nil, err
	}

	// to handle rollback & commit
	defer func(tx *sqlx.Tx) {
		if err != nil {
			tx.Rollback()
			log.Println("Rolling back transaction due to:", err)
		} else {
			err = tx.Commit()
			if err != nil {
				log.Println("Failed to commit transaction:", err.Error())
			}
		}
	}(tx)

	err = tx.QueryRow(Register, data.UserName).Scan(&resultData.ID, &resultData.UserName)

	if err != nil {
		log.Println("Failed Query Register: ", err.Error())
		return nil, err
	}

	err = tx.QueryRow(CreateWallet, resultData.ID).Scan(&resultData.WalletID)

	if err != nil {
		log.Println("Failed Query Create ex_service : ", err.Error())
		return nil, err
	}

	return &resultData, nil
}

func (p *userRepo) Login(data *dto.LoginReqDTO) (*dto.RegisterRespDTO, error) {
	var resultData []*dto.RegisterRespDTO

	err := statement.login.Select(&resultData, data.UserName)

	if err != nil {
		return nil, err
	}

	if len(resultData) < 1 {
		return nil, errors.New("no rows returned from the query")
	}

	if err != nil {
		return nil, err
	}

	return resultData[0], nil
}
