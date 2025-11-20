# 🚀 Quick Start Guide - Agent Team Activation

## ✅ Team Successfully Loaded!

Your specialized agent team for the **ServiceNow Mapping Service** is now ready to work on your Golang microservice with MySQL database.

## 📦 What's Been Created

```
service-now-mapping-service/
├── README.md                          # Project overview
└── .agents/                           # Agent team directory
    ├── README.md                      # Team documentation
    ├── team-manifest.json             # Team configuration
    ├── PROJECT_CONTEXT.md             # Shared project context
    ├── sql-pro.md                     # Database expert
    ├── golang-pro.md                  # Go developer
    ├── api-designer.md                # API architect
    ├── context-manager.md             # Knowledge coordinator
    ├── microservices-architect.md     # System designer
    └── knowledge-synthesizer.md       # Intelligence builder
```

## 👥 Your Agent Team

| Agent | Role | Expertise |
|-------|------|-----------|
| **sql-pro** | Database Expert | MySQL optimization, schema design, query performance |
| **golang-pro** | Backend Developer | Idiomatic Go, concurrency, testing, microservices |
| **api-designer** | Interface Architect | REST/GraphQL, OpenAPI, developer experience |
| **context-manager** | Information Coordinator | State management, knowledge base, synchronization |
| **microservices-architect** | System Designer | Distributed systems, resilience, cloud-native |
| **knowledge-synthesizer** | Intelligence Builder | Pattern recognition, best practices, improvement |

## 🎯 How to Work with Your Team

### 1. Query Context Manager
Get current project state and shared information:

```json
{
  "requesting_agent": "your-request",
  "request_type": "get_project_context",
  "payload": {
    "query": "What's the current database schema?"
  }
}
```

### 2. Consult Specific Experts

**For Database Work** - Talk to `sql-pro`:
- "Design a schema for mapping configurations"
- "Optimize this query for better performance"
- "Create indexes for the mappings table"

**For Go Implementation** - Talk to `golang-pro`:
- "Implement the repository layer"
- "Write tests for the service package"
- "Add concurrent processing for transformations"

**For API Design** - Talk to `api-designer`:
- "Create OpenAPI spec for mapping endpoints"
- "Design REST API for transformation service"
- "Define error response formats"

**For Architecture** - Talk to `microservices-architect`:
- "Define service boundaries"
- "Design resilience patterns"
- "Plan scaling strategy"

**For Knowledge** - Talk to `knowledge-synthesizer`:
- "What patterns are emerging?"
- "What are the best practices so far?"
- "Generate recommendations for improvement"

### 3. Collaborative Work Patterns

**Building a Feature:**
```
1. microservices-architect → Design approach
2. api-designer → Define API contract
3. sql-pro → Design data model
4. golang-pro → Implement code
5. knowledge-synthesizer → Extract learnings
```

**Optimizing Performance:**
```
1. knowledge-synthesizer → Identify bottlenecks
2. sql-pro → Optimize queries
3. golang-pro → Optimize Go code
4. microservices-architect → Review system-wide
```

## 📋 Next Steps

### Immediate Actions

1. **Review Team Documentation**
   ```bash
   cat .agents/README.md
   cat .agents/PROJECT_CONTEXT.md
   ```

2. **Initialize Project Structure**
   - Ask `microservices-architect` to define folder structure
   - Ask `golang-pro` to set up Go modules
   - Ask `sql-pro` to design database schema

3. **Define API Contracts**
   - Ask `api-designer` to create OpenAPI specification
   - Define request/response models
   - Establish error handling patterns

4. **Set Up Development Environment**
   - Initialize Go module
   - Set up MySQL database
   - Configure development tools

### Suggested First Tasks

#### Task 1: Database Schema Design
**Assign to**: `sql-pro`  
**Request**: "Design a MySQL schema for storing mapping configurations, transformation rules, and audit logs. Optimize for high-read performance."

#### Task 2: API Specification
**Assign to**: `api-designer`  
**Request**: "Create an OpenAPI 3.1 specification for a RESTful API to manage mappings and execute transformations."

#### Task 3: Project Scaffolding
**Assign to**: `golang-pro`  
**Request**: "Set up a Go project structure following clean architecture principles with api, service, and repository layers."

#### Task 4: Architecture Documentation
**Assign to**: `microservices-architect`  
**Request**: "Document the system architecture, define service boundaries, and establish communication patterns."

## 💡 Tips for Success

1. **Always Query Context First**
   - Before starting work, check `PROJECT_CONTEXT.md`
   - Use `context-manager` to retrieve shared state

2. **Coordinate Major Decisions**
   - Consult `microservices-architect` for system-level changes
   - Get `knowledge-synthesizer` input on patterns

3. **Follow Quality Standards**
   - >80% test coverage (golang-pro)
   - <100ms query time (sql-pro)
   - OpenAPI 3.1 compliance (api-designer)

4. **Document Learnings**
   - `knowledge-synthesizer` tracks all insights
   - Best practices are continuously extracted
   - Patterns are automatically identified

## 🔄 Agent Communication

### Standard Request Format
```json
{
  "requesting_agent": "golang-pro",
  "request_type": "get_database_schema",
  "payload": {
    "query": "Current schema for mappings table"
  }
}
```

### Progress Updates
```json
{
  "agent": "golang-pro",
  "status": "implementing",
  "progress": {
    "completed": ["repository layer", "unit tests"],
    "in_progress": ["service layer"],
    "pending": ["API handlers"]
  }
}
```

## 📚 Resources

- **Team Documentation**: `.agents/README.md`
- **Project Context**: `.agents/PROJECT_CONTEXT.md`
- **Team Configuration**: `.agents/team-manifest.json`
- **Agent Profiles**: `.agents/*.md`

## 🆘 Need Help?

1. **For technical questions**: Query specific agent (sql-pro, golang-pro, etc.)
2. **For project context**: Check with `context-manager`
3. **For architecture decisions**: Consult `microservices-architect`
4. **For best practices**: Ask `knowledge-synthesizer`

## ✨ Example Workflow

```bash
# 1. Review current context
cat .agents/PROJECT_CONTEXT.md

# 2. Ask microservices-architect for architecture guidance
"Design the overall system architecture for a mapping service"

# 3. Ask sql-pro to design database
"Create optimized MySQL schema for mapping configurations"

# 4. Ask api-designer for API spec
"Generate OpenAPI spec for mapping CRUD operations"

# 5. Ask golang-pro to implement
"Implement repository layer with GORM"

# 6. Ask knowledge-synthesizer for insights
"What patterns should we follow based on similar projects?"
```

---

## 🎉 You're Ready!

Your agent team is **loaded and active**. Start by assigning tasks to specific agents or asking the `microservices-architect` to create a comprehensive development plan.

**Team Status**: ✅ Active  
**Project**: ServiceNow Mapping Service  
**Technology**: Golang + MySQL  
**Loaded**: 2025-11-21

Start coding! 🚀
