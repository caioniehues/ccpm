---
name: ccpm-testing
description: Configure and execute tests with intelligent framework detection and test-runner agent integration. Supports priming test environment and running tests.
---

<objective>
Provide comprehensive testing support for Claude Code projects by detecting test frameworks, configuring test runners, and executing tests with detailed analysis. The skill ensures consistent, reliable test execution across different languages and frameworks.

This skill handles two core actions:
- **prime**: Detect framework, validate dependencies, configure test-runner
- **run**: Execute tests with the configured test-runner agent
</objective>

<shared_references>
Load before any operation:
- @ccpm/skills/shared-references/datetime.md
</shared_references>

<action name="prime">
<description>
Prepare the testing environment by detecting the test framework, validating dependencies, and configuring the test-runner agent for optimal test execution.
</description>

<preflight>
**Framework Detection by Language:**

JavaScript/Node.js:
- Check package.json for test scripts: `grep -E '"test"|"jest"|"mocha"' package.json 2>/dev/null`
- Config files: `ls jest.config.* .mocharc.* 2>/dev/null`
- Test directories: `find . -type d \( -name "test" -o -name "tests" -o -name "__tests__" \) -maxdepth 3 2>/dev/null`

Python:
- pytest: `find . -name "pytest.ini" -o -name "conftest.py" 2>/dev/null`
- Requirements: `grep -E "pytest|unittest" requirements.txt 2>/dev/null`

Rust:
- Cargo tests: `grep '\[dev-dependencies\]' Cargo.toml 2>/dev/null`
- Test modules: `find . -name "*.rs" -exec grep -l "#\[cfg(test)\]" {} \; 2>/dev/null | head -5`

Go:
- Test files: `find . -name "*_test.go" 2>/dev/null | head -5`

PHP:
- PHPUnit: `find . -name "phpunit.xml*" 2>/dev/null`
- Pest: `grep "pestphp/pest" composer.json 2>/dev/null`

C#/.NET:
- Test frameworks: `find . -name "*.csproj" -exec grep -l -E "Microsoft\.NET\.Test|NUnit|xunit" {} \; 2>/dev/null`

Java/Kotlin:
- Maven: `grep "junit" pom.xml 2>/dev/null`
- Gradle: `grep -E "junit|testImplementation" build.gradle* 2>/dev/null`

Swift:
- XCTest: `grep "XCTest" Package.swift 2>/dev/null`

Dart/Flutter:
- Flutter test: `grep "flutter_test" pubspec.yaml 2>/dev/null`

Ruby:
- RSpec: `find . -name ".rspec" -o -name "spec_helper.rb" 2>/dev/null`
- Minitest: `grep "minitest" Gemfile 2>/dev/null`

C/C++:
- GoogleTest/Catch2: `grep -E "gtest|GTest|Catch2" CMakeLists.txt 2>/dev/null`
</preflight>

<framework_configurations>
**JavaScript/Node.js (Jest):**
```yaml
framework: jest
test_command: npm test
test_directory: __tests__
config_file: jest.config.js
options: [--verbose, --no-coverage, --runInBand]
environment: { NODE_ENV: test }
```

**JavaScript/Node.js (Mocha):**
```yaml
framework: mocha
test_command: npm test
test_directory: test
config_file: .mocharc.js
options: [--reporter spec, --recursive, --bail]
environment: { NODE_ENV: test }
```

**Python (Pytest):**
```yaml
framework: pytest
test_command: pytest
test_directory: tests
config_file: pytest.ini
options: [-v, --tb=short, --strict-markers]
environment: { PYTHONPATH: . }
```

**Rust:**
```yaml
framework: cargo
test_command: cargo test
test_directory: tests
config_file: Cargo.toml
options: [--verbose, --nocapture]
```

**Go:**
```yaml
framework: go
test_command: go test
test_directory: .
config_file: go.mod
options: [-v, ./...]
```

**PHP (PHPUnit):**
```yaml
framework: phpunit
test_command: ./vendor/bin/phpunit
test_directory: tests
config_file: phpunit.xml
options: [--verbose, --testdox]
environment: { APP_ENV: testing }
```

**C#/.NET:**
```yaml
framework: dotnet
test_command: dotnet test
test_directory: .
options: [--verbosity normal]
```

**Java (Maven):**
```yaml
framework: maven
test_command: mvn test
test_directory: src/test/java
config_file: pom.xml
```

**Java (Gradle):**
```yaml
framework: gradle
test_command: ./gradlew test
test_directory: src/test/java
config_file: build.gradle
options: [--info, --continue]
```

**Swift:**
```yaml
framework: swift
test_command: swift test
test_directory: Tests
config_file: Package.swift
options: [--verbose]
```

**Dart/Flutter:**
```yaml
framework: flutter
test_command: flutter test
test_directory: test
config_file: pubspec.yaml
options: [--verbose]
```

**Ruby (RSpec):**
```yaml
framework: rspec
test_command: bundle exec rspec
test_directory: spec
config_file: .rspec
options: [--format documentation, --color]
environment: { RAILS_ENV: test }
```

