---
name: ccpm-testing
description: Manages testing operations for CCPM projects. Prepares testing environments by detecting test frameworks across 12+ languages, validates dependencies, and executes tests using configured test-runner agent. Use when setting up testing or running tests.
---

<essential_principles>
## How CCPM Testing Works

CCPM testing manages two core operations that ensure reliable, debuggable test execution across multiple programming languages.

### 1. Prime - Environment Preparation

Automatically detects test frameworks across JavaScript/Node.js, Python, Rust, Go, PHP, C#/.NET, Java, Kotlin, Swift, Dart/Flutter, C/C++, and Ruby. Validates dependencies, discovers test files, and creates comprehensive configuration at `.claude/testing-config.md` for the test-runner agent.

### 2. Run - Test Execution

Executes tests using the configured test-runner agent from `.claude/agents/test-runner.md`. Supports running all tests, specific test files, or pattern-based test selection with verbose debugging output.

### Key Testing Principles

**Multi-language framework detection**: Automatically identifies Jest, Mocha, Pytest, PHPUnit, Cargo, Go test, Maven, Gradle, dotnet test, XCTest, Flutter test, CMake/CTest, RSpec, and more.

**Verbose debugging output**: All test execution uses maximum verbosity to capture complete stack traces and failure context.

**Real services only**: No mocking - tests run against actual implementations to ensure realistic behavior and catch integration issues.

**Sequential execution**: Tests run one at a time to avoid race conditions, resource conflicts, and hard-to-debug parallel execution issues.

**Comprehensive validation**: Validates framework presence, dependency installation, test file discovery, and environment configuration before execution.
</essential_principles>

<intake>
What would you like to do?

1. **Prime** - Prepare testing environment (detect frameworks, validate dependencies, configure test-runner)
2. **Run** - Execute tests (all tests or specific files/patterns)

**Wait for response before proceeding.**
</intake>

<routing>
| Response | Operation | Command File |
|----------|-----------|--------------|
| 1, "prime", "prepare", "setup", "configure", "detect" | Prime testing environment | Read and execute `/home/caio/Developer/Claude/ccpm-audit/ccpm/commands/testing/prime.md` |
| 2, "run", "execute", "test", "tests" | Run tests | Read and execute `/home/caio/Developer/Claude/ccpm-audit/ccpm/commands/testing/run.md` |

**After identifying the operation, read the command file and follow its instructions exactly.**
</routing>

<supported_frameworks>
## Framework Support Matrix

**JavaScript/Node.js**: Jest, Mocha, Jasmine (via package.json and config files)

**Python**: Pytest, unittest, nose (via pytest.ini, conftest.py, setup.cfg)

**PHP**: PHPUnit, Pest (via phpunit.xml, composer.json)

**Java**: JUnit with Maven or Gradle (via pom.xml, build.gradle)

**Kotlin**: Kotlin test, Spek (via build.gradle.kts)

**C#/.NET**: MSTest, NUnit, xUnit (via .csproj, .sln files)

**Swift**: XCTest (via Package.swift, Xcode projects)

**Dart/Flutter**: Flutter test (via pubspec.yaml)

**C/C++**: GoogleTest, Catch2 (via CMakeLists.txt)

**Ruby**: RSpec, Minitest (via .rspec, Gemfile)

**Go**: go test (via *_test.go files)

**Rust**: cargo test (via Cargo.toml, #[cfg(test)])
</supported_frameworks>

<configuration_output>
## Generated Configuration

Prime operation creates `.claude/testing-config.md` containing:

- **Framework details**: Type, version, configuration file paths
- **Test structure**: Directory locations, file count, naming patterns
- **Execution commands**: Full suite, specific test, debugging modes
- **Environment requirements**: Environment variables, test databases, services
- **Test-runner agent setup**: Verbose output, sequential execution, real services

This configuration is used by the Run operation to execute tests consistently.
</configuration_output>

<success_criteria>
Testing operations are successful when:

**For Prime:**
- ✅ Test framework detected and validated
- ✅ Dependencies confirmed installed
- ✅ Test files discovered and counted
- ✅ Configuration created at `.claude/testing-config.md`
- ✅ Test-runner agent configured with proper settings
- ✅ Validation confirms setup is working

**For Run:**
- ✅ Tests executed with verbose output
- ✅ Results clearly reported (passed/failed/skipped counts)
- ✅ Failures include stack traces and context analysis
- ✅ Test processes properly cleaned up after execution
- ✅ Test structure validated (not assumed to be code issues)
</success_criteria>
