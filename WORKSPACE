load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "51dc53293afe317d2696d4d6433a4c33feedb7748a9e352072e2ec3c0dafd2c6",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.40.1/rules_go-v0.40.1.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.40.1/rules_go-v0.40.1.zip",
    ],
)

http_archive(
    name = "googleapis",
    sha256 = "9d1a930e767c93c825398b8f8692eca3fe353b9aaadedfbcf1fca2282c85df88",
    strip_prefix = "googleapis-64926d52febbf298cb82a8f472ade4a3969ba922",
    urls = [
        "https://github.com/googleapis/googleapis/archive/64926d52febbf298cb82a8f472ade4a3969ba922.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "727f3e4edd96ea20c29e8c2ca9e8d2af724d8c7778e7923a854b2c80952bc405",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.30.0/bazel-gazelle-v0.30.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.30.0/bazel-gazelle-v0.30.0.tar.gz",
    ],
)

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "b1e80761a8a8243d03ebca8845e9cc1ba6c82ce7c5179ce2b295cd36f7e394bf",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.25.0/rules_docker-v0.25.0.tar.gz"],
)

# rules_proto defines abstract rules for building Protocol Buffers.
http_archive(
    name = "rules_proto",
    sha256 = "dc3fb206a2cb3441b485eb1e423165b231235a1ea9b031b4433cf7bc1fa460dd",
    strip_prefix = "rules_proto-5.3.0-21.7",
    urls = [
        "https://github.com/bazelbuild/rules_proto/archive/refs/tags/5.3.0-21.7.tar.gz",
    ],
)

