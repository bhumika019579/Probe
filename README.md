# Probe

CLI tool that scans a local repository and collects **code-level evidence for SOC 2 controls**.

Probe checks authentication, access control, secrets, CI/CD, testing, dependencies, documentation, and application hardening.

**Written in Go**  
**CLI built with Cobra**  
**YAML-based control configuration and detector mapping**  
**Supports Go, JavaScript, TypeScript, Python, and Java**

Probe performs **static repository analysis only**. It does not access cloud accounts or make network calls during a scan.

Probe is not an auditor and does not certify SOC 2 compliance. It complements cloud/SaaS evidence tools such as Osto's Evidence Collector.

## Why Probe?

Probe is designed to answer a simple question:

> What security-control evidence can be found in this repository?

It scans the repository, detects relevant signals, maps them to SOC 2 controls, and produces a terminal or JSON report.

**Probe collects evidence that controls exist, and it doesn't find vulnerabilities. That keeps the claim honest and puts you next to tools like Osto instead of against scanners like Snyk.**

## How Probe Works

![Probe Architecture](probe.png)

## SOC 2 Controls

| Control | What Probe Checks |
| --- | --- |
| CC6.1 | Authentication |
| CC6.3 | Access Control |
| CC6.6 | External Threat Protection |
| CC6.7 | Secrets / Credential Protection |
| CC8.1 | Change Management |
| CC7.1 | Testing / Monitoring |
| CC7.2 | Dependencies / Security Issues |
| CC2.1 | Documentation |

## Install

Requires **Go 1.22+**.

Install Probe:

```bash
go install github.com/bhumika019579/probe@latest
```

On Windows, make sure `%USERPROFILE%\go\bin` is added to your PATH.

## Build From Source

```bash
git clone https://github.com/bhumika019579/probe.git
cd probe
go install .
```

## Usage

### Scan the current repository

```bash
probe scan .
```

### Scan a specific repository

```bash
probe scan ./my-project
```

### Generate a JSON report

```bash
probe scan ./my-project --json report.json
```

## What Probe Scans

Probe looks for code-level signals related to:

- Authentication libraries and authentication patterns
- Access-control and route-guard patterns
- Application hardening and input-validation libraries
- Secrets and credential exposure
- CI/CD configuration
- Test files and testing signals
- Dependency lockfiles
- Security and project documentation

Probe scans local repository files and excludes common generated or dependency directories such as:

- `node_modules`
- `vendor`
- `dist`
- `build`
- `.next`
- `coverage`
- `.git`

## Output

Probe reports each SOC 2 control as:

- **FOUND** — relevant evidence was detected
- **PARTIAL** — some code-level evidence was detected, but it is not enough to establish the control
- **GAP** — no matching evidence was detected

Evidence includes the file, location, description, and confidence level where available.

Example:

```text
[CC6.1] Logical Access Security — FOUND

    - package.json:0
      Node.js JWT library ("jsonwebtoken") found in dependencies
      confidence: high
```

## Confidence Levels

**High** — strong evidence or a directly verifiable repository fact.

**Medium** — stronger code-level evidence, but still not proof of runtime behavior.

**Low** — keyword, pattern, or heuristic signal that requires human verification.

## Tech Stack

- Go — core language
- Cobra — CLI commands and argument handling
- YAML — SOC 2 control configuration and detector mapping
- go-yaml/v3 — YAML parsing
- Git — repository-aware checks
- JSON — optional report output

## Supported Languages

Probe can analyze repositories containing:

- Go
- JavaScript
- TypeScript
- Python
- Java

The Probe CLI itself is written in **Go**.

## Limitations

Probe is a static evidence collector, so it cannot verify runtime or organizational behavior.

It does not currently verify:

- MFA enforcement
- Live IAM / SSO configuration
- Cloud account settings
- WAF or DDoS protection
- Backup and disaster-recovery policies
- Employee onboarding/offboarding processes
- Actual test results or coverage
- CVEs or vulnerable dependency versions

A detected signal does not automatically mean that a control is fully implemented or compliant.

## Out of Scope

Probe does not perform vulnerability scanning, penetration testing, cloud configuration auditing, or compliance certification.

It is intended to provide developers and security teams with an initial view of **what control-related evidence exists inside a repository**.

---

Built with Go by [Bhumika Chanchlani](https://github.com/bhumika019579)
