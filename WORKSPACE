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
    importpath = "github.com/cpuguy83/go-md2man/v2",
    sum = "h1:p1EgwI/C7NhT0JmVkwCD2ZBK8j4aeHQX2pMHHBfMQ6w=",
    version = "v2.0.2",
)

go_repository(
    name = "com_github_russross_blackfriday_v2",
    importpath = "github.com/russross/blackfriday/v2",
    sum = "h1:JIOH55/0cWyOuilr9/qlrm0BSXldqnqwMsf35Ld67mk=",
    version = "v2.1.0",
)

go_repository(
    name = "com_github_xrash_smetrics",
    importpath = "github.com/xrash/smetrics",
    sum = "h1:bAn7/zixMGCfxrRTfdpNzjtPYqr8smhKouy9mxVdGPU=",
    version = "v0.0.0-20201216005158-039620a65673",
)

go_repository(
    name = "com_google_cloud_go_accessapproval",
    importpath = "cloud.google.com/go/accessapproval",
    sum = "h1:/5YjNhR6lzCvmJZAnByYkfEgWjfAKwYP6nkuTk6nKFE=",
    version = "v1.7.1",
)

go_repository(
    name = "com_google_cloud_go_accesscontextmanager",
    importpath = "cloud.google.com/go/accesscontextmanager",
    sum = "h1:WIAt9lW9AXtqw/bnvrEUaE8VG/7bAAeMzRCBGMkc4+w=",
    version = "v1.8.1",
)

go_repository(
    name = "com_google_cloud_go_aiplatform",
    importpath = "cloud.google.com/go/aiplatform",
    sum = "h1:M5davZWCTzE043rJCn+ZLW6hSxfG1KAx4vJTtas2/ec=",
    version = "v1.48.0",
)

go_repository(
    name = "com_google_cloud_go_analytics",
    importpath = "cloud.google.com/go/analytics",
    sum = "h1:TFBC1ZAqX9/jL56GEXdLrVe5vT3I22bDVWyDwZX4IEg=",
    version = "v0.21.3",
)

go_repository(
    name = "com_google_cloud_go_apigateway",
    importpath = "cloud.google.com/go/apigateway",
    sum = "h1:aBSwCQPcp9rZ0zVEUeJbR623palnqtvxJlUyvzsKGQc=",
    version = "v1.6.1",
)

go_repository(
    name = "com_google_cloud_go_apigeeconnect",
    importpath = "cloud.google.com/go/apigeeconnect",
    sum = "h1:6u/jj0P2c3Mcm+H9qLsXI7gYcTiG9ueyQL3n6vCmFJM=",
    version = "v1.6.1",
)

go_repository(
    name = "com_google_cloud_go_appengine",
    importpath = "cloud.google.com/go/appengine",
    sum = "h1:J+aaUZ6IbTpBegXbmEsh8qZZy864ZVnOoWyfa1XSNbI=",
    version = "v1.8.1",
)

go_repository(
    name = "com_google_cloud_go_area120",
    importpath = "cloud.google.com/go/area120",
    sum = "h1:wiOq3KDpdqXmaHzvZwKdpoM+3lDcqsI2Lwhyac7stss=",
    version = "v0.8.1",
)

go_repository(
    name = "com_google_cloud_go_artifactregistry",
    importpath = "cloud.google.com/go/artifactregistry",
    sum = "h1:k6hNqab2CubhWlGcSzunJ7kfxC7UzpAfQ1UPb9PDCKI=",
    version = "v1.14.1",
)

go_repository(
    name = "com_google_cloud_go_asset",
    importpath = "cloud.google.com/go/asset",
    sum = "h1:vlHdznX70eYW4V1y1PxocvF6tEwxJTTarwIGwOhFF3U=",
    version = "v1.14.1",
)

go_repository(
    name = "com_google_cloud_go_assuredworkloads",
    importpath = "cloud.google.com/go/assuredworkloads",
    sum = "h1:yaO0kwS+SnhVSTF7BqTyVGt3DTocI6Jqo+S3hHmCwNk=",
    version = "v1.11.1",
)

go_repository(
    name = "com_google_cloud_go_automl",
    importpath = "cloud.google.com/go/automl",
    sum = "h1:iP9iQurb0qbz+YOOMfKSEjhONA/WcoOIjt6/m+6pIgo=",
    version = "v1.13.1",
)

go_repository(
    name = "com_google_cloud_go_baremetalsolution",
    importpath = "cloud.google.com/go/baremetalsolution",
    sum = "h1:0Ge9PQAy6cZ1tRrkc44UVgYV15nw2TVnzJzYsMHXF+E=",
    version = "v1.1.1",
)

go_repository(
    name = "com_google_cloud_go_batch",
    importpath = "cloud.google.com/go/batch",
    sum = "h1:uE0Q//W7FOGPjf7nuPiP0zoE8wOT3ngoIO2HIet0ilY=",
    version = "v1.3.1",
)

go_repository(
    name = "com_google_cloud_go_beyondcorp",
    importpath = "cloud.google.com/go/beyondcorp",
    sum = "h1:VPg+fZXULQjs8LiMeWdLaB5oe8G9sEoZ0I0j6IMiG1Q=",
    version = "v1.0.0",
)

go_repository(
    name = "com_google_cloud_go_bigquery",
    importpath = "cloud.google.com/go/bigquery",
    sum = "h1:K3wLbjbnSlxhuG5q4pntHv5AEbQM1QqHKGYgwFIqOTg=",
    version = "v1.53.0",
)

go_repository(
    name = "com_google_cloud_go_billing",
    importpath = "cloud.google.com/go/billing",
    sum = "h1:1iktEAIZ2uA6KpebC235zi/rCXDdDYQ0bTXTNetSL80=",
    version = "v1.16.0",
)

go_repository(
    name = "com_google_cloud_go_binaryauthorization",
    importpath = "cloud.google.com/go/binaryauthorization",
    sum = "h1:cAkOhf1ic92zEN4U1zRoSupTmwmxHfklcp1X7CCBKvE=",
    version = "v1.6.1",
)

go_repository(
    name = "com_google_cloud_go_certificatemanager",
    importpath = "cloud.google.com/go/certificatemanager",
    sum = "h1:uKsohpE0hiobx1Eak9jNcPCznwfB6gvyQCcS28Ah9E8=",
    version = "v1.7.1",
)

go_repository(
    name = "com_google_cloud_go_channel",
    importpath = "cloud.google.com/go/channel",
    sum = "h1:dqRkK2k7Ll/HHeYGxv18RrfhozNxuTJRkspW0iaFZoY=",
    version = "v1.16.0",
)

go_repository(
    name = "com_google_cloud_go_cloudbuild",
    importpath = "cloud.google.com/go/cloudbuild",
    sum = "h1:YBbAWcvE4x6xPWTyS+OU4eiUpz5rCS3VCM/aqmfddPA=",
    version = "v1.13.0",
)

go_repository(
    name = "com_google_cloud_go_clouddms",
    importpath = "cloud.google.com/go/clouddms",
    sum = "h1:rjR1nV6oVf2aNNB7B5uz1PDIlBjlOiBgR+q5n7bbB7M=",
    version = "v1.6.1",
)

go_repository(
    name = "com_google_cloud_go_cloudtasks",
    importpath = "cloud.google.com/go/cloudtasks",
    sum = "h1:cMh9Q6dkvh+Ry5LAPbD/U2aw6KAqdiU6FttwhbTo69w=",
    version = "v1.12.1",
)

go_repository(
    name = "com_google_cloud_go_contactcenterinsights",
    importpath = "cloud.google.com/go/contactcenterinsights",
    sum = "h1:YR2aPedGVQPpFBZXJnPkqRj8M//8veIZZH5ZvICoXnI=",
    version = "v1.10.0",
)

go_repository(
    name = "com_google_cloud_go_container",
    importpath = "cloud.google.com/go/container",
    sum = "h1:N51t/cgQJFqDD/W7Mb+IvmAPHrf8AbPx7Bb7aF4lROE=",
    version = "v1.24.0",
)

go_repository(
    name = "com_google_cloud_go_containeranalysis",
    importpath = "cloud.google.com/go/containeranalysis",
    sum = "h1:SM/ibWHWp4TYyJMwrILtcBtYKObyupwOVeceI9pNblw=",
    version = "v0.10.1",
)

go_repository(
    name = "com_google_cloud_go_datacatalog",
    importpath = "cloud.google.com/go/datacatalog",
    sum = "h1:qVeQcw1Cz93/cGu2E7TYUPh8Lz5dn5Ws2siIuQ17Vng=",
    version = "v1.16.0",
)

go_repository(
    name = "com_google_cloud_go_dataflow",
    importpath = "cloud.google.com/go/dataflow",
    sum = "h1:VzG2tqsk/HbmOtq/XSfdF4cBvUWRK+S+oL9k4eWkENQ=",
    version = "v0.9.1",
)

go_repository(
    name = "com_google_cloud_go_dataform",
    importpath = "cloud.google.com/go/dataform",
    sum = "h1:xcWso0hKOoxeW72AjBSIp/UfkvpqHNzzS0/oygHlcqY=",
    version = "v0.8.1",
)

go_repository(
    name = "com_google_cloud_go_datafusion",
    importpath = "cloud.google.com/go/datafusion",
    sum = "h1:eX9CZoyhKQW6g1Xj7+RONeDj1mV8KQDKEB9KLELX9/8=",
    version = "v1.7.1",
)

go_repository(
    name = "com_google_cloud_go_datalabeling",
    importpath = "cloud.google.com/go/datalabeling",
    sum = "h1:zxsCD/BLKXhNuRssen8lVXChUj8VxF3ofN06JfdWOXw=",
    version = "v0.8.1",
)

go_repository(
    name = "com_google_cloud_go_dataplex",
    importpath = "cloud.google.com/go/dataplex",
    sum = "h1:yoBWuuUZklYp7nx26evIhzq8+i/nvKYuZr1jka9EqLs=",
    version = "v1.9.0",
)

go_repository(
    name = "com_google_cloud_go_dataproc",
    importpath = "cloud.google.com/go/dataproc",
    sum = "h1:W47qHL3W4BPkAIbk4SWmIERwsWBaNnWm0P2sdx3YgGU=",
    version = "v1.12.0",
)

go_repository(
    name = "com_google_cloud_go_dataqna",
    importpath = "cloud.google.com/go/dataqna",
    sum = "h1:ITpUJep04hC9V7C+gcK390HO++xesQFSUJ7S4nSnF3U=",
    version = "v0.8.1",
)

go_repository(
    name = "com_google_cloud_go_datastream",
    importpath = "cloud.google.com/go/datastream",
    sum = "h1:ra/+jMv36zTAGPfi8TRne1hXme+UsKtdcK4j6bnqQiw=",
    version = "v1.10.0",
)

go_repository(
    name = "com_google_cloud_go_deploy",
    importpath = "cloud.google.com/go/deploy",
    sum = "h1:A+w/xpWgz99EYzB6e31gMGAI/P5jTZ2UO7veQK5jQ8o=",
    version = "v1.13.0",
)

go_repository(
    name = "com_google_cloud_go_dialogflow",
    importpath = "cloud.google.com/go/dialogflow",
    sum = "h1:yAQXMYywsorTTXNmWxwb7WZCNkdOPJER8O8GQHSqigs=",
    version = "v1.42.0",
)

go_repository(
    name = "com_google_cloud_go_dlp",
    importpath = "cloud.google.com/go/dlp",
    sum = "h1:tF3wsJ2QulRhRLWPzWVkeDz3FkOGVoMl6cmDUHtfYxw=",
    version = "v1.10.1",
)

go_repository(
    name = "com_google_cloud_go_documentai",
    importpath = "cloud.google.com/go/documentai",
    sum = "h1:dW8ex9yb3oT9s1yD2+yLcU8Zq15AquRZ+wd0U+TkxFw=",
    version = "v1.22.0",
)

go_repository(
    name = "com_google_cloud_go_domains",
    importpath = "cloud.google.com/go/domains",
    sum = "h1:rqz6KY7mEg7Zs/69U6m6LMbB7PxFDWmT3QWNXIqhHm0=",
    version = "v0.9.1",
)

go_repository(
    name = "com_google_cloud_go_edgecontainer",
    importpath = "cloud.google.com/go/edgecontainer",
    sum = "h1:zhHWnLzg6AqzE+I3gzJqiIwHfjEBhWctNQEzqb+FaRo=",
    version = "v1.1.1",
)

go_repository(
    name = "com_google_cloud_go_essentialcontacts",
    importpath = "cloud.google.com/go/essentialcontacts",
    sum = "h1:OEJ0MLXXCW/tX1fkxzEZOsv/wRfyFsvDVNaHWBAvoV0=",
    version = "v1.6.2",
)

go_repository(
    name = "com_google_cloud_go_eventarc",
    importpath = "cloud.google.com/go/eventarc",
    sum = "h1:xIP3XZi0Xawx8DEfh++mE2lrIi5kQmCr/KcWhJ1q0J4=",
    version = "v1.13.0",
)

go_repository(
    name = "com_google_cloud_go_filestore",
    importpath = "cloud.google.com/go/filestore",
    sum = "h1:Eiz8xZzMJc5ppBWkuaod/PUdUZGCFR8ku0uS+Ah2fRw=",
    version = "v1.7.1",
)

