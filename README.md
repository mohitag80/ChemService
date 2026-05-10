# ChemService

Microservice for Chemistry exam questions for Grade 9-12 students. Part of the **ExamPlatform** suite.

## Tech Stack

- **Language**: Go 1.17
- **Framework**: Gin v1.7.0
- **Build Tool**: Go modules
- **Container**: Docker

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| GET | `/api/chem/questions/grade/{grade}/top/{n}` | Top N questions for a grade |
| GET | `/api/chem/questions/topic/{topic}/count/{n}` | N questions by topic |
| GET | `/api/chem/questions/complexity/{complexity}/count/{n}` | N questions by complexity |
| GET | `/api/chem/topics` | List all available topics |

## Supported Topics

- Atomic Structure
- Periodic Table
- Chemical Bonding
- Chemical Reactions
- Acids and Bases
- Mole Concept
- Thermodynamics
- Equilibrium
- Electrochemistry
- Organic Chemistry
- Polymers
- Biomolecules
- Industrial Chemistry

## Complexity Levels

- `easy` — 1 mark
- `medium` — 2 marks
- `hard` — 3 marks

## Running Locally

```bash
go mod tidy
go run main.go
```

Service starts on port **8083**.

### Example Requests

```bash
# Top 5 questions for Grade 10
curl http://localhost:8083/api/chem/questions/grade/10/top/5

# 3 Organic Chemistry questions
curl http://localhost:8083/api/chem/questions/topic/Organic%20Chemistry/count/3

# 4 hard questions
curl http://localhost:8083/api/chem/questions/complexity/hard/count/4
```

## Docker

```bash
docker build -t chem-service:1.0.0 .
docker run -p 8083:8083 chem-service:1.0.0
```
