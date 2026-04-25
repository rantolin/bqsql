# 📋 REQUISITOS: Conventional Commits & Changelog

## 🎫 Origen (Ticket/Issue)
> **Issue #37**: Implement automated CHANGELOG.md generation based on Conventional Commits.
> **Roadmap Context**: Phase 1: Performance & Core Stability. P0 priority.

## 📚 Documentación Externa
- [Conventional Commits 1.0.0](https://www.conventionalcommits.org/en/v1.0.0/)
- [git-chglog](https://github.com/git-chglog/git-chglog) - Herramienta recomendada para Go.
- [GitHub Actions: Release Drafter](https://github.com/marketplace/actions/release-drafter) - Alternativa para automatización en CI.

## 🔍 Código de Referencia (MIRROR)
- N/A (Primera implementación de estándares de release).

## 🎯 Archivos Objetivo (TARGET)
- `file: CHANGELOG.md` # Archivo a generar/mantener.
- `file: .chglog/config.yml` # Configuración de git-chglog (si se elige esta herramienta).
- `file: .chglog/CHANGELOG.tpl.md` # Plantilla para el changelog.
- `file: .github/workflows/changelog.yml` # Workflow para automatizar la actualización.

## ⚠️ Gotchas
- [ ] Asegurarse de que los commits existentes no rompan la generación inicial (podría requerir un "base tag").
- [ ] Definir si la generación será manual (pre-release) o automática vía CI.
- [ ] El formato de los tags debe ser consistente (vX.Y.Z).

## 📝 Notas Adicionales
- Se recomienda el uso de `git-chglog` por ser una herramienta escrita en Go, alineada con el stack del proyecto.
- Es fundamental que el equipo empiece a usar el prefijo `feat:`, `fix:`, `chore:`, etc.
- Esta tarea es un "Quick Win" fundacional para la Phase 7 (CI/CD).
