/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package boot

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var versionedAPIs string

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates source code for a repo",
	Long:  `Generates source code for a repo`,
	Run:   RunGenerate,
}

const genericAPI = "k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/apimachinery/pkg/api/resource/,k8s.io/apimachinery/pkg/version/,k8s.io/apimachinery/pkg/runtime/,k8s.io/apimachinery/pkg//util/intstr/"
const extraAPI = "k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/apimachinery/pkg/conversion,k8s.io/apimachinery/pkg/runtime"

func AddGenerate(cmd *cobra.Command) {
	cmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVar(&copyright, "copyright", "boilerplate.go.txt", "path to copyright file.  defaults to boilerplate.go.txt")
	generateCmd.Flags().StringVar(&versionedAPIs, "api-versions", "", "comma separated list of APIs versions.  e.g. foo/v1beta1,bar/v1  defaults to all directories under pkd/apis/group/version")
}

func RunGenerate(cmd *cobra.Command, args []string) {
	if len(versionedAPIs) == 0 {
		groups, err := ioutil.ReadDir(filepath.Join("pkg", "apis"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not read pkg/apis directory to find api versions\n")
			os.Exit(-1)
		}
		versions := []string{}
		for _, g := range groups {
			if g.IsDir() {
				versionFiles, err := ioutil.ReadDir(filepath.Join("pkg", "apis", g.Name()))
				if err != nil {
					fmt.Fprintf(os.Stderr, "could not read pkg/apis/%s directory to find api versions\n", g.Name())
					os.Exit(-1)
				}
				for _, v := range versionFiles {
					if v.IsDir() {
						versions = append(versions)
					}
				}
			}
		}
	}

	getCopyright()

	root, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(-1)
	}
	root = filepath.Dir(root)

	src := filepath.Join(Repo, "pkg", "apis", "...")

	c := exec.Command(filepath.Join(root, "apiregister-gen"),
		"-i", src)
	fmt.Printf("%s\n", strings.Join(c.Args, " "))
	out, err := c.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to run apiregister-gen %s %v\n", out, err)
		os.Exit(-1)
	}

	c = exec.Command(filepath.Join(root, "conversion-gen"),
		"-o", GoSrc,
		"--go-header-file", copyright,
		"-O", "zz_generated.conversion",
		"-i", src,
		"--extra-peer-dirs", extraAPI,
	)
	fmt.Printf("%s\n", strings.Join(c.Args, " "))
	out, err = c.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to run conversion-gen %s %v\n", out, err)
		os.Exit(-1)
	}

	c = exec.Command(filepath.Join(root, "deepcopy-gen"),
		"-o", GoSrc,
		"--go-header-file", copyright,
		"-O", "zz_generated.deepcopy",
		"-i", src,
	)
	fmt.Printf("%s\n", strings.Join(c.Args, " "))
	out, err = c.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to run deepcopy-gen %s %v\n", out, err)
		os.Exit(-1)
	}

	c = exec.Command(filepath.Join(root, "openapi-gen"),
		"-o", GoSrc,
		"--go-header-file", copyright,
		"-i", src+","+genericAPI,
		"--output-package", filepath.Join(Repo, "pkg", "openapi"),
	)
	fmt.Printf("%s\n", strings.Join(c.Args, " "))
	out, err = c.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to run openapi-gen %s %v\n", out, err)
		os.Exit(-1)
	}

	c = exec.Command(filepath.Join(root, "defaulter-gen"),
		"-o", GoSrc,
		"--go-header-file", copyright,
		"-O", "zz_generated.defaults",
		"-i", src,
		"--extra-peer-dirs=", extraAPI,
	)
	fmt.Printf("%s\n", strings.Join(c.Args, " "))
	out, err = c.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to run defaulter-gen %s %v\n", out, err)
		os.Exit(-1)
	}

	// Builder the versioned apis client
	clientPkg := filepath.Join(Repo, "pkg", "client")
	clientset := filepath.Join(clientPkg, "clientset_generated")
	c = exec.Command(filepath.Join(root, "client-gen"),
		"-o", GoSrc,
		"--go-header-file", copyright,
		"--input-base", filepath.Join(Repo, "pkg", "apis"),
		"--input", versionedAPIs,
		"--clientset-path", clientset,
		"--clientset-name", "clientset",
	)
	fmt.Printf("%s\n", strings.Join(c.Args, " "))
	out, err = c.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to run client-gen %s %v\n", out, err)
		os.Exit(-1)
	}

	// Build the unversioned apis client
	unversioned := map[string]bool{}
	for _, a := range strings.Split(versionedAPIs, ",") {
		unversioned[path.Dir(a)] = true
	}
	u := []string{}
	for a, _ := range unversioned {
		fmt.Printf("found %s\n", a)
		u = append(u, a)
	}
	unversionedAPIs := strings.Join(u, ",")

	c = exec.Command(filepath.Join(root, "client-gen"),
		"-o", GoSrc,
		"--go-header-file", copyright,
		"--input-base", filepath.Join(Repo, "pkg", "apis"),
		"--input", unversionedAPIs,
		"--clientset-path", clientset,
		"--clientset-name", "internalclientset")
	fmt.Printf("%s\n", strings.Join(c.Args, " "))
	out, err = c.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to run client-gen for unversioned APIs %s %v\n", out, err)
		os.Exit(-1)
	}

	listerPkg := filepath.Join(clientPkg, "listers_generated")
	c = exec.Command(filepath.Join(root, "lister-gen"),
		"-o", GoSrc,
		"--go-header-file", copyright,
		"-i", src,
		"--output-package", listerPkg,
	)
	fmt.Printf("%s\n", strings.Join(c.Args, " "))
	out, err = c.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to run lister-gen %s %v\n", out, err)
		os.Exit(-1)
	}

	informerPkg := filepath.Join(clientPkg, "informers_generated")
	c = exec.Command(filepath.Join(root, "informer-gen"),
		"-o", GoSrc,
		"--go-header-file", copyright,
		"-i", src,
		"--output-package", informerPkg,
		"--listers-package", listerPkg,
		"--versioned-clientset-package", filepath.Join(clientset, "clientset"),
		"--internal-clientset-package", filepath.Join(clientset, "internalclientset"))
	fmt.Printf("%s\n", strings.Join(c.Args, " "))
	out, err = c.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to run informer-gen %s %v\n", out, err)
		os.Exit(-1)
	}
}