go_repository(
    name = "com_google_cloud_go_functions",
    importpath = "cloud.google.com/go/functions",
    sum = "h1:LtAyqvO1TFmNLcROzHZhV0agEJfBi+zfMZsF4RT/a7U=",
    version = "v1.15.1",
)

go_repository(
    name = "com_google_cloud_go_gaming",
    importpath = "cloud.google.com/go/gaming",
    sum = "h1:PKggmegChZulPW8yvtziF8P9UOuVFwbvylbEucTNups=",
    version = "v1.6.0",
)

go_repository(
    name = "com_google_cloud_go_gkebackup",
    importpath = "cloud.google.com/go/gkebackup",
    sum = "h1:lgyrpdhtJKV7l1GM15YFt+OCyHMxsQZuSydyNmS0Pxo=",
    version = "v1.3.0",
)

go_repository(
    name = "com_google_cloud_go_gkeconnect",
    importpath = "cloud.google.com/go/gkeconnect",
    sum = "h1:a1ckRvVznnuvDWESM2zZDzSVFvggeBaVY5+BVB8tbT0=",
    version = "v0.8.1",
)

go_repository(
    name = "com_google_cloud_go_gkehub",
    importpath = "cloud.google.com/go/gkehub",
    sum = "h1:2BLSb8i+Co1P05IYCKATXy5yaaIw/ZqGvVSBTLdzCQo=",
    version = "v0.14.1",
)

go_repository(
    name = "com_google_cloud_go_gkemulticloud",
    importpath = "cloud.google.com/go/gkemulticloud",
    sum = "h1:MluqhtPVZReoriP5+adGIw+ij/RIeRik8KApCW2WMTw=",
    version = "v1.0.0",
)

go_repository(
    name = "com_google_cloud_go_gsuiteaddons",
    importpath = "cloud.google.com/go/gsuiteaddons",
    sum = "h1:mi9jxZpzVjLQibTS/XfPZvl+Jr6D5Bs8pGqUjllRb00=",
    version = "v1.6.1",
)

go_repository(
    name = "com_google_cloud_go_iam",
    importpath = "cloud.google.com/go/iam",
    sum = "h1:lW7fzj15aVIXYHREOqjRBV9PsH0Z6u8Y46a1YGvQP4Y=",
    version = "v1.1.1",
)

go_repository(
    name = "com_google_cloud_go_iap",
    importpath = "cloud.google.com/go/iap",
    sum = "h1:X1tcp+EoJ/LGX6cUPt3W2D4H2Kbqq0pLAsldnsCjLlE=",
    version = "v1.8.1",
)

go_repository(
    name = "com_google_cloud_go_ids",
    importpath = "cloud.google.com/go/ids",
    sum = "h1:khXYmSoDDhWGEVxHl4c4IgbwSRR+qE/L4hzP3vaU9Hc=",
    version = "v1.4.1",
)

go_repository(
    name = "com_google_cloud_go_iot",
    importpath = "cloud.google.com/go/iot",
    sum = "h1:yrH0OSmicD5bqGBoMlWG8UltzdLkYzNUwNVUVz7OT54=",
    version = "v1.7.1",
)

go_repository(
    name = "com_google_cloud_go_kms",
    importpath = "cloud.google.com/go/kms",
    sum = "h1:xYl5WEaSekKYN5gGRyhjvZKM22GVBBCzegGNVPy+aIs=",
    version = "v1.15.0",
)

go_repository(
    name = "com_google_cloud_go_language",
    importpath = "cloud.google.com/go/language",
    sum = "h1:3MXeGEv8AlX+O2LyV4pO4NGpodanc26AmXwOuipEym0=",
    version = "v1.10.1",
)

go_repository(
    name = "com_google_cloud_go_lifesciences",
    importpath = "cloud.google.com/go/lifesciences",
    sum = "h1:axkANGx1wiBXHiPcJZAE+TDjjYoJRIDzbHC/WYllCBU=",
    version = "v0.9.1",
)

go_repository(
    name = "com_google_cloud_go_longrunning",
    importpath = "cloud.google.com/go/longrunning",
    sum = "h1:Fr7TXftcqTudoyRJa113hyaqlGdiBQkp0Gq7tErFDWI=",
    version = "v0.5.1",
)

go_repository(
    name = "com_google_cloud_go_managedidentities",
    importpath = "cloud.google.com/go/managedidentities",
    sum = "h1:2/qZuOeLgUHorSdxSQGtnOu9xQkBn37+j+oZQv/KHJY=",
    version = "v1.6.1",
)

go_repository(
    name = "com_google_cloud_go_mediatranslation",
    importpath = "cloud.google.com/go/mediatranslation",
    sum = "h1:50cF7c1l3BanfKrpnTCaTvhf+Fo6kdF21DG0byG7gYU=",
    version = "v0.8.1",
)

go_repository(
    name = "com_google_cloud_go_memcache",
    importpath = "cloud.google.com/go/memcache",
    sum = "h1:7lkLsF0QF+Mre0O/NvkD9Q5utUNwtzvIYjrOLOs0HO0=",
    version = "v1.10.1",
)

go_repository(
    name = "com_google_cloud_go_metastore",
    importpath = "cloud.google.com/go/metastore",
    sum = "h1:+9DsxUOHvsqvC0ylrRc/JwzbXJaaBpfIK3tX0Lx8Tcc=",
    version = "v1.12.0",
)

go_repository(
    name = "com_google_cloud_go_monitoring",
    importpath = "cloud.google.com/go/monitoring",
    sum = "h1:65JhLMd+JiYnXr6j5Z63dUYCuOg770p8a/VC+gil/58=",
    version = "v1.15.1",
)

go_repository(
    name = "com_google_cloud_go_networkconnectivity",
    importpath = "cloud.google.com/go/networkconnectivity",
    sum = "h1:LnrYM6lBEeTq+9f2lR4DjBhv31EROSAQi/P5W4Q0AEc=",
    version = "v1.12.1",
)

go_repository(
    name = "com_google_cloud_go_networkmanagement",
    importpath = "cloud.google.com/go/networkmanagement",
    sum = "h1:/3xP37eMxnyvkfLrsm1nv1b2FbMMSAEAOlECTvoeCq4=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_networksecurity",
    importpath = "cloud.google.com/go/networksecurity",
    sum = "h1:TBLEkMp3AE+6IV/wbIGRNTxnqLXHCTEQWoxRVC18TzY=",
    version = "v0.9.1",
)

go_repository(
    name = "com_google_cloud_go_notebooks",
    importpath = "cloud.google.com/go/notebooks",
    sum = "h1:CUqMNEtv4EHFnbogV+yGHQH5iAQLmijOx191innpOcs=",
    version = "v1.9.1",
)

go_repository(
    name = "com_google_cloud_go_optimization",
    importpath = "cloud.google.com/go/optimization",
    sum = "h1:pEwOAmO00mxdbesCRSsfj8Sd4rKY9kBrYW7Vd3Pq7cA=",
    version = "v1.4.1",
)

go_repository(
    name = "com_google_cloud_go_orchestration",
    importpath = "cloud.google.com/go/orchestration",
    sum = "h1:KmN18kE/xa1n91cM5jhCh7s1/UfIguSCisw7nTMUzgE=",
    version = "v1.8.1",
)

go_repository(
    name = "com_google_cloud_go_orgpolicy",
    importpath = "cloud.google.com/go/orgpolicy",
    sum = "h1:I/7dHICQkNwym9erHqmlb50LRU588NPCvkfIY0Bx9jI=",
    version = "v1.11.1",
)

go_repository(
    name = "com_google_cloud_go_osconfig",
    importpath = "cloud.google.com/go/osconfig",
    sum = "h1:dgyEHdfqML6cUW6/MkihNdTVc0INQst0qSE8Ou1ub9c=",
    version = "v1.12.1",
)

go_repository(
    name = "com_google_cloud_go_oslogin",
    importpath = "cloud.google.com/go/oslogin",
    sum = "h1:LdSuG3xBYu2Sgr3jTUULL1XCl5QBx6xwzGqzoDUw1j0=",
    version = "v1.10.1",
)

go_repository(
    name = "com_google_cloud_go_phishingprotection",
    importpath = "cloud.google.com/go/phishingprotection",
    sum = "h1:aK/lNmSd1vtbft/vLe2g7edXK72sIQbqr2QyrZN/iME=",
    version = "v0.8.1",
)

go_repository(
    name = "com_google_cloud_go_policytroubleshooter",
    importpath = "cloud.google.com/go/policytroubleshooter",
    sum = "h1:XTMHy31yFmXgQg57CB3w9YQX8US7irxDX0Fl0VwlZyY=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_privatecatalog",
    importpath = "cloud.google.com/go/privatecatalog",
    sum = "h1:B/18xGo+E0EMS9LOEQ0zXz7F2asMgmVgTYGSI89MHOA=",
    version = "v0.9.1",
)

go_repository(
    name = "com_google_cloud_go_recaptchaenterprise_v2",
    importpath = "cloud.google.com/go/recaptchaenterprise/v2",
    sum = "h1:IGkbudobsTXAwmkEYOzPCQPApUCsN4Gbq3ndGVhHQpI=",
    version = "v2.7.2",
)

go_repository(
    name = "com_google_cloud_go_recommendationengine",
    importpath = "cloud.google.com/go/recommendationengine",
    sum = "h1:nMr1OEVHuDambRn+/y4RmNAmnR/pXCuHtH0Y4tCgGRQ=",
    version = "v0.8.1",
)

go_repository(
    name = "com_google_cloud_go_recommender",
    importpath = "cloud.google.com/go/recommender",
    sum = "h1:UKp94UH5/Lv2WXSQe9+FttqV07x/2p1hFTMMYVFtilg=",
    version = "v1.10.1",
)

go_repository(
    name = "com_google_cloud_go_redis",
    importpath = "cloud.google.com/go/redis",
    sum = "h1:YrjQnCC7ydk+k30op7DSjSHw1yAYhqYXFcOq1bSXRYA=",
    version = "v1.13.1",
)

go_repository(
    name = "com_google_cloud_go_resourcemanager",
    importpath = "cloud.google.com/go/resourcemanager",
    sum = "h1:QIAMfndPOHR6yTmMUB0ZN+HSeRmPjR/21Smq5/xwghI=",
    version = "v1.9.1",
)

go_repository(
    name = "com_google_cloud_go_resourcesettings",
    importpath = "cloud.google.com/go/resourcesettings",
    sum = "h1:Fdyq418U69LhvNPFdlEO29w+DRRjwDA4/pFamm4ksAg=",
    version = "v1.6.1",
)

go_repository(
    name = "com_google_cloud_go_retail",
    importpath = "cloud.google.com/go/retail",
    sum = "h1:gYBrb9u/Hc5s5lUTFXX1Vsbc/9BEvgtioY6ZKaK0DK8=",
    version = "v1.14.1",
)

go_repository(
    name = "com_google_cloud_go_run",
    importpath = "cloud.google.com/go/run",
    sum = "h1:kHeIG8q+N6Zv0nDkBjSOYfK2eWqa5FnaiDPH/7/HirE=",
    version = "v1.2.0",
)

go_repository(
    name = "com_google_cloud_go_scheduler",
    importpath = "cloud.google.com/go/scheduler",
    sum = "h1:yoZbZR8880KgPGLmACOMCiY2tPk+iX4V/dkxqTirlz8=",
    version = "v1.10.1",
)

go_repository(
    name = "com_google_cloud_go_secretmanager",
    importpath = "cloud.google.com/go/secretmanager",
    sum = "h1:cLTCwAjFh9fKvU6F13Y4L9vPcx9yiWPyWXE4+zkuEQs=",
    version = "v1.11.1",
)

go_repository(
    name = "com_google_cloud_go_security",
    importpath = "cloud.google.com/go/security",
    sum = "h1:jR3itwycg/TgGA0uIgTItcVhA55hKWiNJxaNNpQJaZE=",
    version = "v1.15.1",
)

go_repository(
    name = "com_google_cloud_go_securitycenter",
    importpath = "cloud.google.com/go/securitycenter",
    sum = "h1:XOGJ9OpnDtqg8izd7gYk/XUhj8ytjIalyjjsR6oyG0M=",
    version = "v1.23.0",
)

go_repository(
    name = "com_google_cloud_go_servicecontrol",
    importpath = "cloud.google.com/go/servicecontrol",
    sum = "h1:d0uV7Qegtfaa7Z2ClDzr9HJmnbJW7jn0WhZ7wOX6hLE=",
    version = "v1.11.1",
)

go_repository(
    name = "com_google_cloud_go_servicedirectory",
    importpath = "cloud.google.com/go/servicedirectory",
    sum = "h1:pBWpjCFVGWkzVTkqN3TBBIqNSoSHY86/6RL0soSQ4z8=",
    version = "v1.11.0",
)

go_repository(
    name = "com_google_cloud_go_servicemanagement",
    importpath = "cloud.google.com/go/servicemanagement",
    sum = "h1:fopAQI/IAzlxnVeiKn/8WiV6zKndjFkvi+gzu+NjywY=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_serviceusage",
    importpath = "cloud.google.com/go/serviceusage",
    sum = "h1:rXyq+0+RSIm3HFypctp7WoXxIA563rn206CfMWdqXX4=",
    version = "v1.6.0",
)

