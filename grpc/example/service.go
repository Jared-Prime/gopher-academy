package example

import (
	"context"
	"fmt"
	"net/http"
)

type releaseInfo struct {
	ReleaseDate     string `json:"release_date"`
	ReleaseNotesURL string `json:"release_notes_url"`
}

type ExampleService interface {
	GetReleaseInfo(context.Context, *GetReleaseInfoRequest) (*ReleaseInfo, error)
	ListReleases(context.Context, *ListReleasesRequest) *ListReleasesResponse
}

type goReleaseService struct {
	releases map[string]releaseInfo
}

func (service *goReleaseService) GetReleaseInfo(ctx context.Context, request *GetReleaseInfoRequest) (*ReleaseInfo, error) {
	release_info, ok := service.releases[request.GetVersion()]
	if !ok {
		return nil, fmt.Errorf(http.StatusText(http.StatusNotFound), "release versions %s not found", request.GetVersion())
	}

	return &ReleaseInfo{
		Version:         request.GetVersion(),
		ReleaseDate:     release_info.ReleaseDate,
		ReleaseNotesUrl: release_info.ReleaseNotesURL,
	}, nil
}

func (service *goReleaseService) ListReleases(ctx context.Context, request *ListReleasesRequest) (*ListReleasesResponse, error) {
	var releases []*ReleaseInfo

	for version, release := range service.releases {
		release_info := &ReleaseInfo{
			Version:         version,
			ReleaseDate:     release.ReleaseDate,
			ReleaseNotesUrl: release.ReleaseNotesURL,
		}

		releases = append(releases, release_info)
	}

	return &ListReleasesResponse{
		Releases: releases,
	}, nil
}
