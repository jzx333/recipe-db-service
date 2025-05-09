package main

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"webServer/internal/recipes/repo"
	"webServer/internal/server"
)

//nTag := dto.NewTag{
//	"Хайп",
//	"☠️",
//}
//
//res, err := r.TagCreate(ctx, &nTag)
//if err != nil {
//	log.Fatalf("Error creating tag: %v", err)
//}
//fmt.Println(res)

//ingr := []dto.Ingredient{
//	{"Колесо", "4 штуки"},
//	{"Корпус", "1 штука"},
//	{"Пицца в багажник", "5 коробок"},
//}
//fmt.Println(res)
//
//steps := []dto.Step{
//	{1, "Сядь за руль"},
//	{2, "Кайфуй"},
//	{3, "Едь"},
//	{4, "Едь"},
//	{5, "Едь"},
//	{6, "Едь"},
//	{7, "Тормози, куда гонишь"},
//}
//
//tags := []int{
//	1,
//}
//
//test := dto.NewRecipe{
//	"Хонда Сивик", tags, 154, 9999,
//	250000, ingr, steps,
//	"https://autogeizer.ru/wp-content/uploads/2021/08/honda-civic.jpg",
//}
//
//res2, err := r.RecipeCreate(ctx, &test)
//if err != nil {
//	log.Fatalf("Error creating recipe: %v", err)
//}
//fmt.Println(res2)

//ingr = []dto.Ingredient{
//	{"Страшилка", "1 штука"},
//	{"Пугалка", "2 столовых ложки"},
//}
//
//steps = []dto.Step{
//	{1, "Страшно"},
//}
//
//tags = []int{
//	1, 2, 4,
//}
//
//test = dto.NewRecipe{
//	"Ёжиг", tags, 666, 111,
//	1500, ingr, steps,
//	"https://cs10.pikabu.ru/post_img/big/2019/11/06/8/1573044991158234281.jpg",
//}
//
//res3, err := r.RecipeCreate(ctx, &test)
//if err != nil {
//	log.Fatalf("Error creating recipe: %v", err)
//}
//fmt.Println(res3)

type Config struct {
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	DbName   string `env:"DB_NAME"`
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg := Config{}
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("Error parsing config: %v", err)
	}

	ctx := context.Background()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName)
	dbpool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	nativeDB := stdlib.OpenDBFromPool(dbpool)
	sqlxDB := sqlx.NewDb(nativeDB, "pgx")

	r := repo.NewRepo(sqlxDB)

	server.Server(ctx, r)

}
