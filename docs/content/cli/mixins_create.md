---
title: "porter mixins create"
slug: porter_mixins_create
url: /cli/porter_mixins_create/
---
## porter mixins create

Create a new mixin project based on the getporter/skeletor repository

### Synopsis

Create a new mixin project based on the getporter/skeletor repository.
The first argument is the name of the mixin to create and is required.

A flag of --author to declare the author of the mixin is a required input.
A flag of --username to specify the GitHub's username of the mixin's author is a required input.

You can also specify where to put the mixin directory. It will default to the current directory.

```
porter mixins create NAME --author "My Name" --username mygithubusername [--dir /path/to/mixin/dir] [flags]
```

### Examples

```
 porter mixin create MyMixin --author "My Name" --username mygithubusername
		porter mixin create MyMixin --author "My Name" --username mygithubusername --dir path/to/mymixin
		
```

### Options

```
      --author string     Your full name.
      --dir string        Path to the designated location of the mixin's directory.
  -h, --help              help for create
      --username string   Your GitHub username.
```

### Options inherited from parent commands

```
      --debug                  Enable debug logging
      --debug-plugins          Enable plugin debug logging
      --experimental strings   Comma separated list of experimental features to enable. See https://getporter.org/configuration/#experimental-feature-flags for available feature flags.
```

### SEE ALSO

* [porter mixins](/cli/porter_mixins/)	 - Mixin commands. Mixins assist with authoring bundles.

