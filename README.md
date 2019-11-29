# <img src="./images/dkr.svg"/>

[![Go Report Card](https://goreportcard.com/badge/github.com/Roverr/dkr)](https://goreportcard.com/report/github.com/Roverr/dkr)
 [![Maintainability](https://api.codeclimate.com/v1/badges/202152e83296250ab527/maintainability)](https://codeclimate.com/github/Roverr/rtsp-stream/maintainability)
 [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
 ![GitHub last commit](https://img.shields.io/github/last-commit/Roverr/dkr.svg)
 ![GitHub release](https://img.shields.io/github/release/Roverr/dkr.svg)
 
Light CLI application to make it easier for developers to interact with docker containers.

Provides an interactive interface over couple of docker commands so it can be easily utilised while working with docker containers a lot and you prefer something simple and light.

## Table of contents
* [How to use](#how-does-it-work)
* [Install](#install)
    * [Linux](#linux)
    * [OSx](#osx)
* [Credits](#credits)

## How to use

dkr is a really simple CLI application to help you with your everyday tasks around docker container. It is written in Go to ensure proper distribution for the main developer platforms.

<p align="center">
    <img src="images/main.gif" width="100%"/>
</p>

dkr implements 3 commands right now:
* exec - For times when you have to enter the container and manually check things
* logs - When you want to see the logs of the container
* stop - To stop the container

The reason for having this 3 is that in my experience most of the time when I use `docker ps` I want to do one of these commands. As a common user I think an interactive version is better.

You can use it in 2 ways:
* You can simply type `dkr` - This will lead you to the containers where you will be selecting what to do
* You can also type the commands mentioned above like `dkr logs` - This way you will only be asked to choose a container

<p align="center">
    <img src="images/commands.gif" width="100%"/>
</p>

## Credits

Big thanks to [MariaLetta](https://github.com/MariaLetta/free-gophers-pack) for the Gopher in the headline!
