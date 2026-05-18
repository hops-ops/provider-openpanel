### What's changed in v1.1.0

* feat(sso): regen against terraform-provider-openpanel v0.2.0 (SsoConfig CRD) (by @patrickleet)

  - Bump TERRAFORM_PROVIDER_VERSION 0.1.0 → 0.2.0
  - Register the new `openpanel_organization_sso_config` TF resource in
    config/provider.go with `ShortGroup = "organization"`,
    `Kind = "SsoConfig"`, and a cross-resource reference on
    `organization_id` so MR composition can wire the parent Org.
  - `make generate` regenerates the upjet pipeline: 5 resources × 2
    scopes + ProviderConfig/Usage = 15 CRDs (was 13).

  New CRDs:
    organization.openpanel.crossplane.io/SsoConfig
    organization.openpanel.m.crossplane.io/SsoConfig

  OIDC `client_secret` is a `Sensitive` write-only attribute on the TF
  schema; upjet propagates that — the CRD's spec.forProvider.oidcClientSecretSecretRef
  shape lets operators source the cleartext from a Kubernetes Secret
  (typical pattern: ESO-projected from AWS SM). The CRD also exposes
  status.atProvider.hasOidcClientSecret so composition functions can
  gate on whether a config is "set" without ever seeing the value.

  Build clean: `go build ./...` passes.


See full diff: [v1.0.1...v1.1.0](https://github.com/hops-ops/provider-openpanel/compare/v1.0.1...v1.1.0)
