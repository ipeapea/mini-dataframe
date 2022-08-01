# mini-dataframe

## 内容列表

- [背景](#背景)
- [目标](#目标)

## 背景

本项目思想来源于[gota](https://github.com/go-gota/gota)，该包是一个提供了类似于Python中pandas.DataFrame功能。但该项目存在以下缺点：

1. 支持的数据结构过于简单，目前只支持`int`、`string`、`bool`、`float`，不支持用户自定义`struct`
2. 聚合函数只支持几种简单的数字类，不支持用户自定义聚合函数
3. 没有`dataframe`到`struct`的转换功能
4. 一些bug问题，不如`group_by`对于`bool`类型支持不好，`print`不能自定义输出格式

等等问题，这些问题源于该库中底层实现`Series`有比较大的局限，导致后续想增加这些功能都比较困难。因此才有想法重新实现一套。

## 目标

- [ ] 实现基础的`Element`，支持所有通用的基础类型和用户自定义类型，为上层建设提供基础能力】
- [ ] 实现基本的`Series`能力，支持通用的基础类型和用户自定义类型，提供更好的扩展性
- [ ] 实现基本的`Series`聚合能力，并且提供自定义聚合实现
- [ ] 支持通用的聚合逻辑
- [ ] 实现基本的`DataFrame`能力
- [ ] 周边功能扩展
- [ ] 支持`SQL`能力


## 维护者

[@liracle](https://github.com/liracle)

## 如何贡献
非常欢迎你的假如! [提一个Issue](https://github.com/ipeapea/mini-dataframe/issues/new) 或者提交一个 Pull Request。

### 贡献者


## 使用许可

[MIT](LICENSE) © liracle 
