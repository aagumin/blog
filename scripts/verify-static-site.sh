#!/usr/bin/env sh
set -eu

hugo --minify

test -f public/index.html
test -f public/posts/index.html
test -f public/sitemap.xml
test -f public/robots.txt
test -f public/index.xml
test -f public/_redirects
test -f NEWPOST.md
test -f archetypes/posts/index.md
test -f layouts/_default/_markup/render-image.html
test -f layouts/_default/_markup/render-link.html
test -f content/posts/seo-static-blog/index.md
test -f content/posts/seo-static-blog/cover.jpg
test -f content/posts/cloudflare-pages-hugo/index.md
test -f content/posts/cloudflare-pages-hugo/cover.jpg
test -f public/images/default-og.jpg
grep -q '/old-seo-static-blog /posts/seo-static-blog/ 301' public/_redirects
grep -q 'loading=lazy' public/posts/seo-static-blog/index.html
grep -q '<a href=https://gohugo.io/ rel="external noopener noreferrer"' public/posts/seo-static-blog/index.html

if ls content/posts/*.md >/dev/null 2>&1; then
  echo "Posts must use leaf page bundles: content/posts/<slug>/index.md"
  exit 1
fi

if grep -R "<script" public | grep -v 'application/ld+json'; then
  echo "Client-side JavaScript found"
  exit 1
fi

grep -Eq '<meta name="?robots"? content="?index,follow"?' public/index.html
grep -Eq '<meta name="?description"? content="?[^">]+' public/index.html
grep -Eq '<link rel="?canonical"?' public/posts/seo-static-blog/index.html
grep -q 'application/ld+json' public/posts/seo-static-blog/index.html
grep -q '"@context":"https://schema.org"' public/posts/seo-static-blog/index.html
grep -q '"@type":"BlogPosting"' public/posts/seo-static-blog/index.html
grep -Eq 'property="?og:title"?' public/posts/seo-static-blog/index.html
grep -Eq 'property="?og:type"? content="?article"?' public/posts/seo-static-blog/index.html
grep -Eq 'property="?og:image"? content="?https://blog.example.com/posts/seo-static-blog/cover' public/posts/seo-static-blog/index.html
grep -Eq 'property="?og:image:width"? content="?1200"?' public/posts/seo-static-blog/index.html
grep -Eq 'property="?og:image:height"? content="?630"?' public/posts/seo-static-blog/index.html
grep -Eq 'name="?twitter:card"? content="?summary_large_image"?' public/posts/seo-static-blog/index.html
grep -Eq 'name="?twitter:image"? content="?https://blog.example.com/posts/seo-static-blog/cover' public/posts/seo-static-blog/index.html
grep -Eq 'property="?og:image"? content="?https://blog.example.com/images/default-og.jpg"?' public/about/index.html
grep -Eq 'name="?twitter:image"? content="?https://blog.example.com/images/default-og.jpg"?' public/about/index.html
grep -q '<article' public/posts/seo-static-blog/index.html
grep -Eq 'rel="?author"?' public/posts/seo-static-blog/index.html

if grep -R -E 'CLUSTER|A Player|brand-name|github.com/aagumin|t.me/arsengumin|mailto:gumin@live.ru' public/*.html public/posts public/about; then
  echo "Duplicated onepager identity/social elements found in blog HTML"
  exit 1
fi

if grep -Eq 'Technical blog|Cloud-native ML/AI, data, and platform engineering' public/index.html; then
  echo "Oversized home hero copy found"
  exit 1
fi

grep -q 'Notes on systems, data, and AI infrastructure.' public/index.html
grep -q 'width:min(1012px,100%)' public/styles.css
grep -Eq '(&copy;|©) 2026 Arsen Gumin\. All rights reserved\.' public/index.html
grep -q 'hugo new posts/my-post/index.md' NEWPOST.md
grep -q './scripts/verify-posts.sh' NEWPOST.md
grep -q 'cover.jpg' NEWPOST.md
if grep -Eq '<footer[^>]*.*No JavaScript\. Small by default\.|<footer[^>]*.*href=/index.xml' public/index.html; then
  echo "Outdated footer copy or RSS footer link found"
  exit 1
fi

GOCACHE="${GOCACHE:-$(pwd)/.gocache}" go run scripts/verify-posts.go
