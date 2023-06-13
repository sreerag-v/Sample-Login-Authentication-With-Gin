package main

import (
	"net/http"

	// Dependency Libreries
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	Router := gin.Default()
	// the cookie store is responsable for managing session data by storing in it encrypted cookie on the clinte side
	store := cookie.NewStore([]byte("Secret"))

	// this line line sets the middleware for the Router
	Router.Use(sessions.Sessions("session", store))

	// set fuctuon for find the path of the diroictory and Render the Html templates when requested
	Router.LoadHTMLGlob("Templates/*")

	// this section for if the user is aldredy logged in redirct to the user to homepage otherwice redirect to the login page
	// handlerfunction
	Router.GET("/", func(C *gin.Context) {
		// to acces and mainpulate session data
		session := sessions.Default(C)
		// in this session the "username" key aldredy exist Redirect to homepage and Render homepage
		if session.Get("username") != nil {
			C.Redirect(http.StatusSeeOther, "/home")
			return
		}
		// if the user war not logged in this line will exicute and Render the Login Page
		C.HTML(http.StatusOK, "login.html", nil)
	})

	// Authentication and Validation Method

	Router.POST("/login", func(C *gin.Context) {
		session := sessions.Default(C)

		// inserting form values to these variables
		username := C.PostForm("username")
		password := C.PostForm("password")

		ValidUsername := "username"
		ValidPassword := "password"

		// check the values
		if username == ValidUsername && password == ValidPassword {
			// is these two are matche then store the values to session for future authentication
			session.Set("username", username)
			session.Save()
			// and render to the home page
			C.Redirect(http.StatusSeeOther, "/home")
		} else {
			// else show the error
			C.HTML(http.StatusOK, "login.html", gin.H{"Error": "invalid username and password"})
		}
	})

	//  make the function for the user to rener the Home page

	Router.GET("/home", func(C *gin.Context) {
		session := sessions.Default(C)
		// they check the user name and pass is in the session.if this not there they will rediret to the root page
		if session.Get("username") == nil {
			C.Redirect(http.StatusSeeOther, "/")
			return
		}
		// logout section and after logout clear the session data
		if C.Query("logout") == "true" {
			session.Clear()
			session.Save()

			C.Redirect(http.StatusSeeOther, "/")
		}

		// after this  all cheks it will render the home page
		C.HTML(http.StatusOK, "home.html", gin.H{
			"Username": session.Get("username"),
		})

	})

	Router.Run(":8080")
}
