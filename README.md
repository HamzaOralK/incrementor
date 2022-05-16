## Incrementor

Written very quickly to curb a need. Basically for my release plan I am creating branches as `0.1.x` and
I want every commit those branches to create incremented patch tags and eventually those tags should create
releases automatically. 

For this purpose I need to filter `0.1.x` tags from the repository and get the latest one and increment it
then create a tag and a release for this purpose. 

## Usage 

```shell
./incrementor  -prefix=release -separator=/ -owner=venture-justbuild -repository=ProductCatalogManager -branch=release/0.1.x
```

Prefix and separator are there if one is using those to denote branches but apart from it semver representation
must be in `number.number.x` format because incrementor is using x to create necessary regex to filter 
tags. 

Then the output can be used for the tag and release creation.

## Necessary Environment Variables
`GITHUB_TOKEN` must be given in order to get tags for the repository.
