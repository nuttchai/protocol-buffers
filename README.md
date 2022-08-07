# protocol-buffers

`protoc --cpp_out=. *.proto`

We can use above command to test our syntax which it is used to generate c++ code. If the command executes successfully, our syntax is correct

## Updating Pbs, the rules:

- DON'T change tags!
- Add new fields
- Use reserved tags
- Before changing type (int32 to int64) =>
  - check the documentation
  - or add a new field (preferred)

## Integer Types:

- int32 or int64
- signed or unsigned (uint32, uint64, sint32, sint64)
- Varint or not (fixed32, sfixed32: 4 bytes or fixed64, sfixed64: 8 bytes)

## Reference

- [Well Known Types](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#index)
- [Options](https://github.com/protocolbuffers/protobuf/blob/main/src/google/protobuf/descriptor.proto)
- [Naming Convention](https://developers.google.com/protocol-buffers/docs/style)
- [gRPC](https://grpc.io/)
- [Main repository examples](https://github.com/protocolbuffers/protobuf/tree/main/examples)
- [Some Google Apis types](https://github.com/googleapis/googleapis/tree/master/google/type)
- [Protocol Buffer itself](https://github.com/protocolbuffers/protobuf/tree/main/src/google/protobuf)
