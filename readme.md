# Finch

Finch, a simple neural network framework and visualizer that uses genetic networks to train.

## Nix

The use of [Nix](https://nixos.org) is not required but strongly recommended, installation guide available [here](https://github.com/determinatesystems/nix-installer).

```shell
nix flake clone github:stepbrobd/finch --dest finch && cd finch
```

To start a dev shell:

```shell
nix develop .
```

To run directly:

```shell
nix run . -- -help
```

To build:

```shell
nix build .
```

Output binary will located at `./result/bin/finch`.

## License

The contents inside this repository, excluding all submodules, are licensed under the [MIT License](license.md).
Third-party file(s) and/or code(s) are subject to their original term(s) and/or license(s).
