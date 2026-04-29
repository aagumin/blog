# Arsen Gumin Blog

Static personal blog built with Hugo for Cloudflare Pages. It uses Markdown content, semantic HTML, one small local CSS file, no client-side JavaScript, automatic sitemap, robots.txt, RSS, canonical URLs, Open Graph, Twitter Cards, and BlogPosting JSON-LD.

## Local Development

```bash
hugo server
```

Open the local URL printed by Hugo. Drafts are excluded by default.

## Create a Post

```bash
hugo new posts/my-post/index.md
```

This creates a leaf page bundle:

```text
content/posts/my-post/
├── index.md
└── cover.jpg
```

Use this front matter:

```yaml
---
title: "Post title with keyword"
date: 2026-04-28
description: "Meta description between 120 and 160 characters with the main keyword and a concrete reason to click."
tags:
  - tag1
  - tag2
topics:
  - topic1
series:
  - series-name
keywords:
  - keyword one
aliases:
  - /old-url/
draft: false
cover: "cover.jpg"
---
```

Rules:

- keep `title` under 60 characters
- keep `description` under 160 characters
- put the main keyword in the title and first 100 words
- use one clear H1 from the page title, then H2/H3 sections
- link to at least one related post with descriptive anchor text
- add meaningful `alt` text to every image
- use a 1200x630 `cover` image for social previews when publishing an article
- keep article images in the same bundle as `index.md`

## Content Structure

Good technical articles usually follow this shape:

1. State the problem and target reader.
2. Explain why the problem matters.
3. Show the practical implementation or decision.
4. Discuss trade-offs and failure modes.
5. Link to related articles and next steps.

Write narrow posts. A focused article like "Deploy Hugo Blog to Cloudflare Pages" is easier to rank and easier to share than a broad article like "Everything About Static Sites".

## SEO Best Practices

- Use a unique title and description for every page.
- Keep descriptions near 120-160 characters so Telegram, Google, and social cards can show a useful preview.
- Match search intent in the introduction.
- Keep URLs readable: `/posts/my-post/`.
- Use semantic HTML through the Hugo templates: `main`, `article`, `header`, `nav`, `footer`.
- Add internal links between related posts.
- Use `topics`, `series`, and `keywords` so Hugo can generate taxonomy pages and better related posts.
- Add `aliases` when a URL changes; Cloudflare Pages receives generated `_redirects`.
- Keep published posts out of `draft: true`.
- Update `baseURL` in `config.toml` to the production domain before launch.

## Telegram Preview

Each post can define a cover image:

```yaml
cover: "/images/posts/serverless-blog-cover.jpg"
```

For page bundles, prefer:

```yaml
cover: "cover.jpg"
```

The SEO partial uses that image for `og:image` and `twitter:image`. If `cover` is a bundle resource, Hugo processes it to `1200x630` WebP automatically. If `cover` is missing, it falls back to `/images/default-og.jpg`. Hugo converts image paths to absolute URLs using `baseURL`, which Telegram needs for reliable previews.

Prepare cover images at `1200x630`. This ratio is the standard large social preview shape, works well in Telegram link cards, and gives enough room for a recognizable visual without cropping important details. Keep text inside the image minimal because Telegram already shows the article title and description from meta tags.

To check a preview:

1. Deploy the page or open a public preview URL from Cloudflare Pages.
2. Send the article URL to yourself in Telegram.
3. Confirm the card has the article title, description, and large image.

Telegram caches previews aggressively. To refresh the cache, open `@WebpageBot`, send the article URL, and ask it to update the preview. If the card still looks stale, check that `og:image` is absolute, publicly reachable, and returns a valid 1200x630 image.

## Personal Brand

The design intentionally follows the existing one-page site without duplicating its identity block: warm paper background, black border, mono typography, Georgia display headings, and hard shadows. Keep the same voice in articles:

- be specific about engineering decisions
- write from direct experience
- show trade-offs instead of generic advice
- keep author identity consistent across posts

Author metadata lives in `config.toml` under `[params]`.

## Build

```bash
hugo --minify
```

The generated site is written to `public/`.

## Verify

```bash
./scripts/verify-static-site.sh
```

This builds the site and checks key production invariants: sitemap, robots.txt, RSS, canonical URLs, article schema, Open Graph tags, Twitter Card tags, semantic article markup, author links, and absence of client-side JavaScript.

To validate only posts before publishing:

```bash
./scripts/verify-posts.sh
```

This checks every `content/posts/*/index.md` bundle for required front matter, `title` length up to 60 characters, `description` length up to 160 characters, non-empty tags/topics/series/keywords, aliases, `draft: false`, an existing JPEG `cover`, 1200x630 cover dimensions, and generated article HTML with canonical, BlogPosting JSON-LD, Open Graph, Twitter Card, and absolute social image URLs.

## Hugo Features Used

- **Leaf page bundles:** every post lives at `content/posts/<slug>/index.md` with local assets beside it.
- **Image processing:** cover images and Markdown images are processed by Hugo for social cards and responsive content images.
- **Render hooks:** Markdown images require `alt` text and get `loading="lazy"`, `decoding="async"`, dimensions, and `srcset`; external links get `rel="external noopener noreferrer"`.
- **Archetypes:** `archetypes/posts/index.md` creates the required post structure and front matter.
- **Taxonomies:** `tags`, `topics`, and `series` generate index pages and improve internal discovery.
- **Related content:** related posts use tags, topics, series, keywords, and date.
- **Aliases to redirects:** front matter `aliases` generate Cloudflare Pages `_redirects`.

## Cloudflare Pages

Use these settings:

- Build command: `hugo --minify`
- Output directory: `public`
- Environment variable: `HUGO_VERSION=0.161.0`

Deployment flow:

1. Push this repository to GitHub or GitLab.
2. Create a Cloudflare Pages project from the repository.
3. Set the build command, output directory, and `HUGO_VERSION`.
4. Deploy from the selected branch.
5. Add the custom domain in Cloudflare Pages.
6. Let Cloudflare create or validate the DNS record.
7. HTTPS is issued automatically after the domain points to Pages.

Future deployments happen through `git push`.
