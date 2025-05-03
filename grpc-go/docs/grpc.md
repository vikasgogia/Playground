# Welcome to gRPC tutorial

## API (Application Programming Interface)

An API is a mechanism that enables two software components to communicate with each other using a set of definitions and protocols. gRPC and REST are two ways you can design an API.

- In gRPC, one component (the client) calls or invokes specific functions in another software component (the server).
- In REST, instead of calling functions, the client requests or updates data on the server. REST is primarily for CRUD operations on data.

## gRPC (Google Remote Procedure Call)

High-performance, open-source RPC framework that lets you define APIs as services, and clients call them like local functions — over HTTP/2.

- Supports streaming, bi-directional communication, and code generation.
- Uses Protocol Buffers (for serialization) and HTTP/2 for data transmission.

## How REST differs from gRPC?

| | gRPC | REST |
|----------|----------|----------|
| Design | Service-oriented design (may or may not impact server resources)  | Entity-oriented design (impacts server resources) |
| Protocol | HTTP 2 | HTTP 1.1 |
| Communication | Unary, one server to many clients, one client to many servers, and many clients to many servers | single client to single server (unary)|
| Data access | Service (function) calls | Multiple endpoints in the form of URLs to define resources |
| Data returned | As defined in Protobuf file | JSON |
| Bi-directional Streaming| Yes | No |
| Best suited for | High-performance or data-heavy microservice architectures | Simple data sources |
| Example | createNewOrder(customer_id, item_id, item_quantity) -> order_id | POST /orders headers (customer_id, item_id, item_quantity) -> order_id |

### REST interaction with gRPC

No, a traditional REST client cannot directly interact with a gRPC server because they use different protocols and data formats. To bridge the gap, a **gRPC-Gateway** can be used to translate RESTful JSON over HTTP into gRPC calls.
