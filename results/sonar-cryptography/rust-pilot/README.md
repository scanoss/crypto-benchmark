# sonar-cryptography, suite A Rust pilot

Not run: sonar-cryptography does not support Rust.

sonar-cryptography 1.7.0, the version the benchmark pins (in CBOMkit-action v2.3.0 with cbomkit-lib 1.3.0), ships detection rules for Java, Python and Go only, driven through the SonarQube Java, Python and Go frontends. SonarQube has no Rust frontend for the plugin to use, so there is no configuration under which it could report an asset in the 15 Rust entries. `scripts/run-sonar-cryptography.sh` would fall back to its Python mode on a Rust directory and scan nothing.

There are no predictions and no score. In the [Rust comparison](../../rust-pilot.md), sonar-cryptography is listed as not supported rather than as recall 0, so that a missing language is not read as missed crypto.
