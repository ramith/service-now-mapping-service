# Project Context - ServiceNow Mapping Service

**Last Updated**: 2025-11-21T00:00:00Z  
**Managed by**: context-manager  
**Project Status**: Initialization Phase

## 📋 Project Information

### Basic Details
- **Project Name**: ServiceNow Mapping Service
- **Type**: Golang Microservice
- **Database**: MySQL 8.0+
- **API Style**: RESTful
- **Target Environment**: Cloud-native, containerized

### Business Domain
- **Purpose**: Data mapping and transformation for ServiceNow integration
- **Primary Users**: Enterprise integration teams
- **Key Features**:
  - Dynamic mapping configuration
  - Real-time data transformation
  - Audit trail and versioning
  - High-performance processing

## 🏗️ Technical Architecture

### System Components

#### 1. API Layer
- **Framework**: TBD (Chi/Gin/Echo)
- **Authentication**: API Key based
- **Rate Limiting**: Token bucket algorithm
- **Documentation**: OpenAPI 3.1

#### 2. Service Layer
- **Pattern**: Clean architecture
- **Responsibilities**:
  - Business logic orchestration
  - Mapping rule evaluation
  - Transformation execution
  - Validation and error handling

#### 3. Repository Layer
- **ORM**: TBD (GORM/sqlx)
- **Connection Pool**: Configured for high concurrency
- **Query Optimization**: Indexed queries, prepared statements
- **Transactions**: ACID compliance

#### 4. Database Schema
```
Tables:
- mappings: Core mapping definitions
- transformations: Transformation rules
- audit_logs: Change tracking
- api_keys: Authentication
```

## 👥 Agent Team Configuration

### Active Agents

#### sql-pro
- **Status**: Active
- **Current Focus**: Schema design, query optimization
- **Performance Targets**:
  - Query execution < 50ms
  - Index coverage > 90%
  - Zero N+1 queries

#### golang-pro
- **Status**: Active
- **Current Focus**: Service implementation, testing
- **Quality Targets**:
  - Test coverage > 80%
  - Zero race conditions
  - Idiomatic Go patterns

#### api-designer
- **Status**: Active
- **Current Focus**: REST API design, OpenAPI spec
- **Design Targets**:
  - Consistent naming conventions
  - Proper HTTP semantics
  - Comprehensive error responses

#### context-manager
- **Status**: Active
- **Current Focus**: Knowledge coordination, state tracking
- **Service Targets**:
  - Context retrieval < 100ms
  - 100% consistency
  - Complete audit trail

#### microservices-architect
- **Status**: Active
- **Current Focus**: System design, resilience patterns
- **Architecture Targets**:
  - 99.9% availability
  - Horizontal scalability
  - Graceful degradation

#### knowledge-synthesizer
- **Status**: Active
- **Current Focus**: Pattern extraction, best practices
- **Intelligence Targets**:
  - Pattern accuracy > 85%
  - Daily insights generation
  - Continuous improvement tracking

## 📊 Current State

### Project Phase: Initialization

#### Completed
- ✅ Agent team loaded
- ✅ Project structure defined
- ✅ Documentation initialized
- ✅ Quality standards established

#### In Progress
- 🔄 Database schema design (sql-pro)
- 🔄 API specification (api-designer)
- 🔄 Go module setup (golang-pro)
- 🔄 Architecture documentation (microservices-architect)

#### Pending
- ⏳ Service implementation
- ⏳ Testing framework setup
- ⏳ CI/CD pipeline
- ⏳ Container configuration
- ⏳ Monitoring setup

## 🎯 Performance Requirements

### API Performance
- **Response Time**: < 100ms (p99)
- **Throughput**: > 1000 req/s
- **Error Rate**: < 0.1%
- **Availability**: 99.9%

### Database Performance
- **Query Time**: < 50ms (p99)
- **Connection Pool**: 20-100 connections
- **Replication Lag**: < 1s
- **Index Hit Rate**: > 95%

