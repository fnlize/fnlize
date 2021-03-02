# Fission CRD generation

* Clone https://github.com/fission/code-generator to generate fission CRD object deepcopy and client methods.
* MUST run code-generator in the fission root directory.

``` bash
$ cd $GOPATH/src/github.com/fnlize/fnlize/
$ bash $GOPATH/src/k8s.io/code-generator/generate-groups.sh \
    all \
    github.com/fnlize/fnlize/pkg/apis/genclient \
    github.com/fnlize/fnlize/pkg/apis \
    "core:v1" \
    --go-header-file $GOPATH/src/github.com/fnlize/fnlize/pkg/apis/boilerplate.txt
```

# Reference

* https://blog.openshift.com/kubernetes-deep-dive-code-generation-customresources/
