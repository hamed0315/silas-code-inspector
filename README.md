# Silas Code Inspector

[![Go Version](https://img.shields.io/github/go-mod/go-version/silasdev78/silas-code-inspector)](https://go.dev)
[!Build](https://img.shields.io/github/actions/workflow/status/silasdev78/silas-code-inspector/silas.yml?branch=main)(https://github.com/silasdev78/silas-code-inspector/actions)
[!License](https://img.shields.io/badge/license-MIT-blue)](LICENSE)

Multi-language static security scanner for **TON blockchain (Tact & FunC)**, Go, Docker, and Web.
O
---

## 🚀 Project Evolution & Note for Reviewers

**This repository (v2.0) is a complete, 100x stronger rewrite of my previous work:** 
👉 [Original Version (v1.0 - silasdev78)](https://github.com/silasdev78/silas-code-inspector)

### Why the rewrite?
The original version was built during an internet blackout, using local dependencies. In this new version (v2.0), I have:
- **Refactored the entire architecture** for better performance.
- **Fixed Dependency Hell:** Successfully downgraded breaking packages (`fatih/color` and `golang.org/x/sys`) to be 100% compatible with Go 1.22.
- **Improved the AST parsing engine** for Solidity and Go.
- **Zero-config CI/CD:** Now fully integrated with GitHub Actions and SARIF export.

*This proves my ability to take a product from a raw prototype to a production-ready static analysis tool.*
