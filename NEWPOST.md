# New Post Workflow

Use this checklist when adding a new article.

## 1. Create the Bundle

```bash
hugo new posts/my-post/index.md
```

This creates a leaf page bundle:

```text
content/posts/my-post/
├── index.md
└── cover.jpg
```

Put the article Markdown in `index.md`. Put the social preview image next to it as `cover.jpg`.

## 2. Prepare Cover Image

Rules:

- file name: `cover.jpg`
- location: same folder as `index.md`
- size: `1200x630`
- format: JPEG
- keep text inside the image minimal
- make the image readable when cropped in Telegram previews

Hugo will process this image for Open Graph and Twitter Cards.

## 3. Fill Front Matter

```yaml
---
title: "Post title with keyword"
date: 2026-04-29
description: "Meta description up to 160 characters with the main keyword and a concrete reason to click."
tags:
  - hugo
  - cloudflare
topics:
  - static-sites
series:
  - blog-platform
keywords:
  - Hugo SEO
  - static blog
aliases:
  - /old-url-if-needed/
draft: false
cover: "cover.jpg"
---
```

Rules:

- `title` is required and must be 60 characters or less
- `description` is required and must be 160 characters or less
- `date` is required
- `tags`, `topics`, `series`, and `keywords` must not be empty
- `aliases` must include old URLs when renaming or replacing a post
- `draft` must be `false` before publishing
- `cover` must be `cover.jpg`

## 4. Write the Article

Content rules:

- Put the main keyword in the title and first 100 words.
- Use one search intent per post.
- Use H2/H3 sections after the page title.
- Add internal links to related posts.
- Use descriptive anchor text.
- Add alt text to every Markdown image.
- Keep external links as normal Markdown links; Hugo adds safe `rel` attributes.

Example image:

```markdown
![Architecture diagram for Hugo deployment](diagram.jpg)
```

Example internal link:

```markdown
Read the [Cloudflare Pages Hugo deployment guide](/posts/cloudflare-pages-hugo/).
```

## 5. Validate Before Publishing

Validate only posts:

```bash
./scripts/verify-posts.sh
```

Validate the whole site:

```bash
./scripts/verify-static-site.sh
```

The checks verify front matter, cover dimensions, generated HTML, social preview tags, JSON-LD, redirects, render hooks, and absence of client-side JavaScript.

## 6. Preview Locally

```bash
hugo server
```

Open the local URL from Hugo and check:

- article title
- description
- cover image
- headings
- internal links
- related posts
- mobile readability

## 7. Publish

Commit and push. Cloudflare Pages will build the site with:

```bash
hugo --minify
```

After deploy, test the public URL in Telegram. If preview is stale, send the URL to `@WebpageBot` and refresh the preview cache.
