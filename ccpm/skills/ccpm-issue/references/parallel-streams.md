<overview>
Patterns for decomposing work into parallel execution streams. Use these patterns to maximize agent parallelization while managing dependencies and conflicts.
</overview>

<stream_patterns>
**By Architectural Layer**

**Database Layer**
- Scope: Schema changes, migrations, models, seeds
- Files: `db/`, `migrations/`, `models/`, `prisma/`, `drizzle/`
- Agent Type: database-specialist
- Dependencies: Usually first - other streams depend on schema

**Service/Business Logic Layer**
- Scope: Core business logic, domain services, utilities
- Files: `src/services/`, `src/lib/`, `src/utils/`, `src/domain/`
- Agent Type: backend-specialist
- Dependencies: Often depends on Database layer

**API Layer**
- Scope: Endpoints, controllers, middleware, validation
- Files: `src/api/`, `src/routes/`, `src/controllers/`, `src/handlers/`
- Agent Type: backend-specialist
- Dependencies: Depends on Service layer types

**UI/Frontend Layer**
- Scope: Components, pages, styles, state
- Files: `src/components/`, `src/pages/`, `src/views/`, `src/styles/`
- Agent Type: frontend-specialist
- Dependencies: Depends on API types/contracts

**Test Layer**
- Scope: Unit tests, integration tests, e2e tests
- Files: `tests/`, `__tests__/`, `*.test.*`, `*.spec.*`
- Agent Type: test-specialist
- Dependencies: Depends on implementation being complete

**By Feature Slice**

For feature-based architectures:

**Feature A - Backend**
- Files: `src/features/featureA/api/`, `src/features/featureA/services/`

**Feature A - Frontend**
- Files: `src/features/featureA/components/`, `src/features/featureA/pages/`

**Feature A - Tests**
- Files: `src/features/featureA/__tests__/`

**By Concern**

**Type Definitions**
- Scope: Shared types, interfaces, contracts
- Files: `src/types/`, `*.d.ts`, `src/interfaces/`
- Agent Type: fullstack-specialist
- Note: Often a coordination point - modify early

**Configuration**
- Scope: Config files, environment setup
- Files: `config/`, `.env*`, `*.config.*`
- Agent Type: devops-specialist
- Note: Usually independent, low conflict risk

**Documentation**
- Scope: README, API docs, inline docs
- Files: `docs/`, `README.md`, `*.md`
- Agent Type: documentation-specialist
- Dependencies: Can run parallel to implementation
</stream_patterns>

<parallelization_strategies>
**Full Parallel**

When streams have no dependencies:
```
Stream A ─────────────►
Stream B ─────────────►
Stream C ─────────────►
```
Best for: Independent features, documentation, isolated modules

**Pipeline**

When streams have sequential dependencies:
```
Stream A ──► Stream B ──► Stream C
```
Best for: Database → API → UI flows

**Hybrid**

Mix of parallel and sequential:
```
Stream A ─────────────►
                        ├──► Stream C
Stream B ─────────────►
```
Best for: Most real-world scenarios

**Diamond**

Converging dependencies:
```
Stream A ──────────────────┐
                           ├──► Stream D
Stream B ──► Stream C ─────┘
```
Best for: Complex features with integration points
</parallelization_strategies>

<conflict_risk_assessment>
**Low Risk**
- Streams work on completely different directories
- No shared types or interfaces
- Clear module boundaries

**Medium Risk**
- Some shared type files
- Common utility functions
- Configuration changes

**High Risk**
- Multiple streams modifying same files
- Shared state management
- Core utility refactoring
</conflict_risk_assessment>

<coordination_strategies>
**Type-First**
1. Define types/interfaces first in separate PR
2. Merge types
3. Parallel implementation using shared types

**Contract-First**
1. Define API contracts (OpenAPI, GraphQL schema)
2. Parallel backend and frontend implementation
3. Integration testing

**Feature Flags**
1. Implement behind feature flag
2. Parallel work on same files (different flags)
3. Sequential flag enabling
</coordination_strategies>

<agent_assignment>
| Stream Type | Recommended Agent | Rationale |
|-------------|------------------|-----------|
| Database | database-specialist | Schema expertise, migration safety |
| Backend API | backend-specialist | API patterns, validation |
| Frontend UI | frontend-specialist | Component patterns, UX |
| Full Feature | fullstack-specialist | End-to-end understanding |
| Tests | test-specialist | Testing patterns, coverage |
| DevOps | devops-specialist | Infrastructure, CI/CD |
</agent_assignment>

<estimation_guidelines>
**Stream Size**
- **Small**: 1-2 hours, single file or simple changes
- **Medium**: 2-4 hours, multiple files, moderate complexity
- **Large**: 4-8 hours, significant implementation

**Parallelization Factor**
- **1.0x**: No parallelization possible (sequential)
- **1.5x**: Limited parallel streams
- **2.0x**: Two major parallel streams
- **3.0x**: Three or more independent streams
- **4.0x+**: Highly parallelizable, minimal dependencies

**Coordination Overhead**
Add 10-20% to parallel estimates for:
- Communication between streams
- Merge conflict resolution
- Integration testing
</estimation_guidelines>

<anti_patterns>
**Over-Parallelization**
- Too many small streams create coordination overhead
- Better: 2-4 substantial streams than 8 tiny ones

**Ignoring Dependencies**
- Starting dependent work before prerequisites
- Results in blocking and wasted effort

**Shared File Free-for-All**
- Multiple streams modifying same files without coordination
- Results in merge conflicts and bugs

**Missing Integration Stream**
- Parallel work without planned integration point
- Results in "works separately, breaks together"
</anti_patterns>
