package handler

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/google/uuid"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/store"

	pb "github.com/micro/services/project/service/proto"
)

// Project implements the project service interface
type Project struct {
	name  string
	store store.Store
}

// New returns an initialized project handler
func New(service micro.Service) *Project {
	return &Project{
		name:  service.Name(),
		store: store.DefaultStore,
	}
}

const (
	// projectsPrefix is the store prefix for projects. projects are stored with
	// keys in the following format "projects/{namespace}/{id}". This allows
	// us to lookup projects using both namespace and ID. Namespace is used
	// more commonly than ID so we'll use this as the first component of
	// the key.
	projectsPrefix = "projects/"
	// membersPrefix is the stroe prefix for memberships. Memberships are
	// stored with key in the following format "memberships/{projectID}/{userID}".
	// The value is the user ID (string, stored as bytes).
	membersPrefix = "members/"
)

var (
	// reservedNamespaces cannot be used by projects
	reservedNamespaces = []string{"default", "go.micro", "runtime"}
)

// Read looks up a project using ID or namespace
func (p *Project) Read(ctx context.Context, req *pb.ReadRequest, rsp *pb.ReadResponse) error {
	// lookup the project
	var err error
	if len(req.Id) > 0 {
		rsp.Project, err = p.findProjectByID(req.Id)
	}
	if len(req.Namespace) > 0 && rsp.Project == nil {
		rsp.Project, err = p.findProjectByNamespace(req.Namespace)
	}
	if err != nil {
		return err
	}

	// lookup the project members
	recs, err := p.store.Read(membersPrefix+rsp.Project.Id+"/", store.ReadPrefix())
	if err != nil {
		return nil
	}
	rsp.Project.Members = make([]*pb.Member, 0, len(recs))
	for _, r := range recs {
		rsp.Project.Members = append(rsp.Project.Members, &pb.Member{
			Id: string(r.Value),
		})
	}

	return nil
}

// Create a new projects
func (p *Project) Create(ctx context.Context, req *pb.CreateRequest, rsp *pb.CreateResponse) error {
	// validate the request
	if req.Project == nil {
		return errors.BadRequest(p.name, "Missing project")
	}
	if len(req.Project.Name) == 0 {
		return errors.BadRequest(p.name, "Missing project name")
	}
	if err := p.validateNamespace(req.Project.Namespace); err != nil {
		return err
	}

	// add the default fields
	req.Project.Id = uuid.New().String()

	// write to the store
	if err := p.writeProjectToStore(req.Project); err != nil {
		return err
	}

	// return the project in the response
	rsp.Project = req.Project
	return nil
}

// Update a project
func (p *Project) Update(ctx context.Context, req *pb.UpdateRequest, rsp *pb.UpdateResponse) error {
	// validate the request
	if req.Project == nil {
		return errors.BadRequest(p.name, "Missing project")
	}
	if len(req.Project.Id) == 0 {
		return errors.BadRequest(p.name, "Missing project id")
	}
	if len(req.Project.Name) == 0 {
		return errors.BadRequest(p.name, "Missing project name")
	}

	// lookup the project
	project, err := p.findProjectByID(req.Project.Id)
	if err != nil {
		return errors.BadRequest(p.name, "Error finding project: %v", err)
	}

	// assign the update params
	project.Name = req.Project.Name
	project.WebDomain = req.Project.WebDomain
	project.ApiDomain = req.Project.ApiDomain

	// write to the store
	if err := p.writeProjectToStore(req.Project); err != nil {
		return errors.InternalServerError(p.name, "Error writing project: %v", err)
	}

	return nil
}

// List all the projects (does not return membership)
func (p *Project) List(ctx context.Context, req *pb.ListRequest, rsp *pb.ListResponse) error {
	// get the records with the project prefix
	recs, err := p.store.Read(projectsPrefix, store.ReadPrefix())
	if err != nil {
		return err
	}

	// unmarshal and return in the response
	rsp.Projects = make([]*pb.Project, len(recs))
	for i, r := range recs {
		if err := json.Unmarshal(r.Value, &rsp.Projects[i]); err != nil {
			return errors.InternalServerError(p.name, "Error unmarsaling json: %v", err)
		}
	}

	return nil
}

// AddMember to a project
func (p *Project) AddMember(ctx context.Context, req *pb.AddMemberRequest, rsp *pb.AddMemberResponse) error {
	// validate the request
	if _, err := p.findProjectByID(req.ProjectId); err != nil {
		return err
	}
	if len(req.MemberId) == 0 {
		return errors.BadRequest(p.name, "Missing member id")
	}

	// write the membership to the store
	return p.store.Write(&store.Record{
		Key:   membersPrefix + req.ProjectId + "/" + req.MemberId,
		Value: []byte(req.MemberId),
	})
}

