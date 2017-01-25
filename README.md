# Feldküche

Scroll through your soup.io backups, made with [neingeist/soup-backup](https://github.com/neingeist/soup-backup).

## What?

This is a very crude program I made a few years ago and then never released it. If put it in a directory with a `backup.rss` file and an `enclosures/` folder, it will provide you with a simple interface with support for endless scrolling of all files in the backup.

## Help needed!

There are many shortcomings, and I'd be delighted if somebody wants to help with these:

* You have to click *somewhere* in order to load more images
* Older images do not disappear (I bet that makes the browser unhappy quite quickly)
* The backups do not contain videos (maybe one can patch/extend neingeist's script?)
* All images are in the original resolution — one could extend the program to resize pictures on the fly
* The page is ugly (maybe there is a way to export/import soup.io user settings?)


Please do not run it somewhere where it is accessible from the outside; It contains comments like
```go
// TODO: Is this secure?
	filename := r.URL.Path[len("/image/"):]
```
So, well, I bet there is a path traversal in it.

## License

For now I settled for GPLv3 — maybe this is a candidate for AGPL!
