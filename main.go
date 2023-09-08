package main

import (
	mailsender "mailsender/pkg/mailSender"
	"net/http"

	"github.com/gin-gonic/gin"
)

type sendEmail struct {
	Subject string   `json:"subject"`
	Content string   `json:"content"`
	To      []string `json:"to"`
	Cc      []string `json:"cc"`
	Bcc     []string `json:"bcc"`
	Files   []string `json:"files"`
}

func main() {
	r := gin.Default()

	sender := mailsender.NewGmailSender("Sender Name", "Sender Email", "Sender App Password")

	// get the app password from google account management section after enabling 2 step verification

	r.GET("/", Home)
	r.POST("/send-email", func(c *gin.Context) {
		SendEmail(sender, c)
	})

	r.Run(":8080")

}

func Home(c *gin.Context) {

	input := sendEmail{}

	c.JSON(http.StatusOK, input)
}

func SendEmail(sender mailsender.EmailSender, c *gin.Context) {
	input := sendEmail{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := sender.SendEmail(input.Subject, input.Content, input.To, input.Cc, input.Bcc, input.Files); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, "Email send succesfully")
}
