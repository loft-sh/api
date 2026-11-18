package v1

// StackUI declares the resources a stack adds to the platform UI. The platform renders them with
// its own components; a stack ships no frontend code.
type StackUI struct {
	// Placements say where the platform shows the resources. The default is Cluster.
	// +optional
	Placements []StackUIPlacement `json:"placements,omitempty"`

	// Resources are the kinds the stack provides, each shown as its own page on the destination.
	// +optional
	Resources []StackUIResource `json:"resources,omitempty"`
}

// StackUIPlacement is a place in the platform UI that shows a stack's resources.
// +enum
type StackUIPlacement string

const (
	// StackUIPlacementCluster shows the resources on the page of the destination cluster.
	StackUIPlacementCluster StackUIPlacement = "Cluster"
	// StackUIPlacementProject shows the stack as an entry of its project, next to Instances. The
	// entry gathers every installation of the same StackTemplate in the project, so only stacks
	// that use a templateRef appear there.
	StackUIPlacementProject StackUIPlacement = "Project"
)

// StackUIResource is one kind the stack provides on its destination.
type StackUIResource struct {
	// Name identifies the resource within the stack. It appears in UI paths, so it is a DNS label.
	Name string `json:"name"`

	// Label is the plural display name, for example "Sandboxes".
	Label string `json:"label"`

	// Icon is an image URL shown next to the label.
	// +optional
	Icon string `json:"icon,omitempty"`

	// Group is the API group of the kind. Empty means the core group.
	// +optional
	Group string `json:"group,omitempty"`

	// Version is the API version of the kind.
	Version string `json:"version"`

	// Resource is the plural resource name, for example "sandboxes".
	Resource string `json:"resource"`

	// LabelSelector limits the list to the objects that belong to the stack, in kubectl selector
	// syntax, for example "app.kubernetes.io/part-of=web-services". Without it the list shows
	// every object of the kind, which suits a kind the stack owns, such as its own CRD, but not a
	// built-in kind like Deployments. A resource with a selector declares its columns, because the
	// platform applies the selector when it fills them.
	// +optional
	LabelSelector string `json:"labelSelector,omitempty"`

	// Columns replace the printer columns of the kind on the list page. A column reads the listed
	// object only, so it may not set From or When.
	// +optional
	Columns []StackUIRow `json:"columns,omitempty"`

	// Relations name the objects a detail page reads next to the object itself. A relation name is
	// also the CEL variable that holds its objects, so it is a CEL identifier.
	// +optional
	Relations []StackUIRelation `json:"relations,omitempty"`

	// Detail is the layout of the detail page. Without it, the page shows the object as YAML.
	// +optional
	Detail *StackUIDetail `json:"detail,omitempty"`

	// Actions are the operations a detail page offers on the object. The platform runs them on
	// the destination as the user who clicks them.
	// +optional
	Actions []StackUIAction `json:"actions,omitempty"`

	// Form enables create and edit forms for the object's spec, built from the schema the
	// destination serves for the kind. Without it, the resource offers no form.
	// +optional
	Form *StackUIForm `json:"form,omitempty"`
}

// StackUIForm lays out the spec form. Without groups, the form shows every spec field in schema order.
type StackUIForm struct {
	// Groups are shown in order, each with its fields.
	// +optional
	Groups []StackUIFormGroup `json:"groups,omitempty"`
}

// StackUIFormGroup is one titled group of form fields.
type StackUIFormGroup struct {
	// Label is the group heading.
	Label string `json:"label"`

	// Fields are dotted paths under spec, for example "warmPoolRef.name".
	Fields []string `json:"fields"`
}