go_repository(
    name = "com_google_cloud_go_shell",
    importpath = "cloud.google.com/go/shell",
    sum = "h1:aHbwH9LSqs4r2rbay9f6fKEls61TAjT63jSyglsw7sI=",
    version = "v1.7.1",
)

go_repository(
    name = "com_google_cloud_go_speech",
    importpath = "cloud.google.com/go/speech",
    sum = "h1:MCagaq8ObV2tr1kZJcJYgXYbIn8Ai5rp42tyGYw9rls=",
    version = "v1.19.0",
)

go_repository(
    name = "com_google_cloud_go_storagetransfer",
    importpath = "cloud.google.com/go/storagetransfer",
    sum = "h1:+ZLkeXx0K0Pk5XdDmG0MnUVqIR18lllsihU/yq39I8Q=",
    version = "v1.10.0",
)

go_repository(
    name = "com_google_cloud_go_talent",
    importpath = "cloud.google.com/go/talent",
    sum = "h1:j46ZgD6N2YdpFPux9mc7OAf4YK3tiBCsbLKc8rQx+bU=",
    version = "v1.6.2",
)

go_repository(
    name = "com_google_cloud_go_texttospeech",
    importpath = "cloud.google.com/go/texttospeech",
    sum = "h1:S/pR/GZT9p15R7Y2dk2OXD/3AufTct/NSxT4a7nxByw=",
    version = "v1.7.1",
)

go_repository(
    name = "com_google_cloud_go_tpu",
    importpath = "cloud.google.com/go/tpu",
    sum = "h1:kQf1jgPY04UJBYYjNUO+3GrZtIb57MfGAW2bwgLbR3A=",
    version = "v1.6.1",
)

go_repository(
    name = "com_google_cloud_go_trace",
    importpath = "cloud.google.com/go/trace",
    sum = "h1:EwGdOLCNfYOOPtgqo+D2sDLZmRCEO1AagRTJCU6ztdg=",
    version = "v1.10.1",
)

go_repository(
    name = "com_google_cloud_go_translate",
    importpath = "cloud.google.com/go/translate",
    sum = "h1:PQHamiOzlehqLBJMnM72lXk/OsMQewZB12BKJ8zXrU0=",
    version = "v1.8.2",
)

go_repository(
    name = "com_google_cloud_go_video",
    importpath = "cloud.google.com/go/video",
    sum = "h1:BRyyS+wU+Do6VOXnb8WfPr42ZXti9hzmLKLUCkggeK4=",
    version = "v1.19.0",
)

go_repository(
    name = "com_google_cloud_go_videointelligence",
    importpath = "cloud.google.com/go/videointelligence",
    sum = "h1:MBMWnkQ78GQnRz5lfdTAbBq/8QMCF3wahgtHh3s/J+k=",
    version = "v1.11.1",
)

go_repository(
    name = "com_google_cloud_go_vision_v2",
    importpath = "cloud.google.com/go/vision/v2",
    sum = "h1:ccK6/YgPfGHR/CyESz1mvIbsht5Y2xRsWCPqmTNydEw=",
    version = "v2.7.2",
)

go_repository(
    name = "com_google_cloud_go_vmmigration",
    importpath = "cloud.google.com/go/vmmigration",
    sum = "h1:gnjIclgqbEMc+cF5IJuPxp53wjBIlqZ8h9hE8Rkwp7A=",
    version = "v1.7.1",
)

go_repository(
    name = "com_google_cloud_go_vpcaccess",
    importpath = "cloud.google.com/go/vpcaccess",
    sum = "h1:ram0GzjNWElmbxXMIzeOZUkQ9J8ZAahD6V8ilPGqX0Y=",
    version = "v1.7.1",
)

go_repository(
    name = "com_google_cloud_go_webrisk",
    importpath = "cloud.google.com/go/webrisk",
    sum = "h1:Ssy3MkOMOnyRV5H2bkMQ13Umv7CwB/kugo3qkAX83Fk=",
    version = "v1.9.1",
)

go_repository(
    name = "com_google_cloud_go_websecurityscanner",
    importpath = "cloud.google.com/go/websecurityscanner",
    sum = "h1:CfEF/vZ+xXyAR3zC9iaC/QRdf1MEgS20r5UR17Q4gOg=",
    version = "v1.6.1",
)

go_repository(
    name = "com_google_cloud_go_workflows",
    importpath = "cloud.google.com/go/workflows",
    sum = "h1:2akeQ/PgtRhrNuD/n1WvJd5zb7YyuDZrlOanBj2ihPg=",
    version = "v1.11.1",
)

go_repository(
    name = "com_github_urfave_cli_v2",
    importpath = "github.com/urfave/cli/v2",
    sum = "h1:VAzn5oq403l5pHjc4OhD54+XGO9cdKVL/7lDjF+iKUs=",
    version = "v2.25.7",
)

go_repository(
    name = "com_github_aws_aws_lambda_go",
    importpath = "github.com/aws/aws-lambda-go",
    sum = "h1:l/5fyVb6Ud9uYd411xdHZzSf2n86TakxzpvIoz7l+3Y=",
    version = "v1.41.0",
)

go_repository(
    name = "com_github_cenkalti_backoff_v4",
    importpath = "github.com/cenkalti/backoff/v4",
    sum = "h1:y4OZtCnogmCPw98Zjyt5a6+QwPLGkiQsYW5oUqylYbM=",
    version = "v4.2.1",
)

go_repository(
    name = "com_github_jarcoal_httpmock",
    importpath = "github.com/jarcoal/httpmock",
    sum = "h1:gSvTxxFR/MEMfsGrvRbdfpRUMBStovlSRLw0Ep1bwwc=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_kelseyhightower_envconfig",
    importpath = "github.com/kelseyhightower/envconfig",
    sum = "h1:Im6hONhd3pLkfDFsbRgu68RDNkGF1r3dvMUtDTo2cv8=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_okta_okta_sdk_golang_v2",
    importpath = "github.com/okta/okta-sdk-golang/v2",
    sum = "h1:EDKM+uOPfihOMNwgHMdno+NAsIfyXkVnoFAYVPay0YU=",
    version = "v2.20.0",
)

go_repository(
    name = "com_github_patrickmn_go_cache",
    importpath = "github.com/patrickmn/go-cache",
    sum = "h1:pSCLCl6joCFRnjpeojzOpEYs4q7Vditq8fySFG5ap3Y=",
    version = "v0.0.0-20180815053127-5633e0862627",
)

go_repository(
    name = "in_gopkg_square_go_jose_v2",
    importpath = "gopkg.in/square/go-jose.v2",
    sum = "h1:NGk74WTnPKBNUhNzQX7PYcTLUjoq7mzKk2OKbvwk2iI=",
    version = "v2.6.0",
)

go_repository(
    name = "com_google_cloud_go_apigeeregistry",
    importpath = "cloud.google.com/go/apigeeregistry",
    sum = "h1:hgq0ANLDx7t2FDZDJQrCMtCtddR/pjCqVuvQWGrQbXw=",
    version = "v0.7.1",
)

go_repository(
    name = "com_google_cloud_go_apikeys",
    importpath = "cloud.google.com/go/apikeys",
    sum = "h1:B9CdHFZTFjVti89tmyXXrO+7vSNo2jvZuHG8zD5trdQ=",
    version = "v0.6.0",
)

go_repository(
    name = "com_google_cloud_go_errorreporting",
    importpath = "cloud.google.com/go/errorreporting",
    sum = "h1:kj1XEWMu8P0qlLhm3FwcaFsUvXChV/OraZwA70trRR0=",
    version = "v0.3.0",
)

go_repository(
    name = "com_google_cloud_go_logging",
    importpath = "cloud.google.com/go/logging",
    sum = "h1:CJYxlNNNNAMkHp9em/YEXcfJg+rPDg7YfwoRpMU+t5I=",
    version = "v1.7.0",
)

go_repository(
    name = "com_google_cloud_go_maps",
    importpath = "cloud.google.com/go/maps",
    sum = "h1:PdfgpBLhAoSzZrQXP+/zBc78fIPLZSJp5y8+qSMn2UU=",
    version = "v1.4.0",
)

go_repository(
    name = "com_google_cloud_go_pubsublite",
    importpath = "cloud.google.com/go/pubsublite",
    sum = "h1:pX+idpWMIH30/K7c0epN6V703xpIcMXWRjKJsz0tYGY=",
    version = "v1.8.1",
)

go_repository(
    name = "com_google_cloud_go_spanner",
    importpath = "cloud.google.com/go/spanner",
    sum = "h1:aqiMP8dhsEXgn9K5EZBWxPG7dxIiyM2VaikqeU4iteg=",
    version = "v1.47.0",
)

go_repository(
    name = "com_google_cloud_go_vmwareengine",
    importpath = "cloud.google.com/go/vmwareengine",
    sum = "h1:qsJ0CPlOQu/3MFBGklu752v3AkD+Pdu091UmXJ+EjTA=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_jeffpignataro_golang_pkg_do_a_thing_v2",
    importpath = "github.com/jeffpignataro/golang/pkg/do-a-thing/v2",
    sum = "h1:ZRYia/PhSqGsVWBwwKa4vTWr8kKnx55KxT0WNVGxqEY=",
    version = "v2.0.0-20230318170428-6f5c8400faa0",
)

go_repository(
    name = "com_github_jeffpignataro_golang_pkg_do_a_thing_v3",
    importpath = "github.com/jeffpignataro/golang/pkg/do-a-thing/v3",
    sum = "h1:NNo+/YDdsJ+3RCUe8nDPVSgBa9ZW/y8wNmW2VHGnfJY=",
    version = "v3.0.0",
)

go_repository(
    name = "com_github_emicklei_go_restful_v3",
    importpath = "github.com/emicklei/go-restful/v3",
    sum = "h1:rAQeMHw1c7zTmncogyy8VvRZwtkmkZ4FxERmMY4rD+g=",
    version = "v3.11.0",
)

go_repository(
    name = "com_github_evanphx_json_patch",
    importpath = "github.com/evanphx/json-patch",
    sum = "h1:jBYDEEiFBPxA0v50tFdvOzQQTCvpL6mnFh5mB2/l16U=",
    version = "v5.6.0+incompatible",
)

go_repository(
    name = "com_github_go_logr_logr",
    importpath = "github.com/go-logr/logr",
    sum = "h1:g01GSCwiDw2xSZfjJ2/T9M+S6pFdcNtFYsp+Y43HYDQ=",
    version = "v1.2.4",
)

go_repository(
    name = "com_github_go_openapi_jsonpointer",
    importpath = "github.com/go-openapi/jsonpointer",
    sum = "h1:ESKJdU9ASRfaPNOPRx12IUyA1vn3R9GiE3KYD14BXdQ=",
    version = "v0.20.0",
)

go_repository(
    name = "com_github_go_openapi_jsonreference",
    importpath = "github.com/go-openapi/jsonreference",
    sum = "h1:3sVjiK66+uXK/6oQ8xgcRKcFgQ5KXa2KvnJRumpMGbE=",
    version = "v0.20.2",
)

go_repository(
    name = "com_github_go_openapi_swag",
    importpath = "github.com/go-openapi/swag",
    sum = "h1:QLMzNJnMGPRNDCbySlcj1x01tzU8/9LTTL9hZZZogBU=",
    version = "v0.22.4",
)

go_repository(
    name = "com_github_google_gnostic",
    importpath = "github.com/google/gnostic",
    sum = "h1:FhTMOKj2VhjpouxvWJAV1TL304uMlb9zcDqkl6cEI54=",
    version = "v0.5.7-v3refs",
)

go_repository(
    name = "com_github_google_gofuzz",
    importpath = "github.com/google/gofuzz",
    sum = "h1:xRy4A+RhZaiKjJ1bPfwQ8sedCA+YS2YcCHW6ec7JMi0=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_gregjones_httpcache",
    importpath = "github.com/gregjones/httpcache",
    sum = "h1:pdN6V1QBWetyv/0+wjACpqVH+eVULgEjkurDLq3goeM=",
    version = "v0.0.0-20180305231024-9cad4c3443a7",
)

go_repository(
    name = "com_github_imdario_mergo",
    importpath = "github.com/imdario/mergo",
    sum = "h1:xTNEAn+kxVO7dTZGu0CegyqKZmoWFI0rF8UxjlB2d28=",
    version = "v0.3.6",
)

go_repository(
    name = "com_github_josharian_intern",
    importpath = "github.com/josharian/intern",
    sum = "h1:vlS4z54oSdjm0bgjRigI+G1HpF+tI+9rE5LLzOg8HmY=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_mailru_easyjson",
    importpath = "github.com/mailru/easyjson",
    sum = "h1:UGYAvKxe3sBsEDzO8ZeWOSlIQfWFlxbzLZe7hwFURr0=",
    version = "v0.7.7",
)

go_repository(
    name = "com_github_moby_spdystream",
    importpath = "github.com/moby/spdystream",
    sum = "h1:cjW1zVyyoiM0T7b6UoySUFqzXMoqRckQtXwGPiBhOM8=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_munnerz_goautoneg",
    importpath = "github.com/munnerz/goautoneg",
    sum = "h1:C3w9PqII01/Oq1c1nUAm88MOHcQC9l5mIlSMApZMrHA=",
    version = "v0.0.0-20191010083416-a7dc8b61c822",
)

