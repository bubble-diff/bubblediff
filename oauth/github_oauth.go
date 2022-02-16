package oauth

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/bubble-diff/bubblediff/config"
	"github.com/bubble-diff/bubblediff/db"
	"github.com/bubble-diff/bubblediff/models"
)

// GitHubOAuth GitHub OAuth过程
//
// 请求应携带code授权码，若有效(err == nil)，则返回对应User信息。
func GitHubOAuth(c *gin.Context) (user interface{}, err error) {
	accessToken, err := exchange(c.Query("code"))
	if err != nil {
		log.Printf("get access_token failed, %s", err)
		return nil, err
	}
	if len(accessToken) == 0 {
		log.Printf("get empty access_token, code may be invalid.")
		return nil, errors.New("invalid code")
	}

	user, err = getUser(accessToken)
	if err != nil {
		log.Printf("get userinfo failed, %s", err)
		return nil, err
	}

	err = upsertUser(user.(*models.User))
	if err != nil {
		log.Printf("upsert user failed, %s", err)
		return nil, err
	}

	return user, nil
}

type exchangeReq struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
}

type exchangeResp struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

// exchange 使用code获取用户的github access_token，code使用后失效。
func exchange(code string) (accessToken string, err error) {
	conf := config.Get()

	dataReq := &exchangeReq{
		ClientID:     conf.ClientId,
		ClientSecret: conf.ClientSecret,
		Code:         code,
	}
	raw, err := json.Marshal(dataReq)
	if err != nil {
		return "", err
	}

	// retrieve access_token for this user.
	req, err := http.NewRequest(
		http.MethodPost,
		"https://github.com/login/oauth/access_token",
		bytes.NewReader(raw),
	)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	var dataResp exchangeResp
	err = json.NewDecoder(resp.Body).Decode(&dataResp)
	if err != nil {
		return "", err
	}
	resp.Body.Close()

	return dataResp.AccessToken, nil
}

// getUser 调用github api获取用户信息
func getUser(token string) (user *models.User, err error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/user", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))
	req.Header.Set("Accept", "application/json")
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var data models.User
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &data, nil
}

// upsertUser 向mongodb插入user数据，如user已存在，则更新信息。
func upsertUser(user *models.User) (err error) {
	log.Printf("upsert user:\n%+v", user)
	ctx := context.Background()
	coll := db.Mongodb.Database("bubblediff_test").Collection("user")

	filter := bson.D{{"id", user.ID}}
	update := bson.D{
		bson.E{Key: "$set", Value: bson.D{
			{"login", user.Login},
			{"avatar_url", user.AvatarUrl},
			{"email", user.Email},
		}},
	}
	opts := options.Update().SetUpsert(true)
	_, err = coll.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Printf("updateOne failed, %s", err)
		return err
	}
	return nil
}
