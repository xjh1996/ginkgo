package devops

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/caicloud/nubela/logger"
)

//API文档https://docs.atlassian.com/bitbucket-server/rest/5.16.0/bitbucket-rest.html#idm8287391664

type BitbucketTagOptions struct {
	Name       string `json:"name"` //首字母大写，如果小写则表示不传出该参数
	StartPoint string `json:"startPoint"`
	Message    string `json:"message"`
}

type BitbucketProject struct {
	Key string `json:"key"`
}

type BitbucketRepo struct {
	Slug    string           `json:"slug"`
	Project BitbucketProject `json:"project"`
}

type BitbucketRef struct {
	Id         string        `json:"id"`
	Repository BitbucketRepo `json:"repository"`
}

type BitPRCommit struct {
	Title   string       `json:"title"`
	FromRef BitbucketRef `json:"fromRef"`
	ToRef   BitbucketRef `json:"toRef"`
}

const (
	BitbucketServer = "http://192.168.130.29:7990"
	BitbucketToken  = "Mjg2MDE3NDIzMzI3OuCRJ/bgsVe6pVmKXHrXCHZXA3HK"
	projectKey      = "CPPDEM"
	initName        = "cntest"
	commitID        = "e1c2bdbb1f152b33c65805df84a8976890e2e5ed"
	repo            = "cppdemo"
)

//TODO 完成 Bitbucket Branch Push 触发

//完成 Bitbucket Tag 触发
func BitbucketTag() {
	url := fmt.Sprintf("%s/rest/api/1.0/projects/%s/repos/%s/tags", BitbucketServer, projectKey, repo)
	method := "POST"
	nameNew := fmt.Sprintf("%s%d", initName, random())
	creTag := BitbucketTagOptions{
		Name:       nameNew,
		StartPoint: commitID,
		Message:    "自动化脚本创建",
	}

	bodyJson, _ := json.Marshal(creTag)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(bodyJson))

	if err != nil {
		logger.Failf("%v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+BitbucketToken)

	res, err := client.Do(req)
	if err != nil {
		logger.Failf("%v", err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	if body != nil {
		logger.Infof("创建Tag成功")
	}
}

//完成 Bitbucket PR Commit 触发
func BitbucketPRCommit() {
	url := fmt.Sprintf("%s/rest/api/1.0/projects/%s/repos/%s/pull-requests/", BitbucketServer, projectKey, repo)
	method := "POST"
	nameNew := fmt.Sprintf("%s%d", initName, random())

	ref1 := BitbucketRef{
		Id: "refs/heads/autoPr2",
		Repository: BitbucketRepo{
			Slug: repo,
			Project: BitbucketProject{
				Key: projectKey,
			},
		},
	}

	ref2 := BitbucketRef{
		Id: "refs/heads/autoPr1",
		Repository: BitbucketRepo{
			Slug: repo,
			Project: BitbucketProject{
				Key: projectKey,
			},
		},
	}

	crePRCommit := BitPRCommit{
		Title:   nameNew,
		FromRef: ref1,
		ToRef:   ref2,
	}

	bodyJson, _ := json.Marshal(crePRCommit)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(bodyJson))

	if err != nil {
		logger.Failf("%v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+BitbucketToken)

	res, err := client.Do(req)
	if err != nil {
		logger.Failf("%v", err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	if body != nil {
		logger.Infof("创建PR成功")
	}
	BitbucketDelPR(string(body)[6:8]) //TODO 后续优化成正则匹配
}

//完成 Bitbucket PR Comment 触发
func BitbucketPRComment() {
	url := fmt.Sprintf("%s/rest/api/1.0/projects/%s/repos/%s/pull-requests/23/comments", BitbucketServer, projectKey, repo)
	method := "POST"

	payload := strings.NewReader(`{"text": "cntest"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		logger.Failf("%v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+BitbucketToken)

	res, err := client.Do(req)
	if err != nil {
		logger.Failf("%v", err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	if body != nil {
		logger.Infof("创建PR Comment成功")
	}
}

//完成 Bitbucket PR 的删除
func BitbucketDelPR(id string) {
	url := fmt.Sprintf("%s/rest/api/1.0/projects/%s/repos/%s/pull-requests/%s", BitbucketServer, projectKey, repo, id)
	method := "DELETE"

	payload := strings.NewReader(`{"version": 0}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		logger.Failf("%v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+BitbucketToken)

	res, err := client.Do(req)
	if err != nil {
		logger.Failf("%v", err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	if body != nil {
		logger.Infof("删除PR成功")
	}
}
