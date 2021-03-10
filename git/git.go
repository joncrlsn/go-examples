package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"

	"github.com/Masterminds/semver"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

var semverRegex = regexp.MustCompile(`.*/([a-z]?)((\d+)\.(\d+)\.(\d+)).*`)

// Git new
type Git struct {
	repository git.Repository
}

func main() {
	g, err := NewAt("./test")
	if err != nil {
		panic(err)
	}

	tags, err := g.Tags()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", tags)
}

func New() (*Git, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return NewAt(dir)
}

func NewAt(dirPath string) (*Git, error) {

	opt := &git.PlainOpenOptions{DetectDotGit: true}
	repo, err := git.PlainOpenWithOptions(dirPath, opt)
	if err != nil {
		return nil, err
	}

	git := &Git{
		repository: *repo,
	}
	return git, nil
}

func IsGit() bool {
	_, err := New()
	return err == nil
}

func (g *Git) Revision(longSha bool) (string, error) {
	h, err := g.CurrentCommit()
	if longSha || err != nil {
		return h, err
	}
	return h[:7], err
}

func (g *Git) IsDirty() bool {
	w, err := g.repository.Worktree()
	if err != nil {
		return true
	}

	status, err := w.Status()
	if err != nil {
		return true
	}

	return !status.IsClean()

	// res, _ := oneliner("git", "status", "--porcelain")
	// return len(res) > 0
}

func (g *Git) Branches() ([]string, error) {
	var currentBranchesNames []string

	branchRefs, err := g.repository.Branches()
	if err != nil {
		return currentBranchesNames, err
	}

	headRef, err := g.repository.Head()
	if err != nil {
		return currentBranchesNames, err
	}

	err = branchRefs.ForEach(func(branchRef *plumbing.Reference) error {
		if branchRef.Hash() == headRef.Hash() {
			currentBranchesNames = append(currentBranchesNames, branchRef.Name().Short())

			return nil
		}

		return nil
	})
	if err != nil {
		return currentBranchesNames, err
	}

	return currentBranchesNames, nil
}

func (g *Git) CurrentCommit() (string, error) {
	headRef, err := g.repository.Head()
	if err != nil {
		return "", err
	}
	headSha := headRef.Hash().String()

	return headSha, nil
}

func (g *Git) Tags() ([]string, error) {
	tags, _, err := g.tags()
	return tags, err
}

func (g *Git) LatestTag() (string, error) {
	_, tag, err := g.tags()
	return tag, err
}

// tags returns a list of tags sorted with the largest first
func (g *Git) tags() ([]string, string, error) {
	var latestTagName string
	var tags []string
	var semvers semver.Collection

	tagRefs, err := g.repository.Tags()
	if err != nil {
		return tags, latestTagName, err
	}

	var prefix string
	err = tagRefs.ForEach(func(tagRef *plumbing.Reference) error {
		tag := tagRef.Name().String()
		if semverRegex.MatchString(tag) {
			if prefix == "" {
				prefix = semverRegex.ReplaceAllString(tag, `$1`)
			}
			semverPart := semverRegex.ReplaceAllString(tag, `$2`)
			//fmt.Println(semverPart)
			version, err := semver.NewVersion(semverPart)
			if err != nil {
				fmt.Println("Invalid semver tag:", tag)
				return err
			}
			//fmt.Println("Adding version:", version)
			semvers = append(semvers, version)
		}

		return nil
	})

	if err != nil {
		return tags, latestTagName, err
	}

	sort.Sort(sort.Reverse(semvers))

	// Map the Version objects to strings
	tags = make([]string, len(semvers))
	for i, v := range semvers {
		if i == 0 {
			latestTagName = prefix + v.String()
		}
		tags[i] = prefix + v.String()
	}

	//fmt.Println(tags, latestTagName)
	return tags, latestTagName, nil
}

func (g *Git) CreateTag(version string) error {
	head, err := g.repository.Head()
	if err != nil {
		return err
	}

	_, err = g.repository.CreateTag(version, head.Hash(), nil)
	return err

}