go_repository(
    name = "com_github_peterbourgon_diskv",
    importpath = "github.com/peterbourgon/diskv",
    sum = "h1:UBdAOUP5p4RWqPBg048CAvpKN+vxiaj6gdUUzhl4XmI=",
    version = "v2.0.1+incompatible",
)

go_repository(
    name = "in_gopkg_inf_v0",
    importpath = "gopkg.in/inf.v0",
    sum = "h1:73M5CoZyi3ZLMOyDlQh031Cx6N9NDJ2Vvfl76EDAgDc=",
    version = "v0.9.1",
)

go_repository(
    name = "io_k8s_api",
    importpath = "k8s.io/api",
    sum = "h1:i+0O8k2NPBCPYaMB+uCkseEbawEt/eFaiRqUx8aB108=",
    version = "v0.28.1",
)

go_repository(
    name = "io_k8s_apimachinery",
    importpath = "k8s.io/apimachinery",
    sum = "h1:EJD40og3GizBSV3mkIoXQBsws32okPOy+MkRyzh6nPY=",
    version = "v0.28.1",
)

go_repository(
    name = "io_k8s_client_go",
    importpath = "k8s.io/client-go",
    sum = "h1:pRhMzB8HyLfVwpngWKE8hDcXRqifh1ga2Z/PU9SXVK8=",
    version = "v0.28.1",
)

go_repository(
    name = "io_k8s_klog_v2",
    importpath = "k8s.io/klog/v2",
    sum = "h1:7WCHKK6K8fNhTqfBhISHQ97KrnJNFZMcQvKp7gP/tmg=",
    version = "v2.100.1",
)

go_repository(
    name = "io_k8s_kube_openapi",
    importpath = "k8s.io/kube-openapi",
    sum = "h1:CAIciCnJnSOQxPd0xvpV6JU3D4AJvnYbImPpFpO9Hnw=",
    version = "v0.0.0-20230816210353-14e408962443",
)

go_repository(
    name = "io_k8s_sigs_json",
    importpath = "sigs.k8s.io/json",
    sum = "h1:EDPBXCAspyGV4jQlpZSudPeMmr1bNJefnuqLsRAsHZo=",
    version = "v0.0.0-20221116044647-bc3834ca7abd",
)

go_repository(
    name = "io_k8s_sigs_structured_merge_diff_v4",
    importpath = "sigs.k8s.io/structured-merge-diff/v4",
    sum = "h1:UZbZAZfX0wV2zr7YZorDz6GXROfDFj6LvqCRm4VUVKk=",
    version = "v4.3.0",
)

go_repository(
    name = "io_k8s_sigs_yaml",
    importpath = "sigs.k8s.io/yaml",
    sum = "h1:a2VclLzOGrwOHDiV8EfBGhvjHvP46CtW5j6POvhYGGo=",
    version = "v1.3.0",
)

go_repository(
    name = "io_k8s_utils",
    importpath = "k8s.io/utils",
    sum = "h1:sgn3ZU783SCgtaSJjpcVVlRqd6GSnlTLKgpAAttJvpI=",
    version = "v0.0.0-20230726121419-3b25d923346b",
)

go_repository(
    name = "com_github_armon_go_socks5",
    importpath = "github.com/armon/go-socks5",
    sum = "h1:0CwZNZbxp69SHPdPJAN/hZIm0C4OItdklCFmMRWYpio=",
    version = "v0.0.0-20160902184237-e75332964ef5",
)

go_repository(
    name = "com_github_elazarl_goproxy",
    importpath = "github.com/elazarl/goproxy",
    sum = "h1:p1yVGRW3nmb85p1Sh1ZJSDm4A4iKLS5QNbvUHMgGu/M=",
    version = "v0.0.0-20170405201442-c4fc26588b6e",
)

go_repository(
    name = "com_github_mxk_go_flowrate",
    importpath = "github.com/mxk/go-flowrate",
    sum = "h1:y5//uYreIhSUg3J1GEMiLbxo1LJaP8RfCpH6pymGZus=",
    version = "v0.0.0-20140419014527-cca7078d478f",
)

go_repository(
    name = "com_github_niemeyer_pretty",
    importpath = "github.com/niemeyer/pretty",
    sum = "h1:fD57ERR4JtEqsWbfPhv4DMiApHyliiK5xCTNVSPiaAs=",
    version = "v0.0.0-20200227124842-a10e7caefd8e",
)

go_repository(
    name = "com_github_onsi_ginkgo_v2",
    importpath = "github.com/onsi/ginkgo/v2",
    sum = "h1:xR7vG4IXt5RWx6FfIjyAtsoMAtnc3C/rFXBBd2AjZwE=",
    version = "v2.9.4",
)

go_repository(
    name = "com_github_onsi_gomega",
    importpath = "github.com/onsi/gomega",
    sum = "h1:ENqfyGeS5AX/rlXDd/ETokDz93u0YufY1Pgxuy/PvWE=",
    version = "v1.27.6",
)

go_repository(
    name = "com_github_asaskevich_govalidator",
    importpath = "github.com/asaskevich/govalidator",
    sum = "h1:idn718Q4B6AGu/h5Sxe66HYVdqdGu2l9Iebqhi/AEoA=",
    version = "v0.0.0-20190424111038-f61b66f89f4a",
)

go_repository(
    name = "com_github_creack_pty",
    importpath = "github.com/creack/pty",
    sum = "h1:uDmaGzcdjhF4i/plgjmEsriH11Y0o7RKapEf/LDaM3w=",
    version = "v1.1.9",
)

go_repository(
    name = "com_github_docopt_docopt_go",
    importpath = "github.com/docopt/docopt-go",
    sum = "h1:bWDMxwH3px2JBh6AyO7hdCn/PkvCZXii8TGj7sbtEbQ=",
    version = "v0.0.0-20180111231733-ee0de3bc6815",
)

go_repository(
    name = "com_github_kisielk_errcheck",
    importpath = "github.com/kisielk/errcheck",
    sum = "h1:e8esj/e4R+SAOwFwN+n3zr0nYeCyeweozKfO23MvHzY=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_nytimes_gziphandler",
    importpath = "github.com/NYTimes/gziphandler",
    sum = "h1:lsxEuwrXEAokXB9qhlbKWPpo3KMLZQ5WB5WLQRW1uq0=",
    version = "v0.0.0-20170623195520-56545f4a5d46",
)

go_repository(
    name = "com_github_puerkitobio_purell",
    importpath = "github.com/PuerkitoBio/purell",
    sum = "h1:0GoNN3taZV6QI81IXgCbxMyEaJDXMSIjArYBCYzVVvs=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_puerkitobio_urlesc",
    importpath = "github.com/PuerkitoBio/urlesc",
    sum = "h1:JCHLVE3B+kJde7bIEo5N4J+ZbLhp0J1Fs+ulyRws4gE=",
    version = "v0.0.0-20160726150825-5bd2802263f2",
)

go_repository(
    name = "com_github_stoewer_go_strcase",
    importpath = "github.com/stoewer/go-strcase",
    sum = "h1:Z2iHWqGXH00XYgqDmNgQbIBxf3wrNq0F3feEy0ainaU=",
    version = "v1.2.0",
)

go_repository(
    name = "io_k8s_gengo",
    importpath = "k8s.io/gengo",
    sum = "h1:GohjlNKauSai7gN4wsJkeZ3WAJx4Sh+oT/b5IYn5suA=",
    version = "v0.0.0-20210813121822-485abfe95c7c",
)

go_repository(
    name = "com_github_antihax_optional",
    importpath = "github.com/antihax/optional",
    sum = "h1:xK2lYat7ZLaVVcIuj82J8kIro4V6kDe0AUDFboUCwcg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_buger_jsonparser",
    importpath = "github.com/buger/jsonparser",
    sum = "h1:2PnMjfWD7wBILjqQbt530v576A/cAbQvEW9gGIpYMUs=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_cespare_xxhash",
    importpath = "github.com/cespare/xxhash",
    sum = "h1:a6HrQnmkObjyL+Gs60czilIUGqrzKutQD6XZog3p+ko=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_flowstack_go_jsonschema",
    importpath = "github.com/flowstack/go-jsonschema",
    sum = "h1:dCrjGJRXIlbDsLAgTJZTjhwUJnnxVWl1OgNyYh5nyDc=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_ghodss_yaml",
    importpath = "github.com/ghodss/yaml",
    sum = "h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_grpc_ecosystem_grpc_gateway",
    importpath = "github.com/grpc-ecosystem/grpc-gateway",
    sum = "h1:gmcG1KaJ57LophUzW0Hy8NmPhnMZb4M0+kPpLofRdBo=",
    version = "v1.16.0",
)

go_repository(
    name = "com_github_oneofone_xxhash",
    importpath = "github.com/OneOfOne/xxhash",
    sum = "h1:KMrpdQIwFcEqXDklaen+P1axHaj9BSKzvpUUfnHldSE=",
    version = "v1.2.2",
)

go_repository(
    name = "com_github_rogpeppe_fastuuid",
    importpath = "github.com/rogpeppe/fastuuid",
    sum = "h1:Ppwyp6VYCF1nvBTXL3trRso7mXMlRrw9ooo375wvi2s=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_spaolacci_murmur3",
    importpath = "github.com/spaolacci/murmur3",
    sum = "h1:qLC7fQah7D6K1B0ujays3HV9gkFtllcxhzImRR7ArPQ=",
    version = "v0.0.0-20180118202830-f09979ecbc72",
)

go_repository(
    name = "com_github_xeipuuv_gojsonpointer",
    importpath = "github.com/xeipuuv/gojsonpointer",
    sum = "h1:J9EGpcZtP0E/raorCMxlFGSTBrsSlaDGf3jU/qvAE2c=",
    version = "v0.0.0-20180127040702-4e3ac2762d5f",
)

go_repository(
    name = "com_github_xeipuuv_gojsonreference",
    importpath = "github.com/xeipuuv/gojsonreference",
    sum = "h1:EzJWgHovont7NscjpAxXsDA8S8BMYve8Y5+7cuRE7R0=",
    version = "v0.0.0-20180127040603-bd5ef7bd5415",
)

go_repository(
    name = "com_github_xeipuuv_gojsonschema",
    importpath = "github.com/xeipuuv/gojsonschema",
    sum = "h1:LhYJRs+L4fBtjZUfuSZIKGeVu0QRy8e5Xi7D17UxZ74=",
    version = "v1.2.0",
)

go_repository(
    name = "io_opentelemetry_go_proto_otlp",
    importpath = "go.opentelemetry.io/proto/otlp",
    sum = "h1:rwOQPCuKAKmwGKq2aVNnYIibI6wnV7EvzgfTCzcdGg8=",
    version = "v0.7.0",
)

go_repository(
    name = "com_github_go_test_deep",
    importpath = "github.com/go-test/deep",
    sum = "h1:u2CU3YKy9I2pmu9pX0eq50wCgjfGIt539SqR7FbHiho=",
    version = "v1.0.4",
)

go_repository(
    name = "com_github_gorilla_websocket",
    importpath = "github.com/gorilla/websocket",
    sum = "h1:PPwGk2jz7EePpoHN/+ClbZu8SPxiqlu12wZP/3sWmnc=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_slack_go_slack",
    importpath = "github.com/slack-go/slack",
    sum = "h1:92/dfFU8Q5XP6Wp5rr5/T5JHLM5c5Smtn53fhToAP88=",
    version = "v0.12.3",
)

go_repository(
    name = "com_github_sirupsen_logrus",
    importpath = "github.com/sirupsen/logrus",
    sum = "h1:dueUQJ1C2q9oE3F7wvmSGAaVtTmUizReu6fjN8uqzbQ=",
    version = "v1.9.3",
)

