package http

import (
	"fmt"
	"net/http"
	"os"

	"github.com/argoproj/argo-workflows/v3/errors"
	wfv1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"github.com/argoproj/argo-workflows/v3/workflow/artifacts/common"
)

// ArtifactDriver is the artifact driver for a HTTP URL
type ArtifactDriver struct{}

var _ common.ArtifactDriver = &ArtifactDriver{}

// Load download artifacts from an HTTP URL
func (h *ArtifactDriver) Load(inputArtifact *wfv1.Artifact, path string) error {
	// Download the file to a local file path
	req, err := http.NewRequest("GET", inputArtifact.HTTP.URL, nil)
	if err != nil {
		return err
	}

	headers := inputArtifact.HTTP.Headers
	for _, v := range headers {
		req.Header.Set(v.Name, v.Value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	defer resp.Body.Close()

	err = os.MkdirAll(path, os.ModePerm) // Create the directory
	if err != nil {
		return err
	}

	return nil
}

func (h *ArtifactDriver) Save(string, *wfv1.Artifact) error {
	return errors.Errorf(errors.CodeBadRequest, "HTTP output artifacts unsupported")
}

func (h *ArtifactDriver) ListObjects(artifact *wfv1.Artifact) ([]string, error) {
	return nil, fmt.Errorf("ListObjects is currently not supported for this artifact type, but it will be in a future version")
}
