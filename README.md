# Qmgo 

[![Build Status](https://travis-ci.org/qiniu/qmgo.png?branch=master)](https://travis-ci.org/qiniu/qmgo)
[![Coverage Status](https://codecov.io/gh/qiniu/qmgo/branch/master/graph/badge.svg)](https://codecov.io/gh/qiniu/qmgo)
[![Go Report Card](https://goreportcard.com/badge/github.com/qiniu/qmgo)](https://goreportcard.com/report/github.com/qiniu/qmgo)
[![GitHub release](https://img.shields.io/github/v/tag/qiniu/qmgo.svg?label=release)](https://github.com/qiniu/qmgo/releases)
[![GoDoc](https://pkg.go.dev/badge/github.com/qiniu/qmgo?status.svg)](https://pkg.go.dev/github.com/qiniu/qmgo?tab=doc) 

English | [简体中文](README_ZH.md)

`Qmgo` is a `Go` `driver` for `MongoDB` . It is based on [MongoDB official driver](https://github.com/mongodb/mongo-go-driver), but easier to use like [mgo](https://github.com/go-mgo/mgo) (such as the chain call). 

- `Qmgo` allows users to use the new features of `MongoDB` in a more elegant way.

- `Qmgo` is the first choice for migrating from `mgo` to the new `MongoDB driver` with minimal code changes.

## // md5:352d18eff92a20ab## # 特性
- 支持所有官方选项的文档创建、读取、更新和删除（CRUD）
- 排序（Sort）、限制（Limit）、计数（Count）、选择（Select）、唯一（Distinct）
- 事务处理
- 钩子（Hooks）：用于在操作前后执行自定义功能
- 自动默认字段和自定义字段
- 预定义操作符键
- 聚合操作、索引操作、游标管理
- 验证标签：用于数据校验
- 插件支持

# 特性
- 支持所有官方选项的文档创建、读取、更新和删除（CRUD）
- 排序（Sort）、限制（Limit）、计数（Count）、选择（Select）、唯一（Distinct）
- 事务处理
- 钩子（Hooks）：用于在操作前后执行自定义功能
- 自动默认字段和自定义字段
- 预定义操作符键
- 聚合操作、索引操作、游标管理
- 验证标签：用于数据校验
- 插件支持

// md5:41394b5b99ea90b1

- Use `go mod` to automatically install dependencies by `import github.com/qiniu/qmgo`

Or 

- Use `go get github.com/qiniu/qmgo`

## # 

# 

// md5:2bf52807f1a8969d

Below we give an example of multi-file search、sort and limit to illustrate the similarities between `qmgo` and `mgo` and the improvement compare to `go.mongodb.org/mongo-driver`.
How do we do in`go.mongodb.org/mongo-driver`:

```go
// go.mongodb.org/mongo-driver
// find all, sort and limit
findOptions := options.Find()
findOptions.SetLimit(7) // set limit
var sorts D
sorts = append(sorts, E{Key: "weight", Value: 1})
findOptions.SetSort(sorts) // set sort

batch := []UserInfo{}
cur, err := coll.Find(ctx, bson.M{"age": 6}, findOptions)
cur.All(ctx, &batch)
```

How do we do in `Qmgo` and `mgo`:

```go
// qmgo
// find all, sort and limit
batch := []UserInfo{}
cli.Find(ctx, bson.M{"age": 6}).Sort("weight").Limit(7).All(&batch)

// mgo
// find all, sort and limit
coll.Find(bson.M{"age": 6}).Sort("weight").Limit(7).All(&batch)
```

## # `Qmgo`与`mgo`的区别
[QMGO与Mgo之间的差异](https://github.com/qiniu/qmgo/wiki/Qmgo-与-Mgo-之间的差异)

# `Qmgo`与`mgo`的区别
[QMGO与Mgo之间的差异](https://github.com/qiniu/qmgo/wiki/Qmgo-与-Mgo-之间的差异)

// md5:ff308faf71e49c1b

The Qmgo project welcomes all contributors. We appreciate your help! 

## # 交流：

- 参加[qmgo 讨论区](https://github.com/qiniu/qmgo/discussions)

# 交流：

- 参加[qmgo 讨论区](https://github.com/qiniu/qmgo/discussions)

// md5:6b6ac083a753ecb9