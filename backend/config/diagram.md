# TinyTube Database Schema Description

This document provides an overview of the database schema for TinyTube, a distributed video storage and sharing platform. Each table is described with its purpose, key columns, and relationships.

## Tables

### 1. **Users**

Stores information about each registered user on the platform.

| Column      | Type         | Description                        |
|-------------|--------------|------------------------------------|
| `id`        | UUID         | Primary key, unique identifier for each user. |
| `username`  | VARCHAR(50)  | Username chosen by the user, must be unique. |
| `email`     | VARCHAR(100) | Email address, unique and indexed. |
| `password`  | TEXT         | Hashed password for user authentication. |
| `created_at`| TIMESTAMP    | Timestamp for when the user account was created. |

- **Relationships**: 
  - `Videos`: One-to-many (one user can upload multiple videos).
  - `Likes`: One-to-many (one user can like multiple videos).
  - `Comments`: One-to-many (one user can comment on multiple videos).

---

### 2. **Videos**

Stores details about each uploaded video.

| Column         | Type         | Description                                 |
|----------------|--------------|---------------------------------------------|
| `id`           | UUID         | Primary key, unique identifier for each video. |
| `user_id`      | UUID         | Foreign key referencing `Users(id)`, indicating the uploader. |
| `title`        | VARCHAR(200) | Title of the video.                         |
| `description`  | TEXT         | Description of the video content.           |
| `file_path`    | TEXT         | Path to the video file in storage.          |
| `thumbnail_path`| TEXT        | Path to the thumbnail image in storage.     |
| `views`        | INT          | Count of views for the video.               |
| `uploaded_at`  | TIMESTAMP    | Timestamp for when the video was uploaded.  |

- **Relationships**:
  - `Users`: Many-to-one (each video is uploaded by a single user).
  - `Comments`: One-to-many (each video can have multiple comments).
  - `Likes`: One-to-many (each video can have multiple likes).

---

### 3. **Comments**

Stores comments left by users on videos.

| Column        | Type         | Description                                  |
|---------------|--------------|----------------------------------------------|
| `id`          | UUID         | Primary key, unique identifier for each comment. |
| `video_id`    | UUID         | Foreign key referencing `Videos(id)`.        |
| `user_id`     | UUID         | Foreign key referencing `Users(id)`.         |
| `content`     | TEXT         | Text content of the comment.                 |
| `created_at`  | TIMESTAMP    | Timestamp for when the comment was posted.   |

- **Relationships**:
  - `Users`: Many-to-one (each comment is made by one user).
  - `Videos`: Many-to-one (each comment belongs to one video).

---

### 4. **Likes**

Tracks likes on videos by users.

| Column      | Type         | Description                                 |
|-------------|--------------|---------------------------------------------|
| `id`        | UUID         | Primary key, unique identifier for each like. |
| `video_id`  | UUID         | Foreign key referencing `Videos(id)`.       |
| `user_id`   | UUID         | Foreign key referencing `Users(id)`.        |
| `created_at`| TIMESTAMP    | Timestamp for when the like was given.      |

- **Relationships**:
  - `Users`: Many-to-one (each like is given by one user).
  - `Videos`: Many-to-one (each like is associated with one video).

---

### 5. **Video_Metadata**

Stores metadata for videos, including length, resolution, and codec.

| Column         | Type        | Description                                 |
|----------------|-------------|---------------------------------------------|
| `id`           | UUID        | Primary key, unique identifier for each metadata entry. |
| `video_id`     | UUID        | Foreign key referencing `Videos(id)`.       |
| `duration`     | INT         | Duration of the video in seconds.           |
| `resolution`   | VARCHAR(20) | Resolution of the video (e.g., 1080p, 720p).|
| `codec`        | VARCHAR(20) | Codec used for encoding the video.          |
| `size`         | INT         | Size of the video file in bytes.            |

- **Relationships**:
  - `Videos`: One-to-one (each video has one metadata entry).

---

## Relationships Overview

- **Users - Videos**: One-to-many
- **Users - Comments**: One-to-many
- **Users - Likes**: One-to-many
- **Videos - Comments**: One-to-many
- **Videos - Likes**: One-to-many
- **Videos - Video_Metadata**: One-to-one

---

This schema supports the basic operations for TinyTube, including uploading videos, user authentication, viewing, liking, and commenting on videos, and tracking video metadata.
