package main

import (
	"fmt"
	"log"

	"github.com/ghodss/yaml"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-tools/pkg/crd"
	crdmarkers "sigs.k8s.io/controller-tools/pkg/crd/markers"
	"sigs.k8s.io/controller-tools/pkg/loader"
	"sigs.k8s.io/controller-tools/pkg/markers"
)

func main() {
	pkgs, err := loader.LoadRoots("./pkg/apis/storage/v1")
	if err != nil {
		log.Fatal(err)
	}

	// find the virtual cluster package
	v1Pkg := pkgs[0]
	reg := &markers.Registry{}
	err = crdmarkers.Register(reg)
	if err != nil {
		log.Fatal(err)
	}

	parser := &crd.Parser{
		Collector: &markers.Collector{Registry: reg},
		Checker:   &loader.TypeChecker{},
	}
	crd.AddKnownTypes(parser)

	parser.NeedPackage(v1Pkg)

	groupKind := schema.GroupKind{Kind: "VirtualCluster", Group: "storage.loft.sh"}
	parser.NeedCRDFor(groupKind, nil)

	crd, ok := parser.CustomResourceDefinitions[groupKind]
	if ok {
		out, err := yaml.Marshal(crd)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(out))
	} else {
		log.Fatal("Not found")
	}
}
