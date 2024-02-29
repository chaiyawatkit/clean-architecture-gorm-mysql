package controllers

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
)

func BindJSON(c *gin.Context, request any) (err error) {
	buf := make([]byte, 5120)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	c.Request.Body = io.NopCloser(bytes.NewBuffer([]byte(reqBody)))
	err = c.ShouldBindJSON(request)
	c.Request.Body = io.NopCloser(bytes.NewBuffer([]byte(reqBody)))
	return
}

func BindJSONMap(c *gin.Context, request *map[string]any) (err error) {
	buf := make([]byte, 5120)
	num, _ := c.Request.Body.Read(buf)
	reqBody := buf[0:num]
	c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
	err = json.Unmarshal(reqBody, &request)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
	return
}