go_repository(
    name = "com_github_g3n_engine",
    importpath = "github.com/g3n/engine",
    sum = "h1:7dmj4c+3xHcBnYrVmRuVf/oZ2JycxJU9Y+2FQj1Af2Y=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_golang_freetype",
    importpath = "github.com/golang/freetype",
    sum = "h1:DACJavvAHhabrF08vX0COfcOBJRhZ8lUbR+ZWIs0Y5g=",
    version = "v0.0.0-20170609003504-e2365dfdc4a0",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2",
    importpath = "github.com/aws/aws-sdk-go-v2",
    sum = "h1:gMT0IW+03wtYJhRqTVYn0wLzwdnK9sRMcxmtfGzRdJc=",
    version = "v1.21.0",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_config",
    importpath = "github.com/aws/aws-sdk-go-v2/config",
    sum = "h1:RNAfbPqw1CstCooHaTPhScz7z1PyocQj0UL+l95CgzI=",
    version = "v1.18.37",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_credentials",
    importpath = "github.com/aws/aws-sdk-go-v2/credentials",
    sum = "h1:QpsNitYJu0GgvMBLUIYu9H4yryA5kMksjeIVQfgXrt8=",
    version = "v1.13.35",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_feature_ec2_imds",
    importpath = "github.com/aws/aws-sdk-go-v2/feature/ec2/imds",
    sum = "h1:uDZJF1hu0EVT/4bogChk8DyjSF6fof6uL/0Y26Ma7Fg=",
    version = "v1.13.11",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_internal_configsources",
    importpath = "github.com/aws/aws-sdk-go-v2/internal/configsources",
    sum = "h1:22dGT7PneFMx4+b3pz7lMTRyN8ZKH7M2cW4GP9yUS2g=",
    version = "v1.1.41",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_internal_endpoints_v2",
    importpath = "github.com/aws/aws-sdk-go-v2/internal/endpoints/v2",
    sum = "h1:SijA0mgjV8E+8G45ltVHs0fvKpTj8xmZJ3VwhGKtUSI=",
    version = "v2.4.35",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_service_internal_presigned_url",
    importpath = "github.com/aws/aws-sdk-go-v2/service/internal/presigned-url",
    sum = "h1:CdzPW9kKitgIiLV1+MHobfR5Xg25iYnyzWZhyQuSlDI=",
    version = "v1.9.35",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_service_sfn",
    importpath = "github.com/aws/aws-sdk-go-v2/service/sfn",
    sum = "h1:uGuCRiB/3dCGb0iInxJJJTeMvTxM7wIdFv9R0uSFLKQ=",
    version = "v1.19.5",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_service_sso",
    importpath = "github.com/aws/aws-sdk-go-v2/service/sso",
    sum = "h1:oCvTFSDi67AX0pOX3PuPdGFewvLRU2zzFSrTsgURNo0=",
    version = "v1.13.5",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_service_ssooidc",
    importpath = "github.com/aws/aws-sdk-go-v2/service/ssooidc",
    sum = "h1:dnInJb4S0oy8aQuri1mV6ipLlnZPfnsDNB9BGO9PDNY=",
    version = "v1.15.5",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2_service_sts",
    importpath = "github.com/aws/aws-sdk-go-v2/service/sts",
    sum = "h1:CQBFElb0LS8RojMJlxRSo/HXipvTZW2S44Lt9Mk2aYQ=",
    version = "v1.21.5",
)

go_repository(
    name = "com_github_aws_smithy_go",
    importpath = "github.com/aws/smithy-go",
    sum = "h1:MJU9hqBGbvWZdApzpvoF2WAIJDbtjK2NDJSiJP7HblQ=",
    version = "v1.14.2",
)

go_repository(
    name = "com_github_billheroinc_logrus",
    importpath = "github.com/BillHeroInc/logrus",
    sum = "h1:NELeUSqRVUBgqkZJS9oEnfMmtFo0a6imfjRhEP0ELK4=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_jmespath_go_jmespath",
    importpath = "github.com/jmespath/go-jmespath",
    sum = "h1:BEgLn5cpjn8UN1mAw4NjwDrS35OdebyEtFe+9YPoQUg=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_jmespath_go_jmespath_internal_testify",
    importpath = "github.com/jmespath/go-jmespath/internal/testify",
    sum = "h1:shLQSRRSCCPj3f2gpwzGwWFoC7ycTf1rcQZHOlsJ6N8=",
    version = "v1.5.1",
)

go_repository(
    name = "com_github_aymanbagabas_go_osc52",
    importpath = "github.com/aymanbagabas/go-osc52",
    sum = "h1:q2sWUyDcozPLcLabEMd+a+7Ea2DitxZVN9hTxab9L4E=",
    version = "v1.2.1",
)

go_repository(
    name = "com_github_charmbracelet_bubbletea",
    importpath = "github.com/charmbracelet/bubbletea",
    sum = "h1:uaQIKx9Ai6Gdh5zpTbGiWpytMU+CfsPp06RaW2cx/SY=",
    version = "v0.24.2",
)

go_repository(
    name = "com_github_containerd_console",
    importpath = "github.com/containerd/console",
    sum = "h1:q2hJAaP1k2wIvVRd/hEHD7lacgqrCPS+k8g1MndzfWY=",
    version = "v1.0.4-0.20230313162750-1ae8d489ac81",
)

go_repository(
    name = "com_github_lucasb_eyer_go_colorful",
    importpath = "github.com/lucasb-eyer/go-colorful",
    sum = "h1:1nnpGOrhyZZuNyfu1QjKiUICQ74+3FNCN69Aj6K7nkY=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_mattn_go_localereader",
    importpath = "github.com/mattn/go-localereader",
    sum = "h1:ygSAOl7ZXTx4RdPYinUpg6W99U8jWvWi9Ye2JC/oIi4=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_mattn_go_runewidth",
    importpath = "github.com/mattn/go-runewidth",
    sum = "h1:UNAjwbU9l54TA3KzvqLGxwWjHmMgBUVhBiTjelZgg3U=",
    version = "v0.0.15",
)

go_repository(
    name = "com_github_muesli_ansi",
    importpath = "github.com/muesli/ansi",
    sum = "h1:ZK8zHtRHOkbHy6Mmr5D264iyp3TiX5OmNcI5cIARiQI=",
    version = "v0.0.0-20230316100256-276c6243b2f6",
)

go_repository(
    name = "com_github_muesli_cancelreader",
    importpath = "github.com/muesli/cancelreader",
    sum = "h1:3I4Kt4BQjOR54NavqnDogx/MIoWBFa0StPA8ELUXHmA=",
    version = "v0.2.2",
)

go_repository(
    name = "com_github_muesli_reflow",
    importpath = "github.com/muesli/reflow",
    sum = "h1:IFsN6K9NfGtjeggFP+68I4chLZV2yIKsXJFNZ+eWh6s=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_muesli_termenv",
    importpath = "github.com/muesli/termenv",
    sum = "h1:GohcuySI0QmI3wN8Ok9PtKGkgkFIk7y6Vpb5PvrY+Wo=",
    version = "v0.15.2",
)

go_repository(
    name = "com_github_rivo_uniseg",
    importpath = "github.com/rivo/uniseg",
    sum = "h1:8TfxU8dW6PdqD27gjM8MVNuicgxIjxpm4K7x4jp8sis=",
    version = "v0.4.4",
)

go_repository(
    name = "com_github_abdullin_seq",
    importpath = "github.com/abdullin/seq",
    sum = "h1:DBNMBMuMiWYu0b+8KMJuWmfCkcxl09JwdlqwDZZ6U14=",
    version = "v0.0.0-20160510034733-d5467c17e7af",
)

go_repository(
    name = "com_github_agext_levenshtein",
    importpath = "github.com/agext/levenshtein",
    sum = "h1:YB2fHEn0UJagG8T1rrWknE3ZQzWM06O8AMAatNn7lmo=",
    version = "v1.2.3",
)

go_repository(
    name = "com_github_agl_ed25519",
    importpath = "github.com/agl/ed25519",
    sum = "h1:LoeFxdq5zUCBQPhbQKE6zvoGwHMxCBlqwbH9+9kHoHA=",
    version = "v0.0.0-20150830182803-278e1ec8e8a6",
)

go_repository(
    name = "com_github_alecthomas_template",
    importpath = "github.com/alecthomas/template",
    sum = "h1:JYp7IbQjafoB+tBA3gMyHYHrpOtNuDiK/uB5uXxq5wM=",
    version = "v0.0.0-20190718012654-fb15b899a751",
)

go_repository(
    name = "com_github_alecthomas_units",
    importpath = "github.com/alecthomas/units",
    sum = "h1:Hs82Z41s6SdL1CELW+XaDYmOH4hkBN4/N9og/AsOv7E=",
    version = "v0.0.0-20190717042225-c3de453c63f4",
)

go_repository(
    name = "com_github_aliyun_alibaba_cloud_sdk_go",
    importpath = "github.com/aliyun/alibaba-cloud-sdk-go",
    sum = "h1:Ij3S0pNUMgHlhx3Ew8g9RNrt59EKhHYdMODGtFXJfSc=",
    version = "v1.61.1501",
)

go_repository(
    name = "com_github_aliyun_aliyun_oss_go_sdk",
    importpath = "github.com/aliyun/aliyun-oss-go-sdk",
    sum = "h1:FrF4uxA24DF3ARNXVbUin3wa5fDLaB1Cy8mKks/LRz4=",
    version = "v0.0.0-20190103054945-8205d1f41e70",
)

go_repository(
    name = "com_github_aliyun_aliyun_tablestore_go_sdk",
    importpath = "github.com/aliyun/aliyun-tablestore-go-sdk",
    sum = "h1:ABQ7FF+IxSFHDMOTtjCfmMDMHiCq6EsAoCV/9sFinaM=",
    version = "v4.1.2+incompatible",
)

go_repository(
    name = "com_github_antchfx_xpath",
    importpath = "github.com/antchfx/xpath",
    sum = "h1:ptBAamGVd6CfRsUtyHD+goy2JGhv1QC32v3gqM8mYAM=",
    version = "v0.0.0-20190129040759-c8489ed3251e",
)

go_repository(
    name = "com_github_antchfx_xquery",
    importpath = "github.com/antchfx/xquery",
    sum = "h1:JaCC8jz0zdMLk2m+qCCVLLLM/PL93p84w4pK3aJWj60=",
    version = "v0.0.0-20180515051857-ad5b8c7a47b0",
)

go_repository(
    name = "com_github_apparentlymart_go_cidr",
    importpath = "github.com/apparentlymart/go-cidr",
    sum = "h1:2mAhrMoF+nhXqxTzSZMUzDHkLjmIHC+Zzn4tdgBZjnU=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_apparentlymart_go_dump",
    importpath = "github.com/apparentlymart/go-dump",
    sum = "h1:MzVXffFUye+ZcSR6opIgz9Co7WcDx6ZcY+RjfFHoA0I=",
    version = "v0.0.0-20190214190832-042adf3cf4a0",
)

go_repository(
    name = "com_github_apparentlymart_go_textseg",
    importpath = "github.com/apparentlymart/go-textseg",
    sum = "h1:rRmlIsPEEhUTIKQb7T++Nz/A5Q6C9IuX2wFoYVvnCs0=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_apparentlymart_go_textseg_v12",
    importpath = "github.com/apparentlymart/go-textseg/v12",
    sum = "h1:bNEQyAGak9tojivJNkoqWErVCQbjdL7GzRt3F8NvfJ0=",
    version = "v12.0.0",
)

go_repository(
    name = "com_github_apparentlymart_go_textseg_v13",
    importpath = "github.com/apparentlymart/go-textseg/v13",
    sum = "h1:Y+KvPE1NYz0xl601PVImeQfFyEy6iT90AvPUL1NNfNw=",
    version = "v13.0.0",
)

go_repository(
    name = "com_github_apparentlymart_go_userdirs",
    importpath = "github.com/apparentlymart/go-userdirs",
    sum = "h1:JtuelWqyixKApmXm3qghhZ7O96P6NKpyrlSIe8Rwnhw=",
    version = "v0.0.0-20200915174352-b0c018a67c13",
)

go_repository(
    name = "com_github_apparentlymart_go_versions",
    importpath = "github.com/apparentlymart/go-versions",
    sum = "h1:n5Gg9YvSLK8Zzpy743J7abh2jt7z7ammOQ0oTd/5oA4=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_armon_circbuf",
    importpath = "github.com/armon/circbuf",
    sum = "h1:7Ip0wMmLHLRJdrloDxZfhMm0xrLXZS8+COSu2bXmEQs=",
    version = "v0.0.0-20190214190532-5111143e8da2",
)

go_repository(
    name = "com_github_armon_go_radix",
    importpath = "github.com/armon/go-radix",
    sum = "h1:F4z6KzEeeQIMeLFa97iZU6vupzoecKdU5TX24SNppXI=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_aws_aws_sdk_go",
    importpath = "github.com/aws/aws-sdk-go",
    sum = "h1:p6mw01WBaNpbdP2xrisz5tIkcNwzj/HysobNoaAHjgo=",
    version = "v1.44.122",
)

go_repository(
    name = "com_github_azure_azure_sdk_for_go",
    importpath = "github.com/Azure/azure-sdk-for-go",
    sum = "h1:mbxiZy1K820hQ+dI+YIO/+a0wQDYqOu18BAGe4lXjVk=",
    version = "v59.2.0+incompatible",
)

go_repository(
    name = "com_github_azure_go_autorest",
    importpath = "github.com/Azure/go-autorest",
    sum = "h1:V5VMDjClD3GiElqLWO7mz2MxNAK/vTfRHdAubSIPRgs=",
    version = "v14.2.0+incompatible",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest",
    importpath = "github.com/Azure/go-autorest/autorest",
    sum = "h1:1fIGgHKqVm54KIPT+q8Zmd1QlVsmHqeUGso5qm2BqqE=",
    version = "v0.11.24",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest_adal",
    importpath = "github.com/Azure/go-autorest/autorest/adal",
    sum = "h1:kLnPsRjzZZUF3K5REu/Kc+qMQrvuza2bwSnNdhmzLfQ=",
    version = "v0.9.18",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest_azure_cli",
    importpath = "github.com/Azure/go-autorest/autorest/azure/cli",
    sum = "h1:iuooz5cZL6VRcO7DVSFYxRcouqn6bFVE/e77Wts50Zk=",
    version = "v0.4.4",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest_date",
    importpath = "github.com/Azure/go-autorest/autorest/date",
    sum = "h1:7gUk1U5M/CQbp9WoqinNzJar+8KY+LPI6wiWrP/myHw=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest_mocks",
    importpath = "github.com/Azure/go-autorest/autorest/mocks",
    sum = "h1:K0laFcLE6VLTOwNgSxaGbUcLPuGXlNkbVvq4cW4nIHk=",
    version = "v0.4.1",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest_to",
    importpath = "github.com/Azure/go-autorest/autorest/to",
    sum = "h1:oXVqrxakqqV1UZdSazDOPOLvOIz+XA683u8EctwboHk=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest_validation",
    importpath = "github.com/Azure/go-autorest/autorest/validation",
    sum = "h1:AgyqjAd94fwNAoTjl/WQXg4VvFeRFpO+UhNyRXqF1ac=",
    version = "v0.3.1",
)

