# 🎵 Song Randomizer

A lightweight web app that picks a random song from a personal playlist and surfaces a direct Spotify link to it. Built with a Go backend, PostgreSQL database, and a plain HTML/CSS/JS frontend — all containerized with Docker.

---

## How It Works

1. Songs (title, artist, Spotify link) are stored in a PostgreSQL database, imported from a CSV file.
2. The Go backend exposes an API endpoint that queries a random row from the database.
3. The frontend fetches that data and displays the song with a clickable Spotify link.

---

## Tech Stack

| Layer    | Technology              |
|----------|-------------------------|
| Frontend | HTML, CSS, JavaScript (Fetch API) |
| Backend  | Go                      |
| Database | PostgreSQL              |
| DevOps   | Docker, Docker Compose  |

---

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/) and Docker Compose installed

### Setup

1. **Clone the repository**

   ```bash
   git clone https://github.com/AltynayA/songrandomizer.git
   cd songrandomizer
   ```

2. **Configure environment variables**

   Copy the example env file and fill in your values:

   ```bash
   cp .env.example .env
   ```

   Open `.env` and set your PostgreSQL credentials and any other required variables.

3. **Start the app**

   ```bash
   docker compose up --build
   ```

   This will spin up the database and backend, and serve the frontend.

4. **Open in your browser**

   Navigate to `http://localhost:<port>` (check `docker-compose.yml` for the configured port).

---

## Project Structure

```
songrandomizer/
├── backend/          # Go API server
├── init/             # Database initialization scripts (schema + CSV import)
├── index.html        # Frontend markup
├── script.js         # Fetch logic & DOM updates
├── style.css         # Styles
├── docker-compose.yml
├── .env.example
└── README.md
```

---

## Database

Songs are seeded from a CSV file via the `init/` directory scripts. Each record contains:

- **Title** — song name
- **Artist** — artist name
- **Link** — Spotify URL

To use your own playlist, replace the CSV with your own data in the same format before running `docker compose up`.

---

## License

This project is open source and available for personal use.