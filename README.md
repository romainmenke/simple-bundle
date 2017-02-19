# Simple Bundle

A lazy man's bundler

`go get github.com/romainmenke/simple-bundle`

---

I needed a tool to bundle css and js files for golang web projects. I did't want a transpiler and I didn't want to add node as a dependency.

- it loops over all files in a directory
- bundles files with the same extension
- saves to `bundle.` version

I use it with `//go:generate simple-bundle` and [modd](https://github.com/cortesi/modd).

---

### Options

- `-h`            : help
- `-source`       : source directory
- `-out`          : output directory
- `trailing args` : exclusion -> simple `must not contain` logic

---

### Simple

- [simple-mini](https://github.com/romainmenke/simple-mini)
- [simple-bundle](https://github.com/romainmenke/simple-bundle)
- [simple-gzip](https://github.com/romainmenke/simple-gzip)
