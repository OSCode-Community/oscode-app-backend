package controllers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	middlewares "github.com/OSCode-Community/oscode-app-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func GoogleLogin(c *gin.Context) {
	if c.Request.Method != "GET" {
		c.AbortWithStatus(http.StatusMethodNotAllowed)
		return
	}

	oauthState := GenerateStateOauthCookie(c)

	u := middlewares.AppConfig.GoogleLoginConfig.AuthCodeURL(oauthState)
	c.Redirect(http.StatusTemporaryRedirect, u)
}

func GoogleCallback(c *gin.Context) {
	if c.Request.Method != "GET" {
		c.AbortWithStatus(http.StatusMethodNotAllowed)
		return
	}

	oauthState, _ := c.Cookie("oauthstate") // handle error here
	state := c.Request.FormValue("state")
	code := c.Request.FormValue("code")

	c.Header("Content-Type", "application/json")

	if state != oauthState {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		c.JSON(http.StatusOK, gin.H{
			"error": "invalid oauth google state",
		})
		return
	}

	token, err := middlewares.AppConfig.GoogleLoginConfig.Exchange(context.Background(), code)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": fmt.Sprintf("failed code exchange: %s", err.Error()),
		})
		return
	}

	response, err := http.Get(middlewares.OauthGoogleUrlAPI + token.AccessToken)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": fmt.Sprintf("failed getting user info: %s", err.Error()),
		})
		return
	}

	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": fmt.Sprintf("failed read response: %s", err.Error()),
		})
		return
	}

	var userData map[string]interface{}
	err = json.Unmarshal(contents, &userData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": fmt.Sprintf("failed parse user data: %s", err.Error()),
		})
		return
	}

	c.Set("user_data", userData)
	c.JSON(http.StatusOK, userData)
}

func GenerateStateOauthCookie(c *gin.Context) string {
	var expiration = 3600
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	c.SetCookie("oauthstate", state, expiration, "/", "", true, true) // fix params here
	return state
}
