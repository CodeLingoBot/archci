package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"

	"fmt"
)

const (
	STATUS_NOT_TEST = 0
	STATUS_TESTING  = 1
	STATUS_SUCESS   = 2
	STATUS_FAIL     = 3
)

// More setting in http://beego.me/docs/mvc/model/models.md
type Project struct {
	Id          int64  `orm:"auto"`
	ProjectName string `orm:"unique"` // TODO(tobe): add size(1024)
	RepoUrl     string `orm:"size(1024)"`
	Status      int
}

type Build struct {
	Id          int64  `orm:"pk;auto"`
	ProjectName string `orm:"size(1024)"`
	Branch      string `orm:"size(1024)"`
	Commit      string `orm:"size(1024"`
	CommitTime  time.Time
	Committer   string `orm:"size(1024)"`
	Status      int
}

func RegisterModels() {
	orm.RegisterModel(new(Project), new(Build))
}

// For advanced usage in http://beego.me/docs/mvc/model/query.md#all
func GetAllProjects() []*Project {
	o := orm.NewOrm()

	var projects []*Project
	num, err := o.QueryTable("project").All(&projects)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)
	return projects
}

func GetAllBuilds() []*Build {
	o := orm.NewOrm()

	var builds []*Build
	// o.QueryTable("build").Filter("name", "slene").All(&builds) to filter with build status
	num, err := o.QueryTable("build").All(&builds)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)
	return builds
}

// For more usage in http://beego.me/docs/mvc/model/overview.md
func AddProject(projectName string, repoUrl string) error {
	o := orm.NewOrm()
	project := Project{ProjectName: projectName, RepoUrl: repoUrl}
	_, err := o.Insert(&project)
	return err
}

func AddBuild(projectName string, branch string, commit string, commitTime time.Time, committer string) error {
	o := orm.NewOrm()

	build := Build{ProjectName: projectName, Branch: branch, Commit: commit, CommitTime: commitTime, Committer: committer}
	_, err := o.Insert(&build)
	return err
}

func TestOrm() {
	o := orm.NewOrm()

	project := Project{ProjectName: "tobegit3hub/seagull", RepoUrl: "https://github.com/tobegit3hub/seagull", Status: 0}

	// insert
	id, err := o.Insert(&project)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	project.ProjectName = "ArchCI/archci"
	num, err := o.Update(&project)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read one
	u := Project{Id: project.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// delete
	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
