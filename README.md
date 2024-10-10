# TinyTube 🎥

**TinyTube** is a lightweight video sharing platform, similar to YouTube but with minimal features and a microservice-based architecture.

## Tech Stack
- **Backend**:
  - Golang (User, Upload, Streaming, etc.)
  - Python (Recommendation Service)
  - PostgreSQL & MongoDB for databases
  - Apache Kafka for event-driven architecture
  - WebSockets for real-time interactions
- **Frontend**: TBA (React, Vue, etc.)
  
## Services
1. **User Management Service**
   - JWT-based authentication
   - Profile management
2. **User Interaction Service**
   - Likes, comments, subscriptions
3. **Recommendation Service**
   - Video suggestions based on user behavior
4. **Video Upload & Transcoding Service**
   - Transcoding videos into multiple formats
5. **Video Streaming Service**
   - HLS or DASH streaming
6. **Thumbnail Management Service**
   - Thumbnail creation and storage

## How to Run the Project
_(Instructions on running services locally with Docker Compose, etc.)_

## Contributing
Feel free to open issues and submit pull requests! See our [CONTRIBUTING.md](CONTRIBUTING.md) for more information.
