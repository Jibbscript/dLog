# dLog
Cloud-native distributed commit-log service
# Features
- Commit-log library package providing core functionality
- Client-Server communication over [gRPC](https://grpc.io/) (ProtoBuf binary payloads)
- Authentication w/ mTLS and self-operated CA w/ [CloudFlare CFSSL toolkit](https://github.com/cloudflare/cfssl) 
- Authorization thru simple ACL implementation w/ [casbin](https://github.com/casbin/casbin)
- Instrumented for observability (Metrics/Tracing w/ [OpenCensus](https://opencensus.io/); Logging w/ [Zap](https://github.com/uber-go/zap))
- Compatible w/ [Prometheus](https://prometheus.io/), [Jaegar](https://www.jaegertracing.io/), [Elasticsearch](https://www.elastic.co/) for Production deployment
# In Development
- Server2Server Service Discovery w/ [HashiCorp Serf](https://github.com/hashicorp/serf)
- Service clustering w/ consensus backed by [Raft](https://raft.github.io/) algorithm provided by [HashiCorp Raft library](https://github.com/hashicorp/raft)
- Client-side discovery and load-balancing (Preferable to proxy-based LB for higher perf communication between trusted internal-services)
- k8s-ready deployment Helm charts and Custom Controller logic w/ [MetaController](https://metacontroller.github.io/metacontroller/intro.html)
- [Operator-based](https://operatorhub.io/what-is-an-operator) deployment to k8s
# Further reading
- [Protocol Buffers](https://developers.google.com/protocol-buffers)
- [Memory-mapped files Go pkg](https://pkg.go.dev/github.com/tysontate/gommap)
- [CFSSL](https://blog.cloudflare.com/introducing-cfssl/)
- [gRPC server options](https://godoc.org/google.golang.org/grpc#ServerOption)
- [Serf](https://www.serf.io/)
- [distributed-services-with-go](https://www.oreilly.com/library/view/distributed-services-with/9781680508376/)
