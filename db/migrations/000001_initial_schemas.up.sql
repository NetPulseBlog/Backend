CREATE TYPE user_account_type_enum as ENUM ('personal', 'system_topic');

CREATE TABLE "user"
(
    "id"                 uuid PRIMARY KEY       NOT NULL,
    "encrypted_password" character varying               DEFAULT null,

    "created_at"         TIMESTAMP              NOT NULL DEFAULT now(),
    "updated_at"         TIMESTAMP                       DEFAULT now(),

    "type"               user_account_type_enum NOT NULL DEFAULT 'personal',
    "email"              character varying,
    "name"               character varying               DEFAULT null,
    "description"        character varying               DEFAULT null,

    "avatar_url"         character varying               DEFAULT null,
    "cover_url"          character varying               DEFAULT null
);

CREATE TYPE news_line_default_enum as ENUM ('nld_popular', 'nld_fresh');

CREATE TYPE news_line_sort_enum as ENUM ('nls_by_popular', 'nls_by_date');

CREATE TABLE "user_settings"
(
    "user_id"           uuid                   NOT NULL,
    "news_line_default" news_line_default_enum NOT NULL DEFAULT 'nld_popular',
    "news_line_sort"    news_line_sort_enum    NOT NULL DEFAULT 'nls_by_popular',

    FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE
);

CREATE TABLE "user_auth"
(
    "id"            uuid              NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    "user_id"       uuid              NOT NULL,

    "refresh_token" character varying NOT NULL,
    "access_token"  character varying NOT NULL,

    "device_name"   character varying NOT NULL,

    "expires_at"    TIMESTAMP         NOT NULL,
    "created_at"    TIMESTAMP         NOT NULL             DEFAULT now(),

    FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE
);

CREATE TYPE bookmark_type_enum as ENUM ('bt_article', 'bt_comment');

CREATE TABLE "user_bookmark"
(
    "user_id"       uuid               NOT NULL,
    "resource_id"   uuid               NOT NULL,
    "resource_type" bookmark_type_enum NOT NULL,
    "created_at"    TIMESTAMP          NOT NULL DEFAULT now(),

    FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE
);

CREATE TABLE "user_subscription"
(
    "subscriber_id"      uuid      NOT NULL,
    "subscribed_user_id" uuid      NOT NULL,

    "created_at"         TIMESTAMP NOT NULL DEFAULT now(),

    FOREIGN KEY ("subscriber_id") REFERENCES "user" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("subscribed_user_id") REFERENCES "user" ("id") ON DELETE CASCADE
);

CREATE TYPE article_status as ENUM ('published', 'draft');

CREATE TABLE "article"
(
    "id"           uuid              NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    "author_id"    uuid              NOT NULL,

    "status"       article_status    NOT NULL             DEFAULT 'draft',

    "created_at"   TIMESTAMP         NOT NULL             DEFAULT now(),
    "updated_at"   TIMESTAMP                              DEFAULT now(),

    title          character varying NOT NULL,
    topic_id       uuid,

    content_blocks jsonb             NOT NULL             DEFAULT '{}',
    cover_url      character varying,
    SubTitle       character varying,

    views_count    int               NOT NULL             DEFAULT 0,

    FOREIGN KEY ("author_id") REFERENCES "user" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("topic_id") REFERENCES "user" ("id") ON DELETE CASCADE
);

CREATE TABLE "article_comment"
(
    "id"               uuid              NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    "reply_comment_id" uuid,

    "article_id"       uuid              NOT NULL,
    "author_id"        uuid              NOT NULL,

    "created_at"       TIMESTAMP         NOT NULL             DEFAULT now(),
    "updated_at"       TIMESTAMP                              DEFAULT now(),

    "content"          character varying NOT NULL,

    "is_edited"        bool              NOT NULL             DEFAULT false,

    FOREIGN KEY ("article_id") REFERENCES "article" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("reply_comment_id") REFERENCES "article_comment" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("author_id") REFERENCES "user" ("id") ON DELETE CASCADE
);