go_repository(
    name = "com_github_azure_go_autorest_logger",
    importpath = "github.com/Azure/go-autorest/logger",
    sum = "h1:IG7i4p/mDa2Ce4TRyAO8IHnVhAVF3RFU+ZtXWSmf4Tg=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_azure_go_autorest_tracing",
    importpath = "github.com/Azure/go-autorest/tracing",
    sum = "h1:TYi4+3m5t6K48TGI9AUdb+IzbnSxvnvUMfuitfgcfuo=",
    version = "v0.6.0",
)

go_repository(
    name = "com_github_azure_go_ntlmssp",
    importpath = "github.com/Azure/go-ntlmssp",
    sum = "h1:/IBSNwUN8+eKzUzbJPqhK839ygXJ82sde8x3ogr6R28=",
    version = "v0.0.0-20200615164410-66371956d46c",
)

go_repository(
    name = "com_github_baiyubin_aliyun_sts_go_sdk",
    importpath = "github.com/baiyubin/aliyun-sts-go-sdk",
    sum = "h1:ZNv7On9kyUzm7fvRZumSyy/IUiSC7AzL0I1jKKtwooA=",
    version = "v0.0.0-20180326062324-cfa1a18b161f",
)

go_repository(
    name = "com_github_beorn7_perks",
    importpath = "github.com/beorn7/perks",
    sum = "h1:VlbKKnNfV8bJzeqoa4cOKqO6bYr3WgKZxO8Z16+hsOM=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_bgentry_go_netrc",
    importpath = "github.com/bgentry/go-netrc",
    sum = "h1:xDfNPAt8lFiC1UJrqV3uuy861HCTo708pDMbjHHdCas=",
    version = "v0.0.0-20140422174119-9fd32a8b3d3d",
)

go_repository(
    name = "com_github_bgentry_speakeasy",
    importpath = "github.com/bgentry/speakeasy",
    sum = "h1:ByYyxL9InA1OWqxJqqp2A5pYHUrCiAL6K3J+LKSsQkY=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_blang_semver",
    importpath = "github.com/blang/semver",
    sum = "h1:cQNTCjp13qL8KC3Nbxr/y2Bqb63oX6wdnnjpJbkM4JQ=",
    version = "v3.5.1+incompatible",
)

go_repository(
    name = "com_github_bmatcuk_doublestar",
    importpath = "github.com/bmatcuk/doublestar",
    sum = "h1:gPypJ5xD31uhX6Tf54sDPUOBXTqKH4c9aPY66CyQrS0=",
    version = "v1.3.4",
)

go_repository(
    name = "com_github_boltdb_bolt",
    importpath = "github.com/boltdb/bolt",
    sum = "h1:JQmyP4ZBrce+ZQu0dY660FMfatumYDLun9hBCUVIkF4=",
    version = "v1.3.1",
)

go_repository(
    name = "com_github_cheggaaa_pb",
    importpath = "github.com/cheggaaa/pb",
    sum = "h1:wIkZHkNfC7R6GI5w7l/PdAdzXzlrbcI3p8OAlnkTsnc=",
    version = "v1.0.27",
)

go_repository(
    name = "com_github_christrenkamp_goxpath",
    importpath = "github.com/ChrisTrenkamp/goxpath",
    sum = "h1:W1diKnDQkXxNDhghdBSbQ4LI/E1aJNTwpqPp3KtlB8w=",
    version = "v0.0.0-20190607011252-c5096ec8773d",
)

go_repository(
    name = "com_github_coreos_bbolt",
    importpath = "github.com/coreos/bbolt",
    sum = "h1:HIgH5xUWXT914HCI671AxuTTqjj64UOFr7pHn48LUTI=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_coreos_etcd",
    importpath = "github.com/coreos/etcd",
    sum = "h1:jFneRYjIvLMLhDLCzuTuU4rSJUjRplcJQ7pD7MnhC04=",
    version = "v3.3.10+incompatible",
)

go_repository(
    name = "com_github_coreos_go_systemd",
    importpath = "github.com/coreos/go-systemd",
    sum = "h1:t5Wuyh53qYyg9eqn4BbnlIT+vmhyww0TatL+zT3uWgI=",
    version = "v0.0.0-20181012123002-c6f51f82210d",
)

go_repository(
    name = "com_github_coreos_pkg",
    importpath = "github.com/coreos/pkg",
    sum = "h1:lBNOc5arjvs8E5mO2tbpBpLoyyu8B6e44T7hJy6potg=",
    version = "v0.0.0-20180928190104-399ea9e2e55f",
)

go_repository(
    name = "com_github_dgrijalva_jwt_go",
    importpath = "github.com/dgrijalva/jwt-go",
    sum = "h1:7qlOGliEKZXTDg6OTjfoBKDXWrumCAMpl/TFQ4/5kLM=",
    version = "v3.2.0+incompatible",
)

go_repository(
    name = "com_github_dimchansky_utfbom",
    importpath = "github.com/dimchansky/utfbom",
    sum = "h1:vV6w1AhK4VMnhBno/TPVCoK9U/LP0PkLCS9tbxHdi/U=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_docker_spdystream",
    importpath = "github.com/docker/spdystream",
    sum = "h1:cenwrSVm+Z7QLSV/BsnenAOcDXdX4cMv4wP0B/5QbPg=",
    version = "v0.0.0-20160310174837-449fdfce4d96",
)

go_repository(
    name = "com_github_dylanmei_iso8601",
    importpath = "github.com/dylanmei/iso8601",
    sum = "h1:812NGQDBcqquTfH5Yeo7lwR0nzx/cKdsmf3qMjPURUI=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_dylanmei_winrmtest",
    importpath = "github.com/dylanmei/winrmtest",
    sum = "h1:r1oACdS2XYiAWcfF8BJXkoU8l1J71KehGR+d99yWEDA=",
    version = "v0.0.0-20190225150635-99b7fe2fddf1",
)

go_repository(
    name = "com_github_emicklei_go_restful",
    importpath = "github.com/emicklei/go-restful",
    sum = "h1:H2pdYOb3KQ1/YsqVWoWNLQO+fusocsw354rqGTZtAgw=",
    version = "v0.0.0-20170410110728-ff4f55a20633",
)

go_repository(
    name = "com_github_go_kit_kit",
    importpath = "github.com/go-kit/kit",
    sum = "h1:wDJmvq38kDhkVxi50ni9ykkdUr1PKgqKOoi01fa0Mdk=",
    version = "v0.9.0",
)

go_repository(
    name = "com_github_go_logfmt_logfmt",
    importpath = "github.com/go-logfmt/logfmt",
    sum = "h1:MP4Eh7ZCb31lleYCFuwm0oe4/YGak+5l1vA2NOE80nA=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_go_openapi_spec",
    importpath = "github.com/go-openapi/spec",
    sum = "h1:C1JKChikHGpXwT5UQDFaryIpDtyyGL/CR6C2kB7F1oc=",
    version = "v0.0.0-20160808142527-6aced65f8501",
)

go_repository(
    name = "com_github_go_sql_driver_mysql",
    importpath = "github.com/go-sql-driver/mysql",
    sum = "h1:ozyZYNQW3x3HtqT1jira07DN2PArx2v7/mN66gGcHOs=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_go_stack_stack",
    importpath = "github.com/go-stack/stack",
    sum = "h1:5SgMzNM5HxrEjV0ww2lTmX6E2Izsfxas4+YHWRs3Lsk=",
    version = "v1.8.0",
)

go_repository(
    name = "com_github_golang_snappy",
    importpath = "github.com/golang/snappy",
    sum = "h1:yAGX7huGHXlcLOEtBnF4w7FQwA26wojNCwOYAEhLjQM=",
    version = "v0.0.4",
)

go_repository(
    name = "com_github_google_go_querystring",
    importpath = "github.com/google/go-querystring",
    sum = "h1:AnCroh3fv4ZBgVIf1Iwtovgjaw/GiKJo8M8yD/fhyJ8=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_googleapis_go_type_adapters",
    importpath = "github.com/googleapis/go-type-adapters",
    sum = "h1:9XdMn+d/G57qq1s8dNc5IesGCXHf6V2HZ2JwRxfA2tA=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_gophercloud_gophercloud",
    importpath = "github.com/gophercloud/gophercloud",
    sum = "h1:Ciwp7ro4LyptUOkili/TX/ecuYr7vGtEIFnOOOKUjD8=",
    version = "v0.10.1-0.20200424014253-c3bfe50899e5",
)

go_repository(
    name = "com_github_gophercloud_utils",
    importpath = "github.com/gophercloud/utils",
    sum = "h1:fduaPzWwIfvOMLuHk2Al3GZH0XbUqG8MbElPop+Igzs=",
    version = "v0.0.0-20200423144003-7c72efc7435d",
)

go_repository(
    name = "com_github_gopherjs_gopherjs",
    importpath = "github.com/gopherjs/gopherjs",
    sum = "h1:EGx4pi6eqNxGaHF6qqu48+N2wcFQ5qg5FXgOdqsJ5d8=",
    version = "v0.0.0-20181017120253-0766667cb4d1",
)

go_repository(
    name = "com_github_grpc_ecosystem_go_grpc_middleware",
    importpath = "github.com/grpc-ecosystem/go-grpc-middleware",
    sum = "h1:Iju5GlWwrvL6UBg4zJJt3btmonfrMlCDdsejg4CZE7c=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_grpc_ecosystem_go_grpc_prometheus",
    importpath = "github.com/grpc-ecosystem/go-grpc-prometheus",
    sum = "h1:Ovs26xHkKqVztRpIrF/92BcuyuQ/YW4NSIpoGtfXNho=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_hashicorp_aws_sdk_go_base",
    importpath = "github.com/hashicorp/aws-sdk-go-base",
    sum = "h1:7s/aR3hFn74tYPVihzDyZe7y/+BorN70rr9ZvpV3j3o=",
    version = "v0.7.1",
)

go_repository(
    name = "com_github_hashicorp_consul",
    importpath = "github.com/hashicorp/consul",
    sum = "h1:1eDpXAxTh0iPv+1kc9/gfSI2pxRERDsTk/lNGolwHn8=",
    version = "v0.0.0-20171026175957-610f3c86a089",
)

go_repository(
    name = "com_github_hashicorp_errwrap",
    importpath = "github.com/hashicorp/errwrap",
    sum = "h1:OxrOeh75EUXMY8TBjag2fzXGZ40LB6IKw45YeGUDY2I=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_hashicorp_go_azure_helpers",
    importpath = "github.com/hashicorp/go-azure-helpers",
    sum = "h1:larj4ZgwO3hKzA9xIOTXRW4NBpI6F3K8wpig8eikNOw=",
    version = "v0.43.0",
)

go_repository(
    name = "com_github_hashicorp_go_checkpoint",
    importpath = "github.com/hashicorp/go-checkpoint",
    sum = "h1:MFYpPZCnQqQTE18jFwSII6eUQrD/oxMFp3mlgcqk5mU=",
    version = "v0.5.0",
)

go_repository(
    name = "com_github_hashicorp_go_getter",
    importpath = "github.com/hashicorp/go-getter",
    sum = "h1:uJDtyXwEfalmp1PqdxuhZqrNkUyClZAhVeZYTArbqkg=",
    version = "v1.7.2",
)

go_repository(
    name = "com_github_hashicorp_go_msgpack",
    importpath = "github.com/hashicorp/go-msgpack",
    sum = "h1:SFT72YqIkOcLdWJUYcriVX7hbrZpwc/f7h8aW2NUqrA=",
    version = "v0.5.4",
)

go_repository(
    name = "com_github_hashicorp_go_multierror",
    importpath = "github.com/hashicorp/go-multierror",
    sum = "h1:H5DkEtf6CXdFp0N0Em5UCwQpXMWke8IA0+lD48awMYo=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_hashicorp_go_plugin",
    importpath = "github.com/hashicorp/go-plugin",
    sum = "h1:DXmvivbWD5qdiBts9TpBC7BYL1Aia5sxbRgQB+v6UZM=",
    version = "v1.4.3",
)

go_repository(
    name = "com_github_hashicorp_go_retryablehttp",
    importpath = "github.com/hashicorp/go-retryablehttp",
    sum = "h1:ZQgVdpTdAL7WpMIwLzCfbalOcSUdkDZnpUv3/+BxzFA=",
    version = "v0.7.4",
)

