package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
    "go_test",
)

go_library(
    name = "go_default_library",
    srcs = ["strategy.go"],
    importpath = "k8s.io/kubernetes/pkg/security/podsecuritypolicy/apparmor",
    deps = [
        "//vendor/k8s.io/apimachinery/pkg/util/validation/field:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/api:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/security/apparmor:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/util/maps:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["strategy_test.go"],
    importpath = "k8s.io/kubernetes/pkg/security/podsecuritypolicy/apparmor",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/davecgh/go-spew/spew:go_default_library",
        "//vendor/github.com/stretchr/testify/assert:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/api:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/security/apparmor:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/util/maps:go_default_library",
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
    srcs = [":package-srcs"],
    tags = ["automanaged"],
)
