# Price (WIP)

[![Go Reference](https://pkg.go.dev/badge/github.com/inuoshios/price-action.svg)](https://pkg.go.dev/github.com/inuoshios/price-action)
[![build](https://github.com/inuoshios/price-action/actions/workflows/release.yml/badge.svg)](https://github.com/inuoshios/price-action/actions/workflows/release.yml)

Price is a CLI application that helps you display real time prices of Crypto/NFTs directly in your terminal.

This project was implemented in Golang using the Cobra Framework

## Installation

Download the binary file from [here](https://github.com/inuoshios/price-action/releases)

### For Windows

- Add the installation folder to your environment variable
- Open your terminal (Tested on git bash, it works on other windows terminals too)
- Run the `price` command, it'll generate the list of available commands

```sh
Available Commands:
  crypto      Generates real-time information on different crypto prices 🚀
  nft         Generates over a thousand list of NFTs, and it is blazingly fast
  stocks      Generate details of a particular stock price based on the input given to it
```

The command below will genrate the list of NFTs - you can use both

```sh
price nft
```

```sh
price n
```

This command will generate the prices of Cryptocurrency

```sh
price crypto
```

```sh
price c
```

To get the Price of a specific stock on the stock market, the command below can be used to generate it

```sh
price stocks --abbreviation {stockabbreviation eg: GME TSLA MSFT}
```

```sh
price s -a {stockabbreviation eg: GME TSLA MSFT}
```

### Contribution

If you find an issue, or would like to submit an improvement to this project, please submit an issue [here](https://github.com/inuoshios/price-action/issues). If you would like to submit a Pull request, please reference your issue or the issue you intend to fix.

### License

This project is under the MIT License. See the [License](LICENSE) file for the full license text.
