package devops

import (
	"crypto/rand"

	"github.com/caicloud/nubela/logger"

	"github.com/xanzy/go-gitlab"

	"math/big"
	"net/http"
	"strconv"
)

const (
	gitLabServer = "http://192.168.130.29:81"
	v4APIVersion = "v4"
	pid          = 37
	branch       = "master"
	name         = "cntest"
	gitLabRepo   = "192.168.130.29:81/fufu/test/blob/master/"
	gitLabToken  = "c5o6YxuraCcrBtiAgGW4"
)

func random() int64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(10000))
	return n.Int64()
}

//完成 Gitlab Tag 自动触发
func GitlabTag() {
	//GO文档https://godoc.org/github.com/xanzy/go-gitlab
	//API文档https://docs.gitlab.com/ee/api
	nameNew := name + strconv.FormatInt(random(), 10)
	httpClient := &http.Client{}
	git, err := gitlab.NewClient(gitLabToken, gitlab.WithHTTPClient(httpClient), gitlab.WithBaseURL(gitLabServer+"/api/"+v4APIVersion))
	if err != nil {
		logger.Failf("v", err)
	}

	//tags := &gitlab.ListTagsOptions{}
	//tag, _, _ := git.Tags.ListTags(pid, tags)
	//logger.Infof("%v",tag)

	//TODO tag触发
	creTag := &gitlab.CreateTagOptions{
		TagName: gitlab.String(nameNew),
		Ref:     gitlab.String(branch),
		Message: gitlab.String("自动化脚本创建"),
	}
	_, _, err = git.Tags.CreateTag(pid, creTag)
	if err != nil {
		logger.Failf("%v", err)
	} else {
		logger.Infof("创建Tag成功")
	}

	_, err = git.Tags.DeleteTag(pid, nameNew)
	if err == nil {
		logger.Infof("删除Tag成功")
	} else {
		logger.Errorf("删除Tag失败")
	}
}

//完成 Gitlab Branch Push 触发
func GitlabBranchPush() {
	nameNew := name + strconv.FormatInt(random(), 10)
	httpClient := &http.Client{}
	git, err := gitlab.NewClient(gitLabToken, gitlab.WithHTTPClient(httpClient), gitlab.WithBaseURL(gitLabServer+"/api/"+v4APIVersion))
	if err != nil {
		logger.Failf("v", err)
	}

	//TODO Branch Push触发
	creFile := &gitlab.CreateFileOptions{
		Branch:        gitlab.String(branch),
		Content:       gitlab.String("自动化脚本创建"),
		CommitMessage: gitlab.String(nameNew),
	}
	_, _, err = git.RepositoryFiles.CreateFile(pid, nameNew, creFile)
	if err != nil {
		logger.Failf("%v", err)
	} else {
		logger.Infof("创建File成功")
	}

	delFile := &gitlab.DeleteFileOptions{
		Branch:        gitlab.String(branch),
		CommitMessage: gitlab.String("删除"),
	}
	_, err = git.RepositoryFiles.DeleteFile(pid, nameNew, delFile)
	if err == nil {
		logger.Infof("删除File成功")
	} else {
		logger.Errorf("删除File失败")
	}
}

//完成 Gitlab PR Commit 触发
func GitlabPRCommit() {
	nameNew := name + strconv.FormatInt(random(), 10)
	httpClient := &http.Client{}
	git, err := gitlab.NewClient(gitLabToken, gitlab.WithHTTPClient(httpClient), gitlab.WithBaseURL(gitLabServer+"/api/"+v4APIVersion))
	if err != nil {
		logger.Failf("v", err)
	}

	//TODO PR Commit触发
	fileAction := gitlab.FileCreate
	action := &gitlab.CommitActionOptions{
		Action:   &fileAction,
		FilePath: gitlab.String(gitLabRepo + nameNew),
	}
	var actionArr []*gitlab.CommitActionOptions
	actionArr = append(actionArr, action)
	creCommit := &gitlab.CreateCommitOptions{
		Branch:        gitlab.String(branch),
		CommitMessage: gitlab.String(nameNew),
		Actions:       actionArr,
	}
	_, _, err = git.Commits.CreateCommit(pid, creCommit)
	if err != nil {
		logger.Failf("%v", err)
	} else {
		logger.Infof("创建Commit成功")
	}
}

//完成 Gitlab PR Comment 触发
func GitlabPRComment() {
	httpClient := &http.Client{}
	git, err := gitlab.NewClient(gitLabToken, gitlab.WithHTTPClient(httpClient), gitlab.WithBaseURL(gitLabServer+"/api/"+v4APIVersion))
	if err != nil {
		logger.Failf("v", err)
	}

	//getPR := &gitlab.GetMergeRequestsOptions{}
	//getMerge, _, _ := git.MergeRequests.GetMergeRequest(37, 1, getPR)
	//logger.Infof("%v",getMerge)

	//TODO PR Comment触发
	creComment := &gitlab.CreateMergeRequestDiscussionOptions{
		Body: gitlab.String(name),
	}
	_, _, err = git.Discussions.CreateMergeRequestDiscussion(pid, 1, creComment)
	if err != nil {
		logger.Failf("%v", err)
	} else {
		logger.Infof("创建Comment成功")
	}
}
