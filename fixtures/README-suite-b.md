# Suite B fixtures: transitive reachability

Suite B asks one question: does the tool know whether the application can reach crypto that sits in a dependency? Each fixture is an application that depends on library A, which may depend on library B, and so on. The crypto lives in a library. Every scenario comes as a pair of entries that share the libraries and differ in one change to the application:

- **reachable**: the application's entry point calls into the libraries along a path that ends in the crypto call;
- **unreachable**: the dependency graph is the same and the crypto is still in the library, but no path from the entry point reaches it.

The fixtures live in `fixtures/<language>/suite-b/<scenario>/`: the libraries `liba/`, `libb/` and, in depth3, `libc/`, next to the two applications `app-reachable/` and `app-unreachable/`. The manifest entries are `<language>-b-<directory>-<variant>`, with the application directory as their path. The directory is the scenario's name, except for the handler scenario, which has one directory per framework (`handler-gin`), all with `scenario: handler` in the manifest.

## Scenarios

The first seven scenarios are the same in all five languages, with the same algorithms, so a row of the score reads the same across languages. handler and callback were added after the first results ([METHODOLOGY.md](../METHODOLOGY.md#matching-rules)). handler has one pair per framework below, in every language; callback has a pair in Go and Python only.

| Scenario | Shape | Crypto | Distance | What the unreachable variant changes |
|---|---|---|---|---|
| direct | app -> liba | SHA-256 in liba | 1 | the call to liba moves into an application function that the entry point never calls |
| transitive | app -> liba -> libb | SHA-512 in libb | 2 | same |
| depth3 | app -> liba -> libb -> libc | HMAC-SHA-256 in libc | 3 | same |
| wrapper | app -> liba (two nested wrappers) -> libb (two nested wrappers) | SHA-384 in libb | 2 | same |
| dispatch | app calls a method of an interface, trait or base class liba declares; liba's implementation calls libb | SHA-256 in libb | 2 | the entry point asks liba for its other implementation, which does no crypto |
| deadfn | one libb file holds a used crypto function, a non-crypto sibling and a crypto function nothing calls | HMAC-SHA-256 used, SHA-1 dead | 2 | the entry point calls liba's non-crypto function, which calls the non-crypto sibling in the same libb file |
| unused-import | app imports liba, which imports libb | SHA-256 in libb | 2 | the entry point only reads a constant of liba |
| handler | the entry point starts a web application and registers an application handler with the framework; the handler calls liba -> libb | SHA-256 in libb | 2 | the handler function is still there, but the entry point never registers it and nothing calls it |
| callback | the entry point passes an application function by value to a standard-library function, which calls it; the function calls liba -> libb | SHA-256 in libb | 2 | the entry point passes a different function, and nothing calls the crypto one |

The crypto library is the standard one where the language has it: Go's `crypto/*`, Python's `hashlib` and `hmac`, Java's JCA, Node's `node:crypto`. Rust has none, so its libraries call the RustCrypto crates `sha2` 0.10.8, `hmac` 0.12.1 and `sha1` 0.11.0, which are real registry dependencies of the fixture libraries. Their own code is outside the ground truth (`scope.components`).

The handler pairs, one per framework:

| Language | Directory | Framework | Registration in the reachable variant |
|---|---|---|---|
| Go | `handler-nethttp` | net/http (standard library) | `http.HandleFunc("/fingerprint", fingerprint)` |
| Go | `handler-gin` | gin v1.12.0 | `r.GET("/fingerprint", fingerprint)` |
| Python | `handler-flask` | Flask 3.1.3 | `app.add_url_rule("/fingerprint", view_func=fingerprint)` |
| Python | `handler-fastapi` | FastAPI 0.162.0, uvicorn 0.54.0 | `app.add_api_route("/fingerprint", fingerprint, methods=["GET"])` |
| Java | `handler-spring` | Spring Boot 4.1.2 (Spring MVC, embedded Tomcat) | `FingerprintController` is a `@RestController` and `fingerprint` has `@GetMapping("/fingerprint")`; the unreachable variant's class has neither annotation |
| JavaScript | `handler-express` | Express 5.3.0 | `app.get('/fingerprint', fingerprint)` |
| Rust | `handler-axum` | axum 0.8.9, tokio 1.54.1 | `.route("/fingerprint", get(fingerprint))` |
| Rust | `handler-actix` | actix-web 4.18.1 (default features off, `macros` on) | `.route("/fingerprint", web::get().to(fingerprint))` |

The callback pairs pass `byFingerprint` to `slices.SortFunc` in Go and `report` to `threading.Thread(target=...)` in Python; the unreachable variants pass `strings.Compare` and `greet`. Go's `sort.Slice` takes an index comparator, which in practice is a closure over the slice, so the Go pair uses `slices.SortFunc`, which takes a named function as it is.

Every handler registration happens in the entry point, and the handler is passed by reference (by annotation in Spring). Python's decorator form (`@app.route`, `@app.get`) is not covered. Each application serves on 127.0.0.1:8080 when given the argument `serve`; without it, it sends one request for `/fingerprint?data=hello` through the framework, in process or over an ephemeral local port, prints the status and body, and exits, so `scripts/prepare-suite-b.sh` can run it once. The reachable variant prints `200` and the digest, the unreachable variant `404`. Spring, Express and actix-web answer the request over a local port, so that path runs the framework's real server; the other frameworks answer it through their router in process (`ServeHTTP`, a test client, `oneshot`).

The frameworks are real dependencies from the public registries, pinned exactly: Go in `go.sum`, Python with every transitive package in `requirements.txt`, Rust in the committed `Cargo.lock`, Java through the Spring Boot BOM, JavaScript by Express's exact version (its own dependencies resolve at install time). On 2026-11-05 none of the resolved packages had a known vulnerability in OSV, except the deprecated `golang.org/x/crypto/openpgp` package (GO-2026-5932), which gin does not import and which every released `golang.org/x/crypto` still ships. The frameworks are not in `scope.components`, so findings in them are not scored.

The scenarios separate failure modes. In direct, transitive, depth3 and wrapper, the only caller of the crypto path in the unreachable variant is an application function that nothing calls, so a tool that treats every uncalled function as an entry point reports it reachable. In dispatch, the crypto implementation exists in the reachable dependency graph but is never instantiated in the unreachable variant. deadfn tests function-level precision inside a reachable file, and unused-import tests a dependency that is imported but never called. handler and callback test the reverse of the first four: the crypto path starts at a function the application never calls itself, which a framework or the standard library calls because the entry point handed it over. A tool reads the reachable variant right only if it follows that hand-over, and the unreachable variant right only if the hand-over, not the function's existence, is what makes it reachable.

## How the libraries are dependencies

Each library is a package of its ecosystem, installed through the package manager the way a published package is, not a source directory next to the application. `scripts/prepare-suite-b.sh ENTRY_DIR OUT` builds one entry into `OUT` and runs the application once:

| Language | Mechanism |
|---|---|
| Go | each library is a module; the application's `go.mod` requires it at `v1.0.0` and `replace`s it with the library directory. A framework comes from the Go module proxy |
| Python | each library is a setuptools project; `pip wheel` builds a wheel of each, and `pip install --no-index --find-links` installs the application's `requirements.txt` from those wheels into a virtual environment outside the application. Library A's wheel declares library B in `Requires-Dist`. A framework's pinned packages, every line of `requirements.txt` that does not start with `cbench-`, are downloaded from PyPI into the same wheel directory first, so the install never asks PyPI for a fixture library |
| Java | each library is a Maven project; `mvn install` puts its jar and its sources jar (maven-source-plugin, as on Maven Central) into a local Maven repository; the application's `pom.xml` depends on library A at `1.0.0`. A framework comes from Maven Central |
| JavaScript | each library is an npm package; `npm pack` makes its tarball. The application depends on library A's tarball, library A depends on library B by version (`"1.0.0"`) as a published package does, and the application's `overrides` point that version at library B's tarball. `npm install` writes `package-lock.json` and copies the packages into `node_modules`. A framework comes from the npm registry |
| Rust | each library is a crate; the application and the libraries use path dependencies with a version. `cargo fetch --locked` downloads the registry crates `Cargo.lock` pins, then the build runs offline. `Cargo.lock` is committed for each application |

The libraries are named per scenario (`cbench-transitive-libb`, `org.cbench.transitive:libb`), so no tool cache can serve one scenario's library for another. Run `scripts/prepare-suite-b.sh fixtures/go/suite-b/transitive/app-reachable work/x` to build one; every application prints a check value.

Mechanisms not used:

- **npm `file:` directory dependencies.** npm links them into `node_modules` instead of installing them, which no published package is. crypto-finder's npm resolver also reads the linked directory as a workspace member, part of the application, rather than as a dependency (`internal/dependency/npm_resolver.go` in crypto-finder), so the fixture would not test dependency reachability. Tarballs install the way a registry package does.
- **Python editable installs.** They put a path finder into the environment, not the package files, which no published package does.
