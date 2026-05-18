### What's changed in v1.0.1

* fix(ci): use crossplane xpkg --help to verify CLI install (no kubeconfig) (by @patrickleet)

  `crossplane version` reports the server version too, which needs
  a kubeconfig — fails 'no configuration has been provided' in CI.
  Switch to `crossplane xpkg --help` so we verify the binary runs
  without making a cluster call.


See full diff: [v1.0.0...v1.0.1](https://github.com/hops-ops/provider-openpanel/compare/v1.0.0...v1.0.1)
