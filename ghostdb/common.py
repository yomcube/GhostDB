from os import listdir
from os.path import isfile, join
import re

tracks = [
	[
		[
			[ "LC", "Luigi Circuit" ],
			[ "MMM", "Moo Moo Meadows" ],
			[ "MG", "Mushroom Gorge" ],
			[ "TF", "Toad's Factory" ],
		],
		[
			[ "MC", "Mario Circuit" ],
			[ "CM", "Coconut Mall" ],
			[ "DKSC", "DK's Snowboard Cross" ],
			[ "WGM", "Wario's Gold Mine" ],
		],
		[
			[ "DC", "Daisy Circuit" ],
			[ "KC", "Koopa Cape" ],
			[ "MT", "Maple Treeway" ],
			[ "GV", "Grumble Volcano" ],
		],
		[	
			[ "DDR", "Dry Dry Ruins" ],
			[ "MH", "Moonview Highway" ],
			[ "BC", "Bowser's Castle" ],
			[ "RR", "Rainbow Road" ],
		],
	],
	[
		[
			[ "rPB", "GCN Peach Beach" ],
			[ "rYF", "DS Yoshi Falls" ],
			[ "rGV2", "SNES Ghost Valley 2" ],
			[ "rMR", "N64 Mario Raceway" ],
		],
		[
			[ "rSL", "N64 Sherbet Land" ],
			[ "rSGB", "GBA Shy Guy Beach" ],
			[ "rDS", "DS Delfino Square" ],
			[ "rWS", "GCN Waluigi Stadium" ],
		],
		[
			[ "rDH", "DS Desert Hills" ],
			[ "rBC3", "GBA Bowser Castle 3" ],
			[ "rDKJP", "N64 DK's Jungle Parkway" ],
			[ "rMC", "GCN Mario Circuit" ],
		],
		[
			[ "rMC3", "SNES Mario Circuit 3" ],
			[ "rPG", "DS Peach Gardens" ],
			[ "rDKM", "GCN DK Mountain" ],
			[ "rBC", "N64 Bowser's Castle" ],
		]
	]
]
trackdict = {}
for g in tracks:
	for c in g:
		for t in c:
			trackdict[t[0]] = t[1]

ctmkfdict = {
	'0': '\u2460',
	'1': '\u2461',
	'2': '\u2462',
	'3': '\u2463',
	'4': '\u2464',
	'5': '\u2465',
	'6': '\u2466',
	'7': '\u2467',
	'8': '\u2468',
	'9': '\u2469',
	':': '\u246A',
	'.': '\u246B',
	'/': '\u246C',
	'-': '\u246D',
	'+': '\u246E',
}

pagetemplate = """<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>{0} | GhostDB</title>
		<link rel="stylesheet" href="/s/style.css">
	</head>
	<body>
		<!-- Header -->
		<div class="header">
			<a id="mainlink" href="/">GhostDB</a>
			<hr>
			<a href="/">Home</a> |
			<a href="/db/o/">Original Tracks</a> |
			<a href="/db/c/">Custom Track Distributions</a>
		</div>
		<div class="altheader">&nbsp;<hr>&nbsp;</div>
		<div class="flex-container">
			<div> <!-- Contain the page contents -->
				<h2 style="text-align: center;">{2}</h2>
				<!-- Code generated -->
{1}
			</div>
		</div>
	</body>
</html>"""

def makehtml(title, body, bodytitle = None):
	if bodytitle == None: bodytitle = title 
	return pagetemplate.format(title, body, bodytitle)

def makeccpage(p, course, cc, namedict = None):
	pat = f'{p}/{cc}/{course}'
	path = 'DBs/' + pat
	files = [f for f in listdir(path) if isfile(join(path, f))]
	c = course if namedict == None else namedict[course]
	if len(files) == 0:
		return makehtml(c, f'<p>No ghosts are available for {c}.</p>')
	
	body = '<ul id="trackul">\n'
	files.sort()
	for f in files:
		n = f.rstrip('.rkg')
		minutes = n[0:2].lstrip('0') or '0'
		time = f'{minutes}:{n[3:9]}'
		creator = n[13:]
		body += f'<li><a href="/rkg/{pat}/{f}">{time} by {creator}</a></li>'
	body += '</ul>'
	return makehtml(c, body)


def formattime(tobj):
	return f"{tobj['minutes']}:{tobj['seconds']}.{tobj['milliseconds']}"