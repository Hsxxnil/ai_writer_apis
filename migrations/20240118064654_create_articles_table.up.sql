CREATE EXTENSION IF NOT EXISTS "pg_trgm";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE articles
(
    id                 UUID    NOT NULL PRIMARY KEY,
    -- 文章設定
    content_type       text    not null,               -- 內容形式
    forum              text    not null,               -- 論壇
    board              text    not null,               -- 版位
    type               text    not null,               -- 文章類型
    style              text    not null,               -- 文章風格
    sponsorship        text    not null,               -- 業配程度
    title              text    not null,               -- 文章標題
    word_limit         integer not null,               -- 文章字數
    key_message        text    not null,               -- 關鍵訊息
    story              text    not null,               -- 口碑切角
    -- 人物設定
    gender             text    not null,               -- 人物性別
    from_age           integer not null,               -- 人物年齡(起)
    to_age             integer not null,               -- 人物年齡(迄)
    character_trait    text,                           -- 人物特質
    character_remarks  text,                           -- 人物細部描述
    -- 產品資訊
    product_name       text    not null,               -- 產品名稱
    product_feature    text    not null,               -- 產品特性
    product_highlights text    not null,               -- 產品亮點
    has_comparative    boolean not null default false, -- 是否有競品比較

    ai_article         text,                           -- 生成文章
    modify_article     text,                           -- 修改後的文章
    rating             integer,                        -- 文章分數
    created_at         TIMESTAMP        default now(),
    created_by         UUID,
    updated_at         TIMESTAMP,
    updated_by         UUID,
    deleted_at         TIMESTAMP
);

create index idx_articles_id
    on articles using hash (id);

create index idx_articles_forum
    on articles (forum);

create index idx_articles_board
    on articles (board);

create index idx_articles_type
    on articles (type);

create index idx_articles_style
    on articles (style);

create index idx_articles_sponsorship
    on articles (sponsorship);

create index idx_articles_title
    on articles using gin (title gin_trgm_ops);

create index idx_articles_word_limit
    on articles (word_limit);

create index idx_articles_key_message
    on articles using gin (key_message gin_trgm_ops);

create index idx_articles_story
    on articles using gin (story gin_trgm_ops);

create index idx_articles_gender
    on articles (gender);

create index idx_articles_from_age
    on articles (from_age);

create index idx_articles_to_age
    on articles (to_age);

create index idx_articles_character_trait
    on articles using gin (character_trait gin_trgm_ops);

create index idx_articles_character_remarks
    on articles using gin (character_remarks gin_trgm_ops);

create index idx_articles_product_name
    on articles using gin (product_name gin_trgm_ops);

create index idx_articles_product_feature
    on articles using gin (product_feature gin_trgm_ops);

create index idx_articles_product_highlights
    on articles using gin (product_highlights gin_trgm_ops);

create index idx_articles_has_comparative
    on articles (has_comparative);

create index idx_articles_ai_article
    on articles using gin (ai_article gin_trgm_ops);

create index idx_articles_modify_article
    on articles using gin (modify_article gin_trgm_ops);

create index idx_articles_rating
    on articles (rating);

create index idx_articles_created_at
    on articles (created_at desc);

create index idx_articles_created_by
    on articles using hash (created_by);

create index idx_articles_updated_at
    on articles (updated_at desc);

create index idx_articles_updated_by
    on articles using hash (updated_by);