go_repository(
    name = "com_github_hashicorp_go_safetemp",
    importpath = "github.com/hashicorp/go-safetemp",
    sum = "h1:2HR189eFNrjHQyENnQMMpCiBAsRxzbTMIgBhEyExpmo=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_slug",
    importpath = "github.com/hashicorp/go-slug",
    sum = "h1:c6lLdQnlhUWbS5I7hw8SvfymoFuy6EmiFDedy6ir994=",
    version = "v0.11.1",
)

go_repository(
    name = "com_github_hashicorp_go_sockaddr",
    importpath = "github.com/hashicorp/go-sockaddr",
    sum = "h1:ztczhD1jLxIRjVejw8gFomI1BQZOe2WoVOu0SyteCQc=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_hashicorp_go_tfe",
    importpath = "github.com/hashicorp/go-tfe",
    sum = "h1:aacguqCENg6Z7ttfhAxdbbY2vm/jKrntl5sUUY0h6EM=",
    version = "v1.26.0",
)

go_repository(
    name = "com_github_hashicorp_go_uuid",
    importpath = "github.com/hashicorp/go-uuid",
    sum = "h1:2gKiV6YVmrJ1i2CKKa9obLvRieoRGviZFL26PcT/Co8=",
    version = "v1.0.3",
)

go_repository(
    name = "com_github_hashicorp_go_version",
    importpath = "github.com/hashicorp/go-version",
    sum = "h1:feTTfFNnjP967rlCxM/I9g701jU+RN74YKx2mOkIeek=",
    version = "v1.6.0",
)

go_repository(
    name = "com_github_hashicorp_hcl_v2",
    importpath = "github.com/hashicorp/hcl/v2",
    sum = "h1:z1XvSUyXd1HP10U4lrLg5e0JMVz6CPaJvAgxM0KNZVY=",
    version = "v2.17.0",
)

go_repository(
    name = "com_github_hashicorp_hil",
    importpath = "github.com/hashicorp/hil",
    sum = "h1:ZffYoGCZ7E9PuawPRlSyMdC9xDMmKVB8LUzCikoWAyQ=",
    version = "v0.0.0-20230714223007-037c8ce1181b",
)

go_repository(
    name = "com_github_hashicorp_memberlist",
    importpath = "github.com/hashicorp/memberlist",
    sum = "h1:EtYPN8DpAURiapus508I4n9CzHs2W+8NZGbmmR/prTM=",
    version = "v0.5.0",
)

go_repository(
    name = "com_github_hashicorp_terraform",
    importpath = "github.com/hashicorp/terraform",
    sum = "h1:2QWbTj2xJ/8W1gCyIrd0WAqVF4weKPMYjx8nKjbkQjA=",
    version = "v0.15.3",
)

go_repository(
    name = "com_github_hashicorp_terraform_config_inspect",
    importpath = "github.com/hashicorp/terraform-config-inspect",
    sum = "h1:1uPcotqoL4TjcGKlgIe7OFSRplf7BMVtUjekwmCrvuM=",
    version = "v0.0.0-20230614215431-f32df32a01cd",
)

go_repository(
    name = "com_github_hashicorp_terraform_svchost",
    importpath = "github.com/hashicorp/terraform-svchost",
    sum = "h1:EZZimZ1GxdqFRinZ1tpJwVxxt49xc/S52uzrw4x0jKQ=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_hashicorp_vault",
    importpath = "github.com/hashicorp/vault",
    sum = "h1:4x0lHxui/ZRp/B3E0Auv1QNBJpzETqHR2kQD3mHSBJU=",
    version = "v0.10.4",
)

go_repository(
    name = "com_github_hashicorp_yamux",
    importpath = "github.com/hashicorp/yamux",
    sum = "h1:kJCB4vdITiW1eC1vq2e6IsrXKrZit1bv/TDYFGMp4BQ=",
    version = "v0.0.0-20181012175058-2f1d1f20f75d",
)

go_repository(
    name = "com_github_hpcloud_tail",
    importpath = "github.com/hpcloud/tail",
    sum = "h1:nfCOvKYfkgYP8hkirhJocXT2+zOD8yUNjXaWfTlyFKI=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_jhump_protoreflect",
    importpath = "github.com/jhump/protoreflect",
    sum = "h1:h5jfMVslIg6l29nsMs0D8Wj17RDVdNYti0vDN/PZZoE=",
    version = "v1.6.0",
)

go_repository(
    name = "com_github_jonboulle_clockwork",
    importpath = "github.com/jonboulle/clockwork",
    sum = "h1:VKV+ZcuP6l3yW9doeqz6ziZGgcynBVQO+obU0+0hcPo=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_joyent_triton_go",
    importpath = "github.com/joyent/triton-go",
    sum = "h1:kie3qOosvRKqwij2HGzXWffwpXvcqfPPXRUw8I4F/mg=",
    version = "v0.0.0-20180313100802-d8f9c0314926",
)

go_repository(
    name = "com_github_jtolds_gls",
    importpath = "github.com/jtolds/gls",
    sum = "h1:fSuqC+Gmlu6l/ZYAoZzx2pyucC8Xza35fpRVWLVmUEE=",
    version = "v4.2.1+incompatible",
)

go_repository(
    name = "com_github_julienschmidt_httprouter",
    importpath = "github.com/julienschmidt/httprouter",
    sum = "h1:TDTW5Yz1mjftljbcKqRcrYhd4XeOoI98t+9HbQbYf7g=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_kardianos_osext",
    importpath = "github.com/kardianos/osext",
    sum = "h1:iQTw/8FWTuc7uiaSepXwyf3o52HaUYcV+Tu66S3F5GA=",
    version = "v0.0.0-20190222173326-2bc1f35cddc0",
)

go_repository(
    name = "com_github_keybase_go_crypto",
    importpath = "github.com/keybase/go-crypto",
    sum = "h1:NARVGAAgEXvoMeNPHhPFt1SBt1VMznA3Gnz9d0qj+co=",
    version = "v0.0.0-20161004153544-93f5b35093ba",
)

go_repository(
    name = "com_github_klauspost_compress",
    importpath = "github.com/klauspost/compress",
    sum = "h1:Lcadnb3RKGin4FYM/orgq0qde+nc15E5Cbqg4B9Sx9c=",
    version = "v1.15.11",
)

go_repository(
    name = "com_github_konsorten_go_windows_terminal_sequences",
    importpath = "github.com/konsorten/go-windows-terminal-sequences",
    sum = "h1:mweAR1A6xJ3oS2pRaGiHgQ4OO8tzTaLawm8vnODuwDk=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_kr_logfmt",
    importpath = "github.com/kr/logfmt",
    sum = "h1:T+h1c/A9Gawja4Y9mFVWj2vyii2bbUNDw3kt9VxK2EY=",
    version = "v0.0.0-20140226030751-b84e30acd515",
)

go_repository(
    name = "com_github_kylelemons_godebug",
    importpath = "github.com/kylelemons/godebug",
    sum = "h1:RPNrshWIDI6G2gRW9EHilWtl7Z6Sb1BR0xunSBf0SNc=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_lib_pq",
    importpath = "github.com/lib/pq",
    sum = "h1:v9QZf2Sn6AmjXtQeFpdoq/eaNtYP6IN+7lcrygsIAtg=",
    version = "v1.10.3",
)

go_repository(
    name = "com_github_likexian_gokit",
    importpath = "github.com/likexian/gokit",
    sum = "h1:DgtIqqTRFqtbiLJFzuRESwVrxWxfs8OlY6hnPYBa3BM=",
    version = "v0.20.15",
)

go_repository(
    name = "com_github_likexian_simplejson_go",
    importpath = "github.com/likexian/simplejson-go",
    sum = "h1:tNa5q0zLsg6EzUI4npXeUYwk/xuTtMcKdVnRsYdOHzk=",
    version = "v0.0.0-20190502021454-d8787b4bfa0b",
)

go_repository(
    name = "com_github_lusis_go_artifactory",
    importpath = "github.com/lusis/go-artifactory",
    sum = "h1:wnfcqULT+N2seWf6y4yHzmi7GD2kNx4Ute0qArktD48=",
    version = "v0.0.0-20160115162124-7e4ce345df82",
)

go_repository(
    name = "com_github_masterzen_simplexml",
    importpath = "github.com/masterzen/simplexml",
    sum = "h1:2ZKn+w/BJeL43sCxI2jhPLRv73oVVOjEKZjKkflyqxg=",
    version = "v0.0.0-20190410153822-31eea3082786",
)

go_repository(
    name = "com_github_masterzen_winrm",
    importpath = "github.com/masterzen/winrm",
    sum = "h1:cxuVcCvCLD9yYDbRCWw0jSgh1oT6P6mv3aJDKK5o7X4=",
    version = "v0.0.0-20200615185753-c42b5136ff88",
)

go_repository(
    name = "com_github_mattn_go_shellwords",
    importpath = "github.com/mattn/go-shellwords",
    sum = "h1:xmZZyxuP+bYKAKkA9ABYXVNJ+G/Wf3R8d8vAP3LDJJk=",
    version = "v1.0.4",
)

go_repository(
    name = "com_github_matttproud_golang_protobuf_extensions",
    importpath = "github.com/matttproud/golang_protobuf_extensions",
    sum = "h1:4hp9jkHxhMHkqkrB3Ix0jegS5sx/RkqARlsWZ6pIwiU=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_miekg_dns",
    importpath = "github.com/miekg/dns",
    sum = "h1:WMszZWJG0XmzbK9FEmzH2TVcqYzFesusSIB41b8KHxY=",
    version = "v1.1.41",
)

go_repository(
    name = "com_github_mitchellh_cli",
    importpath = "github.com/mitchellh/cli",
    sum = "h1:OxRIeJXpAMztws/XHlN2vu6imG5Dpq+j61AzAX5fLng=",
    version = "v1.1.5",
)

go_repository(
    name = "com_github_mitchellh_colorstring",
    importpath = "github.com/mitchellh/colorstring",
    sum = "h1:62I3jR2EmQ4l5rM/4FEfDWcRD+abF5XlKShorW5LRoQ=",
    version = "v0.0.0-20190213212951-d06e56a500db",
)

go_repository(
    name = "com_github_mitchellh_copystructure",
    importpath = "github.com/mitchellh/copystructure",
    sum = "h1:vpKXTN4ewci03Vljg/q9QvCGUDttBOGBIa15WveJJGw=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_mitchellh_go_linereader",
    importpath = "github.com/mitchellh/go-linereader",
    sum = "h1:GRiLv4rgyqjqzxbhJke65IYUf4NCOOvrPOJbV/sPxkM=",
    version = "v0.0.0-20190213213312-1b945b3263eb",
)

go_repository(
    name = "com_github_mitchellh_go_testing_interface",
    importpath = "github.com/mitchellh/go-testing-interface",
    sum = "h1:jrgshOhYAUVNMAJiKbEu7EqAwgJJ2JqpQmpLJOu07cU=",
    version = "v1.14.1",
)

go_repository(
    name = "com_github_mitchellh_go_wordwrap",
    importpath = "github.com/mitchellh/go-wordwrap",
    sum = "h1:TLuKupo69TCn6TQSyGxwI1EblZZEsQ0vMlAFQflz0v0=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_mitchellh_hashstructure",
    importpath = "github.com/mitchellh/hashstructure",
    sum = "h1:P6P1hdjqAAknpY/M1CGipelZgp+4y9ja9kmUZPXP+H0=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_mitchellh_panicwrap",
    importpath = "github.com/mitchellh/panicwrap",
    sum = "h1:67zIyVakCIvcs69A0FGfZjBdPleaonSgGlXRSRlb6fE=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_mitchellh_prefixedio",
    importpath = "github.com/mitchellh/prefixedio",
    sum = "h1:eD92Am0Qf3rqhsOeA1zwBHSfRkoHrt4o6uORamdmJP8=",
    version = "v0.0.0-20190213213902-5733675afd51",
)

go_repository(
    name = "com_github_mitchellh_reflectwalk",
    importpath = "github.com/mitchellh/reflectwalk",
    sum = "h1:G2LzWKi524PWgd3mLHV8Y5k7s6XUvT0Gef6zxSIeXaQ=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_mozillazg_go_httpheader",
    importpath = "github.com/mozillazg/go-httpheader",
    sum = "h1:3brX5z8HTH+0RrNA1362Rc3HsaxyWEKtGY45YrhuINM=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_mwitkow_go_conntrack",
    importpath = "github.com/mwitkow/go-conntrack",
    sum = "h1:F9x/1yl3T2AeKLr2AMdilSD8+f9bvMnNN8VS5iDtovc=",
    version = "v0.0.0-20161129095857-cc309e4a2223",
)

go_repository(
    name = "com_github_nu7hatch_gouuid",
    importpath = "github.com/nu7hatch/gouuid",
    sum = "h1:VhgPp6v9qf9Agr/56bj7Y/xa04UccTW04VP0Qed4vnQ=",
    version = "v0.0.0-20131221200532-179d4d0c4d8d",
)

go_repository(
    name = "com_github_oklog_run",
    importpath = "github.com/oklog/run",
    sum = "h1:Ru7dDtJNOyC66gQ5dQmaCa0qIsAUFY3sFpK1Xk8igrw=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_onsi_ginkgo",
    importpath = "github.com/onsi/ginkgo",
    sum = "h1:VkHVNpR4iVnU8XQR6DBm8BqYjN7CRzw+xKUbVVbbW9w=",
    version = "v1.8.0",
)

