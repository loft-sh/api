package v1

import (
	"encoding/json"
	"strings"
	"testing"
)

// omitempty never omits a struct, so a nested struct held by value is written out even when
// nothing was set, and the difference shows nowhere but in the serialized object. These pin which
// of the host validation fields disappear when unset and which are there whatever happens.
func TestHostValidationFieldsOmitWhatIsUnset(t *testing.T) {
	for _, test := range []struct {
		name    string
		object  any
		key     string
		present bool
	}{
		{
			name:   "a node type that requires no profile",
			object: MachineNodeTypeSpec{NamedNodeTypeSpec: NamedNodeTypeSpec{Name: "small"}},
			key:    "hostValidationProfiles",
		},
		{
			name: "a node type that requires profiles",
			object: MachineNodeTypeSpec{
				NamedNodeTypeSpec:      NamedNodeTypeSpec{Name: "small"},
				HostValidationProfiles: &NodeTypeHostValidationProfiles{},
			},
			key:     "hostValidationProfiles",
			present: true,
		},
		{
			name:   "a provider whose machines are not host validated",
			object: NodeProviderNICo{Endpoints: &NICoEndpoints{API: ServiceEndpoint{URL: "https://nico.example.com"}}},
			key:    "bmcProxy",
		},
		{
			name:    "the REST API endpoint, which every endpoints block names",
			object:  NICoEndpoints{},
			key:     "api",
			present: true,
		},
		{
			name:   "a tenant that never reached its provider",
			object: TenantNICoStatus{},
			key:    "api",
		},
		{
			name:   "a machine never checked against a profile",
			object: MachineStatus{},
			key:    "hostValidation",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			encoded, err := json.Marshal(test.object)
			if err != nil {
				t.Fatal(err)
			}
			if got := strings.Contains(string(encoded), `"`+test.key+`"`); got != test.present {
				if test.present {
					t.Fatalf("%q is missing: %s", test.key, encoded)
				}
				t.Fatalf("%q is written although nothing set it: %s", test.key, encoded)
			}
		})
	}
}
