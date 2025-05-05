# Protobuf (Protocol Buffer)

Protobuf is a language-agnostic, compact binary format developed by Google for serializing structured data.

- Used to define gRPC messages and services via .proto files.
- Much faster and smaller than JSON or XML.
- Comes with a compiler (protoc) that generates strongly typed code (e.g., Go structs).

```bash
protoc --go_out=. --go-grpc_out=. proto/note.proto
```
