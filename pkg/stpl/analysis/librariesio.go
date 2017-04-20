package analysis

type LibrariesIoComponent struct {
	Name                     string   `json:"name"`
	Platform                 string   `json:"platform"`
	Description              string   `json:"description"`
	Homepage                 string   `json:"homepage"`
	RepositoryURL            string   `json:"repository_url"`
	Rank                     int      `json:"rank"`
	LatestReleasePublishedAt string   `json:"latest_release_published_at"`
	LatestReleaseNumber      string   `json:"latest_release_number"`
	Language                 string   `json:"language"`
	Status                   string   `json:"status"`
	PackageManagerURL        string   `json:"package_manager_url"`
	Stars                    int      `json:"stars"`
	Forks                    int      `json:"forks"`
	Keywords                 []string `json:"keywords"`
	// LatestStableRelease object   `json:"latest_stable_release"`
}
