package delivery

import (
	"database/sql"
	"field_work/config"
	"field_work/delivery/controller"
	"field_work/delivery/middleware"
	"field_work/repository"
	"field_work/shared/service"
	"field_work/usecase"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	usersUC    usecase.UsersUseCase
	authUC     usecase.AuthUseCase
	productsUC     usecase.ProductsUseCase
	jwtService service.JwtService
	engine     *gin.Engine
	host       string
}

func (s *Server) initRoute() {
	rg := s.engine.Group(config.ApiGroup)

	authMiddleware := middleware.NewAuthMiddleware(s.jwtService)
	controller.NewAuthController(s.authUC, rg).Route()
	controller.NewUsersController(s.usersUC, rg, authMiddleware).Route()
	controller.NewProductsController(s.productsUC, rg, authMiddleware).Route()
}

func (s *Server) Run() {
	s.initRoute()
	if err := s.engine.Run(s.host); err != nil {
		panic(fmt.Errorf("server not running on host %s, because error %v", s.host, err.Error()))
	}
}

func NewServer() *Server {
	cfg, _ := config.NewConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		panic(err.Error())
	}

	// service
	jwtService := service.NewJwtService(cfg.TokenConfig)

	// Repo
	usersRepo := repository.NewUsersRepository(db)
	productsRepo := repository.NewProductsRepository(db)
  
	// Usecase
	usersUC := usecase.NewUsersUseCase(usersRepo)
	authUC := usecase.NewAuthUseCase(usersUC, jwtService)
	productsUC := usecase.NewProductsUseCase(productsRepo)

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		usersUC:    usersUC,
		authUC:     authUC,
		productsUC:     productsUC,
		jwtService: jwtService,
		engine:     engine,
		host:       host,
	}
}
