package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Carlitonchin/Backend-Tesis/handler"
	"github.com/Carlitonchin/Backend-Tesis/repository"
	"github.com/Carlitonchin/Backend-Tesis/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func inject(data_source *dataSource) (*gin.Engine, error) {

	router := gin.Default()
	user_repo := repository.NewUserRepository(data_source.DB)
	token_repo := repository.NewTokenRepository(data_source.RedisClient)
	us_config := &service.USConfig{UserRepository: user_repo}

	user_serv := service.NewUserService(us_config)

	priv, err := ioutil.ReadFile(os.Getenv("PRIVATE_KEY_FILE"))

	if err != nil {
		log.Fatalf("Error when reading private key file, error: %v", err)
	}

	pub, err := ioutil.ReadFile(os.Getenv("PUBLIC_KEY_FILE"))

	if err != nil {
		log.Fatalf("Error when reading public key file, error: %v", err)
	}

	refresh_secret := os.Getenv("REFRESH_TOKEN_SECRET")

	priv_key, err := jwt.ParseRSAPrivateKeyFromPEM(priv)

	if err != nil {
		log.Fatalf("Error when parsing private key, error: %v", err)
	}

	pub_key, err := jwt.ParseRSAPublicKeyFromPEM(pub)

	if err != nil {
		log.Fatalf("Error when parsing public key, error: %v", err)
	}

	id_exp, err := strconv.ParseInt(os.Getenv("ID_TOKEN_EXP"), 0, 64)

	if err != nil {
		log.Fatalf("error when readinga and parsing id_token_exp, error: %v", err)
	}

	refresh_exp, err := strconv.ParseInt(os.Getenv("REFRESH_TOKEN_EXP"), 0, 64)

	if err != nil {
		log.Fatalf("error when readinga and parsing refresh_exp, error: %v", err)
	}

	handlerTimeout, err := strconv.ParseInt(os.Getenv("HANDLER_TIMEOUT"), 0, 64)

	if err != nil {
		log.Fatalf("Error when reading and parsing handlerTimeout, error:%v", err)
	}

	tsc := &service.TSConfig{
		PrivateKey:           priv_key,
		PublicKey:            pub_key,
		RefreshSecret:        refresh_secret,
		IDExpirationSec:      id_exp,
		RefreshExpirationSec: refresh_exp,
		TokenRepository:      token_repo,
	}

	token_service := service.NewTokenService(tsc)
	role_repo := repository.NewRoleRepository(data_source.DB)
	role_service := service.NewRoleService(role_repo)

	area_repo := repository.NewAreaRepository(data_source.DB)
	area_serv := service.NewAreaService(area_repo)

	question_repo := repository.NewQuestionRepository(data_source.DB)
	question_serv := service.NewQuestionService(question_repo)

	c := handler.Config{
		R:               router,
		UserService:     user_serv,
		TokenService:    token_service,
		TimeOut:         time.Duration(time.Duration(handlerTimeout) * time.Second),
		RoleService:     role_service,
		AreaService:     area_serv,
		QuestionService: question_serv,
	}

	handler.NewHandler(&c)

	return router, nil
}
