1. `Goland` Import Proto
  * `File`->`Settings`->`Languages & Frameworks`->`Protocol Buffers`

2. plugins

  ```bash
  1. https://github.com/git-for-windows/git/releases/download/v2.40.1.windows.1/Git-2.40.1-64-bit.exe
  2. https://github.com/yoheimuta/protolint
  3. https://github.com/bufbuild/protoc-gen-validate/releases/tag/v1.0.1
  4. https://github.com/protocolbuffers/protobuf/releases
  5. https://github.com/pseudomuto/protoc-gen-doc
  6. npm install ts-protoc-gen -g #https://github.com/improbable-eng/ts-protoc-gen
  ```

```bash
# mac
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

```


## Windows编译注意事项 - 貌似只是windows上使用`Git for Windows`

1. 新版 validate_out 不支持 ./tmp 的形式
2. js_out 和 ts_out 不支持 source_relative
