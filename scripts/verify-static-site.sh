#!/usr/bin/env sh
set -eu

hugo --minify

test -f public/index.html
test -f public/posts/index.html
test -f public/sitemap.xml
test -f public/robots.txt
test -f public/index.xml

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
grep -Eq 'name="?twitter:card"? content="?summary_large_image"?' public/posts/seo-static-blog/index.html
grep -q '<article' public/posts/seo-static-blog/index.html
grep -Eq 'rel="?author"?' public/posts/seo-static-blog/index.html