// StackUIAction is one operation on a detail page. Exactly one of Delete, Patch and Create is set.
type StackUIAction struct {
	// Name identifies the action within the resource. It is a DNS label.
	Name string `json:"name"`

	// Label is the button text.
	Label string `json:"label"`

	// When is a CEL expression. The action is offered only when it is true.
	// +optional
	When string `json:"when,omitempty"`

	// Confirm is a question the user answers before the action runs.
	// +optional
	Confirm string `json:"confirm,omitempty"`

	// Primary marks the main action of the page. A client shows it as the one prominent button.
	// At most one action of a resource is primary, and a delete cannot be.
	// +optional
	Primary bool `json:"primary,omitempty"`

	// Parameters are the inputs the user fills in. Expressions read them under params.
	// +optional
	Parameters []StackUIActionParameter `json:"parameters,omitempty"`

	// Delete deletes the object, or the objects of a relation.
	// +optional
	Delete *StackUIActionTarget `json:"delete,omitempty"`

	// Patch applies a JSON merge patch to the object, or to the objects of a relation.
	// +optional
	Patch *StackUIPatchAction `json:"patch,omitempty"`

	// Create creates an object of a resource of this stack, in the object's namespace.
	// +optional
	Create *StackUICreateAction `json:"create,omitempty"`
}

// StackUIActionTarget names what an action changes.
type StackUIActionTarget struct {
	// Relation names a relation whose objects the action changes. Empty changes the object itself.
	// +optional
	Relation string `json:"relation,omitempty"`
}

// StackUIPatchAction changes objects with a JSON merge patch.
type StackUIPatchAction struct {
	StackUIActionTarget `json:",inline"`

	// Patch is a CEL expression that builds the merge patch as a map, for example
	// "{'spec': {'replicas': params.replicas}}".
	Patch string `json:"patch"`
}

// StackUICreateAction creates one object.
type StackUICreateAction struct {
	// Resource names a resource of this stack. The platform sets apiVersion, kind and namespace.
	Resource string `json:"resource"`

	// Object is a CEL expression that builds the object as a map, for example
	// "{'metadata': {'generateName': self.metadata.name + '-'}, 'spec': {...}}".
	Object string `json:"object"`
}

// StackUIActionParameter is one input of an action.
type StackUIActionParameter struct {
	// Name is the key under params. It is a CEL identifier.
	Name string `json:"name"`

	// Label is the input label.
	Label string `json:"label"`

	// Type is the value type. The default is String.
	// +optional
	Type StackUIParameterType `json:"type,omitempty"`

	// Required rejects an empty value.
	// +optional
	Required bool `json:"required,omitempty"`

	// Default is a CEL expression for the initial value, for example "self.spec.replicas".
	// +optional
	Default string `json:"default,omitempty"`
}

// StackUIParameterType is the type of an action parameter.
// +enum
type StackUIParameterType string

const (
	StackUIParameterString  StackUIParameterType = "String"
	StackUIParameterNumber  StackUIParameterType = "Number"
	StackUIParameterBoolean StackUIParameterType = "Boolean"
)

// StackUIRelation finds objects related to the object a detail page shows. Exactly one of
// Owned, FieldRef and LabelSelector is set. Related objects are read from the object's namespace.
type StackUIRelation struct {
	// Name identifies the relation within the resource. Rows read from it by this name.
	Name string `json:"name"`

	// Group is the API group of the related kind. Empty means the core group.
	// +optional
	Group string `json:"group,omitempty"`

	// Version is the API version of the related kind.
	Version string `json:"version"`

	// Resource is the plural resource name of the related kind.
	Resource string `json:"resource"`

	// Owned selects the objects that carry an owner reference to the object.
	// +optional
	Owned bool `json:"owned,omitempty"`

	// FieldRef is a JSONPath on the object that holds the name of the related object.
	// +optional
	FieldRef string `json:"fieldRef,omitempty"`

	// LabelSelector selects the related objects by a label whose value comes from the object.
	// +optional
	LabelSelector *StackUILabelRelation `json:"labelSelector,omitempty"`
}

// StackUILabelRelation selects related objects whose label Key equals the value at ValuePath.
type StackUILabelRelation struct {
	// Key is the label key on the related objects.
	Key string `json:"key"`

	// ValuePath is a JSONPath on the object that gives the label value.
	ValuePath string `json:"valuePath"`
}

// StackUIDetail is the layout of a detail page: cards of labelled rows.
type StackUIDetail struct {
	// Cards are shown in order.
	Cards []StackUICard `json:"cards"`
}

