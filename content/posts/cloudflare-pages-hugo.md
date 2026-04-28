---
title: "Deploy Hugo Blog to Cloudflare Pages"
date: 2026-04-27
description: "Deploy a Hugo blog to Cloudflare Pages with git push, HTTPS, custom domains, and production build settings."
tags:
  - Cloudflare Pages
  - Hugo
  - Deployment
draft: false
slug: "cloudflare-pages-hugo"
---

Deploy Hugo blog to Cloudflare Pages when you want a fast static publishing workflow with global delivery, automatic HTTPS, and no backend to maintain. Hugo generates the files, Cloudflare Pages serves them, and git push becomes the deployment trigger.

This approach pairs well with a technical personal brand because the publishing pipeline stays boring. The energy goes into writing useful articles instead of operating infrastructure.

## Production Build Settings

Use these Cloudflare Pages settings:

- build command: `hugo --minify`
- output directory: `public`
- environment variable: `HUGO_VERSION=0.161.0`

The version should match local development. Pinning it avoids subtle rendering differences between your machine and the Cloudflare build environment.

## Domain and HTTPS

After the first deployment, add a custom domain in the Cloudflare Pages dashboard. Cloudflare will guide you through DNS records and issue HTTPS automatically once the domain points to the Pages project.

Keep the same domain in `baseURL` so canonical URLs, RSS links, sitemap entries, and social previews all point to the production site.

## Publishing Workflow

Write the post in Markdown, check it locally with `hugo server`, then push to the connected branch. Cloudflare Pages builds the site and publishes the new static files.

For more on the template structure, read the [SEO static blog architecture](/posts/seo-static-blog/).
