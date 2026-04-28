# Arsen Gumin Blog

Static personal blog built with Hugo for Cloudflare Pages. It uses Markdown content, semantic HTML, one small local CSS file, no client-side JavaScript, automatic sitemap, robots.txt, RSS, canonical URLs, Open Graph, Twitter Cards, and BlogPosting JSON-LD.

## Local Development

```bash
hugo server
```

Open the local URL printed by Hugo. Drafts are excluded by default.

## Create a Post

```bash
hugo new posts/my-post.md
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
draft: false
---
```

Rules:

- keep `title` under 60 characters
- keep `description` under 160 characters
- put the main keyword in the title and first 100 words
- use one clear H1 from the page title, then H2/H3 sections
- link to at least one related post with descriptive anchor text
- add meaningful `alt` text to every image

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
- Match search intent in the introduction.
- Keep URLs readable: `/posts/my-post/`.
- Use semantic HTML through the Hugo templates: `main`, `article`, `header`, `nav`, `footer`.
- Add internal links between related posts.
- Keep published posts out of `draft: true`.
- Update `baseURL` in `config.toml` to the production domain before launch.

## Personal Brand

The design intentionally matches the existing one-page site: warm paper background, black border, mono typography, Georgia display headings, hard shadows, and the cluster badge. Keep the same voice in articles:

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
