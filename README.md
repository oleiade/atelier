<p align="center"><img src="logo.png" alt="atelier logo"/></p>
<h1 align="center">Golang algorithms, data structures, APIs and helpers to resort to when in need</h3>

<p align="center">
    <a href="http://github.com/oleiade/atelier/releases"><img src="https://img.shields.io/github/release/oleiade/atelier.svg" alt="release"></a>
    <a href="https://github.com/oleiade/atelier/actions/workflows/build.yml"><img src="https://github.com/oleiade/atelier/actions/workflows/build.yml/badge.svg" alt="Build status"></a>
</p>

Atelier is a Go library providing a set of algorithms, data structures, APIs and helpers to resort to when in need.

It is essentially my personal library of Go code, in a single place, easily accessible. Types, tools, helpers, that I use on a recurrent basis, and that I don't want to rewrite every time I need them.

## Content

### Data structures

#### [Trie](./trie.go)

The Trie is a versatile tree-like structure optimized for **storing and searching strings**. Tries are ideal for operations such as autocomplete, prefix lookup, and spell checking, as they provide efficient means of storing and retrieving words based on their prefixes. This Trie implementation supports insertion, search, autocomplete, and also offers specialized methods to find words that start or end with specific substrings, making it suitable for use cases that require fast and flexible string searching capabilities.

### Tooling

#### Debugging

##### [MapAddressToWord](./debug.go)

Provides a way to **map the memory address of a given pointer to a human-readable word**. This utility is particularly useful in debugging scenarios, where tracking and identifying pointers by their raw addresses can be cumbersome. By leveraging a hash function, this method calculates an index into a predefined WordList, transforming a numeric memory address into a more memorable and recognizable word. This approach simplifies the process of monitoring and distinguishing different pointer variables during debugging, making it easier to follow their behavior and interactions within the program.