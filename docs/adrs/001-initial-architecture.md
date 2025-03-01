# ADR 001: Initial Architecture Decision

## Status
Accepted

## Context
Bot Racer needs to be designed as a real-time racing game that supports both human and bot players, with the following key requirements:

1. Game Environment:
   - Procedurally generated tracks using GUID inputs for reproducibility
   - Tracks that loop back after approximately 30 seconds of racing
   - Support for multiple laps
   - Top-down view perspective

2. User Interface:
   - Mobile-friendly web interface
   - Terminal-based client for *NIX systems
   - Car positioned near bottom center of view (web) or center of viewport (terminal)
   - Support for keyboard input (arrow keys) and touch (on-screen buttons) in web client
   - Support for keyboard input in terminal client
   - Real-time updates of race state

3. Multiplayer:
   - Support for both human and bot players
   - Real-time race state synchronization
   - Fair and consistent racing environment

## Decision

### System Architecture
We will implement a client-server architecture with the following components:

1. Backend (Go):
   - WebSocket server for real-time communication
   - Track generation system using GUIDs as seeds
   - Race state management
   - PostgreSQL database integration
   - Bot API interface
   - Common protocol for web, terminal, and bot clients

2. Frontend:
   a. Web Client:
      - Mobile-responsive design
      - WebSocket client for real-time updates
      - Canvas-based rendering for track and racers
      - Touch and keyboard input handling

   b. Terminal Client (Go):
      - ASCII/Unicode art rendering of track and racers
      - Terminal UI framework (e.g., Bubble Tea)
      - Keyboard input handling
      - Support for different terminal sizes
      - Configurable refresh rate

3. Database (PostgreSQL):
   - User accounts and authentication
   - Bot registrations and metadata
   - Race history and records
   - Track seeds and metadata

### Key Technical Decisions

1. Communication Protocol:
   - WebSocket for real-time game state updates
   - REST API for non-real-time operations (user management, history, etc.)
   - JSON message format for all communications

2. Track Generation:
   - Deterministic procedural generation using GUID as seed
   - Fixed-width track with varying curvature
   - Guaranteed loop completion after ~30 seconds of racing
   - Track difficulty metadata for matchmaking
   - Different generation algorithms for local and server play:
     * Local play: Consistent generation across all clients for same GUID
     * Server play: Server-specific generation to prevent pre-learning tracks
     * No correlation between local and server tracks with same GUID
   - Clients include track generation capability for offline/local play
   - Server maintains its own track generation state separate from clients

3. Game State Management:
   - Server-authoritative architecture
   - Client prediction for smooth gameplay
   - Regular state synchronization
   - Conflict resolution favoring server state

## Consequences

### Positive
1. WebSocket provides low-latency, bi-directional communication
2. Go backend offers good performance and concurrency support
3. PostgreSQL provides robust data persistence and querying
4. Multiple client options increase accessibility
5. Terminal client enables headless testing and development
6. Canvas-based rendering allows efficient updates and scaling
7. Deterministic track generation enables fair competition
8. Separate local/server track generation prevents exploitation
9. Offline play support increases accessibility

### Negative
1. Requires managing WebSocket connection state and reconnection
2. Need to handle varying network conditions and latency
3. Must ensure fair play between different input methods
4. Track generation needs careful tuning for consistent difficulty
5. Users may be confused by different tracks from same GUID in local vs server play
6. Additional complexity in maintaining separate track generation systems

### Risks
1. Network latency affecting gameplay fairness
2. Browser compatibility issues with WebSocket/Canvas
3. Mobile input lag compared to keyboard
4. Potential for race conditions in multiplayer scenarios

## Compliance
This ADR complies with the project constitution by:
1. Establishing clear interfaces between components
2. Supporting both human and bot players fairly
3. Providing reproducible game states
4. Following best practices for separation of concerns

## Implementation Notes
Initial implementation will focus on:
1. Basic track generation system
2. Single-player functionality
3. Core WebSocket communication
4. Essential database schema
5. Basic terminal client implementation

Future ADRs will detail:
1. Specific track generation algorithms
   - Local generation algorithm
   - Server generation algorithm
   - Difficulty metrics and controls
2. Bot API specifications
3. Authentication system
4. Matchmaking system
5. Terminal rendering approach
6. Cross-client fairness measures
