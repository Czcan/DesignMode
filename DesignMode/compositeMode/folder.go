package main

import "fmt"

type folder struct {
	components []component
	name       string
}

func (f *folder) search(keyword string) {
	fmt.Printf("Searching for keyword %s in Folder %s \n", keyword, f.name)
	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *folder) add(cpt component) {
	f.components = append(f.components, cpt)
}
