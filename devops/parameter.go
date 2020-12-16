package devops

const (
	configPath  = "/Users/chenneng/Desktop/autotest/config"
	nameSpace   = "cyclone--cntest"
	configSpace = "cyclone-cntest"
	tenantName  = "cntest"
	pvcName     = "cyclone-pvc-" + tenantName
	pvcQuota    = "20Gi"
	podName     = "cd"

	gitlabProjectName  = "gitlab-auto"
	gitlabWorkflowName = "gitlab-auto-pipeline-20201026103357"
	gitlabServer       = "http://192.168.130.29:81"
	gitlabUrl          = "http://192.168.130.29:81/fufu/cntest20200103.git"
	gitlabRun          = "gitlab-auto-pipeline-20201026103357-v9rb9"

	githubProjectName  = "github-auto"
	githubWorkflowName = "github-auto-pipeline-20201026163420"
	githubServer       = "https://github.com"
	githubUrl          = "https://github.com/chenne69163/cppdemo.git"
	githubRun          = "github-auto-pipeline-20201026163420-bmwh5"

	bitbucketProjectName  = "bitbucket-auto"
	bitbucketWorkflowName = "bitbucket-auto-pipeline-20201027164102"
	bitbucketServer       = "http://192.168.21.106:7979"
	bitbucketUrl          = "http://admin@192.168.21.106:7979/scm/cntes/cntest.git"
	bitbucketRun          = "bitbucket-auto-pipeline-20201027164102-7f7g6"

	svnProjectName  = "svn-auto"
	svnWorkflowName = "svn-auto-pipeline-20201027164603"
	svnServer       = "svn://192.168.131.18:3691/svn"
	svnUrl          = "svn://192.168.131.18:3691/svn/test/cppdemo"
	svnRun          = "svn-auto-pipeline-20201027164603-gs95s"
)
