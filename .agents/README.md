# ServiceNow Mapping Service - Agent Team

## 🚀 Team Overview

This is a specialized multi-agent team designed to build and maintain a high-performance Golang microservice with MySQL database integration.

## 👥 Team Members

### 1. **SQL Professional** (`sql-pro`)
- **Role**: Database Expert
- **Specialization**: MySQL optimization, schema design, query performance
- **Key Responsibilities**:
  - Design efficient database schemas
  - Optimize queries for < 100ms response time
  - Implement indexing strategies
  - Ensure data integrity and ACID compliance
  - Handle migrations and versioning

### 2. **Golang Professional** (`golang-pro`)
- **Role**: Backend Developer
- **Specialization**: Idiomatic Go, concurrency, microservices
- **Key Responsibilities**:
  - Implement clean, performant Go code
  - Build concurrent services with goroutines
  - Write comprehensive tests (>80% coverage)
  - Ensure race-condition-free code
  - Follow Go best practices

### 3. **API Designer** (`api-designer`)
- **Role**: Interface Architect
- **Specialization**: REST/GraphQL, OpenAPI, developer experience
- **Key Responsibilities**:
  - Design RESTful API endpoints
  - Create OpenAPI 3.1 specifications
  - Ensure consistent API patterns
  - Implement proper error handling
  - Optimize for developer experience

### 4. **Context Manager** (`context-manager`)
- **Role**: Information Coordinator
- **Specialization**: State management, synchronization, knowledge base
- **Key Responsibilities**:
  - Manage shared context across agents
  - Maintain project state and history
  - Coordinate information flow
  - Ensure < 100ms retrieval times
  - Track system evolution

### 5. **Microservices Architect** (`microservices-architect`)
- **Role**: System Designer
- **Specialization**: Distributed systems, service mesh, cloud-native
- **Key Responsibilities**:
  - Define service boundaries
  - Design communication patterns
  - Implement resilience strategies
  - Ensure scalability and reliability
  - Establish operational excellence

### 6. **Knowledge Synthesizer** (`knowledge-synthesizer`)
- **Role**: Intelligence Builder
- **Specialization**: Pattern recognition, continuous improvement
- **Key Responsibilities**:
  - Extract insights from team interactions
  - Identify success patterns
  - Generate recommendations
  - Track improvement metrics
  - Enable collective learning

## 🔄 Collaboration Patterns

### Architecture Design Phase
**Team**: Microservices Architect → API Designer → Context Manager
- Define service boundaries and communication patterns
- Design API contracts and data models
- Establish context management strategy

### Database Development Phase
**Team**: SQL Pro → Golang Pro → Context Manager
- Design database schema and indexes
- Implement repository layer in Go
- Manage database migrations and state

### API Implementation Phase
**Team**: API Designer → Golang Pro → Microservices Architect
- Create OpenAPI specifications
- Implement REST endpoints
- Ensure resilience patterns

### Quality Improvement Phase
**Team**: Knowledge Synthesizer → Golang Pro → SQL Pro
- Analyze performance metrics
- Identify optimization opportunities
- Implement improvements

### System Optimization Phase
**Team**: Microservices Architect → SQL Pro → Golang Pro → Knowledge Synthesizer
- Review system performance
- Optimize critical paths
- Document best practices

## 📡 Communication Protocol

### Context Requests
All agents query the `context-manager` for shared information:

```json
{
  "requesting_agent": "golang-pro",
  "request_type": "get_golang_context",
  "payload": {
    "query": "Current module structure and dependencies"
  }
}
```

### Status Updates
Agents report progress through standardized JSON messages:

```json
{
  "agent": "golang-pro",
  "status": "implementing",
  "progress": {
    "packages_created": ["api", "service", "repository"],
    "tests_written": 47,
    "coverage": "87%"
  }
}
```

### Knowledge Sharing
`knowledge-synthesizer` extracts and distributes insights to all team members.

## 🛠️ Technology Stack

### Backend
- **Language**: Go 1.21+
- **Framework**: Chi/Gin/Echo (to be determined)
- **Database**: MySQL 8.0+
- **ORM**: GORM or sqlx

### API
- **Style**: RESTful
- **Documentation**: OpenAPI 3.1
- **Format**: JSON

### Infrastructure
- **Containerization**: Docker
- **Orchestration**: Kubernetes (optional)
- **Service Mesh**: Istio (optional)
- **Monitoring**: Prometheus + Grafana

### Development Tools
- **Testing**: Go testing, Testify
- **Linting**: golangci-lint
- **Formatting**: gofmt, goimports
- **Debugging**: Delve

## 📋 Development Workflow

1. **Initialization**
   - Microservices Architect defines system boundaries
   - Context Manager sets up knowledge base
   - API Designer creates initial specifications

2. **Implementation**
   - SQL Pro designs database schema
   - Golang Pro implements services
   - API Designer validates endpoints

3. **Testing**
   - Golang Pro writes comprehensive tests
   - SQL Pro validates query performance
   - Microservices Architect tests resilience

4. **Optimization**
   - Knowledge Synthesizer analyzes metrics
   - Team implements improvements
   - Context Manager tracks progress

5. **Deployment**
   - Microservices Architect validates infrastructure
   - Team reviews production readiness
   - Knowledge Synthesizer documents lessons learned

## 🎯 Quality Standards

### Code Quality
- ✅ Idiomatic Go code
- ✅ Test coverage > 80%
- ✅ Zero race conditions
- ✅ Comprehensive error handling
- ✅ Full documentation

### API Quality
- ✅ OpenAPI 3.1 compliant
- ✅ RESTful principles
- ✅ Consistent naming
- ✅ Proper status codes
- ✅ Comprehensive examples

### Database Quality
- ✅ Query performance < 100ms
- ✅ Optimal indexing
- ✅ Data integrity enforced
- ✅ Migration versioning
- ✅ Backup strategy

### System Quality
- ✅ High availability (99.9%+)
- ✅ Scalability validated
- ✅ Security best practices
- ✅ Observability enabled
- ✅ Documentation complete

## 📚 Resources

### Agent Definitions
- [`sql-pro.md`](./.agents/sql-pro.md)
- [`golang-pro.md`](./.agents/golang-pro.md)
- [`api-designer.md`](./.agents/api-designer.md)
- [`context-manager.md`](./.agents/context-manager.md)
- [`microservices-architect.md`](./.agents/microservices-architect.md)
- [`knowledge-synthesizer.md`](./.agents/knowledge-synthesizer.md)

### Project Context
- [`team-manifest.json`](./.agents/team-manifest.json) - Team configuration
- Project structure and conventions
- API specifications
- Database schema

## 🚦 Getting Started

To activate the agent team:

1. Review agent definitions in `.agents/` directory
2. Initialize project context via `context-manager`
3. Consult `microservices-architect` for system design
4. Coordinate with `api-designer` for API contracts
5. Implement with `golang-pro` and `sql-pro`
6. Synthesize learnings with `knowledge-synthesizer`

---

**Team Status**: ✅ Loaded and Ready
**Project**: ServiceNow Mapping Service
**Technology**: Golang + MySQL Microservice
**Created**: 2025-11-21
