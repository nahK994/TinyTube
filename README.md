# TinyTube 🎥🚀

**TinyTube** is a lightweight video sharing platform, kind of like YouTube but with minimal features and a cool microservice-based architecture. 📦✨

## Tech Stack 💻🔧
- **Backend**:
  - Golang (User, Upload, Streaming, etc.)
  - Python 🐍 (Recommendation Service)
  - PostgreSQL 🐘 & MongoDB for databases
  - RabbitMQ 🐰📬 for event-driven architecture
  - WebSockets 🔄 for real-time interactions
- **Frontend**: Not yet decided 🤔 (TBD, any ideas? 🧠💡)

## Project Milestones 🏁📅

### Milestone 1: Initial Setup 🔨
- Set up monorepo structure 🗂️
- Implement basic User Management Service 👤🔑
- Set up PostgreSQL database 🐘

### Milestone 2: Video Upload and Transcoding 📹➡️📂
- Build Video Upload Service 🎬
- Add transcoding for multiple formats 🔄📺
- Store video metadata in PostgreSQL 📑🐘

### Milestone 3: Streaming Service 📡🎥
- Implement HLS or DASH for video streaming 🖥️💨
- Ensure scalability of streaming service 🚀🌍

### Milestone 4: User Interaction and Real-Time Updates 💬💖
- Implement like, comment, and subscribe features 👍💬🔔
- Integrate WebSockets for real-time feedback 📡🔄

## Services 🛠️

1. **User Management Service** 👥🔐
   - JWT-based authentication 🔑🔒
   - Profile management (name, email, profile pic, etc.) 🖼️📧

2. **User Interaction Service** 💬👍
   - Likes, comments, subscriptions ❤️💬🔔

3. **Recommendation Service** 🤖🎯
   - Video suggestions based on user behavior and viewing history 📈🎬

4. **Video Upload & Transcoding Service** ⬆️📹➡️📂
   - Transcoding videos into multiple formats 🖥️➡️📺

5. **Video Streaming Service** 📡🎥
   - HLS or DASH video streaming 🔄📺

6. **Thumbnail Management Service** 🖼️📸
   - Generate or upload thumbnails for uploaded videos 📂🖼️

## How to Run the Project 🏃‍♂️🛠️
_(Instructions on running services locally with Docker Compose, etc. coming soon... Stay tuned! 📻)_

## Contributing 🤝💻
Feel free to open issues 🐛📥 and submit pull requests 🚀! See our [CONTRIBUTING.md](CONTRIBUTING.md) for more info on how to join the TinyTube journey! 🎉
