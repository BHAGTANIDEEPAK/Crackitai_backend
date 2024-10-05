package main

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Interview struct {
	Email      string `json:"Email"`
	Position   string `json:"Position"`
	Techstack  string `json:"Techstack"`
	Experience int    `json:"Experience"`
}

var Interviews = []Interview{
	{Email: "deepak@gmail.com", Position: "Front End Developer", Techstack: "Javascript", Experience: 1},
	{Email: "deepak@gmail.com", Position: "Backend End Developer", Techstack: "Python", Experience: 1},
	{Email: "ashutosh@gmail.com", Position: "Full Stack Developer", Techstack: "Javascript, Nodejs, cloud", Experience: 1},
}

func getInterviews(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, Interviews)
}

func addInterview(context *gin.Context) {
	var newInterview Interview

	if err := context.BindJSON(&newInterview); err != nil {
		return
	}

	Interviews = append(Interviews, newInterview)

	context.IndentedJSON(http.StatusCreated, Interviews)
}

func getInterviewByEmail(email string) ([]*Interview, error) {
	var matchingInterviews []*Interview

	for i, t := range Interviews {
		if t.Email == email {
			matchingInterviews = append(matchingInterviews, &Interviews[i])
		}
	}

	if len(matchingInterviews) == 0 {
		return nil, errors.New("Interview history not found")
	}

	return matchingInterviews, nil

}

func getInterview(context *gin.Context) {
	email := context.Param("email")
	interview, error := getInterviewByEmail(email)

	if error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Interview history not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, interview)

}

func main() {
	router := gin.Default()

	//Cors Configuration
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	router.GET("/getInterviews", getInterviews)
	router.GET("/getInterviews/:email", getInterview)
	router.POST("/addInterview", addInterview)
	router.Run("localhost:9090")
}
