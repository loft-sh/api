package v1

import (
	storagev1 "github.com/loft-sh/api/v5/pkg/apis/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StackInstanceView renders the detail page of one object that a stack declares in spec.ui. The
// caller posts the object to show; the platform reads it and its relations on the destination as
// the caller, and returns the declared cards with their values filled in.
// +subresource-request
type StackInstanceView struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Spec StackInstanceViewSpec `json:"spec,omitempty"`

	// +optional
	Status StackInstanceViewStatus `json:"status,omitempty"`
}

// StackInstanceViewSpec names what to show. Without a resource, the view lists the resources the
// stack declares, so a client can build its navigation. With a resource and no name, it lists the
// objects of a resource that declares columns. With both, it renders one object's detail page.
type StackInstanceViewSpec struct {
	// Resource is the name of a resource in the stack's spec.ui.
	// +optional
	Resource string `json:"resource,omitempty"`

	// Namespace is the namespace of the object on the destination. In a list, it limits the list
	// to that namespace.
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// Name is the name of the object.
	// +optional
	Name string `json:"name,omitempty"`

	// Action names an action of the resource to run on the object.
	// +optional
	Action string `json:"action,omitempty"`

	// Parameters are the values for the action's parameters, as text. The platform converts them
	// to the declared types and rejects a value that does not convert.
	// +optional
	Parameters map[string]string `json:"parameters,omitempty"`

	// Form asks for the create or edit form of the resource. With Object, it creates the object
	// named by namespace and name, or replaces the spec of that object.
	// +optional
	Form StackViewFormMode `json:"form,omitempty"`

	// Object is the submitted spec as JSON.
	// +optional
	Object string `json:"object,omitempty"`
}

// StackViewFormMode says whether a form creates an object or edits one.
// +enum
type StackViewFormMode string

const (
	StackViewFormCreate StackViewFormMode = "Create"
	StackViewFormEdit   StackViewFormMode = "Edit"
)

// StackViewForm is a spec form: the schema, the declared groups and the data to start from.
type StackViewForm struct {
	// Schema is the JSON Schema of the spec, as JSON.
	Schema string `json:"schema"`

	// Groups are the declared groups. Without them, a client shows every field in schema order.
	// +optional
	Groups []storagev1.StackUIFormGroup `json:"groups,omitempty"`

	// Data is the current spec for an edit, as JSON.
	// +optional
	Data string `json:"data,omitempty"`
}

// StackInstanceViewStatus holds the declared resources, or the rendered detail page.
type StackInstanceViewStatus struct {
	// Resources are the resources the stack declares, as the destination serves them.
	// +optional
	Resources []StackViewResource `json:"resources,omitempty"`

	// Cards are the declared cards whose condition holds, in order, with their values.
	// +optional
	Cards []StackViewCard `json:"cards,omitempty"`

	// Items are the listed objects, each with one cell per declared column.
	// +optional
	Items []StackViewItem `json:"items,omitempty"`

	// Actions are the actions offered on the shown object, with their parameter defaults.
	// +optional
	Actions []StackViewAction `json:"actions,omitempty"`

	// Result says what a run action or a submitted form changed.
	// +optional
	Result string `json:"result,omitempty"`

	// Form is the requested form.
	// +optional
	Form *StackViewForm `json:"form,omitempty"`

	// App describes the stack as an app of the platform UI. The view sets it with Resources.
	// +optional
	App *StackViewApp `json:"app,omitempty"`
}

// StackViewApp is what a client needs to show a stack as an app: its identity across
// installations, its name and icon, and where it appears.
type StackViewApp struct {
	// Template is the StackTemplate the stack uses. Installations of the same template are one
	// app in a project. Empty for a stack with an inline template.
	// +optional
	Template string `json:"template,omitempty"`

	DisplayName string `json:"displayName"`

	// +optional
	Description string `json:"description,omitempty"`

	// +optional
	Icon string `json:"icon,omitempty"`

	// Placements are the declared placements, or Cluster when the stack declares none.
	Placements []storagev1.StackUIPlacement `json:"placements"`
}

// StackViewAction is one action offered on the shown object.
type StackViewAction struct {
	Name string `json:"name"`

	Label string `json:"label"`

	// +optional
	Confirm string `json:"confirm,omitempty"`

	// Destructive marks a delete, so a client can style it.
	Destructive bool `json:"destructive"`

	// Primary marks the main action of the page.
	Primary bool `json:"primary"`

	// +optional
	Parameters []StackViewParameter `json:"parameters,omitempty"`
}

