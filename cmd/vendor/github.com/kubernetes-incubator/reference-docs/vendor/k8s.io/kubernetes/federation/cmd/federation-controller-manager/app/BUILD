package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "go_default_library",
    srcs = [
        "controllermanager.go",
        "plugins.go",
    ],
    importpath = "k8s.io/kubernetes/federation/cmd/federation-controller-manager/app",
    deps = [
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/prometheus/client_golang/prometheus:go_default_library",
        "//vendor/github.com/spf13/cobra:go_default_library",
        "//vendor/github.com/spf13/pflag:go_default_library",
        "//vendor/k8s.io/api/core/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/api/errors:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime/schema:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/util/wait:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/server/healthz:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/util/flag:go_default_library",
        "//vendor/k8s.io/client-go/discovery:go_default_library",
        "//vendor/k8s.io/client-go/kubernetes:go_default_library",
        "//vendor/k8s.io/client-go/rest:go_default_library",
        "//vendor/k8s.io/client-go/tools/clientcmd:go_default_library",
        "//vendor/k8s.io/client-go/tools/leaderelection:go_default_library",
        "//vendor/k8s.io/client-go/tools/leaderelection/resourcelock:go_default_library",
        "//vendor/k8s.io/client-go/tools/record:go_default_library",
        "//vendor/k8s.io/kubernetes/federation/apis/federation/v1beta1:go_default_library",
        "//vendor/k8s.io/kubernetes/federation/client/clientset_generated/federation_clientset:go_default_library",
        "//vendor/k8s.io/kubernetes/federation/cmd/federation-controller-manager/app/options:go_default_library",
        "//vendor/k8s.io/kubernetes/federation/pkg/dnsprovider/providers/aws/route53:go_default_library",
        "//vendor/k8s.io/kubernetes/federation/pkg/dnsprovider/providers/coredns:go_default_library",
        "//vendor/k8s.io/kubernetes/federation/pkg/dnsprovider/providers/google/clouddns:go_default_library",
        "//vendor/k8s.io/kubernetes/federation/pkg/federatedtypes:go_default_library",
        "//vendor/k8s.io/kubernetes/federation/pkg/federation-controller/cluster:go_default_library",
        "//vendor/k8s.io/kubernetes/federation/pkg/federation-controller/ingress:go_default_library",
        "//vendor/k8s.io/kubernetes/federation/pkg/federation-controller/job:go_default_library",
        "//vendor/k8s.io/kubernetes/federation/pkg/federation-controller/service:go_default_library",
        "//vendor/k8s.io/kubernetes/federation/pkg/federation-controller/service/dns:go_default_library",
        "//vendor/k8s.io/kubernetes/federation/pkg/federation-controller/sync:go_default_library",
        "//vendor/k8s.io/kubernetes/federation/pkg/federation-controller/util/eventsink:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/api:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/util/configz:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/version:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["controllermanager_test.go"],
    importpath = "k8s.io/kubernetes/federation/cmd/federation-controller-manager/app",
    library = ":go_default_library",
    deps = [
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime/schema:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/util/flag:go_default_library",
        "//vendor/k8s.io/kubernetes/federation/pkg/federation-controller/ingress:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//federation/cmd/federation-controller-manager/app/options:all-srcs",
    ],
    tags = ["automanaged"],
)
