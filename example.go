package main

import (
    "encoding/json"
    "log"
    "os"
    "bytes"
    "fmt"
    "time"
    "io/ioutil"
)

type SecretJsonbody struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Labels struct {
			Heritage  string `json:"heritage"`
			Project   string `json:"project"`
			Build     string `json:"build"`
			Component string `json:"component"`
		} `json:"labels"`
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
	} `json:"metadata"`
	Type string `json:"type"`
	Data struct {
		Commit_id      string `json:"commit"`
		Commit_ref     string `json:"ref"`
		Event_provider string `json:"event_provider"`
		Event_type     string `json:"event_type"`
		Project_id     string `json:"project_id"`
		Build_id       string `json:"build_id"`
		Payload        string `json:"payload"`
	} `json:"data"`
}

type GiteaJsonbody struct {
	Secret     string `json:"secret,omitempty"`
	Ref        string `json:"ref,omitempty,omitempty"`
	Before     string `json:"before,omitempty,omitempty"`
	After      string `json:"after,omitempty,omitempty"`
	CompareURL string `json:"compare_url,omitempty,omitempty"`
	Commits    *[]struct {
		ID      string `json:"id,omitempty"`
		Message string `json:"message,omitempty"`
		URL     string `json:"url,omitempty"`
		Author  *struct {
			Name     string `json:"name,omitempty"`
			Email    string `json:"email,omitempty"`
			Username string `json:"username,omitempty"`
		} `json:"author,omitempty"`
		Committer *struct {
			Name     string `json:"name,omitempty"`
			Email    string `json:"email,omitempty"`
			Username string `json:"username,omitempty"`
		} `json:"committer,omitempty"`
		Verification interface{} `json:"verification,omitempty"`
		Timestamp    *time.Time   `json:"timestamp,omitempty"`
	} `json:"commits,omitempty"`
	Repository *struct {
		ID    int `json:"id,omitempty"`
		Owner *struct {
			ID        int    `json:"id,omitempty"`
			Login     string `json:"login,omitempty"`
			FullName  string `json:"full_name,omitempty"`
			Email     string `json:"email,omitempty"`
			AvatarURL string `json:"avatar_url,omitempty"`
			Language  string `json:"language,omitempty"`
			Username  string `json:"username,omitempty"`
		} `json:"owner,omitempty"`
		Name            string      `json:"name,omitempty"`
		FullName        string      `json:"full_name,omitempty"`
		Description     string      `json:"description,omitempty"`
		Empty           bool        `json:"empty,omitempty"`
		Private         bool        `json:"private,omitempty"`
		Fork            bool        `json:"fork,omitempty"`
		Parent          interface{} `json:"parent,omitempty"`
		Mirror          bool        `json:"mirror,omitempty"`
		Size            int         `json:"size,omitempty"`
		HTMLURL         string      `json:"html_url,omitempty"`
		SSHURL          string      `json:"ssh_url,omitempty"`
		CloneURL        string      `json:"clone_url,omitempty"`
		Website         string      `json:"website,omitempty"`
		StarsCount      int         `json:"stars_count,omitempty"`
		ForksCount      int         `json:"forks_count,omitempty"`
		WatchersCount   int         `json:"watchers_count,omitempty"`
		OpenIssuesCount int         `json:"open_issues_count,omitempty"`
		DefaultBranch   string      `json:"default_branch,omitempty"`
		CreatedAt       *time.Time   `json:"created_at,omitempty"`
		UpdatedAt       *time.Time   `json:"updated_at,omitempty"`
		Permissions     *struct {
			Admin bool `json:"admin,omitempty"`
			Push  bool `json:"push,omitempty"`
			Pull  bool `json:"pull,omitempty"`
		} `json:"permissions,omitempty"`
	} `json:"repository,omitempty"`
	Pusher *struct {
		ID        int    `json:"id,omitempty"`
		Login     string `json:"login,omitempty"`
		FullName  string `json:"full_name,omitempty"`
		Email     string `json:"email,omitempty"`
		AvatarURL string `json:"avatar_url,omitempty"`
		Language  string `json:"language,omitempty"`
		Username  string `json:"username,omitempty"`
	} `json:"pusher,omitempty"`
	Action      string `json:"action,omitempty"`
	Number      int    `json:"number,omitempty"`
	PullRequest *struct {
		ID     int    `json:"id,omitempty"`
		URL    string `json:"url,omitempty"`
		Number int    `json:"number,omitempty"`
		User   *struct {
			ID        int    `json:"id,omitempty"`
			Login     string `json:"login,omitempty"`
			FullName  string `json:"full_name,omitempty"`
			Email     string `json:"email,omitempty"`
			AvatarURL string `json:"avatar_url,omitempty"`
			Language  string `json:"language,omitempty"`
			Username  string `json:"username,omitempty"`
		} `json:"user,omitempty"`
		Title          string        `json:"title,omitempty"`
		Body           string        `json:"body,omitempty"`
		Labels         []interface{} `json:"labels,omitempty"`
		Milestone      interface{}   `json:"milestone,omitempty"`
		Assignee       interface{}   `json:"assignee,omitempty"`
		Assignees      interface{}   `json:"assignees,omitempty"`
		State          string        `json:"state,omitempty"`
		Comments       int           `json:"comments,omitempty"`
		HTMLURL        string        `json:"html_url,omitempty"`
		DiffURL        string        `json:"diff_url,omitempty"`
		PatchURL       string        `json:"patch_url,omitempty"`
		Mergeable      bool          `json:"mergeable,omitempty"`
		Merged         bool          `json:"merged,omitempty"`
		MergedAt       *time.Time     `json:"merged_at,omitempty"`
		MergeCommitSha string        `json:"merge_commit_sha,omitempty"`
		MergedBy       *struct {
			ID        int    `json:"id,omitempty"`
			Login     string `json:"login,omitempty"`
			FullName  string `json:"full_name,omitempty"`
			Email     string `json:"email,omitempty"`
			AvatarURL string `json:"avatar_url,omitempty"`
			Language  string `json:"language,omitempty"`
			Username  string `json:"username,omitempty"`
		} `json:"merged_by,omitempty"`
		Base *struct {
			Label  string `json:"label,omitempty"`
			Ref    string `json:"ref,omitempty"`
			Sha    string `json:"sha,omitempty"`
			RepoID int    `json:"repo_id,omitempty"`
			Repo   *struct {
				ID    int `json:"id,omitempty"`
				Owner *struct {
					ID        int    `json:"id,omitempty"`
					Login     string `json:"login,omitempty"`
					FullName  string `json:"full_name,omitempty"`
					Email     string `json:"email,omitempty"`
					AvatarURL string `json:"avatar_url,omitempty"`
					Language  string `json:"language,omitempty"`
					Username  string `json:"username,omitempty"`
				} `json:"owner,omitempty"`
				Name            string      `json:"name,omitempty"`
				FullName        string      `json:"full_name,omitempty"`
				Description     string      `json:"description,omitempty"`
				Empty           bool        `json:"empty,omitempty"`
				Private         bool        `json:"private,omitempty"`
				Fork            bool        `json:"fork,omitempty"`
				Parent          interface{} `json:"parent,omitempty"`
				Mirror          bool        `json:"mirror,omitempty"`
				Size            int         `json:"size,omitempty"`
				HTMLURL         string      `json:"html_url,omitempty"`
				SSHURL          string      `json:"ssh_url,omitempty"`
				CloneURL        string      `json:"clone_url,omitempty"`
				Website         string      `json:"website,omitempty"`
				StarsCount      int         `json:"stars_count,omitempty"`
				ForksCount      int         `json:"forks_count,omitempty"`
				WatchersCount   int         `json:"watchers_count,omitempty"`
				OpenIssuesCount int         `json:"open_issues_count,omitempty"`
				DefaultBranch   string      `json:"default_branch,omitempty"`
				CreatedAt       *time.Time   `json:"created_at,omitempty"`
				UpdatedAt       *time.Time   `json:"updated_at,omitempty"`
				Permissions     *struct {
					Admin bool `json:"admin,omitempty"`
					Push  bool `json:"push,omitempty"`
					Pull  bool `json:"pull,omitempty"`
				} `json:"permissions,omitempty"`
			} `json:"repo,omitempty"`
		} `json:"base,omitempty"`
		Head *struct {
			Label  string `json:"label,omitempty"`
			Ref    string `json:"ref,omitempty"`
			Sha    string `json:"sha,omitempty"`
			RepoID int    `json:"repo_id,omitempty"`
			Repo   *struct {
				ID    int `json:"id,omitempty"`
				Owner *struct {
					ID        int    `json:"id,omitempty"`
					Login     string `json:"login,omitempty"`
					FullName  string `json:"full_name,omitempty"`
					Email     string `json:"email,omitempty"`
					AvatarURL string `json:"avatar_url,omitempty"`
					Language  string `json:"language,omitempty"`
					Username  string `json:"username,omitempty"`
				} `json:"owner,omitempty"`
				Name            string      `json:"name,omitempty"`
				FullName        string      `json:"full_name,omitempty"`
				Description     string      `json:"description,omitempty"`
				Empty           bool        `json:"empty,omitempty"`
				Private         bool        `json:"private,omitempty"`
				Fork            bool        `json:"fork,omitempty"`
				Parent          interface{} `json:"parent,omitempty"`
				Mirror          bool        `json:"mirror,omitempty"`
				Size            int         `json:"size,omitempty"`
				HTMLURL         string      `json:"html_url,omitempty"`
				SSHURL          string      `json:"ssh_url,omitempty"`
				CloneURL        string      `json:"clone_url,omitempty"`
				Website         string      `json:"website,omitempty"`
				StarsCount      int         `json:"stars_count,omitempty"`
				ForksCount      int         `json:"forks_count,omitempty"`
				WatchersCount   int         `json:"watchers_count,omitempty"`
				OpenIssuesCount int         `json:"open_issues_count,omitempty"`
				DefaultBranch   string      `json:"default_branch,omitempty"`
				CreatedAt       *time.Time   `json:"created_at,omitempty"`
				UpdatedAt       *time.Time   `json:"updated_at,omitempty"`
				Permissions     *struct {
					Admin bool `json:"admin,omitempty"`
					Push  bool `json:"push,omitempty"`
					Pull  bool `json:"pull,omitempty"`
				} `json:"permissions,omitempty"`
			} `json:"repo,omitempty"`
		} `json:"head,omitempty"`
		MergeBase string      `json:"merge_base,omitempty"`
		DueDate   interface{} `json:"due_date,omitempty"`
		CreatedAt *time.Time   `json:"created_at,omitempty"`
		UpdatedAt *time.Time   `json:"updated_at,omitempty"`
		ClosedAt  interface{} `json:"closed_at,omitempty"`
	} `json:"pull_request,omitempty"`
	Sender *struct {
		ID        int    `json:"id,omitempty"`
		Login     string `json:"login,omitempty"`
		FullName  string `json:"full_name,omitempty"`
		Email     string `json:"email,omitempty"`
		AvatarURL string `json:"avatar_url,omitempty"`
		Language  string `json:"language,omitempty"`
		Username  string `json:"username,omitempty"`
	} `json:"sender,omitempty"`
}

