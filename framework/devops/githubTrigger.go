package devops

import (
	"context"

	"github.com/caicloud/nubela/logger"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// 完成 Github Branch push 触发
func GithubBranchPush(githubToken, githubBranch, githubOwner, githubRepo, githubInitName string) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	b := []byte{1}
	fileOpt := &github.RepositoryContentFileOptions{
		Message: github.String("自动化脚本创建"),
		Branch:  github.String(githubBranch),
		Content: b,
	}
	creFile, _, err := client.Repositories.CreateFile(ctx, githubOwner, githubRepo, githubInitName, fileOpt)
	if err == nil {
		logger.Infof("创建File成功")
	} else {
		return err
	}
	delSHA := creFile.Content.SHA

	delFileOpt := &github.RepositoryContentFileOptions{
		Branch:  github.String(githubBranch),
		SHA:     delSHA,
		Message: github.String("自动化脚本删除"),
	}
	_, _, err = client.Repositories.DeleteFile(ctx, githubOwner, githubRepo, githubInitName, delFileOpt)
	if err == nil {
		logger.Infof("删除File成功")
	} else {
		return err
	}
	return nil
}

// 完成 Github PR Commit 触发
func GithubPRCommit(githubToken, githubOwner, githubRepo string) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	PROpt := &github.NewPullRequest{
		Title: github.String("自动化脚本创建"),
		Head:  github.String("pr2"),
		Base:  github.String("pr1"),
	}

	_, _, err := client.PullRequests.Create(ctx, githubOwner, githubRepo, PROpt)
	if err == nil {
		logger.Infof("创建PR成功，测试结束必须手动关闭PR")
	} else {
		return err
	}
	return nil
}

// 完成 Github PR Comment 触发
func GithubPRComment(githubToken, githubOwner, githubRepo string) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	CMTOpt := &github.PullRequestComment{
		CommitID: github.String("42b71bcb1a092ceab92bf6c42543efeae9a6db96"),
		Path:     github.String("中文测试"),
		Position: github.Int(1),
		Body:     github.String("cntest"),
	}

	creCom, _, err := client.PullRequests.CreateComment(ctx, githubOwner, githubRepo, 1, CMTOpt)
	if err == nil {
		logger.Infof("创建 PR Comment 成功")
	} else {
		return err
	}
	PRID := *creCom.ID

	_, err = client.PullRequests.DeleteComment(ctx, githubOwner, githubRepo, PRID)
	if err == nil {
		logger.Infof("删除 PR Comment 成功")
	} else {
		return err
	}
	return nil
}
