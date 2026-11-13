package vclusterconfig

import (
	"fmt"
	"testing"

	"k8s.io/apimachinery/pkg/util/validation/field"
)

// wantErr asserts the error kind and the exact field path, which a count alone would not catch.
type wantErr struct {
	typ  field.ErrorType
	path string
}

func stackWithTemplate(name string, tasks ...StackTaskConfig) StackConfig {
	return StackConfig{Name: name, Template: &StackTemplateDefinitionConfig{Tasks: tasks}}
}

func stackWithRef(name, ref string) StackConfig {
	return StackConfig{Name: name, TemplateRef: &StackTemplateRefConfig{Name: ref}}
}

func TestValidateStacks(t *testing.T) {
	fldPath := field.NewPath("spec")

	tests := []struct {
		name   string
		stacks []StackConfig
		want   []wantErr
	}{
		{
			name:   "no stacks",
			stacks: nil,
		},
		{
			name:   "valid inline template",
			stacks: []StackConfig{stackWithTemplate("monitoring", StackTaskConfig{Name: "install"})},
		},
		{
			name:   "valid template ref",
			stacks: []StackConfig{stackWithRef("monitoring", "kube-prometheus")},
		},
		{
			name: "valid with every optional field set",
			stacks: []StackConfig{{
				Name:        "monitoring",
				DisplayName: "Monitoring",
				Description: "prometheus",
				PrunePolicy: "Prune",
				Defaults:    &StackDefaultsConfig{TaskTimeout: "10m"},
				Parameters:  map[string]interface{}{"replicas": 2},
				Template: &StackTemplateDefinitionConfig{
					Tasks: []StackTaskConfig{{Name: "install", Timeout: "5m"}},
				},
			}},
		},
		{
			name:   "missing name",
			stacks: []StackConfig{stackWithTemplate("")},
			want:   []wantErr{{field.ErrorTypeRequired, "spec.deploy.stacks[0].name"}},
		},
		{
			name:   "name is not a DNS label",
			stacks: []StackConfig{stackWithTemplate("Not_A_Label")},
			want:   []wantErr{{field.ErrorTypeInvalid, "spec.deploy.stacks[0].name"}},
		},
		{
			// Trimming instead would accept a name the apiserver later rejects.
			name:   "leading space in name is rejected, not trimmed",
			stacks: []StackConfig{stackWithTemplate(" monitoring")},
			want:   []wantErr{{field.ErrorTypeInvalid, "spec.deploy.stacks[0].name"}},
		},
		{
			name: "duplicate name reported on the second entry",
			stacks: []StackConfig{
				stackWithTemplate("monitoring"),
				stackWithTemplate("monitoring"),
			},
			want: []wantErr{{field.ErrorTypeDuplicate, "spec.deploy.stacks[1].name"}},
		},
		{
			name:   "neither template nor templateRef",
			stacks: []StackConfig{{Name: "monitoring"}},
			want:   []wantErr{{field.ErrorTypeRequired, "spec.deploy.stacks[0]"}},
		},
		{
			name: "both template and templateRef",
			stacks: []StackConfig{{
				Name:        "monitoring",
				Template:    &StackTemplateDefinitionConfig{},
				TemplateRef: &StackTemplateRefConfig{Name: "kube-prometheus"},
			}},
			want: []wantErr{{field.ErrorTypeForbidden, "spec.deploy.stacks[0]"}},
		},
		{
			name:   "templateRef without a name",
			stacks: []StackConfig{stackWithRef("monitoring", "")},
			want:   []wantErr{{field.ErrorTypeRequired, "spec.deploy.stacks[0].templateRef.name"}},
		},
		{
			name: "unsupported prunePolicy",
			stacks: []StackConfig{{
				Name:        "monitoring",
				PrunePolicy: "Nope",
				Template:    &StackTemplateDefinitionConfig{},
			}},
			want: []wantErr{{field.ErrorTypeNotSupported, "spec.deploy.stacks[0].prunePolicy"}},
		},
		{
			name: "unparseable defaults taskTimeout",
			stacks: []StackConfig{{
				Name:     "monitoring",
				Defaults: &StackDefaultsConfig{TaskTimeout: "5 bananas"},
				Template: &StackTemplateDefinitionConfig{},
			}},
			want: []wantErr{{field.ErrorTypeInvalid, "spec.deploy.stacks[0].defaults.taskTimeout"}},
		},
		{
			name: "unparseable task timeout names the task index",
			stacks: []StackConfig{stackWithTemplate("monitoring",
				StackTaskConfig{Name: "first", Timeout: "5m"},
				StackTaskConfig{Name: "second", Timeout: "nope"},
			)},
			want: []wantErr{{field.ErrorTypeInvalid, "spec.deploy.stacks[0].template.tasks[1].timeout"}},
		},
		{
			// Early returns instead of collecting would drop all but the first.
			name: "several problems in one entry are all reported",
			stacks: []StackConfig{{
				Name:        "Bad_Name",
				PrunePolicy: "Nope",
				Defaults:    &StackDefaultsConfig{TaskTimeout: "nope"},
			}},
			want: []wantErr{
				{field.ErrorTypeInvalid, "spec.deploy.stacks[0].name"},
				{field.ErrorTypeRequired, "spec.deploy.stacks[0]"},
				{field.ErrorTypeNotSupported, "spec.deploy.stacks[0].prunePolicy"},
				{field.ErrorTypeInvalid, "spec.deploy.stacks[0].defaults.taskTimeout"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertErrs(t, ValidateStacks(fldPath, tt.stacks), tt.want)
		})
	}
}

// TestValidateStacks_OverCapShortCircuits checks the cap replaces the per-entry errors, not adds to them.
func TestValidateStacks_OverCapShortCircuits(t *testing.T) {
	stacks := make([]StackConfig, 0, MaxStacks+1)
	// Also invalid, so a missing short-circuit shows up as extra errors.
	stacks = append(stacks, StackConfig{})
	for i := 1; i <= MaxStacks; i++ {
		stacks = append(stacks, stackWithTemplate(fmt.Sprintf("stack-%d", i)))
	}

	errs := ValidateStacks(field.NewPath("spec"), stacks)
	assertErrs(t, errs, []wantErr{{field.ErrorTypeTooMany, "spec.deploy.stacks"}})
}

// TestValidateStacks_AtCapIsAccepted checks the boundary is > and not >=.
func TestValidateStacks_AtCapIsAccepted(t *testing.T) {
	stacks := make([]StackConfig, 0, MaxStacks)
	for i := range MaxStacks {
		stacks = append(stacks, stackWithTemplate(fmt.Sprintf("stack-%d", i)))
	}

	if errs := ValidateStacks(field.NewPath("spec"), stacks); len(errs) != 0 {
		t.Fatalf("exactly MaxStacks stacks must be accepted, got %v", errs)
	}
}

func assertErrs(t *testing.T, errs field.ErrorList, want []wantErr) {
	t.Helper()

	if len(errs) != len(want) {
		t.Fatalf("got %d errors, want %d: %v", len(errs), len(want), errs)
	}
	for i, w := range want {
		if errs[i].Type != w.typ {
			t.Errorf("error %d: got type %q, want %q (%v)", i, errs[i].Type, w.typ, errs[i])
		}
		if errs[i].Field != w.path {
			t.Errorf("error %d: got path %q, want %q", i, errs[i].Field, w.path)
		}
	}
}

// TestValidateStackList pins which rules fail the whole sync. Moving one into ValidateStack would
// quietly downgrade it to a single skipped stack.
func TestValidateStackList(t *testing.T) {
	fldPath := field.NewPath("spec")

	tests := []struct {
		name   string
		stacks []StackConfig
		want   []wantErr
	}{
		{
			name:   "no stacks",
			stacks: nil,
		},
		{
			name:   "distinct names",
			stacks: []StackConfig{stackWithTemplate("monitoring"), stackWithTemplate("runai")},
		},
		{
			name:   "missing name",
			stacks: []StackConfig{stackWithTemplate("")},
			want:   []wantErr{{field.ErrorTypeRequired, "spec.deploy.stacks[0].name"}},
		},
		{
			name:   "two missing names are reported separately",
			stacks: []StackConfig{stackWithTemplate(""), stackWithTemplate("")},
			want: []wantErr{
				{field.ErrorTypeRequired, "spec.deploy.stacks[0].name"},
				{field.ErrorTypeRequired, "spec.deploy.stacks[1].name"},
			},
		},
		{
			name:   "duplicate name reported on the second entry",
			stacks: []StackConfig{stackWithTemplate("monitoring"), stackWithTemplate("monitoring")},
			want:   []wantErr{{field.ErrorTypeDuplicate, "spec.deploy.stacks[1].name"}},
		},
		{
			// Both are per-entry problems. Reporting either here would abort the sync instead of
			// isolating the one bad entry.
			name:   "per-entry problems are not list level",
			stacks: []StackConfig{{Name: "monitoring", PrunePolicy: "Nope"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertErrs(t, ValidateStackList(fldPath, tt.stacks), tt.want)
		})
	}
}

// TestValidateStackList_OverCapIsListLevel keeps the cap on the list. No single entry owns it.
func TestValidateStackList_OverCapIsListLevel(t *testing.T) {
	stacks := make([]StackConfig, 0, MaxStacks+1)
	for i := 0; i <= MaxStacks; i++ {
		stacks = append(stacks, stackWithTemplate(fmt.Sprintf("stack-%d", i)))
	}

	assertErrs(t, ValidateStackList(field.NewPath("spec"), stacks), []wantErr{
		{field.ErrorTypeTooMany, "spec.deploy.stacks"},
	})
}

// TestValidateStack_LeavesNameRulesToTheList checks one name problem is never reported twice.
func TestValidateStack_LeavesNameRulesToTheList(t *testing.T) {
	stackPath := field.NewPath("spec", "deploy", "stacks").Index(0)

	if errs := ValidateStack(stackPath, stackWithTemplate("")); len(errs) != 0 {
		t.Fatalf("an empty name is ValidateStackList's to report, got %v", errs)
	}
	if errs := ValidateStack(stackPath, stackWithTemplate("monitoring")); len(errs) != 0 {
		t.Fatalf("a valid entry must produce no errors, got %v", errs)
	}
}
