# netconf

[![GoDoc](https://godoc.org/github.com/einarnn/go-netconf/netconf?status.svg)](https://godoc.org/github.com/einarnn/go-netconf/netconf)
[![Report Card](https://goreportcard.com/badge/github.com/einarnn/go-netconf)](https://goreportcard.com/report/github.com/einarnn/go-netconf)
![Build Status](https://github.com/einarnn/go-netconf/actions/workflows/go.yml/badge.svg)

This library is a **fork** of a simple NETCONF client based on [RFC6241](http://tools.ietf.org/html/rfc6241) and [RFC6242](http://tools.ietf.org/html/rfc6242) (although not fully compliant yet). The original fork is at [https://github.com/Juniper/go-netconf](https://github.com/Juniper/go-netconf)

> **Note:** This is currently pre-alpha release.  API and features may and probably will change.  Suggestions and pull requests are welcome.

## Features
* Support for SSH transport using `golang.org/x/crypto/ssh`. (Other transports are planned).
* Built in RPC support (in progress).
* Support for custom RPCs.
* Independent of XML library.  Free to choose encoding/xml or another third party library to parse the results.

Since the fork the following minor additions have been made:

* Removed use of deprecated standard library x509 functions.
* Added stripping of NETCONF 1.1 chunk separators.
* Now support XPath filters for `get` and `get-config`.


## Install
* Requires Go 1.19 or later!
```bash
$ go get github.com/einarnn/go-netconf/netconf
```

## Example
* See examples in `examples/` directory.

## Documentation

You can view full API documentation at GoDoc: http://godoc.org/github.com/einarnn/go-netconf/netconf

## License
(BSD 2)

Copyright © 2013-2018 Juniper Networks, Inc. All rights reserved.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

(1) Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

(2) Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS “AS IS” AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

The views and conclusions contained in the software and documentation are those of the authors and should not be interpreted as representing official policies, either expressed or implied, of Juniper Networks.

Authors and Contributors
------------------------
* [Brandon Bennett](https://github.com/nemith), Facebook
* [Charl Matthee](https://github.com/charl)
* [Jade Auer](https://github.com/jda)
* [Wayne Tucker](https://github.com/wtucker)
* [Christian Giese](https://github.com/GIC-de), Juniper Networks
* [Einar Nilsen-Nygaard](https://github.com/einarnn), Cisco
