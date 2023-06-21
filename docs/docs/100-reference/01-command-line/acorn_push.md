---
title: "acorn push"
---
## acorn push

Push an image to a remote registry

```
acorn push [flags] IMAGE
```

### Options

```
  -h, --help                            help for push
  -k, --key string                      Key to use for signing (default "./cosign.key")
  -s, --sign                            Sign the image before pushing
  -a, --signature-annotations strings   Annotations to add to the signature
```

### Options inherited from parent commands

```
  -A, --all-projects        Use all known projects
      --debug               Enable debug logging
      --debug-level int     Debug log level (valid 0-9) (default 7)
      --kubeconfig string   Explicitly use kubeconfig file, overriding current project
  -j, --project string      Project to work in
```

### SEE ALSO

* [acorn](acorn.md)	 - 

