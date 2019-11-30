# <img height=150 src="./images/dkr.png"/>

[![Go Report Card](https://goreportcard.com/badge/github.com/Roverr/dkr)](https://goreportcard.com/report/github.com/Roverr/dkr)
 [![Maintainability](https://api.codeclimate.com/v1/badges/202152e83296250ab527/maintainability)](https://codeclimate.com/github/Roverr/rtsp-stream/maintainability)
 [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
 ![GitHub last commit](https://img.shields.io/github/last-commit/Roverr/dkr.svg)
 ![GitHub release](https://img.shields.io/github/release/Roverr/dkr.svg)
 
Light CLI application to make it easier for developers to interact with docker containers.

It can be utilisied while working regularly with containers as a developer. Provides a nice CLI user interface to do basics.

<p align="center">
    <img src="images/summary.gif" width="100%"/>
</p>


## Table of contents
* [How to use](#how-does-it-work)
* [Install](#install)
    * [Linux](#linux)
    * [OSx](#osx)
* [Improvements](#improvements)
* [Credits](#credits)

## How to use

dkr is a really simple CLI application to help you with your everyday tasks around docker containers.<br/>
It is written in Go to ensure proper distribution for the main developer platforms.

It does not try to be an all-around solution for providing CLI interface over docker functions. The reason for this is that there are a lot of docker commands that do not need any simplification because they are not used regularly or just plain simple to use. Consider this application as a shortcut for doing something that you can do anyway but takes longer time. In the long run the plan is to eliminate use cases that are pretty similiar to the current functionality.

On select screens you can use your arrow keys and enter to navigate however this is also stated everytime you encounter a select situation.

<p align="center">
    <img src="images/main.gif" width="100%"/>
</p>

dkr implements 3 commands right now:
* **exec** - For times when you have to enter the container and manually check things
* **logs** - When you want to see the logs of the container
* **stop** - To stop the container

The reason for having this 3 is that in my experience most of the time when I use `docker ps` I want to do one of these commands. As a common user I think an interactive version is better.

**Two potential way to use the application:**
* `dkr`
    * This will lead you to the list of containers where you can select the container and action
* `dkr logs` / `dkr exec` / `dkr stop`
    * Takes you to the list of containers where you can choose which one should be the target of the given action command

<p align="center">
    <img src="images/commands.gif" width="100%"/>
</p>

## Install

**Docker has to be installed beforehand** 


* Manual installation
    * clone the repository under `github.com/Roverr/dkr`
    * `dep ensure` (install dep before)
    * `go build .` (install go before)
    * Move the binary into your path
* Choose a binary packed under [the latest release](https://github.com/Roverr/dkr/releases)
* Use the following scripts below

### Linux

```s
```

### OSx

```s
brew tap roverr/opensource
brew install dkr
```

## Improvements

Have a use case similiar to what this application is made for?<br/>
[Open an issue](https://github.com/Roverr/dkr/issues/new) describing it!

## Credits

Big thanks to [MariaLetta](https://github.com/MariaLetta/free-gophers-pack) for the Gopher in the headline!
