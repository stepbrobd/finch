{
  inputs = {
    nixpkgs.url = "flake:nixpkgs/nixpkgs-unstable";
    flake-utils.url = "flake:flake-utils";
    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.flake-utils.follows = "flake-utils";
    };
  };

  outputs =
    { self
    , nixpkgs
    , flake-utils
    , gomod2nix
    , ...
    }: flake-utils.lib.eachDefaultSystem
      (system:
      let
        pkgs = import nixpkgs {
          inherit system;
          overlays = [
            gomod2nix.overlays.default
          ];
        };

        lib = pkgs.lib;

        finch = pkgs.buildGoApplication rec {
          name = "finch";
          src = lib.cleanSource ./.;
          modules = ./gomod2nix.toml;
        };
      in
      {
        packages.default = finch;

        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            go_1_21
            gopls
            delve
            gomod2nix.packages.${system}.default
          ];
        };

        formatter = pkgs.nixpkgs-fmt;
      }
      );
}
