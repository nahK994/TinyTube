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

<!-- ## Project Milestones 🏁📅

### Milestone 1: Initial Setup 🔨
- Set up monorepo structure 🗂️
- Implement Auth Service 🔑
- Implement User Management Service 👤

### Milestone 2: Video Upload and Transcoding 📹➡️📂
- Build Video Upload Service 🎬
- Add transcoding for multiple formats 🔄📺
- Store video metadata in PostgreSQL 📑🐘

### Milestone 3: Streaming Service 📡🎥
- Implement HLS or DASH for video streaming 🖥️💨
- Ensure scalability of streaming service 🚀🌍

### Milestone 4: User Interaction and Real-Time Updates 💬💖
- Implement like, comment, and subscribe features 👍💬🔔
- Integrate WebSockets for real-time feedback 📡🔄 -->

## Services 🛠️

1. **Auth Service** 👥🔐
   - Manages user authentication, registration, and access tokens. 🔑🔒

2. **User Management Service** 👥🔐
   - Handles user profile information, account settings, and possibly subscription details. 🖼️📧

3. **User Interaction Service** 💬👍
   - Manages comments, likes, views, and other user interactions. ❤️💬🔔

4. **Recommendation Service** 🤖🎯
   - Provides personalized video recommendations based on user behavior and preferences. 📈🎬

5. **Video Upload Service** ⬆️📹➡️📂
   - Manages the uploading, transcoding into multiple formats, data replication and storage of video files. 🖥️➡️📺

6. **Video Streaming Service** 📡🎥
   - HLS or DASH video streaming 🔄📺

7. **Video Metadata Service** 🗂️📝
   - Stores and retrieves metadata (e.g., title, description, tags, timestamps) associated with videos. 🖼️📌

8. **Notification Service** 🔔📲
   - Sends notifications related to new video uploads, comments, or other relevant updates. 📬📩

9. **Analytics Service** 📊📈
   - Tracks user engagement metrics, video performance, and other analytics. 📉🔍

10. **RabbitMQ Service** 🐇📬
   - Facilitates asynchronous communication between services, especially useful for tasks like video processing and notification dispatch. 📤🔄

## How to Run the Project 🏃‍♂️🛠️
_(Instructions on running services locally with Docker Compose, etc. coming soon... Stay tuned! 📻)_

## Contributing 🤝💻
Feel free to open issues 🐛📥 and submit pull requests 🚀! See our [CONTRIBUTING.md](CONTRIBUTING.md) for more info on how to join the TinyTube journey! 🎉