func main() {
    //m := []byte(`{"type":"example","data": {"name": "abc","labels": {"key": "value"}},"subsets": [{"addresses": [{"ip": "192.168.103.178"}],"ports": [{"port": 80}]}]}`)
    reqbody := []byte(`{"commits": [{"author": {"name":"Abulfat"}}],"pull_request": {"user": {"login":"amasimov"},"base": {"repo": {"owner": {"login":"abulfat"}}}},"secret":"example","repository": {"full_name":"devops/hello-world"},"action":"open","ref":"develop","After": "262ff7f24c9d129e39e6ad0974c66006d98d64a5","sender": {"login":"Neigh","email":"a@b.c"}}`)

    //val := &SecretJsonbody{}
    giteajsonbody := &GiteaJsonbody{}

    r := bytes.NewReader(reqbody)
    err := json.NewDecoder(r).Decode(&giteajsonbody)
    if err != nil {
        log.Fatal(err)
    }

//eventtype := r.Header.Get("X-Gitea-Event")

//if (eventtype == "push") {
//        fmt.Println("Push request: " + giteajsonbody.after)
//    } else if (eventtype == "pull_request") && (giteajsonbody.action == "opened" ) {
//        fmt.Println("PR opened: " + giteajsonbody.pull_request.id)
//    } else if (eventtype == "pull_request") && (giteajsonbody.action == "
//        fmt.Println("PR merged: " + giteajsonbody.pull_request.id + giteajsonbody.pull_request.merge_commit_sha)
//    } else {
//        fmt.Println("Unknown event type" + eventtype)
//}

    z := new(bytes.Buffer)
//    JSON.Indent(&out, b, "", "    ")
    err1 := json.NewEncoder(z).Encode(&giteajsonbody)
    if err1 != nil {
        log.Fatal(err1)
    }

    var out bytes.Buffer
    json.Indent(&out, z.Bytes(), "", "    ")
    out.WriteTo(os.Stdout)
    var out1 bytes.Buffer
    json.Indent(&out1, z.Bytes(), "", "    ")
    out1.WriteTo(os.Stdout)
    body, _ := ioutil.ReadAll(z)
    os.Stdout.Write(body)
    fmt.Println("request Body:", string(body))

}
