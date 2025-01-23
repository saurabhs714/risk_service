# Risk Service

This project is a Go-based HTTP service for managing risks. It provides the following features:
- Create a new risk.
- Retrieve a list of all risks.
- Fetch a specific risk by its ID.

The service stores risks in memory and uses JSON for data exchange.

---

## **Features**
1. **Endpoints**:
   - `GET /v1/risks`: Fetch all risks.
   - `POST /v1/risks`: Create a new risk.
   - `GET /v1/risks/{id}`: Fetch a specific risk by ID.

2. **Data Structure**:
   - `ID`: UUID auto-generated on creation.
   - `Title`: Risk title.
   - `Description`: Risk description.
   - `State`: Must be one of `open`, `closed`, `accepted`, `investigating`.

---

## **Running the Service**

### Prerequisites:
- Go 1.18+ installed.
- `GOROOT` and `GOPATH` configured.

### Steps to Run:
1. Clone this repository(BRANCH NAME :- final_branch):
   ```bash
   git clone https://github.com/your-repo/risk_service.git
   cd risk_service
