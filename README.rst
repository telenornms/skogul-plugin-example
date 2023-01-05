Skogul length plugin demo
=========================

This repo is an example of a minimal working Skogul plugin built in an
external repo, loaded with 'skogul -experimental-plugins skogul-length.so'.

Note that go.mod has::

	replace github.com/telenornms/skogul => ../skogul

This was because this repo was written BEFORE plugins were supported, and I
needed to test locally. Remove this or check out skogul to ../skogul to
make things work.

This repo and README will be used to show-case the plugin integration over
time and is aimed at potential plugin-developers... But thus far this
support is in its infancy.
