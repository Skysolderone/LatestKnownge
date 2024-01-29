package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)

var userMap = map[string]UserInfo{}

const (
	githubUri         = "https://github.com/login/oauth/authorize"
	githubAccessToken = "https://github.com/login/oauth/access_token"
	githubUserApi     = "https://api.github.com/user"
	redirectUri       = "http://localhost:8080/token" //地址必须注册到github的配置中
	clientID          = ""                            //TODO 填写自己的clientID
	clientSecret      = ""                            //TODO 填写自己的clientSecret
	sessionKey        = "test"
)

func main() {
	userMap = make(map[string]UserInfo)
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		uri := fmt.Sprintf("%s?clietnt_id=%sredirect_uri=%s", githubUri, clientID, redirectUri)
		http.Redirect(w, r, uri, http.StatusFound)
	})
	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		httpClient := http.Client{}
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		code := r.FormValue("code")
		reqUrl := fmt.Sprintf("%s?&client_id=%s&client_secret=%s&code=%s", githubAccessToken, clientID, clientSecret, code)
		req, err := http.NewRequest(http.MethodPost, reqUrl, nil)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		req.Header.Set("accept", "application/json")
		res, err := httpClient.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer func() {
			_ = res.Body.Close()
		}()

		var t OAuthAccessResponse
		if err = json.NewDecoder(res.Body).Decode(&t); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		cookie, err := genCookie(t.AccessToken)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{Name: sessionKey, Value: cookie, Path: "/", Domain: "localhost", Expires: time.Now().Add(time.Second * 3600)})

		w.Header().Set("Location", "/welcome.html")
		w.WriteHeader(http.StatusFound)
	})
	http.HandleFunc("/userinfo", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(sessionKey)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		bytes, err := json.Marshal(userMap[cookie.Value])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, _ = w.Write(bytes)
	})
	_ = http.ListenAndServe(":8080", nil)
}

// 根据token获取userInfo，置换出自定义的cookie
func genCookie(token string) (string, error) {
	httpClient := http.Client{}
	req, err := http.NewRequest(http.MethodGet, githubUserApi, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "token "+token)
	res, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	cookie := uuid.NewString()
	var userInfo UserInfo
	if err = json.Unmarshal(bytes, &userInfo); err != nil {
		return "", err
	}

	userMap[cookie] = userInfo
	return cookie, nil
}

type OAuthAccessResponse struct {
	AccessToken string `json:"access_token"`
}