// StackViewParameter is one input of an offered action.
type StackViewParameter struct {
	Name string `json:"name"`

	Label string `json:"label"`

	Type storagev1.StackUIParameterType `json:"type"`

	Required bool `json:"required"`

	// Value is the initial value as text, from the declared default.
	// +optional
	Value string `json:"value,omitempty"`
}

// StackViewItem is one listed object.
type StackViewItem struct {
	// +optional
	Namespace string `json:"namespace,omitempty"`

	Name string `json:"name"`

	// Cells follow the order of the declared columns.
	// +optional
	Cells []StackViewRow `json:"cells,omitempty"`
}

// StackViewResource is a declared resource with the kind and scope the destination reports.
type StackViewResource struct {
	Name string `json:"name"`

	Label string `json:"label"`

	// +optional
	Icon string `json:"icon,omitempty"`

	// +optional
	Group string `json:"group,omitempty"`

	Version string `json:"version"`

	Resource string `json:"resource"`

	// Kind is empty when the destination does not serve the resource.
	// +optional
	Kind string `json:"kind,omitempty"`

	Namespaced bool `json:"namespaced"`

	// Columns are the labels of the declared columns. Without them, a client lists the resource
	// with the printer columns of its kind.
	// +optional
	Columns []string `json:"columns,omitempty"`

	// HasForm says the resource declares create and edit forms.
	HasForm bool `json:"hasForm"`

	// Error says why the destination does not serve the resource, for example a missing CRD.
	// +optional
	Error string `json:"error,omitempty"`
}

// StackViewCard is one rendered card: rows, a timeline or a log.
type StackViewCard struct {
	Title string `json:"title"`

	Type StackViewCardType `json:"type"`

	// +optional
	Rows []StackViewRow `json:"rows,omitempty"`

	// Events are the timeline entries, newest first.
	// +optional
	Events []StackViewEvent `json:"events,omitempty"`

	// Logs is the log of a logs card.
	// +optional
	Logs *StackViewLogs `json:"logs,omitempty"`

	// Error says why a timeline or a log could not be read.
	// +optional
	Error string `json:"error,omitempty"`
}

// StackViewCardType says what a card shows.
// +enum
type StackViewCardType string

const (
	StackViewCardRows     StackViewCardType = "Rows"
	StackViewCardTimeline StackViewCardType = "Timeline"
	StackViewCardLogs     StackViewCardType = "Logs"
)

// StackViewEvent is one Kubernetes event on a timeline.
type StackViewEvent struct {
	// Type is Normal or Warning.
	Type string `json:"type"`

	Reason string `json:"reason"`

	// +optional
	Message string `json:"message,omitempty"`

	// Object is the kind and name the event is about, for example "Pod/python-pool-7xk2p".
	Object string `json:"object"`

	// Time is when the event last happened, in RFC 3339.
	Time string `json:"time"`

	// Count is how often the event happened.
	Count int32 `json:"count"`
}

// StackViewLogs is the end of one container log.
type StackViewLogs struct {
	Pod string `json:"pod"`

	Container string `json:"container"`

	// +optional
	Lines string `json:"lines,omitempty"`
}

// StackViewRow is one rendered row. Values holds one entry per matched object; it is empty when
// the path selects nothing.
type StackViewRow struct {
	Label string `json:"label"`

	Widget storagev1.StackUIWidget `json:"widget"`

	// +optional
	Values []string `json:"values,omitempty"`

	// Tones holds the declared tone of each shown value that has one.
	// +optional
	Tones map[string]storagev1.StackUITone `json:"tones,omitempty"`

	// LinkTo is the stack resource that Link values point to.
	// +optional
	LinkTo string `json:"linkTo,omitempty"`

	// Conditions holds the conditions for the Conditions widget.
	// +optional
	Conditions []StackViewCondition `json:"conditions,omitempty"`

	// Error says why the row has no value, for example a relation the caller may not read.
	// +optional
	Error string `json:"error,omitempty"`
}

// StackViewCondition is one Kubernetes condition of the shown object.
type StackViewCondition struct {
	Type string `json:"type"`

	Status string `json:"status"`

	// +optional
	Reason string `json:"reason,omitempty"`

	// +optional
	Message string `json:"message,omitempty"`
}
