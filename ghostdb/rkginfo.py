from ghostdb.common import *

def rkginfo(by):
	fti = int.from_bytes(by[4:7], 'big')
	ftmi = (fti >> 17) & 0x7F
	ftse = (fti >> 10) & 0x7F
	ftms = fti & 0x3FF
	
	d = {
		"finishtime": {
			"minutes": ftmi,
			"seconds": ftse,
			"milliseconds": ftms,
		},
	}
	return d

def genrkgbody(d):
	finishtime = formattime(d['finishtime'])
	body = f'<span class="ctmkf" style="color:yellow">{finishtime}</span>'
	return body