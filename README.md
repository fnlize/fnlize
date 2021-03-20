<p align="center">
  <img src="https://docs.fission.io/images/logo.png" width="300" />
  <br>
  <h1 align="center">Fission: Serverless Functions for Kubernetes</h1>
</p>
<p align="center">
  <a href="https://github.com/fnlize/fnlize/actions?query=workflow%3A%22Fission+CI%22">
    <img src="https://github.com/fnlize/fnlize/workflows/Fission%20CI/badge.svg?branch=master" alt="Build Status" />
  </a>
  
  <a href="https://goreportcard.com/report/github.com/fnlize/fnlize">
    <img src="https://goreportcard.com/badge/github.com/fnlize/fnlize" alt="Go Report Card" />
  </a>
  <a href="https://codecov.io/gh/fission/fission">
    <img src="https://codecov.io/gh/fission/fission/branch/master/graph/badge.svg" alt="codecov" />
  </a>
  <br>
  <a href="http://fission.io">fission.io</a> | <a href="http://twitter.com/fissionio">@fissionio</a> | <a href="https://join.slack.com/t/fissionio/shared_invite/enQtOTI3NjgyMjE5NzE3LTllODJiODBmYTBiYWUwMWQxZWRhNDhiZDMyN2EyNjAzMTFiYjE2Nzc1NzE0MTU4ZTg2MzVjMDQ1NWY3MGJhZmE">Slack</a>
</p>

--------------

Fission is a fast serverless framework for Kubernetes with a focus on
developer productivity and high performance.

Fission operates on _just the code_: Docker and Kubernetes are
abstracted away under normal operation, though you can use both to
extend Fission if you want to.

Fission is extensible to any language; the core is written in Go, and
language-specific parts are isolated in something called
_environments_ (more below).  Fission currently supports NodeJS, Python, Ruby, Go, 
PHP, Bash, and any Linux executable, with more languages coming soon.

Table of Contents
=================

* [Fission: Serverless Functions for Kubernetes](#fission-serverless-functions-for-kubernetes)
  * [Performance: 100msec cold start](#performance-100msec-cold-start)
  * [Kubernetes is the right place for Serverless](#kubernetes-is-the-right-place-for-serverless)
  * [Getting Started](#getting-started)
  * [Learn More](#learn-more)
  * [Contributing](#contributing)
  * [Get Help &amp; Community Meeting](#get-help--community-meeting)
  * [Official Releases](#official-releases)
* [Licensing](#licensing)

## Performance: 100msec cold start

Fission maintains a pool of "warm" containers that each contain a
small dynamic loader.  When a function is first called,
i.e. "cold-started", a running container is chosen and the function is
loaded.  This pool is what makes Fission fast: cold-start latencies
are typically about 100msec.

## Kubernetes is the right place for Serverless

We're built on Kubernetes because we think any non-trivial app will
use a combination of serverless functions and more conventional
microservices, and Kubernetes is a great framework to bring these
together seamlessly.

Building on Kubernetes also means that anything you do for operations
on your Kubernetes cluster &mdash; such as monitoring or log
aggregation &mdash; also helps with ops on your Fission deployment.

## Getting Started

```bash
  $ minikube start

  $ skaffold run

  # Add the stock NodeJS env to your Fission deployment
  $ fission env create --name nodejs --image fission/node-env

  # Create a function with a javascript one-liner that prints "hello world"
  $ fission function create --name hello --env nodejs --code https://raw.githubusercontent.com/fission/fission/master/examples/nodejs/hello.js

  # Run the function. This takes about 100msec the first time.
  $ fission function test --name hello
  # > Hello, world!
```

## Learn More

* Understand [Fission Concepts](https://docs.fission.io/docs/concepts/).
* See the [installation guide](https://docs.fission.io/docs/installation/) for installing and running Fission.
* You can learn more about Fission and get started from [Fission Docs](https://docs.fission.io/docs).
* See the [troubleshooting guide](https://docs.fission.io/docs/trouble-shooting/) for debugging your functions and Fission installation.

## Contributing

Check out the [contributing guide](CONTRIBUTING.md).

## Official Releases

Official releases of Fission can be found on [the releases page](https://github.com/fnlize/fnlize/releases).
Please note that it is strongly recommended that you use official releases of Fission, as unreleased versions from
the master branch are subject to changes and incompatibilities that will not be supported in the official releases.

## Licensing

Fission is under the Apache 2.0 license.
