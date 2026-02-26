package releaser

import (
	"fmt"
	"path/filepath"

	"github.com/goodylabs/releaser/adapters/httpconnector"
	"github.com/goodylabs/releaser/adapters/oshelper"
	"github.com/goodylabs/releaser/ports"
)

type GithubOpts struct {
	User string
	Repo string
}

var releaseRes struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		Name               string `json:"name"`
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

type githubApp struct {
	opts           GithubOpts
	newReleaseName string
	newReleaseUrl  string
	httpconnector  *httpconnector.HttpClient
	oshelper       *oshelper.OsHelper
}

func NewGithubApp(opts *GithubOpts) ports.Provider {
	return &githubApp{
		opts:          *opts,
		httpconnector: httpconnector.NewHttpClient(),
	}
}

func (g *githubApp) GetNewestReleaseName() (string, error) {
	lastestReleaseUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", g.opts.User, g.opts.Repo)
	if err := g.httpconnector.DoGet(lastestReleaseUrl, &releaseRes); err != nil {
		return "", err
	}

	g.newReleaseName = releaseRes.TagName

	osType := g.oshelper.GetOSType()
	osArch, err := g.oshelper.GetArch()
	if err != nil {
		return "", err
	}

	assetName := fmt.Sprintf("%s-%s", osType, osArch)
	for _, asset := range releaseRes.Assets {
		if asset.Name == assetName {
			g.newReleaseUrl = asset.BrowserDownloadURL
			break
		}
	}
	if g.newReleaseUrl == "" {
		return "", fmt.Errorf("no compatible binary found for %s-%s while looking for asset %s", osType, osArch, assetName)
	}

	return g.newReleaseName, nil
}

func (g *githubApp) PerformUpdate(appDir string) error {
	osType := g.oshelper.GetOSType()
	osArch, err := g.oshelper.GetArch()
	if err != nil {
		return err
	}

	if g.newReleaseUrl == "" {
		return fmt.Errorf("no compatible binary found for %s-%s", osType, osArch)
	}

	fmt.Println("Downloading binary from:", g.newReleaseUrl)

	binnaryDir := filepath.Join(appDir, "bin")
	if err := g.oshelper.MakeDirIfNotExist(binnaryDir); err != nil {
		return err
	}

	binnaryPath := filepath.Join(binnaryDir, g.opts.Repo)
	if err := g.oshelper.DownloadBinary(g.newReleaseUrl, binnaryPath); err != nil {
		return err
	}

	fmt.Println("Updated binary at:", binnaryPath)

	return nil
}
