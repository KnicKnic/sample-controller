// This is a generated file. Do not edit directly.

module k8s.io/sample-controller

go 1.13

require (
	github.com/KnicKnic/go-powershell v0.0.10
	github.com/flosch/pongo2 v0.0.0-20190707114632-bbf5a6c351f4 // indirect
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/juju/go4 v0.0.0-20160222163258-40d72ab9641a // indirect
	github.com/juju/persistent-cookiejar v0.0.0-20171026135701-d5e5a8405ef9 // indirect
	github.com/juju/webbrowser v1.0.0 // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	github.com/lxc/lxd v0.0.0-20200115020223-dd8970dfc7af
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rogpeppe/fastuuid v1.2.0 // indirect
	github.com/scylladb/go-set v1.0.2
	github.com/yyyar/gobetween v0.0.0-00010101000000-000000000000
	gopkg.in/httprequest.v1 v1.2.0 // indirect
	gopkg.in/juju/environschema.v1 v1.0.0 // indirect
	gopkg.in/macaroon-bakery.v2 v2.1.0 // indirect
	gopkg.in/macaroon.v2 v2.1.0 // indirect
	gopkg.in/retry.v1 v1.0.3 // indirect
	gopkg.in/robfig/cron.v2 v2.0.0-20150107220207-be2e0b0deed5 // indirect
	k8s.io/api v0.0.0-20200113233642-3946df5ca773
	k8s.io/apimachinery v0.0.0-20200113233504-44bd77c24ef9
	k8s.io/client-go v0.0.0-20200113233857-bcaa73156d59
	k8s.io/code-generator v0.0.0-20200113233325-0826954c61ed
	k8s.io/klog v1.0.0
)

replace (
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a // pinned to release-branch.go1.13
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190821162956-65e3620a7ae7 // pinned to release-branch.go1.13
	k8s.io/api => k8s.io/api v0.0.0-20200113233642-3946df5ca773
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20200113233504-44bd77c24ef9
	k8s.io/client-go => k8s.io/client-go v0.0.0-20200113233857-bcaa73156d59
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20200113233325-0826954c61ed
)

replace github.com/yyyar/gobetween => github.com/KnicKnic/gobetween/src v0.0.0-20200115081610-fd37f831237b
