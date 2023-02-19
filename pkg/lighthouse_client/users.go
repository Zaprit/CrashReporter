package lighthouse_client

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/Zaprit/CrashReporter/pkg/config"
    "github.com/Zaprit/CrashReporter/pkg/model"
    "io"
    "net/http"
)

//TODO: write user search API in Lighthouse

func GetUserRaw(name string) ([]byte,int, error) {
    resp, err := http.Get(fmt.Sprintf("%s/api/v1/username/%s", config.LoadedConfig.LighthouseURL, name))
    if err != nil {
        return nil, 500, err
    }

    if resp.StatusCode != 200 {
        return nil, resp.StatusCode, errors.New(resp.Status)
    }

    body, err := io.ReadAll(resp.Body)
    defer resp.Body.Close()

    if err != nil {
        return nil, resp.StatusCode, err
    }

    return body, resp.StatusCode, nil
}

func GetUser(name string) (model.LighthouseUser, error) {
    jsonData, err := GetUserRaw(name)

    var user model.LighthouseUser

    err = json.Unmarshal(jsonData, &user)

    return user, err
}