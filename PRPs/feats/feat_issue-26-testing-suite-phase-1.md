# 📋 REQUISITOS: Testing Suite (Phase 1: Utilities)

## 🎫 Origen (Ticket/Issue)
> **Issue #26**: Establish a comprehensive testing suite (Unit, Integration) for the new SDK.
> **Roadmap Context**: Phase 4: Architecture & Internal SDK Refactor. P0 priority (Acelerado por decisión estratégica para asegurar calidad desde el inicio).

## 📚 Documentación Externa
- [Go Testing Package](https://pkg.go.dev/testing)
- [Table-Driven Tests](https://go.dev/wiki/TableDrivenTests)
- [BigQuery Go Client Documentation](https://cloud.google.com/go/bigquery)

## 🔍 Código de Referencia (MIRROR)
- N/A (No existen tests actualmente en el repositorio). Se utilizará el patrón de "Table-Driven Tests" como estándar del proyecto.

## 🎯 Archivos Objetivo (TARGET)
- `path: utils/` # Foco inicial de la suite de pruebas.
- `file: utils/formats_test.go` # Suite de pruebas (TDD).
- `file: utils/formats.go` # Definición de la interfaz `RowProvider` y refactorización.

## 🛠 Metodología (TDD)
- **Ciclo Rojo-Verde-Refactor**: Cada cambio en la lógica debe ser precedido por un test que falle.
- **Contract Testing**: Las implementaciones de `RowProvider` (Mock y Real) deben ser validadas contra un test de contrato común.
- **Table-Driven Tests**: Obligatorio para probar múltiples casos de borde en el cálculo de anchos.

## 🏗 Decisiones Arquitectónicas
- [x] **Abstracción de Datos**: Se implementará la interfaz `RowProvider` con métodos `Next()` y `Schema()`.
- [x] **Estrategia de Mocking**: Se utilizará un `MockRowProvider` manual para desacoplar los tests de la SDK de BigQuery.
- [x] **Wrapper Pattern**: El `RowIterator` real de BigQuery será envuelto en un `BigQueryRowProvider` para cumplir la interfaz sin modificar la SDK externa.

## ⚠️ Restricciones
- **Compatibilidad**: Los comandos en `cmd/query.go` y `cmd/head.go` deben seguir funcionando tras la refactorización (uso del wrapper).
- **Pureza**: No se permite el uso de dependencias externas de mocking (como `testify` o `mockery`) en esta fase para mantener la suite ligera y entender los fundamentos.

## 📝 Notas Adicionales
- Esta tarea es el requisito previo para el "Library Decoupling" (#25).
- Se empezará por las funciones más "puras" (como `max` y `PrintFormatedRow`) antes de abordar las que tienen dependencias de BigQuery.
- Se verificará la ejecución con `go test ./utils/...`.
