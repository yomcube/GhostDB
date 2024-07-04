#!/usr/bin/env python3

from flask import Flask, make_response, send_from_directory

from ghostdb.common import *
from ghostdb.rkginfo import *

app = Flask(__name__, static_url_path='/s', static_folder='static')

@app.route('/')
def root():
	body = """
<p>GhostDB is a database of Mario Kart Wii ghost files.</p>
<ul id="ocul">
	<li><a href="/db/o/">Original Mario Kart Wii</a></li>
	<li><a href="/db/c/">Custom Track Distributions</a></li>
</ul>
"""
	
	htm = makehtml('Home', body, 'GhostDB')
	res = make_response(htm)
	return res

# RKG info
@app.route('/rkg/<path:filename>')
def rkginfopage(filename):
	if not filename.endswith('.rkg'):
		return make_response('Bad request')
	
	rkgd = {}
	with open('DBs/' + filename, 'rb') as f:
		by = f.read(0x88)
		if not by[0:4] == b'RKGD':
			return make_response('File is not a valid RKG')
		rkgd = rkginfo(by)
		body = genrkgbody(rkgd)
	return makehtml('Ghost Info', body)

@app.route('/db/o/')
def orig():
	body = '<div id="trackdiv">\n'
	
	i = 0
	while i < 4:
		body += '<div>\n'
		j = 0
		while j < 2:
			body += '<ul>\n'
			
			k = 0
			while k < 4:
				t = tracks[j][i][k]
				body += f'<li><a href="/db/o/{t[0]}/">{t[1]}</a></li>\n'
				k += 1
			body += '</ul><br>\n'
			j += 1
		body += '</div>\n'
		i += 1
	
	body += '</div>'
	return makehtml('Original Tracks', body)

@app.route('/db/o/<course>/')
def origcourse(course):
	return makeccpage('o', course, '150', trackdict)

@app.route('/db/o/<course>/<cc>/')
def origcoursecc(course, cc):
	return makeccpage('o', course, cc, trackdict)

"""@app.route('/db/c/')
def cust():
	pass"""








if __name__ == '__main__':
	app.run()