// Code generated by protoc-gen-gogo.
// source: graph.proto
// DO NOT EDIT!

/*
Package graph2 is a generated protocol buffer package.

It is generated from these files:
	graph.proto
	output.proto

It has these top-level messages:
	GitHubRepo
	RepoPermissions
	TreeKey
	Tree
	TreeRev
	RawDep
	UnitKey
	UnitInfo
	Unit
	NodeKey
	Node
	DefData
	DocData
	Edge
	Output
*/
package graph2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto"
import pbtypes "sourcegraph.com/sqs/pbtypes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// GitHubRepo holds additional metadata about GitHub repos.
type GitHubRepo struct {
	Stars int32 `protobuf:"varint,1,opt,name=stars,proto3" json:"stars,omitempty"`
}

func (m *GitHubRepo) Reset()         { *m = GitHubRepo{} }
func (m *GitHubRepo) String() string { return proto.CompactTextString(m) }
func (*GitHubRepo) ProtoMessage()    {}

// RepoPermissions describes the possible permissions that a user (or an anonymous
// user) can be granted to a repository.
type RepoPermissions struct {
	Read  bool `protobuf:"varint,1,opt,name=read,proto3" json:"read,omitempty"`
	Write bool `protobuf:"varint,2,opt,name=write,proto3" json:"write,omitempty"`
	Admin bool `protobuf:"varint,3,opt,name=admin,proto3" json:"admin,omitempty"`
}

func (m *RepoPermissions) Reset()         { *m = RepoPermissions{} }
func (m *RepoPermissions) String() string { return proto.CompactTextString(m) }
func (*RepoPermissions) ProtoMessage()    {}

