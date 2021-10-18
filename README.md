# OTERON

This is a simple app that converts SMILES (Simplified Molecular-Input Line-Entry System) format into IUPAC (International Union of Pure and Applied Chemistry) format

## Requirements

- Go v1.17
- Node v16.9.1

## Directory structure

- `backend`

  Package for API server, written in Golang

- `frontend`

  Package for frontend, React app

## How to configure environment

- Make a file `frontend/.env` as a copy of `frontend/.env.example`. \
  Set `REACT_APP_SERVER_API_URL=<SERVER_API_URL>`. For example, `REACT_APP_SERVER_API_URL="http://localhost:8080/api"`.

## How to run for development

- In `backend` directory, run the following command to run server in development mode

  ```
  go mod tidy   # install dependencies
  go run .      # build and start server
  ```

- In `frontend` directory, run the following command to start app

  ```
  npm install   # install dependencies
  npm start     # start development server
  ```

## How to run tests

- In `backend` directory, run the following command

  ```
  go test ./tests      # unit tests
  ```

- In `frontend` directory, run the following command

  - Run unit tests
    ```
    npm run test
    ```
  - Run e2e tests

    - Start backend and frontend
    - Run `npm run cypress:run` to start cypress app. Cypress window will open.
    - Start tests by clicking a test file or clicking on `Run` button.

## Current progress

### Front end

- Set up front end with create-react-app
- Uses TypeScript.
- Uses REST API endpoint to request SMILES to IUPAC nomenclature conversion to back end

### Back end

- Set up back end with Gin framework of Go language.
- Built engine to convert SMILES TO IUPAC. API part (using GIN) uses this engine.
- Engine only handles straight-chain and branched alkanes. i.e. no cyclics.
