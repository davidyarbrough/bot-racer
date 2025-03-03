# Bot Racer

Bot Racer is a racing game where humans and bots compete on procedurally generated tracks, either solo or against each other. The game features both a web interface and a terminal-based client for *NIX systems.

## Features

- Procedurally generated tracks using GUIDs as seeds
- Support for both human and bot racers
- Multiple game modes: solo practice, multiplayer races
- Web and terminal-based interfaces
- Local offline play and online multiplayer
- Fair competition between humans and bots

## Clients

### Web Client
- Mobile-friendly web interface
- Real-time racing with WebSocket communication
- Touch controls for mobile, keyboard controls for desktop
- Canvas-based rendering with smooth animations

### Terminal Client
- ASCII/Unicode art rendering
- Keyboard controls using arrow keys
- Support for different terminal sizes
- Real-time updates with configurable refresh rate

## Architecture

The system uses a client-server architecture with:
- Go backend with PostgreSQL database
- WebSocket communication for real-time updates
- REST API for non-real-time operations
- Common protocol for web, terminal, and bot clients

## Track Generation

Tracks are procedurally generated using GUIDs as seeds. Note that:
- Local play tracks are consistent across all clients for the same GUID
- Server-generated tracks are unique to prevent pre-learning
- The same GUID will generate different tracks in local vs server play

## Glossary

### Game Elements

#### Track
- 9-lane grid view
- Player's vehicle positioned on bottom row
- 10 rows visible above player
- Scrolling view as player progresses
- Track loops back after 30 seconds of racing

#### Instrument Panel
- Located below the track view
- Contains speedometer showing current speed, position, and lap
- Additional instruments may be added in future versions

#### Hazards

##### Trees 🌳
- Static obstacle
- Appear randomly on track
- Collision causes immediate crash

##### Chickens 🐔
- Mobile obstacle
- Walks across the track
- Motives unclear

## Constitutional Development

This project employs Constitutional Development, where all code changes must comply with the project's constitution.

### What is Constitutional Development?

Constitutional Development is a structured approach where:

1. All code changes and system behaviors must comply with a written constitution
2. The constitution is maintained in `/docs/constitution/`
3. AI assistants and developers must verify compliance with the constitution
4. Changes that would violate the constitution are not permitted

### Amendment Process

1. Amendments can introduce new sections/subsections or revise existing ones
2. When a section is amended, the previous version becomes invalid
3. A current, consolidated version of the constitution is maintained at `/docs/constitution/constitution.md`
4. The original constitution is preserved for historical reference

### Purpose

This approach ensures:
- Consistent system behavior
- Clear governance of AI development
- Traceable evolution of system rules
- Protected core principles
- Controlled modification of fundamental rules

## Project Status

This project is under active development. The initial constitution has been established and serves as the foundation for all future development.
