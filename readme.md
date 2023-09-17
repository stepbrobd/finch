# Finch

Finch, a simple neural net framework and visualizer that uses genetic networks to train.

## Running Finch on example datasets

**OR, NOR, XOR**:

`-input=2 -output=1 -hidden=2`

2 input neurons, 1 output neuron, 1 hidden layer with two neurons

`-population=128 -mutation=0.025`

128 individules in the population, with 2.5% mutation rate

`-example=./data/gates/input_data.csv -expected=./data/gates/{or,nor,xor}_label_data.csv`

dataset paths

```shell
finch -input=2 -output=1 -hidden=2 -population=128 -mutation=0.025 -example=./data/gates/input_data.csv -expected=./data/gates/{or,nor,xor}_label_data.csv
```

Remember to change which operation you want to run: OR, NOR, XOR.

**MNIST**:

`-input=784 -output=10 -hidden=16,16`

784 input neurons (28x28 greyscale images), 10 output neurons (numbers 1-10), 2 hidden layers with 16 neurons each

`-population=4096 -mutation=0.1`

4096 individules in the population, with 10% mutation rate

`-example=./data/mnist/mnist_pixel_data_{32,64,128,256,512,1024,2048,4096,8192}.csv -expected=./data/mnist/mnist_label_data_{32,64,128,256,512,1024,2048,4096,8192}.csv`

dataset paths

```shell
finch -input=784 -output=10 -hidden=16,16 -population=4096 -mutation=0.1 -example=./data/mnist/mnist_pixel_data_{32,64,128,256,512,1024,2048,4096,8192}.csv -expected=./data/mnist/mnist_label_data_{32,64,128,256,512,1024,2048,4096,8192}.csv
```

Remember to change the size of MNIST dataset: 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192.

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