**C/C++ (CMake):**
```yaml
framework: cmake
test_command: ctest
test_directory: build
config_file: CMakeLists.txt
options: [--verbose, --output-on-failure]
```
</framework_configurations>

<process>
1. **Detect Framework**
   - Run detection commands for each supported language
   - Identify primary test framework

2. **Validate Dependencies**
   - Check if test dependencies are installed
   - If missing, suggest installation:
     - Node.js: `npm install` or `pnpm install`
     - Python: `pip install -r requirements.txt` or `poetry install`
     - PHP: `composer install`
     - Java: `mvn clean install` or `./gradlew build`
     - C#/.NET: `dotnet restore`
     - Ruby: `bundle install`
     - Dart/Flutter: `flutter pub get`
     - Swift: `swift package resolve`
     - C/C++: `cmake .. && make`

3. **Discover Test Files**
   - Count test files for detected framework
   - Identify test naming patterns

4. **Create Configuration**
   Save to `.claude/testing-config.md`:
   ```markdown
   ---
   framework: {detected_framework}
   test_command: {detected_command}
   created: {REAL datetime}
   ---

   # Testing Configuration

   ## Framework
   - Type: {framework_name}
   - Config File: {config_file_path}

   ## Test Structure
   - Test Directory: {test_dir}
   - Test Files: {count} files found

   ## Commands
   - Run All Tests: `{full_test_command}`
   - Run Specific Test: `{specific_test_command}`

   ## Test Runner Agent Configuration
   - Use verbose output for debugging
   - Run tests sequentially (no parallel)
   - Capture full stack traces
   - No mocking - use real implementations
   ```

5. **Output Summary**
   ```
   🧪 Testing Environment Primed

   🔍 Detection Results:
     ✅ Framework: {framework_name}
     ✅ Test Files: {count} files
     ✅ Dependencies: All installed

   ⚡ Ready Commands:
     - Run all tests: /testing:run
     - Run specific: /testing:run {test_file}

   💡 Tips:
     - Always run tests with verbose output
     - Check test structure if tests fail
   ```
</process>

<error_handling>
- **No Framework Detected**: "⚠️ No test framework found. Please specify your testing setup."
- **Missing Dependencies**: "❌ Test framework not installed. Install dependencies first."
- **No Test Files**: "⚠️ No test files found. Create tests first or check directory location."
</error_handling>
</action>

<action name="run">
<description>
Execute tests with the configured test-runner agent, providing detailed analysis of results.
</description>

<preflight>
1. **Check Configuration**
   ```bash
   test -f .claude/testing-config.md || echo "❌ Testing not configured. Run /testing:prime first"
   ```

2. **Validate Target** (if provided)
   ```bash
   # For file targets
   test -f "$ARGUMENTS" || echo "⚠️ Test file not found: $ARGUMENTS"
   ```
</preflight>

<process>
1. **Determine Test Command**
   - No arguments → Run full test suite from config
   - File path → Run specific test file
   - Pattern → Run tests matching pattern

2. **Execute Tests**
   Use the test-runner agent from `.claude/agents/test-runner.md`:
   ```markdown
   Execute tests for: $ARGUMENTS (or "all" if empty)

   Requirements:
   - Run with verbose output for debugging
   - No mocks - use real services
   - Capture full output including stack traces
   - If test fails, check test structure before assuming code issue
   ```

3. **Monitor Execution**
   - Show test progress
   - Capture stdout and stderr
   - Note execution time

4. **Report Results**

   **Success:**
   ```
   ✅ All tests passed ({count} tests in {time}s)
   ```

   **Failure:**
   ```
   ❌ Test failures: {failed_count} of {total_count}

   {test_name} - {file}:{line}
     Error: {error_message}
     Likely: {test issue | code issue}
     Fix: {suggestion}

   Run with more detail: /testing:run {specific_test}
   ```

   **Mixed:**
   ```
   Tests complete: {passed} passed, {failed} failed, {skipped} skipped

   Failed:
   - {test_1}: {brief_reason}
   - {test_2}: {brief_reason}
   ```

5. **Cleanup**
   ```bash
   # Kill any hanging test processes
   pkill -f "jest|mocha|pytest|phpunit|rspec|ctest" 2>/dev/null || true
   pkill -f "mvn.*test|gradle.*test|dotnet.*test|cargo.*test|go.*test" 2>/dev/null || true
   ```
</process>

<error_handling>
- **Test command fails**: "❌ Test execution failed: {error}. Check test framework is installed."
- **Timeout**: Kill process and report: "❌ Tests timed out after {time}s"
- **No tests found**: "❌ No tests found matching: $ARGUMENTS"
</error_handling>
</action>

<success_criteria>
- **prime**: Framework detected, dependencies validated, configuration saved
- **run**: Tests executed successfully with detailed reporting
- Test-runner agent properly configured
- Clear, actionable feedback on failures
- No hanging test processes
</success_criteria>
