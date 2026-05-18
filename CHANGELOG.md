### What's changed in v1.0.0

* feat: initial scaffold — Crossplane provider generated from terraform-provider-openpanel via upjet (by @patrickleet)

  Source: cloned and adapted from crossplane-contrib/provider-upjet-zitadel
  (v2.2.0 upjet, same toolchain as Zitadel + AWS upjet families).

  Generates 4 managed resources × 2 API scopes (cluster + namespaced
  per Crossplane v2):

    organizations.organization.openpanel.crossplane.io
    organizations.organization.openpanel.m.crossplane.io
    projects.project.openpanel.crossplane.io
    projects.project.openpanel.m.crossplane.io
    clients.client.openpanel.crossplane.io
    clients.client.openpanel.m.crossplane.io
    references.reference.openpanel.crossplane.io
    references.reference.openpanel.m.crossplane.io

  ProviderConfig (cluster + namespaced) takes client-pair credentials
  from a referenced Kubernetes Secret matching the JSON shape:

    { "host": "...", "client_id": "...", "client_secret": "..." }

  Compatible with the credential AnalyticsStack already pushes to AWS
  Secrets Manager at `push/<cluster>/openpanel-credentials` — consumers
  build their own ExternalSecret + ProviderConfig referencing that path.

  Sources `hops-ops/openpanel` v0.1.0 from the Terraform Registry at
  runtime; bumping the upstream TF provider is a Makefile var change.

* ci: replace Zitadel-template workflows with vnext + ghcr xpkg publish (by @patrickleet)

  The initial scaffold inherited 7 Zitadel-template workflows that
  reference Upbound's xpkg.upbound.io registry + crossplane-contrib's
  backport / chatops machinery — none of which apply here. They fired
  red on every push (CI workflow failing at scaffold).

  Replace with the same vnext-driven pattern the rest of hops-ops uses:

    on-pr.yaml            build + lint + go-vet on PRs to main
    on-push-main.yaml     same gates + vnext version-and-tag (DEPLOY_KEY
                          installed via `vnext generate-deploy-key`)
    on-version-tagged.yaml release notes via simple-release + build the
                          xpkg + push to ghcr.io/hops-ops/provider-upjet-openpanel

  xpkg push uses `crossplane xpkg push` against GHCR with the repo's
  GITHUB_TOKEN (read/write packages permission already granted via the
  workflow `permissions:` block).

  Conventional commits drive future tags. First push after this lands
  becomes v0.1.0.

* feat: rename to provider-openpanel (by @patrickleet)

  BREAKING CHANGE: Drop the `-upjet-` infix from the repo name and all in-tree
  references. Mechanical sweep — module path, package name in
  Makefile + crossplane.yaml, generated CRD imports (zz_*.go),
  image directory, README, install example, on-version-tagged
  workflow.

  The `upjet` framework is an implementation detail; consumers care
  that this is the Crossplane provider for OpenPanel. Matches the
  `provider-helm`, `provider-kubernetes` naming convention rather
  than the explicit `provider-upjet-zitadel` convention from
  crossplane-contrib.

  BREAKING CHANGE: import path and OCI image change:

    github.com/hops-ops/provider-upjet-openpanel    →
    github.com/hops-ops/provider-openpanel

    ghcr.io/hops-ops/provider-upjet-openpanel:vN.M.P →
    ghcr.io/hops-ops/provider-openpanel:vN.M.P

  There are no existing releases under the old name (the v0.1.0
  build was cancelled mid-flight when the rename happened), so this
  is a fresh start. Next vnext tag becomes v0.1.0 under the new
  name.

* fix(lint): add doc comments on v1alpha1 stub exports + tighten zz_ exclusion (by @patrickleet)

  The CI run on the rename commit (92e9dc6) failed lint with 6
  revive errors: 4 in the v1alpha1 stub register.go files (missing
  doc comments on exported Group / SchemeGroupVersion / SchemeBuilder)
  and 2 package-comment complaints on upjet-generated zz_setup.go files.

  For the stubs: add proper doc comments. For zz_setup.go: the existing
  .golangci.yml exclusion `path: zz_` was too lax for v2 regex
  matching apparently — tighten to `(^|/)zz_.*\.go$` and also add it
  to exclusions.paths so the files are skipped at file-walk time, not
  just per-rule.


