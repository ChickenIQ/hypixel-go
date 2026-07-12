{
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs =
    { nixpkgs, ... }:
    let
      eachSystem =
        func:
        nixpkgs.lib.genAttrs nixpkgs.lib.systems.flakeExposed (
          system:
          let
            pkgs = nixpkgs.legacyPackages.${system};
            tools = with pkgs; [
              go
              protobuf
              protoc-gen-go
              protoc-gen-go-grpc
            ];
          in
          func pkgs tools
        );
    in
    {
      packages = eachSystem (
        pkgs: tools: {
          default = pkgs.writeShellApplication {
            name = "generate-proto";
            runtimeInputs = tools;

            text = ''
              shopt -s globstar nullglob

              files=(proto/**/*.proto)
              module="$(go list -m)"

              for file in "''${files[@]}"; do
                relative="''${file#proto/}"
                package="$module/$(dirname "$file")"

                protoc -I proto \
                  --go_out=proto \
                  --go_opt=paths=source_relative \
                  --go_opt="M$relative=$package" \
                  --go-grpc_out=proto \
                  --go-grpc_opt=paths=source_relative \
                  --go-grpc_opt="M$relative=$package" \
                  "$file"
              done
            '';
          };
        }
      );

      devShells = eachSystem (
        pkgs: tools: {
          default = pkgs.mkShell {
            packages = tools;
          };
        }
      );
    };
}
