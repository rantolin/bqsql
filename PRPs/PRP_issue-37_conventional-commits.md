# PRP: Conventional Commits & Changelog

## 🎯 Goal
- **Objective:** Implement automated `CHANGELOG.md` generation using `git-chglog` and establish Conventional Commits as the project standard.
- **Why:** To improve release transparency, automate documentation, and prepare for future CI/CD integration.
- **Success Criteria:**
  - [ ] `git-chglog` is configured with `bqsql` specific settings.
  - [ ] `CHANGELOG.md` is successfully generated from existing tags/commits.
  - [ ] A GitHub Action automatically updates the changelog on new version tags (`v*`).
  - [ ] Project documentation (README or DEVELOPMENT.md) is updated to reflect the new commit standard.

## 📚 Context
```yaml
docs:
  - https://www.conventionalcommits.org/en/v1.0.0/
  - https://github.com/git-chglog/git-chglog
codebase:
  examples:
    - N/A (Initial implementation of this standard)
  targets:
    - file: CHANGELOG.md
    - file: .chglog/config.yml
    - file: .chglog/CHANGELOG.tpl.md
    - file: .github/workflows/changelog.yml
    - file: DEVELOPMENT.md
  context:
    - path: .
gotchas:
  - Existing commits (before v0.0.5) are inconsistent; the template needs to handle or skip them gracefully.
  - `git-chglog` is officially archived but still widely used in Go projects; we stick to it as per requirements.
  - The GitHub Action requires `fetch-depth: 0` to see all tags.
```

## 🗺️ Implementation Blueprint
### Data Models / Schema
- **.chglog/config.yml**: Configuration mapping commit types (feat, fix, perf, etc.) to human-readable sections.
- **.chglog/CHANGELOG.tpl.md**: Go template for the Markdown output.

### Integration Points
- **GitHub Actions**: Trigger on `push: tags: ['v*']`.

### Pseudocode / Logic
```bash
# Local generation
git-chglog -o CHANGELOG.md

# CI/CD (GitHub Action)
1. Checkout repo with full history.
2. Install git-chglog.
3. Run git-chglog -o CHANGELOG.md.
4. Commit and push CHANGELOG.md back to main.
```

### Sequential Tasks
- [ ] **Task 1: Initial Setup**: Install `git-chglog` locally and initialize the `.chglog` directory.
- [ ] **Task 2: Configuration**: Customize `.chglog/config.yml` to match Conventional Commits 1.0.0 and `bqsql` repository URL.
- [ ] **Task 3: Templating**: Create `.chglog/CHANGELOG.tpl.md` with links to commits and diffs.
- [ ] **Task 4: Baseline Generation**: Generate the initial `CHANGELOG.md` and verify it handles old commits reasonably.
- [ ] **Task 5: Automation**: Create `.github/workflows/changelog.yml` for automated updates.
- [ ] **Task 6: Documentation**: Update `DEVELOPMENT.md` to instruct contributors on using Conventional Commits.

## 🚀 Future Improvements (Option A: Release PR)
- **Problem**: Current push-to-main strategy fails when `main` is protected.
- **Solution**: Move to a "Release PR" workflow (e.g., using `release-please` or `standard-version`).
- **Workflow**:
    1. Instead of manual tagging, a GitHub Action detects changes and opens a PR with the version bump and `CHANGELOG.md` update.
    2. Merging the PR triggers the actual release and tagging.
- **Benefits**: Respects branch protections, allows review of the changelog before it's permanent, and ensures consistent versioning.

## ✅ Validation Loop
### Level 1: Structure & Syntax
```bash
# Verify YAML and Template existence
ls .chglog/config.yml .chglog/CHANGELOG.tpl.md
# Run git-chglog in dry-run mode (to stdout)
git-chglog
```

### Level 2: Functional Correctness
```bash
# Generate the file and check for content
git-chglog -o CHANGELOG.md
grep "Features" CHANGELOG.md
grep "Bug Fixes" CHANGELOG.md
```

### Level 3: Integration/E2E
```bash
# (Optional/Simulated) Manual tag push to verify local flow
# In a real environment, this would be verified by the GitHub Action run.
```

## 🚫 Anti-Patterns
- **Do not** manually edit `CHANGELOG.md` after automation is in place.
- **Do not** use `git-chglog` versions that aren't compatible with the Go template syntax used.
- **Avoid** complex commit parsers that might fail on non-standard older commits; prefer simple fallbacks.