go_repository(
    name = "com_github_packer_community_winrmcp",
    importpath = "github.com/packer-community/winrmcp",
    sum = "h1:9uViuKtx1jrlXLBW/pMnhOfzn3iSEdLase/But/IZRU=",
    version = "v0.0.0-20180921211025-c76d91c1e7db",
)

go_repository(
    name = "com_github_pascaldekloe_goe",
    importpath = "github.com/pascaldekloe/goe",
    sum = "h1:cBOtyMzM9HTpWjXfbbunk26uA6nG3a8n06Wieeh0MwY=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_pkg_browser",
    importpath = "github.com/pkg/browser",
    sum = "h1:dofHuld+js7eKSemxqTVIo8yRlpRw+H1SdpzZxWruBc=",
    version = "v0.0.0-20201207095918-0426ae3fba23",
)

go_repository(
    name = "com_github_posener_complete",
    importpath = "github.com/posener/complete",
    sum = "h1:NP0eAhjcjImqslEwo/1hq7gpajME0fTLTezBKDqfXqo=",
    version = "v1.2.3",
)

go_repository(
    name = "com_github_prometheus_client_golang",
    importpath = "github.com/prometheus/client_golang",
    sum = "h1:YVIb/fVcOTMSqtqZWSKnHpSLBxu8DKgxq8z6RuBZwqI=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_prometheus_common",
    importpath = "github.com/prometheus/common",
    sum = "h1:KOMtN28tlbam3/7ZKEYKHhKoJZYYj3gMH4uc62x7X7U=",
    version = "v0.9.1",
)

go_repository(
    name = "com_github_prometheus_procfs",
    importpath = "github.com/prometheus/procfs",
    sum = "h1:+fpWZdT24pJBiqJdAwYBjPSk+5YmQzYNPYzQsdzLkt8=",
    version = "v0.0.8",
)

go_repository(
    name = "com_github_qcloudapi_qcloud_sign_golang",
    importpath = "github.com/QcloudApi/qcloud_sign_golang",
    sum = "h1:DTQ/38ao/CfXsrK0cSAL+h4R/u0VVvfWLZEOlLwEROI=",
    version = "v0.0.0-20141224014652-e4130a326409",
)

go_repository(
    name = "com_github_satori_go_uuid",
    importpath = "github.com/satori/go.uuid",
    sum = "h1:0uYX9dsZ2yD7q2RtLRtPSdGDWzjeM3TbMJP9utgA0ww=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_sean_seed",
    importpath = "github.com/sean-/seed",
    sum = "h1:nn5Wsu0esKSJiIVhscUtVbo7ada43DJhG55ua/hjS5I=",
    version = "v0.0.0-20170313163322-e2103e2c3529",
)

go_repository(
    name = "com_github_sergi_go_diff",
    importpath = "github.com/sergi/go-diff",
    sum = "h1:Kpca3qRNrduNnOQeazBd0ysaKrUJiIuISHxogkT9RPQ=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_smartystreets_assertions",
    importpath = "github.com/smartystreets/assertions",
    sum = "h1:zE9ykElWQ6/NYmHa3jpm/yHnI4xSofP+UP6SpjHcSeM=",
    version = "v0.0.0-20180927180507-b2de0cb4f26d",
)

go_repository(
    name = "com_github_smartystreets_goconvey",
    importpath = "github.com/smartystreets/goconvey",
    sum = "h1:JSvGDIbmil4Ui/dDdFBExb7/cmkNjyX5F97oglmvCDo=",
    version = "v0.0.0-20180222194500-ef6db91d284a",
)

go_repository(
    name = "com_github_soheilhy_cmux",
    importpath = "github.com/soheilhy/cmux",
    sum = "h1:0HKaf1o97UwFjHH9o5XsHUOF+tqmdA7KEzXLpiyaw0E=",
    version = "v0.1.4",
)

go_repository(
    name = "com_github_svanharmelen_jsonapi",
    importpath = "github.com/svanharmelen/jsonapi",
    sum = "h1:Z4EH+5EffvBEhh37F0C0DnpklTMh00JOkjW5zK3ofBI=",
    version = "v0.0.0-20180618144545-0c0828c3f16d",
)

go_repository(
    name = "com_github_tencentcloud_tencentcloud_sdk_go",
    importpath = "github.com/tencentcloud/tencentcloud-sdk-go",
    sum = "h1:5Td2b0yfaOvw9M9nZ5Oav6Li9bxUNxt4DgxMfIPpsa0=",
    version = "v3.0.82+incompatible",
)

go_repository(
    name = "com_github_tencentyun_cos_go_sdk_v5",
    importpath = "github.com/tencentyun/cos-go-sdk-v5",
    sum = "h1:uwRBzc70Wgtc5iQQCowqecfRT0OpCXUOZzodZHOOEDs=",
    version = "v0.7.29",
)

go_repository(
    name = "com_github_tmc_grpc_websocket_proxy",
    importpath = "github.com/tmc/grpc-websocket-proxy",
    sum = "h1:lYIiVDtZnyTWlNwiAxLj0bbpTcx1BWCFhXjfsvmPdNc=",
    version = "v0.0.0-20171017195756-830351dc03c6",
)

go_repository(
    name = "com_github_tombuildsstuff_giovanni",
    importpath = "github.com/tombuildsstuff/giovanni",
    sum = "h1:CVRaLOJ7C/eercCrKIsarfJ4SZoGMdBL9Q2deFDUXco=",
    version = "v0.15.1",
)

go_repository(
    name = "com_github_ugorji_go",
    importpath = "github.com/ugorji/go",
    sum = "h1:cMjKdf4PxEBN9K5HaD9UMW8gkTbM0kMzkTa9SJe0WNQ=",
    version = "v0.0.0-20180813092308-00b869d2f4a5",
)

go_repository(
    name = "com_github_ulikunitz_xz",
    importpath = "github.com/ulikunitz/xz",
    sum = "h1:t92gobL9l3HE202wg3rlk19F6X+JOxl9BBrCCMYEYd8=",
    version = "v0.5.10",
)

go_repository(
    name = "com_github_vmihailenco_msgpack",
    importpath = "github.com/vmihailenco/msgpack",
    sum = "h1:dSLoQfGFAo3F6OoNhwUmLwVgaUXK79GlxNBwueZn0xI=",
    version = "v4.0.4+incompatible",
)

go_repository(
    name = "com_github_vmihailenco_msgpack_v5",
    importpath = "github.com/vmihailenco/msgpack/v5",
    sum = "h1:5gO0H1iULLWGhs2H5tbAHIZTV8/cYafcFOr9znI5mJU=",
    version = "v5.3.5",
)

go_repository(
    name = "com_github_vmihailenco_tagparser_v2",
    importpath = "github.com/vmihailenco/tagparser/v2",
    sum = "h1:y09buUbR+b5aycVFQs/g70pqKVZNBmxwAhO7/IwNM9g=",
    version = "v2.0.0",
)

go_repository(
    name = "com_github_xanzy_ssh_agent",
    importpath = "github.com/xanzy/ssh-agent",
    sum = "h1:AmzO1SSWxw73zxFZPRwaMN1MohDw8UyHnmuxyceTEGo=",
    version = "v0.3.1",
)

go_repository(
    name = "com_github_xiang90_probing",
    importpath = "github.com/xiang90/probing",
    sum = "h1:MPPkRncZLN9Kh4MEFmbnK4h3BD7AUmskWv2+EeZJCCs=",
    version = "v0.0.0-20160813154853-07dd2e8dfe18",
)

go_repository(
    name = "com_github_xlab_treeprint",
    importpath = "github.com/xlab/treeprint",
    sum = "h1:Jpn2j6wHkC9wJv5iMfJhKqrZJx3TahFx+7sbZ7zQdxs=",
    version = "v0.0.0-20161029104018-1d6e34225557",
)

go_repository(
    name = "com_github_zclconf_go_cty",
    importpath = "github.com/zclconf/go-cty",
    sum = "h1:4GvrUxe/QUDYuJKAav4EYqdM47/kZa672LwmXFmEKT0=",
    version = "v1.13.2",
)

go_repository(
    name = "com_github_zclconf_go_cty_debug",
    importpath = "github.com/zclconf/go-cty-debug",
    sum = "h1:FosyBZYxY34Wul7O/MSKey3txpPYyCqVO5ZyceuQJEI=",
    version = "v0.0.0-20191215020915-b22d67c1ba0b",
)

go_repository(
    name = "com_github_zclconf_go_cty_yaml",
    importpath = "github.com/zclconf/go-cty-yaml",
    sum = "h1:og/eOQ7lvA/WWhHGFETVWNduJM7Rjsv2RRpx1sdFMLc=",
    version = "v1.0.3",
)

go_repository(
    name = "com_google_cloud_go_grafeas",
    importpath = "cloud.google.com/go/grafeas",
    sum = "h1:CYjC+xzdPvbV65gi6Dr4YowKcmLo045pm18L0DhdELM=",
    version = "v0.2.0",
)

go_repository(
    name = "com_google_cloud_go_recaptchaenterprise",
    importpath = "cloud.google.com/go/recaptchaenterprise",
    sum = "h1:u6EznTGzIdsyOsvm+Xkw0aSuKFXQlyjGE9a4exk6iNQ=",
    version = "v1.3.1",
)

go_repository(
    name = "com_google_cloud_go_vision",
    importpath = "cloud.google.com/go/vision",
    sum = "h1:/CsSTkbmO9HC8iQpxbK8ATms3OQaX3YQUeTMGCxlaK4=",
    version = "v1.2.0",
)

go_repository(
    name = "in_gopkg_alecthomas_kingpin_v2",
    importpath = "gopkg.in/alecthomas/kingpin.v2",
    sum = "h1:jMFz6MfLP0/4fUyZle81rXUoxOBFi19VUFKVDOQfozc=",
    version = "v2.2.6",
)

go_repository(
    name = "in_gopkg_cheggaaa_pb_v1",
    importpath = "gopkg.in/cheggaaa/pb.v1",
    sum = "h1:kJdccidYzt3CaHD1crCFTS1hxyhSi059NhOFUf03YFo=",
    version = "v1.0.27",
)

go_repository(
    name = "in_gopkg_fsnotify_v1",
    importpath = "gopkg.in/fsnotify.v1",
    sum = "h1:xOHLXZwVvI9hhs+cLKq5+I5onOuwQLhQwiu63xxlHs4=",
    version = "v1.4.7",
)

go_repository(
    name = "in_gopkg_resty_v1",
    importpath = "gopkg.in/resty.v1",
    sum = "h1:CuXP0Pjfw9rOuY6EP+UvtNvt5DSqHpIxILZKT/quCZI=",
    version = "v1.12.0",
)

go_repository(
    name = "in_gopkg_tomb_v1",
    importpath = "gopkg.in/tomb.v1",
    sum = "h1:uRGJdciOHaEIrze2W8Q3AKkepLTh2hOroT7a+7czfdQ=",
    version = "v1.0.0-20141024135613-dd632973f1e7",
)

go_repository(
    name = "org_golang_google_grpc_cmd_protoc_gen_go_grpc",
    importpath = "google.golang.org/grpc/cmd/protoc-gen-go-grpc",
    sum = "h1:M1YKkFIboKNieVO5DLUEVzQfGwJD30Nv2jfUgzb5UcE=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_sashabaranov_go_openai",
    importpath = "github.com/sashabaranov/go-openai",
    sum = "h1:5DPTtR9JBjKPJS008/A409I5ntFhUPPGCmaAihcPRyo=",
    version = "v1.14.2",
)

go_repository(
    name = "com_github_atotto_clipboard",
    importpath = "github.com/atotto/clipboard",
    sum = "h1:EH0zSVneZPSuFR11BlR9YppQTVDbh5+16AmcJi4g1z4=",
    version = "v0.1.4",
)

go_repository(
    name = "com_github_aymanbagabas_go_osc52_v2",
    importpath = "github.com/aymanbagabas/go-osc52/v2",
    sum = "h1:HwpRHbFMcZLEVr42D4p7XBqjyuxQH5SMiErDT4WkJ2k=",
    version = "v2.0.1",
)

go_repository(
    name = "com_github_charmbracelet_bubbles",
    importpath = "github.com/charmbracelet/bubbles",
    sum = "h1:6uzpAAaT9ZqKssntbvZMlksWHruQLNxg49H5WdeuYSY=",
    version = "v0.16.1",
)

go_repository(
    name = "com_github_charmbracelet_harmonica",
    importpath = "github.com/charmbracelet/harmonica",
    sum = "h1:8NxJWRWg/bzKqqEaaeFNipOu77YR5t8aSwG4pgaUBiQ=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_charmbracelet_lipgloss",
    importpath = "github.com/charmbracelet/lipgloss",
    sum = "h1:IS00fk4XAHcf8uZKc3eHeMUTCxUH6NkaTrdyCQk84RU=",
    version = "v0.8.0",
)

go_repository(
    name = "com_github_dustin_go_humanize",
    importpath = "github.com/dustin/go-humanize",
    sum = "h1:GzkhY7T5VNhEkwH0PVJgjz+fX1rhBrR7pRT3mDkpeCY=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_sahilm_fuzzy",
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
