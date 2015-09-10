package repository_fetcher

import (
	"net/url"
	"strings"

	"github.com/docker/docker/registry"
	"github.com/pivotal-golang/lager"
)

//go:generate counterfeiter -o fake_pinger/fake_pinger.go . Pinger
type Pinger interface {
	Ping(*registry.Endpoint) (registry.RegistryInfo, error)
}

type EndpointPinger struct{}

func (EndpointPinger) Ping(e *registry.Endpoint) (registry.RegistryInfo, error) {
	return e.Ping()
}

type RemoteFetchRequestCreator struct {
	RegistryProvider RegistryProvider
	Pinger           Pinger
}

func (creator *RemoteFetchRequestCreator) CreateFetchRequest(logger lager.Logger, repoURL *url.URL, tag string, diskQuota int64) (*FetchRequest, error) {
	fLog := logger.Session("fetch", lager.Data{
		"repo": repoURL,
		"tag":  tag,
	})

	fLog.Debug("fetching")

	if len(repoURL.Path) == 0 {
		return nil, ErrInvalidDockerURL
	}

	path := repoURL.Path[1:]
	remotePath := path

	r, endpoint, err := creator.RegistryProvider.ProvideRegistry(repoURL.Host)
	if err != nil {
		logger.Error("failed-to-construct-registry-endpoint", err)
		return nil, FetchError("RemoteFetchRequestCreator", repoURL.Host, path, err)
	}

	if regInfo, err := creator.Pinger.Ping(endpoint); err == nil {
		logger.Debug("pinged-registry", lager.Data{
			"info":             regInfo,
			"endpoint-version": endpoint.Version,
		})
		if !regInfo.Standalone && strings.IndexRune(remotePath, '/') == -1 {
			remotePath = "library/" + remotePath
		}
	} else {
		return nil, FetchError("RemoteFetchRequestCreator", repoURL.Host, path, err)
	}

	return &FetchRequest{
		Session:    r,
		Endpoint:   endpoint,
		Logger:     fLog,
		Path:       path,
		RemotePath: remotePath,
		Tag:        tag,
		MaxSize:    diskQuota,
	}, nil
}