// RemoveMember from a project
func (p *Project) RemoveMember(ctx context.Context, req *pb.RemoveMemberRequest, rsp *pb.RemoveMemberResponse) error {
	return p.store.Delete(membersPrefix + req.ProjectId + "/" + req.MemberId)
}

// ListMemberships returns all the projects a member belongs to
func (p *Project) ListMemberships(ctx context.Context, req *pb.ListMembershipsRequest, rsp *pb.ListMembershipsResponse) error {
	// member id is the last component of the key, so list all
	// the keys in the store which relate to memberships
	keys, err := p.store.List(store.ListPrefix(membersPrefix))
	if err != nil {
		return err
	}

	// filter to get the project ids which the member belongs to
	var projectIDs []string
	for _, k := range keys {
		if strings.HasSuffix(k, "/"+req.MemberId) {
			projectIDs = append(projectIDs, strings.Split(k, "/")[1])
		}
	}

	// get each of the projects
	rsp.Projects = make([]*pb.Project, 0, len(projectIDs))
	for _, id := range projectIDs {
		project, err := p.findProjectByID(id)
		if err != nil {
			return err
		}
		rsp.Projects = append(rsp.Projects, project)
	}

	return nil
}

func (p *Project) findProjectByID(id string) (*pb.Project, error) {
	// ID is stored as the last component of the key, so list
	// all the keys in the store which relate to projects
	keys, err := p.store.List(store.ListPrefix(projectsPrefix))
	if err != nil {
		return nil, err
	}

	// Check each key to see if it ends in the ID. If the key
	// is not found, return an error.
	var projectKey string
	for _, k := range keys {
		if strings.HasSuffix(k, "/"+id) {
			projectKey = k
			break
		}
	}
	if len(projectKey) == 0 {
		return nil, store.ErrNotFound
	}

	// Lookup the record and then decode the value
	recs, err := p.store.Read(projectKey)
	if err != nil {
		return nil, err
	}
	var project *pb.Project
	err = json.Unmarshal(recs[0].Value, &project)
	return project, err
}

// writeProjectToStore marshals a project and writes it to the store under
// the corresponding key (prefix + namespace + id)
func (p *Project) writeProjectToStore(project *pb.Project) error {
	bytes, err := json.Marshal(project)
	if err != nil {
		return errors.InternalServerError(p.name, "Error marsaling json: %v", err)
	}

	key := projectsPrefix + project.Namespace + "/" + project.Id
	if err := p.store.Write(&store.Record{Key: key, Value: bytes}); err != nil {
		return errors.InternalServerError(p.name, "Error writing to the store: %v", err)
	}

	return nil
}

func (p *Project) findProjectByNamespace(ns string) (*pb.Project, error) {
	// Namespace is the first component of the key, so lookup records which
	// have this as a prefix. Read does't return an error when using the
	// ReadPrefix option, so we also need to check for an empty slice.
	recs, err := p.store.Read(projectsPrefix+ns+"/", store.ReadPrefix())
	if err != nil {
		return nil, err
	} else if len(recs) != 1 {
		return nil, store.ErrNotFound
	}

	// Unmarshal and return the result
	var project *pb.Project
	err = json.Unmarshal(recs[0].Value, &project)
	return project, err
}

// validateNamespace returns an error if the namespace provided is invalid.
func (p *Project) validateNamespace(ns string) error {
	// compare namespaces in lowercase
	ns = strings.ToLower(ns)

	// validate the length of the namespace
	if len(ns) < 3 {
		return errors.BadRequest(p.name, "Namespaces must be at least 3 characters long")
	}
	if len(ns) > 20 {
		return errors.BadRequest(p.name, "Namespaces must be at no more than 20 characters long")
	}

	// check against reserved namespaces
	for _, v := range reservedNamespaces {
		if v == ns {
			return errors.BadRequest(p.name, "%v is a reserved namespace", ns)
		}
	}

	// check against existing namespaces. The namespace is used
	// as the key in the store
	recs, err := p.store.Read(projectsPrefix+ns+"/", store.ReadPrefix())
	if err != nil {
		return err
	} else if len(recs) > 0 {
		return errors.BadRequest(p.name, "The namespace %v has already been taken", ns)
	}

	return nil
}
