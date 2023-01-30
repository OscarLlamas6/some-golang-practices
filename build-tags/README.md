# Go | Build Tags

```bash

# OR Rule
# // +build tag1 tag2
# To call build tags
> go run -tags tag1 .
> go run -tags tag2 .

# AND Rule
# // +build tag1,tag2
# To call build" tags
> go run -tags "tag1 tag2" .

# NOT Rule
# // +build !tag2

```