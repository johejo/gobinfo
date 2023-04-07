# gobinfo

Print build infomation embedded in a Go binary file.

## Install

```
go install github.com/johejo/gobinfo@latest
```

## Usage

```
gobinfo -file /path/to/bin
# or
gobinfo < /path/to/bin
```

## Examples

### gobinfo itself

```
$ gobinfo -file ~/go/bin/gobinfo
go      go1.20.3
path    github.com/johejo/gobinfo
mod     github.com/johejo/gobinfo       (devel)
build   -buildmode=exe
build   -compiler=gc
build   CGO_ENABLED=1
build   CGO_CFLAGS=
build   CGO_CPPFLAGS=
build   CGO_CXXFLAGS=
build   CGO_LDFLAGS=
build   GOARCH=amd64
build   GOOS=linux
build   GOAMD64=v1
build   vcs=git
build   vcs.modified=true
```

### kubectl v1.26.3

```
$ curl -fsSL "https://storage.googleapis.com/kubernetes-release/release/v1.26.3/bin/linux/amd64/kubectl" | gobinfo
go	go1.19.7
path	k8s.io/kubernetes/cmd/kubectl
mod	k8s.io/kubernetes	(devel)	
dep	github.com/MakeNowJust/heredoc	v1.0.0	
dep	github.com/beorn7/perks	v1.0.1	
dep	github.com/blang/semver/v4	v4.0.0	
dep	github.com/cespare/xxhash/v2	v2.1.2	
dep	github.com/chai2010/gettext-go	v1.0.2	
dep	github.com/davecgh/go-spew	v1.1.1	
dep	github.com/daviddengcn/go-colortext	v1.0.0	
dep	github.com/docker/distribution	v2.8.1+incompatible	
dep	github.com/emicklei/go-restful/v3	v3.9.0	
dep	github.com/evanphx/json-patch	v4.12.0+incompatible	
dep	github.com/exponent-io/jsonpath	v0.0.0-20151013193312-d6023ce2651d	
dep	github.com/fatih/camelcase	v1.0.0	
dep	github.com/fvbommel/sortorder	v1.0.1	
dep	github.com/go-errors/errors	v1.0.1	
dep	github.com/go-logr/logr	v1.2.3	
dep	github.com/go-openapi/jsonpointer	v0.19.5	
dep	github.com/go-openapi/jsonreference	v0.20.0	
dep	github.com/go-openapi/swag	v0.19.14	
dep	github.com/gogo/protobuf	v1.3.2	
dep	github.com/golang/protobuf	v1.5.2	
dep	github.com/google/btree	v1.0.1	
dep	github.com/google/gnostic	v0.5.7-v3refs	
dep	github.com/google/go-cmp	v0.5.9	
dep	github.com/google/gofuzz	v1.1.0	
dep	github.com/google/shlex	v0.0.0-20191202100458-e7afc7fbc510	
dep	github.com/google/uuid	v1.1.2	
dep	github.com/gregjones/httpcache	v0.0.0-20180305231024-9cad4c3443a7	
dep	github.com/imdario/mergo	v0.3.6	
dep	github.com/jonboulle/clockwork	v0.2.2	
dep	github.com/josharian/intern	v1.0.0	
dep	github.com/json-iterator/go	v1.1.12	
dep	github.com/liggitt/tabwriter	v0.0.0-20181228230101-89fcab3d43de	
dep	github.com/lithammer/dedent	v1.1.0	
dep	github.com/mailru/easyjson	v0.7.6	
dep	github.com/matttproud/golang_protobuf_extensions	v1.0.2	
dep	github.com/mitchellh/go-wordwrap	v1.0.0	
dep	github.com/moby/spdystream	v0.2.0	
dep	github.com/moby/term	v0.0.0-20220808134915-39b0c02b01ae	
dep	github.com/modern-go/concurrent	v0.0.0-20180306012644-bacd9c7ef1dd	
dep	github.com/modern-go/reflect2	v1.0.2	
dep	github.com/monochromegane/go-gitignore	v0.0.0-20200626010858-205db1a8cc00	
dep	github.com/munnerz/goautoneg	v0.0.0-20191010083416-a7dc8b61c822	
dep	github.com/mxk/go-flowrate	v0.0.0-20140419014527-cca7078d478f	
dep	github.com/opencontainers/go-digest	v1.0.0	
dep	github.com/peterbourgon/diskv	v2.0.1+incompatible	
dep	github.com/pkg/errors	v0.9.1	
dep	github.com/prometheus/client_golang	v1.14.0	
dep	github.com/prometheus/client_model	v0.3.0	
dep	github.com/prometheus/common	v0.37.0	
dep	github.com/prometheus/procfs	v0.8.0	
dep	github.com/russross/blackfriday/v2	v2.1.0	
dep	github.com/spf13/cobra	v1.6.0	
dep	github.com/spf13/pflag	v1.0.5	
dep	github.com/xlab/treeprint	v1.1.0	
dep	go.starlark.net	v0.0.0-20200306205701-8dd3e2ee1dd5	
dep	golang.org/x/net	v0.7.0	
dep	golang.org/x/oauth2	v0.0.0-20220223155221-ee480838109b	
dep	golang.org/x/sys	v0.5.0	
dep	golang.org/x/term	v0.5.0	
dep	golang.org/x/text	v0.7.0	
dep	golang.org/x/time	v0.0.0-20220210224613-90d013bbcef8	
dep	google.golang.org/protobuf	v1.28.1	
dep	gopkg.in/inf.v0	v0.9.1	
dep	gopkg.in/yaml.v2	v2.4.0	
dep	gopkg.in/yaml.v3	v3.0.1	
dep	k8s.io/api	v0.0.0
=>	./staging/src/k8s.io/api	(devel)	

dep	k8s.io/apimachinery	v0.0.0
=>	./staging/src/k8s.io/apimachinery	(devel)	

dep	k8s.io/cli-runtime	v0.0.0
=>	./staging/src/k8s.io/cli-runtime	(devel)	

dep	k8s.io/client-go	v0.0.0
=>	./staging/src/k8s.io/client-go	(devel)	

dep	k8s.io/component-base	v0.0.0
=>	./staging/src/k8s.io/component-base	(devel)	

dep	k8s.io/component-helpers	v0.0.0
=>	./staging/src/k8s.io/component-helpers	(devel)	

dep	k8s.io/klog/v2	v2.80.1	
dep	k8s.io/kube-openapi	v0.0.0-20221012153701-172d655c2280	
dep	k8s.io/kubectl	v0.0.0
=>	./staging/src/k8s.io/kubectl	(devel)	

dep	k8s.io/metrics	v0.0.0
=>	./staging/src/k8s.io/metrics	(devel)	

dep	k8s.io/utils	v0.0.0-20221107191617-1a15be271d1d	
dep	sigs.k8s.io/json	v0.0.0-20220713155537-f223a00ba0e2	
dep	sigs.k8s.io/kustomize/api	v0.12.1	
dep	sigs.k8s.io/kustomize/kustomize/v4	v4.5.7	
dep	sigs.k8s.io/kustomize/kyaml	v0.13.9	
dep	sigs.k8s.io/structured-merge-diff/v4	v4.2.3	
dep	sigs.k8s.io/yaml	v1.3.0	
build	-asmflags=all=-trimpath=/workspace/src/k8s.io/kubernetes/_output/dockerized/go/src/k8s.io/kubernetes
build	-compiler=gc
build	-gcflags="all=-trimpath=/workspace/src/k8s.io/kubernetes/_output/dockerized/go/src/k8s.io/kubernetes "
build	-ldflags="all=-X 'k8s.io/kubernetes/vendor/k8s.io/client-go/pkg/version.buildDate=2023-03-15T13:40:17Z' -X 'k8s.io/kubernetes/vendor/k8s.io/component-base/version.buildDate=2023-03-15T13:40:17Z' -X 'k8s.io/client-go/pkg/version.buildDate=2023-03-15T13:40:17Z' -X 'k8s.io/component-base/version.buildDate=2023-03-15T13:40:17Z' -X 'k8s.io/kubernetes/vendor/k8s.io/client-go/pkg/version.gitCommit=9e644106593f3f4aa98f8a84b23db5fa378900bd' -X 'k8s.io/kubernetes/vendor/k8s.io/component-base/version.gitCommit=9e644106593f3f4aa98f8a84b23db5fa378900bd' -X 'k8s.io/client-go/pkg/version.gitCommit=9e644106593f3f4aa98f8a84b23db5fa378900bd' -X 'k8s.io/component-base/version.gitCommit=9e644106593f3f4aa98f8a84b23db5fa378900bd' -X 'k8s.io/kubernetes/vendor/k8s.io/client-go/pkg/version.gitTreeState=clean' -X 'k8s.io/kubernetes/vendor/k8s.io/component-base/version.gitTreeState=clean' -X 'k8s.io/client-go/pkg/version.gitTreeState=clean' -X 'k8s.io/component-base/version.gitTreeState=clean' -X 'k8s.io/kubernetes/vendor/k8s.io/client-go/pkg/version.gitVersion=v1.26.3' -X 'k8s.io/kubernetes/vendor/k8s.io/component-base/version.gitVersion=v1.26.3' -X 'k8s.io/client-go/pkg/version.gitVersion=v1.26.3' -X 'k8s.io/component-base/version.gitVersion=v1.26.3' -X 'k8s.io/kubernetes/vendor/k8s.io/client-go/pkg/version.gitMajor=1' -X 'k8s.io/kubernetes/vendor/k8s.io/component-base/version.gitMajor=1' -X 'k8s.io/client-go/pkg/version.gitMajor=1' -X 'k8s.io/component-base/version.gitMajor=1' -X 'k8s.io/kubernetes/vendor/k8s.io/client-go/pkg/version.gitMinor=26' -X 'k8s.io/kubernetes/vendor/k8s.io/component-base/version.gitMinor=26' -X 'k8s.io/client-go/pkg/version.gitMinor=26' -X 'k8s.io/component-base/version.gitMinor=26'  -s -w"
build	-tags=selinux,notest
build	CGO_ENABLED=0
build	GOARCH=amd64
build	GOOS=linux
build	GOAMD64=v1
build	vcs=git
build	vcs.revision=9e644106593f3f4aa98f8a84b23db5fa378900bd
build	vcs.time=2023-03-15T13:33:11Z
build	vcs.modified=false
```

## License

MIT
