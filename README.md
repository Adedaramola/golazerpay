# Lazerpay Golang SDK
A simple golang library for interacting with lazerpay's crypto payment gateway.

## Table of contents
- [Get started](#get-started)
- [Basic Usage](#basic-usage)
- [Payments](#payments)
- [Transfers](#transfers)
- [Swap](#swap)

## Get started
Although this is a simple api wrapper helper, it is advised you are familiar with the lazerpay official API [documentation](https://docs.lazerpay.finance)

### Install
```bash
go get github/adedaramola/golazerpay
```

## Basic usage
```golang
package main

import "github/adedaramola/golazerpay"

func main() {
    lazerpay := lazerpay.Setup(YOUR_SECRET_KEY, YOUR_PUBLIC_KEY)

    // Get all coins on your integration
    res, _ := lazerpay.Miscellaneous.GetCoins()
}
```