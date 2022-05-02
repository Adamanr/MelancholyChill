package HandlerRoot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Handler(r *gin.Engine) {
	r = gin.New()
	r.Use(gin.Logger())
	defer r.Run()

	r.Static("/assets", "client/assets")
	r.LoadHTMLGlob("client/templates/**/*.tmpl")

	/*****************Routing************************/

	//Главная страница
	r.GET("/", mainPage)

	//Страница профиля
	r.GET("/profile", profile)

	//О нас
	r.GET("/information", information)

	//Библиотека
	r.Group("/library")
	{
		r.GET("/library-page", libraryPage)
		r.POST("/library-add", libraryAdd)
	}

	//Плейлисты
	r.Group("/playlist")
	{
		r.GET("/playlist-page", playlistPage)
		r.POST("/playlist-add", playlistAdd)
	}

	//Авторизация
	r.Group("/auth")
	{
		r.GET("/auth-page", authPage)
		r.POST("/auth-answer", authAnswer)
	}

	//Регистрация
	r.Group("/register")
	{
		r.GET("/register-page", registerPage)
		r.POST("/register-answer", registerAnswer)
	}

	//Обработка несуществующей страницы
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "404.tmpl", gin.H{
			"title": "Страницы не существует",
		})
	})

}

func init() {
}

var db, _ = sqlx.Connect("postgres", "user=# password=# dbname=# sslmode=disable")

//Переменная подключения к БД

type User struct {
	Id       int
	Name     string
	Login    string
	Password string
}

//Переменная в которую добавляется музыка из папки и выводится
var musicList []string

//Функция поиска музыки формата .mp3 и добавляющая её в массив musicList для дальнейшего вывода
func listTrack() {
	files, err := ioutil.ReadDir("client/assets/files/")
	if err != nil {
		log.Fatalln("Папка не читается")
	}
	for _, file := range files {
		fmt.Println(file.Name())
		fileName := strings.Trim(file.Name(), ".mp3")
		musicList = append(musicList, fileName)
	}
}

/********************************Страницы***********************/

//Главная страница
func mainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Главная страница",
	})
}

//Страница плейлистов
func playlistPage(c *gin.Context) {
	c.HTML(http.StatusOK, "playlist.tmpl", gin.H{
		"title": "Главная страница",
	})
}

//Метод для добавления плейлиста
func playlistAdd(c *gin.Context) {
	//Данные для добавления плейлиста
}

//Страница "О нас"
func information(c *gin.Context) {
	c.HTML(http.StatusOK, "information.tmpl", gin.H{
		"title": "О нас",
	})
}

var userInfo string

//Страница авторизации
func authPage(c *gin.Context) {
	c.HTML(http.StatusOK, "authPage.tmpl", gin.H{
		"title":     "Главная страница",
		"checkAuth": userInfo,
	})

}

//Метод отправки запроса для проверки на существование пользователя
func authAnswer(c *gin.Context) {
	login := c.PostForm("login")
	password := c.PostForm("password")
	u := User{}
	sql := fmt.Sprintf("SELECT * FROM users WHERE login = '%s' and password = '%s'", login, password)
	fmt.Println(login + ":" + password)
	if db.Get(&u, sql) == nil {
		c.Redirect(http.StatusMovedPermanently, "/")

	} else {
		userInfo = "Пользователя нет!"
		c.Redirect(http.StatusMovedPermanently, "/auth-page")
	}
}

//Страница регистрации
func registerPage(c *gin.Context) {
	userInfo = ""
	c.HTML(http.StatusOK, "registerPage.tmpl", gin.H{
		"title": "Регистрация",
	})
}

//Метод отправки данных для регистрации
func registerAnswer(c *gin.Context) {
	name := c.PostForm("name")
	login := c.PostForm("login")
	password := c.PostForm("password")
	u := User{Name: name, Login: login, Password: password}
	_, err := db.NamedExec(`INSERT INTO users(name,login,password) VALUES (:name,:login,:password)`, &u)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/register-page")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/auth-page")
	}
	db.Close()

}

//Страница личного профиля
func profile(c *gin.Context) {
	c.HTML(http.StatusOK, "profile.tmpl", gin.H{
		"title": "Профиль",
	})
}

//Страница библиотеки
func libraryPage(c *gin.Context) {
	c.HTML(http.StatusOK, "library.tmpl", gin.H{
		"title": "Библиотека",
	})
}

var libraryList []string

//Страница добавления в библиотеку
func libraryAdd(c *gin.Context) {
	//Добавление в библиотеку
}