// StackUICard is one titled card. Exactly one of Rows, Timeline and Logs is set.
type StackUICard struct {
	// Title is the card heading.
	Title string `json:"title"`

	// When is a CEL expression. The card shows only when it is true.
	// +optional
	When string `json:"when,omitempty"`

	// Rows are labelled values, shown in order.
	// +optional
	Rows []StackUIRow `json:"rows,omitempty"`

	// Timeline shows the Kubernetes events of the object, newest first.
	// +optional
	Timeline *StackUITimeline `json:"timeline,omitempty"`

	// Logs shows the end of a container log.
	// +optional
	Logs *StackUILogs `json:"logs,omitempty"`
}

// StackUITimeline lists events. It always includes the events of the object itself.
type StackUITimeline struct {
	// Relations name relations whose objects' events join the timeline, for example the pod of
	// a sandbox.
	// +optional
	Relations []string `json:"relations,omitempty"`
}

// StackUILogs reads the log of one pod container.
type StackUILogs struct {
	// From names a relation whose first object is the pod. Empty reads the object itself, which
	// must then be a pod.
	// +optional
	From string `json:"from,omitempty"`

	// Container is the container name. Empty reads the pod's only container.
	// +optional
	Container string `json:"container,omitempty"`

	// TailLines is how many lines to show. The default is 100 and the maximum is 1000.
	// +optional
	TailLines int64 `json:"tailLines,omitempty"`
}

// StackUIRow is one labelled value on a card.
type StackUIRow struct {
	// Label is the row name.
	Label string `json:"label"`

	// From names the relation the row reads. Empty reads the object itself. A relation that
	// matches more than one object gives one value per object.
	// +optional
	From string `json:"from,omitempty"`

	// Path is a JSONPath, for example "{.status.podIP}". Exactly one of Path and Expression is set.
	// +optional
	Path string `json:"path,omitempty"`

	// Expression is a CEL expression for a derived value. It sees the object as self and each
	// relation as a list under its name, for example
	// "size(claim) > 0 ? 'Adopted' : 'Waiting'". A list result gives one value per entry.
	// +optional
	Expression string `json:"expression,omitempty"`

	// When is a CEL expression. The row shows only when it is true.
	// +optional
	When string `json:"when,omitempty"`

	// Widget selects how the value is shown. The default is Text.
	// +optional
	Widget StackUIWidget `json:"widget,omitempty"`

	// LinkTo names a resource of this stack. With the Link widget, the value is the name of an
	// object of that resource, in the object's namespace.
	// +optional
	LinkTo string `json:"linkTo,omitempty"`

	// Values maps raw values to display values, for example "vnode" to "vNode".
	// +optional
	Values map[string]string `json:"values,omitempty"`

	// Tones maps display values to the tone of their Badge, for example "Ready" to Success. A
	// value without a tone is Neutral.
	// +optional
	Tones map[string]StackUITone `json:"tones,omitempty"`
}

// StackUITone is the color meaning of a badge.
// +enum
type StackUITone string

const (
	StackUIToneNeutral StackUITone = "Neutral"
	StackUIToneInfo    StackUITone = "Info"
	StackUIToneSuccess StackUITone = "Success"
	StackUIToneWarning StackUITone = "Warning"
	StackUIToneDanger  StackUITone = "Danger"
)

// StackUIWidget is how a row shows its value.
// +enum
type StackUIWidget string

const (
	// StackUIWidgetText shows the value as text.
	StackUIWidgetText StackUIWidget = "Text"
	// StackUIWidgetMono shows the value in a monospace font.
	StackUIWidgetMono StackUIWidget = "Mono"
	// StackUIWidgetBadge shows the value as a status badge.
	StackUIWidgetBadge StackUIWidget = "Badge"
	// StackUIWidgetTime shows an RFC 3339 time as a relative time.
	StackUIWidgetTime StackUIWidget = "Time"
	// StackUIWidgetLink shows the value as a link to an object of the LinkTo resource.
	StackUIWidgetLink StackUIWidget = "Link"
	// StackUIWidgetConditions shows a list of Kubernetes conditions as chips.
	StackUIWidgetConditions StackUIWidget = "Conditions"
)
