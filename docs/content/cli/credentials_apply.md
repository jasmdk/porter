---
title: "porter credentials apply"
slug: porter_credentials_apply
url: /cli/porter_credentials_apply/
---
## porter credentials apply

Apply changes to a credential set

### Synopsis

Apply changes from the specified file to a credential set. If the credential set doesn't already exist, it is created.

Supported file extensions: json and yaml.

You can use the generate and show commands to create the initial file:
  porter credentials generate mycreds --reference SOME_BUNDLE
  porter credentials show mycreds --output yaml > mycreds.yaml


```
porter credentials apply FILE [flags]
```

### Examples

```
  porter credentials apply mycreds.yaml
```

### Options

```
  -h, --help               help for apply
  -n, --namespace string   Namespace in which the credential set is defined. The namespace in the file, if set, takes precedence.
```

### Options inherited from parent commands

```
      --debug                  Enable debug logging
      --debug-plugins          Enable plugin debug logging
      --experimental strings   Comma separated list of experimental features to enable. See https://getporter.org/configuration/#experimental-feature-flags for available feature flags.
```

### SEE ALSO

* [porter credentials](/cli/porter_credentials/)	 - Credentials commands

