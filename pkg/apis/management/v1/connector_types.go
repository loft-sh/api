package v1

import (
	agentstoragev1 "github.com/loft-sh/agentapi/v4/pkg/apis/loft/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +genclient:noStatus
// +genclient:skipVerbs=watch,deleteCollection
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Connector represents an integration connector of the platform (shared database,
// Argo CD, observability). It is a cluster-scoped virtual resource: it has no CRD of
// its own and is served straight from the connector Secret (labeled with
// loft.sh/connector-type) in the platform namespace, through the same translating
// client every other Secret-backed management kind uses. Watch is not served in v1: a
// watch request fails instead of hanging, so list and poll instead.
//
// Rejection happens at write time only: create and update through this API validate
// the request, while a read never rejects, so pre-existing Secrets with missing
// optional fields still project. A read does re-run that validation, but only to
// report the outcome in status, so a malformed Secret is described rather than hidden.
// Writers that bypass this API and write the backing Secrets directly (the UI and
// platform controllers today) are trusted to keep the Secrets well-formed; this API
// does not defend against out-of-band malformed Secrets.
//
// Status is read-only on the wire. Whatever a request carries is ignored: a create stores
// no status and an update keeps the stored object's. Status is assembled on every read,
// part of it computed from the stored payload and part of it unpacked from an annotation
// controllers publish on the backing Secret.
//
// Metadata is the backing Secret's. Labels and annotations are served and written as they are
// stored, as on every other management kind, with the exceptions each connector type declares.
// Projected keys become fields and are not served as metadata: the loft.sh/connector-type
// label is spec.type, the loft.sh/display-name annotation is spec.displayName. Server-owned keys
// are the platform's, never served, and a write that names one is refused: the status
// annotation controllers publish, unpacked into status. Withheld keys are never served: kubectl's
// last-applied-configuration annotation, because on a Secret written by hand it carries the
// whole manifest, credentials included. A write that sends a projected key disagreeing with the
// field it mirrors is refused. Labels the
// tenancy layer owns (tenant.vcluster.com/*) are that layer's, not this type's: an operator
// sees them and assigns a connector to a tenant by writing one, while a tenant never sees them
// and a tenant's write naming one is refused. Finalizers and ownerReferences round-trip.
// managedFields are served with this API's apiVersion. A label selector on
// loft.sh/connector-type still narrows a list to that type although the served object does not
// carry the label. Because the last-applied annotation is withheld, a client-side kubectl apply
// never finds it and reports a permanent diff: use server-side apply, imperative writes or
// full-replace updates.
//
// Credentials live in the type-specific payload sections of spec and are served only to
// a caller who could update this connector: one holding the update verb on it. Every
// other caller reads a redacted projection with the credential fields empty.
//
// Tenancy in this release: a platform admin creates connectors and assigns one to a
// tenant through the tenancy layer's exclusive assignment. A tenant reads the connectors
// assigned to it, redacted, and cannot create, update or delete a connector; every other
// connector is not found. Tenant-authored connectors come later, without a change to this
// type.
//
// An update through this API replaces the backing Secret from the request; nothing is read
// from the stored Secret to merge with. Passthrough labels and annotations, finalizers and
// ownerReferences round-trip because they are served and sent back. A credential the
// request leaves empty keeps its stored value, and the conditions controllers published
// are carried over from the stored object. What this API never serves does not survive an
// update: Data keys the connector type does not model, and the withheld annotations. The
// keys this API writes are the ones the platform's consumers of a connector read.
//
// The verbs behave as the generic path serves every Secret-backed kind. A delete is by
// name, with no precondition on the object's UID, and is reported as completed even when
// a finalizer on the backing Secret holds the Secret back. NotFound, AlreadyExists and
// Conflict speak of connectors; an Invalid or Forbidden raised by the store itself is
// passed through as the store phrased it.
// +k8s:openapi-gen=true
// +resource:path=connectors,rest=ConnectorREST
type Connector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConnectorSpec   `json:"spec,omitempty"`
	Status ConnectorStatus `json:"status,omitempty"`
}

// ConnectorType discriminates the connector types. The values are exactly the values
// of the loft.sh/connector-type label on the backing Secret, so projecting between
// spec.type and the label never requires a mapping.
// +enum
type ConnectorType string

const (
	// ConnectorTypeSharedDatabase connects the platform to a shared database server
	// used to provision databases as tenant cluster backing stores.
	ConnectorTypeSharedDatabase ConnectorType = "shared-database"

	// ConnectorTypeArgoCD connects the platform to an Argo CD (or Akuity) instance.
	ConnectorTypeArgoCD ConnectorType = "argocd"

	// ConnectorTypeObservability connects the platform to an observability stack.
	ConnectorTypeObservability ConnectorType = "observability"
)

// ConnectorSpec holds the specification
type ConnectorSpec struct {
	// Type is the connector type and selects which payload section below applies.
	// Required: create rejects an empty or unknown type and update rejects a change,
	// so the field is immutable after create. It is projected from and stamped to the
	// loft.sh/connector-type label on the backing Secret. Required fields carry no
	// omitempty, so a read always serializes them (as "" for a sparse pre-existing
	// Secret) and the object stays valid against the published schema, which lists
	// them as required.
	// +required
	Type ConnectorType `json:"type"`

	// DisplayName is the human-readable name shown in the UI. It is projected from
	// and stamped to the loft.sh/display-name annotation on the backing Secret.
	// +optional
	DisplayName string `json:"displayName,omitempty"`

	// SharedDatabase is the shared-database connector payload. It is required when
	// type is shared-database and must be unset for every other type. Its fields
	// mirror the Data keys of the backing Secret one to one, so a connector written
	// through this API is byte-compatible with one written by the UI directly.
	// +optional
	SharedDatabase *ConnectorSharedDatabaseSpec `json:"sharedDatabase,omitempty"`
}

// SharedDatabaseDialect is the database server dialect of a shared-database
// connector. It is stored in the "type" Data key of the backing Secret; the field is
// named dialect on the wire because spec.type is the connector type discriminator and
// must not be conflated with the database dialect.
type SharedDatabaseDialect string

const (
	// SharedDatabaseDialectMySQL is the MySQL dialect. It is the platform's
	// historical default: consumers treat a connector Secret without an explicit
	// dialect as MySQL.
	SharedDatabaseDialectMySQL SharedDatabaseDialect = "mysql"

	// SharedDatabaseDialectPostgres is the PostgreSQL dialect.
	SharedDatabaseDialectPostgres SharedDatabaseDialect = "postgres"
)

// ConnectorSharedDatabaseSpec configures a connection to a shared database server
// used to provision databases as tenant cluster backing stores. Every field maps to
// one Data key of the backing Secret (given in each field comment), so direct Secret
// writers and this API stay byte-compatible in both directions.
type ConnectorSharedDatabaseSpec struct {
	// Dialect is the database server dialect, mysql or postgres. Data key "type"
	// (distinct from spec.type, the connector type). Required on every write: there
	// is no server-side default and unknown values are rejected. Reads of a
	// pre-existing Secret without the key still project mysql, matching how the read
	// consumers treat such Secrets.
	// +required
	Dialect SharedDatabaseDialect `json:"dialect"`

	// Endpoint is the host of the database server. Data key "endpoint". Required.
	// +required
	Endpoint string `json:"endpoint"`

	// Port is the port of the database server, kept as a string exactly as stored in
	// the Data key "port". An empty port defaults by dialect on CREATE ONLY: 5432
	// for postgres, 3306 for mysql (the same defaults the UI applies). Updates never
	// default it. For postgres an empty port on update removes the key and the driver
	// falls back to 5432; for mysql an empty port is rejected on update, because the
	// mysql driver has no fallback and would dial port 0. A stored mysql connector
	// without a port is therefore rejected on every update, whatever else the update
	// changes, until the port is set; such a connector cannot connect until then
	// anyway, and its Configured condition says so.
	// +optional
	Port string `json:"port,omitempty"`

	// User is the database user the platform connects as. Data key "user". Required.
	// +required
	User string `json:"user"`

	// Password is the password of the database user. Data key "password". Required
	// on create unless identityProvider is set. It is a credential: served only to a
	// caller who could update this connector and empty for every other caller (see
	// the Connector type). On update an empty password keeps the stored one, so a
	// redacted read written back never clears a credential; switching to
	// identityProvider removes it. Switching back from identityProvider to password
	// authentication requires a non-empty password in the request: the stored one
	// was removed by the earlier switch, so there is nothing to keep.
	// +optional
	Password string `json:"password,omitempty"`

	// IdentityProvider selects password-less IAM authentication. Data key
	// "identityProvider". The only supported value is "aws"; other values are
	// rejected. When set, password must be empty and the stored password key is
	// removed, matching the UI behavior.
	// +optional
	IdentityProvider string `json:"identityProvider,omitempty"`

	// AWSPermissionBoundaryPolicyARN is the ARN of the IAM policy attached as the
	// permissions boundary of every IAM role the platform provisions for a tenant
	// cluster through this connector. It is only read when identityProvider is set.
	// Unlike the other fields it maps to an annotation on the backing Secret
	// (connector.platform.vcluster.com/aws-permission-boundary-policy-arn), not to a
	// Data key, and this field is the one way to write it: the annotation is
	// projected, so a metadata write carrying a different value is refused. Empty
	// means no boundary, which is a supported configuration, and an empty value on
	// update removes the annotation, so a read-modify-write that drops the field
	// clears a boundary that was set.
	// +optional
	AWSPermissionBoundaryPolicyARN string `json:"awsPermissionBoundaryPolicyArn,omitempty"`

	// CACert is the PEM-encoded CA bundle used to verify the database server's TLS
	// certificate. Data key "caCert".
	// +optional
	CACert string `json:"caCert,omitempty"`

	// SSLMode is an explicit Postgres sslmode value (disable, allow, prefer, require,
	// verify-ca or verify-full) overriding the default policy derived from caCert.
	// Data key "sslMode".
	// +optional
	SSLMode string `json:"sslMode,omitempty"`
}

// ConnectorStatus reports what the platform knows about a connector. Nothing here is
// settable through this API: a create drops whatever status the request carried and an
// update replaces it with the stored object's, so a client can never write status and
// can never clobber what a controller published.
//
// Conditions are an open list in the house Cluster-API shape, so a controller may add a
// condition type this API does not name without an API change. Two sources fill the
// list. The Configured condition is computed on read by validating the stored payload
// with the type's codec, so it needs no controller and costs no persistence. Everything
// only a controller can know (reachability, reconcile outcomes, the observability
// gateway observations) is published by controllers into a server-owned annotation on
// the backing Secret and unpacked here on read. No caller ever sees that annotation: it
// is server-owned, so it is never served, a write naming it is refused, and it stays on
// the Secret as the controller left it.
type ConnectorStatus struct {
	// Conditions holds the conditions observed for this connector. The list is open:
	// v1 names Configured, Reachable (and Authenticated where credentials fail
	// separately from connectivity) and Reconciled, but a controller may publish
	// other types.
	//
	// The computed Configured condition reports the connector's own
	// metadata.creationTimestamp as its lastTransitionTime, unless a controller has
	// published a Configured in the same state, in which case that condition's
	// transition time is adopted. The API recomputes this condition on every read and
	// keeps no record of when the outcome last flipped, so the creation time is a
	// stand-in that the real transition can only have happened at or after, never a
	// transition the server observed.
	// +optional
	Conditions agentstoragev1.Conditions `json:"conditions,omitempty"`
}