### System Performance
- **CPU Usage**: < 70% average
- **Memory Usage**: < 80% average
- **Network Latency**: < 10ms internal
- **Disk I/O**: Optimized with caching

## 🔐 Security Requirements

### Authentication & Authorization
- API key-based authentication
- Role-based access control (future)
- Token expiration and rotation
- Rate limiting per key

### Data Security
- Encryption at rest (database)
- Encryption in transit (TLS)
- SQL injection prevention
- Input validation and sanitization
- Audit logging for all changes

### Compliance
- GDPR considerations
- Data retention policies
- Access logging
- Security headers

## 📦 Dependencies

### Core Dependencies (Planned)
```go
- github.com/go-chi/chi/v5         // HTTP router
- github.com/go-sql-driver/mysql   // MySQL driver
- github.com/jmoiron/sqlx          // SQL extensions
- github.com/golang-migrate/migrate // Database migrations
- github.com/sirupsen/logrus       // Structured logging
- github.com/prometheus/client_golang // Metrics
- github.com/stretchr/testify      // Testing
```

### Development Dependencies
```go
- github.com/golangci/golangci-lint // Linting
- github.com/golang/mock            // Mocking
- github.com/go-delve/delve         // Debugging
```

## 🚀 Development Workflow

### Phase 1: Foundation (Current)
1. Database schema design (sql-pro)
2. API specification (api-designer)
3. Project scaffolding (golang-pro)
4. Architecture documentation (microservices-architect)

### Phase 2: Implementation
1. Repository layer (golang-pro + sql-pro)
2. Service layer (golang-pro)
3. API layer (golang-pro + api-designer)
4. Testing (golang-pro)

### Phase 3: Integration
1. Database integration (sql-pro)
2. API testing (api-designer)
3. Performance testing (microservices-architect)
4. Security review (all agents)

### Phase 4: Deployment
1. Container setup (microservices-architect)
2. CI/CD pipeline (microservices-architect)
3. Monitoring setup (microservices-architect)
4. Documentation (knowledge-synthesizer)

## 📈 Success Metrics

### Code Quality
- Test coverage > 80%
- Linting score: 100%
- Code review approval: All agents
- Documentation coverage: 100%

### Performance
- API latency < 100ms (p99)
- Database queries < 50ms (p99)
- Throughput > 1000 req/s
- Zero race conditions

### Reliability
- Uptime > 99.9%
- Error rate < 0.1%
- MTTR < 30 minutes
- Zero data loss

### Developer Experience
- API documentation: Complete
- Examples: Comprehensive
- Onboarding time < 1 day
- Setup time < 30 minutes

## 🔄 Knowledge Management

### Context Storage
- **Location**: Managed by context-manager
- **Update Frequency**: Real-time
- **Retention**: Full history
- **Access Pattern**: < 100ms retrieval

### Pattern Recognition
- **Managed by**: knowledge-synthesizer
- **Analysis Frequency**: Daily
- **Insight Generation**: Continuous
- **Distribution**: All agents

### Best Practices
- **Source**: Team interactions
- **Validation**: Performance data
- **Evolution**: Continuous
- **Documentation**: Automated

## 📞 Agent Communication

### Standard Message Format
```json
{
  "requesting_agent": "agent-id",
  "request_type": "action_type",
  "timestamp": "2025-11-21T00:00:00Z",
  "payload": {
    "query": "specific request details"
  }
}
```

### Response Format
```json
{
  "responding_agent": "agent-id",
  "status": "success|error",
  "timestamp": "2025-11-21T00:00:00Z",
  "data": {},
  "metadata": {}
}
```

## 🎓 Learning & Improvement

### Continuous Improvement Areas
1. Query optimization techniques
2. Go concurrency patterns
3. API design best practices
4. Microservices resilience
5. Performance tuning strategies

### Knowledge Evolution
- Daily pattern analysis
- Weekly performance reviews
- Monthly architecture assessments
- Quarterly strategy updates

---

**Maintained by**: context-manager  
**Next Update**: After each significant milestone  
**Accessible to**: All team agents
