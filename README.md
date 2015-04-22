# go-itc

[![Build Status](https://travis-ci.org/asp2insp/go-itc.svg?branch=master)](https://travis-ci.org/asp2insp/go-itc)

Interval Tree Clocks written in Go

[Interval Tree Clocks](https://github.com/ricardobcl/Interval-Tree-Clocks) are a generalization of [Version Vectors](http://en.wikipedia.org/wiki/Version_vector) and [Vector Clocks](http://en.wikipedia.org/wiki/Vector_clock) proposed in [2008](http://gsd.di.uminho.pt/members/cbm/ps/itc2008.pdf).

They provide causality context for distributed event processing. This is an implementation in Go which follows the author's original paper.