# drumkit
## Overview

Drumkit is a web application designed for freight brokers to streamline their workflow by integrating with their Transportation Management System (TMS). For this project, Drumkit connects with Turvo, allowing users to view and create loads directly from the web app.

## Features

- **View Loads:** Fetch and display all loads from the connected Turvo account.
- **Create Loads:** Fill out a form to create new loads in Turvo.

## Architecture

- **Backend:** Go (Golang) REST API that communicates with Turvo using their API.
- **Frontend:** React with TypeScript for a modern, responsive UI.
- **Cloud Hosting:** Deployed on AWS for easy access and testing.

## Data Model

Drumkit follows the load object format as described in the [Drumkit API Reference](https://drumkit.readme.io/reference/post_integrations-webhooks-loads). Please refer to this documentation for details on required fields and structure.

## Setup & Running Locally

### Prerequisites

- Go (>=1.18)
- Node.js (>=16)
- Turvo sandbox credentials (provided separately)
- AWS account (for deployment)

### Backend

1. Clone the repository.
2. Set Turvo credentials in environment variables.
3. Run:
    ```bash
    go run main.go
    ```

### Frontend

1. Navigate to the `frontend` directory.
2. Install dependencies:
    ```bash
    npm install
    ```
3. Start the development server:
    ```bash
    npm start
    ```

### Deployment

- Deploy backend and frontend to AWS (e.g., using Elastic Beanstalk for Go and S3/CloudFront for React).
- Update environment variables with production Turvo credentials.

## Usage

- **View Loads:** Navigate to the Loads page to see all shipments from Turvo.
- **Create Load:** Use the "New Load" form, fill in required details, and submit to create a load in Turvo.

## Inputs & Outputs

- **Inputs:** Load details via form (origin, destination, goods, carrier, etc.).
- **Outputs:** List of loads, confirmation of successful load creation, error messages if any issues occur.

## Testing

- Unit and integration tests are included for both backend and frontend.
- Run backend tests:
  ```bash
  go test ./...
  ```
- Run frontend tests:
  ```bash
  npm test
  ```

## Submission

- Add [@jinyanzang](https://github.com/jinyanzang) and [@dhruv4](https://github.com/dhruv4) as contributors.
- Ensure the README covers setup, usage, and testing instructions.

## Resources

- [Turvo API Documentation](https://app.turvo.com/lobby/documentation#tag/Shipments)
- [Drumkit API Reference](https://drumkit.readme.io/reference/post_integrations-webhooks-loads)

---

Feel free to reach out with any questions or clarifications!