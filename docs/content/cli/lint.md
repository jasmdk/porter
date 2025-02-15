---
title: "porter lint"
slug: porter_lint
url: /cli/porter_lint/
---
## porter lint

Lint a bundle

### Synopsis

Check the bundle for problems and adherence to best practices by running linters for porter and the mixins used in the bundle.

The lint command is run automatically when you build a bundle. The command is available separately so that you can just lint your bundle without also building it.

```
porter lint [flags]
```

### Examples

```
  porter lint
  porter lint --file path/to/porter.yaml
  porter lint --output plaintext

```

### Options

```
  -f, --file string     Path to the porter manifest file. Defaults to the bundle in the current directory.
  -h, --help            help for lint
  -o, --output string   Specify an output format.  Allowed values: plaintext, json (default "plaintext")
  -v, --verbose         Enable verbose logging
```

### Options inherited from parent commands

```
      --debug                  Enable debug logging
      --debug-plugins          Enable plugin debug logging
      --experimental strings   Comma separated list of experimental features to enable. See https://getporter.org/configuration/#experimental-feature-flags for available feature flags.
```

### SEE ALSO

* [porter](/cli/porter/)	 - With Porter you can package your application artifact, client tools, configuration and deployment logic together as a versioned bundle that you can distribute, and then install with a single command.

Most commands require a Docker daemon, either local or remote.

Try our QuickStart https://getporter.org/quickstart to learn how to use Porter.


