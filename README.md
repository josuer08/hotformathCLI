# hotformathCLI
This is a CLI version of the [HOTFORMATH](https://hotformath.com) project.

This tool is a google advanced index search helper, mainly intended at helping people find archives of videos, pictures, documents,
zips, or executables by scrapping google and using it advance search futures.

_this project does not intend to be a piracy tool, this is a tool that i maintain and use personally for research and as a medium_
_to get content that can be used under the fair use act_

# Install
For instalation you just need to download and run the most recent release directly form github releases tab.
## Compile from source
`go build hotformathCLI.go` should be enough

# How does this work?
This project uses [rocketlaunchr's google-search module](github.com/rocketlaunchr/google-search) to search google for a term introduced by the user and then scraps the first 3 results
the query sent to google shoul look a little something like `intext:"[YOUR SEARCH TERM]" -inurl:(jsp|pl|php|html|aspx|htm|cf|shtml) -inurl:(index_of|listen77|mp3raid|mp3toss|mp3drug|index_of|wallywashis) intitle:"index.of./" [specific modifier]`
where the specific modifier depends on the flag you use

| Type             | Flag | Modifier                                                     |
| ---------------- | ---- | ------------------------------------------------------------ |
| Videos           | -v   | `(avi\|mkv\|mov\|mp4\|mpg\|wmv)`                             |
| Audios           | -a   | `(ac3\|flac\|m4a\|mp3\|ogg\|wav\|wma)`                       |
| Ebooks           | -e   | `(CBZ\|CBR\|CHM\|DOC\|DOCX\|EPUB\|MOBI\|ODT\|PDF\|RTF\|txt)` |
| Pictures         | -p   | `(bmp\|gif\|jpg\|png\|psd\|tif\|tiff)`                       |
| Software         | -s   | `(apk\|exe\|iso\|rar\|tar\|zip)`                             |
| Compressed files | -c   | `(7z\|bz2\|gz\|iso\|rar\|zip)`                               |

You can read more about google's advance querying features [here](https://ahrefs.com/blog/google-advanced-search-operators/)

# Roamap
In the future we plan to:
  - [x] scrap from google the results
  - [ ] integrate with fzf
  - [ ] open links into browser or copy to clipboard directly
  - [ ] advanced manual search with other flags
  - [ ] TUI with https://github.com/charmbracelet/ repos

# Credits
![HotForMath](logo.png)
[HOTFORMATH ON GITHUB](https://github.com/marcosfermin/hotformath)
