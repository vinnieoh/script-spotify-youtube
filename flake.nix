{
  description = "Script Spotify YouTube - Projeto Go com ambiente Nix";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs = { self, nixpkgs }:
    let
      system = "x86_64-linux";
      pkgs = import nixpkgs { inherit system; };
    in {
      devShells.${system}.default = pkgs.mkShell {
        buildInputs = [
          pkgs.go_1_23         # Go 1.23.x
          pkgs.golangci-lint    # Linter para Go
        ];

        shellHook = ''
          echo "🚀 Ambiente Nix para Script Spotify YouTube pronto!"
          echo "Go version: $(go version)"
          echo "GolangCI-Lint version: $(golangci-lint --version || true)"
        '';
      };
    };
}