http_archive(
    name = "com_google_protobuf",
    sha256 = "d0f5f605d0d656007ce6c8b5a82df3037e1d8fe8b121ed42e536f569dec16113",
    strip_prefix = "protobuf-3.14.0",
    urls = [
        "https://mirror.bazel.build/github.com/protocolbuffers/protobuf/archive/v3.14.0.tar.gz",
        "https://github.com/protocolbuffers/protobuf/archive/v3.14.0.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")
load("@googleapis//:repository_rules.bzl", "switched_rules_by_language")
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")
load("//:deps.bzl", "go_dependencies")

rules_proto_dependencies()

rules_proto_toolchains()

switched_rules_by_language(
    name = "com_google_googleapis_imports",
)

go_repository(
    name = "com_github_cpuguy83_go_md2man_v2",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/cpuguy83/go-md2man/v2",
    sum = "h1:p1EgwI/C7NhT0JmVkwCD2ZBK8j4aeHQX2pMHHBfMQ6w=",
    version = "v2.0.2",
)

go_repository(
    name = "com_github_russross_blackfriday_v2",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/russross/blackfriday/v2",
    sum = "h1:JIOH55/0cWyOuilr9/qlrm0BSXldqnqwMsf35Ld67mk=",
    version = "v2.1.0",
)

go_repository(
    name = "com_github_xrash_smetrics",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/xrash/smetrics",
    sum = "h1:bAn7/zixMGCfxrRTfdpNzjtPYqr8smhKouy9mxVdGPU=",
    version = "v0.0.0-20201216005158-039620a65673",
)

go_repository(
    name = "com_google_cloud_go_accessapproval",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/accessapproval",
    sum = "h1:/5YjNhR6lzCvmJZAnByYkfEgWjfAKwYP6nkuTk6nKFE=",
    version = "v1.7.1",
)

go_repository(
    name = "com_google_cloud_go_accesscontextmanager",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/accesscontextmanager",
    sum = "h1:WIAt9lW9AXtqw/bnvrEUaE8VG/7bAAeMzRCBGMkc4+w=",
    version = "v1.8.1",
)

go_repository(
    name = "com_google_cloud_go_aiplatform",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/aiplatform",
    sum = "h1:M5davZWCTzE043rJCn+ZLW6hSxfG1KAx4vJTtas2/ec=",
    version = "v1.48.0",
)

go_repository(
    name = "com_google_cloud_go_analytics",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/analytics",
    sum = "h1:TFBC1ZAqX9/jL56GEXdLrVe5vT3I22bDVWyDwZX4IEg=",
    version = "v0.21.3",
)

go_repository(
    name = "com_google_cloud_go_apigateway",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/apigateway",
    sum = "h1:aBSwCQPcp9rZ0zVEUeJbR623palnqtvxJlUyvzsKGQc=",
    version = "v1.6.1",
)

go_repository(
    name = "com_google_cloud_go_apigeeconnect",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/apigeeconnect",
    sum = "h1:6u/jj0P2c3Mcm+H9qLsXI7gYcTiG9ueyQL3n6vCmFJM=",
    version = "v1.6.1",
)

go_repository(
    name = "com_google_cloud_go_appengine",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/appengine",
    sum = "h1:J+aaUZ6IbTpBegXbmEsh8qZZy864ZVnOoWyfa1XSNbI=",
    version = "v1.8.1",
)

go_repository(
    name = "com_google_cloud_go_area120",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/area120",
    sum = "h1:wiOq3KDpdqXmaHzvZwKdpoM+3lDcqsI2Lwhyac7stss=",
    version = "v0.8.1",
)

go_repository(
    name = "com_google_cloud_go_artifactregistry",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/artifactregistry",
    sum = "h1:k6hNqab2CubhWlGcSzunJ7kfxC7UzpAfQ1UPb9PDCKI=",
    version = "v1.14.1",
)

go_repository(
    name = "com_google_cloud_go_asset",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/asset",
    sum = "h1:vlHdznX70eYW4V1y1PxocvF6tEwxJTTarwIGwOhFF3U=",
    version = "v1.14.1",
)

go_repository(
    name = "com_google_cloud_go_assuredworkloads",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/assuredworkloads",
    sum = "h1:yaO0kwS+SnhVSTF7BqTyVGt3DTocI6Jqo+S3hHmCwNk=",
    version = "v1.11.1",
)

go_repository(
    name = "com_google_cloud_go_automl",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/automl",
    sum = "h1:iP9iQurb0qbz+YOOMfKSEjhONA/WcoOIjt6/m+6pIgo=",
    version = "v1.13.1",
)

go_repository(
    name = "com_google_cloud_go_baremetalsolution",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/baremetalsolution",
    sum = "h1:0Ge9PQAy6cZ1tRrkc44UVgYV15nw2TVnzJzYsMHXF+E=",
    version = "v1.1.1",
)

go_repository(
    name = "com_google_cloud_go_batch",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/batch",
    sum = "h1:uE0Q//W7FOGPjf7nuPiP0zoE8wOT3ngoIO2HIet0ilY=",
    version = "v1.3.1",
)

go_repository(
    name = "com_google_cloud_go_beyondcorp",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/beyondcorp",
    sum = "h1:VPg+fZXULQjs8LiMeWdLaB5oe8G9sEoZ0I0j6IMiG1Q=",
    version = "v1.0.0",
)

go_repository(
    name = "com_google_cloud_go_bigquery",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/bigquery",
    sum = "h1:K3wLbjbnSlxhuG5q4pntHv5AEbQM1QqHKGYgwFIqOTg=",
    version = "v1.53.0",
)

go_repository(
    name = "com_google_cloud_go_billing",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/billing",
    sum = "h1:1iktEAIZ2uA6KpebC235zi/rCXDdDYQ0bTXTNetSL80=",
    version = "v1.16.0",
)

go_repository(
    name = "com_google_cloud_go_binaryauthorization",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/binaryauthorization",
    sum = "h1:cAkOhf1ic92zEN4U1zRoSupTmwmxHfklcp1X7CCBKvE=",
    version = "v1.6.1",
)

go_repository(
    name = "com_google_cloud_go_certificatemanager",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/certificatemanager",
    sum = "h1:uKsohpE0hiobx1Eak9jNcPCznwfB6gvyQCcS28Ah9E8=",
    version = "v1.7.1",
)

go_repository(
    name = "com_google_cloud_go_channel",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/channel",
    sum = "h1:dqRkK2k7Ll/HHeYGxv18RrfhozNxuTJRkspW0iaFZoY=",
    version = "v1.16.0",
)

go_repository(
    name = "com_google_cloud_go_cloudbuild",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/cloudbuild",
    sum = "h1:YBbAWcvE4x6xPWTyS+OU4eiUpz5rCS3VCM/aqmfddPA=",
    version = "v1.13.0",
)

go_repository(
    name = "com_google_cloud_go_clouddms",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/clouddms",
    sum = "h1:rjR1nV6oVf2aNNB7B5uz1PDIlBjlOiBgR+q5n7bbB7M=",
    version = "v1.6.1",
)

go_repository(
    name = "com_google_cloud_go_cloudtasks",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/cloudtasks",
    sum = "h1:cMh9Q6dkvh+Ry5LAPbD/U2aw6KAqdiU6FttwhbTo69w=",
    version = "v1.12.1",
)

go_repository(
    name = "com_google_cloud_go_contactcenterinsights",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/contactcenterinsights",
    sum = "h1:YR2aPedGVQPpFBZXJnPkqRj8M//8veIZZH5ZvICoXnI=",
    version = "v1.10.0",
)

go_repository(
    name = "com_google_cloud_go_container",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/container",
    sum = "h1:N51t/cgQJFqDD/W7Mb+IvmAPHrf8AbPx7Bb7aF4lROE=",
    version = "v1.24.0",
)

go_repository(
    name = "com_google_cloud_go_containeranalysis",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/containeranalysis",
    sum = "h1:SM/ibWHWp4TYyJMwrILtcBtYKObyupwOVeceI9pNblw=",
    version = "v0.10.1",
)

go_repository(
    name = "com_google_cloud_go_datacatalog",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/datacatalog",
    sum = "h1:qVeQcw1Cz93/cGu2E7TYUPh8Lz5dn5Ws2siIuQ17Vng=",
    version = "v1.16.0",
)

go_repository(
    name = "com_google_cloud_go_dataflow",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/dataflow",
    sum = "h1:VzG2tqsk/HbmOtq/XSfdF4cBvUWRK+S+oL9k4eWkENQ=",
    version = "v0.9.1",
)

go_repository(
    name = "com_google_cloud_go_dataform",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/dataform",
    sum = "h1:xcWso0hKOoxeW72AjBSIp/UfkvpqHNzzS0/oygHlcqY=",
    version = "v0.8.1",
)

go_repository(
    name = "com_google_cloud_go_datafusion",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/datafusion",
    sum = "h1:eX9CZoyhKQW6g1Xj7+RONeDj1mV8KQDKEB9KLELX9/8=",
    version = "v1.7.1",
)

go_repository(
    name = "com_google_cloud_go_datalabeling",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/datalabeling",
    sum = "h1:zxsCD/BLKXhNuRssen8lVXChUj8VxF3ofN06JfdWOXw=",
    version = "v0.8.1",
)

go_repository(
    name = "com_google_cloud_go_dataplex",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/dataplex",
    sum = "h1:yoBWuuUZklYp7nx26evIhzq8+i/nvKYuZr1jka9EqLs=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_dataqna",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/dataqna",
    sum = "h1:ITpUJep04hC9V7C+gcK390HO++xesQFSUJ7S4nSnF3U=",
    version = "v0.8.1",
)

go_repository(
    name = "com_google_cloud_go_datastream",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/datastream",
    sum = "h1:ra/+jMv36zTAGPfi8TRne1hXme+UsKtdcK4j6bnqQiw=",
    version = "v1.10.0",
)

go_repository(
    name = "com_google_cloud_go_deploy",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/deploy",
    sum = "h1:A+w/xpWgz99EYzB6e31gMGAI/P5jTZ2UO7veQK5jQ8o=",
    version = "v1.13.0",
)

go_repository(
    name = "com_google_cloud_go_dialogflow",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/dialogflow",
    sum = "h1:yAQXMYywsorTTXNmWxwb7WZCNkdOPJER8O8GQHSqigs=",
    version = "v1.42.0",
)

go_repository(
    name = "com_google_cloud_go_dlp",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/dlp",
    sum = "h1:tF3wsJ2QulRhRLWPzWVkeDz3FkOGVoMl6cmDUHtfYxw=",
    version = "v1.10.1",
)

go_repository(
    name = "com_google_cloud_go_documentai",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/documentai",
    sum = "h1:dW8ex9yb3oT9s1yD2+yLcU8Zq15AquRZ+wd0U+TkxFw=",
    version = "v1.22.0",
)

go_repository(
    name = "com_google_cloud_go_domains",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/domains",
    sum = "h1:rqz6KY7mEg7Zs/69U6m6LMbB7PxFDWmT3QWNXIqhHm0=",
    version = "v0.9.1",
)

go_repository(
    name = "com_google_cloud_go_edgecontainer",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/edgecontainer",
    sum = "h1:zhHWnLzg6AqzE+I3gzJqiIwHfjEBhWctNQEzqb+FaRo=",
    version = "v1.1.1",
)

go_repository(
    name = "com_google_cloud_go_essentialcontacts",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/essentialcontacts",
    sum = "h1:OEJ0MLXXCW/tX1fkxzEZOsv/wRfyFsvDVNaHWBAvoV0=",
    version = "v1.6.2",
)

go_repository(
    name = "com_google_cloud_go_eventarc",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/eventarc",
    sum = "h1:xIP3XZi0Xawx8DEfh++mE2lrIi5kQmCr/KcWhJ1q0J4=",
    version = "v1.13.0",
)

go_repository(
    name = "com_google_cloud_go_filestore",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/filestore",
    sum = "h1:Eiz8xZzMJc5ppBWkuaod/PUdUZGCFR8ku0uS+Ah2fRw=",
    version = "v1.7.1",
)

go_repository(
    name = "com_google_cloud_go_functions",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/functions",
    sum = "h1:LtAyqvO1TFmNLcROzHZhV0agEJfBi+zfMZsF4RT/a7U=",
    version = "v1.15.1",
)

go_repository(
    name = "com_google_cloud_go_gkebackup",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/gkebackup",
    sum = "h1:lgyrpdhtJKV7l1GM15YFt+OCyHMxsQZuSydyNmS0Pxo=",
    version = "v1.3.0",
)

go_repository(
    name = "com_google_cloud_go_gkeconnect",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/gkeconnect",
    sum = "h1:a1ckRvVznnuvDWESM2zZDzSVFvggeBaVY5+BVB8tbT0=",
    version = "v0.8.1",
)

go_repository(
    name = "com_google_cloud_go_gkehub",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/gkehub",
    sum = "h1:2BLSb8i+Co1P05IYCKATXy5yaaIw/ZqGvVSBTLdzCQo=",
    version = "v0.14.1",
)

go_repository(
    name = "com_google_cloud_go_gkemulticloud",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/gkemulticloud",
    sum = "h1:MluqhtPVZReoriP5+adGIw+ij/RIeRik8KApCW2WMTw=",
    version = "v1.0.0",
)

go_repository(
    name = "com_google_cloud_go_gsuiteaddons",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/gsuiteaddons",
    sum = "h1:mi9jxZpzVjLQibTS/XfPZvl+Jr6D5Bs8pGqUjllRb00=",
    version = "v1.6.1",
)

go_repository(
    name = "com_google_cloud_go_iam",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/iam",
    sum = "h1:lW7fzj15aVIXYHREOqjRBV9PsH0Z6u8Y46a1YGvQP4Y=",
    version = "v1.1.1",
)

go_repository(
    name = "com_google_cloud_go_iap",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/iap",
    sum = "h1:X1tcp+EoJ/LGX6cUPt3W2D4H2Kbqq0pLAsldnsCjLlE=",
    version = "v1.8.1",
)

go_repository(
    name = "com_google_cloud_go_ids",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/ids",
    sum = "h1:khXYmSoDDhWGEVxHl4c4IgbwSRR+qE/L4hzP3vaU9Hc=",
    version = "v1.4.1",
)

go_repository(
    name = "com_google_cloud_go_iot",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/iot",
    sum = "h1:yrH0OSmicD5bqGBoMlWG8UltzdLkYzNUwNVUVz7OT54=",
    version = "v1.7.1",
)

go_repository(
    name = "com_google_cloud_go_kms",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/kms",
    sum = "h1:xYl5WEaSekKYN5gGRyhjvZKM22GVBBCzegGNVPy+aIs=",
    version = "v1.15.0",
)

go_repository(
    name = "com_google_cloud_go_language",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/language",
    sum = "h1:3MXeGEv8AlX+O2LyV4pO4NGpodanc26AmXwOuipEym0=",
    version = "v1.10.1",
)

go_repository(
    name = "com_google_cloud_go_lifesciences",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/lifesciences",
    sum = "h1:axkANGx1wiBXHiPcJZAE+TDjjYoJRIDzbHC/WYllCBU=",
    version = "v0.9.1",
)

go_repository(
    name = "com_google_cloud_go_longrunning",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/longrunning",
    sum = "h1:Fr7TXftcqTudoyRJa113hyaqlGdiBQkp0Gq7tErFDWI=",
    version = "v0.5.1",
)

go_repository(
    name = "com_google_cloud_go_managedidentities",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/managedidentities",
    sum = "h1:2/qZuOeLgUHorSdxSQGtnOu9xQkBn37+j+oZQv/KHJY=",
    version = "v1.6.1",
)

go_repository(
    name = "com_google_cloud_go_mediatranslation",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/mediatranslation",
    sum = "h1:50cF7c1l3BanfKrpnTCaTvhf+Fo6kdF21DG0byG7gYU=",
    version = "v0.8.1",
)

go_repository(
    name = "com_google_cloud_go_memcache",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/memcache",
    sum = "h1:7lkLsF0QF+Mre0O/NvkD9Q5utUNwtzvIYjrOLOs0HO0=",
    version = "v1.10.1",
)

go_repository(
    name = "com_google_cloud_go_metastore",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/metastore",
    sum = "h1:+9DsxUOHvsqvC0ylrRc/JwzbXJaaBpfIK3tX0Lx8Tcc=",
    version = "v1.12.0",
)

go_repository(
    name = "com_google_cloud_go_monitoring",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/monitoring",
    sum = "h1:65JhLMd+JiYnXr6j5Z63dUYCuOg770p8a/VC+gil/58=",
    version = "v1.15.1",
)

go_repository(
    name = "com_google_cloud_go_networkconnectivity",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/networkconnectivity",
    sum = "h1:LnrYM6lBEeTq+9f2lR4DjBhv31EROSAQi/P5W4Q0AEc=",
    version = "v1.12.1",
)

go_repository(
    name = "com_google_cloud_go_networkmanagement",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/networkmanagement",
    sum = "h1:/3xP37eMxnyvkfLrsm1nv1b2FbMMSAEAOlECTvoeCq4=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_networksecurity",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/networksecurity",
    sum = "h1:TBLEkMp3AE+6IV/wbIGRNTxnqLXHCTEQWoxRVC18TzY=",
    version = "v0.9.1",
)

go_repository(
    name = "com_google_cloud_go_notebooks",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/notebooks",
    sum = "h1:CUqMNEtv4EHFnbogV+yGHQH5iAQLmijOx191innpOcs=",
    version = "v1.9.1",
)

go_repository(
    name = "com_google_cloud_go_optimization",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/optimization",
    sum = "h1:pEwOAmO00mxdbesCRSsfj8Sd4rKY9kBrYW7Vd3Pq7cA=",
    version = "v1.4.1",
)

go_repository(
    name = "com_google_cloud_go_orchestration",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/orchestration",
    sum = "h1:KmN18kE/xa1n91cM5jhCh7s1/UfIguSCisw7nTMUzgE=",
    version = "v1.8.1",
)

go_repository(
    name = "com_google_cloud_go_orgpolicy",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/orgpolicy",
    sum = "h1:I/7dHICQkNwym9erHqmlb50LRU588NPCvkfIY0Bx9jI=",
    version = "v1.11.1",
)

go_repository(
    name = "com_google_cloud_go_osconfig",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/osconfig",
    sum = "h1:dgyEHdfqML6cUW6/MkihNdTVc0INQst0qSE8Ou1ub9c=",
    version = "v1.12.1",
)

go_repository(
    name = "com_google_cloud_go_oslogin",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/oslogin",
    sum = "h1:LdSuG3xBYu2Sgr3jTUULL1XCl5QBx6xwzGqzoDUw1j0=",
    version = "v1.10.1",
)

go_repository(
    name = "com_google_cloud_go_phishingprotection",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/phishingprotection",
    sum = "h1:aK/lNmSd1vtbft/vLe2g7edXK72sIQbqr2QyrZN/iME=",
    version = "v0.8.1",
)

go_repository(
    name = "com_google_cloud_go_policytroubleshooter",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/policytroubleshooter",
    sum = "h1:XTMHy31yFmXgQg57CB3w9YQX8US7irxDX0Fl0VwlZyY=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_privatecatalog",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/privatecatalog",
    sum = "h1:B/18xGo+E0EMS9LOEQ0zXz7F2asMgmVgTYGSI89MHOA=",
    version = "v0.9.1",
)

go_repository(
    name = "com_google_cloud_go_recaptchaenterprise_v2",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/recaptchaenterprise/v2",
    sum = "h1:IGkbudobsTXAwmkEYOzPCQPApUCsN4Gbq3ndGVhHQpI=",
    version = "v2.7.2",
)

go_repository(
    name = "com_google_cloud_go_recommendationengine",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/recommendationengine",
    sum = "h1:nMr1OEVHuDambRn+/y4RmNAmnR/pXCuHtH0Y4tCgGRQ=",
    version = "v0.8.1",
)

go_repository(
    name = "com_google_cloud_go_recommender",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/recommender",
    sum = "h1:UKp94UH5/Lv2WXSQe9+FttqV07x/2p1hFTMMYVFtilg=",
    version = "v1.10.1",
)

go_repository(
    name = "com_google_cloud_go_redis",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/redis",
    sum = "h1:YrjQnCC7ydk+k30op7DSjSHw1yAYhqYXFcOq1bSXRYA=",
    version = "v1.13.1",
)

go_repository(
    name = "com_google_cloud_go_resourcemanager",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/resourcemanager",
    sum = "h1:QIAMfndPOHR6yTmMUB0ZN+HSeRmPjR/21Smq5/xwghI=",
    version = "v1.9.1",
)

go_repository(
    name = "com_google_cloud_go_resourcesettings",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/resourcesettings",
    sum = "h1:Fdyq418U69LhvNPFdlEO29w+DRRjwDA4/pFamm4ksAg=",
    version = "v1.6.1",
)

go_repository(
    name = "com_google_cloud_go_retail",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/retail",
    sum = "h1:gYBrb9u/Hc5s5lUTFXX1Vsbc/9BEvgtioY6ZKaK0DK8=",
    version = "v1.14.1",
)

go_repository(
    name = "com_google_cloud_go_run",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/run",
    sum = "h1:kHeIG8q+N6Zv0nDkBjSOYfK2eWqa5FnaiDPH/7/HirE=",
    version = "v1.2.0",
)

go_repository(
    name = "com_google_cloud_go_scheduler",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/scheduler",
    sum = "h1:yoZbZR8880KgPGLmACOMCiY2tPk+iX4V/dkxqTirlz8=",
    version = "v1.10.1",
)

go_repository(
    name = "com_google_cloud_go_secretmanager",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/secretmanager",
    sum = "h1:cLTCwAjFh9fKvU6F13Y4L9vPcx9yiWPyWXE4+zkuEQs=",
    version = "v1.11.1",
)

go_repository(
    name = "com_google_cloud_go_security",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/security",
    sum = "h1:jR3itwycg/TgGA0uIgTItcVhA55hKWiNJxaNNpQJaZE=",
    version = "v1.15.1",
)

go_repository(
    name = "com_google_cloud_go_securitycenter",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/securitycenter",
    sum = "h1:XOGJ9OpnDtqg8izd7gYk/XUhj8ytjIalyjjsR6oyG0M=",
    version = "v1.23.0",
)

go_repository(
    name = "com_google_cloud_go_servicedirectory",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/servicedirectory",
    sum = "h1:pBWpjCFVGWkzVTkqN3TBBIqNSoSHY86/6RL0soSQ4z8=",
    version = "v1.11.0",
)

go_repository(
    name = "com_google_cloud_go_shell",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/shell",
    sum = "h1:aHbwH9LSqs4r2rbay9f6fKEls61TAjT63jSyglsw7sI=",
    version = "v1.7.1",
)

go_repository(
    name = "com_google_cloud_go_speech",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/speech",
    sum = "h1:MCagaq8ObV2tr1kZJcJYgXYbIn8Ai5rp42tyGYw9rls=",
    version = "v1.19.0",
)

go_repository(
    name = "com_google_cloud_go_storagetransfer",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/storagetransfer",
    sum = "h1:+ZLkeXx0K0Pk5XdDmG0MnUVqIR18lllsihU/yq39I8Q=",
    version = "v1.10.0",
)

go_repository(
    name = "com_google_cloud_go_talent",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/talent",
    sum = "h1:j46ZgD6N2YdpFPux9mc7OAf4YK3tiBCsbLKc8rQx+bU=",
    version = "v1.6.2",
)

go_repository(
    name = "com_google_cloud_go_texttospeech",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/texttospeech",
    sum = "h1:S/pR/GZT9p15R7Y2dk2OXD/3AufTct/NSxT4a7nxByw=",
    version = "v1.7.1",
)

go_repository(
    name = "com_google_cloud_go_tpu",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/tpu",
    sum = "h1:kQf1jgPY04UJBYYjNUO+3GrZtIb57MfGAW2bwgLbR3A=",
    version = "v1.6.1",
)

go_repository(
    name = "com_google_cloud_go_trace",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/trace",
    sum = "h1:EwGdOLCNfYOOPtgqo+D2sDLZmRCEO1AagRTJCU6ztdg=",
    version = "v1.10.1",
)

go_repository(
    name = "com_google_cloud_go_translate",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/translate",
    sum = "h1:PQHamiOzlehqLBJMnM72lXk/OsMQewZB12BKJ8zXrU0=",
    version = "v1.8.2",
)

go_repository(
    name = "com_google_cloud_go_video",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/video",
    sum = "h1:BRyyS+wU+Do6VOXnb8WfPr42ZXti9hzmLKLUCkggeK4=",
    version = "v1.19.0",
)

go_repository(
    name = "com_google_cloud_go_videointelligence",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/videointelligence",
    sum = "h1:MBMWnkQ78GQnRz5lfdTAbBq/8QMCF3wahgtHh3s/J+k=",
    version = "v1.11.1",
)

go_repository(
    name = "com_google_cloud_go_vision_v2",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/vision/v2",
    sum = "h1:ccK6/YgPfGHR/CyESz1mvIbsht5Y2xRsWCPqmTNydEw=",
    version = "v2.7.2",
)

go_repository(
    name = "com_google_cloud_go_vmmigration",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/vmmigration",
    sum = "h1:gnjIclgqbEMc+cF5IJuPxp53wjBIlqZ8h9hE8Rkwp7A=",
    version = "v1.7.1",
)

go_repository(
    name = "com_google_cloud_go_vpcaccess",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/vpcaccess",
    sum = "h1:ram0GzjNWElmbxXMIzeOZUkQ9J8ZAahD6V8ilPGqX0Y=",
    version = "v1.7.1",
)

go_repository(
    name = "com_google_cloud_go_webrisk",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/webrisk",
    sum = "h1:Ssy3MkOMOnyRV5H2bkMQ13Umv7CwB/kugo3qkAX83Fk=",
    version = "v1.9.1",
)

go_repository(
    name = "com_google_cloud_go_websecurityscanner",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/websecurityscanner",
    sum = "h1:CfEF/vZ+xXyAR3zC9iaC/QRdf1MEgS20r5UR17Q4gOg=",
    version = "v1.6.1",
)

go_repository(
    name = "com_google_cloud_go_workflows",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/workflows",
    sum = "h1:2akeQ/PgtRhrNuD/n1WvJd5zb7YyuDZrlOanBj2ihPg=",
    version = "v1.11.1",
)

go_repository(
    name = "com_github_urfave_cli_v2",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/urfave/cli/v2",
    sum = "h1:VAzn5oq403l5pHjc4OhD54+XGO9cdKVL/7lDjF+iKUs=",
    version = "v2.25.7",
)

go_repository(
    name = "com_github_aws_aws_lambda_go",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/aws/aws-lambda-go",
    sum = "h1:l/5fyVb6Ud9uYd411xdHZzSf2n86TakxzpvIoz7l+3Y=",
    version = "v1.41.0",
)

go_repository(
    name = "com_github_cenkalti_backoff_v4",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/cenkalti/backoff/v4",
    sum = "h1:y4OZtCnogmCPw98Zjyt5a6+QwPLGkiQsYW5oUqylYbM=",
    version = "v4.2.1",
)

go_repository(
    name = "com_github_jarcoal_httpmock",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/jarcoal/httpmock",
    sum = "h1:gSvTxxFR/MEMfsGrvRbdfpRUMBStovlSRLw0Ep1bwwc=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_kelseyhightower_envconfig",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/kelseyhightower/envconfig",
    sum = "h1:Im6hONhd3pLkfDFsbRgu68RDNkGF1r3dvMUtDTo2cv8=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_okta_okta_sdk_golang_v2",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/okta/okta-sdk-golang/v2",
    sum = "h1:EDKM+uOPfihOMNwgHMdno+NAsIfyXkVnoFAYVPay0YU=",
    version = "v2.20.0",
)

go_repository(
    name = "com_github_patrickmn_go_cache",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/patrickmn/go-cache",
    sum = "h1:pSCLCl6joCFRnjpeojzOpEYs4q7Vditq8fySFG5ap3Y=",
    version = "v0.0.0-20180815053127-5633e0862627",
)

go_repository(
    name = "com_google_cloud_go_apigeeregistry",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/apigeeregistry",
    sum = "h1:hgq0ANLDx7t2FDZDJQrCMtCtddR/pjCqVuvQWGrQbXw=",
    version = "v0.7.1",
)

go_repository(
    name = "com_google_cloud_go_errorreporting",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/errorreporting",
    sum = "h1:kj1XEWMu8P0qlLhm3FwcaFsUvXChV/OraZwA70trRR0=",
    version = "v0.3.0",
)

go_repository(
    name = "com_google_cloud_go_logging",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/logging",
    sum = "h1:CJYxlNNNNAMkHp9em/YEXcfJg+rPDg7YfwoRpMU+t5I=",
    version = "v1.7.0",
)

go_repository(
    name = "com_google_cloud_go_maps",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/maps",
    sum = "h1:PdfgpBLhAoSzZrQXP+/zBc78fIPLZSJp5y8+qSMn2UU=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_pubsublite",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/pubsublite",
    sum = "h1:pX+idpWMIH30/K7c0epN6V703xpIcMXWRjKJsz0tYGY=",
    version = "v1.8.1",
)

go_repository(
    name = "com_google_cloud_go_spanner",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/spanner",
    sum = "h1:aqiMP8dhsEXgn9K5EZBWxPG7dxIiyM2VaikqeU4iteg=",
    version = "v1.47.0",
)

go_repository(
    name = "com_google_cloud_go_vmwareengine",
    build_file_proto_mode = "disable_global",
    importpath = "cloud.google.com/go/vmwareengine",
    sum = "h1:qsJ0CPlOQu/3MFBGklu752v3AkD+Pdu091UmXJ+EjTA=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_emicklei_go_restful_v3",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/emicklei/go-restful/v3",
    sum = "h1:rAQeMHw1c7zTmncogyy8VvRZwtkmkZ4FxERmMY4rD+g=",
    version = "v3.11.0",
)

go_repository(
    name = "com_github_evanphx_json_patch",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/evanphx/json-patch",
    sum = "h1:jBYDEEiFBPxA0v50tFdvOzQQTCvpL6mnFh5mB2/l16U=",
    version = "v5.6.0+incompatible",
)

go_repository(
    name = "com_github_go_logr_logr",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/go-logr/logr",
    sum = "h1:g01GSCwiDw2xSZfjJ2/T9M+S6pFdcNtFYsp+Y43HYDQ=",
    version = "v1.2.4",
)

go_repository(
    name = "com_github_go_openapi_jsonpointer",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/go-openapi/jsonpointer",
    sum = "h1:ESKJdU9ASRfaPNOPRx12IUyA1vn3R9GiE3KYD14BXdQ=",
    version = "v0.20.0",
)

go_repository(
    name = "com_github_go_openapi_jsonreference",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/go-openapi/jsonreference",
    sum = "h1:3sVjiK66+uXK/6oQ8xgcRKcFgQ5KXa2KvnJRumpMGbE=",
    version = "v0.20.2",
)

go_repository(
    name = "com_github_go_openapi_swag",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/go-openapi/swag",
    sum = "h1:QLMzNJnMGPRNDCbySlcj1x01tzU8/9LTTL9hZZZogBU=",
    version = "v0.22.4",
)

go_repository(
    name = "com_github_google_gofuzz",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/google/gofuzz",
    sum = "h1:xRy4A+RhZaiKjJ1bPfwQ8sedCA+YS2YcCHW6ec7JMi0=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_gregjones_httpcache",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/gregjones/httpcache",
    sum = "h1:pdN6V1QBWetyv/0+wjACpqVH+eVULgEjkurDLq3goeM=",
    version = "v0.0.0-20180305231024-9cad4c3443a7",
)

go_repository(
    name = "com_github_imdario_mergo",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/imdario/mergo",
    sum = "h1:xTNEAn+kxVO7dTZGu0CegyqKZmoWFI0rF8UxjlB2d28=",
    version = "v0.3.6",
)

go_repository(
    name = "com_github_josharian_intern",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/josharian/intern",
    sum = "h1:vlS4z54oSdjm0bgjRigI+G1HpF+tI+9rE5LLzOg8HmY=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_mailru_easyjson",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/mailru/easyjson",
    sum = "h1:UGYAvKxe3sBsEDzO8ZeWOSlIQfWFlxbzLZe7hwFURr0=",
    version = "v0.7.7",
)

go_repository(
    name = "com_github_moby_spdystream",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/moby/spdystream",
    sum = "h1:cjW1zVyyoiM0T7b6UoySUFqzXMoqRckQtXwGPiBhOM8=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_munnerz_goautoneg",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/munnerz/goautoneg",
    sum = "h1:C3w9PqII01/Oq1c1nUAm88MOHcQC9l5mIlSMApZMrHA=",
    version = "v0.0.0-20191010083416-a7dc8b61c822",
)

go_repository(
    name = "com_github_peterbourgon_diskv",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/peterbourgon/diskv",
    sum = "h1:UBdAOUP5p4RWqPBg048CAvpKN+vxiaj6gdUUzhl4XmI=",
    version = "v2.0.1+incompatible",
)

go_repository(
    name = "in_gopkg_inf_v0",
    build_file_proto_mode = "disable_global",
    importpath = "gopkg.in/inf.v0",
    sum = "h1:73M5CoZyi3ZLMOyDlQh031Cx6N9NDJ2Vvfl76EDAgDc=",
    version = "v0.9.1",
)

go_repository(
    name = "io_k8s_api",
    build_file_proto_mode = "disable_global",
    importpath = "k8s.io/api",
    sum = "h1:i+0O8k2NPBCPYaMB+uCkseEbawEt/eFaiRqUx8aB108=",
    version = "v0.28.1",
)

go_repository(
    name = "io_k8s_apimachinery",
    build_file_proto_mode = "disable_global",
    importpath = "k8s.io/apimachinery",
    sum = "h1:EJD40og3GizBSV3mkIoXQBsws32okPOy+MkRyzh6nPY=",
    version = "v0.28.1",
)

go_repository(
    name = "io_k8s_client_go",
    build_file_proto_mode = "disable_global",
    importpath = "k8s.io/client-go",
    sum = "h1:pRhMzB8HyLfVwpngWKE8hDcXRqifh1ga2Z/PU9SXVK8=",
    version = "v0.28.1",
)

go_repository(
    name = "io_k8s_klog_v2",
    build_file_proto_mode = "disable_global",
    importpath = "k8s.io/klog/v2",
    sum = "h1:7WCHKK6K8fNhTqfBhISHQ97KrnJNFZMcQvKp7gP/tmg=",
    version = "v2.100.1",
)

go_repository(
    name = "io_k8s_kube_openapi",
    build_file_proto_mode = "disable_global",
    importpath = "k8s.io/kube-openapi",
    sum = "h1:CAIciCnJnSOQxPd0xvpV6JU3D4AJvnYbImPpFpO9Hnw=",
    version = "v0.0.0-20230816210353-14e408962443",
)

go_repository(
    name = "io_k8s_sigs_json",
    build_file_proto_mode = "disable_global",
    importpath = "sigs.k8s.io/json",
    sum = "h1:EDPBXCAspyGV4jQlpZSudPeMmr1bNJefnuqLsRAsHZo=",
    version = "v0.0.0-20221116044647-bc3834ca7abd",
)

go_repository(
    name = "io_k8s_sigs_structured_merge_diff_v4",
    build_file_proto_mode = "disable_global",
    importpath = "sigs.k8s.io/structured-merge-diff/v4",
    sum = "h1:UZbZAZfX0wV2zr7YZorDz6GXROfDFj6LvqCRm4VUVKk=",
    version = "v4.3.0",
)

go_repository(
    name = "io_k8s_sigs_yaml",
    build_file_proto_mode = "disable_global",
    importpath = "sigs.k8s.io/yaml",
    sum = "h1:a2VclLzOGrwOHDiV8EfBGhvjHvP46CtW5j6POvhYGGo=",
    version = "v1.3.0",
)

go_repository(
    name = "io_k8s_utils",
    build_file_proto_mode = "disable_global",
    importpath = "k8s.io/utils",
    sum = "h1:sgn3ZU783SCgtaSJjpcVVlRqd6GSnlTLKgpAAttJvpI=",
    version = "v0.0.0-20230726121419-3b25d923346b",
)

go_repository(
    name = "com_github_armon_go_socks5",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/armon/go-socks5",
    sum = "h1:0CwZNZbxp69SHPdPJAN/hZIm0C4OItdklCFmMRWYpio=",
    version = "v0.0.0-20160902184237-e75332964ef5",
)

go_repository(
    name = "com_github_mxk_go_flowrate",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/mxk/go-flowrate",
    sum = "h1:y5//uYreIhSUg3J1GEMiLbxo1LJaP8RfCpH6pymGZus=",
    version = "v0.0.0-20140419014527-cca7078d478f",
)

go_repository(
    name = "com_github_onsi_ginkgo_v2",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/onsi/ginkgo/v2",
    sum = "h1:xR7vG4IXt5RWx6FfIjyAtsoMAtnc3C/rFXBBd2AjZwE=",
    version = "v2.9.4",
)

go_repository(
    name = "com_github_onsi_gomega",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/onsi/gomega",
    sum = "h1:ENqfyGeS5AX/rlXDd/ETokDz93u0YufY1Pgxuy/PvWE=",
    version = "v1.27.6",
)

go_repository(
    name = "com_github_asaskevich_govalidator",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/asaskevich/govalidator",
    sum = "h1:idn718Q4B6AGu/h5Sxe66HYVdqdGu2l9Iebqhi/AEoA=",
    version = "v0.0.0-20190424111038-f61b66f89f4a",
)

go_repository(
    name = "com_github_creack_pty",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/creack/pty",
    sum = "h1:uDmaGzcdjhF4i/plgjmEsriH11Y0o7RKapEf/LDaM3w=",
    version = "v1.1.9",
)

go_repository(
    name = "com_github_kisielk_errcheck",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/kisielk/errcheck",
    sum = "h1:e8esj/e4R+SAOwFwN+n3zr0nYeCyeweozKfO23MvHzY=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_nytimes_gziphandler",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/NYTimes/gziphandler",
    sum = "h1:lsxEuwrXEAokXB9qhlbKWPpo3KMLZQ5WB5WLQRW1uq0=",
    version = "v0.0.0-20170623195520-56545f4a5d46",
)

go_repository(
    name = "io_k8s_gengo",
    build_file_proto_mode = "disable_global",
    importpath = "k8s.io/gengo",
    sum = "h1:GohjlNKauSai7gN4wsJkeZ3WAJx4Sh+oT/b5IYn5suA=",
    version = "v0.0.0-20210813121822-485abfe95c7c",
)

go_repository(
    name = "com_github_antihax_optional",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/antihax/optional",
    sum = "h1:xK2lYat7ZLaVVcIuj82J8kIro4V6kDe0AUDFboUCwcg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_ghodss_yaml",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/ghodss/yaml",
    sum = "h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_grpc_ecosystem_grpc_gateway",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/grpc-ecosystem/grpc-gateway",
    sum = "h1:gmcG1KaJ57LophUzW0Hy8NmPhnMZb4M0+kPpLofRdBo=",
    version = "v1.16.0",
)

go_repository(
    name = "com_github_rogpeppe_fastuuid",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/rogpeppe/fastuuid",
    sum = "h1:Ppwyp6VYCF1nvBTXL3trRso7mXMlRrw9ooo375wvi2s=",
    version = "v1.2.0",
)

go_repository(
    name = "io_opentelemetry_go_proto_otlp",
    build_file_proto_mode = "disable_global",
    importpath = "go.opentelemetry.io/proto/otlp",
    sum = "h1:rwOQPCuKAKmwGKq2aVNnYIibI6wnV7EvzgfTCzcdGg8=",
    version = "v0.7.0",
)

go_repository(
    name = "com_github_go_test_deep",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/go-test/deep",
    sum = "h1:u2CU3YKy9I2pmu9pX0eq50wCgjfGIt539SqR7FbHiho=",
    version = "v1.0.4",
)

go_repository(
    name = "com_github_gorilla_websocket",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/gorilla/websocket",
    sum = "h1:PPwGk2jz7EePpoHN/+ClbZu8SPxiqlu12wZP/3sWmnc=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_slack_go_slack",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/slack-go/slack",
    sum = "h1:92/dfFU8Q5XP6Wp5rr5/T5JHLM5c5Smtn53fhToAP88=",
    version = "v0.12.3",
)

go_repository(
    name = "com_github_sirupsen_logrus",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/sirupsen/logrus",
    sum = "h1:dueUQJ1C2q9oE3F7wvmSGAaVtTmUizReu6fjN8uqzbQ=",
    version = "v1.9.3",
)

go_repository(
    name = "com_github_g3n_engine",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/g3n/engine",
    sum = "h1:7dmj4c+3xHcBnYrVmRuVf/oZ2JycxJU9Y+2FQj1Af2Y=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_golang_freetype",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/golang/freetype",
    sum = "h1:DACJavvAHhabrF08vX0COfcOBJRhZ8lUbR+ZWIs0Y5g=",
    version = "v0.0.0-20170609003504-e2365dfdc4a0",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/aws/aws-sdk-go-v2",
    sum = "h1:gMT0IW+03wtYJhRqTVYn0wLzwdnK9sRMcxmtfGzRdJc=",
    version = "v1.21.0",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_config",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/aws/aws-sdk-go-v2/config",
    sum = "h1:RNAfbPqw1CstCooHaTPhScz7z1PyocQj0UL+l95CgzI=",
    version = "v1.18.37",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_credentials",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/aws/aws-sdk-go-v2/credentials",
    sum = "h1:QpsNitYJu0GgvMBLUIYu9H4yryA5kMksjeIVQfgXrt8=",
    version = "v1.13.35",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_feature_ec2_imds",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/aws/aws-sdk-go-v2/feature/ec2/imds",
    sum = "h1:uDZJF1hu0EVT/4bogChk8DyjSF6fof6uL/0Y26Ma7Fg=",
    version = "v1.13.11",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_internal_configsources",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/aws/aws-sdk-go-v2/internal/configsources",
    sum = "h1:22dGT7PneFMx4+b3pz7lMTRyN8ZKH7M2cW4GP9yUS2g=",
    version = "v1.1.41",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_internal_endpoints_v2",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/aws/aws-sdk-go-v2/internal/endpoints/v2",
    sum = "h1:SijA0mgjV8E+8G45ltVHs0fvKpTj8xmZJ3VwhGKtUSI=",
    version = "v2.4.35",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_service_internal_presigned_url",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/aws/aws-sdk-go-v2/service/internal/presigned-url",
    sum = "h1:CdzPW9kKitgIiLV1+MHobfR5Xg25iYnyzWZhyQuSlDI=",
    version = "v1.9.35",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_service_sfn",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/aws/aws-sdk-go-v2/service/sfn",
    sum = "h1:uGuCRiB/3dCGb0iInxJJJTeMvTxM7wIdFv9R0uSFLKQ=",
    version = "v1.19.5",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_service_sso",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/aws/aws-sdk-go-v2/service/sso",
    sum = "h1:oCvTFSDi67AX0pOX3PuPdGFewvLRU2zzFSrTsgURNo0=",
    version = "v1.13.5",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_service_ssooidc",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/aws/aws-sdk-go-v2/service/ssooidc",
    sum = "h1:dnInJb4S0oy8aQuri1mV6ipLlnZPfnsDNB9BGO9PDNY=",
    version = "v1.15.5",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_service_sts",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/aws/aws-sdk-go-v2/service/sts",
    sum = "h1:CQBFElb0LS8RojMJlxRSo/HXipvTZW2S44Lt9Mk2aYQ=",
    version = "v1.21.5",
)

go_repository(
    name = "com_github_aws_smithy_go",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/aws/smithy-go",
    sum = "h1:MJU9hqBGbvWZdApzpvoF2WAIJDbtjK2NDJSiJP7HblQ=",
    version = "v1.14.2",
)

go_repository(
    name = "com_github_billheroinc_logrus",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/BillHeroInc/logrus",
    sum = "h1:NELeUSqRVUBgqkZJS9oEnfMmtFo0a6imfjRhEP0ELK4=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_jmespath_go_jmespath",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/jmespath/go-jmespath",
    sum = "h1:BEgLn5cpjn8UN1mAw4NjwDrS35OdebyEtFe+9YPoQUg=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_jmespath_go_jmespath_internal_testify",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/jmespath/go-jmespath/internal/testify",
    sum = "h1:shLQSRRSCCPj3f2gpwzGwWFoC7ycTf1rcQZHOlsJ6N8=",
    version = "v1.5.1",
)

go_repository(
    name = "com_github_charmbracelet_bubbletea",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/charmbracelet/bubbletea",
    sum = "h1:uaQIKx9Ai6Gdh5zpTbGiWpytMU+CfsPp06RaW2cx/SY=",
    version = "v0.24.2",
)

go_repository(
    name = "com_github_containerd_console",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/containerd/console",
    sum = "h1:q2hJAaP1k2wIvVRd/hEHD7lacgqrCPS+k8g1MndzfWY=",
    version = "v1.0.4-0.20230313162750-1ae8d489ac81",
)

go_repository(
    name = "com_github_lucasb_eyer_go_colorful",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/lucasb-eyer/go-colorful",
    sum = "h1:1nnpGOrhyZZuNyfu1QjKiUICQ74+3FNCN69Aj6K7nkY=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_mattn_go_localereader",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/mattn/go-localereader",
    sum = "h1:ygSAOl7ZXTx4RdPYinUpg6W99U8jWvWi9Ye2JC/oIi4=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_mattn_go_runewidth",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/mattn/go-runewidth",
    sum = "h1:UNAjwbU9l54TA3KzvqLGxwWjHmMgBUVhBiTjelZgg3U=",
    version = "v0.0.15",
)

go_repository(
    name = "com_github_muesli_ansi",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/muesli/ansi",
    sum = "h1:ZK8zHtRHOkbHy6Mmr5D264iyp3TiX5OmNcI5cIARiQI=",
    version = "v0.0.0-20230316100256-276c6243b2f6",
)

go_repository(
    name = "com_github_muesli_cancelreader",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/muesli/cancelreader",
    sum = "h1:3I4Kt4BQjOR54NavqnDogx/MIoWBFa0StPA8ELUXHmA=",
    version = "v0.2.2",
)

go_repository(
    name = "com_github_muesli_reflow",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/muesli/reflow",
    sum = "h1:IFsN6K9NfGtjeggFP+68I4chLZV2yIKsXJFNZ+eWh6s=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_muesli_termenv",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/muesli/termenv",
    sum = "h1:GohcuySI0QmI3wN8Ok9PtKGkgkFIk7y6Vpb5PvrY+Wo=",
    version = "v0.15.2",
)

go_repository(
    name = "com_github_rivo_uniseg",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/rivo/uniseg",
    sum = "h1:8TfxU8dW6PdqD27gjM8MVNuicgxIjxpm4K7x4jp8sis=",
    version = "v0.4.4",
)

go_repository(
    name = "com_github_golang_snappy",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/golang/snappy",
    sum = "h1:yAGX7huGHXlcLOEtBnF4w7FQwA26wojNCwOYAEhLjQM=",
    version = "v0.0.4",
)

go_repository(
    name = "com_github_kylelemons_godebug",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/kylelemons/godebug",
    sum = "h1:RPNrshWIDI6G2gRW9EHilWtl7Z6Sb1BR0xunSBf0SNc=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_sashabaranov_go_openai",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/sashabaranov/go-openai",
    sum = "h1:5DPTtR9JBjKPJS008/A409I5ntFhUPPGCmaAihcPRyo=",
    version = "v1.14.2",
)

go_repository(
    name = "com_github_atotto_clipboard",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/atotto/clipboard",
    sum = "h1:EH0zSVneZPSuFR11BlR9YppQTVDbh5+16AmcJi4g1z4=",
    version = "v0.1.4",
)

go_repository(
    name = "com_github_aymanbagabas_go_osc52_v2",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/aymanbagabas/go-osc52/v2",
    sum = "h1:HwpRHbFMcZLEVr42D4p7XBqjyuxQH5SMiErDT4WkJ2k=",
    version = "v2.0.1",
)

go_repository(
    name = "com_github_charmbracelet_bubbles",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/charmbracelet/bubbles",
    sum = "h1:6uzpAAaT9ZqKssntbvZMlksWHruQLNxg49H5WdeuYSY=",
    version = "v0.16.1",
)

go_repository(
    name = "com_github_charmbracelet_harmonica",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/charmbracelet/harmonica",
    sum = "h1:8NxJWRWg/bzKqqEaaeFNipOu77YR5t8aSwG4pgaUBiQ=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_charmbracelet_lipgloss",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/charmbracelet/lipgloss",
    sum = "h1:IS00fk4XAHcf8uZKc3eHeMUTCxUH6NkaTrdyCQk84RU=",
    version = "v0.8.0",
)

go_repository(
    name = "com_github_dustin_go_humanize",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/dustin/go-humanize",
    sum = "h1:GzkhY7T5VNhEkwH0PVJgjz+fX1rhBrR7pRT3mDkpeCY=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_sahilm_fuzzy",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/sahilm/fuzzy",
    sum = "h1:FzWGaw2Opqyu+794ZQ9SYifWv2EIXpwP4q8dY1kDAwI=",
    version = "v0.1.0",
)

# gazelle:repository_macro deps.bzl%go_dependencies
go_dependencies()

go_rules_dependencies()

go_register_toolchains(version = "1.20.7")

gazelle_dependencies()

protobuf_deps()

## Container configuration setup
# Load the macro that allows you to customize the docker toolchain configuration.
load(
    "@io_bazel_rules_docker//toolchains/docker:toolchain.bzl",
    docker_toolchain_configure = "toolchain_configure",
)

docker_toolchain_configure(
    name = "docker_config",
    # Replace this with a Bazel label to the config.json file. Note absolute or relative
    # paths are not supported. Docker allows you to specify custom authentication credentials
    # in the client configuration JSON file.
    # See https://docs.docker.com/engine/reference/commandline/cli/#configuration-files
    # for more details.
    client_config = "@//.docker:config.json",
)

load("@io_bazel_rules_docker//repositories:repositories.bzl", container_repositories = "repositories")

container_repositories()

load("@io_bazel_rules_docker//go:image.bzl", _go_image_repos = "repositories")

_go_image_repos()

load("@io_bazel_rules_docker//repositories:deps.bzl", container_deps = "deps")

container_deps()

load("@io_bazel_rules_docker//container:container.bzl", "container_pull")

container_pull(
    name = "static-debian11",
    # 'tag' is also supported, but digest is encouraged for reproducibility.
    digest = "sha256:58b04fe33b80e96132fe3d28d7f08d5792ebf28a346da13376fcddcab3a9410e",
    registry = "gcr.io",
    repository = "distroless/static-debian11",
)

container_pull(
    name = "golang_amd64",
    digest = "sha256:f2e0acaf7c628cd819b73541d7c1ea8f888d51edb0a58935a3c46a084fa953fa",
    registry = "index.docker.io",
    repository = "library/golang",
)
## End of Container configuration section