type TreeKey struct {
	// Genus is the type of source tree
	// (git, hg, build-system specific like pip or npm, ftp)
	Genus string `protobuf:"bytes,1,opt,name=genus,proto3" json:"genus,omitempty"`
	// A URL of one of the following forms:
	//   src://{hostname}/{path} (absolute)
	//   src:///{path} (relative)
	URI string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (m *TreeKey) Reset()         { *m = TreeKey{} }
func (m *TreeKey) String() string { return proto.CompactTextString(m) }
func (*TreeKey) ProtoMessage()    {}

type Tree struct {
	TreeKey   `protobuf:"bytes,1,opt,name=key,embedded=key" json:""`
	FetchUrls []string `protobuf:"bytes,2,rep,name=fetch_urls" json:"fetch_urls,omitempty"`
	// Origin is populated for repos fetched via federation or
	// discovery. It is the hostname of the host that owns the repo.
	Origin string `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	// Name is the base name (the final path component) of the source tree,
	// typically the name of the directory that the repository would be cloned
	// into. (For example, for git://example.com/foo.git, the name is "foo".)
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description is a brief description of the repository.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// HomepageURL is the URL to the repository's homepage, if any.
	HomepageURL string `protobuf:"bytes,6,opt,name=homepage_url,proto3" json:"homepage_url,omitempty"`
	// HTMLURL is the URL to the repository's main page on the
	// Sourcegraph server.
	HTMLURL string `protobuf:"bytes,7,opt,name=html_url,proto3" json:"html_url,omitempty"`
	// DefaultBranch is the default VCS branch used (typically "master" for git
	// repositories and "default" for hg repositories).
	DefaultBranch string `protobuf:"bytes,8,opt,name=default_branch,proto3" json:"default_branch,omitempty"`
	// Language is the primary programming language used in this source tree.
	Language string `protobuf:"bytes,9,opt,name=language,proto3" json:"language,omitempty"`
	// Blocked is whether this repo has been blocked by an admin (and
	// will not be returned via the external API).
	Blocked bool `protobuf:"varint,10,opt,name=blocked,proto3" json:"blocked,omitempty"`
	// Deprecated source trees are labeled as such and hidden from global search
	// results.
	Deprecated bool `protobuf:"varint,11,opt,name=deprecated,proto3" json:"deprecated,omitempty"`
	// Fork is whether this source tree is a VCS repository fork.
	Fork bool `protobuf:"varint,12,opt,name=fork,proto3" json:"fork,omitempty"`
	// Mirror indicates whether this source tree's canonical location is on
	// another server. Mirror source trees track their upstream.
	Mirror bool `protobuf:"varint,13,opt,name=mirror,proto3" json:"mirror,omitempty"`
	// Private is whether this source tree is private.
	Private bool `protobuf:"varint,14,opt,name=private,proto3" json:"private,omitempty"`
	// CreatedAt is when this source tree was created. If it represents an
	// externally hosted (e.g., GitHub) source tree, the creation date is when it
	// was created at that origin.
	CreatedAt *pbtypes.Timestamp `protobuf:"bytes,15,opt,name=created_at" json:"created_at,omitempty"`
	// UpdatedAt is when this source tree's metadata was last updated (on its origin if
	// it's an externally hosted source tree).
	UpdatedAt *pbtypes.Timestamp `protobuf:"bytes,16,opt,name=updated_at" json:"updated_at,omitempty"`
	// PushedAt is when this source tree's source code was last updated (e.g., via
	// VCS-push).
	PushedAt *pbtypes.Timestamp `protobuf:"bytes,17,opt,name=pushed_at" json:"pushed_at,omitempty"`
	// Permissions describes the permissions that the current user (or anonymous users,
	// if there is no current user) is granted to this repository.
	Permissions *RepoPermissions `protobuf:"bytes,18,opt,name=permissions" json:"permissions,omitempty"`
	GitHub      *GitHubRepo      `protobuf:"bytes,19,opt,name=github" json:"github,omitempty"`
}

func (m *Tree) Reset()         { *m = Tree{} }
func (m *Tree) String() string { return proto.CompactTextString(m) }
func (*Tree) ProtoMessage()    {}

type TreeRev struct {
	Tree *Tree `protobuf:"bytes,1,opt,name=tree" json:"tree,omitempty"`
	// Rev is the version of the source tree
	Rev string `protobuf:"bytes,2,opt,name=rev,proto3" json:"rev,omitempty"`
	// Type of revision (e.g., commit ID, branch name, or tag for VCS
	// repositories; version string for build system packages)
	RevType string `protobuf:"bytes,3,opt,name=rev_type,proto3" json:"rev_type,omitempty"`
}

func (m *TreeRev) Reset()         { *m = TreeRev{} }
func (m *TreeRev) String() string { return proto.CompactTextString(m) }
func (*TreeRev) ProtoMessage()    {}

type RawDep struct {
	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *RawDep) Reset()         { *m = RawDep{} }
func (m *RawDep) String() string { return proto.CompactTextString(m) }
func (*RawDep) ProtoMessage()    {}

type UnitKey struct {
	TreeKey `protobuf:"bytes,1,opt,name=key,embedded=key" json:""`
	// Version is the build-system-level version string. It can also be a commit
	// ID, but no client code should depend on implied VCS semantics (e.g., never
	// run `git clone buildUnit.Version`).
	Version  string `protobuf:"bytes,2,opt,name=Version,proto3" json:"Version,omitempty"`
	UnitName string `protobuf:"bytes,3,opt,name=UnitName,proto3" json:"UnitName,omitempty"`
	UnitType string `protobuf:"bytes,4,opt,name=UnitType,proto3" json:"UnitType,omitempty"`
}

func (m *UnitKey) Reset()         { *m = UnitKey{} }
func (m *UnitKey) String() string { return proto.CompactTextString(m) }
func (*UnitKey) ProtoMessage()    {}

type UnitInfo struct {
	// NameInRepository is the name to use when displaying the source unit in
	// the context of the repository in which it is defined. This name
	// typically needs less qualification than GlobalName.
	//
	// For example, a Go package's GlobalName is its repository URI basename
	// plus its directory path within the repository (e.g.,
	// "github.com/user/repo/x/y"'s NameInRepository is "repo/x/y"). Because npm
	// and pip packages are named globally, their name is probably appropriate
	// to use as both the unit's NameInRepository and GlobalName.
	NameInRepository string `protobuf:"bytes,1,opt,name=NameInRepository,proto3" json:"NameInRepository,omitempty"`
	// GlobalName is the name to use when displaying the source unit *OUTSIDE OF*
	// the context of the repository in which it is defined.
	//
	// For example, a Go package's GlobalName is its full import path. Because
	// npm and pip packages are named globally, their name is probably
	// appropriate to use as both the unit's NameInRepository and GlobalName.
	GlobalName string `protobuf:"bytes,2,opt,name=GlobalName,proto3" json:"GlobalName,omitempty"`
	// Description is a short (~1-sentence) description of the source unit.
	Description string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	// TypeName is the human-readable name of the type of source unit; e.g., "Go
	// package".
	TypeName string `protobuf:"bytes,4,opt,name=TypeName,proto3" json:"TypeName,omitempty"`
}

func (m *UnitInfo) Reset()         { *m = UnitInfo{} }
func (m *UnitInfo) String() string { return proto.CompactTextString(m) }
func (*UnitInfo) ProtoMessage()    {}

type Unit struct {
	UnitKey `protobuf:"bytes,1,opt,name=key,embedded=key" json:""`
	// Globs is a list of patterns that match files that make up this source
	// unit. It is used to detect when the source unit definition is out of date
	// (e.g., when a file matches the glob but is not in the Files list).
	//
	// TODO(sqs): implement this in the Makefiles
	Globs []string `protobuf:"bytes,2,rep,name=globs" json:"globs,omitempty"`
	// Files is all of the files that make up this source unit. Filepaths should
	// be relative to the repository root.
	Files []string `protobuf:"bytes,3,rep,name=files" json:"files,omitempty"`
	// Dir is the root directory of this build unit. It is optional and maybe
	// empty.
	Dir string `protobuf:"bytes,4,opt,name=dir,proto3" json:"dir,omitempty"`
	// Dependencies is a list of dependencies that this source unit has. The
	// schema for these dependencies is internal to the scanner that produced
	// this source unit. The dependency resolver is expected to know how to
	// interpret this schema.
	//
	// The dependency information stored in this field should be able to be very
	// quickly determined by the scanner. The scanner should not perform any
	// dependency resolution on these entries. This is because the scanner is
	// run frequently and should execute very quickly, and dependency resolution
	// is often slow (requiring network access, etc.).
	RawDeps []*RawDep `protobuf:"bytes,5,rep,name=raw_deps" json:"raw_deps,omitempty"`
	Deps    []*Unit   `protobuf:"bytes,6,rep,name=deps" json:"deps,omitempty"`
	// DerivedFrom is the Unit from which this Unit is derived. The build unit
	// composed of the pip package "django", for instance, is derived from the
	// corresponding source unit in the repository github.com/django/django.
	DerivedFrom *Unit `protobuf:"bytes,7,opt,name=derived_from" json:"derived_from,omitempty"`
	// Info is an optional field that contains additional information used to
	// display the source unit
	Info *UnitInfo `protobuf:"bytes,8,opt,name=info" json:"info,omitempty"`
	// Data is additional data dumped by the scanner about this source unit. It
	// typically holds information that the scanner wants to make available to
	// other components in the toolchain (grapher, dep resolver, etc.).
	Data []byte `protobuf:"bytes,9,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Unit) Reset()         { *m = Unit{} }
func (m *Unit) String() string { return proto.CompactTextString(m) }
func (*Unit) ProtoMessage()    {}

type NodeKey struct {
	// Unit is the build unit that defines this definition.
	UnitKey `protobuf:"bytes,1,opt,name=unit,embedded=unit" json:""`
	// Path is a unique identifier for the def, relative to the build unit.
	// It should remain stable across commits as long as the def is the
	// "same" def. Its Elasticsearch mapping is defined separately (because
	// it is a multi_field, which the struct tag can't currently represent).
	//
	// Path encodes no structural semantics. Its only meaning is to be a stable
	// unique identifier within a given source unit. In many languages, it is
	// convenient to use the namespace hierarchy (with some modifications) as
	// the Path, but this may not always be the case. I.e., don't rely on Path
	// to find parents or children or any other structural propreties of the
	// def hierarchy). See Def.TreePath instead.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"Path"`
}

func (m *NodeKey) Reset()         { *m = NodeKey{} }
func (m *NodeKey) String() string { return proto.CompactTextString(m) }
func (*NodeKey) ProtoMessage()    {}

type Node struct {
	NodeKey `protobuf:"bytes,1,opt,name=key,embedded=key" json:""`
	Kind    string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	File    string `protobuf:"bytes,3,opt,name=file,proto3" json:"file,omitempty"`
	Start   uint32 `protobuf:"varint,4,opt,name=start,proto3" json:"start,omitempty"`
	End     uint32 `protobuf:"varint,5,opt,name=end,proto3" json:"end,omitempty"`
	// Def is definition metadata associated with the Node if it has kind "def".
	Def *DefData `protobuf:"bytes,6,opt,name=def" json:"def,omitempty"`
	// Doc is documentation metadata associated with the Node if it has kind "doc".
	Doc *DocData `protobuf:"bytes,7,opt,name=doc" json:"doc,omitempty"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}

type DefData struct {
	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Kind     string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Exported bool   `protobuf:"varint,3,opt,name=exported,proto3" json:"exported,omitempty"`
	Local    bool   `protobuf:"varint,4,opt,name=local,proto3" json:"local,omitempty"`
	Test     bool   `protobuf:"varint,5,opt,name=test,proto3" json:"test,omitempty"`
	Data     []byte `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *DefData) Reset()         { *m = DefData{} }
func (m *DefData) String() string { return proto.CompactTextString(m) }
func (*DefData) ProtoMessage()    {}

type DocData struct {
	// Format is the the MIME-type that the documentation is stored in. Valid
	// formats include 'text/html', 'text/plain', 'text/x-markdown', text/x-rst'.
	Format string `protobuf:"bytes,1,opt,name=format,proto3" json:"format,omitempty"`
	// Data is the actual documentation text.
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *DocData) Reset()         { *m = DocData{} }
func (m *DocData) String() string { return proto.CompactTextString(m) }
func (*DocData) ProtoMessage()    {}

// Edge describes the relationship between two nodes in the source graph.
// These can extend across build units.
type Edge struct {
	Src  NodeKey `protobuf:"bytes,1,opt,name=src" json:""`
	Dst  NodeKey `protobuf:"bytes,2,opt,name=dst" json:""`
	Kind string  `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
}

func (m *Edge) Reset()         { *m = Edge{} }
func (m *Edge) String() string { return proto.CompactTextString(m) }
func (*Edge) ProtoMessage()    {}
