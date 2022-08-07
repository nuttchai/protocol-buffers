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
