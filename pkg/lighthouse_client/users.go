package lighthouse_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Zaprit/CrashReporter/pkg/config"
	"github.com/Zaprit/CrashReporter/pkg/model"
	"io"
	"net/http"
	"net/url"
	"strings"
)

//TODO: write user search API in Lighthouse

func GetUserRaw(name string) ([]byte, int, error) {
	resp, err := http.Get(fmt.Sprintf("%s/api/v1/username/%s", config.LoadedConfig.LighthouseURL, name))
	if err != nil {
		return nil, 500, err
	}

	if resp.StatusCode != 200 {
		return nil, resp.StatusCode, errors.New(resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()

	if err != nil {
		return nil, resp.StatusCode, err
	}

	return body, resp.StatusCode, nil
}

func GetUser(name string) (model.LighthouseUser, int, error) {
	jsonData, status, err := GetUserRaw(name)

	var user model.LighthouseUser

	err = json.Unmarshal(jsonData, &user)

	return user, status, err
}

func SearchUsers(query string) ([]model.LighthouseUser, error) {
	resp, err := http.Get(fmt.Sprintf("%s/api/v1/username?query=%s", config.LoadedConfig.LighthouseURL, url.QueryEscape(query)))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()

	if err != nil {
		return nil, err
	}

	var users []model.LighthouseUser

	err = json.Unmarshal(body, &users)

	return users, err
}

func UserAvatar(user model.LighthouseUser) string {
	avatarHash := user.IconHash

	if avatarHash == "" || strings.HasPrefix(avatarHash, "g") {
		avatarHash = user.YayHash
	}
	if avatarHash == "" {
		avatarHash = user.MehHash
	}
	if avatarHash == "" {
		avatarHash = user.BooHash
	}

	if avatarHash != "" {
		return fmt.Sprintf("%s/gameAssets/%s", config.LoadedConfig.LighthouseURL, avatarHash)
	} else {
		return "/static/img/missing.png"
	}
}
