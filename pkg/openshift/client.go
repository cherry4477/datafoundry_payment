package openshift

import (
	"net/http"

	projectapi "github.com/openshift/origin/pkg/project/api/v1"
	"github.com/zonesan/clog"
)

type OClient struct {
	client *OpenshiftREST
	user   string
}

func NewOClient(host, token, username string) *OClient {

	clog.Debugf("%v:(%v)@%v", username, token, host)

	client := NewOpenshiftREST(NewOpenshiftTokenClient(host, token))
	return &OClient{client: client, user: username}
}

func (oc *OClient) CreateProject(r *http.Request, name string) (*projectapi.ProjectRequest, error) {

	uri := "/projectrequests"

	proj := new(projectapi.ProjectRequest)
	{
		proj.DisplayName = name
		proj.Name = oc.user + "-org-" + genRandomName(8)
		proj.Annotations = make(map[string]string)
		proj.Annotations["datafoundry.io/requester"] = oc.user
	}

	oc.client.OPost(uri, proj, proj)
	if oc.client.Err != nil {
		clog.Error(oc.client.Err)
		return nil, oc.client.Err
	}

	return proj, nil
}
