```mermaid
erDiagram
    Users ||--o{ Posts : creates
    Users ||--o{ Comments : writes
    Users ||--o{ Reactions : makes
    Users ||--o{ Sessions : has
    Posts ||--o{ Comments : receives
    Posts ||--o{ Reactions : receives
    Posts }|--|{ Post_Categories : has
    Categories }|--|{ Post_Categories : has

    Users {
        int id PK
        string username
        string email UK
        string password
    }

    Posts {
        int id PK
        int user_id FK
        string title
        text content
    }

    Categories {
        int id PK
        string name UK
    }

    Comments {
        int id PK
        int post_id FK
        int user_id FK
        text content
    }

    Reactions {
        int id PK
        int post_id FK
        int user_id FK
        string type "Possible values: like, dislike"
        string constraint "Unique: (post_id, user_id)"
    }

    Sessions {
        uuid session_id PK
        int user_id FK
        datetime expires_at
    }

    Post_Categories {
        int post_id FK
        int category_id FK
    }